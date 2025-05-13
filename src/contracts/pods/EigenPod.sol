// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "../libraries/BeaconChainProofs.sol";

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

    /// @notice The address of the EIP-7002 withdrawal request predeploy
    /// (See https://eips.ethereum.org/EIPS/eip-7002)
    address internal constant WITHDRAWAL_REQUEST_ADDRESS = 0x00000961Ef480Eb55e80D19ad83579A64c007002;

    /// @notice The address of the EIP-7251 consolidation request predeploy
    /// (See https://eips.ethereum.org/EIPS/eip-7251)
    address internal constant CONSOLIDATION_REQUEST_ADDRESS = 0x0000BBdDc7CE488642fb579F8B00f3a590007251;

    /// @notice The length of the EIP-4788 beacon block root ring buffer
    uint256 internal constant BEACON_ROOTS_HISTORY_BUFFER_LENGTH = 8191;

    /// @notice The beacon chain deposit contract
    IETHPOSDeposit public immutable ethPOS;

    /// @notice The single EigenPodManager for EigenLayer
    IEigenPodManager public immutable eigenPodManager;

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
        string memory _version
    ) SemVerMixin(_version) {
        ethPOS = _ethPOS;
        eigenPodManager = _eigenPodManager;
        _disableInitializers();
    }

    /// @inheritdoc IEigenPod
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

    /// @notice payable fallback function used to receive ETH sent directly to the pod
    receive() external payable {
        emit NonBeaconChainETHReceived(msg.value);
    }

    /// @inheritdoc IEigenPod
    function startCheckpoint(
        bool revertIfNoBalance
    ) external onlyOwnerOrProofSubmitter onlyWhenNotPaused(PAUSED_START_CHECKPOINT) {
        _startCheckpoint(revertIfNoBalance);
    }

    /// @inheritdoc IEigenPod
    function verifyCheckpointProofs(
        BeaconChainProofs.BalanceContainerProof calldata balanceContainerProof,
        BeaconChainProofs.BalanceProof[] calldata proofs
    ) external onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS) {
        uint64 checkpointTimestamp = currentCheckpointTimestamp;
        require(checkpointTimestamp != 0, NoActiveCheckpoint());

        Checkpoint memory checkpoint = _currentCheckpoint;

        // Verify `balanceContainerProof` against `beaconBlockRoot`
        BeaconChainProofs.verifyBalanceContainer({
            proofVersion: _getProofVersion(checkpointTimestamp),
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

    /// @inheritdoc IEigenPod
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
                beaconTimestamp,
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

    /// @inheritdoc IEigenPod
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
            proofVersion: _getProofVersion(beaconTimestamp),
            beaconStateRoot: stateRootProof.beaconStateRoot,
            validatorFields: proof.validatorFields,
            validatorFieldsProof: proof.proof,
            validatorIndex: uint40(validatorInfo.validatorIndex)
        });

        // Validator verified to be stale - start a checkpoint
        _startCheckpoint(false);
    }

    /// @inheritdoc IEigenPod
    function requestConsolidation(
        ConsolidationRequest[] calldata requests
    ) external payable onlyWhenNotPaused(PAUSED_CONSOLIDATIONS) onlyOwnerOrProofSubmitter {
        uint256 fee = getConsolidationRequestFee();
        require(msg.value >= fee * requests.length, InsufficientFunds());
        uint256 remainder = msg.value - (fee * requests.length);

        for (uint256 i = 0; i < requests.length; i++) {
            ConsolidationRequest calldata request = requests[i];
            // Validate pubkeys are well-formed
            require(request.srcPubkey.length == 48, InvalidPubKeyLength());
            require(request.targetPubkey.length == 48, InvalidPubKeyLength());

            // Ensure target has verified withdrawal credentials pointed at this pod
            bytes32 sourcePubkeyHash = _calcPubkeyHash(request.srcPubkey);
            bytes32 targetPubkeyHash = _calcPubkeyHash(request.targetPubkey);
            ValidatorInfo memory target = validatorPubkeyHashToInfo(targetPubkeyHash);
            require(target.status == VALIDATOR_STATUS.ACTIVE, ValidatorNotActiveInPod());

            // Call the predeploy
            bytes memory callData = bytes.concat(request.srcPubkey, request.targetPubkey);
            (bool ok,) = CONSOLIDATION_REQUEST_ADDRESS.call{value: fee}(callData);
            require(ok, PredeployFailed());

            // Emit event depending on whether this is a switch to 0x02, or a regular consolidation
            if (sourcePubkeyHash == targetPubkeyHash) emit SwitchToCompoundingRequested(sourcePubkeyHash);
            else emit ConsolidationRequested(sourcePubkeyHash, targetPubkeyHash);
        }

        // Refund remainder of msg.value
        if (remainder > 0) {
            Address.sendValue(payable(msg.sender), remainder);
        }
    }

    /// @inheritdoc IEigenPod
    function requestWithdrawal(
        WithdrawalRequest[] calldata requests
    ) external payable onlyWhenNotPaused(PAUSED_WITHDRAWAL_REQUESTS) onlyOwnerOrProofSubmitter {
        uint256 fee = getWithdrawalRequestFee();
        require(msg.value >= fee * requests.length, InsufficientFunds());
        uint256 remainder = msg.value - (fee * requests.length);

        for (uint256 i = 0; i < requests.length; i++) {
            WithdrawalRequest calldata request = requests[i];
            // Validate pubkey is well-formed.
            //
            // It's not necessary to perform any additional validation; the worst-case
            // scenario is just that the consensus layer skips an invalid request.
            require(request.pubkey.length == 48, InvalidPubKeyLength());

            // Call the predeploy
            bytes memory callData = abi.encodePacked(request.pubkey, request.amountGwei);
            (bool ok,) = WITHDRAWAL_REQUEST_ADDRESS.call{value: fee}(callData);
            require(ok, PredeployFailed());

            // Emit event depending on whether the request is a full exit or a partial withdrawal
            bytes32 pubkeyHash = _calcPubkeyHash(request.pubkey);
            if (request.amountGwei == 0) emit ExitRequested(pubkeyHash);
            else emit WithdrawalRequested(pubkeyHash, request.amountGwei);
        }

        // Refund remainder of msg.value
        if (remainder > 0) {
            Address.sendValue(payable(msg.sender), remainder);
        }
    }

    /// @inheritdoc IEigenPod
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

    /// @inheritdoc IEigenPod
    function setProofSubmitter(
        address newProofSubmitter
    ) external onlyEigenPodOwner {
        emit ProofSubmitterUpdated(proofSubmitter, newProofSubmitter);
        proofSubmitter = newProofSubmitter;
    }

    /// @inheritdoc IEigenPod
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

    /// @inheritdoc IEigenPod
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
        uint64 beaconTimestamp,
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
            validatorFields.getWithdrawalCredentials() == bytes32(_podWithdrawalCredentials())
                || validatorFields.getWithdrawalCredentials() == bytes32(_podCompoundingWithdrawalCredentials()),
            WithdrawalCredentialsNotForEigenPod()
        );

        // Get the validator's effective balance. Note that this method uses effective balance, while
        // `verifyCheckpointProofs` uses current balance. Effective balance is updated per-epoch - so it's
        // less accurate, but is good enough for verifying withdrawal credentials.
        uint64 restakedBalanceGwei = validatorFields.getEffectiveBalanceGwei();

        // Verify passed-in validatorFields against verified beaconStateRoot:
        BeaconChainProofs.verifyValidatorFields({
            proofVersion: _getProofVersion(beaconTimestamp),
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
     * - `currentCheckpointTimestamp` is set to zero
     */
    function _updateCheckpoint(
        Checkpoint memory checkpoint
    ) internal {
        _currentCheckpoint = checkpoint;
        if (checkpoint.proofsRemaining != 0) {
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

    function _podCompoundingWithdrawalCredentials() internal view returns (bytes memory) {
        return abi.encodePacked(bytes1(uint8(2)), bytes11(0), address(this));
    }

    ///@notice Calculates the pubkey hash of a validator's pubkey as per SSZ spec
    function _calcPubkeyHash(
        bytes memory validatorPubkey
    ) internal pure returns (bytes32) {
        require(validatorPubkey.length == 48, InvalidPubKeyLength());
        return sha256(abi.encodePacked(validatorPubkey, bytes16(0)));
    }

    /// @dev Returns the current fee required to query either the EIP-7002 or EIP-7251 predeploy
    function _getFee(
        address predeploy
    ) internal view returns (uint256) {
        (bool success, bytes memory result) = predeploy.staticcall("");
        require(success && result.length == 32, FeeQueryFailed());

        return uint256(bytes32(result));
    }

    /// @notice Returns the PROOF_TYPE depending on the `proofTimestamp` in relation to the fork timestamp.
    function _getProofVersion(
        uint64 proofTimestamp
    ) internal view returns (BeaconChainProofs.ProofVersion) {
        /// Get the timestamp of the Pectra fork, read from the `EigenPodManager`
        /// This returns the timestamp of the first non-missed slot at or after the Pectra hard fork
        uint64 forkTimestamp = eigenPodManager.pectraForkTimestamp();
        require(forkTimestamp != 0, ForkTimestampZero());

        /// We check if the proofTimestamp is <= pectraForkTimestamp because a `proofTimestamp` at the `pectraForkTimestamp`
        /// is considered to be Pre-Pectra given the EIP-4788 oracle returns the parent block.
        return proofTimestamp <= forkTimestamp
            ? BeaconChainProofs.ProofVersion.DENEB
            : BeaconChainProofs.ProofVersion.PECTRA;
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

    /// @inheritdoc IEigenPod
    function validatorPubkeyHashToInfo(
        bytes32 validatorPubkeyHash
    ) public view returns (ValidatorInfo memory) {
        return _validatorPubkeyHashToInfo[validatorPubkeyHash];
    }

    /// @inheritdoc IEigenPod
    function validatorPubkeyToInfo(
        bytes calldata validatorPubkey
    ) public view returns (ValidatorInfo memory) {
        return _validatorPubkeyHashToInfo[_calcPubkeyHash(validatorPubkey)];
    }

    /// @inheritdoc IEigenPod
    function validatorStatus(
        bytes32 pubkeyHash
    ) external view returns (VALIDATOR_STATUS) {
        return _validatorPubkeyHashToInfo[pubkeyHash].status;
    }

    /// @inheritdoc IEigenPod
    function validatorStatus(
        bytes calldata validatorPubkey
    ) external view returns (VALIDATOR_STATUS) {
        bytes32 validatorPubkeyHash = _calcPubkeyHash(validatorPubkey);
        return _validatorPubkeyHashToInfo[validatorPubkeyHash].status;
    }

    /// @inheritdoc IEigenPod
    function currentCheckpoint() public view returns (Checkpoint memory) {
        return _currentCheckpoint;
    }

    /// @inheritdoc IEigenPod
    function getParentBlockRoot(
        uint64 timestamp
    ) public view returns (bytes32) {
        require(block.timestamp - timestamp < BEACON_ROOTS_HISTORY_BUFFER_LENGTH * 12, TimestampOutOfRange());

        (bool success, bytes memory result) = BEACON_ROOTS_ADDRESS.staticcall(abi.encode(timestamp));

        require(success && result.length > 0, InvalidEIP4788Response());
        return abi.decode(result, (bytes32));
    }

    /// @inheritdoc IEigenPod
    function getConsolidationRequestFee() public view returns (uint256) {
        return _getFee(CONSOLIDATION_REQUEST_ADDRESS);
    }

    /// @inheritdoc IEigenPod
    function getWithdrawalRequestFee() public view returns (uint256) {
        return _getFee(WITHDRAWAL_REQUEST_ADDRESS);
    }
}
