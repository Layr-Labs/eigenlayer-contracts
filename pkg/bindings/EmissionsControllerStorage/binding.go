// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EmissionsControllerStorage

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

// EmissionsControllerStorageMetaData contains all meta data concerning the EmissionsControllerStorage contract.
var EmissionsControllerStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ALLOCATION_MANAGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BACKING_EIGEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBackingEigen\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EIGEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigen\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EMISSIONS_EPOCH_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EMISSIONS_INFLATION_RATE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EMISSIONS_START_TIME\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_TOTAL_WEIGHT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REWARDS_COORDINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRewardsCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addDistribution\",\"inputs\":[{\"name\":\"distribution\",\"type\":\"tuple\",\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"components\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalEpochs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"distributionType\",\"type\":\"uint8\",\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"outputs\":[{\"name\":\"distributionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistribution\",\"inputs\":[{\"name\":\"distributionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"components\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalEpochs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"distributionType\",\"type\":\"uint8\",\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributions\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIEmissionsControllerTypes.Distribution[]\",\"components\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalEpochs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"distributionType\",\"type\":\"uint8\",\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalProcessableDistributions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"incentiveCouncil\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"incentiveCouncil\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isButtonPressable\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastTimeButtonPressable\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextTimeButtonPressable\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pressButton\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIncentiveCouncil\",\"inputs\":[{\"name\":\"incentiveCouncil\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sweep\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalWeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateDistribution\",\"inputs\":[{\"name\":\"distributionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"distribution\",\"type\":\"tuple\",\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"components\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalEpochs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"distributionType\",\"type\":\"uint8\",\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DistributionAdded\",\"inputs\":[{\"name\":\"distributionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"distribution\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"components\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalEpochs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"distributionType\",\"type\":\"uint8\",\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionProcessed\",\"inputs\":[{\"name\":\"distributionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"distribution\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"components\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalEpochs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"distributionType\",\"type\":\"uint8\",\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionUpdated\",\"inputs\":[{\"name\":\"distributionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"distribution\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIEmissionsControllerTypes.Distribution\",\"components\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalEpochs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"distributionType\",\"type\":\"uint8\",\"internalType\":\"enumIEmissionsControllerTypes.DistributionType\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[][]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[][]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncentiveCouncilUpdated\",\"inputs\":[{\"name\":\"incentiveCouncil\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Swept\",\"inputs\":[{\"name\":\"incentiveCouncil\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllDistributionsMustBeProcessed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AllDistributionsProcessed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotIncentiveCouncil\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotAddDisabledDistribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmissionsNotStarted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EpochLengthNotAlignedWithCalculationInterval\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDistributionType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaliciousCallDetected\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorSetNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardsSubmissionsCannotBeEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartEpochMustBeInTheFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartTimeNotAlignedWithCalculationInterval\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TotalWeightExceedsMax\",\"inputs\":[]}]",
}

// EmissionsControllerStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use EmissionsControllerStorageMetaData.ABI instead.
var EmissionsControllerStorageABI = EmissionsControllerStorageMetaData.ABI

// EmissionsControllerStorage is an auto generated Go binding around an Ethereum contract.
type EmissionsControllerStorage struct {
	EmissionsControllerStorageCaller     // Read-only binding to the contract
	EmissionsControllerStorageTransactor // Write-only binding to the contract
	EmissionsControllerStorageFilterer   // Log filterer for contract events
}

// EmissionsControllerStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type EmissionsControllerStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmissionsControllerStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EmissionsControllerStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmissionsControllerStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EmissionsControllerStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmissionsControllerStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EmissionsControllerStorageSession struct {
	Contract     *EmissionsControllerStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EmissionsControllerStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EmissionsControllerStorageCallerSession struct {
	Contract *EmissionsControllerStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// EmissionsControllerStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EmissionsControllerStorageTransactorSession struct {
	Contract     *EmissionsControllerStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// EmissionsControllerStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type EmissionsControllerStorageRaw struct {
	Contract *EmissionsControllerStorage // Generic contract binding to access the raw methods on
}

// EmissionsControllerStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EmissionsControllerStorageCallerRaw struct {
	Contract *EmissionsControllerStorageCaller // Generic read-only contract binding to access the raw methods on
}

// EmissionsControllerStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EmissionsControllerStorageTransactorRaw struct {
	Contract *EmissionsControllerStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEmissionsControllerStorage creates a new instance of EmissionsControllerStorage, bound to a specific deployed contract.
func NewEmissionsControllerStorage(address common.Address, backend bind.ContractBackend) (*EmissionsControllerStorage, error) {
	contract, err := bindEmissionsControllerStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EmissionsControllerStorage{EmissionsControllerStorageCaller: EmissionsControllerStorageCaller{contract: contract}, EmissionsControllerStorageTransactor: EmissionsControllerStorageTransactor{contract: contract}, EmissionsControllerStorageFilterer: EmissionsControllerStorageFilterer{contract: contract}}, nil
}

// NewEmissionsControllerStorageCaller creates a new read-only instance of EmissionsControllerStorage, bound to a specific deployed contract.
func NewEmissionsControllerStorageCaller(address common.Address, caller bind.ContractCaller) (*EmissionsControllerStorageCaller, error) {
	contract, err := bindEmissionsControllerStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EmissionsControllerStorageCaller{contract: contract}, nil
}

// NewEmissionsControllerStorageTransactor creates a new write-only instance of EmissionsControllerStorage, bound to a specific deployed contract.
func NewEmissionsControllerStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*EmissionsControllerStorageTransactor, error) {
	contract, err := bindEmissionsControllerStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EmissionsControllerStorageTransactor{contract: contract}, nil
}

// NewEmissionsControllerStorageFilterer creates a new log filterer instance of EmissionsControllerStorage, bound to a specific deployed contract.
func NewEmissionsControllerStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*EmissionsControllerStorageFilterer, error) {
	contract, err := bindEmissionsControllerStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EmissionsControllerStorageFilterer{contract: contract}, nil
}

// bindEmissionsControllerStorage binds a generic wrapper to an already deployed contract.
func bindEmissionsControllerStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EmissionsControllerStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmissionsControllerStorage *EmissionsControllerStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmissionsControllerStorage.Contract.EmissionsControllerStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmissionsControllerStorage *EmissionsControllerStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.EmissionsControllerStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmissionsControllerStorage *EmissionsControllerStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.EmissionsControllerStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmissionsControllerStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.contract.Transact(opts, method, params...)
}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) ALLOCATIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "ALLOCATION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) ALLOCATIONMANAGER() (common.Address, error) {
	return _EmissionsControllerStorage.Contract.ALLOCATIONMANAGER(&_EmissionsControllerStorage.CallOpts)
}

// ALLOCATIONMANAGER is a free data retrieval call binding the contract method 0x31232bc9.
//
// Solidity: function ALLOCATION_MANAGER() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) ALLOCATIONMANAGER() (common.Address, error) {
	return _EmissionsControllerStorage.Contract.ALLOCATIONMANAGER(&_EmissionsControllerStorage.CallOpts)
}

// BACKINGEIGEN is a free data retrieval call binding the contract method 0xd455724e.
//
// Solidity: function BACKING_EIGEN() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) BACKINGEIGEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "BACKING_EIGEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BACKINGEIGEN is a free data retrieval call binding the contract method 0xd455724e.
//
// Solidity: function BACKING_EIGEN() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) BACKINGEIGEN() (common.Address, error) {
	return _EmissionsControllerStorage.Contract.BACKINGEIGEN(&_EmissionsControllerStorage.CallOpts)
}

// BACKINGEIGEN is a free data retrieval call binding the contract method 0xd455724e.
//
// Solidity: function BACKING_EIGEN() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) BACKINGEIGEN() (common.Address, error) {
	return _EmissionsControllerStorage.Contract.BACKINGEIGEN(&_EmissionsControllerStorage.CallOpts)
}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) EIGEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "EIGEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) EIGEN() (common.Address, error) {
	return _EmissionsControllerStorage.Contract.EIGEN(&_EmissionsControllerStorage.CallOpts)
}

// EIGEN is a free data retrieval call binding the contract method 0xfdc371ce.
//
// Solidity: function EIGEN() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) EIGEN() (common.Address, error) {
	return _EmissionsControllerStorage.Contract.EIGEN(&_EmissionsControllerStorage.CallOpts)
}

// EMISSIONSEPOCHLENGTH is a free data retrieval call binding the contract method 0xc2f208e4.
//
// Solidity: function EMISSIONS_EPOCH_LENGTH() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) EMISSIONSEPOCHLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "EMISSIONS_EPOCH_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EMISSIONSEPOCHLENGTH is a free data retrieval call binding the contract method 0xc2f208e4.
//
// Solidity: function EMISSIONS_EPOCH_LENGTH() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) EMISSIONSEPOCHLENGTH() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.EMISSIONSEPOCHLENGTH(&_EmissionsControllerStorage.CallOpts)
}

// EMISSIONSEPOCHLENGTH is a free data retrieval call binding the contract method 0xc2f208e4.
//
// Solidity: function EMISSIONS_EPOCH_LENGTH() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) EMISSIONSEPOCHLENGTH() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.EMISSIONSEPOCHLENGTH(&_EmissionsControllerStorage.CallOpts)
}

// EMISSIONSINFLATIONRATE is a free data retrieval call binding the contract method 0x47a28ea2.
//
// Solidity: function EMISSIONS_INFLATION_RATE() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) EMISSIONSINFLATIONRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "EMISSIONS_INFLATION_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EMISSIONSINFLATIONRATE is a free data retrieval call binding the contract method 0x47a28ea2.
//
// Solidity: function EMISSIONS_INFLATION_RATE() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) EMISSIONSINFLATIONRATE() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.EMISSIONSINFLATIONRATE(&_EmissionsControllerStorage.CallOpts)
}

// EMISSIONSINFLATIONRATE is a free data retrieval call binding the contract method 0x47a28ea2.
//
// Solidity: function EMISSIONS_INFLATION_RATE() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) EMISSIONSINFLATIONRATE() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.EMISSIONSINFLATIONRATE(&_EmissionsControllerStorage.CallOpts)
}

// EMISSIONSSTARTTIME is a free data retrieval call binding the contract method 0xc9d3eff9.
//
// Solidity: function EMISSIONS_START_TIME() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) EMISSIONSSTARTTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "EMISSIONS_START_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EMISSIONSSTARTTIME is a free data retrieval call binding the contract method 0xc9d3eff9.
//
// Solidity: function EMISSIONS_START_TIME() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) EMISSIONSSTARTTIME() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.EMISSIONSSTARTTIME(&_EmissionsControllerStorage.CallOpts)
}

// EMISSIONSSTARTTIME is a free data retrieval call binding the contract method 0xc9d3eff9.
//
// Solidity: function EMISSIONS_START_TIME() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) EMISSIONSSTARTTIME() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.EMISSIONSSTARTTIME(&_EmissionsControllerStorage.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "MAX_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.MAXTOTALWEIGHT(&_EmissionsControllerStorage.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.MAXTOTALWEIGHT(&_EmissionsControllerStorage.CallOpts)
}

// REWARDSCOORDINATOR is a free data retrieval call binding the contract method 0x71e2c264.
//
// Solidity: function REWARDS_COORDINATOR() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) REWARDSCOORDINATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "REWARDS_COORDINATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDSCOORDINATOR is a free data retrieval call binding the contract method 0x71e2c264.
//
// Solidity: function REWARDS_COORDINATOR() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) REWARDSCOORDINATOR() (common.Address, error) {
	return _EmissionsControllerStorage.Contract.REWARDSCOORDINATOR(&_EmissionsControllerStorage.CallOpts)
}

// REWARDSCOORDINATOR is a free data retrieval call binding the contract method 0x71e2c264.
//
// Solidity: function REWARDS_COORDINATOR() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) REWARDSCOORDINATOR() (common.Address, error) {
	return _EmissionsControllerStorage.Contract.REWARDSCOORDINATOR(&_EmissionsControllerStorage.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) GetCurrentEpoch() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.GetCurrentEpoch(&_EmissionsControllerStorage.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.GetCurrentEpoch(&_EmissionsControllerStorage.CallOpts)
}

// GetDistribution is a free data retrieval call binding the contract method 0x3b345a87.
//
// Solidity: function getDistribution(uint256 distributionId) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]))
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) GetDistribution(opts *bind.CallOpts, distributionId *big.Int) (IEmissionsControllerTypesDistribution, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "getDistribution", distributionId)

	if err != nil {
		return *new(IEmissionsControllerTypesDistribution), err
	}

	out0 := *abi.ConvertType(out[0], new(IEmissionsControllerTypesDistribution)).(*IEmissionsControllerTypesDistribution)

	return out0, err

}

// GetDistribution is a free data retrieval call binding the contract method 0x3b345a87.
//
// Solidity: function getDistribution(uint256 distributionId) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]))
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) GetDistribution(distributionId *big.Int) (IEmissionsControllerTypesDistribution, error) {
	return _EmissionsControllerStorage.Contract.GetDistribution(&_EmissionsControllerStorage.CallOpts, distributionId)
}

// GetDistribution is a free data retrieval call binding the contract method 0x3b345a87.
//
// Solidity: function getDistribution(uint256 distributionId) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]))
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) GetDistribution(distributionId *big.Int) (IEmissionsControllerTypesDistribution, error) {
	return _EmissionsControllerStorage.Contract.GetDistribution(&_EmissionsControllerStorage.CallOpts, distributionId)
}

// GetDistributions is a free data retrieval call binding the contract method 0x147a7a5b.
//
// Solidity: function getDistributions(uint256 start, uint256 length) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][])[])
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) GetDistributions(opts *bind.CallOpts, start *big.Int, length *big.Int) ([]IEmissionsControllerTypesDistribution, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "getDistributions", start, length)

	if err != nil {
		return *new([]IEmissionsControllerTypesDistribution), err
	}

	out0 := *abi.ConvertType(out[0], new([]IEmissionsControllerTypesDistribution)).(*[]IEmissionsControllerTypesDistribution)

	return out0, err

}

// GetDistributions is a free data retrieval call binding the contract method 0x147a7a5b.
//
// Solidity: function getDistributions(uint256 start, uint256 length) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][])[])
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) GetDistributions(start *big.Int, length *big.Int) ([]IEmissionsControllerTypesDistribution, error) {
	return _EmissionsControllerStorage.Contract.GetDistributions(&_EmissionsControllerStorage.CallOpts, start, length)
}

// GetDistributions is a free data retrieval call binding the contract method 0x147a7a5b.
//
// Solidity: function getDistributions(uint256 start, uint256 length) view returns((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][])[])
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) GetDistributions(start *big.Int, length *big.Int) ([]IEmissionsControllerTypesDistribution, error) {
	return _EmissionsControllerStorage.Contract.GetDistributions(&_EmissionsControllerStorage.CallOpts, start, length)
}

// GetTotalProcessableDistributions is a free data retrieval call binding the contract method 0xbe851337.
//
// Solidity: function getTotalProcessableDistributions() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) GetTotalProcessableDistributions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "getTotalProcessableDistributions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalProcessableDistributions is a free data retrieval call binding the contract method 0xbe851337.
//
// Solidity: function getTotalProcessableDistributions() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) GetTotalProcessableDistributions() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.GetTotalProcessableDistributions(&_EmissionsControllerStorage.CallOpts)
}

// GetTotalProcessableDistributions is a free data retrieval call binding the contract method 0xbe851337.
//
// Solidity: function getTotalProcessableDistributions() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) GetTotalProcessableDistributions() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.GetTotalProcessableDistributions(&_EmissionsControllerStorage.CallOpts)
}

// IncentiveCouncil is a free data retrieval call binding the contract method 0xc44cb727.
//
// Solidity: function incentiveCouncil() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) IncentiveCouncil(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "incentiveCouncil")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IncentiveCouncil is a free data retrieval call binding the contract method 0xc44cb727.
//
// Solidity: function incentiveCouncil() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) IncentiveCouncil() (common.Address, error) {
	return _EmissionsControllerStorage.Contract.IncentiveCouncil(&_EmissionsControllerStorage.CallOpts)
}

// IncentiveCouncil is a free data retrieval call binding the contract method 0xc44cb727.
//
// Solidity: function incentiveCouncil() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) IncentiveCouncil() (common.Address, error) {
	return _EmissionsControllerStorage.Contract.IncentiveCouncil(&_EmissionsControllerStorage.CallOpts)
}

// IsButtonPressable is a free data retrieval call binding the contract method 0xd8393150.
//
// Solidity: function isButtonPressable() view returns(bool)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) IsButtonPressable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "isButtonPressable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsButtonPressable is a free data retrieval call binding the contract method 0xd8393150.
//
// Solidity: function isButtonPressable() view returns(bool)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) IsButtonPressable() (bool, error) {
	return _EmissionsControllerStorage.Contract.IsButtonPressable(&_EmissionsControllerStorage.CallOpts)
}

// IsButtonPressable is a free data retrieval call binding the contract method 0xd8393150.
//
// Solidity: function isButtonPressable() view returns(bool)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) IsButtonPressable() (bool, error) {
	return _EmissionsControllerStorage.Contract.IsButtonPressable(&_EmissionsControllerStorage.CallOpts)
}

// LastTimeButtonPressable is a free data retrieval call binding the contract method 0xd44b1c9e.
//
// Solidity: function lastTimeButtonPressable() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) LastTimeButtonPressable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "lastTimeButtonPressable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeButtonPressable is a free data retrieval call binding the contract method 0xd44b1c9e.
//
// Solidity: function lastTimeButtonPressable() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) LastTimeButtonPressable() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.LastTimeButtonPressable(&_EmissionsControllerStorage.CallOpts)
}

// LastTimeButtonPressable is a free data retrieval call binding the contract method 0xd44b1c9e.
//
// Solidity: function lastTimeButtonPressable() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) LastTimeButtonPressable() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.LastTimeButtonPressable(&_EmissionsControllerStorage.CallOpts)
}

// NextTimeButtonPressable is a free data retrieval call binding the contract method 0xf769479f.
//
// Solidity: function nextTimeButtonPressable() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) NextTimeButtonPressable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "nextTimeButtonPressable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTimeButtonPressable is a free data retrieval call binding the contract method 0xf769479f.
//
// Solidity: function nextTimeButtonPressable() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) NextTimeButtonPressable() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.NextTimeButtonPressable(&_EmissionsControllerStorage.CallOpts)
}

// NextTimeButtonPressable is a free data retrieval call binding the contract method 0xf769479f.
//
// Solidity: function nextTimeButtonPressable() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) NextTimeButtonPressable() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.NextTimeButtonPressable(&_EmissionsControllerStorage.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) Paused(index uint8) (bool, error) {
	return _EmissionsControllerStorage.Contract.Paused(&_EmissionsControllerStorage.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) Paused(index uint8) (bool, error) {
	return _EmissionsControllerStorage.Contract.Paused(&_EmissionsControllerStorage.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) Paused0() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.Paused0(&_EmissionsControllerStorage.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) Paused0() (*big.Int, error) {
	return _EmissionsControllerStorage.Contract.Paused0(&_EmissionsControllerStorage.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) PauserRegistry() (common.Address, error) {
	return _EmissionsControllerStorage.Contract.PauserRegistry(&_EmissionsControllerStorage.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) PauserRegistry() (common.Address, error) {
	return _EmissionsControllerStorage.Contract.PauserRegistry(&_EmissionsControllerStorage.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint16)
func (_EmissionsControllerStorage *EmissionsControllerStorageCaller) TotalWeight(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EmissionsControllerStorage.contract.Call(opts, &out, "totalWeight")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint16)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) TotalWeight() (uint16, error) {
	return _EmissionsControllerStorage.Contract.TotalWeight(&_EmissionsControllerStorage.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint16)
func (_EmissionsControllerStorage *EmissionsControllerStorageCallerSession) TotalWeight() (uint16, error) {
	return _EmissionsControllerStorage.Contract.TotalWeight(&_EmissionsControllerStorage.CallOpts)
}

// AddDistribution is a paid mutator transaction binding the contract method 0xcd1e341b.
//
// Solidity: function addDistribution((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns(uint256 distributionId)
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactor) AddDistribution(opts *bind.TransactOpts, distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _EmissionsControllerStorage.contract.Transact(opts, "addDistribution", distribution)
}

// AddDistribution is a paid mutator transaction binding the contract method 0xcd1e341b.
//
// Solidity: function addDistribution((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns(uint256 distributionId)
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) AddDistribution(distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.AddDistribution(&_EmissionsControllerStorage.TransactOpts, distribution)
}

// AddDistribution is a paid mutator transaction binding the contract method 0xcd1e341b.
//
// Solidity: function addDistribution((uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns(uint256 distributionId)
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactorSession) AddDistribution(distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.AddDistribution(&_EmissionsControllerStorage.TransactOpts, distribution)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address incentiveCouncil, uint256 initialPausedStatus) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, incentiveCouncil common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _EmissionsControllerStorage.contract.Transact(opts, "initialize", initialOwner, incentiveCouncil, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address incentiveCouncil, uint256 initialPausedStatus) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) Initialize(initialOwner common.Address, incentiveCouncil common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.Initialize(&_EmissionsControllerStorage.TransactOpts, initialOwner, incentiveCouncil, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address incentiveCouncil, uint256 initialPausedStatus) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactorSession) Initialize(initialOwner common.Address, incentiveCouncil common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.Initialize(&_EmissionsControllerStorage.TransactOpts, initialOwner, incentiveCouncil, initialPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EmissionsControllerStorage.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.Pause(&_EmissionsControllerStorage.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.Pause(&_EmissionsControllerStorage.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmissionsControllerStorage.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) PauseAll() (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.PauseAll(&_EmissionsControllerStorage.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactorSession) PauseAll() (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.PauseAll(&_EmissionsControllerStorage.TransactOpts)
}

// PressButton is a paid mutator transaction binding the contract method 0x400efa85.
//
// Solidity: function pressButton(uint256 length) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactor) PressButton(opts *bind.TransactOpts, length *big.Int) (*types.Transaction, error) {
	return _EmissionsControllerStorage.contract.Transact(opts, "pressButton", length)
}

// PressButton is a paid mutator transaction binding the contract method 0x400efa85.
//
// Solidity: function pressButton(uint256 length) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) PressButton(length *big.Int) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.PressButton(&_EmissionsControllerStorage.TransactOpts, length)
}

// PressButton is a paid mutator transaction binding the contract method 0x400efa85.
//
// Solidity: function pressButton(uint256 length) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactorSession) PressButton(length *big.Int) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.PressButton(&_EmissionsControllerStorage.TransactOpts, length)
}

// SetIncentiveCouncil is a paid mutator transaction binding the contract method 0xc695acdb.
//
// Solidity: function setIncentiveCouncil(address incentiveCouncil) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactor) SetIncentiveCouncil(opts *bind.TransactOpts, incentiveCouncil common.Address) (*types.Transaction, error) {
	return _EmissionsControllerStorage.contract.Transact(opts, "setIncentiveCouncil", incentiveCouncil)
}

// SetIncentiveCouncil is a paid mutator transaction binding the contract method 0xc695acdb.
//
// Solidity: function setIncentiveCouncil(address incentiveCouncil) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) SetIncentiveCouncil(incentiveCouncil common.Address) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.SetIncentiveCouncil(&_EmissionsControllerStorage.TransactOpts, incentiveCouncil)
}

// SetIncentiveCouncil is a paid mutator transaction binding the contract method 0xc695acdb.
//
// Solidity: function setIncentiveCouncil(address incentiveCouncil) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactorSession) SetIncentiveCouncil(incentiveCouncil common.Address) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.SetIncentiveCouncil(&_EmissionsControllerStorage.TransactOpts, incentiveCouncil)
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactor) Sweep(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmissionsControllerStorage.contract.Transact(opts, "sweep")
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) Sweep() (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.Sweep(&_EmissionsControllerStorage.TransactOpts)
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactorSession) Sweep() (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.Sweep(&_EmissionsControllerStorage.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EmissionsControllerStorage.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.Unpause(&_EmissionsControllerStorage.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.Unpause(&_EmissionsControllerStorage.TransactOpts, newPausedStatus)
}

// UpdateDistribution is a paid mutator transaction binding the contract method 0x44a32028.
//
// Solidity: function updateDistribution(uint256 distributionId, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactor) UpdateDistribution(opts *bind.TransactOpts, distributionId *big.Int, distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _EmissionsControllerStorage.contract.Transact(opts, "updateDistribution", distributionId, distribution)
}

// UpdateDistribution is a paid mutator transaction binding the contract method 0x44a32028.
//
// Solidity: function updateDistribution(uint256 distributionId, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageSession) UpdateDistribution(distributionId *big.Int, distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.UpdateDistribution(&_EmissionsControllerStorage.TransactOpts, distributionId, distribution)
}

// UpdateDistribution is a paid mutator transaction binding the contract method 0x44a32028.
//
// Solidity: function updateDistribution(uint256 distributionId, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution) returns()
func (_EmissionsControllerStorage *EmissionsControllerStorageTransactorSession) UpdateDistribution(distributionId *big.Int, distribution IEmissionsControllerTypesDistribution) (*types.Transaction, error) {
	return _EmissionsControllerStorage.Contract.UpdateDistribution(&_EmissionsControllerStorage.TransactOpts, distributionId, distribution)
}

// EmissionsControllerStorageDistributionAddedIterator is returned from FilterDistributionAdded and is used to iterate over the raw logs and unpacked data for DistributionAdded events raised by the EmissionsControllerStorage contract.
type EmissionsControllerStorageDistributionAddedIterator struct {
	Event *EmissionsControllerStorageDistributionAdded // Event containing the contract specifics and raw log

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
func (it *EmissionsControllerStorageDistributionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmissionsControllerStorageDistributionAdded)
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
		it.Event = new(EmissionsControllerStorageDistributionAdded)
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
func (it *EmissionsControllerStorageDistributionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmissionsControllerStorageDistributionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmissionsControllerStorageDistributionAdded represents a DistributionAdded event raised by the EmissionsControllerStorage contract.
type EmissionsControllerStorageDistributionAdded struct {
	DistributionId *big.Int
	Epoch          *big.Int
	Distribution   IEmissionsControllerTypesDistribution
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDistributionAdded is a free log retrieval operation binding the contract event 0x006f7ba35643ecf5852cfe01b66220b1fe04a4cd4866923d5f3e66c7fcb390ef.
//
// Solidity: event DistributionAdded(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) FilterDistributionAdded(opts *bind.FilterOpts, distributionId []*big.Int, epoch []*big.Int) (*EmissionsControllerStorageDistributionAddedIterator, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _EmissionsControllerStorage.contract.FilterLogs(opts, "DistributionAdded", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &EmissionsControllerStorageDistributionAddedIterator{contract: _EmissionsControllerStorage.contract, event: "DistributionAdded", logs: logs, sub: sub}, nil
}

// WatchDistributionAdded is a free log subscription operation binding the contract event 0x006f7ba35643ecf5852cfe01b66220b1fe04a4cd4866923d5f3e66c7fcb390ef.
//
// Solidity: event DistributionAdded(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) WatchDistributionAdded(opts *bind.WatchOpts, sink chan<- *EmissionsControllerStorageDistributionAdded, distributionId []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _EmissionsControllerStorage.contract.WatchLogs(opts, "DistributionAdded", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmissionsControllerStorageDistributionAdded)
				if err := _EmissionsControllerStorage.contract.UnpackLog(event, "DistributionAdded", log); err != nil {
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
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) ParseDistributionAdded(log types.Log) (*EmissionsControllerStorageDistributionAdded, error) {
	event := new(EmissionsControllerStorageDistributionAdded)
	if err := _EmissionsControllerStorage.contract.UnpackLog(event, "DistributionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmissionsControllerStorageDistributionProcessedIterator is returned from FilterDistributionProcessed and is used to iterate over the raw logs and unpacked data for DistributionProcessed events raised by the EmissionsControllerStorage contract.
type EmissionsControllerStorageDistributionProcessedIterator struct {
	Event *EmissionsControllerStorageDistributionProcessed // Event containing the contract specifics and raw log

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
func (it *EmissionsControllerStorageDistributionProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmissionsControllerStorageDistributionProcessed)
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
		it.Event = new(EmissionsControllerStorageDistributionProcessed)
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
func (it *EmissionsControllerStorageDistributionProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmissionsControllerStorageDistributionProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmissionsControllerStorageDistributionProcessed represents a DistributionProcessed event raised by the EmissionsControllerStorage contract.
type EmissionsControllerStorageDistributionProcessed struct {
	DistributionId *big.Int
	Epoch          *big.Int
	Distribution   IEmissionsControllerTypesDistribution
	Success        bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDistributionProcessed is a free log retrieval operation binding the contract event 0xba5d66336bc9a4f8459242151a4da4d5020ac581243d98403bb55f7f348e071b.
//
// Solidity: event DistributionProcessed(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution, bool success)
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) FilterDistributionProcessed(opts *bind.FilterOpts, distributionId []*big.Int, epoch []*big.Int) (*EmissionsControllerStorageDistributionProcessedIterator, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _EmissionsControllerStorage.contract.FilterLogs(opts, "DistributionProcessed", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &EmissionsControllerStorageDistributionProcessedIterator{contract: _EmissionsControllerStorage.contract, event: "DistributionProcessed", logs: logs, sub: sub}, nil
}

// WatchDistributionProcessed is a free log subscription operation binding the contract event 0xba5d66336bc9a4f8459242151a4da4d5020ac581243d98403bb55f7f348e071b.
//
// Solidity: event DistributionProcessed(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution, bool success)
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) WatchDistributionProcessed(opts *bind.WatchOpts, sink chan<- *EmissionsControllerStorageDistributionProcessed, distributionId []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _EmissionsControllerStorage.contract.WatchLogs(opts, "DistributionProcessed", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmissionsControllerStorageDistributionProcessed)
				if err := _EmissionsControllerStorage.contract.UnpackLog(event, "DistributionProcessed", log); err != nil {
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
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) ParseDistributionProcessed(log types.Log) (*EmissionsControllerStorageDistributionProcessed, error) {
	event := new(EmissionsControllerStorageDistributionProcessed)
	if err := _EmissionsControllerStorage.contract.UnpackLog(event, "DistributionProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmissionsControllerStorageDistributionUpdatedIterator is returned from FilterDistributionUpdated and is used to iterate over the raw logs and unpacked data for DistributionUpdated events raised by the EmissionsControllerStorage contract.
type EmissionsControllerStorageDistributionUpdatedIterator struct {
	Event *EmissionsControllerStorageDistributionUpdated // Event containing the contract specifics and raw log

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
func (it *EmissionsControllerStorageDistributionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmissionsControllerStorageDistributionUpdated)
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
		it.Event = new(EmissionsControllerStorageDistributionUpdated)
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
func (it *EmissionsControllerStorageDistributionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmissionsControllerStorageDistributionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmissionsControllerStorageDistributionUpdated represents a DistributionUpdated event raised by the EmissionsControllerStorage contract.
type EmissionsControllerStorageDistributionUpdated struct {
	DistributionId *big.Int
	Epoch          *big.Int
	Distribution   IEmissionsControllerTypesDistribution
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDistributionUpdated is a free log retrieval operation binding the contract event 0x548fb50d6be978df2bacbf48c6840e4a4743672408921282117f3f00555b2b4c.
//
// Solidity: event DistributionUpdated(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) FilterDistributionUpdated(opts *bind.FilterOpts, distributionId []*big.Int, epoch []*big.Int) (*EmissionsControllerStorageDistributionUpdatedIterator, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _EmissionsControllerStorage.contract.FilterLogs(opts, "DistributionUpdated", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &EmissionsControllerStorageDistributionUpdatedIterator{contract: _EmissionsControllerStorage.contract, event: "DistributionUpdated", logs: logs, sub: sub}, nil
}

// WatchDistributionUpdated is a free log subscription operation binding the contract event 0x548fb50d6be978df2bacbf48c6840e4a4743672408921282117f3f00555b2b4c.
//
// Solidity: event DistributionUpdated(uint256 indexed distributionId, uint256 indexed epoch, (uint64,uint64,uint64,uint8,(address,uint32),(address,uint96)[][]) distribution)
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) WatchDistributionUpdated(opts *bind.WatchOpts, sink chan<- *EmissionsControllerStorageDistributionUpdated, distributionId []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var distributionIdRule []interface{}
	for _, distributionIdItem := range distributionId {
		distributionIdRule = append(distributionIdRule, distributionIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _EmissionsControllerStorage.contract.WatchLogs(opts, "DistributionUpdated", distributionIdRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmissionsControllerStorageDistributionUpdated)
				if err := _EmissionsControllerStorage.contract.UnpackLog(event, "DistributionUpdated", log); err != nil {
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
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) ParseDistributionUpdated(log types.Log) (*EmissionsControllerStorageDistributionUpdated, error) {
	event := new(EmissionsControllerStorageDistributionUpdated)
	if err := _EmissionsControllerStorage.contract.UnpackLog(event, "DistributionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmissionsControllerStorageIncentiveCouncilUpdatedIterator is returned from FilterIncentiveCouncilUpdated and is used to iterate over the raw logs and unpacked data for IncentiveCouncilUpdated events raised by the EmissionsControllerStorage contract.
type EmissionsControllerStorageIncentiveCouncilUpdatedIterator struct {
	Event *EmissionsControllerStorageIncentiveCouncilUpdated // Event containing the contract specifics and raw log

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
func (it *EmissionsControllerStorageIncentiveCouncilUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmissionsControllerStorageIncentiveCouncilUpdated)
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
		it.Event = new(EmissionsControllerStorageIncentiveCouncilUpdated)
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
func (it *EmissionsControllerStorageIncentiveCouncilUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmissionsControllerStorageIncentiveCouncilUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmissionsControllerStorageIncentiveCouncilUpdated represents a IncentiveCouncilUpdated event raised by the EmissionsControllerStorage contract.
type EmissionsControllerStorageIncentiveCouncilUpdated struct {
	IncentiveCouncil common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterIncentiveCouncilUpdated is a free log retrieval operation binding the contract event 0x8befac6896b786e67b23cefc473bfabd36e7fc013125c883dfeec8e3a9636216.
//
// Solidity: event IncentiveCouncilUpdated(address indexed incentiveCouncil)
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) FilterIncentiveCouncilUpdated(opts *bind.FilterOpts, incentiveCouncil []common.Address) (*EmissionsControllerStorageIncentiveCouncilUpdatedIterator, error) {

	var incentiveCouncilRule []interface{}
	for _, incentiveCouncilItem := range incentiveCouncil {
		incentiveCouncilRule = append(incentiveCouncilRule, incentiveCouncilItem)
	}

	logs, sub, err := _EmissionsControllerStorage.contract.FilterLogs(opts, "IncentiveCouncilUpdated", incentiveCouncilRule)
	if err != nil {
		return nil, err
	}
	return &EmissionsControllerStorageIncentiveCouncilUpdatedIterator{contract: _EmissionsControllerStorage.contract, event: "IncentiveCouncilUpdated", logs: logs, sub: sub}, nil
}

// WatchIncentiveCouncilUpdated is a free log subscription operation binding the contract event 0x8befac6896b786e67b23cefc473bfabd36e7fc013125c883dfeec8e3a9636216.
//
// Solidity: event IncentiveCouncilUpdated(address indexed incentiveCouncil)
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) WatchIncentiveCouncilUpdated(opts *bind.WatchOpts, sink chan<- *EmissionsControllerStorageIncentiveCouncilUpdated, incentiveCouncil []common.Address) (event.Subscription, error) {

	var incentiveCouncilRule []interface{}
	for _, incentiveCouncilItem := range incentiveCouncil {
		incentiveCouncilRule = append(incentiveCouncilRule, incentiveCouncilItem)
	}

	logs, sub, err := _EmissionsControllerStorage.contract.WatchLogs(opts, "IncentiveCouncilUpdated", incentiveCouncilRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmissionsControllerStorageIncentiveCouncilUpdated)
				if err := _EmissionsControllerStorage.contract.UnpackLog(event, "IncentiveCouncilUpdated", log); err != nil {
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
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) ParseIncentiveCouncilUpdated(log types.Log) (*EmissionsControllerStorageIncentiveCouncilUpdated, error) {
	event := new(EmissionsControllerStorageIncentiveCouncilUpdated)
	if err := _EmissionsControllerStorage.contract.UnpackLog(event, "IncentiveCouncilUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmissionsControllerStoragePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EmissionsControllerStorage contract.
type EmissionsControllerStoragePausedIterator struct {
	Event *EmissionsControllerStoragePaused // Event containing the contract specifics and raw log

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
func (it *EmissionsControllerStoragePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmissionsControllerStoragePaused)
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
		it.Event = new(EmissionsControllerStoragePaused)
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
func (it *EmissionsControllerStoragePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmissionsControllerStoragePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmissionsControllerStoragePaused represents a Paused event raised by the EmissionsControllerStorage contract.
type EmissionsControllerStoragePaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*EmissionsControllerStoragePausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EmissionsControllerStorage.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &EmissionsControllerStoragePausedIterator{contract: _EmissionsControllerStorage.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EmissionsControllerStoragePaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EmissionsControllerStorage.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmissionsControllerStoragePaused)
				if err := _EmissionsControllerStorage.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) ParsePaused(log types.Log) (*EmissionsControllerStoragePaused, error) {
	event := new(EmissionsControllerStoragePaused)
	if err := _EmissionsControllerStorage.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmissionsControllerStorageSweptIterator is returned from FilterSwept and is used to iterate over the raw logs and unpacked data for Swept events raised by the EmissionsControllerStorage contract.
type EmissionsControllerStorageSweptIterator struct {
	Event *EmissionsControllerStorageSwept // Event containing the contract specifics and raw log

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
func (it *EmissionsControllerStorageSweptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmissionsControllerStorageSwept)
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
		it.Event = new(EmissionsControllerStorageSwept)
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
func (it *EmissionsControllerStorageSweptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmissionsControllerStorageSweptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmissionsControllerStorageSwept represents a Swept event raised by the EmissionsControllerStorage contract.
type EmissionsControllerStorageSwept struct {
	IncentiveCouncil common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSwept is a free log retrieval operation binding the contract event 0xc36b5179cb9c303b200074996eab2b3473eac370fdd7eba3bec636fe35109696.
//
// Solidity: event Swept(address indexed incentiveCouncil, uint256 amount)
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) FilterSwept(opts *bind.FilterOpts, incentiveCouncil []common.Address) (*EmissionsControllerStorageSweptIterator, error) {

	var incentiveCouncilRule []interface{}
	for _, incentiveCouncilItem := range incentiveCouncil {
		incentiveCouncilRule = append(incentiveCouncilRule, incentiveCouncilItem)
	}

	logs, sub, err := _EmissionsControllerStorage.contract.FilterLogs(opts, "Swept", incentiveCouncilRule)
	if err != nil {
		return nil, err
	}
	return &EmissionsControllerStorageSweptIterator{contract: _EmissionsControllerStorage.contract, event: "Swept", logs: logs, sub: sub}, nil
}

// WatchSwept is a free log subscription operation binding the contract event 0xc36b5179cb9c303b200074996eab2b3473eac370fdd7eba3bec636fe35109696.
//
// Solidity: event Swept(address indexed incentiveCouncil, uint256 amount)
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) WatchSwept(opts *bind.WatchOpts, sink chan<- *EmissionsControllerStorageSwept, incentiveCouncil []common.Address) (event.Subscription, error) {

	var incentiveCouncilRule []interface{}
	for _, incentiveCouncilItem := range incentiveCouncil {
		incentiveCouncilRule = append(incentiveCouncilRule, incentiveCouncilItem)
	}

	logs, sub, err := _EmissionsControllerStorage.contract.WatchLogs(opts, "Swept", incentiveCouncilRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmissionsControllerStorageSwept)
				if err := _EmissionsControllerStorage.contract.UnpackLog(event, "Swept", log); err != nil {
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
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) ParseSwept(log types.Log) (*EmissionsControllerStorageSwept, error) {
	event := new(EmissionsControllerStorageSwept)
	if err := _EmissionsControllerStorage.contract.UnpackLog(event, "Swept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmissionsControllerStorageUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EmissionsControllerStorage contract.
type EmissionsControllerStorageUnpausedIterator struct {
	Event *EmissionsControllerStorageUnpaused // Event containing the contract specifics and raw log

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
func (it *EmissionsControllerStorageUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmissionsControllerStorageUnpaused)
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
		it.Event = new(EmissionsControllerStorageUnpaused)
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
func (it *EmissionsControllerStorageUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmissionsControllerStorageUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmissionsControllerStorageUnpaused represents a Unpaused event raised by the EmissionsControllerStorage contract.
type EmissionsControllerStorageUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*EmissionsControllerStorageUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EmissionsControllerStorage.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &EmissionsControllerStorageUnpausedIterator{contract: _EmissionsControllerStorage.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EmissionsControllerStorageUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EmissionsControllerStorage.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmissionsControllerStorageUnpaused)
				if err := _EmissionsControllerStorage.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_EmissionsControllerStorage *EmissionsControllerStorageFilterer) ParseUnpaused(log types.Log) (*EmissionsControllerStorageUnpaused, error) {
	event := new(EmissionsControllerStorageUnpaused)
	if err := _EmissionsControllerStorage.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
