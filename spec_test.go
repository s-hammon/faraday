package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

var defaultDelims = []byte("^~\\&")

func TestFieldSpecValidate(t *testing.T) {
	spec := NewFieldSpec(0, reflect.TypeOf(ST("")))
	spec.validate([]byte("pass"), nil)
	require.NoError(t, spec.validationErr)
	require.Equal(t, ST("pass"), spec.Val.Interface().(ST))

	spec = NewFieldSpec(0, reflect.TypeOf(CM{}))
	spec.validate([]byte(""), defaultDelims)
	require.NoError(t, spec.validationErr)

	spec = NewFieldSpec(0, reflect.TypeOf(CM{}))
	spec.Optionality = Required
	spec.validate([]byte(""), defaultDelims)
	require.Error(t, spec.validationErr)

	spec = NewFieldSpec(0, reflect.TypeOf(CM{}))
	spec.validate([]byte("pass"), defaultDelims)
	require.NoError(t, spec.validationErr)
	require.Equal(t, CM{Type: "pass"}, spec.Val.Interface().(CM))

	spec = NewFieldSpec(0, reflect.TypeOf(CM{}))
	spec.validate([]byte("ADT^A01"), defaultDelims)
	require.NoError(t, spec.validationErr)
	require.Equal(t, CM{Type: "ADT", Event: "A01"}, spec.Val.Interface().(CM))

	spec = NewFieldSpec(0, reflect.TypeOf(CQ{}))
	spec.validate([]byte("1234^ABC&millimeters&MEDITECH"), defaultDelims)
	require.NoError(t, spec.validationErr)
	require.Equal(t,
		CQ{
			Quantity: "1234",
			Units: CE{
				Identifier:   "ABC",
				Text:         "millimeters",
				CodingSystem: "MEDITECH",
			},
		},
		spec.Val.Interface().(CQ),
	)

	spec = NewFieldSpec(0, reflect.TypeOf(HD{}))
	spec.validate([]byte("this^has^too^many^components"), defaultDelims)
	require.Error(t, spec.validationErr)
	require.Equal(t,
		HD{
			NamespaceId:     "this",
			UniversalId:     "has",
			UniversalIdType: "too",
		},
		spec.Val.Interface().(HD),
	)
}
