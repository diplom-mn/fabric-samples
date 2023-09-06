/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/dms/diploma-basic/chaincode-go/chaincode"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	diplomaBasicChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating diploma-basic chaincode: %v", err)
	}

	if err := diplomaBasicChaincode.Start(); err != nil {
		log.Panicf("Error starting diploma-basic chaincode: %v", err)
	}
}
