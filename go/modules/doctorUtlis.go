package modules

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// Define the createDoctor method for the healthcare smart contract
func (hc *HealthcareContract) createDoctor(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4.")
	}
	args[0] = "D" + args[0]
	exists, err := stub.GetState(args[0])
    if err != nil {
        return shim.Error(err.Error())
    }
    if exists != nil {
        return shim.Error(fmt.Sprintf("the docotor is already exists with ID %s", args[0]))
    }
	// Create a new doctor object
	doctor := Doctor{
		ID:          args[0],
		Name:        args[1],
		Specialty:   args[2],
		PhoneNumber: args[3],
	}

	// Marshal the doctor object to JSON
	doctorJSON, err := json.Marshal(doctor)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Store the doctor in the ledger
	err = stub.PutState(doctor.ID, doctorJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return a success response
	return shim.Success([]byte(fmt.Sprintf("Doctor has been created and id is %s", args[0])))
}

// Define the getDoctor method for the healthcare smart contract
func (hc *HealthcareContract) getDoctor(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}
	args[0] = "D" + args[0]
	// Retrieve the doctor from the ledger
	doctorJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if doctorJSON == nil {
		return shim.Error(fmt.Sprintf("Doctor has not found for the following id %s",args[0]))
	}

	// Unmarshal the doctor object from JSON
	var doctor Doctor
	err = json.Unmarshal(doctorJSON, &doctor)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Marshal the doctor object to JSON
	doctorJSON, err = json.Marshal(doctor)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return the doctor object
	return shim.Success(doctorJSON)
}

// Define the updateDoctor method for the healthcare smart contract
func (hc *HealthcareContract) updateDoctor(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4.")
	}
	args[0] = "D" + args[0]
	// Retrieve the doctor from the ledger
	doctorJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if doctorJSON == nil {
		return shim.Error("Doctor has not been not found.")
	}

	// Unmarshal the doctor object from JSON
	var doctor Doctor
	err = json.Unmarshal(doctorJSON, &doctor)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Update the doctor object
	doctor.Name = args[1]
	doctor.Specialty = args[2]
	doctor.PhoneNumber = args[3]

	// Marshal the doctor object to JSON
	doctorJSON, err = json.Marshal(doctor)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Store the updated doctor in the ledger
	err = stub.PutState(doctor.ID, doctorJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return a success response
	return shim.Success([]byte(fmt.Sprintf("Doctor has been updated and id is %s", args[0])))
}

// Define the deleteDoctor method for the healthcare smart contract
func (hc *HealthcareContract) deleteDoctor(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}
	args[0] = "D" + args[0]
	// Retrieve the doctor from the ledger
	doctorJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if doctorJSON == nil {
		return shim.Error("Doctor has not been found.")
	}
	// Delete the doctor from the ledger
	err = stub.DelState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	// Return a success response
	return shim.Success([]byte(fmt.Sprintf("Doctor Id %s has been deteted.", args[0])))
}