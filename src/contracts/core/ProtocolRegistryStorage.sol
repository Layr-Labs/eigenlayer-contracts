// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/utils/ShortStringsUpgradeable.sol";
import "../interfaces/IProtocolRegistry.sol";

abstract contract ProtocolRegistryStorage is IProtocolRegistry {
    /// @notice Mapping from an address to its semantic version.
    mapping(address addr => ShortString semver) internal _semver;

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}
