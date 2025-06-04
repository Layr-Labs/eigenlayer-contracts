// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

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

    /// @notice Thrown when a transport reservation already exists for the operator set
    error TransportReservationAlreadyExists();

    /// @notice Thrown when a transport reservation does not exist for the operator set
    error TransportReservationDoesNotExist();

    /// @notice Thrown when a chain ID is already whitelisted
    error ChainIDAlreadyWhitelisted();

    /// @notice Thrown when a chain ID is not whitelisted
    error ChainIDNotWhitelisted();

    /// @notice Thrown when the staleness period is zero
    error StalenessPeriodZero();

    /// @notice Thrown when the operator set is not valid
    error InvalidOperatorSet();
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
    /// @notice Emitted when a generation reservation is made
    event GenerationReservationMade(OperatorSet operatorSet);

    /// @notice Emitted when a generation reservation is removed
    event GenerationReservationRemoved(OperatorSet operatorSet);

    /// @notice Emitted when an operatorTableCalculator is set
    event OperatorTableCalculatorSet(OperatorSet operatorSet, IOperatorTableCalculator operatorTableCalculator);

    /// @notice Emitted when a transport reservation is made
    event TransportReservationMade(OperatorSet operatorSet);

    /// @notice Emitted when a transport reservation is removed
    event TransportReservationRemoved(OperatorSet operatorSet);

    /// @notice Emitted when a transport destination is added
    event TransportDestinationAdded(OperatorSet operatorSet, uint32 chainID);

    /// @notice Emitted when a transport destination is removed
    event TransportDestinationRemoved(OperatorSet operatorSet, uint32 chainID);

    /// @notice Emitted when an operatorSetConfig is set
    event OperatorSetConfigSet(OperatorSet operatorSet, OperatorSetConfig config);

    /// @notice Emitted when a chainID is added to the whitelist
    event ChainIDAddedToWhitelist(uint32 chainID);

    /// @notice Emitted when a chainID is removed from the whitelist
    event ChainIDRemovedFromWhitelist(uint32 chainID);
}

interface ICrossChainRegistry is ICrossChainRegistryErrors, ICrossChainRegistryEvents {
    /**
     * @notice Initiates a generation reservation
     * @param operatorSet the operatorSet to make a reservation for
     * @param operatorTableCalculator the address of the operatorTableCalculator
     * @param config the config to set for the operatorSet
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     */
    function requestGenerationReservation(
        OperatorSet calldata operatorSet,
        IOperatorTableCalculator operatorTableCalculator,
        OperatorSetConfig calldata config
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
     * @dev operatorSet must have an active transport reservation
     */
    function setOperatorSetConfig(OperatorSet calldata operatorSet, OperatorSetConfig calldata config) external;

    /**
     * @notice Initiates a transport reservation
     * @param operatorSet the operatorSet to make a reservation for
     * @param chainIDs the chainIDs to transport to
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     */
    function requestTransportReservation(OperatorSet calldata operatorSet, uint32[] calldata chainIDs) external;

    /**
     * @notice Removes a transport reservation for a given operatorSet
     * @param operatorSet the operatorSet to remove
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     */
    function removeTransportReservation(
        OperatorSet calldata operatorSet
    ) external;

    /**
     * @notice Adds a destination chain to transport to
     * @param chainIDs to add transport to
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     * @dev operatorSet must have an active transport reservation
     */
    function addTransportDestinations(OperatorSet calldata operatorSet, uint32[] calldata chainIDs) external;

    /**
     * @notice Removes a destination chain to transport to
     * @param chainIDs to remove transport to
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     * @dev operatorSet must have an active transport reservation
     */
    function removeTransportDestinations(OperatorSet calldata operatorSet, uint32[] calldata chainIDs) external;

    /**
     * @notice Adds a chainID to the whitelist of chainIDs that can be transported to
     * @param chainID the chainID to add to the whitelist
     * @dev msg.sender must be the owner of the CrossChainRegistry
     */
    function addChainIDToWhitelist(
        uint32 chainID
    ) external;

    /**
     * @notice Removes a chainID from the whitelist of chainIDs that can be transported to
     * @param chainID the chainID to remove from the whitelist
     * @dev msg.sender must be the owner of the CrossChainRegistry
     */
    function removeChainIDFromWhitelist(
        uint32 chainID
    ) external;

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     * @notice Gets the active generation reservations
     * @return An array of operatorSets with active generationReservations
     * @return An array of the corresponding operatorTableCalculators
     */
    function getActiveGenerationReservations()
        external
        view
        returns (OperatorSet[] memory, IOperatorTableCalculator[] memory);

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
     * @notice Gets the active transport reservations
     * @return An array of operatorSets with active transport reservations
     * @return An array of chainIDs that the operatorSet is configured to transport to
     */
    function getActiveTransportReservations() external view returns (OperatorSet[] memory, uint32[][] memory);

    /**
     * @notice Gets the transport destinations for a given operatorSet
     * @param operatorSet the operatorSet to get the transport destinations for
     * @return An array of chainIDs that the operatorSet is configured to transport to
     */
    function getTransportDestinations(
        OperatorSet memory operatorSet
    ) external view returns (uint32[] memory);

    /**
     * @notice Gets the list of chains that are supported by the CrossChainRegistry
     * @return An array of chainIDs that are supported by the CrossChainRegistry
     */
    function getSupportedChains() external view returns (uint32[] memory);
}
