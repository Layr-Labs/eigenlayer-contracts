// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IProxyAdmin

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

// IProxyAdminMetaData contains all meta data concerning the IProxyAdmin contract.
var IProxyAdminMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"changeProxyAdmin\",\"inputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"internalType\":\"contractITransparentUpgradeableProxy\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getProxyAdmin\",\"inputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"internalType\":\"contractITransparentUpgradeableProxy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProxyImplementation\",\"inputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"internalType\":\"contractITransparentUpgradeableProxy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgrade\",\"inputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"internalType\":\"contractITransparentUpgradeableProxy\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeAndCall\",\"inputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"internalType\":\"contractITransparentUpgradeableProxy\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"}]",
}

// IProxyAdminABI is the input ABI used to generate the binding from.
// Deprecated: Use IProxyAdminMetaData.ABI instead.
var IProxyAdminABI = IProxyAdminMetaData.ABI

// IProxyAdmin is an auto generated Go binding around an Ethereum contract.
type IProxyAdmin struct {
	IProxyAdminCaller     // Read-only binding to the contract
	IProxyAdminTransactor // Write-only binding to the contract
	IProxyAdminFilterer   // Log filterer for contract events
}

// IProxyAdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type IProxyAdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProxyAdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IProxyAdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProxyAdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IProxyAdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProxyAdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IProxyAdminSession struct {
	Contract     *IProxyAdmin      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IProxyAdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IProxyAdminCallerSession struct {
	Contract *IProxyAdminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IProxyAdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IProxyAdminTransactorSession struct {
	Contract     *IProxyAdminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IProxyAdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type IProxyAdminRaw struct {
	Contract *IProxyAdmin // Generic contract binding to access the raw methods on
}

// IProxyAdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IProxyAdminCallerRaw struct {
	Contract *IProxyAdminCaller // Generic read-only contract binding to access the raw methods on
}

// IProxyAdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IProxyAdminTransactorRaw struct {
	Contract *IProxyAdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIProxyAdmin creates a new instance of IProxyAdmin, bound to a specific deployed contract.
func NewIProxyAdmin(address common.Address, backend bind.ContractBackend) (*IProxyAdmin, error) {
	contract, err := bindIProxyAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IProxyAdmin{IProxyAdminCaller: IProxyAdminCaller{contract: contract}, IProxyAdminTransactor: IProxyAdminTransactor{contract: contract}, IProxyAdminFilterer: IProxyAdminFilterer{contract: contract}}, nil
}

// NewIProxyAdminCaller creates a new read-only instance of IProxyAdmin, bound to a specific deployed contract.
func NewIProxyAdminCaller(address common.Address, caller bind.ContractCaller) (*IProxyAdminCaller, error) {
	contract, err := bindIProxyAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IProxyAdminCaller{contract: contract}, nil
}

// NewIProxyAdminTransactor creates a new write-only instance of IProxyAdmin, bound to a specific deployed contract.
func NewIProxyAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*IProxyAdminTransactor, error) {
	contract, err := bindIProxyAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IProxyAdminTransactor{contract: contract}, nil
}

// NewIProxyAdminFilterer creates a new log filterer instance of IProxyAdmin, bound to a specific deployed contract.
func NewIProxyAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*IProxyAdminFilterer, error) {
	contract, err := bindIProxyAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IProxyAdminFilterer{contract: contract}, nil
}

// bindIProxyAdmin binds a generic wrapper to an already deployed contract.
func bindIProxyAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IProxyAdminMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProxyAdmin *IProxyAdminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProxyAdmin.Contract.IProxyAdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProxyAdmin *IProxyAdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProxyAdmin.Contract.IProxyAdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProxyAdmin *IProxyAdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProxyAdmin.Contract.IProxyAdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProxyAdmin *IProxyAdminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProxyAdmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProxyAdmin *IProxyAdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProxyAdmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProxyAdmin *IProxyAdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProxyAdmin.Contract.contract.Transact(opts, method, params...)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_IProxyAdmin *IProxyAdminCaller) GetProxyAdmin(opts *bind.CallOpts, proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _IProxyAdmin.contract.Call(opts, &out, "getProxyAdmin", proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_IProxyAdmin *IProxyAdminSession) GetProxyAdmin(proxy common.Address) (common.Address, error) {
	return _IProxyAdmin.Contract.GetProxyAdmin(&_IProxyAdmin.CallOpts, proxy)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_IProxyAdmin *IProxyAdminCallerSession) GetProxyAdmin(proxy common.Address) (common.Address, error) {
	return _IProxyAdmin.Contract.GetProxyAdmin(&_IProxyAdmin.CallOpts, proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_IProxyAdmin *IProxyAdminCaller) GetProxyImplementation(opts *bind.CallOpts, proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _IProxyAdmin.contract.Call(opts, &out, "getProxyImplementation", proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_IProxyAdmin *IProxyAdminSession) GetProxyImplementation(proxy common.Address) (common.Address, error) {
	return _IProxyAdmin.Contract.GetProxyImplementation(&_IProxyAdmin.CallOpts, proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_IProxyAdmin *IProxyAdminCallerSession) GetProxyImplementation(proxy common.Address) (common.Address, error) {
	return _IProxyAdmin.Contract.GetProxyImplementation(&_IProxyAdmin.CallOpts, proxy)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_IProxyAdmin *IProxyAdminTransactor) ChangeProxyAdmin(opts *bind.TransactOpts, proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _IProxyAdmin.contract.Transact(opts, "changeProxyAdmin", proxy, newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_IProxyAdmin *IProxyAdminSession) ChangeProxyAdmin(proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _IProxyAdmin.Contract.ChangeProxyAdmin(&_IProxyAdmin.TransactOpts, proxy, newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_IProxyAdmin *IProxyAdminTransactorSession) ChangeProxyAdmin(proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _IProxyAdmin.Contract.ChangeProxyAdmin(&_IProxyAdmin.TransactOpts, proxy, newAdmin)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_IProxyAdmin *IProxyAdminTransactor) Upgrade(opts *bind.TransactOpts, proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _IProxyAdmin.contract.Transact(opts, "upgrade", proxy, implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_IProxyAdmin *IProxyAdminSession) Upgrade(proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _IProxyAdmin.Contract.Upgrade(&_IProxyAdmin.TransactOpts, proxy, implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_IProxyAdmin *IProxyAdminTransactorSession) Upgrade(proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _IProxyAdmin.Contract.Upgrade(&_IProxyAdmin.TransactOpts, proxy, implementation)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_IProxyAdmin *IProxyAdminTransactor) UpgradeAndCall(opts *bind.TransactOpts, proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _IProxyAdmin.contract.Transact(opts, "upgradeAndCall", proxy, implementation, data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_IProxyAdmin *IProxyAdminSession) UpgradeAndCall(proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _IProxyAdmin.Contract.UpgradeAndCall(&_IProxyAdmin.TransactOpts, proxy, implementation, data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_IProxyAdmin *IProxyAdminTransactorSession) UpgradeAndCall(proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _IProxyAdmin.Contract.UpgradeAndCall(&_IProxyAdmin.TransactOpts, proxy, implementation, data)
}
