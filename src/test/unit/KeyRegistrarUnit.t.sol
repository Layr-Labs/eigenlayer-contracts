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
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/interfaces/ISignatureUtilsMixin.sol";

contract KeyRegistrarUnitTests is EigenLayerUnitTestSetup {
    using BN254 for BN254.G1Point;

    KeyRegistrar public keyRegistrar;

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

    event KeyRegistered(OperatorSet operatorSet, address indexed operator, IKeyRegistrarTypes.CurveType curveType, bytes pubkey);
    event KeyDeregistered(OperatorSet operatorSet, address indexed operator, IKeyRegistrarTypes.CurveType curveType);
    event OperatorSetConfigured(OperatorSet operatorSet, IKeyRegistrarTypes.CurveType curveType);

    function setUp() public virtual override {
        EigenLayerUnitTestSetup.setUp();

        keyRegistrar = new KeyRegistrar(
            IPermissionController(address(permissionController)), IAllocationManager(address(allocationManagerMock)), "1.0.0"
        );

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
        bytes32 structHash =
            keccak256(abi.encode(keyRegistrar.ECDSA_KEY_REGISTRATION_TYPEHASH(), operator, operatorSet.avs, operatorSet.id, keyAddress));
        bytes32 messageHash = keyRegistrar.domainSeparator();
        messageHash = keccak256(abi.encodePacked("\x19\x01", messageHash, structHash));

        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privKey, messageHash);
        return abi.encodePacked(r, s, v);
    }

    function _generateBN254Signature(address operator, OperatorSet memory operatorSet, bytes memory pubkey, uint privKey)
        internal
        view
        returns (bytes memory)
    {
        bytes32 structHash = keccak256(
            abi.encode(keyRegistrar.BN254_KEY_REGISTRATION_TYPEHASH(), operator, operatorSet.avs, operatorSet.id, keccak256(pubkey))
        );
        bytes32 messageHash = keyRegistrar.domainSeparator();
        messageHash = keccak256(abi.encodePacked("\x19\x01", messageHash, structHash));

        BN254.G1Point memory msgPoint = BN254.hashToG1(messageHash);
        BN254.G1Point memory signature = msgPoint.scalar_mul(privKey);

        return abi.encode(signature.X, signature.Y);
    }

    // ============ Operator Set Configuration Tests ============

    function testConfigureOperatorSet() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        vm.expectEmit(true, true, true, true);
        emit OperatorSetConfigured(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        IKeyRegistrarTypes.CurveType curveType = keyRegistrar.getOperatorSetCurveType(operatorSet);
        assertEq(uint8(curveType), uint8(IKeyRegistrarTypes.CurveType.ECDSA));
    }

    function testConfigureOperatorSet_RevertUnauthorized() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(operator1);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
    }

    function testConfigureOperatorSet_BN254() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        vm.expectEmit(true, true, true, true);
        emit OperatorSetConfigured(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);

        IKeyRegistrarTypes.CurveType curveType = keyRegistrar.getOperatorSetCurveType(operatorSet);
        assertEq(uint8(curveType), uint8(IKeyRegistrarTypes.CurveType.BN254));
    }

    function testConfigureOperatorSet_RevertConfigurationAlreadySet() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        vm.prank(avs1);
        vm.expectRevert(IKeyRegistrarErrors.ConfigurationAlreadySet.selector);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
    }

    function testConfigureOperatorSet_RevertInvalidCurveType() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        vm.expectRevert(IKeyRegistrarErrors.InvalidCurveType.selector);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.NONE);
    }

    // ============ ECDSA Key Registration Tests ============

    function testRegisterECDSAKey() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        bytes memory signature = _generateECDSASignature(operator1, operatorSet, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        vm.expectEmit(true, true, true, true);
        emit KeyRegistered(operatorSet, operator1, IKeyRegistrarTypes.CurveType.ECDSA, ecdsaKey1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature);

        assertTrue(keyRegistrar.isRegistered(operatorSet, operator1));
        bytes memory storedKey = keyRegistrar.getECDSAKey(operatorSet, operator1);
        assertEq(storedKey, ecdsaKey1);

        address storedAddress = keyRegistrar.getECDSAAddress(operatorSet, operator1);
        assertEq(storedAddress, ecdsaAddress1);
    }

    function testRegisterECDSAKey_RevertInvalidFormat() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        // Invalid length (not 20 bytes)
        bytes memory invalidKey = hex"04b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4";

        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.InvalidKeyFormat.selector);
        keyRegistrar.registerKey(operator1, operatorSet, invalidKey, "");

        // Another invalid length
        bytes memory invalidKey2 =
            hex"b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c86a57d8a35ebfd17be4b8d44d0f7b8d5b5e5e5e5e5e5e5e5e5e5e5e5e5e5e5e5e5e5";

        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.InvalidKeyFormat.selector);
        keyRegistrar.registerKey(operator1, operatorSet, invalidKey2, "");
    }

    function testRegisterECDSAKey_RevertZeroPubkey() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        bytes memory zeroKey = abi.encodePacked(address(0));

        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.ZeroPubkey.selector);
        keyRegistrar.registerKey(operator1, operatorSet, zeroKey, "");
    }

    function testRegisterECDSAKey_RevertInvalidSignature() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        // Use wrong private key for signature
        bytes memory wrongSignature = _generateECDSASignature(
            operator1,
            operatorSet,
            ecdsaAddress1,
            ecdsaPrivKey2 // Wrong private key
        );

        vm.prank(operator1);
        vm.expectRevert(ISignatureUtilsMixinErrors.InvalidSignature.selector);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, wrongSignature);
    }

    function testRegisterECDSAKey_RevertAlreadyRegistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        bytes memory signature = _generateECDSASignature(operator1, operatorSet, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature);

        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature);
    }

    function testRegisterECDSAKey_RevertGloballyRegistered() public {
        OperatorSet memory operatorSet1 = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        OperatorSet memory operatorSet2 = _createOperatorSet(avs2, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet1, IKeyRegistrarTypes.CurveType.ECDSA);

        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(operatorSet2, IKeyRegistrarTypes.CurveType.ECDSA);

        bytes memory signature1 = _generateECDSASignature(operator1, operatorSet1, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, ecdsaKey1, signature1);

        bytes memory signature2 = _generateECDSASignature(operator2, operatorSet2, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator2);
        vm.expectRevert(IKeyRegistrarErrors.KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator2, operatorSet2, ecdsaKey1, signature2);
    }

    // ============ BN254 Key Registration Tests ============

    function testRegisterBN254Key() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);

        bytes memory signature = _generateBN254Signature(operator1, operatorSet, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        vm.expectEmit(true, true, true, true);
        emit KeyRegistered(operatorSet, operator1, IKeyRegistrarTypes.CurveType.BN254, bn254Key1);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);

        assertTrue(keyRegistrar.isRegistered(operatorSet, operator1));
        (BN254.G1Point memory storedG1, BN254.G2Point memory storedG2) = keyRegistrar.getBN254Key(operatorSet, operator1);
        assertEq(storedG1.X, bn254G1Key1.X);
        assertEq(storedG1.Y, bn254G1Key1.Y);
    }

    function testRegisterBN254Key_RevertZeroPubkey() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);

        bytes memory zeroKey = abi.encode(uint(0), uint(0), bn254G2Key1.X, bn254G2Key1.Y);
        bytes memory signature = _generateBN254Signature(operator1, operatorSet, zeroKey, bn254PrivKey1);

        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.ZeroPubkey.selector);
        keyRegistrar.registerKey(operator1, operatorSet, zeroKey, signature);
    }

    function testRegisterBN254Key_RevertInvalidSignature() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);

        bytes memory invalidSignature = abi.encode(uint(1), uint(2));

        vm.prank(operator1);
        vm.expectRevert(ISignatureUtilsMixinErrors.InvalidSignature.selector);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, invalidSignature);
    }

    function testRegisterBN254Key_RevertAlreadyRegistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);

        bytes memory signature = _generateBN254Signature(operator1, operatorSet, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);

        bytes32 keyHash = keyRegistrar.getKeyHash(operatorSet, operator1);
        bytes32 expectedHash = BN254.hashG1Point(bn254G1Key1);
        assertEq(keyHash, expectedHash);
    }

    function testGetKeyHash_UnregisteredOperator() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        bytes32 zeroHash = keyRegistrar.getKeyHash(operatorSet, operator2);
        assertEq(zeroHash, bytes32(0));
    }

    function testIsKeyGloballyRegistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        bytes32 keyHash = keccak256(ecdsaKey1);
        assertFalse(keyRegistrar.isKeyGloballyRegistered(keyHash));

        bytes memory signature = _generateECDSASignature(operator1, operatorSet, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature);

        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));
    }

    function testGetBN254Key_EmptyForUnregistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);

        (BN254.G1Point memory g1Point, BN254.G2Point memory g2Point) = keyRegistrar.getBN254Key(operatorSet, operator1);

        assertEq(g1Point.X, 0);
        assertEq(g1Point.Y, 0);
        assertEq(g2Point.X[0], 0);
        assertEq(g2Point.X[1], 0);
        assertEq(g2Point.Y[0], 0);
        assertEq(g2Point.Y[1], 0);
    }

    function testGetECDSAKey_EmptyForUnregistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        bytes memory emptyKey = keyRegistrar.getECDSAKey(operatorSet, operator1);
        assertEq(emptyKey.length, 0);
    }

    function testGetECDSAAddress_ZeroForUnregistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        address zeroAddress = keyRegistrar.getECDSAAddress(operatorSet, operator1);
        assertEq(zeroAddress, address(0));
    }

    // ============ Authorization Tests ============

    function testRegisterKey_RevertOperatorSetNotConfigured() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.OperatorSetNotConfigured.selector);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, "");
    }

    function testRegisterKey_RevertUnauthorized() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        vm.prank(operator2);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, "");
    }

    function testDeregisterKey_RevertOperatorSetNotConfigured() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        vm.expectRevert(IKeyRegistrarErrors.OperatorSetNotConfigured.selector);
        keyRegistrar.deregisterKey(operator1, operatorSet);
    }

    function testCheckKey_RevertOperatorSetNotConfigured() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        vm.expectRevert(IKeyRegistrarErrors.OperatorSetNotConfigured.selector);
        keyRegistrar.checkKey(operatorSet, operator1);
    }

    // ============ Multiple Operator Sets Tests ============

    function testMultipleOperatorSets() public {
        uint32 operatorSetId1 = 0;
        uint32 operatorSetId2 = 1;

        OperatorSet memory operatorSet1 = _createOperatorSet(avs1, operatorSetId1);
        OperatorSet memory operatorSet2 = _createOperatorSet(avs1, operatorSetId2);

        vm.startPrank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet1, IKeyRegistrarTypes.CurveType.ECDSA);
        keyRegistrar.configureOperatorSet(operatorSet2, IKeyRegistrarTypes.CurveType.BN254);
        vm.stopPrank();

        IKeyRegistrarTypes.CurveType curveType1 = keyRegistrar.getOperatorSetCurveType(operatorSet1);
        IKeyRegistrarTypes.CurveType curveType2 = keyRegistrar.getOperatorSetCurveType(operatorSet2);

        assertEq(uint8(curveType1), uint8(IKeyRegistrarTypes.CurveType.ECDSA));
        assertEq(uint8(curveType2), uint8(IKeyRegistrarTypes.CurveType.BN254));

        bytes memory signature = _generateECDSASignature(operator1, operatorSet1, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, ecdsaKey1, signature);

        assertTrue(keyRegistrar.isRegistered(operatorSet1, operator1));
        assertFalse(keyRegistrar.isRegistered(operatorSet2, operator1));
    }

    function testMultipleOperatorSets_DifferentKeyTypes() public {
        uint32 ecdsaSetId = 0;
        uint32 bn254SetId = 1;

        OperatorSet memory ecdsaOperatorSet = _createOperatorSet(avs1, ecdsaSetId);
        OperatorSet memory bn254OperatorSet = _createOperatorSet(avs1, bn254SetId);

        vm.startPrank(avs1);
        keyRegistrar.configureOperatorSet(ecdsaOperatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
        keyRegistrar.configureOperatorSet(bn254OperatorSet, IKeyRegistrarTypes.CurveType.BN254);
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

    // ============ Global Key Persistence Tests ============

    function testGlobalKeyPersistence() public {
        OperatorSet memory operatorSet1 = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        OperatorSet memory operatorSet2 = _createOperatorSet(avs2, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet1, IKeyRegistrarTypes.CurveType.ECDSA);

        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(operatorSet2, IKeyRegistrarTypes.CurveType.ECDSA);

        bytes memory signature1 = _generateECDSASignature(operator1, operatorSet1, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, ecdsaKey1, signature1);

        bytes32 keyHash = keccak256(ecdsaKey1);
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));

        vm.prank(avs1);
        keyRegistrar.deregisterKey(operator1, operatorSet1);

        // Key should still be globally registered after deregistration
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));

        bytes memory signature2 = _generateECDSASignature(operator2, operatorSet2, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator2);
        vm.expectRevert(IKeyRegistrarErrors.KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator2, operatorSet2, ecdsaKey1, signature2);
    }

    function testGlobalKeyPersistence_BN254() public {
        OperatorSet memory operatorSet1 = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        OperatorSet memory operatorSet2 = _createOperatorSet(avs2, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet1, IKeyRegistrarTypes.CurveType.BN254);

        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(operatorSet2, IKeyRegistrarTypes.CurveType.BN254);

        bytes memory signature1 = _generateBN254Signature(operator1, operatorSet1, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, bn254Key1, signature1);

        bytes32 keyHash = BN254.hashG1Point(bn254G1Key1);
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));

        vm.prank(avs1);
        keyRegistrar.deregisterKey(operator1, operatorSet1);

        // Key should still be globally registered after deregistration
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));

        bytes memory signature2 = _generateBN254Signature(operator2, operatorSet2, bn254Key1, bn254PrivKey1);

        vm.prank(operator2);
        vm.expectRevert(IKeyRegistrarErrors.KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator2, operatorSet2, bn254Key1, signature2);
    }

    // ============ Cross-Curve Type Tests ============

    function testCrossCurveGlobalUniqueness() public {
        // Configure ECDSA and BN254 operator sets
        OperatorSet memory ecdsaOperatorSet = _createOperatorSet(avs1, 0);
        OperatorSet memory bn254OperatorSet = _createOperatorSet(avs1, 1);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(ecdsaOperatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(bn254OperatorSet, IKeyRegistrarTypes.CurveType.BN254);

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

    // ============ Error Condition Tests ============

    function testRegisterKey_RevertWrongCurveType() public {
        // Configure for ECDSA but try to register BN254 key
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        bytes memory signature = _generateBN254Signature(operator1, operatorSet, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.InvalidKeyFormat.selector);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);
    }

    function testGetECDSAKey_RevertWrongCurveType() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);

        vm.expectRevert(IKeyRegistrarErrors.InvalidCurveType.selector);
        keyRegistrar.getECDSAKey(operatorSet, operator1);
    }

    function testGetBN254Key_RevertWrongCurveType() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        vm.expectRevert(IKeyRegistrarErrors.InvalidCurveType.selector);
        keyRegistrar.getBN254Key(operatorSet, operator1);
    }

    // ============ Version Tests ============

    function testVersion() public {
        string memory version = keyRegistrar.version();
        assertEq(version, "1.0.0");
    }

    // ============ Signature Verification Tests ============

    function testVerifyBN254Signature() public {
        bytes32 messageHash = keccak256("test message");

        // Generate signature with private key
        BN254.G1Point memory msgPoint = BN254.hashToG1(messageHash);
        BN254.G1Point memory signature = msgPoint.scalar_mul(bn254PrivKey2);

        (bool success, bool pairingSuccessful) =
            BN254SignatureVerifier.verifySignature(messageHash, signature, bn254G1Key2, bn254G2Key2, false, 0);
        assertTrue(success && pairingSuccessful);
    }

    function testVerifyBN254Signature_RevertInvalid() public {
        bytes32 messageHash = keccak256("test message");

        // Use a signature generated with the wrong private key - this should fail verification
        BN254.G1Point memory msgPoint = BN254.hashToG1(messageHash);
        BN254.G1Point memory invalidSignature = msgPoint.scalar_mul(bn254PrivKey1); // Wrong private key

        (bool success, bool pairingSuccessful) =
            BN254SignatureVerifier.verifySignature(messageHash, invalidSignature, bn254G1Key2, bn254G2Key2, false, 0);
        assertFalse(success && pairingSuccessful);
    }

    function testRegisterBN254Key_RevertGloballyRegistered() public {
        OperatorSet memory operatorSet1 = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        OperatorSet memory operatorSet2 = _createOperatorSet(avs2, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet1, IKeyRegistrarTypes.CurveType.BN254);

        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(operatorSet2, IKeyRegistrarTypes.CurveType.BN254);

        bytes memory signature1 = _generateBN254Signature(operator1, operatorSet1, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, bn254Key1, signature1);

        bytes memory signature2 = _generateBN254Signature(operator2, operatorSet2, bn254Key1, bn254PrivKey1);

        vm.prank(operator2);
        vm.expectRevert(IKeyRegistrarErrors.KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator2, operatorSet2, bn254Key1, signature2);
    }

    function testRegisterBN254Key_WrongSignature() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);

        // Use wrong private key for signature
        bytes memory wrongSignature = _generateBN254Signature(
            operator1,
            operatorSet,
            bn254Key1,
            bn254PrivKey2 // Wrong private key
        );

        vm.prank(operator1);
        vm.expectRevert(ISignatureUtilsMixinErrors.InvalidSignature.selector);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, wrongSignature);
    }

    // ============ Key Deregistration Tests ============

    function testDeregisterKey_ECDSA() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        bytes memory signature = _generateECDSASignature(operator1, operatorSet, ecdsaAddress1, ecdsaPrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, signature);

        vm.prank(avs1);
        vm.expectEmit(true, true, true, true);
        emit KeyDeregistered(operatorSet, operator1, IKeyRegistrarTypes.CurveType.ECDSA);
        keyRegistrar.deregisterKey(operator1, operatorSet);

        assertFalse(keyRegistrar.isRegistered(operatorSet, operator1));
    }

    function testDeregisterKey_BN254() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);

        bytes memory signature = _generateBN254Signature(operator1, operatorSet, bn254Key1, bn254PrivKey1);

        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);

        vm.prank(avs1);
        vm.expectEmit(true, true, true, true);
        emit KeyDeregistered(operatorSet, operator1, IKeyRegistrarTypes.CurveType.BN254);
        keyRegistrar.deregisterKey(operator1, operatorSet);

        assertFalse(keyRegistrar.isRegistered(operatorSet, operator1));
    }

    function testDeregisterKey_RevertKeyNotFound() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        vm.prank(avs1);
        vm.expectRevert(abi.encodeWithSelector(IKeyRegistrarErrors.KeyNotFound.selector, operatorSet, operator1));
        keyRegistrar.deregisterKey(operator1, operatorSet);
    }

    function testDeregisterKey_RevertUnauthorized() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);

        vm.prank(operator1);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        keyRegistrar.deregisterKey(operator1, operatorSet);
    }
}
