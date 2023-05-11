package modules

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)


//create medical records

//Read medical records

func (hc *HealthcareContract) doctorLabTechnicianFunctionality(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	
	function := args[0]

	args = args[1:]

	switch function {
		case "createMedicalRecord":
			return hc.createMedicalRecord(stub, args)
		case "readMedicalRecords":
			return hc.getAllMedicalRecordByPatientIdWithAccess(stub, args)
		default:
			return shim.Error("Patient Function is not valid")
	}
}