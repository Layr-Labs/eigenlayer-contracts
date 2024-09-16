// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/math/Math.sol";

library SlashingLib {
    using SlashingLib for uint256;

    /// @dev The initial total magnitude for an operator
    uint64 public constant INITIAL_TOTAL_MAGNITUDE = 1e18;

    /// @dev that stakerScalingFactor and totalMagnitude have initial default values to 1e18 as "1"
    /// to preserve precision with uint256 math. We use `PRECISION_FACTOR` where these variables are used
    /// and divide to represent as 1
    uint256 public constant PRECISION_FACTOR = 1e18;

    /// @dev Delay before deallocations are completable and can be added back into freeMagnitude
    /// This is also the same delay for withdrawals to be completable
    uint32 public constant DEALLOCATION_DELAY = 17.5 days;

    /**
     * @dev Helper pure function to calculate the scaledShares to store in queue withdrawal
     * and the shares to decrement from StrategyManager/EigenPodManager
     */
    function calculateSharesToQueueWithdraw(
        uint256 sharesToWithdraw,
        uint256 stakerScalingFactor,
        uint64 totalMagnitude
    ) internal pure returns (uint256 sharesToDecrement, uint256 scaledStakerShares) {
        // TODO: DOUBLE CHECK THIS BEHAVIOR
        sharesToDecrement = sharesToWithdraw.divWad(stakerScalingFactor).divWad(totalMagnitude);
        scaledStakerShares = sharesToWithdraw.divWad(totalMagnitude);
    }

    /**
     * @dev helper pure to calculate the shares to complete a withdrawal after slashing
     * We use the totalMagnitude of the delegated operator at the time of withdrawal completion to scale the shares
     * back to real StrategyManager/EigenPodManager shares
     */
    function calculateSharesToCompleteWithdraw(
        uint256 scaledStakerShares,
        uint64 totalMagnitudeAtCompletion
    ) internal pure returns (uint256) {
        return scaledStakerShares.mulWad(totalMagnitudeAtCompletion);
    }

    /**
     * @dev helper to calculate the withdrawable shares for a staker given their shares and the
     * current totalMagnitude. `shares` should be the staker's shares in storage in the StrategyManager/EigenPodManager.
     */
    function getWithdrawableShares(
        address staker,
        uint256 stakerScalingFactor,
        uint256 shares,
        uint64 currTotalMagnitude
    ) internal pure returns (uint256) {
        return stakerScalingFactor.mulWad(shares).mulWad(currTotalMagnitude);
    }

    /**
     * @dev helper pure to return scaledShares given shares and current totalMagnitude. Used for
     * adding/removing staker shares from operatorScaledShares
     */
    function scaleShares(uint256 shares, uint64 totalMagnitude) internal pure returns (uint256) {
        return shares.divWad(totalMagnitude);
    }

    /**
     * @dev helper pure to return real strategy shares given operator scaledShares and current totalMagnitude.
     * Used for returning the total delegated shares for an operator and strategy
     */
    function descaleShares(uint256 scaledShares, uint64 totalMagnitude) internal pure returns (uint256) {
        return scaledShares.mulWad(totalMagnitude);
    }

    /**
     * @dev helper to calculate the new staker scaling factor after adding shares. This is only used
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
        if (existingShares == 0) {
            // existing shares are 0, meaning no existing delegated shares. In this case, the new staker scaling factor
            // is re-initialized to
            return PRECISION_FACTOR / totalMagnitude;
        } else {
            // TODO: DOUBLE CHECK THIS BEHAVIOR AND OVERFLOWS
            // staker scaling factor is initialized to PRECISION_FACTOR(1e18) and totalMagnitude is initialized to INITIAL_TOTAL_MAGNITUDE(1e18)
            // and is monotonically decreasing. You can deduce that the newStakerScalingFactor will never decrease to less than the PRECISION_FACTOR
            // so this won't round to 0.
            return (existingShares.mulWad(currStakerScalingFactor).mulWad(totalMagnitude) + addedShares)
                / (existingShares + addedShares).mulWad(totalMagnitude);
        }
    }

    /// WAD MATH

    /// @dev Multiplys two unsigned intergers that are denominated in wad (18 decimal places).
    /// Intermediate overflow is avoided, read more here: https://xn--2-umb.com/21/muldiv
    function mulWad(uint256 a, uint256 b) internal pure returns (uint256) {
        return Math.mulDiv(a, b, PRECISION_FACTOR);
    }

    /// @dev Divides two unsigned intergers that are denominated in wad (18 decimal places).
    /// Intermediate overflow is avoided, read more here: https://xn--2-umb.com/21/muldiv
    function divWad(uint256 a, uint256 b) internal pure returns (uint256) {
        return Math.mulDiv(a, PRECISION_FACTOR, b);
    }
}
