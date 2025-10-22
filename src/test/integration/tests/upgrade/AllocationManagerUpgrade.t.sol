// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";
import {OperatorSet} from "src/contracts/libraries/OperatorSetLib.sol";

contract Integration_Upgrade_AllocationManager is UpgradeTest {
    using ArrayLib for *;
    using StdStyle for *;
    // Storage state snapshots

    struct AllocationManagerState {
        // Storage variables that should persist across upgrades
        bool avsRegisteredMetadata;
        uint slashCount;
        address redistributionRecipient;
        bool isRedistributingOperatorSet;
        bool isOperatorSet;
        uint operatorSetCount;
        bool isMemberOfOperatorSet;
        uint memberCount;
        IStrategy[] strategiesInOperatorSet;
        bool isOperatorSlashable;
        bool isOperatorRedistributable;
        uint64 encumberedMagnitude;
        uint64 allocatableMagnitude;
        uint64 maxMagnitude;
        bool allocationDelayIsSet;
        uint32 allocationDelay;
        OperatorSet[] allocatedSets;
        OperatorSet[] registeredSets;
        IStrategy[] allocatedStrategies;
        Allocation allocation;
        OperatorSet[] strategyAllocationsOperatorSets;
        Allocation[] strategyAllocations;
    }

    User staker;
    User operator;
    AllocateParams allocateParams;

    AVS avs;
    IStrategy[] strategies;
    uint[] initTokenBalances;
    uint[] initDepositShares;
    OperatorSet operatorSet;

    /// Shared setup:
    ///
    /// 1. Generate staker with assets, operator, and AVS
    /// 2. Staker deposits assets and delegates to operator
    /// 3. AVS creates an operator set containing the strategies held by the staker
    /// 4. Operator registers for the operator set
    /// 5. Operator allocates to the operator set
    /// 6. Roll blocks to complete allocation
    function _init() internal override {
        // 1. Create entities
        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator_NoAssets();
        (avs,) = _newRandomAVS();

        operator.setAllocationDelay(ALLOCATION_CONFIGURATION_DELAY);
        rollForward({blocks: ALLOCATION_CONFIGURATION_DELAY + 1});

        // 2. Staker deposits into EigenLayer
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, initDepositShares);

        // 3. Staker delegates to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // 4. AVS creates an operator set containing the strategies held by the staker
        operatorSet = avs.createOperatorSet(strategies);

        // 5. Operator registers for the operator set
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, allStrats);

        // 6. Operator allocates to the operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);

        // 7. Roll blocks to complete allocation
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
    }

    function testFuzz_query_upgrade_compare(uint24 r) public rand(r) {
        AllocationManagerState memory preUpgrade = _captureAllocationManagerState();
        _upgradeEigenLayerContracts();
        AllocationManagerState memory postUpgrade = _captureAllocationManagerState();
        _verifyStorageStateConsistency(preUpgrade, postUpgrade);
    }

    function _captureAllocationManagerState() internal view returns (AllocationManagerState memory state) {
        // Capture storage variables that should persist across upgrades
        state.isOperatorSet = allocationManager.isOperatorSet(operatorSet);
        state.operatorSetCount = allocationManager.getOperatorSetCount(address(avs));
        state.isMemberOfOperatorSet = allocationManager.isMemberOfOperatorSet(address(operator), operatorSet);
        state.memberCount = allocationManager.getMemberCount(operatorSet);
        state.strategiesInOperatorSet = allocationManager.getStrategiesInOperatorSet(operatorSet);
        state.isOperatorSlashable = allocationManager.isOperatorSlashable(address(operator), operatorSet);
        state.isOperatorRedistributable = allocationManager.isOperatorRedistributable(address(operator));

        // Only query strategy-related data if we have at least one strategy
        if (strategies.length > 0 && address(strategies[0]) != address(0)) {
            state.encumberedMagnitude = allocationManager.getEncumberedMagnitude(address(operator), strategies[0]);
            state.allocatableMagnitude = allocationManager.getAllocatableMagnitude(address(operator), strategies[0]);
            state.maxMagnitude = allocationManager.getMaxMagnitude(address(operator), strategies[0]);
        } else {
            state.encumberedMagnitude = 0;
            state.allocatableMagnitude = 0;
            state.maxMagnitude = 0;
        }
        (state.allocationDelayIsSet, state.allocationDelay) = allocationManager.getAllocationDelay(address(operator));
        state.allocatedSets = allocationManager.getAllocatedSets(address(operator));
        state.registeredSets = allocationManager.getRegisteredSets(address(operator));
        state.allocatedStrategies = allocationManager.getAllocatedStrategies(address(operator), operatorSet);
        // Only query allocation data if we have at least one strategy
        if (strategies.length > 0 && address(strategies[0]) != address(0)) {
            state.allocation = allocationManager.getAllocation(address(operator), operatorSet, strategies[0]);
            (state.strategyAllocationsOperatorSets, state.strategyAllocations) =
                allocationManager.getStrategyAllocations(address(operator), strategies[0]);
        } else {
            state.allocation = Allocation(0, 0, 0);
            state.strategyAllocationsOperatorSets = new OperatorSet[](0);
            state.strategyAllocations = new Allocation[](0);
        }
        state.redistributionRecipient = allocationManager.getRedistributionRecipient(operatorSet);
        state.isRedistributingOperatorSet = allocationManager.isRedistributingOperatorSet(operatorSet);
        state.slashCount = allocationManager.getSlashCount(operatorSet);
    }

    function _verifyStorageStateConsistency(AllocationManagerState memory preUpgrade, AllocationManagerState memory postUpgrade) internal {
        // Verify storage variables that should persist across upgrades
        assertEq(preUpgrade.isOperatorSet, postUpgrade.isOperatorSet, "isOperatorSet should be the same");
        assertEq(preUpgrade.operatorSetCount, postUpgrade.operatorSetCount, "operatorSetCount should be the same");
        assertEq(preUpgrade.isMemberOfOperatorSet, postUpgrade.isMemberOfOperatorSet, "isMemberOfOperatorSet should be the same");
        assertEq(preUpgrade.memberCount, postUpgrade.memberCount, "memberCount should be the same");
        assertEq(preUpgrade.isOperatorSlashable, postUpgrade.isOperatorSlashable, "isOperatorSlashable should be the same");
        assertEq(
            preUpgrade.isOperatorRedistributable, postUpgrade.isOperatorRedistributable, "isOperatorRedistributable should be the same"
        );
        assertEq(preUpgrade.encumberedMagnitude, postUpgrade.encumberedMagnitude, "encumberedMagnitude should be the same");
        assertEq(preUpgrade.allocatableMagnitude, postUpgrade.allocatableMagnitude, "allocatableMagnitude should be the same");
        assertEq(preUpgrade.maxMagnitude, postUpgrade.maxMagnitude, "maxMagnitude should be the same");
        assertEq(preUpgrade.allocationDelayIsSet, postUpgrade.allocationDelayIsSet, "allocationDelayIsSet should be the same");
        assertEq(preUpgrade.allocationDelay, postUpgrade.allocationDelay, "allocationDelay should be the same");
        assertEq(preUpgrade.redistributionRecipient, postUpgrade.redistributionRecipient, "redistributionRecipient should be the same");
        assertEq(
            preUpgrade.isRedistributingOperatorSet,
            postUpgrade.isRedistributingOperatorSet,
            "isRedistributingOperatorSet should be the same"
        );
        assertEq(preUpgrade.slashCount, postUpgrade.slashCount, "slashCount should be the same");

        // Verify array lengths
        assertEq(
            preUpgrade.strategiesInOperatorSet.length,
            postUpgrade.strategiesInOperatorSet.length,
            "strategiesInOperatorSet length should be the same"
        );
        assertEq(preUpgrade.allocatedSets.length, postUpgrade.allocatedSets.length, "allocatedSets length should be the same");
        assertEq(preUpgrade.registeredSets.length, postUpgrade.registeredSets.length, "registeredSets length should be the same");
        assertEq(
            preUpgrade.allocatedStrategies.length, postUpgrade.allocatedStrategies.length, "allocatedStrategies length should be the same"
        );
        assertEq(
            preUpgrade.strategyAllocationsOperatorSets.length,
            postUpgrade.strategyAllocationsOperatorSets.length,
            "strategyAllocationsOperatorSets length should be the same"
        );
        assertEq(
            preUpgrade.strategyAllocations.length, postUpgrade.strategyAllocations.length, "strategyAllocations length should be the same"
        );

        // Verify allocation struct
        assertEq(
            preUpgrade.allocation.currentMagnitude,
            postUpgrade.allocation.currentMagnitude,
            "allocation.currentMagnitude should be the same"
        );
        assertEq(preUpgrade.allocation.pendingDiff, postUpgrade.allocation.pendingDiff, "allocation.pendingDiff should be the same");
        assertEq(preUpgrade.allocation.effectBlock, postUpgrade.allocation.effectBlock, "allocation.effectBlock should be the same");

        // Verify array contents (if arrays have elements)
        if (preUpgrade.strategiesInOperatorSet.length > 0) {
            for (uint i = 0; i < preUpgrade.strategiesInOperatorSet.length; i++) {
                assertEq(
                    address(preUpgrade.strategiesInOperatorSet[i]),
                    address(postUpgrade.strategiesInOperatorSet[i]),
                    "strategiesInOperatorSet element should be the same"
                );
            }
        }

        if (preUpgrade.allocatedSets.length > 0) {
            for (uint i = 0; i < preUpgrade.allocatedSets.length; i++) {
                assertEq(preUpgrade.allocatedSets[i].avs, postUpgrade.allocatedSets[i].avs, "allocatedSets avs should be the same");
                assertEq(preUpgrade.allocatedSets[i].id, postUpgrade.allocatedSets[i].id, "allocatedSets id should be the same");
            }
        }

        if (preUpgrade.registeredSets.length > 0) {
            for (uint i = 0; i < preUpgrade.registeredSets.length; i++) {
                assertEq(preUpgrade.registeredSets[i].avs, postUpgrade.registeredSets[i].avs, "registeredSets avs should be the same");
                assertEq(preUpgrade.registeredSets[i].id, postUpgrade.registeredSets[i].id, "registeredSets id should be the same");
            }
        }

        if (preUpgrade.allocatedStrategies.length > 0) {
            for (uint i = 0; i < preUpgrade.allocatedStrategies.length; i++) {
                assertEq(
                    address(preUpgrade.allocatedStrategies[i]),
                    address(postUpgrade.allocatedStrategies[i]),
                    "allocatedStrategies element should be the same"
                );
            }
        }

        if (preUpgrade.strategyAllocations.length > 0) {
            for (uint i = 0; i < preUpgrade.strategyAllocations.length; i++) {
                assertEq(
                    preUpgrade.strategyAllocations[i].currentMagnitude,
                    postUpgrade.strategyAllocations[i].currentMagnitude,
                    "strategyAllocations currentMagnitude should be the same"
                );
                assertEq(
                    preUpgrade.strategyAllocations[i].pendingDiff,
                    postUpgrade.strategyAllocations[i].pendingDiff,
                    "strategyAllocations pendingDiff should be the same"
                );
                assertEq(
                    preUpgrade.strategyAllocations[i].effectBlock,
                    postUpgrade.strategyAllocations[i].effectBlock,
                    "strategyAllocations effectBlock should be the same"
                );
            }
        }

        console.log("All storage state consistency checks passed!".green());
    }
}
