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
    using EnumerableMap for EnumerableMap.UintToAddressMap;
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

    modifier hasActiveGenerationReservation(
        OperatorSet calldata operatorSet
    ) {
        require(_activeGenerationReservations.contains(operatorSet.key()), GenerationReservationDoesNotExist());
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
     * @param initialTableUpdateCadence The initial table update cadence
     * @param initialPausedStatus The initial paused status bitmap
     */
    function initialize(
        address initialOwner,
        uint32 initialTableUpdateCadence,
        uint256 initialPausedStatus
    ) external initializer {
        _transferOwnership(initialOwner);
        _setTableUpdateCadence(initialTableUpdateCadence);
        _setPausedStatus(initialPausedStatus);
    }

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

    /// @inheritdoc ICrossChainRegistry
    function createGenerationReservation(
        OperatorSet calldata operatorSet,
        IOperatorTableCalculator operatorTableCalculator,
        OperatorSetConfig calldata config,
        uint256[] calldata chainIDs
    )
        external
        onlyWhenNotPaused(PAUSED_GENERATION_RESERVATIONS)
        checkCanCall(operatorSet.avs)
        isValidOperatorSet(operatorSet)
    {
        // Add to active generation reservations
        require(_activeGenerationReservations.add(operatorSet.key()), GenerationReservationAlreadyExists());
        emit GenerationReservationCreated(operatorSet);

        // Set the operator table calculator
        _setOperatorTableCalculator(operatorSet, operatorTableCalculator);
        // Set the operator set config
        _setOperatorSetConfig(operatorSet, config);
        // Add transport destinations
        _addTransportDestinations(operatorSet, chainIDs);
    }

    /// @inheritdoc ICrossChainRegistry
    function removeGenerationReservation(
        OperatorSet calldata operatorSet
    )
        external
        onlyWhenNotPaused(PAUSED_GENERATION_RESERVATIONS)
        checkCanCall(operatorSet.avs)
        isValidOperatorSet(operatorSet)
        hasActiveGenerationReservation(operatorSet)
    {
        bytes32 operatorSetKey = operatorSet.key();

        // Clear all storage for the operator set
        // 1. Remove the operator table calculator
        delete _operatorTableCalculators[operatorSetKey];
        emit OperatorTableCalculatorRemoved(operatorSet);

        // 2. Remove the operator set config
        delete _operatorSetConfigs[operatorSetKey];
        emit OperatorSetConfigRemoved(operatorSet);

        // 3. Remove all transport destinations
        delete _transportDestinations[operatorSetKey];
        emit TransportDestinationsRemoved(operatorSet);

        // 4. Remove from active generation reservations
        _activeGenerationReservations.remove(operatorSetKey);
        emit GenerationReservationRemoved(operatorSet);
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
        hasActiveGenerationReservation(operatorSet)
    {
        // Set the operator table calculator
        _setOperatorTableCalculator(operatorSet, operatorTableCalculator);
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
        hasActiveGenerationReservation(operatorSet)
    {
        // Set the operator set config
        _setOperatorSetConfig(operatorSet, config);
    }

    /// @inheritdoc ICrossChainRegistry
    function addTransportDestinations(
        OperatorSet calldata operatorSet,
        uint256[] calldata chainIDs
    )
        external
        onlyWhenNotPaused(PAUSED_TRANSPORT_DESTINATIONS)
        checkCanCall(operatorSet.avs)
        isValidOperatorSet(operatorSet)
        hasActiveGenerationReservation(operatorSet)
    {
        _addTransportDestinations(operatorSet, chainIDs);
    }

    /// @inheritdoc ICrossChainRegistry
    function removeTransportDestinations(
        OperatorSet calldata operatorSet,
        uint256[] calldata chainIDs
    )
        external
        onlyWhenNotPaused(PAUSED_TRANSPORT_DESTINATIONS)
        checkCanCall(operatorSet.avs)
        isValidOperatorSet(operatorSet)
        hasActiveGenerationReservation(operatorSet)
    {
        _removeTransportDestinations(operatorSet, chainIDs);
    }

    /// @inheritdoc ICrossChainRegistry
    function addChainIDsToWhitelist(
        uint256[] calldata chainIDs,
        address[] calldata operatorTableUpdaters
    ) external onlyOwner onlyWhenNotPaused(PAUSED_CHAIN_WHITELIST) {
        require(chainIDs.length == operatorTableUpdaters.length, ArrayLengthMismatch());

        for (uint256 i = 0; i < chainIDs.length; i++) {
            uint256 chainID = chainIDs[i];

            // Validate chainID
            require(chainID != 0, InvalidChainId());

            // Add to whitelist
            require(_whitelistedChainIDs.set(chainID, operatorTableUpdaters[i]), ChainIDAlreadyWhitelisted());

            emit ChainIDAddedToWhitelist(chainID, operatorTableUpdaters[i]);
        }
    }

    /// @inheritdoc ICrossChainRegistry
    function removeChainIDsFromWhitelist(
        uint256[] calldata chainIDs
    ) external onlyOwner onlyWhenNotPaused(PAUSED_CHAIN_WHITELIST) {
        for (uint256 i = 0; i < chainIDs.length; i++) {
            uint256 chainID = chainIDs[i];

            // Remove from whitelist
            require(_whitelistedChainIDs.remove(chainID), ChainIDNotWhitelisted());

            emit ChainIDRemovedFromWhitelist(chainID);
        }
    }

    /// @inheritdoc ICrossChainRegistry
    function setTableUpdateCadence(
        uint32 tableUpdateCadence
    ) external onlyOwner {
        _setTableUpdateCadence(tableUpdateCadence);
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
     */
    function _setOperatorTableCalculator(
        OperatorSet memory operatorSet,
        IOperatorTableCalculator operatorTableCalculator
    ) internal {
        _operatorTableCalculators[operatorSet.key()] = operatorTableCalculator;
        emit OperatorTableCalculatorSet(operatorSet, operatorTableCalculator);
    }

    /**
     * @dev Internal function to set the operator set config for an operator set
     * @param operatorSet The operator set to set the config for
     * @param config The operator set config
     * @dev The 0 staleness period is special case and is allowed, since it allows for certificates to ALWAYS be valid
     */
    function _setOperatorSetConfig(OperatorSet memory operatorSet, OperatorSetConfig memory config) internal {
        require(
            config.maxStalenessPeriod == 0 || config.maxStalenessPeriod >= _tableUpdateCadence, InvalidStalenessPeriod()
        );
        _operatorSetConfigs[operatorSet.key()] = config;
        emit OperatorSetConfigSet(operatorSet, config);
    }

    /**
     * @dev Internal function to add transport destinations for an operator set
     * @param operatorSet The operator set to add destinations for
     * @param chainIDs The chain IDs to add as destinations
     */
    function _addTransportDestinations(OperatorSet memory operatorSet, uint256[] memory chainIDs) internal {
        // Validate chainIDs array is not empty. This is to prevent users from
        // creating a generation reservation with no destinations.
        require(chainIDs.length > 0, EmptyChainIDsArray());

        bytes32 operatorSetKey = operatorSet.key();

        for (uint256 i = 0; i < chainIDs.length; i++) {
            uint256 chainID = chainIDs[i];

            // Check if chainID is whitelisted
            require(_whitelistedChainIDs.contains(chainID), ChainIDNotWhitelisted());

            // Add transport destination
            require(_transportDestinations[operatorSetKey].add(chainID), TransportDestinationAlreadyAdded());

            emit TransportDestinationChainAdded(operatorSet, chainID);
        }
    }

    /**
     * @dev Internal function to remove transport destinations for an operator set
     * @param operatorSet The operator set to remove destinations from
     * @param chainIDs The chain IDs to remove as destinations
     */
    function _removeTransportDestinations(OperatorSet memory operatorSet, uint256[] memory chainIDs) internal {
        bytes32 operatorSetKey = operatorSet.key();

        for (uint256 i = 0; i < chainIDs.length; i++) {
            uint256 chainID = chainIDs[i];

            // Remove transport destination
            require(_transportDestinations[operatorSetKey].remove(chainID), TransportDestinationNotFound());

            emit TransportDestinationChainRemoved(operatorSet, chainID);
        }

        // Ensure that at least one destination remains
        // If a user wants to remove all destinations, they should call `removeGenerationReservation` instead
        require(_transportDestinations[operatorSetKey].length() > 0, RequireAtLeastOneTransportDestination());
    }

    /**
     * @dev Internal function to set the table update cadence
     * @param tableUpdateCadence the table update cadence
     * @dev The table update cadence cannot be 0 as that is special-cased to allow for certificates to ALWAYS be valid
     */
    function _setTableUpdateCadence(
        uint32 tableUpdateCadence
    ) internal {
        require(tableUpdateCadence > 0, InvalidTableUpdateCadence());
        _tableUpdateCadence = tableUpdateCadence;
        emit TableUpdateCadenceSet(tableUpdateCadence);
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc ICrossChainRegistry
    function getActiveGenerationReservations() external view returns (OperatorSet[] memory) {
        uint256 length = _activeGenerationReservations.length();
        OperatorSet[] memory operatorSets = new OperatorSet[](length);

        for (uint256 i = 0; i < length; i++) {
            bytes32 operatorSetKey = _activeGenerationReservations.at(i);
            OperatorSet memory operatorSet = OperatorSetLib.decode(operatorSetKey);

            operatorSets[i] = operatorSet;
        }

        return operatorSets;
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
    function getActiveTransportReservations() external view returns (OperatorSet[] memory, uint256[][] memory) {
        uint256 length = _activeGenerationReservations.length();
        OperatorSet[] memory operatorSets = new OperatorSet[](length);
        uint256[][] memory chainIDs = new uint256[][](length);

        for (uint256 i = 0; i < length; i++) {
            bytes32 operatorSetKey = _activeGenerationReservations.at(i);
            OperatorSet memory operatorSet = OperatorSetLib.decode(operatorSetKey);

            operatorSets[i] = operatorSet;
            chainIDs[i] = getTransportDestinations(operatorSet);
        }

        return (operatorSets, chainIDs);
    }

    /// @inheritdoc ICrossChainRegistry
    function getTransportDestinations(
        OperatorSet memory operatorSet
    ) public view returns (uint256[] memory) {
        EnumerableSet.UintSet storage chainIDs = _transportDestinations[operatorSet.key()];
        uint256 length = chainIDs.length();

        // Create result array with maximum possible size
        uint256[] memory result = new uint256[](length);
        uint256 count = 0;

        // Single loop to filter whitelisted chains
        for (uint256 i = 0; i < length; i++) {
            uint256 chainID = chainIDs.at(i);
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
    function getSupportedChains() external view returns (uint256[] memory, address[] memory) {
        uint256 length = _whitelistedChainIDs.length();
        uint256[] memory chainIDs = new uint256[](length);
        address[] memory operatorTableUpdaters = new address[](length);

        for (uint256 i = 0; i < length; i++) {
            (uint256 chainID, address operatorTableUpdater) = _whitelistedChainIDs.at(i);
            chainIDs[i] = chainID;
            operatorTableUpdaters[i] = operatorTableUpdater;
        }

        return (chainIDs, operatorTableUpdaters);
    }

    /// @inheritdoc ICrossChainRegistry
    function getTableUpdateCadence() external view returns (uint32) {
        return _tableUpdateCadence;
    }
}
