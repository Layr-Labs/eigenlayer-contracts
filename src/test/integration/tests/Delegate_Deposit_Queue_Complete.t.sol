// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/IntegrationBase.t.sol";
import "src/test/integration/User.t.sol";
import "src/test/integration/tests/utils.t.sol";

contract Integration_Delegate_Deposit_QueueWithdrawal_Complete is IntegrationTestUtils {
    
    function testFuzz_delegate_deposit_queueWithdrawal_completeAsShares(uint24 _random) public {
        // Configure the random parameters for the test
        _configRand({
            _randomSeed: _random,
            _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
            _userTypes: DEFAULT | ALT_METHODS
        });

        // Create a staker and an operator with a nonzero balance and corresponding strategies
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator, ,) = _newRandomOperator();

        // 1. Delegate to operator
        staker.delegateTo(operator);
        assertDelegationState(staker, operator, strategies, new uint256[](strategies.length)); // Initial shares are zero

        // 2. Deposit into strategy
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        // Check that the deposit increased operator shares the staker is delegated to
        assertDelegationState(staker, operator, strategies, shares);

        // 3. Queue Withdrawal
        IDelegationManager.Withdrawal[] memory withdrawals;
        bytes32[] memory withdrawalRoots;
        (withdrawals, withdrawalRoots) = staker.queueWithdrawals(strategies, shares);
        assertQueuedWithdrawalState(staker, operator, strategies, shares, withdrawals, withdrawalRoots);

        // 4. Complete Queued Withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeQueuedWithdrawal(withdrawals[i], false); // 'false' indicates completion as shares
            assertWithdrawalAsSharesState(staker, withdrawals[i], strategies, shares);
        }
    }

    function testFuzz_delegate_deposit_queueWithdrawal_completeAsTokens(uint24 _random) public {
        // Configure the random parameters for the test
        _configRand({
            _randomSeed: _random,
            _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
            _userTypes: DEFAULT | ALT_METHODS
        });

        // Create a staker and an operator with a nonzero balance and corresponding strategies
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator, ,) = _newRandomOperator();

        // 1. Delegate to operator
        staker.delegateTo(operator);
        assertDelegationState(staker, operator, strategies, new uint256[](strategies.length)); // Initial shares are zero

        // 2. Deposit into strategy
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        // Check that the deposit increased operator shares the staker is delegated to
        assertDelegationState(staker, operator, strategies, shares);

        // 3. Queue Withdrawal
        IDelegationManager.Withdrawal[] memory withdrawals;
        bytes32[] memory withdrawalRoots;
        (withdrawals, withdrawalRoots) = staker.queueWithdrawals(strategies, shares);
        assertQueuedWithdrawalState(staker, operator, strategies, shares, withdrawals, withdrawalRoots);

        // 4. Complete Queued Withdrawal
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());
        for (uint i = 0; i < withdrawals.length; i++) {
            IERC20[] memory tokens = staker.completeQueuedWithdrawal(withdrawals[i], true); // Assuming 'true' for tokens
            uint[] memory expectedTokens = _calculateExpectedTokens(strategies, shares);
            assertWithdrawalAsTokensState(staker, withdrawals[i], strategies, shares, tokens, expectedTokens);
        }
    }
}
