// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

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
    function testFuzz_slash_upgrade_clearGrandfatheredSlash(uint24 _random) public rand(_random) {
        IERC20 token = IStrategy(strategies[0]).underlyingToken();
        address redistributionRecipient = allocationManager.getRedistributionRecipient(operatorSet);
        uint recipientBalanceBefore = token.balanceOf(redistributionRecipient);
        uint pendingSharesBefore = strategyManager.getBurnOrRedistributableShares(operatorSet, slashId, strategies[0]);
        IStrategy[] memory clearedStrategies = new IStrategy[](1);
        clearedStrategies[0] = strategies[0];
        uint[] memory expectedRedistributions = _calculateExpectedTokens(clearedStrategies, sharesSlashed);

        // 6. Upgrade contracts
        _upgradeEigenLayerContracts();

        // 7. Grandfathered slashes should have a zero resolution block and remain immediately clearable
        assertEq(strategyManager.getSlashResolutionBlock(operatorSet, slashId), 0, "grandfathered slash should have zero resolution block");
        assertEq(strategyManager.getBurnOrRedistributableCount(operatorSet, slashId), 1, "grandfathered slash should still be pending");
        assertEq(
            strategyManager.getBurnOrRedistributableShares(operatorSet, slashId, strategies[0]),
            pendingSharesBefore,
            "grandfathered slash shares should persist across upgrade"
        );

        // 8. Clear immediately after the upgrade without rolling past the new delay
        avs.clearBurnOrRedistributableSharesByStrategy(operatorSet, slashId, strategies[0]);

        // 9. Clearing should pay the recipient and remove the pending slash bookkeeping
        assertEq(strategyManager.getBurnOrRedistributableCount(operatorSet, slashId), 0, "grandfathered slash should be cleared");
        assertEq(strategyManager.getPendingSlashIds(operatorSet).length, 0, "pending slash ids should be empty");
        assertEq(strategyManager.getPendingOperatorSets().length, 0, "pending operator sets should be empty");
        assertEq(strategyManager.getSlashResolutionBlock(operatorSet, slashId), 0, "resolution block should remain zero after clear");
        assertEq(
            token.balanceOf(redistributionRecipient),
            recipientBalanceBefore + expectedRedistributions[0],
            "recipient should receive the grandfathered slash funds"
        );
    }
}
