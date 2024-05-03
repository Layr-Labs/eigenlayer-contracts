// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IEigenPod

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

// BeaconChainProofsStateRootProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsStateRootProof struct {
	BeaconStateRoot [32]byte
	Proof           []byte
}

// BeaconChainProofsWithdrawalProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsWithdrawalProof struct {
	WithdrawalProof                 []byte
	SlotProof                       []byte
	ExecutionPayloadProof           []byte
	TimestampProof                  []byte
	HistoricalSummaryBlockRootProof []byte
	BlockRootIndex                  uint64
	HistoricalSummaryIndex          uint64
	WithdrawalIndex                 uint64
	BlockRoot                       [32]byte
	SlotRoot                        [32]byte
	TimestampRoot                   [32]byte
	ExecutionPayloadRoot            [32]byte
}

// IEigenPodValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IEigenPodValidatorInfo struct {
	ValidatorIndex                   uint64
	RestakedBalanceGwei              uint64
	MostRecentBalanceUpdateTimestamp uint64
	Status                           uint8
}

// IEigenPodMetaData contains all meta data concerning the IEigenPod contract.
var IEigenPodMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activateRestaking\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRestaked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mostRecentWithdrawalTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonBeaconChainETHBalanceWei\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"podOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provenWithdrawal\",\"inputs\":[{\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverTokens\",\"inputs\":[{\"name\":\"tokenList\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"amountsToWithdraw\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validatorPubkeyHashToInfo\",\"inputs\":[{\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPod.ValidatorInfo\",\"components\":[{\"name\":\"validatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"restakedBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"mostRecentBalanceUpdateTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorPubkeyToInfo\",\"inputs\":[{\"name\":\"validatorPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPod.ValidatorInfo\",\"components\":[{\"name\":\"validatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"restakedBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"mostRecentBalanceUpdateTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorStatus\",\"inputs\":[{\"name\":\"validatorPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorStatus\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyAndProcessWithdrawals\",\"inputs\":[{\"name\":\"oracleTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"withdrawalProofs\",\"type\":\"tuple[]\",\"internalType\":\"structBeaconChainProofs.WithdrawalProof[]\",\"components\":[{\"name\":\"withdrawalProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"slotProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"executionPayloadProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timestampProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"historicalSummaryBlockRootProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockRootIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"historicalSummaryIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"withdrawalIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"slotRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestampRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"executionPayloadRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"validatorFieldsProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"validatorFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"},{\"name\":\"withdrawalFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyBalanceUpdates\",\"inputs\":[{\"name\":\"oracleTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validatorIndices\",\"type\":\"uint40[]\",\"internalType\":\"uint40[]\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"validatorFieldsProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"validatorFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyWithdrawalCredentials\",\"inputs\":[{\"name\":\"oracleTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"validatorIndices\",\"type\":\"uint40[]\",\"internalType\":\"uint40[]\"},{\"name\":\"withdrawalCredentialProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"validatorFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawBeforeRestaking\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawNonBeaconChainETHBalanceWei\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountToWithdraw\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawRestakedBeaconChainETH\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawableRestakedExecutionLayerGwei\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EigenPodStaked\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FullWithdrawalRedeemed\",\"inputs\":[{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"withdrawalTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawalAmountGwei\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonBeaconChainETHReceived\",\"inputs\":[{\"name\":\"amountReceived\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonBeaconChainETHWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountWithdrawn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialWithdrawalRedeemed\",\"inputs\":[{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"withdrawalTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"partialWithdrawalAmountGwei\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RestakedBeaconChainETHWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RestakingActivated\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorBalanceUpdated\",\"inputs\":[{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"balanceTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newValidatorBalanceGwei\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRestaked\",\"inputs\":[{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false}]",
}

// IEigenPodABI is the input ABI used to generate the binding from.
// Deprecated: Use IEigenPodMetaData.ABI instead.
var IEigenPodABI = IEigenPodMetaData.ABI

// IEigenPod is an auto generated Go binding around an Ethereum contract.
type IEigenPod struct {
	IEigenPodCaller     // Read-only binding to the contract
	IEigenPodTransactor // Write-only binding to the contract
	IEigenPodFilterer   // Log filterer for contract events
}

// IEigenPodCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEigenPodCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEigenPodTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEigenPodTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEigenPodFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEigenPodFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEigenPodSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEigenPodSession struct {
	Contract     *IEigenPod        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEigenPodCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEigenPodCallerSession struct {
	Contract *IEigenPodCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IEigenPodTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEigenPodTransactorSession struct {
	Contract     *IEigenPodTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IEigenPodRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEigenPodRaw struct {
	Contract *IEigenPod // Generic contract binding to access the raw methods on
}

// IEigenPodCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEigenPodCallerRaw struct {
	Contract *IEigenPodCaller // Generic read-only contract binding to access the raw methods on
}

// IEigenPodTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEigenPodTransactorRaw struct {
	Contract *IEigenPodTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEigenPod creates a new instance of IEigenPod, bound to a specific deployed contract.
func NewIEigenPod(address common.Address, backend bind.ContractBackend) (*IEigenPod, error) {
	contract, err := bindIEigenPod(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEigenPod{IEigenPodCaller: IEigenPodCaller{contract: contract}, IEigenPodTransactor: IEigenPodTransactor{contract: contract}, IEigenPodFilterer: IEigenPodFilterer{contract: contract}}, nil
}

// NewIEigenPodCaller creates a new read-only instance of IEigenPod, bound to a specific deployed contract.
func NewIEigenPodCaller(address common.Address, caller bind.ContractCaller) (*IEigenPodCaller, error) {
	contract, err := bindIEigenPod(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEigenPodCaller{contract: contract}, nil
}

// NewIEigenPodTransactor creates a new write-only instance of IEigenPod, bound to a specific deployed contract.
func NewIEigenPodTransactor(address common.Address, transactor bind.ContractTransactor) (*IEigenPodTransactor, error) {
	contract, err := bindIEigenPod(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEigenPodTransactor{contract: contract}, nil
}

// NewIEigenPodFilterer creates a new log filterer instance of IEigenPod, bound to a specific deployed contract.
func NewIEigenPodFilterer(address common.Address, filterer bind.ContractFilterer) (*IEigenPodFilterer, error) {
	contract, err := bindIEigenPod(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEigenPodFilterer{contract: contract}, nil
}

// bindIEigenPod binds a generic wrapper to an already deployed contract.
func bindIEigenPod(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEigenPodMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEigenPod *IEigenPodRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEigenPod.Contract.IEigenPodCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEigenPod *IEigenPodRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEigenPod.Contract.IEigenPodTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEigenPod *IEigenPodRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEigenPod.Contract.IEigenPodTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEigenPod *IEigenPodCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEigenPod.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEigenPod *IEigenPodTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEigenPod.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEigenPod *IEigenPodTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEigenPod.Contract.contract.Transact(opts, method, params...)
}

// MAXRESTAKEDBALANCEGWEIPERVALIDATOR is a free data retrieval call binding the contract method 0x1d905d5c.
//
// Solidity: function MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR() view returns(uint64)
func (_IEigenPod *IEigenPodCaller) MAXRESTAKEDBALANCEGWEIPERVALIDATOR(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXRESTAKEDBALANCEGWEIPERVALIDATOR is a free data retrieval call binding the contract method 0x1d905d5c.
//
// Solidity: function MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR() view returns(uint64)
func (_IEigenPod *IEigenPodSession) MAXRESTAKEDBALANCEGWEIPERVALIDATOR() (uint64, error) {
	return _IEigenPod.Contract.MAXRESTAKEDBALANCEGWEIPERVALIDATOR(&_IEigenPod.CallOpts)
}

// MAXRESTAKEDBALANCEGWEIPERVALIDATOR is a free data retrieval call binding the contract method 0x1d905d5c.
//
// Solidity: function MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR() view returns(uint64)
func (_IEigenPod *IEigenPodCallerSession) MAXRESTAKEDBALANCEGWEIPERVALIDATOR() (uint64, error) {
	return _IEigenPod.Contract.MAXRESTAKEDBALANCEGWEIPERVALIDATOR(&_IEigenPod.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_IEigenPod *IEigenPodCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_IEigenPod *IEigenPodSession) EigenPodManager() (common.Address, error) {
	return _IEigenPod.Contract.EigenPodManager(&_IEigenPod.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_IEigenPod *IEigenPodCallerSession) EigenPodManager() (common.Address, error) {
	return _IEigenPod.Contract.EigenPodManager(&_IEigenPod.CallOpts)
}

// HasRestaked is a free data retrieval call binding the contract method 0x3106ab53.
//
// Solidity: function hasRestaked() view returns(bool)
func (_IEigenPod *IEigenPodCaller) HasRestaked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "hasRestaked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRestaked is a free data retrieval call binding the contract method 0x3106ab53.
//
// Solidity: function hasRestaked() view returns(bool)
func (_IEigenPod *IEigenPodSession) HasRestaked() (bool, error) {
	return _IEigenPod.Contract.HasRestaked(&_IEigenPod.CallOpts)
}

// HasRestaked is a free data retrieval call binding the contract method 0x3106ab53.
//
// Solidity: function hasRestaked() view returns(bool)
func (_IEigenPod *IEigenPodCallerSession) HasRestaked() (bool, error) {
	return _IEigenPod.Contract.HasRestaked(&_IEigenPod.CallOpts)
}

// MostRecentWithdrawalTimestamp is a free data retrieval call binding the contract method 0x87e0d289.
//
// Solidity: function mostRecentWithdrawalTimestamp() view returns(uint64)
func (_IEigenPod *IEigenPodCaller) MostRecentWithdrawalTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "mostRecentWithdrawalTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MostRecentWithdrawalTimestamp is a free data retrieval call binding the contract method 0x87e0d289.
//
// Solidity: function mostRecentWithdrawalTimestamp() view returns(uint64)
func (_IEigenPod *IEigenPodSession) MostRecentWithdrawalTimestamp() (uint64, error) {
	return _IEigenPod.Contract.MostRecentWithdrawalTimestamp(&_IEigenPod.CallOpts)
}

// MostRecentWithdrawalTimestamp is a free data retrieval call binding the contract method 0x87e0d289.
//
// Solidity: function mostRecentWithdrawalTimestamp() view returns(uint64)
func (_IEigenPod *IEigenPodCallerSession) MostRecentWithdrawalTimestamp() (uint64, error) {
	return _IEigenPod.Contract.MostRecentWithdrawalTimestamp(&_IEigenPod.CallOpts)
}

// NonBeaconChainETHBalanceWei is a free data retrieval call binding the contract method 0xfe80b087.
//
// Solidity: function nonBeaconChainETHBalanceWei() view returns(uint256)
func (_IEigenPod *IEigenPodCaller) NonBeaconChainETHBalanceWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "nonBeaconChainETHBalanceWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonBeaconChainETHBalanceWei is a free data retrieval call binding the contract method 0xfe80b087.
//
// Solidity: function nonBeaconChainETHBalanceWei() view returns(uint256)
func (_IEigenPod *IEigenPodSession) NonBeaconChainETHBalanceWei() (*big.Int, error) {
	return _IEigenPod.Contract.NonBeaconChainETHBalanceWei(&_IEigenPod.CallOpts)
}

// NonBeaconChainETHBalanceWei is a free data retrieval call binding the contract method 0xfe80b087.
//
// Solidity: function nonBeaconChainETHBalanceWei() view returns(uint256)
func (_IEigenPod *IEigenPodCallerSession) NonBeaconChainETHBalanceWei() (*big.Int, error) {
	return _IEigenPod.Contract.NonBeaconChainETHBalanceWei(&_IEigenPod.CallOpts)
}

// PodOwner is a free data retrieval call binding the contract method 0x0b18ff66.
//
// Solidity: function podOwner() view returns(address)
func (_IEigenPod *IEigenPodCaller) PodOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "podOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PodOwner is a free data retrieval call binding the contract method 0x0b18ff66.
//
// Solidity: function podOwner() view returns(address)
func (_IEigenPod *IEigenPodSession) PodOwner() (common.Address, error) {
	return _IEigenPod.Contract.PodOwner(&_IEigenPod.CallOpts)
}

// PodOwner is a free data retrieval call binding the contract method 0x0b18ff66.
//
// Solidity: function podOwner() view returns(address)
func (_IEigenPod *IEigenPodCallerSession) PodOwner() (common.Address, error) {
	return _IEigenPod.Contract.PodOwner(&_IEigenPod.CallOpts)
}

// ProvenWithdrawal is a free data retrieval call binding the contract method 0x34bea20a.
//
// Solidity: function provenWithdrawal(bytes32 validatorPubkeyHash, uint64 slot) view returns(bool)
func (_IEigenPod *IEigenPodCaller) ProvenWithdrawal(opts *bind.CallOpts, validatorPubkeyHash [32]byte, slot uint64) (bool, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "provenWithdrawal", validatorPubkeyHash, slot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProvenWithdrawal is a free data retrieval call binding the contract method 0x34bea20a.
//
// Solidity: function provenWithdrawal(bytes32 validatorPubkeyHash, uint64 slot) view returns(bool)
func (_IEigenPod *IEigenPodSession) ProvenWithdrawal(validatorPubkeyHash [32]byte, slot uint64) (bool, error) {
	return _IEigenPod.Contract.ProvenWithdrawal(&_IEigenPod.CallOpts, validatorPubkeyHash, slot)
}

// ProvenWithdrawal is a free data retrieval call binding the contract method 0x34bea20a.
//
// Solidity: function provenWithdrawal(bytes32 validatorPubkeyHash, uint64 slot) view returns(bool)
func (_IEigenPod *IEigenPodCallerSession) ProvenWithdrawal(validatorPubkeyHash [32]byte, slot uint64) (bool, error) {
	return _IEigenPod.Contract.ProvenWithdrawal(&_IEigenPod.CallOpts, validatorPubkeyHash, slot)
}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_IEigenPod *IEigenPodCaller) ValidatorPubkeyHashToInfo(opts *bind.CallOpts, validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "validatorPubkeyHashToInfo", validatorPubkeyHash)

	if err != nil {
		return *new(IEigenPodValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodValidatorInfo)).(*IEigenPodValidatorInfo)

	return out0, err

}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_IEigenPod *IEigenPodSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error) {
	return _IEigenPod.Contract.ValidatorPubkeyHashToInfo(&_IEigenPod.CallOpts, validatorPubkeyHash)
}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_IEigenPod *IEigenPodCallerSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error) {
	return _IEigenPod.Contract.ValidatorPubkeyHashToInfo(&_IEigenPod.CallOpts, validatorPubkeyHash)
}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_IEigenPod *IEigenPodCaller) ValidatorPubkeyToInfo(opts *bind.CallOpts, validatorPubkey []byte) (IEigenPodValidatorInfo, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "validatorPubkeyToInfo", validatorPubkey)

	if err != nil {
		return *new(IEigenPodValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodValidatorInfo)).(*IEigenPodValidatorInfo)

	return out0, err

}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_IEigenPod *IEigenPodSession) ValidatorPubkeyToInfo(validatorPubkey []byte) (IEigenPodValidatorInfo, error) {
	return _IEigenPod.Contract.ValidatorPubkeyToInfo(&_IEigenPod.CallOpts, validatorPubkey)
}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_IEigenPod *IEigenPodCallerSession) ValidatorPubkeyToInfo(validatorPubkey []byte) (IEigenPodValidatorInfo, error) {
	return _IEigenPod.Contract.ValidatorPubkeyToInfo(&_IEigenPod.CallOpts, validatorPubkey)
}

// ValidatorStatus is a free data retrieval call binding the contract method 0x58eaee79.
//
// Solidity: function validatorStatus(bytes validatorPubkey) view returns(uint8)
func (_IEigenPod *IEigenPodCaller) ValidatorStatus(opts *bind.CallOpts, validatorPubkey []byte) (uint8, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "validatorStatus", validatorPubkey)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ValidatorStatus is a free data retrieval call binding the contract method 0x58eaee79.
//
// Solidity: function validatorStatus(bytes validatorPubkey) view returns(uint8)
func (_IEigenPod *IEigenPodSession) ValidatorStatus(validatorPubkey []byte) (uint8, error) {
	return _IEigenPod.Contract.ValidatorStatus(&_IEigenPod.CallOpts, validatorPubkey)
}

// ValidatorStatus is a free data retrieval call binding the contract method 0x58eaee79.
//
// Solidity: function validatorStatus(bytes validatorPubkey) view returns(uint8)
func (_IEigenPod *IEigenPodCallerSession) ValidatorStatus(validatorPubkey []byte) (uint8, error) {
	return _IEigenPod.Contract.ValidatorStatus(&_IEigenPod.CallOpts, validatorPubkey)
}

// ValidatorStatus0 is a free data retrieval call binding the contract method 0x7439841f.
//
// Solidity: function validatorStatus(bytes32 pubkeyHash) view returns(uint8)
func (_IEigenPod *IEigenPodCaller) ValidatorStatus0(opts *bind.CallOpts, pubkeyHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "validatorStatus0", pubkeyHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ValidatorStatus0 is a free data retrieval call binding the contract method 0x7439841f.
//
// Solidity: function validatorStatus(bytes32 pubkeyHash) view returns(uint8)
func (_IEigenPod *IEigenPodSession) ValidatorStatus0(pubkeyHash [32]byte) (uint8, error) {
	return _IEigenPod.Contract.ValidatorStatus0(&_IEigenPod.CallOpts, pubkeyHash)
}

// ValidatorStatus0 is a free data retrieval call binding the contract method 0x7439841f.
//
// Solidity: function validatorStatus(bytes32 pubkeyHash) view returns(uint8)
func (_IEigenPod *IEigenPodCallerSession) ValidatorStatus0(pubkeyHash [32]byte) (uint8, error) {
	return _IEigenPod.Contract.ValidatorStatus0(&_IEigenPod.CallOpts, pubkeyHash)
}

// WithdrawableRestakedExecutionLayerGwei is a free data retrieval call binding the contract method 0x3474aa16.
//
// Solidity: function withdrawableRestakedExecutionLayerGwei() view returns(uint64)
func (_IEigenPod *IEigenPodCaller) WithdrawableRestakedExecutionLayerGwei(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "withdrawableRestakedExecutionLayerGwei")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// WithdrawableRestakedExecutionLayerGwei is a free data retrieval call binding the contract method 0x3474aa16.
//
// Solidity: function withdrawableRestakedExecutionLayerGwei() view returns(uint64)
func (_IEigenPod *IEigenPodSession) WithdrawableRestakedExecutionLayerGwei() (uint64, error) {
	return _IEigenPod.Contract.WithdrawableRestakedExecutionLayerGwei(&_IEigenPod.CallOpts)
}

// WithdrawableRestakedExecutionLayerGwei is a free data retrieval call binding the contract method 0x3474aa16.
//
// Solidity: function withdrawableRestakedExecutionLayerGwei() view returns(uint64)
func (_IEigenPod *IEigenPodCallerSession) WithdrawableRestakedExecutionLayerGwei() (uint64, error) {
	return _IEigenPod.Contract.WithdrawableRestakedExecutionLayerGwei(&_IEigenPod.CallOpts)
}

// ActivateRestaking is a paid mutator transaction binding the contract method 0x0cd4649e.
//
// Solidity: function activateRestaking() returns()
func (_IEigenPod *IEigenPodTransactor) ActivateRestaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEigenPod.contract.Transact(opts, "activateRestaking")
}

// ActivateRestaking is a paid mutator transaction binding the contract method 0x0cd4649e.
//
// Solidity: function activateRestaking() returns()
func (_IEigenPod *IEigenPodSession) ActivateRestaking() (*types.Transaction, error) {
	return _IEigenPod.Contract.ActivateRestaking(&_IEigenPod.TransactOpts)
}

// ActivateRestaking is a paid mutator transaction binding the contract method 0x0cd4649e.
//
// Solidity: function activateRestaking() returns()
func (_IEigenPod *IEigenPodTransactorSession) ActivateRestaking() (*types.Transaction, error) {
	return _IEigenPod.Contract.ActivateRestaking(&_IEigenPod.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_IEigenPod *IEigenPodTransactor) Initialize(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _IEigenPod.contract.Transact(opts, "initialize", owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_IEigenPod *IEigenPodSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _IEigenPod.Contract.Initialize(&_IEigenPod.TransactOpts, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_IEigenPod *IEigenPodTransactorSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _IEigenPod.Contract.Initialize(&_IEigenPod.TransactOpts, owner)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_IEigenPod *IEigenPodTransactor) RecoverTokens(opts *bind.TransactOpts, tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IEigenPod.contract.Transact(opts, "recoverTokens", tokenList, amountsToWithdraw, recipient)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_IEigenPod *IEigenPodSession) RecoverTokens(tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IEigenPod.Contract.RecoverTokens(&_IEigenPod.TransactOpts, tokenList, amountsToWithdraw, recipient)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_IEigenPod *IEigenPodTransactorSession) RecoverTokens(tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _IEigenPod.Contract.RecoverTokens(&_IEigenPod.TransactOpts, tokenList, amountsToWithdraw, recipient)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_IEigenPod *IEigenPodTransactor) Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _IEigenPod.contract.Transact(opts, "stake", pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_IEigenPod *IEigenPodSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _IEigenPod.Contract.Stake(&_IEigenPod.TransactOpts, pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_IEigenPod *IEigenPodTransactorSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _IEigenPod.Contract.Stake(&_IEigenPod.TransactOpts, pubkey, signature, depositDataRoot)
}

// VerifyAndProcessWithdrawals is a paid mutator transaction binding the contract method 0xe251ef52.
//
// Solidity: function verifyAndProcessWithdrawals(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, (bytes,bytes,bytes,bytes,bytes,uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32)[] withdrawalProofs, bytes[] validatorFieldsProofs, bytes32[][] validatorFields, bytes32[][] withdrawalFields) returns()
func (_IEigenPod *IEigenPodTransactor) VerifyAndProcessWithdrawals(opts *bind.TransactOpts, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, withdrawalProofs []BeaconChainProofsWithdrawalProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte, withdrawalFields [][][32]byte) (*types.Transaction, error) {
	return _IEigenPod.contract.Transact(opts, "verifyAndProcessWithdrawals", oracleTimestamp, stateRootProof, withdrawalProofs, validatorFieldsProofs, validatorFields, withdrawalFields)
}

// VerifyAndProcessWithdrawals is a paid mutator transaction binding the contract method 0xe251ef52.
//
// Solidity: function verifyAndProcessWithdrawals(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, (bytes,bytes,bytes,bytes,bytes,uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32)[] withdrawalProofs, bytes[] validatorFieldsProofs, bytes32[][] validatorFields, bytes32[][] withdrawalFields) returns()
func (_IEigenPod *IEigenPodSession) VerifyAndProcessWithdrawals(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, withdrawalProofs []BeaconChainProofsWithdrawalProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte, withdrawalFields [][][32]byte) (*types.Transaction, error) {
	return _IEigenPod.Contract.VerifyAndProcessWithdrawals(&_IEigenPod.TransactOpts, oracleTimestamp, stateRootProof, withdrawalProofs, validatorFieldsProofs, validatorFields, withdrawalFields)
}

// VerifyAndProcessWithdrawals is a paid mutator transaction binding the contract method 0xe251ef52.
//
// Solidity: function verifyAndProcessWithdrawals(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, (bytes,bytes,bytes,bytes,bytes,uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32)[] withdrawalProofs, bytes[] validatorFieldsProofs, bytes32[][] validatorFields, bytes32[][] withdrawalFields) returns()
func (_IEigenPod *IEigenPodTransactorSession) VerifyAndProcessWithdrawals(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, withdrawalProofs []BeaconChainProofsWithdrawalProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte, withdrawalFields [][][32]byte) (*types.Transaction, error) {
	return _IEigenPod.Contract.VerifyAndProcessWithdrawals(&_IEigenPod.TransactOpts, oracleTimestamp, stateRootProof, withdrawalProofs, validatorFieldsProofs, validatorFields, withdrawalFields)
}

// VerifyBalanceUpdates is a paid mutator transaction binding the contract method 0xa50600f4.
//
// Solidity: function verifyBalanceUpdates(uint64 oracleTimestamp, uint40[] validatorIndices, (bytes32,bytes) stateRootProof, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_IEigenPod *IEigenPodTransactor) VerifyBalanceUpdates(opts *bind.TransactOpts, oracleTimestamp uint64, validatorIndices []*big.Int, stateRootProof BeaconChainProofsStateRootProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _IEigenPod.contract.Transact(opts, "verifyBalanceUpdates", oracleTimestamp, validatorIndices, stateRootProof, validatorFieldsProofs, validatorFields)
}

// VerifyBalanceUpdates is a paid mutator transaction binding the contract method 0xa50600f4.
//
// Solidity: function verifyBalanceUpdates(uint64 oracleTimestamp, uint40[] validatorIndices, (bytes32,bytes) stateRootProof, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_IEigenPod *IEigenPodSession) VerifyBalanceUpdates(oracleTimestamp uint64, validatorIndices []*big.Int, stateRootProof BeaconChainProofsStateRootProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _IEigenPod.Contract.VerifyBalanceUpdates(&_IEigenPod.TransactOpts, oracleTimestamp, validatorIndices, stateRootProof, validatorFieldsProofs, validatorFields)
}

// VerifyBalanceUpdates is a paid mutator transaction binding the contract method 0xa50600f4.
//
// Solidity: function verifyBalanceUpdates(uint64 oracleTimestamp, uint40[] validatorIndices, (bytes32,bytes) stateRootProof, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_IEigenPod *IEigenPodTransactorSession) VerifyBalanceUpdates(oracleTimestamp uint64, validatorIndices []*big.Int, stateRootProof BeaconChainProofsStateRootProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _IEigenPod.Contract.VerifyBalanceUpdates(&_IEigenPod.TransactOpts, oracleTimestamp, validatorIndices, stateRootProof, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] withdrawalCredentialProofs, bytes32[][] validatorFields) returns()
func (_IEigenPod *IEigenPodTransactor) VerifyWithdrawalCredentials(opts *bind.TransactOpts, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, withdrawalCredentialProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _IEigenPod.contract.Transact(opts, "verifyWithdrawalCredentials", oracleTimestamp, stateRootProof, validatorIndices, withdrawalCredentialProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] withdrawalCredentialProofs, bytes32[][] validatorFields) returns()
func (_IEigenPod *IEigenPodSession) VerifyWithdrawalCredentials(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, withdrawalCredentialProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _IEigenPod.Contract.VerifyWithdrawalCredentials(&_IEigenPod.TransactOpts, oracleTimestamp, stateRootProof, validatorIndices, withdrawalCredentialProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] withdrawalCredentialProofs, bytes32[][] validatorFields) returns()
func (_IEigenPod *IEigenPodTransactorSession) VerifyWithdrawalCredentials(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, withdrawalCredentialProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _IEigenPod.Contract.VerifyWithdrawalCredentials(&_IEigenPod.TransactOpts, oracleTimestamp, stateRootProof, validatorIndices, withdrawalCredentialProofs, validatorFields)
}

// WithdrawBeforeRestaking is a paid mutator transaction binding the contract method 0xbaa7145a.
//
// Solidity: function withdrawBeforeRestaking() returns()
func (_IEigenPod *IEigenPodTransactor) WithdrawBeforeRestaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEigenPod.contract.Transact(opts, "withdrawBeforeRestaking")
}

// WithdrawBeforeRestaking is a paid mutator transaction binding the contract method 0xbaa7145a.
//
// Solidity: function withdrawBeforeRestaking() returns()
func (_IEigenPod *IEigenPodSession) WithdrawBeforeRestaking() (*types.Transaction, error) {
	return _IEigenPod.Contract.WithdrawBeforeRestaking(&_IEigenPod.TransactOpts)
}

// WithdrawBeforeRestaking is a paid mutator transaction binding the contract method 0xbaa7145a.
//
// Solidity: function withdrawBeforeRestaking() returns()
func (_IEigenPod *IEigenPodTransactorSession) WithdrawBeforeRestaking() (*types.Transaction, error) {
	return _IEigenPod.Contract.WithdrawBeforeRestaking(&_IEigenPod.TransactOpts)
}

// WithdrawNonBeaconChainETHBalanceWei is a paid mutator transaction binding the contract method 0xe2c83445.
//
// Solidity: function withdrawNonBeaconChainETHBalanceWei(address recipient, uint256 amountToWithdraw) returns()
func (_IEigenPod *IEigenPodTransactor) WithdrawNonBeaconChainETHBalanceWei(opts *bind.TransactOpts, recipient common.Address, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _IEigenPod.contract.Transact(opts, "withdrawNonBeaconChainETHBalanceWei", recipient, amountToWithdraw)
}

// WithdrawNonBeaconChainETHBalanceWei is a paid mutator transaction binding the contract method 0xe2c83445.
//
// Solidity: function withdrawNonBeaconChainETHBalanceWei(address recipient, uint256 amountToWithdraw) returns()
func (_IEigenPod *IEigenPodSession) WithdrawNonBeaconChainETHBalanceWei(recipient common.Address, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _IEigenPod.Contract.WithdrawNonBeaconChainETHBalanceWei(&_IEigenPod.TransactOpts, recipient, amountToWithdraw)
}

// WithdrawNonBeaconChainETHBalanceWei is a paid mutator transaction binding the contract method 0xe2c83445.
//
// Solidity: function withdrawNonBeaconChainETHBalanceWei(address recipient, uint256 amountToWithdraw) returns()
func (_IEigenPod *IEigenPodTransactorSession) WithdrawNonBeaconChainETHBalanceWei(recipient common.Address, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _IEigenPod.Contract.WithdrawNonBeaconChainETHBalanceWei(&_IEigenPod.TransactOpts, recipient, amountToWithdraw)
}

// WithdrawRestakedBeaconChainETH is a paid mutator transaction binding the contract method 0xc4907442.
//
// Solidity: function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) returns()
func (_IEigenPod *IEigenPodTransactor) WithdrawRestakedBeaconChainETH(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IEigenPod.contract.Transact(opts, "withdrawRestakedBeaconChainETH", recipient, amount)
}

// WithdrawRestakedBeaconChainETH is a paid mutator transaction binding the contract method 0xc4907442.
//
// Solidity: function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) returns()
func (_IEigenPod *IEigenPodSession) WithdrawRestakedBeaconChainETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IEigenPod.Contract.WithdrawRestakedBeaconChainETH(&_IEigenPod.TransactOpts, recipient, amount)
}

// WithdrawRestakedBeaconChainETH is a paid mutator transaction binding the contract method 0xc4907442.
//
// Solidity: function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) returns()
func (_IEigenPod *IEigenPodTransactorSession) WithdrawRestakedBeaconChainETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IEigenPod.Contract.WithdrawRestakedBeaconChainETH(&_IEigenPod.TransactOpts, recipient, amount)
}

// IEigenPodEigenPodStakedIterator is returned from FilterEigenPodStaked and is used to iterate over the raw logs and unpacked data for EigenPodStaked events raised by the IEigenPod contract.
type IEigenPodEigenPodStakedIterator struct {
	Event *IEigenPodEigenPodStaked // Event containing the contract specifics and raw log

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
func (it *IEigenPodEigenPodStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodEigenPodStaked)
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
		it.Event = new(IEigenPodEigenPodStaked)
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
func (it *IEigenPodEigenPodStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodEigenPodStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodEigenPodStaked represents a EigenPodStaked event raised by the IEigenPod contract.
type IEigenPodEigenPodStaked struct {
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEigenPodStaked is a free log retrieval operation binding the contract event 0x606865b7934a25d4aed43f6cdb426403353fa4b3009c4d228407474581b01e23.
//
// Solidity: event EigenPodStaked(bytes pubkey)
func (_IEigenPod *IEigenPodFilterer) FilterEigenPodStaked(opts *bind.FilterOpts) (*IEigenPodEigenPodStakedIterator, error) {

	logs, sub, err := _IEigenPod.contract.FilterLogs(opts, "EigenPodStaked")
	if err != nil {
		return nil, err
	}
	return &IEigenPodEigenPodStakedIterator{contract: _IEigenPod.contract, event: "EigenPodStaked", logs: logs, sub: sub}, nil
}

// WatchEigenPodStaked is a free log subscription operation binding the contract event 0x606865b7934a25d4aed43f6cdb426403353fa4b3009c4d228407474581b01e23.
//
// Solidity: event EigenPodStaked(bytes pubkey)
func (_IEigenPod *IEigenPodFilterer) WatchEigenPodStaked(opts *bind.WatchOpts, sink chan<- *IEigenPodEigenPodStaked) (event.Subscription, error) {

	logs, sub, err := _IEigenPod.contract.WatchLogs(opts, "EigenPodStaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodEigenPodStaked)
				if err := _IEigenPod.contract.UnpackLog(event, "EigenPodStaked", log); err != nil {
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

// ParseEigenPodStaked is a log parse operation binding the contract event 0x606865b7934a25d4aed43f6cdb426403353fa4b3009c4d228407474581b01e23.
//
// Solidity: event EigenPodStaked(bytes pubkey)
func (_IEigenPod *IEigenPodFilterer) ParseEigenPodStaked(log types.Log) (*IEigenPodEigenPodStaked, error) {
	event := new(IEigenPodEigenPodStaked)
	if err := _IEigenPod.contract.UnpackLog(event, "EigenPodStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodFullWithdrawalRedeemedIterator is returned from FilterFullWithdrawalRedeemed and is used to iterate over the raw logs and unpacked data for FullWithdrawalRedeemed events raised by the IEigenPod contract.
type IEigenPodFullWithdrawalRedeemedIterator struct {
	Event *IEigenPodFullWithdrawalRedeemed // Event containing the contract specifics and raw log

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
func (it *IEigenPodFullWithdrawalRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodFullWithdrawalRedeemed)
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
		it.Event = new(IEigenPodFullWithdrawalRedeemed)
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
func (it *IEigenPodFullWithdrawalRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodFullWithdrawalRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodFullWithdrawalRedeemed represents a FullWithdrawalRedeemed event raised by the IEigenPod contract.
type IEigenPodFullWithdrawalRedeemed struct {
	ValidatorIndex       *big.Int
	WithdrawalTimestamp  uint64
	Recipient            common.Address
	WithdrawalAmountGwei uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterFullWithdrawalRedeemed is a free log retrieval operation binding the contract event 0xb76a93bb649ece524688f1a01d184e0bbebcda58eae80c28a898bec3fb5a0963.
//
// Solidity: event FullWithdrawalRedeemed(uint40 validatorIndex, uint64 withdrawalTimestamp, address indexed recipient, uint64 withdrawalAmountGwei)
func (_IEigenPod *IEigenPodFilterer) FilterFullWithdrawalRedeemed(opts *bind.FilterOpts, recipient []common.Address) (*IEigenPodFullWithdrawalRedeemedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IEigenPod.contract.FilterLogs(opts, "FullWithdrawalRedeemed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodFullWithdrawalRedeemedIterator{contract: _IEigenPod.contract, event: "FullWithdrawalRedeemed", logs: logs, sub: sub}, nil
}

// WatchFullWithdrawalRedeemed is a free log subscription operation binding the contract event 0xb76a93bb649ece524688f1a01d184e0bbebcda58eae80c28a898bec3fb5a0963.
//
// Solidity: event FullWithdrawalRedeemed(uint40 validatorIndex, uint64 withdrawalTimestamp, address indexed recipient, uint64 withdrawalAmountGwei)
func (_IEigenPod *IEigenPodFilterer) WatchFullWithdrawalRedeemed(opts *bind.WatchOpts, sink chan<- *IEigenPodFullWithdrawalRedeemed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IEigenPod.contract.WatchLogs(opts, "FullWithdrawalRedeemed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodFullWithdrawalRedeemed)
				if err := _IEigenPod.contract.UnpackLog(event, "FullWithdrawalRedeemed", log); err != nil {
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

// ParseFullWithdrawalRedeemed is a log parse operation binding the contract event 0xb76a93bb649ece524688f1a01d184e0bbebcda58eae80c28a898bec3fb5a0963.
//
// Solidity: event FullWithdrawalRedeemed(uint40 validatorIndex, uint64 withdrawalTimestamp, address indexed recipient, uint64 withdrawalAmountGwei)
func (_IEigenPod *IEigenPodFilterer) ParseFullWithdrawalRedeemed(log types.Log) (*IEigenPodFullWithdrawalRedeemed, error) {
	event := new(IEigenPodFullWithdrawalRedeemed)
	if err := _IEigenPod.contract.UnpackLog(event, "FullWithdrawalRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodNonBeaconChainETHReceivedIterator is returned from FilterNonBeaconChainETHReceived and is used to iterate over the raw logs and unpacked data for NonBeaconChainETHReceived events raised by the IEigenPod contract.
type IEigenPodNonBeaconChainETHReceivedIterator struct {
	Event *IEigenPodNonBeaconChainETHReceived // Event containing the contract specifics and raw log

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
func (it *IEigenPodNonBeaconChainETHReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodNonBeaconChainETHReceived)
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
		it.Event = new(IEigenPodNonBeaconChainETHReceived)
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
func (it *IEigenPodNonBeaconChainETHReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodNonBeaconChainETHReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodNonBeaconChainETHReceived represents a NonBeaconChainETHReceived event raised by the IEigenPod contract.
type IEigenPodNonBeaconChainETHReceived struct {
	AmountReceived *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNonBeaconChainETHReceived is a free log retrieval operation binding the contract event 0x6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49.
//
// Solidity: event NonBeaconChainETHReceived(uint256 amountReceived)
func (_IEigenPod *IEigenPodFilterer) FilterNonBeaconChainETHReceived(opts *bind.FilterOpts) (*IEigenPodNonBeaconChainETHReceivedIterator, error) {

	logs, sub, err := _IEigenPod.contract.FilterLogs(opts, "NonBeaconChainETHReceived")
	if err != nil {
		return nil, err
	}
	return &IEigenPodNonBeaconChainETHReceivedIterator{contract: _IEigenPod.contract, event: "NonBeaconChainETHReceived", logs: logs, sub: sub}, nil
}

// WatchNonBeaconChainETHReceived is a free log subscription operation binding the contract event 0x6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49.
//
// Solidity: event NonBeaconChainETHReceived(uint256 amountReceived)
func (_IEigenPod *IEigenPodFilterer) WatchNonBeaconChainETHReceived(opts *bind.WatchOpts, sink chan<- *IEigenPodNonBeaconChainETHReceived) (event.Subscription, error) {

	logs, sub, err := _IEigenPod.contract.WatchLogs(opts, "NonBeaconChainETHReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodNonBeaconChainETHReceived)
				if err := _IEigenPod.contract.UnpackLog(event, "NonBeaconChainETHReceived", log); err != nil {
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

// ParseNonBeaconChainETHReceived is a log parse operation binding the contract event 0x6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49.
//
// Solidity: event NonBeaconChainETHReceived(uint256 amountReceived)
func (_IEigenPod *IEigenPodFilterer) ParseNonBeaconChainETHReceived(log types.Log) (*IEigenPodNonBeaconChainETHReceived, error) {
	event := new(IEigenPodNonBeaconChainETHReceived)
	if err := _IEigenPod.contract.UnpackLog(event, "NonBeaconChainETHReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodNonBeaconChainETHWithdrawnIterator is returned from FilterNonBeaconChainETHWithdrawn and is used to iterate over the raw logs and unpacked data for NonBeaconChainETHWithdrawn events raised by the IEigenPod contract.
type IEigenPodNonBeaconChainETHWithdrawnIterator struct {
	Event *IEigenPodNonBeaconChainETHWithdrawn // Event containing the contract specifics and raw log

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
func (it *IEigenPodNonBeaconChainETHWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodNonBeaconChainETHWithdrawn)
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
		it.Event = new(IEigenPodNonBeaconChainETHWithdrawn)
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
func (it *IEigenPodNonBeaconChainETHWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodNonBeaconChainETHWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodNonBeaconChainETHWithdrawn represents a NonBeaconChainETHWithdrawn event raised by the IEigenPod contract.
type IEigenPodNonBeaconChainETHWithdrawn struct {
	Recipient       common.Address
	AmountWithdrawn *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNonBeaconChainETHWithdrawn is a free log retrieval operation binding the contract event 0x30420aacd028abb3c1fd03aba253ae725d6ddd52d16c9ac4cb5742cd43f53096.
//
// Solidity: event NonBeaconChainETHWithdrawn(address indexed recipient, uint256 amountWithdrawn)
func (_IEigenPod *IEigenPodFilterer) FilterNonBeaconChainETHWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*IEigenPodNonBeaconChainETHWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IEigenPod.contract.FilterLogs(opts, "NonBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodNonBeaconChainETHWithdrawnIterator{contract: _IEigenPod.contract, event: "NonBeaconChainETHWithdrawn", logs: logs, sub: sub}, nil
}

// WatchNonBeaconChainETHWithdrawn is a free log subscription operation binding the contract event 0x30420aacd028abb3c1fd03aba253ae725d6ddd52d16c9ac4cb5742cd43f53096.
//
// Solidity: event NonBeaconChainETHWithdrawn(address indexed recipient, uint256 amountWithdrawn)
func (_IEigenPod *IEigenPodFilterer) WatchNonBeaconChainETHWithdrawn(opts *bind.WatchOpts, sink chan<- *IEigenPodNonBeaconChainETHWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IEigenPod.contract.WatchLogs(opts, "NonBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodNonBeaconChainETHWithdrawn)
				if err := _IEigenPod.contract.UnpackLog(event, "NonBeaconChainETHWithdrawn", log); err != nil {
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

// ParseNonBeaconChainETHWithdrawn is a log parse operation binding the contract event 0x30420aacd028abb3c1fd03aba253ae725d6ddd52d16c9ac4cb5742cd43f53096.
//
// Solidity: event NonBeaconChainETHWithdrawn(address indexed recipient, uint256 amountWithdrawn)
func (_IEigenPod *IEigenPodFilterer) ParseNonBeaconChainETHWithdrawn(log types.Log) (*IEigenPodNonBeaconChainETHWithdrawn, error) {
	event := new(IEigenPodNonBeaconChainETHWithdrawn)
	if err := _IEigenPod.contract.UnpackLog(event, "NonBeaconChainETHWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodPartialWithdrawalRedeemedIterator is returned from FilterPartialWithdrawalRedeemed and is used to iterate over the raw logs and unpacked data for PartialWithdrawalRedeemed events raised by the IEigenPod contract.
type IEigenPodPartialWithdrawalRedeemedIterator struct {
	Event *IEigenPodPartialWithdrawalRedeemed // Event containing the contract specifics and raw log

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
func (it *IEigenPodPartialWithdrawalRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodPartialWithdrawalRedeemed)
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
		it.Event = new(IEigenPodPartialWithdrawalRedeemed)
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
func (it *IEigenPodPartialWithdrawalRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodPartialWithdrawalRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodPartialWithdrawalRedeemed represents a PartialWithdrawalRedeemed event raised by the IEigenPod contract.
type IEigenPodPartialWithdrawalRedeemed struct {
	ValidatorIndex              *big.Int
	WithdrawalTimestamp         uint64
	Recipient                   common.Address
	PartialWithdrawalAmountGwei uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterPartialWithdrawalRedeemed is a free log retrieval operation binding the contract event 0x8a7335714231dbd551aaba6314f4a97a14c201e53a3e25e1140325cdf67d7a4e.
//
// Solidity: event PartialWithdrawalRedeemed(uint40 validatorIndex, uint64 withdrawalTimestamp, address indexed recipient, uint64 partialWithdrawalAmountGwei)
func (_IEigenPod *IEigenPodFilterer) FilterPartialWithdrawalRedeemed(opts *bind.FilterOpts, recipient []common.Address) (*IEigenPodPartialWithdrawalRedeemedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IEigenPod.contract.FilterLogs(opts, "PartialWithdrawalRedeemed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodPartialWithdrawalRedeemedIterator{contract: _IEigenPod.contract, event: "PartialWithdrawalRedeemed", logs: logs, sub: sub}, nil
}

// WatchPartialWithdrawalRedeemed is a free log subscription operation binding the contract event 0x8a7335714231dbd551aaba6314f4a97a14c201e53a3e25e1140325cdf67d7a4e.
//
// Solidity: event PartialWithdrawalRedeemed(uint40 validatorIndex, uint64 withdrawalTimestamp, address indexed recipient, uint64 partialWithdrawalAmountGwei)
func (_IEigenPod *IEigenPodFilterer) WatchPartialWithdrawalRedeemed(opts *bind.WatchOpts, sink chan<- *IEigenPodPartialWithdrawalRedeemed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IEigenPod.contract.WatchLogs(opts, "PartialWithdrawalRedeemed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodPartialWithdrawalRedeemed)
				if err := _IEigenPod.contract.UnpackLog(event, "PartialWithdrawalRedeemed", log); err != nil {
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

// ParsePartialWithdrawalRedeemed is a log parse operation binding the contract event 0x8a7335714231dbd551aaba6314f4a97a14c201e53a3e25e1140325cdf67d7a4e.
//
// Solidity: event PartialWithdrawalRedeemed(uint40 validatorIndex, uint64 withdrawalTimestamp, address indexed recipient, uint64 partialWithdrawalAmountGwei)
func (_IEigenPod *IEigenPodFilterer) ParsePartialWithdrawalRedeemed(log types.Log) (*IEigenPodPartialWithdrawalRedeemed, error) {
	event := new(IEigenPodPartialWithdrawalRedeemed)
	if err := _IEigenPod.contract.UnpackLog(event, "PartialWithdrawalRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodRestakedBeaconChainETHWithdrawnIterator is returned from FilterRestakedBeaconChainETHWithdrawn and is used to iterate over the raw logs and unpacked data for RestakedBeaconChainETHWithdrawn events raised by the IEigenPod contract.
type IEigenPodRestakedBeaconChainETHWithdrawnIterator struct {
	Event *IEigenPodRestakedBeaconChainETHWithdrawn // Event containing the contract specifics and raw log

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
func (it *IEigenPodRestakedBeaconChainETHWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodRestakedBeaconChainETHWithdrawn)
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
		it.Event = new(IEigenPodRestakedBeaconChainETHWithdrawn)
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
func (it *IEigenPodRestakedBeaconChainETHWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodRestakedBeaconChainETHWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodRestakedBeaconChainETHWithdrawn represents a RestakedBeaconChainETHWithdrawn event raised by the IEigenPod contract.
type IEigenPodRestakedBeaconChainETHWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRestakedBeaconChainETHWithdrawn is a free log retrieval operation binding the contract event 0x8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e.
//
// Solidity: event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount)
func (_IEigenPod *IEigenPodFilterer) FilterRestakedBeaconChainETHWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*IEigenPodRestakedBeaconChainETHWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IEigenPod.contract.FilterLogs(opts, "RestakedBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodRestakedBeaconChainETHWithdrawnIterator{contract: _IEigenPod.contract, event: "RestakedBeaconChainETHWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRestakedBeaconChainETHWithdrawn is a free log subscription operation binding the contract event 0x8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e.
//
// Solidity: event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount)
func (_IEigenPod *IEigenPodFilterer) WatchRestakedBeaconChainETHWithdrawn(opts *bind.WatchOpts, sink chan<- *IEigenPodRestakedBeaconChainETHWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IEigenPod.contract.WatchLogs(opts, "RestakedBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodRestakedBeaconChainETHWithdrawn)
				if err := _IEigenPod.contract.UnpackLog(event, "RestakedBeaconChainETHWithdrawn", log); err != nil {
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

// ParseRestakedBeaconChainETHWithdrawn is a log parse operation binding the contract event 0x8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e.
//
// Solidity: event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount)
func (_IEigenPod *IEigenPodFilterer) ParseRestakedBeaconChainETHWithdrawn(log types.Log) (*IEigenPodRestakedBeaconChainETHWithdrawn, error) {
	event := new(IEigenPodRestakedBeaconChainETHWithdrawn)
	if err := _IEigenPod.contract.UnpackLog(event, "RestakedBeaconChainETHWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodRestakingActivatedIterator is returned from FilterRestakingActivated and is used to iterate over the raw logs and unpacked data for RestakingActivated events raised by the IEigenPod contract.
type IEigenPodRestakingActivatedIterator struct {
	Event *IEigenPodRestakingActivated // Event containing the contract specifics and raw log

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
func (it *IEigenPodRestakingActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodRestakingActivated)
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
		it.Event = new(IEigenPodRestakingActivated)
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
func (it *IEigenPodRestakingActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodRestakingActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodRestakingActivated represents a RestakingActivated event raised by the IEigenPod contract.
type IEigenPodRestakingActivated struct {
	PodOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRestakingActivated is a free log retrieval operation binding the contract event 0xca8dfc8c5e0a67a74501c072a3325f685259bebbae7cfd230ab85198a78b70cd.
//
// Solidity: event RestakingActivated(address indexed podOwner)
func (_IEigenPod *IEigenPodFilterer) FilterRestakingActivated(opts *bind.FilterOpts, podOwner []common.Address) (*IEigenPodRestakingActivatedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _IEigenPod.contract.FilterLogs(opts, "RestakingActivated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodRestakingActivatedIterator{contract: _IEigenPod.contract, event: "RestakingActivated", logs: logs, sub: sub}, nil
}

// WatchRestakingActivated is a free log subscription operation binding the contract event 0xca8dfc8c5e0a67a74501c072a3325f685259bebbae7cfd230ab85198a78b70cd.
//
// Solidity: event RestakingActivated(address indexed podOwner)
func (_IEigenPod *IEigenPodFilterer) WatchRestakingActivated(opts *bind.WatchOpts, sink chan<- *IEigenPodRestakingActivated, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _IEigenPod.contract.WatchLogs(opts, "RestakingActivated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodRestakingActivated)
				if err := _IEigenPod.contract.UnpackLog(event, "RestakingActivated", log); err != nil {
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

// ParseRestakingActivated is a log parse operation binding the contract event 0xca8dfc8c5e0a67a74501c072a3325f685259bebbae7cfd230ab85198a78b70cd.
//
// Solidity: event RestakingActivated(address indexed podOwner)
func (_IEigenPod *IEigenPodFilterer) ParseRestakingActivated(log types.Log) (*IEigenPodRestakingActivated, error) {
	event := new(IEigenPodRestakingActivated)
	if err := _IEigenPod.contract.UnpackLog(event, "RestakingActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodValidatorBalanceUpdatedIterator is returned from FilterValidatorBalanceUpdated and is used to iterate over the raw logs and unpacked data for ValidatorBalanceUpdated events raised by the IEigenPod contract.
type IEigenPodValidatorBalanceUpdatedIterator struct {
	Event *IEigenPodValidatorBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *IEigenPodValidatorBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodValidatorBalanceUpdated)
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
		it.Event = new(IEigenPodValidatorBalanceUpdated)
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
func (it *IEigenPodValidatorBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodValidatorBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodValidatorBalanceUpdated represents a ValidatorBalanceUpdated event raised by the IEigenPod contract.
type IEigenPodValidatorBalanceUpdated struct {
	ValidatorIndex          *big.Int
	BalanceTimestamp        uint64
	NewValidatorBalanceGwei uint64
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterValidatorBalanceUpdated is a free log retrieval operation binding the contract event 0x0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df.
//
// Solidity: event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_IEigenPod *IEigenPodFilterer) FilterValidatorBalanceUpdated(opts *bind.FilterOpts) (*IEigenPodValidatorBalanceUpdatedIterator, error) {

	logs, sub, err := _IEigenPod.contract.FilterLogs(opts, "ValidatorBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return &IEigenPodValidatorBalanceUpdatedIterator{contract: _IEigenPod.contract, event: "ValidatorBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorBalanceUpdated is a free log subscription operation binding the contract event 0x0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df.
//
// Solidity: event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_IEigenPod *IEigenPodFilterer) WatchValidatorBalanceUpdated(opts *bind.WatchOpts, sink chan<- *IEigenPodValidatorBalanceUpdated) (event.Subscription, error) {

	logs, sub, err := _IEigenPod.contract.WatchLogs(opts, "ValidatorBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodValidatorBalanceUpdated)
				if err := _IEigenPod.contract.UnpackLog(event, "ValidatorBalanceUpdated", log); err != nil {
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

// ParseValidatorBalanceUpdated is a log parse operation binding the contract event 0x0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df.
//
// Solidity: event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_IEigenPod *IEigenPodFilterer) ParseValidatorBalanceUpdated(log types.Log) (*IEigenPodValidatorBalanceUpdated, error) {
	event := new(IEigenPodValidatorBalanceUpdated)
	if err := _IEigenPod.contract.UnpackLog(event, "ValidatorBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodValidatorRestakedIterator is returned from FilterValidatorRestaked and is used to iterate over the raw logs and unpacked data for ValidatorRestaked events raised by the IEigenPod contract.
type IEigenPodValidatorRestakedIterator struct {
	Event *IEigenPodValidatorRestaked // Event containing the contract specifics and raw log

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
func (it *IEigenPodValidatorRestakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodValidatorRestaked)
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
		it.Event = new(IEigenPodValidatorRestaked)
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
func (it *IEigenPodValidatorRestakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodValidatorRestakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodValidatorRestaked represents a ValidatorRestaked event raised by the IEigenPod contract.
type IEigenPodValidatorRestaked struct {
	ValidatorIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorRestaked is a free log retrieval operation binding the contract event 0x2d0800bbc377ea54a08c5db6a87aafff5e3e9c8fead0eda110e40e0c10441449.
//
// Solidity: event ValidatorRestaked(uint40 validatorIndex)
func (_IEigenPod *IEigenPodFilterer) FilterValidatorRestaked(opts *bind.FilterOpts) (*IEigenPodValidatorRestakedIterator, error) {

	logs, sub, err := _IEigenPod.contract.FilterLogs(opts, "ValidatorRestaked")
	if err != nil {
		return nil, err
	}
	return &IEigenPodValidatorRestakedIterator{contract: _IEigenPod.contract, event: "ValidatorRestaked", logs: logs, sub: sub}, nil
}

// WatchValidatorRestaked is a free log subscription operation binding the contract event 0x2d0800bbc377ea54a08c5db6a87aafff5e3e9c8fead0eda110e40e0c10441449.
//
// Solidity: event ValidatorRestaked(uint40 validatorIndex)
func (_IEigenPod *IEigenPodFilterer) WatchValidatorRestaked(opts *bind.WatchOpts, sink chan<- *IEigenPodValidatorRestaked) (event.Subscription, error) {

	logs, sub, err := _IEigenPod.contract.WatchLogs(opts, "ValidatorRestaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodValidatorRestaked)
				if err := _IEigenPod.contract.UnpackLog(event, "ValidatorRestaked", log); err != nil {
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

// ParseValidatorRestaked is a log parse operation binding the contract event 0x2d0800bbc377ea54a08c5db6a87aafff5e3e9c8fead0eda110e40e0c10441449.
//
// Solidity: event ValidatorRestaked(uint40 validatorIndex)
func (_IEigenPod *IEigenPodFilterer) ParseValidatorRestaked(log types.Log) (*IEigenPodValidatorRestaked, error) {
	event := new(IEigenPodValidatorRestaked)
	if err := _IEigenPod.contract.UnpackLog(event, "ValidatorRestaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
