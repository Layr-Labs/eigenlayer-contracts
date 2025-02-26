// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_SlashedEigenpod_BC is IntegrationCheckUtils {
    using ArrayLib for *;
    
    AVS avs;
    OperatorSet operatorSet;

    User operator;
    AllocateParams allocateParams;
    SlashingParams slashParams;

    User staker;
    IStrategy[] strategies;
    uint[] initTokenBalances;
    uint[] initDepositShares;

    function _init() internal override {
        _configAssetTypes(HOLDS_ETH);
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        (operator,,) = _newRandomOperator();
        (avs,) = _newRandomAVS();

        cheats.assume(initTokenBalances[0] >= 64 ether);

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, initDepositShares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // 3. Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);

        // 4. Register for operator set
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, allStrats);

        // 5. Allocate to operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // 6. Slash operatorSet
        slashParams = _genSlashing_Half(operator, operatorSet);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);
    }

    /// @dev Asserts that the DSF isn't updated after a queued withdrawal and a checkpoint with 0 balance
    function testFuzz_deposit_delegate_allocate_slash_queueWithdrawal_checkpointZeroBalance(uint24 _rand) public rand(_rand) {
        beaconChain.advanceEpoch_NoRewards();

        // 7. Queue Withdrawal for all shares
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        // TODO: assert this properly

        // 8. Start & complete checkpoint
        staker.startCheckpoint();
        check_StartCheckpoint_State(staker);
        staker.completeCheckpoint();
        check_CompleteCheckpoint_ZeroBalanceDelta_State(staker);
        require(false==true);
    }

}