/*
This module contains the standard message types in v2.3
*/
package faraday

type ORM_O01 struct {
	MSH     MSH `hl7:"opt=R"`
	NTE     NTE
	Patient PatientGroup
	Order   OrderGroup
}

type ADT_A01 struct {
	MSH       MSH `hl7:"opt=R"`
	EVN       EVN `hl7:"opt=R"`
	PID       PID `hl7:"opt=R"`
	PD1       PD1
	NK1       []NK1
	PV1       PV1 `hl7:"opt=R"`
	PV2       PV2
	DB1       []DB1
	OBX       []OBX
	AL1       []AL1
	DG1       []DG1
	DRG       DRG
	Procedure []ProcedureGroup
	GT1       []GT1
	Insurance []InsuranceGroup
	ACC       ACC
	UB1       UB1
	UB2       UB2
}

type ORU_R01 struct {
	MSH     MSH           `hl7:"opt=R"`
	Results []ResultGroup `hl7:"opt=R"`
	DSC     DSC
}
