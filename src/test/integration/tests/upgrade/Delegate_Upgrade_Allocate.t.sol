// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/tests/upgrade/UpgradeTest.t.sol";

contract Integration_Upgrade_Deposit_Delegate_Allocate is UpgradeTest {
    struct TestState {
        User staker;
        User operator;
        AVS avs;
        IStrategy[] strategies;
        uint[] initTokenBalances;
        OperatorSet operatorSet;
        AllocateParams allocateParams;
    }

    function _init_() internal returns (TestState memory state) {
        (state.staker, state.strategies, state.initTokenBalances) = _newRandomStaker();
        state.operator = _newRandomOperator();

        // Pre-upgrade:
        // 1. Create staker and operator with assets, then deposit into EigenLayer
        // 2. Delegate to operator
        state.staker.depositIntoEigenlayer(state.strategies, state.initTokenBalances);
        state.staker.delegateTo(state.operator);
    }

    function _setupAllocation(TestState memory state) internal {
        // 3. Set allocation delay for operator
        state.operator.setAllocationDelay(1);
        rollForward({blocks: ALLOCATION_CONFIGURATION_DELAY + 1});

        // 4. Create an operator set and register an operator.
        state.operatorSet = state.avs.createOperatorSet(state.strategies);
        state.operator.registerForOperatorSet(state.operatorSet);
        check_Registration_State_NoAllocation(state.operator, state.operatorSet, allStrats);

        // 5. Allocate to operator set.
        state.allocateParams = _genAllocation_Rand(state.operator, state.operatorSet);
        state.operator.modifyAllocations(state.allocateParams);
        check_IncrAlloc_State_Slashable(state.operator, state.allocateParams);
    }

    function testFuzz_deposit_delegate_upgrade_allocate(uint24) public {
        TestState memory state = _init_();
        _upgradeEigenLayerContracts();
        state.avs = _newRandomAVS();
        _setupAllocation(state);
    }
}
