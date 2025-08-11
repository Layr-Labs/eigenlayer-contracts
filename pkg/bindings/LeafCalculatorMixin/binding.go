// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LeafCalculatorMixin

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// IOperatorTableCalculatorTypesBN254OperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IOperatorTableCalculatorTypesBN254OperatorInfo struct {
	Pubkey  BN254G1Point
	Weights []*big.Int
}

// LeafCalculatorMixinMetaData contains all meta data concerning the LeafCalculatorMixin contract.
var LeafCalculatorMixinMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"OPERATOR_INFO_LEAF_SALT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_TABLE_LEAF_SALT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorInfoLeaf\",\"inputs\":[{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIOperatorTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"calculateOperatorTableLeaf\",\"inputs\":[{\"name\":\"operatorTableBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"}]",
}

// LeafCalculatorMixinABI is the input ABI used to generate the binding from.
// Deprecated: Use LeafCalculatorMixinMetaData.ABI instead.
var LeafCalculatorMixinABI = LeafCalculatorMixinMetaData.ABI

// LeafCalculatorMixin is an auto generated Go binding around an Ethereum contract.
type LeafCalculatorMixin struct {
	LeafCalculatorMixinCaller     // Read-only binding to the contract
	LeafCalculatorMixinTransactor // Write-only binding to the contract
	LeafCalculatorMixinFilterer   // Log filterer for contract events
}

// LeafCalculatorMixinCaller is an auto generated read-only Go binding around an Ethereum contract.
type LeafCalculatorMixinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeafCalculatorMixinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LeafCalculatorMixinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeafCalculatorMixinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LeafCalculatorMixinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeafCalculatorMixinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LeafCalculatorMixinSession struct {
	Contract     *LeafCalculatorMixin // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LeafCalculatorMixinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LeafCalculatorMixinCallerSession struct {
	Contract *LeafCalculatorMixinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// LeafCalculatorMixinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LeafCalculatorMixinTransactorSession struct {
	Contract     *LeafCalculatorMixinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// LeafCalculatorMixinRaw is an auto generated low-level Go binding around an Ethereum contract.
type LeafCalculatorMixinRaw struct {
	Contract *LeafCalculatorMixin // Generic contract binding to access the raw methods on
}

// LeafCalculatorMixinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LeafCalculatorMixinCallerRaw struct {
	Contract *LeafCalculatorMixinCaller // Generic read-only contract binding to access the raw methods on
}

// LeafCalculatorMixinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LeafCalculatorMixinTransactorRaw struct {
	Contract *LeafCalculatorMixinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLeafCalculatorMixin creates a new instance of LeafCalculatorMixin, bound to a specific deployed contract.
func NewLeafCalculatorMixin(address common.Address, backend bind.ContractBackend) (*LeafCalculatorMixin, error) {
	contract, err := bindLeafCalculatorMixin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LeafCalculatorMixin{LeafCalculatorMixinCaller: LeafCalculatorMixinCaller{contract: contract}, LeafCalculatorMixinTransactor: LeafCalculatorMixinTransactor{contract: contract}, LeafCalculatorMixinFilterer: LeafCalculatorMixinFilterer{contract: contract}}, nil
}

// NewLeafCalculatorMixinCaller creates a new read-only instance of LeafCalculatorMixin, bound to a specific deployed contract.
func NewLeafCalculatorMixinCaller(address common.Address, caller bind.ContractCaller) (*LeafCalculatorMixinCaller, error) {
	contract, err := bindLeafCalculatorMixin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LeafCalculatorMixinCaller{contract: contract}, nil
}

// NewLeafCalculatorMixinTransactor creates a new write-only instance of LeafCalculatorMixin, bound to a specific deployed contract.
func NewLeafCalculatorMixinTransactor(address common.Address, transactor bind.ContractTransactor) (*LeafCalculatorMixinTransactor, error) {
	contract, err := bindLeafCalculatorMixin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LeafCalculatorMixinTransactor{contract: contract}, nil
}

// NewLeafCalculatorMixinFilterer creates a new log filterer instance of LeafCalculatorMixin, bound to a specific deployed contract.
func NewLeafCalculatorMixinFilterer(address common.Address, filterer bind.ContractFilterer) (*LeafCalculatorMixinFilterer, error) {
	contract, err := bindLeafCalculatorMixin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LeafCalculatorMixinFilterer{contract: contract}, nil
}

// bindLeafCalculatorMixin binds a generic wrapper to an already deployed contract.
func bindLeafCalculatorMixin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LeafCalculatorMixinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeafCalculatorMixin *LeafCalculatorMixinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeafCalculatorMixin.Contract.LeafCalculatorMixinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeafCalculatorMixin *LeafCalculatorMixinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeafCalculatorMixin.Contract.LeafCalculatorMixinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeafCalculatorMixin *LeafCalculatorMixinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeafCalculatorMixin.Contract.LeafCalculatorMixinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeafCalculatorMixin *LeafCalculatorMixinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeafCalculatorMixin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeafCalculatorMixin *LeafCalculatorMixinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeafCalculatorMixin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeafCalculatorMixin *LeafCalculatorMixinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeafCalculatorMixin.Contract.contract.Transact(opts, method, params...)
}

// OPERATORINFOLEAFSALT is a free data retrieval call binding the contract method 0xa2c902f5.
//
// Solidity: function OPERATOR_INFO_LEAF_SALT() view returns(uint8)
func (_LeafCalculatorMixin *LeafCalculatorMixinCaller) OPERATORINFOLEAFSALT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LeafCalculatorMixin.contract.Call(opts, &out, "OPERATOR_INFO_LEAF_SALT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPERATORINFOLEAFSALT is a free data retrieval call binding the contract method 0xa2c902f5.
//
// Solidity: function OPERATOR_INFO_LEAF_SALT() view returns(uint8)
func (_LeafCalculatorMixin *LeafCalculatorMixinSession) OPERATORINFOLEAFSALT() (uint8, error) {
	return _LeafCalculatorMixin.Contract.OPERATORINFOLEAFSALT(&_LeafCalculatorMixin.CallOpts)
}

// OPERATORINFOLEAFSALT is a free data retrieval call binding the contract method 0xa2c902f5.
//
// Solidity: function OPERATOR_INFO_LEAF_SALT() view returns(uint8)
func (_LeafCalculatorMixin *LeafCalculatorMixinCallerSession) OPERATORINFOLEAFSALT() (uint8, error) {
	return _LeafCalculatorMixin.Contract.OPERATORINFOLEAFSALT(&_LeafCalculatorMixin.CallOpts)
}

// OPERATORTABLELEAFSALT is a free data retrieval call binding the contract method 0x121409ea.
//
// Solidity: function OPERATOR_TABLE_LEAF_SALT() view returns(uint8)
func (_LeafCalculatorMixin *LeafCalculatorMixinCaller) OPERATORTABLELEAFSALT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LeafCalculatorMixin.contract.Call(opts, &out, "OPERATOR_TABLE_LEAF_SALT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPERATORTABLELEAFSALT is a free data retrieval call binding the contract method 0x121409ea.
//
// Solidity: function OPERATOR_TABLE_LEAF_SALT() view returns(uint8)
func (_LeafCalculatorMixin *LeafCalculatorMixinSession) OPERATORTABLELEAFSALT() (uint8, error) {
	return _LeafCalculatorMixin.Contract.OPERATORTABLELEAFSALT(&_LeafCalculatorMixin.CallOpts)
}

// OPERATORTABLELEAFSALT is a free data retrieval call binding the contract method 0x121409ea.
//
// Solidity: function OPERATOR_TABLE_LEAF_SALT() view returns(uint8)
func (_LeafCalculatorMixin *LeafCalculatorMixinCallerSession) OPERATORTABLELEAFSALT() (uint8, error) {
	return _LeafCalculatorMixin.Contract.OPERATORTABLELEAFSALT(&_LeafCalculatorMixin.CallOpts)
}

// CalculateOperatorInfoLeaf is a free data retrieval call binding the contract method 0x538a3790.
//
// Solidity: function calculateOperatorInfoLeaf(((uint256,uint256),uint256[]) operatorInfo) pure returns(bytes32)
func (_LeafCalculatorMixin *LeafCalculatorMixinCaller) CalculateOperatorInfoLeaf(opts *bind.CallOpts, operatorInfo IOperatorTableCalculatorTypesBN254OperatorInfo) ([32]byte, error) {
	var out []interface{}
	err := _LeafCalculatorMixin.contract.Call(opts, &out, "calculateOperatorInfoLeaf", operatorInfo)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorInfoLeaf is a free data retrieval call binding the contract method 0x538a3790.
//
// Solidity: function calculateOperatorInfoLeaf(((uint256,uint256),uint256[]) operatorInfo) pure returns(bytes32)
func (_LeafCalculatorMixin *LeafCalculatorMixinSession) CalculateOperatorInfoLeaf(operatorInfo IOperatorTableCalculatorTypesBN254OperatorInfo) ([32]byte, error) {
	return _LeafCalculatorMixin.Contract.CalculateOperatorInfoLeaf(&_LeafCalculatorMixin.CallOpts, operatorInfo)
}

// CalculateOperatorInfoLeaf is a free data retrieval call binding the contract method 0x538a3790.
//
// Solidity: function calculateOperatorInfoLeaf(((uint256,uint256),uint256[]) operatorInfo) pure returns(bytes32)
func (_LeafCalculatorMixin *LeafCalculatorMixinCallerSession) CalculateOperatorInfoLeaf(operatorInfo IOperatorTableCalculatorTypesBN254OperatorInfo) ([32]byte, error) {
	return _LeafCalculatorMixin.Contract.CalculateOperatorInfoLeaf(&_LeafCalculatorMixin.CallOpts, operatorInfo)
}

// CalculateOperatorTableLeaf is a free data retrieval call binding the contract method 0xa2f2e24d.
//
// Solidity: function calculateOperatorTableLeaf(bytes operatorTableBytes) pure returns(bytes32)
func (_LeafCalculatorMixin *LeafCalculatorMixinCaller) CalculateOperatorTableLeaf(opts *bind.CallOpts, operatorTableBytes []byte) ([32]byte, error) {
	var out []interface{}
	err := _LeafCalculatorMixin.contract.Call(opts, &out, "calculateOperatorTableLeaf", operatorTableBytes)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorTableLeaf is a free data retrieval call binding the contract method 0xa2f2e24d.
//
// Solidity: function calculateOperatorTableLeaf(bytes operatorTableBytes) pure returns(bytes32)
func (_LeafCalculatorMixin *LeafCalculatorMixinSession) CalculateOperatorTableLeaf(operatorTableBytes []byte) ([32]byte, error) {
	return _LeafCalculatorMixin.Contract.CalculateOperatorTableLeaf(&_LeafCalculatorMixin.CallOpts, operatorTableBytes)
}

// CalculateOperatorTableLeaf is a free data retrieval call binding the contract method 0xa2f2e24d.
//
// Solidity: function calculateOperatorTableLeaf(bytes operatorTableBytes) pure returns(bytes32)
func (_LeafCalculatorMixin *LeafCalculatorMixinCallerSession) CalculateOperatorTableLeaf(operatorTableBytes []byte) ([32]byte, error) {
	return _LeafCalculatorMixin.Contract.CalculateOperatorTableLeaf(&_LeafCalculatorMixin.CallOpts, operatorTableBytes)
}
