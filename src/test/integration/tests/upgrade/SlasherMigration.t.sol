// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

contract Integration_Upgrade_SlasherMigration_Base is UpgradeTest {
    AVS avs;
    OperatorSet operatorSet;

    User operator;
    AllocateParams allocateParams;
    SlashingParams slashParams;

    User staker;
    IStrategy[] strategies;
    IERC20[] tokens; // underlying token for each strategy
    uint[] initTokenBalances;
    uint[] initDepositShares;

    // List of all avss/operatorSets to be created, these do not have any stakers/operators/allocations, except for the first one
    AVS[] avsList;
    OperatorSet[] operatorSetList;

    /// Shared setup:
    ///
    /// 1. Generate staker, operator, and AVS
    /// 2. Staker deposits and delegates to operator
    /// 3. AVS creates an operator set containing the strategies held by the staker
    /// 4. Operator allocates to operator set
    /// 5. Operator registers for operator set
    /// 6. Create 6 more avss/operatorSets

    function _init() internal override {
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator_NoAssets();
        (avs,) = _newRandomAVS_v1CreateSet();
        tokens = _getUnderlyingTokens(strategies);
        avsList.push(avs);

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);

        // 2. Delegate to an operator
        staker.delegateTo(operator);

        // 3. Create an operator set with strategies and register an operator.
        operatorSet = avs.createOperatorSet_v1(strategies);
        operatorSetList.push(operatorSet);
        // Roll forward to after the allocation delay completes. We need to roll forward since the deploy version on mainnet
        // still has to wait a `ALLOCATION_CONFIGURATION_DELAY` to allocate slashable stake.
        _rollForward_AllocationConfigurationDelay();

        // 4. Operator allocates to operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);

        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // 5. Operator registers for operator set
        operator.registerForOperatorSet(operatorSet);

        // 6. Create 5 more avss/operatorSets (pre-upgrade)
        for (uint i = 0; i < 5; i++) {
            (AVS newAVS, OperatorSet[] memory newOperatorSets) = _newRandomAVS_v1CreateSet();
            avsList.push(newAVS);
            for (uint j = 0; j < newOperatorSets.length; j++) {
                operatorSetList.push(newOperatorSets[j]);
            }
        }
    }
}

contract Integration_Upgrade_SlasherMigration is Integration_Upgrade_SlasherMigration_Base {
    function testFuzz_upgrade_migrate(uint24 r) public rand(r) {
        // 1. Upgrade contracts
        _upgradeEigenLayerContracts();

        // Expect all slashers are 0 prior to migration
        for (uint i = 0; i < operatorSetList.length; i++) {
            assertEq(allocationManager.getSlasher(operatorSetList[i]), address(0), "slasher should be the 0 address");
        }

        // 2. Migrate slashers
        allocationManager.migrateSlashers(operatorSetList);

        // Check that the slashers are set to the AVS
        for (uint i = 0; i < operatorSetList.length; i++) {
            assertEq(allocationManager.getSlasher(operatorSetList[i]), address(operatorSetList[i].avs), "slasher should be the AVS");
        }
    }

    function testFuzz_setAppointee_upgrade_migrate(uint24 r) public rand(r) {
        // 1. Set appointee for the slasher for the first operatorSet
        address appointee = address(0x1);
        avsList[0].addAppointee(appointee, address(allocationManager), allocationManager.slashOperator.selector);

        // 2. Upgrade contracts
        _upgradeEigenLayerContracts();

        // 3. Migrate slashers
        allocationManager.migrateSlashers(operatorSetList);

        // Check the slashers are properly set
        for (uint i = 0; i < operatorSetList.length; i++) {
            OperatorSet memory operatorSet = operatorSetList[i];
            if (operatorSet.avs == address(avsList[0])) {
                assertEq(allocationManager.getSlasher(operatorSet), appointee, "slasher should be the appointee");
            } else {
                assertEq(allocationManager.getSlasher(operatorSet), address(operatorSet.avs), "slasher should be the AVS");
            }
        }
    }

    function testFuzz_upgrade_migrate_updateSlasher_slash(uint24 r) public rand(r) {
        // 1. Upgrade contracts
        _upgradeEigenLayerContracts();

        // 2. Migrate slashers
        allocationManager.migrateSlashers(operatorSetList);

        // 3. Update the slasher
        address slasher = address(0x2);
        avs.updateSlasher(operatorSet.id, slasher);
        _rollForward_AllocationConfigurationDelay();

        // 4. Slash the operator
        slashParams = _genSlashing_Rand(operator, operatorSet);
        cheats.prank(slasher);
        (uint slashId,) = allocationManager.slashOperator(address(avs), slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams, slashId);
    }
}
