// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../libraries/OperatorSetLib.sol";

interface IReleaseManagerErrors {
    /// @notice Thrown when the upgrade by time is in the past.
    error InvalidUpgradeByTime();
}

interface IReleaseManagerTypes {
    /// @notice Represents a software artifact with its digest and registry URL.
    /// @param digest The hash digest of the artifact.
    /// @param registryUrl The URL where the artifact can be found.
    struct Artifact {
        bytes32 digest;
        string registryUrl;
    }

    /// @notice Represents a release containing multiple artifacts and an upgrade deadline.
    /// @param artifacts Array of artifacts included in this release.
    /// @param upgradeByTime Timestamp by which operators must upgrade to this release.
    struct Release {
        Artifact[] artifacts;
        uint32 upgradeByTime;
    }
}

interface IReleaseManagerEvents is IReleaseManagerTypes {
    /// @notice Emitted when a new release is published.
    /// @param operatorSet The operator set this release is for.
    /// @param releaseId The id of the release that was published.
    /// @param release The release that was published.
    event ReleasePublished(OperatorSet indexed operatorSet, uint256 indexed releaseId, Release release);
}

interface IReleaseManager is IReleaseManagerErrors, IReleaseManagerEvents {
    /**
     *
     *                         WRITE FUNCTIONS
     *
     */

    /// @notice Publishes a new release for an operator set.
    /// @param operatorSet The operator set this release is for.
    /// @param release The release that was published.
    /// @return releaseId The index of the newly published release.
    function publishRelease(
        OperatorSet calldata operatorSet,
        Release calldata release
    ) external returns (uint256 releaseId);

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @notice Returns the total number of releases for an operator set.
    /// @param operatorSet The operator set to query.
    /// @return The number of releases.
    function getTotalReleases(
        OperatorSet memory operatorSet
    ) external view returns (uint256);

    /// @notice Returns a specific release by index.
    /// @param operatorSet The operator set to query.
    /// @param releaseId The id of the release to get.
    /// @return The release at the specified index.
    function getRelease(OperatorSet memory operatorSet, uint256 releaseId) external view returns (Release memory);

    /// @notice Returns the latest release for an operator set.
    /// @param operatorSet The operator set to query.
    /// @return The id of the latest release.
    /// @return The latest release.
    function getLatestRelease(
        OperatorSet memory operatorSet
    ) external view returns (uint256, Release memory);

    /// @notice Returns the upgrade by time for the latest release.
    /// @param operatorSet The operator set to query.
    /// @return The upgrade by time for the latest release.
    function getLatestUpgradeByTime(
        OperatorSet memory operatorSet
    ) external view returns (uint32);

    /// @notice Returns true if the release is the latest release, false otherwise.
    /// @param operatorSet The operator set to query.
    /// @param releaseId The id of the release to check.
    /// @return True if the release is the latest release, false otherwise.
    function isValidRelease(OperatorSet memory operatorSet, uint256 releaseId) external view returns (bool);
}
