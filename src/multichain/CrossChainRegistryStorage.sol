// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "../contracts/interfaces/ICrossChainRegistry.sol";
import "../contracts/interfaces/IOperatorTableCalculator.sol";
import "../contracts/interfaces/IAllocationManager.sol";
import "../contracts/libraries/OperatorSetLib.sol";

/**
 * @title CrossChainRegistryStorage
 * @author Layr Labs, Inc.
 * @notice Storage contract for the CrossChainRegistry, containing all storage variables and immutables
 * @dev This abstract contract is designed to be inherited by the CrossChainRegistry implementation
 */
abstract contract CrossChainRegistryStorage is ICrossChainRegistry {
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableSet for EnumerableSet.UintSet;
    using OperatorSetLib for OperatorSet;

    // Constants

    /// @dev Index for flag that pauses generation reservations when set
    uint8 internal constant PAUSED_GENERATION_RESERVATIONS = 0;

    /// @dev Index for flag that pauses operator table calculator modifications when set
    uint8 internal constant PAUSED_OPERATOR_TABLE_CALCULATOR = 1;

    /// @dev Index for flag that pauses operator set config modifications when set
    uint8 internal constant PAUSED_OPERATOR_SET_CONFIG = 2;

    /// @dev Index for flag that pauses transport destination modifications when set
    uint8 internal constant PAUSED_TRANSPORT_DESTINATIONS = 3;

    /// @dev Index for flag that pauses chain whitelist modifications when set
    uint8 internal constant PAUSED_CHAIN_WHITELIST = 4;

    // Immutables

    /// @notice The AllocationManager contract for EigenLayer
    IAllocationManager public immutable allocationManager;

    // Mutatables

    /// GENERATION RESERVATIONS

    /// @dev Set of operator sets with active generation reservations
    EnumerableSet.Bytes32Set internal _activeGenerationReservations;

    /// @dev Mapping from operator set key to operator table calculator for active reservations
    mapping(bytes32 operatorSetKey => IOperatorTableCalculator) internal _operatorTableCalculators;

    /// TRANSPORT RESERVATIONS

    /// @dev Set of operator sets with active transport reservations
    EnumerableSet.Bytes32Set internal _activeTransportReservations;

    /// @dev Mapping from operator set key to set of chain IDs for transport destinations
    mapping(bytes32 operatorSetKey => EnumerableSet.UintSet) internal _transportDestinations;

    /// @dev Mapping from operator set key to operator set configuration
    mapping(bytes32 operatorSetKey => OperatorSetConfig) internal _operatorSetConfigs;

    /// CHAIN WHITELISTING

    /// @dev Set of whitelisted chain IDs that can be used as transport destinations
    EnumerableSet.UintSet internal _whitelistedChainIDs;

    // Construction

    constructor(
        IAllocationManager _allocationManager
    ) {
        allocationManager = _allocationManager;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[41] private __gap;
}
