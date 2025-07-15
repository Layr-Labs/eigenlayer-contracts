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

    /// @notice Thrown when a chain ID is already whitelisted
    error ChainIDAlreadyWhitelisted();

    /// @notice Thrown when a chain ID is not whitelisted
    error ChainIDNotWhitelisted();

    /// @notice Thrown when the operator set is not valid
    error InvalidOperatorSet();

    /// @notice Thrown when the chainIDs array is empty
    error EmptyChainIDsArray();

    /// @notice Thrown when the lengths between two arrays are not the same
    error ArrayLengthMismatch();

    /// @notice Thrown when the staleness period set by an operatorSet is invalid
    error InvalidStalenessPeriod();

    /// @notice Thrown when the table update cadence is invalid
    error InvalidTableUpdateCadence();
}

interface ICrossChainRegistryTypes {
    /**
     * @notice A per-operatorSet configuration struct that is transported from the CrossChainRegistry on L1.
     * @param owner the permissioned owner of the OperatorSet on L2 that can call the CertificateVerifier specific setters
     * @param maxStalenessPeriod the maximum staleness period of the operatorSet
     *
     * @dev A staleness period of 0 allows for certificates to be verified against any timestamp in the past
     * @dev Staleness periods should not be greater than 0 and less than the update cadence of the `OperatorTables`, since
     *      certificates would be unable to be validated against. The update cadence is communicated off-chain
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

    /// @notice Emitted when an operatorTableCalculator is removed, when a generation reservation is removed
    event OperatorTableCalculatorRemoved(OperatorSet operatorSet);

    /// @notice Emitted when an operatorSetConfig is set
    event OperatorSetConfigSet(OperatorSet operatorSet, OperatorSetConfig config);

    /// @notice Emitted when an operatorSetConfig is removed, when a generation reservation is removed
    event OperatorSetConfigRemoved(OperatorSet operatorSet);

    /// @notice Emitted when a chainID is added to the whitelist
    event ChainIDAddedToWhitelist(uint256 chainID, address operatorTableUpdater);

    /// @notice Emitted when a chainID is removed from the whitelist
    event ChainIDRemovedFromWhitelist(uint256 chainID);

    /// @notice Emitted when the table update cadence is set
    event TableUpdateCadenceSet(uint32 tableUpdateCadence);
}

interface ICrossChainRegistry is ICrossChainRegistryErrors, ICrossChainRegistryEvents {
    /**
     * @notice Creates a generation reservation
     * @param operatorSet the operatorSet to make a reservation for
     * @param operatorTableCalculator the address of the operatorTableCalculator
     * @param config the config to set for the operatorSet
     * @dev msg.sender must be UAM permissioned for operatorSet.avs
     * @dev Once a generation reservation is created, the operator table will be transported to all chains that are whitelisted
     */
    function createGenerationReservation(
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
     * @dev operatorSet must have an active generation reservation
     * @dev The max staleness period is NOT checkpointed and is applied globally regardless of the reference timestamp of a certificate
     */
    function setOperatorSetConfig(OperatorSet calldata operatorSet, OperatorSetConfig calldata config) external;

    /**
     * @notice Adds chainIDs to the whitelist of chainIDs that are transported to by the multichain protocol
     * @param chainIDs the chainIDs to add to the whitelist
     * @param operatorTableUpdaters the operatorTableUpdaters for each whitelisted chainID
     * @dev msg.sender must be the owner of the CrossChainRegistry
     */
    function addChainIDsToWhitelist(uint256[] calldata chainIDs, address[] calldata operatorTableUpdaters) external;

    /**
     * @notice Removes chainIDs from the whitelist of chainIDs
     * @param chainIDs the chainIDs to remove from the whitelist
     * @dev msg.sender must be the owner of the CrossChainRegistry
     */
    function removeChainIDsFromWhitelist(
        uint256[] calldata chainIDs
    ) external;

    /**
     * @notice Sets the table update cadence in seconds
     * @param tableUpdateCadence the table update cadence
     * @dev msg.sender must be the owner of the CrossChainRegistry
     * @dev The table update cadence cannot be 0
     */
    function setTableUpdateCadence(
        uint32 tableUpdateCadence
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
     * @notice Gets the list of chains that are supported by the CrossChainRegistry
     * @return An array of chainIDs that are supported by the CrossChainRegistry
     * @return An array of operatorTableUpdaters corresponding to each chainID
     */
    function getSupportedChains() external view returns (uint256[] memory, address[] memory);

    /**
     * @notice Gets the table update cadence
     * @return The table update cadence
     * @dev The table update cadence is applicable to all chains
     */
    function getTableUpdateCadence() external view returns (uint32);
}
