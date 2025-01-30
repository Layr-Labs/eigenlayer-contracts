// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract SlashedEigenpod is IntegrationCheckUtils {
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

        uint40[] memory slashedValidators = _chooseSubset(validators);
        slashedGwei = beaconChain.slashValidators(slashedValidators);
        beaconChain.advanceEpoch_NoRewards();
        
        staker.startCheckpoint();
        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithSlashing_State(staker, slashedValidators, slashedGwei);
    }

    function testFuzz_delegateSlashedStaker_dsfWad(uint24 _random) public rand(_random) {

        uint256[] memory initDepositShares = _getStakerDepositShares(staker, strategies);

        // 1. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // 2. Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);
        // TODO invariant checks here
        operator.registerForOperatorSet(operatorSet);
        // TODO invariant checks here
        
        // 3. Allocate to operator set
        allocateParams = operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
        assert_Snap_Allocations_Modified(
            operator, allocateParams, false, "operator allocations should be updated before delay"
        );
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
        assert_Snap_Allocations_Modified(
            operator, allocateParams, true, "operator allocations should be updated after delay"
        );

        // 4. Undelegate from an operator
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, initDepositShares);

        // 5.. Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, withdrawals[i].scaledShares);
        }

        (uint256[] memory withdrawableSharesAfter, uint256[] memory depositSharesAfter) = delegationManager.getWithdrawableShares(address(staker), strategies);
        assertEq(depositSharesAfter[0], initDepositShares[0] - slashedGwei, "Deposit shares should reset to reflect slash(es)");
        assertEq(withdrawableSharesAfter[0], depositSharesAfter[0], "Withdrawable shares should equal deposit shares after withdrawal");
    }

    function testFuzz_delegateSlashedStaker_dsfNonWad(uint24 _random) public rand(_random) {

        //Generate rewards on beacon chain so dsf is nonWad
        beaconChain.advanceEpoch();
        
        staker.startCheckpoint();
        staker.completeCheckpoint();


        uint256[] memory initDepositShares = _getStakerDepositShares(staker, strategies);

        // Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);
        // TODO invariant checks here
        operator.registerForOperatorSet(operatorSet);
        // TODO invariant checks here
        
        // Allocate to operator set
        allocateParams = operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
        assert_Snap_Allocations_Modified(
            operator, allocateParams, false, "operator allocations should be updated before delay"
        );
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
        assert_Snap_Allocations_Modified(
            operator, allocateParams, true, "operator allocations should be updated after delay"
        );

        // Undelegate from an operator
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, initDepositShares);

        // Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Undelegated_State(staker, operator, withdrawals[i], withdrawals[i].strategies, withdrawals[i].scaledShares);
        }

        (uint256[] memory withdrawableSharesAfter, uint256[] memory depositSharesAfter) = delegationManager.getWithdrawableShares(address(staker), strategies);
        assertEq(depositSharesAfter[0], initDepositShares[0] - slashedGwei, "Deposit shares should reset to reflect slash(es)");
        assertEq(withdrawableSharesAfter[0], depositSharesAfter[0], "Withdrawable shares should equal deposit shares after withdrawal");
    }

    function testFuzz_delegateSlashedStaker_slashedOperator(uint24 _random) public rand(_random) {
        //randomize additional deposit to eigenpod
        if(_randBool()){
            beaconChain.advanceEpoch();
        
            staker.startCheckpoint();
            staker.completeCheckpoint();
        }

        // Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);
        // TODO invariant checks here
        operator.registerForOperatorSet(operatorSet);
        // TODO invariant checks here

        //Slash operator before delegation
        IAllocationManagerTypes.SlashingParams memory slashingParams;
        uint wadToSlash = _randWadToSlash();
        slashingParams = avs.slashOperator(operator, operatorSet.id, strategies, wadToSlash.toArrayU256());
        assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");

        uint256[] memory initDepositShares = _getStakerDepositShares(staker, strategies);

        // Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);
        
        // Allocate to operator set
        allocateParams = operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
        assert_Snap_Allocations_Modified(
            operator, allocateParams, false, "operator allocations should be updated before delay"
        );
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
        assert_Snap_Allocations_Modified(
            operator, allocateParams, true, "operator allocations should be updated after delay"
        );

        // Undelegate from an operator
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, initDepositShares);

        // Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State_AfterSlash(staker, operator, withdrawals[i], allocateParams, slashingParams);
        }

        (uint256[] memory withdrawableSharesAfter, uint256[] memory depositSharesAfter) = delegationManager.getWithdrawableShares(address(staker), strategies);
        assertEq(depositSharesAfter[0], initDepositShares[0] - slashedGwei, "Deposit shares should reset to reflect slash(es)");
        assertEq(withdrawableSharesAfter[0], depositSharesAfter[0], "Withdrawable shares should equal deposit shares after withdrawal");
    }

    function testFuzz_delegateSlashedStaker_redelegate_complete(uint24 _random) public rand(_random){

        (User operator2, ,) = _newRandomOperator();

        //Generate rewards on beacon chain so dsf is nonWad
        beaconChain.advanceEpoch();
        
        staker.startCheckpoint();
        staker.completeCheckpoint();


        uint256[] memory initDepositShares = _getStakerDepositShares(staker, strategies);

        // Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // Create an operator set and register an operator.
        operatorSet = avs.createOperatorSet(strategies);
        // TODO invariant checks here
        operator.registerForOperatorSet(operatorSet);
        // TODO invariant checks here
        
        // Allocate to operator set
        allocateParams = operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
        assert_Snap_Allocations_Modified(
            operator, allocateParams, false, "operator allocations should be updated before delay"
        );
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
        assert_Snap_Allocations_Modified(
            operator, allocateParams, true, "operator allocations should be updated after delay"
        );

        // Undelegate from an operator
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.redelegate(operator2);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        check_Redelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, initDepositShares);

        // Complete withdrawal as shares
        // Fast forward to when we can complete the withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_Redelegated_State(staker, operator, operator2, withdrawals[i], withdrawals[i].strategies, withdrawals[i].scaledShares);
        }

        (uint256[] memory withdrawableSharesAfter, uint256[] memory depositSharesAfter) = delegationManager.getWithdrawableShares(address(staker), strategies);
        assertEq(depositSharesAfter[0], initDepositShares[0] - slashedGwei, "Deposit shares should reset to reflect slash(es)");
        assertEq(withdrawableSharesAfter[0], depositSharesAfter[0], "Withdrawable shares should equal deposit shares after withdrawal");
    }
}