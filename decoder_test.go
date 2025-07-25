package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecoder_MultipleOBX(t *testing.T) {
	raw := []byte(
		"MSH|^~\\&|LabSys|MainLab|EHR|Hospital|20250724000008||ORU^R01|MSGID123|P|2.3\r" +
			"PID|1||123456||DOE^JOHN\r" +
			"ORC|RE|999|||\r" +
			"OBR|1|999|ABC|CBC^Complete Blood Count\r" +
			"OBX|1|ST|WBC^White Blood Cells||5.4|10^9/L\r" +
			"OBX|2|ST|HGB^Hemoglobin||13.7|g/dL\r")

	var msg struct {
		MSH    MSH
		PID    PID
		Orders []Order `hl7:"ORC"`
	}

	err := NewDecoder(bytes.NewReader(raw)).Decode(&msg)
	require.NoError(t, err)
	require.Len(t, msg.Orders, 1)

	order := msg.Orders[0]
	require.Equal(t, ST("999"), order.ORC.PlacerOrderNumber.EntityIdentifier)
	require.Equal(t, CE{Identifier: "CBC", Text: "Complete Blood Count"}, order.OBR.UniversalServiceID)

	require.Len(t, order.OBX, 2)
	require.Equal(t, CE{Identifier: "WBC", Text: "White Blood Cells"}, order.OBX[0].ObservationIdentifier)
	require.Equal(t, FT("5.4"), order.OBX[0].ObservationValue)

	require.Equal(t, CE{Identifier: "HGB", Text: "Hemoglobin"}, order.OBX[1].ObservationIdentifier)
	require.Equal(t, FT("13.7"), order.OBX[1].ObservationValue)
}

func TestDecoder_MultipleORC(t *testing.T) {
	raw := []byte("MSH|^~\\&|SendingApp|SendingFac|ReceivingApp||20250724000008||ORM^O01|MSG00004|T|2.3\r" +
		"PID|1||123456||DOE^JOHN\r" +
		"ORC|RE|123|||\r" +
		"OBR|1|123|456|CT1^CT Head\r" +
		"ORC|NW|124|||\r" +
		"OBR|2|124|457|CT2^CT Abdomen\r")

	var msg struct {
		MSH    MSH
		PID    PID
		Orders []Order `hl7:"ORC"`
	}

	err := NewDecoder(bytes.NewReader(raw)).Decode(&msg)
	require.NoError(t, err)
	require.Len(t, msg.Orders, 2)
	require.Equal(t, ID("RE"), msg.Orders[0].ORC.OrderControl)
	require.Equal(t, SI("2"), msg.Orders[1].OBR.SetId)
}

func TestDecoder_ORM(t *testing.T) {
	raw := []byte("MSH|^~\\&|SendingApp|SendingFac|ReceivingApp||20250724000008||ORM^O01|MSG00003|T|2.3\rPID|1|W01536038|W02813944^2^6^^^S||DOE^JANE||19950101|F||W^White/Caucasian|123 MAIN ST^^ANYWHERE^TX^12345^USA||(123)456-7890^PRN|(999)999-9999^WPN|||NON^None|W182253636||||||||||||N\rPV1|1|E|^AER^^^^^^^Acme ER||||735360^Graham^Joshua^M^^^M.D.^^NPI&1234567890||||||||U|||E|W182253636|||||||||||||||||||||||||20250723214200\rGT1|1||DOE^JANE||123 MAIN ST^^ANYWHERE^TX^12345^USA|||19950101|||SELF^Self\rAL1|1||06004977^Penicillin G^penicillin G^^^penicillin G\rORC|XO|003953316|30120800||SC||^^^20250723221330||20250724000009|^Decrad^Support^^^^System.||735360^Graham^Joshua^M^^^M.D.^^NPI&1234567890|AERC^Acme ER Center||||||IDX\rOBR|1|003953316|30120800|XABDP1^CT Abdomen/Pelvis w/IV Cont^XABDP1^CT Abdomen/Pelvis w/IV Cont|Y^N||||||||SEVERE UPPER ABDOMINAL PAIN    DX:  ABDOMINAL PAIN    Comments: BARIATRIC PROTOCOL PLEASE|||735360^Graham^Joshua^M^^^M.D.^^NPI&1234567890||003953316||GRAJOS|MPE4660|||CT|C||^^15^20250723221330^^S^Stat|||A||||||20250723221500|||45893^^321 Niam St^Suite 404^Anywhere^TX^12345^321-654-0987^NORPT^(10/22dm)Meditech Orders^000-000-0000||0|0||||ISO300&Isovue - 300^GASTRO&Gastrografin - Iodinated Oral Contrast~5A90529^00600014~~~~100^60~&Farkas&Julie&&&&Technologist^&Farkas&Julie&&&&Technologist~07/23/2025 23:45:00^07/23/2025 23:15:00~002&Intravenous^001&Oral~ACTIVE^ACTIVE~A^A~CONTRAST^CONTRAST~~CC&cc^CC&cc|~~~~27.83~1733.80~~CTHABITUS\rNTE|1||BARIATRIC PROTOCOL PLEASE")

	dec := NewDecoder(bytes.NewReader(raw))

	orm := struct {
		MSH MSH
		PV1 PV1
		GT1 GT1
		AL1 AL1
		OBR OBR
	}{}
	err := dec.Decode(&orm)
	require.NoError(t, err)
	require.Equal(t,
		PL{
			Room:                IS("AER"),
			LocationDescription: ST("Acme ER"),
		},
		orm.PV1.AssignedPatientLocation,
	)
	require.Equal(t,
		IS("SELF^Self"),
		orm.GT1.RelationshipToPatient,
	)
	require.Equal(t,
		CE{
			Identifier:            ST("06004977"),
			Text:                  ST("Penicillin G"),
			CodingSystem:          ST("penicillin G"),
			AlternateCodingSystem: ST("penicillin G"),
		},
		orm.AL1.AllergyCode,
	)
}

func TestDecoder_ADT(t *testing.T) {
	raw := []byte("MSH|^~\\&|SendingApp|SendingFac|ReceivingApp|ReceivingFac|20250724000001||ADT^A02|MSG00002|P|2.3\rEVN|A02|20250724000001\rPID|1|W01222379|W02257226^^^^^SendingFac||DOE^JANE||19910101|F|||123 MAIN ST^^ANYWHERE^TX^12345^USA||(999)123-4567^PRN|(123)456-7890^WPN|||CHR^Christian|W182254551|987-65-4321|||||||||||N\rPV1|1|O|AER^Acme ER||||GRAJOS^Graham^Joshua^^^^M.D.||||||||||||W182254551|||||||||||||||||||||||||20250723234200")

	dec := NewDecoder(bytes.NewReader(raw))

	adt := struct {
		MSH MSH
		EVN EVN
		PID PID
		PV1 PV1
	}{}
	err := dec.Decode(&adt)
	require.NoError(t, err)
	require.Equal(t,
		MSH{
			FieldSeparator:       ST("|"),
			EncodingCharacters:   ST("^~\\&"),
			SendingApplication:   HD{NamespaceId: "SendingApp"},
			SendingFacility:      HD{NamespaceId: "SendingFac"},
			ReceivingApplication: HD{NamespaceId: "ReceivingApp"},
			ReceivingFacility:    HD{NamespaceId: "ReceivingFac"},
			DateTime:             TS("20250724000001"),
			MessageType:          CM_MSG{Type: "ADT", Event: "A02"},
			MessageControlId:     ST("MSG00002"),
			ProcessingId:         PT{ProcessingId: "P"},
			VersionId:            ID("2.3"),
		},
		adt.MSH,
	)
	require.Equal(t,
		EVN{
			EventTypeCode:    ID("A02"),
			RecordedDateTime: TS("20250724000001"),
		},
		adt.EVN,
	)
	require.Equal(t,
		PID{
			SetId:             SI("1"),
			ExternalPatientId: CX{IdNumber: ST("W01222379")},
			InternalPatientId: CX{
				IdNumber:          ST("W02257226"),
				AssigningFacility: HD{NamespaceId: IS("SendingFac")},
			},
			PatientName: XPN{
				FamilyName: ST("DOE"),
				GivenName:  ST("JANE"),
			},
			DOB: TS("19910101"),
			Sex: IS("F"),
			PatientAddress: XAD{
				StreetAddress: ST("123 MAIN ST"),
				City:          ST("ANYWHERE"),
				State:         ST("TX"),
				Zip:           ST("12345"),
				Country:       ID("USA"),
			},
			HomePhoneNumber: XTN{
				Number:                   TN("(999)123-4567"),
				TelecommunicationUseCode: ID("PRN"),
			},
			WorkPhoneNumber: XTN{
				Number:                   TN("(123)456-7890"),
				TelecommunicationUseCode: ID("WPN"),
			},
			Religion:              IS("CHR^Christian"),
			PatientAccountNumber:  CX{IdNumber: ST("W182254551")},
			SSN:                   ST("987-65-4321"),
			PatientDeathIndicator: ID("N"),
		},
		adt.PID,
	)
	require.Equal(t,
		PV1{
			SetId:        SI("1"),
			PatientClass: IS("O"),
			AssignedPatientLocation: PL{
				PointOfCare: IS("AER"),
				Room:        IS("Acme ER"),
			},
			AttendingDoctor: XCN{
				IdNumber:   ST("GRAJOS"),
				FamilyName: ST("Graham"),
				GivenName:  ST("Joshua"),
				Degree:     ST("M.D."),
			},
			VisitNumber:   CX{IdNumber: ST("W182254551")},
			AdmitDateTime: TS("20250723234200"),
		},
		adt.PV1,
	)
}

func BenchmarkDecoder_LargeORU(b *testing.B) {
	raw := []byte(
		"MSH|^~\\&|LIS|LabDept|EHR|MainHospital|20250724121200||ORU^R01|MSG123456|P|2.3\r" +
			"PID|1||12345678^^^MRN||Smith^Jane^A||19751225|F|||123 Main St^^Metropolis^NY^10001\r" +
			"PV1|1|O|AMB^101^1^Clinic||||1234^Primary^Care|||||||||||1234567\r" +
			"ORC|RE|ORD123456|||\r" +
			"OBR|1|ORD123456|LAB123456|80048^Basic Metabolic Panel^L||20250724120000|||||||1234^Primary^Care|||||F\r" +
			"OBX|1|NM|2345-7^Glucose^LN||98|mg/dL|70-99|N|||F\r" +
			"OBX|2|NM|3094-0^Urea Nitrogen (BUN)^LN||15|mg/dL|7-20|N|||F\r" +
			"OBX|3|NM|2951-2^Sodium^LN||139|mmol/L|135-145|N|||F\r" +
			"OBX|4|NM|2823-3^Potassium^LN||4.1|mmol/L|3.5-5.1|N|||F\r" +
			"OBX|5|NM|2075-0^Chloride^LN||102|mmol/L|98-107|N|||F\r" +
			"OBX|6|NM|2028-9^Calcium^LN||9.3|mg/dL|8.5-10.5|N|||F\r" +
			"OBX|7|NM|2160-0^Creatinine^LN||0.9|mg/dL|0.6-1.3|N|||F\r" +
			"OBX|8|NM|1863-0^eGFR^LN||90|mL/min/1.73m2|>=60|N|||F\r" +
			"NTE|1||Patient hydrated; levels normal.\r" +
			"OBX|9|NM|2093-3^Cholesterol, Total^LN||182|mg/dL|<200|N|||F\r" +
			"OBX|10|NM|2571-8^HDL Cholesterol^LN||52|mg/dL|>=40|N|||F\r" +
			"OBX|11|NM|2089-1^LDL Cholesterol^LN||110|mg/dL|<130|N|||F\r" +
			"OBX|12|NM|3043-7^Triglycerides^LN||140|mg/dL|<150|N|||F\r" +
			"NTE|2||Mildly elevated triglycerides, recommend follow-up.\r")

	type Message struct {
		MSH    MSH
		PID    PID
		PV1    PV1
		Orders []Order `hl7:"ORC"`
	}

	b.ResetTimer()
	for b.Loop() {
		var msg Message
		dec := NewDecoder(bytes.NewReader(raw))
		if err := dec.Decode(&msg); err != nil {
			b.Fatalf("Decode failed: %v", err)
		}
	}
}
