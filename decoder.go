package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"reflect"

	"github.com/s-hammon/p"
)

type Decoder struct {
	r       io.Reader
	scanned int64
	err     error

	parsedMSH     bool
	segmentSchema map[string]int // map of segment name to struct field index
	delimiters    []byte
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r:             r,
		parsedMSH:     false,
		segmentSchema: make(map[string]int),
	}
}

func (dec *Decoder) Decode(val any) error {
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Pointer || v.IsNil() {
		return fmt.Errorf("Decoder.Decode(%T): non-pointer", val)
	}

	elem := v.Elem()
	// TODO: also implement support for map[string]any
	if elem.Kind() != reflect.Struct {
		return fmt.Errorf("Decoder.Decode(%T): not a struct", val)
	}

	for i := range elem.NumField() {
		t := elem.Type()
		if t.Kind() != reflect.Struct {
			continue
		}
		name := exportSegmentName(t.Field(i))
		if name != "" {
			dec.segmentSchema[name] = i
		}
	}

	scanner := bufio.NewScanner(dec.r)
	scanner.Split(SegmentSplitter('\r'))

	for scanner.Scan() {
		data := scanner.Bytes()
		// first line should always be the message header segment (MSH)
		if !dec.parsedMSH {
			if len(data) < 8 {
				return fmt.Errorf("malformed message data: must be at least 8 bytes")
			}
			name := string(data[:3])
			if name != "MSH" {
				return fmt.Errorf("invalid segment name '%s' (expecting 'MSH')", name)
			}

			idx, ok := dec.segmentSchema[name]
			if ok {
				segment := elem.Field(idx)
				if msh, ok := segment.Interface().(MSH); ok {
					if err := msh.UnmarshalHL7(data[3:]); err != nil {
						return fmt.Errorf("MSH.UnmarshalHL7: %v", err)
					}
					segment.Set(reflect.ValueOf(msh))
				} // else {
				// TODO: deserialize non-standard MSH
				// for i := range segment.NumField() {
				// 	...
				// }
			}
			dec.delimiters = data[3:8]
			dec.parsedMSH = true
			continue
		}
		if len(data) < 3 {
			if len(data) == 0 {
				return io.EOF
			} else {
				return fmt.Errorf("malformed segment: %s", string(data))
			}
		}
		name := string(data[:3])

		idx, ok := dec.segmentSchema[name]
		if !ok || idx > 254 {
			// either unexported/unsupported struct field or unwanted segment
			// ....or index is too high
			// TODO: some ORUs/MDMs may genuinely be very, very long
			// check if any do exceed 255 lines
			continue
		}
		segment := elem.Field(idx)
		fields := bytes.Split(data[4:], dec.delimiters[:1])
		for i := range min(segment.NumField(), len(fields)) {
			if len(fields[i]) == 0 {
				continue
			}
			fVal := segment.Field(i)
			spec := NewFieldSpec(uint8(i+1), fVal)
			spec.ParseTag(segment.Type().Field(i).Tag.Get("hl7"))
			if err := spec.parse(fields[i], dec.delimiters); err != nil {
				return fmt.Errorf("FieldSpec.parse: %v", err)
			}
			fVal.Set(spec.Val)
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
