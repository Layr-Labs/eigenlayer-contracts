// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IReleaseManager

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

// IReleaseManagerTypesArtifact is an auto generated low-level Go binding around an user-defined struct.
type IReleaseManagerTypesArtifact struct {
	Digest   [32]byte
	Registry string
}

// IReleaseManagerTypesRelease is an auto generated low-level Go binding around an user-defined struct.
type IReleaseManagerTypesRelease struct {
	Artifacts     []IReleaseManagerTypesArtifact
	UpgradeByTime uint32
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// IReleaseManagerMetaData contains all meta data concerning the IReleaseManager contract.
var IReleaseManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getLatestRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestUpgradeByTime\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetadataURI\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalReleases\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishMetadataURI\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"publishRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"release\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"MetadataURIPublished\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasePublished\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"release\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidMetadataURI\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidUpgradeByTime\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustPublishMetadataURI\",\"inputs\":[]}]",
}

// IReleaseManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IReleaseManagerMetaData.ABI instead.
var IReleaseManagerABI = IReleaseManagerMetaData.ABI

// IReleaseManager is an auto generated Go binding around an Ethereum contract.
type IReleaseManager struct {
	IReleaseManagerCaller     // Read-only binding to the contract
	IReleaseManagerTransactor // Write-only binding to the contract
	IReleaseManagerFilterer   // Log filterer for contract events
}

// IReleaseManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IReleaseManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReleaseManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IReleaseManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReleaseManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IReleaseManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReleaseManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IReleaseManagerSession struct {
	Contract     *IReleaseManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IReleaseManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IReleaseManagerCallerSession struct {
	Contract *IReleaseManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IReleaseManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IReleaseManagerTransactorSession struct {
	Contract     *IReleaseManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IReleaseManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IReleaseManagerRaw struct {
	Contract *IReleaseManager // Generic contract binding to access the raw methods on
}

// IReleaseManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IReleaseManagerCallerRaw struct {
	Contract *IReleaseManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IReleaseManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IReleaseManagerTransactorRaw struct {
	Contract *IReleaseManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIReleaseManager creates a new instance of IReleaseManager, bound to a specific deployed contract.
func NewIReleaseManager(address common.Address, backend bind.ContractBackend) (*IReleaseManager, error) {
	contract, err := bindIReleaseManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IReleaseManager{IReleaseManagerCaller: IReleaseManagerCaller{contract: contract}, IReleaseManagerTransactor: IReleaseManagerTransactor{contract: contract}, IReleaseManagerFilterer: IReleaseManagerFilterer{contract: contract}}, nil
}

// NewIReleaseManagerCaller creates a new read-only instance of IReleaseManager, bound to a specific deployed contract.
func NewIReleaseManagerCaller(address common.Address, caller bind.ContractCaller) (*IReleaseManagerCaller, error) {
	contract, err := bindIReleaseManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IReleaseManagerCaller{contract: contract}, nil
}

// NewIReleaseManagerTransactor creates a new write-only instance of IReleaseManager, bound to a specific deployed contract.
func NewIReleaseManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IReleaseManagerTransactor, error) {
	contract, err := bindIReleaseManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IReleaseManagerTransactor{contract: contract}, nil
}

// NewIReleaseManagerFilterer creates a new log filterer instance of IReleaseManager, bound to a specific deployed contract.
func NewIReleaseManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IReleaseManagerFilterer, error) {
	contract, err := bindIReleaseManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IReleaseManagerFilterer{contract: contract}, nil
}

// bindIReleaseManager binds a generic wrapper to an already deployed contract.
func bindIReleaseManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IReleaseManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReleaseManager *IReleaseManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReleaseManager.Contract.IReleaseManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReleaseManager *IReleaseManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReleaseManager.Contract.IReleaseManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReleaseManager *IReleaseManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReleaseManager.Contract.IReleaseManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReleaseManager *IReleaseManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReleaseManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReleaseManager *IReleaseManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReleaseManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReleaseManager *IReleaseManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReleaseManager.Contract.contract.Transact(opts, method, params...)
}

// GetLatestRelease is a free data retrieval call binding the contract method 0xd30eeb88.
//
// Solidity: function getLatestRelease((address,uint32) operatorSet) view returns(uint256, ((bytes32,string)[],uint32))
func (_IReleaseManager *IReleaseManagerCaller) GetLatestRelease(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, IReleaseManagerTypesRelease, error) {
	var out []interface{}
	err := _IReleaseManager.contract.Call(opts, &out, "getLatestRelease", operatorSet)

	if err != nil {
		return *new(*big.Int), *new(IReleaseManagerTypesRelease), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(IReleaseManagerTypesRelease)).(*IReleaseManagerTypesRelease)

	return out0, out1, err

}

// GetLatestRelease is a free data retrieval call binding the contract method 0xd30eeb88.
//
// Solidity: function getLatestRelease((address,uint32) operatorSet) view returns(uint256, ((bytes32,string)[],uint32))
func (_IReleaseManager *IReleaseManagerSession) GetLatestRelease(operatorSet OperatorSet) (*big.Int, IReleaseManagerTypesRelease, error) {
	return _IReleaseManager.Contract.GetLatestRelease(&_IReleaseManager.CallOpts, operatorSet)
}

// GetLatestRelease is a free data retrieval call binding the contract method 0xd30eeb88.
//
// Solidity: function getLatestRelease((address,uint32) operatorSet) view returns(uint256, ((bytes32,string)[],uint32))
func (_IReleaseManager *IReleaseManagerCallerSession) GetLatestRelease(operatorSet OperatorSet) (*big.Int, IReleaseManagerTypesRelease, error) {
	return _IReleaseManager.Contract.GetLatestRelease(&_IReleaseManager.CallOpts, operatorSet)
}

// GetLatestUpgradeByTime is a free data retrieval call binding the contract method 0xa9e0ed68.
//
// Solidity: function getLatestUpgradeByTime((address,uint32) operatorSet) view returns(uint32)
func (_IReleaseManager *IReleaseManagerCaller) GetLatestUpgradeByTime(opts *bind.CallOpts, operatorSet OperatorSet) (uint32, error) {
	var out []interface{}
	err := _IReleaseManager.contract.Call(opts, &out, "getLatestUpgradeByTime", operatorSet)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetLatestUpgradeByTime is a free data retrieval call binding the contract method 0xa9e0ed68.
//
// Solidity: function getLatestUpgradeByTime((address,uint32) operatorSet) view returns(uint32)
func (_IReleaseManager *IReleaseManagerSession) GetLatestUpgradeByTime(operatorSet OperatorSet) (uint32, error) {
	return _IReleaseManager.Contract.GetLatestUpgradeByTime(&_IReleaseManager.CallOpts, operatorSet)
}

// GetLatestUpgradeByTime is a free data retrieval call binding the contract method 0xa9e0ed68.
//
// Solidity: function getLatestUpgradeByTime((address,uint32) operatorSet) view returns(uint32)
func (_IReleaseManager *IReleaseManagerCallerSession) GetLatestUpgradeByTime(operatorSet OperatorSet) (uint32, error) {
	return _IReleaseManager.Contract.GetLatestUpgradeByTime(&_IReleaseManager.CallOpts, operatorSet)
}

// GetMetadataURI is a free data retrieval call binding the contract method 0xb053b56d.
//
// Solidity: function getMetadataURI((address,uint32) operatorSet) view returns(string)
func (_IReleaseManager *IReleaseManagerCaller) GetMetadataURI(opts *bind.CallOpts, operatorSet OperatorSet) (string, error) {
	var out []interface{}
	err := _IReleaseManager.contract.Call(opts, &out, "getMetadataURI", operatorSet)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMetadataURI is a free data retrieval call binding the contract method 0xb053b56d.
//
// Solidity: function getMetadataURI((address,uint32) operatorSet) view returns(string)
func (_IReleaseManager *IReleaseManagerSession) GetMetadataURI(operatorSet OperatorSet) (string, error) {
	return _IReleaseManager.Contract.GetMetadataURI(&_IReleaseManager.CallOpts, operatorSet)
}

// GetMetadataURI is a free data retrieval call binding the contract method 0xb053b56d.
//
// Solidity: function getMetadataURI((address,uint32) operatorSet) view returns(string)
func (_IReleaseManager *IReleaseManagerCallerSession) GetMetadataURI(operatorSet OperatorSet) (string, error) {
	return _IReleaseManager.Contract.GetMetadataURI(&_IReleaseManager.CallOpts, operatorSet)
}

// GetRelease is a free data retrieval call binding the contract method 0x3acab5fc.
//
// Solidity: function getRelease((address,uint32) operatorSet, uint256 releaseId) view returns(((bytes32,string)[],uint32))
func (_IReleaseManager *IReleaseManagerCaller) GetRelease(opts *bind.CallOpts, operatorSet OperatorSet, releaseId *big.Int) (IReleaseManagerTypesRelease, error) {
	var out []interface{}
	err := _IReleaseManager.contract.Call(opts, &out, "getRelease", operatorSet, releaseId)

	if err != nil {
		return *new(IReleaseManagerTypesRelease), err
	}

	out0 := *abi.ConvertType(out[0], new(IReleaseManagerTypesRelease)).(*IReleaseManagerTypesRelease)

	return out0, err

}

// GetRelease is a free data retrieval call binding the contract method 0x3acab5fc.
//
// Solidity: function getRelease((address,uint32) operatorSet, uint256 releaseId) view returns(((bytes32,string)[],uint32))
func (_IReleaseManager *IReleaseManagerSession) GetRelease(operatorSet OperatorSet, releaseId *big.Int) (IReleaseManagerTypesRelease, error) {
	return _IReleaseManager.Contract.GetRelease(&_IReleaseManager.CallOpts, operatorSet, releaseId)
}

// GetRelease is a free data retrieval call binding the contract method 0x3acab5fc.
//
// Solidity: function getRelease((address,uint32) operatorSet, uint256 releaseId) view returns(((bytes32,string)[],uint32))
func (_IReleaseManager *IReleaseManagerCallerSession) GetRelease(operatorSet OperatorSet, releaseId *big.Int) (IReleaseManagerTypesRelease, error) {
	return _IReleaseManager.Contract.GetRelease(&_IReleaseManager.CallOpts, operatorSet, releaseId)
}

// GetTotalReleases is a free data retrieval call binding the contract method 0x66f409f7.
//
// Solidity: function getTotalReleases((address,uint32) operatorSet) view returns(uint256)
func (_IReleaseManager *IReleaseManagerCaller) GetTotalReleases(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _IReleaseManager.contract.Call(opts, &out, "getTotalReleases", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalReleases is a free data retrieval call binding the contract method 0x66f409f7.
//
// Solidity: function getTotalReleases((address,uint32) operatorSet) view returns(uint256)
func (_IReleaseManager *IReleaseManagerSession) GetTotalReleases(operatorSet OperatorSet) (*big.Int, error) {
	return _IReleaseManager.Contract.GetTotalReleases(&_IReleaseManager.CallOpts, operatorSet)
}

// GetTotalReleases is a free data retrieval call binding the contract method 0x66f409f7.
//
// Solidity: function getTotalReleases((address,uint32) operatorSet) view returns(uint256)
func (_IReleaseManager *IReleaseManagerCallerSession) GetTotalReleases(operatorSet OperatorSet) (*big.Int, error) {
	return _IReleaseManager.Contract.GetTotalReleases(&_IReleaseManager.CallOpts, operatorSet)
}

// IsValidRelease is a free data retrieval call binding the contract method 0x517e4068.
//
// Solidity: function isValidRelease((address,uint32) operatorSet, uint256 releaseId) view returns(bool)
func (_IReleaseManager *IReleaseManagerCaller) IsValidRelease(opts *bind.CallOpts, operatorSet OperatorSet, releaseId *big.Int) (bool, error) {
	var out []interface{}
	err := _IReleaseManager.contract.Call(opts, &out, "isValidRelease", operatorSet, releaseId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidRelease is a free data retrieval call binding the contract method 0x517e4068.
//
// Solidity: function isValidRelease((address,uint32) operatorSet, uint256 releaseId) view returns(bool)
func (_IReleaseManager *IReleaseManagerSession) IsValidRelease(operatorSet OperatorSet, releaseId *big.Int) (bool, error) {
	return _IReleaseManager.Contract.IsValidRelease(&_IReleaseManager.CallOpts, operatorSet, releaseId)
}

// IsValidRelease is a free data retrieval call binding the contract method 0x517e4068.
//
// Solidity: function isValidRelease((address,uint32) operatorSet, uint256 releaseId) view returns(bool)
func (_IReleaseManager *IReleaseManagerCallerSession) IsValidRelease(operatorSet OperatorSet, releaseId *big.Int) (bool, error) {
	return _IReleaseManager.Contract.IsValidRelease(&_IReleaseManager.CallOpts, operatorSet, releaseId)
}

// PublishMetadataURI is a paid mutator transaction binding the contract method 0x4840a67c.
//
// Solidity: function publishMetadataURI((address,uint32) operatorSet, string metadataURI) returns()
func (_IReleaseManager *IReleaseManagerTransactor) PublishMetadataURI(opts *bind.TransactOpts, operatorSet OperatorSet, metadataURI string) (*types.Transaction, error) {
	return _IReleaseManager.contract.Transact(opts, "publishMetadataURI", operatorSet, metadataURI)
}

// PublishMetadataURI is a paid mutator transaction binding the contract method 0x4840a67c.
//
// Solidity: function publishMetadataURI((address,uint32) operatorSet, string metadataURI) returns()
func (_IReleaseManager *IReleaseManagerSession) PublishMetadataURI(operatorSet OperatorSet, metadataURI string) (*types.Transaction, error) {
	return _IReleaseManager.Contract.PublishMetadataURI(&_IReleaseManager.TransactOpts, operatorSet, metadataURI)
}

// PublishMetadataURI is a paid mutator transaction binding the contract method 0x4840a67c.
//
// Solidity: function publishMetadataURI((address,uint32) operatorSet, string metadataURI) returns()
func (_IReleaseManager *IReleaseManagerTransactorSession) PublishMetadataURI(operatorSet OperatorSet, metadataURI string) (*types.Transaction, error) {
	return _IReleaseManager.Contract.PublishMetadataURI(&_IReleaseManager.TransactOpts, operatorSet, metadataURI)
}

// PublishRelease is a paid mutator transaction binding the contract method 0x7c09ea82.
//
// Solidity: function publishRelease((address,uint32) operatorSet, ((bytes32,string)[],uint32) release) returns(uint256 releaseId)
func (_IReleaseManager *IReleaseManagerTransactor) PublishRelease(opts *bind.TransactOpts, operatorSet OperatorSet, release IReleaseManagerTypesRelease) (*types.Transaction, error) {
	return _IReleaseManager.contract.Transact(opts, "publishRelease", operatorSet, release)
}

// PublishRelease is a paid mutator transaction binding the contract method 0x7c09ea82.
//
// Solidity: function publishRelease((address,uint32) operatorSet, ((bytes32,string)[],uint32) release) returns(uint256 releaseId)
func (_IReleaseManager *IReleaseManagerSession) PublishRelease(operatorSet OperatorSet, release IReleaseManagerTypesRelease) (*types.Transaction, error) {
	return _IReleaseManager.Contract.PublishRelease(&_IReleaseManager.TransactOpts, operatorSet, release)
}

// PublishRelease is a paid mutator transaction binding the contract method 0x7c09ea82.
//
// Solidity: function publishRelease((address,uint32) operatorSet, ((bytes32,string)[],uint32) release) returns(uint256 releaseId)
func (_IReleaseManager *IReleaseManagerTransactorSession) PublishRelease(operatorSet OperatorSet, release IReleaseManagerTypesRelease) (*types.Transaction, error) {
	return _IReleaseManager.Contract.PublishRelease(&_IReleaseManager.TransactOpts, operatorSet, release)
}

// IReleaseManagerMetadataURIPublishedIterator is returned from FilterMetadataURIPublished and is used to iterate over the raw logs and unpacked data for MetadataURIPublished events raised by the IReleaseManager contract.
type IReleaseManagerMetadataURIPublishedIterator struct {
	Event *IReleaseManagerMetadataURIPublished // Event containing the contract specifics and raw log

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
func (it *IReleaseManagerMetadataURIPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IReleaseManagerMetadataURIPublished)
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
		it.Event = new(IReleaseManagerMetadataURIPublished)
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
func (it *IReleaseManagerMetadataURIPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IReleaseManagerMetadataURIPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IReleaseManagerMetadataURIPublished represents a MetadataURIPublished event raised by the IReleaseManager contract.
type IReleaseManagerMetadataURIPublished struct {
	OperatorSet OperatorSet
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMetadataURIPublished is a free log retrieval operation binding the contract event 0x209e95fbe8dd14c5e1fbf791ee0a83234f45f20cb85504c7068d5ca0d6224588.
//
// Solidity: event MetadataURIPublished((address,uint32) indexed operatorSet, string metadataURI)
func (_IReleaseManager *IReleaseManagerFilterer) FilterMetadataURIPublished(opts *bind.FilterOpts, operatorSet []OperatorSet) (*IReleaseManagerMetadataURIPublishedIterator, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}

	logs, sub, err := _IReleaseManager.contract.FilterLogs(opts, "MetadataURIPublished", operatorSetRule)
	if err != nil {
		return nil, err
	}
	return &IReleaseManagerMetadataURIPublishedIterator{contract: _IReleaseManager.contract, event: "MetadataURIPublished", logs: logs, sub: sub}, nil
}

// WatchMetadataURIPublished is a free log subscription operation binding the contract event 0x209e95fbe8dd14c5e1fbf791ee0a83234f45f20cb85504c7068d5ca0d6224588.
//
// Solidity: event MetadataURIPublished((address,uint32) indexed operatorSet, string metadataURI)
func (_IReleaseManager *IReleaseManagerFilterer) WatchMetadataURIPublished(opts *bind.WatchOpts, sink chan<- *IReleaseManagerMetadataURIPublished, operatorSet []OperatorSet) (event.Subscription, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}

	logs, sub, err := _IReleaseManager.contract.WatchLogs(opts, "MetadataURIPublished", operatorSetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IReleaseManagerMetadataURIPublished)
				if err := _IReleaseManager.contract.UnpackLog(event, "MetadataURIPublished", log); err != nil {
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

// ParseMetadataURIPublished is a log parse operation binding the contract event 0x209e95fbe8dd14c5e1fbf791ee0a83234f45f20cb85504c7068d5ca0d6224588.
//
// Solidity: event MetadataURIPublished((address,uint32) indexed operatorSet, string metadataURI)
func (_IReleaseManager *IReleaseManagerFilterer) ParseMetadataURIPublished(log types.Log) (*IReleaseManagerMetadataURIPublished, error) {
	event := new(IReleaseManagerMetadataURIPublished)
	if err := _IReleaseManager.contract.UnpackLog(event, "MetadataURIPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IReleaseManagerReleasePublishedIterator is returned from FilterReleasePublished and is used to iterate over the raw logs and unpacked data for ReleasePublished events raised by the IReleaseManager contract.
type IReleaseManagerReleasePublishedIterator struct {
	Event *IReleaseManagerReleasePublished // Event containing the contract specifics and raw log

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
func (it *IReleaseManagerReleasePublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IReleaseManagerReleasePublished)
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
		it.Event = new(IReleaseManagerReleasePublished)
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
func (it *IReleaseManagerReleasePublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IReleaseManagerReleasePublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IReleaseManagerReleasePublished represents a ReleasePublished event raised by the IReleaseManager contract.
type IReleaseManagerReleasePublished struct {
	OperatorSet OperatorSet
	ReleaseId   *big.Int
	Release     IReleaseManagerTypesRelease
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReleasePublished is a free log retrieval operation binding the contract event 0x2decd15222f7c4a8c3d4d2e14dcfdc5a0b52eb2d4b81796bfd010ee5cd972fd3.
//
// Solidity: event ReleasePublished((address,uint32) indexed operatorSet, uint256 indexed releaseId, ((bytes32,string)[],uint32) release)
func (_IReleaseManager *IReleaseManagerFilterer) FilterReleasePublished(opts *bind.FilterOpts, operatorSet []OperatorSet, releaseId []*big.Int) (*IReleaseManagerReleasePublishedIterator, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var releaseIdRule []interface{}
	for _, releaseIdItem := range releaseId {
		releaseIdRule = append(releaseIdRule, releaseIdItem)
	}

	logs, sub, err := _IReleaseManager.contract.FilterLogs(opts, "ReleasePublished", operatorSetRule, releaseIdRule)
	if err != nil {
		return nil, err
	}
	return &IReleaseManagerReleasePublishedIterator{contract: _IReleaseManager.contract, event: "ReleasePublished", logs: logs, sub: sub}, nil
}

// WatchReleasePublished is a free log subscription operation binding the contract event 0x2decd15222f7c4a8c3d4d2e14dcfdc5a0b52eb2d4b81796bfd010ee5cd972fd3.
//
// Solidity: event ReleasePublished((address,uint32) indexed operatorSet, uint256 indexed releaseId, ((bytes32,string)[],uint32) release)
func (_IReleaseManager *IReleaseManagerFilterer) WatchReleasePublished(opts *bind.WatchOpts, sink chan<- *IReleaseManagerReleasePublished, operatorSet []OperatorSet, releaseId []*big.Int) (event.Subscription, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var releaseIdRule []interface{}
	for _, releaseIdItem := range releaseId {
		releaseIdRule = append(releaseIdRule, releaseIdItem)
	}

	logs, sub, err := _IReleaseManager.contract.WatchLogs(opts, "ReleasePublished", operatorSetRule, releaseIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IReleaseManagerReleasePublished)
				if err := _IReleaseManager.contract.UnpackLog(event, "ReleasePublished", log); err != nil {
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

// ParseReleasePublished is a log parse operation binding the contract event 0x2decd15222f7c4a8c3d4d2e14dcfdc5a0b52eb2d4b81796bfd010ee5cd972fd3.
//
// Solidity: event ReleasePublished((address,uint32) indexed operatorSet, uint256 indexed releaseId, ((bytes32,string)[],uint32) release)
func (_IReleaseManager *IReleaseManagerFilterer) ParseReleasePublished(log types.Log) (*IReleaseManagerReleasePublished, error) {
	event := new(IReleaseManagerReleasePublished)
	if err := _IReleaseManager.contract.UnpackLog(event, "ReleasePublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
