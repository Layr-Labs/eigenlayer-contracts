// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/utils/AddressUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/utils/math/MathUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "../libraries/BeaconChainProofs.sol";
import "../libraries/BytesLib.sol";
import "../libraries/Endian.sol";

import "../interfaces/IETHPOSDeposit.sol";
import "../interfaces/IEigenPodManager.sol";
import "../interfaces/IDelayedWithdrawalRouter.sol";
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
contract EigenPod is EigenPodStorage, Initializable, ReentrancyGuardUpgradeable, EigenPodPausingConstants {
    
    using BytesLib for bytes;
    using SafeERC20 for IERC20;
    using BeaconChainProofs for *;

    /*******************************************************************************
                               CONSTANTS / IMMUTABLES
    *******************************************************************************/

    // @notice Internal constant used in calculations, since the beacon chain stores balances in Gwei rather than wei
    uint256 internal constant GWEI_TO_WEI = 1e9;

    /// @notice This is the beacon chain deposit contract
    IETHPOSDeposit public immutable ethPOS;

    /// @notice Contract used for withdrawal routing, to provide an extra "safety net" mechanism
    IDelayedWithdrawalRouter public immutable delayedWithdrawalRouter;

    /// @notice The single EigenPodManager for EigenLayer
    IEigenPodManager public immutable eigenPodManager;

    /// @notice This is the genesis time of the beacon state, to help us calculate conversions between slot and timestamp
    uint64 public immutable GENESIS_TIME;

    /// @notice If a validator is slashed on the beacon chain and their balance has not been checkpointed
    /// within `TIME_TILL_STALE_BALANCE` of the current block, they are eligible to be marked "stale"
    /// via `verifyStaleBalance`.
    uint256 internal constant TIME_TILL_STALE_BALANCE = 2 weeks;

    /// @notice A `verifyStaleBalance` proof MUST be older than `STALENESS_GRACE_PERIOD`, to give time
    /// to a pod owner to prove the slashed validator's balance
    uint256 internal constant STALENESS_GRACE_PERIOD = 6 hours;

    /*******************************************************************************
                                     MODIFIERS
    *******************************************************************************/

    modifier onlyEigenPodManager() {
        require(msg.sender == address(eigenPodManager), "EigenPod.onlyEigenPodManager: not eigenPodManager");
        _;
    }

    modifier onlyEigenPodOwner() {
        require(msg.sender == podOwner, "EigenPod.onlyEigenPodOwner: not podOwner");
        _;
    }

    modifier hasNeverRestaked() {
        require(!hasRestaked, "EigenPod.hasNeverRestaked: restaking is enabled");
        _;
    }

    /// @notice Checks that `timestamp` is greater than or equal to `mostRecentWithdrawalTimestamp`
    modifier afterRestaking(uint64 timestamp) {
        require(
            hasRestaked &&
            timestamp >= mostRecentWithdrawalTimestamp,
            "EigenPod.afterRestaking: timestamp must be at or after restaking was activated"
        );
        _;
    }

    /**
     * @notice Based on 'Pausable' code, but uses the storage of the EigenPodManager instead of this contract. This construction
     * is necessary for enabling pausing all EigenPods at the same time (due to EigenPods being Beacon Proxies).
     * Modifier throws if the `indexed`th bit of `_paused` in the EigenPodManager is 1, i.e. if the `index`th pause switch is flipped.
     */
    modifier onlyWhenNotPaused(uint8 index) {
        require(
            !IPausable(address(eigenPodManager)).paused(index),
            "EigenPod.onlyWhenNotPaused: index is paused in EigenPodManager"
        );
        _;
    }

    /*******************************************************************************
                                  CONSTRUCTOR / INIT
    *******************************************************************************/

    constructor(
        IETHPOSDeposit _ethPOS,
        IDelayedWithdrawalRouter _delayedWithdrawalRouter,
        IEigenPodManager _eigenPodManager,
        uint64 _GENESIS_TIME
    ) {
        ethPOS = _ethPOS;
        delayedWithdrawalRouter = _delayedWithdrawalRouter;
        eigenPodManager = _eigenPodManager;
        GENESIS_TIME = _GENESIS_TIME;
        _disableInitializers();
    }

    /// @notice Used to initialize the pointers to addresses crucial to the pod's functionality. Called on construction by the EigenPodManager.
    function initialize(address _podOwner) external initializer {
        require(_podOwner != address(0), "EigenPod.initialize: podOwner cannot be zero address");
        podOwner = _podOwner;
        /**
         * From the M2 deployment onwards, we are requiring that pods deployed are by default enabled with restaking
         * In prior deployments without proofs, EigenPods could be deployed with restaking disabled so as to allow
         * simple (proof-free) withdrawals.  However, this is no longer the case.  Thus going forward, all pods are
         * initialized with hasRestaked set to true.
         */
        hasRestaked = true;
        emit RestakingActivated(podOwner);
    }

    /*******************************************************************************
                                    EXTERNAL METHODS
    *******************************************************************************/

    /// @notice payable fallback function that receives ether deposited to the eigenpods contract
    receive() external payable {
        nonBeaconChainETHBalanceWei += msg.value;
        emit NonBeaconChainETHReceived(msg.value);
    }

    /**
     * @dev Create a checkpoint used to prove this pod's active validator set. Checkpoints are completed 
     * by submitting one checkpoint proof per ACTIVE validator. During the checkpoint process, the total 
     * change in ACTIVE validator balance is tracked, and any validators with 0 balance are marked `WITHDRAWN`.
     * @dev Once finalized, the pod owner is awarded shares corresponding to:
     * - the total change in their ACTIVE validator balances
     * - any ETH in the pod not already awarded shares.
     * @dev A checkpoint cannot be created if the pod already has an outstanding checkpoint. If
     * this is the case, the pod owner MUST complete the existing checkpoint before starting a new one.
     */
    function startCheckpoint() 
        onlyEigenPodOwner() 
        onlyWhenNotPaused(PAUSED_START_CHECKPOINT) 
        afterRestaking(uint64(block.timestamp)) /// TODO - this is the wrong condition
    {
        require(
            currentCheckpointTimestamp == 0, 
            "EigenPod.startCheckpoint: must finish previous checkpoint before starting another"
        );

        // Snapshot pod balance at the start of the checkpoint. Once the checkpoint is finalized,
        // this amount will be added to the total validator balance delta and credited as shares.
        uint256 podBalanceWei = 
            address(this).balance
                - nonBeaconChainETHBalanceWei
                - (withdrawableRestakedExecutionLayerGwei * GWEI_TO_WEI);

        Checkpoint memory checkpoint = Checkpoint({
            beaconBlockRoot: eigenPodManager.getParentBlockRoot(uint64(block.timestamp)),
            beaconStateRoot: bytes32(0),
            podBalanceGwei: podBalanceWei / GWEI_TO_WEI,
            balanceDeltas: 0,
            proofsRemaining: activeValidatorCount
        });

        // Place checkpoint in storage
        currentCheckpointTimestamp = uint64(block.timestamp);
        _updateCheckpoint(checkpoint);
    }

    /**
     * @dev Progress the current checkpoint towards completion by submitting one or more validator
     * checkpoint proofs. Anyone can call this method to submit proofs towards the current checkpoint.
     * For each validator proven, the current checkpoint's `proofsRemaining` decreases.
     * @dev If the checkpoint's `proofsRemaining` reaches 0, the checkpoint is finalized.
     * (see `_updateCheckpoint` for more details)
     * @dev This method can only be called when there is a currently-active checkpoint.
     * @param stateRootProof proves a beacon state root against the checkpoint's `beaconBlockRoot`
     * @dev Note that if the `beaconStateRoot` has already been verified for this checkpoint, it does
     * not need to be verfied again.
     * @param proofs Proofs for one or more validator current balances against the checkpoint's `beaconStateRoot`
     */
    function verifyCheckpointProofs(
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        BeaconChainProofs.BalanceProof[] calldata proofs
    ) 
        external 
        onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS) 
    {
        uint64 beaconTimestamp = currentCheckpointTimestamp;
        require(
            beaconTimestamp != 0, 
            "EigenPod.verifyCheckpointProofs: must have active checkpoint to perform checkpoint proof"
        );

        Checkpoint memory checkpoint = currentCheckpoint;

        // If we haven't already proven a state root for this checkpoint, verify
        // `stateRootProof` against the block root and cache the result
        if (checkpoint.beaconStateRoot == bytes32(0)) {
            BeaconChainProofs.verifyStateRootAgainstLatestBlockRoot({
                latestBlockRoot: checkpoint.beaconBlockRoot,
                beaconStateRoot: stateRootProof.beaconStateRoot,
                stateRootProof: stateRootProof.proof
            });

            checkpoint.beaconStateRoot = stateRootProof.beaconStateRoot;
        }

        // Process each checkpoint proof submitted
        int256 totalStaleValidatorsProven;
        for (uint256 i = 0; i < proofs.length; i++) {

            // Process a checkpoint proof for a validator. 
            // - the validator MUST be in the ACTIVE state
            // - the validator MUST NOT have already been proven for this checkpoint
            //
            // If the proof shows the validator has a balance of 0, they are marked `WITHDRAWN`.
            // The assumption is that if this is the case, any withdrawn ETH was already in
            // the pod when `startCheckpoint` was originally called.
            (int256 balanceDeltaWei, bool noLongerStale) = _verifyCheckpointProof({
                beaconTimestamp: beaconTimestamp,
                beaconStateRoot: checkpoint.beaconStateRoot,
                proof: proofs[i]
            });

            checkpoint.proofsRemaining--;
            checkpoint.balanceDeltas += balanceDeltaWei;

            if (noLongerStale) {
                totalStaleValidatorsProven++;
            }
        }

        // Update and/or finalize the checkpoint
        _updateCheckpoint(checkpoint);

        if (totalStaleValidatorsProven != 0) {
            eigenPodManager.updateStaleValidatorCount(podOwner, totalStaleValidatorsProven);
        }
    }

    /**
     * @dev Verify one or more validators have their withdrawal credentials pointed at this EigenPod, and award
     * shares based on their effective balance. Proven validators are marked `ACTIVE` within the EigenPod, and
     * future checkpoint proofs will need to include them.
     * @dev Withdrawal credential proofs MUST NOT be older than the `lastFinalizedCheckpoint` OR `currentCheckpointTimestamp`.
     * @dev Validators proven via this method MUST NOT have an exit epoch set already.
     * @param beaconTimestamp the beacon chain timestamp sent to the 4788 oracle contract. Corresponds
     * to the parent beacon block root against which the proof is verified.
     * @param stateRootProof proves a beacon state root against a beacon block root
     * @param validatorIndices a list of validator indices being proven stale
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
    )
        external
        onlyEigenPodOwner
        onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS)
        afterRestaking(beaconTimestamp)
    {
        require(
            (validatorIndices.length == validatorFieldsProofs.length) &&
                (validatorFieldsProofs.length == validatorFields.length),
            "EigenPod.verifyWithdrawalCredentials: validatorIndices and proofs must be same length"
        );

        // TODO - is == valid?
        // I think YES, because we enforce no exit epoch is set
        // However, we could also do strictly greater than if we want to be safe.
        // TODO pt 2 - what about sub 12 on checkpointTimestamp?
        require(
            beaconTimestamp > lastFinalizedCheckpoint && beaconTimestamp > currentCheckpointTimestamp,
            "EigenPod.verifyWithdrawalCredentials: specified timestamp is too far in past"
        );

        // Verify passed-in beaconStateRoot against oracle-provided block root:
        BeaconChainProofs.verifyStateRootAgainstLatestBlockRoot({
            latestBlockRoot: eigenPodManager.getParentBlockRoot(beaconTimestamp),
            beaconStateRoot: stateRootProof.beaconStateRoot,
            stateRootProof: stateRootProof.proof
        });

        uint256 totalAmountToBeRestakedWei;
        for (uint256 i = 0; i < validatorIndices.length; i++) {
            totalAmountToBeRestakedWei += _verifyWithdrawalCredentials(
                beaconTimestamp,
                stateRootProof.beaconStateRoot,
                validatorIndices[i],
                validatorFieldsProofs[i],
                validatorFields[i]
            );
        }

        // Update the EigenPodManager on this pod's new balance
        eigenPodManager.recordBeaconChainETHBalanceUpdate(podOwner, int256(totalAmountToBeRestakedWei));
    }

    /**
     * @dev Prove that one or more validators were slashed on the beacon chain and have not had timely
     * checkpoint proofs since being slashed. If successful, this increases the pod owner's `staleValidatorCount`
     * in the `EigenPodManager`. Stale validators can be restored by proving their balance in a checkpoint via
     * `startCheckpoint` and `verifyCheckpointProofs`.
     * @param beaconTimestamp the beacon chain timestamp sent to the 4788 oracle contract. Corresponds
     * to the parent beacon block root against which the proof is verified.
     * @param stateRootProof proves a beacon state root against a beacon block root
     * @param proofs the fields of the beacon chain "Validator" container, along with a merkle proof against
     * the beacon state root. See the consensus specs for more details:
     * https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator
     *
     * @dev Staleness conditions:
     * - `beaconTimestamp` MUST NOT fall within `STALENESS_GRACE_PERIOD` seconds of `block.timestamp`
     * - Validator's last balance update is older than `beaconTimestamp` by `TIME_TILL_STALE_BALANCE`
     * - Validator MUST be in `ACTIVE` status in the pod
     * - Validator MUST NOT already be marked stale
     * - Validator MUST be slashed on the beacon chain
     */
    function verifyStaleBalances(
        uint64 beaconTimestamp,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        BeaconChainProofs.ValidatorProof calldata proofs
    )
        external
        onlyWhenNotPaused(PAUSED_VERIFY_STALE_BALANCE)
    {
        require(
            beaconTimestamp + STALENESS_GRACE_PERIOD < block.timestamp,
            "EigenPod.verifyStaleBalance: staleness grace period not elapsed"
        );

        // Process each staleness proof
        for (uint256 i = 0; i < proofs.length; i++) {
            BeaconChainProofs.ValidatorProof memory proof = proofs[i];

            bytes32 validatorPubkeyHash = proof.validatorFields.getPubkeyHash();
            ValidatorInfo memory validatorInfo = _validatorPubkeyHashToInfo[validatorPubkeyHash];

            // Validator must be eligible for a staleness proof
            require(
                beaconTimestamp > validatorInfo.mostRecentBalanceUpdateTimestamp + TIME_TILL_STALE_BALANCE,
                "EigenPod.verifyStaleBalance: validator balance is not stale yet"
            );

            // Validator must be checkpoint-able
            require(
                validatorInfo.status == VALIDATOR_STATUS.ACTIVE,
                "EigenPod.verifyStaleBalance: validator is not active"
            );

            // Validator should not already be stale
            require(
                !isValidatorStale[validatorPubkeyHash],
                "EigenPod.verifyStaleBalance: validator already marked stale"
            );

            // Validator must be slashed on the beacon chain
            require(
                proof.validatorFields.getSlashStatus() == true,
                "EigenPod.verifyStaleBalance: validator must be slashed to be marked stale"
            );

            // Verify `beaconStateRoot` against beacon block root
            BeaconChainProofs.verifyStateRootAgainstLatestBlockRoot({
                latestBlockRoot: eigenPodManager.getParentBlockRoot(beaconTimestamp),
                beaconStateRoot: stateRootProof.beaconStateRoot,
                stateRootProof: stateRootProof.proof
            });

            // Verify Validator container proof against `beaconStateRoot`
            BeaconChainProofs.verifyValidatorFields({
                beaconStateRoot: stateRootProof.beaconStateRoot,
                validatorFields: proof.validatorFields,
                validatorFieldsProof: proof.proof,
                validatorIndex: validatorInfo.validatorIndex
            });

            isValidatorStale[validatorPubkey] = true;
        }
        
        // Increment stale validator count by one for each successful proof
        eigenPodManager.updateStaleValidatorCount(podOwner, int256(proofs.length));
    }

    /// @notice Called by the pod owner to withdraw the nonBeaconChainETHBalanceWei
    function withdrawNonBeaconChainETHBalanceWei(
        address recipient,
        uint256 amountToWithdraw
    ) external onlyEigenPodOwner onlyWhenNotPaused(PAUSED_NON_PROOF_WITHDRAWALS) {
        require(
            amountToWithdraw <= nonBeaconChainETHBalanceWei,
            "EigenPod.withdrawnonBeaconChainETHBalanceWei: amountToWithdraw is greater than nonBeaconChainETHBalanceWei"
        );
        nonBeaconChainETHBalanceWei -= amountToWithdraw;
        emit NonBeaconChainETHWithdrawn(recipient, amountToWithdraw);
        _sendETH_AsDelayedWithdrawal(recipient, amountToWithdraw);
    }

    /// @notice called by owner of a pod to remove any ERC20s deposited in the pod
    function recoverTokens(
        IERC20[] memory tokenList,
        uint256[] memory amountsToWithdraw,
        address recipient
    ) external onlyEigenPodOwner onlyWhenNotPaused(PAUSED_NON_PROOF_WITHDRAWALS) {
        require(
            tokenList.length == amountsToWithdraw.length,
            "EigenPod.recoverTokens: tokenList and amountsToWithdraw must be same length"
        );
        for (uint256 i = 0; i < tokenList.length; i++) {
            tokenList[i].safeTransfer(recipient, amountsToWithdraw[i]);
        }
    }

    /**
     * @notice Called by the pod owner to activate restaking by withdrawing
     * all existing ETH from the pod and preventing further withdrawals via
     * "withdrawBeforeRestaking()"
     */
    function activateRestaking()
        external
        onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_CREDENTIALS)
        onlyEigenPodOwner
        hasNeverRestaked
    {
        hasRestaked = true;
        _processWithdrawalBeforeRestaking(podOwner);

        emit RestakingActivated(podOwner);
    }

    /// @notice Called by the pod owner to withdraw the balance of the pod when `hasRestaked` is set to false
    function withdrawBeforeRestaking() external onlyEigenPodOwner hasNeverRestaked {
        _processWithdrawalBeforeRestaking(podOwner);
    }

    /// @notice Called by EigenPodManager when the owner wants to create another ETH validator.
    function stake(
        bytes calldata pubkey,
        bytes calldata signature,
        bytes32 depositDataRoot
    ) external payable onlyEigenPodManager {
        // stake on ethpos
        require(msg.value == 32 ether, "EigenPod.stake: must initially stake for any validator with 32 ether");
        ethPOS.deposit{value: 32 ether}(pubkey, _podWithdrawalCredentials(), signature, depositDataRoot);
        emit EigenPodStaked(pubkey);
    }

    /**
     * @notice Transfers `amountWei` in ether from this contract to the specified `recipient` address
     * @notice Called by EigenPodManager to withdrawBeaconChainETH that has been added to the EigenPod's balance due to a withdrawal from the beacon chain.
     * @dev The podOwner must have already proved sufficient withdrawals, so that this pod's `withdrawableRestakedExecutionLayerGwei` exceeds the
     * `amountWei` input (when converted to GWEI).
     * @dev Reverts if `amountWei` is not a whole Gwei amount
     */
    function withdrawRestakedBeaconChainETH(address recipient, uint256 amountWei) external onlyEigenPodManager {
        require(amountWei % GWEI_TO_WEI == 0, "EigenPod.withdrawRestakedBeaconChainETH: amountWei must be a whole Gwei amount");
        uint64 amountGwei = uint64(amountWei / GWEI_TO_WEI);
        require(amountGwei <= withdrawableRestakedExecutionLayerGwei, "EigenPod.withdrawRestakedBeaconChainETH: amountGwei exceeds withdrawableRestakedExecutionLayerGwei");
        withdrawableRestakedExecutionLayerGwei -= amountGwei;
        emit RestakedBeaconChainETHWithdrawn(recipient, amountWei);
        // transfer ETH from pod to `recipient` directly
        _sendETH(recipient, amountWei);
    }

    /*******************************************************************************
                                INTERNAL FUNCTIONS
    *******************************************************************************/

    /**
     * @notice internal function that proves an individual validator's withdrawal credentials
     * @param beaconTimestamp is the timestamp whose state root the `proof` will be proven against.
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
        bytes32 validatorPubkeyHash = validatorFields.getPubkeyHash();
        ValidatorInfo memory validatorInfo = _validatorPubkeyHashToInfo[validatorPubkeyHash];

        // Withdrawal credential proofs should only be processed for "INACTIVE" validators
        require(
            validatorInfo.status == VALIDATOR_STATUS.INACTIVE,
            "EigenPod._verifyWithdrawalCredentials: validator must be inactive to prove withdrawal credentials"
        );

        // Validator should not already be in the process of exiting
        require(
            validatorFields.getExitEpoch() == FAR_FUTURE_EPOCH,
            "EigenPod._verifyWithdrawalCredentials: validator must not be exiting"
        );

        // Ensure the validator's withdrawal credentials are pointed at this pod
        require(
            validatorFields.getWithdrawalCredentials() == bytes32(_podWithdrawalCredentials()),
            "EigenPod._verifyWithdrawalCredentials: proof is not for this EigenPod"
        );

        // Get the validator's effective balance. Note that this method uses effective balance, while
        // `verifyBalanceUpdates` uses current balance. Effective balance is updated per-epoch - so it's
        // less accurate, but is good enough for verifying withdrawal credentials.
        uint64 restakedBalanceGwei = validatorFields.getEffectiveBalanceGwei();

        // Verify passed-in validatorFields against verified beaconStateRoot:
        BeaconChainProofs.verifyValidatorFields({
            beaconStateRoot: beaconStateRoot,
            validatorFields: validatorFields,
            validatorFieldsProof: validatorFieldsProof,
            validatorIndex: validatorIndex
        });

        // Proofs complete - update this validator's status, record its proven balance, and save in state:
        activeValidatorCount++;
        validatorInfo.status = VALIDATOR_STATUS.ACTIVE;
        validatorInfo.validatorIndex = validatorIndex;
        validatorInfo.mostRecentBalanceUpdateTimestamp = beaconTimestamp;
        validatorInfo.restakedBalanceGwei = restakedBalanceGwei;

        _validatorPubkeyHashToInfo[validatorPubkeyHash] = validatorInfo;

        emit ValidatorRestaked(validatorIndex);
        emit ValidatorBalanceUpdated(validatorIndex, beaconTimestamp, restakedBalanceGwei);

        return restakedBalanceGwei * GWEI_TO_WEI;
    }

    function _verifyCheckpointProof(
        uint64 beaconTimestamp,
        bytes32 beaconStateRoot,
        BeaconChainProofs.BalanceProof memory proof
    ) internal returns (int256 balanceDeltaWei, bool noLongerStale) {
        ValidatorInfo memory validatorInfo = _validatorPubkeyHashToInfo[proof.pubkeyHash];

        require(
            validatorInfo.status == VALIDATOR_STATUS.ACTIVE,
            "EigenPod._verifyCheckpointProof: validator must be ACTIVE"
        );

        // Ensure we aren't proving a validator twice for the same checkpoint. This will fail if:
        // - validator submitted twice during this checkpoint
        // - validator withdrawal credentials verified after checkpoint starts, then submitted
        //   as a checkpoint proof
        // TODO - we might want to "skip" and emit an event, rather than revert?
        require(
            validatorInfo.mostRecentBalanceUpdateTimestamp < beaconTimestamp,
            "EigenPod._verifyCheckpointProof: validator already proven for this checkpoint"
        );
        
        // Verify validator balance against beaconStateRoot
        uint64 prevBalanceGwei = validatorInfo.restakedBalanceGwei;
        uint64 newBalanceGwei = BeaconChainProofs.verifyValidatorBalance({
            beaconStateRoot: beaconStateRoot,
            validatorIndex: validatorInfo.validatorIndex,
            proof: proof
        });
        
        // Update validator state. If their new balance is 0, they are marked `WITHDRAWN`
        validatorInfo.restakedBalanceGwei = newBalanceGwei;
        validatorInfo.mostRecentBalanceUpdateTimestamp = beaconTimestamp;
        if (newBalanceGwei == 0) {
            activeValidatorCount--;
            validatorInfo.status = VALIDATOR_STATUS.WITHDRAWN;
        }

        _validatorPubkeyHashToInfo[proof.pubkeyHash] = validatorInfo;

        // Calculate change in the validator's balance since the last proof
        if (newBalanceGwei != prevBalanceGwei) {
            emit ValidatorBalanceUpdated(validatorIndex, beaconTimestamp, newBalanceGwei);

            balanceDeltaWei = _calcBalanceDeltaWei({
                newAmountGwei: newBalanceGwei,
                previousAmountGwei: prevBalanceGwei
            });
        }

        // If the validator was marked stale and this proof is younger than
        // `TIME_TILL_STALE_BALANCE`, the validator is no longer stale.
        if (
            beaconTimestamp + TIME_TILL_STALE_BALANCE >= block.timestamp && 
            isValidatorStale[proof.pubkeyHash]
        ) {
            isValidatorStale[proof.pubkeyHash] = false;
            noLongerStale = true;
        }

        return (balanceDeltaWei, noLongerStale);
    }

    /**
     * @dev Finish progress on a checkpoint and store it in state.
     * @dev If the checkpoint has no proofs remaining, it is finalized:
     * - a share delta is calculated and sent to the `EigenPodManager`
     * - the checkpointed `podBalance` is added to `withdrawableRestakedExecutionLayerGwei`
     * - `lastFinalizedCheckpoint` is updated
     * - `currentCheckpoint` and `currentCheckpointTimestamp` are deleted
     */
    function _updateCheckpoint(Checkpoint memory checkpoint) internal {
        if (checkpoint.proofsRemaining == 0) {
            int256 totalShareDelta =
                int256(checkpoint.podBalance) + checkpoint.balanceDeltas;

            // Add any native ETH in the pod to `withdrawableRestakedExecutionLayerGwei`
            // ... this amount can be withdrawn via the `DelegationManager` withdrawal queue
            uint64 podBalanceGwei = checkpoint.podBalanceWei / GWEI_TO_WEI;
            withdrawableRestakedExecutionLayerGwei += podBalanceGwei;

            // Finalize the checkpoint
            lastFinalizedCheckpoint = currentCheckpointTimestamp;
            delete currentCheckpointTimestamp;
            delete currentCheckpoint;

            // Update pod owner's shares
            eigenPodManager.recordBeaconChainETHBalanceUpdate(podOwner, totalShareDelta);
        } else {
            currentCheckpoint = checkpoint;
        }
    }

    function _processWithdrawalBeforeRestaking(address _podOwner) internal {
        mostRecentWithdrawalTimestamp = uint32(block.timestamp);
        nonBeaconChainETHBalanceWei = 0;
        _sendETH_AsDelayedWithdrawal(_podOwner, address(this).balance);
    }

    function _sendETH(address recipient, uint256 amountWei) internal {
        Address.sendValue(payable(recipient), amountWei);
    }

    function _sendETH_AsDelayedWithdrawal(address recipient, uint256 amountWei) internal {
        delayedWithdrawalRouter.createDelayedWithdrawal{value: amountWei}(podOwner, recipient);
    }

    function _podWithdrawalCredentials() internal view returns (bytes memory) {
        return abi.encodePacked(bytes1(uint8(1)), bytes11(0), address(this));
    }

    ///@notice Calculates the pubkey hash of a validator's pubkey as per SSZ spec
    function _calculateValidatorPubkeyHash(bytes memory validatorPubkey) internal pure returns (bytes32){
        require(validatorPubkey.length == 48, "EigenPod._calculateValidatorPubkeyHash must be a 48-byte BLS public key");
        return sha256(abi.encodePacked(validatorPubkey, bytes16(0)));
    }

    /// @dev Calculates the delta between two Gwei amounts, converts to Wei, and returns as an int256
    function _calcBalanceDeltaWei(uint64 newAmountGwei, uint64 previousAmountGwei) internal pure returns (int256) {
        return
            (int256(uint256(newAmountGwei)) - int256(uint256(previousAmountGwei))) * int256(GWEI_TO_WEI);
    }

    /*******************************************************************************
                            VIEW FUNCTIONS
    *******************************************************************************/

    function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) external view returns (ValidatorInfo memory) {
        return _validatorPubkeyHashToInfo[validatorPubkeyHash];
    }

    /// @notice Returns the validatorInfo for a given validatorPubkey
    function validatorPubkeyToInfo(bytes calldata validatorPubkey) external view returns (ValidatorInfo memory) {
        return _validatorPubkeyHashToInfo[_calculateValidatorPubkeyHash(validatorPubkey)];
    }

    function validatorStatus(bytes32 pubkeyHash) external view returns (VALIDATOR_STATUS) {
        return _validatorPubkeyHashToInfo[pubkeyHash].status;
    }

        /// @notice Returns the validator status for a given validatorPubkey
    function validatorStatus(bytes calldata validatorPubkey) external view returns (VALIDATOR_STATUS) {
        bytes32 validatorPubkeyHash = _calculateValidatorPubkeyHash(validatorPubkey);
        return _validatorPubkeyHashToInfo[validatorPubkeyHash].status;
    }
}
