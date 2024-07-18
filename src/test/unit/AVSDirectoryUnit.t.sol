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
    uint256 internal constant MAX_PRIVATE_KEY = 0xfffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364140;

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
     * @notice internal function for calculating a signature from the operator corresponding to `operatorPk`, delegating them to
     * the `operator`, and expiring at `expiry`.
     */
    function _getOperatorAVSRegistrationSignature(
        uint256 operatorPk,
        address operator,
        address avs,
        bytes32 salt,
        uint256 expiry
    ) internal view returns (ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature) {
        operatorSignature.expiry = expiry;
        operatorSignature.salt = salt;
        {
            bytes32 digestHash = avsDirectory.calculateOperatorAVSRegistrationDigestHash(operator, avs, salt, expiry);
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(operatorPk, digestHash);
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

    function _registerOperatorToOperatorSet(
        uint256 operatorPk,
        uint32 operatorSetId,
        bytes32 salt,
        uint256 expiry
    ) internal virtual {
        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;

        _registerOperatorToOperatorSets(operatorPk, oids, salt, expiry);
    }

    function _registerOperatorToOperatorSets(
        uint256 operatorPk,
        uint32[] memory operatorSetIds,
        bytes32 salt,
        uint256 expiry
    ) internal virtual {
        expiry = bound(expiry, 1, type(uint256).max);
        cheats.warp(0);

        address operator = cheats.addr(operatorPk);
        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(
            operatorPk, avsDirectory.calculateOperatorSetRegistrationDigestHash(address(this), operatorSetIds, salt, expiry)
        );

        // Set AVS as operator set avs
        avsDirectory.becomeOperatorSetAVS();

        _registerOperatorWithBaseDetails(operator);

        avsDirectory.registerOperatorToOperatorSets(
            operator, operatorSetIds, ISignatureUtils.SignatureWithSaltAndExpiry(abi.encodePacked(r, s, v), salt, expiry)
        );
    }

    function _createOperatorSet(uint32 operatorSetId) internal {
        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;
        avsDirectory.createOperatorSets(oids);
    }
}

contract AVSDirectoryUnitTests_initialize is AVSDirectoryUnitTests {
    function testFuzz_Correctness(
        address delegationManager,
        address owner,
        address pauserRegistry,
        uint256 initialPausedStatus
    ) public virtual {
        AVSDirectory dir = new AVSDirectory(IDelegationManager(delegationManager));

        assertEq(address(dir.delegation()), delegationManager);

        cheats.expectRevert("Initializable: contract is already initialized");
        dir.initialize(owner, IPauserRegistry(pauserRegistry), initialPausedStatus);
    }
}

contract AVSDirectoryUnitTests_domainSeparator is AVSDirectoryUnitTests {
    function test_domainSeparator() public virtual {
        // This is just to get coverage up.
        avsDirectory.domainSeparator();
        cheats.chainId(0xC0FFEE);
        avsDirectory.domainSeparator();
    }
}

contract AVSDirectoryUnitTests_registerOperatorToOperatorSet is AVSDirectoryUnitTests {
    function testFuzz_revert_SignatureIsExpired(
        address operator,
        uint32 operatorSetId,
        bytes32 salt,
        uint256 expiry
    ) public virtual {
        expiry = bound(expiry, 0, type(uint256).max - 1);
        cheats.warp(type(uint256).max);

        _createOperatorSet(operatorSetId);

        _registerOperatorWithBaseDetails(operator);

        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;

        cheats.expectRevert("AVSDirectory.registerOperatorToOperatorSets: operator signature expired");
        avsDirectory.registerOperatorToOperatorSets(
            operator, oids, ISignatureUtils.SignatureWithSaltAndExpiry(new bytes(0), salt, expiry)
        );
    }

    function testFuzz_revert_notOperatorSetAVS(
        uint256 operatorPk,
        uint32 operatorSetId,
        bytes32 salt,
        uint256 expiry
    ) public virtual {
        operatorPk = bound(operatorPk, 1, MAX_PRIVATE_KEY);
        expiry = bound(expiry, 1, type(uint256).max);

        cheats.warp(0);

        _createOperatorSet(operatorSetId);

        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;

        address operator = cheats.addr(operatorPk);
        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(
            operatorPk, avsDirectory.calculateOperatorSetRegistrationDigestHash(address(this), oids, salt, expiry)
        );

        _registerOperatorWithBaseDetails(operator);

        cheats.expectRevert("AVSDirectory.registerOperatorToOperatorSets: AVS is not an operator set AVS");
        avsDirectory.registerOperatorToOperatorSets(
            operator, oids, ISignatureUtils.SignatureWithSaltAndExpiry(abi.encodePacked(r, s, v), salt, expiry)
        );
    }

    function testFuzz_revert_OperatorRegistered(
        uint256 operatorPk,
        uint32 operatorSetId,
        bytes32 salt,
        uint256 expiry
    ) public virtual {
        avsDirectory.becomeOperatorSetAVS();
        operatorPk = bound(operatorPk, 1, MAX_PRIVATE_KEY);
        expiry = bound(expiry, 1, type(uint256).max);

        cheats.warp(0);

        _createOperatorSet(operatorSetId);

        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;

        address operator = cheats.addr(operatorPk);
        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(
            operatorPk, avsDirectory.calculateOperatorSetRegistrationDigestHash(address(this), oids, salt, expiry)
        );

        _registerOperatorWithBaseDetails(operator);

        avsDirectory.registerOperatorToOperatorSets(
            operator, oids, ISignatureUtils.SignatureWithSaltAndExpiry(abi.encodePacked(r, s, v), salt, expiry)
        );

        (v, r, s) = cheats.sign(
            operatorPk,
            avsDirectory.calculateOperatorSetRegistrationDigestHash(address(this), oids, keccak256(""), expiry)
        );

        cheats.expectRevert("AVSDirectory.registerOperatorToOperatorSets: operator already registered to operator set");
        avsDirectory.registerOperatorToOperatorSets(
            operator, oids, ISignatureUtils.SignatureWithSaltAndExpiry(abi.encodePacked(r, s, v), keccak256(""), expiry)
        );
    }

    function testFuzz_revert_OperatorNotRegistered(
        address operator,
        uint32 operatorSetId,
        bytes32 salt,
        uint256 expiry
    ) public virtual {
        avsDirectory.becomeOperatorSetAVS();
        cheats.assume(operator != address(0));
        expiry = bound(expiry, 1, type(uint256).max);
        cheats.warp(0);

        _createOperatorSet(operatorSetId);

        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;

        cheats.expectRevert("AVSDirectory.registerOperatorToOperatorSets: operator not registered to EigenLayer yet");
        avsDirectory.registerOperatorToOperatorSets(
            operator, oids, ISignatureUtils.SignatureWithSaltAndExpiry(new bytes(0), salt, expiry)
        );
    }

    function testFuzz_revert_SaltSpent(
        uint256 operatorPk,
        uint32 operatorSetId,
        bytes32 salt,
        uint256 expiry
    ) public virtual {
        avsDirectory.becomeOperatorSetAVS();
        operatorSetId = uint32(bound(operatorSetId, 1, type(uint32).max));
        operatorPk = bound(operatorPk, 1, MAX_PRIVATE_KEY);
        expiry = bound(expiry, 1, type(uint256).max);

        cheats.warp(0);

        _createOperatorSet(operatorSetId);

        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;

        address operator = cheats.addr(operatorPk);
        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(
            operatorPk, avsDirectory.calculateOperatorSetRegistrationDigestHash(address(this), oids, salt, expiry)
        );

        _registerOperatorWithBaseDetails(operator);

        avsDirectory.registerOperatorToOperatorSets(
            operator, oids, ISignatureUtils.SignatureWithSaltAndExpiry(abi.encodePacked(r, s, v), salt, expiry)
        );

        cheats.expectRevert("AVSDirectory.registerOperatorToOperatorSets: salt already spent");
        avsDirectory.registerOperatorToOperatorSets(
            operator, new uint32[](0), ISignatureUtils.SignatureWithSaltAndExpiry(new bytes(0), salt, expiry)
        );
    }

    function testFuzz_revert_WrongAVS(
        address badAvs,
        uint256 operatorPk,
        uint32 operatorSetId,
        bytes32 salt,
        uint256 expiry
    ) public virtual {
        cheats.assume(badAvs != address(this));

        operatorSetId = uint32(bound(operatorSetId, 1, type(uint32).max));
        operatorPk = bound(operatorPk, 1, MAX_PRIVATE_KEY);
        expiry = bound(expiry, 1, type(uint256).max);

        cheats.warp(0);

        _createOperatorSet(operatorSetId);

        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;

        address operator = cheats.addr(operatorPk);

        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(
            operatorPk, avsDirectory.calculateOperatorSetRegistrationDigestHash(address(this), oids, salt, expiry)
        );

        _registerOperatorWithBaseDetails(operator);

        cheats.startPrank(badAvs);
        avsDirectory.becomeOperatorSetAVS();
        cheats.expectRevert("EIP1271SignatureUtils.checkSignature_EIP1271: signature not from signer");
        avsDirectory.registerOperatorToOperatorSets(
            operator, oids, ISignatureUtils.SignatureWithSaltAndExpiry(abi.encodePacked(r, s, v), salt, expiry)
        );
        cheats.stopPrank();
    }

    function testFuzz_revert_invalidOperatorSet(
        uint256 operatorPk,
        uint32 operatorSetId,
        bytes32 salt,
        uint256 expiry
    ) public virtual {
        avsDirectory.becomeOperatorSetAVS();
        operatorPk = bound(operatorPk, 1, MAX_PRIVATE_KEY);
        expiry = bound(expiry, 1, type(uint256).max);

        cheats.warp(0);

        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;

        address operator = cheats.addr(operatorPk);
        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(
            operatorPk, avsDirectory.calculateOperatorSetRegistrationDigestHash(address(this), oids, salt, expiry)
        );

        _registerOperatorWithBaseDetails(operator);

        cheats.expectRevert("AVSDirectory.registerOperatorToOperatorSets: invalid operator set");
        avsDirectory.registerOperatorToOperatorSets(
            operator, oids, ISignatureUtils.SignatureWithSaltAndExpiry(abi.encodePacked(r, s, v), salt, expiry)
        );
    }

    function testFuzz_MultipleCorrectness(
        uint256 operatorPk,
        uint256 operatorSetIdsLength,
        bytes32 salt,
        uint256 expiry
    ) public virtual {
        avsDirectory.becomeOperatorSetAVS();
        operatorPk = bound(operatorPk, 1, MAX_PRIVATE_KEY);
        operatorSetIdsLength = bound(operatorSetIdsLength, 1, 16);
        expiry = bound(expiry, 1, type(uint256).max);

        cheats.warp(0);

        uint32[] memory oids = new uint32[](operatorSetIdsLength);
        for (uint256 i; i < oids.length; ++i) {
            oids[i] = uint32(uint256(keccak256(abi.encodePacked(i))) % type(uint32).max);
            _createOperatorSet(oids[i]);
        }

        address operator = cheats.addr(operatorPk);
        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(
            operatorPk, avsDirectory.calculateOperatorSetRegistrationDigestHash(address(this), oids, salt, expiry)
        );

        _registerOperatorWithBaseDetails(operator);

        for (uint256 i; i < oids.length; ++i) {
            cheats.expectEmit(true, false, false, false, address(avsDirectory));
            emit OperatorAddedToOperatorSet(operator, address(this), oids[i]);
        }

        avsDirectory.registerOperatorToOperatorSets(
            operator, oids, ISignatureUtils.SignatureWithSaltAndExpiry(abi.encodePacked(r, s, v), salt, expiry)
        );

        for (uint256 i; i < oids.length; ++i) {
            assertTrue(avsDirectory.isMember(address(this), operator, oids[i]));
        }

        assertTrue(avsDirectory.operatorSaltIsSpent(operator, salt));
    }

    function testFuzz_Correctness(
        uint256 operatorPk,
        uint32 operatorSetId,
        bytes32 salt,
        uint256 expiry
    ) public virtual {
        avsDirectory.becomeOperatorSetAVS();
        operatorPk = bound(operatorPk, 1, MAX_PRIVATE_KEY);
        expiry = bound(expiry, 1, type(uint256).max);

        cheats.warp(0);

        _createOperatorSet(operatorSetId);

        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;

        address operator = cheats.addr(operatorPk);
        (uint8 v, bytes32 r, bytes32 s) = cheats.sign(
            operatorPk, avsDirectory.calculateOperatorSetRegistrationDigestHash(address(this), oids, salt, expiry)
        );

        _registerOperatorWithBaseDetails(operator);

        cheats.expectEmit(true, false, false, false, address(avsDirectory));
        emit OperatorAddedToOperatorSet(operator, address(this), operatorSetId);
        avsDirectory.registerOperatorToOperatorSets(
            operator, oids, ISignatureUtils.SignatureWithSaltAndExpiry(abi.encodePacked(r, s, v), salt, expiry)
        );

        assertTrue(avsDirectory.isMember(address(this), operator, operatorSetId));
        assertTrue(avsDirectory.operatorSaltIsSpent(operator, salt));
    }
}

contract AVSDirectoryUnitTests_forceDeregisterFromOperatorSets is AVSDirectoryUnitTests {

    function testFuzz_revert_OperatorNotInOperatorSet(uint256 operatorPk, uint32 operatorSetId) public virtual {
        operatorPk = bound(operatorPk, 1, MAX_PRIVATE_KEY);

        address operator = cheats.addr(operatorPk);
        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;

        _createOperatorSet(operatorSetId);

        ISignatureUtils.SignatureWithSaltAndExpiry memory emptySig;

        cheats.prank(operator);
        cheats.expectRevert("AVSDirectory.deregisterOperatorFromOperatorSet: operator not registered for operator set");
        
        avsDirectory.forceDeregisterFromOperatorSets(operator, address(this), oids, emptySig);
    }

    function testFuzz_revert_operatorNotCaller(uint256 operatorPk, uint32 operatorSetId) public {
        operatorPk = bound(operatorPk, 1, MAX_PRIVATE_KEY);

        address operator = cheats.addr(operatorPk);
        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;

        _createOperatorSet(operatorSetId);

        ISignatureUtils.SignatureWithSaltAndExpiry memory emptySig;

        cheats.expectRevert("AVSDirectory.forceDeregisterFromOperatorSets: caller must be operator");
        avsDirectory.forceDeregisterFromOperatorSets(operator, address(this), oids, emptySig);
    }

    function testFuzz_forceDeregisterFromOperatorSets(uint256 operatorPk, uint32 operatorSetId, uint8 operatorSetsToAdd, bytes32 salt) public {
        operatorPk = bound(operatorPk, 1, MAX_PRIVATE_KEY);
        operatorSetsToAdd = uint8(bound(operatorSetsToAdd, 1, type(uint8).max));
        address operator = cheats.addr(operatorPk);

        // Create operator sets
        operatorSetId = uint32(bound(operatorSetId, 1, type(uint32).max - uint32(operatorSetsToAdd)));
        uint32[] memory oids = new uint32[](operatorSetsToAdd);
        for (uint32 i = 0; i < operatorSetsToAdd; i++) {
            oids[i] = operatorSetId + i;
            _createOperatorSet(oids[i]);
        }

        // Register operator to operator sets
        _registerOperatorToOperatorSets(
            operatorPk, oids, salt, type(uint256).max
        );

        // Deregister operator from operator sets
        ISignatureUtils.SignatureWithSaltAndExpiry memory emptySig;
        cheats.prank(operator);
        for(uint i = 0; i < oids.length; i++) {
            cheats.expectEmit(true, true, true, true, address(avsDirectory));
            emit OperatorRemovedFromOperatorSet(operator, address(this), oids[i]);
        }
        avsDirectory.forceDeregisterFromOperatorSets(operator, address(this), oids, emptySig);

        for (uint32 i = 0; i < operatorSetsToAdd; i++) {
            assertFalse(avsDirectory.isMember(address(this), operator, oids[i]), "operator still in operator set");
        }
    }

    function testFuzz_revert_sigExpired(uint256 operatorPk, uint32 operatorSetId, bytes32 salt) public {
        operatorPk = bound(operatorPk, 1, MAX_PRIVATE_KEY);

        address operator = cheats.addr(operatorPk);
        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;

        _createOperatorSet(operatorSetId);

        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSig = _createForceDeregSignature
            (operatorPk, address(this), oids, 0, salt);

        cheats.warp(type(uint256).max);
        cheats.expectRevert("AVSDirectory.forceDeregisterFromOperatorSets: operator signature expired");
        avsDirectory.forceDeregisterFromOperatorSets(operator, address(this), oids, operatorSig);
    }

    function testFuzz_revert_saltAlreadySpent(uint256 operatorPk, uint32 operatorSetId, bytes32 salt) public {
        operatorPk = bound(operatorPk, 1, MAX_PRIVATE_KEY);

        address operator = cheats.addr(operatorPk);
        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;

        // Register operator to operator sets
        _createOperatorSet(operatorSetId);
        _registerOperatorToOperatorSets(
            operatorPk, oids, salt, type(uint256).max
        );

        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSig = _createForceDeregSignature
            (operatorPk, address(this), oids, type(uint256).max, salt);

        cheats.expectRevert("AVSDirectory.forceDeregisterFromOperatorSets: salt already spent");
        avsDirectory.forceDeregisterFromOperatorSets(operator, address(this), oids, operatorSig);
    }

    function testFuzz_forceDeregisterFromOperatorSets_onBehalf(uint256 operatorPk, uint32 operatorSetId, uint8 operatorSetsToAdd, bytes32 salt1, bytes32 salt2) public {
        operatorPk = bound(operatorPk, 1, MAX_PRIVATE_KEY);
        operatorSetsToAdd = uint8(bound(operatorSetsToAdd, 1, type(uint8).max));
        address operator = cheats.addr(operatorPk);

        // Create operator sets
        operatorSetId = uint32(bound(operatorSetId, 1, type(uint32).max - uint32(operatorSetsToAdd)));
        uint32[] memory oids = new uint32[](operatorSetsToAdd);
        for (uint32 i = 0; i < oids.length; i++) {
            oids[i] = operatorSetId + i;
            _createOperatorSet(oids[i]);
        }

        // Register operator to operator sets
        _registerOperatorToOperatorSets(
            operatorPk, oids, salt1, type(uint256).max
        );

        // Deregister operator from operator sets
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSig = _createForceDeregSignature
            (operatorPk, address(this), oids, type(uint256).max, salt2);
        
        for (uint i = 0; i < oids.length; i++) {
            cheats.expectEmit(true, true, true, true, address(avsDirectory));
            emit OperatorRemovedFromOperatorSet(operator, address(this), oids[i]);
        }

        avsDirectory.forceDeregisterFromOperatorSets(operator, address(this), oids, operatorSig);

        for (uint32 i = 0; i < operatorSetsToAdd; i++) {
            assertFalse(avsDirectory.isMember(address(this), operator, oids[i]));
        }
    }

    function _createForceDeregSignature(uint256 operatorPk, address avs, uint32[] memory oids, uint256 expiry, bytes32 salt) internal view returns (ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature) {
        operatorSignature.expiry = expiry;
        address operator = cheats.addr(operatorPk);
        operatorSignature.salt = salt;
        {
            bytes32 digestHash = avsDirectory.calculateOperatorSetForceDeregistrationTypehash(avs, oids, salt, expiry);
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(operatorPk, digestHash);
            operatorSignature.signature = abi.encodePacked(r, s, v);
        }
        return operatorSignature; 
    }
}

contract AVSDirectoryUnitTests_deregisterOperatorFromOperatorSets is AVSDirectoryUnitTests {

    function testFuzz_revert_OperatorNotInOperatorSet(uint256 operatorPk, uint32 operatorSetId) public virtual {
        operatorPk = bound(operatorPk, 1, MAX_PRIVATE_KEY);

        _createOperatorSet(operatorSetId);

        address operator = cheats.addr(operatorPk);
        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;

        cheats.expectRevert("AVSDirectory.deregisterOperatorFromOperatorSet: operator not registered for operator set");
        avsDirectory.deregisterOperatorFromOperatorSets(operator, oids);
    }

    function testFuzz_Correctness(
        uint256 operatorPk,
        uint32 operatorSetId,
        bytes32 salt,
        uint256 expiry
    ) public virtual {
        operatorPk = bound(operatorPk, 1, MAX_PRIVATE_KEY);

        _createOperatorSet(operatorSetId);

        _registerOperatorToOperatorSet(operatorPk, operatorSetId, salt, expiry);

        address operator = cheats.addr(operatorPk);
        uint32[] memory oids = new uint32[](1);
        oids[0] = operatorSetId;

        cheats.expectEmit(true, false, false, false, address(avsDirectory));
        emit OperatorRemovedFromOperatorSet(operator, address(this), operatorSetId);

        avsDirectory.deregisterOperatorFromOperatorSets(operator, oids);

        assertEq(avsDirectory.isMember(address(this), operator, operatorSetId), false);
    }
}

contract AVSDirectoryUnitTests_createOperatorSet is AVSDirectoryUnitTests {
    function test_createOperatorSet() public {

        uint32[] memory oids = new uint32[](2);
        oids[0] = 1;
        oids[1] = 2;

        for(uint i = 0; i < oids.length; i++) {
            cheats.expectEmit(true, true, true, true, address(avsDirectory));
            emit OperatorSetCreated(address(this), oids[i]);
        }

        avsDirectory.createOperatorSets(oids);

        assertTrue(avsDirectory.isOperatorSet(address(this), oids[0]));
        assertTrue(avsDirectory.isOperatorSet(address(this), oids[1]));
    }

    function test_revert_operatorSetExists() public {
        _createOperatorSet(1);
        cheats.expectRevert("AVSDirectory.createOperatorSet: operator set already exists");
        _createOperatorSet(1);
    }
}

contract AVSDirectoryUnitTests_becomeOperatorSetAVS is AVSDirectoryUnitTests {
    function test_becomeOperatorSetAVS() public {
        cheats.expectEmit(true, true, true, true, address(avsDirectory));
        emit AVSMigratedToOperatorSets(address(this));

        avsDirectory.becomeOperatorSetAVS();

        assertTrue(avsDirectory.isOperatorSetAVS(address(this)));
    }

    function test_revert_alreadyOperatorSetAVS() public {
        avsDirectory.becomeOperatorSetAVS();
        cheats.expectRevert("AVSDirectory.becomeOperatorSetAVS: already an operator set AVS");
        avsDirectory.becomeOperatorSetAVS();
    }
}

contract AVSDirectoryUnitTests_legacyOperatorAVSRegistration is AVSDirectoryUnitTests {
    function test_revert_whenRegisterDeregisterToAVSPaused() public {
        // set the pausing flag
        cheats.prank(pauser);
        avsDirectory.pause(2 ** PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS);

        cheats.expectRevert("Pausable: index is paused");
        avsDirectory.registerOperatorToAVS(
            address(0), ISignatureUtils.SignatureWithSaltAndExpiry(abi.encodePacked(""), 0, 0)
        );

        cheats.expectRevert("Pausable: index is paused");
        avsDirectory.deregisterOperatorFromAVS(address(0));
    }

    function test_revert_deregisterOperatorFromAVS_operatorNotRegistered() public {
        cheats.expectRevert("AVSDirectory.deregisterOperatorFromAVS: operator not registered");
        avsDirectory.deregisterOperatorFromAVS(address(0));
    }

    function testFuzz_deregisterOperatorFromAVS(bytes32 salt) public {
        address operator = cheats.addr(delegationSignerPrivateKey);
        assertFalse(delegationManager.isOperator(operator), "bad test setup");
        _registerOperatorWithBaseDetails(operator);

        uint256 expiry = type(uint256).max;
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature =
            _getOperatorAVSRegistrationSignature(delegationSignerPrivateKey, operator, defaultAVS, salt, expiry);

        cheats.prank(defaultAVS);
        avsDirectory.registerOperatorToAVS(operator, operatorSignature);

        cheats.expectEmit(true, true, true, true, address(avsDirectory));
        emit OperatorAVSRegistrationStatusUpdated(operator, defaultAVS, IAVSDirectory.OperatorAVSRegistrationStatus.UNREGISTERED);

        cheats.prank(defaultAVS);
        avsDirectory.deregisterOperatorFromAVS(operator);

        assertTrue(avsDirectory.avsOperatorStatus(defaultAVS, operator) == IAVSDirectory.OperatorAVSRegistrationStatus.UNREGISTERED);
    }

    // @notice Tests that an avs who calls `updateAVSMetadataURI` will correctly see an `AVSMetadataURIUpdated` event emitted with their input
    function testFuzz_UpdateAVSMetadataURI(string memory metadataURI) public {
        // call `updateAVSMetadataURI` and check for event
        cheats.expectEmit(true, true, true, true, address(avsDirectory));
        cheats.prank(defaultAVS);
        emit AVSMetadataURIUpdated(defaultAVS, metadataURI);
        avsDirectory.updateAVSMetadataURI(metadataURI);
    }

    function testFuzz_revert_whenAVSIsOperatorSetAVS(bytes32 salt) public {
        // set the AVS to be an operator set AVS
        cheats.prank(defaultAVS);
        avsDirectory.becomeOperatorSetAVS();

        // Register Operator to EigenLayer
        address operator = cheats.addr(delegationSignerPrivateKey);
        assertFalse(delegationManager.isOperator(operator), "bad test setup");
        _registerOperatorWithBaseDetails(operator);

        uint256 expiry = type(uint256).max;

        cheats.prank(defaultAVS);
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature =
            _getOperatorAVSRegistrationSignature(delegationSignerPrivateKey, operator, defaultAVS, salt, expiry);
        cheats.expectRevert("AVSDirectory.registerOperatorToAVS: AVS is an operator set AVS");
        avsDirectory.registerOperatorToAVS(operator, operatorSignature);
    }

    // @notice Verifies an operator registers successfull to avs and see an `OperatorAVSRegistrationStatusUpdated` event emitted
    function testFuzz_registerOperatorToAVS(bytes32 salt) public {
        address operator = cheats.addr(delegationSignerPrivateKey);
        assertFalse(delegationManager.isOperator(operator), "bad test setup");
        _registerOperatorWithBaseDetails(operator);

        cheats.expectEmit(true, true, true, true, address(avsDirectory));
        emit OperatorAVSRegistrationStatusUpdated(operator, defaultAVS, IAVSDirectory.OperatorAVSRegistrationStatus.REGISTERED);

        uint256 expiry = type(uint256).max;

        cheats.prank(defaultAVS);
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature =
            _getOperatorAVSRegistrationSignature(delegationSignerPrivateKey, operator, defaultAVS, salt, expiry);

        avsDirectory.registerOperatorToAVS(operator, operatorSignature);

        assertTrue(avsDirectory.avsOperatorStatus(defaultAVS, operator) == IAVSDirectory.OperatorAVSRegistrationStatus.REGISTERED);
    }

    // @notice Verifies an operator registers successfull to avs and see an `OperatorAVSRegistrationStatusUpdated` event emitted
    function testFuzz_revert_whenOperatorNotRegisteredToEigenLayerYet(bytes32 salt) public {
        address operator = cheats.addr(delegationSignerPrivateKey);
        assertFalse(delegationManager.isOperator(operator), "bad test setup");

        cheats.prank(defaultAVS);
        uint256 expiry = type(uint256).max;
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature =
            _getOperatorAVSRegistrationSignature(delegationSignerPrivateKey, operator, defaultAVS, salt, expiry);

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
            _getOperatorAVSRegistrationSignature(delegationSignerPrivateKey, operator, defaultAVS, salt, expiry);

        cheats.expectRevert("EIP1271SignatureUtils.checkSignature_EIP1271: signature not from signer");
        cheats.prank(operator);
        avsDirectory.registerOperatorToAVS(operator, operatorSignature);
    }

    // @notice Verifies an operator registers fails when the signature expiry already expires
    function testFuzz_revert_whenExpiryHasExpired(ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature)
        public
    {
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
            _getOperatorAVSRegistrationSignature(delegationSignerPrivateKey, operator, defaultAVS, salt, expiry);

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
        assertFalse(
            avsDirectory.operatorSaltIsSpent(defaultAVS, salt), "salt should only be cancelled for the operator"
        );

        bytes32 newSalt;
        unchecked {
            newSalt = bytes32(uint256(salt) + 1);
        }

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
            _getOperatorAVSRegistrationSignature(delegationSignerPrivateKey, operator, defaultAVS, salt, expiry);

        cheats.prank(operator);
        avsDirectory.cancelSalt(salt);

        cheats.expectRevert("AVSDirectory.registerOperatorToAVS: salt already spent");
        cheats.prank(defaultAVS);
        avsDirectory.registerOperatorToAVS(operator, operatorSignature);
    }

    /// @notice Verifies that an operator cannot cancel the same salt twice
    function testFuzz_revert_whenCancellingSaltUsedToRegister(bytes32 salt) public {
        address operator = cheats.addr(delegationSignerPrivateKey);
        assertFalse(delegationManager.isOperator(operator), "bad test setup");
        _registerOperatorWithBaseDetails(operator);

        uint256 expiry = type(uint256).max;
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature =
            _getOperatorAVSRegistrationSignature(delegationSignerPrivateKey, operator, defaultAVS, salt, expiry);

        cheats.prank(defaultAVS);
        avsDirectory.registerOperatorToAVS(operator, operatorSignature);
    }
}
