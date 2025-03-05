// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

contract Integration_Deposit_Delegate_UpdateBalance is IntegrationCheckUtils {
// TODO: fix for slashing
/// Generates a random stake and operator. The staker:
/// 1. deposits all assets into strategies
/// 2. delegates to an operator
/// 3. queues a withdrawal for a ALL shares
/// 4. updates their balance randomly
/// 5. completes the queued withdrawal as tokens
// function testFuzz_deposit_delegate_updateBalance_completeAsTokens(uint24 _random) public {
//     _configRand({
//         _randomSeed: _random,
//         _assetTypes: HOLDS_ETH, // HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
//         _userTypes: DEFAULT //| ALT_METHODS
//     });

//     /// 0. Create an operator and staker with some underlying assets
//     (
//         User staker,
//         IStrategy[] memory strategies,
//         uint[] memory tokenBalances
//     ) = _newRandomStaker();
//     (User operator, ,) = _newRandomOperator();
//     // Upgrade contracts if forkType is not local
//     _upgradeEigenLayerContracts();

//     uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

//     assert_HasNoDelegatableShares(staker, "staker should not have delegatable shares before depositing");
//     assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");

//     /// 1. Deposit into strategies
//     staker.depositIntoEigenlayer(strategies, tokenBalances);
//     check_Deposit_State(staker, strategies, shares);

//     /// 2. Delegate to an operator
//     staker.delegateTo(operator);
//     check_Delegation_State(staker, operator, strategies, shares);

//     /// 3. Queue withdrawals for ALL shares
//     Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);
//     bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
//     check_QueuedWithdrawal_State(staker, operator, strategies, shares, withdrawals, withdrawalRoots);

//     // Generate a random balance update:
//     // - For LSTs, the tokenDelta is positive tokens minted to the staker
//     // - For ETH, the tokenDelta is a positive or negative change in beacon chain balance
//     (
//         int[] memory tokenDeltas,
//         int[] memory stakerShareDeltas,
//         int[] memory operatorShareDeltas
//     ) = _randBalanceUpdate(staker, strategies);

//     // 4. Update LST balance by depositing, and beacon balance by submitting a proof
//     staker.updateBalances(strategies, tokenDeltas);
//     assert_Snap_Delta_StakerShares(staker, strategies, stakerShareDeltas, "staker should have applied deltas correctly");
//     assert_Snap_Delta_OperatorShares(operator, strategies, operatorShareDeltas, "operator should have applied deltas correctly");

//     console.log("withdrawble: ", staker.pod().withdrawableRestakedExecutionLayerGwei());

//     // Fast forward to when we can complete the withdrawal
//     _rollBlocksForCompleteWithdrawals(withdrawals);

//     // 5. Complete queued withdrawals as tokens
//     staker.completeWithdrawalsAsTokens(withdrawals);
//     assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should still be delegated to operator");
//     assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
//     assert_Snap_Unchanged_TokenBalances(operator, "operator token balances should not have changed");
//     assert_Snap_Unchanged_OperatorShares(operator, "operator shares should not have changed");
// }
}
