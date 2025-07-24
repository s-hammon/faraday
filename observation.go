/*
This module contains the standard for order entry messages.
Primarily, ORUs.
*/
package main

/*
This is a really stupid part of the HL7 "standard", as the OBX-5 field
(Observation Value) does not have a distinct type defined for it. Rather, they
let you define the type that makes the most sense for the observation type
and/or specialty. See section 7.3.2.5, Chapter 7 of the HL7 v2.3 standard.

FT (as well as ST) is the most permissible data type, so we will just use it
here as the "standard" data type.
*/
type OBX struct {
	SetId                        SI
	ValueType                    ID `hl7:"opt=C"`
	ObservationIdentifier        CE `hl7:"opt=R"`
	ObservationSubId             ST `hl7:"opt=C"`
	ObservationValue             FT `hl7:"opt=C,rep=Y"`
	Units                        CE
	ReferencesRange              ST
	AbnormalFlags                ID `hl7:"rep=Y5"`
	Probability                  NM
	AbnormalTestNature           ID `hl7:"rep=Y"`
	ResultStatus                 ID `hl7:"opt=R"`
	LastDateObservedNormalValues TS
	UserDefinedAccessChecks      ST
	ObservationDateTime          TS
	ProducerId                   CE
	ResponsibleObserver          XCN
	ObservationMethod            CE `hl7:"rep = Y"`
}
