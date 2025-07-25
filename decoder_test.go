package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecoder_MultipleORC(t *testing.T) {
	raw := []byte("MSH|^~\\&|SendingApp|SendingFac|ReceivingApp||20250724000008||ORM^O01|MSG00004|T|2.3\r" +
		"PID|1||123456||DOE^JOHN\r" +
		"ORC|RE|123|||\r" +
		"OBR|1|123|456|CT1^CT Head\r" +
		"ORC|NW|124|||\r" +
		"OBR|2|124|457|CT2^CT Abdomen\r")

	var msg struct {
		MSH MSH
		PID PID
		ORC []ORC
		OBR []OBR
	}

	err := NewDecoder(bytes.NewReader(raw)).Decode(&msg)
	require.NoError(t, err)
	require.Len(t, msg.ORC, 2)
	require.Len(t, msg.OBR, 2)
	require.Equal(t, ST("123"), msg.ORC[0].PlacerOrderNumber.EntityIdentifier)
	require.Equal(t, ST("124"), msg.ORC[1].PlacerOrderNumber.EntityIdentifier)
	require.Equal(t, ST("CT1"), msg.OBR[0].UniversalServiceID.Identifier)
	require.Equal(t, ST("CT2"), msg.OBR[1].UniversalServiceID.Identifier)
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
