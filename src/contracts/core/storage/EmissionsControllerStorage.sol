// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../interfaces/IEmissionsController.sol";
import "../../interfaces/IRewardsCoordinator.sol";
import "../../interfaces/IEigen.sol";
import "../../libraries/OperatorSetLib.sol";

abstract contract EmissionsControllerStorage is IEmissionsController {
    // Constants

    /// @inheritdoc IEmissionsController
    uint256 public constant MAX_TOTAL_WEIGHT = 10_000;

    // Immutables

    /// @dev The EIGEN token that will be minted for emissions.
    IEigen public immutable override EIGEN;
    /// @dev The RewardsCoordinator contract for submitting rewards.
    IRewardsCoordinator public immutable override REWARDS_COORDINATOR;

    /// @inheritdoc IEmissionsController
    uint256 public immutable EMISSIONS_INFLATION_RATE;
    /// @inheritdoc IEmissionsController
    uint256 public immutable EMISSIONS_START_TIME;
    /// @inheritdoc IEmissionsController
    uint256 public immutable EMISSIONS_EPOCH_LENGTH;

    // Mutatables

    /// @inheritdoc IEmissionsController
    address public incentiveCouncil;
    /// @inheritdoc IEmissionsController
    uint256 public totalWeight;

    /// @dev Returns an append-only array of distributions.
    Distribution[] internal _distributions;
    /// @dev Mapping from epoch to the total number of distributions processed for that epoch.
    mapping(uint256 epoch => uint256 totalProcessed) internal _totalProcessed;

    // Construction

    constructor(
        IEigen eigen,
        IRewardsCoordinator rewardsCoordinator,
        uint256 inflationRate,
        uint256 startTime,
        uint256 epochLength
    ) {
        EIGEN = eigen;
        REWARDS_COORDINATOR = rewardsCoordinator;

        EMISSIONS_INFLATION_RATE = inflationRate;
        EMISSIONS_START_TIME = startTime;
        EMISSIONS_EPOCH_LENGTH = epochLength;
    }

    /// @dev This empty reserved space is put in place to allow future versions to add new
    /// variables without shifting down storage in the inheritance chain.
    /// See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
    uint256[45] private __gap;
}
