package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"healthcare/modules"
)



// Define the main function for the healthcare smart contract
func main() {
	err := shim.Start(new(modules.HealthcareContract))
	if err != nil {
		fmt.Printf("Error starting HealthcareContract chaincode: %s", err)
	}
	// fmt.Printf("Hii therere")
}
