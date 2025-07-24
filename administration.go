/*
This module contains the standard for patient administration messages.
Primarly, ADTs.
*/
package main

// The standard EVN segment
type EVN struct {
	EventTypeCode        ID `hl7:"opt=B,tbl=0003"`
	RecordedDateTime     TS `hl7:"opt=R"`
	PlannedEventDateTime TS
	EventReasonCode      IS  `hl7:"tbl=0062"`
	OperatorID           XCN `hl7:"tbl=0188"`
	EventOccurred        TS
}

// The standard PID segment
type PID struct {
	SetId                  SI
	ExternalPatientId      CX
	InternalPatientId      CX  `hl7:"opt=R,rep=Y"`
	AlternatePatientId     CX  `hl7:"rep=Y"`
	PatientName            XPN `hl7:"opt=R,rep=Y"`
	MotherMaidenName       XPN
	DOB                    TS
	Sex                    IS
	PatientAlias           XPN `hl7:"rep=Y"`
	Race                   IS
	PatientAddress         XAD `hl7:"rep=Y"`
	CountyCode             IS  `hl7:"opt=B"`
	HomePhoneNumber        XTN `hl7:"rep=Y"`
	WorkPhoneNumber        XTN `hl7:"rep=Y"`
	PrimaryLanguage        CE
	MaritalStatus          IS
	Religion               IS
	PatientAccountNumber   CX
	SSN                    ST
	DriversLicenseNumber   DLN
	MotherIdentifier       CX `hl7:"rep=Y"`
	EthnicGroup            IS
	BirthPlace             ST
	MultipleBirthIndicator ID
	BirthOrder             NM
	Citizenship            IS `hl7:"rep=Y"`
	VeteranStatus          CE
	Nationality            CE
	PatientDeathDateTime   TS
	PatientDeathIndicator  ID
}

// The standard PV1 segment
type PV1 struct {
	SetId                   SI
	PatientClass            IS `hl7:"opt=R"`
	AssignedPatientLocation PL
	AdmissionType           IS
	PreadmitNumber          CX
	PriorPatientLocation    PL
	AttendingDoctor         XCN `hl7:"rep=Y"`
	ReferringDoctor         XCN `hl7:"rep=Y"`
	ConsultingDoctor        XCN `hl7:"rep=Y"`
	HospitalService         IS
	TemporaryLocation       PL
	PreadmitTestIndicator   IS
	ReadmissionIndicator    IS
	AdmitSource             IS
	AmbulatoryStatus        IS `hl7:"rep=Y"`
	VipIndicator            IS
	AdmittingDoctor         XCN `hl7:"rep=Y"`
	PatientType             IS
	VisitNumber             CX
	FinancialClass          FC `hl7:"rep=Y"`
	ChargePriceIndicator    IS
	CourtesyCode            IS
	CreditRating            IS
	ContractCode            IS `hl7:"rep=Y"`
	ContractEffectiveDate   DT `hl7:"rep=Y"`
	ContractAmount          NM `hl7:"rep=Y"`
	ContractPeriod          NM `hl7:"rep=Y"`
	InterestCode            IS
	TransferBadDebtCode     IS
	TransferBadDebtDate     DT
	BadDebtAgencyCode       IS
	BadDebtTransferAmount   NM
	BadDebtRecoveryAmount   NM
	DeleteAccountIndicator  IS
	DeleteAccountDate       DT
	DischargeDisposition    IS
	DischargedToLocation    CM_DSL
	DietType                IS
	ServicingFacility       IS
	BedStatus               IS `hl7:"opt=B"`
	AccountStatus           IS
	PendingLocation         PL
	PriorTemporaryLocation  PL
	AdmitDateTime           TS
	DischargeDateTime       TS
	CurrentPatientBalance   NM
	TotalCharges            NM
	TotalAdjustments        NM
	TotalPayments           NM
	AlternateVisitId        CX
	VisitIndicator          IS
	OtherHealthcareProvider XCN `hl7:"rep=Y"`
}

// The standard PV2 segment
type PV2 struct {
	PriorPendingLocation              PL `hl7:"opt=C"`
	AccomodationCode                  CE
	AdmitReason                       CE
	TransferReason                    CE
	PatientValuables                  ST `hl7:"rep=Y"`
	PatientValuablesLocation          ST
	VisitUserCode                     IS
	ExpectedAdmitDateTime             TS
	ExpectedDischargeDateTime         TS
	EstimatedLengthInpatientStay      NM
	ActualLengthInpatientStay         NM
	VisitDescription                  ST
	ReferralSourceCode                XCN
	PreviousServiceDAte               DT
	EmploymentIllnessRelatedIndicator ID
	PurgeStatusCode                   IS
	PurgeStatusDate                   DT
	SpecialProgramCode                IS
	RetentionIndicator                ID
	ExpectedCountInsurancePlans       NM
	VisitPublicityCode                IS
	VisitProtectionIndicator          ID
	ClinicOrganizationName            XON `hl7:"rep=Y"`
	PatientStatusCode                 IS
	VisitPriorityCode                 IS
	PreviousTreatmentDAte             DT
	ExpectedDischargeDisposition      IS
	FileSignatureDate                 DT
	FirstSimilarIllnessDate           DT
	PatientChargeAdjustmentCode       IS
	RecurringServiceCode              IS
	BillingMediaCode                  ID
	ExpectedSurgeryDateTime           TS
	MilitaryPartnershipCode           ID
	MilitaryNonAvailabilityCode       ID
	NewbornBabyIndicator              ID
	BabyDetainedIndicator             ID
}

// The standard NK1 segment
type NK1 struct {
	SetId                   SI  `hl7:"opt=R"`
	Name                    XPN `hl7:"rep=Y"`
	Relationship            CE
	Address                 XAD `hl7:"rep=Y"`
	PhoneNumber             XTN `hl7:"rep=Y"`
	WorkPhoneNumber         XTN `hl7:"rep=Y"`
	ContactRole             CE
	StartDate               DT
	EndDate                 DT
	NextOfKinJobTitle       ST
	NextOfKinJobCode        JCC
	NextOfKinEmployeeNumber CX
	OrganizationName        XON `hl7:"rep=Y"`
	MaritalStatus           IS
	Sex                     IS
	DOB                     TS
	LivingDependency        IS `hl7:"rep=Y"`
	AmbulatoryStatus        IS `hl7:"rep=Y"`
	Citizenship             IS `hl7:"rep=Y"`
	PrimaryLanguage         CE
	LivingArrangement       IS
	PublicityIndicator      CE
	ProtectionIndicator     ID
	StudentIndicator        IS
	Religion                IS
	MotherMaidenName        XPN
	Nationality             CE
	EthnicGroup             IS
	ContactReason           CE  `hl7:"rep=Y"`
	ContactName             XPN `hl7:"rep=Y"`
	ContactTelephoneNumber  XTN `hl7:"rep=Y"`
	ContactAddress          XAD `hl7:"rep=Y"`
	NextOfKinIdentifiers    CX  `hl7:"rep=Y"`
	JobStatus               IS
	Race                    IS
	Handicap                IS
	ContactSSN              ST
}

// The standard AL1 segment
type AL1 struct {
	SetId              SI `hl7:"opt=R"`
	AllergyType        IS
	AllergyCode        CE `hl7:"opt=R"`
	AllergySeverity    IS
	AllergyReaction    ST
	IdentificationDate DT
}

// The standard NPU segment
type NPU struct {
	BedLocation PL `hl7:"opt=R"`
	BedStatus   IS
}

// The standard MRG segment
type MRG struct {
	PriorInternalPatientId    CX `hl7:"opt=R,rep=Y"`
	PriorAlternatePatientId   CX `hl7:"rep=Y"`
	PriorPatientAccountNumber CX
	PriorExternalPatientId    CX
	PriorVisitNumber          CX
	PriorAlternateVisitId     CX
	PriorPatientName          XPN
}

// The standard PD1 segment
type PD1 struct {
	LivingDependency       IS `hl7:"rep=Y"`
	LivingArrangement      IS
	PatientPrimaryFacility XON `hl7:"rep=Y"`
	PatientPCPName         XCN `hl7:"rep=Y"`
	StudentIndicator       IS
	Handicap               IS
	LivingWill             IS
	OrganDonor             IS
	SeparateBill           ID
	DuplicatePatient       CX `hl7:"rep=Y"`
	PublicityIndicator     CE
	ProtectionIndicator    ID
}

// The standard DB1 segment
type DB1 struct {
	SetId            SI `hl7:"opt=R"`
	PersonCode       IS
	PersonIdentifier CX `hl7:"rep=Y"`
	Indicator        ID
	StartDate        DT
	EndDate          DT
	ReturnToWorkDate DT
	UnableToWorkDate DT
}
