package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/s-hammon/p"
)

type FieldSpec struct {
	Position     uint8
	Typ          reflect.Type
	Val          reflect.Value
	Optionality  optionality
	Repeats      bool
	RepeatCount  uint8
	ControlTable *ControlTable

	validationErr error
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

// TODO: rework this a bit to support logging for whenever tag key was found,
// but could not parse the value (or it is not supported)
func (spec *FieldSpec) ParseTag(tag string) *FieldSpec {
	for pair := range strings.SplitSeq(tag, ",") {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) == 2 {
			updateSpec(spec, parts...)
		}
	}
	return spec
}

func (spec *FieldSpec) validate(v []byte, delimiters []byte) {
	if len(v) == 0 && spec.Optionality == Required {
		spec.validationErr = fmt.Errorf("no data provided for required field at position %d", spec.Position)
		return
	}

	val := reflect.New(spec.Typ).Elem()
	if val.Kind() == reflect.Struct {
		i := 0
		for component := range bytes.SplitSeq(v, delimiters[:1]) {
			if i >= val.NumField() {
				i++
				break
			}
			fVal := val.Field(i)
			if fVal.Kind() == reflect.Struct {
				j := 0
				for subcomponent := range bytes.SplitSeq(component, delimiters[3:]) {
					if j > fVal.NumField() {
						spec.validationErr = fmt.Errorf("expected max %d subcomponents for field number %d", fVal.NumField(), spec.Position)
						break
					}
					sfVal := fVal.Field(j)
					sfVal.SetString(string(subcomponent))
					j++
				}
			} else {
				fVal.SetString(string(component))
			}
			i++
		}
		if i > val.NumField() && spec.validationErr == nil {
			spec.validationErr = fmt.Errorf("expected max %d components for field number %d", val.NumField(), spec.Position)
		}
	} else if val.Kind() == reflect.String {
		val.SetString(string(v))
	}
	if !p.IsZero(val) {
		spec.Val = val
	}
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
		spec.Repeats = parts[1][:1] == "Y"
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

type SegmentSpec map[string]*FieldSpec
