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
}
