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

        // Add to active generation reservations
        _activeGenerationReservations.add(operatorSetKey);
        emit GenerationReservationMade(operatorSet);

        // Set the operator table calculator
        _setOperatorTableCalculator(operatorSet, operatorTableCalculator, false);
    }

    /// @inheritdoc ICrossChainRegistry
    function removeGenerationReservation(
        OperatorSet calldata operatorSet
    ) external onlyWhenNotPaused(PAUSED_GENERATION_RESERVATIONS) checkCanCall(operatorSet.avs) nonReentrant {
        bytes32 operatorSetKey = operatorSet.key();

        // Check if generation reservation exists
        require(_activeGenerationReservations.contains(operatorSetKey), GenerationReservationDoesNotExist());

        // Remove from active generation reservations
        _activeGenerationReservations.remove(operatorSetKey);
        emit GenerationReservationRemoved(operatorSet);

        // Remove the operator table calculator
        _setOperatorTableCalculator(operatorSet, IOperatorTableCalculator(address(0)), true);
    }

    /// @inheritdoc ICrossChainRegistry
    function setOperatorTableCalculator(
        OperatorSet calldata operatorSet,
        IOperatorTableCalculator operatorTableCalculator
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_TABLE_CALCULATOR) checkCanCall(operatorSet.avs) nonReentrant {
        bytes32 operatorSetKey = operatorSet.key();

        // Check if generation reservation exists
        require(_activeGenerationReservations.contains(operatorSetKey), GenerationReservationDoesNotExist());

        // Set the operator table calculator
        _setOperatorTableCalculator(operatorSet, operatorTableCalculator, false);
    }

    /// @inheritdoc ICrossChainRegistry
    function requestTransportReservation(
        OperatorSet calldata operatorSet,
        uint32[] calldata chainIDs,
        OperatorSetConfig calldata config
    ) external onlyWhenNotPaused(PAUSED_TRANSPORT_RESERVATIONS) checkCanCall(operatorSet.avs) nonReentrant {
        bytes32 operatorSetKey = operatorSet.key();

        // Check if transport reservation already exists
        require(!_activeTransportReservations.contains(operatorSetKey), TransportReservationAlreadyExists());

        // Add to active transport reservations
        _activeTransportReservations.add(operatorSetKey);
        emit TransportReservationMade(operatorSet);

        // Add transport destinations
        for (uint256 i = 0; i < chainIDs.length; i++) {
            _addTransportDestination(operatorSet, chainIDs[i]);
        }

        // Set the operator set config
        _setOperatorSetConfig(operatorSet, config, false);
    }

    /// @inheritdoc ICrossChainRegistry
    function removeTransportReservation(
        OperatorSet calldata operatorSet
    ) external onlyWhenNotPaused(PAUSED_TRANSPORT_RESERVATIONS) checkCanCall(operatorSet.avs) nonReentrant {
        bytes32 operatorSetKey = operatorSet.key();

        // Check if transport reservation exists
        require(_activeTransportReservations.contains(operatorSetKey), TransportReservationDoesNotExist());

        // Remove from active transport reservations
        _activeTransportReservations.remove(operatorSetKey);
        emit TransportReservationRemoved(operatorSet);

        // Get the transport destinations
        EnumerableSet.UintSet storage chainIDs = _transportDestinations[operatorSetKey];
        // Remove the transport destinations
        for (uint256 i = 0; i < chainIDs.length(); i++) {
            _removeTransportDestination(operatorSet, uint32(chainIDs.at(i)));
        }

        // Remove the operator set config
        _setOperatorSetConfig(operatorSet, OperatorSetConfig(address(0), 0), true);
    }

    /// @inheritdoc ICrossChainRegistry
    function addTransportDestinations(
        OperatorSet calldata operatorSet,
        uint32[] calldata chainIDs
    ) external onlyWhenNotPaused(PAUSED_TRANSPORT_DESTINATIONS) checkCanCall(operatorSet.avs) nonReentrant {
        bytes32 operatorSetKey = operatorSet.key();

        // Check if transport reservation exists
        require(_activeTransportReservations.contains(operatorSetKey), TransportReservationDoesNotExist());

        for (uint256 i = 0; i < chainIDs.length; i++) {
            _addTransportDestination(operatorSet, chainIDs[i]);
        }
    }

    /// @inheritdoc ICrossChainRegistry
    function removeTransportDestinations(
        OperatorSet calldata operatorSet,
        uint32[] calldata chainIDs
    ) external onlyWhenNotPaused(PAUSED_TRANSPORT_DESTINATIONS) checkCanCall(operatorSet.avs) nonReentrant {
        bytes32 operatorSetKey = operatorSet.key();

        // Check if transport reservation exists
        require(_activeTransportReservations.contains(operatorSetKey), TransportReservationDoesNotExist());

        for (uint256 i = 0; i < chainIDs.length; i++) {
            _removeTransportDestination(operatorSet, chainIDs[i]);
        }
    }

    /// @inheritdoc ICrossChainRegistry
    function setOperatorSetConfig(
        OperatorSet calldata operatorSet,
        OperatorSetConfig calldata config
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_SET_CONFIG) checkCanCall(operatorSet.avs) nonReentrant {
        bytes32 operatorSetKey = operatorSet.key();

        // Check if transport reservation exists
        require(_activeTransportReservations.contains(operatorSetKey), TransportReservationDoesNotExist());

        // Set the operator set config
        _setOperatorSetConfig(operatorSet, config, false);
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
        // Validate chainID
        require(chainID != 0, InvalidChainId());

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
        bytes32 operatorSetKey = operatorSet.key();
        _operatorTableCalculators[operatorSetKey] = operatorTableCalculator;
        emit OperatorTableCalculatorSet(operatorSet, operatorTableCalculator);
    }

    /**
     * @dev Internal function to add a transport destination for an operator set
     * @param operatorSet The operator set to add the transport destination for
     * @param chainID The chain ID to add the transport destination for
     */
    function _addTransportDestination(OperatorSet calldata operatorSet, uint32 chainID) internal {
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

    /**
     * @dev Internal function to remove a transport destination for an operator set
     * @param operatorSet The operator set to remove the transport destination for
     * @param chainID The chain ID to remove the transport destination for
     */
    function _removeTransportDestination(OperatorSet calldata operatorSet, uint32 chainID) internal {
        // Validate chainID
        require(chainID != 0, InvalidChainId());

        bytes32 operatorSetKey = operatorSet.key();

        // Check if transport destination exists
        require(_transportDestinations[operatorSetKey].contains(chainID), TransportDestinationNotFound());

        // Remove transport destination
        _transportDestinations[operatorSetKey].remove(chainID);

        emit TransportDestinationRemoved(operatorSet, chainID);
    }

    function _setOperatorSetConfig(
        OperatorSet calldata operatorSet,
        OperatorSetConfig memory config,
        bool isDelete
    ) internal {
        if (!isDelete) {
            // Validate the operator set config
            require(config.owner != address(0), InputAddressZero());
            require(config.maxStalenessPeriod != 0, StalenessPeriodZero());
        }
        bytes32 operatorSetKey = operatorSet.key();
        _operatorSetConfigs[operatorSetKey] = config;
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
    function getSupportedChains() external view returns (uint32[] memory) {
        return _getUint32Array(_whitelistedChainIDs);
    }

    /// @inheritdoc ICrossChainRegistry
    function getOperatorTableCalculator(
        OperatorSet memory operatorSet
    ) public view returns (IOperatorTableCalculator) {
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
            OperatorSet memory operatorSet = OperatorSetLib.decode(operatorSetKey);

            operatorSets[i] = operatorSet;
            calculators[i] = getOperatorTableCalculator(operatorSet);
        }

        return (operatorSets, calculators);
    }

    /// @inheritdoc ICrossChainRegistry
    function getTransportDestinations(
        OperatorSet memory operatorSet
    ) public view returns (uint32[] memory) {
        // Only return chains that are whitelisted
        EnumerableSet.UintSet storage chainIDs = _transportDestinations[operatorSet.key()];

        // First count how many chain IDs are whitelisted
        uint256 count = 0;
        for (uint256 i = 0; i < chainIDs.length(); i++) {
            if (_whitelistedChainIDs.contains(uint32(chainIDs.at(i)))) {
                count++;
            }
        }

        // Create result array with correct size
        uint32[] memory result = new uint32[](count);
        uint256 j = 0;
        for (uint256 i = 0; i < chainIDs.length(); i++) {
            uint32 chainID = uint32(chainIDs.at(i));
            if (_whitelistedChainIDs.contains(chainID)) {
                result[j] = chainID;
                j++;
            }
        }
        return result;
    }

    /// @inheritdoc ICrossChainRegistry
    function getOperatorSetConfig(
        OperatorSet memory operatorSet
    ) public view returns (OperatorSetConfig memory) {
        return _operatorSetConfigs[operatorSet.key()];
    }

    /// @inheritdoc ICrossChainRegistry
    function getActiveTransportReservations()
        external
        view
        returns (OperatorSet[] memory, uint32[][] memory, OperatorSetConfig[] memory)
    {
        uint256 length = _activeTransportReservations.length();
        OperatorSet[] memory operatorSets = new OperatorSet[](length);
        uint32[][] memory chainIDs = new uint32[][](length);
        OperatorSetConfig[] memory configs = new OperatorSetConfig[](length);

        for (uint256 i = 0; i < length; i++) {
            bytes32 operatorSetKey = _activeTransportReservations.at(i);
            OperatorSet memory operatorSet = OperatorSetLib.decode(operatorSetKey);

            operatorSets[i] = operatorSet;
            chainIDs[i] = getTransportDestinations(operatorSet);
            configs[i] = getOperatorSetConfig(operatorSet);
        }

        return (operatorSets, chainIDs, configs);
    }
}
