package instances

import (
	"github.com/google/uuid"
	"github.com/rarimo/zkp-iden3-exposer/types"
)

type AuthZkpConfig struct {
	ChainInfo types.ChainZkpInfo
	// TODO: mb use files straight away
	CircuitUrls struct {
		wasmUrl string
		zkeyUrl string
	}
}

type AuthZkp struct {
	Config   AuthZkpConfig
	Identity Identity
}

// TODO: create W3Credential type
func (a *AuthZkp) GetVerifiableCredentials(claimOffer types.ClaimOffer) (*string, error) {
	type ClaimDetailsBody struct {
		Id string `json:"id"`
	}

	type ClaimDetails struct {
		Id       string           `json:"id"`
		Typ      string           `json:"typ"`
		Type     string           `json:"type"`
		ThreadID string           `json:"threadID"`
		Body     ClaimDetailsBody `json:"body"`
		From     string           `json:"from"`
		To       string           `json:"to"`
	}

	guid, err := uuid.NewRandom()

	if err != nil {
		return nil, error
	}

	claimOfferId := claimOffer.Id

	if claimOfferId == "" {
		claimOfferId, err = uuid.NewRandom()
	}

	claimDetailsBody := ClaimDetailsBody{}

	claimDetailsBody.Id = claimOffer.Body.Credentials[0].Id

	claimDetails := ClaimDetails{
		Id:       claimOffer.Id,
		Typ:      *claimOffer.Typ,
		Type:     claimOffer.Type,
		ThreadID: *claimOffer.ThreadID,
		Body:     claimDetailsBody,

		// TODO: explain why
		From: claimOffer.To,
		To:   claimOffer.From,
	}
}
