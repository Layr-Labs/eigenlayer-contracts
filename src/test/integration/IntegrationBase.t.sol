// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

import "src/contracts/libraries/BeaconChainProofs.sol";
import "src/contracts/libraries/SlashingLib.sol";

import "src/test/integration/IntegrationDeployer.t.sol";
import "src/test/integration/TimeMachine.t.sol";
import "src/test/integration/users/User.t.sol";
import "src/test/integration/users/User_M1.t.sol";

abstract contract IntegrationBase is IntegrationDeployer {
    using StdStyle for *;
    using SlashingLib for *;
    using Strings for *;
    using print for *;

    using ArrayLib for IStrategy[];

    uint numStakers = 0;
    uint numOperators = 0;
    uint numAVSs = 0;

    // Lists of operators created before the m2 (not slashing) upgrade
    //
    // When we call _upgradeEigenLayerContracts, we iterate over
    // these lists and migrate perform the standard migration actions
    // for each user
    User[] operatorsToMigrate;
    User[] stakersToMigrate;

    /**
     * Gen/Init methods:
     */

    /**
     * @dev Create a new user according to configured random variants.
     * This user is ready to deposit into some strategies and has some underlying token balances
     */
    function _newRandomStaker() internal returns (User, IStrategy[] memory, uint[] memory) {
        string memory stakerName;

        User staker;
        IStrategy[] memory strategies;
        uint[] memory tokenBalances;

        if (forkType == MAINNET && !isUpgraded) {
            stakerName = string.concat("M2Staker", cheats.toString(numStakers));

            (staker, strategies, tokenBalances) = _randUser(stakerName);

            stakersToMigrate.push(staker);
        } else {
            stakerName = string.concat("staker", cheats.toString(numStakers));

            (staker, strategies, tokenBalances) = _randUser(stakerName);
        }

        assert_HasUnderlyingTokenBalances(staker, strategies, tokenBalances, "_newRandomStaker: failed to award token balances");

        numStakers++;
        return (staker, strategies, tokenBalances);
    }

    /**
     * @dev Create a new operator according to configured random variants.
     * This user will immediately deposit their randomized assets into eigenlayer.
     * @notice If forktype is mainnet and not upgraded, then the operator will only randomize LSTs assets and deposit them
     * as ETH podowner shares are not available yet. 
     */
    function _newRandomOperator() internal returns (User, IStrategy[] memory, uint[] memory) {
        User operator;
        IStrategy[] memory strategies;
        uint[] memory tokenBalances;
        uint[] memory addedShares;

        if (forkType == MAINNET && !isUpgraded) {
            string memory operatorName = string.concat("M2Operator", numOperators.toString());

            // Create an operator for M2.
            // TODO: Allow this operator to have ETH 
            (operator, strategies, tokenBalances) = _randUser_NoETH(operatorName);

            addedShares = _calculateExpectedShares(strategies, tokenBalances);
            
            User_M2(payable(operator)).registerAsOperator_M2();
            operator.depositIntoEigenlayer(strategies, tokenBalances); // Deposits interface doesn't change between M2 and slashing

            operatorsToMigrate.push(operator);
        } else {
            string memory operatorName = string.concat("operator", numOperators.toString());

            (operator, strategies, tokenBalances) = _randUser_NoETH(operatorName);

            addedShares = _calculateExpectedShares(strategies, tokenBalances);

            operator.registerAsOperator();
            operator.depositIntoEigenlayer(strategies, tokenBalances);

            // Roll passed the allocation configuration delay
            rollForward({blocks: ALLOCATION_CONFIGURATION_DELAY});
        }        

        assert_Snap_Added_Staker_DepositShares(operator, strategies, addedShares, "_newRandomOperator: failed to add delegatable shares");
        assert_Snap_Added_OperatorShares(operator, strategies, addedShares, "_newRandomOperator: failed to award shares to operator");
        assertTrue(delegationManager.isOperator(address(operator)), "_newRandomOperator: operator should be registered");

        numOperators++;
        return (operator, strategies, tokenBalances);
    }

    function _newRandomAVS() internal returns (AVS avs, OperatorSet[] memory operatorSets) {
        string memory avsName = string.concat("avs", numAVSs.toString());
        avs = _genRandAVS(avsName);
        operatorSets = avs.createOperatorSets(_randomStrategies());
        ++numAVSs;
    }

    /// @dev Send a random amount of ETH (up to 10 gwei) to the destination via `call`,
    /// triggering its fallback function. Sends a gwei-divisible amount as well as a
    /// non-gwei-divisible amount.
    ///
    /// Total sent == `gweiSent + remainderSent`
    function _sendRandomETH(address destination) internal returns (uint64 gweiSent, uint remainderSent) {
        gweiSent = uint64(_randUint({ min: 1 , max: 10 }));
        remainderSent = _randUint({ min: 1, max: 100 });
        uint totalSent = (gweiSent * GWEI_TO_WEI) + remainderSent;

        cheats.deal(address(this), address(this).balance + totalSent);
        bool r;
        bytes memory d;
        (r, d) = destination.call{ value: totalSent }("");

        return (gweiSent, remainderSent);
    }

    /// @dev If we're on mainnet, upgrade contracts to slashing and migrate stakers/operators
    function _upgradeEigenLayerContracts() internal {
        if (forkType == MAINNET) {
            require(!isUpgraded, "_upgradeEigenLayerContracts: already performed slashing upgrade");

            emit log("_upgradeEigenLayerContracts: upgrading mainnet to slashing");
            _upgradeMainnetContracts();

            // Bump block.timestamp forward to allow verifyWC proofs for migrated pods
            emit log("advancing block time to start of next epoch:");

            beaconChain.advanceEpoch_NoRewards();

            emit log("======");

            isUpgraded = true;
            emit log("_upgradeEigenLayerContracts: slashing upgrade complete");
        } else if (forkType == HOLESKY) {
            require(!isUpgraded, "_upgradeEigenLayerContracts: already performed slashing upgrade");

            emit log("_upgradeEigenLayerContracts: upgrading holesky to slashing");
            _upgradeHoleskyContracts();

            isUpgraded = true;
            emit log("_upgradeEigenLayerContracts: slashing upgrade complete");
        }
    }

    /// @dev Choose a random subset of validators (selects AT LEAST ONE)
    function _choose(uint40[] memory validators) internal returns (uint40[] memory) {
        uint rand = _randUint({ min: 1, max: validators.length ** 2 });

        uint40[] memory result = new uint40[](validators.length);
        uint newLen;
        for (uint i = 0; i < validators.length; i++) {
            // if bit set, add validator
            if (rand >> i & 1 == 1) {
                result[newLen] = validators[i];
                newLen++;
            }
        }

        // Manually update length of result
        assembly { mstore(result, newLen) }

        return result;
    }
    
    function _getTokenName(IERC20 token) internal view returns (string memory) {
        if (token == NATIVE_ETH) {
            return "Native ETH";
        }
        return IERC20Metadata(address(token)).name();
    }
    /*******************************************************************************
                                COMMON ASSERTIONS
    *******************************************************************************/

    function assert_HasNoDelegatableShares(User user, string memory err) internal view {
        (IStrategy[] memory strategies, uint[] memory shares) = 
            delegationManager.getDepositedShares(address(user));
        
        assertEq(strategies.length, 0, err);
        assertEq(strategies.length, shares.length, "assert_HasNoDelegatableShares: return length mismatch");
    }

    function assert_HasUnderlyingTokenBalances(
        User user, 
        IStrategy[] memory strategies, 
        uint[] memory expectedBalances, 
        string memory err
    ) internal view {
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            
            uint expectedBalance = expectedBalances[i];
            uint tokenBalance;
            if (strat == BEACONCHAIN_ETH_STRAT) {
                tokenBalance = address(user).balance;
            } else {
                tokenBalance = strat.underlyingToken().balanceOf(address(user));
            }

            assertApproxEqAbs(expectedBalance, tokenBalance, 1, err);
        }
    }

    function assert_HasNoUnderlyingTokenBalance(User user, IStrategy[] memory strategies, string memory err) internal view {
        assert_HasUnderlyingTokenBalances(user, strategies, new uint[](strategies.length), err);
    }

    function assert_HasExpectedShares(
        User user, 
        IStrategy[] memory strategies, 
        uint[] memory expectedShares, 
        string memory err
    ) internal view {
        uint[] memory actualShares = _getStakerDepositShares(user, strategies);
        for(uint i = 0; i < strategies.length; i++) {
            assertApproxEqAbs(expectedShares[i], actualShares[i], 1, err);
        }
    }

    function assert_HasOperatorShares(
        User user, 
        IStrategy[] memory strategies, 
        uint[] memory expectedShares, 
        string memory err
    ) internal view {
        for (uint i = 0; i < strategies.length; i++) {
            uint actualShares = delegationManager.operatorShares(address(user), strategies[i]);
            assertEq(expectedShares[i], actualShares, err);
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

    function assert_ValidWithdrawalHashes(
        IDelegationManagerTypes.Withdrawal[] memory withdrawals,
        bytes32[] memory withdrawalRoots,
        string memory err
    ) internal view {
        for (uint i = 0; i < withdrawals.length; i++) {
            assert_ValidWithdrawalHash(withdrawals[i], withdrawalRoots[i], err);
        }
    }

    function assert_ValidWithdrawalHash(
        IDelegationManagerTypes.Withdrawal memory withdrawal,
        bytes32 withdrawalRoot,
        string memory err
    ) internal view {
        assertEq(withdrawalRoot, delegationManager.calculateWithdrawalRoot(withdrawal), err);
    }

    function assert_PodBalance_Eq(
        User staker,
        uint expectedBalance,
        string memory err
    ) internal view {
        EigenPod pod = staker.pod();
        assertEq(address(pod).balance, expectedBalance, err);
    }

    function assert_ProofsRemainingEqualsActive(
        User staker,
        string memory err
    ) internal view {
        EigenPod pod = staker.pod();
        assertEq(pod.currentCheckpoint().proofsRemaining, pod.activeValidatorCount(), err);
    }

    function assert_CheckpointPodBalance(
        User staker,
        uint64 expectedPodBalanceGwei,
        string memory err
    ) internal view {
        EigenPod pod = staker.pod();
        assertEq(pod.currentCheckpoint().podBalanceGwei, expectedPodBalanceGwei, err);
    }
    
    /*******************************************************************************
                                SNAPSHOT ASSERTIONS
                       TIME TRAVELERS ONLY BEYOND THIS POINT
    *******************************************************************************/

    /*******************************************************************************
                         SNAPSHOT ASSERTIONS: ALLOCATIONS
    *******************************************************************************/

    function assert_Snap_Allocations_Modified(
        User operator,
        IAllocationManagerTypes.AllocateParams memory allocateParams,
        bool completed,
        string memory err
    ) internal {
        IAllocationManagerTypes.Allocation[] memory curAllocs = _getAllocations(operator, allocateParams.operatorSet, allocateParams.strategies);
        IAllocationManagerTypes.Allocation[] memory prevAllocs = _getPrevAllocations(operator, allocateParams.operatorSet, allocateParams.strategies);
        Magnitudes[] memory curMagnitudes = _getMagnitudes(operator, allocateParams.strategies);
        Magnitudes[] memory prevMagnitudes = _getPrevMagnitudes(operator, allocateParams.strategies);

        for (uint i = 0; i < allocateParams.strategies.length; i++) {
            IAllocationManagerTypes.Allocation memory curAlloc = curAllocs[i];
            IAllocationManagerTypes.Allocation memory prevAlloc = prevAllocs[i]; 
            
            // Check allocations
            if (completed) {
                assertEq(curAlloc.currentMagnitude, allocateParams.newMagnitudes[i], string.concat(err, " (currentMagnitude)"));
                assertEq(curAlloc.pendingDiff, 0, string.concat(err, " (pendingDiff)"));
                assertEq(curAlloc.effectBlock, 0, string.concat(err, " (effectBlock)"));
            } else {
                assertEq(
                    curAlloc.currentMagnitude, 
                    prevAlloc.currentMagnitude, 
                    string.concat(err, " (currentMagnitude)")
                );
                assertEq(
                    curAlloc.pendingDiff, 
                    prevAlloc.pendingDiff + int128(int64(allocateParams.newMagnitudes[i])), 
                    string.concat(err, " (pendingDiff)")
                );

                (, uint32 delay) = allocationManager.getAllocationDelay(address(operator));

                assertEq(
                    curAlloc.effectBlock, 
                    block.number + delay, 
                    string.concat(err, " (effectBlock)")
                );
            }

            // Check magnitudes
            Magnitudes memory curMagnitude = curMagnitudes[i];
            Magnitudes memory prevMagnitude = prevMagnitudes[i];

            // if (isAllocation) ...
            if (prevAlloc.currentMagnitude < allocateParams.newMagnitudes[i]) {
                uint256 delta = allocateParams.newMagnitudes[i] - prevAlloc.currentMagnitude;
                
                // FIXME: Fails on `testFuzz_deposit_delegate_allocate` when completed == true.
                if (!completed) {
                    assertEq(
                        curMagnitude.encumbered, 
                        prevMagnitude.encumbered + delta, 
                        string.concat(err, " (encumbered magnitude)")
                    );

                    assertEq(
                        curMagnitude.allocatable, 
                        prevMagnitude.allocatable - delta, 
                        string.concat(err, " (allocatable magnitude)")
                    );
                }
            } else {
                uint256 delta = prevAlloc.currentMagnitude - allocateParams.newMagnitudes[i];

                // if deallocated, and completed
                if (completed) {
                    assertEq(
                        curMagnitude.encumbered - delta, 
                        prevMagnitude.encumbered, 
                        string.concat(err, " (encumbered magnitude)")
                    );

                    assertEq(
                        curMagnitude.allocatable, 
                        prevMagnitude.allocatable + delta, 
                        string.concat(err, " (allocatable magnitude)")
                    );
                } else {
                    assertEq(
                        curMagnitude.encumbered, 
                        prevMagnitude.encumbered, 
                        string.concat(err, " (encumbered magnitude)")
                    );

                    assertEq(
                        curMagnitude.allocatable, 
                        prevMagnitude.allocatable, 
                        string.concat(err, " (allocatable magnitude)")
                    );
                }
            }

            // Max magnitude should not have mutated.
            assertEq(
                curMagnitude.max, 
                prevMagnitude.max, 
                string.concat(err, " (max magnitude)")
            );
        }
    }

    function assert_Snap_Allocations_Slashed(
        IAllocationManagerTypes.SlashingParams memory slashingParams,
        OperatorSet memory operatorSet,
        bool completed,
        string memory err
    ) internal {
        User op = User(payable(slashingParams.operator));
        
        IAllocationManagerTypes.Allocation[] memory curAllocs = _getAllocations(op, operatorSet, slashingParams.strategies);
        IAllocationManagerTypes.Allocation[] memory prevAllocs = _getPrevAllocations(op, operatorSet, slashingParams.strategies);
        Magnitudes[] memory curMagnitudes = _getMagnitudes(op, slashingParams.strategies);
        Magnitudes[] memory prevMagnitudes = _getPrevMagnitudes(op, slashingParams.strategies);

        (, uint32 delay) = allocationManager.getAllocationDelay(slashingParams.operator);

        for (uint i = 0; i < slashingParams.strategies.length; i++) {
            IAllocationManagerTypes.Allocation memory curAlloc = curAllocs[i];
            IAllocationManagerTypes.Allocation memory prevAlloc = prevAllocs[i]; 

            uint64 slashedMagnitude = uint64(uint256(prevAlloc.currentMagnitude).mulWadRoundUp(slashingParams.wadsToSlash[i]));
            
            // Check Allocations
            assertEq(
                curAlloc.currentMagnitude, 
                prevAlloc.currentMagnitude - slashedMagnitude, 
                string.concat(err, " (currentMagnitude)")
            );

            if (completed) {
                assertEq(curAlloc.pendingDiff, 0, string.concat(err, " (pendingDiff)"));
                assertEq(curAlloc.effectBlock, 0, string.concat(err, " (effectBlock)"));
            } else {
                assertEq(
                    curAlloc.currentMagnitude, 
                    prevAlloc.currentMagnitude - slashedMagnitude, 
                    string.concat(err, " (currentMagnitude)")
                );

                // If (isDeallocation) ... 
                if (prevAlloc.pendingDiff < 0) {
                    uint64 slashedPending =
                        uint64(uint256(uint128(-prevAlloc.pendingDiff)).mulWadRoundUp(slashingParams.wadsToSlash[i]));
                    
                    assertEq(
                        curAlloc.pendingDiff, 
                        prevAlloc.pendingDiff + int128(uint128(slashedPending)), 
                        string.concat(err, " (pendingDiff)")
                    );

                    delay = DEALLOCATION_DELAY + 1;
                }

                assertEq(
                    curAlloc.effectBlock, 
                    block.number + delay, 
                    string.concat(err, " (effectBlock)")
                );
            }

            // Check Magnitudes
            Magnitudes memory curMagnitude = curMagnitudes[i];
            Magnitudes memory prevMagnitude = prevMagnitudes[i];

            assertEq(
                curMagnitude.encumbered, 
                prevMagnitude.encumbered - slashedMagnitude, 
                string.concat(err, " (encumbered magnitude)")
            );

            assertEq(
                curMagnitude.allocatable, 
                prevMagnitude.allocatable, 
                string.concat(err, " (allocatable magnitude)")
            );

            assertEq(
                curMagnitude.max, 
                prevMagnitude.max - slashedMagnitude, 
                string.concat(err, " (max magnitude)")
            );
        }
    }

    function assert_HasUnderlyingTokenBalances_AfterSlash(
        User staker,
        IAllocationManagerTypes.AllocateParams memory allocateParams,
        IAllocationManagerTypes.SlashingParams memory slashingParams,
        uint[] memory expectedBalances,
        string memory err
    ) internal view {
        for (uint i; i < allocateParams.strategies.length; ++i) {
            IStrategy strat = allocateParams.strategies[i];
            
            uint balance = strat == BEACONCHAIN_ETH_STRAT 
                ? address(staker).balance 
                : strat.underlyingToken().balanceOf(address(staker));

            uint256 maxDelta = strat == BEACONCHAIN_ETH_STRAT ? 1 gwei : 3;

            if (slashingParams.strategies.contains(strat)) {
                uint256 wadToSlash = slashingParams.wadsToSlash[slashingParams.strategies.indexOf(strat)];

                expectedBalances[i] -= expectedBalances[i]
                    .mulWadRoundUp(allocateParams.newMagnitudes[i].mulWadRoundUp(wadToSlash));
            } 

            assertApproxEqAbs(expectedBalances[i], balance, maxDelta, err);
        }
    }

    function assert_Snap_StakerWithdrawableShares_AfterSlash(
        User staker,
        IAllocationManagerTypes.AllocateParams memory allocateParams,
        IAllocationManagerTypes.SlashingParams memory slashingParams,
        string memory err
    ) internal {
        uint[] memory curShares = _getWithdrawableShares(staker, allocateParams.strategies);
        uint[] memory prevShares = _getPrevWithdrawableShares(staker, allocateParams.strategies);

        for (uint i = 0; i < allocateParams.strategies.length; i++) {
            IStrategy strat = allocateParams.strategies[i];

            uint256 slashedShares = 0;

            if (slashingParams.strategies.contains(strat)) {
                uint wadToSlash = slashingParams.wadsToSlash[slashingParams.strategies.indexOf(strat)];
                slashedShares = prevShares[i].mulWadRoundUp(allocateParams.newMagnitudes[i].mulWadRoundUp(wadToSlash));
            }

            assertApproxEqAbs(prevShares[i] - slashedShares, curShares[i], 1, err);
        }
    }

    // TODO: slashable stake

    /*******************************************************************************
                        SNAPSHOT ASSERTIONS: OPERATOR SHARES
    *******************************************************************************/

    /// @dev Check that the operator has `addedShares` additional operator shares 
    // for each strategy since the last snapshot
    function assert_Snap_Added_OperatorShares(
        User operator, 
        IStrategy[] memory strategies, 
        uint[] memory addedShares,
        string memory err
    ) internal {
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
            assertEq(prevShares[i] - removedShares[i], curShares[i], err);
        }
    }

    /// @dev Check that the operator's shares in ALL strategies have not changed
    /// since the last snapshot
    function assert_Snap_Unchanged_OperatorShares(
        User operator,
        string memory err
    ) internal {
        IStrategy[] memory strategies = allStrats;

        uint[] memory curShares = _getOperatorShares(operator, strategies);
        // Use timewarp to get previous operator shares
        uint[] memory prevShares = _getPrevOperatorShares(operator, strategies);

        // For each strategy, check (prev == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i], curShares[i], err);
        }
    }

    function assert_Snap_Delta_OperatorShares(
        User operator, 
        IStrategy[] memory strategies, 
        int[] memory shareDeltas,
        string memory err
    ) internal {
        uint[] memory curShares = _getOperatorShares(operator, strategies);
        // Use timewarp to get previous operator shares
        uint[] memory prevShares = _getPrevOperatorShares(operator, strategies);

        // For each strategy, check (prev + added == cur)
        for (uint i = 0; i < strategies.length; i++) {
            uint expectedShares;
            if (shareDeltas[i] < 0) {
                expectedShares = prevShares[i] - uint(-shareDeltas[i]);
            } else {
                expectedShares = prevShares[i] + uint(shareDeltas[i]);
            }
            assertEq(expectedShares, curShares[i], err);
        }
    }

    /*******************************************************************************
                            SNAPSHOT ASSERTIONS: STAKER SHARES
    *******************************************************************************/

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

    function assert_Snap_Added_Staker_DepositShares(
        User staker, 
        IStrategy strat, 
        uint _addedShares,
        string memory err
    ) internal {
        IStrategy[] memory strategies = new IStrategy[](1);
        uint[] memory addedShares = new uint[](1);
        strategies[0] = strat;
        addedShares[0] = _addedShares;

        assert_Snap_Added_Staker_DepositShares(staker, strategies, addedShares, err);
    }

    /// @dev Check that the staker has `removedShares` fewer delegatable shares
    /// for each strategy since the last snapshot
    function assert_Snap_Removed_StakerDepositShares(
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

    function assert_Snap_Removed_StakerDepositShares(
        User staker, 
        IStrategy strat, 
        uint _removedShares,
        string memory err
    ) internal {
        IStrategy[] memory strategies = new IStrategy[](1);
        uint[] memory removedShares = new uint[](1);
        strategies[0] = strat;
        removedShares[0] = _removedShares;

        assert_Snap_Removed_StakerDepositShares(staker, strategies, removedShares, err);
    }

    /// @dev Check that the staker's withdrawable shares have decreased by `removedShares`
    function assert_Snap_Removed_StakerWithdrawableShares(
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

    function assert_Snap_Removed_StakerWithdrawableShares(
        User staker, 
        IStrategy strat, 
        uint _removedShares,
        string memory err
    ) internal {
        IStrategy[] memory strategies = new IStrategy[](1);
        uint[] memory removedShares = new uint[](1);
        strategies[0] = strat;
        removedShares[0] = _removedShares;

        assert_Snap_Removed_StakerWithdrawableShares(staker, strategies, removedShares, err);
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

        // For each strategy, check diff between (prev-removed) and curr is at most 1 gwei
        for (uint i = 0; i < strategies.length; i++) {
            assertApproxEqAbs(prevShares[i] - removedShares[i], curShares[i], 1e9, err);
        }
    }

    function assert_Snap_Removed_Staker_WithdrawableShares_AtLeast(
        User staker,
        IStrategy strat,
        uint removedShares,
        string memory err
    ) internal {
        IStrategy[] memory strategies = new IStrategy[](1);
        uint[] memory removedSharesArr = new uint[](1);
        strategies[0] = strat;
        removedSharesArr[0] = removedShares;

        assert_Snap_Removed_Staker_WithdrawableShares_AtLeast(staker, strategies, removedSharesArr, err);
    }

    /// @dev Check that the staker's delegatable shares in ALL strategies have not changed
    /// since the last snapshot
    function assert_Snap_Unchanged_StakerDepositShares(
        User staker,
        string memory err
    ) internal {
        IStrategy[] memory strategies = allStrats;

        uint[] memory curShares = _getStakerDepositShares(staker, strategies);
        // Use timewarp to get previous staker shares
        uint[] memory prevShares = _getPrevStakerDepositShares(staker, strategies);

        // For each strategy, check (prev == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i], curShares[i], err);
        }
    }

    function assert_Snap_Delta_StakerShares(
        User staker, 
        IStrategy[] memory strategies, 
        int[] memory shareDeltas,
        string memory err
    ) internal {
        int[] memory curShares = _getStakerDepositSharesInt(staker, strategies);
        // Use timewarp to get previous staker shares
        int[] memory prevShares = _getPrevStakerDepositSharesInt(staker, strategies);

        // For each strategy, check (prev + added == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] + shareDeltas[i], curShares[i], err);
        }
    }

    /*******************************************************************************
                        SNAPSHOT ASSERTIONS: STRATEGY SHARES
    *******************************************************************************/

    function assert_Snap_Removed_StrategyShares(
        IStrategy[] memory strategies,
        uint[] memory removedShares,
        string memory err
    ) internal {
        uint[] memory curShares = _getTotalStrategyShares(strategies);

        // Use timewarp to get previous strategy shares
        uint[] memory prevShares = _getPrevTotalStrategyShares(strategies);

        for (uint i = 0; i < strategies.length; i++) {
            // Ignore BeaconChainETH strategy since it doesn't keep track of global strategy shares
            if (strategies[i] == BEACONCHAIN_ETH_STRAT) {
                continue;
            }
            uint prevShare = prevShares[i];
            uint curShare = curShares[i];

            assertApproxEqAbs(prevShare - removedShares[i], curShare, 1, err);
        }
    }

    function assert_Snap_Unchanged_StrategyShares(
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        uint[] memory curShares = _getTotalStrategyShares(strategies);

        // Use timewarp to get previous strategy shares
        uint[] memory prevShares = _getPrevTotalStrategyShares(strategies);

        for (uint i = 0; i < strategies.length; i++) {
            uint prevShare = prevShares[i];
            uint curShare = curShares[i];

            assertEq(prevShare, curShare, err);
        }
    }

    /*******************************************************************************
                      SNAPSHOT ASSERTIONS: UNDERLYING TOKEN
    *******************************************************************************/

    /// @dev Check that the staker has `addedTokens` additional underlying tokens 
    // since the last snapshot
    function assert_Snap_Added_TokenBalances(
        User staker,
        IERC20[] memory tokens,
        uint[] memory addedTokens,
        string memory err
    ) internal {
        uint[] memory curTokenBalances = _getTokenBalances(staker, tokens);
        // Use timewarp to get previous token balances
        uint[] memory prevTokenBalances = _getPrevTokenBalances(staker, tokens);

        for (uint i = 0; i < tokens.length; i++) {
            uint prevBalance = prevTokenBalances[i];
            uint curBalance = curTokenBalances[i];

            assertApproxEqAbs(prevBalance + addedTokens[i], curBalance, 1, err);
        }
    }

    /// @dev Check that the staker has `removedTokens` fewer underlying tokens 
    // since the last snapshot
    function assert_Snap_Removed_TokenBalances(
        User staker,
        IStrategy[] memory strategies,
        uint[] memory removedTokens,
        string memory err
    ) internal {
        IERC20[] memory tokens = _getUnderlyingTokens(strategies);

        uint[] memory curTokenBalances = _getTokenBalances(staker, tokens);
        // Use timewarp to get previous token balances
        uint[] memory prevTokenBalances = _getPrevTokenBalances(staker, tokens);

        for (uint i = 0; i < tokens.length; i++) {
            uint prevBalance = prevTokenBalances[i];
            uint curBalance = curTokenBalances[i];

            assertEq(prevBalance - removedTokens[i], curBalance, err);
        }
    }

    /// @dev Check that the staker's underlying token balance for ALL tokens have
    /// not changed since the last snapshot
    function assert_Snap_Unchanged_TokenBalances(
        User staker,
        string memory err
    ) internal {
        IERC20[] memory tokens = allTokens;

        uint[] memory curTokenBalances = _getTokenBalances(staker, tokens);
        // Use timewarp to get previous token balances
        uint[] memory prevTokenBalances = _getPrevTokenBalances(staker, tokens);

        for (uint i = 0; i < tokens.length; i++) {
            assertEq(prevTokenBalances[i], curTokenBalances[i], err);
        }
    }

    /*******************************************************************************
                      SNAPSHOT ASSERTIONS: QUEUED WITHDRAWALS
    *******************************************************************************/

    function assert_Snap_Added_QueuedWithdrawals(
        User staker, 
        IDelegationManagerTypes.Withdrawal[] memory withdrawals,
        string memory err
    ) internal {
        uint curQueuedWithdrawals = _getCumulativeWithdrawals(staker);
        // Use timewarp to get previous cumulative withdrawals
        uint prevQueuedWithdrawals = _getPrevCumulativeWithdrawals(staker);

        assertEq(prevQueuedWithdrawals + withdrawals.length, curQueuedWithdrawals, err);
    }

    function assert_Snap_Added_QueuedWithdrawal(
        User staker, 
        IDelegationManagerTypes.Withdrawal memory /*withdrawal*/,
        string memory err
    ) internal {
        uint curQueuedWithdrawal = _getCumulativeWithdrawals(staker);
        // Use timewarp to get previous cumulative withdrawals
        uint prevQueuedWithdrawal = _getPrevCumulativeWithdrawals(staker);

        assertEq(prevQueuedWithdrawal + 1, curQueuedWithdrawal, err);
    }

    /*******************************************************************************
                         SNAPSHOT ASSERTIONS: EIGENPODS
    *******************************************************************************/

    function assert_Snap_Added_ActiveValidatorCount(
        User staker,
        uint addedValidators,
        string memory err
    ) internal {
        uint curActiveValidatorCount = _getActiveValidatorCount(staker);
        uint prevActiveValidatorCount = _getPrevActiveValidatorCount(staker);

        assertEq(prevActiveValidatorCount + addedValidators, curActiveValidatorCount, err);
    }

    function assert_Snap_Removed_ActiveValidatorCount(
        User staker,
        uint exitedValidators,
        string memory err
    ) internal {
        uint curActiveValidatorCount = _getActiveValidatorCount(staker);
        uint prevActiveValidatorCount = _getPrevActiveValidatorCount(staker);

        assertEq(curActiveValidatorCount + exitedValidators, prevActiveValidatorCount, err);
    }

    function assert_Snap_Unchanged_ActiveValidatorCount(
        User staker,
        string memory err
    ) internal {
        uint curActiveValidatorCount = _getActiveValidatorCount(staker);
        uint prevActiveValidatorCount = _getPrevActiveValidatorCount(staker);

        assertEq(curActiveValidatorCount, prevActiveValidatorCount, err);
    }

    function assert_Snap_Added_ActiveValidators(
        User staker,
        uint40[] memory addedValidators,
        string memory err
    ) internal {
        bytes32[] memory pubkeyHashes = beaconChain.getPubkeyHashes(addedValidators);

        IEigenPodTypes.VALIDATOR_STATUS[] memory curStatuses = _getValidatorStatuses(staker, pubkeyHashes);
        IEigenPodTypes.VALIDATOR_STATUS[] memory prevStatuses = _getPrevValidatorStatuses(staker, pubkeyHashes);

        for (uint i = 0; i < curStatuses.length; i++) {
            assertTrue(prevStatuses[i] == IEigenPodTypes.VALIDATOR_STATUS.INACTIVE, err);
            assertTrue(curStatuses[i] == IEigenPodTypes.VALIDATOR_STATUS.ACTIVE, err);
        }
    }

    function assert_Snap_Removed_ActiveValidators(
        User staker,
        uint40[] memory exitedValidators,
        string memory err
    ) internal {
        bytes32[] memory pubkeyHashes = beaconChain.getPubkeyHashes(exitedValidators);

        IEigenPodTypes.VALIDATOR_STATUS[] memory curStatuses = _getValidatorStatuses(staker, pubkeyHashes);
        IEigenPodTypes.VALIDATOR_STATUS[] memory prevStatuses = _getPrevValidatorStatuses(staker, pubkeyHashes);

        for (uint i = 0; i < curStatuses.length; i++) {
            assertTrue(prevStatuses[i] == IEigenPodTypes.VALIDATOR_STATUS.ACTIVE, err);
            assertTrue(curStatuses[i] == IEigenPodTypes.VALIDATOR_STATUS.WITHDRAWN, err);
        }
    }

    function assert_Snap_Created_Checkpoint(
        User staker,
        string memory err
    ) internal {
        uint64 curCheckpointTimestamp = _getCheckpointTimestamp(staker);
        uint64 prevCheckpointTimestamp = _getPrevCheckpointTimestamp(staker);

        assertEq(prevCheckpointTimestamp, 0, err);
        assertTrue(curCheckpointTimestamp != 0, err);
    }

    function assert_Snap_Removed_Checkpoint(
        User staker,
        string memory err
    ) internal {
        uint64 curCheckpointTimestamp = _getCheckpointTimestamp(staker);
        uint64 prevCheckpointTimestamp = _getPrevCheckpointTimestamp(staker);

        assertEq(curCheckpointTimestamp, 0, err);
        assertTrue(prevCheckpointTimestamp != 0, err);
    }

    function assert_Snap_Unchanged_Checkpoint(
        User staker,
        string memory err
    ) internal {
        uint64 curCheckpointTimestamp = _getCheckpointTimestamp(staker);
        uint64 prevCheckpointTimestamp = _getPrevCheckpointTimestamp(staker);

        assertEq(curCheckpointTimestamp, prevCheckpointTimestamp, err);
    }

    function assert_Snap_Updated_LastCheckpoint(
        User staker,
        string memory err
    ) internal {
        // Sorry for the confusing naming... the pod variable is called `lastCheckpointTimestamp`
        uint64 curLastCheckpointTimestamp = _getLastCheckpointTimestamp(staker);
        uint64 prevLastCheckpointTimestamp = _getPrevLastCheckpointTimestamp(staker);

        assertTrue(curLastCheckpointTimestamp > prevLastCheckpointTimestamp, err);
    }

    function assert_Snap_Added_PodBalanceToWithdrawable(
        User staker,
        string memory err
    ) internal {
        uint64 curWithdrawableRestakedGwei = _getWithdrawableRestakedGwei(staker);
        uint64 prevWithdrawableRestakedGwei = _getPrevWithdrawableRestakedGwei(staker);

        uint64 prevCheckpointPodBalanceGwei = _getPrevCheckpointPodBalanceGwei(staker);

        assertEq(prevWithdrawableRestakedGwei + prevCheckpointPodBalanceGwei, curWithdrawableRestakedGwei, err);
    }

    function assert_Snap_Added_WithdrawableGwei(
        User staker,
        uint64 addedGwei,
        string memory err
    ) internal {
        uint64 curWithdrawableRestakedGwei = _getWithdrawableRestakedGwei(staker);
        uint64 prevWithdrawableRestakedGwei = _getPrevWithdrawableRestakedGwei(staker);

        assertEq(prevWithdrawableRestakedGwei + addedGwei, curWithdrawableRestakedGwei, err);
    }

    function assert_Snap_Added_BalanceExitedGwei(
        User staker,
        uint64 addedGwei,
        string memory err
    ) internal {
        uint64 curCheckpointTimestamp = _getCheckpointTimestamp(staker);
        uint64 prevCheckpointTimestamp = _getPrevCheckpointTimestamp(staker);

        // If we just finalized a checkpoint, that's the timestamp we want to use
        // to look up checkpoint balances exited
        uint64 targetTimestamp = curCheckpointTimestamp;
        if (curCheckpointTimestamp != prevCheckpointTimestamp) {
            targetTimestamp = prevCheckpointTimestamp;
        }

        uint64 curExitedBalanceGwei = _getCheckpointBalanceExited(staker, targetTimestamp);
        uint64 prevExitedBalanceGwei = _getPrevCheckpointBalanceExited(staker, targetTimestamp);

        assertEq(prevExitedBalanceGwei + addedGwei, curExitedBalanceGwei, err);
    }

    /*******************************************************************************
                                UTILITY METHODS
    *******************************************************************************/
    
    function _randWadToSlash() internal returns (uint) {
        return _randUint({ min: 0.01 ether, max: 1 ether });
    }

    function _randStrategiesAndWadsToSlash(
        OperatorSet memory operatorSet
    ) internal returns (IStrategy[] memory strategies, uint[] memory wadsToSlash) {
        // Get list of all strategies in an operator set.
        strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);

        // Randomly select a subset of strategies to slash.
        uint len = _randUint({ min: 1, max: strategies.length });

        // Update length of strategies array.
        assembly {
            mstore(strategies, len)
        }
        
        wadsToSlash = new uint[](len);
        
        // Randomly select a `wadToSlash` for each strategy.
        for (uint i; i < len; ++i) {
            wadsToSlash[i] = _randWadToSlash();
        }

        return (strategies.sort(), wadsToSlash);
    }
    
    function _randMagnitudes(uint64 sum, uint256 len) internal returns (uint64[] memory magnitudes) {
        magnitudes = new uint64[](len);

        if (sum == 0 || len == 0) return magnitudes;
        
        uint64 remaining = sum;

        for (uint256 i; i < len; ++i) {
            if (i == len - 1) {
                magnitudes[i] = remaining;
            } else {
                magnitudes[i] = uint64(_randUint(0, remaining / (len - i)));
                remaining -= magnitudes[i];
            }
        }
    }

    function _randWithdrawal(
        IStrategy[] memory strategies, 
        uint[] memory shares
    ) internal returns (IStrategy[] memory, uint[] memory) {
        uint stratsToWithdraw = _randUint({ min: 1, max: strategies.length });

        IStrategy[] memory withdrawStrats = new IStrategy[](stratsToWithdraw);
        uint[] memory withdrawShares = new uint[](stratsToWithdraw);

        for (uint i = 0; i < stratsToWithdraw; i++) {
            uint sharesToWithdraw;

            if (strategies[i] == BEACONCHAIN_ETH_STRAT) {
                // For native eth, withdraw a random amount of gwei (at least 1)
                uint portion = _randUint({ min: 1, max: shares[i] / GWEI_TO_WEI });
                portion *= GWEI_TO_WEI;

                sharesToWithdraw = shares[i] - portion;
            } else {
                // For LSTs, withdraw a random amount of shares (at least 1)
                uint portion = _randUint({ min: 1, max: shares[i] });

                sharesToWithdraw = shares[i] - portion;
            }

            withdrawStrats[i] = strategies[i];
            withdrawShares[i] = sharesToWithdraw;
        }

        return (withdrawStrats, withdrawShares);
    }

    /**
     * Helpful getters:
     */
    function _randBalanceUpdate(
        User staker,
        IStrategy[] memory strategies
    ) internal returns (int[] memory, int[] memory, int[] memory) {

        int[] memory tokenDeltas = new int[](strategies.length);
        int[] memory stakerShareDeltas = new int[](strategies.length);
        int[] memory operatorShareDeltas = new int[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                // For native ETH, we're either going to slash the staker's validators,
                // or award them consensus rewards. In either case, the magnitude of
                // the balance update depends on the staker's active validator count
                uint activeValidatorCount = staker.pod().activeValidatorCount();
                int64 deltaGwei;
                if (_randBool()) {
                    uint40[] memory validators = staker.getActiveValidators();
                    emit log_named_uint("slashing validators", validators.length);

                    deltaGwei = -int64(beaconChain.slashValidators(validators));
                    beaconChain.advanceEpoch_NoRewards();

                    emit log_named_int("slashed amount", deltaGwei);
                } else {
                    emit log("generating consensus rewards for validators");

                    deltaGwei = int64(uint64(activeValidatorCount) * beaconChain.CONSENSUS_REWARD_AMOUNT_GWEI());
                    beaconChain.advanceEpoch_NoWithdraw();
                }
                
                tokenDeltas[i] = int(deltaGwei) * int(GWEI_TO_WEI);

                stakerShareDeltas[i] = tokenDeltas[i];
                operatorShareDeltas[i] = _calcNativeETHOperatorShareDelta(staker, stakerShareDeltas[i]);

                emit log_named_int("beacon balance delta (gwei): ", deltaGwei);
                emit log_named_int("staker share delta (gwei): ", stakerShareDeltas[i] / int(GWEI_TO_WEI));
                emit log_named_int("operator share delta (gwei): ", operatorShareDeltas[i] / int(GWEI_TO_WEI));
            } else {
                // For LSTs, mint a random token amount
                uint portion = _randUint({ min: MIN_BALANCE, max: MAX_BALANCE });
                StdCheats.deal(address(strat.underlyingToken()), address(staker), portion);

                int delta = int(portion);
                tokenDeltas[i] = delta;
                stakerShareDeltas[i] = int(strat.underlyingToShares(uint(delta)));
                operatorShareDeltas[i] = int(strat.underlyingToShares(uint(delta)));
            }
        }
        return (tokenDeltas, stakerShareDeltas, operatorShareDeltas);
    }

    function _calcNativeETHOperatorShareDelta(User staker, int shareDelta) internal view returns (int) {
        int curPodOwnerShares = eigenPodManager.podOwnerDepositShares(address(staker));
        int newPodOwnerShares = curPodOwnerShares + shareDelta;

        if (curPodOwnerShares <= 0) {
            // if the shares started negative and stayed negative, then there cannot have been an increase in delegateable shares
            if (newPodOwnerShares <= 0) {
                return 0;
            // if the shares started negative and became positive, then the increase in delegateable shares is the ending share amount
            } else {
                return newPodOwnerShares;
            }
        } else {
            // if the shares started positive and became negative, then the decrease in delegateable shares is the starting share amount
            if (newPodOwnerShares <= 0) {
                return (-curPodOwnerShares);
            // if the shares started positive and stayed positive, then the change in delegateable shares
            // is the difference between starting and ending amounts
            } else {
                return (newPodOwnerShares - curPodOwnerShares);
            }
        }
    }

    /// @dev For some strategies/underlying token balances, calculate the expected shares received
    /// from depositing all tokens
    function _calculateExpectedShares(IStrategy[] memory strategies, uint[] memory tokenBalances) internal returns (uint[] memory) {
        uint[] memory expectedShares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            uint tokenBalance = tokenBalances[i];
            if (strat == BEACONCHAIN_ETH_STRAT) {
                expectedShares[i] = tokenBalance;
            } else {
                expectedShares[i] = strat.underlyingToShares(tokenBalance);
            }
        }

        return expectedShares;
    }

    /// @dev For some strategies/underlying token balances, calculate the expected shares received
    /// from depositing all tokens
    function _calculateExpectedTokens(IStrategy[] memory strategies, uint[] memory shares) internal returns (uint[] memory) {
        uint[] memory expectedTokens = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                expectedTokens[i] = shares[i];
            } else {
                expectedTokens[i] = strat.sharesToUnderlying(shares[i]);
            }
        }

        return expectedTokens;
    }

    function _getWithdrawalHashes(
        IDelegationManagerTypes.Withdrawal[] memory withdrawals
    ) internal view returns (bytes32[] memory) {
        bytes32[] memory withdrawalRoots = new bytes32[](withdrawals.length);

        for (uint i = 0; i < withdrawals.length; i++) {
            withdrawalRoots[i] = delegationManager.calculateWithdrawalRoot(withdrawals[i]);
        }

        return withdrawalRoots;
    }

    /// @dev Converts a list of strategies to underlying tokens
    function _getUnderlyingTokens(IStrategy[] memory strategies) internal view returns (IERC20[] memory) {
        IERC20[] memory tokens = new IERC20[](strategies.length);

        for (uint i = 0; i < tokens.length; i++) {
            IStrategy strat = strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                tokens[i] = NATIVE_ETH;
            } else {
                tokens[i] = strat.underlyingToken();
            }
        }

        return tokens;
    }

    modifier timewarp() {
        uint curState = timeMachine.travelToLast();
        _;
        timeMachine.travel(curState);
    }

    /// @dev Rolls forward by the minimum withdrawal delay blocks.
    function _rollBlocksForCompleteWithdrawals(IDelegationManagerTypes.Withdrawal[] memory withdrawals) internal {     
        uint256 latest;
        for (uint i = 0; i < withdrawals.length; ++i) {
            if (withdrawals[i].startBlock > latest) latest = withdrawals[i].startBlock;
        }
        cheats.roll(latest + delegationManager.minWithdrawalDelayBlocks() + 1);
    }

    /// @dev Rolls forward by the default allocation delay blocks.
    function _rollBlocksForCompleteAllocation(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies
    ) internal {
        uint256 latest;
        for (uint i = 0; i < strategies.length; ++i) {
            uint effectBlock = allocationManager.getAllocation(address(operator), operatorSet, strategies[i]).effectBlock;
            if (effectBlock > latest) latest = effectBlock;
        }
        cheats.roll(latest + 1);
    }

    /// @dev Rolls forward by the default allocation delay blocks.
    function _rollBlocksForCompleteAllocation(
        User operator,
        OperatorSet[] memory operatorSets,
        IStrategy[] memory strategies
    ) internal {
        uint256 latest;
        for (uint i = 0; i < operatorSets.length; ++i) {
            for (uint j = 0; j < strategies.length; ++j) {
                uint effectBlock = allocationManager.getAllocation(address(operator), operatorSets[i], strategies[j]).effectBlock;
                if (effectBlock > latest) latest = effectBlock;
            }
        }
        cheats.roll(latest + 1);
    }

    /// @dev Uses timewarp modifier to get the operator set strategy allocations at the last snapshot.
    function _getPrevAllocations(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies
    ) internal timewarp() returns (IAllocationManagerTypes.Allocation[] memory) {
        return _getAllocations(operator, operatorSet, strategies);
    }

    /// @dev Looks up each strategy for an operator set and returns a list of operator allocations.
    function _getAllocations(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies
    ) internal view returns (IAllocationManagerTypes.Allocation[] memory allocations) {
        allocations = new IAllocationManagerTypes.Allocation[](strategies.length);
        for (uint i = 0; i < strategies.length; ++i) {
            allocations[i] = allocationManager.getAllocation(address(operator), operatorSet, strategies[i]);
        }
    }

    struct Magnitudes {
        uint256 encumbered;
        uint256 allocatable;
        uint256 max;
    }

    function _getPrevMagnitudes(
        User operator,
        IStrategy[] memory strategies
    ) internal timewarp() returns (Magnitudes[] memory) {
        return _getMagnitudes(operator, strategies);
    }

    function _getMagnitudes(
        User operator,
        IStrategy[] memory strategies
    ) internal view returns (Magnitudes[] memory magnitudes) {
        magnitudes = new Magnitudes[](strategies.length);
        for (uint i = 0; i < strategies.length; ++i) {
            magnitudes[i] = Magnitudes({
                encumbered: allocationManager.encumberedMagnitude(address(operator), strategies[i]),
                allocatable: allocationManager.getAllocatableMagnitude(address(operator), strategies[i]),
                max: allocationManager.getMaxMagnitude(address(operator), strategies[i])
            });
        }
    }

    /// @dev Uses timewarp modifier to get operator shares at the last snapshot
    function _getPrevOperatorShares(
        User operator, 
        IStrategy[] memory strategies
    ) internal timewarp() returns (uint[] memory) {
        return _getOperatorShares(operator, strategies);
    }

    /// @dev Looks up each strategy and returns a list of the operator's shares
    function _getOperatorShares(User operator, IStrategy[] memory strategies) internal view returns (uint[] memory) {
        uint[] memory curShares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            curShares[i] = delegationManager.operatorShares(address(operator), strategies[i]);
        }

        return curShares;
    }

    /// @dev Uses timewarp modifier to get staker shares at the last snapshot
    function _getPrevStakerDepositShares(
        User staker, 
        IStrategy[] memory strategies
    ) internal timewarp() returns (uint[] memory) {
        return _getStakerDepositShares(staker, strategies);
    }

    /// @dev Looks up each strategy and returns a list of the staker's shares
    function _getStakerDepositShares(User staker, IStrategy[] memory strategies) internal view returns (uint[] memory) {
        uint[] memory curShares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                // This method should only be used for tests that handle positive
                // balances. Negative balances are an edge case that require
                // the own tests and helper methods.
                int shares;
                if (forkType != LOCAL && !isUpgraded) {
                    shares = int(IEigenPodManager_DeprecatedM2(address(eigenPodManager)).podOwnerShares(address(staker)));
                } else {
                    shares = int(eigenPodManager.podOwnerDepositShares(address(staker)));
                }

                if (shares < 0) {
                    revert("_getStakerDepositShares: negative shares");
                }

                curShares[i] = uint(shares);
            } else {
                if (forkType != LOCAL && !isUpgraded) {
                    curShares[i] = IStrategyManager_DeprecatedM2(address(strategyManager)).stakerStrategyShares(address(staker), strat);
                } else {
                    curShares[i] = strategyManager.stakerDepositShares(address(staker), strat);
                }
            }
        }

        return curShares;
    }

    /// @dev Uses timewarp modifier to get staker shares at the last snapshot
    function _getPrevStakerDepositSharesInt(
        User staker, 
        IStrategy[] memory strategies
    ) internal timewarp() returns (int[] memory) {
        return _getStakerDepositSharesInt(staker, strategies);
    }

    /// @dev Looks up each strategy and returns a list of the staker's shares
    function _getStakerDepositSharesInt(User staker, IStrategy[] memory strategies) internal view returns (int[] memory) {
        int[] memory curShares = new int[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            if (strat == BEACONCHAIN_ETH_STRAT) {
                curShares[i] = eigenPodManager.podOwnerDepositShares(address(staker));
            } else {
                curShares[i] = int(strategyManager.stakerDepositShares(address(staker), strat));
            }
        }

        return curShares;
    }

    function _getPrevStakerWithdrawableShares(User staker, IStrategy[] memory strategies) internal timewarp() returns (uint[] memory) {
        return _getStakerWithdrawableShares(staker, strategies);
    }

    function _getStakerWithdrawableShares(User staker, IStrategy[] memory strategies) internal view returns (uint[] memory) {
        (uint256[] memory withdrawableShares, ) = delegationManager.getWithdrawableShares(address(staker), strategies);
        return withdrawableShares; 
    }

    /// @dev Uses timewarp modifier to get staker beacon chain scaling factor at the last snapshot
    function _getPrevBeaconChainSlashingFactor(User staker) internal timewarp() returns (uint64) {
        return _getBeaconChainSlashingFactor(staker);
    }

    /// @dev Looks up the staker's beacon chain scaling factor
    function _getBeaconChainSlashingFactor(User staker) internal view returns (uint64) {
        return eigenPodManager.beaconChainSlashingFactor(address(staker));
    }

    function _getPrevCumulativeWithdrawals(User staker) internal timewarp() returns (uint) {
        return _getCumulativeWithdrawals(staker);
    }

    function _getCumulativeWithdrawals(User staker) internal view returns (uint) {
        return delegationManager.cumulativeWithdrawalsQueued(address(staker));
    }

    function _getPrevTokenBalances(User staker, IERC20[] memory tokens) internal timewarp() returns (uint[] memory) {
        return _getTokenBalances(staker, tokens);
    }

    function _getTokenBalances(User staker, IERC20[] memory tokens) internal view returns (uint[] memory) {
        uint[] memory balances = new uint[](tokens.length);
        
        for (uint i = 0; i < tokens.length; i++) {
            if (tokens[i] == NATIVE_ETH) {
                balances[i] = address(staker).balance;
            } else {
                balances[i] = tokens[i].balanceOf(address(staker));
            }
        }

        return balances;
    }

    function _getPrevTotalStrategyShares(IStrategy[] memory strategies) internal timewarp() returns (uint[] memory) {
        return _getTotalStrategyShares(strategies);
    }

    function _getTotalStrategyShares(IStrategy[] memory strategies) internal view returns (uint[] memory) {
        uint[] memory shares = new uint[](strategies.length);

        for (uint i = 0; i < strategies.length; i++) {
            if (strategies[i] != BEACONCHAIN_ETH_STRAT) {
                shares[i] = strategies[i].totalShares();
            }
            // BeaconChainETH strategy doesn't keep track of global strategy shares, so we ignore
        }

        return shares;
    }

    function _getPrevWithdrawableShares(User staker, IStrategy[] memory strategies) internal timewarp() returns (uint[] memory) {
        return _getWithdrawableShares(staker, strategies);
    }

    function _getWithdrawableShares(User staker, IStrategy[] memory strategies) internal view returns (uint[] memory withdrawableShares) {
        (withdrawableShares, ) =  delegationManager.getWithdrawableShares(address(staker), strategies);
    }

    function _getActiveValidatorCount(User staker) internal view returns (uint) {
        EigenPod pod = staker.pod();
        return pod.activeValidatorCount();
    }

    function _getPrevActiveValidatorCount(User staker) internal timewarp() returns (uint) {
        return _getActiveValidatorCount(staker);
    }

    function _getValidatorStatuses(User staker, bytes32[] memory pubkeyHashes) internal view returns (IEigenPodTypes.VALIDATOR_STATUS[] memory) {
        EigenPod pod = staker.pod();
        IEigenPodTypes.VALIDATOR_STATUS[] memory statuses = new IEigenPodTypes.VALIDATOR_STATUS[](pubkeyHashes.length);

        for (uint i = 0; i < statuses.length; i++) {
            statuses[i] = pod.validatorStatus(pubkeyHashes[i]);
        }

        return statuses;
    }

    function _getPrevValidatorStatuses(User staker, bytes32[] memory pubkeyHashes) internal timewarp() returns (IEigenPodTypes.VALIDATOR_STATUS[] memory) {
        return _getValidatorStatuses(staker, pubkeyHashes);
    }

    function _getCheckpointTimestamp(User staker) internal view returns (uint64) {
        EigenPod pod = staker.pod();
        return pod.currentCheckpointTimestamp();
    }

    function _getPrevCheckpointTimestamp(User staker) internal timewarp() returns (uint64) {
        return _getCheckpointTimestamp(staker);
    }

    function _getLastCheckpointTimestamp(User staker) internal view returns (uint64) {
        EigenPod pod = staker.pod();
        return pod.lastCheckpointTimestamp();
    }

    function _getPrevLastCheckpointTimestamp(User staker) internal timewarp() returns (uint64) {
        return _getLastCheckpointTimestamp(staker);
    }

    function _getWithdrawableRestakedGwei(User staker) internal view returns (uint64) {
        EigenPod pod = staker.pod();
        return pod.withdrawableRestakedExecutionLayerGwei();
    }

    function _getPrevWithdrawableRestakedGwei(User staker) internal timewarp() returns (uint64) {
        return _getWithdrawableRestakedGwei(staker);
    }

    function _getCheckpointPodBalanceGwei(User staker) internal view returns (uint64) {
        if (forkType != LOCAL && !isUpgraded) {
            IEigenPod_DeprecatedM2 pod = IEigenPod_DeprecatedM2(address(staker.pod()));
            return uint64(pod.currentCheckpoint().podBalanceGwei);
        } else {
            EigenPod pod = staker.pod();
            return uint64(pod.currentCheckpoint().podBalanceGwei);
        }
    }

    function _getPrevCheckpointPodBalanceGwei(User staker) internal timewarp() returns (uint64) {
        return _getCheckpointPodBalanceGwei(staker);
    }

    function _getCheckpointBalanceExited(User staker, uint64 checkpointTimestamp) internal view returns (uint64) {
        EigenPod pod = staker.pod();
        return pod.checkpointBalanceExitedGwei(checkpointTimestamp);
    }

    function _getPrevCheckpointBalanceExited(User staker, uint64 checkpointTimestamp) internal timewarp() returns (uint64) {
        return _getCheckpointBalanceExited(staker, checkpointTimestamp);
    }
}