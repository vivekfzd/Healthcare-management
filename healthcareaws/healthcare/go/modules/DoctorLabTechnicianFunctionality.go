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
		case "createMedicalRecord2":
			return hc.createMedicalRecord2(stub,args)
		case "getMedicalRecordById":
			return hc.getMedicalRecordById(stub, args)
		case "readMedicalRecords":
			return hc.getAllMedicalRecordByPatientIdWithAccess(stub, args)
		default:
			return shim.Error("Entity Function is not valid")
	}
}