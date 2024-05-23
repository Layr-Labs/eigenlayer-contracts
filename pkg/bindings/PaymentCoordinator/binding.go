// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PaymentCoordinator

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

// IPaymentCoordinatorEarnerTreeMerkleLeaf is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorEarnerTreeMerkleLeaf struct {
	Earner          common.Address
	EarnerTokenRoot [32]byte
}

// IPaymentCoordinatorPaymentMerkleClaim is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorPaymentMerkleClaim struct {
	RootIndex       uint32
	EarnerIndex     uint32
	EarnerTreeProof []byte
	EarnerLeaf      IPaymentCoordinatorEarnerTreeMerkleLeaf
	TokenIndices    []uint32
	TokenTreeProofs [][]byte
	TokenLeaves     []IPaymentCoordinatorTokenTreeMerkleLeaf
}

// IPaymentCoordinatorRangePayment is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorRangePayment struct {
	StrategiesAndMultipliers []IPaymentCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IPaymentCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// IPaymentCoordinatorTokenTreeMerkleLeaf is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorTokenTreeMerkleLeaf struct {
	Token              common.Address
	CumulativeEarnings *big.Int
}

// PaymentCoordinatorMetaData contains all meta data concerning the PaymentCoordinator contract.
var PaymentCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_CALCULATION_INTERVAL_SECONDS\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_MAX_PAYMENT_DURATION\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_MAX_RETROACTIVE_LENGTH\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_MAX_FUTURE_LENGTH\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_GENESIS_PAYMENT_TIMESTAMP\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CALCULATION_INTERVAL_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GENESIS_PAYMENT_TIMESTAMP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FUTURE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PAYMENT_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_RETROACTIVE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activationDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateEarnerLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"calculateTokenLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.TokenTreeMerkleLeaf\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"checkClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.PaymentMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimerFor\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cumulativeClaimed\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currPaymentCalculationEndTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"distributionRoots\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymentCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributionRootsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRootIndexFromHash\",\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalOperatorCommissionBips\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_paymentUpdater\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_globalCommissionBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPayAllForRangeSubmitter\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRangePaymentForAllHash\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRangePaymentHash\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorCommissionBips\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payAllForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paymentNonce\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paymentUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.PaymentMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIPaymentCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setActivationDelay\",\"inputs\":[{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimerFor\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalOperatorCommission\",\"inputs\":[{\"name\":\"_globalCommissionBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPayAllForRangeSubmitter\",\"inputs\":[{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_newValue\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPaymentUpdater\",\"inputs\":[{\"name\":\"_paymentUpdater\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymentCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ActivationDelaySet\",\"inputs\":[{\"name\":\"oldActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimerForSet\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldClaimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootSubmitted\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"paymentCalculationEndTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalCommissionBipsSet\",\"inputs\":[{\"name\":\"oldGlobalCommissionBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newGlobalCommissionBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PayAllForRangeSubmitterSet\",\"inputs\":[{\"name\":\"payAllForRangeSubmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"newValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentClaimed\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"claimedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentUpdaterSet\",\"inputs\":[{\"name\":\"oldPaymentUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPaymentUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RangePaymentCreated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymentNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rangePaymentHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rangePayment\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPaymentCoordinator.RangePayment\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RangePaymentForAllCreated\",\"inputs\":[{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"paymentNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rangePaymentHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rangePayment\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPaymentCoordinator.RangePayment\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101806040523480156200001257600080fd5b5060405162003da838038062003da88339810160408190526200003591620002e4565b868686868686866200004885826200037e565b63ffffffff1615620000ed5760405162461bcd60e51b815260206004820152606060248201527f5061796d656e74436f6f7264696e61746f723a2047454e455349535f5041594d60448201527f454e545f54494d455354414d50206d7573742062652061206d756c7469706c6560648201527f206f662043414c43554c4154494f4e5f494e54455256414c5f5345434f4e4453608482015260a4015b60405180910390fd5b620000fc62015180866200037e565b63ffffffff16156200019d5760405162461bcd60e51b815260206004820152605760248201527f5061796d656e74436f6f7264696e61746f723a2043414c43554c4154494f4e5f60448201527f494e54455256414c5f5345434f4e4453206d7573742062652061206d756c746960648201527f706c65206f6620534e415053484f545f434144454e4345000000000000000000608482015260a401620000e4565b6001600160a01b0396871661012052949095166101405263ffffffff92831660805290821660a052811660c05291821660e0521661010052620001df620001f2565b5050466101605250620003b09350505050565b600054610100900460ff16156200025c5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401620000e4565b60005460ff9081161015620002af576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114620002c757600080fd5b50565b805163ffffffff81168114620002df57600080fd5b919050565b600080600080600080600060e0888a0312156200030057600080fd5b87516200030d81620002b1565b60208901519097506200032081620002b1565b95506200033060408901620002ca565b94506200034060608901620002ca565b93506200035060808901620002ca565b92506200036060a08901620002ca565b91506200037060c08901620002ca565b905092959891949750929550565b600063ffffffff80841680620003a457634e487b7160e01b600052601260045260246000fd5b92169190910692915050565b60805160a05160c05160e051610100516101205161014051610160516139616200044760003960006119a501526000818161044f015261228e015260006106e701526000818161040101526120d9015260008181610294015261217701526000818161042801526120880152600081816107210152611e2e01526000818161061301528181611ed70152611f9f01526139616000f3fe608060405234801561001057600080fd5b506004361061028a5760003560e01c806367ef85851161015c578063c8371b46116100ce578063ec1680de11610087578063ec1680de14610709578063ee6195971461071c578063f2fde38b14610743578063f698da2514610756578063f8cd84481461075e578063fabc1cbc1461077157600080fd5b8063c8371b461461065b578063d4540a5514610689578063d588cefa1461069c578063e221b245146106bc578063e810ce21146106cf578063ea4d3c9b146106e257600080fd5b8063886f119511610120578063886f1195146105cf5780638da5cb5b146105e25780639104c319146105f35780639d45c2811461060e578063a0169ddd14610635578063b57152841461064857600080fd5b806367ef85851461054f578063715018a61461056657806373f2fbea1461056e5780637b8f8b051461059c578063865c6953146105a457600080fd5b806337838ed011610200578063595c6a67116101b9578063595c6a67146104c15780635971b3f8146104c95780635ac86ab7146104fe5780635c975abb146105215780635e9d83481461052957806366d3b16b1461053c57600080fd5b806337838ed01461042357806339b70e381461044a5780633a8c0786146104715780633ccc861d146104885780633efe1db61461049b57806358baaa3e146104ae57600080fd5b8063149bc87211610252578063149bc8721461035357806318190f53146103745780631b4455161461038757806322f19a641461039a5780632b9f64a4146103bb5780632cfd45eb146103fc57600080fd5b806304a0c5021461028f578063092db007146102d057806310d67a2f146102f8578063136439dd1461030d578063146cd61d14610320575b600080fd5b6102b67f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020015b60405180910390f35b60cb546102e590600160e01b900461ffff1681565b60405161ffff90911681526020016102c7565b61030b610306366004613075565b610784565b005b61030b61031b366004613092565b610840565b61034361032e366004613075565b60d16020526000908152604090205460ff1681565b60405190151581526020016102c7565b6103666103613660046130c3565b61097f565b6040519081526020016102c7565b61030b610382366004613075565b6109f5565b61030b6103953660046130df565b610a06565b6102e56103a8366004613154565b505060cb54600160e01b900461ffff1690565b6103e46103c9366004613075565b60cc602052600090815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020016102c7565b6102b67f000000000000000000000000000000000000000000000000000000000000000081565b6102b67f000000000000000000000000000000000000000000000000000000000000000081565b6103e47f000000000000000000000000000000000000000000000000000000000000000081565b60cb546102b690600160a01b900463ffffffff1681565b61030b6104963660046131a0565b610ba1565b61030b6104a9366004613200565b610f56565b61030b6104bc36600461322c565b611256565b61030b611267565b6104dc6104d7366004613092565b61132e565b6040805193845263ffffffff92831660208501529116908201526060016102c7565b61034361050c366004613247565b606654600160ff9092169190911b9081161490565b606654610366565b61034361053736600461326a565b61136e565b60cb546103e4906001600160a01b031681565b60cb546102b690600160c01b900463ffffffff1681565b61030b6113ec565b61034361057c36600461329f565b60d060209081526000928352604080842090915290825290205460ff1681565b60ca54610366565b6103666105b2366004613154565b60cd60209081526000928352604080842090915290825290205481565b6065546103e4906001600160a01b031681565b6033546001600160a01b03166103e4565b6103e473beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b6102b67f000000000000000000000000000000000000000000000000000000000000000081565b61030b610643366004613075565b611400565b61030b6106563660046130df565b61145f565b61034361066936600461329f565b60cf60209081526000928352604080842090915290825290205460ff1681565b61030b6106973660046132e8565b61166e565b6103666106aa366004613075565b60ce6020526000908152604090205481565b61030b6106ca36600461335b565b6117b6565b6102b66106dd366004613092565b6117c7565b6103e47f000000000000000000000000000000000000000000000000000000000000000081565b61030b610717366004613384565b6118a9565b6102b67f000000000000000000000000000000000000000000000000000000000000000081565b61030b610751366004613075565b61192b565b6103666119a1565b61036661076c3660046130c3565b6119df565b61030b61077f366004613092565b6119f0565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107fb91906133b2565b6001600160a01b0316336001600160a01b0316146108345760405162461bcd60e51b815260040161082b906133cf565b60405180910390fd5b61083d81611b4c565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610888573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ac9190613419565b6108c85760405162461bcd60e51b815260040161082b90613436565b606654818116146109415760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c6974790000000000000000606482015260840161082b565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008061098f6020840184613075565b83602001356040516020016109d89392919060f89390931b6001600160f81b031916835260609190911b6bffffffffffffffffffffffff19166001830152601582015260350190565b604051602081830303815290604052805190602001209050919050565b6109fd611c43565b61083d81611c9d565b60665460009060019081161415610a2f5760405162461bcd60e51b815260040161082b9061347e565b60026097541415610a525760405162461bcd60e51b815260040161082b906134b5565b600260975560005b82811015610b965736848483818110610a7557610a756134ec565b9050602002810190610a879190613502565b33600081815260ce60209081526040808320549051949550939192610ab29290918591879101613647565b604051602081830303815290604052805190602001209050610ad383611cf9565b33600090815260cf602090815260408083208484529091529020805460ff19166001908117909155610b0690839061368d565b33600081815260ce602052604090819020929092559051829184917f2a0659fa4c38e0075469a0e0dd737e045dc316ffd6cb6e68755c119ee0882aea90610b4e9088906136a5565b60405180910390a4610b80333060408601803590610b6f9060208901613075565b6001600160a01b0316929190612434565b5050508080610b8e906136b8565b915050610a5a565b505060016097555050565b60665460029060049081161415610bca5760405162461bcd60e51b815260040161082b9061347e565b60026097541415610bed5760405162461bcd60e51b815260040161082b906134b5565b6002609755600060ca610c03602086018661322c565b63ffffffff1681548110610c1957610c196134ec565b6000918252602091829020604080516060810182526002909302909101805483526001015463ffffffff808216948401949094526401000000009004909216918101919091529050610c6b84826124a5565b6000610c7d6080860160608701613075565b6001600160a01b03808216600090815260cc60205260409020549192501680610ca35750805b336001600160a01b03821614610d215760405162461bcd60e51b815260206004820152603c60248201527f5061796d656e74436f6f7264696e61746f722e70726f63657373436c61696d3a60448201527f2063616c6c6572206973206e6f742076616c696420636c61696d657200000000606482015260840161082b565b60005b610d3160a08801886136d3565b9050811015610f485736610d4860e0890189613724565b83818110610d5857610d586134ec565b6001600160a01b038716600090815260cd602090815260408083209302949094019450929091508290610d8d90850185613075565b6001600160a01b03166001600160a01b0316815260200190815260200160002054905080826020013511610e475760405162461bcd60e51b815260206004820152605560248201527f5061796d656e74436f6f7264696e61746f722e70726f63657373436c61696d3a60448201527f2063756d756c61746976654561726e696e6773206d75737420626520677420746064820152741a185b8818dd5b5d5b185d1a5d9950db185a5b5959605a1b608482015260a40161082b565b6000610e5782602085013561376e565b6001600160a01b038716600090815260cd60209081526040822092935085018035929190610e859087613075565b6001600160a01b0316815260208082019290925260400160002091909155610ec7908a908390610eb790870187613075565b6001600160a01b0316919061274e565b86516001600160a01b03808b1691878216918916907fbff1e5a32b3f6d3b3c0a7e675ead2091fea820852f35a77abdd6d2420bec477890610f0b6020890189613075565b604080519283526001600160a01b039091166020830152810186905260600160405180910390a45050508080610f40906136b8565b915050610d24565b505060016097555050505050565b60665460039060089081161415610f7f5760405162461bcd60e51b815260040161082b9061347e565b60cb546001600160a01b03163314610ff65760405162461bcd60e51b815260206004820152603460248201527f5061796d656e74436f6f7264696e61746f723a2063616c6c6572206973206e6f6044820152733a103a3432903830bcb6b2b73a2ab83230ba32b960611b606482015260840161082b565b60cb5463ffffffff600160c01b9091048116908316116110925760405162461bcd60e51b815260206004820152604b60248201527f5061796d656e74436f6f7264696e61746f722e7375626d6974526f6f743a206e60448201527f657720726f6f74206d75737420626520666f72206e657765722063616c63756c60648201526a185d1959081c195c9a5bd960aa1b608482015260a40161082b565b428263ffffffff161061112b5760405162461bcd60e51b815260206004820152605560248201527f5061796d656e74436f6f7264696e61746f722e7375626d6974526f6f743a207060448201527f61796d656e7443616c63756c6174696f6e456e6454696d657374616d702063616064820152746e6e6f7420626520696e207468652066757475726560581b608482015260a40161082b565b60ca5460cb5460009061114b90600160a01b900463ffffffff1642613785565b6040805160608101825287815263ffffffff878116602080840182815286841685870181815260ca805460018101825560009190915296517f42d72674974f694b5f5159593243114d38a5c39c89d6b62fee061ff523240ee160029098029788015591517f42d72674974f694b5f5159593243114d38a5c39c89d6b62fee061ff523240ee29096018054925196861667ffffffffffffffff19909316929092176401000000009686169690960295909517905560cb805463ffffffff60c01b1916600160c01b840217905593519283529394508892908616917fecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08910160405180910390a45050505050565b61125e611c43565b61083d8161277e565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156112af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112d39190613419565b6112ef5760405162461bcd60e51b815260040161082b90613436565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60ca818154811061133e57600080fd5b60009182526020909120600290910201805460019091015490915063ffffffff8082169164010000000090041683565b60006113e48260ca611383602083018361322c565b63ffffffff1681548110611399576113996134ec565b6000918252602091829020604080516060810182526002909302909101805483526001015463ffffffff808216948401949094526401000000009004909216918101919091526124a5565b506001919050565b6113f4611c43565b6113fe60006127ef565b565b33600081815260cc602052604080822080546001600160a01b031981166001600160a01b038781169182179093559251911692839185917fbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca31291a4505050565b606654600190600290811614156114885760405162461bcd60e51b815260040161082b9061347e565b33600090815260d1602052604090205460ff166115185760405162461bcd60e51b815260206004820152604260248201527f5061796d656e74436f6f7264696e61746f723a2063616c6c6572206973206e6f60448201527f7420612076616c696420706179416c6c466f7252616e6765207375626d69747460648201526132b960f11b608482015260a40161082b565b6002609754141561153b5760405162461bcd60e51b815260040161082b906134b5565b600260975560005b82811015610b96573684848381811061155e5761155e6134ec565b90506020028101906115709190613502565b33600081815260ce6020908152604080832054905194955093919261159b9290918591879101613647565b6040516020818303038152906040528051906020012090506115bc83611cf9565b33600090815260d0602090815260408083208484529091529020805460ff191660019081179091556115ef90839061368d565b33600081815260ce602052604090819020929092559051829184917fbc0782940ec5871f66c5490ef957f44d19a9adac1dac18b946ad0dd6579c30d6906116379088906136a5565b60405180910390a4611658333060408601803590610b6f9060208901613075565b5050508080611666906136b8565b915050611543565b600054610100900460ff161580801561168e5750600054600160ff909116105b806116a85750303b1580156116a8575060005460ff166001145b61170b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161082b565b6000805460ff19166001179055801561172e576000805461ff0019166101001790555b611736612841565b60c95561174386866128d8565b61174c876127ef565b61175584611c9d565b61175e8361277e565b611767826129c2565b80156117ad576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b6117be611c43565b61083d816129c2565b60ca546000905b63ffffffff81161561183a578260ca6117e86001846137ad565b63ffffffff16815481106117fe576117fe6134ec565b9060005260206000209060020201600001541415611828576118216001826137ad565b9392505050565b80611832816137d2565b9150506117ce565b5060405162461bcd60e51b815260206004820152603760248201527f5061796d656e74436f6f7264696e61746f722e676574526f6f74496e6465784660448201527f726f6d486173683a20726f6f74206e6f7420666f756e64000000000000000000606482015260840161082b565b6118b1611c43565b6001600160a01b038216600081815260d1602052604080822054905160ff9091169284151592841515927f17b06be02af2803593116eb96121b9c6e8bee1cc1b145e7c31c19c180e86189b9190a4506001600160a01b0391909116600090815260d160205260409020805460ff1916911515919091179055565b611933611c43565b6001600160a01b0381166119985760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161082b565b61083d816127ef565b60007f00000000000000000000000000000000000000000000000000000000000000004614156119d2575060c95490565b6119da612841565b905090565b6000600161098f6020840184613075565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a6791906133b2565b6001600160a01b0316336001600160a01b031614611a975760405162461bcd60e51b815260040161082b906133cf565b606654198119606654191614611b155760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c6974790000000000000000606482015260840161082b565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610974565b6001600160a01b038116611bda5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a40161082b565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6033546001600160a01b031633146113fe5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161082b565b60cb546040516001600160a01b038084169216907f07d2890b3eb1206e7c3cb6bf8d46da31385ace3ce99abf85e5b690c83aa4967890600090a360cb80546001600160a01b0319166001600160a01b0392909216919091179055565b6000611d058280613724565b905011611d5d5760405162461bcd60e51b8152602060048201526032602482015260008051602061390c833981519152604482015271081b9bc81cdd1c985d1959da595cc81cd95d60721b606482015260840161082b565b6000816040013511611dbb5760405162461bcd60e51b8152602060048201526033602482015260008051602061390c833981519152604482015272020616d6f756e742063616e6e6f74206265203606c1b606482015260840161082b565b6f4b3b4ca85a86c47a098a223fffffffff81604001351115611e275760405162461bcd60e51b8152602060048201526031602482015260008051602061390c83398151915260448201527020616d6f756e7420746f6f206c6172676560781b606482015260840161082b565b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016611e5e60a083016080840161322c565b63ffffffff161115611ed55760405162461bcd60e51b8152602060048201526046602482015260008051602061390c83398151915260448201527f206475726174696f6e2065786365656473204d41585f5041594d454e545f44556064820152652920aa24a7a760d11b608482015260a40161082b565b7f0000000000000000000000000000000000000000000000000000000000000000611f0660a083016080840161322c565b611f109190613808565b63ffffffff1615611f9d5760405162461bcd60e51b815260206004820152605c602482015260008051602061390c83398151915260448201527f206475726174696f6e206d7573742062652061206d756c7469706c65206f662060648201527f43414c43554c4154494f4e5f494e54455256414c5f5345434f4e445300000000608482015260a40161082b565b7f0000000000000000000000000000000000000000000000000000000000000000611fce608083016060840161322c565b611fd89190613808565b63ffffffff16156120705760405162461bcd60e51b8152602060048201526062602482015260008051602061390c83398151915260448201527f20737461727454696d657374616d70206d7573742062652061206d756c74697060648201527f6c65206f662043414c43554c4154494f4e5f494e54455256414c5f5345434f4e608482015261445360f01b60a482015260c40161082b565b612080608082016060830161322c565b63ffffffff167f000000000000000000000000000000000000000000000000000000000000000063ffffffff16426120b8919061376e565b1115801561210157506120d1608082016060830161322c565b63ffffffff167f000000000000000000000000000000000000000000000000000000000000000063ffffffff1611155b61216d5760405162461bcd60e51b8152602060048201526043602482015260008051602061390c83398151915260448201527f20737461727454696d657374616d7020746f6f2066617220696e207468652070606482015262185cdd60ea1b608482015260a40161082b565b61219d63ffffffff7f0000000000000000000000000000000000000000000000000000000000000000164261368d565b6121ad608083016060840161322c565b63ffffffff1611156122235760405162461bcd60e51b8152602060048201526045602482015260008051602061390c83398151915260448201527f20737461727454696d657374616d7020746f6f2066617220696e207468652066606482015264757475726560d81b608482015260a40161082b565b6000805b6122318380613724565b905081101561242f5760006122468480613724565b83818110612256576122566134ec565b61226c9260206040909202019081019150613075565b60405163198f077960e21b81526001600160a01b0380831660048301529192507f00000000000000000000000000000000000000000000000000000000000000009091169063663c1de490602401602060405180830381865afa1580156122d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122fb9190613419565b8061232257506001600160a01b03811673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0145b6123825760405162461bcd60e51b815260206004820152603c602482015260008051602061390c83398151915260448201527f20696e76616c696420737472617465677920636f6e7369646572656400000000606482015260840161082b565b806001600160a01b0316836001600160a01b03161061241d5760405162461bcd60e51b815260206004820152605b602482015260008051602061390c83398151915260448201527f2073747261746567696573206d75737420626520696e20617363656e64696e6760648201527f206f7264657220746f2068616e646c65206475706c6963617465730000000000608482015260a40161082b565b9150612428816136b8565b9050612227565b505050565b6040516001600160a01b038085166024830152831660448201526064810182905261249f9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152612a2d565b50505050565b806040015163ffffffff1642101561251e5760405162461bcd60e51b815260206004820152603660248201527f5061796d656e74436f6f7264696e61746f722e5f636865636b436c61696d3a206044820152751c9bdbdd081b9bdd081858dd1a5d985d1959081e595d60521b606482015260840161082b565b61252b60c08301836136d3565b905061253a60a08401846136d3565b9050146125c45760405162461bcd60e51b815260206004820152604c60248201527f5061796d656e74436f6f7264696e61746f722e5f636865636b436c61696d3a2060448201527f746f6b656e496e646963657320616e6420746f6b656e50726f6f6673206c656e60648201526b0cee8d040dad2e6dac2e8c6d60a31b608482015260a40161082b565b6125d160e0830183613724565b90506125e060c08401846136d3565b9050146126685760405162461bcd60e51b815260206004820152604a60248201527f5061796d656e74436f6f7264696e61746f722e5f636865636b436c61696d3a2060448201527f746f6b656e5472656550726f6f667320616e64206c6561766573206c656e67746064820152690d040dad2e6dac2e8c6d60b31b608482015260a40161082b565b80516126949061267e604085016020860161322c565b61268b604086018661382b565b86606001612aff565b60005b6126a460a08401846136d3565b905081101561242f5761273e60808401356126c260a08601866136d3565b848181106126d2576126d26134ec565b90506020020160208101906126e7919061322c565b6126f460c08701876136d3565b85818110612704576127046134ec565b9050602002810190612716919061382b565b61272360e0890189613724565b87818110612733576127336134ec565b905060400201612c73565b612747816136b8565b9050612697565b6040516001600160a01b03831660248201526044810182905261242f90849063a9059cbb60e01b90606401612468565b60cb546040805163ffffffff600160a01b9093048316815291831660208301527faf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3910160405180910390a160cb805463ffffffff909216600160a01b0263ffffffff60a01b19909216919091179055565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b6065546001600160a01b03161580156128f957506001600160a01b03821615155b61297b5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a40161082b565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26129be82611b4c565b5050565b60cb546040805161ffff600160e01b9093048316815291831660208301527f8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06910160405180910390a160cb805461ffff909216600160e01b0261ffff60e01b19909216919091179055565b6000612a82826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316612dc49092919063ffffffff16565b80519091501561242f5780806020019051810190612aa09190613419565b61242f5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161082b565b612b0a602083613872565b6001901b8463ffffffff1610612b945760405162461bcd60e51b815260206004820152604360248201527f5061796d656e74436f6f7264696e61746f722e5f7665726966794561726e657260448201527f436c61696d50726f6f663a20696e76616c6964206561726e65724c656166496e6064820152620c8caf60eb1b608482015260a40161082b565b6000612b9f8261097f565b9050612bea84848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a92508591505063ffffffff8916612ddb565b612c6b5760405162461bcd60e51b815260206004820152604660248201527f5061796d656e74436f6f7264696e61746f722e5f7665726966794561726e657260448201527f436c61696d50726f6f663a20696e76616c6964206561726e657220636c61696d60648201526510383937b7b360d11b608482015260a40161082b565b505050505050565b612c7e602083613872565b6001901b8463ffffffff1610612cfc5760405162461bcd60e51b815260206004820152603c60248201527f5061796d656e74436f6f7264696e61746f722e5f766572696679546f6b656e4360448201527f6c61696d3a20696e76616c696420746f6b656e4c656166496e64657800000000606482015260840161082b565b6000612d07826119df565b9050612d5284848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a92508591505063ffffffff8916612ddb565b612c6b5760405162461bcd60e51b815260206004820152603f60248201527f5061796d656e74436f6f7264696e61746f722e5f766572696679546f6b656e4360448201527f6c61696d3a20696e76616c696420746f6b656e20636c61696d2070726f6f6600606482015260840161082b565b6060612dd38484600085612df3565b949350505050565b600083612de9868585612f24565b1495945050505050565b606082471015612e545760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161082b565b6001600160a01b0385163b612eab5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161082b565b600080866001600160a01b03168587604051612ec791906138b2565b60006040518083038185875af1925050503d8060008114612f04576040519150601f19603f3d011682016040523d82523d6000602084013e612f09565b606091505b5091509150612f19828286613027565b979650505050505050565b600060208451612f3491906138c4565b15612fbb5760405162461bcd60e51b815260206004820152604b60248201527f4d65726b6c652e70726f63657373496e636c7573696f6e50726f6f664b65636360448201527f616b3a2070726f6f66206c656e6774682073686f756c642062652061206d756c60648201526a3a34b836329037b310199960a91b608482015260a40161082b565b8260205b8551811161301e57612fd26002856138c4565b612ff35781600052808601516020526040600020915060028404935061300c565b8086015160005281602052604060002091506002840493505b61301760208261368d565b9050612fbf565b50949350505050565b60608315613036575081611821565b8251156130465782518084602001fd5b8160405162461bcd60e51b815260040161082b91906138d8565b6001600160a01b038116811461083d57600080fd5b60006020828403121561308757600080fd5b813561182181613060565b6000602082840312156130a457600080fd5b5035919050565b6000604082840312156130bd57600080fd5b50919050565b6000604082840312156130d557600080fd5b61182183836130ab565b600080602083850312156130f257600080fd5b823567ffffffffffffffff8082111561310a57600080fd5b818501915085601f83011261311e57600080fd5b81358181111561312d57600080fd5b8660208260051b850101111561314257600080fd5b60209290920196919550909350505050565b6000806040838503121561316757600080fd5b823561317281613060565b9150602083013561318281613060565b809150509250929050565b600061010082840312156130bd57600080fd5b600080604083850312156131b357600080fd5b823567ffffffffffffffff8111156131ca57600080fd5b6131d68582860161318d565b925050602083013561318281613060565b803563ffffffff811681146131fb57600080fd5b919050565b6000806040838503121561321357600080fd5b82359150613223602084016131e7565b90509250929050565b60006020828403121561323e57600080fd5b611821826131e7565b60006020828403121561325957600080fd5b813560ff8116811461182157600080fd5b60006020828403121561327c57600080fd5b813567ffffffffffffffff81111561329357600080fd5b612dd38482850161318d565b600080604083850312156132b257600080fd5b82356132bd81613060565b946020939093013593505050565b80356131fb81613060565b803561ffff811681146131fb57600080fd5b60008060008060008060c0878903121561330157600080fd5b863561330c81613060565b9550602087013561331c81613060565b945060408701359350606087013561333381613060565b9250613341608088016131e7565b915061334f60a088016132d6565b90509295509295509295565b60006020828403121561336d57600080fd5b611821826132d6565b801515811461083d57600080fd5b6000806040838503121561339757600080fd5b82356133a281613060565b9150602083013561318281613376565b6000602082840312156133c457600080fd5b815161182181613060565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561342b57600080fd5b815161182181613376565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b60208082526019908201527f5061757361626c653a20696e6465782069732070617573656400000000000000604082015260600190565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b634e487b7160e01b600052603260045260246000fd5b60008235609e1983360301811261351857600080fd5b9190910192915050565b818352600060208085019450826000805b8681101561358757823561354681613060565b6001600160a01b03168852828401356bffffffffffffffffffffffff811680821461356f578384fd5b89860152506040978801979290920191600101613533565b50959695505050505050565b60008135601e198336030181126135a957600080fd5b8201803567ffffffffffffffff8111156135c257600080fd5b8060061b36038413156135d457600080fd5b60a085526135e960a086018260208501613522565b9150506135f8602084016132cb565b6001600160a01b031660208501526040838101359085015261361c606084016131e7565b63ffffffff166060850152613633608084016131e7565b63ffffffff81166080860152509392505050565b60018060a01b038416815282602082015260606040820152600061366e6060830184613593565b95945050505050565b634e487b7160e01b600052601160045260246000fd5b600082198211156136a0576136a0613677565b500190565b6020815260006118216020830184613593565b60006000198214156136cc576136cc613677565b5060010190565b6000808335601e198436030181126136ea57600080fd5b83018035915067ffffffffffffffff82111561370557600080fd5b6020019150600581901b360382131561371d57600080fd5b9250929050565b6000808335601e1984360301811261373b57600080fd5b83018035915067ffffffffffffffff82111561375657600080fd5b6020019150600681901b360382131561371d57600080fd5b60008282101561378057613780613677565b500390565b600063ffffffff8083168185168083038211156137a4576137a4613677565b01949350505050565b600063ffffffff838116908316818110156137ca576137ca613677565b039392505050565b600063ffffffff8216806137e8576137e8613677565b6000190192915050565b634e487b7160e01b600052601260045260246000fd5b600063ffffffff8084168061381f5761381f6137f2565b92169190910692915050565b6000808335601e1984360301811261384257600080fd5b83018035915067ffffffffffffffff82111561385d57600080fd5b60200191503681900382131561371d57600080fd5b600082613881576138816137f2565b500490565b60005b838110156138a1578181015183820152602001613889565b8381111561249f5750506000910152565b60008251613518818460208701613886565b6000826138d3576138d36137f2565b500690565b60208152600082518060208401526138f7816040850160208701613886565b601f01601f1916919091016040019291505056fe5061796d656e74436f6f7264696e61746f722e5f706179466f7252616e67653aa26469706673582212209769e876fa54eaa8705c0622de1ea7a50d6bcd9bcc443f25be008e17d65b4f4c64736f6c634300080c0033",
}

// PaymentCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymentCoordinatorMetaData.ABI instead.
var PaymentCoordinatorABI = PaymentCoordinatorMetaData.ABI

// PaymentCoordinatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PaymentCoordinatorMetaData.Bin instead.
var PaymentCoordinatorBin = PaymentCoordinatorMetaData.Bin

// DeployPaymentCoordinator deploys a new Ethereum contract, binding an instance of PaymentCoordinator to it.
func DeployPaymentCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _delegationManager common.Address, _strategyManager common.Address, _CALCULATION_INTERVAL_SECONDS uint32, _MAX_PAYMENT_DURATION uint32, _MAX_RETROACTIVE_LENGTH uint32, _MAX_FUTURE_LENGTH uint32, _GENESIS_PAYMENT_TIMESTAMP uint32) (common.Address, *types.Transaction, *PaymentCoordinator, error) {
	parsed, err := PaymentCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PaymentCoordinatorBin), backend, _delegationManager, _strategyManager, _CALCULATION_INTERVAL_SECONDS, _MAX_PAYMENT_DURATION, _MAX_RETROACTIVE_LENGTH, _MAX_FUTURE_LENGTH, _GENESIS_PAYMENT_TIMESTAMP)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PaymentCoordinator{PaymentCoordinatorCaller: PaymentCoordinatorCaller{contract: contract}, PaymentCoordinatorTransactor: PaymentCoordinatorTransactor{contract: contract}, PaymentCoordinatorFilterer: PaymentCoordinatorFilterer{contract: contract}}, nil
}

// PaymentCoordinator is an auto generated Go binding around an Ethereum contract.
type PaymentCoordinator struct {
	PaymentCoordinatorCaller     // Read-only binding to the contract
	PaymentCoordinatorTransactor // Write-only binding to the contract
	PaymentCoordinatorFilterer   // Log filterer for contract events
}

// PaymentCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentCoordinatorSession struct {
	Contract     *PaymentCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PaymentCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentCoordinatorCallerSession struct {
	Contract *PaymentCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PaymentCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentCoordinatorTransactorSession struct {
	Contract     *PaymentCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PaymentCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentCoordinatorRaw struct {
	Contract *PaymentCoordinator // Generic contract binding to access the raw methods on
}

// PaymentCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentCoordinatorCallerRaw struct {
	Contract *PaymentCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentCoordinatorTransactorRaw struct {
	Contract *PaymentCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentCoordinator creates a new instance of PaymentCoordinator, bound to a specific deployed contract.
func NewPaymentCoordinator(address common.Address, backend bind.ContractBackend) (*PaymentCoordinator, error) {
	contract, err := bindPaymentCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinator{PaymentCoordinatorCaller: PaymentCoordinatorCaller{contract: contract}, PaymentCoordinatorTransactor: PaymentCoordinatorTransactor{contract: contract}, PaymentCoordinatorFilterer: PaymentCoordinatorFilterer{contract: contract}}, nil
}

// NewPaymentCoordinatorCaller creates a new read-only instance of PaymentCoordinator, bound to a specific deployed contract.
func NewPaymentCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*PaymentCoordinatorCaller, error) {
	contract, err := bindPaymentCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorCaller{contract: contract}, nil
}

// NewPaymentCoordinatorTransactor creates a new write-only instance of PaymentCoordinator, bound to a specific deployed contract.
func NewPaymentCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentCoordinatorTransactor, error) {
	contract, err := bindPaymentCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorTransactor{contract: contract}, nil
}

// NewPaymentCoordinatorFilterer creates a new log filterer instance of PaymentCoordinator, bound to a specific deployed contract.
func NewPaymentCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentCoordinatorFilterer, error) {
	contract, err := bindPaymentCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorFilterer{contract: contract}, nil
}

// bindPaymentCoordinator binds a generic wrapper to an already deployed contract.
func bindPaymentCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PaymentCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentCoordinator *PaymentCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentCoordinator.Contract.PaymentCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentCoordinator *PaymentCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.PaymentCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentCoordinator *PaymentCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.PaymentCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentCoordinator *PaymentCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentCoordinator *PaymentCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentCoordinator *PaymentCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.contract.Transact(opts, method, params...)
}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCaller) CALCULATIONINTERVALSECONDS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "CALCULATION_INTERVAL_SECONDS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorSession) CALCULATIONINTERVALSECONDS() (uint32, error) {
	return _PaymentCoordinator.Contract.CALCULATIONINTERVALSECONDS(&_PaymentCoordinator.CallOpts)
}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) CALCULATIONINTERVALSECONDS() (uint32, error) {
	return _PaymentCoordinator.Contract.CALCULATIONINTERVALSECONDS(&_PaymentCoordinator.CallOpts)
}

// GENESISPAYMENTTIMESTAMP is a free data retrieval call binding the contract method 0x2cfd45eb.
//
// Solidity: function GENESIS_PAYMENT_TIMESTAMP() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCaller) GENESISPAYMENTTIMESTAMP(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "GENESIS_PAYMENT_TIMESTAMP")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GENESISPAYMENTTIMESTAMP is a free data retrieval call binding the contract method 0x2cfd45eb.
//
// Solidity: function GENESIS_PAYMENT_TIMESTAMP() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorSession) GENESISPAYMENTTIMESTAMP() (uint32, error) {
	return _PaymentCoordinator.Contract.GENESISPAYMENTTIMESTAMP(&_PaymentCoordinator.CallOpts)
}

// GENESISPAYMENTTIMESTAMP is a free data retrieval call binding the contract method 0x2cfd45eb.
//
// Solidity: function GENESIS_PAYMENT_TIMESTAMP() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) GENESISPAYMENTTIMESTAMP() (uint32, error) {
	return _PaymentCoordinator.Contract.GENESISPAYMENTTIMESTAMP(&_PaymentCoordinator.CallOpts)
}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCaller) MAXFUTURELENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "MAX_FUTURE_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorSession) MAXFUTURELENGTH() (uint32, error) {
	return _PaymentCoordinator.Contract.MAXFUTURELENGTH(&_PaymentCoordinator.CallOpts)
}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) MAXFUTURELENGTH() (uint32, error) {
	return _PaymentCoordinator.Contract.MAXFUTURELENGTH(&_PaymentCoordinator.CallOpts)
}

// MAXPAYMENTDURATION is a free data retrieval call binding the contract method 0xee619597.
//
// Solidity: function MAX_PAYMENT_DURATION() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCaller) MAXPAYMENTDURATION(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "MAX_PAYMENT_DURATION")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXPAYMENTDURATION is a free data retrieval call binding the contract method 0xee619597.
//
// Solidity: function MAX_PAYMENT_DURATION() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorSession) MAXPAYMENTDURATION() (uint32, error) {
	return _PaymentCoordinator.Contract.MAXPAYMENTDURATION(&_PaymentCoordinator.CallOpts)
}

// MAXPAYMENTDURATION is a free data retrieval call binding the contract method 0xee619597.
//
// Solidity: function MAX_PAYMENT_DURATION() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) MAXPAYMENTDURATION() (uint32, error) {
	return _PaymentCoordinator.Contract.MAXPAYMENTDURATION(&_PaymentCoordinator.CallOpts)
}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCaller) MAXRETROACTIVELENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "MAX_RETROACTIVE_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorSession) MAXRETROACTIVELENGTH() (uint32, error) {
	return _PaymentCoordinator.Contract.MAXRETROACTIVELENGTH(&_PaymentCoordinator.CallOpts)
}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) MAXRETROACTIVELENGTH() (uint32, error) {
	return _PaymentCoordinator.Contract.MAXRETROACTIVELENGTH(&_PaymentCoordinator.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCaller) ActivationDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "activationDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorSession) ActivationDelay() (uint32, error) {
	return _PaymentCoordinator.Contract.ActivationDelay(&_PaymentCoordinator.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) ActivationDelay() (uint32, error) {
	return _PaymentCoordinator.Contract.ActivationDelay(&_PaymentCoordinator.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorCaller) BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "beaconChainETHStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorSession) BeaconChainETHStrategy() (common.Address, error) {
	return _PaymentCoordinator.Contract.BeaconChainETHStrategy(&_PaymentCoordinator.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _PaymentCoordinator.Contract.BeaconChainETHStrategy(&_PaymentCoordinator.CallOpts)
}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_PaymentCoordinator *PaymentCoordinatorCaller) CalculateEarnerLeafHash(opts *bind.CallOpts, leaf IPaymentCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "calculateEarnerLeafHash", leaf)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_PaymentCoordinator *PaymentCoordinatorSession) CalculateEarnerLeafHash(leaf IPaymentCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _PaymentCoordinator.Contract.CalculateEarnerLeafHash(&_PaymentCoordinator.CallOpts, leaf)
}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) CalculateEarnerLeafHash(leaf IPaymentCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _PaymentCoordinator.Contract.CalculateEarnerLeafHash(&_PaymentCoordinator.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_PaymentCoordinator *PaymentCoordinatorCaller) CalculateTokenLeafHash(opts *bind.CallOpts, leaf IPaymentCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "calculateTokenLeafHash", leaf)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_PaymentCoordinator *PaymentCoordinatorSession) CalculateTokenLeafHash(leaf IPaymentCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	return _PaymentCoordinator.Contract.CalculateTokenLeafHash(&_PaymentCoordinator.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) CalculateTokenLeafHash(leaf IPaymentCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	return _PaymentCoordinator.Contract.CalculateTokenLeafHash(&_PaymentCoordinator.CallOpts, leaf)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_PaymentCoordinator *PaymentCoordinatorCaller) CheckClaim(opts *bind.CallOpts, claim IPaymentCoordinatorPaymentMerkleClaim) (bool, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "checkClaim", claim)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_PaymentCoordinator *PaymentCoordinatorSession) CheckClaim(claim IPaymentCoordinatorPaymentMerkleClaim) (bool, error) {
	return _PaymentCoordinator.Contract.CheckClaim(&_PaymentCoordinator.CallOpts, claim)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) CheckClaim(claim IPaymentCoordinatorPaymentMerkleClaim) (bool, error) {
	return _PaymentCoordinator.Contract.CheckClaim(&_PaymentCoordinator.CallOpts, claim)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address ) view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorCaller) ClaimerFor(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "claimerFor", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address ) view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorSession) ClaimerFor(arg0 common.Address) (common.Address, error) {
	return _PaymentCoordinator.Contract.ClaimerFor(&_PaymentCoordinator.CallOpts, arg0)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address ) view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) ClaimerFor(arg0 common.Address) (common.Address, error) {
	return _PaymentCoordinator.Contract.ClaimerFor(&_PaymentCoordinator.CallOpts, arg0)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address , address ) view returns(uint256)
func (_PaymentCoordinator *PaymentCoordinatorCaller) CumulativeClaimed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "cumulativeClaimed", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address , address ) view returns(uint256)
func (_PaymentCoordinator *PaymentCoordinatorSession) CumulativeClaimed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _PaymentCoordinator.Contract.CumulativeClaimed(&_PaymentCoordinator.CallOpts, arg0, arg1)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address , address ) view returns(uint256)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) CumulativeClaimed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _PaymentCoordinator.Contract.CumulativeClaimed(&_PaymentCoordinator.CallOpts, arg0, arg1)
}

// CurrPaymentCalculationEndTimestamp is a free data retrieval call binding the contract method 0x67ef8585.
//
// Solidity: function currPaymentCalculationEndTimestamp() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCaller) CurrPaymentCalculationEndTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "currPaymentCalculationEndTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurrPaymentCalculationEndTimestamp is a free data retrieval call binding the contract method 0x67ef8585.
//
// Solidity: function currPaymentCalculationEndTimestamp() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorSession) CurrPaymentCalculationEndTimestamp() (uint32, error) {
	return _PaymentCoordinator.Contract.CurrPaymentCalculationEndTimestamp(&_PaymentCoordinator.CallOpts)
}

// CurrPaymentCalculationEndTimestamp is a free data retrieval call binding the contract method 0x67ef8585.
//
// Solidity: function currPaymentCalculationEndTimestamp() view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) CurrPaymentCalculationEndTimestamp() (uint32, error) {
	return _PaymentCoordinator.Contract.CurrPaymentCalculationEndTimestamp(&_PaymentCoordinator.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorSession) DelegationManager() (common.Address, error) {
	return _PaymentCoordinator.Contract.DelegationManager(&_PaymentCoordinator.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) DelegationManager() (common.Address, error) {
	return _PaymentCoordinator.Contract.DelegationManager(&_PaymentCoordinator.CallOpts)
}

// DistributionRoots is a free data retrieval call binding the contract method 0x5971b3f8.
//
// Solidity: function distributionRoots(uint256 ) view returns(bytes32 root, uint32 paymentCalculationEndTimestamp, uint32 activatedAt)
func (_PaymentCoordinator *PaymentCoordinatorCaller) DistributionRoots(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Root                           [32]byte
	PaymentCalculationEndTimestamp uint32
	ActivatedAt                    uint32
}, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "distributionRoots", arg0)

	outstruct := new(struct {
		Root                           [32]byte
		PaymentCalculationEndTimestamp uint32
		ActivatedAt                    uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.PaymentCalculationEndTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ActivatedAt = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// DistributionRoots is a free data retrieval call binding the contract method 0x5971b3f8.
//
// Solidity: function distributionRoots(uint256 ) view returns(bytes32 root, uint32 paymentCalculationEndTimestamp, uint32 activatedAt)
func (_PaymentCoordinator *PaymentCoordinatorSession) DistributionRoots(arg0 *big.Int) (struct {
	Root                           [32]byte
	PaymentCalculationEndTimestamp uint32
	ActivatedAt                    uint32
}, error) {
	return _PaymentCoordinator.Contract.DistributionRoots(&_PaymentCoordinator.CallOpts, arg0)
}

// DistributionRoots is a free data retrieval call binding the contract method 0x5971b3f8.
//
// Solidity: function distributionRoots(uint256 ) view returns(bytes32 root, uint32 paymentCalculationEndTimestamp, uint32 activatedAt)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) DistributionRoots(arg0 *big.Int) (struct {
	Root                           [32]byte
	PaymentCalculationEndTimestamp uint32
	ActivatedAt                    uint32
}, error) {
	return _PaymentCoordinator.Contract.DistributionRoots(&_PaymentCoordinator.CallOpts, arg0)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_PaymentCoordinator *PaymentCoordinatorCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_PaymentCoordinator *PaymentCoordinatorSession) DomainSeparator() ([32]byte, error) {
	return _PaymentCoordinator.Contract.DomainSeparator(&_PaymentCoordinator.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) DomainSeparator() ([32]byte, error) {
	return _PaymentCoordinator.Contract.DomainSeparator(&_PaymentCoordinator.CallOpts)
}

// GetDistributionRootsLength is a free data retrieval call binding the contract method 0x7b8f8b05.
//
// Solidity: function getDistributionRootsLength() view returns(uint256)
func (_PaymentCoordinator *PaymentCoordinatorCaller) GetDistributionRootsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "getDistributionRootsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistributionRootsLength is a free data retrieval call binding the contract method 0x7b8f8b05.
//
// Solidity: function getDistributionRootsLength() view returns(uint256)
func (_PaymentCoordinator *PaymentCoordinatorSession) GetDistributionRootsLength() (*big.Int, error) {
	return _PaymentCoordinator.Contract.GetDistributionRootsLength(&_PaymentCoordinator.CallOpts)
}

// GetDistributionRootsLength is a free data retrieval call binding the contract method 0x7b8f8b05.
//
// Solidity: function getDistributionRootsLength() view returns(uint256)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) GetDistributionRootsLength() (*big.Int, error) {
	return _PaymentCoordinator.Contract.GetDistributionRootsLength(&_PaymentCoordinator.CallOpts)
}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCaller) GetRootIndexFromHash(opts *bind.CallOpts, rootHash [32]byte) (uint32, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "getRootIndexFromHash", rootHash)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorSession) GetRootIndexFromHash(rootHash [32]byte) (uint32, error) {
	return _PaymentCoordinator.Contract.GetRootIndexFromHash(&_PaymentCoordinator.CallOpts, rootHash)
}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) GetRootIndexFromHash(rootHash [32]byte) (uint32, error) {
	return _PaymentCoordinator.Contract.GetRootIndexFromHash(&_PaymentCoordinator.CallOpts, rootHash)
}

// GlobalOperatorCommissionBips is a free data retrieval call binding the contract method 0x092db007.
//
// Solidity: function globalOperatorCommissionBips() view returns(uint16)
func (_PaymentCoordinator *PaymentCoordinatorCaller) GlobalOperatorCommissionBips(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "globalOperatorCommissionBips")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GlobalOperatorCommissionBips is a free data retrieval call binding the contract method 0x092db007.
//
// Solidity: function globalOperatorCommissionBips() view returns(uint16)
func (_PaymentCoordinator *PaymentCoordinatorSession) GlobalOperatorCommissionBips() (uint16, error) {
	return _PaymentCoordinator.Contract.GlobalOperatorCommissionBips(&_PaymentCoordinator.CallOpts)
}

// GlobalOperatorCommissionBips is a free data retrieval call binding the contract method 0x092db007.
//
// Solidity: function globalOperatorCommissionBips() view returns(uint16)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) GlobalOperatorCommissionBips() (uint16, error) {
	return _PaymentCoordinator.Contract.GlobalOperatorCommissionBips(&_PaymentCoordinator.CallOpts)
}

// IsPayAllForRangeSubmitter is a free data retrieval call binding the contract method 0x146cd61d.
//
// Solidity: function isPayAllForRangeSubmitter(address ) view returns(bool)
func (_PaymentCoordinator *PaymentCoordinatorCaller) IsPayAllForRangeSubmitter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "isPayAllForRangeSubmitter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPayAllForRangeSubmitter is a free data retrieval call binding the contract method 0x146cd61d.
//
// Solidity: function isPayAllForRangeSubmitter(address ) view returns(bool)
func (_PaymentCoordinator *PaymentCoordinatorSession) IsPayAllForRangeSubmitter(arg0 common.Address) (bool, error) {
	return _PaymentCoordinator.Contract.IsPayAllForRangeSubmitter(&_PaymentCoordinator.CallOpts, arg0)
}

// IsPayAllForRangeSubmitter is a free data retrieval call binding the contract method 0x146cd61d.
//
// Solidity: function isPayAllForRangeSubmitter(address ) view returns(bool)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) IsPayAllForRangeSubmitter(arg0 common.Address) (bool, error) {
	return _PaymentCoordinator.Contract.IsPayAllForRangeSubmitter(&_PaymentCoordinator.CallOpts, arg0)
}

// IsRangePaymentForAllHash is a free data retrieval call binding the contract method 0x73f2fbea.
//
// Solidity: function isRangePaymentForAllHash(address , bytes32 ) view returns(bool)
func (_PaymentCoordinator *PaymentCoordinatorCaller) IsRangePaymentForAllHash(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "isRangePaymentForAllHash", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRangePaymentForAllHash is a free data retrieval call binding the contract method 0x73f2fbea.
//
// Solidity: function isRangePaymentForAllHash(address , bytes32 ) view returns(bool)
func (_PaymentCoordinator *PaymentCoordinatorSession) IsRangePaymentForAllHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _PaymentCoordinator.Contract.IsRangePaymentForAllHash(&_PaymentCoordinator.CallOpts, arg0, arg1)
}

// IsRangePaymentForAllHash is a free data retrieval call binding the contract method 0x73f2fbea.
//
// Solidity: function isRangePaymentForAllHash(address , bytes32 ) view returns(bool)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) IsRangePaymentForAllHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _PaymentCoordinator.Contract.IsRangePaymentForAllHash(&_PaymentCoordinator.CallOpts, arg0, arg1)
}

// IsRangePaymentHash is a free data retrieval call binding the contract method 0xc8371b46.
//
// Solidity: function isRangePaymentHash(address , bytes32 ) view returns(bool)
func (_PaymentCoordinator *PaymentCoordinatorCaller) IsRangePaymentHash(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "isRangePaymentHash", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRangePaymentHash is a free data retrieval call binding the contract method 0xc8371b46.
//
// Solidity: function isRangePaymentHash(address , bytes32 ) view returns(bool)
func (_PaymentCoordinator *PaymentCoordinatorSession) IsRangePaymentHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _PaymentCoordinator.Contract.IsRangePaymentHash(&_PaymentCoordinator.CallOpts, arg0, arg1)
}

// IsRangePaymentHash is a free data retrieval call binding the contract method 0xc8371b46.
//
// Solidity: function isRangePaymentHash(address , bytes32 ) view returns(bool)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) IsRangePaymentHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _PaymentCoordinator.Contract.IsRangePaymentHash(&_PaymentCoordinator.CallOpts, arg0, arg1)
}

// OperatorCommissionBips is a free data retrieval call binding the contract method 0x22f19a64.
//
// Solidity: function operatorCommissionBips(address operator, address avs) view returns(uint16)
func (_PaymentCoordinator *PaymentCoordinatorCaller) OperatorCommissionBips(opts *bind.CallOpts, operator common.Address, avs common.Address) (uint16, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "operatorCommissionBips", operator, avs)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// OperatorCommissionBips is a free data retrieval call binding the contract method 0x22f19a64.
//
// Solidity: function operatorCommissionBips(address operator, address avs) view returns(uint16)
func (_PaymentCoordinator *PaymentCoordinatorSession) OperatorCommissionBips(operator common.Address, avs common.Address) (uint16, error) {
	return _PaymentCoordinator.Contract.OperatorCommissionBips(&_PaymentCoordinator.CallOpts, operator, avs)
}

// OperatorCommissionBips is a free data retrieval call binding the contract method 0x22f19a64.
//
// Solidity: function operatorCommissionBips(address operator, address avs) view returns(uint16)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) OperatorCommissionBips(operator common.Address, avs common.Address) (uint16, error) {
	return _PaymentCoordinator.Contract.OperatorCommissionBips(&_PaymentCoordinator.CallOpts, operator, avs)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorSession) Owner() (common.Address, error) {
	return _PaymentCoordinator.Contract.Owner(&_PaymentCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) Owner() (common.Address, error) {
	return _PaymentCoordinator.Contract.Owner(&_PaymentCoordinator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_PaymentCoordinator *PaymentCoordinatorCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_PaymentCoordinator *PaymentCoordinatorSession) Paused(index uint8) (bool, error) {
	return _PaymentCoordinator.Contract.Paused(&_PaymentCoordinator.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) Paused(index uint8) (bool, error) {
	return _PaymentCoordinator.Contract.Paused(&_PaymentCoordinator.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_PaymentCoordinator *PaymentCoordinatorCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_PaymentCoordinator *PaymentCoordinatorSession) Paused0() (*big.Int, error) {
	return _PaymentCoordinator.Contract.Paused0(&_PaymentCoordinator.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) Paused0() (*big.Int, error) {
	return _PaymentCoordinator.Contract.Paused0(&_PaymentCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorSession) PauserRegistry() (common.Address, error) {
	return _PaymentCoordinator.Contract.PauserRegistry(&_PaymentCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) PauserRegistry() (common.Address, error) {
	return _PaymentCoordinator.Contract.PauserRegistry(&_PaymentCoordinator.CallOpts)
}

// PaymentNonce is a free data retrieval call binding the contract method 0xd588cefa.
//
// Solidity: function paymentNonce(address ) view returns(uint256)
func (_PaymentCoordinator *PaymentCoordinatorCaller) PaymentNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "paymentNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PaymentNonce is a free data retrieval call binding the contract method 0xd588cefa.
//
// Solidity: function paymentNonce(address ) view returns(uint256)
func (_PaymentCoordinator *PaymentCoordinatorSession) PaymentNonce(arg0 common.Address) (*big.Int, error) {
	return _PaymentCoordinator.Contract.PaymentNonce(&_PaymentCoordinator.CallOpts, arg0)
}

// PaymentNonce is a free data retrieval call binding the contract method 0xd588cefa.
//
// Solidity: function paymentNonce(address ) view returns(uint256)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) PaymentNonce(arg0 common.Address) (*big.Int, error) {
	return _PaymentCoordinator.Contract.PaymentNonce(&_PaymentCoordinator.CallOpts, arg0)
}

// PaymentUpdater is a free data retrieval call binding the contract method 0x66d3b16b.
//
// Solidity: function paymentUpdater() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorCaller) PaymentUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "paymentUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentUpdater is a free data retrieval call binding the contract method 0x66d3b16b.
//
// Solidity: function paymentUpdater() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorSession) PaymentUpdater() (common.Address, error) {
	return _PaymentCoordinator.Contract.PaymentUpdater(&_PaymentCoordinator.CallOpts)
}

// PaymentUpdater is a free data retrieval call binding the contract method 0x66d3b16b.
//
// Solidity: function paymentUpdater() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) PaymentUpdater() (common.Address, error) {
	return _PaymentCoordinator.Contract.PaymentUpdater(&_PaymentCoordinator.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentCoordinator.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorSession) StrategyManager() (common.Address, error) {
	return _PaymentCoordinator.Contract.StrategyManager(&_PaymentCoordinator.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_PaymentCoordinator *PaymentCoordinatorCallerSession) StrategyManager() (common.Address, error) {
	return _PaymentCoordinator.Contract.StrategyManager(&_PaymentCoordinator.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xd4540a55.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus, address _paymentUpdater, uint32 _activationDelay, uint16 _globalCommissionBips) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int, _paymentUpdater common.Address, _activationDelay uint32, _globalCommissionBips uint16) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "initialize", initialOwner, _pauserRegistry, initialPausedStatus, _paymentUpdater, _activationDelay, _globalCommissionBips)
}

// Initialize is a paid mutator transaction binding the contract method 0xd4540a55.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus, address _paymentUpdater, uint32 _activationDelay, uint16 _globalCommissionBips) returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int, _paymentUpdater common.Address, _activationDelay uint32, _globalCommissionBips uint16) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.Initialize(&_PaymentCoordinator.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus, _paymentUpdater, _activationDelay, _globalCommissionBips)
}

// Initialize is a paid mutator transaction binding the contract method 0xd4540a55.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus, address _paymentUpdater, uint32 _activationDelay, uint16 _globalCommissionBips) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int, _paymentUpdater common.Address, _activationDelay uint32, _globalCommissionBips uint16) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.Initialize(&_PaymentCoordinator.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus, _paymentUpdater, _activationDelay, _globalCommissionBips)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.Pause(&_PaymentCoordinator.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.Pause(&_PaymentCoordinator.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) PauseAll() (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.PauseAll(&_PaymentCoordinator.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) PauseAll() (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.PauseAll(&_PaymentCoordinator.TransactOpts)
}

// PayAllForRange is a paid mutator transaction binding the contract method 0xb5715284.
//
// Solidity: function payAllForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) PayAllForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "payAllForRange", rangePayments)
}

// PayAllForRange is a paid mutator transaction binding the contract method 0xb5715284.
//
// Solidity: function payAllForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) PayAllForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.PayAllForRange(&_PaymentCoordinator.TransactOpts, rangePayments)
}

// PayAllForRange is a paid mutator transaction binding the contract method 0xb5715284.
//
// Solidity: function payAllForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) PayAllForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.PayAllForRange(&_PaymentCoordinator.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.PayForRange(&_PaymentCoordinator.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.PayForRange(&_PaymentCoordinator.TransactOpts, rangePayments)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) ProcessClaim(opts *bind.TransactOpts, claim IPaymentCoordinatorPaymentMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "processClaim", claim, recipient)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) ProcessClaim(claim IPaymentCoordinatorPaymentMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.ProcessClaim(&_PaymentCoordinator.TransactOpts, claim, recipient)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) ProcessClaim(claim IPaymentCoordinatorPaymentMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.ProcessClaim(&_PaymentCoordinator.TransactOpts, claim, recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.RenounceOwnership(&_PaymentCoordinator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.RenounceOwnership(&_PaymentCoordinator.TransactOpts)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) SetActivationDelay(opts *bind.TransactOpts, _activationDelay uint32) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "setActivationDelay", _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.SetActivationDelay(&_PaymentCoordinator.TransactOpts, _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.SetActivationDelay(&_PaymentCoordinator.TransactOpts, _activationDelay)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) SetClaimerFor(opts *bind.TransactOpts, claimer common.Address) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "setClaimerFor", claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.SetClaimerFor(&_PaymentCoordinator.TransactOpts, claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.SetClaimerFor(&_PaymentCoordinator.TransactOpts, claimer)
}

// SetGlobalOperatorCommission is a paid mutator transaction binding the contract method 0xe221b245.
//
// Solidity: function setGlobalOperatorCommission(uint16 _globalCommissionBips) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) SetGlobalOperatorCommission(opts *bind.TransactOpts, _globalCommissionBips uint16) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "setGlobalOperatorCommission", _globalCommissionBips)
}

// SetGlobalOperatorCommission is a paid mutator transaction binding the contract method 0xe221b245.
//
// Solidity: function setGlobalOperatorCommission(uint16 _globalCommissionBips) returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) SetGlobalOperatorCommission(_globalCommissionBips uint16) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.SetGlobalOperatorCommission(&_PaymentCoordinator.TransactOpts, _globalCommissionBips)
}

// SetGlobalOperatorCommission is a paid mutator transaction binding the contract method 0xe221b245.
//
// Solidity: function setGlobalOperatorCommission(uint16 _globalCommissionBips) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) SetGlobalOperatorCommission(_globalCommissionBips uint16) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.SetGlobalOperatorCommission(&_PaymentCoordinator.TransactOpts, _globalCommissionBips)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.SetPauserRegistry(&_PaymentCoordinator.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.SetPauserRegistry(&_PaymentCoordinator.TransactOpts, newPauserRegistry)
}

// SetPayAllForRangeSubmitter is a paid mutator transaction binding the contract method 0xec1680de.
//
// Solidity: function setPayAllForRangeSubmitter(address _submitter, bool _newValue) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) SetPayAllForRangeSubmitter(opts *bind.TransactOpts, _submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "setPayAllForRangeSubmitter", _submitter, _newValue)
}

// SetPayAllForRangeSubmitter is a paid mutator transaction binding the contract method 0xec1680de.
//
// Solidity: function setPayAllForRangeSubmitter(address _submitter, bool _newValue) returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) SetPayAllForRangeSubmitter(_submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.SetPayAllForRangeSubmitter(&_PaymentCoordinator.TransactOpts, _submitter, _newValue)
}

// SetPayAllForRangeSubmitter is a paid mutator transaction binding the contract method 0xec1680de.
//
// Solidity: function setPayAllForRangeSubmitter(address _submitter, bool _newValue) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) SetPayAllForRangeSubmitter(_submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.SetPayAllForRangeSubmitter(&_PaymentCoordinator.TransactOpts, _submitter, _newValue)
}

// SetPaymentUpdater is a paid mutator transaction binding the contract method 0x18190f53.
//
// Solidity: function setPaymentUpdater(address _paymentUpdater) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) SetPaymentUpdater(opts *bind.TransactOpts, _paymentUpdater common.Address) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "setPaymentUpdater", _paymentUpdater)
}

// SetPaymentUpdater is a paid mutator transaction binding the contract method 0x18190f53.
//
// Solidity: function setPaymentUpdater(address _paymentUpdater) returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) SetPaymentUpdater(_paymentUpdater common.Address) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.SetPaymentUpdater(&_PaymentCoordinator.TransactOpts, _paymentUpdater)
}

// SetPaymentUpdater is a paid mutator transaction binding the contract method 0x18190f53.
//
// Solidity: function setPaymentUpdater(address _paymentUpdater) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) SetPaymentUpdater(_paymentUpdater common.Address) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.SetPaymentUpdater(&_PaymentCoordinator.TransactOpts, _paymentUpdater)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 paymentCalculationEndTimestamp) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) SubmitRoot(opts *bind.TransactOpts, root [32]byte, paymentCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "submitRoot", root, paymentCalculationEndTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 paymentCalculationEndTimestamp) returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) SubmitRoot(root [32]byte, paymentCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.SubmitRoot(&_PaymentCoordinator.TransactOpts, root, paymentCalculationEndTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 paymentCalculationEndTimestamp) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) SubmitRoot(root [32]byte, paymentCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.SubmitRoot(&_PaymentCoordinator.TransactOpts, root, paymentCalculationEndTimestamp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.TransferOwnership(&_PaymentCoordinator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.TransferOwnership(&_PaymentCoordinator.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PaymentCoordinator.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_PaymentCoordinator *PaymentCoordinatorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.Unpause(&_PaymentCoordinator.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_PaymentCoordinator *PaymentCoordinatorTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _PaymentCoordinator.Contract.Unpause(&_PaymentCoordinator.TransactOpts, newPausedStatus)
}

// PaymentCoordinatorActivationDelaySetIterator is returned from FilterActivationDelaySet and is used to iterate over the raw logs and unpacked data for ActivationDelaySet events raised by the PaymentCoordinator contract.
type PaymentCoordinatorActivationDelaySetIterator struct {
	Event *PaymentCoordinatorActivationDelaySet // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorActivationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorActivationDelaySet)
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
		it.Event = new(PaymentCoordinatorActivationDelaySet)
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
func (it *PaymentCoordinatorActivationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorActivationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorActivationDelaySet represents a ActivationDelaySet event raised by the PaymentCoordinator contract.
type PaymentCoordinatorActivationDelaySet struct {
	OldActivationDelay uint32
	NewActivationDelay uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterActivationDelaySet is a free log retrieval operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) FilterActivationDelaySet(opts *bind.FilterOpts) (*PaymentCoordinatorActivationDelaySetIterator, error) {

	logs, sub, err := _PaymentCoordinator.contract.FilterLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorActivationDelaySetIterator{contract: _PaymentCoordinator.contract, event: "ActivationDelaySet", logs: logs, sub: sub}, nil
}

// WatchActivationDelaySet is a free log subscription operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) WatchActivationDelaySet(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorActivationDelaySet) (event.Subscription, error) {

	logs, sub, err := _PaymentCoordinator.contract.WatchLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorActivationDelaySet)
				if err := _PaymentCoordinator.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
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

// ParseActivationDelaySet is a log parse operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) ParseActivationDelaySet(log types.Log) (*PaymentCoordinatorActivationDelaySet, error) {
	event := new(PaymentCoordinatorActivationDelaySet)
	if err := _PaymentCoordinator.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorClaimerForSetIterator is returned from FilterClaimerForSet and is used to iterate over the raw logs and unpacked data for ClaimerForSet events raised by the PaymentCoordinator contract.
type PaymentCoordinatorClaimerForSetIterator struct {
	Event *PaymentCoordinatorClaimerForSet // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorClaimerForSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorClaimerForSet)
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
		it.Event = new(PaymentCoordinatorClaimerForSet)
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
func (it *PaymentCoordinatorClaimerForSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorClaimerForSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorClaimerForSet represents a ClaimerForSet event raised by the PaymentCoordinator contract.
type PaymentCoordinatorClaimerForSet struct {
	Earner     common.Address
	OldClaimer common.Address
	Claimer    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimerForSet is a free log retrieval operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) FilterClaimerForSet(opts *bind.FilterOpts, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (*PaymentCoordinatorClaimerForSetIterator, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var oldClaimerRule []interface{}
	for _, oldClaimerItem := range oldClaimer {
		oldClaimerRule = append(oldClaimerRule, oldClaimerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.FilterLogs(opts, "ClaimerForSet", earnerRule, oldClaimerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorClaimerForSetIterator{contract: _PaymentCoordinator.contract, event: "ClaimerForSet", logs: logs, sub: sub}, nil
}

// WatchClaimerForSet is a free log subscription operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) WatchClaimerForSet(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorClaimerForSet, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (event.Subscription, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var oldClaimerRule []interface{}
	for _, oldClaimerItem := range oldClaimer {
		oldClaimerRule = append(oldClaimerRule, oldClaimerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.WatchLogs(opts, "ClaimerForSet", earnerRule, oldClaimerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorClaimerForSet)
				if err := _PaymentCoordinator.contract.UnpackLog(event, "ClaimerForSet", log); err != nil {
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

// ParseClaimerForSet is a log parse operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) ParseClaimerForSet(log types.Log) (*PaymentCoordinatorClaimerForSet, error) {
	event := new(PaymentCoordinatorClaimerForSet)
	if err := _PaymentCoordinator.contract.UnpackLog(event, "ClaimerForSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorDistributionRootSubmittedIterator is returned from FilterDistributionRootSubmitted and is used to iterate over the raw logs and unpacked data for DistributionRootSubmitted events raised by the PaymentCoordinator contract.
type PaymentCoordinatorDistributionRootSubmittedIterator struct {
	Event *PaymentCoordinatorDistributionRootSubmitted // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorDistributionRootSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorDistributionRootSubmitted)
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
		it.Event = new(PaymentCoordinatorDistributionRootSubmitted)
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
func (it *PaymentCoordinatorDistributionRootSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorDistributionRootSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorDistributionRootSubmitted represents a DistributionRootSubmitted event raised by the PaymentCoordinator contract.
type PaymentCoordinatorDistributionRootSubmitted struct {
	RootIndex                      uint32
	Root                           [32]byte
	PaymentCalculationEndTimestamp uint32
	ActivatedAt                    uint32
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterDistributionRootSubmitted is a free log retrieval operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed paymentCalculationEndTimestamp, uint32 activatedAt)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) FilterDistributionRootSubmitted(opts *bind.FilterOpts, rootIndex []uint32, root [][32]byte, paymentCalculationEndTimestamp []uint32) (*PaymentCoordinatorDistributionRootSubmittedIterator, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}
	var paymentCalculationEndTimestampRule []interface{}
	for _, paymentCalculationEndTimestampItem := range paymentCalculationEndTimestamp {
		paymentCalculationEndTimestampRule = append(paymentCalculationEndTimestampRule, paymentCalculationEndTimestampItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.FilterLogs(opts, "DistributionRootSubmitted", rootIndexRule, rootRule, paymentCalculationEndTimestampRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorDistributionRootSubmittedIterator{contract: _PaymentCoordinator.contract, event: "DistributionRootSubmitted", logs: logs, sub: sub}, nil
}

// WatchDistributionRootSubmitted is a free log subscription operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed paymentCalculationEndTimestamp, uint32 activatedAt)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) WatchDistributionRootSubmitted(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorDistributionRootSubmitted, rootIndex []uint32, root [][32]byte, paymentCalculationEndTimestamp []uint32) (event.Subscription, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}
	var paymentCalculationEndTimestampRule []interface{}
	for _, paymentCalculationEndTimestampItem := range paymentCalculationEndTimestamp {
		paymentCalculationEndTimestampRule = append(paymentCalculationEndTimestampRule, paymentCalculationEndTimestampItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.WatchLogs(opts, "DistributionRootSubmitted", rootIndexRule, rootRule, paymentCalculationEndTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorDistributionRootSubmitted)
				if err := _PaymentCoordinator.contract.UnpackLog(event, "DistributionRootSubmitted", log); err != nil {
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

// ParseDistributionRootSubmitted is a log parse operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed paymentCalculationEndTimestamp, uint32 activatedAt)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) ParseDistributionRootSubmitted(log types.Log) (*PaymentCoordinatorDistributionRootSubmitted, error) {
	event := new(PaymentCoordinatorDistributionRootSubmitted)
	if err := _PaymentCoordinator.contract.UnpackLog(event, "DistributionRootSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorGlobalCommissionBipsSetIterator is returned from FilterGlobalCommissionBipsSet and is used to iterate over the raw logs and unpacked data for GlobalCommissionBipsSet events raised by the PaymentCoordinator contract.
type PaymentCoordinatorGlobalCommissionBipsSetIterator struct {
	Event *PaymentCoordinatorGlobalCommissionBipsSet // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorGlobalCommissionBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorGlobalCommissionBipsSet)
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
		it.Event = new(PaymentCoordinatorGlobalCommissionBipsSet)
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
func (it *PaymentCoordinatorGlobalCommissionBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorGlobalCommissionBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorGlobalCommissionBipsSet represents a GlobalCommissionBipsSet event raised by the PaymentCoordinator contract.
type PaymentCoordinatorGlobalCommissionBipsSet struct {
	OldGlobalCommissionBips uint16
	NewGlobalCommissionBips uint16
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterGlobalCommissionBipsSet is a free log retrieval operation binding the contract event 0x8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06.
//
// Solidity: event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) FilterGlobalCommissionBipsSet(opts *bind.FilterOpts) (*PaymentCoordinatorGlobalCommissionBipsSetIterator, error) {

	logs, sub, err := _PaymentCoordinator.contract.FilterLogs(opts, "GlobalCommissionBipsSet")
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorGlobalCommissionBipsSetIterator{contract: _PaymentCoordinator.contract, event: "GlobalCommissionBipsSet", logs: logs, sub: sub}, nil
}

// WatchGlobalCommissionBipsSet is a free log subscription operation binding the contract event 0x8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06.
//
// Solidity: event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) WatchGlobalCommissionBipsSet(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorGlobalCommissionBipsSet) (event.Subscription, error) {

	logs, sub, err := _PaymentCoordinator.contract.WatchLogs(opts, "GlobalCommissionBipsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorGlobalCommissionBipsSet)
				if err := _PaymentCoordinator.contract.UnpackLog(event, "GlobalCommissionBipsSet", log); err != nil {
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

// ParseGlobalCommissionBipsSet is a log parse operation binding the contract event 0x8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06.
//
// Solidity: event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) ParseGlobalCommissionBipsSet(log types.Log) (*PaymentCoordinatorGlobalCommissionBipsSet, error) {
	event := new(PaymentCoordinatorGlobalCommissionBipsSet)
	if err := _PaymentCoordinator.contract.UnpackLog(event, "GlobalCommissionBipsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PaymentCoordinator contract.
type PaymentCoordinatorInitializedIterator struct {
	Event *PaymentCoordinatorInitialized // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorInitialized)
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
		it.Event = new(PaymentCoordinatorInitialized)
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
func (it *PaymentCoordinatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorInitialized represents a Initialized event raised by the PaymentCoordinator contract.
type PaymentCoordinatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*PaymentCoordinatorInitializedIterator, error) {

	logs, sub, err := _PaymentCoordinator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorInitializedIterator{contract: _PaymentCoordinator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorInitialized) (event.Subscription, error) {

	logs, sub, err := _PaymentCoordinator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorInitialized)
				if err := _PaymentCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PaymentCoordinator *PaymentCoordinatorFilterer) ParseInitialized(log types.Log) (*PaymentCoordinatorInitialized, error) {
	event := new(PaymentCoordinatorInitialized)
	if err := _PaymentCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PaymentCoordinator contract.
type PaymentCoordinatorOwnershipTransferredIterator struct {
	Event *PaymentCoordinatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorOwnershipTransferred)
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
		it.Event = new(PaymentCoordinatorOwnershipTransferred)
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
func (it *PaymentCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorOwnershipTransferred represents a OwnershipTransferred event raised by the PaymentCoordinator contract.
type PaymentCoordinatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaymentCoordinatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorOwnershipTransferredIterator{contract: _PaymentCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorOwnershipTransferred)
				if err := _PaymentCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PaymentCoordinator *PaymentCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*PaymentCoordinatorOwnershipTransferred, error) {
	event := new(PaymentCoordinatorOwnershipTransferred)
	if err := _PaymentCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PaymentCoordinator contract.
type PaymentCoordinatorPausedIterator struct {
	Event *PaymentCoordinatorPaused // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorPaused)
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
		it.Event = new(PaymentCoordinatorPaused)
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
func (it *PaymentCoordinatorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorPaused represents a Paused event raised by the PaymentCoordinator contract.
type PaymentCoordinatorPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*PaymentCoordinatorPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorPausedIterator{contract: _PaymentCoordinator.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorPaused)
				if err := _PaymentCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PaymentCoordinator *PaymentCoordinatorFilterer) ParsePaused(log types.Log) (*PaymentCoordinatorPaused, error) {
	event := new(PaymentCoordinatorPaused)
	if err := _PaymentCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the PaymentCoordinator contract.
type PaymentCoordinatorPauserRegistrySetIterator struct {
	Event *PaymentCoordinatorPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorPauserRegistrySet)
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
		it.Event = new(PaymentCoordinatorPauserRegistrySet)
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
func (it *PaymentCoordinatorPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorPauserRegistrySet represents a PauserRegistrySet event raised by the PaymentCoordinator contract.
type PaymentCoordinatorPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*PaymentCoordinatorPauserRegistrySetIterator, error) {

	logs, sub, err := _PaymentCoordinator.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorPauserRegistrySetIterator{contract: _PaymentCoordinator.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _PaymentCoordinator.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorPauserRegistrySet)
				if err := _PaymentCoordinator.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_PaymentCoordinator *PaymentCoordinatorFilterer) ParsePauserRegistrySet(log types.Log) (*PaymentCoordinatorPauserRegistrySet, error) {
	event := new(PaymentCoordinatorPauserRegistrySet)
	if err := _PaymentCoordinator.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorPayAllForRangeSubmitterSetIterator is returned from FilterPayAllForRangeSubmitterSet and is used to iterate over the raw logs and unpacked data for PayAllForRangeSubmitterSet events raised by the PaymentCoordinator contract.
type PaymentCoordinatorPayAllForRangeSubmitterSetIterator struct {
	Event *PaymentCoordinatorPayAllForRangeSubmitterSet // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorPayAllForRangeSubmitterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorPayAllForRangeSubmitterSet)
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
		it.Event = new(PaymentCoordinatorPayAllForRangeSubmitterSet)
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
func (it *PaymentCoordinatorPayAllForRangeSubmitterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorPayAllForRangeSubmitterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorPayAllForRangeSubmitterSet represents a PayAllForRangeSubmitterSet event raised by the PaymentCoordinator contract.
type PaymentCoordinatorPayAllForRangeSubmitterSet struct {
	PayAllForRangeSubmitter common.Address
	OldValue                bool
	NewValue                bool
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterPayAllForRangeSubmitterSet is a free log retrieval operation binding the contract event 0x17b06be02af2803593116eb96121b9c6e8bee1cc1b145e7c31c19c180e86189b.
//
// Solidity: event PayAllForRangeSubmitterSet(address indexed payAllForRangeSubmitter, bool indexed oldValue, bool indexed newValue)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) FilterPayAllForRangeSubmitterSet(opts *bind.FilterOpts, payAllForRangeSubmitter []common.Address, oldValue []bool, newValue []bool) (*PaymentCoordinatorPayAllForRangeSubmitterSetIterator, error) {

	var payAllForRangeSubmitterRule []interface{}
	for _, payAllForRangeSubmitterItem := range payAllForRangeSubmitter {
		payAllForRangeSubmitterRule = append(payAllForRangeSubmitterRule, payAllForRangeSubmitterItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}
	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.FilterLogs(opts, "PayAllForRangeSubmitterSet", payAllForRangeSubmitterRule, oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorPayAllForRangeSubmitterSetIterator{contract: _PaymentCoordinator.contract, event: "PayAllForRangeSubmitterSet", logs: logs, sub: sub}, nil
}

// WatchPayAllForRangeSubmitterSet is a free log subscription operation binding the contract event 0x17b06be02af2803593116eb96121b9c6e8bee1cc1b145e7c31c19c180e86189b.
//
// Solidity: event PayAllForRangeSubmitterSet(address indexed payAllForRangeSubmitter, bool indexed oldValue, bool indexed newValue)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) WatchPayAllForRangeSubmitterSet(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorPayAllForRangeSubmitterSet, payAllForRangeSubmitter []common.Address, oldValue []bool, newValue []bool) (event.Subscription, error) {

	var payAllForRangeSubmitterRule []interface{}
	for _, payAllForRangeSubmitterItem := range payAllForRangeSubmitter {
		payAllForRangeSubmitterRule = append(payAllForRangeSubmitterRule, payAllForRangeSubmitterItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}
	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.WatchLogs(opts, "PayAllForRangeSubmitterSet", payAllForRangeSubmitterRule, oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorPayAllForRangeSubmitterSet)
				if err := _PaymentCoordinator.contract.UnpackLog(event, "PayAllForRangeSubmitterSet", log); err != nil {
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

// ParsePayAllForRangeSubmitterSet is a log parse operation binding the contract event 0x17b06be02af2803593116eb96121b9c6e8bee1cc1b145e7c31c19c180e86189b.
//
// Solidity: event PayAllForRangeSubmitterSet(address indexed payAllForRangeSubmitter, bool indexed oldValue, bool indexed newValue)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) ParsePayAllForRangeSubmitterSet(log types.Log) (*PaymentCoordinatorPayAllForRangeSubmitterSet, error) {
	event := new(PaymentCoordinatorPayAllForRangeSubmitterSet)
	if err := _PaymentCoordinator.contract.UnpackLog(event, "PayAllForRangeSubmitterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorPaymentClaimedIterator is returned from FilterPaymentClaimed and is used to iterate over the raw logs and unpacked data for PaymentClaimed events raised by the PaymentCoordinator contract.
type PaymentCoordinatorPaymentClaimedIterator struct {
	Event *PaymentCoordinatorPaymentClaimed // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorPaymentClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorPaymentClaimed)
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
		it.Event = new(PaymentCoordinatorPaymentClaimed)
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
func (it *PaymentCoordinatorPaymentClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorPaymentClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorPaymentClaimed represents a PaymentClaimed event raised by the PaymentCoordinator contract.
type PaymentCoordinatorPaymentClaimed struct {
	Root          [32]byte
	Earner        common.Address
	Claimer       common.Address
	Recipient     common.Address
	Token         common.Address
	ClaimedAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaymentClaimed is a free log retrieval operation binding the contract event 0xbff1e5a32b3f6d3b3c0a7e675ead2091fea820852f35a77abdd6d2420bec4778.
//
// Solidity: event PaymentClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) FilterPaymentClaimed(opts *bind.FilterOpts, earner []common.Address, claimer []common.Address, recipient []common.Address) (*PaymentCoordinatorPaymentClaimedIterator, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.FilterLogs(opts, "PaymentClaimed", earnerRule, claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorPaymentClaimedIterator{contract: _PaymentCoordinator.contract, event: "PaymentClaimed", logs: logs, sub: sub}, nil
}

// WatchPaymentClaimed is a free log subscription operation binding the contract event 0xbff1e5a32b3f6d3b3c0a7e675ead2091fea820852f35a77abdd6d2420bec4778.
//
// Solidity: event PaymentClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) WatchPaymentClaimed(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorPaymentClaimed, earner []common.Address, claimer []common.Address, recipient []common.Address) (event.Subscription, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.WatchLogs(opts, "PaymentClaimed", earnerRule, claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorPaymentClaimed)
				if err := _PaymentCoordinator.contract.UnpackLog(event, "PaymentClaimed", log); err != nil {
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

// ParsePaymentClaimed is a log parse operation binding the contract event 0xbff1e5a32b3f6d3b3c0a7e675ead2091fea820852f35a77abdd6d2420bec4778.
//
// Solidity: event PaymentClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) ParsePaymentClaimed(log types.Log) (*PaymentCoordinatorPaymentClaimed, error) {
	event := new(PaymentCoordinatorPaymentClaimed)
	if err := _PaymentCoordinator.contract.UnpackLog(event, "PaymentClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorPaymentUpdaterSetIterator is returned from FilterPaymentUpdaterSet and is used to iterate over the raw logs and unpacked data for PaymentUpdaterSet events raised by the PaymentCoordinator contract.
type PaymentCoordinatorPaymentUpdaterSetIterator struct {
	Event *PaymentCoordinatorPaymentUpdaterSet // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorPaymentUpdaterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorPaymentUpdaterSet)
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
		it.Event = new(PaymentCoordinatorPaymentUpdaterSet)
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
func (it *PaymentCoordinatorPaymentUpdaterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorPaymentUpdaterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorPaymentUpdaterSet represents a PaymentUpdaterSet event raised by the PaymentCoordinator contract.
type PaymentCoordinatorPaymentUpdaterSet struct {
	OldPaymentUpdater common.Address
	NewPaymentUpdater common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPaymentUpdaterSet is a free log retrieval operation binding the contract event 0x07d2890b3eb1206e7c3cb6bf8d46da31385ace3ce99abf85e5b690c83aa49678.
//
// Solidity: event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) FilterPaymentUpdaterSet(opts *bind.FilterOpts, oldPaymentUpdater []common.Address, newPaymentUpdater []common.Address) (*PaymentCoordinatorPaymentUpdaterSetIterator, error) {

	var oldPaymentUpdaterRule []interface{}
	for _, oldPaymentUpdaterItem := range oldPaymentUpdater {
		oldPaymentUpdaterRule = append(oldPaymentUpdaterRule, oldPaymentUpdaterItem)
	}
	var newPaymentUpdaterRule []interface{}
	for _, newPaymentUpdaterItem := range newPaymentUpdater {
		newPaymentUpdaterRule = append(newPaymentUpdaterRule, newPaymentUpdaterItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.FilterLogs(opts, "PaymentUpdaterSet", oldPaymentUpdaterRule, newPaymentUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorPaymentUpdaterSetIterator{contract: _PaymentCoordinator.contract, event: "PaymentUpdaterSet", logs: logs, sub: sub}, nil
}

// WatchPaymentUpdaterSet is a free log subscription operation binding the contract event 0x07d2890b3eb1206e7c3cb6bf8d46da31385ace3ce99abf85e5b690c83aa49678.
//
// Solidity: event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) WatchPaymentUpdaterSet(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorPaymentUpdaterSet, oldPaymentUpdater []common.Address, newPaymentUpdater []common.Address) (event.Subscription, error) {

	var oldPaymentUpdaterRule []interface{}
	for _, oldPaymentUpdaterItem := range oldPaymentUpdater {
		oldPaymentUpdaterRule = append(oldPaymentUpdaterRule, oldPaymentUpdaterItem)
	}
	var newPaymentUpdaterRule []interface{}
	for _, newPaymentUpdaterItem := range newPaymentUpdater {
		newPaymentUpdaterRule = append(newPaymentUpdaterRule, newPaymentUpdaterItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.WatchLogs(opts, "PaymentUpdaterSet", oldPaymentUpdaterRule, newPaymentUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorPaymentUpdaterSet)
				if err := _PaymentCoordinator.contract.UnpackLog(event, "PaymentUpdaterSet", log); err != nil {
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

// ParsePaymentUpdaterSet is a log parse operation binding the contract event 0x07d2890b3eb1206e7c3cb6bf8d46da31385ace3ce99abf85e5b690c83aa49678.
//
// Solidity: event PaymentUpdaterSet(address indexed oldPaymentUpdater, address indexed newPaymentUpdater)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) ParsePaymentUpdaterSet(log types.Log) (*PaymentCoordinatorPaymentUpdaterSet, error) {
	event := new(PaymentCoordinatorPaymentUpdaterSet)
	if err := _PaymentCoordinator.contract.UnpackLog(event, "PaymentUpdaterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorRangePaymentCreatedIterator is returned from FilterRangePaymentCreated and is used to iterate over the raw logs and unpacked data for RangePaymentCreated events raised by the PaymentCoordinator contract.
type PaymentCoordinatorRangePaymentCreatedIterator struct {
	Event *PaymentCoordinatorRangePaymentCreated // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorRangePaymentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorRangePaymentCreated)
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
		it.Event = new(PaymentCoordinatorRangePaymentCreated)
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
func (it *PaymentCoordinatorRangePaymentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorRangePaymentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorRangePaymentCreated represents a RangePaymentCreated event raised by the PaymentCoordinator contract.
type PaymentCoordinatorRangePaymentCreated struct {
	Avs              common.Address
	PaymentNonce     *big.Int
	RangePaymentHash [32]byte
	RangePayment     IPaymentCoordinatorRangePayment
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRangePaymentCreated is a free log retrieval operation binding the contract event 0x2a0659fa4c38e0075469a0e0dd737e045dc316ffd6cb6e68755c119ee0882aea.
//
// Solidity: event RangePaymentCreated(address indexed avs, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) FilterRangePaymentCreated(opts *bind.FilterOpts, avs []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (*PaymentCoordinatorRangePaymentCreatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var paymentNonceRule []interface{}
	for _, paymentNonceItem := range paymentNonce {
		paymentNonceRule = append(paymentNonceRule, paymentNonceItem)
	}
	var rangePaymentHashRule []interface{}
	for _, rangePaymentHashItem := range rangePaymentHash {
		rangePaymentHashRule = append(rangePaymentHashRule, rangePaymentHashItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.FilterLogs(opts, "RangePaymentCreated", avsRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorRangePaymentCreatedIterator{contract: _PaymentCoordinator.contract, event: "RangePaymentCreated", logs: logs, sub: sub}, nil
}

// WatchRangePaymentCreated is a free log subscription operation binding the contract event 0x2a0659fa4c38e0075469a0e0dd737e045dc316ffd6cb6e68755c119ee0882aea.
//
// Solidity: event RangePaymentCreated(address indexed avs, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) WatchRangePaymentCreated(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorRangePaymentCreated, avs []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var paymentNonceRule []interface{}
	for _, paymentNonceItem := range paymentNonce {
		paymentNonceRule = append(paymentNonceRule, paymentNonceItem)
	}
	var rangePaymentHashRule []interface{}
	for _, rangePaymentHashItem := range rangePaymentHash {
		rangePaymentHashRule = append(rangePaymentHashRule, rangePaymentHashItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.WatchLogs(opts, "RangePaymentCreated", avsRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorRangePaymentCreated)
				if err := _PaymentCoordinator.contract.UnpackLog(event, "RangePaymentCreated", log); err != nil {
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

// ParseRangePaymentCreated is a log parse operation binding the contract event 0x2a0659fa4c38e0075469a0e0dd737e045dc316ffd6cb6e68755c119ee0882aea.
//
// Solidity: event RangePaymentCreated(address indexed avs, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) ParseRangePaymentCreated(log types.Log) (*PaymentCoordinatorRangePaymentCreated, error) {
	event := new(PaymentCoordinatorRangePaymentCreated)
	if err := _PaymentCoordinator.contract.UnpackLog(event, "RangePaymentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorRangePaymentForAllCreatedIterator is returned from FilterRangePaymentForAllCreated and is used to iterate over the raw logs and unpacked data for RangePaymentForAllCreated events raised by the PaymentCoordinator contract.
type PaymentCoordinatorRangePaymentForAllCreatedIterator struct {
	Event *PaymentCoordinatorRangePaymentForAllCreated // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorRangePaymentForAllCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorRangePaymentForAllCreated)
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
		it.Event = new(PaymentCoordinatorRangePaymentForAllCreated)
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
func (it *PaymentCoordinatorRangePaymentForAllCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorRangePaymentForAllCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorRangePaymentForAllCreated represents a RangePaymentForAllCreated event raised by the PaymentCoordinator contract.
type PaymentCoordinatorRangePaymentForAllCreated struct {
	Submitter        common.Address
	PaymentNonce     *big.Int
	RangePaymentHash [32]byte
	RangePayment     IPaymentCoordinatorRangePayment
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRangePaymentForAllCreated is a free log retrieval operation binding the contract event 0xbc0782940ec5871f66c5490ef957f44d19a9adac1dac18b946ad0dd6579c30d6.
//
// Solidity: event RangePaymentForAllCreated(address indexed submitter, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) FilterRangePaymentForAllCreated(opts *bind.FilterOpts, submitter []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (*PaymentCoordinatorRangePaymentForAllCreatedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var paymentNonceRule []interface{}
	for _, paymentNonceItem := range paymentNonce {
		paymentNonceRule = append(paymentNonceRule, paymentNonceItem)
	}
	var rangePaymentHashRule []interface{}
	for _, rangePaymentHashItem := range rangePaymentHash {
		rangePaymentHashRule = append(rangePaymentHashRule, rangePaymentHashItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.FilterLogs(opts, "RangePaymentForAllCreated", submitterRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorRangePaymentForAllCreatedIterator{contract: _PaymentCoordinator.contract, event: "RangePaymentForAllCreated", logs: logs, sub: sub}, nil
}

// WatchRangePaymentForAllCreated is a free log subscription operation binding the contract event 0xbc0782940ec5871f66c5490ef957f44d19a9adac1dac18b946ad0dd6579c30d6.
//
// Solidity: event RangePaymentForAllCreated(address indexed submitter, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) WatchRangePaymentForAllCreated(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorRangePaymentForAllCreated, submitter []common.Address, paymentNonce []*big.Int, rangePaymentHash [][32]byte) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var paymentNonceRule []interface{}
	for _, paymentNonceItem := range paymentNonce {
		paymentNonceRule = append(paymentNonceRule, paymentNonceItem)
	}
	var rangePaymentHashRule []interface{}
	for _, rangePaymentHashItem := range rangePaymentHash {
		rangePaymentHashRule = append(rangePaymentHashRule, rangePaymentHashItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.WatchLogs(opts, "RangePaymentForAllCreated", submitterRule, paymentNonceRule, rangePaymentHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorRangePaymentForAllCreated)
				if err := _PaymentCoordinator.contract.UnpackLog(event, "RangePaymentForAllCreated", log); err != nil {
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

// ParseRangePaymentForAllCreated is a log parse operation binding the contract event 0xbc0782940ec5871f66c5490ef957f44d19a9adac1dac18b946ad0dd6579c30d6.
//
// Solidity: event RangePaymentForAllCreated(address indexed submitter, uint256 indexed paymentNonce, bytes32 indexed rangePaymentHash, ((address,uint96)[],address,uint256,uint32,uint32) rangePayment)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) ParseRangePaymentForAllCreated(log types.Log) (*PaymentCoordinatorRangePaymentForAllCreated, error) {
	event := new(PaymentCoordinatorRangePaymentForAllCreated)
	if err := _PaymentCoordinator.contract.UnpackLog(event, "RangePaymentForAllCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentCoordinatorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PaymentCoordinator contract.
type PaymentCoordinatorUnpausedIterator struct {
	Event *PaymentCoordinatorUnpaused // Event containing the contract specifics and raw log

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
func (it *PaymentCoordinatorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentCoordinatorUnpaused)
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
		it.Event = new(PaymentCoordinatorUnpaused)
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
func (it *PaymentCoordinatorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentCoordinatorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentCoordinatorUnpaused represents a Unpaused event raised by the PaymentCoordinator contract.
type PaymentCoordinatorUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*PaymentCoordinatorUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &PaymentCoordinatorUnpausedIterator{contract: _PaymentCoordinator.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_PaymentCoordinator *PaymentCoordinatorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PaymentCoordinatorUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PaymentCoordinator.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentCoordinatorUnpaused)
				if err := _PaymentCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PaymentCoordinator *PaymentCoordinatorFilterer) ParseUnpaused(log types.Log) (*PaymentCoordinatorUnpaused, error) {
	event := new(PaymentCoordinatorUnpaused)
	if err := _PaymentCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
