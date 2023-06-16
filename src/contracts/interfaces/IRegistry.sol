// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./IRegistryCoordinator.sol";

/**
 * @title Minimal interface for a `Registry`-type contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice Functions related to the registration process itself have been intentionally excluded
 * because their function signatures may vary significantly.
 */
interface IRegistry {
    function registryCoordinator() external view returns (IRegistryCoordinator);
}
