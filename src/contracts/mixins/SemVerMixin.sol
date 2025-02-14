// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "../interfaces/ISemVerMixin.sol";
import "@openzeppelin-upgrades/contracts/utils/ShortStringsUpgradeable.sol";

/// @title SemVerMixin
/// @notice A mixin contract that provides semantic versioning functionality.
/// @dev Follows SemVer 2.0.0 specification (https://semver.org/).
abstract contract SemVerMixin is ISemVerMixin {
    using ShortStringsUpgradeable for *;

    /// @notice The semantic version string for this contract, stored as a ShortString for gas efficiency.
    /// @dev Follows SemVer 2.0.0 specification (https://semver.org/). Prefixed with 'v' (e.g., "v1.2.3").
    ShortString internal immutable _VERSION;

    /// @notice Initializes the contract with a semantic version string.
    /// @param _version The SemVer-formatted version string (e.g., "v1.2.3")
    /// @dev Version should follow SemVer 2.0.0 format with 'v' prefix: vMAJOR.MINOR.PATCH
    constructor(
        string memory _version
    ) {
        _VERSION = _version.toShortString();
    }

    /// @inheritdoc ISemVerMixin
    function version() public view virtual returns (string memory) {
        return _VERSION.toString();
    }

    /// @notice Returns the major version of the contract.
    /// @dev Supports single digit major versions (e.g., "v1" for version "v1.2.3")
    /// @return The major version string (e.g., "v1" for version "v1.2.3")
    function _majorVersion() internal view returns (string memory) {
        bytes memory v = bytes(_VERSION.toString());
        return string(bytes.concat(v[0], v[1]));
    }
}
