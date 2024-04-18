package zkp_iden3_exposer

import (
	"encoding/hex"
	"encoding/json"
	"github.com/iden3/go-circuits/v2"
	core "github.com/iden3/go-iden3-core/v2"
	"github.com/iden3/go-iden3-core/v2/w3c"
	"github.com/iden3/go-schema-processor/v2/verifiable"
	"github.com/pkg/errors"
	"github.com/rarimo/go-jwz"
	"github.com/rarimo/zkp-iden3-exposer/client"
	"github.com/rarimo/zkp-iden3-exposer/wallet"
	"github.com/rarimo/zkp-iden3-exposer/zkp/instances"
	"github.com/rarimo/zkp-iden3-exposer/zkp/overrides"
	"github.com/rarimo/zkp-iden3-exposer/zkp/types"
	"net/http"
)

type Connector struct {
	ChainId     string `json:"chainId"`
	Denom       string `json:"denom"`
	Addr        string `json:"addr"`
	MinGasPrice uint64 `json:"minGasPrice"`
	GasLimit    uint64 `json:"gasLimit"`
	TLS         bool   `json:"tls"`
}

func NewConnector(chainId string, denom string, addr string, minGasPrice uint64, gasLimit uint64, tls bool) *Connector {
	return &Connector{
		ChainId:     chainId,
		Denom:       denom,
		Addr:        addr,
		MinGasPrice: minGasPrice,
		GasLimit:    gasLimit,
		TLS:         tls,
	}
}

func getIdentity(identityConfig []byte) (*instances.Identity, error) {
	config := types.IdentityConfig{}

	if err := json.Unmarshal(identityConfig, &config); err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling identity config")
	}

	if config.PkHex == "" || &config.PkHex == nil {
		return nil, errors.New("Private key is required")
	}

	identity, err := instances.NewIdentity(instances.IdentityConfig{
		IdType:        config.IdType,
		SchemaHashHex: config.SchemaHashHex,
		ChainInfo: types.ChainZkpInfo{
			TargetChainId:              config.TargetChainId,
			TargetRpcUrl:               config.TargetRpcUrl,
			TargetStateContractAddress: config.TargetStateContractAddress,

			CoreApiUrl:               config.CoreApiUrl,
			CoreEvmRpcApiUrl:         config.CoreEvmRpcApiUrl,
			CoreStateContractAddress: config.CoreStateContractAddress,
		},
	}, &config.PkHex)

	if err != nil {
		return nil, errors.Wrap(err, "Error creating identity")
	}

	return identity, nil
}

func (c *Connector) GetOfferJson(issuerApi string, identityDidString string, claimType string) ([]byte, error) {
	offer := types.ClaimOffer{}

	response, err := http.Get(issuerApi + "/v1/credentials/" + identityDidString + "/" + claimType)

	if err != nil {
		return nil, errors.Wrap(err, "Error getting offer")
	}

	if err := json.NewDecoder(response.Body).Decode(&offer); err != nil {
		return nil, errors.Wrap(err, "Error decoding offer")
	}

	offerJson, err := json.Marshal(offer)

	if err != nil {
		return nil, errors.Wrap(err, "Error marshalling offer")
	}

	return offerJson, nil
}

func (c *Connector) GetDidString(identityConfig []byte) (string, error) {
	identity, err := getIdentity(identityConfig)

	if err != nil {
		return "", errors.Wrap(err, "Error getting identity")
	}

	return identity.DID.String(), nil
}

func (c *Connector) GetIdBigIntString(identityConfig []byte) (string, error) {
	identity, err := getIdentity(identityConfig)

	if err != nil {
		return "", errors.Wrap(err, "Error getting identity")
	}

	id, err := identity.ID()

	if err != nil {
		return "", nil
	}

	return id.BigInt().String(), nil
}

func (c *Connector) GetAuthV2Inputs(
	identityConfig []byte,
	offerJson []byte,
) ([]byte, error) {
	identity, err := getIdentity(identityConfig)

	if err != nil {
		return nil, errors.Wrap(err, "Error getting identity")
	}

	offer := types.ClaimOffer{}

	if err := json.Unmarshal(offerJson, &offer); err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling offer")
	}

	claimDetailsJson, err := instances.GetClaimDetailsJson(offer)

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

func (c *Connector) GetVC(
	identityConfig []byte,
	offerJson []byte,
	proofRaw []byte,
) ([]byte, error) {
	identity, err := getIdentity(identityConfig)

	if err != nil {
		return nil, errors.Wrap(err, "Error getting identity")
	}

	offer := types.ClaimOffer{}

	if err := json.Unmarshal(offerJson, &offer); err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling offer")
	}

	claimDetailsJson, err := instances.GetClaimDetailsJson(offer)

	jwzToken, err := instances.GetJWZToken(*identity, claimDetailsJson, proofRaw)

	if err != nil {
		return nil, errors.Wrap(err, "Error getting JWZ token")
	}

	vc, err := instances.LoadVC(offer.Body.Url, *jwzToken)

	if err != nil {
		return nil, errors.Wrap(err, "Error loading VC")
	}

	vcJson, err := json.Marshal(vc)

	if err != nil {
		return nil, errors.Wrap(err, "Error marshalling VC")
	}

	return vcJson, nil
}

func (c *Connector) GetAtomicQueryMTVV2OnChainInputs(
	identityConfig []byte,
	jsonVC []byte,

	circuitId string,
	challenge string,

	subjectFieldName string,
	subjectFieldValue string,
	operator int,
) ([]byte, error) {
	identity, err := getIdentity(identityConfig)

	if err != nil {
		return nil, errors.Wrap(err, "Error getting identity")
	}

	proofRequest := types.CreateProofRequest{
		CircuitId: circuits.CircuitID(circuitId),
		Challenge: challenge,
		Query: types.ProofQuery{
			SubjectFieldName:  subjectFieldName,
			SubjectFieldValue: subjectFieldValue,
			Operator:          operator,
		},
	}

	vc := overrides.W3CCredential{}

	if err := json.Unmarshal(jsonVC, &vc); err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling vc")
	}

	vc.W3CCredential.Proof = verifiable.CredentialProofs(vc.Proof)

	type StateInfo struct {
		Index                    string `json:"index"`
		Hash                     string `json:"hash"`
		CreatedAtTimestamp       string `json:"createdAtTimestamp"`
		CreatedAtBlock           string `json:"createdAtBlock"`
		LastUpdateOperationIndex string `json:"lastUpdateOperationIndex"`
	}

	type StateInfoResponse struct {
		State StateInfo `json:"state"`
	}

	issuerDID, _ := w3c.ParseDID(vc.Issuer)

	issuerID, _ := core.IDFromDID(*issuerDID)

	issuerHexId := "0x" + hex.EncodeToString(issuerID.BigInt().Bytes())

	response, _ := http.Get(identity.Config.ChainInfo.CoreApiUrl + "/rarimo/rarimo-core/identity/state/" + issuerHexId)

	stateInfoResponse := StateInfoResponse{}

	if err := json.NewDecoder(response.Body).Decode(&stateInfoResponse); err != nil {
		return nil, errors.Wrap(err, "Error decoding response")
	}

	type OperationStatus string

	const (
		Signed      OperationStatus = "SIGNED"
		Initialized OperationStatus = "INITIALIZED"
		Approved    OperationStatus = "APPROVED"
		NotApproved OperationStatus = "NOT_APPROVED"
	)

	type Operation struct {
		Index         string `json:"index"`
		OperationType string `json:"operationType"`
		Details       struct {
			AtType        string `json:"@type"`
			Contract      string `json:"contract"`
			Chain         string `json:"chain"`
			GISTHash      string `json:"GISTHash"`
			StateRootHash string `json:"stateRootHash"`
			Timestamp     string `json:"timestamp"`
		} `json:"details"`
		Status    OperationStatus `json:"status"`
		Creator   string          `json:"creator"`
		Timestamp string          `json:"timestamp"`
	}

	type OperationResponse struct {
		Operation Operation `json:"operation"`
	}

	response, _ = http.Get(identity.Config.ChainInfo.CoreApiUrl + "/rarimo/rarimo-core/rarimocore/operation/" + stateInfoResponse.State.LastUpdateOperationIndex)

	operationResponse := OperationResponse{}

	if err := json.NewDecoder(response.Body).Decode(&operationResponse); err != nil {
		return nil, errors.Wrap(err, "Error decoding operation response")
	}

	atomicQueryMTPV2OnChainProof := instances.NewAtomicQueryMTPV2OnChainProof(
		*identity,

		stateInfoResponse.State.Hash,
		operationResponse.Operation.Details.GISTHash,
		vc,
		proofRequest,
	)

	inputs, err := atomicQueryMTPV2OnChainProof.GetInputs()

	if err != nil {
		return nil, errors.Wrap(err, "Error getting inputs")
	}

	return inputs, nil
}

func (c *Connector) WalletGetAddress(pk string, addressPrefix string) (string, error) {
	w, err := wallet.NewWallet(pk, addressPrefix)
	if err != nil {
		return "", errors.Wrap(err, "Error creating wallet")
	}

	return w.Address, nil
}

func (c *Connector) WalletSignDirect(pk string, addressPrefix string, signDoc []byte) ([]byte, error) {
	w, err := wallet.NewWallet(pk, addressPrefix)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating wallet")
	}

	signedBytes, err := w.SignDirect(w.Address, signDoc)
	if err != nil {
		return nil, errors.Wrap(err, "Error signing")
	}

	return signedBytes, nil
}

func (c *Connector) WalletSend(pk string, addressPrefix, fromAddr, toAddr string, amount int64, denom string) ([]byte, error) {
	w, err := wallet.NewWallet(pk, addressPrefix)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating wallet")
	}

	rarimoClient, err := client.NewClient(
		client.ChainConfig{
			ChainId:     "rarimo_42-1",
			Denom:       "stake",
			Addr:        "104.196.227.66:9090",
			MinGasPrice: 0,
			GasLimit:    1000000,
			TLS:         true,
		},
		*w,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating client")
	}

	txResp, err := rarimoClient.Send(
		fromAddr,
		toAddr,
		amount,
		denom,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error sending tx")
	}

	return txResp, nil
}

//func (c *Connector) RemoveCredentials() {}

//func (c *Connector) GetCredentials() {}

//func (c *Connector) CheckStateContractSync() {}

//func (c *Connector) CheckCredentialExistence() {}
