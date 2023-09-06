package chaincode_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dms/diploma-basic/chaincode-go/chaincode"
	"github.com/dms/diploma-basic/chaincode-go/chaincode/mocks"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/stretchr/testify/require"
)

//go:generate counterfeiter -o mocks/transaction.go -fake-name TransactionContext . transactionContext
type transactionContext interface {
	contractapi.TransactionContextInterface
}

//go:generate counterfeiter -o mocks/chaincodestub.go -fake-name ChaincodeStub . chaincodeStub
type chaincodeStub interface {
	shim.ChaincodeStubInterface
}

//go:generate counterfeiter -o mocks/statequeryiterator.go -fake-name StateQueryIterator . stateQueryIterator
type stateQueryIterator interface {
	shim.StateQueryIteratorInterface
}

func TestCreateCertificate(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	diplomaBasic := chaincode.SmartContract{}

	err := diplomaBasic.InitLedger(transactionContext)
	require.NoError(t, err)

	err = diplomaBasic.CreateCertificate(transactionContext,
		"https://example.org/assertions/123",
		"Assertion",
		"https://example.org/badges/5",
		"BadgeClass",
		"Bachelor of Computer Science",
		"Bachelor of Computer Science Description",
		"Has satisfactorily completed all degree requirements from 2017 to 2020.",
		"email",
		"battulga.dev@gmail.com",
		"https://example.org/issuer",
		"Profile",
		"Example ORG",
		"https://example.org",
		"contact@example.org",
		"example.org",
	)

	require.NoError(t, err)

	chaincodeStub.GetStateReturns([]byte{}, nil)
	err = diplomaBasic.CreateCertificate(transactionContext,
		"https://example.org/assertions/123",
		"Assertion",
		"https://example.org/badges/5",
		"BadgeClass",
		"Bachelor of Computer Science",
		"Bachelor of Computer Science Description",
		"Has satisfactorily completed all degree requirements from 2017 to 2020.",
		"email",
		"battulga.dev@gmail.com",
		"https://example.org/issuer",
		"Profile",
		"Example ORG",
		"https://example.org",
		"contact@example.org",
		"example.org")
	require.EqualError(t, err, "the certificate https://example.org/assertions/123 already exists")

	chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve asset"))
	err = diplomaBasic.CreateCertificate(transactionContext, "https://example.org/assertions/123",
		"Assertion",
		"https://example.org/badges/5",
		"BadgeClass",
		"Bachelor of Computer Science",
		"Bachelor of Computer Science Description",
		"Has satisfactorily completed all degree requirements from 2017 to 2020.",
		"email",
		"battulga.dev@gmail.com",
		"https://example.org/issuer",
		"Profile",
		"Example ORG",
		"https://example.org",
		"contact@example.org",
		"example.org")
	require.EqualError(t, err, "failed to read from world state: unable to retrieve asset")
}

func TestReadAsset(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	expectedAsset := &chaincode.Certificate{ID: "https://example.org/assertions/123"}
	bytes, err := json.Marshal(expectedAsset)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(bytes, nil)
	diplomaBasic := chaincode.SmartContract{}
	asset, err := diplomaBasic.ReadCertificate(transactionContext, "")
	require.NoError(t, err)
	require.Equal(t, expectedAsset, asset)

	chaincodeStub.GetStateReturns(nil, fmt.Errorf("unable to retrieve asset"))
	_, err = diplomaBasic.ReadCertificate(transactionContext, "")
	require.EqualError(t, err, "failed to read from world state: unable to retrieve asset")

	chaincodeStub.GetStateReturns(nil, nil)
	asset, err = diplomaBasic.ReadCertificate(transactionContext, "asset1")
	require.EqualError(t, err, "the certificate asset1 does not exist")
	require.Nil(t, asset)
}
