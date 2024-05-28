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

    function getEpochFromTimestamp(uint256 timestamp) internal pure returns (uint32) {
        require(timestamp >= EPOCH_GENESIS_TIMESTAMP, "EpochUtils.getEpochFromTimestamp: timestamp is before genesis");
        return uint32((timestamp - EPOCH_GENESIS_TIMESTAMP) / EPOCH_LENGTH_SECONDS);
    }

    function currentEpoch() internal view returns (uint32) {
        return getEpochFromTimestamp(block.timestamp);
    }
}