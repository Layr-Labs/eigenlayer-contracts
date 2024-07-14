// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ISlasher

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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ISlasherMagnitudeAdjustment is an auto generated low-level Go binding around an user-defined struct.
type ISlasherMagnitudeAdjustment struct {
	OperatorSet   ISlasherOperatorSet
	MagnitudeDiff uint64
}

// ISlasherMagnitudeAdjustmentsParam is an auto generated low-level Go binding around an user-defined struct.
type ISlasherMagnitudeAdjustmentsParam struct {
	Strategy                common.Address
	MagnitudeAdjustmentType uint8
	MagnitudeAdjustments    []ISlasherMagnitudeAdjustment
}

// ISlasherOperatorSet is an auto generated low-level Go binding around an user-defined struct.
type ISlasherOperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ISlasherMetaData contains all meta data concerning the ISlasher contract.
var ISlasherMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashedRate\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structISlasher.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueMagnitudeConcentration\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"nonslashableToDecrement\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"allocatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"newNonslashableMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newTotalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queueMagnitudeDilution\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"nonslashableToAdd\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"allocatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"newNonslashableMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newTotalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queueReallocation\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"adjustmentParams\",\"type\":\"tuple[]\",\"internalType\":\"structISlasher.MagnitudeAdjustmentsParam[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"magnitudeAdjustmentType\",\"type\":\"uint8\",\"internalType\":\"enumISlasher.MagnitudeAdjustmentType\"},{\"name\":\"magnitudeAdjustments\",\"type\":\"tuple[]\",\"internalType\":\"structISlasher.MagnitudeAdjustment[]\",\"components\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structISlasher.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeDiff\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]},{\"name\":\"allocatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shareScalingFactor\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shareScalingFactorAtEpoch\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"epoch\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"bipsToSlash\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MagnitudeDecremented\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISlasher.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"updatedSlashableMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISlasher.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"slashableMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonslashableMagnitudeDecremented\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"updatedNonslashableMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISlasher.OperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"bipsToSlash\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"contractIStrategy[]\"},{\"name\":\"slashingRates\",\"type\":\"uint64[]\",\"indexed\":false,\"internalType\":\"uint64[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TotalAndNonslashableMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"totalSlashableMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"nonslashableMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"cumulativeAllocationSum\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false}]",
}

// ISlasherABI is the input ABI used to generate the binding from.
// Deprecated: Use ISlasherMetaData.ABI instead.
var ISlasherABI = ISlasherMetaData.ABI

// ISlasher is an auto generated Go binding around an Ethereum contract.
type ISlasher struct {
	ISlasherCaller     // Read-only binding to the contract
	ISlasherTransactor // Write-only binding to the contract
	ISlasherFilterer   // Log filterer for contract events
}

// ISlasherCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISlasherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlasherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISlasherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlasherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISlasherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlasherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISlasherSession struct {
	Contract     *ISlasher         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISlasherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISlasherCallerSession struct {
	Contract *ISlasherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ISlasherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISlasherTransactorSession struct {
	Contract     *ISlasherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ISlasherRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISlasherRaw struct {
	Contract *ISlasher // Generic contract binding to access the raw methods on
}

// ISlasherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISlasherCallerRaw struct {
	Contract *ISlasherCaller // Generic read-only contract binding to access the raw methods on
}

// ISlasherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISlasherTransactorRaw struct {
	Contract *ISlasherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISlasher creates a new instance of ISlasher, bound to a specific deployed contract.
func NewISlasher(address common.Address, backend bind.ContractBackend) (*ISlasher, error) {
	contract, err := bindISlasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISlasher{ISlasherCaller: ISlasherCaller{contract: contract}, ISlasherTransactor: ISlasherTransactor{contract: contract}, ISlasherFilterer: ISlasherFilterer{contract: contract}}, nil
}

// NewISlasherCaller creates a new read-only instance of ISlasher, bound to a specific deployed contract.
func NewISlasherCaller(address common.Address, caller bind.ContractCaller) (*ISlasherCaller, error) {
	contract, err := bindISlasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISlasherCaller{contract: contract}, nil
}

// NewISlasherTransactor creates a new write-only instance of ISlasher, bound to a specific deployed contract.
func NewISlasherTransactor(address common.Address, transactor bind.ContractTransactor) (*ISlasherTransactor, error) {
	contract, err := bindISlasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISlasherTransactor{contract: contract}, nil
}

// NewISlasherFilterer creates a new log filterer instance of ISlasher, bound to a specific deployed contract.
func NewISlasherFilterer(address common.Address, filterer bind.ContractFilterer) (*ISlasherFilterer, error) {
	contract, err := bindISlasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISlasherFilterer{contract: contract}, nil
}

// bindISlasher binds a generic wrapper to an already deployed contract.
func bindISlasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISlasherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlasher *ISlasherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlasher.Contract.ISlasherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlasher *ISlasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlasher.Contract.ISlasherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlasher *ISlasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlasher.Contract.ISlasherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlasher *ISlasherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlasher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlasher *ISlasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlasher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlasher *ISlasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlasher.Contract.contract.Transact(opts, method, params...)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ISlasher *ISlasherCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ISlasher *ISlasherSession) Delegation() (common.Address, error) {
	return _ISlasher.Contract.Delegation(&_ISlasher.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ISlasher *ISlasherCallerSession) Delegation() (common.Address, error) {
	return _ISlasher.Contract.Delegation(&_ISlasher.CallOpts)
}

// GetSlashedRate is a free data retrieval call binding the contract method 0x10c0a606.
//
// Solidity: function getSlashedRate(address operator, address strategy, (address,uint32) operatorSet, uint32 epoch) view returns(uint64)
func (_ISlasher *ISlasherCaller) GetSlashedRate(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSet ISlasherOperatorSet, epoch uint32) (uint64, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "getSlashedRate", operator, strategy, operatorSet, epoch)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetSlashedRate is a free data retrieval call binding the contract method 0x10c0a606.
//
// Solidity: function getSlashedRate(address operator, address strategy, (address,uint32) operatorSet, uint32 epoch) view returns(uint64)
func (_ISlasher *ISlasherSession) GetSlashedRate(operator common.Address, strategy common.Address, operatorSet ISlasherOperatorSet, epoch uint32) (uint64, error) {
	return _ISlasher.Contract.GetSlashedRate(&_ISlasher.CallOpts, operator, strategy, operatorSet, epoch)
}

// GetSlashedRate is a free data retrieval call binding the contract method 0x10c0a606.
//
// Solidity: function getSlashedRate(address operator, address strategy, (address,uint32) operatorSet, uint32 epoch) view returns(uint64)
func (_ISlasher *ISlasherCallerSession) GetSlashedRate(operator common.Address, strategy common.Address, operatorSet ISlasherOperatorSet, epoch uint32) (uint64, error) {
	return _ISlasher.Contract.GetSlashedRate(&_ISlasher.CallOpts, operator, strategy, operatorSet, epoch)
}

// ShareScalingFactor is a free data retrieval call binding the contract method 0x334f00d6.
//
// Solidity: function shareScalingFactor(address operator, address strategy) view returns(uint64)
func (_ISlasher *ISlasherCaller) ShareScalingFactor(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "shareScalingFactor", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ShareScalingFactor is a free data retrieval call binding the contract method 0x334f00d6.
//
// Solidity: function shareScalingFactor(address operator, address strategy) view returns(uint64)
func (_ISlasher *ISlasherSession) ShareScalingFactor(operator common.Address, strategy common.Address) (uint64, error) {
	return _ISlasher.Contract.ShareScalingFactor(&_ISlasher.CallOpts, operator, strategy)
}

// ShareScalingFactor is a free data retrieval call binding the contract method 0x334f00d6.
//
// Solidity: function shareScalingFactor(address operator, address strategy) view returns(uint64)
func (_ISlasher *ISlasherCallerSession) ShareScalingFactor(operator common.Address, strategy common.Address) (uint64, error) {
	return _ISlasher.Contract.ShareScalingFactor(&_ISlasher.CallOpts, operator, strategy)
}

// ShareScalingFactorAtEpoch is a free data retrieval call binding the contract method 0xe49a1e84.
//
// Solidity: function shareScalingFactorAtEpoch(address operator, address strategy, uint32 epoch) view returns(uint64)
func (_ISlasher *ISlasherCaller) ShareScalingFactorAtEpoch(opts *bind.CallOpts, operator common.Address, strategy common.Address, epoch uint32) (uint64, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "shareScalingFactorAtEpoch", operator, strategy, epoch)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ShareScalingFactorAtEpoch is a free data retrieval call binding the contract method 0xe49a1e84.
//
// Solidity: function shareScalingFactorAtEpoch(address operator, address strategy, uint32 epoch) view returns(uint64)
func (_ISlasher *ISlasherSession) ShareScalingFactorAtEpoch(operator common.Address, strategy common.Address, epoch uint32) (uint64, error) {
	return _ISlasher.Contract.ShareScalingFactorAtEpoch(&_ISlasher.CallOpts, operator, strategy, epoch)
}

// ShareScalingFactorAtEpoch is a free data retrieval call binding the contract method 0xe49a1e84.
//
// Solidity: function shareScalingFactorAtEpoch(address operator, address strategy, uint32 epoch) view returns(uint64)
func (_ISlasher *ISlasherCallerSession) ShareScalingFactorAtEpoch(operator common.Address, strategy common.Address, epoch uint32) (uint64, error) {
	return _ISlasher.Contract.ShareScalingFactorAtEpoch(&_ISlasher.CallOpts, operator, strategy, epoch)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ISlasher *ISlasherCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISlasher.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ISlasher *ISlasherSession) StrategyManager() (common.Address, error) {
	return _ISlasher.Contract.StrategyManager(&_ISlasher.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ISlasher *ISlasherCallerSession) StrategyManager() (common.Address, error) {
	return _ISlasher.Contract.StrategyManager(&_ISlasher.CallOpts)
}

// QueueMagnitudeConcentration is a paid mutator transaction binding the contract method 0xe565c621.
//
// Solidity: function queueMagnitudeConcentration(address operator, address[] strategies, uint64[] nonslashableToDecrement, (bytes,bytes32,uint256) allocatorSignature) returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude)
func (_ISlasher *ISlasherTransactor) QueueMagnitudeConcentration(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, nonslashableToDecrement []uint64, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ISlasher.contract.Transact(opts, "queueMagnitudeConcentration", operator, strategies, nonslashableToDecrement, allocatorSignature)
}

// QueueMagnitudeConcentration is a paid mutator transaction binding the contract method 0xe565c621.
//
// Solidity: function queueMagnitudeConcentration(address operator, address[] strategies, uint64[] nonslashableToDecrement, (bytes,bytes32,uint256) allocatorSignature) returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude)
func (_ISlasher *ISlasherSession) QueueMagnitudeConcentration(operator common.Address, strategies []common.Address, nonslashableToDecrement []uint64, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ISlasher.Contract.QueueMagnitudeConcentration(&_ISlasher.TransactOpts, operator, strategies, nonslashableToDecrement, allocatorSignature)
}

// QueueMagnitudeConcentration is a paid mutator transaction binding the contract method 0xe565c621.
//
// Solidity: function queueMagnitudeConcentration(address operator, address[] strategies, uint64[] nonslashableToDecrement, (bytes,bytes32,uint256) allocatorSignature) returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude)
func (_ISlasher *ISlasherTransactorSession) QueueMagnitudeConcentration(operator common.Address, strategies []common.Address, nonslashableToDecrement []uint64, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ISlasher.Contract.QueueMagnitudeConcentration(&_ISlasher.TransactOpts, operator, strategies, nonslashableToDecrement, allocatorSignature)
}

// QueueMagnitudeDilution is a paid mutator transaction binding the contract method 0x017ec6f9.
//
// Solidity: function queueMagnitudeDilution(address operator, address[] strategies, uint64[] nonslashableToAdd, (bytes,bytes32,uint256) allocatorSignature) returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude)
func (_ISlasher *ISlasherTransactor) QueueMagnitudeDilution(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, nonslashableToAdd []uint64, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ISlasher.contract.Transact(opts, "queueMagnitudeDilution", operator, strategies, nonslashableToAdd, allocatorSignature)
}

// QueueMagnitudeDilution is a paid mutator transaction binding the contract method 0x017ec6f9.
//
// Solidity: function queueMagnitudeDilution(address operator, address[] strategies, uint64[] nonslashableToAdd, (bytes,bytes32,uint256) allocatorSignature) returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude)
func (_ISlasher *ISlasherSession) QueueMagnitudeDilution(operator common.Address, strategies []common.Address, nonslashableToAdd []uint64, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ISlasher.Contract.QueueMagnitudeDilution(&_ISlasher.TransactOpts, operator, strategies, nonslashableToAdd, allocatorSignature)
}

// QueueMagnitudeDilution is a paid mutator transaction binding the contract method 0x017ec6f9.
//
// Solidity: function queueMagnitudeDilution(address operator, address[] strategies, uint64[] nonslashableToAdd, (bytes,bytes32,uint256) allocatorSignature) returns(uint64 newNonslashableMagnitude, uint64 newTotalMagnitude)
func (_ISlasher *ISlasherTransactorSession) QueueMagnitudeDilution(operator common.Address, strategies []common.Address, nonslashableToAdd []uint64, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ISlasher.Contract.QueueMagnitudeDilution(&_ISlasher.TransactOpts, operator, strategies, nonslashableToAdd, allocatorSignature)
}

// QueueReallocation is a paid mutator transaction binding the contract method 0x9292a40d.
//
// Solidity: function queueReallocation(address operator, (address,uint8,((address,uint32),uint64)[])[] adjustmentParams, (bytes,bytes32,uint256) allocatorSignature) returns(uint32 effectTimestamp)
func (_ISlasher *ISlasherTransactor) QueueReallocation(opts *bind.TransactOpts, operator common.Address, adjustmentParams []ISlasherMagnitudeAdjustmentsParam, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ISlasher.contract.Transact(opts, "queueReallocation", operator, adjustmentParams, allocatorSignature)
}

// QueueReallocation is a paid mutator transaction binding the contract method 0x9292a40d.
//
// Solidity: function queueReallocation(address operator, (address,uint8,((address,uint32),uint64)[])[] adjustmentParams, (bytes,bytes32,uint256) allocatorSignature) returns(uint32 effectTimestamp)
func (_ISlasher *ISlasherSession) QueueReallocation(operator common.Address, adjustmentParams []ISlasherMagnitudeAdjustmentsParam, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ISlasher.Contract.QueueReallocation(&_ISlasher.TransactOpts, operator, adjustmentParams, allocatorSignature)
}

// QueueReallocation is a paid mutator transaction binding the contract method 0x9292a40d.
//
// Solidity: function queueReallocation(address operator, (address,uint8,((address,uint32),uint64)[])[] adjustmentParams, (bytes,bytes32,uint256) allocatorSignature) returns(uint32 effectTimestamp)
func (_ISlasher *ISlasherTransactorSession) QueueReallocation(operator common.Address, adjustmentParams []ISlasherMagnitudeAdjustmentsParam, allocatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ISlasher.Contract.QueueReallocation(&_ISlasher.TransactOpts, operator, adjustmentParams, allocatorSignature)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x5c120057.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint32 bipsToSlash) returns()
func (_ISlasher *ISlasherTransactor) SlashOperator(opts *bind.TransactOpts, operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint32) (*types.Transaction, error) {
	return _ISlasher.contract.Transact(opts, "slashOperator", operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x5c120057.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint32 bipsToSlash) returns()
func (_ISlasher *ISlasherSession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint32) (*types.Transaction, error) {
	return _ISlasher.Contract.SlashOperator(&_ISlasher.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x5c120057.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint32 bipsToSlash) returns()
func (_ISlasher *ISlasherTransactorSession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint32) (*types.Transaction, error) {
	return _ISlasher.Contract.SlashOperator(&_ISlasher.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// ISlasherMagnitudeDecrementedIterator is returned from FilterMagnitudeDecremented and is used to iterate over the raw logs and unpacked data for MagnitudeDecremented events raised by the ISlasher contract.
type ISlasherMagnitudeDecrementedIterator struct {
	Event *ISlasherMagnitudeDecremented // Event containing the contract specifics and raw log

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
func (it *ISlasherMagnitudeDecrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlasherMagnitudeDecremented)
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
		it.Event = new(ISlasherMagnitudeDecremented)
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
func (it *ISlasherMagnitudeDecrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlasherMagnitudeDecrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlasherMagnitudeDecremented represents a MagnitudeDecremented event raised by the ISlasher contract.
type ISlasherMagnitudeDecremented struct {
	Operator                  common.Address
	OperatorSet               ISlasherOperatorSet
	Strategy                  common.Address
	UpdatedSlashableMagnitude uint64
	EffectTimestamp           uint32
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeDecremented is a free log retrieval operation binding the contract event 0x543a5d2a6758e26af80ad687683e71136dbf9d41ba11bea98120b3a5be24652d.
//
// Solidity: event MagnitudeDecremented(address operator, (address,uint32) operatorSet, address strategy, uint64 updatedSlashableMagnitude, uint32 effectTimestamp)
func (_ISlasher *ISlasherFilterer) FilterMagnitudeDecremented(opts *bind.FilterOpts) (*ISlasherMagnitudeDecrementedIterator, error) {

	logs, sub, err := _ISlasher.contract.FilterLogs(opts, "MagnitudeDecremented")
	if err != nil {
		return nil, err
	}
	return &ISlasherMagnitudeDecrementedIterator{contract: _ISlasher.contract, event: "MagnitudeDecremented", logs: logs, sub: sub}, nil
}

// WatchMagnitudeDecremented is a free log subscription operation binding the contract event 0x543a5d2a6758e26af80ad687683e71136dbf9d41ba11bea98120b3a5be24652d.
//
// Solidity: event MagnitudeDecremented(address operator, (address,uint32) operatorSet, address strategy, uint64 updatedSlashableMagnitude, uint32 effectTimestamp)
func (_ISlasher *ISlasherFilterer) WatchMagnitudeDecremented(opts *bind.WatchOpts, sink chan<- *ISlasherMagnitudeDecremented) (event.Subscription, error) {

	logs, sub, err := _ISlasher.contract.WatchLogs(opts, "MagnitudeDecremented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlasherMagnitudeDecremented)
				if err := _ISlasher.contract.UnpackLog(event, "MagnitudeDecremented", log); err != nil {
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

// ParseMagnitudeDecremented is a log parse operation binding the contract event 0x543a5d2a6758e26af80ad687683e71136dbf9d41ba11bea98120b3a5be24652d.
//
// Solidity: event MagnitudeDecremented(address operator, (address,uint32) operatorSet, address strategy, uint64 updatedSlashableMagnitude, uint32 effectTimestamp)
func (_ISlasher *ISlasherFilterer) ParseMagnitudeDecremented(log types.Log) (*ISlasherMagnitudeDecremented, error) {
	event := new(ISlasherMagnitudeDecremented)
	if err := _ISlasher.contract.UnpackLog(event, "MagnitudeDecremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlasherMagnitudeUpdatedIterator is returned from FilterMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for MagnitudeUpdated events raised by the ISlasher contract.
type ISlasherMagnitudeUpdatedIterator struct {
	Event *ISlasherMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *ISlasherMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlasherMagnitudeUpdated)
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
		it.Event = new(ISlasherMagnitudeUpdated)
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
func (it *ISlasherMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlasherMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlasherMagnitudeUpdated represents a MagnitudeUpdated event raised by the ISlasher contract.
type ISlasherMagnitudeUpdated struct {
	Operator           common.Address
	Strategy           common.Address
	OperatorSet        ISlasherOperatorSet
	EffectTimestamp    uint32
	SlashableMagnitude uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeUpdated is a free log retrieval operation binding the contract event 0x770e5720664253f46b99dcb3b922235f8ee78191bdb00bf7efb8c48ba3ab396d.
//
// Solidity: event MagnitudeUpdated(address operator, address strategy, (address,uint32) operatorSet, uint32 effectTimestamp, uint64 slashableMagnitude)
func (_ISlasher *ISlasherFilterer) FilterMagnitudeUpdated(opts *bind.FilterOpts) (*ISlasherMagnitudeUpdatedIterator, error) {

	logs, sub, err := _ISlasher.contract.FilterLogs(opts, "MagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &ISlasherMagnitudeUpdatedIterator{contract: _ISlasher.contract, event: "MagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchMagnitudeUpdated is a free log subscription operation binding the contract event 0x770e5720664253f46b99dcb3b922235f8ee78191bdb00bf7efb8c48ba3ab396d.
//
// Solidity: event MagnitudeUpdated(address operator, address strategy, (address,uint32) operatorSet, uint32 effectTimestamp, uint64 slashableMagnitude)
func (_ISlasher *ISlasherFilterer) WatchMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ISlasherMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _ISlasher.contract.WatchLogs(opts, "MagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlasherMagnitudeUpdated)
				if err := _ISlasher.contract.UnpackLog(event, "MagnitudeUpdated", log); err != nil {
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

// ParseMagnitudeUpdated is a log parse operation binding the contract event 0x770e5720664253f46b99dcb3b922235f8ee78191bdb00bf7efb8c48ba3ab396d.
//
// Solidity: event MagnitudeUpdated(address operator, address strategy, (address,uint32) operatorSet, uint32 effectTimestamp, uint64 slashableMagnitude)
func (_ISlasher *ISlasherFilterer) ParseMagnitudeUpdated(log types.Log) (*ISlasherMagnitudeUpdated, error) {
	event := new(ISlasherMagnitudeUpdated)
	if err := _ISlasher.contract.UnpackLog(event, "MagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlasherNonslashableMagnitudeDecrementedIterator is returned from FilterNonslashableMagnitudeDecremented and is used to iterate over the raw logs and unpacked data for NonslashableMagnitudeDecremented events raised by the ISlasher contract.
type ISlasherNonslashableMagnitudeDecrementedIterator struct {
	Event *ISlasherNonslashableMagnitudeDecremented // Event containing the contract specifics and raw log

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
func (it *ISlasherNonslashableMagnitudeDecrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlasherNonslashableMagnitudeDecremented)
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
		it.Event = new(ISlasherNonslashableMagnitudeDecremented)
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
func (it *ISlasherNonslashableMagnitudeDecrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlasherNonslashableMagnitudeDecrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlasherNonslashableMagnitudeDecremented represents a NonslashableMagnitudeDecremented event raised by the ISlasher contract.
type ISlasherNonslashableMagnitudeDecremented struct {
	Operator                     common.Address
	Strategy                     common.Address
	UpdatedNonslashableMagnitude uint64
	EffectTimestamp              uint32
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterNonslashableMagnitudeDecremented is a free log retrieval operation binding the contract event 0x96111829d0d3e9dd48c4e65cecac2a6ce99675f75740a2cbef2a99dca965a713.
//
// Solidity: event NonslashableMagnitudeDecremented(address operator, address strategy, uint64 updatedNonslashableMagnitude, uint32 effectTimestamp)
func (_ISlasher *ISlasherFilterer) FilterNonslashableMagnitudeDecremented(opts *bind.FilterOpts) (*ISlasherNonslashableMagnitudeDecrementedIterator, error) {

	logs, sub, err := _ISlasher.contract.FilterLogs(opts, "NonslashableMagnitudeDecremented")
	if err != nil {
		return nil, err
	}
	return &ISlasherNonslashableMagnitudeDecrementedIterator{contract: _ISlasher.contract, event: "NonslashableMagnitudeDecremented", logs: logs, sub: sub}, nil
}

// WatchNonslashableMagnitudeDecremented is a free log subscription operation binding the contract event 0x96111829d0d3e9dd48c4e65cecac2a6ce99675f75740a2cbef2a99dca965a713.
//
// Solidity: event NonslashableMagnitudeDecremented(address operator, address strategy, uint64 updatedNonslashableMagnitude, uint32 effectTimestamp)
func (_ISlasher *ISlasherFilterer) WatchNonslashableMagnitudeDecremented(opts *bind.WatchOpts, sink chan<- *ISlasherNonslashableMagnitudeDecremented) (event.Subscription, error) {

	logs, sub, err := _ISlasher.contract.WatchLogs(opts, "NonslashableMagnitudeDecremented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlasherNonslashableMagnitudeDecremented)
				if err := _ISlasher.contract.UnpackLog(event, "NonslashableMagnitudeDecremented", log); err != nil {
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

// ParseNonslashableMagnitudeDecremented is a log parse operation binding the contract event 0x96111829d0d3e9dd48c4e65cecac2a6ce99675f75740a2cbef2a99dca965a713.
//
// Solidity: event NonslashableMagnitudeDecremented(address operator, address strategy, uint64 updatedNonslashableMagnitude, uint32 effectTimestamp)
func (_ISlasher *ISlasherFilterer) ParseNonslashableMagnitudeDecremented(log types.Log) (*ISlasherNonslashableMagnitudeDecremented, error) {
	event := new(ISlasherNonslashableMagnitudeDecremented)
	if err := _ISlasher.contract.UnpackLog(event, "NonslashableMagnitudeDecremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlasherOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the ISlasher contract.
type ISlasherOperatorSlashedIterator struct {
	Event *ISlasherOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *ISlasherOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlasherOperatorSlashed)
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
		it.Event = new(ISlasherOperatorSlashed)
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
func (it *ISlasherOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlasherOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlasherOperatorSlashed represents a OperatorSlashed event raised by the ISlasher contract.
type ISlasherOperatorSlashed struct {
	Operator      common.Address
	OperatorSet   ISlasherOperatorSet
	BipsToSlash   uint32
	Strategies    []common.Address
	SlashingRates []uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0x715df5b5f92bd7692569b1129f8a601730c048c3daecf410b60b52d8a6015d40.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, uint32 bipsToSlash, address[] strategies, uint64[] slashingRates)
func (_ISlasher *ISlasherFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*ISlasherOperatorSlashedIterator, error) {

	logs, sub, err := _ISlasher.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &ISlasherOperatorSlashedIterator{contract: _ISlasher.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0x715df5b5f92bd7692569b1129f8a601730c048c3daecf410b60b52d8a6015d40.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, uint32 bipsToSlash, address[] strategies, uint64[] slashingRates)
func (_ISlasher *ISlasherFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *ISlasherOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _ISlasher.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlasherOperatorSlashed)
				if err := _ISlasher.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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

// ParseOperatorSlashed is a log parse operation binding the contract event 0x715df5b5f92bd7692569b1129f8a601730c048c3daecf410b60b52d8a6015d40.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, uint32 bipsToSlash, address[] strategies, uint64[] slashingRates)
func (_ISlasher *ISlasherFilterer) ParseOperatorSlashed(log types.Log) (*ISlasherOperatorSlashed, error) {
	event := new(ISlasherOperatorSlashed)
	if err := _ISlasher.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlasherTotalAndNonslashableMagnitudeUpdatedIterator is returned from FilterTotalAndNonslashableMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for TotalAndNonslashableMagnitudeUpdated events raised by the ISlasher contract.
type ISlasherTotalAndNonslashableMagnitudeUpdatedIterator struct {
	Event *ISlasherTotalAndNonslashableMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *ISlasherTotalAndNonslashableMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlasherTotalAndNonslashableMagnitudeUpdated)
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
		it.Event = new(ISlasherTotalAndNonslashableMagnitudeUpdated)
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
func (it *ISlasherTotalAndNonslashableMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlasherTotalAndNonslashableMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlasherTotalAndNonslashableMagnitudeUpdated represents a TotalAndNonslashableMagnitudeUpdated event raised by the ISlasher contract.
type ISlasherTotalAndNonslashableMagnitudeUpdated struct {
	Operator                common.Address
	Strategy                common.Address
	EffectTimestamp         uint32
	TotalSlashableMagnitude uint64
	NonslashableMagnitude   uint64
	CumulativeAllocationSum uint64
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterTotalAndNonslashableMagnitudeUpdated is a free log retrieval operation binding the contract event 0x2bc1462fd4652ef49aae855563647b6564ff41eb9db10038d4b08e01d34f30e1.
//
// Solidity: event TotalAndNonslashableMagnitudeUpdated(address operator, address strategy, uint32 effectTimestamp, uint64 totalSlashableMagnitude, uint64 nonslashableMagnitude, uint64 cumulativeAllocationSum)
func (_ISlasher *ISlasherFilterer) FilterTotalAndNonslashableMagnitudeUpdated(opts *bind.FilterOpts) (*ISlasherTotalAndNonslashableMagnitudeUpdatedIterator, error) {

	logs, sub, err := _ISlasher.contract.FilterLogs(opts, "TotalAndNonslashableMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &ISlasherTotalAndNonslashableMagnitudeUpdatedIterator{contract: _ISlasher.contract, event: "TotalAndNonslashableMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchTotalAndNonslashableMagnitudeUpdated is a free log subscription operation binding the contract event 0x2bc1462fd4652ef49aae855563647b6564ff41eb9db10038d4b08e01d34f30e1.
//
// Solidity: event TotalAndNonslashableMagnitudeUpdated(address operator, address strategy, uint32 effectTimestamp, uint64 totalSlashableMagnitude, uint64 nonslashableMagnitude, uint64 cumulativeAllocationSum)
func (_ISlasher *ISlasherFilterer) WatchTotalAndNonslashableMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ISlasherTotalAndNonslashableMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _ISlasher.contract.WatchLogs(opts, "TotalAndNonslashableMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlasherTotalAndNonslashableMagnitudeUpdated)
				if err := _ISlasher.contract.UnpackLog(event, "TotalAndNonslashableMagnitudeUpdated", log); err != nil {
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

// ParseTotalAndNonslashableMagnitudeUpdated is a log parse operation binding the contract event 0x2bc1462fd4652ef49aae855563647b6564ff41eb9db10038d4b08e01d34f30e1.
//
// Solidity: event TotalAndNonslashableMagnitudeUpdated(address operator, address strategy, uint32 effectTimestamp, uint64 totalSlashableMagnitude, uint64 nonslashableMagnitude, uint64 cumulativeAllocationSum)
func (_ISlasher *ISlasherFilterer) ParseTotalAndNonslashableMagnitudeUpdated(log types.Log) (*ISlasherTotalAndNonslashableMagnitudeUpdated, error) {
	event := new(ISlasherTotalAndNonslashableMagnitudeUpdated)
	if err := _ISlasher.contract.UnpackLog(event, "TotalAndNonslashableMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
