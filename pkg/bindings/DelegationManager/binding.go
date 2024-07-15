// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DelegationManager

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

// IDelegationManagerOperatorDetails is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerOperatorDetails struct {
	DeprecatedEarningsReceiver common.Address
	DelegationApprover         common.Address
	StakerOptOutWindowBlocks   uint32
}

// IDelegationManagerQueuedWithdrawalParams is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerQueuedWithdrawalParams struct {
	Strategies []common.Address
	Shares     []*big.Int
	Withdrawer common.Address
}

// IDelegationManagerWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerWithdrawal struct {
	Staker      common.Address
	DelegatedTo common.Address
	Withdrawer  common.Address
	Nonce       *big.Int
	StartBlock  uint32
	Strategies  []common.Address
	Shares      []*big.Int
}

// ISignatureUtilsSignatureWithExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithExpiry struct {
	Signature []byte
	Expiry    *big.Int
}

// DelegationManagerMetaData contains all meta data concerning the DelegationManager contract.
var DelegationManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_slasher\",\"type\":\"address\",\"internalType\":\"contractISlasher\"},{\"name\":\"_eigenPodManager\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DELEGATION_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_STAKER_OPT_OUT_WINDOW_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_WITHDRAWAL_DELAY_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKER_DELEGATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateCurrentStakerDelegationDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateDelegationApprovalDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateStakerDelegationDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stakerNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateWithdrawalRoot\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawal\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"middlewareTimesIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawals\",\"inputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManager.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"middlewareTimesIndexes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cumulativeWithdrawalsQueued\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateTo\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateToBySignature\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegatedTo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApprover\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApproverSaltIsSpent\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelegatableShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawalDelay\",\"inputs\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minWithdrawalDelayBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"_withdrawalDelayBlocks\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minWithdrawalDelayBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyOperatorDetails\",\"inputs\":[{\"name\":\"newOperatorDetails\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"__deprecated_earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorDetails\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"__deprecated_earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingWithdrawals\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueWithdrawals\",\"inputs\":[{\"name\":\"queuedWithdrawalParams\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManager.QueuedWithdrawalParams[]\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerAsOperator\",\"inputs\":[{\"name\":\"registeringOperatorDetails\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"__deprecated_earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinWithdrawalDelayBlocks\",\"inputs\":[{\"name\":\"newMinWithdrawalDelayBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyWithdrawalDelayBlocks\",\"inputs\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"withdrawalDelayBlocks\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerNonce\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerOptOutWindowBlocks\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyWithdrawalDelayBlocks\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinWithdrawalDelayBlocksSet\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDetailsModified\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOperatorDetails\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"__deprecated_earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMetadataURIUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorDetails\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"__deprecated_earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesDecreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesIncreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerForceUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyWithdrawalDelayBlocksSet\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalCompleted\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalQueued\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"withdrawal\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationManager.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200a4bd3803806200a4bd8339818101604052810190620000389190620002ca565b8282828273ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1681525050505050620000ea620000fb60201b60201c565b4660e081815250505050506200040a565b600060019054906101000a900460ff16156200014e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200014590620003ad565b60405180910390fd5b60ff801660008054906101000a900460ff1660ff161015620001c05760ff6000806101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff604051620001b79190620003ed565b60405180910390a15b565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001f482620001c7565b9050919050565b60006200020882620001e7565b9050919050565b6200021a81620001fb565b81146200022657600080fd5b50565b6000815190506200023a816200020f565b92915050565b60006200024d82620001e7565b9050919050565b6200025f8162000240565b81146200026b57600080fd5b50565b6000815190506200027f8162000254565b92915050565b60006200029282620001e7565b9050919050565b620002a48162000285565b8114620002b057600080fd5b50565b600081519050620002c48162000299565b92915050565b600080600060608486031215620002e657620002e5620001c2565b5b6000620002f68682870162000229565b935050602062000309868287016200026e565b92505060406200031c86828701620002b3565b9150509250925092565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b60006200039560278362000326565b9150620003a28262000337565b604082019050919050565b60006020820190508181036000830152620003c88162000386565b9050919050565b600060ff82169050919050565b620003e781620003cf565b82525050565b6000602082019050620004046000830184620003dc565b92915050565b60805160a05160c05160e05161a01e6200049f600039600061325c01526000818161127c015281816117cf01528181611ba7015281816125a9015281816135d601528181614db701526153410152600061234c0152600081816112270152818161177a01528181611a5b01528181612648015281816136b7015281816137ac01528181614f6001526153d5015261a01e6000f3fe608060405234801561001057600080fd5b50600436106103425760003560e01c8063635bbd10116101b8578063b7f06ebe11610104578063cf80873e116100a2578063f16172b01161007c578063f16172b014610a6e578063f2fde38b14610a8a578063f698da2514610aa6578063fabc1cbc14610ac457610342565b8063cf80873e146109f1578063da8be86414610a22578063eea9064b14610a5257610342565b8063c488375a116100de578063c488375a14610943578063c5e480db14610973578063c94b5111146109a3578063ca661c04146109d357610342565b8063b7f06ebe146108c5578063bb45fef2146108f5578063c448feb81461092557610342565b8063886f1195116101715780639104c3191161014b5780639104c3191461083d57806399be81c81461085b578063a178848414610877578063b1344271146108a757610342565b8063886f1195146107d15780638da5cb5b146107ef578063900413471461080d57610342565b8063635bbd10146106ff57806365da12641461071b5780636d70f7ae1461074b578063715018a61461077b578063778e55f3146107855780637f548071146107b557610342565b806328a573ae116102925780634665bcda11610230578063597b36da1161020a578063597b36da146106655780635ac86ab7146106955780635c975abb146106c557806360d7faed146106e357610342565b80634665bcda1461061f5780634fc40b611461063d578063595c6a671461065b57610342565b806339b70e381161026c57806339b70e38146105835780633cdeb5e0146105a15780633e28391d146105d1578063433773821461060157610342565b806328a573ae1461051b57806329c77d4f14610537578063334043961461056757610342565b8063132d4967116102ff57806316928365116102d957806316928365146104815780631bbce091146104b157806320606b70146104e157806322bf40e4146104ff57610342565b8063132d49671461042d578063136439dd146104495780631522bf021461046557610342565b80630449ca391461034757806304a4f979146103775780630b9f487a146103955780630dd8dd02146103c55780630f589e59146103f557806310d67a2f14610411575b600080fd5b610361600480360381019061035c9190615960565b610ae0565b60405161036e91906159c6565b60405180910390f35b61037f610b8a565b60405161038c91906159fa565b60405180910390f35b6103af60048036038101906103aa9190615acb565b610bae565b6040516103bc91906159fa565b60405180910390f35b6103df60048036038101906103da9190615b9c565b610c46565b6040516103ec9190615ca7565b60405180910390f35b61040f600480360381019061040a9190615d43565b61100d565b005b61042b60048036038101906104269190615de1565b61111b565b005b61044760048036038101906104429190615e4c565b611225565b005b610463600480360381019061045e9190615e9f565b61138e565b005b61047f600480360381019061047a9190615f22565b611509565b005b61049b60048036038101906104969190615fa3565b611523565b6040516104a891906159c6565b60405180910390f35b6104cb60048036038101906104c69190615fd0565b611585565b6040516104d891906159fa565b60405180910390f35b6104e96115e0565b6040516104f691906159fa565b60405180910390f35b61051960048036038101906105149190616023565b611604565b005b61053560048036038101906105309190615e4c565b611778565b005b610551600480360381019061054c9190615fa3565b6118e1565b60405161055e91906159c6565b60405180910390f35b610581600480360381019061057c91906161f4565b6118f9565b005b61058b611a59565b604051610598919061633c565b60405180910390f35b6105bb60048036038101906105b69190615fa3565b611a7d565b6040516105c89190616366565b60405180910390f35b6105eb60048036038101906105e69190615fa3565b611ae9565b6040516105f8919061639c565b60405180910390f35b610609611b81565b60405161061691906159fa565b60405180910390f35b610627611ba5565b60405161063491906163d8565b60405180910390f35b610645611bc9565b60405161065291906159c6565b60405180910390f35b610663611bd0565b005b61067f600480360381019061067a9190616737565b611d42565b60405161068c91906159fa565b60405180910390f35b6106af60048036038101906106aa91906167b9565b611d72565b6040516106bc919061639c565b60405180910390f35b6106cd611d8e565b6040516106da91906159c6565b60405180910390f35b6106fd60048036038101906106f89190616887565b611d98565b005b61071960048036038101906107149190615e9f565b611e4e565b005b61073560048036038101906107309190615fa3565b611e62565b6040516107429190616366565b60405180910390f35b61076560048036038101906107609190615fa3565b611e95565b604051610772919061639c565b60405180910390f35b610783611f2c565b005b61079f600480360381019061079a919061692b565b611f40565b6040516107ac91906159c6565b60405180910390f35b6107cf60048036038101906107ca9190616a8c565b611f65565b005b6107d9612100565b6040516107e69190616b60565b60405180910390f35b6107f7612126565b6040516108049190616366565b60405180910390f35b61082760048036038101906108229190616b7b565b612150565b6040516108349190616c95565b60405180910390f35b61084561227e565b6040516108529190616cd8565b60405180910390f35b61087560048036038101906108709190616cf3565b612296565b005b610891600480360381019061088c9190615fa3565b612332565b60405161089e91906159c6565b60405180910390f35b6108af61234a565b6040516108bc9190616d61565b60405180910390f35b6108df60048036038101906108da9190616d7c565b61236e565b6040516108ec919061639c565b60405180910390f35b61090f600480360381019061090a9190616da9565b61238e565b60405161091c919061639c565b60405180910390f35b61092d6123bd565b60405161093a91906159c6565b60405180910390f35b61095d60048036038101906109589190616de9565b6123c3565b60405161096a91906159c6565b60405180910390f35b61098d60048036038101906109889190615fa3565b6123db565b60405161099a9190616e76565b60405180910390f35b6109bd60048036038101906109b89190616e91565b612506565b6040516109ca91906159fa565b60405180910390f35b6109db61259b565b6040516109e891906159c6565b60405180910390f35b610a0b6004803603810190610a069190615fa3565b6125a2565b604051610a19929190616fb6565b60405180910390f35b610a3c6004803603810190610a379190615fa3565b612a39565b604051610a499190615ca7565b60405180910390f35b610a6c6004803603810190610a679190616fed565b6130dd565b005b610a886004803603810190610a83919061705c565b61317f565b005b610aa46004803603810190610a9f9190615fa3565b6131d4565b005b610aae613258565b604051610abb91906159fa565b60405180910390f35b610ade6004803603810190610ad99190615e9f565b61329a565b005b600080609d54905060005b84849050811015610b7f57600060a16000878785818110610b0f57610b0e617089565b5b9050602002016020810190610b249190616de9565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811115610b6d578092505b5080610b78906170e7565b9050610aeb565b508091505092915050565b7f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad81565b6000807f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad8588888787604051602001610bec96959493929190617130565b6040516020818303038152906040528051906020012090506000610c0e613258565b82604051602001610c20929190617209565b604051602081830303815290604052805190602001209050809250505095945050505050565b60606001610c5381611d72565b15610c93576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8a9061729d565b60405180910390fd5b60008484905067ffffffffffffffff811115610cb257610cb1616409565b5b604051908082528060200260200182016040528015610ce05781602001602082028036833780820191505090505b5090506000609a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060005b8686905081101561100057868682818110610d6857610d67617089565b5b9050602002810190610d7a91906172cc565b8060200190610d8991906172f4565b9050878783818110610d9e57610d9d617089565b5b9050602002810190610db091906172cc565b8060000190610dbf9190617357565b905014610e01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610df89061742c565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16878783818110610e2b57610e2a617089565b5b9050602002810190610e3d91906172cc565b6040016020810190610e4f9190615fa3565b73ffffffffffffffffffffffffffffffffffffffff1614610ea5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e9c906174be565b60405180910390fd5b610fce3383898985818110610ebd57610ebc617089565b5b9050602002810190610ecf91906172cc565b6040016020810190610ee19190615fa3565b8a8a86818110610ef457610ef3617089565b5b9050602002810190610f0691906172cc565b8060000190610f159190617357565b80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508b8b87818110610f6857610f67617089565b5b9050602002810190610f7a91906172cc565b8060200190610f8991906172f4565b80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505061343b565b838281518110610fe157610fe0617089565b5b6020026020010181815250508080610ff8906170e7565b915050610d4a565b5081935050505092915050565b61101633611ae9565b15611056576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161104d90617576565b60405180910390fd5b6110603384613a13565b61106861587a565b6110773333836000801b613bc5565b3373ffffffffffffffffffffffffffffffffffffffff167f8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e2856040516110bd919061761e565b60405180910390a23373ffffffffffffffffffffffffffffffffffffffff167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090848460405161110d929190617666565b60405180910390a250505050565b606560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611188573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111ac919061769f565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611219576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112109061773e565b60405180910390fd5b61122281613fdb565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806112ca57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b611309576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611300906177d0565b60405180910390fd5b61131283611ae9565b15611389576000609a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050611387818585856140ea565b505b505050565b606560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166346fbf68e336040518263ffffffff1660e01b81526004016113e99190616366565b602060405180830381865afa158015611406573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061142a9190617805565b611469576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611460906178a4565b60405180910390fd5b6066548160665416146114b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114a890617936565b60405180910390fd5b806066819055503373ffffffffffffffffffffffffffffffffffffffff167fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d826040516114fe91906159c6565b60405180910390a250565b6115116141d5565b61151d84848484614253565b50505050565b6000609960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160149054906101000a900463ffffffff1663ffffffff169050919050565b600080609b60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506115d685828686612506565b9150509392505050565b7f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a86681565b60008060019054906101000a900460ff161590508080156116355750600160008054906101000a900460ff1660ff16105b8061166257506116443061441d565b1580156116615750600160008054906101000a900460ff1660ff16145b5b6116a1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611698906179c8565b60405180910390fd5b60016000806101000a81548160ff021916908360ff16021790555080156116de576001600060016101000a81548160ff0219169083151502179055505b6116e88888614440565b6116f061456c565b6097819055506116ff896145fc565b611708866146c2565b61171485858585614253565b801561176d5760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516117649190617a23565b60405180910390a15b505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061181d57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b61185c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611853906177d0565b60405180910390fd5b61186583611ae9565b156118dc576000609a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506118da8185858561474d565b505b505050565b609b6020528060005260406000206000915090505481565b600261190481611d72565b15611944576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161193b9061729d565b60405180910390fd5b600260c954141561198a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161198190617a8a565b60405180910390fd5b600260c98190555060005b89899050811015611a4557611a348a8a838181106119b6576119b5617089565b5b90506020028101906119c89190617aaa565b8989848181106119db576119da617089565b5b90506020028101906119ed9190617ad2565b898986818110611a00576119ff617089565b5b90506020020135888887818110611a1a57611a19617089565b5b9050602002016020810190611a2f9190617b35565b614838565b80611a3e906170e7565b9050611995565b50600160c981905550505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000609960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60008073ffffffffffffffffffffffffffffffffffffffff16609a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b7f39111bc4a4d688e1f685123d7497d4615370152a8ee4a0593e647bd06ad8bb0b81565b7f000000000000000000000000000000000000000000000000000000000000000081565b6213c68081565b606560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166346fbf68e336040518263ffffffff1660e01b8152600401611c2b9190616366565b602060405180830381865afa158015611c48573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c6c9190617805565b611cab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ca2906178a4565b60405180910390fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6066819055503373ffffffffffffffffffffffffffffffffffffffff167fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff604051611d3891906159c6565b60405180910390a2565b600081604051602001611d559190617ce3565b604051602081830303815290604052805190602001209050919050565b6000808260ff166001901b905080816066541614915050919050565b6000606654905090565b6002611da381611d72565b15611de3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dda9061729d565b60405180910390fd5b600260c9541415611e29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e2090617a8a565b60405180910390fd5b600260c981905550611e3e8686868686614838565b600160c981905550505050505050565b611e566141d5565b611e5f816146c2565b50565b609a6020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008173ffffffffffffffffffffffffffffffffffffffff16609a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16149050919050565b611f346141d5565b611f3e60006145fc565b565b6098602052816000526040600020602052806000526040600020600091509150505481565b4283602001511015611fac576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fa390617d9d565b60405180910390fd5b611fb585611ae9565b15611ff5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fec90617e55565b60405180910390fd5b611ffe84611e95565b61203d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161203490617f0d565b60405180910390fd5b6000609b60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060006120938783888860200151612506565b905060018201609b60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506120eb8782876000015161515d565b6120f787878686613bc5565b50505050505050565b606560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606000825167ffffffffffffffff81111561216f5761216e616409565b5b60405190808252806020026020018201604052801561219d5781602001602082028036833780820191505090505b50905060005b835181101561227357609860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000858381518110612200576121ff617089565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482828151811061225657612255617089565b5b6020026020010181815250508061226c906170e7565b90506121a3565b508091505092915050565b73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b61229f33611e95565b6122de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122d590617fc5565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b67080908383604051612326929190617666565b60405180910390a25050565b609f6020528060005260406000206000915090505481565b7f000000000000000000000000000000000000000000000000000000000000000081565b609e6020528060005260406000206000915054906101000a900460ff1681565b609c6020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b609d5481565b60a16020528060005260406000206000915090505481565b6123e3615894565b609960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900463ffffffff1663ffffffff1663ffffffff16815250509050919050565b6000807f39111bc4a4d688e1f685123d7497d4615370152a8ee4a0593e647bd06ad8bb0b86858786604051602001612542959493929190617fe5565b6040516020818303038152906040528051906020012090506000612564613258565b82604051602001612576929190617209565b6040516020818303038152906040528051906020012090508092505050949350505050565b62034bc081565b60608060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166360f4062b856040518263ffffffff1660e01b81526004016126009190616366565b602060405180830381865afa15801561261d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612641919061806e565b90506000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166394f649dd876040518263ffffffff1660e01b815260040161269f9190616366565b600060405180830381865afa1580156126bc573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906126e591906181f3565b91509150600083136126ff57818194509450505050612a34565b60608060008451141561282d57600167ffffffffffffffff81111561272757612726616409565b5b6040519080825280602002602001820160405280156127555781602001602082028036833780820191505090505b509150600167ffffffffffffffff81111561277357612772616409565b5b6040519080825280602002602001820160405280156127a15781602001602082028036833780820191505090505b50905073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0826000815181106127cd576127cc617089565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050848160008151811061281c5761281b617089565b5b602002602001018181525050612a28565b6001845161283b919061826b565b67ffffffffffffffff81111561285457612853616409565b5b6040519080825280602002602001820160405280156128825781602001602082028036833780820191505090505b509150815167ffffffffffffffff8111156128a05761289f616409565b5b6040519080825280602002602001820160405280156128ce5781602001602082028036833780820191505090505b50905060005b845181101561298a578481815181106128f0576128ef617089565b5b602002602001015183828151811061290b5761290a617089565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083818151811061295857612957617089565b5b602002602001015182828151811061297357612972617089565b5b6020026020010181815250508060010190506128d4565b5073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac082600184516129af91906182c1565b815181106129c0576129bf617089565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050848160018451612a0a91906182c1565b81518110612a1b57612a1a617089565b5b6020026020010181815250505b81819650965050505050505b915091565b60606001612a4681611d72565b15612a86576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a7d9061729d565b60405180910390fd5b612a8f83611ae9565b612ace576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ac59061838d565b60405180910390fd5b612ad783611e95565b15612b17576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b0e9061841f565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415612b87576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b7e906184b1565b60405180910390fd5b6000609a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480612c5057508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80612ce85750609960008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b612d27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d1e90618543565b60405180910390fd5b600080612d33866125a2565b915091508573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612dc5578273ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167ff0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a60405160405180910390a35b8273ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167ffee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af4467660405160405180910390a36000609a60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600082511415612ef957600067ffffffffffffffff811115612ec357612ec2616409565b5b604051908082528060200260200182016040528015612ef15781602001602082028036833780820191505090505b5094506130d4565b815167ffffffffffffffff811115612f1457612f13616409565b5b604051908082528060200260200182016040528015612f425781602001602082028036833780820191505090505b50945060005b82518110156130d2576000600167ffffffffffffffff811115612f6e57612f6d616409565b5b604051908082528060200260200182016040528015612f9c5781602001602082028036833780820191505090505b5090506000600167ffffffffffffffff811115612fbc57612fbb616409565b5b604051908082528060200260200182016040528015612fea5781602001602082028036833780820191505090505b50905084838151811061300057612fff617089565b5b60200260200101518260008151811061301c5761301b617089565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505083838151811061306957613068617089565b5b60200260200101518160008151811061308557613084617089565b5b60200260200101818152505061309e89878b858561343b565b8884815181106130b1576130b0617089565b5b602002602001018181525050505080806130ca906170e7565b915050612f48565b505b50505050919050565b6130e633611ae9565b15613126576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161311d906185fb565b60405180910390fd5b61312f83611e95565b61316e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613165906186b3565b60405180910390fd5b61317a33848484613bc5565b505050565b61318833611e95565b6131c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131be9061876b565b60405180910390fd5b6131d13382613a13565b50565b6131dc6141d5565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561324c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613243906187fd565b60405180910390fd5b613255816145fc565b50565b60007f000000000000000000000000000000000000000000000000000000000000000046141561328c576097549050613297565b61329461456c565b90505b90565b606560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613307573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061332b919061769f565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614613398576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161338f9061773e565b60405180910390fd5b6066541981196066541916146133e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133da9061888f565b60405180910390fd5b806066819055503373ffffffffffffffffffffffffffffffffffffffff167f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c8260405161343091906159c6565b60405180910390a250565b60008073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614156134ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134a390618947565b60405180910390fd5b6000835114156134f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134e8906189ff565b60405180910390fd5b60005b835181101561387957600073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161461357257613571868886848151811061354957613548617089565b5b602002602001015186858151811061356457613563617089565b5b60200260200101516140ea565b5b73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac073ffffffffffffffffffffffffffffffffffffffff168482815181106135b0576135af617089565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff161415613680577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663beffbb898885848151811061362457613623617089565b5b60200260200101516040518363ffffffff1660e01b8152600401613649929190618a1f565b600060405180830381600087803b15801561366357600080fd5b505af1158015613677573d6000803e3d6000fd5b5050505061386e565b8473ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16148061376b57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639b4da03d85838151811061370457613703617089565b5b60200260200101516040518263ffffffff1660e01b81526004016137289190616cd8565b602060405180830381865afa158015613745573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137699190617805565b155b6137aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137a190618b2c565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16638c80d4e5888684815181106137fa576137f9617089565b5b602002602001015186858151811061381557613814617089565b5b60200260200101516040518463ffffffff1660e01b815260040161383b93929190618b4c565b600060405180830381600087803b15801561385557600080fd5b505af1158015613869573d6000803e3d6000fd5b505050505b8060010190506134f4565b506000609f60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050609f60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081548092919061390e906170e7565b919050555060006040518060e001604052808973ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018381526020014363ffffffff168152602001868152602001858152509050600061399d82611d42565b90506001609e600083815260200190815260200160002060006101000a81548160ff0219169083151502179055507f9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f981836040516139fc929190618b83565b60405180910390a180935050505095945050505050565b6213c680816040016020810190613a2a9190618bb3565b63ffffffff161115613a71576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613a6890618c9e565b60405180910390fd5b609960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160149054906101000a900463ffffffff1663ffffffff16816040016020810190613add9190618bb3565b63ffffffff161015613b24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613b1b90618d56565b60405180910390fd5b80609960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208181613b709190618f21565b9050503373ffffffffffffffffffffffffffffffffffffffff167ffebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac82604051613bb9919061761e565b60405180910390a25050565b6000613bd081611d72565b15613c10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613c079061729d565b60405180910390fd5b6000609960008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158015613ce057508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b8015613d1857508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15613e91574284602001511015613d64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613d5b90618fa1565b60405180910390fd5b609c60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060009054906101000a900460ff1615613e02576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613df990619033565b60405180910390fd5b6001609c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060006101000a81548160ff0219169083151502179055506000613e7e878784878960200151610bae565b9050613e8f8282876000015161515d565b505b84609a60008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d874330460405160405180910390a3600080613f75886125a2565b9150915060005b8251811015613fd057613fc5888a858481518110613f9d57613f9c617089565b5b6020026020010151858581518110613fb857613fb7617089565b5b602002602001015161474d565b806001019050613f7c565b505050505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561404b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614042906190eb565b60405180910390fd5b7f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6606560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168260405161409e92919061910b565b60405180910390a180606560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b80609860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461417691906182c1565b925050819055508373ffffffffffffffffffffffffffffffffffffffff167f6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd8484846040516141c793929190618b4c565b60405180910390a250505050565b6141dd6152ef565b73ffffffffffffffffffffffffffffffffffffffff166141fb612126565b73ffffffffffffffffffffffffffffffffffffffff1614614251576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161424890619180565b60405180910390fd5b565b81819050848490501461429b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161429290619238565b60405180910390fd5b600084849050905060005b818110156144155760008686838181106142c3576142c2617089565b5b90506020020160208101906142d89190616de9565b9050600060a160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600086868581811061433357614332617089565b5b90506020020135905062034bc0811115614382576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161437990619316565b60405180910390fd5b8060a160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f0e7efa738e8b0ce6376a0c1af471655540d2e9a81647d7b09ed823018426576d8383836040516143f993929190619336565b60405180910390a15050508061440e906170e7565b90506142a6565b505050505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600073ffffffffffffffffffffffffffffffffffffffff16606560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161480156144cb5750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b61450a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161450190619405565b60405180910390fd5b806066819055503373ffffffffffffffffffffffffffffffffffffffff167fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d8260405161455791906159c6565b60405180910390a261456882613fdb565b5050565b60007f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a8666040518060400160405280600a81526020017f456967656e4c61796572000000000000000000000000000000000000000000008152508051906020012046306040516020016145e19493929190619425565b60405160208183030381529060405280519060200120905090565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b62034bc0811115614708576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016146ff90619528565b60405180910390fd5b7fafa003cd76f87ff9d62b35beea889920f33c0c42b8d45b74954d61d50f4b6b69609d548260405161473b929190619548565b60405180910390a180609d8190555050565b80609860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546147d9919061826b565b925050819055508373ffffffffffffffffffffffffffffffffffffffff167f1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c84848460405161482a93929190618b4c565b60405180910390a250505050565b600061484c8661484790619571565b611d42565b9050609e600082815260200190815260200160002060009054906101000a900460ff166148ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016148a59061961c565b60405180910390fd5b43609d548760800160208101906148c59190618bb3565b63ffffffff166148d5919061826b565b1115614916576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161490d906196d4565b60405180910390fd5b8560400160208101906149299190615fa3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614614996576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161498d9061978c565b60405180910390fd5b81156149f357858060a001906149ac9190617357565b905085859050146149f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016149e990619844565b60405180910390fd5b5b609e600082815260200190815260200160002060006101000a81549060ff02191690558115614bc55760005b868060a00190614a2f9190617357565b9050811015614bbf574360a16000898060a00190614a4d9190617357565b85818110614a5e57614a5d617089565b5b9050602002016020810190614a739190616de9565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054886080016020810190614ac19190618bb3565b63ffffffff16614ad1919061826b565b1115614b12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614b0990619922565b60405180910390fd5b614bb4876000016020810190614b289190615fa3565b33898060a00190614b399190617357565b85818110614b4a57614b49617089565b5b9050602002016020810190614b5f9190616de9565b8a8060c00190614b6f91906172f4565b86818110614b8057614b7f617089565b5b905060200201358a8a87818110614b9a57614b99617089565b5b9050602002016020810190614baf9190619980565b6152f7565b806001019050614a1f565b5061511e565b6000609a60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060005b878060a00190614c3c9190617357565b905081101561511b574360a160008a8060a00190614c5a9190617357565b85818110614c6b57614c6a617089565b5b9050602002016020810190614c809190616de9565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054896080016020810190614cce9190618bb3565b63ffffffff16614cde919061826b565b1115614d1f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614d1690619922565b60405180910390fd5b73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac073ffffffffffffffffffffffffffffffffffffffff16888060a00190614d5a9190617357565b83818110614d6b57614d6a617089565b5b9050602002016020810190614d809190616de9565b73ffffffffffffffffffffffffffffffffffffffff161415614f5e576000886000016020810190614db19190615fa3565b905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16630e81073c838c8060c00190614e0291906172f4565b87818110614e1357614e12617089565b5b905060200201356040518363ffffffff1660e01b8152600401614e37929190618a1f565b6020604051808303816000875af1158015614e56573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614e7a91906199ad565b90506000609a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614614f5657614f5581848d8060a00190614f299190617357565b88818110614f3a57614f39617089565b5b9050602002016020810190614f4f9190616de9565b8561474d565b5b505050615110565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c4623ea133898985818110614fae57614fad617089565b5b9050602002016020810190614fc39190619980565b8b8060a00190614fd39190617357565b86818110614fe457614fe3617089565b5b9050602002016020810190614ff99190616de9565b8c8060c0019061500991906172f4565b8781811061501a57615019617089565b5b905060200201356040518563ffffffff1660e01b815260040161504094939291906199fb565b600060405180830381600087803b15801561505a57600080fd5b505af115801561506e573d6000803e3d6000fd5b50505050600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161461510f5761510e82338a8060a001906150bb9190617357565b858181106150cc576150cb617089565b5b90506020020160208101906150e19190616de9565b8b8060c001906150f191906172f4565b8681811061510257615101617089565b5b9050602002013561474d565b5b5b806001019050614c2c565b50505b7fc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d8160405161514d91906159fa565b60405180910390a1505050505050565b6151668361546c565b1561527257631626ba7e60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168373ffffffffffffffffffffffffffffffffffffffff16631626ba7e84846040518363ffffffff1660e01b81526004016151cd929190619ac8565b602060405180830381865afa1580156151ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061520e9190619b50565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461526d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161526490619c15565b60405180910390fd5b6152ea565b8273ffffffffffffffffffffffffffffffffffffffff16615293838361548f565b73ffffffffffffffffffffffffffffffffffffffff16146152e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016152e090619ccd565b60405180910390fd5b5b505050565b600033905090565b73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156153d3577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663387b13008686856040518463ffffffff1660e01b815260040161539c93929190619ced565b600060405180830381600087803b1580156153b657600080fd5b505af11580156153ca573d6000803e3d6000fd5b50505050615465565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c608c7f3858585856040518563ffffffff1660e01b81526004016154329493929190619d24565b600060405180830381600087803b15801561544c57600080fd5b505af1158015615460573d6000803e3d6000fd5b505050505b5050505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600080600061549e85856154b6565b915091506154ab81615539565b819250505092915050565b6000806041835114156154f85760008060006020860151925060408601519150606086015160001a90506154ec8782858561570e565b94509450505050615532565b60408351141561552957600080602085015191506040850151905061551e86838361581b565b935093505050615532565b60006002915091505b9250929050565b6000600481111561554d5761554c619d69565b5b8160048111156155605761555f619d69565b5b141561556b5761570b565b6001600481111561557f5761557e619d69565b5b81600481111561559257615591619d69565b5b14156155d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016155ca90619de4565b60405180910390fd5b600260048111156155e7576155e6619d69565b5b8160048111156155fa576155f9619d69565b5b141561563b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161563290619e50565b60405180910390fd5b6003600481111561564f5761564e619d69565b5b81600481111561566257615661619d69565b5b14156156a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161569a90619ee2565b60405180910390fd5b6004808111156156b6576156b5619d69565b5b8160048111156156c9576156c8619d69565b5b141561570a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161570190619f74565b60405180910390fd5b5b50565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08360001c1115615749576000600391509150615812565b601b8560ff16141580156157615750601c8560ff1614155b15615773576000600491509150615812565b6000600187878787604051600081526020016040526040516157989493929190619fa3565b6020604051602081039080840390855afa1580156157ba573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561580957600060019250925050615812565b80600092509250505b94509492505050565b60008060007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60001b841690506000601b60ff8660001c901c61585e919061826b565b905061586c8782888561570e565b935093505050935093915050565b604051806040016040528060608152602001600081525090565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600063ffffffff1681525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f8401126159205761591f6158fb565b5b8235905067ffffffffffffffff81111561593d5761593c615900565b5b60208301915083602082028301111561595957615958615905565b5b9250929050565b60008060208385031215615977576159766158f1565b5b600083013567ffffffffffffffff811115615995576159946158f6565b5b6159a18582860161590a565b92509250509250929050565b6000819050919050565b6159c0816159ad565b82525050565b60006020820190506159db60008301846159b7565b92915050565b6000819050919050565b6159f4816159e1565b82525050565b6000602082019050615a0f60008301846159eb565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000615a4082615a15565b9050919050565b615a5081615a35565b8114615a5b57600080fd5b50565b600081359050615a6d81615a47565b92915050565b615a7c816159e1565b8114615a8757600080fd5b50565b600081359050615a9981615a73565b92915050565b615aa8816159ad565b8114615ab357600080fd5b50565b600081359050615ac581615a9f565b92915050565b600080600080600060a08688031215615ae757615ae66158f1565b5b6000615af588828901615a5e565b9550506020615b0688828901615a5e565b9450506040615b1788828901615a5e565b9350506060615b2888828901615a8a565b9250506080615b3988828901615ab6565b9150509295509295909350565b60008083601f840112615b5c57615b5b6158fb565b5b8235905067ffffffffffffffff811115615b7957615b78615900565b5b602083019150836020820283011115615b9557615b94615905565b5b9250929050565b60008060208385031215615bb357615bb26158f1565b5b600083013567ffffffffffffffff811115615bd157615bd06158f6565b5b615bdd85828601615b46565b92509250509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b615c1e816159e1565b82525050565b6000615c308383615c15565b60208301905092915050565b6000602082019050919050565b6000615c5482615be9565b615c5e8185615bf4565b9350615c6983615c05565b8060005b83811015615c9a578151615c818882615c24565b9750615c8c83615c3c565b925050600181019050615c6d565b5085935050505092915050565b60006020820190508181036000830152615cc18184615c49565b905092915050565b600080fd5b600060608284031215615ce457615ce3615cc9565b5b81905092915050565b60008083601f840112615d0357615d026158fb565b5b8235905067ffffffffffffffff811115615d2057615d1f615900565b5b602083019150836001820283011115615d3c57615d3b615905565b5b9250929050565b600080600060808486031215615d5c57615d5b6158f1565b5b6000615d6a86828701615cce565b935050606084013567ffffffffffffffff811115615d8b57615d8a6158f6565b5b615d9786828701615ced565b92509250509250925092565b6000615dae82615a35565b9050919050565b615dbe81615da3565b8114615dc957600080fd5b50565b600081359050615ddb81615db5565b92915050565b600060208284031215615df757615df66158f1565b5b6000615e0584828501615dcc565b91505092915050565b6000615e1982615a35565b9050919050565b615e2981615e0e565b8114615e3457600080fd5b50565b600081359050615e4681615e20565b92915050565b600080600060608486031215615e6557615e646158f1565b5b6000615e7386828701615a5e565b9350506020615e8486828701615e37565b9250506040615e9586828701615ab6565b9150509250925092565b600060208284031215615eb557615eb46158f1565b5b6000615ec384828501615ab6565b91505092915050565b60008083601f840112615ee257615ee16158fb565b5b8235905067ffffffffffffffff811115615eff57615efe615900565b5b602083019150836020820283011115615f1b57615f1a615905565b5b9250929050565b60008060008060408587031215615f3c57615f3b6158f1565b5b600085013567ffffffffffffffff811115615f5a57615f596158f6565b5b615f668782880161590a565b9450945050602085013567ffffffffffffffff811115615f8957615f886158f6565b5b615f9587828801615ecc565b925092505092959194509250565b600060208284031215615fb957615fb86158f1565b5b6000615fc784828501615a5e565b91505092915050565b600080600060608486031215615fe957615fe86158f1565b5b6000615ff786828701615a5e565b935050602061600886828701615a5e565b925050604061601986828701615ab6565b9150509250925092565b60008060008060008060008060c0898b031215616043576160426158f1565b5b60006160518b828c01615a5e565b98505060206160628b828c01615dcc565b97505060406160738b828c01615ab6565b96505060606160848b828c01615ab6565b955050608089013567ffffffffffffffff8111156160a5576160a46158f6565b5b6160b18b828c0161590a565b945094505060a089013567ffffffffffffffff8111156160d4576160d36158f6565b5b6160e08b828c01615ecc565b92509250509295985092959890939650565b60008083601f840112616108576161076158fb565b5b8235905067ffffffffffffffff81111561612557616124615900565b5b60208301915083602082028301111561614157616140615905565b5b9250929050565b60008083601f84011261615e5761615d6158fb565b5b8235905067ffffffffffffffff81111561617b5761617a615900565b5b60208301915083602082028301111561619757616196615905565b5b9250929050565b60008083601f8401126161b4576161b36158fb565b5b8235905067ffffffffffffffff8111156161d1576161d0615900565b5b6020830191508360208202830111156161ed576161ec615905565b5b9250929050565b6000806000806000806000806080898b031215616214576162136158f1565b5b600089013567ffffffffffffffff811115616232576162316158f6565b5b61623e8b828c016160f2565b9850985050602089013567ffffffffffffffff811115616261576162606158f6565b5b61626d8b828c01616148565b9650965050604089013567ffffffffffffffff8111156162905761628f6158f6565b5b61629c8b828c01615ecc565b9450945050606089013567ffffffffffffffff8111156162bf576162be6158f6565b5b6162cb8b828c0161619e565b92509250509295985092959890939650565b6000819050919050565b60006163026162fd6162f884615a15565b6162dd565b615a15565b9050919050565b6000616314826162e7565b9050919050565b600061632682616309565b9050919050565b6163368161631b565b82525050565b6000602082019050616351600083018461632d565b92915050565b61636081615a35565b82525050565b600060208201905061637b6000830184616357565b92915050565b60008115159050919050565b61639681616381565b82525050565b60006020820190506163b1600083018461638d565b92915050565b60006163c282616309565b9050919050565b6163d2816163b7565b82525050565b60006020820190506163ed60008301846163c9565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b616441826163f8565b810181811067ffffffffffffffff821117156164605761645f616409565b5b80604052505050565b60006164736158e7565b905061647f8282616438565b919050565b600080fd5b600063ffffffff82169050919050565b6164a281616489565b81146164ad57600080fd5b50565b6000813590506164bf81616499565b92915050565b600067ffffffffffffffff8211156164e0576164df616409565b5b602082029050602081019050919050565b60006165046164ff846164c5565b616469565b9050808382526020820190506020840283018581111561652757616526615905565b5b835b81811015616550578061653c8882615e37565b845260208401935050602081019050616529565b5050509392505050565b600082601f83011261656f5761656e6158fb565b5b813561657f8482602086016164f1565b91505092915050565b600067ffffffffffffffff8211156165a3576165a2616409565b5b602082029050602081019050919050565b60006165c76165c284616588565b616469565b905080838252602082019050602084028301858111156165ea576165e9615905565b5b835b8181101561661357806165ff8882615ab6565b8452602084019350506020810190506165ec565b5050509392505050565b600082601f830112616632576166316158fb565b5b81356166428482602086016165b4565b91505092915050565b600060e08284031215616661576166606163f3565b5b61666b60e0616469565b9050600061667b84828501615a5e565b600083015250602061668f84828501615a5e565b60208301525060406166a384828501615a5e565b60408301525060606166b784828501615ab6565b60608301525060806166cb848285016164b0565b60808301525060a082013567ffffffffffffffff8111156166ef576166ee616484565b5b6166fb8482850161655a565b60a08301525060c082013567ffffffffffffffff81111561671f5761671e616484565b5b61672b8482850161661d565b60c08301525092915050565b60006020828403121561674d5761674c6158f1565b5b600082013567ffffffffffffffff81111561676b5761676a6158f6565b5b6167778482850161664b565b91505092915050565b600060ff82169050919050565b61679681616780565b81146167a157600080fd5b50565b6000813590506167b38161678d565b92915050565b6000602082840312156167cf576167ce6158f1565b5b60006167dd848285016167a4565b91505092915050565b600060e082840312156167fc576167fb615cc9565b5b81905092915050565b60008083601f84011261681b5761681a6158fb565b5b8235905067ffffffffffffffff81111561683857616837615900565b5b60208301915083602082028301111561685457616853615905565b5b9250929050565b61686481616381565b811461686f57600080fd5b50565b6000813590506168818161685b565b92915050565b6000806000806000608086880312156168a3576168a26158f1565b5b600086013567ffffffffffffffff8111156168c1576168c06158f6565b5b6168cd888289016167e6565b955050602086013567ffffffffffffffff8111156168ee576168ed6158f6565b5b6168fa88828901616805565b9450945050604061690d88828901615ab6565b925050606061691e88828901616872565b9150509295509295909350565b60008060408385031215616942576169416158f1565b5b600061695085828601615a5e565b925050602061696185828601615e37565b9150509250929050565b600080fd5b600067ffffffffffffffff82111561698b5761698a616409565b5b616994826163f8565b9050602081019050919050565b82818337600083830152505050565b60006169c36169be84616970565b616469565b9050828152602081018484840111156169df576169de61696b565b5b6169ea8482856169a1565b509392505050565b600082601f830112616a0757616a066158fb565b5b8135616a178482602086016169b0565b91505092915050565b600060408284031215616a3657616a356163f3565b5b616a406040616469565b9050600082013567ffffffffffffffff811115616a6057616a5f616484565b5b616a6c848285016169f2565b6000830152506020616a8084828501615ab6565b60208301525092915050565b600080600080600060a08688031215616aa857616aa76158f1565b5b6000616ab688828901615a5e565b9550506020616ac788828901615a5e565b945050604086013567ffffffffffffffff811115616ae857616ae76158f6565b5b616af488828901616a20565b935050606086013567ffffffffffffffff811115616b1557616b146158f6565b5b616b2188828901616a20565b9250506080616b3288828901615a8a565b9150509295509295909350565b6000616b4a82616309565b9050919050565b616b5a81616b3f565b82525050565b6000602082019050616b756000830184616b51565b92915050565b60008060408385031215616b9257616b916158f1565b5b6000616ba085828601615a5e565b925050602083013567ffffffffffffffff811115616bc157616bc06158f6565b5b616bcd8582860161655a565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b616c0c816159ad565b82525050565b6000616c1e8383616c03565b60208301905092915050565b6000602082019050919050565b6000616c4282616bd7565b616c4c8185616be2565b9350616c5783616bf3565b8060005b83811015616c88578151616c6f8882616c12565b9750616c7a83616c2a565b925050600181019050616c5b565b5085935050505092915050565b60006020820190508181036000830152616caf8184616c37565b905092915050565b6000616cc282616309565b9050919050565b616cd281616cb7565b82525050565b6000602082019050616ced6000830184616cc9565b92915050565b60008060208385031215616d0a57616d096158f1565b5b600083013567ffffffffffffffff811115616d2857616d276158f6565b5b616d3485828601615ced565b92509250509250929050565b6000616d4b82616309565b9050919050565b616d5b81616d40565b82525050565b6000602082019050616d766000830184616d52565b92915050565b600060208284031215616d9257616d916158f1565b5b6000616da084828501615a8a565b91505092915050565b60008060408385031215616dc057616dbf6158f1565b5b6000616dce85828601615a5e565b9250506020616ddf85828601615a8a565b9150509250929050565b600060208284031215616dff57616dfe6158f1565b5b6000616e0d84828501615e37565b91505092915050565b616e1f81615a35565b82525050565b616e2e81616489565b82525050565b606082016000820151616e4a6000850182616e16565b506020820151616e5d6020850182616e16565b506040820151616e706040850182616e25565b50505050565b6000606082019050616e8b6000830184616e34565b92915050565b60008060008060808587031215616eab57616eaa6158f1565b5b6000616eb987828801615a5e565b9450506020616eca87828801615ab6565b9350506040616edb87828801615a5e565b9250506060616eec87828801615ab6565b91505092959194509250565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b616f2d81616cb7565b82525050565b6000616f3f8383616f24565b60208301905092915050565b6000602082019050919050565b6000616f6382616ef8565b616f6d8185616f03565b9350616f7883616f14565b8060005b83811015616fa9578151616f908882616f33565b9750616f9b83616f4b565b925050600181019050616f7c565b5085935050505092915050565b60006040820190508181036000830152616fd08185616f58565b90508181036020830152616fe48184616c37565b90509392505050565b600080600060608486031215617006576170056158f1565b5b600061701486828701615a5e565b935050602084013567ffffffffffffffff811115617035576170346158f6565b5b61704186828701616a20565b925050604061705286828701615a8a565b9150509250925092565b600060608284031215617072576170716158f1565b5b600061708084828501615cce565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006170f2826159ad565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415617125576171246170b8565b5b600182019050919050565b600060c08201905061714560008301896159eb565b6171526020830188616357565b61715f6040830187616357565b61716c6060830186616357565b61717960808301856159eb565b61718660a08301846159b7565b979650505050505050565b600081905092915050565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b60006171d2600283617191565b91506171dd8261719c565b600282019050919050565b6000819050919050565b6172036171fe826159e1565b6171e8565b82525050565b6000617214826171c5565b915061722082856171f2565b60208201915061723082846171f2565b6020820191508190509392505050565b600082825260208201905092915050565b7f5061757361626c653a20696e6465782069732070617573656400000000000000600082015250565b6000617287601983617240565b915061729282617251565b602082019050919050565b600060208201905081810360008301526172b68161727a565b9050919050565b600080fd5b600080fd5b600080fd5b6000823560016060038336030381126172e8576172e76172bd565b5b80830191505092915050565b60008083356001602003843603038112617311576173106172bd565b5b80840192508235915067ffffffffffffffff821115617333576173326172c2565b5b60208301925060208202360383131561734f5761734e6172c7565b5b509250929050565b60008083356001602003843603038112617374576173736172bd565b5b80840192508235915067ffffffffffffffff821115617396576173956172c2565b5b6020830192506020820236038313156173b2576173b16172c7565b5b509250929050565b7f44656c65676174696f6e4d616e616765722e717565756557697468647261776160008201527f6c3a20696e707574206c656e677468206d69736d617463680000000000000000602082015250565b6000617416603883617240565b9150617421826173ba565b604082019050919050565b6000602082019050818103600083015261744581617409565b9050919050565b7f44656c65676174696f6e4d616e616765722e717565756557697468647261776160008201527f6c3a2077697468647261776572206d757374206265207374616b657200000000602082015250565b60006174a8603c83617240565b91506174b38261744c565b604082019050919050565b600060208201905081810360008301526174d78161749b565b9050919050565b7f44656c65676174696f6e4d616e616765722e726567697374657241734f70657260008201527f61746f723a2063616c6c657220697320616c7265616479206163746976656c7960208201527f2064656c65676174656400000000000000000000000000000000000000000000604082015250565b6000617560604a83617240565b915061756b826174de565b606082019050919050565b6000602082019050818103600083015261758f81617553565b9050919050565b60006175a56020840184615a5e565b905092915050565b60006175bc60208401846164b0565b905092915050565b606082016175d56000830183617596565b6175e26000850182616e16565b506175f06020830183617596565b6175fd6020850182616e16565b5061760b60408301836175ad565b6176186040850182616e25565b50505050565b600060608201905061763360008301846175c4565b92915050565b60006176458385617240565b93506176528385846169a1565b61765b836163f8565b840190509392505050565b60006020820190508181036000830152617681818486617639565b90509392505050565b60008151905061769981615a47565b92915050565b6000602082840312156176b5576176b46158f1565b5b60006176c38482850161768a565b91505092915050565b7f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160008201527f7320756e70617573657200000000000000000000000000000000000000000000602082015250565b6000617728602a83617240565b9150617733826176cc565b604082019050919050565b600060208201905081810360008301526177578161771b565b9050919050565b7f44656c65676174696f6e4d616e616765723a206f6e6c7953747261746567794d60008201527f616e616765724f72456967656e506f644d616e61676572000000000000000000602082015250565b60006177ba603783617240565b91506177c58261775e565b604082019050919050565b600060208201905081810360008301526177e9816177ad565b9050919050565b6000815190506177ff8161685b565b92915050565b60006020828403121561781b5761781a6158f1565b5b6000617829848285016177f0565b91505092915050565b7f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160008201527f7320706175736572000000000000000000000000000000000000000000000000602082015250565b600061788e602883617240565b915061789982617832565b604082019050919050565b600060208201905081810360008301526178bd81617881565b9050919050565b7f5061757361626c652e70617573653a20696e76616c696420617474656d70742060008201527f746f20756e70617573652066756e6374696f6e616c6974790000000000000000602082015250565b6000617920603883617240565b915061792b826178c4565b604082019050919050565b6000602082019050818103600083015261794f81617913565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b60006179b2602e83617240565b91506179bd82617956565b604082019050919050565b600060208201905081810360008301526179e1816179a5565b9050919050565b6000819050919050565b6000617a0d617a08617a03846179e8565b6162dd565b616780565b9050919050565b617a1d816179f2565b82525050565b6000602082019050617a386000830184617a14565b92915050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b6000617a74601f83617240565b9150617a7f82617a3e565b602082019050919050565b60006020820190508181036000830152617aa381617a67565b9050919050565b60008235600160e003833603038112617ac657617ac56172bd565b5b80830191505092915050565b60008083356001602003843603038112617aef57617aee6172bd565b5b80840192508235915067ffffffffffffffff821115617b1157617b106172c2565b5b602083019250602082023603831315617b2d57617b2c6172c7565b5b509250929050565b600060208284031215617b4b57617b4a6158f1565b5b6000617b5984828501616872565b91505092915050565b600082825260208201905092915050565b6000617b7e82616ef8565b617b888185617b62565b9350617b9383616f14565b8060005b83811015617bc4578151617bab8882616f33565b9750617bb683616f4b565b925050600181019050617b97565b5085935050505092915050565b600082825260208201905092915050565b6000617bed82616bd7565b617bf78185617bd1565b9350617c0283616bf3565b8060005b83811015617c33578151617c1a8882616c12565b9750617c2583616c2a565b925050600181019050617c06565b5085935050505092915050565b600060e083016000830151617c586000860182616e16565b506020830151617c6b6020860182616e16565b506040830151617c7e6040860182616e16565b506060830151617c916060860182616c03565b506080830151617ca46080860182616e25565b5060a083015184820360a0860152617cbc8282617b73565b91505060c083015184820360c0860152617cd68282617be2565b9150508091505092915050565b60006020820190508181036000830152617cfd8184617c40565b905092915050565b7f44656c65676174696f6e4d616e616765722e64656c6567617465546f4279536960008201527f676e61747572653a207374616b6572207369676e61747572652065787069726560208201527f6400000000000000000000000000000000000000000000000000000000000000604082015250565b6000617d87604183617240565b9150617d9282617d05565b606082019050919050565b60006020820190508181036000830152617db681617d7a565b9050919050565b7f44656c65676174696f6e4d616e616765722e64656c6567617465546f4279536960008201527f676e61747572653a207374616b657220697320616c726561647920616374697660208201527f656c792064656c65676174656400000000000000000000000000000000000000604082015250565b6000617e3f604d83617240565b9150617e4a82617dbd565b606082019050919050565b60006020820190508181036000830152617e6e81617e32565b9050919050565b7f44656c65676174696f6e4d616e616765722e64656c6567617465546f4279536960008201527f676e61747572653a206f70657261746f72206973206e6f74207265676973746560208201527f72656420696e20456967656e4c61796572000000000000000000000000000000604082015250565b6000617ef7605183617240565b9150617f0282617e75565b606082019050919050565b60006020820190508181036000830152617f2681617eea565b9050919050565b7f44656c65676174696f6e4d616e616765722e7570646174654f70657261746f7260008201527f4d657461646174615552493a2063616c6c6572206d75737420626520616e206f60208201527f70657261746f7200000000000000000000000000000000000000000000000000604082015250565b6000617faf604783617240565b9150617fba82617f2d565b606082019050919050565b60006020820190508181036000830152617fde81617fa2565b9050919050565b600060a082019050617ffa60008301886159eb565b6180076020830187616357565b6180146040830186616357565b61802160608301856159b7565b61802e60808301846159b7565b9695505050505050565b6000819050919050565b61804b81618038565b811461805657600080fd5b50565b60008151905061806881618042565b92915050565b600060208284031215618084576180836158f1565b5b600061809284828501618059565b91505092915050565b6000815190506180aa81615e20565b92915050565b60006180c36180be846164c5565b616469565b905080838252602082019050602084028301858111156180e6576180e5615905565b5b835b8181101561810f57806180fb888261809b565b8452602084019350506020810190506180e8565b5050509392505050565b600082601f83011261812e5761812d6158fb565b5b815161813e8482602086016180b0565b91505092915050565b60008151905061815681615a9f565b92915050565b600061816f61816a84616588565b616469565b9050808382526020820190506020840283018581111561819257618191615905565b5b835b818110156181bb57806181a78882618147565b845260208401935050602081019050618194565b5050509392505050565b600082601f8301126181da576181d96158fb565b5b81516181ea84826020860161815c565b91505092915050565b6000806040838503121561820a576182096158f1565b5b600083015167ffffffffffffffff811115618228576182276158f6565b5b61823485828601618119565b925050602083015167ffffffffffffffff811115618255576182546158f6565b5b618261858286016181c5565b9150509250929050565b6000618276826159ad565b9150618281836159ad565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156182b6576182b56170b8565b5b828201905092915050565b60006182cc826159ad565b91506182d7836159ad565b9250828210156182ea576182e96170b8565b5b828203905092915050565b7f44656c65676174696f6e4d616e616765722e756e64656c65676174653a20737460008201527f616b6572206d7573742062652064656c65676174656420746f20756e64656c6560208201527f6761746500000000000000000000000000000000000000000000000000000000604082015250565b6000618377604483617240565b9150618382826182f5565b606082019050919050565b600060208201905081810360008301526183a68161836a565b9050919050565b7f44656c65676174696f6e4d616e616765722e756e64656c65676174653a206f7060008201527f657261746f72732063616e6e6f7420626520756e64656c656761746564000000602082015250565b6000618409603d83617240565b9150618414826183ad565b604082019050919050565b60006020820190508181036000830152618438816183fc565b9050919050565b7f44656c65676174696f6e4d616e616765722e756e64656c65676174653a20636160008201527f6e6e6f7420756e64656c6567617465207a65726f206164647265737300000000602082015250565b600061849b603c83617240565b91506184a68261843f565b604082019050919050565b600060208201905081810360008301526184ca8161848e565b9050919050565b7f44656c65676174696f6e4d616e616765722e756e64656c65676174653a20636160008201527f6c6c65722063616e6e6f7420756e64656c6567617465207374616b6572000000602082015250565b600061852d603d83617240565b9150618538826184d1565b604082019050919050565b6000602082019050818103600083015261855c81618520565b9050919050565b7f44656c65676174696f6e4d616e616765722e64656c6567617465546f3a20737460008201527f616b657220697320616c7265616479206163746976656c792064656c6567617460208201527f6564000000000000000000000000000000000000000000000000000000000000604082015250565b60006185e5604283617240565b91506185f082618563565b606082019050919050565b60006020820190508181036000830152618614816185d8565b9050919050565b7f44656c65676174696f6e4d616e616765722e64656c6567617465546f3a206f7060008201527f657261746f72206973206e6f74207265676973746572656420696e204569676560208201527f6e4c617965720000000000000000000000000000000000000000000000000000604082015250565b600061869d604683617240565b91506186a88261861b565b606082019050919050565b600060208201905081810360008301526186cc81618690565b9050919050565b7f44656c65676174696f6e4d616e616765722e6d6f646966794f70657261746f7260008201527f44657461696c733a2063616c6c6572206d75737420626520616e206f7065726160208201527f746f720000000000000000000000000000000000000000000000000000000000604082015250565b6000618755604383617240565b9150618760826186d3565b606082019050919050565b6000602082019050818103600083015261878481618748565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006187e7602683617240565b91506187f28261878b565b604082019050919050565b60006020820190508181036000830152618816816187da565b9050919050565b7f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060008201527f7420746f2070617573652066756e6374696f6e616c6974790000000000000000602082015250565b6000618879603883617240565b91506188848261881d565b604082019050919050565b600060208201905081810360008301526188a88161886c565b9050919050565b7f44656c65676174696f6e4d616e616765722e5f72656d6f76655368617265734160008201527f6e6451756575655769746864726177616c3a207374616b65722063616e6e6f7460208201527f206265207a65726f206164647265737300000000000000000000000000000000604082015250565b6000618931605083617240565b915061893c826188af565b606082019050919050565b6000602082019050818103600083015261896081618924565b9050919050565b7f44656c65676174696f6e4d616e616765722e5f72656d6f76655368617265734160008201527f6e6451756575655769746864726177616c3a207374726174656769657320636160208201527f6e6e6f7420626520656d70747900000000000000000000000000000000000000604082015250565b60006189e9604d83617240565b91506189f482618967565b606082019050919050565b60006020820190508181036000830152618a18816189dc565b9050919050565b6000604082019050618a346000830185616357565b618a4160208301846159b7565b9392505050565b7f44656c65676174696f6e4d616e616765722e5f72656d6f76655368617265734160008201527f6e6451756575655769746864726177616c3a2077697468647261776572206d7560208201527f73742062652073616d652061646472657373206173207374616b65722069662060408201527f746869726450617274795472616e7366657273466f7262696464656e2061726560608201527f2073657400000000000000000000000000000000000000000000000000000000608082015250565b6000618b16608483617240565b9150618b2182618a48565b60a082019050919050565b60006020820190508181036000830152618b4581618b09565b9050919050565b6000606082019050618b616000830186616357565b618b6e6020830185616cc9565b618b7b60408301846159b7565b949350505050565b6000604082019050618b9860008301856159eb565b8181036020830152618baa8184617c40565b90509392505050565b600060208284031215618bc957618bc86158f1565b5b6000618bd7848285016164b0565b91505092915050565b7f44656c65676174696f6e4d616e616765722e5f7365744f70657261746f72446560008201527f7461696c733a207374616b65724f70744f757457696e646f77426c6f636b732060208201527f63616e6e6f74206265203e204d41585f5354414b45525f4f50545f4f55545f5760408201527f494e444f575f424c4f434b530000000000000000000000000000000000000000606082015250565b6000618c88606c83617240565b9150618c9382618be0565b608082019050919050565b60006020820190508181036000830152618cb781618c7b565b9050919050565b7f44656c65676174696f6e4d616e616765722e5f7365744f70657261746f72446560008201527f7461696c733a207374616b65724f70744f757457696e646f77426c6f636b732060208201527f63616e6e6f742062652064656372656173656400000000000000000000000000604082015250565b6000618d40605383617240565b9150618d4b82618cbe565b606082019050919050565b60006020820190508181036000830152618d6f81618d33565b9050919050565b60008135618d8381615a47565b80915050919050565b60008160001b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff618db984618d8c565b9350801983169250808416831791505092915050565b6000618dda82616309565b9050919050565b6000819050919050565b618df482618dcf565b618e07618e0082618de1565b8354618d99565b8255505050565b60008135618e1b81616499565b80915050919050565b60008160a01b9050919050565b600077ffffffff0000000000000000000000000000000000000000618e5584618e24565b9350801983169250808416831791505092915050565b6000618e86618e81618e7c84616489565b6162dd565b616489565b9050919050565b6000819050919050565b618ea082618e6b565b618eb3618eac82618e8d565b8354618e31565b8255505050565b600081016000830180618ecc81618d76565b9050618ed88184618deb565b505050600181016020830180618eed81618d76565b9050618ef98184618deb565b505050600181016040830180618f0e81618e0e565b9050618f1a8184618e97565b5050505050565b618f2b8282618eba565b5050565b7f44656c65676174696f6e4d616e616765722e5f64656c65676174653a2061707060008201527f726f766572207369676e61747572652065787069726564000000000000000000602082015250565b6000618f8b603783617240565b9150618f9682618f2f565b604082019050919050565b60006020820190508181036000830152618fba81618f7e565b9050919050565b7f44656c65676174696f6e4d616e616765722e5f64656c65676174653a2061707060008201527f726f76657253616c7420616c7265616479207370656e74000000000000000000602082015250565b600061901d603783617240565b915061902882618fc1565b604082019050919050565b6000602082019050818103600083015261904c81619010565b9050919050565b7f5061757361626c652e5f73657450617573657252656769737472793a206e657760008201527f50617573657252656769737472792063616e6e6f7420626520746865207a657260208201527f6f20616464726573730000000000000000000000000000000000000000000000604082015250565b60006190d5604983617240565b91506190e082619053565b606082019050919050565b60006020820190508181036000830152619104816190c8565b9050919050565b60006040820190506191206000830185616b51565b61912d6020830184616b51565b9392505050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b600061916a602083617240565b915061917582619134565b602082019050919050565b600060208201905081810360008301526191998161915d565b9050919050565b7f44656c65676174696f6e4d616e616765722e5f7365745374726174656779576960008201527f746864726177616c44656c6179426c6f636b733a20696e707574206c656e677460208201527f68206d69736d6174636800000000000000000000000000000000000000000000604082015250565b6000619222604a83617240565b915061922d826191a0565b606082019050919050565b6000602082019050818103600083015261925181619215565b9050919050565b7f44656c65676174696f6e4d616e616765722e5f7365745374726174656779576960008201527f746864726177616c44656c6179426c6f636b733a205f7769746864726177616c60208201527f44656c6179426c6f636b732063616e6e6f74206265203e204d41585f5749544860408201527f44524157414c5f44454c41595f424c4f434b5300000000000000000000000000606082015250565b6000619300607383617240565b915061930b82619258565b608082019050919050565b6000602082019050818103600083015261932f816192f3565b9050919050565b600060608201905061934b6000830186616cc9565b61935860208301856159b7565b61936560408301846159b7565b949350505050565b7f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960008201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c60208201527f6564206f6e636500000000000000000000000000000000000000000000000000604082015250565b60006193ef604783617240565b91506193fa8261936d565b606082019050919050565b6000602082019050818103600083015261941e816193e2565b9050919050565b600060808201905061943a60008301876159eb565b61944760208301866159eb565b61945460408301856159b7565b6194616060830184616357565b95945050505050565b7f44656c65676174696f6e4d616e616765722e5f7365744d696e5769746864726160008201527f77616c44656c6179426c6f636b733a205f6d696e5769746864726177616c446560208201527f6c6179426c6f636b732063616e6e6f74206265203e204d41585f57495448445260408201527f4157414c5f44454c41595f424c4f434b53000000000000000000000000000000606082015250565b6000619512607183617240565b915061951d8261946a565b608082019050919050565b6000602082019050818103600083015261954181619505565b9050919050565b600060408201905061955d60008301856159b7565b61956a60208301846159b7565b9392505050565b600061957d368361664b565b9050919050565b7f44656c65676174696f6e4d616e616765722e5f636f6d706c657465517565756560008201527f645769746864726177616c3a20616374696f6e206973206e6f7420696e20717560208201527f6575650000000000000000000000000000000000000000000000000000000000604082015250565b6000619606604383617240565b915061961182619584565b606082019050919050565b60006020820190508181036000830152619635816195f9565b9050919050565b7f44656c65676174696f6e4d616e616765722e5f636f6d706c657465517565756560008201527f645769746864726177616c3a206d696e5769746864726177616c44656c61794260208201527f6c6f636b7320706572696f6420686173206e6f74207965742070617373656400604082015250565b60006196be605f83617240565b91506196c98261963c565b606082019050919050565b600060208201905081810360008301526196ed816196b1565b9050919050565b7f44656c65676174696f6e4d616e616765722e5f636f6d706c657465517565756560008201527f645769746864726177616c3a206f6e6c7920776974686472617765722063616e60208201527f20636f6d706c65746520616374696f6e00000000000000000000000000000000604082015250565b6000619776605083617240565b9150619781826196f4565b606082019050919050565b600060208201905081810360008301526197a581619769565b9050919050565b7f44656c65676174696f6e4d616e616765722e5f636f6d706c657465517565756560008201527f645769746864726177616c3a20696e707574206c656e677468206d69736d617460208201527f6368000000000000000000000000000000000000000000000000000000000000604082015250565b600061982e604283617240565b9150619839826197ac565b606082019050919050565b6000602082019050818103600083015261985d81619821565b9050919050565b7f44656c65676174696f6e4d616e616765722e5f636f6d706c657465517565756560008201527f645769746864726177616c3a207769746864726177616c44656c6179426c6f6360208201527f6b7320706572696f6420686173206e6f74207965742070617373656420666f7260408201527f2074686973207374726174656779000000000000000000000000000000000000606082015250565b600061990c606e83617240565b915061991782619864565b608082019050919050565b6000602082019050818103600083015261993b816198ff565b9050919050565b600061994d82615a35565b9050919050565b61995d81619942565b811461996857600080fd5b50565b60008135905061997a81619954565b92915050565b600060208284031215619996576199956158f1565b5b60006199a48482850161996b565b91505092915050565b6000602082840312156199c3576199c26158f1565b5b60006199d184828501618147565b91505092915050565b60006199e582616309565b9050919050565b6199f5816199da565b82525050565b6000608082019050619a106000830187616357565b619a1d60208301866199ec565b619a2a6040830185616cc9565b619a3760608301846159b7565b95945050505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015619a7a578082015181840152602081019050619a5f565b83811115619a89576000848401525b50505050565b6000619a9a82619a40565b619aa48185619a4b565b9350619ab4818560208601619a5c565b619abd816163f8565b840191505092915050565b6000604082019050619add60008301856159eb565b8181036020830152619aef8184619a8f565b90509392505050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b619b2d81619af8565b8114619b3857600080fd5b50565b600081519050619b4a81619b24565b92915050565b600060208284031215619b6657619b656158f1565b5b6000619b7484828501619b3b565b91505092915050565b7f454950313237315369676e61747572655574696c732e636865636b5369676e6160008201527f747572655f454950313237313a2045524331323731207369676e61747572652060208201527f766572696669636174696f6e206661696c656400000000000000000000000000604082015250565b6000619bff605383617240565b9150619c0a82619b7d565b606082019050919050565b60006020820190508181036000830152619c2e81619bf2565b9050919050565b7f454950313237315369676e61747572655574696c732e636865636b5369676e6160008201527f747572655f454950313237313a207369676e6174757265206e6f742066726f6d60208201527f207369676e657200000000000000000000000000000000000000000000000000604082015250565b6000619cb7604783617240565b9150619cc282619c35565b606082019050919050565b60006020820190508181036000830152619ce681619caa565b9050919050565b6000606082019050619d026000830186616357565b619d0f6020830185616357565b619d1c60408301846159b7565b949350505050565b6000608082019050619d396000830187616357565b619d466020830186616cc9565b619d5360408301856159b7565b619d6060608301846199ec565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f45434453413a20696e76616c6964207369676e61747572650000000000000000600082015250565b6000619dce601883617240565b9150619dd982619d98565b602082019050919050565b60006020820190508181036000830152619dfd81619dc1565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800600082015250565b6000619e3a601f83617240565b9150619e4582619e04565b602082019050919050565b60006020820190508181036000830152619e6981619e2d565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b6000619ecc602283617240565b9150619ed782619e70565b604082019050919050565b60006020820190508181036000830152619efb81619ebf565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b6000619f5e602283617240565b9150619f6982619f02565b604082019050919050565b60006020820190508181036000830152619f8d81619f51565b9050919050565b619f9d81616780565b82525050565b6000608082019050619fb860008301876159eb565b619fc56020830186619f94565b619fd260408301856159eb565b619fdf60608301846159eb565b9594505050505056fea2646970667358221220d3e8c2997b42ae4c5ef03378383284b1e02237fc7ac866c0cca3cbf39599746f64736f6c634300080c0033",
}

// DelegationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DelegationManagerMetaData.ABI instead.
var DelegationManagerABI = DelegationManagerMetaData.ABI

// DelegationManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DelegationManagerMetaData.Bin instead.
var DelegationManagerBin = DelegationManagerMetaData.Bin

// DeployDelegationManager deploys a new Ethereum contract, binding an instance of DelegationManager to it.
func DeployDelegationManager(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _slasher common.Address, _eigenPodManager common.Address) (common.Address, *types.Transaction, *DelegationManager, error) {
	parsed, err := DelegationManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DelegationManagerBin), backend, _strategyManager, _slasher, _eigenPodManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DelegationManager{DelegationManagerCaller: DelegationManagerCaller{contract: contract}, DelegationManagerTransactor: DelegationManagerTransactor{contract: contract}, DelegationManagerFilterer: DelegationManagerFilterer{contract: contract}}, nil
}

// DelegationManager is an auto generated Go binding around an Ethereum contract.
type DelegationManager struct {
	DelegationManagerCaller     // Read-only binding to the contract
	DelegationManagerTransactor // Write-only binding to the contract
	DelegationManagerFilterer   // Log filterer for contract events
}

// DelegationManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegationManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegationManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegationManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegationManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegationManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegationManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegationManagerSession struct {
	Contract     *DelegationManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DelegationManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegationManagerCallerSession struct {
	Contract *DelegationManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DelegationManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegationManagerTransactorSession struct {
	Contract     *DelegationManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DelegationManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegationManagerRaw struct {
	Contract *DelegationManager // Generic contract binding to access the raw methods on
}

// DelegationManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegationManagerCallerRaw struct {
	Contract *DelegationManagerCaller // Generic read-only contract binding to access the raw methods on
}

// DelegationManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegationManagerTransactorRaw struct {
	Contract *DelegationManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegationManager creates a new instance of DelegationManager, bound to a specific deployed contract.
func NewDelegationManager(address common.Address, backend bind.ContractBackend) (*DelegationManager, error) {
	contract, err := bindDelegationManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelegationManager{DelegationManagerCaller: DelegationManagerCaller{contract: contract}, DelegationManagerTransactor: DelegationManagerTransactor{contract: contract}, DelegationManagerFilterer: DelegationManagerFilterer{contract: contract}}, nil
}

// NewDelegationManagerCaller creates a new read-only instance of DelegationManager, bound to a specific deployed contract.
func NewDelegationManagerCaller(address common.Address, caller bind.ContractCaller) (*DelegationManagerCaller, error) {
	contract, err := bindDelegationManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerCaller{contract: contract}, nil
}

// NewDelegationManagerTransactor creates a new write-only instance of DelegationManager, bound to a specific deployed contract.
func NewDelegationManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegationManagerTransactor, error) {
	contract, err := bindDelegationManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerTransactor{contract: contract}, nil
}

// NewDelegationManagerFilterer creates a new log filterer instance of DelegationManager, bound to a specific deployed contract.
func NewDelegationManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegationManagerFilterer, error) {
	contract, err := bindDelegationManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerFilterer{contract: contract}, nil
}

// bindDelegationManager binds a generic wrapper to an already deployed contract.
func bindDelegationManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DelegationManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegationManager *DelegationManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelegationManager.Contract.DelegationManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegationManager *DelegationManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegationManager.Contract.DelegationManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegationManager *DelegationManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegationManager.Contract.DelegationManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegationManager *DelegationManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelegationManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegationManager *DelegationManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegationManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegationManager *DelegationManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegationManager.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_DelegationManager *DelegationManagerCaller) DELEGATIONAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "DELEGATION_APPROVAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_DelegationManager *DelegationManagerSession) DELEGATIONAPPROVALTYPEHASH() ([32]byte, error) {
	return _DelegationManager.Contract.DELEGATIONAPPROVALTYPEHASH(&_DelegationManager.CallOpts)
}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_DelegationManager *DelegationManagerCallerSession) DELEGATIONAPPROVALTYPEHASH() ([32]byte, error) {
	return _DelegationManager.Contract.DELEGATIONAPPROVALTYPEHASH(&_DelegationManager.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_DelegationManager *DelegationManagerCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_DelegationManager *DelegationManagerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _DelegationManager.Contract.DOMAINTYPEHASH(&_DelegationManager.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_DelegationManager *DelegationManagerCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _DelegationManager.Contract.DOMAINTYPEHASH(&_DelegationManager.CallOpts)
}

// MAXSTAKEROPTOUTWINDOWBLOCKS is a free data retrieval call binding the contract method 0x4fc40b61.
//
// Solidity: function MAX_STAKER_OPT_OUT_WINDOW_BLOCKS() view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) MAXSTAKEROPTOUTWINDOWBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "MAX_STAKER_OPT_OUT_WINDOW_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTAKEROPTOUTWINDOWBLOCKS is a free data retrieval call binding the contract method 0x4fc40b61.
//
// Solidity: function MAX_STAKER_OPT_OUT_WINDOW_BLOCKS() view returns(uint256)
func (_DelegationManager *DelegationManagerSession) MAXSTAKEROPTOUTWINDOWBLOCKS() (*big.Int, error) {
	return _DelegationManager.Contract.MAXSTAKEROPTOUTWINDOWBLOCKS(&_DelegationManager.CallOpts)
}

// MAXSTAKEROPTOUTWINDOWBLOCKS is a free data retrieval call binding the contract method 0x4fc40b61.
//
// Solidity: function MAX_STAKER_OPT_OUT_WINDOW_BLOCKS() view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) MAXSTAKEROPTOUTWINDOWBLOCKS() (*big.Int, error) {
	return _DelegationManager.Contract.MAXSTAKEROPTOUTWINDOWBLOCKS(&_DelegationManager.CallOpts)
}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) MAXWITHDRAWALDELAYBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "MAX_WITHDRAWAL_DELAY_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_DelegationManager *DelegationManagerSession) MAXWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _DelegationManager.Contract.MAXWITHDRAWALDELAYBLOCKS(&_DelegationManager.CallOpts)
}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) MAXWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _DelegationManager.Contract.MAXWITHDRAWALDELAYBLOCKS(&_DelegationManager.CallOpts)
}

// STAKERDELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0x43377382.
//
// Solidity: function STAKER_DELEGATION_TYPEHASH() view returns(bytes32)
func (_DelegationManager *DelegationManagerCaller) STAKERDELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "STAKER_DELEGATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKERDELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0x43377382.
//
// Solidity: function STAKER_DELEGATION_TYPEHASH() view returns(bytes32)
func (_DelegationManager *DelegationManagerSession) STAKERDELEGATIONTYPEHASH() ([32]byte, error) {
	return _DelegationManager.Contract.STAKERDELEGATIONTYPEHASH(&_DelegationManager.CallOpts)
}

// STAKERDELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0x43377382.
//
// Solidity: function STAKER_DELEGATION_TYPEHASH() view returns(bytes32)
func (_DelegationManager *DelegationManagerCallerSession) STAKERDELEGATIONTYPEHASH() ([32]byte, error) {
	return _DelegationManager.Contract.STAKERDELEGATIONTYPEHASH(&_DelegationManager.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_DelegationManager *DelegationManagerCaller) BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "beaconChainETHStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_DelegationManager *DelegationManagerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _DelegationManager.Contract.BeaconChainETHStrategy(&_DelegationManager.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_DelegationManager *DelegationManagerCallerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _DelegationManager.Contract.BeaconChainETHStrategy(&_DelegationManager.CallOpts)
}

// CalculateCurrentStakerDelegationDigestHash is a free data retrieval call binding the contract method 0x1bbce091.
//
// Solidity: function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationManager *DelegationManagerCaller) CalculateCurrentStakerDelegationDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "calculateCurrentStakerDelegationDigestHash", staker, operator, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateCurrentStakerDelegationDigestHash is a free data retrieval call binding the contract method 0x1bbce091.
//
// Solidity: function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationManager *DelegationManagerSession) CalculateCurrentStakerDelegationDigestHash(staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _DelegationManager.Contract.CalculateCurrentStakerDelegationDigestHash(&_DelegationManager.CallOpts, staker, operator, expiry)
}

// CalculateCurrentStakerDelegationDigestHash is a free data retrieval call binding the contract method 0x1bbce091.
//
// Solidity: function calculateCurrentStakerDelegationDigestHash(address staker, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationManager *DelegationManagerCallerSession) CalculateCurrentStakerDelegationDigestHash(staker common.Address, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _DelegationManager.Contract.CalculateCurrentStakerDelegationDigestHash(&_DelegationManager.CallOpts, staker, operator, expiry)
}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address _delegationApprover, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_DelegationManager *DelegationManagerCaller) CalculateDelegationApprovalDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "calculateDelegationApprovalDigestHash", staker, operator, _delegationApprover, approverSalt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address _delegationApprover, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_DelegationManager *DelegationManagerSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _DelegationManager.Contract.CalculateDelegationApprovalDigestHash(&_DelegationManager.CallOpts, staker, operator, _delegationApprover, approverSalt, expiry)
}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address _delegationApprover, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_DelegationManager *DelegationManagerCallerSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, _delegationApprover common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _DelegationManager.Contract.CalculateDelegationApprovalDigestHash(&_DelegationManager.CallOpts, staker, operator, _delegationApprover, approverSalt, expiry)
}

// CalculateStakerDelegationDigestHash is a free data retrieval call binding the contract method 0xc94b5111.
//
// Solidity: function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationManager *DelegationManagerCaller) CalculateStakerDelegationDigestHash(opts *bind.CallOpts, staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "calculateStakerDelegationDigestHash", staker, _stakerNonce, operator, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateStakerDelegationDigestHash is a free data retrieval call binding the contract method 0xc94b5111.
//
// Solidity: function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationManager *DelegationManagerSession) CalculateStakerDelegationDigestHash(staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _DelegationManager.Contract.CalculateStakerDelegationDigestHash(&_DelegationManager.CallOpts, staker, _stakerNonce, operator, expiry)
}

// CalculateStakerDelegationDigestHash is a free data retrieval call binding the contract method 0xc94b5111.
//
// Solidity: function calculateStakerDelegationDigestHash(address staker, uint256 _stakerNonce, address operator, uint256 expiry) view returns(bytes32)
func (_DelegationManager *DelegationManagerCallerSession) CalculateStakerDelegationDigestHash(staker common.Address, _stakerNonce *big.Int, operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _DelegationManager.Contract.CalculateStakerDelegationDigestHash(&_DelegationManager.CallOpts, staker, _stakerNonce, operator, expiry)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_DelegationManager *DelegationManagerCaller) CalculateWithdrawalRoot(opts *bind.CallOpts, withdrawal IDelegationManagerWithdrawal) ([32]byte, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "calculateWithdrawalRoot", withdrawal)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_DelegationManager *DelegationManagerSession) CalculateWithdrawalRoot(withdrawal IDelegationManagerWithdrawal) ([32]byte, error) {
	return _DelegationManager.Contract.CalculateWithdrawalRoot(&_DelegationManager.CallOpts, withdrawal)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_DelegationManager *DelegationManagerCallerSession) CalculateWithdrawalRoot(withdrawal IDelegationManagerWithdrawal) ([32]byte, error) {
	return _DelegationManager.Contract.CalculateWithdrawalRoot(&_DelegationManager.CallOpts, withdrawal)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address ) view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) CumulativeWithdrawalsQueued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "cumulativeWithdrawalsQueued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address ) view returns(uint256)
func (_DelegationManager *DelegationManagerSession) CumulativeWithdrawalsQueued(arg0 common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.CumulativeWithdrawalsQueued(&_DelegationManager.CallOpts, arg0)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address ) view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) CumulativeWithdrawalsQueued(arg0 common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.CumulativeWithdrawalsQueued(&_DelegationManager.CallOpts, arg0)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address ) view returns(address)
func (_DelegationManager *DelegationManagerCaller) DelegatedTo(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "delegatedTo", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address ) view returns(address)
func (_DelegationManager *DelegationManagerSession) DelegatedTo(arg0 common.Address) (common.Address, error) {
	return _DelegationManager.Contract.DelegatedTo(&_DelegationManager.CallOpts, arg0)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address ) view returns(address)
func (_DelegationManager *DelegationManagerCallerSession) DelegatedTo(arg0 common.Address) (common.Address, error) {
	return _DelegationManager.Contract.DelegatedTo(&_DelegationManager.CallOpts, arg0)
}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_DelegationManager *DelegationManagerCaller) DelegationApprover(opts *bind.CallOpts, operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "delegationApprover", operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_DelegationManager *DelegationManagerSession) DelegationApprover(operator common.Address) (common.Address, error) {
	return _DelegationManager.Contract.DelegationApprover(&_DelegationManager.CallOpts, operator)
}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_DelegationManager *DelegationManagerCallerSession) DelegationApprover(operator common.Address) (common.Address, error) {
	return _DelegationManager.Contract.DelegationApprover(&_DelegationManager.CallOpts, operator)
}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address , bytes32 ) view returns(bool)
func (_DelegationManager *DelegationManagerCaller) DelegationApproverSaltIsSpent(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "delegationApproverSaltIsSpent", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address , bytes32 ) view returns(bool)
func (_DelegationManager *DelegationManagerSession) DelegationApproverSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _DelegationManager.Contract.DelegationApproverSaltIsSpent(&_DelegationManager.CallOpts, arg0, arg1)
}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address , bytes32 ) view returns(bool)
func (_DelegationManager *DelegationManagerCallerSession) DelegationApproverSaltIsSpent(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _DelegationManager.Contract.DelegationApproverSaltIsSpent(&_DelegationManager.CallOpts, arg0, arg1)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_DelegationManager *DelegationManagerCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_DelegationManager *DelegationManagerSession) DomainSeparator() ([32]byte, error) {
	return _DelegationManager.Contract.DomainSeparator(&_DelegationManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_DelegationManager *DelegationManagerCallerSession) DomainSeparator() ([32]byte, error) {
	return _DelegationManager.Contract.DomainSeparator(&_DelegationManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_DelegationManager *DelegationManagerCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_DelegationManager *DelegationManagerSession) EigenPodManager() (common.Address, error) {
	return _DelegationManager.Contract.EigenPodManager(&_DelegationManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_DelegationManager *DelegationManagerCallerSession) EigenPodManager() (common.Address, error) {
	return _DelegationManager.Contract.EigenPodManager(&_DelegationManager.CallOpts)
}

// GetDelegatableShares is a free data retrieval call binding the contract method 0xcf80873e.
//
// Solidity: function getDelegatableShares(address staker) view returns(address[], uint256[])
func (_DelegationManager *DelegationManagerCaller) GetDelegatableShares(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "getDelegatableShares", staker)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDelegatableShares is a free data retrieval call binding the contract method 0xcf80873e.
//
// Solidity: function getDelegatableShares(address staker) view returns(address[], uint256[])
func (_DelegationManager *DelegationManagerSession) GetDelegatableShares(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _DelegationManager.Contract.GetDelegatableShares(&_DelegationManager.CallOpts, staker)
}

// GetDelegatableShares is a free data retrieval call binding the contract method 0xcf80873e.
//
// Solidity: function getDelegatableShares(address staker) view returns(address[], uint256[])
func (_DelegationManager *DelegationManagerCallerSession) GetDelegatableShares(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _DelegationManager.Contract.GetDelegatableShares(&_DelegationManager.CallOpts, staker)
}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] strategies) view returns(uint256[])
func (_DelegationManager *DelegationManagerCaller) GetOperatorShares(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "getOperatorShares", operator, strategies)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] strategies) view returns(uint256[])
func (_DelegationManager *DelegationManagerSession) GetOperatorShares(operator common.Address, strategies []common.Address) ([]*big.Int, error) {
	return _DelegationManager.Contract.GetOperatorShares(&_DelegationManager.CallOpts, operator, strategies)
}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] strategies) view returns(uint256[])
func (_DelegationManager *DelegationManagerCallerSession) GetOperatorShares(operator common.Address, strategies []common.Address) ([]*big.Int, error) {
	return _DelegationManager.Contract.GetOperatorShares(&_DelegationManager.CallOpts, operator, strategies)
}

// GetWithdrawalDelay is a free data retrieval call binding the contract method 0x0449ca39.
//
// Solidity: function getWithdrawalDelay(address[] strategies) view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) GetWithdrawalDelay(opts *bind.CallOpts, strategies []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "getWithdrawalDelay", strategies)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawalDelay is a free data retrieval call binding the contract method 0x0449ca39.
//
// Solidity: function getWithdrawalDelay(address[] strategies) view returns(uint256)
func (_DelegationManager *DelegationManagerSession) GetWithdrawalDelay(strategies []common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.GetWithdrawalDelay(&_DelegationManager.CallOpts, strategies)
}

// GetWithdrawalDelay is a free data retrieval call binding the contract method 0x0449ca39.
//
// Solidity: function getWithdrawalDelay(address[] strategies) view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) GetWithdrawalDelay(strategies []common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.GetWithdrawalDelay(&_DelegationManager.CallOpts, strategies)
}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_DelegationManager *DelegationManagerCaller) IsDelegated(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "isDelegated", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_DelegationManager *DelegationManagerSession) IsDelegated(staker common.Address) (bool, error) {
	return _DelegationManager.Contract.IsDelegated(&_DelegationManager.CallOpts, staker)
}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_DelegationManager *DelegationManagerCallerSession) IsDelegated(staker common.Address) (bool, error) {
	return _DelegationManager.Contract.IsDelegated(&_DelegationManager.CallOpts, staker)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_DelegationManager *DelegationManagerCaller) IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "isOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_DelegationManager *DelegationManagerSession) IsOperator(operator common.Address) (bool, error) {
	return _DelegationManager.Contract.IsOperator(&_DelegationManager.CallOpts, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_DelegationManager *DelegationManagerCallerSession) IsOperator(operator common.Address) (bool, error) {
	return _DelegationManager.Contract.IsOperator(&_DelegationManager.CallOpts, operator)
}

// MinWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0xc448feb8.
//
// Solidity: function minWithdrawalDelayBlocks() view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) MinWithdrawalDelayBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "minWithdrawalDelayBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0xc448feb8.
//
// Solidity: function minWithdrawalDelayBlocks() view returns(uint256)
func (_DelegationManager *DelegationManagerSession) MinWithdrawalDelayBlocks() (*big.Int, error) {
	return _DelegationManager.Contract.MinWithdrawalDelayBlocks(&_DelegationManager.CallOpts)
}

// MinWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0xc448feb8.
//
// Solidity: function minWithdrawalDelayBlocks() view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) MinWithdrawalDelayBlocks() (*big.Int, error) {
	return _DelegationManager.Contract.MinWithdrawalDelayBlocks(&_DelegationManager.CallOpts)
}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address operator) view returns((address,address,uint32))
func (_DelegationManager *DelegationManagerCaller) OperatorDetails(opts *bind.CallOpts, operator common.Address) (IDelegationManagerOperatorDetails, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "operatorDetails", operator)

	if err != nil {
		return *new(IDelegationManagerOperatorDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelegationManagerOperatorDetails)).(*IDelegationManagerOperatorDetails)

	return out0, err

}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address operator) view returns((address,address,uint32))
func (_DelegationManager *DelegationManagerSession) OperatorDetails(operator common.Address) (IDelegationManagerOperatorDetails, error) {
	return _DelegationManager.Contract.OperatorDetails(&_DelegationManager.CallOpts, operator)
}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address operator) view returns((address,address,uint32))
func (_DelegationManager *DelegationManagerCallerSession) OperatorDetails(operator common.Address) (IDelegationManagerOperatorDetails, error) {
	return _DelegationManager.Contract.OperatorDetails(&_DelegationManager.CallOpts, operator)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address , address ) view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) OperatorShares(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "operatorShares", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address , address ) view returns(uint256)
func (_DelegationManager *DelegationManagerSession) OperatorShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.OperatorShares(&_DelegationManager.CallOpts, arg0, arg1)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address , address ) view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) OperatorShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.OperatorShares(&_DelegationManager.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelegationManager *DelegationManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelegationManager *DelegationManagerSession) Owner() (common.Address, error) {
	return _DelegationManager.Contract.Owner(&_DelegationManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelegationManager *DelegationManagerCallerSession) Owner() (common.Address, error) {
	return _DelegationManager.Contract.Owner(&_DelegationManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_DelegationManager *DelegationManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_DelegationManager *DelegationManagerSession) Paused(index uint8) (bool, error) {
	return _DelegationManager.Contract.Paused(&_DelegationManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_DelegationManager *DelegationManagerCallerSession) Paused(index uint8) (bool, error) {
	return _DelegationManager.Contract.Paused(&_DelegationManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_DelegationManager *DelegationManagerSession) Paused0() (*big.Int, error) {
	return _DelegationManager.Contract.Paused0(&_DelegationManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) Paused0() (*big.Int, error) {
	return _DelegationManager.Contract.Paused0(&_DelegationManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DelegationManager *DelegationManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DelegationManager *DelegationManagerSession) PauserRegistry() (common.Address, error) {
	return _DelegationManager.Contract.PauserRegistry(&_DelegationManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_DelegationManager *DelegationManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _DelegationManager.Contract.PauserRegistry(&_DelegationManager.CallOpts)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 ) view returns(bool)
func (_DelegationManager *DelegationManagerCaller) PendingWithdrawals(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "pendingWithdrawals", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 ) view returns(bool)
func (_DelegationManager *DelegationManagerSession) PendingWithdrawals(arg0 [32]byte) (bool, error) {
	return _DelegationManager.Contract.PendingWithdrawals(&_DelegationManager.CallOpts, arg0)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 ) view returns(bool)
func (_DelegationManager *DelegationManagerCallerSession) PendingWithdrawals(arg0 [32]byte) (bool, error) {
	return _DelegationManager.Contract.PendingWithdrawals(&_DelegationManager.CallOpts, arg0)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_DelegationManager *DelegationManagerCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_DelegationManager *DelegationManagerSession) Slasher() (common.Address, error) {
	return _DelegationManager.Contract.Slasher(&_DelegationManager.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_DelegationManager *DelegationManagerCallerSession) Slasher() (common.Address, error) {
	return _DelegationManager.Contract.Slasher(&_DelegationManager.CallOpts)
}

// StakerNonce is a free data retrieval call binding the contract method 0x29c77d4f.
//
// Solidity: function stakerNonce(address ) view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) StakerNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "stakerNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerNonce is a free data retrieval call binding the contract method 0x29c77d4f.
//
// Solidity: function stakerNonce(address ) view returns(uint256)
func (_DelegationManager *DelegationManagerSession) StakerNonce(arg0 common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.StakerNonce(&_DelegationManager.CallOpts, arg0)
}

// StakerNonce is a free data retrieval call binding the contract method 0x29c77d4f.
//
// Solidity: function stakerNonce(address ) view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) StakerNonce(arg0 common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.StakerNonce(&_DelegationManager.CallOpts, arg0)
}

// StakerOptOutWindowBlocks is a free data retrieval call binding the contract method 0x16928365.
//
// Solidity: function stakerOptOutWindowBlocks(address operator) view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) StakerOptOutWindowBlocks(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "stakerOptOutWindowBlocks", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerOptOutWindowBlocks is a free data retrieval call binding the contract method 0x16928365.
//
// Solidity: function stakerOptOutWindowBlocks(address operator) view returns(uint256)
func (_DelegationManager *DelegationManagerSession) StakerOptOutWindowBlocks(operator common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.StakerOptOutWindowBlocks(&_DelegationManager.CallOpts, operator)
}

// StakerOptOutWindowBlocks is a free data retrieval call binding the contract method 0x16928365.
//
// Solidity: function stakerOptOutWindowBlocks(address operator) view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) StakerOptOutWindowBlocks(operator common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.StakerOptOutWindowBlocks(&_DelegationManager.CallOpts, operator)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_DelegationManager *DelegationManagerCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_DelegationManager *DelegationManagerSession) StrategyManager() (common.Address, error) {
	return _DelegationManager.Contract.StrategyManager(&_DelegationManager.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_DelegationManager *DelegationManagerCallerSession) StrategyManager() (common.Address, error) {
	return _DelegationManager.Contract.StrategyManager(&_DelegationManager.CallOpts)
}

// StrategyWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0xc488375a.
//
// Solidity: function strategyWithdrawalDelayBlocks(address ) view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) StrategyWithdrawalDelayBlocks(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "strategyWithdrawalDelayBlocks", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StrategyWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0xc488375a.
//
// Solidity: function strategyWithdrawalDelayBlocks(address ) view returns(uint256)
func (_DelegationManager *DelegationManagerSession) StrategyWithdrawalDelayBlocks(arg0 common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.StrategyWithdrawalDelayBlocks(&_DelegationManager.CallOpts, arg0)
}

// StrategyWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0xc488375a.
//
// Solidity: function strategyWithdrawalDelayBlocks(address ) view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) StrategyWithdrawalDelayBlocks(arg0 common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.StrategyWithdrawalDelayBlocks(&_DelegationManager.CallOpts, arg0)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0x60d7faed.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerTransactor) CompleteQueuedWithdrawal(opts *bind.TransactOpts, withdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "completeQueuedWithdrawal", withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0x60d7faed.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerSession) CompleteQueuedWithdrawal(withdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _DelegationManager.Contract.CompleteQueuedWithdrawal(&_DelegationManager.TransactOpts, withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0x60d7faed.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, uint256 middlewareTimesIndex, bool receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerTransactorSession) CompleteQueuedWithdrawal(withdrawal IDelegationManagerWithdrawal, tokens []common.Address, middlewareTimesIndex *big.Int, receiveAsTokens bool) (*types.Transaction, error) {
	return _DelegationManager.Contract.CompleteQueuedWithdrawal(&_DelegationManager.TransactOpts, withdrawal, tokens, middlewareTimesIndex, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x33404396.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerTransactor) CompleteQueuedWithdrawals(opts *bind.TransactOpts, withdrawals []IDelegationManagerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "completeQueuedWithdrawals", withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x33404396.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerSession) CompleteQueuedWithdrawals(withdrawals []IDelegationManagerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _DelegationManager.Contract.CompleteQueuedWithdrawals(&_DelegationManager.TransactOpts, withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x33404396.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerTransactorSession) CompleteQueuedWithdrawals(withdrawals []IDelegationManagerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _DelegationManager.Contract.CompleteQueuedWithdrawals(&_DelegationManager.TransactOpts, withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address strategy, uint256 shares) returns()
func (_DelegationManager *DelegationManagerTransactor) DecreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "decreaseDelegatedShares", staker, strategy, shares)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address strategy, uint256 shares) returns()
func (_DelegationManager *DelegationManagerSession) DecreaseDelegatedShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.DecreaseDelegatedShares(&_DelegationManager.TransactOpts, staker, strategy, shares)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address strategy, uint256 shares) returns()
func (_DelegationManager *DelegationManagerTransactorSession) DecreaseDelegatedShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.DecreaseDelegatedShares(&_DelegationManager.TransactOpts, staker, strategy, shares)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationManager *DelegationManagerTransactor) DelegateTo(opts *bind.TransactOpts, operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "delegateTo", operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationManager *DelegationManagerSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationManager.Contract.DelegateTo(&_DelegationManager.TransactOpts, operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationManager *DelegationManagerTransactorSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationManager.Contract.DelegateTo(&_DelegationManager.TransactOpts, operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateToBySignature is a paid mutator transaction binding the contract method 0x7f548071.
//
// Solidity: function delegateToBySignature(address staker, address operator, (bytes,uint256) stakerSignatureAndExpiry, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationManager *DelegationManagerTransactor) DelegateToBySignature(opts *bind.TransactOpts, staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "delegateToBySignature", staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
}

// DelegateToBySignature is a paid mutator transaction binding the contract method 0x7f548071.
//
// Solidity: function delegateToBySignature(address staker, address operator, (bytes,uint256) stakerSignatureAndExpiry, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationManager *DelegationManagerSession) DelegateToBySignature(staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationManager.Contract.DelegateToBySignature(&_DelegationManager.TransactOpts, staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
}

// DelegateToBySignature is a paid mutator transaction binding the contract method 0x7f548071.
//
// Solidity: function delegateToBySignature(address staker, address operator, (bytes,uint256) stakerSignatureAndExpiry, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationManager *DelegationManagerTransactorSession) DelegateToBySignature(staker common.Address, operator common.Address, stakerSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationManager.Contract.DelegateToBySignature(&_DelegationManager.TransactOpts, staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, approverSalt)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x28a573ae.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 shares) returns()
func (_DelegationManager *DelegationManagerTransactor) IncreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "increaseDelegatedShares", staker, strategy, shares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x28a573ae.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 shares) returns()
func (_DelegationManager *DelegationManagerSession) IncreaseDelegatedShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.IncreaseDelegatedShares(&_DelegationManager.TransactOpts, staker, strategy, shares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x28a573ae.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 shares) returns()
func (_DelegationManager *DelegationManagerTransactorSession) IncreaseDelegatedShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.IncreaseDelegatedShares(&_DelegationManager.TransactOpts, staker, strategy, shares)
}

// Initialize is a paid mutator transaction binding the contract method 0x22bf40e4.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus, uint256 _minWithdrawalDelayBlocks, address[] _strategies, uint256[] _withdrawalDelayBlocks) returns()
func (_DelegationManager *DelegationManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int, _minWithdrawalDelayBlocks *big.Int, _strategies []common.Address, _withdrawalDelayBlocks []*big.Int) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "initialize", initialOwner, _pauserRegistry, initialPausedStatus, _minWithdrawalDelayBlocks, _strategies, _withdrawalDelayBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0x22bf40e4.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus, uint256 _minWithdrawalDelayBlocks, address[] _strategies, uint256[] _withdrawalDelayBlocks) returns()
func (_DelegationManager *DelegationManagerSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int, _minWithdrawalDelayBlocks *big.Int, _strategies []common.Address, _withdrawalDelayBlocks []*big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.Initialize(&_DelegationManager.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus, _minWithdrawalDelayBlocks, _strategies, _withdrawalDelayBlocks)
}

// Initialize is a paid mutator transaction binding the contract method 0x22bf40e4.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus, uint256 _minWithdrawalDelayBlocks, address[] _strategies, uint256[] _withdrawalDelayBlocks) returns()
func (_DelegationManager *DelegationManagerTransactorSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int, _minWithdrawalDelayBlocks *big.Int, _strategies []common.Address, _withdrawalDelayBlocks []*big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.Initialize(&_DelegationManager.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus, _minWithdrawalDelayBlocks, _strategies, _withdrawalDelayBlocks)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xf16172b0.
//
// Solidity: function modifyOperatorDetails((address,address,uint32) newOperatorDetails) returns()
func (_DelegationManager *DelegationManagerTransactor) ModifyOperatorDetails(opts *bind.TransactOpts, newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "modifyOperatorDetails", newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xf16172b0.
//
// Solidity: function modifyOperatorDetails((address,address,uint32) newOperatorDetails) returns()
func (_DelegationManager *DelegationManagerSession) ModifyOperatorDetails(newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _DelegationManager.Contract.ModifyOperatorDetails(&_DelegationManager.TransactOpts, newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xf16172b0.
//
// Solidity: function modifyOperatorDetails((address,address,uint32) newOperatorDetails) returns()
func (_DelegationManager *DelegationManagerTransactorSession) ModifyOperatorDetails(newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _DelegationManager.Contract.ModifyOperatorDetails(&_DelegationManager.TransactOpts, newOperatorDetails)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_DelegationManager *DelegationManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_DelegationManager *DelegationManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.Pause(&_DelegationManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_DelegationManager *DelegationManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.Pause(&_DelegationManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_DelegationManager *DelegationManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_DelegationManager *DelegationManagerSession) PauseAll() (*types.Transaction, error) {
	return _DelegationManager.Contract.PauseAll(&_DelegationManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_DelegationManager *DelegationManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _DelegationManager.Contract.PauseAll(&_DelegationManager.TransactOpts)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes32[])
func (_DelegationManager *DelegationManagerTransactor) QueueWithdrawals(opts *bind.TransactOpts, queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "queueWithdrawals", queuedWithdrawalParams)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes32[])
func (_DelegationManager *DelegationManagerSession) QueueWithdrawals(queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _DelegationManager.Contract.QueueWithdrawals(&_DelegationManager.TransactOpts, queuedWithdrawalParams)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes32[])
func (_DelegationManager *DelegationManagerTransactorSession) QueueWithdrawals(queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _DelegationManager.Contract.QueueWithdrawals(&_DelegationManager.TransactOpts, queuedWithdrawalParams)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x0f589e59.
//
// Solidity: function registerAsOperator((address,address,uint32) registeringOperatorDetails, string metadataURI) returns()
func (_DelegationManager *DelegationManagerTransactor) RegisterAsOperator(opts *bind.TransactOpts, registeringOperatorDetails IDelegationManagerOperatorDetails, metadataURI string) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "registerAsOperator", registeringOperatorDetails, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x0f589e59.
//
// Solidity: function registerAsOperator((address,address,uint32) registeringOperatorDetails, string metadataURI) returns()
func (_DelegationManager *DelegationManagerSession) RegisterAsOperator(registeringOperatorDetails IDelegationManagerOperatorDetails, metadataURI string) (*types.Transaction, error) {
	return _DelegationManager.Contract.RegisterAsOperator(&_DelegationManager.TransactOpts, registeringOperatorDetails, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x0f589e59.
//
// Solidity: function registerAsOperator((address,address,uint32) registeringOperatorDetails, string metadataURI) returns()
func (_DelegationManager *DelegationManagerTransactorSession) RegisterAsOperator(registeringOperatorDetails IDelegationManagerOperatorDetails, metadataURI string) (*types.Transaction, error) {
	return _DelegationManager.Contract.RegisterAsOperator(&_DelegationManager.TransactOpts, registeringOperatorDetails, metadataURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelegationManager *DelegationManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelegationManager *DelegationManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _DelegationManager.Contract.RenounceOwnership(&_DelegationManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelegationManager *DelegationManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DelegationManager.Contract.RenounceOwnership(&_DelegationManager.TransactOpts)
}

// SetMinWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x635bbd10.
//
// Solidity: function setMinWithdrawalDelayBlocks(uint256 newMinWithdrawalDelayBlocks) returns()
func (_DelegationManager *DelegationManagerTransactor) SetMinWithdrawalDelayBlocks(opts *bind.TransactOpts, newMinWithdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "setMinWithdrawalDelayBlocks", newMinWithdrawalDelayBlocks)
}

// SetMinWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x635bbd10.
//
// Solidity: function setMinWithdrawalDelayBlocks(uint256 newMinWithdrawalDelayBlocks) returns()
func (_DelegationManager *DelegationManagerSession) SetMinWithdrawalDelayBlocks(newMinWithdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.SetMinWithdrawalDelayBlocks(&_DelegationManager.TransactOpts, newMinWithdrawalDelayBlocks)
}

// SetMinWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x635bbd10.
//
// Solidity: function setMinWithdrawalDelayBlocks(uint256 newMinWithdrawalDelayBlocks) returns()
func (_DelegationManager *DelegationManagerTransactorSession) SetMinWithdrawalDelayBlocks(newMinWithdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.SetMinWithdrawalDelayBlocks(&_DelegationManager.TransactOpts, newMinWithdrawalDelayBlocks)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_DelegationManager *DelegationManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_DelegationManager *DelegationManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _DelegationManager.Contract.SetPauserRegistry(&_DelegationManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_DelegationManager *DelegationManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _DelegationManager.Contract.SetPauserRegistry(&_DelegationManager.TransactOpts, newPauserRegistry)
}

// SetStrategyWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x1522bf02.
//
// Solidity: function setStrategyWithdrawalDelayBlocks(address[] strategies, uint256[] withdrawalDelayBlocks) returns()
func (_DelegationManager *DelegationManagerTransactor) SetStrategyWithdrawalDelayBlocks(opts *bind.TransactOpts, strategies []common.Address, withdrawalDelayBlocks []*big.Int) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "setStrategyWithdrawalDelayBlocks", strategies, withdrawalDelayBlocks)
}

// SetStrategyWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x1522bf02.
//
// Solidity: function setStrategyWithdrawalDelayBlocks(address[] strategies, uint256[] withdrawalDelayBlocks) returns()
func (_DelegationManager *DelegationManagerSession) SetStrategyWithdrawalDelayBlocks(strategies []common.Address, withdrawalDelayBlocks []*big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.SetStrategyWithdrawalDelayBlocks(&_DelegationManager.TransactOpts, strategies, withdrawalDelayBlocks)
}

// SetStrategyWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x1522bf02.
//
// Solidity: function setStrategyWithdrawalDelayBlocks(address[] strategies, uint256[] withdrawalDelayBlocks) returns()
func (_DelegationManager *DelegationManagerTransactorSession) SetStrategyWithdrawalDelayBlocks(strategies []common.Address, withdrawalDelayBlocks []*big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.SetStrategyWithdrawalDelayBlocks(&_DelegationManager.TransactOpts, strategies, withdrawalDelayBlocks)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelegationManager *DelegationManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelegationManager *DelegationManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelegationManager.Contract.TransferOwnership(&_DelegationManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelegationManager *DelegationManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelegationManager.Contract.TransferOwnership(&_DelegationManager.TransactOpts, newOwner)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoots)
func (_DelegationManager *DelegationManagerTransactor) Undelegate(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "undelegate", staker)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoots)
func (_DelegationManager *DelegationManagerSession) Undelegate(staker common.Address) (*types.Transaction, error) {
	return _DelegationManager.Contract.Undelegate(&_DelegationManager.TransactOpts, staker)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoots)
func (_DelegationManager *DelegationManagerTransactorSession) Undelegate(staker common.Address) (*types.Transaction, error) {
	return _DelegationManager.Contract.Undelegate(&_DelegationManager.TransactOpts, staker)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_DelegationManager *DelegationManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_DelegationManager *DelegationManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.Unpause(&_DelegationManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_DelegationManager *DelegationManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.Unpause(&_DelegationManager.TransactOpts, newPausedStatus)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string metadataURI) returns()
func (_DelegationManager *DelegationManagerTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "updateOperatorMetadataURI", metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string metadataURI) returns()
func (_DelegationManager *DelegationManagerSession) UpdateOperatorMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _DelegationManager.Contract.UpdateOperatorMetadataURI(&_DelegationManager.TransactOpts, metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x99be81c8.
//
// Solidity: function updateOperatorMetadataURI(string metadataURI) returns()
func (_DelegationManager *DelegationManagerTransactorSession) UpdateOperatorMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _DelegationManager.Contract.UpdateOperatorMetadataURI(&_DelegationManager.TransactOpts, metadataURI)
}

// DelegationManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DelegationManager contract.
type DelegationManagerInitializedIterator struct {
	Event *DelegationManagerInitialized // Event containing the contract specifics and raw log

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
func (it *DelegationManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerInitialized)
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
		it.Event = new(DelegationManagerInitialized)
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
func (it *DelegationManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerInitialized represents a Initialized event raised by the DelegationManager contract.
type DelegationManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DelegationManager *DelegationManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*DelegationManagerInitializedIterator, error) {

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DelegationManagerInitializedIterator{contract: _DelegationManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DelegationManager *DelegationManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DelegationManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerInitialized)
				if err := _DelegationManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DelegationManager *DelegationManagerFilterer) ParseInitialized(log types.Log) (*DelegationManagerInitialized, error) {
	event := new(DelegationManagerInitialized)
	if err := _DelegationManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerMinWithdrawalDelayBlocksSetIterator is returned from FilterMinWithdrawalDelayBlocksSet and is used to iterate over the raw logs and unpacked data for MinWithdrawalDelayBlocksSet events raised by the DelegationManager contract.
type DelegationManagerMinWithdrawalDelayBlocksSetIterator struct {
	Event *DelegationManagerMinWithdrawalDelayBlocksSet // Event containing the contract specifics and raw log

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
func (it *DelegationManagerMinWithdrawalDelayBlocksSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerMinWithdrawalDelayBlocksSet)
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
		it.Event = new(DelegationManagerMinWithdrawalDelayBlocksSet)
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
func (it *DelegationManagerMinWithdrawalDelayBlocksSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerMinWithdrawalDelayBlocksSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerMinWithdrawalDelayBlocksSet represents a MinWithdrawalDelayBlocksSet event raised by the DelegationManager contract.
type DelegationManagerMinWithdrawalDelayBlocksSet struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMinWithdrawalDelayBlocksSet is a free log retrieval operation binding the contract event 0xafa003cd76f87ff9d62b35beea889920f33c0c42b8d45b74954d61d50f4b6b69.
//
// Solidity: event MinWithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_DelegationManager *DelegationManagerFilterer) FilterMinWithdrawalDelayBlocksSet(opts *bind.FilterOpts) (*DelegationManagerMinWithdrawalDelayBlocksSetIterator, error) {

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "MinWithdrawalDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return &DelegationManagerMinWithdrawalDelayBlocksSetIterator{contract: _DelegationManager.contract, event: "MinWithdrawalDelayBlocksSet", logs: logs, sub: sub}, nil
}

// WatchMinWithdrawalDelayBlocksSet is a free log subscription operation binding the contract event 0xafa003cd76f87ff9d62b35beea889920f33c0c42b8d45b74954d61d50f4b6b69.
//
// Solidity: event MinWithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_DelegationManager *DelegationManagerFilterer) WatchMinWithdrawalDelayBlocksSet(opts *bind.WatchOpts, sink chan<- *DelegationManagerMinWithdrawalDelayBlocksSet) (event.Subscription, error) {

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "MinWithdrawalDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerMinWithdrawalDelayBlocksSet)
				if err := _DelegationManager.contract.UnpackLog(event, "MinWithdrawalDelayBlocksSet", log); err != nil {
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

// ParseMinWithdrawalDelayBlocksSet is a log parse operation binding the contract event 0xafa003cd76f87ff9d62b35beea889920f33c0c42b8d45b74954d61d50f4b6b69.
//
// Solidity: event MinWithdrawalDelayBlocksSet(uint256 previousValue, uint256 newValue)
func (_DelegationManager *DelegationManagerFilterer) ParseMinWithdrawalDelayBlocksSet(log types.Log) (*DelegationManagerMinWithdrawalDelayBlocksSet, error) {
	event := new(DelegationManagerMinWithdrawalDelayBlocksSet)
	if err := _DelegationManager.contract.UnpackLog(event, "MinWithdrawalDelayBlocksSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerOperatorDetailsModifiedIterator is returned from FilterOperatorDetailsModified and is used to iterate over the raw logs and unpacked data for OperatorDetailsModified events raised by the DelegationManager contract.
type DelegationManagerOperatorDetailsModifiedIterator struct {
	Event *DelegationManagerOperatorDetailsModified // Event containing the contract specifics and raw log

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
func (it *DelegationManagerOperatorDetailsModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerOperatorDetailsModified)
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
		it.Event = new(DelegationManagerOperatorDetailsModified)
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
func (it *DelegationManagerOperatorDetailsModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerOperatorDetailsModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerOperatorDetailsModified represents a OperatorDetailsModified event raised by the DelegationManager contract.
type DelegationManagerOperatorDetailsModified struct {
	Operator           common.Address
	NewOperatorDetails IDelegationManagerOperatorDetails
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOperatorDetailsModified is a free log retrieval operation binding the contract event 0xfebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac.
//
// Solidity: event OperatorDetailsModified(address indexed operator, (address,address,uint32) newOperatorDetails)
func (_DelegationManager *DelegationManagerFilterer) FilterOperatorDetailsModified(opts *bind.FilterOpts, operator []common.Address) (*DelegationManagerOperatorDetailsModifiedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "OperatorDetailsModified", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerOperatorDetailsModifiedIterator{contract: _DelegationManager.contract, event: "OperatorDetailsModified", logs: logs, sub: sub}, nil
}

// WatchOperatorDetailsModified is a free log subscription operation binding the contract event 0xfebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac.
//
// Solidity: event OperatorDetailsModified(address indexed operator, (address,address,uint32) newOperatorDetails)
func (_DelegationManager *DelegationManagerFilterer) WatchOperatorDetailsModified(opts *bind.WatchOpts, sink chan<- *DelegationManagerOperatorDetailsModified, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "OperatorDetailsModified", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerOperatorDetailsModified)
				if err := _DelegationManager.contract.UnpackLog(event, "OperatorDetailsModified", log); err != nil {
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

// ParseOperatorDetailsModified is a log parse operation binding the contract event 0xfebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac.
//
// Solidity: event OperatorDetailsModified(address indexed operator, (address,address,uint32) newOperatorDetails)
func (_DelegationManager *DelegationManagerFilterer) ParseOperatorDetailsModified(log types.Log) (*DelegationManagerOperatorDetailsModified, error) {
	event := new(DelegationManagerOperatorDetailsModified)
	if err := _DelegationManager.contract.UnpackLog(event, "OperatorDetailsModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerOperatorMetadataURIUpdatedIterator is returned from FilterOperatorMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for OperatorMetadataURIUpdated events raised by the DelegationManager contract.
type DelegationManagerOperatorMetadataURIUpdatedIterator struct {
	Event *DelegationManagerOperatorMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *DelegationManagerOperatorMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerOperatorMetadataURIUpdated)
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
		it.Event = new(DelegationManagerOperatorMetadataURIUpdated)
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
func (it *DelegationManagerOperatorMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerOperatorMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerOperatorMetadataURIUpdated represents a OperatorMetadataURIUpdated event raised by the DelegationManager contract.
type DelegationManagerOperatorMetadataURIUpdated struct {
	Operator    common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorMetadataURIUpdated is a free log retrieval operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_DelegationManager *DelegationManagerFilterer) FilterOperatorMetadataURIUpdated(opts *bind.FilterOpts, operator []common.Address) (*DelegationManagerOperatorMetadataURIUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "OperatorMetadataURIUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerOperatorMetadataURIUpdatedIterator{contract: _DelegationManager.contract, event: "OperatorMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorMetadataURIUpdated is a free log subscription operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_DelegationManager *DelegationManagerFilterer) WatchOperatorMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *DelegationManagerOperatorMetadataURIUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "OperatorMetadataURIUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerOperatorMetadataURIUpdated)
				if err := _DelegationManager.contract.UnpackLog(event, "OperatorMetadataURIUpdated", log); err != nil {
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

// ParseOperatorMetadataURIUpdated is a log parse operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_DelegationManager *DelegationManagerFilterer) ParseOperatorMetadataURIUpdated(log types.Log) (*DelegationManagerOperatorMetadataURIUpdated, error) {
	event := new(DelegationManagerOperatorMetadataURIUpdated)
	if err := _DelegationManager.contract.UnpackLog(event, "OperatorMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the DelegationManager contract.
type DelegationManagerOperatorRegisteredIterator struct {
	Event *DelegationManagerOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *DelegationManagerOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerOperatorRegistered)
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
		it.Event = new(DelegationManagerOperatorRegistered)
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
func (it *DelegationManagerOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerOperatorRegistered represents a OperatorRegistered event raised by the DelegationManager contract.
type DelegationManagerOperatorRegistered struct {
	Operator        common.Address
	OperatorDetails IDelegationManagerOperatorDetails
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e2.
//
// Solidity: event OperatorRegistered(address indexed operator, (address,address,uint32) operatorDetails)
func (_DelegationManager *DelegationManagerFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*DelegationManagerOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerOperatorRegisteredIterator{contract: _DelegationManager.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e2.
//
// Solidity: event OperatorRegistered(address indexed operator, (address,address,uint32) operatorDetails)
func (_DelegationManager *DelegationManagerFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *DelegationManagerOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerOperatorRegistered)
				if err := _DelegationManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e2.
//
// Solidity: event OperatorRegistered(address indexed operator, (address,address,uint32) operatorDetails)
func (_DelegationManager *DelegationManagerFilterer) ParseOperatorRegistered(log types.Log) (*DelegationManagerOperatorRegistered, error) {
	event := new(DelegationManagerOperatorRegistered)
	if err := _DelegationManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerOperatorSharesDecreasedIterator is returned from FilterOperatorSharesDecreased and is used to iterate over the raw logs and unpacked data for OperatorSharesDecreased events raised by the DelegationManager contract.
type DelegationManagerOperatorSharesDecreasedIterator struct {
	Event *DelegationManagerOperatorSharesDecreased // Event containing the contract specifics and raw log

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
func (it *DelegationManagerOperatorSharesDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerOperatorSharesDecreased)
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
		it.Event = new(DelegationManagerOperatorSharesDecreased)
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
func (it *DelegationManagerOperatorSharesDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerOperatorSharesDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerOperatorSharesDecreased represents a OperatorSharesDecreased event raised by the DelegationManager contract.
type DelegationManagerOperatorSharesDecreased struct {
	Operator common.Address
	Staker   common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesDecreased is a free log retrieval operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_DelegationManager *DelegationManagerFilterer) FilterOperatorSharesDecreased(opts *bind.FilterOpts, operator []common.Address) (*DelegationManagerOperatorSharesDecreasedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "OperatorSharesDecreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerOperatorSharesDecreasedIterator{contract: _DelegationManager.contract, event: "OperatorSharesDecreased", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesDecreased is a free log subscription operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_DelegationManager *DelegationManagerFilterer) WatchOperatorSharesDecreased(opts *bind.WatchOpts, sink chan<- *DelegationManagerOperatorSharesDecreased, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "OperatorSharesDecreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerOperatorSharesDecreased)
				if err := _DelegationManager.contract.UnpackLog(event, "OperatorSharesDecreased", log); err != nil {
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

// ParseOperatorSharesDecreased is a log parse operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_DelegationManager *DelegationManagerFilterer) ParseOperatorSharesDecreased(log types.Log) (*DelegationManagerOperatorSharesDecreased, error) {
	event := new(DelegationManagerOperatorSharesDecreased)
	if err := _DelegationManager.contract.UnpackLog(event, "OperatorSharesDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerOperatorSharesIncreasedIterator is returned from FilterOperatorSharesIncreased and is used to iterate over the raw logs and unpacked data for OperatorSharesIncreased events raised by the DelegationManager contract.
type DelegationManagerOperatorSharesIncreasedIterator struct {
	Event *DelegationManagerOperatorSharesIncreased // Event containing the contract specifics and raw log

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
func (it *DelegationManagerOperatorSharesIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerOperatorSharesIncreased)
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
		it.Event = new(DelegationManagerOperatorSharesIncreased)
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
func (it *DelegationManagerOperatorSharesIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerOperatorSharesIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerOperatorSharesIncreased represents a OperatorSharesIncreased event raised by the DelegationManager contract.
type DelegationManagerOperatorSharesIncreased struct {
	Operator common.Address
	Staker   common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesIncreased is a free log retrieval operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_DelegationManager *DelegationManagerFilterer) FilterOperatorSharesIncreased(opts *bind.FilterOpts, operator []common.Address) (*DelegationManagerOperatorSharesIncreasedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "OperatorSharesIncreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerOperatorSharesIncreasedIterator{contract: _DelegationManager.contract, event: "OperatorSharesIncreased", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesIncreased is a free log subscription operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_DelegationManager *DelegationManagerFilterer) WatchOperatorSharesIncreased(opts *bind.WatchOpts, sink chan<- *DelegationManagerOperatorSharesIncreased, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "OperatorSharesIncreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerOperatorSharesIncreased)
				if err := _DelegationManager.contract.UnpackLog(event, "OperatorSharesIncreased", log); err != nil {
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

// ParseOperatorSharesIncreased is a log parse operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_DelegationManager *DelegationManagerFilterer) ParseOperatorSharesIncreased(log types.Log) (*DelegationManagerOperatorSharesIncreased, error) {
	event := new(DelegationManagerOperatorSharesIncreased)
	if err := _DelegationManager.contract.UnpackLog(event, "OperatorSharesIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DelegationManager contract.
type DelegationManagerOwnershipTransferredIterator struct {
	Event *DelegationManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DelegationManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerOwnershipTransferred)
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
		it.Event = new(DelegationManagerOwnershipTransferred)
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
func (it *DelegationManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerOwnershipTransferred represents a OwnershipTransferred event raised by the DelegationManager contract.
type DelegationManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelegationManager *DelegationManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DelegationManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerOwnershipTransferredIterator{contract: _DelegationManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelegationManager *DelegationManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DelegationManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerOwnershipTransferred)
				if err := _DelegationManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DelegationManager *DelegationManagerFilterer) ParseOwnershipTransferred(log types.Log) (*DelegationManagerOwnershipTransferred, error) {
	event := new(DelegationManagerOwnershipTransferred)
	if err := _DelegationManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the DelegationManager contract.
type DelegationManagerPausedIterator struct {
	Event *DelegationManagerPaused // Event containing the contract specifics and raw log

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
func (it *DelegationManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerPaused)
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
		it.Event = new(DelegationManagerPaused)
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
func (it *DelegationManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerPaused represents a Paused event raised by the DelegationManager contract.
type DelegationManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_DelegationManager *DelegationManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*DelegationManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerPausedIterator{contract: _DelegationManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_DelegationManager *DelegationManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *DelegationManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerPaused)
				if err := _DelegationManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_DelegationManager *DelegationManagerFilterer) ParsePaused(log types.Log) (*DelegationManagerPaused, error) {
	event := new(DelegationManagerPaused)
	if err := _DelegationManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the DelegationManager contract.
type DelegationManagerPauserRegistrySetIterator struct {
	Event *DelegationManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *DelegationManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerPauserRegistrySet)
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
		it.Event = new(DelegationManagerPauserRegistrySet)
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
func (it *DelegationManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerPauserRegistrySet represents a PauserRegistrySet event raised by the DelegationManager contract.
type DelegationManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_DelegationManager *DelegationManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*DelegationManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &DelegationManagerPauserRegistrySetIterator{contract: _DelegationManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_DelegationManager *DelegationManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *DelegationManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerPauserRegistrySet)
				if err := _DelegationManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_DelegationManager *DelegationManagerFilterer) ParsePauserRegistrySet(log types.Log) (*DelegationManagerPauserRegistrySet, error) {
	event := new(DelegationManagerPauserRegistrySet)
	if err := _DelegationManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerStakerDelegatedIterator is returned from FilterStakerDelegated and is used to iterate over the raw logs and unpacked data for StakerDelegated events raised by the DelegationManager contract.
type DelegationManagerStakerDelegatedIterator struct {
	Event *DelegationManagerStakerDelegated // Event containing the contract specifics and raw log

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
func (it *DelegationManagerStakerDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerStakerDelegated)
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
		it.Event = new(DelegationManagerStakerDelegated)
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
func (it *DelegationManagerStakerDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerStakerDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerStakerDelegated represents a StakerDelegated event raised by the DelegationManager contract.
type DelegationManagerStakerDelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerDelegated is a free log retrieval operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_DelegationManager *DelegationManagerFilterer) FilterStakerDelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*DelegationManagerStakerDelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "StakerDelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerStakerDelegatedIterator{contract: _DelegationManager.contract, event: "StakerDelegated", logs: logs, sub: sub}, nil
}

// WatchStakerDelegated is a free log subscription operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_DelegationManager *DelegationManagerFilterer) WatchStakerDelegated(opts *bind.WatchOpts, sink chan<- *DelegationManagerStakerDelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "StakerDelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerStakerDelegated)
				if err := _DelegationManager.contract.UnpackLog(event, "StakerDelegated", log); err != nil {
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

// ParseStakerDelegated is a log parse operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_DelegationManager *DelegationManagerFilterer) ParseStakerDelegated(log types.Log) (*DelegationManagerStakerDelegated, error) {
	event := new(DelegationManagerStakerDelegated)
	if err := _DelegationManager.contract.UnpackLog(event, "StakerDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerStakerForceUndelegatedIterator is returned from FilterStakerForceUndelegated and is used to iterate over the raw logs and unpacked data for StakerForceUndelegated events raised by the DelegationManager contract.
type DelegationManagerStakerForceUndelegatedIterator struct {
	Event *DelegationManagerStakerForceUndelegated // Event containing the contract specifics and raw log

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
func (it *DelegationManagerStakerForceUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerStakerForceUndelegated)
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
		it.Event = new(DelegationManagerStakerForceUndelegated)
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
func (it *DelegationManagerStakerForceUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerStakerForceUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerStakerForceUndelegated represents a StakerForceUndelegated event raised by the DelegationManager contract.
type DelegationManagerStakerForceUndelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerForceUndelegated is a free log retrieval operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_DelegationManager *DelegationManagerFilterer) FilterStakerForceUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*DelegationManagerStakerForceUndelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "StakerForceUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerStakerForceUndelegatedIterator{contract: _DelegationManager.contract, event: "StakerForceUndelegated", logs: logs, sub: sub}, nil
}

// WatchStakerForceUndelegated is a free log subscription operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_DelegationManager *DelegationManagerFilterer) WatchStakerForceUndelegated(opts *bind.WatchOpts, sink chan<- *DelegationManagerStakerForceUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "StakerForceUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerStakerForceUndelegated)
				if err := _DelegationManager.contract.UnpackLog(event, "StakerForceUndelegated", log); err != nil {
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

// ParseStakerForceUndelegated is a log parse operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_DelegationManager *DelegationManagerFilterer) ParseStakerForceUndelegated(log types.Log) (*DelegationManagerStakerForceUndelegated, error) {
	event := new(DelegationManagerStakerForceUndelegated)
	if err := _DelegationManager.contract.UnpackLog(event, "StakerForceUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerStakerUndelegatedIterator is returned from FilterStakerUndelegated and is used to iterate over the raw logs and unpacked data for StakerUndelegated events raised by the DelegationManager contract.
type DelegationManagerStakerUndelegatedIterator struct {
	Event *DelegationManagerStakerUndelegated // Event containing the contract specifics and raw log

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
func (it *DelegationManagerStakerUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerStakerUndelegated)
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
		it.Event = new(DelegationManagerStakerUndelegated)
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
func (it *DelegationManagerStakerUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerStakerUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerStakerUndelegated represents a StakerUndelegated event raised by the DelegationManager contract.
type DelegationManagerStakerUndelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerUndelegated is a free log retrieval operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_DelegationManager *DelegationManagerFilterer) FilterStakerUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*DelegationManagerStakerUndelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "StakerUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerStakerUndelegatedIterator{contract: _DelegationManager.contract, event: "StakerUndelegated", logs: logs, sub: sub}, nil
}

// WatchStakerUndelegated is a free log subscription operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_DelegationManager *DelegationManagerFilterer) WatchStakerUndelegated(opts *bind.WatchOpts, sink chan<- *DelegationManagerStakerUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "StakerUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerStakerUndelegated)
				if err := _DelegationManager.contract.UnpackLog(event, "StakerUndelegated", log); err != nil {
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

// ParseStakerUndelegated is a log parse operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_DelegationManager *DelegationManagerFilterer) ParseStakerUndelegated(log types.Log) (*DelegationManagerStakerUndelegated, error) {
	event := new(DelegationManagerStakerUndelegated)
	if err := _DelegationManager.contract.UnpackLog(event, "StakerUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerStrategyWithdrawalDelayBlocksSetIterator is returned from FilterStrategyWithdrawalDelayBlocksSet and is used to iterate over the raw logs and unpacked data for StrategyWithdrawalDelayBlocksSet events raised by the DelegationManager contract.
type DelegationManagerStrategyWithdrawalDelayBlocksSetIterator struct {
	Event *DelegationManagerStrategyWithdrawalDelayBlocksSet // Event containing the contract specifics and raw log

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
func (it *DelegationManagerStrategyWithdrawalDelayBlocksSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerStrategyWithdrawalDelayBlocksSet)
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
		it.Event = new(DelegationManagerStrategyWithdrawalDelayBlocksSet)
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
func (it *DelegationManagerStrategyWithdrawalDelayBlocksSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerStrategyWithdrawalDelayBlocksSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerStrategyWithdrawalDelayBlocksSet represents a StrategyWithdrawalDelayBlocksSet event raised by the DelegationManager contract.
type DelegationManagerStrategyWithdrawalDelayBlocksSet struct {
	Strategy      common.Address
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStrategyWithdrawalDelayBlocksSet is a free log retrieval operation binding the contract event 0x0e7efa738e8b0ce6376a0c1af471655540d2e9a81647d7b09ed823018426576d.
//
// Solidity: event StrategyWithdrawalDelayBlocksSet(address strategy, uint256 previousValue, uint256 newValue)
func (_DelegationManager *DelegationManagerFilterer) FilterStrategyWithdrawalDelayBlocksSet(opts *bind.FilterOpts) (*DelegationManagerStrategyWithdrawalDelayBlocksSetIterator, error) {

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "StrategyWithdrawalDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return &DelegationManagerStrategyWithdrawalDelayBlocksSetIterator{contract: _DelegationManager.contract, event: "StrategyWithdrawalDelayBlocksSet", logs: logs, sub: sub}, nil
}

// WatchStrategyWithdrawalDelayBlocksSet is a free log subscription operation binding the contract event 0x0e7efa738e8b0ce6376a0c1af471655540d2e9a81647d7b09ed823018426576d.
//
// Solidity: event StrategyWithdrawalDelayBlocksSet(address strategy, uint256 previousValue, uint256 newValue)
func (_DelegationManager *DelegationManagerFilterer) WatchStrategyWithdrawalDelayBlocksSet(opts *bind.WatchOpts, sink chan<- *DelegationManagerStrategyWithdrawalDelayBlocksSet) (event.Subscription, error) {

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "StrategyWithdrawalDelayBlocksSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerStrategyWithdrawalDelayBlocksSet)
				if err := _DelegationManager.contract.UnpackLog(event, "StrategyWithdrawalDelayBlocksSet", log); err != nil {
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

// ParseStrategyWithdrawalDelayBlocksSet is a log parse operation binding the contract event 0x0e7efa738e8b0ce6376a0c1af471655540d2e9a81647d7b09ed823018426576d.
//
// Solidity: event StrategyWithdrawalDelayBlocksSet(address strategy, uint256 previousValue, uint256 newValue)
func (_DelegationManager *DelegationManagerFilterer) ParseStrategyWithdrawalDelayBlocksSet(log types.Log) (*DelegationManagerStrategyWithdrawalDelayBlocksSet, error) {
	event := new(DelegationManagerStrategyWithdrawalDelayBlocksSet)
	if err := _DelegationManager.contract.UnpackLog(event, "StrategyWithdrawalDelayBlocksSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the DelegationManager contract.
type DelegationManagerUnpausedIterator struct {
	Event *DelegationManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *DelegationManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerUnpaused)
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
		it.Event = new(DelegationManagerUnpaused)
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
func (it *DelegationManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerUnpaused represents a Unpaused event raised by the DelegationManager contract.
type DelegationManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_DelegationManager *DelegationManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*DelegationManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerUnpausedIterator{contract: _DelegationManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_DelegationManager *DelegationManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *DelegationManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerUnpaused)
				if err := _DelegationManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_DelegationManager *DelegationManagerFilterer) ParseUnpaused(log types.Log) (*DelegationManagerUnpaused, error) {
	event := new(DelegationManagerUnpaused)
	if err := _DelegationManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerWithdrawalCompletedIterator is returned from FilterWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for WithdrawalCompleted events raised by the DelegationManager contract.
type DelegationManagerWithdrawalCompletedIterator struct {
	Event *DelegationManagerWithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *DelegationManagerWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerWithdrawalCompleted)
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
		it.Event = new(DelegationManagerWithdrawalCompleted)
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
func (it *DelegationManagerWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerWithdrawalCompleted represents a WithdrawalCompleted event raised by the DelegationManager contract.
type DelegationManagerWithdrawalCompleted struct {
	WithdrawalRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCompleted is a free log retrieval operation binding the contract event 0xc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d.
//
// Solidity: event WithdrawalCompleted(bytes32 withdrawalRoot)
func (_DelegationManager *DelegationManagerFilterer) FilterWithdrawalCompleted(opts *bind.FilterOpts) (*DelegationManagerWithdrawalCompletedIterator, error) {

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return &DelegationManagerWithdrawalCompletedIterator{contract: _DelegationManager.contract, event: "WithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCompleted is a free log subscription operation binding the contract event 0xc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d.
//
// Solidity: event WithdrawalCompleted(bytes32 withdrawalRoot)
func (_DelegationManager *DelegationManagerFilterer) WatchWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *DelegationManagerWithdrawalCompleted) (event.Subscription, error) {

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerWithdrawalCompleted)
				if err := _DelegationManager.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
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

// ParseWithdrawalCompleted is a log parse operation binding the contract event 0xc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d.
//
// Solidity: event WithdrawalCompleted(bytes32 withdrawalRoot)
func (_DelegationManager *DelegationManagerFilterer) ParseWithdrawalCompleted(log types.Log) (*DelegationManagerWithdrawalCompleted, error) {
	event := new(DelegationManagerWithdrawalCompleted)
	if err := _DelegationManager.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerWithdrawalQueuedIterator is returned from FilterWithdrawalQueued and is used to iterate over the raw logs and unpacked data for WithdrawalQueued events raised by the DelegationManager contract.
type DelegationManagerWithdrawalQueuedIterator struct {
	Event *DelegationManagerWithdrawalQueued // Event containing the contract specifics and raw log

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
func (it *DelegationManagerWithdrawalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerWithdrawalQueued)
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
		it.Event = new(DelegationManagerWithdrawalQueued)
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
func (it *DelegationManagerWithdrawalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerWithdrawalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerWithdrawalQueued represents a WithdrawalQueued event raised by the DelegationManager contract.
type DelegationManagerWithdrawalQueued struct {
	WithdrawalRoot [32]byte
	Withdrawal     IDelegationManagerWithdrawal
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalQueued is a free log retrieval operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_DelegationManager *DelegationManagerFilterer) FilterWithdrawalQueued(opts *bind.FilterOpts) (*DelegationManagerWithdrawalQueuedIterator, error) {

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return &DelegationManagerWithdrawalQueuedIterator{contract: _DelegationManager.contract, event: "WithdrawalQueued", logs: logs, sub: sub}, nil
}

// WatchWithdrawalQueued is a free log subscription operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_DelegationManager *DelegationManagerFilterer) WatchWithdrawalQueued(opts *bind.WatchOpts, sink chan<- *DelegationManagerWithdrawalQueued) (event.Subscription, error) {

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "WithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerWithdrawalQueued)
				if err := _DelegationManager.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
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

// ParseWithdrawalQueued is a log parse operation binding the contract event 0x9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9.
//
// Solidity: event WithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_DelegationManager *DelegationManagerFilterer) ParseWithdrawalQueued(log types.Log) (*DelegationManagerWithdrawalQueued, error) {
	event := new(DelegationManagerWithdrawalQueued)
	if err := _DelegationManager.contract.UnpackLog(event, "WithdrawalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
