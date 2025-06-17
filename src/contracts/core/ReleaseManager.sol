// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "../mixins/PermissionControllerMixin.sol";
import "../mixins/SemVerMixin.sol";
import "./ReleaseManagerStorage.sol";

contract ReleaseManager is Initializable, ReleaseManagerStorage, PermissionControllerMixin, SemVerMixin {
    using OperatorSetLib for OperatorSet;

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */
    constructor(
        IPermissionController _permissionController,
        string memory _version
    ) PermissionControllerMixin(_permissionController) SemVerMixin(_version) {
        _disableInitializers();
    }

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

    /// @inheritdoc IReleaseManager
    function publishRelease(
        OperatorSet calldata operatorSet,
        Release calldata release
    ) external checkCanCall(operatorSet.avs) returns (uint256 releaseId) {
        Release[] storage releases = _operatorSetReleases[operatorSet.key()];

        require(release.upgradeByTime >= block.timestamp, InvalidUpgradeByTime());

        // New release id is the length of the array before this call.
        releaseId = releases.length;
        // Increment total releases for this operator set.
        releases.push();
        // Copy the release to storage.
        for (uint256 i = 0; i < release.artifacts.length; ++i) {
            releases[releaseId].artifacts.push(release.artifacts[i]);
        }
        releases[releaseId].upgradeByTime = release.upgradeByTime;

        emit ReleasePublished(operatorSet, releaseId, release);
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IReleaseManager
    function getTotalReleases(
        OperatorSet memory operatorSet
    ) public view returns (uint256) {
        return _operatorSetReleases[operatorSet.key()].length;
    }

    /// @inheritdoc IReleaseManager
    function getRelease(OperatorSet memory operatorSet, uint256 releaseId) external view returns (Release memory) {
        return _operatorSetReleases[operatorSet.key()][releaseId];
    }

    /// @inheritdoc IReleaseManager
    function getLatestRelease(
        OperatorSet memory operatorSet
    ) public view returns (uint256, Release memory) {
        Release[] storage releases = _operatorSetReleases[operatorSet.key()];
        uint256 latestReleaseId = releases.length - 1;
        return (latestReleaseId, releases[latestReleaseId]);
    }

    /// @inheritdoc IReleaseManager
    function getLatestUpgradeByTime(
        OperatorSet memory operatorSet
    ) external view returns (uint32) {
        Release[] storage releases = _operatorSetReleases[operatorSet.key()];
        uint256 latestReleaseId = releases.length - 1;
        return releases[latestReleaseId].upgradeByTime;
    }

    /// @inheritdoc IReleaseManager
    function isValidRelease(OperatorSet memory operatorSet, uint256 releaseId) external view returns (bool) {
        return releaseId == getTotalReleases(operatorSet) - 1;
    }
}
