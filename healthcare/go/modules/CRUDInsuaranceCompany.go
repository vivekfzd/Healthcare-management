package modules

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func (hc *HealthcareContract) createInsuranceCompany(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3.")
	}

	id := args[0]
	name := args[1]
	company := args[2]

	// Check if the insurance company already exists
	existingCompanyJSON, err := stub.GetState(id)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get insurance company: %s", err.Error()))
	}
	if existingCompanyJSON != nil {
		return shim.Error(fmt.Sprintf("Insurance company with ID %s already exists", id))
	}

	// Create a new insurance company instance
	insuranceCompany := InsuranceCompany{
		ID:      id,
		Name:    name,
		Company: company,
	}

	// Convert the insurance company to JSON
	companyJSON, err := json.Marshal(insuranceCompany)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to marshal insurance company: %s", err.Error()))
	}

	// Save the insurance company to the ledger
	err = stub.PutState(id, companyJSON)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to save insurance company: %s", err.Error()))
	}

	return shim.Success([]byte("Insurance company created successfully"))
}

func (hc *HealthcareContract) getInsuranceCompany(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	id := args[0]

	// Retrieve the insurance company from the ledger
	companyJSON, err := stub.GetState(id)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get insurance company: %s", err.Error()))
	}
	if companyJSON == nil {
		return shim.Error(fmt.Sprintf("Insurance company with ID %s does not exist", id))
	}

	return shim.Success(companyJSON)
}

func (hc *HealthcareContract) updateInsuranceCompany(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4.")
	}

	id := args[0]
	newName := args[1]
	newCompany := args[2]

	// Retrieve the insurance company from the ledger
	companyJSON, err := stub.GetState(id)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get insurance company: %s", err.Error()))
	}
	if companyJSON == nil {
		return shim.Error(fmt.Sprintf("Insurance company with ID %s does not exist", id))
	}

	// Update the insurance company's name and company
	company := InsuranceCompany{}
	err = json.Unmarshal(companyJSON, &company)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to unmarshal insurance company: %s", err.Error()))
	}
	company.Name = newName
	company.Company = newCompany

	// Convert the updated insurance company back to JSON
	updatedCompanyJSON, err := json.Marshal(company)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to marshal updated insurance company: %s", err.Error()))
	}

	// Save the updated insurance company to the ledger
	err = stub.PutState(id, updatedCompanyJSON)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to save updated insurance company: %s", err.Error()))
	}

	return shim.Success([]byte("Insurance company updated successfully"))
}

func (hc *HealthcareContract) deleteInsuranceCompany(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	id := args[0]

	// Check if the insurance company exists
	existingCompanyJSON, err := stub.GetState(id)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get insurance company: %s", err.Error()))
	}
	if existingCompanyJSON == nil {
		return shim.Error(fmt.Sprintf("Insurance company with ID %s does not exist", id))
	}

	// Delete the insurance company from the ledger
	err = stub.DelState(id)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to delete insurance company: %s", err.Error()))
	}

	return shim.Success([]byte("Insurance company deleted successfully"))
}
