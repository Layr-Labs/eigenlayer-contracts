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

// IRewardsCoordinatorTypesRewardsCoordinatorConstructorParams is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesRewardsCoordinatorConstructorParams struct {
	DelegationManager          common.Address
	StrategyManager            common.Address
	AllocationManager          common.Address
	PauserRegistry             common.Address
	PermissionController       common.Address
	CALCULATIONINTERVALSECONDS uint32
	MAXREWARDSDURATION         uint32
	MAXRETROACTIVELENGTH       uint32
	MAXFUTURELENGTH            uint32
	GENESISREWARDSTIMESTAMP    uint32
	Version                    string
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsCoordinatorConstructorParams\",\"components\":[{\"name\":\"delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"CALCULATION_INTERVAL_SECONDS\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"MAX_REWARDS_DURATION\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"MAX_RETROACTIVE_LENGTH\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"MAX_FUTURE_LENGTH\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"GENESIS_REWARDS_TIMESTAMP\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CALCULATION_INTERVAL_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GENESIS_REWARDS_TIMESTAMP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FUTURE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_RETROACTIVE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_REWARDS_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activationDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateEarnerLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"calculateTokenLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"checkClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimerFor\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAVSRewardsSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorDirectedAVSRewardsSubmission\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorDirectedRewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorDirectedOperatorSetRewardsSubmission\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorDirectedRewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRewardsForAllEarners\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRewardsForAllSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cumulativeClaimed\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"totalClaimed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currRewardsCalculationEndTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultOperatorSplitBips\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableRoot\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentClaimableDistributionRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentDistributionRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributionRootAtIndex\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributionRootsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorAVSSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorPISplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRootIndexFromHash\",\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_rewardsUpdater\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_defaultSplitBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAVSRewardsSubmissionHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorDirectedAVSRewardsSubmissionHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorDirectedOperatorSetRewardsSubmissionHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRewardsForAllSubmitter\",\"inputs\":[{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRewardsSubmissionForAllEarnersHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRewardsSubmissionForAllHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"processClaims\",\"inputs\":[{\"name\":\"claims\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsMerkleClaim[]\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardsUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setActivationDelay\",\"inputs\":[{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimerFor\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimerFor\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultOperatorSplit\",\"inputs\":[{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorAVSSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorPISplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsForAllSubmitter\",\"inputs\":[{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_newValue\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsUpdater\",\"inputs\":[{\"name\":\"_rewardsUpdater\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submissionNonce\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AVSRewardsSubmissionCreated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ActivationDelaySet\",\"inputs\":[{\"name\":\"oldActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimerForSet\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldClaimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultOperatorSplitBipsSet\",\"inputs\":[{\"name\":\"oldDefaultOperatorSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newDefaultOperatorSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootDisabled\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootSubmitted\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSSplitBipsSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"oldOperatorAVSSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newOperatorAVSSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDirectedAVSRewardsSubmissionCreated\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorDirectedRewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"operatorDirectedRewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDirectedOperatorSetRewardsSubmissionCreated\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorDirectedRewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"operatorDirectedRewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorPISplitBipsSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"oldOperatorPISplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newOperatorPISplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetSplitBipsSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"oldOperatorSetSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newOperatorSetSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsClaimed\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"claimedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsForAllSubmitterSet\",\"inputs\":[{\"name\":\"rewardsForAllSubmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"newValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsSubmissionForAllCreated\",\"inputs\":[{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsSubmissionForAllEarnersCreated\",\"inputs\":[{\"name\":\"tokenHopper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsUpdaterSet\",\"inputs\":[{\"name\":\"oldRewardsUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newRewardsUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AmountExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AmountIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DurationExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DurationIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EarningsNotGreaterThanClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCalculationIntervalSecondsRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidClaimProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDurationRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEarner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEarnerLeafIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGenesisRewardsTimestampRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRootIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidStartTimestampRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTokenLeafIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewRootMustBeForNewCalculatedPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorsNotInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PreviousSplitPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardsEndTimestampNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootAlreadyActivated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootNotActivated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SplitExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartTimestampTooFarInFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartTimestampTooFarInPast\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategiesNotInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"SubmissionNotRetroactive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[]}]",
	Bin: "0x6101e0604052348015610010575f5ffd5b50604051614b90380380614b9083398101604081905261002f9161034a565b610140810151608082015182516020840151604085015160a086015160c087015160e08801516101008901516101208a015160608b01516001600160a01b03811661008d576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b03166080526100a38582610465565b63ffffffff16156100c757604051630e06bd3160e01b815260040160405180910390fd5b6100d46201518086610465565b63ffffffff16156100f85760405163223c7b3960e11b815260040160405180910390fd5b6001600160a01b0397881660a05295871660c05293861660e05263ffffffff9283166101005290821661012052811661014052908116610160521661018052166101a05261014581610158565b6101c0525061015261019e565b506104f3565b5f5f829050601f8151111561018b578260405163305a27a960e01b81526004016101829190610498565b60405180910390fd5b8051610196826104cd565b179392505050565b5f54610100900460ff16156102055760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610182565b5f5460ff90811614610254575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b634e487b7160e01b5f52604160045260245ffd5b60405161016081016001600160401b038111828210171561028d5761028d610256565b60405290565b80516001600160a01b03811681146102a9575f5ffd5b919050565b805163ffffffff811681146102a9575f5ffd5b5f82601f8301126102d0575f5ffd5b81516001600160401b038111156102e9576102e9610256565b604051601f8201601f19908116603f011681016001600160401b038111828210171561031757610317610256565b60405281815283820160200185101561032e575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f6020828403121561035a575f5ffd5b81516001600160401b0381111561036f575f5ffd5b82016101608185031215610381575f5ffd5b61038961026a565b61039282610293565b81526103a060208301610293565b60208201526103b160408301610293565b60408201526103c260608301610293565b60608201526103d360808301610293565b60808201526103e460a083016102ae565b60a08201526103f560c083016102ae565b60c082015261040660e083016102ae565b60e082015261041861010083016102ae565b61010082015261042b61012083016102ae565b6101208201526101408201516001600160401b0381111561044a575f5ffd5b610456868285016102c1565b61014083015250949350505050565b5f63ffffffff83168061048657634e487b7160e01b5f52601260045260245ffd5b8063ffffffff84160691505092915050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156104ed575f198160200360031b1b821691505b50919050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516101c0516145c26105ce5f395f61138701525f81816105f001526124b301525f81816104a4015261341101525f81816103f3015261299801525f818161055201526133cf01525f818161087e01526132b901525f81816107bf01528181613309015261337d01525f81816108d201528181610adc01528181611bff0152611e8c01525f818161057901526134ac01525f81816109450152611b6f01525f81816107510152818161283d01526131ea01526145c25ff3fe608060405234801561000f575f5ffd5b50600436106103b3575f3560e01c8063886f1195116101f5578063de02e50311610114578063f6efbb59116100a9578063fabc1cbc11610079578063fabc1cbc14610a33578063fbf1e2c114610a46578063fce36c7d14610a59578063ff9f6cce14610a6c575f5ffd5b8063f6efbb59146109e7578063f74e8eac146109fa578063f8cd844814610a0d578063f96abf2e14610a20575f5ffd5b8063ed71e6a2116100e4578063ed71e6a214610967578063f22cef8514610994578063f2f07ab4146109a7578063f2fde38b146109d4575f5ffd5b8063de02e50314610907578063e063f81f1461091a578063e810ce211461092d578063ea4d3c9b14610940575f5ffd5b8063a50a1d9c1161018a578063bf21a8aa1161015a578063bf21a8aa14610879578063c46db606146108a0578063ca8aa7c7146108cd578063dcbb03b3146108f4575f5ffd5b8063a50a1d9c14610807578063aebd8bae1461081a578063b3dbb0e014610847578063bb7e451f1461085a575f5ffd5b80639cb9a5fa116101c55780639cb9a5fa146107a75780639d45c281146107ba5780639de4b35f146107e1578063a0169ddd146107f4575f5ffd5b8063886f11951461074c5780638da5cb5b146107735780639104c319146107845780639be3d4e41461079f575f5ffd5b80634596021c116102e15780635c975abb11610276578063715018a611610246578063715018a6146106ff5780637b8f8b0514610707578063863cb9a91461070f578063865c695314610722575f5ffd5b80635c975abb146106a25780635e9d8348146106aa57806363f6a798146106bd5780636d21117e146106d2575f5ffd5b806354fd4d50116102b157806354fd4d501461064f57806358baaa3e14610664578063595c6a67146106775780635ac86ab71461067f575f5ffd5b80634596021c146105d85780634657e26a146105eb5780634b943960146106125780634d18cc3514610638575f5ffd5b8063149bc8721161035757806339b70e381161032757806339b70e38146105745780633a8c07861461059b5780633ccc861d146105b25780633efe1db6146105c5575f5ffd5b8063149bc872146104d95780632b9f64a4146104fa57806336af41fa1461053a57806337838ed01461054d575f5ffd5b80630e9a53cf116103925780630e9a53cf1461043f5780630eb383451461048c578063131433b41461049f578063136439dd146104c6575f5ffd5b806218572c146103b757806304a0c502146103ee5780630ca298991461042a575b5f5ffd5b6103d96103c53660046139ef565b60d16020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b6104157f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016103e5565b61043d610438366004613a67565b610a7f565b005b610447610d25565b6040516103e591905f6080820190508251825263ffffffff602084015116602083015263ffffffff604084015116604083015260608301511515606083015292915050565b61043d61049a366004613ac3565b610e25565b6104157f000000000000000000000000000000000000000000000000000000000000000081565b61043d6104d4366004613afa565b610ea5565b6104ec6104e7366004613b11565b610edf565b6040519081526020016103e5565b6105226105083660046139ef565b60cc6020525f90815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020016103e5565b61043d610548366004613b2b565b610f54565b6104157f000000000000000000000000000000000000000000000000000000000000000081565b6105227f000000000000000000000000000000000000000000000000000000000000000081565b60cb5461041590600160a01b900463ffffffff1681565b61043d6105c0366004613b7a565b6110c5565b61043d6105d3366004613bd0565b6110ec565b61043d6105e6366004613bfa565b6112c2565b6105227f000000000000000000000000000000000000000000000000000000000000000081565b6106256106203660046139ef565b611325565b60405161ffff90911681526020016103e5565b60cb5461041590600160c01b900463ffffffff1681565b610657611380565b6040516103e59190613c4c565b61043d610672366004613c81565b6113b0565b61043d6113c4565b6103d961068d366004613c9a565b606654600160ff9092169190911b9081161490565b6066546104ec565b6103d96106b8366004613cba565b6113d8565b60cb5461062590600160e01b900461ffff1681565b6103d96106e0366004613ceb565b60cf60209081525f928352604080842090915290825290205460ff1681565b61043d611463565b60ca546104ec565b61043d61071d3660046139ef565b611474565b6104ec610730366004613d15565b60cd60209081525f928352604080842090915290825290205481565b6105227f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b0316610522565b61052273beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b610447611485565b61043d6107b5366004613d41565b611521565b6104157f000000000000000000000000000000000000000000000000000000000000000081565b6106256107ef366004613d78565b61169c565b61043d6108023660046139ef565b61171f565b61043d610815366004613db4565b61172a565b6103d9610828366004613ceb565b60d260209081525f928352604080842090915290825290205460ff1681565b61043d610855366004613dcd565b61173b565b6104ec6108683660046139ef565b60ce6020525f908152604090205481565b6104157f000000000000000000000000000000000000000000000000000000000000000081565b6103d96108ae366004613ceb565b60d060209081525f928352604080842090915290825290205460ff1681565b6105227f000000000000000000000000000000000000000000000000000000000000000081565b61043d610902366004613df7565b611866565b610447610915366004613afa565b6119b3565b610625610928366004613d15565b611a43565b61041561093b366004613afa565b611aa8565b6105227f000000000000000000000000000000000000000000000000000000000000000081565b6103d9610975366004613ceb565b60d360209081525f928352604080842090915290825290205460ff1681565b61043d6109a2366004613d15565b611b29565b6103d96109b5366004613ceb565b60d760209081525f928352604080842090915290825290205460ff1681565b61043d6109e23660046139ef565b611c93565b61043d6109f5366004613e3b565b611d0e565b61043d610a08366004613e99565b611e43565b6104ec610a1b366004613b11565b61200a565b61043d610a2e366004613c81565b61201a565b61043d610a41366004613afa565b61214b565b60cb54610522906001600160a01b031681565b61043d610a67366004613b2b565b6121b8565b61043d610a7a366004613b2b565b6122e9565b6009610a8a8161244a565b610a9760208501856139ef565b610aa081612475565b610abd5760405163932d94f760e01b815260040160405180910390fd5b610ac561251f565b6040516304c1b8eb60e31b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc75890610b11908890600401613f04565b602060405180830381865afa158015610b2c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b509190613f12565b610b6d57604051631fb1705560e21b815260040160405180910390fd5b5f5b83811015610d135736858583818110610b8a57610b8a613f2d565b9050602002810190610b9c9190613f41565b90505f60ce81610baf60208b018b6139ef565b6001600160a01b031681526020808201929092526040015f90812054925090610bda908a018a6139ef565b8284604051602001610bee93929190614169565b6040516020818303038152906040528051906020012090505f610c1084612578565b9050600160d75f610c2460208e018e6139ef565b6001600160a01b0316815260208082019290925260409081015f9081208682529092529020805460ff1916911515919091179055610c638360016141ac565b60ce5f610c7360208e018e6139ef565b6001600160a01b03166001600160a01b031681526020019081526020015f208190555081336001600160a01b03167ffff0759ccb371dfb5691798724e70b4fa61cb3bfe730a33ac19fb86a48efc7568c8688604051610cd4939291906141bf565b60405180910390a3610d03333083610cf26040890160208a016139ef565b6001600160a01b0316929190612763565b505060019092019150610b6f9050565b50610d1e6001609755565b5050505050565b604080516080810182525f80825260208201819052918101829052606081019190915260ca545b8015610dfd575f60ca610d606001846141e4565b81548110610d7057610d70613f2d565b5f91825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff161580156060830181905291925090610ddf5750806040015163ffffffff164210155b15610dea5792915050565b5080610df5816141f7565b915050610d4c565b5050604080516080810182525f80825260208201819052918101829052606081019190915290565b610e2d6127ce565b6001600160a01b0382165f81815260d1602052604080822054905160ff9091169284151592841515927f4de6293e668df1398422e1def12118052c1539a03cbfedc145895d48d7685f1c9190a4506001600160a01b03919091165f90815260d160205260409020805460ff1916911515919091179055565b610ead612828565b6066548181168114610ed25760405163c61dca5d60e01b815260040160405180910390fd5b610edb826128cb565b5050565b5f80610eee60208401846139ef565b8360200135604051602001610f379392919060f89390931b6001600160f81b031916835260609190911b6bffffffffffffffffffffffff19166001830152601582015260350190565b604051602081830303815290604052805190602001209050919050565b6001610f5f8161244a565b335f90815260d1602052604090205460ff16610f8e57604051635c427cd960e01b815260040160405180910390fd5b610f9661251f565b5f5b828110156110b55736848483818110610fb357610fb3613f2d565b9050602002810190610fc5919061420c565b335f81815260ce60209081526040808320549051949550939192610fef929091859187910161429e565b60405160208183030381529060405280519060200120905061101083612908565b335f90815260d0602090815260408083208484529091529020805460ff191660019081179091556110429083906141ac565b335f81815260ce602052604090819020929092559051829184917f51088b8c89628df3a8174002c2a034d0152fce6af8415d651b2a4734bf270482906110899088906142c4565b60405180910390a46110aa333060408601803590610cf290602089016139ef565b505050600101610f98565b506110c06001609755565b505050565b60026110d08161244a565b6110d861251f565b6110e283836129f3565b6110c06001609755565b60036110f78161244a565b60cb546001600160a01b0316331461112257604051635c427cd960e01b815260040160405180910390fd5b60cb5463ffffffff600160c01b90910481169083161161115557604051631ca7e50b60e21b815260040160405180910390fd5b428263ffffffff161061117b576040516306957c9160e11b815260040160405180910390fd5b60ca5460cb545f9061119a90600160a01b900463ffffffff16426142d6565b6040805160808101825287815263ffffffff87811660208084018281528684168587018181525f6060880181815260ca8054600181018255925297517f42d72674974f694b5f5159593243114d38a5c39c89d6b62fee061ff523240ee160029092029182015592517f42d72674974f694b5f5159593243114d38a5c39c89d6b62fee061ff523240ee290930180549151975193871667ffffffffffffffff1990921691909117600160201b978716979097029690961760ff60401b1916600160401b921515929092029190911790945560cb805463ffffffff60c01b1916600160c01b840217905593519283529394508892908616917fecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08910160405180910390a45050505050565b60026112cd8161244a565b6112d561251f565b5f5b838110156113145761130c8585838181106112f4576112f4613f2d565b905060200281019061130691906142f2565b846129f3565b6001016112d7565b5061131f6001609755565b50505050565b6001600160a01b0381165f90815260d5602090815260408083208151606081018352905461ffff80821683526201000082041693820193909352600160201b90920463ffffffff169082015261137a90612c7b565b92915050565b60606113ab7f0000000000000000000000000000000000000000000000000000000000000000612ceb565b905090565b6113b86127ce565b6113c181612d28565b50565b6113cc612828565b6113d65f196128cb565b565b5f61145b8260ca6113ec6020830183613c81565b63ffffffff168154811061140257611402613f2d565b5f91825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff1615156060820152612d99565b506001919050565b61146b6127ce565b6113d65f612f3c565b61147c6127ce565b6113c181612f8d565b604080516080810182525f80825260208201819052918101829052606081019190915260ca80546114b8906001906141e4565b815481106114c8576114c8613f2d565b5f91825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff1615156060820152919050565b600561152c8161244a565b8361153681612475565b6115535760405163932d94f760e01b815260040160405180910390fd5b61155b61251f565b5f5b83811015610d13573685858381811061157857611578613f2d565b905060200281019061158a9190613f41565b6001600160a01b0388165f90815260ce60209081526040808320549051939450926115bb918b918591879101614169565b6040516020818303038152906040528051906020012090505f6115dd84612578565b6001600160a01b038b165f90815260d3602090815260408083208684529091529020805460ff1916600190811790915590915061161b9084906141ac565b6001600160a01b038b165f81815260ce60205260409081902092909255905183919033907ffc8888bffd711da60bc5092b33f677d81896fe80ecc677b84cfab8184462b6e09061166e9088908a90614306565b60405180910390a461168c333083610cf26040890160208a016139ef565b50506001909201915061155d9050565b6001600160a01b0382165f90815260d66020526040812061171890826116cf6116ca3687900387018761431e565b612fe8565b815260208082019290925260409081015f208151606081018352905461ffff80821683526201000082041693820193909352600160201b90920463ffffffff1690820152612c7b565b9392505050565b33610edb818361304b565b6117326127ce565b6113c1816130ae565b60076117468161244a565b8261175081612475565b61176d5760405163932d94f760e01b815260040160405180910390fd5b60cb545f9061178990600160a01b900463ffffffff16426142d6565b6001600160a01b0386165f90815260d5602090815260408083208151606081018352905461ffff80821683526201000082041693820193909352600160201b90920463ffffffff1690820152919250906117e290612c7b565b6001600160a01b0387165f90815260d560205260409020909150611807908684613119565b6040805163ffffffff8416815261ffff838116602083015287168183015290516001600160a01b0388169133917fd1e028bd664486a46ad26040e999cd2d22e1e9a094ee6afe19fcf64678f16f749181900360600190a3505050505050565b60066118718161244a565b8361187b81612475565b6118985760405163932d94f760e01b815260040160405180910390fd5b60cb545f906118b490600160a01b900463ffffffff16426142d6565b6001600160a01b038781165f90815260d460209081526040808320938a1683529281528282208351606081018552905461ffff80821683526201000082041692820192909252600160201b90910463ffffffff169281019290925291925061191b90612c7b565b6001600160a01b038089165f90815260d460209081526040808320938b1683529290522090915061194d908684613119565b6040805163ffffffff8416815261ffff838116602083015287168183015290516001600160a01b0388811692908a169133917f48e198b6ae357e529204ee53a8e514c470ff77d9cc8e4f7207f8b5d490ae6934919081900360600190a450505050505050565b604080516080810182525f80825260208201819052918101829052606081019190915260ca82815481106119e9576119e9613f2d565b5f91825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff161515606082015292915050565b6001600160a01b038281165f90815260d46020908152604080832093851683529281528282208351606081018552905461ffff80821683526201000082041692820192909252600160201b90910463ffffffff16928101929092529061171890612c7b565b60ca545f905b63ffffffff811615611b0f578260ca611ac8600184614386565b63ffffffff1681548110611ade57611ade613f2d565b905f5260205f2090600202015f015403611afd57611718600182614386565b80611b07816143a2565b915050611aae565b5060405163504570e360e01b815260040160405180910390fd5b81611b3381612475565b611b505760405163932d94f760e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0384811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015611bb4573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611bd89190613f12565b80611c6c575060405163ba1a84e560e01b81526001600160a01b0384811660048301525f917f00000000000000000000000000000000000000000000000000000000000000009091169063ba1a84e590602401602060405180830381865afa158015611c46573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c6a91906143c0565b115b611c895760405163fb494ea160e01b815260040160405180910390fd5b6110c0838361304b565b611c9b6127ce565b6001600160a01b038116611d055760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b6113c181612f3c565b5f54610100900460ff1615808015611d2c57505f54600160ff909116105b80611d455750303b158015611d4557505f5460ff166001145b611da85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401611cfc565b5f805460ff191660011790558015611dc9575f805461ff0019166101001790555b611dd2856128cb565b611ddb86612f3c565b611de484612f8d565b611ded83612d28565b611df6826130ae565b8015611e3b575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b6008611e4e8161244a565b83611e5881612475565b611e755760405163932d94f760e01b815260040160405180910390fd5b6040516304c1b8eb60e31b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc75890611ec1908790600401613f04565b602060405180830381865afa158015611edc573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f009190613f12565b611f1d57604051631fb1705560e21b815260040160405180910390fd5b60cb545f90611f3990600160a01b900463ffffffff16426142d6565b6001600160a01b0387165f90815260d66020526040812091925090611f6b90826116cf6116ca368b90038b018b61431e565b6001600160a01b0388165f90815260d660205260408120919250611fb09190611f9c6116ca368b90038b018b61431e565b81526020019081526020015f208684613119565b866001600160a01b0316336001600160a01b03167f14918b3834ab6752eb2e1b489b6663a67810efb5f56f3944a97ede8ecf1fd9f18885858a604051611ff994939291906143d7565b60405180910390a350505050505050565b5f6001610eee60208401846139ef565b60036120258161244a565b60cb546001600160a01b0316331461205057604051635c427cd960e01b815260040160405180910390fd5b60ca5463ffffffff831610612078576040516394a8d38960e01b815260040160405180910390fd5b5f60ca8363ffffffff168154811061209257612092613f2d565b905f5260205f20906002020190508060010160089054906101000a900460ff16156120d057604051631b14174b60e01b815260040160405180910390fd5b6001810154600160201b900463ffffffff16421061210157604051630c36f66560e21b815260040160405180910390fd5b60018101805460ff60401b1916600160401b17905560405163ffffffff8416907fd850e6e5dfa497b72661fa73df2923464eaed9dc2ff1d3cb82bccbfeabe5c41e905f90a2505050565b6121536131e8565b6066548019821981161461217a5760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b5f6121c28161244a565b6121ca61251f565b5f5b828110156110b557368484838181106121e7576121e7613f2d565b90506020028101906121f9919061420c565b335f81815260ce60209081526040808320549051949550939192612223929091859187910161429e565b60405160208183030381529060405280519060200120905061224483612908565b335f90815260cf602090815260408083208484529091529020805460ff191660019081179091556122769083906141ac565b335f81815260ce602052604090819020929092559051829184917f450a367a380c4e339e5ae7340c8464ef27af7781ad9945cfe8abd828f89e6281906122bd9088906142c4565b60405180910390a46122de333060408601803590610cf290602089016139ef565b5050506001016121cc565b60046122f48161244a565b335f90815260d1602052604090205460ff1661232357604051635c427cd960e01b815260040160405180910390fd5b61232b61251f565b5f5b828110156110b5573684848381811061234857612348613f2d565b905060200281019061235a919061420c565b335f81815260ce60209081526040808320549051949550939192612384929091859187910161429e565b6040516020818303038152906040528051906020012090506123a583612908565b335f90815260d2602090815260408083208484529091529020805460ff191660019081179091556123d79083906141ac565b335f81815260ce602052604090819020929092559051829184917f5251b6fdefcb5d81144e735f69ea4c695fd43b0289ca53dc075033f5fc80068b9061241e9088906142c4565b60405180910390a461243f333060408601803590610cf290602089016139ef565b50505060010161232d565b606654600160ff83161b908116036113c15760405163840a48d560e01b815260040160405180910390fd5b604051631beb2b9760e31b81526001600160a01b0382811660048301523360248301523060448301525f80356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303815f875af11580156124fb573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061137a9190613f12565b6002609754036125715760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611cfc565b6002609755565b5f6125ab612586838061440b565b6125966080860160608701613c81565b6125a660a0870160808801613c81565b613299565b5f6125b9604084018461440b565b9050116125d95760405163796cc52560e01b815260040160405180910390fd5b426125ea60a0840160808501613c81565b6125fa6080850160608601613c81565b61260491906142d6565b63ffffffff16106126285760405163150358a160e21b815260040160405180910390fd5b5f80805b612639604086018661440b565b905081101561272a5736612650604087018761440b565b8381811061266057612660613f2d565b6040029190910191505f905061267960208301836139ef565b6001600160a01b0316036126a057604051630863a45360e11b815260040160405180910390fd5b6126ad60208201826139ef565b6001600160a01b0316836001600160a01b0316106126de576040516310fb47f160e31b815260040160405180910390fd5b5f816020013511612702576040516310eb483f60e21b815260040160405180910390fd5b61270f60208201826139ef565b925061271f6020820135856141ac565b93505060010161262c565b506f4b3b4ca85a86c47a098a223fffffffff82111561275c5760405163070b5a6f60e21b815260040160405180910390fd5b5092915050565b6040516001600160a01b038085166024830152831660448201526064810182905261131f9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613597565b6033546001600160a01b031633146113d65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401611cfc565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa15801561288a573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128ae9190613f12565b6113d657604051631d77d47760e21b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b612935612915828061440b565b6129256080850160608601613c81565b6125a660a0860160808701613c81565b5f816040013511612959576040516310eb483f60e21b815260040160405180910390fd5b6f4b3b4ca85a86c47a098a223fffffffff8160400135111561298e5760405163070b5a6f60e21b815260040160405180910390fd5b6129be63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016426141ac565b6129ce6080830160608401613c81565b63ffffffff1611156113c157604051637ee2b44360e01b815260040160405180910390fd5b5f60ca612a036020850185613c81565b63ffffffff1681548110612a1957612a19613f2d565b5f91825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff16151560608201529050612a798382612d99565b5f612a8a60808501606086016139ef565b6001600160a01b038082165f90815260cc60205260409020549192501680612aaf5750805b336001600160a01b03821614612ad857604051635c427cd960e01b815260040160405180910390fd5b5f5b612ae760a0870187614450565b9050811015611e3b5736612afe60e088018861440b565b83818110612b0e57612b0e613f2d565b6001600160a01b0387165f90815260cd602090815260408083209302949094019450929091508290612b42908501856139ef565b6001600160a01b03166001600160a01b031681526020019081526020015f2054905080826020013511612b885760405163aa385e8160e01b815260040160405180910390fd5b5f612b978260208501356141e4565b6001600160a01b0387165f90815260cd60209081526040822092935085018035929190612bc490876139ef565b6001600160a01b031681526020808201929092526040015f2091909155612c059089908390612bf5908701876139ef565b6001600160a01b0316919061366a565b86516001600160a01b03808a1691878216918916907f9543dbd55580842586a951f0386e24d68a5df99ae29e3b216588b45fd684ce3190612c4960208901896139ef565b604080519283526001600160a01b039091166020830152810186905260600160405180910390a4505050600101612ada565b5f816040015163ffffffff165f1480612cad5750815161ffff908116148015612cad5750816040015163ffffffff1642105b15612cc557505060cb54600160e01b900461ffff1690565b816040015163ffffffff16421015612cde57815161137a565b506020015190565b919050565b60605f612cf78361369a565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b60cb546040805163ffffffff600160a01b9093048316815291831660208301527faf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3910160405180910390a160cb805463ffffffff909216600160a01b0263ffffffff60a01b19909216919091179055565b806060015115612dbc57604051631b14174b60e01b815260040160405180910390fd5b806040015163ffffffff16421015612de757604051631437a2bb60e31b815260040160405180910390fd5b612df460c0830183614450565b9050612e0360a0840184614450565b905014612e23576040516343714afd60e01b815260040160405180910390fd5b612e3060e083018361440b565b9050612e3f60c0840184614450565b905014612e5f576040516343714afd60e01b815260040160405180910390fd5b8051612e8b90612e756040850160208601613c81565b612e826040860186614495565b866060016136c1565b5f5b612e9a60a0840184614450565b90508110156110c057612f346080840135612eb860a0860186614450565b84818110612ec857612ec8613f2d565b9050602002016020810190612edd9190613c81565b612eea60c0870187614450565b85818110612efa57612efa613f2d565b9050602002810190612f0c9190614495565b612f1960e089018961440b565b87818110612f2957612f29613f2d565b905060400201613765565b600101612e8d565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60cb546040516001600160a01b038084169216907f237b82f438d75fc568ebab484b75b01d9287b9e98b490b7c23221623b6705dbb905f90a360cb80546001600160a01b0319166001600160a01b0392909216919091179055565b5f815f0151826020015163ffffffff1660405160200161303392919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b60405160208183030381529060405261137a906144d7565b6001600160a01b038083165f81815260cc602052604080822080548686166001600160a01b0319821681179092559151919094169392849290917fbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca3129190a4505050565b60cb546040805161ffff600160e01b9093048316815291831660208301527fe6cd4edfdcc1f6d130ab35f73d72378f3a642944fb4ee5bd84b7807a81ea1c4e910160405180910390a160cb805461ffff909216600160e01b0261ffff60e01b19909216919091179055565b61271061ffff831611156131405760405163891c63df60e01b815260040160405180910390fd5b8254600160201b900463ffffffff16421161316e57604051637b1e25c560e01b815260040160405180910390fd5b8254600160201b900463ffffffff165f0361319557825461ffff191661ffff1783556131ac565b825462010000810461ffff1661ffff199091161783555b825463ffffffff909116600160201b0267ffffffff000000001961ffff90931662010000029290921667ffffffffffff00001990911617179055565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613244573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061326891906144fa565b6001600160a01b0316336001600160a01b0316146113d65760405163794821ff60e01b815260040160405180910390fd5b826132b75760405163796cc52560e01b815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168163ffffffff16111561330457604051630dd0b9f560e21b815260040160405180910390fd5b61332e7f000000000000000000000000000000000000000000000000000000000000000082614529565b63ffffffff16156133525760405163ee66470560e01b815260040160405180910390fd5b5f8163ffffffff16116133785760405163cb3f434d60e01b815260040160405180910390fd5b6133a27f000000000000000000000000000000000000000000000000000000000000000083614529565b63ffffffff16156133c657604051633c1a94f160e21b815260040160405180910390fd5b8163ffffffff167f000000000000000000000000000000000000000000000000000000000000000063ffffffff16426133ff91906141e4565b1115801561343957508163ffffffff167f000000000000000000000000000000000000000000000000000000000000000063ffffffff1611155b6134565760405163041aa75760e11b815260040160405180910390fd5b5f805b84811015611e3b575f86868381811061347457613474613f2d565b61348a92602060409092020190810191506139ef565b60405163198f077960e21b81526001600160a01b0380831660048301529192507f00000000000000000000000000000000000000000000000000000000000000009091169063663c1de490602401602060405180830381865afa1580156134f3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135179190613f12565b8061353e57506001600160a01b03811673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0145b61355b57604051632efd965160e11b815260040160405180910390fd5b806001600160a01b0316836001600160a01b03161061358d5760405163dfad9ca160e01b815260040160405180910390fd5b9150600101613459565b5f6135eb826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166137a39092919063ffffffff16565b905080515f148061360b57508080602001905181019061360b9190613f12565b6110c05760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401611cfc565b6040516001600160a01b0383166024820152604481018290526110c090849063a9059cbb60e01b90606401612797565b5f60ff8216601f81111561137a57604051632cd44ac360e21b815260040160405180910390fd5b6136cc602083614550565b6001901b8463ffffffff16106136f45760405162c6c39d60e71b815260040160405180910390fd5b5f6136fe82610edf565b905061374884848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508a92508591505063ffffffff89166137b9565b611e3b576040516369ca16c960e01b815260040160405180910390fd5b613770602083614550565b6001901b8463ffffffff16106137995760405163054ff4df60e51b815260040160405180910390fd5b5f6136fe8261200a565b60606137b184845f856137d0565b949350505050565b5f836137c68685856138a7565b1495945050505050565b6060824710156138315760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401611cfc565b5f5f866001600160a01b0316858760405161384c9190614563565b5f6040518083038185875af1925050503d805f8114613886576040519150601f19603f3d011682016040523d82523d5f602084013e61388b565b606091505b509150915061389c8783838761393e565b979650505050505050565b5f602084516138b69190614579565b156138d4576040516313717da960e21b815260040160405180910390fd5b8260205b85518111613935576138eb600285614579565b5f0361390c57815f528086015160205260405f209150600284049350613923565b808601515f528160205260405f2091506002840493505b61392e6020826141ac565b90506138d8565b50949350505050565b606083156139ac5782515f036139a5576001600160a01b0385163b6139a55760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611cfc565b50816137b1565b6137b183838151156139c15781518083602001fd5b8060405162461bcd60e51b8152600401611cfc9190613c4c565b6001600160a01b03811681146113c1575f5ffd5b5f602082840312156139ff575f5ffd5b8135611718816139db565b5f60408284031215613a1a575f5ffd5b50919050565b5f5f83601f840112613a30575f5ffd5b5081356001600160401b03811115613a46575f5ffd5b6020830191508360208260051b8501011115613a60575f5ffd5b9250929050565b5f5f5f60608486031215613a79575f5ffd5b613a838585613a0a565b925060408401356001600160401b03811115613a9d575f5ffd5b613aa986828701613a20565b9497909650939450505050565b80151581146113c1575f5ffd5b5f5f60408385031215613ad4575f5ffd5b8235613adf816139db565b91506020830135613aef81613ab6565b809150509250929050565b5f60208284031215613b0a575f5ffd5b5035919050565b5f60408284031215613b21575f5ffd5b6117188383613a0a565b5f5f60208385031215613b3c575f5ffd5b82356001600160401b03811115613b51575f5ffd5b613b5d85828601613a20565b90969095509350505050565b5f6101008284031215613a1a575f5ffd5b5f5f60408385031215613b8b575f5ffd5b82356001600160401b03811115613ba0575f5ffd5b613bac85828601613b69565b9250506020830135613aef816139db565b803563ffffffff81168114612ce6575f5ffd5b5f5f60408385031215613be1575f5ffd5b82359150613bf160208401613bbd565b90509250929050565b5f5f5f60408486031215613c0c575f5ffd5b83356001600160401b03811115613c21575f5ffd5b613c2d86828701613a20565b9094509250506020840135613c41816139db565b809150509250925092565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f60208284031215613c91575f5ffd5b61171882613bbd565b5f60208284031215613caa575f5ffd5b813560ff81168114611718575f5ffd5b5f60208284031215613cca575f5ffd5b81356001600160401b03811115613cdf575f5ffd5b6137b184828501613b69565b5f5f60408385031215613cfc575f5ffd5b8235613d07816139db565b946020939093013593505050565b5f5f60408385031215613d26575f5ffd5b8235613d31816139db565b91506020830135613aef816139db565b5f5f5f60408486031215613d53575f5ffd5b8335613d5e816139db565b925060208401356001600160401b03811115613a9d575f5ffd5b5f5f60608385031215613d89575f5ffd5b8235613d94816139db565b9150613bf18460208501613a0a565b803561ffff81168114612ce6575f5ffd5b5f60208284031215613dc4575f5ffd5b61171882613da3565b5f5f60408385031215613dde575f5ffd5b8235613de9816139db565b9150613bf160208401613da3565b5f5f5f60608486031215613e09575f5ffd5b8335613e14816139db565b92506020840135613e24816139db565b9150613e3260408501613da3565b90509250925092565b5f5f5f5f5f60a08688031215613e4f575f5ffd5b8535613e5a816139db565b9450602086013593506040860135613e71816139db565b9250613e7f60608701613bbd565b9150613e8d60808701613da3565b90509295509295909350565b5f5f5f60808486031215613eab575f5ffd5b8335613eb6816139db565b9250613ec58560208601613a0a565b9150613e3260608501613da3565b8035613ede816139db565b6001600160a01b0316825263ffffffff613efa60208301613bbd565b1660208301525050565b6040810161137a8284613ed3565b5f60208284031215613f22575f5ffd5b815161171881613ab6565b634e487b7160e01b5f52603260045260245ffd5b5f823560be19833603018112613f55575f5ffd5b9190910192915050565b5f5f8335601e19843603018112613f74575f5ffd5b83016020810192503590506001600160401b03811115613f92575f5ffd5b8060061b3603821315613a60575f5ffd5b8183526020830192505f815f5b84811015614006578135613fc3816139db565b6001600160a01b0316865260208201356bffffffffffffffffffffffff8116808214613fed575f5ffd5b6020880152506040958601959190910190600101613fb0565b5093949350505050565b5f5f8335601e19843603018112614025575f5ffd5b83016020810192503590506001600160401b03811115614043575f5ffd5b803603821315613a60575f5ffd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f6140848283613f5f565b60c0855261409660c086018284613fa3565b91505060208301356140a7816139db565b6001600160a01b031660208501526140c26040840184613f5f565b858303604087015280835290915f91906020015b818310156141115783356140e9816139db565b6001600160a01b031681526020848101359082015260409384019360019390930192016140d6565b61411d60608701613bbd565b63ffffffff81166060890152935061413760808701613bbd565b63ffffffff81166080890152935061415260a0870187614010565b9450925086810360a088015261389c818585614051565b60018060a01b0384168152826020820152606060408201525f61418f6060830184614079565b95945050505050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561137a5761137a614198565b6141c98185613ed3565b826040820152608060608201525f61418f6080830184614079565b8181038181111561137a5761137a614198565b5f8161420557614205614198565b505f190190565b5f8235609e19833603018112613f55575f5ffd5b5f61422b8283613f5f565b60a0855261423d60a086018284613fa3565b915050602083013561424e816139db565b6001600160a01b031660208501526040838101359085015263ffffffff61427760608501613bbd565b16606085015263ffffffff61428e60808501613bbd565b1660808501528091505092915050565b60018060a01b0384168152826020820152606060408201525f61418f6060830184614220565b602081525f6117186020830184614220565b63ffffffff818116838216019081111561137a5761137a614198565b5f823560fe19833603018112613f55575f5ffd5b828152604060208201525f6137b16040830184614079565b5f604082840312801561432f575f5ffd5b50604080519081016001600160401b038111828210171561435e57634e487b7160e01b5f52604160045260245ffd5b604052823561436c816139db565b815261437a60208401613bbd565b60208201529392505050565b63ffffffff828116828216039081111561137a5761137a614198565b5f63ffffffff8216806143b7576143b7614198565b5f190192915050565b5f602082840312156143d0575f5ffd5b5051919050565b60a081016143e58287613ed3565b63ffffffff94909416604082015261ffff92831660608201529116608090910152919050565b5f5f8335601e19843603018112614420575f5ffd5b8301803591506001600160401b03821115614439575f5ffd5b6020019150600681901b3603821315613a60575f5ffd5b5f5f8335601e19843603018112614465575f5ffd5b8301803591506001600160401b0382111561447e575f5ffd5b6020019150600581901b3603821315613a60575f5ffd5b5f5f8335601e198436030181126144aa575f5ffd5b8301803591506001600160401b038211156144c3575f5ffd5b602001915036819003821315613a60575f5ffd5b80516020808301519190811015613a1a575f1960209190910360031b1b16919050565b5f6020828403121561450a575f5ffd5b8151611718816139db565b634e487b7160e01b5f52601260045260245ffd5b5f63ffffffff83168061453e5761453e614515565b8063ffffffff84160691505092915050565b5f8261455e5761455e614515565b500490565b5f82518060208501845e5f920191825250919050565b5f8261458757614587614515565b50069056fea2646970667358221220319adaf6a1466ae377b31189860d1faa51f67ddd645e00b4f2a02f66148a3f2964736f6c634300081b0033",
}

// RewardsCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardsCoordinatorMetaData.ABI instead.
var RewardsCoordinatorABI = RewardsCoordinatorMetaData.ABI

// RewardsCoordinatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RewardsCoordinatorMetaData.Bin instead.
var RewardsCoordinatorBin = RewardsCoordinatorMetaData.Bin

// DeployRewardsCoordinator deploys a new Ethereum contract, binding an instance of RewardsCoordinator to it.
func DeployRewardsCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, params IRewardsCoordinatorTypesRewardsCoordinatorConstructorParams) (common.Address, *types.Transaction, *RewardsCoordinator, error) {
	parsed, err := RewardsCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RewardsCoordinatorBin), backend, params)
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_RewardsCoordinator *RewardsCoordinatorCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RewardsCoordinator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_RewardsCoordinator *RewardsCoordinatorSession) Version() (string, error) {
	return _RewardsCoordinator.Contract.Version(&_RewardsCoordinator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_RewardsCoordinator *RewardsCoordinatorCallerSession) Version() (string, error) {
	return _RewardsCoordinator.Contract.Version(&_RewardsCoordinator.CallOpts)
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
