// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../libraries/OperatorSetLib.sol";

interface IReleaseManagerErrors {
    /// @notice Thrown when a release with the same digest already exists.
    error ReleaseAlreadyExists();
    /// @notice Thrown when a release with the same digest is already deprecated.
    error ReleaseAlreadyDeprecated();
    /// @notice Thrown when a deprecation is attempted that leaves an operator
    /// set without a stable version.
    error MustHaveAtLeastOneStableVersion();
}

interface IReleaseManagerTypes {
    enum ReleaseType {
        MAJOR,
        MINOR,
        PATCH
    }

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
    event ReleasePublished(
        bytes32 indexed operatorSetKey,
        Version indexed version,
        ReleaseType indexed releaseType,
        uint32 upgradeByTime,
        Artifact[] artifacts
    );

    event ReleaseDeprecated(bytes32 indexed operatorSetKey, Version indexed version, uint32 deprecateAtTime);

    event UpgradeWindowExtended(bytes32 indexed operatorSetKey, Version indexed version, uint32 upgradeWindow);
}

interface IReleaseManager is IReleaseManagerErrors, IReleaseManagerEvents {
// WRITE

// READ
}
