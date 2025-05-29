//SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "../../contracts/core/KeyRegistrar.sol";
import {BN254} from "../../contracts/libraries/BN254.sol";
import {IPermissionController} from "../../contracts/interfaces/IPermissionController.sol";
import {PermissionControllerMixin} from "../../contracts/mixins/PermissionControllerMixin.sol";

contract MockPermissionController is IPermissionController {
    function addPendingAdmin(address, address) external override {}
    function removePendingAdmin(address, address) external override {}
    function acceptAdmin(address) external override {}
    function removeAdmin(address, address) external override {}
    function setAppointee(address, address, address, bytes4) external override {}
    function removeAppointee(address, address, address, bytes4) external override {}
    
    function isAdmin(address, address) external pure override returns (bool) {
        return true;
    }
    
    function isPendingAdmin(address, address) external pure override returns (bool) {
        return false;
    }
    
    function getAdmins(address account) external pure override returns (address[] memory) {
        address[] memory admins = new address[](1);
        admins[0] = account;
        return admins;
    }
    
    function getPendingAdmins(address) external pure override returns (address[] memory) {
        return new address[](0);
    }
    
    function canCall(address account, address caller, address target, bytes4 selector) external override returns (bool) {
        return caller == account;
    }
    function getAppointeePermissions(address, address) external pure override returns (address[] memory, bytes4[] memory) {
        return (new address[](0), new bytes4[](0));
    }
    
    function getAppointees(address, address, bytes4) external pure override returns (address[] memory) {
        return new address[](0);
    }
    
    function version() external pure override returns (string memory) {
        return "1.0.0";
    }
}

contract KeyRegistrarUnitTests is Test {
    using BN254 for BN254.G1Point;

    KeyRegistrar public keyRegistrar;
    MockPermissionController public permissionController;
    
    address public owner = address(0x1);
    address public operator1 = address(0x2);
    address public operator2 = address(0x3);
    address public avs1 = address(0x4);
    address public avs2 = address(0x5);
    
    uint32 public constant DEFAULT_OPERATOR_SET_ID = 0;
    uint256 public constant DEFAULT_ROTATION_DELAY = 100;
    
    // Test keys for ECDSA
    bytes public ecdsaKey1 = hex"04b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c86a57d8a35ebfd17be4b8d44d0f7b8d5b5e5e5e5e5e5e5e5e5e5e5e5e5e5e5e5";
    bytes public ecdsaKey2 = hex"04c6cc8d9d0251a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4955d96a68e8a36fbef18ce5c9e55e1f8c9e6c6f6f6f6f6f6f6f6f6f6f6f6f6f6f6";
    
    // Test keys for BN254
    uint256 public bn254PrivKey1 = 69;
    uint256 public bn254PrivKey2 = 123;
    
    BN254.G1Point bn254G1Key1;
    BN254.G1Point bn254G1Key2;
    BN254.G2Point bn254G2Key1;
    BN254.G2Point bn254G2Key2;
    
    bytes public bn254Key1;
    bytes public bn254Key2;

    event KeyRegistered(address indexed avs, uint32 indexed operatorSetId, address indexed operator, KeyRegistrar.CurveType curveType, bytes pubkey);
    event KeyDeregistered(address indexed avs, uint32 indexed operatorSetId, address indexed operator, KeyRegistrar.CurveType curveType);
    event KeyRotationInitiated(address indexed avs, uint32 indexed operatorSetId, address indexed operator, KeyRegistrar.CurveType curveType, bytes oldKey, bytes newKey, uint256 effectiveBlock);
    event AggregateBN254KeyUpdated(address indexed avs, uint32 indexed operatorSetId, BN254.G1Point newAggregateKey);
    event OperatorSetConfigured(address indexed avs, uint32 indexed operatorSetId, KeyRegistrar.CurveType curveType, uint256 rotationDelay);

    function setUp() public {
        permissionController = new MockPermissionController();
        keyRegistrar = new KeyRegistrar(IPermissionController(address(permissionController)));
        
        // Set up BN254 keys with proper G2 components
        bn254G1Key1 = BN254.generatorG1().scalar_mul(bn254PrivKey1);
        bn254G1Key2 = BN254.generatorG1().scalar_mul(bn254PrivKey2);
        
        // Valid G2 points that correspond to the private keys
        bn254G2Key1.X[1] = 19101821850089705274637533855249918363070101489527618151493230256975900223847;
        bn254G2Key1.X[0] = 5334410886741819556325359147377682006012228123419628681352847439302316235957;
        bn254G2Key1.Y[1] = 354176189041917478648604979334478067325821134838555150300539079146482658331;
        bn254G2Key1.Y[0] = 4185483097059047421902184823581361466320657066600218863748375739772335928910;

        bn254G2Key2.X[1] = 19276105129625393659655050515259006463014579919681138299520812914148935621072;
        bn254G2Key2.X[0] = 14066454060412929535985836631817650877381034334390275410072431082437297539867;
        bn254G2Key2.Y[1] = 12642665914920339463975152321804664028480770144655934937445922690262428344269;
        bn254G2Key2.Y[0] = 10109651107942685361120988628892759706059655669161016107907096760613704453218;
        
        bn254Key1 = abi.encode(bn254G1Key1.X, bn254G1Key1.Y, bn254G2Key1.X, bn254G2Key1.Y);
        bn254Key2 = abi.encode(bn254G1Key2.X, bn254G1Key2.Y, bn254G2Key2.X, bn254G2Key2.Y);
    }

    function _generateBN254Signature(
        address operator,
        address avs,
        uint32 operatorSetId,
        bytes memory pubkey,
        uint256 privKey
    ) internal view returns (bytes memory) {
        bytes32 messageHash = keccak256(abi.encodePacked(
            "EigenLayer.KeyRegistrar.v1",
            address(keyRegistrar),
            avs, 
            operatorSetId, 
            operator, 
            pubkey
        ));
        
        BN254.G1Point memory msgPoint = BN254.hashToG1(messageHash);
        BN254.G1Point memory signature = msgPoint.scalar_mul(privKey);
        
        return abi.encode(signature.X, signature.Y);
    }

    // ============ Operator Set Configuration Tests ============

    function testConfigureOperatorSet() public {
        vm.prank(avs1);
        vm.expectEmit(true, true, true, true);
        emit OperatorSetConfigured(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        KeyRegistrar.OperatorSetConfig memory config = keyRegistrar.getOperatorSetConfig(avs1, DEFAULT_OPERATOR_SET_ID);
        assertEq(uint8(config.curveType), uint8(KeyRegistrar.CurveType.ECDSA));
        assertTrue(config.isActive);
        assertEq(config.rotationDelay, DEFAULT_ROTATION_DELAY);
    }

    function testConfigureOperatorSet_RevertUnauthorized() public {
        vm.prank(operator1);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
    }

    function testConfigureOperatorSet_BN254() public {
        vm.prank(avs1);
        vm.expectEmit(true, true, true, true);
        emit OperatorSetConfigured(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        KeyRegistrar.OperatorSetConfig memory config = keyRegistrar.getOperatorSetConfig(avs1, DEFAULT_OPERATOR_SET_ID);
        assertEq(uint8(config.curveType), uint8(KeyRegistrar.CurveType.BN254));
        assertTrue(config.isActive);
        assertEq(config.rotationDelay, DEFAULT_ROTATION_DELAY);
    }

    // ============ ECDSA Key Registration Tests ============

    function testRegisterECDSAKey() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator1);
        vm.expectEmit(true, true, true, true);
        emit KeyRegistered(avs1, DEFAULT_OPERATOR_SET_ID, operator1, KeyRegistrar.CurveType.ECDSA, ecdsaKey1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey1, "");
        
        assertTrue(keyRegistrar.isRegistered(avs1, DEFAULT_OPERATOR_SET_ID, operator1));
        bytes memory storedKey = keyRegistrar.getECDSAKey(avs1, DEFAULT_OPERATOR_SET_ID, operator1);
        assertEq(storedKey, ecdsaKey1);
    }

    function testRegisterECDSAKey_RevertInvalidFormat() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        bytes memory invalidKey = hex"04b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4";
        
        vm.prank(operator1);
        vm.expectRevert(KeyRegistrar.InvalidKeyFormat.selector);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, invalidKey, "");
        
        bytes memory invalidPrefix = hex"03b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c86a57d8a35ebfd17be4b8d44d0f7b8d5b5e5e5e5e5e5e5e5e5e5e5e5e5e5e5e5e5";
        
        vm.prank(operator1);
        vm.expectRevert(KeyRegistrar.InvalidKeyFormat.selector);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, invalidPrefix, "");
    }

    function testRegisterECDSAKey_RevertZeroPubkey() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        bytes memory zeroKey = new bytes(65);
        zeroKey[0] = 0x04;
        
        vm.prank(operator1);
        vm.expectRevert(KeyRegistrar.ZeroPubkey.selector);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, zeroKey, "");
    }

    function testRegisterECDSAKey_RevertAlreadyRegistered() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey1, "");
        
        vm.prank(operator1);
        vm.expectRevert(KeyRegistrar.KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey1, "");
    }

    function testRegisterECDSAKey_RevertGloballyRegistered() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(avs2, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey1, "");
        
        vm.prank(operator2);
        vm.expectRevert(KeyRegistrar.KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator2, avs2, DEFAULT_OPERATOR_SET_ID, ecdsaKey1, "");
    }

    // ============ BN254 Key Registration Tests ============

    function testRegisterBN254Key() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        bytes memory signature = _generateBN254Signature(
            operator1, 
            avs1, 
            DEFAULT_OPERATOR_SET_ID, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        vm.expectEmit(true, true, true, true);
        emit KeyRegistered(avs1, DEFAULT_OPERATOR_SET_ID, operator1, KeyRegistrar.CurveType.BN254, bn254Key1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, bn254Key1, signature);
        
        assertTrue(keyRegistrar.isRegistered(avs1, DEFAULT_OPERATOR_SET_ID, operator1));
        (BN254.G1Point memory storedG1, BN254.G2Point memory storedG2) = keyRegistrar.getBN254Key(avs1, DEFAULT_OPERATOR_SET_ID, operator1);
        assertEq(storedG1.X, bn254G1Key1.X);
        assertEq(storedG1.Y, bn254G1Key1.Y);
    }

    function testRegisterBN254Key_RevertZeroPubkey() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        bytes memory zeroKey = abi.encode(uint256(0), uint256(0), bn254G2Key1.X, bn254G2Key1.Y);
        bytes memory signature = _generateBN254Signature(
            operator1, 
            avs1, 
            DEFAULT_OPERATOR_SET_ID, 
            zeroKey, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        vm.expectRevert(KeyRegistrar.ZeroPubkey.selector);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, zeroKey, signature);
    }

    function testRegisterBN254Key_RevertInvalidSignature() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        bytes memory invalidSignature = abi.encode(uint256(1), uint256(2));
        
        vm.prank(operator1);
        vm.expectRevert(KeyRegistrar.InvalidSignature.selector);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, bn254Key1, invalidSignature);
    }

    function testRegisterBN254Key_RevertAlreadyRegistered() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        bytes memory signature = _generateBN254Signature(
            operator1, 
            avs1, 
            DEFAULT_OPERATOR_SET_ID, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, bn254Key1, signature);
        
        vm.prank(operator1);
        vm.expectRevert(KeyRegistrar.KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, bn254Key1, signature);
    }

    function testRegisterBN254Key_RevertGloballyRegistered() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(avs2, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        bytes memory signature1 = _generateBN254Signature(
            operator1, 
            avs1, 
            DEFAULT_OPERATOR_SET_ID, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, bn254Key1, signature1);
        
        bytes memory signature2 = _generateBN254Signature(
            operator2, 
            avs2, 
            DEFAULT_OPERATOR_SET_ID, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator2);
        vm.expectRevert(KeyRegistrar.KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator2, avs2, DEFAULT_OPERATOR_SET_ID, bn254Key1, signature2);
    }

    function testRegisterBN254Key_WrongSignature() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        // Use wrong private key for signature
        bytes memory wrongSignature = _generateBN254Signature(
            operator1, 
            avs1, 
            DEFAULT_OPERATOR_SET_ID, 
            bn254Key1, 
            bn254PrivKey2  // Wrong private key
        );
        
        vm.prank(operator1);
        vm.expectRevert(KeyRegistrar.InvalidSignature.selector);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, bn254Key1, wrongSignature);
    }

    // ============ Key Deregistration Tests ============

    function testDeregisterECDSAKey() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey1, "");
        
        vm.prank(operator1);
        vm.expectEmit(true, true, true, true);
        emit KeyDeregistered(avs1, DEFAULT_OPERATOR_SET_ID, operator1, KeyRegistrar.CurveType.ECDSA);
        keyRegistrar.deregisterKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID);
        
        assertFalse(keyRegistrar.isRegistered(avs1, DEFAULT_OPERATOR_SET_ID, operator1));
    }

    function testDeregisterBN254Key() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        bytes memory signature = _generateBN254Signature(
            operator1, 
            avs1, 
            DEFAULT_OPERATOR_SET_ID, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, bn254Key1, signature);
        
        vm.prank(operator1);
        vm.expectEmit(true, true, true, true);
        emit KeyDeregistered(avs1, DEFAULT_OPERATOR_SET_ID, operator1, KeyRegistrar.CurveType.BN254);
        keyRegistrar.deregisterKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID);
        
        assertFalse(keyRegistrar.isRegistered(avs1, DEFAULT_OPERATOR_SET_ID, operator1));
    }

    function testDeregisterKey_SilentSuccessWhenNotRegistered() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator1);
        keyRegistrar.deregisterKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID);
        
        assertFalse(keyRegistrar.isRegistered(avs1, DEFAULT_OPERATOR_SET_ID, operator1));
    }

    // ============ Key Rotation Tests ============

    function testRotateECDSAKey() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey1, "");
        
        vm.roll(block.number + DEFAULT_ROTATION_DELAY + 1);
        
        vm.prank(operator1);
        vm.expectEmit(true, true, true, true);
        emit KeyRotationInitiated(avs1, DEFAULT_OPERATOR_SET_ID, operator1, KeyRegistrar.CurveType.ECDSA, ecdsaKey1, ecdsaKey2, block.number + DEFAULT_ROTATION_DELAY);
        keyRegistrar.rotateKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey2, "");
        
        bytes memory storedKey = keyRegistrar.getECDSAKey(avs1, DEFAULT_OPERATOR_SET_ID, operator1);
        assertEq(storedKey, ecdsaKey2);
    }

    function testRotateBN254Key() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        bytes memory signature1 = _generateBN254Signature(
            operator1, 
            avs1, 
            DEFAULT_OPERATOR_SET_ID, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, bn254Key1, signature1);
        
        vm.roll(block.number + DEFAULT_ROTATION_DELAY + 1);
        
        bytes memory signature2 = _generateBN254Signature(
            operator1, 
            avs1, 
            DEFAULT_OPERATOR_SET_ID, 
            bn254Key2, 
            bn254PrivKey2
        );
        
        vm.prank(operator1);
        vm.expectEmit(true, true, true, true);
        emit KeyRotationInitiated(avs1, DEFAULT_OPERATOR_SET_ID, operator1, KeyRegistrar.CurveType.BN254, bn254Key1, bn254Key2, block.number + DEFAULT_ROTATION_DELAY);
        keyRegistrar.rotateKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, bn254Key2, signature2);
        
        (BN254.G1Point memory storedG1,) = keyRegistrar.getBN254Key(avs1, DEFAULT_OPERATOR_SET_ID, operator1);
        assertEq(storedG1.X, bn254G1Key2.X);
        assertEq(storedG1.Y, bn254G1Key2.Y);
    }

    function testRotateBN254Key_RevertInvalidSignature() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        bytes memory signature1 = _generateBN254Signature(
            operator1, 
            avs1, 
            DEFAULT_OPERATOR_SET_ID, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, bn254Key1, signature1);
        
        vm.roll(block.number + DEFAULT_ROTATION_DELAY + 1);
        
        bytes memory invalidSignature = abi.encode(uint256(1), uint256(2));
        
        vm.prank(operator1);
        vm.expectRevert(KeyRegistrar.InvalidSignature.selector);
        keyRegistrar.rotateKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, bn254Key2, invalidSignature);
    }

    function testRotateKey_RevertNotRegistered() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator1);
        vm.expectRevert(KeyRegistrar.KeyNotRegistered.selector);
        keyRegistrar.rotateKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey2, "");
    }

    function testRotateKey_RevertRotationTooSoon() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey1, "");
        
        vm.prank(operator1);
        vm.expectRevert(abi.encodeWithSelector(KeyRegistrar.RotationTooSoon.selector, block.number, DEFAULT_ROTATION_DELAY));
        keyRegistrar.rotateKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey2, "");
    }

    function testRotationAtExactBoundary() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);

        uint256 registrationBlock = block.number;
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey1, "");
        
        vm.roll(registrationBlock + DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator1);
        vm.expectRevert(abi.encodeWithSelector(KeyRegistrar.RotationTooSoon.selector, registrationBlock, DEFAULT_ROTATION_DELAY));
        keyRegistrar.rotateKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey2, "");
        
        vm.roll(registrationBlock + DEFAULT_ROTATION_DELAY + 1);
        
        vm.prank(operator1);
        keyRegistrar.rotateKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey2, "");
        
        bytes memory storedKey = keyRegistrar.getECDSAKey(avs1, DEFAULT_OPERATOR_SET_ID, operator1);
        assertEq(storedKey, ecdsaKey2);
    }

    // ============ APK Management Tests ============

    function testUpdateAPK_OnlyAuthorized() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator1);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        keyRegistrar.updateAPK(avs1, DEFAULT_OPERATOR_SET_ID, operator1, false);
    }

    function testUpdateAPK_AuthorizedCall() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);

        bytes memory signature = _generateBN254Signature(
            operator1, 
            avs1, 
            DEFAULT_OPERATOR_SET_ID, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, bn254Key1, signature);
        
        vm.prank(avs1);
        keyRegistrar.updateAPK(avs1, DEFAULT_OPERATOR_SET_ID, operator1, false);
    }

    function testUpdateAPK_ECDSAType_NoOp() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(avs1);
        keyRegistrar.updateAPK(avs1, DEFAULT_OPERATOR_SET_ID, operator1, false);
    }

    function testRemoveFromAPK_OnlyAuthorized() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator1);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        keyRegistrar.removeFromAPK(avs1, DEFAULT_OPERATOR_SET_ID, operator1);
    }

    function testRemoveFromAPK_AuthorizedCall() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        vm.prank(avs1);
        keyRegistrar.removeFromAPK(avs1, DEFAULT_OPERATOR_SET_ID, operator1);
    }

    // ============ View Function Tests ============

    function testGetOperatorSetConfig() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        KeyRegistrar.OperatorSetConfig memory config = keyRegistrar.getOperatorSetConfig(avs1, DEFAULT_OPERATOR_SET_ID);
        assertEq(uint8(config.curveType), uint8(KeyRegistrar.CurveType.ECDSA));
        assertTrue(config.isActive);
        assertEq(config.rotationDelay, DEFAULT_ROTATION_DELAY);
    }

    function testGetKeyHash_ECDSA() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey1, "");
        
        bytes32 keyHash = keyRegistrar.getKeyHash(avs1, DEFAULT_OPERATOR_SET_ID, operator1);
        bytes32 expectedHash = keccak256(ecdsaKey1);
        assertEq(keyHash, expectedHash);
    }

    function testGetKeyHash_BN254() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        bytes memory signature = _generateBN254Signature(
            operator1, 
            avs1, 
            DEFAULT_OPERATOR_SET_ID, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, bn254Key1, signature);
        
        bytes32 keyHash = keyRegistrar.getKeyHash(avs1, DEFAULT_OPERATOR_SET_ID, operator1);
        bytes32 expectedHash = BN254.hashG1Point(bn254G1Key1);
        assertEq(keyHash, expectedHash);
    }

    function testGetKeyHash_UnregisteredOperator() public {
        bytes32 zeroHash = keyRegistrar.getKeyHash(avs1, DEFAULT_OPERATOR_SET_ID, operator2);
        assertEq(zeroHash, bytes32(0));
    }

    function testIsKeyGloballyRegistered() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        bytes32 keyHash = keccak256(ecdsaKey1);
        assertFalse(keyRegistrar.isKeyGloballyRegistered(keyHash));
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey1, "");
        
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));
    }

    function testGetBN254Key_EmptyForUnregistered() public {
        (BN254.G1Point memory g1Point, BN254.G2Point memory g2Point) = keyRegistrar.getBN254Key(avs1, DEFAULT_OPERATOR_SET_ID, operator1);
        
        assertEq(g1Point.X, 0);
        assertEq(g1Point.Y, 0);
        assertEq(g2Point.X[0], 0);
        assertEq(g2Point.X[1], 0);
        assertEq(g2Point.Y[0], 0);
        assertEq(g2Point.Y[1], 0);
    }

    function testGetECDSAKey_EmptyForUnregistered() public {
        bytes memory key = keyRegistrar.getECDSAKey(avs1, DEFAULT_OPERATOR_SET_ID, operator1);
        assertEq(key.length, 0);
    }

    function testGetApk_ReturnsZeroForUninitialized() public {
        BN254.G1Point memory apk = keyRegistrar.getApk(avs1, DEFAULT_OPERATOR_SET_ID);
        assertEq(apk.X, 0);
        assertEq(apk.Y, 0);
    }

    // ============ Authorization Tests ============

    function testRegisterKey_RevertOperatorSetNotConfigured() public {
        vm.prank(operator1);
        vm.expectRevert(KeyRegistrar.OperatorSetNotConfigured.selector);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey1, "");
    }

    function testRegisterKey_RevertUnauthorized() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator2);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey1, "");
    }

    function testDeregisterKey_RevertUnauthorized() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator2);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        keyRegistrar.deregisterKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID);
    }

    function testRotateKey_RevertUnauthorized() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator2);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        keyRegistrar.rotateKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey2, "");
    }

    // ============ Multiple Operator Sets Tests ============

    function testMultipleOperatorSets() public {
        uint32 operatorSetId1 = 0;
        uint32 operatorSetId2 = 1;
        
        vm.startPrank(avs1);
        keyRegistrar.configureOperatorSet(avs1, operatorSetId1, KeyRegistrar.CurveType.ECDSA, 100);
        keyRegistrar.configureOperatorSet(avs1, operatorSetId2, KeyRegistrar.CurveType.BN254, 200);
        vm.stopPrank();
        
        KeyRegistrar.OperatorSetConfig memory config1 = keyRegistrar.getOperatorSetConfig(avs1, operatorSetId1);
        KeyRegistrar.OperatorSetConfig memory config2 = keyRegistrar.getOperatorSetConfig(avs1, operatorSetId2);
        
        assertEq(uint8(config1.curveType), uint8(KeyRegistrar.CurveType.ECDSA));
        assertEq(config1.rotationDelay, 100);
        
        assertEq(uint8(config2.curveType), uint8(KeyRegistrar.CurveType.BN254));
        assertEq(config2.rotationDelay, 200);
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, operatorSetId1, ecdsaKey1, "");
        
        assertTrue(keyRegistrar.isRegistered(avs1, operatorSetId1, operator1));
        assertFalse(keyRegistrar.isRegistered(avs1, operatorSetId2, operator1));
    }

    function testMultipleOperatorSets_DifferentKeyTypes() public {
        uint32 ecdsaSetId = 0;
        uint32 bn254SetId = 1;
        
        vm.startPrank(avs1);
        keyRegistrar.configureOperatorSet(avs1, ecdsaSetId, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        keyRegistrar.configureOperatorSet(avs1, bn254SetId, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        vm.stopPrank();
        
        // Register ECDSA key for one operator set
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, ecdsaSetId, ecdsaKey1, "");
        
        // Register BN254 key for another operator set
        bytes memory signature = _generateBN254Signature(
            operator1, 
            avs1, 
            bn254SetId, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, bn254SetId, bn254Key1, signature);
        
        // Verify both registrations
        assertTrue(keyRegistrar.isRegistered(avs1, ecdsaSetId, operator1));
        assertTrue(keyRegistrar.isRegistered(avs1, bn254SetId, operator1));
        
        // Verify key retrieval
        bytes memory ecdsaKey = keyRegistrar.getECDSAKey(avs1, ecdsaSetId, operator1);
        assertEq(ecdsaKey, ecdsaKey1);
        
        (BN254.G1Point memory g1Point,) = keyRegistrar.getBN254Key(avs1, bn254SetId, operator1);
        assertEq(g1Point.X, bn254G1Key1.X);
        assertEq(g1Point.Y, bn254G1Key1.Y);
    }

    // ============ Global Key Persistence Tests ============

    function testGlobalKeyPersistence() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(avs2, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, ecdsaKey1, "");
        
        bytes32 keyHash = keccak256(ecdsaKey1);
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));
        
        vm.prank(operator1);
        keyRegistrar.deregisterKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID);
        
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));
        
        vm.prank(operator2);
        vm.expectRevert(KeyRegistrar.KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator2, avs2, DEFAULT_OPERATOR_SET_ID, ecdsaKey1, "");
    }

    function testGlobalKeyPersistence_BN254() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(avs2, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        bytes memory signature1 = _generateBN254Signature(
            operator1, 
            avs1, 
            DEFAULT_OPERATOR_SET_ID, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, bn254Key1, signature1);
        
        bytes32 keyHash = BN254.hashG1Point(bn254G1Key1);
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));
        
        vm.prank(operator1);
        keyRegistrar.deregisterKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID);
        
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));
        
        bytes memory signature2 = _generateBN254Signature(
            operator2, 
            avs2, 
            DEFAULT_OPERATOR_SET_ID, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator2);
        vm.expectRevert(KeyRegistrar.KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator2, avs2, DEFAULT_OPERATOR_SET_ID, bn254Key1, signature2);
    }

    // ============ Cross-Curve Type Tests ============

    function testCrossCurveGlobalUniqueness() public {
        // Configure ECDSA and BN254 operator sets
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, 0, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, 1, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        // Register ECDSA key
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, 0, ecdsaKey1, "");
        
        // Register BN254 key (should succeed as they have different hashes)
        bytes memory signature = _generateBN254Signature(
            operator1, 
            avs1, 
            1, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, avs1, 1, bn254Key1, signature);
        
        // Both should be registered
        assertTrue(keyRegistrar.isRegistered(avs1, 0, operator1));
        assertTrue(keyRegistrar.isRegistered(avs1, 1, operator1));
    }

    // ============ Error Condition Tests ============

    function testRegisterKey_RevertWrongCurveType() public {
        // Configure for ECDSA but try to register BN254 key
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.ECDSA, DEFAULT_ROTATION_DELAY);
        
        bytes memory signature = _generateBN254Signature(
            operator1, 
            avs1, 
            DEFAULT_OPERATOR_SET_ID, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        vm.expectRevert(KeyRegistrar.InvalidKeyFormat.selector);
        keyRegistrar.registerKey(operator1, avs1, DEFAULT_OPERATOR_SET_ID, bn254Key1, signature);
    }

    function testUpdateAPK_RevertOperatorSetNotConfigured() public {
        vm.prank(avs1);
        vm.expectRevert(KeyRegistrar.OperatorSetNotConfigured.selector);
        keyRegistrar.updateAPK(avs1, DEFAULT_OPERATOR_SET_ID, operator1, false);
    }

    function testUpdateAPK_RevertKeyNotRegistered() public {
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID, KeyRegistrar.CurveType.BN254, DEFAULT_ROTATION_DELAY);
        
        vm.prank(avs1);
        vm.expectRevert(KeyRegistrar.KeyNotRegistered.selector);
        keyRegistrar.updateAPK(avs1, DEFAULT_OPERATOR_SET_ID, operator1, false);
    }
}