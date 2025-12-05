// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/console2.sol";

import {BN254} from "src/contracts/libraries/BN254.sol";
import {Merkle} from "src/contracts/libraries/Merkle.sol";
import {LeafCalculatorMixin} from "src/contracts/mixins/LeafCalculatorMixin.sol";
import {IOperatorTableCalculatorTypes as Types} from "src/contracts/interfaces/IOperatorTableCalculator.sol";

contract BuildNonsignerProof is Script, LeafCalculatorMixin {
    using BN254 for *;

    // Build a keccak Merkle proof for a 2-operator tree (index 0 or 1)
    // Inputs are operator G1 pubkeys and weight arrays
    function runTwoPacked(
        uint256[2] memory op1XY,
        uint256[] memory w1,
        uint256[2] memory op2XY,
        uint256[] memory w2,
        uint32 indexToProve
    ) external pure returns (bytes memory proof, bytes32 root, bytes32 leaf) {
        Types.BN254OperatorInfo[] memory ops = new Types.BN254OperatorInfo[](2);
        ops[0].pubkey = BN254.G1Point(op1XY[0], op1XY[1]);
        ops[0].weights = w1;
        ops[1].pubkey = BN254.G1Point(op2XY[0], op2XY[1]);
        ops[1].weights = w2;

        bytes32[] memory leaves = new bytes32[](2);
        for (uint256 i = 0; i < 2; i++) {
            leaves[i] = calculateOperatorInfoLeaf(ops[i]);
        }

        root = Merkle.merkleizeKeccak(leaves);
        proof = Merkle.getProofKeccak(leaves, indexToProve);
        leaf = leaves[indexToProve];

        console2.log("root");
        console2.logBytes32(root);
        console2.log("leaf");
        console2.logBytes32(leaf);
        console2.log("proof bytes len");
        console2.logUint(proof.length);
    }
}
