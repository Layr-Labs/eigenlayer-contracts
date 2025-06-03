// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "../contracts/mixins/PermissionControllerMixin.sol";
import "../contracts/mixins/SemVerMixin.sol";
import "../contracts/permissions/Pausable.sol";
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
    ReentrancyGuardUpgradeable,
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
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the CrossChainRegistry with immutable dependencies
     * @param _permissionController The permission controller for access control
     * @param _allocationManager The allocation manager for operator set validation
     * @param _pauserRegistry The pauser registry for pause functionality
     * @param _version The semantic version of the contract
     */
    constructor(
        IPermissionController _permissionController,
        IAllocationManager _allocationManager,
        IPauserRegistry _pauserRegistry,
        string memory _version
    )
        CrossChainRegistryStorage(_allocationManager)
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
        IOperatorTableCalculator operatorTableCalculator
    ) external onlyWhenNotPaused(PAUSED_GENERATION_RESERVATIONS) checkCanCall(operatorSet.avs) nonReentrant {
        bytes32 operatorSetKey = operatorSet.key();

        // Check if generation reservation already exists
        require(!_activeGenerationReservations.contains(operatorSetKey), GenerationReservationAlreadyExists());

        // Validate the operator table calculator
        require(address(operatorTableCalculator) != address(0), InvalidOperatorTableCalculator());

        // Add to active generation reservations
        _activeGenerationReservations.add(operatorSetKey);

        // Set the operator table calculator
        _setOperatorTableCalculator(operatorSet, operatorTableCalculator);

        emit GenerationReservationMade(operatorSet, operatorTableCalculator);
    }

    /// @inheritdoc ICrossChainRegistry
    function removeGenerationReservation(
        OperatorSet calldata operatorSet
    ) external onlyWhenNotPaused(PAUSED_GENERATION_RESERVATIONS) checkCanCall(operatorSet.avs) nonReentrant {
        bytes32 operatorSetKey = operatorSet.key();

        // Check if generation reservation exists
        require(_activeGenerationReservations.contains(operatorSetKey), GenerationReservationDoesNotExist());

        // Get the current operator table calculator before removal
        IOperatorTableCalculator calculator = _operatorTableCalculators[operatorSetKey];

        // Remove from active generation reservations
        _activeGenerationReservations.remove(operatorSetKey);

        // Remove the operator table calculator
        _setOperatorTableCalculator(operatorSet, IOperatorTableCalculator(address(0)));

        emit GenerationReservationRemoved(operatorSet, calculator);
    }

    /// @inheritdoc ICrossChainRegistry
    function addTransportDestination(
        OperatorSet calldata operatorSet,
        uint32 chainID
    ) external onlyWhenNotPaused(PAUSED_TRANSPORT_DESTINATIONS) checkCanCall(operatorSet.avs) nonReentrant {
        // Validate chainID
        require(chainID != 0, InvalidChainId());

        // Check if chainID is whitelisted
        require(_whitelistedChainIDs.contains(chainID), ChainIDNotWhitelisted());

        bytes32 operatorSetKey = operatorSet.key();

        // Check if already added
        require(!_transportDestinations[operatorSetKey].contains(chainID), TransportDestinationAlreadyAdded());

        // Add transport destination
        _transportDestinations[operatorSetKey].add(chainID);

        emit TransportDestinationAdded(operatorSet, chainID);
    }

    /// @inheritdoc ICrossChainRegistry
    function removeTransportDestination(
        OperatorSet calldata operatorSet,
        uint32 chainID
    ) external onlyWhenNotPaused(PAUSED_TRANSPORT_DESTINATIONS) checkCanCall(operatorSet.avs) nonReentrant {
        bytes32 operatorSetKey = operatorSet.key();

        // Check if transport destination exists
        require(_transportDestinations[operatorSetKey].contains(chainID), TransportDestinationNotFound());

        // Remove transport destination
        _transportDestinations[operatorSetKey].remove(chainID);

        emit TransportDestinationRemoved(operatorSet, chainID);
    }

    /// @inheritdoc ICrossChainRegistry
    function setOperatorTableCalculator(
        OperatorSet calldata operatorSet,
        IOperatorTableCalculator calculator
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_TABLE_CALCULATOR) checkCanCall(operatorSet.avs) nonReentrant {
        bytes32 operatorSetKey = operatorSet.key();

        // Check if generation reservation exists
        require(_activeGenerationReservations.contains(operatorSetKey), GenerationReservationDoesNotExist());

        // Validate the operator table calculator
        require(address(calculator) != address(0), InvalidOperatorTableCalculator());

        // Set the operator table calculator
        _setOperatorTableCalculator(operatorSet, calculator);
    }

    /// @inheritdoc ICrossChainRegistry
    function setOperatorSetConfig(
        OperatorSet calldata operatorSet,
        OperatorSetConfig calldata config
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_SET_CONFIG) checkCanCall(operatorSet.avs) nonReentrant {
        // Validate config
        require(config.owner != address(0), InputAddressZero());

        bytes32 operatorSetKey = operatorSet.key();

        // Set the operator set config
        _operatorSetConfigs[operatorSetKey] = config;

        emit OperatorSetConfigSet(operatorSet, config);
    }

    /// @inheritdoc ICrossChainRegistry
    function addChainIDToWhitelist(
        uint32 chainID
    ) external onlyOwner onlyWhenNotPaused(PAUSED_CHAIN_WHITELIST) nonReentrant {
        // Validate chainID
        require(chainID != 0, InvalidChainId());

        // Check if already whitelisted
        require(!_whitelistedChainIDs.contains(chainID), ChainIDAlreadyWhitelisted());

        // Add to whitelist
        _whitelistedChainIDs.add(chainID);

        emit ChainIDAddedToWhitelist(chainID);
    }

    /// @inheritdoc ICrossChainRegistry
    function removeChainIDFromWhitelist(
        uint32 chainID
    ) external onlyOwner onlyWhenNotPaused(PAUSED_CHAIN_WHITELIST) nonReentrant {
        // Check if whitelisted
        require(_whitelistedChainIDs.contains(chainID), ChainIDNotWhitelisted());

        // Remove from whitelist
        _whitelistedChainIDs.remove(chainID);

        emit ChainIDRemovedFromWhitelist(chainID);
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @dev Internal function to set the operator table calculator for an operator set
     * @param operatorSet The operator set to set the calculator for
     * @param calculator The operator table calculator contract
     */
    function _setOperatorTableCalculator(
        OperatorSet memory operatorSet,
        IOperatorTableCalculator calculator
    ) internal {
        bytes32 operatorSetKey = operatorSet.key();
        _operatorTableCalculators[operatorSetKey] = calculator;
        emit OperatorTableCalculatorSet(operatorSet, calculator);
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
    function getSupportedChains() external view returns (uint32[] memory) {
        return _getUint32Array(_whitelistedChainIDs);
    }

    /// @inheritdoc ICrossChainRegistry
    function getOperatorTableCalculator(
        OperatorSet calldata operatorSet
    ) external view returns (IOperatorTableCalculator) {
        return _operatorTableCalculators[operatorSet.key()];
    }

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
            operatorSets[i] = OperatorSetLib.decode(operatorSetKey);
            calculators[i] = _operatorTableCalculators[operatorSetKey];
        }

        return (operatorSets, calculators);
    }

    /// @inheritdoc ICrossChainRegistry
    function getTransportDestinations(
        OperatorSet calldata operatorSet
    ) external view returns (uint32[] memory) {
        return _getUint32Array(_transportDestinations[operatorSet.key()]);
    }

    /// @inheritdoc ICrossChainRegistry
    function getOperatorSetConfig(
        OperatorSet calldata operatorSet
    ) external view returns (OperatorSetConfig memory) {
        return _operatorSetConfigs[operatorSet.key()];
    }

    /// @inheritdoc ICrossChainRegistry
    function getWhitelistedChainIDs() external view returns (uint32[] memory) {
        return _getUint32Array(_whitelistedChainIDs);
    }
}
