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
    using StdStyle for *;

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
        // TODO - currently only used after a `NoWithdrawNoRewards` action. Investigate re-adding in future.
        // assert_Snap_Removed_ActiveValidatorCount(staker, slashedValidators.length, "should have decreased active validator count");
        // assert_Snap_Removed_ActiveValidators(staker, slashedValidators, "exited validators should each be WITHDRAWN");
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
        assert_Snap_Unchanged_Staker_WithdrawableShares_Delegation(staker, "withdrawable shares should be unchanged after delegating");
        uint256[] memory delegatableShares = _getPrevStakerWithdrawableShares(staker, strategies);
        assert_Snap_Added_OperatorShares(operator, strategies, delegatableShares, "operator should have received shares");
    }

    function check_QueuedWithdrawal_State(
        User staker, 
        User operator, 
        IStrategy[] memory strategies, 
        uint[] memory shares, 
        Withdrawal[] memory withdrawals, 
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
        Withdrawal[] memory withdrawals,
        bytes32[] memory withdrawalRoots,
        IStrategy[] memory strategies,
        uint[] memory stakerDepositShares,
        uint[] memory stakerDelegatedShares 
    ) internal {
        /// Undelegate from an operator
        //
        // ... check that the staker is undelegated, all strategies from which the staker is deposited are unqueued,
        //     that the returned root matches the hashes for each strategy and share amounts, and that the staker
        //     and operator have reduced shares
        assertFalse(delegationManager.isDelegated(address(staker)),
            "check_Undelegate_State: staker should not be delegated");
        assert_ValidWithdrawalHashes(withdrawals, withdrawalRoots,
            "check_Undelegate_State: calculated withdrawal should match returned root");
        assert_AllWithdrawalsPending(withdrawalRoots,
            "check_Undelegate_State: stakers withdrawal should now be pending");
        assert_DSF_Reset(staker, strategies,
            "check_Undelegate_State: staker dsfs should be reset to wad");
        assert_Snap_Added_QueuedWithdrawals(staker, withdrawals,
            "check_Undelegate_State: staker should have increased nonce by withdrawals.length");
        assert_Snap_Removed_OperatorShares(operator, strategies, stakerDelegatedShares,
            "check_Undelegate_State: failed to remove operator shares");
        assert_Snap_Removed_Staker_DepositShares(staker, strategies, stakerDepositShares,
            "check_Undelegate_State: failed to remove staker shares");
        assert_Snap_RemovedAll_Staker_WithdrawableShares(staker, strategies,
            "check_QueuedWithdrawal_State: failed to remove staker withdrawable shares");
    }

    function check_Redelegate_State(
        User staker, 
        User operator,
        IDelegationManagerTypes.Withdrawal[] memory withdrawals,
        bytes32[] memory withdrawalRoots,
        IStrategy[] memory strategies,
        uint[] memory stakerDepositShares,
        uint[] memory stakerDelegatedShares  
    ) internal {
        /// Redelegate to a new operator
        //
        // ... check that the staker is delegated to new operator, all strategies from which the staker is deposited are unqueued,
        //     that the returned root matches the hashes for each strategy and share amounts, and that the staker
        //     and operator have reduced shares
        assertTrue(delegationManager.isDelegated(address(staker)),
            "check_Redelegate_State: staker should not be delegated");
        assert_ValidWithdrawalHashes(withdrawals, withdrawalRoots,
            "check_Redelegate_State: calculated withdrawal should match returned root");
        assert_AllWithdrawalsPending(withdrawalRoots,
            "check_Redelegate_State: stakers withdrawal should now be pending");
        assert_Snap_Added_QueuedWithdrawals(staker, withdrawals,
            "check_Redelegate_State: staker should have increased nonce by withdrawals.length");
        assert_Snap_Removed_OperatorShares(operator, strategies, stakerDelegatedShares,
            "check_Redelegate_State: failed to remove operator shares");
        assert_Snap_Removed_Staker_DepositShares(staker, strategies, stakerDepositShares,
            "check_Redelegate_State: failed to remove staker shares");
        assert_Snap_RemovedAll_Staker_WithdrawableShares(staker, strategies,
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
        Withdrawal memory withdrawal,
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
        Withdrawal memory withdrawal,
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
        Withdrawal memory withdrawal,
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
        assert_Snap_Added_Staker_DepositShares(staker, strategies, shares, "staker should have received expected deposit shares");
        uint[] memory expectedWithdrawableShares = _getExpectedWithdrawableSharesUndelegate(staker, strategies, shares);
        assert_Snap_Added_Staker_WithdrawableShares(staker, strategies, expectedWithdrawableShares, "staker should have received expected withdrawable shares");
        assert_Snap_Unchanged_OperatorShares(operator, "operator should have shares unchanged");
        assert_Snap_Unchanged_StrategyShares(strategies, "strategies should have total shares unchanged");
    }

    function check_Withdrawal_AsShares_Redelegated_State(
        User staker,
        User operator,
        User newOperator,
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
                                ALM - BASIC INVARIANTS
    *******************************************************************************/

    /// @dev Run a method as if the user's allocation delay had passed
    /// When done, reset block number so other tests are not affected
    modifier activateAllocation(User operator) {
        _rollForward_AllocationDelay(operator);

        _;

        _rollBackward_AllocationDelay(operator);
    }

    /// @dev Run a method as if the deallocation delay has passed
    /// When done, reset block number so other tests are not affected
    modifier activateDeallocation() {
        _rollForward_DeallocationDelay();

        _;

        _rollBackward_DeallocationDelay();
    }

    /// @dev Run a method ONLY IF the operator has a nonzero activation delay
    modifier skipIfInstantAlloc(User operator) {
        /// Note: if the ALM says the allocation delay is "not set", this will revert
        uint32 delay = _getExistingAllocationDelay(operator);
        
        if (delay != 0) {
            _;
        } else {
            console.log("%s", "skipping checks for operator with allocation delay of 0".italic());
        }
    }

    /// @dev Check global max magnitude invariants - these should ALWAYS hold
    function check_MaxMag_Invariants(
        User operator
    ) internal {
        assert_MaxMagsEqualMaxMagsAtCurrentBlock(operator, allStrats, "max magnitudes should equal upperlookup at current block");
        assert_MaxEqualsAllocatablePlusEncumbered(operator, "max magnitude should equal encumbered plus allocatable");
    }

    /// @dev Check that the last call to modifyAllocations resulted in a non-pending modification
    function check_ActiveModification_State(
        User operator,
        AllocateParams memory params
    ) internal {
        OperatorSet memory operatorSet = params.operatorSet;
        IStrategy[] memory strategies = params.strategies;

        assert_CurrentMagnitude(operator, params, "current magnitude should match allocate params");
        assert_NoPendingModification(operator, operatorSet, strategies, "there should not be a pending modification for any strategy");
    }

    function check_IsSlashable_State(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies
    ) internal {
        assert_IsSlashable(operator, operatorSet, "operator should be slashable for operator set");
        assert_CurMinSlashableEqualsMinAllocated(operator, operatorSet, strategies, "minimum slashable stake should equal allocated stake at current block");
    }

    function check_NotSlashable_State(
        User operator,
        OperatorSet memory operatorSet
    ) internal {
        assert_NotSlashable(operator, operatorSet, "operator should not be slashable for operator set");
        assert_NoSlashableStake(operator, operatorSet, "operator should not have any slashable stake");
    }

    /*******************************************************************************
                                 ALM - REGISTRATION
    *******************************************************************************/
    
    /// @dev Basic invariants that should hold after EVERY call to `registerForOperatorSets`
    /// NOTE: These are only slightly modified from check_Base_Deregistration_State
    /// If you add invariants here, consider adding them there (and vice-versa)
    function check_Base_Registration_State(
        User operator,
        OperatorSet memory operatorSet
    ) internal {
        check_MaxMag_Invariants(operator);
        check_IsSlashable_State(operator, operatorSet, allocationManager.getStrategiesInOperatorSet(operatorSet));

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
        assert_Snap_Unchanged_AllocatedStake(operator, operatorSet, unallocated, "should not have updated allocated stake in any way");
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
        AllocateParams memory active
    ) internal {
        OperatorSet memory operatorSet = active.operatorSet;
        IStrategy[] memory strategies = active.strategies;

        /// Basic registerForOperatorSets invariants
        check_Base_Registration_State(operator, operatorSet);

        /// Given an active allocation, check that the allocation is reflected in state
        assert_IsAllocatedToSet(operator, operatorSet, "operatorSet should be included in allocatedSets");
        assert_IsAllocatedToSetStrats(operator, operatorSet, strategies, "strategies should be included in allocatedStrategies");
        assert_CurrentMagnitude(operator, active, "queried allocation should equal active allocation");
        
        /// Check that additional stake just became slashable
        assert_Snap_Unchanged_AllocatedStake(operator, operatorSet, strategies, "should not have updated allocated stake in any way");
        assert_Snap_StakeBecameSlashable(operator, operatorSet, strategies, "registration should make entirety of active allocation slashable");
    }

    /// @dev Check registration invariants. Assumes the operator has a PENDING allocation
    /// to the set, but that the allocation's effect block has not yet been reached
    function check_Registration_State_PendingAllocation(
        User operator,
        AllocateParams memory params
    ) internal {
        OperatorSet memory operatorSet = params.operatorSet;
        IStrategy[] memory strategies = params.strategies;

        check_Base_Registration_State(operator, operatorSet);

        assert_IsAllocatedToSet(operator, operatorSet, "operator should be allocated to set, even while pending");
        assert_IsAllocatedToSetStrats(operator, operatorSet, strategies, "strategies should be included in allocatedStrategies");

        /// Skip pending checks if operator has no allocation delay
        uint32 delay = _getExistingAllocationDelay(operator);
        if (delay == 0) {
            return;
        }

        assert_Snap_Unchanged_AllocatedStake(operator, operatorSet, strategies, "should not have updated allocated stake in any way");
        assert_Snap_Unchanged_SlashableStake(operator, operatorSet, allStrats, "operator should not have increased slashable stake");
    }

    /*******************************************************************************
                                 ALM - DEREGISTRATION
    *******************************************************************************/

    /// @dev Basic invariants that should hold after EVERY call to `deregisterFromOperatorSets`
    /// NOTE: These are only slightly modified from check_Base_Registration_State
    /// If you add invariants here, consider adding them there (and vice-versa)
    function check_Base_Deregistration_State(
        User operator,
        OperatorSet memory operatorSet
    ) internal {
        check_MaxMag_Invariants(operator);

        // Deregistration SHOULD remove the operator as a member of the set
        assert_Snap_Became_Deregistered(operator, operatorSet, "operator should have been registered before, and is now deregistered");
        assert_Snap_Removed_RegisteredSet(operator, operatorSet, "should have removed operator set from list of registered sets");
        assert_Snap_Removed_MemberOfSet(operator, operatorSet, "should have removed operator from list of set members");

        // Deregistration should NOT change slashability, magnitude, allocations, or allocated sets
        assert_Snap_Remains_Slashable(operator, operatorSet, "operator should have been slashable already, and should still be slashable");
        assert_Snap_Unchanged_AllocatedSets(operator, "should not have updated allocated sets");
        assert_Snap_Unchanged_AllocatedStrats(operator, operatorSet, "should not have updated allocated strategies");
        assert_Snap_Unchanged_MaxMagnitude(operator, allStrats, "should not have updated max magnitudes in any way");
        assert_Snap_Unchanged_AllocatedStake(operator, operatorSet, allStrats, "should not have updated allocated stake in any way");
        assert_Snap_Unchanged_StrategyAllocations(operator, operatorSet, allStrats, "should not have updated any individual allocations");
        assert_Snap_Unchanged_EncumberedMagnitude(operator, allStrats, "should not have updated encumbered magnitude");
        assert_Snap_Unchanged_AllocatableMagnitude(operator, allStrats, "should not have updated allocatable magnitude");

        _rollForward_DeallocationDelay();
        {
            check_NotSlashable_State(operator, operatorSet);
        }
        _rollBackward_DeallocationDelay();
    }

    function check_Deregistration_State_NoAllocation(
        User operator,
        OperatorSet memory operatorSet
    ) internal {
        check_Base_Deregistration_State(operator, operatorSet);

        assert_Snap_Unchanged_AllocatedStake(operator, operatorSet, allStrats, "should not have updated allocated stake in any way");
        assert_Snap_Unchanged_SlashableStake(operator, operatorSet, allStrats, "operator should not have increased slashable stake for any given strategy");
    }

    function check_Deregistration_State_ActiveAllocation(
        User operator,
        OperatorSet memory operatorSet
    ) internal {
        check_Base_Deregistration_State(operator, operatorSet);

        assert_Snap_Unchanged_AllocatedStake(operator, operatorSet, allStrats, "should not have updated allocated stake in any way");
        assert_Snap_Unchanged_SlashableStake(operator, operatorSet, allStrats, "operator should not have increased slashable stake for any given strategy");
    }

    function check_Deregistration_State_PendingAllocation(
        User operator,
        OperatorSet memory operatorSet
    ) internal {
        check_Base_Deregistration_State(operator, operatorSet);

        assert_Snap_Unchanged_AllocatedStake(operator, operatorSet, allStrats, "should not have updated allocated stake in any way");
        assert_Snap_Unchanged_SlashableStake(operator, operatorSet, allStrats, "operator should not have increased slashable stake for any given strategy");
    }

    /*******************************************************************************
                                ALM - INCREASE ALLOCATION
    *******************************************************************************/

    /// @dev Basic invariants that should hold after all calls to `modifyAllocations`
    /// where the input `params` represent an _increase_ in magnitude
    function check_Base_IncrAlloc_State(
        User operator,
        AllocateParams memory params
    ) internal {
        check_MaxMag_Invariants(operator);

        OperatorSet memory operatorSet = params.operatorSet;
        IStrategy[] memory strategies = params.strategies;
        uint64[] memory newMagnitudes = params.newMagnitudes;

        // Increasing Allocation should NOT change operator set registration, max magnitude
        assert_Snap_Unchanged_Registration(operator, operatorSet, "operator registration status should be unchanged");
        assert_Snap_Unchanged_Slashability(operator, operatorSet, "operator slashability should be unchanged");
        assert_Snap_Unchanged_RegisteredSet(operator, "list of registered sets should remain unchanged");
        assert_Snap_Unchanged_MemberOfSet(operatorSet, "list of set members should remain unchanged");
        assert_Snap_Unchanged_MaxMagnitude(operator, allStrats, "should not have updated max magnitudes in any way");

        // Increasing Allocation SHOULD consume magnitude and mark the operator as being allocated to the set
        assert_IsAllocatedToSet(operator, operatorSet, "operator should be allocated to set");
        assert_IsAllocatedToSetStrats(operator, operatorSet, strategies, "operator should be allocated to strategies for set");
        assert_Snap_Allocated_Magnitude(operator, strategies, "operator should have allocated magnitude");
    }

    /// @dev Invariants for modifyAllocations. Use when:
    /// - operator is NOT slashable for this operator set
    /// - last call to modifyAllocations created an INCREASE in allocation
    function check_IncrAlloc_State_NotSlashable(
        User operator,
        AllocateParams memory params
    ) internal {
        check_Base_IncrAlloc_State(operator, params);
        check_NotSlashable_State(operator, params.operatorSet);

        /// Run checks on pending allocation, if the operator has a nonzero delay
        check_IncrAlloc_State_NotSlashable_Pending(operator, params);

        /// Run checks on active allocation
        check_IncrAlloc_State_NotSlashable_Active(operator, params);
    }

    /// @dev Invariants for modifyAllocations. Used when:
    /// - operator is NOT slashable for this operator set
    /// - last call to modifyAllocations created an INCREASE in allocation
    /// - effectBlock for the increase HAS NOT been reached
    function check_IncrAlloc_State_NotSlashable_Pending(
        User operator,
        AllocateParams memory params
    ) private skipIfInstantAlloc(operator) {
        // Validate operator allocation is pending
        assert_HasPendingIncrease(operator, params, "params should reflect a pending modification");

        // Should not have allocated magnitude
        assert_Snap_Unchanged_AllocatedStake(operator, params.operatorSet, params.strategies, "should not have updated allocated stake in any way");
        assert_Snap_Unchanged_SlashableStake(operator, params.operatorSet, params.strategies, "should not have updated allocated stake in any way");
    }

    /// @dev Invariants for modifyAllocations. Used when:
    /// - operator is NOT slashable for this operator set
    /// - last call to modifyAllocations created an INCREASE in allocation
    /// - effectBlock for the increase HAS been reached
    function check_IncrAlloc_State_NotSlashable_Active(
        User operator,
        AllocateParams memory params
    ) private activateAllocation(operator) {
        // Validate allocation is active
        check_ActiveModification_State(operator, params);

        // SHOULD set current magnitude and increase allocated stake
        assert_Snap_Set_CurrentMagnitude(operator, params, "should have updated the operator's magnitude");
        assert_HasAllocatedStake(operator, params, "operator should have expected allocated stake for each strategy");
        assert_Snap_StakeBecameAllocated(operator, params.operatorSet, params.strategies, "allocated stake should have increased");

        // Should NOT change slashable stake
        assert_Snap_Unchanged_SlashableStake(operator, params.operatorSet, params.strategies, "slashable stake should not be changed");
    }

    /// @dev Invariants for modifyAllocations. Use when:
    /// - operator IS slashable for this operator set
    /// - last call to modifyAllocations created an INCREASE in allocation
    function check_IncrAlloc_State_Slashable(
        User operator,
        AllocateParams memory params
    ) internal {
        check_Base_IncrAlloc_State(operator, params);
        check_IsSlashable_State(operator, params.operatorSet, params.strategies);        

        /// Run checks on pending allocation, if the operator has a nonzero delay
        check_IncrAlloc_State_Slashable_Pending(operator, params);

        /// Run checks on active allocation
        check_IncrAlloc_State_Slashable_Active(operator, params);
    }

    /// @dev Invariants for modifyAllocations. Used when:
    /// - operator IS slashable for this operator set
    /// - last call to modifyAllocations created an INCREASE in allocation
    /// - effectBlock for the increase HAS NOT been reached
    function check_IncrAlloc_State_Slashable_Pending(
        User operator,
        AllocateParams memory params
    ) private skipIfInstantAlloc(operator) {
        OperatorSet memory operatorSet = params.operatorSet;
        IStrategy[] memory strategies = params.strategies;

        // Validate operator has pending allocation and unchanged allocated/slashable stake
        assert_HasPendingIncrease(operator, params, "params should reflect a pending modification");

        // Should not have allocated magnitude
        assert_Snap_Unchanged_AllocatedStake(operator, operatorSet, strategies, "should not have updated allocated stake in any way");
        assert_Snap_Unchanged_SlashableStake(operator, operatorSet, strategies, "should not have updated allocated stake in any way");
    }

    /// @dev Invariants for modifyAllocations. Used when:
    /// - operator IS slashable for this operator set
    /// - last call to modifyAllocations created an INCREASE in allocation
    /// - effectBlock for the increase HAS been reached
    function check_IncrAlloc_State_Slashable_Active(
        User operator,
        AllocateParams memory params
    ) private activateAllocation(operator) {
        // Validate operator does not have a pending modification, and has expected slashable stake
        check_ActiveModification_State(operator, params);

        // SHOULD set current magnitude and increase slashable/allocated stake
        assert_Snap_Set_CurrentMagnitude(operator, params, "should have updated the operator's magnitude");
        assert_HasAllocatedStake(operator, params, "operator should have expected allocated stake for each strategy");
        assert_HasSlashableStake(operator, params, "operator should have expected slashable stake for each strategy");
        assert_Snap_StakeBecameAllocated(operator, params.operatorSet, params.strategies, "allocated stake should have increased");
        assert_Snap_StakeBecameSlashable(operator, params.operatorSet, params.strategies, "slashable stake should have increased");
    }

    /*******************************************************************************
                                ALM - DECREASE ALLOCATION
    *******************************************************************************/

    /// @dev Basic invariants that should hold after all calls to `modifyAllocations`
    /// where the input `params` represent a decrease in magnitude
    function check_Base_DecrAlloc_State(
        User operator,
        AllocateParams memory params
    ) internal {
        check_MaxMag_Invariants(operator);

        OperatorSet memory operatorSet = params.operatorSet;
        IStrategy[] memory strategies = params.strategies;
        uint64[] memory newMagnitudes = params.newMagnitudes;

        // Decreasing Allocation should NOT change operator set registration, max magnitude
        assert_Snap_Unchanged_Registration(operator, operatorSet, "operator registration status should be unchanged");
        assert_Snap_Unchanged_Slashability(operator, operatorSet, "operator slashability should be unchanged");
        assert_Snap_Unchanged_RegisteredSet(operator, "list of registered sets should remain unchanged");
        assert_Snap_Unchanged_MemberOfSet(operatorSet, "list of set members should remain unchanged");
        assert_Snap_Unchanged_MaxMagnitude(operator, allStrats, "should not have updated max magnitudes in any way");
    }

    function check_DecrAlloc_State_NotSlashable(
        User operator,
        AllocateParams memory params
    ) internal {
        OperatorSet memory operatorSet = params.operatorSet;
        IStrategy[] memory strategies = params.strategies;
        
        check_Base_DecrAlloc_State(operator, params);
        check_NotSlashable_State(operator, operatorSet);
        check_ActiveModification_State(operator, params);
         
        // SHOULD set current magnitude and decrease allocated stake
        assert_HasAllocatedStake(operator, params, "operator should have expected allocated stake for each strategy");        
        assert_Snap_Set_CurrentMagnitude(operator, params, "should have updated the operator's magnitude");
        assert_Snap_StakeBecameDeallocated(operator, operatorSet, strategies, "allocated stake should have increased");
        assert_Snap_Deallocated_Magnitude(operator, strategies, "should have deallocated magnitude");
    }

    function check_DecrAlloc_State_Slashable(
        User operator,
        AllocateParams memory params
    ) internal {
        check_Base_DecrAlloc_State(operator, params);
        check_IsSlashable_State(operator, params.operatorSet, params.strategies);

        // Run checks on pending deallocation
        check_DecrAlloc_State_Slashable_Pending(operator, params);

        // Run checks on active deallocation
        check_DecrAlloc_State_Slashable_Active(operator, params);
    }

    function check_DecrAlloc_State_Slashable_Pending(
        User operator,
        AllocateParams memory params
    ) private {
        // Validate deallocation is pending
        assert_HasPendingDecrease(operator, params, "params should reflect a pending modification");

        // Should NOT have changed allocated magnitude or stake
        assert_Snap_Unchanged_EncumberedMagnitude(operator, params.strategies, "should not have changed encumbered magnitude");
        assert_Snap_Unchanged_AllocatableMagnitude(operator, params.strategies, "should not have changed allocatable magnitude");
        assert_Snap_Unchanged_AllocatedSets(operator, "should not have removed operator set from allocated sets");
        assert_Snap_Unchanged_AllocatedStake(operator, params.operatorSet, params.strategies, "should not have updated allocated stake in any way");
        assert_Snap_Unchanged_SlashableStake(operator, params.operatorSet, params.strategies, "should not have updated allocated stake in any way");
    }

    function check_DecrAlloc_State_Slashable_Active(
        User operator,
        AllocateParams memory params
    ) private activateDeallocation() {
        OperatorSet memory operatorSet = params.operatorSet;
        IStrategy[] memory strategies = params.strategies;

        check_ActiveModification_State(operator, params);

        // SHOULD set current magnitude and decrease allocated stake
        assert_Snap_Set_CurrentMagnitude(operator, params, "should have updated the operator's magnitude");
        assert_HasAllocatedStake(operator, params, "operator should have expected allocated stake for each strategy");
        assert_HasSlashableStake(operator, params, "operator should have expected slashable stake for each strategy");

        assert_Snap_StakeBecomeUnslashable(operator, operatorSet, strategies, "slashable stake should have decreased");
        assert_Snap_StakeBecameDeallocated(operator, params.operatorSet, params.strategies, "allocated stake should have decreased");
        assert_Snap_Deallocated_Magnitude(operator, strategies, "should have deallocated magnitude");
    }

    function check_FullyDeallocated_State(
        User operator,
        AllocateParams memory allocateParams,
        AllocateParams memory deallocateParams
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
        check_ActiveModification_State(operator, deallocateParams); 
    }

    /*******************************************************************************
                                ALM - SLASHING
    *******************************************************************************/

    function check_Base_Slashing_State(
        User operator,
        AllocateParams memory allocateParams,
        SlashingParams memory slashParams
    ) internal {
        OperatorSet memory operatorSet = allocateParams.operatorSet;

        check_MaxMag_Invariants(operator);
        check_IsSlashable_State(operator, operatorSet, allocateParams.strategies);

        // Slashing SHOULD change max magnitude and current allocation
        assert_Snap_Slashed_MaxMagnitude(operator, slashParams, "slash should lower max magnitude");
        assert_Snap_Slashed_EncumberedMagnitude(operator, slashParams, "slash should lower max magnitude");
        assert_Snap_Slashed_AllocatedStake(operator, operatorSet, slashParams, "slash should lower allocated stake");
        assert_Snap_Slashed_SlashableStake(operator, operatorSet, slashParams, "slash should lower slashable stake");
        assert_Snap_Slashed_OperatorShares(operator, slashParams, "slash should remove operator shares");
        assert_Snap_Slashed_Allocation(operator, operatorSet, slashParams, "slash should reduce current magnitude");
        assert_Snap_Increased_BurnableShares(operator, slashParams, "slash should increase burnable shares");

        // Slashing SHOULD NOT change allocatable magnitude, registration, and slashability status
        assert_Snap_Unchanged_AllocatableMagnitude(operator, allStrats, "slashing should not change allocatable magnitude");
        assert_Snap_Unchanged_Registration(operator, operatorSet, "slash should not change registration status");
        assert_Snap_Unchanged_Slashability(operator, operatorSet, "slash should not change slashability status");
        assert_Snap_Unchanged_AllocatedSets(operator, "should not have updated allocated sets");
        assert_Snap_Unchanged_AllocatedStrats(operator, operatorSet, "should not have updated allocated strategies");
    }

    // TODO: improvement needed 

    function check_Withdrawal_AsTokens_State_AfterSlash(
        User staker,
        User operator,
        Withdrawal memory withdrawal,
        AllocateParams memory allocateParams,
        SlashingParams memory slashingParams,
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
        Withdrawal memory withdrawal,
        AllocateParams memory allocateParams, // TODO - was this needed?
        SlashingParams memory slashingParams
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
