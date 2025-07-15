// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IAVSTaskHook

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

// ITaskMailboxTypesTaskParams is an auto generated low-level Go binding around an user-defined struct.
type ITaskMailboxTypesTaskParams struct {
	RefundCollector     common.Address
	ExecutorOperatorSet OperatorSet
	Payload             []byte
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// IAVSTaskHookMetaData contains all meta data concerning the IAVSTaskHook contract.
var IAVSTaskHookMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"calculateTaskFee\",\"inputs\":[{\"name\":\"taskParams\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.TaskParams\",\"components\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handlePostTaskCreation\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"handlePostTaskResultSubmission\",\"inputs\":[{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatePreTaskCreation\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskParams\",\"type\":\"tuple\",\"internalType\":\"structITaskMailboxTypes.TaskParams\",\"components\":[{\"name\":\"refundCollector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executorOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatePreTaskResultSubmission\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taskHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"cert\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"}]",
}

// IAVSTaskHookABI is the input ABI used to generate the binding from.
// Deprecated: Use IAVSTaskHookMetaData.ABI instead.
var IAVSTaskHookABI = IAVSTaskHookMetaData.ABI

// IAVSTaskHook is an auto generated Go binding around an Ethereum contract.
type IAVSTaskHook struct {
	IAVSTaskHookCaller     // Read-only binding to the contract
	IAVSTaskHookTransactor // Write-only binding to the contract
	IAVSTaskHookFilterer   // Log filterer for contract events
}

// IAVSTaskHookCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAVSTaskHookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAVSTaskHookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAVSTaskHookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAVSTaskHookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAVSTaskHookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAVSTaskHookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAVSTaskHookSession struct {
	Contract     *IAVSTaskHook     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAVSTaskHookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAVSTaskHookCallerSession struct {
	Contract *IAVSTaskHookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IAVSTaskHookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAVSTaskHookTransactorSession struct {
	Contract     *IAVSTaskHookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IAVSTaskHookRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAVSTaskHookRaw struct {
	Contract *IAVSTaskHook // Generic contract binding to access the raw methods on
}

// IAVSTaskHookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAVSTaskHookCallerRaw struct {
	Contract *IAVSTaskHookCaller // Generic read-only contract binding to access the raw methods on
}

// IAVSTaskHookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAVSTaskHookTransactorRaw struct {
	Contract *IAVSTaskHookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAVSTaskHook creates a new instance of IAVSTaskHook, bound to a specific deployed contract.
func NewIAVSTaskHook(address common.Address, backend bind.ContractBackend) (*IAVSTaskHook, error) {
	contract, err := bindIAVSTaskHook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAVSTaskHook{IAVSTaskHookCaller: IAVSTaskHookCaller{contract: contract}, IAVSTaskHookTransactor: IAVSTaskHookTransactor{contract: contract}, IAVSTaskHookFilterer: IAVSTaskHookFilterer{contract: contract}}, nil
}

// NewIAVSTaskHookCaller creates a new read-only instance of IAVSTaskHook, bound to a specific deployed contract.
func NewIAVSTaskHookCaller(address common.Address, caller bind.ContractCaller) (*IAVSTaskHookCaller, error) {
	contract, err := bindIAVSTaskHook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAVSTaskHookCaller{contract: contract}, nil
}

// NewIAVSTaskHookTransactor creates a new write-only instance of IAVSTaskHook, bound to a specific deployed contract.
func NewIAVSTaskHookTransactor(address common.Address, transactor bind.ContractTransactor) (*IAVSTaskHookTransactor, error) {
	contract, err := bindIAVSTaskHook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAVSTaskHookTransactor{contract: contract}, nil
}

// NewIAVSTaskHookFilterer creates a new log filterer instance of IAVSTaskHook, bound to a specific deployed contract.
func NewIAVSTaskHookFilterer(address common.Address, filterer bind.ContractFilterer) (*IAVSTaskHookFilterer, error) {
	contract, err := bindIAVSTaskHook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAVSTaskHookFilterer{contract: contract}, nil
}

// bindIAVSTaskHook binds a generic wrapper to an already deployed contract.
func bindIAVSTaskHook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAVSTaskHookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAVSTaskHook *IAVSTaskHookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAVSTaskHook.Contract.IAVSTaskHookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAVSTaskHook *IAVSTaskHookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAVSTaskHook.Contract.IAVSTaskHookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAVSTaskHook *IAVSTaskHookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAVSTaskHook.Contract.IAVSTaskHookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAVSTaskHook *IAVSTaskHookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAVSTaskHook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAVSTaskHook *IAVSTaskHookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAVSTaskHook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAVSTaskHook *IAVSTaskHookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAVSTaskHook.Contract.contract.Transact(opts, method, params...)
}

// CalculateTaskFee is a free data retrieval call binding the contract method 0xe06cd27e.
//
// Solidity: function calculateTaskFee((address,(address,uint32),bytes) taskParams) view returns(uint96)
func (_IAVSTaskHook *IAVSTaskHookCaller) CalculateTaskFee(opts *bind.CallOpts, taskParams ITaskMailboxTypesTaskParams) (*big.Int, error) {
	var out []interface{}
	err := _IAVSTaskHook.contract.Call(opts, &out, "calculateTaskFee", taskParams)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTaskFee is a free data retrieval call binding the contract method 0xe06cd27e.
//
// Solidity: function calculateTaskFee((address,(address,uint32),bytes) taskParams) view returns(uint96)
func (_IAVSTaskHook *IAVSTaskHookSession) CalculateTaskFee(taskParams ITaskMailboxTypesTaskParams) (*big.Int, error) {
	return _IAVSTaskHook.Contract.CalculateTaskFee(&_IAVSTaskHook.CallOpts, taskParams)
}

// CalculateTaskFee is a free data retrieval call binding the contract method 0xe06cd27e.
//
// Solidity: function calculateTaskFee((address,(address,uint32),bytes) taskParams) view returns(uint96)
func (_IAVSTaskHook *IAVSTaskHookCallerSession) CalculateTaskFee(taskParams ITaskMailboxTypesTaskParams) (*big.Int, error) {
	return _IAVSTaskHook.Contract.CalculateTaskFee(&_IAVSTaskHook.CallOpts, taskParams)
}

// ValidatePreTaskCreation is a free data retrieval call binding the contract method 0x51fe3098.
//
// Solidity: function validatePreTaskCreation(address caller, (address,(address,uint32),bytes) taskParams) view returns()
func (_IAVSTaskHook *IAVSTaskHookCaller) ValidatePreTaskCreation(opts *bind.CallOpts, caller common.Address, taskParams ITaskMailboxTypesTaskParams) error {
	var out []interface{}
	err := _IAVSTaskHook.contract.Call(opts, &out, "validatePreTaskCreation", caller, taskParams)

	if err != nil {
		return err
	}

	return err

}

// ValidatePreTaskCreation is a free data retrieval call binding the contract method 0x51fe3098.
//
// Solidity: function validatePreTaskCreation(address caller, (address,(address,uint32),bytes) taskParams) view returns()
func (_IAVSTaskHook *IAVSTaskHookSession) ValidatePreTaskCreation(caller common.Address, taskParams ITaskMailboxTypesTaskParams) error {
	return _IAVSTaskHook.Contract.ValidatePreTaskCreation(&_IAVSTaskHook.CallOpts, caller, taskParams)
}

// ValidatePreTaskCreation is a free data retrieval call binding the contract method 0x51fe3098.
//
// Solidity: function validatePreTaskCreation(address caller, (address,(address,uint32),bytes) taskParams) view returns()
func (_IAVSTaskHook *IAVSTaskHookCallerSession) ValidatePreTaskCreation(caller common.Address, taskParams ITaskMailboxTypesTaskParams) error {
	return _IAVSTaskHook.Contract.ValidatePreTaskCreation(&_IAVSTaskHook.CallOpts, caller, taskParams)
}

// ValidatePreTaskResultSubmission is a free data retrieval call binding the contract method 0xba33565d.
//
// Solidity: function validatePreTaskResultSubmission(address caller, bytes32 taskHash, bytes cert, bytes result) view returns()
func (_IAVSTaskHook *IAVSTaskHookCaller) ValidatePreTaskResultSubmission(opts *bind.CallOpts, caller common.Address, taskHash [32]byte, cert []byte, result []byte) error {
	var out []interface{}
	err := _IAVSTaskHook.contract.Call(opts, &out, "validatePreTaskResultSubmission", caller, taskHash, cert, result)

	if err != nil {
		return err
	}

	return err

}

// ValidatePreTaskResultSubmission is a free data retrieval call binding the contract method 0xba33565d.
//
// Solidity: function validatePreTaskResultSubmission(address caller, bytes32 taskHash, bytes cert, bytes result) view returns()
func (_IAVSTaskHook *IAVSTaskHookSession) ValidatePreTaskResultSubmission(caller common.Address, taskHash [32]byte, cert []byte, result []byte) error {
	return _IAVSTaskHook.Contract.ValidatePreTaskResultSubmission(&_IAVSTaskHook.CallOpts, caller, taskHash, cert, result)
}

// ValidatePreTaskResultSubmission is a free data retrieval call binding the contract method 0xba33565d.
//
// Solidity: function validatePreTaskResultSubmission(address caller, bytes32 taskHash, bytes cert, bytes result) view returns()
func (_IAVSTaskHook *IAVSTaskHookCallerSession) ValidatePreTaskResultSubmission(caller common.Address, taskHash [32]byte, cert []byte, result []byte) error {
	return _IAVSTaskHook.Contract.ValidatePreTaskResultSubmission(&_IAVSTaskHook.CallOpts, caller, taskHash, cert, result)
}

// HandlePostTaskCreation is a paid mutator transaction binding the contract method 0x09c5c450.
//
// Solidity: function handlePostTaskCreation(bytes32 taskHash) returns()
func (_IAVSTaskHook *IAVSTaskHookTransactor) HandlePostTaskCreation(opts *bind.TransactOpts, taskHash [32]byte) (*types.Transaction, error) {
	return _IAVSTaskHook.contract.Transact(opts, "handlePostTaskCreation", taskHash)
}

// HandlePostTaskCreation is a paid mutator transaction binding the contract method 0x09c5c450.
//
// Solidity: function handlePostTaskCreation(bytes32 taskHash) returns()
func (_IAVSTaskHook *IAVSTaskHookSession) HandlePostTaskCreation(taskHash [32]byte) (*types.Transaction, error) {
	return _IAVSTaskHook.Contract.HandlePostTaskCreation(&_IAVSTaskHook.TransactOpts, taskHash)
}

// HandlePostTaskCreation is a paid mutator transaction binding the contract method 0x09c5c450.
//
// Solidity: function handlePostTaskCreation(bytes32 taskHash) returns()
func (_IAVSTaskHook *IAVSTaskHookTransactorSession) HandlePostTaskCreation(taskHash [32]byte) (*types.Transaction, error) {
	return _IAVSTaskHook.Contract.HandlePostTaskCreation(&_IAVSTaskHook.TransactOpts, taskHash)
}

// HandlePostTaskResultSubmission is a paid mutator transaction binding the contract method 0xdb6ecf67.
//
// Solidity: function handlePostTaskResultSubmission(bytes32 taskHash) returns()
func (_IAVSTaskHook *IAVSTaskHookTransactor) HandlePostTaskResultSubmission(opts *bind.TransactOpts, taskHash [32]byte) (*types.Transaction, error) {
	return _IAVSTaskHook.contract.Transact(opts, "handlePostTaskResultSubmission", taskHash)
}

// HandlePostTaskResultSubmission is a paid mutator transaction binding the contract method 0xdb6ecf67.
//
// Solidity: function handlePostTaskResultSubmission(bytes32 taskHash) returns()
func (_IAVSTaskHook *IAVSTaskHookSession) HandlePostTaskResultSubmission(taskHash [32]byte) (*types.Transaction, error) {
	return _IAVSTaskHook.Contract.HandlePostTaskResultSubmission(&_IAVSTaskHook.TransactOpts, taskHash)
}

// HandlePostTaskResultSubmission is a paid mutator transaction binding the contract method 0xdb6ecf67.
//
// Solidity: function handlePostTaskResultSubmission(bytes32 taskHash) returns()
func (_IAVSTaskHook *IAVSTaskHookTransactorSession) HandlePostTaskResultSubmission(taskHash [32]byte) (*types.Transaction, error) {
	return _IAVSTaskHook.Contract.HandlePostTaskResultSubmission(&_IAVSTaskHook.TransactOpts, taskHash)
}
