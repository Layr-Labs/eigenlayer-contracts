// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

contract Integration_Upgrade_Deposit_Delegate_Allocate is UpgradeTest {

    function testFuzz_deposit_delegate_upgrade_allocate(uint24 _random) public rand(_random) {
        (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();

        // Pre-upgrade:
        // 1. Create staker and operator with assets, then deposit into EigenLayer
        // 2. Delegate to operator
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        staker.delegateTo(operator);

        // Upgrade to slashing release
        _upgradeEigenLayerContracts();
        (AVS avs,) = _newRandomAVS();

        // 3. Set allocation delay for operator
        operator.setAllocationDelay(1);
        rollForward({blocks: ALLOCATION_CONFIGURATION_DELAY + 1});

        // 4. Create an operator set and register an operator.
        OperatorSet memory operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, allStrats);

        // 5. Allocate to operator set.
        AllocateParams memory allocateParams = AllocateParams({
            operatorSet: operatorSet,
            strategies: strategies,
            newMagnitudes: _randMagnitudes({sum: 1 ether, len: strategies.length})
        });
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);
    }
}