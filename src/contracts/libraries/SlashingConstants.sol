// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";

library SlashingConstants {
    /// @dev The initial total magnitude for an operator
    uint64 public constant INITIAL_TOTAL_MAGNITUDE = 1e18;

    /// @notice that stakerScalingFactor and totalMagnitude have initial default values to 1e18 as "1"
    /// to preserve precision with uint256 math. We use `PRECISION_FACTOR` where these variables are used
    /// and divide to represent as 1
    uint256 public constant PRECISION_FACTOR = 1e18;
    uint256 public constant PRECISION_FACTOR_SQUARED = 1e36;

    /// @dev Delay before deallocations are completable and can be added back into freeMagnitude
    uint32 public constant DEALLOCATION_DELAY = 17.5 days;
}
