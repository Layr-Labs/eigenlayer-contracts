// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ISlasher

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

// ISlasherMiddlewareTimes is an auto generated low-level Go binding around an user-defined struct.
type ISlasherMiddlewareTimes struct {
	StalestUpdateBlock    uint32
	LatestServeUntilBlock uint32
}

// ISlasherMetaData contains all meta data concerning the ISlasher contract.
var ISlasherMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"canSlash\",\"inputs\":[{\"name\":\"toBeSlashed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slashingContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canWithdraw\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawalStartBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"middlewareTimesIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"contractCanSlashOperatorUntilBlock\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"serviceContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"toBeFrozen\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCorrectValueForInsertAfter\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"updateBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMiddlewareTimesIndexServeUntilBlock\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMiddlewareTimesIndexStalestUpdateBlock\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFrozen\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestUpdateBlock\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"serviceContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"middlewareTimesLength\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorToMiddlewareTimes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arrayIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structISlasher.MiddlewareTimes\",\"components\":[{\"name\":\"stalestUpdateBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"latestServeUntilBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorWhitelistedContractsLinkedListEntry\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorWhitelistedContractsLinkedListSize\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"optIntoSlashing\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recordFirstStakeUpdate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"serveUntilBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recordLastStakeUpdateAndRevokeSlashingAbility\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"serveUntilBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recordStakeUpdate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"updateBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"serveUntilBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"insertAfter\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resetFrozenStatus\",\"inputs\":[{\"name\":\"frozenAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FrozenStatusReset\",\"inputs\":[{\"name\":\"previouslySlashedAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MiddlewareTimesAdded\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"stalestUpdateBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"latestServeUntilBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorFrozen\",\"inputs\":[{\"name\":\"slashedOperator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slashingContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OptedIntoSlashing\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashingAbilityRevoked\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contractCanSlashOperatorUntilBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false}]",
}

// ISlasherABI is the input ABI used to generate the binding from.
// Deprecated: Use ISlasherMetaData.ABI instead.
var ISlasherABI = ISlasherMetaData.ABI

// ISlasher is an auto generated Go binding around an Ethereum contract.
type ISlasher struct {
	ISlasherCaller     // Read-only binding to the contract
	ISlasherTransactor // Write-only binding to the contract
	ISlasherFilterer   // Log filterer for contract events
}

// ISlasherCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISlasherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlasherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISlasherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlasherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISlasherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlasherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISlasherSession struct {
	Contract     *ISlasher         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISlasherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISlasherCallerSession struct {
	Contract *ISlasherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ISlasherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISlasherTransactorSession struct {
	Contract     *ISlasherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ISlasherRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISlasherRaw struct {
	Contract *ISlasher // Generic contract binding to access the raw methods on
}

// ISlasherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISlasherCallerRaw struct {
	Contract *ISlasherCaller // Generic read-only contract binding to access the raw methods on
}

// ISlasherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISlasherTransactorRaw struct {
	Contract *ISlasherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISlasher creates a new instance of ISlasher, bound to a specific deployed contract.
func NewISlasher(address common.Address, backend bind.ContractBackend) (*ISlasher, error) {
	contract, err := bindISlasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISlasher{ISlasherCaller: ISlasherCaller{contract: contract}, ISlasherTransactor: ISlasherTransactor{contract: contract}, ISlasherFilterer: ISlasherFilterer{contract: contract}}, nil
}

// NewISlasherCaller creates a new read-only instance of ISlasher, bound to a specific deployed contract.
func NewISlasherCaller(address common.Address, caller bind.ContractCaller) (*ISlasherCaller, error) {
	contract, err := bindISlasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISlasherCaller{contract: contract}, nil
}

// NewISlasherTransactor creates a new write-only instance of ISlasher, bound to a specific deployed contract.
func NewISlasherTransactor(address common.Address, transactor bind.ContractTransactor) (*ISlasherTransactor, error) {
	contract, err := bindISlasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISlasherTransactor{contract: contract}, nil
}

// NewISlasherFilterer creates a new log filterer instance of ISlasher, bound to a specific deployed contract.
func NewISlasherFilterer(address common.Address, filterer bind.ContractFilterer) (*ISlasherFilterer, error) {
	contract, err := bindISlasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISlasherFilterer{contract: contract}, nil
}

// bindISlasher binds a generic wrapper to an already deployed contract.
func bindISlasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISlasherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlasher *ISlasherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlasher.Contract.ISlasherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlasher *ISlasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlasher.Contract.ISlasherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlasher *ISlasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlasher.Contract.ISlasherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlasher *ISlasherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlasher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlasher *ISlasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlasher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlasher *ISlasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlasher.Contract.contract.Transact(opts, method, params...)
}

// CanSlash is a free data retrieval call binding the contract method 0xd98128c0.
//
// Solidity: function canSlash(address toBeSlashed, address slashingContract) view returns(bool)
func (_ISlasher *ISlasherCaller) CanSlash(opts *bind.CallOpts, toBeSlashed common.Address, slashingContract common.Address) (bool, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "canSlash", toBeSlashed, slashingContract)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanSlash is a free data retrieval call binding the contract method 0xd98128c0.
//
// Solidity: function canSlash(address toBeSlashed, address slashingContract) view returns(bool)
func (_ISlasher *ISlasherSession) CanSlash(toBeSlashed common.Address, slashingContract common.Address) (bool, error) {
	return _ISlasher.Contract.CanSlash(&_ISlasher.CallOpts, toBeSlashed, slashingContract)
}

// CanSlash is a free data retrieval call binding the contract method 0xd98128c0.
//
// Solidity: function canSlash(address toBeSlashed, address slashingContract) view returns(bool)
func (_ISlasher *ISlasherCallerSession) CanSlash(toBeSlashed common.Address, slashingContract common.Address) (bool, error) {
	return _ISlasher.Contract.CanSlash(&_ISlasher.CallOpts, toBeSlashed, slashingContract)
}

// ContractCanSlashOperatorUntilBlock is a free data retrieval call binding the contract method 0x6f0c2f74.
//
// Solidity: function contractCanSlashOperatorUntilBlock(address operator, address serviceContract) view returns(uint32)
func (_ISlasher *ISlasherCaller) ContractCanSlashOperatorUntilBlock(opts *bind.CallOpts, operator common.Address, serviceContract common.Address) (uint32, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "contractCanSlashOperatorUntilBlock", operator, serviceContract)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ContractCanSlashOperatorUntilBlock is a free data retrieval call binding the contract method 0x6f0c2f74.
//
// Solidity: function contractCanSlashOperatorUntilBlock(address operator, address serviceContract) view returns(uint32)
func (_ISlasher *ISlasherSession) ContractCanSlashOperatorUntilBlock(operator common.Address, serviceContract common.Address) (uint32, error) {
	return _ISlasher.Contract.ContractCanSlashOperatorUntilBlock(&_ISlasher.CallOpts, operator, serviceContract)
}

// ContractCanSlashOperatorUntilBlock is a free data retrieval call binding the contract method 0x6f0c2f74.
//
// Solidity: function contractCanSlashOperatorUntilBlock(address operator, address serviceContract) view returns(uint32)
func (_ISlasher *ISlasherCallerSession) ContractCanSlashOperatorUntilBlock(operator common.Address, serviceContract common.Address) (uint32, error) {
	return _ISlasher.Contract.ContractCanSlashOperatorUntilBlock(&_ISlasher.CallOpts, operator, serviceContract)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ISlasher *ISlasherCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ISlasher *ISlasherSession) Delegation() (common.Address, error) {
	return _ISlasher.Contract.Delegation(&_ISlasher.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ISlasher *ISlasherCallerSession) Delegation() (common.Address, error) {
	return _ISlasher.Contract.Delegation(&_ISlasher.CallOpts)
}

// GetCorrectValueForInsertAfter is a free data retrieval call binding the contract method 0x723e59c7.
//
// Solidity: function getCorrectValueForInsertAfter(address operator, uint32 updateBlock) view returns(uint256)
func (_ISlasher *ISlasherCaller) GetCorrectValueForInsertAfter(opts *bind.CallOpts, operator common.Address, updateBlock uint32) (*big.Int, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "getCorrectValueForInsertAfter", operator, updateBlock)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCorrectValueForInsertAfter is a free data retrieval call binding the contract method 0x723e59c7.
//
// Solidity: function getCorrectValueForInsertAfter(address operator, uint32 updateBlock) view returns(uint256)
func (_ISlasher *ISlasherSession) GetCorrectValueForInsertAfter(operator common.Address, updateBlock uint32) (*big.Int, error) {
	return _ISlasher.Contract.GetCorrectValueForInsertAfter(&_ISlasher.CallOpts, operator, updateBlock)
}

// GetCorrectValueForInsertAfter is a free data retrieval call binding the contract method 0x723e59c7.
//
// Solidity: function getCorrectValueForInsertAfter(address operator, uint32 updateBlock) view returns(uint256)
func (_ISlasher *ISlasherCallerSession) GetCorrectValueForInsertAfter(operator common.Address, updateBlock uint32) (*big.Int, error) {
	return _ISlasher.Contract.GetCorrectValueForInsertAfter(&_ISlasher.CallOpts, operator, updateBlock)
}

// GetMiddlewareTimesIndexServeUntilBlock is a free data retrieval call binding the contract method 0x7259a45c.
//
// Solidity: function getMiddlewareTimesIndexServeUntilBlock(address operator, uint32 index) view returns(uint32)
func (_ISlasher *ISlasherCaller) GetMiddlewareTimesIndexServeUntilBlock(opts *bind.CallOpts, operator common.Address, index uint32) (uint32, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "getMiddlewareTimesIndexServeUntilBlock", operator, index)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMiddlewareTimesIndexServeUntilBlock is a free data retrieval call binding the contract method 0x7259a45c.
//
// Solidity: function getMiddlewareTimesIndexServeUntilBlock(address operator, uint32 index) view returns(uint32)
func (_ISlasher *ISlasherSession) GetMiddlewareTimesIndexServeUntilBlock(operator common.Address, index uint32) (uint32, error) {
	return _ISlasher.Contract.GetMiddlewareTimesIndexServeUntilBlock(&_ISlasher.CallOpts, operator, index)
}

// GetMiddlewareTimesIndexServeUntilBlock is a free data retrieval call binding the contract method 0x7259a45c.
//
// Solidity: function getMiddlewareTimesIndexServeUntilBlock(address operator, uint32 index) view returns(uint32)
func (_ISlasher *ISlasherCallerSession) GetMiddlewareTimesIndexServeUntilBlock(operator common.Address, index uint32) (uint32, error) {
	return _ISlasher.Contract.GetMiddlewareTimesIndexServeUntilBlock(&_ISlasher.CallOpts, operator, index)
}

// GetMiddlewareTimesIndexStalestUpdateBlock is a free data retrieval call binding the contract method 0x1874e5ae.
//
// Solidity: function getMiddlewareTimesIndexStalestUpdateBlock(address operator, uint32 index) view returns(uint32)
func (_ISlasher *ISlasherCaller) GetMiddlewareTimesIndexStalestUpdateBlock(opts *bind.CallOpts, operator common.Address, index uint32) (uint32, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "getMiddlewareTimesIndexStalestUpdateBlock", operator, index)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMiddlewareTimesIndexStalestUpdateBlock is a free data retrieval call binding the contract method 0x1874e5ae.
//
// Solidity: function getMiddlewareTimesIndexStalestUpdateBlock(address operator, uint32 index) view returns(uint32)
func (_ISlasher *ISlasherSession) GetMiddlewareTimesIndexStalestUpdateBlock(operator common.Address, index uint32) (uint32, error) {
	return _ISlasher.Contract.GetMiddlewareTimesIndexStalestUpdateBlock(&_ISlasher.CallOpts, operator, index)
}

// GetMiddlewareTimesIndexStalestUpdateBlock is a free data retrieval call binding the contract method 0x1874e5ae.
//
// Solidity: function getMiddlewareTimesIndexStalestUpdateBlock(address operator, uint32 index) view returns(uint32)
func (_ISlasher *ISlasherCallerSession) GetMiddlewareTimesIndexStalestUpdateBlock(operator common.Address, index uint32) (uint32, error) {
	return _ISlasher.Contract.GetMiddlewareTimesIndexStalestUpdateBlock(&_ISlasher.CallOpts, operator, index)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address staker) view returns(bool)
func (_ISlasher *ISlasherCaller) IsFrozen(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "isFrozen", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address staker) view returns(bool)
func (_ISlasher *ISlasherSession) IsFrozen(staker common.Address) (bool, error) {
	return _ISlasher.Contract.IsFrozen(&_ISlasher.CallOpts, staker)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address staker) view returns(bool)
func (_ISlasher *ISlasherCallerSession) IsFrozen(staker common.Address) (bool, error) {
	return _ISlasher.Contract.IsFrozen(&_ISlasher.CallOpts, staker)
}

// LatestUpdateBlock is a free data retrieval call binding the contract method 0xda16e29b.
//
// Solidity: function latestUpdateBlock(address operator, address serviceContract) view returns(uint32)
func (_ISlasher *ISlasherCaller) LatestUpdateBlock(opts *bind.CallOpts, operator common.Address, serviceContract common.Address) (uint32, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "latestUpdateBlock", operator, serviceContract)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestUpdateBlock is a free data retrieval call binding the contract method 0xda16e29b.
//
// Solidity: function latestUpdateBlock(address operator, address serviceContract) view returns(uint32)
func (_ISlasher *ISlasherSession) LatestUpdateBlock(operator common.Address, serviceContract common.Address) (uint32, error) {
	return _ISlasher.Contract.LatestUpdateBlock(&_ISlasher.CallOpts, operator, serviceContract)
}

// LatestUpdateBlock is a free data retrieval call binding the contract method 0xda16e29b.
//
// Solidity: function latestUpdateBlock(address operator, address serviceContract) view returns(uint32)
func (_ISlasher *ISlasherCallerSession) LatestUpdateBlock(operator common.Address, serviceContract common.Address) (uint32, error) {
	return _ISlasher.Contract.LatestUpdateBlock(&_ISlasher.CallOpts, operator, serviceContract)
}

// MiddlewareTimesLength is a free data retrieval call binding the contract method 0xa49db732.
//
// Solidity: function middlewareTimesLength(address operator) view returns(uint256)
func (_ISlasher *ISlasherCaller) MiddlewareTimesLength(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "middlewareTimesLength", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiddlewareTimesLength is a free data retrieval call binding the contract method 0xa49db732.
//
// Solidity: function middlewareTimesLength(address operator) view returns(uint256)
func (_ISlasher *ISlasherSession) MiddlewareTimesLength(operator common.Address) (*big.Int, error) {
	return _ISlasher.Contract.MiddlewareTimesLength(&_ISlasher.CallOpts, operator)
}

// MiddlewareTimesLength is a free data retrieval call binding the contract method 0xa49db732.
//
// Solidity: function middlewareTimesLength(address operator) view returns(uint256)
func (_ISlasher *ISlasherCallerSession) MiddlewareTimesLength(operator common.Address) (*big.Int, error) {
	return _ISlasher.Contract.MiddlewareTimesLength(&_ISlasher.CallOpts, operator)
}

// OperatorToMiddlewareTimes is a free data retrieval call binding the contract method 0x282670fc.
//
// Solidity: function operatorToMiddlewareTimes(address operator, uint256 arrayIndex) view returns((uint32,uint32))
func (_ISlasher *ISlasherCaller) OperatorToMiddlewareTimes(opts *bind.CallOpts, operator common.Address, arrayIndex *big.Int) (ISlasherMiddlewareTimes, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "operatorToMiddlewareTimes", operator, arrayIndex)

	if err != nil {
		return *new(ISlasherMiddlewareTimes), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlasherMiddlewareTimes)).(*ISlasherMiddlewareTimes)

	return out0, err

}

// OperatorToMiddlewareTimes is a free data retrieval call binding the contract method 0x282670fc.
//
// Solidity: function operatorToMiddlewareTimes(address operator, uint256 arrayIndex) view returns((uint32,uint32))
func (_ISlasher *ISlasherSession) OperatorToMiddlewareTimes(operator common.Address, arrayIndex *big.Int) (ISlasherMiddlewareTimes, error) {
	return _ISlasher.Contract.OperatorToMiddlewareTimes(&_ISlasher.CallOpts, operator, arrayIndex)
}

// OperatorToMiddlewareTimes is a free data retrieval call binding the contract method 0x282670fc.
//
// Solidity: function operatorToMiddlewareTimes(address operator, uint256 arrayIndex) view returns((uint32,uint32))
func (_ISlasher *ISlasherCallerSession) OperatorToMiddlewareTimes(operator common.Address, arrayIndex *big.Int) (ISlasherMiddlewareTimes, error) {
	return _ISlasher.Contract.OperatorToMiddlewareTimes(&_ISlasher.CallOpts, operator, arrayIndex)
}

// OperatorWhitelistedContractsLinkedListEntry is a free data retrieval call binding the contract method 0x855fcc4a.
//
// Solidity: function operatorWhitelistedContractsLinkedListEntry(address operator, address node) view returns(bool, uint256, uint256)
func (_ISlasher *ISlasherCaller) OperatorWhitelistedContractsLinkedListEntry(opts *bind.CallOpts, operator common.Address, node common.Address) (bool, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "operatorWhitelistedContractsLinkedListEntry", operator, node)

	if err != nil {
		return *new(bool), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// OperatorWhitelistedContractsLinkedListEntry is a free data retrieval call binding the contract method 0x855fcc4a.
//
// Solidity: function operatorWhitelistedContractsLinkedListEntry(address operator, address node) view returns(bool, uint256, uint256)
func (_ISlasher *ISlasherSession) OperatorWhitelistedContractsLinkedListEntry(operator common.Address, node common.Address) (bool, *big.Int, *big.Int, error) {
	return _ISlasher.Contract.OperatorWhitelistedContractsLinkedListEntry(&_ISlasher.CallOpts, operator, node)
}

// OperatorWhitelistedContractsLinkedListEntry is a free data retrieval call binding the contract method 0x855fcc4a.
//
// Solidity: function operatorWhitelistedContractsLinkedListEntry(address operator, address node) view returns(bool, uint256, uint256)
func (_ISlasher *ISlasherCallerSession) OperatorWhitelistedContractsLinkedListEntry(operator common.Address, node common.Address) (bool, *big.Int, *big.Int, error) {
	return _ISlasher.Contract.OperatorWhitelistedContractsLinkedListEntry(&_ISlasher.CallOpts, operator, node)
}

// OperatorWhitelistedContractsLinkedListSize is a free data retrieval call binding the contract method 0xe921d4fa.
//
// Solidity: function operatorWhitelistedContractsLinkedListSize(address operator) view returns(uint256)
func (_ISlasher *ISlasherCaller) OperatorWhitelistedContractsLinkedListSize(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "operatorWhitelistedContractsLinkedListSize", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorWhitelistedContractsLinkedListSize is a free data retrieval call binding the contract method 0xe921d4fa.
//
// Solidity: function operatorWhitelistedContractsLinkedListSize(address operator) view returns(uint256)
func (_ISlasher *ISlasherSession) OperatorWhitelistedContractsLinkedListSize(operator common.Address) (*big.Int, error) {
	return _ISlasher.Contract.OperatorWhitelistedContractsLinkedListSize(&_ISlasher.CallOpts, operator)
}

// OperatorWhitelistedContractsLinkedListSize is a free data retrieval call binding the contract method 0xe921d4fa.
//
// Solidity: function operatorWhitelistedContractsLinkedListSize(address operator) view returns(uint256)
func (_ISlasher *ISlasherCallerSession) OperatorWhitelistedContractsLinkedListSize(operator common.Address) (*big.Int, error) {
	return _ISlasher.Contract.OperatorWhitelistedContractsLinkedListSize(&_ISlasher.CallOpts, operator)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ISlasher *ISlasherCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ISlasher *ISlasherSession) StrategyManager() (common.Address, error) {
	return _ISlasher.Contract.StrategyManager(&_ISlasher.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ISlasher *ISlasherCallerSession) StrategyManager() (common.Address, error) {
	return _ISlasher.Contract.StrategyManager(&_ISlasher.CallOpts)
}

// CanWithdraw is a paid mutator transaction binding the contract method 0x8105e043.
//
// Solidity: function canWithdraw(address operator, uint32 withdrawalStartBlock, uint256 middlewareTimesIndex) returns(bool)
func (_ISlasher *ISlasherTransactor) CanWithdraw(opts *bind.TransactOpts, operator common.Address, withdrawalStartBlock uint32, middlewareTimesIndex *big.Int) (*types.Transaction, error) {
	return _ISlasher.contract.Transact(opts, "canWithdraw", operator, withdrawalStartBlock, middlewareTimesIndex)
}

// CanWithdraw is a paid mutator transaction binding the contract method 0x8105e043.
//
// Solidity: function canWithdraw(address operator, uint32 withdrawalStartBlock, uint256 middlewareTimesIndex) returns(bool)
func (_ISlasher *ISlasherSession) CanWithdraw(operator common.Address, withdrawalStartBlock uint32, middlewareTimesIndex *big.Int) (*types.Transaction, error) {
	return _ISlasher.Contract.CanWithdraw(&_ISlasher.TransactOpts, operator, withdrawalStartBlock, middlewareTimesIndex)
}

// CanWithdraw is a paid mutator transaction binding the contract method 0x8105e043.
//
// Solidity: function canWithdraw(address operator, uint32 withdrawalStartBlock, uint256 middlewareTimesIndex) returns(bool)
func (_ISlasher *ISlasherTransactorSession) CanWithdraw(operator common.Address, withdrawalStartBlock uint32, middlewareTimesIndex *big.Int) (*types.Transaction, error) {
	return _ISlasher.Contract.CanWithdraw(&_ISlasher.TransactOpts, operator, withdrawalStartBlock, middlewareTimesIndex)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address toBeFrozen) returns()
func (_ISlasher *ISlasherTransactor) FreezeOperator(opts *bind.TransactOpts, toBeFrozen common.Address) (*types.Transaction, error) {
	return _ISlasher.contract.Transact(opts, "freezeOperator", toBeFrozen)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address toBeFrozen) returns()
func (_ISlasher *ISlasherSession) FreezeOperator(toBeFrozen common.Address) (*types.Transaction, error) {
	return _ISlasher.Contract.FreezeOperator(&_ISlasher.TransactOpts, toBeFrozen)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address toBeFrozen) returns()
func (_ISlasher *ISlasherTransactorSession) FreezeOperator(toBeFrozen common.Address) (*types.Transaction, error) {
	return _ISlasher.Contract.FreezeOperator(&_ISlasher.TransactOpts, toBeFrozen)
}

// OptIntoSlashing is a paid mutator transaction binding the contract method 0xf73b7519.
//
// Solidity: function optIntoSlashing(address contractAddress) returns()
func (_ISlasher *ISlasherTransactor) OptIntoSlashing(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ISlasher.contract.Transact(opts, "optIntoSlashing", contractAddress)
}

// OptIntoSlashing is a paid mutator transaction binding the contract method 0xf73b7519.
//
// Solidity: function optIntoSlashing(address contractAddress) returns()
func (_ISlasher *ISlasherSession) OptIntoSlashing(contractAddress common.Address) (*types.Transaction, error) {
	return _ISlasher.Contract.OptIntoSlashing(&_ISlasher.TransactOpts, contractAddress)
}

// OptIntoSlashing is a paid mutator transaction binding the contract method 0xf73b7519.
//
// Solidity: function optIntoSlashing(address contractAddress) returns()
func (_ISlasher *ISlasherTransactorSession) OptIntoSlashing(contractAddress common.Address) (*types.Transaction, error) {
	return _ISlasher.Contract.OptIntoSlashing(&_ISlasher.TransactOpts, contractAddress)
}

// RecordFirstStakeUpdate is a paid mutator transaction binding the contract method 0x175d3205.
//
// Solidity: function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock) returns()
func (_ISlasher *ISlasherTransactor) RecordFirstStakeUpdate(opts *bind.TransactOpts, operator common.Address, serveUntilBlock uint32) (*types.Transaction, error) {
	return _ISlasher.contract.Transact(opts, "recordFirstStakeUpdate", operator, serveUntilBlock)
}

// RecordFirstStakeUpdate is a paid mutator transaction binding the contract method 0x175d3205.
//
// Solidity: function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock) returns()
func (_ISlasher *ISlasherSession) RecordFirstStakeUpdate(operator common.Address, serveUntilBlock uint32) (*types.Transaction, error) {
	return _ISlasher.Contract.RecordFirstStakeUpdate(&_ISlasher.TransactOpts, operator, serveUntilBlock)
}

// RecordFirstStakeUpdate is a paid mutator transaction binding the contract method 0x175d3205.
//
// Solidity: function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock) returns()
func (_ISlasher *ISlasherTransactorSession) RecordFirstStakeUpdate(operator common.Address, serveUntilBlock uint32) (*types.Transaction, error) {
	return _ISlasher.Contract.RecordFirstStakeUpdate(&_ISlasher.TransactOpts, operator, serveUntilBlock)
}

// RecordLastStakeUpdateAndRevokeSlashingAbility is a paid mutator transaction binding the contract method 0x0ffabbce.
//
// Solidity: function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) returns()
func (_ISlasher *ISlasherTransactor) RecordLastStakeUpdateAndRevokeSlashingAbility(opts *bind.TransactOpts, operator common.Address, serveUntilBlock uint32) (*types.Transaction, error) {
	return _ISlasher.contract.Transact(opts, "recordLastStakeUpdateAndRevokeSlashingAbility", operator, serveUntilBlock)
}

// RecordLastStakeUpdateAndRevokeSlashingAbility is a paid mutator transaction binding the contract method 0x0ffabbce.
//
// Solidity: function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) returns()
func (_ISlasher *ISlasherSession) RecordLastStakeUpdateAndRevokeSlashingAbility(operator common.Address, serveUntilBlock uint32) (*types.Transaction, error) {
	return _ISlasher.Contract.RecordLastStakeUpdateAndRevokeSlashingAbility(&_ISlasher.TransactOpts, operator, serveUntilBlock)
}

// RecordLastStakeUpdateAndRevokeSlashingAbility is a paid mutator transaction binding the contract method 0x0ffabbce.
//
// Solidity: function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) returns()
func (_ISlasher *ISlasherTransactorSession) RecordLastStakeUpdateAndRevokeSlashingAbility(operator common.Address, serveUntilBlock uint32) (*types.Transaction, error) {
	return _ISlasher.Contract.RecordLastStakeUpdateAndRevokeSlashingAbility(&_ISlasher.TransactOpts, operator, serveUntilBlock)
}

// RecordStakeUpdate is a paid mutator transaction binding the contract method 0xc747075b.
//
// Solidity: function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 insertAfter) returns()
func (_ISlasher *ISlasherTransactor) RecordStakeUpdate(opts *bind.TransactOpts, operator common.Address, updateBlock uint32, serveUntilBlock uint32, insertAfter *big.Int) (*types.Transaction, error) {
	return _ISlasher.contract.Transact(opts, "recordStakeUpdate", operator, updateBlock, serveUntilBlock, insertAfter)
}

// RecordStakeUpdate is a paid mutator transaction binding the contract method 0xc747075b.
//
// Solidity: function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 insertAfter) returns()
func (_ISlasher *ISlasherSession) RecordStakeUpdate(operator common.Address, updateBlock uint32, serveUntilBlock uint32, insertAfter *big.Int) (*types.Transaction, error) {
	return _ISlasher.Contract.RecordStakeUpdate(&_ISlasher.TransactOpts, operator, updateBlock, serveUntilBlock, insertAfter)
}

// RecordStakeUpdate is a paid mutator transaction binding the contract method 0xc747075b.
//
// Solidity: function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 insertAfter) returns()
func (_ISlasher *ISlasherTransactorSession) RecordStakeUpdate(operator common.Address, updateBlock uint32, serveUntilBlock uint32, insertAfter *big.Int) (*types.Transaction, error) {
	return _ISlasher.Contract.RecordStakeUpdate(&_ISlasher.TransactOpts, operator, updateBlock, serveUntilBlock, insertAfter)
}

// ResetFrozenStatus is a paid mutator transaction binding the contract method 0x7cf72bba.
//
// Solidity: function resetFrozenStatus(address[] frozenAddresses) returns()
func (_ISlasher *ISlasherTransactor) ResetFrozenStatus(opts *bind.TransactOpts, frozenAddresses []common.Address) (*types.Transaction, error) {
	return _ISlasher.contract.Transact(opts, "resetFrozenStatus", frozenAddresses)
}

// ResetFrozenStatus is a paid mutator transaction binding the contract method 0x7cf72bba.
//
// Solidity: function resetFrozenStatus(address[] frozenAddresses) returns()
func (_ISlasher *ISlasherSession) ResetFrozenStatus(frozenAddresses []common.Address) (*types.Transaction, error) {
	return _ISlasher.Contract.ResetFrozenStatus(&_ISlasher.TransactOpts, frozenAddresses)
}

// ResetFrozenStatus is a paid mutator transaction binding the contract method 0x7cf72bba.
//
// Solidity: function resetFrozenStatus(address[] frozenAddresses) returns()
func (_ISlasher *ISlasherTransactorSession) ResetFrozenStatus(frozenAddresses []common.Address) (*types.Transaction, error) {
	return _ISlasher.Contract.ResetFrozenStatus(&_ISlasher.TransactOpts, frozenAddresses)
}

// ISlasherFrozenStatusResetIterator is returned from FilterFrozenStatusReset and is used to iterate over the raw logs and unpacked data for FrozenStatusReset events raised by the ISlasher contract.
type ISlasherFrozenStatusResetIterator struct {
	Event *ISlasherFrozenStatusReset // Event containing the contract specifics and raw log

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
func (it *ISlasherFrozenStatusResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlasherFrozenStatusReset)
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
		it.Event = new(ISlasherFrozenStatusReset)
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
func (it *ISlasherFrozenStatusResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlasherFrozenStatusResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlasherFrozenStatusReset represents a FrozenStatusReset event raised by the ISlasher contract.
type ISlasherFrozenStatusReset struct {
	PreviouslySlashedAddress common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterFrozenStatusReset is a free log retrieval operation binding the contract event 0xd4cef0af27800d466fcacd85779857378b85cb61569005ff1464fa6e5ced69d8.
//
// Solidity: event FrozenStatusReset(address indexed previouslySlashedAddress)
func (_ISlasher *ISlasherFilterer) FilterFrozenStatusReset(opts *bind.FilterOpts, previouslySlashedAddress []common.Address) (*ISlasherFrozenStatusResetIterator, error) {

	var previouslySlashedAddressRule []interface{}
	for _, previouslySlashedAddressItem := range previouslySlashedAddress {
		previouslySlashedAddressRule = append(previouslySlashedAddressRule, previouslySlashedAddressItem)
	}

	logs, sub, err := _ISlasher.contract.FilterLogs(opts, "FrozenStatusReset", previouslySlashedAddressRule)
	if err != nil {
		return nil, err
	}
	return &ISlasherFrozenStatusResetIterator{contract: _ISlasher.contract, event: "FrozenStatusReset", logs: logs, sub: sub}, nil
}

// WatchFrozenStatusReset is a free log subscription operation binding the contract event 0xd4cef0af27800d466fcacd85779857378b85cb61569005ff1464fa6e5ced69d8.
//
// Solidity: event FrozenStatusReset(address indexed previouslySlashedAddress)
func (_ISlasher *ISlasherFilterer) WatchFrozenStatusReset(opts *bind.WatchOpts, sink chan<- *ISlasherFrozenStatusReset, previouslySlashedAddress []common.Address) (event.Subscription, error) {

	var previouslySlashedAddressRule []interface{}
	for _, previouslySlashedAddressItem := range previouslySlashedAddress {
		previouslySlashedAddressRule = append(previouslySlashedAddressRule, previouslySlashedAddressItem)
	}

	logs, sub, err := _ISlasher.contract.WatchLogs(opts, "FrozenStatusReset", previouslySlashedAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlasherFrozenStatusReset)
				if err := _ISlasher.contract.UnpackLog(event, "FrozenStatusReset", log); err != nil {
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

// ParseFrozenStatusReset is a log parse operation binding the contract event 0xd4cef0af27800d466fcacd85779857378b85cb61569005ff1464fa6e5ced69d8.
//
// Solidity: event FrozenStatusReset(address indexed previouslySlashedAddress)
func (_ISlasher *ISlasherFilterer) ParseFrozenStatusReset(log types.Log) (*ISlasherFrozenStatusReset, error) {
	event := new(ISlasherFrozenStatusReset)
	if err := _ISlasher.contract.UnpackLog(event, "FrozenStatusReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlasherMiddlewareTimesAddedIterator is returned from FilterMiddlewareTimesAdded and is used to iterate over the raw logs and unpacked data for MiddlewareTimesAdded events raised by the ISlasher contract.
type ISlasherMiddlewareTimesAddedIterator struct {
	Event *ISlasherMiddlewareTimesAdded // Event containing the contract specifics and raw log

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
func (it *ISlasherMiddlewareTimesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlasherMiddlewareTimesAdded)
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
		it.Event = new(ISlasherMiddlewareTimesAdded)
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
func (it *ISlasherMiddlewareTimesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlasherMiddlewareTimesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlasherMiddlewareTimesAdded represents a MiddlewareTimesAdded event raised by the ISlasher contract.
type ISlasherMiddlewareTimesAdded struct {
	Operator              common.Address
	Index                 *big.Int
	StalestUpdateBlock    uint32
	LatestServeUntilBlock uint32
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMiddlewareTimesAdded is a free log retrieval operation binding the contract event 0x1b62ba64c72d01e41a2b8c46e6aeeff728ef3a4438cf1cac3d92ee12189d5649.
//
// Solidity: event MiddlewareTimesAdded(address operator, uint256 index, uint32 stalestUpdateBlock, uint32 latestServeUntilBlock)
func (_ISlasher *ISlasherFilterer) FilterMiddlewareTimesAdded(opts *bind.FilterOpts) (*ISlasherMiddlewareTimesAddedIterator, error) {

	logs, sub, err := _ISlasher.contract.FilterLogs(opts, "MiddlewareTimesAdded")
	if err != nil {
		return nil, err
	}
	return &ISlasherMiddlewareTimesAddedIterator{contract: _ISlasher.contract, event: "MiddlewareTimesAdded", logs: logs, sub: sub}, nil
}

// WatchMiddlewareTimesAdded is a free log subscription operation binding the contract event 0x1b62ba64c72d01e41a2b8c46e6aeeff728ef3a4438cf1cac3d92ee12189d5649.
//
// Solidity: event MiddlewareTimesAdded(address operator, uint256 index, uint32 stalestUpdateBlock, uint32 latestServeUntilBlock)
func (_ISlasher *ISlasherFilterer) WatchMiddlewareTimesAdded(opts *bind.WatchOpts, sink chan<- *ISlasherMiddlewareTimesAdded) (event.Subscription, error) {

	logs, sub, err := _ISlasher.contract.WatchLogs(opts, "MiddlewareTimesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlasherMiddlewareTimesAdded)
				if err := _ISlasher.contract.UnpackLog(event, "MiddlewareTimesAdded", log); err != nil {
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

// ParseMiddlewareTimesAdded is a log parse operation binding the contract event 0x1b62ba64c72d01e41a2b8c46e6aeeff728ef3a4438cf1cac3d92ee12189d5649.
//
// Solidity: event MiddlewareTimesAdded(address operator, uint256 index, uint32 stalestUpdateBlock, uint32 latestServeUntilBlock)
func (_ISlasher *ISlasherFilterer) ParseMiddlewareTimesAdded(log types.Log) (*ISlasherMiddlewareTimesAdded, error) {
	event := new(ISlasherMiddlewareTimesAdded)
	if err := _ISlasher.contract.UnpackLog(event, "MiddlewareTimesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlasherOperatorFrozenIterator is returned from FilterOperatorFrozen and is used to iterate over the raw logs and unpacked data for OperatorFrozen events raised by the ISlasher contract.
type ISlasherOperatorFrozenIterator struct {
	Event *ISlasherOperatorFrozen // Event containing the contract specifics and raw log

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
func (it *ISlasherOperatorFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlasherOperatorFrozen)
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
		it.Event = new(ISlasherOperatorFrozen)
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
func (it *ISlasherOperatorFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlasherOperatorFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlasherOperatorFrozen represents a OperatorFrozen event raised by the ISlasher contract.
type ISlasherOperatorFrozen struct {
	SlashedOperator  common.Address
	SlashingContract common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOperatorFrozen is a free log retrieval operation binding the contract event 0x444a84f512816ae7be8ed8a66aa88e362eb54d0988e83acc9d81746622b3ba51.
//
// Solidity: event OperatorFrozen(address indexed slashedOperator, address indexed slashingContract)
func (_ISlasher *ISlasherFilterer) FilterOperatorFrozen(opts *bind.FilterOpts, slashedOperator []common.Address, slashingContract []common.Address) (*ISlasherOperatorFrozenIterator, error) {

	var slashedOperatorRule []interface{}
	for _, slashedOperatorItem := range slashedOperator {
		slashedOperatorRule = append(slashedOperatorRule, slashedOperatorItem)
	}
	var slashingContractRule []interface{}
	for _, slashingContractItem := range slashingContract {
		slashingContractRule = append(slashingContractRule, slashingContractItem)
	}

	logs, sub, err := _ISlasher.contract.FilterLogs(opts, "OperatorFrozen", slashedOperatorRule, slashingContractRule)
	if err != nil {
		return nil, err
	}
	return &ISlasherOperatorFrozenIterator{contract: _ISlasher.contract, event: "OperatorFrozen", logs: logs, sub: sub}, nil
}

// WatchOperatorFrozen is a free log subscription operation binding the contract event 0x444a84f512816ae7be8ed8a66aa88e362eb54d0988e83acc9d81746622b3ba51.
//
// Solidity: event OperatorFrozen(address indexed slashedOperator, address indexed slashingContract)
func (_ISlasher *ISlasherFilterer) WatchOperatorFrozen(opts *bind.WatchOpts, sink chan<- *ISlasherOperatorFrozen, slashedOperator []common.Address, slashingContract []common.Address) (event.Subscription, error) {

	var slashedOperatorRule []interface{}
	for _, slashedOperatorItem := range slashedOperator {
		slashedOperatorRule = append(slashedOperatorRule, slashedOperatorItem)
	}
	var slashingContractRule []interface{}
	for _, slashingContractItem := range slashingContract {
		slashingContractRule = append(slashingContractRule, slashingContractItem)
	}

	logs, sub, err := _ISlasher.contract.WatchLogs(opts, "OperatorFrozen", slashedOperatorRule, slashingContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlasherOperatorFrozen)
				if err := _ISlasher.contract.UnpackLog(event, "OperatorFrozen", log); err != nil {
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

// ParseOperatorFrozen is a log parse operation binding the contract event 0x444a84f512816ae7be8ed8a66aa88e362eb54d0988e83acc9d81746622b3ba51.
//
// Solidity: event OperatorFrozen(address indexed slashedOperator, address indexed slashingContract)
func (_ISlasher *ISlasherFilterer) ParseOperatorFrozen(log types.Log) (*ISlasherOperatorFrozen, error) {
	event := new(ISlasherOperatorFrozen)
	if err := _ISlasher.contract.UnpackLog(event, "OperatorFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlasherOptedIntoSlashingIterator is returned from FilterOptedIntoSlashing and is used to iterate over the raw logs and unpacked data for OptedIntoSlashing events raised by the ISlasher contract.
type ISlasherOptedIntoSlashingIterator struct {
	Event *ISlasherOptedIntoSlashing // Event containing the contract specifics and raw log

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
func (it *ISlasherOptedIntoSlashingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlasherOptedIntoSlashing)
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
		it.Event = new(ISlasherOptedIntoSlashing)
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
func (it *ISlasherOptedIntoSlashingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlasherOptedIntoSlashingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlasherOptedIntoSlashing represents a OptedIntoSlashing event raised by the ISlasher contract.
type ISlasherOptedIntoSlashing struct {
	Operator        common.Address
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOptedIntoSlashing is a free log retrieval operation binding the contract event 0xefa9fb38e813d53c15edf501e03852843a3fed691960523391d71a092b3627d8.
//
// Solidity: event OptedIntoSlashing(address indexed operator, address indexed contractAddress)
func (_ISlasher *ISlasherFilterer) FilterOptedIntoSlashing(opts *bind.FilterOpts, operator []common.Address, contractAddress []common.Address) (*ISlasherOptedIntoSlashingIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ISlasher.contract.FilterLogs(opts, "OptedIntoSlashing", operatorRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &ISlasherOptedIntoSlashingIterator{contract: _ISlasher.contract, event: "OptedIntoSlashing", logs: logs, sub: sub}, nil
}

// WatchOptedIntoSlashing is a free log subscription operation binding the contract event 0xefa9fb38e813d53c15edf501e03852843a3fed691960523391d71a092b3627d8.
//
// Solidity: event OptedIntoSlashing(address indexed operator, address indexed contractAddress)
func (_ISlasher *ISlasherFilterer) WatchOptedIntoSlashing(opts *bind.WatchOpts, sink chan<- *ISlasherOptedIntoSlashing, operator []common.Address, contractAddress []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ISlasher.contract.WatchLogs(opts, "OptedIntoSlashing", operatorRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlasherOptedIntoSlashing)
				if err := _ISlasher.contract.UnpackLog(event, "OptedIntoSlashing", log); err != nil {
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

// ParseOptedIntoSlashing is a log parse operation binding the contract event 0xefa9fb38e813d53c15edf501e03852843a3fed691960523391d71a092b3627d8.
//
// Solidity: event OptedIntoSlashing(address indexed operator, address indexed contractAddress)
func (_ISlasher *ISlasherFilterer) ParseOptedIntoSlashing(log types.Log) (*ISlasherOptedIntoSlashing, error) {
	event := new(ISlasherOptedIntoSlashing)
	if err := _ISlasher.contract.UnpackLog(event, "OptedIntoSlashing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlasherSlashingAbilityRevokedIterator is returned from FilterSlashingAbilityRevoked and is used to iterate over the raw logs and unpacked data for SlashingAbilityRevoked events raised by the ISlasher contract.
type ISlasherSlashingAbilityRevokedIterator struct {
	Event *ISlasherSlashingAbilityRevoked // Event containing the contract specifics and raw log

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
func (it *ISlasherSlashingAbilityRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlasherSlashingAbilityRevoked)
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
		it.Event = new(ISlasherSlashingAbilityRevoked)
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
func (it *ISlasherSlashingAbilityRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlasherSlashingAbilityRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlasherSlashingAbilityRevoked represents a SlashingAbilityRevoked event raised by the ISlasher contract.
type ISlasherSlashingAbilityRevoked struct {
	Operator                           common.Address
	ContractAddress                    common.Address
	ContractCanSlashOperatorUntilBlock uint32
	Raw                                types.Log // Blockchain specific contextual infos
}

// FilterSlashingAbilityRevoked is a free log retrieval operation binding the contract event 0x9aa1b1391f35c672ed1f3b7ece632f4513e618366bef7a2f67b7c6bc1f2d2b14.
//
// Solidity: event SlashingAbilityRevoked(address indexed operator, address indexed contractAddress, uint32 contractCanSlashOperatorUntilBlock)
func (_ISlasher *ISlasherFilterer) FilterSlashingAbilityRevoked(opts *bind.FilterOpts, operator []common.Address, contractAddress []common.Address) (*ISlasherSlashingAbilityRevokedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ISlasher.contract.FilterLogs(opts, "SlashingAbilityRevoked", operatorRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &ISlasherSlashingAbilityRevokedIterator{contract: _ISlasher.contract, event: "SlashingAbilityRevoked", logs: logs, sub: sub}, nil
}

// WatchSlashingAbilityRevoked is a free log subscription operation binding the contract event 0x9aa1b1391f35c672ed1f3b7ece632f4513e618366bef7a2f67b7c6bc1f2d2b14.
//
// Solidity: event SlashingAbilityRevoked(address indexed operator, address indexed contractAddress, uint32 contractCanSlashOperatorUntilBlock)
func (_ISlasher *ISlasherFilterer) WatchSlashingAbilityRevoked(opts *bind.WatchOpts, sink chan<- *ISlasherSlashingAbilityRevoked, operator []common.Address, contractAddress []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ISlasher.contract.WatchLogs(opts, "SlashingAbilityRevoked", operatorRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlasherSlashingAbilityRevoked)
				if err := _ISlasher.contract.UnpackLog(event, "SlashingAbilityRevoked", log); err != nil {
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

// ParseSlashingAbilityRevoked is a log parse operation binding the contract event 0x9aa1b1391f35c672ed1f3b7ece632f4513e618366bef7a2f67b7c6bc1f2d2b14.
//
// Solidity: event SlashingAbilityRevoked(address indexed operator, address indexed contractAddress, uint32 contractCanSlashOperatorUntilBlock)
func (_ISlasher *ISlasherFilterer) ParseSlashingAbilityRevoked(log types.Log) (*ISlasherSlashingAbilityRevoked, error) {
	event := new(ISlasherSlashingAbilityRevoked)
	if err := _ISlasher.contract.UnpackLog(event, "SlashingAbilityRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
