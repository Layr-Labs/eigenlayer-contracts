// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IProtocolRegistry

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

// IProtocolRegistryMetaData contains all meta data concerning the IProtocolRegistry contract.
var IProtocolRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"configure\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAddress\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllDeployments\",\"inputs\":[],\"outputs\":[{\"name\":\"names\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"configs\",\"type\":\"tuple[]\",\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig[]\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeployment\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pauserMultisig\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"majorVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ship\",\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"configs\",\"type\":\"tuple[]\",\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig[]\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"contractNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"semanticVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalDeployments\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DeploymentConfigured\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeploymentShipped\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIProtocolRegistryTypes.DeploymentConfig\",\"components\":[{\"name\":\"pausable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"deprecated\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SemanticVersionUpdated\",\"inputs\":[{\"name\":\"previousSemanticVersion\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"semanticVersion\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]",
}

// IProtocolRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IProtocolRegistryMetaData.ABI instead.
var IProtocolRegistryABI = IProtocolRegistryMetaData.ABI

// IProtocolRegistry is an auto generated Go binding around an Ethereum contract.
type IProtocolRegistry struct {
	IProtocolRegistryCaller     // Read-only binding to the contract
	IProtocolRegistryTransactor // Write-only binding to the contract
	IProtocolRegistryFilterer   // Log filterer for contract events
}

// IProtocolRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IProtocolRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProtocolRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IProtocolRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProtocolRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IProtocolRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProtocolRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IProtocolRegistrySession struct {
	Contract     *IProtocolRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IProtocolRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IProtocolRegistryCallerSession struct {
	Contract *IProtocolRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IProtocolRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IProtocolRegistryTransactorSession struct {
	Contract     *IProtocolRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IProtocolRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IProtocolRegistryRaw struct {
	Contract *IProtocolRegistry // Generic contract binding to access the raw methods on
}

// IProtocolRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IProtocolRegistryCallerRaw struct {
	Contract *IProtocolRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IProtocolRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IProtocolRegistryTransactorRaw struct {
	Contract *IProtocolRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIProtocolRegistry creates a new instance of IProtocolRegistry, bound to a specific deployed contract.
func NewIProtocolRegistry(address common.Address, backend bind.ContractBackend) (*IProtocolRegistry, error) {
	contract, err := bindIProtocolRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IProtocolRegistry{IProtocolRegistryCaller: IProtocolRegistryCaller{contract: contract}, IProtocolRegistryTransactor: IProtocolRegistryTransactor{contract: contract}, IProtocolRegistryFilterer: IProtocolRegistryFilterer{contract: contract}}, nil
}

// NewIProtocolRegistryCaller creates a new read-only instance of IProtocolRegistry, bound to a specific deployed contract.
func NewIProtocolRegistryCaller(address common.Address, caller bind.ContractCaller) (*IProtocolRegistryCaller, error) {
	contract, err := bindIProtocolRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IProtocolRegistryCaller{contract: contract}, nil
}

// NewIProtocolRegistryTransactor creates a new write-only instance of IProtocolRegistry, bound to a specific deployed contract.
func NewIProtocolRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IProtocolRegistryTransactor, error) {
	contract, err := bindIProtocolRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IProtocolRegistryTransactor{contract: contract}, nil
}

// NewIProtocolRegistryFilterer creates a new log filterer instance of IProtocolRegistry, bound to a specific deployed contract.
func NewIProtocolRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IProtocolRegistryFilterer, error) {
	contract, err := bindIProtocolRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IProtocolRegistryFilterer{contract: contract}, nil
}

// bindIProtocolRegistry binds a generic wrapper to an already deployed contract.
func bindIProtocolRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IProtocolRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProtocolRegistry *IProtocolRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProtocolRegistry.Contract.IProtocolRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProtocolRegistry *IProtocolRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProtocolRegistry.Contract.IProtocolRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProtocolRegistry *IProtocolRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProtocolRegistry.Contract.IProtocolRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProtocolRegistry *IProtocolRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProtocolRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProtocolRegistry *IProtocolRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProtocolRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProtocolRegistry *IProtocolRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProtocolRegistry.Contract.contract.Transact(opts, method, params...)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_IProtocolRegistry *IProtocolRegistryCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IProtocolRegistry.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_IProtocolRegistry *IProtocolRegistrySession) PAUSERROLE() ([32]byte, error) {
	return _IProtocolRegistry.Contract.PAUSERROLE(&_IProtocolRegistry.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_IProtocolRegistry *IProtocolRegistryCallerSession) PAUSERROLE() ([32]byte, error) {
	return _IProtocolRegistry.Contract.PAUSERROLE(&_IProtocolRegistry.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string name) view returns(address)
func (_IProtocolRegistry *IProtocolRegistryCaller) GetAddress(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _IProtocolRegistry.contract.Call(opts, &out, "getAddress", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string name) view returns(address)
func (_IProtocolRegistry *IProtocolRegistrySession) GetAddress(name string) (common.Address, error) {
	return _IProtocolRegistry.Contract.GetAddress(&_IProtocolRegistry.CallOpts, name)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string name) view returns(address)
func (_IProtocolRegistry *IProtocolRegistryCallerSession) GetAddress(name string) (common.Address, error) {
	return _IProtocolRegistry.Contract.GetAddress(&_IProtocolRegistry.CallOpts, name)
}

// GetAllDeployments is a free data retrieval call binding the contract method 0x8eec00b8.
//
// Solidity: function getAllDeployments() view returns(string[] names, address[] addresses, (bool,bool)[] configs)
func (_IProtocolRegistry *IProtocolRegistryCaller) GetAllDeployments(opts *bind.CallOpts) (struct {
	Names     []string
	Addresses []common.Address
	Configs   []IProtocolRegistryTypesDeploymentConfig
}, error) {
	var out []interface{}
	err := _IProtocolRegistry.contract.Call(opts, &out, "getAllDeployments")

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
func (_IProtocolRegistry *IProtocolRegistrySession) GetAllDeployments() (struct {
	Names     []string
	Addresses []common.Address
	Configs   []IProtocolRegistryTypesDeploymentConfig
}, error) {
	return _IProtocolRegistry.Contract.GetAllDeployments(&_IProtocolRegistry.CallOpts)
}

// GetAllDeployments is a free data retrieval call binding the contract method 0x8eec00b8.
//
// Solidity: function getAllDeployments() view returns(string[] names, address[] addresses, (bool,bool)[] configs)
func (_IProtocolRegistry *IProtocolRegistryCallerSession) GetAllDeployments() (struct {
	Names     []string
	Addresses []common.Address
	Configs   []IProtocolRegistryTypesDeploymentConfig
}, error) {
	return _IProtocolRegistry.Contract.GetAllDeployments(&_IProtocolRegistry.CallOpts)
}

// GetDeployment is a free data retrieval call binding the contract method 0xa8091d97.
//
// Solidity: function getDeployment(string name) view returns(address addr, (bool,bool) config)
func (_IProtocolRegistry *IProtocolRegistryCaller) GetDeployment(opts *bind.CallOpts, name string) (struct {
	Addr   common.Address
	Config IProtocolRegistryTypesDeploymentConfig
}, error) {
	var out []interface{}
	err := _IProtocolRegistry.contract.Call(opts, &out, "getDeployment", name)

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
func (_IProtocolRegistry *IProtocolRegistrySession) GetDeployment(name string) (struct {
	Addr   common.Address
	Config IProtocolRegistryTypesDeploymentConfig
}, error) {
	return _IProtocolRegistry.Contract.GetDeployment(&_IProtocolRegistry.CallOpts, name)
}

// GetDeployment is a free data retrieval call binding the contract method 0xa8091d97.
//
// Solidity: function getDeployment(string name) view returns(address addr, (bool,bool) config)
func (_IProtocolRegistry *IProtocolRegistryCallerSession) GetDeployment(name string) (struct {
	Addr   common.Address
	Config IProtocolRegistryTypesDeploymentConfig
}, error) {
	return _IProtocolRegistry.Contract.GetDeployment(&_IProtocolRegistry.CallOpts, name)
}

// MajorVersion is a free data retrieval call binding the contract method 0xaf05a5c5.
//
// Solidity: function majorVersion() view returns(string)
func (_IProtocolRegistry *IProtocolRegistryCaller) MajorVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IProtocolRegistry.contract.Call(opts, &out, "majorVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MajorVersion is a free data retrieval call binding the contract method 0xaf05a5c5.
//
// Solidity: function majorVersion() view returns(string)
func (_IProtocolRegistry *IProtocolRegistrySession) MajorVersion() (string, error) {
	return _IProtocolRegistry.Contract.MajorVersion(&_IProtocolRegistry.CallOpts)
}

// MajorVersion is a free data retrieval call binding the contract method 0xaf05a5c5.
//
// Solidity: function majorVersion() view returns(string)
func (_IProtocolRegistry *IProtocolRegistryCallerSession) MajorVersion() (string, error) {
	return _IProtocolRegistry.Contract.MajorVersion(&_IProtocolRegistry.CallOpts)
}

// TotalDeployments is a free data retrieval call binding the contract method 0xfb35b4e4.
//
// Solidity: function totalDeployments() view returns(uint256)
func (_IProtocolRegistry *IProtocolRegistryCaller) TotalDeployments(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IProtocolRegistry.contract.Call(opts, &out, "totalDeployments")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeployments is a free data retrieval call binding the contract method 0xfb35b4e4.
//
// Solidity: function totalDeployments() view returns(uint256)
func (_IProtocolRegistry *IProtocolRegistrySession) TotalDeployments() (*big.Int, error) {
	return _IProtocolRegistry.Contract.TotalDeployments(&_IProtocolRegistry.CallOpts)
}

// TotalDeployments is a free data retrieval call binding the contract method 0xfb35b4e4.
//
// Solidity: function totalDeployments() view returns(uint256)
func (_IProtocolRegistry *IProtocolRegistryCallerSession) TotalDeployments() (*big.Int, error) {
	return _IProtocolRegistry.Contract.TotalDeployments(&_IProtocolRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IProtocolRegistry *IProtocolRegistryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IProtocolRegistry.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IProtocolRegistry *IProtocolRegistrySession) Version() (string, error) {
	return _IProtocolRegistry.Contract.Version(&_IProtocolRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IProtocolRegistry *IProtocolRegistryCallerSession) Version() (string, error) {
	return _IProtocolRegistry.Contract.Version(&_IProtocolRegistry.CallOpts)
}

// Configure is a paid mutator transaction binding the contract method 0xd3466389.
//
// Solidity: function configure(address addr, (bool,bool) config) returns()
func (_IProtocolRegistry *IProtocolRegistryTransactor) Configure(opts *bind.TransactOpts, addr common.Address, config IProtocolRegistryTypesDeploymentConfig) (*types.Transaction, error) {
	return _IProtocolRegistry.contract.Transact(opts, "configure", addr, config)
}

// Configure is a paid mutator transaction binding the contract method 0xd3466389.
//
// Solidity: function configure(address addr, (bool,bool) config) returns()
func (_IProtocolRegistry *IProtocolRegistrySession) Configure(addr common.Address, config IProtocolRegistryTypesDeploymentConfig) (*types.Transaction, error) {
	return _IProtocolRegistry.Contract.Configure(&_IProtocolRegistry.TransactOpts, addr, config)
}

// Configure is a paid mutator transaction binding the contract method 0xd3466389.
//
// Solidity: function configure(address addr, (bool,bool) config) returns()
func (_IProtocolRegistry *IProtocolRegistryTransactorSession) Configure(addr common.Address, config IProtocolRegistryTypesDeploymentConfig) (*types.Transaction, error) {
	return _IProtocolRegistry.Contract.Configure(&_IProtocolRegistry.TransactOpts, addr, config)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialAdmin, address pauserMultisig) returns()
func (_IProtocolRegistry *IProtocolRegistryTransactor) Initialize(opts *bind.TransactOpts, initialAdmin common.Address, pauserMultisig common.Address) (*types.Transaction, error) {
	return _IProtocolRegistry.contract.Transact(opts, "initialize", initialAdmin, pauserMultisig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialAdmin, address pauserMultisig) returns()
func (_IProtocolRegistry *IProtocolRegistrySession) Initialize(initialAdmin common.Address, pauserMultisig common.Address) (*types.Transaction, error) {
	return _IProtocolRegistry.Contract.Initialize(&_IProtocolRegistry.TransactOpts, initialAdmin, pauserMultisig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialAdmin, address pauserMultisig) returns()
func (_IProtocolRegistry *IProtocolRegistryTransactorSession) Initialize(initialAdmin common.Address, pauserMultisig common.Address) (*types.Transaction, error) {
	return _IProtocolRegistry.Contract.Initialize(&_IProtocolRegistry.TransactOpts, initialAdmin, pauserMultisig)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_IProtocolRegistry *IProtocolRegistryTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProtocolRegistry.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_IProtocolRegistry *IProtocolRegistrySession) PauseAll() (*types.Transaction, error) {
	return _IProtocolRegistry.Contract.PauseAll(&_IProtocolRegistry.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_IProtocolRegistry *IProtocolRegistryTransactorSession) PauseAll() (*types.Transaction, error) {
	return _IProtocolRegistry.Contract.PauseAll(&_IProtocolRegistry.TransactOpts)
}

// Ship is a paid mutator transaction binding the contract method 0xfc434a8a.
//
// Solidity: function ship(address[] addresses, (bool,bool)[] configs, string[] contractNames, string semanticVersion) returns()
func (_IProtocolRegistry *IProtocolRegistryTransactor) Ship(opts *bind.TransactOpts, addresses []common.Address, configs []IProtocolRegistryTypesDeploymentConfig, contractNames []string, semanticVersion string) (*types.Transaction, error) {
	return _IProtocolRegistry.contract.Transact(opts, "ship", addresses, configs, contractNames, semanticVersion)
}

// Ship is a paid mutator transaction binding the contract method 0xfc434a8a.
//
// Solidity: function ship(address[] addresses, (bool,bool)[] configs, string[] contractNames, string semanticVersion) returns()
func (_IProtocolRegistry *IProtocolRegistrySession) Ship(addresses []common.Address, configs []IProtocolRegistryTypesDeploymentConfig, contractNames []string, semanticVersion string) (*types.Transaction, error) {
	return _IProtocolRegistry.Contract.Ship(&_IProtocolRegistry.TransactOpts, addresses, configs, contractNames, semanticVersion)
}

// Ship is a paid mutator transaction binding the contract method 0xfc434a8a.
//
// Solidity: function ship(address[] addresses, (bool,bool)[] configs, string[] contractNames, string semanticVersion) returns()
func (_IProtocolRegistry *IProtocolRegistryTransactorSession) Ship(addresses []common.Address, configs []IProtocolRegistryTypesDeploymentConfig, contractNames []string, semanticVersion string) (*types.Transaction, error) {
	return _IProtocolRegistry.Contract.Ship(&_IProtocolRegistry.TransactOpts, addresses, configs, contractNames, semanticVersion)
}

// IProtocolRegistryDeploymentConfiguredIterator is returned from FilterDeploymentConfigured and is used to iterate over the raw logs and unpacked data for DeploymentConfigured events raised by the IProtocolRegistry contract.
type IProtocolRegistryDeploymentConfiguredIterator struct {
	Event *IProtocolRegistryDeploymentConfigured // Event containing the contract specifics and raw log

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
func (it *IProtocolRegistryDeploymentConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProtocolRegistryDeploymentConfigured)
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
		it.Event = new(IProtocolRegistryDeploymentConfigured)
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
func (it *IProtocolRegistryDeploymentConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProtocolRegistryDeploymentConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProtocolRegistryDeploymentConfigured represents a DeploymentConfigured event raised by the IProtocolRegistry contract.
type IProtocolRegistryDeploymentConfigured struct {
	Addr   common.Address
	Config IProtocolRegistryTypesDeploymentConfig
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeploymentConfigured is a free log retrieval operation binding the contract event 0x6c5879ec82f910f6d12844857cfb8eb474dcecc9aa5b8257c7a77dcb42990e9d.
//
// Solidity: event DeploymentConfigured(address indexed addr, (bool,bool) config)
func (_IProtocolRegistry *IProtocolRegistryFilterer) FilterDeploymentConfigured(opts *bind.FilterOpts, addr []common.Address) (*IProtocolRegistryDeploymentConfiguredIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _IProtocolRegistry.contract.FilterLogs(opts, "DeploymentConfigured", addrRule)
	if err != nil {
		return nil, err
	}
	return &IProtocolRegistryDeploymentConfiguredIterator{contract: _IProtocolRegistry.contract, event: "DeploymentConfigured", logs: logs, sub: sub}, nil
}

// WatchDeploymentConfigured is a free log subscription operation binding the contract event 0x6c5879ec82f910f6d12844857cfb8eb474dcecc9aa5b8257c7a77dcb42990e9d.
//
// Solidity: event DeploymentConfigured(address indexed addr, (bool,bool) config)
func (_IProtocolRegistry *IProtocolRegistryFilterer) WatchDeploymentConfigured(opts *bind.WatchOpts, sink chan<- *IProtocolRegistryDeploymentConfigured, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _IProtocolRegistry.contract.WatchLogs(opts, "DeploymentConfigured", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProtocolRegistryDeploymentConfigured)
				if err := _IProtocolRegistry.contract.UnpackLog(event, "DeploymentConfigured", log); err != nil {
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
func (_IProtocolRegistry *IProtocolRegistryFilterer) ParseDeploymentConfigured(log types.Log) (*IProtocolRegistryDeploymentConfigured, error) {
	event := new(IProtocolRegistryDeploymentConfigured)
	if err := _IProtocolRegistry.contract.UnpackLog(event, "DeploymentConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProtocolRegistryDeploymentShippedIterator is returned from FilterDeploymentShipped and is used to iterate over the raw logs and unpacked data for DeploymentShipped events raised by the IProtocolRegistry contract.
type IProtocolRegistryDeploymentShippedIterator struct {
	Event *IProtocolRegistryDeploymentShipped // Event containing the contract specifics and raw log

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
func (it *IProtocolRegistryDeploymentShippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProtocolRegistryDeploymentShipped)
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
		it.Event = new(IProtocolRegistryDeploymentShipped)
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
func (it *IProtocolRegistryDeploymentShippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProtocolRegistryDeploymentShippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProtocolRegistryDeploymentShipped represents a DeploymentShipped event raised by the IProtocolRegistry contract.
type IProtocolRegistryDeploymentShipped struct {
	Addr   common.Address
	Config IProtocolRegistryTypesDeploymentConfig
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeploymentShipped is a free log retrieval operation binding the contract event 0x4e81e5495dbc85919bf1b9c12a9a4bb6d546f75d3f92fa0a36867fc29a51467c.
//
// Solidity: event DeploymentShipped(address indexed addr, (bool,bool) config)
func (_IProtocolRegistry *IProtocolRegistryFilterer) FilterDeploymentShipped(opts *bind.FilterOpts, addr []common.Address) (*IProtocolRegistryDeploymentShippedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _IProtocolRegistry.contract.FilterLogs(opts, "DeploymentShipped", addrRule)
	if err != nil {
		return nil, err
	}
	return &IProtocolRegistryDeploymentShippedIterator{contract: _IProtocolRegistry.contract, event: "DeploymentShipped", logs: logs, sub: sub}, nil
}

// WatchDeploymentShipped is a free log subscription operation binding the contract event 0x4e81e5495dbc85919bf1b9c12a9a4bb6d546f75d3f92fa0a36867fc29a51467c.
//
// Solidity: event DeploymentShipped(address indexed addr, (bool,bool) config)
func (_IProtocolRegistry *IProtocolRegistryFilterer) WatchDeploymentShipped(opts *bind.WatchOpts, sink chan<- *IProtocolRegistryDeploymentShipped, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _IProtocolRegistry.contract.WatchLogs(opts, "DeploymentShipped", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProtocolRegistryDeploymentShipped)
				if err := _IProtocolRegistry.contract.UnpackLog(event, "DeploymentShipped", log); err != nil {
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
func (_IProtocolRegistry *IProtocolRegistryFilterer) ParseDeploymentShipped(log types.Log) (*IProtocolRegistryDeploymentShipped, error) {
	event := new(IProtocolRegistryDeploymentShipped)
	if err := _IProtocolRegistry.contract.UnpackLog(event, "DeploymentShipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IProtocolRegistrySemanticVersionUpdatedIterator is returned from FilterSemanticVersionUpdated and is used to iterate over the raw logs and unpacked data for SemanticVersionUpdated events raised by the IProtocolRegistry contract.
type IProtocolRegistrySemanticVersionUpdatedIterator struct {
	Event *IProtocolRegistrySemanticVersionUpdated // Event containing the contract specifics and raw log

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
func (it *IProtocolRegistrySemanticVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IProtocolRegistrySemanticVersionUpdated)
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
		it.Event = new(IProtocolRegistrySemanticVersionUpdated)
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
func (it *IProtocolRegistrySemanticVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IProtocolRegistrySemanticVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IProtocolRegistrySemanticVersionUpdated represents a SemanticVersionUpdated event raised by the IProtocolRegistry contract.
type IProtocolRegistrySemanticVersionUpdated struct {
	PreviousSemanticVersion string
	SemanticVersion         string
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSemanticVersionUpdated is a free log retrieval operation binding the contract event 0x769656e5cb3218f54758f85bd184d41e987639374d6cb9e244439cc9d1abe1e2.
//
// Solidity: event SemanticVersionUpdated(string previousSemanticVersion, string semanticVersion)
func (_IProtocolRegistry *IProtocolRegistryFilterer) FilterSemanticVersionUpdated(opts *bind.FilterOpts) (*IProtocolRegistrySemanticVersionUpdatedIterator, error) {

	logs, sub, err := _IProtocolRegistry.contract.FilterLogs(opts, "SemanticVersionUpdated")
	if err != nil {
		return nil, err
	}
	return &IProtocolRegistrySemanticVersionUpdatedIterator{contract: _IProtocolRegistry.contract, event: "SemanticVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchSemanticVersionUpdated is a free log subscription operation binding the contract event 0x769656e5cb3218f54758f85bd184d41e987639374d6cb9e244439cc9d1abe1e2.
//
// Solidity: event SemanticVersionUpdated(string previousSemanticVersion, string semanticVersion)
func (_IProtocolRegistry *IProtocolRegistryFilterer) WatchSemanticVersionUpdated(opts *bind.WatchOpts, sink chan<- *IProtocolRegistrySemanticVersionUpdated) (event.Subscription, error) {

	logs, sub, err := _IProtocolRegistry.contract.WatchLogs(opts, "SemanticVersionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IProtocolRegistrySemanticVersionUpdated)
				if err := _IProtocolRegistry.contract.UnpackLog(event, "SemanticVersionUpdated", log); err != nil {
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
func (_IProtocolRegistry *IProtocolRegistryFilterer) ParseSemanticVersionUpdated(log types.Log) (*IProtocolRegistrySemanticVersionUpdated, error) {
	event := new(IProtocolRegistrySemanticVersionUpdated)
	if err := _IProtocolRegistry.contract.UnpackLog(event, "SemanticVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
