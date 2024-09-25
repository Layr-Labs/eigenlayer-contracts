// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/math/Math.sol";

/// @dev the stakerScalingFactor and totalMagnitude have initial default values to 1e18 as "1"
/// to preserve precision with uint256 math. We use `WAD` where these variables are used
/// and divide to represent as 1
uint64 constant WAD = 1e18;

/*
 * There are 3 types of shares:
 *      1. ownedShares
 *          - These can be converted to an amount of tokens given a strategy
 *              - by calling `sharesToUnderlying` on the strategy address (they're already tokens 
 *              in the case of EigenPods)
 *          - These are comparable between operators and stakers.
 *          - These live in the storage of StrategyManager strategies: 
 *              - `totalShares` is the total amount of shares delegated to a strategy
 *      2. delegatedShares
 *          - These can be converted to shares given an operator and a strategy
 *              - by multiplying by the operator's totalMagnitude for the strategy
 *          - These values automatically update their conversion into tokens
 *              - when the operator's total magnitude for the strategy is decreased upon slashing
 *          - These live in the storage of the DelegationManager:
 *              - `delegatedShares` is the total amount of delegatedShares delegated to an operator for a strategy
 *              - `withdrawal.delegatedShares` is the amount of delegatedShares in a withdrawal          
 *      3. shares 
 *          - These can be converted into delegatedShares given a staker and a strategy
 *              - by multiplying by the staker's depositScalingFactor for the strategy
 *          - These values automatically update their conversion into tokens
 *             - when the staker's depositScalingFactor for the strategy is increased upon new deposits
 *             - or when the staker's operator's total magnitude for the strategy is decreased upon slashing
 *          - These represent the total amount of shares the staker would have of a strategy if they were never slashed
 *          - These live in the storage of the StrategyManager/EigenPodManager
 *              - `stakerStrategyShares` in the SM is the staker's shares that have not been queued for withdrawal in a strategy
 *              - `podOwnerShares` in the EPM is the staker's shares that have not been queued for withdrawal in the beaconChainETHStrategy
 * 
 * Note that `withdrawal.delegatedShares` is scaled for the beaconChainETHStrategy to divide by the beaconChainScalingFactor upon queueing
 * and multiply by the beaconChainScalingFactor upon withdrawal
 */

type OwnedShares is uint256;
type DelegatedShares is uint256;
type Shares is uint256;
struct StakerScalingFactors {
    uint256 depositScalingFactor;

    // we need to know if the beaconChainScalingFactor is set because it can be set to 0 through 100% slashing
    bool isBeaconChainScalingFactorSet;
    uint64 beaconChainScalingFactor;
}


using SlashingLib for OwnedShares global;
using SlashingLib for DelegatedShares global;
using SlashingLib for Shares global;
using SlashingLib for StakerScalingFactors global;

// TODO: validate order of operations everywhere
library SlashingLib {
    using Math for uint256;
    using SlashingLib for uint256;

    // MATH

    function add(Shares x, uint256 y) internal pure returns (uint256) {
        return x.unwrap() + y;
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

    function sub(DelegatedShares x, uint256 y) internal pure returns (uint256) {
        return x.unwrap() - y;
    }

    function sub(DelegatedShares x, DelegatedShares y) internal pure returns (DelegatedShares) {
        return (x.unwrap() - y.unwrap()).wrapDelegated();
    }

    function sub(OwnedShares x, uint256 y) internal pure returns (uint256) {
        return x.unwrap() - y;
    }

    /// @dev beaconChainScalingFactor = 0 -> WAD for all non beaconChainETH strategies
    function toShares(
        DelegatedShares delegatedShares,
        StakerScalingFactors storage ssf
    ) internal view returns (Shares) {
        return delegatedShares.unwrap()
                .divWad(ssf.getDepositScalingFactor())
                .divWad(ssf.getBeaconChainScalingFactor())
                .wrapShares();
    }

    function toDelegatedShares(
        OwnedShares shares, 
        uint256 magnitude
    ) internal pure returns (DelegatedShares) {
        // forgefmt: disable-next-item
        return shares
            .unwrap()
            .divWad(magnitude)
            .wrapDelegated();
    }

    function toOwnedShares(DelegatedShares delegatedShares, uint256 magnitude) internal view returns (OwnedShares) {
        return delegatedShares
                .unwrap()
                .mulWad(magnitude)
                .wrapOwned();
    }

    function scaleForQueueWithdrawal(DelegatedShares delegatedShares, StakerScalingFactors storage ssf) internal view returns (DelegatedShares) {
        return delegatedShares
                .unwrap()
                .divWad(ssf.getBeaconChainScalingFactor())
                .wrapDelegated();
    }

    function scaleForCompleteWithdrawal(DelegatedShares delegatedShares, StakerScalingFactors storage ssf) internal view returns (DelegatedShares) {
        return delegatedShares
                .unwrap()
                .mulWad(ssf.getBeaconChainScalingFactor())
                .wrapDelegated();
    }

    function decreaseBeaconChainScalingFactor(StakerScalingFactors storage ssf, uint64 proportionOfOldBalance) internal {
        ssf.beaconChainScalingFactor = uint64(uint256(ssf.beaconChainScalingFactor).mulWad(proportionOfOldBalance));
        ssf.isBeaconChainScalingFactorSet = true;
    }

    /// @dev beaconChainScalingFactor = 0 -> WAD for all non beaconChainETH strategies
    function toDelegatedShares(
        Shares shares,
        StakerScalingFactors storage ssf
    ) internal view returns (DelegatedShares) {
        return shares.unwrap()
                .mulWad(ssf.getDepositScalingFactor())
                .mulWad(ssf.getBeaconChainScalingFactor())
                .wrapDelegated();
    }

    function toDelegatedShares(Shares shares, uint256 magnitude) internal view returns (DelegatedShares) {
        return shares.unwrap()
                .mulWad(magnitude)
                .wrapDelegated();
    }

    // WAD MATH

    function mulWad(uint256 x, uint256 y) internal pure returns (uint256) {
        return x.mulDiv(y, WAD);
    }

    function divWad(uint256 x, uint256 y) internal pure returns (uint256) {
        return x.mulDiv(WAD, y);
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

    function getDepositScalingFactor(StakerScalingFactors storage ssf) internal view returns (uint256) {
        return ssf.depositScalingFactor == 0 ? WAD : ssf.depositScalingFactor;
    }
    
    function getBeaconChainScalingFactor(StakerScalingFactors storage ssf) internal view returns (uint64) {
        return !ssf.isBeaconChainScalingFactorSet && ssf.beaconChainScalingFactor == 0 ? WAD : ssf.beaconChainScalingFactor;
    }
}
