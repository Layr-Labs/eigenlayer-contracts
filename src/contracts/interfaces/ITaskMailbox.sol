// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {IAVSTaskHook} from "./IAVSTaskHook.sol";
import {IBN254CertificateVerifierTypes} from "./IBN254CertificateVerifier.sol";
import {IECDSACertificateVerifierTypes} from "./IECDSACertificateVerifier.sol";
import {IKeyRegistrarTypes} from "./IKeyRegistrar.sol";
import {OperatorSet, OperatorSetLib} from "../libraries/OperatorSetLib.sol";

/**
 * @title ITaskMailboxTypes
 * @notice Interface defining the type structures used in the TaskMailbox
 */
interface ITaskMailboxTypes {
    /**
     * @notice Enum defining the type of consensus mechanism
     */
    enum ConsensusType {
        NONE,
        STAKE_PROPORTION_THRESHOLD
    }

    /**
     * @notice Consensus configuration for task verification
     * @param consensusType The type of consensus mechanism
     * @param value Encoded consensus parameters based on consensusType
     */
    struct Consensus {
        ConsensusType consensusType;
        bytes value;
    }

    /**
     * @notice Configuration for the executor operator set
     * @param taskHook Address of the AVS task hook contract
     * @param taskSLA Time (in seconds) within which the task must be completed
     * @param feeToken ERC20 token used for task fees
     * @param curveType The curve type used for signature verification
     * @param feeCollector Address to receive AVS fees
     * @param consensus Consensus configuration for task verification
     * @param taskMetadata Additional metadata for task execution
     */
    struct ExecutorOperatorSetTaskConfig {
        IAVSTaskHook taskHook;
        uint96 taskSLA;
        IERC20 feeToken;
        IKeyRegistrarTypes.CurveType curveType;
        address feeCollector;
        Consensus consensus;
        bytes taskMetadata;
    }

    /**
     * @notice Parameters for creating a new task
     * @param refundCollector Address to receive refunds if task is not completed
     * @param executorOperatorSet The operator set that will execute the task
     * @param payload Task payload
     */
    struct TaskParams {
        address refundCollector;
        OperatorSet executorOperatorSet;
        bytes payload;
    }

    /**
     * @notice Status of a task in the system
     */
    enum TaskStatus {
        NONE, // 0 - Default value for uninitialized tasks
        CREATED, // 1 - Task has been created
        VERIFIED, // 2 - Task has been verified
        EXPIRED // 3 - Task has expired

    }

    /**
     * @notice Complete task information
     * @param creator Address that created the task
     * @param creationTime Block timestamp when task was created
     * @param avs Address of the AVS handling the task
     * @param avsFee Fee paid to the AVS
     * @param refundCollector Address to receive refunds
     * @param executorOperatorSetId ID of the operator set executing the task
     * @param feeSplit Percentage split of fees taken by the TaskMailbox in basis points
     * @param status Current status of the task
     * @param isFeeRefunded Whether the fee has been refunded
     * @param executorOperatorSetTaskConfig Configuration for executor operator set task execution
     * @param payload Task payload
     * @param executorCert Executor certificate
     * @param result Task execution result data
     */
    struct Task {
        address creator;
        uint96 creationTime;
        address avs;
        uint96 avsFee;
        address refundCollector;
        uint32 executorOperatorSetId;
        uint16 feeSplit;
        TaskStatus status;
        bool isFeeRefunded;
        uint32 operatorTableReferenceTimestamp;
        ExecutorOperatorSetTaskConfig executorOperatorSetTaskConfig;
        bytes payload;
        bytes executorCert;
        bytes result;
    }
}

/**
 * @title ITaskMailboxErrors
 * @notice Interface defining errors that can be thrown by the TaskMailbox
 */
interface ITaskMailboxErrors is ITaskMailboxTypes {
    /// @notice Thrown when a certificate verification fails
    error CertificateVerificationFailed();

    /// @notice Thrown when an executor operator set is not registered
    error ExecutorOperatorSetNotRegistered();

    /// @notice Thrown when an executor operator set task config is not set
    error ExecutorOperatorSetTaskConfigNotSet();

    /// @notice Thrown when an input address is zero
    error InvalidAddressZero();

    /// @notice Thrown when an invalid curve type is provided
    error InvalidCurveType();

    /// @notice Thrown when a task creator is invalid
    error InvalidTaskCreator();

    /// @notice Thrown when a task status is invalid
    /// @param expected The expected task status
    /// @param actual The actual task status
    error InvalidTaskStatus(TaskStatus expected, TaskStatus actual);

    /// @notice Thrown when a payload is empty
    error PayloadIsEmpty();

    /// @notice Thrown when a task SLA is zero
    error TaskSLAIsZero();

    /// @notice Thrown when a timestamp is at creation
    error TimestampAtCreation();

    /// @notice Thrown when an operator set owner is invalid
    error InvalidOperatorSetOwner();

    /// @notice Thrown when an invalid consensus type is provided
    error InvalidConsensusType();

    /// @notice Thrown when an invalid consensus value is provided
    error InvalidConsensusValue();

    /// @notice Thrown when a certificate has an invalid reference timestamp
    error InvalidReferenceTimestamp();

    /// @notice Thrown when a certificate has an empty signature
    error EmptyCertificateSignature();

    /// @notice Thrown when fee receiver is zero address
    error InvalidFeeReceiver();

    /// @notice Thrown when fee has already been refunded
    error FeeAlreadyRefunded();

    /// @notice Thrown when caller is not the refund collector
    error OnlyRefundCollector();

    /// @notice Thrown when fee split value is invalid (> 10000 bips)
    error InvalidFeeSplit();

    /// @notice Thrown when a certificate has an invalid message hash
    error InvalidMessageHash();
}

/**
 * @title ITaskMailboxEvents
 * @notice Interface defining events emitted by the TaskMailbox
 */
interface ITaskMailboxEvents is ITaskMailboxTypes {
    /**
     * @notice Emitted when an executor operator set is registered
     * @param caller Address that called the registration function
     * @param avs Address of the AVS being registered
     * @param executorOperatorSetId ID of the executor operator set
     * @param isRegistered Whether the operator set is registered
     */
    event ExecutorOperatorSetRegistered(
        address indexed caller, address indexed avs, uint32 indexed executorOperatorSetId, bool isRegistered
    );

    /**
     * @notice Emitted when an executor operator set task configuration is set
     * @param caller Address that called the configuration function
     * @param avs Address of the AVS being configured
     * @param executorOperatorSetId ID of the executor operator set
     * @param config The task configuration for the executor operator set
     */
    event ExecutorOperatorSetTaskConfigSet(
        address indexed caller,
        address indexed avs,
        uint32 indexed executorOperatorSetId,
        ExecutorOperatorSetTaskConfig config
    );

    /**
     * @notice Emitted when a new task is created
     * @param creator Address that created the task
     * @param taskHash Unique identifier of the task
     * @param avs Address of the AVS handling the task
     * @param executorOperatorSetId ID of the executor operator set
     * @param operatorTableReferenceTimestamp Reference timestamp of the operator table
     * @param refundCollector Address to receive refunds
     * @param avsFee Fee paid to the AVS
     * @param taskDeadline Timestamp by which the task must be completed
     * @param payload Task payload
     */
    event TaskCreated(
        address indexed creator,
        bytes32 indexed taskHash,
        address indexed avs,
        uint32 executorOperatorSetId,
        uint32 operatorTableReferenceTimestamp,
        address refundCollector,
        uint96 avsFee,
        uint256 taskDeadline,
        bytes payload
    );

    /**
     * @notice Emitted when a task is verified
     * @param aggregator Address that submitted the verification
     * @param taskHash Unique identifier of the task
     * @param avs Address of the AVS handling the task
     * @param executorOperatorSetId ID of the executor operator set
     * @param executorCert Executor certificate
     * @param result Task execution result data
     */
    event TaskVerified(
        address indexed aggregator,
        bytes32 indexed taskHash,
        address indexed avs,
        uint32 executorOperatorSetId,
        bytes executorCert,
        bytes result
    );

    /**
     * @notice Emitted when a task fee is refunded
     * @param refundCollector Address that received the refund
     * @param taskHash Unique identifier of the task
     * @param avsFee Amount of fee refunded
     */
    event FeeRefunded(address indexed refundCollector, bytes32 indexed taskHash, uint96 avsFee);

    /**
     * @notice Emitted when the fee split is updated
     * @param feeSplit The new fee split value in basis points
     */
    event FeeSplitSet(uint16 feeSplit);

    /**
     * @notice Emitted when the fee split collector is updated
     * @param feeSplitCollector The new fee split collector address
     */
    event FeeSplitCollectorSet(address indexed feeSplitCollector);
}

/**
 * @title ITaskMailbox
 * @author Layr Labs, Inc.
 * @notice Interface for the TaskMailbox contract.
 */
interface ITaskMailbox is ITaskMailboxErrors, ITaskMailboxEvents {
    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Initializes the TaskMailbox
     * @param owner The owner of the contract
     * @param feeSplit The initial fee split in basis points
     * @param feeSplitCollector The initial fee split collector address
     */
    function initialize(address owner, uint16 feeSplit, address feeSplitCollector) external;

    /**
     * @notice Sets the task configuration for an executor operator set
     * @param operatorSet The operator set to configure
     * @param config Task configuration for the operator set
     * @dev Fees can be switched off by setting the fee token to the zero address.
     */
    function setExecutorOperatorSetTaskConfig(
        OperatorSet memory operatorSet,
        ExecutorOperatorSetTaskConfig memory config
    ) external;

    /**
     * @notice Registers an executor operator set with the TaskMailbox
     * @param operatorSet The operator set to register
     * @param isRegistered Whether the operator set is registered
     * @dev This function can be called to toggle the registration once the task config has been set.
     */
    function registerExecutorOperatorSet(OperatorSet memory operatorSet, bool isRegistered) external;

    /**
     * @notice Creates a new task
     * @param taskParams Parameters for the task
     * @return taskHash Unique identifier of the created task
     * @dev If the operator set has its fee token set, call `IAVSTaskHook.calculateTaskFee()` to get
     * the fee for the task and approve the TaskMailbox for the fee before calling this function.
     */
    function createTask(
        TaskParams memory taskParams
    ) external returns (bytes32 taskHash);

    /**
     * @notice Submits the result of a task execution
     * @param taskHash Unique identifier of the task
     * @param executorCert Certificate proving the validity of the result
     * @param result Task execution result data
     */
    function submitResult(bytes32 taskHash, bytes memory executorCert, bytes memory result) external;

    /**
     * @notice Refunds the fee for an expired task
     * @param taskHash Unique identifier of the task
     * @dev Can only be called by the refund collector for expired tasks
     */
    function refundFee(
        bytes32 taskHash
    ) external;

    /**
     * @notice Sets the fee split percentage
     * @param feeSplit The fee split in basis points (0-10000)
     * @dev Only callable by the owner
     */
    function setFeeSplit(
        uint16 feeSplit
    ) external;

    /**
     * @notice Sets the fee split collector address
     * @param feeSplitCollector The address to receive fee splits
     * @dev Only callable by the owner
     */
    function setFeeSplitCollector(
        address feeSplitCollector
    ) external;

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     * @notice Checks if an executor operator set is registered
     * @param operatorSetKey Key of the operator set to check
     * @return True if the executor operator set is registered, false otherwise
     */
    function isExecutorOperatorSetRegistered(
        bytes32 operatorSetKey
    ) external view returns (bool);

    /**
     * @notice Gets the task configuration for an executor operator set
     * @param operatorSet The operator set to get configuration for
     * @return Task configuration for the operator set
     */
    function getExecutorOperatorSetTaskConfig(
        OperatorSet memory operatorSet
    ) external view returns (ExecutorOperatorSetTaskConfig memory);

    /**
     * @notice Gets complete information about a task
     * @param taskHash Unique identifier of the task
     * @return Complete task information
     */
    function getTaskInfo(
        bytes32 taskHash
    ) external view returns (Task memory);

    /**
     * @notice Gets the current status of a task
     * @param taskHash Unique identifier of the task
     * @return Current status of the task
     */
    function getTaskStatus(
        bytes32 taskHash
    ) external view returns (TaskStatus);

    /**
     * @notice Gets the result of a verified task
     * @param taskHash Unique identifier of the task
     * @return Result data of the task
     */
    function getTaskResult(
        bytes32 taskHash
    ) external view returns (bytes memory);

    /**
     * @notice Gets the bytes of a BN254 certificate
     * @param cert The certificate to get the bytes of
     * @return The bytes of the certificate
     */
    function getBN254CertificateBytes(
        IBN254CertificateVerifierTypes.BN254Certificate memory cert
    ) external pure returns (bytes memory);

    /**
     * @notice Gets the bytes of a ECDSA certificate
     * @param cert The certificate to get the bytes of
     * @return The bytes of the certificate
     */
    function getECDSACertificateBytes(
        IECDSACertificateVerifierTypes.ECDSACertificate memory cert
    ) external pure returns (bytes memory);

    /**
     * @notice Gets the current fee split percentage
     * @return The fee split in basis points
     */
    function feeSplit() external view returns (uint16);

    /**
     * @notice Gets the current fee split collector address
     * @return The address that receives fee splits
     */
    function feeSplitCollector() external view returns (address);

    /**
     * @notice Gets the BN254 certificate verifier address
     * @return The address of the BN254 certificate verifier
     */
    function BN254_CERTIFICATE_VERIFIER() external view returns (address);

    /**
     * @notice Gets the ECDSA certificate verifier address
     * @return The address of the ECDSA certificate verifier
     */
    function ECDSA_CERTIFICATE_VERIFIER() external view returns (address);
}
