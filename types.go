package main

import (
	"strconv"

	"github.com/s-hammon/p"
)

// these are the basic HL7 "types"

/*
	ALPHANUMERIC
*/

// String
// TODO: technically only includes hexadecimal 20 - 7E (ASCII 32 - 126) except
// for those defined in MSH.2 (applies to ST, TX, FT)
type ST string

// Text
// TODO: implement method to limit this to 64kb in length
type TX string

// Formatted Text
// TODO: implement method to limit this to 64kb in length
type FT string

/*
	NUMERICAL
*/

// Numeric
// TODO: implement ToNumeric
// parse +/-123.792
type NM string

// Sequence ID (integer)
type SI string

func (s SI) Int() int {
	i, err := strconv.Atoi(string(s))
	if err != nil {
		panic(p.Format("non-numeric value for SI: %v", s))
	}
	return i
}

// Composite quantity
type CQ struct {
	Quantity NM
	Units    CE
}

// Money
// TODO: MSH.17 country code determines denomination if not provided, so need
// to handle a cross walk & default if MSH.17 isn't specified (USA -> USD)
type MO struct {
	Quantity     NM
	Denomination ID // ISO 4217
}

// Composite Price
type CP struct {
	Price      MO
	PriceType  ID // HL7 0205
	FromValue  NM
	ToValue    NM
	RangeUnits CE
	RangeType  ID // HL7 0298
}

type comparator string

const (
	GT comparator = ">"
	LT comparator = "<"
	GE comparator = ">="
	LE comparator = "<="
	EQ comparator = "="
	NE comparator = "<>"
)

type separator string

const (
	S separator = "-"
	A separator = "+"
	D separator = "/"
	R separator = ":"
)

// Structured Numeric
type SN struct {
	Comparator comparator
	FirstNum   NM
	SecondNum  NM
	Separator  ST
}

/*
	IDENTIFIER
*/

// Coded value for user-defined tables
type IS string

// Coded value for HL7 tables
type ID string

// Hierarchic Designator
type HD struct {
	NamespaceId     IS
	UniversalId     ST
	UniversalIdType ID // HL7 0301
}

// Entity Identifier
type EI struct {
	EntityIdentifier ST
	NamespaceId      IS
	UniversalId      ST
}

// Reference Pointer
type RP struct {
	Pointer       ST
	ApplicationId HD
	DataType      ID // HL7 0191
	DataSubType   ID // HL7 0191
}

// Person Location
type PL struct {
	PointOfCare         IS
	Room                IS
	Bed                 IS
	Facility            HD
	LocationStatus      IS
	PersonLocationType  IS
	Building            IS
	Floor               IS
	LocationDescription ST
}

// Processing Type
type PT struct {
	ProcessingId   ID // HL7 0103
	ProcessingMode ID // HL7 0207
}

/*
	DATETIME
*/

// Date YYYY[MM[dd]]
type DT string

// Time HH[MM[SS[.S[S[S[S]]]]]][+/-ZZZZ]
type TM string

// Timestamp YYYY[MM[dd]]HH[MM[SS[.S[S[S[S]]]]]][+/-ZZZZ]
type TS string

// TODO: implement ToDatetime

/*
	CODE VALUES
*/

// Coded Element
type CE struct {
	Identifier            ST
	Text                  ST
	CodingSystem          ST
	AlternateIdentifier   ST
	AlternateText         ST
	AlternateCodingSystem ST
}

// Coded Element with formatted values
type CF struct {
	Identifier            ID // HL7 0203
	FormattedText         FT
	CodingSystem          ST
	AlternateIdentifier   ID // HL7 0203
	AlternateText         FT
	AlternateCodingSystem ST
}

// Composite ID with check digit
type CK struct {
	IdNumber             NM
	CheckDigit           NM
	CheckDigitSchemeCode ID // HL7 0061
	AssigningAuthority   HD
}

// Composite ID number and name
type CN struct {
	IdNumber           ST
	FamilyName         ST
	GivenName          ST
	MiddleName         ST
	Suffix             ST
	Prefix             ST
	Degree             ST
	SourceTable        IS
	AssigningAuthority HD
}

// Extended Composite ID with check digit
type CX struct {
	IdNumber             ST
	CheckDigit           ST
	CheckDigitSchemeCode ID // HL7 0061
	AssigningAuthority   HD
	IdentifierTypeCode   IS
	AssigningFacility    HD
}

// Extended Composite ID number and name
type XCN struct {
	IdNumber             ST
	FamilyName           ST
	GivenName            ST
	MiddleName           ST
	Suffix               ST
	Prefix               ST
	Degree               ST
	SourceTable          IS
	AssigningAuthority   HD
	NameTypeCode         ID // HL7 0200
	IdentifierCheckDigit ST
	CheckDigitSchemeCode ID // HL7 0061
	IdentifierTypeCode   IS
	AssigningFacility    HD
}

// Composite -- apparently they aren't allowed
type CM struct {
	Type  ID // HL7 0076
	Event ID // HL7 0003
}

/*
	DEMOGRAPHICS
*/

// Address
type AD struct {
	StreetAddress              ST
	OtherDesignation           ST
	City                       ST
	State                      ST
	Zip                        ST
	Country                    ID // ISO 3166
	AddressType                ID // HL7 0190
	OtherGeographicDesignation ST
}

// Person Name
type PN struct {
	FamilyName ST
	GivenName  ST
	MiddleName ST
	Suffix     ST
	Prefix     ST
	Degree     ST
}

// Telephone number [NN][(999)]999-9999[X99999][B99999]
type TN string

// Extended Address
type XAD struct {
	StreetAddress              ST
	OtherDesignation           ST
	City                       ST
	State                      ST
	Zip                        ST
	Country                    ID // ISO 3166
	AddressType                ID // HL7 0190
	OtherGeographicDesignation ST
	CountyCode                 IS
	CensusTract                IS
}

// Extended Person Name
type XPN struct {
	FamilyName             ST
	GivenName              ST
	MiddleName             ST
	Suffix                 ST
	Prefix                 ST
	Degree                 ST
	NameTypeCode           ID // HL7 0200
	NameRepresentationCode ID // HL7 4000
}

// Extended Composite Name and ID number for organizations
type XON struct {
	OrganizationName     ST
	TypeCode             IS
	IdNumber             NM
	CheckDigit           NM
	CheckDigitSchemeCode ID // HL7 0061
	AssigningAuthority   HD
	IdentifierTypeCode   IS
	AssigningFacility    HD
}

// Extended Telecommunication Number
type XTN struct {
	Number                   TN
	TelecommunicationUseCode ID // HL7 0201
	EquipmentType            ID // HL7 0202
	EmailAddress             ST
	CountryCode              NM
	AreaCode                 NM
	PhoneNumber              NM
	Extension                NM
	AnyText                  ST
}

/*
	EXTENDED QUERIES
*/

// Query Input Parameter
// TODO: if this ever comes up, the convention for the field name is such that
// it begints with '@' (must be escaped if it is being used as a delimiter as
// specified in the MSH.2 field). Also, the Values field could contain
// subcomponent values (e.g "value1&value2&value3") or a single value.
type QIP struct {
	FieldName ST
	Values    ST
}

// Query Selection Criteria
type QSC struct {
	FieldName             ST
	RelationalOperator    ID // HL7 0209
	Value                 ST
	RelationalConjunction ID // HL7 0102
}

// Row Column Definition
type RCD struct {
	Hl7ItemNumber      ST
	Hl7DataType        ST
	MaximumColumnWidth NM
}

/*
	MASTER FILES
*/

// Driver's License Number
type DLN struct {
	LicenseNumber  ST
	IssuingState   IS
	ExpirationDate DT
}

// Job Code/Class
type JCC struct {
	JobCode  IS
	JobClass IS
}

// Visiting hours
type VH struct {
	StartDayRange  ID // HL7 0267
	EndDayRange    ID // HL7 0267
	StartHourRange TM
	EndHourRange   TM
}

/*
	MEDICAL RECORDS/INFORMATION MANAGEMENT
*/

// Performing Person time stamp
type PPN struct {
	FamilyName              ST
	GivenName               ST
	MiddleName              ST
	Suffix                  ST
	Prefix                  ST
	Degree                  ST
	SourceTable             IS
	AssigningAuthority      HD
	NameTypeCode            ID // HL7 Table 0200
	IdentifierCheckDigit    ST
	CheckDigitSchemeCode    ID // HL7 Table 0061
	IdentifierTypeCode      IS
	AssigningFacility       HD
	DateTimeActionPerformed TS
}

/*
	TIME SERIES
*/

// Date/Time Range
type DR struct {
	StartDateTime TS
	EndDateTime   TS
}

// Repeat Interval
// TODO: parse TimeInterval as:
// HHMM,HHMM,HHMM,....
// Section 4.4.2.2
type RI struct {
	RepeatPattern IS
	TimeInterval  ST
}

// Scheduling Class Value Pair
type SCV struct {
	ParameterClass IS
	ParameterValue IS
}

type interval string
type duration string
type sequencing string

// Timing/quantity
// TODO: Section 4.4
type TQ struct {
	Quantity           CQ
	Interval           interval
	Duration           duration
	StartDateTime      TS
	EndDateTime        TS
	Priority           ID
	Condition          ST
	Text               TX
	Conjunction        ID
	OrderingSequencing sequencing
}

// Message Type
type CM_MSG struct {
	MessageType  ID
	TriggerEvent ID
}
