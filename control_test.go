package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMSHUnmarshalHL7_Basic(t *testing.T) {
	raw := []byte("MSH|^~\\&|SendingApp|SendingFac|ReceivingApp|ReceivingFac|202507231230||ADT^A01|MSG00001|P|2.5|||||USA|ASCII")

	var msh MSH
	err := msh.UnmarshalHeader(raw[3:])
	require.NoError(t, err)
	require.Equal(t, "|", string(msh.FieldSeparator))
	require.Equal(t, "^~\\&", string(msh.EncodingCharacters))
	require.Equal(t, "SendingApp", string(msh.SendingApplication.NamespaceId))
	require.Equal(t, "SendingFac", string(msh.SendingFacility.NamespaceId))
	require.Equal(t, "ReceivingApp", string(msh.ReceivingApplication.NamespaceId))
	require.Equal(t, "ReceivingFac", string(msh.ReceivingFacility.NamespaceId))
	require.Equal(t, "202507231230", string(msh.DateTime))
	require.Equal(t, "", string(msh.Security))
	require.Equal(t, "ADT", string(msh.MessageType.Type))
	require.Equal(t, "A01", string(msh.MessageType.Event))
	require.Equal(t, "MSG00001", string(msh.MessageControlId))
	require.Equal(t, "P", string(msh.ProcessingId.ProcessingId))
	require.Equal(t, "2.5", string(msh.VersionId))
	require.Equal(t, "", string(msh.SequenceNumber))
	require.Equal(t, "", string(msh.ContinuationPointer))
	require.Equal(t, "", string(msh.AcceptAcknowledgmentType))
	require.Equal(t, "", string(msh.ApplicationAcknowledgmentType))
	require.Equal(t, "USA", string(msh.CountryCode))
	require.Equal(t, "ASCII", string(msh.CharacterSet))
}
