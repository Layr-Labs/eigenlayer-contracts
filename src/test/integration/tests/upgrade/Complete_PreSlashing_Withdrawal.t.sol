// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

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
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
    
        /// Upgrade to slashing contracts
        _upgradeEigenLayerContracts();
    
        // Complete pre-slashing withdrawals as shares
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, operator, withdrawals[i], strategies, shares);
        }
    }

    function testFuzz_delegate_deposit_queue_upgrade_completeAsShares(uint24 _random) public rand(_random) {
        /// Pre-upgrade:
        /// 1. Create staker and operator with some asset amounts
        /// 2. Staker delegates to operator
        /// 3. Staker deposits into EigenLayer
        /// 4. Staker queues a withdrawal
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator, ,) = _newRandomOperator();

        staker.delegateTo(operator);
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);

        /// Upgrade to slashing contracts
        _upgradeEigenLayerContracts();

        // Complete pre-slashing withdrawals as shares
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, operator, withdrawals[i], strategies, shares);
        }
    }

    function testFuzz_deposit_queue_upgrade_completeAsTokens(uint24 _random) public rand(_random) {
        /// Pre-upgrade:
        /// 1. Create staker with some assets
        /// 2. Staker deposits into EigenLayer
        /// 3. Staker queues a withdrawal
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        User operator = User(payable(0));

        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);
        uint[] memory expectedTokens = _calculateExpectedTokens(strategies, shares);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
    
        /// Upgrade to slashing contracts
        _upgradeEigenLayerContracts();
    
        // Complete pre-slashing withdrawals as tokens
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, shares, tokens, expectedTokens);
        }
    }

    function testFuzz_delegate_deposit_queue_upgrade_completeAsTokens(uint24 _random) public rand(_random) {
        /// Pre-upgrade:
        /// 1. Create staker and operator with some asset amounts
        /// 2. Staker delegates to operator
        /// 3. Staker deposits into EigenLayer
        /// 4. Staker queues a withdrawal
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator, ,) = _newRandomOperator();

        staker.delegateTo(operator);
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);
        uint[] memory expectedTokens = _calculateExpectedTokens(strategies, shares);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);

        /// Upgrade to slashing contracts
        _upgradeEigenLayerContracts();

        // Complete pre-slashing withdrawals as tokens
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, shares, tokens, expectedTokens);
        }
    }

    function testFuzz_upgrade_delegate_queuePartial_completeAsTokens(uint24 _random) public rand(_random) {
        /// 1. Create staker and operator with some asset amounts
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();
        (User operator, ,) = _newRandomOperator();

        /// 2. Upgrade to slashing contracts
        _upgradeEigenLayerContracts();

        /// 3. Staker delegates to operator and deposits
        staker.delegateTo(operator);
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        /// 4. Queue partial withdrawal
        uint[] memory partialShares = new uint[](shares.length);
        for (uint i = 0; i < shares.length; i++) {
            partialShares[i] = shares[i] / 2;
        }
        uint[] memory expectedTokens = _calculateExpectedTokens(strategies, partialShares);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, partialShares);

        /// 5. Complete withdrawals as tokens
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, partialShares, tokens, expectedTokens);
        }
    }

    /// TODO - complete pre-upgrade withdrawal after earliest possible operator slashing
}
