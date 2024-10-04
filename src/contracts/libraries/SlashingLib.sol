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
 * Note that `withdrawal.scaledSharesToWithdraw` is scaled for the beaconChainETHStrategy to divide by the beaconChainScalingFactor upon queueing
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

    // GETTERS


    function getDepositScalingFactor(StakerScalingFactors memory ssf) internal pure returns (uint256) {
        return ssf.depositScalingFactor == 0 ? WAD : ssf.depositScalingFactor;
    }
    
    function getBeaconChainScalingFactor(StakerScalingFactors memory ssf) internal pure returns (uint64) {
        return !ssf.isBeaconChainScalingFactorSet && ssf.beaconChainScalingFactor == 0 ? WAD : ssf.beaconChainScalingFactor;
    }


    function scaleSharesForQueuedWithdrawal(uint256 sharesToWithdraw, StakerScalingFactors memory ssf, uint64 operatorMagnitude) internal pure returns (uint256) {
        return 
            sharesToWithdraw
                .divWad(uint256(ssf.getBeaconChainScalingFactor()))
                .divWad(uint256(operatorMagnitude));
    }

    function scaleSharesForCompleteWithdrawal(uint256 scaledSharesToWithdraw, StakerScalingFactors memory ssf, uint64 operatorMagnitude) internal pure returns (uint256) {
        return 
            scaledSharesToWithdraw
                .mulWad(uint256(ssf.getBeaconChainScalingFactor()))
                .mulWad(uint256(operatorMagnitude));
    }

    function decreaseOperatorShares(uint256 operatorShares, uint64 previousMagnitude, uint64 newMagnitude) internal pure returns (uint256) {
        return operatorShares.divWad(previousMagnitude).mulWad(newMagnitude);    
    }

    function decreaseBeaconChainScalingFactor(StakerScalingFactors storage ssf, uint64 proportionOfOldBalance) internal {
        ssf.beaconChainScalingFactor = uint64(uint256(ssf.beaconChainScalingFactor).mulWad(proportionOfOldBalance));
        ssf.isBeaconChainScalingFactorSet = true;
    }

    function calculateNewDepositScalingFactor(
        uint256 currentDepositShares,
        uint256 addedDepositShares,
        StakerScalingFactors memory ssf,
        uint64 magnitude
    ) internal pure returns (uint64) {
        // TODO: update equation for beacon chain scaling factor
        /**
         * Base Equations:
         * (1) newShares = currentShares + addedDepositShares
         * (2) newOwnedShares = currentOwnedShares + addedDepositShares
         * (3) newOwnedShares = newShares * newStakerDepositScalingFactor * newMagnitude
         * 
         * Plugging (3) into (2):
         * (4) newShares * newStakerDepositScalingFactor * newMagnitude = currentOwnedShares + addedDepositShares
         * 
         * Solving for newStakerDepositScalingFactor
         * (5) newStakerDepositScalingFactor = (ownedShares + addedDepositShares) / (newShares * newMagnitude)
         * 
         * We also know that the magnitude remains constant on deposits, thus
         * (6) newMagnitude = oldMagnitude
         * 
         * Plugging in (6) and (1) into (5):
         * (7) newStakerDepositScalingFactor = (ownedShares + addedDepositShares) / ((currentShares + addedDepositShares) * oldMagnitude)
         * Note that magnitudes must be divided by WAD for precision. Thus,
         * 
         * (8) newStakerDepositScalingFactor = (ownedShares + addedDepositShares) / ((currentShares + addedDepositShares) * oldMagnitude / WAD)
         * (9) newStakerDepositScalingFactor = (ownedShares + addedDepositShares) / ((currentShares + addedDepositShares) / WAD) * (oldMagnitude / WAD))
         */

        // Step 1: Calculate Numerator
        uint256 currentShares = currentDepositShares.toShares(ssf, magnitude);

        // Step 2: Compute currentShares + addedShares
        uint256 newShares = currentShares + addedDepositShares;

        // Step 3: Calculate newStakerDepositScalingFactor
        // Note: We divide by magnitude to preserve 

        //TODO: figure out if we only need to do one divWad here
        uint256 newStakerDepositScalingFactor = 
            newShares
            .divWad(currentShares + addedDepositShares)
            .divWad(magnitude)
            .divWad(uint256(ssf.getBeaconChainScalingFactor()));

        return uint64(newStakerDepositScalingFactor);
    }

    // CONVERSION

    function toDepositShares(
        uint256 shares,
        StakerScalingFactors memory ssf,
        uint64 magnitude
    ) internal pure returns (uint256 depositShares) {
        depositShares = 
            shares
            .divWad(ssf.getDepositScalingFactor())
            .divWad(uint256(ssf.getBeaconChainScalingFactor()))
            .divWad(uint256(magnitude));
    }

    function toShares(uint256 depositShares, StakerScalingFactors memory ssf, uint64 magnitude) internal pure returns (uint256 shares) {
        shares = 
            depositShares                
                .mulWad(ssf.getDepositScalingFactor())
                .mulWad(uint256(ssf.getBeaconChainScalingFactor()))
                .mulWad(uint256(magnitude));
    }

}
