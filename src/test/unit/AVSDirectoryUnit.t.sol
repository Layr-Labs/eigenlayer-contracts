// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/core/AVSDirectory.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

contract AVSDirectoryUnitTests is EigenLayerUnitTestSetup, IAVSDirectoryEvents, IAVSDirectoryErrors, ISignatureUtilsMixinTypes {
    uint8 constant PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS = 0;

    AVSDirectory avsDirectory;

    address defaultAVS;
    address defaultOperator;
    uint defaultOperatorPk;
    SignatureWithSaltAndExpiry defaultOperatorSignature;

    function setUp() public virtual override {
        EigenLayerUnitTestSetup.setUp();

        avsDirectory = _deployAVSD(address(delegationManagerMock), pauserRegistry);

        defaultAVS = cheats.randomAddress();
        defaultOperatorPk = cheats.randomUint(1, MAX_PRIVATE_KEY);
        defaultOperator = cheats.addr(defaultOperatorPk);
        defaultOperatorSignature = _newOperatorRegistrationSignature({
            operatorPk: defaultOperatorPk,
            avs: defaultAVS,
            salt: bytes32(cheats.randomUint()),
            expiry: type(uint).max
        });

        delegationManagerMock.setIsOperator(defaultOperator, true);
    }

    function _deployAVSD(address delegationManager, IPauserRegistry pauserRegistry) internal returns (AVSDirectory avsd) {
        avsd = AVSDirectory(
            address(
                new TransparentUpgradeableProxy(
                    address(new AVSDirectory(IDelegationManager(delegationManager), pauserRegistry, "v9.9.9")),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        AVSDirectory.initialize.selector,
                        address(this),
                        0 // 0 is initialPausedStatus
                    )
                )
            )
        );
        isExcludedFuzzAddress[address(avsd)] = true;

        bytes memory v = bytes(avsd.version());
        bytes32 expectedDomainSeparator = keccak256(
            abi.encode(
                EIP712_DOMAIN_TYPEHASH, keccak256(bytes("EigenLayer")), keccak256(bytes.concat(v[0], v[1])), block.chainid, address(avsd)
            )
        );

        assertEq(avsd.domainSeparator(), expectedDomainSeparator, "sanity check");
    }

    function _newOperatorRegistrationSignature(uint operatorPk, address avs, bytes32 salt, uint expiry)
        internal
        view
        returns (SignatureWithSaltAndExpiry memory)
    {
        (uint8 v, bytes32 r, bytes32 s) =
            cheats.sign(operatorPk, avsDirectory.calculateOperatorAVSRegistrationDigestHash(cheats.addr(operatorPk), avs, salt, expiry));
        return SignatureWithSaltAndExpiry({signature: abi.encodePacked(r, s, v), salt: salt, expiry: expiry});
    }

    /// -----------------------------------------------------------------------
    /// initialize()
    /// -----------------------------------------------------------------------

    function test_initialize_Correctness() public {
        assertEq(address(avsDirectory.delegation()), address(delegationManagerMock));
        cheats.expectRevert("Initializable: contract is already initialized");
        avsDirectory.initialize(address(this), 0);
    }

    /// -----------------------------------------------------------------------
    /// updateAVSMetadataURI()
    /// -----------------------------------------------------------------------

    function test_updateAVSMetadataURI_Correctness() public {
        cheats.expectEmit(true, true, true, true, address(avsDirectory));
        emit AVSMetadataURIUpdated(address(this), "test");
        avsDirectory.updateAVSMetadataURI("test");
    }

    /// -----------------------------------------------------------------------
    /// cancelSalt()
    /// -----------------------------------------------------------------------

    function test_cancelSalt_Correctness() public {
        bytes32 salt = bytes32(cheats.randomUint());
        cheats.prank(defaultAVS);
        avsDirectory.cancelSalt(salt);
        assertTrue(avsDirectory.operatorSaltIsSpent(defaultAVS, salt));
    }

    /// -----------------------------------------------------------------------
    /// registerOperatorToAVS()
    /// -----------------------------------------------------------------------

    function test_registerOperatorToAVS_Paused() public {
        cheats.prank(pauser);
        avsDirectory.pause(2 ** PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        avsDirectory.registerOperatorToAVS(defaultOperator, defaultOperatorSignature);
    }

    function test_registerOperatorToAVS_SignatureExpired() public {
        defaultOperatorSignature.expiry = block.timestamp - 1;
        cheats.expectRevert(ISignatureUtilsMixinErrors.SignatureExpired.selector);
        avsDirectory.registerOperatorToAVS(defaultOperator, defaultOperatorSignature);
    }

    function test_registerOperatorToAVS_OperatorAlreadyRegistered() public {
        cheats.startPrank(defaultAVS);
        avsDirectory.registerOperatorToAVS(defaultOperator, defaultOperatorSignature);
        cheats.expectRevert(OperatorAlreadyRegisteredToAVS.selector);
        avsDirectory.registerOperatorToAVS(defaultOperator, defaultOperatorSignature);
        cheats.stopPrank();
    }

    function test_registerOperatorToAVS_SaltSpent() public {
        cheats.prank(defaultOperator);
        avsDirectory.cancelSalt(defaultOperatorSignature.salt);
        cheats.prank(defaultAVS);
        cheats.expectRevert(SaltSpent.selector);
        avsDirectory.registerOperatorToAVS(defaultOperator, defaultOperatorSignature);
    }

    function test_registerOperatorToAVS_OperatorNotRegisteredToEigenLayer() public {
        delegationManagerMock.setIsOperator(defaultOperator, false);
        cheats.expectRevert(OperatorNotRegisteredToEigenLayer.selector);
        avsDirectory.registerOperatorToAVS(defaultOperator, defaultOperatorSignature);
    }

    function test_registerOperatorToAVS_Correctness() public {
        cheats.expectEmit(true, true, true, false, address(avsDirectory));
        emit OperatorAVSRegistrationStatusUpdated(defaultOperator, defaultAVS, OperatorAVSRegistrationStatus.REGISTERED);

        cheats.prank(defaultAVS);
        avsDirectory.registerOperatorToAVS(defaultOperator, defaultOperatorSignature);

        assertTrue(avsDirectory.avsOperatorStatus(defaultAVS, defaultOperator) == OperatorAVSRegistrationStatus.REGISTERED);
        assertTrue(avsDirectory.operatorSaltIsSpent(defaultOperator, defaultOperatorSignature.salt));
    }

    /// -----------------------------------------------------------------------
    /// deregisterOperatorFromAVS()
    /// -----------------------------------------------------------------------

    function test_deregisterOperatorFromAVS_Paused() public {
        cheats.prank(pauser);
        avsDirectory.pause(2 ** PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        avsDirectory.deregisterOperatorFromAVS(defaultOperator);
    }

    function test_deregisterOperatorFromAVS_OperatorNotRegisteredToAVS() public {
        cheats.expectRevert(OperatorNotRegisteredToAVS.selector);
        avsDirectory.deregisterOperatorFromAVS(defaultOperator);
    }

    function test_deregisterOperatorFromAVS_Correctness() public {
        cheats.startPrank(defaultAVS);
        avsDirectory.registerOperatorToAVS(defaultOperator, defaultOperatorSignature);

        cheats.expectEmit(true, true, true, false, address(avsDirectory));
        emit OperatorAVSRegistrationStatusUpdated(defaultOperator, defaultAVS, OperatorAVSRegistrationStatus.UNREGISTERED);

        avsDirectory.deregisterOperatorFromAVS(defaultOperator);
        cheats.stopPrank();

        assertTrue(avsDirectory.avsOperatorStatus(defaultAVS, defaultOperator) == OperatorAVSRegistrationStatus.UNREGISTERED);
    }
}
