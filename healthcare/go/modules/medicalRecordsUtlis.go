package modules

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)



//recordID string, doctor string, patient string, name string, description string, date string) error {
func (hc *HealthcareContract) createMedicalRecord(stub shim.ChaincodeStubInterface, args[] string) peer.Response {
	
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 6.")
	}

	args[0] = "M" + args[0]
	args[1] = "P" + args[1]
	args[2] = "D" + args[2]

	//medical data shouldn't exist
    exists, err := stub.GetState(args[0])
    if err != nil {
        return shim.Error(err.Error())
    }
    if exists != nil {
        return shim.Error(fmt.Sprintf("the medical record with ID %s already exists", args[0]))
    }

	//patient should exist
	exists, err = stub.GetState(args[1])
    if err != nil {
        return shim.Error(err.Error())
    }
    if exists == nil {
        return shim.Error(fmt.Sprintf("Patient is not exist for %s id so, unable to create medical record", args[1]))
    }

	//doctor should exist
	exists, err = stub.GetState(args[2])
    if err != nil {
        return shim.Error(err.Error())
    }
    if exists == nil {
        return shim.Error(fmt.Sprintf("Doctor is not exist for %s id so, unable to create medical record", args[2]))
    }


	medicalRecord := MedicalRecord{
		ID : args[0],
		PatientID: args[1],
		DoctorID : args[2],
		Date : args[3],
		Prescription : args[4],
	}

    medicalRecordJSON, err := json.Marshal(medicalRecord)
    if err != nil {
        return shim.Error(err.Error())
    }

	recordID := args[0]

    err = stub.PutState(recordID, medicalRecordJSON)
    if err != nil {
        return shim.Error(err.Error())
    }

	// Retrieve the patient from the ledger
	patientJSON, err := stub.GetState(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}
	if patientJSON == nil {
		return shim.Error("Patient not found.")
	}

	// Unmarshal the patient object from JSON
	var patientRecord Patient
	err = json.Unmarshal(patientJSON, &patientRecord)
	if err != nil {
		return shim.Error(err.Error())
	}
    
	patientRecord.RecordIDs = append(patientRecord.RecordIDs, recordID)
    
	// Marshal the patient object to JSON
	patientJSON, err = json.Marshal(patientRecord)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Store the updated patient in the ledger
	err = stub.PutState(patientRecord.ID, patientJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

    return shim.Success([]byte(fmt.Sprintf("Medical record Id %s has been created for the patient id %s.", args[0],args[1])))
}


//patient Id and medical record id
func (hc *HealthcareContract) deleteMedicalRecord(stub shim.ChaincodeStubInterface, args[] string) peer.Response {
	
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2.")
	}

	args[0] = "P" + args[0]
	args[1] = "M" + args[1]

	//patient should exist
	exists, err := stub.GetState(args[0])
    if err != nil {
        return shim.Error(err.Error())
    }
    if exists == nil {
        return shim.Error(fmt.Sprintf("Patient is not exist for %s id so, unable to delete medical record", args[1]))
    }

	//medical data should exist
    exists, err = stub.GetState(args[1])
    if err != nil {
        return shim.Error(err.Error())
    }
    if exists == nil {
        return shim.Error(fmt.Sprintf("the medical record is not exist", args[1]))
    }

	// Retrieve the patient from the ledger
	patientJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if patientJSON == nil {
		return shim.Error("Patient not found.")
	}

	// Unmarshal the patient object from JSON
	var patientRecord Patient
	err = json.Unmarshal(patientJSON, &patientRecord)
	if err != nil {
		return shim.Error(err.Error())
	}
    
	var found bool   // default is false
	//delete the medical data from patient
	for pos,recordID := range patientRecord.RecordIDs {
		if recordID == args[1] {
			patientRecord.RecordIDs[pos] = patientRecord.RecordIDs[len(patientRecord.RecordIDs)-1]
			found = true		
		}
	}
	
	if found {
		patientRecord.RecordIDs = patientRecord.RecordIDs[0:len(patientRecord.RecordIDs)-1]
	} else {
		return shim.Error(fmt.Sprintf("Medical Record id %s is not found in %s patient Id",args[1],args[0]))
	}
	
	//delete the medical record
	err = stub.DelState(args[1]) 
    if err != nil {
		return shim.Error(err.Error())
	}

	// Marshal the patient object to JSON
	patientJSON, err = json.Marshal(patientRecord)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Store the updated patient in the ledger
	err = stub.PutState(patientRecord.ID, patientJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

    return shim.Success([]byte(fmt.Sprintf("Medical record Id %s has been deleted for the patient id %s.", args[1],args[0])))
}



func (hc *HealthcareContract) getAllMedicalRecordByPatientId(stub shim.ChaincodeStubInterface, args[] string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	args[0] = "P" + args[0]
	// Retrieve the patient from the ledger
	patientJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if patientJSON == nil {
		return shim.Error("Patient not found.")
	}

	// Unmarshal the patient object from JSON
	var patientRecord Patient
	err = json.Unmarshal(patientJSON, &patientRecord)
	if err != nil {
		return shim.Error(err.Error())
	}
	
	//create the medical Record array
	var medicalRecords [] MedicalRecord
    for _, recordID := range patientRecord.RecordIDs {
		//query the medical record data from the id
        medicalRecordJSON, err := stub.GetState(recordID)
        if err != nil {
            return shim.Error(err.Error())
        }
		if medicalRecordJSON == nil {
			return shim.Error(fmt.Sprintf("Medical Record id %s is not found",recordID))
		}
		//change into the object form json
		var medicalRecord MedicalRecord
		err = json.Unmarshal(medicalRecordJSON,&medicalRecord)
		if err != nil {
			return shim.Error(err.Error())
		}
        medicalRecords = append(medicalRecords, medicalRecord)
    }

	//change object to json
	medicalRecordsJSON, err2 := json.Marshal(medicalRecords)

	if err2 != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(medicalRecordsJSON)
}


//patient Id , entity type , entity id
func (hc *HealthcareContract) getAllMedicalRecordByPatientIdWithAccess(stub shim.ChaincodeStubInterface, args[] string) peer.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3.")
	}

	args[0] = "P" + args[0]
	// Retrieve the patient from the ledger
	patientJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if patientJSON == nil {
		return shim.Error("Patient not found.")
	}
	// Unmarshal the patient object from JSON
	var patient Patient
	err = json.Unmarshal(patientJSON, &patient)
	if err != nil {
		return shim.Error(err.Error())
	}

	//entity is exist or not
	if(args[1]=="Doctor") {
		args[2] = "D" + args[2]
	}

	//check access
	var access bool

	for _,e := range(patient.Read) {
		if e==args[2] {
			access = true
			break
		}
	}
	
	if !access {
		return shim.Error(fmt.Sprintf("This %d is not acess to read the medical records",args[2]))
	}

	pass := []string{patient.ID[1:]}
	queryResult := hc.getAllMedicalRecordByPatientId(stub,pass)

	return shim.Success([]byte(string(queryResult.Payload)))
}

func (hc *HealthcareContract) getMedicalRecordById(stub shim.ChaincodeStubInterface, args[] string) peer.Response {
	args[0] = "M" + args[0]
	
	medicalRecordJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if medicalRecordJSON == nil {
		return shim.Error(fmt.Sprintf("Medical Record Id %s has not found",args[0]))
	}
	
	return shim.Success(medicalRecordJSON)
}


