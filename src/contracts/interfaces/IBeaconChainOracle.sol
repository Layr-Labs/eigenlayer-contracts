// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

/**
 * @title Interface for the BeaconStateOracle contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
interface IBeaconChainOracle {
    /// @notice The block number to state root mapping.
    function timestampToBlockRoot(uint256 timestamp) external view returns (bytes32);
}
