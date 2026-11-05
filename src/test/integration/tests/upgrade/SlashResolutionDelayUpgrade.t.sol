// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "src/contracts/interfaces/IStrategyManager.sol";
import "src/test/integration/UpgradeTest.t.sol";

contract Integration_Upgrade_SlashResolutionDelay_Base is UpgradeTest {
    User operator;
    User staker;
    AVS avs;
    OperatorSet operatorSet;

    IStrategy[] strategies;
    uint[] initTokenBalances;

    AllocateParams allocateParams;
    SlashingParams slashParams;
    uint slashId;
    uint[] sharesSlashed;

    function _init() internal override {
        _configAssetTypes(HOLDS_LST);

        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator_NoAssets();
        (avs,) = _newRandomAVS_v1CreateSet();

        operator.setAllocationDelay(ALLOCATION_CONFIGURATION_DELAY);
        rollForward({blocks: ALLOCATION_CONFIGURATION_DELAY + 1});

        // 1. Deposit into strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);

        // 2. Delegate to an operator
        staker.delegateTo(operator);

        // 3. Create an operator set using the pre-upgrade path and register the operator
        operatorSet = avs.createOperatorSet_v1(strategies);
        operator.registerForOperatorSet(operatorSet);

        // 4. Allocate to the operator set and roll until slashable
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // 5. Create a pending slash before the upgrade
        slashParams = _genSlashing_SingleStrategy(operator, operatorSet, strategies[0]);
        slashParams.wadsToSlash[0] = 0.1e18;
        (slashId, sharesSlashed) = avs.slashOperator(slashParams);
    }
}

contract Integration_Upgrade_SlashResolutionDelay is Integration_Upgrade_SlashResolutionDelay_Base {
    function testFuzz_slash_upgrade_pendingSlashPreserved_delayConfigurable(uint24 _random) public rand(_random) {
        IERC20 token = IStrategy(strategies[0]).underlyingToken();
        address redistributionRecipient = allocationManager.getRedistributionRecipient(operatorSet);
        uint recipientBalanceBefore = token.balanceOf(redistributionRecipient);
        uint pendingSharesBefore = strategyManager.getBurnOrRedistributableShares(operatorSet, slashId, strategies[0]);
        uint32 resolutionBlockBefore = _getSlashResolutionBlockOrZero(operatorSet, slashId);
        IStrategy[] memory clearedStrategies = new IStrategy[](1);
        clearedStrategies[0] = strategies[0];
        uint[] memory expectedRedistributions = _calculateExpectedTokens(clearedStrategies, sharesSlashed);

        // 6. Upgrade contracts and migrate the legacy operator set's slasher so the AVS can slash again
        _upgradeEigenLayerContracts();
        OperatorSet[] memory operatorSets = new OperatorSet[](1);
        operatorSets[0] = operatorSet;
        allocationManager.migrateSlashers(operatorSets);

        // 7. The pre-upgrade slash is untouched (a pre-v1.13.0 slash stays grandfathered at 0; a v1.13.0 slash keeps
        //    its resolution block); the delay is now a storage value defaulting to 0 with no setter
        assertEq(
            strategyManager.getSlashResolutionBlock(operatorSet, slashId),
            resolutionBlockBefore,
            "pending slash resolution block changed by upgrade"
        );
        assertEq(strategyManager.getBurnOrRedistributableCount(operatorSet, slashId), 1, "pending slash should persist across upgrade");
        assertEq(
            strategyManager.getBurnOrRedistributableShares(operatorSet, slashId, strategies[0]),
            pendingSharesBefore,
            "pending slash shares should persist across upgrade"
        );
        assertEq(strategyManager.SLASH_RESOLUTION_DELAY_BLOCKS(), 0, "delay should default to zero after upgrade");
        assertEq(strategyManager.slashResolutionDelaySetter(), address(0), "setter should be unset after upgrade");

        // 8. The pre-upgrade slash still honors its original delay: reverts before, clears at the resolution block
        if (block.number < resolutionBlockBefore) {
            cheats.expectRevert(IStrategyManagerErrors.SlashResolutionDelayNotElapsed.selector);
            avs.clearBurnOrRedistributableSharesByStrategy(operatorSet, slashId, strategies[0]);
            cheats.roll(resolutionBlockBefore);
        }
        avs.clearBurnOrRedistributableSharesByStrategy(operatorSet, slashId, strategies[0]);
        assertEq(strategyManager.getBurnOrRedistributableCount(operatorSet, slashId), 0, "pre-upgrade slash should be cleared");
        assertEq(strategyManager.getPendingSlashIds(operatorSet).length, 0, "no pending slash ids after clearing");
        assertEq(strategyManager.getPendingOperatorSets().length, 0, "no pending operator sets after clearing");
        assertEq(strategyManager.getSlashResolutionBlock(operatorSet, slashId), 0, "resolution block should be deleted");
        assertEq(
            token.balanceOf(redistributionRecipient),
            recipientBalanceBefore + expectedRedistributions[0],
            "recipient should receive the pre-upgrade slash funds"
        );

        // 9. With the delay at zero, a new slash is clearable in the block it was recorded
        SlashingParams memory secondSlashParams = _genSlashing_SingleStrategy(operator, operatorSet, strategies[0]);
        secondSlashParams.wadsToSlash[0] = 0.1e18;
        (uint slashId2,) = avs.slashOperator(secondSlashParams);
        assertEq(
            strategyManager.getSlashResolutionBlock(operatorSet, slashId2),
            uint32(block.number),
            "zero-delay slash should resolve at the slash block"
        );
        avs.clearBurnOrRedistributableSharesByStrategy(operatorSet, slashId2, strategies[0]);
        assertEq(strategyManager.getBurnOrRedistributableCount(operatorSet, slashId2), 0, "zero-delay slash should clear immediately");

        // 10. The owner names a setter, the setter restores the delay, and the next slash is delayed again
        cheats.prank(executorMultisig);
        strategyManager.setSlashResolutionDelaySetter(executorMultisig);
        cheats.prank(executorMultisig);
        strategyManager.setSlashResolutionDelayBlocks(DEFAULT_SLASH_RESOLUTION_DELAY_BLOCKS);

        SlashingParams memory thirdSlashParams = _genSlashing_SingleStrategy(operator, operatorSet, strategies[0]);
        thirdSlashParams.wadsToSlash[0] = 0.1e18;
        (uint slashId3,) = avs.slashOperator(thirdSlashParams);
        assertEq(
            strategyManager.getSlashResolutionBlock(operatorSet, slashId3),
            uint32(block.number) + DEFAULT_SLASH_RESOLUTION_DELAY_BLOCKS,
            "delayed slash resolution block incorrect"
        );
        cheats.expectRevert(IStrategyManagerErrors.SlashResolutionDelayNotElapsed.selector);
        avs.clearBurnOrRedistributableSharesByStrategy(operatorSet, slashId3, strategies[0]);
    }

    /// @dev A pre-v1.13.0 StrategyManager has no `getSlashResolutionBlock`; treat that as 0 (grandfathered).
    function _getSlashResolutionBlockOrZero(OperatorSet memory operatorSet, uint slashId) internal view returns (uint32) {
        (bool ok, bytes memory data) =
            address(strategyManager).staticcall(abi.encodeCall(IStrategyManager.getSlashResolutionBlock, (operatorSet, slashId)));
        return ok ? abi.decode(data, (uint32)) : 0;
    }
}
