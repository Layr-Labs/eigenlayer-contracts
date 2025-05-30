//SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "../../contracts/core/KeyRegistrar.sol";
import {IKeyRegistrar, IKeyRegistrarTypes, IKeyRegistrarErrors} from "../../contracts/interfaces/IKeyRegistrar.sol";
import {BN254} from "../../contracts/libraries/BN254.sol";
import {IPermissionController} from "../../contracts/interfaces/IPermissionController.sol";
import {PermissionControllerMixin} from "../../contracts/mixins/PermissionControllerMixin.sol";
import {OperatorSet} from "../../contracts/libraries/OperatorSetLib.sol";
import {AllocationManagerMock} from "../mocks/AllocationManagerMock.sol";

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
    AllocationManagerMock public allocationManager;
    
    address public owner = address(0x1);
    address public operator1 = address(0x2);
    address public operator2 = address(0x3);
    address public avs1 = address(0x4);
    address public avs2 = address(0x5);
    
    uint32 public constant DEFAULT_OPERATOR_SET_ID = 0;
    
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

    event KeyRegistered(OperatorSet operatorSet, address indexed operator, IKeyRegistrarTypes.CurveType curveType, bytes pubkey);
    event KeyDeregistered(OperatorSet operatorSet, address indexed operator, IKeyRegistrarTypes.CurveType curveType);
    event AggregateBN254KeyUpdated(OperatorSet operatorSet, BN254.G1Point newAggregateKey);
    event OperatorSetConfigured(OperatorSet operatorSet, IKeyRegistrarTypes.CurveType curveType);

    function setUp() public {
        permissionController = new MockPermissionController();
        allocationManager = new AllocationManagerMock();
        keyRegistrar = new KeyRegistrar(
            IPermissionController(address(permissionController)), 
            IAllocationManager(address(allocationManager)),
            "1.0.0"
        );
        
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

    function _createOperatorSet(address avs, uint32 operatorSetId) internal pure returns (OperatorSet memory) {
        return OperatorSet({avs: avs, id: operatorSetId});
    }

    function _generateBN254Signature(
        address operator,
        OperatorSet memory operatorSet,
        bytes memory pubkey,
        uint256 privKey
    ) internal view returns (bytes memory) {
        bytes32 messageHash = keccak256(abi.encodePacked(
            "EigenLayer.KeyRegistrar.v1",
            address(keyRegistrar),
            operatorSet.avs, 
            operatorSet.id, 
            operator, 
            pubkey
        ));
        
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
        
        IKeyRegistrarTypes.OperatorSetConfig memory config = keyRegistrar.getOperatorSetConfig(operatorSet);
        assertEq(uint8(config.curveType), uint8(IKeyRegistrarTypes.CurveType.ECDSA));
        assertTrue(config.isActive);
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
        
        IKeyRegistrarTypes.OperatorSetConfig memory config = keyRegistrar.getOperatorSetConfig(operatorSet);
        assertEq(uint8(config.curveType), uint8(IKeyRegistrarTypes.CurveType.BN254));
        assertTrue(config.isActive);
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
        
        vm.prank(operator1);
        vm.expectEmit(true, true, true, true);
        emit KeyRegistered(operatorSet, operator1, IKeyRegistrarTypes.CurveType.ECDSA, ecdsaKey1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, "");
        
        assertTrue(keyRegistrar.isRegistered(operatorSet, operator1));
        bytes memory storedKey = keyRegistrar.getECDSAKey(operatorSet, operator1);
        assertEq(storedKey, ecdsaKey1);
    }

    function testRegisterECDSAKey_RevertInvalidFormat() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
        
        bytes memory invalidKey = hex"04b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4";
        
        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.InvalidKeyFormat.selector);
        keyRegistrar.registerKey(operator1, operatorSet, invalidKey, "");
        
        bytes memory invalidPrefix = hex"03b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c86a57d8a35ebfd17be4b8d44d0f7b8d5b5e5e5e5e5e5e5e5e5e5e5e5e5e5e5e5e5";
        
        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.InvalidKeyFormat.selector);
        keyRegistrar.registerKey(operator1, operatorSet, invalidPrefix, "");
    }

    function testRegisterECDSAKey_RevertZeroPubkey() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
        
        bytes memory zeroKey = new bytes(65);
        zeroKey[0] = 0x04;
        
        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.ZeroPubkey.selector);
        keyRegistrar.registerKey(operator1, operatorSet, zeroKey, "");
    }

    function testRegisterECDSAKey_RevertAlreadyRegistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, "");
        
        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, "");
    }

    function testRegisterECDSAKey_RevertGloballyRegistered() public {
        OperatorSet memory operatorSet1 = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        OperatorSet memory operatorSet2 = _createOperatorSet(avs2, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet1, IKeyRegistrarTypes.CurveType.ECDSA);
        
        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(operatorSet2, IKeyRegistrarTypes.CurveType.ECDSA);
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, ecdsaKey1, "");
        
        vm.prank(operator2);
        vm.expectRevert(IKeyRegistrarErrors.KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator2, operatorSet2, ecdsaKey1, "");
    }

    // ============ BN254 Key Registration Tests ============

    function testRegisterBN254Key() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        
        bytes memory signature = _generateBN254Signature(
            operator1, 
            operatorSet, 
            bn254Key1, 
            bn254PrivKey1
        );
        
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
        
        bytes memory zeroKey = abi.encode(uint256(0), uint256(0), bn254G2Key1.X, bn254G2Key1.Y);
        bytes memory signature = _generateBN254Signature(
            operator1, 
            operatorSet, 
            zeroKey, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.ZeroPubkey.selector);
        keyRegistrar.registerKey(operator1, operatorSet, zeroKey, signature);
    }

    function testRegisterBN254Key_RevertInvalidSignature() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        
        bytes memory invalidSignature = abi.encode(uint256(1), uint256(2));
        
        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.InvalidSignature.selector);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, invalidSignature);
    }

    function testRegisterBN254Key_RevertAlreadyRegistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        
        bytes memory signature = _generateBN254Signature(
            operator1, 
            operatorSet, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);
        
        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);
    }

    function testRegisterBN254Key_RevertGloballyRegistered() public {
        OperatorSet memory operatorSet1 = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        OperatorSet memory operatorSet2 = _createOperatorSet(avs2, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet1, IKeyRegistrarTypes.CurveType.BN254);
        
        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(operatorSet2, IKeyRegistrarTypes.CurveType.BN254);
        
        bytes memory signature1 = _generateBN254Signature(
            operator1, 
            operatorSet1, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, bn254Key1, signature1);
        
        bytes memory signature2 = _generateBN254Signature(
            operator2, 
            operatorSet2, 
            bn254Key1, 
            bn254PrivKey1
        );
        
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
            bn254PrivKey2  // Wrong private key
        );
        
        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.InvalidSignature.selector);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, wrongSignature);
    }

    // ============ Key Deregistration Tests ============

    function testDeregisterKeyAndUpdate_ECDSA() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, "");
        
        vm.prank(avs1);
        vm.expectEmit(true, true, true, true);
        emit KeyDeregistered(operatorSet, operator1, IKeyRegistrarTypes.CurveType.ECDSA);
        keyRegistrar.deregisterKeyAndUpdate(operator1, operatorSet);
        
        assertFalse(keyRegistrar.isRegistered(operatorSet, operator1));
    }

    function testDeregisterKeyAndUpdate_BN254() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        
        bytes memory signature = _generateBN254Signature(
            operator1, 
            operatorSet, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);
        
        // First add to APK
        vm.prank(avs1);
        keyRegistrar.checkAndUpdateKey(operatorSet, operator1);
        
        vm.prank(avs1);
        vm.expectEmit(true, true, true, true);
        emit KeyDeregistered(operatorSet, operator1, IKeyRegistrarTypes.CurveType.BN254);
        keyRegistrar.deregisterKeyAndUpdate(operator1, operatorSet);
        
        assertFalse(keyRegistrar.isRegistered(operatorSet, operator1));
        
        // Verify aggregate key was updated (should be zero after removing the only key)
        BN254.G1Point memory apk = keyRegistrar.getApk(operatorSet);
        assertEq(apk.X, 0);
        assertEq(apk.Y, 0);
    }

    function testDeregisterKeyAndUpdate_RevertKeyNotFound() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
        
        vm.prank(avs1);
        vm.expectRevert(abi.encodeWithSelector(IKeyRegistrarErrors.KeyNotFound.selector, operatorSet, operator1));
        keyRegistrar.deregisterKeyAndUpdate(operator1, operatorSet);
    }

    function testDeregisterKeyAndUpdate_RevertUnauthorized() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(operator1);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        keyRegistrar.deregisterKeyAndUpdate(operator1, operatorSet);
    }

    // ============ APK Management Tests ============

    function testCheckAndUpdateKey_OnlyAuthorized() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        
        vm.prank(operator1);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        keyRegistrar.checkAndUpdateKey(operatorSet, operator1);
    }

    function testCheckAndUpdateKey_AuthorizedCall() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);

        bytes memory signature = _generateBN254Signature(
            operator1, 
            operatorSet, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);
        
        vm.prank(avs1);
        vm.expectEmit(true, true, true, true);
        emit AggregateBN254KeyUpdated(operatorSet, bn254G1Key1);
        keyRegistrar.checkAndUpdateKey(operatorSet, operator1);
        
        // Verify aggregate key was updated
        BN254.G1Point memory apk = keyRegistrar.getApk(operatorSet);
        assertEq(apk.X, bn254G1Key1.X);
        assertEq(apk.Y, bn254G1Key1.Y);
    }

    function testCheckAndUpdateKey_ECDSAType_NoOp() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, "");
        
        vm.prank(avs1);
        keyRegistrar.checkAndUpdateKey(operatorSet, operator1);
        
        // Should not emit any events or change APK for ECDSA
        BN254.G1Point memory apk = keyRegistrar.getApk(operatorSet);
        assertEq(apk.X, 0);
        assertEq(apk.Y, 0);
    }

    function testCheckAndUpdateKey_RevertKeyNotFound() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        
        vm.prank(avs1);
        vm.expectRevert(abi.encodeWithSelector(IKeyRegistrarErrors.KeyNotFound.selector, operatorSet, operator1));
        keyRegistrar.checkAndUpdateKey(operatorSet, operator1);
    }

    // ============ View Function Tests ============

    function testGetOperatorSetConfig() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
        
        IKeyRegistrarTypes.OperatorSetConfig memory config = keyRegistrar.getOperatorSetConfig(operatorSet);
        assertEq(uint8(config.curveType), uint8(IKeyRegistrarTypes.CurveType.ECDSA));
        assertTrue(config.isActive);
    }

    function testGetKeyHash_ECDSA() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, "");
        
        bytes32 keyHash = keyRegistrar.getKeyHash(operatorSet, operator1);
        bytes32 expectedHash = keccak256(ecdsaKey1);
        assertEq(keyHash, expectedHash);
    }

    function testGetKeyHash_BN254() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        
        bytes memory signature = _generateBN254Signature(
            operator1, 
            operatorSet, 
            bn254Key1, 
            bn254PrivKey1
        );
        
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
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, ecdsaKey1, "");
        
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));
    }

    function testGetBN254Key_EmptyForUnregistered() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
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
        bytes memory key = keyRegistrar.getECDSAKey(operatorSet, operator1);
        assertEq(key.length, 0);
    }

    function testGetApk_ReturnsZeroForUninitialized() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        BN254.G1Point memory apk = keyRegistrar.getApk(operatorSet);
        assertEq(apk.X, 0);
        assertEq(apk.Y, 0);
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

    function testDeregisterKeyAndUpdate_RevertOperatorSetNotConfigured() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        vm.expectRevert(IKeyRegistrarErrors.OperatorSetNotConfigured.selector);
        keyRegistrar.deregisterKeyAndUpdate(operator1, operatorSet);
    }

    function testCheckAndUpdateKey_RevertOperatorSetNotConfigured() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        vm.expectRevert(IKeyRegistrarErrors.OperatorSetNotConfigured.selector);
        keyRegistrar.checkAndUpdateKey(operatorSet, operator1);
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
        
        IKeyRegistrarTypes.OperatorSetConfig memory config1 = keyRegistrar.getOperatorSetConfig(operatorSet1);
        IKeyRegistrarTypes.OperatorSetConfig memory config2 = keyRegistrar.getOperatorSetConfig(operatorSet2);
        
        assertEq(uint8(config1.curveType), uint8(IKeyRegistrarTypes.CurveType.ECDSA));
        assertEq(uint8(config2.curveType), uint8(IKeyRegistrarTypes.CurveType.BN254));
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, ecdsaKey1, "");
        
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
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, ecdsaOperatorSet, ecdsaKey1, "");
        
        // Register BN254 key for another operator set
        bytes memory signature = _generateBN254Signature(
            operator1, 
            bn254OperatorSet, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, bn254OperatorSet, bn254Key1, signature);
        
        // Verify both registrations
        assertTrue(keyRegistrar.isRegistered(ecdsaOperatorSet, operator1));
        assertTrue(keyRegistrar.isRegistered(bn254OperatorSet, operator1));
        
        // Verify key retrieval
        bytes memory ecdsaKey = keyRegistrar.getECDSAKey(ecdsaOperatorSet, operator1);
        assertEq(ecdsaKey, ecdsaKey1);
        
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
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, ecdsaKey1, "");
        
        bytes32 keyHash = keccak256(ecdsaKey1);
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));
        
        vm.prank(avs1);
        keyRegistrar.deregisterKeyAndUpdate(operator1, operatorSet1);
        
        // Key should still be globally registered after deregistration
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));
        
        vm.prank(operator2);
        vm.expectRevert(IKeyRegistrarErrors.KeyAlreadyRegistered.selector);
        keyRegistrar.registerKey(operator2, operatorSet2, ecdsaKey1, "");
    }

    function testGlobalKeyPersistence_BN254() public {
        OperatorSet memory operatorSet1 = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        OperatorSet memory operatorSet2 = _createOperatorSet(avs2, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet1, IKeyRegistrarTypes.CurveType.BN254);
        
        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(operatorSet2, IKeyRegistrarTypes.CurveType.BN254);
        
        bytes memory signature1 = _generateBN254Signature(
            operator1, 
            operatorSet1, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet1, bn254Key1, signature1);
        
        bytes32 keyHash = BN254.hashG1Point(bn254G1Key1);
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));
        
        vm.prank(avs1);
        keyRegistrar.deregisterKeyAndUpdate(operator1, operatorSet1);
        
        // Key should still be globally registered after deregistration
        assertTrue(keyRegistrar.isKeyGloballyRegistered(keyHash));
        
        bytes memory signature2 = _generateBN254Signature(
            operator2, 
            operatorSet2, 
            bn254Key1, 
            bn254PrivKey1
        );
        
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
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, ecdsaOperatorSet, ecdsaKey1, "");
        
        // Register BN254 key (should succeed as they have different hashes)
        bytes memory signature = _generateBN254Signature(
            operator1, 
            bn254OperatorSet, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, bn254OperatorSet, bn254Key1, signature);
        
        // Both should be registered
        assertTrue(keyRegistrar.isRegistered(ecdsaOperatorSet, operator1));
        assertTrue(keyRegistrar.isRegistered(bn254OperatorSet, operator1));
    }

    // ============ BN254 Aggregate Key Tests ============

    function testAggregateBN254Key_MultipleOperators() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        
        // Register first operator
        bytes memory signature1 = _generateBN254Signature(
            operator1, 
            operatorSet, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature1);
        
        vm.prank(avs1);
        keyRegistrar.checkAndUpdateKey(operatorSet, operator1);
        
        // Register second operator
        bytes memory signature2 = _generateBN254Signature(
            operator2, 
            operatorSet, 
            bn254Key2, 
            bn254PrivKey2
        );
        
        vm.prank(operator2);
        keyRegistrar.registerKey(operator2, operatorSet, bn254Key2, signature2);
        
        vm.prank(avs1);
        keyRegistrar.checkAndUpdateKey(operatorSet, operator2);
        
        // Verify aggregate key is the sum of both keys
        BN254.G1Point memory apk = keyRegistrar.getApk(operatorSet);
        BN254.G1Point memory expectedApk = bn254G1Key1.plus(bn254G1Key2);
        
        assertEq(apk.X, expectedApk.X);
        assertEq(apk.Y, expectedApk.Y);
        
        // Remove first operator
        vm.prank(avs1);
        keyRegistrar.deregisterKeyAndUpdate(operator1, operatorSet);
        
        // Verify aggregate key is now just the second key
        apk = keyRegistrar.getApk(operatorSet);
        assertEq(apk.X, bn254G1Key2.X);
        assertEq(apk.Y, bn254G1Key2.Y);
    }

    // ============ Error Condition Tests ============

    function testRegisterKey_RevertWrongCurveType() public {
        // Configure for ECDSA but try to register BN254 key
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
        
        bytes memory signature = _generateBN254Signature(
            operator1, 
            operatorSet, 
            bn254Key1, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.InvalidKeyFormat.selector);
        keyRegistrar.registerKey(operator1, operatorSet, bn254Key1, signature);
    }

    function testRegisterBN254Key_RevertInvalidKeypair() public {
        OperatorSet memory operatorSet = _createOperatorSet(avs1, DEFAULT_OPERATOR_SET_ID);
        
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);
        
        // Create invalid keypair (wrong G2 for G1)
        bytes memory invalidKey = abi.encode(bn254G1Key1.X, bn254G1Key1.Y, bn254G2Key2.X, bn254G2Key2.Y);
        bytes memory signature = _generateBN254Signature(
            operator1, 
            operatorSet, 
            invalidKey, 
            bn254PrivKey1
        );
        
        vm.prank(operator1);
        vm.expectRevert(IKeyRegistrarErrors.InvalidKeypair.selector);
        keyRegistrar.registerKey(operator1, operatorSet, invalidKey, signature);
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
        BN254.G1Point memory signature = msgPoint.scalar_mul(bn254PrivKey1);
        bytes memory signatureBytes = abi.encode(signature.X, signature.Y);
        
        // Should not revert for valid signature
        keyRegistrar.verifyBN254Signature(messageHash, signatureBytes, bn254G1Key1, bn254G2Key1);
    }

    function testVerifyBN254Signature_RevertInvalid() public {
        bytes32 messageHash = keccak256("test message");
        bytes memory invalidSignature = abi.encode(uint256(1), uint256(2));
        
        vm.expectRevert(IKeyRegistrarErrors.InvalidSignature.selector);
        keyRegistrar.verifyBN254Signature(messageHash, invalidSignature, bn254G1Key1, bn254G2Key1);
    }
}