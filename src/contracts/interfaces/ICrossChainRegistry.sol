// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {OperatorSet} from "../libraries/OperatorSetLib.sol";
import "./IOperatorTableCalculator.sol";

interface ICrossChainRegistryErrors {
    /// @notice Thrown when the chainId is invalid
    /// @dev Error code: 0x7a47c9a2
    error InvalidChainId();

    /// @notice Thrown when a generation reservation already exists for the operator set
    /// @dev Error code: 0x18834615
    error GenerationReservationAlreadyExists();

    /// @notice Thrown when a generation reservation does not exist for the operator set
    /// @dev Error code: 0x9a575d52
    error GenerationReservationDoesNotExist();

    /// @notice Thrown when a chain ID is already whitelisted
    /// @dev Error code: 0x497ec636
    error ChainIDAlreadyWhitelisted();

    /// @notice Thrown when a chain ID is not whitelisted
    /// @dev Error code: 0xb3f92ba1
    error ChainIDNotWhitelisted();

    /// @notice Thrown when the operator set is not valid
    /// @dev Error code: 0x7ec5c154
    /// @dev We require valid operator sets (existing in AllocationManager) to ensure cross-chain operations are performed on legitimate operator configurations
    error InvalidOperatorSet();

    /// @notice Thrown when the key type is not set for the operatorSet
    /// @dev Error code: 0xe57cacbd
    /// @dev We require operator sets to have configured curve types so the multichain protocol can properly decode the operator table
    error KeyTypeNotSet();

    /// @notice Thrown when the chainIDs array is empty
    /// @dev Error code: 0x8631a075
    error EmptyChainIDsArray();

    /// @notice Thrown when the lengths between two arrays are not the same
    /// @dev Error code: 0xa24a13a6
    error ArrayLengthMismatch();

    /// @notice Thrown when the staleness period set by an operatorSet is invalid
    /// @dev Error code: 0x5c8c9062
    /// @dev We enforce valid staleness periods (0 or >= table update cadence) to ensure certificates are not prematurely invalidated between table updates
    error InvalidStalenessPeriod();

    /// @notice Thrown when the table update cadence is invalid
    /// @dev Error code: 0xb6cc70d8
    error InvalidTableUpdateCadence();

    /// @notice Thrown when the range is invalid for the `getActiveGenerationReservations` function
    /// @dev A valid range is defined as `startIndex` <= `endIndex` in `getActiveGenerationReservationsByRange`
    /// @dev Error Code: 0x561ce9bb
    error InvalidRange();

    /// @notice Thrown when the end index is invalid for the `getActiveGenerationReservations` function
    /// @dev A valid end index is defined as `endIndex` <= `getActiveGenerationReservationCount()` in `getActiveGenerationReservationsByRange`
    /// @dev Error Code: 0xb68d84c0
    error InvalidEndIndex();
}

interface ICrossChainRegistryTypes {
    /**
     * @notice A per-operatorSet configuration struct that is transported from the CrossChainRegistry on L1.
     * @param owner the permissioned owner of the OperatorSet on L2 that can be used by downstream contracts to authorize actions
     * @param maxStalenessPeriod the maximum staleness period of the operatorSet, in seconds
     *
     * @dev A `maxStalenessPeriod` of 0 completely removes staleness checks, allowing certificates to be validated regardless of their timestamp
     * @dev A nonzero `maxStalenessPeriod` has a floor of the table update cadence, which is the frequency at which operator tables are expected
     *      to be updated. The table update cadence is set by the owner of the `CrossChainRegistry`
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

/// @notice The CrossChainRegistry allows AVSs to register their operatorSets and the associated operator table for cross-chain transport by the multichain protocol
interface ICrossChainRegistry is ICrossChainRegistryErrors, ICrossChainRegistryEvents {
    /**
     * @notice Creates a generation reservation, which transports the operator table of an operatorSet to all whitelisted chains
     * @param operatorSet the operatorSet to make a reservation for
     * @param operatorTableCalculator the address of the operatorTableCalculator. This contract is deployed (or a template is used) by the AVS
     *                                to calculate the stake weights for the operatorSet. See `IOperatorTableCalculator` for more details
     * @param config the config to set for the operatorSet, which includes the owner of the operatorSet and the max staleness period
     * @dev Tables are transported at a cadence of `tableUpdateCadence` seconds. The `maxStalenessPeriod` is used to determine the maximum
     * @dev msg.sender must be an authorized caller for operatorSet.avs
     * @dev Once a generation reservation is created, the operator table will be transported to all chains that are whitelisted
     * @dev It is expected that the AVS has:
     *      - Deployed or is using a generalizable `OperatorTableCalculator` to calculate its operator's stake weights
     *      - Set the `KeyType` for the operatorSet in the `KeyRegistrar`, even if the AVS is not using the `KeyRegistrar` for operator key management
     *           - Valid Key Types are given in the `IKeyRegistrarTypes.CurveType` enum. The `KeyType` must not be `NONE`
     *      - Created an operatorSet in the `AllocationManager`
     * @dev Reverts for:
     *      - CurrentlyPaused: Generation reservations are paused
     *      - InvalidPermissions: Caller is not an authorized caller for operatorSet.avs
     *      - InvalidOperatorSet: The operatorSet does not exist in the AllocationManager
     *      - KeyTypeNotSet: The key type is not set for the operatorSet in the KeyRegistrar
     *      - GenerationReservationAlreadyExists: A generation reservation already exists for the operatorSet
     *      - InvalidStalenessPeriod: The maxStalenessPeriod is invalid
     * @dev Emits the following events:
     *      - GenerationReservationCreated: When the generation reservation is successfully created
     *      - OperatorTableCalculatorSet: When the operator table calculator is set for the operatorSet
     *      - OperatorSetConfigSet: When the operator set config is set for the operatorSet
     */
    function createGenerationReservation(
        OperatorSet calldata operatorSet,
        IOperatorTableCalculator operatorTableCalculator,
        OperatorSetConfig calldata config
    ) external;

    /**
     * @notice Removes a generation reservation for a given operatorSet
     * @param operatorSet the operatorSet to remove
     * @dev msg.sender must be an authorized caller for operatorSet.avs
     * @dev Reverts for:
     *      - CurrentlyPaused: Generation reservations are paused
     *      - InvalidPermissions: Caller is not an authorized caller for operatorSet.avs
     *      - InvalidOperatorSet: The operatorSet does not exist in the AllocationManager
     *      - GenerationReservationDoesNotExist: A generation reservation does not exist for the operatorSet
     * @dev Emits the following events:
     *      - OperatorTableCalculatorRemoved: When the operator table calculator is removed
     *      - OperatorSetConfigRemoved: When the operator set config is removed
     *      - GenerationReservationRemoved: When the generation reservation is removed
     */
    function removeGenerationReservation(
        OperatorSet calldata operatorSet
    ) external;

    /**
     * @notice Sets the operatorTableCalculator for the operatorSet
     * @param operatorSet the operatorSet whose operatorTableCalculator is desired to be set
     * @param operatorTableCalculator the contract to call to calculate the operator table
     * @dev msg.sender must be an authorized caller for operatorSet.avs
     * @dev operatorSet must have an active reservation
     * @dev Reverts for:
     *      - CurrentlyPaused: Setting the operatorTableCalculator is paused
     *      - InvalidPermissions: Caller is not an authorized caller for operatorSet.avs
     *      - InvalidOperatorSet: The operatorSet does not exist in the AllocationManager
     *      - GenerationReservationDoesNotExist: A generation reservation does not exist for the operatorSet
     * @dev Emits the following events:
     *      - OperatorTableCalculatorSet: When the operator table calculator is successfully set
     */
    function setOperatorTableCalculator(
        OperatorSet calldata operatorSet,
        IOperatorTableCalculator operatorTableCalculator
    ) external;

    /**
     * @notice Sets the operatorSetConfig for a given operatorSet
     * @param operatorSet the operatorSet to set the operatorSetConfig for
     * @param config the config to set, which includes the owner of the operatorSet and the max staleness period
     * @dev msg.sender must be an authorized caller for operatorSet.avs
     * @dev operatorSet must have an active generation reservation
     * @dev The max staleness period is NOT checkpointed and is applied globally regardless of the reference timestamp of a certificate
     * @dev Reverts for:
     *      - CurrentlyPaused: Setting the operatorSetConfig is paused
     *      - InvalidPermissions: Caller is not an authorized caller for operatorSet.avs
     *      - InvalidOperatorSet: The operatorSet does not exist in the AllocationManager
     *      - GenerationReservationDoesNotExist: A generation reservation does not exist for the operatorSet
     *      - InvalidStalenessPeriod: The maxStalenessPeriod is invalid
     * @dev Emits the following events:
     *      - OperatorSetConfigSet: When the operator set config is successfully set
     */
    function setOperatorSetConfig(OperatorSet calldata operatorSet, OperatorSetConfig calldata config) external;

    /**
     * @notice Adds chainIDs to the whitelist of chainIDs that are transported to by the multichain protocol
     * @param chainIDs the chainIDs to add to the whitelist
     * @param operatorTableUpdaters the operatorTableUpdaters for each whitelisted chainID
     * @dev msg.sender must be the owner of the CrossChainRegistry
     * @dev Reverts for:
     *      - "Ownable: caller is not the owner": Caller is not the owner of the contract
     *      - CurrentlyPaused: Chain whitelisting is paused
     *      - ArrayLengthMismatch: The chainIDs and operatorTableUpdaters arrays have different lengths
     *      - InvalidChainId: Any chainID is zero
     *      - ChainIDAlreadyWhitelisted: Any chainID is already whitelisted
     * @dev Emits the following events:
     *      - ChainIDAddedToWhitelist: When each chainID is successfully added to the whitelist
     */
    function addChainIDsToWhitelist(uint256[] calldata chainIDs, address[] calldata operatorTableUpdaters) external;

    /**
     * @notice Removes chainIDs from the whitelist of chainIDs
     * @param chainIDs the chainIDs to remove from the whitelist
     * @dev msg.sender must be the owner of the CrossChainRegistry
     * @dev Reverts for:
     *      - "Ownable: caller is not the owner": Caller is not the owner of the contract
     *      - CurrentlyPaused: Chain whitelisting is paused
     *      - ChainIDNotWhitelisted: Any chainID is not currently whitelisted
     * @dev Emits the following events:
     *      - ChainIDRemovedFromWhitelist: When each chainID is successfully removed from the whitelist
     */
    function removeChainIDsFromWhitelist(
        uint256[] calldata chainIDs
    ) external;

    /**
     * @notice Sets the table update cadence in seconds. This is the frequency at which operator tables are expected to be updated on all destination chains
     * @param tableUpdateCadence the table update cadence
     * @dev msg.sender must be the owner of the CrossChainRegistry
     * @dev The table update cadence cannot be 0
     * @dev Reverts for:
     *      - "Ownable: caller is not the owner": Caller is not the owner of the contract
     *      - InvalidTableUpdateCadence: The tableUpdateCadence is zero
     * @dev Emits the following events:
     *      - TableUpdateCadenceSet: When the table update cadence is successfully set
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
     * @notice Gets the active generation reservations by range
     * @param startIndex the start index of the range, inclusive
     * @param endIndex the end index of the range, exclusive
     * @return An array of operatorSets with active generationReservations
     * @dev Reverts for:
     *      - InvalidRange: startIndex is greater than endIndex
     *      - InvalidEndIndex: endIndex is greater than the length of the active generation reservations array
     */
    function getActiveGenerationReservationsByRange(
        uint256 startIndex,
        uint256 endIndex
    ) external view returns (OperatorSet[] memory);

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
     * @dev Reverts when the call to the operatorTableCalculator contract call fails
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
     * @notice Gets the table update cadence, in seconds
     * @return The table update cadence, in seconds
     * @dev The table update cadence is applicable to all whitelisted chains, and is the
     *      frequency at which operator tables are expected to be updated on all destination chains
     */
    function getTableUpdateCadence() external view returns (uint32);

    /**
     * @notice Gets the number of active generation reservations
     * @return The number of active generation reservations
     * @dev This function can be used in conjunction with the paginated version of `getActiveGenerationReservations`
     *      to iterate over all active generation reservations
     */
    function getActiveGenerationReservationCount() external view returns (uint256);
}
