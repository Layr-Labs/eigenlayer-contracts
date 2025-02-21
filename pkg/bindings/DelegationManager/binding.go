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

// IDelegationManagerTypesQueuedWithdrawalParams is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerTypesQueuedWithdrawalParams struct {
	Strategies           []common.Address
	DepositShares        []*big.Int
	DeprecatedWithdrawer common.Address
}

// IDelegationManagerTypesWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerTypesWithdrawal struct {
	Staker       common.Address
	DelegatedTo  common.Address
	Withdrawer   common.Address
	Nonce        *big.Int
	StartBlock   uint32
	Strategies   []common.Address
	ScaledShares []*big.Int
}

// ISignatureUtilsSignatureWithExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithExpiry struct {
	Signature []byte
	Expiry    *big.Int
}

// DelegationManagerMetaData contains all meta data concerning the DelegationManager contract.
var DelegationManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_eigenPodManager\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_MIN_WITHDRAWAL_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DELEGATION_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burnOperatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"prevMaxMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newMaxMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateDelegationApprovalDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateWithdrawalRoot\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawal\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawals\",\"inputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertToDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"withdrawableShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cumulativeWithdrawalsQueued\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"totalQueued\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"curDepositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beaconChainSlashingFactorDecrease\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateTo\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegatedTo\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApprover\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApproverSaltIsSpent\",\"inputs\":[{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"spent\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositScalingFactor\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorsShares\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQueuedWithdrawal\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQueuedWithdrawalRoots\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQueuedWithdrawals\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"shares\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashableSharesInQueue\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawableShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"withdrawableShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"depositShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"prevDepositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addedShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minWithdrawalDelayBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyOperatorDetails\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newDelegationApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingWithdrawals\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"pending\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueWithdrawals\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.QueuedWithdrawalParams[]\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"depositShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"__deprecated_withdrawer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redelegate\",\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newOperatorApproverSig\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerAsOperator\",\"inputs\":[{\"name\":\"initDelegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorMetadataURI\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DelegationApproverUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newDelegationApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositScalingFactorUpdated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"newDepositScalingFactor\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMetadataURIUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesBurned\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesDecreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesIncreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashingWithdrawalCompleted\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashingWithdrawalQueued\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"withdrawal\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"sharesToWithdraw\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerForceUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ActivelyDelegated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerCannotUndelegate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FullySlashed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSnapshotOrdering\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotActivelyDelegated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAllocationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEigenPodManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManagerOrEigenPodManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorsCannotUndelegate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaltSpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalDelayNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalNotQueued\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawerNotCaller\",\"inputs\":[]}]",
	Bin: "0x610180604052348015610010575f5ffd5b50604051615e7b380380615e7b83398101604081905261002f9161021c565b8186868684876001600160a01b03811661005c576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805293841660a05291831660c05290911660e05263ffffffff16610100524661012052610125604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b610140526001600160a01b03166101605261013e610149565b5050505050506102a7565b5f54610100900460ff16156101b45760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614610203575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610219575f5ffd5b50565b5f5f5f5f5f5f60c08789031215610231575f5ffd5b865161023c81610205565b602088015190965061024d81610205565b604088015190955061025e81610205565b606088015190945061026f81610205565b608088015190935061028081610205565b60a088015190925063ffffffff81168114610299575f5ffd5b809150509295509295509295565b60805160a05160c05160e05161010051610120516101405161016051615ad86103a35f395f818161042c015261336501525f61279201525f6126d201525f81816106ed01528181611504015281816135f0015261380e01525f818161073d01528181610da901528181610f5a0152818161173801528181611b8c015281816123b701528181612967015261341c01525f818161045301528181610ee00152818161169f015281816118fd0152818161315f0152613c5801525f818161038901528181610eae01528181611851015281816124a70152613c3201525f81816105db01528181610b410152818161107a01526127b60152615ad85ff3fe608060405234801561000f575f5ffd5b50600436106102cb575f3560e01c8063778e55f31161017b578063c448feb8116100e4578063ee74937f1161009e578063f2fde38b11610079578063f2fde38b146107de578063f698da25146107f1578063fabc1cbc146107f9578063fd8aa88d1461080c575f5ffd5b8063ee74937f14610798578063eea9064b146107ab578063f0e0e676146107be575f5ffd5b8063c448feb8146106e3578063c978f7ac14610717578063ca8aa7c714610738578063cd6dc6871461075f578063da8be86414610772578063e4cc3f9014610785575f5ffd5b80639435bb43116101355780639435bb431461063c578063a17884841461064f578063a33a34331461066e578063b7f06ebe14610681578063bb45fef2146106a3578063bfae3fd2146106d0575f5ffd5b8063778e55f31461059957806378296ec5146105c3578063886f1195146105d65780638da5cb5b146105fd578063900413471461060e5780639104c31914610621575f5ffd5b806354b7c96c116102375780635dd68579116101f157806366d5ba93116101cc57806366d5ba931461054a5780636d70f7ae1461056b5780636e1744481461057e578063715018a614610591575f5ffd5b80635dd68579146104ee57806360a0d1ce1461050f57806365da126414610522575f5ffd5b806354b7c96c14610475578063595c6a6714610488578063597b36da146104905780635ac86ab7146104a35780635c975abb146104c65780635d975e88146104ce575f5ffd5b806339b70e381161028857806339b70e38146103845780633c651cf2146103c35780633cdeb5e0146103d65780633e28391d146104045780634657e26a146104275780634665bcda1461044e575f5ffd5b806304a4f979146102cf5780630b9f487a146103095780630dd8dd021461031c578063136439dd1461033c57806325df922e146103515780632aa6d88814610371575b5f5ffd5b6102f67f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad81565b6040519081526020015b60405180910390f35b6102f66103173660046149f5565b61081f565b61032f61032a366004614a8c565b6108a7565b6040516103009190614aca565b61034f61034a366004614b01565b610b2c565b005b61036461035f366004614c96565b610c01565b6040516103009190614d44565b61034f61037f366004614da6565b610d61565b6103ab7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610300565b61034f6103d1366004614e04565b610ea3565b6103ab6103e4366004614e47565b6001600160a01b039081165f908152609960205260409020600101541690565b610417610412366004614e47565b610fea565b6040519015158152602001610300565b6103ab7f000000000000000000000000000000000000000000000000000000000000000081565b6103ab7f000000000000000000000000000000000000000000000000000000000000000081565b61034f610483366004614e62565b611009565b61034f611065565b6102f661049e366004614f55565b611114565b6104176104b1366004614f86565b606654600160ff9092169190911b9081161490565b6066546102f6565b6104e16104dc366004614b01565b611143565b604051610300919061505d565b6105016104fc366004614e47565b61125f565b6040516103009291906150bd565b61034f61051d36600461513e565b611694565b6103ab610530366004614e47565b609a6020525f90815260409020546001600160a01b031681565b61055d610558366004614e47565b611829565b60405161030092919061517d565b610417610579366004614e47565b611b29565b6102f661058c366004614e62565b611b61565b61034f611c0b565b6102f66105a7366004614e62565b609860209081525f928352604080842090915290825290205481565b61034f6105d13660046151a1565b611c1c565b6103ab7f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b03166103ab565b61036461061c3660046151f1565b611ca4565b6103ab73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b61034f61064a36600461523d565b611d7a565b6102f661065d366004614e47565b609f6020525f908152604090205481565b61032f61067c3660046152d9565b611e4a565b61041761068f366004614b01565b609e6020525f908152604090205460ff1681565b6104176106b13660046153c0565b609c60209081525f928352604080842090915290825290205460ff1681565b6102f66106de366004614e62565b611e62565b60405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610300565b61072a6107253660046151f1565b611e9e565b6040516103009291906153ea565b6103ab7f000000000000000000000000000000000000000000000000000000000000000081565b61034f61076d3660046153c0565b61212b565b61032f610780366004614e47565b612246565b61034f610793366004615409565b612356565b61034f6107a6366004615487565b6123ac565b61034f6107b93660046152d9565b61254e565b6107d16107cc3660046154d5565b6125b1565b6040516103009190615582565b61034f6107ec366004614e47565b612656565b6102f66126cf565b61034f610807366004614b01565b6127b4565b61032f61081a366004614e47565b6128cb565b604080517f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad60208201526001600160a01b03808616928201929092528187166060820152908516608082015260a0810183905260c081018290525f9061089d9060e001604051602081830303815290604052805190602001206128ee565b9695505050505050565b6066546060906001906002908116036108d35760405163840a48d560e01b815260040160405180910390fd5b5f836001600160401b038111156108ec576108ec614b18565b604051908082528060200260200182016040528015610915578160200160208202803683370190505b50335f908152609a60205260408120549192506001600160a01b03909116905b85811015610b215786868281811061094f5761094f615594565b905060200281019061096191906155a8565b61096f9060208101906155c6565b905087878381811061098357610983615594565b905060200281019061099591906155a8565b61099f90806155c6565b9050146109bf576040516343714afd60e01b815260040160405180910390fd5b5f610a2933848a8a868181106109d7576109d7615594565b90506020028101906109e991906155a8565b6109f390806155c6565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525061291c92505050565b9050610afb33848a8a86818110610a4257610a42615594565b9050602002810190610a5491906155a8565b610a5e90806155c6565b808060200260200160405190810160405280939291908181526020018383602002808284375f920191909152508e92508d9150889050818110610aa357610aa3615594565b9050602002810190610ab591906155a8565b610ac39060208101906155c6565b808060200260200160405190810160405280939291908181526020018383602002808284375f92019190915250889250612a63915050565b848381518110610b0d57610b0d615594565b602090810291909101015250600101610935565b509095945050505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610b8e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bb2919061560b565b610bcf57604051631d77d47760e21b815260040160405180910390fd5b6066548181168114610bf45760405163c61dca5d60e01b815260040160405180910390fd5b610bfd82612f58565b5050565b6001600160a01b038084165f908152609a60205260408120546060921690610c2a86838761291c565b90505f85516001600160401b03811115610c4657610c46614b18565b604051908082528060200260200182016040528015610c6f578160200160208202803683370190505b5090505f5b8651811015610d54576001600160a01b0388165f90815260a260205260408120885182908a9085908110610caa57610caa615594565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f206040518060200160405290815f820154815250509050610d2e878381518110610cfc57610cfc615594565b6020026020010151858481518110610d1657610d16615594565b602002602001015183612f959092919063ffffffff16565b838381518110610d4057610d40615594565b602090810291909101015250600101610c74565b50925050505b9392505050565b610d6a33610fea565b15610d8857604051633bf2b50360e11b815260040160405180910390fd5b604051632b6241f360e11b815233600482015263ffffffff841660248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906356c483e6906044015f604051808303815f87803b158015610df2575f5ffd5b505af1158015610e04573d5f5f3e3d5ffd5b50505050610e123385612fb3565b610e1c3333613015565b6040516001600160a01b038516815233907fa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c19060200160405180910390a2336001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b67080908383604051610e95929190615626565b60405180910390a250505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610f025750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b610f1f5760405163045206a560e21b815260040160405180910390fd5b6001600160a01b038481165f908152609a602052604080822054905163152667d960e31b8152908316600482018190528684166024830152927f0000000000000000000000000000000000000000000000000000000000000000169063a9333ec890604401602060405180830381865afa158015610f9f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fc39190615654565b90505f610fd1878784613118565b9050610fe18388888888866131fa565b50505050505050565b6001600160a01b039081165f908152609a602052604090205416151590565b8161101381613327565b6110305760405163932d94f760e01b815260040160405180910390fd5b61103983611b29565b611056576040516325ec6c1f60e01b815260040160405180910390fd5b6110608383612fb3565b505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156110c7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110eb919061560b565b61110857604051631d77d47760e21b815260040160405180910390fd5b6111125f19612f58565b565b5f81604051602001611126919061505d565b604051602081830303815290604052805190602001209050919050565b61114b6148b1565b5f82815260a46020908152604091829020825160e08101845281546001600160a01b03908116825260018301548116828501526002830154168185015260038201546060820152600482015463ffffffff1660808201526005820180548551818602810186019096528086529194929360a086019392908301828280156111f957602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116111db575b505050505081526020016006820180548060200260200160405190810160405280929190818152602001828054801561124f57602002820191905f5260205f20905b81548152602001906001019080831161123b575b5050505050815250509050919050565b6060805f61126c846128cb565b8051909150806001600160401b0381111561128957611289614b18565b6040519080825280602002602001820160405280156112c257816020015b6112af6148b1565b8152602001906001900390816112a75790505b509350806001600160401b038111156112dd576112dd614b18565b60405190808252806020026020018201604052801561131057816020015b60608152602001906001900390816112fb5790505b506001600160a01b038087165f908152609a60205260408120549295509116905b8281101561168b5760a45f85838151811061134e5761134e615594565b60209081029190910181015182528181019290925260409081015f20815160e08101835281546001600160a01b03908116825260018301548116828601526002830154168184015260038201546060820152600482015463ffffffff1660808201526005820180548451818702810187019095528085529194929360a086019390929083018282801561140857602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116113ea575b505050505081526020016006820180548060200260200160405190810160405280929190818152602001828054801561145e57602002820191905f5260205f20905b81548152602001906001019080831161144a575b50505050508152505086828151811061147957611479615594565b602002602001018190525085818151811061149657611496615594565b602002602001015160a00151516001600160401b038111156114ba576114ba614b18565b6040519080825280602002602001820160405280156114e3578160200160208202803683370190505b508582815181106114f6576114f6615594565b60200260200101819052505f7f000000000000000000000000000000000000000000000000000000000000000087838151811061153557611535615594565b60200260200101516080015161154b9190615683565b905060604363ffffffff168263ffffffff1610156115935761158c89858a868151811061157a5761157a615594565b602002602001015160a00151856133d1565b90506115be565b6115bb89858a86815181106115aa576115aa615594565b602002602001015160a0015161291c565b90505b5f5b8884815181106115d2576115d2615594565b602002602001015160a001515181101561167d5761163f8985815181106115fb576115fb615594565b602002602001015160c00151828151811061161857611618615594565b602002602001015183838151811061163257611632615594565b60200260200101516134ff565b88858151811061165157611651615594565b6020026020010151828151811061166a5761166a615594565b60209081029190910101526001016115c0565b505050806001019050611331565b50505050915091565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146116dd57604051633213a66160e21b815260040160405180910390fd5b6116e683610fea565b15611060576001600160a01b038381165f908152609a602052604080822054905163152667d960e31b81529083166004820181905273beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac06024830152927f0000000000000000000000000000000000000000000000000000000000000000169063a9333ec890604401602060405180830381865afa15801561177d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117a19190615654565b6001600160a01b0386165f90815260a26020908152604080832073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0845282528083208151928301909152548152919250611807866117ff6001600160401b03808716908916613506565b84919061351a565b9050610fe1848873beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac084613538565b6040516394f649dd60e01b81526001600160a01b03828116600483015260609182915f9182917f000000000000000000000000000000000000000000000000000000000000000016906394f649dd906024015f60405180830381865afa158015611895573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526118bc91908101906156fa565b60405163fe243a1760e01b81526001600160a01b03888116600483015273beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac060248301529294509092505f917f0000000000000000000000000000000000000000000000000000000000000000169063fe243a1790604401602060405180830381865afa158015611942573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061196691906157b5565b9050805f0361197a57509094909350915050565b5f8351600161198991906157cc565b6001600160401b038111156119a0576119a0614b18565b6040519080825280602002602001820160405280156119c9578160200160208202803683370190505b5090505f845160016119db91906157cc565b6001600160401b038111156119f2576119f2614b18565b604051908082528060200260200182016040528015611a1b578160200160208202803683370190505b50905073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac082865181518110611a4657611a46615594565b60200260200101906001600160a01b031690816001600160a01b0316815250508281865181518110611a7a57611a7a615594565b60209081029190910101525f5b8551811015611b1b57858181518110611aa257611aa2615594565b6020026020010151838281518110611abc57611abc615594565b60200260200101906001600160a01b031690816001600160a01b031681525050848181518110611aee57611aee615594565b6020026020010151828281518110611b0857611b08615594565b6020908102919091010152600101611a87565b509097909650945050505050565b5f6001600160a01b03821615801590611b5b57506001600160a01b038083165f818152609a6020526040902054909116145b92915050565b60405163152667d960e31b81526001600160a01b03838116600483015282811660248301525f9182917f0000000000000000000000000000000000000000000000000000000000000000169063a9333ec890604401602060405180830381865afa158015611bd1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611bf59190615654565b9050611c038484835f6135b2565b949350505050565b611c1361366f565b6111125f6136c9565b82611c2681613327565b611c435760405163932d94f760e01b815260040160405180910390fd5b611c4c84611b29565b611c69576040516325ec6c1f60e01b815260040160405180910390fd5b836001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b67080908484604051610e95929190615626565b60605f82516001600160401b03811115611cc057611cc0614b18565b604051908082528060200260200182016040528015611ce9578160200160208202803683370190505b5090505f5b8351811015611d72576001600160a01b0385165f9081526098602052604081208551909190869084908110611d2557611d25615594565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f2054828281518110611d5f57611d5f615594565b6020908102919091010152600101611cee565b509392505050565b606654600290600490811603611da35760405163840a48d560e01b815260040160405180910390fd5b611dab61371a565b855f5b81811015611e3e57611e36898983818110611dcb57611dcb615594565b9050602002810190611ddd91906157df565b611de6906157f3565b888884818110611df857611df8615594565b9050602002810190611e0a91906155c6565b888886818110611e1c57611e1c615594565b9050602002016020810190611e3191906157fe565b613773565b600101611dae565b5050610fe1600160c955565b6060611e5533612246565b9050610d5a84848461254e565b6001600160a01b038083165f90815260a260209081526040808320938516835292815282822083519182019093529154825290610d5a90613bec565b60608082516001600160401b03811115611eba57611eba614b18565b604051908082528060200260200182016040528015611ee3578160200160208202803683370190505b50915082516001600160401b03811115611eff57611eff614b18565b604051908082528060200260200182016040528015611f28578160200160208202803683370190505b506001600160a01b038086165f908152609a6020526040812054929350911690611f5386838761291c565b90505f5b8551811015612120575f611f83878381518110611f7657611f76615594565b6020026020010151613c0b565b9050806001600160a01b031663fe243a1789898581518110611fa757611fa7615594565b60200260200101516040518363ffffffff1660e01b8152600401611fe19291906001600160a01b0392831681529116602082015260400190565b602060405180830381865afa158015611ffc573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061202091906157b5565b85838151811061203257612032615594565b6020026020010181815250505f60a25f8a6001600160a01b03166001600160a01b031681526020019081526020015f205f89858151811061207557612075615594565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f206040518060200160405290815f8201548152505090506120f98684815181106120c7576120c7615594565b60200260200101518585815181106120e1576120e1615594565b60200260200101518361351a9092919063ffffffff16565b87848151811061210b5761210b615594565b60209081029190910101525050600101611f57565b5050505b9250929050565b5f54610100900460ff161580801561214957505f54600160ff909116105b806121625750303b15801561216257505f5460ff166001145b6121ca5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff1916600117905580156121eb575f805461ff0019166101001790555b6121f482612f58565b6121fd836136c9565b8015611060575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b606061225182610fea565b61226e5760405163a5c7c44560e01b815260040160405180910390fd5b61227782611b29565b15612295576040516311ca333560e31b815260040160405180910390fd5b336001600160a01b0383161461234d576001600160a01b038083165f908152609a6020526040902054166122c881613327565b806122ee57506001600160a01b038181165f908152609960205260409020600101541633145b61230b57604051631e499a2360e11b815260040160405180910390fd5b806001600160a01b0316836001600160a01b03167ff0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a60405160405180910390a3505b611b5b82613c7d565b60665460029060049081160361237f5760405163840a48d560e01b815260040160405180910390fd5b61238761371a565b61239b612393866157f3565b858585613773565b6123a5600160c955565b5050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146123f5576040516323d871a560e01b815260040160405180910390fd5b6001600160a01b038085165f908152609860209081526040808320938716835292905290812054612433906001600160401b03808616908516613edc565b90505f612442868686866135b2565b90505f61244f82846157cc565b905061245d875f8886613538565b6001600160a01b03861673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac014610fe157604051633b9e9f0160e21b81526001600160a01b038781166004830152602482018390527f0000000000000000000000000000000000000000000000000000000000000000169063ee7a7c04906044015f604051808303815f87803b1580156124e8575f5ffd5b505af11580156124fa573d5f5f3e3d5ffd5b5050604080516001600160a01b038a81168252602082018690528b1693507feff6aab2bc3f7c648896e1522eae71d6c22e3b0e218206b3f40af0e4d204716b9250015b60405180910390a250505050505050565b61255733610fea565b1561257557604051633bf2b50360e11b815260040160405180910390fd5b61257e83611b29565b61259b576040516325ec6c1f60e01b815260040160405180910390fd5b6125a733848484613ef4565b6110603384613015565b60605f83516001600160401b038111156125cd576125cd614b18565b60405190808252806020026020018201604052801561260057816020015b60608152602001906001900390816125eb5790505b5090505f5b8451811015611d725761263185828151811061262357612623615594565b602002602001015185611ca4565b82828151811061264357612643615594565b6020908102919091010152600101612605565b61265e61366f565b6001600160a01b0381166126c35760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016121c1565b6126cc816136c9565b50565b5f7f0000000000000000000000000000000000000000000000000000000000000000461461278f5750604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b507f000000000000000000000000000000000000000000000000000000000000000090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612810573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128349190615819565b6001600160a01b0316336001600160a01b0316146128655760405163794821ff60e01b815260040160405180910390fd5b6066548019821981161461288c5760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020015b60405180910390a25050565b6001600160a01b0381165f90815260a360205260409020606090611b5b90613fb9565b5f6128f76126cf565b60405161190160f01b6020820152602281019190915260428101839052606201611126565b60605f82516001600160401b0381111561293857612938614b18565b604051908082528060200260200182016040528015612961578160200160208202803683370190505b5090505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663547afb8786866040518363ffffffff1660e01b81526004016129b3929190615834565b5f60405180830381865afa1580156129cd573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526129f49190810190615857565b90505f5b8451811015610b2157612a3e87868381518110612a1757612a17615594565b6020026020010151848481518110612a3157612a31615594565b6020026020010151613118565b838281518110612a5057612a50615594565b60209081029190910101526001016129f8565b5f6001600160a01b038616612a8b576040516339b190bb60e11b815260040160405180910390fd5b83515f03612aac5760405163796cc52560e01b815260040160405180910390fd5b5f84516001600160401b03811115612ac657612ac6614b18565b604051908082528060200260200182016040528015612aef578160200160208202803683370190505b5090505f85516001600160401b03811115612b0c57612b0c614b18565b604051908082528060200260200182016040528015612b35578160200160208202803683370190505b5090505f5b8651811015612d8b575f612b59888381518110611f7657611f76615594565b90505f60a25f8c6001600160a01b03166001600160a01b031681526020019081526020015f205f8a8581518110612b9257612b92615594565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f206040518060200160405290815f820154815250509050612bfe888481518110612be457612be4615594565b60200260200101518885815181106120e1576120e1615594565b848481518110612c1057612c10615594565b602002602001018181525050612c48888481518110612c3157612c31615594565b602002602001015182613fc590919063ffffffff16565b858481518110612c5a57612c5a615594565b60209081029190910101526001600160a01b038a1615612cef57612cb18a8a8581518110612c8a57612c8a615594565b6020026020010151878681518110612ca457612ca4615594565b6020026020010151613fd9565b612cef8a8c8b8681518110612cc857612cc8615594565b6020026020010151878781518110612ce257612ce2615594565b6020026020010151613538565b816001600160a01b031663724af4238c8b8681518110612d1157612d11615594565b60200260200101518b8781518110612d2b57612d2b615594565b60200260200101516040518463ffffffff1660e01b8152600401612d51939291906158e6565b5f604051808303815f87803b158015612d68575f5ffd5b505af1158015612d7a573d5f5f3e3d5ffd5b505050505050806001019050612b3a565b506001600160a01b0388165f908152609f60205260408120805491829190612db28361590a565b91905055505f6040518060e001604052808b6001600160a01b031681526020018a6001600160a01b031681526020018b6001600160a01b031681526020018381526020014363ffffffff1681526020018981526020018581525090505f612e1882611114565b5f818152609e602090815260408083208054600160ff19909116811790915560a4835292819020865181546001600160a01b03199081166001600160a01b039283161783558885015195830180548216968316969096179095559187015160028201805490951692169190911790925560608501516003830155608085015160048301805463ffffffff191663ffffffff90921691909117905560a085015180519394508593612ece926005850192019061490a565b5060c08201518051612eea91600684019160209091019061496d565b5050506001600160a01b038b165f90815260a360205260409020612f0e9082614067565b507f26b2aae26516e8719ef50ea2f6831a2efbd4e37dccdf0f6936b27bc08e793e30818386604051612f4293929190615922565b60405180910390a19a9950505050505050505050565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b5f611c0382612fad612fa687613bec565b8690614072565b90614072565b6001600160a01b038281165f8181526099602090815260409182902060010180546001600160a01b0319169486169485179055905192835290917f773b54c04d756fcc5e678111f7d730de3be98192000799eee3d63716055a87c691016128bf565b6066545f9060019081160361303d5760405163840a48d560e01b815260040160405180910390fd5b6001600160a01b038381165f818152609a602052604080822080546001600160a01b0319169487169485179055517fc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d87433049190a35f5f61309a85611829565b915091505f6130aa86868561291c565b90505f5b8351811015610fe15761311086888684815181106130ce576130ce615594565b60200260200101515f8786815181106130e9576130e9615594565b602002602001015187878151811061310357613103615594565b60200260200101516131fa565b6001016130ae565b5f73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeabf196001600160a01b038416016131ea5760405163a3d75e0960e01b81526001600160a01b0385811660048301525f917f00000000000000000000000000000000000000000000000000000000000000009091169063a3d75e0990602401602060405180830381865afa1580156131a6573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131ca9190615654565b90506131e26001600160401b03848116908316613506565b915050610d5a565b506001600160401b031692915050565b805f0361321a57604051630a33bc6960e21b815260040160405180910390fd5b6001600160a01b038086165f90815260a26020908152604080832093881683529290522061324a81858585614086565b6040805160208101909152815481527f8be932bac54561f27260f95463d9b8ab37e06b2842e5ee2404157cc13df6eb8f908790879061328890613bec565b604051613297939291906158e6565b60405180910390a16132a886610fea565b15610fe1576001600160a01b038088165f908152609860209081526040808320938916835292905290812080548592906132e39084906157cc565b92505081905550866001600160a01b03167f1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c87878660405161253d939291906158e6565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af11580156133ad573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b5b919061560b565b60605f83516001600160401b038111156133ed576133ed614b18565b604051908082528060200260200182016040528015613416578160200160208202803683370190505b5090505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166394d7d00c8787876040518463ffffffff1660e01b815260040161346a9392919061594c565b5f60405180830381865afa158015613484573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526134ab9190810190615857565b90505f5b85518110156134f3576134ce88878381518110612a1757612a17615594565b8382815181106134e0576134e0615594565b60209081029190910101526001016134af565b50909695505050505050565b5f610d5a83835b5f610d5a8383670de0b6b3a76400006140f5565b5f611c038261353261352b87613bec565b8690613506565b90613506565b6001600160a01b038085165f9081526098602090815260408083209386168352929052908120805483929061356e908490615985565b92505081905550836001600160a01b03167f6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd848484604051610e95939291906158e6565b6001600160a01b038085165f90815260a560209081526040808320938716835292905290812081906135e3906141da565b90505f61364960016136157f000000000000000000000000000000000000000000000000000000000000000043615998565b61361f9190615998565b6001600160a01b03808a165f90815260a560209081526040808320938c16835292905220906141f4565b90505f6136568284615985565b9050613663818787614210565b98975050505050505050565b6033546001600160a01b031633146111125760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016121c1565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b600260c9540361376c5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016121c1565b600260c955565b60a0840151518214613798576040516343714afd60e01b815260040160405180910390fd5b83604001516001600160a01b0316336001600160a01b0316146137ce576040516316110d3560e21b815260040160405180910390fd5b5f6137d885611114565b5f818152609e602052604090205490915060ff16613809576040516387c9d21960e01b815260040160405180910390fd5b60605f7f0000000000000000000000000000000000000000000000000000000000000000876080015161383c9190615683565b90508063ffffffff164363ffffffff161161386a576040516378f67ae160e11b815260040160405180910390fd5b613881875f015188602001518960a00151846133d1565b87516001600160a01b03165f90815260a3602052604090209092506138a791508361422e565b505f82815260a46020526040812080546001600160a01b031990811682556001820180548216905560028201805490911690556003810182905560048101805463ffffffff19169055906138fe60058301826149a6565b61390b600683015f6149a6565b50505f828152609e602052604090819020805460ff19169055517f1f40400889274ed07b24845e5054a87a0cab969eb1277aafe61ae352e7c32a00906139549084815260200190565b60405180910390a185516001600160a01b039081165f908152609a6020526040812054885160a08a0151919093169261398e91849061291c565b90505f5b8860a0015151811015613be1575f6139b98a60a001518381518110611f7657611f76615594565b90505f6139ef8b60c0015184815181106139d5576139d5615594565b602002602001015187858151811061163257611632615594565b90508715613abf57816001600160a01b0316632eae418c8c5f01518d60a001518681518110613a2057613a20615594565b60200260200101518d8d88818110613a3a57613a3a615594565b9050602002016020810190613a4f9190614e47565b60405160e085901b6001600160e01b03191681526001600160a01b03938416600482015291831660248301529091166044820152606481018490526084015f604051808303815f87803b158015613aa4575f5ffd5b505af1158015613ab6573d5f5f3e3d5ffd5b50505050613bd7565b5f5f836001600160a01b031663c4623ea18e5f01518f60a001518881518110613aea57613aea615594565b60200260200101518f8f8a818110613b0457613b04615594565b9050602002016020810190613b199190614e47565b60405160e085901b6001600160e01b03191681526001600160a01b039384166004820152918316602483015290911660448201526064810186905260840160408051808303815f875af1158015613b72573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b9691906159b4565b91509150613bd4878e5f01518f60a001518881518110613bb857613bb8615594565b602002602001015185858b8b8151811061310357613103615594565b50505b5050600101613992565b505050505050505050565b80515f9015613bfc578151611b5b565b670de0b6b3a764000092915050565b5f6001600160a01b03821673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac014613c56577f0000000000000000000000000000000000000000000000000000000000000000611b5b565b7f000000000000000000000000000000000000000000000000000000000000000092915050565b606654606090600190600290811603613ca95760405163840a48d560e01b815260040160405180910390fd5b6001600160a01b038084165f818152609a602052604080822080546001600160a01b0319811690915590519316928392917ffee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af4467691a35f5f613d0886611829565b9150915081515f03613d1c57505050613ed6565b81516001600160401b03811115613d3557613d35614b18565b604051908082528060200260200182016040528015613d5e578160200160208202803683370190505b5094505f613d6d87858561291c565b90505f5b8351811015613ed0576040805160018082528183019092525f916020808301908036833750506040805160018082528183019092529293505f9291506020808301908036833750506040805160018082528183019092529293505f92915060208083019080368337019050509050868481518110613df157613df1615594565b6020026020010151835f81518110613e0b57613e0b615594565b60200260200101906001600160a01b031690816001600160a01b031681525050858481518110613e3d57613e3d615594565b6020026020010151825f81518110613e5757613e57615594565b602002602001018181525050848481518110613e7557613e75615594565b6020026020010151815f81518110613e8f57613e8f615594565b602002602001018181525050613ea88b89858585612a63565b8a8581518110613eba57613eba615594565b6020908102919091010152505050600101613d71565b50505050505b50919050565b5f613eea8483856001614239565b611c039085615985565b6001600160a01b038084165f908152609960205260409020600101541680613f1c5750613fb3565b6001600160a01b0381165f908152609c6020908152604080832085845290915290205460ff1615613f6057604051630d4c4c9160e21b815260040160405180910390fd5b6001600160a01b0381165f908152609c602090815260408083208584528252909120805460ff191660011790558301516123a5908290613fa790889088908490889061081f565b85516020870151614294565b50505050565b60605f610d5a836142e6565b5f610d5a613fd284613bec565b8390613506565b6001600160a01b03821673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac014611060576001600160a01b038084165f90815260a560209081526040808320938616835292905290812061402c906141da565b9050613fb34361403c84846157cc565b6001600160a01b038088165f90815260a560209081526040808320938a16835292905220919061433f565b5f610d5a838361434a565b5f610d5a83670de0b6b3a7640000846140f5565b825f036140a65761409f670de0b6b3a764000082614072565b8455613fb3565b6040805160208101909152845481525f906140c290858461351a565b90505f6140cf84836157cc565b90505f6140ea84612fad6140e3888a6157cc565b8590614072565b875550505050505050565b5f80805f19858709858702925082811083820303915050805f0361412c57838281614122576141226159d6565b0492505050610d5a565b8084116141735760405162461bcd60e51b81526020600482015260156024820152744d6174683a206d756c446976206f766572666c6f7760581b60448201526064016121c1565b5f8486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091025f889003889004909101858311909403939093029303949094049190911702949350505050565b5f6141e58282614396565b6001600160e01b031692915050565b5f6142008383836143db565b6001600160e01b03169392505050565b5f611c0361421e83856159ea565b85906001600160401b0316613506565b5f610d5a8383614424565b5f5f6142468686866140f5565b9050600183600281111561425c5761425c615a09565b14801561427857505f8480614273576142736159d6565b868809115b1561428b576142886001826157cc565b90505b95945050505050565b428110156142b557604051630819bdcd60e01b815260040160405180910390fd5b6142c96001600160a01b0385168484614507565b613fb357604051638baa579f60e01b815260040160405180910390fd5b6060815f0180548060200260200160405190810160405280929190818152602001828054801561433357602002820191905f5260205f20905b81548152602001906001019080831161431f575b50505050509050919050565b61106083838361455b565b5f81815260018301602052604081205461438f57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155611b5b565b505f611b5b565b81545f9080156143d3576143bc846143af600184615985565b5f91825260209091200190565b5464010000000090046001600160e01b0316611c03565b509092915050565b82545f90816143ec86868385614661565b9050801561441a57614403866143af600184615985565b5464010000000090046001600160e01b031661089d565b5091949350505050565b5f81815260018301602052604081205480156144fe575f614446600183615985565b85549091505f9061445990600190615985565b90508181146144b8575f865f01828154811061447757614477615594565b905f5260205f200154905080875f01848154811061449757614497615594565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806144c9576144c9615a1d565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050611b5b565b5f915050611b5b565b5f5f5f61451485856146b4565b90925090505f81600481111561452c5761452c615a09565b14801561454a5750856001600160a01b0316826001600160a01b0316145b8061089d575061089d8686866146f3565b82548015614613575f614573856143af600185615985565b60408051808201909152905463ffffffff8082168084526401000000009092046001600160e01b0316602084015291925090851610156145c65760405163151b8e3f60e11b815260040160405180910390fd5b805163ffffffff80861691160361461157826145e7866143af600186615985565b80546001600160e01b03929092166401000000000263ffffffff9092169190911790555050505050565b505b506040805180820190915263ffffffff92831681526001600160e01b03918216602080830191825285546001810187555f968752952091519051909216640100000000029190921617910155565b5f5b81831015611d72575f61467684846147da565b5f8781526020902090915063ffffffff86169082015463ffffffff1611156146a0578092506146ae565b6146ab8160016157cc565b93505b50614663565b5f5f82516041036146e8576020830151604084015160608501515f1a6146dc878285856147f4565b94509450505050612124565b505f90506002612124565b5f5f5f856001600160a01b0316631626ba7e60e01b868660405160240161471b929190615a31565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516147599190615a6d565b5f60405180830381855afa9150503d805f8114614791576040519150601f19603f3d011682016040523d82523d5f602084013e614796565b606091505b50915091508180156147aa57506020815110155b801561089d57508051630b135d3f60e11b906147cf90830160209081019084016157b5565b149695505050505050565b5f6147e86002848418615a83565b610d5a908484166157cc565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561482957505f905060036148a8565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561487a573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166148a2575f600192509250506148a8565b91505f90505b94509492505050565b6040518060e001604052805f6001600160a01b031681526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f81526020015f63ffffffff16815260200160608152602001606081525090565b828054828255905f5260205f2090810192821561495d579160200282015b8281111561495d57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614928565b506149699291506149bd565b5090565b828054828255905f5260205f2090810192821561495d579160200282015b8281111561495d57825182559160200191906001019061498b565b5080545f8255905f5260205f20908101906126cc91905b5b80821115614969575f81556001016149be565b6001600160a01b03811681146126cc575f5ffd5b80356149f0816149d1565b919050565b5f5f5f5f5f60a08688031215614a09575f5ffd5b8535614a14816149d1565b94506020860135614a24816149d1565b93506040860135614a34816149d1565b94979396509394606081013594506080013592915050565b5f5f83601f840112614a5c575f5ffd5b5081356001600160401b03811115614a72575f5ffd5b6020830191508360208260051b8501011115612124575f5ffd5b5f5f60208385031215614a9d575f5ffd5b82356001600160401b03811115614ab2575f5ffd5b614abe85828601614a4c565b90969095509350505050565b602080825282518282018190525f918401906040840190835b81811015610b21578351835260209384019390920191600101614ae3565b5f60208284031215614b11575f5ffd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b60405160e081016001600160401b0381118282101715614b4e57614b4e614b18565b60405290565b604080519081016001600160401b0381118282101715614b4e57614b4e614b18565b604051601f8201601f191681016001600160401b0381118282101715614b9e57614b9e614b18565b604052919050565b5f6001600160401b03821115614bbe57614bbe614b18565b5060051b60200190565b5f82601f830112614bd7575f5ffd5b8135614bea614be582614ba6565b614b76565b8082825260208201915060208360051b860101925085831115614c0b575f5ffd5b602085015b83811015614c31578035614c23816149d1565b835260209283019201614c10565b5095945050505050565b5f82601f830112614c4a575f5ffd5b8135614c58614be582614ba6565b8082825260208201915060208360051b860101925085831115614c79575f5ffd5b602085015b83811015614c31578035835260209283019201614c7e565b5f5f5f60608486031215614ca8575f5ffd5b8335614cb3816149d1565b925060208401356001600160401b03811115614ccd575f5ffd5b614cd986828701614bc8565b92505060408401356001600160401b03811115614cf4575f5ffd5b614d0086828701614c3b565b9150509250925092565b5f8151808452602084019350602083015f5b82811015614d3a578151865260209586019590910190600101614d1c565b5093949350505050565b602081525f610d5a6020830184614d0a565b803563ffffffff811681146149f0575f5ffd5b5f5f83601f840112614d79575f5ffd5b5081356001600160401b03811115614d8f575f5ffd5b602083019150836020828501011115612124575f5ffd5b5f5f5f5f60608587031215614db9575f5ffd5b8435614dc4816149d1565b9350614dd260208601614d56565b925060408501356001600160401b03811115614dec575f5ffd5b614df887828801614d69565b95989497509550505050565b5f5f5f5f60808587031215614e17575f5ffd5b8435614e22816149d1565b93506020850135614e32816149d1565b93969395505050506040820135916060013590565b5f60208284031215614e57575f5ffd5b8135610d5a816149d1565b5f5f60408385031215614e73575f5ffd5b8235614e7e816149d1565b91506020830135614e8e816149d1565b809150509250929050565b5f60e08284031215614ea9575f5ffd5b614eb1614b2c565b9050614ebc826149e5565b8152614eca602083016149e5565b6020820152614edb604083016149e5565b604082015260608281013590820152614ef660808301614d56565b608082015260a08201356001600160401b03811115614f13575f5ffd5b614f1f84828501614bc8565b60a08301525060c08201356001600160401b03811115614f3d575f5ffd5b614f4984828501614c3b565b60c08301525092915050565b5f60208284031215614f65575f5ffd5b81356001600160401b03811115614f7a575f5ffd5b611c0384828501614e99565b5f60208284031215614f96575f5ffd5b813560ff81168114610d5a575f5ffd5b5f8151808452602084019350602083015f5b82811015614d3a5781516001600160a01b0316865260209586019590910190600101614fb8565b80516001600160a01b03908116835260208083015182169084015260408083015190911690830152606080820151908301526080808201515f9161502a9085018263ffffffff169052565b5060a082015160e060a085015261504460e0850182614fa6565b905060c083015184820360c086015261428b8282614d0a565b602081525f610d5a6020830184614fdf565b5f82825180855260208501945060208160051b830101602085015f5b838110156134f357601f198584030188526150a7838351614d0a565b602098890198909350919091019060010161508b565b5f604082016040835280855180835260608501915060608160051b8601019250602087015f5b8281101561511457605f198786030184526150ff858351614fdf565b945060209384019391909101906001016150e3565b50505050828103602084015261428b818561506f565b6001600160401b03811681146126cc575f5ffd5b5f5f5f60608486031215615150575f5ffd5b833561515b816149d1565b92506020840135915060408401356151728161512a565b809150509250925092565b604081525f61518f6040830185614fa6565b828103602084015261428b8185614d0a565b5f5f5f604084860312156151b3575f5ffd5b83356151be816149d1565b925060208401356001600160401b038111156151d8575f5ffd5b6151e486828701614d69565b9497909650939450505050565b5f5f60408385031215615202575f5ffd5b823561520d816149d1565b915060208301356001600160401b03811115615227575f5ffd5b61523385828601614bc8565b9150509250929050565b5f5f5f5f5f5f60608789031215615252575f5ffd5b86356001600160401b03811115615267575f5ffd5b61527389828a01614a4c565b90975095505060208701356001600160401b03811115615291575f5ffd5b61529d89828a01614a4c565b90955093505060408701356001600160401b038111156152bb575f5ffd5b6152c789828a01614a4c565b979a9699509497509295939492505050565b5f5f5f606084860312156152eb575f5ffd5b83356152f6816149d1565b925060208401356001600160401b03811115615310575f5ffd5b840160408187031215615321575f5ffd5b615329614b54565b81356001600160401b0381111561533e575f5ffd5b8201601f8101881361534e575f5ffd5b80356001600160401b0381111561536757615367614b18565b61537a601f8201601f1916602001614b76565b81815289602083850101111561538e575f5ffd5b816020840160208301375f60209282018301528352928301359282019290925293969395505050506040919091013590565b5f5f604083850312156153d1575f5ffd5b82356153dc816149d1565b946020939093013593505050565b604081525f61518f6040830185614d0a565b80151581146126cc575f5ffd5b5f5f5f5f6060858703121561541c575f5ffd5b84356001600160401b03811115615431575f5ffd5b850160e08188031215615442575f5ffd5b935060208501356001600160401b0381111561545c575f5ffd5b61546887828801614a4c565b909450925050604085013561547c816153fc565b939692955090935050565b5f5f5f5f6080858703121561549a575f5ffd5b84356154a5816149d1565b935060208501356154b5816149d1565b925060408501356154c58161512a565b9150606085013561547c8161512a565b5f5f604083850312156154e6575f5ffd5b82356001600160401b038111156154fb575f5ffd5b8301601f8101851361550b575f5ffd5b8035615519614be582614ba6565b8082825260208201915060208360051b85010192508783111561553a575f5ffd5b6020840193505b82841015615565578335615554816149d1565b825260209384019390910190615541565b945050505060208301356001600160401b03811115615227575f5ffd5b602081525f610d5a602083018461506f565b634e487b7160e01b5f52603260045260245ffd5b5f8235605e198336030181126155bc575f5ffd5b9190910192915050565b5f5f8335601e198436030181126155db575f5ffd5b8301803591506001600160401b038211156155f4575f5ffd5b6020019150600581901b3603821315612124575f5ffd5b5f6020828403121561561b575f5ffd5b8151610d5a816153fc565b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b5f60208284031215615664575f5ffd5b8151610d5a8161512a565b634e487b7160e01b5f52601160045260245ffd5b63ffffffff8181168382160190811115611b5b57611b5b61566f565b5f82601f8301126156ae575f5ffd5b81516156bc614be582614ba6565b8082825260208201915060208360051b8601019250858311156156dd575f5ffd5b602085015b83811015614c315780518352602092830192016156e2565b5f5f6040838503121561570b575f5ffd5b82516001600160401b03811115615720575f5ffd5b8301601f81018513615730575f5ffd5b805161573e614be582614ba6565b8082825260208201915060208360051b85010192508783111561575f575f5ffd5b6020840193505b8284101561578a578351615779816149d1565b825260209384019390910190615766565b8095505050505060208301516001600160401b038111156157a9575f5ffd5b6152338582860161569f565b5f602082840312156157c5575f5ffd5b5051919050565b80820180821115611b5b57611b5b61566f565b5f823560de198336030181126155bc575f5ffd5b5f611b5b3683614e99565b5f6020828403121561580e575f5ffd5b8135610d5a816153fc565b5f60208284031215615829575f5ffd5b8151610d5a816149d1565b6001600160a01b03831681526040602082018190525f90611c0390830184614fa6565b5f60208284031215615867575f5ffd5b81516001600160401b0381111561587c575f5ffd5b8201601f8101841361588c575f5ffd5b805161589a614be582614ba6565b8082825260208201915060208360051b8501019250868311156158bb575f5ffd5b6020840193505b8284101561089d5783516158d58161512a565b8252602093840193909101906158c2565b6001600160a01b039384168152919092166020820152604081019190915260600190565b5f6001820161591b5761591b61566f565b5060010190565b838152606060208201525f61593a6060830185614fdf565b828103604084015261089d8185614d0a565b6001600160a01b03841681526060602082018190525f9061596f90830185614fa6565b905063ffffffff83166040830152949350505050565b81810381811115611b5b57611b5b61566f565b63ffffffff8281168282160390811115611b5b57611b5b61566f565b5f5f604083850312156159c5575f5ffd5b505080516020909101519092909150565b634e487b7160e01b5f52601260045260245ffd5b6001600160401b038281168282160390811115611b5b57611b5b61566f565b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52603160045260245ffd5b828152604060208201525f82518060408401528060208501606085015e5f606082850101526060601f19601f8301168401019150509392505050565b5f82518060208501845e5f920191825250919050565b5f82615a9d57634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220941e9e6e8593291e8d6ceb84ad7c824f3ae57f6b4b3f6b1f5510b2c2e26719a864736f6c634300081b0033",
}

// DelegationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DelegationManagerMetaData.ABI instead.
var DelegationManagerABI = DelegationManagerMetaData.ABI

// DelegationManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DelegationManagerMetaData.Bin instead.
var DelegationManagerBin = DelegationManagerMetaData.Bin

// DeployDelegationManager deploys a new Ethereum contract, binding an instance of DelegationManager to it.
func DeployDelegationManager(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _eigenPodManager common.Address, _allocationManager common.Address, _pauserRegistry common.Address, _permissionController common.Address, _MIN_WITHDRAWAL_DELAY uint32) (common.Address, *types.Transaction, *DelegationManager, error) {
	parsed, err := DelegationManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DelegationManagerBin), backend, _strategyManager, _eigenPodManager, _allocationManager, _pauserRegistry, _permissionController, _MIN_WITHDRAWAL_DELAY)
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

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address approver, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_DelegationManager *DelegationManagerCaller) CalculateDelegationApprovalDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, approver common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "calculateDelegationApprovalDigestHash", staker, operator, approver, approverSalt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address approver, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_DelegationManager *DelegationManagerSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, approver common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _DelegationManager.Contract.CalculateDelegationApprovalDigestHash(&_DelegationManager.CallOpts, staker, operator, approver, approverSalt, expiry)
}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address approver, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_DelegationManager *DelegationManagerCallerSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, approver common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _DelegationManager.Contract.CalculateDelegationApprovalDigestHash(&_DelegationManager.CallOpts, staker, operator, approver, approverSalt, expiry)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_DelegationManager *DelegationManagerCaller) CalculateWithdrawalRoot(opts *bind.CallOpts, withdrawal IDelegationManagerTypesWithdrawal) ([32]byte, error) {
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
func (_DelegationManager *DelegationManagerSession) CalculateWithdrawalRoot(withdrawal IDelegationManagerTypesWithdrawal) ([32]byte, error) {
	return _DelegationManager.Contract.CalculateWithdrawalRoot(&_DelegationManager.CallOpts, withdrawal)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_DelegationManager *DelegationManagerCallerSession) CalculateWithdrawalRoot(withdrawal IDelegationManagerTypesWithdrawal) ([32]byte, error) {
	return _DelegationManager.Contract.CalculateWithdrawalRoot(&_DelegationManager.CallOpts, withdrawal)
}

// ConvertToDepositShares is a free data retrieval call binding the contract method 0x25df922e.
//
// Solidity: function convertToDepositShares(address staker, address[] strategies, uint256[] withdrawableShares) view returns(uint256[])
func (_DelegationManager *DelegationManagerCaller) ConvertToDepositShares(opts *bind.CallOpts, staker common.Address, strategies []common.Address, withdrawableShares []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "convertToDepositShares", staker, strategies, withdrawableShares)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ConvertToDepositShares is a free data retrieval call binding the contract method 0x25df922e.
//
// Solidity: function convertToDepositShares(address staker, address[] strategies, uint256[] withdrawableShares) view returns(uint256[])
func (_DelegationManager *DelegationManagerSession) ConvertToDepositShares(staker common.Address, strategies []common.Address, withdrawableShares []*big.Int) ([]*big.Int, error) {
	return _DelegationManager.Contract.ConvertToDepositShares(&_DelegationManager.CallOpts, staker, strategies, withdrawableShares)
}

// ConvertToDepositShares is a free data retrieval call binding the contract method 0x25df922e.
//
// Solidity: function convertToDepositShares(address staker, address[] strategies, uint256[] withdrawableShares) view returns(uint256[])
func (_DelegationManager *DelegationManagerCallerSession) ConvertToDepositShares(staker common.Address, strategies []common.Address, withdrawableShares []*big.Int) ([]*big.Int, error) {
	return _DelegationManager.Contract.ConvertToDepositShares(&_DelegationManager.CallOpts, staker, strategies, withdrawableShares)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address staker) view returns(uint256 totalQueued)
func (_DelegationManager *DelegationManagerCaller) CumulativeWithdrawalsQueued(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "cumulativeWithdrawalsQueued", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address staker) view returns(uint256 totalQueued)
func (_DelegationManager *DelegationManagerSession) CumulativeWithdrawalsQueued(staker common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.CumulativeWithdrawalsQueued(&_DelegationManager.CallOpts, staker)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address staker) view returns(uint256 totalQueued)
func (_DelegationManager *DelegationManagerCallerSession) CumulativeWithdrawalsQueued(staker common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.CumulativeWithdrawalsQueued(&_DelegationManager.CallOpts, staker)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address staker) view returns(address operator)
func (_DelegationManager *DelegationManagerCaller) DelegatedTo(opts *bind.CallOpts, staker common.Address) (common.Address, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "delegatedTo", staker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address staker) view returns(address operator)
func (_DelegationManager *DelegationManagerSession) DelegatedTo(staker common.Address) (common.Address, error) {
	return _DelegationManager.Contract.DelegatedTo(&_DelegationManager.CallOpts, staker)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address staker) view returns(address operator)
func (_DelegationManager *DelegationManagerCallerSession) DelegatedTo(staker common.Address) (common.Address, error) {
	return _DelegationManager.Contract.DelegatedTo(&_DelegationManager.CallOpts, staker)
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
// Solidity: function delegationApproverSaltIsSpent(address delegationApprover, bytes32 salt) view returns(bool spent)
func (_DelegationManager *DelegationManagerCaller) DelegationApproverSaltIsSpent(opts *bind.CallOpts, delegationApprover common.Address, salt [32]byte) (bool, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "delegationApproverSaltIsSpent", delegationApprover, salt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address delegationApprover, bytes32 salt) view returns(bool spent)
func (_DelegationManager *DelegationManagerSession) DelegationApproverSaltIsSpent(delegationApprover common.Address, salt [32]byte) (bool, error) {
	return _DelegationManager.Contract.DelegationApproverSaltIsSpent(&_DelegationManager.CallOpts, delegationApprover, salt)
}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address delegationApprover, bytes32 salt) view returns(bool spent)
func (_DelegationManager *DelegationManagerCallerSession) DelegationApproverSaltIsSpent(delegationApprover common.Address, salt [32]byte) (bool, error) {
	return _DelegationManager.Contract.DelegationApproverSaltIsSpent(&_DelegationManager.CallOpts, delegationApprover, salt)
}

// DepositScalingFactor is a free data retrieval call binding the contract method 0xbfae3fd2.
//
// Solidity: function depositScalingFactor(address staker, address strategy) view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) DepositScalingFactor(opts *bind.CallOpts, staker common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "depositScalingFactor", staker, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositScalingFactor is a free data retrieval call binding the contract method 0xbfae3fd2.
//
// Solidity: function depositScalingFactor(address staker, address strategy) view returns(uint256)
func (_DelegationManager *DelegationManagerSession) DepositScalingFactor(staker common.Address, strategy common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.DepositScalingFactor(&_DelegationManager.CallOpts, staker, strategy)
}

// DepositScalingFactor is a free data retrieval call binding the contract method 0xbfae3fd2.
//
// Solidity: function depositScalingFactor(address staker, address strategy) view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) DepositScalingFactor(staker common.Address, strategy common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.DepositScalingFactor(&_DelegationManager.CallOpts, staker, strategy)
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

// GetDepositedShares is a free data retrieval call binding the contract method 0x66d5ba93.
//
// Solidity: function getDepositedShares(address staker) view returns(address[], uint256[])
func (_DelegationManager *DelegationManagerCaller) GetDepositedShares(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "getDepositedShares", staker)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDepositedShares is a free data retrieval call binding the contract method 0x66d5ba93.
//
// Solidity: function getDepositedShares(address staker) view returns(address[], uint256[])
func (_DelegationManager *DelegationManagerSession) GetDepositedShares(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _DelegationManager.Contract.GetDepositedShares(&_DelegationManager.CallOpts, staker)
}

// GetDepositedShares is a free data retrieval call binding the contract method 0x66d5ba93.
//
// Solidity: function getDepositedShares(address staker) view returns(address[], uint256[])
func (_DelegationManager *DelegationManagerCallerSession) GetDepositedShares(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _DelegationManager.Contract.GetDepositedShares(&_DelegationManager.CallOpts, staker)
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

// GetOperatorsShares is a free data retrieval call binding the contract method 0xf0e0e676.
//
// Solidity: function getOperatorsShares(address[] operators, address[] strategies) view returns(uint256[][])
func (_DelegationManager *DelegationManagerCaller) GetOperatorsShares(opts *bind.CallOpts, operators []common.Address, strategies []common.Address) ([][]*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "getOperatorsShares", operators, strategies)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetOperatorsShares is a free data retrieval call binding the contract method 0xf0e0e676.
//
// Solidity: function getOperatorsShares(address[] operators, address[] strategies) view returns(uint256[][])
func (_DelegationManager *DelegationManagerSession) GetOperatorsShares(operators []common.Address, strategies []common.Address) ([][]*big.Int, error) {
	return _DelegationManager.Contract.GetOperatorsShares(&_DelegationManager.CallOpts, operators, strategies)
}

// GetOperatorsShares is a free data retrieval call binding the contract method 0xf0e0e676.
//
// Solidity: function getOperatorsShares(address[] operators, address[] strategies) view returns(uint256[][])
func (_DelegationManager *DelegationManagerCallerSession) GetOperatorsShares(operators []common.Address, strategies []common.Address) ([][]*big.Int, error) {
	return _DelegationManager.Contract.GetOperatorsShares(&_DelegationManager.CallOpts, operators, strategies)
}

// GetQueuedWithdrawal is a free data retrieval call binding the contract method 0x5d975e88.
//
// Solidity: function getQueuedWithdrawal(bytes32 withdrawalRoot) view returns((address,address,address,uint256,uint32,address[],uint256[]))
func (_DelegationManager *DelegationManagerCaller) GetQueuedWithdrawal(opts *bind.CallOpts, withdrawalRoot [32]byte) (IDelegationManagerTypesWithdrawal, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "getQueuedWithdrawal", withdrawalRoot)

	if err != nil {
		return *new(IDelegationManagerTypesWithdrawal), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelegationManagerTypesWithdrawal)).(*IDelegationManagerTypesWithdrawal)

	return out0, err

}

// GetQueuedWithdrawal is a free data retrieval call binding the contract method 0x5d975e88.
//
// Solidity: function getQueuedWithdrawal(bytes32 withdrawalRoot) view returns((address,address,address,uint256,uint32,address[],uint256[]))
func (_DelegationManager *DelegationManagerSession) GetQueuedWithdrawal(withdrawalRoot [32]byte) (IDelegationManagerTypesWithdrawal, error) {
	return _DelegationManager.Contract.GetQueuedWithdrawal(&_DelegationManager.CallOpts, withdrawalRoot)
}

// GetQueuedWithdrawal is a free data retrieval call binding the contract method 0x5d975e88.
//
// Solidity: function getQueuedWithdrawal(bytes32 withdrawalRoot) view returns((address,address,address,uint256,uint32,address[],uint256[]))
func (_DelegationManager *DelegationManagerCallerSession) GetQueuedWithdrawal(withdrawalRoot [32]byte) (IDelegationManagerTypesWithdrawal, error) {
	return _DelegationManager.Contract.GetQueuedWithdrawal(&_DelegationManager.CallOpts, withdrawalRoot)
}

// GetQueuedWithdrawalRoots is a free data retrieval call binding the contract method 0xfd8aa88d.
//
// Solidity: function getQueuedWithdrawalRoots(address staker) view returns(bytes32[])
func (_DelegationManager *DelegationManagerCaller) GetQueuedWithdrawalRoots(opts *bind.CallOpts, staker common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "getQueuedWithdrawalRoots", staker)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetQueuedWithdrawalRoots is a free data retrieval call binding the contract method 0xfd8aa88d.
//
// Solidity: function getQueuedWithdrawalRoots(address staker) view returns(bytes32[])
func (_DelegationManager *DelegationManagerSession) GetQueuedWithdrawalRoots(staker common.Address) ([][32]byte, error) {
	return _DelegationManager.Contract.GetQueuedWithdrawalRoots(&_DelegationManager.CallOpts, staker)
}

// GetQueuedWithdrawalRoots is a free data retrieval call binding the contract method 0xfd8aa88d.
//
// Solidity: function getQueuedWithdrawalRoots(address staker) view returns(bytes32[])
func (_DelegationManager *DelegationManagerCallerSession) GetQueuedWithdrawalRoots(staker common.Address) ([][32]byte, error) {
	return _DelegationManager.Contract.GetQueuedWithdrawalRoots(&_DelegationManager.CallOpts, staker)
}

// GetQueuedWithdrawals is a free data retrieval call binding the contract method 0x5dd68579.
//
// Solidity: function getQueuedWithdrawals(address staker) view returns((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, uint256[][] shares)
func (_DelegationManager *DelegationManagerCaller) GetQueuedWithdrawals(opts *bind.CallOpts, staker common.Address) (struct {
	Withdrawals []IDelegationManagerTypesWithdrawal
	Shares      [][]*big.Int
}, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "getQueuedWithdrawals", staker)

	outstruct := new(struct {
		Withdrawals []IDelegationManagerTypesWithdrawal
		Shares      [][]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Withdrawals = *abi.ConvertType(out[0], new([]IDelegationManagerTypesWithdrawal)).(*[]IDelegationManagerTypesWithdrawal)
	outstruct.Shares = *abi.ConvertType(out[1], new([][]*big.Int)).(*[][]*big.Int)

	return *outstruct, err

}

// GetQueuedWithdrawals is a free data retrieval call binding the contract method 0x5dd68579.
//
// Solidity: function getQueuedWithdrawals(address staker) view returns((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, uint256[][] shares)
func (_DelegationManager *DelegationManagerSession) GetQueuedWithdrawals(staker common.Address) (struct {
	Withdrawals []IDelegationManagerTypesWithdrawal
	Shares      [][]*big.Int
}, error) {
	return _DelegationManager.Contract.GetQueuedWithdrawals(&_DelegationManager.CallOpts, staker)
}

// GetQueuedWithdrawals is a free data retrieval call binding the contract method 0x5dd68579.
//
// Solidity: function getQueuedWithdrawals(address staker) view returns((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, uint256[][] shares)
func (_DelegationManager *DelegationManagerCallerSession) GetQueuedWithdrawals(staker common.Address) (struct {
	Withdrawals []IDelegationManagerTypesWithdrawal
	Shares      [][]*big.Int
}, error) {
	return _DelegationManager.Contract.GetQueuedWithdrawals(&_DelegationManager.CallOpts, staker)
}

// GetSlashableSharesInQueue is a free data retrieval call binding the contract method 0x6e174448.
//
// Solidity: function getSlashableSharesInQueue(address operator, address strategy) view returns(uint256)
func (_DelegationManager *DelegationManagerCaller) GetSlashableSharesInQueue(opts *bind.CallOpts, operator common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "getSlashableSharesInQueue", operator, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlashableSharesInQueue is a free data retrieval call binding the contract method 0x6e174448.
//
// Solidity: function getSlashableSharesInQueue(address operator, address strategy) view returns(uint256)
func (_DelegationManager *DelegationManagerSession) GetSlashableSharesInQueue(operator common.Address, strategy common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.GetSlashableSharesInQueue(&_DelegationManager.CallOpts, operator, strategy)
}

// GetSlashableSharesInQueue is a free data retrieval call binding the contract method 0x6e174448.
//
// Solidity: function getSlashableSharesInQueue(address operator, address strategy) view returns(uint256)
func (_DelegationManager *DelegationManagerCallerSession) GetSlashableSharesInQueue(operator common.Address, strategy common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.GetSlashableSharesInQueue(&_DelegationManager.CallOpts, operator, strategy)
}

// GetWithdrawableShares is a free data retrieval call binding the contract method 0xc978f7ac.
//
// Solidity: function getWithdrawableShares(address staker, address[] strategies) view returns(uint256[] withdrawableShares, uint256[] depositShares)
func (_DelegationManager *DelegationManagerCaller) GetWithdrawableShares(opts *bind.CallOpts, staker common.Address, strategies []common.Address) (struct {
	WithdrawableShares []*big.Int
	DepositShares      []*big.Int
}, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "getWithdrawableShares", staker, strategies)

	outstruct := new(struct {
		WithdrawableShares []*big.Int
		DepositShares      []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WithdrawableShares = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.DepositShares = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetWithdrawableShares is a free data retrieval call binding the contract method 0xc978f7ac.
//
// Solidity: function getWithdrawableShares(address staker, address[] strategies) view returns(uint256[] withdrawableShares, uint256[] depositShares)
func (_DelegationManager *DelegationManagerSession) GetWithdrawableShares(staker common.Address, strategies []common.Address) (struct {
	WithdrawableShares []*big.Int
	DepositShares      []*big.Int
}, error) {
	return _DelegationManager.Contract.GetWithdrawableShares(&_DelegationManager.CallOpts, staker, strategies)
}

// GetWithdrawableShares is a free data retrieval call binding the contract method 0xc978f7ac.
//
// Solidity: function getWithdrawableShares(address staker, address[] strategies) view returns(uint256[] withdrawableShares, uint256[] depositShares)
func (_DelegationManager *DelegationManagerCallerSession) GetWithdrawableShares(staker common.Address, strategies []common.Address) (struct {
	WithdrawableShares []*big.Int
	DepositShares      []*big.Int
}, error) {
	return _DelegationManager.Contract.GetWithdrawableShares(&_DelegationManager.CallOpts, staker, strategies)
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
// Solidity: function minWithdrawalDelayBlocks() view returns(uint32)
func (_DelegationManager *DelegationManagerCaller) MinWithdrawalDelayBlocks(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "minWithdrawalDelayBlocks")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MinWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0xc448feb8.
//
// Solidity: function minWithdrawalDelayBlocks() view returns(uint32)
func (_DelegationManager *DelegationManagerSession) MinWithdrawalDelayBlocks() (uint32, error) {
	return _DelegationManager.Contract.MinWithdrawalDelayBlocks(&_DelegationManager.CallOpts)
}

// MinWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0xc448feb8.
//
// Solidity: function minWithdrawalDelayBlocks() view returns(uint32)
func (_DelegationManager *DelegationManagerCallerSession) MinWithdrawalDelayBlocks() (uint32, error) {
	return _DelegationManager.Contract.MinWithdrawalDelayBlocks(&_DelegationManager.CallOpts)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address operator, address strategy) view returns(uint256 shares)
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
// Solidity: function operatorShares(address operator, address strategy) view returns(uint256 shares)
func (_DelegationManager *DelegationManagerSession) OperatorShares(operator common.Address, strategy common.Address) (*big.Int, error) {
	return _DelegationManager.Contract.OperatorShares(&_DelegationManager.CallOpts, operator, strategy)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address operator, address strategy) view returns(uint256 shares)
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
// Solidity: function pendingWithdrawals(bytes32 withdrawalRoot) view returns(bool pending)
func (_DelegationManager *DelegationManagerCaller) PendingWithdrawals(opts *bind.CallOpts, withdrawalRoot [32]byte) (bool, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "pendingWithdrawals", withdrawalRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 withdrawalRoot) view returns(bool pending)
func (_DelegationManager *DelegationManagerSession) PendingWithdrawals(withdrawalRoot [32]byte) (bool, error) {
	return _DelegationManager.Contract.PendingWithdrawals(&_DelegationManager.CallOpts, withdrawalRoot)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 withdrawalRoot) view returns(bool pending)
func (_DelegationManager *DelegationManagerCallerSession) PendingWithdrawals(withdrawalRoot [32]byte) (bool, error) {
	return _DelegationManager.Contract.PendingWithdrawals(&_DelegationManager.CallOpts, withdrawalRoot)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_DelegationManager *DelegationManagerCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_DelegationManager *DelegationManagerSession) PermissionController() (common.Address, error) {
	return _DelegationManager.Contract.PermissionController(&_DelegationManager.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_DelegationManager *DelegationManagerCallerSession) PermissionController() (common.Address, error) {
	return _DelegationManager.Contract.PermissionController(&_DelegationManager.CallOpts)
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

// BurnOperatorShares is a paid mutator transaction binding the contract method 0xee74937f.
//
// Solidity: function burnOperatorShares(address operator, address strategy, uint64 prevMaxMagnitude, uint64 newMaxMagnitude) returns()
func (_DelegationManager *DelegationManagerTransactor) BurnOperatorShares(opts *bind.TransactOpts, operator common.Address, strategy common.Address, prevMaxMagnitude uint64, newMaxMagnitude uint64) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "burnOperatorShares", operator, strategy, prevMaxMagnitude, newMaxMagnitude)
}

// BurnOperatorShares is a paid mutator transaction binding the contract method 0xee74937f.
//
// Solidity: function burnOperatorShares(address operator, address strategy, uint64 prevMaxMagnitude, uint64 newMaxMagnitude) returns()
func (_DelegationManager *DelegationManagerSession) BurnOperatorShares(operator common.Address, strategy common.Address, prevMaxMagnitude uint64, newMaxMagnitude uint64) (*types.Transaction, error) {
	return _DelegationManager.Contract.BurnOperatorShares(&_DelegationManager.TransactOpts, operator, strategy, prevMaxMagnitude, newMaxMagnitude)
}

// BurnOperatorShares is a paid mutator transaction binding the contract method 0xee74937f.
//
// Solidity: function burnOperatorShares(address operator, address strategy, uint64 prevMaxMagnitude, uint64 newMaxMagnitude) returns()
func (_DelegationManager *DelegationManagerTransactorSession) BurnOperatorShares(operator common.Address, strategy common.Address, prevMaxMagnitude uint64, newMaxMagnitude uint64) (*types.Transaction, error) {
	return _DelegationManager.Contract.BurnOperatorShares(&_DelegationManager.TransactOpts, operator, strategy, prevMaxMagnitude, newMaxMagnitude)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0xe4cc3f90.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, bool receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerTransactor) CompleteQueuedWithdrawal(opts *bind.TransactOpts, withdrawal IDelegationManagerTypesWithdrawal, tokens []common.Address, receiveAsTokens bool) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "completeQueuedWithdrawal", withdrawal, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0xe4cc3f90.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, bool receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerSession) CompleteQueuedWithdrawal(withdrawal IDelegationManagerTypesWithdrawal, tokens []common.Address, receiveAsTokens bool) (*types.Transaction, error) {
	return _DelegationManager.Contract.CompleteQueuedWithdrawal(&_DelegationManager.TransactOpts, withdrawal, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0xe4cc3f90.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, bool receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerTransactorSession) CompleteQueuedWithdrawal(withdrawal IDelegationManagerTypesWithdrawal, tokens []common.Address, receiveAsTokens bool) (*types.Transaction, error) {
	return _DelegationManager.Contract.CompleteQueuedWithdrawal(&_DelegationManager.TransactOpts, withdrawal, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x9435bb43.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerTransactor) CompleteQueuedWithdrawals(opts *bind.TransactOpts, withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "completeQueuedWithdrawals", withdrawals, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x9435bb43.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerSession) CompleteQueuedWithdrawals(withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _DelegationManager.Contract.CompleteQueuedWithdrawals(&_DelegationManager.TransactOpts, withdrawals, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x9435bb43.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_DelegationManager *DelegationManagerTransactorSession) CompleteQueuedWithdrawals(withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _DelegationManager.Contract.CompleteQueuedWithdrawals(&_DelegationManager.TransactOpts, withdrawals, tokens, receiveAsTokens)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x60a0d1ce.
//
// Solidity: function decreaseDelegatedShares(address staker, uint256 curDepositShares, uint64 beaconChainSlashingFactorDecrease) returns()
func (_DelegationManager *DelegationManagerTransactor) DecreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, curDepositShares *big.Int, beaconChainSlashingFactorDecrease uint64) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "decreaseDelegatedShares", staker, curDepositShares, beaconChainSlashingFactorDecrease)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x60a0d1ce.
//
// Solidity: function decreaseDelegatedShares(address staker, uint256 curDepositShares, uint64 beaconChainSlashingFactorDecrease) returns()
func (_DelegationManager *DelegationManagerSession) DecreaseDelegatedShares(staker common.Address, curDepositShares *big.Int, beaconChainSlashingFactorDecrease uint64) (*types.Transaction, error) {
	return _DelegationManager.Contract.DecreaseDelegatedShares(&_DelegationManager.TransactOpts, staker, curDepositShares, beaconChainSlashingFactorDecrease)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x60a0d1ce.
//
// Solidity: function decreaseDelegatedShares(address staker, uint256 curDepositShares, uint64 beaconChainSlashingFactorDecrease) returns()
func (_DelegationManager *DelegationManagerTransactorSession) DecreaseDelegatedShares(staker common.Address, curDepositShares *big.Int, beaconChainSlashingFactorDecrease uint64) (*types.Transaction, error) {
	return _DelegationManager.Contract.DecreaseDelegatedShares(&_DelegationManager.TransactOpts, staker, curDepositShares, beaconChainSlashingFactorDecrease)
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

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x3c651cf2.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 prevDepositShares, uint256 addedShares) returns()
func (_DelegationManager *DelegationManagerTransactor) IncreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, prevDepositShares *big.Int, addedShares *big.Int) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "increaseDelegatedShares", staker, strategy, prevDepositShares, addedShares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x3c651cf2.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 prevDepositShares, uint256 addedShares) returns()
func (_DelegationManager *DelegationManagerSession) IncreaseDelegatedShares(staker common.Address, strategy common.Address, prevDepositShares *big.Int, addedShares *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.IncreaseDelegatedShares(&_DelegationManager.TransactOpts, staker, strategy, prevDepositShares, addedShares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x3c651cf2.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 prevDepositShares, uint256 addedShares) returns()
func (_DelegationManager *DelegationManagerTransactorSession) IncreaseDelegatedShares(staker common.Address, strategy common.Address, prevDepositShares *big.Int, addedShares *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.IncreaseDelegatedShares(&_DelegationManager.TransactOpts, staker, strategy, prevDepositShares, addedShares)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_DelegationManager *DelegationManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "initialize", initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_DelegationManager *DelegationManagerSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.Initialize(&_DelegationManager.TransactOpts, initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_DelegationManager *DelegationManagerTransactorSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.Initialize(&_DelegationManager.TransactOpts, initialOwner, initialPausedStatus)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x54b7c96c.
//
// Solidity: function modifyOperatorDetails(address operator, address newDelegationApprover) returns()
func (_DelegationManager *DelegationManagerTransactor) ModifyOperatorDetails(opts *bind.TransactOpts, operator common.Address, newDelegationApprover common.Address) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "modifyOperatorDetails", operator, newDelegationApprover)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x54b7c96c.
//
// Solidity: function modifyOperatorDetails(address operator, address newDelegationApprover) returns()
func (_DelegationManager *DelegationManagerSession) ModifyOperatorDetails(operator common.Address, newDelegationApprover common.Address) (*types.Transaction, error) {
	return _DelegationManager.Contract.ModifyOperatorDetails(&_DelegationManager.TransactOpts, operator, newDelegationApprover)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x54b7c96c.
//
// Solidity: function modifyOperatorDetails(address operator, address newDelegationApprover) returns()
func (_DelegationManager *DelegationManagerTransactorSession) ModifyOperatorDetails(operator common.Address, newDelegationApprover common.Address) (*types.Transaction, error) {
	return _DelegationManager.Contract.ModifyOperatorDetails(&_DelegationManager.TransactOpts, operator, newDelegationApprover)
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
// Solidity: function queueWithdrawals((address[],uint256[],address)[] params) returns(bytes32[])
func (_DelegationManager *DelegationManagerTransactor) QueueWithdrawals(opts *bind.TransactOpts, params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "queueWithdrawals", params)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] params) returns(bytes32[])
func (_DelegationManager *DelegationManagerSession) QueueWithdrawals(params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error) {
	return _DelegationManager.Contract.QueueWithdrawals(&_DelegationManager.TransactOpts, params)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] params) returns(bytes32[])
func (_DelegationManager *DelegationManagerTransactorSession) QueueWithdrawals(params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error) {
	return _DelegationManager.Contract.QueueWithdrawals(&_DelegationManager.TransactOpts, params)
}

// Redelegate is a paid mutator transaction binding the contract method 0xa33a3433.
//
// Solidity: function redelegate(address newOperator, (bytes,uint256) newOperatorApproverSig, bytes32 approverSalt) returns(bytes32[] withdrawalRoots)
func (_DelegationManager *DelegationManagerTransactor) Redelegate(opts *bind.TransactOpts, newOperator common.Address, newOperatorApproverSig ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "redelegate", newOperator, newOperatorApproverSig, approverSalt)
}

// Redelegate is a paid mutator transaction binding the contract method 0xa33a3433.
//
// Solidity: function redelegate(address newOperator, (bytes,uint256) newOperatorApproverSig, bytes32 approverSalt) returns(bytes32[] withdrawalRoots)
func (_DelegationManager *DelegationManagerSession) Redelegate(newOperator common.Address, newOperatorApproverSig ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationManager.Contract.Redelegate(&_DelegationManager.TransactOpts, newOperator, newOperatorApproverSig, approverSalt)
}

// Redelegate is a paid mutator transaction binding the contract method 0xa33a3433.
//
// Solidity: function redelegate(address newOperator, (bytes,uint256) newOperatorApproverSig, bytes32 approverSalt) returns(bytes32[] withdrawalRoots)
func (_DelegationManager *DelegationManagerTransactorSession) Redelegate(newOperator common.Address, newOperatorApproverSig ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationManager.Contract.Redelegate(&_DelegationManager.TransactOpts, newOperator, newOperatorApproverSig, approverSalt)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x2aa6d888.
//
// Solidity: function registerAsOperator(address initDelegationApprover, uint32 allocationDelay, string metadataURI) returns()
func (_DelegationManager *DelegationManagerTransactor) RegisterAsOperator(opts *bind.TransactOpts, initDelegationApprover common.Address, allocationDelay uint32, metadataURI string) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "registerAsOperator", initDelegationApprover, allocationDelay, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x2aa6d888.
//
// Solidity: function registerAsOperator(address initDelegationApprover, uint32 allocationDelay, string metadataURI) returns()
func (_DelegationManager *DelegationManagerSession) RegisterAsOperator(initDelegationApprover common.Address, allocationDelay uint32, metadataURI string) (*types.Transaction, error) {
	return _DelegationManager.Contract.RegisterAsOperator(&_DelegationManager.TransactOpts, initDelegationApprover, allocationDelay, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x2aa6d888.
//
// Solidity: function registerAsOperator(address initDelegationApprover, uint32 allocationDelay, string metadataURI) returns()
func (_DelegationManager *DelegationManagerTransactorSession) RegisterAsOperator(initDelegationApprover common.Address, allocationDelay uint32, metadataURI string) (*types.Transaction, error) {
	return _DelegationManager.Contract.RegisterAsOperator(&_DelegationManager.TransactOpts, initDelegationApprover, allocationDelay, metadataURI)
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

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x78296ec5.
//
// Solidity: function updateOperatorMetadataURI(address operator, string metadataURI) returns()
func (_DelegationManager *DelegationManagerTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, operator common.Address, metadataURI string) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "updateOperatorMetadataURI", operator, metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x78296ec5.
//
// Solidity: function updateOperatorMetadataURI(address operator, string metadataURI) returns()
func (_DelegationManager *DelegationManagerSession) UpdateOperatorMetadataURI(operator common.Address, metadataURI string) (*types.Transaction, error) {
	return _DelegationManager.Contract.UpdateOperatorMetadataURI(&_DelegationManager.TransactOpts, operator, metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x78296ec5.
//
// Solidity: function updateOperatorMetadataURI(address operator, string metadataURI) returns()
func (_DelegationManager *DelegationManagerTransactorSession) UpdateOperatorMetadataURI(operator common.Address, metadataURI string) (*types.Transaction, error) {
	return _DelegationManager.Contract.UpdateOperatorMetadataURI(&_DelegationManager.TransactOpts, operator, metadataURI)
}

// DelegationManagerDelegationApproverUpdatedIterator is returned from FilterDelegationApproverUpdated and is used to iterate over the raw logs and unpacked data for DelegationApproverUpdated events raised by the DelegationManager contract.
type DelegationManagerDelegationApproverUpdatedIterator struct {
	Event *DelegationManagerDelegationApproverUpdated // Event containing the contract specifics and raw log

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
func (it *DelegationManagerDelegationApproverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerDelegationApproverUpdated)
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
		it.Event = new(DelegationManagerDelegationApproverUpdated)
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
func (it *DelegationManagerDelegationApproverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerDelegationApproverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerDelegationApproverUpdated represents a DelegationApproverUpdated event raised by the DelegationManager contract.
type DelegationManagerDelegationApproverUpdated struct {
	Operator              common.Address
	NewDelegationApprover common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDelegationApproverUpdated is a free log retrieval operation binding the contract event 0x773b54c04d756fcc5e678111f7d730de3be98192000799eee3d63716055a87c6.
//
// Solidity: event DelegationApproverUpdated(address indexed operator, address newDelegationApprover)
func (_DelegationManager *DelegationManagerFilterer) FilterDelegationApproverUpdated(opts *bind.FilterOpts, operator []common.Address) (*DelegationManagerDelegationApproverUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "DelegationApproverUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerDelegationApproverUpdatedIterator{contract: _DelegationManager.contract, event: "DelegationApproverUpdated", logs: logs, sub: sub}, nil
}

// WatchDelegationApproverUpdated is a free log subscription operation binding the contract event 0x773b54c04d756fcc5e678111f7d730de3be98192000799eee3d63716055a87c6.
//
// Solidity: event DelegationApproverUpdated(address indexed operator, address newDelegationApprover)
func (_DelegationManager *DelegationManagerFilterer) WatchDelegationApproverUpdated(opts *bind.WatchOpts, sink chan<- *DelegationManagerDelegationApproverUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "DelegationApproverUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerDelegationApproverUpdated)
				if err := _DelegationManager.contract.UnpackLog(event, "DelegationApproverUpdated", log); err != nil {
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

// ParseDelegationApproverUpdated is a log parse operation binding the contract event 0x773b54c04d756fcc5e678111f7d730de3be98192000799eee3d63716055a87c6.
//
// Solidity: event DelegationApproverUpdated(address indexed operator, address newDelegationApprover)
func (_DelegationManager *DelegationManagerFilterer) ParseDelegationApproverUpdated(log types.Log) (*DelegationManagerDelegationApproverUpdated, error) {
	event := new(DelegationManagerDelegationApproverUpdated)
	if err := _DelegationManager.contract.UnpackLog(event, "DelegationApproverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerDepositScalingFactorUpdatedIterator is returned from FilterDepositScalingFactorUpdated and is used to iterate over the raw logs and unpacked data for DepositScalingFactorUpdated events raised by the DelegationManager contract.
type DelegationManagerDepositScalingFactorUpdatedIterator struct {
	Event *DelegationManagerDepositScalingFactorUpdated // Event containing the contract specifics and raw log

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
func (it *DelegationManagerDepositScalingFactorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerDepositScalingFactorUpdated)
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
		it.Event = new(DelegationManagerDepositScalingFactorUpdated)
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
func (it *DelegationManagerDepositScalingFactorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerDepositScalingFactorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerDepositScalingFactorUpdated represents a DepositScalingFactorUpdated event raised by the DelegationManager contract.
type DelegationManagerDepositScalingFactorUpdated struct {
	Staker                  common.Address
	Strategy                common.Address
	NewDepositScalingFactor *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterDepositScalingFactorUpdated is a free log retrieval operation binding the contract event 0x8be932bac54561f27260f95463d9b8ab37e06b2842e5ee2404157cc13df6eb8f.
//
// Solidity: event DepositScalingFactorUpdated(address staker, address strategy, uint256 newDepositScalingFactor)
func (_DelegationManager *DelegationManagerFilterer) FilterDepositScalingFactorUpdated(opts *bind.FilterOpts) (*DelegationManagerDepositScalingFactorUpdatedIterator, error) {

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "DepositScalingFactorUpdated")
	if err != nil {
		return nil, err
	}
	return &DelegationManagerDepositScalingFactorUpdatedIterator{contract: _DelegationManager.contract, event: "DepositScalingFactorUpdated", logs: logs, sub: sub}, nil
}

// WatchDepositScalingFactorUpdated is a free log subscription operation binding the contract event 0x8be932bac54561f27260f95463d9b8ab37e06b2842e5ee2404157cc13df6eb8f.
//
// Solidity: event DepositScalingFactorUpdated(address staker, address strategy, uint256 newDepositScalingFactor)
func (_DelegationManager *DelegationManagerFilterer) WatchDepositScalingFactorUpdated(opts *bind.WatchOpts, sink chan<- *DelegationManagerDepositScalingFactorUpdated) (event.Subscription, error) {

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "DepositScalingFactorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerDepositScalingFactorUpdated)
				if err := _DelegationManager.contract.UnpackLog(event, "DepositScalingFactorUpdated", log); err != nil {
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

// ParseDepositScalingFactorUpdated is a log parse operation binding the contract event 0x8be932bac54561f27260f95463d9b8ab37e06b2842e5ee2404157cc13df6eb8f.
//
// Solidity: event DepositScalingFactorUpdated(address staker, address strategy, uint256 newDepositScalingFactor)
func (_DelegationManager *DelegationManagerFilterer) ParseDepositScalingFactorUpdated(log types.Log) (*DelegationManagerDepositScalingFactorUpdated, error) {
	event := new(DelegationManagerDepositScalingFactorUpdated)
	if err := _DelegationManager.contract.UnpackLog(event, "DepositScalingFactorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Operator           common.Address
	DelegationApprover common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed operator, address delegationApprover)
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

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed operator, address delegationApprover)
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed operator, address delegationApprover)
func (_DelegationManager *DelegationManagerFilterer) ParseOperatorRegistered(log types.Log) (*DelegationManagerOperatorRegistered, error) {
	event := new(DelegationManagerOperatorRegistered)
	if err := _DelegationManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerOperatorSharesBurnedIterator is returned from FilterOperatorSharesBurned and is used to iterate over the raw logs and unpacked data for OperatorSharesBurned events raised by the DelegationManager contract.
type DelegationManagerOperatorSharesBurnedIterator struct {
	Event *DelegationManagerOperatorSharesBurned // Event containing the contract specifics and raw log

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
func (it *DelegationManagerOperatorSharesBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerOperatorSharesBurned)
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
		it.Event = new(DelegationManagerOperatorSharesBurned)
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
func (it *DelegationManagerOperatorSharesBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerOperatorSharesBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerOperatorSharesBurned represents a OperatorSharesBurned event raised by the DelegationManager contract.
type DelegationManagerOperatorSharesBurned struct {
	Operator common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesBurned is a free log retrieval operation binding the contract event 0xeff6aab2bc3f7c648896e1522eae71d6c22e3b0e218206b3f40af0e4d204716b.
//
// Solidity: event OperatorSharesBurned(address indexed operator, address strategy, uint256 shares)
func (_DelegationManager *DelegationManagerFilterer) FilterOperatorSharesBurned(opts *bind.FilterOpts, operator []common.Address) (*DelegationManagerOperatorSharesBurnedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "OperatorSharesBurned", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerOperatorSharesBurnedIterator{contract: _DelegationManager.contract, event: "OperatorSharesBurned", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesBurned is a free log subscription operation binding the contract event 0xeff6aab2bc3f7c648896e1522eae71d6c22e3b0e218206b3f40af0e4d204716b.
//
// Solidity: event OperatorSharesBurned(address indexed operator, address strategy, uint256 shares)
func (_DelegationManager *DelegationManagerFilterer) WatchOperatorSharesBurned(opts *bind.WatchOpts, sink chan<- *DelegationManagerOperatorSharesBurned, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "OperatorSharesBurned", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerOperatorSharesBurned)
				if err := _DelegationManager.contract.UnpackLog(event, "OperatorSharesBurned", log); err != nil {
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

// ParseOperatorSharesBurned is a log parse operation binding the contract event 0xeff6aab2bc3f7c648896e1522eae71d6c22e3b0e218206b3f40af0e4d204716b.
//
// Solidity: event OperatorSharesBurned(address indexed operator, address strategy, uint256 shares)
func (_DelegationManager *DelegationManagerFilterer) ParseOperatorSharesBurned(log types.Log) (*DelegationManagerOperatorSharesBurned, error) {
	event := new(DelegationManagerOperatorSharesBurned)
	if err := _DelegationManager.contract.UnpackLog(event, "OperatorSharesBurned", log); err != nil {
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

// DelegationManagerSlashingWithdrawalCompletedIterator is returned from FilterSlashingWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for SlashingWithdrawalCompleted events raised by the DelegationManager contract.
type DelegationManagerSlashingWithdrawalCompletedIterator struct {
	Event *DelegationManagerSlashingWithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *DelegationManagerSlashingWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerSlashingWithdrawalCompleted)
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
		it.Event = new(DelegationManagerSlashingWithdrawalCompleted)
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
func (it *DelegationManagerSlashingWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerSlashingWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerSlashingWithdrawalCompleted represents a SlashingWithdrawalCompleted event raised by the DelegationManager contract.
type DelegationManagerSlashingWithdrawalCompleted struct {
	WithdrawalRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSlashingWithdrawalCompleted is a free log retrieval operation binding the contract event 0x1f40400889274ed07b24845e5054a87a0cab969eb1277aafe61ae352e7c32a00.
//
// Solidity: event SlashingWithdrawalCompleted(bytes32 withdrawalRoot)
func (_DelegationManager *DelegationManagerFilterer) FilterSlashingWithdrawalCompleted(opts *bind.FilterOpts) (*DelegationManagerSlashingWithdrawalCompletedIterator, error) {

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "SlashingWithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return &DelegationManagerSlashingWithdrawalCompletedIterator{contract: _DelegationManager.contract, event: "SlashingWithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchSlashingWithdrawalCompleted is a free log subscription operation binding the contract event 0x1f40400889274ed07b24845e5054a87a0cab969eb1277aafe61ae352e7c32a00.
//
// Solidity: event SlashingWithdrawalCompleted(bytes32 withdrawalRoot)
func (_DelegationManager *DelegationManagerFilterer) WatchSlashingWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *DelegationManagerSlashingWithdrawalCompleted) (event.Subscription, error) {

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "SlashingWithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerSlashingWithdrawalCompleted)
				if err := _DelegationManager.contract.UnpackLog(event, "SlashingWithdrawalCompleted", log); err != nil {
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

// ParseSlashingWithdrawalCompleted is a log parse operation binding the contract event 0x1f40400889274ed07b24845e5054a87a0cab969eb1277aafe61ae352e7c32a00.
//
// Solidity: event SlashingWithdrawalCompleted(bytes32 withdrawalRoot)
func (_DelegationManager *DelegationManagerFilterer) ParseSlashingWithdrawalCompleted(log types.Log) (*DelegationManagerSlashingWithdrawalCompleted, error) {
	event := new(DelegationManagerSlashingWithdrawalCompleted)
	if err := _DelegationManager.contract.UnpackLog(event, "SlashingWithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegationManagerSlashingWithdrawalQueuedIterator is returned from FilterSlashingWithdrawalQueued and is used to iterate over the raw logs and unpacked data for SlashingWithdrawalQueued events raised by the DelegationManager contract.
type DelegationManagerSlashingWithdrawalQueuedIterator struct {
	Event *DelegationManagerSlashingWithdrawalQueued // Event containing the contract specifics and raw log

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
func (it *DelegationManagerSlashingWithdrawalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerSlashingWithdrawalQueued)
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
		it.Event = new(DelegationManagerSlashingWithdrawalQueued)
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
func (it *DelegationManagerSlashingWithdrawalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerSlashingWithdrawalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerSlashingWithdrawalQueued represents a SlashingWithdrawalQueued event raised by the DelegationManager contract.
type DelegationManagerSlashingWithdrawalQueued struct {
	WithdrawalRoot   [32]byte
	Withdrawal       IDelegationManagerTypesWithdrawal
	SharesToWithdraw []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSlashingWithdrawalQueued is a free log retrieval operation binding the contract event 0x26b2aae26516e8719ef50ea2f6831a2efbd4e37dccdf0f6936b27bc08e793e30.
//
// Solidity: event SlashingWithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal, uint256[] sharesToWithdraw)
func (_DelegationManager *DelegationManagerFilterer) FilterSlashingWithdrawalQueued(opts *bind.FilterOpts) (*DelegationManagerSlashingWithdrawalQueuedIterator, error) {

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "SlashingWithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return &DelegationManagerSlashingWithdrawalQueuedIterator{contract: _DelegationManager.contract, event: "SlashingWithdrawalQueued", logs: logs, sub: sub}, nil
}

// WatchSlashingWithdrawalQueued is a free log subscription operation binding the contract event 0x26b2aae26516e8719ef50ea2f6831a2efbd4e37dccdf0f6936b27bc08e793e30.
//
// Solidity: event SlashingWithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal, uint256[] sharesToWithdraw)
func (_DelegationManager *DelegationManagerFilterer) WatchSlashingWithdrawalQueued(opts *bind.WatchOpts, sink chan<- *DelegationManagerSlashingWithdrawalQueued) (event.Subscription, error) {

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "SlashingWithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerSlashingWithdrawalQueued)
				if err := _DelegationManager.contract.UnpackLog(event, "SlashingWithdrawalQueued", log); err != nil {
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

// ParseSlashingWithdrawalQueued is a log parse operation binding the contract event 0x26b2aae26516e8719ef50ea2f6831a2efbd4e37dccdf0f6936b27bc08e793e30.
//
// Solidity: event SlashingWithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal, uint256[] sharesToWithdraw)
func (_DelegationManager *DelegationManagerFilterer) ParseSlashingWithdrawalQueued(log types.Log) (*DelegationManagerSlashingWithdrawalQueued, error) {
	event := new(DelegationManagerSlashingWithdrawalQueued)
	if err := _DelegationManager.contract.UnpackLog(event, "SlashingWithdrawalQueued", log); err != nil {
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
