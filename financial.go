/*
This module contains the standard for segments found in financial management
messages.
*/
package main

// The standard GT1 segment
type GT1 struct {
	SetId                    SI  `hl7:"opt=R"`
	GuarantorNumber          CX  `hl7:"rep=Y"`
	Name                     XPN `hl7:"opt=R,rep=Y"`
	SpouseName               XPN `hl7:"rep=Y"`
	Address                  XAD `hl7:"rep=Y"`
	HomePhoneNumber          XTN `hl7:"rep=Y"`
	WorkPhoneNumber          XTN `hl7:"rep=Y"`
	DOB                      TS
	Sex                      IS
	Type                     IS
	RelationshipToPatient    IS
	SSN                      IS
	BeginDate                DT
	EndDate                  DT
	Priority                 NM
	EmployerName             XPN `hl7:"rep=Y"`
	EmployerAddress          XAD `hl7:"rep=Y"`
	EmployerPhoneNumber      XTN `hl7:"rep=Y"`
	EmployeeIdNumber         CX  `hl7:"rep=Y"`
	EmploymentStatus         IS
	OrganizationName         XON `hl7:"rep=Y"`
	BillingHoldFlag          ID
	CreditRatingCode         CE
	DeathDateTime            TS
	DeathFlag                ID
	ChargeAdjustmentCode     CE
	HouseholdAnnualIncome    CP
	HouseholdSize            NM
	EmployerIdNumber         CX `hl7:"rep=Y"`
	MaritalStatus            IS
	HireEffectiveDate        DT
	EmploymentStopDate       DT
	LivingDependency         IS
	AmbulatoryStatus         IS
	Citizenship              IS
	PrimaryLanguage          CE
	LivingArrangement        IS
	PublicityIndicator       CE
	ProtectionIndicator      ID
	StudentIndicator         IS
	Religion                 IS
	MotherMaidenName         XPN
	Nationality              CE
	EthnicGroup              IS
	ContactName              XPN `hl7:"rep=Y"`
	ContactPhoneNumber       XTN `hl7:"rep=Y"`
	ContactReason            CE
	ContactRelationship      IS
	JobTitle                 ST
	JobCode                  JCC
	EmployerOrganizationName XON `hl7:"rep=Y"`
	Handicap                 IS
	JobStatus                IS
	FinancialClass           FC
	Race                     IS
}

// The standard IN1 segment
type IN1 struct {
	SetId                    SI  `hl7:"opt=R"`
	PlanId                   CE  `hl7:"opt=R"`
	CompanyId                CX  `hl7:"opt=R,rep=Y"`
	CompanyName              XON `hl7:"rep=Y"`
	CompanyAddress           XAD `hl7:"rep=Y"`
	CompanyContact           XPN `hl7:"rep=Y"`
	CompanyPhoneNumber       XTN `hl7:"rep=Y"`
	GroupNumber              ST
	GroupName                XON `hl7:"rep=Y"`
	GroupEmployerId          CX  `hl7:"rep=Y"`
	GroupEmployerName        XON `hl7:"rep=Y"`
	PlanEffectiveDate        DT
	PlanExpirationDate       DT
	AuthorizationInformation CM_AUI
	PlanType                 IS
	InsuredName              XPN `hl7:"rep=Y"`
	RelationshipToPatient    IS
	InsuredDOB               TS
	InsuredAddress           XAD `hl7:"rep=Y"`
	AOB                      IS
	COB                      IS
	COBPriority              ST
	AdmissionFlag            ID
	AdmissionDate            DT
	EligibilityFlag          ID
	EligibilityDate          DT
	ReleaseInformationCode   IS
	PAC                      ST
	VerificationDateTime     TS
	VerificationBy           XCN
	AgreementCode            IS
	BillingStatus            IS
	LifetimeReserveDays      NM
	DelayBeforeLRDay         NM
	CompanyPlanCode          IS
	PolicyNumber             ST
	PolicyDeductible         CP
	PolicyLimitAmount        CP `hl7:"opt=B"`
	PolicyLimitDays          NM
	RoomRateSemiPrivate      CP `hl7:"opt=B"`
	RoomRatePrivate          CP `hl7:"opt=B"`
	InsuredEmploymentStatus  CE
	InsuredSex               IS
	InsuredEmployerAddress   XAD `hl7:"rep=Y"`
	VerificationStatus       ST
	PriorInsturancePlanId    IS
	CoverageType             IS
	Handicap                 IS
	InsuredIdNumber          CX `hl7:"rep=Y"`
}

// The standard IN2 segment
type IN2 struct {
	InsuredEmployeeId                  CX `hl7:"rep=Y"`
	InsuredSSN                         ST
	InsuredEmployerName                XCN `hl7:"rep=Y"`
	EmployerInformationData            IS
	MailClaimParty                     IS `hl7:"rep=Y"`
	MedicareCardNumber                 ST
	MedicaidCaseName                   XPN `hl7:"rep=Y"`
	MedicaidCaseNumber                 ST
	ChampuSponsorName                  XPN `hl7:"rep=Y"`
	ChampusIdNumber                    ST
	ChampusDependentRecipient          CE
	ChampusOrganization                ST
	ChampusStation                     ST
	ChampusService                     IS
	ChampusRank                        IS
	ChampusStatus                      IS
	ChampusRetireDate                  DT
	ChampusNonAvailCertOnFile          ID
	BabyCoverage                       ID
	CombineBabyBill                    ID
	BloodDeductible                    ST
	SpecialCoverageApprovalName        XPN `hl7:"rep=Y"`
	SpecialCoverageApprovalTitle       ST
	NoncoveredInsuranceCode            IS `hl7:"rep=Y"`
	PayorId                            CX `hl7:"rep=Y"`
	PayorSubscriberId                  CX `hl7:"rep=Y"`
	EligibilitySource                  IS
	RoomCoverageType                   CM_PLT `hl7:"rep=Y"`
	PolicyType                         CM_PLT `hl7:"rep=Y"`
	DailyDeductible                    CM_DDE
	LivingDependency                   IS
	AmbulatoryStatus                   IS
	Citizenship                        IS
	PrimaryLanguage                    CE
	LivingArrangement                  IS
	PublicityIndicator                 CE
	ProtectionIndicator                ID
	StudentIndicator                   IS
	Religion                           IS
	MotherMaidenName                   XPN
	Nationality                        CE
	EthnicGroup                        IS
	MaritalStatus                      IS `hl7:"rep=Y"`
	InsuredEmploymentStartDate         DT
	InsuredEmploymentStopDate          DT
	JobTitle                           ST
	JobCode                            JCC
	JobStatus                          IS
	EmployerContactName                XPN `hl7:"rep=Y"`
	EmployerContactPhoneNumber         XTN `hl7:"rep=Y"`
	EmployerContactReason              IS
	InsuredContactName                 XPN `hl7:"rep=Y"`
	InsuredContactPhoneNumbet          XTN `hl7:"rep=Y"`
	InsuredContactReason               IS  `hl7:"rep=Y"`
	RelationshipToPatientStartDate     DT
	RelationshipToPatientStopDate      DT `hl7:"rep=Y"`
	InsuranceCompanyContactReason      IS
	InsuranceCompanyContactPhoneNumber XTN
	PolicyScope                        IS
	PolicySource                       IS
	PatientMemberNumber                CX
	GuarantorRelationship              IS
	InsuredHomePhoneNumber             XTN `hl7:"rep=Y"`
	InsuredHomeWorkNumber              XTN `hl7:"rep=Y"`
	MilitaryHandicappedProgram         CE
	SuspendFlag                        ID
	CopayLimitFlag                     ID
	StoplossLimitFlag                  ID
	InsuredOrganizationName            XON `hl7:"rep=Y"`
	InsuredEmployerOrganizationName    XON `hl7:"rep=Y"`
	Race                               IS
	HcfaPatientRelationshipToInsured   CE
}

// The standard IN3 segment
type IN3 struct {
	SetId                              SI `hl7:"opt=R"`
	CertificationNumber                CX
	CertifiedBy                        XCN `hl7:"rep=Y"`
	CertificationRequired              ID
	Penalty                            CM_VAL
	CertificationDateTime              TS
	CertificationModalityDateTime      TS
	Operator                           XCN `hl7:"rep=Y"`
	CertificationBeginDate             DT
	CertificationEndDate               DT
	Days                               CM_VAL
	NonConcurCodeDescription           CE
	NonConcurEffectiveDateTime         TS
	PhysicianReviewer                  XCN `hl7:"rep=Y"`
	CertificationContact               ST
	CertificationContactPhoneNumber    XTN `hl7:"rep=Y"`
	AppealReason                       CE
	CertificationAgency                CE
	CertificationAgencyPhoneNumber     XTN    `hl7:"rep=Y"`
	PreCertRequirementWindow           CM_PCR `hl7:"rep=Y"`
	CaseManager                        ST
	SecondOpinionDate                  DT
	SecondOpinionStatus                IS
	SecondOpinionDocumentationReceived IS  `hl7:"rep=Y"`
	SecondOpinionPhysician             XCN `hl7:"rep=Y"`
}

// The standard ACC segment
type ACC struct {
	DateTime            TS
	Code                CE
	Location            ST
	AutoAccidentState   CE
	JobRelatedIndicator ID
	DeathIndicator      ID
}

// The standard UB1 segment
type UB1 struct {
	SetId                SI
	BloodDeductible      NM `hl7:"opt=B"`
	BloodFurnishedOf     NM
	BluodReplaced        NM
	BloodNotReplaced     NM
	CoInsuranceDays      NM
	ConditionCode        IS `hl7:"rep=Y5"`
	CoveredDays          NM
	NonCoveredDays       NM
	ValueAmount          CM_VAL `hl7:"rep=Y8"`
	GraceDays            NM
	SpecProgramIndicator CE
	ApprovalIndicator    CE
	ApprovedStayFrom     DT
	ApprovedStayTo       DT
	Occurrence           CE `hl7:"rep=Y5"` // NOTE: defind as CM in spec, but is actually the same structure as a CE
	OccurrenceSpan       CE
	OccurSpanStartDate   DT
	OccurSpanEndDate     DT
	Locator2             ST
	Locator9             ST
	Locator27            ST
	Locator45            ST
}

// The standard UB2 segment
type UB2 struct {
	SetId                 SI
	CoInsuranceDays       ST
	ConditionCode         IS `hl7:"rep=Y7"`
	CoveredDays           ST
	NonCoveredDays        ST
	ValueAmountCode       CM_VAL `hl7:"rep=Y12"`
	Occurrence            CM_OCD `hl7:"rep=Y8"`
	OccurrenceSpanCode    ST     `hl7:"rep=Y2"`
	Locator2              ST     `hl7:"rep=Y2"`
	Locator11             ST     `hl7:"rep=Y2"`
	Locator31             ST
	DocumentControlNumber ST `hl7:"rep=Y3"`
	Locator49             ST `hl7:"rep=Y23"`
	Locator56             ST `hl7:"rep=Y5"`
	Locator57             ST
	Locator78             ST `hl7:"rep=Y2"`
	SpecialVisitCount     NM
}

// The standard DG1 segment
type DG1 struct {
	SetId                   SI `hl7:"opt=R"`
	CodingMethod            ID `hl7:"opt=R"`
	Code                    CE
	Description             ST `hl7:"opt=B"`
	DateTime                TS
	Type                    IS  `hl7:"opt=R"`
	MajorDiagnosticCategory CE  `hl7:"opt=B"`
	DiagnosticRelatedGroup  CE  `hl7:"opt=B"`
	DRGApprovalIndicator    ID  `hl7:"opt=B"`
	DRGGrouperReviewCode    IS  `hl7:"opt=B"`
	OutlierType             CE  `hl7:"opt=B"`
	OutlierDays             NM  `hl7:"opt=B"`
	OutlierCost             CP  `hl7:"opt=B"`
	GoruperVersion          ST  `hl7:"opt=B"`
	Priority                NM  `hl7:"opt=B"`
	DiagnosingClinician     XCN `hl7:"rep=Y"`
	Classification          IS
	ConfidentialIndicator   ID
	AttestationDateTime     TS
}

// The standard DRG segment
type DRG struct {
	DiagnosticRelatedGroup CE
	AssignedDateTime       TS
	ApprovalIndicator      ID
	GrouperReviewCode      IS
	OutlierType            CE
	OutlierDays            NM
	OutlierCost            CP
	Payor                  IS
	OutlierReimbursement   CP
	ConfidentialIndicator  ID
}
