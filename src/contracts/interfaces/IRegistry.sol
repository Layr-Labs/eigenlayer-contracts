// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

/**
 * @title Minimal interface for a `Registry`-type contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice Functions related to the registration process itself have been intentionally excluded
 * because their function signatures may vary significantly.
 */
interface IRegistry {
    /// @notice Returns 'true' if `operator` is registered as an active operator, and 'false' otherwise.
    function isActiveOperator(address operator) external view returns (bool);
}
