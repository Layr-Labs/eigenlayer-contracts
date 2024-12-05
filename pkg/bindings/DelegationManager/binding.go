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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_slasher\",\"type\":\"address\",\"internalType\":\"contractISlasher\"},{\"name\":\"_eigenPodManager\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DELEGATION_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_STAKER_OPT_OUT_WINDOW_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_WITHDRAWAL_DELAY_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKER_DELEGATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateCurrentStakerDelegationDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateDelegationApprovalDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateStakerDelegationDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stakerNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateWithdrawalRoot\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawal\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"middlewareTimesIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawals\",\"inputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManager.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"middlewareTimesIndexes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cumulativeWithdrawalsQueued\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateTo\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateToBySignature\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegatedTo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApprover\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApproverSaltIsSpent\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelegatableShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawalDelay\",\"inputs\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minWithdrawalDelayBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"_withdrawalDelayBlocks\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minWithdrawalDelayBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyOperatorDetails\",\"inputs\":[{\"name\":\"newOperatorDetails\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"__deprecated_earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorDetails\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"__deprecated_earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingWithdrawals\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueWithdrawals\",\"inputs\":[{\"name\":\"queuedWithdrawalParams\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManager.QueuedWithdrawalParams[]\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerAsOperator\",\"inputs\":[{\"name\":\"registeringOperatorDetails\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"__deprecated_earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinWithdrawalDelayBlocks\",\"inputs\":[{\"name\":\"newMinWithdrawalDelayBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyWithdrawalDelayBlocks\",\"inputs\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"withdrawalDelayBlocks\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerNonce\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerOptOutWindowBlocks\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyWithdrawalDelayBlocks\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinWithdrawalDelayBlocksSet\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDetailsModified\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOperatorDetails\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"__deprecated_earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMetadataURIUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorDetails\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"__deprecated_earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesDecreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesIncreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerForceUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyWithdrawalDelayBlocksSet\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalCompleted\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalQueued\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"withdrawal\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationManager.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyDelegated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignatureEIP1271\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignatureEOA\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCurrentlyDelegated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorsCannotUndelegate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureSaltSpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakerOptOutWindowBlocksCannotDecrease\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakerOptOutWindowBlocksExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalDelayExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalDelayNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawerNotStaker\",\"inputs\":[]}]",
	Bin: "0x61010060405234801561001157600080fd5b50604051614b63380380614b6383398101604081905261003091610134565b6001600160a01b0380841660805280821660c052821660a05261005161005d565b50504660e05250610181565b600054610100900460ff16156100c95760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161461011a576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461013157600080fd5b50565b60008060006060848603121561014957600080fd5b83516101548161011c565b60208501519093506101658161011c565b60408501519092506101768161011c565b809150509250925092565b60805160a05160c05160e05161494e61021560003960006120460152600081816105b101528181610f04015281816112320152818161185b0152818161227301528181613156015261356c015260006107620152600081816104f901528181610ed201528181611200015281816118ef01528181612340015281816124140152818161327c0152613612015261494e6000f3fe608060405234801561001057600080fd5b50600436106103425760003560e01c8063635bbd10116101b8578063b7f06ebe11610104578063cf80873e116100a2578063f16172b01161007c578063f16172b014610908578063f2fde38b1461091b578063f698da251461092e578063fabc1cbc1461093657600080fd5b8063cf80873e146108c1578063da8be864146108e2578063eea9064b146108f557600080fd5b8063c488375a116100de578063c488375a146107de578063c5e480db146107fe578063c94b5111146108a4578063ca661c04146108b757600080fd5b8063b7f06ebe14610784578063bb45fef2146107a7578063c448feb8146107d557600080fd5b8063886f1195116101715780639104c3191161014b5780639104c3191461070f57806399be81c81461072a578063a17884841461073d578063b13442711461075d57600080fd5b8063886f1195146106cb5780638da5cb5b146106de57806390041347146106ef57600080fd5b8063635bbd101461063657806365da1264146106495780636d70f7ae14610672578063715018a614610685578063778e55f31461068d5780637f548071146106b857600080fd5b806328a573ae116102925780634665bcda11610230578063597b36da1161020a578063597b36da146105e55780635ac86ab7146105f85780635c975abb1461061b57806360d7faed1461062357600080fd5b80634665bcda146105ac5780634fc40b61146105d3578063595c6a67146105dd57600080fd5b806339b70e381161026c57806339b70e38146104f45780633cdeb5e0146105335780633e28391d14610562578063433773821461058557600080fd5b806328a573ae146104ae57806329c77d4f146104c157806333404396146104e157600080fd5b8063132d4967116102ff57806316928365116102d957806316928365146104285780631bbce0911461046157806320606b701461047457806322bf40e41461049b57600080fd5b8063132d4967146103ef578063136439dd146104025780631522bf021461041557600080fd5b80630449ca391461034757806304a4f9791461036d5780630b9f487a146103945780630dd8dd02146103a75780630f589e59146103c757806310d67a2f146103dc575b600080fd5b61035a610355366004613924565b610949565b6040519081526020015b60405180910390f35b61035a7f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad81565b61035a6103a236600461398a565b6109c8565b6103ba6103b5366004613924565b610a8a565b60405161036491906139e5565b6103da6103d5366004613a76565b610d36565b005b6103da6103ea366004613ac9565b610e13565b6103da6103fd366004613aed565b610ec7565b6103da610410366004613b2e565b610f7f565b6103da610423366004613b47565b61106a565b61035a610436366004613ac9565b6001600160a01b0316600090815260996020526040902060010154600160a01b900463ffffffff1690565b61035a61046f366004613aed565b61107e565b61035a7f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a86681565b6103da6104a9366004613bb6565b6110ac565b6103da6104bc366004613aed565b6111f5565b61035a6104cf366004613ac9565b609b6020526000908152604090205481565b6103da6104ef366004613c61565b6112a6565b61051b7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610364565b61051b610541366004613ac9565b6001600160a01b039081166000908152609960205260409020600101541690565b610575610570366004613ac9565b611385565b6040519015158152602001610364565b61035a7f39111bc4a4d688e1f685123d7497d4615370152a8ee4a0593e647bd06ad8bb0b81565b61051b7f000000000000000000000000000000000000000000000000000000000000000081565b61035a6213c68081565b6103da6113a5565b61035a6105f3366004613f78565b61146d565b610575610606366004613fb4565b606654600160ff9092169190911b9081161490565b60665461035a565b6103da610631366004613fe5565b61149d565b6103da610644366004613b2e565b6114ed565b61051b610657366004613ac9565b609a602052600090815260409020546001600160a01b031681565b610575610680366004613ac9565b6114fe565b6103da611535565b61035a61069b366004614075565b609860209081526000928352604080842090915290825290205481565b6103da6106c6366004614159565b611549565b60655461051b906001600160a01b031681565b6033546001600160a01b031661051b565b6107026106fd3660046141ed565b61163c565b6040516103649190614278565b61051b73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b6103da61073836600461428b565b61170e565b61035a61074b366004613ac9565b609f6020526000908152604090205481565b61051b7f000000000000000000000000000000000000000000000000000000000000000081565b610575610792366004613b2e565b609e6020526000908152604090205460ff1681565b6105756107b53660046142c0565b609c60209081526000928352604080842090915290825290205460ff1681565b61035a609d5481565b61035a6107ec366004613ac9565b60a16020526000908152604090205481565b61086e61080c366004613ac9565b6040805160608082018352600080835260208084018290529284018190526001600160a01b03948516815260998352839020835191820184528054851682526001015493841691810191909152600160a01b90920463ffffffff169082015290565b6040805182516001600160a01b039081168252602080850151909116908201529181015163ffffffff1690820152606001610364565b61035a6108b23660046142ec565b61177b565b61035a62034bc081565b6108d46108cf366004613ac9565b611834565b60405161036492919061436f565b6103ba6108f0366004613ac9565b611beb565b6103da610903366004614394565b611f43565b6103da6109163660046143ed565b611f9c565b6103da610929366004613ac9565b611fcc565b61035a612042565b6103da610944366004613b2e565b61207f565b609d54600090815b838110156109be57600060a1600087878581811061097157610971614409565b90506020020160208101906109869190613ac9565b6001600160a01b03166001600160a01b03168152602001908152602001600020549050828111156109b5578092505b50600101610951565b5090505b92915050565b604080517f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad6020808301919091526001600160a01b038681168385015288811660608401528716608083015260a0820185905260c08083018590528351808403909101815260e0909201909252805191012060009081610a46612042565b60405161190160f01b602082015260228101919091526042810183905260620160408051808303601f19018152919052805160209091012098975050505050505050565b606654606090600190600290811603610ab65760405163840a48d560e01b815260040160405180910390fd5b6000836001600160401b03811115610ad057610ad0613d14565b604051908082528060200260200182016040528015610af9578160200160208202803683370190505b50336000908152609a60205260408120549192506001600160a01b03909116905b85811015610d2b57868682818110610b3457610b34614409565b9050602002810190610b46919061441f565b610b5490602081019061443f565b9050878783818110610b6857610b68614409565b9050602002810190610b7a919061441f565b610b84908061443f565b905014610ba4576040516343714afd60e01b815260040160405180910390fd5b33878783818110610bb757610bb7614409565b9050602002810190610bc9919061441f565b610bda906060810190604001613ac9565b6001600160a01b031614610c01576040516330c4716960e21b815260040160405180910390fd5b610d063383898985818110610c1857610c18614409565b9050602002810190610c2a919061441f565b610c3b906060810190604001613ac9565b8a8a86818110610c4d57610c4d614409565b9050602002810190610c5f919061441f565b610c69908061443f565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508e92508d9150889050818110610caf57610caf614409565b9050602002810190610cc1919061441f565b610ccf90602081019061443f565b8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061218792505050565b838281518110610d1857610d18614409565b6020908102919091010152600101610b1a565b509095945050505050565b610d3f33611385565b15610d5d57604051631a30c08d60e31b815260040160405180910390fd5b610d6733846125c1565b604080518082019091526060815260006020820152610d8933808360006126b3565b336001600160a01b03167f8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e285604051610dc29190614488565b60405180910390a2336001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b67080908484604051610e059291906144df565b60405180910390a250505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8a919061450e565b6001600160a01b0316336001600160a01b031614610ebb5760405163794821ff60e01b815260040160405180910390fd5b610ec48161289f565b50565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610f265750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b610f4357604051635c427cd960e01b815260040160405180910390fd5b610f4c83611385565b15610f7a576001600160a01b038084166000908152609a602052604090205416610f788185858561292f565b505b505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610fc7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610feb919061452b565b61100857604051631d77d47760e21b815260040160405180910390fd5b6066548181161461102c5760405163c61dca5d60e01b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b6110726129aa565b610f7884848484612a04565b6001600160a01b0383166000908152609b60205260408120546110a38582868661177b565b95945050505050565b600054610100900460ff16158080156110cc5750600054600160ff909116105b806110e65750303b1580156110e6575060005460ff166001145b61114e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015611171576000805461ff0019166101001790555b61117b8888612b1d565b611183612ba2565b60975561118f89612c39565b61119886612c8b565b6111a485858585612a04565b80156111ea576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806112545750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b61127157604051635c427cd960e01b815260040160405180910390fd5b61127a83611385565b15610f7a576001600160a01b038084166000908152609a602052604090205416610f7881858585612cf0565b6066546002906004908116036112cf5760405163840a48d560e01b815260040160405180910390fd5b6112d7612d6b565b60005b8881101561137a576113728a8a838181106112f7576112f7614409565b90506020028101906113099190614548565b89898481811061131b5761131b614409565b905060200281019061132d919061443f565b89898681811061133f5761133f614409565b9050602002013588888781811061135857611358614409565b905060200201602081019061136d919061455e565b612dc4565b6001016112da565b506111ea600160c955565b6001600160a01b039081166000908152609a602052604090205416151590565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156113ed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611411919061452b565b61142e57604051631d77d47760e21b815260040160405180910390fd5b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60008160405160200161148091906145fa565b604051602081830303815290604052805190602001209050919050565b6066546002906004908116036114c65760405163840a48d560e01b815260040160405180910390fd5b6114ce612d6b565b6114db8686868686612dc4565b6114e5600160c955565b505050505050565b6114f56129aa565b610ec481612c8b565b60006001600160a01b038216158015906109c25750506001600160a01b039081166000818152609a60205260409020549091161490565b61153d6129aa565b6115476000612c39565b565b428360200151101561156e57604051630819bdcd60e01b815260040160405180910390fd5b61157785611385565b1561159557604051631a30c08d60e31b815260040160405180910390fd5b61159e846114fe565b6115bb576040516325878fa360e21b815260040160405180910390fd5b6000609b6000876001600160a01b03166001600160a01b0316815260200190815260200160002054905060006115f7878388886020015161177b565b6001600160a01b0388166000908152609b602052604090206001840190558551909150611627908890839061344d565b611633878786866126b3565b50505050505050565b6060600082516001600160401b0381111561165957611659613d14565b604051908082528060200260200182016040528015611682578160200160208202803683370190505b50905060005b83518110156109be576001600160a01b038516600090815260986020526040812085519091908690849081106116c0576116c0614409565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020548282815181106116fb576116fb614409565b6020908102919091010152600101611688565b611717336114fe565b611734576040516325878fa360e21b815260040160405180910390fd5b336001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090838360405161176f9291906144df565b60405180910390a25050565b604080517f39111bc4a4d688e1f685123d7497d4615370152a8ee4a0593e647bd06ad8bb0b6020808301919091526001600160a01b0387811683850152851660608301526080820186905260a08083018590528351808403909101815260c09092019092528051910120600090816117f1612042565b60405161190160f01b602082015260228101919091526042810183905260620160408051808303601f190181529190528051602090910120979650505050505050565b6040516360f4062b60e01b81526001600160a01b03828116600483015260609182916000917f0000000000000000000000000000000000000000000000000000000000000000909116906360f4062b90602401602060405180830381865afa1580156118a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118c8919061460d565b6040516394f649dd60e01b81526001600160a01b03868116600483015291925060009182917f0000000000000000000000000000000000000000000000000000000000000000909116906394f649dd90602401600060405180830381865afa158015611938573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526119609190810190614684565b915091506000831361197757909590945092505050565b6060808351600003611a30576040805160018082528183019092529060208083019080368337505060408051600180825281830190925292945090506020808301908036833701905050905073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0826000815181106119eb576119eb614409565b60200260200101906001600160a01b031690816001600160a01b0316815250508481600081518110611a1f57611a1f614409565b602002602001018181525050611bde565b8351611a3d90600161475b565b6001600160401b03811115611a5457611a54613d14565b604051908082528060200260200182016040528015611a7d578160200160208202803683370190505b50915081516001600160401b03811115611a9957611a99613d14565b604051908082528060200260200182016040528015611ac2578160200160208202803683370190505b50905060005b8451811015611b5c57848181518110611ae357611ae3614409565b6020026020010151838281518110611afd57611afd614409565b60200260200101906001600160a01b031690816001600160a01b031681525050838181518110611b2f57611b2f614409565b6020026020010151828281518110611b4957611b49614409565b6020908102919091010152600101611ac8565b5073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac08260018451611b81919061476e565b81518110611b9157611b91614409565b60200260200101906001600160a01b031690816001600160a01b031681525050848160018451611bc1919061476e565b81518110611bd157611bd1614409565b6020026020010181815250505b9097909650945050505050565b606654606090600190600290811603611c175760405163840a48d560e01b815260040160405180910390fd5b611c2083611385565b611c3d57604051630b3d63f960e11b815260040160405180910390fd5b611c46836114fe565b15611c64576040516311ca333560e31b815260040160405180910390fd5b6001600160a01b038316611c8b576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b038084166000818152609a602052604090205490911690331480611cbe5750336001600160a01b038216145b80611ce557506001600160a01b038181166000908152609960205260409020600101541633145b611d0257604051635c427cd960e01b815260040160405180910390fd5b600080611d0e86611834565b9092509050336001600160a01b03871614611d6457826001600160a01b0316866001600160a01b03167ff0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a60405160405180910390a35b826001600160a01b0316866001600160a01b03167ffee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af4467660405160405180910390a36001600160a01b0386166000908152609a6020526040812080546001600160a01b031916905582519003611de8576040805160008152602081019091529450611f3a565b81516001600160401b03811115611e0157611e01613d14565b604051908082528060200260200182016040528015611e2a578160200160208202803683370190505b50945060005b8251811015611f3857604080516001808252818301909252600091602080830190803683375050604080516001808252818301909252929350600092915060208083019080368337019050509050848381518110611e9057611e90614409565b602002602001015182600081518110611eab57611eab614409565b60200260200101906001600160a01b031690816001600160a01b031681525050838381518110611edd57611edd614409565b602002602001015181600081518110611ef857611ef8614409565b602002602001018181525050611f1189878b8585612187565b888481518110611f2357611f23614409565b60209081029190910101525050600101611e30565b505b50505050919050565b611f4c33611385565b15611f6a57604051631a30c08d60e31b815260040160405180910390fd5b611f73836114fe565b611f90576040516325878fa360e21b815260040160405180910390fd5b610f7a338484846126b3565b611fa5336114fe565b611fc2576040516325878fa360e21b815260040160405180910390fd5b610ec433826125c1565b611fd46129aa565b6001600160a01b0381166120395760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401611145565b610ec481612c39565b60007f00000000000000000000000000000000000000000000000000000000000000004603612072575060975490565b61207a612ba2565b905090565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156120d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120f6919061450e565b6001600160a01b0316336001600160a01b0316146121275760405163794821ff60e01b815260040160405180910390fd5b6066541981196066541916146121505760405163c61dca5d60e01b815260040160405180910390fd5b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200161105f565b60006001600160a01b0386166121b0576040516339b190bb60e11b815260040160405180910390fd5b82516000036121d25760405163796cc52560e01b815260040160405180910390fd5b60005b83518110156124cf576001600160a01b0386161561222b5761222b868886848151811061220457612204614409565b602002602001015186858151811061221e5761221e614409565b602002602001015161292f565b73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac06001600160a01b031684828151811061225b5761225b614409565b60200260200101516001600160a01b031603612323577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663beffbb89888584815181106122b3576122b3614409565b60200260200101516040518363ffffffff1660e01b81526004016122ec9291906001600160a01b03929092168252602082015260400190565b600060405180830381600087803b15801561230657600080fd5b505af115801561231a573d6000803e3d6000fd5b505050506124c7565b846001600160a01b0316876001600160a01b031614806123f557507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639b4da03d85838151811061237f5761237f614409565b60200260200101516040518263ffffffff1660e01b81526004016123b291906001600160a01b0391909116815260200190565b602060405180830381865afa1580156123cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123f3919061452b565b155b612412576040516330c4716960e21b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638c80d4e58886848151811061245457612454614409565b602002602001015186858151811061246e5761246e614409565b60200260200101516040518463ffffffff1660e01b815260040161249493929190614781565b600060405180830381600087803b1580156124ae57600080fd5b505af11580156124c2573d6000803e3d6000fd5b505050505b6001016121d5565b506001600160a01b0386166000908152609f602052604081208054918291906124f7836147a5565b919050555060006040518060e00160405280896001600160a01b03168152602001886001600160a01b03168152602001876001600160a01b031681526020018381526020014363ffffffff168152602001868152602001858152509050600061255f8261146d565b6000818152609e602052604090819020805460ff19166001179055519091507f9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f9906125ad90839085906147be565b60405180910390a198975050505050505050565b6213c6806125d560608301604084016147d7565b63ffffffff1611156125fa57604051638b11ba2f60e01b815260040160405180910390fd5b6001600160a01b0382166000908152609960205260409081902060010154600160a01b900463ffffffff169061263690606084019084016147d7565b63ffffffff16101561265b5760405163102b92a760e01b815260040160405180910390fd5b6001600160a01b0382166000908152609960205260409020819061267f8282614814565b505060405133907ffebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac9061176f908490614488565b6066546000906001908116036126dc5760405163840a48d560e01b815260040160405180910390fd5b6001600160a01b038085166000908152609960205260409020600101541680158015906127125750336001600160a01b03821614155b80156127275750336001600160a01b03861614155b156127ea57428460200151101561275157604051630819bdcd60e01b815260040160405180910390fd5b6001600160a01b0381166000908152609c6020908152604080832086845290915290205460ff161561279657604051639a5afeb760e01b815260040160405180910390fd5b6001600160a01b0381166000908152609c6020908152604080832086845282528220805460ff191660011790558501516127d79088908890859088906109c8565b90506127e88282876000015161344d565b505b6001600160a01b038681166000818152609a602052604080822080546001600160a01b031916948a169485179055517fc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d87433049190a360008061284988611834565b9150915060005b82518110156111ea57612897888a85848151811061287057612870614409565b602002602001015185858151811061288a5761288a614409565b6020026020010151612cf0565b600101612850565b6001600160a01b0381166128c6576040516339b190bb60e11b815260040160405180910390fd5b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0380851660009081526098602090815260408083209386168352929052908120805483929061296690849061476e565b92505081905550836001600160a01b03167f6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd848484604051610e0593929190614781565b6033546001600160a01b031633146115475760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401611145565b828114612a24576040516343714afd60e01b815260040160405180910390fd5b8260005b818110156114e5576000868683818110612a4457612a44614409565b9050602002016020810190612a599190613ac9565b6001600160a01b038116600090815260a16020526040812054919250868685818110612a8757612a87614409565b90506020020135905062034bc0811115612ab457604051639373e7cf60e01b815260040160405180910390fd5b6001600160a01b038316600081815260a160209081526040918290208490558151928352820184905281018290527f0e7efa738e8b0ce6376a0c1af471655540d2e9a81647d7b09ed823018426576d9060600160405180910390a1505050806001019050612a28565b6065546001600160a01b0316158015612b3e57506001600160a01b03821615155b612b5b576040516339b190bb60e11b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2612b9e8261289f565b5050565b604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b62034bc0811115612caf57604051639373e7cf60e01b815260040160405180910390fd5b609d5460408051918252602082018390527fafa003cd76f87ff9d62b35beea889920f33c0c42b8d45b74954d61d50f4b6b69910160405180910390a1609d55565b6001600160a01b03808516600090815260986020908152604080832093861683529290529081208054839290612d2790849061475b565b92505081905550836001600160a01b03167f1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c848484604051610e0593929190614781565b600260c95403612dbd5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611145565b600260c955565b6000612dd26105f387614877565b6000818152609e602052604090205490915060ff16612e045760405163295a66e160e11b815260040160405180910390fd5b609d544390612e1960a0890160808a016147d7565b63ffffffff16612e29919061475b565b1115612e48576040516378f67ae160e11b815260040160405180910390fd5b612e586060870160408801613ac9565b6001600160a01b0316336001600160a01b031614612e8957604051635c427cd960e01b815260040160405180910390fd5b8115612ebd57612e9c60a087018761443f565b85149050612ebd576040516343714afd60e01b815260040160405180910390fd5b6000818152609e60205260409020805460ff1916905581156130235760005b612ee960a088018861443f565b905081101561301d574360a16000612f0460a08b018b61443f565b85818110612f1457612f14614409565b9050602002016020810190612f299190613ac9565b6001600160a01b03168152602081019190915260400160002054612f5360a08a0160808b016147d7565b63ffffffff16612f63919061475b565b1115612f82576040516378f67ae160e11b815260040160405180910390fd5b613015612f926020890189613ac9565b33612fa060a08b018b61443f565b85818110612fb057612fb0614409565b9050602002016020810190612fc59190613ac9565b612fd260c08c018c61443f565b86818110612fe257612fe2614409565b905060200201358a8a87818110612ffb57612ffb614409565b90506020020160208101906130109190613ac9565b613531565b600101612edc565b50613412565b336000908152609a60205260408120546001600160a01b0316905b61304b60a089018961443f565b905081101561340f574360a1600061306660a08c018c61443f565b8581811061307657613076614409565b905060200201602081019061308b9190613ac9565b6001600160a01b031681526020810191909152604001600020546130b560a08b0160808c016147d7565b63ffffffff166130c5919061475b565b11156130e4576040516378f67ae160e11b815260040160405180910390fd5b73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac061310660a08a018a61443f565b8381811061311657613116614409565b905060200201602081019061312b9190613ac9565b6001600160a01b03160361327a57600061314860208a018a613ac9565b905060006001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016630e81073c8361318960c08e018e61443f565b8781811061319957613199614409565b6040516001600160e01b031960e087901b1681526001600160a01b03909416600485015260200291909101356024830152506044016020604051808303816000875af11580156131ed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613211919061460d565b6001600160a01b038084166000908152609a602052604090205491925016801561327257613272818461324760a08f018f61443f565b8881811061325757613257614409565b905060200201602081019061326c9190613ac9565b85612cf0565b505050613407565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c4623ea1338989858181106132bc576132bc614409565b90506020020160208101906132d19190613ac9565b6132de60a08d018d61443f565b868181106132ee576132ee614409565b90506020020160208101906133039190613ac9565b61331060c08e018e61443f565b8781811061332057613320614409565b60405160e088901b6001600160e01b03191681526001600160a01b03968716600482015294861660248601529290941660448401526020909102013560648201526084019050600060405180830381600087803b15801561338057600080fd5b505af1158015613394573d6000803e3d6000fd5b505050506001600160a01b038216156134075761340782336133b960a08c018c61443f565b858181106133c9576133c9614409565b90506020020160208101906133de9190613ac9565b6133eb60c08d018d61443f565b868181106133fb576133fb614409565b90506020020135612cf0565b60010161303e565b50505b6040518181527fc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d9060200160405180910390a1505050505050565b6001600160a01b0383163b156134f657604051630b135d3f60e11b808252906001600160a01b03851690631626ba7e9061348d9086908690600401614883565b602060405180830381865afa1580156134aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134ce91906148d8565b6001600160e01b03191614610f7a5760405163026aa43360e31b815260040160405180910390fd5b826001600160a01b031661350a8383613671565b6001600160a01b031614610f7a576040516306c653c360e31b815260040160405180910390fd5b73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeabf196001600160a01b038416016135dc5760405162387b1360e81b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063387b1300906135a590889088908790600401614781565b600060405180830381600087803b1580156135bf57600080fd5b505af11580156135d3573d6000803e3d6000fd5b5050505061366a565b60405163c608c7f360e01b81526001600160a01b03858116600483015284811660248301526044820184905282811660648301527f0000000000000000000000000000000000000000000000000000000000000000169063c608c7f390608401600060405180830381600087803b15801561365657600080fd5b505af11580156111ea573d6000803e3d6000fd5b5050505050565b6000806000613680858561368d565b915091506109be816136d2565b60008082516041036136c35760208301516040840151606085015160001a6136b78782858561381c565b945094505050506136cb565b506000905060025b9250929050565b60008160048111156136e6576136e6614902565b036136ee5750565b600181600481111561370257613702614902565b0361374f5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401611145565b600281600481111561376357613763614902565b036137b05760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401611145565b60038160048111156137c4576137c4614902565b03610ec45760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401611145565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561385357506000905060036138d7565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156138a7573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166138d0576000600192509250506138d7565b9150600090505b94509492505050565b60008083601f8401126138f257600080fd5b5081356001600160401b0381111561390957600080fd5b6020830191508360208260051b85010111156136cb57600080fd5b6000806020838503121561393757600080fd5b82356001600160401b0381111561394d57600080fd5b613959858286016138e0565b90969095509350505050565b6001600160a01b0381168114610ec457600080fd5b803561398581613965565b919050565b600080600080600060a086880312156139a257600080fd5b85356139ad81613965565b945060208601356139bd81613965565b935060408601356139cd81613965565b94979396509394606081013594506080013592915050565b602080825282518282018190526000918401906040840190835b81811015610d2b5783518352602093840193909201916001016139ff565b600060608284031215613a2f57600080fd5b50919050565b60008083601f840112613a4757600080fd5b5081356001600160401b03811115613a5e57600080fd5b6020830191508360208285010111156136cb57600080fd5b600080600060808486031215613a8b57600080fd5b613a958585613a1d565b925060608401356001600160401b03811115613ab057600080fd5b613abc86828701613a35565b9497909650939450505050565b600060208284031215613adb57600080fd5b8135613ae681613965565b9392505050565b600080600060608486031215613b0257600080fd5b8335613b0d81613965565b92506020840135613b1d81613965565b929592945050506040919091013590565b600060208284031215613b4057600080fd5b5035919050565b60008060008060408587031215613b5d57600080fd5b84356001600160401b03811115613b7357600080fd5b613b7f878288016138e0565b90955093505060208501356001600160401b03811115613b9e57600080fd5b613baa878288016138e0565b95989497509550505050565b60008060008060008060008060c0898b031215613bd257600080fd5b8835613bdd81613965565b97506020890135613bed81613965565b9650604089013595506060890135945060808901356001600160401b03811115613c1657600080fd5b613c228b828c016138e0565b90955093505060a08901356001600160401b03811115613c4157600080fd5b613c4d8b828c016138e0565b999c989b5096995094979396929594505050565b6000806000806000806000806080898b031215613c7d57600080fd5b88356001600160401b03811115613c9357600080fd5b613c9f8b828c016138e0565b90995097505060208901356001600160401b03811115613cbe57600080fd5b613cca8b828c016138e0565b90975095505060408901356001600160401b03811115613ce957600080fd5b613cf58b828c016138e0565b90955093505060608901356001600160401b03811115613c4157600080fd5b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b0381118282101715613d4c57613d4c613d14565b60405290565b604080519081016001600160401b0381118282101715613d4c57613d4c613d14565b604051601f8201601f191681016001600160401b0381118282101715613d9c57613d9c613d14565b604052919050565b63ffffffff81168114610ec457600080fd5b803561398581613da4565b60006001600160401b03821115613dda57613dda613d14565b5060051b60200190565b600082601f830112613df557600080fd5b8135613e08613e0382613dc1565b613d74565b8082825260208201915060208360051b860101925085831115613e2a57600080fd5b602085015b83811015613e50578035613e4281613965565b835260209283019201613e2f565b5095945050505050565b600082601f830112613e6b57600080fd5b8135613e79613e0382613dc1565b8082825260208201915060208360051b860101925085831115613e9b57600080fd5b602085015b83811015613e50578035835260209283019201613ea0565b600060e08284031215613eca57600080fd5b613ed2613d2a565b9050613edd8261397a565b8152613eeb6020830161397a565b6020820152613efc6040830161397a565b604082015260608281013590820152613f1760808301613db6565b608082015260a08201356001600160401b03811115613f3557600080fd5b613f4184828501613de4565b60a08301525060c08201356001600160401b03811115613f6057600080fd5b613f6c84828501613e5a565b60c08301525092915050565b600060208284031215613f8a57600080fd5b81356001600160401b03811115613fa057600080fd5b613fac84828501613eb8565b949350505050565b600060208284031215613fc657600080fd5b813560ff81168114613ae657600080fd5b8015158114610ec457600080fd5b600080600080600060808688031215613ffd57600080fd5b85356001600160401b0381111561401357600080fd5b860160e0818903121561402557600080fd5b945060208601356001600160401b0381111561404057600080fd5b61404c888289016138e0565b90955093505060408601359150606086013561406781613fd7565b809150509295509295909350565b6000806040838503121561408857600080fd5b823561409381613965565b915060208301356140a381613965565b809150509250929050565b6000604082840312156140c057600080fd5b6140c8613d52565b905081356001600160401b038111156140e057600080fd5b8201601f810184136140f157600080fd5b80356001600160401b0381111561410a5761410a613d14565b61411d601f8201601f1916602001613d74565b81815285602083850101111561413257600080fd5b81602084016020830137600060209282018301528352928301359282019290925292915050565b600080600080600060a0868803121561417157600080fd5b853561417c81613965565b9450602086013561418c81613965565b935060408601356001600160401b038111156141a757600080fd5b6141b3888289016140ae565b93505060608601356001600160401b038111156141cf57600080fd5b6141db888289016140ae565b95989497509295608001359392505050565b6000806040838503121561420057600080fd5b823561420b81613965565b915060208301356001600160401b0381111561422657600080fd5b61423285828601613de4565b9150509250929050565b600081518084526020840193506020830160005b8281101561426e578151865260209586019590910190600101614250565b5093949350505050565b602081526000613ae6602083018461423c565b6000806020838503121561429e57600080fd5b82356001600160401b038111156142b457600080fd5b61395985828601613a35565b600080604083850312156142d357600080fd5b82356142de81613965565b946020939093013593505050565b6000806000806080858703121561430257600080fd5b843561430d81613965565b935060208501359250604085013561432481613965565b9396929550929360600135925050565b600081518084526020840193506020830160005b8281101561426e5781516001600160a01b0316865260209586019590910190600101614348565b6040815260006143826040830185614334565b82810360208401526110a3818561423c565b6000806000606084860312156143a957600080fd5b83356143b481613965565b925060208401356001600160401b038111156143cf57600080fd5b6143db868287016140ae565b93969395505050506040919091013590565b6000606082840312156143ff57600080fd5b613ae68383613a1d565b634e487b7160e01b600052603260045260246000fd5b60008235605e1983360301811261443557600080fd5b9190910192915050565b6000808335601e1984360301811261445657600080fd5b8301803591506001600160401b0382111561447057600080fd5b6020019150600581901b36038213156136cb57600080fd5b60608101823561449781613965565b6001600160a01b0316825260208301356144b081613965565b6001600160a01b0316602083015260408301356144cc81613da4565b63ffffffff811660408401525092915050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60006020828403121561452057600080fd5b8151613ae681613965565b60006020828403121561453d57600080fd5b8151613ae681613fd7565b6000823560de1983360301811261443557600080fd5b60006020828403121561457057600080fd5b8135613ae681613fd7565b80516001600160a01b03908116835260208083015182169084015260408083015190911690830152606080820151908301526080808201516000916145c79085018263ffffffff169052565b5060a082015160e060a08501526145e160e0850182614334565b905060c083015184820360c08601526110a3828261423c565b602081526000613ae6602083018461457b565b60006020828403121561461f57600080fd5b5051919050565b600082601f83011261463757600080fd5b8151614645613e0382613dc1565b8082825260208201915060208360051b86010192508583111561466757600080fd5b602085015b83811015613e5057805183526020928301920161466c565b6000806040838503121561469757600080fd5b82516001600160401b038111156146ad57600080fd5b8301601f810185136146be57600080fd5b80516146cc613e0382613dc1565b8082825260208201915060208360051b8501019250878311156146ee57600080fd5b6020840193505b8284101561471957835161470881613965565b8252602093840193909101906146f5565b8095505050505060208301516001600160401b0381111561473957600080fd5b61423285828601614626565b634e487b7160e01b600052601160045260246000fd5b808201808211156109c2576109c2614745565b818103818111156109c2576109c2614745565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6000600182016147b7576147b7614745565b5060010190565b828152604060208201526000613fac604083018461457b565b6000602082840312156147e957600080fd5b8135613ae681613da4565b80546001600160a01b0319166001600160a01b0392909216919091179055565b813561481f81613965565b61482981836147f4565b5060018101602083013561483c81613965565b61484681836147f4565b50604083013561485581613da4565b815463ffffffff60a01b191660a09190911b63ffffffff60a01b161790555050565b60006109c23683613eb8565b828152604060208201526000825180604084015260005b818110156148b7576020818601810151606086840101520161489a565b506000606082850101526060601f19601f8301168401019150509392505050565b6000602082840312156148ea57600080fd5b81516001600160e01b031981168114613ae657600080fd5b634e487b7160e01b600052602160045260246000fdfea26469706673582212207eff95b58bd0423a50ea943d372bfc0aa51f5a496e204874795e3046ea0b53ee64736f6c634300081b0033",
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
