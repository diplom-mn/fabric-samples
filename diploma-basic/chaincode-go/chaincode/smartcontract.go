package chaincode

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

type IssuerVerification struct {
	AllowedOrigins string `json:"AllowedOrigins"`
}

type Issuer struct {
	ID           string             `json:"id"`
	Type         string             `json:"type"`
	Name         string             `json:"name"`
	Url          string             `json:"url"`
	Email        string             `json:"email"`
	Verification IssuerVerification `json:"verification"`
}

type Recipient struct {
	Hashed   bool   `json:"hashed"`
	Identity string `json:"identity"`
	Type     string `json:"type"`
}

type Criteria struct {
	Narrative string `json:"narrative"`
}

type Badge struct {
	Type        string   `json:"type"`
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Criteria    Criteria `json:"criteria"`
	Issuer      Issuer   `json:"issuer"`
}

type OpenBadgeV2VerificationType struct {
	Type string `json:"type"`
}

// OpenBadge 2.0 schema
// https://www.imsglobal.org/sites/default/files/Badges/OBv2p0Final/index.html
type Certificate struct {
	ID           string                      `json:"id"`
	Context      string                      `json:"@context"`
	Type         string                      `json:"type"`
	Recipient    Recipient                   `json:"recipient"`
	IssuedOn     string                      `json:"issuedOn"`
	Verification OpenBadgeV2VerificationType `json:"verification"`
	Badge        Badge                       `json:"badge"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	return nil
}

func (s *SmartContract) CreateCertificate(ctx contractapi.TransactionContextInterface, id string, certType string, badgeId string, badgeType string, badgeName string, badgeDesc string, criteriaNarrative string, recipientType string, recipientId string, issuerId string, issuerType string, issuerName string, issuerUrl string, issuerEmail string, issuerVerificationAllowedOrigins string) error {
	exists, err := s.CertificateExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the certificate %s already exists", id)
	}

	txTimestamp, _ := ctx.GetStub().GetTxTimestamp()

	var issuedOn string
	if txTimestamp != nil {
		issuedOn = time.Unix(txTimestamp.Seconds, int64(txTimestamp.Nanos)).String()
	}

	certificate := Certificate{
		ID:      id,
		Context: "https://w3id.org/openbadges/v2",
		Type:    certType,
		Recipient: Recipient{
			Type:     recipientType,
			Identity: recipientId,
		},
		IssuedOn: issuedOn,
		Verification: OpenBadgeV2VerificationType{
			Type: "hosted",
		},
		Badge: Badge{
			Type:        badgeType,
			ID:          badgeId,
			Name:        badgeName,
			Description: badgeDesc,
			Criteria: Criteria{
				Narrative: criteriaNarrative,
			},
			Issuer: Issuer{
				ID:    issuerId,
				Type:  issuerType,
				Name:  issuerName,
				Url:   issuerUrl,
				Email: issuerEmail,
				Verification: IssuerVerification{
					AllowedOrigins: issuerVerificationAllowedOrigins,
				},
			},
		},
	}
	certificateJSON, err := json.Marshal(certificate)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, certificateJSON)
}

func (s *SmartContract) ReadCertificate(ctx contractapi.TransactionContextInterface, id string) (*Certificate, error) {
	certificateJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if certificateJSON == nil {
		return nil, fmt.Errorf("the certificate %s does not exist", id)
	}

	var certificate Certificate
	err = json.Unmarshal(certificateJSON, &certificate)
	if err != nil {
		return nil, err
	}

	return &certificate, nil
}

func (s *SmartContract) CertificateExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	certificateJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return certificateJSON != nil, nil
}
