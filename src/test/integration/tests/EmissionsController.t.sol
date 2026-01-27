// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/IncentiveCouncil.t.sol";

/// @notice Base contract for EmissionsController integration tests with shared setup
contract Integration_EmissionsController_Base is IntegrationCheckUtils, IEmissionsControllerTypes {
    using ArrayLib for *;

    address incentiveCouncil;
    IncentiveCouncil ic;
    AVS avs;
    OperatorSet[] operatorSets;
    IStrategy[] strategies;

    function _init() internal virtual override {
        _configAssetTypes(HOLDS_LST);
        OperatorSet[] memory _operatorSets;
        // Get incentive council address
        incentiveCouncil = emissionsController.incentiveCouncil();
        ic = new IncentiveCouncil(emissionsController, incentiveCouncil, timeMachine);
        // Enable emissions controller as rewards submitter
        cheats.prank(rewardsCoordinator.owner());
        rewardsCoordinator.setRewardsForAllSubmitter(address(emissionsController), true);
        // Create AVS and operator for operator set distributions
        (avs, _operatorSets) = _newRandomAVS();
        strategies = allStrats;
        for (uint i = 0; i < _operatorSets.length; i++) {
            operatorSets.push(_operatorSets[i]);
        }
    }

    function _genEvenlyDistributedWeights(uint numWeights, uint16 totalWeight) internal returns (uint16[] memory weights) {
        weights = new uint16[](numWeights);
        uint16 weightPerWeight = uint16(totalWeight / numWeights);
        uint16 remainder = uint16(totalWeight % numWeights);

        for (uint i = 0; i < numWeights; ++i) {
            weights[i] = weightPerWeight;
            // Distribute remainder across first few weights
            if (i < remainder) weights[i] += 1;
        }
        return weights;
    }

    function _genRandParams() internal returns (uint64 startEpoch, uint16 totalWeight, uint numDistributions) {
        startEpoch = uint64(_randUint({min: 0, max: 2**12 - 1}));
        totalWeight = uint16(_randUint({min: 1, max: 10_000})); // 0.01-100% (bips)
        numDistributions = _randUint({min: 1, max: 32});
        return (startEpoch, totalWeight, numDistributions);
    }
}

contract Integration_EmissionsController_E2E is Integration_EmissionsController_Base {
    /// -----------------------------------------------------------------------
    /// ALL BEFORE emissions start
    /// -----------------------------------------------------------------------
    function testFuzz_addDists_pressButton_noneProcessed(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();

        // 1. Add distributions
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributions(operatorSets, strategies, numDistributions, startEpoch, totalWeight);
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Press button (should revert `EmissionsNotStarted()`)
        vm.expectRevert(IEmissionsControllerErrors.EmissionsNotStarted.selector);
        ic.pressButton(numDistributions);
    }

    function testFuzz_addDists_updateDists_pressButton_noneProcessed(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();

        // 1. Add distributions
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributions(operatorSets, strategies, numDistributions, startEpoch, totalWeight);
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Update distributions
        Distribution[] memory updatedDistributions =
            ic.updateDistributions(distributionIds, distributions, _genEvenlyDistributedWeights(numDistributions, totalWeight));
        check_updateDists_State(distributionIds, distributions, updatedDistributions);

        // 3. Press button
        vm.expectRevert(IEmissionsControllerErrors.EmissionsNotStarted.selector);
        ic.pressButton(numDistributions);
    }

    function testFuzz_addDists_pressButton_noneProcessed_sweep_noneSwept(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();

        // 1. Add distributions
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributions(operatorSets, strategies, numDistributions, startEpoch, totalWeight);
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Press button
        vm.expectRevert(IEmissionsControllerErrors.EmissionsNotStarted.selector);
        ic.pressButton(numDistributions);

        // 3. Sweep
        bool swept = ic.sweep();
        check_sweep_State(distributionIds, distributions, swept, false);
    }

    function testFuzz_addDists_updateDists_pressButton_noneProcessed_sweep_noneSwept(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();

        // 1. Add distributions
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributions(operatorSets, strategies, numDistributions, startEpoch, totalWeight);
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Update distributions
        Distribution[] memory updatedDistributions =
            ic.updateDistributions(distributionIds, distributions, _genEvenlyDistributedWeights(numDistributions, totalWeight));
        check_updateDists_State(distributionIds, distributions, updatedDistributions);

        // 3. Press button
        vm.expectRevert(IEmissionsControllerErrors.EmissionsNotStarted.selector);
        ic.pressButton(numDistributions);

        // 4. Sweep
        bool swept = ic.sweep();
        check_sweep_State(distributionIds, updatedDistributions, swept, false);
    }

    function testFuzz_addDists_sweep_noneSwept(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();

        // 1. Add distributions
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributions(operatorSets, strategies, numDistributions, startEpoch, totalWeight);
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Sweep
        bool swept = ic.sweep();
        check_sweep_State(distributionIds, distributions, swept, false);
    }

    function testFuzz_addDists_updateDists_sweep_noneSwept(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();

        // 1. Add distributions
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributions(operatorSets, strategies, numDistributions, startEpoch, totalWeight);
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Update distributions
        Distribution[] memory updatedDistributions =
            ic.updateDistributions(distributionIds, distributions, _genEvenlyDistributedWeights(numDistributions, totalWeight));
        check_updateDists_State(distributionIds, distributions, updatedDistributions);

        // 3. Sweep
        bool swept = ic.sweep();
        check_sweep_State(distributionIds, updatedDistributions, swept, false);
    }

    /// -----------------------------------------------------------------------
    /// BEFORE → AFTER transition (add before, then something after)
    /// -----------------------------------------------------------------------

    function testFuzz_addDists_warp_pressButton(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();

        // 1. Add distributions
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributions(operatorSets, strategies, numDistributions, startEpoch, totalWeight);
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Warp to start epoch
        ic.warpToEpoch(startEpoch);
        check_warpToEpoch_State(startEpoch);

        // 3. Press button
        uint processed = ic.pressButton(numDistributions);
        check_pressButton_State(distributionIds, distributions, processed, numDistributions);

        // 4. Press button again (should revert).
        vm.expectRevert(IEmissionsControllerErrors.AllDistributionsProcessed.selector);
        ic.pressButton(numDistributions);
    }

    function testFuzz_addDists_updateDists_warp_pressButton(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();

        // 1. Add distributions
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributions(operatorSets, strategies, numDistributions, startEpoch, totalWeight);
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Update distributions
        Distribution[] memory updatedDistributions =
            ic.updateDistributions(distributionIds, distributions, _genEvenlyDistributedWeights(numDistributions, totalWeight));
        check_updateDists_State(distributionIds, distributions, updatedDistributions);

        // 3. Warp to start epoch
        ic.warpToEpoch(startEpoch);
        check_warpToEpoch_State(startEpoch);

        // 4. Press button
        uint processed = ic.pressButton(numDistributions);
        check_pressButton_State(distributionIds, updatedDistributions, processed, numDistributions);

        // 4. Press button again (should revert).
        vm.expectRevert(IEmissionsControllerErrors.AllDistributionsProcessed.selector);
        ic.pressButton(numDistributions);
    }

    function testFuzz_addDists_warp_pressButton_sweep(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();

        // 1. Add distributions
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributions(operatorSets, strategies, numDistributions, startEpoch, totalWeight);
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Warp to start epoch
        ic.warpToEpoch(startEpoch);
        check_warpToEpoch_State(startEpoch);

        // 3. Press button
        uint processed = ic.pressButton(numDistributions);
        check_pressButton_State(distributionIds, distributions, processed, numDistributions);

        // 4. Sweep
        bool expectedSwept = emissionsController.totalWeight() < 10_000;
        bool swept = ic.sweep();
        check_sweep_State(distributionIds, distributions, swept, expectedSwept);

        // 5. Press button again (should revert).
        vm.expectRevert(IEmissionsControllerErrors.AllDistributionsProcessed.selector);
        ic.pressButton(numDistributions);
    }

    function testFuzz_addDists_updateDists_warp_pressButton_sweep(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();

        // 1. Add distributions
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributions(operatorSets, strategies, numDistributions, startEpoch, totalWeight);
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Update distributions
        Distribution[] memory updatedDistributions =
            ic.updateDistributions(distributionIds, distributions, _genEvenlyDistributedWeights(numDistributions, totalWeight));
        check_updateDists_State(distributionIds, distributions, updatedDistributions);

        // 3. Warp to start epoch
        ic.warpToEpoch(startEpoch);
        check_warpToEpoch_State(startEpoch);

        // 4. Press button
        uint processed = ic.pressButton(numDistributions);
        check_pressButton_State(distributionIds, updatedDistributions, processed, numDistributions);

        // 5. Sweep
        bool expectedSwept = emissionsController.totalWeight() < 10_000;
        bool swept = ic.sweep();
        check_sweep_State(distributionIds, updatedDistributions, swept, expectedSwept);

        // 6. Press button again (should revert).
        vm.expectRevert(IEmissionsControllerErrors.AllDistributionsProcessed.selector);
        ic.pressButton(numDistributions);
    }

    function testFuzz_addDists_warp_updateDists_pressButton(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();

        // 1. Add distributions
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributions(operatorSets, strategies, numDistributions, startEpoch, totalWeight);
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Warp to start epoch
        ic.warpToEpoch(startEpoch);
        check_warpToEpoch_State(startEpoch);

        // 3. Update distributions
        vm.expectRevert(IEmissionsControllerErrors.AllDistributionsMustBeProcessed.selector);
        ic.updateDistributions(distributionIds, distributions, _genEvenlyDistributedWeights(numDistributions, totalWeight));

        // 4. Press button
        uint processed = ic.pressButton(numDistributions);
        check_pressButton_State(distributionIds, distributions, processed, numDistributions);

        // 5. Press button again (should revert).
        vm.expectRevert(IEmissionsControllerErrors.AllDistributionsProcessed.selector);
        ic.pressButton(numDistributions);
    }

    function testFuzz_addDists_warp_updateDists_pressButton_sweep(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();

        // 1. Add distributions
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributions(operatorSets, strategies, numDistributions, startEpoch, totalWeight);
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Warp to start epoch
        ic.warpToEpoch(startEpoch);
        check_warpToEpoch_State(startEpoch);

        // 3. Update distributions
        vm.expectRevert(IEmissionsControllerErrors.AllDistributionsMustBeProcessed.selector);
        ic.updateDistributions(distributionIds, distributions, _genEvenlyDistributedWeights(numDistributions, totalWeight));

        // 4. Press button
        uint processed = ic.pressButton(numDistributions);
        check_pressButton_State(distributionIds, distributions, processed, numDistributions);

        // 5. Sweep
        bool expectedSwept = emissionsController.totalWeight() < 10_000;
        bool swept = ic.sweep();
        check_sweep_State(distributionIds, distributions, swept, expectedSwept);

        // 6. Press button again (should revert).
        vm.expectRevert(IEmissionsControllerErrors.AllDistributionsProcessed.selector);
        ic.pressButton(numDistributions);
    }

    /// -----------------------------------------------------------------------
    /// Invariants for pressButton
    /// -----------------------------------------------------------------------

    /// @dev Assert that all distributions succeed given valid inputs.
    function testFuzz_addDists_pressButton_allDistributionTypes(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight,) = _genRandParams();

        DistributionType[] memory types = new DistributionType[](5);
        types[0] = DistributionType.RewardsForAllEarners;
        types[1] = DistributionType.OperatorSetTotalStake;
        types[2] = DistributionType.OperatorSetUniqueStake;
        types[3] = DistributionType.EigenDA;
        types[4] = DistributionType.Manual;

        // 1. Add distributions with all types
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributionsWithTypes(operatorSets, strategies, types, startEpoch, totalWeight);
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Warp to start epoch
        ic.warpToEpoch(startEpoch);
        check_warpToEpoch_State(startEpoch);

        // 3. Press button - should process all distributions
        uint processed = ic.pressButton(types.length);
        check_pressButton_State(distributionIds, distributions, processed, types.length);

        // 4. Press button again (should revert).
        vm.expectRevert(IEmissionsControllerErrors.AllDistributionsProcessed.selector);
        ic.pressButton(types.length);
    }

    /// @dev Assert that the function skips disabled distributions.
    function testFuzz_addDists_pressButton_skipsDisabledDistributions(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();
        numDistributions = _randUint({min: 2, max: 32}); // Need at least 2 distributions

        // Create distribution types (all enabled initially)
        DistributionType[] memory types = new DistributionType[](numDistributions);
        uint numToDisable = _randUint({min: 1, max: numDistributions - 1}); // At least 1 disabled, at least 1 enabled

        for (uint i = 0; i < numDistributions; ++i) {
            types[i] = DistributionType(uint8(_randUint({min: 1, max: 5}))); // Non-disabled types
        }

        // 1. Add distributions (all enabled)
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributionsWithTypes(operatorSets, strategies, types, startEpoch, totalWeight);
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Update some distributions to disabled
        for (uint i = 0; i < numToDisable; ++i) {
            distributions[i].distributionType = DistributionType.Disabled;
            cheats.prank(incentiveCouncil);
            emissionsController.updateDistribution(distributionIds[i], distributions[i]);
        }

        // Note: Disabling distributions doesn't change totalWeight - it just affects processing
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 3. Warp to start epoch
        ic.warpToEpoch(startEpoch);
        check_warpToEpoch_State(startEpoch);

        // 4. Press button - should only process non-disabled distributions
        uint processed = ic.pressButton(numDistributions);
        uint expectedProcessed = numDistributions - numToDisable;
        assertEq(processed, expectedProcessed, "Should skip disabled distributions");
    }

    /// @dev Assert that the function skips distributions that have not started yet.
    function testFuzz_addDists_pressButton_skipsNotYetStartedDistributions(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();
        numDistributions = _randUint({min: 2, max: 32}); // Need at least 2 distributions
        startEpoch = uint64(_randUint({min: 2, max: 2**12 - 1})); // Ensure startEpoch > 0

        // Generate random distribution types (non-disabled)
        DistributionType[] memory types = new DistributionType[](numDistributions);
        for (uint i = 0; i < numDistributions; ++i) {
            types[i] = DistributionType(uint8(_randUint({min: 1, max: 5})));
        }

        // 1. Add distributions with staggered start epochs
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributionsWithTypes(operatorSets, strategies, types, startEpoch, totalWeight);

        // Manually update some distributions to have later start epochs
        uint numLaterStart = _randUint({min: 1, max: numDistributions - 1});
        for (uint i = 0; i < numLaterStart; ++i) {
            distributions[i].startEpoch = startEpoch + 1; // Start one epoch later
            cheats.prank(incentiveCouncil);
            emissionsController.updateDistribution(distributionIds[i], distributions[i]);
        }
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Warp to start epoch (some distributions haven't started yet)
        ic.warpToEpoch(startEpoch);
        check_warpToEpoch_State(startEpoch);

        // 3. Press button - should only process distributions that have started
        uint processed = ic.pressButton(numDistributions);
        uint expectedProcessed = numDistributions - numLaterStart;
        assertEq(processed, expectedProcessed, "Should skip distributions that haven't started");
    }

    /// @dev Assert that the function skips distributions that have ended.
    function testFuzz_addDists_pressButton_skipsEndedDistributions(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();
        numDistributions = _randUint({min: 2, max: 32}); // Need at least 2 distributions
        startEpoch = uint64(_randUint({min: 1, max: 100})); // Keep epochs reasonable

        // Generate random distribution types (non-disabled)
        DistributionType[] memory types = new DistributionType[](numDistributions);
        for (uint i = 0; i < numDistributions; ++i) {
            types[i] = DistributionType(uint8(_randUint({min: 1, max: 5})));
        }

        // 1. Add distributions (all have totalEpochs = 1 by default, meaning they end after first epoch)
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributionsWithTypes(operatorSets, strategies, types, startEpoch, totalWeight);

        // Update some distributions to have totalEpochs = 0 (infinite, so they DON'T end)
        // The rest keep totalEpochs = 1 (end after startEpoch)
        uint numInfinite = _randUint({min: 1, max: numDistributions - 1});
        for (uint i = 0; i < numInfinite; ++i) {
            distributions[i].totalEpochs = 0; // Infinite - will be active in all epochs
            cheats.prank(incentiveCouncil);
            emissionsController.updateDistribution(distributionIds[i], distributions[i]);
        }
        uint numTimeLimited = numDistributions - numInfinite;
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Warp to the first epoch where time-limited distributions are active
        ic.warpToEpoch(startEpoch);
        check_warpToEpoch_State(startEpoch);

        // 3. Press button - all distributions should be processed
        uint processed = ic.pressButton(numDistributions);
        assertEq(processed, numDistributions, "All distributions should be processed in their start epoch");

        // 4. Verify button is not pressable (all distributions processed for this epoch)
        assertFalse(emissionsController.isButtonPressable(), "Button should not be pressable after all distributions processed");

        // 5. Warp to the next epoch where time-limited distributions have ended
        // Distributions with totalEpochs = 1 end when currentEpoch >= startEpoch + totalEpochs
        // i.e., currentEpoch >= startEpoch + 1
        ic.warpToEpoch(startEpoch + 1);
        check_warpToEpoch_State(startEpoch + 1);

        // 6. Press button - should skip ended distributions
        // In the new epoch, only infinite distributions (totalEpochs = 0) should be processed
        // Time-limited distributions (totalEpochs = 1) have ended
        processed = ic.pressButton(numDistributions);
        assertEq(processed, numInfinite, "Should skip ended distributions and only process infinite ones");
    }

    /// @dev Assert that the function skips distributions with zero weight.
    function testFuzz_addDists_pressButton_skipsZeroWeightDistributions(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();
        numDistributions = _randUint({min: 2, max: 32}); // Need at least 2 distributions

        // Generate random distribution types (non-disabled)
        DistributionType[] memory types = new DistributionType[](numDistributions);
        for (uint i = 0; i < numDistributions; ++i) {
            types[i] = DistributionType(uint8(_randUint({min: 1, max: 5})));
        }

        // 1. Add distributions
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributionsWithTypes(operatorSets, strategies, types, startEpoch, totalWeight);

        // Update some distributions to have zero weight
        uint numZeroWeight = _randUint({min: 1, max: numDistributions - 1});
        for (uint i = 0; i < numZeroWeight; ++i) {
            distributions[i].weight = 0;
            cheats.prank(incentiveCouncil);
            emissionsController.updateDistribution(distributionIds[i], distributions[i]);
        }

        // Recalculate expected totalWeight after zero-weight updates
        uint16 expectedTotalWeight = 0;
        for (uint i = 0; i < distributions.length; ++i) {
            expectedTotalWeight += uint16(distributions[i].weight);
        }
        check_addDists_State(distributionIds, distributions, expectedTotalWeight);

        // 2. Warp to start epoch
        ic.warpToEpoch(startEpoch);
        check_warpToEpoch_State(startEpoch);

        // 3. Press button - should skip zero-weight distributions
        // Note: Zero-weight distributions don't emit DistributionProcessed events
        uint processed = ic.pressButton(numDistributions);
        uint expectedProcessed = numDistributions - numZeroWeight;
        assertEq(processed, expectedProcessed, "Should skip zero-weight distributions");
    }

    /// @dev Assert that the function processes zero distributions when length is zero.
    function testFuzz_addDists_pressButtonLengthZero_NoneProcessed(uint24 r) public rand(r) {
        (uint64 startEpoch, uint16 totalWeight, uint numDistributions) = _genRandParams();

        // 1. Add distributions
        (uint[] memory distributionIds, Distribution[] memory distributions) =
            ic.addDistributions(operatorSets, strategies, numDistributions, startEpoch, totalWeight);
        check_addDists_State(distributionIds, distributions, totalWeight);

        // 2. Warp to start epoch
        ic.warpToEpoch(startEpoch);
        check_warpToEpoch_State(startEpoch);

        // 3. Press button with length 0 - should process zero distributions
        uint processed = ic.pressButton(0);
        assertEq(processed, 0, "Should process zero distributions when length is zero");

        // 4. Button should still be pressable since distributions remain
        assertTrue(emissionsController.isButtonPressable(), "Button should still be pressable");
    }

    /// -----------------------------------------------------------------------
    /// Edge cases
    /// -----------------------------------------------------------------------

    /// @dev Assert that sweep returns false when there are no distributions and no emissions minted.
    function testFuzz_noDistributions_pressButton_NoneProcessed_sweep_AllEmissionsShouldBeSwept(uint24 r) public rand(r) {
        uint64 startEpoch = uint64(_randUint({min: 0, max: 100}));

        // 1. Warp to start epoch (no distributions added)
        ic.warpToEpoch(startEpoch);
        check_warpToEpoch_State(startEpoch);

        // 2. When there are no distributions, pressButton should revert with AllDistributionsProcessed
        // because totalProcessable = 0 and totalProcessed = 0 (i.e., all 0 distributions are "processed")
        vm.expectRevert(IEmissionsControllerErrors.AllDistributionsProcessed.selector);
        ic.pressButton(1);

        // 3. Sweep - since pressButton was never successfully called, no emissions were minted,
        // so sweep should return false (nothing to sweep)
        bool swept = ic.sweep();
        assertFalse(swept, "Should not sweep when no emissions were minted");
    }
}
