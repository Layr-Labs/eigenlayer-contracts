// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "forge-std/Test.sol";
import "../../contracts/middleware/BLSPublicKeyCompendium.sol";

contract BLSPublicKeyCompendiumUnitTests is Test {
    using BN254 for BN254.G1Point;

    BLSPublicKeyCompendium compendium;
    
    uint256 privKey = 69;
    
    BN254.G1Point pubKeyG1;
    BN254.G2Point pubKeyG2;
    BN254.G1Point signedMessageHash;

    address alice = address(1);
    address bob = address(2);

    function setUp() public {
        compendium = new BLSPublicKeyCompendium();

        pubKeyG1 = BN254.generatorG1().scalar_mul(privKey);
        
        //privKey*G2
        pubKeyG2.X[1] = 19101821850089705274637533855249918363070101489527618151493230256975900223847;
        pubKeyG2.X[0] = 5334410886741819556325359147377682006012228123419628681352847439302316235957;
        pubKeyG2.Y[1] = 354176189041917478648604979334478067325821134838555150300539079146482658331;
        pubKeyG2.Y[0] = 4185483097059047421902184823581361466320657066600218863748375739772335928910;
    }

    function testRegisterBLSPublicKey() public {
        signedMessageHash = _signMessage(alice);
        vm.prank(alice);
        compendium.registerBLSPublicKey(signedMessageHash, pubKeyG1, pubKeyG2);

        assertEq(compendium.operatorToPubkeyHash(alice), BN254.hashG1Point(pubKeyG1), "pubkey hash not stored correctly");
        assertEq(compendium.pubkeyHashToOperator(BN254.hashG1Point(pubKeyG1)), alice, "operator address not stored correctly");
    }

    function testRegisterBLSPublicKey_NoMatch_Reverts() public {
        signedMessageHash = _signMessage(alice);
        BN254.G1Point memory badPubKeyG1 = BN254.generatorG1().scalar_mul(420); // mismatch public keys

        vm.prank(alice);
        vm.expectRevert(bytes("BLSPublicKeyCompendium.registerBLSPublicKey: G1 and G2 private key do not match"));
        compendium.registerBLSPublicKey(signedMessageHash, badPubKeyG1, pubKeyG2);
    }

    function testRegisterBLSPublicKey_BadSig_Reverts() public {
        signedMessageHash = _signMessage(bob); // sign with wrong private key

        vm.prank(alice); 
        vm.expectRevert(bytes("BLSPublicKeyCompendium.registerBLSPublicKey: G1 and G2 private key do not match"));
        compendium.registerBLSPublicKey(signedMessageHash, pubKeyG1, pubKeyG2);
    }

    function testRegisterBLSPublicKey_OpRegistered_Reverts() public {
        testRegisterBLSPublicKey(); // register alice

        vm.prank(alice); 
        vm.expectRevert(bytes("BLSPublicKeyCompendium.registerBLSPublicKey: operator already registered pubkey"));
        compendium.registerBLSPublicKey(signedMessageHash, pubKeyG1, pubKeyG2);
    }

    function testRegisterBLSPublicKey_PkRegistered_Reverts() public {
        testRegisterBLSPublicKey(); 
        signedMessageHash = _signMessage(bob); // same private key different operator

        vm.prank(bob); 
        vm.expectRevert(bytes("BLSPublicKeyCompendium.registerBLSPublicKey: public key already registered"));
        compendium.registerBLSPublicKey(signedMessageHash, pubKeyG1, pubKeyG2);
    }

    function _signMessage(address signer) internal view returns(BN254.G1Point memory) {
        BN254.G1Point memory messageHash = BN254.hashToG1(keccak256(abi.encodePacked(signer, block.chainid, "EigenLayer_BN254_Pubkey_Registration")));
        return BN254.scalar_mul(messageHash, privKey);
    }

}