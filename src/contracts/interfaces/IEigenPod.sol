// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "../libraries/BeaconChainProofs.sol";
import "./ISemVerMixin.sol";
import "./IEigenPodManager.sol";

interface IEigenPodErrors {
    /// @dev Thrown when msg.sender is not the EPM.
    error OnlyEigenPodManager();
    /// @dev Thrown when msg.sender is not the pod owner.
    error OnlyEigenPodOwner();
    /// @dev Thrown when msg.sender is not owner or the proof submitter.
    error OnlyEigenPodOwnerOrProofSubmitter();
    /// @dev Thrown when attempting an action that is currently paused.
    error CurrentlyPaused();

    /// Invalid Inputs

    /// @dev Thrown when an address of zero is provided.
    error InputAddressZero();
    /// @dev Thrown when two array parameters have mismatching lengths.
    error InputArrayLengthMismatch();
    /// @dev Thrown when `validatorPubKey` length is not equal to 48-bytes.
    error InvalidPubKeyLength();
    /// @dev Thrown when provided timestamp is out of range.
    error TimestampOutOfRange();

    /// Checkpoints

    /// @dev Thrown when no active checkpoints are found.
    error NoActiveCheckpoint();
    /// @dev Thrown if an uncompleted checkpoint exists.
    error CheckpointAlreadyActive();
    /// @dev Thrown if there's not a balance available to checkpoint.
    error NoBalanceToCheckpoint();
    /// @dev Thrown when attempting to create a checkpoint twice within a given block.
    error CannotCheckpointTwiceInSingleBlock();

    /// Withdrawing

    /// @dev Thrown when amount exceeds `restakedExecutionLayerGwei`.
    error InsufficientWithdrawableBalance();

    /// Validator Status

    /// @dev Thrown when a validator's withdrawal credentials have already been verified.
    error CredentialsAlreadyVerified();
    /// @dev Thrown if the provided proof is not valid for this EigenPod.
    error WithdrawalCredentialsNotForEigenPod();
    /// @dev Thrown when a validator is not in the ACTIVE status in the pod.
    error ValidatorNotActiveInPod();
    /// @dev Thrown when validator is not active yet on the beacon chain.
    error ValidatorInactiveOnBeaconChain();
    /// @dev Thrown if a validator is exiting the beacon chain.
    error ValidatorIsExitingBeaconChain();
    /// @dev Thrown when a validator has not been slashed on the beacon chain.
    error ValidatorNotSlashedOnBeaconChain();

    /// Consolidation and Withdrawal Requests

    /// @dev Thrown when a predeploy request is initiated with insufficient msg.value
    error InsufficientFunds();
    /// @dev Thrown when calling the predeploy fails
    error PredeployFailed();
    /// @dev Thrown when querying a predeploy for its current fee fails
    error FeeQueryFailed();

    /// Misc

    /// @dev Thrown when an invalid block root is returned by the EIP-4788 oracle.
    error InvalidEIP4788Response();
    /// @dev Thrown when attempting to send an invalid amount to the beacon deposit contract.
    error MsgValueNot32ETH();
    /// @dev Thrown when provided `beaconTimestamp` is too far in the past.
    error BeaconTimestampTooFarInPast();
    /// @dev Thrown when provided `beaconTimestamp` is before the last checkpoint
    error BeaconTimestampBeforeLatestCheckpoint();
    /// @dev Thrown when the pectraForkTimestamp returned from the EigenPodManager is zero
    error ForkTimestampZero();
}

interface IEigenPodTypes {
    enum VALIDATOR_STATUS {
        INACTIVE, // doesnt exist
        ACTIVE, // staked on ethpos and withdrawal credentials are pointed to the EigenPod
        WITHDRAWN // withdrawn from the Beacon Chain

    }

    /**
     * @param validatorIndex index of the validator on the beacon chain
     * @param restakedBalanceGwei amount of beacon chain ETH restaked on EigenLayer in gwei
     * @param lastCheckpointedAt timestamp of the validator's most recent balance update
     * @param status last recorded status of the validator
     */
    struct ValidatorInfo {
        uint64 validatorIndex;
        uint64 restakedBalanceGwei;
        uint64 lastCheckpointedAt;
        VALIDATOR_STATUS status;
    }

    struct Checkpoint {
        bytes32 beaconBlockRoot;
        uint24 proofsRemaining;
        uint64 podBalanceGwei;
        int64 balanceDeltasGwei;
        uint64 prevBeaconBalanceGwei;
    }

    /**
     * @param srcPubkey the pubkey of the source validator for the consolidation
     * @param targetPubkey the pubkey of the target validator for the consolidation
     * @dev Note that if srcPubkey == targetPubkey, this is a "switch request," and will
     * change the validator's withdrawal credential type from 0x01 to 0x02.
     * For more notes on usage, see `requestConsolidation`
     */
    struct ConsolidationRequest {
        bytes srcPubkey;
        bytes targetPubkey;
    }

    /**
     * @param pubkey the pubkey of the validator to withdraw from
     * @param amountGwei the amount (in gwei) to withdraw from the beacon chain to the pod
     * @dev Note that if amountGwei == 0, this is a "full exit request," and will fully exit
     * the validator to the pod.
     * For more notes on usage, see `requestWithdrawal`
     */
    struct WithdrawalRequest {
        bytes pubkey;
        uint64 amountGwei;
    }
}

interface IEigenPodEvents is IEigenPodTypes {
    /// @notice Emitted when an ETH validator stakes via this eigenPod
    event EigenPodStaked(bytes32 pubkeyHash);

    /// @notice Emitted when a pod owner updates the proof submitter address
    event ProofSubmitterUpdated(address prevProofSubmitter, address newProofSubmitter);

    /// @notice Emitted when an ETH validator's withdrawal credentials are successfully verified to be pointed to this eigenPod
    event ValidatorRestaked(bytes32 pubkeyHash);

    /// @notice Emitted when an ETH validator's  balance is proven to be updated.  Here newValidatorBalanceGwei
    //  is the validator's balance that is credited on EigenLayer.
    event ValidatorBalanceUpdated(bytes32 pubkeyHash, uint64 balanceTimestamp, uint64 newValidatorBalanceGwei);

    /// @notice Emitted when restaked beacon chain ETH is withdrawn from the eigenPod.
    event RestakedBeaconChainETHWithdrawn(address indexed recipient, uint256 amount);

    /// @notice Emitted when ETH is received via the `receive` fallback
    event NonBeaconChainETHReceived(uint256 amountReceived);

    /// @notice Emitted when a checkpoint is created
    event CheckpointCreated(
        uint64 indexed checkpointTimestamp, bytes32 indexed beaconBlockRoot, uint256 validatorCount
    );

    /// @notice Emitted when a checkpoint is finalized
    event CheckpointFinalized(uint64 indexed checkpointTimestamp, int256 totalShareDeltaWei);

    /// @notice Emitted when a validator is proven for a given checkpoint
    event ValidatorCheckpointed(uint64 indexed checkpointTimestamp, bytes32 indexed pubkeyHash);

    /// @notice Emitted when a validator is proven to have 0 balance at a given checkpoint
    event ValidatorWithdrawn(uint64 indexed checkpointTimestamp, bytes32 indexed pubkeyHash);

    /// @notice Emitted when a consolidation request is initiated where source == target
    event SwitchToCompoundingRequested(bytes32 indexed validatorPubkeyHash);

    /// @notice Emitted when a standard consolidation request is initiated
    event ConsolidationRequested(bytes32 indexed sourcePubkeyHash, bytes32 indexed targetPubkeyHash);

    /// @notice Emitted when a withdrawal request is initiated where request.amountGwei == 0
    event ExitRequested(bytes32 indexed validatorPubkeyHash);

    /// @notice Emitted when a partial withdrawal request is initiated
    event WithdrawalRequested(bytes32 indexed validatorPubkeyHash, uint64 withdrawalAmountGwei);
}

/**
 * @title The implementation contract used for restaking beacon chain ETH on EigenLayer
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @dev Note that all beacon chain balances are stored as gwei within the beacon chain datastructures. We choose
 *   to account balances in terms of gwei in the EigenPod contract and convert to wei when making calls to other contracts
 */
interface IEigenPod is IEigenPodErrors, IEigenPodEvents, ISemVerMixin {
    /// @notice Used to initialize the pointers to contracts crucial to the pod's functionality, in beacon proxy construction from EigenPodManager
    function initialize(
        address owner
    ) external;

    /// @notice Called by EigenPodManager when the owner wants to create another ETH validator.
    /// @dev This function only supports staking to a 0x01 validator. For compounding validators, please interact directly with the deposit contract.
    function stake(bytes calldata pubkey, bytes calldata signature, bytes32 depositDataRoot) external payable;

    /**
     * @notice Transfers `amountWei` from this contract to the `recipient`. Only callable by the EigenPodManager as part
     * of the DelegationManager's withdrawal flow.
     * @dev `amountWei` is not required to be a whole Gwei amount. Amounts less than a Gwei multiple may be unrecoverable due to Gwei conversion.
     */
    function withdrawRestakedBeaconChainETH(address recipient, uint256 amount) external;

    /**
     * @dev Create a checkpoint used to prove this pod's active validator set. Checkpoints are completed
     * by submitting one checkpoint proof per ACTIVE validator. During the checkpoint process, the total
     * change in ACTIVE validator balance is tracked, and any validators with 0 balance are marked `WITHDRAWN`.
     * @dev Once finalized, the pod owner is awarded shares corresponding to:
     * - the total change in their ACTIVE validator balances
     * - any ETH in the pod not already awarded shares
     * @dev A checkpoint cannot be created if the pod already has an outstanding checkpoint. If
     * this is the case, the pod owner MUST complete the existing checkpoint before starting a new one.
     * @param revertIfNoBalance Forces a revert if the pod ETH balance is 0. This allows the pod owner
     * to prevent accidentally starting a checkpoint that will not increase their shares
     */
    function startCheckpoint(
        bool revertIfNoBalance
    ) external;

    /**
     * @dev Progress the current checkpoint towards completion by submitting one or more validator
     * checkpoint proofs. Anyone can call this method to submit proofs towards the current checkpoint.
     * For each validator proven, the current checkpoint's `proofsRemaining` decreases.
     * @dev If the checkpoint's `proofsRemaining` reaches 0, the checkpoint is finalized.
     * (see `_updateCheckpoint` for more details)
     * @dev This method can only be called when there is a currently-active checkpoint.
     * @param balanceContainerProof proves the beacon's current balance container root against a checkpoint's `beaconBlockRoot`
     * @param proofs Proofs for one or more validator current balances against the `balanceContainerRoot`
     */
    function verifyCheckpointProofs(
        BeaconChainProofs.BalanceContainerProof calldata balanceContainerProof,
        BeaconChainProofs.BalanceProof[] calldata proofs
    ) external;

    /**
     * @dev Verify one or more validators have their withdrawal credentials pointed at this EigenPod, and award
     * shares based on their effective balance. Proven validators are marked `ACTIVE` within the EigenPod, and
     * future checkpoint proofs will need to include them.
     * @dev Withdrawal credential proofs MUST NOT be older than `currentCheckpointTimestamp`.
     * @dev Validators proven via this method MUST NOT have an exit epoch set already.
     * @param beaconTimestamp the beacon chain timestamp sent to the 4788 oracle contract. Corresponds
     * to the parent beacon block root against which the proof is verified.
     * @param stateRootProof proves a beacon state root against a beacon block root
     * @param validatorIndices a list of validator indices being proven
     * @param validatorFieldsProofs proofs of each validator's `validatorFields` against the beacon state root
     * @param validatorFields the fields of the beacon chain "Validator" container. See consensus specs for
     * details: https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator
     */
    function verifyWithdrawalCredentials(
        uint64 beaconTimestamp,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        uint40[] calldata validatorIndices,
        bytes[] calldata validatorFieldsProofs,
        bytes32[][] calldata validatorFields
    ) external;

    /**
     * @dev Prove that one of this pod's active validators was slashed on the beacon chain. A successful
     * staleness proof allows the caller to start a checkpoint.
     *
     * @dev Note that in order to start a checkpoint, any existing checkpoint must already be completed!
     * (See `_startCheckpoint` for details)
     *
     * @dev Note that this method allows anyone to start a checkpoint as soon as a slashing occurs on the beacon
     * chain. This is intended to make it easier to external watchers to keep a pod's balance up to date.
     *
     * @dev Note too that beacon chain slashings are not instant. There is a delay between the initial slashing event
     * and the validator's final exit back to the execution layer. During this time, the validator's balance may or
     * may not drop further due to a correlation penalty. This method allows proof of a slashed validator
     * to initiate a checkpoint for as long as the validator remains on the beacon chain. Once the validator
     * has exited and been checkpointed at 0 balance, they are no longer "checkpoint-able" and cannot be proven
     * "stale" via this method.
     * See https://eth2book.info/capella/part3/transition/epoch/#slashings for more info.
     *
     * @param beaconTimestamp the beacon chain timestamp sent to the 4788 oracle contract. Corresponds
     * to the parent beacon block root against which the proof is verified.
     * @param stateRootProof proves a beacon state root against a beacon block root
     * @param proof the fields of the beacon chain "Validator" container, along with a merkle proof against
     * the beacon state root. See the consensus specs for more details:
     * https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator
     *
     * @dev Staleness conditions:
     * - Validator's last checkpoint is older than `beaconTimestamp`
     * - Validator MUST be in `ACTIVE` status in the pod
     * - Validator MUST be slashed on the beacon chain
     */
    function verifyStaleBalance(
        uint64 beaconTimestamp,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        BeaconChainProofs.ValidatorProof calldata proof
    ) external;

    /// @notice Allows the owner or proof submitter to initiate one or more requests to
    /// consolidate their validators on the beacon chain.
    /// @param requests An array of requests consisting of the source and target pubkeys
    /// of the validators to be consolidated
    /// @dev The target validator MUST have ACTIVE (proven) withdrawal credentials pointed at
    /// the pod. This prevents cross-pod consolidations.
    /// @dev The consolidation request predeploy requires a fee is sent with each request;
    /// this is pulled from msg.value. After submitting all requests, any remaining fee is
    /// refunded to the caller by calling its fallback function.
    /// @dev This contract exposes `getConsolidationRequestFee` to query the current fee for
    /// a single request. If submitting multiple requests in a single block, the total fee
    /// is equal to (fee * requests.length). This fee is updated at the end of each block.
    ///
    /// (See https://eips.ethereum.org/EIPS/eip-7251#fee-calculation for details)
    ///
    /// @dev Note on beacon chain behavior:
    /// - If request.srcPubkey == request.targetPubkey, this is a "switch" consolidation. Once
    ///   processed on the beacon chain, the validator's withdrawal credentials will be changed
    ///   to compounding (0x02).
    /// - The rest of the notes assume src != target.
    /// - The target validator MUST already have 0x02 credentials. The source validator can have either.
    /// - Consolidation sets the source validator's exit_epoch and withdrawable_epoch, similar to an exit.
    ///   When the exit epoch is reached, an epoch sweep will process the consolidation and transfer balance
    ///   from the source to the target validator.
    /// - Consolidation transfers min(srcValidator.effective_balance, state.balance[srcIndex]) to the target.
    ///   This may not be the entirety of the source validator's balance; any remainder will be moved to the
    ///   pod when hit by a subsequent withdrawal sweep.
    ///
    /// @dev Note that consolidation requests CAN FAIL for a variety of reasons. Failures occur when the request
    /// is processed on the beacon chain, and are invisible to the pod. The pod and predeploy cannot guarantee
    /// a request will succeed; it's up to the pod owner to determine this for themselves. If your request fails,
    /// you can retry by initiating another request via this method.
    ///
    /// Some requirements that are NOT checked by the pod:
    /// - If request.srcPubkey == request.targetPubkey, the validator MUST have 0x01 credentials
    /// - If request.srcPubkey != request.targetPubkey, the target validator MUST have 0x02 credentials
    /// - Both the source and target validators MUST be active on the beacon chain and MUST NOT have
    ///   initiated exits
    /// - The source validator MUST NOT have pending partial withdrawal requests (via `requestWithdrawal`)
    /// - If the source validator is slashed after requesting consolidation (but before processing),
    ///   the consolidation will be skipped.
    ///
    /// For further reference, see consolidation processing at block and epoch boundaries:
    /// - Block: https://github.com/ethereum/consensus-specs/blob/dev/specs/electra/beacon-chain.md#new-process_consolidation_request
    /// - Epoch: https://github.com/ethereum/consensus-specs/blob/dev/specs/electra/beacon-chain.md#new-process_pending_consolidations
    function requestConsolidation(
        ConsolidationRequest[] calldata requests
    ) external payable;

    /// @notice Allows the owner or proof submitter to initiate one or more requests to
    /// withdraw funds from validators on the beacon chain.
    /// @param requests An array of requests consisting of the source validator and an
    /// amount to withdraw
    /// @dev The withdrawal request predeploy requires a fee is sent with each request;
    /// this is pulled from msg.value. After submitting all requests, any remaining fee is
    /// refunded to the caller by calling its fallback function.
    /// @dev This contract exposes `getWithdrawalRequestFee` to query the current fee for
    /// a single request. If submitting multiple requests in a single block, the total fee
    /// is equal to (fee * requests.length). This fee is updated at the end of each block.
    ///
    /// (See https://eips.ethereum.org/EIPS/eip-7002#fee-update-rule for details)
    ///
    /// @dev Note on beacon chain behavior:
    /// - Withdrawal requests have two types: full exit requests, and partial exit requests.
    ///   Partial exit requests will be skipped if the validator has 0x01 withdrawal credentials.
    ///   If you want your validators to have access to partial exits, use `requestConsolidation`
    ///   to change their withdrawal credentials to compounding (0x02).
    /// - If request.amount == 0, this is a FULL exit request. A full exit request initiates a
    ///   standard validator exit.
    /// - Other amounts are treated as PARTIAL exit requests. A partial exit request will NOT result
    ///   in a validator with less than 32 ETH balance. Any requested amount above this is ignored.
    /// - The actual amount withdrawn for a partial exit is given by the formula:
    ///   min(request.amount, state.balances[vIdx] - 32 ETH - pending_balance_to_withdraw)
    ///   (where `pending_balance_to_withdraw` is the sum of any outstanding partial exit requests)
    ///   (Note that this means you may request more than is actually withdrawn!)
    ///
    /// @dev Note that withdrawal requests CAN FAIL for a variety of reasons. Failures occur when the request
    /// is processed on the beacon chain, and are invisible to the pod. The pod and predeploy cannot guarantee
    /// a request will succeed; it's up to the pod owner to determine this for themselves. If your request fails,
    /// you can retry by initiating another request via this method.
    ///
    /// Some requirements that are NOT checked by the pod:
    /// - request.pubkey MUST be a valid validator pubkey
    /// - request.pubkey MUST belong to a validator whose withdrawal credentials are this pod
    /// - If request.amount is for a partial exit, the validator MUST have 0x02 withdrawal credentials
    /// - If request.amount is for a full exit, the validator MUST NOT have any pending partial exits
    /// - The validator MUST be active and MUST NOT have initiated exit
    ///
    /// For further reference: https://github.com/ethereum/consensus-specs/blob/dev/specs/electra/beacon-chain.md#new-process_withdrawal_request
    function requestWithdrawal(
        WithdrawalRequest[] calldata requests
    ) external payable;

    /// @notice called by owner of a pod to remove any ERC20s deposited in the pod
    function recoverTokens(IERC20[] memory tokenList, uint256[] memory amountsToWithdraw, address recipient) external;

    /// @notice Allows the owner of a pod to update the proof submitter, a permissioned
    /// address that can call various EigenPod methods, but cannot trigger asset withdrawals
    /// from the DelegationManager.
    /// @dev Note that EITHER the podOwner OR proofSubmitter can access these methods,
    /// so it's fine to set your proofSubmitter to 0 if you want the podOwner to be the
    /// only address that can call these methods.
    /// @param newProofSubmitter The new proof submitter address. If set to 0, only the
    /// pod owner will be able to call EigenPod methods.
    function setProofSubmitter(
        address newProofSubmitter
    ) external;

    /**
     *
     *                                VIEW METHODS
     *
     */

    /// @notice An address with permissions to call `startCheckpoint` and `verifyWithdrawalCredentials`, set
    /// by the podOwner. This role exists to allow a podOwner to designate a hot wallet that can call
    /// these methods, allowing the podOwner to remain a cold wallet that is only used to manage funds.
    /// @dev If this address is NOT set, only the podOwner can call `startCheckpoint` and `verifyWithdrawalCredentials`
    function proofSubmitter() external view returns (address);

    /// @notice Native ETH in the pod that has been accounted for in a checkpoint (denominated in gwei).
    /// This amount is withdrawable from the pod via the DelegationManager withdrawal flow.
    function withdrawableRestakedExecutionLayerGwei() external view returns (uint64);

    /// @notice The single EigenPodManager for EigenLayer
    function eigenPodManager() external view returns (IEigenPodManager);

    /// @notice The owner of this EigenPod
    function podOwner() external view returns (address);

    /// @notice Returns the validatorInfo struct for the provided pubkeyHash
    function validatorPubkeyHashToInfo(
        bytes32 validatorPubkeyHash
    ) external view returns (ValidatorInfo memory);

    /// @notice Returns the validatorInfo struct for the provided pubkey
    function validatorPubkeyToInfo(
        bytes calldata validatorPubkey
    ) external view returns (ValidatorInfo memory);

    /// @notice Returns the validator status for a given validator pubkey hash
    function validatorStatus(
        bytes32 pubkeyHash
    ) external view returns (VALIDATOR_STATUS);

    /// @notice Returns the validator status for a given validator pubkey
    function validatorStatus(
        bytes calldata validatorPubkey
    ) external view returns (VALIDATOR_STATUS);

    /// @notice Number of validators with proven withdrawal credentials, who do not have proven full withdrawals
    function activeValidatorCount() external view returns (uint256);

    /// @notice The timestamp of the last checkpoint finalized
    function lastCheckpointTimestamp() external view returns (uint64);

    /// @notice The timestamp of the currently-active checkpoint. Will be 0 if there is not active checkpoint
    function currentCheckpointTimestamp() external view returns (uint64);

    /// @notice Returns the currently-active checkpoint
    /// If there's not an active checkpoint, this method returns the checkpoint that was last active.
    function currentCheckpoint() external view returns (Checkpoint memory);

    /// @notice For each checkpoint, the total balance attributed to exited validators, in gwei
    ///
    /// NOTE that the values added to this mapping are NOT guaranteed to capture the entirety of a validator's
    /// exit - rather, they capture the total change in a validator's balance when a checkpoint shows their
    /// balance change from nonzero to zero. While a change from nonzero to zero DOES guarantee that a validator
    /// has been fully exited, it is possible that the magnitude of this change does not capture what is
    /// typically thought of as a "full exit."
    ///
    /// For example:
    /// 1. Consider a validator was last checkpointed at 32 ETH before exiting. Once the exit has been processed,
    /// it is expected that the validator's exited balance is calculated to be `32 ETH`.
    /// 2. However, before `startCheckpoint` is called, a deposit is made to the validator for 1 ETH. The beacon
    /// chain will automatically withdraw this ETH, but not until the withdrawal sweep passes over the validator
    /// again. Until this occurs, the validator's current balance (used for checkpointing) is 1 ETH.
    /// 3. If `startCheckpoint` is called at this point, the balance delta calculated for this validator will be
    /// `-31 ETH`, and because the validator has a nonzero balance, it is not marked WITHDRAWN.
    /// 4. After the exit is processed by the beacon chain, a subsequent `startCheckpoint` and checkpoint proof
    /// will calculate a balance delta of `-1 ETH` and attribute a 1 ETH exit to the validator.
    ///
    /// If this edge case impacts your usecase, it should be possible to mitigate this by monitoring for deposits
    /// to your exited validators, and waiting to call `startCheckpoint` until those deposits have been automatically
    /// exited.
    ///
    /// Additional edge cases this mapping does not cover:
    /// - If a validator is slashed, their balance exited will reflect their original balance rather than the slashed amount
    /// - The final partial withdrawal for an exited validator will be likely be included in this mapping.
    ///   i.e. if a validator was last checkpointed at 32.1 ETH before exiting, the next checkpoint will calculate their
    ///   "exited" amount to be 32.1 ETH rather than 32 ETH.
    function checkpointBalanceExitedGwei(
        uint64
    ) external view returns (uint64);

    /// @notice Query the 4788 oracle to get the parent block root of the slot with the given `timestamp`
    /// @param timestamp of the block for which the parent block root will be returned. MUST correspond
    /// to an existing slot within the last 24 hours. If the slot at `timestamp` was skipped, this method
    /// will revert.
    function getParentBlockRoot(
        uint64 timestamp
    ) external view returns (bytes32);

    /// @notice Returns the fee required to add a consolidation request to the EIP-7251 predeploy this block.
    /// @dev Note that the predeploy updates its fee every block according to https://eips.ethereum.org/EIPS/eip-7251#fee-calculation
    /// Consider overestimating the amount sent to ensure the fee does not update before your transaction.
    function getConsolidationRequestFee() external view returns (uint256);

    /// @notice Returns the current fee required to add a withdrawal request to the EIP-7002 predeploy.
    /// @dev Note that the predeploy updates its fee every block according to https://eips.ethereum.org/EIPS/eip-7002#fee-update-rule
    /// Consider overestimating the amount sent to ensure the fee does not update before your transaction.
    function getWithdrawalRequestFee() external view returns (uint256);
}
