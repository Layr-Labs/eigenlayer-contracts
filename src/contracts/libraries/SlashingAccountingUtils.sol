// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

/**
 * @title Library of utilities for calculations related to slashing in EigenLayer
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
library SlashingAccountingUtils {

    // TODO: set this based on payments epochs
    uint256 internal constant EPOCH_GENESIS_TIMESTAMP = 1606824000;

    // TODO: define appropriately
    uint256 internal constant EPOCH_LENGTH_SECONDS = 2 weeks;

    uint256 internal constant SHARE_CONVERSION_SCALE = 1e18;

    // an amount of shares over this will cause overflow when multiplying by `SHARE_CONVERSION_SCALE`
    uint256 internal constant MAX_VALID_SHARES = type(uint96).max;

    uint256 internal constant BIPS_FACTOR = 10000;

    // TODO: explain this better. basically seems like we may need to set some max factor beyond which shares are just zeroed out
    uint256 internal constant MAX_SCALING_FACTOR = type(uint256).max / (MAX_VALID_SHARES * SHARE_CONVERSION_SCALE);

    function getEpochFromTimestamp(uint256 timestamp) internal pure returns (int256) {
        return (int256(timestamp) - int256(EPOCH_GENESIS_TIMESTAMP)) / int256(EPOCH_LENGTH_SECONDS);
    }

    function currentEpoch() internal view returns (int256) {
        return getEpochFromTimestamp(block.timestamp);
    }

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

    function findNewScalingFactor(uint256 scalingFactorBefore, uint256 bipsToSlash) internal pure returns (uint256) {
        require(bipsToSlash != 0, "cannot slash for 0%");
        require(bipsToSlash < BIPS_FACTOR, "cannot slash more than 99.99% at once");
        uint256 scalingFactorAfter;
        // deal with edge case of operator being slashed repeatedly, inflating scalingFactor to max uint size
        // TODO: figure out more nuanced / appropriate way to handle this 'edge case', e.g. deciding if deposits should be blocked when close to limit
        if (MAX_SCALING_FACTOR / scalingFactorBefore >= bipsToSlash) {
            scalingFactorAfter = type(uint256).max;
        } else {
            scalingFactorAfter = scalingFactorBefore * BIPS_FACTOR / (BIPS_FACTOR - bipsToSlash);
        }
        return scalingFactorAfter;
    }
}