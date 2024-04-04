package zkp_iden3

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/iden3/go-circuits/v2"
	"github.com/iden3/go-jwz/v2"
	"github.com/iden3/go-schema-processor/v2/verifiable"
	"github.com/pkg/errors"
	"github.com/rarimo/zkp-iden3-exposer/overrides"
	"github.com/rarimo/zkp-iden3-exposer/types"
	"net/http"
	"strings"
)

type AuthZkpConfig struct {
	ChainInfo types.ChainZkpInfo
	// TODO: mb use files straight away
	Circuits types.CircuitPair
}

type AuthZkp struct {
	Config   AuthZkpConfig
	Identity Identity
}

func NewAuthZkp(config AuthZkpConfig, identity Identity) *AuthZkp {
	return &AuthZkp{
		Config:   config,
		Identity: identity,
	}
}

func (a *AuthZkp) GetVerifiableCredentials(claimOffer types.ClaimOffer) (*overrides.W3CCredential, error) {
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

	//if claimDetails.Type = claimOffer.Type; claimDetails.Type == "" {
	//	fallbackType := "https://iden3-communication.io/credentials/1.0/fetch-request"
	//	claimDetails.Type = fallbackType
	//}
	claimDetails.Type = "https://iden3-communication.io/credentials/1.0/fetch-request"

	if claimDetails.ThreadID = *claimOffer.ThreadID; claimDetails.ThreadID == "" {
		claimDetails.ThreadID = guid.String()
	}

	claimDetailsBody := ClaimDetailsBody{}

	if claimDetailsBody.Id = claimOffer.Body.Credentials[0].Id; claimDetailsBody.Id == "" {
		return nil, errors.New("claimOffer.Body is nil")
	}

	claimDetails.Body = claimDetailsBody

	// TODO: explain why from = to, and to = from
	if claimDetails.From = claimOffer.To; claimDetails.From == "" {
		return nil, errors.New("claimOffer.To is nil")
	}

	if claimDetails.To = claimOffer.From; claimDetails.To == "" {
		return nil, errors.New("claimOffer.From is nil")
	}

	claimDetailsJson, err := json.Marshal(claimDetails)

	if err != nil {
		return nil, err
	}

	preparer := jwz.ProofInputsPreparerHandlerFunc(func(hash []byte, circuitID circuits.CircuitID) ([]byte, error) {
		return a.Identity.PrepareAuthV2Inputs(hash, circuitID)
	})

	token, err := jwz.NewWithPayload(
		jwz.ProvingMethodGroth16AuthV2Instance,
		claimDetailsJson,
		preparer,
	)

	if err != nil {
		return nil, err
	}

	jwzTokenRaw, err := token.Prove(a.Config.Circuits.ProvingKey, a.Config.Circuits.Wasm)

	if err != nil {
		return nil, err
	}

	response, err := http.Post(claimOffer.Body.Url, "application/json", strings.NewReader(jwzTokenRaw))

	if err != nil {
		return nil, errors.Wrap(err, "failed to post")
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("response status is not OK")
	}

	defer response.Body.Close()

	type AgentResponse struct {
		Body struct {
			Credential overrides.W3CCredential `json:"credential"`
		} `json:"body"`
	}

	agentResponse := AgentResponse{}

	if err := json.NewDecoder(response.Body).Decode(&agentResponse); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	agentResponse.Body.Credential.W3CCredential.Proof = verifiable.CredentialProofs(agentResponse.Body.Credential.Proof)

	return &agentResponse.Body.Credential, nil
}
