// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import "../interfaces/ICrossChainRegistry.sol";
import "../interfaces/IOperatorTableCalculator.sol";
import "../interfaces/IAllocationManager.sol";
import "../interfaces/IKeyRegistrar.sol";
import "../libraries/OperatorSetLib.sol";

/**
 * @title CrossChainRegistryStorage
 * @author Layr Labs, Inc.
 * @notice Storage contract for the CrossChainRegistry, containing all storage variables and immutables
 * @dev This abstract contract is designed to be inherited by the CrossChainRegistry implementation
 */
abstract contract CrossChainRegistryStorage is ICrossChainRegistry {
    using EnumerableMap for EnumerableMap.UintToAddressMap;
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

    /// @notice The KeyRegistrar contract for EigenLayer
    IKeyRegistrar public immutable keyRegistrar;

    // Mutatables

    /// GENERATION RESERVATIONS

    /// @notice Mapping of generation reservations for operator sets
    EnumerableSet.Bytes32Set internal _activeGenerationReservations;

    /// @dev Mapping from operator set key to operator table calculator for active reservations
    mapping(bytes32 operatorSetKey => IOperatorTableCalculator) internal _operatorTableCalculators;

    /// @dev Mapping from operator set key to operator set configuration
    mapping(bytes32 operatorSetKey => OperatorSetConfig) internal _operatorSetConfigs;

    /// @dev Mapping from operator set key to set of chain IDs for transport destinations
    mapping(bytes32 operatorSetKey => EnumerableSet.UintSet) internal _transportDestinations;

    /// CHAIN WHITELISTING

    /// @dev Map of whitelisted chain IDs to operator table updaters
    EnumerableMap.UintToAddressMap internal _whitelistedChainIDs;

    /// @notice Table update cadence for all chains
    uint32 internal _tableUpdateCadence;

    // Construction

    constructor(IAllocationManager _allocationManager, IKeyRegistrar _keyRegistrar) {
        allocationManager = _allocationManager;
        keyRegistrar = _keyRegistrar;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[42] private __gap;
}
