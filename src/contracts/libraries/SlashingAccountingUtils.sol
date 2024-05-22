// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

/**
 * @title Library of utilities for calculations related to slashing in EigenLayer
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
library SlashingAccountingUtils {

    uint256 internal constant SHARE_CONVERSION_SCALE = 1e18;

    // TODO: evaluate the correct max and work on ensuring this is enforced
    // an amount of shares over this will cause overflow when multiplying by `SHARE_CONVERSION_SCALE`
    uint256 internal constant MAX_VALID_SHARES = type(uint96).max;

    uint256 internal constant BIPS_FACTOR = 10000;

    uint256 internal constant BIPS_FACTOR_SQUARED = 1e8;

    // TODO: explain this better. basically seems like we may need to set some max factor beyond which shares are just zeroed out
    // at present this is (2^160)/1e18 ~ 1.46e30, which would be reached after ~659 consecutive 10% slashings, or ~15 consecutive 99% slashings
    uint256 internal constant MAX_SCALING_FACTOR = type(uint256).max / (MAX_VALID_SHARES * SHARE_CONVERSION_SCALE);

    function denormalize(uint256 shares, uint256 scalingFactor) internal pure returns (uint256) {
        return (shares * scalingFactor) / SHARE_CONVERSION_SCALE;
    }

    function normalize(uint256 nonNormalizedShares, uint256 scalingFactor) internal pure returns (uint256) {
        return (nonNormalizedShares * SHARE_CONVERSION_SCALE) / scalingFactor;
    }

    // @notice Overloaded version of `scaleUp` that accepts a signed integer shares amount
    function denormalize(int256 shares, uint256 scalingFactor) internal pure returns (int256) {
        return (shares * int256(scalingFactor)) / int256(SHARE_CONVERSION_SCALE);
    }

    // @notice Overloaded version of `scaleDown` that accepts a signed integer shares amount
    function normalize(int256 nonNormalizedShares, uint256 scalingFactor) internal pure returns (int256) {
        return (nonNormalizedShares * int256(SHARE_CONVERSION_SCALE)) / int256(scalingFactor);
    }

    // TODO: consider possible loss of precision and its consequences
    // @dev note that rateToSlash is in parts per BIPS_FACTOR_SQUARED, i.e. in parts per 1e8
    function findNewScalingFactor(uint256 scalingFactorBefore, uint256 rateToSlash) internal pure returns (uint256) {
        require(rateToSlash != 0, "cannot slash for 0%");
        require(rateToSlash <= BIPS_FACTOR_SQUARED, "cannot slash more than 100% at once");
        uint256 scalingFactorAfter;
        // deal with edge case of operator being slashed repeatedly, inflating scalingFactor to max uint size
        // TODO: figure out more nuanced / appropriate way to handle this 'edge case', e.g. deciding if deposits should be blocked when close to limit
        if (rateToSlash == BIPS_FACTOR_SQUARED || MAX_SCALING_FACTOR / scalingFactorBefore >= rateToSlash) {
            scalingFactorAfter = type(uint256).max;
        } else {
            scalingFactorAfter = scalingFactorBefore * BIPS_FACTOR_SQUARED / (BIPS_FACTOR_SQUARED - rateToSlash);
        }
        return scalingFactorAfter;
    }
}