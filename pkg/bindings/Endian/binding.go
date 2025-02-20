// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Endian

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

// EndianMetaData contains all meta data concerning the Endian contract.
var EndianMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea26469706673582212207ab822e4bec68ffbff791aaa98829dc58d61a18245bb4328599d5053f59bc75764736f6c634300081b0033",
}

// EndianABI is the input ABI used to generate the binding from.
// Deprecated: Use EndianMetaData.ABI instead.
var EndianABI = EndianMetaData.ABI

// EndianBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EndianMetaData.Bin instead.
var EndianBin = EndianMetaData.Bin

// DeployEndian deploys a new Ethereum contract, binding an instance of Endian to it.
func DeployEndian(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Endian, error) {
	parsed, err := EndianMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EndianBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Endian{EndianCaller: EndianCaller{contract: contract}, EndianTransactor: EndianTransactor{contract: contract}, EndianFilterer: EndianFilterer{contract: contract}}, nil
}

// Endian is an auto generated Go binding around an Ethereum contract.
type Endian struct {
	EndianCaller     // Read-only binding to the contract
	EndianTransactor // Write-only binding to the contract
	EndianFilterer   // Log filterer for contract events
}

// EndianCaller is an auto generated read-only Go binding around an Ethereum contract.
type EndianCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EndianTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EndianTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EndianFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EndianFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EndianSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EndianSession struct {
	Contract     *Endian           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EndianCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EndianCallerSession struct {
	Contract *EndianCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EndianTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EndianTransactorSession struct {
	Contract     *EndianTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EndianRaw is an auto generated low-level Go binding around an Ethereum contract.
type EndianRaw struct {
	Contract *Endian // Generic contract binding to access the raw methods on
}

// EndianCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EndianCallerRaw struct {
	Contract *EndianCaller // Generic read-only contract binding to access the raw methods on
}

// EndianTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EndianTransactorRaw struct {
	Contract *EndianTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEndian creates a new instance of Endian, bound to a specific deployed contract.
func NewEndian(address common.Address, backend bind.ContractBackend) (*Endian, error) {
	contract, err := bindEndian(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Endian{EndianCaller: EndianCaller{contract: contract}, EndianTransactor: EndianTransactor{contract: contract}, EndianFilterer: EndianFilterer{contract: contract}}, nil
}

// NewEndianCaller creates a new read-only instance of Endian, bound to a specific deployed contract.
func NewEndianCaller(address common.Address, caller bind.ContractCaller) (*EndianCaller, error) {
	contract, err := bindEndian(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EndianCaller{contract: contract}, nil
}

// NewEndianTransactor creates a new write-only instance of Endian, bound to a specific deployed contract.
func NewEndianTransactor(address common.Address, transactor bind.ContractTransactor) (*EndianTransactor, error) {
	contract, err := bindEndian(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EndianTransactor{contract: contract}, nil
}

// NewEndianFilterer creates a new log filterer instance of Endian, bound to a specific deployed contract.
func NewEndianFilterer(address common.Address, filterer bind.ContractFilterer) (*EndianFilterer, error) {
	contract, err := bindEndian(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EndianFilterer{contract: contract}, nil
}

// bindEndian binds a generic wrapper to an already deployed contract.
func bindEndian(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EndianMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Endian *EndianRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Endian.Contract.EndianCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Endian *EndianRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Endian.Contract.EndianTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Endian *EndianRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Endian.Contract.EndianTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Endian *EndianCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Endian.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Endian *EndianTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Endian.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Endian *EndianTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Endian.Contract.contract.Transact(opts, method, params...)
}
