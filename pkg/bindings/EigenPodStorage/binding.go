// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EigenPodStorage

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

// IEigenPodTypesConsolidationRequest is an auto generated low-level Go binding around an user-defined struct.
type IEigenPodTypesConsolidationRequest struct {
	SrcPubkey    []byte
	TargetPubkey []byte
}

// IEigenPodTypesValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IEigenPodTypesValidatorInfo struct {
	ValidatorIndex      uint64
	RestakedBalanceGwei uint64
	LastCheckpointedAt  uint64
	Status              uint8
}

// IEigenPodTypesWithdrawalRequest is an auto generated low-level Go binding around an user-defined struct.
type IEigenPodTypesWithdrawalRequest struct {
	Pubkey     []byte
	AmountGwei uint64
}

// EigenPodStorageMetaData contains all meta data concerning the EigenPodStorage contract.
var EigenPodStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"activeValidatorCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpointBalanceExitedGwei\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentCheckpoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPodTypes.Checkpoint\",\"components\":[{\"name\":\"beaconBlockRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proofsRemaining\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"podBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"balanceDeltasGwei\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"prevBeaconBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentCheckpointTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConsolidationRequestFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParentBlockRoot\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawalRequestFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastCheckpointTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"podOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proofSubmitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverTokens\",\"inputs\":[{\"name\":\"tokenList\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"amountsToWithdraw\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestConsolidation\",\"inputs\":[{\"name\":\"requests\",\"type\":\"tuple[]\",\"internalType\":\"structIEigenPodTypes.ConsolidationRequest[]\",\"components\":[{\"name\":\"srcPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"targetPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestWithdrawal\",\"inputs\":[{\"name\":\"requests\",\"type\":\"tuple[]\",\"internalType\":\"structIEigenPodTypes.WithdrawalRequest[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amountGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setProofSubmitter\",\"inputs\":[{\"name\":\"newProofSubmitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"startCheckpoint\",\"inputs\":[{\"name\":\"revertIfNoBalance\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorPubkeyHashToInfo\",\"inputs\":[{\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPodTypes.ValidatorInfo\",\"components\":[{\"name\":\"validatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"restakedBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lastCheckpointedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPodTypes.VALIDATOR_STATUS\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorPubkeyToInfo\",\"inputs\":[{\"name\":\"validatorPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIEigenPodTypes.ValidatorInfo\",\"components\":[{\"name\":\"validatorIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"restakedBalanceGwei\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lastCheckpointedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPodTypes.VALIDATOR_STATUS\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorStatus\",\"inputs\":[{\"name\":\"validatorPubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPodTypes.VALIDATOR_STATUS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorStatus\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIEigenPodTypes.VALIDATOR_STATUS\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyCheckpointProofs\",\"inputs\":[{\"name\":\"balanceContainerProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.BalanceContainerProof\",\"components\":[{\"name\":\"balanceContainerRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"proofs\",\"type\":\"tuple[]\",\"internalType\":\"structBeaconChainProofs.BalanceProof[]\",\"components\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"balanceRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyStaleBalance\",\"inputs\":[{\"name\":\"beaconTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.ValidatorProof\",\"components\":[{\"name\":\"validatorFields\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyWithdrawalCredentials\",\"inputs\":[{\"name\":\"beaconTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"validatorIndices\",\"type\":\"uint40[]\",\"internalType\":\"uint40[]\"},{\"name\":\"validatorFieldsProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"validatorFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawRestakedBeaconChainETH\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawableRestakedExecutionLayerGwei\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CheckpointCreated\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"beaconBlockRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validatorCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckpointFinalized\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"totalShareDeltaWei\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConsolidationRequested\",\"inputs\":[{\"name\":\"sourcePubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"targetPubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EigenPodStaked\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExitRequested\",\"inputs\":[{\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonBeaconChainETHReceived\",\"inputs\":[{\"name\":\"amountReceived\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProofSubmitterUpdated\",\"inputs\":[{\"name\":\"prevProofSubmitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newProofSubmitter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RestakedBeaconChainETHWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SwitchToCompoundingRequested\",\"inputs\":[{\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorBalanceUpdated\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"balanceTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newValidatorBalanceGwei\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorCheckpointed\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRestaked\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorWithdrawn\",\"inputs\":[{\"name\":\"checkpointTimestamp\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalRequested\",\"inputs\":[{\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"withdrawalAmountGwei\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BeaconTimestampTooFarInPast\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotCheckpointTwiceInSingleBlock\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CheckpointAlreadyActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CredentialsAlreadyVerified\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeQueryFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ForkTimestampZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFunds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientWithdrawableBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEIP4788Response\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPubKeyLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MsgValueNot32ETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoActiveCheckpoint\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoBalanceToCheckpoint\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEigenPodManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEigenPodOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEigenPodOwnerOrProofSubmitter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PredeployFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RefundFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TimestampOutOfRange\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorInactiveOnBeaconChain\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorIsExitingBeaconChain\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorNotActiveInPod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidatorNotSlashedOnBeaconChain\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalCredentialsNotForEigenPod\",\"inputs\":[]}]",
}

// EigenPodStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use EigenPodStorageMetaData.ABI instead.
var EigenPodStorageABI = EigenPodStorageMetaData.ABI

// EigenPodStorage is an auto generated Go binding around an Ethereum contract.
type EigenPodStorage struct {
	EigenPodStorageCaller     // Read-only binding to the contract
	EigenPodStorageTransactor // Write-only binding to the contract
	EigenPodStorageFilterer   // Log filterer for contract events
}

// EigenPodStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type EigenPodStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenPodStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EigenPodStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenPodStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EigenPodStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EigenPodStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EigenPodStorageSession struct {
	Contract     *EigenPodStorage  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EigenPodStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EigenPodStorageCallerSession struct {
	Contract *EigenPodStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EigenPodStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EigenPodStorageTransactorSession struct {
	Contract     *EigenPodStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EigenPodStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type EigenPodStorageRaw struct {
	Contract *EigenPodStorage // Generic contract binding to access the raw methods on
}

// EigenPodStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EigenPodStorageCallerRaw struct {
	Contract *EigenPodStorageCaller // Generic read-only contract binding to access the raw methods on
}

// EigenPodStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EigenPodStorageTransactorRaw struct {
	Contract *EigenPodStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEigenPodStorage creates a new instance of EigenPodStorage, bound to a specific deployed contract.
func NewEigenPodStorage(address common.Address, backend bind.ContractBackend) (*EigenPodStorage, error) {
	contract, err := bindEigenPodStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EigenPodStorage{EigenPodStorageCaller: EigenPodStorageCaller{contract: contract}, EigenPodStorageTransactor: EigenPodStorageTransactor{contract: contract}, EigenPodStorageFilterer: EigenPodStorageFilterer{contract: contract}}, nil
}

// NewEigenPodStorageCaller creates a new read-only instance of EigenPodStorage, bound to a specific deployed contract.
func NewEigenPodStorageCaller(address common.Address, caller bind.ContractCaller) (*EigenPodStorageCaller, error) {
	contract, err := bindEigenPodStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageCaller{contract: contract}, nil
}

// NewEigenPodStorageTransactor creates a new write-only instance of EigenPodStorage, bound to a specific deployed contract.
func NewEigenPodStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*EigenPodStorageTransactor, error) {
	contract, err := bindEigenPodStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageTransactor{contract: contract}, nil
}

// NewEigenPodStorageFilterer creates a new log filterer instance of EigenPodStorage, bound to a specific deployed contract.
func NewEigenPodStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*EigenPodStorageFilterer, error) {
	contract, err := bindEigenPodStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageFilterer{contract: contract}, nil
}

// bindEigenPodStorage binds a generic wrapper to an already deployed contract.
func bindEigenPodStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EigenPodStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EigenPodStorage *EigenPodStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EigenPodStorage.Contract.EigenPodStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EigenPodStorage *EigenPodStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.EigenPodStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EigenPodStorage *EigenPodStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.EigenPodStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EigenPodStorage *EigenPodStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EigenPodStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EigenPodStorage *EigenPodStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EigenPodStorage *EigenPodStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.contract.Transact(opts, method, params...)
}

// ActiveValidatorCount is a free data retrieval call binding the contract method 0x2340e8d3.
//
// Solidity: function activeValidatorCount() view returns(uint256)
func (_EigenPodStorage *EigenPodStorageCaller) ActiveValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "activeValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveValidatorCount is a free data retrieval call binding the contract method 0x2340e8d3.
//
// Solidity: function activeValidatorCount() view returns(uint256)
func (_EigenPodStorage *EigenPodStorageSession) ActiveValidatorCount() (*big.Int, error) {
	return _EigenPodStorage.Contract.ActiveValidatorCount(&_EigenPodStorage.CallOpts)
}

// ActiveValidatorCount is a free data retrieval call binding the contract method 0x2340e8d3.
//
// Solidity: function activeValidatorCount() view returns(uint256)
func (_EigenPodStorage *EigenPodStorageCallerSession) ActiveValidatorCount() (*big.Int, error) {
	return _EigenPodStorage.Contract.ActiveValidatorCount(&_EigenPodStorage.CallOpts)
}

// CheckpointBalanceExitedGwei is a free data retrieval call binding the contract method 0x52396a59.
//
// Solidity: function checkpointBalanceExitedGwei(uint64 ) view returns(uint64)
func (_EigenPodStorage *EigenPodStorageCaller) CheckpointBalanceExitedGwei(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "checkpointBalanceExitedGwei", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CheckpointBalanceExitedGwei is a free data retrieval call binding the contract method 0x52396a59.
//
// Solidity: function checkpointBalanceExitedGwei(uint64 ) view returns(uint64)
func (_EigenPodStorage *EigenPodStorageSession) CheckpointBalanceExitedGwei(arg0 uint64) (uint64, error) {
	return _EigenPodStorage.Contract.CheckpointBalanceExitedGwei(&_EigenPodStorage.CallOpts, arg0)
}

// CheckpointBalanceExitedGwei is a free data retrieval call binding the contract method 0x52396a59.
//
// Solidity: function checkpointBalanceExitedGwei(uint64 ) view returns(uint64)
func (_EigenPodStorage *EigenPodStorageCallerSession) CheckpointBalanceExitedGwei(arg0 uint64) (uint64, error) {
	return _EigenPodStorage.Contract.CheckpointBalanceExitedGwei(&_EigenPodStorage.CallOpts, arg0)
}

// CurrentCheckpoint is a free data retrieval call binding the contract method 0x47d28372.
//
// Solidity: function currentCheckpoint() view returns((bytes32,uint24,uint64,int64,uint64))
func (_EigenPodStorage *EigenPodStorageCaller) CurrentCheckpoint(opts *bind.CallOpts) (IEigenPodTypesCheckpoint, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "currentCheckpoint")

	if err != nil {
		return *new(IEigenPodTypesCheckpoint), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodTypesCheckpoint)).(*IEigenPodTypesCheckpoint)

	return out0, err

}

// CurrentCheckpoint is a free data retrieval call binding the contract method 0x47d28372.
//
// Solidity: function currentCheckpoint() view returns((bytes32,uint24,uint64,int64,uint64))
func (_EigenPodStorage *EigenPodStorageSession) CurrentCheckpoint() (IEigenPodTypesCheckpoint, error) {
	return _EigenPodStorage.Contract.CurrentCheckpoint(&_EigenPodStorage.CallOpts)
}

// CurrentCheckpoint is a free data retrieval call binding the contract method 0x47d28372.
//
// Solidity: function currentCheckpoint() view returns((bytes32,uint24,uint64,int64,uint64))
func (_EigenPodStorage *EigenPodStorageCallerSession) CurrentCheckpoint() (IEigenPodTypesCheckpoint, error) {
	return _EigenPodStorage.Contract.CurrentCheckpoint(&_EigenPodStorage.CallOpts)
}

// CurrentCheckpointTimestamp is a free data retrieval call binding the contract method 0x42ecff2a.
//
// Solidity: function currentCheckpointTimestamp() view returns(uint64)
func (_EigenPodStorage *EigenPodStorageCaller) CurrentCheckpointTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "currentCheckpointTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CurrentCheckpointTimestamp is a free data retrieval call binding the contract method 0x42ecff2a.
//
// Solidity: function currentCheckpointTimestamp() view returns(uint64)
func (_EigenPodStorage *EigenPodStorageSession) CurrentCheckpointTimestamp() (uint64, error) {
	return _EigenPodStorage.Contract.CurrentCheckpointTimestamp(&_EigenPodStorage.CallOpts)
}

// CurrentCheckpointTimestamp is a free data retrieval call binding the contract method 0x42ecff2a.
//
// Solidity: function currentCheckpointTimestamp() view returns(uint64)
func (_EigenPodStorage *EigenPodStorageCallerSession) CurrentCheckpointTimestamp() (uint64, error) {
	return _EigenPodStorage.Contract.CurrentCheckpointTimestamp(&_EigenPodStorage.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_EigenPodStorage *EigenPodStorageCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_EigenPodStorage *EigenPodStorageSession) EigenPodManager() (common.Address, error) {
	return _EigenPodStorage.Contract.EigenPodManager(&_EigenPodStorage.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_EigenPodStorage *EigenPodStorageCallerSession) EigenPodManager() (common.Address, error) {
	return _EigenPodStorage.Contract.EigenPodManager(&_EigenPodStorage.CallOpts)
}

// GetConsolidationRequestFee is a free data retrieval call binding the contract method 0x1e515533.
//
// Solidity: function getConsolidationRequestFee() view returns(uint256)
func (_EigenPodStorage *EigenPodStorageCaller) GetConsolidationRequestFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "getConsolidationRequestFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConsolidationRequestFee is a free data retrieval call binding the contract method 0x1e515533.
//
// Solidity: function getConsolidationRequestFee() view returns(uint256)
func (_EigenPodStorage *EigenPodStorageSession) GetConsolidationRequestFee() (*big.Int, error) {
	return _EigenPodStorage.Contract.GetConsolidationRequestFee(&_EigenPodStorage.CallOpts)
}

// GetConsolidationRequestFee is a free data retrieval call binding the contract method 0x1e515533.
//
// Solidity: function getConsolidationRequestFee() view returns(uint256)
func (_EigenPodStorage *EigenPodStorageCallerSession) GetConsolidationRequestFee() (*big.Int, error) {
	return _EigenPodStorage.Contract.GetConsolidationRequestFee(&_EigenPodStorage.CallOpts)
}

// GetParentBlockRoot is a free data retrieval call binding the contract method 0x6c0d2d5a.
//
// Solidity: function getParentBlockRoot(uint64 timestamp) view returns(bytes32)
func (_EigenPodStorage *EigenPodStorageCaller) GetParentBlockRoot(opts *bind.CallOpts, timestamp uint64) ([32]byte, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "getParentBlockRoot", timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetParentBlockRoot is a free data retrieval call binding the contract method 0x6c0d2d5a.
//
// Solidity: function getParentBlockRoot(uint64 timestamp) view returns(bytes32)
func (_EigenPodStorage *EigenPodStorageSession) GetParentBlockRoot(timestamp uint64) ([32]byte, error) {
	return _EigenPodStorage.Contract.GetParentBlockRoot(&_EigenPodStorage.CallOpts, timestamp)
}

// GetParentBlockRoot is a free data retrieval call binding the contract method 0x6c0d2d5a.
//
// Solidity: function getParentBlockRoot(uint64 timestamp) view returns(bytes32)
func (_EigenPodStorage *EigenPodStorageCallerSession) GetParentBlockRoot(timestamp uint64) ([32]byte, error) {
	return _EigenPodStorage.Contract.GetParentBlockRoot(&_EigenPodStorage.CallOpts, timestamp)
}

// GetWithdrawalRequestFee is a free data retrieval call binding the contract method 0xc44e30dc.
//
// Solidity: function getWithdrawalRequestFee() view returns(uint256)
func (_EigenPodStorage *EigenPodStorageCaller) GetWithdrawalRequestFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "getWithdrawalRequestFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawalRequestFee is a free data retrieval call binding the contract method 0xc44e30dc.
//
// Solidity: function getWithdrawalRequestFee() view returns(uint256)
func (_EigenPodStorage *EigenPodStorageSession) GetWithdrawalRequestFee() (*big.Int, error) {
	return _EigenPodStorage.Contract.GetWithdrawalRequestFee(&_EigenPodStorage.CallOpts)
}

// GetWithdrawalRequestFee is a free data retrieval call binding the contract method 0xc44e30dc.
//
// Solidity: function getWithdrawalRequestFee() view returns(uint256)
func (_EigenPodStorage *EigenPodStorageCallerSession) GetWithdrawalRequestFee() (*big.Int, error) {
	return _EigenPodStorage.Contract.GetWithdrawalRequestFee(&_EigenPodStorage.CallOpts)
}

// LastCheckpointTimestamp is a free data retrieval call binding the contract method 0xee94d67c.
//
// Solidity: function lastCheckpointTimestamp() view returns(uint64)
func (_EigenPodStorage *EigenPodStorageCaller) LastCheckpointTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "lastCheckpointTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastCheckpointTimestamp is a free data retrieval call binding the contract method 0xee94d67c.
//
// Solidity: function lastCheckpointTimestamp() view returns(uint64)
func (_EigenPodStorage *EigenPodStorageSession) LastCheckpointTimestamp() (uint64, error) {
	return _EigenPodStorage.Contract.LastCheckpointTimestamp(&_EigenPodStorage.CallOpts)
}

// LastCheckpointTimestamp is a free data retrieval call binding the contract method 0xee94d67c.
//
// Solidity: function lastCheckpointTimestamp() view returns(uint64)
func (_EigenPodStorage *EigenPodStorageCallerSession) LastCheckpointTimestamp() (uint64, error) {
	return _EigenPodStorage.Contract.LastCheckpointTimestamp(&_EigenPodStorage.CallOpts)
}

// PodOwner is a free data retrieval call binding the contract method 0x0b18ff66.
//
// Solidity: function podOwner() view returns(address)
func (_EigenPodStorage *EigenPodStorageCaller) PodOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "podOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PodOwner is a free data retrieval call binding the contract method 0x0b18ff66.
//
// Solidity: function podOwner() view returns(address)
func (_EigenPodStorage *EigenPodStorageSession) PodOwner() (common.Address, error) {
	return _EigenPodStorage.Contract.PodOwner(&_EigenPodStorage.CallOpts)
}

// PodOwner is a free data retrieval call binding the contract method 0x0b18ff66.
//
// Solidity: function podOwner() view returns(address)
func (_EigenPodStorage *EigenPodStorageCallerSession) PodOwner() (common.Address, error) {
	return _EigenPodStorage.Contract.PodOwner(&_EigenPodStorage.CallOpts)
}

// ProofSubmitter is a free data retrieval call binding the contract method 0x58753357.
//
// Solidity: function proofSubmitter() view returns(address)
func (_EigenPodStorage *EigenPodStorageCaller) ProofSubmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "proofSubmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProofSubmitter is a free data retrieval call binding the contract method 0x58753357.
//
// Solidity: function proofSubmitter() view returns(address)
func (_EigenPodStorage *EigenPodStorageSession) ProofSubmitter() (common.Address, error) {
	return _EigenPodStorage.Contract.ProofSubmitter(&_EigenPodStorage.CallOpts)
}

// ProofSubmitter is a free data retrieval call binding the contract method 0x58753357.
//
// Solidity: function proofSubmitter() view returns(address)
func (_EigenPodStorage *EigenPodStorageCallerSession) ProofSubmitter() (common.Address, error) {
	return _EigenPodStorage.Contract.ProofSubmitter(&_EigenPodStorage.CallOpts)
}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_EigenPodStorage *EigenPodStorageCaller) ValidatorPubkeyHashToInfo(opts *bind.CallOpts, validatorPubkeyHash [32]byte) (IEigenPodTypesValidatorInfo, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "validatorPubkeyHashToInfo", validatorPubkeyHash)

	if err != nil {
		return *new(IEigenPodTypesValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodTypesValidatorInfo)).(*IEigenPodTypesValidatorInfo)

	return out0, err

}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_EigenPodStorage *EigenPodStorageSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodTypesValidatorInfo, error) {
	return _EigenPodStorage.Contract.ValidatorPubkeyHashToInfo(&_EigenPodStorage.CallOpts, validatorPubkeyHash)
}

// ValidatorPubkeyHashToInfo is a free data retrieval call binding the contract method 0x6fcd0e53.
//
// Solidity: function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) view returns((uint64,uint64,uint64,uint8))
func (_EigenPodStorage *EigenPodStorageCallerSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodTypesValidatorInfo, error) {
	return _EigenPodStorage.Contract.ValidatorPubkeyHashToInfo(&_EigenPodStorage.CallOpts, validatorPubkeyHash)
}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_EigenPodStorage *EigenPodStorageCaller) ValidatorPubkeyToInfo(opts *bind.CallOpts, validatorPubkey []byte) (IEigenPodTypesValidatorInfo, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "validatorPubkeyToInfo", validatorPubkey)

	if err != nil {
		return *new(IEigenPodTypesValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodTypesValidatorInfo)).(*IEigenPodTypesValidatorInfo)

	return out0, err

}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_EigenPodStorage *EigenPodStorageSession) ValidatorPubkeyToInfo(validatorPubkey []byte) (IEigenPodTypesValidatorInfo, error) {
	return _EigenPodStorage.Contract.ValidatorPubkeyToInfo(&_EigenPodStorage.CallOpts, validatorPubkey)
}

// ValidatorPubkeyToInfo is a free data retrieval call binding the contract method 0xb522538a.
//
// Solidity: function validatorPubkeyToInfo(bytes validatorPubkey) view returns((uint64,uint64,uint64,uint8))
func (_EigenPodStorage *EigenPodStorageCallerSession) ValidatorPubkeyToInfo(validatorPubkey []byte) (IEigenPodTypesValidatorInfo, error) {
	return _EigenPodStorage.Contract.ValidatorPubkeyToInfo(&_EigenPodStorage.CallOpts, validatorPubkey)
}

// ValidatorStatus is a free data retrieval call binding the contract method 0x58eaee79.
//
// Solidity: function validatorStatus(bytes validatorPubkey) view returns(uint8)
func (_EigenPodStorage *EigenPodStorageCaller) ValidatorStatus(opts *bind.CallOpts, validatorPubkey []byte) (uint8, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "validatorStatus", validatorPubkey)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ValidatorStatus is a free data retrieval call binding the contract method 0x58eaee79.
//
// Solidity: function validatorStatus(bytes validatorPubkey) view returns(uint8)
func (_EigenPodStorage *EigenPodStorageSession) ValidatorStatus(validatorPubkey []byte) (uint8, error) {
	return _EigenPodStorage.Contract.ValidatorStatus(&_EigenPodStorage.CallOpts, validatorPubkey)
}

// ValidatorStatus is a free data retrieval call binding the contract method 0x58eaee79.
//
// Solidity: function validatorStatus(bytes validatorPubkey) view returns(uint8)
func (_EigenPodStorage *EigenPodStorageCallerSession) ValidatorStatus(validatorPubkey []byte) (uint8, error) {
	return _EigenPodStorage.Contract.ValidatorStatus(&_EigenPodStorage.CallOpts, validatorPubkey)
}

// ValidatorStatus0 is a free data retrieval call binding the contract method 0x7439841f.
//
// Solidity: function validatorStatus(bytes32 pubkeyHash) view returns(uint8)
func (_EigenPodStorage *EigenPodStorageCaller) ValidatorStatus0(opts *bind.CallOpts, pubkeyHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "validatorStatus0", pubkeyHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ValidatorStatus0 is a free data retrieval call binding the contract method 0x7439841f.
//
// Solidity: function validatorStatus(bytes32 pubkeyHash) view returns(uint8)
func (_EigenPodStorage *EigenPodStorageSession) ValidatorStatus0(pubkeyHash [32]byte) (uint8, error) {
	return _EigenPodStorage.Contract.ValidatorStatus0(&_EigenPodStorage.CallOpts, pubkeyHash)
}

// ValidatorStatus0 is a free data retrieval call binding the contract method 0x7439841f.
//
// Solidity: function validatorStatus(bytes32 pubkeyHash) view returns(uint8)
func (_EigenPodStorage *EigenPodStorageCallerSession) ValidatorStatus0(pubkeyHash [32]byte) (uint8, error) {
	return _EigenPodStorage.Contract.ValidatorStatus0(&_EigenPodStorage.CallOpts, pubkeyHash)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_EigenPodStorage *EigenPodStorageCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_EigenPodStorage *EigenPodStorageSession) Version() (string, error) {
	return _EigenPodStorage.Contract.Version(&_EigenPodStorage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_EigenPodStorage *EigenPodStorageCallerSession) Version() (string, error) {
	return _EigenPodStorage.Contract.Version(&_EigenPodStorage.CallOpts)
}

// WithdrawableRestakedExecutionLayerGwei is a free data retrieval call binding the contract method 0x3474aa16.
//
// Solidity: function withdrawableRestakedExecutionLayerGwei() view returns(uint64)
func (_EigenPodStorage *EigenPodStorageCaller) WithdrawableRestakedExecutionLayerGwei(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EigenPodStorage.contract.Call(opts, &out, "withdrawableRestakedExecutionLayerGwei")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// WithdrawableRestakedExecutionLayerGwei is a free data retrieval call binding the contract method 0x3474aa16.
//
// Solidity: function withdrawableRestakedExecutionLayerGwei() view returns(uint64)
func (_EigenPodStorage *EigenPodStorageSession) WithdrawableRestakedExecutionLayerGwei() (uint64, error) {
	return _EigenPodStorage.Contract.WithdrawableRestakedExecutionLayerGwei(&_EigenPodStorage.CallOpts)
}

// WithdrawableRestakedExecutionLayerGwei is a free data retrieval call binding the contract method 0x3474aa16.
//
// Solidity: function withdrawableRestakedExecutionLayerGwei() view returns(uint64)
func (_EigenPodStorage *EigenPodStorageCallerSession) WithdrawableRestakedExecutionLayerGwei() (uint64, error) {
	return _EigenPodStorage.Contract.WithdrawableRestakedExecutionLayerGwei(&_EigenPodStorage.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_EigenPodStorage *EigenPodStorageTransactor) Initialize(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _EigenPodStorage.contract.Transact(opts, "initialize", owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_EigenPodStorage *EigenPodStorageSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.Initialize(&_EigenPodStorage.TransactOpts, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_EigenPodStorage *EigenPodStorageTransactorSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.Initialize(&_EigenPodStorage.TransactOpts, owner)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_EigenPodStorage *EigenPodStorageTransactor) RecoverTokens(opts *bind.TransactOpts, tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _EigenPodStorage.contract.Transact(opts, "recoverTokens", tokenList, amountsToWithdraw, recipient)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_EigenPodStorage *EigenPodStorageSession) RecoverTokens(tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.RecoverTokens(&_EigenPodStorage.TransactOpts, tokenList, amountsToWithdraw, recipient)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_EigenPodStorage *EigenPodStorageTransactorSession) RecoverTokens(tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.RecoverTokens(&_EigenPodStorage.TransactOpts, tokenList, amountsToWithdraw, recipient)
}

// RequestConsolidation is a paid mutator transaction binding the contract method 0x6691954e.
//
// Solidity: function requestConsolidation((bytes,bytes)[] requests) payable returns()
func (_EigenPodStorage *EigenPodStorageTransactor) RequestConsolidation(opts *bind.TransactOpts, requests []IEigenPodTypesConsolidationRequest) (*types.Transaction, error) {
	return _EigenPodStorage.contract.Transact(opts, "requestConsolidation", requests)
}

// RequestConsolidation is a paid mutator transaction binding the contract method 0x6691954e.
//
// Solidity: function requestConsolidation((bytes,bytes)[] requests) payable returns()
func (_EigenPodStorage *EigenPodStorageSession) RequestConsolidation(requests []IEigenPodTypesConsolidationRequest) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.RequestConsolidation(&_EigenPodStorage.TransactOpts, requests)
}

// RequestConsolidation is a paid mutator transaction binding the contract method 0x6691954e.
//
// Solidity: function requestConsolidation((bytes,bytes)[] requests) payable returns()
func (_EigenPodStorage *EigenPodStorageTransactorSession) RequestConsolidation(requests []IEigenPodTypesConsolidationRequest) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.RequestConsolidation(&_EigenPodStorage.TransactOpts, requests)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x3f5fa57a.
//
// Solidity: function requestWithdrawal((bytes,uint64)[] requests) payable returns()
func (_EigenPodStorage *EigenPodStorageTransactor) RequestWithdrawal(opts *bind.TransactOpts, requests []IEigenPodTypesWithdrawalRequest) (*types.Transaction, error) {
	return _EigenPodStorage.contract.Transact(opts, "requestWithdrawal", requests)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x3f5fa57a.
//
// Solidity: function requestWithdrawal((bytes,uint64)[] requests) payable returns()
func (_EigenPodStorage *EigenPodStorageSession) RequestWithdrawal(requests []IEigenPodTypesWithdrawalRequest) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.RequestWithdrawal(&_EigenPodStorage.TransactOpts, requests)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x3f5fa57a.
//
// Solidity: function requestWithdrawal((bytes,uint64)[] requests) payable returns()
func (_EigenPodStorage *EigenPodStorageTransactorSession) RequestWithdrawal(requests []IEigenPodTypesWithdrawalRequest) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.RequestWithdrawal(&_EigenPodStorage.TransactOpts, requests)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xd06d5587.
//
// Solidity: function setProofSubmitter(address newProofSubmitter) returns()
func (_EigenPodStorage *EigenPodStorageTransactor) SetProofSubmitter(opts *bind.TransactOpts, newProofSubmitter common.Address) (*types.Transaction, error) {
	return _EigenPodStorage.contract.Transact(opts, "setProofSubmitter", newProofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xd06d5587.
//
// Solidity: function setProofSubmitter(address newProofSubmitter) returns()
func (_EigenPodStorage *EigenPodStorageSession) SetProofSubmitter(newProofSubmitter common.Address) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.SetProofSubmitter(&_EigenPodStorage.TransactOpts, newProofSubmitter)
}

// SetProofSubmitter is a paid mutator transaction binding the contract method 0xd06d5587.
//
// Solidity: function setProofSubmitter(address newProofSubmitter) returns()
func (_EigenPodStorage *EigenPodStorageTransactorSession) SetProofSubmitter(newProofSubmitter common.Address) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.SetProofSubmitter(&_EigenPodStorage.TransactOpts, newProofSubmitter)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_EigenPodStorage *EigenPodStorageTransactor) Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _EigenPodStorage.contract.Transact(opts, "stake", pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_EigenPodStorage *EigenPodStorageSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.Stake(&_EigenPodStorage.TransactOpts, pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_EigenPodStorage *EigenPodStorageTransactorSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.Stake(&_EigenPodStorage.TransactOpts, pubkey, signature, depositDataRoot)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x88676cad.
//
// Solidity: function startCheckpoint(bool revertIfNoBalance) returns()
func (_EigenPodStorage *EigenPodStorageTransactor) StartCheckpoint(opts *bind.TransactOpts, revertIfNoBalance bool) (*types.Transaction, error) {
	return _EigenPodStorage.contract.Transact(opts, "startCheckpoint", revertIfNoBalance)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x88676cad.
//
// Solidity: function startCheckpoint(bool revertIfNoBalance) returns()
func (_EigenPodStorage *EigenPodStorageSession) StartCheckpoint(revertIfNoBalance bool) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.StartCheckpoint(&_EigenPodStorage.TransactOpts, revertIfNoBalance)
}

// StartCheckpoint is a paid mutator transaction binding the contract method 0x88676cad.
//
// Solidity: function startCheckpoint(bool revertIfNoBalance) returns()
func (_EigenPodStorage *EigenPodStorageTransactorSession) StartCheckpoint(revertIfNoBalance bool) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.StartCheckpoint(&_EigenPodStorage.TransactOpts, revertIfNoBalance)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0xf074ba62.
//
// Solidity: function verifyCheckpointProofs((bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_EigenPodStorage *EigenPodStorageTransactor) VerifyCheckpointProofs(opts *bind.TransactOpts, balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _EigenPodStorage.contract.Transact(opts, "verifyCheckpointProofs", balanceContainerProof, proofs)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0xf074ba62.
//
// Solidity: function verifyCheckpointProofs((bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_EigenPodStorage *EigenPodStorageSession) VerifyCheckpointProofs(balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.VerifyCheckpointProofs(&_EigenPodStorage.TransactOpts, balanceContainerProof, proofs)
}

// VerifyCheckpointProofs is a paid mutator transaction binding the contract method 0xf074ba62.
//
// Solidity: function verifyCheckpointProofs((bytes32,bytes) balanceContainerProof, (bytes32,bytes32,bytes)[] proofs) returns()
func (_EigenPodStorage *EigenPodStorageTransactorSession) VerifyCheckpointProofs(balanceContainerProof BeaconChainProofsBalanceContainerProof, proofs []BeaconChainProofsBalanceProof) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.VerifyCheckpointProofs(&_EigenPodStorage.TransactOpts, balanceContainerProof, proofs)
}

// VerifyStaleBalance is a paid mutator transaction binding the contract method 0x039157d2.
//
// Solidity: function verifyStaleBalance(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, (bytes32[],bytes) proof) returns()
func (_EigenPodStorage *EigenPodStorageTransactor) VerifyStaleBalance(opts *bind.TransactOpts, beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, proof BeaconChainProofsValidatorProof) (*types.Transaction, error) {
	return _EigenPodStorage.contract.Transact(opts, "verifyStaleBalance", beaconTimestamp, stateRootProof, proof)
}

// VerifyStaleBalance is a paid mutator transaction binding the contract method 0x039157d2.
//
// Solidity: function verifyStaleBalance(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, (bytes32[],bytes) proof) returns()
func (_EigenPodStorage *EigenPodStorageSession) VerifyStaleBalance(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, proof BeaconChainProofsValidatorProof) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.VerifyStaleBalance(&_EigenPodStorage.TransactOpts, beaconTimestamp, stateRootProof, proof)
}

// VerifyStaleBalance is a paid mutator transaction binding the contract method 0x039157d2.
//
// Solidity: function verifyStaleBalance(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, (bytes32[],bytes) proof) returns()
func (_EigenPodStorage *EigenPodStorageTransactorSession) VerifyStaleBalance(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, proof BeaconChainProofsValidatorProof) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.VerifyStaleBalance(&_EigenPodStorage.TransactOpts, beaconTimestamp, stateRootProof, proof)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_EigenPodStorage *EigenPodStorageTransactor) VerifyWithdrawalCredentials(opts *bind.TransactOpts, beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _EigenPodStorage.contract.Transact(opts, "verifyWithdrawalCredentials", beaconTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_EigenPodStorage *EigenPodStorageSession) VerifyWithdrawalCredentials(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.VerifyWithdrawalCredentials(&_EigenPodStorage.TransactOpts, beaconTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 beaconTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_EigenPodStorage *EigenPodStorageTransactorSession) VerifyWithdrawalCredentials(beaconTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.VerifyWithdrawalCredentials(&_EigenPodStorage.TransactOpts, beaconTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// WithdrawRestakedBeaconChainETH is a paid mutator transaction binding the contract method 0xc4907442.
//
// Solidity: function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) returns()
func (_EigenPodStorage *EigenPodStorageTransactor) WithdrawRestakedBeaconChainETH(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EigenPodStorage.contract.Transact(opts, "withdrawRestakedBeaconChainETH", recipient, amount)
}

// WithdrawRestakedBeaconChainETH is a paid mutator transaction binding the contract method 0xc4907442.
//
// Solidity: function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) returns()
func (_EigenPodStorage *EigenPodStorageSession) WithdrawRestakedBeaconChainETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.WithdrawRestakedBeaconChainETH(&_EigenPodStorage.TransactOpts, recipient, amount)
}

// WithdrawRestakedBeaconChainETH is a paid mutator transaction binding the contract method 0xc4907442.
//
// Solidity: function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) returns()
func (_EigenPodStorage *EigenPodStorageTransactorSession) WithdrawRestakedBeaconChainETH(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EigenPodStorage.Contract.WithdrawRestakedBeaconChainETH(&_EigenPodStorage.TransactOpts, recipient, amount)
}

// EigenPodStorageCheckpointCreatedIterator is returned from FilterCheckpointCreated and is used to iterate over the raw logs and unpacked data for CheckpointCreated events raised by the EigenPodStorage contract.
type EigenPodStorageCheckpointCreatedIterator struct {
	Event *EigenPodStorageCheckpointCreated // Event containing the contract specifics and raw log

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
func (it *EigenPodStorageCheckpointCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodStorageCheckpointCreated)
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
		it.Event = new(EigenPodStorageCheckpointCreated)
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
func (it *EigenPodStorageCheckpointCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodStorageCheckpointCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodStorageCheckpointCreated represents a CheckpointCreated event raised by the EigenPodStorage contract.
type EigenPodStorageCheckpointCreated struct {
	CheckpointTimestamp uint64
	BeaconBlockRoot     [32]byte
	ValidatorCount      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCheckpointCreated is a free log retrieval operation binding the contract event 0x575796133bbed337e5b39aa49a30dc2556a91e0c6c2af4b7b886ae77ebef1076.
//
// Solidity: event CheckpointCreated(uint64 indexed checkpointTimestamp, bytes32 indexed beaconBlockRoot, uint256 validatorCount)
func (_EigenPodStorage *EigenPodStorageFilterer) FilterCheckpointCreated(opts *bind.FilterOpts, checkpointTimestamp []uint64, beaconBlockRoot [][32]byte) (*EigenPodStorageCheckpointCreatedIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var beaconBlockRootRule []interface{}
	for _, beaconBlockRootItem := range beaconBlockRoot {
		beaconBlockRootRule = append(beaconBlockRootRule, beaconBlockRootItem)
	}

	logs, sub, err := _EigenPodStorage.contract.FilterLogs(opts, "CheckpointCreated", checkpointTimestampRule, beaconBlockRootRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageCheckpointCreatedIterator{contract: _EigenPodStorage.contract, event: "CheckpointCreated", logs: logs, sub: sub}, nil
}

// WatchCheckpointCreated is a free log subscription operation binding the contract event 0x575796133bbed337e5b39aa49a30dc2556a91e0c6c2af4b7b886ae77ebef1076.
//
// Solidity: event CheckpointCreated(uint64 indexed checkpointTimestamp, bytes32 indexed beaconBlockRoot, uint256 validatorCount)
func (_EigenPodStorage *EigenPodStorageFilterer) WatchCheckpointCreated(opts *bind.WatchOpts, sink chan<- *EigenPodStorageCheckpointCreated, checkpointTimestamp []uint64, beaconBlockRoot [][32]byte) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var beaconBlockRootRule []interface{}
	for _, beaconBlockRootItem := range beaconBlockRoot {
		beaconBlockRootRule = append(beaconBlockRootRule, beaconBlockRootItem)
	}

	logs, sub, err := _EigenPodStorage.contract.WatchLogs(opts, "CheckpointCreated", checkpointTimestampRule, beaconBlockRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodStorageCheckpointCreated)
				if err := _EigenPodStorage.contract.UnpackLog(event, "CheckpointCreated", log); err != nil {
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
func (_EigenPodStorage *EigenPodStorageFilterer) ParseCheckpointCreated(log types.Log) (*EigenPodStorageCheckpointCreated, error) {
	event := new(EigenPodStorageCheckpointCreated)
	if err := _EigenPodStorage.contract.UnpackLog(event, "CheckpointCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodStorageCheckpointFinalizedIterator is returned from FilterCheckpointFinalized and is used to iterate over the raw logs and unpacked data for CheckpointFinalized events raised by the EigenPodStorage contract.
type EigenPodStorageCheckpointFinalizedIterator struct {
	Event *EigenPodStorageCheckpointFinalized // Event containing the contract specifics and raw log

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
func (it *EigenPodStorageCheckpointFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodStorageCheckpointFinalized)
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
		it.Event = new(EigenPodStorageCheckpointFinalized)
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
func (it *EigenPodStorageCheckpointFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodStorageCheckpointFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodStorageCheckpointFinalized represents a CheckpointFinalized event raised by the EigenPodStorage contract.
type EigenPodStorageCheckpointFinalized struct {
	CheckpointTimestamp uint64
	TotalShareDeltaWei  *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCheckpointFinalized is a free log retrieval operation binding the contract event 0x525408c201bc1576eb44116f6478f1c2a54775b19a043bcfdc708364f74f8e44.
//
// Solidity: event CheckpointFinalized(uint64 indexed checkpointTimestamp, int256 totalShareDeltaWei)
func (_EigenPodStorage *EigenPodStorageFilterer) FilterCheckpointFinalized(opts *bind.FilterOpts, checkpointTimestamp []uint64) (*EigenPodStorageCheckpointFinalizedIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}

	logs, sub, err := _EigenPodStorage.contract.FilterLogs(opts, "CheckpointFinalized", checkpointTimestampRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageCheckpointFinalizedIterator{contract: _EigenPodStorage.contract, event: "CheckpointFinalized", logs: logs, sub: sub}, nil
}

// WatchCheckpointFinalized is a free log subscription operation binding the contract event 0x525408c201bc1576eb44116f6478f1c2a54775b19a043bcfdc708364f74f8e44.
//
// Solidity: event CheckpointFinalized(uint64 indexed checkpointTimestamp, int256 totalShareDeltaWei)
func (_EigenPodStorage *EigenPodStorageFilterer) WatchCheckpointFinalized(opts *bind.WatchOpts, sink chan<- *EigenPodStorageCheckpointFinalized, checkpointTimestamp []uint64) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}

	logs, sub, err := _EigenPodStorage.contract.WatchLogs(opts, "CheckpointFinalized", checkpointTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodStorageCheckpointFinalized)
				if err := _EigenPodStorage.contract.UnpackLog(event, "CheckpointFinalized", log); err != nil {
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
func (_EigenPodStorage *EigenPodStorageFilterer) ParseCheckpointFinalized(log types.Log) (*EigenPodStorageCheckpointFinalized, error) {
	event := new(EigenPodStorageCheckpointFinalized)
	if err := _EigenPodStorage.contract.UnpackLog(event, "CheckpointFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodStorageConsolidationRequestedIterator is returned from FilterConsolidationRequested and is used to iterate over the raw logs and unpacked data for ConsolidationRequested events raised by the EigenPodStorage contract.
type EigenPodStorageConsolidationRequestedIterator struct {
	Event *EigenPodStorageConsolidationRequested // Event containing the contract specifics and raw log

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
func (it *EigenPodStorageConsolidationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodStorageConsolidationRequested)
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
		it.Event = new(EigenPodStorageConsolidationRequested)
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
func (it *EigenPodStorageConsolidationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodStorageConsolidationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodStorageConsolidationRequested represents a ConsolidationRequested event raised by the EigenPodStorage contract.
type EigenPodStorageConsolidationRequested struct {
	SourcePubkeyHash [32]byte
	TargetPubkeyHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterConsolidationRequested is a free log retrieval operation binding the contract event 0x42f9c9db2ca443e9ec62f4588bd0c9b241065c02c2a8001ac164ae1282dc7b94.
//
// Solidity: event ConsolidationRequested(bytes32 indexed sourcePubkeyHash, bytes32 indexed targetPubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) FilterConsolidationRequested(opts *bind.FilterOpts, sourcePubkeyHash [][32]byte, targetPubkeyHash [][32]byte) (*EigenPodStorageConsolidationRequestedIterator, error) {

	var sourcePubkeyHashRule []interface{}
	for _, sourcePubkeyHashItem := range sourcePubkeyHash {
		sourcePubkeyHashRule = append(sourcePubkeyHashRule, sourcePubkeyHashItem)
	}
	var targetPubkeyHashRule []interface{}
	for _, targetPubkeyHashItem := range targetPubkeyHash {
		targetPubkeyHashRule = append(targetPubkeyHashRule, targetPubkeyHashItem)
	}

	logs, sub, err := _EigenPodStorage.contract.FilterLogs(opts, "ConsolidationRequested", sourcePubkeyHashRule, targetPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageConsolidationRequestedIterator{contract: _EigenPodStorage.contract, event: "ConsolidationRequested", logs: logs, sub: sub}, nil
}

// WatchConsolidationRequested is a free log subscription operation binding the contract event 0x42f9c9db2ca443e9ec62f4588bd0c9b241065c02c2a8001ac164ae1282dc7b94.
//
// Solidity: event ConsolidationRequested(bytes32 indexed sourcePubkeyHash, bytes32 indexed targetPubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) WatchConsolidationRequested(opts *bind.WatchOpts, sink chan<- *EigenPodStorageConsolidationRequested, sourcePubkeyHash [][32]byte, targetPubkeyHash [][32]byte) (event.Subscription, error) {

	var sourcePubkeyHashRule []interface{}
	for _, sourcePubkeyHashItem := range sourcePubkeyHash {
		sourcePubkeyHashRule = append(sourcePubkeyHashRule, sourcePubkeyHashItem)
	}
	var targetPubkeyHashRule []interface{}
	for _, targetPubkeyHashItem := range targetPubkeyHash {
		targetPubkeyHashRule = append(targetPubkeyHashRule, targetPubkeyHashItem)
	}

	logs, sub, err := _EigenPodStorage.contract.WatchLogs(opts, "ConsolidationRequested", sourcePubkeyHashRule, targetPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodStorageConsolidationRequested)
				if err := _EigenPodStorage.contract.UnpackLog(event, "ConsolidationRequested", log); err != nil {
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

// ParseConsolidationRequested is a log parse operation binding the contract event 0x42f9c9db2ca443e9ec62f4588bd0c9b241065c02c2a8001ac164ae1282dc7b94.
//
// Solidity: event ConsolidationRequested(bytes32 indexed sourcePubkeyHash, bytes32 indexed targetPubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) ParseConsolidationRequested(log types.Log) (*EigenPodStorageConsolidationRequested, error) {
	event := new(EigenPodStorageConsolidationRequested)
	if err := _EigenPodStorage.contract.UnpackLog(event, "ConsolidationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodStorageEigenPodStakedIterator is returned from FilterEigenPodStaked and is used to iterate over the raw logs and unpacked data for EigenPodStaked events raised by the EigenPodStorage contract.
type EigenPodStorageEigenPodStakedIterator struct {
	Event *EigenPodStorageEigenPodStaked // Event containing the contract specifics and raw log

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
func (it *EigenPodStorageEigenPodStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodStorageEigenPodStaked)
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
		it.Event = new(EigenPodStorageEigenPodStaked)
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
func (it *EigenPodStorageEigenPodStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodStorageEigenPodStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodStorageEigenPodStaked represents a EigenPodStaked event raised by the EigenPodStorage contract.
type EigenPodStorageEigenPodStaked struct {
	PubkeyHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEigenPodStaked is a free log retrieval operation binding the contract event 0xa01003766d3cd97cf2ade5429690bf5d206be7fb01ef9d3a0089ecf67bc11219.
//
// Solidity: event EigenPodStaked(bytes32 pubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) FilterEigenPodStaked(opts *bind.FilterOpts) (*EigenPodStorageEigenPodStakedIterator, error) {

	logs, sub, err := _EigenPodStorage.contract.FilterLogs(opts, "EigenPodStaked")
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageEigenPodStakedIterator{contract: _EigenPodStorage.contract, event: "EigenPodStaked", logs: logs, sub: sub}, nil
}

// WatchEigenPodStaked is a free log subscription operation binding the contract event 0xa01003766d3cd97cf2ade5429690bf5d206be7fb01ef9d3a0089ecf67bc11219.
//
// Solidity: event EigenPodStaked(bytes32 pubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) WatchEigenPodStaked(opts *bind.WatchOpts, sink chan<- *EigenPodStorageEigenPodStaked) (event.Subscription, error) {

	logs, sub, err := _EigenPodStorage.contract.WatchLogs(opts, "EigenPodStaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodStorageEigenPodStaked)
				if err := _EigenPodStorage.contract.UnpackLog(event, "EigenPodStaked", log); err != nil {
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

// ParseEigenPodStaked is a log parse operation binding the contract event 0xa01003766d3cd97cf2ade5429690bf5d206be7fb01ef9d3a0089ecf67bc11219.
//
// Solidity: event EigenPodStaked(bytes32 pubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) ParseEigenPodStaked(log types.Log) (*EigenPodStorageEigenPodStaked, error) {
	event := new(EigenPodStorageEigenPodStaked)
	if err := _EigenPodStorage.contract.UnpackLog(event, "EigenPodStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodStorageExitRequestedIterator is returned from FilterExitRequested and is used to iterate over the raw logs and unpacked data for ExitRequested events raised by the EigenPodStorage contract.
type EigenPodStorageExitRequestedIterator struct {
	Event *EigenPodStorageExitRequested // Event containing the contract specifics and raw log

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
func (it *EigenPodStorageExitRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodStorageExitRequested)
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
		it.Event = new(EigenPodStorageExitRequested)
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
func (it *EigenPodStorageExitRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodStorageExitRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodStorageExitRequested represents a ExitRequested event raised by the EigenPodStorage contract.
type EigenPodStorageExitRequested struct {
	ValidatorPubkeyHash [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterExitRequested is a free log retrieval operation binding the contract event 0x60d8ca014d4765a2b8b389e25714cb1cef83b574222911a01d90c1bd69d2d320.
//
// Solidity: event ExitRequested(bytes32 indexed validatorPubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) FilterExitRequested(opts *bind.FilterOpts, validatorPubkeyHash [][32]byte) (*EigenPodStorageExitRequestedIterator, error) {

	var validatorPubkeyHashRule []interface{}
	for _, validatorPubkeyHashItem := range validatorPubkeyHash {
		validatorPubkeyHashRule = append(validatorPubkeyHashRule, validatorPubkeyHashItem)
	}

	logs, sub, err := _EigenPodStorage.contract.FilterLogs(opts, "ExitRequested", validatorPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageExitRequestedIterator{contract: _EigenPodStorage.contract, event: "ExitRequested", logs: logs, sub: sub}, nil
}

// WatchExitRequested is a free log subscription operation binding the contract event 0x60d8ca014d4765a2b8b389e25714cb1cef83b574222911a01d90c1bd69d2d320.
//
// Solidity: event ExitRequested(bytes32 indexed validatorPubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) WatchExitRequested(opts *bind.WatchOpts, sink chan<- *EigenPodStorageExitRequested, validatorPubkeyHash [][32]byte) (event.Subscription, error) {

	var validatorPubkeyHashRule []interface{}
	for _, validatorPubkeyHashItem := range validatorPubkeyHash {
		validatorPubkeyHashRule = append(validatorPubkeyHashRule, validatorPubkeyHashItem)
	}

	logs, sub, err := _EigenPodStorage.contract.WatchLogs(opts, "ExitRequested", validatorPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodStorageExitRequested)
				if err := _EigenPodStorage.contract.UnpackLog(event, "ExitRequested", log); err != nil {
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

// ParseExitRequested is a log parse operation binding the contract event 0x60d8ca014d4765a2b8b389e25714cb1cef83b574222911a01d90c1bd69d2d320.
//
// Solidity: event ExitRequested(bytes32 indexed validatorPubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) ParseExitRequested(log types.Log) (*EigenPodStorageExitRequested, error) {
	event := new(EigenPodStorageExitRequested)
	if err := _EigenPodStorage.contract.UnpackLog(event, "ExitRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodStorageNonBeaconChainETHReceivedIterator is returned from FilterNonBeaconChainETHReceived and is used to iterate over the raw logs and unpacked data for NonBeaconChainETHReceived events raised by the EigenPodStorage contract.
type EigenPodStorageNonBeaconChainETHReceivedIterator struct {
	Event *EigenPodStorageNonBeaconChainETHReceived // Event containing the contract specifics and raw log

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
func (it *EigenPodStorageNonBeaconChainETHReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodStorageNonBeaconChainETHReceived)
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
		it.Event = new(EigenPodStorageNonBeaconChainETHReceived)
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
func (it *EigenPodStorageNonBeaconChainETHReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodStorageNonBeaconChainETHReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodStorageNonBeaconChainETHReceived represents a NonBeaconChainETHReceived event raised by the EigenPodStorage contract.
type EigenPodStorageNonBeaconChainETHReceived struct {
	AmountReceived *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNonBeaconChainETHReceived is a free log retrieval operation binding the contract event 0x6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49.
//
// Solidity: event NonBeaconChainETHReceived(uint256 amountReceived)
func (_EigenPodStorage *EigenPodStorageFilterer) FilterNonBeaconChainETHReceived(opts *bind.FilterOpts) (*EigenPodStorageNonBeaconChainETHReceivedIterator, error) {

	logs, sub, err := _EigenPodStorage.contract.FilterLogs(opts, "NonBeaconChainETHReceived")
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageNonBeaconChainETHReceivedIterator{contract: _EigenPodStorage.contract, event: "NonBeaconChainETHReceived", logs: logs, sub: sub}, nil
}

// WatchNonBeaconChainETHReceived is a free log subscription operation binding the contract event 0x6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49.
//
// Solidity: event NonBeaconChainETHReceived(uint256 amountReceived)
func (_EigenPodStorage *EigenPodStorageFilterer) WatchNonBeaconChainETHReceived(opts *bind.WatchOpts, sink chan<- *EigenPodStorageNonBeaconChainETHReceived) (event.Subscription, error) {

	logs, sub, err := _EigenPodStorage.contract.WatchLogs(opts, "NonBeaconChainETHReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodStorageNonBeaconChainETHReceived)
				if err := _EigenPodStorage.contract.UnpackLog(event, "NonBeaconChainETHReceived", log); err != nil {
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
func (_EigenPodStorage *EigenPodStorageFilterer) ParseNonBeaconChainETHReceived(log types.Log) (*EigenPodStorageNonBeaconChainETHReceived, error) {
	event := new(EigenPodStorageNonBeaconChainETHReceived)
	if err := _EigenPodStorage.contract.UnpackLog(event, "NonBeaconChainETHReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodStorageProofSubmitterUpdatedIterator is returned from FilterProofSubmitterUpdated and is used to iterate over the raw logs and unpacked data for ProofSubmitterUpdated events raised by the EigenPodStorage contract.
type EigenPodStorageProofSubmitterUpdatedIterator struct {
	Event *EigenPodStorageProofSubmitterUpdated // Event containing the contract specifics and raw log

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
func (it *EigenPodStorageProofSubmitterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodStorageProofSubmitterUpdated)
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
		it.Event = new(EigenPodStorageProofSubmitterUpdated)
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
func (it *EigenPodStorageProofSubmitterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodStorageProofSubmitterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodStorageProofSubmitterUpdated represents a ProofSubmitterUpdated event raised by the EigenPodStorage contract.
type EigenPodStorageProofSubmitterUpdated struct {
	PrevProofSubmitter common.Address
	NewProofSubmitter  common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProofSubmitterUpdated is a free log retrieval operation binding the contract event 0xfb8129080a19d34dceac04ba253fc50304dc86c729bd63cdca4a969ad19a5eac.
//
// Solidity: event ProofSubmitterUpdated(address prevProofSubmitter, address newProofSubmitter)
func (_EigenPodStorage *EigenPodStorageFilterer) FilterProofSubmitterUpdated(opts *bind.FilterOpts) (*EigenPodStorageProofSubmitterUpdatedIterator, error) {

	logs, sub, err := _EigenPodStorage.contract.FilterLogs(opts, "ProofSubmitterUpdated")
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageProofSubmitterUpdatedIterator{contract: _EigenPodStorage.contract, event: "ProofSubmitterUpdated", logs: logs, sub: sub}, nil
}

// WatchProofSubmitterUpdated is a free log subscription operation binding the contract event 0xfb8129080a19d34dceac04ba253fc50304dc86c729bd63cdca4a969ad19a5eac.
//
// Solidity: event ProofSubmitterUpdated(address prevProofSubmitter, address newProofSubmitter)
func (_EigenPodStorage *EigenPodStorageFilterer) WatchProofSubmitterUpdated(opts *bind.WatchOpts, sink chan<- *EigenPodStorageProofSubmitterUpdated) (event.Subscription, error) {

	logs, sub, err := _EigenPodStorage.contract.WatchLogs(opts, "ProofSubmitterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodStorageProofSubmitterUpdated)
				if err := _EigenPodStorage.contract.UnpackLog(event, "ProofSubmitterUpdated", log); err != nil {
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
func (_EigenPodStorage *EigenPodStorageFilterer) ParseProofSubmitterUpdated(log types.Log) (*EigenPodStorageProofSubmitterUpdated, error) {
	event := new(EigenPodStorageProofSubmitterUpdated)
	if err := _EigenPodStorage.contract.UnpackLog(event, "ProofSubmitterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodStorageRestakedBeaconChainETHWithdrawnIterator is returned from FilterRestakedBeaconChainETHWithdrawn and is used to iterate over the raw logs and unpacked data for RestakedBeaconChainETHWithdrawn events raised by the EigenPodStorage contract.
type EigenPodStorageRestakedBeaconChainETHWithdrawnIterator struct {
	Event *EigenPodStorageRestakedBeaconChainETHWithdrawn // Event containing the contract specifics and raw log

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
func (it *EigenPodStorageRestakedBeaconChainETHWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodStorageRestakedBeaconChainETHWithdrawn)
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
		it.Event = new(EigenPodStorageRestakedBeaconChainETHWithdrawn)
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
func (it *EigenPodStorageRestakedBeaconChainETHWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodStorageRestakedBeaconChainETHWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodStorageRestakedBeaconChainETHWithdrawn represents a RestakedBeaconChainETHWithdrawn event raised by the EigenPodStorage contract.
type EigenPodStorageRestakedBeaconChainETHWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRestakedBeaconChainETHWithdrawn is a free log retrieval operation binding the contract event 0x8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e.
//
// Solidity: event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount)
func (_EigenPodStorage *EigenPodStorageFilterer) FilterRestakedBeaconChainETHWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*EigenPodStorageRestakedBeaconChainETHWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EigenPodStorage.contract.FilterLogs(opts, "RestakedBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageRestakedBeaconChainETHWithdrawnIterator{contract: _EigenPodStorage.contract, event: "RestakedBeaconChainETHWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRestakedBeaconChainETHWithdrawn is a free log subscription operation binding the contract event 0x8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e.
//
// Solidity: event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount)
func (_EigenPodStorage *EigenPodStorageFilterer) WatchRestakedBeaconChainETHWithdrawn(opts *bind.WatchOpts, sink chan<- *EigenPodStorageRestakedBeaconChainETHWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EigenPodStorage.contract.WatchLogs(opts, "RestakedBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodStorageRestakedBeaconChainETHWithdrawn)
				if err := _EigenPodStorage.contract.UnpackLog(event, "RestakedBeaconChainETHWithdrawn", log); err != nil {
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
func (_EigenPodStorage *EigenPodStorageFilterer) ParseRestakedBeaconChainETHWithdrawn(log types.Log) (*EigenPodStorageRestakedBeaconChainETHWithdrawn, error) {
	event := new(EigenPodStorageRestakedBeaconChainETHWithdrawn)
	if err := _EigenPodStorage.contract.UnpackLog(event, "RestakedBeaconChainETHWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodStorageSwitchToCompoundingRequestedIterator is returned from FilterSwitchToCompoundingRequested and is used to iterate over the raw logs and unpacked data for SwitchToCompoundingRequested events raised by the EigenPodStorage contract.
type EigenPodStorageSwitchToCompoundingRequestedIterator struct {
	Event *EigenPodStorageSwitchToCompoundingRequested // Event containing the contract specifics and raw log

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
func (it *EigenPodStorageSwitchToCompoundingRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodStorageSwitchToCompoundingRequested)
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
		it.Event = new(EigenPodStorageSwitchToCompoundingRequested)
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
func (it *EigenPodStorageSwitchToCompoundingRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodStorageSwitchToCompoundingRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodStorageSwitchToCompoundingRequested represents a SwitchToCompoundingRequested event raised by the EigenPodStorage contract.
type EigenPodStorageSwitchToCompoundingRequested struct {
	ValidatorPubkeyHash [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSwitchToCompoundingRequested is a free log retrieval operation binding the contract event 0xc97b965b92ae7fd20095fe8eb7b99f81f95f8c4adffb22a19116d8eb2846b016.
//
// Solidity: event SwitchToCompoundingRequested(bytes32 indexed validatorPubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) FilterSwitchToCompoundingRequested(opts *bind.FilterOpts, validatorPubkeyHash [][32]byte) (*EigenPodStorageSwitchToCompoundingRequestedIterator, error) {

	var validatorPubkeyHashRule []interface{}
	for _, validatorPubkeyHashItem := range validatorPubkeyHash {
		validatorPubkeyHashRule = append(validatorPubkeyHashRule, validatorPubkeyHashItem)
	}

	logs, sub, err := _EigenPodStorage.contract.FilterLogs(opts, "SwitchToCompoundingRequested", validatorPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageSwitchToCompoundingRequestedIterator{contract: _EigenPodStorage.contract, event: "SwitchToCompoundingRequested", logs: logs, sub: sub}, nil
}

// WatchSwitchToCompoundingRequested is a free log subscription operation binding the contract event 0xc97b965b92ae7fd20095fe8eb7b99f81f95f8c4adffb22a19116d8eb2846b016.
//
// Solidity: event SwitchToCompoundingRequested(bytes32 indexed validatorPubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) WatchSwitchToCompoundingRequested(opts *bind.WatchOpts, sink chan<- *EigenPodStorageSwitchToCompoundingRequested, validatorPubkeyHash [][32]byte) (event.Subscription, error) {

	var validatorPubkeyHashRule []interface{}
	for _, validatorPubkeyHashItem := range validatorPubkeyHash {
		validatorPubkeyHashRule = append(validatorPubkeyHashRule, validatorPubkeyHashItem)
	}

	logs, sub, err := _EigenPodStorage.contract.WatchLogs(opts, "SwitchToCompoundingRequested", validatorPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodStorageSwitchToCompoundingRequested)
				if err := _EigenPodStorage.contract.UnpackLog(event, "SwitchToCompoundingRequested", log); err != nil {
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

// ParseSwitchToCompoundingRequested is a log parse operation binding the contract event 0xc97b965b92ae7fd20095fe8eb7b99f81f95f8c4adffb22a19116d8eb2846b016.
//
// Solidity: event SwitchToCompoundingRequested(bytes32 indexed validatorPubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) ParseSwitchToCompoundingRequested(log types.Log) (*EigenPodStorageSwitchToCompoundingRequested, error) {
	event := new(EigenPodStorageSwitchToCompoundingRequested)
	if err := _EigenPodStorage.contract.UnpackLog(event, "SwitchToCompoundingRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodStorageValidatorBalanceUpdatedIterator is returned from FilterValidatorBalanceUpdated and is used to iterate over the raw logs and unpacked data for ValidatorBalanceUpdated events raised by the EigenPodStorage contract.
type EigenPodStorageValidatorBalanceUpdatedIterator struct {
	Event *EigenPodStorageValidatorBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *EigenPodStorageValidatorBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodStorageValidatorBalanceUpdated)
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
		it.Event = new(EigenPodStorageValidatorBalanceUpdated)
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
func (it *EigenPodStorageValidatorBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodStorageValidatorBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodStorageValidatorBalanceUpdated represents a ValidatorBalanceUpdated event raised by the EigenPodStorage contract.
type EigenPodStorageValidatorBalanceUpdated struct {
	PubkeyHash              [32]byte
	BalanceTimestamp        uint64
	NewValidatorBalanceGwei uint64
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterValidatorBalanceUpdated is a free log retrieval operation binding the contract event 0xcdae700d7241bc027168c53cf6f889763b0a2c88a65d77fc13a8a9fef0d8605f.
//
// Solidity: event ValidatorBalanceUpdated(bytes32 pubkeyHash, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_EigenPodStorage *EigenPodStorageFilterer) FilterValidatorBalanceUpdated(opts *bind.FilterOpts) (*EigenPodStorageValidatorBalanceUpdatedIterator, error) {

	logs, sub, err := _EigenPodStorage.contract.FilterLogs(opts, "ValidatorBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageValidatorBalanceUpdatedIterator{contract: _EigenPodStorage.contract, event: "ValidatorBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorBalanceUpdated is a free log subscription operation binding the contract event 0xcdae700d7241bc027168c53cf6f889763b0a2c88a65d77fc13a8a9fef0d8605f.
//
// Solidity: event ValidatorBalanceUpdated(bytes32 pubkeyHash, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_EigenPodStorage *EigenPodStorageFilterer) WatchValidatorBalanceUpdated(opts *bind.WatchOpts, sink chan<- *EigenPodStorageValidatorBalanceUpdated) (event.Subscription, error) {

	logs, sub, err := _EigenPodStorage.contract.WatchLogs(opts, "ValidatorBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodStorageValidatorBalanceUpdated)
				if err := _EigenPodStorage.contract.UnpackLog(event, "ValidatorBalanceUpdated", log); err != nil {
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

// ParseValidatorBalanceUpdated is a log parse operation binding the contract event 0xcdae700d7241bc027168c53cf6f889763b0a2c88a65d77fc13a8a9fef0d8605f.
//
// Solidity: event ValidatorBalanceUpdated(bytes32 pubkeyHash, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei)
func (_EigenPodStorage *EigenPodStorageFilterer) ParseValidatorBalanceUpdated(log types.Log) (*EigenPodStorageValidatorBalanceUpdated, error) {
	event := new(EigenPodStorageValidatorBalanceUpdated)
	if err := _EigenPodStorage.contract.UnpackLog(event, "ValidatorBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodStorageValidatorCheckpointedIterator is returned from FilterValidatorCheckpointed and is used to iterate over the raw logs and unpacked data for ValidatorCheckpointed events raised by the EigenPodStorage contract.
type EigenPodStorageValidatorCheckpointedIterator struct {
	Event *EigenPodStorageValidatorCheckpointed // Event containing the contract specifics and raw log

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
func (it *EigenPodStorageValidatorCheckpointedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodStorageValidatorCheckpointed)
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
		it.Event = new(EigenPodStorageValidatorCheckpointed)
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
func (it *EigenPodStorageValidatorCheckpointedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodStorageValidatorCheckpointedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodStorageValidatorCheckpointed represents a ValidatorCheckpointed event raised by the EigenPodStorage contract.
type EigenPodStorageValidatorCheckpointed struct {
	CheckpointTimestamp uint64
	PubkeyHash          [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterValidatorCheckpointed is a free log retrieval operation binding the contract event 0xe4866335761a51dcaff766448ab0af6064291ee5dc94e68492bb9cd757c1e350.
//
// Solidity: event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, bytes32 indexed pubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) FilterValidatorCheckpointed(opts *bind.FilterOpts, checkpointTimestamp []uint64, pubkeyHash [][32]byte) (*EigenPodStorageValidatorCheckpointedIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}

	logs, sub, err := _EigenPodStorage.contract.FilterLogs(opts, "ValidatorCheckpointed", checkpointTimestampRule, pubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageValidatorCheckpointedIterator{contract: _EigenPodStorage.contract, event: "ValidatorCheckpointed", logs: logs, sub: sub}, nil
}

// WatchValidatorCheckpointed is a free log subscription operation binding the contract event 0xe4866335761a51dcaff766448ab0af6064291ee5dc94e68492bb9cd757c1e350.
//
// Solidity: event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, bytes32 indexed pubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) WatchValidatorCheckpointed(opts *bind.WatchOpts, sink chan<- *EigenPodStorageValidatorCheckpointed, checkpointTimestamp []uint64, pubkeyHash [][32]byte) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}

	logs, sub, err := _EigenPodStorage.contract.WatchLogs(opts, "ValidatorCheckpointed", checkpointTimestampRule, pubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodStorageValidatorCheckpointed)
				if err := _EigenPodStorage.contract.UnpackLog(event, "ValidatorCheckpointed", log); err != nil {
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

// ParseValidatorCheckpointed is a log parse operation binding the contract event 0xe4866335761a51dcaff766448ab0af6064291ee5dc94e68492bb9cd757c1e350.
//
// Solidity: event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, bytes32 indexed pubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) ParseValidatorCheckpointed(log types.Log) (*EigenPodStorageValidatorCheckpointed, error) {
	event := new(EigenPodStorageValidatorCheckpointed)
	if err := _EigenPodStorage.contract.UnpackLog(event, "ValidatorCheckpointed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodStorageValidatorRestakedIterator is returned from FilterValidatorRestaked and is used to iterate over the raw logs and unpacked data for ValidatorRestaked events raised by the EigenPodStorage contract.
type EigenPodStorageValidatorRestakedIterator struct {
	Event *EigenPodStorageValidatorRestaked // Event containing the contract specifics and raw log

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
func (it *EigenPodStorageValidatorRestakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodStorageValidatorRestaked)
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
		it.Event = new(EigenPodStorageValidatorRestaked)
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
func (it *EigenPodStorageValidatorRestakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodStorageValidatorRestakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodStorageValidatorRestaked represents a ValidatorRestaked event raised by the EigenPodStorage contract.
type EigenPodStorageValidatorRestaked struct {
	PubkeyHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorRestaked is a free log retrieval operation binding the contract event 0x101790c2993f6a4d962bd17c786126823ba1c4cf04ff4cccb2659d50fb20aee8.
//
// Solidity: event ValidatorRestaked(bytes32 pubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) FilterValidatorRestaked(opts *bind.FilterOpts) (*EigenPodStorageValidatorRestakedIterator, error) {

	logs, sub, err := _EigenPodStorage.contract.FilterLogs(opts, "ValidatorRestaked")
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageValidatorRestakedIterator{contract: _EigenPodStorage.contract, event: "ValidatorRestaked", logs: logs, sub: sub}, nil
}

// WatchValidatorRestaked is a free log subscription operation binding the contract event 0x101790c2993f6a4d962bd17c786126823ba1c4cf04ff4cccb2659d50fb20aee8.
//
// Solidity: event ValidatorRestaked(bytes32 pubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) WatchValidatorRestaked(opts *bind.WatchOpts, sink chan<- *EigenPodStorageValidatorRestaked) (event.Subscription, error) {

	logs, sub, err := _EigenPodStorage.contract.WatchLogs(opts, "ValidatorRestaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodStorageValidatorRestaked)
				if err := _EigenPodStorage.contract.UnpackLog(event, "ValidatorRestaked", log); err != nil {
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

// ParseValidatorRestaked is a log parse operation binding the contract event 0x101790c2993f6a4d962bd17c786126823ba1c4cf04ff4cccb2659d50fb20aee8.
//
// Solidity: event ValidatorRestaked(bytes32 pubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) ParseValidatorRestaked(log types.Log) (*EigenPodStorageValidatorRestaked, error) {
	event := new(EigenPodStorageValidatorRestaked)
	if err := _EigenPodStorage.contract.UnpackLog(event, "ValidatorRestaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodStorageValidatorWithdrawnIterator is returned from FilterValidatorWithdrawn and is used to iterate over the raw logs and unpacked data for ValidatorWithdrawn events raised by the EigenPodStorage contract.
type EigenPodStorageValidatorWithdrawnIterator struct {
	Event *EigenPodStorageValidatorWithdrawn // Event containing the contract specifics and raw log

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
func (it *EigenPodStorageValidatorWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodStorageValidatorWithdrawn)
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
		it.Event = new(EigenPodStorageValidatorWithdrawn)
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
func (it *EigenPodStorageValidatorWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodStorageValidatorWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodStorageValidatorWithdrawn represents a ValidatorWithdrawn event raised by the EigenPodStorage contract.
type EigenPodStorageValidatorWithdrawn struct {
	CheckpointTimestamp uint64
	PubkeyHash          [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterValidatorWithdrawn is a free log retrieval operation binding the contract event 0x5ce0aa04ae51d52da6e680fbe0336d2e2432f7c3dc2d4f3193204c57b9072107.
//
// Solidity: event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, bytes32 indexed pubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) FilterValidatorWithdrawn(opts *bind.FilterOpts, checkpointTimestamp []uint64, pubkeyHash [][32]byte) (*EigenPodStorageValidatorWithdrawnIterator, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}

	logs, sub, err := _EigenPodStorage.contract.FilterLogs(opts, "ValidatorWithdrawn", checkpointTimestampRule, pubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageValidatorWithdrawnIterator{contract: _EigenPodStorage.contract, event: "ValidatorWithdrawn", logs: logs, sub: sub}, nil
}

// WatchValidatorWithdrawn is a free log subscription operation binding the contract event 0x5ce0aa04ae51d52da6e680fbe0336d2e2432f7c3dc2d4f3193204c57b9072107.
//
// Solidity: event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, bytes32 indexed pubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) WatchValidatorWithdrawn(opts *bind.WatchOpts, sink chan<- *EigenPodStorageValidatorWithdrawn, checkpointTimestamp []uint64, pubkeyHash [][32]byte) (event.Subscription, error) {

	var checkpointTimestampRule []interface{}
	for _, checkpointTimestampItem := range checkpointTimestamp {
		checkpointTimestampRule = append(checkpointTimestampRule, checkpointTimestampItem)
	}
	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}

	logs, sub, err := _EigenPodStorage.contract.WatchLogs(opts, "ValidatorWithdrawn", checkpointTimestampRule, pubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodStorageValidatorWithdrawn)
				if err := _EigenPodStorage.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
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

// ParseValidatorWithdrawn is a log parse operation binding the contract event 0x5ce0aa04ae51d52da6e680fbe0336d2e2432f7c3dc2d4f3193204c57b9072107.
//
// Solidity: event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, bytes32 indexed pubkeyHash)
func (_EigenPodStorage *EigenPodStorageFilterer) ParseValidatorWithdrawn(log types.Log) (*EigenPodStorageValidatorWithdrawn, error) {
	event := new(EigenPodStorageValidatorWithdrawn)
	if err := _EigenPodStorage.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EigenPodStorageWithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the EigenPodStorage contract.
type EigenPodStorageWithdrawalRequestedIterator struct {
	Event *EigenPodStorageWithdrawalRequested // Event containing the contract specifics and raw log

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
func (it *EigenPodStorageWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EigenPodStorageWithdrawalRequested)
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
		it.Event = new(EigenPodStorageWithdrawalRequested)
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
func (it *EigenPodStorageWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EigenPodStorageWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EigenPodStorageWithdrawalRequested represents a WithdrawalRequested event raised by the EigenPodStorage contract.
type EigenPodStorageWithdrawalRequested struct {
	ValidatorPubkeyHash  [32]byte
	WithdrawalAmountGwei uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0x8b2737bb64ab2f2dc09552dfa1c250399e6a42c7ea9f0e1c658f5d65d708ec05.
//
// Solidity: event WithdrawalRequested(bytes32 indexed validatorPubkeyHash, uint64 withdrawalAmountGwei)
func (_EigenPodStorage *EigenPodStorageFilterer) FilterWithdrawalRequested(opts *bind.FilterOpts, validatorPubkeyHash [][32]byte) (*EigenPodStorageWithdrawalRequestedIterator, error) {

	var validatorPubkeyHashRule []interface{}
	for _, validatorPubkeyHashItem := range validatorPubkeyHash {
		validatorPubkeyHashRule = append(validatorPubkeyHashRule, validatorPubkeyHashItem)
	}

	logs, sub, err := _EigenPodStorage.contract.FilterLogs(opts, "WithdrawalRequested", validatorPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return &EigenPodStorageWithdrawalRequestedIterator{contract: _EigenPodStorage.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0x8b2737bb64ab2f2dc09552dfa1c250399e6a42c7ea9f0e1c658f5d65d708ec05.
//
// Solidity: event WithdrawalRequested(bytes32 indexed validatorPubkeyHash, uint64 withdrawalAmountGwei)
func (_EigenPodStorage *EigenPodStorageFilterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *EigenPodStorageWithdrawalRequested, validatorPubkeyHash [][32]byte) (event.Subscription, error) {

	var validatorPubkeyHashRule []interface{}
	for _, validatorPubkeyHashItem := range validatorPubkeyHash {
		validatorPubkeyHashRule = append(validatorPubkeyHashRule, validatorPubkeyHashItem)
	}

	logs, sub, err := _EigenPodStorage.contract.WatchLogs(opts, "WithdrawalRequested", validatorPubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EigenPodStorageWithdrawalRequested)
				if err := _EigenPodStorage.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
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

// ParseWithdrawalRequested is a log parse operation binding the contract event 0x8b2737bb64ab2f2dc09552dfa1c250399e6a42c7ea9f0e1c658f5d65d708ec05.
//
// Solidity: event WithdrawalRequested(bytes32 indexed validatorPubkeyHash, uint64 withdrawalAmountGwei)
func (_EigenPodStorage *EigenPodStorageFilterer) ParseWithdrawalRequested(log types.Log) (*EigenPodStorageWithdrawalRequested, error) {
	event := new(EigenPodStorageWithdrawalRequested)
	if err := _EigenPodStorage.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
