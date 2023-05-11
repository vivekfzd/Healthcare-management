package modules

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func (hc *HealthcareContract) createLabTechnician(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2.")
	}

	id := args[0]
	name := args[1]

	// Check if the lab technician already exists
	existingTechnician, err := stub.GetState(id)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get lab technician: %s", err.Error()))
	}
	if existingTechnician != nil {
		return shim.Error(fmt.Sprintf("Lab technician with ID %s already exists", id))
	}

	// Create a new lab technician instance
	technician := LabTechnician{
		ID:   id,
		Name: name,
	}

	// Convert the lab technician to JSON
	technicianJSON, err := json.Marshal(technician)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to marshal lab technician: %s", err.Error()))
	}

	// Save the lab technician to the ledger
	err = stub.PutState(id, technicianJSON)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to save lab technician: %s", err.Error()))
	}

	return shim.Success([]byte("Lab technician created successfully"))
}

func (hc *HealthcareContract) getLabTechnician(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	id := args[0]

	// Retrieve the lab technician from the ledger
	technicianJSON, err := stub.GetState(id)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get lab technician: %s", err.Error()))
	}
	if technicianJSON == nil {
		return shim.Error(fmt.Sprintf("Lab technician with ID %s does not exist", id))
	}

	return shim.Success(technicianJSON)
}

func (hc *HealthcareContract) updateLabTechnician(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2.")
	}

	id := args[0]
	newName := args[1]

	// Retrieve the lab technician from the ledger
	technicianJSON, err := stub.GetState(id)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get lab technician: %s", err.Error()))
	}
	if technicianJSON == nil {
		return shim.Error(fmt.Sprintf("Lab technician with ID %s does not exist", id))
	}

	// Update the lab technician's name
	technician := LabTechnician{}
	err = json.Unmarshal(technicianJSON, &technician)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to unmarshal lab technician: %s", err.Error()))
	}
	technician.Name = newName

	// Convert the updated lab technician back to JSON
	updatedTechnicianJSON, err := json.Marshal(technician)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to marshal updated lab technician: %s", err.Error()))
	}

	// Save the updated lab technician to the ledger
	err = stub.PutState(id, updatedTechnicianJSON)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to save updated lab technician: %s", err.Error()))
	}

	return shim.Success([]byte("Lab technician updated successfully"))
}
