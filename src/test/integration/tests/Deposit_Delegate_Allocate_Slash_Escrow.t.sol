// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";
import {console} from "forge-std/console.sol";

contract Integration_Deposit_Delegate_Allocate_Slash_Escrow is IntegrationCheckUtils {
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
    uint[] slashShares;

    address redistributionRecipient;

    function _init() internal virtual override {
        _configAssetTypes(HOLDS_LST);

        (staker, strategies, initTokenBalances) = _newRandomStaker();
        operator = _newRandomOperator_NoAssets();
        (avs,) = _newRandomAVS();
        redistributionRecipient = cheats.randomAddress();

        operatorSet = avs.createRedistributingOperatorSet(strategies, redistributionRecipient);
        tokens = _getUnderlyingTokens(strategies);

        // 1) Register operator for operator set.
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, strategies);

        // 2) Operator allocates to operator set.
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet);
        operator.modifyAllocations(allocateParams);
        check_Base_IncrAlloc_State(operator, allocateParams);

        // 3) Roll forward to complete allocation.
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);
    }

    function testFuzz_fullSlash_EscrowTiming_EscrowDelayNotElapsed(uint24 _random) public rand(_random) {
        // 4) Operator is full slashed.
        slashParams = _genSlashing_Full(operator, operatorSet);
        (slashId, slashShares) = avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams, slashId); // TODO: add state, balance checks

        // Roll forward to just before the escrow delay.
        cheats.roll(block.number + INITIAL_GLOBAL_DELAY_BLOCKS);

        // Attempt to release escrow before delay has elapsed, expect revert.
        cheats.prank(redistributionRecipient);
        cheats.expectRevert(ISlashEscrowFactoryErrors.EscrowDelayNotElapsed.selector);
        slashEscrowFactory.releaseSlashEscrow({operatorSet: operatorSet, slashId: 1});
    }

    function testFuzz_fullSlash_EscrowTiming_EscrowElapsed(uint24 _random) public rand(_random) {
        // 4) Operator is full slashed.
        slashParams = _genSlashing_Full(operator, operatorSet);
        (slashId, slashShares) = avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams, slashId); // TODO: add state, balance checks

        // Roll forward to just before the escrow delay.
        cheats.roll(block.number + INITIAL_GLOBAL_DELAY_BLOCKS + 1);

        // Release escrow, expect success.
        cheats.prank(redistributionRecipient);
        slashEscrowFactory.releaseSlashEscrow({operatorSet: operatorSet, slashId: 1});
    }
}
