// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

contract Integration_Deposit_Delegate_Allocate is IntegrationCheckUtils, IDelegationManagerTypes {
    // function testFuzz_deposit_delegate_allocate(
    //     uint24 _random
    // ) public {
    //     // Configure the random parameters for the test
    //     _configRand({
    //         _randomSeed: _random,
    //         _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
    //         _userTypes: DEFAULT | ALT_METHODS
    //     });

    //     // Create a staker and an operator with a nonzero balance and corresponding strategies
    //     (AVS avs, OperatorSet[] memory operatorSets) = _newRandomAVS();
    //     (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
    //     (User operator,,) = _newRandomOperator();

    //     _upgradeEigenLayerContracts();

    //     // 1. Delegate to operator
    //     staker.delegateTo(operator);
    //     check_Delegation_State(staker, operator, strategies, new uint256[](strategies.length)); // Initial shares are zero

    //     // 2. Deposit into strategies
    //     staker.depositIntoEigenlayer(strategies, tokenBalances);
    //     uint256[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

    //     // Check that the deposit increased operator shares the staker is delegated to
    //     check_Deposit_State(staker, strategies, shares);
    //     assert_Snap_Added_OperatorShares(operator, strategies, shares, "operator should have received shares");

    //     operator.registerForOperatorSets(operatorSets);

    //     for (uint256 i; i < operatorSets.length; ++i) {
    //         uint256 len = allocationManager.getStrategiesInOperatorSet(operatorSets[i]).length;
    //         operator.modifyAllocations(
    //             operatorSets[i], _randMagnitudes({sum: 1 ether / uint64(operatorSets.length), len: len})
    //         );
    //         allocationManager.getAllocation(operatorSets[i]);
    //     }

    //     // TODO: slashing checks
    // }

    // FIXME: Current fails with `AmountMustBeMultipleOfGwei` since slashed beaconchain deposits
    // are not garanteed to be multiple of gwei after being slashed.
    function testFuzz_deposit_delegate_allocate_slash_undelegate_completeAsTokens(
        uint24 _random
    ) public {
        _configRand({_randomSeed: _random, _assetTypes: HOLDS_MAX, _userTypes: DEFAULT});
        _upgradeEigenLayerContracts(); // Upgrade contracts if forkType is not local

        (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        (AVS avs,) = _newRandomAVS();

        assertEq(strategies.length, 33, "sanity");

        uint256[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        assert_HasNoDelegatableShares(staker, "staker should not have delegatable shares before depositing");
        assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");

        /// 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // 2. Delegate to an operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, shares);

        // Create an operator set and register an operator.
        OperatorSet memory operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);

        // 3. Allocate to operator set.
        AllocateParams memory allocateParams =
            operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));

        assert_Snap_Allocations_Updated(
            operator,
            operatorSet,
            strategies,
            allocateParams.newMagnitudes,
            false,
            "operator allocations should be updated before delay"
        );

        _rollBlocksForCompleteAllocation();

        assert_Snap_Allocations_Updated(
            operator,
            operatorSet,
            strategies,
            allocateParams.newMagnitudes,
            true,
            "operator allocations should be updated after delay"
        );

        // TODO: check_Allocation_State()

        // 4. Slash operator
        SlashingParams memory slashingParams = avs.slashOperator(operator, operatorSet.id, _randWadToSlash());
        // TODO: check_Slash_State()

        // 5. Undelegate from an operator
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        // FIXME: check_Undelegate_State(staker, operator, withdrawals, withdrawalRoots, strategies, shares);

        // 6. Complete withdrawal
        _rollBlocksForCompleteWithdrawals();
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint256[] memory expectedTokens =
                _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            // FIXME: check_Withdrawal_AsTokens_State(staker, operator, withdrawals[i], withdrawals[i].strategies, withdrawals[i].scaledShares, tokens, expectedTokens);
        }

        // Check Final State
        assert_HasNoDelegatableShares(staker, "staker should have withdrawn all shares");
        assert_HasUnderlyingTokenBalances(
            staker, strategies, tokenBalances, "staker should once again have original token tokenBalances"
        );
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    // function testFuzz_deposit_delegate_allocate_queue_slash_completeAsTokens(
    //     uint24 _random
    // ) public {
    //     _configRand({_randomSeed: _random, _assetTypes: HOLDS_MAX, _userTypes: DEFAULT});
    //     _upgradeEigenLayerContracts(); // Upgrade contracts if forkType is not local

    //     (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
    //     (User operator,,) = _newRandomOperator();
    //     (AVS avs,) = _newRandomAVS();

    //     assertEq(strategies.length, 33, "sanity");

    //     uint256[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

    //     assert_HasNoDelegatableShares(staker, "staker should not have delegatable shares before depositing");
    //     assertFalse(delegationManager.isDelegated(address(staker)), "staker should not be delegated");

    //     /// 1. Deposit Into Strategies
    //     staker.depositIntoEigenlayer(strategies, tokenBalances);
    //     check_Deposit_State(staker, strategies, shares);

    //     // 2. Delegate to an operator
    //     staker.delegateTo(operator);
    //     check_Delegation_State(staker, operator, strategies, shares);

    //     // Create an operator set and register an operator.
    //     OperatorSet memory operatorSet = avs.createOperatorSet(strategies);
    //     operator.registerForOperatorSet(operatorSet);

    //     // 3. Allocate to operator set.
    //     operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
    //     cheats.roll(block.number + allocationManager.ALLOCATION_CONFIGURATION_DELAY());

    //     // 4. Slash operator
    //     SlashingParams memory slashingParams = avs.slashOperator(operator, operatorSet.id, _randWadToSlash());

    //     // 5. Queue withdrawal
    // }

    // function testFuzz_deposit_delegate_allocate_deallocate_slash_queue_completeAsTokens(
    //     uint24 _random
    // ) public {}

    // function testFuzz_deposit_delegate_allocate_deregister_slash(
    //     uint24 _random
    // ) public {}
}
