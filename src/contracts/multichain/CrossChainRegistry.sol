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

    modifier checkHasActiveGenerationReservation(
        OperatorSet calldata operatorSet
    ) {
        require(hasActiveGenerationReservation(operatorSet), GenerationReservationDoesNotExist());
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
        OperatorSetConfig calldata config
    )
        external
        onlyWhenNotPaused(PAUSED_GENERATION_RESERVATIONS)
        checkCanCall(operatorSet.avs)
        isValidOperatorSet(operatorSet)
    {
        // Validate the the KeyType has been set in the KeyRegistrar for the OperatorSet
        require(keyRegistrar.getOperatorSetCurveType(operatorSet) != IKeyRegistrarTypes.CurveType.NONE, KeyTypeNotSet());

        // Add to active generation reservations
        require(_activeGenerationReservations.add(operatorSet.key()), GenerationReservationAlreadyExists());
        emit GenerationReservationCreated(operatorSet);

        // Set the operator table calculator
        _setOperatorTableCalculator(operatorSet, operatorTableCalculator);
        // Set the operator set config
        _setOperatorSetConfig(operatorSet, config);
    }

    /// @inheritdoc ICrossChainRegistry
    function removeGenerationReservation(
        OperatorSet calldata operatorSet
    )
        external
        onlyWhenNotPaused(PAUSED_GENERATION_RESERVATIONS)
        checkCanCall(operatorSet.avs)
        isValidOperatorSet(operatorSet)
        checkHasActiveGenerationReservation(operatorSet)
    {
        bytes32 operatorSetKey = operatorSet.key();

        // Clear all storage for the operator set
        // 1. Remove the operator table calculator
        delete _operatorTableCalculators[operatorSetKey];
        emit OperatorTableCalculatorRemoved(operatorSet);

        // 2. Remove the operator set config
        delete _operatorSetConfigs[operatorSetKey];
        emit OperatorSetConfigRemoved(operatorSet);

        // 3. Remove from active generation reservations
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
        checkHasActiveGenerationReservation(operatorSet)
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
        checkHasActiveGenerationReservation(operatorSet)
    {
        // Set the operator set config
        _setOperatorSetConfig(operatorSet, config);
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
    function hasActiveGenerationReservation(
        OperatorSet memory operatorSet
    ) public view returns (bool) {
        return _activeGenerationReservations.contains(operatorSet.key());
    }

    /// @inheritdoc ICrossChainRegistry
    function getActiveGenerationReservationsByRange(
        uint256 startIndex,
        uint256 endIndex
    ) external view returns (OperatorSet[] memory) {
        require(startIndex <= endIndex, InvalidRange());
        require(endIndex <= _activeGenerationReservations.length(), InvalidEndIndex());

        uint256 length = endIndex - startIndex;
        OperatorSet[] memory operatorSets = new OperatorSet[](length);

        for (uint256 i = 0; i < length; i++) {
            bytes32 operatorSetKey = _activeGenerationReservations.at(startIndex + i);
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

    /// @inheritdoc ICrossChainRegistry
    function getActiveGenerationReservationCount() external view returns (uint256) {
        return _activeGenerationReservations.length();
    }
}
