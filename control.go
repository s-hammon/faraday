/*
This module contains the standard for control messages (MSH, ACK, MSA, etc).
*/
package faraday

import (
	"bytes"
	"fmt"
	"reflect"
)

// The standard MSH segment
/*
A struct field = a segment field, and we've defined a Go type for each HL7
type. The remaining specifications for a field is handled as follows:

	Position: `pos` tag (default = n+1 position in struct)
	Optionality: `opt` tag (default = O)
	Repetition: `rep` tag (default = N)
	Table: `tbl` tab (default = nil)

Tags:
	opt: R , O, C, X, B		opt=R
	tbl: 0104, 0155, etc		tbl=0104
	rep: Y or N with number		rep=Y
	pos: 1, 2, 3, etc.
*/
type MSH struct {
	FieldSeparator                ST `hl7:"opt=R"`
	EncodingCharacters            ST `hl7:"opt=R"`
	SendingApplication            HD
	SendingFacility               HD
	ReceivingApplication          HD
	ReceivingFacility             HD
	DateTime                      TS
	Security                      ST
	MessageType                   CM_MSG `hl7:"opt=R"`
	MessageControlId              ST     `hl7:"opt=R"`
	ProcessingId                  PT     `hl7:"opt=R"`
	VersionId                     ID     `hl7:"opt=R,tbl=0104"`
	SequenceNumber                NM
	ContinuationPointer           ST
	AcceptAcknowledgmentType      ID `hl7:"tbl=0155"`
	ApplicationAcknowledgmentType ID `hl7:"tbl=0155"`
	CountryCode                   ID
	CharacterSet                  ID `hl7:"rep=Y3,tbl=0211"`
	PrincipalLanguage             CE
}

func (seg *MSH) UnmarshalHeader(b []byte) error {
	if len(b) < 6 {
		return fmt.Errorf("input '%s' too short--must be at least 6 bytes", string(b))
	}
	seg.FieldSeparator = ST(b[:1])
	seg.EncodingCharacters = ST(b[1:5])

	v := reflect.ValueOf(seg).Elem()
	t := v.Type()

	fields := bytes.Split(b[6:], seg.fieldSeparator())
	for i, j := 0, 2; i < len(fields) && j < v.NumField(); i, j = i+1, j+1 {
		fVal := v.Field(j)
		fTyp := t.Field(j).Type

		raw := fields[i]
		switch fTyp {
		case reflect.TypeOf(ST("")):
			fVal.SetString(string(raw))
		case reflect.TypeOf(HD{}):
			parts := bytes.Split(raw, seg.componentSeparator())
			fVal.Set(reflect.ValueOf(HD{
				NamespaceId:     IS(partsSafe(parts, 0)),
				UniversalId:     ST(partsSafe(parts, 1)),
				UniversalIdType: ID(partsSafe(parts, 2)),
			}))
		case reflect.TypeOf(TS("")):
			fVal.Set(reflect.ValueOf(TS(raw)))
		case reflect.TypeOf(ID("")):
			fVal.Set(reflect.ValueOf(ID(raw)))
		case reflect.TypeOf(NM("")):
			fVal.Set(reflect.ValueOf(NM(raw)))
		case reflect.TypeOf(CE{}):
			parts := bytes.Split(raw, seg.componentSeparator())
			fVal.Set(reflect.ValueOf(CE{
				Identifier:            ST(partsSafe(parts, 0)),
				Text:                  ST(partsSafe(parts, 1)),
				CodingSystem:          ST(partsSafe(parts, 2)),
				AlternateIdentifier:   ST(partsSafe(parts, 3)),
				AlternateText:         ST(partsSafe(parts, 4)),
				AlternateCodingSystem: ST(partsSafe(parts, 5)),
			}))
		case reflect.TypeOf(CM_MSG{}):
			parts := bytes.Split(raw, seg.componentSeparator())
			fVal.Set(reflect.ValueOf(CM_MSG{
				Type:  ID(partsSafe(parts, 0)),
				Event: ID(partsSafe(parts, 1)),
			}))
		case reflect.TypeOf(PT{}):
			parts := bytes.Split(raw, seg.componentSeparator())
			fVal.Set(reflect.ValueOf(PT{
				ProcessingId:   ID(partsSafe(parts, 0)),
				ProcessingMode: ID(partsSafe(parts, 1)),
			}))
		default:
		}
	}
	return nil
}

func (seg *MSH) fieldSeparator() []byte {
	return []byte(seg.FieldSeparator)
}

func (seg *MSH) componentSeparator() []byte {
	return []byte(seg.EncodingCharacters[:1])
}

func partsSafe(parts [][]byte, i int) string {
	if i < len(parts) {
		return string(parts[i])
	}
	return ""
}

// The standard NTE segment
type NTE struct {
	SetId           SI
	SourceOfComment ID
	Comment         FT `hl7:"rep=Y"`
}
