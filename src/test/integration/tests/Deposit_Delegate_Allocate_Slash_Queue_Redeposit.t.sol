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

    function testFuzz_deposit_delegate_allocate_partialSlash_redeposit_queue_complete(uint24 r) public rand(r) {
        // Partially slash operator
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

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
    
    function testFuzz_deposit_delegate_undelegate_partialSlash_complete(uint24 r) public rand(r) {
        // Undelegate from operator
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        assert_AllWithdrawalsPending(withdrawalRoots, "withdrawals should be pending after undelegating");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated after undelegating");
        
        // Partially slash operator
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);
        
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

    function testFuzz_deposit_delegate_deallocate_partialSlash_queue_complete(uint24 r) public rand(r) {
        // Deallocate from operator set
        operator.modifyAllocations(AllocateParams({
            operatorSet: operatorSet,
            strategies: allocateParams.strategies,
            newMagnitudes: new uint64[](allocateParams.strategies.length)
        }));
        
        // Partially slash operator
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

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

    function testFuzz_deposit_delegate_deregister_partialSlash_queue_complete(uint24 r) public rand(r) {
        // Deregister operator from operator set
        operator.deregisterFromOperatorSet(operatorSet);
        
        // Partially slash operator
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

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

    function testFuzz_delegate_zeroShares_partialSlash_deposit_undelegate_complete(uint24 r) public rand(r) {
        // Create a new staker with 0 shares
        (User zeroSharesStaker, uint[] memory tokensToDeposit) = _newStaker(strategies);
        
        // Delegate to operator (with 0 shares)
        zeroSharesStaker.delegateTo(operator);
        
        // Partially slash operator
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);
        
        // Deposit tokens
        uint[] memory depositShares = _calculateExpectedShares(strategies, tokensToDeposit);
        zeroSharesStaker.depositIntoEigenlayer(strategies, tokensToDeposit);
        check_Deposit_State(zeroSharesStaker, strategies, depositShares);
        
        // Undelegate
        Withdrawal[] memory withdrawals = zeroSharesStaker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        
        // Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            zeroSharesStaker.completeWithdrawalAsTokens(withdrawals[i]);
        }
        
        // Final state checks
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_deposit_delegate_allocate_partialSlash_deallocate(uint24 r) public rand(r) {
        // Partially slash operator
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);
        
        // Deallocate from operator set
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_DecrAlloc_State_Slashable(operator, deallocateParams);
    }
}