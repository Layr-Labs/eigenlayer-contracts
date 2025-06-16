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
        Release[] storage releases = _operatorSetReleases[operatorSet.key()];

        uint256 newTotalReleases = releases.length + 1;

        assembly {
            sstore(releases.slot, newTotalReleases)
        }

        Release storage release = releases[newTotalReleases - 1];

        for (uint256 i = 0; i < artifacts.length; ++i) {
            release.artifacts.push(artifacts[i]);
        }

        release.upgradeByTime = upgradeByTime;

        emit ReleasePublished(operatorSet, artifacts, upgradeByTime);
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */
    function getTotalReleases(
        OperatorSet memory operatorSet
    ) external view returns (uint256) {
        return _operatorSetReleases[operatorSet.key()].length;
    }

    function getReleaseArtifacts(
        OperatorSet memory operatorSet,
        uint256 index
    ) external view returns (Artifact[] memory) {
        return _operatorSetReleases[operatorSet.key()][index].artifacts;
    }

    function getLatestReleaseArtifacts(
        OperatorSet memory operatorSet
    ) external view returns (Artifact[] memory) {
        Release[] storage releases = _operatorSetReleases[operatorSet.key()];

        return releases[releases.length - 1].artifacts;
    }
}
