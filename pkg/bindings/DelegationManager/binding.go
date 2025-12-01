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

// ISignatureUtilsMixinTypesSignatureWithExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsMixinTypesSignatureWithExpiry struct {
	Signature []byte
	Expiry    *big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// DelegationManagerMetaData contains all meta data concerning the DelegationManager contract.
var DelegationManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_eigenPodManager\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_MIN_WITHDRAWAL_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DELEGATION_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateDelegationApprovalDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateWithdrawalRoot\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawal\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawals\",\"inputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertToDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"withdrawableShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cumulativeWithdrawalsQueued\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"totalQueued\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"curDepositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beaconChainSlashingFactorDecrease\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateTo\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtilsMixinTypes.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegatedTo\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApprover\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApproverSaltIsSpent\",\"inputs\":[{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"spent\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositScalingFactor\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorsShares\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQueuedWithdrawal\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQueuedWithdrawalRoots\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQueuedWithdrawals\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"shares\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashableSharesInQueue\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawableShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"withdrawableShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"depositShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"prevDepositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addedShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minWithdrawalDelayBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyOperatorDetails\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newDelegationApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingWithdrawals\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"pending\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueWithdrawals\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.QueuedWithdrawalParams[]\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"depositShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"__deprecated_withdrawer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queuedWithdrawals\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redelegate\",\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newOperatorApproverSig\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtilsMixinTypes.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerAsOperator\",\"inputs\":[{\"name\":\"initDelegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"slashId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"prevMaxMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newMaxMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"totalDepositSharesToSlash\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorMetadataURI\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DelegationApproverUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newDelegationApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositScalingFactorUpdated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"newDepositScalingFactor\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMetadataURIUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesDecreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesIncreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"totalSlashedShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashingWithdrawalCompleted\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashingWithdrawalQueued\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"withdrawal\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"sharesToWithdraw\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerForceUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ActivelyDelegated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerCannotUndelegate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FullySlashed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDepositScalingFactor\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSnapshotOrdering\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotActivelyDelegated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAllocationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEigenPodManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManagerOrEigenPodManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorsCannotUndelegate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaltSpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"WithdrawalDelayNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalNotQueued\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawerNotCaller\",\"inputs\":[]}]",
	Bin: "0x610160604052348015610010575f5ffd5b5060405161605638038061605683398101604081905261002f916101d9565b808084898989878a6001600160a01b03811661005e576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805293841660a05291831660c052821660e05263ffffffff16610100521661012052610095816100b0565b61014052506100a490506100f6565b50505050505050610364565b5f5f829050601f815111156100e3578260405163305a27a960e01b81526004016100da9190610309565b60405180910390fd5b80516100ee8261033e565b179392505050565b5f54610100900460ff161561015d5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100da565b5f5460ff908116146101ac575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101c2575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f5f5f5f60e0888a0312156101ef575f5ffd5b87516101fa816101ae565b602089015190975061020b816101ae565b604089015190965061021c816101ae565b606089015190955061022d816101ae565b608089015190945061023e816101ae565b60a089015190935063ffffffff81168114610257575f5ffd5b60c08901519092506001600160401b03811115610272575f5ffd5b88015f601f82018b13610283575f5ffd5b81516001600160401b0381111561029c5761029c6101c5565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102ca576102ca6101c5565b6040528181528382016020018d10156102e1575f5ffd5b8160208501602083015e5f602083830101528092508094505050505092959891949750929550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b8051602080830151919081101561035e575f198160200360031b1b821691505b50919050565b60805160a05160c05160e051610100516101205161014051615c006104565f395f8181610fc10152613db201525f81816104120152613a4001525f8181610703015281816131b101528181613483015261363c01525f818161075301528181610cf301528181610eb60152818161103901528181611392015281816117f401528181612446015261413001525f818161043901528181610e34015281816112f10152818161156501528181612d3701528181612f1801526132f701525f818161036f01528181610e02015281816114b901526132d101525f81816105e201528181612afd0152613e210152615c005ff3fe608060405234801561000f575f5ffd5b50600436106102b1575f3560e01c80636d70f7ae1161017b578063bb45fef2116100e4578063e4cc3f901161009e578063f698da2511610079578063f698da25146107ce578063fabc1cbc146107d6578063fd8aa88d146107e9578063fe4b84df146107fc575f5ffd5b8063e4cc3f9014610788578063eea9064b1461079b578063f0e0e676146107ae575f5ffd5b8063bb45fef2146106b9578063bfae3fd2146106e6578063c448feb8146106f9578063c978f7ac1461072d578063ca8aa7c71461074e578063da8be86414610775575f5ffd5b80639104c319116101355780639104c319146106175780639435bb431461063257806399f5371b14610645578063a178848414610665578063a33a343314610684578063b7f06ebe14610697575f5ffd5b80636d70f7ae1461057a5780636e1744481461058d578063778e55f3146105a057806378296ec5146105ca578063886f1195146105dd5780639004134714610604575f5ffd5b806354b7c96c1161021d5780635c975abb116101d75780635c975abb146104d45780635d975e88146104dc5780635dd68579146104fd57806360a0d1ce1461051e57806365da12641461053157806366d5ba9314610559575f5ffd5b806354b7c96c1461045b57806354fd4d501461046e578063595c6a6714610483578063597b36da1461048b5780635ac86ab71461049e5780635ae679a7146104c1575f5ffd5b806339b70e381161026e57806339b70e381461036a5780633c651cf2146103a95780633cdeb5e0146103bc5780633e28391d146103ea5780634657e26a1461040d5780634665bcda14610434575f5ffd5b806304a4f979146102b55780630b9f487a146102ef5780630dd8dd0214610302578063136439dd1461032257806325df922e146103375780632aa6d88814610357575b5f5ffd5b6102dc7f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad81565b6040519081526020015b60405180910390f35b6102dc6102fd366004614a6a565b61080f565b610315610310366004614b01565b610897565b6040516102e69190614b3f565b610335610330366004614b76565b610b09565b005b61034a610345366004614d0b565b610b43565b6040516102e69190614db9565b610335610365366004614e1b565b610ca3565b6103917f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016102e6565b6103356103b7366004614e79565b610df7565b6103916103ca366004614ebc565b6001600160a01b039081165f908152609960205260409020600101541690565b6103fd6103f8366004614ebc565b610f4a565b60405190151581526020016102e6565b6103917f000000000000000000000000000000000000000000000000000000000000000081565b6103917f000000000000000000000000000000000000000000000000000000000000000081565b610335610469366004614ed7565b610f69565b610476610fba565b6040516102e69190614f3c565b610335610fea565b6102dc61049936600461500a565b610ffe565b6103fd6104ac36600461503b565b606654600160ff9092169190911b9081161490565b6102dc6104cf36600461506f565b61102d565b6066546102dc565b6104ef6104ea366004614b76565b61119f565b6040516102e69291906151a4565b61051061050b366004614ebc565b6111bc565b6040516102e6929190615216565b61033561052c366004615283565b6112e6565b61039161053f366004614ebc565b609a6020525f90815260409020546001600160a01b031681565b61056c610567366004614ebc565b611491565b6040516102e69291906152c2565b6103fd610588366004614ebc565b611791565b6102dc61059b366004614ed7565b6117c9565b6102dc6105ae366004614ed7565b609860209081525f928352604080842090915290825290205481565b6103356105d83660046152d4565b611873565b6103917f000000000000000000000000000000000000000000000000000000000000000081565b61034a610612366004615324565b6118ec565b61039173beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b610335610640366004615370565b6119c2565b610658610653366004614b76565b611a7d565b6040516102e6919061540c565b6102dc610673366004614ebc565b609f6020525f908152604090205481565b61031561069236600461541e565b611b99565b6103fd6106a5366004614b76565b609e6020525f908152604090205460ff1681565b6103fd6106c7366004615505565b609c60209081525f928352604080842090915290825290205460ff1681565b6102dc6106f4366004614ed7565b611bb1565b60405163ffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526020016102e6565b61074061073b366004615324565b611bed565b6040516102e692919061552f565b6103917f000000000000000000000000000000000000000000000000000000000000000081565b610315610783366004614ebc565b611e7a565b61033561079636600461554e565b611fa3565b6103356107a936600461541e565b611fdb565b6107c16107bc3660046155cc565b612046565b6040516102e69190615679565b6102dc6120eb565b6103356107e4366004614b76565b6121a4565b6103156107f7366004614ebc565b612212565b61033561080a366004614b76565b612235565b604080517f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad60208201526001600160a01b03808616928201929092528187166060820152908516608082015260a0810183905260c081018290525f9061088d9060e00160405160208183030381529060405280519060200120612346565b9695505050505050565b606060016108a481612374565b6108ac6123a2565b5f836001600160401b038111156108c5576108c5614b8d565b6040519080825280602002602001820160405280156108ee578160200160208202803683370190505b50335f908152609a60205260408120549192506001600160a01b03909116905b85811015610afa578686828181106109285761092861568b565b905060200281019061093a919061569f565b6109489060208101906156bd565b905087878381811061095c5761095c61568b565b905060200281019061096e919061569f565b61097890806156bd565b905014610998576040516343714afd60e01b815260040160405180910390fd5b5f610a0233848a8a868181106109b0576109b061568b565b90506020028101906109c2919061569f565b6109cc90806156bd565b808060200260200160405190810160405280939291908181526020018383602002808284375f920191909152506123fb92505050565b9050610ad433848a8a86818110610a1b57610a1b61568b565b9050602002810190610a2d919061569f565b610a3790806156bd565b808060200260200160405190810160405280939291908181526020018383602002808284375f920191909152508e92508d9150889050818110610a7c57610a7c61568b565b9050602002810190610a8e919061569f565b610a9c9060208101906156bd565b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525088925061254d915050565b848381518110610ae657610ae661568b565b60209081029190910101525060010161090e565b5050600160c955949350505050565b610b11612ae8565b6066548181168114610b365760405163c61dca5d60e01b815260040160405180910390fd5b610b3f82612b8b565b5050565b6001600160a01b038084165f908152609a60205260408120546060921690610b6c8683876123fb565b90505f85516001600160401b03811115610b8857610b88614b8d565b604051908082528060200260200182016040528015610bb1578160200160208202803683370190505b5090505f5b8651811015610c96576001600160a01b0388165f90815260a260205260408120885182908a9085908110610bec57610bec61568b565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f206040518060200160405290815f820154815250509050610c70878381518110610c3e57610c3e61568b565b6020026020010151858481518110610c5857610c5861568b565b602002602001015183612bc89092919063ffffffff16565b838381518110610c8257610c8261568b565b602090810291909101015250600101610bb6565b50925050505b9392505050565b610cab6123a2565b610cb433610f4a565b15610cd257604051633bf2b50360e11b815260040160405180910390fd5b604051632b6241f360e11b815233600482015263ffffffff841660248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906356c483e6906044015f604051808303815f87803b158015610d3c575f5ffd5b505af1158015610d4e573d5f5f3e3d5ffd5b50505050610d5c3385612be6565b610d663333612c48565b6040516001600160a01b038516815233907fa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c19060200160405180910390a2336001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b67080908383604051610ddf929190615702565b60405180910390a2610df1600160c955565b50505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610e565750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b610e735760405163045206a560e21b815260040160405180910390fd5b610e7b6123a2565b6001600160a01b038481165f908152609a602052604080822054905163152667d960e31b8152908316600482018190528684166024830152927f0000000000000000000000000000000000000000000000000000000000000000169063a9333ec890604401602060405180830381865afa158015610efb573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f1f9190615730565b90505f610f2d878784612ed1565b9050610f3d838888888886612fb3565b505050610df1600160c955565b6001600160a01b039081165f908152609a602052604090205416151590565b81610f73816130f8565b610f7b6123a2565b610f8483611791565b610fa1576040516325ec6c1f60e01b815260040160405180910390fd5b610fab8383612be6565b610fb5600160c955565b505050565b6060610fe57f000000000000000000000000000000000000000000000000000000000000000061311e565b905090565b610ff2612ae8565b610ffc5f19612b8b565b565b5f81604051602001611010919061540c565b604051602081830303815290604052805190602001209050919050565b5f336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611077576040516323d871a560e01b815260040160405180910390fd5b61107f6123a2565b6001600160a01b038088165f9081526098602090815260408083209388168352929052908120546110bd906001600160401b0380871690861661315b565b90505f6110cc89878787613173565b90506110d8818361575f565b92506110e6895f8885613230565b604080516001600160a01b038881168252602082018690528b16917fdd611f4ef63f4385f1756c86ce1f1f389a9013ba6fa07daba8528291bc2d3c30910160405180910390a2611135866132aa565b6001600160a01b0316633fb99ca5898989876040518563ffffffff1660e01b81526004016111669493929190615772565b5f604051808303815f87803b15801561117d575f5ffd5b505af115801561118f573d5f5f3e3d5ffd5b50505050505061088d600160c955565b6111a761492b565b60606111b28361331c565b9094909350915050565b6060805f6111c984612212565b8051909150806001600160401b038111156111e6576111e6614b8d565b60405190808252806020026020018201604052801561121f57816020015b61120c61492b565b8152602001906001900390816112045790505b509350806001600160401b0381111561123a5761123a614b8d565b60405190808252806020026020018201604052801561126d57816020015b60608152602001906001900390816112585790505b5092505f5b818110156112de5761129c83828151811061128f5761128f61568b565b602002602001015161331c565b8683815181106112ae576112ae61568b565b602002602001018684815181106112c7576112c761568b565b602090810291909101019190915252600101611272565b505050915091565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461132f57604051633213a66160e21b815260040160405180910390fd5b6113376123a2565b61134083610f4a565b15610fab576001600160a01b038381165f908152609a602052604080822054905163152667d960e31b81529083166004820181905273beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac06024830152927f0000000000000000000000000000000000000000000000000000000000000000169063a9333ec890604401602060405180830381865afa1580156113d7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113fb9190615730565b6001600160a01b0386165f90815260a26020908152604080832073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0845282528083208151928301909152548152919250611461866114596001600160401b0380871690891661356f565b849190613583565b9050611483848873beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac084613230565b50505050610fb5600160c955565b6040516394f649dd60e01b81526001600160a01b03828116600483015260609182915f9182917f000000000000000000000000000000000000000000000000000000000000000016906394f649dd906024015f60405180830381865afa1580156114fd573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526115249190810190615822565b60405163fe243a1760e01b81526001600160a01b03888116600483015273beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac060248301529294509092505f917f0000000000000000000000000000000000000000000000000000000000000000169063fe243a1790604401602060405180830381865afa1580156115aa573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115ce91906158dd565b9050805f036115e257509094909350915050565b5f835160016115f1919061575f565b6001600160401b0381111561160857611608614b8d565b604051908082528060200260200182016040528015611631578160200160208202803683370190505b5090505f84516001611643919061575f565b6001600160401b0381111561165a5761165a614b8d565b604051908082528060200260200182016040528015611683578160200160208202803683370190505b50905073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0828651815181106116ae576116ae61568b565b60200260200101906001600160a01b031690816001600160a01b03168152505082818651815181106116e2576116e261568b565b60209081029190910101525f5b85518110156117835785818151811061170a5761170a61568b565b60200260200101518382815181106117245761172461568b565b60200260200101906001600160a01b031690816001600160a01b0316815250508481815181106117565761175661568b565b60200260200101518282815181106117705761177061568b565b60209081029190910101526001016116ef565b509097909650945050505050565b5f6001600160a01b038216158015906117c357506001600160a01b038083165f818152609a6020526040902054909116145b92915050565b60405163152667d960e31b81526001600160a01b03838116600483015282811660248301525f9182917f0000000000000000000000000000000000000000000000000000000000000000169063a9333ec890604401602060405180830381865afa158015611839573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061185d9190615730565b905061186b8484835f613173565b949350505050565b8261187d816130f8565b61188684611791565b6118a3576040516325ec6c1f60e01b815260040160405180910390fd5b836001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b670809084846040516118de929190615702565b60405180910390a250505050565b60605f82516001600160401b0381111561190857611908614b8d565b604051908082528060200260200182016040528015611931578160200160208202803683370190505b5090505f5b83518110156119ba576001600160a01b0385165f908152609860205260408120855190919086908490811061196d5761196d61568b565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f20548282815181106119a7576119a761568b565b6020908102919091010152600101611936565b509392505050565b60026119cd81612374565b6119d56123a2565b855f5b81811015611a6857611a608989838181106119f5576119f561568b565b9050602002810190611a0791906158f4565b611a1090615908565b888884818110611a2257611a2261568b565b9050602002810190611a3491906156bd565b888886818110611a4657611a4661568b565b9050602002016020810190611a5b9190615913565b6135a1565b6001016119d8565b5050611a74600160c955565b50505050505050565b611a8561492b565b5f82815260a46020908152604091829020825160e08101845281546001600160a01b03908116825260018301548116828501526002830154168185015260038201546060820152600482015463ffffffff1660808201526005820180548551818602810186019096528086529194929360a08601939290830182828015611b3357602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311611b15575b5050505050815260200160068201805480602002602001604051908101604052809291908181526020018280548015611b8957602002820191905f5260205f20905b815481526020019060010190808311611b75575b5050505050815250509050919050565b6060611ba433611e7a565b9050610c9c848484611fdb565b6001600160a01b038083165f90815260a260209081526040808320938516835292815282822083519182019093529154825290610c9c906139e3565b60608082516001600160401b03811115611c0957611c09614b8d565b604051908082528060200260200182016040528015611c32578160200160208202803683370190505b50915082516001600160401b03811115611c4e57611c4e614b8d565b604051908082528060200260200182016040528015611c77578160200160208202803683370190505b506001600160a01b038086165f908152609a6020526040812054929350911690611ca28683876123fb565b90505f5b8551811015611e6f575f611cd2878381518110611cc557611cc561568b565b60200260200101516132aa565b9050806001600160a01b031663fe243a1789898581518110611cf657611cf661568b565b60200260200101516040518363ffffffff1660e01b8152600401611d309291906001600160a01b0392831681529116602082015260400190565b602060405180830381865afa158015611d4b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d6f91906158dd565b858381518110611d8157611d8161568b565b6020026020010181815250505f60a25f8a6001600160a01b03166001600160a01b031681526020019081526020015f205f898581518110611dc457611dc461568b565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f206040518060200160405290815f820154815250509050611e48868481518110611e1657611e1661568b565b6020026020010151858581518110611e3057611e3061568b565b6020026020010151836135839092919063ffffffff16565b878481518110611e5a57611e5a61568b565b60209081029190910101525050600101611ca6565b5050505b9250929050565b6060611e846123a2565b611e8d82610f4a565b611eaa5760405163a5c7c44560e01b815260040160405180910390fd5b611eb382611791565b15611ed1576040516311ca333560e31b815260040160405180910390fd5b336001600160a01b03831614611f89576001600160a01b038083165f908152609a602052604090205416611f0481613a02565b80611f2a57506001600160a01b038181165f908152609960205260409020600101541633145b611f4757604051631e499a2360e11b815260040160405180910390fd5b806001600160a01b0316836001600160a01b03167ff0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a60405160405180910390a3505b611f9282613aab565b9050611f9e600160c955565b919050565b6002611fae81612374565b611fb66123a2565b611fca611fc286615908565b8585856135a1565b611fd4600160c955565b5050505050565b611fe36123a2565b611fec33610f4a565b1561200a57604051633bf2b50360e11b815260040160405180910390fd5b61201383611791565b612030576040516325ec6c1f60e01b815260040160405180910390fd5b61203c33848484613ceb565b610fab3384612c48565b60605f83516001600160401b0381111561206257612062614b8d565b60405190808252806020026020018201604052801561209557816020015b60608152602001906001900390816120805790505b5090505f5b84518110156119ba576120c68582815181106120b8576120b861568b565b6020026020010151856118ec565b8282815181106120d8576120d861568b565b602090810291909101015260010161209a565b60408051808201909152600a81526922b4b3b2b72630bcb2b960b11b6020909101525f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea612158613daa565b805160209182012060408051928301949094529281019190915260608101919091524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b6121ac613e1f565b606654801982198116146121d35760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020015b60405180910390a25050565b6001600160a01b0381165f90815260a3602052604090206060906117c390613ed0565b5f54610100900460ff161580801561225357505f54600160ff909116105b8061226c5750303b15801561226c57505f5460ff166001145b6122d45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff1916600117905580156122f5575f805461ff0019166101001790555b6122fe82612b8b565b8015610b3f575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b5f61234f6120eb565b60405161190160f01b6020820152602281019190915260428101839052606201611010565b606654600160ff83161b9081160361239f5760405163840a48d560e01b815260040160405180910390fd5b50565b600260c954036123f45760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016122cb565b600260c955565b60605f82516001600160401b0381111561241757612417614b8d565b604051908082528060200260200182016040528015612440578160200160208202803683370190505b5090505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663547afb8786866040518363ffffffff1660e01b815260040161249292919061592e565b5f60405180830381865afa1580156124ac573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526124d39190810190615951565b90505f5b84518110156125425761251d878683815181106124f6576124f661568b565b60200260200101518484815181106125105761251061568b565b6020026020010151612ed1565b83828151811061252f5761252f61568b565b60209081029190910101526001016124d7565b509095945050505050565b5f6001600160a01b038616612575576040516339b190bb60e11b815260040160405180910390fd5b83515f036125965760405163796cc52560e01b815260040160405180910390fd5b5f84516001600160401b038111156125b0576125b0614b8d565b6040519080825280602002602001820160405280156125d9578160200160208202803683370190505b5090505f85516001600160401b038111156125f6576125f6614b8d565b60405190808252806020026020018201604052801561261f578160200160208202803683370190505b5090505f5b865181101561291b575f612643888381518110611cc557611cc561568b565b90505f60a25f8c6001600160a01b03166001600160a01b031681526020019081526020015f205f8a858151811061267c5761267c61568b565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f2090506126f58884815181106126ba576126ba61568b565b60200260200101518885815181106126d4576126d461568b565b60209081029190910181015160408051928301905284548252909190613583565b8484815181106127075761270761568b565b6020026020010181815250506127468884815181106127285761272861568b565b60209081029190910181015160408051928301905283548252613edc565b8584815181106127585761275861568b565b60209081029190910101526001600160a01b038a16156127ed576127af8a8a85815181106127885761278861568b565b60200260200101518786815181106127a2576127a261568b565b6020026020010151613ef0565b6127ed8a8c8b86815181106127c6576127c661568b565b60200260200101518787815181106127e0576127e061568b565b6020026020010151613230565b5f826001600160a01b031663724af4238d8c87815181106128105761281061568b565b60200260200101518c888151811061282a5761282a61568b565b60200260200101516040518463ffffffff1660e01b8152600401612850939291906159e0565b6020604051808303815f875af115801561286c573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061289091906158dd565b9050805f0361290d575f82557f8be932bac54561f27260f95463d9b8ab37e06b2842e5ee2404157cc13df6eb8f8c8b86815181106128d0576128d061568b565b60200260200101516128f5856040518060200160405290815f820154815250506139e3565b604051612904939291906159e0565b60405180910390a15b505050806001019050612624565b506001600160a01b0388165f908152609f6020526040812080549182919061294283615a04565b91905055505f6040518060e001604052808b6001600160a01b031681526020018a6001600160a01b031681526020018b6001600160a01b031681526020018381526020014363ffffffff1681526020018981526020018581525090505f6129a882610ffe565b5f818152609e602090815260408083208054600160ff19909116811790915560a4835292819020865181546001600160a01b03199081166001600160a01b039283161783558885015195830180548216968316969096179095559187015160028201805490951692169190911790925560608501516003830155608085015160048301805463ffffffff191663ffffffff90921691909117905560a085015180519394508593612a5e9260058501920190614984565b5060c08201518051612a7a9160068401916020909101906149e7565b5050506001600160a01b038b165f90815260a360205260409020612a9e9082613f5a565b507f26b2aae26516e8719ef50ea2f6831a2efbd4e37dccdf0f6936b27bc08e793e30818386604051612ad293929190615a1c565b60405180910390a19a9950505050505050505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015612b4a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612b6e9190615a46565b610ffc57604051631d77d47760e21b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b5f61186b82612be0612bd9876139e3565b8690613f65565b90613f65565b6001600160a01b038281165f8181526099602090815260409182902060010180546001600160a01b0319169486169485179055905192835290917f773b54c04d756fcc5e678111f7d730de3be98192000799eee3d63716055a87c69101612206565b5f612c5281612374565b5f5f612c5d85611491565b915091505f612c6d5f86856123fb565b6001600160a01b038781165f818152609a602052604080822080546001600160a01b031916948b16948517905551939450919290917fc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d874330491a35f5b8351811015611a745773beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac06001600160a01b0316848281518110612d0057612d0061568b565b60200260200101516001600160a01b031603612e705760405163a3d75e0960e01b81526001600160a01b0388811660048301525f917f00000000000000000000000000000000000000000000000000000000000000009091169063a3d75e0990602401602060405180830381865afa158015612d7e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612da29190615730565b90505f60a25f8a6001600160a01b03166001600160a01b031681526020019081526020015f205f878581518110612ddb57612ddb61568b565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020015f206040518060200160405290815f820154815250509050612e4f858481518110612e2d57612e2d61568b565b6020026020010151836001600160401b0316836135839092919063ffffffff16565b858481518110612e6157612e6161568b565b60200260200101818152505050505b612ec98688868481518110612e8757612e8761568b565b60200260200101515f878681518110612ea257612ea261568b565b6020026020010151878781518110612ebc57612ebc61568b565b6020026020010151612fb3565b600101612cc7565b5f73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeabf196001600160a01b03841601612fa35760405163a3d75e0960e01b81526001600160a01b0385811660048301525f917f00000000000000000000000000000000000000000000000000000000000000009091169063a3d75e0990602401602060405180830381865afa158015612f5f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f839190615730565b9050612f9b6001600160401b0384811690831661356f565b915050610c9c565b506001600160401b031692915050565b805f03612fd357604051630a33bc6960e21b815260040160405180910390fd5b81156130f0576001600160a01b038086165f90815260a26020908152604080832093881683529290522061300981858585613f79565b6040805160208101909152815481527f8be932bac54561f27260f95463d9b8ab37e06b2842e5ee2404157cc13df6eb8f9087908790613047906139e3565b604051613056939291906159e0565b60405180910390a161306786610f4a565b15611a74576001600160a01b038088165f908152609860209081526040808320938916835292905290812080548592906130a290849061575f565b92505081905550866001600160a01b03167f1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c8787866040516130e6939291906159e0565b60405180910390a2505b505050505050565b61310181613a02565b61239f5760405163932d94f760e01b815260040160405180910390fd5b60605f61312a8361400f565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f6131698483856001614036565b61186b9085615a61565b6001600160a01b038085165f90815260a560209081526040808320938716835292905290812081906131a490614091565b90505f61320a60016131d67f000000000000000000000000000000000000000000000000000000000000000043615a74565b6131e09190615a74565b6001600160a01b03808a165f90815260a560209081526040808320938c16835292905220906140ab565b90505f6132178284615a61565b90506132248187876140c7565b98975050505050505050565b6001600160a01b038085165f90815260986020908152604080832093861683529290529081208054839290613266908490615a61565b92505081905550836001600160a01b03167f6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd8484846040516118de939291906159e0565b5f6001600160a01b03821673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0146132f5577f00000000000000000000000000000000000000000000000000000000000000006117c3565b7f000000000000000000000000000000000000000000000000000000000000000092915050565b61332461492b565b5f82815260a46020908152604091829020825160e08101845281546001600160a01b0390811682526001830154811682850152600283015416818501526003820154606082810191909152600483015463ffffffff1660808301526005830180548651818702810187019097528087529195929460a086019392908301828280156133d657602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116133b8575b505050505081526020016006820180548060200260200160405190810160405280929190818152602001828054801561342c57602002820191905f5260205f20905b815481526020019060010190808311613418575b50505050508152505091508160a00151516001600160401b0381111561345457613454614b8d565b60405190808252806020026020018201604052801561347d578160200160208202803683370190505b5090505f7f000000000000000000000000000000000000000000000000000000000000000083608001516134b19190615a90565b90505f4363ffffffff168263ffffffff16106134e2576134dd845f015185602001518660a001516123fb565b6134f9565b6134f9845f015185602001518660a00151856140e5565b90505f5b8460a00151518110156112de5761354a8560c0015182815181106135235761352361568b565b602002602001015183838151811061353d5761353d61568b565b6020026020010151614213565b84828151811061355c5761355c61568b565b60209081029190910101526001016134fd565b5f610c9c8383670de0b6b3a764000061421e565b5f61186b8261359b613594876139e3565b869061356f565b9061356f565b60a08401515182146135c6576040516343714afd60e01b815260040160405180910390fd5b83604001516001600160a01b0316336001600160a01b0316146135fc576040516316110d3560e21b815260040160405180910390fd5b5f61360685610ffe565b5f818152609e602052604090205490915060ff16613637576040516387c9d21960e01b815260040160405180910390fd5b60605f7f0000000000000000000000000000000000000000000000000000000000000000876080015161366a9190615a90565b90508063ffffffff164363ffffffff1611613698576040516378f67ae160e11b815260040160405180910390fd5b6136af875f015188602001518960a00151846140e5565b87516001600160a01b03165f90815260a3602052604090209092506136d5915083614303565b505f82815260a46020526040812080546001600160a01b031990811682556001820180548216905560028201805490911690556003810182905560048101805463ffffffff191690559061372c6005830182614a20565b613739600683015f614a20565b50505f828152609e602052604090819020805460ff19169055517f1f40400889274ed07b24845e5054a87a0cab969eb1277aafe61ae352e7c32a00906137829084815260200190565b60405180910390a185516001600160a01b039081165f908152609a6020526040812054885160a08a015191909316926137bc9184906123fb565b90505f5b8860a00151518110156139d8575f6137e78a60a001518381518110611cc557611cc561568b565b90505f61381d8b60c0015184815181106138035761380361568b565b602002602001015187858151811061353d5761353d61568b565b9050805f0361382d5750506139d0565b87156138fb57816001600160a01b0316632eae418c8c5f01518d60a00151868151811061385c5761385c61568b565b60200260200101518d8d888181106138765761387661568b565b905060200201602081019061388b9190614ebc565b60405160e085901b6001600160e01b03191681526001600160a01b03938416600482015291831660248301529091166044820152606481018490526084015f604051808303815f87803b1580156138e0575f5ffd5b505af11580156138f2573d5f5f3e3d5ffd5b505050506139cd565b5f5f836001600160a01b03166350ff72258e5f01518f60a0015188815181106139265761392661568b565b6020026020010151866040518463ffffffff1660e01b815260040161394d939291906159e0565b60408051808303815f875af1158015613968573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061398c9190615aac565b915091506139ca878e5f01518f60a0015188815181106139ae576139ae61568b565b602002602001015185858b8b81518110612ebc57612ebc61568b565b50505b50505b6001016137c0565b505050505050505050565b80515f90156139f35781516117c3565b670de0b6b3a764000092915050565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb890608401602060405180830381865afa158015613a87573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117c39190615a46565b60606001613ab881612374565b6001600160a01b038084165f818152609a602052604080822080546001600160a01b0319811690915590519316928392917ffee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af4467691a35f5f613b1786611491565b9150915081515f03613b2b57505050613ce5565b81516001600160401b03811115613b4457613b44614b8d565b604051908082528060200260200182016040528015613b6d578160200160208202803683370190505b5094505f613b7c8785856123fb565b90505f5b8351811015613cdf576040805160018082528183019092525f916020808301908036833750506040805160018082528183019092529293505f9291506020808301908036833750506040805160018082528183019092529293505f92915060208083019080368337019050509050868481518110613c0057613c0061568b565b6020026020010151835f81518110613c1a57613c1a61568b565b60200260200101906001600160a01b031690816001600160a01b031681525050858481518110613c4c57613c4c61568b565b6020026020010151825f81518110613c6657613c6661568b565b602002602001018181525050848481518110613c8457613c8461568b565b6020026020010151815f81518110613c9e57613c9e61568b565b602002602001018181525050613cb78b8985858561254d565b8a8581518110613cc957613cc961568b565b6020908102919091010152505050600101613b80565b50505050505b50919050565b6001600160a01b038084165f908152609960205260409020600101541680613d135750610df1565b6001600160a01b0381165f908152609c6020908152604080832085845290915290205460ff1615613d5757604051630d4c4c9160e21b815260040160405180910390fd5b6001600160a01b0381165f908152609c602090815260408083208584528252909120805460ff19166001179055830151611fd4908290613d9e90889088908490889061080f565b8551602087015161430e565b60605f613dd67f000000000000000000000000000000000000000000000000000000000000000061311e565b9050805f81518110613dea57613dea61568b565b016020908101516040516001600160f81b03199091169181019190915260210160405160208183030381529060405291505090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613e7b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613e9f9190615ace565b6001600160a01b0316336001600160a01b031614610ffc5760405163794821ff60e01b815260040160405180910390fd5b60605f610c9c83614360565b5f610c9c613ee9846139e3565b839061356f565b6001600160a01b038084165f90815260a5602090815260408083209386168352929052908120613f1f90614091565b9050610df143613f2f848461575f565b6001600160a01b038088165f90815260a560209081526040808320938a1683529290522091906143b9565b5f610c9c83836143c4565b5f610c9c83670de0b6b3a76400008461421e565b825f03613fa557604080516020810190915284548152613f9e908290612be0906139e3565b8455610df1565b6040805160208101909152845481525f90613fc1908584613583565b90505f613fce848361575f565b90505f613fe984612be0613fe2888a61575f565b8590613f65565b80885590505f819003611a745760405163172cec7360e31b815260040160405180910390fd5b5f60ff8216601f8111156117c357604051632cd44ac360e21b815260040160405180910390fd5b5f5f61404386868661421e565b9050600183600281111561405957614059615ae9565b14801561407557505f848061407057614070615afd565b868809115b156140885761408560018261575f565b90505b95945050505050565b5f61409c8282614410565b6001600160e01b031692915050565b5f6140b7838383614455565b6001600160e01b03169392505050565b5f61186b6140d58385615b11565b85906001600160401b031661356f565b60605f83516001600160401b0381111561410157614101614b8d565b60405190808252806020026020018201604052801561412a578160200160208202803683370190505b5090505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166394d7d00c8787876040518463ffffffff1660e01b815260040161417e93929190615b30565b5f60405180830381865afa158015614198573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526141bf9190810190615951565b90505f5b8551811015614207576141e2888783815181106124f6576124f661568b565b8382815181106141f4576141f461568b565b60209081029190910101526001016141c3565b50909695505050505050565b5f610c9c838361356f565b5f80805f19858709858702925082811083820303915050805f036142555783828161424b5761424b615afd565b0492505050610c9c565b80841161429c5760405162461bcd60e51b81526020600482015260156024820152744d6174683a206d756c446976206f766572666c6f7760581b60448201526064016122cb565b5f8486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091025f889003889004909101858311909403939093029303949094049190911702949350505050565b5f610c9c838361449e565b4281101561432f57604051630819bdcd60e01b815260040160405180910390fd5b6143436001600160a01b0385168484614581565b610df157604051638baa579f60e01b815260040160405180910390fd5b6060815f018054806020026020016040519081016040528092919081815260200182805480156143ad57602002820191905f5260205f20905b815481526020019060010190808311614399575b50505050509050919050565b610fb58383836145d5565b5f81815260018301602052604081205461440957508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556117c3565b505f6117c3565b81545f90801561444d5761443684614429600184615a61565b5f91825260209091200190565b5464010000000090046001600160e01b031661186b565b509092915050565b82545f9081614466868683856146db565b905080156144945761447d86614429600184615a61565b5464010000000090046001600160e01b031661088d565b5091949350505050565b5f8181526001830160205260408120548015614578575f6144c0600183615a61565b85549091505f906144d390600190615a61565b9050818114614532575f865f0182815481106144f1576144f161568b565b905f5260205f200154905080875f0184815481106145115761451161568b565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061454357614543615b69565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506117c3565b5f9150506117c3565b5f5f5f61458e858561472e565b90925090505f8160048111156145a6576145a6615ae9565b1480156145c45750856001600160a01b0316826001600160a01b0316145b8061088d575061088d86868661476d565b8254801561468d575f6145ed85614429600185615a61565b60408051808201909152905463ffffffff8082168084526401000000009092046001600160e01b0316602084015291925090851610156146405760405163151b8e3f60e11b815260040160405180910390fd5b805163ffffffff80861691160361468b578261466186614429600186615a61565b80546001600160e01b03929092166401000000000263ffffffff9092169190911790555050505050565b505b506040805180820190915263ffffffff92831681526001600160e01b03918216602080830191825285546001810187555f968752952091519051909216640100000000029190921617910155565b5f5b818310156119ba575f6146f08484614854565b5f8781526020902090915063ffffffff86169082015463ffffffff16111561471a57809250614728565b61472581600161575f565b93505b506146dd565b5f5f8251604103614762576020830151604084015160608501515f1a6147568782858561486e565b94509450505050611e73565b505f90506002611e73565b5f5f5f856001600160a01b0316631626ba7e60e01b8686604051602401614795929190615b7d565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516147d39190615b95565b5f60405180830381855afa9150503d805f811461480b576040519150601f19603f3d011682016040523d82523d5f602084013e614810565b606091505b509150915081801561482457506020815110155b801561088d57508051630b135d3f60e11b9061484990830160209081019084016158dd565b149695505050505050565b5f6148626002848418615bab565b610c9c9084841661575f565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156148a357505f90506003614922565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156148f4573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661491c575f60019250925050614922565b91505f90505b94509492505050565b6040518060e001604052805f6001600160a01b031681526020015f6001600160a01b031681526020015f6001600160a01b031681526020015f81526020015f63ffffffff16815260200160608152602001606081525090565b828054828255905f5260205f209081019282156149d7579160200282015b828111156149d757825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906149a2565b506149e3929150614a37565b5090565b828054828255905f5260205f209081019282156149d7579160200282015b828111156149d7578251825591602001919060010190614a05565b5080545f8255905f5260205f209081019061239f91905b5b808211156149e3575f8155600101614a38565b6001600160a01b038116811461239f575f5ffd5b8035611f9e81614a4b565b5f5f5f5f5f60a08688031215614a7e575f5ffd5b8535614a8981614a4b565b94506020860135614a9981614a4b565b93506040860135614aa981614a4b565b94979396509394606081013594506080013592915050565b5f5f83601f840112614ad1575f5ffd5b5081356001600160401b03811115614ae7575f5ffd5b6020830191508360208260051b8501011115611e73575f5ffd5b5f5f60208385031215614b12575f5ffd5b82356001600160401b03811115614b27575f5ffd5b614b3385828601614ac1565b90969095509350505050565b602080825282518282018190525f918401906040840190835b81811015612542578351835260209384019390920191600101614b58565b5f60208284031215614b86575f5ffd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b60405160e081016001600160401b0381118282101715614bc357614bc3614b8d565b60405290565b604080519081016001600160401b0381118282101715614bc357614bc3614b8d565b604051601f8201601f191681016001600160401b0381118282101715614c1357614c13614b8d565b604052919050565b5f6001600160401b03821115614c3357614c33614b8d565b5060051b60200190565b5f82601f830112614c4c575f5ffd5b8135614c5f614c5a82614c1b565b614beb565b8082825260208201915060208360051b860101925085831115614c80575f5ffd5b602085015b83811015614ca6578035614c9881614a4b565b835260209283019201614c85565b5095945050505050565b5f82601f830112614cbf575f5ffd5b8135614ccd614c5a82614c1b565b8082825260208201915060208360051b860101925085831115614cee575f5ffd5b602085015b83811015614ca6578035835260209283019201614cf3565b5f5f5f60608486031215614d1d575f5ffd5b8335614d2881614a4b565b925060208401356001600160401b03811115614d42575f5ffd5b614d4e86828701614c3d565b92505060408401356001600160401b03811115614d69575f5ffd5b614d7586828701614cb0565b9150509250925092565b5f8151808452602084019350602083015f5b82811015614daf578151865260209586019590910190600101614d91565b5093949350505050565b602081525f610c9c6020830184614d7f565b803563ffffffff81168114611f9e575f5ffd5b5f5f83601f840112614dee575f5ffd5b5081356001600160401b03811115614e04575f5ffd5b602083019150836020828501011115611e73575f5ffd5b5f5f5f5f60608587031215614e2e575f5ffd5b8435614e3981614a4b565b9350614e4760208601614dcb565b925060408501356001600160401b03811115614e61575f5ffd5b614e6d87828801614dde565b95989497509550505050565b5f5f5f5f60808587031215614e8c575f5ffd5b8435614e9781614a4b565b93506020850135614ea781614a4b565b93969395505050506040820135916060013590565b5f60208284031215614ecc575f5ffd5b8135610c9c81614a4b565b5f5f60408385031215614ee8575f5ffd5b8235614ef381614a4b565b91506020830135614f0381614a4b565b809150509250929050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f610c9c6020830184614f0e565b5f60e08284031215614f5e575f5ffd5b614f66614ba1565b9050614f7182614a5f565b8152614f7f60208301614a5f565b6020820152614f9060408301614a5f565b604082015260608281013590820152614fab60808301614dcb565b608082015260a08201356001600160401b03811115614fc8575f5ffd5b614fd484828501614c3d565b60a08301525060c08201356001600160401b03811115614ff2575f5ffd5b614ffe84828501614cb0565b60c08301525092915050565b5f6020828403121561501a575f5ffd5b81356001600160401b0381111561502f575f5ffd5b61186b84828501614f4e565b5f6020828403121561504b575f5ffd5b813560ff81168114610c9c575f5ffd5b6001600160401b038116811461239f575f5ffd5b5f5f5f5f5f5f86880360e0811215615085575f5ffd5b873561509081614a4b565b96506040601f19820112156150a3575f5ffd5b506020870194506060870135935060808701356150bf81614a4b565b925060a08701356150cf8161505b565b915060c08701356150df8161505b565b809150509295509295509295565b5f8151808452602084019350602083015f5b82811015614daf5781516001600160a01b03168652602095860195909101906001016150ff565b80516001600160a01b03908116835260208083015182169084015260408083015190911690830152606080820151908301526080808201515f916151719085018263ffffffff169052565b5060a082015160e060a085015261518b60e08501826150ed565b905060c083015184820360c08601526140888282614d7f565b604081525f6151b66040830185615126565b82810360208401526140888185614d7f565b5f82825180855260208501945060208160051b830101602085015f5b8381101561420757601f19858403018852615200838351614d7f565b60209889019890935091909101906001016151e4565b5f604082016040835280855180835260608501915060608160051b8601019250602087015f5b8281101561526d57605f19878603018452615258858351615126565b9450602093840193919091019060010161523c565b50505050828103602084015261408881856151c8565b5f5f5f60608486031215615295575f5ffd5b83356152a081614a4b565b92506020840135915060408401356152b78161505b565b809150509250925092565b604081525f6151b660408301856150ed565b5f5f5f604084860312156152e6575f5ffd5b83356152f181614a4b565b925060208401356001600160401b0381111561530b575f5ffd5b61531786828701614dde565b9497909650939450505050565b5f5f60408385031215615335575f5ffd5b823561534081614a4b565b915060208301356001600160401b0381111561535a575f5ffd5b61536685828601614c3d565b9150509250929050565b5f5f5f5f5f5f60608789031215615385575f5ffd5b86356001600160401b0381111561539a575f5ffd5b6153a689828a01614ac1565b90975095505060208701356001600160401b038111156153c4575f5ffd5b6153d089828a01614ac1565b90955093505060408701356001600160401b038111156153ee575f5ffd5b6153fa89828a01614ac1565b979a9699509497509295939492505050565b602081525f610c9c6020830184615126565b5f5f5f60608486031215615430575f5ffd5b833561543b81614a4b565b925060208401356001600160401b03811115615455575f5ffd5b840160408187031215615466575f5ffd5b61546e614bc9565b81356001600160401b03811115615483575f5ffd5b8201601f81018813615493575f5ffd5b80356001600160401b038111156154ac576154ac614b8d565b6154bf601f8201601f1916602001614beb565b8181528960208385010111156154d3575f5ffd5b816020840160208301375f60209282018301528352928301359282019290925293969395505050506040919091013590565b5f5f60408385031215615516575f5ffd5b823561552181614a4b565b946020939093013593505050565b604081525f6151b66040830185614d7f565b801515811461239f575f5ffd5b5f5f5f5f60608587031215615561575f5ffd5b84356001600160401b03811115615576575f5ffd5b850160e08188031215615587575f5ffd5b935060208501356001600160401b038111156155a1575f5ffd5b6155ad87828801614ac1565b90945092505060408501356155c181615541565b939692955090935050565b5f5f604083850312156155dd575f5ffd5b82356001600160401b038111156155f2575f5ffd5b8301601f81018513615602575f5ffd5b8035615610614c5a82614c1b565b8082825260208201915060208360051b850101925087831115615631575f5ffd5b6020840193505b8284101561565c57833561564b81614a4b565b825260209384019390910190615638565b945050505060208301356001600160401b0381111561535a575f5ffd5b602081525f610c9c60208301846151c8565b634e487b7160e01b5f52603260045260245ffd5b5f8235605e198336030181126156b3575f5ffd5b9190910192915050565b5f5f8335601e198436030181126156d2575f5ffd5b8301803591506001600160401b038211156156eb575f5ffd5b6020019150600581901b3603821315611e73575f5ffd5b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b5f60208284031215615740575f5ffd5b8151610c9c8161505b565b634e487b7160e01b5f52601160045260245ffd5b808201808211156117c3576117c361574b565b60a08101853561578181614a4b565b6001600160a01b0316825263ffffffff61579d60208801614dcb565b16602083015260408201949094526001600160a01b03929092166060830152608090910152919050565b5f82601f8301126157d6575f5ffd5b81516157e4614c5a82614c1b565b8082825260208201915060208360051b860101925085831115615805575f5ffd5b602085015b83811015614ca657805183526020928301920161580a565b5f5f60408385031215615833575f5ffd5b82516001600160401b03811115615848575f5ffd5b8301601f81018513615858575f5ffd5b8051615866614c5a82614c1b565b8082825260208201915060208360051b850101925087831115615887575f5ffd5b6020840193505b828410156158b25783516158a181614a4b565b82526020938401939091019061588e565b8095505050505060208301516001600160401b038111156158d1575f5ffd5b615366858286016157c7565b5f602082840312156158ed575f5ffd5b5051919050565b5f823560de198336030181126156b3575f5ffd5b5f6117c33683614f4e565b5f60208284031215615923575f5ffd5b8135610c9c81615541565b6001600160a01b03831681526040602082018190525f9061186b908301846150ed565b5f60208284031215615961575f5ffd5b81516001600160401b03811115615976575f5ffd5b8201601f81018413615986575f5ffd5b8051615994614c5a82614c1b565b8082825260208201915060208360051b8501019250868311156159b5575f5ffd5b6020840193505b8284101561088d5783516159cf8161505b565b8252602093840193909101906159bc565b6001600160a01b039384168152919092166020820152604081019190915260600190565b5f60018201615a1557615a1561574b565b5060010190565b838152606060208201525f615a346060830185615126565b828103604084015261088d8185614d7f565b5f60208284031215615a56575f5ffd5b8151610c9c81615541565b818103818111156117c3576117c361574b565b63ffffffff82811682821603908111156117c3576117c361574b565b63ffffffff81811683821601908111156117c3576117c361574b565b5f5f60408385031215615abd575f5ffd5b505080516020909101519092909150565b5f60208284031215615ade575f5ffd5b8151610c9c81614a4b565b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52601260045260245ffd5b6001600160401b0382811682821603908111156117c3576117c361574b565b6001600160a01b03841681526060602082018190525f90615b53908301856150ed565b905063ffffffff83166040830152949350505050565b634e487b7160e01b5f52603160045260245ffd5b828152604060208201525f61186b6040830184614f0e565b5f82518060208501845e5f920191825250919050565b5f82615bc557634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220283a0a4f383d9daac7229b580defe060026a3ab9e204f046a0718b60a5ba843a64736f6c634300081e0033",
}

// DelegationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DelegationManagerMetaData.ABI instead.
var DelegationManagerABI = DelegationManagerMetaData.ABI

// DelegationManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DelegationManagerMetaData.Bin instead.
var DelegationManagerBin = DelegationManagerMetaData.Bin

// DeployDelegationManager deploys a new Ethereum contract, binding an instance of DelegationManager to it.
func DeployDelegationManager(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _eigenPodManager common.Address, _allocationManager common.Address, _pauserRegistry common.Address, _permissionController common.Address, _MIN_WITHDRAWAL_DELAY uint32, _version string) (common.Address, *types.Transaction, *DelegationManager, error) {
	parsed, err := DelegationManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DelegationManagerBin), backend, _strategyManager, _eigenPodManager, _allocationManager, _pauserRegistry, _permissionController, _MIN_WITHDRAWAL_DELAY, _version)
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
// Solidity: function getQueuedWithdrawal(bytes32 withdrawalRoot) view returns((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, uint256[] shares)
func (_DelegationManager *DelegationManagerCaller) GetQueuedWithdrawal(opts *bind.CallOpts, withdrawalRoot [32]byte) (struct {
	Withdrawal IDelegationManagerTypesWithdrawal
	Shares     []*big.Int
}, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "getQueuedWithdrawal", withdrawalRoot)

	outstruct := new(struct {
		Withdrawal IDelegationManagerTypesWithdrawal
		Shares     []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Withdrawal = *abi.ConvertType(out[0], new(IDelegationManagerTypesWithdrawal)).(*IDelegationManagerTypesWithdrawal)
	outstruct.Shares = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetQueuedWithdrawal is a free data retrieval call binding the contract method 0x5d975e88.
//
// Solidity: function getQueuedWithdrawal(bytes32 withdrawalRoot) view returns((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, uint256[] shares)
func (_DelegationManager *DelegationManagerSession) GetQueuedWithdrawal(withdrawalRoot [32]byte) (struct {
	Withdrawal IDelegationManagerTypesWithdrawal
	Shares     []*big.Int
}, error) {
	return _DelegationManager.Contract.GetQueuedWithdrawal(&_DelegationManager.CallOpts, withdrawalRoot)
}

// GetQueuedWithdrawal is a free data retrieval call binding the contract method 0x5d975e88.
//
// Solidity: function getQueuedWithdrawal(bytes32 withdrawalRoot) view returns((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, uint256[] shares)
func (_DelegationManager *DelegationManagerCallerSession) GetQueuedWithdrawal(withdrawalRoot [32]byte) (struct {
	Withdrawal IDelegationManagerTypesWithdrawal
	Shares     []*big.Int
}, error) {
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

// QueuedWithdrawals is a free data retrieval call binding the contract method 0x99f5371b.
//
// Solidity: function queuedWithdrawals(bytes32 withdrawalRoot) view returns((address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_DelegationManager *DelegationManagerCaller) QueuedWithdrawals(opts *bind.CallOpts, withdrawalRoot [32]byte) (IDelegationManagerTypesWithdrawal, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "queuedWithdrawals", withdrawalRoot)

	if err != nil {
		return *new(IDelegationManagerTypesWithdrawal), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelegationManagerTypesWithdrawal)).(*IDelegationManagerTypesWithdrawal)

	return out0, err

}

// QueuedWithdrawals is a free data retrieval call binding the contract method 0x99f5371b.
//
// Solidity: function queuedWithdrawals(bytes32 withdrawalRoot) view returns((address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_DelegationManager *DelegationManagerSession) QueuedWithdrawals(withdrawalRoot [32]byte) (IDelegationManagerTypesWithdrawal, error) {
	return _DelegationManager.Contract.QueuedWithdrawals(&_DelegationManager.CallOpts, withdrawalRoot)
}

// QueuedWithdrawals is a free data retrieval call binding the contract method 0x99f5371b.
//
// Solidity: function queuedWithdrawals(bytes32 withdrawalRoot) view returns((address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_DelegationManager *DelegationManagerCallerSession) QueuedWithdrawals(withdrawalRoot [32]byte) (IDelegationManagerTypesWithdrawal, error) {
	return _DelegationManager.Contract.QueuedWithdrawals(&_DelegationManager.CallOpts, withdrawalRoot)
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DelegationManager *DelegationManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DelegationManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DelegationManager *DelegationManagerSession) Version() (string, error) {
	return _DelegationManager.Contract.Version(&_DelegationManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DelegationManager *DelegationManagerCallerSession) Version() (string, error) {
	return _DelegationManager.Contract.Version(&_DelegationManager.CallOpts)
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
func (_DelegationManager *DelegationManagerTransactor) DelegateTo(opts *bind.TransactOpts, operator common.Address, approverSignatureAndExpiry ISignatureUtilsMixinTypesSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "delegateTo", operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationManager *DelegationManagerSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsMixinTypesSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationManager.Contract.DelegateTo(&_DelegationManager.TransactOpts, operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_DelegationManager *DelegationManagerTransactorSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsMixinTypesSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
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

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 initialPausedStatus) returns()
func (_DelegationManager *DelegationManagerTransactor) Initialize(opts *bind.TransactOpts, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "initialize", initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 initialPausedStatus) returns()
func (_DelegationManager *DelegationManagerSession) Initialize(initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.Initialize(&_DelegationManager.TransactOpts, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 initialPausedStatus) returns()
func (_DelegationManager *DelegationManagerTransactorSession) Initialize(initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _DelegationManager.Contract.Initialize(&_DelegationManager.TransactOpts, initialPausedStatus)
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
func (_DelegationManager *DelegationManagerTransactor) Redelegate(opts *bind.TransactOpts, newOperator common.Address, newOperatorApproverSig ISignatureUtilsMixinTypesSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "redelegate", newOperator, newOperatorApproverSig, approverSalt)
}

// Redelegate is a paid mutator transaction binding the contract method 0xa33a3433.
//
// Solidity: function redelegate(address newOperator, (bytes,uint256) newOperatorApproverSig, bytes32 approverSalt) returns(bytes32[] withdrawalRoots)
func (_DelegationManager *DelegationManagerSession) Redelegate(newOperator common.Address, newOperatorApproverSig ISignatureUtilsMixinTypesSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _DelegationManager.Contract.Redelegate(&_DelegationManager.TransactOpts, newOperator, newOperatorApproverSig, approverSalt)
}

// Redelegate is a paid mutator transaction binding the contract method 0xa33a3433.
//
// Solidity: function redelegate(address newOperator, (bytes,uint256) newOperatorApproverSig, bytes32 approverSalt) returns(bytes32[] withdrawalRoots)
func (_DelegationManager *DelegationManagerTransactorSession) Redelegate(newOperator common.Address, newOperatorApproverSig ISignatureUtilsMixinTypesSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
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

// SlashOperatorShares is a paid mutator transaction binding the contract method 0x5ae679a7.
//
// Solidity: function slashOperatorShares(address operator, (address,uint32) operatorSet, uint256 slashId, address strategy, uint64 prevMaxMagnitude, uint64 newMaxMagnitude) returns(uint256 totalDepositSharesToSlash)
func (_DelegationManager *DelegationManagerTransactor) SlashOperatorShares(opts *bind.TransactOpts, operator common.Address, operatorSet OperatorSet, slashId *big.Int, strategy common.Address, prevMaxMagnitude uint64, newMaxMagnitude uint64) (*types.Transaction, error) {
	return _DelegationManager.contract.Transact(opts, "slashOperatorShares", operator, operatorSet, slashId, strategy, prevMaxMagnitude, newMaxMagnitude)
}

// SlashOperatorShares is a paid mutator transaction binding the contract method 0x5ae679a7.
//
// Solidity: function slashOperatorShares(address operator, (address,uint32) operatorSet, uint256 slashId, address strategy, uint64 prevMaxMagnitude, uint64 newMaxMagnitude) returns(uint256 totalDepositSharesToSlash)
func (_DelegationManager *DelegationManagerSession) SlashOperatorShares(operator common.Address, operatorSet OperatorSet, slashId *big.Int, strategy common.Address, prevMaxMagnitude uint64, newMaxMagnitude uint64) (*types.Transaction, error) {
	return _DelegationManager.Contract.SlashOperatorShares(&_DelegationManager.TransactOpts, operator, operatorSet, slashId, strategy, prevMaxMagnitude, newMaxMagnitude)
}

// SlashOperatorShares is a paid mutator transaction binding the contract method 0x5ae679a7.
//
// Solidity: function slashOperatorShares(address operator, (address,uint32) operatorSet, uint256 slashId, address strategy, uint64 prevMaxMagnitude, uint64 newMaxMagnitude) returns(uint256 totalDepositSharesToSlash)
func (_DelegationManager *DelegationManagerTransactorSession) SlashOperatorShares(operator common.Address, operatorSet OperatorSet, slashId *big.Int, strategy common.Address, prevMaxMagnitude uint64, newMaxMagnitude uint64) (*types.Transaction, error) {
	return _DelegationManager.Contract.SlashOperatorShares(&_DelegationManager.TransactOpts, operator, operatorSet, slashId, strategy, prevMaxMagnitude, newMaxMagnitude)
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

// DelegationManagerOperatorSharesSlashedIterator is returned from FilterOperatorSharesSlashed and is used to iterate over the raw logs and unpacked data for OperatorSharesSlashed events raised by the DelegationManager contract.
type DelegationManagerOperatorSharesSlashedIterator struct {
	Event *DelegationManagerOperatorSharesSlashed // Event containing the contract specifics and raw log

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
func (it *DelegationManagerOperatorSharesSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegationManagerOperatorSharesSlashed)
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
		it.Event = new(DelegationManagerOperatorSharesSlashed)
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
func (it *DelegationManagerOperatorSharesSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegationManagerOperatorSharesSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegationManagerOperatorSharesSlashed represents a OperatorSharesSlashed event raised by the DelegationManager contract.
type DelegationManagerOperatorSharesSlashed struct {
	Operator           common.Address
	Strategy           common.Address
	TotalSlashedShares *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesSlashed is a free log retrieval operation binding the contract event 0xdd611f4ef63f4385f1756c86ce1f1f389a9013ba6fa07daba8528291bc2d3c30.
//
// Solidity: event OperatorSharesSlashed(address indexed operator, address strategy, uint256 totalSlashedShares)
func (_DelegationManager *DelegationManagerFilterer) FilterOperatorSharesSlashed(opts *bind.FilterOpts, operator []common.Address) (*DelegationManagerOperatorSharesSlashedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.FilterLogs(opts, "OperatorSharesSlashed", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegationManagerOperatorSharesSlashedIterator{contract: _DelegationManager.contract, event: "OperatorSharesSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesSlashed is a free log subscription operation binding the contract event 0xdd611f4ef63f4385f1756c86ce1f1f389a9013ba6fa07daba8528291bc2d3c30.
//
// Solidity: event OperatorSharesSlashed(address indexed operator, address strategy, uint256 totalSlashedShares)
func (_DelegationManager *DelegationManagerFilterer) WatchOperatorSharesSlashed(opts *bind.WatchOpts, sink chan<- *DelegationManagerOperatorSharesSlashed, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DelegationManager.contract.WatchLogs(opts, "OperatorSharesSlashed", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegationManagerOperatorSharesSlashed)
				if err := _DelegationManager.contract.UnpackLog(event, "OperatorSharesSlashed", log); err != nil {
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

// ParseOperatorSharesSlashed is a log parse operation binding the contract event 0xdd611f4ef63f4385f1756c86ce1f1f389a9013ba6fa07daba8528291bc2d3c30.
//
// Solidity: event OperatorSharesSlashed(address indexed operator, address strategy, uint256 totalSlashedShares)
func (_DelegationManager *DelegationManagerFilterer) ParseOperatorSharesSlashed(log types.Log) (*DelegationManagerOperatorSharesSlashed, error) {
	event := new(DelegationManagerOperatorSharesSlashed)
	if err := _DelegationManager.contract.UnpackLog(event, "OperatorSharesSlashed", log); err != nil {
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
