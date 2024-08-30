// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ShareScalingLib

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

// ShareScalingLibMetaData contains all meta data concerning the ShareScalingLib contract.
var ShareScalingLibMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"INITIAL_TOTAL_MAGNITUDE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x6099610038600b82828239805160001a607314602b57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c80639a543ca4146038575b600080fd5b6046670de0b6b3a764000081565b60405167ffffffffffffffff909116815260200160405180910390f3fea26469706673582212202a7c161dfde6ddff6dab226692633b8d1a15545a6aea30d4a2cc2caf7061b6f064736f6c634300080c0033",
}

// ShareScalingLibABI is the input ABI used to generate the binding from.
// Deprecated: Use ShareScalingLibMetaData.ABI instead.
var ShareScalingLibABI = ShareScalingLibMetaData.ABI

// ShareScalingLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ShareScalingLibMetaData.Bin instead.
var ShareScalingLibBin = ShareScalingLibMetaData.Bin

// DeployShareScalingLib deploys a new Ethereum contract, binding an instance of ShareScalingLib to it.
func DeployShareScalingLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ShareScalingLib, error) {
	parsed, err := ShareScalingLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ShareScalingLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ShareScalingLib{ShareScalingLibCaller: ShareScalingLibCaller{contract: contract}, ShareScalingLibTransactor: ShareScalingLibTransactor{contract: contract}, ShareScalingLibFilterer: ShareScalingLibFilterer{contract: contract}}, nil
}

// ShareScalingLib is an auto generated Go binding around an Ethereum contract.
type ShareScalingLib struct {
	ShareScalingLibCaller     // Read-only binding to the contract
	ShareScalingLibTransactor // Write-only binding to the contract
	ShareScalingLibFilterer   // Log filterer for contract events
}

// ShareScalingLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShareScalingLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareScalingLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShareScalingLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareScalingLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShareScalingLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareScalingLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShareScalingLibSession struct {
	Contract     *ShareScalingLib  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShareScalingLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShareScalingLibCallerSession struct {
	Contract *ShareScalingLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ShareScalingLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShareScalingLibTransactorSession struct {
	Contract     *ShareScalingLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ShareScalingLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShareScalingLibRaw struct {
	Contract *ShareScalingLib // Generic contract binding to access the raw methods on
}

// ShareScalingLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShareScalingLibCallerRaw struct {
	Contract *ShareScalingLibCaller // Generic read-only contract binding to access the raw methods on
}

// ShareScalingLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShareScalingLibTransactorRaw struct {
	Contract *ShareScalingLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShareScalingLib creates a new instance of ShareScalingLib, bound to a specific deployed contract.
func NewShareScalingLib(address common.Address, backend bind.ContractBackend) (*ShareScalingLib, error) {
	contract, err := bindShareScalingLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShareScalingLib{ShareScalingLibCaller: ShareScalingLibCaller{contract: contract}, ShareScalingLibTransactor: ShareScalingLibTransactor{contract: contract}, ShareScalingLibFilterer: ShareScalingLibFilterer{contract: contract}}, nil
}

// NewShareScalingLibCaller creates a new read-only instance of ShareScalingLib, bound to a specific deployed contract.
func NewShareScalingLibCaller(address common.Address, caller bind.ContractCaller) (*ShareScalingLibCaller, error) {
	contract, err := bindShareScalingLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShareScalingLibCaller{contract: contract}, nil
}

// NewShareScalingLibTransactor creates a new write-only instance of ShareScalingLib, bound to a specific deployed contract.
func NewShareScalingLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ShareScalingLibTransactor, error) {
	contract, err := bindShareScalingLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShareScalingLibTransactor{contract: contract}, nil
}

// NewShareScalingLibFilterer creates a new log filterer instance of ShareScalingLib, bound to a specific deployed contract.
func NewShareScalingLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ShareScalingLibFilterer, error) {
	contract, err := bindShareScalingLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShareScalingLibFilterer{contract: contract}, nil
}

// bindShareScalingLib binds a generic wrapper to an already deployed contract.
func bindShareScalingLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShareScalingLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShareScalingLib *ShareScalingLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShareScalingLib.Contract.ShareScalingLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShareScalingLib *ShareScalingLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShareScalingLib.Contract.ShareScalingLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShareScalingLib *ShareScalingLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShareScalingLib.Contract.ShareScalingLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShareScalingLib *ShareScalingLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShareScalingLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShareScalingLib *ShareScalingLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShareScalingLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShareScalingLib *ShareScalingLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShareScalingLib.Contract.contract.Transact(opts, method, params...)
}

// INITIALTOTALMAGNITUDE is a free data retrieval call binding the contract method 0x9a543ca4.
//
// Solidity: function INITIAL_TOTAL_MAGNITUDE() view returns(uint64)
func (_ShareScalingLib *ShareScalingLibCaller) INITIALTOTALMAGNITUDE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ShareScalingLib.contract.Call(opts, &out, "INITIAL_TOTAL_MAGNITUDE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// INITIALTOTALMAGNITUDE is a free data retrieval call binding the contract method 0x9a543ca4.
//
// Solidity: function INITIAL_TOTAL_MAGNITUDE() view returns(uint64)
func (_ShareScalingLib *ShareScalingLibSession) INITIALTOTALMAGNITUDE() (uint64, error) {
	return _ShareScalingLib.Contract.INITIALTOTALMAGNITUDE(&_ShareScalingLib.CallOpts)
}

// INITIALTOTALMAGNITUDE is a free data retrieval call binding the contract method 0x9a543ca4.
//
// Solidity: function INITIAL_TOTAL_MAGNITUDE() view returns(uint64)
func (_ShareScalingLib *ShareScalingLibCallerSession) INITIALTOTALMAGNITUDE() (uint64, error) {
	return _ShareScalingLib.Contract.INITIALTOTALMAGNITUDE(&_ShareScalingLib.CallOpts)
}
