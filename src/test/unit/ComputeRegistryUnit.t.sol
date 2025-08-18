// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/cloud/ComputeRegistry.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/test/mocks/MockAVSRegistrar.sol";
import "src/test/mocks/ReleaseManagerMock.sol";

contract ComputeRegistryUnitTests is EigenLayerUnitTestSetup, IComputeRegistryErrors, IComputeRegistryEvents {
    using StdStyle for *;
    using ArrayLib for *;
    using OperatorSetLib for OperatorSet;

    // Constants
    bytes32 constant TOS_HASH = keccak256("Terms of Service v1.0");
    string constant VERSION = "1.0.0";
    uint constant MAX_EXPIRY = type(uint).max;

    // Contracts
    ComputeRegistry computeRegistry;
    ReleaseManagerMock releaseManagerMock;

    // Test variables
    address defaultAVS;
    address defaultSigner;
    uint defaultSignerPrivateKey;
    OperatorSet defaultOperatorSet;

    function setUp() public virtual override {
        EigenLayerUnitTestSetup.setUp();

        // Setup mock contracts
        releaseManagerMock = new ReleaseManagerMock();

        // Setup default test accounts
        defaultSignerPrivateKey = 0x1234;
        defaultSigner = vm.addr(defaultSignerPrivateKey);
        defaultAVS = makeAddr("defaultAVS");

        // Deploy ComputeRegistry
        computeRegistry = new ComputeRegistry(
            IReleaseManager(address(releaseManagerMock)),
            IAllocationManager(address(allocationManagerMock)),
            IPermissionController(address(permissionController)),
            TOS_HASH,
            VERSION
        );

        // Setup default operator set
        defaultOperatorSet = OperatorSet(defaultAVS, 0);

        // Configure mocks
        allocationManagerMock.setIsOperatorSet(defaultOperatorSet, true);
        releaseManagerMock.setHasRelease(defaultOperatorSet, true);

        // Setup permissions for default signer to call registerForCompute and deregisterFromCompute
        vm.startPrank(defaultAVS);
        permissionController.setAppointee(defaultAVS, defaultSigner, address(computeRegistry), computeRegistry.registerForCompute.selector);
        permissionController.setAppointee(
            defaultAVS, defaultSigner, address(computeRegistry), computeRegistry.deregisterFromCompute.selector
        );
        vm.stopPrank();
    }

    // Helper functions
    function _generateTOSSignature(OperatorSet memory operatorSet, address signer, uint signerPrivateKey)
        internal
        view
        returns (bytes memory)
    {
        bytes32 digest = computeRegistry.calculateTOSAgreementDigest(operatorSet, signer);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(signerPrivateKey, digest);
        return abi.encodePacked(r, s, v);
    }

    function _generateInvalidSignature() internal pure returns (bytes memory) {
        return abi.encodePacked(bytes32(0), bytes32(0), uint8(0));
    }
}

contract ComputeRegistryUnitTests_Initialization is ComputeRegistryUnitTests {
    function test_initialization() public view {
        assertEq(address(computeRegistry.RELEASE_MANAGER()), address(releaseManagerMock));
        assertEq(address(computeRegistry.ALLOCATION_MANAGER()), address(allocationManagerMock));
        assertEq(address(computeRegistry.permissionController()), address(permissionController));
        assertEq(computeRegistry.TOS_HASH(), TOS_HASH);
        assertEq(computeRegistry.MAX_EXPIRY(), MAX_EXPIRY);
        assertEq(
            computeRegistry.TOS_AGREEMENT_TYPEHASH(),
            keccak256("TOSAgreement(bytes32 tosHash,address avs,uint32 operatorSetId,address signer,uint256 expiry)")
        );
    }

    function test_domainSeparator() public view {
        bytes32 expectedDomainSeparator = keccak256(
            abi.encode(
                keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"),
                keccak256(bytes("EigenLayer")),
                keccak256(bytes("1")), // Major version only
                block.chainid,
                address(computeRegistry)
            )
        );
        assertEq(computeRegistry.domainSeparator(), expectedDomainSeparator);
    }
}

contract ComputeRegistryUnitTests_RegisterForCompute is ComputeRegistryUnitTests {
    function test_registerForCompute_success() public {
        // Generate valid signature
        bytes memory signature = _generateTOSSignature(defaultOperatorSet, defaultSigner, defaultSignerPrivateKey);

        // Register
        vm.expectEmit(true, true, true, true);
        emit OperatorSetRegistered(defaultOperatorSet, defaultSigner, TOS_HASH, signature);

        vm.prank(defaultSigner);
        computeRegistry.registerForCompute(defaultOperatorSet, signature);

        // Verify registration
        assertTrue(computeRegistry.isOperatorSetRegistered(defaultOperatorSet.key()));

        // Verify TOS signature storage
        IComputeRegistryTypes.TOSSignature memory tosSignature = computeRegistry.getOperatorSetTosSignature(defaultOperatorSet);
        assertEq(tosSignature.signer, defaultSigner);
        assertEq(tosSignature.tosHash, TOS_HASH);
        assertEq(tosSignature.signature, signature);
    }

    function test_registerForCompute_revert_invalidPermissions() public {
        address unauthorizedCaller = address(0x999);
        bytes memory signature = _generateTOSSignature(defaultOperatorSet, unauthorizedCaller, defaultSignerPrivateKey);

        vm.prank(unauthorizedCaller);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        computeRegistry.registerForCompute(defaultOperatorSet, signature);
    }

    function test_registerForCompute_revert_invalidOperatorSet() public {
        OperatorSet memory invalidOperatorSet = OperatorSet(defaultAVS, 999);
        allocationManagerMock.setIsOperatorSet(invalidOperatorSet, false);

        bytes memory signature = _generateTOSSignature(invalidOperatorSet, defaultSigner, defaultSignerPrivateKey);

        vm.prank(defaultSigner);
        vm.expectRevert(InvalidOperatorSet.selector);
        computeRegistry.registerForCompute(invalidOperatorSet, signature);
    }

    function test_registerForCompute_revert_alreadyRegistered() public {
        bytes memory signature = _generateTOSSignature(defaultOperatorSet, defaultSigner, defaultSignerPrivateKey);

        // First registration succeeds
        vm.prank(defaultSigner);
        computeRegistry.registerForCompute(defaultOperatorSet, signature);

        // Second registration fails
        vm.prank(defaultSigner);
        vm.expectRevert(OperatorSetAlreadyRegistered.selector);
        computeRegistry.registerForCompute(defaultOperatorSet, signature);
    }

    function test_registerForCompute_revert_noReleases() public {
        OperatorSet memory operatorSet = OperatorSet(defaultAVS, 1);
        allocationManagerMock.setIsOperatorSet(operatorSet, true);
        releaseManagerMock.setHasRelease(operatorSet, false);

        bytes memory signature = _generateTOSSignature(operatorSet, defaultSigner, defaultSignerPrivateKey);

        vm.prank(defaultSigner);
        vm.expectRevert(IReleaseManagerErrors.NoReleases.selector);
        computeRegistry.registerForCompute(operatorSet, signature);
    }

    function test_registerForCompute_revert_invalidSignature() public {
        bytes memory invalidSignature = _generateInvalidSignature();

        vm.prank(defaultSigner);
        vm.expectRevert(ISignatureUtilsMixinErrors.InvalidSignature.selector);
        computeRegistry.registerForCompute(defaultOperatorSet, invalidSignature);
    }

    function test_registerForCompute_revert_wrongSigner() public {
        // Generate signature with different private key
        uint wrongPrivateKey = 0x5678;
        bytes memory signature = _generateTOSSignature(defaultOperatorSet, defaultSigner, wrongPrivateKey);

        vm.prank(defaultSigner);
        vm.expectRevert(ISignatureUtilsMixinErrors.InvalidSignature.selector);
        computeRegistry.registerForCompute(defaultOperatorSet, signature);
    }
}

contract ComputeRegistryUnitTests_DeregisterFromCompute is ComputeRegistryUnitTests {
    function setUp() public override {
        super.setUp();

        // Pre-register default operator set
        bytes memory signature = _generateTOSSignature(defaultOperatorSet, defaultSigner, defaultSignerPrivateKey);
        vm.prank(defaultSigner);
        computeRegistry.registerForCompute(defaultOperatorSet, signature);
    }

    function test_deregisterFromCompute_success() public {
        assertTrue(computeRegistry.isOperatorSetRegistered(defaultOperatorSet.key()));

        vm.expectEmit(true, true, true, true);
        emit OperatorSetDeregistered(defaultOperatorSet);

        vm.prank(defaultSigner);
        computeRegistry.deregisterFromCompute(defaultOperatorSet);

        assertFalse(computeRegistry.isOperatorSetRegistered(defaultOperatorSet.key()));
    }

    function test_deregisterFromCompute_revert_invalidPermissions() public {
        address unauthorizedCaller = address(0x999);

        vm.prank(unauthorizedCaller);
        vm.expectRevert(PermissionControllerMixin.InvalidPermissions.selector);
        computeRegistry.deregisterFromCompute(defaultOperatorSet);
    }

    function test_deregisterFromCompute_revert_invalidOperatorSet() public {
        OperatorSet memory invalidOperatorSet = OperatorSet(defaultAVS, 999);
        allocationManagerMock.setIsOperatorSet(invalidOperatorSet, false);

        vm.prank(defaultSigner);
        vm.expectRevert(InvalidOperatorSet.selector);
        computeRegistry.deregisterFromCompute(invalidOperatorSet);
    }

    function test_deregisterFromCompute_revert_notRegistered() public {
        OperatorSet memory unregisteredOperatorSet = OperatorSet(defaultAVS, 1);
        allocationManagerMock.setIsOperatorSet(unregisteredOperatorSet, true);

        vm.prank(defaultSigner);
        vm.expectRevert(OperatorSetNotRegistered.selector);
        computeRegistry.deregisterFromCompute(unregisteredOperatorSet);
    }

    function test_deregisterFromCompute_canReregisterAfterDeregister() public {
        // Deregister
        vm.prank(defaultSigner);
        computeRegistry.deregisterFromCompute(defaultOperatorSet);
        assertFalse(computeRegistry.isOperatorSetRegistered(defaultOperatorSet.key()));

        // Re-register
        bytes memory signature = _generateTOSSignature(defaultOperatorSet, defaultSigner, defaultSignerPrivateKey);
        vm.prank(defaultSigner);
        computeRegistry.registerForCompute(defaultOperatorSet, signature);
        assertTrue(computeRegistry.isOperatorSetRegistered(defaultOperatorSet.key()));
    }
}

contract ComputeRegistryUnitTests_ViewFunctions is ComputeRegistryUnitTests {
    function test_getOperatorSetTosSignature_registered() public {
        bytes memory signature = _generateTOSSignature(defaultOperatorSet, defaultSigner, defaultSignerPrivateKey);

        vm.prank(defaultSigner);
        computeRegistry.registerForCompute(defaultOperatorSet, signature);

        IComputeRegistryTypes.TOSSignature memory tosSignature = computeRegistry.getOperatorSetTosSignature(defaultOperatorSet);
        assertEq(tosSignature.signer, defaultSigner);
        assertEq(tosSignature.tosHash, TOS_HASH);
        assertEq(tosSignature.signature, signature);
    }

    function test_getOperatorSetTosSignature_notRegistered() public {
        IComputeRegistryTypes.TOSSignature memory tosSignature = computeRegistry.getOperatorSetTosSignature(defaultOperatorSet);
        assertEq(tosSignature.signer, address(0));
        assertEq(tosSignature.tosHash, bytes32(0));
        assertEq(tosSignature.signature.length, 0);
    }

    function test_calculateTOSAgreementDigest() public view {
        bytes32 digest = computeRegistry.calculateTOSAgreementDigest(defaultOperatorSet, defaultSigner);

        // Verify digest is deterministic
        bytes32 digest2 = computeRegistry.calculateTOSAgreementDigest(defaultOperatorSet, defaultSigner);
        assertEq(digest, digest2);

        // Verify digest changes with different parameters
        bytes32 digestDifferentSigner = computeRegistry.calculateTOSAgreementDigest(defaultOperatorSet, address(0x999));
        assertTrue(digest != digestDifferentSigner);

        OperatorSet memory differentOperatorSet = OperatorSet(defaultAVS, 1);
        bytes32 digestDifferentOperatorSet = computeRegistry.calculateTOSAgreementDigest(differentOperatorSet, defaultSigner);
        assertTrue(digest != digestDifferentOperatorSet);
    }

    function test_isOperatorSetRegistered() public {
        assertFalse(computeRegistry.isOperatorSetRegistered(defaultOperatorSet.key()));

        bytes memory signature = _generateTOSSignature(defaultOperatorSet, defaultSigner, defaultSignerPrivateKey);
        vm.prank(defaultSigner);
        computeRegistry.registerForCompute(defaultOperatorSet, signature);

        assertTrue(computeRegistry.isOperatorSetRegistered(defaultOperatorSet.key()));
    }
}
