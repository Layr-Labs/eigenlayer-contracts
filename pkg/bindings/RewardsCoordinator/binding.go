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

// IRewardsCoordinatorTypesDistributionRoot is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesDistributionRoot struct {
	Root                           [32]byte
	RewardsCalculationEndTimestamp uint32
	ActivatedAt                    uint32
	Disabled                       bool
}

// IRewardsCoordinatorTypesEarnerTreeMerkleLeaf is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesEarnerTreeMerkleLeaf struct {
	Earner          common.Address
	EarnerTokenRoot [32]byte
}

// IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission struct {
	StrategiesAndMultipliers []IRewardsCoordinatorTypesStrategyAndMultiplier
	Token                    common.Address
	OperatorRewards          []IRewardsCoordinatorTypesOperatorReward
	StartTimestamp           uint32
	Duration                 uint32
	Description              string
}

// IRewardsCoordinatorTypesOperatorReward is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesOperatorReward struct {
	Operator common.Address
	Amount   *big.Int
}

// IRewardsCoordinatorTypesRewardsMerkleClaim is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesRewardsMerkleClaim struct {
	RootIndex       uint32
	EarnerIndex     uint32
	EarnerTreeProof []byte
	EarnerLeaf      IRewardsCoordinatorTypesEarnerTreeMerkleLeaf
	TokenIndices    []uint32
	TokenTreeProofs [][]byte
	TokenLeaves     []IRewardsCoordinatorTypesTokenTreeMerkleLeaf
}

// IRewardsCoordinatorTypesRewardsSubmission is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesRewardsSubmission struct {
	StrategiesAndMultipliers []IRewardsCoordinatorTypesStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IRewardsCoordinatorTypesStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// IRewardsCoordinatorTypesTokenTreeMerkleLeaf is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesTokenTreeMerkleLeaf struct {
	Token              common.Address
	CumulativeEarnings *big.Int
}

// RewardsCoordinatorMetaData contains all meta data concerning the RewardsCoordinator contract.
var RewardsCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_CALCULATION_INTERVAL_SECONDS\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_MAX_REWARDS_DURATION\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_MAX_RETROACTIVE_LENGTH\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_MAX_FUTURE_LENGTH\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_GENESIS_REWARDS_TIMESTAMP\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CALCULATION_INTERVAL_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GENESIS_REWARDS_TIMESTAMP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FUTURE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_RETROACTIVE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_REWARDS_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activationDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateEarnerLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"calculateTokenLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"checkClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimerFor\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAVSRewardsSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorDirectedAVSRewardsSubmission\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorDirectedRewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRewardsForAllEarners\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRewardsForAllSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cumulativeClaimed\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"totalClaimed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currRewardsCalculationEndTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultOperatorSplitBips\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableRoot\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentClaimableDistributionRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentDistributionRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributionRootAtIndex\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributionRootsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorAVSSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorPISplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRootIndexFromHash\",\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_rewardsUpdater\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_defaultSplitBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAVSRewardsSubmissionHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorDirectedAVSRewardsSubmissionHash\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRewardsForAllSubmitter\",\"inputs\":[{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRewardsSubmissionForAllEarnersHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRewardsSubmissionForAllHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"processClaims\",\"inputs\":[{\"name\":\"claims\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsMerkleClaim[]\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardsUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setActivationDelay\",\"inputs\":[{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimerFor\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimerFor\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultOperatorSplit\",\"inputs\":[{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorAVSSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorPISplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsForAllSubmitter\",\"inputs\":[{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_newValue\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsUpdater\",\"inputs\":[{\"name\":\"_rewardsUpdater\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submissionNonce\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSRewardsSubmissionCreated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ActivationDelaySet\",\"inputs\":[{\"name\":\"oldActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimerForSet\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldClaimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultOperatorSplitBipsSet\",\"inputs\":[{\"name\":\"oldDefaultOperatorSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newDefaultOperatorSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootDisabled\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootSubmitted\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSSplitBipsSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"oldOperatorAVSSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newOperatorAVSSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDirectedAVSRewardsSubmissionCreated\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorDirectedRewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"operatorDirectedRewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorPISplitBipsSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"oldOperatorPISplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newOperatorPISplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsClaimed\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"claimedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsForAllSubmitterSet\",\"inputs\":[{\"name\":\"rewardsForAllSubmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"newValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsSubmissionForAllCreated\",\"inputs\":[{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsSubmissionForAllEarnersCreated\",\"inputs\":[{\"name\":\"tokenHopper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsUpdaterSet\",\"inputs\":[{\"name\":\"oldRewardsUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newRewardsUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AmountExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AmountIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DurationExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EarningsNotGreaterThanClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCalculationIntervalSecondsRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidClaimProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDurationRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEarner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEarnerLeafIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGenesisRewardsTimestampRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRootIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidStartTimestampRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTokenLeafIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewRootMustBeForNewCalculatedPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorsNotInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PreviousSplitPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardsEndTimestampNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootAlreadyActivated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootNotActivated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SplitExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartTimestampTooFarInFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartTimestampTooFarInPast\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategiesNotInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SubmissionNotRetroactive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[]}]",
	Bin: "0x6101c0604052348015610010575f5ffd5b506040516142e63803806142e683398101604081905261002f91610211565b858a8a8a88888888888f6001600160a01b038116610060576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b031660805261007685826102d3565b63ffffffff161561009a57604051630e06bd3160e01b815260040160405180910390fd5b6100a762015180866102d3565b63ffffffff16156100cb5760405163223c7b3960e11b815260040160405180910390fd5b6001600160a01b0397881660a05295871660c05293861660e05263ffffffff9283166101005290821661012052811661014052908116610160521661018052166101a052610117610126565b50505050505050505050610306565b5f54610100900460ff16156101915760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146101e0575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101f6575f5ffd5b50565b805163ffffffff8116811461020c575f5ffd5b919050565b5f5f5f5f5f5f5f5f5f5f6101408b8d03121561022b575f5ffd5b8a51610236816101e2565b60208c0151909a50610247816101e2565b60408c0151909950610258816101e2565b60608c0151909850610269816101e2565b60808c015190975061027a816101e2565b955061028860a08c016101f9565b945061029660c08c016101f9565b93506102a460e08c016101f9565b92506102b36101008c016101f9565b91506102c26101208c016101f9565b90509295989b9194979a5092959850565b5f63ffffffff8316806102f457634e487b7160e01b5f52601260045260245ffd5b8063ffffffff84160691505092915050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a051613f166103d05f395f81816105a60152612a1701525f818161045a0152612f3f01525f81816103bc015261238001525f81816105080152612efd01525f818161080c0152612e0d01525f818161076001528181612e5d0152612eab01525f81816108600152611a7401525f818161052f0152612fda01525f81816108d301526119e401525f81816106f201528181610b62015281816111750152611e190152613f165ff3fe608060405234801561000f575f5ffd5b506004361061037c575f3560e01c8063886f1195116101d4578063dcbb03b311610109578063f2fde38b116100a9578063fabc1cbc11610079578063fabc1cbc14610981578063fbf1e2c114610994578063fce36c7d146109a7578063ff9f6cce146109ba575f5ffd5b8063f2fde38b14610935578063f6efbb5914610948578063f8cd84481461095b578063f96abf2e1461096e575f5ffd5b8063e810ce21116100e4578063e810ce21146108bb578063ea4d3c9b146108ce578063ed71e6a2146108f5578063f22cef8514610922575f5ffd5b8063dcbb03b314610882578063de02e50314610895578063e063f81f146108a8575f5ffd5b8063a50a1d9c11610174578063bb7e451f1161014f578063bb7e451f146107e8578063bf21a8aa14610807578063c46db6061461082e578063ca8aa7c71461085b575f5ffd5b8063a50a1d9c14610795578063aebd8bae146107a8578063b3dbb0e0146107d5575f5ffd5b80639be3d4e4116101af5780639be3d4e4146107405780639cb9a5fa146107485780639d45c2811461075b578063a0169ddd14610782575f5ffd5b8063886f1195146106ed5780638da5cb5b146107145780639104c31914610725575f5ffd5b80634596021c116102b55780635c975abb11610255578063715018a611610225578063715018a6146106a05780637b8f8b05146106a8578063863cb9a9146106b0578063865c6953146106c3575f5ffd5b80635c975abb146106435780635e9d83481461064b57806363f6a7981461065e5780636d21117e14610673575f5ffd5b80634d18cc35116102905780634d18cc35146105ee57806358baaa3e14610605578063595c6a67146106185780635ac86ab714610620575f5ffd5b80634596021c1461058e5780634657e26a146105a15780634b943960146105c8575f5ffd5b80632b9f64a41161032057806339b70e38116102fb57806339b70e381461052a5780633a8c0786146105515780633ccc861d146105685780633efe1db61461057b575f5ffd5b80632b9f64a4146104b057806336af41fa146104f057806337838ed014610503575f5ffd5b80630eb383451161035b5780630eb3834514610440578063131433b414610455578063136439dd1461047c578063149bc8721461048f575f5ffd5b806218572c1461038057806304a0c502146103b75780630e9a53cf146103f3575b5f5ffd5b6103a261038e3660046134f6565b60d16020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b6103de7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016103ae565b6103fb6109cd565b6040516103ae91905f6080820190508251825263ffffffff602084015116602083015263ffffffff604084015116604083015260608301511515606083015292915050565b61045361044e36600461351e565b610acd565b005b6103de7f000000000000000000000000000000000000000000000000000000000000000081565b61045361048a366004613555565b610b4d565b6104a261049d366004613582565b610c22565b6040519081526020016103ae565b6104d86104be3660046134f6565b60cc6020525f90815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020016103ae565b6104536104fe3660046135e4565b610c97565b6103de7f000000000000000000000000000000000000000000000000000000000000000081565b6104d87f000000000000000000000000000000000000000000000000000000000000000081565b60cb546103de90600160a01b900463ffffffff1681565b610453610576366004613634565b610e37565b61045361058936600461368b565b610e7c565b61045361059c3660046136b5565b611070565b6104d87f000000000000000000000000000000000000000000000000000000000000000081565b6105db6105d63660046134f6565b6110f1565b60405161ffff90911681526020016103ae565b60cb546103de90600160c01b900463ffffffff1681565b610453610613366004613708565b61114c565b610453611160565b6103a261062e366004613721565b606654600160ff9092169190911b9081161490565b6066546104a2565b6103a2610659366004613741565b61120f565b60cb546105db90600160e01b900461ffff1681565b6103a2610681366004613773565b60cf60209081525f928352604080842090915290825290205460ff1681565b61045361129a565b60ca546104a2565b6104536106be3660046134f6565b6112ab565b6104a26106d136600461379d565b60cd60209081525f928352604080842090915290825290205481565b6104d87f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b03166104d8565b6104d873beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b6103fb6112bc565b6104536107563660046137c9565b611358565b6103de7f000000000000000000000000000000000000000000000000000000000000000081565b6104536107903660046134f6565b611503565b6104536107a336600461382b565b61150e565b6103a26107b6366004613773565b60d260209081525f928352604080842090915290825290205460ff1681565b6104536107e3366004613844565b61151f565b6104a26107f63660046134f6565b60ce6020525f908152604090205481565b6103de7f000000000000000000000000000000000000000000000000000000000000000081565b6103a261083c366004613773565b60d060209081525f928352604080842090915290825290205460ff1681565b6104d87f000000000000000000000000000000000000000000000000000000000000000081565b61045361089036600461386e565b61168f565b6103fb6108a3366004613555565b611821565b6105db6108b636600461379d565b6118b1565b6103de6108c9366004613555565b61191d565b6104d87f000000000000000000000000000000000000000000000000000000000000000081565b6103a2610903366004613773565b60d360209081525f928352604080842090915290825290205460ff1681565b61045361093036600461379d565b61199e565b6104536109433660046134f6565b611b08565b6104536109563660046138b2565b611b83565b6104a2610969366004613582565b611cb8565b61045361097c366004613708565b611cc8565b61045361098f366004613555565b611e17565b60cb546104d8906001600160a01b031681565b6104536109b53660046135e4565b611f2d565b6104536109c83660046135e4565b61207c565b604080516080810182525f80825260208201819052918101829052606081019190915260ca545b8015610aa5575f60ca610a08600184613924565b81548110610a1857610a18613937565b5f91825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff161580156060830181905291925090610a875750806040015163ffffffff164210155b15610a925792915050565b5080610a9d8161394b565b9150506109f4565b5050604080516080810182525f80825260208201819052918101829052606081019190915290565b610ad56121fb565b6001600160a01b0382165f81815260d1602052604080822054905160ff9091169284151592841515927f4de6293e668df1398422e1def12118052c1539a03cbfedc145895d48d7685f1c9190a4506001600160a01b03919091165f90815260d160205260409020805460ff1916911515919091179055565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610baf573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bd39190613960565b610bf057604051631d77d47760e21b815260040160405180910390fd5b6066548181168114610c155760405163c61dca5d60e01b815260040160405180910390fd5b610c1e82612255565b5050565b5f80610c3160208401846134f6565b8360200135604051602001610c7a9392919060f89390931b6001600160f81b031916835260609190911b6bffffffffffffffffffffffff19166001830152601582015260350190565b604051602081830303815290604052805190602001209050919050565b606654600190600290811603610cc05760405163840a48d560e01b815260040160405180910390fd5b335f90815260d1602052604090205460ff16610cef57604051635c427cd960e01b815260040160405180910390fd5b610cf7612292565b5f5b82811015610e275736848483818110610d1457610d14613937565b9050602002810190610d26919061397b565b335f81815260ce60209081526040808320549051949550939192610d509290918591879101613ac9565b604051602081830303815290604052805190602001209050610d71836122eb565b335f90815260d0602090815260408083208484529091529020805460ff19166001908117909155610da3908390613af8565b335f81815260ce602052604090819020929092559051829184917f51088b8c89628df3a8174002c2a034d0152fce6af8415d651b2a4734bf27048290610dea908890613b0b565b60405180910390a4610e1c333060408601803590610e0b90602089016134f6565b6001600160a01b03169291906123db565b505050600101610cf9565b50610e326001609755565b505050565b606654600290600490811603610e605760405163840a48d560e01b815260040160405180910390fd5b610e68612292565b610e728383612446565b610e326001609755565b606654600390600890811603610ea55760405163840a48d560e01b815260040160405180910390fd5b60cb546001600160a01b03163314610ed057604051635c427cd960e01b815260040160405180910390fd5b60cb5463ffffffff600160c01b909104811690831611610f0357604051631ca7e50b60e21b815260040160405180910390fd5b428263ffffffff1610610f29576040516306957c9160e11b815260040160405180910390fd5b60ca5460cb545f90610f4890600160a01b900463ffffffff1642613b1d565b6040805160808101825287815263ffffffff87811660208084018281528684168587018181525f6060880181815260ca8054600181018255925297517f42d72674974f694b5f5159593243114d38a5c39c89d6b62fee061ff523240ee160029092029182015592517f42d72674974f694b5f5159593243114d38a5c39c89d6b62fee061ff523240ee290930180549151975193871667ffffffffffffffff1990921691909117600160201b978716979097029690961760ff60401b1916600160401b921515929092029190911790945560cb805463ffffffff60c01b1916600160c01b840217905593519283529394508892908616917fecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08910160405180910390a45050505050565b6066546002906004908116036110995760405163840a48d560e01b815260040160405180910390fd5b6110a1612292565b5f5b838110156110e0576110d88585838181106110c0576110c0613937565b90506020028101906110d29190613b39565b84612446565b6001016110a3565b506110eb6001609755565b50505050565b6001600160a01b0381165f90815260d5602090815260408083208151606081018352905461ffff80821683526201000082041693820193909352600160201b90920463ffffffff1690820152611146906126ce565b92915050565b6111546121fb565b61115d81612719565b50565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156111c2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111e69190613960565b61120357604051631d77d47760e21b815260040160405180910390fd5b61120d5f19612255565b565b5f6112928260ca6112236020830183613708565b63ffffffff168154811061123957611239613937565b5f91825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff161515606082015261278a565b506001919050565b6112a26121fb565b61120d5f61292d565b6112b36121fb565b61115d8161297e565b604080516080810182525f80825260208201819052918101829052606081019190915260ca80546112ef90600190613924565b815481106112ff576112ff613937565b5f91825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff1615156060820152919050565b6066546005906020908116036113815760405163840a48d560e01b815260040160405180910390fd5b8361138b816129d9565b6113a85760405163932d94f760e01b815260040160405180910390fd5b6113b0612292565b5f5b838110156114f157368585838181106113cd576113cd613937565b90506020028101906113df9190613b4d565b6001600160a01b0388165f90815260ce6020908152604080832054905193945092611410918b918591879101613cbb565b6040516020818303038152906040528051906020012090505f61143284612a83565b6001600160a01b038b165f90815260d3602090815260408083208684529091529020805460ff19166001908117909155909150611470908490613af8565b6001600160a01b038b165f81815260ce60205260409081902092909255905183919033907ffc8888bffd711da60bc5092b33f677d81896fe80ecc677b84cfab8184462b6e0906114c39088908a90613ce1565b60405180910390a46114e1333083610e0b6040890160208a016134f6565b5050600190920191506113b29050565b506114fc6001609755565b5050505050565b33610c1e8183612c69565b6115166121fb565b61115d81612ccc565b6066546007906080908116036115485760405163840a48d560e01b815260040160405180910390fd5b82611552816129d9565b61156f5760405163932d94f760e01b815260040160405180910390fd5b61271061ffff841611156115965760405163891c63df60e01b815260040160405180910390fd5b60cb545f906115b290600160a01b900463ffffffff1642613b1d565b6001600160a01b0386165f90815260d5602090815260408083208151606081018352905461ffff80821683526201000082041693820193909352600160201b90920463ffffffff16908201529192509061160b906126ce565b6001600160a01b0387165f90815260d560205260409020909150611630908684612d37565b6040805163ffffffff8416815261ffff838116602083015287168183015290516001600160a01b0388169133917fd1e028bd664486a46ad26040e999cd2d22e1e9a094ee6afe19fcf64678f16f749181900360600190a3505050505050565b6066546006906040908116036116b85760405163840a48d560e01b815260040160405180910390fd5b836116c2816129d9565b6116df5760405163932d94f760e01b815260040160405180910390fd5b61271061ffff841611156117065760405163891c63df60e01b815260040160405180910390fd5b60cb545f9061172290600160a01b900463ffffffff1642613b1d565b6001600160a01b038781165f90815260d460209081526040808320938a1683529281528282208351606081018552905461ffff80821683526201000082041692820192909252600160201b90910463ffffffff1692810192909252919250611789906126ce565b6001600160a01b038089165f90815260d460209081526040808320938b168352929052209091506117bb908684612d37565b6040805163ffffffff8416815261ffff838116602083015287168183015290516001600160a01b0388811692908a169133917f48e198b6ae357e529204ee53a8e514c470ff77d9cc8e4f7207f8b5d490ae6934919081900360600190a450505050505050565b604080516080810182525f80825260208201819052918101829052606081019190915260ca828154811061185757611857613937565b5f91825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff161515606082015292915050565b6001600160a01b038281165f90815260d46020908152604080832093851683529281528282208351606081018552905461ffff80821683526201000082041692820192909252600160201b90910463ffffffff169281019290925290611916906126ce565b9392505050565b60ca545f905b63ffffffff811615611984578260ca61193d600184613cf9565b63ffffffff168154811061195357611953613937565b905f5260205f2090600202015f01540361197257611916600182613cf9565b8061197c81613d15565b915050611923565b5060405163504570e360e01b815260040160405180910390fd5b816119a8816129d9565b6119c55760405163932d94f760e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0384811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015611a29573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a4d9190613960565b80611ae1575060405163ba1a84e560e01b81526001600160a01b0384811660048301525f917f00000000000000000000000000000000000000000000000000000000000000009091169063ba1a84e590602401602060405180830381865afa158015611abb573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611adf9190613d33565b115b611afe5760405163fb494ea160e01b815260040160405180910390fd5b610e328383612c69565b611b106121fb565b6001600160a01b038116611b7a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b61115d8161292d565b5f54610100900460ff1615808015611ba157505f54600160ff909116105b80611bba5750303b158015611bba57505f5460ff166001145b611c1d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401611b71565b5f805460ff191660011790558015611c3e575f805461ff0019166101001790555b611c4785612255565b611c508661292d565b611c598461297e565b611c6283612719565b611c6b82612ccc565b8015611cb0575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b5f6001610c3160208401846134f6565b606654600390600890811603611cf15760405163840a48d560e01b815260040160405180910390fd5b60cb546001600160a01b03163314611d1c57604051635c427cd960e01b815260040160405180910390fd5b60ca5463ffffffff831610611d44576040516394a8d38960e01b815260040160405180910390fd5b5f60ca8363ffffffff1681548110611d5e57611d5e613937565b905f5260205f20906002020190508060010160089054906101000a900460ff1615611d9c57604051631b14174b60e01b815260040160405180910390fd5b6001810154600160201b900463ffffffff164210611dcd57604051630c36f66560e21b815260040160405180910390fd5b60018101805460ff60401b1916600160401b17905560405163ffffffff8416907fd850e6e5dfa497b72661fa73df2923464eaed9dc2ff1d3cb82bccbfeabe5c41e905f90a2505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611e73573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e979190613d4a565b6001600160a01b0316336001600160a01b031614611ec85760405163794821ff60e01b815260040160405180910390fd5b60665480198219811614611eef5760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b6066545f90600190811603611f555760405163840a48d560e01b815260040160405180910390fd5b611f5d612292565b5f5b82811015610e275736848483818110611f7a57611f7a613937565b9050602002810190611f8c919061397b565b335f81815260ce60209081526040808320549051949550939192611fb69290918591879101613ac9565b604051602081830303815290604052805190602001209050611fd7836122eb565b335f90815260cf602090815260408083208484529091529020805460ff19166001908117909155612009908390613af8565b335f81815260ce602052604090819020929092559051829184917f450a367a380c4e339e5ae7340c8464ef27af7781ad9945cfe8abd828f89e628190612050908890613b0b565b60405180910390a4612071333060408601803590610e0b90602089016134f6565b505050600101611f5f565b6066546004906010908116036120a55760405163840a48d560e01b815260040160405180910390fd5b335f90815260d1602052604090205460ff166120d457604051635c427cd960e01b815260040160405180910390fd5b6120dc612292565b5f5b82811015610e2757368484838181106120f9576120f9613937565b905060200281019061210b919061397b565b335f81815260ce602090815260408083205490519495509391926121359290918591879101613ac9565b604051602081830303815290604052805190602001209050612156836122eb565b335f90815260d2602090815260408083208484529091529020805460ff19166001908117909155612188908390613af8565b335f81815260ce602052604090819020929092559051829184917f5251b6fdefcb5d81144e735f69ea4c695fd43b0289ca53dc075033f5fc80068b906121cf908890613b0b565b60405180910390a46121f0333060408601803590610e0b90602089016134f6565b5050506001016120de565b6033546001600160a01b0316331461120d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401611b71565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b6002609754036122e45760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611b71565b6002609755565b61231d6122f88280613d65565b6123086080850160608601613708565b61231860a0860160808701613708565b612ded565b5f816040013511612341576040516310eb483f60e21b815260040160405180910390fd5b6f4b3b4ca85a86c47a098a223fffffffff816040013511156123765760405163070b5a6f60e21b815260040160405180910390fd5b6123a663ffffffff7f00000000000000000000000000000000000000000000000000000000000000001642613af8565b6123b66080830160608401613708565b63ffffffff16111561115d57604051637ee2b44360e01b815260040160405180910390fd5b6040516001600160a01b03808516602483015283166044820152606481018290526110eb9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526130c5565b5f60ca6124566020850185613708565b63ffffffff168154811061246c5761246c613937565b5f91825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff161515606082015290506124cc838261278a565b5f6124dd60808501606086016134f6565b6001600160a01b038082165f90815260cc602052604090205491925016806125025750805b336001600160a01b0382161461252b57604051635c427cd960e01b815260040160405180910390fd5b5f5b61253a60a0870187613dab565b9050811015611cb0573661255160e0880188613d65565b8381811061256157612561613937565b6001600160a01b0387165f90815260cd602090815260408083209302949094019450929091508290612595908501856134f6565b6001600160a01b03166001600160a01b031681526020019081526020015f20549050808260200135116125db5760405163aa385e8160e01b815260040160405180910390fd5b5f6125ea826020850135613924565b6001600160a01b0387165f90815260cd6020908152604082209293508501803592919061261790876134f6565b6001600160a01b031681526020808201929092526040015f20919091556126589089908390612648908701876134f6565b6001600160a01b03169190613198565b86516001600160a01b03808a1691878216918916907f9543dbd55580842586a951f0386e24d68a5df99ae29e3b216588b45fd684ce319061269c60208901896134f6565b604080519283526001600160a01b039091166020830152810186905260600160405180910390a450505060010161252d565b5f816040015163ffffffff165f036126f357505060cb54600160e01b900461ffff1690565b816040015163ffffffff1642101561270c578151611146565b506020015190565b919050565b60cb546040805163ffffffff600160a01b9093048316815291831660208301527faf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3910160405180910390a160cb805463ffffffff909216600160a01b0263ffffffff60a01b19909216919091179055565b8060600151156127ad57604051631b14174b60e01b815260040160405180910390fd5b806040015163ffffffff164210156127d857604051631437a2bb60e31b815260040160405180910390fd5b6127e560c0830183613dab565b90506127f460a0840184613dab565b905014612814576040516343714afd60e01b815260040160405180910390fd5b61282160e0830183613d65565b905061283060c0840184613dab565b905014612850576040516343714afd60e01b815260040160405180910390fd5b805161287c906128666040850160208601613708565b6128736040860186613df1565b866060016131c8565b5f5b61288b60a0840184613dab565b9050811015610e325761292560808401356128a960a0860186613dab565b848181106128b9576128b9613937565b90506020020160208101906128ce9190613708565b6128db60c0870187613dab565b858181106128eb576128eb613937565b90506020028101906128fd9190613df1565b61290a60e0890189613d65565b8781811061291a5761291a613937565b90506040020161326c565b60010161287e565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60cb546040516001600160a01b038084169216907f237b82f438d75fc568ebab484b75b01d9287b9e98b490b7c23221623b6705dbb905f90a360cb80546001600160a01b0319166001600160a01b0392909216919091179055565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af1158015612a5f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111469190613960565b5f612ab1612a918380613d65565b612aa16080860160608701613708565b61231860a0870160808801613708565b5f612abf6040840184613d65565b905011612adf5760405163796cc52560e01b815260040160405180910390fd5b42612af060a0840160808501613708565b612b006080850160608601613708565b612b0a9190613b1d565b63ffffffff1610612b2e5760405163150358a160e21b815260040160405180910390fd5b5f80805b612b3f6040860186613d65565b9050811015612c305736612b566040870187613d65565b83818110612b6657612b66613937565b6040029190910191505f9050612b7f60208301836134f6565b6001600160a01b031603612ba657604051630863a45360e11b815260040160405180910390fd5b612bb360208201826134f6565b6001600160a01b0316836001600160a01b031610612be4576040516310fb47f160e31b815260040160405180910390fd5b5f816020013511612c08576040516310eb483f60e21b815260040160405180910390fd5b612c1560208201826134f6565b9250612c25602082013585613af8565b935050600101612b32565b506f4b3b4ca85a86c47a098a223fffffffff821115612c625760405163070b5a6f60e21b815260040160405180910390fd5b5092915050565b6001600160a01b038083165f81815260cc602052604080822080548686166001600160a01b0319821681179092559151919094169392849290917fbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca3129190a4505050565b60cb546040805161ffff600160e01b9093048316815291831660208301527fe6cd4edfdcc1f6d130ab35f73d72378f3a642944fb4ee5bd84b7807a81ea1c4e910160405180910390a160cb805461ffff909216600160e01b0261ffff60e01b19909216919091179055565b8254600160201b900463ffffffff164211612d6557604051637b1e25c560e01b815260040160405180910390fd5b8254600160201b900463ffffffff165f03612d9a5760cb548354600160e01b90910461ffff1661ffff19909116178355612db1565b825462010000810461ffff1661ffff199091161783555b825463ffffffff909116600160201b0267ffffffff000000001961ffff90931662010000029290921667ffffffffffff00001990911617179055565b82612e0b5760405163796cc52560e01b815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168163ffffffff161115612e5857604051630dd0b9f560e21b815260040160405180910390fd5b612e827f000000000000000000000000000000000000000000000000000000000000000082613e48565b63ffffffff1615612ea65760405163ee66470560e01b815260040160405180910390fd5b612ed07f000000000000000000000000000000000000000000000000000000000000000083613e48565b63ffffffff1615612ef457604051633c1a94f160e21b815260040160405180910390fd5b8163ffffffff167f000000000000000000000000000000000000000000000000000000000000000063ffffffff1642612f2d9190613924565b11158015612f6757508163ffffffff167f000000000000000000000000000000000000000000000000000000000000000063ffffffff1611155b612f845760405163041aa75760e11b815260040160405180910390fd5b5f805b84811015611cb0575f868683818110612fa257612fa2613937565b612fb892602060409092020190810191506134f6565b60405163198f077960e21b81526001600160a01b0380831660048301529192507f00000000000000000000000000000000000000000000000000000000000000009091169063663c1de490602401602060405180830381865afa158015613021573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130459190613960565b8061306c57506001600160a01b03811673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0145b61308957604051632efd965160e11b815260040160405180910390fd5b806001600160a01b0316836001600160a01b0316106130bb5760405163dfad9ca160e01b815260040160405180910390fd5b9150600101612f87565b5f613119826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166132aa9092919063ffffffff16565b905080515f14806131395750808060200190518101906131399190613960565b610e325760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401611b71565b6040516001600160a01b038316602482015260448101829052610e3290849063a9059cbb60e01b9060640161240f565b6131d3602083613e6f565b6001901b8463ffffffff16106131fb5760405162c6c39d60e71b815260040160405180910390fd5b5f61320582610c22565b905061324f84848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508a92508591505063ffffffff89166132c0565b611cb0576040516369ca16c960e01b815260040160405180910390fd5b613277602083613e6f565b6001901b8463ffffffff16106132a05760405163054ff4df60e51b815260040160405180910390fd5b5f61320582611cb8565b60606132b884845f856132d7565b949350505050565b5f836132cd8685856133ae565b1495945050505050565b6060824710156133385760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401611b71565b5f5f866001600160a01b031685876040516133539190613e82565b5f6040518083038185875af1925050503d805f811461338d576040519150601f19603f3d011682016040523d82523d5f602084013e613392565b606091505b50915091506133a387838387613445565b979650505050505050565b5f602084516133bd9190613e98565b156133db576040516313717da960e21b815260040160405180910390fd5b8260205b8551811161343c576133f2600285613e98565b5f0361341357815f528086015160205260405f20915060028404935061342a565b808601515f528160205260405f2091506002840493505b613435602082613af8565b90506133df565b50949350505050565b606083156134b35782515f036134ac576001600160a01b0385163b6134ac5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611b71565b50816132b8565b6132b883838151156134c85781518083602001fd5b8060405162461bcd60e51b8152600401611b719190613eab565b6001600160a01b038116811461115d575f5ffd5b5f60208284031215613506575f5ffd5b8135611916816134e2565b801515811461115d575f5ffd5b5f5f6040838503121561352f575f5ffd5b823561353a816134e2565b9150602083013561354a81613511565b809150509250929050565b5f60208284031215613565575f5ffd5b5035919050565b5f6040828403121561357c575f5ffd5b50919050565b5f60408284031215613592575f5ffd5b611916838361356c565b5f5f83601f8401126135ac575f5ffd5b50813567ffffffffffffffff8111156135c3575f5ffd5b6020830191508360208260051b85010111156135dd575f5ffd5b9250929050565b5f5f602083850312156135f5575f5ffd5b823567ffffffffffffffff81111561360b575f5ffd5b6136178582860161359c565b90969095509350505050565b5f610100828403121561357c575f5ffd5b5f5f60408385031215613645575f5ffd5b823567ffffffffffffffff81111561365b575f5ffd5b61366785828601613623565b925050602083013561354a816134e2565b803563ffffffff81168114612714575f5ffd5b5f5f6040838503121561369c575f5ffd5b823591506136ac60208401613678565b90509250929050565b5f5f5f604084860312156136c7575f5ffd5b833567ffffffffffffffff8111156136dd575f5ffd5b6136e98682870161359c565b90945092505060208401356136fd816134e2565b809150509250925092565b5f60208284031215613718575f5ffd5b61191682613678565b5f60208284031215613731575f5ffd5b813560ff81168114611916575f5ffd5b5f60208284031215613751575f5ffd5b813567ffffffffffffffff811115613767575f5ffd5b6132b884828501613623565b5f5f60408385031215613784575f5ffd5b823561378f816134e2565b946020939093013593505050565b5f5f604083850312156137ae575f5ffd5b82356137b9816134e2565b9150602083013561354a816134e2565b5f5f5f604084860312156137db575f5ffd5b83356137e6816134e2565b9250602084013567ffffffffffffffff811115613801575f5ffd5b61380d8682870161359c565b9497909650939450505050565b803561ffff81168114612714575f5ffd5b5f6020828403121561383b575f5ffd5b6119168261381a565b5f5f60408385031215613855575f5ffd5b8235613860816134e2565b91506136ac6020840161381a565b5f5f5f60608486031215613880575f5ffd5b833561388b816134e2565b9250602084013561389b816134e2565b91506138a96040850161381a565b90509250925092565b5f5f5f5f5f60a086880312156138c6575f5ffd5b85356138d1816134e2565b94506020860135935060408601356138e8816134e2565b92506138f660608701613678565b91506139046080870161381a565b90509295509295909350565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561114657611146613910565b634e487b7160e01b5f52603260045260245ffd5b5f8161395957613959613910565b505f190190565b5f60208284031215613970575f5ffd5b815161191681613511565b5f8235609e1983360301811261398f575f5ffd5b9190910192915050565b5f5f8335601e198436030181126139ae575f5ffd5b830160208101925035905067ffffffffffffffff8111156139cd575f5ffd5b8060061b36038213156135dd575f5ffd5b8183526020830192505f815f5b84811015613a415781356139fe816134e2565b6001600160a01b0316865260208201356bffffffffffffffffffffffff8116808214613a28575f5ffd5b60208801525060409586019591909101906001016139eb565b5093949350505050565b5f613a568283613999565b60a08552613a6860a0860182846139de565b9150506020830135613a79816134e2565b6001600160a01b031660208501526040838101359085015263ffffffff613aa260608501613678565b16606085015263ffffffff613ab960808501613678565b1660808501528091505092915050565b60018060a01b0384168152826020820152606060408201525f613aef6060830184613a4b565b95945050505050565b8082018082111561114657611146613910565b602081525f6119166020830184613a4b565b63ffffffff818116838216019081111561114657611146613910565b5f823560fe1983360301811261398f575f5ffd5b5f823560be1983360301811261398f575f5ffd5b5f5f8335601e19843603018112613b76575f5ffd5b830160208101925035905067ffffffffffffffff811115613b95575f5ffd5b8036038213156135dd575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f613bd68283613999565b60c08552613be860c0860182846139de565b9150506020830135613bf9816134e2565b6001600160a01b03166020850152613c146040840184613999565b858303604087015280835290915f91906020015b81831015613c63578335613c3b816134e2565b6001600160a01b03168152602084810135908201526040938401936001939093019201613c28565b613c6f60608701613678565b63ffffffff811660608901529350613c8960808701613678565b63ffffffff811660808901529350613ca460a0870187613b61565b9450925086810360a08801526133a3818585613ba3565b60018060a01b0384168152826020820152606060408201525f613aef6060830184613bcb565b828152604060208201525f6132b86040830184613bcb565b63ffffffff828116828216039081111561114657611146613910565b5f63ffffffff821680613d2a57613d2a613910565b5f190192915050565b5f60208284031215613d43575f5ffd5b5051919050565b5f60208284031215613d5a575f5ffd5b8151611916816134e2565b5f5f8335601e19843603018112613d7a575f5ffd5b83018035915067ffffffffffffffff821115613d94575f5ffd5b6020019150600681901b36038213156135dd575f5ffd5b5f5f8335601e19843603018112613dc0575f5ffd5b83018035915067ffffffffffffffff821115613dda575f5ffd5b6020019150600581901b36038213156135dd575f5ffd5b5f5f8335601e19843603018112613e06575f5ffd5b83018035915067ffffffffffffffff821115613e20575f5ffd5b6020019150368190038213156135dd575f5ffd5b634e487b7160e01b5f52601260045260245ffd5b5f63ffffffff831680613e5d57613e5d613e34565b8063ffffffff84160691505092915050565b5f82613e7d57613e7d613e34565b500490565b5f82518060208501845e5f920191825250919050565b5f82613ea657613ea6613e34565b500690565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f8301168401019150509291505056fea26469706673582212209ae8e3cc13bb5aa90cb31ce2f6b06a545376d52a7949d5bd3cda2c01b7b783b764736f6c634300081b0033",
}

// RewardsCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardsCoordinatorMetaData.ABI instead.
var RewardsCoordinatorABI = RewardsCoordinatorMetaData.ABI

// RewardsCoordinatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RewardsCoordinatorMetaData.Bin instead.
var RewardsCoordinatorBin = RewardsCoordinatorMetaData.Bin

// DeployRewardsCoordinator deploys a new Ethereum contract, binding an instance of RewardsCoordinator to it.
func DeployRewardsCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, _delegationManager common.Address, _strategyManager common.Address, _allocationManager common.Address, _pauserRegistry common.Address, _permissionController common.Address, _CALCULATION_INTERVAL_SECONDS uint32, _MAX_REWARDS_DURATION uint32, _MAX_RETROACTIVE_LENGTH uint32, _MAX_FUTURE_LENGTH uint32, _GENESIS_REWARDS_TIMESTAMP uint32) (common.Address, *types.Transaction, *RewardsCoordinator, error) {
	parsed, err := RewardsCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RewardsCoordinatorBin), backend, _delegationManager, _strategyManager, _allocationManager, _pauserRegistry, _permissionController, _CALCULATION_INTERVAL_SECONDS, _MAX_REWARDS_DURATION, _MAX_RETROACTIVE_LENGTH, _MAX_FUTURE_LENGTH, _GENESIS_REWARDS_TIMESTAMP)
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

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorSession) AllocationManager() (common.Address, error) {
	return _RewardsCoordinator.Contract.AllocationManager(&_RewardsCoordinator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) AllocationManager() (common.Address, error) {
	return _RewardsCoordinator.Contract.AllocationManager(&_RewardsCoordinator.CallOpts)
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
func (_RewardsCoordinator *RewardsCoordinatorCaller) CalculateEarnerLeafHash(opts *bind.CallOpts, leaf IRewardsCoordinatorTypesEarnerTreeMerkleLeaf) ([32]byte, error) {
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
func (_RewardsCoordinator *RewardsCoordinatorSession) CalculateEarnerLeafHash(leaf IRewardsCoordinatorTypesEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _RewardsCoordinator.Contract.CalculateEarnerLeafHash(&_RewardsCoordinator.CallOpts, leaf)
}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) CalculateEarnerLeafHash(leaf IRewardsCoordinatorTypesEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _RewardsCoordinator.Contract.CalculateEarnerLeafHash(&_RewardsCoordinator.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_RewardsCoordinator *RewardsCoordinatorCaller) CalculateTokenLeafHash(opts *bind.CallOpts, leaf IRewardsCoordinatorTypesTokenTreeMerkleLeaf) ([32]byte, error) {
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
func (_RewardsCoordinator *RewardsCoordinatorSession) CalculateTokenLeafHash(leaf IRewardsCoordinatorTypesTokenTreeMerkleLeaf) ([32]byte, error) {
	return _RewardsCoordinator.Contract.CalculateTokenLeafHash(&_RewardsCoordinator.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) CalculateTokenLeafHash(leaf IRewardsCoordinatorTypesTokenTreeMerkleLeaf) ([32]byte, error) {
	return _RewardsCoordinator.Contract.CalculateTokenLeafHash(&_RewardsCoordinator.CallOpts, leaf)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCaller) CheckClaim(opts *bind.CallOpts, claim IRewardsCoordinatorTypesRewardsMerkleClaim) (bool, error) {
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
func (_RewardsCoordinator *RewardsCoordinatorSession) CheckClaim(claim IRewardsCoordinatorTypesRewardsMerkleClaim) (bool, error) {
	return _RewardsCoordinator.Contract.CheckClaim(&_RewardsCoordinator.CallOpts, claim)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) CheckClaim(claim IRewardsCoordinatorTypesRewardsMerkleClaim) (bool, error) {
	return _RewardsCoordinator.Contract.CheckClaim(&_RewardsCoordinator.CallOpts, claim)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address earner) view returns(address claimer)
func (_RewardsCoordinator *RewardsCoordinatorCaller) ClaimerFor(opts *bind.CallOpts, earner common.Address) (common.Address, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "claimerFor", earner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address earner) view returns(address claimer)
func (_RewardsCoordinator *RewardsCoordinatorSession) ClaimerFor(earner common.Address) (common.Address, error) {
	return _RewardsCoordinator.Contract.ClaimerFor(&_RewardsCoordinator.CallOpts, earner)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address earner) view returns(address claimer)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) ClaimerFor(earner common.Address) (common.Address, error) {
	return _RewardsCoordinator.Contract.ClaimerFor(&_RewardsCoordinator.CallOpts, earner)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address earner, address token) view returns(uint256 totalClaimed)
func (_RewardsCoordinator *RewardsCoordinatorCaller) CumulativeClaimed(opts *bind.CallOpts, earner common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "cumulativeClaimed", earner, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address earner, address token) view returns(uint256 totalClaimed)
func (_RewardsCoordinator *RewardsCoordinatorSession) CumulativeClaimed(earner common.Address, token common.Address) (*big.Int, error) {
	return _RewardsCoordinator.Contract.CumulativeClaimed(&_RewardsCoordinator.CallOpts, earner, token)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address earner, address token) view returns(uint256 totalClaimed)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) CumulativeClaimed(earner common.Address, token common.Address) (*big.Int, error) {
	return _RewardsCoordinator.Contract.CumulativeClaimed(&_RewardsCoordinator.CallOpts, earner, token)
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

// DefaultOperatorSplitBips is a free data retrieval call binding the contract method 0x63f6a798.
//
// Solidity: function defaultOperatorSplitBips() view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorCaller) DefaultOperatorSplitBips(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "defaultOperatorSplitBips")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DefaultOperatorSplitBips is a free data retrieval call binding the contract method 0x63f6a798.
//
// Solidity: function defaultOperatorSplitBips() view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorSession) DefaultOperatorSplitBips() (uint16, error) {
	return _RewardsCoordinator.Contract.DefaultOperatorSplitBips(&_RewardsCoordinator.CallOpts)
}

// DefaultOperatorSplitBips is a free data retrieval call binding the contract method 0x63f6a798.
//
// Solidity: function defaultOperatorSplitBips() view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) DefaultOperatorSplitBips() (uint16, error) {
	return _RewardsCoordinator.Contract.DefaultOperatorSplitBips(&_RewardsCoordinator.CallOpts)
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

// GetCurrentClaimableDistributionRoot is a free data retrieval call binding the contract method 0x0e9a53cf.
//
// Solidity: function getCurrentClaimableDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorCaller) GetCurrentClaimableDistributionRoot(opts *bind.CallOpts) (IRewardsCoordinatorTypesDistributionRoot, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "getCurrentClaimableDistributionRoot")

	if err != nil {
		return *new(IRewardsCoordinatorTypesDistributionRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(IRewardsCoordinatorTypesDistributionRoot)).(*IRewardsCoordinatorTypesDistributionRoot)

	return out0, err

}

// GetCurrentClaimableDistributionRoot is a free data retrieval call binding the contract method 0x0e9a53cf.
//
// Solidity: function getCurrentClaimableDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorSession) GetCurrentClaimableDistributionRoot() (IRewardsCoordinatorTypesDistributionRoot, error) {
	return _RewardsCoordinator.Contract.GetCurrentClaimableDistributionRoot(&_RewardsCoordinator.CallOpts)
}

// GetCurrentClaimableDistributionRoot is a free data retrieval call binding the contract method 0x0e9a53cf.
//
// Solidity: function getCurrentClaimableDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) GetCurrentClaimableDistributionRoot() (IRewardsCoordinatorTypesDistributionRoot, error) {
	return _RewardsCoordinator.Contract.GetCurrentClaimableDistributionRoot(&_RewardsCoordinator.CallOpts)
}

// GetCurrentDistributionRoot is a free data retrieval call binding the contract method 0x9be3d4e4.
//
// Solidity: function getCurrentDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorCaller) GetCurrentDistributionRoot(opts *bind.CallOpts) (IRewardsCoordinatorTypesDistributionRoot, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "getCurrentDistributionRoot")

	if err != nil {
		return *new(IRewardsCoordinatorTypesDistributionRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(IRewardsCoordinatorTypesDistributionRoot)).(*IRewardsCoordinatorTypesDistributionRoot)

	return out0, err

}

// GetCurrentDistributionRoot is a free data retrieval call binding the contract method 0x9be3d4e4.
//
// Solidity: function getCurrentDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorSession) GetCurrentDistributionRoot() (IRewardsCoordinatorTypesDistributionRoot, error) {
	return _RewardsCoordinator.Contract.GetCurrentDistributionRoot(&_RewardsCoordinator.CallOpts)
}

// GetCurrentDistributionRoot is a free data retrieval call binding the contract method 0x9be3d4e4.
//
// Solidity: function getCurrentDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) GetCurrentDistributionRoot() (IRewardsCoordinatorTypesDistributionRoot, error) {
	return _RewardsCoordinator.Contract.GetCurrentDistributionRoot(&_RewardsCoordinator.CallOpts)
}

// GetDistributionRootAtIndex is a free data retrieval call binding the contract method 0xde02e503.
//
// Solidity: function getDistributionRootAtIndex(uint256 index) view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorCaller) GetDistributionRootAtIndex(opts *bind.CallOpts, index *big.Int) (IRewardsCoordinatorTypesDistributionRoot, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "getDistributionRootAtIndex", index)

	if err != nil {
		return *new(IRewardsCoordinatorTypesDistributionRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(IRewardsCoordinatorTypesDistributionRoot)).(*IRewardsCoordinatorTypesDistributionRoot)

	return out0, err

}

// GetDistributionRootAtIndex is a free data retrieval call binding the contract method 0xde02e503.
//
// Solidity: function getDistributionRootAtIndex(uint256 index) view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorSession) GetDistributionRootAtIndex(index *big.Int) (IRewardsCoordinatorTypesDistributionRoot, error) {
	return _RewardsCoordinator.Contract.GetDistributionRootAtIndex(&_RewardsCoordinator.CallOpts, index)
}

// GetDistributionRootAtIndex is a free data retrieval call binding the contract method 0xde02e503.
//
// Solidity: function getDistributionRootAtIndex(uint256 index) view returns((bytes32,uint32,uint32,bool))
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) GetDistributionRootAtIndex(index *big.Int) (IRewardsCoordinatorTypesDistributionRoot, error) {
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

// GetOperatorAVSSplit is a free data retrieval call binding the contract method 0xe063f81f.
//
// Solidity: function getOperatorAVSSplit(address operator, address avs) view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorCaller) GetOperatorAVSSplit(opts *bind.CallOpts, operator common.Address, avs common.Address) (uint16, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "getOperatorAVSSplit", operator, avs)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetOperatorAVSSplit is a free data retrieval call binding the contract method 0xe063f81f.
//
// Solidity: function getOperatorAVSSplit(address operator, address avs) view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorSession) GetOperatorAVSSplit(operator common.Address, avs common.Address) (uint16, error) {
	return _RewardsCoordinator.Contract.GetOperatorAVSSplit(&_RewardsCoordinator.CallOpts, operator, avs)
}

// GetOperatorAVSSplit is a free data retrieval call binding the contract method 0xe063f81f.
//
// Solidity: function getOperatorAVSSplit(address operator, address avs) view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) GetOperatorAVSSplit(operator common.Address, avs common.Address) (uint16, error) {
	return _RewardsCoordinator.Contract.GetOperatorAVSSplit(&_RewardsCoordinator.CallOpts, operator, avs)
}

// GetOperatorPISplit is a free data retrieval call binding the contract method 0x4b943960.
//
// Solidity: function getOperatorPISplit(address operator) view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorCaller) GetOperatorPISplit(opts *bind.CallOpts, operator common.Address) (uint16, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "getOperatorPISplit", operator)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetOperatorPISplit is a free data retrieval call binding the contract method 0x4b943960.
//
// Solidity: function getOperatorPISplit(address operator) view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorSession) GetOperatorPISplit(operator common.Address) (uint16, error) {
	return _RewardsCoordinator.Contract.GetOperatorPISplit(&_RewardsCoordinator.CallOpts, operator)
}

// GetOperatorPISplit is a free data retrieval call binding the contract method 0x4b943960.
//
// Solidity: function getOperatorPISplit(address operator) view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) GetOperatorPISplit(operator common.Address) (uint16, error) {
	return _RewardsCoordinator.Contract.GetOperatorPISplit(&_RewardsCoordinator.CallOpts, operator)
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

// IsAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0x6d21117e.
//
// Solidity: function isAVSRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorCaller) IsAVSRewardsSubmissionHash(opts *bind.CallOpts, avs common.Address, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "isAVSRewardsSubmissionHash", avs, hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0x6d21117e.
//
// Solidity: function isAVSRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorSession) IsAVSRewardsSubmissionHash(avs common.Address, hash [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsAVSRewardsSubmissionHash(&_RewardsCoordinator.CallOpts, avs, hash)
}

// IsAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0x6d21117e.
//
// Solidity: function isAVSRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) IsAVSRewardsSubmissionHash(avs common.Address, hash [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsAVSRewardsSubmissionHash(&_RewardsCoordinator.CallOpts, avs, hash)
}

// IsOperatorDirectedAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0xed71e6a2.
//
// Solidity: function isOperatorDirectedAVSRewardsSubmissionHash(address , bytes32 ) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCaller) IsOperatorDirectedAVSRewardsSubmissionHash(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "isOperatorDirectedAVSRewardsSubmissionHash", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorDirectedAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0xed71e6a2.
//
// Solidity: function isOperatorDirectedAVSRewardsSubmissionHash(address , bytes32 ) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorSession) IsOperatorDirectedAVSRewardsSubmissionHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsOperatorDirectedAVSRewardsSubmissionHash(&_RewardsCoordinator.CallOpts, arg0, arg1)
}

// IsOperatorDirectedAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0xed71e6a2.
//
// Solidity: function isOperatorDirectedAVSRewardsSubmissionHash(address , bytes32 ) view returns(bool)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) IsOperatorDirectedAVSRewardsSubmissionHash(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsOperatorDirectedAVSRewardsSubmissionHash(&_RewardsCoordinator.CallOpts, arg0, arg1)
}

// IsRewardsForAllSubmitter is a free data retrieval call binding the contract method 0x0018572c.
//
// Solidity: function isRewardsForAllSubmitter(address submitter) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorCaller) IsRewardsForAllSubmitter(opts *bind.CallOpts, submitter common.Address) (bool, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "isRewardsForAllSubmitter", submitter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewardsForAllSubmitter is a free data retrieval call binding the contract method 0x0018572c.
//
// Solidity: function isRewardsForAllSubmitter(address submitter) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorSession) IsRewardsForAllSubmitter(submitter common.Address) (bool, error) {
	return _RewardsCoordinator.Contract.IsRewardsForAllSubmitter(&_RewardsCoordinator.CallOpts, submitter)
}

// IsRewardsForAllSubmitter is a free data retrieval call binding the contract method 0x0018572c.
//
// Solidity: function isRewardsForAllSubmitter(address submitter) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) IsRewardsForAllSubmitter(submitter common.Address) (bool, error) {
	return _RewardsCoordinator.Contract.IsRewardsForAllSubmitter(&_RewardsCoordinator.CallOpts, submitter)
}

// IsRewardsSubmissionForAllEarnersHash is a free data retrieval call binding the contract method 0xaebd8bae.
//
// Solidity: function isRewardsSubmissionForAllEarnersHash(address avs, bytes32 hash) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorCaller) IsRewardsSubmissionForAllEarnersHash(opts *bind.CallOpts, avs common.Address, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "isRewardsSubmissionForAllEarnersHash", avs, hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewardsSubmissionForAllEarnersHash is a free data retrieval call binding the contract method 0xaebd8bae.
//
// Solidity: function isRewardsSubmissionForAllEarnersHash(address avs, bytes32 hash) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorSession) IsRewardsSubmissionForAllEarnersHash(avs common.Address, hash [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsRewardsSubmissionForAllEarnersHash(&_RewardsCoordinator.CallOpts, avs, hash)
}

// IsRewardsSubmissionForAllEarnersHash is a free data retrieval call binding the contract method 0xaebd8bae.
//
// Solidity: function isRewardsSubmissionForAllEarnersHash(address avs, bytes32 hash) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) IsRewardsSubmissionForAllEarnersHash(avs common.Address, hash [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsRewardsSubmissionForAllEarnersHash(&_RewardsCoordinator.CallOpts, avs, hash)
}

// IsRewardsSubmissionForAllHash is a free data retrieval call binding the contract method 0xc46db606.
//
// Solidity: function isRewardsSubmissionForAllHash(address avs, bytes32 hash) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorCaller) IsRewardsSubmissionForAllHash(opts *bind.CallOpts, avs common.Address, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "isRewardsSubmissionForAllHash", avs, hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewardsSubmissionForAllHash is a free data retrieval call binding the contract method 0xc46db606.
//
// Solidity: function isRewardsSubmissionForAllHash(address avs, bytes32 hash) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorSession) IsRewardsSubmissionForAllHash(avs common.Address, hash [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsRewardsSubmissionForAllHash(&_RewardsCoordinator.CallOpts, avs, hash)
}

// IsRewardsSubmissionForAllHash is a free data retrieval call binding the contract method 0xc46db606.
//
// Solidity: function isRewardsSubmissionForAllHash(address avs, bytes32 hash) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) IsRewardsSubmissionForAllHash(avs common.Address, hash [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsRewardsSubmissionForAllHash(&_RewardsCoordinator.CallOpts, avs, hash)
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

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorSession) PermissionController() (common.Address, error) {
	return _RewardsCoordinator.Contract.PermissionController(&_RewardsCoordinator.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) PermissionController() (common.Address, error) {
	return _RewardsCoordinator.Contract.PermissionController(&_RewardsCoordinator.CallOpts)
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
// Solidity: function submissionNonce(address avs) view returns(uint256 nonce)
func (_RewardsCoordinator *RewardsCoordinatorCaller) SubmissionNonce(opts *bind.CallOpts, avs common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "submissionNonce", avs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmissionNonce is a free data retrieval call binding the contract method 0xbb7e451f.
//
// Solidity: function submissionNonce(address avs) view returns(uint256 nonce)
func (_RewardsCoordinator *RewardsCoordinatorSession) SubmissionNonce(avs common.Address) (*big.Int, error) {
	return _RewardsCoordinator.Contract.SubmissionNonce(&_RewardsCoordinator.CallOpts, avs)
}

// SubmissionNonce is a free data retrieval call binding the contract method 0xbb7e451f.
//
// Solidity: function submissionNonce(address avs) view returns(uint256 nonce)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) SubmissionNonce(avs common.Address) (*big.Int, error) {
	return _RewardsCoordinator.Contract.SubmissionNonce(&_RewardsCoordinator.CallOpts, avs)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) CreateAVSRewardsSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "createAVSRewardsSubmission", rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.CreateAVSRewardsSubmission(&_RewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.CreateAVSRewardsSubmission(&_RewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateOperatorDirectedAVSRewardsSubmission is a paid mutator transaction binding the contract method 0x9cb9a5fa.
//
// Solidity: function createOperatorDirectedAVSRewardsSubmission(address avs, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) CreateOperatorDirectedAVSRewardsSubmission(opts *bind.TransactOpts, avs common.Address, operatorDirectedRewardsSubmissions []IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "createOperatorDirectedAVSRewardsSubmission", avs, operatorDirectedRewardsSubmissions)
}

// CreateOperatorDirectedAVSRewardsSubmission is a paid mutator transaction binding the contract method 0x9cb9a5fa.
//
// Solidity: function createOperatorDirectedAVSRewardsSubmission(address avs, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) CreateOperatorDirectedAVSRewardsSubmission(avs common.Address, operatorDirectedRewardsSubmissions []IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.CreateOperatorDirectedAVSRewardsSubmission(&_RewardsCoordinator.TransactOpts, avs, operatorDirectedRewardsSubmissions)
}

// CreateOperatorDirectedAVSRewardsSubmission is a paid mutator transaction binding the contract method 0x9cb9a5fa.
//
// Solidity: function createOperatorDirectedAVSRewardsSubmission(address avs, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) CreateOperatorDirectedAVSRewardsSubmission(avs common.Address, operatorDirectedRewardsSubmissions []IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.CreateOperatorDirectedAVSRewardsSubmission(&_RewardsCoordinator.TransactOpts, avs, operatorDirectedRewardsSubmissions)
}

// CreateRewardsForAllEarners is a paid mutator transaction binding the contract method 0xff9f6cce.
//
// Solidity: function createRewardsForAllEarners(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) CreateRewardsForAllEarners(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "createRewardsForAllEarners", rewardsSubmissions)
}

// CreateRewardsForAllEarners is a paid mutator transaction binding the contract method 0xff9f6cce.
//
// Solidity: function createRewardsForAllEarners(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) CreateRewardsForAllEarners(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.CreateRewardsForAllEarners(&_RewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateRewardsForAllEarners is a paid mutator transaction binding the contract method 0xff9f6cce.
//
// Solidity: function createRewardsForAllEarners(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) CreateRewardsForAllEarners(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.CreateRewardsForAllEarners(&_RewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateRewardsForAllSubmission is a paid mutator transaction binding the contract method 0x36af41fa.
//
// Solidity: function createRewardsForAllSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) CreateRewardsForAllSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "createRewardsForAllSubmission", rewardsSubmissions)
}

// CreateRewardsForAllSubmission is a paid mutator transaction binding the contract method 0x36af41fa.
//
// Solidity: function createRewardsForAllSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) CreateRewardsForAllSubmission(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.CreateRewardsForAllSubmission(&_RewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateRewardsForAllSubmission is a paid mutator transaction binding the contract method 0x36af41fa.
//
// Solidity: function createRewardsForAllSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) CreateRewardsForAllSubmission(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
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

// Initialize is a paid mutator transaction binding the contract method 0xf6efbb59.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus, address _rewardsUpdater, uint32 _activationDelay, uint16 _defaultSplitBips) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int, _rewardsUpdater common.Address, _activationDelay uint32, _defaultSplitBips uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "initialize", initialOwner, initialPausedStatus, _rewardsUpdater, _activationDelay, _defaultSplitBips)
}

// Initialize is a paid mutator transaction binding the contract method 0xf6efbb59.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus, address _rewardsUpdater, uint32 _activationDelay, uint16 _defaultSplitBips) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int, _rewardsUpdater common.Address, _activationDelay uint32, _defaultSplitBips uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.Initialize(&_RewardsCoordinator.TransactOpts, initialOwner, initialPausedStatus, _rewardsUpdater, _activationDelay, _defaultSplitBips)
}

// Initialize is a paid mutator transaction binding the contract method 0xf6efbb59.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus, address _rewardsUpdater, uint32 _activationDelay, uint16 _defaultSplitBips) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int, _rewardsUpdater common.Address, _activationDelay uint32, _defaultSplitBips uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.Initialize(&_RewardsCoordinator.TransactOpts, initialOwner, initialPausedStatus, _rewardsUpdater, _activationDelay, _defaultSplitBips)
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
func (_RewardsCoordinator *RewardsCoordinatorTransactor) ProcessClaim(opts *bind.TransactOpts, claim IRewardsCoordinatorTypesRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "processClaim", claim, recipient)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) ProcessClaim(claim IRewardsCoordinatorTypesRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.ProcessClaim(&_RewardsCoordinator.TransactOpts, claim, recipient)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) ProcessClaim(claim IRewardsCoordinatorTypesRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.ProcessClaim(&_RewardsCoordinator.TransactOpts, claim, recipient)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0x4596021c.
//
// Solidity: function processClaims((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[])[] claims, address recipient) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) ProcessClaims(opts *bind.TransactOpts, claims []IRewardsCoordinatorTypesRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "processClaims", claims, recipient)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0x4596021c.
//
// Solidity: function processClaims((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[])[] claims, address recipient) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) ProcessClaims(claims []IRewardsCoordinatorTypesRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.ProcessClaims(&_RewardsCoordinator.TransactOpts, claims, recipient)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0x4596021c.
//
// Solidity: function processClaims((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[])[] claims, address recipient) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) ProcessClaims(claims []IRewardsCoordinatorTypesRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.ProcessClaims(&_RewardsCoordinator.TransactOpts, claims, recipient)
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

// SetClaimerFor0 is a paid mutator transaction binding the contract method 0xf22cef85.
//
// Solidity: function setClaimerFor(address earner, address claimer) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) SetClaimerFor0(opts *bind.TransactOpts, earner common.Address, claimer common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "setClaimerFor0", earner, claimer)
}

// SetClaimerFor0 is a paid mutator transaction binding the contract method 0xf22cef85.
//
// Solidity: function setClaimerFor(address earner, address claimer) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) SetClaimerFor0(earner common.Address, claimer common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetClaimerFor0(&_RewardsCoordinator.TransactOpts, earner, claimer)
}

// SetClaimerFor0 is a paid mutator transaction binding the contract method 0xf22cef85.
//
// Solidity: function setClaimerFor(address earner, address claimer) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) SetClaimerFor0(earner common.Address, claimer common.Address) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetClaimerFor0(&_RewardsCoordinator.TransactOpts, earner, claimer)
}

// SetDefaultOperatorSplit is a paid mutator transaction binding the contract method 0xa50a1d9c.
//
// Solidity: function setDefaultOperatorSplit(uint16 split) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) SetDefaultOperatorSplit(opts *bind.TransactOpts, split uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "setDefaultOperatorSplit", split)
}

// SetDefaultOperatorSplit is a paid mutator transaction binding the contract method 0xa50a1d9c.
//
// Solidity: function setDefaultOperatorSplit(uint16 split) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) SetDefaultOperatorSplit(split uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetDefaultOperatorSplit(&_RewardsCoordinator.TransactOpts, split)
}

// SetDefaultOperatorSplit is a paid mutator transaction binding the contract method 0xa50a1d9c.
//
// Solidity: function setDefaultOperatorSplit(uint16 split) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) SetDefaultOperatorSplit(split uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetDefaultOperatorSplit(&_RewardsCoordinator.TransactOpts, split)
}

// SetOperatorAVSSplit is a paid mutator transaction binding the contract method 0xdcbb03b3.
//
// Solidity: function setOperatorAVSSplit(address operator, address avs, uint16 split) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) SetOperatorAVSSplit(opts *bind.TransactOpts, operator common.Address, avs common.Address, split uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "setOperatorAVSSplit", operator, avs, split)
}

// SetOperatorAVSSplit is a paid mutator transaction binding the contract method 0xdcbb03b3.
//
// Solidity: function setOperatorAVSSplit(address operator, address avs, uint16 split) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) SetOperatorAVSSplit(operator common.Address, avs common.Address, split uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetOperatorAVSSplit(&_RewardsCoordinator.TransactOpts, operator, avs, split)
}

// SetOperatorAVSSplit is a paid mutator transaction binding the contract method 0xdcbb03b3.
//
// Solidity: function setOperatorAVSSplit(address operator, address avs, uint16 split) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) SetOperatorAVSSplit(operator common.Address, avs common.Address, split uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetOperatorAVSSplit(&_RewardsCoordinator.TransactOpts, operator, avs, split)
}

// SetOperatorPISplit is a paid mutator transaction binding the contract method 0xb3dbb0e0.
//
// Solidity: function setOperatorPISplit(address operator, uint16 split) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) SetOperatorPISplit(opts *bind.TransactOpts, operator common.Address, split uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "setOperatorPISplit", operator, split)
}

// SetOperatorPISplit is a paid mutator transaction binding the contract method 0xb3dbb0e0.
//
// Solidity: function setOperatorPISplit(address operator, uint16 split) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) SetOperatorPISplit(operator common.Address, split uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetOperatorPISplit(&_RewardsCoordinator.TransactOpts, operator, split)
}

// SetOperatorPISplit is a paid mutator transaction binding the contract method 0xb3dbb0e0.
//
// Solidity: function setOperatorPISplit(address operator, uint16 split) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) SetOperatorPISplit(operator common.Address, split uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetOperatorPISplit(&_RewardsCoordinator.TransactOpts, operator, split)
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
	RewardsSubmission     IRewardsCoordinatorTypesRewardsSubmission
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

// RewardsCoordinatorDefaultOperatorSplitBipsSetIterator is returned from FilterDefaultOperatorSplitBipsSet and is used to iterate over the raw logs and unpacked data for DefaultOperatorSplitBipsSet events raised by the RewardsCoordinator contract.
type RewardsCoordinatorDefaultOperatorSplitBipsSetIterator struct {
	Event *RewardsCoordinatorDefaultOperatorSplitBipsSet // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorDefaultOperatorSplitBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorDefaultOperatorSplitBipsSet)
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
		it.Event = new(RewardsCoordinatorDefaultOperatorSplitBipsSet)
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
func (it *RewardsCoordinatorDefaultOperatorSplitBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorDefaultOperatorSplitBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorDefaultOperatorSplitBipsSet represents a DefaultOperatorSplitBipsSet event raised by the RewardsCoordinator contract.
type RewardsCoordinatorDefaultOperatorSplitBipsSet struct {
	OldDefaultOperatorSplitBips uint16
	NewDefaultOperatorSplitBips uint16
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterDefaultOperatorSplitBipsSet is a free log retrieval operation binding the contract event 0xe6cd4edfdcc1f6d130ab35f73d72378f3a642944fb4ee5bd84b7807a81ea1c4e.
//
// Solidity: event DefaultOperatorSplitBipsSet(uint16 oldDefaultOperatorSplitBips, uint16 newDefaultOperatorSplitBips)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterDefaultOperatorSplitBipsSet(opts *bind.FilterOpts) (*RewardsCoordinatorDefaultOperatorSplitBipsSetIterator, error) {

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "DefaultOperatorSplitBipsSet")
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorDefaultOperatorSplitBipsSetIterator{contract: _RewardsCoordinator.contract, event: "DefaultOperatorSplitBipsSet", logs: logs, sub: sub}, nil
}

// WatchDefaultOperatorSplitBipsSet is a free log subscription operation binding the contract event 0xe6cd4edfdcc1f6d130ab35f73d72378f3a642944fb4ee5bd84b7807a81ea1c4e.
//
// Solidity: event DefaultOperatorSplitBipsSet(uint16 oldDefaultOperatorSplitBips, uint16 newDefaultOperatorSplitBips)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchDefaultOperatorSplitBipsSet(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorDefaultOperatorSplitBipsSet) (event.Subscription, error) {

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "DefaultOperatorSplitBipsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorDefaultOperatorSplitBipsSet)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "DefaultOperatorSplitBipsSet", log); err != nil {
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

// ParseDefaultOperatorSplitBipsSet is a log parse operation binding the contract event 0xe6cd4edfdcc1f6d130ab35f73d72378f3a642944fb4ee5bd84b7807a81ea1c4e.
//
// Solidity: event DefaultOperatorSplitBipsSet(uint16 oldDefaultOperatorSplitBips, uint16 newDefaultOperatorSplitBips)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseDefaultOperatorSplitBipsSet(log types.Log) (*RewardsCoordinatorDefaultOperatorSplitBipsSet, error) {
	event := new(RewardsCoordinatorDefaultOperatorSplitBipsSet)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "DefaultOperatorSplitBipsSet", log); err != nil {
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

// RewardsCoordinatorOperatorAVSSplitBipsSetIterator is returned from FilterOperatorAVSSplitBipsSet and is used to iterate over the raw logs and unpacked data for OperatorAVSSplitBipsSet events raised by the RewardsCoordinator contract.
type RewardsCoordinatorOperatorAVSSplitBipsSetIterator struct {
	Event *RewardsCoordinatorOperatorAVSSplitBipsSet // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorOperatorAVSSplitBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorOperatorAVSSplitBipsSet)
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
		it.Event = new(RewardsCoordinatorOperatorAVSSplitBipsSet)
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
func (it *RewardsCoordinatorOperatorAVSSplitBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorOperatorAVSSplitBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorOperatorAVSSplitBipsSet represents a OperatorAVSSplitBipsSet event raised by the RewardsCoordinator contract.
type RewardsCoordinatorOperatorAVSSplitBipsSet struct {
	Caller                  common.Address
	Operator                common.Address
	Avs                     common.Address
	ActivatedAt             uint32
	OldOperatorAVSSplitBips uint16
	NewOperatorAVSSplitBips uint16
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterOperatorAVSSplitBipsSet is a free log retrieval operation binding the contract event 0x48e198b6ae357e529204ee53a8e514c470ff77d9cc8e4f7207f8b5d490ae6934.
//
// Solidity: event OperatorAVSSplitBipsSet(address indexed caller, address indexed operator, address indexed avs, uint32 activatedAt, uint16 oldOperatorAVSSplitBips, uint16 newOperatorAVSSplitBips)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterOperatorAVSSplitBipsSet(opts *bind.FilterOpts, caller []common.Address, operator []common.Address, avs []common.Address) (*RewardsCoordinatorOperatorAVSSplitBipsSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "OperatorAVSSplitBipsSet", callerRule, operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorOperatorAVSSplitBipsSetIterator{contract: _RewardsCoordinator.contract, event: "OperatorAVSSplitBipsSet", logs: logs, sub: sub}, nil
}

// WatchOperatorAVSSplitBipsSet is a free log subscription operation binding the contract event 0x48e198b6ae357e529204ee53a8e514c470ff77d9cc8e4f7207f8b5d490ae6934.
//
// Solidity: event OperatorAVSSplitBipsSet(address indexed caller, address indexed operator, address indexed avs, uint32 activatedAt, uint16 oldOperatorAVSSplitBips, uint16 newOperatorAVSSplitBips)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchOperatorAVSSplitBipsSet(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorOperatorAVSSplitBipsSet, caller []common.Address, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "OperatorAVSSplitBipsSet", callerRule, operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorOperatorAVSSplitBipsSet)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "OperatorAVSSplitBipsSet", log); err != nil {
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

// ParseOperatorAVSSplitBipsSet is a log parse operation binding the contract event 0x48e198b6ae357e529204ee53a8e514c470ff77d9cc8e4f7207f8b5d490ae6934.
//
// Solidity: event OperatorAVSSplitBipsSet(address indexed caller, address indexed operator, address indexed avs, uint32 activatedAt, uint16 oldOperatorAVSSplitBips, uint16 newOperatorAVSSplitBips)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseOperatorAVSSplitBipsSet(log types.Log) (*RewardsCoordinatorOperatorAVSSplitBipsSet, error) {
	event := new(RewardsCoordinatorOperatorAVSSplitBipsSet)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "OperatorAVSSplitBipsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator is returned from FilterOperatorDirectedAVSRewardsSubmissionCreated and is used to iterate over the raw logs and unpacked data for OperatorDirectedAVSRewardsSubmissionCreated events raised by the RewardsCoordinator contract.
type RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator struct {
	Event *RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated)
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
		it.Event = new(RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated)
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
func (it *RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated represents a OperatorDirectedAVSRewardsSubmissionCreated event raised by the RewardsCoordinator contract.
type RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated struct {
	Caller                                common.Address
	Avs                                   common.Address
	OperatorDirectedRewardsSubmissionHash [32]byte
	SubmissionNonce                       *big.Int
	OperatorDirectedRewardsSubmission     IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission
	Raw                                   types.Log // Blockchain specific contextual infos
}

// FilterOperatorDirectedAVSRewardsSubmissionCreated is a free log retrieval operation binding the contract event 0xfc8888bffd711da60bc5092b33f677d81896fe80ecc677b84cfab8184462b6e0.
//
// Solidity: event OperatorDirectedAVSRewardsSubmissionCreated(address indexed caller, address indexed avs, bytes32 indexed operatorDirectedRewardsSubmissionHash, uint256 submissionNonce, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string) operatorDirectedRewardsSubmission)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterOperatorDirectedAVSRewardsSubmissionCreated(opts *bind.FilterOpts, caller []common.Address, avs []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (*RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var operatorDirectedRewardsSubmissionHashRule []interface{}
	for _, operatorDirectedRewardsSubmissionHashItem := range operatorDirectedRewardsSubmissionHash {
		operatorDirectedRewardsSubmissionHashRule = append(operatorDirectedRewardsSubmissionHashRule, operatorDirectedRewardsSubmissionHashItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "OperatorDirectedAVSRewardsSubmissionCreated", callerRule, avsRule, operatorDirectedRewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator{contract: _RewardsCoordinator.contract, event: "OperatorDirectedAVSRewardsSubmissionCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorDirectedAVSRewardsSubmissionCreated is a free log subscription operation binding the contract event 0xfc8888bffd711da60bc5092b33f677d81896fe80ecc677b84cfab8184462b6e0.
//
// Solidity: event OperatorDirectedAVSRewardsSubmissionCreated(address indexed caller, address indexed avs, bytes32 indexed operatorDirectedRewardsSubmissionHash, uint256 submissionNonce, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string) operatorDirectedRewardsSubmission)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchOperatorDirectedAVSRewardsSubmissionCreated(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated, caller []common.Address, avs []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var operatorDirectedRewardsSubmissionHashRule []interface{}
	for _, operatorDirectedRewardsSubmissionHashItem := range operatorDirectedRewardsSubmissionHash {
		operatorDirectedRewardsSubmissionHashRule = append(operatorDirectedRewardsSubmissionHashRule, operatorDirectedRewardsSubmissionHashItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "OperatorDirectedAVSRewardsSubmissionCreated", callerRule, avsRule, operatorDirectedRewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "OperatorDirectedAVSRewardsSubmissionCreated", log); err != nil {
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

// ParseOperatorDirectedAVSRewardsSubmissionCreated is a log parse operation binding the contract event 0xfc8888bffd711da60bc5092b33f677d81896fe80ecc677b84cfab8184462b6e0.
//
// Solidity: event OperatorDirectedAVSRewardsSubmissionCreated(address indexed caller, address indexed avs, bytes32 indexed operatorDirectedRewardsSubmissionHash, uint256 submissionNonce, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string) operatorDirectedRewardsSubmission)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseOperatorDirectedAVSRewardsSubmissionCreated(log types.Log) (*RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated, error) {
	event := new(RewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "OperatorDirectedAVSRewardsSubmissionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsCoordinatorOperatorPISplitBipsSetIterator is returned from FilterOperatorPISplitBipsSet and is used to iterate over the raw logs and unpacked data for OperatorPISplitBipsSet events raised by the RewardsCoordinator contract.
type RewardsCoordinatorOperatorPISplitBipsSetIterator struct {
	Event *RewardsCoordinatorOperatorPISplitBipsSet // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorOperatorPISplitBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorOperatorPISplitBipsSet)
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
		it.Event = new(RewardsCoordinatorOperatorPISplitBipsSet)
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
func (it *RewardsCoordinatorOperatorPISplitBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorOperatorPISplitBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorOperatorPISplitBipsSet represents a OperatorPISplitBipsSet event raised by the RewardsCoordinator contract.
type RewardsCoordinatorOperatorPISplitBipsSet struct {
	Caller                 common.Address
	Operator               common.Address
	ActivatedAt            uint32
	OldOperatorPISplitBips uint16
	NewOperatorPISplitBips uint16
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterOperatorPISplitBipsSet is a free log retrieval operation binding the contract event 0xd1e028bd664486a46ad26040e999cd2d22e1e9a094ee6afe19fcf64678f16f74.
//
// Solidity: event OperatorPISplitBipsSet(address indexed caller, address indexed operator, uint32 activatedAt, uint16 oldOperatorPISplitBips, uint16 newOperatorPISplitBips)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterOperatorPISplitBipsSet(opts *bind.FilterOpts, caller []common.Address, operator []common.Address) (*RewardsCoordinatorOperatorPISplitBipsSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "OperatorPISplitBipsSet", callerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorOperatorPISplitBipsSetIterator{contract: _RewardsCoordinator.contract, event: "OperatorPISplitBipsSet", logs: logs, sub: sub}, nil
}

// WatchOperatorPISplitBipsSet is a free log subscription operation binding the contract event 0xd1e028bd664486a46ad26040e999cd2d22e1e9a094ee6afe19fcf64678f16f74.
//
// Solidity: event OperatorPISplitBipsSet(address indexed caller, address indexed operator, uint32 activatedAt, uint16 oldOperatorPISplitBips, uint16 newOperatorPISplitBips)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchOperatorPISplitBipsSet(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorOperatorPISplitBipsSet, caller []common.Address, operator []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "OperatorPISplitBipsSet", callerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorOperatorPISplitBipsSet)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "OperatorPISplitBipsSet", log); err != nil {
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

// ParseOperatorPISplitBipsSet is a log parse operation binding the contract event 0xd1e028bd664486a46ad26040e999cd2d22e1e9a094ee6afe19fcf64678f16f74.
//
// Solidity: event OperatorPISplitBipsSet(address indexed caller, address indexed operator, uint32 activatedAt, uint16 oldOperatorPISplitBips, uint16 newOperatorPISplitBips)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseOperatorPISplitBipsSet(log types.Log) (*RewardsCoordinatorOperatorPISplitBipsSet, error) {
	event := new(RewardsCoordinatorOperatorPISplitBipsSet)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "OperatorPISplitBipsSet", log); err != nil {
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
	RewardsSubmission     IRewardsCoordinatorTypesRewardsSubmission
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
	RewardsSubmission     IRewardsCoordinatorTypesRewardsSubmission
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
