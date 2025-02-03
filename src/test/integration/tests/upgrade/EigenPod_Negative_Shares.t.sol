// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

contract Integration_Upgrade_EigenPod_Negative_Shares is UpgradeTest {
    function _init() internal override {
        _configAssetTypes(HOLDS_ETH);
        _configUserTypes(DEFAULT);
    }

    function testFuzz_deposit_delegate_updateBalance_upgrade_completeAsShares(
        uint24 _rand
    ) public rand(_rand) {
        /// 0. Create an operator and staker with some underlying assets
        (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        uint256[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        ///  1. Deposit into strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);

        /// 2. Delegate to operator
        staker.delegateTo(operator);

        /// 3. Queue a withdrawal for all shares
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
        IDelegationManagerTypes.Withdrawal memory withdrawal = withdrawals[0];

        /// 4. Update balance randomly (can be positive or negative)
        (int256[] memory tokenDeltas, int256[] memory balanceUpdateShareDelta,) = _randBalanceUpdate(staker, strategies);
        staker.updateBalances(strategies, tokenDeltas);

        /// 5. Upgrade contracts
        _upgradeEigenLayerContracts();

        /// 6. Complete the withdrawal as shares
        _rollBlocksForCompleteWithdrawals(withdrawals);
        staker.completeWithdrawalAsShares(withdrawal);

        // Manually complete checks since we could still negative shares prior to the upgrade, causing a revert in the share check
        (uint256[] memory expectedOperatorShareDelta, int256[] memory expectedStakerShareDelta) =
            _getPostWithdrawalExpectedShareDeltas(balanceUpdateShareDelta[0], withdrawal);
        assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawal), "staker withdrawal should no longer be pending");
        assert_Snap_Unchanged_TokenBalances(staker, "staker should not have any change in underlying token balances");
        assert_Snap_Added_OperatorShares(operator, withdrawal.strategies, expectedOperatorShareDelta, "operator should have received shares");
        assert_Snap_Delta_StakerShares(staker, strategies, expectedStakerShareDelta, "staker should have received expected shares");
    }

    function testFuzz_deposit_delegate_updateBalance_upgrade_completeAsTokens(
        uint24 _rand
    ) public rand(_rand) {
        /// 0. Create an operator and staker with some underlying assets
        (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        uint256[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        ///  1. Deposit into strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);

        /// 2. Delegate to operator
        staker.delegateTo(operator);

        /// 3. Queue a withdrawal for all shares
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
        IDelegationManagerTypes.Withdrawal memory withdrawal = withdrawals[0];

        /// 4. Update balance randomly (can be positive or negative)
        (int256[] memory tokenDeltas, int256[] memory balanceUpdateShareDelta,) = _randBalanceUpdate(staker, strategies);
        staker.updateBalances(strategies, tokenDeltas);

        /// 5. Upgrade contracts
        _upgradeEigenLayerContracts();

        /// 6. Complete the withdrawal as shares
        _rollBlocksForCompleteWithdrawals(withdrawals);
        IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawal);
        uint256[] memory expectedTokens = _getPostWithdrawalExpectedTokenDeltas(balanceUpdateShareDelta[0], withdrawal);

        // Manually complete checks since we could still negative shares prior to the upgrade, causing a revert in the share check
        assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawal), "staker withdrawal should no longer be pending");
        assert_Snap_Added_TokenBalances(staker, tokens, expectedTokens, "staker should have received expected tokens");
        assert_Snap_Unchanged_OperatorShares(operator, "operator shares should not have changed");

        // If we had a positive balance update, then the staker shares should not have changed
        if (balanceUpdateShareDelta[0] > 0) {
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker shares should not have changed");
        }
        // Else, the staker shares should have increased by the magnitude of the negative share delta
        else {
            int256[] memory expectedStakerShareDelta = new int256[](1);
            expectedStakerShareDelta[0] = -balanceUpdateShareDelta[0];
            assert_Snap_Delta_StakerShares(staker, strategies, expectedStakerShareDelta, "staker should have received expected shares");
        }
    }



    function _getPostWithdrawalExpectedShareDeltas(
        int256 balanceUpdateShareDelta,
        IDelegationManagerTypes.Withdrawal memory withdrawal
    ) internal pure returns (uint256[] memory, int256[] memory) {
        uint256[] memory operatorShareDelta = new uint256[](1);
        int256[] memory stakerShareDelta = new int256[](1);
        // The staker share delta is the withdrawal scaled shares since it can go from negative to positive
        stakerShareDelta[0] = int256(withdrawal.scaledShares[0]);

        if (balanceUpdateShareDelta > 0) {
            // If balanceUpdateShareDelta is positive, then the operator delta is the withdrawal scaled shares
            operatorShareDelta[0] = withdrawal.scaledShares[0];
        } else {
            // Operator shares never went negative, so we can just add the withdrawal scaled shares and the negative share delta
            operatorShareDelta[0] = uint256(int256(withdrawal.scaledShares[0]) + balanceUpdateShareDelta);
        }

        return (operatorShareDelta, stakerShareDelta);
    }

    function _getPostWithdrawalExpectedTokenDeltas(
        int256 balanceUpdateShareDelta,
        IDelegationManagerTypes.Withdrawal memory withdrawal
    ) internal pure returns (uint256[] memory) {
        uint256[] memory expectedTokenDeltas = new uint256[](1);
        if (balanceUpdateShareDelta > 0) {
            // If we had a positive balance update, then the expected token delta is the withdrawal scaled shares
            expectedTokenDeltas[0] = withdrawal.scaledShares[0];
        } else {
            // If we had a negative balance update, then the expected token delta is the withdrawal scaled shares plus the negative share delta
            expectedTokenDeltas[0] = uint256(int256(withdrawal.scaledShares[0]) + balanceUpdateShareDelta);
        }
        return expectedTokenDeltas;
    }
}
