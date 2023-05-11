package modules

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)


// CRUD of all entity
func (hc *HealthcareContract) adminFunctionality(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	
	function := args[0]

	args = args[1:]

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
		case "createLabTechnician":
			return hc.createLabTechnician(stub,args)
		case "getLabTechnician":
			return hc.createLabTechnician(stub,args)
		case "deleteLabTechnician":
			return hc.createLabTechnician(stub,args)
		case "createInsuranceCompany":
			return hc.createInsuranceCompany(stub,args)
		case "getInsuranceCompany":
			return hc.getInsuranceCompany(stub,args)
		case "updateInsuranceCompany":
			return hc.updateInsuranceCompany(stub,args)
		case "deleteInsuranceCompany":
			return hc.deleteInsuranceCompany(stub,args)
		default:
			return shim.Error("Admin Function is not valid")
	}
}