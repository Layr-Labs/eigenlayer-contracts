// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ISlashEscrowFactory

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

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ISlashEscrowFactoryMetaData contains all meta data concerning the ISlashEscrowFactory contract.
var ISlashEscrowFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"computeSlashEscrowSalt\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getEscrowCompleteBlock\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEscrowStartBlock\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalEscrowDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingEscrows\",\"inputs\":[],\"outputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"isRedistributing\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"slashIds\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"},{\"name\":\"completeBlocks\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingOperatorSets\",\"inputs\":[],\"outputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingSlashIds\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingStrategiesForSlashId\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingStrategiesForSlashIds\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address[][]\",\"internalType\":\"contractIStrategy[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingUnderlyingAmountForStrategy\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlashEscrow\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategyEscrowDelay\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalPendingOperatorSets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalPendingSlashIds\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalPendingStrategiesForSlashId\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialGlobalDelayBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateSlashEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDeployedSlashEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isDeployedSlashEscrow\",\"inputs\":[{\"name\":\"slashEscrow\",\"type\":\"address\",\"internalType\":\"contractISlashEscrow\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isEscrowPaused\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPendingOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPendingSlashId\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseSlashEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseSlashEscrowByStrategy\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalEscrowDelay\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyEscrowDelay\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EscrowComplete\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EscrowPaused\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EscrowUnpaused\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalEscrowDelaySet\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StartEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyEscrowDelaySet\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EscrowDelayNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EscrowNotMature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRedistributionRecipient\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManager\",\"inputs\":[]}]",
}

// ISlashEscrowFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ISlashEscrowFactoryMetaData.ABI instead.
var ISlashEscrowFactoryABI = ISlashEscrowFactoryMetaData.ABI

// ISlashEscrowFactory is an auto generated Go binding around an Ethereum contract.
type ISlashEscrowFactory struct {
	ISlashEscrowFactoryCaller     // Read-only binding to the contract
	ISlashEscrowFactoryTransactor // Write-only binding to the contract
	ISlashEscrowFactoryFilterer   // Log filterer for contract events
}

// ISlashEscrowFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISlashEscrowFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlashEscrowFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISlashEscrowFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlashEscrowFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISlashEscrowFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISlashEscrowFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISlashEscrowFactorySession struct {
	Contract     *ISlashEscrowFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ISlashEscrowFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISlashEscrowFactoryCallerSession struct {
	Contract *ISlashEscrowFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ISlashEscrowFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISlashEscrowFactoryTransactorSession struct {
	Contract     *ISlashEscrowFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ISlashEscrowFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISlashEscrowFactoryRaw struct {
	Contract *ISlashEscrowFactory // Generic contract binding to access the raw methods on
}

// ISlashEscrowFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISlashEscrowFactoryCallerRaw struct {
	Contract *ISlashEscrowFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ISlashEscrowFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISlashEscrowFactoryTransactorRaw struct {
	Contract *ISlashEscrowFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISlashEscrowFactory creates a new instance of ISlashEscrowFactory, bound to a specific deployed contract.
func NewISlashEscrowFactory(address common.Address, backend bind.ContractBackend) (*ISlashEscrowFactory, error) {
	contract, err := bindISlashEscrowFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISlashEscrowFactory{ISlashEscrowFactoryCaller: ISlashEscrowFactoryCaller{contract: contract}, ISlashEscrowFactoryTransactor: ISlashEscrowFactoryTransactor{contract: contract}, ISlashEscrowFactoryFilterer: ISlashEscrowFactoryFilterer{contract: contract}}, nil
}

// NewISlashEscrowFactoryCaller creates a new read-only instance of ISlashEscrowFactory, bound to a specific deployed contract.
func NewISlashEscrowFactoryCaller(address common.Address, caller bind.ContractCaller) (*ISlashEscrowFactoryCaller, error) {
	contract, err := bindISlashEscrowFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISlashEscrowFactoryCaller{contract: contract}, nil
}

// NewISlashEscrowFactoryTransactor creates a new write-only instance of ISlashEscrowFactory, bound to a specific deployed contract.
func NewISlashEscrowFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ISlashEscrowFactoryTransactor, error) {
	contract, err := bindISlashEscrowFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISlashEscrowFactoryTransactor{contract: contract}, nil
}

// NewISlashEscrowFactoryFilterer creates a new log filterer instance of ISlashEscrowFactory, bound to a specific deployed contract.
func NewISlashEscrowFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ISlashEscrowFactoryFilterer, error) {
	contract, err := bindISlashEscrowFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISlashEscrowFactoryFilterer{contract: contract}, nil
}

// bindISlashEscrowFactory binds a generic wrapper to an already deployed contract.
func bindISlashEscrowFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISlashEscrowFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlashEscrowFactory *ISlashEscrowFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlashEscrowFactory.Contract.ISlashEscrowFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlashEscrowFactory *ISlashEscrowFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.ISlashEscrowFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlashEscrowFactory *ISlashEscrowFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.ISlashEscrowFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISlashEscrowFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.contract.Transact(opts, method, params...)
}

// ComputeSlashEscrowSalt is a free data retrieval call binding the contract method 0xc2de7096.
//
// Solidity: function computeSlashEscrowSalt((address,uint32) operatorSet, uint256 slashId) pure returns(bytes32)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) ComputeSlashEscrowSalt(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "computeSlashEscrowSalt", operatorSet, slashId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeSlashEscrowSalt is a free data retrieval call binding the contract method 0xc2de7096.
//
// Solidity: function computeSlashEscrowSalt((address,uint32) operatorSet, uint256 slashId) pure returns(bytes32)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) ComputeSlashEscrowSalt(operatorSet OperatorSet, slashId *big.Int) ([32]byte, error) {
	return _ISlashEscrowFactory.Contract.ComputeSlashEscrowSalt(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// ComputeSlashEscrowSalt is a free data retrieval call binding the contract method 0xc2de7096.
//
// Solidity: function computeSlashEscrowSalt((address,uint32) operatorSet, uint256 slashId) pure returns(bytes32)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) ComputeSlashEscrowSalt(operatorSet OperatorSet, slashId *big.Int) ([32]byte, error) {
	return _ISlashEscrowFactory.Contract.ComputeSlashEscrowSalt(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetEscrowCompleteBlock is a free data retrieval call binding the contract method 0x0310f3e6.
//
// Solidity: function getEscrowCompleteBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint32)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) GetEscrowCompleteBlock(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (uint32, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "getEscrowCompleteBlock", operatorSet, slashId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetEscrowCompleteBlock is a free data retrieval call binding the contract method 0x0310f3e6.
//
// Solidity: function getEscrowCompleteBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint32)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) GetEscrowCompleteBlock(operatorSet OperatorSet, slashId *big.Int) (uint32, error) {
	return _ISlashEscrowFactory.Contract.GetEscrowCompleteBlock(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetEscrowCompleteBlock is a free data retrieval call binding the contract method 0x0310f3e6.
//
// Solidity: function getEscrowCompleteBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint32)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) GetEscrowCompleteBlock(operatorSet OperatorSet, slashId *big.Int) (uint32, error) {
	return _ISlashEscrowFactory.Contract.GetEscrowCompleteBlock(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetEscrowStartBlock is a free data retrieval call binding the contract method 0x05a4dfbb.
//
// Solidity: function getEscrowStartBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) GetEscrowStartBlock(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "getEscrowStartBlock", operatorSet, slashId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEscrowStartBlock is a free data retrieval call binding the contract method 0x05a4dfbb.
//
// Solidity: function getEscrowStartBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) GetEscrowStartBlock(operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	return _ISlashEscrowFactory.Contract.GetEscrowStartBlock(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetEscrowStartBlock is a free data retrieval call binding the contract method 0x05a4dfbb.
//
// Solidity: function getEscrowStartBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) GetEscrowStartBlock(operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	return _ISlashEscrowFactory.Contract.GetEscrowStartBlock(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetGlobalEscrowDelay is a free data retrieval call binding the contract method 0x8a65d2d2.
//
// Solidity: function getGlobalEscrowDelay() view returns(uint32)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) GetGlobalEscrowDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "getGlobalEscrowDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetGlobalEscrowDelay is a free data retrieval call binding the contract method 0x8a65d2d2.
//
// Solidity: function getGlobalEscrowDelay() view returns(uint32)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) GetGlobalEscrowDelay() (uint32, error) {
	return _ISlashEscrowFactory.Contract.GetGlobalEscrowDelay(&_ISlashEscrowFactory.CallOpts)
}

// GetGlobalEscrowDelay is a free data retrieval call binding the contract method 0x8a65d2d2.
//
// Solidity: function getGlobalEscrowDelay() view returns(uint32)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) GetGlobalEscrowDelay() (uint32, error) {
	return _ISlashEscrowFactory.Contract.GetGlobalEscrowDelay(&_ISlashEscrowFactory.CallOpts)
}

// GetPendingEscrows is a free data retrieval call binding the contract method 0xc50f4e48.
//
// Solidity: function getPendingEscrows() view returns((address,uint32)[] operatorSets, bool[] isRedistributing, uint256[][] slashIds, uint32[][] completeBlocks)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) GetPendingEscrows(opts *bind.CallOpts) (struct {
	OperatorSets     []OperatorSet
	IsRedistributing []bool
	SlashIds         [][]*big.Int
	CompleteBlocks   [][]uint32
}, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "getPendingEscrows")

	outstruct := new(struct {
		OperatorSets     []OperatorSet
		IsRedistributing []bool
		SlashIds         [][]*big.Int
		CompleteBlocks   [][]uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OperatorSets = *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)
	outstruct.IsRedistributing = *abi.ConvertType(out[1], new([]bool)).(*[]bool)
	outstruct.SlashIds = *abi.ConvertType(out[2], new([][]*big.Int)).(*[][]*big.Int)
	outstruct.CompleteBlocks = *abi.ConvertType(out[3], new([][]uint32)).(*[][]uint32)

	return *outstruct, err

}

// GetPendingEscrows is a free data retrieval call binding the contract method 0xc50f4e48.
//
// Solidity: function getPendingEscrows() view returns((address,uint32)[] operatorSets, bool[] isRedistributing, uint256[][] slashIds, uint32[][] completeBlocks)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) GetPendingEscrows() (struct {
	OperatorSets     []OperatorSet
	IsRedistributing []bool
	SlashIds         [][]*big.Int
	CompleteBlocks   [][]uint32
}, error) {
	return _ISlashEscrowFactory.Contract.GetPendingEscrows(&_ISlashEscrowFactory.CallOpts)
}

// GetPendingEscrows is a free data retrieval call binding the contract method 0xc50f4e48.
//
// Solidity: function getPendingEscrows() view returns((address,uint32)[] operatorSets, bool[] isRedistributing, uint256[][] slashIds, uint32[][] completeBlocks)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) GetPendingEscrows() (struct {
	OperatorSets     []OperatorSet
	IsRedistributing []bool
	SlashIds         [][]*big.Int
	CompleteBlocks   [][]uint32
}, error) {
	return _ISlashEscrowFactory.Contract.GetPendingEscrows(&_ISlashEscrowFactory.CallOpts)
}

// GetPendingOperatorSets is a free data retrieval call binding the contract method 0x3f292b08.
//
// Solidity: function getPendingOperatorSets() view returns((address,uint32)[] operatorSets)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) GetPendingOperatorSets(opts *bind.CallOpts) ([]OperatorSet, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "getPendingOperatorSets")

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetPendingOperatorSets is a free data retrieval call binding the contract method 0x3f292b08.
//
// Solidity: function getPendingOperatorSets() view returns((address,uint32)[] operatorSets)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) GetPendingOperatorSets() ([]OperatorSet, error) {
	return _ISlashEscrowFactory.Contract.GetPendingOperatorSets(&_ISlashEscrowFactory.CallOpts)
}

// GetPendingOperatorSets is a free data retrieval call binding the contract method 0x3f292b08.
//
// Solidity: function getPendingOperatorSets() view returns((address,uint32)[] operatorSets)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) GetPendingOperatorSets() ([]OperatorSet, error) {
	return _ISlashEscrowFactory.Contract.GetPendingOperatorSets(&_ISlashEscrowFactory.CallOpts)
}

// GetPendingSlashIds is a free data retrieval call binding the contract method 0x7def1564.
//
// Solidity: function getPendingSlashIds((address,uint32) operatorSet) view returns(uint256[])
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) GetPendingSlashIds(opts *bind.CallOpts, operatorSet OperatorSet) ([]*big.Int, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "getPendingSlashIds", operatorSet)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPendingSlashIds is a free data retrieval call binding the contract method 0x7def1564.
//
// Solidity: function getPendingSlashIds((address,uint32) operatorSet) view returns(uint256[])
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) GetPendingSlashIds(operatorSet OperatorSet) ([]*big.Int, error) {
	return _ISlashEscrowFactory.Contract.GetPendingSlashIds(&_ISlashEscrowFactory.CallOpts, operatorSet)
}

// GetPendingSlashIds is a free data retrieval call binding the contract method 0x7def1564.
//
// Solidity: function getPendingSlashIds((address,uint32) operatorSet) view returns(uint256[])
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) GetPendingSlashIds(operatorSet OperatorSet) ([]*big.Int, error) {
	return _ISlashEscrowFactory.Contract.GetPendingSlashIds(&_ISlashEscrowFactory.CallOpts, operatorSet)
}

// GetPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x7130c423.
//
// Solidity: function getPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(address[] strategies)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) GetPendingStrategiesForSlashId(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "getPendingStrategiesForSlashId", operatorSet, slashId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x7130c423.
//
// Solidity: function getPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(address[] strategies)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) GetPendingStrategiesForSlashId(operatorSet OperatorSet, slashId *big.Int) ([]common.Address, error) {
	return _ISlashEscrowFactory.Contract.GetPendingStrategiesForSlashId(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x7130c423.
//
// Solidity: function getPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(address[] strategies)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) GetPendingStrategiesForSlashId(operatorSet OperatorSet, slashId *big.Int) ([]common.Address, error) {
	return _ISlashEscrowFactory.Contract.GetPendingStrategiesForSlashId(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetPendingStrategiesForSlashIds is a free data retrieval call binding the contract method 0xa56b21e4.
//
// Solidity: function getPendingStrategiesForSlashIds((address,uint32) operatorSet) view returns(address[][] strategies)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) GetPendingStrategiesForSlashIds(opts *bind.CallOpts, operatorSet OperatorSet) ([][]common.Address, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "getPendingStrategiesForSlashIds", operatorSet)

	if err != nil {
		return *new([][]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([][]common.Address)).(*[][]common.Address)

	return out0, err

}

// GetPendingStrategiesForSlashIds is a free data retrieval call binding the contract method 0xa56b21e4.
//
// Solidity: function getPendingStrategiesForSlashIds((address,uint32) operatorSet) view returns(address[][] strategies)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) GetPendingStrategiesForSlashIds(operatorSet OperatorSet) ([][]common.Address, error) {
	return _ISlashEscrowFactory.Contract.GetPendingStrategiesForSlashIds(&_ISlashEscrowFactory.CallOpts, operatorSet)
}

// GetPendingStrategiesForSlashIds is a free data retrieval call binding the contract method 0xa56b21e4.
//
// Solidity: function getPendingStrategiesForSlashIds((address,uint32) operatorSet) view returns(address[][] strategies)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) GetPendingStrategiesForSlashIds(operatorSet OperatorSet) ([][]common.Address, error) {
	return _ISlashEscrowFactory.Contract.GetPendingStrategiesForSlashIds(&_ISlashEscrowFactory.CallOpts, operatorSet)
}

// GetPendingUnderlyingAmountForStrategy is a free data retrieval call binding the contract method 0xb23ff83b.
//
// Solidity: function getPendingUnderlyingAmountForStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) view returns(uint256)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) GetPendingUnderlyingAmountForStrategy(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "getPendingUnderlyingAmountForStrategy", operatorSet, slashId, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingUnderlyingAmountForStrategy is a free data retrieval call binding the contract method 0xb23ff83b.
//
// Solidity: function getPendingUnderlyingAmountForStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) view returns(uint256)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) GetPendingUnderlyingAmountForStrategy(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*big.Int, error) {
	return _ISlashEscrowFactory.Contract.GetPendingUnderlyingAmountForStrategy(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId, strategy)
}

// GetPendingUnderlyingAmountForStrategy is a free data retrieval call binding the contract method 0xb23ff83b.
//
// Solidity: function getPendingUnderlyingAmountForStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) view returns(uint256)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) GetPendingUnderlyingAmountForStrategy(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*big.Int, error) {
	return _ISlashEscrowFactory.Contract.GetPendingUnderlyingAmountForStrategy(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId, strategy)
}

// GetSlashEscrow is a free data retrieval call binding the contract method 0x3453b234.
//
// Solidity: function getSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(address)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) GetSlashEscrow(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "getSlashEscrow", operatorSet, slashId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSlashEscrow is a free data retrieval call binding the contract method 0x3453b234.
//
// Solidity: function getSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(address)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) GetSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (common.Address, error) {
	return _ISlashEscrowFactory.Contract.GetSlashEscrow(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetSlashEscrow is a free data retrieval call binding the contract method 0x3453b234.
//
// Solidity: function getSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(address)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) GetSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (common.Address, error) {
	return _ISlashEscrowFactory.Contract.GetSlashEscrow(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetStrategyEscrowDelay is a free data retrieval call binding the contract method 0x8d5d4036.
//
// Solidity: function getStrategyEscrowDelay(address strategy) view returns(uint32)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) GetStrategyEscrowDelay(opts *bind.CallOpts, strategy common.Address) (uint32, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "getStrategyEscrowDelay", strategy)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetStrategyEscrowDelay is a free data retrieval call binding the contract method 0x8d5d4036.
//
// Solidity: function getStrategyEscrowDelay(address strategy) view returns(uint32)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) GetStrategyEscrowDelay(strategy common.Address) (uint32, error) {
	return _ISlashEscrowFactory.Contract.GetStrategyEscrowDelay(&_ISlashEscrowFactory.CallOpts, strategy)
}

// GetStrategyEscrowDelay is a free data retrieval call binding the contract method 0x8d5d4036.
//
// Solidity: function getStrategyEscrowDelay(address strategy) view returns(uint32)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) GetStrategyEscrowDelay(strategy common.Address) (uint32, error) {
	return _ISlashEscrowFactory.Contract.GetStrategyEscrowDelay(&_ISlashEscrowFactory.CallOpts, strategy)
}

// GetTotalPendingOperatorSets is a free data retrieval call binding the contract method 0x78cb9600.
//
// Solidity: function getTotalPendingOperatorSets() view returns(uint256)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) GetTotalPendingOperatorSets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "getTotalPendingOperatorSets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPendingOperatorSets is a free data retrieval call binding the contract method 0x78cb9600.
//
// Solidity: function getTotalPendingOperatorSets() view returns(uint256)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) GetTotalPendingOperatorSets() (*big.Int, error) {
	return _ISlashEscrowFactory.Contract.GetTotalPendingOperatorSets(&_ISlashEscrowFactory.CallOpts)
}

// GetTotalPendingOperatorSets is a free data retrieval call binding the contract method 0x78cb9600.
//
// Solidity: function getTotalPendingOperatorSets() view returns(uint256)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) GetTotalPendingOperatorSets() (*big.Int, error) {
	return _ISlashEscrowFactory.Contract.GetTotalPendingOperatorSets(&_ISlashEscrowFactory.CallOpts)
}

// GetTotalPendingSlashIds is a free data retrieval call binding the contract method 0x6729b5db.
//
// Solidity: function getTotalPendingSlashIds((address,uint32) operatorSet) view returns(uint256)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) GetTotalPendingSlashIds(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "getTotalPendingSlashIds", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPendingSlashIds is a free data retrieval call binding the contract method 0x6729b5db.
//
// Solidity: function getTotalPendingSlashIds((address,uint32) operatorSet) view returns(uint256)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) GetTotalPendingSlashIds(operatorSet OperatorSet) (*big.Int, error) {
	return _ISlashEscrowFactory.Contract.GetTotalPendingSlashIds(&_ISlashEscrowFactory.CallOpts, operatorSet)
}

// GetTotalPendingSlashIds is a free data retrieval call binding the contract method 0x6729b5db.
//
// Solidity: function getTotalPendingSlashIds((address,uint32) operatorSet) view returns(uint256)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) GetTotalPendingSlashIds(operatorSet OperatorSet) (*big.Int, error) {
	return _ISlashEscrowFactory.Contract.GetTotalPendingSlashIds(&_ISlashEscrowFactory.CallOpts, operatorSet)
}

// GetTotalPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x277a9f0e.
//
// Solidity: function getTotalPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) GetTotalPendingStrategiesForSlashId(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "getTotalPendingStrategiesForSlashId", operatorSet, slashId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x277a9f0e.
//
// Solidity: function getTotalPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) GetTotalPendingStrategiesForSlashId(operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	return _ISlashEscrowFactory.Contract.GetTotalPendingStrategiesForSlashId(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// GetTotalPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x277a9f0e.
//
// Solidity: function getTotalPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) GetTotalPendingStrategiesForSlashId(operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	return _ISlashEscrowFactory.Contract.GetTotalPendingStrategiesForSlashId(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// IsDeployedSlashEscrow is a free data retrieval call binding the contract method 0x19f3db26.
//
// Solidity: function isDeployedSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) IsDeployedSlashEscrow(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "isDeployedSlashEscrow", operatorSet, slashId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeployedSlashEscrow is a free data retrieval call binding the contract method 0x19f3db26.
//
// Solidity: function isDeployedSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) IsDeployedSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _ISlashEscrowFactory.Contract.IsDeployedSlashEscrow(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// IsDeployedSlashEscrow is a free data retrieval call binding the contract method 0x19f3db26.
//
// Solidity: function isDeployedSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) IsDeployedSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _ISlashEscrowFactory.Contract.IsDeployedSlashEscrow(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// IsDeployedSlashEscrow0 is a free data retrieval call binding the contract method 0xc8b5330c.
//
// Solidity: function isDeployedSlashEscrow(address slashEscrow) view returns(bool)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) IsDeployedSlashEscrow0(opts *bind.CallOpts, slashEscrow common.Address) (bool, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "isDeployedSlashEscrow0", slashEscrow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeployedSlashEscrow0 is a free data retrieval call binding the contract method 0xc8b5330c.
//
// Solidity: function isDeployedSlashEscrow(address slashEscrow) view returns(bool)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) IsDeployedSlashEscrow0(slashEscrow common.Address) (bool, error) {
	return _ISlashEscrowFactory.Contract.IsDeployedSlashEscrow0(&_ISlashEscrowFactory.CallOpts, slashEscrow)
}

// IsDeployedSlashEscrow0 is a free data retrieval call binding the contract method 0xc8b5330c.
//
// Solidity: function isDeployedSlashEscrow(address slashEscrow) view returns(bool)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) IsDeployedSlashEscrow0(slashEscrow common.Address) (bool, error) {
	return _ISlashEscrowFactory.Contract.IsDeployedSlashEscrow0(&_ISlashEscrowFactory.CallOpts, slashEscrow)
}

// IsEscrowPaused is a free data retrieval call binding the contract method 0x78748459.
//
// Solidity: function isEscrowPaused((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) IsEscrowPaused(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "isEscrowPaused", operatorSet, slashId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEscrowPaused is a free data retrieval call binding the contract method 0x78748459.
//
// Solidity: function isEscrowPaused((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) IsEscrowPaused(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _ISlashEscrowFactory.Contract.IsEscrowPaused(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// IsEscrowPaused is a free data retrieval call binding the contract method 0x78748459.
//
// Solidity: function isEscrowPaused((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) IsEscrowPaused(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _ISlashEscrowFactory.Contract.IsEscrowPaused(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// IsPendingOperatorSet is a free data retrieval call binding the contract method 0x71e166e7.
//
// Solidity: function isPendingOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) IsPendingOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "isPendingOperatorSet", operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingOperatorSet is a free data retrieval call binding the contract method 0x71e166e7.
//
// Solidity: function isPendingOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) IsPendingOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _ISlashEscrowFactory.Contract.IsPendingOperatorSet(&_ISlashEscrowFactory.CallOpts, operatorSet)
}

// IsPendingOperatorSet is a free data retrieval call binding the contract method 0x71e166e7.
//
// Solidity: function isPendingOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) IsPendingOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _ISlashEscrowFactory.Contract.IsPendingOperatorSet(&_ISlashEscrowFactory.CallOpts, operatorSet)
}

// IsPendingSlashId is a free data retrieval call binding the contract method 0x87420b07.
//
// Solidity: function isPendingSlashId((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCaller) IsPendingSlashId(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	var out []interface{}
	err := _ISlashEscrowFactory.contract.Call(opts, &out, "isPendingSlashId", operatorSet, slashId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingSlashId is a free data retrieval call binding the contract method 0x87420b07.
//
// Solidity: function isPendingSlashId((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) IsPendingSlashId(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _ISlashEscrowFactory.Contract.IsPendingSlashId(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// IsPendingSlashId is a free data retrieval call binding the contract method 0x87420b07.
//
// Solidity: function isPendingSlashId((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_ISlashEscrowFactory *ISlashEscrowFactoryCallerSession) IsPendingSlashId(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _ISlashEscrowFactory.Contract.IsPendingSlashId(&_ISlashEscrowFactory.CallOpts, operatorSet, slashId)
}

// Initialize is a paid mutator transaction binding the contract method 0x6c5ac81e.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus, uint32 initialGlobalDelayBlocks) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int, initialGlobalDelayBlocks uint32) (*types.Transaction, error) {
	return _ISlashEscrowFactory.contract.Transact(opts, "initialize", initialOwner, initialPausedStatus, initialGlobalDelayBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0x6c5ac81e.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus, uint32 initialGlobalDelayBlocks) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int, initialGlobalDelayBlocks uint32) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.Initialize(&_ISlashEscrowFactory.TransactOpts, initialOwner, initialPausedStatus, initialGlobalDelayBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0x6c5ac81e.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus, uint32 initialGlobalDelayBlocks) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactorSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int, initialGlobalDelayBlocks uint32) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.Initialize(&_ISlashEscrowFactory.TransactOpts, initialOwner, initialPausedStatus, initialGlobalDelayBlocks)
}

// InitiateSlashEscrow is a paid mutator transaction binding the contract method 0x7a967611.
//
// Solidity: function initiateSlashEscrow((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactor) InitiateSlashEscrow(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _ISlashEscrowFactory.contract.Transact(opts, "initiateSlashEscrow", operatorSet, slashId, strategy)
}

// InitiateSlashEscrow is a paid mutator transaction binding the contract method 0x7a967611.
//
// Solidity: function initiateSlashEscrow((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) InitiateSlashEscrow(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.InitiateSlashEscrow(&_ISlashEscrowFactory.TransactOpts, operatorSet, slashId, strategy)
}

// InitiateSlashEscrow is a paid mutator transaction binding the contract method 0x7a967611.
//
// Solidity: function initiateSlashEscrow((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactorSession) InitiateSlashEscrow(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.InitiateSlashEscrow(&_ISlashEscrowFactory.TransactOpts, operatorSet, slashId, strategy)
}

// PauseEscrow is a paid mutator transaction binding the contract method 0xe7ed076b.
//
// Solidity: function pauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactor) PauseEscrow(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _ISlashEscrowFactory.contract.Transact(opts, "pauseEscrow", operatorSet, slashId)
}

// PauseEscrow is a paid mutator transaction binding the contract method 0xe7ed076b.
//
// Solidity: function pauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) PauseEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.PauseEscrow(&_ISlashEscrowFactory.TransactOpts, operatorSet, slashId)
}

// PauseEscrow is a paid mutator transaction binding the contract method 0xe7ed076b.
//
// Solidity: function pauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactorSession) PauseEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.PauseEscrow(&_ISlashEscrowFactory.TransactOpts, operatorSet, slashId)
}

// ReleaseSlashEscrow is a paid mutator transaction binding the contract method 0x5ffa5a81.
//
// Solidity: function releaseSlashEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactor) ReleaseSlashEscrow(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _ISlashEscrowFactory.contract.Transact(opts, "releaseSlashEscrow", operatorSet, slashId)
}

// ReleaseSlashEscrow is a paid mutator transaction binding the contract method 0x5ffa5a81.
//
// Solidity: function releaseSlashEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) ReleaseSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.ReleaseSlashEscrow(&_ISlashEscrowFactory.TransactOpts, operatorSet, slashId)
}

// ReleaseSlashEscrow is a paid mutator transaction binding the contract method 0x5ffa5a81.
//
// Solidity: function releaseSlashEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactorSession) ReleaseSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.ReleaseSlashEscrow(&_ISlashEscrowFactory.TransactOpts, operatorSet, slashId)
}

// ReleaseSlashEscrowByStrategy is a paid mutator transaction binding the contract method 0x0e475b17.
//
// Solidity: function releaseSlashEscrowByStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactor) ReleaseSlashEscrowByStrategy(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _ISlashEscrowFactory.contract.Transact(opts, "releaseSlashEscrowByStrategy", operatorSet, slashId, strategy)
}

// ReleaseSlashEscrowByStrategy is a paid mutator transaction binding the contract method 0x0e475b17.
//
// Solidity: function releaseSlashEscrowByStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) ReleaseSlashEscrowByStrategy(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.ReleaseSlashEscrowByStrategy(&_ISlashEscrowFactory.TransactOpts, operatorSet, slashId, strategy)
}

// ReleaseSlashEscrowByStrategy is a paid mutator transaction binding the contract method 0x0e475b17.
//
// Solidity: function releaseSlashEscrowByStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactorSession) ReleaseSlashEscrowByStrategy(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.ReleaseSlashEscrowByStrategy(&_ISlashEscrowFactory.TransactOpts, operatorSet, slashId, strategy)
}

// SetGlobalEscrowDelay is a paid mutator transaction binding the contract method 0x5e0a64c5.
//
// Solidity: function setGlobalEscrowDelay(uint32 delay) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactor) SetGlobalEscrowDelay(opts *bind.TransactOpts, delay uint32) (*types.Transaction, error) {
	return _ISlashEscrowFactory.contract.Transact(opts, "setGlobalEscrowDelay", delay)
}

// SetGlobalEscrowDelay is a paid mutator transaction binding the contract method 0x5e0a64c5.
//
// Solidity: function setGlobalEscrowDelay(uint32 delay) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) SetGlobalEscrowDelay(delay uint32) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.SetGlobalEscrowDelay(&_ISlashEscrowFactory.TransactOpts, delay)
}

// SetGlobalEscrowDelay is a paid mutator transaction binding the contract method 0x5e0a64c5.
//
// Solidity: function setGlobalEscrowDelay(uint32 delay) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactorSession) SetGlobalEscrowDelay(delay uint32) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.SetGlobalEscrowDelay(&_ISlashEscrowFactory.TransactOpts, delay)
}

// SetStrategyEscrowDelay is a paid mutator transaction binding the contract method 0x8fc46be5.
//
// Solidity: function setStrategyEscrowDelay(address strategy, uint32 delay) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactor) SetStrategyEscrowDelay(opts *bind.TransactOpts, strategy common.Address, delay uint32) (*types.Transaction, error) {
	return _ISlashEscrowFactory.contract.Transact(opts, "setStrategyEscrowDelay", strategy, delay)
}

// SetStrategyEscrowDelay is a paid mutator transaction binding the contract method 0x8fc46be5.
//
// Solidity: function setStrategyEscrowDelay(address strategy, uint32 delay) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) SetStrategyEscrowDelay(strategy common.Address, delay uint32) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.SetStrategyEscrowDelay(&_ISlashEscrowFactory.TransactOpts, strategy, delay)
}

// SetStrategyEscrowDelay is a paid mutator transaction binding the contract method 0x8fc46be5.
//
// Solidity: function setStrategyEscrowDelay(address strategy, uint32 delay) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactorSession) SetStrategyEscrowDelay(strategy common.Address, delay uint32) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.SetStrategyEscrowDelay(&_ISlashEscrowFactory.TransactOpts, strategy, delay)
}

// UnpauseEscrow is a paid mutator transaction binding the contract method 0x9b122356.
//
// Solidity: function unpauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactor) UnpauseEscrow(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _ISlashEscrowFactory.contract.Transact(opts, "unpauseEscrow", operatorSet, slashId)
}

// UnpauseEscrow is a paid mutator transaction binding the contract method 0x9b122356.
//
// Solidity: function unpauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactorySession) UnpauseEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.UnpauseEscrow(&_ISlashEscrowFactory.TransactOpts, operatorSet, slashId)
}

// UnpauseEscrow is a paid mutator transaction binding the contract method 0x9b122356.
//
// Solidity: function unpauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_ISlashEscrowFactory *ISlashEscrowFactoryTransactorSession) UnpauseEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _ISlashEscrowFactory.Contract.UnpauseEscrow(&_ISlashEscrowFactory.TransactOpts, operatorSet, slashId)
}

// ISlashEscrowFactoryEscrowCompleteIterator is returned from FilterEscrowComplete and is used to iterate over the raw logs and unpacked data for EscrowComplete events raised by the ISlashEscrowFactory contract.
type ISlashEscrowFactoryEscrowCompleteIterator struct {
	Event *ISlashEscrowFactoryEscrowComplete // Event containing the contract specifics and raw log

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
func (it *ISlashEscrowFactoryEscrowCompleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlashEscrowFactoryEscrowComplete)
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
		it.Event = new(ISlashEscrowFactoryEscrowComplete)
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
func (it *ISlashEscrowFactoryEscrowCompleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlashEscrowFactoryEscrowCompleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlashEscrowFactoryEscrowComplete represents a EscrowComplete event raised by the ISlashEscrowFactory contract.
type ISlashEscrowFactoryEscrowComplete struct {
	OperatorSet OperatorSet
	SlashId     *big.Int
	Strategy    common.Address
	Recipient   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEscrowComplete is a free log retrieval operation binding the contract event 0x32be306ad5a833e756b7cb9724d5312afe0feda6163bfc2dd98ee713346a9abc.
//
// Solidity: event EscrowComplete((address,uint32) operatorSet, uint256 slashId, address strategy, address recipient)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) FilterEscrowComplete(opts *bind.FilterOpts) (*ISlashEscrowFactoryEscrowCompleteIterator, error) {

	logs, sub, err := _ISlashEscrowFactory.contract.FilterLogs(opts, "EscrowComplete")
	if err != nil {
		return nil, err
	}
	return &ISlashEscrowFactoryEscrowCompleteIterator{contract: _ISlashEscrowFactory.contract, event: "EscrowComplete", logs: logs, sub: sub}, nil
}

// WatchEscrowComplete is a free log subscription operation binding the contract event 0x32be306ad5a833e756b7cb9724d5312afe0feda6163bfc2dd98ee713346a9abc.
//
// Solidity: event EscrowComplete((address,uint32) operatorSet, uint256 slashId, address strategy, address recipient)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) WatchEscrowComplete(opts *bind.WatchOpts, sink chan<- *ISlashEscrowFactoryEscrowComplete) (event.Subscription, error) {

	logs, sub, err := _ISlashEscrowFactory.contract.WatchLogs(opts, "EscrowComplete")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlashEscrowFactoryEscrowComplete)
				if err := _ISlashEscrowFactory.contract.UnpackLog(event, "EscrowComplete", log); err != nil {
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

// ParseEscrowComplete is a log parse operation binding the contract event 0x32be306ad5a833e756b7cb9724d5312afe0feda6163bfc2dd98ee713346a9abc.
//
// Solidity: event EscrowComplete((address,uint32) operatorSet, uint256 slashId, address strategy, address recipient)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) ParseEscrowComplete(log types.Log) (*ISlashEscrowFactoryEscrowComplete, error) {
	event := new(ISlashEscrowFactoryEscrowComplete)
	if err := _ISlashEscrowFactory.contract.UnpackLog(event, "EscrowComplete", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlashEscrowFactoryEscrowPausedIterator is returned from FilterEscrowPaused and is used to iterate over the raw logs and unpacked data for EscrowPaused events raised by the ISlashEscrowFactory contract.
type ISlashEscrowFactoryEscrowPausedIterator struct {
	Event *ISlashEscrowFactoryEscrowPaused // Event containing the contract specifics and raw log

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
func (it *ISlashEscrowFactoryEscrowPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlashEscrowFactoryEscrowPaused)
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
		it.Event = new(ISlashEscrowFactoryEscrowPaused)
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
func (it *ISlashEscrowFactoryEscrowPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlashEscrowFactoryEscrowPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlashEscrowFactoryEscrowPaused represents a EscrowPaused event raised by the ISlashEscrowFactory contract.
type ISlashEscrowFactoryEscrowPaused struct {
	OperatorSet OperatorSet
	SlashId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEscrowPaused is a free log retrieval operation binding the contract event 0x050add19b1a78a4240cdebc8747899275f2dd070c88e83904a37ff7d1a539744.
//
// Solidity: event EscrowPaused((address,uint32) operatorSet, uint256 slashId)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) FilterEscrowPaused(opts *bind.FilterOpts) (*ISlashEscrowFactoryEscrowPausedIterator, error) {

	logs, sub, err := _ISlashEscrowFactory.contract.FilterLogs(opts, "EscrowPaused")
	if err != nil {
		return nil, err
	}
	return &ISlashEscrowFactoryEscrowPausedIterator{contract: _ISlashEscrowFactory.contract, event: "EscrowPaused", logs: logs, sub: sub}, nil
}

// WatchEscrowPaused is a free log subscription operation binding the contract event 0x050add19b1a78a4240cdebc8747899275f2dd070c88e83904a37ff7d1a539744.
//
// Solidity: event EscrowPaused((address,uint32) operatorSet, uint256 slashId)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) WatchEscrowPaused(opts *bind.WatchOpts, sink chan<- *ISlashEscrowFactoryEscrowPaused) (event.Subscription, error) {

	logs, sub, err := _ISlashEscrowFactory.contract.WatchLogs(opts, "EscrowPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlashEscrowFactoryEscrowPaused)
				if err := _ISlashEscrowFactory.contract.UnpackLog(event, "EscrowPaused", log); err != nil {
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

// ParseEscrowPaused is a log parse operation binding the contract event 0x050add19b1a78a4240cdebc8747899275f2dd070c88e83904a37ff7d1a539744.
//
// Solidity: event EscrowPaused((address,uint32) operatorSet, uint256 slashId)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) ParseEscrowPaused(log types.Log) (*ISlashEscrowFactoryEscrowPaused, error) {
	event := new(ISlashEscrowFactoryEscrowPaused)
	if err := _ISlashEscrowFactory.contract.UnpackLog(event, "EscrowPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlashEscrowFactoryEscrowUnpausedIterator is returned from FilterEscrowUnpaused and is used to iterate over the raw logs and unpacked data for EscrowUnpaused events raised by the ISlashEscrowFactory contract.
type ISlashEscrowFactoryEscrowUnpausedIterator struct {
	Event *ISlashEscrowFactoryEscrowUnpaused // Event containing the contract specifics and raw log

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
func (it *ISlashEscrowFactoryEscrowUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlashEscrowFactoryEscrowUnpaused)
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
		it.Event = new(ISlashEscrowFactoryEscrowUnpaused)
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
func (it *ISlashEscrowFactoryEscrowUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlashEscrowFactoryEscrowUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlashEscrowFactoryEscrowUnpaused represents a EscrowUnpaused event raised by the ISlashEscrowFactory contract.
type ISlashEscrowFactoryEscrowUnpaused struct {
	OperatorSet OperatorSet
	SlashId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEscrowUnpaused is a free log retrieval operation binding the contract event 0xb8877c6bf02d5f6603188eb653c9269271836b75b2012a5d7f5f233e7cf2f241.
//
// Solidity: event EscrowUnpaused((address,uint32) operatorSet, uint256 slashId)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) FilterEscrowUnpaused(opts *bind.FilterOpts) (*ISlashEscrowFactoryEscrowUnpausedIterator, error) {

	logs, sub, err := _ISlashEscrowFactory.contract.FilterLogs(opts, "EscrowUnpaused")
	if err != nil {
		return nil, err
	}
	return &ISlashEscrowFactoryEscrowUnpausedIterator{contract: _ISlashEscrowFactory.contract, event: "EscrowUnpaused", logs: logs, sub: sub}, nil
}

// WatchEscrowUnpaused is a free log subscription operation binding the contract event 0xb8877c6bf02d5f6603188eb653c9269271836b75b2012a5d7f5f233e7cf2f241.
//
// Solidity: event EscrowUnpaused((address,uint32) operatorSet, uint256 slashId)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) WatchEscrowUnpaused(opts *bind.WatchOpts, sink chan<- *ISlashEscrowFactoryEscrowUnpaused) (event.Subscription, error) {

	logs, sub, err := _ISlashEscrowFactory.contract.WatchLogs(opts, "EscrowUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlashEscrowFactoryEscrowUnpaused)
				if err := _ISlashEscrowFactory.contract.UnpackLog(event, "EscrowUnpaused", log); err != nil {
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

// ParseEscrowUnpaused is a log parse operation binding the contract event 0xb8877c6bf02d5f6603188eb653c9269271836b75b2012a5d7f5f233e7cf2f241.
//
// Solidity: event EscrowUnpaused((address,uint32) operatorSet, uint256 slashId)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) ParseEscrowUnpaused(log types.Log) (*ISlashEscrowFactoryEscrowUnpaused, error) {
	event := new(ISlashEscrowFactoryEscrowUnpaused)
	if err := _ISlashEscrowFactory.contract.UnpackLog(event, "EscrowUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlashEscrowFactoryGlobalEscrowDelaySetIterator is returned from FilterGlobalEscrowDelaySet and is used to iterate over the raw logs and unpacked data for GlobalEscrowDelaySet events raised by the ISlashEscrowFactory contract.
type ISlashEscrowFactoryGlobalEscrowDelaySetIterator struct {
	Event *ISlashEscrowFactoryGlobalEscrowDelaySet // Event containing the contract specifics and raw log

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
func (it *ISlashEscrowFactoryGlobalEscrowDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlashEscrowFactoryGlobalEscrowDelaySet)
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
		it.Event = new(ISlashEscrowFactoryGlobalEscrowDelaySet)
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
func (it *ISlashEscrowFactoryGlobalEscrowDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlashEscrowFactoryGlobalEscrowDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlashEscrowFactoryGlobalEscrowDelaySet represents a GlobalEscrowDelaySet event raised by the ISlashEscrowFactory contract.
type ISlashEscrowFactoryGlobalEscrowDelaySet struct {
	Delay uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGlobalEscrowDelaySet is a free log retrieval operation binding the contract event 0x67d0077d22e4e06f761dd87f6c9f2310ac879c9ce17de50d381e05b72f45fbf6.
//
// Solidity: event GlobalEscrowDelaySet(uint32 delay)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) FilterGlobalEscrowDelaySet(opts *bind.FilterOpts) (*ISlashEscrowFactoryGlobalEscrowDelaySetIterator, error) {

	logs, sub, err := _ISlashEscrowFactory.contract.FilterLogs(opts, "GlobalEscrowDelaySet")
	if err != nil {
		return nil, err
	}
	return &ISlashEscrowFactoryGlobalEscrowDelaySetIterator{contract: _ISlashEscrowFactory.contract, event: "GlobalEscrowDelaySet", logs: logs, sub: sub}, nil
}

// WatchGlobalEscrowDelaySet is a free log subscription operation binding the contract event 0x67d0077d22e4e06f761dd87f6c9f2310ac879c9ce17de50d381e05b72f45fbf6.
//
// Solidity: event GlobalEscrowDelaySet(uint32 delay)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) WatchGlobalEscrowDelaySet(opts *bind.WatchOpts, sink chan<- *ISlashEscrowFactoryGlobalEscrowDelaySet) (event.Subscription, error) {

	logs, sub, err := _ISlashEscrowFactory.contract.WatchLogs(opts, "GlobalEscrowDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlashEscrowFactoryGlobalEscrowDelaySet)
				if err := _ISlashEscrowFactory.contract.UnpackLog(event, "GlobalEscrowDelaySet", log); err != nil {
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

// ParseGlobalEscrowDelaySet is a log parse operation binding the contract event 0x67d0077d22e4e06f761dd87f6c9f2310ac879c9ce17de50d381e05b72f45fbf6.
//
// Solidity: event GlobalEscrowDelaySet(uint32 delay)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) ParseGlobalEscrowDelaySet(log types.Log) (*ISlashEscrowFactoryGlobalEscrowDelaySet, error) {
	event := new(ISlashEscrowFactoryGlobalEscrowDelaySet)
	if err := _ISlashEscrowFactory.contract.UnpackLog(event, "GlobalEscrowDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlashEscrowFactoryStartEscrowIterator is returned from FilterStartEscrow and is used to iterate over the raw logs and unpacked data for StartEscrow events raised by the ISlashEscrowFactory contract.
type ISlashEscrowFactoryStartEscrowIterator struct {
	Event *ISlashEscrowFactoryStartEscrow // Event containing the contract specifics and raw log

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
func (it *ISlashEscrowFactoryStartEscrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlashEscrowFactoryStartEscrow)
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
		it.Event = new(ISlashEscrowFactoryStartEscrow)
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
func (it *ISlashEscrowFactoryStartEscrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlashEscrowFactoryStartEscrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlashEscrowFactoryStartEscrow represents a StartEscrow event raised by the ISlashEscrowFactory contract.
type ISlashEscrowFactoryStartEscrow struct {
	OperatorSet OperatorSet
	SlashId     *big.Int
	Strategy    common.Address
	StartBlock  uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStartEscrow is a free log retrieval operation binding the contract event 0x3afae24c1d3dd2ce3649054ad82458a8c9967ebd9ce10a9a5a3d059f55bfaedb.
//
// Solidity: event StartEscrow((address,uint32) operatorSet, uint256 slashId, address strategy, uint32 startBlock)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) FilterStartEscrow(opts *bind.FilterOpts) (*ISlashEscrowFactoryStartEscrowIterator, error) {

	logs, sub, err := _ISlashEscrowFactory.contract.FilterLogs(opts, "StartEscrow")
	if err != nil {
		return nil, err
	}
	return &ISlashEscrowFactoryStartEscrowIterator{contract: _ISlashEscrowFactory.contract, event: "StartEscrow", logs: logs, sub: sub}, nil
}

// WatchStartEscrow is a free log subscription operation binding the contract event 0x3afae24c1d3dd2ce3649054ad82458a8c9967ebd9ce10a9a5a3d059f55bfaedb.
//
// Solidity: event StartEscrow((address,uint32) operatorSet, uint256 slashId, address strategy, uint32 startBlock)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) WatchStartEscrow(opts *bind.WatchOpts, sink chan<- *ISlashEscrowFactoryStartEscrow) (event.Subscription, error) {

	logs, sub, err := _ISlashEscrowFactory.contract.WatchLogs(opts, "StartEscrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlashEscrowFactoryStartEscrow)
				if err := _ISlashEscrowFactory.contract.UnpackLog(event, "StartEscrow", log); err != nil {
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

// ParseStartEscrow is a log parse operation binding the contract event 0x3afae24c1d3dd2ce3649054ad82458a8c9967ebd9ce10a9a5a3d059f55bfaedb.
//
// Solidity: event StartEscrow((address,uint32) operatorSet, uint256 slashId, address strategy, uint32 startBlock)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) ParseStartEscrow(log types.Log) (*ISlashEscrowFactoryStartEscrow, error) {
	event := new(ISlashEscrowFactoryStartEscrow)
	if err := _ISlashEscrowFactory.contract.UnpackLog(event, "StartEscrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISlashEscrowFactoryStrategyEscrowDelaySetIterator is returned from FilterStrategyEscrowDelaySet and is used to iterate over the raw logs and unpacked data for StrategyEscrowDelaySet events raised by the ISlashEscrowFactory contract.
type ISlashEscrowFactoryStrategyEscrowDelaySetIterator struct {
	Event *ISlashEscrowFactoryStrategyEscrowDelaySet // Event containing the contract specifics and raw log

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
func (it *ISlashEscrowFactoryStrategyEscrowDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISlashEscrowFactoryStrategyEscrowDelaySet)
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
		it.Event = new(ISlashEscrowFactoryStrategyEscrowDelaySet)
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
func (it *ISlashEscrowFactoryStrategyEscrowDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISlashEscrowFactoryStrategyEscrowDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISlashEscrowFactoryStrategyEscrowDelaySet represents a StrategyEscrowDelaySet event raised by the ISlashEscrowFactory contract.
type ISlashEscrowFactoryStrategyEscrowDelaySet struct {
	Strategy common.Address
	Delay    uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyEscrowDelaySet is a free log retrieval operation binding the contract event 0x5d2b33f07ae22a809e79005f96ffac70c3715df85a3b011b025e0a86a23a007b.
//
// Solidity: event StrategyEscrowDelaySet(address strategy, uint32 delay)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) FilterStrategyEscrowDelaySet(opts *bind.FilterOpts) (*ISlashEscrowFactoryStrategyEscrowDelaySetIterator, error) {

	logs, sub, err := _ISlashEscrowFactory.contract.FilterLogs(opts, "StrategyEscrowDelaySet")
	if err != nil {
		return nil, err
	}
	return &ISlashEscrowFactoryStrategyEscrowDelaySetIterator{contract: _ISlashEscrowFactory.contract, event: "StrategyEscrowDelaySet", logs: logs, sub: sub}, nil
}

// WatchStrategyEscrowDelaySet is a free log subscription operation binding the contract event 0x5d2b33f07ae22a809e79005f96ffac70c3715df85a3b011b025e0a86a23a007b.
//
// Solidity: event StrategyEscrowDelaySet(address strategy, uint32 delay)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) WatchStrategyEscrowDelaySet(opts *bind.WatchOpts, sink chan<- *ISlashEscrowFactoryStrategyEscrowDelaySet) (event.Subscription, error) {

	logs, sub, err := _ISlashEscrowFactory.contract.WatchLogs(opts, "StrategyEscrowDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISlashEscrowFactoryStrategyEscrowDelaySet)
				if err := _ISlashEscrowFactory.contract.UnpackLog(event, "StrategyEscrowDelaySet", log); err != nil {
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

// ParseStrategyEscrowDelaySet is a log parse operation binding the contract event 0x5d2b33f07ae22a809e79005f96ffac70c3715df85a3b011b025e0a86a23a007b.
//
// Solidity: event StrategyEscrowDelaySet(address strategy, uint32 delay)
func (_ISlashEscrowFactory *ISlashEscrowFactoryFilterer) ParseStrategyEscrowDelaySet(log types.Log) (*ISlashEscrowFactoryStrategyEscrowDelaySet, error) {
	event := new(ISlashEscrowFactoryStrategyEscrowDelaySet)
	if err := _ISlashEscrowFactory.contract.UnpackLog(event, "StrategyEscrowDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
