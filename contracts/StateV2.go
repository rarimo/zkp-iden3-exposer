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

// IStateGistProof is an auto generated low-level Go binding around an user-defined struct.
type IStateGistProof struct {
	Root         *big.Int
	Existence    bool
	Siblings     [64]*big.Int
	Index        *big.Int
	Value        *big.Int
	AuxExistence bool
	AuxIndex     *big.Int
	AuxValue     *big.Int
}

// IStateGistRootInfo is an auto generated low-level Go binding around an user-defined struct.
type IStateGistRootInfo struct {
	Root                *big.Int
	ReplacedByRoot      *big.Int
	CreatedAtTimestamp  *big.Int
	ReplacedAtTimestamp *big.Int
	CreatedAtBlock      *big.Int
	ReplacedAtBlock     *big.Int
}

// IStateStateInfo is an auto generated low-level Go binding around an user-defined struct.
type IStateStateInfo struct {
	Id                  *big.Int
	State               *big.Int
	ReplacedByState     *big.Int
	CreatedAtTimestamp  *big.Int
	ReplacedAtTimestamp *big.Int
	CreatedAtBlock      *big.Int
	ReplacedAtBlock     *big.Int
}

// StateV2MetaData contains all meta data concerning the StateV2 contract.
var StateV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getGISTProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getGISTProofByBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getGISTProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getGISTProofByTime\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGISTRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getGISTRootHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGISTRootHistoryLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getGISTRootInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getGISTRootInfoByBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getGISTRootInfoByTime\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getStateInfoById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.StateInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"getStateInfoByIdAndState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.StateInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getStateInfoHistoryById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.StateInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getStateInfoHistoryLengthById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"idExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStateTransitionVerifier\",\"name\":\"verifierContractAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVerifierAddr\",\"type\":\"address\"}],\"name\":\"setVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"stateExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newState\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOldStateGenesis\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"name\":\"transitState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StateV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use StateV2MetaData.ABI instead.
var StateV2ABI = StateV2MetaData.ABI

// StateV2 is an auto generated Go binding around an Ethereum contract.
type StateV2 struct {
	StateV2Caller     // Read-only binding to the contract
	StateV2Transactor // Write-only binding to the contract
	StateV2Filterer   // Log filterer for contract events
}

// StateV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type StateV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type StateV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateV2Session struct {
	Contract     *StateV2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateV2CallerSession struct {
	Contract *StateV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StateV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateV2TransactorSession struct {
	Contract     *StateV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StateV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type StateV2Raw struct {
	Contract *StateV2 // Generic contract binding to access the raw methods on
}

// StateV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateV2CallerRaw struct {
	Contract *StateV2Caller // Generic read-only contract binding to access the raw methods on
}

// StateV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateV2TransactorRaw struct {
	Contract *StateV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewStateV2 creates a new instance of StateV2, bound to a specific deployed contract.
func NewStateV2(address common.Address, backend bind.ContractBackend) (*StateV2, error) {
	contract, err := bindStateV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateV2{StateV2Caller: StateV2Caller{contract: contract}, StateV2Transactor: StateV2Transactor{contract: contract}, StateV2Filterer: StateV2Filterer{contract: contract}}, nil
}

// NewStateV2Caller creates a new read-only instance of StateV2, bound to a specific deployed contract.
func NewStateV2Caller(address common.Address, caller bind.ContractCaller) (*StateV2Caller, error) {
	contract, err := bindStateV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateV2Caller{contract: contract}, nil
}

// NewStateV2Transactor creates a new write-only instance of StateV2, bound to a specific deployed contract.
func NewStateV2Transactor(address common.Address, transactor bind.ContractTransactor) (*StateV2Transactor, error) {
	contract, err := bindStateV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateV2Transactor{contract: contract}, nil
}

// NewStateV2Filterer creates a new log filterer instance of StateV2, bound to a specific deployed contract.
func NewStateV2Filterer(address common.Address, filterer bind.ContractFilterer) (*StateV2Filterer, error) {
	contract, err := bindStateV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateV2Filterer{contract: contract}, nil
}

// bindStateV2 binds a generic wrapper to an already deployed contract.
func bindStateV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StateV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateV2 *StateV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateV2.Contract.StateV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateV2 *StateV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateV2.Contract.StateV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateV2 *StateV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateV2.Contract.StateV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateV2 *StateV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateV2 *StateV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateV2 *StateV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateV2.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_StateV2 *StateV2Caller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_StateV2 *StateV2Session) VERSION() (string, error) {
	return _StateV2.Contract.VERSION(&_StateV2.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_StateV2 *StateV2CallerSession) VERSION() (string, error) {
	return _StateV2.Contract.VERSION(&_StateV2.CallOpts)
}

// GetGISTProof is a free data retrieval call binding the contract method 0x3025bb8c.
//
// Solidity: function getGISTProof(uint256 id) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2 *StateV2Caller) GetGISTProof(opts *bind.CallOpts, id *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "getGISTProof", id)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProof is a free data retrieval call binding the contract method 0x3025bb8c.
//
// Solidity: function getGISTProof(uint256 id) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2 *StateV2Session) GetGISTProof(id *big.Int) (IStateGistProof, error) {
	return _StateV2.Contract.GetGISTProof(&_StateV2.CallOpts, id)
}

// GetGISTProof is a free data retrieval call binding the contract method 0x3025bb8c.
//
// Solidity: function getGISTProof(uint256 id) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2 *StateV2CallerSession) GetGISTProof(id *big.Int) (IStateGistProof, error) {
	return _StateV2.Contract.GetGISTProof(&_StateV2.CallOpts, id)
}

// GetGISTProofByBlock is a free data retrieval call binding the contract method 0x046ff140.
//
// Solidity: function getGISTProofByBlock(uint256 id, uint256 blockNumber) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2 *StateV2Caller) GetGISTProofByBlock(opts *bind.CallOpts, id *big.Int, blockNumber *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "getGISTProofByBlock", id, blockNumber)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProofByBlock is a free data retrieval call binding the contract method 0x046ff140.
//
// Solidity: function getGISTProofByBlock(uint256 id, uint256 blockNumber) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2 *StateV2Session) GetGISTProofByBlock(id *big.Int, blockNumber *big.Int) (IStateGistProof, error) {
	return _StateV2.Contract.GetGISTProofByBlock(&_StateV2.CallOpts, id, blockNumber)
}

// GetGISTProofByBlock is a free data retrieval call binding the contract method 0x046ff140.
//
// Solidity: function getGISTProofByBlock(uint256 id, uint256 blockNumber) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2 *StateV2CallerSession) GetGISTProofByBlock(id *big.Int, blockNumber *big.Int) (IStateGistProof, error) {
	return _StateV2.Contract.GetGISTProofByBlock(&_StateV2.CallOpts, id, blockNumber)
}

// GetGISTProofByRoot is a free data retrieval call binding the contract method 0xe12a36c0.
//
// Solidity: function getGISTProofByRoot(uint256 id, uint256 root) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2 *StateV2Caller) GetGISTProofByRoot(opts *bind.CallOpts, id *big.Int, root *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "getGISTProofByRoot", id, root)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProofByRoot is a free data retrieval call binding the contract method 0xe12a36c0.
//
// Solidity: function getGISTProofByRoot(uint256 id, uint256 root) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2 *StateV2Session) GetGISTProofByRoot(id *big.Int, root *big.Int) (IStateGistProof, error) {
	return _StateV2.Contract.GetGISTProofByRoot(&_StateV2.CallOpts, id, root)
}

// GetGISTProofByRoot is a free data retrieval call binding the contract method 0xe12a36c0.
//
// Solidity: function getGISTProofByRoot(uint256 id, uint256 root) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2 *StateV2CallerSession) GetGISTProofByRoot(id *big.Int, root *big.Int) (IStateGistProof, error) {
	return _StateV2.Contract.GetGISTProofByRoot(&_StateV2.CallOpts, id, root)
}

// GetGISTProofByTime is a free data retrieval call binding the contract method 0xd51afebf.
//
// Solidity: function getGISTProofByTime(uint256 id, uint256 timestamp) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2 *StateV2Caller) GetGISTProofByTime(opts *bind.CallOpts, id *big.Int, timestamp *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "getGISTProofByTime", id, timestamp)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProofByTime is a free data retrieval call binding the contract method 0xd51afebf.
//
// Solidity: function getGISTProofByTime(uint256 id, uint256 timestamp) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2 *StateV2Session) GetGISTProofByTime(id *big.Int, timestamp *big.Int) (IStateGistProof, error) {
	return _StateV2.Contract.GetGISTProofByTime(&_StateV2.CallOpts, id, timestamp)
}

// GetGISTProofByTime is a free data retrieval call binding the contract method 0xd51afebf.
//
// Solidity: function getGISTProofByTime(uint256 id, uint256 timestamp) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2 *StateV2CallerSession) GetGISTProofByTime(id *big.Int, timestamp *big.Int) (IStateGistProof, error) {
	return _StateV2.Contract.GetGISTProofByTime(&_StateV2.CallOpts, id, timestamp)
}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_StateV2 *StateV2Caller) GetGISTRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "getGISTRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_StateV2 *StateV2Session) GetGISTRoot() (*big.Int, error) {
	return _StateV2.Contract.GetGISTRoot(&_StateV2.CallOpts)
}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_StateV2 *StateV2CallerSession) GetGISTRoot() (*big.Int, error) {
	return _StateV2.Contract.GetGISTRoot(&_StateV2.CallOpts)
}

// GetGISTRootHistory is a free data retrieval call binding the contract method 0x2f7670e4.
//
// Solidity: function getGISTRootHistory(uint256 start, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateV2 *StateV2Caller) GetGISTRootHistory(opts *bind.CallOpts, start *big.Int, length *big.Int) ([]IStateGistRootInfo, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "getGISTRootHistory", start, length)

	if err != nil {
		return *new([]IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStateGistRootInfo)).(*[]IStateGistRootInfo)

	return out0, err

}

// GetGISTRootHistory is a free data retrieval call binding the contract method 0x2f7670e4.
//
// Solidity: function getGISTRootHistory(uint256 start, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateV2 *StateV2Session) GetGISTRootHistory(start *big.Int, length *big.Int) ([]IStateGistRootInfo, error) {
	return _StateV2.Contract.GetGISTRootHistory(&_StateV2.CallOpts, start, length)
}

// GetGISTRootHistory is a free data retrieval call binding the contract method 0x2f7670e4.
//
// Solidity: function getGISTRootHistory(uint256 start, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateV2 *StateV2CallerSession) GetGISTRootHistory(start *big.Int, length *big.Int) ([]IStateGistRootInfo, error) {
	return _StateV2.Contract.GetGISTRootHistory(&_StateV2.CallOpts, start, length)
}

// GetGISTRootHistoryLength is a free data retrieval call binding the contract method 0xdccbd57a.
//
// Solidity: function getGISTRootHistoryLength() view returns(uint256)
func (_StateV2 *StateV2Caller) GetGISTRootHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "getGISTRootHistoryLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGISTRootHistoryLength is a free data retrieval call binding the contract method 0xdccbd57a.
//
// Solidity: function getGISTRootHistoryLength() view returns(uint256)
func (_StateV2 *StateV2Session) GetGISTRootHistoryLength() (*big.Int, error) {
	return _StateV2.Contract.GetGISTRootHistoryLength(&_StateV2.CallOpts)
}

// GetGISTRootHistoryLength is a free data retrieval call binding the contract method 0xdccbd57a.
//
// Solidity: function getGISTRootHistoryLength() view returns(uint256)
func (_StateV2 *StateV2CallerSession) GetGISTRootHistoryLength() (*big.Int, error) {
	return _StateV2.Contract.GetGISTRootHistoryLength(&_StateV2.CallOpts)
}

// GetGISTRootInfo is a free data retrieval call binding the contract method 0x7c1a66de.
//
// Solidity: function getGISTRootInfo(uint256 root) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2 *StateV2Caller) GetGISTRootInfo(opts *bind.CallOpts, root *big.Int) (IStateGistRootInfo, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "getGISTRootInfo", root)

	if err != nil {
		return *new(IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistRootInfo)).(*IStateGistRootInfo)

	return out0, err

}

// GetGISTRootInfo is a free data retrieval call binding the contract method 0x7c1a66de.
//
// Solidity: function getGISTRootInfo(uint256 root) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2 *StateV2Session) GetGISTRootInfo(root *big.Int) (IStateGistRootInfo, error) {
	return _StateV2.Contract.GetGISTRootInfo(&_StateV2.CallOpts, root)
}

// GetGISTRootInfo is a free data retrieval call binding the contract method 0x7c1a66de.
//
// Solidity: function getGISTRootInfo(uint256 root) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2 *StateV2CallerSession) GetGISTRootInfo(root *big.Int) (IStateGistRootInfo, error) {
	return _StateV2.Contract.GetGISTRootInfo(&_StateV2.CallOpts, root)
}

// GetGISTRootInfoByBlock is a free data retrieval call binding the contract method 0x5845e530.
//
// Solidity: function getGISTRootInfoByBlock(uint256 blockNumber) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2 *StateV2Caller) GetGISTRootInfoByBlock(opts *bind.CallOpts, blockNumber *big.Int) (IStateGistRootInfo, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "getGISTRootInfoByBlock", blockNumber)

	if err != nil {
		return *new(IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistRootInfo)).(*IStateGistRootInfo)

	return out0, err

}

// GetGISTRootInfoByBlock is a free data retrieval call binding the contract method 0x5845e530.
//
// Solidity: function getGISTRootInfoByBlock(uint256 blockNumber) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2 *StateV2Session) GetGISTRootInfoByBlock(blockNumber *big.Int) (IStateGistRootInfo, error) {
	return _StateV2.Contract.GetGISTRootInfoByBlock(&_StateV2.CallOpts, blockNumber)
}

// GetGISTRootInfoByBlock is a free data retrieval call binding the contract method 0x5845e530.
//
// Solidity: function getGISTRootInfoByBlock(uint256 blockNumber) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2 *StateV2CallerSession) GetGISTRootInfoByBlock(blockNumber *big.Int) (IStateGistRootInfo, error) {
	return _StateV2.Contract.GetGISTRootInfoByBlock(&_StateV2.CallOpts, blockNumber)
}

// GetGISTRootInfoByTime is a free data retrieval call binding the contract method 0x0ef6e65b.
//
// Solidity: function getGISTRootInfoByTime(uint256 timestamp) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2 *StateV2Caller) GetGISTRootInfoByTime(opts *bind.CallOpts, timestamp *big.Int) (IStateGistRootInfo, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "getGISTRootInfoByTime", timestamp)

	if err != nil {
		return *new(IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistRootInfo)).(*IStateGistRootInfo)

	return out0, err

}

// GetGISTRootInfoByTime is a free data retrieval call binding the contract method 0x0ef6e65b.
//
// Solidity: function getGISTRootInfoByTime(uint256 timestamp) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2 *StateV2Session) GetGISTRootInfoByTime(timestamp *big.Int) (IStateGistRootInfo, error) {
	return _StateV2.Contract.GetGISTRootInfoByTime(&_StateV2.CallOpts, timestamp)
}

// GetGISTRootInfoByTime is a free data retrieval call binding the contract method 0x0ef6e65b.
//
// Solidity: function getGISTRootInfoByTime(uint256 timestamp) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2 *StateV2CallerSession) GetGISTRootInfoByTime(timestamp *big.Int) (IStateGistRootInfo, error) {
	return _StateV2.Contract.GetGISTRootInfoByTime(&_StateV2.CallOpts, timestamp)
}

// GetStateInfoById is a free data retrieval call binding the contract method 0xb4bdea55.
//
// Solidity: function getStateInfoById(uint256 id) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2 *StateV2Caller) GetStateInfoById(opts *bind.CallOpts, id *big.Int) (IStateStateInfo, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "getStateInfoById", id)

	if err != nil {
		return *new(IStateStateInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateStateInfo)).(*IStateStateInfo)

	return out0, err

}

// GetStateInfoById is a free data retrieval call binding the contract method 0xb4bdea55.
//
// Solidity: function getStateInfoById(uint256 id) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2 *StateV2Session) GetStateInfoById(id *big.Int) (IStateStateInfo, error) {
	return _StateV2.Contract.GetStateInfoById(&_StateV2.CallOpts, id)
}

// GetStateInfoById is a free data retrieval call binding the contract method 0xb4bdea55.
//
// Solidity: function getStateInfoById(uint256 id) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2 *StateV2CallerSession) GetStateInfoById(id *big.Int) (IStateStateInfo, error) {
	return _StateV2.Contract.GetStateInfoById(&_StateV2.CallOpts, id)
}

// GetStateInfoByIdAndState is a free data retrieval call binding the contract method 0x53c87312.
//
// Solidity: function getStateInfoByIdAndState(uint256 id, uint256 state) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2 *StateV2Caller) GetStateInfoByIdAndState(opts *bind.CallOpts, id *big.Int, state *big.Int) (IStateStateInfo, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "getStateInfoByIdAndState", id, state)

	if err != nil {
		return *new(IStateStateInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateStateInfo)).(*IStateStateInfo)

	return out0, err

}

// GetStateInfoByIdAndState is a free data retrieval call binding the contract method 0x53c87312.
//
// Solidity: function getStateInfoByIdAndState(uint256 id, uint256 state) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2 *StateV2Session) GetStateInfoByIdAndState(id *big.Int, state *big.Int) (IStateStateInfo, error) {
	return _StateV2.Contract.GetStateInfoByIdAndState(&_StateV2.CallOpts, id, state)
}

// GetStateInfoByIdAndState is a free data retrieval call binding the contract method 0x53c87312.
//
// Solidity: function getStateInfoByIdAndState(uint256 id, uint256 state) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2 *StateV2CallerSession) GetStateInfoByIdAndState(id *big.Int, state *big.Int) (IStateStateInfo, error) {
	return _StateV2.Contract.GetStateInfoByIdAndState(&_StateV2.CallOpts, id, state)
}

// GetStateInfoHistoryById is a free data retrieval call binding the contract method 0xe99858fe.
//
// Solidity: function getStateInfoHistoryById(uint256 id, uint256 startIndex, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateV2 *StateV2Caller) GetStateInfoHistoryById(opts *bind.CallOpts, id *big.Int, startIndex *big.Int, length *big.Int) ([]IStateStateInfo, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "getStateInfoHistoryById", id, startIndex, length)

	if err != nil {
		return *new([]IStateStateInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStateStateInfo)).(*[]IStateStateInfo)

	return out0, err

}

// GetStateInfoHistoryById is a free data retrieval call binding the contract method 0xe99858fe.
//
// Solidity: function getStateInfoHistoryById(uint256 id, uint256 startIndex, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateV2 *StateV2Session) GetStateInfoHistoryById(id *big.Int, startIndex *big.Int, length *big.Int) ([]IStateStateInfo, error) {
	return _StateV2.Contract.GetStateInfoHistoryById(&_StateV2.CallOpts, id, startIndex, length)
}

// GetStateInfoHistoryById is a free data retrieval call binding the contract method 0xe99858fe.
//
// Solidity: function getStateInfoHistoryById(uint256 id, uint256 startIndex, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateV2 *StateV2CallerSession) GetStateInfoHistoryById(id *big.Int, startIndex *big.Int, length *big.Int) ([]IStateStateInfo, error) {
	return _StateV2.Contract.GetStateInfoHistoryById(&_StateV2.CallOpts, id, startIndex, length)
}

// GetStateInfoHistoryLengthById is a free data retrieval call binding the contract method 0x676d5b5a.
//
// Solidity: function getStateInfoHistoryLengthById(uint256 id) view returns(uint256)
func (_StateV2 *StateV2Caller) GetStateInfoHistoryLengthById(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "getStateInfoHistoryLengthById", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStateInfoHistoryLengthById is a free data retrieval call binding the contract method 0x676d5b5a.
//
// Solidity: function getStateInfoHistoryLengthById(uint256 id) view returns(uint256)
func (_StateV2 *StateV2Session) GetStateInfoHistoryLengthById(id *big.Int) (*big.Int, error) {
	return _StateV2.Contract.GetStateInfoHistoryLengthById(&_StateV2.CallOpts, id)
}

// GetStateInfoHistoryLengthById is a free data retrieval call binding the contract method 0x676d5b5a.
//
// Solidity: function getStateInfoHistoryLengthById(uint256 id) view returns(uint256)
func (_StateV2 *StateV2CallerSession) GetStateInfoHistoryLengthById(id *big.Int) (*big.Int, error) {
	return _StateV2.Contract.GetStateInfoHistoryLengthById(&_StateV2.CallOpts, id)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_StateV2 *StateV2Caller) GetVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "getVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_StateV2 *StateV2Session) GetVerifier() (common.Address, error) {
	return _StateV2.Contract.GetVerifier(&_StateV2.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_StateV2 *StateV2CallerSession) GetVerifier() (common.Address, error) {
	return _StateV2.Contract.GetVerifier(&_StateV2.CallOpts)
}

// IdExists is a free data retrieval call binding the contract method 0x0b8a295a.
//
// Solidity: function idExists(uint256 id) view returns(bool)
func (_StateV2 *StateV2Caller) IdExists(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "idExists", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IdExists is a free data retrieval call binding the contract method 0x0b8a295a.
//
// Solidity: function idExists(uint256 id) view returns(bool)
func (_StateV2 *StateV2Session) IdExists(id *big.Int) (bool, error) {
	return _StateV2.Contract.IdExists(&_StateV2.CallOpts, id)
}

// IdExists is a free data retrieval call binding the contract method 0x0b8a295a.
//
// Solidity: function idExists(uint256 id) view returns(bool)
func (_StateV2 *StateV2CallerSession) IdExists(id *big.Int) (bool, error) {
	return _StateV2.Contract.IdExists(&_StateV2.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StateV2 *StateV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StateV2 *StateV2Session) Owner() (common.Address, error) {
	return _StateV2.Contract.Owner(&_StateV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StateV2 *StateV2CallerSession) Owner() (common.Address, error) {
	return _StateV2.Contract.Owner(&_StateV2.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_StateV2 *StateV2Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_StateV2 *StateV2Session) PendingOwner() (common.Address, error) {
	return _StateV2.Contract.PendingOwner(&_StateV2.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_StateV2 *StateV2CallerSession) PendingOwner() (common.Address, error) {
	return _StateV2.Contract.PendingOwner(&_StateV2.CallOpts)
}

// StateExists is a free data retrieval call binding the contract method 0x233a4d23.
//
// Solidity: function stateExists(uint256 id, uint256 state) view returns(bool)
func (_StateV2 *StateV2Caller) StateExists(opts *bind.CallOpts, id *big.Int, state *big.Int) (bool, error) {
	var out []interface{}
	err := _StateV2.contract.Call(opts, &out, "stateExists", id, state)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StateExists is a free data retrieval call binding the contract method 0x233a4d23.
//
// Solidity: function stateExists(uint256 id, uint256 state) view returns(bool)
func (_StateV2 *StateV2Session) StateExists(id *big.Int, state *big.Int) (bool, error) {
	return _StateV2.Contract.StateExists(&_StateV2.CallOpts, id, state)
}

// StateExists is a free data retrieval call binding the contract method 0x233a4d23.
//
// Solidity: function stateExists(uint256 id, uint256 state) view returns(bool)
func (_StateV2 *StateV2CallerSession) StateExists(id *big.Int, state *big.Int) (bool, error) {
	return _StateV2.Contract.StateExists(&_StateV2.CallOpts, id, state)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_StateV2 *StateV2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateV2.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_StateV2 *StateV2Session) AcceptOwnership() (*types.Transaction, error) {
	return _StateV2.Contract.AcceptOwnership(&_StateV2.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_StateV2 *StateV2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _StateV2.Contract.AcceptOwnership(&_StateV2.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address verifierContractAddr) returns()
func (_StateV2 *StateV2Transactor) Initialize(opts *bind.TransactOpts, verifierContractAddr common.Address) (*types.Transaction, error) {
	return _StateV2.contract.Transact(opts, "initialize", verifierContractAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address verifierContractAddr) returns()
func (_StateV2 *StateV2Session) Initialize(verifierContractAddr common.Address) (*types.Transaction, error) {
	return _StateV2.Contract.Initialize(&_StateV2.TransactOpts, verifierContractAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address verifierContractAddr) returns()
func (_StateV2 *StateV2TransactorSession) Initialize(verifierContractAddr common.Address) (*types.Transaction, error) {
	return _StateV2.Contract.Initialize(&_StateV2.TransactOpts, verifierContractAddr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StateV2 *StateV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StateV2 *StateV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _StateV2.Contract.RenounceOwnership(&_StateV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StateV2 *StateV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StateV2.Contract.RenounceOwnership(&_StateV2.TransactOpts)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address newVerifierAddr) returns()
func (_StateV2 *StateV2Transactor) SetVerifier(opts *bind.TransactOpts, newVerifierAddr common.Address) (*types.Transaction, error) {
	return _StateV2.contract.Transact(opts, "setVerifier", newVerifierAddr)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address newVerifierAddr) returns()
func (_StateV2 *StateV2Session) SetVerifier(newVerifierAddr common.Address) (*types.Transaction, error) {
	return _StateV2.Contract.SetVerifier(&_StateV2.TransactOpts, newVerifierAddr)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address newVerifierAddr) returns()
func (_StateV2 *StateV2TransactorSession) SetVerifier(newVerifierAddr common.Address) (*types.Transaction, error) {
	return _StateV2.Contract.SetVerifier(&_StateV2.TransactOpts, newVerifierAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StateV2 *StateV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StateV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StateV2 *StateV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StateV2.Contract.TransferOwnership(&_StateV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StateV2 *StateV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StateV2.Contract.TransferOwnership(&_StateV2.TransactOpts, newOwner)
}

// TransitState is a paid mutator transaction binding the contract method 0x28f88a65.
//
// Solidity: function transitState(uint256 id, uint256 oldState, uint256 newState, bool isOldStateGenesis, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_StateV2 *StateV2Transactor) TransitState(opts *bind.TransactOpts, id *big.Int, oldState *big.Int, newState *big.Int, isOldStateGenesis bool, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _StateV2.contract.Transact(opts, "transitState", id, oldState, newState, isOldStateGenesis, a, b, c)
}

// TransitState is a paid mutator transaction binding the contract method 0x28f88a65.
//
// Solidity: function transitState(uint256 id, uint256 oldState, uint256 newState, bool isOldStateGenesis, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_StateV2 *StateV2Session) TransitState(id *big.Int, oldState *big.Int, newState *big.Int, isOldStateGenesis bool, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _StateV2.Contract.TransitState(&_StateV2.TransactOpts, id, oldState, newState, isOldStateGenesis, a, b, c)
}

// TransitState is a paid mutator transaction binding the contract method 0x28f88a65.
//
// Solidity: function transitState(uint256 id, uint256 oldState, uint256 newState, bool isOldStateGenesis, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_StateV2 *StateV2TransactorSession) TransitState(id *big.Int, oldState *big.Int, newState *big.Int, isOldStateGenesis bool, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _StateV2.Contract.TransitState(&_StateV2.TransactOpts, id, oldState, newState, isOldStateGenesis, a, b, c)
}

// StateV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StateV2 contract.
type StateV2InitializedIterator struct {
	Event *StateV2Initialized // Event containing the contract specifics and raw log

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
func (it *StateV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateV2Initialized)
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
		it.Event = new(StateV2Initialized)
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
func (it *StateV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateV2Initialized represents a Initialized event raised by the StateV2 contract.
type StateV2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StateV2 *StateV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*StateV2InitializedIterator, error) {

	logs, sub, err := _StateV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StateV2InitializedIterator{contract: _StateV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StateV2 *StateV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StateV2Initialized) (event.Subscription, error) {

	logs, sub, err := _StateV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateV2Initialized)
				if err := _StateV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StateV2 *StateV2Filterer) ParseInitialized(log types.Log) (*StateV2Initialized, error) {
	event := new(StateV2Initialized)
	if err := _StateV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateV2OwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the StateV2 contract.
type StateV2OwnershipTransferStartedIterator struct {
	Event *StateV2OwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *StateV2OwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateV2OwnershipTransferStarted)
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
		it.Event = new(StateV2OwnershipTransferStarted)
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
func (it *StateV2OwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateV2OwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateV2OwnershipTransferStarted represents a OwnershipTransferStarted event raised by the StateV2 contract.
type StateV2OwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_StateV2 *StateV2Filterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StateV2OwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StateV2.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StateV2OwnershipTransferStartedIterator{contract: _StateV2.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_StateV2 *StateV2Filterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *StateV2OwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StateV2.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateV2OwnershipTransferStarted)
				if err := _StateV2.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_StateV2 *StateV2Filterer) ParseOwnershipTransferStarted(log types.Log) (*StateV2OwnershipTransferStarted, error) {
	event := new(StateV2OwnershipTransferStarted)
	if err := _StateV2.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StateV2 contract.
type StateV2OwnershipTransferredIterator struct {
	Event *StateV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StateV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateV2OwnershipTransferred)
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
		it.Event = new(StateV2OwnershipTransferred)
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
func (it *StateV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateV2OwnershipTransferred represents a OwnershipTransferred event raised by the StateV2 contract.
type StateV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StateV2 *StateV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StateV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StateV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StateV2OwnershipTransferredIterator{contract: _StateV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StateV2 *StateV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StateV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StateV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateV2OwnershipTransferred)
				if err := _StateV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StateV2 *StateV2Filterer) ParseOwnershipTransferred(log types.Log) (*StateV2OwnershipTransferred, error) {
	event := new(StateV2OwnershipTransferred)
	if err := _StateV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
