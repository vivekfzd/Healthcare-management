package modules

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)


//grant read access
//args[0] -> patient id  , args[1] -> entity Id
func (hc *HealthcareContract) grantReadAccess(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2.")
	}

	//check that entity is exist or not
	exist, err := stub.GetState(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}
	if exist == nil {
		return shim.Error("Entity not found.")
	}
	
	// Retrieve the patient from the ledger
	patientJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if patientJSON == nil {
		return shim.Error("Patient not found.")
	}


	var patient Patient
	err = json.Unmarshal(patientJSON, &patient)
	if err != nil {
		return shim.Error(err.Error())
	}

	var found bool
	for _,id := range patient.Read {
		if id == args[1] {
			found = true
		}
	}

	if found {
		return shim.Error(fmt.Sprintf("%s Id is also have read access",args[1]))
	} else {
		patient.Read = append(patient.Read,args[1])
		// Marshal the patient object to JSON
		patientJSON, err := json.Marshal(patient)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Store the patient in the ledger
		err = stub.PutState(patient.ID, patientJSON)
		if err != nil {
			return shim.Error(err.Error())
		}

	}

	// Return a success response
	return shim.Success([]byte(fmt.Sprintf("%s Id has been successfully grant read access", args[0])))
}

//revoke read access
//args[0] -> patient id , args[1] -> entity Id
func (hc *HealthcareContract) revokeReadAccess(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2.")
	}

	//check that entity is exist or not
	exist, err := stub.GetState(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}
	if exist == nil {
		return shim.Error("Entity not found.")
	}
	
	// Retrieve the patient from the ledger
	patientJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if patientJSON == nil {
		return shim.Error("Patient not found.")
	}


	var patient Patient
	err = json.Unmarshal(patientJSON, &patient)
	if err != nil {
		return shim.Error(err.Error())
	}

	var found bool
	
	for pos,id := range patient.Read {
		if id == args[1] {
			patient.Read[pos] = patient.Read[len(patient.Read)-1]
			found = true		
		}
	}
	

	if found {
		patient.Read = patient.Read[0:len(patient.Read)-1]
		// Marshal the patient object to JSON
		patientJSON, err := json.Marshal(patient)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Store the patient in the ledger
		err = stub.PutState(patient.ID, patientJSON)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else {
		//logic
		return shim.Error(fmt.Sprintf("%s Id is not found, Unable to delete",args[1]))
	}

	// Return a success response
	return shim.Success([]byte(fmt.Sprintf("%s Id has been successfully revoke read access", args[0])))
}




//grant write access
//args[0] -> patient id , args[1] -> entity Id
func (hc *HealthcareContract) grantWriteAccess(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2.")
	}

	
	//check that entity is exist or not
	exist, err := stub.GetState(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}
	if exist == nil {
		return shim.Error("Entity not found.")
	}
	
	// Retrieve the patient from the ledger
	patientJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if patientJSON == nil {
		return shim.Error("Patient not found.")
	}


	var patient Patient
	err = json.Unmarshal(patientJSON, &patient)
	if err != nil {
		return shim.Error(err.Error())
	}

	var found bool
	for _,id := range patient.Write {
		if id == args[1] {
			found = true 
		}
	}

	if found {
		return shim.Error(fmt.Sprintf("%s Id is also have write access",args[1]))
	} else {
		patient.Write = append(patient.Write,args[1])
		// Marshal the patient object to JSON
		patientJSON, err := json.Marshal(patient)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Store the patient in the ledger
		err = stub.PutState(patient.ID, patientJSON)
		if err != nil {
			return shim.Error(err.Error())
		}

	}

	// Return a success response
	return shim.Success([]byte(fmt.Sprintf("%s Id has been successfully inserted", args[0])))
}

//revoke write access
//args[0] -> patient id ,  args[1] -> entity Id
func (hc *HealthcareContract) revokeWriteAccess(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2.")
	}

	//check that entity is exist or not
	exist, err := stub.GetState(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}
	if exist == nil {
		return shim.Error("Entity not found.")
	}
	
	// Retrieve the patient from the ledger
	patientJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if patientJSON == nil {
		return shim.Error("Patient not found.")
	}


	var patient Patient
	err = json.Unmarshal(patientJSON, &patient)
	if err != nil {
		return shim.Error(err.Error())
	}

	var found bool
	
	for pos,id := range patient.Write {
		if id == args[1] {
			patient.Write[pos] = patient.Write[len(patient.Write)-1]
			found = true		
		}
	}
	

	if found {
		patient.Write = patient.Write[0:len(patient.Write)-1]
		// Marshal the patient object to JSON
		patientJSON, err := json.Marshal(patient)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Store the patient in the ledger
		err = stub.PutState(patient.ID, patientJSON)
		if err != nil {
			return shim.Error(err.Error())
		}

	} else {
		//logic
		return shim.Error(fmt.Sprintf("%s Id is not found, Unable to delete",args[1]))
	}

	// Return a success response
	return shim.Success([]byte(fmt.Sprintf("%s Id has been successfully revoke write access", args[0])))
}