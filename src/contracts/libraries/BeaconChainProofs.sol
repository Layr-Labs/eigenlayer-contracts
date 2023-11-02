// SPDX-License-Identifier: BUSL-1.1

pragma solidity ^0.8.0;

import "./Merkle.sol";
import "../libraries/Endian.sol";

//Utility library for parsing and PHASE0 beacon chain block headers
//SSZ Spec: https://github.com/ethereum/consensus-specs/blob/dev/ssz/simple-serialize.md#merkleization
//BeaconBlockHeader Spec: https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#beaconblockheader
//BeaconState Spec: https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#beaconstate
library BeaconChainProofs {
    // constants are the number of fields and the heights of the different merkle trees used in merkleizing beacon chain containers
    uint256 internal constant NUM_BEACON_BLOCK_HEADER_FIELDS = 5;
    uint256 internal constant BEACON_BLOCK_HEADER_FIELD_TREE_HEIGHT = 3;

    uint256 internal constant NUM_BEACON_BLOCK_BODY_FIELDS = 11;
    uint256 internal constant BEACON_BLOCK_BODY_FIELD_TREE_HEIGHT = 4;

    uint256 internal constant NUM_BEACON_STATE_FIELDS = 21;
    uint256 internal constant BEACON_STATE_FIELD_TREE_HEIGHT = 5;

    uint256 internal constant NUM_ETH1_DATA_FIELDS = 3;
    uint256 internal constant ETH1_DATA_FIELD_TREE_HEIGHT = 2;

    uint256 internal constant NUM_VALIDATOR_FIELDS = 8;
    uint256 internal constant VALIDATOR_FIELD_TREE_HEIGHT = 3;

    uint256 internal constant NUM_EXECUTION_PAYLOAD_HEADER_FIELDS = 15;
    uint256 internal constant EXECUTION_PAYLOAD_HEADER_FIELD_TREE_HEIGHT = 4;

    uint256 internal constant NUM_EXECUTION_PAYLOAD_FIELDS = 15;
    uint256 internal constant EXECUTION_PAYLOAD_FIELD_TREE_HEIGHT = 4;

    // HISTORICAL_ROOTS_LIMIT	 = 2**24, so tree height is 24
    uint256 internal constant HISTORICAL_ROOTS_TREE_HEIGHT = 24;

    // HISTORICAL_BATCH is root of state_roots and block_root, so number of leaves =  2^1
    uint256 internal constant HISTORICAL_BATCH_TREE_HEIGHT = 1;

    // SLOTS_PER_HISTORICAL_ROOT = 2**13, so tree height is 13
    uint256 internal constant STATE_ROOTS_TREE_HEIGHT = 13;
    uint256 internal constant BLOCK_ROOTS_TREE_HEIGHT = 13;

    //HISTORICAL_ROOTS_LIMIT = 2**24, so tree height is 24
    uint256 internal constant HISTORICAL_SUMMARIES_TREE_HEIGHT = 24;

    //Index of block_summary_root in historical_summary container
    uint256 internal constant BLOCK_SUMMARY_ROOT_INDEX = 0;

    uint256 internal constant NUM_WITHDRAWAL_FIELDS = 4;
    // tree height for hash tree of an individual withdrawal container
    uint256 internal constant WITHDRAWAL_FIELD_TREE_HEIGHT = 2;

    uint256 internal constant VALIDATOR_TREE_HEIGHT = 40;

    // MAX_WITHDRAWALS_PER_PAYLOAD = 2**4, making tree height = 4
    uint256 internal constant WITHDRAWALS_TREE_HEIGHT = 4;

    //in beacon block body https://github.com/ethereum/consensus-specs/blob/dev/specs/capella/beacon-chain.md#beaconblockbody
    uint256 internal constant EXECUTION_PAYLOAD_INDEX = 9;

    // in beacon block header https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#beaconblockheader
    uint256 internal constant SLOT_INDEX = 0;
    uint256 internal constant PROPOSER_INDEX_INDEX = 1;
    uint256 internal constant STATE_ROOT_INDEX = 3;
    uint256 internal constant BODY_ROOT_INDEX = 4;
    // in beacon state https://github.com/ethereum/consensus-specs/blob/dev/specs/capella/beacon-chain.md#beaconstate
    uint256 internal constant HISTORICAL_BATCH_STATE_ROOT_INDEX = 1;
    uint256 internal constant BEACON_STATE_SLOT_INDEX = 2;
    uint256 internal constant LATEST_BLOCK_HEADER_ROOT_INDEX = 4;
    uint256 internal constant BLOCK_ROOTS_INDEX = 5;
    uint256 internal constant STATE_ROOTS_INDEX = 6;
    uint256 internal constant HISTORICAL_ROOTS_INDEX = 7;
    uint256 internal constant ETH_1_ROOT_INDEX = 8;
    uint256 internal constant VALIDATOR_TREE_ROOT_INDEX = 11;
    uint256 internal constant EXECUTION_PAYLOAD_HEADER_INDEX = 24;
    uint256 internal constant HISTORICAL_SUMMARIES_INDEX = 27;

    // in validator https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator
    uint256 internal constant VALIDATOR_PUBKEY_INDEX = 0;
    uint256 internal constant VALIDATOR_WITHDRAWAL_CREDENTIALS_INDEX = 1;
    uint256 internal constant VALIDATOR_BALANCE_INDEX = 2;
    uint256 internal constant VALIDATOR_SLASHED_INDEX = 3;
    uint256 internal constant VALIDATOR_WITHDRAWABLE_EPOCH_INDEX = 7;

    // in execution payload header
    uint256 internal constant TIMESTAMP_INDEX = 9;
    uint256 internal constant WITHDRAWALS_ROOT_INDEX = 14;

    //in execution payload
    uint256 internal constant WITHDRAWALS_INDEX = 14;

    // in withdrawal
    uint256 internal constant WITHDRAWAL_VALIDATOR_INDEX_INDEX = 1;
    uint256 internal constant WITHDRAWAL_VALIDATOR_AMOUNT_INDEX = 3;

    //In historicalBatch
    uint256 internal constant HISTORICALBATCH_STATEROOTS_INDEX = 1;

    //Misc Constants

    /// @notice The number of slots each epoch in the beacon chain
    uint64 internal constant SLOTS_PER_EPOCH = 32;

    /// @notice The number of seconds in a slot in the beacon chain
    uint64 internal constant SECONDS_PER_SLOT = 12;

    /// @notice Number of seconds per epoch: 384 == 32 slots/epoch * 12 seconds/slot 
    uint64 internal constant SECONDS_PER_EPOCH = SLOTS_PER_EPOCH * SECONDS_PER_SLOT;

    bytes8 internal constant UINT64_MASK = 0xffffffffffffffff;

    /// @notice This struct contains the merkle proofs and leaves needed to verify a partial/full withdrawal
    struct WithdrawalProof {
        bytes withdrawalProof;
        bytes slotProof;
        bytes executionPayloadProof;
        bytes timestampProof;
        bytes historicalSummaryBlockRootProof;
        uint64 blockRootIndex;
        uint64 historicalSummaryIndex;
        uint64 withdrawalIndex;
        bytes32 blockRoot;
        bytes32 slotRoot;
        bytes32 timestampRoot;
        bytes32 executionPayloadRoot;
    }

    /// @notice This struct contains the root and proof for verifying the state root against the oracle block root
    struct StateRootProof {
        bytes32 beaconStateRoot;
        bytes proof;
    }

    /**
     * @notice This function verifies merkle proofs of the fields of a certain validator against a beacon chain state root
     * @param validatorIndex the index of the proven validator
     * @param beaconStateRoot is the beacon chain state root to be proven against.
     * @param validatorFieldsProof is the data used in proving the validator's fields
     * @param validatorFields the claimed fields of the validator
     */
    function verifyValidatorFields(
        bytes32 beaconStateRoot,
        bytes32[] calldata validatorFields,
        bytes calldata validatorFieldsProof,
        uint40 validatorIndex
    ) internal view {
        require(
            validatorFields.length == 2 ** VALIDATOR_FIELD_TREE_HEIGHT,
            "BeaconChainProofs.verifyValidatorFields: Validator fields has incorrect length"
        );

        /**
         * Note: the length of the validator merkle proof is BeaconChainProofs.VALIDATOR_TREE_HEIGHT + 1.
         * There is an additional layer added by hashing the root with the length of the validator list
         */
        require(
            validatorFieldsProof.length == 32 * ((VALIDATOR_TREE_HEIGHT + 1) + BEACON_STATE_FIELD_TREE_HEIGHT),
            "BeaconChainProofs.verifyValidatorFields: Proof has incorrect length"
        );
        uint256 index = (VALIDATOR_TREE_ROOT_INDEX << (VALIDATOR_TREE_HEIGHT + 1)) | uint256(validatorIndex);
        // merkleize the validatorFields to get the leaf to prove
        bytes32 validatorRoot = Merkle.merkleizeSha256(validatorFields);

        // verify the proof of the validatorRoot against the beaconStateRoot
        require(
            Merkle.verifyInclusionSha256({
                proof: validatorFieldsProof,
                root: beaconStateRoot,
                leaf: validatorRoot,
                index: index
            }),
            "BeaconChainProofs.verifyValidatorFields: Invalid merkle proof"
        );
    }

    /**
     * @notice This function verifies the latestBlockHeader against the state root. the latestBlockHeader is
     * a tracked in the beacon state.
     * @param beaconStateRoot is the beacon chain state root to be proven against.
     * @param stateRootProof is the provided merkle proof
     * @param latestBlockRoot is hashtree root of the latest block header in the beacon state
     */
    function verifyStateRootAgainstLatestBlockRoot(
        bytes32 latestBlockRoot,
        bytes32 beaconStateRoot,
        bytes calldata stateRootProof
    ) internal view {
        require(
            stateRootProof.length == 32 * (BEACON_BLOCK_HEADER_FIELD_TREE_HEIGHT),
            "BeaconChainProofs.verifyStateRootAgainstLatestBlockRoot: Proof has incorrect length"
        );
        //Next we verify the slot against the blockRoot
        require(
            Merkle.verifyInclusionSha256({
                proof: stateRootProof,
                root: latestBlockRoot,
                leaf: beaconStateRoot,
                index: STATE_ROOT_INDEX
            }),
            "BeaconChainProofs.verifyStateRootAgainstLatestBlockRoot: Invalid latest block header root merkle proof"
        );
    }

    /**
     * @notice This function verifies the slot and the withdrawal fields for a given withdrawal
     * @param withdrawalProof is the provided set of merkle proofs
     * @param withdrawalFields is the serialized withdrawal container to be proven
     */
    function verifyWithdrawal(
        bytes32 beaconStateRoot,
        bytes32[] calldata withdrawalFields,
        WithdrawalProof calldata withdrawalProof
    ) internal view {
        require(
            withdrawalFields.length == 2 ** WITHDRAWAL_FIELD_TREE_HEIGHT,
            "BeaconChainProofs.verifyWithdrawal: withdrawalFields has incorrect length"
        );

        require(
            withdrawalProof.blockRootIndex < 2 ** BLOCK_ROOTS_TREE_HEIGHT,
            "BeaconChainProofs.verifyWithdrawal: blockRootIndex is too large"
        );
        require(
            withdrawalProof.withdrawalIndex < 2 ** WITHDRAWALS_TREE_HEIGHT,
            "BeaconChainProofs.verifyWithdrawal: withdrawalIndex is too large"
        );

        require(
            withdrawalProof.historicalSummaryIndex < 2 ** HISTORICAL_SUMMARIES_TREE_HEIGHT,
            "BeaconChainProofs.verifyWithdrawal: historicalSummaryIndex is too large"
        );

        require(
            withdrawalProof.withdrawalProof.length ==
                32 * (EXECUTION_PAYLOAD_HEADER_FIELD_TREE_HEIGHT + WITHDRAWALS_TREE_HEIGHT + 1),
            "BeaconChainProofs.verifyWithdrawal: withdrawalProof has incorrect length"
        );
        require(
            withdrawalProof.executionPayloadProof.length ==
                32 * (BEACON_BLOCK_HEADER_FIELD_TREE_HEIGHT + BEACON_BLOCK_BODY_FIELD_TREE_HEIGHT),
            "BeaconChainProofs.verifyWithdrawal: executionPayloadProof has incorrect length"
        );
        require(
            withdrawalProof.slotProof.length == 32 * (BEACON_BLOCK_HEADER_FIELD_TREE_HEIGHT),
            "BeaconChainProofs.verifyWithdrawal: slotProof has incorrect length"
        );
        require(
            withdrawalProof.timestampProof.length == 32 * (EXECUTION_PAYLOAD_HEADER_FIELD_TREE_HEIGHT),
            "BeaconChainProofs.verifyWithdrawal: timestampProof has incorrect length"
        );

        require(
            withdrawalProof.historicalSummaryBlockRootProof.length ==
                32 *
                    (BEACON_STATE_FIELD_TREE_HEIGHT +
                        (HISTORICAL_SUMMARIES_TREE_HEIGHT + 1) +
                        1 +
                        (BLOCK_ROOTS_TREE_HEIGHT)),
            "BeaconChainProofs.verifyWithdrawal: historicalSummaryBlockRootProof has incorrect length"
        );
        /**
         * Note: Here, the "1" in "1 + (BLOCK_ROOTS_TREE_HEIGHT)" signifies that extra step of choosing the "block_root_summary" within the individual
         * "historical_summary". Everywhere else it signifies merkelize_with_mixin, where the length of an array is hashed with the root of the array,
         * but not here.
         */
        uint256 historicalBlockHeaderIndex = (HISTORICAL_SUMMARIES_INDEX <<
            ((HISTORICAL_SUMMARIES_TREE_HEIGHT + 1) + 1 + (BLOCK_ROOTS_TREE_HEIGHT))) |
            (uint256(withdrawalProof.historicalSummaryIndex) << (1 + (BLOCK_ROOTS_TREE_HEIGHT))) |
            (BLOCK_SUMMARY_ROOT_INDEX << (BLOCK_ROOTS_TREE_HEIGHT)) |
            uint256(withdrawalProof.blockRootIndex);

        require(
            Merkle.verifyInclusionSha256({
                proof: withdrawalProof.historicalSummaryBlockRootProof,
                root: beaconStateRoot,
                leaf: withdrawalProof.blockRoot,
                index: historicalBlockHeaderIndex
            }),
            "BeaconChainProofs.verifyWithdrawal: Invalid historicalsummary merkle proof"
        );

        //Next we verify the slot against the blockRoot
        require(
            Merkle.verifyInclusionSha256({
                proof: withdrawalProof.slotProof,
                root: withdrawalProof.blockRoot,
                leaf: withdrawalProof.slotRoot,
                index: SLOT_INDEX
            }),
            "BeaconChainProofs.verifyWithdrawal: Invalid slot merkle proof"
        );

        {
            // Next we verify the executionPayloadRoot against the blockRoot
            uint256 executionPayloadIndex = (BODY_ROOT_INDEX << (BEACON_BLOCK_BODY_FIELD_TREE_HEIGHT)) |
                EXECUTION_PAYLOAD_INDEX;
            require(
                Merkle.verifyInclusionSha256({
                    proof: withdrawalProof.executionPayloadProof,
                    root: withdrawalProof.blockRoot,
                    leaf: withdrawalProof.executionPayloadRoot,
                    index: executionPayloadIndex
                }),
                "BeaconChainProofs.verifyWithdrawal: Invalid executionPayload merkle proof"
            );
        }

        // Next we verify the timestampRoot against the executionPayload root
        require(
            Merkle.verifyInclusionSha256({
                proof: withdrawalProof.timestampProof,
                root: withdrawalProof.executionPayloadRoot,
                leaf: withdrawalProof.timestampRoot,
                index: TIMESTAMP_INDEX
            }),
            "BeaconChainProofs.verifyWithdrawal: Invalid blockNumber merkle proof"
        );

        {
            /**
             * Next we verify the withdrawal fields against the blockRoot:
             * First we compute the withdrawal_index relative to the blockRoot by concatenating the indexes of all the
             * intermediate root indexes from the bottom of the sub trees (the withdrawal container) to the top, the blockRoot.
             * Then we calculate merkleize the withdrawalFields container to calculate the the withdrawalRoot.
             * Finally we verify the withdrawalRoot against the executionPayloadRoot.
             *
             *
             * Note: Merkleization of the withdrawals root tree uses MerkleizeWithMixin, i.e., the length of the array is hashed with the root of
             * the array.  Thus we shift the WITHDRAWALS_INDEX over by WITHDRAWALS_TREE_HEIGHT + 1 and not just WITHDRAWALS_TREE_HEIGHT.
             */
            uint256 withdrawalIndex = (WITHDRAWALS_INDEX << (WITHDRAWALS_TREE_HEIGHT + 1)) |
                uint256(withdrawalProof.withdrawalIndex);
            bytes32 withdrawalRoot = Merkle.merkleizeSha256(withdrawalFields);
            require(
                Merkle.verifyInclusionSha256({
                    proof: withdrawalProof.withdrawalProof,
                    root: withdrawalProof.executionPayloadRoot,
                    leaf: withdrawalRoot,
                    index: withdrawalIndex
                }),
                "BeaconChainProofs.verifyWithdrawal: Invalid withdrawal merkle proof"
            );
        }
    }

    /**
     * @notice This function replicates the ssz hashing of a validator's pubkey, outlined below:
     *  hh := ssz.NewHasher()
     *  hh.PutBytes(validatorPubkey[:])
     *  validatorPubkeyHash := hh.Hash()
     *  hh.Reset()
     */
    function hashValidatorBLSPubkey(bytes memory validatorPubkey) internal pure returns (bytes32 pubkeyHash) {
        require(validatorPubkey.length == 48, "Input should be 48 bytes in length");
        return sha256(abi.encodePacked(validatorPubkey, bytes16(0)));
    }

    /**
     * @dev Retrieve the withdrawal timestamp
     */
    function getWithdrawalTimestamp(WithdrawalProof memory withdrawalProof) internal pure returns (uint64) {
        return
            Endian.fromLittleEndianUint64(withdrawalProof.timestampRoot);
    }

    /**
     * @dev Converts the withdrawal's slot to an epoch
     */
    function getWithdrawalEpoch(WithdrawalProof memory withdrawalProof) internal pure returns (uint64) {
        return
            Endian.fromLittleEndianUint64(withdrawalProof.slotRoot) / SLOTS_PER_EPOCH;
    }

    /**
     * Indices for validator fields (refer to consensus specs):
     * 0: pubkey
     * 1: withdrawal credentials
     * 2: effective balance
     * 3: slashed?
     * 4: activation elligibility epoch
     * 5: activation epoch
     * 6: exit epoch
     * 7: withdrawable epoch
     */

    /**
     * @dev Retrieves a validator's pubkey hash
     */
    function getPubkeyHash(bytes32[] memory validatorFields) internal pure returns (bytes32) {
        return 
            validatorFields[VALIDATOR_PUBKEY_INDEX];
    }

    function getWithdrawalCredentials(bytes32[] memory validatorFields) internal pure returns (bytes32) {
        return
            validatorFields[VALIDATOR_WITHDRAWAL_CREDENTIALS_INDEX];
    }

    /**
     * @dev Retrieves a validator's effective balance (in gwei)
     */
    function getEffectiveBalanceGwei(bytes32[] memory validatorFields) internal pure returns (uint64) {
        return 
            Endian.fromLittleEndianUint64(validatorFields[VALIDATOR_BALANCE_INDEX]);
    }

    /**
     * @dev Retrieves a validator's withdrawable epoch
     */
    function getWithdrawableEpoch(bytes32[] memory validatorFields) internal pure returns (uint64) {
        return 
            Endian.fromLittleEndianUint64(validatorFields[VALIDATOR_WITHDRAWABLE_EPOCH_INDEX]);
    }

    /**
     * Indices for withdrawal fields (refer to consensus specs):
     * 0: withdrawal index
     * 1: validator index
     * 2: execution address
     * 3: withdrawal amount
     */

    /**
     * @dev Retrieves a withdrawal's validator index
     */
    function getValidatorIndex(bytes32[] memory withdrawalFields) internal pure returns (uint40) {
        return 
            uint40(Endian.fromLittleEndianUint64(withdrawalFields[WITHDRAWAL_VALIDATOR_INDEX_INDEX]));
    }

    /**
     * @dev Retrieves a withdrawal's withdrawal amount (in gwei)
     */
    function getWithdrawalAmountGwei(bytes32[] memory withdrawalFields) internal pure returns (uint64) {
        return
            Endian.fromLittleEndianUint64(withdrawalFields[WITHDRAWAL_VALIDATOR_AMOUNT_INDEX]);
    }
}
