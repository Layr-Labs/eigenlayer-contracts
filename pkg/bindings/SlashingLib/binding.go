// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SlashingLib

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

// SlashingLibMetaData contains all meta data concerning the SlashingLib contract.
var SlashingLibMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INITIAL_TOTAL_MAGNITUDE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRECISION_FACTOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x60e6610038600b82828239805160001a607314602b57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060475760003560e01c80632981eb7714604c5780639a543ca414606f578063ccd34cd5146095575b600080fd5b60556217124081565b60405163ffffffff90911681526020015b60405180910390f35b607d670de0b6b3a764000081565b60405167ffffffffffffffff90911681526020016066565b60a3670de0b6b3a764000081565b604051908152602001606656fea2646970667358221220ef31d2c613ebe47dda6f05a88229016986f1f1262195fa7cfdb883a9aaf71f5264736f6c634300081b0033",
}

// SlashingLibABI is the input ABI used to generate the binding from.
// Deprecated: Use SlashingLibMetaData.ABI instead.
var SlashingLibABI = SlashingLibMetaData.ABI

// SlashingLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlashingLibMetaData.Bin instead.
var SlashingLibBin = SlashingLibMetaData.Bin

// DeploySlashingLib deploys a new Ethereum contract, binding an instance of SlashingLib to it.
func DeploySlashingLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SlashingLib, error) {
	parsed, err := SlashingLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlashingLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SlashingLib{SlashingLibCaller: SlashingLibCaller{contract: contract}, SlashingLibTransactor: SlashingLibTransactor{contract: contract}, SlashingLibFilterer: SlashingLibFilterer{contract: contract}}, nil
}

// SlashingLib is an auto generated Go binding around an Ethereum contract.
type SlashingLib struct {
	SlashingLibCaller     // Read-only binding to the contract
	SlashingLibTransactor // Write-only binding to the contract
	SlashingLibFilterer   // Log filterer for contract events
}

// SlashingLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlashingLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlashingLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlashingLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlashingLibSession struct {
	Contract     *SlashingLib      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlashingLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlashingLibCallerSession struct {
	Contract *SlashingLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SlashingLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlashingLibTransactorSession struct {
	Contract     *SlashingLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SlashingLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlashingLibRaw struct {
	Contract *SlashingLib // Generic contract binding to access the raw methods on
}

// SlashingLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlashingLibCallerRaw struct {
	Contract *SlashingLibCaller // Generic read-only contract binding to access the raw methods on
}

// SlashingLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlashingLibTransactorRaw struct {
	Contract *SlashingLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlashingLib creates a new instance of SlashingLib, bound to a specific deployed contract.
func NewSlashingLib(address common.Address, backend bind.ContractBackend) (*SlashingLib, error) {
	contract, err := bindSlashingLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlashingLib{SlashingLibCaller: SlashingLibCaller{contract: contract}, SlashingLibTransactor: SlashingLibTransactor{contract: contract}, SlashingLibFilterer: SlashingLibFilterer{contract: contract}}, nil
}

// NewSlashingLibCaller creates a new read-only instance of SlashingLib, bound to a specific deployed contract.
func NewSlashingLibCaller(address common.Address, caller bind.ContractCaller) (*SlashingLibCaller, error) {
	contract, err := bindSlashingLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlashingLibCaller{contract: contract}, nil
}

// NewSlashingLibTransactor creates a new write-only instance of SlashingLib, bound to a specific deployed contract.
func NewSlashingLibTransactor(address common.Address, transactor bind.ContractTransactor) (*SlashingLibTransactor, error) {
	contract, err := bindSlashingLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlashingLibTransactor{contract: contract}, nil
}

// NewSlashingLibFilterer creates a new log filterer instance of SlashingLib, bound to a specific deployed contract.
func NewSlashingLibFilterer(address common.Address, filterer bind.ContractFilterer) (*SlashingLibFilterer, error) {
	contract, err := bindSlashingLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlashingLibFilterer{contract: contract}, nil
}

// bindSlashingLib binds a generic wrapper to an already deployed contract.
func bindSlashingLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlashingLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashingLib *SlashingLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashingLib.Contract.SlashingLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashingLib *SlashingLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashingLib.Contract.SlashingLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashingLib *SlashingLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashingLib.Contract.SlashingLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashingLib *SlashingLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashingLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashingLib *SlashingLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashingLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashingLib *SlashingLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashingLib.Contract.contract.Transact(opts, method, params...)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_SlashingLib *SlashingLibCaller) DEALLOCATIONDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SlashingLib.contract.Call(opts, &out, "DEALLOCATION_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_SlashingLib *SlashingLibSession) DEALLOCATIONDELAY() (uint32, error) {
	return _SlashingLib.Contract.DEALLOCATIONDELAY(&_SlashingLib.CallOpts)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_SlashingLib *SlashingLibCallerSession) DEALLOCATIONDELAY() (uint32, error) {
	return _SlashingLib.Contract.DEALLOCATIONDELAY(&_SlashingLib.CallOpts)
}

// INITIALTOTALMAGNITUDE is a free data retrieval call binding the contract method 0x9a543ca4.
//
// Solidity: function INITIAL_TOTAL_MAGNITUDE() view returns(uint64)
func (_SlashingLib *SlashingLibCaller) INITIALTOTALMAGNITUDE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SlashingLib.contract.Call(opts, &out, "INITIAL_TOTAL_MAGNITUDE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// INITIALTOTALMAGNITUDE is a free data retrieval call binding the contract method 0x9a543ca4.
//
// Solidity: function INITIAL_TOTAL_MAGNITUDE() view returns(uint64)
func (_SlashingLib *SlashingLibSession) INITIALTOTALMAGNITUDE() (uint64, error) {
	return _SlashingLib.Contract.INITIALTOTALMAGNITUDE(&_SlashingLib.CallOpts)
}

// INITIALTOTALMAGNITUDE is a free data retrieval call binding the contract method 0x9a543ca4.
//
// Solidity: function INITIAL_TOTAL_MAGNITUDE() view returns(uint64)
func (_SlashingLib *SlashingLibCallerSession) INITIALTOTALMAGNITUDE() (uint64, error) {
	return _SlashingLib.Contract.INITIALTOTALMAGNITUDE(&_SlashingLib.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_SlashingLib *SlashingLibCaller) PRECISIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashingLib.contract.Call(opts, &out, "PRECISION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_SlashingLib *SlashingLibSession) PRECISIONFACTOR() (*big.Int, error) {
	return _SlashingLib.Contract.PRECISIONFACTOR(&_SlashingLib.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_SlashingLib *SlashingLibCallerSession) PRECISIONFACTOR() (*big.Int, error) {
	return _SlashingLib.Contract.PRECISIONFACTOR(&_SlashingLib.CallOpts)
}
