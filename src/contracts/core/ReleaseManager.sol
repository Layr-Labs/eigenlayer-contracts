// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {OwnableUpgradeable} from "./OwnableUpgradeable.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import {Checkpoints} from "@openzeppelin/contracts/utils/structs/Checkpoints.sol";

import {IReleaseManager} from "../interfaces/IReleaseManager.sol";
import {IPermissionController} from "../interfaces/IPermissionController.sol";
import {ReleaseManagerStorage} from "./ReleaseManagerStorage.sol";

/**
 * @title ReleaseManager
 * @author Your Organization
 * @notice Contract for managing an append-only log of OCI artifact releases for AVS deployments.
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
    function publishRelease(
        address avs,
        bytes32 digest,
        string calldata registryUrl,
        string calldata version,
        uint256 deploymentDeadline
    ) external nonReentrant onlyRegistered(avs) onlyAuthorized(avs) {
        if (deploymentDeadline == 0) revert InvalidDeadline();
        if (digest == bytes32(0)) revert InvalidDigest();
        if (bytes(registryUrl).length == 0) revert InvalidDigest();

        PublishedRelease memory release = PublishedRelease({
            digest: digest,
            registryUrl: registryUrl,
            version: version,
            deploymentDeadline: deploymentDeadline,
            publishedAt: block.timestamp
        });

        // Add to the append-only log
        allPublishedReleases[avs].push(release);

        // Update the checkpoint to point to the new index
        uint256 newIndex = allPublishedReleases[avs].length - 1;
        releaseIndexHistory[avs].push(block.number, newIndex);

        emit ReleasePublished(avs, version, digest, registryUrl, deploymentDeadline);
    }

    /// @inheritdoc IReleaseManager
    function deprecateRelease(
        address avs,
        bytes32 digest
    ) external nonReentrant onlyRegistered(avs) onlyAuthorized(avs) {
        if (digest == bytes32(0)) revert InvalidDigest();

        // Mark as deprecated
        bytes32 deprecationKey = keccak256(abi.encodePacked(avs, digest));
        if (isDeprecated[deprecationKey]) revert AlreadyDeprecated();

        isDeprecated[deprecationKey] = true;

        // Add to deprecated list
        deprecatedReleases[avs].push(digest);

        emit ReleaseDeprecated(avs, digest);
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
     *
     *                         VIEW FUNCTIONS
     *
     */

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
        if (releases.length == 0) revert ReleaseNotFound();
        return releases[releases.length - 1];
    }

    /// @inheritdoc IReleaseManager
    function getReleaseAtBlock(
        address avs,
        uint256 blockNumber
    ) external view returns (PublishedRelease memory) {
        uint256 index = releaseIndexHistory[avs].upperLookup(blockNumber);
        PublishedRelease[] memory releases = allPublishedReleases[avs];
        if (releases.length == 0 || index >= releases.length) revert ReleaseNotFound();
        return releases[index];
    }

    /// @inheritdoc IReleaseManager
    function getDeprecatedReleases(
        address avs
    ) external view returns (bytes32[] memory) {
        return deprecatedReleases[avs];
    }

    /// @inheritdoc IReleaseManager
    function isReleaseDeprecated(
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