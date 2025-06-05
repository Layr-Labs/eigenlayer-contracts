// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "../mixins/PermissionControllerMixin.sol";
import "../mixins/SemVerMixin.sol";
import "../permissions/Pausable.sol";
import "../interfaces/IKeyRegistrar.sol";
import "./CrossChainRegistryStorage.sol";

/**
 * @title CrossChainRegistry
 * @author Layr Labs, Inc.
 * @notice Implementation contract for managing cross-chain operator set configurations and generation reservations
 * @dev Manages operator table calculations, transport destinations, and operator set configurations for cross-chain operations
 */
contract CrossChainRegistry is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    CrossChainRegistryStorage,
    PermissionControllerMixin,
    SemVerMixin
{
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableSet for EnumerableSet.UintSet;
    using OperatorSetLib for OperatorSet;

    /**
     *
     *                         MODIFIERS
     *
     */

    /**
     * @dev Validates that the operator set exists in the AllocationManager
     * @param operatorSet The operator set to validate
     */
    modifier isValidOperatorSet(
        OperatorSet calldata operatorSet
    ) {
        require(allocationManager.isOperatorSet(operatorSet), InvalidOperatorSet());
        _;
    }

    /**
     * @dev Validates that the chainIDs array is not empty
     * @param chainIDs The array of chain IDs to validate
     */
    modifier nonEmptyChainIDsArray(
        uint32[] calldata chainIDs
    ) {
        require(chainIDs.length > 0, EmptyChainIDsArray());
        _;
    }

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the CrossChainRegistry with immutable dependencies
     * @param _allocationManager The allocation manager for operator set validation
     * @param _keyRegistrar The key registrar for operator set curve type validation
     * @param _permissionController The permission controller for access control
     * @param _pauserRegistry The pauser registry for pause functionality
     * @param _version The semantic version of the contract
     */
    constructor(
        IAllocationManager _allocationManager,
        IKeyRegistrar _keyRegistrar,
        IPermissionController _permissionController,
        IPauserRegistry _pauserRegistry,
        string memory _version
    )
        CrossChainRegistryStorage(_allocationManager, _keyRegistrar)
        PermissionControllerMixin(_permissionController)
        Pausable(_pauserRegistry)
        SemVerMixin(_version)
    {
        _disableInitializers();
    }

    /**
     * @notice Initializes the contract with the initial paused status and owner
     * @param initialOwner The initial owner of the contract
     * @param initialPausedStatus The initial paused status bitmap
     */
    function initialize(address initialOwner, uint256 initialPausedStatus) external initializer {
        _transferOwnership(initialOwner);
        _setPausedStatus(initialPausedStatus);
    }

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

    /// @inheritdoc ICrossChainRegistry
    function requestGenerationReservation(
        OperatorSet calldata operatorSet,
        IOperatorTableCalculator operatorTableCalculator,
        OperatorSetConfig calldata config
    )
        external
        onlyWhenNotPaused(PAUSED_GENERATION_RESERVATIONS)
        checkCanCall(operatorSet.avs)
        isValidOperatorSet(operatorSet)
    {
        // Add to active generation reservations
        require(_activeGenerationReservations.add(operatorSet.key()), GenerationReservationAlreadyExists());
        emit GenerationReservationMade(operatorSet);

        // Set the operator table calculator
        _setOperatorTableCalculator(operatorSet, operatorTableCalculator, false);
        // Set the operator set config
        _setOperatorSetConfig(operatorSet, config, false);
    }

    /// @inheritdoc ICrossChainRegistry
    function removeGenerationReservation(
        OperatorSet calldata operatorSet
    )
        external
        onlyWhenNotPaused(PAUSED_GENERATION_RESERVATIONS)
        checkCanCall(operatorSet.avs)
        isValidOperatorSet(operatorSet)
    {
        // Remove from active generation reservations
        require(_activeGenerationReservations.remove(operatorSet.key()), GenerationReservationDoesNotExist());
        emit GenerationReservationRemoved(operatorSet);

        // Remove the operator table calculator
        _setOperatorTableCalculator(operatorSet, IOperatorTableCalculator(address(0)), true);
        // Remove the operator set config
        _setOperatorSetConfig(operatorSet, OperatorSetConfig(address(0), 0), true);
    }

    /// @inheritdoc ICrossChainRegistry
    function setOperatorTableCalculator(
        OperatorSet calldata operatorSet,
        IOperatorTableCalculator operatorTableCalculator
    )
        external
        onlyWhenNotPaused(PAUSED_OPERATOR_TABLE_CALCULATOR)
        checkCanCall(operatorSet.avs)
        isValidOperatorSet(operatorSet)
    {
        // Check if generation reservation exists
        require(_activeGenerationReservations.contains(operatorSet.key()), GenerationReservationDoesNotExist());

        // Set the operator table calculator
        _setOperatorTableCalculator(operatorSet, operatorTableCalculator, false);
    }

    /// @inheritdoc ICrossChainRegistry
    function setOperatorSetConfig(
        OperatorSet calldata operatorSet,
        OperatorSetConfig calldata config
    )
        external
        onlyWhenNotPaused(PAUSED_OPERATOR_SET_CONFIG)
        checkCanCall(operatorSet.avs)
        isValidOperatorSet(operatorSet)
    {
        // Check if generation reservation exists
        require(_activeGenerationReservations.contains(operatorSet.key()), GenerationReservationDoesNotExist());

        // Set the operator set config
        _setOperatorSetConfig(operatorSet, config, false);
    }

    /// @inheritdoc ICrossChainRegistry
    function addTransportDestinations(
        OperatorSet calldata operatorSet,
        uint32[] calldata chainIDs
    )
        external
        onlyWhenNotPaused(PAUSED_TRANSPORT_DESTINATIONS)
        checkCanCall(operatorSet.avs)
        isValidOperatorSet(operatorSet)
        nonEmptyChainIDsArray(chainIDs)
    {
        bytes32 operatorSetKey = operatorSet.key();

        // Check if this is the first transport destination (create reservation if needed)
        _activeTransportReservations.add(operatorSetKey);

        for (uint256 i = 0; i < chainIDs.length; i++) {
            uint32 chainID = chainIDs[i];

            // Validate chainID
            require(chainID != 0, InvalidChainId());
            // Check if chainID is whitelisted
            require(_whitelistedChainIDs.contains(chainID), ChainIDNotWhitelisted());

            // Add transport destination
            require(_transportDestinations[operatorSetKey].add(chainID), TransportDestinationAlreadyAdded());

            emit TransportDestinationAdded(operatorSet, chainID);
        }
    }

    /// @inheritdoc ICrossChainRegistry
    function removeTransportDestinations(
        OperatorSet calldata operatorSet,
        uint32[] calldata chainIDs
    )
        external
        onlyWhenNotPaused(PAUSED_TRANSPORT_DESTINATIONS)
        checkCanCall(operatorSet.avs)
        isValidOperatorSet(operatorSet)
        nonEmptyChainIDsArray(chainIDs)
    {
        bytes32 operatorSetKey = operatorSet.key();

        // Check if transport reservation exists
        require(_activeTransportReservations.contains(operatorSetKey), TransportReservationDoesNotExist());

        for (uint256 i = 0; i < chainIDs.length; i++) {
            uint32 chainID = chainIDs[i];

            // Validate chainID
            require(chainID != 0, InvalidChainId());

            // Remove transport destination
            require(_transportDestinations[operatorSetKey].remove(chainID), TransportDestinationNotFound());

            emit TransportDestinationRemoved(operatorSet, chainID);
        }

        // Check if all transport destinations have been removed
        if (_transportDestinations[operatorSetKey].length() == 0) {
            _activeTransportReservations.remove(operatorSetKey);
        }
    }

    /// @inheritdoc ICrossChainRegistry
    function addChainIDsToWhitelist(
        uint32[] calldata chainIDs
    ) external onlyOwner onlyWhenNotPaused(PAUSED_CHAIN_WHITELIST) nonEmptyChainIDsArray(chainIDs) {
        for (uint256 i = 0; i < chainIDs.length; i++) {
            uint32 chainID = chainIDs[i];

            // Validate chainID
            require(chainID != 0, InvalidChainId());

            // Add to whitelist
            require(_whitelistedChainIDs.add(chainID), ChainIDAlreadyWhitelisted());

            emit ChainIDAddedToWhitelist(chainID);
        }
    }

    /// @inheritdoc ICrossChainRegistry
    function removeChainIDsFromWhitelist(
        uint32[] calldata chainIDs
    ) external onlyOwner onlyWhenNotPaused(PAUSED_CHAIN_WHITELIST) nonEmptyChainIDsArray(chainIDs) {
        for (uint256 i = 0; i < chainIDs.length; i++) {
            uint32 chainID = chainIDs[i];

            // Validate chainID
            require(chainID != 0, InvalidChainId());

            // Remove from whitelist
            require(_whitelistedChainIDs.remove(chainID), ChainIDNotWhitelisted());

            emit ChainIDRemovedFromWhitelist(chainID);
        }
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @dev Internal function to set the operator table calculator for an operator set
     * @param operatorSet The operator set to set the calculator for
     * @param operatorTableCalculator The operator table calculator contract
     * @param isDelete Whether to delete the operator table calculator
     */
    function _setOperatorTableCalculator(
        OperatorSet memory operatorSet,
        IOperatorTableCalculator operatorTableCalculator,
        bool isDelete
    ) internal {
        if (!isDelete) {
            // Validate the operator table calculator
            require(address(operatorTableCalculator) != address(0), InvalidOperatorTableCalculator());
        }
        _operatorTableCalculators[operatorSet.key()] = operatorTableCalculator;
        emit OperatorTableCalculatorSet(operatorSet, operatorTableCalculator);
    }

    /**
     * @dev Internal function to set the operator set config for an operator set
     * @param operatorSet The operator set to set the config for
     * @param config The operator set config
     * @param isDelete Whether to delete the operator set config
     */
    function _setOperatorSetConfig(
        OperatorSet memory operatorSet,
        OperatorSetConfig memory config,
        bool isDelete
    ) internal {
        if (!isDelete) {
            // Validate the operator set config
            require(config.owner != address(0), InputAddressZero());
            require(config.maxStalenessPeriod != 0, StalenessPeriodZero());
        }
        _operatorSetConfigs[operatorSet.key()] = config;
        emit OperatorSetConfigSet(operatorSet, config);
    }

    /**
     * @dev Internal helper function to convert EnumerableSet.UintSet to uint32[]
     * @param set The EnumerableSet.UintSet to convert
     * @return result The converted uint32 array
     */
    function _getUint32Array(
        EnumerableSet.UintSet storage set
    ) internal view returns (uint32[] memory result) {
        uint256 length = set.length();
        result = new uint32[](length);
        for (uint256 i = 0; i < length; i++) {
            result[i] = uint32(set.at(i));
        }
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc ICrossChainRegistry
    function getActiveGenerationReservations()
        external
        view
        returns (OperatorSet[] memory, IOperatorTableCalculator[] memory)
    {
        uint256 length = _activeGenerationReservations.length();
        OperatorSet[] memory operatorSets = new OperatorSet[](length);
        IOperatorTableCalculator[] memory calculators = new IOperatorTableCalculator[](length);

        for (uint256 i = 0; i < length; i++) {
            bytes32 operatorSetKey = _activeGenerationReservations.at(i);
            OperatorSet memory operatorSet = OperatorSetLib.decode(operatorSetKey);

            operatorSets[i] = operatorSet;
            calculators[i] = getOperatorTableCalculator(operatorSet);
        }

        return (operatorSets, calculators);
    }

    /// @inheritdoc ICrossChainRegistry
    function getOperatorTableCalculator(
        OperatorSet memory operatorSet
    ) public view returns (IOperatorTableCalculator) {
        return _operatorTableCalculators[operatorSet.key()];
    }

    /// @inheritdoc ICrossChainRegistry
    function getOperatorSetConfig(
        OperatorSet memory operatorSet
    ) public view returns (OperatorSetConfig memory) {
        return _operatorSetConfigs[operatorSet.key()];
    }

    /// @inheritdoc ICrossChainRegistry
    function calculateOperatorTableBytes(
        OperatorSet calldata operatorSet
    ) external view returns (bytes memory) {
        return abi.encode(
            operatorSet,
            keyRegistrar.getOperatorSetCurveType(operatorSet),
            getOperatorSetConfig(operatorSet),
            getOperatorTableCalculator(operatorSet).calculateOperatorTableBytes(operatorSet)
        );
    }

    /// @inheritdoc ICrossChainRegistry
    function getActiveTransportReservations() external view returns (OperatorSet[] memory, uint32[][] memory) {
        uint256 length = _activeTransportReservations.length();
        OperatorSet[] memory operatorSets = new OperatorSet[](length);
        uint32[][] memory chainIDs = new uint32[][](length);

        for (uint256 i = 0; i < length; i++) {
            bytes32 operatorSetKey = _activeTransportReservations.at(i);
            OperatorSet memory operatorSet = OperatorSetLib.decode(operatorSetKey);

            operatorSets[i] = operatorSet;
            chainIDs[i] = getTransportDestinations(operatorSet);
        }

        return (operatorSets, chainIDs);
    }

    /// @inheritdoc ICrossChainRegistry
    function getTransportDestinations(
        OperatorSet memory operatorSet
    ) public view returns (uint32[] memory) {
        EnumerableSet.UintSet storage chainIDs = _transportDestinations[operatorSet.key()];
        uint256 length = chainIDs.length();

        // Create result array with maximum possible size
        uint32[] memory result = new uint32[](length);
        uint256 count = 0;

        // Single loop to filter whitelisted chains
        for (uint256 i = 0; i < length; i++) {
            uint32 chainID = uint32(chainIDs.at(i));
            if (_whitelistedChainIDs.contains(chainID)) {
                result[count] = chainID;
                count++;
            }
        }

        // Resize the array to the actual count using assembly
        assembly {
            mstore(result, count)
        }
        // Only return chains that are whitelisted
        return result;
    }

    /// @inheritdoc ICrossChainRegistry
    function getSupportedChains() external view returns (uint32[] memory) {
        return _getUint32Array(_whitelistedChainIDs);
    }
}
