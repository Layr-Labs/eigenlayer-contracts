// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/Checkpoints.sol";
import "../interfaces/IReleaseManager.sol";

/**
 * @title ReleaseManagerStorage
 * @author Your Organization
 * @notice Storage contract for the ReleaseManager contract.
 */
abstract contract ReleaseManagerStorage is IReleaseManager {
    /// @notice Permission controller for UAM integration
    IPermissionController public permissionController;

    /// @notice Mapping of AVS addresses to registration status
    mapping(address => bool) public registeredAVS;

    /// @notice Mapping of AVS to array of ALL published releases (append-only)
    mapping(address => PublishedRelease[]) public allPublishedReleases;

    /// @notice Checkpoints tracking the index of releases for each AVS
    /// @dev Uses Trace256 to store uint256 indices that point to allPublishedReleases array
    mapping(address => Checkpoints.Trace256) internal releaseIndexHistory;

    /// @notice Mapping to track deprecated releases per AVS
    /// @dev Key is keccak256(avs, digest)
    mapping(bytes32 => bool) public isDeprecated;

    /// @notice List of deprecated release digests per AVS
    mapping(address => bytes32[]) public deprecatedReleases;

    /// @notice Storage gap for future upgrades
    uint256[44] private __gap;
}