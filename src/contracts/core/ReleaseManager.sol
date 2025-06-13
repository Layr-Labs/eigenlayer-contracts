// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "../mixins/PermissionControllerMixin.sol";
import "../mixins/SemVerMixin.sol";
import "./ReleaseManagerStorage.sol";

contract ReleaseManager is Initializable, ReleaseManagerStorage, PermissionControllerMixin, SemVerMixin {
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using OperatorSetLib for OperatorSet;
    using Strings for uint16;

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */
    constructor(
        IAllocationManager _allocationManager,
        IPermissionController _permissionController,
        string memory _version
    )
        ReleaseManagerStorage(_allocationManager)
        PermissionControllerMixin(_permissionController)
        SemVerMixin(_version)
    {
        _disableInitializers();
    }

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */
    function publishRelease(
        OperatorSet calldata operatorSet,
        Artifact[] calldata artifacts,
        uint32 upgradeWindow,
        ReleaseType releaseType
    ) external checkCanCall(operatorSet.avs) returns (Version memory) {
        Version memory version = getLatestVersion(operatorSet);

        uint32 upgradeByTime = uint32(block.timestamp + upgradeWindow);

        if (releaseType == ReleaseType.MAJOR) {
            ++version.major;
            version.minor = 0;
            version.patch = 0;
            _upgradeByTimes[operatorSet.key()][version.major] = upgradeByTime;
        } else if (releaseType == ReleaseType.MINOR) {
            ++version.minor;
            version.patch = 0;
        } else if (releaseType == ReleaseType.PATCH) {
            ++version.patch;
        }

        bytes32 versionKey = _encodeVersion(version);

        // Append the version to the operator set's version history.
        _versions[operatorSet.key()].add(versionKey);
        // Map the version to the release artifacts and deprecation time.
        Release storage release = _releases[operatorSet.key()][versionKey];
        for (uint256 i = 0; i < artifacts.length; i++) {
            release.artifacts.push(artifacts[i]);
        }

        emit ReleasePublished(operatorSet.key(), version, releaseType, upgradeByTime, artifacts);

        return version;
    }

    function deprecateRelease(
        OperatorSet calldata operatorSet,
        Version calldata version,
        uint32 deprecationDelay
    ) public checkCanCall(operatorSet.avs) {
        bytes32 versionKey = _encodeVersion(version);

        // Check that the release is not already deprecated and modify state.
        Release storage release = _releases[operatorSet.key()][versionKey];
        require(release.deprecateAtTime == 0, ReleaseAlreadyDeprecated());
        uint32 deprecateAtTime = uint32(block.timestamp + deprecationDelay);
        release.deprecateAtTime = deprecateAtTime;

        // Checked that there is at least one stable version after deprecating this release.
        Version memory lastStableVersion = getLatestStableVersion(operatorSet);
        require(lastStableVersion.major != type(uint16).max, MustHaveAtLeastOneStableVersion());

        emit ReleaseDeprecated(operatorSet.key(), version, deprecateAtTime);
    }

    function extendUpgradeWindow(
        OperatorSet calldata operatorSet,
        Version calldata version,
        uint32 upgradeWindow
    ) public checkCanCall(operatorSet.avs) {
        _upgradeByTimes[operatorSet.key()][version.major] = uint32(block.timestamp + upgradeWindow);

        // emit UpgradeWindowExtended(operatorSet.key(), _encodeVersion(version), upgradeWindow);
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */
    function _encodeVersion(
        Version memory version
    ) internal pure returns (bytes32) {
        return bytes32(abi.encodePacked(version.major, version.minor, version.patch));
    }

    function _decodeVersion(
        bytes32 encoded
    ) internal pure returns (Version memory) {
        return Version({
            major: uint16((uint256(encoded) >> 208) & type(uint16).max),
            minor: uint16((uint256(encoded) >> 192) & type(uint16).max),
            patch: uint16((uint256(encoded) >> 176) & type(uint16).max)
        });
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */
    function getTotalReleaseCount(
        OperatorSet memory operatorSet
    ) external view returns (uint256) {
        return _versions[operatorSet.key()].length();
    }

    function getVersion(OperatorSet memory operatorSet, uint256 index) external view returns (Version memory) {
        return _decodeVersion(_versions[operatorSet.key()].at(index));
    }

    function getRelease(
        OperatorSet memory operatorSet,
        Version memory version
    ) external view returns (Release memory) {
        return _releases[operatorSet.key()][_encodeVersion(version)];
    }

    function getReleaseStatus(
        OperatorSet memory operatorSet,
        Version memory version
    ) external view returns (ReleaseStatus) {
        // First, check whether the version exists.
        bool exists = _versions[operatorSet.key()].contains(_encodeVersion(version));

        // If the version does not exist, it is not valid.
        if (!exists) return ReleaseStatus.NONEXISTENT;

        // Second, check whether the version is deprecated by a force deprecation.
        uint32 deprecateAtTime = _releases[operatorSet.key()][_encodeVersion(version)].deprecateAtTime;
        if (deprecateAtTime != 0 && block.timestamp >= deprecateAtTime) return (ReleaseStatus.DEPRECATED);

        // Third, check whether the version is deprecated by a major version upgrade.
        uint32 lastUpgradeByTime = _upgradeByTimes[operatorSet.key()][getLatestStableVersion(operatorSet).major];
        if (lastUpgradeByTime != 0 && block.timestamp >= lastUpgradeByTime) return (ReleaseStatus.OUTDATED);

        // Otherwise, the version is live.
        return (ReleaseStatus.LIVE);
    }

    function getLatestVersion(
        OperatorSet memory operatorSet
    ) public view returns (Version memory) {
        EnumerableSet.Bytes32Set storage versions = _versions[operatorSet.key()];
        return _decodeVersion(versions.at(versions.length() - 1));
    }

    function getLatestStableVersion(
        OperatorSet memory operatorSet
    ) public view returns (Version memory) {
        EnumerableSet.Bytes32Set storage versions = _versions[operatorSet.key()];
        uint256 versionCount = versions.length();

        // Linear search backwards for the latest stable version.
        for (uint256 i = versionCount - 1; i >= 0; i--) {
            bytes32 versionKey = versions.at(i);
            uint32 deprecateAtTime = _releases[operatorSet.key()][versionKey].deprecateAtTime;

            // If the release is deprecated, skip it.
            if (block.timestamp >= deprecateAtTime) {
                continue;
            }

            // Othersise, return the release version.
            return _decodeVersion(versionKey);
        }

        // No version has been published.
        return Version({major: type(uint16).max, minor: type(uint16).max, patch: type(uint16).max});
    }
}
