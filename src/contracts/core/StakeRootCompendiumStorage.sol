// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IAllocationManager.sol";
import "../interfaces/IStakeRootCompendium.sol";
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
    uint256 public immutable MIN_BALANCE_THRESHOLD;

    /// @notice the placeholder index used for operator sets that are removed from the StakeTree
    uint32 public constant REMOVED_INDEX = type(uint32).max;

    /// @notice the minimum number of proofs that an operatorSet's deposit balance needs to cover and
    /// the number of proofs they must pay for since their latest reconfiguration
    /// @dev this prevents de-registering an operatorSet immediately after reconfiguring
    uint256 public immutable MIN_PREPAID_PROOFS;

    /// @notice the total number of strategies among all operator sets (with duplicates)
    uint256 public totalStrategies;

    /// @notice the verifier contract that will be used to verify snark proofs
    address public immutable verifier;
    /// @notice the id of the program being verified when roots are posted
    bytes32 public immutable imageId;

    /// @notice the address allowed to confirm roots
    address public rootConfirmer;

    /// @notice list of operator sets that have been configured to be in the StakeTree
    OperatorSet[] public operatorSets;
    /// @notice the total charge for a proofs at a certain time depending on the number of strategies
    Snapshots.History internal chargePerProof;
    /// @dev Contains cumulative charges for operator sets, strategies, and max total charge.
    ChargeParams public chargeParams;
    /// @dev Contains info about cumulative charges.
    CumulativeChargeParams public cumulativeChargeParams;
    /// @notice deposit balance to be deducted for operatorSets
    mapping(address => mapping(uint32 => DepositInfo)) public depositInfos;
    /// @notice the extraData for each operatorSet
    mapping(address => mapping(uint32 => bytes32)) internal operatorSetExtraDatas;
    /// @notice map from operator set to a trace of their index over time
    mapping(address => mapping(uint32 => Snapshots.History)) internal operatorSetToIndex;
    /// @notice the extraData for each operator in each operator set
    mapping(address => mapping(uint32 => mapping(address => bytes32))) internal operatorExtraDatas;
    /// @notice the strategies and multipliers for each operator set
    mapping(address => mapping(uint32 => EnumerableMap.AddressToUintMap)) internal operatorSetToStrategyAndMultipliers;

    /// @notice the stake root submissions
    IStakeRootCompendium.StakeRootSubmission[] public stakeRootSubmissions;

    constructor(
        IDelegationManager _delegationManager,
        IAVSDirectory _avsDirectory,
        IAllocationManager _allocationManager,
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

        // note verifier and imageId are immutable and set by implementation contract
        // since proof verification is in the hot path, this is a gas optimization to avoid calling the storage contract for verifier and imageId
        // however the new impl does not have access to the immutable variables of the last impl so we can't reference the old verifier and imageId
        verifier = _verifier;
        imageId = _imageId;
    }
}
