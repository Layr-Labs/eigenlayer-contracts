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

        require(bytes(_operatorSetMetadataURI[operatorSet.key()]).length != 0, MustPublishMetadataURI());
        require(release.upgradeByTime == 0 || release.upgradeByTime >= block.timestamp, InvalidUpgradeByTime());

        // New release id is the length of the array before this call.
        releaseId = releases.length;
        // Increment total releases for this operator set.
        releases.push(release);

        emit ReleasePublished(operatorSet, releaseId, release);
    }

    /// @inheritdoc IReleaseManager
    function publishMetadataURI(
        OperatorSet calldata operatorSet,
        string calldata metadataURI
    ) external checkCanCall(operatorSet.avs) {
        require(bytes(metadataURI).length != 0, InvalidMetadataURI());
        _operatorSetMetadataURI[operatorSet.key()] = metadataURI;
        emit MetadataURIPublished(operatorSet, metadataURI);
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
    ) external view returns (uint256, Release memory) {
        Release[] storage releases = _operatorSetReleases[operatorSet.key()];
        require(releases.length > 0, NoReleases());

        uint256 latestReleaseId = releases.length - 1;
        return (latestReleaseId, releases[latestReleaseId]);
    }

    /// @inheritdoc IReleaseManager
    function getLatestUpgradeByTime(
        OperatorSet memory operatorSet
    ) external view returns (uint32) {
        Release[] storage releases = _operatorSetReleases[operatorSet.key()];
        require(releases.length > 0, NoReleases());

        uint256 latestReleaseId = releases.length - 1;
        return releases[latestReleaseId].upgradeByTime;
    }

    /// @inheritdoc IReleaseManager
    function isValidRelease(OperatorSet memory operatorSet, uint256 releaseId) external view returns (bool) {
        uint256 totalReleases = getTotalReleases(operatorSet);
        require(totalReleases > 0, NoReleases());
        return releaseId == totalReleases - 1;
    }

    /// @inheritdoc IReleaseManager
    function getMetadataURI(
        OperatorSet memory operatorSet
    ) external view returns (string memory) {
        return _operatorSetMetadataURI[operatorSet.key()];
    }
}
