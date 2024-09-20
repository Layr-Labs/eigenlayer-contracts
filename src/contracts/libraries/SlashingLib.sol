// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

/// @dev the stakerScalingFactor and totalMagnitude have initial default values to 1e18 as "1"
/// to preserve precision with uint256 math. We use `PRECISION_FACTOR` where these variables are used
/// and divide to represent as 1
uint64 constant PRECISION_FACTOR = 1e18;

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
struct StakeShares {
    uint256 stakeShares;
}

struct DepositShares {
    uint256 depositShares;
}

struct Shares {
    uint256 shares;
}

library SlashingLib {
    function createShares(uint256 shares) internal pure returns (Shares memory) {
        return Shares({
            shares: shares
        });
    }

    function toUint256(Shares memory shares) internal pure returns (uint256) {
        return shares.shares;
    }

    function createStakeShares(uint256 stakeShares) internal pure returns (StakeShares memory) {
        return StakeShares({
            stakeShares: stakeShares
        });
    }

    function toUint256(StakeShares memory stakeShares) internal pure returns (uint256) {
        return stakeShares.stakeShares;
    }

    function createDepositShares(uint256 depositShares) internal pure returns (DepositShares memory) {
        return DepositShares({
            depositShares: depositShares
        });
    }

    function toUint256(DepositShares memory depositShares) internal pure returns (uint256) {
        return depositShares.depositShares;
    }

    function toStakeShares(Shares memory shares, uint256 magnitude) internal pure returns (StakeShares memory) {
        return StakeShares({
            stakeShares: shares.shares * PRECISION_FACTOR / magnitude
        });
    }

    function toDepositShares(StakeShares memory stakeShares, uint256 depositScalingFactor) internal pure returns (DepositShares memory) {
        if (depositScalingFactor == 0) {
            depositScalingFactor = PRECISION_FACTOR;
        }
        return DepositShares({
            depositShares: stakeShares.stakeShares * PRECISION_FACTOR / depositScalingFactor
        });
    }

    function toStakeShares(DepositShares memory depositShares, uint256 depositScalingFactor) internal pure returns (StakeShares memory) {
        if (depositScalingFactor == 0) {
            depositScalingFactor = PRECISION_FACTOR;
        }
        return StakeShares({
            stakeShares: depositShares.depositShares * depositScalingFactor / PRECISION_FACTOR
        });
    }

    function toShares(StakeShares memory stakeShares, uint256 magnitude) internal pure returns (Shares memory) {
        return Shares({
            shares: stakeShares.stakeShares * magnitude / PRECISION_FACTOR
        });
    }
}