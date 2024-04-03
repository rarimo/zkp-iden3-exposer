package types

import (
	"encoding/json"
	"github.com/iden3/go-circuits/v2"
)

// custom
type ClaimOffer struct {
	Body struct {
		Credentials []struct {
			Description string `json:"description"`
			Id          string `json:"id"`
		} `json:"Credentials"`
		Url string `json:"url"`
	} `json:"body"`
	From     string  `json:"from"`
	Id       string  `json:"id"`
	ThreadID *string `json:"threadID"`
	To       string  `json:"to"`
	Typ      *string `json:"typ"`
	Type     string  `json:"type"`
}

// custom
type ProofQuery struct {
	AllowedIssuers      []string        `json:"allowedIssuers"`
	CredentialSubject   json.RawMessage `json:"credentialSubject"`
	Schema              string          `json:"schema"`
	ClaimId             string          `json:"claimId"`
	CredentialSubjectId string          `json:"credentialSubjectId"`
	Context             string          `json:"context"`
	Type                []string        `json:"type"`
}

// custom
type CreateProofRequest struct {
	Id             *int
	AccountAddress *string
	CircuitId      circuits.CircuitID
	Challenge      *string // bigint string
	Query          ProofQuery
}
