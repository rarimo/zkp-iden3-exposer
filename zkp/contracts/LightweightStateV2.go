// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ILightweightStateV2GistRootData is an auto generated low-level Go binding around an user-defined struct.
type ILightweightStateV2GistRootData struct {
	Root               *big.Int
	CreatedAtTimestamp *big.Int
}

// ILightweightStateV2IdentitiesStatesRootData is an auto generated low-level Go binding around an user-defined struct.
type ILightweightStateV2IdentitiesStatesRootData struct {
	Root         [32]byte
	SetTimestamp *big.Int
}

// LightweightStateV2MetaData contains all meta data concerning the LightweightStateV2 contract.
var LightweightStateV2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGistRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newIdentitesStatesRoot\",\"type\":\"bytes32\"}],\"name\":\"SignedStateTransited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"P\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceStateContract_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sourceChainName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"}],\"name\":\"__LightweightStateV2_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"}],\"name\":\"__Signers_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"newSignerPubKey_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"changeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSourceStateContract_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"changeSourceStateContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"checkSignatureAndIncrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentGISTRootInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structILightweightStateV2.GistRootData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGISTRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"root_\",\"type\":\"uint256\"}],\"name\":\"getGISTRootInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structILightweightStateV2.GistRootData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"}],\"name\":\"getIdentitiesStatesRootData\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"setTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structILightweightStateV2.IdentitiesStatesRootData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"getSigComponents\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"identitiesStatesRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"}],\"name\":\"isIdentitiesStatesRootExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newIdentitiesStatesRoot_\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structILightweightStateV2.GistRootData\",\"name\":\"gistData_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"}],\"name\":\"signedTransitState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceChainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceStateContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"upgradeToWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"validateChangeAddressSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LightweightStateV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use LightweightStateV2MetaData.ABI instead.
var LightweightStateV2ABI = LightweightStateV2MetaData.ABI

// LightweightStateV2 is an auto generated Go binding around an Ethereum contract.
type LightweightStateV2 struct {
	LightweightStateV2Caller     // Read-only binding to the contract
	LightweightStateV2Transactor // Write-only binding to the contract
	LightweightStateV2Filterer   // Log filterer for contract events
}

// LightweightStateV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type LightweightStateV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightweightStateV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type LightweightStateV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightweightStateV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LightweightStateV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightweightStateV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LightweightStateV2Session struct {
	Contract     *LightweightStateV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LightweightStateV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LightweightStateV2CallerSession struct {
	Contract *LightweightStateV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// LightweightStateV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LightweightStateV2TransactorSession struct {
	Contract     *LightweightStateV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// LightweightStateV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type LightweightStateV2Raw struct {
	Contract *LightweightStateV2 // Generic contract binding to access the raw methods on
}

// LightweightStateV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LightweightStateV2CallerRaw struct {
	Contract *LightweightStateV2Caller // Generic read-only contract binding to access the raw methods on
}

// LightweightStateV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LightweightStateV2TransactorRaw struct {
	Contract *LightweightStateV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewLightweightStateV2 creates a new instance of LightweightStateV2, bound to a specific deployed contract.
func NewLightweightStateV2(address common.Address, backend bind.ContractBackend) (*LightweightStateV2, error) {
	contract, err := bindLightweightStateV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LightweightStateV2{LightweightStateV2Caller: LightweightStateV2Caller{contract: contract}, LightweightStateV2Transactor: LightweightStateV2Transactor{contract: contract}, LightweightStateV2Filterer: LightweightStateV2Filterer{contract: contract}}, nil
}

// NewLightweightStateV2Caller creates a new read-only instance of LightweightStateV2, bound to a specific deployed contract.
func NewLightweightStateV2Caller(address common.Address, caller bind.ContractCaller) (*LightweightStateV2Caller, error) {
	contract, err := bindLightweightStateV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LightweightStateV2Caller{contract: contract}, nil
}

// NewLightweightStateV2Transactor creates a new write-only instance of LightweightStateV2, bound to a specific deployed contract.
func NewLightweightStateV2Transactor(address common.Address, transactor bind.ContractTransactor) (*LightweightStateV2Transactor, error) {
	contract, err := bindLightweightStateV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LightweightStateV2Transactor{contract: contract}, nil
}

// NewLightweightStateV2Filterer creates a new log filterer instance of LightweightStateV2, bound to a specific deployed contract.
func NewLightweightStateV2Filterer(address common.Address, filterer bind.ContractFilterer) (*LightweightStateV2Filterer, error) {
	contract, err := bindLightweightStateV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LightweightStateV2Filterer{contract: contract}, nil
}

// bindLightweightStateV2 binds a generic wrapper to an already deployed contract.
func bindLightweightStateV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LightweightStateV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightweightStateV2 *LightweightStateV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightweightStateV2.Contract.LightweightStateV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightweightStateV2 *LightweightStateV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.LightweightStateV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightweightStateV2 *LightweightStateV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.LightweightStateV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightweightStateV2 *LightweightStateV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightweightStateV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightweightStateV2 *LightweightStateV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightweightStateV2 *LightweightStateV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.contract.Transact(opts, method, params...)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_LightweightStateV2 *LightweightStateV2Caller) P(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightweightStateV2.contract.Call(opts, &out, "P")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_LightweightStateV2 *LightweightStateV2Session) P() (*big.Int, error) {
	return _LightweightStateV2.Contract.P(&_LightweightStateV2.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_LightweightStateV2 *LightweightStateV2CallerSession) P() (*big.Int, error) {
	return _LightweightStateV2.Contract.P(&_LightweightStateV2.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_LightweightStateV2 *LightweightStateV2Caller) ChainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LightweightStateV2.contract.Call(opts, &out, "chainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_LightweightStateV2 *LightweightStateV2Session) ChainName() (string, error) {
	return _LightweightStateV2.Contract.ChainName(&_LightweightStateV2.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_LightweightStateV2 *LightweightStateV2CallerSession) ChainName() (string, error) {
	return _LightweightStateV2.Contract.ChainName(&_LightweightStateV2.CallOpts)
}

// GetCurrentGISTRootInfo is a free data retrieval call binding the contract method 0xaf7a3f59.
//
// Solidity: function getCurrentGISTRootInfo() view returns((uint256,uint256))
func (_LightweightStateV2 *LightweightStateV2Caller) GetCurrentGISTRootInfo(opts *bind.CallOpts) (ILightweightStateV2GistRootData, error) {
	var out []interface{}
	err := _LightweightStateV2.contract.Call(opts, &out, "getCurrentGISTRootInfo")

	if err != nil {
		return *new(ILightweightStateV2GistRootData), err
	}

	out0 := *abi.ConvertType(out[0], new(ILightweightStateV2GistRootData)).(*ILightweightStateV2GistRootData)

	return out0, err

}

// GetCurrentGISTRootInfo is a free data retrieval call binding the contract method 0xaf7a3f59.
//
// Solidity: function getCurrentGISTRootInfo() view returns((uint256,uint256))
func (_LightweightStateV2 *LightweightStateV2Session) GetCurrentGISTRootInfo() (ILightweightStateV2GistRootData, error) {
	return _LightweightStateV2.Contract.GetCurrentGISTRootInfo(&_LightweightStateV2.CallOpts)
}

// GetCurrentGISTRootInfo is a free data retrieval call binding the contract method 0xaf7a3f59.
//
// Solidity: function getCurrentGISTRootInfo() view returns((uint256,uint256))
func (_LightweightStateV2 *LightweightStateV2CallerSession) GetCurrentGISTRootInfo() (ILightweightStateV2GistRootData, error) {
	return _LightweightStateV2.Contract.GetCurrentGISTRootInfo(&_LightweightStateV2.CallOpts)
}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_LightweightStateV2 *LightweightStateV2Caller) GetGISTRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightweightStateV2.contract.Call(opts, &out, "getGISTRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_LightweightStateV2 *LightweightStateV2Session) GetGISTRoot() (*big.Int, error) {
	return _LightweightStateV2.Contract.GetGISTRoot(&_LightweightStateV2.CallOpts)
}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_LightweightStateV2 *LightweightStateV2CallerSession) GetGISTRoot() (*big.Int, error) {
	return _LightweightStateV2.Contract.GetGISTRoot(&_LightweightStateV2.CallOpts)
}

// GetGISTRootInfo is a free data retrieval call binding the contract method 0x7c1a66de.
//
// Solidity: function getGISTRootInfo(uint256 root_) view returns((uint256,uint256))
func (_LightweightStateV2 *LightweightStateV2Caller) GetGISTRootInfo(opts *bind.CallOpts, root_ *big.Int) (ILightweightStateV2GistRootData, error) {
	var out []interface{}
	err := _LightweightStateV2.contract.Call(opts, &out, "getGISTRootInfo", root_)

	if err != nil {
		return *new(ILightweightStateV2GistRootData), err
	}

	out0 := *abi.ConvertType(out[0], new(ILightweightStateV2GistRootData)).(*ILightweightStateV2GistRootData)

	return out0, err

}

// GetGISTRootInfo is a free data retrieval call binding the contract method 0x7c1a66de.
//
// Solidity: function getGISTRootInfo(uint256 root_) view returns((uint256,uint256))
func (_LightweightStateV2 *LightweightStateV2Session) GetGISTRootInfo(root_ *big.Int) (ILightweightStateV2GistRootData, error) {
	return _LightweightStateV2.Contract.GetGISTRootInfo(&_LightweightStateV2.CallOpts, root_)
}

// GetGISTRootInfo is a free data retrieval call binding the contract method 0x7c1a66de.
//
// Solidity: function getGISTRootInfo(uint256 root_) view returns((uint256,uint256))
func (_LightweightStateV2 *LightweightStateV2CallerSession) GetGISTRootInfo(root_ *big.Int) (ILightweightStateV2GistRootData, error) {
	return _LightweightStateV2.Contract.GetGISTRootInfo(&_LightweightStateV2.CallOpts, root_)
}

// GetIdentitiesStatesRootData is a free data retrieval call binding the contract method 0xa055a692.
//
// Solidity: function getIdentitiesStatesRootData(bytes32 root_) view returns((bytes32,uint256))
func (_LightweightStateV2 *LightweightStateV2Caller) GetIdentitiesStatesRootData(opts *bind.CallOpts, root_ [32]byte) (ILightweightStateV2IdentitiesStatesRootData, error) {
	var out []interface{}
	err := _LightweightStateV2.contract.Call(opts, &out, "getIdentitiesStatesRootData", root_)

	if err != nil {
		return *new(ILightweightStateV2IdentitiesStatesRootData), err
	}

	out0 := *abi.ConvertType(out[0], new(ILightweightStateV2IdentitiesStatesRootData)).(*ILightweightStateV2IdentitiesStatesRootData)

	return out0, err

}

// GetIdentitiesStatesRootData is a free data retrieval call binding the contract method 0xa055a692.
//
// Solidity: function getIdentitiesStatesRootData(bytes32 root_) view returns((bytes32,uint256))
func (_LightweightStateV2 *LightweightStateV2Session) GetIdentitiesStatesRootData(root_ [32]byte) (ILightweightStateV2IdentitiesStatesRootData, error) {
	return _LightweightStateV2.Contract.GetIdentitiesStatesRootData(&_LightweightStateV2.CallOpts, root_)
}

// GetIdentitiesStatesRootData is a free data retrieval call binding the contract method 0xa055a692.
//
// Solidity: function getIdentitiesStatesRootData(bytes32 root_) view returns((bytes32,uint256))
func (_LightweightStateV2 *LightweightStateV2CallerSession) GetIdentitiesStatesRootData(root_ [32]byte) (ILightweightStateV2IdentitiesStatesRootData, error) {
	return _LightweightStateV2.Contract.GetIdentitiesStatesRootData(&_LightweightStateV2.CallOpts, root_)
}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_LightweightStateV2 *LightweightStateV2Caller) GetSigComponents(opts *bind.CallOpts, methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	var out []interface{}
	err := _LightweightStateV2.contract.Call(opts, &out, "getSigComponents", methodId_, contractAddress_)

	outstruct := new(struct {
		ChainName string
		Nonce     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Nonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_LightweightStateV2 *LightweightStateV2Session) GetSigComponents(methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	return _LightweightStateV2.Contract.GetSigComponents(&_LightweightStateV2.CallOpts, methodId_, contractAddress_)
}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_LightweightStateV2 *LightweightStateV2CallerSession) GetSigComponents(methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	return _LightweightStateV2.Contract.GetSigComponents(&_LightweightStateV2.CallOpts, methodId_, contractAddress_)
}

// IdentitiesStatesRoot is a free data retrieval call binding the contract method 0xe08e70bb.
//
// Solidity: function identitiesStatesRoot() view returns(bytes32)
func (_LightweightStateV2 *LightweightStateV2Caller) IdentitiesStatesRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LightweightStateV2.contract.Call(opts, &out, "identitiesStatesRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IdentitiesStatesRoot is a free data retrieval call binding the contract method 0xe08e70bb.
//
// Solidity: function identitiesStatesRoot() view returns(bytes32)
func (_LightweightStateV2 *LightweightStateV2Session) IdentitiesStatesRoot() ([32]byte, error) {
	return _LightweightStateV2.Contract.IdentitiesStatesRoot(&_LightweightStateV2.CallOpts)
}

// IdentitiesStatesRoot is a free data retrieval call binding the contract method 0xe08e70bb.
//
// Solidity: function identitiesStatesRoot() view returns(bytes32)
func (_LightweightStateV2 *LightweightStateV2CallerSession) IdentitiesStatesRoot() ([32]byte, error) {
	return _LightweightStateV2.Contract.IdentitiesStatesRoot(&_LightweightStateV2.CallOpts)
}

// IsIdentitiesStatesRootExists is a free data retrieval call binding the contract method 0xbfd73455.
//
// Solidity: function isIdentitiesStatesRootExists(bytes32 root_) view returns(bool)
func (_LightweightStateV2 *LightweightStateV2Caller) IsIdentitiesStatesRootExists(opts *bind.CallOpts, root_ [32]byte) (bool, error) {
	var out []interface{}
	err := _LightweightStateV2.contract.Call(opts, &out, "isIdentitiesStatesRootExists", root_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIdentitiesStatesRootExists is a free data retrieval call binding the contract method 0xbfd73455.
//
// Solidity: function isIdentitiesStatesRootExists(bytes32 root_) view returns(bool)
func (_LightweightStateV2 *LightweightStateV2Session) IsIdentitiesStatesRootExists(root_ [32]byte) (bool, error) {
	return _LightweightStateV2.Contract.IsIdentitiesStatesRootExists(&_LightweightStateV2.CallOpts, root_)
}

// IsIdentitiesStatesRootExists is a free data retrieval call binding the contract method 0xbfd73455.
//
// Solidity: function isIdentitiesStatesRootExists(bytes32 root_) view returns(bool)
func (_LightweightStateV2 *LightweightStateV2CallerSession) IsIdentitiesStatesRootExists(root_ [32]byte) (bool, error) {
	return _LightweightStateV2.Contract.IsIdentitiesStatesRootExists(&_LightweightStateV2.CallOpts, root_)
}

// Nonces is a free data retrieval call binding the contract method 0xed3218a2.
//
// Solidity: function nonces(address , uint8 ) view returns(uint256)
func (_LightweightStateV2 *LightweightStateV2Caller) Nonces(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (*big.Int, error) {
	var out []interface{}
	err := _LightweightStateV2.contract.Call(opts, &out, "nonces", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0xed3218a2.
//
// Solidity: function nonces(address , uint8 ) view returns(uint256)
func (_LightweightStateV2 *LightweightStateV2Session) Nonces(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _LightweightStateV2.Contract.Nonces(&_LightweightStateV2.CallOpts, arg0, arg1)
}

// Nonces is a free data retrieval call binding the contract method 0xed3218a2.
//
// Solidity: function nonces(address , uint8 ) view returns(uint256)
func (_LightweightStateV2 *LightweightStateV2CallerSession) Nonces(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _LightweightStateV2.Contract.Nonces(&_LightweightStateV2.CallOpts, arg0, arg1)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LightweightStateV2 *LightweightStateV2Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LightweightStateV2.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LightweightStateV2 *LightweightStateV2Session) ProxiableUUID() ([32]byte, error) {
	return _LightweightStateV2.Contract.ProxiableUUID(&_LightweightStateV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LightweightStateV2 *LightweightStateV2CallerSession) ProxiableUUID() ([32]byte, error) {
	return _LightweightStateV2.Contract.ProxiableUUID(&_LightweightStateV2.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_LightweightStateV2 *LightweightStateV2Caller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightweightStateV2.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_LightweightStateV2 *LightweightStateV2Session) Signer() (common.Address, error) {
	return _LightweightStateV2.Contract.Signer(&_LightweightStateV2.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_LightweightStateV2 *LightweightStateV2CallerSession) Signer() (common.Address, error) {
	return _LightweightStateV2.Contract.Signer(&_LightweightStateV2.CallOpts)
}

// SourceChainName is a free data retrieval call binding the contract method 0xe4ffd04a.
//
// Solidity: function sourceChainName() view returns(string)
func (_LightweightStateV2 *LightweightStateV2Caller) SourceChainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LightweightStateV2.contract.Call(opts, &out, "sourceChainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SourceChainName is a free data retrieval call binding the contract method 0xe4ffd04a.
//
// Solidity: function sourceChainName() view returns(string)
func (_LightweightStateV2 *LightweightStateV2Session) SourceChainName() (string, error) {
	return _LightweightStateV2.Contract.SourceChainName(&_LightweightStateV2.CallOpts)
}

// SourceChainName is a free data retrieval call binding the contract method 0xe4ffd04a.
//
// Solidity: function sourceChainName() view returns(string)
func (_LightweightStateV2 *LightweightStateV2CallerSession) SourceChainName() (string, error) {
	return _LightweightStateV2.Contract.SourceChainName(&_LightweightStateV2.CallOpts)
}

// SourceStateContract is a free data retrieval call binding the contract method 0xfc228319.
//
// Solidity: function sourceStateContract() view returns(address)
func (_LightweightStateV2 *LightweightStateV2Caller) SourceStateContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightweightStateV2.contract.Call(opts, &out, "sourceStateContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SourceStateContract is a free data retrieval call binding the contract method 0xfc228319.
//
// Solidity: function sourceStateContract() view returns(address)
func (_LightweightStateV2 *LightweightStateV2Session) SourceStateContract() (common.Address, error) {
	return _LightweightStateV2.Contract.SourceStateContract(&_LightweightStateV2.CallOpts)
}

// SourceStateContract is a free data retrieval call binding the contract method 0xfc228319.
//
// Solidity: function sourceStateContract() view returns(address)
func (_LightweightStateV2 *LightweightStateV2CallerSession) SourceStateContract() (common.Address, error) {
	return _LightweightStateV2.Contract.SourceStateContract(&_LightweightStateV2.CallOpts)
}

// LightweightStateV2Init is a paid mutator transaction binding the contract method 0x1e0d9571.
//
// Solidity: function __LightweightStateV2_init(address signer_, address sourceStateContract_, string sourceChainName_, string chainName_) returns()
func (_LightweightStateV2 *LightweightStateV2Transactor) LightweightStateV2Init(opts *bind.TransactOpts, signer_ common.Address, sourceStateContract_ common.Address, sourceChainName_ string, chainName_ string) (*types.Transaction, error) {
	return _LightweightStateV2.contract.Transact(opts, "__LightweightStateV2_init", signer_, sourceStateContract_, sourceChainName_, chainName_)
}

// LightweightStateV2Init is a paid mutator transaction binding the contract method 0x1e0d9571.
//
// Solidity: function __LightweightStateV2_init(address signer_, address sourceStateContract_, string sourceChainName_, string chainName_) returns()
func (_LightweightStateV2 *LightweightStateV2Session) LightweightStateV2Init(signer_ common.Address, sourceStateContract_ common.Address, sourceChainName_ string, chainName_ string) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.LightweightStateV2Init(&_LightweightStateV2.TransactOpts, signer_, sourceStateContract_, sourceChainName_, chainName_)
}

// LightweightStateV2Init is a paid mutator transaction binding the contract method 0x1e0d9571.
//
// Solidity: function __LightweightStateV2_init(address signer_, address sourceStateContract_, string sourceChainName_, string chainName_) returns()
func (_LightweightStateV2 *LightweightStateV2TransactorSession) LightweightStateV2Init(signer_ common.Address, sourceStateContract_ common.Address, sourceChainName_ string, chainName_ string) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.LightweightStateV2Init(&_LightweightStateV2.TransactOpts, signer_, sourceStateContract_, sourceChainName_, chainName_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x509ace95.
//
// Solidity: function __Signers_init(address signer_, string chainName_) returns()
func (_LightweightStateV2 *LightweightStateV2Transactor) SignersInit(opts *bind.TransactOpts, signer_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _LightweightStateV2.contract.Transact(opts, "__Signers_init", signer_, chainName_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x509ace95.
//
// Solidity: function __Signers_init(address signer_, string chainName_) returns()
func (_LightweightStateV2 *LightweightStateV2Session) SignersInit(signer_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.SignersInit(&_LightweightStateV2.TransactOpts, signer_, chainName_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x509ace95.
//
// Solidity: function __Signers_init(address signer_, string chainName_) returns()
func (_LightweightStateV2 *LightweightStateV2TransactorSession) SignersInit(signer_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.SignersInit(&_LightweightStateV2.TransactOpts, signer_, chainName_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_LightweightStateV2 *LightweightStateV2Transactor) ChangeSigner(opts *bind.TransactOpts, newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.contract.Transact(opts, "changeSigner", newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_LightweightStateV2 *LightweightStateV2Session) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.ChangeSigner(&_LightweightStateV2.TransactOpts, newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_LightweightStateV2 *LightweightStateV2TransactorSession) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.ChangeSigner(&_LightweightStateV2.TransactOpts, newSignerPubKey_, signature_)
}

// ChangeSourceStateContract is a paid mutator transaction binding the contract method 0x89aeb0f5.
//
// Solidity: function changeSourceStateContract(address newSourceStateContract_, bytes signature_) returns()
func (_LightweightStateV2 *LightweightStateV2Transactor) ChangeSourceStateContract(opts *bind.TransactOpts, newSourceStateContract_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.contract.Transact(opts, "changeSourceStateContract", newSourceStateContract_, signature_)
}

// ChangeSourceStateContract is a paid mutator transaction binding the contract method 0x89aeb0f5.
//
// Solidity: function changeSourceStateContract(address newSourceStateContract_, bytes signature_) returns()
func (_LightweightStateV2 *LightweightStateV2Session) ChangeSourceStateContract(newSourceStateContract_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.ChangeSourceStateContract(&_LightweightStateV2.TransactOpts, newSourceStateContract_, signature_)
}

// ChangeSourceStateContract is a paid mutator transaction binding the contract method 0x89aeb0f5.
//
// Solidity: function changeSourceStateContract(address newSourceStateContract_, bytes signature_) returns()
func (_LightweightStateV2 *LightweightStateV2TransactorSession) ChangeSourceStateContract(newSourceStateContract_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.ChangeSourceStateContract(&_LightweightStateV2.TransactOpts, newSourceStateContract_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_LightweightStateV2 *LightweightStateV2Transactor) CheckSignatureAndIncrementNonce(opts *bind.TransactOpts, methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.contract.Transact(opts, "checkSignatureAndIncrementNonce", methodId_, contractAddress_, signHash_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_LightweightStateV2 *LightweightStateV2Session) CheckSignatureAndIncrementNonce(methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.CheckSignatureAndIncrementNonce(&_LightweightStateV2.TransactOpts, methodId_, contractAddress_, signHash_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_LightweightStateV2 *LightweightStateV2TransactorSession) CheckSignatureAndIncrementNonce(methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.CheckSignatureAndIncrementNonce(&_LightweightStateV2.TransactOpts, methodId_, contractAddress_, signHash_, signature_)
}

// SignedTransitState is a paid mutator transaction binding the contract method 0x189a5073.
//
// Solidity: function signedTransitState(bytes32 newIdentitiesStatesRoot_, (uint256,uint256) gistData_, bytes proof_) returns()
func (_LightweightStateV2 *LightweightStateV2Transactor) SignedTransitState(opts *bind.TransactOpts, newIdentitiesStatesRoot_ [32]byte, gistData_ ILightweightStateV2GistRootData, proof_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.contract.Transact(opts, "signedTransitState", newIdentitiesStatesRoot_, gistData_, proof_)
}

// SignedTransitState is a paid mutator transaction binding the contract method 0x189a5073.
//
// Solidity: function signedTransitState(bytes32 newIdentitiesStatesRoot_, (uint256,uint256) gistData_, bytes proof_) returns()
func (_LightweightStateV2 *LightweightStateV2Session) SignedTransitState(newIdentitiesStatesRoot_ [32]byte, gistData_ ILightweightStateV2GistRootData, proof_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.SignedTransitState(&_LightweightStateV2.TransactOpts, newIdentitiesStatesRoot_, gistData_, proof_)
}

// SignedTransitState is a paid mutator transaction binding the contract method 0x189a5073.
//
// Solidity: function signedTransitState(bytes32 newIdentitiesStatesRoot_, (uint256,uint256) gistData_, bytes proof_) returns()
func (_LightweightStateV2 *LightweightStateV2TransactorSession) SignedTransitState(newIdentitiesStatesRoot_ [32]byte, gistData_ ILightweightStateV2GistRootData, proof_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.SignedTransitState(&_LightweightStateV2.TransactOpts, newIdentitiesStatesRoot_, gistData_, proof_)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LightweightStateV2 *LightweightStateV2Transactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _LightweightStateV2.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LightweightStateV2 *LightweightStateV2Session) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.UpgradeTo(&_LightweightStateV2.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LightweightStateV2 *LightweightStateV2TransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.UpgradeTo(&_LightweightStateV2.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LightweightStateV2 *LightweightStateV2Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LightweightStateV2.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LightweightStateV2 *LightweightStateV2Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.UpgradeToAndCall(&_LightweightStateV2.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LightweightStateV2 *LightweightStateV2TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.UpgradeToAndCall(&_LightweightStateV2.TransactOpts, newImplementation, data)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_LightweightStateV2 *LightweightStateV2Transactor) UpgradeToWithSig(opts *bind.TransactOpts, newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.contract.Transact(opts, "upgradeToWithSig", newImplementation_, signature_)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_LightweightStateV2 *LightweightStateV2Session) UpgradeToWithSig(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.UpgradeToWithSig(&_LightweightStateV2.TransactOpts, newImplementation_, signature_)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_LightweightStateV2 *LightweightStateV2TransactorSession) UpgradeToWithSig(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.UpgradeToWithSig(&_LightweightStateV2.TransactOpts, newImplementation_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_LightweightStateV2 *LightweightStateV2Transactor) ValidateChangeAddressSignature(opts *bind.TransactOpts, methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.contract.Transact(opts, "validateChangeAddressSignature", methodId_, contractAddress_, newAddress_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_LightweightStateV2 *LightweightStateV2Session) ValidateChangeAddressSignature(methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.ValidateChangeAddressSignature(&_LightweightStateV2.TransactOpts, methodId_, contractAddress_, newAddress_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_LightweightStateV2 *LightweightStateV2TransactorSession) ValidateChangeAddressSignature(methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightStateV2.Contract.ValidateChangeAddressSignature(&_LightweightStateV2.TransactOpts, methodId_, contractAddress_, newAddress_, signature_)
}

// LightweightStateV2AdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the LightweightStateV2 contract.
type LightweightStateV2AdminChangedIterator struct {
	Event *LightweightStateV2AdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightweightStateV2AdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightweightStateV2AdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightweightStateV2AdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightweightStateV2AdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightweightStateV2AdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightweightStateV2AdminChanged represents a AdminChanged event raised by the LightweightStateV2 contract.
type LightweightStateV2AdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LightweightStateV2 *LightweightStateV2Filterer) FilterAdminChanged(opts *bind.FilterOpts) (*LightweightStateV2AdminChangedIterator, error) {

	logs, sub, err := _LightweightStateV2.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &LightweightStateV2AdminChangedIterator{contract: _LightweightStateV2.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LightweightStateV2 *LightweightStateV2Filterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LightweightStateV2AdminChanged) (event.Subscription, error) {

	logs, sub, err := _LightweightStateV2.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightweightStateV2AdminChanged)
				if err := _LightweightStateV2.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LightweightStateV2 *LightweightStateV2Filterer) ParseAdminChanged(log types.Log) (*LightweightStateV2AdminChanged, error) {
	event := new(LightweightStateV2AdminChanged)
	if err := _LightweightStateV2.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightweightStateV2BeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the LightweightStateV2 contract.
type LightweightStateV2BeaconUpgradedIterator struct {
	Event *LightweightStateV2BeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightweightStateV2BeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightweightStateV2BeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightweightStateV2BeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightweightStateV2BeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightweightStateV2BeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightweightStateV2BeaconUpgraded represents a BeaconUpgraded event raised by the LightweightStateV2 contract.
type LightweightStateV2BeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LightweightStateV2 *LightweightStateV2Filterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*LightweightStateV2BeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LightweightStateV2.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &LightweightStateV2BeaconUpgradedIterator{contract: _LightweightStateV2.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LightweightStateV2 *LightweightStateV2Filterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *LightweightStateV2BeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LightweightStateV2.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightweightStateV2BeaconUpgraded)
				if err := _LightweightStateV2.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LightweightStateV2 *LightweightStateV2Filterer) ParseBeaconUpgraded(log types.Log) (*LightweightStateV2BeaconUpgraded, error) {
	event := new(LightweightStateV2BeaconUpgraded)
	if err := _LightweightStateV2.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightweightStateV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LightweightStateV2 contract.
type LightweightStateV2InitializedIterator struct {
	Event *LightweightStateV2Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightweightStateV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightweightStateV2Initialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightweightStateV2Initialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightweightStateV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightweightStateV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightweightStateV2Initialized represents a Initialized event raised by the LightweightStateV2 contract.
type LightweightStateV2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightweightStateV2 *LightweightStateV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*LightweightStateV2InitializedIterator, error) {

	logs, sub, err := _LightweightStateV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LightweightStateV2InitializedIterator{contract: _LightweightStateV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightweightStateV2 *LightweightStateV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LightweightStateV2Initialized) (event.Subscription, error) {

	logs, sub, err := _LightweightStateV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightweightStateV2Initialized)
				if err := _LightweightStateV2.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightweightStateV2 *LightweightStateV2Filterer) ParseInitialized(log types.Log) (*LightweightStateV2Initialized, error) {
	event := new(LightweightStateV2Initialized)
	if err := _LightweightStateV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightweightStateV2SignedStateTransitedIterator is returned from FilterSignedStateTransited and is used to iterate over the raw logs and unpacked data for SignedStateTransited events raised by the LightweightStateV2 contract.
type LightweightStateV2SignedStateTransitedIterator struct {
	Event *LightweightStateV2SignedStateTransited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightweightStateV2SignedStateTransitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightweightStateV2SignedStateTransited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightweightStateV2SignedStateTransited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightweightStateV2SignedStateTransitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightweightStateV2SignedStateTransitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightweightStateV2SignedStateTransited represents a SignedStateTransited event raised by the LightweightStateV2 contract.
type LightweightStateV2SignedStateTransited struct {
	NewGistRoot            *big.Int
	NewIdentitesStatesRoot [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSignedStateTransited is a free log retrieval operation binding the contract event 0x8e8ff16553fbf4a457c08a5e07cc27c8aac14b9a1a8e1f546a6c1b9366304a56.
//
// Solidity: event SignedStateTransited(uint256 newGistRoot, bytes32 newIdentitesStatesRoot)
func (_LightweightStateV2 *LightweightStateV2Filterer) FilterSignedStateTransited(opts *bind.FilterOpts) (*LightweightStateV2SignedStateTransitedIterator, error) {

	logs, sub, err := _LightweightStateV2.contract.FilterLogs(opts, "SignedStateTransited")
	if err != nil {
		return nil, err
	}
	return &LightweightStateV2SignedStateTransitedIterator{contract: _LightweightStateV2.contract, event: "SignedStateTransited", logs: logs, sub: sub}, nil
}

// WatchSignedStateTransited is a free log subscription operation binding the contract event 0x8e8ff16553fbf4a457c08a5e07cc27c8aac14b9a1a8e1f546a6c1b9366304a56.
//
// Solidity: event SignedStateTransited(uint256 newGistRoot, bytes32 newIdentitesStatesRoot)
func (_LightweightStateV2 *LightweightStateV2Filterer) WatchSignedStateTransited(opts *bind.WatchOpts, sink chan<- *LightweightStateV2SignedStateTransited) (event.Subscription, error) {

	logs, sub, err := _LightweightStateV2.contract.WatchLogs(opts, "SignedStateTransited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightweightStateV2SignedStateTransited)
				if err := _LightweightStateV2.contract.UnpackLog(event, "SignedStateTransited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignedStateTransited is a log parse operation binding the contract event 0x8e8ff16553fbf4a457c08a5e07cc27c8aac14b9a1a8e1f546a6c1b9366304a56.
//
// Solidity: event SignedStateTransited(uint256 newGistRoot, bytes32 newIdentitesStatesRoot)
func (_LightweightStateV2 *LightweightStateV2Filterer) ParseSignedStateTransited(log types.Log) (*LightweightStateV2SignedStateTransited, error) {
	event := new(LightweightStateV2SignedStateTransited)
	if err := _LightweightStateV2.contract.UnpackLog(event, "SignedStateTransited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightweightStateV2UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LightweightStateV2 contract.
type LightweightStateV2UpgradedIterator struct {
	Event *LightweightStateV2Upgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightweightStateV2UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightweightStateV2Upgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightweightStateV2Upgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightweightStateV2UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightweightStateV2UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightweightStateV2Upgraded represents a Upgraded event raised by the LightweightStateV2 contract.
type LightweightStateV2Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LightweightStateV2 *LightweightStateV2Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LightweightStateV2UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LightweightStateV2.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LightweightStateV2UpgradedIterator{contract: _LightweightStateV2.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LightweightStateV2 *LightweightStateV2Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LightweightStateV2Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LightweightStateV2.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightweightStateV2Upgraded)
				if err := _LightweightStateV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LightweightStateV2 *LightweightStateV2Filterer) ParseUpgraded(log types.Log) (*LightweightStateV2Upgraded, error) {
	event := new(LightweightStateV2Upgraded)
	if err := _LightweightStateV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
