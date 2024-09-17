// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IStakeRootCompendium

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

// IStakeRootCompendiumMetaData contains all meta data concerning the IStakeRootCompendium contract.
var IStakeRootCompendiumMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"MIN_BALANCE_THRESHOLD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addOrModifyStrategiesAndMultipliers\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRootCompendium.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canWithdrawDepositBalance\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmStakeRoot\",\"inputs\":[{\"name\":\"calculationTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"stakeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getDepositBalance\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumOperatorSets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumStakeRootSubmissions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetLeaves\",\"inputs\":[{\"name\":\"operatorSetIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startOperatorIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRootCompendium.OperatorLeaf[]\",\"components\":[{\"name\":\"delegatedStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slashableStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetRoot\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operatorLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIStakeRootCompendium.OperatorLeaf[]\",\"components\":[{\"name\":\"delegatedStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slashableStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakeRoot\",\"inputs\":[{\"name\":\"operatorSetsInStakeTree\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorSetRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakeRootSubmission\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIStakeRootCompendium.StakeRootSubmission\",\"components\":[{\"name\":\"stakeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"calculationTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"confirmed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakes\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"delegatedStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slashableStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"imageId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minDepositBalance\",\"inputs\":[{\"name\":\"numStrategies\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proofIntervalSeconds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeOperatorSetsFromStakeTree\",\"inputs\":[{\"name\":\"operatorSetsToRemove\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesAndMultipliers\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChargePerProof\",\"inputs\":[{\"name\":\"_chargePerStrategy\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"_chargePerOperatorSet\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorExtraData\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetExtraData\",\"inputs\":[{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProofIntervalSeconds\",\"inputs\":[{\"name\":\"proofIntervalSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRootConfirmer\",\"inputs\":[{\"name\":\"_rootConfirmer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyStakeRoot\",\"inputs\":[{\"name\":\"_calculationTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stakeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_chargeRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexChargePerProof\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_proof\",\"type\":\"tuple\",\"internalType\":\"structIStakeRootCompendium.Proof\",\"components\":[{\"name\":\"x\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ImageIdSet\",\"inputs\":[{\"name\":\"newImageId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SnarkProofVerified\",\"inputs\":[{\"name\":\"journal\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"seal\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifierSet\",\"inputs\":[{\"name\":\"newVerifier\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// IStakeRootCompendiumABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakeRootCompendiumMetaData.ABI instead.
var IStakeRootCompendiumABI = IStakeRootCompendiumMetaData.ABI

// IStakeRootCompendium is an auto generated Go binding around an Ethereum contract.
type IStakeRootCompendium struct {
	IStakeRootCompendiumCaller     // Read-only binding to the contract
	IStakeRootCompendiumTransactor // Write-only binding to the contract
	IStakeRootCompendiumFilterer   // Log filterer for contract events
}

// IStakeRootCompendiumCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakeRootCompendiumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakeRootCompendiumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakeRootCompendiumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakeRootCompendiumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakeRootCompendiumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakeRootCompendiumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakeRootCompendiumSession struct {
	Contract     *IStakeRootCompendium // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IStakeRootCompendiumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakeRootCompendiumCallerSession struct {
	Contract *IStakeRootCompendiumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IStakeRootCompendiumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakeRootCompendiumTransactorSession struct {
	Contract     *IStakeRootCompendiumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IStakeRootCompendiumRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakeRootCompendiumRaw struct {
	Contract *IStakeRootCompendium // Generic contract binding to access the raw methods on
}

// IStakeRootCompendiumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakeRootCompendiumCallerRaw struct {
	Contract *IStakeRootCompendiumCaller // Generic read-only contract binding to access the raw methods on
}

// IStakeRootCompendiumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakeRootCompendiumTransactorRaw struct {
	Contract *IStakeRootCompendiumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStakeRootCompendium creates a new instance of IStakeRootCompendium, bound to a specific deployed contract.
func NewIStakeRootCompendium(address common.Address, backend bind.ContractBackend) (*IStakeRootCompendium, error) {
	contract, err := bindIStakeRootCompendium(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStakeRootCompendium{IStakeRootCompendiumCaller: IStakeRootCompendiumCaller{contract: contract}, IStakeRootCompendiumTransactor: IStakeRootCompendiumTransactor{contract: contract}, IStakeRootCompendiumFilterer: IStakeRootCompendiumFilterer{contract: contract}}, nil
}

// NewIStakeRootCompendiumCaller creates a new read-only instance of IStakeRootCompendium, bound to a specific deployed contract.
func NewIStakeRootCompendiumCaller(address common.Address, caller bind.ContractCaller) (*IStakeRootCompendiumCaller, error) {
	contract, err := bindIStakeRootCompendium(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakeRootCompendiumCaller{contract: contract}, nil
}

// NewIStakeRootCompendiumTransactor creates a new write-only instance of IStakeRootCompendium, bound to a specific deployed contract.
func NewIStakeRootCompendiumTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakeRootCompendiumTransactor, error) {
	contract, err := bindIStakeRootCompendium(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakeRootCompendiumTransactor{contract: contract}, nil
}

// NewIStakeRootCompendiumFilterer creates a new log filterer instance of IStakeRootCompendium, bound to a specific deployed contract.
func NewIStakeRootCompendiumFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakeRootCompendiumFilterer, error) {
	contract, err := bindIStakeRootCompendium(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakeRootCompendiumFilterer{contract: contract}, nil
}

// bindIStakeRootCompendium binds a generic wrapper to an already deployed contract.
func bindIStakeRootCompendium(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStakeRootCompendiumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakeRootCompendium *IStakeRootCompendiumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakeRootCompendium.Contract.IStakeRootCompendiumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakeRootCompendium *IStakeRootCompendiumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.IStakeRootCompendiumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakeRootCompendium *IStakeRootCompendiumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.IStakeRootCompendiumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakeRootCompendium *IStakeRootCompendiumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakeRootCompendium.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakeRootCompendium *IStakeRootCompendiumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakeRootCompendium *IStakeRootCompendiumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.contract.Transact(opts, method, params...)
}

// MINBALANCETHRESHOLD is a free data retrieval call binding the contract method 0xc442daee.
//
// Solidity: function MIN_BALANCE_THRESHOLD() view returns(uint256)
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) MINBALANCETHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "MIN_BALANCE_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBALANCETHRESHOLD is a free data retrieval call binding the contract method 0xc442daee.
//
// Solidity: function MIN_BALANCE_THRESHOLD() view returns(uint256)
func (_IStakeRootCompendium *IStakeRootCompendiumSession) MINBALANCETHRESHOLD() (*big.Int, error) {
	return _IStakeRootCompendium.Contract.MINBALANCETHRESHOLD(&_IStakeRootCompendium.CallOpts)
}

// MINBALANCETHRESHOLD is a free data retrieval call binding the contract method 0xc442daee.
//
// Solidity: function MIN_BALANCE_THRESHOLD() view returns(uint256)
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) MINBALANCETHRESHOLD() (*big.Int, error) {
	return _IStakeRootCompendium.Contract.MINBALANCETHRESHOLD(&_IStakeRootCompendium.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_IStakeRootCompendium *IStakeRootCompendiumSession) AvsDirectory() (common.Address, error) {
	return _IStakeRootCompendium.Contract.AvsDirectory(&_IStakeRootCompendium.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) AvsDirectory() (common.Address, error) {
	return _IStakeRootCompendium.Contract.AvsDirectory(&_IStakeRootCompendium.CallOpts)
}

// CanWithdrawDepositBalance is a free data retrieval call binding the contract method 0x71616631.
//
// Solidity: function canWithdrawDepositBalance((address,uint32) operatorSet) view returns(bool)
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) CanWithdrawDepositBalance(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "canWithdrawDepositBalance", operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanWithdrawDepositBalance is a free data retrieval call binding the contract method 0x71616631.
//
// Solidity: function canWithdrawDepositBalance((address,uint32) operatorSet) view returns(bool)
func (_IStakeRootCompendium *IStakeRootCompendiumSession) CanWithdrawDepositBalance(operatorSet OperatorSet) (bool, error) {
	return _IStakeRootCompendium.Contract.CanWithdrawDepositBalance(&_IStakeRootCompendium.CallOpts, operatorSet)
}

// CanWithdrawDepositBalance is a free data retrieval call binding the contract method 0x71616631.
//
// Solidity: function canWithdrawDepositBalance((address,uint32) operatorSet) view returns(bool)
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) CanWithdrawDepositBalance(operatorSet OperatorSet) (bool, error) {
	return _IStakeRootCompendium.Contract.CanWithdrawDepositBalance(&_IStakeRootCompendium.CallOpts, operatorSet)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_IStakeRootCompendium *IStakeRootCompendiumSession) DelegationManager() (common.Address, error) {
	return _IStakeRootCompendium.Contract.DelegationManager(&_IStakeRootCompendium.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) DelegationManager() (common.Address, error) {
	return _IStakeRootCompendium.Contract.DelegationManager(&_IStakeRootCompendium.CallOpts)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0x2156171a.
//
// Solidity: function getDepositBalance((address,uint32) operatorSet) view returns(uint256 balance)
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) GetDepositBalance(opts *bind.CallOpts, operatorSet OperatorSet) (*big.Int, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "getDepositBalance", operatorSet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositBalance is a free data retrieval call binding the contract method 0x2156171a.
//
// Solidity: function getDepositBalance((address,uint32) operatorSet) view returns(uint256 balance)
func (_IStakeRootCompendium *IStakeRootCompendiumSession) GetDepositBalance(operatorSet OperatorSet) (*big.Int, error) {
	return _IStakeRootCompendium.Contract.GetDepositBalance(&_IStakeRootCompendium.CallOpts, operatorSet)
}

// GetDepositBalance is a free data retrieval call binding the contract method 0x2156171a.
//
// Solidity: function getDepositBalance((address,uint32) operatorSet) view returns(uint256 balance)
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) GetDepositBalance(operatorSet OperatorSet) (*big.Int, error) {
	return _IStakeRootCompendium.Contract.GetDepositBalance(&_IStakeRootCompendium.CallOpts, operatorSet)
}

// GetNumOperatorSets is a free data retrieval call binding the contract method 0x11bea0d1.
//
// Solidity: function getNumOperatorSets() view returns(uint256)
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) GetNumOperatorSets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "getNumOperatorSets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumOperatorSets is a free data retrieval call binding the contract method 0x11bea0d1.
//
// Solidity: function getNumOperatorSets() view returns(uint256)
func (_IStakeRootCompendium *IStakeRootCompendiumSession) GetNumOperatorSets() (*big.Int, error) {
	return _IStakeRootCompendium.Contract.GetNumOperatorSets(&_IStakeRootCompendium.CallOpts)
}

// GetNumOperatorSets is a free data retrieval call binding the contract method 0x11bea0d1.
//
// Solidity: function getNumOperatorSets() view returns(uint256)
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) GetNumOperatorSets() (*big.Int, error) {
	return _IStakeRootCompendium.Contract.GetNumOperatorSets(&_IStakeRootCompendium.CallOpts)
}

// GetNumStakeRootSubmissions is a free data retrieval call binding the contract method 0xd985c7b3.
//
// Solidity: function getNumStakeRootSubmissions() view returns(uint256)
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) GetNumStakeRootSubmissions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "getNumStakeRootSubmissions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumStakeRootSubmissions is a free data retrieval call binding the contract method 0xd985c7b3.
//
// Solidity: function getNumStakeRootSubmissions() view returns(uint256)
func (_IStakeRootCompendium *IStakeRootCompendiumSession) GetNumStakeRootSubmissions() (*big.Int, error) {
	return _IStakeRootCompendium.Contract.GetNumStakeRootSubmissions(&_IStakeRootCompendium.CallOpts)
}

// GetNumStakeRootSubmissions is a free data retrieval call binding the contract method 0xd985c7b3.
//
// Solidity: function getNumStakeRootSubmissions() view returns(uint256)
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) GetNumStakeRootSubmissions() (*big.Int, error) {
	return _IStakeRootCompendium.Contract.GetNumStakeRootSubmissions(&_IStakeRootCompendium.CallOpts)
}

// GetOperatorSetLeaves is a free data retrieval call binding the contract method 0x732d7398.
//
// Solidity: function getOperatorSetLeaves(uint256 operatorSetIndex, uint256 startOperatorIndex, uint256 numOperators) view returns((address,uint32), address[], (uint256,uint256,bytes32)[])
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) GetOperatorSetLeaves(opts *bind.CallOpts, operatorSetIndex *big.Int, startOperatorIndex *big.Int, numOperators *big.Int) (OperatorSet, []common.Address, []IStakeRootCompendiumOperatorLeaf, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "getOperatorSetLeaves", operatorSetIndex, startOperatorIndex, numOperators)

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
func (_IStakeRootCompendium *IStakeRootCompendiumSession) GetOperatorSetLeaves(operatorSetIndex *big.Int, startOperatorIndex *big.Int, numOperators *big.Int) (OperatorSet, []common.Address, []IStakeRootCompendiumOperatorLeaf, error) {
	return _IStakeRootCompendium.Contract.GetOperatorSetLeaves(&_IStakeRootCompendium.CallOpts, operatorSetIndex, startOperatorIndex, numOperators)
}

// GetOperatorSetLeaves is a free data retrieval call binding the contract method 0x732d7398.
//
// Solidity: function getOperatorSetLeaves(uint256 operatorSetIndex, uint256 startOperatorIndex, uint256 numOperators) view returns((address,uint32), address[], (uint256,uint256,bytes32)[])
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) GetOperatorSetLeaves(operatorSetIndex *big.Int, startOperatorIndex *big.Int, numOperators *big.Int) (OperatorSet, []common.Address, []IStakeRootCompendiumOperatorLeaf, error) {
	return _IStakeRootCompendium.Contract.GetOperatorSetLeaves(&_IStakeRootCompendium.CallOpts, operatorSetIndex, startOperatorIndex, numOperators)
}

// GetOperatorSetRoot is a free data retrieval call binding the contract method 0x117def3a.
//
// Solidity: function getOperatorSetRoot((address,uint32) operatorSet, address[] operators, (uint256,uint256,bytes32)[] operatorLeaves) view returns(bytes32)
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) GetOperatorSetRoot(opts *bind.CallOpts, operatorSet OperatorSet, operators []common.Address, operatorLeaves []IStakeRootCompendiumOperatorLeaf) ([32]byte, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "getOperatorSetRoot", operatorSet, operators, operatorLeaves)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorSetRoot is a free data retrieval call binding the contract method 0x117def3a.
//
// Solidity: function getOperatorSetRoot((address,uint32) operatorSet, address[] operators, (uint256,uint256,bytes32)[] operatorLeaves) view returns(bytes32)
func (_IStakeRootCompendium *IStakeRootCompendiumSession) GetOperatorSetRoot(operatorSet OperatorSet, operators []common.Address, operatorLeaves []IStakeRootCompendiumOperatorLeaf) ([32]byte, error) {
	return _IStakeRootCompendium.Contract.GetOperatorSetRoot(&_IStakeRootCompendium.CallOpts, operatorSet, operators, operatorLeaves)
}

// GetOperatorSetRoot is a free data retrieval call binding the contract method 0x117def3a.
//
// Solidity: function getOperatorSetRoot((address,uint32) operatorSet, address[] operators, (uint256,uint256,bytes32)[] operatorLeaves) view returns(bytes32)
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) GetOperatorSetRoot(operatorSet OperatorSet, operators []common.Address, operatorLeaves []IStakeRootCompendiumOperatorLeaf) ([32]byte, error) {
	return _IStakeRootCompendium.Contract.GetOperatorSetRoot(&_IStakeRootCompendium.CallOpts, operatorSet, operators, operatorLeaves)
}

// GetStakeRoot is a free data retrieval call binding the contract method 0x3d5f0fd6.
//
// Solidity: function getStakeRoot((address,uint32)[] operatorSetsInStakeTree, bytes32[] operatorSetRoots) view returns(bytes32)
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) GetStakeRoot(opts *bind.CallOpts, operatorSetsInStakeTree []OperatorSet, operatorSetRoots [][32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "getStakeRoot", operatorSetsInStakeTree, operatorSetRoots)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStakeRoot is a free data retrieval call binding the contract method 0x3d5f0fd6.
//
// Solidity: function getStakeRoot((address,uint32)[] operatorSetsInStakeTree, bytes32[] operatorSetRoots) view returns(bytes32)
func (_IStakeRootCompendium *IStakeRootCompendiumSession) GetStakeRoot(operatorSetsInStakeTree []OperatorSet, operatorSetRoots [][32]byte) ([32]byte, error) {
	return _IStakeRootCompendium.Contract.GetStakeRoot(&_IStakeRootCompendium.CallOpts, operatorSetsInStakeTree, operatorSetRoots)
}

// GetStakeRoot is a free data retrieval call binding the contract method 0x3d5f0fd6.
//
// Solidity: function getStakeRoot((address,uint32)[] operatorSetsInStakeTree, bytes32[] operatorSetRoots) view returns(bytes32)
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) GetStakeRoot(operatorSetsInStakeTree []OperatorSet, operatorSetRoots [][32]byte) ([32]byte, error) {
	return _IStakeRootCompendium.Contract.GetStakeRoot(&_IStakeRootCompendium.CallOpts, operatorSetsInStakeTree, operatorSetRoots)
}

// GetStakeRootSubmission is a free data retrieval call binding the contract method 0x901ad96e.
//
// Solidity: function getStakeRootSubmission(uint32 index) view returns((bytes32,uint32,bool))
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) GetStakeRootSubmission(opts *bind.CallOpts, index uint32) (IStakeRootCompendiumStakeRootSubmission, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "getStakeRootSubmission", index)

	if err != nil {
		return *new(IStakeRootCompendiumStakeRootSubmission), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeRootCompendiumStakeRootSubmission)).(*IStakeRootCompendiumStakeRootSubmission)

	return out0, err

}

// GetStakeRootSubmission is a free data retrieval call binding the contract method 0x901ad96e.
//
// Solidity: function getStakeRootSubmission(uint32 index) view returns((bytes32,uint32,bool))
func (_IStakeRootCompendium *IStakeRootCompendiumSession) GetStakeRootSubmission(index uint32) (IStakeRootCompendiumStakeRootSubmission, error) {
	return _IStakeRootCompendium.Contract.GetStakeRootSubmission(&_IStakeRootCompendium.CallOpts, index)
}

// GetStakeRootSubmission is a free data retrieval call binding the contract method 0x901ad96e.
//
// Solidity: function getStakeRootSubmission(uint32 index) view returns((bytes32,uint32,bool))
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) GetStakeRootSubmission(index uint32) (IStakeRootCompendiumStakeRootSubmission, error) {
	return _IStakeRootCompendium.Contract.GetStakeRootSubmission(&_IStakeRootCompendium.CallOpts, index)
}

// GetStakes is a free data retrieval call binding the contract method 0xdb5a5711.
//
// Solidity: function getStakes((address,uint32) operatorSet, address operator) view returns(uint256 delegatedStake, uint256 slashableStake)
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) GetStakes(opts *bind.CallOpts, operatorSet OperatorSet, operator common.Address) (struct {
	DelegatedStake *big.Int
	SlashableStake *big.Int
}, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "getStakes", operatorSet, operator)

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
func (_IStakeRootCompendium *IStakeRootCompendiumSession) GetStakes(operatorSet OperatorSet, operator common.Address) (struct {
	DelegatedStake *big.Int
	SlashableStake *big.Int
}, error) {
	return _IStakeRootCompendium.Contract.GetStakes(&_IStakeRootCompendium.CallOpts, operatorSet, operator)
}

// GetStakes is a free data retrieval call binding the contract method 0xdb5a5711.
//
// Solidity: function getStakes((address,uint32) operatorSet, address operator) view returns(uint256 delegatedStake, uint256 slashableStake)
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) GetStakes(operatorSet OperatorSet, operator common.Address) (struct {
	DelegatedStake *big.Int
	SlashableStake *big.Int
}, error) {
	return _IStakeRootCompendium.Contract.GetStakes(&_IStakeRootCompendium.CallOpts, operatorSet, operator)
}

// ImageId is a free data retrieval call binding the contract method 0xef3f7dd5.
//
// Solidity: function imageId() view returns(bytes32)
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) ImageId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "imageId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ImageId is a free data retrieval call binding the contract method 0xef3f7dd5.
//
// Solidity: function imageId() view returns(bytes32)
func (_IStakeRootCompendium *IStakeRootCompendiumSession) ImageId() ([32]byte, error) {
	return _IStakeRootCompendium.Contract.ImageId(&_IStakeRootCompendium.CallOpts)
}

// ImageId is a free data retrieval call binding the contract method 0xef3f7dd5.
//
// Solidity: function imageId() view returns(bytes32)
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) ImageId() ([32]byte, error) {
	return _IStakeRootCompendium.Contract.ImageId(&_IStakeRootCompendium.CallOpts)
}

// MinDepositBalance is a free data retrieval call binding the contract method 0xca7bfc9d.
//
// Solidity: function minDepositBalance(uint256 numStrategies) view returns(uint256)
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) MinDepositBalance(opts *bind.CallOpts, numStrategies *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "minDepositBalance", numStrategies)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDepositBalance is a free data retrieval call binding the contract method 0xca7bfc9d.
//
// Solidity: function minDepositBalance(uint256 numStrategies) view returns(uint256)
func (_IStakeRootCompendium *IStakeRootCompendiumSession) MinDepositBalance(numStrategies *big.Int) (*big.Int, error) {
	return _IStakeRootCompendium.Contract.MinDepositBalance(&_IStakeRootCompendium.CallOpts, numStrategies)
}

// MinDepositBalance is a free data retrieval call binding the contract method 0xca7bfc9d.
//
// Solidity: function minDepositBalance(uint256 numStrategies) view returns(uint256)
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) MinDepositBalance(numStrategies *big.Int) (*big.Int, error) {
	return _IStakeRootCompendium.Contract.MinDepositBalance(&_IStakeRootCompendium.CallOpts, numStrategies)
}

// ProofIntervalSeconds is a free data retrieval call binding the contract method 0xa992c74a.
//
// Solidity: function proofIntervalSeconds() view returns(uint32)
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) ProofIntervalSeconds(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "proofIntervalSeconds")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ProofIntervalSeconds is a free data retrieval call binding the contract method 0xa992c74a.
//
// Solidity: function proofIntervalSeconds() view returns(uint32)
func (_IStakeRootCompendium *IStakeRootCompendiumSession) ProofIntervalSeconds() (uint32, error) {
	return _IStakeRootCompendium.Contract.ProofIntervalSeconds(&_IStakeRootCompendium.CallOpts)
}

// ProofIntervalSeconds is a free data retrieval call binding the contract method 0xa992c74a.
//
// Solidity: function proofIntervalSeconds() view returns(uint32)
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) ProofIntervalSeconds() (uint32, error) {
	return _IStakeRootCompendium.Contract.ProofIntervalSeconds(&_IStakeRootCompendium.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_IStakeRootCompendium *IStakeRootCompendiumCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IStakeRootCompendium.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_IStakeRootCompendium *IStakeRootCompendiumSession) Verifier() (common.Address, error) {
	return _IStakeRootCompendium.Contract.Verifier(&_IStakeRootCompendium.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_IStakeRootCompendium *IStakeRootCompendiumCallerSession) Verifier() (common.Address, error) {
	return _IStakeRootCompendium.Contract.Verifier(&_IStakeRootCompendium.CallOpts)
}

// AddOrModifyStrategiesAndMultipliers is a paid mutator transaction binding the contract method 0xd1ee8b08.
//
// Solidity: function addOrModifyStrategiesAndMultipliers(uint32 operatorSetId, (address,uint96)[] strategiesAndMultipliers) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactor) AddOrModifyStrategiesAndMultipliers(opts *bind.TransactOpts, operatorSetId uint32, strategiesAndMultipliers []IStakeRootCompendiumStrategyAndMultiplier) (*types.Transaction, error) {
	return _IStakeRootCompendium.contract.Transact(opts, "addOrModifyStrategiesAndMultipliers", operatorSetId, strategiesAndMultipliers)
}

// AddOrModifyStrategiesAndMultipliers is a paid mutator transaction binding the contract method 0xd1ee8b08.
//
// Solidity: function addOrModifyStrategiesAndMultipliers(uint32 operatorSetId, (address,uint96)[] strategiesAndMultipliers) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumSession) AddOrModifyStrategiesAndMultipliers(operatorSetId uint32, strategiesAndMultipliers []IStakeRootCompendiumStrategyAndMultiplier) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.AddOrModifyStrategiesAndMultipliers(&_IStakeRootCompendium.TransactOpts, operatorSetId, strategiesAndMultipliers)
}

// AddOrModifyStrategiesAndMultipliers is a paid mutator transaction binding the contract method 0xd1ee8b08.
//
// Solidity: function addOrModifyStrategiesAndMultipliers(uint32 operatorSetId, (address,uint96)[] strategiesAndMultipliers) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactorSession) AddOrModifyStrategiesAndMultipliers(operatorSetId uint32, strategiesAndMultipliers []IStakeRootCompendiumStrategyAndMultiplier) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.AddOrModifyStrategiesAndMultipliers(&_IStakeRootCompendium.TransactOpts, operatorSetId, strategiesAndMultipliers)
}

// ConfirmStakeRoot is a paid mutator transaction binding the contract method 0xbf855cc8.
//
// Solidity: function confirmStakeRoot(uint32 calculationTimestamp, bytes32 stakeRoot) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactor) ConfirmStakeRoot(opts *bind.TransactOpts, calculationTimestamp uint32, stakeRoot [32]byte) (*types.Transaction, error) {
	return _IStakeRootCompendium.contract.Transact(opts, "confirmStakeRoot", calculationTimestamp, stakeRoot)
}

// ConfirmStakeRoot is a paid mutator transaction binding the contract method 0xbf855cc8.
//
// Solidity: function confirmStakeRoot(uint32 calculationTimestamp, bytes32 stakeRoot) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumSession) ConfirmStakeRoot(calculationTimestamp uint32, stakeRoot [32]byte) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.ConfirmStakeRoot(&_IStakeRootCompendium.TransactOpts, calculationTimestamp, stakeRoot)
}

// ConfirmStakeRoot is a paid mutator transaction binding the contract method 0xbf855cc8.
//
// Solidity: function confirmStakeRoot(uint32 calculationTimestamp, bytes32 stakeRoot) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactorSession) ConfirmStakeRoot(calculationTimestamp uint32, stakeRoot [32]byte) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.ConfirmStakeRoot(&_IStakeRootCompendium.TransactOpts, calculationTimestamp, stakeRoot)
}

// Deposit is a paid mutator transaction binding the contract method 0xe6291b5b.
//
// Solidity: function deposit((address,uint32) operatorSet) payable returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactor) Deposit(opts *bind.TransactOpts, operatorSet OperatorSet) (*types.Transaction, error) {
	return _IStakeRootCompendium.contract.Transact(opts, "deposit", operatorSet)
}

// Deposit is a paid mutator transaction binding the contract method 0xe6291b5b.
//
// Solidity: function deposit((address,uint32) operatorSet) payable returns()
func (_IStakeRootCompendium *IStakeRootCompendiumSession) Deposit(operatorSet OperatorSet) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.Deposit(&_IStakeRootCompendium.TransactOpts, operatorSet)
}

// Deposit is a paid mutator transaction binding the contract method 0xe6291b5b.
//
// Solidity: function deposit((address,uint32) operatorSet) payable returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactorSession) Deposit(operatorSet OperatorSet) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.Deposit(&_IStakeRootCompendium.TransactOpts, operatorSet)
}

// RemoveOperatorSetsFromStakeTree is a paid mutator transaction binding the contract method 0x7e09a4fd.
//
// Solidity: function removeOperatorSetsFromStakeTree((address,uint32)[] operatorSetsToRemove) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactor) RemoveOperatorSetsFromStakeTree(opts *bind.TransactOpts, operatorSetsToRemove []OperatorSet) (*types.Transaction, error) {
	return _IStakeRootCompendium.contract.Transact(opts, "removeOperatorSetsFromStakeTree", operatorSetsToRemove)
}

// RemoveOperatorSetsFromStakeTree is a paid mutator transaction binding the contract method 0x7e09a4fd.
//
// Solidity: function removeOperatorSetsFromStakeTree((address,uint32)[] operatorSetsToRemove) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumSession) RemoveOperatorSetsFromStakeTree(operatorSetsToRemove []OperatorSet) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.RemoveOperatorSetsFromStakeTree(&_IStakeRootCompendium.TransactOpts, operatorSetsToRemove)
}

// RemoveOperatorSetsFromStakeTree is a paid mutator transaction binding the contract method 0x7e09a4fd.
//
// Solidity: function removeOperatorSetsFromStakeTree((address,uint32)[] operatorSetsToRemove) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactorSession) RemoveOperatorSetsFromStakeTree(operatorSetsToRemove []OperatorSet) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.RemoveOperatorSetsFromStakeTree(&_IStakeRootCompendium.TransactOpts, operatorSetsToRemove)
}

// RemoveStrategiesAndMultipliers is a paid mutator transaction binding the contract method 0x75cff0a5.
//
// Solidity: function removeStrategiesAndMultipliers(uint32 operatorSetId, address[] strategies) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactor) RemoveStrategiesAndMultipliers(opts *bind.TransactOpts, operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _IStakeRootCompendium.contract.Transact(opts, "removeStrategiesAndMultipliers", operatorSetId, strategies)
}

// RemoveStrategiesAndMultipliers is a paid mutator transaction binding the contract method 0x75cff0a5.
//
// Solidity: function removeStrategiesAndMultipliers(uint32 operatorSetId, address[] strategies) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumSession) RemoveStrategiesAndMultipliers(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.RemoveStrategiesAndMultipliers(&_IStakeRootCompendium.TransactOpts, operatorSetId, strategies)
}

// RemoveStrategiesAndMultipliers is a paid mutator transaction binding the contract method 0x75cff0a5.
//
// Solidity: function removeStrategiesAndMultipliers(uint32 operatorSetId, address[] strategies) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactorSession) RemoveStrategiesAndMultipliers(operatorSetId uint32, strategies []common.Address) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.RemoveStrategiesAndMultipliers(&_IStakeRootCompendium.TransactOpts, operatorSetId, strategies)
}

// SetChargePerProof is a paid mutator transaction binding the contract method 0x1016d85b.
//
// Solidity: function setChargePerProof(uint96 _chargePerStrategy, uint96 _chargePerOperatorSet) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactor) SetChargePerProof(opts *bind.TransactOpts, _chargePerStrategy *big.Int, _chargePerOperatorSet *big.Int) (*types.Transaction, error) {
	return _IStakeRootCompendium.contract.Transact(opts, "setChargePerProof", _chargePerStrategy, _chargePerOperatorSet)
}

// SetChargePerProof is a paid mutator transaction binding the contract method 0x1016d85b.
//
// Solidity: function setChargePerProof(uint96 _chargePerStrategy, uint96 _chargePerOperatorSet) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumSession) SetChargePerProof(_chargePerStrategy *big.Int, _chargePerOperatorSet *big.Int) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.SetChargePerProof(&_IStakeRootCompendium.TransactOpts, _chargePerStrategy, _chargePerOperatorSet)
}

// SetChargePerProof is a paid mutator transaction binding the contract method 0x1016d85b.
//
// Solidity: function setChargePerProof(uint96 _chargePerStrategy, uint96 _chargePerOperatorSet) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactorSession) SetChargePerProof(_chargePerStrategy *big.Int, _chargePerOperatorSet *big.Int) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.SetChargePerProof(&_IStakeRootCompendium.TransactOpts, _chargePerStrategy, _chargePerOperatorSet)
}

// SetOperatorExtraData is a paid mutator transaction binding the contract method 0xab384788.
//
// Solidity: function setOperatorExtraData(uint32 operatorSetId, address operator, bytes32 extraData) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactor) SetOperatorExtraData(opts *bind.TransactOpts, operatorSetId uint32, operator common.Address, extraData [32]byte) (*types.Transaction, error) {
	return _IStakeRootCompendium.contract.Transact(opts, "setOperatorExtraData", operatorSetId, operator, extraData)
}

// SetOperatorExtraData is a paid mutator transaction binding the contract method 0xab384788.
//
// Solidity: function setOperatorExtraData(uint32 operatorSetId, address operator, bytes32 extraData) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumSession) SetOperatorExtraData(operatorSetId uint32, operator common.Address, extraData [32]byte) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.SetOperatorExtraData(&_IStakeRootCompendium.TransactOpts, operatorSetId, operator, extraData)
}

// SetOperatorExtraData is a paid mutator transaction binding the contract method 0xab384788.
//
// Solidity: function setOperatorExtraData(uint32 operatorSetId, address operator, bytes32 extraData) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactorSession) SetOperatorExtraData(operatorSetId uint32, operator common.Address, extraData [32]byte) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.SetOperatorExtraData(&_IStakeRootCompendium.TransactOpts, operatorSetId, operator, extraData)
}

// SetOperatorSetExtraData is a paid mutator transaction binding the contract method 0x1e7bba0d.
//
// Solidity: function setOperatorSetExtraData(uint32 operatorSetId, bytes32 extraData) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactor) SetOperatorSetExtraData(opts *bind.TransactOpts, operatorSetId uint32, extraData [32]byte) (*types.Transaction, error) {
	return _IStakeRootCompendium.contract.Transact(opts, "setOperatorSetExtraData", operatorSetId, extraData)
}

// SetOperatorSetExtraData is a paid mutator transaction binding the contract method 0x1e7bba0d.
//
// Solidity: function setOperatorSetExtraData(uint32 operatorSetId, bytes32 extraData) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumSession) SetOperatorSetExtraData(operatorSetId uint32, extraData [32]byte) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.SetOperatorSetExtraData(&_IStakeRootCompendium.TransactOpts, operatorSetId, extraData)
}

// SetOperatorSetExtraData is a paid mutator transaction binding the contract method 0x1e7bba0d.
//
// Solidity: function setOperatorSetExtraData(uint32 operatorSetId, bytes32 extraData) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactorSession) SetOperatorSetExtraData(operatorSetId uint32, extraData [32]byte) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.SetOperatorSetExtraData(&_IStakeRootCompendium.TransactOpts, operatorSetId, extraData)
}

// SetProofIntervalSeconds is a paid mutator transaction binding the contract method 0xafc6c70b.
//
// Solidity: function setProofIntervalSeconds(uint32 proofIntervalSeconds) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactor) SetProofIntervalSeconds(opts *bind.TransactOpts, proofIntervalSeconds uint32) (*types.Transaction, error) {
	return _IStakeRootCompendium.contract.Transact(opts, "setProofIntervalSeconds", proofIntervalSeconds)
}

// SetProofIntervalSeconds is a paid mutator transaction binding the contract method 0xafc6c70b.
//
// Solidity: function setProofIntervalSeconds(uint32 proofIntervalSeconds) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumSession) SetProofIntervalSeconds(proofIntervalSeconds uint32) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.SetProofIntervalSeconds(&_IStakeRootCompendium.TransactOpts, proofIntervalSeconds)
}

// SetProofIntervalSeconds is a paid mutator transaction binding the contract method 0xafc6c70b.
//
// Solidity: function setProofIntervalSeconds(uint32 proofIntervalSeconds) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactorSession) SetProofIntervalSeconds(proofIntervalSeconds uint32) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.SetProofIntervalSeconds(&_IStakeRootCompendium.TransactOpts, proofIntervalSeconds)
}

// SetRootConfirmer is a paid mutator transaction binding the contract method 0x758e3bfb.
//
// Solidity: function setRootConfirmer(address _rootConfirmer) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactor) SetRootConfirmer(opts *bind.TransactOpts, _rootConfirmer common.Address) (*types.Transaction, error) {
	return _IStakeRootCompendium.contract.Transact(opts, "setRootConfirmer", _rootConfirmer)
}

// SetRootConfirmer is a paid mutator transaction binding the contract method 0x758e3bfb.
//
// Solidity: function setRootConfirmer(address _rootConfirmer) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumSession) SetRootConfirmer(_rootConfirmer common.Address) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.SetRootConfirmer(&_IStakeRootCompendium.TransactOpts, _rootConfirmer)
}

// SetRootConfirmer is a paid mutator transaction binding the contract method 0x758e3bfb.
//
// Solidity: function setRootConfirmer(address _rootConfirmer) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactorSession) SetRootConfirmer(_rootConfirmer common.Address) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.SetRootConfirmer(&_IStakeRootCompendium.TransactOpts, _rootConfirmer)
}

// VerifyStakeRoot is a paid mutator transaction binding the contract method 0x308e813e.
//
// Solidity: function verifyStakeRoot(uint256 _calculationTimestamp, bytes32 _stakeRoot, address _chargeRecipient, uint256 _indexChargePerProof, (uint32) _proof) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactor) VerifyStakeRoot(opts *bind.TransactOpts, _calculationTimestamp *big.Int, _stakeRoot [32]byte, _chargeRecipient common.Address, _indexChargePerProof *big.Int, _proof IStakeRootCompendiumProof) (*types.Transaction, error) {
	return _IStakeRootCompendium.contract.Transact(opts, "verifyStakeRoot", _calculationTimestamp, _stakeRoot, _chargeRecipient, _indexChargePerProof, _proof)
}

// VerifyStakeRoot is a paid mutator transaction binding the contract method 0x308e813e.
//
// Solidity: function verifyStakeRoot(uint256 _calculationTimestamp, bytes32 _stakeRoot, address _chargeRecipient, uint256 _indexChargePerProof, (uint32) _proof) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumSession) VerifyStakeRoot(_calculationTimestamp *big.Int, _stakeRoot [32]byte, _chargeRecipient common.Address, _indexChargePerProof *big.Int, _proof IStakeRootCompendiumProof) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.VerifyStakeRoot(&_IStakeRootCompendium.TransactOpts, _calculationTimestamp, _stakeRoot, _chargeRecipient, _indexChargePerProof, _proof)
}

// VerifyStakeRoot is a paid mutator transaction binding the contract method 0x308e813e.
//
// Solidity: function verifyStakeRoot(uint256 _calculationTimestamp, bytes32 _stakeRoot, address _chargeRecipient, uint256 _indexChargePerProof, (uint32) _proof) returns()
func (_IStakeRootCompendium *IStakeRootCompendiumTransactorSession) VerifyStakeRoot(_calculationTimestamp *big.Int, _stakeRoot [32]byte, _chargeRecipient common.Address, _indexChargePerProof *big.Int, _proof IStakeRootCompendiumProof) (*types.Transaction, error) {
	return _IStakeRootCompendium.Contract.VerifyStakeRoot(&_IStakeRootCompendium.TransactOpts, _calculationTimestamp, _stakeRoot, _chargeRecipient, _indexChargePerProof, _proof)
}

// IStakeRootCompendiumImageIdSetIterator is returned from FilterImageIdSet and is used to iterate over the raw logs and unpacked data for ImageIdSet events raised by the IStakeRootCompendium contract.
type IStakeRootCompendiumImageIdSetIterator struct {
	Event *IStakeRootCompendiumImageIdSet // Event containing the contract specifics and raw log

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
func (it *IStakeRootCompendiumImageIdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakeRootCompendiumImageIdSet)
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
		it.Event = new(IStakeRootCompendiumImageIdSet)
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
func (it *IStakeRootCompendiumImageIdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakeRootCompendiumImageIdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakeRootCompendiumImageIdSet represents a ImageIdSet event raised by the IStakeRootCompendium contract.
type IStakeRootCompendiumImageIdSet struct {
	NewImageId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterImageIdSet is a free log retrieval operation binding the contract event 0xf1c38b162bdd25c4ac89067b5e68db75e712c6617f39e3b9b58e399d9463dad8.
//
// Solidity: event ImageIdSet(bytes32 newImageId)
func (_IStakeRootCompendium *IStakeRootCompendiumFilterer) FilterImageIdSet(opts *bind.FilterOpts) (*IStakeRootCompendiumImageIdSetIterator, error) {

	logs, sub, err := _IStakeRootCompendium.contract.FilterLogs(opts, "ImageIdSet")
	if err != nil {
		return nil, err
	}
	return &IStakeRootCompendiumImageIdSetIterator{contract: _IStakeRootCompendium.contract, event: "ImageIdSet", logs: logs, sub: sub}, nil
}

// WatchImageIdSet is a free log subscription operation binding the contract event 0xf1c38b162bdd25c4ac89067b5e68db75e712c6617f39e3b9b58e399d9463dad8.
//
// Solidity: event ImageIdSet(bytes32 newImageId)
func (_IStakeRootCompendium *IStakeRootCompendiumFilterer) WatchImageIdSet(opts *bind.WatchOpts, sink chan<- *IStakeRootCompendiumImageIdSet) (event.Subscription, error) {

	logs, sub, err := _IStakeRootCompendium.contract.WatchLogs(opts, "ImageIdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakeRootCompendiumImageIdSet)
				if err := _IStakeRootCompendium.contract.UnpackLog(event, "ImageIdSet", log); err != nil {
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
func (_IStakeRootCompendium *IStakeRootCompendiumFilterer) ParseImageIdSet(log types.Log) (*IStakeRootCompendiumImageIdSet, error) {
	event := new(IStakeRootCompendiumImageIdSet)
	if err := _IStakeRootCompendium.contract.UnpackLog(event, "ImageIdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakeRootCompendiumSnarkProofVerifiedIterator is returned from FilterSnarkProofVerified and is used to iterate over the raw logs and unpacked data for SnarkProofVerified events raised by the IStakeRootCompendium contract.
type IStakeRootCompendiumSnarkProofVerifiedIterator struct {
	Event *IStakeRootCompendiumSnarkProofVerified // Event containing the contract specifics and raw log

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
func (it *IStakeRootCompendiumSnarkProofVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakeRootCompendiumSnarkProofVerified)
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
		it.Event = new(IStakeRootCompendiumSnarkProofVerified)
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
func (it *IStakeRootCompendiumSnarkProofVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakeRootCompendiumSnarkProofVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakeRootCompendiumSnarkProofVerified represents a SnarkProofVerified event raised by the IStakeRootCompendium contract.
type IStakeRootCompendiumSnarkProofVerified struct {
	Journal []byte
	Seal    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSnarkProofVerified is a free log retrieval operation binding the contract event 0x423df40688ec6e8fc991c66c91422a63c8f423d0643ff565a9ec648a399e8fff.
//
// Solidity: event SnarkProofVerified(bytes journal, bytes seal)
func (_IStakeRootCompendium *IStakeRootCompendiumFilterer) FilterSnarkProofVerified(opts *bind.FilterOpts) (*IStakeRootCompendiumSnarkProofVerifiedIterator, error) {

	logs, sub, err := _IStakeRootCompendium.contract.FilterLogs(opts, "SnarkProofVerified")
	if err != nil {
		return nil, err
	}
	return &IStakeRootCompendiumSnarkProofVerifiedIterator{contract: _IStakeRootCompendium.contract, event: "SnarkProofVerified", logs: logs, sub: sub}, nil
}

// WatchSnarkProofVerified is a free log subscription operation binding the contract event 0x423df40688ec6e8fc991c66c91422a63c8f423d0643ff565a9ec648a399e8fff.
//
// Solidity: event SnarkProofVerified(bytes journal, bytes seal)
func (_IStakeRootCompendium *IStakeRootCompendiumFilterer) WatchSnarkProofVerified(opts *bind.WatchOpts, sink chan<- *IStakeRootCompendiumSnarkProofVerified) (event.Subscription, error) {

	logs, sub, err := _IStakeRootCompendium.contract.WatchLogs(opts, "SnarkProofVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakeRootCompendiumSnarkProofVerified)
				if err := _IStakeRootCompendium.contract.UnpackLog(event, "SnarkProofVerified", log); err != nil {
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
func (_IStakeRootCompendium *IStakeRootCompendiumFilterer) ParseSnarkProofVerified(log types.Log) (*IStakeRootCompendiumSnarkProofVerified, error) {
	event := new(IStakeRootCompendiumSnarkProofVerified)
	if err := _IStakeRootCompendium.contract.UnpackLog(event, "SnarkProofVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakeRootCompendiumVerifierSetIterator is returned from FilterVerifierSet and is used to iterate over the raw logs and unpacked data for VerifierSet events raised by the IStakeRootCompendium contract.
type IStakeRootCompendiumVerifierSetIterator struct {
	Event *IStakeRootCompendiumVerifierSet // Event containing the contract specifics and raw log

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
func (it *IStakeRootCompendiumVerifierSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakeRootCompendiumVerifierSet)
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
		it.Event = new(IStakeRootCompendiumVerifierSet)
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
func (it *IStakeRootCompendiumVerifierSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakeRootCompendiumVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakeRootCompendiumVerifierSet represents a VerifierSet event raised by the IStakeRootCompendium contract.
type IStakeRootCompendiumVerifierSet struct {
	NewVerifier common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerifierSet is a free log retrieval operation binding the contract event 0x480b37e3d134e44cb444c9703493c7db564c707cb8a18cecea165f340431da1f.
//
// Solidity: event VerifierSet(address newVerifier)
func (_IStakeRootCompendium *IStakeRootCompendiumFilterer) FilterVerifierSet(opts *bind.FilterOpts) (*IStakeRootCompendiumVerifierSetIterator, error) {

	logs, sub, err := _IStakeRootCompendium.contract.FilterLogs(opts, "VerifierSet")
	if err != nil {
		return nil, err
	}
	return &IStakeRootCompendiumVerifierSetIterator{contract: _IStakeRootCompendium.contract, event: "VerifierSet", logs: logs, sub: sub}, nil
}

// WatchVerifierSet is a free log subscription operation binding the contract event 0x480b37e3d134e44cb444c9703493c7db564c707cb8a18cecea165f340431da1f.
//
// Solidity: event VerifierSet(address newVerifier)
func (_IStakeRootCompendium *IStakeRootCompendiumFilterer) WatchVerifierSet(opts *bind.WatchOpts, sink chan<- *IStakeRootCompendiumVerifierSet) (event.Subscription, error) {

	logs, sub, err := _IStakeRootCompendium.contract.WatchLogs(opts, "VerifierSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakeRootCompendiumVerifierSet)
				if err := _IStakeRootCompendium.contract.UnpackLog(event, "VerifierSet", log); err != nil {
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
func (_IStakeRootCompendium *IStakeRootCompendiumFilterer) ParseVerifierSet(log types.Log) (*IStakeRootCompendiumVerifierSet, error) {
	event := new(IStakeRootCompendiumVerifierSet)
	if err := _IStakeRootCompendium.contract.UnpackLog(event, "VerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
