// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IAllocationManager.sol";
import "../interfaces/IStrategy.sol";
import "../interfaces/IStakeRootCompendium.sol";
import "../libraries/Merkle.sol";

import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import "../libraries/Snapshots.sol";

abstract contract StakeRootCompendiumStorage is IStakeRootCompendium, OwnableUpgradeable {
    using Snapshots for Snapshots.History;
    using EnumerableMap for EnumerableMap.AddressToUintMap;
    using EnumerableSet for EnumerableSet.Bytes32Set;

    /// @notice the delegation manager contract
    IDelegationManager public immutable delegationManager;
    /// @notice the AVS directory contract
    IAVSDirectory public immutable avsDirectory;
    /// @notice the allocation manager contract
    IAllocationManager public immutable allocationManager;

    /// @notice the minimum balance that must be maintained for an operatorSet
    /// @dev this balance compensates gas costs to deregister an operatorSet
    uint256 immutable public MIN_BALANCE_THRESHOLD;

    /// @notice the placeholder index used for operator sets that are removed from the StakeTree
    uint32 public constant REMOVED_INDEX = type(uint32).max;

    /// @notice the minimum number of proofs that an operatorSet's deposit balance needs to cover and 
    /// the number of proofs they must pay for since their latest reconfiguration
    /// @dev this prevents de-registering an operatorSet immediately after reconfiguring
    uint256 immutable public MIN_PREPAID_PROOFS;

    /// @notice the verifier contract that will be used to verify snark proofs
    address public immutable verifier;
    /// @notice the id of the program being verified when roots are posted
    bytes32 public immutable imageId;

    /// @notice the interval in seconds at which proofs can be posted
    uint32 public proofIntervalSeconds;
    /// @notice the address allowed to confirm roots
    address public rootConfirmer;
    /// @notice the linear charge per proof in the number of strategies
    uint96 public chargePerOperatorSet;
    /// @notice the constant charge per proof
    uint96 public chargePerStrategy;

    /// @notice the max total charge for a stakeRoot proof. used to bound computation offchain
    uint96 public maxTotalCharge;
    /// @notice the last time the cumulative charges were updated
    uint32 public cumulativeChargeLastUpdatedTimestamp;
    /// @notice the cumulative constant charge per operator set since deployment
    uint96 public cumulativeChargePerOperatorSetLastUpdate;
    /// @notice the cumulative linear charge per strategy per operatorSet since deployment
    uint96 public cumulativeChargePerStrategyLastUpdate;
    /// @notice deposit balance to be deducted for operatorSets
    mapping(address => mapping(uint32 => DepositInfo)) public depositInfos;

    /// @notice map from operator set to a trace of their index over time
    mapping(address => mapping(uint32 => Snapshots.History)) internal operatorSetToIndex;
    /// @notice list of operator sets that have been configured to be in the StakeTree
    OperatorSet[] public operatorSets;
    
    /// @notice the total number of strategies among all operator sets (with duplicates)
    uint256 public totalStrategies;
    /// @notice the total charge for a proofs at a certain time depending on the number of strategies
    Snapshots.History internal chargePerProofHistory;

    /// @notice the strategies and multipliers for each operator set
    mapping(address => mapping(uint32 => EnumerableMap.AddressToUintMap)) internal operatorSetToStrategyAndMultipliers;
    /// @notice the extraData for each operatorSet
    mapping(address => mapping(uint32 => bytes32)) internal operatorSetExtraDatas;
    /// @notice the extraData for each operator in each operator set
    mapping(address => mapping(uint32 => mapping(address => bytes32))) internal operatorExtraDatas;

    /// @notice the stake root submissions
    IStakeRootCompendium.StakeRootSubmission[] public stakeRootSubmissions;
    
    constructor(
        IDelegationManager _delegationManager,
        IAVSDirectory _avsDirectory,
        IAllocationManager _allocationManager,
        uint256 _maxTotalCharge,
        uint256 _minBalanceThreshold,
        uint256 _minPrepaidProofs,
        address _verifier,
        bytes32 _imageId
    ) {
        delegationManager = _delegationManager;
        avsDirectory = _avsDirectory;
        allocationManager = _allocationManager;
        MIN_BALANCE_THRESHOLD = _minBalanceThreshold;
        MIN_PREPAID_PROOFS = _minPrepaidProofs;
        verifier = _verifier;
        imageId = _imageId;
    }
}
