// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/test/integration/IntegrationBase.t.sol";
import "src/test/integration/User.t.sol";
import "src/test/integration/tests/utils.t.sol";

contract Integration_Deposit_Register_QueueWithdrawal_Complete is IntegrationTestUtils {
    function testFuzz_deposit_registerOperator_queueWithdrawal_completeAsShares(uint24 _random) public {
        // Configure the random parameters for the test
        _configRand({
            _randomSeed: _random,
            _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
            _userTypes: DEFAULT | ALT_METHODS
        });

        // Create a staker with a nonzero balance and corresponding strategies
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();

        // 1. Staker deposits into strategy
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);
        assertDepositState(staker, strategies, shares);

        // 2. Staker registers as an operator
        staker.registerAsOperator();
        assertRegisteredAsOperatorState(staker);

        // 3. Queue Withdrawal
        IDelegationManager.Withdrawal[] memory withdrawals;
        bytes32[] memory withdrawalRoots;
        (withdrawals, withdrawalRoots) = staker.queueWithdrawals(strategies, shares);

        // 4. Complete Queued Withdrawal as Shares
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeQueuedWithdrawal(withdrawals[i], false); // 'false' indicates completion as shares
            assertWithdrawalAsSharesState(staker, withdrawals[i], strategies, shares);
        }
    }

    function testFuzz_deposit_registerOperator_queueWithdrawal_completeAsTokens(uint24 _random) public {
        // Configure the random parameters for the test
        _configRand({
            _randomSeed: _random,
            _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
            _userTypes: DEFAULT | ALT_METHODS
        });

        // Create a staker with a nonzero balance and corresponding strategies
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();

        // 1. Staker deposits into strategy
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);
        assertDepositState(staker, strategies, shares);

        // 2. Staker registers as an operator
        staker.registerAsOperator();
        assertRegisteredAsOperatorState(staker);

        // 3. Queue Withdrawal
        IDelegationManager.Withdrawal[] memory withdrawals;
        bytes32[] memory withdrawalRoots;
        (withdrawals, withdrawalRoots) = staker.queueWithdrawals(strategies, shares);

        // 4. Complete Queued Withdrawal as Tokens
        cheats.roll(block.number + delegationManager.withdrawalDelayBlocks());
        for (uint i = 0; i < withdrawals.length; i++) {
            IERC20[] memory tokens = staker.completeQueuedWithdrawal(withdrawals[i], true); // 'true' indicates completion as tokens
            uint[] memory expectedTokens = _calculateExpectedTokens(strategies, shares);
            assertWithdrawalAsTokensState(staker, withdrawals[i], strategies, shares, tokens, expectedTokens);
        }
    }
}
