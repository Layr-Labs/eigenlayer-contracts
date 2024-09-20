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
 *      2. stakeShares
 *          - These can be converted to shares given an operator and a strategy
 *              - by multiplying by the operator's totalMagnitude for the strategy
 *          - These values automatically update their conversion into tokens
 *              - when the operator's total magnitude for the strategy is decreased upon slashing
 *          - These live in the storage of the DelegationManager:
 *              - `stakeShares` is the total amount of stakeShares delegated to an operator for a strategy
 *              - `withdrawal.stakeShares` is the amount of stakeShares in a withdrawal          
 *      3. depositShares 
 *          - These can be converted into stakeShares given a staker and a strategy
 *              - by multiplying by the staker's depositScalingFactor for the strategy
 *          - These values automatically update their conversion into tokens
 *             - when the staker's depositScalingFactor for the strategy is increased upon new deposits
 *             - or when the staker's operator's total magnitude for the strategy is decreased upon slashing
 *          - These represent the total amount of shares the staker would have of a strategy if they were never slashed
 *          - These live in the storage of the StrategyManager/EigenPodManager
 *              - `stakerStrategyShares` in the SM is the staker's depositShares that have not been queued for withdrawal in a strategy
 *              - `podOwnerShares` in the EPM is the staker's depositShares that have not been queued for withdrawal in the beaconChainETHStrategy
 */

type Shares is uint256;

type DelegationManagerShares is uint256;

type StakeManagerShares is uint256;

using SlashingLib for Shares global;
using SlashingLib for DelegationManagerShares global;
using SlashingLib for StakeManagerShares global;

library SlashingLib {
    using Math for uint256;
    using SlashingLib for uint256;

    function toDepositShares(
        DelegationManagerShares stakeShares,
        uint256 depositScalingFactor
    ) internal pure returns (StakeManagerShares) {
        if (depositScalingFactor == 0) {
            depositScalingFactor = WAD;
        }
        return StakeManagerShares.wrap(DelegationManagerShares.unwrap(stakeShares).divWad(depositScalingFactor));
    }

    function toStakeShares(
        StakeManagerShares depositShares,
        uint256 depositScalingFactor
    ) internal pure returns (DelegationManagerShares) {
        if (depositScalingFactor == 0) {
            depositScalingFactor = WAD;
        }
        return DelegationManagerShares.wrap(StakeManagerShares.unwrap(depositShares).mulWad(depositScalingFactor));
    }

    function toShares(DelegationManagerShares stakeShares, uint256 magnitude) internal pure returns (Shares) {
        return Shares.wrap(DelegationManagerShares.unwrap(stakeShares).mulWad(magnitude));
    }

    function toStakeShares(Shares shares, uint256 magnitude) internal pure returns (DelegationManagerShares) {
        return DelegationManagerShares.wrap(Shares.unwrap(shares).divWad(magnitude));
    }

    // WAD MATH

    function mulWad(uint256 a, uint256 b) internal pure returns (uint256) {
        return a.mulDiv(b, WAD);
    }

    function divWad(uint256 a, uint256 b) internal pure returns (uint256) {
        return a.mulDiv(WAD, b);
    }
}
