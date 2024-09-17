// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StakeRootCompendiumStorage

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

// IStakeRootCompendiumOperatorLeaf is an auto generated low-level Go binding around an user-defined struct.
type IStakeRootCompendiumOperatorLeaf struct {
	DelegatedStake *big.Int
	SlashableStake *big.Int
	ExtraData      [32]byte
}

// IStakeRootCompendiumProof is an auto generated low-level Go binding around an user-defined struct.
type IStakeRootCompendiumProof struct {
	X uint32
}

// IStakeRootCompendiumStakeRootSubmission is an auto generated low-level Go binding around an user-defined struct.
type IStakeRootCompendiumStakeRootSubmission struct {
	StakeRoot            [32]byte
	CalculationTimestamp uint32
	Confirmed            bool
}

// IStakeRootCompendiumStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IStakeRootCompendiumStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs           common.Address
	OperatorSetId uint32
}

// StakeRootCompendiumStorageMetaData contains all meta data concerning the StakeRootCompendiumStorage contract.
var StakeRootCompendiumStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"MAX_TOTAL_CHARGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_BALANCE_THRESHOLD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_PROOFS_PREPAID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REMOVED_INDEX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addOrModifyStrategiesAndMultipliers\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRootCompendium.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canWithdrawDepositBalance\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chargePerOperatorSet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chargePerStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmStakeRoot\",\"inputs\":[{\"name\":\"calculationTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stakeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositInfos\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"lastDemandIncreaseTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"totalChargePerOperatorSetLastPaid\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"totalChargePerStrategyLastPaid\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositBalance\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumOperatorSets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumStakeRootSubmissions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetLeaves\",\"inputs\":[{\"name\":\"operatorSetIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startOperatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRootCompendium.OperatorLeaf[]\",\"components\":[{\"name\":\"delegatedStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slashableStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetRoot\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRootCompendium.OperatorLeaf[]\",\"components\":[{\"name\":\"delegatedStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slashableStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakeRoot\",\"inputs\":[{\"name\":\"operatorSetsInStakeTree\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorSetRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakeRootSubmission\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIStakeRootCompendium.StakeRootSubmission\",\"components\":[{\"name\":\"stakeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"calculationTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"confirmed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"delegatedStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slashableStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"imageId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minDepositBalance\",\"inputs\":[{\"name\":\"numStrategies\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proofIntervalSeconds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeOperatorSetsFromStakeTree\",\"inputs\":[{\"name\":\"operatorSetsToRemove\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesAndMultipliers\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rootConfirmer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setOperatorExtraData\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetExtraData\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRootSubmissions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"stakeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"calculationTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"confirmed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalChargeLastUpdatedTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalChargePerOperatorSetLastUpdate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalChargePerStrategyLastUpdate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyStakeRoot\",\"inputs\":[{\"name\":\"_calculationTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stakeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_chargeRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexChargePerProof\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_proof\",\"type\":\"tuple\",\"internalType\":\"structIStakeRootCompendium.Proof\",\"components\":[{\"name\":\"x\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ImageIdSet\",\"inputs\":[{\"name\":\"newImageId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SnarkProofVerified\",\"inputs\":[{\"name\":\"journal\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"seal\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifierSet\",\"inputs\":[{\"name\":\"newVerifier\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// StakeRootCompendiumStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeRootCompendiumStorageMetaData.ABI instead.
var StakeRootCompendiumStorageABI = StakeRootCompendiumStorageMetaData.ABI

// StakeRootCompendiumStorage is an auto generated Go binding around an Ethereum contract.
type StakeRootCompendiumStorage struct {
	StakeRootCompendiumStorageCaller     // Read-only binding to the contract
	StakeRootCompendiumStorageTransactor // Write-only binding to the contract
	StakeRootCompendiumStorageFilterer   // Log filterer for contract events
}

// StakeRootCompendiumStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeRootCompendiumStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeRootCompendiumStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeRootCompendiumStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeRootCompendiumStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeRootCompendiumStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeRootCompendiumStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeRootCompendiumStorageSession struct {
	Contract     *StakeRootCompendiumStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StakeRootCompendiumStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeRootCompendiumStorageCallerSession struct {
	Contract *StakeRootCompendiumStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// StakeRootCompendiumStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeRootCompendiumStorageTransactorSession struct {
	Contract     *StakeRootCompendiumStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// StakeRootCompendiumStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeRootCompendiumStorageRaw struct {
	Contract *StakeRootCompendiumStorage // Generic contract binding to access the raw methods on
}

// StakeRootCompendiumStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeRootCompendiumStorageCallerRaw struct {
	Contract *StakeRootCompendiumStorageCaller // Generic read-only contract binding to access the raw methods on
}

// StakeRootCompendiumStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeRootCompendiumStorageTransactorRaw struct {
	Contract *StakeRootCompendiumStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeRootCompendiumStorage creates a new instance of StakeRootCompendiumStorage, bound to a specific deployed contract.
func NewStakeRootCompendiumStorage(address common.Address, backend bind.ContractBackend) (*StakeRootCompendiumStorage, error) {
	contract, err := bindStakeRootCompendiumStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeRootCompendiumStorage{StakeRootCompendiumStorageCaller: StakeRootCompendiumStorageCaller{contract: contract}, StakeRootCompendiumStorageTransactor: StakeRootCompendiumStorageTransactor{contract: contract}, StakeRootCompendiumStorageFilterer: StakeRootCompendiumStorageFilterer{contract: contract}}, nil
}

// NewStakeRootCompendiumStorageCaller creates a new read-only instance of StakeRootCompendiumStorage, bound to a specific deployed contract.
func NewStakeRootCompendiumStorageCaller(address common.Address, caller bind.ContractCaller) (*StakeRootCompendiumStorageCaller, error) {
	contract, err := bindStakeRootCompendiumStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeRootCompendiumStorageCaller{contract: contract}, nil
}

// NewStakeRootCompendiumStorageTransactor creates a new write-only instance of StakeRootCompendiumStorage, bound to a specific deployed contract.
func NewStakeRootCompendiumStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeRootCompendiumStorageTransactor, error) {
	contract, err := bindStakeRootCompendiumStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeRootCompendiumStorageTransactor{contract: contract}, nil
}

// NewStakeRootCompendiumStorageFilterer creates a new log filterer instance of StakeRootCompendiumStorage, bound to a specific deployed contract.
func NewStakeRootCompendiumStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeRootCompendiumStorageFilterer, error) {
	contract, err := bindStakeRootCompendiumStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeRootCompendiumStorageFilterer{contract: contract}, nil
}

// bindStakeRootCompendiumStorage binds a generic wrapper to an already deployed contract.
func bindStakeRootCompendiumStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakeRootCompendiumStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeRootCompendiumStorage.Contract.StakeRootCompendiumStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.StakeRootCompendiumStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.StakeRootCompendiumStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeRootCompendiumStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.contract.Transact(opts, method, params...)
}

// MAXTOTALCHARGE is a free data retrieval call binding the contract method 0x3264c71b.
//
// Solidity: function MAX_TOTAL_CHARGE() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) MAXTOTALCHARGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "MAX_TOTAL_CHARGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOTALCHARGE is a free data retrieval call binding the contract method 0x3264c71b.
//
// Solidity: function MAX_TOTAL_CHARGE() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) MAXTOTALCHARGE() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.MAXTOTALCHARGE(&_StakeRootCompendiumStorage.CallOpts)
}

// MAXTOTALCHARGE is a free data retrieval call binding the contract method 0x3264c71b.
//
// Solidity: function MAX_TOTAL_CHARGE() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) MAXTOTALCHARGE() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.MAXTOTALCHARGE(&_StakeRootCompendiumStorage.CallOpts)
}

// MINBALANCETHRESHOLD is a free data retrieval call binding the contract method 0xc442daee.
//
// Solidity: function MIN_BALANCE_THRESHOLD() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) MINBALANCETHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "MIN_BALANCE_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBALANCETHRESHOLD is a free data retrieval call binding the contract method 0xc442daee.
//
// Solidity: function MIN_BALANCE_THRESHOLD() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) MINBALANCETHRESHOLD() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.MINBALANCETHRESHOLD(&_StakeRootCompendiumStorage.CallOpts)
}

// MINBALANCETHRESHOLD is a free data retrieval call binding the contract method 0xc442daee.
//
// Solidity: function MIN_BALANCE_THRESHOLD() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) MINBALANCETHRESHOLD() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.MINBALANCETHRESHOLD(&_StakeRootCompendiumStorage.CallOpts)
}

// MINPROOFSPREPAID is a free data retrieval call binding the contract method 0x54a21db2.
//
// Solidity: function MIN_PROOFS_PREPAID() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) MINPROOFSPREPAID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "MIN_PROOFS_PREPAID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINPROOFSPREPAID is a free data retrieval call binding the contract method 0x54a21db2.
//
// Solidity: function MIN_PROOFS_PREPAID() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) MINPROOFSPREPAID() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.MINPROOFSPREPAID(&_StakeRootCompendiumStorage.CallOpts)
}

// MINPROOFSPREPAID is a free data retrieval call binding the contract method 0x54a21db2.
//
// Solidity: function MIN_PROOFS_PREPAID() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) MINPROOFSPREPAID() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.MINPROOFSPREPAID(&_StakeRootCompendiumStorage.CallOpts)
}

// REMOVEDINDEX is a free data retrieval call binding the contract method 0x0d6098d6.
//
// Solidity: function REMOVED_INDEX() view returns(uint32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) REMOVEDINDEX(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "REMOVED_INDEX")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// REMOVEDINDEX is a free data retrieval call binding the contract method 0x0d6098d6.
//
// Solidity: function REMOVED_INDEX() view returns(uint32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) REMOVEDINDEX() (uint32, error) {
	return _StakeRootCompendiumStorage.Contract.REMOVEDINDEX(&_StakeRootCompendiumStorage.CallOpts)
}

// REMOVEDINDEX is a free data retrieval call binding the contract method 0x0d6098d6.
//
// Solidity: function REMOVED_INDEX() view returns(uint32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) REMOVEDINDEX() (uint32, error) {
	return _StakeRootCompendiumStorage.Contract.REMOVEDINDEX(&_StakeRootCompendiumStorage.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) AllocationManager() (common.Address, error) {
	return _StakeRootCompendiumStorage.Contract.AllocationManager(&_StakeRootCompendiumStorage.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) AllocationManager() (common.Address, error) {
	return _StakeRootCompendiumStorage.Contract.AllocationManager(&_StakeRootCompendiumStorage.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) AvsDirectory() (common.Address, error) {
	return _StakeRootCompendiumStorage.Contract.AvsDirectory(&_StakeRootCompendiumStorage.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) AvsDirectory() (common.Address, error) {
	return _StakeRootCompendiumStorage.Contract.AvsDirectory(&_StakeRootCompendiumStorage.CallOpts)
}

// CanWithdrawDepositBalance is a free data retrieval call binding the contract method 0x71616631.
//
// Solidity: function canWithdrawDepositBalance((address,uint32) operatorSet) view returns(bool)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) CanWithdrawDepositBalance(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "canWithdrawDepositBalance", operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanWithdrawDepositBalance is a free data retrieval call binding the contract method 0x71616631.
//
// Solidity: function canWithdrawDepositBalance((address,uint32) operatorSet) view returns(bool)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) CanWithdrawDepositBalance(operatorSet OperatorSet) (bool, error) {
	return _StakeRootCompendiumStorage.Contract.CanWithdrawDepositBalance(&_StakeRootCompendiumStorage.CallOpts, operatorSet)
}

// CanWithdrawDepositBalance is a free data retrieval call binding the contract method 0x71616631.
//
// Solidity: function canWithdrawDepositBalance((address,uint32) operatorSet) view returns(bool)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) CanWithdrawDepositBalance(operatorSet OperatorSet) (bool, error) {
	return _StakeRootCompendiumStorage.Contract.CanWithdrawDepositBalance(&_StakeRootCompendiumStorage.CallOpts, operatorSet)
}

// ChargePerOperatorSet is a free data retrieval call binding the contract method 0x2e930d26.
//
// Solidity: function chargePerOperatorSet() view returns(uint96)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) ChargePerOperatorSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "chargePerOperatorSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChargePerOperatorSet is a free data retrieval call binding the contract method 0x2e930d26.
//
// Solidity: function chargePerOperatorSet() view returns(uint96)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) ChargePerOperatorSet() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.ChargePerOperatorSet(&_StakeRootCompendiumStorage.CallOpts)
}

// ChargePerOperatorSet is a free data retrieval call binding the contract method 0x2e930d26.
//
// Solidity: function chargePerOperatorSet() view returns(uint96)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) ChargePerOperatorSet() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.ChargePerOperatorSet(&_StakeRootCompendiumStorage.CallOpts)
}

// ChargePerStrategy is a free data retrieval call binding the contract method 0x08c76694.
//
// Solidity: function chargePerStrategy() view returns(uint96)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) ChargePerStrategy(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "chargePerStrategy")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChargePerStrategy is a free data retrieval call binding the contract method 0x08c76694.
//
// Solidity: function chargePerStrategy() view returns(uint96)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) ChargePerStrategy() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.ChargePerStrategy(&_StakeRootCompendiumStorage.CallOpts)
}

// ChargePerStrategy is a free data retrieval call binding the contract method 0x08c76694.
//
// Solidity: function chargePerStrategy() view returns(uint96)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) ChargePerStrategy() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.ChargePerStrategy(&_StakeRootCompendiumStorage.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) DelegationManager() (common.Address, error) {
	return _StakeRootCompendiumStorage.Contract.DelegationManager(&_StakeRootCompendiumStorage.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) DelegationManager() (common.Address, error) {
	return _StakeRootCompendiumStorage.Contract.DelegationManager(&_StakeRootCompendiumStorage.CallOpts)
}

// DepositInfos is a free data retrieval call binding the contract method 0xbb1bb891.
//
// Solidity: function depositInfos(address , uint32 ) view returns(uint96 balance, uint32 lastDemandIncreaseTimestamp, uint96 totalChargePerOperatorSetLastPaid, uint96 totalChargePerStrategyLastPaid)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) DepositInfos(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	Balance                           *big.Int
	LastDemandIncreaseTimestamp       uint32
	TotalChargePerOperatorSetLastPaid *big.Int
	TotalChargePerStrategyLastPaid    *big.Int
}, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "depositInfos", arg0, arg1)

	outstruct := new(struct {
		Balance                           *big.Int
		LastDemandIncreaseTimestamp       uint32
		TotalChargePerOperatorSetLastPaid *big.Int
		TotalChargePerStrategyLastPaid    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastDemandIncreaseTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.TotalChargePerOperatorSetLastPaid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalChargePerStrategyLastPaid = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DepositInfos is a free data retrieval call binding the contract method 0xbb1bb891.
//
// Solidity: function depositInfos(address , uint32 ) view returns(uint96 balance, uint32 lastDemandIncreaseTimestamp, uint96 totalChargePerOperatorSetLastPaid, uint96 totalChargePerStrategyLastPaid)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) DepositInfos(arg0 common.Address, arg1 uint32) (struct {
	Balance                           *big.Int
	LastDemandIncreaseTimestamp       uint32
	TotalChargePerOperatorSetLastPaid *big.Int
	TotalChargePerStrategyLastPaid    *big.Int
}, error) {
	return _StakeRootCompendiumStorage.Contract.DepositInfos(&_StakeRootCompendiumStorage.CallOpts, arg0, arg1)
}

// DepositInfos is a free data retrieval call binding the contract method 0xbb1bb891.
//
// Solidity: function depositInfos(address , uint32 ) view returns(uint96 balance, uint32 lastDemandIncreaseTimestamp, uint96 totalChargePerOperatorSetLastPaid, uint96 totalChargePerStrategyLastPaid)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) DepositInfos(arg0 common.Address, arg1 uint32) (struct {
	Balance                           *big.Int
	LastDemandIncreaseTimestamp       uint32
	TotalChargePerOperatorSetLastPaid *big.Int
	TotalChargePerStrategyLastPaid    *big.Int
}, error) {
	return _StakeRootCompendiumStorage.Contract.DepositInfos(&_StakeRootCompendiumStorage.CallOpts, arg0, arg1)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0x2156171a.
//
// Solidity: function getDepositBalance((address,uint32) operatorSet) view returns(uint256 balance)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) GetDepositBalance(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "getDepositBalance", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositBalance is a free data retrieval call binding the contract method 0x2156171a.
//
// Solidity: function getDepositBalance((address,uint32) operatorSet) view returns(uint256 balance)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) GetDepositBalance(operatorSet OperatorSet) (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.GetDepositBalance(&_StakeRootCompendiumStorage.CallOpts, operatorSet)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0x2156171a.
//
// Solidity: function getDepositBalance((address,uint32) operatorSet) view returns(uint256 balance)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) GetDepositBalance(operatorSet OperatorSet) (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.GetDepositBalance(&_StakeRootCompendiumStorage.CallOpts, operatorSet)
}

// GetNumOperatorSets is a free data retrieval call binding the contract method 0x11bea0d1.
//
// Solidity: function getNumOperatorSets() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) GetNumOperatorSets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "getNumOperatorSets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumOperatorSets is a free data retrieval call binding the contract method 0x11bea0d1.
//
// Solidity: function getNumOperatorSets() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) GetNumOperatorSets() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.GetNumOperatorSets(&_StakeRootCompendiumStorage.CallOpts)
}

// GetNumOperatorSets is a free data retrieval call binding the contract method 0x11bea0d1.
//
// Solidity: function getNumOperatorSets() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) GetNumOperatorSets() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.GetNumOperatorSets(&_StakeRootCompendiumStorage.CallOpts)
}

// GetNumStakeRootSubmissions is a free data retrieval call binding the contract method 0xd985c7b3.
//
// Solidity: function getNumStakeRootSubmissions() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) GetNumStakeRootSubmissions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "getNumStakeRootSubmissions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumStakeRootSubmissions is a free data retrieval call binding the contract method 0xd985c7b3.
//
// Solidity: function getNumStakeRootSubmissions() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) GetNumStakeRootSubmissions() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.GetNumStakeRootSubmissions(&_StakeRootCompendiumStorage.CallOpts)
}

// GetNumStakeRootSubmissions is a free data retrieval call binding the contract method 0xd985c7b3.
//
// Solidity: function getNumStakeRootSubmissions() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) GetNumStakeRootSubmissions() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.GetNumStakeRootSubmissions(&_StakeRootCompendiumStorage.CallOpts)
}

// GetOperatorSetLeaves is a free data retrieval call binding the contract method 0x732d7398.
//
// Solidity: function getOperatorSetLeaves(uint256 operatorSetIndex, uint256 startOperatorIndex, uint256 numOperators) view returns((address,uint32), address[], (uint256,uint256,bytes32)[])
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) GetOperatorSetLeaves(opts *bind.CallOpts, operatorSetIndex *big.Int, startOperatorIndex *big.Int, numOperators *big.Int) (OperatorSet, []common.Address, []IStakeRootCompendiumOperatorLeaf, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "getOperatorSetLeaves", operatorSetIndex, startOperatorIndex, numOperators)

	if err != nil {
		return *new(OperatorSet), *new([]common.Address), *new([]IStakeRootCompendiumOperatorLeaf), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorSet)).(*OperatorSet)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	out2 := *abi.ConvertType(out[2], new([]IStakeRootCompendiumOperatorLeaf)).(*[]IStakeRootCompendiumOperatorLeaf)

	return out0, out1, out2, err

}

// GetOperatorSetLeaves is a free data retrieval call binding the contract method 0x732d7398.
//
// Solidity: function getOperatorSetLeaves(uint256 operatorSetIndex, uint256 startOperatorIndex, uint256 numOperators) view returns((address,uint32), address[], (uint256,uint256,bytes32)[])
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) GetOperatorSetLeaves(operatorSetIndex *big.Int, startOperatorIndex *big.Int, numOperators *big.Int) (OperatorSet, []common.Address, []IStakeRootCompendiumOperatorLeaf, error) {
	return _StakeRootCompendiumStorage.Contract.GetOperatorSetLeaves(&_StakeRootCompendiumStorage.CallOpts, operatorSetIndex, startOperatorIndex, numOperators)
}

// GetOperatorSetLeaves is a free data retrieval call binding the contract method 0x732d7398.
//
// Solidity: function getOperatorSetLeaves(uint256 operatorSetIndex, uint256 startOperatorIndex, uint256 numOperators) view returns((address,uint32), address[], (uint256,uint256,bytes32)[])
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) GetOperatorSetLeaves(operatorSetIndex *big.Int, startOperatorIndex *big.Int, numOperators *big.Int) (OperatorSet, []common.Address, []IStakeRootCompendiumOperatorLeaf, error) {
	return _StakeRootCompendiumStorage.Contract.GetOperatorSetLeaves(&_StakeRootCompendiumStorage.CallOpts, operatorSetIndex, startOperatorIndex, numOperators)
}

// GetOperatorSetRoot is a free data retrieval call binding the contract method 0x117def3a.
//
// Solidity: function getOperatorSetRoot((address,uint32) operatorSet, address[] operators, (uint256,uint256,bytes32)[] operatorLeaves) view returns(bytes32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) GetOperatorSetRoot(opts *bind.CallOpts, operatorSet OperatorSet, operators []common.Address, operatorLeaves []IStakeRootCompendiumOperatorLeaf) ([32]byte, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "getOperatorSetRoot", operatorSet, operators, operatorLeaves)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorSetRoot is a free data retrieval call binding the contract method 0x117def3a.
//
// Solidity: function getOperatorSetRoot((address,uint32) operatorSet, address[] operators, (uint256,uint256,bytes32)[] operatorLeaves) view returns(bytes32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) GetOperatorSetRoot(operatorSet OperatorSet, operators []common.Address, operatorLeaves []IStakeRootCompendiumOperatorLeaf) ([32]byte, error) {
	return _StakeRootCompendiumStorage.Contract.GetOperatorSetRoot(&_StakeRootCompendiumStorage.CallOpts, operatorSet, operators, operatorLeaves)
}

// GetOperatorSetRoot is a free data retrieval call binding the contract method 0x117def3a.
//
// Solidity: function getOperatorSetRoot((address,uint32) operatorSet, address[] operators, (uint256,uint256,bytes32)[] operatorLeaves) view returns(bytes32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) GetOperatorSetRoot(operatorSet OperatorSet, operators []common.Address, operatorLeaves []IStakeRootCompendiumOperatorLeaf) ([32]byte, error) {
	return _StakeRootCompendiumStorage.Contract.GetOperatorSetRoot(&_StakeRootCompendiumStorage.CallOpts, operatorSet, operators, operatorLeaves)
}

// GetStakeRoot is a free data retrieval call binding the contract method 0x3d5f0fd6.
//
// Solidity: function getStakeRoot((address,uint32)[] operatorSetsInStakeTree, bytes32[] operatorSetRoots) view returns(bytes32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) GetStakeRoot(opts *bind.CallOpts, operatorSetsInStakeTree []OperatorSet, operatorSetRoots [][32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "getStakeRoot", operatorSetsInStakeTree, operatorSetRoots)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStakeRoot is a free data retrieval call binding the contract method 0x3d5f0fd6.
//
// Solidity: function getStakeRoot((address,uint32)[] operatorSetsInStakeTree, bytes32[] operatorSetRoots) view returns(bytes32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) GetStakeRoot(operatorSetsInStakeTree []OperatorSet, operatorSetRoots [][32]byte) ([32]byte, error) {
	return _StakeRootCompendiumStorage.Contract.GetStakeRoot(&_StakeRootCompendiumStorage.CallOpts, operatorSetsInStakeTree, operatorSetRoots)
}

// GetStakeRoot is a free data retrieval call binding the contract method 0x3d5f0fd6.
//
// Solidity: function getStakeRoot((address,uint32)[] operatorSetsInStakeTree, bytes32[] operatorSetRoots) view returns(bytes32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) GetStakeRoot(operatorSetsInStakeTree []OperatorSet, operatorSetRoots [][32]byte) ([32]byte, error) {
	return _StakeRootCompendiumStorage.Contract.GetStakeRoot(&_StakeRootCompendiumStorage.CallOpts, operatorSetsInStakeTree, operatorSetRoots)
}

// GetStakeRootSubmission is a free data retrieval call binding the contract method 0x901ad96e.
//
// Solidity: function getStakeRootSubmission(uint32 index) view returns((bytes32,uint32,bool))
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) GetStakeRootSubmission(opts *bind.CallOpts, index uint32) (IStakeRootCompendiumStakeRootSubmission, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "getStakeRootSubmission", index)

	if err != nil {
		return *new(IStakeRootCompendiumStakeRootSubmission), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeRootCompendiumStakeRootSubmission)).(*IStakeRootCompendiumStakeRootSubmission)

	return out0, err

}

// GetStakeRootSubmission is a free data retrieval call binding the contract method 0x901ad96e.
//
// Solidity: function getStakeRootSubmission(uint32 index) view returns((bytes32,uint32,bool))
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) GetStakeRootSubmission(index uint32) (IStakeRootCompendiumStakeRootSubmission, error) {
	return _StakeRootCompendiumStorage.Contract.GetStakeRootSubmission(&_StakeRootCompendiumStorage.CallOpts, index)
}

// GetStakeRootSubmission is a free data retrieval call binding the contract method 0x901ad96e.
//
// Solidity: function getStakeRootSubmission(uint32 index) view returns((bytes32,uint32,bool))
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) GetStakeRootSubmission(index uint32) (IStakeRootCompendiumStakeRootSubmission, error) {
	return _StakeRootCompendiumStorage.Contract.GetStakeRootSubmission(&_StakeRootCompendiumStorage.CallOpts, index)
}

// GetStakes is a free data retrieval call binding the contract method 0xdb5a5711.
//
// Solidity: function getStakes((address,uint32) operatorSet, address operator) view returns(uint256 delegatedStake, uint256 slashableStake)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) GetStakes(opts *bind.CallOpts, operatorSet OperatorSet, operator common.Address) (struct {
	DelegatedStake *big.Int
	SlashableStake *big.Int
}, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "getStakes", operatorSet, operator)

	outstruct := new(struct {
		DelegatedStake *big.Int
		SlashableStake *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DelegatedStake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SlashableStake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakes is a free data retrieval call binding the contract method 0xdb5a5711.
//
// Solidity: function getStakes((address,uint32) operatorSet, address operator) view returns(uint256 delegatedStake, uint256 slashableStake)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) GetStakes(operatorSet OperatorSet, operator common.Address) (struct {
	DelegatedStake *big.Int
	SlashableStake *big.Int
}, error) {
	return _StakeRootCompendiumStorage.Contract.GetStakes(&_StakeRootCompendiumStorage.CallOpts, operatorSet, operator)
}

// GetStakes is a free data retrieval call binding the contract method 0xdb5a5711.
//
// Solidity: function getStakes((address,uint32) operatorSet, address operator) view returns(uint256 delegatedStake, uint256 slashableStake)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) GetStakes(operatorSet OperatorSet, operator common.Address) (struct {
	DelegatedStake *big.Int
	SlashableStake *big.Int
}, error) {
	return _StakeRootCompendiumStorage.Contract.GetStakes(&_StakeRootCompendiumStorage.CallOpts, operatorSet, operator)
}

// ImageId is a free data retrieval call binding the contract method 0xef3f7dd5.
//
// Solidity: function imageId() view returns(bytes32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) ImageId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "imageId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ImageId is a free data retrieval call binding the contract method 0xef3f7dd5.
//
// Solidity: function imageId() view returns(bytes32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) ImageId() ([32]byte, error) {
	return _StakeRootCompendiumStorage.Contract.ImageId(&_StakeRootCompendiumStorage.CallOpts)
}

// ImageId is a free data retrieval call binding the contract method 0xef3f7dd5.
//
// Solidity: function imageId() view returns(bytes32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) ImageId() ([32]byte, error) {
	return _StakeRootCompendiumStorage.Contract.ImageId(&_StakeRootCompendiumStorage.CallOpts)
}

// MinDepositBalance is a free data retrieval call binding the contract method 0xca7bfc9d.
//
// Solidity: function minDepositBalance(uint256 numStrategies) view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) MinDepositBalance(opts *bind.CallOpts, numStrategies *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "minDepositBalance", numStrategies)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDepositBalance is a free data retrieval call binding the contract method 0xca7bfc9d.
//
// Solidity: function minDepositBalance(uint256 numStrategies) view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) MinDepositBalance(numStrategies *big.Int) (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.MinDepositBalance(&_StakeRootCompendiumStorage.CallOpts, numStrategies)
}

// MinDepositBalance is a free data retrieval call binding the contract method 0xca7bfc9d.
//
// Solidity: function minDepositBalance(uint256 numStrategies) view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) MinDepositBalance(numStrategies *big.Int) (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.MinDepositBalance(&_StakeRootCompendiumStorage.CallOpts, numStrategies)
}

// OperatorSets is a free data retrieval call binding the contract method 0xc45d3389.
//
// Solidity: function operatorSets(uint256 ) view returns(address avs, uint32 operatorSetId)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) OperatorSets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Avs           common.Address
	OperatorSetId uint32
}, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "operatorSets", arg0)

	outstruct := new(struct {
		Avs           common.Address
		OperatorSetId uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Avs = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.OperatorSetId = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// OperatorSets is a free data retrieval call binding the contract method 0xc45d3389.
//
// Solidity: function operatorSets(uint256 ) view returns(address avs, uint32 operatorSetId)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) OperatorSets(arg0 *big.Int) (struct {
	Avs           common.Address
	OperatorSetId uint32
}, error) {
	return _StakeRootCompendiumStorage.Contract.OperatorSets(&_StakeRootCompendiumStorage.CallOpts, arg0)
}

// OperatorSets is a free data retrieval call binding the contract method 0xc45d3389.
//
// Solidity: function operatorSets(uint256 ) view returns(address avs, uint32 operatorSetId)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) OperatorSets(arg0 *big.Int) (struct {
	Avs           common.Address
	OperatorSetId uint32
}, error) {
	return _StakeRootCompendiumStorage.Contract.OperatorSets(&_StakeRootCompendiumStorage.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) Owner() (common.Address, error) {
	return _StakeRootCompendiumStorage.Contract.Owner(&_StakeRootCompendiumStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) Owner() (common.Address, error) {
	return _StakeRootCompendiumStorage.Contract.Owner(&_StakeRootCompendiumStorage.CallOpts)
}

// ProofIntervalSeconds is a free data retrieval call binding the contract method 0xa992c74a.
//
// Solidity: function proofIntervalSeconds() view returns(uint32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) ProofIntervalSeconds(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "proofIntervalSeconds")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ProofIntervalSeconds is a free data retrieval call binding the contract method 0xa992c74a.
//
// Solidity: function proofIntervalSeconds() view returns(uint32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) ProofIntervalSeconds() (uint32, error) {
	return _StakeRootCompendiumStorage.Contract.ProofIntervalSeconds(&_StakeRootCompendiumStorage.CallOpts)
}

// ProofIntervalSeconds is a free data retrieval call binding the contract method 0xa992c74a.
//
// Solidity: function proofIntervalSeconds() view returns(uint32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) ProofIntervalSeconds() (uint32, error) {
	return _StakeRootCompendiumStorage.Contract.ProofIntervalSeconds(&_StakeRootCompendiumStorage.CallOpts)
}

// RootConfirmer is a free data retrieval call binding the contract method 0x3716f5f9.
//
// Solidity: function rootConfirmer() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) RootConfirmer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "rootConfirmer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RootConfirmer is a free data retrieval call binding the contract method 0x3716f5f9.
//
// Solidity: function rootConfirmer() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) RootConfirmer() (common.Address, error) {
	return _StakeRootCompendiumStorage.Contract.RootConfirmer(&_StakeRootCompendiumStorage.CallOpts)
}

// RootConfirmer is a free data retrieval call binding the contract method 0x3716f5f9.
//
// Solidity: function rootConfirmer() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) RootConfirmer() (common.Address, error) {
	return _StakeRootCompendiumStorage.Contract.RootConfirmer(&_StakeRootCompendiumStorage.CallOpts)
}

// StakeRootSubmissions is a free data retrieval call binding the contract method 0xfe225d8b.
//
// Solidity: function stakeRootSubmissions(uint256 ) view returns(bytes32 stakeRoot, uint32 calculationTimestamp, bool confirmed)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) StakeRootSubmissions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StakeRoot            [32]byte
	CalculationTimestamp uint32
	Confirmed            bool
}, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "stakeRootSubmissions", arg0)

	outstruct := new(struct {
		StakeRoot            [32]byte
		CalculationTimestamp uint32
		Confirmed            bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakeRoot = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.CalculationTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Confirmed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// StakeRootSubmissions is a free data retrieval call binding the contract method 0xfe225d8b.
//
// Solidity: function stakeRootSubmissions(uint256 ) view returns(bytes32 stakeRoot, uint32 calculationTimestamp, bool confirmed)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) StakeRootSubmissions(arg0 *big.Int) (struct {
	StakeRoot            [32]byte
	CalculationTimestamp uint32
	Confirmed            bool
}, error) {
	return _StakeRootCompendiumStorage.Contract.StakeRootSubmissions(&_StakeRootCompendiumStorage.CallOpts, arg0)
}

// StakeRootSubmissions is a free data retrieval call binding the contract method 0xfe225d8b.
//
// Solidity: function stakeRootSubmissions(uint256 ) view returns(bytes32 stakeRoot, uint32 calculationTimestamp, bool confirmed)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) StakeRootSubmissions(arg0 *big.Int) (struct {
	StakeRoot            [32]byte
	CalculationTimestamp uint32
	Confirmed            bool
}, error) {
	return _StakeRootCompendiumStorage.Contract.StakeRootSubmissions(&_StakeRootCompendiumStorage.CallOpts, arg0)
}

// TotalChargeLastUpdatedTimestamp is a free data retrieval call binding the contract method 0xcca56019.
//
// Solidity: function totalChargeLastUpdatedTimestamp() view returns(uint32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) TotalChargeLastUpdatedTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "totalChargeLastUpdatedTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalChargeLastUpdatedTimestamp is a free data retrieval call binding the contract method 0xcca56019.
//
// Solidity: function totalChargeLastUpdatedTimestamp() view returns(uint32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) TotalChargeLastUpdatedTimestamp() (uint32, error) {
	return _StakeRootCompendiumStorage.Contract.TotalChargeLastUpdatedTimestamp(&_StakeRootCompendiumStorage.CallOpts)
}

// TotalChargeLastUpdatedTimestamp is a free data retrieval call binding the contract method 0xcca56019.
//
// Solidity: function totalChargeLastUpdatedTimestamp() view returns(uint32)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) TotalChargeLastUpdatedTimestamp() (uint32, error) {
	return _StakeRootCompendiumStorage.Contract.TotalChargeLastUpdatedTimestamp(&_StakeRootCompendiumStorage.CallOpts)
}

// TotalChargePerOperatorSetLastUpdate is a free data retrieval call binding the contract method 0x5d466205.
//
// Solidity: function totalChargePerOperatorSetLastUpdate() view returns(uint96)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) TotalChargePerOperatorSetLastUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "totalChargePerOperatorSetLastUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalChargePerOperatorSetLastUpdate is a free data retrieval call binding the contract method 0x5d466205.
//
// Solidity: function totalChargePerOperatorSetLastUpdate() view returns(uint96)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) TotalChargePerOperatorSetLastUpdate() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.TotalChargePerOperatorSetLastUpdate(&_StakeRootCompendiumStorage.CallOpts)
}

// TotalChargePerOperatorSetLastUpdate is a free data retrieval call binding the contract method 0x5d466205.
//
// Solidity: function totalChargePerOperatorSetLastUpdate() view returns(uint96)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) TotalChargePerOperatorSetLastUpdate() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.TotalChargePerOperatorSetLastUpdate(&_StakeRootCompendiumStorage.CallOpts)
}

// TotalChargePerStrategyLastUpdate is a free data retrieval call binding the contract method 0x016769e0.
//
// Solidity: function totalChargePerStrategyLastUpdate() view returns(uint96)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) TotalChargePerStrategyLastUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "totalChargePerStrategyLastUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalChargePerStrategyLastUpdate is a free data retrieval call binding the contract method 0x016769e0.
//
// Solidity: function totalChargePerStrategyLastUpdate() view returns(uint96)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) TotalChargePerStrategyLastUpdate() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.TotalChargePerStrategyLastUpdate(&_StakeRootCompendiumStorage.CallOpts)
}

// TotalChargePerStrategyLastUpdate is a free data retrieval call binding the contract method 0x016769e0.
//
// Solidity: function totalChargePerStrategyLastUpdate() view returns(uint96)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) TotalChargePerStrategyLastUpdate() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.TotalChargePerStrategyLastUpdate(&_StakeRootCompendiumStorage.CallOpts)
}

// TotalStrategies is a free data retrieval call binding the contract method 0xf96d7b80.
//
// Solidity: function totalStrategies() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) TotalStrategies(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "totalStrategies")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStrategies is a free data retrieval call binding the contract method 0xf96d7b80.
//
// Solidity: function totalStrategies() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) TotalStrategies() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.TotalStrategies(&_StakeRootCompendiumStorage.CallOpts)
}

// TotalStrategies is a free data retrieval call binding the contract method 0xf96d7b80.
//
// Solidity: function totalStrategies() view returns(uint256)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) TotalStrategies() (*big.Int, error) {
	return _StakeRootCompendiumStorage.Contract.TotalStrategies(&_StakeRootCompendiumStorage.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeRootCompendiumStorage.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) Verifier() (common.Address, error) {
	return _StakeRootCompendiumStorage.Contract.Verifier(&_StakeRootCompendiumStorage.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageCallerSession) Verifier() (common.Address, error) {
	return _StakeRootCompendiumStorage.Contract.Verifier(&_StakeRootCompendiumStorage.CallOpts)
}

// AddOrModifyStrategiesAndMultipliers is a paid mutator transaction binding the contract method 0xd1ee8b08.
//
// Solidity: function addOrModifyStrategiesAndMultipliers(uint32 operatorSetId, (address,uint96)[] strategiesAndMultipliers) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactor) AddOrModifyStrategiesAndMultipliers(opts *bind.TransactOpts, operatorSetId uint32, strategiesAndMultipliers []IStakeRootCompendiumStrategyAndMultiplier) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.contract.Transact(opts, "addOrModifyStrategiesAndMultipliers", operatorSetId, strategiesAndMultipliers)
}

// AddOrModifyStrategiesAndMultipliers is a paid mutator transaction binding the contract method 0xd1ee8b08.
//
// Solidity: function addOrModifyStrategiesAndMultipliers(uint32 operatorSetId, (address,uint96)[] strategiesAndMultipliers) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) AddOrModifyStrategiesAndMultipliers(operatorSetId uint32, strategiesAndMultipliers []IStakeRootCompendiumStrategyAndMultiplier) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.AddOrModifyStrategiesAndMultipliers(&_StakeRootCompendiumStorage.TransactOpts, operatorSetId, strategiesAndMultipliers)
}

// AddOrModifyStrategiesAndMultipliers is a paid mutator transaction binding the contract method 0xd1ee8b08.
//
// Solidity: function addOrModifyStrategiesAndMultipliers(uint32 operatorSetId, (address,uint96)[] strategiesAndMultipliers) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactorSession) AddOrModifyStrategiesAndMultipliers(operatorSetId uint32, strategiesAndMultipliers []IStakeRootCompendiumStrategyAndMultiplier) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.AddOrModifyStrategiesAndMultipliers(&_StakeRootCompendiumStorage.TransactOpts, operatorSetId, strategiesAndMultipliers)
}

// ConfirmStakeRoot is a paid mutator transaction binding the contract method 0xbf855cc8.
//
// Solidity: function confirmStakeRoot(uint32 calculationTimestamp, bytes32 stakeRoot) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactor) ConfirmStakeRoot(opts *bind.TransactOpts, calculationTimestamp uint32, stakeRoot [32]byte) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.contract.Transact(opts, "confirmStakeRoot", calculationTimestamp, stakeRoot)
}

// ConfirmStakeRoot is a paid mutator transaction binding the contract method 0xbf855cc8.
//
// Solidity: function confirmStakeRoot(uint32 calculationTimestamp, bytes32 stakeRoot) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) ConfirmStakeRoot(calculationTimestamp uint32, stakeRoot [32]byte) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.ConfirmStakeRoot(&_StakeRootCompendiumStorage.TransactOpts, calculationTimestamp, stakeRoot)
}

// ConfirmStakeRoot is a paid mutator transaction binding the contract method 0xbf855cc8.
//
// Solidity: function confirmStakeRoot(uint32 calculationTimestamp, bytes32 stakeRoot) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactorSession) ConfirmStakeRoot(calculationTimestamp uint32, stakeRoot [32]byte) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.ConfirmStakeRoot(&_StakeRootCompendiumStorage.TransactOpts, calculationTimestamp, stakeRoot)
}

// Deposit is a paid mutator transaction binding the contract method 0xe6291b5b.
//
// Solidity: function deposit((address,uint32) operatorSet) payable returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactor) Deposit(opts *bind.TransactOpts, operatorSet OperatorSet) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.contract.Transact(opts, "deposit", operatorSet)
}

// Deposit is a paid mutator transaction binding the contract method 0xe6291b5b.
//
// Solidity: function deposit((address,uint32) operatorSet) payable returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) Deposit(operatorSet OperatorSet) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.Deposit(&_StakeRootCompendiumStorage.TransactOpts, operatorSet)
}

// Deposit is a paid mutator transaction binding the contract method 0xe6291b5b.
//
// Solidity: function deposit((address,uint32) operatorSet) payable returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactorSession) Deposit(operatorSet OperatorSet) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.Deposit(&_StakeRootCompendiumStorage.TransactOpts, operatorSet)
}

// RemoveOperatorSetsFromStakeTree is a paid mutator transaction binding the contract method 0x7e09a4fd.
//
// Solidity: function removeOperatorSetsFromStakeTree((address,uint32)[] operatorSetsToRemove) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactor) RemoveOperatorSetsFromStakeTree(opts *bind.TransactOpts, operatorSetsToRemove []OperatorSet) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.contract.Transact(opts, "removeOperatorSetsFromStakeTree", operatorSetsToRemove)
}

// RemoveOperatorSetsFromStakeTree is a paid mutator transaction binding the contract method 0x7e09a4fd.
//
// Solidity: function removeOperatorSetsFromStakeTree((address,uint32)[] operatorSetsToRemove) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) RemoveOperatorSetsFromStakeTree(operatorSetsToRemove []OperatorSet) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.RemoveOperatorSetsFromStakeTree(&_StakeRootCompendiumStorage.TransactOpts, operatorSetsToRemove)
}

// RemoveOperatorSetsFromStakeTree is a paid mutator transaction binding the contract method 0x7e09a4fd.
//
// Solidity: function removeOperatorSetsFromStakeTree((address,uint32)[] operatorSetsToRemove) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactorSession) RemoveOperatorSetsFromStakeTree(operatorSetsToRemove []OperatorSet) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.RemoveOperatorSetsFromStakeTree(&_StakeRootCompendiumStorage.TransactOpts, operatorSetsToRemove)
}

// RemoveStrategiesAndMultipliers is a paid mutator transaction binding the contract method 0x75cff0a5.
//
// Solidity: function removeStrategiesAndMultipliers(uint32 operatorSetId, address[] strategies) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactor) RemoveStrategiesAndMultipliers(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.contract.Transact(opts, "removeStrategiesAndMultipliers", operatorSetId, strategies)
}

// RemoveStrategiesAndMultipliers is a paid mutator transaction binding the contract method 0x75cff0a5.
//
// Solidity: function removeStrategiesAndMultipliers(uint32 operatorSetId, address[] strategies) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) RemoveStrategiesAndMultipliers(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.RemoveStrategiesAndMultipliers(&_StakeRootCompendiumStorage.TransactOpts, operatorSetId, strategies)
}

// RemoveStrategiesAndMultipliers is a paid mutator transaction binding the contract method 0x75cff0a5.
//
// Solidity: function removeStrategiesAndMultipliers(uint32 operatorSetId, address[] strategies) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactorSession) RemoveStrategiesAndMultipliers(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.RemoveStrategiesAndMultipliers(&_StakeRootCompendiumStorage.TransactOpts, operatorSetId, strategies)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.RenounceOwnership(&_StakeRootCompendiumStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.RenounceOwnership(&_StakeRootCompendiumStorage.TransactOpts)
}

// SetOperatorExtraData is a paid mutator transaction binding the contract method 0xab384788.
//
// Solidity: function setOperatorExtraData(uint32 operatorSetId, address operator, bytes32 extraData) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactor) SetOperatorExtraData(opts *bind.TransactOpts, operatorSetId uint32, operator common.Address, extraData [32]byte) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.contract.Transact(opts, "setOperatorExtraData", operatorSetId, operator, extraData)
}

// SetOperatorExtraData is a paid mutator transaction binding the contract method 0xab384788.
//
// Solidity: function setOperatorExtraData(uint32 operatorSetId, address operator, bytes32 extraData) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) SetOperatorExtraData(operatorSetId uint32, operator common.Address, extraData [32]byte) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.SetOperatorExtraData(&_StakeRootCompendiumStorage.TransactOpts, operatorSetId, operator, extraData)
}

// SetOperatorExtraData is a paid mutator transaction binding the contract method 0xab384788.
//
// Solidity: function setOperatorExtraData(uint32 operatorSetId, address operator, bytes32 extraData) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactorSession) SetOperatorExtraData(operatorSetId uint32, operator common.Address, extraData [32]byte) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.SetOperatorExtraData(&_StakeRootCompendiumStorage.TransactOpts, operatorSetId, operator, extraData)
}

// SetOperatorSetExtraData is a paid mutator transaction binding the contract method 0x1e7bba0d.
//
// Solidity: function setOperatorSetExtraData(uint32 operatorSetId, bytes32 extraData) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactor) SetOperatorSetExtraData(opts *bind.TransactOpts, operatorSetId uint32, extraData [32]byte) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.contract.Transact(opts, "setOperatorSetExtraData", operatorSetId, extraData)
}

// SetOperatorSetExtraData is a paid mutator transaction binding the contract method 0x1e7bba0d.
//
// Solidity: function setOperatorSetExtraData(uint32 operatorSetId, bytes32 extraData) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) SetOperatorSetExtraData(operatorSetId uint32, extraData [32]byte) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.SetOperatorSetExtraData(&_StakeRootCompendiumStorage.TransactOpts, operatorSetId, extraData)
}

// SetOperatorSetExtraData is a paid mutator transaction binding the contract method 0x1e7bba0d.
//
// Solidity: function setOperatorSetExtraData(uint32 operatorSetId, bytes32 extraData) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactorSession) SetOperatorSetExtraData(operatorSetId uint32, extraData [32]byte) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.SetOperatorSetExtraData(&_StakeRootCompendiumStorage.TransactOpts, operatorSetId, extraData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.TransferOwnership(&_StakeRootCompendiumStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.TransferOwnership(&_StakeRootCompendiumStorage.TransactOpts, newOwner)
}

// VerifyStakeRoot is a paid mutator transaction binding the contract method 0x308e813e.
//
// Solidity: function verifyStakeRoot(uint256 _calculationTimestamp, bytes32 _stakeRoot, address _chargeRecipient, uint256 _indexChargePerProof, (uint32) _proof) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactor) VerifyStakeRoot(opts *bind.TransactOpts, _calculationTimestamp *big.Int, _stakeRoot [32]byte, _chargeRecipient common.Address, _indexChargePerProof *big.Int, _proof IStakeRootCompendiumProof) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.contract.Transact(opts, "verifyStakeRoot", _calculationTimestamp, _stakeRoot, _chargeRecipient, _indexChargePerProof, _proof)
}

// VerifyStakeRoot is a paid mutator transaction binding the contract method 0x308e813e.
//
// Solidity: function verifyStakeRoot(uint256 _calculationTimestamp, bytes32 _stakeRoot, address _chargeRecipient, uint256 _indexChargePerProof, (uint32) _proof) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageSession) VerifyStakeRoot(_calculationTimestamp *big.Int, _stakeRoot [32]byte, _chargeRecipient common.Address, _indexChargePerProof *big.Int, _proof IStakeRootCompendiumProof) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.VerifyStakeRoot(&_StakeRootCompendiumStorage.TransactOpts, _calculationTimestamp, _stakeRoot, _chargeRecipient, _indexChargePerProof, _proof)
}

// VerifyStakeRoot is a paid mutator transaction binding the contract method 0x308e813e.
//
// Solidity: function verifyStakeRoot(uint256 _calculationTimestamp, bytes32 _stakeRoot, address _chargeRecipient, uint256 _indexChargePerProof, (uint32) _proof) returns()
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageTransactorSession) VerifyStakeRoot(_calculationTimestamp *big.Int, _stakeRoot [32]byte, _chargeRecipient common.Address, _indexChargePerProof *big.Int, _proof IStakeRootCompendiumProof) (*types.Transaction, error) {
	return _StakeRootCompendiumStorage.Contract.VerifyStakeRoot(&_StakeRootCompendiumStorage.TransactOpts, _calculationTimestamp, _stakeRoot, _chargeRecipient, _indexChargePerProof, _proof)
}

// StakeRootCompendiumStorageImageIdSetIterator is returned from FilterImageIdSet and is used to iterate over the raw logs and unpacked data for ImageIdSet events raised by the StakeRootCompendiumStorage contract.
type StakeRootCompendiumStorageImageIdSetIterator struct {
	Event *StakeRootCompendiumStorageImageIdSet // Event containing the contract specifics and raw log

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
func (it *StakeRootCompendiumStorageImageIdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRootCompendiumStorageImageIdSet)
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
		it.Event = new(StakeRootCompendiumStorageImageIdSet)
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
func (it *StakeRootCompendiumStorageImageIdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRootCompendiumStorageImageIdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRootCompendiumStorageImageIdSet represents a ImageIdSet event raised by the StakeRootCompendiumStorage contract.
type StakeRootCompendiumStorageImageIdSet struct {
	NewImageId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterImageIdSet is a free log retrieval operation binding the contract event 0xf1c38b162bdd25c4ac89067b5e68db75e712c6617f39e3b9b58e399d9463dad8.
//
// Solidity: event ImageIdSet(bytes32 newImageId)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageFilterer) FilterImageIdSet(opts *bind.FilterOpts) (*StakeRootCompendiumStorageImageIdSetIterator, error) {

	logs, sub, err := _StakeRootCompendiumStorage.contract.FilterLogs(opts, "ImageIdSet")
	if err != nil {
		return nil, err
	}
	return &StakeRootCompendiumStorageImageIdSetIterator{contract: _StakeRootCompendiumStorage.contract, event: "ImageIdSet", logs: logs, sub: sub}, nil
}

// WatchImageIdSet is a free log subscription operation binding the contract event 0xf1c38b162bdd25c4ac89067b5e68db75e712c6617f39e3b9b58e399d9463dad8.
//
// Solidity: event ImageIdSet(bytes32 newImageId)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageFilterer) WatchImageIdSet(opts *bind.WatchOpts, sink chan<- *StakeRootCompendiumStorageImageIdSet) (event.Subscription, error) {

	logs, sub, err := _StakeRootCompendiumStorage.contract.WatchLogs(opts, "ImageIdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRootCompendiumStorageImageIdSet)
				if err := _StakeRootCompendiumStorage.contract.UnpackLog(event, "ImageIdSet", log); err != nil {
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

// ParseImageIdSet is a log parse operation binding the contract event 0xf1c38b162bdd25c4ac89067b5e68db75e712c6617f39e3b9b58e399d9463dad8.
//
// Solidity: event ImageIdSet(bytes32 newImageId)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageFilterer) ParseImageIdSet(log types.Log) (*StakeRootCompendiumStorageImageIdSet, error) {
	event := new(StakeRootCompendiumStorageImageIdSet)
	if err := _StakeRootCompendiumStorage.contract.UnpackLog(event, "ImageIdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRootCompendiumStorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakeRootCompendiumStorage contract.
type StakeRootCompendiumStorageInitializedIterator struct {
	Event *StakeRootCompendiumStorageInitialized // Event containing the contract specifics and raw log

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
func (it *StakeRootCompendiumStorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRootCompendiumStorageInitialized)
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
		it.Event = new(StakeRootCompendiumStorageInitialized)
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
func (it *StakeRootCompendiumStorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRootCompendiumStorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRootCompendiumStorageInitialized represents a Initialized event raised by the StakeRootCompendiumStorage contract.
type StakeRootCompendiumStorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakeRootCompendiumStorageInitializedIterator, error) {

	logs, sub, err := _StakeRootCompendiumStorage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakeRootCompendiumStorageInitializedIterator{contract: _StakeRootCompendiumStorage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakeRootCompendiumStorageInitialized) (event.Subscription, error) {

	logs, sub, err := _StakeRootCompendiumStorage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRootCompendiumStorageInitialized)
				if err := _StakeRootCompendiumStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageFilterer) ParseInitialized(log types.Log) (*StakeRootCompendiumStorageInitialized, error) {
	event := new(StakeRootCompendiumStorageInitialized)
	if err := _StakeRootCompendiumStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRootCompendiumStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakeRootCompendiumStorage contract.
type StakeRootCompendiumStorageOwnershipTransferredIterator struct {
	Event *StakeRootCompendiumStorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakeRootCompendiumStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRootCompendiumStorageOwnershipTransferred)
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
		it.Event = new(StakeRootCompendiumStorageOwnershipTransferred)
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
func (it *StakeRootCompendiumStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRootCompendiumStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRootCompendiumStorageOwnershipTransferred represents a OwnershipTransferred event raised by the StakeRootCompendiumStorage contract.
type StakeRootCompendiumStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakeRootCompendiumStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakeRootCompendiumStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakeRootCompendiumStorageOwnershipTransferredIterator{contract: _StakeRootCompendiumStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakeRootCompendiumStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakeRootCompendiumStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRootCompendiumStorageOwnershipTransferred)
				if err := _StakeRootCompendiumStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageFilterer) ParseOwnershipTransferred(log types.Log) (*StakeRootCompendiumStorageOwnershipTransferred, error) {
	event := new(StakeRootCompendiumStorageOwnershipTransferred)
	if err := _StakeRootCompendiumStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRootCompendiumStorageSnarkProofVerifiedIterator is returned from FilterSnarkProofVerified and is used to iterate over the raw logs and unpacked data for SnarkProofVerified events raised by the StakeRootCompendiumStorage contract.
type StakeRootCompendiumStorageSnarkProofVerifiedIterator struct {
	Event *StakeRootCompendiumStorageSnarkProofVerified // Event containing the contract specifics and raw log

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
func (it *StakeRootCompendiumStorageSnarkProofVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRootCompendiumStorageSnarkProofVerified)
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
		it.Event = new(StakeRootCompendiumStorageSnarkProofVerified)
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
func (it *StakeRootCompendiumStorageSnarkProofVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRootCompendiumStorageSnarkProofVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRootCompendiumStorageSnarkProofVerified represents a SnarkProofVerified event raised by the StakeRootCompendiumStorage contract.
type StakeRootCompendiumStorageSnarkProofVerified struct {
	Journal []byte
	Seal    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSnarkProofVerified is a free log retrieval operation binding the contract event 0x423df40688ec6e8fc991c66c91422a63c8f423d0643ff565a9ec648a399e8fff.
//
// Solidity: event SnarkProofVerified(bytes journal, bytes seal)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageFilterer) FilterSnarkProofVerified(opts *bind.FilterOpts) (*StakeRootCompendiumStorageSnarkProofVerifiedIterator, error) {

	logs, sub, err := _StakeRootCompendiumStorage.contract.FilterLogs(opts, "SnarkProofVerified")
	if err != nil {
		return nil, err
	}
	return &StakeRootCompendiumStorageSnarkProofVerifiedIterator{contract: _StakeRootCompendiumStorage.contract, event: "SnarkProofVerified", logs: logs, sub: sub}, nil
}

// WatchSnarkProofVerified is a free log subscription operation binding the contract event 0x423df40688ec6e8fc991c66c91422a63c8f423d0643ff565a9ec648a399e8fff.
//
// Solidity: event SnarkProofVerified(bytes journal, bytes seal)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageFilterer) WatchSnarkProofVerified(opts *bind.WatchOpts, sink chan<- *StakeRootCompendiumStorageSnarkProofVerified) (event.Subscription, error) {

	logs, sub, err := _StakeRootCompendiumStorage.contract.WatchLogs(opts, "SnarkProofVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRootCompendiumStorageSnarkProofVerified)
				if err := _StakeRootCompendiumStorage.contract.UnpackLog(event, "SnarkProofVerified", log); err != nil {
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

// ParseSnarkProofVerified is a log parse operation binding the contract event 0x423df40688ec6e8fc991c66c91422a63c8f423d0643ff565a9ec648a399e8fff.
//
// Solidity: event SnarkProofVerified(bytes journal, bytes seal)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageFilterer) ParseSnarkProofVerified(log types.Log) (*StakeRootCompendiumStorageSnarkProofVerified, error) {
	event := new(StakeRootCompendiumStorageSnarkProofVerified)
	if err := _StakeRootCompendiumStorage.contract.UnpackLog(event, "SnarkProofVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRootCompendiumStorageVerifierSetIterator is returned from FilterVerifierSet and is used to iterate over the raw logs and unpacked data for VerifierSet events raised by the StakeRootCompendiumStorage contract.
type StakeRootCompendiumStorageVerifierSetIterator struct {
	Event *StakeRootCompendiumStorageVerifierSet // Event containing the contract specifics and raw log

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
func (it *StakeRootCompendiumStorageVerifierSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRootCompendiumStorageVerifierSet)
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
		it.Event = new(StakeRootCompendiumStorageVerifierSet)
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
func (it *StakeRootCompendiumStorageVerifierSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRootCompendiumStorageVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRootCompendiumStorageVerifierSet represents a VerifierSet event raised by the StakeRootCompendiumStorage contract.
type StakeRootCompendiumStorageVerifierSet struct {
	NewVerifier common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerifierSet is a free log retrieval operation binding the contract event 0x480b37e3d134e44cb444c9703493c7db564c707cb8a18cecea165f340431da1f.
//
// Solidity: event VerifierSet(address newVerifier)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageFilterer) FilterVerifierSet(opts *bind.FilterOpts) (*StakeRootCompendiumStorageVerifierSetIterator, error) {

	logs, sub, err := _StakeRootCompendiumStorage.contract.FilterLogs(opts, "VerifierSet")
	if err != nil {
		return nil, err
	}
	return &StakeRootCompendiumStorageVerifierSetIterator{contract: _StakeRootCompendiumStorage.contract, event: "VerifierSet", logs: logs, sub: sub}, nil
}

// WatchVerifierSet is a free log subscription operation binding the contract event 0x480b37e3d134e44cb444c9703493c7db564c707cb8a18cecea165f340431da1f.
//
// Solidity: event VerifierSet(address newVerifier)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageFilterer) WatchVerifierSet(opts *bind.WatchOpts, sink chan<- *StakeRootCompendiumStorageVerifierSet) (event.Subscription, error) {

	logs, sub, err := _StakeRootCompendiumStorage.contract.WatchLogs(opts, "VerifierSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRootCompendiumStorageVerifierSet)
				if err := _StakeRootCompendiumStorage.contract.UnpackLog(event, "VerifierSet", log); err != nil {
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

// ParseVerifierSet is a log parse operation binding the contract event 0x480b37e3d134e44cb444c9703493c7db564c707cb8a18cecea165f340431da1f.
//
// Solidity: event VerifierSet(address newVerifier)
func (_StakeRootCompendiumStorage *StakeRootCompendiumStorageFilterer) ParseVerifierSet(log types.Log) (*StakeRootCompendiumStorageVerifierSet, error) {
	event := new(StakeRootCompendiumStorageVerifierSet)
	if err := _StakeRootCompendiumStorage.contract.UnpackLog(event, "VerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
