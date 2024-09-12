// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SlashingConstants

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

// SlashingConstantsMetaData contains all meta data concerning the SlashingConstants contract.
var SlashingConstantsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEALLOCATION_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INITIAL_TOTAL_MAGNITUDE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRECISION_FACTOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRECISION_FACTOR_SQUARED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x61010561003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060515760003560e01c806321afdf8e1460565780632981eb7714607e5780639a543ca414609b578063ccd34cd51460c1575b600080fd5b606b6ec097ce7bc90715b34b9f100000000081565b6040519081526020015b60405180910390f35b60876217124081565b60405163ffffffff90911681526020016075565b60a9670de0b6b3a764000081565b60405167ffffffffffffffff90911681526020016075565b606b670de0b6b3a76400008156fea264697066735822122051573b877db53ce7fd9202510ed740ac5b29774096c961210de6c69c249f537264736f6c634300080c0033",
}

// SlashingConstantsABI is the input ABI used to generate the binding from.
// Deprecated: Use SlashingConstantsMetaData.ABI instead.
var SlashingConstantsABI = SlashingConstantsMetaData.ABI

// SlashingConstantsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlashingConstantsMetaData.Bin instead.
var SlashingConstantsBin = SlashingConstantsMetaData.Bin

// DeploySlashingConstants deploys a new Ethereum contract, binding an instance of SlashingConstants to it.
func DeploySlashingConstants(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SlashingConstants, error) {
	parsed, err := SlashingConstantsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlashingConstantsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SlashingConstants{SlashingConstantsCaller: SlashingConstantsCaller{contract: contract}, SlashingConstantsTransactor: SlashingConstantsTransactor{contract: contract}, SlashingConstantsFilterer: SlashingConstantsFilterer{contract: contract}}, nil
}

// SlashingConstants is an auto generated Go binding around an Ethereum contract.
type SlashingConstants struct {
	SlashingConstantsCaller     // Read-only binding to the contract
	SlashingConstantsTransactor // Write-only binding to the contract
	SlashingConstantsFilterer   // Log filterer for contract events
}

// SlashingConstantsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlashingConstantsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingConstantsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlashingConstantsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingConstantsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlashingConstantsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingConstantsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlashingConstantsSession struct {
	Contract     *SlashingConstants // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SlashingConstantsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlashingConstantsCallerSession struct {
	Contract *SlashingConstantsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SlashingConstantsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlashingConstantsTransactorSession struct {
	Contract     *SlashingConstantsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SlashingConstantsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlashingConstantsRaw struct {
	Contract *SlashingConstants // Generic contract binding to access the raw methods on
}

// SlashingConstantsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlashingConstantsCallerRaw struct {
	Contract *SlashingConstantsCaller // Generic read-only contract binding to access the raw methods on
}

// SlashingConstantsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlashingConstantsTransactorRaw struct {
	Contract *SlashingConstantsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlashingConstants creates a new instance of SlashingConstants, bound to a specific deployed contract.
func NewSlashingConstants(address common.Address, backend bind.ContractBackend) (*SlashingConstants, error) {
	contract, err := bindSlashingConstants(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlashingConstants{SlashingConstantsCaller: SlashingConstantsCaller{contract: contract}, SlashingConstantsTransactor: SlashingConstantsTransactor{contract: contract}, SlashingConstantsFilterer: SlashingConstantsFilterer{contract: contract}}, nil
}

// NewSlashingConstantsCaller creates a new read-only instance of SlashingConstants, bound to a specific deployed contract.
func NewSlashingConstantsCaller(address common.Address, caller bind.ContractCaller) (*SlashingConstantsCaller, error) {
	contract, err := bindSlashingConstants(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlashingConstantsCaller{contract: contract}, nil
}

// NewSlashingConstantsTransactor creates a new write-only instance of SlashingConstants, bound to a specific deployed contract.
func NewSlashingConstantsTransactor(address common.Address, transactor bind.ContractTransactor) (*SlashingConstantsTransactor, error) {
	contract, err := bindSlashingConstants(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlashingConstantsTransactor{contract: contract}, nil
}

// NewSlashingConstantsFilterer creates a new log filterer instance of SlashingConstants, bound to a specific deployed contract.
func NewSlashingConstantsFilterer(address common.Address, filterer bind.ContractFilterer) (*SlashingConstantsFilterer, error) {
	contract, err := bindSlashingConstants(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlashingConstantsFilterer{contract: contract}, nil
}

// bindSlashingConstants binds a generic wrapper to an already deployed contract.
func bindSlashingConstants(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlashingConstantsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashingConstants *SlashingConstantsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashingConstants.Contract.SlashingConstantsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashingConstants *SlashingConstantsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashingConstants.Contract.SlashingConstantsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashingConstants *SlashingConstantsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashingConstants.Contract.SlashingConstantsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashingConstants *SlashingConstantsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashingConstants.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashingConstants *SlashingConstantsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashingConstants.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashingConstants *SlashingConstantsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashingConstants.Contract.contract.Transact(opts, method, params...)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_SlashingConstants *SlashingConstantsCaller) DEALLOCATIONDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SlashingConstants.contract.Call(opts, &out, "DEALLOCATION_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_SlashingConstants *SlashingConstantsSession) DEALLOCATIONDELAY() (uint32, error) {
	return _SlashingConstants.Contract.DEALLOCATIONDELAY(&_SlashingConstants.CallOpts)
}

// DEALLOCATIONDELAY is a free data retrieval call binding the contract method 0x2981eb77.
//
// Solidity: function DEALLOCATION_DELAY() view returns(uint32)
func (_SlashingConstants *SlashingConstantsCallerSession) DEALLOCATIONDELAY() (uint32, error) {
	return _SlashingConstants.Contract.DEALLOCATIONDELAY(&_SlashingConstants.CallOpts)
}

// INITIALTOTALMAGNITUDE is a free data retrieval call binding the contract method 0x9a543ca4.
//
// Solidity: function INITIAL_TOTAL_MAGNITUDE() view returns(uint64)
func (_SlashingConstants *SlashingConstantsCaller) INITIALTOTALMAGNITUDE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SlashingConstants.contract.Call(opts, &out, "INITIAL_TOTAL_MAGNITUDE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// INITIALTOTALMAGNITUDE is a free data retrieval call binding the contract method 0x9a543ca4.
//
// Solidity: function INITIAL_TOTAL_MAGNITUDE() view returns(uint64)
func (_SlashingConstants *SlashingConstantsSession) INITIALTOTALMAGNITUDE() (uint64, error) {
	return _SlashingConstants.Contract.INITIALTOTALMAGNITUDE(&_SlashingConstants.CallOpts)
}

// INITIALTOTALMAGNITUDE is a free data retrieval call binding the contract method 0x9a543ca4.
//
// Solidity: function INITIAL_TOTAL_MAGNITUDE() view returns(uint64)
func (_SlashingConstants *SlashingConstantsCallerSession) INITIALTOTALMAGNITUDE() (uint64, error) {
	return _SlashingConstants.Contract.INITIALTOTALMAGNITUDE(&_SlashingConstants.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_SlashingConstants *SlashingConstantsCaller) PRECISIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashingConstants.contract.Call(opts, &out, "PRECISION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_SlashingConstants *SlashingConstantsSession) PRECISIONFACTOR() (*big.Int, error) {
	return _SlashingConstants.Contract.PRECISIONFACTOR(&_SlashingConstants.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_SlashingConstants *SlashingConstantsCallerSession) PRECISIONFACTOR() (*big.Int, error) {
	return _SlashingConstants.Contract.PRECISIONFACTOR(&_SlashingConstants.CallOpts)
}

// PRECISIONFACTORSQUARED is a free data retrieval call binding the contract method 0x21afdf8e.
//
// Solidity: function PRECISION_FACTOR_SQUARED() view returns(uint256)
func (_SlashingConstants *SlashingConstantsCaller) PRECISIONFACTORSQUARED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashingConstants.contract.Call(opts, &out, "PRECISION_FACTOR_SQUARED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTORSQUARED is a free data retrieval call binding the contract method 0x21afdf8e.
//
// Solidity: function PRECISION_FACTOR_SQUARED() view returns(uint256)
func (_SlashingConstants *SlashingConstantsSession) PRECISIONFACTORSQUARED() (*big.Int, error) {
	return _SlashingConstants.Contract.PRECISIONFACTORSQUARED(&_SlashingConstants.CallOpts)
}

// PRECISIONFACTORSQUARED is a free data retrieval call binding the contract method 0x21afdf8e.
//
// Solidity: function PRECISION_FACTOR_SQUARED() view returns(uint256)
func (_SlashingConstants *SlashingConstantsCallerSession) PRECISIONFACTORSQUARED() (*big.Int, error) {
	return _SlashingConstants.Contract.PRECISIONFACTORSQUARED(&_SlashingConstants.CallOpts)
}
