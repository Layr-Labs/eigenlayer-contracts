// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IRewardsCoordinator.sol";

contract RewardsCoordinatorMock is Test {
    /// @notice The interval in seconds at which the rewards calculation is done
    /// @dev Defaulting to 1 day to match the real RewardsCoordinator configuration
    uint32 public CALCULATION_INTERVAL_SECONDS = 1 days;

    function createRewardsForAllEarners(IRewardsCoordinatorTypes.RewardsSubmission[] calldata) external {
        // Mock implementation - does nothing for testing
    }

    function createUniqueStakeRewardsSubmission(OperatorSet calldata, IRewardsCoordinatorTypes.RewardsSubmission[] calldata) external {
        // Mock implementation - does nothing for testing
    }

    function createTotalStakeRewardsSubmission(OperatorSet calldata, IRewardsCoordinatorTypes.RewardsSubmission[] calldata) external {
        // Mock implementation - does nothing for testing
    }

    function createAVSRewardsSubmission(IRewardsCoordinatorTypes.RewardsSubmission[] calldata) external {
        // Mock implementation - does nothing for testing
    }
}

