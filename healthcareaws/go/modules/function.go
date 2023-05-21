package modules

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// Define the Init method for the healthcare smart contract
func (hc *HealthcareContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("Healthcare chaincode is successfully initialized"))
}

// Define the Invoke method for the healthcare smart contract
func (hc *HealthcareContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	entity, args := stub.GetFunctionAndParameters()

	switch entity {
		case "Admin":
			return hc.adminFunctionality(stub, args)
		case "Patient":
			return hc.patientFunctionality(stub, args)
		case "Doctor","LabTechnicain":
			return hc.doctorLabTechnicianFunctionality(stub, args)
		case "InsuranceCompany":
			return hc.insuranceCompanyFunctionality(stub, args)
		default:
			return shim.Error("Entity is not valid")
	}
}





