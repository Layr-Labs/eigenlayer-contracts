// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract SlashedEigenpod is IntegrationCheckUtils {
    
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

        (,uint256[] memory depositSharesAfter) = delegationManager.getDepositedShares(address(staker));
        assertEq(depositSharesAfter[0], initDepositShares[0] - slashedGwei, "Deposit shares should not have changed");
    }

}