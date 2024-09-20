// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

/// @dev stakerScalingFactor and totalMagnitude have initial default values to 1e18 as "1"
/// to preserve precision with uint256 math. We use `PRECISION_FACTOR` where these variables are used
/// and divide to represent as 1
uint64 constant PRECISION_FACTOR = 1e18;

// TODO: make this a immutable in contracts
/// @dev Delay before deallocations are completable and can be added back into freeMagnitude
/// This is also the same delay for withdrawals to be completable
uint32 constant DEALLOCATION_DELAY = 17.5 days;

library SlashingLib {

    /**
     * @notice helper to calculate the new staker scaling factor after adding shares. This is only used
     * when a staker is depositing through the StrategyManager or EigenPodManager. A stakers scaling factor
     * is only updated when they have new deposits and their shares are being increased.
     */
    function calculateStakerScalingFactor(
        address staker,
        uint256 currStakerScalingFactor,
        uint64 totalMagnitude,
        uint256 existingShares,
        uint256 addedShares
    ) internal pure returns (uint256) {
        uint256 newStakerScalingFactor;

        if (existingShares == 0) {
            // existing shares are 0, meaning no existing delegated shares. In this case, the new staker scaling factor
            // is re-initialized to
            newStakerScalingFactor = PRECISION_FACTOR / (totalMagnitude);
        } else {
            // TODO: DOUBLE CHECK THIS BEHAVIOR AND OVERFLOWS
            // staker scaling factor is initialized to PRECISION_FACTOR(1e18) and totalMagnitude is initialized to INITIAL_TOTAL_MAGNITUDE(1e18)
            // and is monotonically decreasing. You can deduce that the newStakerScalingFactor will never decrease to less than the PRECISION_FACTOR
            // so this won't round to 0.
            newStakerScalingFactor = (
                existingShares * currStakerScalingFactor / PRECISION_FACTOR * totalMagnitude / PRECISION_FACTOR
                    + addedShares
            ) / ((existingShares + addedShares) * totalMagnitude / PRECISION_FACTOR);
        }

        return newStakerScalingFactor;
    }
}
