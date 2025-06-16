// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../libraries/OperatorSetLib.sol";

interface IReleaseManagerErrors {
    /// @notice Thrown when the upgrade by time is not in the future.
    error UpgradeByTimeNotInFuture();
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
    /// @param artifacts Array of artifacts included in the release.
    /// @param upgradeByTime Timestamp by which operators must upgrade.
    event ReleasePublished(OperatorSet indexed operatorSet, Artifact[] artifacts, uint32 upgradeByTime);
}

interface IReleaseManager is IReleaseManagerErrors, IReleaseManagerEvents {
    /**
     *
     *                         WRITE FUNCTIONS
     *
     */

    /// @notice Publishes a new release for an operator set.
    /// @param operatorSet The operator set this release is for.
    /// @param artifacts Array of artifacts to include in the release.
    /// @param upgradeByTime Timestamp by which operators must upgrade to this release.
    /// @return releaseId The index of the newly published release.
    function publishRelease(
        OperatorSet calldata operatorSet,
        Artifact[] calldata artifacts,
        uint32 upgradeByTime
    ) external returns (uint256 releaseId);

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @notice Gets the total number of releases for an operator set.
    /// @param operatorSet The operator set to query.
    /// @return The number of releases.
    function getTotalReleases(
        OperatorSet memory operatorSet
    ) external view returns (uint256);

    /// @notice Gets a specific release by index.
    /// @param operatorSet The operator set to query.
    /// @param index The index of the release to get.
    /// @return The release at the specified index.
    function getRelease(OperatorSet memory operatorSet, uint256 index) external view returns (Release memory);

    /// @notice Gets the latest release for an operator set.
    /// @param operatorSet The operator set to query.
    /// @return The most recent release.
    function getLatestRelease(
        OperatorSet memory operatorSet
    ) external view returns (Release memory);
}
