// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import {Checkpoints} from "@openzeppelin/contracts/utils/structs/Checkpoints.sol";

import {IReleaseManager} from "../interfaces/IReleaseManager.sol";
import {IPermissionController} from "../interfaces/IPermissionController.sol";
import {ReleaseManagerStorage} from "./ReleaseManagerStorage.sol";

/**
 * @title ReleaseManager
 * @author EigenLabs
 * @notice Contract for managing an append-only log of artifact releases for AVS deployments.
 */
contract ReleaseManager is ReentrancyGuard, OwnableUpgradeable, ReleaseManagerStorage {
    /**
     *
     *                         MODIFIERS
     *
     */

    modifier onlyRegistered(address avs) {
        if (!registeredAVS[avs]) revert AVSNotRegistered();
        _;
    }

    modifier onlyAuthorized(address avs) {
        if (!_isAuthorized(avs, msg.sender)) revert Unauthorized();
        _;
    }

    /**
     *
     *                         INITIALIZATION
     *
     */

    function initialize(address _permissionController) public initializer {
        __Ownable_init();
        permissionController = IPermissionController(_permissionController);
    }

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

    /// @inheritdoc IReleaseManager
    function register(address avs) external onlyAuthorized(avs) {
        if (registeredAVS[avs]) revert AVSAlreadyRegistered();
        registeredAVS[avs] = true;
        emit AVSRegistered(avs);
    }

    /// @inheritdoc IReleaseManager
    function deregister(address avs) external onlyAuthorized(avs) {
        if (!registeredAVS[avs]) revert AVSNotRegistered();
        registeredAVS[avs] = false;
        emit AVSDeregistered(avs);
    }

    /// @inheritdoc IReleaseManager
    function publishArtifacts(
        address avs,
        Artifact[] calldata _artifacts,
        string calldata version,
        uint256 deploymentDeadline
    ) external nonReentrant onlyRegistered(avs) onlyAuthorized(avs) {
        if (deploymentDeadline == 0) revert InvalidDeadline();
        if (_artifacts.length == 0) revert ArrayLengthMismatch();

        bytes32[] memory digests = new bytes32[](_artifacts.length);

        // First, store all artifacts
        for (uint256 i = 0; i < _artifacts.length; i++) {
            Artifact memory artifact = _artifacts[i];
            artifact.publishedAt = block.timestamp;

            // Create unique key: keccak256(abi.encodePacked(avs, digest))
            bytes32 key = _getArtifactKey(avs, artifact.digest);

            // Store artifact
            artifacts[key] = artifact;
            artifactExists[key] = true;
            digests[i] = artifact.digest;
        }

        // Create releases for each artifact
        for (uint256 i = 0; i < _artifacts.length; i++) {
            PublishedRelease memory release = PublishedRelease({
                digest: _artifacts[i].digest,
                registryUrl: _artifacts[i].registryUrl,
                version: version,
                deploymentDeadline: deploymentDeadline,
                publishedAt: block.timestamp
            });

            // Add to the append-only log
            allPublishedReleases[avs].push(release);

            // Update the checkpoint to point to the new index
            uint256 newIndex = allPublishedReleases[avs].length - 1;
            releaseIndexHistory[avs].push(block.number, newIndex);
        }

        emit ArtifactsPublished(avs, version, deploymentDeadline, digests);
    }

    /// @inheritdoc IReleaseManager
    function deprecateArtifact(
        address avs,
        bytes32 digest
    ) external nonReentrant onlyRegistered(avs) onlyAuthorized(avs) {
        // Verify artifact exists
        bytes32 artifactKey = _getArtifactKey(avs, digest);
        if (!artifactExists[artifactKey]) revert ArtifactNotFound();

        // Mark as deprecated
        bytes32 deprecationKey = keccak256(abi.encodePacked(avs, digest));
        if (isDeprecated[deprecationKey]) revert AlreadyDeprecated();

        isDeprecated[deprecationKey] = true;

        // Add to deprecated list
        deprecatedArtifacts[avs].push(digest);

        emit ArtifactDeprecated(avs, digest);
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Check if caller is authorized to act on behalf of AVS
     * @param avs The AVS address
     * @param caller The caller address
     * @return True if authorized
     */
    function _isAuthorized(address avs, address caller) internal view returns (bool) {
        // Check if caller is the AVS itself
        if (avs == caller) return true;

        // Check if caller is an admin via permission controller
        if (address(permissionController) != address(0)) {
            return permissionController.isAdmin(avs, caller);
        }

        return false;
    }

    /**
     * @notice Generate unique key for artifact storage
     * @param avs The AVS address
     * @param digest The artifact digest
     * @return The unique key
     */
    function _getArtifactKey(address avs, bytes32 digest) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(avs, digest));
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IReleaseManager
    function getArtifact(address avs, bytes32 digest) external view returns (Artifact memory) {
        bytes32 key = _getArtifactKey(avs, digest);
        if (!artifactExists[key]) revert ArtifactNotFound();
        return artifacts[key];
    }

    /// @inheritdoc IReleaseManager
    function getPublishedReleases(
        address avs
    ) external view returns (PublishedRelease[] memory) {
        return allPublishedReleases[avs];
    }

    /// @inheritdoc IReleaseManager
    function getLatestRelease(
        address avs
    ) external view returns (PublishedRelease memory) {
        PublishedRelease[] memory releases = allPublishedReleases[avs];
        if (releases.length == 0) revert ArtifactNotFound();
        return releases[releases.length - 1];
    }

    /// @inheritdoc IReleaseManager
    function getReleaseAtBlock(
        address avs,
        uint256 blockNumber
    ) external view returns (PublishedRelease memory) {
        uint256 index = releaseIndexHistory[avs].upperLookup(blockNumber);
        PublishedRelease[] memory releases = allPublishedReleases[avs];
        if (releases.length == 0 || index >= releases.length) revert ArtifactNotFound();
        return releases[index];
    }

    /// @inheritdoc IReleaseManager
    function getDeprecatedArtifacts(
        address avs
    ) external view returns (bytes32[] memory) {
        return deprecatedArtifacts[avs];
    }

    /// @inheritdoc IReleaseManager
    function isArtifactDeprecated(
        address avs,
        bytes32 digest
    ) external view returns (bool) {
        bytes32 deprecationKey = keccak256(abi.encodePacked(avs, digest));
        return isDeprecated[deprecationKey];
    }

    /// @inheritdoc IReleaseManager
    function getReleaseCheckpointCount(
        address avs
    ) external view returns (uint256) {
        return releaseIndexHistory[avs].length();
    }
}