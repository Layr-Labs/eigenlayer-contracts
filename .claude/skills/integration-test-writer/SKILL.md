---
name: integration-test-writer
description: Write Solidity integration tests for EigenLayer contracts. Use when the user asks to write integration tests, test user flows, test cross-contract interactions, or test upgrade scenarios. Follows project conventions with User/AVS actors and numbered action steps.
allowed-tools: Read, Glob, Grep, Edit, Write, Bash(forge:*)
---

# Integration Test Writer

Write comprehensive integration tests for EigenLayer Solidity contracts following the project's established conventions.

## Overview

Integration tests orchestrate the deployment of all EigenLayer core contracts to test high-level user flows across multiple contracts. There are three test modes:

1. **Local Integration Tests** - Deploy fresh contracts and test user flows
2. **Fork Tests** - Fork mainnet, upgrade all contracts to latest implementations, then run the integration test suite
3. **Upgrade Tests** - Fork mainnet, perform actions on OLD contracts, then upgrade and verify compatibility

## Test Function Signature

**All integration test functions MUST:**
1. Be named `testFuzz_action1_action2_...` describing the flow
2. Take `uint24 _random` (or `_r`) as the only parameter - this seeds randomness
3. Use the `rand(_random)` modifier to initialize the random seed

```solidity
function testFuzz_deposit_delegate_queue_complete(uint24 _random) public rand(_random) {
    // Test implementation
}
```

The `rand()` modifier initializes the test's random seed, which is used by helper functions like `_newRandomStaker()` to generate deterministic random values for reproducible tests.

## Test File Locations

| Type | Location |
|------|----------|
| Normal integration tests | `src/test/integration/tests/` |
| Upgrade tests | `src/test/integration/tests/upgrade/` |
| Check functions | `src/test/integration/IntegrationChecks.t.sol` |
| Multichain checks | `src/test/integration/MultichainIntegrationChecks.t.sol` |

## Core Principles

### 1. All Actions Must Be Called Through User Contracts

Never call contracts directly. Use the `User` or `AVS` actor contracts:

```solidity
// ✅ CORRECT - Actions through User/AVS
staker.depositIntoEigenlayer(strategies, tokenBalances);
operator.delegateTo(operator);
avs.createOperatorSet(strategies);

// ❌ WRONG - Direct contract calls
strategyManager.depositIntoStrategy(...);
delegationManager.delegateTo(...);
```

### 2. Every Action Must Be Followed By a Check

After each numbered action, verify state changes using `check_*` functions from `IntegrationChecks.t.sol`:

```solidity
// 1. Deposit Into Strategies
staker.depositIntoEigenlayer(strategies, tokenBalances);
check_Deposit_State(staker, strategies, shares);

// 2. Delegate to an operator
staker.delegateTo(operator);
check_Delegation_State(staker, operator, strategies, shares);
```

### 3. Actions Must Be Numbered

Use comments to number each action step for clarity:

```solidity
// 1. Deposit Into Strategies
staker.depositIntoEigenlayer(strategies, tokenBalances);
check_Deposit_State(staker, strategies, shares);

// 2. Delegate to an operator
staker.delegateTo(operator);
check_Delegation_State(staker, operator, strategies, shares);

// 3. Queue Withdrawals
Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
check_QueuedWithdrawal_State(staker, operator, strategies, shares, withdrawableShares, withdrawals, withdrawalRoots);
```

## User Types

### `User` (Staker/Operator)

Located in `src/test/integration/users/User.t.sol`. A `User` can act as both a staker AND an operator.

**Key methods:**
- `depositIntoEigenlayer(strategies, tokenBalances)` - Deposit tokens
- `delegateTo(operator)` - Delegate to an operator
- `registerAsOperator()` - Register as an operator
- `queueWithdrawals(strategies, shares)` - Queue withdrawals
- `completeWithdrawalAsTokens(withdrawal)` / `completeWithdrawalAsShares(withdrawal)` - Complete withdrawals
- `registerForOperatorSet(operatorSet)` - Register for an operator set
- `modifyAllocations(params)` - Allocate magnitude to operator sets
- `startValidators()` / `verifyWithdrawalCredentials(validators)` - Native ETH staking
- `startCheckpoint()` / `completeCheckpoint()` - Checkpoint EigenPod

### `AVS`

Located in `src/test/integration/users/AVS.t.sol`. Represents an AVS that creates operator sets and slashes operators.

**Key methods:**
- `createOperatorSet(strategies)` - Create an operator set
- `createRedistributingOperatorSets(strategies, recipients)` - Create redistributing operator sets
- `slashOperator(params)` - Slash an operator
- `updateSlasher(operatorSetId, slasher)` - Update the slasher for an operator set

## Test Structure

### Normal Integration Test

```solidity
// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

/// @notice Base contract for shared setup across test variants
contract Integration_FlowName_Base is IntegrationCheckUtils {
    using ArrayLib for *;

    // Declare state used across tests
    AVS avs;
    OperatorSet operatorSet;
    User operator;
    User staker;
    IStrategy[] strategies;
    uint[] initTokenBalances;
    uint[] initDepositShares;

    /// @dev Setup state used across all test functions
    function _init() internal virtual override {
        _configAssetTypes(HOLDS_LST | HOLDS_ETH); // Configure asset types

        // Create actors
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        (operator,,) = _newRandomOperator();
        (avs,) = _newRandomAVS();

        // 1. Deposit into strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, initDepositShares);

        // 2. Delegate staker to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // 3. Create operator set and register
        operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, allStrats);
    }
}

/// @notice Test contract for specific flow variant
contract Integration_FlowName_Variant is Integration_FlowName_Base {
    /// @dev All test functions must:
    /// 1. Be named testFuzz_action1_action2_...
    /// 2. Take uint24 _r as parameter (seeds randomness)
    /// 3. Use rand(_r) modifier
    function testFuzz_action1_action2(uint24 _r) public rand(_r) {
        // 4. Next action
        // ... action ...
        // ... check ...

        // 5. Another action
        // ... action ...
        // ... check ...
    }
}
```

### Upgrade Test

Upgrade tests verify that upgrades correctly handle pre-upgrade state.

```solidity
// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

contract Integration_Upgrade_FeatureName is UpgradeTest {
    User staker;
    IStrategy[] strategies;
    uint[] tokenBalances;

    function _init() internal override {
        // Pre-upgrade setup - NO check_ functions here!
        (staker, strategies, tokenBalances) = _newRandomStaker();
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        // ... more pre-upgrade actions WITHOUT checks
    }

    function testFuzz_upgrade_scenario(uint24 _r) public rand(_r) {
        /// Pre-upgrade actions (no checks - old contracts)
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);

        /// Upgrade to new contracts
        _upgradeEigenLayerContracts();

        /// Post-upgrade actions WITH checks
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, operator, withdrawals[i], strategies, shares);
        }
    }
}
```

**Important for Upgrade Tests:**
- Pre-upgrade actions should NOT have `check_*` functions (old contracts have different invariants)
- Call `_upgradeEigenLayerContracts()` to upgrade to new contracts
- Post-upgrade actions SHOULD have `check_*` functions
- Only run on mainnet forks: `env FOUNDRY_PROFILE=forktest forge t --mc Integration_Upgrade`

## Check Functions

All state verification should be in `IntegrationChecks.t.sol`. There are two types:

### 1. `check_*` Functions

High-level state checks that verify multiple invariants after an action:

```solidity
check_Deposit_State(staker, strategies, shares);
check_Delegation_State(staker, operator, strategies, shares);
check_QueuedWithdrawal_State(staker, operator, strategies, shares, withdrawableShares, withdrawals, withdrawalRoots);
check_Withdrawal_AsTokens_State(staker, operator, withdrawal, strategies, shares, tokens, expectedTokens);
```

### 2. `assert_Snap_*` Functions

Time-machine assertions that compare state before/after an action:

```solidity
assert_Snap_Added_Staker_DepositShares(staker, strategy, amount, "error message");
assert_Snap_Removed_Staker_WithdrawableShares(staker, strategy, amount, "error message");
assert_Snap_Unchanged_Staker_DepositShares(staker, "error message");
```

### Adding New Checks

If a check doesn't exist, add it to `IntegrationChecks.t.sol`:

```solidity
function check_NewAction_State(
    User staker,
    IStrategy[] memory strategies,
    uint[] memory expectedValues
) internal {
    // Use assert_Snap_* for before/after comparisons
    assert_Snap_Added_Staker_DepositShares(
        staker, strategies[0], expectedValues[0], "should have added shares"
    );
    
    // Or use regular assertions for absolute checks
    assertEq(
        someContract.getValue(address(staker)),
        expectedValue,
        "value should match expected"
    );
}
```

## Randomness and Configuration

### The `rand(_r)` Modifier

Every test function takes a `uint24 _r` parameter and uses the `rand(_r)` modifier:

```solidity
function testFuzz_deposit_delegate(uint24 _r) public rand(_r) {
    // _r seeds all random generation in this test
    // This makes tests reproducible - same _r = same test execution
}
```

The `rand()` modifier initializes the random seed used by all `_newRandom*` helper functions. This ensures:
- **Reproducibility**: Same seed produces same random values
- **Fuzz coverage**: Foundry automatically runs with many different seeds

### Asset and User Type Configuration

Use `_configRand` or `_configAssetTypes` to control what types of users/assets are created:

```solidity
function testFuzz_example(uint24 _r) public rand(_r) {
    // Full configuration
    _configRand({
        _randomSeed: _r,
        _assetTypes: HOLDS_LST | HOLDS_ETH,
        _userTypes: DEFAULT | ALT_METHODS
    });
    
    // Or just configure asset types (simpler)
    _configAssetTypes(HOLDS_LST);
    
    // Create users - will use the configured randomization
    (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
}
```

**Asset Types:**
- `HOLDS_LST` - User holds liquid staking tokens
- `HOLDS_ETH` - User holds native ETH (beacon chain)
- `HOLDS_ALL` - User holds both

**User Types:**
- `DEFAULT` - Standard User contract
- `ALT_METHODS` - User that uses alternative method signatures

## Helper Functions

Common helpers available in `IntegrationBase`:

```solidity
// Create actors
(User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
(User operator, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomOperator();
User operator = _newRandomOperator_NoAssets();
(AVS avs, OperatorSet[] memory operatorSets) = _newRandomAVS();

// Calculate expected values
uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);
uint[] memory tokens = _calculateExpectedTokens(strategies, shares);
uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);

// Generate params
AllocateParams memory params = _genAllocation_AllAvailable(operator, operatorSet);
SlashingParams memory slashParams = _genSlashing_Rand(operator, operatorSet);

// Time advancement
_rollBlocksForCompleteWithdrawals(withdrawals);
_rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
_rollForward_AllocationConfigurationDelay();

// Get state
bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
IERC20[] memory tokens = _getUnderlyingTokens(strategies);
```

## Running Tests

```bash
# Run all integration tests locally (fresh contract deployment)
forge t --mc Integration

# Run mainnet fork tests (upgrades mainnet contracts to latest, then runs tests)
# Requires RPC_MAINNET environment variable
env FOUNDRY_PROFILE=forktest forge t --mc Integration

# Run upgrade tests only (tests upgrade compatibility)
env FOUNDRY_PROFILE=forktest forge t --mc Integration_Upgrade

# Run specific test
forge t --match-test testFuzz_deposit_delegate

# Run with verbosity
forge t --mc Integration -vvv
```

### Fork Tests vs Local Tests

| Mode | Command | What Happens |
|------|---------|--------------|
| Local | `forge t --mc Integration` | Deploys fresh contracts, runs tests |
| Fork | `env FOUNDRY_PROFILE=forktest forge t --mc Integration` | Forks mainnet, upgrades ALL contracts to latest repo implementations, runs tests |
| Upgrade | `env FOUNDRY_PROFILE=forktest forge t --mc Integration_Upgrade` | Forks mainnet, runs pre-upgrade actions on OLD contracts, then upgrades and tests compatibility |

**Fork tests** ensure that the latest contract code works correctly when upgrading from the current mainnet state. The test framework automatically upgrades all proxy contracts to the latest implementations before running tests.

## Naming Conventions

### Contract Names

| Pattern | Purpose |
|---------|---------|
| `Integration_FlowName_Base` | Base contract with shared `_init()` setup |
| `Integration_FlowName_Variant` | Test contract for specific flow variant |
| `Integration_Upgrade_FeatureName` | Upgrade test for a feature |

### Test Function Names

All test functions follow the pattern: `testFuzz_action1_action2_...(uint24 _random) public rand(_random)`

| Example | Description |
|---------|-------------|
| `testFuzz_deposit_delegate_queue_completeAsTokens` | Deposit → Delegate → Queue → Complete as tokens |
| `testFuzz_deposit_delegate_undelegate` | Deposit → Delegate → Undelegate |
| `testFuzz_verifyWC_checkpoint_slash` | Verify withdrawal credentials → Checkpoint → Slash |
| `testFuzz_upgrade_migrate_slash` | Upgrade contracts → Migrate → Slash |

### Check/Assert Names

| Pattern | Purpose |
|---------|---------|
| `check_ActionName_State` | High-level check function in IntegrationChecks |
| `assert_Snap_Added_*` | Assert value increased from snapshot |
| `assert_Snap_Removed_*` | Assert value decreased from snapshot |
| `assert_Snap_Unchanged_*` | Assert value unchanged from snapshot |

## Checklist Before Writing Tests

1. Identify the user flow to test
2. Determine if it's a normal test or upgrade test
3. Identify all actors needed (stakers, operators, AVSs)
4. Plan the numbered action steps
5. Identify which `check_*` functions to use after each action
6. If checks don't exist, add them to `IntegrationChecks.t.sol`
7. Use appropriate asset/user type configuration

## Example: Complete Integration Test

Reference: `src/test/integration/tests/DualSlashing.t.sol`

This file demonstrates:
- Base contract with `_init()` for shared setup
- Multiple test contracts inheriting from base
- Numbered action steps
- `check_*` after every action
- Using both `User` and `AVS` actors
- Slashing and checkpoint flows