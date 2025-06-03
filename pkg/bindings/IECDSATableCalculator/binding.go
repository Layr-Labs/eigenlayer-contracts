// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IECDSATableCalculator

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

// IECDSATableCalculatorTypesECDSAOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IECDSATableCalculatorTypesECDSAOperatorInfo struct {
	Pubkey  common.Address
	Weights []*big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// IECDSATableCalculatorMetaData contains all meta data concerning the IECDSATableCalculator contract.
var IECDSATableCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"calculateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorInfos\",\"type\":\"tuple[]\",\"internalType\":\"structIECDSATableCalculatorTypes.ECDSAOperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weights\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorTableBytes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"}]",
}

// IECDSATableCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use IECDSATableCalculatorMetaData.ABI instead.
var IECDSATableCalculatorABI = IECDSATableCalculatorMetaData.ABI

// IECDSATableCalculator is an auto generated Go binding around an Ethereum contract.
type IECDSATableCalculator struct {
	IECDSATableCalculatorCaller     // Read-only binding to the contract
	IECDSATableCalculatorTransactor // Write-only binding to the contract
	IECDSATableCalculatorFilterer   // Log filterer for contract events
}

// IECDSATableCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IECDSATableCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IECDSATableCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IECDSATableCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IECDSATableCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IECDSATableCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IECDSATableCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IECDSATableCalculatorSession struct {
	Contract     *IECDSATableCalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IECDSATableCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IECDSATableCalculatorCallerSession struct {
	Contract *IECDSATableCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IECDSATableCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IECDSATableCalculatorTransactorSession struct {
	Contract     *IECDSATableCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IECDSATableCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IECDSATableCalculatorRaw struct {
	Contract *IECDSATableCalculator // Generic contract binding to access the raw methods on
}

// IECDSATableCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IECDSATableCalculatorCallerRaw struct {
	Contract *IECDSATableCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// IECDSATableCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IECDSATableCalculatorTransactorRaw struct {
	Contract *IECDSATableCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIECDSATableCalculator creates a new instance of IECDSATableCalculator, bound to a specific deployed contract.
func NewIECDSATableCalculator(address common.Address, backend bind.ContractBackend) (*IECDSATableCalculator, error) {
	contract, err := bindIECDSATableCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IECDSATableCalculator{IECDSATableCalculatorCaller: IECDSATableCalculatorCaller{contract: contract}, IECDSATableCalculatorTransactor: IECDSATableCalculatorTransactor{contract: contract}, IECDSATableCalculatorFilterer: IECDSATableCalculatorFilterer{contract: contract}}, nil
}

// NewIECDSATableCalculatorCaller creates a new read-only instance of IECDSATableCalculator, bound to a specific deployed contract.
func NewIECDSATableCalculatorCaller(address common.Address, caller bind.ContractCaller) (*IECDSATableCalculatorCaller, error) {
	contract, err := bindIECDSATableCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IECDSATableCalculatorCaller{contract: contract}, nil
}

// NewIECDSATableCalculatorTransactor creates a new write-only instance of IECDSATableCalculator, bound to a specific deployed contract.
func NewIECDSATableCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IECDSATableCalculatorTransactor, error) {
	contract, err := bindIECDSATableCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IECDSATableCalculatorTransactor{contract: contract}, nil
}

// NewIECDSATableCalculatorFilterer creates a new log filterer instance of IECDSATableCalculator, bound to a specific deployed contract.
func NewIECDSATableCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IECDSATableCalculatorFilterer, error) {
	contract, err := bindIECDSATableCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IECDSATableCalculatorFilterer{contract: contract}, nil
}

// bindIECDSATableCalculator binds a generic wrapper to an already deployed contract.
func bindIECDSATableCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IECDSATableCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IECDSATableCalculator *IECDSATableCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IECDSATableCalculator.Contract.IECDSATableCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IECDSATableCalculator *IECDSATableCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IECDSATableCalculator.Contract.IECDSATableCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IECDSATableCalculator *IECDSATableCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IECDSATableCalculator.Contract.IECDSATableCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IECDSATableCalculator *IECDSATableCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IECDSATableCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IECDSATableCalculator *IECDSATableCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IECDSATableCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IECDSATableCalculator *IECDSATableCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IECDSATableCalculator.Contract.contract.Transact(opts, method, params...)
}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((address,uint96[])[] operatorInfos)
func (_IECDSATableCalculator *IECDSATableCalculatorCaller) CalculateOperatorTable(opts *bind.CallOpts, operatorSet OperatorSet) ([]IECDSATableCalculatorTypesECDSAOperatorInfo, error) {
	var out []interface{}
	err := _IECDSATableCalculator.contract.Call(opts, &out, "calculateOperatorTable", operatorSet)

	if err != nil {
		return *new([]IECDSATableCalculatorTypesECDSAOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IECDSATableCalculatorTypesECDSAOperatorInfo)).(*[]IECDSATableCalculatorTypesECDSAOperatorInfo)

	return out0, err

}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((address,uint96[])[] operatorInfos)
func (_IECDSATableCalculator *IECDSATableCalculatorSession) CalculateOperatorTable(operatorSet OperatorSet) ([]IECDSATableCalculatorTypesECDSAOperatorInfo, error) {
	return _IECDSATableCalculator.Contract.CalculateOperatorTable(&_IECDSATableCalculator.CallOpts, operatorSet)
}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((address,uint96[])[] operatorInfos)
func (_IECDSATableCalculator *IECDSATableCalculatorCallerSession) CalculateOperatorTable(operatorSet OperatorSet) ([]IECDSATableCalculatorTypesECDSAOperatorInfo, error) {
	return _IECDSATableCalculator.Contract.CalculateOperatorTable(&_IECDSATableCalculator.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_IECDSATableCalculator *IECDSATableCalculatorCaller) CalculateOperatorTableBytes(opts *bind.CallOpts, operatorSet OperatorSet) ([]byte, error) {
	var out []interface{}
	err := _IECDSATableCalculator.contract.Call(opts, &out, "calculateOperatorTableBytes", operatorSet)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_IECDSATableCalculator *IECDSATableCalculatorSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _IECDSATableCalculator.Contract.CalculateOperatorTableBytes(&_IECDSATableCalculator.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_IECDSATableCalculator *IECDSATableCalculatorCallerSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _IECDSATableCalculator.Contract.CalculateOperatorTableBytes(&_IECDSATableCalculator.CallOpts, operatorSet)
}
