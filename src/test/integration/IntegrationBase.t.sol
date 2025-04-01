// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

import "src/contracts/libraries/BeaconChainProofs.sol";
import "src/contracts/libraries/SlashingLib.sol";

import "src/test/integration/IntegrationGetters.t.sol";
import "src/test/integration/users/User.t.sol";
import "src/test/integration/users/User_M1.t.sol";

abstract contract IntegrationBase is IntegrationGetters {
    using ArrayLib for *;
    using Math for *;
    using SlashingLib for *;
    using StdStyle for *;

    /**
     *
     *                             COMMON ASSERTIONS
     *
     */
    function assert_HasNoDelegatableShares(User user, string memory err) internal view {
        (IStrategy[] memory strategies, uint[] memory shares) = delegationManager.getDepositedShares(address(user));
        assertEq(strategies.length, 0, err);
        assertEq(strategies.length, shares.length, "assert_HasNoDelegatableShares: return length mismatch");
    }

    function assert_HasUnderlyingTokenBalances(User user, IStrategy[] memory strategies, uint[] memory expectedBalances, string memory err)
        internal
        view
    {
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint expectedBalance = expectedBalances[i];
            uint tokenBalance = strat == BEACONCHAIN_ETH_STRAT ? address(user).balance : strat.underlyingToken().balanceOf(address(user));
            assertApproxEqAbs(expectedBalance, tokenBalance, 1, err);
        }
    }

    function assert_HasNoUnderlyingTokenBalance(User user, IStrategy[] memory strategies, string memory err) internal view {
        assert_HasUnderlyingTokenBalances(user, strategies, new uint[](strategies.length), err);
    }

    function assert_HasExpectedShares(User user, IStrategy[] memory strategies, uint[] memory expectedShares, string memory err)
        internal
        view
    {
        uint[] memory actualShares = _getStakerDepositShares(user, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertApproxEqAbs(expectedShares[i], actualShares[i], 1, err);
        }
    }

    /// @dev Check that all the staker's deposit shares have been removed
    function assert_RemovedAll_Staker_DepositShares(User user, IStrategy[] memory strategies, string memory err) internal view {
        uint[] memory depositShares = _getStakerDepositShares(user, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(depositShares[i], 0, err);
        }
    }

    /// @dev Check that all the staker's withdrawable shares have been removed
    function assert_RemovedAll_Staker_WithdrawableShares(User staker, IStrategy[] memory strategies, string memory err) internal view {
        uint[] memory curShares = _getStakerWithdrawableShares(staker, strategies);
        // For each strategy, check all shares have been withdrawn
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(0, curShares[i], err);
        }
    }

    /// @dev Asserts that ALL of the `withdrawalRoots` is in `delegationManager.pendingWithdrawals`
    function assert_AllWithdrawalsPending(bytes32[] memory withdrawalRoots, string memory err) internal view {
        for (uint i = 0; i < withdrawalRoots.length; i++) {
            assert_WithdrawalPending(withdrawalRoots[i], err);
        }
    }

    /// @dev Asserts that NONE of the `withdrawalRoots` is in `delegationManager.pendingWithdrawals`
    function assert_NoWithdrawalsPending(bytes32[] memory withdrawalRoots, string memory err) internal view {
        for (uint i = 0; i < withdrawalRoots.length; i++) {
            assert_WithdrawalNotPending(withdrawalRoots[i], err);
        }
    }

    /// @dev Asserts that the hash of each withdrawal corresponds to the provided withdrawal root
    function assert_WithdrawalPending(bytes32 withdrawalRoot, string memory err) internal view {
        assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), err);
    }

    function assert_WithdrawalNotPending(bytes32 withdrawalRoot, string memory err) internal view {
        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), err);
    }

    function assert_ValidWithdrawalHashes(Withdrawal[] memory withdrawals, bytes32[] memory withdrawalRoots, string memory err)
        internal
        view
    {
        for (uint i = 0; i < withdrawals.length; i++) {
            assert_ValidWithdrawalHash(withdrawals[i], withdrawalRoots[i], err);
        }
    }

    function assert_ValidWithdrawalHash(Withdrawal memory withdrawal, bytes32 withdrawalRoot, string memory err) internal view {
        assertEq(withdrawalRoot, delegationManager.calculateWithdrawalRoot(withdrawal), err);
    }

    function assert_StakerStrategyListEmpty(User staker, string memory err) internal view {
        IStrategy[] memory strategies = _getStakerStrategyList(staker);
        assertEq(strategies.length, 0, err);
    }

    function assert_StrategyNotInStakerStrategyList(User staker, IStrategy strategy, string memory err) internal view {
        // BEACONCHAIN_ETH_STRAT is not in the staker's strategy list
        if (strategy == BEACONCHAIN_ETH_STRAT) return;
        IStrategy[] memory strategies = _getStakerStrategyList(staker);
        assertFalse(strategies.contains(strategy), err);
    }

    function assert_StrategiesInStakerStrategyList(User staker, IStrategy[] memory strategies, string memory err) internal view {
        for (uint i = 0; i < strategies.length; i++) {
            assert_StrategyInStakerStrategyList(staker, strategies[i], err);
        }
    }

    function assert_StrategyInStakerStrategyList(User staker, IStrategy strategy, string memory err) internal view {
        // BEACONCHAIN_ETH_STRAT is not in the staker's strategy list
        if (strategy == BEACONCHAIN_ETH_STRAT) return;
        IStrategy[] memory strategies = _getStakerStrategyList(staker);
        assertTrue(strategies.contains(strategy), err);
    }

    function assert_PodBalance_Eq(User staker, uint expectedBalance, string memory err) internal view {
        EigenPod pod = staker.pod();
        assertEq(address(pod).balance, expectedBalance, err);
    }

    function assert_ProofsRemainingEqualsActive(User staker, string memory err) internal view {
        EigenPod pod = staker.pod();
        console.log("proofsRemaining: ", pod.currentCheckpoint().proofsRemaining);
        console.log("activeValidatorCount: ", pod.activeValidatorCount());
        assertEq(pod.currentCheckpoint().proofsRemaining, pod.activeValidatorCount(), err);
    }

    function assert_CheckpointPodBalance(User staker, uint64 expectedPodBalanceGwei, string memory err) internal view {
        EigenPod pod = staker.pod();
        assertEq(pod.currentCheckpoint().podBalanceGwei, expectedPodBalanceGwei, err);
    }

    function assert_MaxEqualsAllocatablePlusEncumbered(User operator, string memory err) internal view {
        Magnitudes[] memory mags = _getMagnitudes(operator, allStrats);
        for (uint i = 0; i < allStrats.length; i++) {
            Magnitudes memory m = mags[i];
            assertEq(m.max, m.encumbered + m.allocatable, err);
        }
    }

    function assert_CurMinSlashableEqualsMinAllocated(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal view {
        assertEq(_getMinSlashableStake(operator, operatorSet, strategies), _getAllocatedStake(operator, operatorSet, strategies), err);
    }

    function assert_MaxMagsEqualMaxMagsAtCurrentBlock(User operator, IStrategy[] memory strategies, string memory err) internal view {
        assertEq(
            _getMaxMagnitudes(operator, strategies).toUintArray(),
            _getMaxMagnitudes(operator, strategies, uint32(block.number)).toUintArray(),
            err
        );
    }

    function assert_CurrentMagnitude(User operator, AllocateParams memory params, string memory err) internal view {
        Allocation[] memory allocations = _getAllocations(operator, params.operatorSet, params.strategies);
        for (uint i = 0; i < allocations.length; i++) {
            assertEq(allocations[i].currentMagnitude, params.newMagnitudes[i], err);
        }
    }

    function assert_NoPendingModification(User operator, OperatorSet memory operatorSet, IStrategy[] memory strategies, string memory err)
        internal
        view
    {
        Allocation[] memory allocations = _getAllocations(operator, operatorSet, strategies);
        for (uint i = 0; i < allocations.length; i++) {
            assertEq(0, allocations[i].effectBlock, err);
        }
    }

    function assert_HasPendingIncrease(User operator, AllocateParams memory params, string memory err) internal view {
        uint32 delay = _getExistingAllocationDelay(operator);
        Allocation[] memory allocations = _getAllocations(operator, params.operatorSet, params.strategies);
        for (uint i = 0; i < allocations.length; i++) {
            assertEq(allocations[i].effectBlock, uint32(block.number) + delay, err);
            assertTrue(allocations[i].currentMagnitude != params.newMagnitudes[i], err);
            assertGt(allocations[i].pendingDiff, 0, err);
        }
    }

    function assert_HasPendingDecrease(User operator, AllocateParams memory params, string memory err) internal view {
        uint32 deallocationDelay = allocationManager.DEALLOCATION_DELAY();
        Allocation[] memory allocations = _getAllocations(operator, params.operatorSet, params.strategies);
        for (uint i = 0; i < allocations.length; i++) {
            assertEq(allocations[i].effectBlock, uint32(block.number) + deallocationDelay + 1, err);
            assertTrue(allocations[i].currentMagnitude != params.newMagnitudes[i], err);
            assertLt(allocations[i].pendingDiff, 0, err);
        }
    }

    function assert_IsSlashable(User operator, OperatorSet memory operatorSet, string memory err) internal view {
        assertTrue(allocationManager.isOperatorSlashable(address(operator), operatorSet), err);
    }

    function assert_NotSlashable(User operator, OperatorSet memory operatorSet, string memory err) internal view {
        assertFalse(allocationManager.isOperatorSlashable(address(operator), operatorSet), err);
    }

    function assert_IsAllocatedToSet(User operator, OperatorSet memory operatorSet, string memory err) internal view {
        assertTrue(allocationManager.getAllocatedSets(address(operator)).contains(operatorSet), err);
    }

    function assert_IsAllocatedToSetStrats(User operator, OperatorSet memory operatorSet, IStrategy[] memory strategies, string memory err)
        internal
        view
    {
        IStrategy[] memory allocatedStrategies = allocationManager.getAllocatedStrategies(address(operator), operatorSet);
        for (uint i = 0; i < strategies.length; i++) {
            assertTrue(allocatedStrategies.contains(strategies[i]), err);
        }
    }

    function assert_HasAllocatedStake(User operator, AllocateParams memory params, string memory err) internal view {
        OperatorSet memory operatorSet = params.operatorSet;
        IStrategy[] memory strategies = params.strategies;
        uint64[] memory curMagnitudes = params.newMagnitudes;
        uint64[] memory maxMagnitudes = _getMaxMagnitudes(operator, params.strategies);
        uint[] memory operatorShares = _getOperatorShares(operator, params.strategies);
        uint[] memory allocatedStake = _getAllocatedStake(operator, operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            uint expectedAllocated;
            if (maxMagnitudes[i] == 0) {
                expectedAllocated = 0;
            } else {
                uint slashableProportion = uint(curMagnitudes[i]).divWad(maxMagnitudes[i]);
                expectedAllocated = operatorShares[i].mulWad(slashableProportion);
            }
            assertEq(expectedAllocated, allocatedStake[i], err);
        }
    }

    function assert_HasSlashableStake(User operator, AllocateParams memory params, string memory err) internal view {
        OperatorSet memory operatorSet = params.operatorSet;
        IStrategy[] memory strategies = params.strategies;
        uint64[] memory curMagnitudes = params.newMagnitudes;
        uint64[] memory maxMagnitudes = _getMaxMagnitudes(operator, params.strategies);
        uint[] memory operatorShares = _getOperatorShares(operator, params.strategies);
        uint[] memory slashableStake = _getMinSlashableStake(operator, operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            uint expectedSlashable;
            if (maxMagnitudes[i] == 0) {
                expectedSlashable = 0;
            } else {
                uint slashableProportion = uint(curMagnitudes[i]).divWad(maxMagnitudes[i]);
                expectedSlashable = operatorShares[i].mulWad(slashableProportion);
            }
            assertEq(expectedSlashable, slashableStake[i], err);
        }
    }

    function assert_NoSlashableStake(User operator, OperatorSet memory operatorSet, string memory err) internal view {
        IStrategy[] memory strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);
        uint[] memory slashableStake = _getMinSlashableStake(operator, operatorSet, strategies);
        for (uint i = 0; i < slashableStake.length; i++) {
            assertEq(slashableStake[i], 0, err);
        }
    }

    function assert_DSF_WAD(User staker, IStrategy[] memory strategies, string memory err) internal view {
        uint[] memory depositScalingFactors = _getDepositScalingFactors(staker, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(depositScalingFactors[i], WAD, err);
        }
    }

    function assert_Zero_BCSF(User staker, string memory err) internal view {
        uint64 curBCSF = _getBeaconChainSlashingFactor(staker);
        assertEq(curBCSF, 0, err);
    }

    function assert_BCSF_WAD(User staker, string memory err) internal view {
        uint64 curBCSF = _getBeaconChainSlashingFactor(staker);
        assertEq(curBCSF, WAD, err);
    }

    function assert_ActiveValidatorCount(User staker, uint expectedCount, string memory err) internal view {
        uint curActiveValidatorCount = _getActiveValidatorCount(staker);
        assertEq(curActiveValidatorCount, expectedCount, err);
    }

    function assert_withdrawableSharesDecreasedByAtLeast(
        User staker,
        IStrategy[] memory strategies,
        uint[] memory originalShares,
        uint[] memory expectedDecreases,
        string memory err
    ) internal view {
        for (uint i = 0; i < strategies.length; i++) {
            assert_withdrawableSharesDecreasedByAtLeast(staker, strategies[i], originalShares[i], expectedDecreases[i], err);
        }
    }

    function assert_withdrawableSharesDecreasedByAtLeast(
        User staker,
        IStrategy strategy,
        uint originalShares,
        uint expectedDecrease,
        string memory err
    ) internal view {
        uint currentShares = _getWithdrawableShares(staker, strategy);
        assertLt(currentShares, originalShares - expectedDecrease, err);
    }

    function assert_DepositShares_GTE_WithdrawableShares(User staker, IStrategy[] memory strategies, string memory err) internal view {
        uint[] memory depositShares = _getStakerDepositShares(staker, strategies);
        uint[] memory withdrawableShares = _getWithdrawableShares(staker, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertGe(depositShares[i], withdrawableShares[i], err);
        }
    }

    function assert_Zero_WithdrawableShares(User staker, IStrategy strategy, string memory err) internal view {
        assertEq(_getWithdrawableShares(staker, strategy), 0, err);
    }

    /**
     *
     *                             SNAPSHOT ASSERTIONS
     *                    TIME TRAVELERS ONLY BEYOND THIS POINT
     *
     */

    /**
     *
     *                      SNAPSHOT ASSERTIONS: ALLOCATIONS
     *
     */
    function assert_Snap_Became_Registered(User operator, OperatorSet memory operatorSet, string memory err) internal {
        bool curIsMemberOfSet = _getIsMemberOfSet(operator, operatorSet);
        bool prevIsMemberOfSet = _getPrevIsMemberOfSet(operator, operatorSet);
        assertFalse(prevIsMemberOfSet, err);
        assertTrue(curIsMemberOfSet, err);
    }

    function assert_Snap_Became_Deregistered(User operator, OperatorSet memory operatorSet, string memory err) internal {
        bool curIsMemberOfSet = _getIsMemberOfSet(operator, operatorSet);
        bool prevIsMemberOfSet = _getPrevIsMemberOfSet(operator, operatorSet);
        assertTrue(prevIsMemberOfSet, err);
        assertFalse(curIsMemberOfSet, err);
    }

    function assert_Snap_Unchanged_Registration(User operator, OperatorSet memory operatorSet, string memory err) internal {
        bool curIsMemberOfSet = _getIsMemberOfSet(operator, operatorSet);
        bool prevIsMemberOfSet = _getPrevIsMemberOfSet(operator, operatorSet);
        assertEq(prevIsMemberOfSet, curIsMemberOfSet, err);
    }

    function assert_Snap_Became_Slashable(User operator, OperatorSet memory operatorSet, string memory err) internal {
        bool curIsSlashable = _getIsSlashable(operator, operatorSet);
        bool prevIsSlashable = _getPrevIsSlashable(operator, operatorSet);
        assertFalse(prevIsSlashable, err);
        assertTrue(curIsSlashable, err);
    }

    function assert_Snap_Remains_Slashable(User operator, OperatorSet memory operatorSet, string memory err) internal {
        bool curIsSlashable = _getIsSlashable(operator, operatorSet);
        bool prevIsSlashable = _getPrevIsSlashable(operator, operatorSet);
        assertTrue(prevIsSlashable, err);
        assertTrue(curIsSlashable, err);
    }

    function assert_Snap_Unchanged_Slashability(User operator, OperatorSet memory operatorSet, string memory err) internal {
        bool curIsSlashable = _getIsSlashable(operator, operatorSet);
        bool prevIsSlashable = _getPrevIsSlashable(operator, operatorSet);
        assertEq(prevIsSlashable, curIsSlashable, err);
    }

    function assert_Snap_Unchanged_AllocatedStrats(User operator, OperatorSet memory operatorSet, string memory err) internal {
        assertEq(
            _getAllocatedStrats(operator, operatorSet).toAddressArray(),
            _getPrevAllocatedStrats(operator, operatorSet).toAddressArray(),
            err
        );
    }

    function assert_Snap_Removed_AllocatedStrats(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory removedStrats,
        string memory err
    ) internal {
        IStrategy[] memory curAllocatedStrats = _getAllocatedStrats(operator, operatorSet);
        IStrategy[] memory prevAllocatedStrats = _getPrevAllocatedStrats(operator, operatorSet);
        assertEq(curAllocatedStrats.length + removedStrats.length, prevAllocatedStrats.length, err);
        for (uint i = 0; i < removedStrats.length; i++) {
            assertFalse(curAllocatedStrats.contains(removedStrats[i]), err);
        }
    }

    function assert_Snap_Unchanged_StrategyAllocations(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        Allocation[] memory curAllocations = _getAllocations(operator, operatorSet, strategies);
        Allocation[] memory prevAllocations = _getPrevAllocations(operator, operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            Allocation memory curAllocation = curAllocations[i];
            Allocation memory prevAllocation = prevAllocations[i];
            assertEq(curAllocation.currentMagnitude, prevAllocation.currentMagnitude, err);
            assertEq(curAllocation.pendingDiff, prevAllocation.pendingDiff, err);
            assertEq(curAllocation.effectBlock, prevAllocation.effectBlock, err);
        }
    }

    function assert_Snap_Unchanged_AllocatedSets(User operator, string memory err) internal {
        OperatorSet[] memory curAllocatedSets = _getAllocatedSets(operator);
        OperatorSet[] memory prevAllocatedSets = _getPrevAllocatedSets(operator);
        assertEq(curAllocatedSets.length, prevAllocatedSets.length, err);
    }

    function assert_Snap_Removed_AllocatedSet(User operator, OperatorSet memory operatorSet, string memory err) internal {
        OperatorSet[] memory curAllocatedSets = _getAllocatedSets(operator);
        OperatorSet[] memory prevAllocatedSets = _getPrevAllocatedSets(operator);
        assertEq(curAllocatedSets.length + 1, prevAllocatedSets.length, err);
        assertTrue(prevAllocatedSets.contains(operatorSet), err);
        assertFalse(curAllocatedSets.contains(operatorSet), err);
    }

    function assert_Snap_Added_RegisteredSet(User operator, OperatorSet memory operatorSet, string memory err) internal {
        OperatorSet[] memory curRegisteredSets = _getRegisteredSets(operator);
        OperatorSet[] memory prevRegisteredSets = _getPrevRegisteredSets(operator);
        assertEq(curRegisteredSets.length, prevRegisteredSets.length + 1, err);
        assertFalse(prevRegisteredSets.contains(operatorSet), err);
        assertTrue(curRegisteredSets.contains(operatorSet), err);
    }

    function assert_Snap_Removed_RegisteredSet(User operator, OperatorSet memory operatorSet, string memory err) internal {
        OperatorSet[] memory curRegisteredSets = _getRegisteredSets(operator);
        OperatorSet[] memory prevRegisteredSets = _getPrevRegisteredSets(operator);
        assertEq(curRegisteredSets.length + 1, prevRegisteredSets.length, err);
        assertTrue(prevRegisteredSets.contains(operatorSet), err);
        assertFalse(curRegisteredSets.contains(operatorSet), err);
    }

    function assert_Snap_Unchanged_RegisteredSet(User operator, string memory err) internal {
        OperatorSet[] memory curRegisteredSets = _getRegisteredSets(operator);
        OperatorSet[] memory prevRegisteredSets = _getPrevRegisteredSets(operator);
        assertEq(curRegisteredSets.length, prevRegisteredSets.length, err);
        for (uint i = 0; i < curRegisteredSets.length; i++) {
            assertEq(curRegisteredSets[i].avs, prevRegisteredSets[i].avs, err);
            assertEq(curRegisteredSets[i].id, prevRegisteredSets[i].id, err);
        }
    }

    function assert_Snap_Added_MemberOfSet(User operator, OperatorSet memory operatorSet, string memory err) internal {
        address[] memory curOperators = _getMembers(operatorSet);
        address[] memory prevOperators = _getPrevMembers(operatorSet);
        assertEq(curOperators.length, prevOperators.length + 1, err);
        assertFalse(prevOperators.contains(address(operator)), err);
        assertTrue(curOperators.contains(address(operator)), err);
    }

    function assert_Snap_Removed_MemberOfSet(User operator, OperatorSet memory operatorSet, string memory err) internal {
        address[] memory curOperators = _getMembers(operatorSet);
        address[] memory prevOperators = _getPrevMembers(operatorSet);
        assertEq(curOperators.length + 1, prevOperators.length, err);
        assertTrue(prevOperators.contains(address(operator)), err);
        assertFalse(curOperators.contains(address(operator)), err);
    }

    function assert_Snap_Unchanged_MemberOfSet(OperatorSet memory operatorSet, string memory err) internal {
        assertEq(_getMembers(operatorSet), _getPrevMembers(operatorSet), err);
    }

    function assert_Snap_StakeBecameSlashable(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        uint[] memory curSlashableStake = _getMinSlashableStake(operator, operatorSet, strategies);
        uint[] memory prevSlashableStake = _getPrevMinSlashableStake(operator, operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertTrue(prevSlashableStake[i] < curSlashableStake[i], err);
        }
    }

    function assert_Snap_StakeBecomeUnslashable(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        uint[] memory curSlashableStake = _getMinSlashableStake(address(operator), operatorSet, strategies);
        uint[] memory prevSlashableStake = _getPrevMinSlashableStake(address(operator), operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertTrue(prevSlashableStake[i] > curSlashableStake[i], err);
            assertTrue(prevSlashableStake[i] > curSlashableStake[i], err);
        }
    }

    function assert_Snap_Unchanged_SlashableStake(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        assertEq(
            _getMinSlashableStake(operator, operatorSet, strategies), _getPrevMinSlashableStake(operator, operatorSet, strategies), err
        );
    }

    function assert_Snap_Slashed_SlashableStake(
        User operator,
        OperatorSet memory operatorSet,
        SlashingParams memory params,
        string memory err
    ) internal {
        uint[] memory curSlashableStake = _getMinSlashableStake(operator, operatorSet, params.strategies);
        uint[] memory prevSlashableStake = _getPrevMinSlashableStake(operator, operatorSet, params.strategies);
        Magnitudes[] memory curMagnitudes = _getMagnitudes(operator, params.strategies);
        Magnitudes[] memory prevMagnitudes = _getPrevMagnitudes(operator, params.strategies);
        for (uint i = 0; i < params.strategies.length; i++) {
            // Slashing doesn't occur if the operator has no slashable magnitude
            // This prevents a div by 0 when calculating expected slashed
            uint expectedSlashed = prevMagnitudes[i].max == 0
                ? 0
                : SlashingLib.calcSlashedAmount({
                    operatorShares: prevSlashableStake[i],
                    prevMaxMagnitude: prevMagnitudes[i].max,
                    newMaxMagnitude: curMagnitudes[i].max
                });
            assertEq(curSlashableStake[i], prevSlashableStake[i] - expectedSlashed, err);
        }
    }

    /// @dev requires slashparams strategies to be same as withdrawal strategies meant to be used in check_base_slashing_state
    function assert_Snap_Decreased_SlashableSharesInQueue(
        User operator,
        SlashingParams memory slashParams,
        Withdrawal[] memory withdrawals,
        string memory err
    ) internal {
        IStrategy[] memory strategies = slashParams.strategies;
        uint[] memory curSlashableSharesInQueue = _getSlashableSharesInQueue(operator, strategies);
        uint[] memory prevSlashableSharesInQueue = _getPrevSlashableSharesInQueue(operator, strategies);
        uint[] memory totalScaledShares = new uint[](strategies.length);
        for (uint i = 0; i < withdrawals.length; i++) {
            for (uint j = 0; j < withdrawals[i].strategies.length; j++) {
                totalScaledShares[j] = totalScaledShares[j] + withdrawals[i].scaledShares[j];
            }
        }
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(
                curSlashableSharesInQueue[i], prevSlashableSharesInQueue[i] - totalScaledShares[i].mulWad(slashParams.wadsToSlash[i]), err
            );
        }
    }

    function assert_Snap_Increased_SlashableSharesInQueue(User operator, Withdrawal[] memory withdrawals, string memory err) internal {
        uint[] memory curSlashableSharesInQueue;
        uint[] memory prevSlashableSharesInQueue;
        uint64[] memory maxMagnitudes;
        for (uint i = 0; i < withdrawals.length; i++) {
            curSlashableSharesInQueue = _getSlashableSharesInQueue(operator, withdrawals[i].strategies);
            prevSlashableSharesInQueue = _getPrevSlashableSharesInQueue(operator, withdrawals[i].strategies);
            maxMagnitudes = _getMaxMagnitudes(operator, withdrawals[i].strategies);
            for (uint j = 0; j < withdrawals[i].strategies.length; j++) {
                assertEq(
                    curSlashableSharesInQueue[j],
                    prevSlashableSharesInQueue[j] + withdrawals[i].scaledShares[j].mulWad(maxMagnitudes[j]),
                    err
                );
            }
        }
    }

    function assert_Snap_StakeBecameAllocated(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        uint[] memory curMinAllocatedStake = _getAllocatedStake(operator, operatorSet, strategies);
        uint[] memory prevMinAllocatedStake = _getPrevAllocatedStake(operator, operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertGt(curMinAllocatedStake[i], prevMinAllocatedStake[i], err);
        }
    }

    function assert_Snap_StakeBecameDeallocated(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        uint[] memory curMinAllocatedStake = _getAllocatedStake(operator, operatorSet, strategies);
        uint[] memory prevMinAllocatedStake = _getPrevAllocatedStake(operator, operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertLt(curMinAllocatedStake[i], prevMinAllocatedStake[i], err);
        }
    }

    function assert_Snap_Unchanged_AllocatedStake(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        assertEq(_getAllocatedStake(operator, operatorSet, strategies), _getPrevAllocatedStake(operator, operatorSet, strategies), err);
    }

    function assert_Snap_Slashed_AllocatedStake(
        User operator,
        OperatorSet memory operatorSet,
        SlashingParams memory params,
        string memory err
    ) internal {
        uint[] memory curAllocatedStake = _getAllocatedStake(operator, operatorSet, params.strategies);
        uint[] memory prevAllocatedStake = _getPrevAllocatedStake(operator, operatorSet, params.strategies);
        Magnitudes[] memory curMagnitudes = _getMagnitudes(operator, params.strategies);
        Magnitudes[] memory prevMagnitudes = _getPrevMagnitudes(operator, params.strategies);
        for (uint i = 0; i < curAllocatedStake.length; i++) {
            // Slashing doesn't occur if the operator has no slashable magnitude
            // This prevents a div by 0 when calculating expected slashed
            uint expectedSlashed = prevMagnitudes[i].max == 0
                ? 0
                : SlashingLib.calcSlashedAmount({
                    operatorShares: prevAllocatedStake[i],
                    prevMaxMagnitude: prevMagnitudes[i].max,
                    newMaxMagnitude: curMagnitudes[i].max
                });
            assertEq(curAllocatedStake[i], prevAllocatedStake[i] - expectedSlashed, err);
        }
    }

    function assert_Snap_Unchanged_EncumberedMagnitude(User operator, IStrategy[] memory strategies, string memory err) internal {
        Magnitudes[] memory curMagnitudes = _getMagnitudes(operator, strategies);
        Magnitudes[] memory prevMagnitudes = _getPrevMagnitudes(operator, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(curMagnitudes[i].encumbered, prevMagnitudes[i].encumbered, err);
        }
    }

    function assert_Snap_Removed_EncumberedMagnitude(
        User operator,
        IStrategy[] memory strategies,
        uint64[] memory magnitudeRemoved,
        string memory err
    ) internal {
        Magnitudes[] memory curMagnitudes = _getMagnitudes(operator, strategies);
        Magnitudes[] memory prevMagnitudes = _getPrevMagnitudes(operator, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(curMagnitudes[i].encumbered + magnitudeRemoved[i], prevMagnitudes[i].encumbered, err);
        }
    }

    function assert_Snap_Slashed_EncumberedMagnitude(User operator, SlashingParams memory params, string memory err) internal {
        Magnitudes[] memory curMagnitudes = _getMagnitudes(operator, params.strategies);
        Magnitudes[] memory prevMagnitudes = _getPrevMagnitudes(operator, params.strategies);
        for (uint i = 0; i < params.strategies.length; i++) {
            uint expectedSlashed = prevMagnitudes[i].encumbered.mulWadRoundUp(params.wadsToSlash[i]);
            assertEq(curMagnitudes[i].encumbered, prevMagnitudes[i].encumbered - expectedSlashed, err);
        }
    }

    function assert_Snap_Added_AllocatableMagnitude(
        User operator,
        IStrategy[] memory strategies,
        uint64[] memory magnitudeFreed,
        string memory err
    ) internal {
        Magnitudes[] memory curMagnitudes = _getMagnitudes(operator, strategies);
        Magnitudes[] memory prevMagnitudes = _getPrevMagnitudes(operator, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(curMagnitudes[i].allocatable, prevMagnitudes[i].allocatable + magnitudeFreed[i], err);
        }
    }

    function assert_Snap_Unchanged_AllocatableMagnitude(User operator, IStrategy[] memory strategies, string memory err) internal {
        Magnitudes[] memory curMagnitudes = _getMagnitudes(operator, strategies);
        Magnitudes[] memory prevMagnitudes = _getPrevMagnitudes(operator, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(curMagnitudes[i].allocatable, prevMagnitudes[i].allocatable, err);
        }
    }

    function assert_Snap_Allocated_Magnitude(User operator, IStrategy[] memory strategies, string memory err) internal {
        Magnitudes[] memory curMagnitudes = _getMagnitudes(operator, strategies);
        Magnitudes[] memory prevMagnitudes = _getPrevMagnitudes(operator, strategies);
        /// Check:
        /// allocatable increased
        /// encumbered decreased
        for (uint i = 0; i < strategies.length; i++) {
            assertLt(curMagnitudes[i].allocatable, prevMagnitudes[i].allocatable, err);
            assertGt(curMagnitudes[i].encumbered, prevMagnitudes[i].encumbered, err);
        }
    }

    function assert_Snap_Deallocated_Magnitude(User operator, IStrategy[] memory strategies, string memory err) internal {
        Magnitudes[] memory curMagnitudes = _getMagnitudes(operator, strategies);
        Magnitudes[] memory prevMagnitudes = _getPrevMagnitudes(operator, strategies);
        /// Check:
        /// allocatable increased
        /// encumbered decreased
        for (uint i = 0; i < strategies.length; i++) {
            assertGt(curMagnitudes[i].allocatable, prevMagnitudes[i].allocatable, err);
            assertLt(curMagnitudes[i].encumbered, prevMagnitudes[i].encumbered, err);
        }
    }

    function assert_Snap_Set_CurrentMagnitude(User operator, AllocateParams memory params, string memory err) internal {
        Allocation[] memory curAllocations = _getAllocations(operator, params.operatorSet, params.strategies);
        Allocation[] memory prevAllocations = _getPrevAllocations(operator, params.operatorSet, params.strategies);
        /// Prev allocation.currentMagnitude should NOT equal newly-set magnitude
        /// Cur allocation.currentMagnitude SHOULD
        for (uint i = 0; i < params.strategies.length; i++) {
            assertTrue(prevAllocations[i].currentMagnitude != params.newMagnitudes[i], err);
            assertEq(curAllocations[i].currentMagnitude, params.newMagnitudes[i], err);
        }
    }

    function assert_Snap_Slashed_Allocation(User operator, OperatorSet memory operatorSet, SlashingParams memory params, string memory err)
        internal
    {
        Allocation[] memory curAllocations = _getAllocations(operator, operatorSet, params.strategies);
        Allocation[] memory prevAllocations = _getPrevAllocations(operator, operatorSet, params.strategies);
        for (uint i = 0; i < params.strategies.length; i++) {
            uint expectedSlashed = prevAllocations[i].currentMagnitude.mulWadRoundUp(params.wadsToSlash[i]);
            assertEq(curAllocations[i].currentMagnitude, prevAllocations[i].currentMagnitude - expectedSlashed, err);
        }
    }

    function assert_Snap_Unchanged_MaxMagnitude(User operator, IStrategy[] memory strategies, string memory err) internal {
        Magnitudes[] memory curMagnitudes = _getMagnitudes(operator, strategies);
        Magnitudes[] memory prevMagnitudes = _getPrevMagnitudes(operator, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(curMagnitudes[i].max, prevMagnitudes[i].max, err);
        }
    }

    function assert_Snap_Slashed_MaxMagnitude(
        User operator,
        OperatorSet memory operatorSet,
        SlashingParams memory params,
        string memory err
    ) internal {
        Allocation[] memory prevAllocations = _getPrevAllocations(operator, operatorSet, params.strategies);
        Magnitudes[] memory curMagnitudes = _getMagnitudes(operator, params.strategies);
        Magnitudes[] memory prevMagnitudes = _getPrevMagnitudes(operator, params.strategies);
        for (uint i = 0; i < params.strategies.length; i++) {
            uint expectedSlashed = prevAllocations[i].currentMagnitude.mulWadRoundUp(params.wadsToSlash[i]);
            assertEq(curMagnitudes[i].max, prevMagnitudes[i].max - expectedSlashed, err);
        }
    }

    function assert_Snap_Allocations_Slashed(
        SlashingParams memory slashingParams,
        OperatorSet memory operatorSet,
        bool completed,
        string memory err
    ) internal {
        User op = User(payable(slashingParams.operator));
        Allocation[] memory curAllocs = _getAllocations(op, operatorSet, slashingParams.strategies);
        Allocation[] memory prevAllocs = _getPrevAllocations(op, operatorSet, slashingParams.strategies);
        Magnitudes[] memory curMagnitudes = _getMagnitudes(op, slashingParams.strategies);
        Magnitudes[] memory prevMagnitudes = _getPrevMagnitudes(op, slashingParams.strategies);
        uint32 delay = _getExistingAllocationDelay(User(payable(slashingParams.operator)));
        for (uint i = 0; i < slashingParams.strategies.length; i++) {
            Allocation memory curAlloc = curAllocs[i];
            Allocation memory prevAlloc = prevAllocs[i];
            uint64 slashedMagnitude = uint64(uint(prevAlloc.currentMagnitude).mulWadRoundUp(slashingParams.wadsToSlash[i]));
            // Check Allocations
            assertEq(curAlloc.currentMagnitude, prevAlloc.currentMagnitude - slashedMagnitude, string.concat(err, " (currentMagnitude)"));
            if (completed) {
                assertEq(curAlloc.pendingDiff, 0, string.concat(err, " (pendingDiff)"));
                assertEq(curAlloc.effectBlock, 0, string.concat(err, " (effectBlock)"));
            } else {
                assertEq(
                    curAlloc.currentMagnitude, prevAlloc.currentMagnitude - slashedMagnitude, string.concat(err, " (currentMagnitude)")
                );
                // If (isDeallocation) ...
                if (prevAlloc.pendingDiff < 0) {
                    uint64 slashedPending = uint64(uint(uint128(-prevAlloc.pendingDiff)).mulWadRoundUp(slashingParams.wadsToSlash[i]));
                    assertEq(
                        curAlloc.pendingDiff, prevAlloc.pendingDiff + int128(uint128(slashedPending)), string.concat(err, " (pendingDiff)")
                    );
                    delay = DEALLOCATION_DELAY + 1;
                }
                assertEq(curAlloc.effectBlock, block.number + delay, string.concat(err, " (effectBlock)"));
            }
            // Check Magnitudes
            Magnitudes memory curMagnitude = curMagnitudes[i];
            Magnitudes memory prevMagnitude = prevMagnitudes[i];
            assertEq(curMagnitude.encumbered, prevMagnitude.encumbered - slashedMagnitude, string.concat(err, " (encumbered magnitude)"));
            assertEq(curMagnitude.allocatable, prevMagnitude.allocatable, string.concat(err, " (allocatable magnitude)"));
            assertEq(curMagnitude.max, prevMagnitude.max - slashedMagnitude, string.concat(err, " (max magnitude)"));
        }
    }

    function assert_Snap_StakerWithdrawableShares_AfterSlash(
        User staker,
        AllocateParams memory allocateParams,
        SlashingParams memory slashingParams,
        string memory err
    ) internal {
        uint[] memory curShares = _getWithdrawableShares(staker, allocateParams.strategies);
        uint[] memory prevShares = _getPrevWithdrawableShares(staker, allocateParams.strategies);
        for (uint i = 0; i < allocateParams.strategies.length; i++) {
            IStrategy strat = allocateParams.strategies[i];
            uint slashedShares = 0;
            if (slashingParams.strategies.contains(strat)) {
                uint wadToSlash = slashingParams.wadsToSlash[slashingParams.strategies.indexOf(strat)];
                slashedShares = prevShares[i].mulWadRoundUp(allocateParams.newMagnitudes[i].mulWadRoundUp(wadToSlash));
            }
            assertApproxEqAbs(prevShares[i] - slashedShares, curShares[i], 1, err);
        }
    }

    /**
     *
     *                 SNAPSHOT ASSERTIONS: BEACON CHAIN AND AVS SLASHING
     *
     */

    /// @dev Same as `assert_Snap_StakerWithdrawableShares_AfterSlash`
    /// @dev but when a BC slash occurs before an AVS slash
    /// @dev There is additional rounding error when a BC and AVS slash occur together
    function assert_Snap_StakerWithdrawableShares_AfterBCSlash_AVSSlash(
        User staker,
        AllocateParams memory allocateParams,
        SlashingParams memory slashingParams,
        string memory err
    ) internal {
        require(allocateParams.strategies.length == 1 && slashingParams.strategies.length == 1, "only beacon strategy supported");
        require(allocateParams.strategies[0] == BEACONCHAIN_ETH_STRAT, "only beacon strategy supported");
        uint curShares = _getWithdrawableShares(staker, allocateParams.strategies)[0];
        uint prevShares = _getPrevWithdrawableShares(staker, allocateParams.strategies)[0];
        uint slashedShares = 0;
        uint wadToSlash = slashingParams.wadsToSlash[0];
        slashedShares = prevShares.mulWadRoundUp(allocateParams.newMagnitudes[0].mulWadRoundUp(wadToSlash));
        assertApproxEqAbs(prevShares - slashedShares, curShares, 1e2, err);
    }

    /// @dev Validates behavior of "restaking", ie. that the funds can be slashed twice
    function assert_Snap_StakerWithdrawableShares_AfterAVSSlash_BCSlash(
        User staker,
        AllocateParams memory allocateParams,
        SlashingParams memory slashingParams,
        string memory err
    ) internal {
        require(allocateParams.strategies.length == 1 && slashingParams.strategies.length == 1, "only beacon strategy supported");
        require(allocateParams.strategies[0] == BEACONCHAIN_ETH_STRAT, "only beacon strategy supported");
        uint curShares = _getWithdrawableShares(staker, allocateParams.strategies)[0];
        uint prevShares = _getPrevWithdrawableShares(staker, allocateParams.strategies)[0];
        uint depositShares = _getStakerDepositShares(staker, allocateParams.strategies)[0];
        // 1. The withdrawable shares should decrease by a factor of the BCSF
        assertApproxEqAbs(prevShares.mulWad(_getBeaconChainSlashingFactor(staker)), curShares, 1e2, err);
        /**
         * 2. The delta in shares is given by:
         * (depositShares * operatorMag) - (depositShares * operatorMag * BCSF)
         *  = depositShares * operatorMag * (1 - BCSF)
         */
        uint beaconChainSlashingFactor = _getBeaconChainSlashingFactor(staker);
        uint wadToSlash = slashingParams.wadsToSlash[0];
        uint originalAVSSlashedShares = depositShares.mulWadRoundUp(allocateParams.newMagnitudes[0].mulWadRoundUp(wadToSlash));
        uint withdrawableSharesAfterAVSSlash = depositShares - originalAVSSlashedShares;
        uint expectedDelta = withdrawableSharesAfterAVSSlash.mulWad(WAD - beaconChainSlashingFactor);
        assertApproxEqAbs(prevShares - expectedDelta, curShares, 1e2, err);

        /**
         * 3. The attributable avs slashed shares should decrease by a factor of the BCSF
         * Attributable avs slashed shares = originalWithdrawableShares - bcSlashedShares - curShares
         * Where bcSlashedShares = originalWithdrawableShares * (1 - BCSF)
         */
        uint bcSlashedShares = depositShares.mulWad(WAD - beaconChainSlashingFactor);
        uint attributableAVSSlashedShares = depositShares - bcSlashedShares - curShares;
        assertApproxEqAbs(originalAVSSlashedShares.mulWad(beaconChainSlashingFactor), attributableAVSSlashedShares, 1e2, err);
    }

    /**
     * @dev Validates behavior of "restaking", ie. that the funds can be slashed twice. Also validates
     *      the edge case where a validator is proven prior to the BC slash.
     * @dev These bounds are based off of rounding when avs and bc slashing occur together
     */
    function assert_Snap_StakerWithdrawableShares_AVSSlash_ValidatorProven_BCSlash(
        User staker,
        uint originalWithdrawableShares,
        uint extraValidatorShares,
        AllocateParams memory allocateParams,
        SlashingParams memory slashingParams,
        string memory err
    ) internal {
        require(allocateParams.strategies.length == 1 && slashingParams.strategies.length == 1, "only beacon strategy supported");
        require(allocateParams.strategies[0] == BEACONCHAIN_ETH_STRAT, "only beacon strategy supported");
        uint curShares = _getWithdrawableShares(staker, allocateParams.strategies)[0];
        uint prevShares = _getPrevWithdrawableShares(staker, allocateParams.strategies)[0];
        // 1. The withdrawable shares should decrease by a factor of the BCSF
        assertApproxEqAbs(prevShares.mulWad(_getBeaconChainSlashingFactor(staker)), curShares, 1e5, err);
        /**
         * 2. The delta in shares is given by:
         * (originalWithdrawableShares * operatorMag) + extraValidatorShares - (depositShares * operatorMag * BCSF * dsf)
         */
        uint beaconChainSlashingFactor = _getBeaconChainSlashingFactor(staker);
        uint wadToSlash = slashingParams.wadsToSlash[0];
        uint originalAVSSlashedShares = originalWithdrawableShares.mulWadRoundUp(allocateParams.newMagnitudes[0].mulWadRoundUp(wadToSlash));
        uint withdrawableSharesAfterValidatorProven = originalWithdrawableShares - originalAVSSlashedShares + extraValidatorShares;
        uint expectedDelta = withdrawableSharesAfterValidatorProven.mulWad(WAD - beaconChainSlashingFactor);
        assertApproxEqAbs(prevShares - expectedDelta, curShares, 1e5, err);
        /**
         * 3. The attributable avs slashed shares should decrease by a factor of the BCSF
         * Attributable avs slashed shares = depositShares - bcSlashedShares - curShars
         * Where bcSlashedShares = depositShares * (1 - BCSF)
         */
        uint depositShares = _getStakerDepositShares(staker, allocateParams.strategies)[0];
        uint bcSlashedShares = depositShares.mulWad(WAD - beaconChainSlashingFactor);
        uint attributableAVSSlashedShares = depositShares - bcSlashedShares - curShares;
        assertApproxEqAbs(originalAVSSlashedShares.mulWad(beaconChainSlashingFactor), attributableAVSSlashedShares, 1e5, err);
    }
    // TODO: slashable stake

    /**
     *
     *                     SNAPSHOT ASSERTIONS: OPERATOR SHARES
     *
     */

    /// @dev Check that the operator has `addedShares` additional operator shares
    // for each strategy since the last snapshot
    function assert_Snap_Added_OperatorShares(User operator, IStrategy[] memory strategies, uint[] memory addedShares, string memory err)
        internal
    {
        uint[] memory curShares = _getOperatorShares(operator, strategies);
        // Use timewarp to get previous operator shares
        uint[] memory prevShares = _getPrevOperatorShares(operator, strategies);
        // For each strategy, check (prev + added == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] + addedShares[i], curShares[i], err);
        }
    }

    /// @dev Check that the operator has `removedShares` fewer operator shares
    /// for each strategy since the last snapshot
    function assert_Snap_Removed_OperatorShares(
        User operator,
        IStrategy[] memory strategies,
        uint[] memory removedShares,
        string memory err
    ) internal {
        uint[] memory curShares = _getOperatorShares(operator, strategies);
        // Use timewarp to get previous operator shares
        uint[] memory prevShares = _getPrevOperatorShares(operator, strategies);
        // For each strategy, check (prev - removed == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i], curShares[i] + removedShares[i], err);
        }
    }

    /// @dev Check that the operator's shares in ALL strategies have not changed since the last snapshot
    function assert_Snap_Unchanged_OperatorShares(User operator, string memory err) internal {
        IStrategy[] memory strategies = allStrats;
        assertEq(_getOperatorShares(operator, strategies), _getPrevOperatorShares(operator, strategies), err);
    }

    function assert_Snap_Delta_OperatorShares(User operator, IStrategy[] memory strategies, int[] memory shareDeltas, string memory err)
        internal
    {
        uint[] memory curShares = _getOperatorShares(operator, strategies);
        // Use timewarp to get previous operator shares
        uint[] memory prevShares = _getPrevOperatorShares(operator, strategies);
        // For each strategy, check (prev + added == cur)
        for (uint i = 0; i < strategies.length; i++) {
            uint expectedShares;
            if (shareDeltas[i] < 0) expectedShares = prevShares[i] - uint(-shareDeltas[i]);
            else expectedShares = prevShares[i] + uint(shareDeltas[i]);
            assertEq(expectedShares, curShares[i], err);
        }
    }

    function assert_Snap_Slashed_OperatorShares(User operator, SlashingParams memory params, string memory err) internal {
        uint[] memory curShares = _getOperatorShares(operator, params.strategies);
        uint[] memory prevShares = _getPrevOperatorShares(operator, params.strategies);
        Magnitudes[] memory curMagnitudes = _getMagnitudes(operator, params.strategies);
        Magnitudes[] memory prevMagnitudes = _getPrevMagnitudes(operator, params.strategies);
        for (uint i = 0; i < params.strategies.length; i++) {
            // Slashing doesn't occur if the operator has no slashable magnitude
            // This prevents a div by 0 when calculating expected slashed
            uint expectedSlashed = prevMagnitudes[i].max == 0
                ? 0
                : SlashingLib.calcSlashedAmount({
                    operatorShares: prevShares[i],
                    prevMaxMagnitude: prevMagnitudes[i].max,
                    newMaxMagnitude: curMagnitudes[i].max
                });
            assertEq(curShares[i], prevShares[i] - expectedSlashed, err);
        }
    }

    function assert_Snap_Increased_BurnableShares(User operator, SlashingParams memory params, string memory err) internal {
        uint[] memory curBurnable = _getBurnableShares(params.strategies);
        uint[] memory prevBurnable = _getPrevBurnableShares(params.strategies);
        uint[] memory curShares = _getOperatorShares(operator, params.strategies);
        uint[] memory prevShares = _getPrevOperatorShares(operator, params.strategies);
        for (uint i = 0; i < params.strategies.length; i++) {
            uint slashedAtLeast = prevShares[i] - curShares[i];
            // Not factoring in slashable shares in queue here, because that gets more complex (TODO)
            assertTrue(curBurnable[i] >= (prevBurnable[i] + slashedAtLeast), err);
        }
    }

    /**
     *
     *                         SNAPSHOT ASSERTIONS: STAKER SHARES
     *
     */

    /// @dev Check that the staker has `addedShares` additional deposit shares
    /// for each strategy since the last snapshot
    function assert_Snap_Added_Staker_DepositShares(
        User staker,
        IStrategy[] memory strategies,
        uint[] memory addedShares,
        string memory err
    ) internal {
        uint[] memory curShares = _getStakerDepositShares(staker, strategies);
        // Use timewarp to get previous staker shares
        uint[] memory prevShares = _getPrevStakerDepositShares(staker, strategies);
        // For each strategy, check (prev + added == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertApproxEqAbs(prevShares[i] + addedShares[i], curShares[i], 1, err);
        }
    }

    function assert_Snap_Added_Staker_DepositShares(User staker, IStrategy strat, uint addedShares, string memory err) internal {
        assert_Snap_Added_Staker_DepositShares(staker, strat.toArray(), addedShares.toArrayU256(), err);
    }

    /// @dev Check that the staker has `removedShares` fewer delegatable shares
    /// for each strategy since the last snapshot
    function assert_Snap_Removed_Staker_DepositShares(
        User staker,
        IStrategy[] memory strategies,
        uint[] memory removedShares,
        string memory err
    ) internal {
        uint[] memory curShares = _getStakerDepositShares(staker, strategies);
        // Use timewarp to get previous staker shares
        uint[] memory prevShares = _getPrevStakerDepositShares(staker, strategies);
        // For each strategy, check (prev - removed == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] - removedShares[i], curShares[i], err);
        }
    }

    function assert_Snap_Removed_Staker_DepositShares(User staker, IStrategy strat, uint removedShares, string memory err) internal {
        assert_Snap_Removed_Staker_DepositShares(staker, strat.toArray(), removedShares.toArrayU256(), err);
    }

    /// @dev Check that the staker's delegatable shares in ALL strategies have not changed
    /// since the last snapshot
    function assert_Snap_Unchanged_Staker_DepositShares(User staker, string memory err) internal {
        IStrategy[] memory strategies = allStrats;
        assertEq(_getStakerDepositShares(staker, strategies), _getPrevStakerDepositShares(staker, strategies), err);
    }

    /// @dev Check that the staker's withdrawable shares have increased by `addedShares`
    function assert_Snap_Added_Staker_WithdrawableShares(
        User staker,
        IStrategy[] memory strategies,
        uint[] memory addedShares,
        string memory err
    ) internal {
        uint[] memory curShares = _getStakerWithdrawableShares(staker, strategies);
        // Use timewarp to get previous staker shares
        uint[] memory prevShares = _getPrevStakerWithdrawableShares(staker, strategies);
        // For each strategy, check (prev - removed == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] + addedShares[i], curShares[i], err);
        }
    }

    /// @dev This is currently used by dual slashing tests
    /// TODO: potentially bound better
    function assert_Snap_Added_Staker_WithdrawableShares_AtLeast(
        User staker,
        IStrategy[] memory strategies,
        uint[] memory addedShares,
        string memory err
    ) internal {
        uint[] memory curShares = _getStakerWithdrawableShares(staker, strategies);
        // Use timewarp to get previous staker shares
        uint[] memory prevShares = _getPrevStakerWithdrawableShares(staker, strategies);
        // For each strategy, check (prev - removed == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertApproxEqAbs(prevShares[i] + addedShares[i], curShares[i], 1e3, err);
        }
    }

    /// @dev Check that the staker's withdrawable shares have decreased by `removedShares`
    function assert_Snap_Removed_Staker_WithdrawableShares(
        User staker,
        IStrategy[] memory strategies,
        uint[] memory removedShares,
        string memory err
    ) internal {
        uint[] memory curShares = _getStakerWithdrawableShares(staker, strategies);
        // Use timewarp to get previous staker shares
        uint[] memory prevShares = _getPrevStakerWithdrawableShares(staker, strategies);
        // For each strategy, check (prev - removed == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] - removedShares[i], curShares[i], err);
        }
    }

    function assert_Snap_Removed_Staker_WithdrawableShares(User staker, IStrategy strat, uint removedShares, string memory err) internal {
        assert_Snap_Removed_Staker_WithdrawableShares(staker, strat.toArray(), removedShares.toArrayU256(), err);
    }

    function assert_Snap_Unchanged_Staker_WithdrawableShares(User staker, IStrategy[] memory strategies, string memory err) internal {
        assertEq(_getStakerWithdrawableShares(staker, strategies), _getPrevStakerWithdrawableShares(staker, strategies), err);
    }

    /// @dev Check that the staker's withdrawable shares have changed by the expected amount
    function assert_Snap_Expected_Staker_WithdrawableShares_Delegation(
        User staker,
        User operator,
        IStrategy[] memory strategies,
        uint[] memory depositShares,
        string memory err
    ) internal {
        assertEq(
            _getStakerWithdrawableShares(staker, strategies),
            _getExpectedWithdrawableSharesDelegate(staker, operator, strategies, depositShares),
            err
        );
    }

    function assert_Snap_Expected_Staker_WithdrawableShares_Deposit(
        User staker,
        User operator,
        IStrategy[] memory strategies,
        uint[] memory depositSharesAdded,
        string memory err
    ) internal {
        uint[] memory curShares = _getStakerWithdrawableShares(staker, strategies);
        // Use timewarp to get previous staker shares
        uint[] memory prevShares = _getPrevStakerWithdrawableShares(staker, strategies);
        uint[] memory expectedWithdrawableShares = new uint[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            if (prevShares[i] == 0 && depositSharesAdded[i] > 0) {
                expectedWithdrawableShares[i] =
                    _getExpectedWithdrawableSharesDeposit(staker, operator, strategies[i], depositSharesAdded[i]);
                assertEq(curShares[i], expectedWithdrawableShares[i], err);
            } else {
                uint[] memory prevDepositShares = _getPrevStakerDepositShares(staker, strategies);
                assertEq(
                    (prevDepositShares[i] + depositSharesAdded[i]).mulWad(_getDepositScalingFactor(staker, strategies[i])).mulWad(
                        _getSlashingFactor(staker, strategies[i])
                    ),
                    curShares[i],
                    err
                );
            }
        }
    }

    /// @dev Check that the staker's withdrawable shares have decreased by at least `removedShares`
    /// @dev Used to handle overslashing of beacon chain
    function assert_Snap_Removed_Staker_WithdrawableShares_AtLeast(
        User staker,
        IStrategy[] memory strategies,
        uint[] memory removedShares,
        string memory err
    ) internal {
        uint[] memory curShares = _getStakerWithdrawableShares(staker, strategies);
        // Use timewarp to get previous staker shares
        uint[] memory prevShares = _getPrevStakerWithdrawableShares(staker, strategies);
        // Assert that the decrease in withdrawable shares is at least as much as the removed shares
        // Checking for expected rounding down behavior
        for (uint i = 0; i < strategies.length; i++) {
            assertGe(prevShares[i] - curShares[i], removedShares[i], err);
        }
    }

    /// @dev Check that the staker's withdrawable shares have decreased by at least `removedShares`
    /// @dev Used to handle overslashing of beacon chain with AVS slashings
    function assert_Snap_Removed_Staker_WithdrawableShares_AtLeast(
        User staker,
        IStrategy[] memory strategies,
        uint[] memory removedShares,
        uint errBound,
        string memory err
    ) internal {
        uint[] memory curShares = _getStakerWithdrawableShares(staker, strategies);
        // Use timewarp to get previous staker shares
        uint[] memory prevShares = _getPrevStakerWithdrawableShares(staker, strategies);
        // For each strategy, check diff between (prev-removed) and curr is at most 1 gwei
        for (uint i = 0; i < strategies.length; i++) {
            assertApproxEqAbs(prevShares[i] - removedShares[i], curShares[i], errBound, err);
        }
    }

    function assert_Snap_Removed_Staker_WithdrawableShares_AtLeast(User staker, IStrategy strat, uint removedShares, string memory err)
        internal
    {
        assert_Snap_Removed_Staker_WithdrawableShares_AtLeast(staker, strat.toArray(), removedShares.toArrayU256(), err);
    }

    function assert_Snap_Delta_StakerShares(User staker, IStrategy[] memory strategies, int[] memory shareDeltas, string memory err)
        internal
    {
        int[] memory curShares = _getStakerDepositSharesInt(staker, strategies);
        // Use timewarp to get previous staker shares
        int[] memory prevShares = _getPrevStakerDepositSharesInt(staker, strategies);
        // For each strategy, check (prev + added == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] + shareDeltas[i], curShares[i], err);
        }
    }

    function assert_Snap_Unchanged_DSF(User staker, IStrategy[] memory strategies, string memory err) internal {
        assertEq(_getDepositScalingFactors(staker, strategies), _getPrevDepositScalingFactors(staker, strategies), err);
    }

    function assert_Snap_Increased_DSF(User staker, IStrategy[] memory strategies, string memory err) internal {
        uint[] memory curDSFs = _getDepositScalingFactors(staker, strategies);
        uint[] memory prevDSFs = _getPrevDepositScalingFactors(staker, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertGt(curDSFs[i], prevDSFs[i], err);
        }
    }

    /// @dev Used to assert that the DSF is either increased or unchanged, depending on the slashing factor, on a deposit
    function assert_Snap_DSF_State_Deposit(User staker, IStrategy[] memory strategies, string memory err) internal {
        uint[] memory curDepositShares = _getStakerDepositShares(staker, strategies);
        uint[] memory prevDepositShares = _getPrevStakerDepositShares(staker, strategies);
        uint[] memory curDSFs = _getDepositScalingFactors(staker, strategies);
        uint[] memory prevDSFs = _getPrevDepositScalingFactors(staker, strategies);
        uint[] memory curSlashingFactors = _getSlashingFactors(staker, strategies);
        uint[] memory prevSlashingFactors = _getPrevSlashingFactors(staker, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            // If there was never a slashing, no need to normalize
            if (curSlashingFactors[i] == WAD) {
                assertEq(prevDSFs[i], curDSFs[i], err); // No slashing, so DSF is unchanged
                assertEq(curDSFs[i], WAD, err); // DSF should have always been WAD
            }
            // If there was a slashing, and we deposit, normalize
            else {
                // If the DSF and slashing factor are already normalized against each other from a previous deposit (prevWithdrawableFactor very close to WAD)
                // and there have been no subsequent slashings, DSF "should" stay the same, but recomputing decreases it slightly due to
                // fixed point arithmetic rounding. Outer if is to prevent int underflow errors.
                uint prevWithdrawableFactor = prevDSFs[i].mulWad(prevSlashingFactors[i]);
                require(WAD >= prevWithdrawableFactor, "withdrawableFactor should always be less than or equal to WAD");
                if (WAD - prevWithdrawableFactor < 1e2 && prevDepositShares[i] > 0 && curSlashingFactors[i] == prevSlashingFactors[i]) {
                    assertApproxEqAbs(curDSFs[i], prevDSFs[i], 1e2, err);
                } else {
                    assertGt(curDSFs[i], prevDSFs[i], err); // Slashing, so DSF is increased
                }
            }
        }
    }

    /// @dev When completing withdrawals as shares, we must also handle the case where a staker completes a withdrawal for 0 shares
    function assert_Snap_DSF_State_WithdrawalAsShares(User staker, IStrategy[] memory strategies, string memory err) internal {
        uint[] memory curDepositShares = _getStakerDepositShares(staker, strategies);
        uint[] memory prevDepositShares = _getPrevStakerDepositShares(staker, strategies);
        uint[] memory curDSFs = _getDepositScalingFactors(staker, strategies);
        uint[] memory prevDSFs = _getPrevDepositScalingFactors(staker, strategies);
        uint[] memory curSlashingFactors = _getSlashingFactors(staker, strategies);
        uint[] memory prevSlashingFactors = _getPrevSlashingFactors(staker, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            // If there was never a slashing, no need to normalize
            // If there was a slashing, but we complete a withdrawal for 0 shares, no need to normalize
            if (curSlashingFactors[i] == WAD || curDepositShares[i] == 0) {
                assertEq(prevDSFs[i], curDSFs[i], err);
                assertEq(curDSFs[i], WAD, err);
            }
            // If the staker has a slashingFactor of 0, any withdrawal as shares won't change the DSF
            else if (curSlashingFactors[i] == 0) {
                assertEq(prevDSFs[i], curDSFs[i], err);
            }
            // If there was a slashing and we complete a withdrawal for non-zero shares, normalize the DSF
            else {
                // If the DSF and slashing factor are already normalized against each other from a previous deposit (prevWithdrawableFactor very close to WAD)
                // and there have been no subsequent slashings, DSF "should" stay the same, but recomputing decreases it slightly due to
                // fixed point arithmetic rounding. Outer if is to prevent int underflow errors.
                uint prevWithdrawableFactor = prevDSFs[i].mulWad(prevSlashingFactors[i]);
                require(WAD >= prevWithdrawableFactor, "withdrawableFactor should always be less than or equal to WAD");
                if (WAD - prevWithdrawableFactor < 1e2 && prevDepositShares[i] > 0 && curSlashingFactors[i] == prevSlashingFactors[i]) {
                    assertApproxEqAbs(curDSFs[i], prevDSFs[i], 1e2, err);
                } else {
                    assertGt(curDSFs[i], prevDSFs[i], err); // Slashing, so DSF is increased
                }
            }
        }
    }

    /// @dev On a delegation, the DSF should be increased if the operator magnitude is non-WAD
    function assert_Snap_DSF_State_Delegation(
        User staker,
        IStrategy[] memory strategies,
        uint[] memory delegatableShares,
        string memory err
    ) internal {
        address operator = delegationManager.delegatedTo(address(staker));
        uint64[] memory maxMags = _getMaxMagnitudes(User(payable(operator)), strategies);
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy[] memory stratArray = strategies[i].toArray();
            // If you are delegating with 0 shares, no need to normalize
            // If there was never an operator slashing, no need to normalize
            if (delegatableShares[i] == 0 || maxMags[i] == WAD) {
                assert_Snap_Unchanged_DSF(staker, stratArray, err);
                // If we are not a beaconChainStrategy, we should also have a DSF of WAD
                // We exclude BEACONCHAIN_ETH_STRAT because it could have had a non-WAD DSF from BC slashings
                if (strategies[i] != BEACONCHAIN_ETH_STRAT) assert_DSF_WAD(staker, stratArray, err);
            }
            // If there was an operator slashing, and delegating with non-zero shares, normalize
            else {
                assert_Snap_Increased_DSF(staker, stratArray, err); // Slashing, so DSF is increased
            }
        }
    }

    /**
     *
     *                     SNAPSHOT ASSERTIONS: STRATEGY SHARES
     *
     */
    function assert_Snap_Removed_StrategyShares(IStrategy[] memory strategies, uint[] memory removedShares, string memory err) internal {
        uint[] memory curShares = _getTotalStrategyShares(strategies);
        uint[] memory prevShares = _getPrevTotalStrategyShares(strategies);
        for (uint i = 0; i < strategies.length; i++) {
            // Ignore BeaconChainETH strategy since it doesn't keep track of global strategy shares
            if (strategies[i] == BEACONCHAIN_ETH_STRAT) continue;
            assertApproxEqAbs(prevShares[i] - removedShares[i], curShares[i], 1, err);
        }
    }

    function assert_Snap_Unchanged_StrategyShares(IStrategy[] memory strategies, string memory err) internal {
        assertEq(_getTotalStrategyShares(strategies), _getPrevTotalStrategyShares(strategies), err);
    }

    function assert_SlashableStake_Decrease_BCSlash(User staker) internal {
        if (delegationManager.isDelegated(address(staker))) {
            address operator = delegationManager.delegatedTo(address(staker));
            (OperatorSet[] memory operatorSets, Allocation[] memory allocations) = _getStrategyAllocations(operator, BEACONCHAIN_ETH_STRAT);
            for (uint i = 0; i < operatorSets.length; i++) {
                if (allocations[i].currentMagnitude > 0) {
                    assert_Snap_StakeBecomeUnslashable(
                        User(payable(operator)),
                        operatorSets[i],
                        BEACONCHAIN_ETH_STRAT.toArray(),
                        "operator should have minSlashableStake decreased"
                    );
                }
            }
        }
    }

    /**
     *
     *                   SNAPSHOT ASSERTIONS: UNDERLYING TOKEN
     *
     */

    /// @dev Check that the staker has `addedTokens` additional underlying tokens since the last snapshot
    function assert_Snap_Added_TokenBalances(User staker, IERC20[] memory tokens, uint[] memory addedTokens, string memory err) internal {
        uint[] memory curTokenBalances = _getTokenBalances(staker, tokens);
        uint[] memory prevTokenBalances = _getPrevTokenBalances(staker, tokens);
        for (uint i = 0; i < tokens.length; i++) {
            assertApproxEqAbs(prevTokenBalances[i] + addedTokens[i], curTokenBalances[i], 1, err);
        }
    }

    function assert_Snap_Added_TokenBalances(User staker, IStrategy[] memory strategies, uint[] memory addedTokens, string memory err)
        internal
    {
        IERC20[] memory tokens = _getUnderlyingTokens(strategies);
        uint[] memory curTokenBalances = _getTokenBalances(staker, tokens);
        uint[] memory prevTokenBalances = _getPrevTokenBalances(staker, tokens);
        for (uint i = 0; i < tokens.length; i++) {
            assertApproxEqAbs(prevTokenBalances[i] + addedTokens[i], curTokenBalances[i], 1, err);
        }
    }

    /// @dev Check that the staker's underlying token balance for ALL tokens have not changed since the last snapshot
    function assert_Snap_Unchanged_TokenBalances(User staker, string memory err) internal {
        IERC20[] memory tokens = allTokens;
        assertEq(_getTokenBalances(staker, tokens), _getPrevTokenBalances(staker, tokens), err);
    }

    /**
     *
     *                   SNAPSHOT ASSERTIONS: QUEUED WITHDRAWALS
     *
     */
    function assert_Snap_Added_QueuedWithdrawals(User staker, Withdrawal[] memory withdrawals, string memory err) internal {
        assertEq(_getPrevCumulativeWithdrawals(staker) + withdrawals.length, _getCumulativeWithdrawals(staker), err);
    }

    function assert_Snap_Added_QueuedWithdrawal(User staker, Withdrawal memory, /*withdrawal*/ string memory err) internal {
        assertEq(_getPrevCumulativeWithdrawals(staker) + 1, _getCumulativeWithdrawals(staker), err);
    }
    /**
     *
     *                      SNAPSHOT ASSERTIONS: EIGENPODS
     *
     */

    function assert_Snap_Added_ActiveValidatorCount(User staker, uint addedValidators, string memory err) internal {
        assertEq(_getPrevActiveValidatorCount(staker) + addedValidators, _getActiveValidatorCount(staker), err);
    }

    function assert_Snap_Removed_ActiveValidatorCount(User staker, uint exitedValidators, string memory err) internal {
        assertEq(_getActiveValidatorCount(staker) + exitedValidators, _getPrevActiveValidatorCount(staker), err);
    }

    function assert_Snap_Unchanged_ActiveValidatorCount(User staker, string memory err) internal {
        assertEq(_getActiveValidatorCount(staker), _getPrevActiveValidatorCount(staker), err);
    }

    function assert_Snap_Added_ActiveValidators(User staker, uint40[] memory addedValidators, string memory err) internal {
        bytes32[] memory pubkeyHashes = beaconChain.getPubkeyHashes(addedValidators);
        VALIDATOR_STATUS[] memory curStatuses = _getValidatorStatuses(staker, pubkeyHashes);
        VALIDATOR_STATUS[] memory prevStatuses = _getPrevValidatorStatuses(staker, pubkeyHashes);
        for (uint i = 0; i < curStatuses.length; i++) {
            assertTrue(prevStatuses[i] == VALIDATOR_STATUS.INACTIVE, err);
            assertTrue(curStatuses[i] == VALIDATOR_STATUS.ACTIVE, err);
        }
    }

    function assert_Snap_Removed_ActiveValidators(User staker, uint40[] memory exitedValidators, string memory err) internal {
        bytes32[] memory pubkeyHashes = beaconChain.getPubkeyHashes(exitedValidators);
        VALIDATOR_STATUS[] memory curStatuses = _getValidatorStatuses(staker, pubkeyHashes);
        VALIDATOR_STATUS[] memory prevStatuses = _getPrevValidatorStatuses(staker, pubkeyHashes);
        for (uint i = 0; i < curStatuses.length; i++) {
            assertTrue(prevStatuses[i] == VALIDATOR_STATUS.ACTIVE, err);
            assertTrue(curStatuses[i] == VALIDATOR_STATUS.WITHDRAWN, err);
        }
    }

    function assert_Snap_Created_Checkpoint(User staker, string memory err) internal {
        assertEq(_getPrevCheckpointTimestamp(staker), 0, err);
        assertTrue(_getCheckpointTimestamp(staker) != 0, err);
    }

    function assert_Snap_Removed_Checkpoint(User staker, string memory err) internal {
        assertEq(_getCheckpointTimestamp(staker), 0, err);
        assertTrue(_getPrevCheckpointTimestamp(staker) != 0, err);
    }

    function assert_Snap_Unchanged_Checkpoint(User staker, string memory err) internal {
        assertEq(_getCheckpointTimestamp(staker), _getPrevCheckpointTimestamp(staker), err);
    }

    function assert_Snap_Updated_LastCheckpoint(User staker, string memory err) internal {
        assertTrue(_getLastCheckpointTimestamp(staker) > _getPrevLastCheckpointTimestamp(staker), err);
    }

    function assert_Snap_Added_PodBalanceToWithdrawable(User staker, string memory err) internal {
        assertEq(
            _getPrevWithdrawableRestakedGwei(staker) + _getPrevCheckpointPodBalanceGwei(staker), _getWithdrawableRestakedGwei(staker), err
        );
    }

    function assert_Snap_Added_WithdrawableGwei(User staker, uint64 addedGwei, string memory err) internal {
        assertEq(_getPrevWithdrawableRestakedGwei(staker) + addedGwei, _getWithdrawableRestakedGwei(staker), err);
    }

    function assert_Snap_Added_BalanceExitedGwei(User staker, uint64 addedGwei, string memory err) internal {
        uint64 curCheckpointTimestamp = _getCheckpointTimestamp(staker);
        uint64 prevCheckpointTimestamp = _getPrevCheckpointTimestamp(staker);
        // If we just finalized a checkpoint, that's the timestamp we want to use to look up checkpoint balances exited
        uint64 targetTimestamp = curCheckpointTimestamp;
        if (curCheckpointTimestamp != prevCheckpointTimestamp) targetTimestamp = prevCheckpointTimestamp;
        assertEq(
            _getPrevCheckpointBalanceExited(staker, targetTimestamp) + addedGwei, _getCheckpointBalanceExited(staker, targetTimestamp), err
        );
    }

    function assert_Snap_Decreased_BCSF(User staker, string memory err) internal {
        assertLt(_getBeaconChainSlashingFactor(staker), _getPrevBeaconChainSlashingFactor(staker), err);
    }

    function assert_Snap_Unchanged_BCSF(User staker, string memory err) internal {
        assertEq(_getBeaconChainSlashingFactor(staker), _getPrevBeaconChainSlashingFactor(staker), err);
    }
}
