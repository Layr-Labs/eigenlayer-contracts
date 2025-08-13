// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {ReentrancyGuardUpgradeable} from "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {SafeCast} from "@openzeppelin/contracts/utils/math/SafeCast.sol";

import {IAVSTaskHook} from "../../interfaces/IAVSTaskHook.sol";
import {
    IBN254CertificateVerifier, IBN254CertificateVerifierTypes
} from "../../interfaces/IBN254CertificateVerifier.sol";
import {
    IECDSACertificateVerifier, IECDSACertificateVerifierTypes
} from "../../interfaces/IECDSACertificateVerifier.sol";
import {IBaseCertificateVerifier} from "../../interfaces/IBaseCertificateVerifier.sol";
import {IKeyRegistrarTypes} from "../../interfaces/IKeyRegistrar.sol";
import {ITaskMailbox} from "../../interfaces/ITaskMailbox.sol";
import {OperatorSet} from "../../libraries/OperatorSetLib.sol";
import {SemVerMixin} from "../../mixins/SemVerMixin.sol";
import {TaskMailboxStorage} from "./TaskMailboxStorage.sol";

/**
 * @title TaskMailbox
 * @author Layr Labs, Inc.
 * @notice Contract for managing the lifecycle of tasks that are executed by operator sets of task-based AVSs.
 */
contract TaskMailbox is
    Initializable,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable,
    TaskMailboxStorage,
    SemVerMixin
{
    using SafeERC20 for IERC20;
    using SafeCast for *;

    /**
     * @notice Constructor for TaskMailbox
     * @param _bn254CertificateVerifier Address of the BN254 certificate verifier
     * @param _ecdsaCertificateVerifier Address of the ECDSA certificate verifier
     * @param _maxTaskSLA Maximum task SLA in seconds
     * @param _version The semantic version of the contract
     */
    constructor(
        address _bn254CertificateVerifier,
        address _ecdsaCertificateVerifier,
        uint96 _maxTaskSLA,
        string memory _version
    ) TaskMailboxStorage(_bn254CertificateVerifier, _ecdsaCertificateVerifier, _maxTaskSLA) SemVerMixin(_version) {
        _disableInitializers();
    }

    /**
     * @notice Initializer for TaskMailbox
     * @param _owner The owner of the contract
     * @param _feeSplit The initial fee split in basis points
     * @param _feeSplitCollector The initial fee split collector address
     */
    function initialize(address _owner, uint16 _feeSplit, address _feeSplitCollector) external initializer {
        __Ownable_init();
        __ReentrancyGuard_init();
        _transferOwnership(_owner);
        _setFeeSplit(_feeSplit);
        _setFeeSplitCollector(_feeSplitCollector);
    }

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

    /// @inheritdoc ITaskMailbox
    function setExecutorOperatorSetTaskConfig(
        OperatorSet memory operatorSet,
        ExecutorOperatorSetTaskConfig memory config
    ) external {
        // Validate config is populated with non-zero values
        require(_isConfigPopulated(config), ExecutorOperatorSetTaskConfigNotSet());

        // Validate task SLA is within maximum limit
        require(config.taskSLA <= MAX_TASK_SLA, TaskSLAExceedsMaximum());

        // Validate consensus enum within range
        _validateConsensus(config.consensus);
        // Validate operator set ownership
        _validateOperatorSetOwner(operatorSet, config.curveType);

        _executorOperatorSetTaskConfigs[operatorSet.key()] = config;
        emit ExecutorOperatorSetTaskConfigSet(msg.sender, operatorSet.avs, operatorSet.id, config);

        // If executor operator set is not registered, register it.
        if (!isExecutorOperatorSetRegistered[operatorSet.key()]) {
            _registerExecutorOperatorSet(operatorSet, true);
        }
    }

    /// @inheritdoc ITaskMailbox
    function registerExecutorOperatorSet(OperatorSet memory operatorSet, bool isRegistered) external {
        ExecutorOperatorSetTaskConfig memory taskConfig = _executorOperatorSetTaskConfigs[operatorSet.key()];

        // Validate that task config has been set before registration can be toggled.
        require(_isConfigPopulated(taskConfig), ExecutorOperatorSetTaskConfigNotSet());

        // Validate operator set ownership
        _validateOperatorSetOwner(operatorSet, taskConfig.curveType);

        _registerExecutorOperatorSet(operatorSet, isRegistered);
    }

    /// @inheritdoc ITaskMailbox
    function createTask(
        TaskParams memory taskParams
    ) external nonReentrant returns (bytes32) {
        require(taskParams.payload.length > 0, PayloadIsEmpty());
        require(
            isExecutorOperatorSetRegistered[taskParams.executorOperatorSet.key()], ExecutorOperatorSetNotRegistered()
        );

        ExecutorOperatorSetTaskConfig memory taskConfig =
            _executorOperatorSetTaskConfigs[taskParams.executorOperatorSet.key()];
        require(_isConfigPopulated(taskConfig), ExecutorOperatorSetTaskConfigNotSet());

        // Get the operator table reference timestamp and max staleness period
        IBaseCertificateVerifier certificateVerifier =
            IBaseCertificateVerifier(_getCertificateVerifier(taskConfig.curveType));
        uint32 operatorTableReferenceTimestamp =
            certificateVerifier.latestReferenceTimestamp(taskParams.executorOperatorSet);
        {
            // Scoping to prevent `Stack too deep` error during compilation
            uint32 maxStaleness = certificateVerifier.maxOperatorTableStaleness(taskParams.executorOperatorSet);
            require(
                maxStaleness == 0
                    || (block.timestamp + taskConfig.taskSLA <= operatorTableReferenceTimestamp + maxStaleness),
                CertificateStale()
            );
        }

        // Pre-task submission checks: AVS can validate the caller and task params.
        taskConfig.taskHook.validatePreTaskCreation(msg.sender, taskParams);

        // Calculate the AVS fee using the task hook
        uint96 avsFee = taskConfig.taskHook.calculateTaskFee(taskParams);

        bytes32 taskHash = keccak256(abi.encode(_globalTaskCount, address(this), block.chainid, taskParams));
        _globalTaskCount = _globalTaskCount + 1;

        _tasks[taskHash] = Task({
            creator: msg.sender,
            creationTime: block.timestamp.toUint96(),
            avs: taskParams.executorOperatorSet.avs,
            avsFee: avsFee,
            refundCollector: taskParams.refundCollector,
            executorOperatorSetId: taskParams.executorOperatorSet.id,
            feeSplit: feeSplit,
            status: TaskStatus.CREATED,
            isFeeRefunded: false,
            operatorTableReferenceTimestamp: operatorTableReferenceTimestamp,
            executorOperatorSetTaskConfig: taskConfig,
            payload: taskParams.payload,
            executorCert: bytes(""),
            result: bytes("")
        });

        // Transfer fee to the TaskMailbox if there's a fee to transfer
        if (taskConfig.feeToken != IERC20(address(0)) && avsFee > 0) {
            require(taskConfig.feeCollector != address(0), InvalidFeeReceiver());
            require(taskParams.refundCollector != address(0), InvalidFeeReceiver());
            taskConfig.feeToken.safeTransferFrom(msg.sender, address(this), avsFee);
        }

        // Post-task submission checks: AVS can write to storage in their hook for validating task lifecycle
        taskConfig.taskHook.handlePostTaskCreation(taskHash);

        emit TaskCreated(
            msg.sender,
            taskHash,
            taskParams.executorOperatorSet.avs,
            taskParams.executorOperatorSet.id,
            operatorTableReferenceTimestamp,
            taskParams.refundCollector,
            avsFee,
            block.timestamp + taskConfig.taskSLA,
            taskParams.payload
        );
        return taskHash;
    }

    /// @inheritdoc ITaskMailbox
    function submitResult(bytes32 taskHash, bytes memory executorCert, bytes memory result) external nonReentrant {
        Task storage task = _tasks[taskHash];
        TaskStatus status = _getTaskStatus(task);
        require(status == TaskStatus.CREATED, InvalidTaskStatus(TaskStatus.CREATED, status));
        require(block.timestamp > task.creationTime, TimestampAtCreation());

        // Pre-task result submission checks: AVS can validate the caller, task result, params and certificate.
        task.executorOperatorSetTaskConfig.taskHook.validatePreTaskResultSubmission(
            msg.sender, taskHash, executorCert, result
        );

        // Verify certificate based on consensus configuration
        OperatorSet memory executorOperatorSet = OperatorSet(task.avs, task.executorOperatorSetId);
        _verifyExecutorCertificate(
            task.executorOperatorSetTaskConfig.curveType,
            task.executorOperatorSetTaskConfig.consensus,
            executorOperatorSet,
            task.operatorTableReferenceTimestamp,
            keccak256(result),
            executorCert
        );

        task.status = TaskStatus.VERIFIED;
        task.executorCert = executorCert;
        task.result = result;

        // Transfer fee to the fee collector if there's a fee to transfer
        if (task.executorOperatorSetTaskConfig.feeToken != IERC20(address(0)) && task.avsFee > 0) {
            // Calculate fee split amount
            uint96 feeSplitAmount = ((uint256(task.avsFee) * task.feeSplit) / ONE_HUNDRED_IN_BIPS).toUint96();

            // Transfer split to fee split collector if there's a split
            if (feeSplitAmount > 0) {
                task.executorOperatorSetTaskConfig.feeToken.safeTransfer(feeSplitCollector, feeSplitAmount);
            }

            // Transfer remaining fee to AVS fee collector
            uint96 avsAmount = task.avsFee - feeSplitAmount;
            if (avsAmount > 0) {
                task.executorOperatorSetTaskConfig.feeToken.safeTransfer(
                    task.executorOperatorSetTaskConfig.feeCollector, avsAmount
                );
            }
        }

        // Post-task result submission checks: AVS can update hook storage for task lifecycle if needed.
        task.executorOperatorSetTaskConfig.taskHook.handlePostTaskResultSubmission(msg.sender, taskHash);

        emit TaskVerified(msg.sender, taskHash, task.avs, task.executorOperatorSetId, task.executorCert, task.result);
    }

    /// @inheritdoc ITaskMailbox
    function refundFee(
        bytes32 taskHash
    ) external nonReentrant {
        Task storage task = _tasks[taskHash];
        require(task.refundCollector == msg.sender, OnlyRefundCollector());
        require(!task.isFeeRefunded, FeeAlreadyRefunded());

        TaskStatus status = _getTaskStatus(task);
        require(status == TaskStatus.EXPIRED, InvalidTaskStatus(TaskStatus.EXPIRED, status));

        // Mark fee as refunded to prevent double refunds
        task.isFeeRefunded = true;

        // Transfer fee to refund collector if there's a fee to refund
        if (task.executorOperatorSetTaskConfig.feeToken != IERC20(address(0)) && task.avsFee > 0) {
            task.executorOperatorSetTaskConfig.feeToken.safeTransfer(task.refundCollector, task.avsFee);
        }

        emit FeeRefunded(task.refundCollector, taskHash, task.avsFee);
    }

    /// @inheritdoc ITaskMailbox
    function setFeeSplit(
        uint16 _feeSplit
    ) external onlyOwner {
        _setFeeSplit(_feeSplit);
    }

    /// @inheritdoc ITaskMailbox
    function setFeeSplitCollector(
        address _feeSplitCollector
    ) external onlyOwner {
        _setFeeSplitCollector(_feeSplitCollector);
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Sets the fee split percentage
     * @param _feeSplit The fee split in basis points (0-10000)
     */
    function _setFeeSplit(
        uint16 _feeSplit
    ) internal {
        require(_feeSplit <= ONE_HUNDRED_IN_BIPS, InvalidFeeSplit());
        feeSplit = _feeSplit;
        emit FeeSplitSet(_feeSplit);
    }

    /**
     * @notice Sets the fee split collector address
     * @param _feeSplitCollector The address to receive fee splits
     */
    function _setFeeSplitCollector(
        address _feeSplitCollector
    ) internal {
        require(_feeSplitCollector != address(0), InvalidAddressZero());
        feeSplitCollector = _feeSplitCollector;
        emit FeeSplitCollectorSet(_feeSplitCollector);
    }

    /**
     * @notice Gets the current status of a task
     * @param task The task to get the status for
     * @return The current status of the task, considering expiration
     */
    function _getTaskStatus(
        Task memory task
    ) internal view returns (TaskStatus) {
        if (
            task.status == TaskStatus.CREATED
                && block.timestamp > (task.creationTime + task.executorOperatorSetTaskConfig.taskSLA)
        ) {
            return TaskStatus.EXPIRED;
        }
        return task.status;
    }

    /**
     * @notice Registers an executor operator set with the TaskMailbox
     * @param operatorSet The operator set to register
     * @param isRegistered Whether the operator set is registered
     */
    function _registerExecutorOperatorSet(OperatorSet memory operatorSet, bool isRegistered) internal {
        isExecutorOperatorSetRegistered[operatorSet.key()] = isRegistered;
        emit ExecutorOperatorSetRegistered(msg.sender, operatorSet.avs, operatorSet.id, isRegistered);
    }

    /**
     * @notice Gets the certificate verifier for a given curve type
     * @param curveType The curve type to get the certificate verifier for
     * @return The address of the certificate verifier
     * @dev This function will revert if the curve type is invalid
     */
    function _getCertificateVerifier(
        IKeyRegistrarTypes.CurveType curveType
    ) internal view returns (address) {
        if (curveType == IKeyRegistrarTypes.CurveType.BN254) {
            return BN254_CERTIFICATE_VERIFIER;
        } else if (curveType == IKeyRegistrarTypes.CurveType.ECDSA) {
            return ECDSA_CERTIFICATE_VERIFIER;
        } else {
            revert InvalidCurveType();
        }
    }

    /**
     * @notice Checks if a task config is populated
     * @param taskConfig The task config to check
     * @return True if all elements of the task config are populated, false otherwise
     */
    function _isConfigPopulated(
        ExecutorOperatorSetTaskConfig memory taskConfig
    ) internal pure returns (bool) {
        return taskConfig.curveType != IKeyRegistrarTypes.CurveType.NONE
            && taskConfig.taskHook != IAVSTaskHook(address(0)) && taskConfig.taskSLA > 0;
    }

    /**
     * @notice Validates that the caller is the owner of the operator set
     * @param operatorSet The operator set to validate ownership for
     * @param curveType The curve type used to determine the certificate verifier
     * @dev This function will revert if the curve type is invalid or if msg.sender is not the owner of the operator set
     */
    function _validateOperatorSetOwner(
        OperatorSet memory operatorSet,
        IKeyRegistrarTypes.CurveType curveType
    ) internal view {
        address certificateVerifier = _getCertificateVerifier(curveType);

        require(
            IBaseCertificateVerifier(certificateVerifier).getOperatorSetOwner(operatorSet) == msg.sender,
            InvalidOperatorSetOwner()
        );
    }

    /**
     * @notice Validates the consensus configuration
     * @param consensus The consensus configuration to validate
     */
    function _validateConsensus(
        Consensus memory consensus
    ) internal pure {
        if (consensus.consensusType == ConsensusType.NONE) {
            // For NONE consensus type, value should be empty bytes
            require(consensus.value.length == 0, InvalidConsensusValue());
        } else if (consensus.consensusType == ConsensusType.STAKE_PROPORTION_THRESHOLD) {
            // Decode and validate the stake proportion threshold
            require(consensus.value.length == 32, InvalidConsensusValue());
            uint16 stakeProportionThreshold = abi.decode(consensus.value, (uint16));
            require(stakeProportionThreshold <= ONE_HUNDRED_IN_BIPS, InvalidConsensusValue());
        } else {
            revert InvalidConsensusType();
        }
    }

    /**
     * @notice Validates a BN254 certificate's basic requirements
     * @param cert The BN254 certificate to validate
     * @param operatorTableReferenceTimestamp The expected reference timestamp
     * @param resultHash The expected message hash
     */
    function _validateBN254Certificate(
        IBN254CertificateVerifierTypes.BN254Certificate memory cert,
        uint32 operatorTableReferenceTimestamp,
        bytes32 resultHash
    ) internal pure {
        require(cert.referenceTimestamp == operatorTableReferenceTimestamp, InvalidReferenceTimestamp());
        require(cert.messageHash == resultHash, InvalidMessageHash());
        require(!(cert.signature.X == 0 && cert.signature.Y == 0), EmptyCertificateSignature());
    }

    /**
     * @notice Validates an ECDSA certificate's basic requirements
     * @param cert The ECDSA certificate to validate
     * @param operatorTableReferenceTimestamp The expected reference timestamp
     * @param resultHash The expected message hash
     */
    function _validateECDSACertificate(
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert,
        uint32 operatorTableReferenceTimestamp,
        bytes32 resultHash
    ) internal pure {
        require(cert.referenceTimestamp == operatorTableReferenceTimestamp, InvalidReferenceTimestamp());
        require(cert.messageHash == resultHash, InvalidMessageHash());
        require(cert.sig.length > 0, EmptyCertificateSignature());
    }

    /**
     * @notice Verifies an executor certificate based on the consensus configuration
     * @param curveType The curve type used for signature verification
     * @param consensus The consensus configuration
     * @param executorOperatorSet The executor operator set
     * @param operatorTableReferenceTimestamp The reference timestamp of the operator table
     * @param resultHash The hash of the result of the task
     * @param executorCert The executor certificate to verify
     */
    function _verifyExecutorCertificate(
        IKeyRegistrarTypes.CurveType curveType,
        Consensus memory consensus,
        OperatorSet memory executorOperatorSet,
        uint32 operatorTableReferenceTimestamp,
        bytes32 resultHash,
        bytes memory executorCert
    ) internal {
        if (consensus.consensusType == ConsensusType.NONE) {
            // For NONE consensus type, just verify the certificate.
            // The consensus logic handling is left to the AVS in the `handlePostTaskResultSubmission` hook.

            // Verify certificate based on curve type
            if (curveType == IKeyRegistrarTypes.CurveType.BN254) {
                // BN254 Certificate verification
                IBN254CertificateVerifierTypes.BN254Certificate memory bn254Cert =
                    abi.decode(executorCert, (IBN254CertificateVerifierTypes.BN254Certificate));

                // Validate the certificate
                _validateBN254Certificate(bn254Cert, operatorTableReferenceTimestamp, resultHash);

                IBN254CertificateVerifier(BN254_CERTIFICATE_VERIFIER).verifyCertificate(executorOperatorSet, bn254Cert);
            } else if (curveType == IKeyRegistrarTypes.CurveType.ECDSA) {
                // ECDSA Certificate verification
                IECDSACertificateVerifierTypes.ECDSACertificate memory ecdsaCert =
                    abi.decode(executorCert, (IECDSACertificateVerifierTypes.ECDSACertificate));

                // Validate the certificate
                _validateECDSACertificate(ecdsaCert, operatorTableReferenceTimestamp, resultHash);

                IECDSACertificateVerifier(ECDSA_CERTIFICATE_VERIFIER).verifyCertificate(executorOperatorSet, ecdsaCert);
            } else {
                revert InvalidCurveType();
            }
        } else if (consensus.consensusType == ConsensusType.STAKE_PROPORTION_THRESHOLD) {
            // Decode stake proportion threshold
            uint16 stakeProportionThreshold = abi.decode(consensus.value, (uint16));
            // This assumes that that the `BN254TableCalculator` or `ECDSATableCalculator` have been set as the `OperatorTableCalculator` in the `CrossChainRegistry` on the source chain while creating a generation reservation.
            // These table calculators calculate slashable stake for an operatorSet and value all strategies equally. Hence `totalStakeProportionThresholds` is an array of length 1.
            uint16[] memory totalStakeProportionThresholds = new uint16[](1);
            totalStakeProportionThresholds[0] = stakeProportionThreshold;

            bool isThresholdMet;
            // Verify certificate based on curve type
            if (curveType == IKeyRegistrarTypes.CurveType.BN254) {
                // BN254 Certificate verification
                IBN254CertificateVerifierTypes.BN254Certificate memory bn254Cert =
                    abi.decode(executorCert, (IBN254CertificateVerifierTypes.BN254Certificate));

                // Validate the certificate
                _validateBN254Certificate(bn254Cert, operatorTableReferenceTimestamp, resultHash);

                isThresholdMet = IBN254CertificateVerifier(BN254_CERTIFICATE_VERIFIER).verifyCertificateProportion(
                    executorOperatorSet, bn254Cert, totalStakeProportionThresholds
                );
            } else if (curveType == IKeyRegistrarTypes.CurveType.ECDSA) {
                // ECDSA Certificate verification
                IECDSACertificateVerifierTypes.ECDSACertificate memory ecdsaCert =
                    abi.decode(executorCert, (IECDSACertificateVerifierTypes.ECDSACertificate));

                // Validate the certificate
                _validateECDSACertificate(ecdsaCert, operatorTableReferenceTimestamp, resultHash);

                (isThresholdMet,) = IECDSACertificateVerifier(ECDSA_CERTIFICATE_VERIFIER).verifyCertificateProportion(
                    executorOperatorSet, ecdsaCert, totalStakeProportionThresholds
                );
            } else {
                revert InvalidCurveType();
            }
            // If the threshold is not met, revert
            require(isThresholdMet, ThresholdNotMet());
        } else {
            revert InvalidConsensusType();
        }
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc ITaskMailbox
    function getExecutorOperatorSetTaskConfig(
        OperatorSet memory operatorSet
    ) external view returns (ExecutorOperatorSetTaskConfig memory) {
        return _executorOperatorSetTaskConfigs[operatorSet.key()];
    }

    /// @inheritdoc ITaskMailbox
    function getTaskInfo(
        bytes32 taskHash
    ) external view returns (Task memory) {
        Task memory task = _tasks[taskHash];
        return Task(
            task.creator,
            task.creationTime,
            task.avs,
            task.avsFee,
            task.refundCollector,
            task.executorOperatorSetId,
            task.feeSplit,
            _getTaskStatus(task),
            task.isFeeRefunded,
            task.operatorTableReferenceTimestamp,
            task.executorOperatorSetTaskConfig,
            task.payload,
            task.executorCert,
            task.result
        );
    }

    /// @inheritdoc ITaskMailbox
    function getTaskStatus(
        bytes32 taskHash
    ) external view returns (TaskStatus) {
        Task memory task = _tasks[taskHash];
        return _getTaskStatus(task);
    }

    /// @inheritdoc ITaskMailbox
    function getTaskResult(
        bytes32 taskHash
    ) external view returns (bytes memory) {
        Task memory task = _tasks[taskHash];
        TaskStatus status = _getTaskStatus(task);
        require(status == TaskStatus.VERIFIED, InvalidTaskStatus(TaskStatus.VERIFIED, status));
        return task.result;
    }

    /// @inheritdoc ITaskMailbox
    function getBN254CertificateBytes(
        IBN254CertificateVerifierTypes.BN254Certificate memory cert
    ) external pure returns (bytes memory) {
        return abi.encode(cert);
    }

    /// @inheritdoc ITaskMailbox
    function getECDSACertificateBytes(
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert
    ) external pure returns (bytes memory) {
        return abi.encode(cert);
    }
}
