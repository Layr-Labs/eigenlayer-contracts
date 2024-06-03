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
contract EigenPod is 
    Initializable, 
    ReentrancyGuardUpgradeable, 
    EigenPodPausingConstants, 
    EigenPodStorage 
{

    using BytesLib for bytes;
    using SafeERC20 for IERC20;
    using BeaconChainProofs for *;

    /*******************************************************************************
                               CONSTANTS / IMMUTABLES
    *******************************************************************************/

    /// @notice The beacon chain stores balances in Gwei, rather than wei. This value is used to convert between the two
    uint256 internal constant GWEI_TO_WEI = 1e9;

    /// @notice If a validator is slashed on the beacon chain and their balance has not been checkpointed
    /// within `TIME_TILL_STALE_BALANCE` of the current block, `verifyStaleBalance` allows anyone to start
    /// a checkpoint for the pod.
    uint256 internal constant TIME_TILL_STALE_BALANCE = 2 weeks;

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
        IEigenPodManager _eigenPodManager,
        uint64 _GENESIS_TIME
    ) {
        ethPOS = _ethPOS;
        eigenPodManager = _eigenPodManager;
        GENESIS_TIME = _GENESIS_TIME;
        _disableInitializers();
    }

    /// @notice Used to initialize the pointers to addresses crucial to the pod's functionality. Called on construction by the EigenPodManager.
    function initialize(address _podOwner) external initializer {
        require(_podOwner != address(0), "EigenPod.initialize: podOwner cannot be zero address");
        podOwner = _podOwner;

        /// Pods deployed prior to the M2 release have this variable set to `false`, as before M2, pod owners
        /// did not need to engage with the proof system and were able to withdraw ETH from their pod on demand.
        hasRestaked = true;
        emit RestakingActivated(podOwner);
    }

    /*******************************************************************************
                                    EXTERNAL METHODS
    *******************************************************************************/

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
    function startCheckpoint(bool revertIfNoBalance)
        external
        onlyEigenPodOwner() 
        onlyWhenNotPaused(PAUSED_START_CHECKPOINT) 
    {
        _startCheckpoint(revertIfNoBalance);

        /// Legacy support for pods deployed pre-M2 that never activated restaking. `startCheckpoint`
        /// can activate restaking, allowing them to continue using their pods via the checkpoint
        /// system.
        if (!hasRestaked) {
            hasRestaked = true;
            emit RestakingActivated(podOwner);
        }
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
    ) 
        external 
        onlyWhenNotPaused(PAUSED_EIGENPODS_VERIFY_CHECKPOINT_PROOFS) 
    {
        uint64 beaconTimestamp = currentCheckpointTimestamp;
        require(
            beaconTimestamp != 0, 
            "EigenPod.verifyCheckpointProofs: must have active checkpoint to perform checkpoint proof"
        );

        Checkpoint memory checkpoint = _currentCheckpoint;

        // Verify `balanceContainerProof` against `beaconBlockRoot`
        BeaconChainProofs.verifyBalanceContainer({
            beaconBlockRoot: checkpoint.beaconBlockRoot,
            proof: balanceContainerProof
        });

        // Process each checkpoint proof submitted
        for (uint256 i = 0; i < proofs.length; i++) {
            ValidatorInfo memory validatorInfo = _validatorPubkeyHashToInfo[proofs[i].pubkeyHash];

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
            if (validatorInfo.mostRecentBalanceUpdateTimestamp >= beaconTimestamp) {
                continue;
            }

            // Process a checkpoint proof for a validator. 
            // - the validator MUST be in the ACTIVE state
            // - the validator MUST NOT have already been proven for this checkpoint
            //
            // If the proof shows the validator has a balance of 0, they are marked `WITHDRAWN`.
            // The assumption is that if this is the case, any withdrawn ETH was already in
            // the pod when `startCheckpoint` was originally called.
            int128 balanceDeltaGwei = _verifyCheckpointProof({
                validatorInfo: validatorInfo,
                beaconTimestamp: beaconTimestamp,
                balanceContainerRoot: balanceContainerProof.balanceContainerRoot,
                proof: proofs[i]
            });

            checkpoint.proofsRemaining--;
            checkpoint.balanceDeltasGwei += balanceDeltaGwei;
        }

        // Update and/or finalize the checkpoint
        _updateCheckpoint(checkpoint);
    }

    /**
     * @dev Verify one or more validators have their withdrawal credentials pointed at this EigenPod, and award
     * shares based on their effective balance. Proven validators are marked `ACTIVE` within the EigenPod, and
     * future checkpoint proofs will need to include them.
     * @dev Withdrawal credential proofs MUST NOT be older than the `lastCheckpointTimestamp` OR `currentCheckpointTimestamp`.
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
    {
        require(hasRestaked, "EigenPod.verifyWithdrawalCredentials: restaking not active");

        require(
            (validatorIndices.length == validatorFieldsProofs.length) &&
                (validatorFieldsProofs.length == validatorFields.length),
            "EigenPod.verifyWithdrawalCredentials: validatorIndices and proofs must be same length"
        );
        
        require(
            beaconTimestamp > lastCheckpointTimestamp && beaconTimestamp > currentCheckpointTimestamp,
            "EigenPod.verifyWithdrawalCredentials: specified timestamp is too far in past"
        );

        // Verify passed-in `beaconStateRoot` against the beacon block root
        BeaconChainProofs.verifyStateRoot({
            beaconBlockRoot: _getParentBlockRoot(beaconTimestamp),
            proof: stateRootProof
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
     * checkpoint proofs since being slashed. If successful, this allows the caller to start a checkpoint.
     * @dev Note that in order to start a checkpoint, any existing checkpoint must already be completed!
     * (See `_startCheckpoint` for details)
     * @param beaconTimestamp the beacon chain timestamp sent to the 4788 oracle contract. Corresponds
     * to the parent beacon block root against which the proof is verified.
     * @param stateRootProof proves a beacon state root against a beacon block root
     * @param proof the fields of the beacon chain "Validator" container, along with a merkle proof against
     * the beacon state root. See the consensus specs for more details:
     * https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator
     *
     * @dev Staleness conditions:
     * - Validator's last balance update is older than `beaconTimestamp` by `TIME_TILL_STALE_BALANCE`
     * - Validator MUST be in `ACTIVE` status in the pod
     * - Validator MUST be slashed on the beacon chain
     */
    function verifyStaleBalance(
        uint64 beaconTimestamp,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        BeaconChainProofs.ValidatorProof calldata proof
    )
        external
        onlyWhenNotPaused(PAUSED_VERIFY_STALE_BALANCE)
    {  
        bytes32 validatorPubkey = proof.validatorFields.getPubkeyHash();
        ValidatorInfo memory validatorInfo = _validatorPubkeyHashToInfo[validatorPubkey];

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

        // Validator must be slashed on the beacon chain
        require(
            proof.validatorFields.isValidatorSlashed(),
            "EigenPod.verifyStaleBalance: validator must be slashed to be marked stale"
        );

        // Verify passed-in `beaconStateRoot` against the beacon block root
        BeaconChainProofs.verifyStateRoot({
            beaconBlockRoot: _getParentBlockRoot(beaconTimestamp),
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
        require(
            tokenList.length == amountsToWithdraw.length,
            "EigenPod.recoverTokens: tokenList and amountsToWithdraw must be same length"
        );
        for (uint256 i = 0; i < tokenList.length; i++) {
            tokenList[i].safeTransfer(recipient, amountsToWithdraw[i]);
        }
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
            validatorFields.getExitEpoch() == BeaconChainProofs.FAR_FUTURE_EPOCH,
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
        ValidatorInfo memory validatorInfo,
        uint64 beaconTimestamp,
        bytes32 balanceContainerRoot,
        BeaconChainProofs.BalanceProof calldata proof
    ) internal returns (int128 balanceDeltaGwei) {
        uint40 validatorIndex = uint40(validatorInfo.validatorIndex);
        
        // Verify validator balance against `balanceContainerRoot`
        uint64 prevBalanceGwei = validatorInfo.restakedBalanceGwei;
        uint64 newBalanceGwei = BeaconChainProofs.verifyValidatorBalance({
            balanceContainerRoot: balanceContainerRoot,
            validatorIndex: validatorIndex,
            proof: proof
        });
        
        // Update validator state. If their new balance is 0, they are marked `WITHDRAWN`
        validatorInfo.restakedBalanceGwei = newBalanceGwei;
        validatorInfo.mostRecentBalanceUpdateTimestamp = beaconTimestamp;
        if (newBalanceGwei == 0) {
            activeValidatorCount--;
            validatorInfo.status = VALIDATOR_STATUS.WITHDRAWN;

            emit ValidatorWithdrawn(beaconTimestamp, validatorIndex);
        }

        _validatorPubkeyHashToInfo[proof.pubkeyHash] = validatorInfo;
        emit ValidatorCheckpointed(beaconTimestamp, validatorIndex);

        // Calculate change in the validator's balance since the last proof
        if (newBalanceGwei != prevBalanceGwei) {
            emit ValidatorBalanceUpdated(validatorIndex, beaconTimestamp, newBalanceGwei);

            balanceDeltaGwei = _calcBalanceDelta({
                newAmountGwei: newBalanceGwei,
                previousAmountGwei: prevBalanceGwei
            });
        }

        return balanceDeltaGwei;
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
    function _startCheckpoint(bool revertIfNoBalance) internal {
        require(
            currentCheckpointTimestamp == 0, 
            "EigenPod._startCheckpoint: must finish previous checkpoint before starting another"
        );

        // Snapshot pod balance at the start of the checkpoint, subtracting pod balance that has
        // previously been credited with shares. Once the checkpoint is finalized, `podBalanceGwei` 
        // will be added to the total validator balance delta and credited as shares.
        // 
        // Note: On finalization, `podBalanceGwei` is added to `withdrawableRestakedExecutionLayerGwei`
        // to denote that it has been credited with shares. Because this value is denominated in gwei, 
        // `podBalanceGwei` is also converted to a gwei amount here. This means that any sub-gwei amounts 
        // sent to the pod are not credited with shares and are therefore not withdrawable. 
        // This can be addressed by topping up a pod's balance to a value divisible by 1 gwei.
        uint64 podBalanceGwei = 
            uint64(address(this).balance / GWEI_TO_WEI) - withdrawableRestakedExecutionLayerGwei;

        // If the caller doesn't want a "0 balance" checkpoint, revert
        if (revertIfNoBalance && podBalanceGwei == 0) {
            revert("EigenPod._startCheckpoint: no balance available to checkpoint");
        }

        // Create checkpoint using the previous block's root for proofs, and the current
        // `activeValidatorCount` as the number of checkpoint proofs needed to finalize
        // the checkpoint.
        Checkpoint memory checkpoint = Checkpoint({
            beaconBlockRoot: _getParentBlockRoot(uint64(block.timestamp)),
            proofsRemaining: uint24(activeValidatorCount),
            podBalanceGwei: podBalanceGwei,
            balanceDeltasGwei: 0
        });

        // Place checkpoint in storage. If `proofsRemaining` is 0, the checkpoint
        // is automatically finalized.
        currentCheckpointTimestamp = uint64(block.timestamp);
        _updateCheckpoint(checkpoint);

        emit CheckpointCreated(uint64(block.timestamp), checkpoint.beaconBlockRoot);
    }

    /**
     * @dev Finish progress on a checkpoint and store it in state.
     * @dev If the checkpoint has no proofs remaining, it is finalized:
     * - a share delta is calculated and sent to the `EigenPodManager`
     * - the checkpointed `podBalanceGwei` is added to `withdrawableRestakedExecutionLayerGwei`
     * - `lastCheckpointTimestamp` is updated
     * - `_currentCheckpoint` and `currentCheckpointTimestamp` are deleted
     */
    function _updateCheckpoint(Checkpoint memory checkpoint) internal {
        if (checkpoint.proofsRemaining == 0) {
            int256 totalShareDeltaWei =
                (int128(uint128(checkpoint.podBalanceGwei)) + checkpoint.balanceDeltasGwei) * int256(GWEI_TO_WEI);

            // Add any native ETH in the pod to `withdrawableRestakedExecutionLayerGwei`
            // ... this amount can be withdrawn via the `DelegationManager` withdrawal queue
            withdrawableRestakedExecutionLayerGwei += checkpoint.podBalanceGwei;

            // Finalize the checkpoint
            lastCheckpointTimestamp = currentCheckpointTimestamp;
            delete currentCheckpointTimestamp;
            delete _currentCheckpoint;

            // Update pod owner's shares
            eigenPodManager.recordBeaconChainETHBalanceUpdate(podOwner, totalShareDeltaWei);
            emit CheckpointFinalized(lastCheckpointTimestamp, totalShareDeltaWei);
        } else {
            _currentCheckpoint = checkpoint;
        }
    }

    function _sendETH(address recipient, uint256 amountWei) internal {
        Address.sendValue(payable(recipient), amountWei);
    }

    /// @notice Query the 4788 oracle to get the parent block root of the slot with the given `timestamp`
    /// @param timestamp of the block for which the parent block root will be returned. MUST correspond
    /// to an existing slot within the last 24 hours. If the slot at `timestamp` was skipped, this method
    /// will revert.
    function _getParentBlockRoot(uint64 timestamp) internal view returns (bytes32) {
        require(
            block.timestamp - timestamp < BEACON_ROOTS_HISTORY_BUFFER_LENGTH * 12,
            "EigenPod._getParentBlockRoot: timestamp out of range"
        );

        (bool success, bytes memory result) =
            BEACON_ROOTS_ADDRESS.staticcall(abi.encode(timestamp));

        if (success && result.length > 0) {
            return abi.decode(result, (bytes32));
        } else {
            revert("EigenPod._getParentBlockRoot: invalid block root returned");
        }
    }

    function _podWithdrawalCredentials() internal view returns (bytes memory) {
        return abi.encodePacked(bytes1(uint8(1)), bytes11(0), address(this));
    }

    ///@notice Calculates the pubkey hash of a validator's pubkey as per SSZ spec
    function _calculateValidatorPubkeyHash(bytes memory validatorPubkey) internal pure returns (bytes32){
        require(validatorPubkey.length == 48, "EigenPod._calculateValidatorPubkeyHash must be a 48-byte BLS public key");
        return sha256(abi.encodePacked(validatorPubkey, bytes16(0)));
    }

    /// @dev Calculates the delta between two Gwei amounts and returns as an int256
    function _calcBalanceDelta(uint64 newAmountGwei, uint64 previousAmountGwei) internal pure returns (int128) {
        return
            int128(uint128(newAmountGwei)) - int128(uint128(previousAmountGwei));
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

    /// @notice Returns the currently-active checkpoint
    function currentCheckpoint() public view returns (Checkpoint memory) {
        return _currentCheckpoint;
    }
}
