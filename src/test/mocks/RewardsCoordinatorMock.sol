// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IRewardsCoordinator.sol";

contract RewardsCoordinatorMock is Test {
    function createRewardsForAllEarners(IRewardsCoordinatorTypes.RewardsSubmission[] calldata) external {
        // Mock implementation - does nothing for testing
    }

    function createUniqueStakeRewardsSubmission(OperatorSet calldata, IRewardsCoordinatorTypes.RewardsSubmission[] calldata) external {
        // Mock implementation - does nothing for testing
    }

    function createTotalStakeRewardsSubmission(OperatorSet calldata, IRewardsCoordinatorTypes.RewardsSubmission[] calldata) external {
        // Mock implementation - does nothing for testing
    }
}

