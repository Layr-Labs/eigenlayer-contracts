// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../libraries/OperatorSetLib.sol";

interface IReleaseManagerErrors {}

interface IReleaseManagerTypes {
    struct Artifact {
        bytes32 digest;
        string registryUrl;
    }

    struct Release {
        Artifact[] artifacts;
        uint32 upgradeByTime;
    }
}

interface IReleaseManagerEvents is IReleaseManagerTypes {
    event ReleasePublished(OperatorSet indexed operatorSet, Artifact[] artifacts, uint32 upgradeByTime);
}

interface IReleaseManager is IReleaseManagerErrors, IReleaseManagerEvents {
    // WRITE

    function publishRelease(
        OperatorSet calldata operatorSet,
        Artifact[] calldata artifacts,
        uint32 upgradeByTime
    ) external;

    // READ

    function getTotalReleases(
        OperatorSet memory operatorSet
    ) external view returns (uint256);

    function getReleaseArtifacts(
        OperatorSet memory operatorSet,
        uint256 index
    ) external view returns (Artifact[] memory);

    function getLatestReleaseArtifacts(
        OperatorSet memory operatorSet
    ) external view returns (Artifact[] memory);
}
