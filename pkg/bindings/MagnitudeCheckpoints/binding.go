// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MagnitudeCheckpoints

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

// MagnitudeCheckpointsMetaData contains all meta data concerning the MagnitudeCheckpoints contract.
var MagnitudeCheckpointsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d403c7152d652d1a100c83daf55d72392c025180800bf29f79d386c0af19a05d64736f6c634300080c0033",
}

// MagnitudeCheckpointsABI is the input ABI used to generate the binding from.
// Deprecated: Use MagnitudeCheckpointsMetaData.ABI instead.
var MagnitudeCheckpointsABI = MagnitudeCheckpointsMetaData.ABI

// MagnitudeCheckpointsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MagnitudeCheckpointsMetaData.Bin instead.
var MagnitudeCheckpointsBin = MagnitudeCheckpointsMetaData.Bin

// DeployMagnitudeCheckpoints deploys a new Ethereum contract, binding an instance of MagnitudeCheckpoints to it.
func DeployMagnitudeCheckpoints(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MagnitudeCheckpoints, error) {
	parsed, err := MagnitudeCheckpointsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MagnitudeCheckpointsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MagnitudeCheckpoints{MagnitudeCheckpointsCaller: MagnitudeCheckpointsCaller{contract: contract}, MagnitudeCheckpointsTransactor: MagnitudeCheckpointsTransactor{contract: contract}, MagnitudeCheckpointsFilterer: MagnitudeCheckpointsFilterer{contract: contract}}, nil
}

// MagnitudeCheckpoints is an auto generated Go binding around an Ethereum contract.
type MagnitudeCheckpoints struct {
	MagnitudeCheckpointsCaller     // Read-only binding to the contract
	MagnitudeCheckpointsTransactor // Write-only binding to the contract
	MagnitudeCheckpointsFilterer   // Log filterer for contract events
}

// MagnitudeCheckpointsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MagnitudeCheckpointsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagnitudeCheckpointsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MagnitudeCheckpointsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagnitudeCheckpointsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MagnitudeCheckpointsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagnitudeCheckpointsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MagnitudeCheckpointsSession struct {
	Contract     *MagnitudeCheckpoints // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MagnitudeCheckpointsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MagnitudeCheckpointsCallerSession struct {
	Contract *MagnitudeCheckpointsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// MagnitudeCheckpointsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MagnitudeCheckpointsTransactorSession struct {
	Contract     *MagnitudeCheckpointsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// MagnitudeCheckpointsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MagnitudeCheckpointsRaw struct {
	Contract *MagnitudeCheckpoints // Generic contract binding to access the raw methods on
}

// MagnitudeCheckpointsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MagnitudeCheckpointsCallerRaw struct {
	Contract *MagnitudeCheckpointsCaller // Generic read-only contract binding to access the raw methods on
}

// MagnitudeCheckpointsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MagnitudeCheckpointsTransactorRaw struct {
	Contract *MagnitudeCheckpointsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMagnitudeCheckpoints creates a new instance of MagnitudeCheckpoints, bound to a specific deployed contract.
func NewMagnitudeCheckpoints(address common.Address, backend bind.ContractBackend) (*MagnitudeCheckpoints, error) {
	contract, err := bindMagnitudeCheckpoints(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MagnitudeCheckpoints{MagnitudeCheckpointsCaller: MagnitudeCheckpointsCaller{contract: contract}, MagnitudeCheckpointsTransactor: MagnitudeCheckpointsTransactor{contract: contract}, MagnitudeCheckpointsFilterer: MagnitudeCheckpointsFilterer{contract: contract}}, nil
}

// NewMagnitudeCheckpointsCaller creates a new read-only instance of MagnitudeCheckpoints, bound to a specific deployed contract.
func NewMagnitudeCheckpointsCaller(address common.Address, caller bind.ContractCaller) (*MagnitudeCheckpointsCaller, error) {
	contract, err := bindMagnitudeCheckpoints(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MagnitudeCheckpointsCaller{contract: contract}, nil
}

// NewMagnitudeCheckpointsTransactor creates a new write-only instance of MagnitudeCheckpoints, bound to a specific deployed contract.
func NewMagnitudeCheckpointsTransactor(address common.Address, transactor bind.ContractTransactor) (*MagnitudeCheckpointsTransactor, error) {
	contract, err := bindMagnitudeCheckpoints(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MagnitudeCheckpointsTransactor{contract: contract}, nil
}

// NewMagnitudeCheckpointsFilterer creates a new log filterer instance of MagnitudeCheckpoints, bound to a specific deployed contract.
func NewMagnitudeCheckpointsFilterer(address common.Address, filterer bind.ContractFilterer) (*MagnitudeCheckpointsFilterer, error) {
	contract, err := bindMagnitudeCheckpoints(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MagnitudeCheckpointsFilterer{contract: contract}, nil
}

// bindMagnitudeCheckpoints binds a generic wrapper to an already deployed contract.
func bindMagnitudeCheckpoints(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MagnitudeCheckpointsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MagnitudeCheckpoints *MagnitudeCheckpointsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MagnitudeCheckpoints.Contract.MagnitudeCheckpointsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MagnitudeCheckpoints *MagnitudeCheckpointsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MagnitudeCheckpoints.Contract.MagnitudeCheckpointsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MagnitudeCheckpoints *MagnitudeCheckpointsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MagnitudeCheckpoints.Contract.MagnitudeCheckpointsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MagnitudeCheckpoints *MagnitudeCheckpointsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MagnitudeCheckpoints.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MagnitudeCheckpoints *MagnitudeCheckpointsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MagnitudeCheckpoints.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MagnitudeCheckpoints *MagnitudeCheckpointsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MagnitudeCheckpoints.Contract.contract.Transact(opts, method, params...)
}
