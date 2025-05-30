// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SlashEscrowFactoryStorage

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

// SlashEscrowFactoryStorageMetaData contains all meta data concerning the SlashEscrowFactoryStorage contract.
var SlashEscrowFactoryStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeSlashEscrowSalt\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getEscrowCompleteBlock\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEscrowStartBlock\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalEscrowDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingEscrows\",\"inputs\":[],\"outputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"isRedistributing\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"slashIds\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"},{\"name\":\"completeBlocks\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingOperatorSets\",\"inputs\":[],\"outputs\":[{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingSlashIds\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingStrategiesForSlashId\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingStrategiesForSlashIds\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address[][]\",\"internalType\":\"contractIStrategy[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingUnderlyingAmountForStrategy\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlashEscrow\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStrategyEscrowDelay\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalPendingOperatorSets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalPendingSlashIds\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalPendingStrategiesForSlashId\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialGlobalDelayBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateSlashEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDeployedSlashEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isDeployedSlashEscrow\",\"inputs\":[{\"name\":\"slashEscrow\",\"type\":\"address\",\"internalType\":\"contractISlashEscrow\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isEscrowPaused\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPendingOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPendingSlashId\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseSlashEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseSlashEscrowByStrategy\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalEscrowDelay\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyEscrowDelay\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashEscrowImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlashEscrow\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpauseEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EscrowComplete\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EscrowPaused\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EscrowUnpaused\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalEscrowDelaySet\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StartEscrow\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyEscrowDelaySet\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EscrowDelayNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EscrowNotMature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRedistributionRecipient\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManager\",\"inputs\":[]}]",
}

// SlashEscrowFactoryStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use SlashEscrowFactoryStorageMetaData.ABI instead.
var SlashEscrowFactoryStorageABI = SlashEscrowFactoryStorageMetaData.ABI

// SlashEscrowFactoryStorage is an auto generated Go binding around an Ethereum contract.
type SlashEscrowFactoryStorage struct {
	SlashEscrowFactoryStorageCaller     // Read-only binding to the contract
	SlashEscrowFactoryStorageTransactor // Write-only binding to the contract
	SlashEscrowFactoryStorageFilterer   // Log filterer for contract events
}

// SlashEscrowFactoryStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlashEscrowFactoryStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashEscrowFactoryStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlashEscrowFactoryStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashEscrowFactoryStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlashEscrowFactoryStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashEscrowFactoryStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlashEscrowFactoryStorageSession struct {
	Contract     *SlashEscrowFactoryStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SlashEscrowFactoryStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlashEscrowFactoryStorageCallerSession struct {
	Contract *SlashEscrowFactoryStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// SlashEscrowFactoryStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlashEscrowFactoryStorageTransactorSession struct {
	Contract     *SlashEscrowFactoryStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// SlashEscrowFactoryStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlashEscrowFactoryStorageRaw struct {
	Contract *SlashEscrowFactoryStorage // Generic contract binding to access the raw methods on
}

// SlashEscrowFactoryStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlashEscrowFactoryStorageCallerRaw struct {
	Contract *SlashEscrowFactoryStorageCaller // Generic read-only contract binding to access the raw methods on
}

// SlashEscrowFactoryStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlashEscrowFactoryStorageTransactorRaw struct {
	Contract *SlashEscrowFactoryStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlashEscrowFactoryStorage creates a new instance of SlashEscrowFactoryStorage, bound to a specific deployed contract.
func NewSlashEscrowFactoryStorage(address common.Address, backend bind.ContractBackend) (*SlashEscrowFactoryStorage, error) {
	contract, err := bindSlashEscrowFactoryStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryStorage{SlashEscrowFactoryStorageCaller: SlashEscrowFactoryStorageCaller{contract: contract}, SlashEscrowFactoryStorageTransactor: SlashEscrowFactoryStorageTransactor{contract: contract}, SlashEscrowFactoryStorageFilterer: SlashEscrowFactoryStorageFilterer{contract: contract}}, nil
}

// NewSlashEscrowFactoryStorageCaller creates a new read-only instance of SlashEscrowFactoryStorage, bound to a specific deployed contract.
func NewSlashEscrowFactoryStorageCaller(address common.Address, caller bind.ContractCaller) (*SlashEscrowFactoryStorageCaller, error) {
	contract, err := bindSlashEscrowFactoryStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryStorageCaller{contract: contract}, nil
}

// NewSlashEscrowFactoryStorageTransactor creates a new write-only instance of SlashEscrowFactoryStorage, bound to a specific deployed contract.
func NewSlashEscrowFactoryStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SlashEscrowFactoryStorageTransactor, error) {
	contract, err := bindSlashEscrowFactoryStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryStorageTransactor{contract: contract}, nil
}

// NewSlashEscrowFactoryStorageFilterer creates a new log filterer instance of SlashEscrowFactoryStorage, bound to a specific deployed contract.
func NewSlashEscrowFactoryStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SlashEscrowFactoryStorageFilterer, error) {
	contract, err := bindSlashEscrowFactoryStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryStorageFilterer{contract: contract}, nil
}

// bindSlashEscrowFactoryStorage binds a generic wrapper to an already deployed contract.
func bindSlashEscrowFactoryStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlashEscrowFactoryStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashEscrowFactoryStorage.Contract.SlashEscrowFactoryStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.SlashEscrowFactoryStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.SlashEscrowFactoryStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashEscrowFactoryStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.contract.Transact(opts, method, params...)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) AllocationManager() (common.Address, error) {
	return _SlashEscrowFactoryStorage.Contract.AllocationManager(&_SlashEscrowFactoryStorage.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) AllocationManager() (common.Address, error) {
	return _SlashEscrowFactoryStorage.Contract.AllocationManager(&_SlashEscrowFactoryStorage.CallOpts)
}

// ComputeSlashEscrowSalt is a free data retrieval call binding the contract method 0xc2de7096.
//
// Solidity: function computeSlashEscrowSalt((address,uint32) operatorSet, uint256 slashId) pure returns(bytes32)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) ComputeSlashEscrowSalt(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "computeSlashEscrowSalt", operatorSet, slashId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeSlashEscrowSalt is a free data retrieval call binding the contract method 0xc2de7096.
//
// Solidity: function computeSlashEscrowSalt((address,uint32) operatorSet, uint256 slashId) pure returns(bytes32)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) ComputeSlashEscrowSalt(operatorSet OperatorSet, slashId *big.Int) ([32]byte, error) {
	return _SlashEscrowFactoryStorage.Contract.ComputeSlashEscrowSalt(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// ComputeSlashEscrowSalt is a free data retrieval call binding the contract method 0xc2de7096.
//
// Solidity: function computeSlashEscrowSalt((address,uint32) operatorSet, uint256 slashId) pure returns(bytes32)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) ComputeSlashEscrowSalt(operatorSet OperatorSet, slashId *big.Int) ([32]byte, error) {
	return _SlashEscrowFactoryStorage.Contract.ComputeSlashEscrowSalt(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// GetEscrowCompleteBlock is a free data retrieval call binding the contract method 0x0310f3e6.
//
// Solidity: function getEscrowCompleteBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint32)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) GetEscrowCompleteBlock(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (uint32, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "getEscrowCompleteBlock", operatorSet, slashId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetEscrowCompleteBlock is a free data retrieval call binding the contract method 0x0310f3e6.
//
// Solidity: function getEscrowCompleteBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint32)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) GetEscrowCompleteBlock(operatorSet OperatorSet, slashId *big.Int) (uint32, error) {
	return _SlashEscrowFactoryStorage.Contract.GetEscrowCompleteBlock(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// GetEscrowCompleteBlock is a free data retrieval call binding the contract method 0x0310f3e6.
//
// Solidity: function getEscrowCompleteBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint32)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) GetEscrowCompleteBlock(operatorSet OperatorSet, slashId *big.Int) (uint32, error) {
	return _SlashEscrowFactoryStorage.Contract.GetEscrowCompleteBlock(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// GetEscrowStartBlock is a free data retrieval call binding the contract method 0x05a4dfbb.
//
// Solidity: function getEscrowStartBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) GetEscrowStartBlock(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "getEscrowStartBlock", operatorSet, slashId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEscrowStartBlock is a free data retrieval call binding the contract method 0x05a4dfbb.
//
// Solidity: function getEscrowStartBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) GetEscrowStartBlock(operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	return _SlashEscrowFactoryStorage.Contract.GetEscrowStartBlock(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// GetEscrowStartBlock is a free data retrieval call binding the contract method 0x05a4dfbb.
//
// Solidity: function getEscrowStartBlock((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) GetEscrowStartBlock(operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	return _SlashEscrowFactoryStorage.Contract.GetEscrowStartBlock(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// GetGlobalEscrowDelay is a free data retrieval call binding the contract method 0x8a65d2d2.
//
// Solidity: function getGlobalEscrowDelay() view returns(uint32)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) GetGlobalEscrowDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "getGlobalEscrowDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetGlobalEscrowDelay is a free data retrieval call binding the contract method 0x8a65d2d2.
//
// Solidity: function getGlobalEscrowDelay() view returns(uint32)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) GetGlobalEscrowDelay() (uint32, error) {
	return _SlashEscrowFactoryStorage.Contract.GetGlobalEscrowDelay(&_SlashEscrowFactoryStorage.CallOpts)
}

// GetGlobalEscrowDelay is a free data retrieval call binding the contract method 0x8a65d2d2.
//
// Solidity: function getGlobalEscrowDelay() view returns(uint32)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) GetGlobalEscrowDelay() (uint32, error) {
	return _SlashEscrowFactoryStorage.Contract.GetGlobalEscrowDelay(&_SlashEscrowFactoryStorage.CallOpts)
}

// GetPendingEscrows is a free data retrieval call binding the contract method 0xc50f4e48.
//
// Solidity: function getPendingEscrows() view returns((address,uint32)[] operatorSets, bool[] isRedistributing, uint256[][] slashIds, uint32[][] completeBlocks)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) GetPendingEscrows(opts *bind.CallOpts) (struct {
	OperatorSets     []OperatorSet
	IsRedistributing []bool
	SlashIds         [][]*big.Int
	CompleteBlocks   [][]uint32
}, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "getPendingEscrows")

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
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) GetPendingEscrows() (struct {
	OperatorSets     []OperatorSet
	IsRedistributing []bool
	SlashIds         [][]*big.Int
	CompleteBlocks   [][]uint32
}, error) {
	return _SlashEscrowFactoryStorage.Contract.GetPendingEscrows(&_SlashEscrowFactoryStorage.CallOpts)
}

// GetPendingEscrows is a free data retrieval call binding the contract method 0xc50f4e48.
//
// Solidity: function getPendingEscrows() view returns((address,uint32)[] operatorSets, bool[] isRedistributing, uint256[][] slashIds, uint32[][] completeBlocks)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) GetPendingEscrows() (struct {
	OperatorSets     []OperatorSet
	IsRedistributing []bool
	SlashIds         [][]*big.Int
	CompleteBlocks   [][]uint32
}, error) {
	return _SlashEscrowFactoryStorage.Contract.GetPendingEscrows(&_SlashEscrowFactoryStorage.CallOpts)
}

// GetPendingOperatorSets is a free data retrieval call binding the contract method 0x3f292b08.
//
// Solidity: function getPendingOperatorSets() view returns((address,uint32)[] operatorSets)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) GetPendingOperatorSets(opts *bind.CallOpts) ([]OperatorSet, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "getPendingOperatorSets")

	if err != nil {
		return *new([]OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)

	return out0, err

}

// GetPendingOperatorSets is a free data retrieval call binding the contract method 0x3f292b08.
//
// Solidity: function getPendingOperatorSets() view returns((address,uint32)[] operatorSets)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) GetPendingOperatorSets() ([]OperatorSet, error) {
	return _SlashEscrowFactoryStorage.Contract.GetPendingOperatorSets(&_SlashEscrowFactoryStorage.CallOpts)
}

// GetPendingOperatorSets is a free data retrieval call binding the contract method 0x3f292b08.
//
// Solidity: function getPendingOperatorSets() view returns((address,uint32)[] operatorSets)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) GetPendingOperatorSets() ([]OperatorSet, error) {
	return _SlashEscrowFactoryStorage.Contract.GetPendingOperatorSets(&_SlashEscrowFactoryStorage.CallOpts)
}

// GetPendingSlashIds is a free data retrieval call binding the contract method 0x7def1564.
//
// Solidity: function getPendingSlashIds((address,uint32) operatorSet) view returns(uint256[])
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) GetPendingSlashIds(opts *bind.CallOpts, operatorSet OperatorSet) ([]*big.Int, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "getPendingSlashIds", operatorSet)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPendingSlashIds is a free data retrieval call binding the contract method 0x7def1564.
//
// Solidity: function getPendingSlashIds((address,uint32) operatorSet) view returns(uint256[])
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) GetPendingSlashIds(operatorSet OperatorSet) ([]*big.Int, error) {
	return _SlashEscrowFactoryStorage.Contract.GetPendingSlashIds(&_SlashEscrowFactoryStorage.CallOpts, operatorSet)
}

// GetPendingSlashIds is a free data retrieval call binding the contract method 0x7def1564.
//
// Solidity: function getPendingSlashIds((address,uint32) operatorSet) view returns(uint256[])
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) GetPendingSlashIds(operatorSet OperatorSet) ([]*big.Int, error) {
	return _SlashEscrowFactoryStorage.Contract.GetPendingSlashIds(&_SlashEscrowFactoryStorage.CallOpts, operatorSet)
}

// GetPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x7130c423.
//
// Solidity: function getPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(address[] strategies)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) GetPendingStrategiesForSlashId(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "getPendingStrategiesForSlashId", operatorSet, slashId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x7130c423.
//
// Solidity: function getPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(address[] strategies)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) GetPendingStrategiesForSlashId(operatorSet OperatorSet, slashId *big.Int) ([]common.Address, error) {
	return _SlashEscrowFactoryStorage.Contract.GetPendingStrategiesForSlashId(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// GetPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x7130c423.
//
// Solidity: function getPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(address[] strategies)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) GetPendingStrategiesForSlashId(operatorSet OperatorSet, slashId *big.Int) ([]common.Address, error) {
	return _SlashEscrowFactoryStorage.Contract.GetPendingStrategiesForSlashId(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// GetPendingStrategiesForSlashIds is a free data retrieval call binding the contract method 0xa56b21e4.
//
// Solidity: function getPendingStrategiesForSlashIds((address,uint32) operatorSet) view returns(address[][] strategies)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) GetPendingStrategiesForSlashIds(opts *bind.CallOpts, operatorSet OperatorSet) ([][]common.Address, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "getPendingStrategiesForSlashIds", operatorSet)

	if err != nil {
		return *new([][]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([][]common.Address)).(*[][]common.Address)

	return out0, err

}

// GetPendingStrategiesForSlashIds is a free data retrieval call binding the contract method 0xa56b21e4.
//
// Solidity: function getPendingStrategiesForSlashIds((address,uint32) operatorSet) view returns(address[][] strategies)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) GetPendingStrategiesForSlashIds(operatorSet OperatorSet) ([][]common.Address, error) {
	return _SlashEscrowFactoryStorage.Contract.GetPendingStrategiesForSlashIds(&_SlashEscrowFactoryStorage.CallOpts, operatorSet)
}

// GetPendingStrategiesForSlashIds is a free data retrieval call binding the contract method 0xa56b21e4.
//
// Solidity: function getPendingStrategiesForSlashIds((address,uint32) operatorSet) view returns(address[][] strategies)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) GetPendingStrategiesForSlashIds(operatorSet OperatorSet) ([][]common.Address, error) {
	return _SlashEscrowFactoryStorage.Contract.GetPendingStrategiesForSlashIds(&_SlashEscrowFactoryStorage.CallOpts, operatorSet)
}

// GetPendingUnderlyingAmountForStrategy is a free data retrieval call binding the contract method 0xb23ff83b.
//
// Solidity: function getPendingUnderlyingAmountForStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) view returns(uint256)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) GetPendingUnderlyingAmountForStrategy(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "getPendingUnderlyingAmountForStrategy", operatorSet, slashId, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingUnderlyingAmountForStrategy is a free data retrieval call binding the contract method 0xb23ff83b.
//
// Solidity: function getPendingUnderlyingAmountForStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) view returns(uint256)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) GetPendingUnderlyingAmountForStrategy(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*big.Int, error) {
	return _SlashEscrowFactoryStorage.Contract.GetPendingUnderlyingAmountForStrategy(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId, strategy)
}

// GetPendingUnderlyingAmountForStrategy is a free data retrieval call binding the contract method 0xb23ff83b.
//
// Solidity: function getPendingUnderlyingAmountForStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) view returns(uint256)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) GetPendingUnderlyingAmountForStrategy(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*big.Int, error) {
	return _SlashEscrowFactoryStorage.Contract.GetPendingUnderlyingAmountForStrategy(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId, strategy)
}

// GetSlashEscrow is a free data retrieval call binding the contract method 0x3453b234.
//
// Solidity: function getSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(address)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) GetSlashEscrow(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "getSlashEscrow", operatorSet, slashId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSlashEscrow is a free data retrieval call binding the contract method 0x3453b234.
//
// Solidity: function getSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(address)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) GetSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (common.Address, error) {
	return _SlashEscrowFactoryStorage.Contract.GetSlashEscrow(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// GetSlashEscrow is a free data retrieval call binding the contract method 0x3453b234.
//
// Solidity: function getSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(address)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) GetSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (common.Address, error) {
	return _SlashEscrowFactoryStorage.Contract.GetSlashEscrow(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// GetStrategyEscrowDelay is a free data retrieval call binding the contract method 0x8d5d4036.
//
// Solidity: function getStrategyEscrowDelay(address strategy) view returns(uint32)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) GetStrategyEscrowDelay(opts *bind.CallOpts, strategy common.Address) (uint32, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "getStrategyEscrowDelay", strategy)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetStrategyEscrowDelay is a free data retrieval call binding the contract method 0x8d5d4036.
//
// Solidity: function getStrategyEscrowDelay(address strategy) view returns(uint32)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) GetStrategyEscrowDelay(strategy common.Address) (uint32, error) {
	return _SlashEscrowFactoryStorage.Contract.GetStrategyEscrowDelay(&_SlashEscrowFactoryStorage.CallOpts, strategy)
}

// GetStrategyEscrowDelay is a free data retrieval call binding the contract method 0x8d5d4036.
//
// Solidity: function getStrategyEscrowDelay(address strategy) view returns(uint32)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) GetStrategyEscrowDelay(strategy common.Address) (uint32, error) {
	return _SlashEscrowFactoryStorage.Contract.GetStrategyEscrowDelay(&_SlashEscrowFactoryStorage.CallOpts, strategy)
}

// GetTotalPendingOperatorSets is a free data retrieval call binding the contract method 0x78cb9600.
//
// Solidity: function getTotalPendingOperatorSets() view returns(uint256)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) GetTotalPendingOperatorSets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "getTotalPendingOperatorSets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPendingOperatorSets is a free data retrieval call binding the contract method 0x78cb9600.
//
// Solidity: function getTotalPendingOperatorSets() view returns(uint256)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) GetTotalPendingOperatorSets() (*big.Int, error) {
	return _SlashEscrowFactoryStorage.Contract.GetTotalPendingOperatorSets(&_SlashEscrowFactoryStorage.CallOpts)
}

// GetTotalPendingOperatorSets is a free data retrieval call binding the contract method 0x78cb9600.
//
// Solidity: function getTotalPendingOperatorSets() view returns(uint256)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) GetTotalPendingOperatorSets() (*big.Int, error) {
	return _SlashEscrowFactoryStorage.Contract.GetTotalPendingOperatorSets(&_SlashEscrowFactoryStorage.CallOpts)
}

// GetTotalPendingSlashIds is a free data retrieval call binding the contract method 0x6729b5db.
//
// Solidity: function getTotalPendingSlashIds((address,uint32) operatorSet) view returns(uint256)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) GetTotalPendingSlashIds(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "getTotalPendingSlashIds", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPendingSlashIds is a free data retrieval call binding the contract method 0x6729b5db.
//
// Solidity: function getTotalPendingSlashIds((address,uint32) operatorSet) view returns(uint256)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) GetTotalPendingSlashIds(operatorSet OperatorSet) (*big.Int, error) {
	return _SlashEscrowFactoryStorage.Contract.GetTotalPendingSlashIds(&_SlashEscrowFactoryStorage.CallOpts, operatorSet)
}

// GetTotalPendingSlashIds is a free data retrieval call binding the contract method 0x6729b5db.
//
// Solidity: function getTotalPendingSlashIds((address,uint32) operatorSet) view returns(uint256)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) GetTotalPendingSlashIds(operatorSet OperatorSet) (*big.Int, error) {
	return _SlashEscrowFactoryStorage.Contract.GetTotalPendingSlashIds(&_SlashEscrowFactoryStorage.CallOpts, operatorSet)
}

// GetTotalPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x277a9f0e.
//
// Solidity: function getTotalPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) GetTotalPendingStrategiesForSlashId(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "getTotalPendingStrategiesForSlashId", operatorSet, slashId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x277a9f0e.
//
// Solidity: function getTotalPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) GetTotalPendingStrategiesForSlashId(operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	return _SlashEscrowFactoryStorage.Contract.GetTotalPendingStrategiesForSlashId(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// GetTotalPendingStrategiesForSlashId is a free data retrieval call binding the contract method 0x277a9f0e.
//
// Solidity: function getTotalPendingStrategiesForSlashId((address,uint32) operatorSet, uint256 slashId) view returns(uint256)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) GetTotalPendingStrategiesForSlashId(operatorSet OperatorSet, slashId *big.Int) (*big.Int, error) {
	return _SlashEscrowFactoryStorage.Contract.GetTotalPendingStrategiesForSlashId(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// IsDeployedSlashEscrow is a free data retrieval call binding the contract method 0x19f3db26.
//
// Solidity: function isDeployedSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) IsDeployedSlashEscrow(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "isDeployedSlashEscrow", operatorSet, slashId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeployedSlashEscrow is a free data retrieval call binding the contract method 0x19f3db26.
//
// Solidity: function isDeployedSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) IsDeployedSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _SlashEscrowFactoryStorage.Contract.IsDeployedSlashEscrow(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// IsDeployedSlashEscrow is a free data retrieval call binding the contract method 0x19f3db26.
//
// Solidity: function isDeployedSlashEscrow((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) IsDeployedSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _SlashEscrowFactoryStorage.Contract.IsDeployedSlashEscrow(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// IsDeployedSlashEscrow0 is a free data retrieval call binding the contract method 0xc8b5330c.
//
// Solidity: function isDeployedSlashEscrow(address slashEscrow) view returns(bool)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) IsDeployedSlashEscrow0(opts *bind.CallOpts, slashEscrow common.Address) (bool, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "isDeployedSlashEscrow0", slashEscrow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeployedSlashEscrow0 is a free data retrieval call binding the contract method 0xc8b5330c.
//
// Solidity: function isDeployedSlashEscrow(address slashEscrow) view returns(bool)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) IsDeployedSlashEscrow0(slashEscrow common.Address) (bool, error) {
	return _SlashEscrowFactoryStorage.Contract.IsDeployedSlashEscrow0(&_SlashEscrowFactoryStorage.CallOpts, slashEscrow)
}

// IsDeployedSlashEscrow0 is a free data retrieval call binding the contract method 0xc8b5330c.
//
// Solidity: function isDeployedSlashEscrow(address slashEscrow) view returns(bool)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) IsDeployedSlashEscrow0(slashEscrow common.Address) (bool, error) {
	return _SlashEscrowFactoryStorage.Contract.IsDeployedSlashEscrow0(&_SlashEscrowFactoryStorage.CallOpts, slashEscrow)
}

// IsEscrowPaused is a free data retrieval call binding the contract method 0x78748459.
//
// Solidity: function isEscrowPaused((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) IsEscrowPaused(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "isEscrowPaused", operatorSet, slashId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEscrowPaused is a free data retrieval call binding the contract method 0x78748459.
//
// Solidity: function isEscrowPaused((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) IsEscrowPaused(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _SlashEscrowFactoryStorage.Contract.IsEscrowPaused(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// IsEscrowPaused is a free data retrieval call binding the contract method 0x78748459.
//
// Solidity: function isEscrowPaused((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) IsEscrowPaused(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _SlashEscrowFactoryStorage.Contract.IsEscrowPaused(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// IsPendingOperatorSet is a free data retrieval call binding the contract method 0x71e166e7.
//
// Solidity: function isPendingOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) IsPendingOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "isPendingOperatorSet", operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingOperatorSet is a free data retrieval call binding the contract method 0x71e166e7.
//
// Solidity: function isPendingOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) IsPendingOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _SlashEscrowFactoryStorage.Contract.IsPendingOperatorSet(&_SlashEscrowFactoryStorage.CallOpts, operatorSet)
}

// IsPendingOperatorSet is a free data retrieval call binding the contract method 0x71e166e7.
//
// Solidity: function isPendingOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) IsPendingOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _SlashEscrowFactoryStorage.Contract.IsPendingOperatorSet(&_SlashEscrowFactoryStorage.CallOpts, operatorSet)
}

// IsPendingSlashId is a free data retrieval call binding the contract method 0x87420b07.
//
// Solidity: function isPendingSlashId((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) IsPendingSlashId(opts *bind.CallOpts, operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "isPendingSlashId", operatorSet, slashId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingSlashId is a free data retrieval call binding the contract method 0x87420b07.
//
// Solidity: function isPendingSlashId((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) IsPendingSlashId(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _SlashEscrowFactoryStorage.Contract.IsPendingSlashId(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// IsPendingSlashId is a free data retrieval call binding the contract method 0x87420b07.
//
// Solidity: function isPendingSlashId((address,uint32) operatorSet, uint256 slashId) view returns(bool)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) IsPendingSlashId(operatorSet OperatorSet, slashId *big.Int) (bool, error) {
	return _SlashEscrowFactoryStorage.Contract.IsPendingSlashId(&_SlashEscrowFactoryStorage.CallOpts, operatorSet, slashId)
}

// SlashEscrowImplementation is a free data retrieval call binding the contract method 0xa3c65641.
//
// Solidity: function slashEscrowImplementation() view returns(address)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) SlashEscrowImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "slashEscrowImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlashEscrowImplementation is a free data retrieval call binding the contract method 0xa3c65641.
//
// Solidity: function slashEscrowImplementation() view returns(address)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) SlashEscrowImplementation() (common.Address, error) {
	return _SlashEscrowFactoryStorage.Contract.SlashEscrowImplementation(&_SlashEscrowFactoryStorage.CallOpts)
}

// SlashEscrowImplementation is a free data retrieval call binding the contract method 0xa3c65641.
//
// Solidity: function slashEscrowImplementation() view returns(address)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) SlashEscrowImplementation() (common.Address, error) {
	return _SlashEscrowFactoryStorage.Contract.SlashEscrowImplementation(&_SlashEscrowFactoryStorage.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashEscrowFactoryStorage.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) StrategyManager() (common.Address, error) {
	return _SlashEscrowFactoryStorage.Contract.StrategyManager(&_SlashEscrowFactoryStorage.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageCallerSession) StrategyManager() (common.Address, error) {
	return _SlashEscrowFactoryStorage.Contract.StrategyManager(&_SlashEscrowFactoryStorage.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x6c5ac81e.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus, uint32 initialGlobalDelayBlocks) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int, initialGlobalDelayBlocks uint32) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.contract.Transact(opts, "initialize", initialOwner, initialPausedStatus, initialGlobalDelayBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0x6c5ac81e.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus, uint32 initialGlobalDelayBlocks) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int, initialGlobalDelayBlocks uint32) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.Initialize(&_SlashEscrowFactoryStorage.TransactOpts, initialOwner, initialPausedStatus, initialGlobalDelayBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0x6c5ac81e.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus, uint32 initialGlobalDelayBlocks) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactorSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int, initialGlobalDelayBlocks uint32) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.Initialize(&_SlashEscrowFactoryStorage.TransactOpts, initialOwner, initialPausedStatus, initialGlobalDelayBlocks)
}

// InitiateSlashEscrow is a paid mutator transaction binding the contract method 0x7a967611.
//
// Solidity: function initiateSlashEscrow((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactor) InitiateSlashEscrow(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.contract.Transact(opts, "initiateSlashEscrow", operatorSet, slashId, strategy)
}

// InitiateSlashEscrow is a paid mutator transaction binding the contract method 0x7a967611.
//
// Solidity: function initiateSlashEscrow((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) InitiateSlashEscrow(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.InitiateSlashEscrow(&_SlashEscrowFactoryStorage.TransactOpts, operatorSet, slashId, strategy)
}

// InitiateSlashEscrow is a paid mutator transaction binding the contract method 0x7a967611.
//
// Solidity: function initiateSlashEscrow((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactorSession) InitiateSlashEscrow(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.InitiateSlashEscrow(&_SlashEscrowFactoryStorage.TransactOpts, operatorSet, slashId, strategy)
}

// PauseEscrow is a paid mutator transaction binding the contract method 0xe7ed076b.
//
// Solidity: function pauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactor) PauseEscrow(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.contract.Transact(opts, "pauseEscrow", operatorSet, slashId)
}

// PauseEscrow is a paid mutator transaction binding the contract method 0xe7ed076b.
//
// Solidity: function pauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) PauseEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.PauseEscrow(&_SlashEscrowFactoryStorage.TransactOpts, operatorSet, slashId)
}

// PauseEscrow is a paid mutator transaction binding the contract method 0xe7ed076b.
//
// Solidity: function pauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactorSession) PauseEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.PauseEscrow(&_SlashEscrowFactoryStorage.TransactOpts, operatorSet, slashId)
}

// ReleaseSlashEscrow is a paid mutator transaction binding the contract method 0x5ffa5a81.
//
// Solidity: function releaseSlashEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactor) ReleaseSlashEscrow(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.contract.Transact(opts, "releaseSlashEscrow", operatorSet, slashId)
}

// ReleaseSlashEscrow is a paid mutator transaction binding the contract method 0x5ffa5a81.
//
// Solidity: function releaseSlashEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) ReleaseSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.ReleaseSlashEscrow(&_SlashEscrowFactoryStorage.TransactOpts, operatorSet, slashId)
}

// ReleaseSlashEscrow is a paid mutator transaction binding the contract method 0x5ffa5a81.
//
// Solidity: function releaseSlashEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactorSession) ReleaseSlashEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.ReleaseSlashEscrow(&_SlashEscrowFactoryStorage.TransactOpts, operatorSet, slashId)
}

// ReleaseSlashEscrowByStrategy is a paid mutator transaction binding the contract method 0x0e475b17.
//
// Solidity: function releaseSlashEscrowByStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactor) ReleaseSlashEscrowByStrategy(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.contract.Transact(opts, "releaseSlashEscrowByStrategy", operatorSet, slashId, strategy)
}

// ReleaseSlashEscrowByStrategy is a paid mutator transaction binding the contract method 0x0e475b17.
//
// Solidity: function releaseSlashEscrowByStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) ReleaseSlashEscrowByStrategy(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.ReleaseSlashEscrowByStrategy(&_SlashEscrowFactoryStorage.TransactOpts, operatorSet, slashId, strategy)
}

// ReleaseSlashEscrowByStrategy is a paid mutator transaction binding the contract method 0x0e475b17.
//
// Solidity: function releaseSlashEscrowByStrategy((address,uint32) operatorSet, uint256 slashId, address strategy) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactorSession) ReleaseSlashEscrowByStrategy(operatorSet OperatorSet, slashId *big.Int, strategy common.Address) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.ReleaseSlashEscrowByStrategy(&_SlashEscrowFactoryStorage.TransactOpts, operatorSet, slashId, strategy)
}

// SetGlobalEscrowDelay is a paid mutator transaction binding the contract method 0x5e0a64c5.
//
// Solidity: function setGlobalEscrowDelay(uint32 delay) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactor) SetGlobalEscrowDelay(opts *bind.TransactOpts, delay uint32) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.contract.Transact(opts, "setGlobalEscrowDelay", delay)
}

// SetGlobalEscrowDelay is a paid mutator transaction binding the contract method 0x5e0a64c5.
//
// Solidity: function setGlobalEscrowDelay(uint32 delay) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) SetGlobalEscrowDelay(delay uint32) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.SetGlobalEscrowDelay(&_SlashEscrowFactoryStorage.TransactOpts, delay)
}

// SetGlobalEscrowDelay is a paid mutator transaction binding the contract method 0x5e0a64c5.
//
// Solidity: function setGlobalEscrowDelay(uint32 delay) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactorSession) SetGlobalEscrowDelay(delay uint32) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.SetGlobalEscrowDelay(&_SlashEscrowFactoryStorage.TransactOpts, delay)
}

// SetStrategyEscrowDelay is a paid mutator transaction binding the contract method 0x8fc46be5.
//
// Solidity: function setStrategyEscrowDelay(address strategy, uint32 delay) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactor) SetStrategyEscrowDelay(opts *bind.TransactOpts, strategy common.Address, delay uint32) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.contract.Transact(opts, "setStrategyEscrowDelay", strategy, delay)
}

// SetStrategyEscrowDelay is a paid mutator transaction binding the contract method 0x8fc46be5.
//
// Solidity: function setStrategyEscrowDelay(address strategy, uint32 delay) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) SetStrategyEscrowDelay(strategy common.Address, delay uint32) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.SetStrategyEscrowDelay(&_SlashEscrowFactoryStorage.TransactOpts, strategy, delay)
}

// SetStrategyEscrowDelay is a paid mutator transaction binding the contract method 0x8fc46be5.
//
// Solidity: function setStrategyEscrowDelay(address strategy, uint32 delay) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactorSession) SetStrategyEscrowDelay(strategy common.Address, delay uint32) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.SetStrategyEscrowDelay(&_SlashEscrowFactoryStorage.TransactOpts, strategy, delay)
}

// UnpauseEscrow is a paid mutator transaction binding the contract method 0x9b122356.
//
// Solidity: function unpauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactor) UnpauseEscrow(opts *bind.TransactOpts, operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.contract.Transact(opts, "unpauseEscrow", operatorSet, slashId)
}

// UnpauseEscrow is a paid mutator transaction binding the contract method 0x9b122356.
//
// Solidity: function unpauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageSession) UnpauseEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.UnpauseEscrow(&_SlashEscrowFactoryStorage.TransactOpts, operatorSet, slashId)
}

// UnpauseEscrow is a paid mutator transaction binding the contract method 0x9b122356.
//
// Solidity: function unpauseEscrow((address,uint32) operatorSet, uint256 slashId) returns()
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageTransactorSession) UnpauseEscrow(operatorSet OperatorSet, slashId *big.Int) (*types.Transaction, error) {
	return _SlashEscrowFactoryStorage.Contract.UnpauseEscrow(&_SlashEscrowFactoryStorage.TransactOpts, operatorSet, slashId)
}

// SlashEscrowFactoryStorageEscrowCompleteIterator is returned from FilterEscrowComplete and is used to iterate over the raw logs and unpacked data for EscrowComplete events raised by the SlashEscrowFactoryStorage contract.
type SlashEscrowFactoryStorageEscrowCompleteIterator struct {
	Event *SlashEscrowFactoryStorageEscrowComplete // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryStorageEscrowCompleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryStorageEscrowComplete)
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
		it.Event = new(SlashEscrowFactoryStorageEscrowComplete)
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
func (it *SlashEscrowFactoryStorageEscrowCompleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryStorageEscrowCompleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryStorageEscrowComplete represents a EscrowComplete event raised by the SlashEscrowFactoryStorage contract.
type SlashEscrowFactoryStorageEscrowComplete struct {
	OperatorSet OperatorSet
	SlashId     *big.Int
	Strategy    common.Address
	Recipient   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEscrowComplete is a free log retrieval operation binding the contract event 0x32be306ad5a833e756b7cb9724d5312afe0feda6163bfc2dd98ee713346a9abc.
//
// Solidity: event EscrowComplete((address,uint32) operatorSet, uint256 slashId, address strategy, address recipient)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) FilterEscrowComplete(opts *bind.FilterOpts) (*SlashEscrowFactoryStorageEscrowCompleteIterator, error) {

	logs, sub, err := _SlashEscrowFactoryStorage.contract.FilterLogs(opts, "EscrowComplete")
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryStorageEscrowCompleteIterator{contract: _SlashEscrowFactoryStorage.contract, event: "EscrowComplete", logs: logs, sub: sub}, nil
}

// WatchEscrowComplete is a free log subscription operation binding the contract event 0x32be306ad5a833e756b7cb9724d5312afe0feda6163bfc2dd98ee713346a9abc.
//
// Solidity: event EscrowComplete((address,uint32) operatorSet, uint256 slashId, address strategy, address recipient)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) WatchEscrowComplete(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryStorageEscrowComplete) (event.Subscription, error) {

	logs, sub, err := _SlashEscrowFactoryStorage.contract.WatchLogs(opts, "EscrowComplete")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryStorageEscrowComplete)
				if err := _SlashEscrowFactoryStorage.contract.UnpackLog(event, "EscrowComplete", log); err != nil {
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
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) ParseEscrowComplete(log types.Log) (*SlashEscrowFactoryStorageEscrowComplete, error) {
	event := new(SlashEscrowFactoryStorageEscrowComplete)
	if err := _SlashEscrowFactoryStorage.contract.UnpackLog(event, "EscrowComplete", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashEscrowFactoryStorageEscrowPausedIterator is returned from FilterEscrowPaused and is used to iterate over the raw logs and unpacked data for EscrowPaused events raised by the SlashEscrowFactoryStorage contract.
type SlashEscrowFactoryStorageEscrowPausedIterator struct {
	Event *SlashEscrowFactoryStorageEscrowPaused // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryStorageEscrowPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryStorageEscrowPaused)
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
		it.Event = new(SlashEscrowFactoryStorageEscrowPaused)
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
func (it *SlashEscrowFactoryStorageEscrowPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryStorageEscrowPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryStorageEscrowPaused represents a EscrowPaused event raised by the SlashEscrowFactoryStorage contract.
type SlashEscrowFactoryStorageEscrowPaused struct {
	OperatorSet OperatorSet
	SlashId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEscrowPaused is a free log retrieval operation binding the contract event 0x050add19b1a78a4240cdebc8747899275f2dd070c88e83904a37ff7d1a539744.
//
// Solidity: event EscrowPaused((address,uint32) operatorSet, uint256 slashId)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) FilterEscrowPaused(opts *bind.FilterOpts) (*SlashEscrowFactoryStorageEscrowPausedIterator, error) {

	logs, sub, err := _SlashEscrowFactoryStorage.contract.FilterLogs(opts, "EscrowPaused")
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryStorageEscrowPausedIterator{contract: _SlashEscrowFactoryStorage.contract, event: "EscrowPaused", logs: logs, sub: sub}, nil
}

// WatchEscrowPaused is a free log subscription operation binding the contract event 0x050add19b1a78a4240cdebc8747899275f2dd070c88e83904a37ff7d1a539744.
//
// Solidity: event EscrowPaused((address,uint32) operatorSet, uint256 slashId)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) WatchEscrowPaused(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryStorageEscrowPaused) (event.Subscription, error) {

	logs, sub, err := _SlashEscrowFactoryStorage.contract.WatchLogs(opts, "EscrowPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryStorageEscrowPaused)
				if err := _SlashEscrowFactoryStorage.contract.UnpackLog(event, "EscrowPaused", log); err != nil {
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
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) ParseEscrowPaused(log types.Log) (*SlashEscrowFactoryStorageEscrowPaused, error) {
	event := new(SlashEscrowFactoryStorageEscrowPaused)
	if err := _SlashEscrowFactoryStorage.contract.UnpackLog(event, "EscrowPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashEscrowFactoryStorageEscrowUnpausedIterator is returned from FilterEscrowUnpaused and is used to iterate over the raw logs and unpacked data for EscrowUnpaused events raised by the SlashEscrowFactoryStorage contract.
type SlashEscrowFactoryStorageEscrowUnpausedIterator struct {
	Event *SlashEscrowFactoryStorageEscrowUnpaused // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryStorageEscrowUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryStorageEscrowUnpaused)
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
		it.Event = new(SlashEscrowFactoryStorageEscrowUnpaused)
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
func (it *SlashEscrowFactoryStorageEscrowUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryStorageEscrowUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryStorageEscrowUnpaused represents a EscrowUnpaused event raised by the SlashEscrowFactoryStorage contract.
type SlashEscrowFactoryStorageEscrowUnpaused struct {
	OperatorSet OperatorSet
	SlashId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEscrowUnpaused is a free log retrieval operation binding the contract event 0xb8877c6bf02d5f6603188eb653c9269271836b75b2012a5d7f5f233e7cf2f241.
//
// Solidity: event EscrowUnpaused((address,uint32) operatorSet, uint256 slashId)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) FilterEscrowUnpaused(opts *bind.FilterOpts) (*SlashEscrowFactoryStorageEscrowUnpausedIterator, error) {

	logs, sub, err := _SlashEscrowFactoryStorage.contract.FilterLogs(opts, "EscrowUnpaused")
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryStorageEscrowUnpausedIterator{contract: _SlashEscrowFactoryStorage.contract, event: "EscrowUnpaused", logs: logs, sub: sub}, nil
}

// WatchEscrowUnpaused is a free log subscription operation binding the contract event 0xb8877c6bf02d5f6603188eb653c9269271836b75b2012a5d7f5f233e7cf2f241.
//
// Solidity: event EscrowUnpaused((address,uint32) operatorSet, uint256 slashId)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) WatchEscrowUnpaused(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryStorageEscrowUnpaused) (event.Subscription, error) {

	logs, sub, err := _SlashEscrowFactoryStorage.contract.WatchLogs(opts, "EscrowUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryStorageEscrowUnpaused)
				if err := _SlashEscrowFactoryStorage.contract.UnpackLog(event, "EscrowUnpaused", log); err != nil {
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
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) ParseEscrowUnpaused(log types.Log) (*SlashEscrowFactoryStorageEscrowUnpaused, error) {
	event := new(SlashEscrowFactoryStorageEscrowUnpaused)
	if err := _SlashEscrowFactoryStorage.contract.UnpackLog(event, "EscrowUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashEscrowFactoryStorageGlobalEscrowDelaySetIterator is returned from FilterGlobalEscrowDelaySet and is used to iterate over the raw logs and unpacked data for GlobalEscrowDelaySet events raised by the SlashEscrowFactoryStorage contract.
type SlashEscrowFactoryStorageGlobalEscrowDelaySetIterator struct {
	Event *SlashEscrowFactoryStorageGlobalEscrowDelaySet // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryStorageGlobalEscrowDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryStorageGlobalEscrowDelaySet)
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
		it.Event = new(SlashEscrowFactoryStorageGlobalEscrowDelaySet)
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
func (it *SlashEscrowFactoryStorageGlobalEscrowDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryStorageGlobalEscrowDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryStorageGlobalEscrowDelaySet represents a GlobalEscrowDelaySet event raised by the SlashEscrowFactoryStorage contract.
type SlashEscrowFactoryStorageGlobalEscrowDelaySet struct {
	Delay uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGlobalEscrowDelaySet is a free log retrieval operation binding the contract event 0x67d0077d22e4e06f761dd87f6c9f2310ac879c9ce17de50d381e05b72f45fbf6.
//
// Solidity: event GlobalEscrowDelaySet(uint32 delay)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) FilterGlobalEscrowDelaySet(opts *bind.FilterOpts) (*SlashEscrowFactoryStorageGlobalEscrowDelaySetIterator, error) {

	logs, sub, err := _SlashEscrowFactoryStorage.contract.FilterLogs(opts, "GlobalEscrowDelaySet")
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryStorageGlobalEscrowDelaySetIterator{contract: _SlashEscrowFactoryStorage.contract, event: "GlobalEscrowDelaySet", logs: logs, sub: sub}, nil
}

// WatchGlobalEscrowDelaySet is a free log subscription operation binding the contract event 0x67d0077d22e4e06f761dd87f6c9f2310ac879c9ce17de50d381e05b72f45fbf6.
//
// Solidity: event GlobalEscrowDelaySet(uint32 delay)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) WatchGlobalEscrowDelaySet(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryStorageGlobalEscrowDelaySet) (event.Subscription, error) {

	logs, sub, err := _SlashEscrowFactoryStorage.contract.WatchLogs(opts, "GlobalEscrowDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryStorageGlobalEscrowDelaySet)
				if err := _SlashEscrowFactoryStorage.contract.UnpackLog(event, "GlobalEscrowDelaySet", log); err != nil {
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
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) ParseGlobalEscrowDelaySet(log types.Log) (*SlashEscrowFactoryStorageGlobalEscrowDelaySet, error) {
	event := new(SlashEscrowFactoryStorageGlobalEscrowDelaySet)
	if err := _SlashEscrowFactoryStorage.contract.UnpackLog(event, "GlobalEscrowDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashEscrowFactoryStorageStartEscrowIterator is returned from FilterStartEscrow and is used to iterate over the raw logs and unpacked data for StartEscrow events raised by the SlashEscrowFactoryStorage contract.
type SlashEscrowFactoryStorageStartEscrowIterator struct {
	Event *SlashEscrowFactoryStorageStartEscrow // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryStorageStartEscrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryStorageStartEscrow)
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
		it.Event = new(SlashEscrowFactoryStorageStartEscrow)
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
func (it *SlashEscrowFactoryStorageStartEscrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryStorageStartEscrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryStorageStartEscrow represents a StartEscrow event raised by the SlashEscrowFactoryStorage contract.
type SlashEscrowFactoryStorageStartEscrow struct {
	OperatorSet OperatorSet
	SlashId     *big.Int
	Strategy    common.Address
	StartBlock  uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStartEscrow is a free log retrieval operation binding the contract event 0x3afae24c1d3dd2ce3649054ad82458a8c9967ebd9ce10a9a5a3d059f55bfaedb.
//
// Solidity: event StartEscrow((address,uint32) operatorSet, uint256 slashId, address strategy, uint32 startBlock)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) FilterStartEscrow(opts *bind.FilterOpts) (*SlashEscrowFactoryStorageStartEscrowIterator, error) {

	logs, sub, err := _SlashEscrowFactoryStorage.contract.FilterLogs(opts, "StartEscrow")
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryStorageStartEscrowIterator{contract: _SlashEscrowFactoryStorage.contract, event: "StartEscrow", logs: logs, sub: sub}, nil
}

// WatchStartEscrow is a free log subscription operation binding the contract event 0x3afae24c1d3dd2ce3649054ad82458a8c9967ebd9ce10a9a5a3d059f55bfaedb.
//
// Solidity: event StartEscrow((address,uint32) operatorSet, uint256 slashId, address strategy, uint32 startBlock)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) WatchStartEscrow(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryStorageStartEscrow) (event.Subscription, error) {

	logs, sub, err := _SlashEscrowFactoryStorage.contract.WatchLogs(opts, "StartEscrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryStorageStartEscrow)
				if err := _SlashEscrowFactoryStorage.contract.UnpackLog(event, "StartEscrow", log); err != nil {
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
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) ParseStartEscrow(log types.Log) (*SlashEscrowFactoryStorageStartEscrow, error) {
	event := new(SlashEscrowFactoryStorageStartEscrow)
	if err := _SlashEscrowFactoryStorage.contract.UnpackLog(event, "StartEscrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashEscrowFactoryStorageStrategyEscrowDelaySetIterator is returned from FilterStrategyEscrowDelaySet and is used to iterate over the raw logs and unpacked data for StrategyEscrowDelaySet events raised by the SlashEscrowFactoryStorage contract.
type SlashEscrowFactoryStorageStrategyEscrowDelaySetIterator struct {
	Event *SlashEscrowFactoryStorageStrategyEscrowDelaySet // Event containing the contract specifics and raw log

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
func (it *SlashEscrowFactoryStorageStrategyEscrowDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashEscrowFactoryStorageStrategyEscrowDelaySet)
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
		it.Event = new(SlashEscrowFactoryStorageStrategyEscrowDelaySet)
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
func (it *SlashEscrowFactoryStorageStrategyEscrowDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashEscrowFactoryStorageStrategyEscrowDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashEscrowFactoryStorageStrategyEscrowDelaySet represents a StrategyEscrowDelaySet event raised by the SlashEscrowFactoryStorage contract.
type SlashEscrowFactoryStorageStrategyEscrowDelaySet struct {
	Strategy common.Address
	Delay    uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyEscrowDelaySet is a free log retrieval operation binding the contract event 0x5d2b33f07ae22a809e79005f96ffac70c3715df85a3b011b025e0a86a23a007b.
//
// Solidity: event StrategyEscrowDelaySet(address strategy, uint32 delay)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) FilterStrategyEscrowDelaySet(opts *bind.FilterOpts) (*SlashEscrowFactoryStorageStrategyEscrowDelaySetIterator, error) {

	logs, sub, err := _SlashEscrowFactoryStorage.contract.FilterLogs(opts, "StrategyEscrowDelaySet")
	if err != nil {
		return nil, err
	}
	return &SlashEscrowFactoryStorageStrategyEscrowDelaySetIterator{contract: _SlashEscrowFactoryStorage.contract, event: "StrategyEscrowDelaySet", logs: logs, sub: sub}, nil
}

// WatchStrategyEscrowDelaySet is a free log subscription operation binding the contract event 0x5d2b33f07ae22a809e79005f96ffac70c3715df85a3b011b025e0a86a23a007b.
//
// Solidity: event StrategyEscrowDelaySet(address strategy, uint32 delay)
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) WatchStrategyEscrowDelaySet(opts *bind.WatchOpts, sink chan<- *SlashEscrowFactoryStorageStrategyEscrowDelaySet) (event.Subscription, error) {

	logs, sub, err := _SlashEscrowFactoryStorage.contract.WatchLogs(opts, "StrategyEscrowDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashEscrowFactoryStorageStrategyEscrowDelaySet)
				if err := _SlashEscrowFactoryStorage.contract.UnpackLog(event, "StrategyEscrowDelaySet", log); err != nil {
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
func (_SlashEscrowFactoryStorage *SlashEscrowFactoryStorageFilterer) ParseStrategyEscrowDelaySet(log types.Log) (*SlashEscrowFactoryStorageStrategyEscrowDelaySet, error) {
	event := new(SlashEscrowFactoryStorageStrategyEscrowDelaySet)
	if err := _SlashEscrowFactoryStorage.contract.UnpackLog(event, "StrategyEscrowDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
