// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IAllocationManager.sol";
import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";
import {Snapshots} from "../libraries/Snapshots.sol";

abstract contract AllocationManagerStorage is IAllocationManager {
    using Snapshots for Snapshots.History;

    // Constants

    /// @notice The EIP-712 typehash for the contract's domain
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");

    /// @notice The EIP-712 typehash for the `MagnitudeAdjustments` struct used by the contract
    bytes32 public constant MAGNITUDE_ADJUSTMENT_TYPEHASH = keccak256(
        "MagnitudeAdjustments(address operator,MagnitudeAdjustment(address strategy, OperatorSet(address avs, uint32 operatorSetId)[], uint64[] magnitudeDiffs)[],bytes32 salt,uint256 expiry)"
    );

    /// @dev Index for flag that pauses operator allocations/deallocations when set.
    uint8 internal constant PAUSED_STAKE_ALLOCATIONS_AND_DEALLOCATIONS = 0;

    /// @dev Index for flag that pauses operator register/deregister to operator sets when set.
    uint8 internal constant PAUSED_OPERATOR_SLASHING = 1;

    /// @dev BIPS factor for slashable bips
    uint256 internal constant BIPS_FACTOR = 10_000;

    /// @dev Maximum number of pending updates that can be queued for allocations/deallocations
    uint256 internal constant MAX_PENDING_UPDATES = 1;

    // Immutables

    /// @notice The DelegationManager contract for EigenLayer
    IDelegationManager public immutable delegation;

    /// @notice The AVSDirectory contract for EigenLayer
    IAVSDirectory public immutable avsDirectory;

    /// @notice Delay before deallocations are completable and can be added back into freeMagnitude
    /// In this window, deallocations still remain slashable by the operatorSet they were allocated to.
    uint32 public immutable DEALLOCATION_DELAY;

    /// @dev Delay before alloaction delay modifications take effect.
    uint32 public immutable ALLOCATION_CONFIGURATION_DELAY; // QUESTION: 21 days?

    /// @dev Returns the chain ID from the time the contract was deployed.
    uint256 internal immutable ORIGINAL_CHAIN_ID;

    // Mutatables

    /**
     * @notice Original EIP-712 Domain separator for this contract.
     * @dev The domain separator may change in the event of a fork that modifies the ChainID.
     * Use the getter function `domainSeparator` to get the current domain separator for this contract.
     */
    bytes32 internal _DOMAIN_SEPARATOR;

    /// @notice Mapping: operator => salt => Whether the salt has been used or not.
    mapping(address => mapping(bytes32 => bool)) public operatorSaltIsSpent;

    /// @notice Mapping: operator => strategy => snapshotted totalMagnitude
    /// Note that totalMagnitude is monotonically decreasing and only gets updated upon slashing
    mapping(address => mapping(IStrategy => Snapshots.History)) internal _totalMagnitudeUpdate;

    /// @notice Mapping: operator => strategy => OperatorMagnitudeInfo to keep track of info regarding pending magnitude allocations.
    mapping(address => mapping(IStrategy => FreeMagnitudeInfo)) public operatorFreeMagnitudeInfo;

    

    /// @notice Mapping: operator => strategy => operatorSet (encoded) => MagnitudeInfo
    mapping(address => mapping(IStrategy => mapping(bytes32 => MagnitudeInfo))) internal _operatorMagnitudeInfo;

    /// @notice Mapping: operator => strategy => amount of magnitude they can allocate
    mapping(address => mapping(IStrategy => uint64)) internal _allocatableMagnitude;

    struct PendingModification {
        uint32 effectTimestamp;
        int128 magnitudeDelta;
    }

    /// @notice Mapping: operator => Strategy => operatorSet => pending allocation or deallocation
    mapping(address => mapping(IStrategy => mapping(bytes32 => PendingModification))) _pendingModifications;

    /// @notice Mapping: operator => strategy => operatorSet => currently allocated magnitude
    mapping(address => mapping(IStrategy => mapping(bytes32 => uint64))) internal _allocatedMagnitude;

    struct OperatorSetQueue {
        uint256 nextIndex;
        bytes32[] operatorSets;
    }

    /// @notice Mapping: operator => strategy => operatorSet[] (encoded) to keep track of pending free magnitude for operatorSet from deallocations
    mapping(address => mapping(IStrategy => OperatorSetQueue)) internal deallocationQueue;

    /// @notice Mapping: operator => allocation delay (in seconds) for the operator.
    /// This determines how long it takes for allocations to take effect in the future.
    mapping(address => AllocationDelayInfo) internal _allocationDelayInfo;

    // Construction

    constructor(
        IDelegationManager _delegation,
        IAVSDirectory _avsDirectory,
        uint32 _DEALLOCATION_DELAY,
        uint32 _ALLOCATION_CONFIGURATION_DELAY
    ) {
        delegation = _delegation;
        avsDirectory = _avsDirectory;
        DEALLOCATION_DELAY = _DEALLOCATION_DELAY;
        ALLOCATION_CONFIGURATION_DELAY = _ALLOCATION_CONFIGURATION_DELAY;
        ORIGINAL_CHAIN_ID = block.chainid;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[44] private __gap;
}
