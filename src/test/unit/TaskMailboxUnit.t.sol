// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Test, console, Vm} from "forge-std/Test.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ITransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {IBN254CertificateVerifier, IBN254CertificateVerifierTypes} from "src/contracts/interfaces/IBN254CertificateVerifier.sol";
import {IECDSACertificateVerifier, IECDSACertificateVerifierTypes} from "src/contracts/interfaces/IECDSACertificateVerifier.sol";
import {OperatorSet, OperatorSetLib} from "src/contracts/libraries/OperatorSetLib.sol";
import {BN254} from "src/contracts/libraries/BN254.sol";
import {IKeyRegistrarTypes} from "src/contracts/interfaces/IKeyRegistrar.sol";
import {TaskMailbox} from "src/contracts/avs/task/TaskMailbox.sol";
import {ITaskMailbox, ITaskMailboxTypes, ITaskMailboxErrors, ITaskMailboxEvents} from "src/contracts/interfaces/ITaskMailbox.sol";
import {IAVSTaskHook} from "src/contracts/interfaces/IAVSTaskHook.sol";

import {MockAVSTaskHook} from "src/test/mocks/MockAVSTaskHook.sol";
import {MockBN254CertificateVerifier} from "src/test/mocks/MockBN254CertificateVerifier.sol";
import {MockBN254CertificateVerifierFailure} from "src/test/mocks/MockBN254CertificateVerifierFailure.sol";
import {MockECDSACertificateVerifier} from "src/test/mocks/MockECDSACertificateVerifier.sol";
import {MockECDSACertificateVerifierFailure} from "src/test/mocks/MockECDSACertificateVerifierFailure.sol";
import {MockSimpleERC20} from "src/test/mocks/MockSimpleERC20.sol";
import {AVSTaskHookReentrantAttacker} from "src/test/mocks/AVSTaskHookReentrantAttacker.sol";

contract TaskMailboxUnitTests is Test, ITaskMailboxTypes, ITaskMailboxErrors, ITaskMailboxEvents {
    using OperatorSetLib for OperatorSet;

    // Contracts
    TaskMailbox public taskMailbox;
    ProxyAdmin public proxyAdmin;
    MockAVSTaskHook public mockTaskHook;
    MockBN254CertificateVerifier public mockBN254CertificateVerifier;
    MockECDSACertificateVerifier public mockECDSACertificateVerifier;
    MockSimpleERC20 public mockToken;

    // Test addresses
    address public avs = address(0x1);
    address public avs2 = address(0x2);
    address public feeCollector = address(0x3);
    address public refundCollector = address(0x4);
    address public creator = address(0x5);
    address public aggregator = address(0x6);
    address public owner = address(0x7);
    address public feeSplitCollector = address(0x8);

    // Test operator set IDs
    uint32 public executorOperatorSetId = 1;
    uint32 public executorOperatorSetId2 = 2;

    // Test config values
    uint96 public taskSLA = 60 seconds;
    uint16 public stakeProportionThreshold = 6667; // 66.67%
    uint96 public avsFee = 1 ether;

    function setUp() public virtual {
        // Deploy mock contracts
        mockTaskHook = new MockAVSTaskHook();
        mockBN254CertificateVerifier = new MockBN254CertificateVerifier();
        mockECDSACertificateVerifier = new MockECDSACertificateVerifier();
        mockToken = new MockSimpleERC20();

        // Deploy TaskMailbox with proxy pattern
        proxyAdmin = new ProxyAdmin();
        TaskMailbox taskMailboxImpl = new TaskMailbox(address(mockBN254CertificateVerifier), address(mockECDSACertificateVerifier), "1.0.0");
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(taskMailboxImpl),
            address(proxyAdmin),
            abi.encodeWithSelector(TaskMailbox.initialize.selector, owner, 0, feeSplitCollector)
        );
        taskMailbox = TaskMailbox(address(proxy));

        // Give creator some tokens and approve TaskMailbox
        mockToken.mint(creator, 1000 ether);
        vm.prank(creator);
        mockToken.approve(address(taskMailbox), type(uint).max);
    }

    function _createValidTaskParams() internal view returns (TaskParams memory) {
        return TaskParams({
            refundCollector: refundCollector,
            executorOperatorSet: OperatorSet(avs, executorOperatorSetId),
            payload: bytes("test payload")
        });
    }

    function _createValidExecutorOperatorSetTaskConfig() internal view returns (ExecutorOperatorSetTaskConfig memory) {
        return ExecutorOperatorSetTaskConfig({
            taskHook: IAVSTaskHook(address(mockTaskHook)),
            taskSLA: taskSLA,
            feeToken: IERC20(address(mockToken)),
            curveType: IKeyRegistrarTypes.CurveType.BN254,
            feeCollector: feeCollector,
            consensus: Consensus({consensusType: ConsensusType.STAKE_PROPORTION_THRESHOLD, value: abi.encode(stakeProportionThreshold)}),
            taskMetadata: bytes("test metadata")
        });
    }

    /**
     * @notice Get the ECDSA certificate digest for a given result
     * @param referenceTimestamp The reference timestamp for the certificate
     * @param result The result data
     * @return The certificate digest
     */
    function _getECDSACertificateDigest(uint32 referenceTimestamp, bytes memory result) internal view returns (bytes32) {
        // For testing, we'll use a simple digest calculation that matches what the mock expects
        // In the real implementation, this would call IECDSACertificateVerifier.calculateCertificateDigest
        return keccak256(abi.encode(referenceTimestamp, keccak256(result)));
    }

    function _createValidBN254Certificate(bytes32 messageHash, uint96 referenceTimestamp)
        internal
        view
        returns (IBN254CertificateVerifierTypes.BN254Certificate memory)
    {
        return IBN254CertificateVerifierTypes.BN254Certificate({
            referenceTimestamp: uint32(referenceTimestamp),
            messageHash: messageHash,
            signature: BN254.G1Point(1, 2), // Non-zero signature
            apk: BN254.G2Point([uint(1), uint(2)], [uint(3), uint(4)]),
            nonSignerWitnesses: new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](0)
        });
    }

    function _createValidECDSACertificate(bytes32 messageHash, uint96 referenceTimestamp)
        internal
        view
        returns (IECDSACertificateVerifierTypes.ECDSACertificate memory)
    {
        return IECDSACertificateVerifierTypes.ECDSACertificate({
            referenceTimestamp: uint32(referenceTimestamp),
            messageHash: messageHash,
            sig: bytes("0x1234567890abcdef") // Non-empty signature
        });
    }

    /**
     * @notice Create a valid BN254 certificate for a given result
     * @param result The result data to create certificate for
     * @param referenceTimestamp The reference timestamp
     * @return The BN254 certificate
     */
    function _createValidBN254CertificateForResult(bytes memory result, uint96 referenceTimestamp)
        internal
        view
        returns (IBN254CertificateVerifierTypes.BN254Certificate memory)
    {
        return _createValidBN254Certificate(keccak256(result), referenceTimestamp);
    }

    /**
     * @notice Get the reference timestamp for a task (helper for tests)
     * @param _taskHash The task hash
     * @return The reference timestamp stored in the task
     */
    function _getTaskReferenceTimestamp(bytes32 _taskHash) internal view returns (uint32) {
        Task memory task = taskMailbox.getTaskInfo(_taskHash);
        return task.operatorTableReferenceTimestamp;
    }

    /**
     * @notice Create a valid ECDSA certificate for a given result
     * @param result The result data to create certificate for
     * @param referenceTimestamp The reference timestamp
     * @return The ECDSA certificate
     */
    function _createValidECDSACertificateForResult(bytes memory result, uint96 referenceTimestamp)
        internal
        view
        returns (IECDSACertificateVerifierTypes.ECDSACertificate memory)
    {
        bytes32 digest = _getECDSACertificateDigest(uint32(referenceTimestamp), result);
        return _createValidECDSACertificate(digest, referenceTimestamp);
    }
}

contract TaskMailboxUnitTests_Constructor is TaskMailboxUnitTests {
    function test_Constructor_WithCertificateVerifiers() public {
        address bn254Verifier = address(0x1234);
        address ecdsaVerifier = address(0x5678);

        // Deploy with proxy pattern
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        TaskMailbox taskMailboxImpl = new TaskMailbox(bn254Verifier, ecdsaVerifier, "1.0.0");
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(taskMailboxImpl),
            address(proxyAdmin),
            abi.encodeWithSelector(TaskMailbox.initialize.selector, owner, 0, feeSplitCollector)
        );
        TaskMailbox newTaskMailbox = TaskMailbox(address(proxy));

        assertEq(newTaskMailbox.BN254_CERTIFICATE_VERIFIER(), bn254Verifier);
        assertEq(newTaskMailbox.ECDSA_CERTIFICATE_VERIFIER(), ecdsaVerifier);
        assertEq(newTaskMailbox.version(), "1.0.0");
        assertEq(newTaskMailbox.owner(), owner);
        assertEq(newTaskMailbox.feeSplit(), 0);
        assertEq(newTaskMailbox.feeSplitCollector(), feeSplitCollector);
    }
}

// Test contract for registerExecutorOperatorSet
contract TaskMailboxUnitTests_registerExecutorOperatorSet is TaskMailboxUnitTests {
    function testFuzz_registerExecutorOperatorSet(address fuzzAvs, uint32 fuzzOperatorSetId, bool fuzzIsRegistered) public {
        // Skip if fuzzAvs is the proxy admin to avoid proxy admin access issues
        vm.assume(fuzzAvs != address(proxyAdmin));
        OperatorSet memory operatorSet = OperatorSet(fuzzAvs, fuzzOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();

        // Set config first (requirement for registerExecutorOperatorSet)
        vm.prank(fuzzAvs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // If unregistering, expect event
        if (!fuzzIsRegistered) {
            vm.expectEmit(true, true, true, true, address(taskMailbox));
            emit ExecutorOperatorSetRegistered(fuzzAvs, fuzzAvs, fuzzOperatorSetId, fuzzIsRegistered);

            // Register operator set
            vm.prank(fuzzAvs);
            taskMailbox.registerExecutorOperatorSet(operatorSet, fuzzIsRegistered);
        }

        // Verify registration status
        assertEq(taskMailbox.isExecutorOperatorSetRegistered(operatorSet.key()), fuzzIsRegistered);
    }

    function test_registerExecutorOperatorSet_Unregister() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();

        // Set config (this automatically registers the operator set)
        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);
        assertTrue(taskMailbox.isExecutorOperatorSetRegistered(operatorSet.key()));

        // Then unregister
        vm.expectEmit(true, true, true, true, address(taskMailbox));
        emit ExecutorOperatorSetRegistered(avs, avs, executorOperatorSetId, false);

        vm.prank(avs);
        taskMailbox.registerExecutorOperatorSet(operatorSet, false);
        assertFalse(taskMailbox.isExecutorOperatorSetRegistered(operatorSet.key()));
    }

    function test_Revert_registerExecutorOperatorSet_ConfigNotSet() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);

        vm.prank(avs);
        vm.expectRevert(ExecutorOperatorSetTaskConfigNotSet.selector);
        taskMailbox.registerExecutorOperatorSet(operatorSet, true);
    }

    function test_Revert_registerExecutorOperatorSet_InvalidOperatorSetOwner() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();

        // Set config as avs
        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Try to register as different address
        vm.prank(avs2);
        vm.expectRevert(InvalidOperatorSetOwner.selector);
        taskMailbox.registerExecutorOperatorSet(operatorSet, false);
    }
}

// Test contract for setExecutorOperatorSetTaskConfig
contract TaskMailboxUnitTests_setExecutorOperatorSetTaskConfig is TaskMailboxUnitTests {
    function test_Revert_InvalidOperatorSetOwner() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();

        // Try to set config as wrong address
        vm.prank(avs2);
        vm.expectRevert(InvalidOperatorSetOwner.selector);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);
    }

    function testFuzz_setExecutorOperatorSetTaskConfig(
        address fuzzCertificateVerifier,
        address fuzzTaskHook,
        address fuzzFeeToken,
        address fuzzFeeCollector,
        uint96 fuzzTaskSLA,
        uint16 fuzzStakeProportionThreshold,
        bytes memory fuzzTaskMetadata
    ) public {
        // Bound inputs
        vm.assume(fuzzCertificateVerifier != address(0));
        vm.assume(fuzzTaskHook != address(0));
        vm.assume(fuzzTaskSLA > 0);
        // Bound stake proportion threshold to valid range (0-10000 basis points)
        fuzzStakeProportionThreshold = uint16(bound(fuzzStakeProportionThreshold, 0, 10_000));

        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);

        ExecutorOperatorSetTaskConfig memory config = ExecutorOperatorSetTaskConfig({
            taskHook: IAVSTaskHook(fuzzTaskHook),
            taskSLA: fuzzTaskSLA,
            feeToken: IERC20(fuzzFeeToken),
            curveType: IKeyRegistrarTypes.CurveType.BN254,
            feeCollector: fuzzFeeCollector,
            consensus: Consensus({consensusType: ConsensusType.STAKE_PROPORTION_THRESHOLD, value: abi.encode(fuzzStakeProportionThreshold)}),
            taskMetadata: fuzzTaskMetadata
        });

        // Since setExecutorOperatorSetTaskConfig always registers if not already registered,
        // we expect both events every time for a new operator set
        // Note: The contract emits config event first, then registration event

        // Expect config event first
        vm.expectEmit(true, true, true, true, address(taskMailbox));
        emit ExecutorOperatorSetTaskConfigSet(avs, avs, executorOperatorSetId, config);

        // Expect registration event second
        vm.expectEmit(true, true, true, true, address(taskMailbox));
        emit ExecutorOperatorSetRegistered(avs, avs, executorOperatorSetId, true);

        // Set config
        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Verify config was set
        ExecutorOperatorSetTaskConfig memory retrievedConfig = taskMailbox.getExecutorOperatorSetTaskConfig(operatorSet);

        assertEq(uint8(retrievedConfig.curveType), uint8(IKeyRegistrarTypes.CurveType.BN254));
        assertEq(address(retrievedConfig.taskHook), fuzzTaskHook);
        assertEq(address(retrievedConfig.feeToken), fuzzFeeToken);
        assertEq(retrievedConfig.feeCollector, fuzzFeeCollector);
        assertEq(retrievedConfig.taskSLA, fuzzTaskSLA);
        // Verify consensus configuration
        assertEq(uint8(retrievedConfig.consensus.consensusType), uint8(ConsensusType.STAKE_PROPORTION_THRESHOLD));
        uint16 decodedThreshold = abi.decode(retrievedConfig.consensus.value, (uint16));
        assertEq(decodedThreshold, fuzzStakeProportionThreshold);
        assertEq(retrievedConfig.taskMetadata, fuzzTaskMetadata);

        // Verify operator set is registered
        assertTrue(taskMailbox.isExecutorOperatorSetRegistered(operatorSet.key()));
    }

    function test_setExecutorOperatorSetTaskConfig_AlreadyRegistered() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();

        // First set config (which auto-registers)
        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);
        assertTrue(taskMailbox.isExecutorOperatorSetRegistered(operatorSet.key()));

        // Update config again
        config.taskSLA = 120;

        // Should not emit registration event since already registered
        vm.recordLogs();

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Verify only one event was emitted (config set, not registration)
        Vm.Log[] memory entries = vm.getRecordedLogs();
        assertEq(entries.length, 1);
        assertEq(
            entries[0].topics[0],
            keccak256("ExecutorOperatorSetTaskConfigSet(address,address,uint32,(address,uint96,address,uint8,address,(uint8,bytes),bytes))")
        );

        // Verify the config was updated
        ExecutorOperatorSetTaskConfig memory updatedConfig = taskMailbox.getExecutorOperatorSetTaskConfig(operatorSet);
        assertEq(updatedConfig.taskSLA, 120);
    }

    function test_Revert_WhenCurveTypeIsNone() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.curveType = IKeyRegistrarTypes.CurveType.NONE;

        // Expecting revert due to accessing zero address certificate verifier
        vm.prank(avs);
        vm.expectRevert();
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);
    }

    function test_Revert_WhenTaskHookIsZero() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.taskHook = IAVSTaskHook(address(0));

        vm.prank(avs);
        vm.expectRevert(InvalidAddressZero.selector);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);
    }

    function test_Revert_WhenTaskSLAIsZero() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.taskSLA = 0;

        vm.prank(avs);
        vm.expectRevert(TaskSLAIsZero.selector);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);
    }

    function test_Revert_WhenConsensusValueInvalid_EmptyBytes() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.consensus = Consensus({consensusType: ConsensusType.STAKE_PROPORTION_THRESHOLD, value: bytes("")});

        vm.prank(avs);
        vm.expectRevert(InvalidConsensusValue.selector);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);
    }

    function test_Revert_WhenConsensusValueInvalid_WrongLength() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.consensus = Consensus({
            consensusType: ConsensusType.STAKE_PROPORTION_THRESHOLD,
            value: abi.encodePacked(uint8(50)) // Wrong size - should be 32 bytes
        });

        vm.prank(avs);
        vm.expectRevert(InvalidConsensusValue.selector);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);
    }

    function test_Revert_WhenConsensusValueInvalid_ExceedsMaximum() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.consensus = Consensus({
            consensusType: ConsensusType.STAKE_PROPORTION_THRESHOLD,
            value: abi.encode(uint16(10_001)) // Exceeds 10000 basis points
        });

        vm.prank(avs);
        vm.expectRevert(InvalidConsensusValue.selector);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);
    }

    function test_ConsensusZeroThreshold() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.consensus = Consensus({
            consensusType: ConsensusType.STAKE_PROPORTION_THRESHOLD,
            value: abi.encode(uint16(0)) // Zero threshold is valid
        });

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Verify config was set
        ExecutorOperatorSetTaskConfig memory retrievedConfig = taskMailbox.getExecutorOperatorSetTaskConfig(operatorSet);
        uint16 decodedThreshold = abi.decode(retrievedConfig.consensus.value, (uint16));
        assertEq(decodedThreshold, 0);
    }

    function test_ConsensusMaxThreshold() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.consensus = Consensus({
            consensusType: ConsensusType.STAKE_PROPORTION_THRESHOLD,
            value: abi.encode(uint16(10_000)) // Maximum 100%
        });

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Verify config was set
        ExecutorOperatorSetTaskConfig memory retrievedConfig = taskMailbox.getExecutorOperatorSetTaskConfig(operatorSet);
        uint16 decodedThreshold = abi.decode(retrievedConfig.consensus.value, (uint16));
        assertEq(decodedThreshold, 10_000);
    }

    function test_Revert_WhenConsensusTypeIsNone() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.consensus = Consensus({
            consensusType: ConsensusType.NONE,
            value: bytes("") // Empty value for NONE type
        });

        vm.prank(avs);
        vm.expectRevert(InvalidConsensusType.selector);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);
    }
}

// Test contract for setFeeSplit
contract TaskMailboxUnitTests_setFeeSplit is TaskMailboxUnitTests {
    function test_SetFeeSplit() public {
        uint16 newFeeSplit = 2000; // 20%

        vm.expectEmit(true, true, true, true);
        emit FeeSplitSet(newFeeSplit);

        vm.prank(owner);
        taskMailbox.setFeeSplit(newFeeSplit);
        assertEq(taskMailbox.feeSplit(), newFeeSplit);
    }

    function test_SetFeeSplit_MaxValue() public {
        uint16 maxFeeSplit = 10_000; // 100%

        vm.expectEmit(true, true, true, true);
        emit FeeSplitSet(maxFeeSplit);

        vm.prank(owner);
        taskMailbox.setFeeSplit(maxFeeSplit);
        assertEq(taskMailbox.feeSplit(), maxFeeSplit);
    }

    function test_Revert_SetFeeSplit_NotOwner() public {
        vm.prank(address(0x999));
        vm.expectRevert("Ownable: caller is not the owner");
        taskMailbox.setFeeSplit(1000);
    }

    function test_Revert_SetFeeSplit_ExceedsMax() public {
        vm.prank(owner);
        vm.expectRevert(InvalidFeeSplit.selector);
        taskMailbox.setFeeSplit(10_001); // > 100%
    }
}

// Test contract for setFeeSplitCollector
contract TaskMailboxUnitTests_setFeeSplitCollector is TaskMailboxUnitTests {
    function test_SetFeeSplitCollector() public {
        address newCollector = address(0x123);

        vm.expectEmit(true, true, true, true);
        emit FeeSplitCollectorSet(newCollector);

        vm.prank(owner);
        taskMailbox.setFeeSplitCollector(newCollector);
        assertEq(taskMailbox.feeSplitCollector(), newCollector);
    }

    function test_Revert_SetFeeSplitCollector_NotOwner() public {
        vm.prank(address(0x999));
        vm.expectRevert("Ownable: caller is not the owner");
        taskMailbox.setFeeSplitCollector(address(0x123));
    }

    function test_Revert_SetFeeSplitCollector_ZeroAddress() public {
        vm.prank(owner);
        vm.expectRevert(InvalidAddressZero.selector);
        taskMailbox.setFeeSplitCollector(address(0));
    }
}

// Test contract for createTask
contract TaskMailboxUnitTests_createTask is TaskMailboxUnitTests {
    function setUp() public override {
        super.setUp();

        // Set up executor operator set task config
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);
    }

    function testFuzz_createTask(address fuzzRefundCollector, uint96 fuzzAvsFee, bytes memory fuzzPayload) public {
        // Bound inputs
        vm.assume(fuzzPayload.length > 0);
        // We create two tasks in this test, so need at least 2x the fee
        vm.assume(fuzzAvsFee <= mockToken.balanceOf(creator) / 2);
        // If there's a fee, refund collector cannot be zero address
        if (fuzzAvsFee > 0) vm.assume(fuzzRefundCollector != address(0));

        // Set the mock hook to return the fuzzed fee
        mockTaskHook.setDefaultFee(fuzzAvsFee);

        TaskParams memory taskParams = TaskParams({
            refundCollector: fuzzRefundCollector,
            executorOperatorSet: OperatorSet(avs, executorOperatorSetId),
            payload: fuzzPayload
        });

        // First task will have count 0
        uint expectedTaskCount = 0;
        bytes32 expectedTaskHash = keccak256(abi.encode(expectedTaskCount, address(taskMailbox), block.chainid, taskParams));

        // Expect Transfer event for fee transfer from creator to taskMailbox
        if (fuzzAvsFee > 0) {
            vm.expectEmit(true, true, false, true, address(mockToken));
            emit IERC20.Transfer(creator, address(taskMailbox), fuzzAvsFee);
        }

        // Expect event
        vm.expectEmit(true, true, true, true, address(taskMailbox));
        emit TaskCreated(
            creator,
            expectedTaskHash,
            avs,
            executorOperatorSetId,
            0, // operatorTableReferenceTimestamp
            fuzzRefundCollector,
            fuzzAvsFee,
            block.timestamp + taskSLA,
            fuzzPayload
        );

        // Create task
        vm.prank(creator);
        bytes32 taskHash = taskMailbox.createTask(taskParams);

        // Verify task hash
        assertEq(taskHash, expectedTaskHash);

        // Verify global task count incremented by creating another task and checking its hash
        bytes32 nextExpectedTaskHash = keccak256(abi.encode(expectedTaskCount + 1, address(taskMailbox), block.chainid, taskParams));

        // Expect Transfer event for fee transfer from creator to taskMailbox for second task
        if (fuzzAvsFee > 0) {
            vm.expectEmit(true, true, false, true, address(mockToken));
            emit IERC20.Transfer(creator, address(taskMailbox), fuzzAvsFee);
        }

        vm.prank(creator);
        bytes32 nextTaskHash = taskMailbox.createTask(taskParams);
        assertEq(nextTaskHash, nextExpectedTaskHash);

        // Verify task was created
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        assertEq(task.creator, creator);
        assertEq(task.creationTime, block.timestamp);
        assertEq(uint8(task.status), uint8(TaskStatus.CREATED));
        assertEq(task.avs, avs);
        assertEq(task.executorOperatorSetId, executorOperatorSetId);
        assertEq(task.operatorTableReferenceTimestamp, 0);
        assertEq(task.refundCollector, fuzzRefundCollector);
        assertEq(task.avsFee, fuzzAvsFee);
        assertEq(task.feeSplit, 0);
        assertEq(task.payload, fuzzPayload);

        // Verify token transfer if fee > 0
        // Note: We created two tasks with the same fee, so balance should be 2 * fuzzAvsFee
        if (fuzzAvsFee > 0) assertEq(mockToken.balanceOf(address(taskMailbox)), fuzzAvsFee * 2);
    }

    function test_createTask_ZeroFee() public {
        // Set the mock hook to return 0 fee
        mockTaskHook.setDefaultFee(0);

        TaskParams memory taskParams = _createValidTaskParams();

        uint balanceBefore = mockToken.balanceOf(address(taskMailbox));

        vm.prank(creator);
        bytes32 taskHash = taskMailbox.createTask(taskParams);

        // Verify no token transfer occurred
        assertEq(mockToken.balanceOf(address(taskMailbox)), balanceBefore);

        // Verify task was created with zero fee
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        assertEq(task.avsFee, 0);
    }

    function test_createTask_NoFeeToken() public {
        // Set up config without fee token
        OperatorSet memory operatorSet = OperatorSet(avs2, executorOperatorSetId2);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.feeToken = IERC20(address(0));

        vm.prank(avs2);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        TaskParams memory taskParams =
            TaskParams({refundCollector: refundCollector, executorOperatorSet: operatorSet, payload: bytes("test payload")});

        uint balanceBefore = mockToken.balanceOf(address(taskMailbox));

        vm.prank(creator);
        bytes32 taskHash = taskMailbox.createTask(taskParams);

        // Verify no token transfer occurred even with non-zero fee
        assertEq(mockToken.balanceOf(address(taskMailbox)), balanceBefore);

        // Verify task was created
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        assertEq(task.avsFee, 1 ether);
    }

    function test_Revert_WhenPayloadIsEmpty() public {
        TaskParams memory taskParams = _createValidTaskParams();
        taskParams.payload = bytes("");

        vm.prank(creator);
        vm.expectRevert(PayloadIsEmpty.selector);
        taskMailbox.createTask(taskParams);
    }

    function test_Revert_WhenExecutorOperatorSetNotRegistered() public {
        TaskParams memory taskParams = _createValidTaskParams();
        taskParams.executorOperatorSet.id = 99; // Unregistered operator set

        vm.prank(creator);
        vm.expectRevert(ExecutorOperatorSetNotRegistered.selector);
        taskMailbox.createTask(taskParams);
    }

    function test_Revert_WhenExecutorOperatorSetTaskConfigNotSet() public {
        // Create an operator set that has never been configured
        OperatorSet memory operatorSet = OperatorSet(avs2, executorOperatorSetId2);

        TaskParams memory taskParams =
            TaskParams({refundCollector: refundCollector, executorOperatorSet: operatorSet, payload: bytes("test payload")});

        // Should revert because operator set is not registered (no config set)
        vm.prank(creator);
        vm.expectRevert(ExecutorOperatorSetNotRegistered.selector);
        taskMailbox.createTask(taskParams);
    }

    function test_Revert_ReentrancyOnCreateTask() public {
        // Deploy reentrant attacker as task hook
        AVSTaskHookReentrantAttacker attacker = new AVSTaskHookReentrantAttacker(address(taskMailbox));

        // Set up executor operator set with attacker as hook
        OperatorSet memory operatorSet = OperatorSet(avs2, executorOperatorSetId2);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.taskHook = IAVSTaskHook(address(attacker));

        vm.prank(avs2);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Give attacker tokens and approve
        mockToken.mint(address(attacker), 1000 ether);
        vm.prank(address(attacker));
        mockToken.approve(address(taskMailbox), type(uint).max);

        // Set up attack parameters
        TaskParams memory taskParams =
            TaskParams({refundCollector: refundCollector, executorOperatorSet: operatorSet, payload: bytes("test payload")});

        attacker.setAttackParams(
            taskParams,
            bytes32(0),
            _createValidBN254Certificate(bytes32(0), uint96(block.timestamp)),
            bytes(""),
            true, // attack on post
            true // attack createTask
        );

        // Try to create task - should revert on reentrancy
        vm.prank(creator);
        vm.expectRevert("ReentrancyGuard: reentrant call");
        taskMailbox.createTask(taskParams);
    }

    function test_Revert_createTask_InvalidFeeReceiver_RefundCollector() public {
        // Set up operator set with fee token
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.feeToken = mockToken;
        config.feeCollector = feeCollector;

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create task params with zero refund collector
        TaskParams memory taskParams =
            TaskParams({refundCollector: address(0), executorOperatorSet: operatorSet, payload: bytes("test payload")});

        // Should revert with InvalidFeeReceiver when refundCollector is zero
        vm.prank(creator);
        vm.expectRevert(InvalidFeeReceiver.selector);
        taskMailbox.createTask(taskParams);
    }

    function test_Revert_createTask_InvalidFeeReceiver_FeeCollector() public {
        // Set up operator set with fee token but zero fee collector
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.feeToken = mockToken;
        config.feeCollector = address(0); // Zero fee collector

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create valid task params
        TaskParams memory taskParams = _createValidTaskParams();

        // Should revert with InvalidFeeReceiver when feeCollector is zero
        vm.prank(creator);
        vm.expectRevert(InvalidFeeReceiver.selector);
        taskMailbox.createTask(taskParams);
    }

    function test_createTask_ValidWithZeroFeeReceivers_NoFeeToken() public {
        // When there's no fee token, zero addresses should be allowed
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.feeToken = IERC20(address(0)); // No fee token
        config.feeCollector = address(0); // Zero fee collector is OK when no fee token

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create task params with zero refund collector
        TaskParams memory taskParams = TaskParams({
            refundCollector: address(0), // Zero refund collector is OK when no fee token
            executorOperatorSet: operatorSet,
            payload: bytes("test payload")
        });

        // Should succeed when there's no fee token
        vm.prank(creator);
        bytes32 taskHash = taskMailbox.createTask(taskParams);

        // Verify task was created
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        assertEq(task.creator, creator);
        assertEq(task.refundCollector, address(0));
    }

    function test_createTask_CapturesFeeSplitValues() public {
        // Set fee split values
        uint16 feeSplit = 1500; // 15%
        address localFeeSplitCollector = address(0x456);
        vm.prank(owner);
        taskMailbox.setFeeSplit(feeSplit);
        vm.prank(owner);
        taskMailbox.setFeeSplitCollector(localFeeSplitCollector);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();

        // Expect Transfer event for fee transfer from creator to taskMailbox
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(creator, address(taskMailbox), avsFee);

        vm.prank(creator);
        bytes32 taskHash = taskMailbox.createTask(taskParams);

        // Verify task captured current fee split value
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        assertEq(task.feeSplit, feeSplit);

        // Change fee split values
        uint16 newFeeSplit = 3000; // 30%
        address newFeeSplitCollector = address(0x789);
        vm.prank(owner);
        taskMailbox.setFeeSplit(newFeeSplit);
        vm.prank(owner);
        taskMailbox.setFeeSplitCollector(newFeeSplitCollector);

        // Create another task

        // Expect Transfer event for fee transfer from creator to taskMailbox
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(creator, address(taskMailbox), avsFee);

        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Verify new task has new fee split while old task retains old value
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        assertEq(newTask.feeSplit, newFeeSplit);

        // Verify old task still has old fee split value
        task = taskMailbox.getTaskInfo(taskHash);
        assertEq(task.feeSplit, feeSplit);

        // Verify that the global feeSplitCollector is used (not captured in task)
        assertEq(taskMailbox.feeSplitCollector(), newFeeSplitCollector);
    }
}

// Test contract for submitResult
contract TaskMailboxUnitTests_submitResult is TaskMailboxUnitTests {
    bytes32 public taskHash;

    function setUp() public override {
        super.setUp();

        // Set up executor operator set task config
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create a task
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        taskHash = taskMailbox.createTask(taskParams);
    }

    function testFuzz_submitResult_WithBN254Certificate(bytes memory fuzzResult) public {
        // Advance time by 1 second to pass TimestampAtCreation check
        vm.warp(block.timestamp + 1);

        IBN254CertificateVerifier.BN254Certificate memory cert =
            _createValidBN254CertificateForResult(fuzzResult, _getTaskReferenceTimestamp(taskHash));

        // Expect event
        vm.expectEmit(true, true, true, true, address(taskMailbox));
        emit TaskVerified(aggregator, taskHash, avs, executorOperatorSetId, abi.encode(cert), fuzzResult);

        // Submit result
        vm.prank(aggregator);
        taskMailbox.submitResult(taskHash, abi.encode(cert), fuzzResult);

        // Verify task was verified
        TaskStatus status = taskMailbox.getTaskStatus(taskHash);
        assertEq(uint8(status), uint8(TaskStatus.VERIFIED));

        // Verify result was stored
        bytes memory storedResult = taskMailbox.getTaskResult(taskHash);
        assertEq(storedResult, fuzzResult);

        // Verify certificate was stored
        Task memory verifiedTask = taskMailbox.getTaskInfo(taskHash);
        assertEq(verifiedTask.executorCert, abi.encode(cert));
    }

    function testFuzz_submitResult_WithECDSACertificate(bytes memory fuzzResult) public {
        // Setup executor operator set with ECDSA curve type
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.curveType = IKeyRegistrarTypes.CurveType.ECDSA;

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Advance time by 1 second to pass TimestampAtCreation check
        vm.warp(block.timestamp + 1);

        IECDSACertificateVerifier.ECDSACertificate memory cert =
            _createValidECDSACertificateForResult(fuzzResult, _getTaskReferenceTimestamp(newTaskHash));

        // Expect event
        vm.expectEmit(true, true, true, true, address(taskMailbox));
        emit TaskVerified(aggregator, newTaskHash, avs, executorOperatorSetId, abi.encode(cert), fuzzResult);

        // Submit result
        vm.prank(aggregator);
        taskMailbox.submitResult(newTaskHash, abi.encode(cert), fuzzResult);

        // Verify task was verified
        TaskStatus status = taskMailbox.getTaskStatus(newTaskHash);
        assertEq(uint8(status), uint8(TaskStatus.VERIFIED));

        // Verify result was stored
        bytes memory storedResult = taskMailbox.getTaskResult(newTaskHash);
        assertEq(storedResult, fuzzResult);

        // Verify certificate was stored
        Task memory verifiedTask = taskMailbox.getTaskInfo(newTaskHash);
        assertEq(verifiedTask.executorCert, abi.encode(cert));
    }

    function test_Revert_WhenTimestampAtCreation() public {
        // Get task creation time
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        uint96 taskCreationTime = task.creationTime;

        IBN254CertificateVerifier.BN254Certificate memory cert = _createValidBN254Certificate(taskHash, taskCreationTime);

        // Don't advance time
        vm.prank(aggregator);
        vm.expectRevert(TimestampAtCreation.selector);
        taskMailbox.submitResult(taskHash, abi.encode(cert), bytes("result"));
    }

    function test_Revert_WhenTaskExpired() public {
        // Get task creation time before advancing time
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        uint96 taskCreationTime = task.creationTime;

        // Advance time past task SLA
        vm.warp(block.timestamp + taskSLA + 1);

        IBN254CertificateVerifier.BN254Certificate memory cert = _createValidBN254Certificate(taskHash, taskCreationTime);

        vm.prank(aggregator);
        vm.expectRevert(abi.encodeWithSelector(InvalidTaskStatus.selector, TaskStatus.CREATED, TaskStatus.EXPIRED));
        taskMailbox.submitResult(taskHash, abi.encode(cert), bytes("result"));
    }

    function test_Revert_WhenTaskDoesNotExist() public {
        bytes32 nonExistentHash = keccak256("non-existent");

        // Use block.timestamp as a dummy creation time for non-existent task
        uint96 dummyCreationTime = uint96(block.timestamp);

        // Advance time by 1 second to pass TimestampAtCreation check
        vm.warp(block.timestamp + 1);

        IBN254CertificateVerifier.BN254Certificate memory cert = _createValidBN254Certificate(nonExistentHash, dummyCreationTime);

        vm.prank(aggregator);
        vm.expectRevert(abi.encodeWithSelector(InvalidTaskStatus.selector, TaskStatus.CREATED, TaskStatus.NONE));
        taskMailbox.submitResult(nonExistentHash, abi.encode(cert), bytes("result"));
    }

    function test_Revert_WhenCertificateVerificationFailed_BN254() public {
        // Create a custom mock that returns false for certificate verification
        MockBN254CertificateVerifierFailure mockFailingVerifier = new MockBN254CertificateVerifierFailure();

        // Deploy a new TaskMailbox with the failing verifier using proxy pattern
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        TaskMailbox taskMailboxImpl = new TaskMailbox(address(mockFailingVerifier), address(mockECDSACertificateVerifier), "1.0.0");
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(taskMailboxImpl),
            address(proxyAdmin),
            abi.encodeWithSelector(TaskMailbox.initialize.selector, owner, 0, feeSplitCollector)
        );
        TaskMailbox failingTaskMailbox = TaskMailbox(address(proxy));

        // Give creator tokens and approve the new TaskMailbox
        vm.prank(creator);
        mockToken.approve(address(failingTaskMailbox), type(uint).max);

        // Set config
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();

        vm.prank(avs);
        failingTaskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create new task with this config
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        bytes32 newTaskHash = failingTaskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = failingTaskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Advance time
        vm.warp(block.timestamp + 1);

        bytes memory result = bytes("result");
        IBN254CertificateVerifier.BN254Certificate memory cert =
            _createValidBN254CertificateForResult(result, _getTaskReferenceTimestamp(taskHash));

        vm.prank(aggregator);
        vm.expectRevert(CertificateVerificationFailed.selector);
        failingTaskMailbox.submitResult(newTaskHash, abi.encode(cert), result);
    }

    function test_Revert_WhenCertificateVerificationFailed_ECDSA() public {
        // Create a custom mock that returns false for certificate verification
        MockECDSACertificateVerifierFailure mockECDSACertificateVerifierFailure = new MockECDSACertificateVerifierFailure();

        // Deploy a new TaskMailbox with the failing ECDSA verifier using proxy pattern
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        TaskMailbox taskMailboxImpl =
            new TaskMailbox(address(mockBN254CertificateVerifier), address(mockECDSACertificateVerifierFailure), "1.0.0");
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(taskMailboxImpl),
            address(proxyAdmin),
            abi.encodeWithSelector(TaskMailbox.initialize.selector, owner, 0, feeSplitCollector)
        );
        TaskMailbox failingTaskMailbox = TaskMailbox(address(proxy));

        // Give creator tokens and approve the new TaskMailbox
        vm.prank(creator);
        mockToken.approve(address(failingTaskMailbox), type(uint).max);

        // Setup executor operator set with ECDSA curve type
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.curveType = IKeyRegistrarTypes.CurveType.ECDSA;

        vm.prank(avs);
        failingTaskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        bytes32 newTaskHash = failingTaskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = failingTaskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Advance time
        vm.warp(block.timestamp + 1);

        // Create ECDSA certificate
        bytes memory result = bytes("result");
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert =
            _createValidECDSACertificateForResult(result, _getTaskReferenceTimestamp(newTaskHash));

        // Submit should fail
        vm.prank(aggregator);
        vm.expectRevert(CertificateVerificationFailed.selector);
        failingTaskMailbox.submitResult(newTaskHash, abi.encode(cert), result);
    }

    function test_Revert_WhenInvalidCertificateEncoding() public {
        // Setup executor operator set with ECDSA curve type
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.curveType = IKeyRegistrarTypes.CurveType.ECDSA;

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Create invalid encoded certificate (not properly encoded ECDSA certificate)
        bytes memory invalidCert = abi.encode("invalid", "certificate", "data");
        bytes memory result = bytes("test result");

        // Submit should fail due to decoding error
        vm.prank(aggregator);
        vm.expectRevert(); // Will revert during abi.decode
        taskMailbox.submitResult(newTaskHash, invalidCert, result);
    }

    function test_Revert_AlreadyVerified() public {
        // Get task creation time
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        uint96 taskCreationTime = task.creationTime;

        // First submit a valid result
        vm.warp(block.timestamp + 1);
        bytes memory result = bytes("result");
        IBN254CertificateVerifier.BN254Certificate memory cert =
            _createValidBN254CertificateForResult(result, _getTaskReferenceTimestamp(taskHash));

        vm.prank(aggregator);
        taskMailbox.submitResult(taskHash, abi.encode(cert), result);

        // Try to submit again with a different result
        bytes memory newResult = bytes("new result");
        IBN254CertificateVerifier.BN254Certificate memory cert2 =
            _createValidBN254CertificateForResult(newResult, _getTaskReferenceTimestamp(taskHash));

        vm.prank(aggregator);
        vm.expectRevert(abi.encodeWithSelector(InvalidTaskStatus.selector, TaskStatus.CREATED, TaskStatus.VERIFIED));
        taskMailbox.submitResult(taskHash, abi.encode(cert2), newResult);
    }

    function test_Revert_ReentrancyOnSubmitResult() public {
        // Deploy reentrant attacker as task hook
        AVSTaskHookReentrantAttacker attacker = new AVSTaskHookReentrantAttacker(address(taskMailbox));

        // Set up executor operator set with attacker as hook
        OperatorSet memory operatorSet = OperatorSet(avs2, executorOperatorSetId2);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.taskHook = IAVSTaskHook(address(attacker));

        vm.prank(avs2);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create a task
        TaskParams memory taskParams =
            TaskParams({refundCollector: refundCollector, executorOperatorSet: operatorSet, payload: bytes("test payload")});

        vm.prank(creator);
        bytes32 attackTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory attackTask = taskMailbox.getTaskInfo(attackTaskHash);
        uint96 attackTaskCreationTime = attackTask.creationTime;

        // Advance time
        vm.warp(block.timestamp + 1);

        // Set up attack parameters
        bytes memory result = bytes("result");
        IBN254CertificateVerifier.BN254Certificate memory cert =
            _createValidBN254CertificateForResult(result, _getTaskReferenceTimestamp(taskHash));

        attacker.setAttackParams(
            taskParams,
            attackTaskHash,
            cert,
            result,
            false, // attack on handleTaskResultSubmission
            false // attack submitResult
        );

        // Try to submit result - should revert on reentrancy
        vm.prank(aggregator);
        vm.expectRevert("ReentrancyGuard: reentrant call");
        taskMailbox.submitResult(attackTaskHash, abi.encode(cert), result);
    }

    function test_Revert_ReentrancyOnSubmitResult_TryingToCreateTask() public {
        // Deploy reentrant attacker as task hook
        AVSTaskHookReentrantAttacker attacker = new AVSTaskHookReentrantAttacker(address(taskMailbox));

        // Give attacker tokens and approve
        mockToken.mint(address(attacker), 1000 ether);
        vm.prank(address(attacker));
        mockToken.approve(address(taskMailbox), type(uint).max);

        // Set up executor operator set with attacker as hook
        OperatorSet memory operatorSet = OperatorSet(avs2, executorOperatorSetId2);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.taskHook = IAVSTaskHook(address(attacker));

        vm.prank(avs2);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create a task
        TaskParams memory taskParams =
            TaskParams({refundCollector: refundCollector, executorOperatorSet: operatorSet, payload: bytes("test payload")});

        vm.prank(creator);
        bytes32 attackTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory attackTask = taskMailbox.getTaskInfo(attackTaskHash);
        uint96 attackTaskCreationTime = attackTask.creationTime;

        // Advance time
        vm.warp(block.timestamp + 1);

        // Set up attack parameters
        bytes memory result = bytes("result");
        IBN254CertificateVerifier.BN254Certificate memory cert =
            _createValidBN254CertificateForResult(result, _getTaskReferenceTimestamp(taskHash));

        attacker.setAttackParams(
            taskParams,
            attackTaskHash,
            cert,
            result,
            false, // attack on handleTaskResultSubmission
            true // attack createTask
        );

        // Try to submit result - should revert on reentrancy
        vm.prank(aggregator);
        vm.expectRevert("ReentrancyGuard: reentrant call");
        taskMailbox.submitResult(attackTaskHash, abi.encode(cert), result);
    }

    function test_submitResult_WithZeroStakeThreshold() public {
        // Setup executor operator set with zero stake threshold
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.consensus = Consensus({
            consensusType: ConsensusType.STAKE_PROPORTION_THRESHOLD,
            value: abi.encode(uint16(0)) // Zero threshold
        });

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Advance time
        vm.warp(block.timestamp + 1);

        // Submit result with zero threshold should still work
        bytes memory result = bytes("test result");
        IBN254CertificateVerifier.BN254Certificate memory cert =
            _createValidBN254CertificateForResult(result, _getTaskReferenceTimestamp(taskHash));

        vm.prank(aggregator);
        taskMailbox.submitResult(newTaskHash, abi.encode(cert), result);

        // Verify task was verified
        TaskStatus status = taskMailbox.getTaskStatus(newTaskHash);
        assertEq(uint8(status), uint8(TaskStatus.VERIFIED));
    }

    function test_submitResult_WithMaxStakeThreshold() public {
        // Setup executor operator set with max stake threshold
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.consensus = Consensus({
            consensusType: ConsensusType.STAKE_PROPORTION_THRESHOLD,
            value: abi.encode(uint16(10_000)) // 100% threshold
        });

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Advance time
        vm.warp(block.timestamp + 1);

        // Submit result with max threshold
        bytes memory result = bytes("test result");
        IBN254CertificateVerifier.BN254Certificate memory cert =
            _createValidBN254CertificateForResult(result, _getTaskReferenceTimestamp(taskHash));

        vm.prank(aggregator);
        taskMailbox.submitResult(newTaskHash, abi.encode(cert), result);

        // Verify task was verified
        TaskStatus status = taskMailbox.getTaskStatus(newTaskHash);
        assertEq(uint8(status), uint8(TaskStatus.VERIFIED));
    }

    function test_Revert_WhenBN254CertificateHasEmptySignature() public {
        // Advance time by 1 second to pass TimestampAtCreation check
        vm.warp(block.timestamp + 1);

        // Create BN254 certificate with empty signature (X=0, Y=0)
        bytes memory result = bytes("result");
        IBN254CertificateVerifierTypes.BN254Certificate memory cert = IBN254CertificateVerifierTypes.BN254Certificate({
            referenceTimestamp: _getTaskReferenceTimestamp(taskHash),
            messageHash: keccak256(result),
            signature: BN254.G1Point(0, 0), // Empty signature
            apk: BN254.G2Point([uint(1), uint(2)], [uint(3), uint(4)]),
            nonSignerWitnesses: new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](0)
        });

        // Submit result should fail with EmptyCertificateSignature error
        vm.prank(aggregator);
        vm.expectRevert(EmptyCertificateSignature.selector);
        taskMailbox.submitResult(taskHash, abi.encode(cert), result);
    }

    function test_Revert_WhenBN254CertificateHasEmptySignature_OnlyXZero() public {
        // Get task creation time
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        uint96 taskCreationTime = task.creationTime;

        // Advance time by 1 second to pass TimestampAtCreation check
        vm.warp(block.timestamp + 1);

        // Create BN254 certificate with empty X coordinate
        bytes memory result = bytes("result");
        IBN254CertificateVerifierTypes.BN254Certificate memory cert = IBN254CertificateVerifierTypes.BN254Certificate({
            referenceTimestamp: _getTaskReferenceTimestamp(taskHash),
            messageHash: keccak256(result),
            signature: BN254.G1Point(0, 2), // Only X is zero
            apk: BN254.G2Point([uint(1), uint(2)], [uint(3), uint(4)]),
            nonSignerWitnesses: new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](0)
        });

        // Submit result should fail with EmptyCertificateSignature error
        vm.prank(aggregator);
        vm.expectRevert(EmptyCertificateSignature.selector);
        taskMailbox.submitResult(taskHash, abi.encode(cert), result);
    }

    function test_Revert_WhenBN254CertificateHasEmptySignature_OnlyYZero() public {
        // Get task creation time
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        uint96 taskCreationTime = task.creationTime;

        // Advance time by 1 second to pass TimestampAtCreation check
        vm.warp(block.timestamp + 1);

        // Create BN254 certificate with empty Y coordinate
        bytes memory result = bytes("result");
        IBN254CertificateVerifierTypes.BN254Certificate memory cert = IBN254CertificateVerifierTypes.BN254Certificate({
            referenceTimestamp: _getTaskReferenceTimestamp(taskHash),
            messageHash: keccak256(result),
            signature: BN254.G1Point(1, 0), // Only Y is zero
            apk: BN254.G2Point([uint(1), uint(2)], [uint(3), uint(4)]),
            nonSignerWitnesses: new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](0)
        });

        // Submit result should fail with EmptyCertificateSignature error
        vm.prank(aggregator);
        vm.expectRevert(EmptyCertificateSignature.selector);
        taskMailbox.submitResult(taskHash, abi.encode(cert), result);
    }

    function test_Revert_WhenECDSACertificateHasEmptySignature() public {
        // Setup executor operator set with ECDSA curve type
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.curveType = IKeyRegistrarTypes.CurveType.ECDSA;

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Advance time by 1 second to pass TimestampAtCreation check
        vm.warp(block.timestamp + 1);

        // Create ECDSA certificate with empty signature
        bytes memory result = bytes("result");
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert = IECDSACertificateVerifierTypes.ECDSACertificate({
            referenceTimestamp: _getTaskReferenceTimestamp(newTaskHash),
            messageHash: _getECDSACertificateDigest(_getTaskReferenceTimestamp(newTaskHash), result),
            sig: bytes("") // Empty signature
        });

        // Submit result should fail with EmptyCertificateSignature error
        vm.prank(aggregator);
        vm.expectRevert(EmptyCertificateSignature.selector);
        taskMailbox.submitResult(newTaskHash, abi.encode(cert), result);
    }

    function test_Revert_WhenBN254CertificateHasEmptySignature_WithZeroThreshold() public {
        // Setup executor operator set with zero stake threshold
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.consensus = Consensus({
            consensusType: ConsensusType.STAKE_PROPORTION_THRESHOLD,
            value: abi.encode(uint16(0)) // Zero threshold
        });

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Advance time
        vm.warp(block.timestamp + 1);

        // Create BN254 certificate with empty signature
        bytes memory result = bytes("result");
        IBN254CertificateVerifierTypes.BN254Certificate memory cert = IBN254CertificateVerifierTypes.BN254Certificate({
            referenceTimestamp: _getTaskReferenceTimestamp(taskHash),
            messageHash: keccak256(result),
            signature: BN254.G1Point(0, 0), // Empty signature
            apk: BN254.G2Point([uint(1), uint(2)], [uint(3), uint(4)]),
            nonSignerWitnesses: new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](0)
        });

        // Even with zero threshold, empty signatures should be rejected
        vm.prank(aggregator);
        vm.expectRevert(EmptyCertificateSignature.selector);
        taskMailbox.submitResult(newTaskHash, abi.encode(cert), result);
    }

    function test_Revert_WhenECDSACertificateHasEmptySignature_WithZeroThreshold() public {
        // Setup executor operator set with ECDSA curve type and zero threshold
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.curveType = IKeyRegistrarTypes.CurveType.ECDSA;
        config.consensus = Consensus({
            consensusType: ConsensusType.STAKE_PROPORTION_THRESHOLD,
            value: abi.encode(uint16(0)) // Zero threshold
        });

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Advance time
        vm.warp(block.timestamp + 1);

        // Create ECDSA certificate with empty signature
        bytes memory result = bytes("result");
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert = IECDSACertificateVerifierTypes.ECDSACertificate({
            referenceTimestamp: _getTaskReferenceTimestamp(newTaskHash),
            messageHash: _getECDSACertificateDigest(_getTaskReferenceTimestamp(newTaskHash), result),
            sig: bytes("") // Empty signature
        });

        // Even with zero threshold, empty signatures should be rejected
        vm.prank(aggregator);
        vm.expectRevert(EmptyCertificateSignature.selector);
        taskMailbox.submitResult(newTaskHash, abi.encode(cert), result);
    }

    function test_submitResult_FeeTransferToCollector() public {
        // Set up operator set with fee token
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.feeToken = mockToken;
        config.feeCollector = feeCollector;

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create task with fee
        TaskParams memory taskParams = _createValidTaskParams();

        // Expect Transfer event for fee transfer from creator to taskMailbox
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(creator, address(taskMailbox), avsFee);

        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Check initial balances
        uint mailboxBalanceBefore = mockToken.balanceOf(address(taskMailbox));
        uint feeCollectorBalanceBefore = mockToken.balanceOf(feeCollector);
        // mailboxBalanceBefore should be 2*avsFee (one from setUp, one from this test)

        // Advance time and submit result
        vm.warp(block.timestamp + 1);

        bytes memory result = bytes("result");
        IBN254CertificateVerifierTypes.BN254Certificate memory cert = IBN254CertificateVerifierTypes.BN254Certificate({
            referenceTimestamp: _getTaskReferenceTimestamp(newTaskHash),
            messageHash: keccak256(result),
            signature: BN254.G1Point(1, 2),
            apk: BN254.G2Point([uint(3), uint(4)], [uint(5), uint(6)]),
            nonSignerWitnesses: new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](0)
        });

        // Mock certificate verification
        vm.mockCall(
            address(mockBN254CertificateVerifier),
            abi.encodeWithSelector(IBN254CertificateVerifier.verifyCertificateProportion.selector),
            abi.encode(true)
        );

        // Expect Transfer event for fee transfer from taskMailbox to feeCollector
        // Since feeSplit is 0 by default, all avsFee goes to feeCollector
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(address(taskMailbox), feeCollector, avsFee);

        vm.prank(aggregator);
        taskMailbox.submitResult(newTaskHash, abi.encode(cert), result);

        // Verify fee was transferred to fee collector
        assertEq(mockToken.balanceOf(address(taskMailbox)), mailboxBalanceBefore - avsFee);
        assertEq(mockToken.balanceOf(feeCollector), feeCollectorBalanceBefore + avsFee);

        // Verify task cannot be refunded after verification
        Task memory task = taskMailbox.getTaskInfo(newTaskHash);
        assertEq(uint8(task.status), uint8(TaskStatus.VERIFIED));
        assertFalse(task.isFeeRefunded);
    }

    function test_FeeSplit_10Percent() public {
        // Setup fee split
        uint16 feeSplit = 1000; // 10%
        address localFeeSplitCollector = address(0x789);
        vm.prank(owner);
        taskMailbox.setFeeSplit(feeSplit);
        vm.prank(owner);
        taskMailbox.setFeeSplitCollector(localFeeSplitCollector);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();

        // Expect Transfer event for fee transfer from creator to taskMailbox
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(creator, address(taskMailbox), avsFee);

        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Check initial balances
        uint feeCollectorBalanceBefore = mockToken.balanceOf(feeCollector);
        uint feeSplitCollectorBalanceBefore = mockToken.balanceOf(localFeeSplitCollector);

        // Submit result
        vm.warp(block.timestamp + 1);
        bytes memory result = bytes("result");
        bytes memory executorCert = abi.encode(_createValidBN254CertificateForResult(result, _getTaskReferenceTimestamp(taskHash)));

        // Calculate expected amounts
        uint expectedFeeSplitAmount = (avsFee * feeSplit) / 10_000;
        uint expectedFeeCollectorAmount = avsFee - expectedFeeSplitAmount;

        // Expect Transfer events for fee distribution
        // First, transfer to fee split collector
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(address(taskMailbox), localFeeSplitCollector, expectedFeeSplitAmount);

        // Then, transfer remaining to fee collector
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(address(taskMailbox), feeCollector, expectedFeeCollectorAmount);

        vm.prank(aggregator);
        taskMailbox.submitResult(newTaskHash, executorCert, result);

        assertEq(mockToken.balanceOf(localFeeSplitCollector), feeSplitCollectorBalanceBefore + expectedFeeSplitAmount);
        assertEq(mockToken.balanceOf(feeCollector), feeCollectorBalanceBefore + expectedFeeCollectorAmount);
    }

    function test_FeeSplit_50Percent() public {
        // Setup fee split
        uint16 feeSplit = 5000; // 50%
        address localFeeSplitCollector = address(0x789);
        vm.prank(owner);
        taskMailbox.setFeeSplit(feeSplit);
        vm.prank(owner);
        taskMailbox.setFeeSplitCollector(localFeeSplitCollector);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();

        // Expect Transfer event for fee transfer from creator to taskMailbox
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(creator, address(taskMailbox), avsFee);

        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Check initial balances
        uint feeCollectorBalanceBefore = mockToken.balanceOf(feeCollector);
        uint feeSplitCollectorBalanceBefore = mockToken.balanceOf(localFeeSplitCollector);

        // Submit result
        vm.warp(block.timestamp + 1);
        bytes memory result = bytes("result");
        bytes memory executorCert = abi.encode(_createValidBN254CertificateForResult(result, _getTaskReferenceTimestamp(taskHash)));

        // Calculate expected amounts - should be equal split
        uint expectedFeeSplitAmount = avsFee / 2;
        uint expectedFeeCollectorAmount = avsFee - expectedFeeSplitAmount;

        // Expect Transfer events for fee distribution
        // First, transfer to fee split collector
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(address(taskMailbox), localFeeSplitCollector, expectedFeeSplitAmount);

        // Then, transfer remaining to fee collector
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(address(taskMailbox), feeCollector, expectedFeeCollectorAmount);

        vm.prank(aggregator);
        taskMailbox.submitResult(newTaskHash, executorCert, result);

        assertEq(mockToken.balanceOf(localFeeSplitCollector), feeSplitCollectorBalanceBefore + expectedFeeSplitAmount);
        assertEq(mockToken.balanceOf(feeCollector), feeCollectorBalanceBefore + expectedFeeCollectorAmount);
    }

    function test_FeeSplit_0Percent() public {
        // Setup fee split - 0% means all fees go to fee collector
        uint16 feeSplit = 0;
        address localFeeSplitCollector = address(0x789);
        vm.prank(owner);
        taskMailbox.setFeeSplit(feeSplit);
        vm.prank(owner);
        taskMailbox.setFeeSplitCollector(localFeeSplitCollector);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();

        // Expect Transfer event for fee transfer from creator to taskMailbox
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(creator, address(taskMailbox), avsFee);

        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Check initial balances
        uint feeCollectorBalanceBefore = mockToken.balanceOf(feeCollector);
        uint feeSplitCollectorBalanceBefore = mockToken.balanceOf(localFeeSplitCollector);

        // Submit result
        vm.warp(block.timestamp + 1);
        bytes memory result = bytes("result");
        bytes memory executorCert = abi.encode(_createValidBN254CertificateForResult(result, _getTaskReferenceTimestamp(taskHash)));

        // Expect Transfer event for fee transfer to fee collector only (no fee split)
        // Since feeSplit is 0, all avsFee goes to feeCollector
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(address(taskMailbox), feeCollector, avsFee);

        vm.prank(aggregator);
        taskMailbox.submitResult(newTaskHash, executorCert, result);

        // Verify all fees went to fee collector
        assertEq(mockToken.balanceOf(localFeeSplitCollector), feeSplitCollectorBalanceBefore); // No change
        assertEq(mockToken.balanceOf(feeCollector), feeCollectorBalanceBefore + avsFee);
    }

    function test_FeeSplit_100Percent() public {
        // Setup fee split - 100% means all fees go to fee split collector
        uint16 feeSplit = 10_000;
        address localFeeSplitCollector = address(0x789);
        vm.prank(owner);
        taskMailbox.setFeeSplit(feeSplit);
        vm.prank(owner);
        taskMailbox.setFeeSplitCollector(localFeeSplitCollector);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();

        // Expect Transfer event for fee transfer from creator to taskMailbox
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(creator, address(taskMailbox), avsFee);

        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Check initial balances
        uint feeCollectorBalanceBefore = mockToken.balanceOf(feeCollector);
        uint feeSplitCollectorBalanceBefore = mockToken.balanceOf(localFeeSplitCollector);

        // Submit result
        vm.warp(block.timestamp + 1);
        bytes memory result = bytes("result");
        bytes memory executorCert = abi.encode(_createValidBN254CertificateForResult(result, _getTaskReferenceTimestamp(taskHash)));

        // Expect Transfer event for fee transfer to fee split collector only
        // Since feeSplit is 100%, all avsFee goes to feeSplitCollector
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(address(taskMailbox), localFeeSplitCollector, avsFee);

        vm.prank(aggregator);
        taskMailbox.submitResult(newTaskHash, executorCert, result);

        // Verify all fees went to fee split collector
        assertEq(mockToken.balanceOf(localFeeSplitCollector), feeSplitCollectorBalanceBefore + avsFee);
        assertEq(mockToken.balanceOf(feeCollector), feeCollectorBalanceBefore); // No change
    }

    function test_FeeSplit_ZeroFeeAmount() public {
        // Setup fee split
        uint16 feeSplit = 5000; // 50%
        address localFeeSplitCollector = address(0x789);
        vm.prank(owner);
        taskMailbox.setFeeSplit(feeSplit);
        vm.prank(owner);
        taskMailbox.setFeeSplitCollector(localFeeSplitCollector);

        // Setup operator set with zero fee
        mockTaskHook.setDefaultFee(0);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Check initial balances
        uint feeCollectorBalanceBefore = mockToken.balanceOf(feeCollector);
        uint feeSplitCollectorBalanceBefore = mockToken.balanceOf(localFeeSplitCollector);

        // Submit result
        vm.warp(block.timestamp + 1);
        bytes memory result = bytes("result");
        bytes memory executorCert = abi.encode(_createValidBN254CertificateForResult(result, _getTaskReferenceTimestamp(taskHash)));

        vm.prank(aggregator);
        taskMailbox.submitResult(newTaskHash, executorCert, result);

        // Verify no transfers occurred
        assertEq(mockToken.balanceOf(feeSplitCollector), feeSplitCollectorBalanceBefore);
        assertEq(mockToken.balanceOf(feeCollector), feeCollectorBalanceBefore);
    }

    function test_FeeSplit_WithSmallFee() public {
        // Setup fee split
        uint16 feeSplit = 3333; // 33.33%
        address localFeeSplitCollector = address(0x789);
        vm.prank(owner);
        taskMailbox.setFeeSplit(feeSplit);
        vm.prank(owner);
        taskMailbox.setFeeSplitCollector(localFeeSplitCollector);

        // Setup small fee
        uint96 smallFee = 100; // 100 wei
        mockTaskHook.setDefaultFee(smallFee);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();

        // Expect Transfer event for fee transfer from creator to taskMailbox
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(creator, address(taskMailbox), smallFee);

        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Check initial balances
        uint feeCollectorBalanceBefore = mockToken.balanceOf(feeCollector);
        uint feeSplitCollectorBalanceBefore = mockToken.balanceOf(localFeeSplitCollector);

        // Submit result
        vm.warp(block.timestamp + 1);
        bytes memory result = bytes("result");
        bytes memory executorCert = abi.encode(_createValidBN254CertificateForResult(result, _getTaskReferenceTimestamp(taskHash)));

        // Calculate expected fee distribution
        uint expectedFeeSplitAmount = (smallFee * feeSplit) / 10_000; // 33 wei
        uint expectedFeeCollectorAmount = smallFee - expectedFeeSplitAmount; // 67 wei

        // Expect Transfer events for fee distribution
        // First, transfer to fee split collector
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(address(taskMailbox), localFeeSplitCollector, expectedFeeSplitAmount);

        // Then, transfer remaining to fee collector
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(address(taskMailbox), feeCollector, expectedFeeCollectorAmount);

        vm.prank(aggregator);
        taskMailbox.submitResult(newTaskHash, executorCert, result);

        assertEq(mockToken.balanceOf(localFeeSplitCollector), feeSplitCollectorBalanceBefore + expectedFeeSplitAmount);
        assertEq(mockToken.balanceOf(feeCollector), feeCollectorBalanceBefore + expectedFeeCollectorAmount);
        assertEq(expectedFeeSplitAmount + expectedFeeCollectorAmount, smallFee); // Verify no wei lost
    }

    function test_FeeSplit_Rounding() public {
        // Setup fee split that will cause rounding
        uint16 feeSplit = 1; // 0.01%
        address localFeeSplitCollector = address(0x789);
        vm.prank(owner);
        taskMailbox.setFeeSplit(feeSplit);
        vm.prank(owner);
        taskMailbox.setFeeSplitCollector(localFeeSplitCollector);

        // Setup fee that won't divide evenly
        uint96 oddFee = 10_001; // Will result in 1.0001 wei split
        mockTaskHook.setDefaultFee(oddFee);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();

        // Expect Transfer event for fee transfer from creator to taskMailbox
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(creator, address(taskMailbox), oddFee);

        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Submit result
        vm.warp(block.timestamp + 1);
        bytes memory result = bytes("result");
        bytes memory executorCert = abi.encode(_createValidBN254CertificateForResult(result, _getTaskReferenceTimestamp(taskHash)));

        // Calculate expected fee distribution
        uint expectedFeeSplitAmount = (oddFee * feeSplit) / 10_000; // 1 wei (rounded down from 1.0001)
        uint expectedFeeCollectorAmount = oddFee - expectedFeeSplitAmount; // 10000 wei

        // Expect Transfer events for fee distribution
        // First, transfer to fee split collector
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(address(taskMailbox), localFeeSplitCollector, expectedFeeSplitAmount);

        // Then, transfer remaining to fee collector
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(address(taskMailbox), feeCollector, expectedFeeCollectorAmount);

        vm.prank(aggregator);
        taskMailbox.submitResult(newTaskHash, executorCert, result);

        assertEq(mockToken.balanceOf(localFeeSplitCollector), expectedFeeSplitAmount);
        assertEq(mockToken.balanceOf(feeCollector), expectedFeeCollectorAmount);
        assertEq(expectedFeeSplitAmount + expectedFeeCollectorAmount, oddFee);
    }

    function testFuzz_FeeSplit(uint16 _feeSplit, uint96 _avsFee) public {
        // Bound inputs
        _feeSplit = uint16(bound(_feeSplit, 0, 10_000));
        vm.assume(_avsFee > 0 && _avsFee <= 1000 ether);
        vm.assume(_avsFee <= mockToken.balanceOf(creator));

        // Setup fee split
        address localFeeSplitCollector = address(0x789);
        vm.prank(owner);
        taskMailbox.setFeeSplit(_feeSplit);
        vm.prank(owner);
        taskMailbox.setFeeSplitCollector(localFeeSplitCollector);

        // Setup fee
        mockTaskHook.setDefaultFee(_avsFee);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();

        // Expect Transfer event for fee transfer from creator to taskMailbox
        if (_avsFee > 0) {
            vm.expectEmit(true, true, false, true, address(mockToken));
            emit IERC20.Transfer(creator, address(taskMailbox), _avsFee);
        }

        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Check initial balances
        uint mailboxBalanceBefore = mockToken.balanceOf(address(taskMailbox));
        uint feeCollectorBalanceBefore = mockToken.balanceOf(feeCollector);
        uint feeSplitCollectorBalanceBefore = mockToken.balanceOf(localFeeSplitCollector);

        // Submit result
        vm.warp(block.timestamp + 1);
        bytes memory result = bytes("result");
        bytes memory executorCert = abi.encode(_createValidBN254CertificateForResult(result, _getTaskReferenceTimestamp(taskHash)));

        // Calculate expected amounts
        uint expectedFeeSplitAmount = (uint(_avsFee) * _feeSplit) / 10_000;
        uint expectedAvsAmount = _avsFee - expectedFeeSplitAmount;

        // Expect Transfer events for fee distribution
        if (_avsFee > 0) {
            if (expectedFeeSplitAmount > 0) {
                // First, transfer to fee split collector
                vm.expectEmit(true, true, false, true, address(mockToken));
                emit IERC20.Transfer(address(taskMailbox), localFeeSplitCollector, expectedFeeSplitAmount);
            }

            if (expectedAvsAmount > 0) {
                // Then, transfer remaining to fee collector
                vm.expectEmit(true, true, false, true, address(mockToken));
                emit IERC20.Transfer(address(taskMailbox), feeCollector, expectedAvsAmount);
            }
        }

        vm.prank(aggregator);
        taskMailbox.submitResult(newTaskHash, executorCert, result);

        // Verify balances
        assertEq(mockToken.balanceOf(localFeeSplitCollector), feeSplitCollectorBalanceBefore + expectedFeeSplitAmount);
        assertEq(mockToken.balanceOf(feeCollector), feeCollectorBalanceBefore + expectedAvsAmount);

        // Verify total distribution equals original fee
        assertEq(expectedFeeSplitAmount + expectedAvsAmount, _avsFee);
    }

    function test_FeeSplit_TaskUsesSnapshotFeeSplitAndCurrentCollector() public {
        // Setup initial fee split
        uint16 initialFeeSplit = 2000; // 20%
        address initialFeeSplitCollector = address(0x789);
        vm.prank(owner);
        taskMailbox.setFeeSplit(initialFeeSplit);
        vm.prank(owner);
        taskMailbox.setFeeSplitCollector(initialFeeSplitCollector);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();

        // Expect Transfer event for fee transfer from creator to taskMailbox
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(creator, address(taskMailbox), avsFee);

        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Change fee split after task creation
        address newFeeSplitCollector = address(0xABC);
        vm.prank(owner);
        taskMailbox.setFeeSplit(5000); // 50%
        vm.prank(owner);
        taskMailbox.setFeeSplitCollector(newFeeSplitCollector);

        // Check initial balances
        uint feeCollectorBalanceBefore = mockToken.balanceOf(feeCollector);
        uint initialCollectorBalanceBefore = mockToken.balanceOf(initialFeeSplitCollector);
        uint newCollectorBalanceBefore = mockToken.balanceOf(newFeeSplitCollector);

        // Submit result
        vm.warp(block.timestamp + 1);
        bytes memory executorCert = abi.encode(_createValidBN254CertificateForResult(bytes("result"), _getTaskReferenceTimestamp(taskHash)));

        // Expect Transfer events for fee distribution
        // First, transfer to current fee split collector (not the initial one)
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(address(taskMailbox), newFeeSplitCollector, ((avsFee * initialFeeSplit) / 10_000));

        // Then, transfer remaining to fee collector
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(address(taskMailbox), feeCollector, avsFee - ((avsFee * initialFeeSplit) / 10_000));

        vm.prank(aggregator);
        taskMailbox.submitResult(newTaskHash, executorCert, bytes("result"));

        assertEq(mockToken.balanceOf(initialFeeSplitCollector), initialCollectorBalanceBefore); // No change
        assertEq(mockToken.balanceOf(feeCollector), feeCollectorBalanceBefore + (avsFee - ((avsFee * initialFeeSplit) / 10_000)));
        assertEq(mockToken.balanceOf(newFeeSplitCollector), newCollectorBalanceBefore + ((avsFee * initialFeeSplit) / 10_000)); // Gets the fee split
    }

    function test_Revert_WhenBN254CertificateHasInvalidReferenceTimestamp() public {
        // Get task creation time
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        uint96 taskCreationTime = task.creationTime;

        // Advance time by 1 second to pass TimestampAtCreation check
        vm.warp(block.timestamp + 1);

        // Create BN254 certificate with wrong reference timestamp (current block.timestamp instead of creation time)
        IBN254CertificateVerifierTypes.BN254Certificate memory cert = IBN254CertificateVerifierTypes.BN254Certificate({
            referenceTimestamp: uint32(block.timestamp), // Wrong timestamp
            messageHash: taskHash,
            signature: BN254.G1Point(1, 2),
            apk: BN254.G2Point([uint(1), uint(2)], [uint(3), uint(4)]),
            nonSignerWitnesses: new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](0)
        });

        // Submit result should fail with InvalidReferenceTimestamp error
        vm.prank(aggregator);
        vm.expectRevert(InvalidReferenceTimestamp.selector);
        taskMailbox.submitResult(taskHash, abi.encode(cert), bytes("result"));
    }

    function test_Revert_WhenECDSACertificateHasInvalidReferenceTimestamp() public {
        // Setup executor operator set with ECDSA curve type
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.curveType = IKeyRegistrarTypes.CurveType.ECDSA;

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Advance time by 1 second to pass TimestampAtCreation check
        vm.warp(block.timestamp + 1);

        // Create ECDSA certificate with wrong reference timestamp
        bytes memory result = bytes("result");
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert = IECDSACertificateVerifierTypes.ECDSACertificate({
            referenceTimestamp: uint32(block.timestamp), // Wrong timestamp
            messageHash: _getECDSACertificateDigest(uint32(block.timestamp), result),
            sig: bytes("0x1234567890abcdef")
        });

        // Submit result should fail with InvalidReferenceTimestamp error
        vm.prank(aggregator);
        vm.expectRevert(InvalidReferenceTimestamp.selector);
        taskMailbox.submitResult(newTaskHash, abi.encode(cert), result);
    }

    function test_Revert_WhenBN254CertificateHasInvalidMessageHash() public {
        // Get task creation time
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        uint96 taskCreationTime = task.creationTime;

        // Advance time by 1 second to pass TimestampAtCreation check
        vm.warp(block.timestamp + 1);

        // Create BN254 certificate with wrong message hash
        bytes memory result = bytes("result");
        bytes memory wrongResult = bytes("wrong result");
        IBN254CertificateVerifierTypes.BN254Certificate memory cert = IBN254CertificateVerifierTypes.BN254Certificate({
            referenceTimestamp: _getTaskReferenceTimestamp(taskHash),
            messageHash: keccak256(wrongResult), // Wrong message hash
            signature: BN254.G1Point(1, 2),
            apk: BN254.G2Point([uint(1), uint(2)], [uint(3), uint(4)]),
            nonSignerWitnesses: new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](0)
        });

        // Submit result should fail with InvalidMessageHash error
        vm.prank(aggregator);
        vm.expectRevert(InvalidMessageHash.selector);
        taskMailbox.submitResult(taskHash, abi.encode(cert), result);
    }

    function test_Revert_WhenECDSACertificateHasInvalidMessageHash() public {
        // Setup executor operator set with ECDSA curve type
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.curveType = IKeyRegistrarTypes.CurveType.ECDSA;

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Get task creation time
        Task memory newTask = taskMailbox.getTaskInfo(newTaskHash);
        uint96 taskCreationTime = newTask.creationTime;

        // Advance time by 1 second to pass TimestampAtCreation check
        vm.warp(block.timestamp + 1);

        // Create ECDSA certificate with wrong message hash
        bytes memory result = bytes("result");
        bytes memory wrongResult = bytes("wrong result");
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert = IECDSACertificateVerifierTypes.ECDSACertificate({
            referenceTimestamp: _getTaskReferenceTimestamp(newTaskHash),
            messageHash: _getECDSACertificateDigest(_getTaskReferenceTimestamp(newTaskHash), wrongResult), // Wrong digest
            sig: bytes("0x1234567890abcdef")
        });

        // Submit result should fail with InvalidMessageHash error
        vm.prank(aggregator);
        vm.expectRevert(InvalidMessageHash.selector);
        taskMailbox.submitResult(newTaskHash, abi.encode(cert), result);
    }
}

// Test contract for refundFee function
contract TaskMailboxUnitTests_refundFee is TaskMailboxUnitTests {
    bytes32 public taskHash;

    function setUp() public override {
        super.setUp();

        // Set up operator set and task config with fee token
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.feeToken = mockToken;
        config.feeCollector = feeCollector;

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create a task with fee
        TaskParams memory taskParams = _createValidTaskParams();

        // Expect Transfer event for fee transfer from creator to taskMailbox
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(creator, address(taskMailbox), avsFee);

        vm.prank(creator);
        taskHash = taskMailbox.createTask(taskParams);
    }

    function test_refundFee_Success() public {
        // Move time forward to expire the task
        vm.warp(block.timestamp + taskSLA + 1);

        // Check initial balances
        uint mailboxBalanceBefore = mockToken.balanceOf(address(taskMailbox));
        uint refundCollectorBalanceBefore = mockToken.balanceOf(refundCollector);

        // Expect Transfer event for fee refund from taskMailbox to refundCollector
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(address(taskMailbox), refundCollector, avsFee);

        // Refund fee as refund collector
        vm.expectEmit(true, true, false, true);
        emit FeeRefunded(refundCollector, taskHash, avsFee);

        vm.prank(refundCollector);
        taskMailbox.refundFee(taskHash);

        // Verify balances changed correctly
        assertEq(mockToken.balanceOf(address(taskMailbox)), mailboxBalanceBefore - avsFee);
        assertEq(mockToken.balanceOf(refundCollector), refundCollectorBalanceBefore + avsFee);

        // Verify task state
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        assertTrue(task.isFeeRefunded);
        assertEq(uint8(task.status), uint8(TaskStatus.EXPIRED));
    }

    function test_Revert_refundFee_OnlyRefundCollector() public {
        // Move time forward to expire the task
        vm.warp(block.timestamp + taskSLA + 1);

        // Try to refund as someone else (not refund collector)
        vm.prank(creator);
        vm.expectRevert(OnlyRefundCollector.selector);
        taskMailbox.refundFee(taskHash);

        // Try as a random address
        vm.prank(address(0x1234));
        vm.expectRevert(OnlyRefundCollector.selector);
        taskMailbox.refundFee(taskHash);
    }

    function test_Revert_refundFee_FeeAlreadyRefunded() public {
        // Move time forward to expire the task
        vm.warp(block.timestamp + taskSLA + 1);

        // First refund should succeed
        vm.prank(refundCollector);
        taskMailbox.refundFee(taskHash);

        // Second refund should fail
        vm.prank(refundCollector);
        vm.expectRevert(FeeAlreadyRefunded.selector);
        taskMailbox.refundFee(taskHash);
    }

    function test_Revert_refundFee_TaskNotExpired() public {
        // Try to refund before task expires
        vm.prank(refundCollector);
        vm.expectRevert(abi.encodeWithSelector(InvalidTaskStatus.selector, TaskStatus.EXPIRED, TaskStatus.CREATED));
        taskMailbox.refundFee(taskHash);
    }

    function test_Revert_refundFee_TaskAlreadyVerified() public {
        // Get task creation time
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        uint96 taskCreationTime = task.creationTime;

        // Submit result to verify the task
        bytes memory result = bytes("result");
        IBN254CertificateVerifierTypes.BN254Certificate memory cert = IBN254CertificateVerifierTypes.BN254Certificate({
            referenceTimestamp: _getTaskReferenceTimestamp(taskHash),
            messageHash: keccak256(result),
            signature: BN254.G1Point(1, 2),
            apk: BN254.G2Point([uint(3), uint(4)], [uint(5), uint(6)]),
            nonSignerWitnesses: new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](0)
        });

        // Mock certificate verification
        vm.mockCall(
            address(mockBN254CertificateVerifier),
            abi.encodeWithSelector(IBN254CertificateVerifier.verifyCertificateProportion.selector),
            abi.encode(true)
        );

        vm.prank(aggregator);
        vm.warp(block.timestamp + 1);
        taskMailbox.submitResult(taskHash, abi.encode(cert), result);

        // Move time forward to what would be expiry
        vm.warp(block.timestamp + taskSLA + 1);

        // Try to refund - should fail because task is verified
        vm.prank(refundCollector);
        vm.expectRevert(abi.encodeWithSelector(InvalidTaskStatus.selector, TaskStatus.EXPIRED, TaskStatus.VERIFIED));
        taskMailbox.refundFee(taskHash);
    }

    function test_refundFee_NoFeeToken() public {
        // Create a task without fee token
        OperatorSet memory operatorSet = OperatorSet(avs2, executorOperatorSetId2);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        config.feeToken = IERC20(address(0));

        vm.prank(avs2);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        TaskParams memory taskParams =
            TaskParams({refundCollector: refundCollector, executorOperatorSet: operatorSet, payload: bytes("test payload")});

        vm.prank(creator);
        bytes32 noFeeTaskHash = taskMailbox.createTask(taskParams);

        // Move time forward to expire the task
        vm.warp(block.timestamp + taskSLA + 1);

        // Refund should succeed but no transfer should occur
        uint mailboxBalanceBefore = mockToken.balanceOf(address(taskMailbox));
        uint refundCollectorBalanceBefore = mockToken.balanceOf(refundCollector);

        vm.prank(refundCollector);
        taskMailbox.refundFee(noFeeTaskHash);

        // Balances should not change since there's no fee token
        assertEq(mockToken.balanceOf(address(taskMailbox)), mailboxBalanceBefore);
        assertEq(mockToken.balanceOf(refundCollector), refundCollectorBalanceBefore);

        // Task should still be marked as refunded
        Task memory task = taskMailbox.getTaskInfo(noFeeTaskHash);
        assertTrue(task.isFeeRefunded);
    }

    function test_refundFee_WithFeeSplit() public {
        // Setup fee split
        uint16 feeSplit = 3000; // 30%
        address localFeeSplitCollector = address(0x789);
        vm.prank(owner);
        taskMailbox.setFeeSplit(feeSplit);
        vm.prank(owner);
        taskMailbox.setFeeSplitCollector(localFeeSplitCollector);

        // Create task
        TaskParams memory taskParams = _createValidTaskParams();

        // Expect Transfer event for fee transfer from creator to taskMailbox
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(creator, address(taskMailbox), avsFee);

        vm.prank(creator);
        bytes32 newTaskHash = taskMailbox.createTask(taskParams);

        // Check initial balance
        uint refundCollectorBalanceBefore = mockToken.balanceOf(refundCollector);

        // Move time forward to expire the task
        vm.warp(block.timestamp + taskSLA + 1);

        // Expect Transfer event for fee refund from taskMailbox to refundCollector
        // Note: fee split doesn't apply to refunds, full amount is refunded
        vm.expectEmit(true, true, false, true, address(mockToken));
        emit IERC20.Transfer(address(taskMailbox), refundCollector, avsFee);

        // Refund fee
        vm.prank(refundCollector);
        taskMailbox.refundFee(newTaskHash);

        // Verify full fee was refunded (fee split doesn't apply to refunds)
        assertEq(mockToken.balanceOf(refundCollector), refundCollectorBalanceBefore + avsFee);

        // Verify fee split collector got nothing
        assertEq(mockToken.balanceOf(feeSplitCollector), 0);
    }

    function test_refundFee_ZeroFee() public {
        // Set mock to return 0 fee
        mockTaskHook.setDefaultFee(0);

        // Create a task with 0 fee
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        bytes32 zeroFeeTaskHash = taskMailbox.createTask(taskParams);

        // Move time forward to expire the task
        vm.warp(block.timestamp + taskSLA + 1);

        // Refund should succeed but no transfer should occur
        uint mailboxBalanceBefore = mockToken.balanceOf(address(taskMailbox));
        uint refundCollectorBalanceBefore = mockToken.balanceOf(refundCollector);

        vm.prank(refundCollector);
        taskMailbox.refundFee(zeroFeeTaskHash);

        // Balances should not change since fee is 0
        assertEq(mockToken.balanceOf(address(taskMailbox)), mailboxBalanceBefore);
        assertEq(mockToken.balanceOf(refundCollector), refundCollectorBalanceBefore);

        // Task should still be marked as refunded
        Task memory task = taskMailbox.getTaskInfo(zeroFeeTaskHash);
        assertTrue(task.isFeeRefunded);
    }
}

// Test contract for view functions
contract TaskMailboxUnitTests_ViewFunctions is TaskMailboxUnitTests {
    bytes32 public taskHash;
    OperatorSet public operatorSet;

    function setUp() public override {
        super.setUp();

        // Set up executor operator set task config
        operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create a task
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        taskHash = taskMailbox.createTask(taskParams);
    }

    function test_ViewFunctions() public {
        // Test that we can read the immutable certificate verifiers
        assertEq(taskMailbox.BN254_CERTIFICATE_VERIFIER(), address(mockBN254CertificateVerifier));
        assertEq(taskMailbox.ECDSA_CERTIFICATE_VERIFIER(), address(mockECDSACertificateVerifier));
        assertEq(taskMailbox.version(), "1.0.0");
        assertEq(taskMailbox.owner(), owner);

        // Test fee split getters
        assertEq(taskMailbox.feeSplit(), 0); // Default value from initialization
        assertEq(taskMailbox.feeSplitCollector(), feeSplitCollector); // Default value from initialization
    }

    function test_getExecutorOperatorSetTaskConfig() public {
        ExecutorOperatorSetTaskConfig memory config = taskMailbox.getExecutorOperatorSetTaskConfig(operatorSet);

        assertEq(uint8(config.curveType), uint8(IKeyRegistrarTypes.CurveType.BN254));
        assertEq(address(config.taskHook), address(mockTaskHook));
        assertEq(address(config.feeToken), address(mockToken));
        assertEq(config.feeCollector, feeCollector);
        assertEq(config.taskSLA, taskSLA);
        assertEq(uint8(config.consensus.consensusType), uint8(ConsensusType.STAKE_PROPORTION_THRESHOLD));
        uint16 decodedThreshold = abi.decode(config.consensus.value, (uint16));
        assertEq(decodedThreshold, stakeProportionThreshold);
        assertEq(config.taskMetadata, bytes("test metadata"));
    }

    function test_getExecutorOperatorSetTaskConfig_Unregistered() public {
        OperatorSet memory unregisteredSet = OperatorSet(avs2, 99);
        ExecutorOperatorSetTaskConfig memory config = taskMailbox.getExecutorOperatorSetTaskConfig(unregisteredSet);

        // Should return empty config
        assertEq(uint8(config.curveType), uint8(IKeyRegistrarTypes.CurveType.NONE));
        assertEq(address(config.taskHook), address(0));
        assertEq(address(config.feeToken), address(0));
        assertEq(config.feeCollector, address(0));
        assertEq(config.taskSLA, 0);
        assertEq(config.consensus.value.length, 0);
        assertEq(config.taskMetadata, bytes(""));
    }

    function test_getTaskInfo() public {
        Task memory task = taskMailbox.getTaskInfo(taskHash);

        assertEq(task.creator, creator);
        assertEq(task.creationTime, block.timestamp);
        assertEq(uint8(task.status), uint8(TaskStatus.CREATED));
        assertEq(task.avs, avs);
        assertEq(task.executorOperatorSetId, executorOperatorSetId);
        assertEq(task.refundCollector, refundCollector);
        assertEq(task.avsFee, avsFee);
        assertEq(task.feeSplit, 0);
        assertEq(task.operatorTableReferenceTimestamp, _getTaskReferenceTimestamp(taskHash)); // New field check
        assertEq(task.payload, bytes("test payload"));
        assertEq(task.executorCert, bytes(""));
        assertEq(task.result, bytes(""));
    }

    function test_getTaskInfo_NonExistentTask() public {
        bytes32 nonExistentHash = keccak256("non-existent");
        Task memory task = taskMailbox.getTaskInfo(nonExistentHash);

        // Should return empty task with NONE status (default for non-existent tasks)
        assertEq(task.creator, address(0));
        assertEq(task.creationTime, 0);
        assertEq(uint8(task.status), uint8(TaskStatus.NONE)); // Non-existent tasks show as NONE
        assertEq(task.avs, address(0));
        assertEq(task.executorOperatorSetId, 0);
    }

    function test_getTaskStatus_Created() public {
        TaskStatus status = taskMailbox.getTaskStatus(taskHash);
        assertEq(uint8(status), uint8(TaskStatus.CREATED));
    }

    function test_getTaskStatus_Verified() public {
        // Get task creation time
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        uint96 taskCreationTime = task.creationTime;

        vm.warp(block.timestamp + 1);
        bytes memory result = bytes("result");
        IBN254CertificateVerifier.BN254Certificate memory cert =
            _createValidBN254CertificateForResult(result, _getTaskReferenceTimestamp(taskHash));

        vm.prank(aggregator);
        taskMailbox.submitResult(taskHash, abi.encode(cert), result);

        TaskStatus status = taskMailbox.getTaskStatus(taskHash);
        assertEq(uint8(status), uint8(TaskStatus.VERIFIED));
    }

    function test_getTaskStatus_Expired() public {
        // Advance time past SLA
        vm.warp(block.timestamp + taskSLA + 1);

        TaskStatus status = taskMailbox.getTaskStatus(taskHash);
        assertEq(uint8(status), uint8(TaskStatus.EXPIRED));
    }

    function test_getTaskStatus_None() public {
        // Get status of non-existent task
        bytes32 nonExistentHash = keccak256("non-existent");
        TaskStatus status = taskMailbox.getTaskStatus(nonExistentHash);
        assertEq(uint8(status), uint8(TaskStatus.NONE));
    }

    function test_getTaskInfo_Expired() public {
        // Advance time past SLA
        vm.warp(block.timestamp + taskSLA + 1);

        Task memory task = taskMailbox.getTaskInfo(taskHash);

        // getTaskInfo should return Expired status
        assertEq(uint8(task.status), uint8(TaskStatus.EXPIRED));
    }

    function test_getTaskResult() public {
        // Get task creation time
        Task memory task = taskMailbox.getTaskInfo(taskHash);
        uint96 taskCreationTime = task.creationTime;

        // Submit result first
        vm.warp(block.timestamp + 1);
        bytes memory expectedResult = bytes("test result");
        IBN254CertificateVerifier.BN254Certificate memory cert =
            _createValidBN254CertificateForResult(expectedResult, _getTaskReferenceTimestamp(taskHash));

        vm.prank(aggregator);
        taskMailbox.submitResult(taskHash, abi.encode(cert), expectedResult);

        // Get result
        bytes memory result = taskMailbox.getTaskResult(taskHash);
        assertEq(result, expectedResult);
    }

    function test_Revert_getTaskResult_NotVerified() public {
        vm.expectRevert(abi.encodeWithSelector(InvalidTaskStatus.selector, TaskStatus.VERIFIED, TaskStatus.CREATED));
        taskMailbox.getTaskResult(taskHash);
    }

    function test_Revert_getTaskResult_Expired() public {
        vm.warp(block.timestamp + taskSLA + 1);

        vm.expectRevert(abi.encodeWithSelector(InvalidTaskStatus.selector, TaskStatus.VERIFIED, TaskStatus.EXPIRED));
        taskMailbox.getTaskResult(taskHash);
    }

    function test_Revert_getTaskResult_None() public {
        bytes32 nonExistentHash = keccak256("non-existent");

        vm.expectRevert(abi.encodeWithSelector(InvalidTaskStatus.selector, TaskStatus.VERIFIED, TaskStatus.NONE));
        taskMailbox.getTaskResult(nonExistentHash);
    }
}

// Test contract for storage variables
contract TaskMailboxUnitTests_Storage is TaskMailboxUnitTests {
    function test_globalTaskCount() public {
        // Set up executor operator set task config
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create multiple tasks and verify the count through task hashes
        TaskParams memory taskParams = _createValidTaskParams();

        // First task should have count 0
        bytes32 expectedHash0 = keccak256(abi.encode(0, address(taskMailbox), block.chainid, taskParams));
        vm.prank(creator);
        bytes32 taskHash0 = taskMailbox.createTask(taskParams);
        assertEq(taskHash0, expectedHash0);

        // Second task should have count 1
        bytes32 expectedHash1 = keccak256(abi.encode(1, address(taskMailbox), block.chainid, taskParams));
        vm.prank(creator);
        bytes32 taskHash1 = taskMailbox.createTask(taskParams);
        assertEq(taskHash1, expectedHash1);

        // Third task should have count 2
        bytes32 expectedHash2 = keccak256(abi.encode(2, address(taskMailbox), block.chainid, taskParams));
        vm.prank(creator);
        bytes32 taskHash2 = taskMailbox.createTask(taskParams);
        assertEq(taskHash2, expectedHash2);
    }

    function test_isExecutorOperatorSetRegistered() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();

        // Initially not registered
        assertFalse(taskMailbox.isExecutorOperatorSetRegistered(operatorSet.key()));

        // Set config first (requirement for registerExecutorOperatorSet)
        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // After setting config, it should be automatically registered
        assertTrue(taskMailbox.isExecutorOperatorSetRegistered(operatorSet.key()));

        // Unregister
        vm.prank(avs);
        taskMailbox.registerExecutorOperatorSet(operatorSet, false);
        assertFalse(taskMailbox.isExecutorOperatorSetRegistered(operatorSet.key()));
    }

    function test_getExecutorOperatorSetTaskConfig() public {
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();

        // Set config
        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Access config via getExecutorOperatorSetTaskConfig function
        ExecutorOperatorSetTaskConfig memory storedConfig = taskMailbox.getExecutorOperatorSetTaskConfig(operatorSet);

        assertEq(uint8(storedConfig.curveType), uint8(config.curveType));
        assertEq(address(storedConfig.taskHook), address(config.taskHook));
        assertEq(address(storedConfig.feeToken), address(config.feeToken));
        assertEq(storedConfig.feeCollector, config.feeCollector);
        assertEq(storedConfig.taskSLA, config.taskSLA);
        assertEq(uint8(storedConfig.consensus.consensusType), uint8(config.consensus.consensusType));
        assertEq(storedConfig.consensus.value, config.consensus.value);
        assertEq(storedConfig.taskMetadata, config.taskMetadata);
    }

    function test_tasks() public {
        // Set up executor operator set task config
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();

        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Create a task
        TaskParams memory taskParams = _createValidTaskParams();
        vm.prank(creator);
        bytes32 taskHash = taskMailbox.createTask(taskParams);

        // Access task via getTaskInfo public function
        Task memory task = taskMailbox.getTaskInfo(taskHash);

        assertEq(task.creator, creator);
        assertEq(task.creationTime, block.timestamp);
        assertEq(uint8(task.status), uint8(TaskStatus.CREATED));
        assertEq(task.avs, avs);
        assertEq(task.executorOperatorSetId, executorOperatorSetId);
        assertEq(task.refundCollector, refundCollector);
        assertEq(task.avsFee, avsFee);
        assertEq(task.feeSplit, 0);
        assertEq(task.payload, bytes("test payload"));
        assertEq(task.executorCert, bytes(""));
        assertEq(task.result, bytes(""));
    }
}

// Test contract for upgradeable functionality
contract TaskMailboxUnitTests_Upgradeable is TaskMailboxUnitTests {
    function test_Initialize_OnlyOnce() public {
        // Try to initialize again, should revert
        vm.expectRevert("Initializable: contract is already initialized");
        taskMailbox.initialize(address(0x9999), 0, feeSplitCollector);
    }

    function test_Implementation_CannotBeInitialized() public {
        // Deploy a new implementation
        TaskMailbox newImpl = new TaskMailbox(address(mockBN254CertificateVerifier), address(mockECDSACertificateVerifier), "1.0.1");

        // Try to initialize the implementation directly, should revert
        vm.expectRevert("Initializable: contract is already initialized");
        newImpl.initialize(owner, 0, feeSplitCollector);
    }

    function test_ProxyUpgrade() public {
        address newOwner = address(0x1234);

        // Deploy new implementation with different version
        TaskMailbox newImpl = new TaskMailbox(address(mockBN254CertificateVerifier), address(mockECDSACertificateVerifier), "2.0.0");

        // Check version before upgrade
        assertEq(taskMailbox.version(), "1.0.0");

        // Upgrade proxy to new implementation
        proxyAdmin.upgrade(ITransparentUpgradeableProxy(address(taskMailbox)), address(newImpl));

        // Check version after upgrade
        assertEq(taskMailbox.version(), "2.0.0");

        // Verify state is preserved (owner should still be the same)
        assertEq(taskMailbox.owner(), owner);
    }

    function test_ProxyAdmin_OnlyOwnerCanUpgrade() public {
        address attacker = address(0x9999);

        // Deploy new implementation
        TaskMailbox newImpl = new TaskMailbox(address(mockBN254CertificateVerifier), address(mockECDSACertificateVerifier), "2.0.0");

        // Try to upgrade from non-owner, should revert
        vm.prank(attacker);
        vm.expectRevert("Ownable: caller is not the owner");
        proxyAdmin.upgrade(ITransparentUpgradeableProxy(address(taskMailbox)), address(newImpl));
    }

    function test_ProxyAdmin_CannotCallImplementation() public {
        // ProxyAdmin should not be able to call implementation functions
        vm.prank(address(proxyAdmin));
        vm.expectRevert("TransparentUpgradeableProxy: admin cannot fallback to proxy target");
        TaskMailbox(payable(address(taskMailbox))).owner();
    }

    function test_StorageSlotConsistency_AfterUpgrade() public {
        address newOwner = address(0x1234);

        // First, make some state changes
        vm.prank(owner);
        taskMailbox.transferOwnership(newOwner);
        assertEq(taskMailbox.owner(), newOwner);

        // Set up an executor operator set
        OperatorSet memory operatorSet = OperatorSet(avs, executorOperatorSetId);
        ExecutorOperatorSetTaskConfig memory config = _createValidExecutorOperatorSetTaskConfig();
        vm.prank(avs);
        taskMailbox.setExecutorOperatorSetTaskConfig(operatorSet, config);

        // Verify config is set
        ExecutorOperatorSetTaskConfig memory retrievedConfig = taskMailbox.getExecutorOperatorSetTaskConfig(operatorSet);
        assertEq(address(retrievedConfig.taskHook), address(config.taskHook));

        // Deploy new implementation
        TaskMailbox newImpl = new TaskMailbox(address(mockBN254CertificateVerifier), address(mockECDSACertificateVerifier), "2.0.0");

        // Upgrade
        vm.prank(address(this)); // proxyAdmin owner
        proxyAdmin.upgrade(ITransparentUpgradeableProxy(address(taskMailbox)), address(newImpl));

        // Verify all state is preserved after upgrade
        assertEq(taskMailbox.owner(), newOwner);
        assertEq(taskMailbox.version(), "2.0.0");

        // Verify the executor operator set config is still there
        ExecutorOperatorSetTaskConfig memory configAfterUpgrade = taskMailbox.getExecutorOperatorSetTaskConfig(operatorSet);
        assertEq(address(configAfterUpgrade.taskHook), address(config.taskHook));
        assertEq(configAfterUpgrade.taskSLA, config.taskSLA);
        assertEq(uint8(configAfterUpgrade.consensus.consensusType), uint8(config.consensus.consensusType));
        assertEq(configAfterUpgrade.consensus.value, config.consensus.value);
    }

    function test_InitializerModifier_PreventsReinitialization() public {
        // Deploy a new proxy without initialization data
        TransparentUpgradeableProxy uninitializedProxy = new TransparentUpgradeableProxy(
            address(new TaskMailbox(address(mockBN254CertificateVerifier), address(mockECDSACertificateVerifier), "1.0.0")),
            address(new ProxyAdmin()),
            ""
        );

        TaskMailbox uninitializedTaskMailbox = TaskMailbox(address(uninitializedProxy));

        // Initialize it once
        uninitializedTaskMailbox.initialize(owner, 0, feeSplitCollector);
        assertEq(uninitializedTaskMailbox.owner(), owner);

        // Try to initialize again, should fail
        vm.expectRevert("Initializable: contract is already initialized");
        uninitializedTaskMailbox.initialize(address(0x9999), 0, feeSplitCollector);
    }
}
