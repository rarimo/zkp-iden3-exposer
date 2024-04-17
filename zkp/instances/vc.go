package instances

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/iden3/go-circuits/v2"
	types2 "github.com/iden3/go-rapidsnark/types"
	"github.com/iden3/go-schema-processor/v2/verifiable"
	"github.com/pkg/errors"
	"github.com/rarimo/go-jwz"
	"github.com/rarimo/zkp-iden3-exposer/zkp/overrides"
	"github.com/rarimo/zkp-iden3-exposer/zkp/types"
	"net/http"
	"strings"
)

func GetClaimDetailsJson(claimOffer types.ClaimOffer) ([]byte, error) {
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
	claimDetailsJson, err := GetClaimDetailsJson(claimOffer)

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

func GetJWZToken(
	identity Identity,
	claimDetailsJson []byte,
	proofRaw []byte,
) (*string, error) {
	preparer := jwz.ProofInputsPreparerHandlerFunc(func(hash []byte, circuitID circuits.CircuitID) ([]byte, error) {
		return identity.PrepareAuthV2Inputs(hash, circuitID)
	})

	token, err := jwz.NewWithPayload(
		jwz.ProvingMethodGroth16AuthV2Instance,
		claimDetailsJson,
		preparer,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating token")
	}

	headers, err := json.Marshal(token.Raw.Header)
	if err != nil {
		return nil, errors.Wrap(err, "Error marshalling headers")
	}
	token.Raw.Protected = headers

	proof := types2.ZKProof{}
	if err := json.Unmarshal(proofRaw, &proof); err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling proof")
	}

	token.ZkProof = &proof
	token.Raw.ZKP = proofRaw

	jwzToken, err := token.CompactSerialize()
	if err != nil {
		return nil, errors.Wrap(err, "Error serializing token")
	}

	return &jwzToken, nil
}

func LoadVC(url string, jwzToken string) (*overrides.W3CCredential, error) {
	response, err := http.Post(url, "application/json", strings.NewReader(jwzToken))

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
