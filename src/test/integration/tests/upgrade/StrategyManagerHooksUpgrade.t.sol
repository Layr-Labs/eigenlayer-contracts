// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

/// @title Upgrade test for StrategyManager hooks (beforeAddShares/beforeRemoveShares)
/// @notice Verifies that deposits made before the upgrade can be queued for withdrawal
///         and completed as shares after the upgrade, testing that the new hooks work
///         correctly with pre-existing strategies.
contract Integration_Upgrade_StrategyManagerHooks is UpgradeTest {
    /// @notice Test: deposit -> upgrade -> queue withdrawal -> complete as shares
    /// @dev Ensures the new beforeRemoveShares and beforeAddShares hooks don't break
    ///      the withdrawal flow for deposits made before the upgrade.
    function testFuzz_deposit_upgrade_queue_completeAsShares(uint24 _random) public rand(_random) {
        /// Pre-upgrade:
        /// 1. Create staker with some assets
        /// 2. Staker deposits into EigenLayer
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();

        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        /// Upgrade to new contracts (with beforeAddShares/beforeRemoveShares hooks)
        _upgradeEigenLayerContracts();

        /// Post-upgrade:
        /// Queue withdrawal (tests beforeRemoveShares hook)
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);

        /// Complete withdrawal as shares (tests beforeAddShares hook)
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
        }

        /// Verify staker has their shares back
        for (uint i = 0; i < strategies.length; i++) {
            uint stakerShares;
            if (strategies[i] == BEACONCHAIN_ETH_STRAT) {
                // Beacon chain ETH shares are tracked in EigenPodManager, not StrategyManager
                int podShares = eigenPodManager.podOwnerDepositShares(address(staker));
                stakerShares = podShares > 0 ? uint(podShares) : 0;
            } else {
                stakerShares = strategyManager.stakerDepositShares(address(staker), strategies[i]);
            }
            assertEq(stakerShares, shares[i], "staker should have shares restored after completing as shares");
        }
    }

    /// @notice Test: deposit -> upgrade -> queue withdrawal -> complete as tokens
    /// @dev Ensures the new beforeRemoveShares hook doesn't break the token withdrawal flow.
    function testFuzz_deposit_upgrade_queue_completeAsTokens(uint24 _random) public rand(_random) {
        /// Pre-upgrade:
        /// 1. Create staker with some assets
        /// 2. Staker deposits into EigenLayer
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();

        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        /// Upgrade to new contracts
        _upgradeEigenLayerContracts();

        /// Post-upgrade:
        /// Queue withdrawal (tests beforeRemoveShares hook)
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);

        /// Complete withdrawal as tokens
        _rollBlocksForCompleteWithdrawals(withdrawals);

        // Record token balances before completing withdrawals
        IERC20[] memory tokens = _getUnderlyingTokens(strategies);
        uint[] memory balancesBefore = _getTokenBalances(staker, tokens);

        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsTokens(withdrawals[i]);
        }

        // Get balances after completing withdrawals
        uint[] memory balancesAfter = _getTokenBalances(staker, tokens);

        /// Verify staker has no shares and got tokens back
        for (uint i = 0; i < strategies.length; i++) {
            uint stakerShares = strategyManager.stakerDepositShares(address(staker), strategies[i]);
            assertEq(stakerShares, 0, "staker should have no shares after completing as tokens");

            // Verify token balance increased (skip beacon chain ETH which has different mechanics)
            if (strategies[i] != BEACONCHAIN_ETH_STRAT) {
                assertGt(balancesAfter[i], balancesBefore[i], "staker should have received tokens back");
            }
        }
    }

    /// @notice Test: queue -> upgrade -> complete as shares
    /// @dev Ensures withdrawals queued before the upgrade can still be completed
    ///      as shares after the upgrade.
    function testFuzz_queue_upgrade_completeAsShares(uint24 _random) public rand(_random) {
        /// Pre-upgrade:
        /// 1. Create staker with some assets
        /// 2. Staker deposits into EigenLayer
        /// 3. Queue withdrawal
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();

        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);

        /// Upgrade to new contracts
        _upgradeEigenLayerContracts();

        /// Post-upgrade:
        /// Complete withdrawal as shares (tests beforeAddShares hook for pre-upgrade queued withdrawal)
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
        }

        /// Verify staker has their shares back
        for (uint i = 0; i < strategies.length; i++) {
            uint stakerShares;
            if (strategies[i] == BEACONCHAIN_ETH_STRAT) {
                // Beacon chain ETH shares are tracked in EigenPodManager, not StrategyManager
                int podShares = eigenPodManager.podOwnerDepositShares(address(staker));
                stakerShares = podShares > 0 ? uint(podShares) : 0;
            } else {
                stakerShares = strategyManager.stakerDepositShares(address(staker), strategies[i]);
            }
            assertEq(stakerShares, shares[i], "staker should have shares restored after completing as shares");
        }
    }
}
