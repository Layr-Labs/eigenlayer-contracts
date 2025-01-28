// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../src/contracts/libraries/Merkle.sol";

contract MerkleCaller
{
    function equals(bytes memory b1, bytes memory b2) external pure returns (bool) {
        return keccak256(abi.encode(b1)) == keccak256(abi.encode(b2));
    }

    function equals(bytes32[] memory b1, bytes32[]  memory b2) external pure returns (bool) {
        return keccak256(abi.encode(b1)) == keccak256(abi.encode(b2));
    }

    function getHash(bytes memory b) external pure returns (bytes32) {
        return keccak256(abi.encode(b));
    }

    function merkleizeSha256(bytes32[] memory leaves) external pure returns (bytes32) {
        return Merkle.merkleizeSha256(leaves);
    }

    function processInclusionProofSha256(bytes memory proof, bytes32 leaf, uint256 index) 
        external view returns (bytes32) 
    {
        return Merkle.processInclusionProofSha256(proof, leaf, index);
    }

    function processInclusionProofKeccak(bytes memory proof, bytes32 leaf, uint256 index) 
        external pure returns (bytes32) 
    {
        return Merkle.processInclusionProofKeccak(proof, leaf, index);
    }
}
