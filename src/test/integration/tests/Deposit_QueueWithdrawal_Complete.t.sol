// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/IntegrationBase.t.sol";
import "src/test/integration/User.t.sol";
import "src/test/integration/tests/utils.t.sol";

contract Integration_Deposit_QueueWithdrawal_Complete is IntegrationTestUtils {

    /// Randomly generates a user with different held assets. Then:
    /// 1. deposit into strategy
    /// 2. queueWithdrawal
    /// 3. completeQueuedWithdrawal"
     function testFuzz_deposit_queueWithdrawal_complete(uint24 _random) public {
        // Configure the random parameters for the test
        _configRand({
            _randomSeed: _random,
            _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
            _userTypes: DEFAULT | ALT_METHODS
        });

        // Create a staker with a nonzero balance and corresponding strategies
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();

        // 1. Deposit into strategy
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);
        assertDepositState(staker, strategies, shares);

        // Ensure staker is not delegated to anyone post deposit
        assertFalse(delegationManager.isDelegated(address(staker)), "Staker should not be delegated after deposit");

        // 2. Queue Withdrawal
        IDelegationManager.Withdrawal[] memory withdrawals;
        bytes32[] memory withdrawalRoots;
        (withdrawals, withdrawalRoots) = staker.queueWithdrawals(strategies, shares);

        // 3. Complete Queued Withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());
        for (uint i = 0; i < withdrawals.length; i++) {
            IERC20[] memory tokens = staker.completeQueuedWithdrawal(withdrawals[i], true); // Assume 'true' is to receive as tokens
            uint[] memory expectedTokens = _calculateExpectedTokens(strategies, shares);
            assertWithdrawalAsTokensState(staker, withdrawals[i], strategies, shares, tokens, expectedTokens);
        }

        // Ensure staker is still not delegated to anyone post withdrawal completion
        assertFalse(delegationManager.isDelegated(address(staker)), "Staker should still not be delegated after withdrawal");
    }
}
