package modules

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

//Read medical records

//grant access

//revoke access

//delete medical records

func (hc *HealthcareContract) patientFunctionality(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	
	function := args[0]

	args = args[1:]

	switch function {
		case "readMedicalRecords":
			return hc.getAllMedicalRecordByPatientId(stub, args)
		case "grantReadAccess":
			return hc.grantReadAccess(stub, args)
		case "revokeReadAccess":
			return hc.revokeReadAccess(stub, args)
		case "grantWriteAccess":
			return hc.grantReadAccess(stub, args)
		case "revokeWriteAccess":
			return hc.revokeWriteAccess(stub, args)
		case "deleteMedicalRecords":
			return hc.deleteMedicalRecord(stub, args)
		default:
			return shim.Error("Patient Function is not valid")
	}
}