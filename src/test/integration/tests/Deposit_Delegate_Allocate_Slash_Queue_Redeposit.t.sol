// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";
import {console} from "forge-std/console.sol";

contract Integration_Deposit_Delegate_Allocate_Slash_Queue_Redeposit is IntegrationCheckUtils {

    AVS avs;
    OperatorSet operatorSet;

    User operator;
    AllocateParams allocateParams;

    User staker;
    IStrategy[] strategies;
    IERC20[] tokens;
    uint[] initTokenBalances;
    uint[] initDepositShares;

    uint[] numTokensRemaining;

    function _init() internal override {
        // TODO: Partial deposits don't work when beacon chain eth balance is initialized to < 64 ETH, need to write _newRandomStaker variant that ensures beacon chain ETH balance
        // greater than or equal to 64
        _configAssetTypes(HOLDS_LST);
        _configUserTypes(DEFAULT);

        (staker, strategies, initTokenBalances) = _newRandomStaker();
        (operator,,) = _newRandomOperator();
        (avs,) = _newRandomAVS();
        tokens = _getUnderlyingTokens(strategies);

        uint256[] memory tokensToDeposit = new uint256[](initTokenBalances.length);
        numTokensRemaining = new uint256[](initTokenBalances.length);
        for (uint256 i = 0; i < initTokenBalances.length; i++) {
            if (strategies[i] == BEACONCHAIN_ETH_STRAT) {
                tokensToDeposit[i] = initTokenBalances[i];
                continue;
            }

            tokensToDeposit[i] = initTokenBalances[i]/2;
            numTokensRemaining[i] = initTokenBalances[i] - tokensToDeposit[i];
        }

        uint256[] memory shares = _calculateExpectedShares(strategies, tokensToDeposit);

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokensToDeposit);
        check_Deposit_State_PartialDeposit(staker, strategies, shares, numTokensRemaining);

        // 2. Delegate to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);

        // Create operator set and register operator
        operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, allStrats);

        // 3. Allocate to operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);

        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
    }
  
    function testFuzz_fullSlash_undelegate_complete_redeposit(
        uint24 _random
    ) public rand(_random) {
        // 4. Fully slash operator
        SlashingParams memory slashParams = _genSlashing_Full(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_FullySlashed_State(operator, allocateParams, slashParams);

        // 5. Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // 6. Complete withdrawal. Staker should receive 0 shares/tokens after a full slash
        _rollBlocksForCompleteWithdrawals(withdrawals);
        uint[] memory expectedShares = new uint[](strategies.length);
        uint[] memory expectedTokens = new uint[](strategies.length);

        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, expectedShares, tokens, expectedTokens);
        }

        // 7. Redeposit
        uint[] memory shares = _calculateExpectedShares(strategies, numTokensRemaining);
        staker.depositIntoEigenlayer(strategies, numTokensRemaining);
        check_Deposit_State(staker, strategies, shares);

        // Final state checks
        assert_HasExpectedShares(staker, strategies, shares, "staker should have expected shares after redeposit");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");      
    }

    function testFuzz_undelegate_fullSlash_complete_redeposit(
        uint24 _random
    ) public rand(_random) {
        // 4. Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // 5. Fully slash operator
        SlashingParams memory slashParams = _genSlashing_Full(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_FullySlashed_State(operator, allocateParams, slashParams);

        // 6. Complete withdrawal. Staker should receive 0 shares/tokens after a full slash
        _rollBlocksForCompleteWithdrawals(withdrawals);
        uint[] memory expectedShares = new uint[](strategies.length);
        uint[] memory expectedTokens = new uint[](strategies.length);

        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, expectedShares, tokens, expectedTokens);
        }

        // 7. Redeposit
        uint[] memory shares = _calculateExpectedShares(strategies, numTokensRemaining);
        staker.depositIntoEigenlayer(strategies, numTokensRemaining);
        check_Deposit_State(staker, strategies, shares);

        // Final state checks
        assert_HasExpectedShares(staker, strategies, shares, "staker should have expected shares after redeposit");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");      
    }

    function testFuzz_depositFull_fullSlash_undelegate_completeAsShares(
        uint24 _random
    ) public rand(_random) {
        uint[] memory shares = _calculateExpectedShares(strategies, numTokensRemaining);
        staker.depositIntoEigenlayer(strategies, numTokensRemaining);
        check_Deposit_State(staker, strategies, shares);

        // 4. Fully slash random proper subset of operators strategies
        SlashingParams memory slashingParams;
        {
            (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
                _strategiesAndWadsForRandFullSlash(operatorSet);
            slashingParams = avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
            assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");
            assert_Snap_Unchanged_Staker_DepositShares(staker, "staker deposit shares should be unchanged after slashing");
            assert_Snap_StakerWithdrawableShares_AfterSlash(staker, allocateParams, slashingParams, "staker deposit shares should be slashed");
        }

        // 5. Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // 6. Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedShares = _calculateExpectedShares(withdrawals[i]);
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, expectedShares);
        }

        // Check final state:
        assert_HasNoUnderlyingTokenBalance(staker, slashingParams.strategies, "staker not have any underlying tokens");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /**
     * T-83: Staker with nonzero shares, Operator allocated, Staker Delegated
     * Partially slash operator on opSet, Redeposit, Queue Full Withdrawal, Complete as Tokens
     */
    function testFuzz_deposit_delegate_allocate_partialSlash_redeposit_queue_complete(uint24 r) public rand(r) {
        // Partially slash operator
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);

        // Redeposit remaining tokens
        uint[] memory additionalShares = _calculateExpectedShares(strategies, numTokensRemaining);
        staker.depositIntoEigenlayer(strategies, numTokensRemaining);
        check_Deposit_State(staker, strategies, additionalShares);
        
        // Queue withdrawal
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, withdrawableShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        assert_AllWithdrawalsPending(withdrawalRoots, "withdrawals should be pending after queueing");
        
        // Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            // Check that each withdrawal is properly completed
            assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawals[i]), 
                "withdrawal should not be pending after completion");
        }
        
        // Final state checks
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }
    
    /**
     * T-84: Staker with nonzero shares, Operator allocated, Staker Delegated
     * Undelegate, Partially slash operator on opSet, Complete as Tokens
     */
    function testFuzz_deposit_delegate_undelegate_partialSlash_complete(uint24 r) public rand(r) {
        // Undelegate from operator
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        assert_AllWithdrawalsPending(withdrawalRoots, "withdrawals should be pending after undelegating");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated after undelegating");
        
        // Partially slash operator
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);
        
        // Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            // Check that each withdrawal is properly completed
            assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawals[i]), 
                "withdrawal should not be pending after completion");
        }
        
        // Final state checks
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /**
     * T-85: Staker with nonzero shares, Operator allocated, Staker Delegated
     * Deallocate from opSet, Partially slash operator on opSet, Queue Full Withdrawal, Complete as Tokens
     */
    function testFuzz_deposit_delegate_deallocate_partialSlash_queue_complete(uint24 r) public rand(r) {
        // Deallocate from operator set
        operator.modifyAllocations(AllocateParams({
            operatorSet: operatorSet,
            strategies: allocateParams.strategies,
            newMagnitudes: new uint64[](allocateParams.strategies.length)
        }));
        
        // Check that the operator has deallocated
        for (uint i = 0; i < allocateParams.strategies.length; i++) {
            (uint64 currentMagnitude, ) = allocationManager.getAllocation(
                address(operator), 
                operatorSet.id, 
                allocateParams.strategies[i]
            );
            assertEq(currentMagnitude, 0, "operator should have zero allocation after deallocating");
        }
        
        // Partially slash operator
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);

        // Queue withdrawal
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, withdrawableShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        
        // Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsTokens(withdrawals[i]);
        }
        
        // Final state checks
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    /**
     * T-86: Staker with nonzero shares, Operator allocated, Staker Delegated
     * Deregister from opSet, Partially slash operator on opSet, Queue Full Withdrawal, Complete as Tokens
     */
    function testFuzz_deposit_delegate_deregister_partialSlash_queue_complete(uint24 r) public rand(r) {
        // Deregister operator from operator set
        operator.deregisterFromOperatorSet(operatorSet);
        
        // Partially slash operator
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);

        // Queue withdrawal
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, withdrawableShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        
        // Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsTokens(withdrawals[i]);
        }
        
        // Final state checks
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }
}