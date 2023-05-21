package modules

import (
	"encoding/json"
	"strconv"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// Define the createPatient method for the healthcare smart contract
func (hc *HealthcareContract) createPatient(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7.")
	}
	
	exists, err := stub.GetState(args[0])
    
	if err != nil {
        return shim.Error(err.Error())
    }
    if exists != nil {
        return shim.Error(fmt.Sprintf("the patient is already exists with ID %s", args[0]))
    }

	// var age int
	age,err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error(err.Error())
	}
	
	// Create a new patient object
	patient := Patient{
		ID:          args[0],
		Name:        args[1],
		Age:         age,
		Gender:      args[3],
		BloodType:   args[4],
		Address:     args[5],
		PhoneNumber: args[6],
		RecordIDs:  []string{},
		Read:       []string{},
		Write:      []string{},
	}

	// Marshal the patient object to JSON
	patientJSON, err := json.Marshal(patient)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Store the patient in the ledger
	err = stub.PutState(patient.ID, patientJSON)
	if err != nil {
		return shim.Error(err.Error())
	}
	
	// Return a success response
	return shim.Success([]byte(fmt.Sprintf("Patient Id %s has been successfuly created.", args[0])))
}



// Define the getPatient method for the healthcare smart contract
func (hc *HealthcareContract) getPatient(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}
	// Retrieve the patient from the ledger
	patientJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if patientJSON == nil {
		return shim.Error("Patient not found.")
	}
	// Return the patient object
	return shim.Success(patientJSON)
}



// Define the updatePatient method for the healthcare smart contract
func (hc *HealthcareContract) updatePatient(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7.")
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

	// var age int

	age,err := strconv.Atoi(args[2])

	// Update the patient object
	patient.Name = args[1]
	patient.Age = age
	patient.Gender = args[3]
	patient.BloodType = args[4]
	patient.Address = args[5]
	patient.PhoneNumber = args[6]

	// Marshal the patient object to JSON
	patientJSON, err = json.Marshal(patient)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Store the updated patient in the ledger
	err = stub.PutState(patient.ID, patientJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return a success response
	return shim.Success([]byte(fmt.Sprintf("Patient Id %s has successfully updated", args[0])))

}


// Define the deletePatient method for the healthcare smart contract
func (hc *HealthcareContract) deletePatient(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}
	// Retrieve the patient from the ledger
	patientJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if patientJSON == nil {
		return shim.Error("Patient not found.")
	}

	//unmarshal the patient json
	var patient Patient
	err = json.Unmarshal(patientJSON, &patient)
	if err != nil {
		return shim.Error(err.Error())
	}

	//Delete all the medical records of the patient
	for _,record := range(patient.RecordIDs) {
		err = stub.DelState(record)
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	// Delete the patient from the ledger
	err = stub.DelState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return a success response
	return shim.Success([]byte(fmt.Sprintf("Patient Id %s has been deteted.", args[0])))
}
