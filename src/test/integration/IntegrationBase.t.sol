// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

import "src/contracts/libraries/BeaconChainProofs.sol";
import "src/contracts/libraries/SlashingLib.sol";

import "src/test/integration/TypeImporter.t.sol";
import "src/test/integration/IntegrationDeployer.t.sol";
import "src/test/integration/TimeMachine.t.sol";
import "src/test/integration/users/User.t.sol";
import "src/test/integration/users/User_M1.t.sol";

abstract contract IntegrationBase is IntegrationDeployer, TypeImporter {
    using StdStyle for *;
    using SlashingLib for *;
    using Math for uint256;
    using Strings for *;
    using print for *;

    using ArrayLib for *;

    uint numStakers = 0;
    uint numOperators = 0;
    uint numAVSs = 0;

    // Lists of operators created prev the m2 (not slashing) upgrade
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
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) 
            = _randUser(_getStakerName());

        if (!isUpgraded) stakersToMigrate.push(staker);

        assert_HasUnderlyingTokenBalances(staker, strategies, tokenBalances, "_newRandomStaker: failed to award token balances");
        return (staker, strategies, tokenBalances);
    }

    /// Given a list of strategies, creates a new user with random token balances in each underlying token
    function _newStaker(IStrategy[] memory strategies) internal returns (User, uint[] memory) {
        (User staker, uint[] memory tokenBalances) 
            = _randUser(_getStakerName(), strategies);

        if (!isUpgraded) stakersToMigrate.push(staker);

        assert_HasUnderlyingTokenBalances(staker, strategies, tokenBalances, "_newStaker: failed to award token balances");
        return (staker, tokenBalances);
    }

    /**
     * @dev Create a new operator according to configured random variants.
     * This user will immediately deposit their randomized assets into eigenlayer.
     */
    function _newRandomOperator() internal returns (User, IStrategy[] memory, uint[] memory) {
        /// TODO: Allow operators to have ETH
        (User operator, IStrategy[] memory strategies, uint[] memory tokenBalances) 
            = _randUser_NoETH(_getOperatorName());

        /// Operators are created with all assets already deposited
        uint[] memory addedShares = _calculateExpectedShares(strategies, tokenBalances);
        operator.depositIntoEigenlayer(strategies, tokenBalances);

        /// Registration flow differs for M2 vs Slashing release
        if (!isUpgraded) {
            User_M2(payable(operator)).registerAsOperator_M2();

            operatorsToMigrate.push(operator);
        } else {
            operator.registerAsOperator();

            rollForward({blocks: ALLOCATION_CONFIGURATION_DELAY + 1});
        }
        
        assert_Snap_Added_OperatorShares(operator, strategies, addedShares, "_newRandomOperator: failed to award shares to operator");
        assertTrue(delegationManager.isOperator(address(operator)), "_newRandomOperator: operator should be registered");
        assertEq(delegationManager.delegatedTo(address(operator)), address(operator), "_newRandomOperator: should be self-delegated");
        return (operator, strategies, tokenBalances);
    }

    /// @dev Creates a new operator with no assets
    function _newRandomOperator_NoAssets() internal returns (User) {
        User operator = _randUser_NoAssets(_getOperatorName());

        /// Registration flow differs for M2 vs Slashing release
        if (!isUpgraded) {
            User_M2(payable(operator)).registerAsOperator_M2();

            operatorsToMigrate.push(operator);
        } else {
            operator.registerAsOperator();

            rollForward({blocks: ALLOCATION_CONFIGURATION_DELAY + 1});
        }

        assertTrue(delegationManager.isOperator(address(operator)), "_newRandomOperator: operator should be registered");
        assertEq(delegationManager.delegatedTo(address(operator)), address(operator), "_newRandomOperator: should be self-delegated");
        return operator;
    }

    /// @dev Name a newly-created staker ("staker1", "staker2", ...)
    function _getStakerName() private returns (string memory) {
        numStakers++;

        string memory stakerNum = cheats.toString(numStakers);
        string memory namePrefix = isUpgraded ? "staker" : "m2-staker";

        return string.concat(namePrefix, stakerNum);
    }

    /// @dev Name a newly-created operator ("operator1", "operator2", ...)
    function _getOperatorName() private returns (string memory) {
        numOperators++;

        string memory operatorNum = cheats.toString(numOperators);
        string memory namePrefix = isUpgraded ? "operator" : "m2-operator";

        return string.concat(namePrefix, operatorNum);
    }

    function _newRandomAVS() internal returns (AVS avs, OperatorSet[] memory operatorSets) {
        string memory avsName = string.concat("avs", numAVSs.toString());
        avs = _genRandAVS(avsName);
        avs.updateAVSMetadataURI("https://example.com");
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
    
    /// @dev Choose a random subset of validators (selects AT LEAST ONE)
    function _choose(uint40[] memory validators) internal returns (uint40[] memory) {
        uint _rand = _randUint({ min: 1, max: (2**validators.length) - 1 });

        uint40[] memory result = new uint40[](validators.length);
        uint newLen;
        for (uint i = 0; i < validators.length; i++) {
            // if bit set, add validator
            if (_rand >> i & 1 == 1) {
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
    ) internal {
        (, uint[] memory actualShares) = _getStakerDepositShares(user, strategies);
        for(uint i = 0; i < strategies.length; i++) {
            assertApproxEqAbs(expectedShares[i], actualShares[i], 1, err);
        }
    }

    /// @dev Check that all the staker's deposit shares have been removed
    function assert_RemovedAll_Staker_DepositShares(
        User user,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (, uint[] memory depositShares) = _getStakerDepositShares(user, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(depositShares[i], 0, err);
        }
    }

    /// @dev Check that all the staker's withdrawable shares have been removed
    function assert_RemovedAll_Staker_WithdrawableShares(
        User staker, 
        IStrategy[] memory strategies, 
        string memory err
    ) internal {
        (, uint[] memory curShares) = _getStakerWithdrawableShares(staker, strategies);
        // For each strategy, check all shares have been withdrawn
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(0, curShares[i], err);
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
        Withdrawal[] memory withdrawals,
        bytes32[] memory withdrawalRoots,
        string memory err
    ) internal view {
        for (uint i = 0; i < withdrawals.length; i++) {
            assert_ValidWithdrawalHash(withdrawals[i], withdrawalRoots[i], err);
        }
    }

    function assert_ValidWithdrawalHash(
        Withdrawal memory withdrawal,
        bytes32 withdrawalRoot,
        string memory err
    ) internal view {
        assertEq(withdrawalRoot, delegationManager.calculateWithdrawalRoot(withdrawal), err);
    }

    function assert_StakerStrategyListEmpty(
        User staker,
        string memory err
    ) internal view {
        IStrategy[] memory strategies = _getStakerStrategyList(staker);
        assertEq(strategies.length, 0, err);
    }

    function assert_StrategyNotInStakerStrategyList(
        User staker,
        IStrategy strategy,
        string memory err
    ) internal view {
        // BEACONCHAIN_ETH_STRAT is not in the staker's strategy list
        if (strategy == BEACONCHAIN_ETH_STRAT) return;       
        
        IStrategy[] memory strategies = _getStakerStrategyList(staker);
        assertFalse(strategies.contains(strategy), err);
    }

    function assert_StrategiesInStakerStrategyList(
        User staker,
        IStrategy[] memory strategies,
        string memory err
    ) internal view {
        for(uint i = 0; i < strategies.length; i++) {
            assert_StrategyInStakerStrategyList(staker, strategies[i], err);
        }
    }

    function assert_StrategyInStakerStrategyList(
        User staker,
        IStrategy strategy,
        string memory err
    ) internal view {
        // BEACONCHAIN_ETH_STRAT is not in the staker's strategy list
        if (strategy == BEACONCHAIN_ETH_STRAT) return;

        IStrategy[] memory strategies = _getStakerStrategyList(staker);
        assertTrue(strategies.contains(strategy), err);
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
        console.log("proofsRemaining: ", pod.currentCheckpoint().proofsRemaining);
        console.log("activeValidatorCount: ", pod.activeValidatorCount());
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

    function assert_MaxEqualsAllocatablePlusEncumbered(
        User operator,
        string memory err
    ) internal {
        (, Magnitudes[] memory current) = _getMagnitudes(operator, allStrats);
        for (uint i = 0; i < allStrats.length; i++) {
            assertEq(current[i].max, current[i].encumbered + current[i].allocatable, err);
        }
    }

    function assert_CurMinSlashableEqualsMinAllocated(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (, uint[] memory minSlashableStake) = _getMinSlashableStake(operator, operatorSet, strategies);
        (, uint[] memory minAllocatedStake) = _getAllocatedStake(operator, operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(minSlashableStake[i], minAllocatedStake[i], err);
        }
    }

    function assert_MaxMagsEqualMaxMagsAtCurrentBlock(
        User operator,
        IStrategy[] memory strategies,
        string memory err
    ) internal view {
        uint64[] memory maxMagnitudes = _getMaxMagnitudes(operator, strategies);
        uint64[] memory maxAtCurrentBlock = _getMaxMagnitudes(operator, strategies, uint32(block.number));

        for (uint i = 0; i < strategies.length; i++) {
            assertEq(maxMagnitudes[i], maxAtCurrentBlock[i], err);
        }
    }

    function assert_CurrentMagnitude(
        User operator,
        AllocateParams memory params,
        string memory err
    ) internal {
        (, Allocation[] memory allocations) = _getAllocations(operator, params.operatorSet, params.strategies);
        for (uint i = 0; i < allocations.length; i++) {
            assertEq(allocations[i].currentMagnitude, params.newMagnitudes[i], err);
        }
    }

    function assert_NoPendingModification(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (, Allocation[] memory allocations) = _getAllocations(operator, operatorSet, strategies);

        for (uint i = 0; i < allocations.length; i++) {
            assertEq(0, allocations[i].effectBlock, err);
        }
    }

    function assert_HasPendingIncrease(
        User operator,
        AllocateParams memory params,
        string memory err
    ) internal {
        uint32 delay = _getExistingAllocationDelay(operator);
        (, Allocation[] memory allocations) = _getAllocations(operator, params.operatorSet, params.strategies);

        for (uint i = 0; i < allocations.length; i++) {
            assertEq(allocations[i].effectBlock, uint32(block.number) + delay, err);
            assertTrue(allocations[i].currentMagnitude != params.newMagnitudes[i], err);
            assertGt(allocations[i].pendingDiff, 0, err);
        }
    }

    function assert_HasPendingDecrease(
        User operator,
        AllocateParams memory params,
        string memory err
    ) internal {
        uint32 deallocationDelay = allocationManager.DEALLOCATION_DELAY();
        (, Allocation[] memory allocations) = _getAllocations(operator, params.operatorSet, params.strategies);

        for (uint i = 0; i < allocations.length; i++) {
            assertEq(allocations[i].effectBlock, uint32(block.number) + deallocationDelay + 1, err);
            assertTrue(allocations[i].currentMagnitude != params.newMagnitudes[i], err);
            assertLt(allocations[i].pendingDiff, 0, err);
        }
    }

    function assert_IsRegistered(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal view {
        assertTrue(allocationManager.isMemberOfOperatorSet(address(operator), operatorSet), err);
    }

    function assert_IsSlashable(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal view {
        assertTrue(allocationManager.isOperatorSlashable(address(operator), operatorSet), err);
    }

    function assert_NotSlashable(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal view {
        assertFalse(allocationManager.isOperatorSlashable(address(operator), operatorSet), err);
    }

    function assert_IsAllocatedToSet(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal view {
        assertTrue(allocationManager.getAllocatedSets(address(operator)).contains(operatorSet), err);
    }

    function assert_IsNotAllocated(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal view {
        assertEq(allocationManager.getAllocatedStrategies(address(operator), operatorSet).length, 0, err);
    }

    function assert_IsAllocatedToSetStrats(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal view {
        IStrategy[] memory allocatedStrategies = allocationManager.getAllocatedStrategies(address(operator), operatorSet);

        for (uint i = 0; i < strategies.length; i++) {
            assertTrue(allocatedStrategies.contains(strategies[i]), err);
        }
    }

    function assert_HasAllocatedStake(
        User operator,
        AllocateParams memory params,
        string memory err
    ) internal {
        OperatorSet memory operatorSet = params.operatorSet;
        IStrategy[] memory strategies = params.strategies;
        uint64[] memory curMagnitudes = params.newMagnitudes;

        uint64[] memory maxMagnitudes = _getMaxMagnitudes(operator, params.strategies);
        (, uint[] memory operatorShares) = _getOperatorShares(operator, params.strategies);
        (, uint[] memory allocatedStake) = _getAllocatedStake(operator, operatorSet, strategies);

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

    function assert_HasSlashableStake(
        User operator,
        AllocateParams memory params,
        string memory err
    ) internal {
        OperatorSet memory operatorSet = params.operatorSet;
        IStrategy[] memory strategies = params.strategies;
        uint64[] memory curMagnitudes = params.newMagnitudes;

        uint64[] memory maxMagnitudes = _getMaxMagnitudes(operator, params.strategies);
        (, uint[] memory operatorShares) = _getOperatorShares(operator, params.strategies);
        (, uint[] memory slashableStake) = _getMinSlashableStake(operator, operatorSet, strategies);

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

    function assert_NoSlashableStake(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal {
        IStrategy[] memory strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);
        (, uint[] memory slashableStake) = _getMinSlashableStake(operator, operatorSet, strategies);

        for (uint i = 0; i < slashableStake.length; i++) {
            assertEq(slashableStake[i], 0, err);
        }
    }

    function assert_DSF_WAD(
        User staker,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (, uint[] memory depositScalingFactors) = _getDepositScalingFactors(staker, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(depositScalingFactors[i], WAD, err);
        }
    }

    function assert_BCSF_Zero(
        User staker,
        string memory err
    ) internal {
        (, uint64 curBCSF) = _getBeaconChainSlashingFactor(staker);
        assertEq(curBCSF, 0, err);
    }

    function assert_BCSF_WAD(
        User staker,
        string memory err
    ) internal {
        (, uint64 curBCSF) = _getBeaconChainSlashingFactor(staker);
        assertEq(curBCSF, WAD, err);
    }

    function assert_ActiveValidatorCount(
        User staker,
        uint expectedCount,
        string memory err
    ) internal {
        (, uint curActiveValidatorCount) = _getActiveValidatorCount(staker);
        assertEq(curActiveValidatorCount, expectedCount, err);
    }
    function assert_withdrawableSharesDecreasedByAtLeast(
        User staker,
        IStrategy[] memory strategies,
        uint256[] memory originalShares,
        uint256[] memory expectedDecreases,
        string memory err
    ) internal {
        for (uint i = 0; i < strategies.length; i++) {
            assert_withdrawableSharesDecreasedByAtLeast(staker, strategies[i], originalShares[i], expectedDecreases[i], err);
        }
    }

    function assert_withdrawableSharesDecreasedByAtLeast(
        User staker,
        IStrategy strategy,
        uint256 originalShares,
        uint256 expectedDecrease,
        string memory err
    ) internal {
        (, uint currentShares) = _getWithdrawableShares(staker, strategy);
        assertLt(currentShares, originalShares - expectedDecrease, err);
    }

    function assert_DepositShares_GTE_WithdrawableShares(
        User staker,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (, uint[] memory depositShares) = _getStakerDepositShares(staker, strategies);
        (, uint[] memory withdrawableShares) = _getWithdrawableShares(staker, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertGe(depositShares[i], withdrawableShares[i], err);
        }
    }
    
    /*******************************************************************************
                                SNAPSHOT ASSERTIONS
                       TIME TRAVELERS ONLY BEYOND THIS POINT
    *******************************************************************************/

    /*******************************************************************************
                         SNAPSHOT ASSERTIONS: ALLOCATIONS
    *******************************************************************************/

    function assert_Snap_Became_Registered(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal {
        (bool prevIsMemberOfSet, bool curIsMemberOfSet) = _getIsMemberOfSet(operator, operatorSet);
        assertFalse(prevIsMemberOfSet, err);
        assertTrue(curIsMemberOfSet, err);
    }

    function assert_Snap_Became_Deregistered(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal {
        (bool prevIsMemberOfSet, bool curIsMemberOfSet) = _getIsMemberOfSet(operator, operatorSet);
        assertTrue(prevIsMemberOfSet, err);
        assertFalse(curIsMemberOfSet, err);
    }

    function assert_Snap_Unchanged_Registration(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal {
        (bool prevIsMemberOfSet, bool curIsMemberOfSet) = _getIsMemberOfSet(operator, operatorSet);
        assertEq(prevIsMemberOfSet, curIsMemberOfSet, err);
    }
    
    function assert_Snap_Became_Slashable(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal {
        (bool prevIsSlashable, bool curIsSlashable) = _getIsSlashable(operator, operatorSet);
        assertFalse(prevIsSlashable, err);
        assertTrue(curIsSlashable, err);
    }

    function assert_Snap_Remains_Slashable(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal {
        (bool prevIsSlashable, bool curIsSlashable) = _getIsSlashable(operator, operatorSet);
        assertTrue(prevIsSlashable, err);
        assertTrue(curIsSlashable, err);
    }

    function assert_Snap_Unchanged_Slashability(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal {
        (bool prevIsSlashable, bool curIsSlashable) = _getIsSlashable(operator, operatorSet);
        assertEq(prevIsSlashable, curIsSlashable, err);
    }

    function assert_Snap_Unchanged_AllocatedStrats(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal {
        (IStrategy[] memory prev, IStrategy[] memory current) = _getAllocatedStrats(operator, operatorSet);
        assertEq(current.length, prev.length, err);
        for (uint i = 0; i < current.length; i++) {
            assertEq(address(current[i]), address(prev[i]), err);
        }
    }

    function assert_Snap_Removed_AllocatedStrats(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory removedStrats,
        string memory err
    ) internal {
        (IStrategy[] memory curAllocatedStrats, IStrategy[] memory prevAllocatedStrats) = _getAllocatedStrats(operator, operatorSet);
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
        (Allocation[] memory prev, Allocation[] memory current) = _getAllocations(operator, operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            Allocation memory curAllocation = current[i];
            Allocation memory prevAllocation = prev[i];
            assertEq(curAllocation.currentMagnitude, prevAllocation.currentMagnitude, err);
            assertEq(curAllocation.pendingDiff, prevAllocation.pendingDiff, err);
            assertEq(curAllocation.effectBlock, prevAllocation.effectBlock, err);
        }
    }

    function assert_Snap_Added_AllocatedSet(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal {
        (OperatorSet[] memory prev, OperatorSet[] memory current) = _getAllocatedSets(operator);
        assertEq(current.length, prev.length + 1, err);
        assertFalse(prev.contains(operatorSet), err);
        assertTrue(current.contains(operatorSet), err);
    }

    function assert_Snap_Unchanged_AllocatedSets(
        User operator,
        string memory err
    ) internal {
        (OperatorSet[] memory prev, OperatorSet[] memory current) = _getAllocatedSets(operator);
        assertEq(current.length, prev.length, err);
        for (uint i = 0; i < current.length; i++) {
            assertEq(current[i].avs, prev[i].avs, err);
            assertEq(current[i].id, prev[i].id, err);
        }
    }

    function assert_Snap_Removed_AllocatedSet(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal {
        (OperatorSet[] memory prev, OperatorSet[] memory current) = _getAllocatedSets(operator);
        assertEq(current.length + 1, prev.length, err);
        assertTrue(prev.contains(operatorSet), err);
        assertFalse(current.contains(operatorSet), err);
    }

    function assert_Snap_Added_RegisteredSet(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal {
        (OperatorSet[] memory prev, OperatorSet[] memory current) = _getRegisteredSets(operator);
        assertEq(current.length, prev.length + 1, err);
        assertFalse(prev.contains(operatorSet), err);
        assertTrue(current.contains(operatorSet), err);
    }

    function assert_Snap_Removed_RegisteredSet(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal {
        (OperatorSet[] memory prev, OperatorSet[] memory current) = _getRegisteredSets(operator);
        assertEq(current.length + 1, prev.length, err);
        assertTrue(prev.contains(operatorSet), err);
        assertFalse(current.contains(operatorSet), err);
    }

    function assert_Snap_Unchanged_RegisteredSet(
        User operator,
        string memory err
    ) internal {
        (OperatorSet[] memory prev, OperatorSet[] memory current) = _getRegisteredSets(operator);
        assertEq(current.length, prev.length, err);
        for (uint i = 0; i < current.length; i++) {
            assertEq(current[i].avs, prev[i].avs, err);
            assertEq(current[i].id, prev[i].id, err);
        }
    }

    function assert_Snap_Added_MemberOfSet(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal {
        (address[] memory prev, address[] memory current) = _getMembers(operatorSet);
        assertEq(current.length, prev.length + 1, err);
        assertFalse(prev.contains(address(operator)), err);
        assertTrue(current.contains(address(operator)), err);
    }

    function assert_Snap_Removed_MemberOfSet(
        User operator,
        OperatorSet memory operatorSet,
        string memory err
    ) internal {
        (address[] memory prev, address[] memory current) = _getMembers(operatorSet);
        assertEq(current.length + 1, prev.length, err);
        assertTrue(prev.contains(address(operator)), err);
        assertFalse(current.contains(address(operator)), err);
    }

    function assert_Snap_Unchanged_MemberOfSet(
        OperatorSet memory operatorSet,
        string memory err
    ) internal {
        (address[] memory prev, address[] memory current) = _getMembers(operatorSet);
        assertEq(current.length, prev.length, err);
        for (uint i = 0; i < current.length; i++) {
            assertEq(current[i], prev[i], err);
        }
    }

    function assert_Snap_StakeBecameSlashable(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (uint[] memory prev, uint[] memory current) = _getMinSlashableStake(operator, operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertTrue(prev[i] < current[i], err);
        }
    }

    function assert_Snap_StakeBecomeUnslashable(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (uint[] memory prev, uint[] memory current) = _getMinSlashableStake(operator, operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertTrue(prev[i] > current[i], err);
        }
    }

    function assert_Snap_StakeBecomeUnslashable(
        address operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (uint[] memory curSlashableStake, uint[] memory prevSlashableStake) = _getMinSlashableStake(User(payable(operator)), operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertTrue(prevSlashableStake[i] > curSlashableStake[i], err);
        }
    }

    function assert_Snap_Added_SlashableStake(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        uint[] memory slashableShares,
        string memory err
    ) internal {
        (uint[] memory prev, uint[] memory current) = _getMinSlashableStake(operator, operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(current[i], prev[i] + slashableShares[i], err);
        }
    }

    function assert_Snap_Unchanged_SlashableStake(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (uint[] memory prev, uint[] memory current) = _getMinSlashableStake(operator, operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(current[i], prev[i], err);
        }
    }

    function assert_Snap_Removed_SlashableStake(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        uint[] memory removedSlashableShares,
        string memory err
    ) internal {
        (uint[] memory prev, uint[] memory current) = _getMinSlashableStake(operator, operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(current[i] + removedSlashableShares[i], prev[i], err);
        }
    }

    function assert_Snap_Slashed_SlashableStake(
        User operator,
        OperatorSet memory operatorSet,
        SlashingParams memory params,
        string memory err
    ) internal {
        (uint[] memory prevSlashableStake, uint[] memory currentSlashableStake)
             = _getMinSlashableStake(operator, operatorSet, params.strategies);
        (Magnitudes[] memory prevMagnitudes, Magnitudes[] memory currentMagnitudes)
             = _getMagnitudes(operator, params.strategies);

        for (uint i = 0; i < params.strategies.length; i++) {
            uint expectedSlashed = SlashingLib.calcSlashedAmount({
                operatorShares: prevSlashableStake[i],
                prevMaxMagnitude: prevMagnitudes[i].max,
                newMaxMagnitude: currentMagnitudes[i].max
            });

            assertEq(currentSlashableStake[i], prevSlashableStake[i] - expectedSlashed, err);
        }
    }

    //@dev requires slashparams strategies to be same as withdrawal strategies
    // meant to be used in check_base_slashing_state
    function assert_Snap_Decreased_SlashableSharesInQueue(
        User operator,
        SlashingParams memory slashParams,
        Withdrawal[] memory withdrawals,
        string memory err
    ) internal {
        IStrategy[] memory strategies = slashParams.strategies;
        (uint[] memory curSlashableSharesInQueue, uint[] memory prevSlashableSharesInQueue) = _getSlashableSharesInQueue(operator, strategies);
        uint[] memory totalScaledShares = new uint[](strategies.length);
        for (uint i = 0; i < withdrawals.length; i++) {
            for (uint j = 0; j < withdrawals[i].strategies.length; j++) {
                totalScaledShares[j] = totalScaledShares[j] + withdrawals[i].scaledShares[j];
            }
        }

        for (uint i = 0; i < strategies.length; i++) {
            assertEq(curSlashableSharesInQueue[i], prevSlashableSharesInQueue[i] - totalScaledShares[i].mulWad(slashParams.wadsToSlash[i]), err);
        }
    }

    function assert_Snap_Increased_SlashableSharesInQueue(
        User operator,
        Withdrawal[] memory withdrawals,
        string memory err
    ) internal {
        uint[] memory curSlashableSharesInQueue;
        uint[] memory prevSlashableSharesInQueue;
        uint64[] memory maxMagnitudes;
        for (uint i = 0; i < withdrawals.length; i++) {
            (curSlashableSharesInQueue, prevSlashableSharesInQueue) = _getSlashableSharesInQueue(operator, withdrawals[i].strategies);
            maxMagnitudes = _getMaxMagnitudes(operator, withdrawals[i].strategies);

            for (uint j = 0; j < withdrawals[i].strategies.length; j++) {
                assertEq(curSlashableSharesInQueue[j], prevSlashableSharesInQueue[j] + withdrawals[i].scaledShares[j].mulWad(maxMagnitudes[j]), err);
            }
        }
    }

    function assert_Snap_StakeBecameAllocated(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (uint[] memory prev, uint[] memory current) = _getAllocatedStake(operator, operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertGt(current[i], prev[i], err);
        }
    }

    function assert_Snap_StakeBecameDeallocated(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (uint[] memory prev, uint[] memory current) = _getAllocatedStake(operator, operatorSet, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertLt(current[i], prev[i], err);
        }
    }

    function assert_Snap_Unchanged_AllocatedStake(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (uint[] memory prev, uint[] memory current) = _getAllocatedStake(operator, operatorSet, strategies);
        for (uint i = 0; i < current.length; i++) {
            assertEq(current[i], prev[i], err);
        }
    }

    function assert_Snap_Slashed_AllocatedStake(
        User operator,
        OperatorSet memory operatorSet,
        SlashingParams memory params,
        string memory err
    ) internal {
        (uint[] memory prev, uint[] memory current) = _getAllocatedStake(operator, operatorSet, params.strategies);
        (Magnitudes[] memory prevMagnitudes, Magnitudes[] memory currentMagnitudes) = _getMagnitudes(operator, params.strategies);
        for (uint i = 0; i < current.length; i++) {
            uint expectedSlashed = SlashingLib.calcSlashedAmount({
                operatorShares: prev[i],
                prevMaxMagnitude: prevMagnitudes[i].max,
                newMaxMagnitude: currentMagnitudes[i].max
            });
            assertEq(current[i], prev[i] - expectedSlashed, err);
        }
    }

    function assert_Snap_Added_EncumberedMagnitude(
        User operator,
        IStrategy[] memory strategies,
        uint64[] memory magnitudeAdded,
        string memory err
    ) internal {
        (Magnitudes[] memory prev, Magnitudes[] memory current) = _getMagnitudes(operator, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(current[i].encumbered, prev[i].encumbered + magnitudeAdded[i], err);
        }
    }

    function assert_Snap_Unchanged_EncumberedMagnitude(
        User operator,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (Magnitudes[] memory prev, Magnitudes[] memory current) = _getMagnitudes(operator, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(current[i].encumbered, prev[i].encumbered, err);
        }
    }

    function assert_Snap_Removed_EncumberedMagnitude(
        User operator,
        IStrategy[] memory strategies,
        uint64[] memory magnitudeRemoved,
        string memory err
    ) internal {
        (Magnitudes[] memory prev, Magnitudes[] memory current) = _getMagnitudes(operator, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(current[i].encumbered + magnitudeRemoved[i], prev[i].encumbered, err);
        }
    }

    function assert_Snap_Slashed_EncumberedMagnitude(
        User operator,
        SlashingParams memory params,
        string memory err
    ) internal {
        (Magnitudes[] memory prev, Magnitudes[] memory current) = _getMagnitudes(operator, params.strategies);
        for (uint i = 0; i < params.strategies.length; i++) {
            uint expectedSlashed = prev[i].encumbered.mulWadRoundUp(params.wadsToSlash[i]);
            assertEq(current[i].encumbered, prev[i].encumbered - expectedSlashed, err);
        }
    }

    function assert_Snap_Added_AllocatableMagnitude(
        User operator,
        IStrategy[] memory strategies,
        uint64[] memory magnitudeFreed,
        string memory err
    ) internal {
        (Magnitudes[] memory prev, Magnitudes[] memory current) = _getMagnitudes(operator, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(current[i].allocatable, prev[i].allocatable + magnitudeFreed[i], err);
        }
    }

    function assert_Snap_Unchanged_AllocatableMagnitude(
        User operator,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (Magnitudes[] memory prev, Magnitudes[] memory current) = _getMagnitudes(operator, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(current[i].allocatable, prev[i].allocatable, err);
        }
    }

    function assert_Snap_Removed_AllocatableMagnitude(
        User operator,
        IStrategy[] memory strategies,
        uint64[] memory magnitudeRemoved,
        string memory err
    ) internal {
        (Magnitudes[] memory prev, Magnitudes[] memory current) = _getMagnitudes(operator, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(current[i].allocatable, prev[i].allocatable - magnitudeRemoved[i], err);
        }
    }

    function assert_Snap_Allocated_Magnitude(
        User operator,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (Magnitudes[] memory prev, Magnitudes[] memory current) = _getMagnitudes(operator, strategies);
        /// Check:
        /// allocatable increased
        /// encumbered decreased
        for (uint i = 0; i < strategies.length; i++) {
            assertLt(current[i].allocatable, prev[i].allocatable, err);
            assertGt(current[i].encumbered, prev[i].encumbered, err);
        }
    }

    function assert_Snap_Deallocated_Magnitude(
        User operator,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (Magnitudes[] memory prev, Magnitudes[] memory current) = _getMagnitudes(operator, strategies);

        /// Check:
        /// allocatable increased
        /// encumbered decreased
        for (uint i = 0; i < strategies.length; i++) {
            assertGt(current[i].allocatable, prev[i].allocatable, err);
            assertLt(current[i].encumbered, prev[i].encumbered, err);
        }
    }

    function assert_Snap_Set_CurrentMagnitude(
        User operator,
        AllocateParams memory params,
        string memory err
    ) internal {
        (Allocation[] memory prev, Allocation[] memory current) = _getAllocations(operator, params.operatorSet, params.strategies);
        /// Prev allocation.currentMagnitude should NOT equal newly-set magnitude
        /// Cur allocation.currentMagnitude SHOULD
        for (uint i = 0; i < params.strategies.length; i++) {
            assertTrue(prev[i].currentMagnitude != params.newMagnitudes[i], err);
            assertEq(current[i].currentMagnitude, params.newMagnitudes[i], err);
        }
    }

    function assert_Snap_Slashed_Allocation(
        User operator,
        OperatorSet memory operatorSet,
        SlashingParams memory params,
        string memory err
    ) internal {
        (Allocation[] memory prev, Allocation[] memory current) = _getAllocations(operator, operatorSet, params.strategies);
        for (uint i = 0; i < params.strategies.length; i++) {
            uint expectedSlashed = prev[i].currentMagnitude.mulWadRoundUp(params.wadsToSlash[i]);
            assertEq(current[i].currentMagnitude, prev[i].currentMagnitude - expectedSlashed, err);
        }
    }

    function assert_Snap_Unchanged_MaxMagnitude(
        User operator,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (Magnitudes[] memory prev, Magnitudes[] memory current) = _getMagnitudes(operator, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(current[i].max, prev[i].max, err);
        }
    }

    function assert_Snap_Slashed_MaxMagnitude(
        User operator,
        OperatorSet memory operatorSet,
        SlashingParams memory params,
        string memory err
    ) internal {
        (Magnitudes[] memory prev, Magnitudes[] memory current) = _getMagnitudes(operator, params.strategies);
        for (uint i = 0; i < params.strategies.length; i++) {
            uint expectedSlashed = prev[i].max.mulWadRoundUp(params.wadsToSlash[i]);
            assertEq(current[i].max, prev[i].max - expectedSlashed, err);
        }
    }

    function assert_Snap_Allocations_Slashed(
        SlashingParams memory slashingParams,
        OperatorSet memory operatorSet,
        bool completed,
        string memory err
    ) internal {
        User op = User(payable(slashingParams.operator));
        
        (Allocation[] memory prevAllocs, Allocation[] memory currentAllocs) = _getAllocations(op, operatorSet, slashingParams.strategies);
        (Magnitudes[] memory prevMagnitudes, Magnitudes[] memory currentMagnitudes) = _getMagnitudes(op, slashingParams.strategies);

        uint32 delay = _getExistingAllocationDelay(User(payable(slashingParams.operator)));

        for (uint i = 0; i < slashingParams.strategies.length; i++) {
            Allocation memory curAlloc = currentAllocs[i];
            Allocation memory prevAlloc = prevAllocs[i]; 

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
            Magnitudes memory curMagnitude = currentMagnitudes[i];
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

    function assert_Snap_StakerWithdrawableShares_AfterSlash(
        User staker,
        AllocateParams memory allocateParams,
        SlashingParams memory slashingParams,
        string memory err
    ) internal {
        (uint[] memory curShares, uint[] memory prevShares) = _getWithdrawableShares(staker, allocateParams.strategies);

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

    /*******************************************************************************
                    SNAPSHOT ASSERTIONS: BEACON CHAIN AND AVS SLASHING
    *******************************************************************************/
    
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
        require(slashingParams.strategies[0] == BEACONCHAIN_ETH_STRAT, "only beacon strategy supported");

        (uint[] memory curShares, uint[] memory prevShares) = _getWithdrawableShares(staker, allocateParams.strategies);

        uint256 slashedShares = 0;

        uint wadToSlash = slashingParams.wadsToSlash[0];
        slashedShares = prevShares[0].mulWadRoundUp(allocateParams.newMagnitudes[0].mulWadRoundUp(wadToSlash));

        assertEq(prevShares[0] - slashedShares, curShares[0], err);
    }

    /// @dev Validates behavior of "restaking", ie. that the funds can be slashed twice
    function assert_Snap_StakerWithdrawableShares_AfterAVSSlash_BCSlash(
        User staker,
        AllocateParams memory allocateParams,
        SlashingParams memory slashingParams,
        string memory err
    ) internal {
        require(allocateParams.strategies.length == 1  && slashingParams.strategies.length == 1, "only beacon strategy supported");
        require(allocateParams.strategies[0] == BEACONCHAIN_ETH_STRAT, "only beacon strategy supported");
        require(slashingParams.strategies[0] == BEACONCHAIN_ETH_STRAT, "only beacon strategy supported");

        (uint[] memory curShares, uint[] memory prevShares) = _getWithdrawableShares(staker, allocateParams.strategies);
        (, uint[] memory depositShares) = _getStakerDepositShares(staker, allocateParams.strategies);

        // 1. The withdrawable shares should decrease by a factor of the BCSF
        (, uint64 curBCSF) = _getBeaconChainSlashingFactor(staker);
        assertEq(prevShares[0].mulWad(curBCSF), curShares[0], err);

        /**
         * 2. The delta in shares is given by:
         * (depositShares * operatorMag) - (depositShares * operatorMag * BCSF)
         *  = depositShares * operatorMag * (1 - BCSF)
         */
        (, uint64 beaconChainSlashingFactor) = _getBeaconChainSlashingFactor(staker);
        uint wadToSlash = slashingParams.wadsToSlash[0];
        uint originalAVSSlashedShares = depositShares[0].mulWadRoundUp(allocateParams.newMagnitudes[0].mulWadRoundUp(wadToSlash));
        uint withdrawableSharesAfterAVSSlash = depositShares[0] - originalAVSSlashedShares;
        uint expectedDelta = withdrawableSharesAfterAVSSlash.mulWad(WAD - beaconChainSlashingFactor);
        assertEq(prevShares[0] - expectedDelta, curShares[0], err);

        /**
         * 3. The attributable avs slashed shares should decrease by a factor of the BCSF
         * Attributable avs slashed shares = originalWithdrawableShares - bcSlashedShares - curShares
         * Where bcSlashedShares = originalWithdrawableShares * (1 - BCSF)
         */
        uint bcSlashedShares = depositShares[0].mulWad(WAD - beaconChainSlashingFactor);
        uint attributableAVSSlashedShares = depositShares[0] - bcSlashedShares - curShares[0];
        assertEq(originalAVSSlashedShares.mulWad(beaconChainSlashingFactor), attributableAVSSlashedShares, err);
    }

    /**
     * @dev Validates behavior of "restaking", ie. that the funds can be slashed twice. Also validates
     *      the edge case where a validator is proven prior to the BC slash.
     * @dev These bounds are based off of rounding when avs and bc slashing occur together
     */
    function assert_Snap_StakerWithdrawableShares_AVSSlash_ValidatorProven_BCSlash(
        User staker,
        uint256 originalWithdrawableShares,
        uint256 extraValidatorShares,
        AllocateParams memory allocateParams,
        SlashingParams memory slashingParams,
        string memory err
    ) internal {
        require(allocateParams.strategies.length == 1  && slashingParams.strategies.length == 1, "only beacon strategy supported");
        require(allocateParams.strategies[0] == BEACONCHAIN_ETH_STRAT, "only beacon strategy supported");
        require(slashingParams.strategies[0] == BEACONCHAIN_ETH_STRAT, "only beacon strategy supported");

        (uint[] memory curShares, uint[] memory prevShares) = _getWithdrawableShares(staker, allocateParams.strategies);

        // 1. The withdrawable shares should decrease by a factor of the BCSF
        (, uint64 curBCSF) = _getBeaconChainSlashingFactor(staker);
        assertApproxEqAbs(prevShares[0].mulWad(curBCSF), curShares[0], 1e5, err);

        /**
         * 2. The delta in shares is given by:
         * (originalWithdrawableShares * operatorMag) + extraValidatorShares - (depositShares * operatorMag * BCSF * dsf)
         */
        (, uint beaconChainSlashingFactor) = _getBeaconChainSlashingFactor(staker);
        uint wadToSlash = slashingParams.wadsToSlash[0];
        uint originalAVSSlashedShares = originalWithdrawableShares.mulWadRoundUp(allocateParams.newMagnitudes[0].mulWadRoundUp(wadToSlash));
        uint withdrawableSharesAfterValidatorProven = originalWithdrawableShares - originalAVSSlashedShares + extraValidatorShares;
        uint expectedDelta = withdrawableSharesAfterValidatorProven.mulWad(WAD - beaconChainSlashingFactor);
        assertApproxEqAbs(prevShares[0] - expectedDelta, curShares[0], 1e5, err);

        /**
         * 3. The attributable avs slashed shares should decrease by a factor of the BCSF
         * Attributable avs slashed shares = depositShares - bcSlashedShares - curShars
         * Where bcSlashedShares = depositShares * (1 - BCSF)
         */
        (, uint[] memory depositShares) = _getStakerDepositShares(staker, allocateParams.strategies);
        uint bcSlashedShares = depositShares[0].mulWad(WAD - beaconChainSlashingFactor);
        uint attributableAVSSlashedShares = depositShares[0] - bcSlashedShares - curShares[0];
        assertApproxEqAbs(originalAVSSlashedShares.mulWad(beaconChainSlashingFactor), attributableAVSSlashedShares, 1e5, err);
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
        (uint[] memory curShares, uint[] memory prevShares) = _getOperatorShares(operator, strategies);

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
        (uint[] memory curShares, uint[] memory prevShares) = _getOperatorShares(operator, strategies);

        // For each strategy, check (prev - removed == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i], curShares[i] + removedShares[i], err);
        }
    }

    /// @dev Check that the operator's shares in ALL strategies have not changed
    /// since the last snapshot
    function assert_Snap_Unchanged_OperatorShares(
        User operator,
        string memory err
    ) internal {
        IStrategy[] memory strategies = allStrats;

        (uint[] memory curShares, uint[] memory prevShares) = _getOperatorShares(operator, strategies);

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
        (uint[] memory curShares, uint[] memory prevShares) = _getOperatorShares(operator, strategies);

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

    function assert_Snap_Slashed_OperatorShares(
        User operator,
        SlashingParams memory params,
        string memory err
    ) internal {
        (uint[] memory curShares, uint[] memory prevShares) = _getOperatorShares(operator, params.strategies);
        (Magnitudes[] memory prevMagnitudes, Magnitudes[] memory currentMagnitudes) = _getMagnitudes(operator, params.strategies);

        for (uint i = 0; i < params.strategies.length; i++) {
            // Slashing doesn't occur if the operator has no slashable magnitude
            // This prevents a div by 0 when calculating expected slashed
            uint expectedSlashed = prevMagnitudes[i].max == 0 ? 0 : SlashingLib.calcSlashedAmount({
                operatorShares: prevShares[i],
                prevMaxMagnitude: prevMagnitudes[i].max,
                newMaxMagnitude: currentMagnitudes[i].max
            });

            assertEq(curShares[i], prevShares[i] - expectedSlashed, err);
        }
    }

    function assert_Snap_Increased_BurnableShares(
        User operator,
        SlashingParams memory params,
        string memory err
    ) internal {
        (uint[] memory curBurnable, uint[] memory prevBurnable) = _getBurnableShares(params.strategies);
        (uint[] memory curShares, uint[] memory prevShares) = _getOperatorShares(operator, params.strategies);
        for (uint i = 0; i < params.strategies.length; i++) {
            uint slashedAtLeast = prevShares[i] - curShares[i];
            // Not factoring in slashable shares in queue here, because that gets more complex (TODO)
            assertTrue(curBurnable[i] >= (prevBurnable[i] + slashedAtLeast), err);
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
        (uint[] memory curShares, uint[] memory prevShares) = _getStakerDepositShares(staker, strategies);

        // For each strategy, check (prev + added == cur)
        for (uint i = 0; i < strategies.length; i++) {        
            assertApproxEqAbs(prevShares[i] + addedShares[i], curShares[i], 1, err);     
        }
    }

    function assert_Snap_Added_Staker_DepositShares(
        User staker, 
        IStrategy strat, 
        uint addedShares,
        string memory err
    ) internal {
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
        (uint[] memory curShares, uint[] memory prevShares) = _getStakerDepositShares(staker, strategies);

        // For each strategy, check (prev - removed == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] - removedShares[i], curShares[i], err);
        }
    }

    function assert_Snap_Removed_Staker_DepositShares(
        User staker, 
        IStrategy strat, 
        uint removedShares,
        string memory err
    ) internal {
        assert_Snap_Removed_Staker_DepositShares(staker, strat.toArray(), removedShares.toArrayU256(), err);
    }

    /// @dev Check that the staker's delegatable shares in ALL strategies have not changed
    /// since the last snapshot
    function assert_Snap_Unchanged_Staker_DepositShares(
        User staker,
        string memory err
    ) internal {
        IStrategy[] memory strategies = allStrats;

        (uint[] memory curShares, uint[] memory prevShares) = _getStakerDepositShares(staker, strategies);

        // For each strategy, check (prev == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i], curShares[i], err);
        }
    }

    /// @dev Check that the staker's withdrawable shares have increased by `addedShares`
    function assert_Snap_Added_Staker_WithdrawableShares(
        User staker, 
        IStrategy[] memory strategies, 
        uint[] memory addedShares,
        string memory err
    ) internal {
        (uint[] memory curShares, uint[] memory prevShares) = _getStakerWithdrawableShares(staker, strategies);

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
        (uint[] memory curShares, uint[] memory prevShares) = _getStakerWithdrawableShares(staker, strategies);

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
        (uint[] memory curShares, uint[] memory prevShares) = _getStakerWithdrawableShares(staker, strategies);

        // For each strategy, check (prev - removed == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] - removedShares[i], curShares[i], err);
        }
    }

    function assert_Snap_Removed_Staker_WithdrawableShares(
        User staker, 
        IStrategy strat, 
        uint removedShares,
        string memory err
    ) internal {
        assert_Snap_Removed_Staker_WithdrawableShares(staker, strat.toArray(), removedShares.toArrayU256(), err);
    }

    function assert_Snap_Unchanged_Staker_WithdrawableShares(
        User staker,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (uint[] memory curShares, uint[] memory prevShares) = _getStakerWithdrawableShares(staker, strategies);
        // For each strategy, check all shares have been withdrawn
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i], curShares[i], err);
        }
    }

    /// @dev Check that the staker's withdrawable shares have changed by the expected amount
    function assert_Snap_Expected_Staker_WithdrawableShares_Delegation(
        User staker,
        User operator,
        IStrategy[] memory strategies,
        uint[] memory depositShares,
        string memory err
    ) internal {
        (, uint[] memory curShares) = _getStakerWithdrawableShares(staker, strategies);
        // Use timewarp to get previous staker shares
        uint[] memory expectedShares = _getExpectedWithdrawableSharesDelegate(staker, operator, strategies, depositShares);

        // For each strategy, check expected == current
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(expectedShares[i], curShares[i], err);
        }
    }

    function assert_Snap_Expected_Staker_WithdrawableShares_Deposit(
        User staker,
        User operator,
        IStrategy[] memory strategies,
        uint[] memory depositSharesAdded,
        string memory err
    ) internal {
        (uint[] memory curShares, uint[] memory prevShares) = _getStakerWithdrawableShares(staker, strategies);
        uint[] memory expectedWithdrawableShares = new uint[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            if (prevShares[i] == 0 && depositSharesAdded[i] > 0){
                expectedWithdrawableShares[i] = _getExpectedWithdrawableSharesDeposit(staker, operator, strategies[i], depositSharesAdded[i]);
                assertEq(curShares[i], expectedWithdrawableShares[i], err);
            } else {
                (uint[] memory prevDepositShares, ) = _getStakerDepositShares(staker, strategies);
                (, uint[] memory currentDSF) = _getDepositScalingFactors(staker, strategies[i].toArray());
                assertEq((prevDepositShares[i] + depositSharesAdded[i]).mulWad(currentDSF[0]).mulWad(_getSlashingFactor(staker, strategies[i])), curShares[i], err);
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
        (uint[] memory curShares, uint[] memory prevShares) = _getStakerWithdrawableShares(staker, strategies);

        // For each strategy, check diff between (prev-removed) and curr is at most 1 gwei
        for (uint i = 0; i < strategies.length; i++) {
            assertApproxEqAbs(prevShares[i] - removedShares[i], curShares[i], 1e9, err);
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
        (uint[] memory curShares, uint[] memory prevShares) = _getStakerWithdrawableShares(staker, strategies);

        // For each strategy, check diff between (prev-removed) and curr is at most 1 gwei
        for (uint i = 0; i < strategies.length; i++) {
            assertApproxEqAbs(prevShares[i] - removedShares[i], curShares[i], errBound, err);
        }
    }

    function assert_Snap_Removed_Staker_WithdrawableShares_AtLeast(
        User staker,
        IStrategy strat,
        uint removedShares,
        string memory err
    ) internal {
        assert_Snap_Removed_Staker_WithdrawableShares_AtLeast(staker, strat.toArray(), removedShares.toArrayU256(), err);
    }

    function assert_Snap_Removed_Staker_WithdrawableShares_AtLeast(
        User staker,
        IStrategy strat,
        uint removedShares,
        uint errBound,
        string memory err
    ) internal {
        assert_Snap_Removed_Staker_WithdrawableShares_AtLeast(staker, strat.toArray(), removedShares.toArrayU256(), errBound, err);
    }

    function assert_Snap_Delta_StakerShares(
        User staker, 
        IStrategy[] memory strategies, 
        int[] memory shareDeltas,
        string memory err
    ) internal {
        (int[] memory curShares, int[] memory prevShares) = _getStakerDepositSharesInt(staker, strategies);

        // For each strategy, check (prev + added == cur)
        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevShares[i] + shareDeltas[i], curShares[i], err);
        }
    }

    function assert_Snap_Unchanged_DSF(
        User staker,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (uint[] memory curDSFs, uint[] memory prevDSFs) = _getDepositScalingFactors(staker, strategies);

        for (uint i = 0; i < strategies.length; i++) {
            assertEq(prevDSFs[i], curDSFs[i], err);
        }
    }

    function assert_Snap_Increased_DSF(
        User staker,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (uint[] memory curDSFs, uint[] memory prevDSFs) = _getDepositScalingFactors(staker, strategies);

        for (uint i = 0; i < strategies.length; i++) {
            assertGt(curDSFs[i], prevDSFs[i], err);
        }
    }

    /// @dev Used to assert that the DSF is either increased or unchanged, depending on the slashing factor, on a deposit
    function assert_Snap_DSF_State_Deposit(
        User staker,
        IStrategy[] memory strategies,
        string memory err
    ) internal {        
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy[] memory stratArray = strategies[i].toArray();
            /// @dev We don't need the previous slashing factors as they shouldn't change before/after a deposit
            uint curSlashingFactor = _getSlashingFactor(staker, strategies[i]);

            // If there was never a slashing, no need to normalize
            if (curSlashingFactor == WAD) {
                assert_Snap_Unchanged_DSF(staker, stratArray, err); // No slashing, so DSF is unchanged
                assert_DSF_WAD(staker, stratArray, err); // DSF should have always been WAD
            }
            // If there was a slashing, and we deposit, normalize
            else {
                assert_Snap_Increased_DSF(staker, stratArray, err); // Slashing, so DSF is increased
            }
        }
    }

    /// @dev When completing withdrawals as shares, we must also handle the case where a staker completes a withdrawal for 0 shares
    function assert_Snap_DSF_State_WithdrawalAsShares(
        User staker,
        IStrategy[] memory strategies,
        string memory err
    ) internal {
        (, uint[] memory curDepositShares) = _getStakerDepositShares(staker, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy[] memory stratArray = strategies[i].toArray();
            /// We don't need the previous slashing factors as they shouldn't change before/after a deposit
            uint curSlashingFactor = _getSlashingFactor(staker, strategies[i]);
            
            // If there was never a slashing, no need to normalize
            // If there was a slashing, but we complete a withdrawal for 0 shares, no need to normalize
            if (curSlashingFactor == WAD || curDepositShares[i] == 0) {
                assert_Snap_Unchanged_DSF(staker, stratArray, err);
                assert_DSF_WAD(staker, stratArray, err);
            }
            // If there was a slashing, and we complete a withdrawal for non-zero shares, normalize
            else {
                assert_Snap_Increased_DSF(staker, stratArray, err);
            }
        }
    }

    /// @dev On a delegation, the DSF should be increased if the operator magnitude is non-WAD
    function assert_Snap_DSF_State_Delegation(
        User staker,
        User operator,
        IStrategy[] memory strategies,
        uint[] memory delegatableShares,
        string memory err
    ) internal {
        uint64[] memory maxMags = _getMaxMagnitudes(operator, strategies);

        for (uint i = 0; i < strategies.length; i++) {
            IStrategy[] memory stratArray = strategies[i].toArray();

            // If you are delegating with 0 shares, no need to normalize
            // If there was never an operator slashing, no need to normalize
            if (delegatableShares[i] == 0 || maxMags[i] == WAD) {
                assert_Snap_Unchanged_DSF(staker, stratArray, err);

                // If we are not a beaconChainStrategy, we should also have a DSF of WAD 
                // We exclude BEACONCHAIN_ETH_STRAT because it could have had a non-WAD DSF from BC slashings
                if (strategies[i] != BEACONCHAIN_ETH_STRAT) {
                    assert_DSF_WAD(staker, stratArray, err);
                }
            }
            // If there was an operator slashing, and delegating with non-zero shares, normalize
            else { 
                assert_Snap_Increased_DSF(staker, stratArray, err); // Slashing, so DSF is increased
            }
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
        (uint[] memory curShares, uint[] memory prevShares) = _getTotalStrategyShares(strategies);

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
        (uint[] memory curShares, uint[] memory prevShares) = _getTotalStrategyShares(strategies);

        for (uint i = 0; i < strategies.length; i++) {
            uint prevShare = prevShares[i];
            uint curShare = curShares[i];

            assertEq(prevShare, curShare, err);
        }
    }

    function assert_SlashableStake_Decrease_BCSlash(
        User staker
    ) internal {
        if (delegationManager.isDelegated(address(staker))) {
            address operator = delegationManager.delegatedTo(address(staker));
            (OperatorSet[] memory operatorSets, Allocation[] memory allocations) = _getStrategyAllocations(operator, BEACONCHAIN_ETH_STRAT);
            for (uint i = 0; i < operatorSets.length; i++) {
                if (allocations[i].currentMagnitude > 0) {
                    assert_Snap_StakeBecomeUnslashable(operator, operatorSets[i], BEACONCHAIN_ETH_STRAT.toArray(), "operator should have minSlashableStake decreased");
                }
            }
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
        (uint[] memory curTokenBalances, uint[] memory prevTokenBalances) = _getTokenBalances(staker, tokens);

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
        (uint[] memory curTokenBalances, uint[] memory prevTokenBalances) = _getTokenBalances(staker, tokens);

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

        (uint[] memory curTokenBalances, uint[] memory prevTokenBalances) = _getTokenBalances(staker, tokens);

        for (uint i = 0; i < tokens.length; i++) {
            assertEq(prevTokenBalances[i], curTokenBalances[i], err);
        }
    }

    /*******************************************************************************
                      SNAPSHOT ASSERTIONS: QUEUED WITHDRAWALS
    *******************************************************************************/

    function assert_Snap_Added_QueuedWithdrawals(
        User staker, 
        Withdrawal[] memory withdrawals,
        string memory err
    ) internal {
        (uint curQueuedWithdrawals, uint prevQueuedWithdrawals) = _getCumulativeWithdrawals(staker);
        assertEq(prevQueuedWithdrawals + withdrawals.length, curQueuedWithdrawals, err);
    }

    function assert_Snap_Added_QueuedWithdrawal(
        User staker, 
        Withdrawal memory /*withdrawal*/,
        string memory err
    ) internal {
        (uint curQueuedWithdrawal, uint prevQueuedWithdrawal) = _getCumulativeWithdrawals(staker);
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
        (uint curActiveValidatorCount, uint prevActiveValidatorCount) = _getActiveValidatorCount(staker);
        assertEq(prevActiveValidatorCount + addedValidators, curActiveValidatorCount, err);
    }

    function assert_Snap_Removed_ActiveValidatorCount(
        User staker,
        uint exitedValidators,
        string memory err
    ) internal {
        (uint curActiveValidatorCount, uint prevActiveValidatorCount) = _getActiveValidatorCount(staker);
        assertEq(curActiveValidatorCount + exitedValidators, prevActiveValidatorCount, err);
    }

    function assert_Snap_Unchanged_ActiveValidatorCount(
        User staker,
        string memory err
    ) internal {
        (uint curActiveValidatorCount, uint prevActiveValidatorCount) = _getActiveValidatorCount(staker);
        assertEq(curActiveValidatorCount, prevActiveValidatorCount, err);
    }

    function assert_Snap_Added_ActiveValidators(
        User staker,
        uint40[] memory addedValidators,
        string memory err
    ) internal {
        bytes32[] memory pubkeyHashes = beaconChain.getPubkeyHashes(addedValidators);
        (VALIDATOR_STATUS[] memory curStatuses, VALIDATOR_STATUS[] memory prevStatuses) = _getValidatorStatuses(staker, pubkeyHashes);
        for (uint i = 0; i < curStatuses.length; i++) {
            assertTrue(prevStatuses[i] == VALIDATOR_STATUS.INACTIVE, err);
            assertTrue(curStatuses[i] == VALIDATOR_STATUS.ACTIVE, err);
        }
    }

    function assert_Snap_Removed_ActiveValidators(
        User staker,
        uint40[] memory exitedValidators,
        string memory err
    ) internal {
        bytes32[] memory pubkeyHashes = beaconChain.getPubkeyHashes(exitedValidators);
        (VALIDATOR_STATUS[] memory curStatuses, VALIDATOR_STATUS[] memory prevStatuses) = _getValidatorStatuses(staker, pubkeyHashes);
        for (uint i = 0; i < curStatuses.length; i++) {
            assertTrue(prevStatuses[i] == VALIDATOR_STATUS.ACTIVE, err);
            assertTrue(curStatuses[i] == VALIDATOR_STATUS.WITHDRAWN, err);
        }
    }

    function assert_Snap_Created_Checkpoint(
        User staker,
        string memory err
    ) internal {
        (uint64 curCheckpointTimestamp, uint64 prevCheckpointTimestamp) = _getCheckpointTimestamp(staker);
        assertEq(prevCheckpointTimestamp, 0, err);
        assertTrue(curCheckpointTimestamp != 0, err);
    }

    function assert_Snap_Removed_Checkpoint(
        User staker,
        string memory err
    ) internal {
        (uint64 curCheckpointTimestamp, uint64 prevCheckpointTimestamp) = _getCheckpointTimestamp(staker);
        assertEq(curCheckpointTimestamp, 0, err);
        assertTrue(prevCheckpointTimestamp != 0, err);
    }

    function assert_Snap_Unchanged_Checkpoint(
        User staker,
        string memory err
    ) internal {
        (uint64 curCheckpointTimestamp, uint64 prevCheckpointTimestamp) = _getCheckpointTimestamp(staker);

        assertEq(curCheckpointTimestamp, prevCheckpointTimestamp, err);
    }

    function assert_Snap_Updated_LastCheckpoint(
        User staker,
        string memory err
    ) internal {
        // Sorry for the confusing naming... the pod variable is called `lastCheckpointTimestamp`
        (uint64 curLastCheckpointTimestamp, uint64 prevLastCheckpointTimestamp) = _getLastCheckpointTimestamp(staker);

        assertTrue(curLastCheckpointTimestamp > prevLastCheckpointTimestamp, err);
    }

    function assert_Snap_Added_PodBalanceToWithdrawable(
        User staker,
        string memory err
    ) internal {
        (uint64 curWithdrawableRestakedGwei, uint64 prevWithdrawableRestakedGwei) = _getWithdrawableRestakedGwei(staker);

        (uint64 prevCheckpointPodBalanceGwei, ) = _getCheckpointPodBalanceGwei(staker);

        assertEq(prevWithdrawableRestakedGwei + prevCheckpointPodBalanceGwei, curWithdrawableRestakedGwei, err);
    }

    function assert_Snap_Added_WithdrawableGwei(
        User staker,
        uint64 addedGwei,
        string memory err
    ) internal {
        (uint64 curWithdrawableRestakedGwei, uint64 prevWithdrawableRestakedGwei) = _getWithdrawableRestakedGwei(staker);

        assertEq(prevWithdrawableRestakedGwei + addedGwei, curWithdrawableRestakedGwei, err);
    }

    function assert_Snap_Added_BalanceExitedGwei(
        User staker,
        uint64 addedGwei,
        string memory err
    ) internal {
        (uint64 curCheckpointTimestamp, uint64 prevCheckpointTimestamp) = _getCheckpointTimestamp(staker);

        // If we just finalized a checkpoint, that's the timestamp we want to use
        // to look up checkpoint balances exited
        uint64 targetTimestamp = curCheckpointTimestamp;
        if (curCheckpointTimestamp != prevCheckpointTimestamp) {
            targetTimestamp = prevCheckpointTimestamp;
        }

        (uint64 prevExitedBalanceGwei, uint64 curExitedBalanceGwei) = _getCheckpointBalanceExited(staker, targetTimestamp);
        assertEq(prevExitedBalanceGwei + addedGwei, curExitedBalanceGwei, err);
    }

    function assert_Snap_BCSF_Decreased(
        User staker,
        string memory err
    ) internal {
        (uint64 curBCSF, uint64 prevBCSF) = _getBeaconChainSlashingFactor(staker);
        assertLt(curBCSF, prevBCSF, err);
    }

    function assert_Snap_BCSF_Unchanged(
        User staker,
        string memory err
    ) internal {
        (uint64 curBCSF, uint64 prevBCSF) = _getBeaconChainSlashingFactor(staker);
        assertEq(curBCSF, prevBCSF, err);
    }

    /*******************************************************************************
                                UTILITY METHODS
    *******************************************************************************/
    
    /// @dev Fetches the opreator's allocation delay; asserts that it is set
    function _getExistingAllocationDelay(User operator) internal view returns (uint32) {
        (bool isSet, uint32 delay) = allocationManager.getAllocationDelay(address(operator));
        assertTrue(isSet, "_getExistingAllocationDelay: expected allocation delay to be set");

        return delay;
    }

    /// @dev Generate params to allocate all available magnitude to each strategy in the operator set
    function _genAllocation_AllAvailable(
        User operator, 
        OperatorSet memory operatorSet
    ) internal view returns (AllocateParams memory params) {
        return _genAllocation_AllAvailable({
            operator: operator,
            operatorSet: operatorSet,
            strategies: allocationManager.getStrategiesInOperatorSet(operatorSet)
        });
    }

    /// @dev Generate params to allocate all available magnitude to each strategy in the operator set
    function _genAllocation_AllAvailable(
        User operator, 
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies
    ) internal view returns (AllocateParams memory params) {
        params.operatorSet = operatorSet;
        params.strategies = strategies;
        params.newMagnitudes = new uint64[](params.strategies.length);

        for (uint i = 0; i < params.strategies.length; i++) {
            IStrategy strategy = params.strategies[i];
            params.newMagnitudes[i] = allocationManager.getMaxMagnitude(address(operator), strategy);
        }
    }

    /// @dev Gen params to allocate half of available magnitude to each strategy in the operator set
    /// returns the params to complete this allocation
    function _genAllocation_HalfAvailable(
        User operator, 
        OperatorSet memory operatorSet
    ) internal returns (AllocateParams memory params) {
        return _genAllocation_HalfAvailable({
            operator: operator, 
            operatorSet: operatorSet, 
            strategies: allocationManager.getStrategiesInOperatorSet(operatorSet)
        });
    }

    /// @dev Gen params to allocate half of available magnitude to each strategy in the operator set
    /// returns the params to complete this allocation
    function _genAllocation_HalfAvailable(
        User operator, 
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies
    ) internal returns (AllocateParams memory params) {
        params.operatorSet = operatorSet;
        params.strategies = strategies;
        params.newMagnitudes = new uint64[](params.strategies.length);

        (, Allocation[] memory allocations) = _getAllocations(operator, operatorSet, strategies);
        (, Magnitudes[] memory magnitudes) = _getMagnitudes(operator, strategies);

        for (uint i = 0; i < params.strategies.length; i++) {
            uint64 halfAvailable = uint64(magnitudes[i].allocatable) / 2;
            params.newMagnitudes[i] = allocations[i].currentMagnitude + halfAvailable;
        }
    }

    /// @dev Generate params to allocate a random portion of available magnitude to each strategy
    /// in the operator set. All strategies will have a nonzero allocation, and the minimum allocation
    /// will be 10% of available magnitude
    function _genAllocation_Rand(
        User operator,
        OperatorSet memory operatorSet
    ) internal returns (AllocateParams memory params) {
        params.operatorSet = operatorSet;
        params.strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);
        params.newMagnitudes = new uint64[](params.strategies.length);

        (, Allocation[] memory allocations) = _getAllocations(operator, operatorSet, params.strategies);
        (, Magnitudes[] memory magnitudes) = _getMagnitudes(operator, params.strategies);

        for (uint i = 0; i < params.strategies.length; i++) {
            // minimum of 10%, maximum of 100%. increments of 10%.
            uint r = _randUint({min: 1, max: 10});
            uint64 allocation = uint64(magnitudes[i].allocatable) / uint64(r);

            params.newMagnitudes[i] = allocations[i].currentMagnitude + allocation;
        }
    }

    /// @dev Generates params for a half deallocation from all strategies the operator is allocated to in the operator set
    function _genDeallocation_HalfRemaining(
        User operator,
        OperatorSet memory operatorSet
    ) internal view returns (AllocateParams memory params) {
        return _genDeallocation_HalfRemaining({
            operator: operator, 
            operatorSet: operatorSet, 
            strategies: allocationManager.getStrategiesInOperatorSet(operatorSet)
        });
    }

    /// @dev Generates params for a half deallocation from all strategies the operator is allocated to in the operator set
    function _genDeallocation_HalfRemaining(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies
    ) internal view returns (AllocateParams memory params) {
        params.operatorSet = operatorSet;
        params.strategies = strategies;
        params.newMagnitudes = new uint64[](params.strategies.length);

        for (uint i = 0; i < params.strategies.length; i++) {
            IStrategy strategy = params.strategies[i];
            params.newMagnitudes[i] = allocationManager.getEncumberedMagnitude(address(operator), strategy) / 2;
        }
    }

    /// @dev Generates params for a full deallocation from all strategies the operator is allocated to in the operator set
    function _genDeallocation_Full(
        User operator,
        OperatorSet memory operatorSet
    ) internal view returns (AllocateParams memory params) {
        return _genDeallocation_Full(
            operator, 
            operatorSet, 
            allocationManager.getStrategiesInOperatorSet(operatorSet)
        );
    }

    /// @dev Generates params for a full deallocation from all strategies the operator is allocated to in the operator set
    function _genDeallocation_Full(
        User,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies
    ) internal pure returns (AllocateParams memory params) {
        params.operatorSet = operatorSet;
        params.strategies = strategies;
        params.newMagnitudes = new uint64[](params.strategies.length);
    }

    /// Generate random slashing between 1 and 99%
    function _genSlashing_Rand(
        User operator,
        OperatorSet memory operatorSet
    ) internal returns (SlashingParams memory params) {
        params.operator = address(operator);
        params.operatorSetId = operatorSet.id;
        params.description = "genSlashing_Rand";
        params.strategies = allocationManager.getStrategiesInOperatorSet(operatorSet).sort();
        params.wadsToSlash = new uint[](params.strategies.length);

        /// 1% * rand(1, 99)
        uint slashWad = 1e16 * _randUint({min: 1, max: 99});

        for (uint i = 0; i < params.wadsToSlash.length; i++) {
            params.wadsToSlash[i] = slashWad;
        }
    }

    function _genSlashing_Half(
        User operator,
        OperatorSet memory operatorSet
    ) internal view returns (SlashingParams memory params) {
        params.operator = address(operator);
        params.operatorSetId = operatorSet.id;
        params.description = "genSlashing_Half";
        params.strategies = allocationManager.getStrategiesInOperatorSet(operatorSet).sort();
        params.wadsToSlash = new uint[](params.strategies.length);

        // slash 50%
        for (uint i = 0; i < params.wadsToSlash.length; i++) {
            params.wadsToSlash[i] = 5e17;
        }
    }

    function _genSlashing_Full(
        User operator,
        OperatorSet memory operatorSet
    ) internal view returns (SlashingParams memory params) {
        params.operator = address(operator);
        params.operatorSetId = operatorSet.id;
        params.description = "_genSlashing_Full";
        params.strategies = allocationManager.getStrategiesInOperatorSet(operatorSet).sort();
        params.wadsToSlash = new uint[](params.strategies.length);

        // slash 100%
        for (uint i = 0; i < params.wadsToSlash.length; i++) {
            params.wadsToSlash[i] = 1e18;
        }
    }

    function _randWadToSlash() internal returns (uint) {
        return _randUint({ min: 0.01 ether, max: 1 ether });
    }

    function _strategiesAndWadsForFullSlash(
        OperatorSet memory operatorSet
    ) internal view returns (IStrategy[] memory strategies, uint[] memory wadsToSlash) {
        // Get list of all strategies in an operator set.
        strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);

        wadsToSlash = new uint[](strategies.length);

        for (uint i; i < strategies.length; ++i) {
            wadsToSlash[i] = 1 ether;
        }

        return(strategies.sort(), wadsToSlash);
    }

    function _strategiesAndWadsForRandFullSlash(
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
        
        // Fully slash each selected strategy
        for (uint i; i < len; ++i) {
            wadsToSlash[i] = 1 ether;
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

    function _maxMagnitudes(OperatorSet memory operatorSet, User operator) internal view returns (uint64[] memory magnitudes) {
        IStrategy[] memory strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);
        uint256 len = strategies.length;
        magnitudes = new uint64[](len);

        if (len == 0) return magnitudes;

        for (uint256 i; i < len; ++i) {
            magnitudes[i] = allocationManager.getMaxMagnitude(address(operator), strategies[i]);
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

    function _randSlashType() internal returns (BeaconChainMock.SlashType) {
        return BeaconChainMock.SlashType(_randUint({ min: 0, max: 2 }));
    }

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

                    deltaGwei = -int64(beaconChain.slashValidators(validators, BeaconChainMock.SlashType.Minor));
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
        // TODO: Maybe we update parent method to have an M2 and Slashing version?
        int curPodOwnerShares;
        if (!isUpgraded) {
            curPodOwnerShares = IEigenPodManager_DeprecatedM2(address(eigenPodManager)).podOwnerShares(address(staker));
        } else {
            curPodOwnerShares = eigenPodManager.podOwnerDepositShares(address(staker));
        }
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

    function _calculateExpectedShares(Withdrawal memory withdrawal) internal view returns (uint[] memory) {
        bytes32 root = delegationManager.calculateWithdrawalRoot(withdrawal);

        (, uint[] memory shares) = delegationManager.getQueuedWithdrawal(root);
        return shares;
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
                // We round down expected tokens to the nearest gwei
                expectedTokens[i] = (shares[i] / GWEI_TO_WEI) * GWEI_TO_WEI;
            } else {
                expectedTokens[i] = strat.sharesToUnderlying(shares[i]);
            }
        }

        return expectedTokens;
    }

    function _getWithdrawalHashes(
        Withdrawal[] memory withdrawals
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
    function _rollBlocksForCompleteWithdrawals(Withdrawal[] memory withdrawals) internal {     
        uint256 latest;
        for (uint i = 0; i < withdrawals.length; ++i) {
            if (withdrawals[i].startBlock > latest) latest = withdrawals[i].startBlock;
        }
        cheats.roll(latest + delegationManager.minWithdrawalDelayBlocks() + 1);
    }

    function _rollForward_AllocationDelay(User operator) internal {
        uint32 delay = _getExistingAllocationDelay(operator);
        rollForward(delay);
    }

    function _rollBackward_AllocationDelay(User operator) internal {
        uint32 delay = _getExistingAllocationDelay(operator);
        rollBackward(delay);
    }

    function _rollForward_DeallocationDelay() internal {
        rollForward(allocationManager.DEALLOCATION_DELAY() + 1);
    }

    function _rollBackward_DeallocationDelay() internal {
        rollBackward(allocationManager.DEALLOCATION_DELAY() + 1);
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

    function _getAllocations(User operator, OperatorSet memory operatorSet, IStrategy[] memory strategies) 
        internal 
        returns (Allocation[] memory prev, Allocation[] memory current) 
    {
        prev = new Allocation[](strategies.length);
        current = new Allocation[](strategies.length);
        for (uint i = 0; i < strategies.length; ++i) {
            uint currentState = timeMachine.travelToLast();
            prev[i] = allocationManager.getAllocation(address(operator), operatorSet, strategies[i]);
            timeMachine.travel(currentState);
            current[i] = allocationManager.getAllocation(address(operator), operatorSet, strategies[i]);
        }
    }

    function _getAllocatedStrats(
        User operator,
        OperatorSet memory operatorSet
    ) internal returns (IStrategy[] memory prev, IStrategy[] memory current) {
        uint currentState = timeMachine.travelToLast();
        prev = allocationManager.getAllocatedStrategies(address(operator), operatorSet);
        timeMachine.travel(currentState);
        current = allocationManager.getAllocatedStrategies(address(operator), operatorSet);
    }
    
    function _getAllocatedSets(
        User operator
    ) internal returns (OperatorSet[] memory prev, OperatorSet[] memory current) {
        uint currentState = timeMachine.travelToLast();
        prev = allocationManager.getAllocatedSets(address(operator));
        timeMachine.travel(currentState);
        current = allocationManager.getAllocatedSets(address(operator));
    }

    function _getRegisteredSets(
        User operator
    ) internal returns (OperatorSet[] memory prev, OperatorSet[] memory current) {
        uint currentState = timeMachine.travelToLast();
        prev = allocationManager.getRegisteredSets(address(operator));
        timeMachine.travel(currentState);
        current = allocationManager.getRegisteredSets(address(operator));
    }

    function _getMembers(
        OperatorSet memory operatorSet
    ) internal returns (address[] memory prev, address[] memory current) {
        uint currentState = timeMachine.travelToLast();
        prev = allocationManager.getMembers(operatorSet);
        timeMachine.travel(currentState);
        current = allocationManager.getMembers(operatorSet);
    }

    struct Magnitudes {
        uint256 encumbered;
        uint256 allocatable;
        uint256 max;
    }

    function _getMagnitudes(
        User operator,
        IStrategy[] memory strategies
    ) internal returns (Magnitudes[] memory prev, Magnitudes[] memory current) {
        prev = new Magnitudes[](strategies.length);
        current = new Magnitudes[](strategies.length);
        for (uint i = 0; i < strategies.length; ++i) {
            uint currentState = timeMachine.travelToLast();
            prev[i] = Magnitudes({
                encumbered: allocationManager.getEncumberedMagnitude(address(operator), strategies[i]),
                allocatable: allocationManager.getAllocatableMagnitude(address(operator), strategies[i]),
                max: allocationManager.getMaxMagnitude(address(operator), strategies[i])
            });
            timeMachine.travel(currentState);
            current[i] = Magnitudes({
                encumbered: allocationManager.getEncumberedMagnitude(address(operator), strategies[i]),
                allocatable: allocationManager.getAllocatableMagnitude(address(operator), strategies[i]),
                max: allocationManager.getMaxMagnitude(address(operator), strategies[i])
            });
        }
    }

    function _getMaxMagnitudes(
        User operator,
        IStrategy[] memory strategies
    ) internal view returns (uint64[] memory) {
        return allocationManager.getMaxMagnitudes(address(operator), strategies);
    }

    function _getMaxMagnitudes(
        User operator,
        IStrategy[] memory strategies,
        uint32 blockNum
    ) internal view returns (uint64[] memory) {
        return allocationManager.getMaxMagnitudesAtBlock(address(operator), strategies, blockNum);
    }

    function _getMinSlashableStake(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies
    ) internal returns (uint[] memory prev, uint[] memory current) {
        uint currentState = timeMachine.travelToLast();
        prev = allocationManager.getMinimumSlashableStake({
            operatorSet: operatorSet,
            operators: address(operator).toArray(),
            strategies: strategies,
            futureBlock: uint32(block.number)
        })[0];
        timeMachine.travel(currentState);
        current = allocationManager.getMinimumSlashableStake({
            operatorSet: operatorSet,
            operators: address(operator).toArray(),
            strategies: strategies,
            futureBlock: uint32(block.number)
        })[0];
    }

    function _getAllocatedStake(
        User operator,
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies
    ) internal returns (uint[] memory prev, uint[] memory current) {
        uint currentState = timeMachine.travelToLast();
        prev = allocationManager.getAllocatedStake({
            operatorSet: operatorSet,
            operators: address(operator).toArray(),
            strategies: strategies
        })[0];
        timeMachine.travel(currentState);
        current = allocationManager.getAllocatedStake({
            operatorSet: operatorSet,
            operators: address(operator).toArray(),
            strategies: strategies
        })[0];
    }

    function _getStrategyAllocations(
        User operator,
        IStrategy strategy
    ) internal view returns (OperatorSet[] memory operatorSets, Allocation[] memory allocations) {
        (operatorSets, allocations) = allocationManager.getStrategyAllocations(address(operator), strategy);
    }

    function _getStrategyAllocations(
        address operator,
        IStrategy strategy
    ) internal view returns (OperatorSet[] memory operatorSets, Allocation[] memory allocations) {
        (operatorSets, allocations) = allocationManager.getStrategyAllocations(operator, strategy);
    }

    function _getIsSlashable(
        User operator,
        OperatorSet memory operatorSet
    ) internal returns (bool prev, bool current) {
        uint currentState = timeMachine.travelToLast();
        prev = allocationManager.isOperatorSlashable(address(operator), operatorSet);
        timeMachine.travel(currentState);
        current = allocationManager.isOperatorSlashable(address(operator), operatorSet);
    }

    function _getIsMemberOfSet(
        User operator,
        OperatorSet memory operatorSet
    ) internal returns (bool prev, bool current) {
        uint currentState = timeMachine.travelToLast();
        prev = allocationManager.isMemberOfOperatorSet(address(operator), operatorSet);
        timeMachine.travel(currentState);
        current = allocationManager.isMemberOfOperatorSet(address(operator), operatorSet);
    }

    function _getBurnableShares(IStrategy[] memory strategies) internal returns (uint[] memory prev, uint[] memory current) {
        prev = new uint[](strategies.length);
        current = new uint[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            uint currentState = timeMachine.travelToLast();
            if (strategies[i] == beaconChainETHStrategy) {
                prev[i] = eigenPodManager.burnableETHShares();
                timeMachine.travel(currentState);
                current[i] = eigenPodManager.burnableETHShares();
            } else {
                prev[i] = strategyManager.getBurnableShares(strategies[i]);
                timeMachine.travel(currentState);
                current[i] = strategyManager.getBurnableShares(strategies[i]);
            }
        }
    }

    function _getSlashableSharesInQueue(User operator, IStrategy[] memory strategies) 
        internal 
        returns (uint[] memory prev, uint[] memory current) 
    {
        prev = new uint[](strategies.length);
        current = new uint[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            uint currentState = timeMachine.travelToLast();
            prev[i] = delegationManager.getSlashableSharesInQueue(address(operator), strategies[i]);
            timeMachine.travel(currentState);
            current[i] = delegationManager.getSlashableSharesInQueue(address(operator), strategies[i]);
        }
    }

    /// @dev Looks up each strategy and returns a list of the operator's shares
    function _getOperatorShares(User operator, IStrategy[] memory strategies) 
        internal 
        returns (uint[] memory prev, uint[] memory current) 
    {
        prev = new uint[](strategies.length);
        current = new uint[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            uint currentState = timeMachine.travelToLast();
            prev[i] = delegationManager.operatorShares(address(operator), strategies[i]);
            timeMachine.travel(currentState);
            current[i] = delegationManager.operatorShares(address(operator), strategies[i]);
        }
    }

    /// @dev Looks up each strategy and returns a list of the staker's shares
    function _getStakerDepositShares(User staker, IStrategy[] memory strategies) 
        internal  
        returns (uint[] memory prev, uint[] memory current) 
    {
        prev = new uint[](strategies.length);
        current = new uint[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint currentState = timeMachine.travelToLast();
            if (strat == BEACONCHAIN_ETH_STRAT) {
                int shares = eigenPodManager.podOwnerDepositShares(address(staker));
                require(shares >= 0, "_getStakerDepositShares: negative shares");
                prev[i] = uint(shares);
                timeMachine.travel(currentState);
                shares = eigenPodManager.podOwnerDepositShares(address(staker));
                require(shares >= 0, "_getStakerDepositShares: negative shares");
                current[i] = uint(shares);
            } else {
                prev[i] = strategyManager.stakerDepositShares(address(staker), strat);
                timeMachine.travel(currentState);
                current[i] = strategyManager.stakerDepositShares(address(staker), strat);
            }
        }
    }

    /// @dev Looks up each strategy and returns a list of the staker's shares
    function _getStakerDepositSharesInt(User staker, IStrategy[] memory strategies) 
        internal 
        returns (int[] memory prev, int[] memory current) 
    {
        prev = new int[](strategies.length);
        current = new int[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint currentState = timeMachine.travelToLast();
            if (strat == BEACONCHAIN_ETH_STRAT) {
                prev[i] = eigenPodManager.podOwnerDepositShares(address(staker));
                timeMachine.travel(currentState);
                current[i] = eigenPodManager.podOwnerDepositShares(address(staker));
            } else {
                prev[i] = int(strategyManager.stakerDepositShares(address(staker), strat));
                timeMachine.travel(currentState);
                current[i] = int(strategyManager.stakerDepositShares(address(staker), strat));
            }
        }
    }

    function _getStakerStrategyList(User staker) internal view returns (IStrategy[] memory) {
        return strategyManager.getStakerStrategyList(address(staker));
    }

    function _getStakerWithdrawableShares(User staker, IStrategy[] memory strategies) 
        internal 
        returns (uint[] memory prev, uint[] memory current) 
    {
        uint currentState = timeMachine.travelToLast();
        (prev, ) = delegationManager.getWithdrawableShares(address(staker), strategies);
        timeMachine.travel(currentState);
        (, current) = delegationManager.getWithdrawableShares(address(staker), strategies);
    }

    function _calcWithdrawable(User staker, IStrategy[] memory strategies, uint[] memory depositSharesToWithdraw) internal returns (uint[] memory) {
        uint[] memory withdrawableShares = new uint[](strategies.length);
        (, uint[] memory depositScalingFactors) = _getDepositScalingFactors(staker, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            withdrawableShares[i] = depositSharesToWithdraw[i].mulWad(depositScalingFactors[i]).mulWad(_getSlashingFactor(staker, strategies[i]));
        }
        return withdrawableShares;
    }
    
    /// @dev Looks up the staker's beacon chain scaling factor
    function _getBeaconChainSlashingFactor(User staker) internal returns (uint64 prev, uint64 current) {
        uint currentState = timeMachine.travelToLast();
        prev = eigenPodManager.beaconChainSlashingFactor(address(staker));
        timeMachine.travel(currentState);
        current = eigenPodManager.beaconChainSlashingFactor(address(staker));
    }

    function _getCumulativeWithdrawals(User staker) internal returns (uint prev, uint current) {
        uint currentState = timeMachine.travelToLast();
        prev = delegationManager.cumulativeWithdrawalsQueued(address(staker));
        timeMachine.travel(currentState);
        current = delegationManager.cumulativeWithdrawalsQueued(address(staker));
    }

    function _getTokenBalances(User staker, IERC20[] memory tokens) internal returns (uint[] memory prev, uint[] memory current) {
        prev = new uint[](tokens.length);
        current = new uint[](tokens.length);
        for (uint i = 0; i < tokens.length; i++) {
            uint currentState = timeMachine.travelToLast();
            if (tokens[i] == NATIVE_ETH) {
                prev[i] = address(staker).balance;
                timeMachine.travel(currentState);
                current[i] = address(staker).balance;
            } else {
                prev[i] = tokens[i].balanceOf(address(staker));
                timeMachine.travel(currentState);
                current[i] = tokens[i].balanceOf(address(staker));
            }
        }
    }

    function _getTotalStrategyShares(IStrategy[] memory strategies) internal returns (uint[] memory prev, uint[] memory current) {
        prev = new uint[](strategies.length);
        current = new uint[](strategies.length);
        // BeaconChainETH strategy doesn't keep track of global strategy shares, so we ignore
        for (uint i = 0; i < strategies.length; i++) {
            uint currentState = timeMachine.travelToLast();
            if (strategies[i] != BEACONCHAIN_ETH_STRAT) {
                prev[i] = strategies[i].totalShares();
                timeMachine.travel(currentState);
                current[i] = strategies[i].totalShares();
            }
        }
    }

    function _getDepositScalingFactors(User staker, IStrategy[] memory strategies) internal returns (uint[] memory prev, uint[] memory current) {
        prev = new uint[](strategies.length);
        current = new uint[](strategies.length);
        for (uint i=0; i < strategies.length; i++) {
            uint currentState = timeMachine.travelToLast();
            prev[i] = delegationManager.depositScalingFactor(address(staker), strategies[i]);
            timeMachine.travel(currentState);
            current[i] = delegationManager.depositScalingFactor(address(staker), strategies[i]);
        }
        return (prev, current);
    }

    function _getExpectedDSFUndelegate(User staker) internal returns (uint expectedDepositScalingFactor) {
        (, uint64 curBCSF) = _getBeaconChainSlashingFactor(staker);
        return WAD.divWad(curBCSF);
    }

    function _getExpectedDSFDeposit(User staker, User operator, IStrategy strategy) internal returns (uint expectedDepositScalingFactor) {
        if (strategy == BEACONCHAIN_ETH_STRAT){
            (, uint64 curBCSF) = _getBeaconChainSlashingFactor(staker);
            return WAD.divWad(allocationManager.getMaxMagnitude(address(operator), strategy).mulWad(curBCSF));
        } else {
            return WAD.divWad(allocationManager.getMaxMagnitude(address(operator), strategy));
        }
    }

    function _getExpectedWithdrawableSharesUndelegate(User staker, IStrategy[] memory strategies, uint[] memory shares) internal returns (uint[] memory){
        uint[] memory expectedWithdrawableShares = new uint[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            if (strategies[i] == BEACONCHAIN_ETH_STRAT) {
                (, uint64 curBCSF) = _getBeaconChainSlashingFactor(staker);
                expectedWithdrawableShares[i] = shares[i].mulWad(_getExpectedDSFUndelegate(staker)).mulWad(curBCSF);
            } else {
                expectedWithdrawableShares[i] = shares[i];
             }
        }
        return expectedWithdrawableShares;
    }

    function _getExpectedDSFsDelegate(User staker, User operator, IStrategy[] memory strategies) internal returns (uint[] memory) {
        uint[] memory expectedDepositScalingFactors = new uint[](strategies.length);
        (uint[] memory oldDepositScalingFactors, ) = _getDepositScalingFactors(staker, strategies);
        uint64[] memory maxMagnitudes = _getMaxMagnitudes(operator, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            expectedDepositScalingFactors[i] = oldDepositScalingFactors[i].divWad(maxMagnitudes[i]);
        }
        return expectedDepositScalingFactors;
    }

    function _getExpectedWithdrawableSharesDelegate(User staker, User operator, IStrategy[] memory strategies, uint[] memory depositShares) internal returns (uint[] memory){
        uint[] memory expectedWithdrawableShares = new uint[](strategies.length);
        uint[] memory expectedDSFs = _getExpectedDSFsDelegate(staker, operator, strategies);
        uint64[] memory maxMagnitudes = _getMaxMagnitudes(operator, strategies);
        for (uint i = 0; i < strategies.length; i++) {
            if (strategies[i] == BEACONCHAIN_ETH_STRAT) {
                (, uint64 curBCSF) = _getBeaconChainSlashingFactor(staker);
                expectedWithdrawableShares[i] = depositShares[i].mulWad(expectedDSFs[i]).mulWad(maxMagnitudes[i]).mulWad(curBCSF);
            } else {
                expectedWithdrawableShares[i] = depositShares[i].mulWad(expectedDSFs[i]).mulWad(maxMagnitudes[i]);
            }
        }
        return expectedWithdrawableShares;
    }

    function _getExpectedWithdrawableSharesDeposit(User staker, User operator, IStrategy strategy, uint depositShares) internal returns (uint) {
        return depositShares.mulWad(_getExpectedDSFDeposit(staker, operator, strategy)).mulWad(_getSlashingFactor(staker, strategy));
    }

    function _getSlashingFactor(
        User staker,
        IStrategy strategy
    ) internal view returns (uint256) {
        address operator = delegationManager.delegatedTo(address(staker));
        uint64 maxMagnitude = allocationManager.getMaxMagnitudes(operator, strategy.toArray())[0];
        if (strategy == beaconChainETHStrategy) {
            return maxMagnitude.mulWad(eigenPodManager.beaconChainSlashingFactor(address(staker)));
        }
        return maxMagnitude;
    }

    function _getWithdrawableShares(User staker, IStrategy[] memory strategies) internal returns (uint[] memory prev, uint[] memory current) {
        uint currentState = timeMachine.travelToLast();
        (prev, ) =  delegationManager.getWithdrawableShares(address(staker), strategies);
        timeMachine.travel(currentState);
        (, current) =  delegationManager.getWithdrawableShares(address(staker), strategies);
    }

    function _getWithdrawableShares(User staker, IStrategy strategy) internal returns (uint prev, uint current) {
        uint currentState = timeMachine.travelToLast();
        (uint[] memory prev, ) = delegationManager.getWithdrawableShares(address(staker), strategy.toArray());
        timeMachine.travel(currentState);
        (, uint[] memory current) = delegationManager.getWithdrawableShares(address(staker), strategy.toArray());
        return (prev[0], current[0]);
    }

    function _getActiveValidatorCount(User staker) internal returns (uint prev, uint current) {
        EigenPod pod = staker.pod();
        uint currentState = timeMachine.travelToLast();
        prev = pod.activeValidatorCount();
        timeMachine.travel(currentState);
        current = pod.activeValidatorCount();
    }

    function _getValidatorStatuses(User staker, bytes32[] memory pubkeyHashes) internal returns (VALIDATOR_STATUS[] memory prev, VALIDATOR_STATUS[] memory current) {
        EigenPod pod = staker.pod();
        prev = new VALIDATOR_STATUS[](pubkeyHashes.length);
        current = new VALIDATOR_STATUS[](pubkeyHashes.length);

        for (uint i = 0; i < pubkeyHashes.length; i++) {
            uint currentState = timeMachine.travelToLast();
            prev[i] = pod.validatorStatus(pubkeyHashes[i]);
            timeMachine.travel(currentState);
            current[i] = pod.validatorStatus(pubkeyHashes[i]);
        }
    }

    function _getCheckpointTimestamp(User staker) internal returns (uint64 prev, uint64 current) {
        EigenPod pod = staker.pod();
        uint currentState = timeMachine.travelToLast();
        prev = pod.currentCheckpointTimestamp();
        timeMachine.travel(currentState);
        current = pod.currentCheckpointTimestamp();
    }

    function _getLastCheckpointTimestamp(User staker) internal returns (uint64 prev, uint64 current) {
        EigenPod pod = staker.pod();
        uint currentState = timeMachine.travelToLast();
        prev = pod.lastCheckpointTimestamp();
        timeMachine.travel(currentState);
        current = pod.lastCheckpointTimestamp();
    }

    function _getWithdrawableRestakedGwei(User staker) internal returns (uint64 prev, uint64 current) {
        EigenPod pod = staker.pod();
        uint currentState = timeMachine.travelToLast();
        prev = pod.withdrawableRestakedExecutionLayerGwei();
        timeMachine.travel(currentState);
        current = pod.withdrawableRestakedExecutionLayerGwei();
    }

    function _getCheckpointPodBalanceGwei(User staker) internal returns (uint64 prev, uint64 current) {
        if (forkType != LOCAL && !isUpgraded) {
            IEigenPod_DeprecatedM2 pod = IEigenPod_DeprecatedM2(address(staker.pod()));
            uint currentState = timeMachine.travelToLast();
            prev = uint64(pod.currentCheckpoint().podBalanceGwei);
            timeMachine.travel(currentState);
            current = uint64(pod.currentCheckpoint().podBalanceGwei);
        } else {
            EigenPod pod = staker.pod();
            uint currentState = timeMachine.travelToLast();
            prev = uint64(pod.currentCheckpoint().podBalanceGwei);
            timeMachine.travel(currentState);
            current = uint64(pod.currentCheckpoint().podBalanceGwei);
        }
    }

    function _getCheckpointBalanceExited(User staker, uint64 checkpointTimestamp) internal returns (uint64 prev, uint64 current) {
        EigenPod pod = staker.pod();
        uint currentState = timeMachine.travelToLast();
        prev = pod.checkpointBalanceExitedGwei(checkpointTimestamp);
        timeMachine.travel(currentState);
        current = pod.checkpointBalanceExitedGwei(checkpointTimestamp);
    }

    function _getQueuedWithdrawals(User staker) internal view returns (Withdrawal[] memory) {
        (Withdrawal[] memory withdrawals,) = delegationManager.getQueuedWithdrawals(address(staker));
        return withdrawals;
    }
}