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

    /// @inheritdoc IReleaseManager
    function publishRelease(
        OperatorSet calldata operatorSet,
        Artifact[] calldata artifacts,
        uint32 upgradeByTime
    ) external checkCanCall(operatorSet.avs) {
        // Create a storage pointer to the releases array for this operator set for readibility.
        Release[] storage releases = _operatorSetReleases[operatorSet.key()];

        // Add a new empty release to the end of the array.
        releases.push();

        // Create a storage pointer to the newly added release for readibility.
        Release storage release = releases[releases.length - 1];

        // Copy the release to storage.
        for (uint256 i = 0; i < artifacts.length; ++i) {
            release.artifacts.push(artifacts[i]);
        }

        // Store the deadline by which operators should upgrade to this release.
        release.upgradeByTime = upgradeByTime;

        // Emit event with the release details.
        emit ReleasePublished(operatorSet, artifacts, upgradeByTime);
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IReleaseManager
    function getTotalReleases(
        OperatorSet memory operatorSet
    ) external view returns (uint256) {
        return _operatorSetReleases[operatorSet.key()].length;
    }

    /// @inheritdoc IReleaseManager
    function getRelease(
        OperatorSet memory operatorSet,
        uint256 index
    ) external view returns (Release memory) {
        return _operatorSetReleases[operatorSet.key()][index];
    }

    /// @inheritdoc IReleaseManager
    function getLatestRelease(
        OperatorSet memory operatorSet
    ) external view returns (Release memory) {
        Release[] storage releases = _operatorSetReleases[operatorSet.key()];

        return releases[releases.length - 1];
    }
}
