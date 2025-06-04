// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import "./IOperatorTableCalculator.sol";

interface ICrossChainRegistryErrors {
    /// @notice Thrown when the chainId is invalid
    error InvalidChainId();
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

interface ICrossChainRegistryEvents {
    /// @notice Emitted when a generation reservation is made
    event GenerationReservationMade(OperatorSet operatorSet, IOperatorTableCalculator operatorTableCalculator);

    /// @notice Emitted when a generation reservation is removed
    event GenerationReservationRemoved(OperatorSet operatorSet, IOperatorTableCalculator operatorTableCalculator);

    /// @notice Emitted when a transport destination is added
    event TransportDestinationAdded(OperatorSet operatorSet, uint32 chainID);

    /// @notice Emitted when a transport destination is removed
    event TransportDestinationRemoved(OperatorSet operatorSet, uint32 chainID);

    /// @notice Emitted when a chainID is added to the whitelist
    event ChainIDAddedToWhitelist(uint32 chainID);

    /// @notice Emitted when a chainID is removed from the whitelist
    event ChainIDRemovedFromWhitelist(uint32 chainID);
}

interface ICrossChainRegistry is ICrossChainRegistryErrors, ICrossChainRegistryTypes, ICrossChainRegistryEvents {
    /**
     * @notice Initiates a generation reservation
     * @param operatorSet the operatorSet to make a reservation for
     * @param operatorTableCalculator the address of the operatorTableCalculator
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     */
    function requestGenerationReservation(OperatorSet calldata operatorSet, address operatorTableCalculator) external;

    /**
     * @notice Removes a generation reservation for a given operatorSet
     * @param operatorSet the operatorSet to remove
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     */
    function removeGenerationReservation(
        OperatorSet calldata operatorSet
    ) external;

    /**
     * @notice Adds a destination chain to transport to
     * @param chainID to add transport to
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     */
    function addTransportDestination(OperatorSet calldata operatorSet, uint32 chainID) external;

    /**
     * @notice Removes a destination chain to transport to
     * @param chainID to remove transport to
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     */
    function removeTransportDestination(OperatorSet calldata operatorSet, uint32 chainID) external;

    /**
     * @notice Sets the operatorTableCalculator for the operatorSet
     * @param operatorSet the operatorSet whose operatorTableCalculator is desired to be set
     * @param calculator the contract to call to calculate the operator table
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     * @dev operatorSet must have an active reservation
     */
    function setOperatorTableCalculator(
        OperatorSet calldata operatorSet,
        IOperatorTableCalculator calculator
    ) external;

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
     * @notice Gets the list of chains that are supported by the CrossChainRegistry
     * @return An array of chainIDs that are supported by the CrossChainRegistry
     */
    function getSupportedChains() external view returns (uint32[] memory);

    /**
     * @notice Gets the operatorTableCalculator for a given operatorSet
     * @param operatorSet the operatorSet to get the operatorTableCalculator for
     * @return The operatorTableCalculator for the given operatorSet
     */
    function getOperatorTableCalculator(
        OperatorSet calldata operatorSet
    ) external view returns (IOperatorTableCalculator);

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
     * @notice Gets the transport destinations for a given operatorSet
     * @param operatorSet the operatorSet to get the transport destinations for
     * @return An array of chainIDs that are transport destinations for the given operatorSet
     */
    function getTransportDestinations(
        OperatorSet calldata operatorSet
    ) external view returns (uint32[] memory);
}
