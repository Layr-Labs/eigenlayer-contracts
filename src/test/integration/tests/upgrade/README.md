## Upgrade Tests

This folder contains specific tests we want to conduct to determine upgrade compatibility. Tests in this folder are only run if the `forktest` profile is selected, e.g:

`env FOUNDRY_PROFILE=forktest forge t --mc Upgrade`

#### How to Write

Upgrade tests inherit from the `UpgradeTest` mixin, which imports everything you need for a standard integration test. A standard integration test automatically upgrades all contracts to the latest repo contracts, even when forking from mainnet.

In contrast, the `UpgradeTest` mixin ensures that (when a test begins) the contracts we're working with are _not upgraded_. This allows the test writer to perform some initial setup actions before calling `_upgradeEigenLayerContracts()`, after which we can check that the upgrade correctly handles pre-upgrade state.

#### Example

Say we want to check that withdrawals initiated before the slashing upgrade are completable after the slashing upgrade. This example test shows how:

```solidity
contract Integration_Upgrade_Complete_PreSlashing_Withdrawal is UpgradeTest {
    
    function testFuzz_deposit_queue_upgrade_completeAsShares(uint24 _random) public rand(_random) {
        /// Pre-upgrade:
        /// 1. Create staker with some assets
        /// 2. Staker deposits into EigenLayer
        /// 3. Staker queues a withdrawal
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        User operator = User(payable(0));

        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);

        /// Upgrade to slashing contracts
        _upgradeEigenLayerContracts();

        // Complete pre-slashing withdrawals as shares
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, operator, withdrawals[i], strategies, shares);
        }
    }
}
```

**Note** how the initial staker actions are NOT followed by `check_X_State` methods. This is because, before calling `_upgradeEigenLayerContracts`, the test is being run on old contracts. Adding invariant checks to old state transitions is not the point of this test and would add a lot of maintenance overhead.