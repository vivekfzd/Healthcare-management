package modules


import (
	"encoding/json"
	"strconv"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)


// Define the createPatient method for the healthcare smart contract
func (hc *HealthcareContract) createPatient(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7.")
	}
	args[0] = "P" + args[0]
	exists, err := stub.GetState(args[0])
    if err != nil {
        return shim.Error(err.Error())
    }
    if exists != nil {
        return shim.Error(fmt.Sprintf("the patient is already exists with ID %s", args[0]))
    }

	// var age int
	age,err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error(err.Error())
	}
	
	// Create a new patient object
	patient := Patient{
		ID:          args[0],
		Name:        args[1],
		Age:         age,
		Gender:      args[3],
		BloodType:   args[4],
		Address:     args[5],
		PhoneNumber: args[6],
		RecordIDs:  []string{},
		Read:       []string{},
		Write:      []string{},
	}

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
	
	// Return a success response
	return shim.Success([]byte(fmt.Sprintf("Patient Id %s has been successfuly created.", args[0])))
}



// Define the getPatient method for the healthcare smart contract
func (hc *HealthcareContract) getPatient(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}
	args[0] = "P" + args[0]

	// Retrieve the patient from the ledger
	patientJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if patientJSON == nil {
		return shim.Error("Patient not found.")
	}

	// Unmarshal the patient object from JSON
	var patient Patient
	err = json.Unmarshal(patientJSON, &patient)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Marshal the patient object to JSON
	patientJSON, err = json.Marshal(patient)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return the patient object
	return shim.Success(patientJSON)
}



// Define the updatePatient method for the healthcare smart contract
func (hc *HealthcareContract) updatePatient(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 7 {
	return shim.Error("Incorrect number of arguments. Expecting 7.")
	}
	args[0] = "P" + args[0]
	// Retrieve the patient from the ledger
	patientJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if patientJSON == nil {
		return shim.Error("Patient not found.")
	}

	// Unmarshal the patient object from JSON
	var patient Patient
	err = json.Unmarshal(patientJSON, &patient)
	if err != nil {
		return shim.Error(err.Error())
	}

	// var age int

	age,err := strconv.Atoi(args[2])

	// Update the patient object
	patient.Name = args[1]
	patient.Age = age
	patient.Gender = args[3]
	patient.BloodType = args[4]
	patient.Address = args[5]
	patient.PhoneNumber = args[6]

	// Marshal the patient object to JSON
	patientJSON, err = json.Marshal(patient)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Store the updated patient in the ledger
	err = stub.PutState(patient.ID, patientJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return a success response
	return shim.Success([]byte(fmt.Sprintf("Patient Id %s has successfully updated", args[0])))

}


// Define the deletePatient method for the healthcare smart contract
func (hc *HealthcareContract) deletePatient(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	args[0] = "P" + args[0]
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}
	// Retrieve the patient from the ledger
	patientJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	if patientJSON == nil {
		return shim.Error("Patient not found.")
	}
	// Delete the patient from the ledger
	err = stub.DelState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return a success response
	return shim.Success([]byte(fmt.Sprintf("Patient Id %s has been deteted.", args[0])))
}

//grant read access
//args[0] -> patient id , args[1] -> entity , args[2] -> entity Id
func (hc *HealthcareContract) grantReadAccess(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3.")
	}

	args[0] = "P" + args[0]
	if args[1] == "Doctor" {
		args[2] = "D" + args[2]
	}
	//check that entity is exist or not
	exist, err := stub.GetState(args[2])
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
		if id == args[2] {
			found = true
		}
	}

	if found {
		return shim.Error(fmt.Sprintf("%s Id is also have read access",args[2]))
	} else {
		patient.Read = append(patient.Read,args[2])
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
//args[0] -> patient id , args[1] -> entity , args[2] -> entity Id
func (hc *HealthcareContract) revokeReadAccess(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3.")
	}

	args[0] = "P" + args[0]
	if args[1] == "Doctor" {
		args[2] = "D" + args[2]
	}
	//check that entity is exist or not
	exist, err := stub.GetState(args[2])
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
		if id == args[2] {
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
		return shim.Error(fmt.Sprintf("%s Id is not found, Unable to delete",args[2]))
	}

	// Return a success response
	return shim.Success([]byte(fmt.Sprintf("%s Id has been successfully revoke read access", args[0])))
}




//grant write access
//args[0] -> patient id , args[1] -> entity , args[2] -> entity Id
func (hc *HealthcareContract) grantWriteAccess(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3.")
	}

	args[0] = "P" + args[0]
	if args[1] == "Doctor" {
		args[2] = "D" + args[2]
	}
	//check that entity is exist or not
	exist, err := stub.GetState(args[2])
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
		if id == args[2] {
			found = true 
		}
	}

	if found {
		return shim.Error(fmt.Sprintf("%s Id is also have read access",args[2]))
	} else {
		patient.Write = append(patient.Write,args[2])
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
	return shim.Success([]byte(fmt.Sprintf("%s Id has been successfully revoked", args[0])))
}

//revoke write access
//args[0] -> patient id , args[1] -> entity , args[2] -> entity Id
func (hc *HealthcareContract) revokeWriteAccess(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3.")
	}

	args[0] = "P" + args[0]
	if args[1] == "Doctor" {
		args[2] = "D" + args[2]
	}
	//check that entity is exist or not
	exist, err := stub.GetState(args[2])
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
		if id == args[2] {
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
		return shim.Error(fmt.Sprintf("%s Id is not found, Unable to delete",args[2]))
	}

	// Return a success response
	return shim.Success([]byte(fmt.Sprintf("%s Id has been successfully revoke write access", args[0])))
}