// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

contract Integration_Upgrade_Complete_PreSlashing_Withdrawal is UpgradeTest {
    struct TestState {
        User staker;
        User operator;
        IStrategy[] strategies;
        uint256[] tokenBalances;
        uint256[] shares;
        uint256[] withdrawalShares;
        uint256[] expectedTokens;
        Withdrawal[] withdrawals;
        bool isPartial;
        bool completeAsTokens;
    }

    function _init_(uint24 randomness, bool withOperator, bool withDelegation) internal returns (TestState memory state) {
        (state.staker, state.strategies, state.tokenBalances) = _newRandomStaker();
        state.shares = _calculateExpectedShares(state.strategies, state.tokenBalances);

        state.operator = withOperator ? _newRandomOperator_NoAssets() : User(payable(0));
        if (withDelegation) state.staker.delegateTo(state.operator);

        state.staker.depositIntoEigenlayer(state.strategies, state.tokenBalances);

        // Setup withdrawal shares (full or partial)
        state.isPartial = _randBool();
        state.withdrawalShares = new uint256[](state.shares.length);
        for (uint256 i = 0; i < state.shares.length; i++) {
            state.withdrawalShares[i] = state.isPartial ? state.shares[i] / 2 : state.shares[i];
        }

        state.expectedTokens = _calculateExpectedTokens(state.strategies, state.withdrawalShares);
        state.completeAsTokens = _randBool();
    }

    function _completeWithdrawal(TestState memory state) internal {
        for (uint256 i = 0; i < state.withdrawals.length; i++) {
            if (state.completeAsTokens) {
                IERC20[] memory tokens = state.staker.completeWithdrawalAsTokens(state.withdrawals[i]);
                check_Withdrawal_AsTokens_State(state.staker, state.operator, state.withdrawals[i], state.strategies, state.withdrawalShares, tokens, state.expectedTokens);
            } else {
                state.staker.completeWithdrawalAsShares(state.withdrawals[i]);
                check_Withdrawal_AsShares_State(state.staker, state.operator, state.withdrawals[i], state.strategies, state.withdrawalShares);
            }
        }
    }

    function testFuzz_deposit_queue_upgrade_complete(uint24 r) public rand(r) {
        TestState memory state = _init_({randomness: r, withOperator: false, withDelegation: false});
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _upgradeEigenLayerContracts();
        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        _completeWithdrawal(state);
    }

    function testFuzz_delegate_deposit_queue_upgrade_complete(uint24 r) public rand(r) {
        TestState memory state = _init_({randomness: r, withOperator: true, withDelegation: true});
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _upgradeEigenLayerContracts();
        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        _completeWithdrawal(state);
    }

    function testFuzz_upgrade_delegate_queuePartial_complete(uint24 r) public rand(r) {
        TestState memory state = _init_({randomness: r, withOperator: true, withDelegation: false});
        _upgradeEigenLayerContracts();
        state.staker.delegateTo(state.operator);
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        _completeWithdrawal(state);
    }

    function testFuzz_delegate_deposit_queue_completeBeforeUpgrade(uint24 r) public rand(r) {
        TestState memory state = _init_({randomness: r, withOperator: true, withDelegation: true});
        state.withdrawals = state.staker.queueWithdrawals(state.strategies, state.withdrawalShares);
        _rollBlocksForCompleteWithdrawals(state.withdrawals);
        // We must roll forward by the delay twice since pre-upgrade delay is half as long as post-upgrade delay.
        rollForward(delegationManager.minWithdrawalDelayBlocks() + 1);
        _upgradeEigenLayerContracts();
        _completeWithdrawal(state);
    }
}
