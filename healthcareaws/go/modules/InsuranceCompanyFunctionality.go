package modules

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)



//Read medical records

func (hc *HealthcareContract) insuranceCompanyFunctionality(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	
	function := args[0]

	args = args[1:]

	switch function {
		case "readMedicalRecords":
			return hc.getAllMedicalRecordByPatientIdWithAccess(stub, args)
		default:
			return shim.Error("Patient Function is not valid")
	}
}