// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "./interfaces/IPermissionController.sol";

// ============ Types ============
    enum Architecture {
        AMD64,
        ARM64
    }

    enum OperatingSystem {
        Linux,
        Darwin,
        Windows
    }

    enum ArtifactType {
        Binary,
        Container
    }

interface IReleaseManagerErrors {
    /// @notice Thrown when an AVS is not registered
    error AVSNotRegistered();
    /// @notice Thrown when an AVS is already registered
    error AVSAlreadyRegistered();
    /// @notice Thrown when an artifact digest is not found
    error ArtifactNotFound();
    /// @notice Thrown when trying to publish with an invalid deadline
    error InvalidDeadline();
    /// @notice Thrown when arrays have mismatched lengths
    error ArrayLengthMismatch();
    /// @notice Thrown when caller lacks required permissions
    error Unauthorized();
    /// @notice Thrown when attempting to deprecate an already deprecated artifact
    error AlreadyDeprecated();
}

interface IReleaseManagerEvents {
    /// @notice Emitted when artifacts are published
    event ArtifactsPublished(
        address indexed avs,
        string indexed version,
        uint256 deploymentDeadline,
        bytes32[] digests
    );

    /// @notice Emitted when an artifact is deprecated
    event ArtifactDeprecated(
        address indexed avs,
        bytes32 indexed digest
    );

    /// @notice Emitted when an AVS registers
    event AVSRegistered(address indexed avs);

    /// @notice Emitted when an AVS deregisters
    event AVSDeregistered(address indexed avs);
}

interface IReleaseManager is IReleaseManagerErrors, IReleaseManagerEvents {
    struct Artifact {
        ArtifactType artifactType;
        Architecture architecture;
        OperatingSystem os;
        bytes32 digest;
        string registryUrl;
        uint256 publishedAt;
    }

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
     * @notice Publish artifacts for an AVS with deployment information
     * @param avs The address of the AVS publishing artifacts
     * @param artifacts Array of artifacts to publish
     * @param version Semantic version string
     * @param deploymentDeadline UTC timestamp deadline for deployment
     */
    function publishArtifacts(
        address avs,
        Artifact[] calldata artifacts,
        string calldata version,
        uint256 deploymentDeadline
    ) external;

    /**
     * @notice Deprecate a specific artifact
     * @param avs The address of the AVS
     * @param digest The digest of the artifact to deprecate
     */
    function deprecateArtifact(
        address avs,
        bytes32 digest
    ) external;

    /**
     * @notice Get artifact details by AVS and digest
     * @param avs The address of the AVS
     * @param digest The digest of the artifact
     * @return The artifact details
     */
    function getArtifact(address avs, bytes32 digest) external view returns (Artifact memory);

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
     * @notice Get all deprecated artifact digests for an AVS
     * @param avs The address of the AVS
     * @return Array of deprecated artifact digests
     */
    function getDeprecatedArtifacts(
        address avs
    ) external view returns (bytes32[] memory);

    /**
     * @notice Check if an artifact is deprecated
     * @param avs The address of the AVS
     * @param digest The artifact digest
     * @return True if deprecated, false otherwise
     */
    function isArtifactDeprecated(
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