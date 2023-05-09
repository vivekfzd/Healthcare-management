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
	function, args := stub.GetFunctionAndParameters()

	switch function {
		case "createPatient":
			return hc.createPatient(stub, args)
		case "getPatient":
			return hc.getPatient(stub, args)
		case "updatePatient":
			return hc.updatePatient(stub, args)
		case "deletePatient":
			return hc.deletePatient(stub, args)
		case "createDoctor":
			return hc.createDoctor(stub, args)
		case "getDoctor":
			return hc.getDoctor(stub, args)
		case "updateDoctor":
			return hc.updateDoctor(stub, args)
		case "deleteDoctor":
			return hc.deleteDoctor(stub, args)
		case "createMedicalRecord":
			return hc.createMedicalRecord(stub,args)
		case "getAllMedicalRecordByPatientId":
			return hc.getAllMedicalRecordByPatientId(stub,args)
		case "getMedicalRecordById":
			return hc.getMedicalRecordById(stub,args)
		case "deleteMedicalRecord":
			return hc.deleteMedicalRecord(stub,args)
		case "grantReadAccess":
			return hc.grantReadAccess(stub,args)
		case "revokeReadAccess":
			return hc.revokeReadAccess(stub,args)
		case "grantWriteAccess":
			return hc.grantWriteAccess(stub,args)
		case "revokeWriteAccess":
			return hc.revokeWriteAccess(stub,args)
		case "getAllMedicalRecordByPatientIdWithAccess":
			return hc.getAllMedicalRecordByPatientIdWithAccess(stub,args)
		default:
			return shim.Error("Invalid function name.")
	}
}





