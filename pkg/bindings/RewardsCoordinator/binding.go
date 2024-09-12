// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RewardsCoordinator

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

// IRewardsCoordinatorDistributionRoot is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorDistributionRoot struct {
	Root                           [32]byte
	RewardsCalculationEndTimestamp uint32
	ActivatedAt                    uint32
	Disabled                       bool
}

// IRewardsCoordinatorEarnerTreeMerkleLeaf is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorEarnerTreeMerkleLeaf struct {
	Earner          common.Address
	EarnerTokenRoot [32]byte
}

// IRewardsCoordinatorRewardsMerkleClaim is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorRewardsMerkleClaim struct {
	RootIndex       uint32
	EarnerIndex     uint32
	EarnerTreeProof []byte
	EarnerLeaf      IRewardsCoordinatorEarnerTreeMerkleLeaf
	TokenIndices    []uint32
	TokenTreeProofs [][]byte
	TokenLeaves     []IRewardsCoordinatorTokenTreeMerkleLeaf
}

// IRewardsCoordinatorRewardsSubmission is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorRewardsSubmission struct {
	StrategiesAndMultipliers []IRewardsCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IRewardsCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// IRewardsCoordinatorTokenTreeMerkleLeaf is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTokenTreeMerkleLeaf struct {
	Token              common.Address
	CumulativeEarnings *big.Int
}

// RewardsCoordinatorMetaData contains all meta data concerning the RewardsCoordinator contract.
var RewardsCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_CALCULATION_INTERVAL_SECONDS\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_MAX_REWARDS_DURATION\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_MAX_RETROACTIVE_LENGTH\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_MAX_FUTURE_LENGTH\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"__GENESIS_REWARDS_TIMESTAMP\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CALCULATION_INTERVAL_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GENESIS_REWARDS_TIMESTAMP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FUTURE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_RETROACTIVE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_REWARDS_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activationDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateEarnerLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"calculateTokenLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.TokenTreeMerkleLeaf\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"checkClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.RewardsMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimerFor\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAVSRewardsSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRewardsForAllEarners\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRewardsForAllSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cumulativeClaimed\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currRewardsCalculationEndTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableRoot\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentClaimableDistributionRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentDistributionRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributionRootAtIndex\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributionRootsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRootIndexFromHash\",\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalOperatorCommissionBips\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_rewardsUpdater\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_globalCommissionBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAVSRewardsSubmissionHash\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRewardsForAllSubmitter\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRewardsSubmissionForAllEarnersHash\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRewardsSubmissionForAllHash\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorCommissionBips\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.RewardsMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinator.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardsUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setActivationDelay\",\"inputs\":[{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimerFor\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlobalOperatorCommission\",\"inputs\":[{\"name\":\"_globalCommissionBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsForAllSubmitter\",\"inputs\":[{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_newValue\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsUpdater\",\"inputs\":[{\"name\":\"_rewardsUpdater\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submissionNonce\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSRewardsSubmissionCreated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinator.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ActivationDelaySet\",\"inputs\":[{\"name\":\"oldActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimerForSet\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldClaimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootDisabled\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootSubmitted\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GlobalCommissionBipsSet\",\"inputs\":[{\"name\":\"oldGlobalCommissionBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newGlobalCommissionBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsClaimed\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"claimedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsForAllSubmitterSet\",\"inputs\":[{\"name\":\"rewardsForAllSubmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"newValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsSubmissionForAllCreated\",\"inputs\":[{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinator.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsSubmissionForAllEarnersCreated\",\"inputs\":[{\"name\":\"tokenHopper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinator.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsUpdaterSet\",\"inputs\":[{\"name\":\"oldRewardsUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newRewardsUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AmountExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AmountZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CommissionBipsExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DurationExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DurationNotMultipleOfCalculationIntervalSeconds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EarningsNotGreaterThanClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidClaimProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidClaimer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEarnerLeafIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRootIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidStrategy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTokenLeafIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewRootMustBeForNewCalculatedPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRewardsForAllSubmitter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRewardsUpdater\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofLengthNotMultipleOf32\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardsEndTimestampNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootActivated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootNotActivated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartTimestampNotMultipleOfCalculationIntervalSeconds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartTimestampTooFarInFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartTimestampTooFarInPast\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategiesNotInAscendingOrder\",\"inputs\":[]}]",
	Bin: "0x61018060405234801561001157600080fd5b506040516137c73803806137c7833981016040819052610030916102d0565b86868686868686610041858261035b565b63ffffffff16156100e55760405162461bcd60e51b815260206004820152606060248201527f52657761726473436f6f7264696e61746f723a2047454e455349535f5245574160448201527f5244535f54494d455354414d50206d7573742062652061206d756c7469706c6560648201527f206f662043414c43554c4154494f4e5f494e54455256414c5f5345434f4e4453608482015260a4015b60405180910390fd5b6100f2620151808661035b565b63ffffffff16156101915760405162461bcd60e51b815260206004820152605760248201527f52657761726473436f6f7264696e61746f723a2043414c43554c4154494f4e5f60448201527f494e54455256414c5f5345434f4e4453206d7573742062652061206d756c746960648201527f706c65206f6620534e415053484f545f434144454e4345000000000000000000608482015260a4016100dc565b6001600160a01b0396871661012052949095166101405263ffffffff92831660805290821660a052811660c05291821660e05216610100526101d16101e3565b50504661016052506103919350505050565b600054610100900460ff161561024b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100dc565b60005460ff908116101561029d576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146102b457600080fd5b50565b805163ffffffff811681146102cb57600080fd5b919050565b600080600080600080600060e0888a0312156102eb57600080fd5b87516102f68161029f565b60208901519097506103078161029f565b9550610315604089016102b7565b9450610323606089016102b7565b9350610331608089016102b7565b925061033f60a089016102b7565b915061034d60c089016102b7565b905092959891949750929550565b600063ffffffff83168061037f57634e487b7160e01b600052601260045260246000fd5b8063ffffffff84160691505092915050565b60805160a05160c05160e051610100516101205161014051610160516133a061042760003960006118150152600081816104f501526121d8015260006107b501526000818161040c01526120c301526000818161033201526121120152600081816104ce01526120720152600081816107140152611f4701526000818161068c01528181611f9e0152611ffd01526133a06000f3fe608060405234801561001057600080fd5b50600436106102f05760003560e01c80637b8f8b051161019d578063d4540a55116100e9578063f698da25116100a2578063fabc1cbc1161007c578063fabc1cbc14610818578063fbf1e2c11461082b578063fce36c7d1461083e578063ff9f6cce1461085157600080fd5b8063f698da25146107ea578063f8cd8448146107f2578063f96abf2e1461080557600080fd5b8063d4540a5514610764578063de02e50314610777578063e221b2451461078a578063e810ce211461079d578063ea4d3c9b146107b0578063f2fde38b146107d757600080fd5b80639be3d4e411610156578063aebd8bae11610130578063aebd8bae146106c1578063bb7e451f146106ef578063bf21a8aa1461070f578063c46db6061461073657600080fd5b80639be3d4e41461067f5780639d45c28114610687578063a0169ddd146106ae57600080fd5b80637b8f8b05146105fa578063863cb9a914610602578063865c695314610615578063886f1195146106405780638da5cb5b146106535780639104c3191461066457600080fd5b806337838ed01161025c57806358baaa3e116102155780635c975abb116101ef5780635c975abb146105a95780635e9d8348146105b15780636d21117e146105c4578063715018a6146105f257600080fd5b806358baaa3e1461056b578063595c6a671461057e5780635ac86ab71461058657600080fd5b806337838ed0146104c957806339b70e38146104f05780633a8c0786146105175780633ccc861d1461052e5780633efe1db6146105415780634d18cc351461055457600080fd5b8063131433b4116102ae578063131433b414610407578063136439dd1461042e578063149bc8721461044157806322f19a64146104625780632b9f64a41461047557806336af41fa146104b657600080fd5b806218572c146102f557806304a0c5021461032d578063092db007146103695780630e9a53cf146103915780630eb38345146103df57806310d67a2f146103f4575b600080fd5b610318610303366004612bc0565b60d16020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b6103547f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610324565b60cb5461037e90600160e01b900461ffff1681565b60405161ffff9091168152602001610324565b610399610864565b604051610324919060006080820190508251825263ffffffff602084015116602083015263ffffffff604084015116604083015260608301511515606083015292915050565b6103f26103ed366004612beb565b610943565b005b6103f2610402366004612bc0565b6109c5565b6103547f000000000000000000000000000000000000000000000000000000000000000081565b6103f261043c366004612c24565b610a79565b61045461044f366004612c55565b610b64565b604051908152602001610324565b61037e610470366004612c71565b610bda565b61049e610483366004612bc0565b60cc602052600090815260409020546001600160a01b031681565b6040516001600160a01b039091168152602001610324565b6103f26104c4366004612c9f565b610bef565b6103547f000000000000000000000000000000000000000000000000000000000000000081565b61049e7f000000000000000000000000000000000000000000000000000000000000000081565b60cb5461035490600160a01b900463ffffffff1681565b6103f261053c366004612d29565b610db7565b6103f261054f366004612d89565b6110a6565b60cb5461035490600160c01b900463ffffffff1681565b6103f2610579366004612db5565b61129c565b6103f26112ad565b610318610594366004612dd0565b606654600160ff9092169190911b9081161490565b606654610454565b6103186105bf366004612df3565b611375565b6103186105d2366004612e28565b60cf60209081526000928352604080842090915290825290205460ff1681565b6103f2611402565b60ca54610454565b6103f2610610366004612bc0565b611416565b610454610623366004612c71565b60cd60209081526000928352604080842090915290825290205481565b60655461049e906001600160a01b031681565b6033546001600160a01b031661049e565b61049e73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b610399611427565b6103547f000000000000000000000000000000000000000000000000000000000000000081565b6103f26106bc366004612bc0565b6114c5565b6103186106cf366004612e28565b60d260209081526000928352604080842090915290825290205460ff1681565b6104546106fd366004612bc0565b60ce6020526000908152604090205481565b6103547f000000000000000000000000000000000000000000000000000000000000000081565b610318610744366004612e28565b60d060209081526000928352604080842090915290825290205460ff1681565b6103f2610772366004612e71565b611524565b610399610785366004612c24565b61166c565b6103f2610798366004612ee4565b6116fe565b6103546107ab366004612c24565b61170f565b61049e7f000000000000000000000000000000000000000000000000000000000000000081565b6103f26107e5366004612bc0565b61179b565b610454611811565b610454610800366004612c55565b61184e565b6103f2610813366004612db5565b61185f565b6103f2610826366004612c24565b6119b2565b60cb5461049e906001600160a01b031681565b6103f261084c366004612c9f565b611aba565b6103f261085f366004612c9f565b611c2d565b60408051608081018252600080825260208201819052918101829052606081019190915260ca545b801561093f57600060ca6108a1600184612f15565b815481106108b1576108b1612f28565b600091825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff1615801560608301819052919250906109215750806040015163ffffffff164210155b1561092c5792915050565b508061093781612f3e565b91505061088c565b5090565b61094b611dd0565b6001600160a01b038216600081815260d1602052604080822054905160ff9091169284151592841515927f4de6293e668df1398422e1def12118052c1539a03cbfedc145895d48d7685f1c9190a4506001600160a01b0391909116600090815260d160205260409020805460ff1916911515919091179055565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a18573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a3c9190612f55565b6001600160a01b0316336001600160a01b031614610a6d5760405163794821ff60e01b815260040160405180910390fd5b610a7681611e2a565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610ac1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ae59190612f72565b610b0257604051631d77d47760e21b815260040160405180910390fd5b60665481811614610b265760405163c61dca5d60e01b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600080610b746020840184612bc0565b8360200135604051602001610bbd9392919060f89390931b6001600160f81b031916835260609190911b6bffffffffffffffffffffffff19166001830152601582015260350190565b604051602081830303815290604052805190602001209050919050565b60cb54600160e01b900461ffff165b92915050565b606654600190600290811603610c185760405163840a48d560e01b815260040160405180910390fd5b33600090815260d1602052604090205460ff16610c485760405163289d7fb960e01b815260040160405180910390fd5b600260975403610c735760405162461bcd60e51b8152600401610c6a90612f8f565b60405180910390fd5b600260975560005b82811015610dac5736848483818110610c9657610c96612f28565b9050602002810190610ca89190612fc6565b33600081815260ce60209081526040808320549051949550939192610cd3929091859187910161310b565b604051602081830303815290604052805190602001209050610cf483611eba565b33600090815260d0602090815260408083208484529091529020805460ff19166001908117909155610d2790839061313b565b33600081815260ce602052604090819020929092559051829184917f51088b8c89628df3a8174002c2a034d0152fce6af8415d651b2a4734bf27048290610d6f90889061314e565b60405180910390a4610da1333060408601803590610d909060208901612bc0565b6001600160a01b03169291906122ca565b505050600101610c7b565b505060016097555050565b606654600290600490811603610de05760405163840a48d560e01b815260040160405180910390fd5b600260975403610e025760405162461bcd60e51b8152600401610c6a90612f8f565b6002609755600060ca610e186020860186612db5565b63ffffffff1681548110610e2e57610e2e612f28565b600091825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff16151560608201529050610e8f848261233b565b6000610ea16080860160608701612bc0565b6001600160a01b03808216600090815260cc60205260409020549192501680610ec75750805b336001600160a01b03821614610eef576040516208978560e71b815260040160405180910390fd5b60005b610eff60a0880188613161565b90508110156110985736610f1660e08901896131b2565b83818110610f2657610f26612f28565b6001600160a01b038716600090815260cd602090815260408083209302949094019450929091508290610f5b90850185612bc0565b6001600160a01b03166001600160a01b0316815260200190815260200160002054905080826020013511610fa25760405163aa385e8160e01b815260040160405180910390fd5b6000610fb2826020850135612f15565b6001600160a01b038716600090815260cd60209081526040822092935085018035929190610fe09087612bc0565b6001600160a01b0316815260208082019290925260400160002091909155611022908a90839061101290870187612bc0565b6001600160a01b031691906124df565b86516001600160a01b03808b1691878216918916907f9543dbd55580842586a951f0386e24d68a5df99ae29e3b216588b45fd684ce31906110666020890189612bc0565b604080519283526001600160a01b039091166020830152810186905260600160405180910390a4505050600101610ef2565b505060016097555050505050565b6066546003906008908116036110cf5760405163840a48d560e01b815260040160405180910390fd5b60cb546001600160a01b031633146110fa5760405163f6f1e6f160e01b815260040160405180910390fd5b60cb5463ffffffff600160c01b90910481169083161161112d57604051631ca7e50b60e21b815260040160405180910390fd5b428263ffffffff1610611153576040516306957c9160e11b815260040160405180910390fd5b60ca5460cb5460009061117390600160a01b900463ffffffff16426131fc565b6040805160808101825287815263ffffffff878116602080840182815286841685870181815260006060880181815260ca8054600181018255925297517f42d72674974f694b5f5159593243114d38a5c39c89d6b62fee061ff523240ee160029092029182015592517f42d72674974f694b5f5159593243114d38a5c39c89d6b62fee061ff523240ee290930180549151975193871667ffffffffffffffff1990921691909117600160201b978716979097029690961760ff60401b1916600160401b921515929092029190911790945560cb805463ffffffff60c01b1916600160c01b840217905593519283529394508892908616917fecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08910160405180910390a45050505050565b6112a4611dd0565b610a768161250f565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156112f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113199190612f72565b61133657604051631d77d47760e21b815260040160405180910390fd5b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60006113fa8260ca61138a6020830183612db5565b63ffffffff16815481106113a0576113a0612f28565b600091825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff161515606082015261233b565b506001919050565b61140a611dd0565b6114146000612580565b565b61141e611dd0565b610a76816125d2565b60408051608081018252600080825260208201819052918101829052606081019190915260ca805461145b90600190612f15565b8154811061146b5761146b612f28565b600091825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff1615156060820152919050565b33600081815260cc602052604080822080546001600160a01b031981166001600160a01b038781169182179093559251911692839185917fbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca31291a4505050565b600054610100900460ff16158080156115445750600054600160ff909116105b8061155e5750303b15801561155e575060005460ff166001145b6115c15760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610c6a565b6000805460ff1916600117905580156115e4576000805461ff0019166101001790555b6115ec61262e565b60c9556115f986866126c5565b61160287612580565b61160b846125d2565b6116148361250f565b61161d8261274a565b8015611663576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b60408051608081018252600080825260208201819052918101829052606081019190915260ca82815481106116a3576116a3612f28565b600091825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff161515606082015292915050565b611706611dd0565b610a768161274a565b60ca546000905b63ffffffff811615611781578260ca611730600184613218565b63ffffffff168154811061174657611746612f28565b9060005260206000209060020201600001540361176f57611768600182613218565b9392505050565b8061177981613234565b915050611716565b5060405163504570e360e01b815260040160405180910390fd5b6117a3611dd0565b6001600160a01b0381166118085760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610c6a565b610a7681612580565b60007f00000000000000000000000000000000000000000000000000000000000000004603611841575060c95490565b61184961262e565b905090565b60006001610b746020840184612bc0565b6066546003906008908116036118885760405163840a48d560e01b815260040160405180910390fd5b60cb546001600160a01b031633146118b35760405163f6f1e6f160e01b815260040160405180910390fd5b60ca5463ffffffff8316106118db576040516394a8d38960e01b815260040160405180910390fd5b600060ca8363ffffffff16815481106118f6576118f6612f28565b906000526020600020906002020190508060010160089054906101000a900460ff161561193657604051631b14174b60e01b815260040160405180910390fd5b6001810154600160201b900463ffffffff16421061196757604051631748afff60e01b815260040160405180910390fd5b60018101805460ff60401b1916600160401b17905560405163ffffffff8416907fd850e6e5dfa497b72661fa73df2923464eaed9dc2ff1d3cb82bccbfeabe5c41e90600090a2505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a05573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a299190612f55565b6001600160a01b0316336001600160a01b031614611a5a5760405163794821ff60e01b815260040160405180910390fd5b606654198119606654191614611a835760405163c61dca5d60e01b815260040160405180910390fd5b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610b59565b606654600090600190811603611ae35760405163840a48d560e01b815260040160405180910390fd5b600260975403611b055760405162461bcd60e51b8152600401610c6a90612f8f565b600260975560005b82811015610dac5736848483818110611b2857611b28612f28565b9050602002810190611b3a9190612fc6565b33600081815260ce60209081526040808320549051949550939192611b65929091859187910161310b565b604051602081830303815290604052805190602001209050611b8683611eba565b33600090815260cf602090815260408083208484529091529020805460ff19166001908117909155611bb990839061313b565b33600081815260ce602052604090819020929092559051829184917f450a367a380c4e339e5ae7340c8464ef27af7781ad9945cfe8abd828f89e628190611c0190889061314e565b60405180910390a4611c22333060408601803590610d909060208901612bc0565b505050600101611b0d565b606654600490601090811603611c565760405163840a48d560e01b815260040160405180910390fd5b33600090815260d1602052604090205460ff16611c865760405163289d7fb960e01b815260040160405180910390fd5b600260975403611ca85760405162461bcd60e51b8152600401610c6a90612f8f565b600260975560005b82811015610dac5736848483818110611ccb57611ccb612f28565b9050602002810190611cdd9190612fc6565b33600081815260ce60209081526040808320549051949550939192611d08929091859187910161310b565b604051602081830303815290604052805190602001209050611d2983611eba565b33600090815260d2602090815260408083208484529091529020805460ff19166001908117909155611d5c90839061313b565b33600081815260ce602052604090819020929092559051829184917f5251b6fdefcb5d81144e735f69ea4c695fd43b0289ca53dc075033f5fc80068b90611da490889061314e565b60405180910390a4611dc5333060408601803590610d909060208901612bc0565b505050600101611cb0565b6033546001600160a01b031633146114145760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610c6a565b6001600160a01b038116611e51576040516339b190bb60e11b815260040160405180910390fd5b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6000611ec682806131b2565b905011611ee65760405163796cc52560e01b815260040160405180910390fd5b6000816040013511611f0b576040516365e52d5160e11b815260040160405180910390fd5b6f4b3b4ca85a86c47a098a223fffffffff81604001351115611f405760405163070b5a6f60e21b815260040160405180910390fd5b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016611f7760a0830160808401612db5565b63ffffffff161115611f9c57604051630dd0b9f560e21b815260040160405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000611fcd60a0830160808401612db5565b611fd7919061326a565b63ffffffff1615611ffb57604051632380fd2760e21b815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000061202c6080830160608401612db5565b612036919061326a565b63ffffffff161561205a57604051630ea11a0360e31b815260040160405180910390fd5b61206a6080820160608301612db5565b63ffffffff167f000000000000000000000000000000000000000000000000000000000000000063ffffffff16426120a29190612f15565b111580156120eb57506120bb6080820160608301612db5565b63ffffffff167f000000000000000000000000000000000000000000000000000000000000000063ffffffff1611155b6121085760405163041aa75760e11b815260040160405180910390fd5b61213863ffffffff7f0000000000000000000000000000000000000000000000000000000000000000164261313b565b6121486080830160608401612db5565b63ffffffff16111561216d57604051637ee2b44360e01b815260040160405180910390fd5b6000805b61217b83806131b2565b90508110156122c557600061219084806131b2565b838181106121a0576121a0612f28565b6121b69260206040909202019081019150612bc0565b60405163198f077960e21b81526001600160a01b0380831660048301529192507f00000000000000000000000000000000000000000000000000000000000000009091169063663c1de490602401602060405180830381865afa158015612221573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122459190612f72565b8061226c57506001600160a01b03811673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0145b61228957604051632711b74d60e11b815260040160405180910390fd5b806001600160a01b0316836001600160a01b0316106122bb5760405163dfad9ca160e01b815260040160405180910390fd5b9150600101612171565b505050565b6040516001600160a01b03808516602483015283166044820152606481018290526123359085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526127b5565b50505050565b80606001511561235e57604051631b14174b60e01b815260040160405180910390fd5b806040015163ffffffff1642101561238957604051631437a2bb60e31b815260040160405180910390fd5b61239660c0830183613161565b90506123a560a0840184613161565b9050146123c5576040516343714afd60e01b815260040160405180910390fd5b6123d260e08301836131b2565b90506123e160c0840184613161565b905014612401576040516343714afd60e01b815260040160405180910390fd5b805161242d906124176040850160208601612db5565b6124246040860186613292565b86606001612887565b60005b61243d60a0840184613161565b90508110156122c5576124d7608084013561245b60a0860186613161565b8481811061246b5761246b612f28565b90506020020160208101906124809190612db5565b61248d60c0870187613161565b8581811061249d5761249d612f28565b90506020028101906124af9190613292565b6124bc60e08901896131b2565b878181106124cc576124cc612f28565b905060400201612936565b600101612430565b6040516001600160a01b0383166024820152604481018290526122c590849063a9059cbb60e01b906064016122fe565b60cb546040805163ffffffff600160a01b9093048316815291831660208301527faf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3910160405180910390a160cb805463ffffffff909216600160a01b0263ffffffff60a01b19909216919091179055565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60cb546040516001600160a01b038084169216907f237b82f438d75fc568ebab484b75b01d9287b9e98b490b7c23221623b6705dbb90600090a360cb80546001600160a01b0319166001600160a01b0392909216919091179055565b604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b6065546001600160a01b03161580156126e657506001600160a01b03821615155b612703576040516339b190bb60e11b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261274682611e2a565b5050565b60cb546040805161ffff600160e01b9093048316815291831660208301527f8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06910160405180910390a160cb805461ffff909216600160e01b0261ffff60e01b19909216919091179055565b600061280a826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166129759092919063ffffffff16565b8051909150156122c557808060200190518101906128289190612f72565b6122c55760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610c6a565b6128926020836132d9565b6001901b8463ffffffff16106128bb576040516369ca16c960e01b815260040160405180910390fd5b60006128c682610b64565b905061291184848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a92508591505063ffffffff891661298c565b61292e576040516369ca16c960e01b815260040160405180910390fd5b505050505050565b6129416020836132d9565b6001901b8463ffffffff161061296a5760405163054ff4df60e51b815260040160405180910390fd5b60006128c68261184e565b606061298484846000856129a4565b949350505050565b60008361299a868585612ad5565b1495945050505050565b606082471015612a055760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610c6a565b6001600160a01b0385163b612a5c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610c6a565b600080866001600160a01b03168587604051612a789190613311565b60006040518083038185875af1925050503d8060008114612ab5576040519150601f19603f3d011682016040523d82523d6000602084013e612aba565b606091505b5091509150612aca828286612b72565b979650505050505050565b600060208451612ae59190613323565b15612b0357604051630228be0560e31b815260040160405180910390fd5b8260205b85518111612b6957612b1a600285613323565b600003612b3e57816000528086015160205260406000209150600284049350612b57565b8086015160005281602052604060002091506002840493505b612b6260208261313b565b9050612b07565b50949350505050565b60608315612b81575081611768565b825115612b915782518084602001fd5b8160405162461bcd60e51b8152600401610c6a9190613337565b6001600160a01b0381168114610a7657600080fd5b600060208284031215612bd257600080fd5b813561176881612bab565b8015158114610a7657600080fd5b60008060408385031215612bfe57600080fd5b8235612c0981612bab565b91506020830135612c1981612bdd565b809150509250929050565b600060208284031215612c3657600080fd5b5035919050565b600060408284031215612c4f57600080fd5b50919050565b600060408284031215612c6757600080fd5b6117688383612c3d565b60008060408385031215612c8457600080fd5b8235612c8f81612bab565b91506020830135612c1981612bab565b60008060208385031215612cb257600080fd5b823567ffffffffffffffff811115612cc957600080fd5b8301601f81018513612cda57600080fd5b803567ffffffffffffffff811115612cf157600080fd5b8560208260051b8401011115612d0657600080fd5b6020919091019590945092505050565b60006101008284031215612c4f57600080fd5b60008060408385031215612d3c57600080fd5b823567ffffffffffffffff811115612d5357600080fd5b612d5f85828601612d16565b9250506020830135612c1981612bab565b803563ffffffff81168114612d8457600080fd5b919050565b60008060408385031215612d9c57600080fd5b82359150612dac60208401612d70565b90509250929050565b600060208284031215612dc757600080fd5b61176882612d70565b600060208284031215612de257600080fd5b813560ff8116811461176857600080fd5b600060208284031215612e0557600080fd5b813567ffffffffffffffff811115612e1c57600080fd5b61298484828501612d16565b60008060408385031215612e3b57600080fd5b8235612e4681612bab565b946020939093013593505050565b8035612d8481612bab565b803561ffff81168114612d8457600080fd5b60008060008060008060c08789031215612e8a57600080fd5b8635612e9581612bab565b95506020870135612ea581612bab565b9450604087013593506060870135612ebc81612bab565b9250612eca60808801612d70565b9150612ed860a08801612e5f565b90509295509295509295565b600060208284031215612ef657600080fd5b61176882612e5f565b634e487b7160e01b600052601160045260246000fd5b81810381811115610be957610be9612eff565b634e487b7160e01b600052603260045260246000fd5b600081612f4d57612f4d612eff565b506000190190565b600060208284031215612f6757600080fd5b815161176881612bab565b600060208284031215612f8457600080fd5b815161176881612bdd565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b60008235609e19833603018112612fdc57600080fd5b9190910192915050565b81835260208301925060008160005b8481101561304c57813561300881612bab565b6001600160a01b0316865260208201356bffffffffffffffffffffffff811680821461303357600080fd5b6020880152506040958601959190910190600101612ff5565b5093949350505050565b60008135601e1983360301811261306c57600080fd5b820160208101903567ffffffffffffffff81111561308957600080fd5b8060061b360382131561309b57600080fd5b60a085526130ad60a086018284612fe6565b9150506130bc60208401612e54565b6001600160a01b03166020850152604083810135908501526130e060608401612d70565b63ffffffff1660608501526130f760808401612d70565b63ffffffff81166080860152509392505050565b60018060a01b03841681528260208201526060604082015260006131326060830184613056565b95945050505050565b80820180821115610be957610be9612eff565b6020815260006117686020830184613056565b6000808335601e1984360301811261317857600080fd5b83018035915067ffffffffffffffff82111561319357600080fd5b6020019150600581901b36038213156131ab57600080fd5b9250929050565b6000808335601e198436030181126131c957600080fd5b83018035915067ffffffffffffffff8211156131e457600080fd5b6020019150600681901b36038213156131ab57600080fd5b63ffffffff8181168382160190811115610be957610be9612eff565b63ffffffff8281168282160390811115610be957610be9612eff565b600063ffffffff82168061324a5761324a612eff565b6000190192915050565b634e487b7160e01b600052601260045260246000fd5b600063ffffffff83168061328057613280613254565b8063ffffffff84160691505092915050565b6000808335601e198436030181126132a957600080fd5b83018035915067ffffffffffffffff8211156132c457600080fd5b6020019150368190038213156131ab57600080fd5b6000826132e8576132e8613254565b500490565b60005b838110156133085781810151838201526020016132f0565b50506000910152565b60008251612fdc8184602087016132ed565b60008261333257613332613254565b500690565b60208152600082518060208401526133568160408501602087016132ed565b601f01601f1916919091016040019291505056fea26469706673582212200ced2d5bca1cb606ffb68c411b9ebfe082e7825398031d1abe1ad1dc0257114064736f6c634300081b0033",
}

// RewardsCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardsCoordinatorMetaData.ABI instead.
var RewardsCoordinatorABI = RewardsCoordinatorMetaData.ABI

// RewardsCoordinatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RewardsCoordinatorMetaData.Bin instead.
var RewardsCoordinatorBin = RewardsCoordinatorMetaData.Bin

// DeployRewardsCoordinator deploys a new Ethereum contract, binding an instance of RewardsCoordinator to it.
func DeployRewardsCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _delegationManager common.Address, _strategyManager common.Address, _CALCULATION_INTERVAL_SECONDS uint32, _MAX_REWARDS_DURATION uint32, _MAX_RETROACTIVE_LENGTH uint32, _MAX_FUTURE_LENGTH uint32, __GENESIS_REWARDS_TIMESTAMP uint32) (common.Address, *types.Transaction, *RewardsCoordinator, error) {
	parsed, err := RewardsCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RewardsCoordinatorBin), backend, _delegationManager, _strategyManager, _CALCULATION_INTERVAL_SECONDS, _MAX_REWARDS_DURATION, _MAX_RETROACTIVE_LENGTH, _MAX_FUTURE_LENGTH, __GENESIS_REWARDS_TIMESTAMP)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RewardsCoordinator{RewardsCoordinatorCaller: RewardsCoordinatorCaller{contract: contract}, RewardsCoordinatorTransactor: RewardsCoordinatorTransactor{contract: contract}, RewardsCoordinatorFilterer: RewardsCoordinatorFilterer{contract: contract}}, nil
}

// RewardsCoordinator is an auto generated Go binding around an Ethereum contract.
type RewardsCoordinator struct {
	RewardsCoordinatorCaller     // Read-only binding to the contract
	RewardsCoordinatorTransactor // Write-only binding to the contract
	RewardsCoordinatorFilterer   // Log filterer for contract events
}

// RewardsCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardsCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardsCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardsCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardsCoordinatorSession struct {
	Contract     *RewardsCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RewardsCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardsCoordinatorCallerSession struct {
	Contract *RewardsCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// RewardsCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardsCoordinatorTransactorSession struct {
	Contract     *RewardsCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RewardsCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardsCoordinatorRaw struct {
	Contract *RewardsCoordinator // Generic contract binding to access the raw methods on
}

// RewardsCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardsCoordinatorCallerRaw struct {
	Contract *RewardsCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// RewardsCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardsCoordinatorTransactorRaw struct {
	Contract *RewardsCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardsCoordinator creates a new instance of RewardsCoordinator, bound to a specific deployed contract.
func NewRewardsCoordinator(address common.Address, backend bind.ContractBackend) (*RewardsCoordinator, error) {
	contract, err := bindRewardsCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinator{RewardsCoordinatorCaller: RewardsCoordinatorCaller{contract: contract}, RewardsCoordinatorTransactor: RewardsCoordinatorTransactor{contract: contract}, RewardsCoordinatorFilterer: RewardsCoordinatorFilterer{contract: contract}}, nil
}

// NewRewardsCoordinatorCaller creates a new read-only instance of RewardsCoordinator, bound to a specific deployed contract.
func NewRewardsCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*RewardsCoordinatorCaller, error) {
	contract, err := bindRewardsCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorCaller{contract: contract}, nil
}

// NewRewardsCoordinatorTransactor creates a new write-only instance of RewardsCoordinator, bound to a specific deployed contract.
func NewRewardsCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardsCoordinatorTransactor, error) {
	contract, err := bindRewardsCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorTransactor{contract: contract}, nil
}

// NewRewardsCoordinatorFilterer creates a new log filterer instance of RewardsCoordinator, bound to a specific deployed contract.
func NewRewardsCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardsCoordinatorFilterer, error) {
	contract, err := bindRewardsCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorFilterer{contract: contract}, nil
}

// bindRewardsCoordinator binds a generic wrapper to an already deployed contract.
func bindRewardsCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewardsCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardsCoordinator *RewardsCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardsCoordinator.Contract.RewardsCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardsCoordinator *RewardsCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.RewardsCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardsCoordinator *RewardsCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.RewardsCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardsCoordinator *RewardsCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardsCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardsCoordinator *RewardsCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardsCoordinator *RewardsCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.contract.Transact(opts, method, params...)
}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCaller) CALCULATIONINTERVALSECONDS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "CALCULATION_INTERVAL_SECONDS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorSession) CALCULATIONINTERVALSECONDS() (uint32, error) {
	return _RewardsCoordinator.Contract.CALCULATIONINTERVALSECONDS(&_RewardsCoordinator.CallOpts)
}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) CALCULATIONINTERVALSECONDS() (uint32, error) {
	return _RewardsCoordinator.Contract.CALCULATIONINTERVALSECONDS(&_RewardsCoordinator.CallOpts)
}

// GENESISREWARDSTIMESTAMP is a free data retrieval call binding the contract method 0x131433b4.
//
// Solidity: function GENESIS_REWARDS_TIMESTAMP() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCaller) GENESISREWARDSTIMESTAMP(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "GENESIS_REWARDS_TIMESTAMP")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GENESISREWARDSTIMESTAMP is a free data retrieval call binding the contract method 0x131433b4.
//
// Solidity: function GENESIS_REWARDS_TIMESTAMP() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorSession) GENESISREWARDSTIMESTAMP() (uint32, error) {
	return _RewardsCoordinator.Contract.GENESISREWARDSTIMESTAMP(&_RewardsCoordinator.CallOpts)
}

// GENESISREWARDSTIMESTAMP is a free data retrieval call binding the contract method 0x131433b4.
//
// Solidity: function GENESIS_REWARDS_TIMESTAMP() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) GENESISREWARDSTIMESTAMP() (uint32, error) {
	return _RewardsCoordinator.Contract.GENESISREWARDSTIMESTAMP(&_RewardsCoordinator.CallOpts)
}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCaller) MAXFUTURELENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "MAX_FUTURE_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorSession) MAXFUTURELENGTH() (uint32, error) {
	return _RewardsCoordinator.Contract.MAXFUTURELENGTH(&_RewardsCoordinator.CallOpts)
}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) MAXFUTURELENGTH() (uint32, error) {
	return _RewardsCoordinator.Contract.MAXFUTURELENGTH(&_RewardsCoordinator.CallOpts)
}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCaller) MAXRETROACTIVELENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "MAX_RETROACTIVE_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorSession) MAXRETROACTIVELENGTH() (uint32, error) {
	return _RewardsCoordinator.Contract.MAXRETROACTIVELENGTH(&_RewardsCoordinator.CallOpts)
}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) MAXRETROACTIVELENGTH() (uint32, error) {
	return _RewardsCoordinator.Contract.MAXRETROACTIVELENGTH(&_RewardsCoordinator.CallOpts)
}

// MAXREWARDSDURATION is a free data retrieval call binding the contract method 0xbf21a8aa.
//
// Solidity: function MAX_REWARDS_DURATION() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCaller) MAXREWARDSDURATION(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "MAX_REWARDS_DURATION")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXREWARDSDURATION is a free data retrieval call binding the contract method 0xbf21a8aa.
//
// Solidity: function MAX_REWARDS_DURATION() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorSession) MAXREWARDSDURATION() (uint32, error) {
	return _RewardsCoordinator.Contract.MAXREWARDSDURATION(&_RewardsCoordinator.CallOpts)
}

// MAXREWARDSDURATION is a free data retrieval call binding the contract method 0xbf21a8aa.
//
// Solidity: function MAX_REWARDS_DURATION() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) MAXREWARDSDURATION() (uint32, error) {
	return _RewardsCoordinator.Contract.MAXREWARDSDURATION(&_RewardsCoordinator.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCaller) ActivationDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "activationDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorSession) ActivationDelay() (uint32, error) {
	return _RewardsCoordinator.Contract.ActivationDelay(&_RewardsCoordinator.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) ActivationDelay() (uint32, error) {
	return _RewardsCoordinator.Contract.ActivationDelay(&_RewardsCoordinator.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCaller) BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "beaconChainETHStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorSession) BeaconChainETHStrategy() (common.Address, error) {
	return _RewardsCoordinator.Contract.BeaconChainETHStrategy(&_RewardsCoordinator.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _RewardsCoordinator.Contract.BeaconChainETHStrategy(&_RewardsCoordinator.CallOpts)
}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_RewardsCoordinator *RewardsCoordinatorCaller) CalculateEarnerLeafHash(opts *bind.CallOpts, leaf IRewardsCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "calculateEarnerLeafHash", leaf)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_RewardsCoordinator *RewardsCoordinatorSession) CalculateEarnerLeafHash(leaf IRewardsCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _RewardsCoordinator.Contract.CalculateEarnerLeafHash(&_RewardsCoordinator.CallOpts, leaf)
}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) CalculateEarnerLeafHash(leaf IRewardsCoordinatorEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _RewardsCoordinator.Contract.CalculateEarnerLeafHash(&_RewardsCoordinator.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_RewardsCoordinator *RewardsCoordinatorCaller) CalculateTokenLeafHash(opts *bind.CallOpts, leaf IRewardsCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "calculateTokenLeafHash", leaf)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_RewardsCoordinator *RewardsCoordinatorSession) CalculateTokenLeafHash(leaf IRewardsCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	return _RewardsCoordinator.Contract.CalculateTokenLeafHash(&_RewardsCoordinator.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) CalculateTokenLeafHash(leaf IRewardsCoordinatorTokenTreeMerkleLeaf) ([32]byte, error) {
	return _RewardsCoordinator.Contract.CalculateTokenLeafHash(&_RewardsCoordinator.CallOpts, leaf)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCaller) CheckClaim(opts *bind.CallOpts, claim IRewardsCoordinatorRewardsMerkleClaim) (bool, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "checkClaim", claim)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorSession) CheckClaim(claim IRewardsCoordinatorRewardsMerkleClaim) (bool, error) {
	return _RewardsCoordinator.Contract.CheckClaim(&_RewardsCoordinator.CallOpts, claim)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) CheckClaim(claim IRewardsCoordinatorRewardsMerkleClaim) (bool, error) {
	return _RewardsCoordinator.Contract.CheckClaim(&_RewardsCoordinator.CallOpts, claim)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address ) view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCaller) ClaimerFor(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "claimerFor", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address ) view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorSession) ClaimerFor(arg0 common.Address) (common.Address, error) {
	return _RewardsCoordinator.Contract.ClaimerFor(&_RewardsCoordinator.CallOpts, arg0)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address ) view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) ClaimerFor(arg0 common.Address) (common.Address, error) {
	return _RewardsCoordinator.Contract.ClaimerFor(&_RewardsCoordinator.CallOpts, arg0)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address , address ) view returns(uint256)
func (_RewardsCoordinator *RewardsCoordinatorCaller) CumulativeClaimed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "cumulativeClaimed", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address , address ) view returns(uint256)
func (_RewardsCoordinator *RewardsCoordinatorSession) CumulativeClaimed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RewardsCoordinator.Contract.CumulativeClaimed(&_RewardsCoordinator.CallOpts, arg0, arg1)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address , address ) view returns(uint256)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) CumulativeClaimed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RewardsCoordinator.Contract.CumulativeClaimed(&_RewardsCoordinator.CallOpts, arg0, arg1)
}

// CurrRewardsCalculationEndTimestamp is a free data retrieval call binding the contract method 0x4d18cc35.
//
// Solidity: function currRewardsCalculationEndTimestamp() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCaller) CurrRewardsCalculationEndTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "currRewardsCalculationEndTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurrRewardsCalculationEndTimestamp is a free data retrieval call binding the contract method 0x4d18cc35.
//
// Solidity: function currRewardsCalculationEndTimestamp() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorSession) CurrRewardsCalculationEndTimestamp() (uint32, error) {
	return _RewardsCoordinator.Contract.CurrRewardsCalculationEndTimestamp(&_RewardsCoordinator.CallOpts)
}

// CurrRewardsCalculationEndTimestamp is a free data retrieval call binding the contract method 0x4d18cc35.
//
// Solidity: function currRewardsCalculationEndTimestamp() view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) CurrRewardsCalculationEndTimestamp() (uint32, error) {
	return _RewardsCoordinator.Contract.CurrRewardsCalculationEndTimestamp(&_RewardsCoordinator.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorSession) DelegationManager() (common.Address, error) {
	return _RewardsCoordinator.Contract.DelegationManager(&_RewardsCoordinator.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) DelegationManager() (common.Address, error) {
	return _RewardsCoordinator.Contract.DelegationManager(&_RewardsCoordinator.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_RewardsCoordinator *RewardsCoordinatorCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_RewardsCoordinator *RewardsCoordinatorSession) DomainSeparator() ([32]byte, error) {
	return _RewardsCoordinator.Contract.DomainSeparator(&_RewardsCoordinator.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) DomainSeparator() ([32]byte, error) {
	return _RewardsCoordinator.Contract.DomainSeparator(&_RewardsCoordinator.CallOpts)
}

// GetCurrentClaimableDistributionRoot is a free data retrieval call binding the contract method 0x0e9a53cf.
//
// Solidity: function getCurrentClaimableDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorCaller) GetCurrentClaimableDistributionRoot(opts *bind.CallOpts) (IRewardsCoordinatorDistributionRoot, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "getCurrentClaimableDistributionRoot")

	if err != nil {
		return *new(IRewardsCoordinatorDistributionRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(IRewardsCoordinatorDistributionRoot)).(*IRewardsCoordinatorDistributionRoot)

	return out0, err

}

// GetCurrentClaimableDistributionRoot is a free data retrieval call binding the contract method 0x0e9a53cf.
//
// Solidity: function getCurrentClaimableDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorSession) GetCurrentClaimableDistributionRoot() (IRewardsCoordinatorDistributionRoot, error) {
	return _RewardsCoordinator.Contract.GetCurrentClaimableDistributionRoot(&_RewardsCoordinator.CallOpts)
}

// GetCurrentClaimableDistributionRoot is a free data retrieval call binding the contract method 0x0e9a53cf.
//
// Solidity: function getCurrentClaimableDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) GetCurrentClaimableDistributionRoot() (IRewardsCoordinatorDistributionRoot, error) {
	return _RewardsCoordinator.Contract.GetCurrentClaimableDistributionRoot(&_RewardsCoordinator.CallOpts)
}

// GetCurrentDistributionRoot is a free data retrieval call binding the contract method 0x9be3d4e4.
//
// Solidity: function getCurrentDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorCaller) GetCurrentDistributionRoot(opts *bind.CallOpts) (IRewardsCoordinatorDistributionRoot, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "getCurrentDistributionRoot")

	if err != nil {
		return *new(IRewardsCoordinatorDistributionRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(IRewardsCoordinatorDistributionRoot)).(*IRewardsCoordinatorDistributionRoot)

	return out0, err

}

// GetCurrentDistributionRoot is a free data retrieval call binding the contract method 0x9be3d4e4.
//
// Solidity: function getCurrentDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorSession) GetCurrentDistributionRoot() (IRewardsCoordinatorDistributionRoot, error) {
	return _RewardsCoordinator.Contract.GetCurrentDistributionRoot(&_RewardsCoordinator.CallOpts)
}

// GetCurrentDistributionRoot is a free data retrieval call binding the contract method 0x9be3d4e4.
//
// Solidity: function getCurrentDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) GetCurrentDistributionRoot() (IRewardsCoordinatorDistributionRoot, error) {
	return _RewardsCoordinator.Contract.GetCurrentDistributionRoot(&_RewardsCoordinator.CallOpts)
}

// GetDistributionRootAtIndex is a free data retrieval call binding the contract method 0xde02e503.
//
// Solidity: function getDistributionRootAtIndex(uint256 index) view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorCaller) GetDistributionRootAtIndex(opts *bind.CallOpts, index *big.Int) (IRewardsCoordinatorDistributionRoot, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "getDistributionRootAtIndex", index)

	if err != nil {
		return *new(IRewardsCoordinatorDistributionRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(IRewardsCoordinatorDistributionRoot)).(*IRewardsCoordinatorDistributionRoot)

	return out0, err

}

// GetDistributionRootAtIndex is a free data retrieval call binding the contract method 0xde02e503.
//
// Solidity: function getDistributionRootAtIndex(uint256 index) view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorSession) GetDistributionRootAtIndex(index *big.Int) (IRewardsCoordinatorDistributionRoot, error) {
	return _RewardsCoordinator.Contract.GetDistributionRootAtIndex(&_RewardsCoordinator.CallOpts, index)
}

// GetDistributionRootAtIndex is a free data retrieval call binding the contract method 0xde02e503.
//
// Solidity: function getDistributionRootAtIndex(uint256 index) view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) GetDistributionRootAtIndex(index *big.Int) (IRewardsCoordinatorDistributionRoot, error) {
	return _RewardsCoordinator.Contract.GetDistributionRootAtIndex(&_RewardsCoordinator.CallOpts, index)
}

// GetDistributionRootsLength is a free data retrieval call binding the contract method 0x7b8f8b05.
//
// Solidity: function getDistributionRootsLength() view returns(uint256)
func (_RewardsCoordinator *RewardsCoordinatorCaller) GetDistributionRootsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "getDistributionRootsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistributionRootsLength is a free data retrieval call binding the contract method 0x7b8f8b05.
//
// Solidity: function getDistributionRootsLength() view returns(uint256)
func (_RewardsCoordinator *RewardsCoordinatorSession) GetDistributionRootsLength() (*big.Int, error) {
	return _RewardsCoordinator.Contract.GetDistributionRootsLength(&_RewardsCoordinator.CallOpts)
}

// GetDistributionRootsLength is a free data retrieval call binding the contract method 0x7b8f8b05.
//
// Solidity: function getDistributionRootsLength() view returns(uint256)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) GetDistributionRootsLength() (*big.Int, error) {
	return _RewardsCoordinator.Contract.GetDistributionRootsLength(&_RewardsCoordinator.CallOpts)
}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCaller) GetRootIndexFromHash(opts *bind.CallOpts, rootHash [32]byte) (uint32, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "getRootIndexFromHash", rootHash)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorSession) GetRootIndexFromHash(rootHash [32]byte) (uint32, error) {
	return _RewardsCoordinator.Contract.GetRootIndexFromHash(&_RewardsCoordinator.CallOpts, rootHash)
}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) GetRootIndexFromHash(rootHash [32]byte) (uint32, error) {
	return _RewardsCoordinator.Contract.GetRootIndexFromHash(&_RewardsCoordinator.CallOpts, rootHash)
}

// GlobalOperatorCommissionBips is a free data retrieval call binding the contract method 0x092db007.
//
// Solidity: function globalOperatorCommissionBips() view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorCaller) GlobalOperatorCommissionBips(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "globalOperatorCommissionBips")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GlobalOperatorCommissionBips is a free data retrieval call binding the contract method 0x092db007.
//
// Solidity: function globalOperatorCommissionBips() view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorSession) GlobalOperatorCommissionBips() (uint16, error) {
	return _RewardsCoordinator.Contract.GlobalOperatorCommissionBips(&_RewardsCoordinator.CallOpts)
}

// GlobalOperatorCommissionBips is a free data retrieval call binding the contract method 0x092db007.
//
// Solidity: function globalOperatorCommissionBips() view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) GlobalOperatorCommissionBips() (uint16, error) {
	return _RewardsCoordinator.Contract.GlobalOperatorCommissionBips(&_RewardsCoordinator.CallOpts)
}

// IsAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0x6d21117e.
//
// Solidity: function isAVSRewardsSubmissionHash(address , bytes32 ) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCaller) IsAVSRewardsSubmissionHash(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "isAVSRewardsSubmissionHash", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0x6d21117e.
//
// Solidity: function isAVSRewardsSubmissionHash(address , bytes32 ) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorSession) IsAVSRewardsSubmissionHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsAVSRewardsSubmissionHash(&_RewardsCoordinator.CallOpts, arg0, arg1)
}

// IsAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0x6d21117e.
//
// Solidity: function isAVSRewardsSubmissionHash(address , bytes32 ) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) IsAVSRewardsSubmissionHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsAVSRewardsSubmissionHash(&_RewardsCoordinator.CallOpts, arg0, arg1)
}

// IsRewardsForAllSubmitter is a free data retrieval call binding the contract method 0x0018572c.
//
// Solidity: function isRewardsForAllSubmitter(address ) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCaller) IsRewardsForAllSubmitter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "isRewardsForAllSubmitter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewardsForAllSubmitter is a free data retrieval call binding the contract method 0x0018572c.
//
// Solidity: function isRewardsForAllSubmitter(address ) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorSession) IsRewardsForAllSubmitter(arg0 common.Address) (bool, error) {
	return _RewardsCoordinator.Contract.IsRewardsForAllSubmitter(&_RewardsCoordinator.CallOpts, arg0)
}

// IsRewardsForAllSubmitter is a free data retrieval call binding the contract method 0x0018572c.
//
// Solidity: function isRewardsForAllSubmitter(address ) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) IsRewardsForAllSubmitter(arg0 common.Address) (bool, error) {
	return _RewardsCoordinator.Contract.IsRewardsForAllSubmitter(&_RewardsCoordinator.CallOpts, arg0)
}

// IsRewardsSubmissionForAllEarnersHash is a free data retrieval call binding the contract method 0xaebd8bae.
//
// Solidity: function isRewardsSubmissionForAllEarnersHash(address , bytes32 ) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCaller) IsRewardsSubmissionForAllEarnersHash(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "isRewardsSubmissionForAllEarnersHash", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewardsSubmissionForAllEarnersHash is a free data retrieval call binding the contract method 0xaebd8bae.
//
// Solidity: function isRewardsSubmissionForAllEarnersHash(address , bytes32 ) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorSession) IsRewardsSubmissionForAllEarnersHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsRewardsSubmissionForAllEarnersHash(&_RewardsCoordinator.CallOpts, arg0, arg1)
}

// IsRewardsSubmissionForAllEarnersHash is a free data retrieval call binding the contract method 0xaebd8bae.
//
// Solidity: function isRewardsSubmissionForAllEarnersHash(address , bytes32 ) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) IsRewardsSubmissionForAllEarnersHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsRewardsSubmissionForAllEarnersHash(&_RewardsCoordinator.CallOpts, arg0, arg1)
}

// IsRewardsSubmissionForAllHash is a free data retrieval call binding the contract method 0xc46db606.
//
// Solidity: function isRewardsSubmissionForAllHash(address , bytes32 ) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCaller) IsRewardsSubmissionForAllHash(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "isRewardsSubmissionForAllHash", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewardsSubmissionForAllHash is a free data retrieval call binding the contract method 0xc46db606.
//
// Solidity: function isRewardsSubmissionForAllHash(address , bytes32 ) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorSession) IsRewardsSubmissionForAllHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsRewardsSubmissionForAllHash(&_RewardsCoordinator.CallOpts, arg0, arg1)
}

// IsRewardsSubmissionForAllHash is a free data retrieval call binding the contract method 0xc46db606.
//
// Solidity: function isRewardsSubmissionForAllHash(address , bytes32 ) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) IsRewardsSubmissionForAllHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsRewardsSubmissionForAllHash(&_RewardsCoordinator.CallOpts, arg0, arg1)
}

// OperatorCommissionBips is a free data retrieval call binding the contract method 0x22f19a64.
//
// Solidity: function operatorCommissionBips(address operator, address avs) view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorCaller) OperatorCommissionBips(opts *bind.CallOpts, operator common.Address, avs common.Address) (uint16, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "operatorCommissionBips", operator, avs)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// OperatorCommissionBips is a free data retrieval call binding the contract method 0x22f19a64.
//
// Solidity: function operatorCommissionBips(address operator, address avs) view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorSession) OperatorCommissionBips(operator common.Address, avs common.Address) (uint16, error) {
	return _RewardsCoordinator.Contract.OperatorCommissionBips(&_RewardsCoordinator.CallOpts, operator, avs)
}

// OperatorCommissionBips is a free data retrieval call binding the contract method 0x22f19a64.
//
// Solidity: function operatorCommissionBips(address operator, address avs) view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) OperatorCommissionBips(operator common.Address, avs common.Address) (uint16, error) {
	return _RewardsCoordinator.Contract.OperatorCommissionBips(&_RewardsCoordinator.CallOpts, operator, avs)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorSession) Owner() (common.Address, error) {
	return _RewardsCoordinator.Contract.Owner(&_RewardsCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) Owner() (common.Address, error) {
	return _RewardsCoordinator.Contract.Owner(&_RewardsCoordinator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorSession) Paused(index uint8) (bool, error) {
	return _RewardsCoordinator.Contract.Paused(&_RewardsCoordinator.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) Paused(index uint8) (bool, error) {
	return _RewardsCoordinator.Contract.Paused(&_RewardsCoordinator.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_RewardsCoordinator *RewardsCoordinatorCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_RewardsCoordinator *RewardsCoordinatorSession) Paused0() (*big.Int, error) {
	return _RewardsCoordinator.Contract.Paused0(&_RewardsCoordinator.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) Paused0() (*big.Int, error) {
	return _RewardsCoordinator.Contract.Paused0(&_RewardsCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorSession) PauserRegistry() (common.Address, error) {
	return _RewardsCoordinator.Contract.PauserRegistry(&_RewardsCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) PauserRegistry() (common.Address, error) {
	return _RewardsCoordinator.Contract.PauserRegistry(&_RewardsCoordinator.CallOpts)
}

// RewardsUpdater is a free data retrieval call binding the contract method 0xfbf1e2c1.
//
// Solidity: function rewardsUpdater() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCaller) RewardsUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "rewardsUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsUpdater is a free data retrieval call binding the contract method 0xfbf1e2c1.
//
// Solidity: function rewardsUpdater() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorSession) RewardsUpdater() (common.Address, error) {
	return _RewardsCoordinator.Contract.RewardsUpdater(&_RewardsCoordinator.CallOpts)
}

// RewardsUpdater is a free data retrieval call binding the contract method 0xfbf1e2c1.
//
// Solidity: function rewardsUpdater() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) RewardsUpdater() (common.Address, error) {
	return _RewardsCoordinator.Contract.RewardsUpdater(&_RewardsCoordinator.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorSession) StrategyManager() (common.Address, error) {
	return _RewardsCoordinator.Contract.StrategyManager(&_RewardsCoordinator.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) StrategyManager() (common.Address, error) {
	return _RewardsCoordinator.Contract.StrategyManager(&_RewardsCoordinator.CallOpts)
}

// SubmissionNonce is a free data retrieval call binding the contract method 0xbb7e451f.
//
// Solidity: function submissionNonce(address ) view returns(uint256)
func (_RewardsCoordinator *RewardsCoordinatorCaller) SubmissionNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "submissionNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmissionNonce is a free data retrieval call binding the contract method 0xbb7e451f.
//
// Solidity: function submissionNonce(address ) view returns(uint256)
func (_RewardsCoordinator *RewardsCoordinatorSession) SubmissionNonce(arg0 common.Address) (*big.Int, error) {
	return _RewardsCoordinator.Contract.SubmissionNonce(&_RewardsCoordinator.CallOpts, arg0)
}

// SubmissionNonce is a free data retrieval call binding the contract method 0xbb7e451f.
//
// Solidity: function submissionNonce(address ) view returns(uint256)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) SubmissionNonce(arg0 common.Address) (*big.Int, error) {
	return _RewardsCoordinator.Contract.SubmissionNonce(&_RewardsCoordinator.CallOpts, arg0)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) CreateAVSRewardsSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "createAVSRewardsSubmission", rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.CreateAVSRewardsSubmission(&_RewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.CreateAVSRewardsSubmission(&_RewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateRewardsForAllEarners is a paid mutator transaction binding the contract method 0xff9f6cce.
//
// Solidity: function createRewardsForAllEarners(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) CreateRewardsForAllEarners(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "createRewardsForAllEarners", rewardsSubmissions)
}

// CreateRewardsForAllEarners is a paid mutator transaction binding the contract method 0xff9f6cce.
//
// Solidity: function createRewardsForAllEarners(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) CreateRewardsForAllEarners(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.CreateRewardsForAllEarners(&_RewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateRewardsForAllEarners is a paid mutator transaction binding the contract method 0xff9f6cce.
//
// Solidity: function createRewardsForAllEarners(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) CreateRewardsForAllEarners(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.CreateRewardsForAllEarners(&_RewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateRewardsForAllSubmission is a paid mutator transaction binding the contract method 0x36af41fa.
//
// Solidity: function createRewardsForAllSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) CreateRewardsForAllSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "createRewardsForAllSubmission", rewardsSubmissions)
}

// CreateRewardsForAllSubmission is a paid mutator transaction binding the contract method 0x36af41fa.
//
// Solidity: function createRewardsForAllSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) CreateRewardsForAllSubmission(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.CreateRewardsForAllSubmission(&_RewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateRewardsForAllSubmission is a paid mutator transaction binding the contract method 0x36af41fa.
//
// Solidity: function createRewardsForAllSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) CreateRewardsForAllSubmission(rewardsSubmissions []IRewardsCoordinatorRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.CreateRewardsForAllSubmission(&_RewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// DisableRoot is a paid mutator transaction binding the contract method 0xf96abf2e.
//
// Solidity: function disableRoot(uint32 rootIndex) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) DisableRoot(opts *bind.TransactOpts, rootIndex uint32) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "disableRoot", rootIndex)
}

// DisableRoot is a paid mutator transaction binding the contract method 0xf96abf2e.
//
// Solidity: function disableRoot(uint32 rootIndex) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) DisableRoot(rootIndex uint32) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.DisableRoot(&_RewardsCoordinator.TransactOpts, rootIndex)
}

// DisableRoot is a paid mutator transaction binding the contract method 0xf96abf2e.
//
// Solidity: function disableRoot(uint32 rootIndex) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) DisableRoot(rootIndex uint32) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.DisableRoot(&_RewardsCoordinator.TransactOpts, rootIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0xd4540a55.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus, address _rewardsUpdater, uint32 _activationDelay, uint16 _globalCommissionBips) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int, _rewardsUpdater common.Address, _activationDelay uint32, _globalCommissionBips uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "initialize", initialOwner, _pauserRegistry, initialPausedStatus, _rewardsUpdater, _activationDelay, _globalCommissionBips)
}

// Initialize is a paid mutator transaction binding the contract method 0xd4540a55.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus, address _rewardsUpdater, uint32 _activationDelay, uint16 _globalCommissionBips) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int, _rewardsUpdater common.Address, _activationDelay uint32, _globalCommissionBips uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.Initialize(&_RewardsCoordinator.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus, _rewardsUpdater, _activationDelay, _globalCommissionBips)
}

// Initialize is a paid mutator transaction binding the contract method 0xd4540a55.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus, address _rewardsUpdater, uint32 _activationDelay, uint16 _globalCommissionBips) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int, _rewardsUpdater common.Address, _activationDelay uint32, _globalCommissionBips uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.Initialize(&_RewardsCoordinator.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus, _rewardsUpdater, _activationDelay, _globalCommissionBips)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.Pause(&_RewardsCoordinator.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.Pause(&_RewardsCoordinator.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) PauseAll() (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.PauseAll(&_RewardsCoordinator.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) PauseAll() (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.PauseAll(&_RewardsCoordinator.TransactOpts)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) ProcessClaim(opts *bind.TransactOpts, claim IRewardsCoordinatorRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "processClaim", claim, recipient)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) ProcessClaim(claim IRewardsCoordinatorRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.ProcessClaim(&_RewardsCoordinator.TransactOpts, claim, recipient)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) ProcessClaim(claim IRewardsCoordinatorRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.ProcessClaim(&_RewardsCoordinator.TransactOpts, claim, recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.RenounceOwnership(&_RewardsCoordinator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.RenounceOwnership(&_RewardsCoordinator.TransactOpts)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) SetActivationDelay(opts *bind.TransactOpts, _activationDelay uint32) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "setActivationDelay", _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetActivationDelay(&_RewardsCoordinator.TransactOpts, _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetActivationDelay(&_RewardsCoordinator.TransactOpts, _activationDelay)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) SetClaimerFor(opts *bind.TransactOpts, claimer common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "setClaimerFor", claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetClaimerFor(&_RewardsCoordinator.TransactOpts, claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetClaimerFor(&_RewardsCoordinator.TransactOpts, claimer)
}

// SetGlobalOperatorCommission is a paid mutator transaction binding the contract method 0xe221b245.
//
// Solidity: function setGlobalOperatorCommission(uint16 _globalCommissionBips) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) SetGlobalOperatorCommission(opts *bind.TransactOpts, _globalCommissionBips uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "setGlobalOperatorCommission", _globalCommissionBips)
}

// SetGlobalOperatorCommission is a paid mutator transaction binding the contract method 0xe221b245.
//
// Solidity: function setGlobalOperatorCommission(uint16 _globalCommissionBips) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) SetGlobalOperatorCommission(_globalCommissionBips uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetGlobalOperatorCommission(&_RewardsCoordinator.TransactOpts, _globalCommissionBips)
}

// SetGlobalOperatorCommission is a paid mutator transaction binding the contract method 0xe221b245.
//
// Solidity: function setGlobalOperatorCommission(uint16 _globalCommissionBips) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) SetGlobalOperatorCommission(_globalCommissionBips uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetGlobalOperatorCommission(&_RewardsCoordinator.TransactOpts, _globalCommissionBips)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetPauserRegistry(&_RewardsCoordinator.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetPauserRegistry(&_RewardsCoordinator.TransactOpts, newPauserRegistry)
}

// SetRewardsForAllSubmitter is a paid mutator transaction binding the contract method 0x0eb38345.
//
// Solidity: function setRewardsForAllSubmitter(address _submitter, bool _newValue) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) SetRewardsForAllSubmitter(opts *bind.TransactOpts, _submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "setRewardsForAllSubmitter", _submitter, _newValue)
}

// SetRewardsForAllSubmitter is a paid mutator transaction binding the contract method 0x0eb38345.
//
// Solidity: function setRewardsForAllSubmitter(address _submitter, bool _newValue) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) SetRewardsForAllSubmitter(_submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetRewardsForAllSubmitter(&_RewardsCoordinator.TransactOpts, _submitter, _newValue)
}

// SetRewardsForAllSubmitter is a paid mutator transaction binding the contract method 0x0eb38345.
//
// Solidity: function setRewardsForAllSubmitter(address _submitter, bool _newValue) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) SetRewardsForAllSubmitter(_submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetRewardsForAllSubmitter(&_RewardsCoordinator.TransactOpts, _submitter, _newValue)
}

// SetRewardsUpdater is a paid mutator transaction binding the contract method 0x863cb9a9.
//
// Solidity: function setRewardsUpdater(address _rewardsUpdater) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) SetRewardsUpdater(opts *bind.TransactOpts, _rewardsUpdater common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "setRewardsUpdater", _rewardsUpdater)
}

// SetRewardsUpdater is a paid mutator transaction binding the contract method 0x863cb9a9.
//
// Solidity: function setRewardsUpdater(address _rewardsUpdater) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) SetRewardsUpdater(_rewardsUpdater common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetRewardsUpdater(&_RewardsCoordinator.TransactOpts, _rewardsUpdater)
}

// SetRewardsUpdater is a paid mutator transaction binding the contract method 0x863cb9a9.
//
// Solidity: function setRewardsUpdater(address _rewardsUpdater) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) SetRewardsUpdater(_rewardsUpdater common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetRewardsUpdater(&_RewardsCoordinator.TransactOpts, _rewardsUpdater)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 rewardsCalculationEndTimestamp) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) SubmitRoot(opts *bind.TransactOpts, root [32]byte, rewardsCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "submitRoot", root, rewardsCalculationEndTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 rewardsCalculationEndTimestamp) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) SubmitRoot(root [32]byte, rewardsCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SubmitRoot(&_RewardsCoordinator.TransactOpts, root, rewardsCalculationEndTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 rewardsCalculationEndTimestamp) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) SubmitRoot(root [32]byte, rewardsCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SubmitRoot(&_RewardsCoordinator.TransactOpts, root, rewardsCalculationEndTimestamp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.TransferOwnership(&_RewardsCoordinator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.TransferOwnership(&_RewardsCoordinator.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.Unpause(&_RewardsCoordinator.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.Unpause(&_RewardsCoordinator.TransactOpts, newPausedStatus)
}

// RewardsCoordinatorAVSRewardsSubmissionCreatedIterator is returned from FilterAVSRewardsSubmissionCreated and is used to iterate over the raw logs and unpacked data for AVSRewardsSubmissionCreated events raised by the RewardsCoordinator contract.
type RewardsCoordinatorAVSRewardsSubmissionCreatedIterator struct {
	Event *RewardsCoordinatorAVSRewardsSubmissionCreated // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorAVSRewardsSubmissionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorAVSRewardsSubmissionCreated)
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
		it.Event = new(RewardsCoordinatorAVSRewardsSubmissionCreated)
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
func (it *RewardsCoordinatorAVSRewardsSubmissionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorAVSRewardsSubmissionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorAVSRewardsSubmissionCreated represents a AVSRewardsSubmissionCreated event raised by the RewardsCoordinator contract.
type RewardsCoordinatorAVSRewardsSubmissionCreated struct {
	Avs                   common.Address
	SubmissionNonce       *big.Int
	RewardsSubmissionHash [32]byte
	RewardsSubmission     IRewardsCoordinatorRewardsSubmission
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterAVSRewardsSubmissionCreated is a free log retrieval operation binding the contract event 0x450a367a380c4e339e5ae7340c8464ef27af7781ad9945cfe8abd828f89e6281.
//
// Solidity: event AVSRewardsSubmissionCreated(address indexed avs, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterAVSRewardsSubmissionCreated(opts *bind.FilterOpts, avs []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (*RewardsCoordinatorAVSRewardsSubmissionCreatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "AVSRewardsSubmissionCreated", avsRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorAVSRewardsSubmissionCreatedIterator{contract: _RewardsCoordinator.contract, event: "AVSRewardsSubmissionCreated", logs: logs, sub: sub}, nil
}

// WatchAVSRewardsSubmissionCreated is a free log subscription operation binding the contract event 0x450a367a380c4e339e5ae7340c8464ef27af7781ad9945cfe8abd828f89e6281.
//
// Solidity: event AVSRewardsSubmissionCreated(address indexed avs, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchAVSRewardsSubmissionCreated(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorAVSRewardsSubmissionCreated, avs []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "AVSRewardsSubmissionCreated", avsRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorAVSRewardsSubmissionCreated)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "AVSRewardsSubmissionCreated", log); err != nil {
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

// ParseAVSRewardsSubmissionCreated is a log parse operation binding the contract event 0x450a367a380c4e339e5ae7340c8464ef27af7781ad9945cfe8abd828f89e6281.
//
// Solidity: event AVSRewardsSubmissionCreated(address indexed avs, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseAVSRewardsSubmissionCreated(log types.Log) (*RewardsCoordinatorAVSRewardsSubmissionCreated, error) {
	event := new(RewardsCoordinatorAVSRewardsSubmissionCreated)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "AVSRewardsSubmissionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorActivationDelaySetIterator is returned from FilterActivationDelaySet and is used to iterate over the raw logs and unpacked data for ActivationDelaySet events raised by the RewardsCoordinator contract.
type RewardsCoordinatorActivationDelaySetIterator struct {
	Event *RewardsCoordinatorActivationDelaySet // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorActivationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorActivationDelaySet)
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
		it.Event = new(RewardsCoordinatorActivationDelaySet)
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
func (it *RewardsCoordinatorActivationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorActivationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorActivationDelaySet represents a ActivationDelaySet event raised by the RewardsCoordinator contract.
type RewardsCoordinatorActivationDelaySet struct {
	OldActivationDelay uint32
	NewActivationDelay uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterActivationDelaySet is a free log retrieval operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterActivationDelaySet(opts *bind.FilterOpts) (*RewardsCoordinatorActivationDelaySetIterator, error) {

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorActivationDelaySetIterator{contract: _RewardsCoordinator.contract, event: "ActivationDelaySet", logs: logs, sub: sub}, nil
}

// WatchActivationDelaySet is a free log subscription operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchActivationDelaySet(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorActivationDelaySet) (event.Subscription, error) {

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorActivationDelaySet)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
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
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseActivationDelaySet(log types.Log) (*RewardsCoordinatorActivationDelaySet, error) {
	event := new(RewardsCoordinatorActivationDelaySet)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorClaimerForSetIterator is returned from FilterClaimerForSet and is used to iterate over the raw logs and unpacked data for ClaimerForSet events raised by the RewardsCoordinator contract.
type RewardsCoordinatorClaimerForSetIterator struct {
	Event *RewardsCoordinatorClaimerForSet // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorClaimerForSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorClaimerForSet)
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
		it.Event = new(RewardsCoordinatorClaimerForSet)
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
func (it *RewardsCoordinatorClaimerForSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorClaimerForSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorClaimerForSet represents a ClaimerForSet event raised by the RewardsCoordinator contract.
type RewardsCoordinatorClaimerForSet struct {
	Earner     common.Address
	OldClaimer common.Address
	Claimer    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimerForSet is a free log retrieval operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterClaimerForSet(opts *bind.FilterOpts, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (*RewardsCoordinatorClaimerForSetIterator, error) {

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

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "ClaimerForSet", earnerRule, oldClaimerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorClaimerForSetIterator{contract: _RewardsCoordinator.contract, event: "ClaimerForSet", logs: logs, sub: sub}, nil
}

// WatchClaimerForSet is a free log subscription operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchClaimerForSet(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorClaimerForSet, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "ClaimerForSet", earnerRule, oldClaimerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorClaimerForSet)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "ClaimerForSet", log); err != nil {
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
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseClaimerForSet(log types.Log) (*RewardsCoordinatorClaimerForSet, error) {
	event := new(RewardsCoordinatorClaimerForSet)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "ClaimerForSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorDistributionRootDisabledIterator is returned from FilterDistributionRootDisabled and is used to iterate over the raw logs and unpacked data for DistributionRootDisabled events raised by the RewardsCoordinator contract.
type RewardsCoordinatorDistributionRootDisabledIterator struct {
	Event *RewardsCoordinatorDistributionRootDisabled // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorDistributionRootDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorDistributionRootDisabled)
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
		it.Event = new(RewardsCoordinatorDistributionRootDisabled)
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
func (it *RewardsCoordinatorDistributionRootDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorDistributionRootDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorDistributionRootDisabled represents a DistributionRootDisabled event raised by the RewardsCoordinator contract.
type RewardsCoordinatorDistributionRootDisabled struct {
	RootIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistributionRootDisabled is a free log retrieval operation binding the contract event 0xd850e6e5dfa497b72661fa73df2923464eaed9dc2ff1d3cb82bccbfeabe5c41e.
//
// Solidity: event DistributionRootDisabled(uint32 indexed rootIndex)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterDistributionRootDisabled(opts *bind.FilterOpts, rootIndex []uint32) (*RewardsCoordinatorDistributionRootDisabledIterator, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "DistributionRootDisabled", rootIndexRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorDistributionRootDisabledIterator{contract: _RewardsCoordinator.contract, event: "DistributionRootDisabled", logs: logs, sub: sub}, nil
}

// WatchDistributionRootDisabled is a free log subscription operation binding the contract event 0xd850e6e5dfa497b72661fa73df2923464eaed9dc2ff1d3cb82bccbfeabe5c41e.
//
// Solidity: event DistributionRootDisabled(uint32 indexed rootIndex)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchDistributionRootDisabled(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorDistributionRootDisabled, rootIndex []uint32) (event.Subscription, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "DistributionRootDisabled", rootIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorDistributionRootDisabled)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "DistributionRootDisabled", log); err != nil {
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

// ParseDistributionRootDisabled is a log parse operation binding the contract event 0xd850e6e5dfa497b72661fa73df2923464eaed9dc2ff1d3cb82bccbfeabe5c41e.
//
// Solidity: event DistributionRootDisabled(uint32 indexed rootIndex)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseDistributionRootDisabled(log types.Log) (*RewardsCoordinatorDistributionRootDisabled, error) {
	event := new(RewardsCoordinatorDistributionRootDisabled)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "DistributionRootDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorDistributionRootSubmittedIterator is returned from FilterDistributionRootSubmitted and is used to iterate over the raw logs and unpacked data for DistributionRootSubmitted events raised by the RewardsCoordinator contract.
type RewardsCoordinatorDistributionRootSubmittedIterator struct {
	Event *RewardsCoordinatorDistributionRootSubmitted // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorDistributionRootSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorDistributionRootSubmitted)
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
		it.Event = new(RewardsCoordinatorDistributionRootSubmitted)
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
func (it *RewardsCoordinatorDistributionRootSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorDistributionRootSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorDistributionRootSubmitted represents a DistributionRootSubmitted event raised by the RewardsCoordinator contract.
type RewardsCoordinatorDistributionRootSubmitted struct {
	RootIndex                      uint32
	Root                           [32]byte
	RewardsCalculationEndTimestamp uint32
	ActivatedAt                    uint32
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterDistributionRootSubmitted is a free log retrieval operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed rewardsCalculationEndTimestamp, uint32 activatedAt)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterDistributionRootSubmitted(opts *bind.FilterOpts, rootIndex []uint32, root [][32]byte, rewardsCalculationEndTimestamp []uint32) (*RewardsCoordinatorDistributionRootSubmittedIterator, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}
	var rewardsCalculationEndTimestampRule []interface{}
	for _, rewardsCalculationEndTimestampItem := range rewardsCalculationEndTimestamp {
		rewardsCalculationEndTimestampRule = append(rewardsCalculationEndTimestampRule, rewardsCalculationEndTimestampItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "DistributionRootSubmitted", rootIndexRule, rootRule, rewardsCalculationEndTimestampRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorDistributionRootSubmittedIterator{contract: _RewardsCoordinator.contract, event: "DistributionRootSubmitted", logs: logs, sub: sub}, nil
}

// WatchDistributionRootSubmitted is a free log subscription operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed rewardsCalculationEndTimestamp, uint32 activatedAt)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchDistributionRootSubmitted(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorDistributionRootSubmitted, rootIndex []uint32, root [][32]byte, rewardsCalculationEndTimestamp []uint32) (event.Subscription, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}
	var rewardsCalculationEndTimestampRule []interface{}
	for _, rewardsCalculationEndTimestampItem := range rewardsCalculationEndTimestamp {
		rewardsCalculationEndTimestampRule = append(rewardsCalculationEndTimestampRule, rewardsCalculationEndTimestampItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "DistributionRootSubmitted", rootIndexRule, rootRule, rewardsCalculationEndTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorDistributionRootSubmitted)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "DistributionRootSubmitted", log); err != nil {
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
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed rewardsCalculationEndTimestamp, uint32 activatedAt)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseDistributionRootSubmitted(log types.Log) (*RewardsCoordinatorDistributionRootSubmitted, error) {
	event := new(RewardsCoordinatorDistributionRootSubmitted)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "DistributionRootSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorGlobalCommissionBipsSetIterator is returned from FilterGlobalCommissionBipsSet and is used to iterate over the raw logs and unpacked data for GlobalCommissionBipsSet events raised by the RewardsCoordinator contract.
type RewardsCoordinatorGlobalCommissionBipsSetIterator struct {
	Event *RewardsCoordinatorGlobalCommissionBipsSet // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorGlobalCommissionBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorGlobalCommissionBipsSet)
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
		it.Event = new(RewardsCoordinatorGlobalCommissionBipsSet)
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
func (it *RewardsCoordinatorGlobalCommissionBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorGlobalCommissionBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorGlobalCommissionBipsSet represents a GlobalCommissionBipsSet event raised by the RewardsCoordinator contract.
type RewardsCoordinatorGlobalCommissionBipsSet struct {
	OldGlobalCommissionBips uint16
	NewGlobalCommissionBips uint16
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterGlobalCommissionBipsSet is a free log retrieval operation binding the contract event 0x8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06.
//
// Solidity: event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterGlobalCommissionBipsSet(opts *bind.FilterOpts) (*RewardsCoordinatorGlobalCommissionBipsSetIterator, error) {

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "GlobalCommissionBipsSet")
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorGlobalCommissionBipsSetIterator{contract: _RewardsCoordinator.contract, event: "GlobalCommissionBipsSet", logs: logs, sub: sub}, nil
}

// WatchGlobalCommissionBipsSet is a free log subscription operation binding the contract event 0x8cdc428b0431b82d1619763f443a48197db344ba96905f3949643acd1c863a06.
//
// Solidity: event GlobalCommissionBipsSet(uint16 oldGlobalCommissionBips, uint16 newGlobalCommissionBips)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchGlobalCommissionBipsSet(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorGlobalCommissionBipsSet) (event.Subscription, error) {

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "GlobalCommissionBipsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorGlobalCommissionBipsSet)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "GlobalCommissionBipsSet", log); err != nil {
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
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseGlobalCommissionBipsSet(log types.Log) (*RewardsCoordinatorGlobalCommissionBipsSet, error) {
	event := new(RewardsCoordinatorGlobalCommissionBipsSet)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "GlobalCommissionBipsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RewardsCoordinator contract.
type RewardsCoordinatorInitializedIterator struct {
	Event *RewardsCoordinatorInitialized // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorInitialized)
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
		it.Event = new(RewardsCoordinatorInitialized)
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
func (it *RewardsCoordinatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorInitialized represents a Initialized event raised by the RewardsCoordinator contract.
type RewardsCoordinatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*RewardsCoordinatorInitializedIterator, error) {

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorInitializedIterator{contract: _RewardsCoordinator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorInitialized) (event.Subscription, error) {

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorInitialized)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseInitialized(log types.Log) (*RewardsCoordinatorInitialized, error) {
	event := new(RewardsCoordinatorInitialized)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RewardsCoordinator contract.
type RewardsCoordinatorOwnershipTransferredIterator struct {
	Event *RewardsCoordinatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorOwnershipTransferred)
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
		it.Event = new(RewardsCoordinatorOwnershipTransferred)
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
func (it *RewardsCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorOwnershipTransferred represents a OwnershipTransferred event raised by the RewardsCoordinator contract.
type RewardsCoordinatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RewardsCoordinatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorOwnershipTransferredIterator{contract: _RewardsCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorOwnershipTransferred)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*RewardsCoordinatorOwnershipTransferred, error) {
	event := new(RewardsCoordinatorOwnershipTransferred)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the RewardsCoordinator contract.
type RewardsCoordinatorPausedIterator struct {
	Event *RewardsCoordinatorPaused // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorPaused)
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
		it.Event = new(RewardsCoordinatorPaused)
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
func (it *RewardsCoordinatorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorPaused represents a Paused event raised by the RewardsCoordinator contract.
type RewardsCoordinatorPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*RewardsCoordinatorPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorPausedIterator{contract: _RewardsCoordinator.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorPaused)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParsePaused(log types.Log) (*RewardsCoordinatorPaused, error) {
	event := new(RewardsCoordinatorPaused)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the RewardsCoordinator contract.
type RewardsCoordinatorPauserRegistrySetIterator struct {
	Event *RewardsCoordinatorPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorPauserRegistrySet)
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
		it.Event = new(RewardsCoordinatorPauserRegistrySet)
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
func (it *RewardsCoordinatorPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorPauserRegistrySet represents a PauserRegistrySet event raised by the RewardsCoordinator contract.
type RewardsCoordinatorPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*RewardsCoordinatorPauserRegistrySetIterator, error) {

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorPauserRegistrySetIterator{contract: _RewardsCoordinator.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorPauserRegistrySet)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParsePauserRegistrySet(log types.Log) (*RewardsCoordinatorPauserRegistrySet, error) {
	event := new(RewardsCoordinatorPauserRegistrySet)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the RewardsCoordinator contract.
type RewardsCoordinatorRewardsClaimedIterator struct {
	Event *RewardsCoordinatorRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorRewardsClaimed)
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
		it.Event = new(RewardsCoordinatorRewardsClaimed)
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
func (it *RewardsCoordinatorRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorRewardsClaimed represents a RewardsClaimed event raised by the RewardsCoordinator contract.
type RewardsCoordinatorRewardsClaimed struct {
	Root          [32]byte
	Earner        common.Address
	Claimer       common.Address
	Recipient     common.Address
	Token         common.Address
	ClaimedAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0x9543dbd55580842586a951f0386e24d68a5df99ae29e3b216588b45fd684ce31.
//
// Solidity: event RewardsClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, earner []common.Address, claimer []common.Address, recipient []common.Address) (*RewardsCoordinatorRewardsClaimedIterator, error) {

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

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "RewardsClaimed", earnerRule, claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorRewardsClaimedIterator{contract: _RewardsCoordinator.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0x9543dbd55580842586a951f0386e24d68a5df99ae29e3b216588b45fd684ce31.
//
// Solidity: event RewardsClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorRewardsClaimed, earner []common.Address, claimer []common.Address, recipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "RewardsClaimed", earnerRule, claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorRewardsClaimed)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0x9543dbd55580842586a951f0386e24d68a5df99ae29e3b216588b45fd684ce31.
//
// Solidity: event RewardsClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseRewardsClaimed(log types.Log) (*RewardsCoordinatorRewardsClaimed, error) {
	event := new(RewardsCoordinatorRewardsClaimed)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorRewardsForAllSubmitterSetIterator is returned from FilterRewardsForAllSubmitterSet and is used to iterate over the raw logs and unpacked data for RewardsForAllSubmitterSet events raised by the RewardsCoordinator contract.
type RewardsCoordinatorRewardsForAllSubmitterSetIterator struct {
	Event *RewardsCoordinatorRewardsForAllSubmitterSet // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorRewardsForAllSubmitterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorRewardsForAllSubmitterSet)
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
		it.Event = new(RewardsCoordinatorRewardsForAllSubmitterSet)
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
func (it *RewardsCoordinatorRewardsForAllSubmitterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorRewardsForAllSubmitterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorRewardsForAllSubmitterSet represents a RewardsForAllSubmitterSet event raised by the RewardsCoordinator contract.
type RewardsCoordinatorRewardsForAllSubmitterSet struct {
	RewardsForAllSubmitter common.Address
	OldValue               bool
	NewValue               bool
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRewardsForAllSubmitterSet is a free log retrieval operation binding the contract event 0x4de6293e668df1398422e1def12118052c1539a03cbfedc145895d48d7685f1c.
//
// Solidity: event RewardsForAllSubmitterSet(address indexed rewardsForAllSubmitter, bool indexed oldValue, bool indexed newValue)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterRewardsForAllSubmitterSet(opts *bind.FilterOpts, rewardsForAllSubmitter []common.Address, oldValue []bool, newValue []bool) (*RewardsCoordinatorRewardsForAllSubmitterSetIterator, error) {

	var rewardsForAllSubmitterRule []interface{}
	for _, rewardsForAllSubmitterItem := range rewardsForAllSubmitter {
		rewardsForAllSubmitterRule = append(rewardsForAllSubmitterRule, rewardsForAllSubmitterItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}
	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "RewardsForAllSubmitterSet", rewardsForAllSubmitterRule, oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorRewardsForAllSubmitterSetIterator{contract: _RewardsCoordinator.contract, event: "RewardsForAllSubmitterSet", logs: logs, sub: sub}, nil
}

// WatchRewardsForAllSubmitterSet is a free log subscription operation binding the contract event 0x4de6293e668df1398422e1def12118052c1539a03cbfedc145895d48d7685f1c.
//
// Solidity: event RewardsForAllSubmitterSet(address indexed rewardsForAllSubmitter, bool indexed oldValue, bool indexed newValue)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchRewardsForAllSubmitterSet(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorRewardsForAllSubmitterSet, rewardsForAllSubmitter []common.Address, oldValue []bool, newValue []bool) (event.Subscription, error) {

	var rewardsForAllSubmitterRule []interface{}
	for _, rewardsForAllSubmitterItem := range rewardsForAllSubmitter {
		rewardsForAllSubmitterRule = append(rewardsForAllSubmitterRule, rewardsForAllSubmitterItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}
	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "RewardsForAllSubmitterSet", rewardsForAllSubmitterRule, oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorRewardsForAllSubmitterSet)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "RewardsForAllSubmitterSet", log); err != nil {
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

// ParseRewardsForAllSubmitterSet is a log parse operation binding the contract event 0x4de6293e668df1398422e1def12118052c1539a03cbfedc145895d48d7685f1c.
//
// Solidity: event RewardsForAllSubmitterSet(address indexed rewardsForAllSubmitter, bool indexed oldValue, bool indexed newValue)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseRewardsForAllSubmitterSet(log types.Log) (*RewardsCoordinatorRewardsForAllSubmitterSet, error) {
	event := new(RewardsCoordinatorRewardsForAllSubmitterSet)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "RewardsForAllSubmitterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorRewardsSubmissionForAllCreatedIterator is returned from FilterRewardsSubmissionForAllCreated and is used to iterate over the raw logs and unpacked data for RewardsSubmissionForAllCreated events raised by the RewardsCoordinator contract.
type RewardsCoordinatorRewardsSubmissionForAllCreatedIterator struct {
	Event *RewardsCoordinatorRewardsSubmissionForAllCreated // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorRewardsSubmissionForAllCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorRewardsSubmissionForAllCreated)
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
		it.Event = new(RewardsCoordinatorRewardsSubmissionForAllCreated)
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
func (it *RewardsCoordinatorRewardsSubmissionForAllCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorRewardsSubmissionForAllCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorRewardsSubmissionForAllCreated represents a RewardsSubmissionForAllCreated event raised by the RewardsCoordinator contract.
type RewardsCoordinatorRewardsSubmissionForAllCreated struct {
	Submitter             common.Address
	SubmissionNonce       *big.Int
	RewardsSubmissionHash [32]byte
	RewardsSubmission     IRewardsCoordinatorRewardsSubmission
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRewardsSubmissionForAllCreated is a free log retrieval operation binding the contract event 0x51088b8c89628df3a8174002c2a034d0152fce6af8415d651b2a4734bf270482.
//
// Solidity: event RewardsSubmissionForAllCreated(address indexed submitter, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterRewardsSubmissionForAllCreated(opts *bind.FilterOpts, submitter []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (*RewardsCoordinatorRewardsSubmissionForAllCreatedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "RewardsSubmissionForAllCreated", submitterRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorRewardsSubmissionForAllCreatedIterator{contract: _RewardsCoordinator.contract, event: "RewardsSubmissionForAllCreated", logs: logs, sub: sub}, nil
}

// WatchRewardsSubmissionForAllCreated is a free log subscription operation binding the contract event 0x51088b8c89628df3a8174002c2a034d0152fce6af8415d651b2a4734bf270482.
//
// Solidity: event RewardsSubmissionForAllCreated(address indexed submitter, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchRewardsSubmissionForAllCreated(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorRewardsSubmissionForAllCreated, submitter []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "RewardsSubmissionForAllCreated", submitterRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorRewardsSubmissionForAllCreated)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "RewardsSubmissionForAllCreated", log); err != nil {
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

// ParseRewardsSubmissionForAllCreated is a log parse operation binding the contract event 0x51088b8c89628df3a8174002c2a034d0152fce6af8415d651b2a4734bf270482.
//
// Solidity: event RewardsSubmissionForAllCreated(address indexed submitter, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseRewardsSubmissionForAllCreated(log types.Log) (*RewardsCoordinatorRewardsSubmissionForAllCreated, error) {
	event := new(RewardsCoordinatorRewardsSubmissionForAllCreated)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "RewardsSubmissionForAllCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator is returned from FilterRewardsSubmissionForAllEarnersCreated and is used to iterate over the raw logs and unpacked data for RewardsSubmissionForAllEarnersCreated events raised by the RewardsCoordinator contract.
type RewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator struct {
	Event *RewardsCoordinatorRewardsSubmissionForAllEarnersCreated // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorRewardsSubmissionForAllEarnersCreated)
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
		it.Event = new(RewardsCoordinatorRewardsSubmissionForAllEarnersCreated)
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
func (it *RewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorRewardsSubmissionForAllEarnersCreated represents a RewardsSubmissionForAllEarnersCreated event raised by the RewardsCoordinator contract.
type RewardsCoordinatorRewardsSubmissionForAllEarnersCreated struct {
	TokenHopper           common.Address
	SubmissionNonce       *big.Int
	RewardsSubmissionHash [32]byte
	RewardsSubmission     IRewardsCoordinatorRewardsSubmission
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRewardsSubmissionForAllEarnersCreated is a free log retrieval operation binding the contract event 0x5251b6fdefcb5d81144e735f69ea4c695fd43b0289ca53dc075033f5fc80068b.
//
// Solidity: event RewardsSubmissionForAllEarnersCreated(address indexed tokenHopper, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterRewardsSubmissionForAllEarnersCreated(opts *bind.FilterOpts, tokenHopper []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (*RewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator, error) {

	var tokenHopperRule []interface{}
	for _, tokenHopperItem := range tokenHopper {
		tokenHopperRule = append(tokenHopperRule, tokenHopperItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "RewardsSubmissionForAllEarnersCreated", tokenHopperRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator{contract: _RewardsCoordinator.contract, event: "RewardsSubmissionForAllEarnersCreated", logs: logs, sub: sub}, nil
}

// WatchRewardsSubmissionForAllEarnersCreated is a free log subscription operation binding the contract event 0x5251b6fdefcb5d81144e735f69ea4c695fd43b0289ca53dc075033f5fc80068b.
//
// Solidity: event RewardsSubmissionForAllEarnersCreated(address indexed tokenHopper, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchRewardsSubmissionForAllEarnersCreated(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorRewardsSubmissionForAllEarnersCreated, tokenHopper []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (event.Subscription, error) {

	var tokenHopperRule []interface{}
	for _, tokenHopperItem := range tokenHopper {
		tokenHopperRule = append(tokenHopperRule, tokenHopperItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "RewardsSubmissionForAllEarnersCreated", tokenHopperRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorRewardsSubmissionForAllEarnersCreated)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "RewardsSubmissionForAllEarnersCreated", log); err != nil {
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

// ParseRewardsSubmissionForAllEarnersCreated is a log parse operation binding the contract event 0x5251b6fdefcb5d81144e735f69ea4c695fd43b0289ca53dc075033f5fc80068b.
//
// Solidity: event RewardsSubmissionForAllEarnersCreated(address indexed tokenHopper, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseRewardsSubmissionForAllEarnersCreated(log types.Log) (*RewardsCoordinatorRewardsSubmissionForAllEarnersCreated, error) {
	event := new(RewardsCoordinatorRewardsSubmissionForAllEarnersCreated)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "RewardsSubmissionForAllEarnersCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorRewardsUpdaterSetIterator is returned from FilterRewardsUpdaterSet and is used to iterate over the raw logs and unpacked data for RewardsUpdaterSet events raised by the RewardsCoordinator contract.
type RewardsCoordinatorRewardsUpdaterSetIterator struct {
	Event *RewardsCoordinatorRewardsUpdaterSet // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorRewardsUpdaterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorRewardsUpdaterSet)
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
		it.Event = new(RewardsCoordinatorRewardsUpdaterSet)
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
func (it *RewardsCoordinatorRewardsUpdaterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorRewardsUpdaterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorRewardsUpdaterSet represents a RewardsUpdaterSet event raised by the RewardsCoordinator contract.
type RewardsCoordinatorRewardsUpdaterSet struct {
	OldRewardsUpdater common.Address
	NewRewardsUpdater common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRewardsUpdaterSet is a free log retrieval operation binding the contract event 0x237b82f438d75fc568ebab484b75b01d9287b9e98b490b7c23221623b6705dbb.
//
// Solidity: event RewardsUpdaterSet(address indexed oldRewardsUpdater, address indexed newRewardsUpdater)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterRewardsUpdaterSet(opts *bind.FilterOpts, oldRewardsUpdater []common.Address, newRewardsUpdater []common.Address) (*RewardsCoordinatorRewardsUpdaterSetIterator, error) {

	var oldRewardsUpdaterRule []interface{}
	for _, oldRewardsUpdaterItem := range oldRewardsUpdater {
		oldRewardsUpdaterRule = append(oldRewardsUpdaterRule, oldRewardsUpdaterItem)
	}
	var newRewardsUpdaterRule []interface{}
	for _, newRewardsUpdaterItem := range newRewardsUpdater {
		newRewardsUpdaterRule = append(newRewardsUpdaterRule, newRewardsUpdaterItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "RewardsUpdaterSet", oldRewardsUpdaterRule, newRewardsUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorRewardsUpdaterSetIterator{contract: _RewardsCoordinator.contract, event: "RewardsUpdaterSet", logs: logs, sub: sub}, nil
}

// WatchRewardsUpdaterSet is a free log subscription operation binding the contract event 0x237b82f438d75fc568ebab484b75b01d9287b9e98b490b7c23221623b6705dbb.
//
// Solidity: event RewardsUpdaterSet(address indexed oldRewardsUpdater, address indexed newRewardsUpdater)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchRewardsUpdaterSet(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorRewardsUpdaterSet, oldRewardsUpdater []common.Address, newRewardsUpdater []common.Address) (event.Subscription, error) {

	var oldRewardsUpdaterRule []interface{}
	for _, oldRewardsUpdaterItem := range oldRewardsUpdater {
		oldRewardsUpdaterRule = append(oldRewardsUpdaterRule, oldRewardsUpdaterItem)
	}
	var newRewardsUpdaterRule []interface{}
	for _, newRewardsUpdaterItem := range newRewardsUpdater {
		newRewardsUpdaterRule = append(newRewardsUpdaterRule, newRewardsUpdaterItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "RewardsUpdaterSet", oldRewardsUpdaterRule, newRewardsUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorRewardsUpdaterSet)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "RewardsUpdaterSet", log); err != nil {
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

// ParseRewardsUpdaterSet is a log parse operation binding the contract event 0x237b82f438d75fc568ebab484b75b01d9287b9e98b490b7c23221623b6705dbb.
//
// Solidity: event RewardsUpdaterSet(address indexed oldRewardsUpdater, address indexed newRewardsUpdater)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseRewardsUpdaterSet(log types.Log) (*RewardsCoordinatorRewardsUpdaterSet, error) {
	event := new(RewardsCoordinatorRewardsUpdaterSet)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "RewardsUpdaterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the RewardsCoordinator contract.
type RewardsCoordinatorUnpausedIterator struct {
	Event *RewardsCoordinatorUnpaused // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorUnpaused)
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
		it.Event = new(RewardsCoordinatorUnpaused)
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
func (it *RewardsCoordinatorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorUnpaused represents a Unpaused event raised by the RewardsCoordinator contract.
type RewardsCoordinatorUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*RewardsCoordinatorUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorUnpausedIterator{contract: _RewardsCoordinator.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorUnpaused)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseUnpaused(log types.Log) (*RewardsCoordinatorUnpaused, error) {
	event := new(RewardsCoordinatorUnpaused)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
