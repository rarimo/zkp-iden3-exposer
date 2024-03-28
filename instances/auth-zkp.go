package instances

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rarimo/go-jwz"
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
func (a *AuthZkp) GetVerifiableCredentials(claimOffer types.ClaimOffer) (*types.W3CCredential, error) {
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

	claimDetails := ClaimDetails{}

	guid, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	if claimDetails.Id = claimOffer.Id; claimDetails.Id == "" {
		claimDetails.Id = guid.String()
	}

	if claimDetails.Typ = *claimOffer.Typ; claimDetails.Typ == "" {
		fallbackTyp := "application/iden3-zkp-json"
		claimDetails.Typ = fallbackTyp
	}

	if claimDetails.Type = claimOffer.Type; claimDetails.Type == "" {
		fallbackType := "https://iden3-communication.io/credentials/1.0/fetch-request"
		claimDetails.Type = fallbackType
	}

	if claimDetails.ThreadID = *claimOffer.ThreadID; claimDetails.ThreadID == "" {
		claimDetails.ThreadID = guid.String()
	}

	claimDetailsBody := ClaimDetailsBody{}

	if claimDetailsBody.Id = claimOffer.Body.Credentials[0].Id; claimDetailsBody.Id == "" {
		return nil, errors.New("claimOffer.Body is nil")
	}

	// TODO: explain why from = to, and to = from
	if claimDetails.From = claimOffer.To; claimDetails.From == "" {
		return nil, errors.New("claimOffer.To is nil")
	}

	if claimDetails.To = claimOffer.From; claimDetails.To == "" {
		return nil, errors.New("claimOffer.From is nil")
	}

	claimDetailsBody.Id = claimOffer.Body.Credentials[0].Id

	claimDetailsJson, err := json.Marshal(claimDetails)

	if err != nil {
		return nil, err
	}

	token, err := jwz.NewWithPayload(jwz.ProvingMethodGroth16AuthV2Instance, claimDetailsJson, a.Identity.PrepareAuthV2Inputs)

	if err != nil {
		return nil, err
	}
}
