## CrossChainRegistry

| File | Type | Proxy |
| -------- | -------- |
| [`CrossChainRegistry.sol`](../../src/contracts/multichain/CrossChainRegistry.sol) | Singleton | Transparent Proxy |

Libraries and Mixins:

| File | Notes |
| -------- | -------- |
| [`PermissionControllerMixin.sol`](../../../src/contracts/mixins/PermissionControllerMixin.sol) | account delegation |
| [`SemverMixin.sol`](../../../src/contracts/mixins/SemVerMixin.sol) | versioning |
| [`OperatorSetLib.sol`](../../src/contracts/libraries/OperatorSetLib.sol) | encode/decode operator sets |

## Overview

The `CrossChainRegistry` manages the registration/deregistration of operatorSets to the multichain protocol. The contract also exposes read-only functions to calculate an operator table, which is used offchain by the `Generator` to generate a `GlobalTableRoot` and `Transporter` to transport the operator table. See [ELIP-007](https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-007.md) for more information.

```solidity
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
```  

## Parameterization
* `SupportedChains` (via `getSupportedChains`) are `Mainnet` and `Base`
    * These are chains to which tables can be transported to
    * On Testnet, the supported chains are `Sepolia` and `Base-Sepolia`
* `TableUpdateCadence` is the frequency at which tables are *expected* to be transported to destination chains
    * When setting an operator set config, the `maxStalenessPeriod` must be either:
        * 0 (special case allowing certificates to always be valid)
        * Greater than or equal to the table update cadence
    * The table update cadence itself cannot be 0

---

## Create/Remove Generation Reservation
A generation reservation registers the operatorSet to be included in the `GlobalTableRoot` and transported to an avs-defined set of destination chains. AVSs do not have to pay for reservations.  

### `createGenerationReservation`

```solidity
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
```

Creates a generation reservation for a given `operatorSet`, which enables the operatorSet to be included in the `GlobalTableRoot` generation and transported to specified destination chains. This function sets up the complete configuration for cross-chain operations in a single transaction.

*Effects*:
* Adds the `operatorSet` to `_activeGenerationReservations`
* Sets the `operatorTableCalculator` for the `operatorSet`
* Sets the `OperatorSetConfig` containing the `owner` and `maxStalenessPeriod`
* Adds all specified `chainIDs` as transport destinations
* Emits a `GenerationReservationCreated` event
* Emits an `OperatorTableCalculatorSet` event
* Emits an `OperatorSetConfigSet` event
* Emits a `TransportDestinationChainAdded` event for each added chain

*Requirements*:
* The global paused status MUST NOT be set: `PAUSED_GENERATION_RESERVATIONS`
* Caller MUST be UAM permissioned for `operatorSet.avs`
* The `operatorSet` MUST exist in the `AllocationManager`
* A generation reservation MUST NOT already exist for the `operatorSet`
* At least one `chainID` MUST be provided
* All provided `chainIDs` MUST be whitelisted

### `removeGenerationReservation`

```solidity
/**
 * @notice Removes a generation reservation for a given operatorSet
 * @param operatorSet the operatorSet to remove
 * @dev msg.sender must be UAM permissioned for operatorSet.avs
 */
function removeGenerationReservation(
    OperatorSet calldata operatorSet
) external;
```

Removes a generation reservation for a given `operatorSet` and clears all associated storage including the operator table calculator, operator set config, and transport destinations.

*Effects*:
* Removes the `operatorTableCalculator` mapping for the `operatorSet`
* Removes the `operatorSetConfig` mapping for the `operatorSet`
* Removes all transport destinations for the `operatorSet`
* Removes the `operatorSet` from `_activeGenerationReservations`
* Emits an `OperatorTableCalculatorRemoved` event
* Emits an `OperatorSetConfigRemoved` event
* Emits a `TransportDestinationsRemoved` event
* Emits a `GenerationReservationRemoved` event

*Requirements*:
* The global paused status MUST NOT be set: `PAUSED_GENERATION_RESERVATIONS`
* Caller MUST be UAM permissioned for `operatorSet.avs`
* The `operatorSet` MUST exist in the `AllocationManager`
* A generation reservation MUST exist for the `operatorSet`

---

## Update Configuration
For a given operatorSet, an AVS can set the [`OperatorTableCalculator`](./OperatorTableCalculator.md) and `OperatorSetConfig`, which is an `owner` and `maxStalenessPeriod` transported to each chain. The [`CertificateVerifier`](../destination/CertificateVerifier.md) has more information on the `maxStalenessPeriod`. 

### `setOperatorTableCalculator`

```solidity
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
```

Updates the `operatorTableCalculator` contract for a given `operatorSet`. The `operatorTableCalculator` is deployed by the AVS and is responsible for computing the operator table bytes that will be included in cross-chain transports. For more information on the `operatorTableCalculator`, please see full documentation in the [middleware repository](https://github.com/Layr-Labs/eigenlayer-middleware/tree/dev/docs).

*Effects*:
* Updates the `_operatorTableCalculators` mapping for the `operatorSet`
* Emits an `OperatorTableCalculatorSet` event

*Requirements*:
* The global paused status MUST NOT be set: `PAUSED_OPERATOR_TABLE_CALCULATOR`
* Caller MUST be UAM permissioned for `operatorSet.avs`
* The `operatorSet` MUST exist in the `AllocationManager`
* A generation reservation MUST exist for the `operatorSet`

### `setOperatorSetConfig` 

```solidity
/**
 * @notice Sets the operatorSetConfig for a given operatorSet
 * @param operatorSet the operatorSet to set the operatorSetConfig for
 * @param config the config to set
 * @dev msg.sender must be UAM permissioned for operatorSet.avs
 * @dev operatorSet must have an active generation reservation
 * @dev The max staleness period is NOT checkpointed and is applied globally regardless of the reference timestamp of a certificate
 */
function setOperatorSetConfig(
    OperatorSet calldata operatorSet, 
    OperatorSetConfig calldata config
) external;
```

Updates the operator set configuration for a given `operatorSet`. The config contains an `owner` address and `maxStalenessPeriod` that will be transported to destination chains for use in certificate verification.

*Effects*:
* Updates the `_operatorSetConfigs` mapping with the new `config` containing:
  * `owner`: The permissioned owner of the OperatorSet on L2
  * `maxStalenessPeriod`: The maximum staleness period for the operatorSet
* Emits an `OperatorSetConfigSet` event

*Requirements*:
* The global paused status MUST NOT be set: `PAUSED_OPERATOR_SET_CONFIG`
* Caller MUST be UAM permissioned for `operatorSet.avs`
* The `operatorSet` MUST exist in the `AllocationManager`
* A generation reservation MUST exist for the `operatorSet`
* The `maxStalenessPeriod` MUST be either:
  * 0 (special case allowing certificates to always be valid)
  * Greater than or equal to the table update cadence

### `addTransportDestinations` 

```solidity
/**
 * @notice Adds destination chains to transport to
 * @param operatorSet the operatorSet to add transport destinations for
 * @param chainIDs to add transport to
 * @dev msg.sender must be UAM permissioned for operatorSet.avs
 * @dev Will create a transport reservation if one doesn't exist
 */
function addTransportDestinations(
    OperatorSet calldata operatorSet, 
    uint256[] calldata chainIDs
) external;
```

Adds destination chain IDs to the transport list for a given `operatorSet`. These chains will receive the operator table data during cross-chain transports.

*Effects*:
* Adds each `chainID` to the `_transportDestinations` set for the `operatorSet`
* Emits a `TransportDestinationChainAdded` event for each successfully added chain

*Requirements*:
* The global paused status MUST NOT be set: `PAUSED_TRANSPORT_DESTINATIONS`
* Caller MUST be UAM permissioned for `operatorSet.avs`
* The `operatorSet` MUST exist in the `AllocationManager`
* A generation reservation MUST exist for the `operatorSet`
* At least one `chainID` MUST be provided
* All provided `chainIDs` MUST be whitelisted
* Each `chainID` MUST NOT already be a transport destination

### `removeTransportDestinations`

```solidity
/**
 * @notice Removes destination chains to transport to
 * @param operatorSet the operatorSet to remove transport destinations for
 * @param chainIDs to remove transport to
 * @dev msg.sender must be UAM permissioned for operatorSet.avs
 * @dev Will remove the transport reservation if all destinations are removed
 */
function removeTransportDestinations(
    OperatorSet calldata operatorSet, 
    uint256[] calldata chainIDs
) external;
```

Removes destination chain IDs from the transport list for a given `operatorSet`. 

*Effects*:
* Removes each `chainID` from the `_transportDestinations` set for the `operatorSet`
* Emits a `TransportDestinationChainRemoved` event for each successfully removed chain

*Requirements*:
* The global paused status MUST NOT be set: `PAUSED_TRANSPORT_DESTINATIONS`
* Caller MUST be UAM permissioned for `operatorSet.avs`
* The `operatorSet` MUST exist in the `AllocationManager`
* A generation reservation MUST exist for the `operatorSet`
* Each `chainID` MUST exist in the transport destinations
* At least one transport destination MUST remain after removal (use `removeGenerationReservation` to remove all)

---

## System Configuration
The `owner` of the `CrossChainRegistry` can update the whitelisted chain IDs. If a chainID is not whitelisted, it cannot be transported to. 

### `addChainIDsToWhitelist`

```solidity
/**
 * @notice Adds chainIDs to the whitelist of chainIDs that can be transported to
 * @param chainIDs the chainIDs to add to the whitelist
 * @param operatorTableUpdaters the operatorTableUpdaters for each whitelisted chainID
 * @dev msg.sender must be the owner of the CrossChainRegistry
 */
function addChainIDsToWhitelist(uint256[] calldata chainIDs, address[] calldata operatorTableUpdaters) external;
```

Adds chain IDs to the global whitelist, enabling them as valid transport destinations. Each chain ID is associated with an operator table updater address that will be responsible for updating operator tables on that chain.

*Effects*:
* Adds each `chainID` and its corresponding `operatorTableUpdater` to the `_whitelistedChainIDs` mapping
* Emits a `ChainIDAddedToWhitelist` event for each successfully added chain

*Requirements*:
* Caller MUST be the `owner` of the contract
* The global paused status MUST NOT be set: `PAUSED_CHAIN_WHITELIST`
* The `chainIDs` and `operatorTableUpdaters` arrays MUST have the same length
* Each `chainID` MUST NOT be zero
* Each `chainID` MUST NOT already be whitelisted

### `removeChainIDsFromWhitelist`

```solidity
/**
 * @notice Removes chainIDs from the whitelist of chainIDs that can be transported to
 * @param chainIDs the chainIDs to remove from the whitelist
 * @dev msg.sender must be the owner of the CrossChainRegistry
 */
function removeChainIDsFromWhitelist(
    uint256[] calldata chainIDs
) external;
```

Removes chain IDs from the global whitelist, preventing them from being used as transport destinations. Note that existing transport destinations for operator sets will continue to filter out non-whitelisted chains when queried.

*Effects*:
* Removes each `chainID` from the `_whitelistedChainIDs` mapping
* Emits a `ChainIDRemovedFromWhitelist` event for each successfully removed chain

*Requirements*:
* Caller MUST be the `owner` of the contract
* The global paused status MUST NOT be set: `PAUSED_CHAIN_WHITELIST`
* Each `chainID` MUST be currently whitelisted 

### `setTableUpdateCadence`

```solidity
/**
 * @notice Sets the table update cadence in seconds
 * @param tableUpdateCadence the table update cadence
 * @dev msg.sender must be the owner of the CrossChainRegistry
 * @dev The table update cadence cannot be 0
 */
function setTableUpdateCadence(
    uint32 tableUpdateCadence
) external;
```

Sets the global table update cadence - the cadence at which operator tables are *expected* to be updated. This value acts as a floor for all non-zero `maxStalenessPeriod` values in operator set configurations.

*Effects*:
* Updates the `_tableUpdateCadence` storage variable
* Emits a `TableUpdateCadenceSet` event

*Requirements*:
* Caller MUST be the `owner` of the contract
* The `tableUpdateCadence` MUST be greater than 0

---

## Offchain View Functions

The `Generator` and `Transporter` use the below view functions to generate and transport tables:

1. `getActiveGenerationReservations`: Gets the operatorSets to be included in the `globalTableRoot`

```solidity
/**
 * @notice Gets the active generation reservations
 * @return An array of operatorSets with active generationReservations
 */
function getActiveGenerationReservations() external view returns (OperatorSet[] memory);
```

2. `calculateOperatorTableBytes`: Calculates the operator table bytes, which is then merkleized into the `globalTableRoot` offchain

```solidity
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
```

3. `getActiveTransportReservations`: Gets all operatorSets with active transport reservations and their destinations

```solidity
/**
 * @notice Gets the active transport reservations
 * @return An array of operatorSets with active transport reservations
 * @return An array of chainIDs that the operatorSet is configured to transport to
 */
function getActiveTransportReservations() external view returns (OperatorSet[] memory, uint256[][] memory);
```
