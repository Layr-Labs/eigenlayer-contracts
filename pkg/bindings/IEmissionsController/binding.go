// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IEmissionsController

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

// IEmissionsControllerTypesDistribution is an auto generated low-level Go binding around an user-defined struct.
type IEmissionsControllerTypesDistribution struct {
	Weight                   uint64
	StartEpoch               uint64
	TotalEpochs              uint64
	DistributionType         uint8
	OperatorSet              OperatorSet
	StrategiesAndMultipliers [][]IRewardsCoordinatorTypesStrategyAndMultiplier
}

// IRewardsCoordinatorTypesStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// IEmissionsControllerMetaData contains all meta data concerning the IEmissionsController contract.
var IEmissionsControllerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ALLOCATION_MANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BACKING_EIGEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBackingEigen\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EIGEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigen\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EMISSIONS_EPOCH_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EMISSIONS_INFLATION_RATE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EMISSIONS_START_TIME\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_TOTAL_WEIGHT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REWARDS_COORDINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRewardsCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addDistribution\",\"inputs\":[{\"name\":\"distribution\",\"type\":\"tuple\",\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"components\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalEpochs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"distributionType\",\"type\":\"uint8\",\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"outputs\":[{\"name\":\"distributionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistribution\",\"inputs\":[{\"name\":\"distributionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"components\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalEpochs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"distributionType\",\"type\":\"uint8\",\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributions\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIEmissionsControllerTypes.Distribution[]\",\"components\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalEpochs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"distributionType\",\"type\":\"uint8\",\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalProcessableDistributions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"incentiveCouncil\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"incentiveCouncil\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isButtonPressable\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastTimeButtonPressable\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextTimeButtonPressable\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pressButton\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIncentiveCouncil\",\"inputs\":[{\"name\":\"incentiveCouncil\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sweep\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalWeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateDistribution\",\"inputs\":[{\"name\":\"distributionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"distribution\",\"type\":\"tuple\",\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"components\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalEpochs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"distributionType\",\"type\":\"uint8\",\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DistributionAdded\",\"inputs\":[{\"name\":\"distributionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"distribution\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"components\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalEpochs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"distributionType\",\"type\":\"uint8\",\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionProcessed\",\"inputs\":[{\"name\":\"distributionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"distribution\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"components\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalEpochs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"distributionType\",\"type\":\"uint8\",\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionUpdated\",\"inputs\":[{\"name\":\"distributionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"distribution\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"components\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalEpochs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"distributionType\",\"type\":\"uint8\",\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncentiveCouncilUpdated\",\"inputs\":[{\"name\":\"incentiveCouncil\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Swept\",\"inputs\":[{\"name\":\"incentiveCouncil\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllDistributionsMustBeProcessed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AllDistributionsProcessed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotIncentiveCouncil\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotAddDisabledDistribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmissionsNotStarted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EpochLengthNotAlignedWithCalculationInterval\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDistributionType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaliciousCallDetected\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardsSubmissionsCannotBeEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartEpochMustBeInTheFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartTimeNotAlignedWithCalculationInterval\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TotalWeightExceedsMax\",\"inputs\":[]}]",
}

// IEmissionsControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use IEmissionsControllerMetaData.ABI instead.
var IEmissionsControllerABI = IEmissionsControllerMetaData.ABI

// IEmissionsController is an auto generated Go binding around an Ethereum contract.
type IEmissionsController struct {
	IEmissionsControllerCaller     // Read-only binding to the contract
	IEmissionsControllerTransactor // Write-only binding to the contract
	IEmissionsControllerFilterer   // Log filterer for contract events
}

// IEmissionsControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEmissionsControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEmissionsControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEmissionsControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEmissionsControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEmissionsControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEmissionsControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEmissionsControllerSession struct {
	Contract     *IEmissionsController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IEmissionsControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEmissionsControllerCallerSession struct {
	Contract *IEmissionsControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IEmissionsControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEmissionsControllerTransactorSession struct {
	Contract     *IEmissionsControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IEmissionsControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEmissionsControllerRaw struct {
	Contract *IEmissionsController // Generic contract binding to access the raw methods on
}

// IEmissionsControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEmissionsControllerCallerRaw struct {
	Contract *IEmissionsControllerCaller // Generic read-only contract binding to access the raw methods on
}

// IEmissionsControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEmissionsControllerTransactorRaw struct {
	Contract *IEmissionsControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEmissionsController creates a new instance of IEmissionsController, bound to a specific deployed contract.
func NewIEmissionsController(address common.Address, backend bind.ContractBackend) (*IEmissionsController, error) {
	contract, err := bindIEmissionsController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEmissionsController{IEmissionsControllerCaller: IEmissionsControllerCaller{contract: contract}, IEmissionsControllerTransactor: IEmissionsControllerTransactor{contract: contract}, IEmissionsControllerFilterer: IEmissionsControllerFilterer{contract: contract}}, nil
}

// NewIEmissionsControllerCaller creates a new read-only instance of IEmissionsController, bound to a specific deployed contract.
func NewIEmissionsControllerCaller(address common.Address, caller bind.ContractCaller) (*IEmissionsControllerCaller, error) {
	contract, err := bindIEmissionsController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEmissionsControllerCaller{contract: contract}, nil
}

// NewIEmissionsControllerTransactor creates a new write-only instance of IEmissionsController, bound to a specific deployed contract.
func NewIEmissionsControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*IEmissionsControllerTransactor, error) {
	contract, err := bindIEmissionsController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEmissionsControllerTransactor{contract: contract}, nil
}

// NewIEmissionsControllerFilterer creates a new log filterer instance of IEmissionsController, bound to a specific deployed contract.
func NewIEmissionsControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*IEmissionsControllerFilterer, error) {
	contract, err := bindIEmissionsController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEmissionsControllerFilterer{contract: contract}, nil
}

// bindIEmissionsController binds a generic wrapper to an already deployed contract.
func bindIEmissionsController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEmissionsControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEmissionsController *IEmissionsControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEmissionsController.Contract.IEmissionsControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEmissionsController *IEmissionsControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEmissionsController.Contract.IEmissionsControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEmissionsController *IEmissionsControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEmissionsController.Contract.IEmissionsControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEmissionsController *IEmissionsControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEmissionsController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEmissionsController *IEmissionsControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEmissionsController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEmissionsController *IEmissionsControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEmissionsController.Contract.contract.Transact(opts, method, params...)
}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_IEmissionsController *IEmissionsControllerCaller) ALLOCATIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "ALLOCATION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_IEmissionsController *IEmissionsControllerSession) ALLOCATIONMANAGER() (common.Address, error) {
	return _IEmissionsController.Contract.ALLOCATIONMANAGER(&_IEmissionsController.CallOpts)
}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_IEmissionsController *IEmissionsControllerCallerSession) ALLOCATIONMANAGER() (common.Address, error) {
	return _IEmissionsController.Contract.ALLOCATIONMANAGER(&_IEmissionsController.CallOpts)
}

// BACKINGEIGEN is a free data retrieval call binding the contract method 0xd455724e.
//
// Solidity: function BACKING_EIGEN() view returns(address)
func (_IEmissionsController *IEmissionsControllerCaller) BACKINGEIGEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "BACKING_EIGEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BACKINGEIGEN is a free data retrieval call binding the contract method 0xd455724e.
//
// Solidity: function BACKING_EIGEN() view returns(address)
func (_IEmissionsController *IEmissionsControllerSession) BACKINGEIGEN() (common.Address, error) {
	return _IEmissionsController.Contract.BACKINGEIGEN(&_IEmissionsController.CallOpts)
}

// BACKINGEIGEN is a free data retrieval call binding the contract method 0xd455724e.
//
// Solidity: function BACKING_EIGEN() view returns(address)
func (_IEmissionsController *IEmissionsControllerCallerSession) BACKINGEIGEN() (common.Address, error) {
	return _IEmissionsController.Contract.BACKINGEIGEN(&_IEmissionsController.CallOpts)
}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_IEmissionsController *IEmissionsControllerCaller) EIGEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "EIGEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_IEmissionsController *IEmissionsControllerSession) EIGEN() (common.Address, error) {
	return _IEmissionsController.Contract.EIGEN(&_IEmissionsController.CallOpts)
}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_IEmissionsController *IEmissionsControllerCallerSession) EIGEN() (common.Address, error) {
	return _IEmissionsController.Contract.EIGEN(&_IEmissionsController.CallOpts)
}

// EMISSIONSEPOCHLENGTH is a free data retrieval call binding the contract method 0xc2f208e4.
//
// Solidity: function EMISSIONS_EPOCH_LENGTH() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCaller) EMISSIONSEPOCHLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "EMISSIONS_EPOCH_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EMISSIONSEPOCHLENGTH is a free data retrieval call binding the contract method 0xc2f208e4.
//
// Solidity: function EMISSIONS_EPOCH_LENGTH() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerSession) EMISSIONSEPOCHLENGTH() (*big.Int, error) {
	return _IEmissionsController.Contract.EMISSIONSEPOCHLENGTH(&_IEmissionsController.CallOpts)
}

// EMISSIONSEPOCHLENGTH is a free data retrieval call binding the contract method 0xc2f208e4.
//
// Solidity: function EMISSIONS_EPOCH_LENGTH() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCallerSession) EMISSIONSEPOCHLENGTH() (*big.Int, error) {
	return _IEmissionsController.Contract.EMISSIONSEPOCHLENGTH(&_IEmissionsController.CallOpts)
}

// EMISSIONSINFLATIONRATE is a free data retrieval call binding the contract method 0x47a28ea2.
//
// Solidity: function EMISSIONS_INFLATION_RATE() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCaller) EMISSIONSINFLATIONRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "EMISSIONS_INFLATION_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EMISSIONSINFLATIONRATE is a free data retrieval call binding the contract method 0x47a28ea2.
//
// Solidity: function EMISSIONS_INFLATION_RATE() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerSession) EMISSIONSINFLATIONRATE() (*big.Int, error) {
	return _IEmissionsController.Contract.EMISSIONSINFLATIONRATE(&_IEmissionsController.CallOpts)
}

// EMISSIONSINFLATIONRATE is a free data retrieval call binding the contract method 0x47a28ea2.
//
// Solidity: function EMISSIONS_INFLATION_RATE() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCallerSession) EMISSIONSINFLATIONRATE() (*big.Int, error) {
	return _IEmissionsController.Contract.EMISSIONSINFLATIONRATE(&_IEmissionsController.CallOpts)
}

// EMISSIONSSTARTTIME is a free data retrieval call binding the contract method 0xc9d3eff9.
//
// Solidity: function EMISSIONS_START_TIME() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCaller) EMISSIONSSTARTTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "EMISSIONS_START_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EMISSIONSSTARTTIME is a free data retrieval call binding the contract method 0xc9d3eff9.
//
// Solidity: function EMISSIONS_START_TIME() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerSession) EMISSIONSSTARTTIME() (*big.Int, error) {
	return _IEmissionsController.Contract.EMISSIONSSTARTTIME(&_IEmissionsController.CallOpts)
}

// EMISSIONSSTARTTIME is a free data retrieval call binding the contract method 0xc9d3eff9.
//
// Solidity: function EMISSIONS_START_TIME() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCallerSession) EMISSIONSSTARTTIME() (*big.Int, error) {
	return _IEmissionsController.Contract.EMISSIONSSTARTTIME(&_IEmissionsController.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "MAX_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _IEmissionsController.Contract.MAXTOTALWEIGHT(&_IEmissionsController.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _IEmissionsController.Contract.MAXTOTALWEIGHT(&_IEmissionsController.CallOpts)
}

// REWARDSCOORDINATOR is a free data retrieval call binding the contract method 0x71e2c264.
//
// Solidity: function REWARDS_COORDINATOR() view returns(address)
func (_IEmissionsController *IEmissionsControllerCaller) REWARDSCOORDINATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "REWARDS_COORDINATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDSCOORDINATOR is a free data retrieval call binding the contract method 0x71e2c264.
//
// Solidity: function REWARDS_COORDINATOR() view returns(address)
func (_IEmissionsController *IEmissionsControllerSession) REWARDSCOORDINATOR() (common.Address, error) {
	return _IEmissionsController.Contract.REWARDSCOORDINATOR(&_IEmissionsController.CallOpts)
}

// REWARDSCOORDINATOR is a free data retrieval call binding the contract method 0x71e2c264.
//
// Solidity: function REWARDS_COORDINATOR() view returns(address)
func (_IEmissionsController *IEmissionsControllerCallerSession) REWARDSCOORDINATOR() (common.Address, error) {
	return _IEmissionsController.Contract.REWARDSCOORDINATOR(&_IEmissionsController.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerSession) GetCurrentEpoch() (*big.Int, error) {
	return _IEmissionsController.Contract.GetCurrentEpoch(&_IEmissionsController.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _IEmissionsController.Contract.GetCurrentEpoch(&_IEmissionsController.CallOpts)
}

// GetDistribution is a free data retrieval call binding the contract method 0x3b345a87.
//
// Solidity: function getDistribution(uint256 distributionId) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]))
func (_IEmissionsController *IEmissionsControllerCaller) GetDistribution(opts *bind.CallOpts, distributionId *big.Int) (IEmissionsControllerTypesDistribution, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "getDistribution", distributionId)

	if err != nil {
		return *new(IEmissionsControllerTypesDistribution), err
	}

	out0 := *abi.ConvertType(out[0], new(IEmissionsControllerTypesDistribution)).(*IEmissionsControllerTypesDistribution)

	return out0, err

}

// GetDistribution is a free data retrieval call binding the contract method 0x3b345a87.
//
// Solidity: function getDistribution(uint256 distributionId) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]))
func (_IEmissionsController *IEmissionsControllerSession) GetDistribution(distributionId *big.Int) (IEmissionsControllerTypesDistribution, error) {
	return _IEmissionsController.Contract.GetDistribution(&_IEmissionsController.CallOpts, distributionId)
}

// GetDistribution is a free data retrieval call binding the contract method 0x3b345a87.
//
// Solidity: function getDistribution(uint256 distributionId) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]))
func (_IEmissionsController *IEmissionsControllerCallerSession) GetDistribution(distributionId *big.Int) (IEmissionsControllerTypesDistribution, error) {
	return _IEmissionsController.Contract.GetDistribution(&_IEmissionsController.CallOpts, distributionId)
}

// GetDistributions is a free data retrieval call binding the contract method 0x147a7a5b.
//
// Solidity: function getDistributions(uint256 start, uint256 length) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][])[])
func (_IEmissionsController *IEmissionsControllerCaller) GetDistributions(opts *bind.CallOpts, start *big.Int, length *big.Int) ([]IEmissionsControllerTypesDistribution, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "getDistributions", start, length)

	if err != nil {
		return *new([]IEmissionsControllerTypesDistribution), err
	}

	out0 := *abi.ConvertType(out[0], new([]IEmissionsControllerTypesDistribution)).(*[]IEmissionsControllerTypesDistribution)

	return out0, err

}

// GetDistributions is a free data retrieval call binding the contract method 0x147a7a5b.
//
// Solidity: function getDistributions(uint256 start, uint256 length) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][])[])
func (_IEmissionsController *IEmissionsControllerSession) GetDistributions(start *big.Int, length *big.Int) ([]IEmissionsControllerTypesDistribution, error) {
	return _IEmissionsController.Contract.GetDistributions(&_IEmissionsController.CallOpts, start, length)
}

// GetDistributions is a free data retrieval call binding the contract method 0x147a7a5b.
//
// Solidity: function getDistributions(uint256 start, uint256 length) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][])[])
func (_IEmissionsController *IEmissionsControllerCallerSession) GetDistributions(start *big.Int, length *big.Int) ([]IEmissionsControllerTypesDistribution, error) {
	return _IEmissionsController.Contract.GetDistributions(&_IEmissionsController.CallOpts, start, length)
}

// GetTotalProcessableDistributions is a free data retrieval call binding the contract method 0xbe851337.
//
// Solidity: function getTotalProcessableDistributions() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCaller) GetTotalProcessableDistributions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "getTotalProcessableDistributions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalProcessableDistributions is a free data retrieval call binding the contract method 0xbe851337.
//
// Solidity: function getTotalProcessableDistributions() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerSession) GetTotalProcessableDistributions() (*big.Int, error) {
	return _IEmissionsController.Contract.GetTotalProcessableDistributions(&_IEmissionsController.CallOpts)
}

// GetTotalProcessableDistributions is a free data retrieval call binding the contract method 0xbe851337.
//
// Solidity: function getTotalProcessableDistributions() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCallerSession) GetTotalProcessableDistributions() (*big.Int, error) {
	return _IEmissionsController.Contract.GetTotalProcessableDistributions(&_IEmissionsController.CallOpts)
}

// IncentiveCouncil is a free data retrieval call binding the contract method 0xc44cb727.
//
// Solidity: function incentiveCouncil() view returns(address)
func (_IEmissionsController *IEmissionsControllerCaller) IncentiveCouncil(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "incentiveCouncil")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IncentiveCouncil is a free data retrieval call binding the contract method 0xc44cb727.
//
// Solidity: function incentiveCouncil() view returns(address)
func (_IEmissionsController *IEmissionsControllerSession) IncentiveCouncil() (common.Address, error) {
	return _IEmissionsController.Contract.IncentiveCouncil(&_IEmissionsController.CallOpts)
}

// IncentiveCouncil is a free data retrieval call binding the contract method 0xc44cb727.
//
// Solidity: function incentiveCouncil() view returns(address)
func (_IEmissionsController *IEmissionsControllerCallerSession) IncentiveCouncil() (common.Address, error) {
	return _IEmissionsController.Contract.IncentiveCouncil(&_IEmissionsController.CallOpts)
}

// IsButtonPressable is a free data retrieval call binding the contract method 0xd8393150.
//
// Solidity: function isButtonPressable() view returns(bool)
func (_IEmissionsController *IEmissionsControllerCaller) IsButtonPressable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "isButtonPressable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsButtonPressable is a free data retrieval call binding the contract method 0xd8393150.
//
// Solidity: function isButtonPressable() view returns(bool)
func (_IEmissionsController *IEmissionsControllerSession) IsButtonPressable() (bool, error) {
	return _IEmissionsController.Contract.IsButtonPressable(&_IEmissionsController.CallOpts)
}

// IsButtonPressable is a free data retrieval call binding the contract method 0xd8393150.
//
// Solidity: function isButtonPressable() view returns(bool)
func (_IEmissionsController *IEmissionsControllerCallerSession) IsButtonPressable() (bool, error) {
	return _IEmissionsController.Contract.IsButtonPressable(&_IEmissionsController.CallOpts)
}

// LastTimeButtonPressable is a free data retrieval call binding the contract method 0xd44b1c9e.
//
// Solidity: function lastTimeButtonPressable() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCaller) LastTimeButtonPressable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "lastTimeButtonPressable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeButtonPressable is a free data retrieval call binding the contract method 0xd44b1c9e.
//
// Solidity: function lastTimeButtonPressable() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerSession) LastTimeButtonPressable() (*big.Int, error) {
	return _IEmissionsController.Contract.LastTimeButtonPressable(&_IEmissionsController.CallOpts)
}

// LastTimeButtonPressable is a free data retrieval call binding the contract method 0xd44b1c9e.
//
// Solidity: function lastTimeButtonPressable() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCallerSession) LastTimeButtonPressable() (*big.Int, error) {
	return _IEmissionsController.Contract.LastTimeButtonPressable(&_IEmissionsController.CallOpts)
}

// NextTimeButtonPressable is a free data retrieval call binding the contract method 0xf769479f.
//
// Solidity: function nextTimeButtonPressable() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCaller) NextTimeButtonPressable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "nextTimeButtonPressable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTimeButtonPressable is a free data retrieval call binding the contract method 0xf769479f.
//
// Solidity: function nextTimeButtonPressable() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerSession) NextTimeButtonPressable() (*big.Int, error) {
	return _IEmissionsController.Contract.NextTimeButtonPressable(&_IEmissionsController.CallOpts)
}

// NextTimeButtonPressable is a free data retrieval call binding the contract method 0xf769479f.
//
// Solidity: function nextTimeButtonPressable() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCallerSession) NextTimeButtonPressable() (*big.Int, error) {
	return _IEmissionsController.Contract.NextTimeButtonPressable(&_IEmissionsController.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_IEmissionsController *IEmissionsControllerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_IEmissionsController *IEmissionsControllerSession) Paused(index uint8) (bool, error) {
	return _IEmissionsController.Contract.Paused(&_IEmissionsController.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_IEmissionsController *IEmissionsControllerCallerSession) Paused(index uint8) (bool, error) {
	return _IEmissionsController.Contract.Paused(&_IEmissionsController.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerSession) Paused0() (*big.Int, error) {
	return _IEmissionsController.Contract.Paused0(&_IEmissionsController.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_IEmissionsController *IEmissionsControllerCallerSession) Paused0() (*big.Int, error) {
	return _IEmissionsController.Contract.Paused0(&_IEmissionsController.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_IEmissionsController *IEmissionsControllerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_IEmissionsController *IEmissionsControllerSession) PauserRegistry() (common.Address, error) {
	return _IEmissionsController.Contract.PauserRegistry(&_IEmissionsController.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_IEmissionsController *IEmissionsControllerCallerSession) PauserRegistry() (common.Address, error) {
	return _IEmissionsController.Contract.PauserRegistry(&_IEmissionsController.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint16)
func (_IEmissionsController *IEmissionsControllerCaller) TotalWeight(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _IEmissionsController.contract.Call(opts, &out, "totalWeight")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint16)
func (_IEmissionsController *IEmissionsControllerSession) TotalWeight() (uint16, error) {
	return _IEmissionsController.Contract.TotalWeight(&_IEmissionsController.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint16)
func (_IEmissionsController *IEmissionsControllerCallerSession) TotalWeight() (uint16, error) {
	return _IEmissionsController.Contract.TotalWeight(&_IEmissionsController.CallOpts)
}

// AddDistribution is a paid mutator transaction binding the contract method 0xcd1e341b.
//
// Solidity: function addDistribution((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns(uint256 distributionId)
func (_IEmissionsController *IEmissionsControllerTransactor) AddDistribution(opts *bind.TransactOpts, distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _IEmissionsController.contract.Transact(opts, "addDistribution", distribution)
}

// AddDistribution is a paid mutator transaction binding the contract method 0xcd1e341b.
//
// Solidity: function addDistribution((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns(uint256 distributionId)
func (_IEmissionsController *IEmissionsControllerSession) AddDistribution(distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _IEmissionsController.Contract.AddDistribution(&_IEmissionsController.TransactOpts, distribution)
}

// AddDistribution is a paid mutator transaction binding the contract method 0xcd1e341b.
//
// Solidity: function addDistribution((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns(uint256 distributionId)
func (_IEmissionsController *IEmissionsControllerTransactorSession) AddDistribution(distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _IEmissionsController.Contract.AddDistribution(&_IEmissionsController.TransactOpts, distribution)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address incentiveCouncil, uint256 initialPausedStatus) returns()
func (_IEmissionsController *IEmissionsControllerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, incentiveCouncil common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _IEmissionsController.contract.Transact(opts, "initialize", initialOwner, incentiveCouncil, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address incentiveCouncil, uint256 initialPausedStatus) returns()
func (_IEmissionsController *IEmissionsControllerSession) Initialize(initialOwner common.Address, incentiveCouncil common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _IEmissionsController.Contract.Initialize(&_IEmissionsController.TransactOpts, initialOwner, incentiveCouncil, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address incentiveCouncil, uint256 initialPausedStatus) returns()
func (_IEmissionsController *IEmissionsControllerTransactorSession) Initialize(initialOwner common.Address, incentiveCouncil common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _IEmissionsController.Contract.Initialize(&_IEmissionsController.TransactOpts, initialOwner, incentiveCouncil, initialPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_IEmissionsController *IEmissionsControllerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _IEmissionsController.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_IEmissionsController *IEmissionsControllerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _IEmissionsController.Contract.Pause(&_IEmissionsController.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_IEmissionsController *IEmissionsControllerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _IEmissionsController.Contract.Pause(&_IEmissionsController.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_IEmissionsController *IEmissionsControllerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEmissionsController.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_IEmissionsController *IEmissionsControllerSession) PauseAll() (*types.Transaction, error) {
	return _IEmissionsController.Contract.PauseAll(&_IEmissionsController.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_IEmissionsController *IEmissionsControllerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _IEmissionsController.Contract.PauseAll(&_IEmissionsController.TransactOpts)
}

// PressButton is a paid mutator transaction binding the contract method 0x400efa85.
//
// Solidity: function pressButton(uint256 length) returns()
func (_IEmissionsController *IEmissionsControllerTransactor) PressButton(opts *bind.TransactOpts, length *big.Int) (*types.Transaction, error) {
	return _IEmissionsController.contract.Transact(opts, "pressButton", length)
}

// PressButton is a paid mutator transaction binding the contract method 0x400efa85.
//
// Solidity: function pressButton(uint256 length) returns()
func (_IEmissionsController *IEmissionsControllerSession) PressButton(length *big.Int) (*types.Transaction, error) {
	return _IEmissionsController.Contract.PressButton(&_IEmissionsController.TransactOpts, length)
}

// PressButton is a paid mutator transaction binding the contract method 0x400efa85.
//
// Solidity: function pressButton(uint256 length) returns()
func (_IEmissionsController *IEmissionsControllerTransactorSession) PressButton(length *big.Int) (*types.Transaction, error) {
	return _IEmissionsController.Contract.PressButton(&_IEmissionsController.TransactOpts, length)
}

// SetIncentiveCouncil is a paid mutator transaction binding the contract method 0xc695acdb.
//
// Solidity: function setIncentiveCouncil(address incentiveCouncil) returns()
func (_IEmissionsController *IEmissionsControllerTransactor) SetIncentiveCouncil(opts *bind.TransactOpts, incentiveCouncil common.Address) (*types.Transaction, error) {
	return _IEmissionsController.contract.Transact(opts, "setIncentiveCouncil", incentiveCouncil)
}

// SetIncentiveCouncil is a paid mutator transaction binding the contract method 0xc695acdb.
//
// Solidity: function setIncentiveCouncil(address incentiveCouncil) returns()
func (_IEmissionsController *IEmissionsControllerSession) SetIncentiveCouncil(incentiveCouncil common.Address) (*types.Transaction, error) {
	return _IEmissionsController.Contract.SetIncentiveCouncil(&_IEmissionsController.TransactOpts, incentiveCouncil)
}

// SetIncentiveCouncil is a paid mutator transaction binding the contract method 0xc695acdb.
//
// Solidity: function setIncentiveCouncil(address incentiveCouncil) returns()
func (_IEmissionsController *IEmissionsControllerTransactorSession) SetIncentiveCouncil(incentiveCouncil common.Address) (*types.Transaction, error) {
	return _IEmissionsController.Contract.SetIncentiveCouncil(&_IEmissionsController.TransactOpts, incentiveCouncil)
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_IEmissionsController *IEmissionsControllerTransactor) Sweep(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEmissionsController.contract.Transact(opts, "sweep")
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_IEmissionsController *IEmissionsControllerSession) Sweep() (*types.Transaction, error) {
	return _IEmissionsController.Contract.Sweep(&_IEmissionsController.TransactOpts)
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_IEmissionsController *IEmissionsControllerTransactorSession) Sweep() (*types.Transaction, error) {
	return _IEmissionsController.Contract.Sweep(&_IEmissionsController.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_IEmissionsController *IEmissionsControllerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _IEmissionsController.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_IEmissionsController *IEmissionsControllerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _IEmissionsController.Contract.Unpause(&_IEmissionsController.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_IEmissionsController *IEmissionsControllerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _IEmissionsController.Contract.Unpause(&_IEmissionsController.TransactOpts, newPausedStatus)
}

// UpdateDistribution is a paid mutator transaction binding the contract method 0x44a32028.
//
// Solidity: function updateDistribution(uint256 distributionId, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns()
func (_IEmissionsController *IEmissionsControllerTransactor) UpdateDistribution(opts *bind.TransactOpts, distributionId *big.Int, distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _IEmissionsController.contract.Transact(opts, "updateDistribution", distributionId, distribution)
}

// UpdateDistribution is a paid mutator transaction binding the contract method 0x44a32028.
//
// Solidity: function updateDistribution(uint256 distributionId, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns()
func (_IEmissionsController *IEmissionsControllerSession) UpdateDistribution(distributionId *big.Int, distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _IEmissionsController.Contract.UpdateDistribution(&_IEmissionsController.TransactOpts, distributionId, distribution)
}

// UpdateDistribution is a paid mutator transaction binding the contract method 0x44a32028.
//
// Solidity: function updateDistribution(uint256 distributionId, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns()
func (_IEmissionsController *IEmissionsControllerTransactorSession) UpdateDistribution(distributionId *big.Int, distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _IEmissionsController.Contract.UpdateDistribution(&_IEmissionsController.TransactOpts, distributionId, distribution)
}

// IEmissionsControllerDistributionAddedIterator is returned from FilterDistributionAdded and is used to iterate over the raw logs and unpacked data for DistributionAdded events raised by the IEmissionsController contract.
type IEmissionsControllerDistributionAddedIterator struct {
	Event *IEmissionsControllerDistributionAdded // Event containing the contract specifics and raw log

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
func (it *IEmissionsControllerDistributionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEmissionsControllerDistributionAdded)
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
		it.Event = new(IEmissionsControllerDistributionAdded)
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
func (it *IEmissionsControllerDistributionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEmissionsControllerDistributionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEmissionsControllerDistributionAdded represents a DistributionAdded event raised by the IEmissionsController contract.
type IEmissionsControllerDistributionAdded struct {
	DistributionId *big.Int
	Epoch          *big.Int
	Distribution   IEmissionsControllerTypesDistribution
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDistributionAdded is a free log retrieval operation binding the contract event 0x006f7ba35643ecf5852cfe01b66220b1fe04a4cd4866923d5f3e66c7fcb390ef.
//
// Solidity: event DistributionAdded(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_IEmissionsController *IEmissionsControllerFilterer) FilterDistributionAdded(opts *bind.FilterOpts, distributionId []*big.Int, epoch []*big.Int) (*IEmissionsControllerDistributionAddedIterator, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _IEmissionsController.contract.FilterLogs(opts, "DistributionAdded", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &IEmissionsControllerDistributionAddedIterator{contract: _IEmissionsController.contract, event: "DistributionAdded", logs: logs, sub: sub}, nil
}

// WatchDistributionAdded is a free log subscription operation binding the contract event 0x006f7ba35643ecf5852cfe01b66220b1fe04a4cd4866923d5f3e66c7fcb390ef.
//
// Solidity: event DistributionAdded(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_IEmissionsController *IEmissionsControllerFilterer) WatchDistributionAdded(opts *bind.WatchOpts, sink chan<- *IEmissionsControllerDistributionAdded, distributionId []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _IEmissionsController.contract.WatchLogs(opts, "DistributionAdded", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEmissionsControllerDistributionAdded)
				if err := _IEmissionsController.contract.UnpackLog(event, "DistributionAdded", log); err != nil {
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

// ParseDistributionAdded is a log parse operation binding the contract event 0x006f7ba35643ecf5852cfe01b66220b1fe04a4cd4866923d5f3e66c7fcb390ef.
//
// Solidity: event DistributionAdded(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_IEmissionsController *IEmissionsControllerFilterer) ParseDistributionAdded(log types.Log) (*IEmissionsControllerDistributionAdded, error) {
	event := new(IEmissionsControllerDistributionAdded)
	if err := _IEmissionsController.contract.UnpackLog(event, "DistributionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEmissionsControllerDistributionProcessedIterator is returned from FilterDistributionProcessed and is used to iterate over the raw logs and unpacked data for DistributionProcessed events raised by the IEmissionsController contract.
type IEmissionsControllerDistributionProcessedIterator struct {
	Event *IEmissionsControllerDistributionProcessed // Event containing the contract specifics and raw log

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
func (it *IEmissionsControllerDistributionProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEmissionsControllerDistributionProcessed)
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
		it.Event = new(IEmissionsControllerDistributionProcessed)
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
func (it *IEmissionsControllerDistributionProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEmissionsControllerDistributionProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEmissionsControllerDistributionProcessed represents a DistributionProcessed event raised by the IEmissionsController contract.
type IEmissionsControllerDistributionProcessed struct {
	DistributionId *big.Int
	Epoch          *big.Int
	Distribution   IEmissionsControllerTypesDistribution
	Success        bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDistributionProcessed is a free log retrieval operation binding the contract event 0xba5d66336bc9a4f8459242151a4da4d5020ac581243d98403bb55f7f348e071b.
//
// Solidity: event DistributionProcessed(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution, bool success)
func (_IEmissionsController *IEmissionsControllerFilterer) FilterDistributionProcessed(opts *bind.FilterOpts, distributionId []*big.Int, epoch []*big.Int) (*IEmissionsControllerDistributionProcessedIterator, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _IEmissionsController.contract.FilterLogs(opts, "DistributionProcessed", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &IEmissionsControllerDistributionProcessedIterator{contract: _IEmissionsController.contract, event: "DistributionProcessed", logs: logs, sub: sub}, nil
}

// WatchDistributionProcessed is a free log subscription operation binding the contract event 0xba5d66336bc9a4f8459242151a4da4d5020ac581243d98403bb55f7f348e071b.
//
// Solidity: event DistributionProcessed(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution, bool success)
func (_IEmissionsController *IEmissionsControllerFilterer) WatchDistributionProcessed(opts *bind.WatchOpts, sink chan<- *IEmissionsControllerDistributionProcessed, distributionId []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _IEmissionsController.contract.WatchLogs(opts, "DistributionProcessed", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEmissionsControllerDistributionProcessed)
				if err := _IEmissionsController.contract.UnpackLog(event, "DistributionProcessed", log); err != nil {
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

// ParseDistributionProcessed is a log parse operation binding the contract event 0xba5d66336bc9a4f8459242151a4da4d5020ac581243d98403bb55f7f348e071b.
//
// Solidity: event DistributionProcessed(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution, bool success)
func (_IEmissionsController *IEmissionsControllerFilterer) ParseDistributionProcessed(log types.Log) (*IEmissionsControllerDistributionProcessed, error) {
	event := new(IEmissionsControllerDistributionProcessed)
	if err := _IEmissionsController.contract.UnpackLog(event, "DistributionProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEmissionsControllerDistributionUpdatedIterator is returned from FilterDistributionUpdated and is used to iterate over the raw logs and unpacked data for DistributionUpdated events raised by the IEmissionsController contract.
type IEmissionsControllerDistributionUpdatedIterator struct {
	Event *IEmissionsControllerDistributionUpdated // Event containing the contract specifics and raw log

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
func (it *IEmissionsControllerDistributionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEmissionsControllerDistributionUpdated)
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
		it.Event = new(IEmissionsControllerDistributionUpdated)
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
func (it *IEmissionsControllerDistributionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEmissionsControllerDistributionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEmissionsControllerDistributionUpdated represents a DistributionUpdated event raised by the IEmissionsController contract.
type IEmissionsControllerDistributionUpdated struct {
	DistributionId *big.Int
	Epoch          *big.Int
	Distribution   IEmissionsControllerTypesDistribution
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDistributionUpdated is a free log retrieval operation binding the contract event 0x548fb50d6be978df2bacbf48c6840e4a4743672408921282117f3f00555b2b4c.
//
// Solidity: event DistributionUpdated(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_IEmissionsController *IEmissionsControllerFilterer) FilterDistributionUpdated(opts *bind.FilterOpts, distributionId []*big.Int, epoch []*big.Int) (*IEmissionsControllerDistributionUpdatedIterator, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _IEmissionsController.contract.FilterLogs(opts, "DistributionUpdated", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &IEmissionsControllerDistributionUpdatedIterator{contract: _IEmissionsController.contract, event: "DistributionUpdated", logs: logs, sub: sub}, nil
}

// WatchDistributionUpdated is a free log subscription operation binding the contract event 0x548fb50d6be978df2bacbf48c6840e4a4743672408921282117f3f00555b2b4c.
//
// Solidity: event DistributionUpdated(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_IEmissionsController *IEmissionsControllerFilterer) WatchDistributionUpdated(opts *bind.WatchOpts, sink chan<- *IEmissionsControllerDistributionUpdated, distributionId []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _IEmissionsController.contract.WatchLogs(opts, "DistributionUpdated", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEmissionsControllerDistributionUpdated)
				if err := _IEmissionsController.contract.UnpackLog(event, "DistributionUpdated", log); err != nil {
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

// ParseDistributionUpdated is a log parse operation binding the contract event 0x548fb50d6be978df2bacbf48c6840e4a4743672408921282117f3f00555b2b4c.
//
// Solidity: event DistributionUpdated(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_IEmissionsController *IEmissionsControllerFilterer) ParseDistributionUpdated(log types.Log) (*IEmissionsControllerDistributionUpdated, error) {
	event := new(IEmissionsControllerDistributionUpdated)
	if err := _IEmissionsController.contract.UnpackLog(event, "DistributionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEmissionsControllerIncentiveCouncilUpdatedIterator is returned from FilterIncentiveCouncilUpdated and is used to iterate over the raw logs and unpacked data for IncentiveCouncilUpdated events raised by the IEmissionsController contract.
type IEmissionsControllerIncentiveCouncilUpdatedIterator struct {
	Event *IEmissionsControllerIncentiveCouncilUpdated // Event containing the contract specifics and raw log

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
func (it *IEmissionsControllerIncentiveCouncilUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEmissionsControllerIncentiveCouncilUpdated)
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
		it.Event = new(IEmissionsControllerIncentiveCouncilUpdated)
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
func (it *IEmissionsControllerIncentiveCouncilUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEmissionsControllerIncentiveCouncilUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEmissionsControllerIncentiveCouncilUpdated represents a IncentiveCouncilUpdated event raised by the IEmissionsController contract.
type IEmissionsControllerIncentiveCouncilUpdated struct {
	IncentiveCouncil common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterIncentiveCouncilUpdated is a free log retrieval operation binding the contract event 0x8befac6896b786e67b23cefc473bfabd36e7fc013125c883dfeec8e3a9636216.
//
// Solidity: event IncentiveCouncilUpdated(address indexed incentiveCouncil)
func (_IEmissionsController *IEmissionsControllerFilterer) FilterIncentiveCouncilUpdated(opts *bind.FilterOpts, incentiveCouncil []common.Address) (*IEmissionsControllerIncentiveCouncilUpdatedIterator, error) {

	var incentiveCouncilRule []interface{}
	for _, incentiveCouncilItem := range incentiveCouncil {
		incentiveCouncilRule = append(incentiveCouncilRule, incentiveCouncilItem)
	}

	logs, sub, err := _IEmissionsController.contract.FilterLogs(opts, "IncentiveCouncilUpdated", incentiveCouncilRule)
	if err != nil {
		return nil, err
	}
	return &IEmissionsControllerIncentiveCouncilUpdatedIterator{contract: _IEmissionsController.contract, event: "IncentiveCouncilUpdated", logs: logs, sub: sub}, nil
}

// WatchIncentiveCouncilUpdated is a free log subscription operation binding the contract event 0x8befac6896b786e67b23cefc473bfabd36e7fc013125c883dfeec8e3a9636216.
//
// Solidity: event IncentiveCouncilUpdated(address indexed incentiveCouncil)
func (_IEmissionsController *IEmissionsControllerFilterer) WatchIncentiveCouncilUpdated(opts *bind.WatchOpts, sink chan<- *IEmissionsControllerIncentiveCouncilUpdated, incentiveCouncil []common.Address) (event.Subscription, error) {

	var incentiveCouncilRule []interface{}
	for _, incentiveCouncilItem := range incentiveCouncil {
		incentiveCouncilRule = append(incentiveCouncilRule, incentiveCouncilItem)
	}

	logs, sub, err := _IEmissionsController.contract.WatchLogs(opts, "IncentiveCouncilUpdated", incentiveCouncilRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEmissionsControllerIncentiveCouncilUpdated)
				if err := _IEmissionsController.contract.UnpackLog(event, "IncentiveCouncilUpdated", log); err != nil {
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

// ParseIncentiveCouncilUpdated is a log parse operation binding the contract event 0x8befac6896b786e67b23cefc473bfabd36e7fc013125c883dfeec8e3a9636216.
//
// Solidity: event IncentiveCouncilUpdated(address indexed incentiveCouncil)
func (_IEmissionsController *IEmissionsControllerFilterer) ParseIncentiveCouncilUpdated(log types.Log) (*IEmissionsControllerIncentiveCouncilUpdated, error) {
	event := new(IEmissionsControllerIncentiveCouncilUpdated)
	if err := _IEmissionsController.contract.UnpackLog(event, "IncentiveCouncilUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEmissionsControllerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IEmissionsController contract.
type IEmissionsControllerPausedIterator struct {
	Event *IEmissionsControllerPaused // Event containing the contract specifics and raw log

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
func (it *IEmissionsControllerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEmissionsControllerPaused)
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
		it.Event = new(IEmissionsControllerPaused)
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
func (it *IEmissionsControllerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEmissionsControllerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEmissionsControllerPaused represents a Paused event raised by the IEmissionsController contract.
type IEmissionsControllerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_IEmissionsController *IEmissionsControllerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*IEmissionsControllerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEmissionsController.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &IEmissionsControllerPausedIterator{contract: _IEmissionsController.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_IEmissionsController *IEmissionsControllerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IEmissionsControllerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEmissionsController.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEmissionsControllerPaused)
				if err := _IEmissionsController.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_IEmissionsController *IEmissionsControllerFilterer) ParsePaused(log types.Log) (*IEmissionsControllerPaused, error) {
	event := new(IEmissionsControllerPaused)
	if err := _IEmissionsController.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEmissionsControllerSweptIterator is returned from FilterSwept and is used to iterate over the raw logs and unpacked data for Swept events raised by the IEmissionsController contract.
type IEmissionsControllerSweptIterator struct {
	Event *IEmissionsControllerSwept // Event containing the contract specifics and raw log

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
func (it *IEmissionsControllerSweptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEmissionsControllerSwept)
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
		it.Event = new(IEmissionsControllerSwept)
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
func (it *IEmissionsControllerSweptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEmissionsControllerSweptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEmissionsControllerSwept represents a Swept event raised by the IEmissionsController contract.
type IEmissionsControllerSwept struct {
	IncentiveCouncil common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSwept is a free log retrieval operation binding the contract event 0xc36b5179cb9c303b200074996eab2b3473eac370fdd7eba3bec636fe35109696.
//
// Solidity: event Swept(address indexed incentiveCouncil, uint256 amount)
func (_IEmissionsController *IEmissionsControllerFilterer) FilterSwept(opts *bind.FilterOpts, incentiveCouncil []common.Address) (*IEmissionsControllerSweptIterator, error) {

	var incentiveCouncilRule []interface{}
	for _, incentiveCouncilItem := range incentiveCouncil {
		incentiveCouncilRule = append(incentiveCouncilRule, incentiveCouncilItem)
	}

	logs, sub, err := _IEmissionsController.contract.FilterLogs(opts, "Swept", incentiveCouncilRule)
	if err != nil {
		return nil, err
	}
	return &IEmissionsControllerSweptIterator{contract: _IEmissionsController.contract, event: "Swept", logs: logs, sub: sub}, nil
}

// WatchSwept is a free log subscription operation binding the contract event 0xc36b5179cb9c303b200074996eab2b3473eac370fdd7eba3bec636fe35109696.
//
// Solidity: event Swept(address indexed incentiveCouncil, uint256 amount)
func (_IEmissionsController *IEmissionsControllerFilterer) WatchSwept(opts *bind.WatchOpts, sink chan<- *IEmissionsControllerSwept, incentiveCouncil []common.Address) (event.Subscription, error) {

	var incentiveCouncilRule []interface{}
	for _, incentiveCouncilItem := range incentiveCouncil {
		incentiveCouncilRule = append(incentiveCouncilRule, incentiveCouncilItem)
	}

	logs, sub, err := _IEmissionsController.contract.WatchLogs(opts, "Swept", incentiveCouncilRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEmissionsControllerSwept)
				if err := _IEmissionsController.contract.UnpackLog(event, "Swept", log); err != nil {
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

// ParseSwept is a log parse operation binding the contract event 0xc36b5179cb9c303b200074996eab2b3473eac370fdd7eba3bec636fe35109696.
//
// Solidity: event Swept(address indexed incentiveCouncil, uint256 amount)
func (_IEmissionsController *IEmissionsControllerFilterer) ParseSwept(log types.Log) (*IEmissionsControllerSwept, error) {
	event := new(IEmissionsControllerSwept)
	if err := _IEmissionsController.contract.UnpackLog(event, "Swept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEmissionsControllerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IEmissionsController contract.
type IEmissionsControllerUnpausedIterator struct {
	Event *IEmissionsControllerUnpaused // Event containing the contract specifics and raw log

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
func (it *IEmissionsControllerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEmissionsControllerUnpaused)
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
		it.Event = new(IEmissionsControllerUnpaused)
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
func (it *IEmissionsControllerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEmissionsControllerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEmissionsControllerUnpaused represents a Unpaused event raised by the IEmissionsController contract.
type IEmissionsControllerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_IEmissionsController *IEmissionsControllerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*IEmissionsControllerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEmissionsController.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &IEmissionsControllerUnpausedIterator{contract: _IEmissionsController.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_IEmissionsController *IEmissionsControllerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IEmissionsControllerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEmissionsController.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEmissionsControllerUnpaused)
				if err := _IEmissionsController.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_IEmissionsController *IEmissionsControllerFilterer) ParseUnpaused(log types.Log) (*IEmissionsControllerUnpaused, error) {
	event := new(IEmissionsControllerUnpaused)
	if err := _IEmissionsController.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
