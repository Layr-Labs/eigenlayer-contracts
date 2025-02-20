// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";
import {console} from "forge-std/console.sol";

contract Integration_Deposit_Delegate_Allocate_Slash_Queue_Redeposit is IntegrationCheckUtils, IDelegationManagerTypes {

    function testFuzz_deposit_delegate_allocate_fullSlash_queue_complete_redeposit(
        uint24 _random
    ) public {
        _configRand({_randomSeed: _random, _assetTypes: HOLDS_LST, _userTypes: DEFAULT});
        _upgradeEigenLayerContracts();

        (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        (AVS avs,) = _newRandomAVS();

        uint256[] memory tokensToDeposit = new uint256[](tokenBalances.length);
        uint256[] memory numTokensRemaining = new uint256[](tokenBalances.length);
        for (uint256 i = 0; i < tokenBalances.length; i++) {
                tokensToDeposit[i] = tokenBalances[i]/2;
                numTokensRemaining[i] = tokenBalances[i] - tokensToDeposit[i];
        }

        uint256[] memory shares = _calculateExpectedShares(strategies, tokensToDeposit);

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokensToDeposit);
        check_Deposit_State_PartialDeposit(staker, strategies, shares, numTokensRemaining);

        // 2. Delegate to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);

        // Create operator set and register operator
        OperatorSet memory operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);

        // 3. Allocate to operator set
        IAllocationManagerTypes.AllocateParams memory allocateParams = 
            operator.modifyAllocations(operatorSet, _maxMagnitudes(operatorSet, operator));
        
        assert_Snap_Allocations_Modified(
            operator,
            allocateParams,
            false,
            "operator allocations should be updated before delay"
        );

        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
        
        assert_Snap_Allocations_Modified(
            operator,
            allocateParams,
            true,
            "operator allocations should be updated after delay"
        );

        // 4. Fully slash operator
        IAllocationManagerTypes.SlashingParams memory slashingParams;
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _strategiesAndWadsForFullSlash(operatorSet);

            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_StakerDepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker deposit shares should be slashed");
        }

        // 5. Undelegate from an operator
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // 6. Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint256[] memory expectedTokens =
                _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            for (uint256 i = 0; i < expectedTokens.length; i++) {
            }
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State_AfterSlash(staker, operator, withdrawals[i], allocateParams, slashingParams, expectedTokens);
        }

        // 7. Redeposit
        shares = _calculateExpectedShares(strategies, numTokensRemaining);
        staker.depositIntoEigenlayer(strategies, numTokensRemaining);
        check_Deposit_State(staker, strategies, shares);

        // Final state checks
        assert_HasExpectedShares(staker, strategies, shares, "staker should have expected shares after redeposit");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");      
    }
}