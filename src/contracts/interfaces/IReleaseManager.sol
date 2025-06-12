// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../libraries/OperatorSetLib.sol";

interface IReleaseManagerErrors {
    error ReleaseAlreadyExists();
    error ReleaseAlreadyDeprecated();
}

interface IReleaseManagerTypes {
    struct Version {
        uint16 major;
        uint16 minor;
        uint16 patch;
    }
}

interface IReleaseManagerEvents is IReleaseManagerTypes {
    /// @notice Emitted when a new release is published for an operator set.
    /// @param operatorSet The operator set that the release was published for.
    /// @param releaseDigest The unique identifier for the release content.
    /// @param registryUrl The URL where the release artifacts can be found.
    /// @param version The semantic version of the release.
    event ReleasePublished(OperatorSet operatorSet, bytes32 releaseDigest, string registryUrl, Version version);

    /// @notice Emitted when a release deprecation is scheduled for an operator set.
    /// @param operatorSet The operator set that the release was deprecated for.
    /// @param releaseDigest The unique identifier for the release content.
    /// @param deprecationTimestamp The timestamp when the release will be deprecated.
    event DeprecationScheduled(OperatorSet operatorSet, bytes32 releaseDigest, uint32 deprecationTimestamp);
}

interface IReleaseManager is IReleaseManagerErrors, IReleaseManagerEvents {
    // WRITE

    /// @notice Publishes a new release for an operator set.
    /// @param operatorSet The operator set to publish the release for.
    /// @param digest The unique identifier for the release content.
    /// @param registryUrl The URL where the release artifacts can be found.
    /// @param version The semantic version of the release.
    /// @dev Only callable by addresses with permissioned by the operator set's AVS.
    function publishRelease(
        OperatorSet calldata operatorSet,
        bytes32 digest,
        string calldata registryUrl,
        Version calldata version
    ) external;

    /// @notice Deprecates a release for an operator set.
    /// @param operatorSet The operator set to deprecate the release for.
    /// @param digest The unique identifier for the release content.
    /// @param deprecationTimestamp The timestamp when the release will be deprecated.
    /// @dev Only callable by addresses with permissioned by the operator set's AVS.
    function deprecateRelease(OperatorSet calldata operatorSet, bytes32 digest, uint32 deprecationTimestamp) external;

    // READ

    /// @notice Checks if a release digest exists for an operator set.
    /// @param operatorSet The operator set to check.
    /// @param releaseDigest The release digest to check.
    /// @return bool True if the release digest exists, false otherwise.
    function isOperatorSetRelease(OperatorSet memory operatorSet, bytes32 releaseDigest) external view returns (bool);

    /// @notice Gets a specific release digest for an operator set by its ID.
    /// @param operatorSet The operator set to get the release from.
    /// @param releaseId The ID of the release to get.
    /// @return bytes32 The release digest.
    function getOperatorSetReleases(
        OperatorSet memory operatorSet,
        uint160 releaseId
    ) external view returns (bytes32);

    /// @notice Gets all release digests for an operator set.
    /// @param operatorSet The operator set to get releases from.
    /// @return bytes32[] Array of release digests.
    function getOperatorSetReleases(
        OperatorSet memory operatorSet
    ) external view returns (bytes32[] memory);

    /// @notice Gets the total number of releases for an operator set.
    /// @param operatorSet The operator set to count releases for.
    /// @return uint256 The total number of releases.
    function getTotalOperatorSetReleases(
        OperatorSet memory operatorSet
    ) external view returns (uint256);

    /// @notice Gets the most recent release digest for an operator set.
    /// @param operatorSet The operator set to get the latest release from.
    /// @return bytes32 The latest release digest.
    function getLastOperatorSetRelease(
        OperatorSet memory operatorSet
    ) external view returns (bytes32);

    /// @notice Gets the release digest that was active at a specific timestamp.
    /// @param operatorSet The operator set to get the release from.
    /// @param previousTimestamp The timestamp to check.
    /// @return bytes32 The release digest that was active at the timestamp.
    function getReleaseAtTime(
        OperatorSet memory operatorSet,
        uint32 previousTimestamp
    ) external view returns (bytes32);

    /// @notice Checks if a release is deprecated.
    /// @param operatorSet The operator set to check.
    /// @param releaseDigest The release digest to check.
    /// @return bool True if the release is deprecated, false otherwise.
    function isReleaseDeprecated(OperatorSet memory operatorSet, bytes32 releaseDigest) external view returns (bool);

    /// @notice Returns the registry URL for a release digest.
    /// @param operatorSet The operator set containing the release.
    /// @param releaseDigest The release digest to get the registry URL for.
    /// @return string The registry URL.
    function getReleaseRegistryUrl(
        OperatorSet memory operatorSet,
        bytes32 releaseDigest
    ) external view returns (string memory);

    /// @notice Returns the semantic version string for a release.
    /// @param operatorSet The operator set containing the release.
    /// @param releaseDigest The release digest to get the version for.
    /// @return string The semantic version string (e.g. "1.2.3").
    function getReleaseVersion(
        OperatorSet memory operatorSet,
        bytes32 releaseDigest
    ) external view returns (string memory);
}
