// TODO: rework this a bit to support logging for whenever tag key was found,
// but could not parse the value (or it is not supported)
package main

import (
	"reflect"
	"strconv"
	"strings"
)

type FieldSpec struct {
	Position     uint8
	Typ          reflect.Type
	Optionality  optionality
	Repeats      bool
	RepeatCount  uint8
	ControlTable *ControlTable
}

// Creates a new FieldSpec with default values
func NewFieldSpec(pos uint8, typ reflect.Type) *FieldSpec {
	return &FieldSpec{
		Position:     pos,
		Typ:          typ,
		Optionality:  Optional,
		Repeats:      false,
		RepeatCount:  0,
		ControlTable: nil,
	}
}

func (spec *FieldSpec) ParseTag(tag string) *FieldSpec {
	for pair := range strings.SplitSeq(tag, ",") {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) == 2 {
			updateSpec(spec, parts...)
		}
	}
	return spec
}

func updateSpec(spec *FieldSpec, parts ...string) {
	if parts[1] == "" {
		return
	}

	switch parts[0] {
	default:
		return
	case "pos":
		n, err := strconv.Atoi(parts[1])
		if err != nil || !canInt8(n) {
			return
		}
		spec.Position = uint8(n)
	case "opt":
		spec.Optionality = fromString(parts[1])
	case "rep":
		spec.Repeats = parts[1][:0] == "Y"
		if len(parts[1]) > 1 {
			n, err := strconv.Atoi(parts[1][1:])
			if err != nil || !canInt8(n) {
				spec.RepeatCount = 1
			} else {
				spec.RepeatCount = uint8(n)
			}
		}
	case "tbl":
		spec.ControlTable = TableMap[parts[1]]
	}
}
