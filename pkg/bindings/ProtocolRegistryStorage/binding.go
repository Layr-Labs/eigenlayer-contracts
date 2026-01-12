// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ProtocolRegistryStorage

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

// IProtocolRegistryTypesDeploymentConfig is an auto generated low-level Go binding around an user-defined struct.
type IProtocolRegistryTypesDeploymentConfig struct {
	Pausable   bool
	Deprecated bool
}

// ProtocolRegistryStorageMetaData contains all meta data concerning the ProtocolRegistryStorage contract.
var ProtocolRegistryStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"configure\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAddress\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllDeployments\",\"inputs\":[],\"outputs\":[{\"name\":\"names\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"configs\",\"type\":\"tuple[]\",\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig[]\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeployment\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pauserMultisig\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"majorVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ship\",\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"configs\",\"type\":\"tuple[]\",\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig[]\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"contractNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"semanticVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalDeployments\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DeploymentConfigDeleted\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeploymentConfigured\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeploymentShipped\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SemanticVersionUpdated\",\"inputs\":[{\"name\":\"previousSemanticVersion\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"semanticVersion\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DeploymentNotShipped\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputNameEmpty\",\"inputs\":[]}]",
}

// ProtocolRegistryStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use ProtocolRegistryStorageMetaData.ABI instead.
var ProtocolRegistryStorageABI = ProtocolRegistryStorageMetaData.ABI

// ProtocolRegistryStorage is an auto generated Go binding around an Ethereum contract.
type ProtocolRegistryStorage struct {
	ProtocolRegistryStorageCaller     // Read-only binding to the contract
	ProtocolRegistryStorageTransactor // Write-only binding to the contract
	ProtocolRegistryStorageFilterer   // Log filterer for contract events
}

// ProtocolRegistryStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtocolRegistryStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolRegistryStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtocolRegistryStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolRegistryStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtocolRegistryStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolRegistryStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtocolRegistryStorageSession struct {
	Contract     *ProtocolRegistryStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ProtocolRegistryStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtocolRegistryStorageCallerSession struct {
	Contract *ProtocolRegistryStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ProtocolRegistryStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtocolRegistryStorageTransactorSession struct {
	Contract     *ProtocolRegistryStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ProtocolRegistryStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtocolRegistryStorageRaw struct {
	Contract *ProtocolRegistryStorage // Generic contract binding to access the raw methods on
}

// ProtocolRegistryStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtocolRegistryStorageCallerRaw struct {
	Contract *ProtocolRegistryStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ProtocolRegistryStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtocolRegistryStorageTransactorRaw struct {
	Contract *ProtocolRegistryStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocolRegistryStorage creates a new instance of ProtocolRegistryStorage, bound to a specific deployed contract.
func NewProtocolRegistryStorage(address common.Address, backend bind.ContractBackend) (*ProtocolRegistryStorage, error) {
	contract, err := bindProtocolRegistryStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryStorage{ProtocolRegistryStorageCaller: ProtocolRegistryStorageCaller{contract: contract}, ProtocolRegistryStorageTransactor: ProtocolRegistryStorageTransactor{contract: contract}, ProtocolRegistryStorageFilterer: ProtocolRegistryStorageFilterer{contract: contract}}, nil
}

// NewProtocolRegistryStorageCaller creates a new read-only instance of ProtocolRegistryStorage, bound to a specific deployed contract.
func NewProtocolRegistryStorageCaller(address common.Address, caller bind.ContractCaller) (*ProtocolRegistryStorageCaller, error) {
	contract, err := bindProtocolRegistryStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryStorageCaller{contract: contract}, nil
}

// NewProtocolRegistryStorageTransactor creates a new write-only instance of ProtocolRegistryStorage, bound to a specific deployed contract.
func NewProtocolRegistryStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtocolRegistryStorageTransactor, error) {
	contract, err := bindProtocolRegistryStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryStorageTransactor{contract: contract}, nil
}

// NewProtocolRegistryStorageFilterer creates a new log filterer instance of ProtocolRegistryStorage, bound to a specific deployed contract.
func NewProtocolRegistryStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtocolRegistryStorageFilterer, error) {
	contract, err := bindProtocolRegistryStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryStorageFilterer{contract: contract}, nil
}

// bindProtocolRegistryStorage binds a generic wrapper to an already deployed contract.
func bindProtocolRegistryStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProtocolRegistryStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolRegistryStorage *ProtocolRegistryStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtocolRegistryStorage.Contract.ProtocolRegistryStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolRegistryStorage *ProtocolRegistryStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolRegistryStorage.Contract.ProtocolRegistryStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolRegistryStorage *ProtocolRegistryStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolRegistryStorage.Contract.ProtocolRegistryStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolRegistryStorage *ProtocolRegistryStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtocolRegistryStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolRegistryStorage *ProtocolRegistryStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolRegistryStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolRegistryStorage *ProtocolRegistryStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolRegistryStorage.Contract.contract.Transact(opts, method, params...)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProtocolRegistryStorage.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageSession) PAUSERROLE() ([32]byte, error) {
	return _ProtocolRegistryStorage.Contract.PAUSERROLE(&_ProtocolRegistryStorage.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ProtocolRegistryStorage.Contract.PAUSERROLE(&_ProtocolRegistryStorage.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string name) view returns(address)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageCaller) GetAddress(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _ProtocolRegistryStorage.contract.Call(opts, &out, "getAddress", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string name) view returns(address)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageSession) GetAddress(name string) (common.Address, error) {
	return _ProtocolRegistryStorage.Contract.GetAddress(&_ProtocolRegistryStorage.CallOpts, name)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string name) view returns(address)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageCallerSession) GetAddress(name string) (common.Address, error) {
	return _ProtocolRegistryStorage.Contract.GetAddress(&_ProtocolRegistryStorage.CallOpts, name)
}

// GetAllDeployments is a free data retrieval call binding the contract method 0x8eec00b8.
//
// Solidity: function getAllDeployments() view returns(string[] names, address[] addresses, (bool,bool)[] configs)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageCaller) GetAllDeployments(opts *bind.CallOpts) (struct {
	Names     []string
	Addresses []common.Address
	Configs   []IProtocolRegistryTypesDeploymentConfig
}, error) {
	var out []interface{}
	err := _ProtocolRegistryStorage.contract.Call(opts, &out, "getAllDeployments")

	outstruct := new(struct {
		Names     []string
		Addresses []common.Address
		Configs   []IProtocolRegistryTypesDeploymentConfig
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Names = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Addresses = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Configs = *abi.ConvertType(out[2], new([]IProtocolRegistryTypesDeploymentConfig)).(*[]IProtocolRegistryTypesDeploymentConfig)

	return *outstruct, err

}

// GetAllDeployments is a free data retrieval call binding the contract method 0x8eec00b8.
//
// Solidity: function getAllDeployments() view returns(string[] names, address[] addresses, (bool,bool)[] configs)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageSession) GetAllDeployments() (struct {
	Names     []string
	Addresses []common.Address
	Configs   []IProtocolRegistryTypesDeploymentConfig
}, error) {
	return _ProtocolRegistryStorage.Contract.GetAllDeployments(&_ProtocolRegistryStorage.CallOpts)
}

// GetAllDeployments is a free data retrieval call binding the contract method 0x8eec00b8.
//
// Solidity: function getAllDeployments() view returns(string[] names, address[] addresses, (bool,bool)[] configs)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageCallerSession) GetAllDeployments() (struct {
	Names     []string
	Addresses []common.Address
	Configs   []IProtocolRegistryTypesDeploymentConfig
}, error) {
	return _ProtocolRegistryStorage.Contract.GetAllDeployments(&_ProtocolRegistryStorage.CallOpts)
}

// GetDeployment is a free data retrieval call binding the contract method 0xa8091d97.
//
// Solidity: function getDeployment(string name) view returns(address addr, (bool,bool) config)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageCaller) GetDeployment(opts *bind.CallOpts, name string) (struct {
	Addr   common.Address
	Config IProtocolRegistryTypesDeploymentConfig
}, error) {
	var out []interface{}
	err := _ProtocolRegistryStorage.contract.Call(opts, &out, "getDeployment", name)

	outstruct := new(struct {
		Addr   common.Address
		Config IProtocolRegistryTypesDeploymentConfig
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Config = *abi.ConvertType(out[1], new(IProtocolRegistryTypesDeploymentConfig)).(*IProtocolRegistryTypesDeploymentConfig)

	return *outstruct, err

}

// GetDeployment is a free data retrieval call binding the contract method 0xa8091d97.
//
// Solidity: function getDeployment(string name) view returns(address addr, (bool,bool) config)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageSession) GetDeployment(name string) (struct {
	Addr   common.Address
	Config IProtocolRegistryTypesDeploymentConfig
}, error) {
	return _ProtocolRegistryStorage.Contract.GetDeployment(&_ProtocolRegistryStorage.CallOpts, name)
}

// GetDeployment is a free data retrieval call binding the contract method 0xa8091d97.
//
// Solidity: function getDeployment(string name) view returns(address addr, (bool,bool) config)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageCallerSession) GetDeployment(name string) (struct {
	Addr   common.Address
	Config IProtocolRegistryTypesDeploymentConfig
}, error) {
	return _ProtocolRegistryStorage.Contract.GetDeployment(&_ProtocolRegistryStorage.CallOpts, name)
}

// MajorVersion is a free data retrieval call binding the contract method 0xaf05a5c5.
//
// Solidity: function majorVersion() view returns(string)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageCaller) MajorVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProtocolRegistryStorage.contract.Call(opts, &out, "majorVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MajorVersion is a free data retrieval call binding the contract method 0xaf05a5c5.
//
// Solidity: function majorVersion() view returns(string)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageSession) MajorVersion() (string, error) {
	return _ProtocolRegistryStorage.Contract.MajorVersion(&_ProtocolRegistryStorage.CallOpts)
}

// MajorVersion is a free data retrieval call binding the contract method 0xaf05a5c5.
//
// Solidity: function majorVersion() view returns(string)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageCallerSession) MajorVersion() (string, error) {
	return _ProtocolRegistryStorage.Contract.MajorVersion(&_ProtocolRegistryStorage.CallOpts)
}

// TotalDeployments is a free data retrieval call binding the contract method 0xfb35b4e4.
//
// Solidity: function totalDeployments() view returns(uint256)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageCaller) TotalDeployments(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolRegistryStorage.contract.Call(opts, &out, "totalDeployments")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeployments is a free data retrieval call binding the contract method 0xfb35b4e4.
//
// Solidity: function totalDeployments() view returns(uint256)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageSession) TotalDeployments() (*big.Int, error) {
	return _ProtocolRegistryStorage.Contract.TotalDeployments(&_ProtocolRegistryStorage.CallOpts)
}

// TotalDeployments is a free data retrieval call binding the contract method 0xfb35b4e4.
//
// Solidity: function totalDeployments() view returns(uint256)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageCallerSession) TotalDeployments() (*big.Int, error) {
	return _ProtocolRegistryStorage.Contract.TotalDeployments(&_ProtocolRegistryStorage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProtocolRegistryStorage.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageSession) Version() (string, error) {
	return _ProtocolRegistryStorage.Contract.Version(&_ProtocolRegistryStorage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageCallerSession) Version() (string, error) {
	return _ProtocolRegistryStorage.Contract.Version(&_ProtocolRegistryStorage.CallOpts)
}

// Configure is a paid mutator transaction binding the contract method 0x0b6cc4f9.
//
// Solidity: function configure(string name, (bool,bool) config) returns()
func (_ProtocolRegistryStorage *ProtocolRegistryStorageTransactor) Configure(opts *bind.TransactOpts, name string, config IProtocolRegistryTypesDeploymentConfig) (*types.Transaction, error) {
	return _ProtocolRegistryStorage.contract.Transact(opts, "configure", name, config)
}

// Configure is a paid mutator transaction binding the contract method 0x0b6cc4f9.
//
// Solidity: function configure(string name, (bool,bool) config) returns()
func (_ProtocolRegistryStorage *ProtocolRegistryStorageSession) Configure(name string, config IProtocolRegistryTypesDeploymentConfig) (*types.Transaction, error) {
	return _ProtocolRegistryStorage.Contract.Configure(&_ProtocolRegistryStorage.TransactOpts, name, config)
}

// Configure is a paid mutator transaction binding the contract method 0x0b6cc4f9.
//
// Solidity: function configure(string name, (bool,bool) config) returns()
func (_ProtocolRegistryStorage *ProtocolRegistryStorageTransactorSession) Configure(name string, config IProtocolRegistryTypesDeploymentConfig) (*types.Transaction, error) {
	return _ProtocolRegistryStorage.Contract.Configure(&_ProtocolRegistryStorage.TransactOpts, name, config)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialAdmin, address pauserMultisig) returns()
func (_ProtocolRegistryStorage *ProtocolRegistryStorageTransactor) Initialize(opts *bind.TransactOpts, initialAdmin common.Address, pauserMultisig common.Address) (*types.Transaction, error) {
	return _ProtocolRegistryStorage.contract.Transact(opts, "initialize", initialAdmin, pauserMultisig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialAdmin, address pauserMultisig) returns()
func (_ProtocolRegistryStorage *ProtocolRegistryStorageSession) Initialize(initialAdmin common.Address, pauserMultisig common.Address) (*types.Transaction, error) {
	return _ProtocolRegistryStorage.Contract.Initialize(&_ProtocolRegistryStorage.TransactOpts, initialAdmin, pauserMultisig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialAdmin, address pauserMultisig) returns()
func (_ProtocolRegistryStorage *ProtocolRegistryStorageTransactorSession) Initialize(initialAdmin common.Address, pauserMultisig common.Address) (*types.Transaction, error) {
	return _ProtocolRegistryStorage.Contract.Initialize(&_ProtocolRegistryStorage.TransactOpts, initialAdmin, pauserMultisig)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ProtocolRegistryStorage *ProtocolRegistryStorageTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolRegistryStorage.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ProtocolRegistryStorage *ProtocolRegistryStorageSession) PauseAll() (*types.Transaction, error) {
	return _ProtocolRegistryStorage.Contract.PauseAll(&_ProtocolRegistryStorage.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ProtocolRegistryStorage *ProtocolRegistryStorageTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ProtocolRegistryStorage.Contract.PauseAll(&_ProtocolRegistryStorage.TransactOpts)
}

// Ship is a paid mutator transaction binding the contract method 0xfc434a8a.
//
// Solidity: function ship(address[] addresses, (bool,bool)[] configs, string[] contractNames, string semanticVersion) returns()
func (_ProtocolRegistryStorage *ProtocolRegistryStorageTransactor) Ship(opts *bind.TransactOpts, addresses []common.Address, configs []IProtocolRegistryTypesDeploymentConfig, contractNames []string, semanticVersion string) (*types.Transaction, error) {
	return _ProtocolRegistryStorage.contract.Transact(opts, "ship", addresses, configs, contractNames, semanticVersion)
}

// Ship is a paid mutator transaction binding the contract method 0xfc434a8a.
//
// Solidity: function ship(address[] addresses, (bool,bool)[] configs, string[] contractNames, string semanticVersion) returns()
func (_ProtocolRegistryStorage *ProtocolRegistryStorageSession) Ship(addresses []common.Address, configs []IProtocolRegistryTypesDeploymentConfig, contractNames []string, semanticVersion string) (*types.Transaction, error) {
	return _ProtocolRegistryStorage.Contract.Ship(&_ProtocolRegistryStorage.TransactOpts, addresses, configs, contractNames, semanticVersion)
}

// Ship is a paid mutator transaction binding the contract method 0xfc434a8a.
//
// Solidity: function ship(address[] addresses, (bool,bool)[] configs, string[] contractNames, string semanticVersion) returns()
func (_ProtocolRegistryStorage *ProtocolRegistryStorageTransactorSession) Ship(addresses []common.Address, configs []IProtocolRegistryTypesDeploymentConfig, contractNames []string, semanticVersion string) (*types.Transaction, error) {
	return _ProtocolRegistryStorage.Contract.Ship(&_ProtocolRegistryStorage.TransactOpts, addresses, configs, contractNames, semanticVersion)
}

// ProtocolRegistryStorageDeploymentConfigDeletedIterator is returned from FilterDeploymentConfigDeleted and is used to iterate over the raw logs and unpacked data for DeploymentConfigDeleted events raised by the ProtocolRegistryStorage contract.
type ProtocolRegistryStorageDeploymentConfigDeletedIterator struct {
	Event *ProtocolRegistryStorageDeploymentConfigDeleted // Event containing the contract specifics and raw log

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
func (it *ProtocolRegistryStorageDeploymentConfigDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRegistryStorageDeploymentConfigDeleted)
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
		it.Event = new(ProtocolRegistryStorageDeploymentConfigDeleted)
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
func (it *ProtocolRegistryStorageDeploymentConfigDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRegistryStorageDeploymentConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRegistryStorageDeploymentConfigDeleted represents a DeploymentConfigDeleted event raised by the ProtocolRegistryStorage contract.
type ProtocolRegistryStorageDeploymentConfigDeleted struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeploymentConfigDeleted is a free log retrieval operation binding the contract event 0xa69cccaa8b056f2577aa7e06e1eb14ae0eb526356819b9403f5b31f41f3bc509.
//
// Solidity: event DeploymentConfigDeleted(address indexed addr)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageFilterer) FilterDeploymentConfigDeleted(opts *bind.FilterOpts, addr []common.Address) (*ProtocolRegistryStorageDeploymentConfigDeletedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ProtocolRegistryStorage.contract.FilterLogs(opts, "DeploymentConfigDeleted", addrRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryStorageDeploymentConfigDeletedIterator{contract: _ProtocolRegistryStorage.contract, event: "DeploymentConfigDeleted", logs: logs, sub: sub}, nil
}

// WatchDeploymentConfigDeleted is a free log subscription operation binding the contract event 0xa69cccaa8b056f2577aa7e06e1eb14ae0eb526356819b9403f5b31f41f3bc509.
//
// Solidity: event DeploymentConfigDeleted(address indexed addr)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageFilterer) WatchDeploymentConfigDeleted(opts *bind.WatchOpts, sink chan<- *ProtocolRegistryStorageDeploymentConfigDeleted, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ProtocolRegistryStorage.contract.WatchLogs(opts, "DeploymentConfigDeleted", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRegistryStorageDeploymentConfigDeleted)
				if err := _ProtocolRegistryStorage.contract.UnpackLog(event, "DeploymentConfigDeleted", log); err != nil {
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

// ParseDeploymentConfigDeleted is a log parse operation binding the contract event 0xa69cccaa8b056f2577aa7e06e1eb14ae0eb526356819b9403f5b31f41f3bc509.
//
// Solidity: event DeploymentConfigDeleted(address indexed addr)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageFilterer) ParseDeploymentConfigDeleted(log types.Log) (*ProtocolRegistryStorageDeploymentConfigDeleted, error) {
	event := new(ProtocolRegistryStorageDeploymentConfigDeleted)
	if err := _ProtocolRegistryStorage.contract.UnpackLog(event, "DeploymentConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRegistryStorageDeploymentConfiguredIterator is returned from FilterDeploymentConfigured and is used to iterate over the raw logs and unpacked data for DeploymentConfigured events raised by the ProtocolRegistryStorage contract.
type ProtocolRegistryStorageDeploymentConfiguredIterator struct {
	Event *ProtocolRegistryStorageDeploymentConfigured // Event containing the contract specifics and raw log

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
func (it *ProtocolRegistryStorageDeploymentConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRegistryStorageDeploymentConfigured)
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
		it.Event = new(ProtocolRegistryStorageDeploymentConfigured)
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
func (it *ProtocolRegistryStorageDeploymentConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRegistryStorageDeploymentConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRegistryStorageDeploymentConfigured represents a DeploymentConfigured event raised by the ProtocolRegistryStorage contract.
type ProtocolRegistryStorageDeploymentConfigured struct {
	Addr   common.Address
	Config IProtocolRegistryTypesDeploymentConfig
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeploymentConfigured is a free log retrieval operation binding the contract event 0x6c5879ec82f910f6d12844857cfb8eb474dcecc9aa5b8257c7a77dcb42990e9d.
//
// Solidity: event DeploymentConfigured(address indexed addr, (bool,bool) config)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageFilterer) FilterDeploymentConfigured(opts *bind.FilterOpts, addr []common.Address) (*ProtocolRegistryStorageDeploymentConfiguredIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ProtocolRegistryStorage.contract.FilterLogs(opts, "DeploymentConfigured", addrRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryStorageDeploymentConfiguredIterator{contract: _ProtocolRegistryStorage.contract, event: "DeploymentConfigured", logs: logs, sub: sub}, nil
}

// WatchDeploymentConfigured is a free log subscription operation binding the contract event 0x6c5879ec82f910f6d12844857cfb8eb474dcecc9aa5b8257c7a77dcb42990e9d.
//
// Solidity: event DeploymentConfigured(address indexed addr, (bool,bool) config)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageFilterer) WatchDeploymentConfigured(opts *bind.WatchOpts, sink chan<- *ProtocolRegistryStorageDeploymentConfigured, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ProtocolRegistryStorage.contract.WatchLogs(opts, "DeploymentConfigured", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRegistryStorageDeploymentConfigured)
				if err := _ProtocolRegistryStorage.contract.UnpackLog(event, "DeploymentConfigured", log); err != nil {
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

// ParseDeploymentConfigured is a log parse operation binding the contract event 0x6c5879ec82f910f6d12844857cfb8eb474dcecc9aa5b8257c7a77dcb42990e9d.
//
// Solidity: event DeploymentConfigured(address indexed addr, (bool,bool) config)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageFilterer) ParseDeploymentConfigured(log types.Log) (*ProtocolRegistryStorageDeploymentConfigured, error) {
	event := new(ProtocolRegistryStorageDeploymentConfigured)
	if err := _ProtocolRegistryStorage.contract.UnpackLog(event, "DeploymentConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRegistryStorageDeploymentShippedIterator is returned from FilterDeploymentShipped and is used to iterate over the raw logs and unpacked data for DeploymentShipped events raised by the ProtocolRegistryStorage contract.
type ProtocolRegistryStorageDeploymentShippedIterator struct {
	Event *ProtocolRegistryStorageDeploymentShipped // Event containing the contract specifics and raw log

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
func (it *ProtocolRegistryStorageDeploymentShippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRegistryStorageDeploymentShipped)
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
		it.Event = new(ProtocolRegistryStorageDeploymentShipped)
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
func (it *ProtocolRegistryStorageDeploymentShippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRegistryStorageDeploymentShippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRegistryStorageDeploymentShipped represents a DeploymentShipped event raised by the ProtocolRegistryStorage contract.
type ProtocolRegistryStorageDeploymentShipped struct {
	Addr   common.Address
	Config IProtocolRegistryTypesDeploymentConfig
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeploymentShipped is a free log retrieval operation binding the contract event 0x4e81e5495dbc85919bf1b9c12a9a4bb6d546f75d3f92fa0a36867fc29a51467c.
//
// Solidity: event DeploymentShipped(address indexed addr, (bool,bool) config)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageFilterer) FilterDeploymentShipped(opts *bind.FilterOpts, addr []common.Address) (*ProtocolRegistryStorageDeploymentShippedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ProtocolRegistryStorage.contract.FilterLogs(opts, "DeploymentShipped", addrRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryStorageDeploymentShippedIterator{contract: _ProtocolRegistryStorage.contract, event: "DeploymentShipped", logs: logs, sub: sub}, nil
}

// WatchDeploymentShipped is a free log subscription operation binding the contract event 0x4e81e5495dbc85919bf1b9c12a9a4bb6d546f75d3f92fa0a36867fc29a51467c.
//
// Solidity: event DeploymentShipped(address indexed addr, (bool,bool) config)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageFilterer) WatchDeploymentShipped(opts *bind.WatchOpts, sink chan<- *ProtocolRegistryStorageDeploymentShipped, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _ProtocolRegistryStorage.contract.WatchLogs(opts, "DeploymentShipped", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRegistryStorageDeploymentShipped)
				if err := _ProtocolRegistryStorage.contract.UnpackLog(event, "DeploymentShipped", log); err != nil {
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

// ParseDeploymentShipped is a log parse operation binding the contract event 0x4e81e5495dbc85919bf1b9c12a9a4bb6d546f75d3f92fa0a36867fc29a51467c.
//
// Solidity: event DeploymentShipped(address indexed addr, (bool,bool) config)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageFilterer) ParseDeploymentShipped(log types.Log) (*ProtocolRegistryStorageDeploymentShipped, error) {
	event := new(ProtocolRegistryStorageDeploymentShipped)
	if err := _ProtocolRegistryStorage.contract.UnpackLog(event, "DeploymentShipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRegistryStorageSemanticVersionUpdatedIterator is returned from FilterSemanticVersionUpdated and is used to iterate over the raw logs and unpacked data for SemanticVersionUpdated events raised by the ProtocolRegistryStorage contract.
type ProtocolRegistryStorageSemanticVersionUpdatedIterator struct {
	Event *ProtocolRegistryStorageSemanticVersionUpdated // Event containing the contract specifics and raw log

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
func (it *ProtocolRegistryStorageSemanticVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRegistryStorageSemanticVersionUpdated)
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
		it.Event = new(ProtocolRegistryStorageSemanticVersionUpdated)
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
func (it *ProtocolRegistryStorageSemanticVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRegistryStorageSemanticVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRegistryStorageSemanticVersionUpdated represents a SemanticVersionUpdated event raised by the ProtocolRegistryStorage contract.
type ProtocolRegistryStorageSemanticVersionUpdated struct {
	PreviousSemanticVersion string
	SemanticVersion         string
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSemanticVersionUpdated is a free log retrieval operation binding the contract event 0x769656e5cb3218f54758f85bd184d41e987639374d6cb9e244439cc9d1abe1e2.
//
// Solidity: event SemanticVersionUpdated(string previousSemanticVersion, string semanticVersion)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageFilterer) FilterSemanticVersionUpdated(opts *bind.FilterOpts) (*ProtocolRegistryStorageSemanticVersionUpdatedIterator, error) {

	logs, sub, err := _ProtocolRegistryStorage.contract.FilterLogs(opts, "SemanticVersionUpdated")
	if err != nil {
		return nil, err
	}
	return &ProtocolRegistryStorageSemanticVersionUpdatedIterator{contract: _ProtocolRegistryStorage.contract, event: "SemanticVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchSemanticVersionUpdated is a free log subscription operation binding the contract event 0x769656e5cb3218f54758f85bd184d41e987639374d6cb9e244439cc9d1abe1e2.
//
// Solidity: event SemanticVersionUpdated(string previousSemanticVersion, string semanticVersion)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageFilterer) WatchSemanticVersionUpdated(opts *bind.WatchOpts, sink chan<- *ProtocolRegistryStorageSemanticVersionUpdated) (event.Subscription, error) {

	logs, sub, err := _ProtocolRegistryStorage.contract.WatchLogs(opts, "SemanticVersionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRegistryStorageSemanticVersionUpdated)
				if err := _ProtocolRegistryStorage.contract.UnpackLog(event, "SemanticVersionUpdated", log); err != nil {
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

// ParseSemanticVersionUpdated is a log parse operation binding the contract event 0x769656e5cb3218f54758f85bd184d41e987639374d6cb9e244439cc9d1abe1e2.
//
// Solidity: event SemanticVersionUpdated(string previousSemanticVersion, string semanticVersion)
func (_ProtocolRegistryStorage *ProtocolRegistryStorageFilterer) ParseSemanticVersionUpdated(log types.Log) (*ProtocolRegistryStorageSemanticVersionUpdated, error) {
	event := new(ProtocolRegistryStorageSemanticVersionUpdated)
	if err := _ProtocolRegistryStorage.contract.UnpackLog(event, "SemanticVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
