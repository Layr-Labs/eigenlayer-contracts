// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "./ProtocolRegistryStorage.sol";

contract ProtocolRegistry is Initializable, OwnableUpgradeable, ProtocolRegistryStorage {
    using ShortStringsUpgradeable for *;

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */
    constructor() {
        _disableInitializers();
    }

    /// @inheritdoc IProtocolRegistry
    function initialize(
        address initialOwner
    ) external initializer {
        _transferOwnership(initialOwner);
    }

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /// @inheritdoc IProtocolRegistry
    function setVersion(address addr, string calldata semver) external onlyOwner {
        _setVersion(addr, semver);
    }

    /// @inheritdoc IProtocolRegistry
    function setVersions(address[] calldata addresses, string calldata semver) external onlyOwner {
        for (uint256 i = 0; i < addresses.length; ++i) {
            _setVersion(addresses[i], semver);
        }
    }

    /// @inheritdoc IProtocolRegistry
    function setVersions(address[] calldata addresses, string[] calldata semvers) external onlyOwner {
        require(addresses.length == semvers.length, InputArrayLengthMismatch());
        for (uint256 i = 0; i < addresses.length; ++i) {
            _setVersion(addresses[i], semvers[i]);
        }
    }

    /// @dev Internal function to set the version for a given address.
    function _setVersion(address addr, string calldata semver) internal {
        _semver[addr] = semver.toShortString();
        emit VersionSet(addr, semver);
    }

    /**
     *
     *                              VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IProtocolRegistry
    function version(
        address addr
    ) external view returns (string memory) {
        return _semver[addr].toString();
    }

    /// @inheritdoc IProtocolRegistry
    function majorVersion(
        address addr
    ) external view returns (string memory) {
        bytes memory v = bytes(_semver[addr].toString());
        return string(abi.encodePacked(v[0]));
    }
}
