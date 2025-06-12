// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "../mixins/PermissionControllerMixin.sol";
import "../mixins/SemVerMixin.sol";
import "../libraries/Snapshots.sol";
import "./ReleaseManagerStorage.sol";

contract ReleaseManager is Initializable, ReleaseManagerStorage, PermissionControllerMixin, SemVerMixin {
    using Snapshots for Snapshots.DefaultZeroHistory;
    using EnumerableSet for EnumerableSet.Bytes32Set;
    using EnumerableMap for EnumerableMap.Bytes32ToUintMap;
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

    /// @inheritdoc IReleaseManager
    function publishRelease(
        OperatorSet calldata operatorSet,
        bytes32 releaseDigest,
        string calldata registryUrl,
        Version calldata version
    ) external checkCanCall(operatorSet.avs) {
        // Parse the next releaseId (equal to the total number of releases).
        uint256 releaseId = getTotalOperatorSetReleases(operatorSet);

        // Push the release digest and set the version.
        _setReleaseInfo(operatorSet, releaseDigest, version, uint32(0));

        // Push the release snapshot.
        _releaseSnapshots[operatorSet.key()].push(uint32(block.timestamp), releaseId);

        // Set the release registry URL for the release digest.
        _releaseRegistryUrls[operatorSet.key()][releaseDigest] = registryUrl;

        emit ReleasePublished(operatorSet, releaseDigest, registryUrl, version);
    }

    /// @inheritdoc IReleaseManager
    function deprecateRelease(
        OperatorSet calldata operatorSet,
        bytes32 releaseDigest,
        uint32 deprecationTimestamp
    ) external checkCanCall(operatorSet.avs) {
        // Get the deprecation timestamp and version for the provided digest.
        (uint32 currentDeprecationTimestamp, Version memory version) = _getReleaseInfo(operatorSet, releaseDigest);

        // Assert that the release is not already deprecated.
        require(currentDeprecationTimestamp != uint32(0), ReleaseAlreadyDeprecated());

        // Set the deprecation timestamp for the release digest.
        _setReleaseInfo(operatorSet, releaseDigest, version, uint32(deprecationTimestamp));

        emit DeprecationScheduled(operatorSet, releaseDigest, deprecationTimestamp);
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /// @dev Returns the deprecation timestamp and version for a release digest.
    function _getReleaseInfo(
        OperatorSet memory operatorSet,
        bytes32 releaseDigest
    ) internal view returns (uint32, Version memory) {
        uint256 encoded = _releaseDigests[operatorSet.key()].get(releaseDigest);

        uint32 deprecationTimestamp = uint32(encoded >> 224);
        uint16 major = uint16((encoded >> 192) & 0xFFFF);
        uint16 minor = uint16((encoded >> 160) & 0xFFFF);
        uint16 patch = uint16((encoded >> 128) & 0xFFFF);

        return (deprecationTimestamp, Version(major, minor, patch));
    }

    /// @dev Sets the release info for a release digest.
    function _setReleaseInfo(
        OperatorSet memory operatorSet,
        bytes32 releaseDigest,
        Version memory version,
        uint32 deprecationTimestamp
    ) internal {
        // Push the release digest and set the version.
        require(
            _releaseDigests[operatorSet.key()].set(
                releaseDigest,
                uint256(bytes32(abi.encodePacked(deprecationTimestamp, version.major, version.minor, version.patch)))
            ),
            ReleaseAlreadyExists()
        );
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IReleaseManager
    function isOperatorSetRelease(OperatorSet memory operatorSet, bytes32 releaseDigest) external view returns (bool) {
        return _releaseDigests[operatorSet.key()].contains(releaseDigest);
    }

    /// @inheritdoc IReleaseManager
    function getOperatorSetReleases(
        OperatorSet memory operatorSet,
        uint160 releaseId
    ) external view returns (bytes32) {
        (bytes32 releaseDigest,) = _releaseDigests[operatorSet.key()].at(releaseId);

        return releaseDigest;
    }

    /// @inheritdoc IReleaseManager
    function getOperatorSetReleases(
        OperatorSet memory operatorSet
    ) external view returns (bytes32[] memory) {
        return _releaseDigests[operatorSet.key()].keys();
    }

    /// @inheritdoc IReleaseManager
    function getTotalOperatorSetReleases(
        OperatorSet memory operatorSet
    ) public view returns (uint256) {
        return _releaseDigests[operatorSet.key()].length();
    }

    /// @inheritdoc IReleaseManager
    function getLastOperatorSetRelease(
        OperatorSet memory operatorSet
    ) public view returns (bytes32) {
        uint256 lastReleaseId = getTotalOperatorSetReleases(operatorSet) - 1;

        (bytes32 releaseDigest,) = _releaseDigests[operatorSet.key()].at(lastReleaseId);

        return releaseDigest;
    }

    /// @inheritdoc IReleaseManager
    function getReleaseAtTime(
        OperatorSet memory operatorSet,
        uint32 previousTimestamp
    ) external view returns (bytes32) {
        // Get the most recent snapshot that was published at or before the provided timestamp.
        uint256 releaseId = _releaseSnapshots[operatorSet.key()].upperLookup(previousTimestamp);

        // Return the release releaseDigest for the `releaseId`, revert if nonexistent.
        (bytes32 releaseDigest,) = _releaseDigests[operatorSet.key()].at(releaseId);

        return releaseDigest;
    }

    /// @inheritdoc IReleaseManager
    function isReleaseDeprecated(OperatorSet memory operatorSet, bytes32 releaseDigest) external view returns (bool) {
        // Get the deprecation timestamp and version for the provided digest.
        (uint32 deprecationTimestamp, Version memory version) = _getReleaseInfo(operatorSet, releaseDigest);

        // Get the the current major version for the operator set.
        (, Version memory currentVersion) = _getReleaseInfo(operatorSet, getLastOperatorSetRelease(operatorSet));

        // If the deprecation timestamp is in the future, the release is not deprecated.
        if (deprecationTimestamp > block.timestamp) return false;

        // If the major version is the same as the current major version, the release is not deprecated.
        if (version.major == currentVersion.major) return false;

        return true;
    }

    /// @inheritdoc IReleaseManager
    function getReleaseRegistryUrl(
        OperatorSet memory operatorSet,
        bytes32 releaseDigest
    ) external view returns (string memory) {
        return _releaseRegistryUrls[operatorSet.key()][releaseDigest];
    }

    /// @inheritdoc IReleaseManager
    function getReleaseVersion(
        OperatorSet memory operatorSet,
        bytes32 releaseDigest
    ) external view returns (string memory) {
        (, Version memory version) = _getReleaseInfo(operatorSet, releaseDigest);

        return string.concat(version.major.toString(), ".", version.minor.toString(), ".", version.patch.toString());
    }
}
