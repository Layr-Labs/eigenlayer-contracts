# EigenLayer Security Testing

This directory contains security testing documentation and test specifications for the EigenLayer protocol.

## Contents

- `SECURITY_TESTING_PLAN.md`: Comprehensive security testing plan
- `SlashingLib_Tests.md`: Test specifications for SlashingLib
- `DelegationManager_Tests.md`: Test specifications for DelegationManager
- `StrategyManager_Tests.md`: Test specifications for StrategyManager
- `AllocationManager_Tests.md`: Test specifications for AllocationManager
- `DualSlashing_Tests.md`: Test specifications for dual slashing scenarios
- `PermissionController_Tests.md`: Test specifications for PermissionController
- `run_security_tests.sh`: Script to execute all security tests

## Implementation

The actual test implementations can be found in the `src/test/security/` directory.

## Running Tests

To run all security tests:

```bash
./run_security_tests.sh
```

To run a specific test:

```bash
forge test --match-path src/test/security/SlashingLibSecurity.t.sol -vv
```

## Test Categories

The security tests are organized into the following categories:

1. **Mathematical Precision Tests**: Tests for precision loss, rounding errors, and extreme values
2. **Reentrancy Tests**: Tests for reentrancy vulnerabilities in deposit, withdrawal, and delegation flows
3. **Share Accounting Tests**: Tests for correct share calculations and scaling
4. **Slashing Tests**: Tests for slashing mechanisms and dual slashing scenarios
5. **Permission Tests**: Tests for proper permission enforcement and privilege escalation prevention

## Implementation Strategy

When implementing security tests:

1. Focus on one vulnerability type at a time
2. Use fuzzing for mathematical operations and boundary testing
3. Implement property-based tests to verify invariants
4. Test interactions between components for integration vulnerabilities
5. Verify all tests pass before considering implementation secure

## Reporting Issues

If you discover a security vulnerability during testing:

1. Document the vulnerability with steps to reproduce
2. Assess the severity and potential impact
3. Propose a fix or mitigation strategy
4. Report the issue to the security team
