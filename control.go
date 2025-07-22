package main

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
	MessageType                   CM `hl7:"opt=R"`
	MessageControlId              ST `hl7:"opt=R"`
	ProcessingId                  PT `hl7:"opt=R"`
	VersionId                     ID `hl7:"opt=R,tbl=0104"`
	SequenceNumber                NM
	ContinuationPointer           ST
	AcceptAcknowledgmentType      ID `hl7:"tbl=0155"`
	ApplicationAcknowledgmentType ID `hl7:"tbl=0155"`
	CountryCode                   ID
	CharacterSet                  ID `hl7:"rep=Y3,tbl=0211"`
	PrincipalLanguage             CE
}
