package main

import (
	"strconv"

	"github.com/s-hammon/p"
)

// these are the basic HL7 "types"
// pretty much everything is a string

// string
// usually indicates "freeform" fields like patient names and procedures
type ST string

// numeric
// ....but actually, a string lol
type NM string

// TODO: implement ToNumeric

// sequence ID
// supposed to be a non-negative integer type
type SI string

func (s SI) ToNumeric() int {
	i, err := strconv.Atoi(string(s))
	if err != nil {
		panic(p.Format("non-numeric value for SI: %v", s))
	}
	return i
}

// Coded value from user-defined table
// aka custom values
type IS string

// Coded value from "official" table
// these usually follow a convention but aren't actually enforced lol
type ID string

// Timestamp, in the string format
type TS string

// TODO: implement ToDatetime

// Datetime, like TS
type DT string

// Telephone number
// Yep, that gets its own type but not SSN
type TN string

// TODO: implement ToPhone(?)

// Hierarchic Designator
type HD struct {
	NamespaceId     IS
	UniversalId     ST
	UniversalIdType ID
}

func (hd *HD) IsValidType() bool {
	return UniversalIdTypes.Valid(hd.UniversalIdType)
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

// Processing Type
// perfect example of the nonstandard standard of HL7--the fields are ST type
// but reference a table :)
type PT struct {
	ProcessingId   ST
	ProcessingMode ST
}

// Extended Composite ID w/ Check Digit
type CX struct {
	Id                   ST
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

// Extended Person Name
type XPN struct {
	FamilyName             ST
	GivenName              ST
	MiddleName             ST
	Suffix                 ST
	Prefix                 ST
	Degree                 ST
	NameTypeCode           ID
	NameRepresentationCode ID
}

// Extended Address
// last 2 fields are IS, but have no table--LOL
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

// Extended Telecommunication Number
type XTN struct {
	TelephoneNumber                TN
	TelecommunicationUseCode       ID
	TelecommunicationEquipmentType ID
	EmailAddress                   ST
	CountryCode                    NM
	AreaCode                       NM
	PhoneNumber                    NM
	Extension                      NM
	AAnyText                       ST
}

// Coded Element -- everyone's favorite type because it's very, very liberal
type CE struct {
	Identifier            ST
	Text                  ST
	CodingSystem          ST
	AlternateIdentifier   ST
	AlternateText         ST
	AlternateCodingSystem ST
}

// Driver's License Number -- but again, SSN is just a plain ol' string
type DLN struct {
	LicenseNumber  ST
	IssuingState   IS
	ExpirationDate DT
}
