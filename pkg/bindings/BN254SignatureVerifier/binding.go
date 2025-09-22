// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BN254SignatureVerifier

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

// BN254SignatureVerifierMetaData contains all meta data concerning the BN254SignatureVerifier contract.
var BN254SignatureVerifierMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea264697066735822122099890b68aab0bf243dce6fa1d9c78bbdb3bbd47fb6b3dbc51c4b9691de9707ef64736f6c634300081b0033",
}

// BN254SignatureVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use BN254SignatureVerifierMetaData.ABI instead.
var BN254SignatureVerifierABI = BN254SignatureVerifierMetaData.ABI

// BN254SignatureVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BN254SignatureVerifierMetaData.Bin instead.
var BN254SignatureVerifierBin = BN254SignatureVerifierMetaData.Bin

// DeployBN254SignatureVerifier deploys a new Ethereum contract, binding an instance of BN254SignatureVerifier to it.
func DeployBN254SignatureVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BN254SignatureVerifier, error) {
	parsed, err := BN254SignatureVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BN254SignatureVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BN254SignatureVerifier{BN254SignatureVerifierCaller: BN254SignatureVerifierCaller{contract: contract}, BN254SignatureVerifierTransactor: BN254SignatureVerifierTransactor{contract: contract}, BN254SignatureVerifierFilterer: BN254SignatureVerifierFilterer{contract: contract}}, nil
}

// BN254SignatureVerifier is an auto generated Go binding around an Ethereum contract.
type BN254SignatureVerifier struct {
	BN254SignatureVerifierCaller     // Read-only binding to the contract
	BN254SignatureVerifierTransactor // Write-only binding to the contract
	BN254SignatureVerifierFilterer   // Log filterer for contract events
}

// BN254SignatureVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type BN254SignatureVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254SignatureVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BN254SignatureVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254SignatureVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BN254SignatureVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254SignatureVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BN254SignatureVerifierSession struct {
	Contract     *BN254SignatureVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BN254SignatureVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BN254SignatureVerifierCallerSession struct {
	Contract *BN254SignatureVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// BN254SignatureVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BN254SignatureVerifierTransactorSession struct {
	Contract     *BN254SignatureVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// BN254SignatureVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type BN254SignatureVerifierRaw struct {
	Contract *BN254SignatureVerifier // Generic contract binding to access the raw methods on
}

// BN254SignatureVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BN254SignatureVerifierCallerRaw struct {
	Contract *BN254SignatureVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// BN254SignatureVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BN254SignatureVerifierTransactorRaw struct {
	Contract *BN254SignatureVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBN254SignatureVerifier creates a new instance of BN254SignatureVerifier, bound to a specific deployed contract.
func NewBN254SignatureVerifier(address common.Address, backend bind.ContractBackend) (*BN254SignatureVerifier, error) {
	contract, err := bindBN254SignatureVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BN254SignatureVerifier{BN254SignatureVerifierCaller: BN254SignatureVerifierCaller{contract: contract}, BN254SignatureVerifierTransactor: BN254SignatureVerifierTransactor{contract: contract}, BN254SignatureVerifierFilterer: BN254SignatureVerifierFilterer{contract: contract}}, nil
}

// NewBN254SignatureVerifierCaller creates a new read-only instance of BN254SignatureVerifier, bound to a specific deployed contract.
func NewBN254SignatureVerifierCaller(address common.Address, caller bind.ContractCaller) (*BN254SignatureVerifierCaller, error) {
	contract, err := bindBN254SignatureVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BN254SignatureVerifierCaller{contract: contract}, nil
}

// NewBN254SignatureVerifierTransactor creates a new write-only instance of BN254SignatureVerifier, bound to a specific deployed contract.
func NewBN254SignatureVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*BN254SignatureVerifierTransactor, error) {
	contract, err := bindBN254SignatureVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BN254SignatureVerifierTransactor{contract: contract}, nil
}

// NewBN254SignatureVerifierFilterer creates a new log filterer instance of BN254SignatureVerifier, bound to a specific deployed contract.
func NewBN254SignatureVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*BN254SignatureVerifierFilterer, error) {
	contract, err := bindBN254SignatureVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BN254SignatureVerifierFilterer{contract: contract}, nil
}

// bindBN254SignatureVerifier binds a generic wrapper to an already deployed contract.
func bindBN254SignatureVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BN254SignatureVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN254SignatureVerifier *BN254SignatureVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BN254SignatureVerifier.Contract.BN254SignatureVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN254SignatureVerifier *BN254SignatureVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN254SignatureVerifier.Contract.BN254SignatureVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN254SignatureVerifier *BN254SignatureVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN254SignatureVerifier.Contract.BN254SignatureVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN254SignatureVerifier *BN254SignatureVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BN254SignatureVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN254SignatureVerifier *BN254SignatureVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN254SignatureVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN254SignatureVerifier *BN254SignatureVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN254SignatureVerifier.Contract.contract.Transact(opts, method, params...)
}
