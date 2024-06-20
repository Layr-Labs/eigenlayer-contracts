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
	ABI: "[{\"type\":\"function\",\"name\":\"getEffectiveBalanceGwei\",\"inputs\":[{\"name\":\"validatorFields\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"}]",
	Bin: "0x6103e2610053600b82828239805160001a607314610046577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c80634534711b1461003a575b600080fd5b610054600480360381019061004f91906102f6565b61006a565b6040516100619190610362565b60405180910390f35b6000610090826002815181106100835761008261037d565b5b6020026020010151610097565b9050919050565b600060c082901c60001c905060388160ff1667ffffffffffffffff16901b60288261ff001667ffffffffffffffff16901b60188362ff00001667ffffffffffffffff16901b60088463ff0000001667ffffffffffffffff16901b60088564ff000000001667ffffffffffffffff16901c60188665ff00000000001667ffffffffffffffff16901c60288766ff0000000000001667ffffffffffffffff16901c60388867ffffffffffffffff16901c171717171717179050919050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6101b58261016c565b810181811067ffffffffffffffff821117156101d4576101d361017d565b5b80604052505050565b60006101e7610153565b90506101f382826101ac565b919050565b600067ffffffffffffffff8211156102135761021261017d565b5b602082029050602081019050919050565b600080fd5b6000819050919050565b61023c81610229565b811461024757600080fd5b50565b60008135905061025981610233565b92915050565b600061027261026d846101f8565b6101dd565b9050808382526020820190506020840283018581111561029557610294610224565b5b835b818110156102be57806102aa888261024a565b845260208401935050602081019050610297565b5050509392505050565b600082601f8301126102dd576102dc610167565b5b81356102ed84826020860161025f565b91505092915050565b60006020828403121561030c5761030b61015d565b5b600082013567ffffffffffffffff81111561032a57610329610162565b5b610336848285016102c8565b91505092915050565b600067ffffffffffffffff82169050919050565b61035c8161033f565b82525050565b60006020820190506103776000830184610353565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220151971e99cfc44a4a3b70ef3ecd7413be4243e9929125af9ffb6c94137108b4164736f6c634300080c0033",
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

// GetEffectiveBalanceGwei is a free data retrieval call binding the contract method 0x4534711b.
//
// Solidity: function getEffectiveBalanceGwei(bytes32[] validatorFields) pure returns(uint64)
func (_BeaconChainProofs *BeaconChainProofsCaller) GetEffectiveBalanceGwei(opts *bind.CallOpts, validatorFields [][32]byte) (uint64, error) {
	var out []interface{}
	err := _BeaconChainProofs.contract.Call(opts, &out, "getEffectiveBalanceGwei", validatorFields)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetEffectiveBalanceGwei is a free data retrieval call binding the contract method 0x4534711b.
//
// Solidity: function getEffectiveBalanceGwei(bytes32[] validatorFields) pure returns(uint64)
func (_BeaconChainProofs *BeaconChainProofsSession) GetEffectiveBalanceGwei(validatorFields [][32]byte) (uint64, error) {
	return _BeaconChainProofs.Contract.GetEffectiveBalanceGwei(&_BeaconChainProofs.CallOpts, validatorFields)
}

// GetEffectiveBalanceGwei is a free data retrieval call binding the contract method 0x4534711b.
//
// Solidity: function getEffectiveBalanceGwei(bytes32[] validatorFields) pure returns(uint64)
func (_BeaconChainProofs *BeaconChainProofsCallerSession) GetEffectiveBalanceGwei(validatorFields [][32]byte) (uint64, error) {
	return _BeaconChainProofs.Contract.GetEffectiveBalanceGwei(&_BeaconChainProofs.CallOpts, validatorFields)
}
