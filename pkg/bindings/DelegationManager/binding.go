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
	Staker         common.Address
	DelegatedTo    common.Address
	Withdrawer     common.Address
	Nonce          *big.Int
	StartTimestamp uint32
	Strategies     []common.Address
	ScaledShares   []*big.Int
}

// ISignatureUtilsSignatureWithExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithExpiry struct {
	Signature []byte
	Expiry    *big.Int
}

// DelegationManagerMetaData contains all meta data concerning the DelegationManager contract.
var DelegationManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_slasher\",\"type\":\"address\",\"internalType\":\"contractISlasher\"},{\"name\":\"_eigenPodManager\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"},{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_MIN_WITHDRAWAL_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DELEGATION_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LEGACY_MIN_WITHDRAWAL_DELAY_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LEGACY_WITHDRAWALS_TIMESTAMP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_WITHDRAWAL_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKER_DELEGATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateCurrentStakerDelegationDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateDelegationApprovalDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateStakerDelegationDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stakerNonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateWithdrawalRoot\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawal\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawals\",\"inputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManager.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cumulativeWithdrawalsQueued\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"removedShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateTo\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateToBySignature\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegatedTo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApprover\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApproverSaltIsSpent\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelegatableShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorScaledShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawableStakerShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"existingShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addedShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyOperatorDetails\",\"inputs\":[{\"name\":\"newOperatorDetails\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"__deprecated_earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorDetails\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"__deprecated_earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorScaledShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingWithdrawals\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueWithdrawals\",\"inputs\":[{\"name\":\"queuedWithdrawalParams\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManager.QueuedWithdrawalParams[]\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerAsOperator\",\"inputs\":[{\"name\":\"registeringOperatorDetails\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"__deprecated_earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerNonce\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerOptOutWindowBlocks\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerScalingFactors\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDetailsModified\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOperatorDetails\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"__deprecated_earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMetadataURIUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorDetails\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"__deprecated_earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesDecreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesIncreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerForceUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalCompleted\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalQueued\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"withdrawal\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationManager.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ActivelyDelegated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AllocationDelaySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerCannotUndelegate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignatureEIP1271\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignatureSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotActivelyDelegated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManagerOrEigenPodManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorsCannotUndelegate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaltSpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakerOptOutWindowBlocksCannotDecrease\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakerOptOutWindowBlocksExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalDelayExeedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalDelayNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalExeedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalNotQueued\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawerNotCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawerNotStaker\",\"inputs\":[]}]",
	Bin: "0x61016060405234801561001157600080fd5b5060405161595838038061595883398101604081905261003091610152565b6001600160a01b0380871660c0528085166101005280861660e05280841660a05282166101205263ffffffff811660805261006961007a565b50504661014052506101e292505050565b600054610100900460ff16156100e65760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015610138576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461014f57600080fd5b50565b60008060008060008060c0878903121561016b57600080fd5b86516101768161013a565b60208801519096506101878161013a565b60408801519095506101988161013a565b60608801519094506101a98161013a565b60808801519093506101ba8161013a565b60a088015190925063ffffffff811681146101d457600080fd5b809150509295509295509295565b60805160a05160c05160e0516101005161012051610140516156666102f2600039600061251501526000818161090401528181610c1a0152818161106a015281816113fa01528181611651015281816120f70152818161277f015281816132050152613ae201526000818161059301528181610fe10152818161137101528181611b21015281816125e201528181612b0c01528181612e0f01528181613cbf0152613e75015260006107c80152600081816104c801528181610faf0152818161133f01528181611bb50152818161268901528181612b9701528181612e7201528181613d520152613f20015260006106900152600081816105ba01528181613a780152613b2601526156666000f3fe608060405234801561001057600080fd5b506004361061035d5760003560e01c8063715018a6116101d3578063c0525f0211610104578063e4cc3f90116100a2578063f2fde38b1161007c578063f2fde38b1461099c578063f698da25146109af578063f878776e146109b7578063fabc1cbc146109ca57600080fd5b8063e4cc3f9014610963578063eea9064b14610976578063f16172b01461098957600080fd5b8063ca8aa7c7116100de578063ca8aa7c7146108ff578063cb00387b14610926578063cf80873e1461092f578063da8be8641461095057600080fd5b8063c0525f021461083b578063c5e480db14610846578063c94b5111146108ec57600080fd5b806399be81c811610171578063a17884841161014b578063a1788484146107a3578063b1344271146107c3578063b7f06ebe146107ea578063bb45fef21461080d57600080fd5b806399be81c8146107455780639fb4ee3214610758578063a16107811461078357600080fd5b8063886f1195116101ad578063886f1195146106f35780638da5cb5b146107065780639104c319146107175780639435bb431461073257600080fd5b8063715018a6146106c5578063778e55f3146106cd5780637f548071146106e057600080fd5b80633c651cf2116102ad578063595c6a671161024b5780635c975abb116102255780635c975abb1461065a57806365da1264146106625780636b3aa72e1461068b5780636d70f7ae146106b257600080fd5b8063595c6a671461061c578063597b36da146106245780635ac86ab71461063757600080fd5b8063433773821161028757806343377382146105675780634665bcda1461058e5780634a5f2b5d146105b557806351260e23146105f157600080fd5b80633c651cf2146105025780633cdeb5e0146105155780633e28391d1461054457600080fd5b8063136439dd1161031a5780631bbce091116102f45780631bbce0911461046957806320606b701461047c57806329c77d4f146104a357806339b70e38146104c357600080fd5b8063136439dd1461040a578063169283651461041d5780631794bb3c1461045657600080fd5b806304a4f979146103625780630b9f487a1461039c5780630dd8dd02146103af5780630f589e59146103cf57806310d67a2f146103e4578063132d4967146103f7575b600080fd5b6103897f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad81565b6040519081526020015b60405180910390f35b6103896103aa3660046144d3565b6109dd565b6103c26103bd366004614572565b610a9f565b60405161039391906145b3565b6103e26103dd366004614644565b610e13565b005b6103e26103f2366004614697565b610ef0565b6103e26104053660046146b4565b610fa4565b6103e26104183660046146f5565b6110ec565b61038961042b366004614697565b6001600160a01b0316600090815260996020526040902060010154600160a01b900463ffffffff1690565b6103e26104643660046146b4565b6111d7565b6103896104773660046146b4565b611306565b6103897f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a86681565b6103896104b1366004614697565b609b6020526000908152604090205481565b6104ea7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610393565b6103e261051036600461470e565b611334565b6104ea610523366004614697565b6001600160a01b039081166000908152609960205260409020600101541690565b610557610552366004614697565b6114bf565b6040519015158152602001610393565b6103897f39111bc4a4d688e1f685123d7497d4615370152a8ee4a0593e647bd06ad8bb0b81565b6104ea7f000000000000000000000000000000000000000000000000000000000000000081565b6105dc7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610393565b6103896105ff366004614754565b60a260209081526000928352604080842090915290825290205481565b6103e26114df565b6103896106323660046149f1565b6115a7565b610557610645366004614a25565b606654600160ff9092169190911b9081161490565b606654610389565b6104ea610670366004614697565b609a602052600090815260409020546001600160a01b031681565b6104ea7f000000000000000000000000000000000000000000000000000000000000000081565b6105576106c0366004614697565b6115d7565b6103e2611611565b6103896106db366004614754565b611625565b6103e26106ee366004614af3565b6116f7565b6065546104ea906001600160a01b031681565b6033546001600160a01b03166104ea565b6104ea73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b6103e2610740366004614b87565b6117e1565b6103e2610753366004614c2a565b6118fa565b610389610766366004614754565b609860209081526000928352604080842090915290825290205481565b610796610791366004614c5f565b611967565b6040516103939190614cea565b6103896107b1366004614697565b609f6020526000908152604090205481565b6104ea7f000000000000000000000000000000000000000000000000000000000000000081565b6105576107f83660046146f5565b609e6020526000908152604090205460ff1681565b61055761081b366004614cfd565b609c60209081526000928352604080842090915290825290205460ff1681565b6105dc636775749081565b6108b6610854366004614697565b6040805160608082018352600080835260208084018290529284018190526001600160a01b03948516815260998352839020835191820184528054851682526001015493841691810191909152600160a01b90920463ffffffff169082015290565b6040805182516001600160a01b039081168252602080850151909116908201529181015163ffffffff1690820152606001610393565b6103896108fa366004614d29565b611a41565b6104ea7f000000000000000000000000000000000000000000000000000000000000000081565b61038961c4e081565b61094261093d366004614697565b611afa565b604051610393929190614dac565b6103c261095e366004614697565b611eb1565b6103e2610971366004614ddf565b61237a565b6103e2610984366004614e63565b612412565b6103e2610997366004614ebc565b61246b565b6103e26109aa366004614697565b61249b565b610389612511565b6107966109c5366004614ed8565b61254e565b6103e26109d83660046146f5565b6128ca565b604080517f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad6020808301919091526001600160a01b038681168385015288811660608401528716608083015260a0820185905260c08083018590528351808403909101815260e0909201909252805191012060009081610a5b612511565b60405161190160f01b602082015260228101919091526042810183905260620160408051808303601f19018152919052805160209091012098975050505050505050565b606654606090600190600290811603610acb5760405163840a48d560e01b815260040160405180910390fd5b6000836001600160401b03811115610ae557610ae561478d565b604051908082528060200260200182016040528015610b0e578160200160208202803683370190505b50336000908152609a60205260408120549192506001600160a01b03909116905b85811015610e0857868682818110610b4957610b49614f1f565b9050602002810190610b5b9190614f35565b610b69906020810190614f55565b9050878783818110610b7d57610b7d614f1f565b9050602002810190610b8f9190614f35565b610b999080614f55565b905014610bb9576040516343714afd60e01b815260040160405180910390fd5b33878783818110610bcc57610bcc614f1f565b9050602002810190610bde9190614f35565b610bef906060810190604001614697565b6001600160a01b031614610c16576040516330c4716960e21b815260040160405180910390fd5b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166339a9a3ed848a8a86818110610c5a57610c5a614f1f565b9050602002810190610c6c9190614f35565b610c769080614f55565b6040518463ffffffff1660e01b8152600401610c9493929190614fde565b600060405180830381865afa158015610cb1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610cd9919081019061501a565b9050610de233848a8a86818110610cf257610cf2614f1f565b9050602002810190610d049190614f35565b610d15906060810190604001614697565b8b8b87818110610d2757610d27614f1f565b9050602002810190610d399190614f35565b610d439080614f55565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508f92508e9150899050818110610d8957610d89614f1f565b9050602002810190610d9b9190614f35565b610da9906020810190614f55565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508992506129d2915050565b848381518110610df457610df4614f1f565b602090810291909101015250600101610b2f565b509095945050505050565b610e1c336114bf565b15610e3a57604051633bf2b50360e11b815260040160405180910390fd5b610e44338461300f565b604080518082019091526060815260006020820152610e663380836000613067565b336001600160a01b03167f8e8485583a2310d41f7c82b9427d0bd49bad74bb9cff9d3402a29d8f9b28a0e285604051610e9f91906150ac565b60405180910390a2336001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b67080908484604051610ee2929190615103565b60405180910390a250505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f679190615132565b6001600160a01b0316336001600160a01b031614610f985760405163794821ff60e01b815260040160405180910390fd5b610fa1816133d0565b50565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806110035750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b6110205760405163045206a560e21b815260040160405180910390fd5b611029836114bf565b156110e7576001600160a01b038381166000908152609a6020526040808220549051635a3932f160e11b8152908316600482018190528584166024830152927f0000000000000000000000000000000000000000000000000000000000000000169063b47265e290604401602060405180830381865afa1580156110b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d5919061514f565b90506110e48286868685613460565b50505b505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611134573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611158919061516a565b61117557604051631d77d47760e21b815260040160405180910390fd5b606654818116146111995760405163c61dca5d60e01b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600054610100900460ff16158080156111f75750600054600160ff909116105b806112115750303b158015611211575060005460ff166001145b6112795760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff19166001179055801561129c576000805461ff0019166101001790555b6112a683836134fc565b6112ae613581565b6097556112ba84613618565b8015611300576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b6001600160a01b0383166000908152609b602052604081205461132b85828686611a41565b95945050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806113935750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b6113b05760405163045206a560e21b815260040160405180910390fd5b6113b9846114bf565b15611300576001600160a01b038481166000908152609a6020526040808220549051635a3932f160e11b8152908316600482018190528684166024830152927f0000000000000000000000000000000000000000000000000000000000000000169063b47265e290604401602060405180830381865afa158015611441573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611465919061514f565b9050600061147f87611477898961366a565b8488886136ab565b6001600160a01b03808916600090815260a260209081526040808320938b1683529290522081905590506114b68388888786613767565b50505050505050565b6001600160a01b039081166000908152609a602052604090205416151590565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015611527573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061154b919061516a565b61156857604051631d77d47760e21b815260040160405180910390fd5b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6000816040516020016115ba9190615206565b604051602081830303815290604052805190602001209050919050565b60006001600160a01b0382161580159061160b57506001600160a01b038083166000818152609a6020526040902054909116145b92915050565b6116196137f3565b6116236000613618565b565b604051635a3932f160e11b81526001600160a01b038381166004830152828116602483015260009182917f0000000000000000000000000000000000000000000000000000000000000000169063b47265e290604401602060405180830381865afa158015611698573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116bc919061514f565b6001600160a01b038086166000908152609860209081526040808320938816835292905220549091506116ef908261384d565b949350505050565b428360200151101561171c57604051630819bdcd60e01b815260040160405180910390fd5b611725856114bf565b1561174357604051633bf2b50360e11b815260040160405180910390fd5b61174c846115d7565b611769576040516325ec6c1f60e01b815260040160405180910390fd5b6000609b6000876001600160a01b03166001600160a01b0316815260200190815260200160002054905060006117a58783888860200151611a41565b6001600160a01b0388166000908152609b6020526040902060018401905585519091506117d59088908390613875565b6114b687878686613067565b60665460029060049081160361180a5760405163840a48d560e01b815260040160405180910390fd5b600260c9540361185c5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611270565b600260c95560005b868110156118eb576118e388888381811061188157611881614f1f565b90506020028101906118939190615219565b8787848181106118a5576118a5614f1f565b90506020028101906118b79190614f55565b8787868181106118c9576118c9614f1f565b90506020020160208101906118de919061522f565b613959565b600101611864565b5050600160c955505050505050565b611903336115d7565b611920576040516325ec6c1f60e01b815260040160405180910390fd5b336001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090838360405161195b929190615103565b60405180910390a25050565b6060600082516001600160401b038111156119845761198461478d565b6040519080825280602002602001820160405280156119ad578160200160208202803683370190505b50905060005b8351811015611a39576001600160a01b038516600090815260986020526040812085519091908690849081106119eb576119eb614f1f565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054828281518110611a2657611a26614f1f565b60209081029190910101526001016119b3565b509392505050565b604080517f39111bc4a4d688e1f685123d7497d4615370152a8ee4a0593e647bd06ad8bb0b6020808301919091526001600160a01b0387811683850152851660608301526080820186905260a08083018590528351808403909101815260c0909201909252805191012060009081611ab7612511565b60405161190160f01b602082015260228101919091526042810183905260620160408051808303601f190181529190528051602090910120979650505050505050565b6040516360f4062b60e01b81526001600160a01b03828116600483015260609182916000917f0000000000000000000000000000000000000000000000000000000000000000909116906360f4062b90602401602060405180830381865afa158015611b6a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b8e919061524c565b6040516394f649dd60e01b81526001600160a01b03868116600483015291925060009182917f0000000000000000000000000000000000000000000000000000000000000000909116906394f649dd90602401600060405180830381865afa158015611bfe573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611c2691908101906152c3565b9150915060008313611c3d57909590945092505050565b6060808351600003611cf6576040805160018082528183019092529060208083019080368337505060408051600180825281830190925292945090506020808301908036833701905050905073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac082600081518110611cb157611cb1614f1f565b60200260200101906001600160a01b031690816001600160a01b0316815250508481600081518110611ce557611ce5614f1f565b602002602001018181525050611ea4565b8351611d0390600161539a565b6001600160401b03811115611d1a57611d1a61478d565b604051908082528060200260200182016040528015611d43578160200160208202803683370190505b50915081516001600160401b03811115611d5f57611d5f61478d565b604051908082528060200260200182016040528015611d88578160200160208202803683370190505b50905060005b8451811015611e2257848181518110611da957611da9614f1f565b6020026020010151838281518110611dc357611dc3614f1f565b60200260200101906001600160a01b031690816001600160a01b031681525050838181518110611df557611df5614f1f565b6020026020010151828281518110611e0f57611e0f614f1f565b6020908102919091010152600101611d8e565b5073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac08260018451611e4791906153ad565b81518110611e5757611e57614f1f565b60200260200101906001600160a01b031690816001600160a01b031681525050848160018451611e8791906153ad565b81518110611e9757611e97614f1f565b6020026020010181815250505b9097909650945050505050565b606654606090600190600290811603611edd5760405163840a48d560e01b815260040160405180910390fd5b611ee6836114bf565b611f035760405163a5c7c44560e01b815260040160405180910390fd5b611f0c836115d7565b15611f2a576040516311ca333560e31b815260040160405180910390fd5b6001600160a01b038316611f51576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b038084166000818152609a602052604090205490911690331480611f845750336001600160a01b038216145b80611fab57506001600160a01b038181166000908152609960205260409020600101541633145b611fc857604051631e499a2360e11b815260040160405180910390fd5b600080611fd486611afa565b9092509050336001600160a01b0387161461202a57826001600160a01b0316866001600160a01b03167ff0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a60405160405180910390a35b826001600160a01b0316866001600160a01b03167ffee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af4467660405160405180910390a36001600160a01b0386166000908152609a6020526040812080546001600160a01b0319169055825190036120ae576040805160008152602081019091529450612371565b81516001600160401b038111156120c7576120c761478d565b6040519080825280602002602001820160405280156120f0578160200160208202803683370190505b50945060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166339a9a3ed85856040518363ffffffff1660e01b81526004016121439291906153c0565b600060405180830381865afa158015612160573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612188919081019061501a565b905060005b835181101561236e57604080516001808252818301909252600091602080830190803683375050604080516001808252818301909252929350600092915060208083019080368337505060408051600180825281830190925292935060009291506020808301908036833701905050905086848151811061221057612210614f1f565b60200260200101518360008151811061222b5761222b614f1f565b60200260200101906001600160a01b031690816001600160a01b03168152505085848151811061225d5761225d614f1f565b60200260200101518260008151811061227857612278614f1f565b60200260200101818152505084848151811061229657612296614f1f565b6020026020010151816000815181106122b1576122b1614f1f565b60200260200101906001600160401b031690816001600160401b0316815250506122df8b898d8686866129d2565b8a85815181106122f1576122f1614f1f565b602002602001018181525050670de0b6b3a764000060a260008d6001600160a01b03166001600160a01b03168152602001908152602001600020600089878151811061233f5761233f614f1f565b6020908102919091018101516001600160a01b031682528101919091526040016000205550505060010161218d565b50505b50505050919050565b6066546002906004908116036123a35760405163840a48d560e01b815260040160405180910390fd5b600260c954036123f55760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611270565b600260c95561240685858585613959565b5050600160c955505050565b61241b336114bf565b1561243957604051633bf2b50360e11b815260040160405180910390fd5b612442836115d7565b61245f576040516325ec6c1f60e01b815260040160405180910390fd5b6110e733848484613067565b612474336115d7565b612491576040516325ec6c1f60e01b815260040160405180910390fd5b610fa1338261300f565b6124a36137f3565b6001600160a01b0381166125085760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401611270565b610fa181613618565b60007f00000000000000000000000000000000000000000000000000000000000000004603612541575060975490565b612549613581565b905090565b6001600160a01b038084166000908152609a602052604081205460609216905b838110156128c15773beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac085858381811061259d5761259d614f1f565b90506020020160208101906125b29190614697565b6001600160a01b031603612687576040516360f4062b60e01b81526001600160a01b0387811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000909116906360f4062b90602401602060405180830381865afa15801561262b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061264f919061524c565b905060008113156126605780612663565b60005b84838151811061267557612675614f1f565b6020026020010181815250505061276c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637a7e0d92878787858181106126c9576126c9614f1f565b90506020020160208101906126de9190614697565b6040516001600160e01b031960e085901b1681526001600160a01b03928316600482015291166024820152604401602060405180830381865afa158015612729573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061274d919061524c565b83828151811061275f5761275f614f1f565b6020026020010181815250505b6001600160a01b038216156128b95760007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663b47265e2848888868181106127bf576127bf614f1f565b90506020020160208101906127d49190614697565b6040516001600160e01b031960e085901b1681526001600160a01b03928316600482015291166024820152604401602060405180830381865afa15801561281f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612843919061514f565b9050612899876128798989898781811061285f5761285f614f1f565b90506020020160208101906128749190614697565b61366a565b86858151811061288b5761288b614f1f565b60200260200101518461407b565b8483815181106128ab576128ab614f1f565b602002602001018181525050505b60010161256e565b50509392505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561291d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129419190615132565b6001600160a01b0316336001600160a01b0316146129725760405163794821ff60e01b815260040160405180910390fd5b60665419811960665419161461299b5760405163c61dca5d60e01b815260040160405180910390fd5b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016111cc565b60006001600160a01b0387166129fb576040516339b190bb60e11b815260040160405180910390fd5b8351600003612a1d5760405163796cc52560e01b815260040160405180910390fd5b846001600160a01b0316876001600160a01b031614612a4f576040516330c4716960e21b815260040160405180910390fd5b600084516001600160401b03811115612a6a57612a6a61478d565b604051908082528060200260200182016040528015612a93578160200160208202803683370190505b50905060005b8551811015612f1b57600073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac06001600160a01b0316878381518110612ad457612ad4614f1f565b60200260200101516001600160a01b031603612b95576040516360f4062b60e01b81526001600160a01b038b811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000909116906360f4062b90602401602060405180830381865afa158015612b55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b79919061524c565b90506000811315612b8a5780612b8d565b60005b915050612c55565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637a7e0d928b898581518110612bd757612bd7614f1f565b60200260200101516040518363ffffffff1660e01b8152600401612c119291906001600160a01b0392831681529116602082015260400190565b602060405180830381865afa158015612c2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c52919061524c565b90505b6000612c7a8b898581518110612c6d57612c6d614f1f565b602002602001015161366a565b9050612ca18b8284898781518110612c9457612c94614f1f565b602002602001015161407b565b878481518110612cb357612cb3614f1f565b60200260200101511115612cda576040516306f2da3960e11b815260040160405180910390fd5b600080612d1a898681518110612cf257612cf2614f1f565b6020026020010151848a8881518110612d0d57612d0d614f1f565b60200260200101516140b9565b9150915080868681518110612d3157612d31614f1f565b60209081029190910101526001600160a01b038c1615612da357612da38c8e8c8881518110612d6257612d62614f1f565b60200260200101518c8981518110612d7c57612d7c614f1f565b60200260200101518c8a81518110612d9657612d96614f1f565b6020026020010151613460565b73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac06001600160a01b03168a8681518110612dd357612dd3614f1f565b60200260200101516001600160a01b031603612e705760405163beffbb8960e01b81526001600160a01b038e81166004830152602482018490527f0000000000000000000000000000000000000000000000000000000000000000169063beffbb8990604401600060405180830381600087803b158015612e5357600080fd5b505af1158015612e67573d6000803e3d6000fd5b50505050612f0c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638c80d4e58e8c8881518110612eb257612eb2614f1f565b6020026020010151856040518463ffffffff1660e01b8152600401612ed9939291906153e4565b600060405180830381600087803b158015612ef357600080fd5b505af1158015612f07573d6000803e3d6000fd5b505050505b50505050806001019050612a99565b506001600160a01b0388166000908152609f60205260408120805491829190612f4383615408565b919050555060006040518060e001604052808b6001600160a01b031681526020018a6001600160a01b03168152602001896001600160a01b031681526020018381526020014263ffffffff1681526020018881526020018481525090506000612fab826115a7565b6000818152609e602052604090819020805460ff19166001179055519091507f9009ab153e8014fbfb02f2217f5cde7aa7f9ad734ae85ca3ee3f4ca2fdd499f990612ff99083908590615421565b60405180910390a19a9950505050505050505050565b6001600160a01b03821660009081526099602052604090208190613033828261545a565b505060405133907ffebe5cd24b2cbc7b065b9d0fdeb904461e4afcff57dd57acda1e7832031ba7ac9061195b9084906150ac565b6066546000906001908116036130905760405163840a48d560e01b815260040160405180910390fd5b6001600160a01b038085166000908152609960205260409020600101541680158015906130c65750336001600160a01b03821614155b80156130db5750336001600160a01b03861614155b1561319e57428460200151101561310557604051630819bdcd60e01b815260040160405180910390fd5b6001600160a01b0381166000908152609c6020908152604080832086845290915290205460ff161561314a57604051630d4c4c9160e21b815260040160405180910390fd5b6001600160a01b0381166000908152609c6020908152604080832086845282528220805460ff1916600117905585015161318b9088908890859088906109dd565b905061319c82828760000151613875565b505b6001600160a01b038681166000818152609a602052604080822080546001600160a01b031916948a169485179055517fc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d87433049190a36000806131fd88611afa565b9150915060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166339a9a3ed89856040518363ffffffff1660e01b81526004016132519291906153c0565b600060405180830381865afa15801561326e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613296919081019061501a565b905060005b83518110156133c45760006132fb8b6132c08d888681518110612c6d57612c6d614f1f565b8585815181106132d2576132d2614f1f565b602002602001015160008887815181106132ee576132ee614f1f565b60200260200101516136ab565b90508060a260008d6001600160a01b03166001600160a01b03168152602001908152602001600020600087858151811061333757613337614f1f565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020819055506133bb8a8c87858151811061337a5761337a614f1f565b602002602001015187868151811061339457613394614f1f565b60200260200101518787815181106133ae576133ae614f1f565b6020026020010151613767565b5060010161329b565b50505050505050505050565b6001600160a01b0381166133f7576040516339b190bb60e11b815260040160405180910390fd5b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b600061346c838361412a565b6001600160a01b0380881660009081526098602090815260408083209389168352929052908120805492935083929091906134a89084906153ad565b92505081905550856001600160a01b03167f6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd8686846040516134ec939291906153e4565b60405180910390a2505050505050565b6065546001600160a01b031615801561351d57506001600160a01b03821615155b61353a576040516339b190bb60e11b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261357d826133d0565b5050565b604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b03808316600090815260a2602090815260408083209385168352929052908120548082036136a45750670de0b6b3a76400005b9392505050565b600080836000036136d8576136d16001600160401b038616670de0b6b3a76400006154bd565b905061375d565b670de0b6b3a76400006001600160401b0386166136f5858761539a565b6136ff91906154df565b61370991906154bd565b83670de0b6b3a76400006001600160401b038816816137288b8a6154df565b61373291906154bd565b61373c91906154df565b61374691906154bd565b613750919061539a565b61375a91906154bd565b90505b9695505050505050565b6000613773838361412a565b6001600160a01b0380881660009081526098602090815260408083209389168352929052908120805492935083929091906137af90849061539a565b92505081905550856001600160a01b03167f1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c8686846040516134ec939291906153e4565b6033546001600160a01b031633146116235760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401611270565b6000670de0b6b3a764000061386b6001600160401b038416856154df565b6136a491906154bd565b6001600160a01b0383163b1561391e57604051630b135d3f60e11b808252906001600160a01b03851690631626ba7e906138b590869086906004016154f6565b602060405180830381865afa1580156138d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138f6919061554b565b6001600160e01b031916146110e75760405163026aa43360e31b815260040160405180910390fd5b826001600160a01b03166139328383614148565b6001600160a01b0316146110e757604051636d20896160e11b815260040160405180910390fd5b61396660a0850185614f55565b83149050613987576040516343714afd60e01b815260040160405180910390fd5b6139976060850160408601614697565b6001600160a01b0316336001600160a01b0316146139c8576040516316110d3560e21b815260040160405180910390fd5b60006139d661063286615575565b6000818152609e60205260408120549192509060ff16613a09576040516387c9d21960e01b815260040160405180910390fd5b6367757490613a1e60a0880160808901615581565b63ffffffff161015613a75574361c4e0613a3e60a0890160808a01615581565b63ffffffff16613a4e919061539a565b1115613a6d576040516378f67ae160e11b815260040160405180910390fd5b506001613ad6565b427f0000000000000000000000000000000000000000000000000000000000000000613aa760a0890160808a01615581565b613ab1919061559e565b63ffffffff161115613ad6576040516378f67ae160e11b815260040160405180910390fd5b60006001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663858d0b47613b1760408a0160208b01614697565b613b2460a08b018b614f55565b7f0000000000000000000000000000000000000000000000000000000000000000613b5560a08e0160808f01615581565b613b5f919061559e565b6040518563ffffffff1660e01b8152600401613b7e94939291906155ba565b600060405180830381865afa158015613b9b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613bc3919081019061501a565b905060005b613bd560a0890189614f55565b90508110156140215760008315613c1157613bf360c08a018a614f55565b83818110613c0357613c03614f1f565b905060200201359050613c5a565b613c57613c2160c08b018b614f55565b84818110613c3157613c31614f1f565b90506020020135848481518110613c4a57613c4a614f1f565b602002602001015161384d565b90505b8515613e165773beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0613c8260a08b018b614f55565b84818110613c9257613c92614f1f565b9050602002016020810190613ca79190614697565b6001600160a01b031603613d48576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663387b1300613cf160208c018c614697565b33846040518463ffffffff1660e01b8152600401613d11939291906153e4565b600060405180830381600087803b158015613d2b57600080fd5b505af1158015613d3f573d6000803e3d6000fd5b50505050614018565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663c608c7f333613d8560a08d018d614f55565b86818110613d9557613d95614f1f565b9050602002016020810190613daa9190614697565b848c8c88818110613dbd57613dbd614f1f565b9050602002016020810190613dd29190614697565b60405160e086901b6001600160e01b03191681526001600160a01b039485166004820152928416602484015260448301919091529091166064820152608401613d11565b73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0613e3860a08b018b614f55565b84818110613e4857613e48614f1f565b9050602002016020810190613e5d9190614697565b6001600160a01b031603613f1e576001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016630e81073c613ea760208c018c614697565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024810184905260440160408051808303816000875af1158015613ef3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613f1791906155f6565b5050614018565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c4623ea1338a8a86818110613f6057613f60614f1f565b9050602002016020810190613f759190614697565b613f8260a08e018e614f55565b87818110613f9257613f92614f1f565b9050602002016020810190613fa79190614697565b60405160e085901b6001600160e01b03191681526001600160a01b0393841660048201529183166024830152909116604482015260648101849052608401600060405180830381600087803b158015613fff57600080fd5b505af1158015614013573d6000803e3d6000fd5b505050505b50600101613bc8565b506000838152609e602052604090819020805460ff19169055517fc97098c2f658800b4df29001527f7324bcdffcf6e8751a699ab920a1eced5b1d9061406a9085815260200190565b60405180910390a150505050505050565b6000670de0b6b3a76400006001600160401b0383168161409b86886154df565b6140a591906154bd565b6140af91906154df565b61132b91906154bd565b6000806001600160401b038316670de0b6b3a7640000856140da82896154df565b6140e491906154bd565b6140ee91906154df565b6140f891906154bd565b91506001600160401b038316614116670de0b6b3a7640000876154df565b61412091906154bd565b9050935093915050565b60006001600160401b03821661386b670de0b6b3a7640000856154df565b60008060006141578585614164565b91509150611a39816141d2565b600080825160410361419a5760208301516040840151606085015160001a61418e87828585614388565b945094505050506141cb565b82516040036141c357602083015160408401516141b8868383614475565b9350935050506141cb565b506000905060025b9250929050565b60008160048111156141e6576141e661561a565b036141ee5750565b60018160048111156142025761420261561a565b0361424f5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401611270565b60028160048111156142635761426361561a565b036142b05760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401611270565b60038160048111156142c4576142c461561a565b0361431c5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401611270565b60048160048111156143305761433061561a565b03610fa15760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401611270565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156143bf575060009050600361446c565b8460ff16601b141580156143d757508460ff16601c14155b156143e8575060009050600461446c565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561443c573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166144655760006001925092505061446c565b9150600090505b94509492505050565b6000806001600160ff1b0383168161449260ff86901c601b61539a565b90506144a087828885614388565b935093505050935093915050565b6001600160a01b0381168114610fa157600080fd5b80356144ce816144ae565b919050565b600080600080600060a086880312156144eb57600080fd5b85356144f6816144ae565b94506020860135614506816144ae565b93506040860135614516816144ae565b94979396509394606081013594506080013592915050565b60008083601f84011261454057600080fd5b5081356001600160401b0381111561455757600080fd5b6020830191508360208260051b85010111156141cb57600080fd5b6000806020838503121561458557600080fd5b82356001600160401b0381111561459b57600080fd5b6145a78582860161452e565b90969095509350505050565b602080825282518282018190526000918401906040840190835b81811015610e085783518352602093840193909201916001016145cd565b6000606082840312156145fd57600080fd5b50919050565b60008083601f84011261461557600080fd5b5081356001600160401b0381111561462c57600080fd5b6020830191508360208285010111156141cb57600080fd5b60008060006080848603121561465957600080fd5b61466385856145eb565b925060608401356001600160401b0381111561467e57600080fd5b61468a86828701614603565b9497909650939450505050565b6000602082840312156146a957600080fd5b81356136a4816144ae565b6000806000606084860312156146c957600080fd5b83356146d4816144ae565b925060208401356146e4816144ae565b929592945050506040919091013590565b60006020828403121561470757600080fd5b5035919050565b6000806000806080858703121561472457600080fd5b843561472f816144ae565b9350602085013561473f816144ae565b93969395505050506040820135916060013590565b6000806040838503121561476757600080fd5b8235614772816144ae565b91506020830135614782816144ae565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b03811182821017156147c5576147c561478d565b60405290565b604080519081016001600160401b03811182821017156147c5576147c561478d565b604051601f8201601f191681016001600160401b03811182821017156148155761481561478d565b604052919050565b63ffffffff81168114610fa157600080fd5b80356144ce8161481d565b60006001600160401b038211156148535761485361478d565b5060051b60200190565b600082601f83011261486e57600080fd5b813561488161487c8261483a565b6147ed565b8082825260208201915060208360051b8601019250858311156148a357600080fd5b602085015b838110156148c95780356148bb816144ae565b8352602092830192016148a8565b5095945050505050565b600082601f8301126148e457600080fd5b81356148f261487c8261483a565b8082825260208201915060208360051b86010192508583111561491457600080fd5b602085015b838110156148c9578035835260209283019201614919565b600060e0828403121561494357600080fd5b61494b6147a3565b9050614956826144c3565b8152614964602083016144c3565b6020820152614975604083016144c3565b6040820152606082810135908201526149906080830161482f565b608082015260a08201356001600160401b038111156149ae57600080fd5b6149ba8482850161485d565b60a08301525060c08201356001600160401b038111156149d957600080fd5b6149e5848285016148d3565b60c08301525092915050565b600060208284031215614a0357600080fd5b81356001600160401b03811115614a1957600080fd5b6116ef84828501614931565b600060208284031215614a3757600080fd5b813560ff811681146136a457600080fd5b600060408284031215614a5a57600080fd5b614a626147cb565b905081356001600160401b03811115614a7a57600080fd5b8201601f81018413614a8b57600080fd5b80356001600160401b03811115614aa457614aa461478d565b614ab7601f8201601f19166020016147ed565b818152856020838501011115614acc57600080fd5b81602084016020830137600060209282018301528352928301359282019290925292915050565b600080600080600060a08688031215614b0b57600080fd5b8535614b16816144ae565b94506020860135614b26816144ae565b935060408601356001600160401b03811115614b4157600080fd5b614b4d88828901614a48565b93505060608601356001600160401b03811115614b6957600080fd5b614b7588828901614a48565b95989497509295608001359392505050565b60008060008060008060608789031215614ba057600080fd5b86356001600160401b03811115614bb657600080fd5b614bc289828a0161452e565b90975095505060208701356001600160401b03811115614be157600080fd5b614bed89828a0161452e565b90955093505060408701356001600160401b03811115614c0c57600080fd5b614c1889828a0161452e565b979a9699509497509295939492505050565b60008060208385031215614c3d57600080fd5b82356001600160401b03811115614c5357600080fd5b6145a785828601614603565b60008060408385031215614c7257600080fd5b8235614c7d816144ae565b915060208301356001600160401b03811115614c9857600080fd5b614ca48582860161485d565b9150509250929050565b600081518084526020840193506020830160005b82811015614ce0578151865260209586019590910190600101614cc2565b5093949350505050565b6020815260006136a46020830184614cae565b60008060408385031215614d1057600080fd5b8235614d1b816144ae565b946020939093013593505050565b60008060008060808587031215614d3f57600080fd5b8435614d4a816144ae565b9350602085013592506040850135614d61816144ae565b9396929550929360600135925050565b600081518084526020840193506020830160005b82811015614ce05781516001600160a01b0316865260209586019590910190600101614d85565b604081526000614dbf6040830185614d71565b828103602084015261132b8185614cae565b8015158114610fa157600080fd5b60008060008060608587031215614df557600080fd5b84356001600160401b03811115614e0b57600080fd5b850160e08188031215614e1d57600080fd5b935060208501356001600160401b03811115614e3857600080fd5b614e448782880161452e565b9094509250506040850135614e5881614dd1565b939692955090935050565b600080600060608486031215614e7857600080fd5b8335614e83816144ae565b925060208401356001600160401b03811115614e9e57600080fd5b614eaa86828701614a48565b93969395505050506040919091013590565b600060608284031215614ece57600080fd5b6136a483836145eb565b600080600060408486031215614eed57600080fd5b8335614ef8816144ae565b925060208401356001600160401b03811115614f1357600080fd5b61468a8682870161452e565b634e487b7160e01b600052603260045260246000fd5b60008235605e19833603018112614f4b57600080fd5b9190910192915050565b6000808335601e19843603018112614f6c57600080fd5b8301803591506001600160401b03821115614f8657600080fd5b6020019150600581901b36038213156141cb57600080fd5b81835260208301925060008160005b84811015614ce0578135614fc0816144ae565b6001600160a01b031686526020958601959190910190600101614fad565b6001600160a01b038416815260406020820181905260009061132b9083018486614f9e565b80516001600160401b03811681146144ce57600080fd5b60006020828403121561502c57600080fd5b81516001600160401b0381111561504257600080fd5b8201601f8101841361505357600080fd5b805161506161487c8261483a565b8082825260208201915060208360051b85010192508683111561508357600080fd5b6020840193505b8284101561375d5761509b84615003565b82526020938401939091019061508a565b6060810182356150bb816144ae565b6001600160a01b0316825260208301356150d4816144ae565b6001600160a01b0316602083015260408301356150f08161481d565b63ffffffff811660408401525092915050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60006020828403121561514457600080fd5b81516136a4816144ae565b60006020828403121561516157600080fd5b6136a482615003565b60006020828403121561517c57600080fd5b81516136a481614dd1565b80516001600160a01b03908116835260208083015182169084015260408083015190911690830152606080820151908301526080808201516000916151d39085018263ffffffff169052565b5060a082015160e060a08501526151ed60e0850182614d71565b905060c083015184820360c086015261132b8282614cae565b6020815260006136a46020830184615187565b6000823560de19833603018112614f4b57600080fd5b60006020828403121561524157600080fd5b81356136a481614dd1565b60006020828403121561525e57600080fd5b5051919050565b600082601f83011261527657600080fd5b815161528461487c8261483a565b8082825260208201915060208360051b8601019250858311156152a657600080fd5b602085015b838110156148c95780518352602092830192016152ab565b600080604083850312156152d657600080fd5b82516001600160401b038111156152ec57600080fd5b8301601f810185136152fd57600080fd5b805161530b61487c8261483a565b8082825260208201915060208360051b85010192508783111561532d57600080fd5b6020840193505b82841015615358578351615347816144ae565b825260209384019390910190615334565b8095505050505060208301516001600160401b0381111561537857600080fd5b614ca485828601615265565b634e487b7160e01b600052601160045260246000fd5b8082018082111561160b5761160b615384565b8181038181111561160b5761160b615384565b6001600160a01b03831681526040602082018190526000906116ef90830184614d71565b6001600160a01b039384168152919092166020820152604081019190915260600190565b60006001820161541a5761541a615384565b5060010190565b8281526040602082015260006116ef6040830184615187565b80546001600160a01b0319166001600160a01b0392909216919091179055565b8135615465816144ae565b61546f818361543a565b50600181016020830135615482816144ae565b61548c818361543a565b50604083013561549b8161481d565b815463ffffffff60a01b191660a09190911b63ffffffff60a01b161790555050565b6000826154da57634e487b7160e01b600052601260045260246000fd5b500490565b808202811582820484141761160b5761160b615384565b828152604060208201526000825180604084015260005b8181101561552a576020818601810151606086840101520161550d565b506000606082850101526060601f19601f8301168401019150509392505050565b60006020828403121561555d57600080fd5b81516001600160e01b0319811681146136a457600080fd5b600061160b3683614931565b60006020828403121561559357600080fd5b81356136a48161481d565b63ffffffff818116838216019081111561160b5761160b615384565b6001600160a01b03851681526060602082018190526000906155df9083018587614f9e565b905063ffffffff8316604083015295945050505050565b6000806040838503121561560957600080fd5b505080516020909101519092909150565b634e487b7160e01b600052602160045260246000fdfea2646970667358221220ed9ad13f9ddcb5a00dbf9fee4255469304e5d9a5b4fc67567fbb31db54dd19ab64736f6c634300081b0033",
}

// DelegationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DelegationManagerMetaData.ABI instead.
var DelegationManagerABI = DelegationManagerMetaData.ABI

// DelegationManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DelegationManagerMetaData.Bin instead.
var DelegationManagerBin = DelegationManagerMetaData.Bin

// DeployDelegationManager deploys a new Ethereum contract, binding an instance of DelegationManager to it.
func DeployDelegationManager(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _slasher common.Address, _eigenPodManager common.Address, _avsDirectory common.Address, _allocationManager common.Address, _MIN_WITHDRAWAL_DELAY uint32) (common.Address, *types.Transaction, *DelegationManager, error) {
	parsed, err := DelegationManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DelegationManagerBin), backend, _strategyManager, _slasher, _eigenPodManager, _avsDirectory, _allocationManager, _MIN_WITHDRAWAL_DELAY)
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

// LEGACYMINWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xcb00387b.
//
// Solidity: function LEGACY_MIN_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) LEGACYMINWITHDRAWALDELAYBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "LEGACY_MIN_WITHDRAWAL_DELAY_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LEGACYMINWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xcb00387b.
//
// Solidity: function LEGACY_MIN_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_DelegationManager *DelegationManagerSession) LEGACYMINWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _DelegationManager.Contract.LEGACYMINWITHDRAWALDELAYBLOCKS(&_DelegationManager.CallOpts)
}

// LEGACYMINWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xcb00387b.
//
// Solidity: function LEGACY_MIN_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) LEGACYMINWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _DelegationManager.Contract.LEGACYMINWITHDRAWALDELAYBLOCKS(&_DelegationManager.CallOpts)
}

// LEGACYWITHDRAWALSTIMESTAMP is a free data retrieval call binding the contract method 0xc0525f02.
//
// Solidity: function LEGACY_WITHDRAWALS_TIMESTAMP() view returns(uint32)
func (_DelegationManager *DelegationManagerCaller) LEGACYWITHDRAWALSTIMESTAMP(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "LEGACY_WITHDRAWALS_TIMESTAMP")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LEGACYWITHDRAWALSTIMESTAMP is a free data retrieval call binding the contract method 0xc0525f02.
//
// Solidity: function LEGACY_WITHDRAWALS_TIMESTAMP() view returns(uint32)
func (_DelegationManager *DelegationManagerSession) LEGACYWITHDRAWALSTIMESTAMP() (uint32, error) {
	return _DelegationManager.Contract.LEGACYWITHDRAWALSTIMESTAMP(&_DelegationManager.CallOpts)
}

// LEGACYWITHDRAWALSTIMESTAMP is a free data retrieval call binding the contract method 0xc0525f02.
//
// Solidity: function LEGACY_WITHDRAWALS_TIMESTAMP() view returns(uint32)
func (_DelegationManager *DelegationManagerCallerSession) LEGACYWITHDRAWALSTIMESTAMP() (uint32, error) {
	return _DelegationManager.Contract.LEGACYWITHDRAWALSTIMESTAMP(&_DelegationManager.CallOpts)
}

// MINWITHDRAWALDELAY is a free data retrieval call binding the contract method 0x4a5f2b5d.
//
// Solidity: function MIN_WITHDRAWAL_DELAY() view returns(uint32)
func (_DelegationManager *DelegationManagerCaller) MINWITHDRAWALDELAY(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "MIN_WITHDRAWAL_DELAY")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MINWITHDRAWALDELAY is a free data retrieval call binding the contract method 0x4a5f2b5d.
//
// Solidity: function MIN_WITHDRAWAL_DELAY() view returns(uint32)
func (_DelegationManager *DelegationManagerSession) MINWITHDRAWALDELAY() (uint32, error) {
	return _DelegationManager.Contract.MINWITHDRAWALDELAY(&_DelegationManager.CallOpts)
}

// MINWITHDRAWALDELAY is a free data retrieval call binding the contract method 0x4a5f2b5d.
//
// Solidity: function MIN_WITHDRAWAL_DELAY() view returns(uint32)
func (_DelegationManager *DelegationManagerCallerSession) MINWITHDRAWALDELAY() (uint32, error) {
	return _DelegationManager.Contract.MINWITHDRAWALDELAY(&_DelegationManager.CallOpts)
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

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_DelegationManager *DelegationManagerCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_DelegationManager *DelegationManagerSession) AllocationManager() (common.Address, error) {
	return _DelegationManager.Contract.AllocationManager(&_DelegationManager.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_DelegationManager *DelegationManagerCallerSession) AllocationManager() (common.Address, error) {
	return _DelegationManager.Contract.AllocationManager(&_DelegationManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_DelegationManager *DelegationManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_DelegationManager *DelegationManagerSession) AvsDirectory() (common.Address, error) {
	return _DelegationManager.Contract.AvsDirectory(&_DelegationManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_DelegationManager *DelegationManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _DelegationManager.Contract.AvsDirectory(&_DelegationManager.CallOpts)
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

// GetOperatorScaledShares is a free data retrieval call binding the contract method 0xa1610781.
//
// Solidity: function getOperatorScaledShares(address operator, address[] strategies) view returns(uint256[])
func (_DelegationManager *DelegationManagerCaller) GetOperatorScaledShares(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "getOperatorScaledShares", operator, strategies)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOperatorScaledShares is a free data retrieval call binding the contract method 0xa1610781.
//
// Solidity: function getOperatorScaledShares(address operator, address[] strategies) view returns(uint256[])
func (_DelegationManager *DelegationManagerSession) GetOperatorScaledShares(operator common.Address, strategies []common.Address) ([]*big.Int, error) {
	return _DelegationManager.Contract.GetOperatorScaledShares(&_DelegationManager.CallOpts, operator, strategies)
}

// GetOperatorScaledShares is a free data retrieval call binding the contract method 0xa1610781.
//
// Solidity: function getOperatorScaledShares(address operator, address[] strategies) view returns(uint256[])
func (_DelegationManager *DelegationManagerCallerSession) GetOperatorScaledShares(operator common.Address, strategies []common.Address) ([]*big.Int, error) {
	return _DelegationManager.Contract.GetOperatorScaledShares(&_DelegationManager.CallOpts, operator, strategies)
}

// GetWithdrawableStakerShares is a free data retrieval call binding the contract method 0xf878776e.
//
// Solidity: function getWithdrawableStakerShares(address staker, address[] strategies) view returns(uint256[] shares)
func (_DelegationManager *DelegationManagerCaller) GetWithdrawableStakerShares(opts *bind.CallOpts, staker common.Address, strategies []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "getWithdrawableStakerShares", staker, strategies)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetWithdrawableStakerShares is a free data retrieval call binding the contract method 0xf878776e.
//
// Solidity: function getWithdrawableStakerShares(address staker, address[] strategies) view returns(uint256[] shares)
func (_DelegationManager *DelegationManagerSession) GetWithdrawableStakerShares(staker common.Address, strategies []common.Address) ([]*big.Int, error) {
	return _DelegationManager.Contract.GetWithdrawableStakerShares(&_DelegationManager.CallOpts, staker, strategies)
}

// GetWithdrawableStakerShares is a free data retrieval call binding the contract method 0xf878776e.
//
// Solidity: function getWithdrawableStakerShares(address staker, address[] strategies) view returns(uint256[] shares)
func (_DelegationManager *DelegationManagerCallerSession) GetWithdrawableStakerShares(staker common.Address, strategies []common.Address) ([]*big.Int, error) {
	return _DelegationManager.Contract.GetWithdrawableStakerShares(&_DelegationManager.CallOpts, staker, strategies)
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

// OperatorScaledShares is a free data retrieval call binding the contract method 0x9fb4ee32.
//
// Solidity: function operatorScaledShares(address , address ) view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) OperatorScaledShares(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "operatorScaledShares", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorScaledShares is a free data retrieval call binding the contract method 0x9fb4ee32.
//
// Solidity: function operatorScaledShares(address , address ) view returns(uint256)
func (_DelegationManager *DelegationManagerSession) OperatorScaledShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.OperatorScaledShares(&_DelegationManager.CallOpts, arg0, arg1)
}

// OperatorScaledShares is a free data retrieval call binding the contract method 0x9fb4ee32.
//
// Solidity: function operatorScaledShares(address , address ) view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) OperatorScaledShares(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.OperatorScaledShares(&_DelegationManager.CallOpts, arg0, arg1)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address operator, address strategy) view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) OperatorShares(opts *bind.CallOpts, operator common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "operatorShares", operator, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address operator, address strategy) view returns(uint256)
func (_DelegationManager *DelegationManagerSession) OperatorShares(operator common.Address, strategy common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.OperatorShares(&_DelegationManager.CallOpts, operator, strategy)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address operator, address strategy) view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) OperatorShares(operator common.Address, strategy common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.OperatorShares(&_DelegationManager.CallOpts, operator, strategy)
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

// StakerScalingFactors is a free data retrieval call binding the contract method 0x51260e23.
//
// Solidity: function stakerScalingFactors(address , address ) view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) StakerScalingFactors(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "stakerScalingFactors", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerScalingFactors is a free data retrieval call binding the contract method 0x51260e23.
//
// Solidity: function stakerScalingFactors(address , address ) view returns(uint256)
func (_DelegationManager *DelegationManagerSession) StakerScalingFactors(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.StakerScalingFactors(&_DelegationManager.CallOpts, arg0, arg1)
}

// StakerScalingFactors is a free data retrieval call binding the contract method 0x51260e23.
//
// Solidity: function stakerScalingFactors(address , address ) view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) StakerScalingFactors(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.StakerScalingFactors(&_DelegationManager.CallOpts, arg0, arg1)
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

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0xe4cc3f90.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, bool receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerTransactor) CompleteQueuedWithdrawal(opts *bind.TransactOpts, withdrawal IDelegationManagerWithdrawal, tokens []common.Address, receiveAsTokens bool) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "completeQueuedWithdrawal", withdrawal, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0xe4cc3f90.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, bool receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerSession) CompleteQueuedWithdrawal(withdrawal IDelegationManagerWithdrawal, tokens []common.Address, receiveAsTokens bool) (*types.Transaction, error) {
	return _DelegationManager.Contract.CompleteQueuedWithdrawal(&_DelegationManager.TransactOpts, withdrawal, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0xe4cc3f90.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, bool receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerTransactorSession) CompleteQueuedWithdrawal(withdrawal IDelegationManagerWithdrawal, tokens []common.Address, receiveAsTokens bool) (*types.Transaction, error) {
	return _DelegationManager.Contract.CompleteQueuedWithdrawal(&_DelegationManager.TransactOpts, withdrawal, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x9435bb43.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerTransactor) CompleteQueuedWithdrawals(opts *bind.TransactOpts, withdrawals []IDelegationManagerWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "completeQueuedWithdrawals", withdrawals, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x9435bb43.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerSession) CompleteQueuedWithdrawals(withdrawals []IDelegationManagerWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _DelegationManager.Contract.CompleteQueuedWithdrawals(&_DelegationManager.TransactOpts, withdrawals, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x9435bb43.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerTransactorSession) CompleteQueuedWithdrawals(withdrawals []IDelegationManagerWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _DelegationManager.Contract.CompleteQueuedWithdrawals(&_DelegationManager.TransactOpts, withdrawals, tokens, receiveAsTokens)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address strategy, uint256 removedShares) returns()
func (_DelegationManager *DelegationManagerTransactor) DecreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, removedShares *big.Int) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "decreaseDelegatedShares", staker, strategy, removedShares)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address strategy, uint256 removedShares) returns()
func (_DelegationManager *DelegationManagerSession) DecreaseDelegatedShares(staker common.Address, strategy common.Address, removedShares *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.DecreaseDelegatedShares(&_DelegationManager.TransactOpts, staker, strategy, removedShares)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x132d4967.
//
// Solidity: function decreaseDelegatedShares(address staker, address strategy, uint256 removedShares) returns()
func (_DelegationManager *DelegationManagerTransactorSession) DecreaseDelegatedShares(staker common.Address, strategy common.Address, removedShares *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.DecreaseDelegatedShares(&_DelegationManager.TransactOpts, staker, strategy, removedShares)
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

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x3c651cf2.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 existingShares, uint256 addedShares) returns()
func (_DelegationManager *DelegationManagerTransactor) IncreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, existingShares *big.Int, addedShares *big.Int) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "increaseDelegatedShares", staker, strategy, existingShares, addedShares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x3c651cf2.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 existingShares, uint256 addedShares) returns()
func (_DelegationManager *DelegationManagerSession) IncreaseDelegatedShares(staker common.Address, strategy common.Address, existingShares *big.Int, addedShares *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.IncreaseDelegatedShares(&_DelegationManager.TransactOpts, staker, strategy, existingShares, addedShares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x3c651cf2.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 existingShares, uint256 addedShares) returns()
func (_DelegationManager *DelegationManagerTransactorSession) IncreaseDelegatedShares(staker common.Address, strategy common.Address, existingShares *big.Int, addedShares *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.IncreaseDelegatedShares(&_DelegationManager.TransactOpts, staker, strategy, existingShares, addedShares)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_DelegationManager *DelegationManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "initialize", initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_DelegationManager *DelegationManagerSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.Initialize(&_DelegationManager.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_DelegationManager *DelegationManagerTransactorSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.Initialize(&_DelegationManager.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
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
