// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_QueueSlashAccounting is IntegrationCheckUtils {
    using ArrayLib for *;

    AVS avs;
    OperatorSet operatorSet;
    User operator;
    IStrategy strategy;
    IStrategy[] strategies;
    AllocateParams allocateParams;

    struct QueueRedelegationContext {
        User queueStaker;
        User redelegatingStaker;
        User newOperator;
        uint delegatedBeforeSlash;
        uint slashedBeforeQueue;
        uint extraDepositShares;
        uint activeBeforeQueue;
        Withdrawal[] queuedWithdrawals;
        Withdrawal[] redelegatedWithdrawals;
    }

    function _assertQueueSlashableBacked(User queueOperator, uint activeBeforeQueue, IStrategy queuedStrategy)
        internal
        view
        returns (uint removedFromOperator, uint queueSlashable)
    {
        uint activeAfterQueue = delegationManager.operatorShares(address(queueOperator), queuedStrategy);
        removedFromOperator = activeBeforeQueue - activeAfterQueue;
        queueSlashable = delegationManager.getSlashableSharesInQueue(address(queueOperator), queuedStrategy);
        assertLe(queueSlashable, removedFromOperator, "queue slashable should not exceed removed backing");
    }

    function _init() internal override {
        _configAssetTypes(HOLDS_LST);
        _configUserTypes(DEFAULT);

        operator = _newRandomOperator_NoAssets();
        (avs,) = _newRandomAVS();

        strategy = lstStrats[0];
        strategies = strategy.toArray();

        // 1. Create operator set and register operator
        operatorSet = avs.createRedistributingOperatorSet(strategies, address(0xBEEF));
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, allStrats);
    }

    function testFuzz_deposit_delegate_slash_queue_slash_noQueueOverAccounting(uint24 _random) public rand(_random) {
        uint numStakers = _randUint(2, 10);
        uint depositTokens = _randUint(1, 1e18);
        User[] memory stakers = new User[](numStakers);
        uint[] memory stakerDepositShares = new uint[](numStakers);
        uint delegatedBefore;

        for (uint i = 0; i < numStakers; ++i) {
            stakers[i] = _newEmptyStaker();
            uint[] memory tokenBalances = depositTokens.toArrayU256();

            // 3. Deposit into strategy
            _dealAmounts(stakers[i], strategies, tokenBalances);
            stakers[i].depositIntoEigenlayer(strategies, tokenBalances);
            uint[] memory depositShares = _calculateExpectedShares(strategies, tokenBalances);
            stakerDepositShares[i] = depositShares[0];
            delegatedBefore += depositShares[0];
            check_Deposit_State(stakers[i], strategies, depositShares);

            // 4. Delegate to operator
            stakers[i].delegateTo(operator);
            check_Delegation_State(stakers[i], operator, strategies, depositShares);
        }

        assertEq(
            delegationManager.operatorShares(address(operator), strategy), delegatedBefore, "operator shares should equal delegated shares"
        );

        // 5. Allocate to operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // 6. Slash operator by 50%
        SlashingParams memory slashA = _genSlashing_Custom(operator, operatorSet, 5e17);
        (uint slashIdA, uint[] memory sharesA) = avs.slashOperator(slashA);
        check_Base_Slashing_State(operator, allocateParams, slashA, slashIdA);

        uint activeBeforeQueue = delegationManager.operatorShares(address(operator), strategy);
        for (uint i = 0; i < numStakers; ++i) {
            uint[] memory withdrawableShares = _getStakerWithdrawableShares(stakers[i], strategies);

            // 7. Queue full withdrawal
            Withdrawal[] memory withdrawals = stakers[i].queueWithdrawals(strategies, stakerDepositShares[i].toArrayU256());
            bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
            check_QueuedWithdrawal_State(
                stakers[i], operator, strategies, stakerDepositShares[i].toArrayU256(), withdrawableShares, withdrawals, withdrawalRoots
            );
        }

        uint activeAfterQueue = delegationManager.operatorShares(address(operator), strategy);
        uint removedFromOperator = activeBeforeQueue - activeAfterQueue;
        uint queueSlashable = delegationManager.getSlashableSharesInQueue(address(operator), strategy);
        assertLe(queueSlashable, removedFromOperator, "queue slashable should not exceed removed backing");

        // 8. Fully slash remaining operator magnitude
        SlashingParams memory slashB = _genSlashing_Full(operator, operatorSet);
        (, uint[] memory sharesB) = avs.slashOperator(slashB);

        assertLe(sharesA[0] + sharesB[0], delegatedBefore, "total slashed shares should not exceed delegated shares");
    }

    function testFuzz_slash_deposit_queue_redelegate_slash_completeAsShares_noQueueOverAccounting(uint24 _random) public rand(_random) {
        QueueRedelegationContext memory ctx =
            _createSlashedStakersWithExtraDeposit(_randUint(10, 1e18), _randUint(10, 1e18), _randUint(10, 1e18));
        ctx = _queueFullAndRedelegate(ctx);

        SlashingParams memory slashB = _genSlashing_Custom(operator, operatorSet, 5e17);
        (, uint[] memory sharesB) = avs.slashOperator(slashB);
        assertLe(
            ctx.slashedBeforeQueue + sharesB[0],
            ctx.delegatedBeforeSlash + ctx.extraDepositShares,
            "total slashed shares should not exceed old-operator delegated shares"
        );

        _completeWithdrawalsAsShares(ctx.queueStaker, operator, ctx.queuedWithdrawals);
        _completeWithdrawalsAsSharesAfterRedelegation(ctx.redelegatingStaker, ctx.newOperator, ctx.redelegatedWithdrawals);
    }

    function _createSlashedStakersWithExtraDeposit(
        uint queueInitialTokens,
        uint redelegatingInitialTokens,
        uint extraTokens
    ) internal returns (QueueRedelegationContext memory ctx) {
        ctx.queueStaker = _newEmptyStaker();
        ctx.redelegatingStaker = _newEmptyStaker();
        ctx.newOperator = _newRandomOperator_NoAssets();

        _depositAndDelegate(ctx.queueStaker, queueInitialTokens);
        _depositAndDelegate(ctx.redelegatingStaker, redelegatingInitialTokens);

        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        ctx.delegatedBeforeSlash = delegationManager.operatorShares(address(operator), strategy);
        SlashingParams memory slashA = _genSlashing_Custom(operator, operatorSet, 5e17);
        (uint slashIdA, uint[] memory sharesA) = avs.slashOperator(slashA);
        check_Base_Slashing_State(operator, allocateParams, slashA, slashIdA);
        ctx.slashedBeforeQueue = sharesA[0];

        uint[] memory extraTokenBalances = extraTokens.toArrayU256();
        _dealAmounts(ctx.queueStaker, strategies, extraTokenBalances);
        ctx.queueStaker.depositIntoEigenlayer(strategies, extraTokenBalances);
        uint[] memory extraDepositShares = _calculateExpectedShares(strategies, extraTokenBalances);
        check_Deposit_State(ctx.queueStaker, strategies, extraDepositShares);
        ctx.extraDepositShares = extraDepositShares[0];
    }

    function _depositAndDelegate(User staker, uint tokenAmount) internal {
        uint[] memory tokenBalances = tokenAmount.toArrayU256();
        _dealAmounts(staker, strategies, tokenBalances);
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory depositShares = _calculateExpectedShares(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, depositShares);

        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, depositShares);
    }

    function _queueFullAndRedelegate(
        QueueRedelegationContext memory ctx
    ) internal returns (QueueRedelegationContext memory) {
        uint[] memory depositShares = _getStakerDepositShares(ctx.queueStaker, strategies);
        uint[] memory withdrawableShares = _getStakerWithdrawableShares(ctx.queueStaker, strategies);
        ctx.activeBeforeQueue = delegationManager.operatorShares(address(operator), strategy);
        ctx.queuedWithdrawals = ctx.queueStaker.queueWithdrawals(strategies, depositShares);
        bytes32[] memory queuedWithdrawalRoots = _getWithdrawalHashes(ctx.queuedWithdrawals);
        check_QueuedWithdrawal_State(ctx.queueStaker, operator, strategies, depositShares, withdrawableShares, ctx.queuedWithdrawals, queuedWithdrawalRoots);

        _assertQueueSlashableBacked(operator, ctx.activeBeforeQueue, strategy);

        uint[] memory remainingWithdrawableShares = _getStakerWithdrawableShares(ctx.redelegatingStaker, strategies);
        ctx.redelegatedWithdrawals = ctx.redelegatingStaker.redelegate(ctx.newOperator);
        bytes32[] memory redelegatedWithdrawalRoots = _getWithdrawalHashes(ctx.redelegatedWithdrawals);
        check_Redelegate_State(
            ctx.redelegatingStaker,
            operator,
            ctx.newOperator,
            ctx.redelegatedWithdrawals,
            redelegatedWithdrawalRoots,
            strategies,
            remainingWithdrawableShares
        );

        (uint removedFromOperator,) = _assertQueueSlashableBacked(operator, ctx.activeBeforeQueue, strategy);
        assertEq(
            removedFromOperator,
            ctx.activeBeforeQueue - delegationManager.operatorShares(address(operator), strategy),
            "removed backing should include queue and redelegation"
        );
        return ctx;
    }

    function _completeWithdrawalsAsShares(
        User staker,
        User currentOperator,
        Withdrawal[] memory withdrawals
    ) internal {
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedShares = _calculateExpectedShares(withdrawals[i]);
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, currentOperator, withdrawals[i], withdrawals[i].strategies, expectedShares);
        }
    }

    function _completeWithdrawalsAsSharesAfterRedelegation(
        User staker,
        User newOperator,
        Withdrawal[] memory withdrawals
    ) internal {
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; ++i) {
            uint[] memory expectedShares = _calculateExpectedShares(withdrawals[i]);
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Redelegated_State(
                staker, operator, newOperator, withdrawals[i], withdrawals[i].strategies, expectedShares
            );
        }
    }
}
