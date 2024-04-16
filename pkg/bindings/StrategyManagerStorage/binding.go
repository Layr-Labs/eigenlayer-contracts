// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StrategyManagerStorage

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

// StrategyManagerStorageMetaData contains all meta data concerning the StrategyManagerStorage contract.
var StrategyManagerStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEPOSIT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStrategiesToDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"thirdPartyTransfersForbiddenValues\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositIntoStrategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoStrategyWithSignature\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyList\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyListLength\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyIsWhitelistedForDeposit\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyWhitelister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"thirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyWhitelisterChanged\",\"inputs\":[{\"name\":\"previousAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedThirdPartyTransfersForbidden\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
}

// StrategyManagerStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyManagerStorageMetaData.ABI instead.
var StrategyManagerStorageABI = StrategyManagerStorageMetaData.ABI

// StrategyManagerStorage is an auto generated Go binding around an Ethereum contract.
type StrategyManagerStorage struct {
	StrategyManagerStorageCaller     // Read-only binding to the contract
	StrategyManagerStorageTransactor // Write-only binding to the contract
	StrategyManagerStorageFilterer   // Log filterer for contract events
}

// StrategyManagerStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyManagerStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyManagerStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyManagerStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyManagerStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyManagerStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyManagerStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategyManagerStorageSession struct {
	Contract     *StrategyManagerStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StrategyManagerStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyManagerStorageCallerSession struct {
	Contract *StrategyManagerStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// StrategyManagerStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyManagerStorageTransactorSession struct {
	Contract     *StrategyManagerStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// StrategyManagerStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyManagerStorageRaw struct {
	Contract *StrategyManagerStorage // Generic contract binding to access the raw methods on
}

// StrategyManagerStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyManagerStorageCallerRaw struct {
	Contract *StrategyManagerStorageCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyManagerStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyManagerStorageTransactorRaw struct {
	Contract *StrategyManagerStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategyManagerStorage creates a new instance of StrategyManagerStorage, bound to a specific deployed contract.
func NewStrategyManagerStorage(address common.Address, backend bind.ContractBackend) (*StrategyManagerStorage, error) {
	contract, err := bindStrategyManagerStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorage{StrategyManagerStorageCaller: StrategyManagerStorageCaller{contract: contract}, StrategyManagerStorageTransactor: StrategyManagerStorageTransactor{contract: contract}, StrategyManagerStorageFilterer: StrategyManagerStorageFilterer{contract: contract}}, nil
}

// NewStrategyManagerStorageCaller creates a new read-only instance of StrategyManagerStorage, bound to a specific deployed contract.
func NewStrategyManagerStorageCaller(address common.Address, caller bind.ContractCaller) (*StrategyManagerStorageCaller, error) {
	contract, err := bindStrategyManagerStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageCaller{contract: contract}, nil
}

// NewStrategyManagerStorageTransactor creates a new write-only instance of StrategyManagerStorage, bound to a specific deployed contract.
func NewStrategyManagerStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyManagerStorageTransactor, error) {
	contract, err := bindStrategyManagerStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageTransactor{contract: contract}, nil
}

// NewStrategyManagerStorageFilterer creates a new log filterer instance of StrategyManagerStorage, bound to a specific deployed contract.
func NewStrategyManagerStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyManagerStorageFilterer, error) {
	contract, err := bindStrategyManagerStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageFilterer{contract: contract}, nil
}

// bindStrategyManagerStorage binds a generic wrapper to an already deployed contract.
func bindStrategyManagerStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrategyManagerStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyManagerStorage *StrategyManagerStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyManagerStorage.Contract.StrategyManagerStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyManagerStorage *StrategyManagerStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.StrategyManagerStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyManagerStorage *StrategyManagerStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.StrategyManagerStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyManagerStorage *StrategyManagerStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyManagerStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyManagerStorage *StrategyManagerStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyManagerStorage *StrategyManagerStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) DEPOSITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "DEPOSIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_StrategyManagerStorage *StrategyManagerStorageSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _StrategyManagerStorage.Contract.DEPOSITTYPEHASH(&_StrategyManagerStorage.CallOpts)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _StrategyManagerStorage.Contract.DEPOSITTYPEHASH(&_StrategyManagerStorage.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_StrategyManagerStorage *StrategyManagerStorageSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _StrategyManagerStorage.Contract.DOMAINTYPEHASH(&_StrategyManagerStorage.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _StrategyManagerStorage.Contract.DOMAINTYPEHASH(&_StrategyManagerStorage.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageSession) Delegation() (common.Address, error) {
	return _StrategyManagerStorage.Contract.Delegation(&_StrategyManagerStorage.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) Delegation() (common.Address, error) {
	return _StrategyManagerStorage.Contract.Delegation(&_StrategyManagerStorage.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageSession) EigenPodManager() (common.Address, error) {
	return _StrategyManagerStorage.Contract.EigenPodManager(&_StrategyManagerStorage.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) EigenPodManager() (common.Address, error) {
	return _StrategyManagerStorage.Contract.EigenPodManager(&_StrategyManagerStorage.CallOpts)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_StrategyManagerStorage *StrategyManagerStorageCaller) GetDeposits(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "getDeposits", staker)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_StrategyManagerStorage *StrategyManagerStorageSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _StrategyManagerStorage.Contract.GetDeposits(&_StrategyManagerStorage.CallOpts, staker)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _StrategyManagerStorage.Contract.GetDeposits(&_StrategyManagerStorage.CallOpts, staker)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _StrategyManagerStorage.Contract.Nonces(&_StrategyManagerStorage.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _StrategyManagerStorage.Contract.Nonces(&_StrategyManagerStorage.CallOpts, arg0)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageSession) Slasher() (common.Address, error) {
	return _StrategyManagerStorage.Contract.Slasher(&_StrategyManagerStorage.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) Slasher() (common.Address, error) {
	return _StrategyManagerStorage.Contract.Slasher(&_StrategyManagerStorage.CallOpts)
}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address , uint256 ) view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) StakerStrategyList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "stakerStrategyList", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address , uint256 ) view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageSession) StakerStrategyList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _StrategyManagerStorage.Contract.StakerStrategyList(&_StrategyManagerStorage.CallOpts, arg0, arg1)
}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address , uint256 ) view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) StakerStrategyList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _StrategyManagerStorage.Contract.StakerStrategyList(&_StrategyManagerStorage.CallOpts, arg0, arg1)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) StakerStrategyListLength(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "stakerStrategyListLength", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _StrategyManagerStorage.Contract.StakerStrategyListLength(&_StrategyManagerStorage.CallOpts, staker)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _StrategyManagerStorage.Contract.StakerStrategyListLength(&_StrategyManagerStorage.CallOpts, staker)
}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address , address ) view returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) StakerStrategyShares(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "stakerStrategyShares", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address , address ) view returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageSession) StakerStrategyShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StrategyManagerStorage.Contract.StakerStrategyShares(&_StrategyManagerStorage.CallOpts, arg0, arg1)
}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address , address ) view returns(uint256)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) StakerStrategyShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StrategyManagerStorage.Contract.StakerStrategyShares(&_StrategyManagerStorage.CallOpts, arg0, arg1)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address ) view returns(bool)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) StrategyIsWhitelistedForDeposit(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "strategyIsWhitelistedForDeposit", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address ) view returns(bool)
func (_StrategyManagerStorage *StrategyManagerStorageSession) StrategyIsWhitelistedForDeposit(arg0 common.Address) (bool, error) {
	return _StrategyManagerStorage.Contract.StrategyIsWhitelistedForDeposit(&_StrategyManagerStorage.CallOpts, arg0)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address ) view returns(bool)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) StrategyIsWhitelistedForDeposit(arg0 common.Address) (bool, error) {
	return _StrategyManagerStorage.Contract.StrategyIsWhitelistedForDeposit(&_StrategyManagerStorage.CallOpts, arg0)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) StrategyWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "strategyWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageSession) StrategyWhitelister() (common.Address, error) {
	return _StrategyManagerStorage.Contract.StrategyWhitelister(&_StrategyManagerStorage.CallOpts)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) StrategyWhitelister() (common.Address, error) {
	return _StrategyManagerStorage.Contract.StrategyWhitelister(&_StrategyManagerStorage.CallOpts)
}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address ) view returns(bool)
func (_StrategyManagerStorage *StrategyManagerStorageCaller) ThirdPartyTransfersForbidden(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StrategyManagerStorage.contract.Call(opts, &out, "thirdPartyTransfersForbidden", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address ) view returns(bool)
func (_StrategyManagerStorage *StrategyManagerStorageSession) ThirdPartyTransfersForbidden(arg0 common.Address) (bool, error) {
	return _StrategyManagerStorage.Contract.ThirdPartyTransfersForbidden(&_StrategyManagerStorage.CallOpts, arg0)
}

// ThirdPartyTransfersForbidden is a free data retrieval call binding the contract method 0x9b4da03d.
//
// Solidity: function thirdPartyTransfersForbidden(address ) view returns(bool)
func (_StrategyManagerStorage *StrategyManagerStorageCallerSession) ThirdPartyTransfersForbidden(arg0 common.Address) (bool, error) {
	return _StrategyManagerStorage.Contract.ThirdPartyTransfersForbidden(&_StrategyManagerStorage.CallOpts, arg0)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address token, address strategy, uint256 shares) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) AddShares(opts *bind.TransactOpts, staker common.Address, token common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "addShares", staker, token, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address token, address strategy, uint256 shares) returns()
func (_StrategyManagerStorage *StrategyManagerStorageSession) AddShares(staker common.Address, token common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.AddShares(&_StrategyManagerStorage.TransactOpts, staker, token, strategy, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address token, address strategy, uint256 shares) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) AddShares(staker common.Address, token common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.AddShares(&_StrategyManagerStorage.TransactOpts, staker, token, strategy, shares)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0xdf5b3547.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) AddStrategiesToDepositWhitelist(opts *bind.TransactOpts, strategiesToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "addStrategiesToDepositWhitelist", strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0xdf5b3547.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_StrategyManagerStorage *StrategyManagerStorageSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.AddStrategiesToDepositWhitelist(&_StrategyManagerStorage.TransactOpts, strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0xdf5b3547.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist, bool[] thirdPartyTransfersForbiddenValues) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address, thirdPartyTransfersForbiddenValues []bool) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.AddStrategiesToDepositWhitelist(&_StrategyManagerStorage.TransactOpts, strategiesToWhitelist, thirdPartyTransfersForbiddenValues)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) DepositIntoStrategy(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "depositIntoStrategy", strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.DepositIntoStrategy(&_StrategyManagerStorage.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.DepositIntoStrategy(&_StrategyManagerStorage.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) DepositIntoStrategyWithSignature(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "depositIntoStrategyWithSignature", strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.DepositIntoStrategyWithSignature(&_StrategyManagerStorage.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.DepositIntoStrategyWithSignature(&_StrategyManagerStorage.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) RemoveShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "removeShares", staker, strategy, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_StrategyManagerStorage *StrategyManagerStorageSession) RemoveShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.RemoveShares(&_StrategyManagerStorage.TransactOpts, staker, strategy, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) RemoveShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.RemoveShares(&_StrategyManagerStorage.TransactOpts, staker, strategy, shares)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) RemoveStrategiesFromDepositWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "removeStrategiesFromDepositWhitelist", strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyManagerStorage *StrategyManagerStorageSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.RemoveStrategiesFromDepositWhitelist(&_StrategyManagerStorage.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.RemoveStrategiesFromDepositWhitelist(&_StrategyManagerStorage.TransactOpts, strategiesToRemoveFromWhitelist)
}

// SetThirdPartyTransfersForbidden is a paid mutator transaction binding the contract method 0x4e5a4263.
//
// Solidity: function setThirdPartyTransfersForbidden(address strategy, bool value) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) SetThirdPartyTransfersForbidden(opts *bind.TransactOpts, strategy common.Address, value bool) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "setThirdPartyTransfersForbidden", strategy, value)
}

// SetThirdPartyTransfersForbidden is a paid mutator transaction binding the contract method 0x4e5a4263.
//
// Solidity: function setThirdPartyTransfersForbidden(address strategy, bool value) returns()
func (_StrategyManagerStorage *StrategyManagerStorageSession) SetThirdPartyTransfersForbidden(strategy common.Address, value bool) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.SetThirdPartyTransfersForbidden(&_StrategyManagerStorage.TransactOpts, strategy, value)
}

// SetThirdPartyTransfersForbidden is a paid mutator transaction binding the contract method 0x4e5a4263.
//
// Solidity: function setThirdPartyTransfersForbidden(address strategy, bool value) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) SetThirdPartyTransfersForbidden(strategy common.Address, value bool) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.SetThirdPartyTransfersForbidden(&_StrategyManagerStorage.TransactOpts, strategy, value)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address strategy, uint256 shares, address token) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, recipient common.Address, strategy common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.contract.Transact(opts, "withdrawSharesAsTokens", recipient, strategy, shares, token)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address strategy, uint256 shares, address token) returns()
func (_StrategyManagerStorage *StrategyManagerStorageSession) WithdrawSharesAsTokens(recipient common.Address, strategy common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.WithdrawSharesAsTokens(&_StrategyManagerStorage.TransactOpts, recipient, strategy, shares, token)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0xc608c7f3.
//
// Solidity: function withdrawSharesAsTokens(address recipient, address strategy, uint256 shares, address token) returns()
func (_StrategyManagerStorage *StrategyManagerStorageTransactorSession) WithdrawSharesAsTokens(recipient common.Address, strategy common.Address, shares *big.Int, token common.Address) (*types.Transaction, error) {
	return _StrategyManagerStorage.Contract.WithdrawSharesAsTokens(&_StrategyManagerStorage.TransactOpts, recipient, strategy, shares, token)
}

// StrategyManagerStorageDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the StrategyManagerStorage contract.
type StrategyManagerStorageDepositIterator struct {
	Event *StrategyManagerStorageDeposit // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStorageDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStorageDeposit)
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
		it.Event = new(StrategyManagerStorageDeposit)
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
func (it *StrategyManagerStorageDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStorageDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStorageDeposit represents a Deposit event raised by the StrategyManagerStorage contract.
type StrategyManagerStorageDeposit struct {
	Staker   common.Address
	Token    common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) FilterDeposit(opts *bind.FilterOpts) (*StrategyManagerStorageDepositIterator, error) {

	logs, sub, err := _StrategyManagerStorage.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageDepositIterator{contract: _StrategyManagerStorage.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *StrategyManagerStorageDeposit) (event.Subscription, error) {

	logs, sub, err := _StrategyManagerStorage.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStorageDeposit)
				if err := _StrategyManagerStorage.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) ParseDeposit(log types.Log) (*StrategyManagerStorageDeposit, error) {
	event := new(StrategyManagerStorageDeposit)
	if err := _StrategyManagerStorage.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerStorageStrategyAddedToDepositWhitelistIterator is returned from FilterStrategyAddedToDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyAddedToDepositWhitelist events raised by the StrategyManagerStorage contract.
type StrategyManagerStorageStrategyAddedToDepositWhitelistIterator struct {
	Event *StrategyManagerStorageStrategyAddedToDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStorageStrategyAddedToDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStorageStrategyAddedToDepositWhitelist)
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
		it.Event = new(StrategyManagerStorageStrategyAddedToDepositWhitelist)
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
func (it *StrategyManagerStorageStrategyAddedToDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStorageStrategyAddedToDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStorageStrategyAddedToDepositWhitelist represents a StrategyAddedToDepositWhitelist event raised by the StrategyManagerStorage contract.
type StrategyManagerStorageStrategyAddedToDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToDepositWhitelist is a free log retrieval operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) FilterStrategyAddedToDepositWhitelist(opts *bind.FilterOpts) (*StrategyManagerStorageStrategyAddedToDepositWhitelistIterator, error) {

	logs, sub, err := _StrategyManagerStorage.contract.FilterLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageStrategyAddedToDepositWhitelistIterator{contract: _StrategyManagerStorage.contract, event: "StrategyAddedToDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToDepositWhitelist is a free log subscription operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) WatchStrategyAddedToDepositWhitelist(opts *bind.WatchOpts, sink chan<- *StrategyManagerStorageStrategyAddedToDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _StrategyManagerStorage.contract.WatchLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStorageStrategyAddedToDepositWhitelist)
				if err := _StrategyManagerStorage.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
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

// ParseStrategyAddedToDepositWhitelist is a log parse operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) ParseStrategyAddedToDepositWhitelist(log types.Log) (*StrategyManagerStorageStrategyAddedToDepositWhitelist, error) {
	event := new(StrategyManagerStorageStrategyAddedToDepositWhitelist)
	if err := _StrategyManagerStorage.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerStorageStrategyRemovedFromDepositWhitelistIterator is returned from FilterStrategyRemovedFromDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromDepositWhitelist events raised by the StrategyManagerStorage contract.
type StrategyManagerStorageStrategyRemovedFromDepositWhitelistIterator struct {
	Event *StrategyManagerStorageStrategyRemovedFromDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStorageStrategyRemovedFromDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStorageStrategyRemovedFromDepositWhitelist)
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
		it.Event = new(StrategyManagerStorageStrategyRemovedFromDepositWhitelist)
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
func (it *StrategyManagerStorageStrategyRemovedFromDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStorageStrategyRemovedFromDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStorageStrategyRemovedFromDepositWhitelist represents a StrategyRemovedFromDepositWhitelist event raised by the StrategyManagerStorage contract.
type StrategyManagerStorageStrategyRemovedFromDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromDepositWhitelist is a free log retrieval operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) FilterStrategyRemovedFromDepositWhitelist(opts *bind.FilterOpts) (*StrategyManagerStorageStrategyRemovedFromDepositWhitelistIterator, error) {

	logs, sub, err := _StrategyManagerStorage.contract.FilterLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageStrategyRemovedFromDepositWhitelistIterator{contract: _StrategyManagerStorage.contract, event: "StrategyRemovedFromDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromDepositWhitelist is a free log subscription operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) WatchStrategyRemovedFromDepositWhitelist(opts *bind.WatchOpts, sink chan<- *StrategyManagerStorageStrategyRemovedFromDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _StrategyManagerStorage.contract.WatchLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStorageStrategyRemovedFromDepositWhitelist)
				if err := _StrategyManagerStorage.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
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

// ParseStrategyRemovedFromDepositWhitelist is a log parse operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) ParseStrategyRemovedFromDepositWhitelist(log types.Log) (*StrategyManagerStorageStrategyRemovedFromDepositWhitelist, error) {
	event := new(StrategyManagerStorageStrategyRemovedFromDepositWhitelist)
	if err := _StrategyManagerStorage.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerStorageStrategyWhitelisterChangedIterator is returned from FilterStrategyWhitelisterChanged and is used to iterate over the raw logs and unpacked data for StrategyWhitelisterChanged events raised by the StrategyManagerStorage contract.
type StrategyManagerStorageStrategyWhitelisterChangedIterator struct {
	Event *StrategyManagerStorageStrategyWhitelisterChanged // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStorageStrategyWhitelisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStorageStrategyWhitelisterChanged)
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
		it.Event = new(StrategyManagerStorageStrategyWhitelisterChanged)
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
func (it *StrategyManagerStorageStrategyWhitelisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStorageStrategyWhitelisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStorageStrategyWhitelisterChanged represents a StrategyWhitelisterChanged event raised by the StrategyManagerStorage contract.
type StrategyManagerStorageStrategyWhitelisterChanged struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStrategyWhitelisterChanged is a free log retrieval operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) FilterStrategyWhitelisterChanged(opts *bind.FilterOpts) (*StrategyManagerStorageStrategyWhitelisterChangedIterator, error) {

	logs, sub, err := _StrategyManagerStorage.contract.FilterLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageStrategyWhitelisterChangedIterator{contract: _StrategyManagerStorage.contract, event: "StrategyWhitelisterChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyWhitelisterChanged is a free log subscription operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) WatchStrategyWhitelisterChanged(opts *bind.WatchOpts, sink chan<- *StrategyManagerStorageStrategyWhitelisterChanged) (event.Subscription, error) {

	logs, sub, err := _StrategyManagerStorage.contract.WatchLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStorageStrategyWhitelisterChanged)
				if err := _StrategyManagerStorage.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
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

// ParseStrategyWhitelisterChanged is a log parse operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) ParseStrategyWhitelisterChanged(log types.Log) (*StrategyManagerStorageStrategyWhitelisterChanged, error) {
	event := new(StrategyManagerStorageStrategyWhitelisterChanged)
	if err := _StrategyManagerStorage.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyManagerStorageUpdatedThirdPartyTransfersForbiddenIterator is returned from FilterUpdatedThirdPartyTransfersForbidden and is used to iterate over the raw logs and unpacked data for UpdatedThirdPartyTransfersForbidden events raised by the StrategyManagerStorage contract.
type StrategyManagerStorageUpdatedThirdPartyTransfersForbiddenIterator struct {
	Event *StrategyManagerStorageUpdatedThirdPartyTransfersForbidden // Event containing the contract specifics and raw log

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
func (it *StrategyManagerStorageUpdatedThirdPartyTransfersForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyManagerStorageUpdatedThirdPartyTransfersForbidden)
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
		it.Event = new(StrategyManagerStorageUpdatedThirdPartyTransfersForbidden)
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
func (it *StrategyManagerStorageUpdatedThirdPartyTransfersForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyManagerStorageUpdatedThirdPartyTransfersForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyManagerStorageUpdatedThirdPartyTransfersForbidden represents a UpdatedThirdPartyTransfersForbidden event raised by the StrategyManagerStorage contract.
type StrategyManagerStorageUpdatedThirdPartyTransfersForbidden struct {
	Strategy common.Address
	Value    bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedThirdPartyTransfersForbidden is a free log retrieval operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address strategy, bool value)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) FilterUpdatedThirdPartyTransfersForbidden(opts *bind.FilterOpts) (*StrategyManagerStorageUpdatedThirdPartyTransfersForbiddenIterator, error) {

	logs, sub, err := _StrategyManagerStorage.contract.FilterLogs(opts, "UpdatedThirdPartyTransfersForbidden")
	if err != nil {
		return nil, err
	}
	return &StrategyManagerStorageUpdatedThirdPartyTransfersForbiddenIterator{contract: _StrategyManagerStorage.contract, event: "UpdatedThirdPartyTransfersForbidden", logs: logs, sub: sub}, nil
}

// WatchUpdatedThirdPartyTransfersForbidden is a free log subscription operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address strategy, bool value)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) WatchUpdatedThirdPartyTransfersForbidden(opts *bind.WatchOpts, sink chan<- *StrategyManagerStorageUpdatedThirdPartyTransfersForbidden) (event.Subscription, error) {

	logs, sub, err := _StrategyManagerStorage.contract.WatchLogs(opts, "UpdatedThirdPartyTransfersForbidden")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyManagerStorageUpdatedThirdPartyTransfersForbidden)
				if err := _StrategyManagerStorage.contract.UnpackLog(event, "UpdatedThirdPartyTransfersForbidden", log); err != nil {
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

// ParseUpdatedThirdPartyTransfersForbidden is a log parse operation binding the contract event 0x77d930df4937793473a95024d87a98fd2ccb9e92d3c2463b3dacd65d3e6a5786.
//
// Solidity: event UpdatedThirdPartyTransfersForbidden(address strategy, bool value)
func (_StrategyManagerStorage *StrategyManagerStorageFilterer) ParseUpdatedThirdPartyTransfersForbidden(log types.Log) (*StrategyManagerStorageUpdatedThirdPartyTransfersForbidden, error) {
	event := new(StrategyManagerStorageUpdatedThirdPartyTransfersForbidden)
	if err := _StrategyManagerStorage.contract.UnpackLog(event, "UpdatedThirdPartyTransfersForbidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
