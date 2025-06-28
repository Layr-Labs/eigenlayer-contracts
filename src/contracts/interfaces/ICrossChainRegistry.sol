// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import "./IOperatorTableCalculator.sol";

interface ICrossChainRegistryErrors {
    /// @notice Thrown when the chainId is invalid
    error InvalidChainId();

    /// @notice Thrown when a generation reservation already exists for the operator set
    error GenerationReservationAlreadyExists();

    /// @notice Thrown when a generation reservation does not exist for the operator set
    error GenerationReservationDoesNotExist();

    /// @notice Thrown when the operator table calculator address is invalid
    error InvalidOperatorTableCalculator();

    /// @notice Thrown when a transport destination is already added for the operator set
    error TransportDestinationAlreadyAdded();

    /// @notice Thrown when a transport destination is not found for the operator set
    error TransportDestinationNotFound();

    /// @notice Thrown when a chain ID is already whitelisted
    error ChainIDAlreadyWhitelisted();

    /// @notice Thrown when a chain ID is not whitelisted
    error ChainIDNotWhitelisted();

    /// @notice Thrown when the staleness period is zero
    error StalenessPeriodZero();

    /// @notice Thrown when the operator set is not valid
    error InvalidOperatorSet();

    /// @notice Thrown when the chainIDs array is empty
    error EmptyChainIDsArray();

    /// @notice Thrown when a at least one transport destination is required
    error RequireAtLeastOneTransportDestination();

    /// @notice Thrown when the storage is not cleared
    error NeedToDelete();

    /// @notice Thrown when the lengths between two arrays are not the same
    error ArrayLengthMismatch();
}

interface ICrossChainRegistryTypes {
    /**
     * @notice A per-operatorSet configuration struct that is transported from the CrossChainRegistry on L1.
     * @param owner the permissioned owner of the OperatorSet on L2 that can call the CertificateVerifier specific setters
     * @param maxStalenessPeriod the maximum staleness period of the operatorSet
     */
    struct OperatorSetConfig {
        address owner;
        uint32 maxStalenessPeriod;
    }
}

interface ICrossChainRegistryEvents is ICrossChainRegistryTypes {
    /// @notice Emitted when a generation reservation is created
    event GenerationReservationCreated(OperatorSet operatorSet);

    /// @notice Emitted when a generation reservation is removed
    event GenerationReservationRemoved(OperatorSet operatorSet);

    /// @notice Emitted when an operatorTableCalculator is set
    event OperatorTableCalculatorSet(OperatorSet operatorSet, IOperatorTableCalculator operatorTableCalculator);

    /// @notice Emitted when an operatorSetConfig is set
    event OperatorSetConfigSet(OperatorSet operatorSet, OperatorSetConfig config);

    /// @notice Emitted when a transport destination is added
    event TransportDestinationAdded(OperatorSet operatorSet, uint256 chainID);

    /// @notice Emitted when a transport destination is removed
    event TransportDestinationRemoved(OperatorSet operatorSet, uint256 chainID);

    /// @notice Emitted when a chainID is added to the whitelist
    event ChainIDAddedToWhitelist(uint256 chainID, address operatorTableUpdater);

    /// @notice Emitted when a chainID is removed from the whitelist
    event ChainIDRemovedFromWhitelist(uint256 chainID);
}

interface ICrossChainRegistry is ICrossChainRegistryErrors, ICrossChainRegistryEvents {
    /**
     * @notice Creates a generation reservation
     * @param operatorSet the operatorSet to make a reservation for
     * @param operatorTableCalculator the address of the operatorTableCalculator
     * @param config the config to set for the operatorSet
     * @param chainIDs the chainIDs to add as transport destinations
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     */
    function createGenerationReservation(
        OperatorSet calldata operatorSet,
        IOperatorTableCalculator operatorTableCalculator,
        OperatorSetConfig calldata config,
        uint256[] calldata chainIDs
    ) external;

    /**
     * @notice Removes a generation reservation for a given operatorSet
     * @param operatorSet the operatorSet to remove
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     */
    function removeGenerationReservation(
        OperatorSet calldata operatorSet
    ) external;

    /**
     * @notice Sets the operatorTableCalculator for the operatorSet
     * @param operatorSet the operatorSet whose operatorTableCalculator is desired to be set
     * @param operatorTableCalculator the contract to call to calculate the operator table
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     * @dev operatorSet must have an active reservation
     */
    function setOperatorTableCalculator(
        OperatorSet calldata operatorSet,
        IOperatorTableCalculator operatorTableCalculator
    ) external;

    /**
     * @notice Sets the operatorSetConfig for a given operatorSet
     * @param operatorSet the operatorSet to set the operatorSetConfig for
     * @param config the config to set
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     * @dev operatorSet must have an active generation reservation
     */
    function setOperatorSetConfig(OperatorSet calldata operatorSet, OperatorSetConfig calldata config) external;

    /**
     * @notice Adds destination chains to transport to
     * @param operatorSet the operatorSet to add transport destinations for
     * @param chainIDs to add transport to
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     * @dev Will create a transport reservation if one doesn't exist
     */
    function addTransportDestinations(OperatorSet calldata operatorSet, uint256[] calldata chainIDs) external;

    /**
     * @notice Removes destination chains to transport to
     * @param operatorSet the operatorSet to remove transport destinations for
     * @param chainIDs to remove transport to
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     * @dev Will remove the transport reservation if all destinations are removed
     */
    function removeTransportDestinations(OperatorSet calldata operatorSet, uint256[] calldata chainIDs) external;

    /**
     * @notice Adds chainIDs to the whitelist of chainIDs that can be transported to
     * @param chainIDs the chainIDs to add to the whitelist
     * @param operatorTableUpdaters the operatorTableUpdaters for each whitelisted chainID
     * @dev msg.sender must be the owner of the CrossChainRegistry
     */
    function addChainIDsToWhitelist(uint256[] calldata chainIDs, address[] calldata operatorTableUpdaters) external;

    /**
     * @notice Removes chainIDs from the whitelist of chainIDs that can be transported to
     * @param chainIDs the chainIDs to remove from the whitelist
     * @dev msg.sender must be the owner of the CrossChainRegistry
     */
    function removeChainIDsFromWhitelist(
        uint256[] calldata chainIDs
    ) external;

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     * @notice Gets the active generation reservations
     * @return An array of operatorSets with active generationReservations
     */
    function getActiveGenerationReservations() external view returns (OperatorSet[] memory);

    /**
     * @notice Gets the operatorTableCalculator for a given operatorSet
     * @param operatorSet the operatorSet to get the operatorTableCalculator for
     * @return The operatorTableCalculator for the given operatorSet
     */
    function getOperatorTableCalculator(
        OperatorSet memory operatorSet
    ) external view returns (IOperatorTableCalculator);

    /**
     * @notice Gets the operatorSetConfig for a given operatorSet
     * @param operatorSet the operatorSet to get the operatorSetConfig for
     * @return The operatorSetConfig for the given operatorSet
     */
    function getOperatorSetConfig(
        OperatorSet memory operatorSet
    ) external view returns (OperatorSetConfig memory);

    /**
     * @notice Calculates the operatorTableBytes for a given operatorSet
     * @param operatorSet the operatorSet to calculate the operator table for
     * @return the encoded operatorTableBytes containing:
     *         - operatorSet details
     *         - curve type from KeyRegistrar
     *         - operator set configuration
     *         - calculated operator table from the calculator contract
     * @dev This function aggregates data from multiple sources for cross-chain transport
     */
    function calculateOperatorTableBytes(
        OperatorSet calldata operatorSet
    ) external view returns (bytes memory);

    /**
     * @notice Gets the active transport reservations
     * @return An array of operatorSets with active transport reservations
     * @return An array of chainIDs that the operatorSet is configured to transport to
     */
    function getActiveTransportReservations() external view returns (OperatorSet[] memory, uint256[][] memory);

    /**
     * @notice Gets the transport destinations for a given operatorSet
     * @param operatorSet the operatorSet to get the transport destinations for
     * @return An array of chainIDs that the operatorSet is configured to transport to
     */
    function getTransportDestinations(
        OperatorSet memory operatorSet
    ) external view returns (uint256[] memory);

    /**
     * @notice Gets the list of chains that are supported by the CrossChainRegistry
     * @return An array of chainIDs that are supported by the CrossChainRegistry
     * @return An array of operatorTableUpdaters corresponding to each chainID
     */
    function getSupportedChains() external view returns (uint256[] memory, address[] memory);
}
