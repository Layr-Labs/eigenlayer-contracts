// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BeaconChainProofs

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

// BeaconChainProofsMetaData contains all meta data concerning the BeaconChainProofs contract.
var BeaconChainProofsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValidatorFieldsLength\",\"inputs\":[]}]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea2646970667358221220bc00bd74e89effe2aa4cd060bd92c2a3553695b3a5d129d2f5a88577444944a964736f6c634300081b0033",
}

// BeaconChainProofsABI is the input ABI used to generate the binding from.
// Deprecated: Use BeaconChainProofsMetaData.ABI instead.
var BeaconChainProofsABI = BeaconChainProofsMetaData.ABI

// BeaconChainProofsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BeaconChainProofsMetaData.Bin instead.
var BeaconChainProofsBin = BeaconChainProofsMetaData.Bin

// DeployBeaconChainProofs deploys a new Ethereum contract, binding an instance of BeaconChainProofs to it.
func DeployBeaconChainProofs(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BeaconChainProofs, error) {
	parsed, err := BeaconChainProofsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BeaconChainProofsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BeaconChainProofs{BeaconChainProofsCaller: BeaconChainProofsCaller{contract: contract}, BeaconChainProofsTransactor: BeaconChainProofsTransactor{contract: contract}, BeaconChainProofsFilterer: BeaconChainProofsFilterer{contract: contract}}, nil
}

// BeaconChainProofs is an auto generated Go binding around an Ethereum contract.
type BeaconChainProofs struct {
	BeaconChainProofsCaller     // Read-only binding to the contract
	BeaconChainProofsTransactor // Write-only binding to the contract
	BeaconChainProofsFilterer   // Log filterer for contract events
}

// BeaconChainProofsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeaconChainProofsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconChainProofsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeaconChainProofsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconChainProofsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeaconChainProofsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconChainProofsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeaconChainProofsSession struct {
	Contract     *BeaconChainProofs // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BeaconChainProofsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeaconChainProofsCallerSession struct {
	Contract *BeaconChainProofsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BeaconChainProofsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeaconChainProofsTransactorSession struct {
	Contract     *BeaconChainProofsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BeaconChainProofsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeaconChainProofsRaw struct {
	Contract *BeaconChainProofs // Generic contract binding to access the raw methods on
}

// BeaconChainProofsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeaconChainProofsCallerRaw struct {
	Contract *BeaconChainProofsCaller // Generic read-only contract binding to access the raw methods on
}

// BeaconChainProofsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeaconChainProofsTransactorRaw struct {
	Contract *BeaconChainProofsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeaconChainProofs creates a new instance of BeaconChainProofs, bound to a specific deployed contract.
func NewBeaconChainProofs(address common.Address, backend bind.ContractBackend) (*BeaconChainProofs, error) {
	contract, err := bindBeaconChainProofs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeaconChainProofs{BeaconChainProofsCaller: BeaconChainProofsCaller{contract: contract}, BeaconChainProofsTransactor: BeaconChainProofsTransactor{contract: contract}, BeaconChainProofsFilterer: BeaconChainProofsFilterer{contract: contract}}, nil
}

// NewBeaconChainProofsCaller creates a new read-only instance of BeaconChainProofs, bound to a specific deployed contract.
func NewBeaconChainProofsCaller(address common.Address, caller bind.ContractCaller) (*BeaconChainProofsCaller, error) {
	contract, err := bindBeaconChainProofs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconChainProofsCaller{contract: contract}, nil
}

// NewBeaconChainProofsTransactor creates a new write-only instance of BeaconChainProofs, bound to a specific deployed contract.
func NewBeaconChainProofsTransactor(address common.Address, transactor bind.ContractTransactor) (*BeaconChainProofsTransactor, error) {
	contract, err := bindBeaconChainProofs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconChainProofsTransactor{contract: contract}, nil
}

// NewBeaconChainProofsFilterer creates a new log filterer instance of BeaconChainProofs, bound to a specific deployed contract.
func NewBeaconChainProofsFilterer(address common.Address, filterer bind.ContractFilterer) (*BeaconChainProofsFilterer, error) {
	contract, err := bindBeaconChainProofs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeaconChainProofsFilterer{contract: contract}, nil
}

// bindBeaconChainProofs binds a generic wrapper to an already deployed contract.
func bindBeaconChainProofs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeaconChainProofsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeaconChainProofs *BeaconChainProofsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconChainProofs.Contract.BeaconChainProofsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeaconChainProofs *BeaconChainProofsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeaconChainProofs.Contract.BeaconChainProofsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeaconChainProofs *BeaconChainProofsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeaconChainProofs.Contract.BeaconChainProofsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeaconChainProofs *BeaconChainProofsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconChainProofs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeaconChainProofs *BeaconChainProofsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeaconChainProofs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeaconChainProofs *BeaconChainProofsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeaconChainProofs.Contract.contract.Transact(opts, method, params...)
}
