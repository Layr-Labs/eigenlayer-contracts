// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BN254

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

// BN254MetaData contains all meta data concerning the BN254 contract.
var BN254MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"ECAddFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECMulFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECPairingFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpModFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ScalarTooLarge\",\"inputs\":[]}]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea2646970667358221220d084bedd451fadde14a6590fa5fb3e40f2250b7f49f9228d15f27cae51d34f7464736f6c634300081b0033",
}

// BN254ABI is the input ABI used to generate the binding from.
// Deprecated: Use BN254MetaData.ABI instead.
var BN254ABI = BN254MetaData.ABI

// BN254Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BN254MetaData.Bin instead.
var BN254Bin = BN254MetaData.Bin

// DeployBN254 deploys a new Ethereum contract, binding an instance of BN254 to it.
func DeployBN254(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BN254, error) {
	parsed, err := BN254MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BN254Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BN254{BN254Caller: BN254Caller{contract: contract}, BN254Transactor: BN254Transactor{contract: contract}, BN254Filterer: BN254Filterer{contract: contract}}, nil
}

// BN254 is an auto generated Go binding around an Ethereum contract.
type BN254 struct {
	BN254Caller     // Read-only binding to the contract
	BN254Transactor // Write-only binding to the contract
	BN254Filterer   // Log filterer for contract events
}

// BN254Caller is an auto generated read-only Go binding around an Ethereum contract.
type BN254Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BN254Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BN254Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN254Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BN254Session struct {
	Contract     *BN254            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BN254CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BN254CallerSession struct {
	Contract *BN254Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BN254TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BN254TransactorSession struct {
	Contract     *BN254Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BN254Raw is an auto generated low-level Go binding around an Ethereum contract.
type BN254Raw struct {
	Contract *BN254 // Generic contract binding to access the raw methods on
}

// BN254CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BN254CallerRaw struct {
	Contract *BN254Caller // Generic read-only contract binding to access the raw methods on
}

// BN254TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BN254TransactorRaw struct {
	Contract *BN254Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBN254 creates a new instance of BN254, bound to a specific deployed contract.
func NewBN254(address common.Address, backend bind.ContractBackend) (*BN254, error) {
	contract, err := bindBN254(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BN254{BN254Caller: BN254Caller{contract: contract}, BN254Transactor: BN254Transactor{contract: contract}, BN254Filterer: BN254Filterer{contract: contract}}, nil
}

// NewBN254Caller creates a new read-only instance of BN254, bound to a specific deployed contract.
func NewBN254Caller(address common.Address, caller bind.ContractCaller) (*BN254Caller, error) {
	contract, err := bindBN254(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BN254Caller{contract: contract}, nil
}

// NewBN254Transactor creates a new write-only instance of BN254, bound to a specific deployed contract.
func NewBN254Transactor(address common.Address, transactor bind.ContractTransactor) (*BN254Transactor, error) {
	contract, err := bindBN254(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BN254Transactor{contract: contract}, nil
}

// NewBN254Filterer creates a new log filterer instance of BN254, bound to a specific deployed contract.
func NewBN254Filterer(address common.Address, filterer bind.ContractFilterer) (*BN254Filterer, error) {
	contract, err := bindBN254(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BN254Filterer{contract: contract}, nil
}

// bindBN254 binds a generic wrapper to an already deployed contract.
func bindBN254(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BN254MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN254 *BN254Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BN254.Contract.BN254Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN254 *BN254Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN254.Contract.BN254Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN254 *BN254Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN254.Contract.BN254Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN254 *BN254CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BN254.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN254 *BN254TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN254.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN254 *BN254TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN254.Contract.contract.Transact(opts, method, params...)
}
