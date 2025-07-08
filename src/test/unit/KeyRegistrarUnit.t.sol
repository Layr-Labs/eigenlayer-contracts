//SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

import "forge-std/Test.sol";
import "src/contracts/permissions/KeyRegistrar.sol";
import "src/contracts/interfaces/IKeyRegistrar.sol";
import "src/contracts/libraries/BN254.sol";
import "src/contracts/libraries/BN254SignatureVerifier.sol";
import "src/contracts/interfaces/IPermissionController.sol";
import "src/contracts/mixins/PermissionControllerMixin.sol";
import "src/test/utils/EigenLayerMultichainUnitTestSetup.sol";
import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/interfaces/ISignatureUtilsMixin.sol";

contract KeyRegistrarUnitTests is
    EigenLayerMultichainUnitTestSetup,
    IKeyRegistrarErrors,
    IKeyRegistrarEvents,
    ISignatureUtilsMixinErrors
{
    using BN254 for BN254.G1Point;

    // Test accounts
    address public owner = address(0x1);
    address public operator1 = address(0x2);
    address public operator2 = address(0x3);
    address public avs1 = address(0x4);
    address public avs2 = address(0x5);

    uint32 public constant DEFAULT_OPERATOR_SET_ID = 0;

    // Test addresses for ECDSA (20 bytes each)
    address public ecdsaAddress1 = address(0x1234567890123456789012345678901234567890);
    address public ecdsaAddress2 = address(0x9876543210987654321098765432109876543210);
    bytes public ecdsaKey1 = abi.encodePacked(ecdsaAddress1);
    bytes public ecdsaKey2 = abi.encodePacked(ecdsaAddress2);

    // Private keys for ECDSA addresses (for signature generation)
    uint public ecdsaPrivKey1 = 0x1234567890123456789012345678901234567890123456789012345678901234;
    uint public ecdsaPrivKey2 = 0x9876543210987654321098765432109876543210987654321098765432109876;

    // Test keys for BN254
    uint public bn254PrivKey1 = 69;
    uint public bn254PrivKey2 = 123;

    BN254.G1Point bn254G1Key1;
    BN254.G1Point bn254G1Key2;
    BN254.G2Point bn254G2Key1;
    BN254.G2Point bn254G2Key2;

    bytes public bn254Key1;
    bytes public bn254Key2;

    function setUp() public virtual override {
        // The KeyRegistrar is deployed in the setUp of the EigenLayerMultichainUnitTestSetup
        EigenLayerMultichainUnitTestSetup.setUp();

        // Set up ECDSA addresses that correspond to the private keys
        ecdsaAddress1 = vm.addr(ecdsaPrivKey1);
        ecdsaAddress2 = vm.addr(ecdsaPrivKey2);
        ecdsaKey1 = abi.encodePacked(ecdsaAddress1);
        ecdsaKey2 = abi.encodePacked(ecdsaAddress2);

        // Set up BN254 keys with proper G2 components
        bn254G1Key1 = BN254.generatorG1().scalar_mul(bn254PrivKey1);
        bn254G1Key2 = BN254.generatorG1().scalar_mul(bn254PrivKey2);

        // Valid G2 points that correspond to the private keys
        bn254G2Key1.X[1] = 19_101_821_850_089_705_274_637_533_855_249_918_363_070_101_489_527_618_151_493_230_256_975_900_223_847;
        bn254G2Key1.X[0] = 5_334_410_886_741_819_556_325_359_147_377_682_006_012_228_123_419_628_681_352_847_439_302_316_235_957;
        bn254G2Key1.Y[1] = 354_176_189_041_917_478_648_604_979_334_478_067_325_821_134_838_555_150_300_539_079_146_482_658_331;
        bn254G2Key1.Y[0] = 4_185_483_097_059_047_421_902_184_823_581_361_466_320_657_066_600_218_863_748_375_739_772_335_928_910;

        bn254G2Key2.X[1] = 19_276_105_129_625_393_659_655_050_515_259_006_463_014_579_919_681_138_299_520_812_914_148_935_621_072;
        bn254G2Key2.X[0] = 14_066_454_060_412_929_535_985_836_631_817_650_877_381_034_334_390_275_410_072_431_082_437_297_539_867;
        bn254G2Key2.Y[1] = 12_642_665_914_920_339_463_975_152_321_804_664_028_480_770_144_655_934_937_445_922_690_262_428_344_269;
        bn254G2Key2.Y[0] = 10_109_651_107_942_685_361_120_988_628_892_759_706_059_655_669_161_016_107_907_096_760_613_704_453_218;

        bn254Key1 = abi.encode(bn254G1Key1.X, bn254G1Key1.Y, bn254G2Key1.X, bn254G2Key1.Y);
        bn254Key2 = abi.encode(bn254G1Key2.X, bn254G1Key2.Y, bn254G2Key2.X, bn254G2Key2.Y);

        allocationManagerMock.setAVSRegistrar(avs1, avs1);
        allocationManagerMock.setAVSRegistrar(avs2, avs2);
    }

    function _createOperatorSet(address avs, uint32 operatorSetId) internal pure returns (OperatorSet memory) {
        return OperatorSet({avs: avs, id: operatorSetId});
    }

    function _generateECDSASignature(address operator, OperatorSet memory operatorSet, address keyAddress, uint privKey)
        internal
        view
        returns (bytes memory)
    {
        bytes32 messageHash = keyRegistrar.getECDSAKeyRegistrationMessageHash(operator, operatorSet, keyAddress);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privKey, messageHash);
        return abi.encodePacked(r, s, v);
    }

    function _generateBN254Signature(address operator, OperatorSet memory operatorSet, bytes memory pubkey, uint privKey)
        internal
        view
        returns (bytes memory)
    {
        bytes32 messageHash = keyRegistrar.getBN254KeyRegistrationMessageHash(operator, operatorSet, pubkey);
        BN254.G1Point memory msgPoint = BN254.hashToG1(messageHash);
        BN254.G1Point memory signature = msgPoint.scalar_mul(privKey);
        return abi.encode(signature.X, signature.Y);
    }
}

/**
 * @title KeyRegistrarUnitTests_configureOperatorSet
 * @notice Unit tests for KeyRegistrar.configureOperatorSet
 */
contract KeyRegistrarUnitTests_configureOperatorSet is KeyRegistrarUnitTests {
    function test_revert_unauthorized() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(operator1);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);
    }

    function test_revert_invalidCurveType() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        vm.expectRevert(InvalidCurveType.selector);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.NONE);
    }

    function test_revert_configurationAlreadySet() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.startPrank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        // Try configuring the same curve type again
        vm.expectRevert(ConfigurationAlreadySet.selector);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);
        vm.stopPrank();
    }

    function test_configureOperatorSet_ECDSA() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        vm.expectEmit(true, true, true, true);
        emit OperatorSetConfigured(operatorSet, CurveType.ECDSA);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        CurveType curveType = keyRegistrar.getOperatorSetCurveType(operatorSet);
        assertEq(uint8(curveType), uint8(CurveType.ECDSA));
    }

    function test_configureOperatorSet_BN254() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        vm.expectEmit(true, true, true, true);
        emit OperatorSetConfigured(operatorSet, CurveType.BN254);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

        CurveType curveType = keyRegistrar.getOperatorSetCurveType(operatorSet);
        assertEq(uint8(curveType), uint8(CurveType.BN254));
    }

    function test_multipleOperatorSets() public {
        uint32 operatorSetId1 = 0;
        uint32 operatorSetId2 = 1;

        OperatorSet memory operatorSet1 = _createOperatorSet(avs1, operatorSetId1);
        OperatorSet memory operatorSet2 = _createOperatorSet(avs1, operatorSetId2);

        vm.startPrank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet1, CurveType.ECDSA);
        keyRegistrar.configureOperatorSet(operatorSet2, CurveType.BN254);
        vm.stopPrank();

        CurveType curveType1 = keyRegistrar.getOperatorSetCurveType(operatorSet1);
        CurveType curveType2 = keyRegistrar.getOperatorSetCurveType(operatorSet2);

        assertEq(uint8(curveType1), uint8(CurveType.ECDSA));
        assertEq(uint8(curveType2), uint8(CurveType.BN254));
    }
}

/**
 * @title KeyRegistrarUnitTests_registerKey_ECDSA
 * @notice Unit tests for KeyRegistrar.registerKey with ECDSA keys
 */
contract KeyRegistrarUnitTests_registerKey_ECDSA is KeyRegistrarUnitTests {
    function test_revert_invalidFormat() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        // Invalid length (not 20 bytes)
        bytes memory invalidKey = hex"04b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4";

        vm.prank(operator1);
        vm.expectRevert(InvalidKeyFormat.selector);
        keyRegistrar.registerKey(operator1, operatorSet, invalidKey, "");

        // Another invalid length
        bytes memory invalidKey2 =
            hex"b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c86a57d8a35ebfd17be4b8d44d0f7b8d5b5e5e5e5e5e5e5e5e5e5e5e5e5e5e5e5e5e5";

        vm.prank(operator1);
        vm.expectRevert(InvalidKeyFormat.selector);
        keyRegistrar.registerKey(operator1, operatorSet, invalidKey2, "");
    }

    function test_revert_zeroPubkey() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        bytes memory zeroKey = abi.encodePacked(address(0));

        vm.prank(operator1);
        vm.expectRevert(ZeroPubkey.selector);
        keyRegistrar.registerKey(operator1, operatorSet, zeroKey, "");
    }

    function test_revert_invalidSignature() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        // Use wrong private key for signature
        bytes memory wrongSignature = _generateECDSASignature(
            operator1,
            operatorSet,
            ecdsaAddress1,
            ecdsaPrivKey2 // Wrong private key
        );

        vm.prank(operator1);
        vm.expectRevert(InvalidSignature.selector);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, wrongSignature);
    }

    function test_revert_invalidSignature_malformed() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        // Malformed signature (wrong length)
        bytes memory malformedSignature = hex"deadbeef";

        vm.prank(operator1);
        vm.expectRevert(); // Will revert with ECDSA-specific error
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, malformedSignature);
    }

    function test_revert_invalidSignature_emptySignature() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        // Empty signature
        bytes memory emptySignature = "";

        vm.prank(operator1);
        vm.expectRevert(); // Will revert with ECDSA recovery error
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, emptySignature);
    }

    function test_revert_alreadyRegistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        bytes memory signature = _generateECDSASignature(operator1, operatorSet, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature);

        vm.prank(operator1);
        vm.expectRevert(KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature);
    }

    function test_revert_globallyRegistered() public {
        OperatorSet memory operatorSet1 = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        OperatorSet memory operatorSet2 = _createOperatorSet(avs2, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet1, CurveType.ECDSA);

        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(operatorSet2, CurveType.ECDSA);

        bytes memory signature1 = _generateECDSASignature(operator1, operatorSet1, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, ecdsaKey1, signature1);

        bytes memory signature2 = _generateECDSASignature(operator2, operatorSet2, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator2);
        vm.expectRevert(KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator2, operatorSet2, ecdsaKey1, signature2);
    }

    function test_revert_operatorSetNotConfigured() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(operator1);
        vm.expectRevert(OperatorSetNotConfigured.selector);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, "");
    }

    function test_revert_unauthorized() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        vm.prank(operator2);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, "");
    }

    function test_revert_wrongCurveType() public {
        // Configure for ECDSA but try to register BN254 key
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        bytes memory signature = _generateBN254Signature(operator1, operatorSet, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        vm.expectRevert(InvalidKeyFormat.selector);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);
    }

    function test_registerECDSAKey() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        bytes memory signature = _generateECDSASignature(operator1, operatorSet, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        vm.expectEmit(true, true, true, true);
        emit KeyRegistered(operatorSet, operator1, CurveType.ECDSA, ecdsaKey1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature);

        assertTrue(keyRegistrar.isRegistered(operatorSet, operator1));
        bytes memory storedKey = keyRegistrar.getECDSAKey(operatorSet, operator1);
        assertEq(storedKey, ecdsaKey1);

        address storedAddress = keyRegistrar.getECDSAAddress(operatorSet, operator1);
        assertEq(storedAddress, ecdsaAddress1);

        // Verify getOperatorFromSigningKey returns the correct operator and registration status
        (address retrievedOperator, bool isReg) = keyRegistrar.getOperatorFromSigningKey(operatorSet, ecdsaKey1);
        assertEq(retrievedOperator, operator1);
        assertTrue(isReg);
    }
}

/**
 * @title KeyRegistrarUnitTests_registerKey_BN254
 * @notice Unit tests for KeyRegistrar.registerKey with BN254 keys
 */
contract KeyRegistrarUnitTests_registerKey_BN254 is KeyRegistrarUnitTests {
    function test_revert_zeroPubkey() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

        bytes memory zeroKey = abi.encode(uint(0), uint(0), bn254G2Key1.X, bn254G2Key1.Y);
        bytes memory signature = _generateBN254Signature(operator1, operatorSet, zeroKey, bn254PrivKey1);

        vm.prank(operator1);
        vm.expectRevert(ZeroPubkey.selector);
        keyRegistrar.registerKey(operator1, operatorSet, zeroKey, signature);
    }

    function test_revert_invalidSignature() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

        bytes memory invalidSignature = abi.encode(uint(1), uint(2));

        vm.prank(operator1);
        vm.expectRevert(InvalidSignature.selector);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, invalidSignature);
    }

    function test_revert_invalidSignature_malformedData() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

        // Malformed signature - wrong format
        bytes memory malformedSignature = hex"deadbeef";

        vm.prank(operator1);
        vm.expectRevert(); // Will revert during abi.decode
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, malformedSignature);
    }

    function test_revert_invalidSignature_emptySignature() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

        // Empty signature
        bytes memory emptySignature = "";

        vm.prank(operator1);
        vm.expectRevert(); // Will revert during abi.decode
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, emptySignature);
    }

    function test_revert_alreadyRegistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

        bytes memory signature = _generateBN254Signature(operator1, operatorSet, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);

        // Try to register the same key again
        vm.prank(operator1);
        vm.expectRevert(KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);
    }

    function test_revert_globallyRegistered() public {
        OperatorSet memory operatorSet1 = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        OperatorSet memory operatorSet2 = _createOperatorSet(avs2, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet1, CurveType.BN254);

        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(operatorSet2, CurveType.BN254);

        bytes memory signature1 = _generateBN254Signature(operator1, operatorSet1, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, bn254Key1, signature1);

        bytes memory signature2 = _generateBN254Signature(operator2, operatorSet2, bn254Key1, bn254PrivKey1);

        vm.prank(operator2);
        vm.expectRevert(KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator2, operatorSet2, bn254Key1, signature2);
    }

    function test_registerBN254Key_wrongSignature() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

        // Use wrong private key for signature
        bytes memory wrongSignature = _generateBN254Signature(
            operator1,
            operatorSet,
            bn254Key1,
            bn254PrivKey2 // Wrong private key
        );

        vm.prank(operator1);
        vm.expectRevert(InvalidSignature.selector);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, wrongSignature);
    }

    function test_encodeBN254KeyData() public {
        // Test that encodeBN254KeyData produces the expected encoding
        bytes memory encodedKey = keyRegistrar.encodeBN254KeyData(bn254G1Key1, bn254G2Key1);

        // Verify the encoded data matches our manual encoding
        assertEq(encodedKey, bn254Key1);

        // Decode to verify structure
        (uint g1X, uint g1Y, uint[2] memory g2X, uint[2] memory g2Y) = abi.decode(encodedKey, (uint, uint, uint[2], uint[2]));

        assertEq(g1X, bn254G1Key1.X);
        assertEq(g1Y, bn254G1Key1.Y);
        assertEq(g2X[0], bn254G2Key1.X[0]);
        assertEq(g2X[1], bn254G2Key1.X[1]);
        assertEq(g2Y[0], bn254G2Key1.Y[0]);
        assertEq(g2Y[1], bn254G2Key1.Y[1]);
    }

    function test_encodeBN254KeyData_registerKey() public {
        // Test that encoded data from encodeBN254KeyData can be used with registerKey
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

        // Use encodeBN254KeyData to prepare the key data
        bytes memory encodedKey = keyRegistrar.encodeBN254KeyData(bn254G1Key2, bn254G2Key2);

        // Generate signature for the encoded key
        bytes memory signature = _generateBN254Signature(operator1, operatorSet, encodedKey, bn254PrivKey2);

        // Register the key using the encoded data
        vm.prank(operator1);
        vm.expectEmit(true, true, true, true);
        emit KeyRegistered(operatorSet, operator1, CurveType.BN254, encodedKey);
        keyRegistrar.registerKey(operator1, operatorSet, encodedKey, signature);

        // Verify registration was successful
        assertTrue(keyRegistrar.isRegistered(operatorSet, operator1));

        // Verify the stored key matches what we registered
        (BN254.G1Point memory storedG1, BN254.G2Point memory storedG2) = keyRegistrar.getBN254Key(operatorSet, operator1);
        assertEq(storedG1.X, bn254G1Key2.X);
        assertEq(storedG1.Y, bn254G1Key2.Y);
        assertEq(storedG2.X[0], bn254G2Key2.X[0]);
        assertEq(storedG2.X[1], bn254G2Key2.X[1]);
        assertEq(storedG2.Y[0], bn254G2Key2.Y[0]);
        assertEq(storedG2.Y[1], bn254G2Key2.Y[1]);

        // Verify getOperatorFromSigningKey returns the correct operator and registration status
        (address retrievedOperator, bool isReg) = keyRegistrar.getOperatorFromSigningKey(operatorSet, encodedKey);
        assertEq(retrievedOperator, operator1);
        assertTrue(isReg);
    }

    function test_registerBN254Key() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

        bytes memory signature = _generateBN254Signature(operator1, operatorSet, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        vm.expectEmit(true, true, true, true);
        emit KeyRegistered(operatorSet, operator1, CurveType.BN254, bn254Key1);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);

        assertTrue(keyRegistrar.isRegistered(operatorSet, operator1));
        (BN254.G1Point memory storedG1, BN254.G2Point memory storedG2) = keyRegistrar.getBN254Key(operatorSet, operator1);
        assertEq(storedG1.X, bn254G1Key1.X);
        assertEq(storedG1.Y, bn254G1Key1.Y);
        assertEq(storedG2.X[0], bn254G2Key1.X[0]);
        assertEq(storedG2.X[1], bn254G2Key1.X[1]);
        assertEq(storedG2.Y[0], bn254G2Key1.Y[0]);
        assertEq(storedG2.Y[1], bn254G2Key1.Y[1]);

        // Verify getOperatorFromSigningKey returns the correct operator and registration status
        (address retrievedOperator, bool isReg) = keyRegistrar.getOperatorFromSigningKey(operatorSet, bn254Key1);
        assertEq(retrievedOperator, operator1);
        assertTrue(isReg);
    }
}

/**
 * @title KeyRegistrarUnitTests_deregisterKey
 * @notice Unit tests for KeyRegistrar.deregisterKey
 */
contract KeyRegistrarUnitTests_deregisterKey is KeyRegistrarUnitTests {
    function test_revert_keyNotFound() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        vm.prank(operator1);
        vm.expectRevert(abi.encodeWithSelector(KeyNotFound.selector, operatorSet, operator1));
        keyRegistrar.deregisterKey(operator1, operatorSet);
    }

    function test_revert_unauthorized() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(operator2); // operator2 is not authorized to call on behalf of operator1
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        keyRegistrar.deregisterKey(operator1, operatorSet);
    }

    function test_revert_operatorStillSlashable() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        bytes memory signature = _generateECDSASignature(operator1, operatorSet, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature);

        // Set operator as slashable
        allocationManagerMock.setIsOperatorSlashable(operator1, operatorSet, true);

        // Operator should not be able to deregister their key when still slashable
        vm.prank(operator1);
        vm.expectRevert(abi.encodeWithSelector(OperatorStillSlashable.selector, operatorSet, operator1));
        keyRegistrar.deregisterKey(operator1, operatorSet);

        assertTrue(keyRegistrar.isRegistered(operatorSet, operator1));
    }

    function test_revert_operatorSetNotConfigured() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(operator1);
        vm.expectRevert(OperatorSetNotConfigured.selector);
        keyRegistrar.deregisterKey(operator1, operatorSet);
    }

    function test_deregisterKey_ECDSA() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        bytes memory signature = _generateECDSASignature(operator1, operatorSet, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature);

        // Set operator as not slashable
        allocationManagerMock.setIsOperatorSlashable(operator1, operatorSet, false);

        vm.prank(operator1);
        vm.expectEmit(true, true, true, true);
        emit KeyDeregistered(operatorSet, operator1, CurveType.ECDSA);
        keyRegistrar.deregisterKey(operator1, operatorSet);

        assertFalse(keyRegistrar.isRegistered(operatorSet, operator1));
    }

    function test_deregisterKey_BN254() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

        bytes memory signature = _generateBN254Signature(operator1, operatorSet, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);

        // Set operator as not slashable
        allocationManagerMock.setIsOperatorSlashable(operator1, operatorSet, false);

        vm.prank(operator1);
        vm.expectEmit(true, true, true, true);
        emit KeyDeregistered(operatorSet, operator1, CurveType.BN254);
        keyRegistrar.deregisterKey(operator1, operatorSet);

        assertFalse(keyRegistrar.isRegistered(operatorSet, operator1));
    }
}

/**
 * @title KeyRegistrarUnitTests_checkKey
 * @notice Unit tests for KeyRegistrar.checkKey
 */
contract KeyRegistrarUnitTests_checkKey is KeyRegistrarUnitTests {
    function test_checkKey_ECDSA() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        bytes memory signature = _generateECDSASignature(operator1, operatorSet, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature);

        bool hasKey = keyRegistrar.checkKey(operatorSet, operator1);
        assertTrue(hasKey);
    }

    function test_checkKey_BN254() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

        bytes memory signature = _generateBN254Signature(operator1, operatorSet, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);

        bool hasKey = keyRegistrar.checkKey(operatorSet, operator1);
        assertTrue(hasKey);
    }

    function test_checkKey_notRegistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        bool hasKey = keyRegistrar.checkKey(operatorSet, operator1);
        assertFalse(hasKey);
    }

    function test_revert_operatorSetNotConfigured() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.expectRevert(OperatorSetNotConfigured.selector);
        keyRegistrar.checkKey(operatorSet, operator1);
    }
}

/**
 * @title KeyRegistrarUnitTests_ViewFunctions
 * @notice Unit tests for KeyRegistrar view functions
 */
contract KeyRegistrarUnitTests_ViewFunctions is KeyRegistrarUnitTests {
    function test_getKeyHash_unregisteredOperator() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        bytes32 zeroHash = keyRegistrar.getKeyHash(operatorSet, operator2);
        assertEq(zeroHash, bytes32(0));
    }

    function test_getKeyHash_ECDSA() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        bytes memory signature = _generateECDSASignature(operator1, operatorSet, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature);

        bytes32 keyHash = keyRegistrar.getKeyHash(operatorSet, operator1);
        bytes32 expectedHash = keccak256(ecdsaKey1);
        assertEq(keyHash, expectedHash);
    }

    function test_getKeyHash_BN254() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

        bytes memory signature = _generateBN254Signature(operator1, operatorSet, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);

        bytes32 keyHash = keyRegistrar.getKeyHash(operatorSet, operator1);
        bytes32 expectedHash = BN254.hashG1Point(bn254G1Key1);
        assertEq(keyHash, expectedHash);
    }

    function test_isKeyGloballyRegistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        bytes32 keyHash = keccak256(ecdsaKey1);
        assertFalse(keyRegistrar.isKeyGloballyRegistered(keyHash));

        bytes memory signature = _generateECDSASignature(operator1, operatorSet, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature);

        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));
    }

    function test_getBN254Key_emptyForUnregistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

        (BN254.G1Point memory g1Point, BN254.G2Point memory g2Point) = keyRegistrar.getBN254Key(operatorSet, operator1);

        assertEq(g1Point.X, 0);
        assertEq(g1Point.Y, 0);
        assertEq(g2Point.X[0], 0);
        assertEq(g2Point.X[1], 0);
        assertEq(g2Point.Y[0], 0);
        assertEq(g2Point.Y[1], 0);
    }

    function test_getECDSAKey_emptyForUnregistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        bytes memory emptyKey = keyRegistrar.getECDSAKey(operatorSet, operator1);
        assertEq(emptyKey.length, 0);
    }

    function test_getECDSAAddress_zeroForUnregistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        address zeroAddress = keyRegistrar.getECDSAAddress(operatorSet, operator1);
        assertEq(zeroAddress, address(0));
    }

    function test_getECDSAKey_revertWrongCurveType() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

        vm.expectRevert(InvalidCurveType.selector);
        keyRegistrar.getECDSAKey(operatorSet, operator1);
    }

    function test_getBN254Key_revertWrongCurveType() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        vm.expectRevert(InvalidCurveType.selector);
        keyRegistrar.getBN254Key(operatorSet, operator1);
    }

    function test_version() public {
        string memory version = keyRegistrar.version();
        assertEq(version, "9.9.9");
    }

    function test_getOperatorSetCurveType_unconfigured() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        // Should return NONE for unconfigured operator set
        CurveType curveType = keyRegistrar.getOperatorSetCurveType(operatorSet);
        assertEq(uint8(curveType), uint8(CurveType.NONE));
    }

    function test_getOperatorFromSigningKey_ECDSA() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        // Before registration, should return address(0) and false
        (address retrievedOperator, bool isReg) = keyRegistrar.getOperatorFromSigningKey(operatorSet, ecdsaKey1);
        assertEq(retrievedOperator, address(0));
        assertFalse(isReg);

        // Register key
        bytes memory signature = _generateECDSASignature(operator1, operatorSet, ecdsaAddress1, ecdsaPrivKey1);
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature);

        // After registration, should return the operator and true
        (retrievedOperator, isReg) = keyRegistrar.getOperatorFromSigningKey(operatorSet, ecdsaKey1);
        assertEq(retrievedOperator, operator1);
        assertTrue(isReg);
    }

    function test_getOperatorFromSigningKey_BN254() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.BN254);

        // Before registration, should return address(0) and false
        (address retrievedOperator, bool isReg) = keyRegistrar.getOperatorFromSigningKey(operatorSet, bn254Key1);
        assertEq(retrievedOperator, address(0));
        assertFalse(isReg);

        // Register key
        bytes memory signature = _generateBN254Signature(operator1, operatorSet, bn254Key1, bn254PrivKey1);
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);

        // After registration, should return the operator and true
        // Only pass in the G1 key
        bytes memory g1Key = abi.encode(bn254G1Key1.X, bn254G1Key1.Y);
        (retrievedOperator, isReg) = keyRegistrar.getOperatorFromSigningKey(operatorSet, g1Key);
        assertEq(retrievedOperator, operator1);
        assertTrue(isReg);
    }

    function test_getOperatorFromSigningKey_multipleOperators() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        // Register different keys for different operators
        bytes memory signature1 = _generateECDSASignature(operator1, operatorSet, ecdsaAddress1, ecdsaPrivKey1);
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature1);

        bytes memory signature2 = _generateECDSASignature(operator2, operatorSet, ecdsaAddress2, ecdsaPrivKey2);
        vm.prank(operator2);
        keyRegistrar.registerKey(operator2, operatorSet, ecdsaKey2, signature2);

        // Verify each key returns the correct operator and registration status
        (address retrievedOperator1, bool isReg1) = keyRegistrar.getOperatorFromSigningKey(operatorSet, ecdsaKey1);
        assertEq(retrievedOperator1, operator1);
        assertTrue(isReg1);

        (address retrievedOperator2, bool isReg2) = keyRegistrar.getOperatorFromSigningKey(operatorSet, ecdsaKey2);
        assertEq(retrievedOperator2, operator2);
        assertTrue(isReg2);
    }

    function test_getOperatorFromSigningKey_sameKeyDifferentOperatorSets() public {
        OperatorSet memory operatorSet1 = _createOperatorSet(avs1, 0);
        OperatorSet memory operatorSet2 = _createOperatorSet(avs1, 1);

        vm.startPrank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet1, CurveType.ECDSA);
        keyRegistrar.configureOperatorSet(operatorSet2, CurveType.ECDSA);
        vm.stopPrank();

        // Register same operator with same key in first operator set
        bytes memory signature = _generateECDSASignature(operator1, operatorSet1, ecdsaAddress1, ecdsaPrivKey1);
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, ecdsaKey1, signature);

        // The same key should return the same operator for both operator sets
        // But registration status will differ - true for operatorSet1, false for operatorSet2
        (address retrievedOperator1, bool isReg1) = keyRegistrar.getOperatorFromSigningKey(operatorSet1, ecdsaKey1);
        assertEq(retrievedOperator1, operator1);
        assertTrue(isReg1); // Registered in operatorSet1

        (address retrievedOperator2, bool isReg2) = keyRegistrar.getOperatorFromSigningKey(operatorSet2, ecdsaKey1);
        assertEq(retrievedOperator2, operator1);
        assertFalse(isReg2); // NOT registered in operatorSet2
    }

    function test_getOperatorFromSigningKey_afterDeregistration() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        // Register key
        bytes memory signature = _generateECDSASignature(operator1, operatorSet, ecdsaAddress1, ecdsaPrivKey1);
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature);

        // Verify registration
        (address retrievedOperator, bool isReg) = keyRegistrar.getOperatorFromSigningKey(operatorSet, ecdsaKey1);
        assertEq(retrievedOperator, operator1);
        assertTrue(isReg);

        // Deregister
        allocationManagerMock.setIsOperatorSlashable(operator1, operatorSet, false);
        vm.prank(operator1);
        keyRegistrar.deregisterKey(operator1, operatorSet);

        // After deregistration, the key should still map to the operator (global registry persists)
        // but the registration status should be false
        (retrievedOperator, isReg) = keyRegistrar.getOperatorFromSigningKey(operatorSet, ecdsaKey1);
        assertEq(retrievedOperator, operator1);
        assertFalse(isReg);
    }

    function test_getOperatorFromSigningKey_nonExistentKey() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, CurveType.ECDSA);

        // Query for a key that was never registered
        bytes memory nonExistentKey = abi.encodePacked(address(0xdeadbeef));
        (address retrievedOperator, bool isReg) = keyRegistrar.getOperatorFromSigningKey(operatorSet, nonExistentKey);
        assertEq(retrievedOperator, address(0));
        assertFalse(isReg);
    }

    function test_getOperatorFromSigningKey_revertUnconfiguredOperatorSet() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        // Don't configure the operator set - it will have CurveType.NONE
        bytes memory someKey = abi.encodePacked(address(0xdeadbeef));

        // This should revert because the operator set is not configured
        vm.expectRevert();
        keyRegistrar.getOperatorFromSigningKey(operatorSet, someKey);
    }
}

/**
 * @title KeyRegistrarUnitTests_MultipleOperatorSets
 * @notice Unit tests for KeyRegistrar with multiple operator sets
 */
contract KeyRegistrarUnitTests_MultipleOperatorSets is KeyRegistrarUnitTests {
    function test_multipleOperatorSets_differentKeyTypes() public {
        uint32 ecdsaSetId = 0;
        uint32 bn254SetId = 1;

        OperatorSet memory ecdsaOperatorSet = _createOperatorSet(avs1, ecdsaSetId);
        OperatorSet memory bn254OperatorSet = _createOperatorSet(avs1, bn254SetId);

        vm.startPrank(avs1);
        keyRegistrar.configureOperatorSet(ecdsaOperatorSet, CurveType.ECDSA);
        keyRegistrar.configureOperatorSet(bn254OperatorSet, CurveType.BN254);
        vm.stopPrank();

        // Register ECDSA key for one operator set
        bytes memory ecdsaSignature = _generateECDSASignature(operator1, ecdsaOperatorSet, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, ecdsaOperatorSet, ecdsaKey1, ecdsaSignature);

        // Register BN254 key for another operator set
        bytes memory bn254Signature = _generateBN254Signature(operator1, bn254OperatorSet, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, bn254OperatorSet, bn254Key1, bn254Signature);

        // Verify both registrations
        assertTrue(keyRegistrar.isRegistered(ecdsaOperatorSet, operator1));
        assertTrue(keyRegistrar.isRegistered(bn254OperatorSet, operator1));

        // Verify key retrieval
        bytes memory ecdsaKey = keyRegistrar.getECDSAKey(ecdsaOperatorSet, operator1);
        assertEq(ecdsaKey, ecdsaKey1);

        address ecdsaAddr = keyRegistrar.getECDSAAddress(ecdsaOperatorSet, operator1);
        assertEq(ecdsaAddr, ecdsaAddress1);

        (BN254.G1Point memory g1Point,) = keyRegistrar.getBN254Key(bn254OperatorSet, operator1);
        assertEq(g1Point.X, bn254G1Key1.X);
        assertEq(g1Point.Y, bn254G1Key1.Y);
    }

    function test_multipleOperatorsRegistration() public {
        uint32 operatorSetId1 = 0;
        uint32 operatorSetId2 = 1;

        OperatorSet memory operatorSet1 = _createOperatorSet(avs1, operatorSetId1);
        OperatorSet memory operatorSet2 = _createOperatorSet(avs1, operatorSetId2);

        vm.startPrank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet1, CurveType.ECDSA);
        keyRegistrar.configureOperatorSet(operatorSet2, CurveType.BN254);
        vm.stopPrank();

        CurveType curveType1 = keyRegistrar.getOperatorSetCurveType(operatorSet1);
        CurveType curveType2 = keyRegistrar.getOperatorSetCurveType(operatorSet2);

        assertEq(uint8(curveType1), uint8(CurveType.ECDSA));
        assertEq(uint8(curveType2), uint8(CurveType.BN254));

        bytes memory signature = _generateECDSASignature(operator1, operatorSet1, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, ecdsaKey1, signature);

        assertTrue(keyRegistrar.isRegistered(operatorSet1, operator1));
        assertFalse(keyRegistrar.isRegistered(operatorSet2, operator1));
    }
}

/**
 * @title KeyRegistrarUnitTests_GlobalKeyPersistence
 * @notice Unit tests for global key persistence in KeyRegistrar
 */
contract KeyRegistrarUnitTests_GlobalKeyPersistence is KeyRegistrarUnitTests {
    function test_globalKeyPersistence_ECDSA() public {
        OperatorSet memory operatorSet1 = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        OperatorSet memory operatorSet2 = _createOperatorSet(avs2, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet1, CurveType.ECDSA);

        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(operatorSet2, CurveType.ECDSA);

        bytes memory signature1 = _generateECDSASignature(operator1, operatorSet1, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, ecdsaKey1, signature1);

        bytes32 keyHash = keccak256(ecdsaKey1);
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));

        // Set operator as not slashable
        allocationManagerMock.setIsOperatorSlashable(operator1, operatorSet1, false);

        vm.prank(operator1);
        keyRegistrar.deregisterKey(operator1, operatorSet1);

        // Key should still be globally registered after deregistration
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));

        bytes memory signature2 = _generateECDSASignature(operator2, operatorSet2, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator2);
        vm.expectRevert(KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator2, operatorSet2, ecdsaKey1, signature2);
    }

    function test_globalKeyPersistence_BN254() public {
        OperatorSet memory operatorSet1 = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        OperatorSet memory operatorSet2 = _createOperatorSet(avs2, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet1, CurveType.BN254);

        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(operatorSet2, CurveType.BN254);

        bytes memory signature1 = _generateBN254Signature(operator1, operatorSet1, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, bn254Key1, signature1);

        bytes32 keyHash = BN254.hashG1Point(bn254G1Key1);
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));

        // Set operator as not slashable
        allocationManagerMock.setIsOperatorSlashable(operator1, operatorSet1, false);

        vm.prank(operator1);
        keyRegistrar.deregisterKey(operator1, operatorSet1);

        // Key should still be globally registered after deregistration
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));

        bytes memory signature2 = _generateBN254Signature(operator2, operatorSet2, bn254Key1, bn254PrivKey1);

        vm.prank(operator2);
        vm.expectRevert(KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator2, operatorSet2, bn254Key1, signature2);
    }
}

/**
 * @title KeyRegistrarUnitTests_CrossCurveInteraction
 * @notice Unit tests for cross-curve type interactions in KeyRegistrar
 */
contract KeyRegistrarUnitTests_CrossCurveInteraction is KeyRegistrarUnitTests {
    function test_crossCurveGlobalUniqueness() public {
        // Configure ECDSA and BN254 operator sets
        OperatorSet memory ecdsaOperatorSet = _createOperatorSet(avs1, 0);
        OperatorSet memory bn254OperatorSet = _createOperatorSet(avs1, 1);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(ecdsaOperatorSet, CurveType.ECDSA);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(bn254OperatorSet, CurveType.BN254);

        // Register ECDSA key
        bytes memory ecdsaSignature = _generateECDSASignature(operator1, ecdsaOperatorSet, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, ecdsaOperatorSet, ecdsaKey1, ecdsaSignature);

        // Register BN254 key (should succeed as they have different hashes)
        bytes memory bn254Signature = _generateBN254Signature(operator1, bn254OperatorSet, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, bn254OperatorSet, bn254Key1, bn254Signature);

        // Both should be registered
        assertTrue(keyRegistrar.isRegistered(ecdsaOperatorSet, operator1));
        assertTrue(keyRegistrar.isRegistered(bn254OperatorSet, operator1));
    }
}

/**
 * @title KeyRegistrarUnitTests_SignatureVerification
 * @notice Unit tests for signature verification in KeyRegistrar
 */
contract KeyRegistrarUnitTests_SignatureVerification is KeyRegistrarUnitTests {
    using BN254 for BN254.G1Point;

    function test_verifyBN254Signature() public {
        bytes32 messageHash = keccak256("test message");

        // Generate signature with private key
        BN254.G1Point memory msgPoint = BN254.hashToG1(messageHash);
        BN254.G1Point memory signature = msgPoint.scalar_mul(bn254PrivKey2);

        (bool success, bool pairingSuccessful) =
            BN254SignatureVerifier.verifySignature(messageHash, signature, bn254G1Key2, bn254G2Key2, false, 0);
        assertTrue(success && pairingSuccessful);
    }

    function test_verifyBN254Signature_revertInvalid() public {
        bytes32 messageHash = keccak256("test message");

        // Use a signature generated with the wrong private key - this should fail verification
        BN254.G1Point memory msgPoint = BN254.hashToG1(messageHash);
        BN254.G1Point memory invalidSignature = msgPoint.scalar_mul(bn254PrivKey1); // Wrong private key

        (bool success, bool pairingSuccessful) =
            BN254SignatureVerifier.verifySignature(messageHash, invalidSignature, bn254G1Key2, bn254G2Key2, false, 0);
        assertFalse(success && pairingSuccessful);
    }
}
