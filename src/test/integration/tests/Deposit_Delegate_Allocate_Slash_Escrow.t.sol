// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";
import {console} from "forge-std/console.sol";

contract Integration_Deposit_Delegate_Allocate_Slash_Escrow is IntegrationCheckUtils {
    using ArrayLib for *;

    AVS avs;
    User staker;
    User operator;
    OperatorSet operatorSet;
    AllocateParams allocateParams;
    SlashingParams slashParams;
    uint slashId;
    IStrategy[] strategies;
    IERC20[] tokens;
    uint[] initTokenBalances;
    uint[] initDepositShares;
    address payable redistributionRecipient;

    function _init() internal virtual override {
        _configAssetTypes(HOLDS_LST);

        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator_NoAssets();
        (avs,) = _newRandomAVS();
        redistributionRecipient = payable(cheats.randomAddress());
        cheats.label(redistributionRecipient, "redistributionRecipient");

        operatorSet = avs.createRedistributingOperatorSet(strategies, redistributionRecipient);
        tokens = _getUnderlyingTokens(strategies);

        // 1) Register operator for operator set.
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, strategies);

        // 2) Deposit Into Strategies
        initDepositShares = _calculateExpectedShares(strategies, initTokenBalances);
        staker.depositIntoEigenlayer(strategies, initTokenBalances);

        // 3) Delegate to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // 4) Operator allocates to operator set.
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_Base_IncrAlloc_State(operator, allocateParams);

        // 5) Roll forward to complete allocation.
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
    }

    function testFuzz_fullSlash_EscrowTiming_RightBeforeFails(uint24 _random) public rand(_random) {
        // 6) Operator is full slashed.
        slashParams = _genSlashing_Full(operator, operatorSet);
        (slashId,) = avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams, slashId);

        // Roll forward to just before the escrow delay.
        cheats.roll(block.number + INITIAL_GLOBAL_DELAY_BLOCKS);

        // Attempt to release escrow before delay has elapsed, expect revert.
        cheats.prank(redistributionRecipient);
        cheats.expectRevert(ISlashEscrowFactoryErrors.EscrowDelayNotElapsed.selector);
        slashEscrowFactory.releaseSlashEscrow({operatorSet: operatorSet, slashId: 1});
    }

    function testFuzz_fullSlash_EscrowTiming_RightAfterPasses(uint24 _random) public rand(_random) {
        // 6) Operator is full slashed.
        slashParams = _genSlashing_Full(operator, operatorSet);
        (slashId,) = avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams, slashId);

        // Roll forward to just after the escrow delay.
        cheats.roll(block.number + INITIAL_GLOBAL_DELAY_BLOCKS + 1);

        // 7) Release escrow, expect success.
        cheats.prank(redistributionRecipient);
        slashEscrowFactory.releaseSlashEscrow({operatorSet: operatorSet, slashId: 1});
    }

    function testFuzz_fullSlash_releaseAll(uint24 _random) public rand(_random) {
        // 6) Operator is full slashed.
        slashParams = _genSlashing_Full(operator, operatorSet);
        (slashId,) = avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams, slashId);

        // Roll forward to just before the escrow delay.
        cheats.roll(block.number + INITIAL_GLOBAL_DELAY_BLOCKS + 1);

        // 7) Release escrow, expect success.
        cheats.prank(redistributionRecipient);
        slashEscrowFactory.releaseSlashEscrow({operatorSet: operatorSet, slashId: 1});
        check_releaseSlashEscrow_State_NoneRemaining(operatorSet, slashId, strategies, initTokenBalances, redistributionRecipient);
    }

    function testFuzz_fullSlash_clearAll_releaseAll(uint24 _random) public rand(_random) {
        // 6) Operator is full slashed.
        slashParams = _genSlashing_Full(operator, operatorSet);
        (slashId,) = avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams, slashId);

        // Roll forward to just before the escrow delay.
        cheats.roll(block.number + INITIAL_GLOBAL_DELAY_BLOCKS + 1);

        // 7) Clear burnable shares (transfers tokens to escrow).
        avs.clearBurnOrRedistributableShares(operatorSet, slashId);
        assert_HasUnderlyingTokenBalances(
            User(payable(address(slashEscrowFactory.getSlashEscrow(operatorSet, slashId)))),
            strategies,
            initTokenBalances,
            "slash escrow should have underlying token balances"
        );

        // 8) Release escrow, expect success.
        cheats.prank(redistributionRecipient);
        slashEscrowFactory.releaseSlashEscrow({operatorSet: operatorSet, slashId: 1});
        check_releaseSlashEscrow_State_NoneRemaining(operatorSet, slashId, strategies, initTokenBalances, redistributionRecipient);
    }

    function testFuzz_fullSlash_clearAll_releaseByStrategy(uint24 _random) public rand(_random) {
        // 6) Operator is full slashed.
        slashParams = _genSlashing_Full(operator, operatorSet);
        (slashId,) = avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams, slashId);

        // Roll forward to just before the escrow delay.
        cheats.roll(block.number + INITIAL_GLOBAL_DELAY_BLOCKS + 1);

        avs.clearBurnOrRedistributableShares(operatorSet, slashId);
        assert_HasUnderlyingTokenBalances(
            User(payable(address(slashEscrowFactory.getSlashEscrow(operatorSet, slashId)))),
            strategies,
            initTokenBalances,
            "slash escrow should have underlying token balances"
        );

        // Reorder strategies and initTokenBalances
        for (uint i = 0; i < strategies.length; i++) {
            uint randomIndex = cheats.randomUint(0, strategies.length - 1);
            IStrategy tempStrategy = strategies[i];
            strategies[i] = strategies[randomIndex];
            strategies[randomIndex] = tempStrategy;

            uint tempBalance = initTokenBalances[i];
            initTokenBalances[i] = initTokenBalances[randomIndex];
            initTokenBalances[randomIndex] = tempBalance;
        }

        // 8) Release escrow, expect success.
        for (uint i = 0; i < strategies.length; i++) {
            cheats.prank(redistributionRecipient);
            slashEscrowFactory.releaseSlashEscrowByStrategy({operatorSet: operatorSet, slashId: 1, strategy: strategies[i]});
        }

        check_releaseSlashEscrow_State_NoneRemaining(operatorSet, slashId, strategies, initTokenBalances, redistributionRecipient);
    }

    function testFuzz_fullSlash_clearByStrategy_releaseAll(uint24 _random) public rand(_random) {
        // 6) Operator is full slashed.
        slashParams = _genSlashing_Full(operator, operatorSet);
        (slashId,) = avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams, slashId);

        // Roll forward to just before the escrow delay.
        cheats.roll(block.number + INITIAL_GLOBAL_DELAY_BLOCKS + 1);

        // Reorder strategies and initTokenBalances
        for (uint i = 0; i < strategies.length; i++) {
            uint randomIndex = cheats.randomUint(0, strategies.length - 1);
            IStrategy tempStrategy = strategies[i];
            strategies[i] = strategies[randomIndex];
            strategies[randomIndex] = tempStrategy;

            uint tempBalance = initTokenBalances[i];
            initTokenBalances[i] = initTokenBalances[randomIndex];
            initTokenBalances[randomIndex] = tempBalance;
        }

        // 7) Clear burnable shares (transfers tokens to escrow).
        for (uint i = 0; i < strategies.length; i++) {
            avs.clearBurnOrRedistributableSharesByStrategy(operatorSet, slashId, strategies[i]);
            assert_HasUnderlyingTokenBalance(
                User(payable(address(slashEscrowFactory.getSlashEscrow(operatorSet, slashId)))),
                strategies[i],
                initTokenBalances[i],
                "slash escrow should have underlying token balance"
            );
        }

        // 8) Release escrow, expect success.
        cheats.prank(redistributionRecipient);
        slashEscrowFactory.releaseSlashEscrow({operatorSet: operatorSet, slashId: 1});
        check_releaseSlashEscrow_State_NoneRemaining(operatorSet, slashId, strategies, initTokenBalances, redistributionRecipient);
    }

    function testFuzz_fullSlash_clearByStrategy_releaseByStrategy(uint24 _random) public rand(_random) {
        // 6) Operator is full slashed.
        slashParams = _genSlashing_Full(operator, operatorSet);
        (slashId,) = avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams, slashId);

        // Roll forward to just before the escrow delay.
        cheats.roll(block.number + INITIAL_GLOBAL_DELAY_BLOCKS + 1);

        // Reorder strategies and initTokenBalances
        for (uint i = 0; i < strategies.length; i++) {
            uint randomIndex = cheats.randomUint(0, strategies.length - 1);
            IStrategy tempStrategy = strategies[i];
            strategies[i] = strategies[randomIndex];
            strategies[randomIndex] = tempStrategy;

            uint tempBalance = initTokenBalances[i];
            initTokenBalances[i] = initTokenBalances[randomIndex];
            initTokenBalances[randomIndex] = tempBalance;
        }

        // 7) Clear burnable shares (transfers tokens to escrow).
        for (uint i = 0; i < strategies.length; i++) {
            avs.clearBurnOrRedistributableSharesByStrategy(operatorSet, slashId, strategies[i]);
            assert_HasUnderlyingTokenBalance(
                User(payable(address(slashEscrowFactory.getSlashEscrow(operatorSet, slashId)))),
                strategies[i],
                initTokenBalances[i],
                "slash escrow should have underlying token balance"
            );
        }

        // 8) Release escrow, expect success.
        for (uint i = 0; i < strategies.length; i++) {
            cheats.prank(redistributionRecipient);
            slashEscrowFactory.releaseSlashEscrowByStrategy({operatorSet: operatorSet, slashId: 1, strategy: strategies[i]});
        }

        check_releaseSlashEscrow_State_NoneRemaining(operatorSet, slashId, strategies, initTokenBalances, redistributionRecipient);
    }
}
