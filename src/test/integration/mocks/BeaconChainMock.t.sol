// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";

import "src/contracts/libraries/BeaconChainProofs.sol";
import "src/contracts/libraries/Merkle.sol";
import "src/contracts/pods/EigenPodManager.sol";

import "src/test/integration/TimeMachine.t.sol";
import "src/test/integration/mocks/BeaconChainOracleMock.t.sol";

struct CredentialsProofs {
    uint64 oracleTimestamp;
    BeaconChainProofs.StateRootProof stateRootProof;
    uint40[] validatorIndices;
    bytes[] validatorFieldsProofs;
    bytes32[][] validatorFields;
}

struct ValidatorFieldsProof {
    uint40 validatorIndex;
    bytes32[] validatorFields;
    bytes validatorFieldsProof;
}

// struct BeaconWithdrawal {
//     uint64 oracleTimestamp;
//     BeaconChainProofs.StateRootProof stateRootProof;
//     BeaconChainProofs.WithdrawalProof[] withdrawalProofs;
//     bytes[] validatorFieldsProofs;
//     bytes32[][] validatorFields;
//     bytes32[][] withdrawalFields;
// }

// struct BalanceUpdate {
//     uint64 oracleTimestamp;
//     BeaconChainProofs.StateRootProof stateRootProof;
//     uint40[] validatorIndices;
//     bytes[] validatorFieldsProofs;
//     bytes32[][] validatorFields;
// }

contract BeaconChainMock is Test {

    Vm cheats = Vm(HEVM_ADDRESS);

    struct Validator {
        uint40 validatorIndex;
        bytes32 pubkeyHash;
        bytes withdrawalCreds;
        uint64 effectiveBalanceGwei;
        uint64 exitEpoch;
    }

    /// @dev All withdrawals are processed with index == 0
    uint64 constant WITHDRAWAL_INDEX = 0;
    uint constant GWEI_TO_WEI = 1e9;
    uint constant ZERO_NODES_LENGTH = 50;
    
    uint64 public nextTimestamp;

    BeaconChainOracleMock oracle;
    EigenPodManager eigenPodManager;
    
    /**
     * BeaconState containers, used for proofgen:
     * https://eth2book.info/capella/part3/containers/state/#beaconstate 
     */

    // Validator container, references every validator created so far
    Validator[] validators;

    // Current balance container, keeps a balance for every validator
    //
    // Since balances are stored in groups of 4, it's easier to make
    // this a mapping rather than an array. If it were an array, its
    // length would be validators.length / 4;
    mapping(uint40 => bytes32) balances;

    /**
     * Generated proofs for each block timestamp:
     */

    // Maps block.timestamp -> calculated beacon block roots
    mapping(uint64 => bytes32) beaconBlockRoots;

    // Maps block.timestamp -> beacon state root and proof
    mapping(uint64 => BeaconChainProofs.StateRootProof) public stateRootProofs;

    // Maps block.timestamp -> validator field proofs for that timestamp
    mapping(uint64 => mapping(uint40 => ValidatorFieldsProof)) public validatorFieldsProofs;
    
    bytes32[] zeroNodes;
    
    constructor(TimeMachine timeMachine, BeaconChainOracleMock beaconChainOracle, EigenPodManager _eigenPodManager) {
        nextTimestamp = timeMachine.proofGenStartTime();
        oracle = beaconChainOracle;
        eigenPodManager = _eigenPodManager;

        // Calculate nodes of empty merkle tree
        bytes32 curNode = Merkle.merkleizeSha256(new bytes32[](8));
        zeroNodes = new bytes32[](ZERO_NODES_LENGTH);
        zeroNodes[0] = curNode;

        for (uint i = 1; i < zeroNodes.length; i++) {
            zeroNodes[i] = sha256(abi.encodePacked(curNode, curNode));
            curNode = zeroNodes[i];
        }
    }
    
    /**
     * @dev Processes a deposit for a new validator and returns the
     * information needed to prove withdrawal credentials.
     * 
     * For now, this returns empty proofs that will pass in the oracle,
     * but in the future this should use FFI to return a valid proof.
     */
    function newValidator(
        uint balanceWei,
        bytes memory withdrawalCreds
    ) public returns (uint40) {
        _log("newValidator");

        // These checks mimic the checks made in the beacon chain deposit contract
        //
        // We sanity-check them here because this contract sorta acts like the 
        // deposit contract and this ensures we only create validators that could
        // exist IRL
        require(balanceWei >= 1 ether, "BeaconChainMock.newValidator: deposit value too low");
        require(balanceWei % 1 gwei == 0, "BeaconChainMock.newValidator: value not multiple of gwei");
        uint depositAmount = balanceWei / GWEI_TO_WEI;
        require(depositAmount <= type(uint64).max, "BeaconChainMock.newValidator: deposit value too high");

        // Create new validator with unique index
        uint40 validatorIndex = uint40(validators.length);
        Validator memory v = Validator({
            validatorIndex: validatorIndex,
            pubkeyHash: keccak256(abi.encodePacked(validatorIndex)),
            withdrawalCreds: withdrawalCreds,
            effectiveBalanceGwei: uint64(depositAmount),
            exitEpoch: BeaconChainProofs.FAR_FUTURE_EPOCH
        });

        // Encode current balance and update balance root
        uint40 balanceRootIndex = validatorIndex / 4;
        bytes32 balanceRoot = balances[balanceRootIndex];
        balanceRoot = _setBalanceAtIndex(balanceRoot, validatorIndex, v.effectiveBalanceGwei);

        // Record in state
        validators.push(v);
        balances[balanceRootIndex] = balanceRoot;

        return validatorIndex;
    }

    function updateValidator(uint40 validatorIndex, uint addBalanceWei) public {
        uint amount = addBalanceWei / GWEI_TO_WEI;
        validators[validatorIndex].effectiveBalanceGwei += uint64(amount);
    }

    struct MerkleTree {
        uint numPopulated;
        bytes32[] leaves;
    }

    mapping(uint64 => mapping(bytes32 => bytes32)) siblings;
    mapping(uint64 => mapping(bytes32 => bytes32)) parents;

    // NOTE: currently only does validator container; ignores balances
    function processEpoch() public returns (ValidatorFieldsProof[] memory) {
        _log("processEpoch");

        uint64 timestamp = uint64(block.timestamp);
        cheats.warp(block.timestamp + 12);
        uint numValidators = validators.length;
        
        ValidatorFieldsProof[] memory vfProofs = new ValidatorFieldsProof[](numValidators);
        MerkleTree memory tree = MerkleTree({
            numPopulated: numValidators,
            leaves: new bytes32[](numValidators)
        });

        emit log_named_uint("populated leaves", tree.leaves.length);

        // Calculate validator fields and roots for each validator
        for (uint i = 0; i < numValidators; i++) {
            bytes32[] memory vFields = _getValidatorFields(uint40(i));
            vfProofs[i].validatorIndex = uint40(i);
            vfProofs[i].validatorFields = vFields;
            tree.leaves[i] = Merkle.merkleizeSha256(vFields);
        }

        // Calculate validators tree
        uint depth = 0;
        for (; depth < BeaconChainProofs.VALIDATOR_TREE_HEIGHT; depth++) {
            uint newLength = (tree.leaves.length + 1) / 2;
            bytes32[] memory newLeaves = new bytes32[](newLength);

            // Hash each pair of nodes in this level of the tree
            for (uint i = 0; i < newLength; i++) {
                uint leftIdx = 2 * i;
                uint rightIdx = leftIdx + 1;

                // Get left leaf
                bytes32 leftLeaf = tree.leaves[leftIdx];

                // Calculate right leaf
                bytes32 rightLeaf;
                if (rightIdx < tree.leaves.length) {
                    rightLeaf = tree.leaves[rightIdx];
                } else {
                    rightLeaf = _getZeroNode(depth);
                }

                // Hash left and right
                bytes32 result = sha256(abi.encodePacked(leftLeaf, rightLeaf));
                newLeaves[i] = result;

                // emit log_named_uint("depth", depth);
                // emit log_named_bytes32("left", leftLeaf);
                // emit log_named_bytes32("right", rightLeaf);

                // Record results, used to generate individual proofs later:
                // Record left and right as siblings
                siblings[timestamp][leftLeaf] = rightLeaf;
                siblings[timestamp][rightLeaf] = leftLeaf;
                // Record the result as the parent of left and right
                parents[timestamp][leftLeaf] = result;
                parents[timestamp][rightLeaf] = result;
            }

            // Move up one level
            tree.leaves = newLeaves;
        }

        emit log_named_uint("validator count", numValidators);
        emit log_named_uint("tree depth", depth);
        emit log_named_bytes32("validator container root", tree.leaves[0]);

        // Generate a proof for each validator
        for (uint i = 0; i < numValidators; i++) {

            bytes memory proof = new bytes(VAL_FIELDS_PROOF_LEN);
            bytes32 curNode = Merkle.merkleizeSha256(vfProofs[i].validatorFields);

            for (uint j = 0; j < depth; j++) {
                bytes32 sibling = siblings[timestamp][curNode];

                // proof[j] = sibling;
                assembly {
                    mstore(
                        add(proof, add(32, mul(32, j))),
                        sibling
                    )
                }

                if (sibling == 0) {
                    emit log_named_uint("err generating proof for validator", i);
                    emit log_named_bytes32("found null sibling for leaf", curNode);
                    emit log_named_uint("at depth", j);
                }

                curNode = parents[timestamp][curNode];
            }

            vfProofs[i].validatorFieldsProof = proof;

            emit log("===");
            emit log_named_uint("proof for validator", i);
            for (uint j = 32; j < proof.length; j+=32) {
                bytes32 element;
                assembly { element := mload(add(j, proof)) }
                emit log_bytes32(element);
            }
        }

        return vfProofs;
    }

    function _getValidatorFields(uint40 validatorIndex) internal view returns (bytes32[] memory) {
        bytes32[] memory vFields = new bytes32[](8);
        Validator memory v = validators[validatorIndex];

        vFields[BeaconChainProofs.VALIDATOR_PUBKEY_INDEX] = v.pubkeyHash;
        vFields[BeaconChainProofs.VALIDATOR_WITHDRAWAL_CREDENTIALS_INDEX] = bytes32(v.withdrawalCreds);
        vFields[BeaconChainProofs.VALIDATOR_BALANCE_INDEX] = _toLittleEndianUint64(v.effectiveBalanceGwei);
        vFields[BeaconChainProofs.VALIDATOR_EXIT_EPOCH_INDEX] = bytes32(uint(BeaconChainProofs.FAR_FUTURE_EPOCH));

        return vFields;
    }

    function getBalanceRoot(uint40 validatorIndex) public view returns (bytes32) {
        return balances[validatorIndex / 4];
    }

    function getBalance(uint40 validatorIndex) public view returns (uint64) {
        return BeaconChainProofs.getBalanceAtIndex(
            getBalanceRoot(validatorIndex),
            validatorIndex
        );
    }

    function _log(string memory s) internal {
        emit log_named_string("- beaconChain", s);
    }

    /**
     * @dev Exit a validator from the beacon chain, given its validatorIndex
     * The passed-in validatorIndex should correspond to a validator created
     * via `newValidator` above.
     *
     * This method will return the exit proofs needed to process eigenpod withdrawals.
     * Additionally, it will send the withdrawal amount to the validator's withdrawal
     * destination.
     */
    // function exitValidator(uint40 validatorIndex) public returns (BeaconWithdrawal memory) {
    //     emit log_named_uint("- BeaconChain.exitValidator: ", validatorIndex);

    //     Validator memory validator = validators[validatorIndex];

    //     // Get the withdrawal amount and destination
    //     uint amountToWithdraw = validator.effectiveBalanceGwei * GWEI_TO_WEI;
    //     address destination = _toAddress(validator.withdrawalCreds);

    //     // Generate exit proofs for a full exit
    //     BeaconWithdrawal memory withdrawal = _genExitProof(validator);

    //     // Update state - set validator balance to zero and send balance to withdrawal destination
    //     validators[validatorIndex].effectiveBalanceGwei = 0;
    //     cheats.deal(destination, destination.balance + amountToWithdraw);

    //     return withdrawal;
    // }

    /**
     * Note: `delta` is expected to be a raw token amount. This method will convert the delta to Gwei
     */
    // function updateBalance(uint40 validatorIndex, int delta) public returns (BalanceUpdate memory) {
    //     delta /= int(GWEI_TO_WEI);
        
    //     emit log_named_uint("- BeaconChain.updateBalance for validator: ", validatorIndex);
    //     emit log_named_int("- BeaconChain.updateBalance delta gwei: ", delta);
        
    //     // Apply delta and update validator balance in state
    //     uint64 newBalance;
    //     if (delta <= 0) {
    //         newBalance = validators[validatorIndex].effectiveBalanceGwei - uint64(uint(-delta));
    //     } else {
    //         newBalance = validators[validatorIndex].effectiveBalanceGwei + uint64(uint(delta));
    //     }
    //     validators[validatorIndex].effectiveBalanceGwei = newBalance;
        
    //     // Generate balance update proof
    //     Validator memory validator = validators[validatorIndex];
    //     BalanceUpdate memory update = _genBalanceUpdateProof(validator);

    //     return update;
    // }

    function setNextTimestamp(uint64 timestamp) public {
        nextTimestamp = timestamp;
    }

    function balanceOfGwei(uint40 validatorIndex) public view returns (uint64) {
        return validators[validatorIndex].effectiveBalanceGwei;
    }

    function pubkeyHash(uint40 validatorIndex) public view returns (bytes32) {
        return validators[validatorIndex].pubkeyHash;
    }

    /**
     * INTERNAL/HELPER METHODS:
     */

    /**
     * @dev For a new validator, generate the beacon chain block root and merkle proof
     * needed to prove withdrawal credentials to an EigenPod.
     *
     * The generated block root is sent to the `BeaconChainOracleMock`, and can be
     * queried using `proof.oracleTimestamp` to validate the generated proof.
     */
    function _genCredentialsProof(Validator memory validator) internal returns (CredentialsProofs memory) {
        CredentialsProofs memory proof;

        proof.validatorIndices = new uint40[](1);
        proof.validatorIndices[0] = validator.validatorIndex;

        // Create validatorFields for the new validator
        proof.validatorFields = new bytes32[][](1);
        proof.validatorFields[0] = new bytes32[](2 ** BeaconChainProofs.VALIDATOR_FIELD_TREE_HEIGHT);
        proof.validatorFields[0][BeaconChainProofs.VALIDATOR_PUBKEY_INDEX] = validator.pubkeyHash;
        proof.validatorFields[0][BeaconChainProofs.VALIDATOR_WITHDRAWAL_CREDENTIALS_INDEX] = 
            bytes32(validator.withdrawalCreds);
        proof.validatorFields[0][BeaconChainProofs.VALIDATOR_BALANCE_INDEX] = 
            _toLittleEndianUint64(validator.effectiveBalanceGwei);

        // Calculate beaconStateRoot using validator index and an empty proof:
        proof.validatorFieldsProofs = new bytes[](1);
        proof.validatorFieldsProofs[0] = new bytes(VAL_FIELDS_PROOF_LEN);
        bytes32 validatorRoot = Merkle.merkleizeSha256(proof.validatorFields[0]);
        uint index = _calcValProofIndex(validator.validatorIndex);

        bytes32 beaconStateRoot = Merkle.processInclusionProofSha256({
            proof: proof.validatorFieldsProofs[0],
            leaf: validatorRoot,
            index: index
        });

        // Calculate blockRoot using beaconStateRoot and an empty proof:
        bytes memory blockRootProof = new bytes(BLOCKROOT_PROOF_LEN);
        bytes32 blockRoot = Merkle.processInclusionProofSha256({
            proof: blockRootProof,
            leaf: beaconStateRoot,
            index: BeaconChainProofs.STATE_ROOT_INDEX
        });

        proof.stateRootProof = BeaconChainProofs.StateRootProof({
            beaconStateRoot: beaconStateRoot,
            proof: blockRootProof
        });

        // Send the block root to the oracle and increment timestamp:
        proof.oracleTimestamp = uint64(nextTimestamp);

        oracle.setBlockRoot(nextTimestamp, blockRoot);
        nextTimestamp++;

        return proof;
    }
    
    /**
     * @dev Generates the proofs and roots needed to prove a validator's exit from
     * the beacon chain.
     *
     * The generated beacon block root is sent to `BeaconChainOracleMock`, and can
     * be queried using `withdrawal.oracleTimestamp` to validate the generated proof.
     *
     * Since a withdrawal proof requires proving multiple leaves in the same tree, this
     * method uses `_genConvergentProofs` to calculate proofs and roots for intermediate
     * subtrees, while retaining the information needed to supply an eigenpod with a proof.
     *
     * The overall merkle tree being proven looks like this:
     *
     * - beaconBlockRoot (submitted to oracle at end)
     * -- beaconStateRoot
     * ---- validatorFieldsRoot
     * ---- blockRoot (from historical summaries)
     * -------- slotRoot
     * -------- executionPayloadRoot
     * ---------------- timestampRoot
     * ---------------- withdrawalFieldsRoot
     *
     * This method first generates proofs for the lowest leaves, and uses the resulting
     * intermediate hashes to generate proofs for higher leaves. Eventually, all of these
     * roots are calculated and the final beaconBlockRoot can be calculated and sent to the
     * oracle.
     */
    // function _genExitProof(Validator memory validator) internal returns (BeaconWithdrawal memory) {
    //     BeaconWithdrawal memory withdrawal;
    //     uint64 withdrawalEpoch = uint64(block.timestamp);

    //     // Get a new, unique timestamp for queries to the oracle
    //     withdrawal.oracleTimestamp = uint64(nextTimestamp);
    //     nextTimestamp++;

    //     // Initialize proof arrays
    //     BeaconChainProofs.WithdrawalProof memory withdrawalProof = _initWithdrawalProof({
    //         withdrawalEpoch: withdrawalEpoch,
    //         withdrawalIndex: WITHDRAWAL_INDEX,
    //         oracleTimestamp: withdrawal.oracleTimestamp
    //     });

    //     // Calculate withdrawalFields and record the validator's index and withdrawal amount
    //     withdrawal.withdrawalFields = new bytes32[][](1);
    //     withdrawal.withdrawalFields[0] = new bytes32[](2 ** BeaconChainProofs.WITHDRAWAL_FIELD_TREE_HEIGHT);
    //     withdrawal.withdrawalFields[0][BeaconChainProofs.WITHDRAWAL_VALIDATOR_INDEX_INDEX] =
    //         _toLittleEndianUint64(validator.validatorIndex);
    //     withdrawal.withdrawalFields[0][BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX] =
    //         _toLittleEndianUint64(validator.effectiveBalanceGwei);

    //     {
    //         /**
    //          * Generate proofs then root for subtree:
    //          * 
    //          * executionPayloadRoot
    //          * - timestampRoot         (withdrawalProof.timestampProof)
    //          * - withdrawalFieldsRoot  (withdrawalProof.withdrawalProof)
    //          */            
    //         withdrawalProof.executionPayloadRoot = _genExecPayloadProofs({ 
    //             withdrawalProof: withdrawalProof,
    //             withdrawalRoot: Merkle.merkleizeSha256(withdrawal.withdrawalFields[0])
    //         });
    //     }

    //     {
    //         /**
    //          * Generate proofs then root for subtree:
    //          * 
    //          * blockRoot (historical summaries)
    //          * - slotRoot             (withdrawalProof.slotProof)
    //          * - executionPayloadRoot (withdrawalProof.executionPayloadProof)
    //          */
    //         withdrawalProof.blockRoot = _genBlockRootProofs({ 
    //             withdrawalProof: withdrawalProof
    //         });
    //     }

    //     // validatorFields
    //     withdrawal.validatorFields = new bytes32[][](1);
    //     withdrawal.validatorFields[0] = new bytes32[](2 ** BeaconChainProofs.VALIDATOR_FIELD_TREE_HEIGHT);
    //     withdrawal.validatorFields[0][BeaconChainProofs.VALIDATOR_PUBKEY_INDEX] = validator.pubkeyHash;
    //     withdrawal.validatorFields[0][BeaconChainProofs.VALIDATOR_WITHDRAWABLE_EPOCH_INDEX] = 
    //         _toLittleEndianUint64(withdrawalEpoch);
        
    //     withdrawal.validatorFieldsProofs = new bytes[](1);
    //     withdrawal.validatorFieldsProofs[0] = new bytes(VAL_FIELDS_PROOF_LEN);

    //     {
    //         /**
    //          * Generate proofs then root for subtree:
    //          *
    //          * beaconStateRoot
    //          * - validatorFieldsRoot               (withdrawal.validatorFieldsProofs[0])
    //          * - blockRoot (historical summaries)  (withdrawalProof.historicalSummaryBlockRootProof)
    //          */
    //         withdrawal.stateRootProof.beaconStateRoot = _genBeaconStateRootProofs({ 
    //             withdrawalProof: withdrawalProof,
    //             validatorFieldsProof: withdrawal.validatorFieldsProofs[0],
    //             validatorIndex: validator.validatorIndex,
    //             validatorRoot: Merkle.merkleizeSha256(withdrawal.validatorFields[0])
    //         });
    //     }

    //     withdrawal.withdrawalProofs = new BeaconChainProofs.WithdrawalProof[](1);
    //     withdrawal.withdrawalProofs[0] = withdrawalProof;

    //     // Calculate beaconBlockRoot using beaconStateRoot and an empty proof:
    //     withdrawal.stateRootProof.proof = new bytes(BLOCKROOT_PROOF_LEN);
    //     bytes32 beaconBlockRoot = Merkle.processInclusionProofSha256({
    //         proof: withdrawal.stateRootProof.proof,
    //         leaf: withdrawal.stateRootProof.beaconStateRoot,
    //         index: BeaconChainProofs.STATE_ROOT_INDEX
    //     });

    //     // Send the block root to the oracle
    //     oracle.setBlockRoot(withdrawal.oracleTimestamp, beaconBlockRoot);
    //     return withdrawal;
    // }

    // function _genBalanceUpdateProof(Validator memory validator) internal returns (BalanceUpdate memory) {
    //     BalanceUpdate memory update;

    //     update.validatorIndices = new uint40[](1);
    //     update.validatorIndices[0] = validator.validatorIndex;

    //     // Create validatorFields showing the balance update
    //     update.validatorFields = new bytes32[][](1);
    //     update.validatorFields[0] = new bytes32[](2 ** BeaconChainProofs.VALIDATOR_FIELD_TREE_HEIGHT);
    //     update.validatorFields[0][BeaconChainProofs.VALIDATOR_PUBKEY_INDEX] = validator.pubkeyHash;
    //     update.validatorFields[0][BeaconChainProofs.VALIDATOR_WITHDRAWAL_CREDENTIALS_INDEX] = 
    //         bytes32(validator.withdrawalCreds);
    //     update.validatorFields[0][BeaconChainProofs.VALIDATOR_BALANCE_INDEX] = 
    //         _toLittleEndianUint64(validator.effectiveBalanceGwei);

    //     // Calculate beaconStateRoot using validator index and an empty proof:
    //     update.validatorFieldsProofs = new bytes[](1);
    //     update.validatorFieldsProofs[0] = new bytes(VAL_FIELDS_PROOF_LEN);
    //     bytes32 validatorRoot = Merkle.merkleizeSha256(update.validatorFields[0]);
    //     uint index = _calcValProofIndex(validator.validatorIndex);

    //     bytes32 beaconStateRoot = Merkle.processInclusionProofSha256({
    //         proof: update.validatorFieldsProofs[0],
    //         leaf: validatorRoot,
    //         index: index
    //     });

    //     // Calculate blockRoot using beaconStateRoot and an empty proof:
    //     bytes memory blockRootProof = new bytes(BLOCKROOT_PROOF_LEN);
    //     bytes32 blockRoot = Merkle.processInclusionProofSha256({
    //         proof: blockRootProof,
    //         leaf: beaconStateRoot,
    //         index: BeaconChainProofs.STATE_ROOT_INDEX
    //     });

    //     update.stateRootProof = BeaconChainProofs.StateRootProof({
    //         beaconStateRoot: beaconStateRoot,
    //         proof: blockRootProof
    //     });

    //     // Send the block root to the oracle and increment timestamp:
    //     update.oracleTimestamp = uint64(nextTimestamp);
    //     oracle.setBlockRoot(nextTimestamp, blockRoot);
    //     nextTimestamp++;
        
    //     return update;
    // }

    /**
     * @dev Generates converging merkle proofs for timestampRoot and withdrawalRoot
     * under the executionPayloadRoot.
     *
     * `withdrawalProof.timestampProof` and `withdrawalProof.withdrawalProof` are
     * directly updated here.
     *
     * @return executionPayloadRoot
     */
    // function _genExecPayloadProofs(
    //     BeaconChainProofs.WithdrawalProof memory withdrawalProof,
    //     bytes32 withdrawalRoot
    // ) internal view returns (bytes32) {

    //     uint withdrawalProofIndex = 
    //         (BeaconChainProofs.WITHDRAWALS_INDEX << (BeaconChainProofs.WITHDRAWALS_TREE_HEIGHT + 1)) |
    //         uint(withdrawalProof.withdrawalIndex);

    //     /**
    //      * Generate merkle proofs for timestampRoot and withdrawalRoot
    //      * that converge at or before executionPayloadRoot.
    //      * 
    //      * timestampProof length: 4
    //      * withdrawalProof length: 9
    //      */
    //     _genConvergentProofs({
    //         shortProof: withdrawalProof.timestampProof,
    //         shortIndex: BeaconChainProofs.TIMESTAMP_INDEX,
    //         shortLeaf: withdrawalProof.timestampRoot,
    //         longProof: withdrawalProof.withdrawalProof,
    //         longIndex: withdrawalProofIndex,
    //         longLeaf: withdrawalRoot
    //     });

    //     // Use generated proofs to calculate tree root and verify both proofs
    //     // result in the same root:
    //     bytes32 execPayloadRoot = Merkle.processInclusionProofSha256({
    //         proof: withdrawalProof.timestampProof,
    //         leaf: withdrawalProof.timestampRoot,
    //         index: BeaconChainProofs.TIMESTAMP_INDEX
    //     });

    //     bytes32 expectedRoot = Merkle.processInclusionProofSha256({
    //         proof: withdrawalProof.withdrawalProof,
    //         leaf: withdrawalRoot,
    //         index: withdrawalProofIndex
    //     });

    //     require(execPayloadRoot == expectedRoot, "_genExecPayloadProofs: mismatched roots");
        
    //     return execPayloadRoot;
    // }

    /**
     * @dev Generates converging merkle proofs for slotRoot and executionPayloadRoot
     * under the block root (historical summaries).
     *
     * `withdrawalProof.slotProof` and `withdrawalProof.executionPayloadProof` are
     * directly updated here.
     *
     * @return historical summary block root
     */
    // function _genBlockRootProofs(
    //     BeaconChainProofs.WithdrawalProof memory withdrawalProof
    // ) internal view returns (bytes32) {

    //     uint slotRootIndex = BeaconChainProofs.SLOT_INDEX;
    //     uint execPayloadIndex = 
    //         (BeaconChainProofs.BODY_ROOT_INDEX << BeaconChainProofs.BEACON_BLOCK_BODY_FIELD_TREE_HEIGHT) |
    //         BeaconChainProofs.EXECUTION_PAYLOAD_INDEX;

    //     /**
    //      * Generate merkle proofs for slotRoot and executionPayloadRoot
    //      * that converge at or before block root.
    //      * 
    //      * slotProof length: 3
    //      * executionPayloadProof length: 7
    //      */
    //     _genConvergentProofs({
    //         shortProof: withdrawalProof.slotProof,
    //         shortIndex: slotRootIndex,
    //         shortLeaf: withdrawalProof.slotRoot,
    //         longProof: withdrawalProof.executionPayloadProof,
    //         longIndex: execPayloadIndex,
    //         longLeaf: withdrawalProof.executionPayloadRoot
    //     });

    //     // Use generated proofs to calculate tree root and verify both proofs
    //     // result in the same root:
    //     bytes32 blockRoot = Merkle.processInclusionProofSha256({
    //         proof: withdrawalProof.slotProof,
    //         leaf: withdrawalProof.slotRoot,
    //         index: slotRootIndex
    //     });

    //     bytes32 expectedRoot = Merkle.processInclusionProofSha256({
    //         proof: withdrawalProof.executionPayloadProof,
    //         leaf: withdrawalProof.executionPayloadRoot,
    //         index: execPayloadIndex
    //     });

    //     require(blockRoot == expectedRoot, "_genBlockRootProofs: mismatched roots");
        
    //     return blockRoot;
    // }

    /**
     * @dev Generates converging merkle proofs for block root and validatorRoot
     * under the beaconStateRoot.
     *
     * `withdrawalProof.historicalSummaryBlockRootProof` and `validatorFieldsProof` are
     * directly updated here.
     *
     * @return beaconStateRoot
     */
    // function _genBeaconStateRootProofs(
    //     BeaconChainProofs.WithdrawalProof memory withdrawalProof,
    //     bytes memory validatorFieldsProof,
    //     uint40 validatorIndex, 
    //     bytes32 validatorRoot
    // ) internal view returns (bytes32) {
    //     uint blockHeaderIndex = _calcBlockHeaderIndex(withdrawalProof);
    //     uint validatorProofIndex = _calcValProofIndex(validatorIndex);

    //     /**
    //      * Generate merkle proofs for validatorRoot and blockRoot
    //      * that converge at or before beaconStateRoot.
    //      * 
    //      * historicalSummaryBlockRootProof length: 44
    //      * validatorFieldsProof length: 46
    //      */
    //     _genConvergentProofs({
    //         shortProof: withdrawalProof.historicalSummaryBlockRootProof,
    //         shortIndex: blockHeaderIndex,
    //         shortLeaf: withdrawalProof.blockRoot,
    //         longProof: validatorFieldsProof,
    //         longIndex: validatorProofIndex,
    //         longLeaf: validatorRoot
    //     });

    //     // Use generated proofs to calculate tree root and verify both proofs
    //     // result in the same root:
    //     bytes32 beaconStateRoot = Merkle.processInclusionProofSha256({
    //         proof: withdrawalProof.historicalSummaryBlockRootProof,
    //         leaf: withdrawalProof.blockRoot,
    //         index: blockHeaderIndex
    //     });

    //     bytes32 expectedRoot = Merkle.processInclusionProofSha256({
    //         proof: validatorFieldsProof,
    //         leaf: validatorRoot,
    //         index: validatorProofIndex
    //     });

    //     require(beaconStateRoot == expectedRoot, "_genBeaconStateRootProofs: mismatched roots");
        
    //     return beaconStateRoot;
    // }

    /**
     * @dev Generates converging merkle proofs given two leaves and empty proofs. 
     * Basics:
     * - `shortProof` and `longProof` start as empty proofs initialized to the correct
     *   length for their respective paths.
     * - At the end of the method, `shortProof` and `longProof` are still entirely empty
     *   EXCEPT at the point where the proofs would normally converge under the root hash.
     * - At this point, `shortProof` will be assigned the current hash for the `longLeaf` proof
     *   ... and `longProof` will be assigned the current hash for the `shortLeaf` proof
     * 
     * Steps:
     * 1. Because the beacon chain has trees and leaves at varying heights, this method
     *    first calculates the root of the longer proof's subtree so that the remaining
     *    proof length is the same for both leaves.
     * 2. This method simultaneously computes each leaf's remaining proof step-by-step,
     *    performing effectively the same steps as `Merkle.processInclusionProof256`.
     * 3. At each step, we check to see if the current indices represent sibling leaves.
     * 4. If `shortIndex` and `longIndex` are siblings:
     *    - longProof[longProof_i] = curShortHash
     *    - shortProof[shortProof_i] = curLongHash
     * 
     * ... Once we've found this convergence and placed each sibling's current hash in
     * its opposing sibling's proof, we're done!
     * @param shortProof An empty proof initialized to the correct length for the shorter proof path
     * @param shortIndex The index of the 
     */
    function _genConvergentProofs(
        bytes memory shortProof,
        uint shortIndex,
        bytes32 shortLeaf,
        bytes memory longProof,
        uint longIndex,
        bytes32 longLeaf
    ) internal view {
        require(longProof.length >= shortProof.length, "_genConvergentProofs: invalid input");

        bytes32[1] memory curShortHash = [shortLeaf];
        bytes32[1] memory curLongHash = [longLeaf];

        // Calculate root of long subtree
        uint longProofOffset = longProof.length - shortProof.length;
        for (uint i = 32; i <= longProofOffset; i += 32) {
            if (longIndex % 2 == 0) {
                assembly {
                    mstore(0x00, mload(curLongHash))
                    mstore(0x20, mload(add(longProof, i)))
                }
            } else {
                assembly {
                    mstore(0x00, mload(add(longProof, i)))
                    mstore(0x20, mload(curLongHash))
                }
            }

            // Compute hash and divide index
            assembly {
                if iszero(staticcall(sub(gas(), 2000), 2, 0x00, 0x40, curLongHash, 0x20)) {
                    revert(0, 0)
                }
                longIndex := div(longIndex, 2)
            }
        }

        {
            // Now that we've calculated the longest sub-tree, continue merklizing both trees simultaneously.
            // When we reach two leaf indices s.t. A is even and B == A + 1, or vice versa, we know we have 
            // found the point where the two sub-trees converge.
            uint longProof_i = 32 + longProofOffset;
            uint shortProof_i = 32;
            bool foundConvergence;
            for (; longProof_i <= longProof.length; ) {
                if (_areSiblings(longIndex, shortIndex)) {
                    foundConvergence = true;
                    assembly {
                        mstore(add(longProof, longProof_i), mload(curShortHash))
                        mstore(add(shortProof, shortProof_i), mload(curLongHash))
                    }

                    break;
                }
                
                // Compute next hash for longProof
                {
                    if (longIndex % 2 == 0) {
                        assembly {
                            mstore(0x00, mload(curLongHash))
                            mstore(0x20, mload(add(longProof, longProof_i)))
                        }
                    } else {
                        assembly {
                            mstore(0x00, mload(add(longProof, longProof_i)))
                            mstore(0x20, mload(curLongHash))
                        }
                    }
                    
                    // Compute hash and divide index
                    assembly {
                        if iszero(staticcall(sub(gas(), 2000), 2, 0x00, 0x40, curLongHash, 0x20)) {
                            revert(0, 0)
                        }
                        longIndex := div(longIndex, 2)
                    }
                }
                
                // Compute next hash for shortProof
                {
                    if (shortIndex % 2 == 0) {
                        assembly {
                            mstore(0x00, mload(curShortHash))
                            mstore(0x20, mload(add(shortProof, shortProof_i)))
                        }
                    } else {
                        assembly {
                            mstore(0x00, mload(add(shortProof, shortProof_i)))
                            mstore(0x20, mload(curShortHash))
                        }
                    }

                    // Compute hash and divide index
                    assembly {
                        if iszero(staticcall(sub(gas(), 2000), 2, 0x00, 0x40, curShortHash, 0x20)) {
                            revert(0, 0)
                        }
                        shortIndex := div(shortIndex, 2)
                    }
                }

                longProof_i += 32;
                shortProof_i += 32;
            }

            require(foundConvergence, "proofs did not converge!");
        }
    }

    /**
     * PROOF LENGTHS, MISC CONSTANTS, AND OTHER HELPERS:
     */

    uint immutable BLOCKROOT_PROOF_LEN = 32 * BeaconChainProofs.BEACON_BLOCK_HEADER_FIELD_TREE_HEIGHT;
    uint immutable VAL_FIELDS_PROOF_LEN = 32 * (
        (BeaconChainProofs.VALIDATOR_TREE_HEIGHT + 1) + BeaconChainProofs.BEACON_STATE_FIELD_TREE_HEIGHT
    );

    uint immutable WITHDRAWAL_PROOF_LEN_CAPELLA = 32 * (
        BeaconChainProofs.EXECUTION_PAYLOAD_HEADER_FIELD_TREE_HEIGHT_CAPELLA + 
        BeaconChainProofs.WITHDRAWALS_TREE_HEIGHT + 1
    );

    uint immutable WITHDRAWAL_PROOF_LEN_DENEB= 32 * (
        BeaconChainProofs.EXECUTION_PAYLOAD_HEADER_FIELD_TREE_HEIGHT_DENEB + 
        BeaconChainProofs.WITHDRAWALS_TREE_HEIGHT + 1
    );

    uint immutable EXECPAYLOAD_PROOF_LEN = 32 * (
        BeaconChainProofs.BEACON_BLOCK_HEADER_FIELD_TREE_HEIGHT + 
        BeaconChainProofs.BEACON_BLOCK_BODY_FIELD_TREE_HEIGHT
    );
    uint immutable SLOT_PROOF_LEN = 32 * (
        BeaconChainProofs.BEACON_BLOCK_HEADER_FIELD_TREE_HEIGHT
    );
    uint immutable TIMESTAMP_PROOF_LEN_CAPELLA = 32 * (
        BeaconChainProofs.EXECUTION_PAYLOAD_HEADER_FIELD_TREE_HEIGHT_CAPELLA
    );
    uint immutable TIMESTAMP_PROOF_LEN_DENEB = 32 * (
        BeaconChainProofs.EXECUTION_PAYLOAD_HEADER_FIELD_TREE_HEIGHT_DENEB
    );
    uint immutable HISTSUMMARY_PROOF_LEN = 32 * (
        BeaconChainProofs.BEACON_STATE_FIELD_TREE_HEIGHT +
        BeaconChainProofs.HISTORICAL_SUMMARIES_TREE_HEIGHT +
        BeaconChainProofs.BLOCK_ROOTS_TREE_HEIGHT + 2
    );

    uint immutable HIST_SUMMARIES_PROOF_INDEX = BeaconChainProofs.HISTORICAL_SUMMARIES_INDEX << (
        BeaconChainProofs.HISTORICAL_SUMMARIES_TREE_HEIGHT + 1 +
        BeaconChainProofs.BLOCK_ROOTS_TREE_HEIGHT + 1
    );

    // function _initWithdrawalProof(
    //     uint64 withdrawalEpoch, 
    //     uint64 withdrawalIndex,
    //     uint64 oracleTimestamp
    // ) internal view returns (BeaconChainProofs.WithdrawalProof memory) {
    //     uint256 withdrawalProofLength;
    //     uint256 timestampProofLength;
    //     if (block.timestamp > eigenPodManager.denebForkTimestamp()) {
    //         withdrawalProofLength = WITHDRAWAL_PROOF_LEN_DENEB;
    //         timestampProofLength = TIMESTAMP_PROOF_LEN_DENEB;
    //     } else {
    //         withdrawalProofLength = WITHDRAWAL_PROOF_LEN_CAPELLA;
    //         timestampProofLength = TIMESTAMP_PROOF_LEN_CAPELLA;
    //     }
    //     return BeaconChainProofs.WithdrawalProof({
    //         withdrawalProof: new bytes(withdrawalProofLength),
    //         slotProof: new bytes(SLOT_PROOF_LEN),
    //         executionPayloadProof: new bytes(EXECPAYLOAD_PROOF_LEN),
    //         timestampProof: new bytes(timestampProofLength),
    //         historicalSummaryBlockRootProof: new bytes(HISTSUMMARY_PROOF_LEN),
    //         blockRootIndex: 0,
    //         historicalSummaryIndex: 0,
    //         withdrawalIndex: withdrawalIndex,
    //         blockRoot: bytes32(0),
    //         slotRoot: _toLittleEndianUint64(withdrawalEpoch * BeaconChainProofs.SLOTS_PER_EPOCH),
    //         timestampRoot: _toLittleEndianUint64(oracleTimestamp),
    //         executionPayloadRoot: bytes32(0)
    //     });
    // }

    // function _calcBlockHeaderIndex(BeaconChainProofs.WithdrawalProof memory withdrawalProof) internal view returns (uint) {
    //     return 
    //         HIST_SUMMARIES_PROOF_INDEX |
    //         (uint(withdrawalProof.historicalSummaryIndex) << (BeaconChainProofs.BLOCK_ROOTS_TREE_HEIGHT + 1)) |
    //         (BeaconChainProofs.BLOCK_SUMMARY_ROOT_INDEX << BeaconChainProofs.BLOCK_ROOTS_TREE_HEIGHT) |
    //         uint(withdrawalProof.blockRootIndex);
    // }

    function _calcValProofIndex(uint40 validatorIndex) internal pure returns (uint) {
        return 
            (BeaconChainProofs.VALIDATOR_TREE_ROOT_INDEX << (BeaconChainProofs.VALIDATOR_TREE_HEIGHT + 1)) | 
            uint(validatorIndex);
    }

    /// @dev Returns true if a and b are sibling indices in the same sub-tree.
    /// 
    /// i.e. the indices belong two child nodes that share a parent:
    /// [A, B] or [B, A]
    function _areSiblings(uint a, uint b) internal pure returns (bool) {
        return 
            (a % 2 == 0 && b == a + 1) || (b % 2 == 0 && a == b + 1);
    }

    function _getZeroNode(uint depth) internal view returns (bytes32) {
        require(depth < ZERO_NODES_LENGTH, "_getZeroNode: invalid depth");

        return zeroNodes[depth];
    }

    /// @dev Opposite of Endian.fromLittleEndianUint64
    function _toLittleEndianUint64(uint64 num) internal pure returns (bytes32) {
        uint256 lenum;
    
        // Rearrange the bytes from big-endian to little-endian format
        lenum |= uint256((num & 0xFF) << 56);
        lenum |= uint256((num & 0xFF00) << 40);
        lenum |= uint256((num & 0xFF0000) << 24);
        lenum |= uint256((num & 0xFF000000) << 8);
        lenum |= uint256((num & 0xFF00000000) >> 8);
        lenum |= uint256((num & 0xFF0000000000) >> 24);
        lenum |= uint256((num & 0xFF000000000000) >> 40);
        lenum |= uint256((num & 0xFF00000000000000) >> 56);

        // Shift the little-endian bytes to the end of the bytes32 value
        return bytes32(lenum << 192);
    }

    /// @dev Opposite of BeaconChainProofs.getBalanceAtIndex
    /// @return The new, updated balance root
    function _setBalanceAtIndex(bytes32 balanceRoot, uint40 validatorIndex, uint64 newBalanceGwei) internal returns (bytes32) {
        // Clear out old balance
        uint bitShiftAmount = 256 - (64 * ((validatorIndex % 4) + 1));
        uint mask = ~(uint(0xFFFFFFFFFFFFFFFF) << bitShiftAmount);
        uint clearedRoot = uint(balanceRoot) & mask;

        // Convert validator balance to little endian and shift to correct position
        uint leBalance = uint(_toLittleEndianUint64(newBalanceGwei));
        uint shiftedBalance = leBalance >> (192 - bitShiftAmount);

        return bytes32(clearedRoot | shiftedBalance);
    }

    /// @dev Helper to convert 32-byte withdrawal credentials to an address
    function _toAddress(bytes memory withdrawalCreds) internal pure returns (address a) {
        bytes32 creds = bytes32(withdrawalCreds);
        uint160 mask = type(uint160).max;

        assembly { a := and(creds, mask) }
    }
}
