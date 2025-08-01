// NOTE: HL7 denotes "optionality" of segments as follows:
// PID (no brackets): segment required, appears exactly once
// [PD1] (square brackets): segment optional, appears exactly once
// {[AL1]} (square+curly brackets): segment optional, may appear 1+ times

// NOTE: It then follows that { ... } denotes a required segment which may
// appear 1+ times, although this is rarely observed (if at all). Rather,
// { ... } is more often used for the so-called segment groups. For example, we
// may define a simple OrderGroup as follows:

// NOTE:
// { OrderGroup
//
//	  ORC
//	  [ OrderDetailGroup
//	    OBR
//	    {[NTE]}
//	    {[DG1]}
//	    {[ ObservationGroup
//	      OBX
//	      {[NTE]}
//	    ]}
//	  ]
//	}
//
// NOTE: In faraday, this is represented as:
//
//	type OrderGroup struct {
//		ORC ORC
//		Details OrderDetailGroup
//	}
//	type OrderDetailGroup struct {
//		OBR OBR `hl7:"opt=R"`
//		NTE []NTE
//		DG1 []DG1
//		Results []ObservationGroup
//	}
//	type ObservationGroup struct {
//		OBX OBX `hl7:"opt=R"`
//		NTE []NTE
//	}
//
// NOTE: This would then be included in a struct representing an ORM as:
//
//	type ORM struct {
//		MSH MSH `hl7:"opt=R"`
//		PID PID `hl7:"opt=R"`
//		PV1 PV1 `hl7:"opt=R"`
//		Order []OrderGroup `hl7:"opt=R"`
//	}
//
// "opt=R" indicates that the segment/group is required. If the segment/group
// is a slice, then at least one is required.
// */
package faraday

// the standard PatientGroup
type PatientGroup struct {
	PID       PID `hl7:"opt=R"`
	PD1       PD1
	Visit     PatientVisitGroup
	Insurance []InsuranceGroup
	GT1       GT1
	AL1       []AL1
}

// The standard PatientVisitGroup
type PatientVisitGroup struct {
	PV1 PV1 `hl7:"opt=R"`
	PV2 PV2
}

// The standard InsuranceGroup
type InsuranceGroup struct {
	IN1 IN1 `hl7:"opt=R"`
	IN2 IN2
	IN3 IN3
}

// The standard OrderGroup (ORM)
type OrderGroup struct {
	ORC     ORC
	Details OrderDetailGroup
}

// The standard Order Detail Group
type OrderDetailGroup struct {
	OBR     OBR `hl7:"opt=R"`
	NTE     []NTE
	DG1     []DG1
	Results []ObservationGroup
}

// The standard ObservationGroup
type ObservationGroup struct {
	OBX OBX `hl7:"opt=R"`
	NTE []NTE
}

// The standard ProcedureGroup
type ProcedureGroup struct {
	PR1 PR1 `hl7:"opt=R"`
	// ROL []ROL
}

// The standard ResultGroup
type ResultGroup struct {
	Patient ObsPatientGroup
	Order   []ObsOrderGroup `hl7:"opt=R"`
}

// The standard PatientGroup (ORU)
type ObsPatientGroup struct {
	PID   PID `hl7:"opt=R"`
	PD1   PD1
	NTE   []NTE
	Visit PatientVisitGroup
}

// The standard OrderGroup (ORU)
type ObsOrderGroup struct {
	ORC     ORC
	OBR     OBR `hl7:"opt=R"`
	NTE     []NTE
	Results []ObservationGroup `hl7:"opt=R"`
	// CTI []CTI
}
