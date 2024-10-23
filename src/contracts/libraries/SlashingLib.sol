// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/math/Math.sol";

/// @dev the stakerScalingFactor and operatorMagnitude have initial default values to 1e18 as "1"
/// to preserve precision with uint256 math. We use `WAD` where these variables are used
/// and divide to represent as 1
uint64 constant WAD = 1e18;

/*
 * There are 2 types of shares:
 *      1. depositShares
 *          - These can be converted to an amount of tokens given a strategy
 *              - by calling `sharesToUnderlying` on the strategy address (they're already tokens 
 *              in the case of EigenPods)
 *          - These live in the storage of EPM and SM strategies 
 *      2. shares
 *          - For a staker, this is the amount of shares that they can withdraw
 *          - For an operator, this is the sum of its staker's withdrawable shares       
 * 
 * Note that `withdrawal.scaledShares` is scaled for the beaconChainETHStrategy to divide by the beaconChainScalingFactor upon queueing
 * and multiply by the beaconChainScalingFactor upon withdrawal
 */

struct StakerScalingFactors {
    uint256 depositScalingFactor;
    // we need to know if the beaconChainScalingFactor is set because it can be set to 0 through 100% slashing
    bool isBeaconChainScalingFactorSet;
    uint64 beaconChainScalingFactor;
}

using SlashingLib for StakerScalingFactors global;

// TODO: validate order of operations everywhere
library SlashingLib {
    using Math for uint256;
    using SlashingLib for uint256;

    // WAD MATH

    function mulWad(uint256 x, uint256 y) internal pure returns (uint256) {
        return x.mulDiv(y, WAD);
    }

    function divWad(uint256 x, uint256 y) internal pure returns (uint256) {
        return x.mulDiv(WAD, y);
    }

    /**
     * @notice Used explicitly for calculating slashed magnitude, we want to ensure even in the
     * situation where an operator is slashed several times and precision has been lost over time,
     * an incoming slashing request isn't rounded down to 0 and an operator is able to avoid slashing penalties.
     */
    function mulWadRoundUp(uint256 x, uint256 y) internal pure returns (uint256) {
        return x.mulDiv(y, WAD, Math.Rounding.Up);
    }

    // GETTERS

    function getDepositScalingFactor(
        StakerScalingFactors memory ssf
    ) internal pure returns (uint256) {
        return ssf.depositScalingFactor == 0 ? WAD : ssf.depositScalingFactor;
    }

    function getBeaconChainScalingFactor(
        StakerScalingFactors memory ssf
    ) internal pure returns (uint64) {
        return ssf.isBeaconChainScalingFactorSet ? ssf.beaconChainScalingFactor : WAD;
    }

    function scaleSharesForQueuedWithdrawal(
        uint256 sharesToWithdraw,
        StakerScalingFactors memory ssf,
        uint64 operatorMagnitude
    ) internal pure returns (uint256) {
        /// forgefmt: disable-next-item
        return sharesToWithdraw
            .divWad(uint256(ssf.getBeaconChainScalingFactor()))
            .divWad(uint256(operatorMagnitude));
    }

    function scaleSharesForCompleteWithdrawal(
        uint256 scaledShares,
        StakerScalingFactors memory ssf,
        uint64 operatorMagnitude
    ) internal pure returns (uint256) {
        /// forgefmt: disable-next-item
        return scaledShares
            .mulWad(uint256(ssf.getBeaconChainScalingFactor()))
            .mulWad(uint256(operatorMagnitude));
    }

    function getOperatorSharesToDecrease(
        uint256 operatorShares,
        uint64 previousMaxMagnitude,
        uint64 newMaxMagnitude
    ) internal pure returns (uint256) {
        return operatorShares - operatorShares.divWad(previousMaxMagnitude).mulWad(newMaxMagnitude);
    }

    function decreaseBeaconChainScalingFactor(
        StakerScalingFactors storage ssf,
        uint64 proportionOfOldBalance
    ) internal {
        ssf.beaconChainScalingFactor = uint64(uint256(ssf.getBeaconChainScalingFactor()).mulWad(proportionOfOldBalance));
        ssf.isBeaconChainScalingFactorSet = true;
    }

    function updateDepositScalingFactor(
        StakerScalingFactors storage ssf,
        uint256 existingDepositShares,
        uint256 addedShares,
        uint64 maxMagnitude
    ) internal {
        if (existingDepositShares == 0) {
            // if this is their first deposit for the operator, set the scaling factor to inverse of maxMagnitude
            /// forgefmt: disable-next-item
            ssf.depositScalingFactor = uint256(WAD)
                .divWad(ssf.getBeaconChainScalingFactor())
                .divWad(maxMagnitude);
            return;
        }
        /**
         * Base Equations:
         * (1) newShares = currentShares + addedShares
         * (2) newDepositShares = existingDepositShares + addedShares
         * (3) newShares = newDepositShares * newStakerDepositScalingFactor * beaconChainScalingFactor * maxMagnitude
         *
         * Plugging (1) into (3):
         * (4) newDepositShares * newStakerDepositScalingFactor * beaconChainScalingFactor * maxMagnitude = currentShares + addedShares
         *
         * Solving for newStakerDepositScalingFactor
         * (5) newStakerDepositScalingFactor = (currentShares + addedShares) / (newDepositShares * beaconChainScalingFactor * maxMagnitude)
         *
         * Plugging in (2) into (5):
         * (7) newStakerDepositScalingFactor = (currentShares + addedShares) / ((existingDepositShares + addedShares) * beaconChainScalingFactor * maxMagnitude)
         * Note that magnitudes must be divided by WAD for precision. Thus,
         *
         * (8) newStakerDepositScalingFactor = WAD * (currentShares + addedShares) / ((existingDepositShares + addedShares) * beaconChainScalingFactor / WAD * maxMagnitude / WAD)
         * (9) newStakerDepositScalingFactor = (currentShares + addedShares) * WAD / (existingDepositShares + addedShares) * WAD / beaconChainScalingFactor * WAD / maxMagnitude
         */

        // Step 1: Calculate Numerator
        uint256 currentShares = existingDepositShares.toShares(ssf, maxMagnitude);

        // Step 2: Compute currentShares + addedShares
        uint256 newShares = currentShares + addedShares;

        // Step 3: Calculate newStakerDepositScalingFactor
        /// forgefmt: disable-next-item
        uint256 newStakerDepositScalingFactor = newShares
            .divWad(existingDepositShares + addedShares)
            .divWad(maxMagnitude)
            .divWad(uint256(ssf.getBeaconChainScalingFactor()));

        ssf.depositScalingFactor = newStakerDepositScalingFactor;
    }

    // CONVERSION

    function toDepositShares(
        uint256 shares,
        StakerScalingFactors memory ssf,
        uint64 magnitude
    ) internal pure returns (uint256 depositShares) {
        /// forgefmt: disable-next-item
        depositShares = shares
            .divWad(ssf.getDepositScalingFactor())
            .divWad(uint256(ssf.getBeaconChainScalingFactor()))
            .divWad(uint256(magnitude));
    }

    function toShares(
        uint256 depositShares,
        StakerScalingFactors memory ssf,
        uint64 magnitude
    ) internal pure returns (uint256 shares) {
        /// forgefmt: disable-next-item
        shares = depositShares
            .mulWad(ssf.getDepositScalingFactor())
            .mulWad(uint256(ssf.getBeaconChainScalingFactor()))
            .mulWad(uint256(magnitude));
    }
}
