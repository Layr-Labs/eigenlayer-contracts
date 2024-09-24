// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/math/Math.sol";

/// @dev the stakerScalingFactor and totalMagnitude have initial default values to 1e18 as "1"
/// to preserve precision with uint256 math. We use `WAD` where these variables are used
/// and divide to represent as 1
uint64 constant WAD = 1e18;

/// @dev Delay before deallocations are completable and can be added back into freeMagnitude
/// This is also the same delay for withdrawals to be completable
uint32 constant DEALLOCATION_DELAY = 17.5 days;

/*
 * There are 3 types of shares:
 *      1. shares
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
 *      3. principalShares 
 *          - These can be converted into delegatedShares given a staker and a strategy
 *              - by multiplying by the staker's depositScalingFactor for the strategy
 *          - These values automatically update their conversion into tokens
 *             - when the staker's depositScalingFactor for the strategy is increased upon new deposits
 *             - or when the staker's operator's total magnitude for the strategy is decreased upon slashing
 *          - These represent the total amount of shares the staker would have of a strategy if they were never slashed
 *          - These live in the storage of the StrategyManager/EigenPodManager
 *              - `stakerStrategyShares` in the SM is the staker's principalShares that have not been queued for withdrawal in a strategy
 *              - `podOwnerShares` in the EPM is the staker's principalShares that have not been queued for withdrawal in the beaconChainETHStrategy
 */

type WithdrawableShares is uint256;
type DelegatedShares is uint256;
type Shares is uint256;

using SlashingLib for WithdrawableShares global;
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
        Shares principalShares,
        uint256 depositScalingFactor
    ) internal pure returns (DelegatedShares) {
        if (depositScalingFactor == 0) {
            depositScalingFactor = WAD;
        }

        // forgefmt: disable-next-item
        return principalShares
            .unwrap()
            .mulWad(depositScalingFactor)
            .wrapDelegated();
    }

    function toWithdrawableShares(
        DelegatedShares delegatedShares, 
        uint256 magnitude
    ) internal pure returns (WithdrawableShares) {
        // forgefmt: disable-next-item
        return delegatedShares
            .unwrap()
            .mulWad(magnitude)
            .wrapWithdrawable();
    }

    function toDelegatedShares(
        WithdrawableShares shares, 
        uint256 magnitude
    ) internal pure returns (DelegatedShares) {
        // forgefmt: disable-next-item
        return shares
            .unwrap()
            .divWad(magnitude)
            .wrapDelegated();
    }

    // MATH

    function add(Shares x, uint256 y) internal pure returns (uint256) {
        return x.unwrap() + y;
    }

    function add(DelegatedShares x, uint256 y) internal pure returns (uint256) {
        return x.unwrap() + y;
    }

    function add(WithdrawableShares x, uint256 y) internal pure returns (uint256) {
        return x.unwrap() + y;
    }

    function sub(Shares x, uint256 y) internal pure returns (uint256) {
        return x.unwrap() - y;
    }

    function sub(DelegatedShares x, uint256 y) internal pure returns (uint256) {
        return x.unwrap() - y;
    }

    function sub(WithdrawableShares x, uint256 y) internal pure returns (uint256) {
        return x.unwrap() - y;
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

    function unwrap(WithdrawableShares x) internal pure returns (uint256) {
        return WithdrawableShares.unwrap(x);
    }

    function wrapShares(uint256 x) internal pure returns (Shares) {
        return Shares.wrap(x);
    }

    function wrapDelegated(uint256 x) internal pure returns (DelegatedShares) {
        return DelegatedShares.wrap(x);
    }

    function wrapWithdrawable(uint256 x) internal pure returns (WithdrawableShares) {
        return WithdrawableShares.wrap(x);
    }
}
