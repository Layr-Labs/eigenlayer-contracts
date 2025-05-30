// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ISlashEscrow

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

// ISlashEscrowMetaData contains all meta data concerning the ISlashEscrow contract.
var ISlashEscrowMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"releaseTokens\",\"inputs\":[{\"name\":\"slashEscrowFactory\",\"type\":\"address\",\"internalType\":\"contractISlashEscrowFactory\"},{\"name\":\"slashEscrowImplementation\",\"type\":\"address\",\"internalType\":\"contractISlashEscrow\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyDeploymentParameters\",\"inputs\":[{\"name\":\"slashEscrowFactory\",\"type\":\"address\",\"internalType\":\"contractISlashEscrowFactory\"},{\"name\":\"slashEscrowImplementation\",\"type\":\"address\",\"internalType\":\"contractISlashEscrow\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidDeploymentParameters\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySlashEscrowFactory\",\"inputs\":[]}]",
}

// ISlashEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use ISlashEscrowMetaData.ABI instead.
var ISlashEscrowABI = ISlashEscrowMetaData.ABI

// ISlashEscrow is an auto generated Go binding around an Ethereum contract.
type ISlashEscrow struct {
	ISlashEscrowCaller     // Read-only binding to the contract
	ISlashEscrowTransactor // Write-only binding to the contract
	ISlashEscrowFilterer   // Log filterer for contract events
}

// ISlashEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISlashEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlashEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISlashEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlashEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISlashEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlashEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISlashEscrowSession struct {
	Contract     *ISlashEscrow     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISlashEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISlashEscrowCallerSession struct {
	Contract *ISlashEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ISlashEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISlashEscrowTransactorSession struct {
	Contract     *ISlashEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ISlashEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISlashEscrowRaw struct {
	Contract *ISlashEscrow // Generic contract binding to access the raw methods on
}

// ISlashEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISlashEscrowCallerRaw struct {
	Contract *ISlashEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// ISlashEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISlashEscrowTransactorRaw struct {
	Contract *ISlashEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISlashEscrow creates a new instance of ISlashEscrow, bound to a specific deployed contract.
func NewISlashEscrow(address common.Address, backend bind.ContractBackend) (*ISlashEscrow, error) {
	contract, err := bindISlashEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISlashEscrow{ISlashEscrowCaller: ISlashEscrowCaller{contract: contract}, ISlashEscrowTransactor: ISlashEscrowTransactor{contract: contract}, ISlashEscrowFilterer: ISlashEscrowFilterer{contract: contract}}, nil
}

// NewISlashEscrowCaller creates a new read-only instance of ISlashEscrow, bound to a specific deployed contract.
func NewISlashEscrowCaller(address common.Address, caller bind.ContractCaller) (*ISlashEscrowCaller, error) {
	contract, err := bindISlashEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISlashEscrowCaller{contract: contract}, nil
}

// NewISlashEscrowTransactor creates a new write-only instance of ISlashEscrow, bound to a specific deployed contract.
func NewISlashEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*ISlashEscrowTransactor, error) {
	contract, err := bindISlashEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISlashEscrowTransactor{contract: contract}, nil
}

// NewISlashEscrowFilterer creates a new log filterer instance of ISlashEscrow, bound to a specific deployed contract.
func NewISlashEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*ISlashEscrowFilterer, error) {
	contract, err := bindISlashEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISlashEscrowFilterer{contract: contract}, nil
}

// bindISlashEscrow binds a generic wrapper to an already deployed contract.
func bindISlashEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISlashEscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlashEscrow *ISlashEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlashEscrow.Contract.ISlashEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlashEscrow *ISlashEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlashEscrow.Contract.ISlashEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlashEscrow *ISlashEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlashEscrow.Contract.ISlashEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlashEscrow *ISlashEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlashEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlashEscrow *ISlashEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlashEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlashEscrow *ISlashEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlashEscrow.Contract.contract.Transact(opts, method, params...)
}

// VerifyDeploymentParameters is a free data retrieval call binding the contract method 0x0d9e4ef1.
//
// Solidity: function verifyDeploymentParameters(address slashEscrowFactory, address slashEscrowImplementation, (address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_ISlashEscrow *ISlashEscrowCaller) VerifyDeploymentParameters(opts *bind.CallOpts, slashEscrowFactory common.Address, slashEscrowImplementation common.Address, operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	var out []interface{}
	err := _ISlashEscrow.contract.Call(opts, &out, "verifyDeploymentParameters", slashEscrowFactory, slashEscrowImplementation, operatorSet, slashId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyDeploymentParameters is a free data retrieval call binding the contract method 0x0d9e4ef1.
//
// Solidity: function verifyDeploymentParameters(address slashEscrowFactory, address slashEscrowImplementation, (address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_ISlashEscrow *ISlashEscrowSession) VerifyDeploymentParameters(slashEscrowFactory common.Address, slashEscrowImplementation common.Address, operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _ISlashEscrow.Contract.VerifyDeploymentParameters(&_ISlashEscrow.CallOpts, slashEscrowFactory, slashEscrowImplementation, operatorSet, slashId)
}

// VerifyDeploymentParameters is a free data retrieval call binding the contract method 0x0d9e4ef1.
//
// Solidity: function verifyDeploymentParameters(address slashEscrowFactory, address slashEscrowImplementation, (address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_ISlashEscrow *ISlashEscrowCallerSession) VerifyDeploymentParameters(slashEscrowFactory common.Address, slashEscrowImplementation common.Address, operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _ISlashEscrow.Contract.VerifyDeploymentParameters(&_ISlashEscrow.CallOpts, slashEscrowFactory, slashEscrowImplementation, operatorSet, slashId)
}

// ReleaseTokens is a paid mutator transaction binding the contract method 0xff491e65.
//
// Solidity: function releaseTokens(address slashEscrowFactory, address slashEscrowImplementation, (address,uint32) operatorSet, uint256 slashId, address recipient, address strategy) returns()
func (_ISlashEscrow *ISlashEscrowTransactor) ReleaseTokens(opts *bind.TransactOpts, slashEscrowFactory common.Address, slashEscrowImplementation common.Address, operatorSet OperatorSet, slashId *big.Int, recipient common.Address, strategy common.Address) (*types.Transaction, error) {
	return _ISlashEscrow.contract.Transact(opts, "releaseTokens", slashEscrowFactory, slashEscrowImplementation, operatorSet, slashId, recipient, strategy)
}

// ReleaseTokens is a paid mutator transaction binding the contract method 0xff491e65.
//
// Solidity: function releaseTokens(address slashEscrowFactory, address slashEscrowImplementation, (address,uint32) operatorSet, uint256 slashId, address recipient, address strategy) returns()
func (_ISlashEscrow *ISlashEscrowSession) ReleaseTokens(slashEscrowFactory common.Address, slashEscrowImplementation common.Address, operatorSet OperatorSet, slashId *big.Int, recipient common.Address, strategy common.Address) (*types.Transaction, error) {
	return _ISlashEscrow.Contract.ReleaseTokens(&_ISlashEscrow.TransactOpts, slashEscrowFactory, slashEscrowImplementation, operatorSet, slashId, recipient, strategy)
}

// ReleaseTokens is a paid mutator transaction binding the contract method 0xff491e65.
//
// Solidity: function releaseTokens(address slashEscrowFactory, address slashEscrowImplementation, (address,uint32) operatorSet, uint256 slashId, address recipient, address strategy) returns()
func (_ISlashEscrow *ISlashEscrowTransactorSession) ReleaseTokens(slashEscrowFactory common.Address, slashEscrowImplementation common.Address, operatorSet OperatorSet, slashId *big.Int, recipient common.Address, strategy common.Address) (*types.Transaction, error) {
	return _ISlashEscrow.Contract.ReleaseTokens(&_ISlashEscrow.TransactOpts, slashEscrowFactory, slashEscrowImplementation, operatorSet, slashId, recipient, strategy)
}
