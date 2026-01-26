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

    /// @dev Assert that all distributions should succeed given valid inputs.
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

    // /// @dev Assert that the function mints only once per epoch.
    // function testFuzz_addDists_pressButton_onlyMintedOncePerEpoch(uint24 r) public rand(r) {}

    // /// @dev Assert that the function skips disabled distributions.
    // function testFuzz_addDists_pressButton_skipsDisabledDistributions(uint24 r) public rand(r) {}

    // /// @dev Assert that the function skips distributions that have not started yet.
    // function testFuzz_addDists_pressButton_skipsNotYetStartedDistributions(uint24 r) public rand(r) {}

    // /// @dev Assert that the function skips distributions that have ended.
    // function testFuzz_addDists_pressButton_skipsEndedDistributions(uint24 r) public rand(r) {}

    // /// @dev Assert that the function skips distributions with zero weight.
    // function testFuzz_addDists_pressButton_skipsZeroWeightDistributions(uint24 r) public rand(r) {}

    // /// @dev Assert that the function processes zero distributions when length is zero.
    // function testFuzz_addDists_pressButtonLengthZero_NoneProcessed(uint24 r) public rand(r) {}

    // /// @dev Assert that the function can be called at any time within an epoch.
    // function testFuzz_addDists_pressButton_canPressButtonAtAnyTimeWithinEpoch(uint24 r) public rand(r) {}

    /// -----------------------------------------------------------------------
    /// Edge cases
    /// -----------------------------------------------------------------------

    // /// @dev Assert that the function processes zero distributions when there are no distributions.
    // function testFuzz_noDistributions_pressButton_NoneProcessed_sweep_AllEmissionsShouldBeSwept(uint24 r) public rand(r) {}
}
