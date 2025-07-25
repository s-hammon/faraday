package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"reflect"

	"github.com/s-hammon/p"
)

type MSHUnmarshaller interface {
	UnmarshalHeader([]byte) error
}

type Decoder struct {
	r       io.Reader
	scanned int64
	err     error

	parsedMSH     bool
	segmentSchema map[string]int // map of segment name to struct field index
	delims        delimiters
}

type delimiters struct {
	field        byte
	component    byte
	repeat       byte
	escape       byte
	subcomponent byte
}

func (d delimiters) toSlice() []byte {
	return []byte{
		d.component,
		d.repeat,
		d.escape,
		d.subcomponent,
	}
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r:         r,
		parsedMSH: false,
	}
}

func (dec *Decoder) Decode(val any) error {
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Pointer || v.IsNil() {
		return fmt.Errorf("Decode: expected non-nil pointer, got %T", val)
	}

	elem := v.Elem()
	// TODO: also implement support for map[string]any
	if elem.Kind() != reflect.Struct {
		return fmt.Errorf("Decoder: not a pointer to struct (got %T)", val)
	}

	dec.segmentSchema = make(map[string]int)
	for i := range elem.NumField() {
		field := elem.Type().Field(i)
		if name := exportSegmentName(field); name != "" {
			dec.segmentSchema[name] = i
		}
	}

	scanner := bufio.NewScanner(dec.r)
	scanner.Split(SegmentSplitter('\r'))

	if !scanner.Scan() {
		return io.EOF
	}
	header := scanner.Bytes()

	if len(header) < 8 || string(header[:3]) != "MSH" {
		return fmt.Errorf("Decode: expected first segment to be MSH")
	}

	fieldSep := header[3]
	encChars := header[4:8]

	dec.delims = delimiters{
		field:        fieldSep,
		component:    encChars[0],
		repeat:       encChars[1],
		escape:       encChars[2],
		subcomponent: encChars[3],
	}

	if idx, ok := dec.segmentSchema["MSH"]; ok {
		field := elem.Field(idx)
		ptr := reflect.New(field.Type())

		u, ok := ptr.Interface().(MSHUnmarshaller)
		if !ok {
			return fmt.Errorf("MSH field type %s does not implement UnmarshalHeader([]byte)", field)
		}
		if err := u.UnmarshalHeader(header[3:]); err != nil {
			return fmt.Errorf("MSH.UnmarshalHeader: %w", err)
		}
		field.Set(ptr.Elem())
	}

	for scanner.Scan() {
		data := scanner.Bytes()
		if len(data) < 3 {
			if len(data) == 0 {
				break
			}
			return fmt.Errorf("malformed segment: %q", data)
		}

		name := string(data[:3])
		idx, ok := dec.segmentSchema[name]
		if !ok {
			continue
		}

		field := elem.Field(idx)
		isSlice := field.Kind() == reflect.Slice
		typ := field.Type()
		if isSlice {
			typ = typ.Elem()
		}

		segVal := reflect.New(typ).Elem()
		rawFields := bytes.Split(data[4:], []byte{dec.delims.field})
		for i := range min(segVal.NumField(), len(rawFields)) {
			structField := segVal.Type().Field(i)
			fVal := segVal.Field(i)

			spec := NewFieldSpec(uint8(i+1), fVal)
			spec.ParseTag(structField.Tag.Get("hl7"))
			if err := spec.parse(rawFields[spec.Position-1], dec.delims.toSlice()); err != nil {
				return fmt.Errorf("%s field %d: %w", name, i+1, err)
			}
			fVal.Set(spec.Val)
		}

		if isSlice {
			field.Set(reflect.Append(field, segVal))
		} else {
			field.Set(segVal)
		}
	}
	return nil
}

func exportSegmentName(field reflect.StructField) string {
	name := p.Coalesce(field.Tag.Get("hl7"), field.Name)
	if _, ok := SegmentTypes[name]; ok {
		return name
	}
	return ""
}

func SegmentSplitter(delim byte) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if i := bytes.IndexByte(data, delim); i >= 0 {
			return i + 1, data[:i], nil
		}
		if atEOF && len(data) > 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}
