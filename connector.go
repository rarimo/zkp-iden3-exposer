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
	zkpTypes "github.com/rarimo/zkp-iden3-exposer/zkp/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net/http"
	"time"
)

type Connector struct {
	PkHex string `json:"pkHex"`

	IdType        []byte `json:"idType"`
	SchemaHashHex string `json:"schemaHashHex"`

	TargetChainId              int    `json:"targetChainId"`
	TargetRpcUrl               string `json:"targetRpcUrl"`
	TargetStateContractAddress string `json:"targetStateContractAddress"`

	CoreApiUrl               string `json:"coreApiUrl"`
	CoreEvmRpcApiUrl         string `json:"coreEvmRpcApiUrl"`
	CoreStateContractAddress string `json:"coreStateContractAddress"`

	ChainId     string `json:"chainId"`
	AddrPrefix  string `json:"addrPrefix"`
	Denom       string `json:"denom"`
	RpcApi      string `json:"rpcApi"`
	MinGasPrice int    `json:"minGasPrice"`
	GasLimit    int    `json:"gasLimit"`
	IsTLS       bool   `json:"tls"`
}

func NewConnector(
	pkHex string,
	idType []byte,
	schemaHashHex string,
	targetChainId int,
	targetRpcUrl string,
	targetStateContractAddress string,
	coreApiUrl string,
	coreEvmRpcApiUrl string,
	coreStateContractAddress string,
	chainId string,
	addrPrefix string,
	denom string,
	rpcApi string,
	minGasPrice int,
	gasLimit int,
	isTls bool,
) *Connector {
	return &Connector{
		PkHex:                      pkHex,
		IdType:                     idType,
		SchemaHashHex:              schemaHashHex,
		TargetChainId:              targetChainId,
		TargetRpcUrl:               targetRpcUrl,
		TargetStateContractAddress: targetStateContractAddress,
		CoreApiUrl:                 coreApiUrl,
		CoreEvmRpcApiUrl:           coreEvmRpcApiUrl,
		CoreStateContractAddress:   coreStateContractAddress,

		ChainId:     chainId,
		AddrPrefix:  addrPrefix,
		Denom:       denom,
		RpcApi:      rpcApi,
		MinGasPrice: minGasPrice,
		GasLimit:    gasLimit,
		IsTLS:       isTls,
	}
}

func getIdentityInstance(identityConfig zkpTypes.IdentityConfig) (*instances.Identity, error) {
	if identityConfig.PkHex == "" || &identityConfig.PkHex == nil {
		return nil, errors.New("Private key is required")
	}

	identity, err := instances.NewIdentity(instances.IdentityConfig{
		IdType:        identityConfig.IdType,
		SchemaHashHex: identityConfig.SchemaHashHex,
		ChainInfo: zkpTypes.ChainZkpInfo{
			TargetChainId:              identityConfig.TargetChainId,
			TargetRpcUrl:               identityConfig.TargetRpcUrl,
			TargetStateContractAddress: identityConfig.TargetStateContractAddress,

			CoreApiUrl:               identityConfig.CoreApiUrl,
			CoreEvmRpcApiUrl:         identityConfig.CoreEvmRpcApiUrl,
			CoreStateContractAddress: identityConfig.CoreStateContractAddress,
		},
	}, &identityConfig.PkHex)

	if err != nil {
		return nil, errors.Wrap(err, "Error creating identity")
	}

	return identity, nil
}

func (c *Connector) getIdentityConfig() *zkpTypes.IdentityConfig {
	return &zkpTypes.IdentityConfig{
		PkHex:                      c.PkHex,
		IdType:                     [2]byte(c.IdType),
		SchemaHashHex:              c.SchemaHashHex,
		TargetChainId:              c.TargetChainId,
		TargetRpcUrl:               c.TargetRpcUrl,
		TargetStateContractAddress: c.TargetStateContractAddress,
		CoreApiUrl:                 c.CoreApiUrl,
		CoreEvmRpcApiUrl:           c.CoreEvmRpcApiUrl,
		CoreStateContractAddress:   c.CoreStateContractAddress,
	}
}

func (c *Connector) GetOfferJson(issuerApi string, identityDidString string, claimType string) ([]byte, error) {
	offer := zkpTypes.ClaimOffer{}

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

func (c *Connector) GetDidString() (string, error) {
	identity, err := getIdentityInstance(*c.getIdentityConfig())

	if err != nil {
		return "", errors.Wrap(err, "Error getting identity")
	}

	return identity.DID.String(), nil
}

func (c *Connector) GetIdBigIntString() (string, error) {
	identity, err := getIdentityInstance(*c.getIdentityConfig())
	if err != nil {
		return "", errors.Wrap(err, "Error getting identity")
	}

	id, err := identity.ID()
	if err != nil {
		return "", errors.Wrap(err, "Error getting ID")
	}

	return id.BigInt().String(), nil
}

func (c *Connector) GetAuthV2Inputs(offerJson []byte) ([]byte, error) {
	identity, err := getIdentityInstance(*c.getIdentityConfig())
	if err != nil {
		return nil, errors.Wrap(err, "Error getting identity")
	}

	offer := zkpTypes.ClaimOffer{}
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
	if err != nil {
		return nil, errors.Wrap(err, "Error creating JWZ token")
	}

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
	offerJson []byte,
	proofRaw []byte,
) ([]byte, error) {
	identity, err := getIdentityInstance(*c.getIdentityConfig())
	if err != nil {
		return nil, errors.Wrap(err, "Error getting identity")
	}

	offer := zkpTypes.ClaimOffer{}
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
	jsonVC []byte,

	circuitId string,
	challenge string,

	subjectFieldName string,
	subjectFieldValue string,
	operator int,
) ([]byte, error) {
	identity, err := getIdentityInstance(*c.getIdentityConfig())
	if err != nil {
		return nil, errors.Wrap(err, "Error getting identity")
	}

	proofRequest := zkpTypes.CreateProofRequest{
		CircuitId: circuits.CircuitID(circuitId),
		Challenge: challenge,
		Query: zkpTypes.ProofQuery{
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

func (c *Connector) WalletGetAddress() (string, error) {
	w, err := wallet.NewWallet(c.PkHex, c.AddrPrefix)
	if err != nil {
		return "", errors.Wrap(err, "Error creating wallet")
	}

	return w.Address, nil
}

func (c *Connector) WalletSend(fromAddr, toAddr string, amount int64) ([]byte, error) {
	w, err := wallet.NewWallet(c.PkHex, c.AddrPrefix)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating wallet")
	}

	chainConfig := client.ChainConfig{
		ChainId:     c.ChainId,
		Denom:       c.Denom,
		MinGasPrice: uint64(c.MinGasPrice),
		GasLimit:    uint64(c.GasLimit),
	}

	grpcClient, err := grpc.Dial(
		c.RpcApi,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    10 * time.Second, // wait time before ping if no activity
			Timeout: 20 * time.Second, // ping timeout
		}),
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error dialing grpc")
	}

	rarimoClient, err := client.NewClient(
		grpcClient,
		chainConfig,
		*w,
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating client")
	}

	txResp, err := rarimoClient.Send(
		fromAddr,
		toAddr,
		amount,
		c.Denom,
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
