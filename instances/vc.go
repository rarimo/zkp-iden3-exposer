package instances

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/iden3/go-circuits/v2"
	"github.com/pkg/errors"
	"github.com/rarimo/go-jwz"
	"github.com/rarimo/zkp-iden3-exposer/types"
)

func GetClaimDetailsJson(
	identity Identity,
	claimOffer types.ClaimOffer,
) ([]byte, error) {
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

	return claimDetailsJson, nil
}

func GetAuthV2Inputs(
	identity Identity,
	claimOffer types.ClaimOffer,
) ([]byte, error) {
	claimDetailsJson, err := GetClaimDetailsJson(identity, claimOffer)

	if err != nil {
		return nil, errors.Wrap(err, "Error getting claim details")
	}

	preparer := jwz.ProofInputsPreparerHandlerFunc(func(hash []byte, circuitID circuits.CircuitID) ([]byte, error) {
		return identity.PrepareAuthV2Inputs(hash, circuitID)
	})

	token, err := jwz.NewWithPayload(
		jwz.ProvingMethodGroth16AuthV2Instance,
		claimDetailsJson,
		preparer,
	)

	messageHash, err := token.GetMessageHash()

	if err != nil {
		return nil, errors.Wrap(err, "Error getting message hash")
	}

	authV2Inputs, err := identity.PrepareAuthV2Inputs(messageHash, circuits.AuthV2CircuitID)

	if err != nil {
		return nil, errors.Wrap(err, "Error getting AuthV2Inputs")
	}

	return authV2Inputs, nil
}

//func GetVerifiableCredentials(
//	identity Identity,
//	claimOffer types.ClaimOffer,
//	circuitsPair types.CircuitPair,
//) (*overrides.W3CCredential, error) {
//	claimDetailsJson, err := GetClaimDetailsJson(identity, claimOffer)
//
//	if err != nil {
//		return nil, err
//	}
//
//	preparer := jwz.ProofInputsPreparerHandlerFunc(func(hash []byte, circuitID circuits.CircuitID) ([]byte, error) {
//		return identity.PrepareAuthV2Inputs(hash, circuitID)
//	})
//
//	token, err := jwz.NewWithPayload(
//		jwz.ProvingMethodGroth16AuthV2Instance,
//		claimDetailsJson,
//		preparer,
//	)
//
//	if err != nil {
//		return nil, err
//	}
//
//	jwzTokenRaw, err := token.Prove(circuitsPair.ProvingKey, circuitsPair.Wasm)
//
//	if err != nil {
//		return nil, err
//	}
//
//	response, err := http.Post(claimOffer.Body.Url, "application/json", strings.NewReader(jwzTokenRaw))
//
//	if err != nil {
//		return nil, errors.Wrap(err, "failed to post")
//	}
//
//	if response.StatusCode != http.StatusOK {
//		return nil, errors.New("response status is not OK")
//	}
//
//	defer response.Body.Close()
//
//	type AgentResponse struct {
//		Body struct {
//			Credential overrides.W3CCredential `json:"credential"`
//		} `json:"body"`
//	}
//
//	agentResponse := AgentResponse{}
//
//	if err := json.NewDecoder(response.Body).Decode(&agentResponse); err != nil {
//		return nil, errors.Wrap(err, "failed to unmarshal")
//	}
//
//	agentResponse.Body.Credential.W3CCredential.Proof = verifiable.CredentialProofs(agentResponse.Body.Credential.Proof)
//
//	return &agentResponse.Body.Credential, nil
//}
