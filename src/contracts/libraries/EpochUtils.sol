// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

/**
 * @title Library of utilities for calculations related to epochs in EigenLayer
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
library EpochUtils {

    // TODO: set this based on payments epochs
    uint256 internal constant EPOCH_GENESIS_TIMESTAMP = 1606824000;

    // TODO: define appropriately
    uint256 internal constant EPOCH_LENGTH_SECONDS = 1 weeks;

    function getEpochFromTimestamp(uint256 timestamp) internal pure returns (int256) {
        return (int256(timestamp) - int256(EPOCH_GENESIS_TIMESTAMP)) / int256(EPOCH_LENGTH_SECONDS);
    }

    function currentEpoch() internal view returns (int256) {
        return getEpochFromTimestamp(block.timestamp);
    }
}