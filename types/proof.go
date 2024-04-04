package types

import (
	"github.com/iden3/go-circuits/v2"
)

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

type ProofQuery struct {
	SubjectFieldName  string `json:"subjectFieldName"`
	Operator          int    `json:"operator"`
	SubjectFieldValue string `json:"subjectFieldValue"`

	Type []string `json:"type"`
}

type CreateProofRequest struct {
	Id        string
	CircuitId circuits.CircuitID
	Challenge string
	Query     ProofQuery
}
