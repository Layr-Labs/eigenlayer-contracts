// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/utils/ShortStringsUpgradeable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import "../../interfaces/IProtocolRegistry.sol";

abstract contract ProtocolRegistryStorage is IProtocolRegistry {
    ///
    ///                           CONSTANTS
    ///
    /// @inheritdoc IProtocolRegistry
    bytes32 public constant PAUSER_ROLE = hex"01";

    ///
    ///                          MUTABLE STATE
    ///

    /// @notice Returns the semantic version of the protocol.
    ShortString internal _semanticVersion;

    /// @notice Maps deployment name hashes to addresses (enumerable for iteration).
    EnumerableMap.UintToAddressMap internal _deployments;

    /// @notice Maps deployment addresses to their configurations.
    mapping(address => DeploymentConfig) internal _deploymentConfigs;

    /// @dev This empty reserved space is put in place to allow future versions to add new
    /// variables without shifting down storage in the inheritance chain.
    /// See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
    uint256[47] private __gap;
}
