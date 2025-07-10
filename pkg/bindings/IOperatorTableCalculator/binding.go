// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IOperatorTableCalculator

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

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// IOperatorTableCalculatorMetaData contains all meta data concerning the IOperatorTableCalculator contract.
var IOperatorTableCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"calculateOperatorTableBytes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetWeights\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"weights\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeights\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"}]",
}

// IOperatorTableCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use IOperatorTableCalculatorMetaData.ABI instead.
var IOperatorTableCalculatorABI = IOperatorTableCalculatorMetaData.ABI

// IOperatorTableCalculator is an auto generated Go binding around an Ethereum contract.
type IOperatorTableCalculator struct {
	IOperatorTableCalculatorCaller     // Read-only binding to the contract
	IOperatorTableCalculatorTransactor // Write-only binding to the contract
	IOperatorTableCalculatorFilterer   // Log filterer for contract events
}

// IOperatorTableCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOperatorTableCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOperatorTableCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOperatorTableCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOperatorTableCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOperatorTableCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOperatorTableCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOperatorTableCalculatorSession struct {
	Contract     *IOperatorTableCalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IOperatorTableCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOperatorTableCalculatorCallerSession struct {
	Contract *IOperatorTableCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IOperatorTableCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOperatorTableCalculatorTransactorSession struct {
	Contract     *IOperatorTableCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IOperatorTableCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOperatorTableCalculatorRaw struct {
	Contract *IOperatorTableCalculator // Generic contract binding to access the raw methods on
}

// IOperatorTableCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOperatorTableCalculatorCallerRaw struct {
	Contract *IOperatorTableCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// IOperatorTableCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOperatorTableCalculatorTransactorRaw struct {
	Contract *IOperatorTableCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOperatorTableCalculator creates a new instance of IOperatorTableCalculator, bound to a specific deployed contract.
func NewIOperatorTableCalculator(address common.Address, backend bind.ContractBackend) (*IOperatorTableCalculator, error) {
	contract, err := bindIOperatorTableCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOperatorTableCalculator{IOperatorTableCalculatorCaller: IOperatorTableCalculatorCaller{contract: contract}, IOperatorTableCalculatorTransactor: IOperatorTableCalculatorTransactor{contract: contract}, IOperatorTableCalculatorFilterer: IOperatorTableCalculatorFilterer{contract: contract}}, nil
}

// NewIOperatorTableCalculatorCaller creates a new read-only instance of IOperatorTableCalculator, bound to a specific deployed contract.
func NewIOperatorTableCalculatorCaller(address common.Address, caller bind.ContractCaller) (*IOperatorTableCalculatorCaller, error) {
	contract, err := bindIOperatorTableCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOperatorTableCalculatorCaller{contract: contract}, nil
}

// NewIOperatorTableCalculatorTransactor creates a new write-only instance of IOperatorTableCalculator, bound to a specific deployed contract.
func NewIOperatorTableCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IOperatorTableCalculatorTransactor, error) {
	contract, err := bindIOperatorTableCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOperatorTableCalculatorTransactor{contract: contract}, nil
}

// NewIOperatorTableCalculatorFilterer creates a new log filterer instance of IOperatorTableCalculator, bound to a specific deployed contract.
func NewIOperatorTableCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IOperatorTableCalculatorFilterer, error) {
	contract, err := bindIOperatorTableCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOperatorTableCalculatorFilterer{contract: contract}, nil
}

// bindIOperatorTableCalculator binds a generic wrapper to an already deployed contract.
func bindIOperatorTableCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOperatorTableCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOperatorTableCalculator *IOperatorTableCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOperatorTableCalculator.Contract.IOperatorTableCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOperatorTableCalculator *IOperatorTableCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOperatorTableCalculator.Contract.IOperatorTableCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOperatorTableCalculator *IOperatorTableCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOperatorTableCalculator.Contract.IOperatorTableCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOperatorTableCalculator *IOperatorTableCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOperatorTableCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOperatorTableCalculator *IOperatorTableCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOperatorTableCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOperatorTableCalculator *IOperatorTableCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOperatorTableCalculator.Contract.contract.Transact(opts, method, params...)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_IOperatorTableCalculator *IOperatorTableCalculatorCaller) CalculateOperatorTableBytes(opts *bind.CallOpts, operatorSet OperatorSet) ([]byte, error) {
	var out []interface{}
	err := _IOperatorTableCalculator.contract.Call(opts, &out, "calculateOperatorTableBytes", operatorSet)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_IOperatorTableCalculator *IOperatorTableCalculatorSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _IOperatorTableCalculator.Contract.CalculateOperatorTableBytes(&_IOperatorTableCalculator.CallOpts, operatorSet)
}

// CalculateOperatorTableBytes is a free data retrieval call binding the contract method 0x41ee6d0e.
//
// Solidity: function calculateOperatorTableBytes((address,uint32) operatorSet) view returns(bytes operatorTableBytes)
func (_IOperatorTableCalculator *IOperatorTableCalculatorCallerSession) CalculateOperatorTableBytes(operatorSet OperatorSet) ([]byte, error) {
	return _IOperatorTableCalculator.Contract.CalculateOperatorTableBytes(&_IOperatorTableCalculator.CallOpts, operatorSet)
}

// GetOperatorSetWeights is a free data retrieval call binding the contract method 0xbff2de25.
//
// Solidity: function getOperatorSetWeights((address,uint32) operatorSet) view returns(address[] operators, uint256[][] weights)
func (_IOperatorTableCalculator *IOperatorTableCalculatorCaller) GetOperatorSetWeights(opts *bind.CallOpts, operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	var out []interface{}
	err := _IOperatorTableCalculator.contract.Call(opts, &out, "getOperatorSetWeights", operatorSet)

	outstruct := new(struct {
		Operators []common.Address
		Weights   [][]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Operators = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Weights = *abi.ConvertType(out[1], new([][]*big.Int)).(*[][]*big.Int)

	return *outstruct, err

}

// GetOperatorSetWeights is a free data retrieval call binding the contract method 0xbff2de25.
//
// Solidity: function getOperatorSetWeights((address,uint32) operatorSet) view returns(address[] operators, uint256[][] weights)
func (_IOperatorTableCalculator *IOperatorTableCalculatorSession) GetOperatorSetWeights(operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	return _IOperatorTableCalculator.Contract.GetOperatorSetWeights(&_IOperatorTableCalculator.CallOpts, operatorSet)
}

// GetOperatorSetWeights is a free data retrieval call binding the contract method 0xbff2de25.
//
// Solidity: function getOperatorSetWeights((address,uint32) operatorSet) view returns(address[] operators, uint256[][] weights)
func (_IOperatorTableCalculator *IOperatorTableCalculatorCallerSession) GetOperatorSetWeights(operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	return _IOperatorTableCalculator.Contract.GetOperatorSetWeights(&_IOperatorTableCalculator.CallOpts, operatorSet)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0xbdf3ad27.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet, address operator) view returns(uint256[] weights)
func (_IOperatorTableCalculator *IOperatorTableCalculatorCaller) GetOperatorWeights(opts *bind.CallOpts, operatorSet OperatorSet, operator common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _IOperatorTableCalculator.contract.Call(opts, &out, "getOperatorWeights", operatorSet, operator)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOperatorWeights is a free data retrieval call binding the contract method 0xbdf3ad27.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet, address operator) view returns(uint256[] weights)
func (_IOperatorTableCalculator *IOperatorTableCalculatorSession) GetOperatorWeights(operatorSet OperatorSet, operator common.Address) ([]*big.Int, error) {
	return _IOperatorTableCalculator.Contract.GetOperatorWeights(&_IOperatorTableCalculator.CallOpts, operatorSet, operator)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0xbdf3ad27.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet, address operator) view returns(uint256[] weights)
func (_IOperatorTableCalculator *IOperatorTableCalculatorCallerSession) GetOperatorWeights(operatorSet OperatorSet, operator common.Address) ([]*big.Int, error) {
	return _IOperatorTableCalculator.Contract.GetOperatorWeights(&_IOperatorTableCalculator.CallOpts, operatorSet, operator)
}
