// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";

import "src/contracts/core/DelegationManager.sol";
import "src/contracts/core/AVSDirectory.sol";

import "src/test/events/IAVSDirectoryEvents.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

/**
 * @notice Unit testing of the AVSDirectory contract. An AVSs' service manager contract will
 * call this to register an operator with the AVS.
 * Contracts tested: AVSDirectory
 * Contracts not mocked: DelegationManager
 */
contract AVSDirectoryUnitTests is EigenLayerUnitTestSetup, IAVSDirectoryEvents {
    // Contract under test
    AVSDirectory avsDirectory;
    AVSDirectory avsDirectoryImplementation;

    // Contract dependencies
    DelegationManager delegationManager;
    DelegationManager delegationManagerImplementation;

    // Delegation signer
    uint256 delegationSignerPrivateKey = uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);
    uint256 stakerPrivateKey = uint256(123_456_789);

    // empty string reused across many tests
    string emptyStringForMetadataURI;

    // reused in various tests. in storage to help handle stack-too-deep errors
    address defaultAVS = address(this);

    uint256 minWithdrawalDelayBlocks = 216_000;
    IStrategy[] public initializeStrategiesToSetDelayBlocks;
    uint256[] public initializeWithdrawalDelayBlocks;

    // Index for flag that pauses registering/deregistering for AVSs
    uint8 internal constant PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS = 0;

    function setUp() public virtual override {
        // Setup
        EigenLayerUnitTestSetup.setUp();

        // Deploy DelegationManager implmentation and proxy
        initializeStrategiesToSetDelayBlocks = new IStrategy[](0);
        initializeWithdrawalDelayBlocks = new uint256[](0);
        delegationManagerImplementation = new DelegationManager(strategyManagerMock, slasherMock, eigenPodManagerMock);
        delegationManager = DelegationManager(
            address(
                new TransparentUpgradeableProxy(
                    address(delegationManagerImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        DelegationManager.initialize.selector,
                        address(this),
                        pauserRegistry,
                        0, // 0 is initialPausedStatus
                        minWithdrawalDelayBlocks,
                        initializeStrategiesToSetDelayBlocks,
                        initializeWithdrawalDelayBlocks
                    )
                )
            )
        );

        // Deploy AVSDirectory implmentation and proxy
        avsDirectoryImplementation = new AVSDirectory(delegationManager);
        avsDirectory = AVSDirectory(
            address(
                new TransparentUpgradeableProxy(
                    address(avsDirectoryImplementation),
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

        // Exclude delegation manager from fuzzed tests
        addressIsExcludedFromFuzzedInputs[address(avsDirectory)] = true;
    }

    /**
     * INTERNAL / HELPER FUNCTIONS
     */

    /**
     * @notice internal function for calculating a signature from the operator corresponding to `_operatorPrivateKey`, delegating them to
     * the `operator`, and expiring at `expiry`.
     */
    function _getOperatorSignature(
        uint256 _operatorPrivateKey,
        address operator,
        address avs,
        bytes32 salt,
        uint256 expiry
    ) internal view returns (ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature) {
        operatorSignature.expiry = expiry;
        operatorSignature.salt = salt;
        {
            bytes32 digestHash = avsDirectory.calculateOperatorAVSRegistrationDigestHash(operator, avs, salt, expiry);
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(_operatorPrivateKey, digestHash);
            operatorSignature.signature = abi.encodePacked(r, s, v);
        }
        return operatorSignature;
    }

    function _registerOperatorWithBaseDetails(address operator) internal {
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            __deprecated_earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        _registerOperator(operator, operatorDetails, emptyStringForMetadataURI);
    }

    function _registerOperatorWithDelegationApprover(address operator) internal {
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            __deprecated_earningsReceiver: operator,
            delegationApprover: cheats.addr(delegationSignerPrivateKey),
            stakerOptOutWindowBlocks: 0
        });
        _registerOperator(operator, operatorDetails, emptyStringForMetadataURI);
    }

    function _registerOperatorWith1271DelegationApprover(address operator) internal returns (ERC1271WalletMock) {
        address delegationSigner = cheats.addr(delegationSignerPrivateKey);
        /**
         * deploy a ERC1271WalletMock contract with the `delegationSigner` address as the owner,
         * so that we can create valid signatures from the `delegationSigner` for the contract to check when called
         */
        ERC1271WalletMock wallet = new ERC1271WalletMock(delegationSigner);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            __deprecated_earningsReceiver: operator,
            delegationApprover: address(wallet),
            stakerOptOutWindowBlocks: 0
        });
        _registerOperator(operator, operatorDetails, emptyStringForMetadataURI);

        return wallet;
    }

    function _registerOperator(
        address operator,
        IDelegationManager.OperatorDetails memory operatorDetails,
        string memory metadataURI
    ) internal filterFuzzedAddressInputs(operator) {
        _filterOperatorDetails(operator, operatorDetails);
        cheats.prank(operator);
        delegationManager.registerAsOperator(operatorDetails, metadataURI);
    }

    function _filterOperatorDetails(
        address operator,
        IDelegationManager.OperatorDetails memory operatorDetails
    ) internal view {
        // filter out zero address since people can't delegate to the zero address and operators are delegated to themselves
        cheats.assume(operator != address(0));
        // filter out disallowed stakerOptOutWindowBlocks values
        cheats.assume(operatorDetails.stakerOptOutWindowBlocks <= delegationManager.MAX_STAKER_OPT_OUT_WINDOW_BLOCKS());
    }
}

contract AVSDirectoryUnitTests_operatorAVSRegisterationStatus is AVSDirectoryUnitTests {
    function test_revert_whenRegisterDeregisterToAVSPaused() public {
        // set the pausing flag
        cheats.prank(pauser);
        avsDirectory.pause(2 ** PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS);

        cheats.expectRevert("Pausable: index is paused");
        avsDirectory.registerOperatorToAVS(address(0), ISignatureUtils.SignatureWithSaltAndExpiry(abi.encodePacked(""), 0, 0));

        cheats.expectRevert("Pausable: index is paused");
        avsDirectory.deregisterOperatorFromAVS(address(0));
    }

    // @notice Tests that an avs who calls `updateAVSMetadataURI` will correctly see an `AVSMetadataURIUpdated` event emitted with their input
    function testFuzz_UpdateAVSMetadataURI(string memory metadataURI) public {
        // call `updateAVSMetadataURI` and check for event
        cheats.expectEmit(true, true, true, true, address(avsDirectory));
        cheats.prank(defaultAVS);
        emit AVSMetadataURIUpdated(defaultAVS, metadataURI);
        avsDirectory.updateAVSMetadataURI(metadataURI);
    }

    // @notice Verifies an operator registers successfull to avs and see an `OperatorAVSRegistrationStatusUpdated` event emitted
    function testFuzz_registerOperatorToAVS(bytes32 salt) public {
        address operator = cheats.addr(delegationSignerPrivateKey);
        assertFalse(delegationManager.isOperator(operator), "bad test setup");
        _registerOperatorWithBaseDetails(operator);

        cheats.expectEmit(true, true, true, true, address(avsDirectory));
        emit OperatorAVSRegistrationStatusUpdated(
            operator, defaultAVS, IAVSDirectory.OperatorAVSRegistrationStatus.REGISTERED
        );

        uint256 expiry = type(uint256).max;

        cheats.prank(defaultAVS);
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature =
            _getOperatorSignature(delegationSignerPrivateKey, operator, defaultAVS, salt, expiry);

        avsDirectory.registerOperatorToAVS(operator, operatorSignature);
    }

    // @notice Verifies an operator registers successfull to avs and see an `OperatorAVSRegistrationStatusUpdated` event emitted
    function testFuzz_revert_whenOperatorNotRegisteredToEigenLayerYet(bytes32 salt) public {
        address operator = cheats.addr(delegationSignerPrivateKey);
        assertFalse(delegationManager.isOperator(operator), "bad test setup");

        cheats.prank(defaultAVS);
        uint256 expiry = type(uint256).max;
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature =
            _getOperatorSignature(delegationSignerPrivateKey, operator, defaultAVS, salt, expiry);

        cheats.expectRevert("AVSDirectory.registerOperatorToAVS: operator not registered to EigenLayer yet");
        avsDirectory.registerOperatorToAVS(operator, operatorSignature);
    }

    // @notice Verifies an operator registers fails when the signature is not from the operator
    function testFuzz_revert_whenSignatureAddressIsNotOperator(bytes32 salt) public {
        address operator = cheats.addr(delegationSignerPrivateKey);
        assertFalse(delegationManager.isOperator(operator), "bad test setup");
        _registerOperatorWithBaseDetails(operator);

        uint256 expiry = type(uint256).max;
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature =
            _getOperatorSignature(delegationSignerPrivateKey, operator, defaultAVS, salt, expiry);

        cheats.expectRevert("EIP1271SignatureUtils.checkSignature_EIP1271: signature not from signer");
        cheats.prank(operator);
        avsDirectory.registerOperatorToAVS(operator, operatorSignature);
    }

    // @notice Verifies an operator registers fails when the signature expiry already expires
    function testFuzz_revert_whenExpiryHasExpired(
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) public {
        address operator = cheats.addr(delegationSignerPrivateKey);
        operatorSignature.expiry = bound(operatorSignature.expiry, 0, block.timestamp - 1);

        cheats.expectRevert("AVSDirectory.registerOperatorToAVS: operator signature expired");
        avsDirectory.registerOperatorToAVS(operator, operatorSignature);
    }

    // @notice Verifies an operator registers fails when it's already registered to the avs
    function testFuzz_revert_whenOperatorAlreadyRegisteredToAVS(bytes32 salt) public {
        address operator = cheats.addr(delegationSignerPrivateKey);
        assertFalse(delegationManager.isOperator(operator), "bad test setup");
        _registerOperatorWithBaseDetails(operator);

        uint256 expiry = type(uint256).max;
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature =
            _getOperatorSignature(delegationSignerPrivateKey, operator, defaultAVS, salt, expiry);

        cheats.startPrank(defaultAVS);
        avsDirectory.registerOperatorToAVS(operator, operatorSignature);

        cheats.expectRevert("AVSDirectory.registerOperatorToAVS: operator already registered");
        avsDirectory.registerOperatorToAVS(operator, operatorSignature);
        cheats.stopPrank();
    }

    /// @notice Checks that cancelSalt updates the operatorSaltIsSpent mapping correctly
    function testFuzz_cancelSalt(bytes32 salt) public {
        address operator = cheats.addr(delegationSignerPrivateKey);
        assertFalse(delegationManager.isOperator(operator), "bad test setup");
        _registerOperatorWithBaseDetails(operator);

        assertFalse(avsDirectory.operatorSaltIsSpent(operator, salt), "bad test setup");
        assertFalse(avsDirectory.operatorSaltIsSpent(defaultAVS, salt), "bad test setup");

        cheats.prank(operator);
        avsDirectory.cancelSalt(salt);

        assertTrue(avsDirectory.operatorSaltIsSpent(operator, salt), "salt was not successfully cancelled");
        assertFalse(avsDirectory.operatorSaltIsSpent(defaultAVS, salt), "salt should only be cancelled for the operator");

        bytes32 newSalt; 
        unchecked { newSalt = bytes32(uint(salt) + 1); }

        assertFalse(salt == newSalt, "bad test setup");

        cheats.prank(operator);
        avsDirectory.cancelSalt(newSalt);

        assertTrue(avsDirectory.operatorSaltIsSpent(operator, salt), "original salt should still be cancelled");
        assertTrue(avsDirectory.operatorSaltIsSpent(operator, newSalt), "new salt should be cancelled");
    }

    /// @notice Verifies that registration fails when the salt has been cancelled via cancelSalt
    function testFuzz_revert_whenRegisteringWithCancelledSalt(bytes32 salt) public {
        address operator = cheats.addr(delegationSignerPrivateKey);
        assertFalse(delegationManager.isOperator(operator), "bad test setup");
        _registerOperatorWithBaseDetails(operator);

        uint256 expiry = type(uint256).max;
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature =
            _getOperatorSignature(delegationSignerPrivateKey, operator, defaultAVS, salt, expiry);

        cheats.prank(operator);
        avsDirectory.cancelSalt(salt);

        cheats.expectRevert("AVSDirectory.registerOperatorToAVS: salt already spent");
        cheats.prank(defaultAVS);
        avsDirectory.registerOperatorToAVS(operator, operatorSignature);
    }

    /// @notice Verifies that an operator cannot cancel the same salt twice
    function testFuzz_revert_whenSaltCancelledTwice(bytes32 salt) public {
        address operator = cheats.addr(delegationSignerPrivateKey);
        assertFalse(delegationManager.isOperator(operator), "bad test setup");
        _registerOperatorWithBaseDetails(operator);

        uint256 expiry = type(uint256).max;
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature =
            _getOperatorSignature(delegationSignerPrivateKey, operator, defaultAVS, salt, expiry);

        cheats.startPrank(operator);
        avsDirectory.cancelSalt(salt);

        cheats.expectRevert("AVSDirectory.cancelSalt: cannot cancel spent salt");
        avsDirectory.cancelSalt(salt);
        cheats.stopPrank();
    }

    /// @notice Verifies that an operator cannot cancel the same salt twice
    function testFuzz_revert_whenCancellingSaltUsedToRegister(bytes32 salt) public {
        address operator = cheats.addr(delegationSignerPrivateKey);
        assertFalse(delegationManager.isOperator(operator), "bad test setup");
        _registerOperatorWithBaseDetails(operator);

        uint256 expiry = type(uint256).max;
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature =
            _getOperatorSignature(delegationSignerPrivateKey, operator, defaultAVS, salt, expiry);

        cheats.prank(defaultAVS);
        avsDirectory.registerOperatorToAVS(operator, operatorSignature);

        cheats.prank(operator);
        cheats.expectRevert("AVSDirectory.cancelSalt: cannot cancel spent salt");
        avsDirectory.cancelSalt(salt);
    }
}
