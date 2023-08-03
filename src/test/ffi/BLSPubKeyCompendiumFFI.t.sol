// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "forge-std/Test.sol";
import "../../contracts/middleware/BLSPublicKeyCompendium.sol";
import "openzeppelin-contracts/contracts/utils/Strings.sol";

contract BLSPublicKeyCompendiumFFITests is Test {
    using BN254 for BN254.G1Point;
    using Strings for uint256;

    BLSPublicKeyCompendium compendium;

    uint256 privKey;
    BN254.G1Point pubKeyG1;
    BN254.G2Point pubKeyG2;
    BN254.G1Point signedMessageHash;

    address alice = address(0x69);

    function setUp() public {
        compendium = new BLSPublicKeyCompendium();
    }

    function testRegisterBLSPublicKey(uint256 _privKey) public {
        vm.assume(_privKey != 0);

        _setKeys(_privKey);
        signedMessageHash = _signMessage(alice);

        vm.prank(alice);
        compendium.registerBLSPublicKey(signedMessageHash, pubKeyG1, pubKeyG2);

        assertEq(compendium.operatorToPubkeyHash(alice), BN254.hashG1Point(pubKeyG1), "pubkey hash not stored correctly");
        assertEq(compendium.pubkeyHashToOperator(BN254.hashG1Point(pubKeyG1)), alice, "operator address not stored correctly");
    }

    function _setKeys(uint256 _privKey) internal {
        privKey = _privKey;
        pubKeyG1 = BN254.generatorG1().scalar_mul(_privKey);

        string[] memory inputs = new string[](5);
        inputs[0] = "go";
        inputs[1] = "run";
        inputs[2] = "src/test/ffi/g2pubkey.go";
        inputs[3] = _privKey.toString(); 

        inputs[4] = "1";
        bytes memory res = vm.ffi(inputs);
        pubKeyG2.X[1] = abi.decode(res, (uint256));

        inputs[4] = "2";
        res = vm.ffi(inputs);
        pubKeyG2.X[0] = abi.decode(res, (uint256));

        inputs[4] = "3";
        res = vm.ffi(inputs);
        pubKeyG2.Y[1] = abi.decode(res, (uint256));

        inputs[4] = "4";
        res = vm.ffi(inputs);
        pubKeyG2.Y[0] = abi.decode(res, (uint256));
    }

    function _signMessage(address signer) internal view returns(BN254.G1Point memory) {
        BN254.G1Point memory messageHash = BN254.hashToG1(keccak256(abi.encodePacked(signer, block.chainid, "EigenLayer_BN254_Pubkey_Registration")));
        return BN254.scalar_mul(messageHash, privKey);
    }
}