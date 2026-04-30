// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "src/contracts/interfaces/IStrategyManager.sol";
import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_SlashResolutionDelay_Base is IntegrationCheckUtils {
    AVS avs;
    OperatorSet operatorSet;

    User operator;
    User staker;

    IStrategy[] strategies;
    uint[] initTokenBalances;
    uint[] initDepositShares;

    AllocateParams allocateParams;
    address redistributionRecipient = address(0xBEEF);

    function _init() internal virtual override {
        _configAssetTypes(HOLDS_LST);

        (staker, strategies, initTokenBalances) = _newRandomStaker();
        (operator,,) = _newRandomOperator();
        (avs,) = _newRandomAVS();

        // 1. Deposit into strategies
        staker.depositIntoEigenlayer(strategies, initTokenBalances);
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
        check_Deposit_State(staker, strategies, initDepositShares);

        // 2. Delegate staker to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // 3. Create a redistributing operator set
        operatorSet = avs.createRedistributingOperatorSet(strategies, redistributionRecipient);

        // 4. Register operator to the operator set
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, allStrats);

        // 5. Allocate all available magnitude to the operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator, allocateParams);

        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
    }

    function _rollPastSlashResolutionDelay(uint32 resolutionBlock) internal {
        if (block.number <= resolutionBlock) cheats.roll(uint(resolutionBlock) + 1);
    }
}

contract Integration_SlashResolutionDelay is Integration_SlashResolutionDelay_Base {
    function testFuzz_slash_clearBeforeDelay_wait_clearAfterDelay(uint24 _random) public rand(_random) {
        // 6. Slash the operator for a single strategy
        uint32 slashBlock = uint32(block.number);
        SlashingParams memory slashParams = _genSlashing_SingleStrategy(operator, operatorSet, strategies[0]);
        (uint slashId, uint[] memory sharesSlashed) = avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams, slashId);

        uint32 resolutionBlock = strategyManager.getSlashResolutionBlock(operatorSet, slashId);
        assertEq(
            resolutionBlock,
            slashBlock + strategyManager.SLASH_RESOLUTION_DELAY_BLOCKS(),
            "resolution block should be slash block plus delay"
        );
        assertEq(strategyManager.getBurnOrRedistributableCount(operatorSet, slashId), 1, "slash should have one pending strategy");
        assertEq(
            strategyManager.getBurnOrRedistributableShares(operatorSet, slashId, strategies[0]), sharesSlashed[0], "pending shares mismatch"
        );

        // 7. Clearing before the resolution delay elapses should revert
        IERC20 token = IStrategy(strategies[0]).underlyingToken();
        uint recipientBalanceBefore = token.balanceOf(redistributionRecipient);
        uint pendingSharesBefore = strategyManager.getBurnOrRedistributableShares(operatorSet, slashId, strategies[0]);
        IStrategy[] memory clearedStrategies = new IStrategy[](1);
        clearedStrategies[0] = strategies[0];
        uint[] memory expectedRedistributions = _calculateExpectedTokens(clearedStrategies, sharesSlashed);

        cheats.expectRevert(IStrategyManagerErrors.SlashResolutionDelayNotElapsed.selector);
        avs.clearBurnOrRedistributableSharesByStrategy(operatorSet, slashId, strategies[0]);

        assertEq(strategyManager.getSlashResolutionBlock(operatorSet, slashId), resolutionBlock, "resolution block should be unchanged");
        assertEq(strategyManager.getBurnOrRedistributableCount(operatorSet, slashId), 1, "slash should still be pending");
        assertEq(
            strategyManager.getBurnOrRedistributableShares(operatorSet, slashId, strategies[0]),
            pendingSharesBefore,
            "pending shares should be unchanged before delay"
        );
        assertEq(token.balanceOf(redistributionRecipient), recipientBalanceBefore, "recipient should not be paid before delay");

        // 8. Roll past the delay and clear the slash
        _rollPastSlashResolutionDelay(resolutionBlock);
        avs.clearBurnOrRedistributableSharesByStrategy(operatorSet, slashId, strategies[0]);

        // 9. Cleared state should remove the slash and pay the redistribution recipient
        assertEq(strategyManager.getBurnOrRedistributableCount(operatorSet, slashId), 0, "slash should be fully cleared");
        assertEq(strategyManager.getPendingSlashIds(operatorSet).length, 0, "pending slash ids should be empty");
        assertEq(strategyManager.getPendingOperatorSets().length, 0, "pending operator sets should be empty");
        assertEq(strategyManager.getSlashResolutionBlock(operatorSet, slashId), 0, "resolution block should be deleted");
        assertEq(
            token.balanceOf(redistributionRecipient),
            recipientBalanceBefore + expectedRedistributions[0],
            "recipient should receive slashed funds"
        );
    }

    function testFuzz_slash_slash_clearFirst_wait_clearSecond(uint24 _random) public rand(_random) {
        // 6. Slash the same operator set twice for the same strategy
        IERC20 token = IStrategy(strategies[0]).underlyingToken();
        uint recipientBalanceBefore = token.balanceOf(redistributionRecipient);

        SlashingParams memory firstSlashParams = _genSlashing_SingleStrategy(operator, operatorSet, strategies[0]);
        firstSlashParams.wadsToSlash[0] = 0.1e18;
        (uint slashId1, uint[] memory sharesSlashed1) = avs.slashOperator(firstSlashParams);
        check_Base_Slashing_State(operator, allocateParams, firstSlashParams, slashId1);
        IStrategy[] memory clearedStrategies = new IStrategy[](1);
        clearedStrategies[0] = strategies[0];
        uint[] memory expectedRedistributions1 = _calculateExpectedTokens(clearedStrategies, sharesSlashed1);

        uint32 resolutionBlock1 = strategyManager.getSlashResolutionBlock(operatorSet, slashId1);

        cheats.roll(block.number + 17);

        SlashingParams memory secondSlashParams = _genSlashing_SingleStrategy(operator, operatorSet, strategies[0]);
        secondSlashParams.wadsToSlash[0] = 0.1e18;
        (uint slashId2, uint[] memory sharesSlashed2) = avs.slashOperator(secondSlashParams);
        uint[] memory expectedRedistributions2 = _calculateExpectedTokens(clearedStrategies, sharesSlashed2);

        uint32 resolutionBlock2 = strategyManager.getSlashResolutionBlock(operatorSet, slashId2);
        assertGt(resolutionBlock2, resolutionBlock1, "second slash should resolve later");
        assertEq(strategyManager.getPendingSlashIds(operatorSet).length, 2, "both slash ids should be pending");

        // 7. Clear the first slash after its delay; the second must remain pending
        _rollPastSlashResolutionDelay(resolutionBlock1);
        avs.clearBurnOrRedistributableSharesByStrategy(operatorSet, slashId1, strategies[0]);

        assertEq(strategyManager.getSlashResolutionBlock(operatorSet, slashId1), 0, "first resolution block should be deleted");
        assertEq(strategyManager.getSlashResolutionBlock(operatorSet, slashId2), resolutionBlock2, "second resolution block should remain");
        assertEq(strategyManager.getBurnOrRedistributableCount(operatorSet, slashId2), 1, "second slash should remain pending");
        assertEq(strategyManager.getPendingSlashIds(operatorSet).length, 1, "one slash id should remain pending");
        assertEq(
            token.balanceOf(redistributionRecipient), recipientBalanceBefore + expectedRedistributions1[0], "first slash payout incorrect"
        );

        // 8. Clearing the second slash before its own delay should still revert
        cheats.expectRevert(IStrategyManagerErrors.SlashResolutionDelayNotElapsed.selector);
        avs.clearBurnOrRedistributableSharesByStrategy(operatorSet, slashId2, strategies[0]);

        // 9. Roll past the second delay and clear it
        _rollPastSlashResolutionDelay(resolutionBlock2);
        avs.clearBurnOrRedistributableSharesByStrategy(operatorSet, slashId2, strategies[0]);

        assertEq(strategyManager.getBurnOrRedistributableCount(operatorSet, slashId2), 0, "second slash should be fully cleared");
        assertEq(strategyManager.getPendingSlashIds(operatorSet).length, 0, "pending slash ids should be empty");
        assertEq(strategyManager.getPendingOperatorSets().length, 0, "pending operator sets should be empty");
        assertEq(strategyManager.getSlashResolutionBlock(operatorSet, slashId2), 0, "second resolution block should be deleted");
        assertEq(
            token.balanceOf(redistributionRecipient),
            recipientBalanceBefore + expectedRedistributions1[0] + expectedRedistributions2[0],
            "total payout incorrect"
        );
    }

    function testFuzz_slash_wait_batchClearAfterDelay(uint24 _random) public rand(_random) {
        // 6. Slash the operator for all strategies in the operator set
        SlashingParams memory slashParams = _genSlashing_Half(operator, operatorSet);
        (uint slashId, uint[] memory sharesSlashed) = avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams, slashId);

        uint32 resolutionBlock = strategyManager.getSlashResolutionBlock(operatorSet, slashId);
        assertEq(strategyManager.getBurnOrRedistributableCount(operatorSet, slashId), slashParams.strategies.length, "pending count mismatch");

        IERC20[] memory tokens = new IERC20[](slashParams.strategies.length);
        uint[] memory recipientBalancesBefore = new uint[](slashParams.strategies.length);
        uint[] memory expectedRedistributions = _calculateExpectedTokens(slashParams.strategies, sharesSlashed);

        for (uint i = 0; i < slashParams.strategies.length; ++i) {
            tokens[i] = IStrategy(slashParams.strategies[i]).underlyingToken();
            recipientBalancesBefore[i] = tokens[i].balanceOf(redistributionRecipient);
            assertEq(
                strategyManager.getBurnOrRedistributableShares(operatorSet, slashId, slashParams.strategies[i]),
                sharesSlashed[i],
                "pending shares mismatch"
            );
        }

        // 7. Roll past the delay and batch-clear every strategy in the slash
        _rollPastSlashResolutionDelay(resolutionBlock);
        avs.clearBurnOrRedistributableShares(operatorSet, slashId);

        // 8. Cleared state should remove the slash and pay the redistribution recipient for every strategy
        assertEq(strategyManager.getBurnOrRedistributableCount(operatorSet, slashId), 0, "slash should be fully cleared");
        assertEq(strategyManager.getPendingSlashIds(operatorSet).length, 0, "pending slash ids should be empty");
        assertEq(strategyManager.getPendingOperatorSets().length, 0, "pending operator sets should be empty");
        assertEq(strategyManager.getSlashResolutionBlock(operatorSet, slashId), 0, "resolution block should be deleted");

        for (uint i = 0; i < slashParams.strategies.length; ++i) {
            assertEq(
                tokens[i].balanceOf(redistributionRecipient),
                recipientBalancesBefore[i] + expectedRedistributions[i],
                "recipient balance mismatch"
            );
        }
    }
}
