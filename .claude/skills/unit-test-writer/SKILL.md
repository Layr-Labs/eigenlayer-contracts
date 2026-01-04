---
name: unit-test-writer
description: Write Solidity unit tests for EigenLayer contracts. Use when the user asks to write tests, add test coverage, create unit tests, or test a function. Follows project conventions with per-function test contracts and mock dependencies.
allowed-tools: Read, Glob, Grep, Edit, Write, Bash(forge:*)
---

# Unit Test Writer

Write comprehensive unit tests for EigenLayer Solidity contracts following the project's established conventions.

## Test File Structure

Each test file follows this structure:

```solidity
// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// Import the contract under test
import "src/contracts/path/to/ContractUnderTest.sol";
// Import the appropriate test setup
import "src/test/utils/EigenLayerUnitTestSetup.sol";
// Import any required mocks
import "src/test/mocks/SomeMock.sol";

/// @title ContractUnderTestUnitTests
/// @notice Base contract for all ContractUnderTest unit tests
contract ContractUnderTestUnitTests is EigenLayerUnitTestSetup, IContractErrors, IContractEvents, IContractTypes {
    // Test state variables
    ContractUnderTest contractUnderTest;

    function setUp() public virtual override {
        EigenLayerUnitTestSetup.setUp();
        // Deploy and initialize contract under test
        // Set up default test values
        // Configure mocks
    }

    // Helper functions
}

/// @title ContractUnderTestUnitTests_functionName
/// @notice Unit tests for ContractUnderTest.functionName
contract ContractUnderTestUnitTests_functionName is ContractUnderTestUnitTests {
    function setUp() public override {
        super.setUp();
        // Function-specific setup
    }

    // Revert tests
    function test_Revert_Paused() public { }
    function test_Revert_NotPermissioned() public { }
    function test_Revert_InvalidInput() public { }

    // Success tests
    function test_functionName_Success() public { }

    // Fuzz tests
    function testFuzz_functionName_VariableName(uint256 value) public { }
}
```

## Test Contract Naming Convention

- **Base contract**: `{ContractName}UnitTests`
- **Per-function contracts**: `{ContractName}UnitTests_{functionName}`

## Test Function Naming Convention

| Pattern | Purpose |
|---------|---------|
| `test_Revert_Paused` | Test function reverts when paused |
| `test_Revert_NotPermissioned` | Test function reverts for unauthorized callers |
| `test_Revert_NotOwner` | Test function reverts for non-owner |
| `test_Revert_Invalid{Thing}` | Test function reverts for invalid inputs |
| `test_Revert_{ErrorName}` | Test function reverts with specific error |
| `test_{functionName}_Success` | Test successful execution (happy path) |
| `test_{functionName}_{Scenario}` | Test specific scenario |
| `testFuzz_{functionName}_{Scenario}` | Fuzz test with bounded variable |

## Coverage Requirements

For each function, ensure:

### 1. Revert Cases (Branch Coverage)
- **Pausable functions**: Test `CurrentlyPaused` revert
- **Permissioned functions**: Test `InvalidPermissions` or `NotOwner` revert
- **Input validation**: Test each require/revert condition
- **State checks**: Test precondition failures

### 2. Happy Path (Line Coverage)
- Call function with valid inputs
- Verify emitted events with `cheats.expectEmit(true, true, true, true, address(contract))`
- Verify state changes with assertions

### 3. Fuzz Tests
- Use `bound()` to constrain fuzz inputs to valid ranges
- Test edge cases and variable inputs
- For complex tests or when standard fuzz inputs are too slow, use the `Randomness` type from `src/test/utils/Random.sol` 

## Using Mocks

External contract calls should use mocks from `src/test/mocks/`:

```solidity
// In setUp()
allocationManagerMock.setIsOperatorSet(operatorSet, true);

// Mock pattern: Mocks expose setters to control return values
mock.setSomeValue(expectedValue);
// Then the contract under test calls mock.getSomeValue() and gets expectedValue
```

### Creating New Mock Contracts

If a mock doesn't exist in `src/test/mocks/`, create one following this pattern:

**Location**: `src/test/mocks/{ContractName}Mock.sol`

**Structure**:

```solidity
// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "src/contracts/interfaces/IContractName.sol";

contract ContractNameMock is Test {
    receive() external payable {}
    fallback() external payable {}

    // Storage for mock return values
    mapping(bytes32 => bool) public _someMapping;

    // Setter to configure mock behavior
    function setSomeValue(bytes32 key, bool value) external {
        _someMapping[key] = value;
    }

    // Interface method that returns configured value
    function someValue(bytes32 key) external view returns (bool) {
        return _someMapping[key];
    }
}
```

**Key principles**:
1. **Inherit from `Test`** - Gives access to cheatcodes if needed
2. **Include `receive()` and `fallback()`** - Allows the mock to receive ETH and handle unknown calls gracefully
3. **Only implement what's needed** - Add functions on a need-to-implement basis as tests require them
4. **Prefix storage with `_`** - Use `_variableName` for internal mock storage to distinguish from interface getters
5. **Create setters for each value** - Pattern: `setX()` to configure, `getX()` or `x()` to return the configured value 

## Test Setup Inheritance

Choose the appropriate base setup:
- `EigenLayerUnitTestSetup` - Standard core contract tests
- `EigenLayerMultichainUnitTestSetup` - Multichain/cross-chain tests

## Event Verification

```solidity
// Expect event emission BEFORE the call
cheats.expectEmit(true, true, true, true, address(contractUnderTest));
emit SomeEvent(param1, param2);

// Make the call
contractUnderTest.someFunction(param1, param2);
```

## State Verification

```solidity
// After the call, verify state
assertEq(contract.getValue(), expectedValue, "Value mismatch");
assertTrue(contract.isEnabled(), "Should be enabled");
assertFalse(contract.isDisabled(), "Should not be disabled");
```

## Fuzz Test Patterns

### Standard Fuzz Tests (using `bound()`)

```solidity
function testFuzz_functionName_Amount(uint256 amount) public {
    // Bound to valid range
    amount = bound(amount, 1, type(uint128).max);

    // Or for uint8
    uint8 smallValue = uint8(bound(value, 1, 100));

    // Test with bounded value
    contractUnderTest.functionName(amount);

    // Verify
    assertEq(contractUnderTest.getAmount(), amount, "Amount mismatch");
}
```

### Randomness generation for Fuzz Tests

For tests that need multiple random values or complex random data structures, use the `Randomness` type from `src/test/utils/Random.sol`. This is preferred when:
- You need multiple correlated random values
- Standard fuzz inputs reject too many cases
- You need random arrays or complex types (addresses, bytes32, OperatorSets, etc.)

**Setup**: The base test contract must have the `rand` modifier and `random()` helper (already in `EigenLayerUnitTestSetup`):

```solidity
modifier rand(Randomness r) {
    r.set();
    _;
}

function random() internal returns (Randomness) {
    return Randomness.wrap(Random.SEED).shuffle();
}
```

**Usage Pattern**:

```solidity
function testFuzz_functionName_ComplexScenario(Randomness r) public rand(r) {
    // Generate random values using r.Type() or r.Type(min, max)
    address staker = r.Address();
    bytes32 salt = r.Bytes32();
    uint256 shares = r.Uint256(1, MAX_SHARES);
    uint64 magnitude = r.Uint64(1, WAD);
    uint32 count = r.Uint32(1, 32);
    bool flag = r.Boolean();

    // Use random values in test
    contractUnderTest.someFunction(staker, shares);

    // Verify behavior
    assertEq(contractUnderTest.getShares(staker), shares);
}
```

**Available Random Methods**:

| Method | Description |
|--------|-------------|
| `r.Uint256()` | Random uint256 |
| `r.Uint256(min, max)` | Random uint256 in range [min, max) |
| `r.Uint128()`, `r.Uint64()`, `r.Uint32()` | Smaller uint types |
| `r.Int256()`, `r.Int128()`, etc. | Signed integers |
| `r.Address()` | Random non-zero address |
| `r.Bytes32()` | Random bytes32 |
| `r.Boolean()` | Random true/false |
| `r.StrategyArray(len)` | Array of random strategy addresses |
| `r.StakerArray(len)` | Array of random staker addresses |
| `r.Uint256Array(len, min, max)` | Array of random uint256 values |

**Helper Functions for Complex Random Data**:

When you need multiple correlated random values (e.g., deposit/withdrawal amounts), create helper functions:

```solidity
/// @notice Generate correlated random amounts for deposits and withdrawals
function _fuzzDepositWithdrawalAmounts(Randomness r, uint32 numStrategies)
    internal
    returns (uint[] memory depositAmounts, uint[] memory withdrawalAmounts)
{
    depositAmounts = new uint[](numStrategies);
    withdrawalAmounts = new uint[](numStrategies);
    for (uint i = 0; i < numStrategies; i++) {
        depositAmounts[i] = r.Uint256(1, MAX_STRATEGY_SHARES);
        // Withdrawal must be <= deposit
        withdrawalAmounts[i] = r.Uint256(1, depositAmounts[i]);
    }
}

// Usage in test:
function testFuzz_queueWithdrawals(Randomness r) public rand(r) {
    uint32 numStrategies = r.Uint32(1, 5);
    (uint[] memory deposits, uint[] memory withdrawals) = 
        _fuzzDepositWithdrawalAmounts(r, numStrategies);
    // ... rest of test
}
```

## Example: Complete Test Contract

Reference: `src/test/unit/CrossChainRegistryUnit.t.sol`

This file demonstrates:
- Base test contract with setUp and helpers
- Per-function test contracts
- Comprehensive revert testing
- Event verification
- State verification
- Fuzz testing

## Checklist Before Writing Tests

1. Read the contract under test to understand all functions
2. Identify all external dependencies (need mocks)
3. Identify all revert conditions (modifiers, requires)
4. Identify all events emitted
5. Identify all state changes
6. Check if similar tests exist

## Running Tests

```bash
# Run all unit tests
forge test --no-match-contract Integration

# Run specific test file
forge test --match-path src/test/unit/ContractUnit.t.sol

# Run specific test
forge test --match-test test_functionName_Success

# Run with verbosity
forge test --match-path src/test/unit/ContractUnit.t.sol -vvv

# Check coverage
forge coverage --match-path src/test/unit/ContractUnit.t.sol
```
