// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/core/AVSDirectory.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

contract AVSDirectoryUnitTests is EigenLayerUnitTestSetup, IAVSDirectoryEvents, IAVSDirectoryErrors {
    uint8 constant PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS = 0;

    AVSDirectory avsDirectory;

    address defaultAVS;
    address defaultOperator;
    uint256 defaultOperatorPk;
    ISignatureUtils.SignatureWithSaltAndExpiry defaultOperatorSignature;

    function setUp() public virtual override {
        EigenLayerUnitTestSetup.setUp();

        avsDirectory = _deployAVSD(address(delegationManagerMock));

        defaultAVS = cheats.randomAddress();
        defaultOperatorPk = cheats.randomUint(1, MAX_PRIVATE_KEY);
        defaultOperator = cheats.addr(defaultOperatorPk);
        defaultOperatorSignature = _newOperatorRegistrationSignature(
            defaultOperatorPk, defaultOperator, defaultAVS, bytes32(cheats.randomUint()), type(uint256).max
        );

        delegationManagerMock.setIsOperator(defaultOperator, true);
        isExcludedFuzzAddress[address(avsDirectory)] = true;
    }

    function _deployAVSD(
        address delegationManager
    ) internal returns (AVSDirectory) {
        return AVSDirectory(
            address(
                new TransparentUpgradeableProxy(
                    address(new AVSDirectory(IDelegationManager(delegationManager))),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        AVSDirectory.initialize.selector,
                        address(this),
                        pauserRegistry,
                        0 // 0 is initialPausedStatus
                    )
                )
            )
        );
    }

    function _newOperatorRegistrationSignature(
        uint256 operatorPk,
        address operator,
        address avs,
        bytes32 salt,
        uint256 expiry
    ) internal view returns (ISignatureUtils.SignatureWithSaltAndExpiry memory) {
        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(
            operatorPk, avsDirectory.calculateOperatorAVSRegistrationDigestHash(operator, avs, salt, expiry)
        );
        return ISignatureUtils.SignatureWithSaltAndExpiry({
            signature: abi.encodePacked(r, s, v),
            salt: salt,
            expiry: expiry
        });
    }

    /// -----------------------------------------------------------------------
    /// initialize()
    /// -----------------------------------------------------------------------

    function test_initialize_Correctness() public view {
        assertEq(address(avsDirectory.delegation()), address(delegationManagerMock));
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
        cheats.expectRevert(SignatureExpired.selector);
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
        emit OperatorAVSRegistrationStatusUpdated(
            defaultOperator, defaultAVS, OperatorAVSRegistrationStatus.REGISTERED
        );
        
        cheats.prank(defaultAVS);
        avsDirectory.registerOperatorToAVS(defaultOperator, defaultOperatorSignature);
        
        assertTrue(
            avsDirectory.avsOperatorStatus(defaultAVS, defaultOperator)
                == OperatorAVSRegistrationStatus.REGISTERED
        );
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
        cheats.expectEmit(true, true, true, false, address(avsDirectory));
        emit OperatorAVSRegistrationStatusUpdated(
            defaultOperator, defaultAVS, OperatorAVSRegistrationStatus.UNREGISTERED
        );
        
        cheats.startPrank(defaultAVS);
        avsDirectory.registerOperatorToAVS(defaultOperator, defaultOperatorSignature);
        avsDirectory.deregisterOperatorFromAVS(defaultOperator);
        cheats.stopPrank();

        assertTrue(
            avsDirectory.avsOperatorStatus(defaultAVS, defaultOperator)
                == OperatorAVSRegistrationStatus.UNREGISTERED
        );
    }
}