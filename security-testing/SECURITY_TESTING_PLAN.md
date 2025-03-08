# EigenLayer Security Testing Plan

## Overview
This document outlines a comprehensive security testing plan for the EigenLayer protocol, focusing on identifying and testing potential vulnerabilities in each system component.

## Core Components and Vulnerabilities

### 1. SlashingLib and Share Accounting
#### Potential Vulnerabilities:
- Mathematical precision loss in nested operations
- Rounding errors in share calculations
- Overflow/underflow in extreme scenarios
- Edge cases in deposit scaling factor calculations
- Division by zero when operators are fully slashed

#### Test Strategies:
- Fuzz testing with extreme values
- Property-based testing for mathematical invariants
- Boundary testing for rounding behavior
- Regression testing for known edge cases
- Tests for division by zero prevention

### 2. DelegationManager
#### Potential Vulnerabilities:
- Reentrancy in delegation and withdrawal flows
- Incorrect share accounting during delegation changes
- Vulnerabilities in withdrawal queue processing
- Improper handling of slashed operators
- Delegation to fully slashed operators

#### Test Strategies:
- Reentrancy attack simulations
- Integration tests for delegation state transitions
- Edge case testing for fully slashed scenarios
- Withdrawal process testing under various conditions
- Tests for delegation to slashed operators

### 3. StrategyManager
#### Potential Vulnerabilities:
- Reentrancy in deposit and withdrawal flows
- Incorrect share calculations for different token types
- Vulnerabilities in strategy upgrades
- Improper handling of rebasing tokens
- Deposit validation failures

#### Test Strategies:
- Deposit/withdrawal flow security testing
- Strategy upgrade security testing
- Token handling edge case testing
- Integration testing with DelegationManager
- Tests for deposit validation

### 4. EigenPodManager and EigenPod
#### Potential Vulnerabilities:
- Beacon chain slashing factor calculation errors
- Vulnerabilities in validator credential verification
- Improper handling of validator withdrawals
- Dual slashing edge cases
- Fully slashed EigenPod handling

#### Test Strategies:
- Beacon chain slashing simulation tests
- Validator credential verification testing
- Withdrawal process security testing
- Dual slashing scenario testing
- Tests for fully slashed EigenPod behavior

### 5. AllocationManager
#### Potential Vulnerabilities:
- Incorrect allocation and deallocation calculations
- Vulnerabilities in operator set management
- Slashing calculation errors
- Improper handling of registration/deregistration
- Rounding errors in slashing calculations

#### Test Strategies:
- Allocation/deallocation flow testing
- Operator set management security testing
- Slashing mechanism testing
- Registration/deregistration security testing
- Tests for rounding behavior in slashing

### 6. PermissionController
#### Potential Vulnerabilities:
- Privilege escalation
- Unauthorized access to restricted functions
- Improper permission delegation
- Vulnerabilities in admin transitions
- Permission boundary violations

#### Test Strategies:
- Permission boundary testing
- Admin role transition testing
- Unauthorized access attempt testing
- Permission delegation security testing
- Tests for permission revocation

## Test Implementation Plan

### Unit Tests
- Implement focused tests for each vulnerability type
- Test individual components in isolation
- Use mocks for external dependencies
- Focus on mathematical precision and edge cases
- Verify component-specific invariants

### Integration Tests
- Test interactions between components
- Simulate complex scenarios (e.g., dual slashing)
- Test end-to-end flows (deposit → delegate → allocate → slash → withdraw)
- Verify system-wide invariants
- Test cross-component dependencies

### Fuzz Tests
- Implement property-based tests with random inputs
- Focus on mathematical operations and share calculations
- Test boundary conditions and extreme values
- Verify invariants hold under all conditions
- Identify edge cases through randomized testing

## Known Edge Cases from SharesAccountingEdgeCases.md

### Division by Zero Scenarios
- When an operator is 100% slashed for a strategy (m_n = 0)
- When a staker's EigenPod has 0 ETH balance and all validators slashed (l_n = 0)
- Tests should verify proper handling of these cases

### Deposit Scaling Factor Bounds
- Upper bound on k_n is approximately 1e74 in worst case
- Tests should verify no overflow occurs in extreme scenarios

### Rounding Behavior
- Rounding up on slashing (to prevent DOS attacks)
- Deposits can actually reduce withdrawable shares due to rounding errors
- Operator shares are rounded up when slashing, staker withdrawable shares rounded down
- Tests should verify the invariant: op_n ≥ sum(a_n,i)

### Residual Operator Shares
- Maximum rounding error is approximately 1/1e18 of deposit amount
- Tests should verify error bounds are within acceptable limits

## Test Execution Strategy
- Use Foundry's forge for test execution
- Implement tight testing loops for efficient debugging
- Isolate failing tests for focused investigation
- Verify all tests pass before implementation is considered secure
- Prioritize testing of critical components and known edge cases
