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

// BeaconChainProofsBalanceContainerProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsBalanceContainerProof struct {
	BalanceContainerRoot [32]byte
	Proof                []byte
}

// BeaconChainProofsBalanceProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsBalanceProof struct {
	PubkeyHash  [32]byte
	BalanceRoot [32]byte
	Proof       []byte
}

// BeaconChainProofsStateRootProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsStateRootProof struct {
	BeaconStateRoot [32]byte
	Proof           []byte
}

// BeaconChainProofsValidatorProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsValidatorProof struct {
	ValidatorFields [][32]byte
	Proof           []byte
}

// IEigenPodTypesCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type IEigenPodTypesCheckpoint struct {
	BeaconBlockRoot       [32]byte
	ProofsRemaining       *big.Int
	PodBalanceGwei        uint64
	BalanceDeltasGwei     int64
	PrevBeaconBalanceGwei uint64
}

// IEigenPodTypesValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IEigenPodTypesValidatorInfo struct {
	ValidatorIndex      uint64
	RestakedBalanceGwei uint64
	LastCheckpointedAt  uint64
	Status              uint8
}

// IEigenPodMetaData contains all meta data concerning the IEigenPod contract.
var IEigenPodMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"activeValidatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpointBalanceExitedGwei\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentCheckpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPodTypes.Checkpoint\",\"components\":[{\"name\":\"beaconBlockRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proofsRemaining\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"podBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"balanceDeltasGwei\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"prevBeaconBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentCheckpointTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParentBlockRoot\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastCheckpointTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"podOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proofSubmitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverTokens\",\"inputs\":[{\"name\":\"tokenList\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"amountsToWithdraw\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProofSubmitter\",\"inputs\":[{\"name\":\"newProofSubmitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"startCheckpoint\",\"inputs\":[{\"name\":\"revertIfNoBalance\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorPubkeyHashToInfo\",\"inputs\":[{\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPodTypes.ValidatorInfo\",\"components\":[{\"name\":\"validatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"restakedBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lastCheckpointedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPodTypes.VALIDATOR_STATUS\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorPubkeyToInfo\",\"inputs\":[{\"name\":\"validatorPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPodTypes.ValidatorInfo\",\"components\":[{\"name\":\"validatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"restakedBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lastCheckpointedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPodTypes.VALIDATOR_STATUS\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorStatus\",\"inputs\":[{\"name\":\"validatorPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPodTypes.VALIDATOR_STATUS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorStatus\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPodTypes.VALIDATOR_STATUS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyCheckpointProofs\",\"inputs\":[{\"name\":\"balanceContainerProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.BalanceContainerProof\",\"components\":[{\"name\":\"balanceContainerRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"proofs\",\"type\":\"tuple[]\",\"internalType\":\"structBeaconChainProofs.BalanceProof[]\",\"components\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"balanceRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyStaleBalance\",\"inputs\":[{\"name\":\"beaconTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.ValidatorProof\",\"components\":[{\"name\":\"validatorFields\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyWithdrawalCredentials\",\"inputs\":[{\"name\":\"beaconTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"validatorIndices\",\"type\":\"uint40[]\",\"internalType\":\"uint40[]\"},{\"name\":\"validatorFieldsProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"validatorFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawRestakedBeaconChainETH\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawableRestakedExecutionLayerGwei\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CheckpointCreated\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"beaconBlockRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validatorCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointFinalized\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"totalShareDeltaWei\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EigenPodStaked\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonBeaconChainETHReceived\",\"inputs\":[{\"name\":\"amountReceived\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProofSubmitterUpdated\",\"inputs\":[{\"name\":\"prevProofSubmitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newProofSubmitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RestakedBeaconChainETHWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorBalanceUpdated\",\"inputs\":[{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"balanceTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newValidatorBalanceGwei\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorCheckpointed\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":true,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRestaked\",\"inputs\":[{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorWithdrawn\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"validatorIndex\",\"type\":\"uint40\",\"indexed\":true,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BeaconTimestampTooFarInPast\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotCheckpointTwiceInSingleBlock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CheckpointAlreadyActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CredentialsAlreadyVerified\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientWithdrawableBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEIP4788Response\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPubKeyLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MsgValueNot32ETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoActiveCheckpoint\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoBalanceToCheckpoint\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEigenPodManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEigenPodOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEigenPodOwnerOrProofSubmitter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TimestampOutOfRange\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorInactiveOnBeaconChain\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorIsExitingBeaconChain\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorNotActiveInPod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorNotSlashedOnBeaconChain\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalCredentialsNotForEigenPod\",\"inputs\":[]}]",
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

// ActiveValidatorCount is a free data retrieval call binding the contract method 0x2340e8d3.
//
// Solidity: function activeValidatorCount() view returns(uint256)
func (_IEigenPod *IEigenPodCaller) ActiveValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "activeValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveValidatorCount is a free data retrieval call binding the contract method 0x2340e8d3.
//
// Solidity: function activeValidatorCount() view returns(uint256)
func (_IEigenPod *IEigenPodSession) ActiveValidatorCount() (*big.Int, error) {
	return _IEigenPod.Contract.ActiveValidatorCount(&_IEigenPod.CallOpts)
}

// ActiveValidatorCount is a free data retrieval call binding the contract method 0x2340e8d3.
//
// Solidity: function activeValidatorCount() view returns(uint256)
func (_IEigenPod *IEigenPodCallerSession) ActiveValidatorCount() (*big.Int, error) {
	return _IEigenPod.Contract.ActiveValidatorCount(&_IEigenPod.CallOpts)
}

// CheckpointBalanceExitedGwei is a free data retrieval call binding the contract method 0x52396a59.
//
// Solidity: function checkpointBalanceExitedGwei(uint64 ) view returns(uint64)
func (_IEigenPod *IEigenPodCaller) CheckpointBalanceExitedGwei(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "checkpointBalanceExitedGwei", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CheckpointBalanceExitedGwei is a free data retrieval call binding the contract method 0x52396a59.
//
// Solidity: function checkpointBalanceExitedGwei(uint64 ) view returns(uint64)
func (_IEigenPod *IEigenPodSession) CheckpointBalanceExitedGwei(arg0 uint64) (uint64, error) {
	return _IEigenPod.Contract.CheckpointBalanceExitedGwei(&_IEigenPod.CallOpts, arg0)
}

// CheckpointBalanceExitedGwei is a free data retrieval call binding the contract method 0x52396a59.
//
// Solidity: function checkpointBalanceExitedGwei(uint64 ) view returns(uint64)
func (_IEigenPod *IEigenPodCallerSession) CheckpointBalanceExitedGwei(arg0 uint64) (uint64, error) {
	return _IEigenPod.Contract.CheckpointBalanceExitedGwei(&_IEigenPod.CallOpts, arg0)
}

// CurrentCheckpoint is a free data retrieval call binding the contract method 0x47d28372.
//
// Solidity: function currentCheckpoint() view returns((bytes32,uint24,uint64,int64,uint64))
func (_IEigenPod *IEigenPodCaller) CurrentCheckpoint(opts *bind.CallOpts) (IEigenPodTypesCheckpoint, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "currentCheckpoint")

	if err != nil {
		return *new(IEigenPodTypesCheckpoint), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodTypesCheckpoint)).(*IEigenPodTypesCheckpoint)

	return out0, err

}

// CurrentCheckpoint is a free data retrieval call binding the contract method 0x47d28372.
//
// Solidity: function currentCheckpoint() view returns((bytes32,uint24,uint64,int64,uint64))
func (_IEigenPod *IEigenPodSession) CurrentCheckpoint() (IEigenPodTypesCheckpoint, error) {
	return _IEigenPod.Contract.CurrentCheckpoint(&_IEigenPod.CallOpts)
}

// CurrentCheckpoint is a free data retrieval call binding the contract method 0x47d28372.
//
// Solidity: function currentCheckpoint() view returns((bytes32,uint24,uint64,int64,uint64))
func (_IEigenPod *IEigenPodCallerSession) CurrentCheckpoint() (IEigenPodTypesCheckpoint, error) {
	return _IEigenPod.Contract.CurrentCheckpoint(&_IEigenPod.CallOpts)
}

// CurrentCheckpointTimestamp is a free data retrieval call binding the contract method 0x42ecff2a.
//
// Solidity: function currentCheckpointTimestamp() view returns(uint64)
func (_IEigenPod *IEigenPodCaller) CurrentCheckpointTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "currentCheckpointTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurrentCheckpointTimestamp is a free data retrieval call binding the contract method 0x42ecff2a.
//
// Solidity: function currentCheckpointTimestamp() view returns(uint64)
func (_IEigenPod *IEigenPodSession) CurrentCheckpointTimestamp() (uint64, error) {
	return _IEigenPod.Contract.CurrentCheckpointTimestamp(&_IEigenPod.CallOpts)
}

// CurrentCheckpointTimestamp is a free data retrieval call binding the contract method 0x42ecff2a.
//
// Solidity: function currentCheckpointTimestamp() view returns(uint64)
func (_IEigenPod *IEigenPodCallerSession) CurrentCheckpointTimestamp() (uint64, error) {
	return _IEigenPod.Contract.CurrentCheckpointTimestamp(&_IEigenPod.CallOpts)
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

// GetParentBlockRoot is a free data retrieval call binding the contract method 0x6c0d2d5a.
//
// Solidity: function getParentBlockRoot(uint64 timestamp) view returns(bytes32)
func (_IEigenPod *IEigenPodCaller) GetParentBlockRoot(opts *bind.CallOpts, timestamp uint64) ([32]byte, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "getParentBlockRoot", timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetParentBlockRoot is a free data retrieval call binding the contract method 0x6c0d2d5a.
//
// Solidity: function getParentBlockRoot(uint64 timestamp) view returns(bytes32)
func (_IEigenPod *IEigenPodSession) GetParentBlockRoot(timestamp uint64) ([32]byte, error) {
	return _IEigenPod.Contract.GetParentBlockRoot(&_IEigenPod.CallOpts, timestamp)
}

// GetParentBlockRoot is a free data retrieval call binding the contract method 0x6c0d2d5a.
//
// Solidity: function getParentBlockRoot(uint64 timestamp) view returns(bytes32)
func (_IEigenPod *IEigenPodCallerSession) GetParentBlockRoot(timestamp uint64) ([32]byte, error) {
	return _IEigenPod.Contract.GetParentBlockRoot(&_IEigenPod.CallOpts, timestamp)
}

// LastCheckpointTimestamp is a free data retrieval call binding the contract method 0xee94d67c.
//
// Solidity: function lastCheckpointTimestamp() view returns(uint64)
func (_IEigenPod *IEigenPodCaller) LastCheckpointTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "lastCheckpointTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastCheckpointTimestamp is a free data retrieval call binding the contract method 0xee94d67c.
//
// Solidity: function lastCheckpointTimestamp() view returns(uint64)
func (_IEigenPod *IEigenPodSession) LastCheckpointTimestamp() (uint64, error) {
	return _IEigenPod.Contract.LastCheckpointTimestamp(&_IEigenPod.CallOpts)
}

// LastCheckpointTimestamp is a free data retrieval call binding the contract method 0xee94d67c.
//
// Solidity: function lastCheckpointTimestamp() view returns(uint64)
func (_IEigenPod *IEigenPodCallerSession) LastCheckpointTimestamp() (uint64, error) {
	return _IEigenPod.Contract.LastCheckpointTimestamp(&_IEigenPod.CallOpts)
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

// ProofSubmitter is a free data retrieval call binding the contract method 0x58753357.
//
// Solidity: function proofSubmitter() view returns(address)
func (_IEigenPod *IEigenPodCaller) ProofSubmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "proofSubmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProofSubmitter is a free data retrieval call binding the contract method 0x58753357.
//
// Solidity: function proofSubmitter() view returns(address)
func (_IEigenPod *IEigenPodSession) ProofSubmitter() (common.Address, error) {
	return _IEigenPod.Contract.ProofSubmitter(&_IEigenPod.CallOpts)
}

// ProofSubmitter is a free data retrieval call binding the contract method 0x58753357.
//
// Solidity: function proofSubmitter() view returns(address)
func (_IEigenPod *IEigenPodCallerSession) ProofSubmitter() (common.Address, error) {
	return _IEigenPod.Contract.ProofSubmitter(&_IEigenPod.CallOpts)
}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_IEigenPod *IEigenPodCaller) ValidatorPubkeyHashToInfo(opts *bind.CallOpts, validatorPubkeyHash [32]byte) (IEigenPodTypesValidatorInfo, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "validatorPubkeyHashToInfo", validatorPubkeyHash)

	if err != nil {
		return *new(IEigenPodTypesValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodTypesValidatorInfo)).(*IEigenPodTypesValidatorInfo)

	return out0, err

}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_IEigenPod *IEigenPodSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodTypesValidatorInfo, error) {
	return _IEigenPod.Contract.ValidatorPubkeyHashToInfo(&_IEigenPod.CallOpts, validatorPubkeyHash)
}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_IEigenPod *IEigenPodCallerSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodTypesValidatorInfo, error) {
	return _IEigenPod.Contract.ValidatorPubkeyHashToInfo(&_IEigenPod.CallOpts, validatorPubkeyHash)
}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_IEigenPod *IEigenPodCaller) ValidatorPubkeyToInfo(opts *bind.CallOpts, validatorPubkey []byte) (IEigenPodTypesValidatorInfo, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "validatorPubkeyToInfo", validatorPubkey)

	if err != nil {
		return *new(IEigenPodTypesValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodTypesValidatorInfo)).(*IEigenPodTypesValidatorInfo)

	return out0, err

}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_IEigenPod *IEigenPodSession) ValidatorPubkeyToInfo(validatorPubkey []byte) (IEigenPodTypesValidatorInfo, error) {
	return _IEigenPod.Contract.ValidatorPubkeyToInfo(&_IEigenPod.CallOpts, validatorPubkey)
}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_IEigenPod *IEigenPodCallerSession) ValidatorPubkeyToInfo(validatorPubkey []byte) (IEigenPodTypesValidatorInfo, error) {
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IEigenPod *IEigenPodCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IEigenPod.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IEigenPod *IEigenPodSession) Version() (string, error) {
	return _IEigenPod.Contract.Version(&_IEigenPod.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IEigenPod *IEigenPodCallerSession) Version() (string, error) {
	return _IEigenPod.Contract.Version(&_IEigenPod.CallOpts)
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

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xd06d5587.
//
// Solidity: function setProofSubmitter(address newProofSubmitter) returns()
func (_IEigenPod *IEigenPodTransactor) SetProofSubmitter(opts *bind.TransactOpts, newProofSubmitter common.Address) (*types.Transaction, error) {
	return _IEigenPod.contract.Transact(opts, "setProofSubmitter", newProofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xd06d5587.
//
// Solidity: function setProofSubmitter(address newProofSubmitter) returns()
func (_IEigenPod *IEigenPodSession) SetProofSubmitter(newProofSubmitter common.Address) (*types.Transaction, error) {
	return _IEigenPod.Contract.SetProofSubmitter(&_IEigenPod.TransactOpts, newProofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xd06d5587.
//
// Solidity: function setProofSubmitter(address newProofSubmitter) returns()
func (_IEigenPod *IEigenPodTransactorSession) SetProofSubmitter(newProofSubmitter common.Address) (*types.Transaction, error) {
	return _IEigenPod.Contract.SetProofSubmitter(&_IEigenPod.TransactOpts, newProofSubmitter)
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

// StartCheckpoint is a paid mutator transaction binding the contract method 0x88676cad.
//
// Solidity: function startCheckpoint(bool revertIfNoBalance) returns()
func (_IEigenPod *IEigenPodTransactor) StartCheckpoint(opts *bind.TransactOpts, revertIfNoBalance bool) (*types.Transaction, error) {
	return _IEigenPod.contract.Transact(opts, "startCheckpoint", revertIfNoBalance)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x88676cad.
//
// Solidity: function startCheckpoint(bool revertIfNoBalance) returns()
func (_IEigenPod *IEigenPodSession) StartCheckpoint(revertIfNoBalance bool) (*types.Transaction, error) {
	return _IEigenPod.Contract.StartCheckpoint(&_IEigenPod.TransactOpts, revertIfNoBalance)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x88676cad.
//
// Solidity: function startCheckpoint(bool revertIfNoBalance) returns()
func (_IEigenPod *IEigenPodTransactorSession) StartCheckpoint(revertIfNoBalance bool) (*types.Transaction, error) {
	return _IEigenPod.Contract.StartCheckpoint(&_IEigenPod.TransactOpts, revertIfNoBalance)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0xf074ba62.
//
// Solidity: function verifyCheckpointProofs((bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_IEigenPod *IEigenPodTransactor) VerifyCheckpointProofs(opts *bind.TransactOpts, balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _IEigenPod.contract.Transact(opts, "verifyCheckpointProofs", balanceContainerProof, proofs)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0xf074ba62.
//
// Solidity: function verifyCheckpointProofs((bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_IEigenPod *IEigenPodSession) VerifyCheckpointProofs(balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _IEigenPod.Contract.VerifyCheckpointProofs(&_IEigenPod.TransactOpts, balanceContainerProof, proofs)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0xf074ba62.
//
// Solidity: function verifyCheckpointProofs((bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_IEigenPod *IEigenPodTransactorSession) VerifyCheckpointProofs(balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _IEigenPod.Contract.VerifyCheckpointProofs(&_IEigenPod.TransactOpts, balanceContainerProof, proofs)
}

// VerifyStaleBalance is a paid mutator transaction binding the contract method 0x039157d2.
//
// Solidity: function verifyStaleBalance(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, (bytes32[],bytes) proof) returns()
func (_IEigenPod *IEigenPodTransactor) VerifyStaleBalance(opts *bind.TransactOpts, beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, proof BeaconChainProofsValidatorProof) (*types.Transaction, error) {
	return _IEigenPod.contract.Transact(opts, "verifyStaleBalance", beaconTimestamp, stateRootProof, proof)
}

// VerifyStaleBalance is a paid mutator transaction binding the contract method 0x039157d2.
//
// Solidity: function verifyStaleBalance(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, (bytes32[],bytes) proof) returns()
func (_IEigenPod *IEigenPodSession) VerifyStaleBalance(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, proof BeaconChainProofsValidatorProof) (*types.Transaction, error) {
	return _IEigenPod.Contract.VerifyStaleBalance(&_IEigenPod.TransactOpts, beaconTimestamp, stateRootProof, proof)
}

// VerifyStaleBalance is a paid mutator transaction binding the contract method 0x039157d2.
//
// Solidity: function verifyStaleBalance(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, (bytes32[],bytes) proof) returns()
func (_IEigenPod *IEigenPodTransactorSession) VerifyStaleBalance(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, proof BeaconChainProofsValidatorProof) (*types.Transaction, error) {
	return _IEigenPod.Contract.VerifyStaleBalance(&_IEigenPod.TransactOpts, beaconTimestamp, stateRootProof, proof)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_IEigenPod *IEigenPodTransactor) VerifyWithdrawalCredentials(opts *bind.TransactOpts, beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _IEigenPod.contract.Transact(opts, "verifyWithdrawalCredentials", beaconTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_IEigenPod *IEigenPodSession) VerifyWithdrawalCredentials(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _IEigenPod.Contract.VerifyWithdrawalCredentials(&_IEigenPod.TransactOpts, beaconTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_IEigenPod *IEigenPodTransactorSession) VerifyWithdrawalCredentials(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _IEigenPod.Contract.VerifyWithdrawalCredentials(&_IEigenPod.TransactOpts, beaconTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
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

// IEigenPodCheckpointCreatedIterator is returned from FilterCheckpointCreated and is used to iterate over the raw logs and unpacked data for CheckpointCreated events raised by the IEigenPod contract.
type IEigenPodCheckpointCreatedIterator struct {
	Event *IEigenPodCheckpointCreated // Event containing the contract specifics and raw log

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
func (it *IEigenPodCheckpointCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodCheckpointCreated)
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
		it.Event = new(IEigenPodCheckpointCreated)
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
func (it *IEigenPodCheckpointCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodCheckpointCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodCheckpointCreated represents a CheckpointCreated event raised by the IEigenPod contract.
type IEigenPodCheckpointCreated struct {
	CheckpointTimestamp uint64
	BeaconBlockRoot     [32]byte
	ValidatorCount      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCheckpointCreated is a free log retrieval operation binding the contract event 0x575796133bbed337e5b39aa49a30dc2556a91e0c6c2af4b7b886ae77ebef1076.
//
// Solidity: event CheckpointCreated(uint64 indexed checkpointTimestamp, bytes32 indexed beaconBlockRoot, uint256 validatorCount)
func (_IEigenPod *IEigenPodFilterer) FilterCheckpointCreated(opts *bind.FilterOpts, checkpointTimestamp []uint64, beaconBlockRoot [][32]byte) (*IEigenPodCheckpointCreatedIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var beaconBlockRootRule []interface{}
	for _, beaconBlockRootItem := range beaconBlockRoot {
		beaconBlockRootRule = append(beaconBlockRootRule, beaconBlockRootItem)
	}

	logs, sub, err := _IEigenPod.contract.FilterLogs(opts, "CheckpointCreated", checkpointTimestampRule, beaconBlockRootRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodCheckpointCreatedIterator{contract: _IEigenPod.contract, event: "CheckpointCreated", logs: logs, sub: sub}, nil
}

// WatchCheckpointCreated is a free log subscription operation binding the contract event 0x575796133bbed337e5b39aa49a30dc2556a91e0c6c2af4b7b886ae77ebef1076.
//
// Solidity: event CheckpointCreated(uint64 indexed checkpointTimestamp, bytes32 indexed beaconBlockRoot, uint256 validatorCount)
func (_IEigenPod *IEigenPodFilterer) WatchCheckpointCreated(opts *bind.WatchOpts, sink chan<- *IEigenPodCheckpointCreated, checkpointTimestamp []uint64, beaconBlockRoot [][32]byte) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var beaconBlockRootRule []interface{}
	for _, beaconBlockRootItem := range beaconBlockRoot {
		beaconBlockRootRule = append(beaconBlockRootRule, beaconBlockRootItem)
	}

	logs, sub, err := _IEigenPod.contract.WatchLogs(opts, "CheckpointCreated", checkpointTimestampRule, beaconBlockRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodCheckpointCreated)
				if err := _IEigenPod.contract.UnpackLog(event, "CheckpointCreated", log); err != nil {
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

// ParseCheckpointCreated is a log parse operation binding the contract event 0x575796133bbed337e5b39aa49a30dc2556a91e0c6c2af4b7b886ae77ebef1076.
//
// Solidity: event CheckpointCreated(uint64 indexed checkpointTimestamp, bytes32 indexed beaconBlockRoot, uint256 validatorCount)
func (_IEigenPod *IEigenPodFilterer) ParseCheckpointCreated(log types.Log) (*IEigenPodCheckpointCreated, error) {
	event := new(IEigenPodCheckpointCreated)
	if err := _IEigenPod.contract.UnpackLog(event, "CheckpointCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEigenPodCheckpointFinalizedIterator is returned from FilterCheckpointFinalized and is used to iterate over the raw logs and unpacked data for CheckpointFinalized events raised by the IEigenPod contract.
type IEigenPodCheckpointFinalizedIterator struct {
	Event *IEigenPodCheckpointFinalized // Event containing the contract specifics and raw log

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
func (it *IEigenPodCheckpointFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodCheckpointFinalized)
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
		it.Event = new(IEigenPodCheckpointFinalized)
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
func (it *IEigenPodCheckpointFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodCheckpointFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodCheckpointFinalized represents a CheckpointFinalized event raised by the IEigenPod contract.
type IEigenPodCheckpointFinalized struct {
	CheckpointTimestamp uint64
	TotalShareDeltaWei  *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCheckpointFinalized is a free log retrieval operation binding the contract event 0x525408c201bc1576eb44116f6478f1c2a54775b19a043bcfdc708364f74f8e44.
//
// Solidity: event CheckpointFinalized(uint64 indexed checkpointTimestamp, int256 totalShareDeltaWei)
func (_IEigenPod *IEigenPodFilterer) FilterCheckpointFinalized(opts *bind.FilterOpts, checkpointTimestamp []uint64) (*IEigenPodCheckpointFinalizedIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}

	logs, sub, err := _IEigenPod.contract.FilterLogs(opts, "CheckpointFinalized", checkpointTimestampRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodCheckpointFinalizedIterator{contract: _IEigenPod.contract, event: "CheckpointFinalized", logs: logs, sub: sub}, nil
}

// WatchCheckpointFinalized is a free log subscription operation binding the contract event 0x525408c201bc1576eb44116f6478f1c2a54775b19a043bcfdc708364f74f8e44.
//
// Solidity: event CheckpointFinalized(uint64 indexed checkpointTimestamp, int256 totalShareDeltaWei)
func (_IEigenPod *IEigenPodFilterer) WatchCheckpointFinalized(opts *bind.WatchOpts, sink chan<- *IEigenPodCheckpointFinalized, checkpointTimestamp []uint64) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}

	logs, sub, err := _IEigenPod.contract.WatchLogs(opts, "CheckpointFinalized", checkpointTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodCheckpointFinalized)
				if err := _IEigenPod.contract.UnpackLog(event, "CheckpointFinalized", log); err != nil {
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

// ParseCheckpointFinalized is a log parse operation binding the contract event 0x525408c201bc1576eb44116f6478f1c2a54775b19a043bcfdc708364f74f8e44.
//
// Solidity: event CheckpointFinalized(uint64 indexed checkpointTimestamp, int256 totalShareDeltaWei)
func (_IEigenPod *IEigenPodFilterer) ParseCheckpointFinalized(log types.Log) (*IEigenPodCheckpointFinalized, error) {
	event := new(IEigenPodCheckpointFinalized)
	if err := _IEigenPod.contract.UnpackLog(event, "CheckpointFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// IEigenPodProofSubmitterUpdatedIterator is returned from FilterProofSubmitterUpdated and is used to iterate over the raw logs and unpacked data for ProofSubmitterUpdated events raised by the IEigenPod contract.
type IEigenPodProofSubmitterUpdatedIterator struct {
	Event *IEigenPodProofSubmitterUpdated // Event containing the contract specifics and raw log

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
func (it *IEigenPodProofSubmitterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodProofSubmitterUpdated)
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
		it.Event = new(IEigenPodProofSubmitterUpdated)
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
func (it *IEigenPodProofSubmitterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodProofSubmitterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodProofSubmitterUpdated represents a ProofSubmitterUpdated event raised by the IEigenPod contract.
type IEigenPodProofSubmitterUpdated struct {
	PrevProofSubmitter common.Address
	NewProofSubmitter  common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProofSubmitterUpdated is a free log retrieval operation binding the contract event 0xfb8129080a19d34dceac04ba253fc50304dc86c729bd63cdca4a969ad19a5eac.
//
// Solidity: event ProofSubmitterUpdated(address prevProofSubmitter, address newProofSubmitter)
func (_IEigenPod *IEigenPodFilterer) FilterProofSubmitterUpdated(opts *bind.FilterOpts) (*IEigenPodProofSubmitterUpdatedIterator, error) {

	logs, sub, err := _IEigenPod.contract.FilterLogs(opts, "ProofSubmitterUpdated")
	if err != nil {
		return nil, err
	}
	return &IEigenPodProofSubmitterUpdatedIterator{contract: _IEigenPod.contract, event: "ProofSubmitterUpdated", logs: logs, sub: sub}, nil
}

// WatchProofSubmitterUpdated is a free log subscription operation binding the contract event 0xfb8129080a19d34dceac04ba253fc50304dc86c729bd63cdca4a969ad19a5eac.
//
// Solidity: event ProofSubmitterUpdated(address prevProofSubmitter, address newProofSubmitter)
func (_IEigenPod *IEigenPodFilterer) WatchProofSubmitterUpdated(opts *bind.WatchOpts, sink chan<- *IEigenPodProofSubmitterUpdated) (event.Subscription, error) {

	logs, sub, err := _IEigenPod.contract.WatchLogs(opts, "ProofSubmitterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodProofSubmitterUpdated)
				if err := _IEigenPod.contract.UnpackLog(event, "ProofSubmitterUpdated", log); err != nil {
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

// ParseProofSubmitterUpdated is a log parse operation binding the contract event 0xfb8129080a19d34dceac04ba253fc50304dc86c729bd63cdca4a969ad19a5eac.
//
// Solidity: event ProofSubmitterUpdated(address prevProofSubmitter, address newProofSubmitter)
func (_IEigenPod *IEigenPodFilterer) ParseProofSubmitterUpdated(log types.Log) (*IEigenPodProofSubmitterUpdated, error) {
	event := new(IEigenPodProofSubmitterUpdated)
	if err := _IEigenPod.contract.UnpackLog(event, "ProofSubmitterUpdated", log); err != nil {
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

// IEigenPodValidatorCheckpointedIterator is returned from FilterValidatorCheckpointed and is used to iterate over the raw logs and unpacked data for ValidatorCheckpointed events raised by the IEigenPod contract.
type IEigenPodValidatorCheckpointedIterator struct {
	Event *IEigenPodValidatorCheckpointed // Event containing the contract specifics and raw log

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
func (it *IEigenPodValidatorCheckpointedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodValidatorCheckpointed)
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
		it.Event = new(IEigenPodValidatorCheckpointed)
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
func (it *IEigenPodValidatorCheckpointedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodValidatorCheckpointedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodValidatorCheckpointed represents a ValidatorCheckpointed event raised by the IEigenPod contract.
type IEigenPodValidatorCheckpointed struct {
	CheckpointTimestamp uint64
	ValidatorIndex      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterValidatorCheckpointed is a free log retrieval operation binding the contract event 0xa91c59033c3423e18b54d0acecebb4972f9ea95aedf5f4cae3b677b02eaf3a3f.
//
// Solidity: event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_IEigenPod *IEigenPodFilterer) FilterValidatorCheckpointed(opts *bind.FilterOpts, checkpointTimestamp []uint64, validatorIndex []*big.Int) (*IEigenPodValidatorCheckpointedIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _IEigenPod.contract.FilterLogs(opts, "ValidatorCheckpointed", checkpointTimestampRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodValidatorCheckpointedIterator{contract: _IEigenPod.contract, event: "ValidatorCheckpointed", logs: logs, sub: sub}, nil
}

// WatchValidatorCheckpointed is a free log subscription operation binding the contract event 0xa91c59033c3423e18b54d0acecebb4972f9ea95aedf5f4cae3b677b02eaf3a3f.
//
// Solidity: event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_IEigenPod *IEigenPodFilterer) WatchValidatorCheckpointed(opts *bind.WatchOpts, sink chan<- *IEigenPodValidatorCheckpointed, checkpointTimestamp []uint64, validatorIndex []*big.Int) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _IEigenPod.contract.WatchLogs(opts, "ValidatorCheckpointed", checkpointTimestampRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodValidatorCheckpointed)
				if err := _IEigenPod.contract.UnpackLog(event, "ValidatorCheckpointed", log); err != nil {
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

// ParseValidatorCheckpointed is a log parse operation binding the contract event 0xa91c59033c3423e18b54d0acecebb4972f9ea95aedf5f4cae3b677b02eaf3a3f.
//
// Solidity: event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_IEigenPod *IEigenPodFilterer) ParseValidatorCheckpointed(log types.Log) (*IEigenPodValidatorCheckpointed, error) {
	event := new(IEigenPodValidatorCheckpointed)
	if err := _IEigenPod.contract.UnpackLog(event, "ValidatorCheckpointed", log); err != nil {
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

// IEigenPodValidatorWithdrawnIterator is returned from FilterValidatorWithdrawn and is used to iterate over the raw logs and unpacked data for ValidatorWithdrawn events raised by the IEigenPod contract.
type IEigenPodValidatorWithdrawnIterator struct {
	Event *IEigenPodValidatorWithdrawn // Event containing the contract specifics and raw log

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
func (it *IEigenPodValidatorWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEigenPodValidatorWithdrawn)
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
		it.Event = new(IEigenPodValidatorWithdrawn)
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
func (it *IEigenPodValidatorWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEigenPodValidatorWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEigenPodValidatorWithdrawn represents a ValidatorWithdrawn event raised by the IEigenPod contract.
type IEigenPodValidatorWithdrawn struct {
	CheckpointTimestamp uint64
	ValidatorIndex      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterValidatorWithdrawn is a free log retrieval operation binding the contract event 0x2a02361ffa66cf2c2da4682c2355a6adcaa9f6c227b6e6563e68480f9587626a.
//
// Solidity: event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_IEigenPod *IEigenPodFilterer) FilterValidatorWithdrawn(opts *bind.FilterOpts, checkpointTimestamp []uint64, validatorIndex []*big.Int) (*IEigenPodValidatorWithdrawnIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _IEigenPod.contract.FilterLogs(opts, "ValidatorWithdrawn", checkpointTimestampRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return &IEigenPodValidatorWithdrawnIterator{contract: _IEigenPod.contract, event: "ValidatorWithdrawn", logs: logs, sub: sub}, nil
}

// WatchValidatorWithdrawn is a free log subscription operation binding the contract event 0x2a02361ffa66cf2c2da4682c2355a6adcaa9f6c227b6e6563e68480f9587626a.
//
// Solidity: event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_IEigenPod *IEigenPodFilterer) WatchValidatorWithdrawn(opts *bind.WatchOpts, sink chan<- *IEigenPodValidatorWithdrawn, checkpointTimestamp []uint64, validatorIndex []*big.Int) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _IEigenPod.contract.WatchLogs(opts, "ValidatorWithdrawn", checkpointTimestampRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEigenPodValidatorWithdrawn)
				if err := _IEigenPod.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
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

// ParseValidatorWithdrawn is a log parse operation binding the contract event 0x2a02361ffa66cf2c2da4682c2355a6adcaa9f6c227b6e6563e68480f9587626a.
//
// Solidity: event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, uint40 indexed validatorIndex)
func (_IEigenPod *IEigenPodFilterer) ParseValidatorWithdrawn(log types.Log) (*IEigenPodValidatorWithdrawn, error) {
	event := new(IEigenPodValidatorWithdrawn)
	if err := _IEigenPod.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
