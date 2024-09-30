// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/math/Math.sol";

/// @dev the stakerScalingFactor and totalMagnitude have initial default values to 1e18 as "1"
/// to preserve precision with uint256 math. We use `WAD` where these variables are used
/// and divide to represent as 1
uint64 constant WAD = 1e18;

/*
 * There are 3 types of shares:
 *      1. shares
 *          - These can be converted to an amount of tokens given a strategy
 *              - by calling `sharesToUnderlying` on the strategy address (they're already tokens 
 *              in the case of EigenPods)
 *          - These are comparable between operators and stakers.
 *          - These live in the storage of StrategyManager strategies: 
 *              - `totalShares` is the total amount of shares that exist in a strategy
 *      2. delegatedShares
 *          - These can be converted to shares given an operator and a strategy
 *              - by multiplying by the operator's totalMagnitude for the strategy
 *          - These values automatically update their conversion into tokens
 *              - when the operator's total magnitude for the strategy is decreased upon slashing
 *          - These live in the storage of the DelegationManager:
 *              - `delegatedShares` is the total amount of delegatedShares delegated to an operator for a strategy
 *              - `withdrawal.delegatedShares` is the amount of delegatedShares in a withdrawal          
 *      3. ownedShares 
 *          - These can be converted into delegatedShares given a staker and a strategy
 *              - by multiplying by the staker's depositScalingFactor for the strategy
 *          - These values automatically update their conversion into tokens
 *             - when the staker's depositScalingFactor for the strategy is increased upon new deposits
 *             - or when the staker's operator's total magnitude for the strategy is decreased upon slashing
 *          - These represent the total amount of shares the staker would have of a strategy if they were never slashed
 *          - These live in the storage of the StrategyManager/EigenPodManager
 *              - `stakerStrategyShares` in the SM is the staker's shares that have not been queued for withdrawal in a strategy
 *              - `podOwnerShares` in the EPM is the staker's shares that have not been queued for withdrawal in the beaconChainETHStrategy
 */

type OwnedShares is uint256;
type DelegatedShares is uint256;
type Shares is uint256;

using SlashingLib for OwnedShares global;
using SlashingLib for DelegatedShares global;
using SlashingLib for Shares global;

library SlashingLib {
    using Math for uint256;
    using SlashingLib for uint256;

    function toShares(
        DelegatedShares delegatedShares,
        uint256 depositScalingFactor
    ) internal pure returns (Shares) {
        if (depositScalingFactor == 0) {
            depositScalingFactor = WAD;
        }

        // forgefmt: disable-next-item
        return delegatedShares
            .unwrap()
            .divWad(depositScalingFactor)
            .wrapShares();
    }

    function toDelegatedShares(
        OwnedShares ownedShares, 
        uint256 magnitude
    ) internal pure returns (DelegatedShares) {
        // forgefmt: disable-next-item
        return ownedShares
            .unwrap()
            .mulWad(magnitude)
            .wrapDelegated();
    }

    function toDelegatedShares(
        Shares shares,
        uint256 magnitude
    ) internal pure returns (DelegatedShares) {
        // forgefmt: disable-next-item
        return shares
            .unwrap()
            .mulWad(magnitude)
            .wrapDelegated();
    }

    function toOwnedShares(
        DelegatedShares delegatedShares, 
        uint256 magnitude
    ) internal pure returns (OwnedShares) {
        // forgefmt: disable-next-item
        return delegatedShares
            .unwrap()
            .mulWad(magnitude)
            .wrapOwned();
    }

    function toOwnedShares(
        Shares shares,
        uint256 stakerScalingFactor,
        uint256 operatorMagnitude
    ) internal pure returns (OwnedShares) {
        /**
         * ownedShares = shares * stakerScalingFactor * operatorMagnitude
         * stakerScalingFactor = storedScalingFactor / WAD
         * operatorMagnitude = storedOperatorMagnitude / WAD
         * ownedShares = shares * (storedScalingFactor / WAD) * (storedOperatorMagnitude / WAD)
         */
        return shares
            .unwrap()
            .mulWad(stakerScalingFactor)
            .mulWad(operatorMagnitude)
            .wrapOwned();
    }

    function calculateNewDepositScalingFactor(
        Shares currentShares,
        Shares addedShares,
        uint256 stakerScalingFactor,
        uint256 operatorMagnitude
    ) internal pure returns (uint256) {
        /**
         * Base Equations:
         * (1) newShares = currentShares + addedShares
         * (2) newOwnedShares = currentOwnedShares + addedShares
         * (3) newOwnedShares = newShares * newStakerScalingFactor * newOperatorMagnitude
         * 
         * Plugging (3) into (2):
         * (4) newShares * newStakerScalingFactor * newOperatorMagnitude = currentOwnedShares + addedShares
         * 
         * Solving for newStakerScalingFactor
         * (5) newStakerScalingFactor = (ownedShares + addedShares) / (newShares * newOperatorMagnitude)
         * 
         * We also know that the operatorMagnitude remains constant on deposits, thus
         * (6) newOperatorMagnitude = oldOperatorMagnitude
         * 
         * Plugging in (6) and (1) into (5):
         * (7) newStakerScalingFactor = (ownedShares + addedShares) / ((currentShares + addedShares) * oldOperatorMagnitude)
         * Note that magnitudes must be divided by WAD for precision. Thus,
         * 
         * (8) newStakerScalingFactor = (ownedShares + addedShares) / ((currentShares + addedShares) * oldOperatorMagnitude / WAD)
         * (9) newStakerScalingFactor = (ownedShares + addedShares) / ((currentShares + addedShares) / WAD) * (oldOperatorMagnitude / WAD))
         */

        // Step 1: Calculate Numerator
        uint256 currentOwnedShares = currentShares.toOwnedShares(stakerScalingFactor, operatorMagnitude).unwrap();

        // Step 2: Compute currentShares + addedShares
        uint256 ownedPlusAddedShares = currentOwnedShares + addedShares.unwrap();

        // Step 3: Calculate newStakerScalingFactor
        // Note: We divide by operatorMagnitude to preserve 

        //TODO: figure out if we only need to do one divWad here
        uint256 newStakerScalingFactor = 
            ownedPlusAddedShares
            .divWad(currentShares.unwrap() + addedShares.unwrap())
            .divWad(operatorMagnitude);

        return newStakerScalingFactor;
    }

    // MATH

    function add(Shares x, uint256 y) internal pure returns (uint256) {
        return x.unwrap() + y;
    }

    function add(Shares x, Shares y) internal pure returns (Shares) {
        return (x.unwrap() + y.unwrap()).wrapShares();
    }

    function add(DelegatedShares x, uint256 y) internal pure returns (uint256) {
        return x.unwrap() + y;
    }

    function add(DelegatedShares x, DelegatedShares y) internal pure returns (DelegatedShares) {
        return (x.unwrap() + y.unwrap()).wrapDelegated();
    }

    function add(OwnedShares x, uint256 y) internal pure returns (uint256) {
        return x.unwrap() + y;
    }

    function add(OwnedShares x, OwnedShares y) internal pure returns (OwnedShares) {
        return (x.unwrap() + y.unwrap()).wrapOwned();
    }

    function sub(Shares x, uint256 y) internal pure returns (uint256) {
        return x.unwrap() - y;
    }

    function sub(Shares x, Shares y) internal pure returns (Shares) {
        return (x.unwrap() - y.unwrap()).wrapShares();
    }

    function sub(DelegatedShares x, uint256 y) internal pure returns (uint256) {
        return x.unwrap() - y;
    }

    function sub(DelegatedShares x, DelegatedShares y) internal pure returns (DelegatedShares) {
        return (x.unwrap() - y.unwrap()).wrapDelegated();
    }

    function sub(OwnedShares x, uint256 y) internal pure returns (uint256) {
        return x.unwrap() - y;
    }

    // WAD MATH

    function mulWad(uint256 x, uint256 y) internal pure returns (uint256) {
        return x.mulDiv(y, WAD);
    }

    function divWad(uint256 x, uint256 y) internal pure returns (uint256) {
        return x.mulDiv(WAD, y);
    }

    // COMPARISONS

    function isZero(Shares x) internal pure returns (bool) {
        return x.unwrap() == 0;
    }

    function isZero(OwnedShares x) internal pure returns (bool) {
        return x.unwrap() == 0;
    }

    function lte(Shares x, Shares y) internal pure returns (bool) {
        return x.unwrap() <= y.unwrap();
    }

    // TYPE CASTING

    function unwrap(Shares x) internal pure returns (uint256) {
        return Shares.unwrap(x);
    }

    function unwrap(DelegatedShares x) internal pure returns (uint256) {
        return DelegatedShares.unwrap(x);
    }

    function unwrap(OwnedShares x) internal pure returns (uint256) {
        return OwnedShares.unwrap(x);
    }

    function wrapShares(uint256 x) internal pure returns (Shares) {
        return Shares.wrap(x);
    }

    function wrapDelegated(uint256 x) internal pure returns (DelegatedShares) {
        return DelegatedShares.wrap(x);
    }

    function wrapOwned(uint256 x) internal pure returns (OwnedShares) {
        return OwnedShares.wrap(x);
    }
}
