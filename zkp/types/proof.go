package types

import (
	"github.com/iden3/go-circuits/v2"
)

// ClaimOffer Data to retrieve from issuer to generate AuthProof and get Verifiable Credentials
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

// ProofQuery Data to build circuits.Query
type ProofQuery struct {
	SubjectFieldName  string `json:"subjectFieldName"`
	Operator          int    `json:"operator"`
	SubjectFieldValue string `json:"subjectFieldValue"`

	Type []string `json:"type"`
}

// CreateProofRequest Data to fill inputs for ZkpGen.GenerateProof
type CreateProofRequest struct {
	Id        string
	CircuitId circuits.CircuitID
	Challenge string
	Query     ProofQuery
}
