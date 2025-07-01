// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ReleaseManagerStorage

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

// ReleaseManagerStorageMetaData contains all meta data concerning the ReleaseManagerStorage contract.
var ReleaseManagerStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getLatestRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestUpgradeByTime\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetadataURI\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalReleases\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishMetadataURI\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"publishRelease\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"release\",\"type\":\"tuple\",\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"releaseId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"MetadataURIPublished\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasePublished\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"releaseId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"release\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIReleaseManagerTypes.Release\",\"components\":[{\"name\":\"artifacts\",\"type\":\"tuple[]\",\"internalType\":\"structIReleaseManagerTypes.Artifact[]\",\"components\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"registry\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"upgradeByTime\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidMetadataURI\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidUpgradeByTime\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustPublishMetadataURI\",\"inputs\":[]}]",
}

// ReleaseManagerStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use ReleaseManagerStorageMetaData.ABI instead.
var ReleaseManagerStorageABI = ReleaseManagerStorageMetaData.ABI

// ReleaseManagerStorage is an auto generated Go binding around an Ethereum contract.
type ReleaseManagerStorage struct {
	ReleaseManagerStorageCaller     // Read-only binding to the contract
	ReleaseManagerStorageTransactor // Write-only binding to the contract
	ReleaseManagerStorageFilterer   // Log filterer for contract events
}

// ReleaseManagerStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReleaseManagerStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseManagerStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReleaseManagerStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseManagerStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReleaseManagerStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseManagerStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReleaseManagerStorageSession struct {
	Contract     *ReleaseManagerStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ReleaseManagerStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReleaseManagerStorageCallerSession struct {
	Contract *ReleaseManagerStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ReleaseManagerStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReleaseManagerStorageTransactorSession struct {
	Contract     *ReleaseManagerStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ReleaseManagerStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReleaseManagerStorageRaw struct {
	Contract *ReleaseManagerStorage // Generic contract binding to access the raw methods on
}

// ReleaseManagerStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReleaseManagerStorageCallerRaw struct {
	Contract *ReleaseManagerStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ReleaseManagerStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReleaseManagerStorageTransactorRaw struct {
	Contract *ReleaseManagerStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReleaseManagerStorage creates a new instance of ReleaseManagerStorage, bound to a specific deployed contract.
func NewReleaseManagerStorage(address common.Address, backend bind.ContractBackend) (*ReleaseManagerStorage, error) {
	contract, err := bindReleaseManagerStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerStorage{ReleaseManagerStorageCaller: ReleaseManagerStorageCaller{contract: contract}, ReleaseManagerStorageTransactor: ReleaseManagerStorageTransactor{contract: contract}, ReleaseManagerStorageFilterer: ReleaseManagerStorageFilterer{contract: contract}}, nil
}

// NewReleaseManagerStorageCaller creates a new read-only instance of ReleaseManagerStorage, bound to a specific deployed contract.
func NewReleaseManagerStorageCaller(address common.Address, caller bind.ContractCaller) (*ReleaseManagerStorageCaller, error) {
	contract, err := bindReleaseManagerStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerStorageCaller{contract: contract}, nil
}

// NewReleaseManagerStorageTransactor creates a new write-only instance of ReleaseManagerStorage, bound to a specific deployed contract.
func NewReleaseManagerStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ReleaseManagerStorageTransactor, error) {
	contract, err := bindReleaseManagerStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerStorageTransactor{contract: contract}, nil
}

// NewReleaseManagerStorageFilterer creates a new log filterer instance of ReleaseManagerStorage, bound to a specific deployed contract.
func NewReleaseManagerStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ReleaseManagerStorageFilterer, error) {
	contract, err := bindReleaseManagerStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerStorageFilterer{contract: contract}, nil
}

// bindReleaseManagerStorage binds a generic wrapper to an already deployed contract.
func bindReleaseManagerStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReleaseManagerStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReleaseManagerStorage *ReleaseManagerStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReleaseManagerStorage.Contract.ReleaseManagerStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReleaseManagerStorage *ReleaseManagerStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.ReleaseManagerStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReleaseManagerStorage *ReleaseManagerStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.ReleaseManagerStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReleaseManagerStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.contract.Transact(opts, method, params...)
}

// GetLatestRelease is a free data retrieval call binding the contract method 0xd30eeb88.
//
// Solidity: function getLatestRelease((address,uint32) operatorSet) view returns(uint256, ((bytes32,string)[],uint32))
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) GetLatestRelease(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, IReleaseManagerTypesRelease, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "getLatestRelease", operatorSet)

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
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) GetLatestRelease(operatorSet OperatorSet) (*big.Int, IReleaseManagerTypesRelease, error) {
	return _ReleaseManagerStorage.Contract.GetLatestRelease(&_ReleaseManagerStorage.CallOpts, operatorSet)
}

// GetLatestRelease is a free data retrieval call binding the contract method 0xd30eeb88.
//
// Solidity: function getLatestRelease((address,uint32) operatorSet) view returns(uint256, ((bytes32,string)[],uint32))
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) GetLatestRelease(operatorSet OperatorSet) (*big.Int, IReleaseManagerTypesRelease, error) {
	return _ReleaseManagerStorage.Contract.GetLatestRelease(&_ReleaseManagerStorage.CallOpts, operatorSet)
}

// GetLatestUpgradeByTime is a free data retrieval call binding the contract method 0xa9e0ed68.
//
// Solidity: function getLatestUpgradeByTime((address,uint32) operatorSet) view returns(uint32)
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) GetLatestUpgradeByTime(opts *bind.CallOpts, operatorSet OperatorSet) (uint32, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "getLatestUpgradeByTime", operatorSet)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetLatestUpgradeByTime is a free data retrieval call binding the contract method 0xa9e0ed68.
//
// Solidity: function getLatestUpgradeByTime((address,uint32) operatorSet) view returns(uint32)
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) GetLatestUpgradeByTime(operatorSet OperatorSet) (uint32, error) {
	return _ReleaseManagerStorage.Contract.GetLatestUpgradeByTime(&_ReleaseManagerStorage.CallOpts, operatorSet)
}

// GetLatestUpgradeByTime is a free data retrieval call binding the contract method 0xa9e0ed68.
//
// Solidity: function getLatestUpgradeByTime((address,uint32) operatorSet) view returns(uint32)
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) GetLatestUpgradeByTime(operatorSet OperatorSet) (uint32, error) {
	return _ReleaseManagerStorage.Contract.GetLatestUpgradeByTime(&_ReleaseManagerStorage.CallOpts, operatorSet)
}

// GetMetadataURI is a free data retrieval call binding the contract method 0xb053b56d.
//
// Solidity: function getMetadataURI((address,uint32) operatorSet) view returns(string)
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) GetMetadataURI(opts *bind.CallOpts, operatorSet OperatorSet) (string, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "getMetadataURI", operatorSet)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMetadataURI is a free data retrieval call binding the contract method 0xb053b56d.
//
// Solidity: function getMetadataURI((address,uint32) operatorSet) view returns(string)
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) GetMetadataURI(operatorSet OperatorSet) (string, error) {
	return _ReleaseManagerStorage.Contract.GetMetadataURI(&_ReleaseManagerStorage.CallOpts, operatorSet)
}

// GetMetadataURI is a free data retrieval call binding the contract method 0xb053b56d.
//
// Solidity: function getMetadataURI((address,uint32) operatorSet) view returns(string)
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) GetMetadataURI(operatorSet OperatorSet) (string, error) {
	return _ReleaseManagerStorage.Contract.GetMetadataURI(&_ReleaseManagerStorage.CallOpts, operatorSet)
}

// GetRelease is a free data retrieval call binding the contract method 0x3acab5fc.
//
// Solidity: function getRelease((address,uint32) operatorSet, uint256 releaseId) view returns(((bytes32,string)[],uint32))
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) GetRelease(opts *bind.CallOpts, operatorSet OperatorSet, releaseId *big.Int) (IReleaseManagerTypesRelease, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "getRelease", operatorSet, releaseId)

	if err != nil {
		return *new(IReleaseManagerTypesRelease), err
	}

	out0 := *abi.ConvertType(out[0], new(IReleaseManagerTypesRelease)).(*IReleaseManagerTypesRelease)

	return out0, err

}

// GetRelease is a free data retrieval call binding the contract method 0x3acab5fc.
//
// Solidity: function getRelease((address,uint32) operatorSet, uint256 releaseId) view returns(((bytes32,string)[],uint32))
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) GetRelease(operatorSet OperatorSet, releaseId *big.Int) (IReleaseManagerTypesRelease, error) {
	return _ReleaseManagerStorage.Contract.GetRelease(&_ReleaseManagerStorage.CallOpts, operatorSet, releaseId)
}

// GetRelease is a free data retrieval call binding the contract method 0x3acab5fc.
//
// Solidity: function getRelease((address,uint32) operatorSet, uint256 releaseId) view returns(((bytes32,string)[],uint32))
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) GetRelease(operatorSet OperatorSet, releaseId *big.Int) (IReleaseManagerTypesRelease, error) {
	return _ReleaseManagerStorage.Contract.GetRelease(&_ReleaseManagerStorage.CallOpts, operatorSet, releaseId)
}

// GetTotalReleases is a free data retrieval call binding the contract method 0x66f409f7.
//
// Solidity: function getTotalReleases((address,uint32) operatorSet) view returns(uint256)
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) GetTotalReleases(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "getTotalReleases", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalReleases is a free data retrieval call binding the contract method 0x66f409f7.
//
// Solidity: function getTotalReleases((address,uint32) operatorSet) view returns(uint256)
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) GetTotalReleases(operatorSet OperatorSet) (*big.Int, error) {
	return _ReleaseManagerStorage.Contract.GetTotalReleases(&_ReleaseManagerStorage.CallOpts, operatorSet)
}

// GetTotalReleases is a free data retrieval call binding the contract method 0x66f409f7.
//
// Solidity: function getTotalReleases((address,uint32) operatorSet) view returns(uint256)
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) GetTotalReleases(operatorSet OperatorSet) (*big.Int, error) {
	return _ReleaseManagerStorage.Contract.GetTotalReleases(&_ReleaseManagerStorage.CallOpts, operatorSet)
}

// IsValidRelease is a free data retrieval call binding the contract method 0x517e4068.
//
// Solidity: function isValidRelease((address,uint32) operatorSet, uint256 releaseId) view returns(bool)
func (_ReleaseManagerStorage *ReleaseManagerStorageCaller) IsValidRelease(opts *bind.CallOpts, operatorSet OperatorSet, releaseId *big.Int) (bool, error) {
	var out []interface{}
	err := _ReleaseManagerStorage.contract.Call(opts, &out, "isValidRelease", operatorSet, releaseId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidRelease is a free data retrieval call binding the contract method 0x517e4068.
//
// Solidity: function isValidRelease((address,uint32) operatorSet, uint256 releaseId) view returns(bool)
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) IsValidRelease(operatorSet OperatorSet, releaseId *big.Int) (bool, error) {
	return _ReleaseManagerStorage.Contract.IsValidRelease(&_ReleaseManagerStorage.CallOpts, operatorSet, releaseId)
}

// IsValidRelease is a free data retrieval call binding the contract method 0x517e4068.
//
// Solidity: function isValidRelease((address,uint32) operatorSet, uint256 releaseId) view returns(bool)
func (_ReleaseManagerStorage *ReleaseManagerStorageCallerSession) IsValidRelease(operatorSet OperatorSet, releaseId *big.Int) (bool, error) {
	return _ReleaseManagerStorage.Contract.IsValidRelease(&_ReleaseManagerStorage.CallOpts, operatorSet, releaseId)
}

// PublishMetadataURI is a paid mutator transaction binding the contract method 0x4840a67c.
//
// Solidity: function publishMetadataURI((address,uint32) operatorSet, string metadataURI) returns()
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactor) PublishMetadataURI(opts *bind.TransactOpts, operatorSet OperatorSet, metadataURI string) (*types.Transaction, error) {
	return _ReleaseManagerStorage.contract.Transact(opts, "publishMetadataURI", operatorSet, metadataURI)
}

// PublishMetadataURI is a paid mutator transaction binding the contract method 0x4840a67c.
//
// Solidity: function publishMetadataURI((address,uint32) operatorSet, string metadataURI) returns()
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) PublishMetadataURI(operatorSet OperatorSet, metadataURI string) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.PublishMetadataURI(&_ReleaseManagerStorage.TransactOpts, operatorSet, metadataURI)
}

// PublishMetadataURI is a paid mutator transaction binding the contract method 0x4840a67c.
//
// Solidity: function publishMetadataURI((address,uint32) operatorSet, string metadataURI) returns()
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactorSession) PublishMetadataURI(operatorSet OperatorSet, metadataURI string) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.PublishMetadataURI(&_ReleaseManagerStorage.TransactOpts, operatorSet, metadataURI)
}

// PublishRelease is a paid mutator transaction binding the contract method 0x7c09ea82.
//
// Solidity: function publishRelease((address,uint32) operatorSet, ((bytes32,string)[],uint32) release) returns(uint256 releaseId)
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactor) PublishRelease(opts *bind.TransactOpts, operatorSet OperatorSet, release IReleaseManagerTypesRelease) (*types.Transaction, error) {
	return _ReleaseManagerStorage.contract.Transact(opts, "publishRelease", operatorSet, release)
}

// PublishRelease is a paid mutator transaction binding the contract method 0x7c09ea82.
//
// Solidity: function publishRelease((address,uint32) operatorSet, ((bytes32,string)[],uint32) release) returns(uint256 releaseId)
func (_ReleaseManagerStorage *ReleaseManagerStorageSession) PublishRelease(operatorSet OperatorSet, release IReleaseManagerTypesRelease) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.PublishRelease(&_ReleaseManagerStorage.TransactOpts, operatorSet, release)
}

// PublishRelease is a paid mutator transaction binding the contract method 0x7c09ea82.
//
// Solidity: function publishRelease((address,uint32) operatorSet, ((bytes32,string)[],uint32) release) returns(uint256 releaseId)
func (_ReleaseManagerStorage *ReleaseManagerStorageTransactorSession) PublishRelease(operatorSet OperatorSet, release IReleaseManagerTypesRelease) (*types.Transaction, error) {
	return _ReleaseManagerStorage.Contract.PublishRelease(&_ReleaseManagerStorage.TransactOpts, operatorSet, release)
}

// ReleaseManagerStorageMetadataURIPublishedIterator is returned from FilterMetadataURIPublished and is used to iterate over the raw logs and unpacked data for MetadataURIPublished events raised by the ReleaseManagerStorage contract.
type ReleaseManagerStorageMetadataURIPublishedIterator struct {
	Event *ReleaseManagerStorageMetadataURIPublished // Event containing the contract specifics and raw log

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
func (it *ReleaseManagerStorageMetadataURIPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseManagerStorageMetadataURIPublished)
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
		it.Event = new(ReleaseManagerStorageMetadataURIPublished)
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
func (it *ReleaseManagerStorageMetadataURIPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseManagerStorageMetadataURIPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseManagerStorageMetadataURIPublished represents a MetadataURIPublished event raised by the ReleaseManagerStorage contract.
type ReleaseManagerStorageMetadataURIPublished struct {
	OperatorSet OperatorSet
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMetadataURIPublished is a free log retrieval operation binding the contract event 0x209e95fbe8dd14c5e1fbf791ee0a83234f45f20cb85504c7068d5ca0d6224588.
//
// Solidity: event MetadataURIPublished((address,uint32) indexed operatorSet, string metadataURI)
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) FilterMetadataURIPublished(opts *bind.FilterOpts, operatorSet []OperatorSet) (*ReleaseManagerStorageMetadataURIPublishedIterator, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}

	logs, sub, err := _ReleaseManagerStorage.contract.FilterLogs(opts, "MetadataURIPublished", operatorSetRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerStorageMetadataURIPublishedIterator{contract: _ReleaseManagerStorage.contract, event: "MetadataURIPublished", logs: logs, sub: sub}, nil
}

// WatchMetadataURIPublished is a free log subscription operation binding the contract event 0x209e95fbe8dd14c5e1fbf791ee0a83234f45f20cb85504c7068d5ca0d6224588.
//
// Solidity: event MetadataURIPublished((address,uint32) indexed operatorSet, string metadataURI)
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) WatchMetadataURIPublished(opts *bind.WatchOpts, sink chan<- *ReleaseManagerStorageMetadataURIPublished, operatorSet []OperatorSet) (event.Subscription, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}

	logs, sub, err := _ReleaseManagerStorage.contract.WatchLogs(opts, "MetadataURIPublished", operatorSetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseManagerStorageMetadataURIPublished)
				if err := _ReleaseManagerStorage.contract.UnpackLog(event, "MetadataURIPublished", log); err != nil {
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
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) ParseMetadataURIPublished(log types.Log) (*ReleaseManagerStorageMetadataURIPublished, error) {
	event := new(ReleaseManagerStorageMetadataURIPublished)
	if err := _ReleaseManagerStorage.contract.UnpackLog(event, "MetadataURIPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReleaseManagerStorageReleasePublishedIterator is returned from FilterReleasePublished and is used to iterate over the raw logs and unpacked data for ReleasePublished events raised by the ReleaseManagerStorage contract.
type ReleaseManagerStorageReleasePublishedIterator struct {
	Event *ReleaseManagerStorageReleasePublished // Event containing the contract specifics and raw log

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
func (it *ReleaseManagerStorageReleasePublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseManagerStorageReleasePublished)
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
		it.Event = new(ReleaseManagerStorageReleasePublished)
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
func (it *ReleaseManagerStorageReleasePublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseManagerStorageReleasePublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseManagerStorageReleasePublished represents a ReleasePublished event raised by the ReleaseManagerStorage contract.
type ReleaseManagerStorageReleasePublished struct {
	OperatorSet OperatorSet
	ReleaseId   *big.Int
	Release     IReleaseManagerTypesRelease
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReleasePublished is a free log retrieval operation binding the contract event 0x2decd15222f7c4a8c3d4d2e14dcfdc5a0b52eb2d4b81796bfd010ee5cd972fd3.
//
// Solidity: event ReleasePublished((address,uint32) indexed operatorSet, uint256 indexed releaseId, ((bytes32,string)[],uint32) release)
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) FilterReleasePublished(opts *bind.FilterOpts, operatorSet []OperatorSet, releaseId []*big.Int) (*ReleaseManagerStorageReleasePublishedIterator, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var releaseIdRule []interface{}
	for _, releaseIdItem := range releaseId {
		releaseIdRule = append(releaseIdRule, releaseIdItem)
	}

	logs, sub, err := _ReleaseManagerStorage.contract.FilterLogs(opts, "ReleasePublished", operatorSetRule, releaseIdRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseManagerStorageReleasePublishedIterator{contract: _ReleaseManagerStorage.contract, event: "ReleasePublished", logs: logs, sub: sub}, nil
}

// WatchReleasePublished is a free log subscription operation binding the contract event 0x2decd15222f7c4a8c3d4d2e14dcfdc5a0b52eb2d4b81796bfd010ee5cd972fd3.
//
// Solidity: event ReleasePublished((address,uint32) indexed operatorSet, uint256 indexed releaseId, ((bytes32,string)[],uint32) release)
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) WatchReleasePublished(opts *bind.WatchOpts, sink chan<- *ReleaseManagerStorageReleasePublished, operatorSet []OperatorSet, releaseId []*big.Int) (event.Subscription, error) {

	var operatorSetRule []interface{}
	for _, operatorSetItem := range operatorSet {
		operatorSetRule = append(operatorSetRule, operatorSetItem)
	}
	var releaseIdRule []interface{}
	for _, releaseIdItem := range releaseId {
		releaseIdRule = append(releaseIdRule, releaseIdItem)
	}

	logs, sub, err := _ReleaseManagerStorage.contract.WatchLogs(opts, "ReleasePublished", operatorSetRule, releaseIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseManagerStorageReleasePublished)
				if err := _ReleaseManagerStorage.contract.UnpackLog(event, "ReleasePublished", log); err != nil {
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
func (_ReleaseManagerStorage *ReleaseManagerStorageFilterer) ParseReleasePublished(log types.Log) (*ReleaseManagerStorageReleasePublished, error) {
	event := new(ReleaseManagerStorageReleasePublished)
	if err := _ReleaseManagerStorage.contract.UnpackLog(event, "ReleasePublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
