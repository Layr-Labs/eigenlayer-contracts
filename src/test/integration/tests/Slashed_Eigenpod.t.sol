// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_SlashedEigenpod is IntegrationCheckUtils {
    using ArrayLib for *;
    
    AVS avs;
    OperatorSet operatorSet;

    User operator;
    IAllocationManagerTypes.AllocateParams allocateParams;

    User staker;
    IStrategy[] strategies;
    uint[] initTokenBalances;
    uint64 slashedGwei;

    function _init() internal override {
        _configAssetTypes(HOLDS_ETH);
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        (operator,,) = _newRandomOperator();
        (avs,) = _newRandomAVS();

        cheats.assume(initTokenBalances[0] >= 64 ether);

        //Slash on Beacon chain
        (uint40[] memory validators,) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(validators);

        uint[] memory shares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, shares);

        uint40[] memory slashedValidators = _choose(validators);
        slashedGwei = beaconChain.slashValidators(slashedValidators);
        console.log(slashedGwei);
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        
        staker.startCheckpoint();
        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithSlashing_HandleRoundDown_State(staker, slashedValidators, slashedGwei);
    }

    function testFuzz_delegateSlashedStaker_dsfWad(uint24 _random) public rand(_random) {

        uint256[] memory initDelegatableShares = _getWithdrawableShares(staker, strategies);
        uint256[] memory initDepositShares = _getStakerDepositShares(staker, strategies);

        // Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, strategies);
        
        // Allocate to operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet, strategies);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        (uint256[] memory withdrawableSharesBefore, ) = delegationManager.getWithdrawableShares(address(staker), strategies);

        // Undelegate from an operator
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, initDepositShares, initDelegatableShares);

        // Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, initDelegatableShares);
        }

        (uint256[] memory withdrawableSharesAfter, uint256[] memory depositSharesAfter) = delegationManager.getWithdrawableShares(address(staker), strategies);
        assertEq(depositSharesAfter[0], initDelegatableShares[0], "Deposit shares should reset to reflect slash(es)");
        assertApproxEqAbs(withdrawableSharesAfter[0], depositSharesAfter[0], 100, "Withdrawable shares should equal deposit shares after withdrawal");
    }

    function testFuzz_delegateSlashedStaker_dsfNonWad(uint24 _random) public rand(_random) {

        //Additional deposit on beacon chain so dsf is nonwad
        uint amount = 32 ether * _randUint({min: 1, max: 5});
        cheats.deal(address(staker), amount);
        (uint40[] memory validators,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.verifyWithdrawalCredentials(validators);
        
        staker.startCheckpoint();
        staker.completeCheckpoint();


        uint256[] memory initDelegatableShares = _getWithdrawableShares(staker, strategies);
        uint256[] memory initDepositShares = _getStakerDepositShares(staker, strategies);
        // Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        uint256[] memory withdrawableSharesAfterDelegation = _getWithdrawableShares(staker, strategies);

        // Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, strategies);
        
        // Allocate to operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet, strategies);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        (uint256[] memory withdrawableSharesBefore, ) = delegationManager.getWithdrawableShares(address(staker), strategies);
        // Undelegate from an operator
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, initDepositShares, initDelegatableShares);

        // Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, initDelegatableShares);
        }

        (uint256[] memory withdrawableSharesAfter, uint256[] memory depositSharesAfter) = delegationManager.getWithdrawableShares(address(staker), strategies);
        assertEq(depositSharesAfter[0], initDelegatableShares[0], "Deposit shares should reset to reflect slash(es)");
        assertApproxEqAbs(withdrawableSharesAfter[0], depositSharesAfter[0], 100, "Withdrawable shares should equal deposit shares after withdrawal");
    }

    // TODO: Fix this test
    // function testFuzz_delegateSlashedStaker_slashedOperator(uint24 _random) public rand(_random) {


    //     (User staker2,,) = _newRandomStaker();
    //     (uint40[] memory validators2,) = staker2.startValidators();
    //     beaconChain.advanceEpoch_NoWithdrawNoRewards();
    //     staker2.verifyWithdrawalCredentials(validators2);
    //     staker2.startCheckpoint();
    //     staker2.completeCheckpoint();
    //     staker2.delegateTo(operator);

    //     //randomize additional deposit to eigenpod
    //     if(_randBool()){
    //         uint amount = 32 ether * _randUint({min: 1, max: 5});
    //         cheats.deal(address(staker), amount);
    //         (uint40[] memory validators,) = staker.startValidators();
    //         beaconChain.advanceEpoch_NoWithdrawNoRewards();
    //         staker.verifyWithdrawalCredentials(validators);
            
    //         staker.startCheckpoint();
    //         staker.completeCheckpoint();
    //     }

    //     // Create an operator set and register an operator.
    //     operatorSet = avs.createOperatorSet(strategies);
    //     operator.registerForOperatorSet(operatorSet);
    //     check_Registration_State_NoAllocation(operator, operatorSet, strategies);

    //     // Allocate to operator set
    //     allocateParams = _genAllocation_AllAvailable(operator, operatorSet, strategies);
    //     operator.modifyAllocations(allocateParams);
    //     check_IncrAlloc_State_Slashable(operator, allocateParams);
    //     _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

    //     //Slash operator before delegation
    //     IAllocationManagerTypes.SlashingParams memory slashingParams;
    //     uint wadToSlash = _randWadToSlash();
    //     slashingParams = avs.slashOperator(operator, operatorSet.id, strategies, wadToSlash.toArrayU256());
    //     assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");

    //     uint256[] memory initDelegatableShares = _getWithdrawableShares(staker, strategies);
    //     uint256[] memory initDepositShares = _getStakerDepositShares(staker, strategies);

    //     // Delegate to an operator
    //     staker.delegateTo(operator);
    //     (uint256[] memory delegatedShares, ) = delegationManager.getWithdrawableShares(address(staker), strategies);
    //     check_Delegation_State(staker, operator, strategies, initDepositShares);
        
    //     // Undelegate from an operator
    //     IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.undelegate();
    //     bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
    //     check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, initDepositShares, delegatedShares);

    //     // Complete withdrawal as shares
    //     // Fast forward to when we can complete the withdrawal
    //     _rollBlocksForCompleteWithdrawals(withdrawals);
    //     for (uint256 i = 0; i < withdrawals.length; ++i) {
    //         staker.completeWithdrawalAsShares(withdrawals[i]);
    //         check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, delegatedShares);
    //     }

    //     (uint256[] memory withdrawableSharesAfter, uint256[] memory depositSharesAfter) = delegationManager.getWithdrawableShares(address(staker), strategies);
    //     assertEq(depositSharesAfter[0], delegatedShares[0], "Deposit shares should reset to reflect slash(es)");
    //     assertApproxEqAbs(withdrawableSharesAfter[0], depositSharesAfter[0], 100, "Withdrawable shares should equal deposit shares after withdrawal");
    // }

    function testFuzz_delegateSlashedStaker_redelegate_complete(uint24 _random) public rand(_random){

        (User operator2, ,) = _newRandomOperator();

        //Additional deposit on beacon chain so dsf is nonwad
        uint amount = 32 ether * _randUint({min: 1, max: 5});
        cheats.deal(address(staker), amount);
        (uint40[] memory validators,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.verifyWithdrawalCredentials(validators);
        
        staker.startCheckpoint();
        staker.completeCheckpoint();


        uint256[] memory initDepositShares = _getStakerDepositShares(staker, strategies);

        // Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);
        (uint256[] memory delegatedShares, ) = delegationManager.getWithdrawableShares(address(staker), strategies);

        // Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, strategies);
        
        // Allocate to operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet, strategies);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // Undelegate from an operator
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.redelegate(operator2);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Redelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, initDepositShares, delegatedShares);

        // Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Redelegated_State(staker, operator, operator2, withdrawals[i], withdrawals[i].strategies, delegatedShares);
        }

        (uint256[] memory withdrawableSharesAfter, uint256[] memory depositSharesAfter) = delegationManager.getWithdrawableShares(address(staker), strategies);
        assertEq(depositSharesAfter[0], delegatedShares[0], "Deposit shares should reset to reflect slash(es)");
        assertApproxEqAbs(withdrawableSharesAfter[0], depositSharesAfter[0], 100, "Withdrawable shares should equal deposit shares after withdrawal");
    }

    
    function testFuzz_delegateSlashedStaker_slashedOperator_withdrawAllShares_complete(uint24 _random) public rand(_random){ 
        //Additional deposit on beacon chain so dsf is nonwad
        uint amount = 32 ether * _randUint({min: 1, max: 5});
        cheats.deal(address(staker), amount);
        (uint40[] memory validators,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.verifyWithdrawalCredentials(validators);
        
        staker.startCheckpoint();
        staker.completeCheckpoint();

        uint256[] memory initDepositShares = _getStakerDepositShares(staker, strategies);

        // Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, strategies);

        //Slash operator before delegation
        IAllocationManagerTypes.SlashingParams memory slashingParams;
        uint wadToSlash = _randWadToSlash();
        slashingParams = avs.slashOperator(operator, operatorSet.id, strategies, wadToSlash.toArrayU256());
        assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");

        // Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);
        (uint256[] memory delegatedShares, ) = delegationManager.getWithdrawableShares(address(staker), strategies);
        
        // Allocate to operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet, strategies);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        //Withdraw all shares
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, initDepositShares);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
        }

        (uint256[] memory withdrawableSharesAfter, uint256[] memory depositSharesAfter) = delegationManager.getWithdrawableShares(address(staker), strategies);
        assertEq(depositSharesAfter[0], delegatedShares[0], "Deposit shares should reset to reflect slash(es)");
        assertApproxEqAbs(withdrawableSharesAfter[0], depositSharesAfter[0], 100, "Withdrawable shares should equal deposit shares after withdrawal");
    }
    
}