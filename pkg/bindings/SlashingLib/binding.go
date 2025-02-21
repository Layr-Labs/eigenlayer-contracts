// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SlashingLib

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

// SlashingLibMetaData contains all meta data concerning the SlashingLib contract.
var SlashingLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea26469706673582212208f1bde89f8f3aa0f34d6051969b229fa83efd16a534afe1769532efc8662438a64736f6c634300081b0033",
}

// SlashingLibABI is the input ABI used to generate the binding from.
// Deprecated: Use SlashingLibMetaData.ABI instead.
var SlashingLibABI = SlashingLibMetaData.ABI

// SlashingLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlashingLibMetaData.Bin instead.
var SlashingLibBin = SlashingLibMetaData.Bin

// DeploySlashingLib deploys a new Ethereum contract, binding an instance of SlashingLib to it.
func DeploySlashingLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SlashingLib, error) {
	parsed, err := SlashingLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlashingLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SlashingLib{SlashingLibCaller: SlashingLibCaller{contract: contract}, SlashingLibTransactor: SlashingLibTransactor{contract: contract}, SlashingLibFilterer: SlashingLibFilterer{contract: contract}}, nil
}

// SlashingLib is an auto generated Go binding around an Ethereum contract.
type SlashingLib struct {
	SlashingLibCaller     // Read-only binding to the contract
	SlashingLibTransactor // Write-only binding to the contract
	SlashingLibFilterer   // Log filterer for contract events
}

// SlashingLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlashingLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlashingLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlashingLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlashingLibSession struct {
	Contract     *SlashingLib      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlashingLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlashingLibCallerSession struct {
	Contract *SlashingLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SlashingLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlashingLibTransactorSession struct {
	Contract     *SlashingLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SlashingLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlashingLibRaw struct {
	Contract *SlashingLib // Generic contract binding to access the raw methods on
}

// SlashingLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlashingLibCallerRaw struct {
	Contract *SlashingLibCaller // Generic read-only contract binding to access the raw methods on
}

// SlashingLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlashingLibTransactorRaw struct {
	Contract *SlashingLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlashingLib creates a new instance of SlashingLib, bound to a specific deployed contract.
func NewSlashingLib(address common.Address, backend bind.ContractBackend) (*SlashingLib, error) {
	contract, err := bindSlashingLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlashingLib{SlashingLibCaller: SlashingLibCaller{contract: contract}, SlashingLibTransactor: SlashingLibTransactor{contract: contract}, SlashingLibFilterer: SlashingLibFilterer{contract: contract}}, nil
}

// NewSlashingLibCaller creates a new read-only instance of SlashingLib, bound to a specific deployed contract.
func NewSlashingLibCaller(address common.Address, caller bind.ContractCaller) (*SlashingLibCaller, error) {
	contract, err := bindSlashingLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlashingLibCaller{contract: contract}, nil
}

// NewSlashingLibTransactor creates a new write-only instance of SlashingLib, bound to a specific deployed contract.
func NewSlashingLibTransactor(address common.Address, transactor bind.ContractTransactor) (*SlashingLibTransactor, error) {
	contract, err := bindSlashingLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlashingLibTransactor{contract: contract}, nil
}

// NewSlashingLibFilterer creates a new log filterer instance of SlashingLib, bound to a specific deployed contract.
func NewSlashingLibFilterer(address common.Address, filterer bind.ContractFilterer) (*SlashingLibFilterer, error) {
	contract, err := bindSlashingLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlashingLibFilterer{contract: contract}, nil
}

// bindSlashingLib binds a generic wrapper to an already deployed contract.
func bindSlashingLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlashingLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashingLib *SlashingLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashingLib.Contract.SlashingLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashingLib *SlashingLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashingLib.Contract.SlashingLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashingLib *SlashingLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashingLib.Contract.SlashingLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashingLib *SlashingLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashingLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashingLib *SlashingLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashingLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashingLib *SlashingLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashingLib.Contract.contract.Transact(opts, method, params...)
}
