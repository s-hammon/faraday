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
	var (
		groupFieldIdx int
		groupType     reflect.Type
		groupSchema   map[string]int
		activeGroup   reflect.Value
		hasGroup      bool
	)

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

			if field.Type.Kind() == reflect.Slice && field.Type.Elem().Kind() == reflect.Struct {
				groupFieldIdx = i
				groupType = field.Type.Elem()
				groupSchema = make(map[string]int)
				for j := range groupType.NumField() {
					sf := groupType.Field(j)
					if seg := exportSegmentName(sf); seg != "" {
						groupSchema[seg] = j
					}
				}
				hasGroup = true
			}
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

	var groupSlice reflect.Value
	if hasGroup {
		groupSlice = reflect.MakeSlice(elem.Field(groupFieldIdx).Type(), 0, 0)
	}

	for scanner.Scan() {
		data := scanner.Bytes()
		if len(data) < 3 {
			continue
		}
		name := string(data[:3])

		if hasGroup {
			_, isGroupSeg := groupSchema[name]
			if isGroupSeg {
				if name == exportSegmentName(groupType.Field(0)) && activeGroup.IsValid() {
					groupSlice = reflect.Append(groupSlice, activeGroup)
					activeGroup = reflect.Value{}
				}
				if !activeGroup.IsValid() {
					activeGroup = reflect.New(groupType).Elem()
				}
				idx := groupSchema[name]
				if err := decodeSegmentInto(activeGroup.Field(idx), groupType.Field(idx), data[4:], dec.delims); err != nil {
					return fmt.Errorf("decode group %s: %w", name, err)
				}
				continue
			}

			if activeGroup.IsValid() {
				groupSlice = reflect.Append(groupSlice, activeGroup)
				activeGroup = reflect.Value{}
			}
		}

		idx, ok := dec.segmentSchema[name]
		if !ok {
			continue
		}

		field := elem.Field(idx)
		if err := decodeSegmentInto(field, elem.Type().Field(idx), data[4:], dec.delims); err != nil {
			return fmt.Errorf("decode top-level %s: %w", name, err)
		}
	}
	if hasGroup && activeGroup.IsValid() {
		groupSlice = reflect.Append(groupSlice, activeGroup)
	}
	if hasGroup {
		elem.Field(groupFieldIdx).Set(groupSlice)
	}
	return nil
}

func decodeSegmentInto(field reflect.Value, sf reflect.StructField, raw []byte, delims delimiters) error {
	isSlice := field.Kind() == reflect.Slice
	typ := field.Type()
	if isSlice {
		typ = typ.Elem()
	}

	segVal := reflect.New(typ).Elem()
	rawFields := bytes.Split(raw, []byte{delims.field})

	for i := range min(segVal.NumField(), len(rawFields)) {
		fVal := segVal.Field(i)
		spec := NewFieldSpec(uint8(i+1), fVal)
		spec.ParseTag(segVal.Type().Field(i).Tag.Get("hl7"))
		if err := spec.parse(rawFields[i], delims.toSlice()); err != nil {
			return err
		}
		fVal.Set(spec.Val)
	}

	if isSlice {
		field.Set(reflect.Append(field, segVal))
	} else {
		field.Set(segVal)
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
