// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../interfaces/IEmissionsController.sol";

abstract contract EmissionsControllerStorage is IEmissionsController {
    // Immutables

    /// @inheritdoc IEmissionsController
    uint256 public immutable EMISSIONS_INFLATION_RATE;
    /// @inheritdoc IEmissionsController
    uint256 public immutable EMISSIONS_START_TIME;
    /// @inheritdoc IEmissionsController
    uint256 public immutable EMISSIONS_EPOCH_LENGTH;

    // Mutatables

    /// @inheritdoc IEmissionsController
    address public incentiveCouncil;

    /// @dev Returns an append-only array of distributions.
    Distribution[] internal _distributions;
    /// @dev Mapping from epoch to whether it has been triggered.
    mapping(uint256 epoch => bool triggered) internal _epochTriggered;

    // Construction

    constructor(
        uint256 inflationRate,
        uint256 startTime,
        uint256 cooldownSeconds
    ) {
        EMISSIONS_INFLATION_RATE = inflationRate;
        EMISSIONS_START_TIME = startTime;
        EMISSIONS_EPOCH_LENGTH = cooldownSeconds;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[46] private __gap;
}
