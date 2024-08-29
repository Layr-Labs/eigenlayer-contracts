// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IStrategy.sol";
import "../interfaces/IStakeRootCompendium.sol";
import "../libraries/Merkle.sol";

import "@risc0-ethereum/IRiscZeroVerifier.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import "../libraries/Checkpoints.sol";

abstract contract StakeRootCompendiumStorage is IStakeRootCompendium, OwnableUpgradeable {
    using Checkpoints for Checkpoints.History;
    using EnumerableMap for EnumerableMap.AddressToUintMap;
    using EnumerableSet for EnumerableSet.Bytes32Set;

    /// @notice the maximum number of operators that can be in an operator set in the StakeTree
    uint32 public constant MAX_OPERATOR_SET_SIZE = 2048;
    /// @notice the maximum number of operator sets that can be in the StakeTree
    uint32 public constant MAX_NUM_OPERATOR_SETS = 2048;
    /// @notice the maximum number of strategies that each operator set in the StakeTree can use to weight their operator stakes
    uint32 public constant MAX_NUM_STRATEGIES = 20;
    /// @notice the placeholder index used for operator sets that are removed from the StakeTree
    uint32 public constant REMOVED_INDEX = type(uint32).max;

    /// @notice the minimum balance that must be maintained for an operatorSet
    uint256 public constant MIN_DEPOSIT_BALANCE = 0.1 ether;

    /// @notice the delegation manager contract
    IDelegationManager public immutable delegationManager;
    /// @notice the AVS directory contract
    IAVSDirectory public immutable avsDirectory;

    /// @notice the interval at which proofs can be posted, to not overcharge the operatorSets
    uint32 public immutable proofInterval;
    /// @notice the period of time within which a root can be marked as blacklisted
    uint32 public immutable blacklistWindow;

    /// @notice the linear charge per proof in the number of strategies
    uint64 public linearChargePerProof;
    /// @notice the constant charge per proof
    uint64 public constantChargePerProof;
    /// @notice the total amount that a operatorSet that had been in the stakeTree since genesis would have been charged in constant charges
    Checkpoints.History internal cumulativeContantChargeSnapshot;
    /// @notice the total amount that a operatorSet that had been in the stakeTree since genesis would have been charged in linear charges
    Checkpoints.History internal cumulativeLinearChargeSnapshot;

    /// @notice deposit balance to be deducted for operatorSets
    mapping(address => mapping(uint32 => DepositBalanceInfo)) public depositBalanceInfo;

    /// @notice map from operator set to a trace of their index over time
    mapping(address => mapping(uint32 => Checkpoints.History)) internal operatorSetToIndex;
    /// @notice list of operator sets that have been configured to be in the StakeTree
    IAVSDirectory.OperatorSet[] public operatorSets;
    
    /// @notice the total number of strategies among all operator sets (with duplicates)
    uint256 public totalStrategies;
    /// @notice the total charge for a proofs at a certain time depending on the number of strategies
    Checkpoints.History internal totalChargeSnapshot;

    /// @notice the strategies and multipliers for each operator set
    mapping(address => mapping(uint32 => EnumerableMap.AddressToUintMap)) internal operatorSetToStrategyAndMultipliers;
    /// @notice the extraData for each operatorSet
    mapping(address => mapping(uint32 => bytes32)) internal operatorSetExtraDatas;
    /// @notice the extraData for each operator in each operator set
    mapping(address => mapping(uint32 => mapping(address => bytes32))) internal operatorExtraDatas;

    /// @notice the verifier contract that will be used to verify snark proofs
    address public verifier;
    /// @notice the id of the program being verified when roots are posted
    bytes32 public imageId;

    /// @notice the index of the latest stake root submission that has been charged
    uint256 public latestChargedSubmissionIndex;
    /// @notice the stake root submissions that have been posted
    StakeRootSubmission[] public stakeRootSubmissions;

    constructor(
        IDelegationManager _delegationManager,
        IAVSDirectory _avsDirectory,
        uint32 _proofInterval,
        uint32 _blacklistWindow
    ) {
        // _disableInitializers();
        delegationManager = _delegationManager;
        avsDirectory = _avsDirectory;
        proofInterval = _proofInterval;
        blacklistWindow = _blacklistWindow;
    }
}
