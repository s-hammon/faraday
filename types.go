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
type ST string

// Text
type TX string

// Formatted Text
type FT string

/*
	NUMERICAL
*/

// Numeric
type NM string

// TODO: implement ToNumeric

// Sequence ID (integer)
type SI string

func (s SI) ToNumeric() int {
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
type MO struct {
	Quantity     NM
	Denomination ID
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

// Structured Numeric
type SN struct {
	Comparator comparator
	FirstNum   NM
	SecondNum  NM
	Separator  ST
}

func (sn *SN) IsValidType() bool {
	switch comparator(sn.Comparator) {
	default:
		return false
	case GT, LT, GE, LE, EQ, NE:
		return true
	}
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
	UniversalIdType ID
}

func (hd *HD) IsValidType() bool {
	return UniversalIdTypes.Valid(hd.UniversalIdType)
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
	DataType      ID
	DataSubType   ID
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
	ProcessingId   ID
	ProcessingMode ID
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
	AlternateIdentifier   ID
	AlternateText         FT
	AlternateCodingSystem ST
}

// Coded Element with formatted values
type CF struct {
	Identifier            ID
	FormattedText         FT
	CodingSystem          ST
	AlternateIdentifier   ID
	AlternateText         FT
	AlternateCodingSystem ST
}

// Composite ID with check digit
type CK struct {
	IdNumber             NM
	CheckDigit           NM
	CheckDigitSchemeCode ID
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
	CheckDigitSchemeCode ID
	AssigningAuthority   HD
	IdentifierTypeCode   IS
	AssigningFacility    HD
}

func (cx *CX) IsValidType() bool {
	return cx.AssigningAuthority.IsValidType() &&
		cx.AssigningFacility.IsValidType() &&
		IdentifierTypeCodes.Valid(cx.CheckDigitSchemeCode)
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
	NameTypeCode         ID
	IdentifierCheckDigit ST
	CheckDigitSchemeCode ID
	IdentifierTypeCode   IS
	AssigningFacility    HD
}

// Composite -- apparently they aren't allowed
type CM struct{}

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
	Country                    ID
	AddressType                ID
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
	Country                    ID
	AddressType                ID
	OtherGeographicDesignation ST
	CountyCode                 IS
	CensusTract                IS
}

// Extended Person Name
type XPN struct {
	FamilyName   ST
	GivenName    ST
	MiddleName   ST
	Suffix       ST
	Prefix       ST
	Degree       ST
	NameTypeCode ID
}

// Extended Composite Name and ID number for organizations
type XON struct {
	OrganizationName     ST
	TypeCode             IS
	IdNumber             NM
	CheckDigit           NM
	CheckDigitSchemeCode ID
	AssigningAuthority   HD
	IdentifierTypeCode   IS
	AssigningFacility    HD
}

// Extended Telecommunication Number
type XTN struct {
	TelecommunicationUseCode ID
	EquipmentType            ID
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

// Query Selection Criteria
type QSC struct {
	FieldName             ST
	RelationalOperator    ID
	Value                 ST
	RelationalConjunction ID
}

// Query Input Parameter
type QIP struct {
	FieldName ST
	Values    []ST // TODO: check this
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
	StartDayRange  ID
	EndDayRange    ID
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
	NameTypeCode            ID
	IdentifierCheckDigit    ST
	CheckDigitSchemeCode    ID
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

func (cm *CM_MSG) IsValidType() bool {
	return UniversalIdTypes.Valid(cm.MessageType) &&
		UniversalIdTypes.Valid(cm.TriggerEvent)
}
