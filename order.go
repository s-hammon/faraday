/*
This module contains the standard for segments found in order entry messages.
*/
package faraday

// The standard ORC segment
type ORC struct {
	OrderControl           ID `hl7:"opt=R"`
	PlacerOrderNumber      EI `hl7:"opt=C"`
	FillerOrderNumber      EI `hl7:"opt=C"`
	PlacerGroupNumber      EI
	OrderStatus            ID
	ResponseFlag           ID
	QuantityTiming         TQ
	Parent                 CM_POR
	TransactionDateTime    TS
	EnteredBy              XCN
	VerifiedBy             XCN
	OrderingProvider       XCN
	EntryLocation          PL
	CallbackPhoneNumber    XTN `hl7:"rep=Y2"`
	EffectiveDateTime      TS
	OrderControlCodeReason CE
	EnteringOrganization   CE
	EnteringDevice         CE
	ActionBy               XCN
}

// The standard OBR segment
type OBR struct {
	SetId                              SI `hl7:"opt=C"`
	PlacerOrderNumber                  EI `hl7:"opt=C"`
	FillerOrderNumber                  EI `hl7:"opt=C"`
	UniversalServiceID                 CE `hl7:"opt=R"`
	Priority                           ID `hl7:"opt=B"`
	RequestedDateTime                  TS `hl7:"opt=B"`
	ObservationDateTime                TS `hl7:"opt=C"`
	ObservationEndDateTime             TS
	CollectionVolume                   CQ
	CollectorIdentifier                XCN `hl7:"rep=Y"`
	SpecimenActionCode                 ID
	DangerCode                         CE
	RelevantClinicalInfo               ST
	SpecimenReceivedDateTime           TS `hl7:"opt=C"`
	SpecimenSource                     CM_SPE
	OrderingProvider                   XCN `hl7:"rep=Y"`
	OrderCallbackPhoneNumber           XTN `hl7:"rep=Y2"`
	PlacerField1                       ST
	PlacerField2                       ST
	FillerField1                       ST
	FillerField2                       ST
	StatusChangeDatTime                TS `hl7:"opt=C"`
	ChargeToPractice                   CM_CHP
	DiagnosticServiceSectionId         ID
	ResultStatus                       ID `hl7:"opt=C"`
	ParentResult                       CM_PRE
	QuantityTiming                     TQ  `hl7:"rep=Y"`
	ResultCopiesTo                     XCN `hl7:"rep=Y5"`
	Parent                             CM_POR
	TransportationMode                 ID
	ReasonForStudy                     CE `hl7:"rep=Y"`
	PrincipalResultInterpreter         CM_OBS
	AssistantResultInterpreter         CM_OBS `hl7:"rep=Y"`
	Technician                         CM_OBS `hl7:"rep=Y"`
	Transcriptionist                   CM_OBS `hl7:"rep=Y"`
	ScheduledDateTime                  TS
	SampleContainersCount              NM
	SampleTransportLogistics           CE `hl7:"rep=Y"`
	CollectorComment                   CE `hl7:"rep=Y"`
	TransportArrangementResponsibility CE
	TransportArranged                  ID
	EscortRequired                     ID
	PlannedPatientTransportComment     CE `hl7:"rep=Y"`
}

// An Order Group--contains an ORC, optionally followed by an OBR and then
// pootentially many OBX, NTE, etc
type Order struct {
	ORC ORC
	OBR OBR
	OBX []OBX
}
