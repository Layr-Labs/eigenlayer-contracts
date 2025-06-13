// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../libraries/OperatorSetLib.sol";

interface IReleaseManagerErrors {
    /// @notice Thrown when a release with the same digest is already deprecated.
    error ReleaseAlreadyDeprecated();
    /// @notice Thrown when a deprecation is attempted that leaves an operator
    /// set without a stable version.
    error MustHaveAtLeastOneStableVersion();
}

interface IReleaseManagerTypes {
    /// @notice The type of release, following semantic versioning principles.
    /// @custom:constant MAJOR A breaking change that requires operators to upgrade.
    /// @custom:constant MINOR A backward-compatible feature addition.
    /// @custom:constant PATCH A backward-compatible bug fix.
    enum ReleaseType {
        MAJOR,
        MINOR,
        PATCH
    }

    /// @notice The current status of a release in the system.
    /// @custom:constant NONEXISTENT The release has not been published.
    /// @custom:constant DEPRECATED The release has been explicitly deprecated.
    /// @custom:constant OUTDATED A newer major version has been released.
    /// @custom:constant LIVE The release is currently active and valid.
    enum ReleaseStatus {
        NONEXISTENT,
        DEPRECATED,
        OUTDATED,
        LIVE
    }

    /// @notice A semantic version.
    /// @param major The major version.
    /// @param minor The minor version.
    /// @param patch The patch version.
    struct Version {
        uint16 major;
        uint16 minor;
        uint16 patch;
    }

    /// @notice An artifact.
    /// @param digest The digest of the artifact.
    /// @param registryUrl The registry URL of the artifact.
    struct Artifact {
        bytes32 digest;
        string registryUrl;
    }

    /// @notice A release.
    /// @param artifacts The artifacts of the release.
    /// @param deprecateAtTime The time at which the release is deprecated.
    struct Release {
        Artifact[] artifacts;
        uint32 deprecateAtTime;
    }
}

interface IReleaseManagerEvents is IReleaseManagerTypes {
    /// @notice Emitted when a new release is published for an operator set.
    /// @param operatorSetKey The key identifying the operator set.
    /// @param version The version number of the published release.
    /// @param releaseType The type of release (MAJOR, MINOR, or PATCH).
    /// @param upgradeByTime The timestamp by which operators must upgrade to this release.
    /// @param artifacts The artifacts included in this release.
    event ReleasePublished(
        bytes32 indexed operatorSetKey,
        Version indexed version,
        ReleaseType indexed releaseType,
        uint32 upgradeByTime,
        Artifact[] artifacts
    );

    /// @notice Emitted when a release is deprecated.
    /// @param operatorSetKey The key identifying the operator set.
    /// @param version The version number of the deprecated release.
    /// @param deprecateAtTime The timestamp at which the release becomes deprecated.
    event ReleaseDeprecated(bytes32 indexed operatorSetKey, Version indexed version, uint32 deprecateAtTime);

    /// @notice Emitted when the upgrade window for a major version is extended.
    /// @param operatorSetKey The key identifying the operator set.
    /// @param version The version number of the release.
    /// @param upgradeWindow The new upgrade window duration in seconds.
    event UpgradeWindowExtended(bytes32 indexed operatorSetKey, Version indexed version, uint32 upgradeWindow);
}

interface IReleaseManager is IReleaseManagerErrors, IReleaseManagerEvents {
    // WRITE

    /// @notice Publishes a new release for an operator set with the specified artifacts and upgrade window.
    /// @param operatorSet The operator set to publish the release for.
    /// @param artifacts The artifacts included in this release.
    /// @param upgradeWindow The time window (in seconds) for operators to upgrade to this release.
    /// @param releaseType The type of release (MAJOR, MINOR, or PATCH).
    /// @return The version number of the published release.
    function publishRelease(
        OperatorSet calldata operatorSet,
        Artifact[] calldata artifacts,
        uint32 upgradeWindow,
        ReleaseType releaseType
    ) external returns (Version memory);

    /// @notice Deprecates a specific release version after a given delay.
    /// @param operatorSet The operator set containing the release to deprecate.
    /// @param version The version to deprecate.
    /// @param deprecationDelay The delay (in seconds) before the release becomes deprecated.
    function deprecateRelease(
        OperatorSet calldata operatorSet,
        Version calldata version,
        uint32 deprecationDelay
    ) external;

    /// @notice Extends the upgrade window for a major version release.
    /// @param operatorSet The operator set containing the version.
    /// @param version The version to extend the upgrade window for.
    /// @param upgradeWindow The new upgrade window duration (in seconds).
    function extendUpgradeWindow(
        OperatorSet calldata operatorSet,
        Version calldata version,
        uint32 upgradeWindow
    ) external;

    // READ

    /// @notice Returns the total number of releases published for an operator set.
    /// @param operatorSet The operator set to query.
    /// @return The total number of releases.
    function getTotalVersions(
        OperatorSet memory operatorSet
    ) external view returns (uint256);

    /// @notice Returns the version at a specific index in the operator set's release history.
    /// @param operatorSet The operator set to query.
    /// @param index The index of the version to retrieve.
    /// @return The version at the specified index.
    function getVersion(OperatorSet memory operatorSet, uint256 index) external view returns (Version memory);

    /// @notice Returns the release details for a specific version.
    /// @param operatorSet The operator set containing the release.
    /// @param version The version to query.
    /// @return The release details including artifacts and deprecation time.
    function getRelease(
        OperatorSet memory operatorSet,
        Version memory version
    ) external view returns (Release memory);

    /// @notice Returns the current status of a release version.
    /// @param operatorSet The operator set containing the release.
    /// @param version The version to check.
    /// @return The status of the release (NONEXISTENT, DEPRECATED, OUTDATED, or LIVE).
    function getReleaseStatus(
        OperatorSet memory operatorSet,
        Version memory version
    ) external view returns (ReleaseStatus);

    /// @notice Returns the latest version published for an operator set.
    /// @param operatorSet The operator set to query.
    /// @return The latest version.
    function getLatestVersion(
        OperatorSet memory operatorSet
    ) external view returns (Version memory);

    /// @notice Returns the latest stable (non-deprecated) version for an operator set.
    /// @param operatorSet The operator set to query.
    /// @return The latest stable version, or max values if no stable version exists.
    function getLatestStableVersion(
        OperatorSet memory operatorSet
    ) external view returns (Version memory);
}
