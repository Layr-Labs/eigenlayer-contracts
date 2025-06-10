// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";

interface IReleaseManagerErrors {
    /// @notice Thrown when an AVS is not registered
    error AVSNotRegistered();
    /// @notice Thrown when an AVS is already registered
    error AVSAlreadyRegistered();
    /// @notice Thrown when a release is not found
    error ReleaseNotFound();
    /// @notice Thrown when trying to publish with an invalid deadline
    error InvalidDeadline();
    /// @notice Thrown when caller lacks required permissions
    error Unauthorized();
    /// @notice Thrown when attempting to deprecate an already deprecated release
    error AlreadyDeprecated();
    /// @notice Thrown when digest is invalid
    error InvalidDigest();
}

interface IReleaseManagerEvents {
    /// @notice Emitted when a release is published
    event ReleasePublished(
        address indexed avs,
        string indexed version,
        bytes32 indexed digest,
        string registryUrl,
        uint256 deploymentDeadline
    );

    /// @notice Emitted when a release is deprecated
    event ReleaseDeprecated(
        address indexed avs,
        bytes32 indexed digest
    );

    /// @notice Emitted when an AVS registers
    event AVSRegistered(address indexed avs);

    /// @notice Emitted when an AVS deregisters
    event AVSDeregistered(address indexed avs);
}

interface IReleaseManager is IReleaseManagerErrors, IReleaseManagerEvents {
    struct PublishedRelease {
        bytes32 digest;
        string registryUrl;
        string version;
        uint256 deploymentDeadline;
        uint256 publishedAt;
    }

    // ============ Functions ============

    /**
     * @notice Register an AVS to use the ReleaseManager
     * @param avs The address of the AVS to register
     */
    function register(address avs) external;

    /**
     * @notice Deregister an AVS from the ReleaseManager
     * @param avs The address of the AVS to deregister
     */
    function deregister(address avs) external;

    /**
     * @notice Publish a release for an AVS
     * @param avs The address of the AVS publishing the release
     * @param digest The OCI artifact digest (sha256 hash)
     * @param registryUrl The full OCI registry URL
     * @param version Semantic version string
     * @param deploymentDeadline UTC timestamp deadline for deployment
     */
    function publishRelease(
        address avs,
        bytes32 digest,
        string calldata registryUrl,
        string calldata version,
        uint256 deploymentDeadline
    ) external;

    /**
     * @notice Deprecate a specific release
     * @param avs The address of the AVS
     * @param digest The digest of the release to deprecate
     */
    function deprecateRelease(
        address avs,
        bytes32 digest
    ) external;

    /**
     * @notice Get all published releases for an AVS
     * @param avs The address of the AVS
     * @return Array of published releases
     */
    function getPublishedReleases(
        address avs
    ) external view returns (PublishedRelease[] memory);

    /**
     * @notice Get the latest published release for an AVS
     * @param avs The address of the AVS
     * @return The latest published release
     */
    function getLatestRelease(
        address avs
    ) external view returns (PublishedRelease memory);

    /**
     * @notice Get the published release that was active at a specific block
     * @param avs The address of the AVS
     * @param blockNumber The block number to query
     * @return The published release that was active at that block
     */
    function getReleaseAtBlock(
        address avs,
        uint256 blockNumber
    ) external view returns (PublishedRelease memory);

    /**
     * @notice Get all deprecated release digests for an AVS
     * @param avs The address of the AVS
     * @return Array of deprecated release digests
     */
    function getDeprecatedReleases(
        address avs
    ) external view returns (bytes32[] memory);

    /**
     * @notice Check if a release is deprecated
     * @param avs The address of the AVS
     * @param digest The release digest
     * @return True if deprecated, false otherwise
     */
    function isReleaseDeprecated(
        address avs,
        bytes32 digest
    ) external view returns (bool);

    /**
     * @notice Get the number of checkpoints in the release history
     * @param avs The address of the AVS
     * @return The number of release checkpoints
     */
    function getReleaseCheckpointCount(
        address avs
    ) external view returns (uint256);
}