// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "../libraries/BeaconChainProofs.sol";
import "../libraries/BytesLib.sol";

import "../mixins/SemVerMixin.sol";

import "../interfaces/IETHPOSDeposit.sol";
import "../interfaces/IEigenPodManager.sol";
import "../interfaces/IPausable.sol";

import "./EigenPodPausingConstants.sol";
import "./EigenPodStorage.sol";

/**
 * @title The implementation contract used for restaking beacon chain ETH on EigenLayer
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This EigenPod Beacon Proxy implementation adheres to the current Deneb consensus specs
 * @dev Note that all beacon chain balances are stored as gwei within the beacon chain datastructures. We choose
 *   to account balances in terms of gwei in the EigenPod contract and convert to wei when making calls to other contracts
 */
contract EigenPod is
    Initializable,
    ReentrancyGuardUpgradeable,
    EigenPodPausingConstants,
    EigenPodStorage,
    SemVerMixin
{
    using BytesLib for bytes;
    using SafeERC20 for IERC20;
    using BeaconChainProofs for *;

    /**
     *
     *                            CONSTANTS / IMMUTABLES
     *
     */

    /// @notice The beacon chain stores balances in Gwei, rather than wei. This value is used to convert between the two
    uint256 internal constant GWEI_TO_WEI = 1e9;

    /// @notice The address of the EIP-4788 beacon block root oracle
    /// (See https://eips.ethereum.org/EIPS/eip-4788)
    address internal constant BEACON_ROOTS_ADDRESS = 0x000F3df6D732807Ef1319fB7B8bB8522d0Beac02;

    /// @notice The length of the EIP-4788 beacon block root ring buffer
    uint256 internal constant BEACON_ROOTS_HISTORY_BUFFER_LENGTH = 8191;

    /// @notice The beacon chain deposit contract
    IETHPOSDeposit public immutable ethPOS;

    /// @notice The single EigenPodManager for EigenLayer
    IEigenPodManager public immutable eigenPodManager;

    /// @notice This is the genesis time of the beacon state, to help us calculate conversions between slot and timestamp
    uint64 public immutable GENESIS_TIME;

    /**
     *
     *                                  MODIFIERS
     *
     */

    /// @notice Callable only by the EigenPodManager
    modifier onlyEigenPodManager() {
        require(msg.sender == address(eigenPodManager), OnlyEigenPodManager());
        _;
    }

    /// @notice Callable only by the pod's owner
    modifier onlyEigenPodOwner() {
        require(msg.sender == podOwner, OnlyEigenPodOwner());
        _;
    }

    /// @notice Callable only by the pod's owner or proof submitter
    modifier onlyOwnerOrProofSubmitter() {
        require(msg.sender == podOwner || msg.sender == proofSubmitter, OnlyEigenPodOwnerOrProofSubmitter());
        _;
    }

    /**
     * @notice Based on 'Pausable' code, but uses the storage of the EigenPodManager instead of this contract. This construction
     * is necessary for enabling pausing all EigenPods at the same time (due to EigenPods being Beacon Proxies).
     * Modifier throws if the `indexed`th bit of `_paused` in the EigenPodManager is 1, i.e. if the `index`th pause switch is flipped.
     */
    modifier onlyWhenNotPaused(
        uint8 index
    ) {
        require(!IPausable(address(eigenPodManager)).paused(index), CurrentlyPaused());
        _;
    }

    /**
     *
     *                               CONSTRUCTOR / INIT
     *
     */
    constructor(
        IETHPOSDeposit _ethPOS,
        IEigenPodManager _eigenPodManager,
        uint64 _GENESIS_TIME,
        string memory _version
    ) SemVerMixin(_version) {
        ethPOS = _ethPOS;
        eigenPodManager = _eigenPodManager;
        GENESIS_TIME = _GENESIS_TIME;
        _disableInitializers();
    }

    /// @notice Used to initialize the pointers to addresses crucial to the pod's functionality. Called on construction by the EigenPodManager.
    function initialize(
        address _podOwner
    ) external initializer {
        require(_podOwner != address(0), InputAddressZero());
        podOwner = _podOwner;
    }

    /**
     *
     *                                 EXTERNAL METHODS
     *
     */

    /// @notice payable fallback function that receives ether deposited to the eigenpods contract
    receive() external payable {
        emit NonBeaconChainETHReceived(msg.value);
    }

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
    ) external onlyOwnerOrProofSubmitter onlyWhenNotPaused(PAUSED_START_CHECKPOINT) {
        _startCheckpoint(revertIfNoBalance);
    }

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
    ) external onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS) {
        uint64 checkpointTimestamp = currentCheckpointTimestamp;
        require(checkpointTimestamp != 0, NoActiveCheckpoint());

        Checkpoint memory checkpoint = _currentCheckpoint;

        // Verify `balanceContainerProof` against `beaconBlockRoot`
        BeaconChainProofs.verifyBalanceContainer({
            beaconBlockRoot: checkpoint.beaconBlockRoot,
            proof: balanceContainerProof
        });

        // Process each checkpoint proof submitted
        uint64 exitedBalancesGwei;
        for (uint256 i = 0; i < proofs.length; i++) {
            BeaconChainProofs.BalanceProof calldata proof = proofs[i];
            ValidatorInfo memory validatorInfo = _validatorPubkeyHashToInfo[proof.pubkeyHash];

            // Validator must be in the ACTIVE state to be provable during a checkpoint.
            // Validators become ACTIVE when initially proven via verifyWithdrawalCredentials
            // Validators become WITHDRAWN when a checkpoint proof shows they have 0 balance
            if (validatorInfo.status != VALIDATOR_STATUS.ACTIVE) {
                continue;
            }

            // Ensure we aren't proving a validator twice for the same checkpoint. This will fail if:
            // - validator submitted twice during this checkpoint
            // - validator withdrawal credentials verified after checkpoint starts, then submitted
            //   as a checkpoint proof
            if (validatorInfo.lastCheckpointedAt >= checkpointTimestamp) {
                continue;
            }

            // Process a checkpoint proof for a validator and update its balance.
            //
            // If the proof shows the validator has a balance of 0, they are marked `WITHDRAWN`.
            // The assumption is that if this is the case, any withdrawn ETH was already in
            // the pod when `startCheckpoint` was originally called.
            (uint64 prevBalanceGwei, int64 balanceDeltaGwei, uint64 exitedBalanceGwei) = _verifyCheckpointProof({
                validatorInfo: validatorInfo,
                checkpointTimestamp: checkpointTimestamp,
                balanceContainerRoot: balanceContainerProof.balanceContainerRoot,
                proof: proof
            });

            checkpoint.proofsRemaining--;
            checkpoint.prevBeaconBalanceGwei += prevBalanceGwei;
            checkpoint.balanceDeltasGwei += balanceDeltaGwei;
            exitedBalancesGwei += exitedBalanceGwei;

            // Record the updated validator in state
            _validatorPubkeyHashToInfo[proof.pubkeyHash] = validatorInfo;
            emit ValidatorCheckpointed(checkpointTimestamp, uint40(validatorInfo.validatorIndex));
        }

        // Update the checkpoint and the total amount attributed to exited validators
        checkpointBalanceExitedGwei[checkpointTimestamp] += exitedBalancesGwei;
        _updateCheckpoint(checkpoint);
    }

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
    ) external onlyOwnerOrProofSubmitter onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS) {
        require(
            (validatorIndices.length == validatorFieldsProofs.length)
                && (validatorFieldsProofs.length == validatorFields.length),
            InputArrayLengthMismatch()
        );

        // Calling this method using a `beaconTimestamp` <= `currentCheckpointTimestamp` would allow
        // a newly-verified validator to be submitted to `verifyCheckpointProofs`, making progress
        // on an existing checkpoint.
        require(beaconTimestamp > currentCheckpointTimestamp, BeaconTimestampTooFarInPast());

        // Verify passed-in `beaconStateRoot` against the beacon block root
        // forgefmt: disable-next-item
        BeaconChainProofs.verifyStateRoot({
            beaconBlockRoot: getParentBlockRoot(beaconTimestamp),
            proof: stateRootProof
        });

        uint256 totalAmountToBeRestakedWei;
        for (uint256 i = 0; i < validatorIndices.length; i++) {
            // forgefmt: disable-next-item
            totalAmountToBeRestakedWei += _verifyWithdrawalCredentials(
                stateRootProof.beaconStateRoot,
                validatorIndices[i],
                validatorFieldsProofs[i],
                validatorFields[i]
            );
        }

        // Update the EigenPodManager on this pod's new balance
        eigenPodManager.recordBeaconChainETHBalanceUpdate({
            podOwner: podOwner,
            prevRestakedBalanceWei: 0, // only used for checkpoint balance updates
            balanceDeltaWei: int256(totalAmountToBeRestakedWei)
        });
    }

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
    ) external onlyWhenNotPaused(PAUSED_START_CHECKPOINT) onlyWhenNotPaused(PAUSED_VERIFY_STALE_BALANCE) {
        bytes32 validatorPubkey = proof.validatorFields.getPubkeyHash();
        ValidatorInfo memory validatorInfo = _validatorPubkeyHashToInfo[validatorPubkey];

        // Validator must be eligible for a staleness proof. Generally, this condition
        // ensures that the staleness proof is newer than the last time we got an update
        // on this validator.
        //
        // Note: It is possible for `validatorInfo.lastCheckpointedAt` to be 0 if
        // a validator's withdrawal credentials are verified when no checkpoint has
        // ever been completed in this pod. Technically, this would mean that `beaconTimestamp`
        // can be any valid EIP-4788 timestamp - because any nonzero value satisfies the
        // require below.
        //
        // However, in practice, if the only update we've seen from a validator is their
        // `verifyWithdrawalCredentials` proof, any valid `verifyStaleBalance` proof is
        // necessarily newer. This is because when a validator is initially slashed, their
        // exit epoch is set. And because `verifyWithdrawalCredentials` rejects validators
        // that have initiated exits, we know that if we're seeing a proof where the validator
        // is slashed that it MUST be newer than the `verifyWithdrawalCredentials` proof
        // (regardless of the relationship between `beaconTimestamp` and `lastCheckpointedAt`).
        require(beaconTimestamp > validatorInfo.lastCheckpointedAt, BeaconTimestampTooFarInPast());

        // Validator must be checkpoint-able
        require(validatorInfo.status == VALIDATOR_STATUS.ACTIVE, ValidatorNotActiveInPod());

        // Validator must be slashed on the beacon chain
        require(proof.validatorFields.isValidatorSlashed(), ValidatorNotSlashedOnBeaconChain());

        // Verify passed-in `beaconStateRoot` against the beacon block root
        // forgefmt: disable-next-item
        BeaconChainProofs.verifyStateRoot({
            beaconBlockRoot: getParentBlockRoot(beaconTimestamp),
            proof: stateRootProof
        });

        // Verify Validator container proof against `beaconStateRoot`
        BeaconChainProofs.verifyValidatorFields({
            beaconStateRoot: stateRootProof.beaconStateRoot,
            validatorFields: proof.validatorFields,
            validatorFieldsProof: proof.proof,
            validatorIndex: uint40(validatorInfo.validatorIndex)
        });

        // Validator verified to be stale - start a checkpoint
        _startCheckpoint(false);
    }

    /// @notice called by owner of a pod to remove any ERC20s deposited in the pod
    function recoverTokens(
        IERC20[] memory tokenList,
        uint256[] memory amountsToWithdraw,
        address recipient
    ) external onlyEigenPodOwner onlyWhenNotPaused(PAUSED_NON_PROOF_WITHDRAWALS) {
        require(tokenList.length == amountsToWithdraw.length, InputArrayLengthMismatch());
        for (uint256 i = 0; i < tokenList.length; i++) {
            tokenList[i].safeTransfer(recipient, amountsToWithdraw[i]);
        }
    }

    /// @notice Allows the owner of a pod to update the proof submitter, a permissioned
    /// address that can call `startCheckpoint` and `verifyWithdrawalCredentials`.
    /// @dev Note that EITHER the podOwner OR proofSubmitter can access these methods,
    /// so it's fine to set your proofSubmitter to 0 if you want the podOwner to be the
    /// only address that can call these methods.
    /// @param newProofSubmitter The new proof submitter address. If set to 0, only the
    /// pod owner will be able to call `startCheckpoint` and `verifyWithdrawalCredentials`
    function setProofSubmitter(
        address newProofSubmitter
    ) external onlyEigenPodOwner {
        emit ProofSubmitterUpdated(proofSubmitter, newProofSubmitter);
        proofSubmitter = newProofSubmitter;
    }

    /// @notice Called by EigenPodManager when the owner wants to create another ETH validator.
    function stake(
        bytes calldata pubkey,
        bytes calldata signature,
        bytes32 depositDataRoot
    ) external payable onlyEigenPodManager {
        // stake on ethpos
        require(msg.value == 32 ether, MsgValueNot32ETH());
        ethPOS.deposit{value: 32 ether}(pubkey, _podWithdrawalCredentials(), signature, depositDataRoot);
        emit EigenPodStaked(pubkey);
    }

    /**
     * @notice Transfers `amountWei` in ether from this contract to the specified `recipient` address
     * @notice Called by EigenPodManager to withdrawBeaconChainETH that has been added to the EigenPod's balance due to a withdrawal from the beacon chain.
     * @dev The podOwner must have already proved sufficient withdrawals, so that this pod's `restakedExecutionLayerGwei` exceeds the
     * `amountWei` input (when converted to GWEI).
     * @dev Reverts if `amountWei` is not a whole Gwei amount
     */
    function withdrawRestakedBeaconChainETH(address recipient, uint256 amountWei) external onlyEigenPodManager {
        uint64 amountGwei = uint64(amountWei / GWEI_TO_WEI);
        amountWei = amountGwei * GWEI_TO_WEI;
        require(amountGwei <= restakedExecutionLayerGwei, InsufficientWithdrawableBalance());
        restakedExecutionLayerGwei -= amountGwei;
        emit RestakedBeaconChainETHWithdrawn(recipient, amountWei);
        // transfer ETH from pod to `recipient` directly
        Address.sendValue(payable(recipient), amountWei);
    }

    /**
     *
     *                             INTERNAL FUNCTIONS
     *
     */

    /**
     * @notice internal function that proves an individual validator's withdrawal credentials
     * @param validatorIndex is the index of the validator being proven
     * @param validatorFieldsProof is the bytes that prove the ETH validator's  withdrawal credentials against a beacon chain state root
     * @param validatorFields are the fields of the "Validator Container", refer to consensus specs
     */
    function _verifyWithdrawalCredentials(
        bytes32 beaconStateRoot,
        uint40 validatorIndex,
        bytes calldata validatorFieldsProof,
        bytes32[] calldata validatorFields
    ) internal returns (uint256) {
        bytes32 pubkeyHash = validatorFields.getPubkeyHash();
        ValidatorInfo memory validatorInfo = _validatorPubkeyHashToInfo[pubkeyHash];

        // Withdrawal credential proofs should only be processed for "INACTIVE" validators
        require(validatorInfo.status == VALIDATOR_STATUS.INACTIVE, CredentialsAlreadyVerified());

        // Validator should be active on the beacon chain, or in the process of activating.
        // This implies the validator has reached the minimum effective balance required
        // to become active on the beacon chain.
        //
        // This check is important because the Pectra upgrade will move any validators that
        // do NOT have an activation epoch to a "pending deposit queue," temporarily resetting
        // their current and effective balances to 0. This balance can be restored if a deposit
        // is made to bring the validator's balance above the minimum activation balance.
        // (See https://github.com/ethereum/consensus-specs/blob/dev/specs/electra/fork.md#upgrading-the-state)
        //
        // In the context of EigenLayer slashing, this temporary reset would allow pod shares
        // to temporarily decrease, then be restored later. This would effectively prevent these
        // shares from being slashable on EigenLayer for a short period of time.
        require(
            validatorFields.getActivationEpoch() != BeaconChainProofs.FAR_FUTURE_EPOCH, ValidatorInactiveOnBeaconChain()
        );

        // Validator should not already be in the process of exiting. This is an important property
        // this method needs to enforce to ensure a validator cannot be already-exited by the time
        // its withdrawal credentials are verified.
        //
        // Note that when a validator initiates an exit, two values are set:
        // - exit_epoch
        // - withdrawable_epoch
        //
        // The latter of these two values describes an epoch after which the validator's ETH MIGHT
        // have been exited to the EigenPod, depending on the state of the beacon chain withdrawal
        // queue.
        //
        // Requiring that a validator has not initiated exit by the time the EigenPod sees their
        // withdrawal credentials guarantees that the validator has not fully exited at this point.
        //
        // This is because:
        // - the earliest beacon chain slot allowed for withdrawal credential proofs is the earliest
        //   slot available in the EIP-4788 oracle, which keeps the last 8192 slots.
        // - when initiating an exit, a validator's earliest possible withdrawable_epoch is equal to
        //   1 + MAX_SEED_LOOKAHEAD + MIN_VALIDATOR_WITHDRAWABILITY_DELAY == 261 epochs (8352 slots).
        //
        // (See https://eth2book.info/capella/part3/helper/mutators/#initiate_validator_exit)
        require(validatorFields.getExitEpoch() == BeaconChainProofs.FAR_FUTURE_EPOCH, ValidatorIsExitingBeaconChain());

        // Ensure the validator's withdrawal credentials are pointed at this pod
        require(
            validatorFields.getWithdrawalCredentials() == bytes32(_podWithdrawalCredentials()),
            WithdrawalCredentialsNotForEigenPod()
        );

        // Get the validator's effective balance. Note that this method uses effective balance, while
        // `verifyCheckpointProofs` uses current balance. Effective balance is updated per-epoch - so it's
        // less accurate, but is good enough for verifying withdrawal credentials.
        uint64 restakedBalanceGwei = validatorFields.getEffectiveBalanceGwei();

        // Verify passed-in validatorFields against verified beaconStateRoot:
        BeaconChainProofs.verifyValidatorFields({
            beaconStateRoot: beaconStateRoot,
            validatorFields: validatorFields,
            validatorFieldsProof: validatorFieldsProof,
            validatorIndex: validatorIndex
        });

        // Account for validator in future checkpoints. Note that if this pod has never started a
        // checkpoint before, `lastCheckpointedAt` will be zero here. This is fine because the main
        // purpose of `lastCheckpointedAt` is to enforce that newly-verified validators are not
        // eligible to progress already-existing checkpoints - however in this case, no checkpoints exist.
        activeValidatorCount++;
        uint64 lastCheckpointedAt =
            currentCheckpointTimestamp == 0 ? lastCheckpointTimestamp : currentCheckpointTimestamp;

        // Proofs complete - create the validator in state
        _validatorPubkeyHashToInfo[pubkeyHash] = ValidatorInfo({
            validatorIndex: validatorIndex,
            restakedBalanceGwei: restakedBalanceGwei,
            lastCheckpointedAt: lastCheckpointedAt,
            status: VALIDATOR_STATUS.ACTIVE
        });

        // Add the validator's balance to the checkpoint's previous beacon balance
        // Note that even if this checkpoint is not active, the next one will include
        // the validator's restaked balance during the checkpoint process
        _currentCheckpoint.prevBeaconBalanceGwei += restakedBalanceGwei;

        emit ValidatorRestaked(validatorIndex);
        emit ValidatorBalanceUpdated(validatorIndex, lastCheckpointedAt, restakedBalanceGwei);
        return restakedBalanceGwei * GWEI_TO_WEI;
    }

    function _verifyCheckpointProof(
        ValidatorInfo memory validatorInfo,
        uint64 checkpointTimestamp,
        bytes32 balanceContainerRoot,
        BeaconChainProofs.BalanceProof calldata proof
    ) internal returns (uint64 prevBalanceGwei, int64 balanceDeltaGwei, uint64 exitedBalanceGwei) {
        uint40 validatorIndex = uint40(validatorInfo.validatorIndex);

        // Verify validator balance against `balanceContainerRoot`
        prevBalanceGwei = validatorInfo.restakedBalanceGwei;
        uint64 newBalanceGwei = BeaconChainProofs.verifyValidatorBalance({
            balanceContainerRoot: balanceContainerRoot,
            validatorIndex: validatorIndex,
            proof: proof
        });

        // Calculate change in the validator's balance since the last proof
        if (newBalanceGwei != prevBalanceGwei) {
            balanceDeltaGwei = int64(newBalanceGwei) - int64(prevBalanceGwei);
            emit ValidatorBalanceUpdated(validatorIndex, checkpointTimestamp, newBalanceGwei);
        }

        validatorInfo.restakedBalanceGwei = newBalanceGwei;
        validatorInfo.lastCheckpointedAt = checkpointTimestamp;

        // If the validator's new balance is 0, mark them withdrawn
        if (newBalanceGwei == 0) {
            activeValidatorCount--;
            validatorInfo.status = VALIDATOR_STATUS.WITHDRAWN;
            // If we reach this point, `balanceDeltaGwei` should always be negative,
            // so this should be a safe conversion
            exitedBalanceGwei = uint64(-balanceDeltaGwei);

            emit ValidatorWithdrawn(checkpointTimestamp, validatorIndex);
        }

        return (prevBalanceGwei, balanceDeltaGwei, exitedBalanceGwei);
    }

    /**
     * @dev Initiate a checkpoint proof by snapshotting both the pod's ETH balance and the
     * current block's parent block root. After providing a checkpoint proof for each of the
     * pod's ACTIVE validators, the pod's ETH balance is awarded shares and can be withdrawn.
     * @dev ACTIVE validators are validators with verified withdrawal credentials (See
     * `verifyWithdrawalCredentials` for details)
     * @dev If the pod does not have any ACTIVE validators, the checkpoint is automatically
     * finalized.
     * @dev Once started, a checkpoint MUST be completed! It is not possible to start a
     * checkpoint if the existing one is incomplete.
     * @param revertIfNoBalance If the available ETH balance for checkpointing is 0 and this is
     * true, this method will revert
     */
    function _startCheckpoint(
        bool revertIfNoBalance
    ) internal {
        require(currentCheckpointTimestamp == 0, CheckpointAlreadyActive());

        // Prevent a checkpoint being completable twice in the same block. This prevents an edge case
        // where the second checkpoint would not be completable.
        //
        // This is because the validators checkpointed in the first checkpoint would have a `lastCheckpointedAt`
        // value equal to the second checkpoint, causing their proofs to get skipped in `verifyCheckpointProofs`
        require(lastCheckpointTimestamp != uint64(block.timestamp), CannotCheckpointTwiceInSingleBlock());

        // Snapshot pod balance at the start of the checkpoint, subtracting pod balance that has
        // previously been credited with shares. Once the checkpoint is finalized, `podBalanceGwei`
        // will be added to the total validator balance delta and credited as shares.
        //
        // Note: On finalization, `podBalanceGwei` is added to `restakedExecutionLayerGwei`
        // to denote that it has been credited with shares. Because this value is denominated in gwei,
        // `podBalanceGwei` is also converted to a gwei amount here. This means that any sub-gwei amounts
        // sent to the pod are not credited with shares and are therefore not withdrawable.
        // This can be addressed by topping up a pod's balance to a value divisible by 1 gwei.
        uint64 podBalanceGwei = uint64(address(this).balance / GWEI_TO_WEI) - restakedExecutionLayerGwei;

        // If the caller doesn't want a "0 balance" checkpoint, revert
        if (revertIfNoBalance && podBalanceGwei == 0) {
            revert NoBalanceToCheckpoint();
        }

        // Create checkpoint using the previous block's root for proofs, and the current
        // `activeValidatorCount` as the number of checkpoint proofs needed to finalize
        // the checkpoint.
        Checkpoint memory checkpoint = Checkpoint({
            beaconBlockRoot: getParentBlockRoot(uint64(block.timestamp)),
            proofsRemaining: uint24(activeValidatorCount),
            podBalanceGwei: podBalanceGwei,
            balanceDeltasGwei: 0,
            prevBeaconBalanceGwei: 0
        });

        // Place checkpoint in storage. If `proofsRemaining` is 0, the checkpoint
        // is automatically finalized.
        currentCheckpointTimestamp = uint64(block.timestamp);
        _updateCheckpoint(checkpoint);

        emit CheckpointCreated(uint64(block.timestamp), checkpoint.beaconBlockRoot, checkpoint.proofsRemaining);
    }

    /**
     * @dev Finish progress on a checkpoint and store it in state.
     * @dev If the checkpoint has no proofs remaining, it is finalized:
     * - a share delta is calculated and sent to the `EigenPodManager`
     * - the checkpointed `podBalanceGwei` is added to `restakedExecutionLayerGwei`
     * - `lastCheckpointTimestamp` is updated
     * - `_currentCheckpoint` and `currentCheckpointTimestamp` are deleted
     */
    function _updateCheckpoint(
        Checkpoint memory checkpoint
    ) internal {
        if (checkpoint.proofsRemaining != 0) {
            _currentCheckpoint = checkpoint;
            return;
        }

        // Calculate the previous total restaked balance and change in restaked balance
        // Note: due to how these values are calculated, a negative `balanceDeltaGwei`
        // should NEVER be greater in magnitude than `prevRestakedBalanceGwei`
        uint64 prevRestakedBalanceGwei = restakedExecutionLayerGwei + checkpoint.prevBeaconBalanceGwei;
        int64 balanceDeltaGwei = int64(checkpoint.podBalanceGwei) + checkpoint.balanceDeltasGwei;

        // And native ETH when the checkpoint was started is now considered restaked.
        // Add it to `restakedExecutionLayerGwei`, which allows it to be withdrawn via
        // the `DelegationManager` withdrawal queue.
        restakedExecutionLayerGwei += checkpoint.podBalanceGwei;

        // Finalize the checkpoint by resetting `currentCheckpointTimestamp`.
        // Note: `_currentCheckpoint` is not deleted, as it is overwritten
        // when a new checkpoint is started
        lastCheckpointTimestamp = currentCheckpointTimestamp;
        delete currentCheckpointTimestamp;

        // Convert shares and delta to wei
        uint256 prevRestakedBalanceWei = prevRestakedBalanceGwei * GWEI_TO_WEI;
        int256 balanceDeltaWei = balanceDeltaGwei * int256(GWEI_TO_WEI);

        // Update pod owner's shares
        emit CheckpointFinalized(lastCheckpointTimestamp, balanceDeltaWei);
        eigenPodManager.recordBeaconChainETHBalanceUpdate({
            podOwner: podOwner,
            prevRestakedBalanceWei: prevRestakedBalanceWei,
            balanceDeltaWei: balanceDeltaWei
        });
    }

    function _podWithdrawalCredentials() internal view returns (bytes memory) {
        return abi.encodePacked(bytes1(uint8(1)), bytes11(0), address(this));
    }

    ///@notice Calculates the pubkey hash of a validator's pubkey as per SSZ spec
    function _calculateValidatorPubkeyHash(
        bytes memory validatorPubkey
    ) internal pure returns (bytes32) {
        require(validatorPubkey.length == 48, InvalidPubKeyLength());
        return sha256(abi.encodePacked(validatorPubkey, bytes16(0)));
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /// @inheritdoc IEigenPod
    function withdrawableRestakedExecutionLayerGwei() external view returns (uint64) {
        return restakedExecutionLayerGwei;
    }

    /// @notice Returns the validatorInfo for a given validatorPubkeyHash
    function validatorPubkeyHashToInfo(
        bytes32 validatorPubkeyHash
    ) external view returns (ValidatorInfo memory) {
        return _validatorPubkeyHashToInfo[validatorPubkeyHash];
    }

    /// @notice Returns the validatorInfo for a given validatorPubkey
    function validatorPubkeyToInfo(
        bytes calldata validatorPubkey
    ) external view returns (ValidatorInfo memory) {
        return _validatorPubkeyHashToInfo[_calculateValidatorPubkeyHash(validatorPubkey)];
    }

    function validatorStatus(
        bytes32 pubkeyHash
    ) external view returns (VALIDATOR_STATUS) {
        return _validatorPubkeyHashToInfo[pubkeyHash].status;
    }

    /// @notice Returns the validator status for a given validatorPubkey
    function validatorStatus(
        bytes calldata validatorPubkey
    ) external view returns (VALIDATOR_STATUS) {
        bytes32 validatorPubkeyHash = _calculateValidatorPubkeyHash(validatorPubkey);
        return _validatorPubkeyHashToInfo[validatorPubkeyHash].status;
    }

    /// @notice Returns the currently-active checkpoint
    function currentCheckpoint() public view returns (Checkpoint memory) {
        return _currentCheckpoint;
    }

    /// @notice Query the 4788 oracle to get the parent block root of the slot with the given `timestamp`
    /// @param timestamp of the block for which the parent block root will be returned. MUST correspond
    /// to an existing slot within the last 24 hours. If the slot at `timestamp` was skipped, this method
    /// will revert.
    function getParentBlockRoot(
        uint64 timestamp
    ) public view returns (bytes32) {
        require(block.timestamp - timestamp < BEACON_ROOTS_HISTORY_BUFFER_LENGTH * 12, TimestampOutOfRange());

        (bool success, bytes memory result) = BEACON_ROOTS_ADDRESS.staticcall(abi.encode(timestamp));

        require(success && result.length > 0, InvalidEIP4788Response());
        return abi.decode(result, (bytes32));
    }
}
