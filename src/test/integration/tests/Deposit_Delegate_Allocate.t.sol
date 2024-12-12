// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";

// TODO: move randomness from tests

contract Integration_Deposit_Delegate_Allocate is IntegrationCheckUtils {
    function testFuzz_deposit_delegate_allocate(
        uint24 _random
    ) public {
        // Configure the random parameters for the test
        _configRand({
            _randomSeed: _random,
            _assetTypes: HOLDS_LST | HOLDS_ETH | HOLDS_ALL,
            _userTypes: DEFAULT | ALT_METHODS
        });

        _upgradeEigenLayerContracts();

        // Create a staker and an operator with a nonzero balance and corresponding strategies
        (, OperatorSet[] memory operatorSets) = _newRandomAVS();
        (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();

        // 1. Delegate to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, new uint256[](strategies.length)); // Initial shares are zero

        // 2. Deposit into strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint256[] memory shares = _calculateExpectedShares(strategies, tokenBalances);

        // Check that the deposit increased operator shares the staker is delegated to
        check_Deposit_State(staker, strategies, shares);
        assert_Snap_Added_OperatorShares(operator, strategies, shares, "operator should have received shares");

        operator.registerForOperatorSets(operatorSets);

        IAllocationManagerTypes.AllocateParams[] memory allocateParams =
            new IAllocationManagerTypes.AllocateParams[](operatorSets.length);

        for (uint256 i; i < operatorSets.length; ++i) {
            uint256 len = allocationManager.getStrategiesInOperatorSet(operatorSets[i]).length;
            allocateParams[i] = operator.modifyAllocations(
                operatorSets[i], _randMagnitudes({sum: WAD / uint64(operatorSets.length), len: len})
            );
            assert_Snap_Allocations_Modified(
                operator, allocateParams[i], false, "operator allocations should be updated before delay"
            );
        }

        _rollBlocksForCompleteAllocation();

        for (uint256 i; i < operatorSets.length; ++i) {
            assert_Snap_Allocations_Modified(
                operator, allocateParams[i], true, "operator allocations should be updated after delay"
            );
        }
    }

    function testFuzz_deposit_delegate_allocate_slash_undelegate_completeAsTokens(
        uint24 _random
    ) public {
        _configRand({_randomSeed: _random, _assetTypes: HOLDS_MAX, _userTypes: DEFAULT});
        _upgradeEigenLayerContracts(); // Upgrade contracts if forkType is not local

        (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        (AVS avs,) = _newRandomAVS();

        if (forkType == LOCAL) assertEq(strategies.length, 33, "sanity");

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);

        // 2. Delegate to an operator
        staker.delegateTo(operator);

        // Create an operator set and register an operator.
        OperatorSet memory operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);

        // 3. Allocate to operator set.
        IAllocationManagerTypes.AllocateParams memory allocateParams =
            operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
        assert_Snap_Allocations_Modified(
            operator, allocateParams, false, "operator allocations should be updated before delay"
        );
        _rollBlocksForCompleteAllocation();
        assert_Snap_Allocations_Modified(
            operator, allocateParams, true, "operator allocations should be updated after delay"
        );

        (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
            _randStrategiesAndWadsToSlash(operatorSet);

        // 4. Slash operator
        IAllocationManagerTypes.SlashingParams memory slashingParams =
            avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);
        assert_Snap_Allocations_Slashed(slashingParams, operatorSet, true, "operator allocations should be slashed");

        // 5. Undelegate from an operator
        IDelegationManagerTypes.Withdrawal[] memory withdrawals = staker.undelegate();
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);
        // 6. Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint256[] memory expectedTokens =
                _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            // FIXME: check_Withdrawal_AsTokens_State_AfterSlash(staker, operator, withdrawals[i], allocateParams, slashingParams, expectedTokens);
        }

        // Check Final State
        assert_HasNoDelegatableShares(staker, "staker should have withdrawn all shares");
        assert_HasUnderlyingTokenBalances_AfterSlash(
            staker,
            allocateParams,
            slashingParams,
            tokenBalances,
            "staker should once again have original token tokenBalances minus slashed"
        );
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_deposit_delegate_allocate_queue_slash_completeAsTokens(
        uint24 _random
    ) public {
        _configRand({_randomSeed: _random, _assetTypes: HOLDS_MAX, _userTypes: DEFAULT});
        _upgradeEigenLayerContracts(); // Upgrade contracts if forkType is not local

        (User staker, IStrategy[] memory strategies, uint256[] memory tokenBalances) = _newRandomStaker();
        (User operator,,) = _newRandomOperator();
        (AVS avs,) = _newRandomAVS();

        if (forkType == LOCAL) assertEq(strategies.length, 33, "sanity");

        // 1. Deposit Into Strategies
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        // 2. Delegate to an operator
        staker.delegateTo(operator);

        // Create an operator set and register an operator.
        OperatorSet memory operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);

        // 3. Allocate to operator set.
        IAllocationManagerTypes.AllocateParams memory allocateParams =
            operator.modifyAllocations(operatorSet, _randMagnitudes({sum: 1 ether, len: strategies.length}));
        assert_Snap_Allocations_Modified(
            operator, allocateParams, false, "operator allocations should be updated before delay"
        );
        _rollBlocksForCompleteAllocation();
        assert_Snap_Allocations_Modified(
            operator, allocateParams, true, "operator allocations should be updated after delay"
        );

        // 4. Queue withdrawal
        IDelegationManagerTypes.Withdrawal[] memory withdrawals =
            staker.queueWithdrawals(strategies, _calculateExpectedShares(strategies, tokenBalances));
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        // 5. Slash operator
        (IStrategy[] memory strategiesToSlash, uint256[] memory wadsToSlash) =
            _randStrategiesAndWadsToSlash(operatorSet);
        IAllocationManagerTypes.SlashingParams memory slashingParams =
            avs.slashOperator(operator, operatorSet.id, strategiesToSlash, wadsToSlash);

        // 6. Complete withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint256 i = 0; i < withdrawals.length; ++i) {
            uint256[] memory expectedTokens =
                _calculateExpectedTokens(withdrawals[i].strategies, withdrawals[i].scaledShares);
            staker.completeWithdrawalAsTokens(withdrawals[i]);
            // FIXME: check_Withdrawal_AsTokens_State_AfterSlash(staker, operator, withdrawals[i], allocateParams, slashingParams, expectedTokens);
        }

        // Check Final State
        assert_HasNoDelegatableShares(staker, "staker should have withdrawn all shares");
        assert_HasUnderlyingTokenBalances_AfterSlash(
            staker,
            allocateParams,
            slashingParams,
            tokenBalances,
            "staker should once again have original token tokenBalances minus slashed"
        );
        assert_NoWithdrawalsPending(withdrawalRoots, "all withdrawals should be removed from pending");
    }

    function testFuzz_deposit_delegate_allocate_deallocate_slash_queue_completeAsTokens(
        uint24 _random
    ) public {}

    function testFuzz_deposit_delegate_allocate_deregister_slash(
        uint24 _random
    ) public {}
}
