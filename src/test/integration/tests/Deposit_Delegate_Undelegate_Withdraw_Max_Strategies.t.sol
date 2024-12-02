// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/users/User.t.sol";
import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_Deposit_Delegate_Undelegate_Withdraw_Max_Strategies is
    IntegrationCheckUtils,
    IDelegationManagerTypes
{
    function testFuzz_deposit_delegate_undelegate_withdraw_max_strategies(
        uint24 _random
    ) public {
        // When new Users are created, they will choose a random configuration from these params:
        _configRand({_randomSeed: _random, _assetTypes: HOLDS_MAX, _userTypes: DEFAULT});
        _upgradeEigenLayerContracts(); // Upgrade contracts if forkType is not local

        (User staker, IStrategy[] memory strategies, uint256[] memory balances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        (AVS avs,) = _newRandomAVS();

        OperatorSet memory operatorSet = avs.createOperatorSet(strategies);

        assertEq(strategies.length, 33, "sanity");

        /// 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, balances);

        // 2. Delegate to an operator
        staker.delegateTo(operator);

        // 3. Register for operator set
        operator.registerForOperatorSet(operatorSet);

        // 4. Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.undelegate();

        // 5. Complete withdrawal
        _rollBlocksForCompleteWithdrawals();
        staker.completeWithdrawalsAsTokens(withdrawals);
    }

    function testFuzz_deposit_delegate_allocate_slash_undelegate_withdraw_max_strategies(
        uint24 _random
    ) public {
        // When new Users are created, they will choose a random configuration from these params:
        _configRand({_randomSeed: _random, _assetTypes: HOLDS_MAX, _userTypes: DEFAULT});
        _upgradeEigenLayerContracts(); // Upgrade contracts if forkType is not local

        (User staker, IStrategy[] memory strategies, uint256[] memory balances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        (AVS avs,) = _newRandomAVS();

        OperatorSet memory operatorSet = avs.createOperatorSet(strategies);

        assertEq(strategies.length, 33, "sanity");

        /// 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, balances);

        // 2. Delegate to an operator
        staker.delegateTo(operator);

        // 3. Register for operator set
        operator.registerForOperatorSet(operatorSet);

        // 4. Allocate to operator set.
        operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
        cheats.roll(block.number + allocationManager.ALLOCATION_CONFIGURATION_DELAY());

        // 5. Slash operator
        avs.slashOperator(operator, operatorSet.id, 0.5 ether);

        // 6. Undelegate from an operator
        Withdrawal[] memory withdrawals = staker.undelegate();

        // 7. Complete withdrawal
        _rollBlocksForCompleteWithdrawals();
        staker.completeWithdrawalsAsTokens(withdrawals);
    }
}
