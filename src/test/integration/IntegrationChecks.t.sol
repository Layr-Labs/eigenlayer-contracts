// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationBase.t.sol";
import "src/test/integration/users/User.t.sol";
import "src/test/integration/users/User_M1.t.sol";
import "src/test/integration/users/User_M2.t.sol";

/// @notice Contract that provides utility functions to reuse common test blocks & checks
contract IntegrationCheckUtils is IntegrationBase {
    using ArrayLib for IStrategy[];
    using SlashingLib for *;

    /*******************************************************************************
                                 EIGENPOD CHECKS
    *******************************************************************************/

    function check_VerifyWC_State(
        User staker,
        uint40[] memory validators,
        uint64 beaconBalanceGwei
    ) internal {
        uint beaconBalanceWei = beaconBalanceGwei * GWEI_TO_WEI;
        assert_Snap_Added_Staker_DepositShares(staker, BEACONCHAIN_ETH_STRAT, beaconBalanceWei, "staker should have added deposit shares to beacon chain strat");
        assert_Snap_Added_ActiveValidatorCount(staker, validators.length, "staker should have increased active validator count");
        assert_Snap_Added_ActiveValidators(staker, validators, "validators should each be active");
    }

    function check_StartCheckpoint_State(
        User staker
    ) internal {
        assert_ProofsRemainingEqualsActive(staker, "checkpoint proofs remaining should equal active validator count");
        assert_Snap_Created_Checkpoint(staker, "staker should have created a new checkpoint");
    }

    function check_StartCheckpoint_WithPodBalance_State(
        User staker,
        uint64 expectedPodBalanceGwei
    ) internal {
        check_StartCheckpoint_State(staker);

        assert_CheckpointPodBalance(staker, expectedPodBalanceGwei, "checkpoint podBalanceGwei should equal expected");
    }

    function check_StartCheckpoint_NoValidators_State(
        User staker,
        uint64 sharesAddedGwei
    ) internal {
        assert_Snap_Added_Staker_DepositShares(staker, BEACONCHAIN_ETH_STRAT, sharesAddedGwei * GWEI_TO_WEI, "should have added staker shares");
        assert_Snap_Added_WithdrawableGwei(staker, sharesAddedGwei, "should have added to withdrawable restaked gwei");
        
        assert_Snap_Unchanged_ActiveValidatorCount(staker, "active validator count should remain 0");
        assert_Snap_Updated_LastCheckpoint(staker, "last checkpoint timestamp should have increased");
        assert_Snap_Unchanged_Checkpoint(staker, "current checkpoint timestamp should be unchanged");
    }

    function check_CompleteCheckpoint_State(
        User staker
    ) internal {
        assert_Snap_Removed_Checkpoint(staker, "should have deleted active checkpoint");
        assert_Snap_Updated_LastCheckpoint(staker, "last checkpoint timestamp should be updated");
        assert_Snap_Added_PodBalanceToWithdrawable(staker, "pod balance should have been added to withdrawable restaked exec layer gwei");
    }

    function check_CompleteCheckpoint_EarnOnBeacon_State(
        User staker,
        uint64 beaconBalanceAdded
    ) internal {
        check_CompleteCheckpoint_State(staker);

        uint balanceAddedWei = beaconBalanceAdded * GWEI_TO_WEI;
        assert_Snap_Added_Staker_DepositShares(staker, BEACONCHAIN_ETH_STRAT, balanceAddedWei, "should have increased shares by excess beacon balance");
    }

    function check_CompleteCheckpoint_WithPodBalance_State(
        User staker,
        uint64 expectedPodBalanceGwei
    ) internal {
        check_CompleteCheckpoint_State(staker);

        assert_Snap_Added_WithdrawableGwei(staker, expectedPodBalanceGwei, "should have added expected gwei to withdrawable restaked exec layer gwei");
    }

    function check_CompleteCheckpoint_WithSlashing_State(
        User staker,
        uint40[] memory slashedValidators,
        uint64 slashedAmountGwei
    ) internal {
        check_CompleteCheckpoint_State(staker);

        assert_Snap_Unchanged_Staker_DepositShares(staker, "staker shares should not have decreased");
        assert_Snap_Removed_Staker_WithdrawableShares(staker, BEACONCHAIN_ETH_STRAT, slashedAmountGwei * GWEI_TO_WEI, "should have decreased withdrawable shares by slashed amount");
        assert_Snap_Removed_ActiveValidatorCount(staker, slashedValidators.length, "should have decreased active validator count");
        assert_Snap_Removed_ActiveValidators(staker, slashedValidators, "exited validators should each be WITHDRAWN");
    }

    function check_CompleteCheckpoint_WithSlashing_HandleRoundDown_State(
        User staker,
        uint40[] memory slashedValidators,
        uint64 slashedAmountGwei
    ) internal {
        check_CompleteCheckpoint_State(staker);

        assert_Snap_Unchanged_Staker_DepositShares(staker, "staker shares should not have decreased");
        assert_Snap_Removed_Staker_WithdrawableShares_AtLeast(staker, BEACONCHAIN_ETH_STRAT, slashedAmountGwei * GWEI_TO_WEI, "should have decreased withdrawable shares by at least slashed amount");
        assert_Snap_Removed_ActiveValidatorCount(staker, slashedValidators.length, "should have decreased active validator count");
        assert_Snap_Removed_ActiveValidators(staker, slashedValidators, "exited validators should each be WITHDRAWN");
    }

    function check_CompleteCheckpoint_WithCLSlashing_HandleRoundDown_State(
        User staker,
        uint64 slashedAmountGwei
    ) internal {
        check_CompleteCheckpoint_State(staker);

        assert_Snap_Unchanged_Staker_DepositShares(staker, "staker shares should not have decreased");
        assert_Snap_Removed_Staker_WithdrawableShares_AtLeast(staker, BEACONCHAIN_ETH_STRAT, slashedAmountGwei * GWEI_TO_WEI, "should have decreased withdrawable shares by at least slashed amount");
        assert_Snap_Unchanged_ActiveValidatorCount(staker, "should not have changed active validator count");
    }

    function check_CompleteCheckpoint_WithCLSlashing_State(
        User staker,
        uint64 slashedAmountGwei
    ) internal {
        check_CompleteCheckpoint_State(staker);

        assert_Snap_Unchanged_Staker_DepositShares(staker, "staker shares should not have decreased");
        assert_Snap_Removed_Staker_WithdrawableShares(staker, BEACONCHAIN_ETH_STRAT, slashedAmountGwei * GWEI_TO_WEI, "should have decreased withdrawable shares by slashed amount");
        assert_Snap_Unchanged_ActiveValidatorCount(staker, "should not have changed active validator count");
    }

    function check_CompleteCheckpoint_WithExits_State(
        User staker,
        uint40[] memory exitedValidators,
        uint64 exitedBalanceGwei
    ) internal {
        check_CompleteCheckpoint_WithPodBalance_State(staker, exitedBalanceGwei);

        assert_Snap_Unchanged_Staker_DepositShares(staker, "staker should not have changed shares");
        assert_Snap_Added_BalanceExitedGwei(staker, exitedBalanceGwei, "should have attributed expected gwei to exited balance");
        assert_Snap_Removed_ActiveValidatorCount(staker, exitedValidators.length, "should have decreased active validator count");
        assert_Snap_Removed_ActiveValidators(staker, exitedValidators, "exited validators should each be WITHDRAWN");
    }

    /*******************************************************************************
                              LST/DELEGATION CHECKS
    *******************************************************************************/

    function check_Deposit_State(
        User staker, 
        IStrategy[] memory strategies, 
        uint[] memory shares
    ) internal {
        /// Deposit into strategies:
        // For each of the assets held by the staker (either StrategyManager or EigenPodManager),
        // the staker calls the relevant deposit function, depositing all held assets.
        //
        // ... check that all underlying tokens were transferred to the correct destination
        //     and that the staker now has the expected amount of delegated shares in each strategy
        assert_HasNoUnderlyingTokenBalance(staker, strategies, "staker should have transferred all underlying tokens");
        assert_Snap_Added_Staker_DepositShares(staker, strategies, shares, "staker should expect shares in each strategy after depositing");
        assert_Snap_Added_Staker_WithdrawableShares(staker, strategies, shares, "deposit should increase withdrawable shares");
    }

    function check_Deposit_State_PartialDeposit(User staker, IStrategy[] memory strategies, uint[] memory shares, uint[] memory tokenBalances) internal {
        /// Deposit into strategies:
        // For each of the assets held by the staker (either StrategyManager or EigenPodManager),
        // the staker calls the relevant deposit function, depositing some subset of held assets
        //
        // ... check that some underlying tokens were transferred to the correct destination
        //     and that the staker now has the expected amount of delegated shares in each strategy
        assert_HasUnderlyingTokenBalances(staker, strategies, tokenBalances, "staker should have transferred some underlying tokens");
        assert_Snap_Added_Staker_DepositShares(staker, strategies, shares, "staker should expected shares in each strategy after depositing");
        assert_Snap_Added_Staker_WithdrawableShares(staker, strategies, shares, "deposit should increase withdrawable shares");
    }

    function check_Delegation_State(
        User staker, 
        User operator, 
        IStrategy[] memory strategies, 
        uint[] memory shares
    ) internal {
        /// Delegate to an operator:
        //
        // ... check that the staker is now delegated to the operator, and that the operator
        //     was awarded the staker shares
        assertTrue(delegationManager.isDelegated(address(staker)), "staker should be delegated");
        assertEq(address(operator), delegationManager.delegatedTo(address(staker)), "staker should be delegated to operator");
        assert_HasExpectedShares(staker, strategies, shares, "staker should still have expected shares after delegating");
        assert_Snap_Unchanged_Staker_DepositShares(staker, "staker shares should be unchanged after delegating");
        assert_Snap_Unchanged_Staker_WithdrawableShares(staker, "withdrawable shares should be unchanged after delegating");
        assert_Snap_Added_OperatorShares(operator, strategies, shares, "operator should have received shares");
    }

    function check_QueuedWithdrawal_State(
        User staker, 
        User operator, 
        IStrategy[] memory strategies, 
        uint[] memory shares, 
        IDelegationManagerTypes.Withdrawal[] memory withdrawals, 
        bytes32[] memory withdrawalRoots
    ) internal {
        // The staker will queue one or more withdrawals for the selected strategies and shares
        //
        // ... check that each withdrawal was successfully enqueued, that the returned roots
        //     match the hashes of each withdrawal, and that the staker and operator have
        //     reduced shares.
        assertEq(withdrawalRoots.length, 1, "check_QueuedWithdrawal_State: should only have 1 withdrawal root after queueing"); 
        assert_AllWithdrawalsPending(withdrawalRoots,
            "check_QueuedWithdrawal_State: staker withdrawals should now be pending");
        assert_ValidWithdrawalHashes(withdrawals, withdrawalRoots,
            "check_QueuedWithdrawal_State: calculated withdrawals should match returned roots");
        assert_Snap_Added_QueuedWithdrawals(staker, withdrawals,
            "check_QueuedWithdrawal_State: staker should have increased nonce by withdrawals.length");
        assert_Snap_Removed_OperatorShares(operator, strategies, shares,
            "check_QueuedWithdrawal_State: failed to remove operator shares");
        assert_Snap_Removed_Staker_DepositShares(staker, strategies, shares,
            "check_QueuedWithdrawal_State: failed to remove staker shares");
        assert_Snap_Removed_Staker_WithdrawableShares(staker, strategies, shares,
            "check_QueuedWithdrawal_State: failed to remove staker withdrawable shares");
    }

    function check_Undelegate_State(
        User staker, 
        User operator, 
        IDelegationManagerTypes.Withdrawal[] memory withdrawals,
        bytes32[] memory withdrawalRoots,
        IStrategy[] memory strategies,
        uint[] memory shares 
    ) internal {
        /// Undelegate from an operator
        //
        // ... check that the staker is undelegated, all strategies from which the staker is deposited are unqeuued,
        //     that the returned root matches the hashes for each strategy and share amounts, and that the staker
        //     and operator have reduced shares
        assertFalse(delegationManager.isDelegated(address(staker)),
            "check_Undelegate_State: staker should not be delegated");
        assert_ValidWithdrawalHashes(withdrawals, withdrawalRoots,
            "check_Undelegate_State: calculated withdrawl should match returned root");
        assert_AllWithdrawalsPending(withdrawalRoots,
            "check_Undelegate_State: stakers withdrawal should now be pending");
        assert_Snap_Added_QueuedWithdrawals(staker, withdrawals,
            "check_Undelegate_State: staker should have increased nonce by withdrawals.length");
        assert_Snap_Removed_OperatorShares(operator, strategies, shares,
            "check_Undelegate_State: failed to remove operator shares");
        assert_Snap_Removed_Staker_DepositShares(staker, strategies, shares,
            "check_Undelegate_State: failed to remove staker shares");
        assert_Snap_Removed_Staker_WithdrawableShares(staker, strategies, shares,
            "check_QueuedWithdrawal_State: failed to remove staker withdrawable shares");
    }

    /**
     * @notice Overloaded function to check the state after a withdrawal as tokens, accepting a non-user type for the operator.
     * @param staker The staker who completed the withdrawal.
     * @param operator The operator address, which can be a non-user type like address(0).
     * @param withdrawal The details of the withdrawal that was completed.
     * @param strategies The strategies from which the withdrawal was made.
     * @param shares The number of shares involved in the withdrawal.
     * @param tokens The tokens received after the withdrawal.
     * @param expectedTokens The expected tokens to be received after the withdrawal.
     */
    function check_Withdrawal_AsTokens_State(
        User staker,
        User operator,
        IDelegationManagerTypes.Withdrawal memory withdrawal,
        IStrategy[] memory strategies,
        uint[] memory shares,
        IERC20[] memory tokens,
        uint[] memory expectedTokens
    ) internal {
        // Common checks
        assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawal), "staker withdrawal should no longer be pending");
        
        assert_Snap_Added_TokenBalances(staker, tokens, expectedTokens, "staker should have received expected tokens");
        assert_Snap_Unchanged_Staker_DepositShares(staker, "staker shares should not have changed");
        assert_Snap_Removed_StrategyShares(strategies, shares, "strategies should have total shares decremented");

        // Checks specific to an operator that the Staker has delegated to
        if (operator != User(payable(0))) {
            if (operator != staker) {
                assert_Snap_Unchanged_TokenBalances(operator, "operator token balances should not have changed");
            }
            assert_Snap_Unchanged_OperatorShares(operator, "operator shares should not have changed");
        }
    }

    function check_Withdrawal_AsShares_State(
        User staker,
        User operator,
        IDelegationManagerTypes.Withdrawal memory withdrawal,
        IStrategy[] memory strategies,
        uint[] memory shares
    ) internal {
        // Common checks applicable to both user and non-user operator types
        assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawal), "staker withdrawal should no longer be pending");
        assert_Snap_Unchanged_TokenBalances(staker, "staker should not have any change in underlying token balances");
        assert_Snap_Added_Staker_DepositShares(staker, strategies, shares, "staker should have received expected shares");
        assert_Snap_Unchanged_StrategyShares(strategies, "strategies should have total shares unchanged");

        // Additional checks or handling for the non-user operator scenario
        if (operator != User(User(payable(0)))) {
            if (operator != staker) {
                assert_Snap_Unchanged_TokenBalances(operator, "operator should not have any change in underlying token balances");
            }
            assert_Snap_Added_OperatorShares(operator, withdrawal.strategies, withdrawal.scaledShares, "operator should have received shares");
        }
    }

    /// @notice Difference from above is that operator shares do not increase since staker is not delegated
    function check_Withdrawal_AsShares_Undelegated_State(
        User staker,
        User operator,
        IDelegationManagerTypes.Withdrawal memory withdrawal,
        IStrategy[] memory strategies,
        uint[] memory shares
    ) internal {
        /// Complete withdrawal(s):
        // The staker will complete the withdrawal as shares
        // 
        // ... check that the withdrawal is not pending, that the token balances of the staker and operator are unchanged,
        //     that the withdrawer received the expected shares, and that that the total shares of each o
        //     strategy withdrawn remains unchanged 
        assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawal), "staker withdrawal should no longer be pending");
        assert_Snap_Unchanged_TokenBalances(staker, "staker should not have any change in underlying token balances");
        assert_Snap_Unchanged_TokenBalances(operator, "operator should not have any change in underlying token balances");
        assert_Snap_Added_Staker_DepositShares(staker, strategies, shares, "staker should have received expected shares");
        assert_Snap_Unchanged_OperatorShares(operator, "operator should have shares unchanged");
        assert_Snap_Unchanged_StrategyShares(strategies, "strategies should have total shares unchanged");
    }

    /*******************************************************************************
                                 ALLOCATION MANAGER CHECKS
    *******************************************************************************/
    
    /// @dev Basic invariants that should hold after EVERY call to `registerForOperatorSet`
    function check_Base_Registration_State(
        User operator,
        OperatorSet memory operatorSet
    ) internal {
        // Global invariant
        assert_MaxMagsEqualMaxMagsAtCurrentBlock(operator, allStrats, "max magnitudes should equal upperlookup at current block");

        // Registration SHOULD register the operator, making them slashable and adding them as a member of the set
        assert_Snap_Became_Registered(operator, operatorSet, "operator should not have been registered before, and is now registered");
        assert_Snap_Became_Slashable(operator, operatorSet, "operator should not have been slashable before, and is now slashable");
        assert_Snap_Added_RegisteredSet(operator, operatorSet, "should have added operator sets to list of registered sets");
        assert_Snap_Added_MemberOfSet(operator, operatorSet, "should have added operator to list of set members");

        // Registration should NOT change anything about magnitude, allocations, or allocated sets
        assert_Snap_Unchanged_AllocatedSets(operator, "should not have updated allocated sets");
        assert_Snap_Unchanged_AllocatedStrats(operator, operatorSet, "should not have updated allocated strategies");
        assert_Snap_Unchanged_MaxMagnitude(operator, allStrats, "should not have updated max magnitudes in any way");
        assert_Snap_Unchanged_AllocatedStake(operator, operatorSet, allStrats, "should not have updated allocated stake in any way");
        assert_Snap_Unchanged_StrategyAllocations(operator, operatorSet, allStrats, "should not have updated any individual allocations");
        assert_Snap_Unchanged_EncumberedMagnitude(operator, allStrats, "should not have updated encumbered magnitude");
        assert_Snap_Unchanged_AllocatableMagnitude(operator, allStrats, "should not have updated allocatable magnitude");
    }

    // /// @dev Checks invariants for registration for a variety of allocation states
    // /// 
    // function check_Registration_State(
    //     User operator,
    //     OperatorSet memory operatorSet,
    //     IStrategy[] memory unallocated,
    //     IAllocationManagerTypes.AllocateParams memory pending,
    //     IAllocationManagerTypes.AllocateParams memory active
    // ) internal {
    //     check_Base_Registration_State(operator, operatorSet);
    // }

    /// @dev Check invariants for registerForOperatorSets given a set of strategies
    /// for which NO allocation exists (currentMag/pendingDiff are 0)
    /// @param unallocated For the given operatorSet, a list of strategies for which NO allocation exists
    function check_Registration_State_NoAllocation(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory unallocated
    ) internal {
        check_Base_Registration_State(operator, operatorSet);

        /// The operator is NOT allocated, ensure their slashable stake and magnitudes are unchanged
        IAllocationManagerTypes.AllocateParams memory params = IAllocationManagerTypes.AllocateParams({
            operatorSet: operatorSet,
            strategies: unallocated,
            newMagnitudes: new uint64[](unallocated.length)
        });
        assert_CurrentMagnitude(operator, params, "operator should have 0 allocation for each strategy");

        // TODO - actually this would break if we did a full deallocation and checked this, because allocatedStrategies isn't updated
        // assert_IsNotAllocated(operator, operatorSet, unallocated, "operator should not have any allocations from given strategies");
        assert_Snap_Unchanged_SlashableStake(operator, operatorSet, unallocated, "operator should not have increased slashable stake for any given strategy");
    }

    /// @dev Check invariants for registerForOperatorSets AFTER a prior allocation becomes active
    /// @param active allocation params to the last call to modifyAllocations
    ///
    /// ASSUMES:
    /// - the effect block for `params` has already passed
    /// - params.newMagnitudes does NOT contain any `0` entries
    function check_Registration_State_ActiveAllocation(
        User operator,
        IAllocationManagerTypes.AllocateParams memory active
    ) internal {
        OperatorSet memory operatorSet = active.operatorSet;

        /// Basic registerForOperatorSets invariants
        check_Base_Registration_State(operator, operatorSet);

        /// Given an active allocation, check that the allocation is reflected in state
        assert_IsAllocatedToSet(operator, operatorSet, "operatorSet should be included in allocatedSets");
        assert_IsAllocatedToStrats(operator, operatorSet, active.strategies, "strategies should be included in allocatedStrategies");
        assert_CurrentMagnitude(operator, active, "queried allocation should equal active allocation");
        assert_CurMinSlashableEqualsMinAllocated(operator, operatorSet, active.strategies, "minimum slashable stake should equal allocated stake at current block");
        
        /// Check that additional stake just became slashable
        assert_Snap_StakeBecameSlashable(operator, active.operatorSet, active.strategies, "registration should make entirety of active allocation slashable");
    }

    /// @dev Check registration invariants. Assumes the operator has a PENDING allocation
    /// to the set, but that the allocation's effect block has not yet been reached
    function check_Registration_State_PendingAllocation(
        User operator,
        IAllocationManagerTypes.AllocateParams memory params
    ) internal {
        OperatorSet memory operatorSet = params.operatorSet;
        check_Base_Registration_State(operator, operatorSet);

        assert_IsAllocated(operator, operatorSet, "operator should be allocated to set, even while pending");
        assert_Snap_Unchanged_SlashableStake(operator, operatorSet, allStrats, "operator should not have increased slashable stake");

        /// Roll forward to check status when the allocation is completed
        _rollForward_AllocationDelay(operator);

        /// NOTE: Mimics checks made in `..._ActiveAllocation`, minus the base check
        {
            /// Given an active allocation, check that the allocation is reflected in state
            assert_IsAllocatedToSet(operator, operatorSet, "operatorSet should be included in allocatedSets");
            assert_IsAllocatedToStrats(operator, operatorSet, params.strategies, "strategies should be included in allocatedStrategies");
            assert_CurrentMagnitude(operator, params, "queried allocation should equal active allocation");
            assert_CurMinSlashableEqualsMinAllocated(operator, operatorSet, params.strategies, "minimum slashable stake should equal allocated stake at current block");
        
            /// Check that additional stake just became slashable
            assert_Snap_StakeBecameSlashable(operator, params.operatorSet, params.strategies, "registration should make entirety of active allocation slashable");
        }

        /// Reset block number so test is not affected
        _rollBackward_AllocationDelay(operator);
    }

    /// @dev NOTE - this is for basic ALLOCATION invariants, NOT DEALLOCATION
    function check_Base_Allocation_State(
        User operator,
        IAllocationManagerTypes.AllocateParams memory params
    ) internal {
        assert_IsAllocated(operator, params.operatorSet, "operator should be allocated to set");
        
        // Any call that allocates more magnitude
        assert_Snap_Created_PendingIncrease(operator, params, "should have created a pending allocation");
        assert_Snap_Added_EncumberedMagnitude(operator, params.strategies, params.newMagnitudes, "encumbered mag should have increased by allocated amount");
        assert_Snap_Removed_AllocatableMagnitude(operator, params.strategies, params.newMagnitudes, "allocatable mag should have decreased by allocated amount");

        // All calls to modifyAllocations
        assert_Snap_Unchanged_MaxMagnitude(operator, params.strategies, "max magnitude should not have changed");
        assert_MaxEqualsAllocatablePlusEncumbered(operator, "max magnitude should equal encumbered plus allocatable");

        // Roll forward so we can check the updated allocation
        _rollForward_AllocationDelay(operator);

        assert_MaxEqualsAllocatablePlusEncumbered(operator, "max magnitude should equal encumbered plus allocatable");
        assert_CurrentMagnitude(operator, params, "current magnitude should match allocate params");
        assert_NoPendingModification(operator, params.operatorSet, params.strategies, "there should not be a pending modification for any strategy");

        // Reset block number so test is not affected
        _rollBackward_AllocationDelay(operator);
    }

    /// @dev Check an operator's first allocation to an operator set
    function check_NotSlashable_Allocation_State(
        User operator,
        IAllocationManagerTypes.AllocateParams memory params
    ) internal {
        check_Base_Allocation_State(operator, params);

        // Roll forward so we can check the updated allocation
        _rollForward_AllocationDelay(operator);

        assert_NoSlashableStake(operator, params.operatorSet, "operator should not have any slashable stake");
        assert_Snap_Unchanged_SlashableStake(operator, params.operatorSet, params.strategies, "slashable stake should not be changed");

        // Reset block number so test is not affected
        _rollBackward_AllocationDelay(operator);
    }

    /// @dev Check an operator's first allocation to an operator set
    function check_Slashable_Allocation_State(
        User operator,
        IAllocationManagerTypes.AllocateParams memory params,
        uint[] memory slashableShares
    ) internal {
        check_Base_Allocation_State(operator, params);

        // Roll forward so we can check the updated allocation
        _rollForward_AllocationDelay(operator);

        assert_Snap_Added_SlashableStake(operator, params.operatorSet, params.strategies, slashableShares, "operator should have added slashable stake after allocation is active");

        // Reset block number so test is not affected
        _rollBackward_AllocationDelay(operator);
    }

    /// @param slashableShares for each strategy in params.strategies, the number of shares that will NO LONGER BE SLASHABLE
    /// when the deallocation is completed
    function check_Slashable_Deallocation_State(
        User operator,
        IAllocationManagerTypes.AllocateParams memory params,
        uint[] memory slashableShares
    ) internal {
        assert_Snap_Unchanged_AllocatedSets(operator, "should not have removed operator set from allocated sets");

        // Any call that deallocates magnitude
        assert_Snap_Created_PendingDecrease(operator, params, "should have created a pending deallocation");
        assert_Snap_Unchanged_EncumberedMagnitude(operator, params.strategies, "should not have changed encumbered magnitude");
        assert_Snap_Unchanged_AllocatableMagnitude(operator, params.strategies, "should not have changed allocatable magnitude");
        assert_Snap_Unchanged_SlashableStake(operator, params.operatorSet, params.strategies, "operator should have unchanged slashable stake");

        // All calls to modifyAllocations
        assert_Snap_Unchanged_MaxMagnitude(operator, params.strategies, "max magnitude should not have changed");
        assert_MaxEqualsAllocatablePlusEncumbered(operator, "max magnitude should equal encumbered plus allocatable");

        // Roll forward so we can check the updated deallocation
        _rollForward_DeallocationDelay();

        assert_MaxEqualsAllocatablePlusEncumbered(operator, "max magnitude should equal encumbered plus allocatable");
        assert_CurrentMagnitude(operator, params, "current magnitude should match deallocate params");
        assert_NoPendingModification(operator, params.operatorSet, params.strategies, "there should not be a pending modification for any strategy");
        assert_Snap_Removed_SlashableStake(operator, params.operatorSet, params.strategies, slashableShares, "operator should have removed slashable shares");

        // Reset block number so test is not affected
        _rollBackward_DeallocationDelay();
    }

    /// @dev Checks invariants for a deallocation that was created when the operator was NOT slashable
    /// (This means the deallocation should have been completed immediately)
    function check_NotSlashable_Deallocation_State(
        User operator,
        IAllocationManagerTypes.AllocateParams memory allocateParams,
        IAllocationManagerTypes.AllocateParams memory deallocateParams
    ) internal {
        OperatorSet memory operatorSet = allocateParams.operatorSet;

        // Any instant deallocation
        assert_NoPendingModification(operator, operatorSet, allocateParams.strategies, "there should not be a pending modification for any strategy");
        assert_Snap_Removed_EncumberedMagnitude(operator, allocateParams.strategies, allocateParams.newMagnitudes, "should have removed allocation from encumbered magnitude");
        assert_Snap_Added_AllocatableMagnitude(operator, allocateParams.strategies, allocateParams.newMagnitudes, "should have added allocation to allocatable magnitude");
        assert_Snap_Unchanged_MaxMagnitude(operator, allocateParams.strategies, "max magnitude should not have changed");
        
        assert_MaxEqualsAllocatablePlusEncumbered(operator, "max magnitude should equal encumbered plus allocatable");
        assert_CurrentMagnitude(operator, deallocateParams, "current magnitude should match deallocate params");
        
        assert_Snap_Unchanged_SlashableStake(operator, operatorSet, allocateParams.strategies, "slashable shares should remain unchanged because operator was not slashable");
    }

    function check_FullyDeallocated_State(
        User operator,
        IAllocationManagerTypes.AllocateParams memory allocateParams,
        IAllocationManagerTypes.AllocateParams memory deallocateParams
    ) internal {
        OperatorSet memory operatorSet = allocateParams.operatorSet;
        assert_NoSlashableStake(operator, operatorSet, "should not have any slashable stake");
        // TODO - broken; do we want to fix this?
        // assert_Snap_Removed_AllocatedSet(operator, operatorSet, "should have removed operator set from allocated sets");

        // Any instant deallocation
        assert_Snap_Removed_EncumberedMagnitude(operator, allocateParams.strategies, allocateParams.newMagnitudes, "should have removed allocation from encumbered magnitude");
        assert_Snap_Added_AllocatableMagnitude(operator, allocateParams.strategies, allocateParams.newMagnitudes, "should have added allocation to allocatable magnitude");
        assert_Snap_Unchanged_MaxMagnitude(operator, allStrats, "max magnitude should not have changed");
        
        assert_MaxEqualsAllocatablePlusEncumbered(operator, "max magnitude should equal encumbered plus allocatable");
        assert_CurrentMagnitude(operator, deallocateParams, "current magnitude should match deallocate params");
        assert_NoPendingModification(operator, operatorSet, deallocateParams.strategies, "there should not be a pending modification for any strategy");
    }

    // TODO: improvement needed 

    function check_Withdrawal_AsTokens_State_AfterSlash(
        User staker,
        User operator,
        IDelegationManagerTypes.Withdrawal memory withdrawal,
        IAllocationManagerTypes.AllocateParams memory allocateParams,
        IAllocationManagerTypes.SlashingParams memory slashingParams,
        uint[] memory expectedTokens
    ) internal {
        IERC20[] memory tokens = new IERC20[](withdrawal.strategies.length);

        for (uint i; i < withdrawal.strategies.length; i++) {
            IStrategy strat = withdrawal.strategies[i];

            bool isBeaconChainETHStrategy = strat == beaconChainETHStrategy;

            tokens[i] = isBeaconChainETHStrategy ? NATIVE_ETH : withdrawal.strategies[i].underlyingToken();
            
            if (slashingParams.strategies.contains(strat)) {
                uint wadToSlash = slashingParams.wadsToSlash[slashingParams.strategies.indexOf(strat)];

                expectedTokens[i] -= expectedTokens[i]
                    .mulWadRoundUp(allocateParams.newMagnitudes[i].mulWadRoundUp(wadToSlash));

                uint256 max = allocationManager.getMaxMagnitude(address(operator), strat);

                withdrawal.scaledShares[i] -= withdrawal.scaledShares[i].calcSlashedAmount(WAD, max);

                // Round down to the nearest gwei for beaconchain ETH strategy.
                if (isBeaconChainETHStrategy) {
                    expectedTokens[i] -= expectedTokens[i] % 1 gwei;
                }
            }
        }

        // Common checks
        assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawal), "staker withdrawal should no longer be pending");
        
        // TODO FIXME
        // assert_Snap_Added_TokenBalances(staker, tokens, expectedTokens, "staker should have received expected tokens");
        assert_Snap_Unchanged_Staker_DepositShares(staker, "staker shares should not have changed");
        assert_Snap_Removed_StrategyShares(withdrawal.strategies, withdrawal.scaledShares, "strategies should have total shares decremented");

        // Checks specific to an operator that the Staker has delegated to
        if (operator != User(payable(0))) {
            if (operator != staker) {
                assert_Snap_Unchanged_TokenBalances(operator, "operator token balances should not have changed");
            }
            assert_Snap_Unchanged_OperatorShares(operator, "operator shares should not have changed");
        }
    }

    function check_Withdrawal_AsShares_State_AfterSlash(
        User staker,
        User operator,
        IDelegationManagerTypes.Withdrawal memory withdrawal,
        IAllocationManagerTypes.AllocateParams memory allocateParams, // TODO - was this needed?
        IAllocationManagerTypes.SlashingParams memory slashingParams
    ) internal {
        IERC20[] memory tokens = new IERC20[](withdrawal.strategies.length);

        for (uint i; i < withdrawal.strategies.length; i++) {
            IStrategy strat = withdrawal.strategies[i];

            bool isBeaconChainETHStrategy = strat == beaconChainETHStrategy;

            tokens[i] = isBeaconChainETHStrategy ? NATIVE_ETH : withdrawal.strategies[i].underlyingToken();
            
            if (slashingParams.strategies.contains(strat)) {
                uint256 max = allocationManager.getMaxMagnitude(address(operator), strat);

                withdrawal.scaledShares[i] -= withdrawal.scaledShares[i].calcSlashedAmount(WAD, max);
            }
        }
        
        // Common checks applicable to both user and non-user operator types
        assert_WithdrawalNotPending(delegationManager.calculateWithdrawalRoot(withdrawal), "staker withdrawal should no longer be pending");
        assert_Snap_Unchanged_TokenBalances(staker, "staker should not have any change in underlying token balances");
        assert_Snap_Added_Staker_DepositShares(staker, withdrawal.strategies,  withdrawal.scaledShares, "staker should have received expected shares");
        assert_Snap_Unchanged_StrategyShares(withdrawal.strategies, "strategies should have total shares unchanged");

        // Additional checks or handling for the non-user operator scenario
        if (operator != User(User(payable(0)))) {
            if (operator != staker) {
                assert_Snap_Unchanged_TokenBalances(operator, "operator should not have any change in underlying token balances");
            }
        }
    }
}
