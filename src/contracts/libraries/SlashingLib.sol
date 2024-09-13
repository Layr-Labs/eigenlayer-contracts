// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

library SlashingLib {
    /// @dev The initial total magnitude for an operator
    uint64 public constant INITIAL_TOTAL_MAGNITUDE = 1e18;

    /// @notice that stakerScalingFactor and totalMagnitude have initial default values to 1e18 as "1"
    /// to preserve precision with uint256 math. We use `PRECISION_FACTOR` where these variables are used
    /// and divide to represent as 1
    uint256 public constant PRECISION_FACTOR = 1e18;

    /// @dev Delay before deallocations are completable and can be added back into freeMagnitude
    /// This is also the same delay for withdrawals to be completable
    uint32 public constant DEALLOCATION_DELAY = 17.5 days;

    /**
     * @notice helper pure to calculate the scaledShares to store in queue withdrawal
     * and the shares to decrement from StrategyManager/EigenPodManager
     */
    function calculateSharesToQueueWithdraw(
        uint256 sharesToWithdraw,
        uint256 stakerScalingFactor,
        uint64 totalMagnitude
    ) internal pure returns (uint256 sharesToDecrement, uint256 scaledStakerShares) {
        // TODO: DOUBLE CHECK THIS BEHAVIOR
        // NOTE that to prevent numerator overflow, the max sharesToWithdraw is
        // x*1e36 <= 2^256-1
        // => x <= 1.1579e41
        // however we know that the max shares in a strategy is 1e38-1 and all the ETH on the beaconchain staked
        // in one eigenPod would be ~3e25. This product here is safe
        sharesToDecrement =
            (sharesToWithdraw * PRECISION_FACTOR / stakerScalingFactor) * PRECISION_FACTOR / totalMagnitude;
        scaledStakerShares = sharesToWithdraw * PRECISION_FACTOR / totalMagnitude;
        return (sharesToDecrement, scaledStakerShares);
    }

    /**
     * @notice helper pure to calculate the shares to complete a withdrawal after slashing
     * We use the totalMagnitude of the delegated operator at the time of withdrawal completion to scale the shares
     * back to real StrategyManager/EigenPodManager shares
     */
    function calculateSharesToCompleteWithdraw(
        uint256 scaledStakerShares,
        uint64 totalMagnitudeAtCompletion
    ) internal pure returns (uint256 shares) {
        shares = scaledStakerShares * totalMagnitudeAtCompletion / PRECISION_FACTOR;
    }

    /**
     * @notice helper to calculate the withdrawable shares for a staker given their shares and the
     * current totalMagnitude. `shares` should be the staker's shares in storage in the StrategyManager/EigenPodManager.
     */
    function getWithdrawableShares(
        address staker,
        uint256 stakerScalingFactor,
        uint256 shares,
        uint64 currTotalMagnitude
    ) internal pure returns (uint256) {
        return (stakerScalingFactor * shares / PRECISION_FACTOR) * currTotalMagnitude / PRECISION_FACTOR;
    }

    /**
     * @notice helper pure to return scaledShares given shares and current totalMagnitude. Used for
     * adding/removing staker shares from operatorScaledShares
     */
    function scaleShares(uint256 shares, uint64 totalMagnitude) internal pure returns (uint256) {
        return shares * PRECISION_FACTOR / totalMagnitude;
    }

    /**
     * @notice helper pure to return real strategy shares given operator scaledShares and current totalMagnitude.
     * Used for returning the total delegated shares for an operator and strategy
     */
    function descaleShares(uint256 scaledShares, uint64 totalMagnitude) internal pure returns (uint256) {
        return scaledShares * totalMagnitude / PRECISION_FACTOR;
    }

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
            newStakerScalingFactor = SlashingLib.PRECISION_FACTOR / (totalMagnitude);
        } else {
            // TODO: DOUBLE CHECK THIS BEHAVIOR AND OVERFLOWS
            // staker scaling factor is initialized to PRECISION_FACTOR(1e18) and totalMagnitude is initialized to INITIAL_TOTAL_MAGNITUDE(1e18)
            // and is monotonically decreasing. You can deduce that the newStakerScalingFactor will never decrease to less than the PRECISION_FACTOR
            // so this won't round to 0.
            newStakerScalingFactor = (
                currStakerScalingFactor * existingShares * totalMagnitude / PRECISION_FACTOR
                    + addedShares * PRECISION_FACTOR
            ) / ((existingShares + addedShares) * totalMagnitude);
        }

        return newStakerScalingFactor;
    }
}
