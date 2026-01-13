---
name: interface-lint
description: Format and lint Solidity interface files following EigenLayer conventions. Use when the user asks to format an interface, add documentation to an interface, or create a new interface. Ensures proper organization with Errors/Events/Types sub-interfaces.
allowed-tools: Read, Glob, Grep, Edit, Write, Bash(cast:*)
---

# Interface Lint

Format Solidity interface files following EigenLayer's established conventions for organization, documentation, and error code generation.

## Interface Structure

Every interface file should contain four interface contracts in this order:

1. **`I{ContractName}Errors`** - All custom errors
2. **`I{ContractName}Types`** - All structs and enums
3. **`I{ContractName}Events`** - All events (inherits Types for struct access)
4. **`I{ContractName}`** - Main interface (inherits Errors and Events)

```solidity
// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// Imports

interface I{ContractName}Errors {
    // All errors with error codes
}

interface I{ContractName}Types {
    // All structs and enums
}

interface I{ContractName}Events is I{ContractName}Types {
    // All events
}

interface I{ContractName} is I{ContractName}Errors, I{ContractName}Events {
    // All function declarations
}
```

## Error Documentation

Each error MUST include:
1. `@notice` describing when the error is thrown
2. `@dev Error code: 0x...` with the 4-byte selector from `cast sig`

### Generating Error Codes

Use `cast sig` to generate the error code:

```bash
cast sig "InvalidOperatorSet()"
# Output: 0x7ec5c154
```

### Error Format

```solidity
interface IContractNameErrors {
    /// @notice Thrown when the operator set is not valid
    /// @dev Error code: 0x7ec5c154
    error InvalidOperatorSet();

    /// @notice Thrown when the chainId is invalid
    /// @dev Error code: 0x7a47c9a2
    error InvalidChainId();

    /// @notice Thrown when the key type is not set for the operatorSet
    /// @dev Error code: 0xe57cacbd
    /// @dev Additional context about why this is required
    error KeyTypeNotSet();
}
```

## Types Documentation

Each struct/enum MUST include:
1. `@notice` with a brief description
2. `@param` for each field in structs

```solidity
interface IContractNameTypes {
    /// @notice A per-operatorSet configuration struct
    /// @param owner the permissioned owner of the OperatorSet
    /// @param maxStalenessPeriod the maximum staleness period in seconds
    struct OperatorSetConfig {
        address owner;
        uint32 maxStalenessPeriod;
    }

    /// @notice Represents the status of an operator's registration
    /// @param registered Whether the operator is currently registered
    /// @param slashableUntil Block until which the operator remains slashable
    struct RegistrationStatus {
        bool registered;
        uint32 slashableUntil;
    }
}
```

## Event Documentation

Each event MUST include a singular `@notice` describing when the event is emitted:

```solidity
interface IContractNameEvents is IContractNameTypes {
    /// @notice Emitted when a generation reservation is created
    event GenerationReservationCreated(OperatorSet operatorSet);

    /// @notice Emitted when an operator set config is set
    event OperatorSetConfigSet(OperatorSet operatorSet, OperatorSetConfig config);

    /// @notice Emitted when a chainID is added to the whitelist
    event ChainIDAddedToWhitelist(uint256 chainID, address operatorTableUpdater);
}
```

## Function Documentation

Each function in the main interface MUST include:

1. `@notice` - What the function does
2. `@param` - Description for each parameter
3. `@return` - Description for each return value (for view functions)
4. `@dev` - Additional context (optional, but include caller requirements)
5. `@dev Reverts for:` - List ALL revert conditions
6. `@dev Emits the following events:` - List ALL events emitted

### Function Documentation Format

```solidity
interface IContractName is IContractNameErrors, IContractNameEvents {
    /// @notice Creates a generation reservation for cross-chain transport
    /// @param operatorSet the operatorSet to make a reservation for
    /// @param operatorTableCalculator the calculator contract address
    /// @param config the config containing owner and staleness period
    /// @dev msg.sender must be an authorized caller for operatorSet.avs
    /// @dev Reverts for:
    ///      - CurrentlyPaused: Generation reservations are paused
    ///      - InvalidPermissions: Caller is not authorized
    ///      - InvalidOperatorSet: The operatorSet does not exist
    ///      - GenerationReservationAlreadyExists: Reservation already exists
    ///      - InvalidStalenessPeriod: The maxStalenessPeriod is invalid
    /// @dev Emits the following events:
    ///      - GenerationReservationCreated: When the reservation is created
    ///      - OperatorTableCalculatorSet: When the calculator is set
    ///      - OperatorSetConfigSet: When the config is set
    function createGenerationReservation(
        OperatorSet calldata operatorSet,
        IOperatorTableCalculator operatorTableCalculator,
        OperatorSetConfig calldata config
    ) external;

    /// @notice Gets the operator set config
    /// @param operatorSet the operatorSet to query
    /// @return The OperatorSetConfig for the given operatorSet
    /// @dev You should check if an operatorSet has an active reservation first
    function getOperatorSetConfig(
        OperatorSet memory operatorSet
    ) external view returns (OperatorSetConfig memory);
}
```

## Revert Conditions Format

List each revert condition with the error name and when it occurs:

```solidity
/// @dev Reverts for:
///      - CurrentlyPaused: Generation reservations are paused
///      - InvalidPermissions: Caller is not an authorized caller for operatorSet.avs
///      - InvalidOperatorSet: The operatorSet does not exist in the AllocationManager
///      - GenerationReservationAlreadyExists: A generation reservation already exists
```

For Ownable errors, use the string format:
```solidity
/// @dev Reverts for:
///      - "Ownable: caller is not the owner": Caller is not the owner
```

## Events Emitted Format

List each event with a brief description of when it's emitted:

```solidity
/// @dev Emits the following events:
///      - GenerationReservationCreated: When the reservation is successfully created
///      - OperatorTableCalculatorSet: When the calculator is set for the operatorSet
///      - OperatorSetConfigSet: When the config is set for the operatorSet
```

## View Functions

View functions typically don't revert (except for input validation) and don't emit events. Document them simpler:

```solidity
/// @notice Gets the active generation reservations
/// @return An array of operatorSets with active generationReservations
function getActiveGenerationReservations() external view returns (OperatorSet[] memory);

/// @notice Gets reservations by range for pagination
/// @param startIndex the start index of the range, inclusive
/// @param endIndex the end index of the range, exclusive
/// @return An array of operatorSets in the specified range
/// @dev Reverts for:
///      - InvalidRange: startIndex is greater than endIndex
///      - InvalidEndIndex: endIndex exceeds array length
function getActiveGenerationReservationsByRange(
    uint256 startIndex,
    uint256 endIndex
) external view returns (OperatorSet[] memory);
```

## Complete Example

Reference: `src/contracts/interfaces/ICrossChainRegistry.sol`

This file demonstrates the full pattern with:
- Errors interface with `cast sig` error codes
- Types interface with documented structs
- Events interface inheriting Types
- Main interface with full function documentation

## Checklist for Linting an Interface

1. ☐ Split into Errors, Types, Events, and main interface contracts
2. ☐ Errors have `@notice` and `@dev Error code: 0x...` (use `cast sig`)
3. ☐ Types have `@notice` and `@param` for each field
4. ☐ Events have singular `@notice` describing when emitted
5. ☐ Events interface inherits Types interface
6. ☐ Main interface inherits Errors and Events interfaces
7. ☐ Functions have `@notice`, `@param`, and `@return` (for views)
8. ☐ Functions list ALL revert conditions under `@dev Reverts for:`
9. ☐ Functions list ALL events under `@dev Emits the following events:`
10. ☐ View functions document reverts only when they can revert

## Generating All Error Codes

To generate error codes for an entire interface:

```bash
# For each error in the interface, run:
cast sig "ErrorName()"
cast sig "ErrorName(uint256)"  # Include params if error has them

# Example output:
# 0x7ec5c154  # InvalidOperatorSet()
# 0x7a47c9a2  # InvalidChainId()
```

## Common Error Patterns

| Error | Typical Code | Usage |
|-------|--------------|-------|
| `InvalidOperatorSet()` | `0x7ec5c154` | OperatorSet doesn't exist |
| `InvalidPermissions()` | varies | Caller not authorized |
| `InvalidCaller()` | varies | Wrong msg.sender |
| `ArrayLengthMismatch()` | `0xa24a13a6` | Input arrays have different lengths |