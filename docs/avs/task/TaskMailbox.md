# TaskMailbox

| File | Type | Proxy |
| -------- | -------- | -------- |
| [`TaskMailbox.sol`](../../../src/contracts/avs/task/TaskMailbox.sol) | Singleton | Transparent proxy |
| [`TaskMailboxStorage.sol`](../../../src/contracts/avs/task/TaskMailboxStorage.sol) | Storage | - |
| [`ITaskMailbox.sol`](../../../src/contracts/interfaces/ITaskMailbox.sol) | Interface | - |

Libraries and Mixins:

| File | Notes |
| -------- | -------- |
| [`SemVerMixin.sol`](../../../src/contracts/mixins/SemVerMixin.sol) | semantic versioning |
| [`OwnableUpgradeable`](https://github.com/OpenZeppelin/openzeppelin-contracts-upgradeable/blob/v4.9.0/contracts/access/OwnableUpgradeable.sol) | ownership management |
| [`ReentrancyGuardUpgradeable`](https://github.com/OpenZeppelin/openzeppelin-contracts-upgradeable/blob/v4.9.0/contracts/security/ReentrancyGuardUpgradeable.sol) | reentrancy protection |
| [`Initializable`](https://github.com/OpenZeppelin/openzeppelin-contracts-upgradeable/blob/v4.9.0/contracts/proxy/utils/Initializable.sol) | upgradeable initialization |
| [`SafeERC20`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/token/ERC20/utils/SafeERC20.sol) | safe token transfers |
| [`SafeCast`](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/utils/math/SafeCast.sol) | safe type casting |
| [`OperatorSetLib.sol`](../../../src/contracts/libraries/OperatorSetLib.sol) | operator set encoding/decoding |

## Prior Reading

* [Hourglass Framework](https://github.com/Layr-Labs/hourglass-monorepo/blob/master/README.md) - for understanding the task-based AVS execution model
* [AllocationManager](../../core/AllocationManager.md) - for understanding operator sets
* [KeyRegistrar](../../permissions/KeyRegistrar.md) - for understanding operator key management
* [CertificateVerifier](../../multichain/destination/CertificateVerifier.md) - for understanding certificate verification

## Overview
The `TaskMailbox` is a core infrastructure contract that enables task-based AVS (Autonomous Verifiable Service) execution models. It provides a standardized way for AVSs to accept tasks created by users or apps, have operators execute them, and submit verified results on-chain. The contract acts as a mailbox system where task creators post tasks with fees, and operators compete to execute and submit results with proper consensus verification.

The `TaskMailbox`'s responsibilities are broken down into the following concepts:

* [Task Creation and Lifecycle](#task-creation-and-lifecycle)
* [Executor Operator Sets](#executor-operator-sets)
* [Result Submission and Verification](#result-submission-and-verification)
* [Fee Management](#fee-management)
* [Task Hooks and AVS Integration](#task-hooks-and-avs-integration)

## Parameterization

The `TaskMailbox` uses the following key parameters:

* **Fee Split**: Configurable percentage (0-10000 basis points) that determines how task fees are split between the protocol fee collector and the AVS fee collector
* **Task SLA**: Service Level Agreement duration (in seconds) set per executor operator set, defining how long operators have to complete a task
* **Consensus Thresholds**: Configurable per operator set, defining the required stake proportion for result verification

---

## Task Creation and Lifecycle

Tasks in the TaskMailbox system follow a well-defined lifecycle with specific states and transitions. Each task is uniquely identified by a hash computed from the task parameters, block context, and a global counter.

**State Variables:**
* `_globalTaskCount`: Counter ensuring unique task hashes
* `_tasks`: Mapping from task hash to task details

**Methods:**
* [`createTask`](#createtask)
* [`getTaskInfo`](#gettaskinfo)
* [`getTaskStatus`](#gettaskstatus)

### Task Status Flow

Tasks can be in one of the following states:
1. **NONE**: Task does not exist
2. **CREATED**: Task has been created and is waiting for execution
3. **EXPIRED**: Task SLA has passed without result submission
4. **VERIFIED**: Task result has been submitted and verified

#### `createTask`

```solidity
/**
 * @notice Creates a new task for execution by operators
 * @param taskParams The parameters for the task including refund collector, executor operator set, and payload
 * @return taskHash The unique identifier for the created task
 */
function createTask(TaskParams memory taskParams) external nonReentrant returns (bytes32 taskHash)
```

Creates a new task in the system. The method performs several validations and operations:

1. Validates that the executor operator set is registered and has a valid configuration
2. Calls the AVS task hook for pre-creation validation
3. Calculates the task fee using the AVS task hook
4. Generates a unique task hash
5. Transfers the fee from the caller to the TaskMailbox
6. Stores the task with its current configuration snapshot
7. Calls the AVS task hook for post-creation handling

*Effects*:
* Increments `_globalTaskCount`
* Stores task in `_tasks` mapping
* Transfers fee tokens from caller
* Emits `TaskCreated` event
* Calls [`IAVSTaskHook.validatePreTaskCreation`](../../../src/contracts/interfaces/IAVSTaskHook.sol) and [`IAVSTaskHook.handlePostTaskCreation`](../../../src/contracts/interfaces/IAVSTaskHook.sol)

*Requirements*:
* Executor operator set must be registered
* Executor operator set must have valid task configuration
* Task payload must not be empty
* Fee transfer must succeed
* AVS validation must pass

#### `getTaskInfo`

```solidity
/**
 * @notice Retrieves the complete information for a task
 * @param taskHash The unique identifier of the task
 * @return Task struct containing all task details
 */
function getTaskInfo(bytes32 taskHash) external view returns (Task memory)
```

Returns the complete task information including its current status. The status is computed dynamically based on the current block timestamp and task state.

#### `getTaskStatus`

```solidity
/**
 * @notice Gets the current status of a task
 * @param taskHash The unique identifier of the task
 * @return TaskStatus enum value
 */
function getTaskStatus(bytes32 taskHash) external view returns (TaskStatus)
```

Returns only the current status of a task, useful for lightweight status checks.

---

## Executor Operator Sets

Executor operator sets define which operators are eligible to execute tasks and under what conditions. Each operator set must be configured before it can be used for task execution.

**State Variables:**
* `_executorOperatorSetTaskConfigs`: Mapping from operator set to its task configuration
* `_isExecutorOperatorSetRegistered`: Tracks registered operator sets

**Methods:**
* [`setExecutorOperatorSetTaskConfig`](#setexecutoroperatorsettaskconfig)
* [`registerExecutorOperatorSet`](#registerexecutoroperatorset)
* [`getExecutorOperatorSetTaskConfig`](#getexecutoroperatorsettaskconfig)

### ExecutorOperatorSetTaskConfig Structure

The task configuration for an executor operator set contains the following fields:

```solidity
struct ExecutorOperatorSetTaskConfig {
    IAVSTaskHook taskHook;        // AVS-specific contract for custom validation
    uint96 taskSLA;               // Time limit in seconds for task completion
    IERC20 feeToken;              // Token used for task fees (zero address = no fees)
    address feeCollector;         // Address to receive AVS portion of fees
    CurveType curveType;          // Cryptographic curve (BN254 or ECDSA)
    Consensus consensus;          // Consensus type and parameters
    bytes taskMetadata;           // AVS-specific metadata
}
```

### Consensus Configuration

The consensus configuration determines how operator signatures are validated:

```solidity
struct Consensus {
    ConsensusType consensusType;  // Type of consensus validation
    bytes value;                  // Type-specific parameters
}

enum ConsensusType {
    NONE,                         // AVS handles consensus validation
    STAKE_PROPORTION_THRESHOLD    // Require minimum stake percentage
}
```

**Consensus Types:**

1. **NONE**: 
   - The TaskMailbox only verifies certificate validity (signature, timestamp, message hash)
   - AVS is responsible for implementing custom consensus logic in task hooks
   - `value` must be empty bytes
   - Useful for AVSs with custom consensus mechanisms or off-chain validation

2. **STAKE_PROPORTION_THRESHOLD**:
   - Requires a minimum percentage of total stake to sign the result
   - `value` contains `abi.encode(uint16)` representing threshold in basis points (0-10000)
   - Example: 6667 = 66.67% stake required
   - Certificate verification will fail if threshold is not met

#### `setExecutorOperatorSetTaskConfig`

```solidity
/**
 * @notice Sets the task configuration for an executor operator set
 * @param executorOperatorSet The operator set to configure
 * @param config The configuration including task hook, SLA, fee token, etc.
 */
function setExecutorOperatorSetTaskConfig(
    OperatorSet memory executorOperatorSet,
    ExecutorOperatorSetTaskConfig memory config
) external
```

Configures how tasks should be executed by a specific operator set. The configuration includes:
- **Task Hook**: AVS-specific contract for custom validation and handling
- **Task SLA**: Time limit for task completion
- **Fee Token**: Token used for task fees (can be zero address for no fees). **Fees will not be collected if this is the zero address.**
- **Fee Collector**: Address to receive AVS portion of fees
- **Curve Type**: Cryptographic curve used by operators (BN254 or ECDSA)
- **Consensus**: Type and threshold for result verification
- **Task Metadata**: AVS-specific metadata

*Effects*:
* Stores configuration in `_executorOperatorSetTaskConfigs`
* Sets `_isExecutorOperatorSetRegistered` to true if not already registered
* Emits `ExecutorOperatorSetTaskConfigSet` event

*Requirements*:
* Caller must be the operator set owner (verified via certificate verifier)
* Task hook must not be zero address
* Task SLA must be greater than zero
* Consensus type and curve type must be valid
* Consensus value must be properly formatted

**Example Configuration:**

```solidity
// Example 1: Configure with 66.67% stake threshold
ExecutorOperatorSetTaskConfig memory config = ExecutorOperatorSetTaskConfig({
    taskHook: IAVSTaskHook(0x...),
    taskSLA: 3600, // 1 hour
    feeToken: IERC20(0x...),
    feeCollector: 0x...,
    curveType: CurveType.BN254,
    consensus: Consensus({
        consensusType: ConsensusType.STAKE_PROPORTION_THRESHOLD,
        value: abi.encode(uint16(6667)) // 66.67%
    }),
    taskMetadata: bytes("")
});

// Example 2: Configure with NONE consensus (AVS handles validation)
ExecutorOperatorSetTaskConfig memory config = ExecutorOperatorSetTaskConfig({
    taskHook: IAVSTaskHook(0x...),
    taskSLA: 3600,
    feeToken: IERC20(address(0)), // No fees
    feeCollector: address(0),
    curveType: CurveType.ECDSA,
    consensus: Consensus({
        consensusType: ConsensusType.NONE,
        value: bytes("") // Must be empty
    }),
    taskMetadata: bytes("")
});
```

#### `registerExecutorOperatorSet`

```solidity
/**
 * @notice Registers or unregisters an executor operator set
 * @param executorOperatorSet The operator set to register/unregister
 * @param isRegistered Whether to register (true) or unregister (false)
 */
function registerExecutorOperatorSet(
    OperatorSet memory executorOperatorSet,
    bool isRegistered
) external
```

Allows operator set owners to explicitly register or unregister their operator sets for task execution.

*Effects*:
* Updates `_isExecutorOperatorSetRegistered`
* Emits `ExecutorOperatorSetRegistered` event

*Requirements*:
* Caller must be the operator set owner
* Operator set must have a valid configuration if registering

---

## Result Submission and Verification

Task results are submitted along with cryptographic certificates that prove consensus among the operator set. The verification process depends on the curve type configured for the operator set.

**Methods:**
* [`submitResult`](#submitresult)
* [`getTaskResult`](#gettaskresult)

#### `submitResult`

```solidity
/**
 * @notice Submits the result of a task execution with consensus proof
 * @param taskHash The unique identifier of the task
 * @param executorCert Certificate proving operator consensus
 * @param result The execution result data
 */
function submitResult(
    bytes32 taskHash,
    bytes memory executorCert,
    bytes memory result
) external nonReentrant
```

Submits the result of task execution along with proof of consensus. The method:

1. Validates the task exists and hasn't expired or been verified
2. Calls AVS hook for pre-submission validation
3. Verifies the certificate based on both curve type and consensus type:
   - **ConsensusType.NONE**: 
     - Validates certificate fields (reference timestamp, message hash, non-empty signature)
     - Calls certificate verifier to validate signature authenticity
     - Does NOT enforce any stake threshold requirements
     - AVS can implement custom consensus logic in `handlePostTaskResultSubmission`
   - **ConsensusType.STAKE_PROPORTION_THRESHOLD**:
     - Performs all validations from NONE
     - Additionally verifies that signers meet the configured stake threshold
     - Uses `verifyCertificateProportion` to ensure minimum stake percentage signed
4. Distributes fees according to fee split configuration
5. Stores the result and marks task as verified
6. Calls AVS hook for post-submission handling

*Effects*:
* Updates task status to VERIFIED
* Stores executor certificate and result
* Transfers fees to collectors
* Emits `TaskVerified` event
* Calls [`IAVSTaskHook.validatePreTaskResultSubmission`](../../../src/contracts/interfaces/IAVSTaskHook.sol) and [`IAVSTaskHook.handlePostTaskResultSubmission`](../../../src/contracts/interfaces/IAVSTaskHook.sol)

*Requirements*:
* Task must exist and be in CREATED status
* Current timestamp must be after task creation time
* Certificate must have valid signature(s)
* Certificate verification must pass consensus threshold (reverts with `ThresholdNotMet` if not)
* AVS validation must pass

#### `getTaskResult`

```solidity
/**
 * @notice Retrieves the result of a verified task
 * @param taskHash The unique identifier of the task
 * @return result The task execution result data
 */
function getTaskResult(bytes32 taskHash) external view returns (bytes memory result)
```

Returns the result data for tasks that have been successfully verified.

*Requirements*:
* Task must be in VERIFIED status

---

## Fee Management

The TaskMailbox implements a flexible fee system that supports fee splitting between the protocol and AVS, as well as refunds for expired tasks.

**State Variables:**
* `feeSplit`: Global fee split percentage (basis points)
* `feeSplitCollector`: Address receiving protocol portion of fees

**Methods:**
* [`setFeeSplit`](#setfeesplit)
* [`setFeeSplitCollector`](#setfeesplitcollector)
* [`refundFee`](#refundfee)

#### `setFeeSplit`

```solidity
/**
 * @notice Sets the global fee split percentage
 * @param _feeSplit The fee split in basis points (0-10000)
 */
function setFeeSplit(uint16 _feeSplit) external onlyOwner
```

Configures what percentage of task fees goes to the protocol fee collector versus the AVS fee collector.

*Effects*:
* Updates `feeSplit` state variable
* Emits `FeeSplitSet` event

*Requirements*:
* Caller must be contract owner
* Fee split must not exceed 10000 (100%)

#### `setFeeSplitCollector`

```solidity
/**
 * @notice Sets the protocol fee collector address
 * @param _feeSplitCollector The address to receive protocol fees
 */
function setFeeSplitCollector(address _feeSplitCollector) external onlyOwner
```

Sets the address that receives the protocol portion of task fees.

*Effects*:
* Updates `feeSplitCollector` state variable
* Emits `FeeSplitCollectorSet` event

*Requirements*:
* Caller must be contract owner
* Address must not be zero

#### `refundFee`

```solidity
/**
 * @notice Refunds the fee for an expired task
 * @param taskHash The unique identifier of the task
 */
function refundFee(bytes32 taskHash) external nonReentrant
```

Allows the designated refund collector to reclaim fees from expired tasks that were never executed.

*Effects*:
* Marks task fee as refunded
* Transfers full fee amount to refund collector
* Emits `FeeRefunded` event

*Requirements*:
* Task must be in EXPIRED status
* Caller must be the task's refund collector
* Fee must not have been previously refunded
* Task must have a fee token configured

---

## Task Hooks and AVS Integration

Task hooks provide AVSs with customization points throughout the task lifecycle. This allows AVSs to implement custom validation logic, fee calculations, and side effects.

**Integration Points:**

### IAVSTaskHook Interface

AVSs implement the [`IAVSTaskHook`](../../../src/contracts/interfaces/IAVSTaskHook.sol) interface to customize task behavior:

```solidity
interface IAVSTaskHook {
    // Called before task creation to validate the caller and parameters
    function validatePreTaskCreation(
        address caller,
        TaskParams memory taskParams
    ) external view;

    // Called after task creation for any side effects
    function handlePostTaskCreation(bytes32 taskHash) external;

    // Called before result submission to validate the submitter and data
    function validatePreTaskResultSubmission(
        address caller,
        bytes32 taskHash,
        bytes memory cert,
        bytes memory result
    ) external view;

    // Called after successful result submission for any side effects
    function handlePostTaskResultSubmission(
        address caller,
        bytes32 taskHash
    ) external;

    // Calculates the fee for a task based on its parameters
    function calculateTaskFee(
        TaskParams memory taskParams
    ) external view returns (uint96);
}
```

### Integration Flow

1. **Task Creation**:
   - AVS validates caller permissions and task parameters
   - AVS calculates appropriate fee based on task complexity
   - AVS performs any initialization after task creation

2. **Result Submission**:
   - AVS validates submitter permissions
   - AVS can validate result format and content
   - AVS processes verified results for their application logic

This hook system enables AVSs to build sophisticated task-based systems while leveraging the TaskMailbox's core infrastructure for consensus verification and fee management.

---

## Security Considerations

The TaskMailbox implements several security measures:

1. **Reentrancy Protection**: All state-changing functions use OpenZeppelin's `ReentrancyGuardUpgradeable`
2. **Certificate Verification**: Results must include valid consensus proofs verified by the appropriate certificate verifier
3. **Fee Safety**: Uses OpenZeppelin's `SafeERC20` for all token transfers
4. **Access Control**: Operator set owners control their configurations; contract owner controls global parameters
5. **Timestamp Validation**: Tasks cannot be verified at their creation timestamp to prevent same-block manipulation