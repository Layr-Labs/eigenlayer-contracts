// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/libraries/BN254.sol";
import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";

contract Operators is Test {
    string internal operatorConfigJson;

    constructor() {
        operatorConfigJson = vm.readFile("./src/test/test-data/operators.json");
    }

    function operatorPrefix(uint256 index) public pure returns(string memory) {
        return string.concat(".operators[", string.concat(vm.toString(index), "]."));
    }

    function getNumOperators() public returns(uint256) {
        return stdJson.readUint(operatorConfigJson, ".numOperators");
    }

    function getOperatorAddress(uint256 index) public returns(address) {
        return stdJson.readAddress(operatorConfigJson, string.concat(operatorPrefix(index), "Address"));
    }

    function getOperatorSchnorrSignature(uint256 index) public returns(uint256, BN254.G1Point memory) {
        uint256 s = readUint(operatorConfigJson, index, "SField");
        BN254.G1Point memory pubkey = BN254.G1Point({
            X: readUint(operatorConfigJson, index, "RPoint.X"),
            Y: readUint(operatorConfigJson, index, "RPoint.Y")
        });
        return (s, pubkey);
    }

    function getOperatorSecretKey(uint256 index) public returns(uint256) {
        return readUint(operatorConfigJson, index, "SecretKey");
    }

    function getOperatorPubkeyG1(uint256 index) public returns(BN254.G1Point memory) {
        BN254.G1Point memory pubkey = BN254.G1Point({
            X: readUint(operatorConfigJson, index, "PubkeyG1.X"),
            Y: readUint(operatorConfigJson, index, "PubkeyG1.Y")
        });
        return pubkey;
    }

    function getOperatorPubkeyG2(uint256 index) public returns(BN254.G2Point memory) {
        BN254.G2Point memory pubkey = BN254.G2Point({
            X: [
                readUint(operatorConfigJson, index, "PubkeyG2.X.A1"),
                readUint(operatorConfigJson, index, "PubkeyG2.X.A0")
            ],   
            Y: [
                readUint(operatorConfigJson, index, "PubkeyG2.Y.A1"),
                readUint(operatorConfigJson, index, "PubkeyG2.Y.A0")
            ]
        }); 
        return pubkey;
    }

    function readUint(string memory json, uint256 index, string memory key) public returns (uint) {
        return stringToUint(stdJson.readString(json, string.concat(operatorPrefix(index), key)));
    }

    function stringToUint(string memory s) public pure returns (uint) {
        bytes memory b = bytes(s);
        uint256 result = 0;
        for (uint256 i = 0; i < b.length; i++) {
            if (uint256(uint8(b[i])) >= 48 && uint256(uint8(b[i])) <= 57) {
                result = result * 10 + (uint256(uint8(b[i])) - 48); 
            }
        }
        return result;
    }
    function setOperatorJsonFilePath(string memory filepath) public {
        operatorConfigJson = vm.readFile(filepath);
    }
}
