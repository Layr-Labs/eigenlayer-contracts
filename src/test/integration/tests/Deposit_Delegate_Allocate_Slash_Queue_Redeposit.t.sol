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
        _configUserTypes(DEFAULT);

        (staker, strategies, initTokenBalances) = _newRandomStaker();
        (operator,,) = _newRandomOperator();
        (avs,) = _newRandomAVS();
        tokens = _getUnderlyingTokens(strategies);

        uint[] memory tokensToDeposit = new uint[](initTokenBalances.length);
        numTokensRemaining = new uint[](initTokenBalances.length);
        uint eth_to_deal;
        for (uint i = 0; i < initTokenBalances.length; i++) {
            if (strategies[i] == BEACONCHAIN_ETH_STRAT) {
                tokensToDeposit[i] = initTokenBalances[i];
                //user.depositIntoEigenlayer uses all ETH balance for a pod, so we deal back staker's
                //starting ETH to replicate partial deposit state
                eth_to_deal = initTokenBalances[i];
                numTokensRemaining[i] = initTokenBalances[i];
                continue;
            }

            tokensToDeposit[i] = initTokenBalances[i] / 2;
            numTokensRemaining[i] = initTokenBalances[i] - tokensToDeposit[i];
        }

        // 1. Deposit Into Strategies
        initDepositShares = _calculateExpectedShares(strategies, tokensToDeposit);
        staker.depositIntoEigenlayer(strategies, tokensToDeposit);
        //dealing back ETH
        cheats.deal(address(staker), eth_to_deal);
        check_Deposit_State_PartialDeposit(staker, strategies, initDepositShares, numTokensRemaining);

        // 2. Delegate to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

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

    function testFuzz_fullSlash_undelegate_complete_redeposit(uint24 _random) public rand(_random) {
        // 4. Fully slash operator
        SlashingParams memory slashParams = _genSlashing_Full(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_FullySlashed_State(operator, allocateParams, slashParams);

        // 5. Undelegate from an operator
        uint[] memory shares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory roots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, roots, strategies, shares);

        // 6. Complete withdrawal. Staker should receive 0 shares/tokens after a full slash
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(
                staker, operator, withdrawals[i], strategies, new uint[](strategies.length), tokens, new uint[](strategies.length)
            );
        }

        // 7. Redeposit
        uint[] memory depositShares = _calculateExpectedShares(strategies, numTokensRemaining);
        staker.depositIntoEigenlayer(strategies, numTokensRemaining);
        check_Deposit_State(staker, strategies, depositShares);
    }

    function testFuzz_undelegate_fullSlash_complete_redeposit(uint24 _random) public rand(_random) {
        // 4. Undelegate from an operator
        uint[] memory shares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory roots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, roots, strategies, shares);

        // 5. Fully slash operator
        SlashingParams memory slashParams = _genSlashing_Full(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_FullySlashed_State(operator, allocateParams, slashParams);

        // 6. Complete withdrawal. Staker should receive 0 shares/tokens after a full slash
        uint[] memory expectedShares = new uint[](strategies.length);
        uint[] memory expectedTokens = new uint[](strategies.length);

        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], strategies, expectedShares, tokens, expectedTokens);
        }

        // 7. Redeposit
        uint[] memory depositShares = _calculateExpectedShares(strategies, numTokensRemaining);
        staker.depositIntoEigenlayer(strategies, numTokensRemaining);
        check_Deposit_State(staker, strategies, depositShares);
    }

    function testFuzz_depositFull_fullSlash_undelegate_completeAsShares(uint24 _random) public rand(_random) {
        uint[] memory depositShares = _calculateExpectedShares(strategies, numTokensRemaining);
        staker.depositIntoEigenlayer(strategies, numTokensRemaining);
        check_Deposit_State(staker, strategies, depositShares);

        // 4. Fully slash random proper subset of operators strategies
        SlashingParams memory slashParams = _genSlashing_Full(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_FullySlashed_State(operator, allocateParams, slashParams);

        // add deposit shares to initDepositShares
        for (uint i = 0; i < depositShares.length; i++) {
            initDepositShares[i] += depositShares[i];
        }

        // 5. Undelegate from an operator
        uint[] memory shares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory roots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, roots, strategies, shares);

        // 6. Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedShares = _calculateExpectedShares(withdrawals[i]);
            tokens = staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, expectedShares);
        }
    }

    function testFuzz_deposit_delegate_allocate_partialSlash_redeposit_queue_complete(uint24 r) public rand(r) {
        // Partially slash operator
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

        // Redeposit remaining tokens
        uint[] memory depositShares = _calculateExpectedShares(strategies, numTokensRemaining);
        staker.depositIntoEigenlayer(strategies, numTokensRemaining);
        check_Deposit_State(staker, strategies, depositShares);

        for (uint i = 0; i < depositShares.length; i++) {
            initDepositShares[i] += depositShares[i];
        }

        // Queue withdrawal
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, initDepositShares, withdrawableShares, withdrawals, withdrawalRoots);

        // Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedShares = _calculateExpectedShares(withdrawals[i]);
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, expectedShares);
            tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(
                staker, operator, withdrawals[i], withdrawals[i].strategies, expectedShares, tokens, expectedTokens
            );
        }
    }

    function testFuzz_deposit_delegate_undelegate_partialSlash_complete(uint24 r) public rand(r) {
        // Undelegate from operator
        uint[] memory shares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory roots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, roots, strategies, shares);

        // Partially slash operator
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

        // Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedShares = _calculateExpectedShares(withdrawals[i]);
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, expectedShares);
            tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(
                staker, operator, withdrawals[i], withdrawals[i].strategies, expectedShares, tokens, expectedTokens
            );
        }
    }

    function testFuzz_deposit_delegate_deallocate_partialSlash_queue_complete(uint24 r) public rand(r) {
        // Deallocate from operator set
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator, operatorSet);
        operator.modifyAllocations(deallocateParams);
        check_DecrAlloc_State_Slashable(operator, deallocateParams);

        // Partially slash operator
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

        // Queue withdrawal
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, initDepositShares, withdrawableShares, withdrawals, withdrawalRoots);

        // Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedShares = _calculateExpectedShares(withdrawals[i]);
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, expectedShares);
            tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(
                staker, operator, withdrawals[i], withdrawals[i].strategies, expectedShares, tokens, expectedTokens
            );
        }
    }

    function testFuzz_deposit_delegate_deregister_partialSlash_queue_complete(uint24 r) public rand(r) {
        // Deregister operator from operator set
        operator.deregisterFromOperatorSet(operatorSet);
        check_Deregistration_State_ActiveAllocation(operator, operatorSet);

        // Partially slash operator
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

        // Queue withdrawal
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_QueuedWithdrawal_State(staker, operator, strategies, initDepositShares, withdrawableShares, withdrawals, withdrawalRoots);

        // Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedShares = _calculateExpectedShares(withdrawals[i]);
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, expectedShares);
            tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(
                staker, operator, withdrawals[i], withdrawals[i].strategies, expectedShares, tokens, expectedTokens
            );
        }
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
        uint[] memory shares = _getStakerWithdrawableShares(zeroSharesStaker, strategies);
        Withdrawal[] memory withdrawals = zeroSharesStaker.undelegate();
        bytes32[] memory roots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(zeroSharesStaker, operator, withdrawals, roots, strategies, shares);

        // Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedShares = _calculateExpectedShares(withdrawals[i]);
            uint[] memory expectedTokens = _calculateExpectedTokens(withdrawals[i].strategies, expectedShares);
            tokens = zeroSharesStaker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(
                zeroSharesStaker, operator, withdrawals[i], withdrawals[i].strategies, expectedShares, tokens, expectedTokens
            );
        }
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

    function testFuzz_fullSlash_undelegate_redeposit_complete(uint24 _random) public rand(_random) {
        initDepositShares = _getStakerDepositShares(staker, strategies);

        // 4. Fully slash operator
        SlashingParams memory slashParams = _genSlashing_Full(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_FullySlashed_State(operator, allocateParams, slashParams);

        // 5. Undelegate from an operator
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, withdrawableShares);

        // 6. Redeposit
        uint[] memory shares = _calculateExpectedShares(strategies, numTokensRemaining);
        staker.depositIntoEigenlayer(strategies, numTokensRemaining);
        check_Deposit_State(staker, strategies, shares);

        // 7. Complete withdrawal. Staker should receive 0 shares/tokens after a full slash
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedShares = _calculateExpectedShares(withdrawals[i]);
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, expectedShares);
        }

        // Final state checks
        assert_HasExpectedShares(staker, strategies, shares, "staker should have expected shares after redeposit");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_fullSlash_redelegate_redeposit_complete(uint24 _random) public rand(_random) {
        (User operator2,,) = _newRandomOperator();
        initDepositShares = _getStakerDepositShares(staker, strategies);

        // 4. Fully slash operator
        SlashingParams memory slashParams = _genSlashing_Full(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_FullySlashed_State(operator, allocateParams, slashParams);

        // 5. Undelegate from an operator
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(staker, strategies);
        Withdrawal[] memory withdrawals = staker.redelegate(operator2);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Redelegate_State(staker, operator, operator2, withdrawals, withdrawalRoots, strategies, withdrawableShares);

        // 6. Redeposit
        uint[] memory shares = _calculateExpectedShares(strategies, numTokensRemaining);
        staker.depositIntoEigenlayer(strategies, numTokensRemaining);
        check_Deposit_State(staker, strategies, shares);

        // 7. Complete withdrawal. Staker should receive 0 shares/tokens after a full slash
        _rollBlocksForCompleteWithdrawals(withdrawals);

        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedShares = _calculateExpectedShares(withdrawals[i]);
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, expectedShares);
        }

        // Final state checks
        assert_HasExpectedShares(staker, strategies, shares, "staker should have expected shares after redeposit");
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }
}
