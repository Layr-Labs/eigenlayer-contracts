// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./IServiceManager.sol";

/**
 * @title Interface for a `ServiceManager`-type contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
interface IServiceManagerWithRegistryCoordinator is IServiceManager {
    function registryCoordinator() external view returns (IRegistryCoordinator);
}
