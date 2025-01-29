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

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// RewardsCoordinatorMetaData contains all meta data concerning the RewardsCoordinator contract.
var RewardsCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_CALCULATION_INTERVAL_SECONDS\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_MAX_REWARDS_DURATION\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_MAX_RETROACTIVE_LENGTH\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_MAX_FUTURE_LENGTH\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_GENESIS_REWARDS_TIMESTAMP\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CALCULATION_INTERVAL_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GENESIS_REWARDS_TIMESTAMP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FUTURE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_RETROACTIVE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_REWARDS_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activationDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateEarnerLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"calculateTokenLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"checkClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimerFor\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAVSRewardsSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorDirectedAVSRewardsSubmission\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorDirectedRewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorDirectedOperatorSetRewardsSubmission\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorDirectedRewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRewardsForAllEarners\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRewardsForAllSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cumulativeClaimed\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"totalClaimed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currRewardsCalculationEndTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultOperatorSplitBips\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableRoot\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentClaimableDistributionRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentDistributionRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributionRootAtIndex\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributionRootsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorAVSSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorPISplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRootIndexFromHash\",\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_rewardsUpdater\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_defaultSplitBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAVSRewardsSubmissionHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorDirectedAVSRewardsSubmissionHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorDirectedOperatorSetRewardsSubmissionHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRewardsForAllSubmitter\",\"inputs\":[{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRewardsSubmissionForAllEarnersHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRewardsSubmissionForAllHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"processClaims\",\"inputs\":[{\"name\":\"claims\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsMerkleClaim[]\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardsUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setActivationDelay\",\"inputs\":[{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimerFor\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimerFor\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultOperatorSplit\",\"inputs\":[{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorAVSSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorPISplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsForAllSubmitter\",\"inputs\":[{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_newValue\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsUpdater\",\"inputs\":[{\"name\":\"_rewardsUpdater\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submissionNonce\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSRewardsSubmissionCreated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ActivationDelaySet\",\"inputs\":[{\"name\":\"oldActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimerForSet\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldClaimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultOperatorSplitBipsSet\",\"inputs\":[{\"name\":\"oldDefaultOperatorSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newDefaultOperatorSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootDisabled\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootSubmitted\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSSplitBipsSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"oldOperatorAVSSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newOperatorAVSSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDirectedAVSRewardsSubmissionCreated\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorDirectedRewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"operatorDirectedRewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDirectedOperatorSetRewardsSubmissionCreated\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorDirectedRewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"operatorDirectedRewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorPISplitBipsSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"oldOperatorPISplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newOperatorPISplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetSplitBipsSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"oldOperatorSetSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newOperatorSetSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsClaimed\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"claimedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsForAllSubmitterSet\",\"inputs\":[{\"name\":\"rewardsForAllSubmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"newValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsSubmissionForAllCreated\",\"inputs\":[{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsSubmissionForAllEarnersCreated\",\"inputs\":[{\"name\":\"tokenHopper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsUpdaterSet\",\"inputs\":[{\"name\":\"oldRewardsUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newRewardsUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AmountExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AmountIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DurationExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EarningsNotGreaterThanClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCalculationIntervalSecondsRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidClaimProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDurationRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEarner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEarnerLeafIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGenesisRewardsTimestampRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRootIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidStartTimestampRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTokenLeafIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewRootMustBeForNewCalculatedPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorsNotInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PreviousSplitPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardsEndTimestampNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootAlreadyActivated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootNotActivated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SplitExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartTimestampTooFarInFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartTimestampTooFarInPast\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategiesNotInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SubmissionNotRetroactive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[]}]",
	Bin: "0x6101c0604052348015610010575f5ffd5b50604051614a90380380614a9083398101604081905261002f91610211565b858a8a8a88888888888f6001600160a01b038116610060576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b031660805261007685826102d3565b63ffffffff161561009a57604051630e06bd3160e01b815260040160405180910390fd5b6100a762015180866102d3565b63ffffffff16156100cb5760405163223c7b3960e11b815260040160405180910390fd5b6001600160a01b0397881660a05295871660c05293861660e05263ffffffff9283166101005290821661012052811661014052908116610160521661018052166101a052610117610126565b50505050505050505050610306565b5f54610100900460ff16156101915760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146101e0575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101f6575f5ffd5b50565b805163ffffffff8116811461020c575f5ffd5b919050565b5f5f5f5f5f5f5f5f5f5f6101408b8d03121561022b575f5ffd5b8a51610236816101e2565b60208c0151909a50610247816101e2565b60408c0151909950610258816101e2565b60608c0151909850610269816101e2565b60808c015190975061027a816101e2565b955061028860a08c016101f9565b945061029660c08c016101f9565b93506102a460e08c016101f9565b92506102b36101008c016101f9565b91506102c26101208c016101f9565b90509295989b9194979a5092959850565b5f63ffffffff8316806102f457634e487b7160e01b5f52601260045260245ffd5b8063ffffffff84160691505092915050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516146b26103de5f395f81816105e5015261278101525f8181610499015261352801525f81816103e80152612bc301525f818161054701526134e601525f818161085e01526133f601525f818161079f01528181613446015261349401525f81816108b201528181610adb01528181611dd6015261208201525f818161056e01526135c301525f81816109250152611d4601525f818161073101528181610eb9015281816114bb015261236101526146b25ff3fe608060405234801561000f575f5ffd5b50600436106103a8575f3560e01c80638da5cb5b116101ea578063de02e50311610114578063f6efbb59116100a9578063fabc1cbc11610079578063fabc1cbc14610a13578063fbf1e2c114610a26578063fce36c7d14610a39578063ff9f6cce14610a4c575f5ffd5b8063f6efbb59146109c7578063f74e8eac146109da578063f8cd8448146109ed578063f96abf2e14610a00575f5ffd5b8063ed71e6a2116100e4578063ed71e6a214610947578063f22cef8514610974578063f2f07ab414610987578063f2fde38b146109b4575f5ffd5b8063de02e503146108e7578063e063f81f146108fa578063e810ce211461090d578063ea4d3c9b14610920575f5ffd5b8063a50a1d9c1161018a578063bf21a8aa1161015a578063bf21a8aa14610859578063c46db60614610880578063ca8aa7c7146108ad578063dcbb03b3146108d4575f5ffd5b8063a50a1d9c146107e7578063aebd8bae146107fa578063b3dbb0e014610827578063bb7e451f1461083a575f5ffd5b80639cb9a5fa116101c55780639cb9a5fa146107875780639d45c2811461079a5780639de4b35f146107c1578063a0169ddd146107d4575f5ffd5b80638da5cb5b146107535780639104c319146107645780639be3d4e41461077f575f5ffd5b80634596021c116102d65780635e9d83481161026b5780637b8f8b051161023b5780637b8f8b05146106e7578063863cb9a9146106ef578063865c695314610702578063886f11951461072c575f5ffd5b80635e9d83481461068a57806363f6a7981461069d5780636d21117e146106b2578063715018a6146106df575f5ffd5b806358baaa3e116102a657806358baaa3e14610644578063595c6a67146106575780635ac86ab71461065f5780635c975abb14610682575f5ffd5b80634596021c146105cd5780634657e26a146105e05780634b943960146106075780634d18cc351461062d575f5ffd5b8063149bc8721161034c57806339b70e381161031c57806339b70e38146105695780633a8c0786146105905780633ccc861d146105a75780633efe1db6146105ba575f5ffd5b8063149bc872146104ce5780632b9f64a4146104ef57806336af41fa1461052f57806337838ed014610542575f5ffd5b80630e9a53cf116103875780630e9a53cf146104345780630eb3834514610481578063131433b414610494578063136439dd146104bb575f5ffd5b806218572c146103ac57806304a0c502146103e35780630ca298991461041f575b5f5ffd5b6103ce6103ba366004613adf565b60d16020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61040a7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016103da565b61043261042d366004613b57565b610a5f565b005b61043c610d24565b6040516103da91905f6080820190508251825263ffffffff602084015116602083015263ffffffff604084015116604083015260608301511515606083015292915050565b61043261048f366004613bb3565b610e24565b61040a7f000000000000000000000000000000000000000000000000000000000000000081565b6104326104c9366004613bea565b610ea4565b6104e16104dc366004613c01565b610f79565b6040519081526020016103da565b6105176104fd366004613adf565b60cc6020525f90815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020016103da565b61043261053d366004613c1b565b610fee565b61040a7f000000000000000000000000000000000000000000000000000000000000000081565b6105177f000000000000000000000000000000000000000000000000000000000000000081565b60cb5461040a90600160a01b900463ffffffff1681565b6104326105b5366004613c6a565b61117d565b6104326105c8366004613cc0565b6111c2565b6104326105db366004613cea565b6113b6565b6105177f000000000000000000000000000000000000000000000000000000000000000081565b61061a610615366004613adf565b611437565b60405161ffff90911681526020016103da565b60cb5461040a90600160c01b900463ffffffff1681565b610432610652366004613d3c565b611492565b6104326114a6565b6103ce61066d366004613d55565b606654600160ff9092169190911b9081161490565b6066546104e1565b6103ce610698366004613d75565b611555565b60cb5461061a90600160e01b900461ffff1681565b6103ce6106c0366004613da6565b60cf60209081525f928352604080842090915290825290205460ff1681565b6104326115e0565b60ca546104e1565b6104326106fd366004613adf565b6115f1565b6104e1610710366004613dd0565b60cd60209081525f928352604080842090915290825290205481565b6105177f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b0316610517565b61051773beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b61043c611602565b610432610795366004613dfc565b61169e565b61040a7f000000000000000000000000000000000000000000000000000000000000000081565b61061a6107cf366004613e33565b611837565b6104326107e2366004613adf565b6118ba565b6104326107f5366004613e6f565b6118c5565b6103ce610808366004613da6565b60d260209081525f928352604080842090915290825290205460ff1681565b610432610835366004613e88565b6118d6565b6104e1610848366004613adf565b60ce6020525f908152604090205481565b61040a7f000000000000000000000000000000000000000000000000000000000000000081565b6103ce61088e366004613da6565b60d060209081525f928352604080842090915290825290205460ff1681565b6105177f000000000000000000000000000000000000000000000000000000000000000081565b6104326108e2366004613eb2565b611a1f565b61043c6108f5366004613bea565b611b8a565b61061a610908366004613dd0565b611c1a565b61040a61091b366004613bea565b611c7f565b6105177f000000000000000000000000000000000000000000000000000000000000000081565b6103ce610955366004613da6565b60d360209081525f928352604080842090915290825290205460ff1681565b610432610982366004613dd0565b611d00565b6103ce610995366004613da6565b60d760209081525f928352604080842090915290825290205460ff1681565b6104326109c2366004613adf565b611e6a565b6104326109d5366004613ef6565b611ee5565b6104326109e8366004613f54565b61201a565b6104e16109fb366004613c01565b612200565b610432610a0e366004613d3c565b612210565b610432610a21366004613bea565b61235f565b60cb54610517906001600160a01b031681565b610432610a47366004613c1b565b612475565b610432610a5a366004613c1b565b6125c4565b60665460099061020090811603610a895760405163840a48d560e01b815260040160405180910390fd5b610a966020850185613adf565b610a9f81612743565b610abc5760405163932d94f760e01b815260040160405180910390fd5b610ac46127ed565b6040516304c1b8eb60e31b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc75890610b10908890600401613fbf565b602060405180830381865afa158015610b2b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b4f9190613fcd565b610b6c57604051631fb1705560e21b815260040160405180910390fd5b5f5b83811015610d125736858583818110610b8957610b89613fe8565b9050602002810190610b9b9190613ffc565b90505f60ce81610bae60208b018b613adf565b6001600160a01b031681526020808201929092526040015f90812054925090610bd9908a018a613adf565b8284604051602001610bed93929190614224565b6040516020818303038152906040528051906020012090505f610c0f84612846565b9050600160d75f610c2360208e018e613adf565b6001600160a01b0316815260208082019290925260409081015f9081208682529092529020805460ff1916911515919091179055610c62836001614267565b60ce5f610c7260208e018e613adf565b6001600160a01b03166001600160a01b031681526020019081526020015f208190555081336001600160a01b03167ffff0759ccb371dfb5691798724e70b4fa61cb3bfe730a33ac19fb86a48efc7568c8688604051610cd39392919061427a565b60405180910390a3610d02333083610cf16040890160208a01613adf565b6001600160a01b0316929190612a31565b505060019092019150610b6e9050565b50610d1d6001609755565b5050505050565b604080516080810182525f80825260208201819052918101829052606081019190915260ca545b8015610dfc575f60ca610d5f60018461429f565b81548110610d6f57610d6f613fe8565b5f91825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff161580156060830181905291925090610dde5750806040015163ffffffff164210155b15610de95792915050565b5080610df4816142b2565b915050610d4b565b5050604080516080810182525f80825260208201819052918101829052606081019190915290565b610e2c612a9c565b6001600160a01b0382165f81815260d1602052604080822054905160ff9091169284151592841515927f4de6293e668df1398422e1def12118052c1539a03cbfedc145895d48d7685f1c9190a4506001600160a01b03919091165f90815260d160205260409020805460ff1916911515919091179055565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610f06573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f2a9190613fcd565b610f4757604051631d77d47760e21b815260040160405180910390fd5b6066548181168114610f6c5760405163c61dca5d60e01b815260040160405180910390fd5b610f7582612af6565b5050565b5f80610f886020840184613adf565b8360200135604051602001610fd19392919060f89390931b6001600160f81b031916835260609190911b6bffffffffffffffffffffffff19166001830152601582015260350190565b604051602081830303815290604052805190602001209050919050565b6066546001906002908116036110175760405163840a48d560e01b815260040160405180910390fd5b335f90815260d1602052604090205460ff1661104657604051635c427cd960e01b815260040160405180910390fd5b61104e6127ed565b5f5b8281101561116d573684848381811061106b5761106b613fe8565b905060200281019061107d91906142c7565b335f81815260ce602090815260408083205490519495509391926110a79290918591879101614359565b6040516020818303038152906040528051906020012090506110c883612b33565b335f90815260d0602090815260408083208484529091529020805460ff191660019081179091556110fa908390614267565b335f81815260ce602052604090819020929092559051829184917f51088b8c89628df3a8174002c2a034d0152fce6af8415d651b2a4734bf2704829061114190889061437f565b60405180910390a4611162333060408601803590610cf19060208901613adf565b505050600101611050565b506111786001609755565b505050565b6066546002906004908116036111a65760405163840a48d560e01b815260040160405180910390fd5b6111ae6127ed565b6111b88383612c1e565b6111786001609755565b6066546003906008908116036111eb5760405163840a48d560e01b815260040160405180910390fd5b60cb546001600160a01b0316331461121657604051635c427cd960e01b815260040160405180910390fd5b60cb5463ffffffff600160c01b90910481169083161161124957604051631ca7e50b60e21b815260040160405180910390fd5b428263ffffffff161061126f576040516306957c9160e11b815260040160405180910390fd5b60ca5460cb545f9061128e90600160a01b900463ffffffff1642614391565b6040805160808101825287815263ffffffff87811660208084018281528684168587018181525f6060880181815260ca8054600181018255925297517f42d72674974f694b5f5159593243114d38a5c39c89d6b62fee061ff523240ee160029092029182015592517f42d72674974f694b5f5159593243114d38a5c39c89d6b62fee061ff523240ee290930180549151975193871667ffffffffffffffff1990921691909117600160201b978716979097029690961760ff60401b1916600160401b921515929092029190911790945560cb805463ffffffff60c01b1916600160c01b840217905593519283529394508892908616917fecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08910160405180910390a45050505050565b6066546002906004908116036113df5760405163840a48d560e01b815260040160405180910390fd5b6113e76127ed565b5f5b838110156114265761141e85858381811061140657611406613fe8565b905060200281019061141891906143ad565b84612c1e565b6001016113e9565b506114316001609755565b50505050565b6001600160a01b0381165f90815260d5602090815260408083208151606081018352905461ffff80821683526201000082041693820193909352600160201b90920463ffffffff169082015261148c90612ea6565b92915050565b61149a612a9c565b6114a381612f16565b50565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611508573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061152c9190613fcd565b61154957604051631d77d47760e21b815260040160405180910390fd5b6115535f19612af6565b565b5f6115d88260ca6115696020830183613d3c565b63ffffffff168154811061157f5761157f613fe8565b5f91825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff1615156060820152612f87565b506001919050565b6115e8612a9c565b6115535f61312a565b6115f9612a9c565b6114a38161317b565b604080516080810182525f80825260208201819052918101829052606081019190915260ca80546116359060019061429f565b8154811061164557611645613fe8565b5f91825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff1615156060820152919050565b6066546005906020908116036116c75760405163840a48d560e01b815260040160405180910390fd5b836116d181612743565b6116ee5760405163932d94f760e01b815260040160405180910390fd5b6116f66127ed565b5f5b83811015610d12573685858381811061171357611713613fe8565b90506020028101906117259190613ffc565b6001600160a01b0388165f90815260ce6020908152604080832054905193945092611756918b918591879101614224565b6040516020818303038152906040528051906020012090505f61177884612846565b6001600160a01b038b165f90815260d3602090815260408083208684529091529020805460ff191660019081179091559091506117b6908490614267565b6001600160a01b038b165f81815260ce60205260409081902092909255905183919033907ffc8888bffd711da60bc5092b33f677d81896fe80ecc677b84cfab8184462b6e0906118099088908a906143c1565b60405180910390a4611827333083610cf16040890160208a01613adf565b5050600190920191506116f89050565b6001600160a01b0382165f90815260d6602052604081206118b3908261186a611865368790038701876143d9565b6131d6565b815260208082019290925260409081015f208151606081018352905461ffff80821683526201000082041693820193909352600160201b90920463ffffffff1690820152612ea6565b9392505050565b33610f758183613239565b6118cd612a9c565b6114a38161329c565b6066546007906080908116036118ff5760405163840a48d560e01b815260040160405180910390fd5b8261190981612743565b6119265760405163932d94f760e01b815260040160405180910390fd5b60cb545f9061194290600160a01b900463ffffffff1642614391565b6001600160a01b0386165f90815260d5602090815260408083208151606081018352905461ffff80821683526201000082041693820193909352600160201b90920463ffffffff16908201529192509061199b90612ea6565b6001600160a01b0387165f90815260d5602052604090209091506119c0908684613307565b6040805163ffffffff8416815261ffff838116602083015287168183015290516001600160a01b0388169133917fd1e028bd664486a46ad26040e999cd2d22e1e9a094ee6afe19fcf64678f16f749181900360600190a3505050505050565b606654600690604090811603611a485760405163840a48d560e01b815260040160405180910390fd5b83611a5281612743565b611a6f5760405163932d94f760e01b815260040160405180910390fd5b60cb545f90611a8b90600160a01b900463ffffffff1642614391565b6001600160a01b038781165f90815260d460209081526040808320938a1683529281528282208351606081018552905461ffff80821683526201000082041692820192909252600160201b90910463ffffffff1692810192909252919250611af290612ea6565b6001600160a01b038089165f90815260d460209081526040808320938b16835292905220909150611b24908684613307565b6040805163ffffffff8416815261ffff838116602083015287168183015290516001600160a01b0388811692908a169133917f48e198b6ae357e529204ee53a8e514c470ff77d9cc8e4f7207f8b5d490ae6934919081900360600190a450505050505050565b604080516080810182525f80825260208201819052918101829052606081019190915260ca8281548110611bc057611bc0613fe8565b5f91825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff161515606082015292915050565b6001600160a01b038281165f90815260d46020908152604080832093851683529281528282208351606081018552905461ffff80821683526201000082041692820192909252600160201b90910463ffffffff1692810192909252906118b390612ea6565b60ca545f905b63ffffffff811615611ce6578260ca611c9f600184614441565b63ffffffff1681548110611cb557611cb5613fe8565b905f5260205f2090600202015f015403611cd4576118b3600182614441565b80611cde8161445d565b915050611c85565b5060405163504570e360e01b815260040160405180910390fd5b81611d0a81612743565b611d275760405163932d94f760e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0384811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015611d8b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611daf9190613fcd565b80611e43575060405163ba1a84e560e01b81526001600160a01b0384811660048301525f917f00000000000000000000000000000000000000000000000000000000000000009091169063ba1a84e590602401602060405180830381865afa158015611e1d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e41919061447b565b115b611e605760405163fb494ea160e01b815260040160405180910390fd5b6111788383613239565b611e72612a9c565b6001600160a01b038116611edc5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b6114a38161312a565b5f54610100900460ff1615808015611f0357505f54600160ff909116105b80611f1c5750303b158015611f1c57505f5460ff166001145b611f7f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401611ed3565b5f805460ff191660011790558015611fa0575f805461ff0019166101001790555b611fa985612af6565b611fb28661312a565b611fbb8461317b565b611fc483612f16565b611fcd8261329c565b8015612012575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b606654600890610100908116036120445760405163840a48d560e01b815260040160405180910390fd5b8361204e81612743565b61206b5760405163932d94f760e01b815260040160405180910390fd5b6040516304c1b8eb60e31b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc758906120b7908790600401613fbf565b602060405180830381865afa1580156120d2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120f69190613fcd565b61211357604051631fb1705560e21b815260040160405180910390fd5b60cb545f9061212f90600160a01b900463ffffffff1642614391565b6001600160a01b0387165f90815260d66020526040812091925090612161908261186a611865368b90038b018b6143d9565b6001600160a01b0388165f90815260d6602052604081209192506121a69190612192611865368b90038b018b6143d9565b81526020019081526020015f208684613307565b866001600160a01b0316336001600160a01b03167f14918b3834ab6752eb2e1b489b6663a67810efb5f56f3944a97ede8ecf1fd9f18885858a6040516121ef9493929190614492565b60405180910390a350505050505050565b5f6001610f886020840184613adf565b6066546003906008908116036122395760405163840a48d560e01b815260040160405180910390fd5b60cb546001600160a01b0316331461226457604051635c427cd960e01b815260040160405180910390fd5b60ca5463ffffffff83161061228c576040516394a8d38960e01b815260040160405180910390fd5b5f60ca8363ffffffff16815481106122a6576122a6613fe8565b905f5260205f20906002020190508060010160089054906101000a900460ff16156122e457604051631b14174b60e01b815260040160405180910390fd5b6001810154600160201b900463ffffffff16421061231557604051630c36f66560e21b815260040160405180910390fd5b60018101805460ff60401b1916600160401b17905560405163ffffffff8416907fd850e6e5dfa497b72661fa73df2923464eaed9dc2ff1d3cb82bccbfeabe5c41e905f90a2505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156123bb573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123df91906144c6565b6001600160a01b0316336001600160a01b0316146124105760405163794821ff60e01b815260040160405180910390fd5b606654801982198116146124375760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b6066545f9060019081160361249d5760405163840a48d560e01b815260040160405180910390fd5b6124a56127ed565b5f5b8281101561116d57368484838181106124c2576124c2613fe8565b90506020028101906124d491906142c7565b335f81815260ce602090815260408083205490519495509391926124fe9290918591879101614359565b60405160208183030381529060405280519060200120905061251f83612b33565b335f90815260cf602090815260408083208484529091529020805460ff19166001908117909155612551908390614267565b335f81815260ce602052604090819020929092559051829184917f450a367a380c4e339e5ae7340c8464ef27af7781ad9945cfe8abd828f89e62819061259890889061437f565b60405180910390a46125b9333060408601803590610cf19060208901613adf565b5050506001016124a7565b6066546004906010908116036125ed5760405163840a48d560e01b815260040160405180910390fd5b335f90815260d1602052604090205460ff1661261c57604051635c427cd960e01b815260040160405180910390fd5b6126246127ed565b5f5b8281101561116d573684848381811061264157612641613fe8565b905060200281019061265391906142c7565b335f81815260ce6020908152604080832054905194955093919261267d9290918591879101614359565b60405160208183030381529060405280519060200120905061269e83612b33565b335f90815260d2602090815260408083208484529091529020805460ff191660019081179091556126d0908390614267565b335f81815260ce602052604090819020929092559051829184917f5251b6fdefcb5d81144e735f69ea4c695fd43b0289ca53dc075033f5fc80068b9061271790889061437f565b60405180910390a4612738333060408601803590610cf19060208901613adf565b505050600101612626565b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af11580156127c9573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061148c9190613fcd565b60026097540361283f5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611ed3565b6002609755565b5f61287961285483806144e1565b6128646080860160608701613d3c565b61287460a0870160808801613d3c565b6133d6565b5f61288760408401846144e1565b9050116128a75760405163796cc52560e01b815260040160405180910390fd5b426128b860a0840160808501613d3c565b6128c86080850160608601613d3c565b6128d29190614391565b63ffffffff16106128f65760405163150358a160e21b815260040160405180910390fd5b5f80805b61290760408601866144e1565b90508110156129f8573661291e60408701876144e1565b8381811061292e5761292e613fe8565b6040029190910191505f90506129476020830183613adf565b6001600160a01b03160361296e57604051630863a45360e11b815260040160405180910390fd5b61297b6020820182613adf565b6001600160a01b0316836001600160a01b0316106129ac576040516310fb47f160e31b815260040160405180910390fd5b5f8160200135116129d0576040516310eb483f60e21b815260040160405180910390fd5b6129dd6020820182613adf565b92506129ed602082013585614267565b9350506001016128fa565b506f4b3b4ca85a86c47a098a223fffffffff821115612a2a5760405163070b5a6f60e21b815260040160405180910390fd5b5092915050565b6040516001600160a01b03808516602483015283166044820152606481018290526114319085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526136ae565b6033546001600160a01b031633146115535760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401611ed3565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b612b60612b4082806144e1565b612b506080850160608601613d3c565b61287460a0860160808701613d3c565b5f816040013511612b84576040516310eb483f60e21b815260040160405180910390fd5b6f4b3b4ca85a86c47a098a223fffffffff81604001351115612bb95760405163070b5a6f60e21b815260040160405180910390fd5b612be963ffffffff7f00000000000000000000000000000000000000000000000000000000000000001642614267565b612bf96080830160608401613d3c565b63ffffffff1611156114a357604051637ee2b44360e01b815260040160405180910390fd5b5f60ca612c2e6020850185613d3c565b63ffffffff1681548110612c4457612c44613fe8565b5f91825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff16151560608201529050612ca48382612f87565b5f612cb56080850160608601613adf565b6001600160a01b038082165f90815260cc60205260409020549192501680612cda5750805b336001600160a01b03821614612d0357604051635c427cd960e01b815260040160405180910390fd5b5f5b612d1260a0870187614526565b90508110156120125736612d2960e08801886144e1565b83818110612d3957612d39613fe8565b6001600160a01b0387165f90815260cd602090815260408083209302949094019450929091508290612d6d90850185613adf565b6001600160a01b03166001600160a01b031681526020019081526020015f2054905080826020013511612db35760405163aa385e8160e01b815260040160405180910390fd5b5f612dc282602085013561429f565b6001600160a01b0387165f90815260cd60209081526040822092935085018035929190612def9087613adf565b6001600160a01b031681526020808201929092526040015f2091909155612e309089908390612e2090870187613adf565b6001600160a01b03169190613781565b86516001600160a01b03808a1691878216918916907f9543dbd55580842586a951f0386e24d68a5df99ae29e3b216588b45fd684ce3190612e746020890189613adf565b604080519283526001600160a01b039091166020830152810186905260600160405180910390a4505050600101612d05565b5f816040015163ffffffff165f1480612ed85750815161ffff908116148015612ed85750816040015163ffffffff1642105b15612ef057505060cb54600160e01b900461ffff1690565b816040015163ffffffff16421015612f0957815161148c565b506020015190565b919050565b60cb546040805163ffffffff600160a01b9093048316815291831660208301527faf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3910160405180910390a160cb805463ffffffff909216600160a01b0263ffffffff60a01b19909216919091179055565b806060015115612faa57604051631b14174b60e01b815260040160405180910390fd5b806040015163ffffffff16421015612fd557604051631437a2bb60e31b815260040160405180910390fd5b612fe260c0830183614526565b9050612ff160a0840184614526565b905014613011576040516343714afd60e01b815260040160405180910390fd5b61301e60e08301836144e1565b905061302d60c0840184614526565b90501461304d576040516343714afd60e01b815260040160405180910390fd5b8051613079906130636040850160208601613d3c565b613070604086018661456b565b866060016137b1565b5f5b61308860a0840184614526565b90508110156111785761312260808401356130a660a0860186614526565b848181106130b6576130b6613fe8565b90506020020160208101906130cb9190613d3c565b6130d860c0870187614526565b858181106130e8576130e8613fe8565b90506020028101906130fa919061456b565b61310760e08901896144e1565b8781811061311757613117613fe8565b905060400201613855565b60010161307b565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60cb546040516001600160a01b038084169216907f237b82f438d75fc568ebab484b75b01d9287b9e98b490b7c23221623b6705dbb905f90a360cb80546001600160a01b0319166001600160a01b0392909216919091179055565b5f815f0151826020015163ffffffff1660405160200161322192919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261148c906145ad565b6001600160a01b038083165f81815260cc602052604080822080548686166001600160a01b0319821681179092559151919094169392849290917fbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca3129190a4505050565b60cb546040805161ffff600160e01b9093048316815291831660208301527fe6cd4edfdcc1f6d130ab35f73d72378f3a642944fb4ee5bd84b7807a81ea1c4e910160405180910390a160cb805461ffff909216600160e01b0261ffff60e01b19909216919091179055565b61271061ffff8316111561332e5760405163891c63df60e01b815260040160405180910390fd5b8254600160201b900463ffffffff16421161335c57604051637b1e25c560e01b815260040160405180910390fd5b8254600160201b900463ffffffff165f0361338357825461ffff191661ffff17835561339a565b825462010000810461ffff1661ffff199091161783555b825463ffffffff909116600160201b0267ffffffff000000001961ffff90931662010000029290921667ffffffffffff00001990911617179055565b826133f45760405163796cc52560e01b815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168163ffffffff16111561344157604051630dd0b9f560e21b815260040160405180910390fd5b61346b7f0000000000000000000000000000000000000000000000000000000000000000826145e4565b63ffffffff161561348f5760405163ee66470560e01b815260040160405180910390fd5b6134b97f0000000000000000000000000000000000000000000000000000000000000000836145e4565b63ffffffff16156134dd57604051633c1a94f160e21b815260040160405180910390fd5b8163ffffffff167f000000000000000000000000000000000000000000000000000000000000000063ffffffff1642613516919061429f565b1115801561355057508163ffffffff167f000000000000000000000000000000000000000000000000000000000000000063ffffffff1611155b61356d5760405163041aa75760e11b815260040160405180910390fd5b5f805b84811015612012575f86868381811061358b5761358b613fe8565b6135a19260206040909202019081019150613adf565b60405163198f077960e21b81526001600160a01b0380831660048301529192507f00000000000000000000000000000000000000000000000000000000000000009091169063663c1de490602401602060405180830381865afa15801561360a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061362e9190613fcd565b8061365557506001600160a01b03811673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0145b61367257604051632efd965160e11b815260040160405180910390fd5b806001600160a01b0316836001600160a01b0316106136a45760405163dfad9ca160e01b815260040160405180910390fd5b9150600101613570565b5f613702826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166138939092919063ffffffff16565b905080515f14806137225750808060200190518101906137229190613fcd565b6111785760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401611ed3565b6040516001600160a01b03831660248201526044810182905261117890849063a9059cbb60e01b90606401612a65565b6137bc60208361460b565b6001901b8463ffffffff16106137e45760405162c6c39d60e71b815260040160405180910390fd5b5f6137ee82610f79565b905061383884848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508a92508591505063ffffffff89166138a9565b612012576040516369ca16c960e01b815260040160405180910390fd5b61386060208361460b565b6001901b8463ffffffff16106138895760405163054ff4df60e51b815260040160405180910390fd5b5f6137ee82612200565b60606138a184845f856138c0565b949350505050565b5f836138b6868585613997565b1495945050505050565b6060824710156139215760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401611ed3565b5f5f866001600160a01b0316858760405161393c919061461e565b5f6040518083038185875af1925050503d805f8114613976576040519150601f19603f3d011682016040523d82523d5f602084013e61397b565b606091505b509150915061398c87838387613a2e565b979650505050505050565b5f602084516139a69190614634565b156139c4576040516313717da960e21b815260040160405180910390fd5b8260205b85518111613a25576139db600285614634565b5f036139fc57815f528086015160205260405f209150600284049350613a13565b808601515f528160205260405f2091506002840493505b613a1e602082614267565b90506139c8565b50949350505050565b60608315613a9c5782515f03613a95576001600160a01b0385163b613a955760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611ed3565b50816138a1565b6138a18383815115613ab15781518083602001fd5b8060405162461bcd60e51b8152600401611ed39190614647565b6001600160a01b03811681146114a3575f5ffd5b5f60208284031215613aef575f5ffd5b81356118b381613acb565b5f60408284031215613b0a575f5ffd5b50919050565b5f5f83601f840112613b20575f5ffd5b5081356001600160401b03811115613b36575f5ffd5b6020830191508360208260051b8501011115613b50575f5ffd5b9250929050565b5f5f5f60608486031215613b69575f5ffd5b613b738585613afa565b925060408401356001600160401b03811115613b8d575f5ffd5b613b9986828701613b10565b9497909650939450505050565b80151581146114a3575f5ffd5b5f5f60408385031215613bc4575f5ffd5b8235613bcf81613acb565b91506020830135613bdf81613ba6565b809150509250929050565b5f60208284031215613bfa575f5ffd5b5035919050565b5f60408284031215613c11575f5ffd5b6118b38383613afa565b5f5f60208385031215613c2c575f5ffd5b82356001600160401b03811115613c41575f5ffd5b613c4d85828601613b10565b90969095509350505050565b5f6101008284031215613b0a575f5ffd5b5f5f60408385031215613c7b575f5ffd5b82356001600160401b03811115613c90575f5ffd5b613c9c85828601613c59565b9250506020830135613bdf81613acb565b803563ffffffff81168114612f11575f5ffd5b5f5f60408385031215613cd1575f5ffd5b82359150613ce160208401613cad565b90509250929050565b5f5f5f60408486031215613cfc575f5ffd5b83356001600160401b03811115613d11575f5ffd5b613d1d86828701613b10565b9094509250506020840135613d3181613acb565b809150509250925092565b5f60208284031215613d4c575f5ffd5b6118b382613cad565b5f60208284031215613d65575f5ffd5b813560ff811681146118b3575f5ffd5b5f60208284031215613d85575f5ffd5b81356001600160401b03811115613d9a575f5ffd5b6138a184828501613c59565b5f5f60408385031215613db7575f5ffd5b8235613dc281613acb565b946020939093013593505050565b5f5f60408385031215613de1575f5ffd5b8235613dec81613acb565b91506020830135613bdf81613acb565b5f5f5f60408486031215613e0e575f5ffd5b8335613e1981613acb565b925060208401356001600160401b03811115613b8d575f5ffd5b5f5f60608385031215613e44575f5ffd5b8235613e4f81613acb565b9150613ce18460208501613afa565b803561ffff81168114612f11575f5ffd5b5f60208284031215613e7f575f5ffd5b6118b382613e5e565b5f5f60408385031215613e99575f5ffd5b8235613ea481613acb565b9150613ce160208401613e5e565b5f5f5f60608486031215613ec4575f5ffd5b8335613ecf81613acb565b92506020840135613edf81613acb565b9150613eed60408501613e5e565b90509250925092565b5f5f5f5f5f60a08688031215613f0a575f5ffd5b8535613f1581613acb565b9450602086013593506040860135613f2c81613acb565b9250613f3a60608701613cad565b9150613f4860808701613e5e565b90509295509295909350565b5f5f5f60808486031215613f66575f5ffd5b8335613f7181613acb565b9250613f808560208601613afa565b9150613eed60608501613e5e565b8035613f9981613acb565b6001600160a01b0316825263ffffffff613fb560208301613cad565b1660208301525050565b6040810161148c8284613f8e565b5f60208284031215613fdd575f5ffd5b81516118b381613ba6565b634e487b7160e01b5f52603260045260245ffd5b5f823560be19833603018112614010575f5ffd5b9190910192915050565b5f5f8335601e1984360301811261402f575f5ffd5b83016020810192503590506001600160401b0381111561404d575f5ffd5b8060061b3603821315613b50575f5ffd5b8183526020830192505f815f5b848110156140c157813561407e81613acb565b6001600160a01b0316865260208201356bffffffffffffffffffffffff81168082146140a8575f5ffd5b602088015250604095860195919091019060010161406b565b5093949350505050565b5f5f8335601e198436030181126140e0575f5ffd5b83016020810192503590506001600160401b038111156140fe575f5ffd5b803603821315613b50575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f61413f828361401a565b60c0855261415160c08601828461405e565b915050602083013561416281613acb565b6001600160a01b0316602085015261417d604084018461401a565b858303604087015280835290915f91906020015b818310156141cc5783356141a481613acb565b6001600160a01b03168152602084810135908201526040938401936001939093019201614191565b6141d860608701613cad565b63ffffffff8116606089015293506141f260808701613cad565b63ffffffff81166080890152935061420d60a08701876140cb565b9450925086810360a088015261398c81858561410c565b60018060a01b0384168152826020820152606060408201525f61424a6060830184614134565b95945050505050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561148c5761148c614253565b6142848185613f8e565b826040820152608060608201525f61424a6080830184614134565b8181038181111561148c5761148c614253565b5f816142c0576142c0614253565b505f190190565b5f8235609e19833603018112614010575f5ffd5b5f6142e6828361401a565b60a085526142f860a08601828461405e565b915050602083013561430981613acb565b6001600160a01b031660208501526040838101359085015263ffffffff61433260608501613cad565b16606085015263ffffffff61434960808501613cad565b1660808501528091505092915050565b60018060a01b0384168152826020820152606060408201525f61424a60608301846142db565b602081525f6118b360208301846142db565b63ffffffff818116838216019081111561148c5761148c614253565b5f823560fe19833603018112614010575f5ffd5b828152604060208201525f6138a16040830184614134565b5f60408284031280156143ea575f5ffd5b50604080519081016001600160401b038111828210171561441957634e487b7160e01b5f52604160045260245ffd5b604052823561442781613acb565b815261443560208401613cad565b60208201529392505050565b63ffffffff828116828216039081111561148c5761148c614253565b5f63ffffffff82168061447257614472614253565b5f190192915050565b5f6020828403121561448b575f5ffd5b5051919050565b60a081016144a08287613f8e565b63ffffffff94909416604082015261ffff92831660608201529116608090910152919050565b5f602082840312156144d6575f5ffd5b81516118b381613acb565b5f5f8335601e198436030181126144f6575f5ffd5b8301803591506001600160401b0382111561450f575f5ffd5b6020019150600681901b3603821315613b50575f5ffd5b5f5f8335601e1984360301811261453b575f5ffd5b8301803591506001600160401b03821115614554575f5ffd5b6020019150600581901b3603821315613b50575f5ffd5b5f5f8335601e19843603018112614580575f5ffd5b8301803591506001600160401b03821115614599575f5ffd5b602001915036819003821315613b50575f5ffd5b80516020808301519190811015613b0a575f1960209190910360031b1b16919050565b634e487b7160e01b5f52601260045260245ffd5b5f63ffffffff8316806145f9576145f96145d0565b8063ffffffff84160691505092915050565b5f82614619576146196145d0565b500490565b5f82518060208501845e5f920191825250919050565b5f82614642576146426145d0565b500690565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f8301168401019150509291505056fea26469706673582212204175e2e90f0c601bffe8c289e572af14203425a6a08e9c100886005e1cd5f2a164736f6c634300081b0033",
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

// GetOperatorSetSplit is a free data retrieval call binding the contract method 0x9de4b35f.
//
// Solidity: function getOperatorSetSplit(address operator, (address,uint32) operatorSet) view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorCaller) GetOperatorSetSplit(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (uint16, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "getOperatorSetSplit", operator, operatorSet)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetOperatorSetSplit is a free data retrieval call binding the contract method 0x9de4b35f.
//
// Solidity: function getOperatorSetSplit(address operator, (address,uint32) operatorSet) view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorSession) GetOperatorSetSplit(operator common.Address, operatorSet OperatorSet) (uint16, error) {
	return _RewardsCoordinator.Contract.GetOperatorSetSplit(&_RewardsCoordinator.CallOpts, operator, operatorSet)
}

// GetOperatorSetSplit is a free data retrieval call binding the contract method 0x9de4b35f.
//
// Solidity: function getOperatorSetSplit(address operator, (address,uint32) operatorSet) view returns(uint16)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) GetOperatorSetSplit(operator common.Address, operatorSet OperatorSet) (uint16, error) {
	return _RewardsCoordinator.Contract.GetOperatorSetSplit(&_RewardsCoordinator.CallOpts, operator, operatorSet)
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
// Solidity: function isOperatorDirectedAVSRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorCaller) IsOperatorDirectedAVSRewardsSubmissionHash(opts *bind.CallOpts, avs common.Address, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "isOperatorDirectedAVSRewardsSubmissionHash", avs, hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorDirectedAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0xed71e6a2.
//
// Solidity: function isOperatorDirectedAVSRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorSession) IsOperatorDirectedAVSRewardsSubmissionHash(avs common.Address, hash [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsOperatorDirectedAVSRewardsSubmissionHash(&_RewardsCoordinator.CallOpts, avs, hash)
}

// IsOperatorDirectedAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0xed71e6a2.
//
// Solidity: function isOperatorDirectedAVSRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) IsOperatorDirectedAVSRewardsSubmissionHash(avs common.Address, hash [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsOperatorDirectedAVSRewardsSubmissionHash(&_RewardsCoordinator.CallOpts, avs, hash)
}

// IsOperatorDirectedOperatorSetRewardsSubmissionHash is a free data retrieval call binding the contract method 0xf2f07ab4.
//
// Solidity: function isOperatorDirectedOperatorSetRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorCaller) IsOperatorDirectedOperatorSetRewardsSubmissionHash(opts *bind.CallOpts, avs common.Address, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "isOperatorDirectedOperatorSetRewardsSubmissionHash", avs, hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorDirectedOperatorSetRewardsSubmissionHash is a free data retrieval call binding the contract method 0xf2f07ab4.
//
// Solidity: function isOperatorDirectedOperatorSetRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorSession) IsOperatorDirectedOperatorSetRewardsSubmissionHash(avs common.Address, hash [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsOperatorDirectedOperatorSetRewardsSubmissionHash(&_RewardsCoordinator.CallOpts, avs, hash)
}

// IsOperatorDirectedOperatorSetRewardsSubmissionHash is a free data retrieval call binding the contract method 0xf2f07ab4.
//
// Solidity: function isOperatorDirectedOperatorSetRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) IsOperatorDirectedOperatorSetRewardsSubmissionHash(avs common.Address, hash [32]byte) (bool, error) {
	return _RewardsCoordinator.Contract.IsOperatorDirectedOperatorSetRewardsSubmissionHash(&_RewardsCoordinator.CallOpts, avs, hash)
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

// CreateOperatorDirectedOperatorSetRewardsSubmission is a paid mutator transaction binding the contract method 0x0ca29899.
//
// Solidity: function createOperatorDirectedOperatorSetRewardsSubmission((address,uint32) operatorSet, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) CreateOperatorDirectedOperatorSetRewardsSubmission(opts *bind.TransactOpts, operatorSet OperatorSet, operatorDirectedRewardsSubmissions []IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "createOperatorDirectedOperatorSetRewardsSubmission", operatorSet, operatorDirectedRewardsSubmissions)
}

// CreateOperatorDirectedOperatorSetRewardsSubmission is a paid mutator transaction binding the contract method 0x0ca29899.
//
// Solidity: function createOperatorDirectedOperatorSetRewardsSubmission((address,uint32) operatorSet, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) CreateOperatorDirectedOperatorSetRewardsSubmission(operatorSet OperatorSet, operatorDirectedRewardsSubmissions []IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.CreateOperatorDirectedOperatorSetRewardsSubmission(&_RewardsCoordinator.TransactOpts, operatorSet, operatorDirectedRewardsSubmissions)
}

// CreateOperatorDirectedOperatorSetRewardsSubmission is a paid mutator transaction binding the contract method 0x0ca29899.
//
// Solidity: function createOperatorDirectedOperatorSetRewardsSubmission((address,uint32) operatorSet, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) CreateOperatorDirectedOperatorSetRewardsSubmission(operatorSet OperatorSet, operatorDirectedRewardsSubmissions []IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.CreateOperatorDirectedOperatorSetRewardsSubmission(&_RewardsCoordinator.TransactOpts, operatorSet, operatorDirectedRewardsSubmissions)
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

// SetOperatorSetSplit is a paid mutator transaction binding the contract method 0xf74e8eac.
//
// Solidity: function setOperatorSetSplit(address operator, (address,uint32) operatorSet, uint16 split) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactor) SetOperatorSetSplit(opts *bind.TransactOpts, operator common.Address, operatorSet OperatorSet, split uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.contract.Transact(opts, "setOperatorSetSplit", operator, operatorSet, split)
}

// SetOperatorSetSplit is a paid mutator transaction binding the contract method 0xf74e8eac.
//
// Solidity: function setOperatorSetSplit(address operator, (address,uint32) operatorSet, uint16 split) returns()
func (_RewardsCoordinator *RewardsCoordinatorSession) SetOperatorSetSplit(operator common.Address, operatorSet OperatorSet, split uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetOperatorSetSplit(&_RewardsCoordinator.TransactOpts, operator, operatorSet, split)
}

// SetOperatorSetSplit is a paid mutator transaction binding the contract method 0xf74e8eac.
//
// Solidity: function setOperatorSetSplit(address operator, (address,uint32) operatorSet, uint16 split) returns()
func (_RewardsCoordinator *RewardsCoordinatorTransactorSession) SetOperatorSetSplit(operator common.Address, operatorSet OperatorSet, split uint16) (*types.Transaction, error) {
	return _RewardsCoordinator.Contract.SetOperatorSetSplit(&_RewardsCoordinator.TransactOpts, operator, operatorSet, split)
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

// RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreatedIterator is returned from FilterOperatorDirectedOperatorSetRewardsSubmissionCreated and is used to iterate over the raw logs and unpacked data for OperatorDirectedOperatorSetRewardsSubmissionCreated events raised by the RewardsCoordinator contract.
type RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreatedIterator struct {
	Event *RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated)
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
		it.Event = new(RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated)
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
func (it *RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated represents a OperatorDirectedOperatorSetRewardsSubmissionCreated event raised by the RewardsCoordinator contract.
type RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated struct {
	Caller                                common.Address
	OperatorDirectedRewardsSubmissionHash [32]byte
	OperatorSet                           OperatorSet
	SubmissionNonce                       *big.Int
	OperatorDirectedRewardsSubmission     IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission
	Raw                                   types.Log // Blockchain specific contextual infos
}

// FilterOperatorDirectedOperatorSetRewardsSubmissionCreated is a free log retrieval operation binding the contract event 0xfff0759ccb371dfb5691798724e70b4fa61cb3bfe730a33ac19fb86a48efc756.
//
// Solidity: event OperatorDirectedOperatorSetRewardsSubmissionCreated(address indexed caller, bytes32 indexed operatorDirectedRewardsSubmissionHash, (address,uint32) operatorSet, uint256 submissionNonce, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string) operatorDirectedRewardsSubmission)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterOperatorDirectedOperatorSetRewardsSubmissionCreated(opts *bind.FilterOpts, caller []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (*RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreatedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorDirectedRewardsSubmissionHashRule []interface{}
	for _, operatorDirectedRewardsSubmissionHashItem := range operatorDirectedRewardsSubmissionHash {
		operatorDirectedRewardsSubmissionHashRule = append(operatorDirectedRewardsSubmissionHashRule, operatorDirectedRewardsSubmissionHashItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "OperatorDirectedOperatorSetRewardsSubmissionCreated", callerRule, operatorDirectedRewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreatedIterator{contract: _RewardsCoordinator.contract, event: "OperatorDirectedOperatorSetRewardsSubmissionCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorDirectedOperatorSetRewardsSubmissionCreated is a free log subscription operation binding the contract event 0xfff0759ccb371dfb5691798724e70b4fa61cb3bfe730a33ac19fb86a48efc756.
//
// Solidity: event OperatorDirectedOperatorSetRewardsSubmissionCreated(address indexed caller, bytes32 indexed operatorDirectedRewardsSubmissionHash, (address,uint32) operatorSet, uint256 submissionNonce, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string) operatorDirectedRewardsSubmission)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchOperatorDirectedOperatorSetRewardsSubmissionCreated(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated, caller []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorDirectedRewardsSubmissionHashRule []interface{}
	for _, operatorDirectedRewardsSubmissionHashItem := range operatorDirectedRewardsSubmissionHash {
		operatorDirectedRewardsSubmissionHashRule = append(operatorDirectedRewardsSubmissionHashRule, operatorDirectedRewardsSubmissionHashItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "OperatorDirectedOperatorSetRewardsSubmissionCreated", callerRule, operatorDirectedRewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "OperatorDirectedOperatorSetRewardsSubmissionCreated", log); err != nil {
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

// ParseOperatorDirectedOperatorSetRewardsSubmissionCreated is a log parse operation binding the contract event 0xfff0759ccb371dfb5691798724e70b4fa61cb3bfe730a33ac19fb86a48efc756.
//
// Solidity: event OperatorDirectedOperatorSetRewardsSubmissionCreated(address indexed caller, bytes32 indexed operatorDirectedRewardsSubmissionHash, (address,uint32) operatorSet, uint256 submissionNonce, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string) operatorDirectedRewardsSubmission)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseOperatorDirectedOperatorSetRewardsSubmissionCreated(log types.Log) (*RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated, error) {
	event := new(RewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "OperatorDirectedOperatorSetRewardsSubmissionCreated", log); err != nil {
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

// RewardsCoordinatorOperatorSetSplitBipsSetIterator is returned from FilterOperatorSetSplitBipsSet and is used to iterate over the raw logs and unpacked data for OperatorSetSplitBipsSet events raised by the RewardsCoordinator contract.
type RewardsCoordinatorOperatorSetSplitBipsSetIterator struct {
	Event *RewardsCoordinatorOperatorSetSplitBipsSet // Event containing the contract specifics and raw log

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
func (it *RewardsCoordinatorOperatorSetSplitBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsCoordinatorOperatorSetSplitBipsSet)
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
		it.Event = new(RewardsCoordinatorOperatorSetSplitBipsSet)
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
func (it *RewardsCoordinatorOperatorSetSplitBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsCoordinatorOperatorSetSplitBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsCoordinatorOperatorSetSplitBipsSet represents a OperatorSetSplitBipsSet event raised by the RewardsCoordinator contract.
type RewardsCoordinatorOperatorSetSplitBipsSet struct {
	Caller                  common.Address
	Operator                common.Address
	OperatorSet             OperatorSet
	ActivatedAt             uint32
	OldOperatorSetSplitBips uint16
	NewOperatorSetSplitBips uint16
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetSplitBipsSet is a free log retrieval operation binding the contract event 0x14918b3834ab6752eb2e1b489b6663a67810efb5f56f3944a97ede8ecf1fd9f1.
//
// Solidity: event OperatorSetSplitBipsSet(address indexed caller, address indexed operator, (address,uint32) operatorSet, uint32 activatedAt, uint16 oldOperatorSetSplitBips, uint16 newOperatorSetSplitBips)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) FilterOperatorSetSplitBipsSet(opts *bind.FilterOpts, caller []common.Address, operator []common.Address) (*RewardsCoordinatorOperatorSetSplitBipsSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.FilterLogs(opts, "OperatorSetSplitBipsSet", callerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &RewardsCoordinatorOperatorSetSplitBipsSetIterator{contract: _RewardsCoordinator.contract, event: "OperatorSetSplitBipsSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSetSplitBipsSet is a free log subscription operation binding the contract event 0x14918b3834ab6752eb2e1b489b6663a67810efb5f56f3944a97ede8ecf1fd9f1.
//
// Solidity: event OperatorSetSplitBipsSet(address indexed caller, address indexed operator, (address,uint32) operatorSet, uint32 activatedAt, uint16 oldOperatorSetSplitBips, uint16 newOperatorSetSplitBips)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) WatchOperatorSetSplitBipsSet(opts *bind.WatchOpts, sink chan<- *RewardsCoordinatorOperatorSetSplitBipsSet, caller []common.Address, operator []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RewardsCoordinator.contract.WatchLogs(opts, "OperatorSetSplitBipsSet", callerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsCoordinatorOperatorSetSplitBipsSet)
				if err := _RewardsCoordinator.contract.UnpackLog(event, "OperatorSetSplitBipsSet", log); err != nil {
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

// ParseOperatorSetSplitBipsSet is a log parse operation binding the contract event 0x14918b3834ab6752eb2e1b489b6663a67810efb5f56f3944a97ede8ecf1fd9f1.
//
// Solidity: event OperatorSetSplitBipsSet(address indexed caller, address indexed operator, (address,uint32) operatorSet, uint32 activatedAt, uint16 oldOperatorSetSplitBips, uint16 newOperatorSetSplitBips)
func (_RewardsCoordinator *RewardsCoordinatorFilterer) ParseOperatorSetSplitBipsSet(log types.Log) (*RewardsCoordinatorOperatorSetSplitBipsSet, error) {
	event := new(RewardsCoordinatorOperatorSetSplitBipsSet)
	if err := _RewardsCoordinator.contract.UnpackLog(event, "OperatorSetSplitBipsSet", log); err != nil {
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
