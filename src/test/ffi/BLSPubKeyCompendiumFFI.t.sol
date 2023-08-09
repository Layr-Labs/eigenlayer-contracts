// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/middleware/BLSPublicKeyCompendium.sol";
import "./util/G2Operations.sol";

contract BLSPublicKeyCompendiumFFITests is G2Operations {
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

    function testRegisterBLSPublicKey(/*uint256 _privKey*/) public {
        // _setKeys(_privKey);

        _setKeys(666);
        signedMessageHash = _signMessage(alice);

        vm.prank(alice);
        compendium.registerBLSPublicKey(signedMessageHash, pubKeyG1, pubKeyG2);

        assertEq(compendium.operatorToPubkeyHash(alice), BN254.hashG1Point(pubKeyG1), "pubkey hash not stored correctly");
        assertEq(compendium.pubkeyHashToOperator(BN254.hashG1Point(pubKeyG1)), alice, "operator address not stored correctly");
    }

    function _setKeys(uint256 _privKey) internal {
        privKey = _privKey;
        pubKeyG1 = BN254.generatorG1().scalar_mul(_privKey);
        pubKeyG2 = G2Operations.mul(_privKey);
    }

    function _signMessage(address signer) internal view returns(BN254.G1Point memory) {
        BN254.G1Point memory messageHash = BN254.hashToG1(keccak256(abi.encodePacked(signer, block.chainid, "EigenLayer_BN254_Pubkey_Registration")));
        return BN254.scalar_mul(messageHash, privKey);
    }
}