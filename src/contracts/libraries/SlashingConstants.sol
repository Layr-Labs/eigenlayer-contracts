// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

/// @dev the stakerScalingFactor and totalMagnitude have initial default values to 1e18 as "1"
/// to preserve precision with uint256 math. We use `PRECISION_FACTOR` where these variables are used
/// and divide to represent as 1
uint64 constant PRECISION_FACTOR = 1e18;

/// @dev Delay before deallocations are completable and can be added back into freeMagnitude
/// This is also the same delay for withdrawals to be completable
uint32 constant DEALLOCATION_DELAY = 17.5 days;
