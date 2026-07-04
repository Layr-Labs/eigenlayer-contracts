// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/interfaces/IReleaseManager.sol";

contract ReleaseManagerMock is IReleaseManager {
    mapping(bytes32 => bool) public hasRelease;
    mapping(bytes32 => string) public metadataURIs;

    function publishRelease(OperatorSet calldata operatorSet, Release calldata release) external override returns (uint releaseId) {
        bytes32 key = keccak256(abi.encode(operatorSet.avs, operatorSet.id));
        hasRelease[key] = true;
        return 1;
    }

    function publishMetadataURI(OperatorSet calldata operatorSet, string calldata metadataURI) external override {
        bytes32 key = keccak256(abi.encode(operatorSet.avs, operatorSet.id));
        metadataURIs[key] = metadataURI;
    }

    function getLatestRelease(OperatorSet memory operatorSet) external view override returns (uint, Release memory) {
        bytes32 key = keccak256(abi.encode(operatorSet.avs, operatorSet.id));
        if (!hasRelease[key]) revert NoReleases();

        Artifact[] memory artifacts = new Artifact[](1);
        artifacts[0] = Artifact({digest: keccak256("test-artifact"), registry: "https://example.com/registry"});

        return (1, Release({artifacts: artifacts, upgradeByTime: uint32(block.timestamp + 7 days)}));
    }

    function getRelease(OperatorSet memory operatorSet, uint releaseId) external view override returns (Release memory) {
        Artifact[] memory artifacts = new Artifact[](1);
        artifacts[0] = Artifact({digest: keccak256("test-artifact"), registry: "https://example.com/registry"});

        return Release({artifacts: artifacts, upgradeByTime: uint32(block.timestamp + 7 days)});
    }

    function getTotalReleases(OperatorSet memory operatorSet) external view override returns (uint) {
        bytes32 key = keccak256(abi.encode(operatorSet.avs, operatorSet.id));
        return hasRelease[key] ? 1 : 0;
    }

    function getLatestUpgradeByTime(OperatorSet memory operatorSet) external view override returns (uint32) {
        bytes32 key = keccak256(abi.encode(operatorSet.avs, operatorSet.id));
        if (!hasRelease[key]) return 0;
        return uint32(block.timestamp + 7 days);
    }

    function isValidRelease(OperatorSet memory operatorSet, uint releaseId) external view override returns (bool) {
        bytes32 key = keccak256(abi.encode(operatorSet.avs, operatorSet.id));
        return hasRelease[key] && releaseId == 1;
    }

    function getMetadataURI(OperatorSet memory operatorSet) external view override returns (string memory) {
        bytes32 key = keccak256(abi.encode(operatorSet.avs, operatorSet.id));
        return metadataURIs[key];
    }

    function setHasRelease(OperatorSet memory operatorSet, bool _hasRelease) external {
        bytes32 key = keccak256(abi.encode(operatorSet.avs, operatorSet.id));
        hasRelease[key] = _hasRelease;
    }
}
