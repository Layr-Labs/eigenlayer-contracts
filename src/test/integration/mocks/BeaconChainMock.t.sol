// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "forge-std/Test.sol";

import "src/contracts/libraries/BeaconChainProofs.sol";
import "src/contracts/libraries/Merkle.sol";

import "src/test/integration/TimeMachine.t.sol";
import "src/test/integration/mocks/BeaconChainOracleMock.t.sol";

struct CredentialsProofs {
    uint64 oracleTimestamp;
    BeaconChainProofs.StateRootProof stateRootProof;
    uint40[] validatorIndices;
    bytes[] validatorFieldsProofs;
    bytes32[][] validatorFields;
}

struct BeaconWithdrawal {
    uint64 oracleTimestamp;
    BeaconChainProofs.StateRootProof stateRootProof;
    BeaconChainProofs.WithdrawalProof[] withdrawalProofs;
    bytes[] validatorFieldsProofs;
    bytes32[][] validatorFields;
    bytes32[][] withdrawalFields;
}

contract BeaconChainMock is Test {

    Vm cheats = Vm(HEVM_ADDRESS);

    struct Validator {
        bytes32 pubkeyHash;
        uint40 validatorIndex;
        bytes withdrawalCreds;
        uint64 effectiveBalanceGwei;
    }
    
    uint40 nextValidatorIndex = 0;
    uint64 constant WITHDRAWAL_INDEX = 0;
    uint64 nextTimestamp;
    
    // Sequential list of created Validators
    // Validator[] validators;
    // mapping(uint40 => Validator) validators;

    mapping(uint40 => Validator) validators;

    BeaconChainOracleMock oracle;

    uint constant GWEI_TO_WEI = 1e9;
    uint immutable BLOCKROOT_PROOF_LEN = 32 * BeaconChainProofs.BEACON_BLOCK_HEADER_FIELD_TREE_HEIGHT;
    uint immutable VAL_FIELDS_PROOF_LEN = 
        32 * (
            (BeaconChainProofs.VALIDATOR_TREE_HEIGHT + 1) + BeaconChainProofs.BEACON_STATE_FIELD_TREE_HEIGHT
        );
    
    constructor(TimeMachine timeMachine, BeaconChainOracleMock beaconChainOracle) {
        nextTimestamp = timeMachine.proofGenStartTime();
        oracle = beaconChainOracle;
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
    ) public returns (uint40, CredentialsProofs memory) {

        // Create unique index for new validator
        uint40 validatorIndex = nextValidatorIndex;
        nextValidatorIndex++;

        // Create new validator and record in state
        Validator memory validator = Validator({
            pubkeyHash: keccak256(abi.encodePacked(validatorIndex)),
            validatorIndex: validatorIndex,
            withdrawalCreds: withdrawalCreds,
            effectiveBalanceGwei: uint64(balanceWei / GWEI_TO_WEI)
        });
        validators[validatorIndex] = validator;

        return (validator.validatorIndex, _genCredentialsProof(validator));
    }

    function exitValidator(uint40 validatorIndex, address pod) public returns (BeaconWithdrawal memory) {
        Validator memory validator = validators[validatorIndex];

        uint amountToWithdraw = validators[validatorIndex].effectiveBalanceGwei * GWEI_TO_WEI;

        BeaconWithdrawal memory withdrawal = _genExitProof(validator);

        // Update state - set validator balance to zero and send balance to pod
        validators[validatorIndex].effectiveBalanceGwei = 0;
        cheats.deal(pod, amountToWithdraw);

        return withdrawal;
    }

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

    uint immutable EXECPAYLOAD_INDEX = 
        (BeaconChainProofs.BODY_ROOT_INDEX << BeaconChainProofs.BEACON_BLOCK_BODY_FIELD_TREE_HEIGHT) |
        BeaconChainProofs.EXECUTION_PAYLOAD_INDEX;

    uint immutable WITHDRAWAL_PROOF_LEN = 32 * (
        BeaconChainProofs.EXECUTION_PAYLOAD_HEADER_FIELD_TREE_HEIGHT + 
        BeaconChainProofs.WITHDRAWALS_TREE_HEIGHT + 1
    );
    uint immutable EXECPAYLOAD_PROOF_LEN = 32 * (
        BeaconChainProofs.BEACON_BLOCK_HEADER_FIELD_TREE_HEIGHT + 
        BeaconChainProofs.BEACON_BLOCK_BODY_FIELD_TREE_HEIGHT
    );
    uint immutable SLOT_PROOF_LEN = 32 * (
        BeaconChainProofs.BEACON_BLOCK_HEADER_FIELD_TREE_HEIGHT
    );
    uint immutable TIMESTAMP_PROOF_LEN = 32 * (
        BeaconChainProofs.EXECUTION_PAYLOAD_HEADER_FIELD_TREE_HEIGHT
    );
    uint immutable HISTSUMMARY_PROOF_LEN = 32 * (
        BeaconChainProofs.BEACON_STATE_FIELD_TREE_HEIGHT +
        BeaconChainProofs.HISTORICAL_SUMMARIES_TREE_HEIGHT +
        BeaconChainProofs.BLOCK_ROOTS_TREE_HEIGHT + 2
    );

    function _initWithdrawalProof(
        uint64 withdrawalEpoch, 
        uint64 withdrawalIndex,
        uint64 oracleTimestamp
    ) internal view returns (BeaconChainProofs.WithdrawalProof memory) {
        return BeaconChainProofs.WithdrawalProof({
            withdrawalProof: new bytes(WITHDRAWAL_PROOF_LEN),
            slotProof: new bytes(SLOT_PROOF_LEN),
            executionPayloadProof: new bytes(EXECPAYLOAD_PROOF_LEN),
            timestampProof: new bytes(TIMESTAMP_PROOF_LEN),
            historicalSummaryBlockRootProof: new bytes(HISTSUMMARY_PROOF_LEN),
            blockRootIndex: 0,
            historicalSummaryIndex: 0,
            withdrawalIndex: withdrawalIndex,
            blockRoot: bytes32(0),
            slotRoot: _toLittleEndianUint64(withdrawalEpoch * BeaconChainProofs.SLOTS_PER_EPOCH),
            timestampRoot: _toLittleEndianUint64(oracleTimestamp),
            executionPayloadRoot: bytes32(0)
        });
    }
    
    function _genExitProof(Validator memory validator) internal returns (BeaconWithdrawal memory) {
        // revert("exitValidator: unimplemented");
        BeaconWithdrawal memory withdrawal;
        uint64 withdrawalEpoch = uint64(block.timestamp);

        withdrawal.oracleTimestamp = uint64(nextTimestamp);
        nextTimestamp++;

        BeaconChainProofs.WithdrawalProof memory withdrawalProof = _initWithdrawalProof({
            withdrawalEpoch: withdrawalEpoch,
            withdrawalIndex: WITHDRAWAL_INDEX,
            oracleTimestamp: withdrawal.oracleTimestamp
        });

        // withdrawalFields 
        withdrawal.withdrawalFields = new bytes32[][](1);
        withdrawal.withdrawalFields[0] = new bytes32[](2 ** BeaconChainProofs.WITHDRAWAL_FIELD_TREE_HEIGHT);
        withdrawal.withdrawalFields[0][BeaconChainProofs.WITHDRAWAL_VALIDATOR_INDEX_INDEX] =
            _toLittleEndianUint64(validator.validatorIndex);
        withdrawal.withdrawalFields[0][BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX] =
            _toLittleEndianUint64(validator.effectiveBalanceGwei);

        {
            /**
             * Generate root and proofs:
             * execPayloadRoot
             * - timestampRoot
             * - withdrawalFieldsRoot
             */
            (
                bytes32 execPayloadRoot,
                bytes memory timestampProof,
                bytes memory withdrawalFieldsProof
            ) = _genExecPayloadProofs({ 
                withdrawalIndex: withdrawalProof.withdrawalIndex,
                withdrawalRoot: Merkle.merkleizeSha256(withdrawal.withdrawalFields[0]),
                timestampRoot: withdrawalProof.timestampRoot
            });

            withdrawalProof.executionPayloadRoot = execPayloadRoot;

            withdrawalProof.timestampProof = timestampProof;
            withdrawalProof.withdrawalProof = withdrawalFieldsProof;
        }

        {
            /**
             * Generate root and proofs:
             * blockRoot
             * - slotRoot
             * - execPayloadRoot
             */
            (
                bytes32 blockRoot,
                bytes memory slotRootProof,
                bytes memory execPayloadProof
            ) = _genBlockRootProofs({ 
                slotRoot: withdrawalProof.slotRoot, 
                execPayloadRoot: withdrawalProof.executionPayloadRoot 
            });

            withdrawalProof.blockRoot = blockRoot;

            withdrawalProof.slotProof = slotRootProof;
            withdrawalProof.executionPayloadProof = execPayloadProof;
        }

        // validatorFields
        withdrawal.validatorFields = new bytes32[][](1);
        withdrawal.validatorFields[0] = new bytes32[](2 ** BeaconChainProofs.VALIDATOR_FIELD_TREE_HEIGHT);
        withdrawal.validatorFields[0][BeaconChainProofs.VALIDATOR_PUBKEY_INDEX] = validator.pubkeyHash;
        withdrawal.validatorFields[0][BeaconChainProofs.VALIDATOR_WITHDRAWABLE_EPOCH_INDEX] = 
            _toLittleEndianUint64(withdrawalEpoch);

        {
            /**
             * Generate root and proofs:
             * beaconStateRoot
             * - blockRoot
             * - validatorFieldsRoot
             */
            (
                bytes32 beaconStateRoot,
                bytes memory validatorFieldsProof,
                bytes memory blockRootProof
            ) = _genBeaconStateRootProofs({ 
                validatorIndex: validator.validatorIndex,
                validatorRoot: Merkle.merkleizeSha256(withdrawal.validatorFields[0]), 
                withdrawalProof: withdrawalProof
            });
    
            withdrawal.stateRootProof.beaconStateRoot = beaconStateRoot;

            withdrawal.validatorFieldsProofs = new bytes[](1);
            withdrawal.validatorFieldsProofs[0] = validatorFieldsProof;

            withdrawalProof.historicalSummaryBlockRootProof = blockRootProof;
        }

        withdrawal.withdrawalProofs = new BeaconChainProofs.WithdrawalProof[](1);
        withdrawal.withdrawalProofs[0] = withdrawalProof;

        // Calculate blockRoot using beaconStateRoot and an empty proof:
        withdrawal.stateRootProof.proof = new bytes(BLOCKROOT_PROOF_LEN);
        bytes32 beaconBlockRoot = Merkle.processInclusionProofSha256({
            proof: withdrawal.stateRootProof.proof,
            leaf: withdrawal.stateRootProof.beaconStateRoot,
            index: BeaconChainProofs.STATE_ROOT_INDEX
        });

        // Send the block root to the oracle
        oracle.setBlockRoot(withdrawal.oracleTimestamp, beaconBlockRoot);
        return withdrawal;
    }

    function _genExecPayloadProofs(
        uint64 withdrawalIndex,
        bytes32 withdrawalRoot,
        bytes32 timestampRoot
    ) internal view returns (bytes32, bytes memory, bytes memory) {

        uint withdrawalProofIndex = 
            (BeaconChainProofs.WITHDRAWALS_INDEX << (BeaconChainProofs.WITHDRAWALS_TREE_HEIGHT + 1)) |
            uint(withdrawalIndex);

        (bytes memory timestampProof, bytes memory withdrawalFieldsProof) = _calcConvergentProofs({
            shortProof: new bytes(TIMESTAMP_PROOF_LEN),
            shortIndex: BeaconChainProofs.TIMESTAMP_INDEX,
            shortLeaf: timestampRoot,
            longProof: new bytes(WITHDRAWAL_PROOF_LEN),
            longIndex: withdrawalProofIndex,
            longLeaf: withdrawalRoot
        });

        // Use generated proofs to calculate tree root and verify both proofs
        // result in the same root:
        bytes32 execPayloadRoot = Merkle.processInclusionProofSha256({
            proof: timestampProof,
            leaf: timestampRoot,
            index: BeaconChainProofs.TIMESTAMP_INDEX
        });

        bytes32 expectedRoot = Merkle.processInclusionProofSha256({
            proof: withdrawalFieldsProof,
            leaf: withdrawalRoot,
            index: withdrawalProofIndex
        });

        require(execPayloadRoot == expectedRoot, "_genBlockRootProofs: mismatched roots");
        
        return (execPayloadRoot, timestampProof, withdrawalFieldsProof);
    }

    uint immutable HIST_SUMMARIES_PROOF_INDEX = BeaconChainProofs.HISTORICAL_SUMMARIES_INDEX << (
        BeaconChainProofs.HISTORICAL_SUMMARIES_TREE_HEIGHT + 1 +
        BeaconChainProofs.BLOCK_ROOTS_TREE_HEIGHT + 1
    );

    function _genBlockRootProofs(
        bytes32 slotRoot, 
        bytes32 execPayloadRoot
    ) internal view returns (bytes32, bytes memory, bytes memory) {

        uint slotRootIndex = BeaconChainProofs.SLOT_INDEX;
        uint execPayloadIndex = 
            (BeaconChainProofs.BODY_ROOT_INDEX << BeaconChainProofs.BEACON_BLOCK_BODY_FIELD_TREE_HEIGHT) |
            BeaconChainProofs.EXECUTION_PAYLOAD_INDEX;

        (bytes memory slotProof, bytes memory execPayloadProof) = _calcConvergentProofs({
            shortProof: new bytes(SLOT_PROOF_LEN),
            shortIndex: slotRootIndex,
            shortLeaf: slotRoot,
            longProof: new bytes(EXECPAYLOAD_PROOF_LEN),
            longIndex: execPayloadIndex,
            longLeaf: execPayloadRoot
        });

        // Use generated proofs to calculate tree root and verify both proofs
        // result in the same root:
        bytes32 blockRoot = Merkle.processInclusionProofSha256({
            proof: slotProof,
            leaf: slotRoot,
            index: slotRootIndex
        });

        bytes32 expectedRoot = Merkle.processInclusionProofSha256({
            proof: execPayloadProof,
            leaf: execPayloadRoot,
            index: execPayloadIndex
        });

        require(blockRoot == expectedRoot, "_genBlockRootProofs: mismatched roots");
        
        return (blockRoot, slotProof, execPayloadProof);
    }

    function _genBeaconStateRootProofs(
        uint40 validatorIndex, 
        bytes32 validatorRoot,
        BeaconChainProofs.WithdrawalProof memory withdrawalProof
    ) internal view returns (bytes32, bytes memory, bytes memory) {
        uint blockHeaderIndex = _calcBlockHeaderIndex(withdrawalProof);
        uint validatorProofIndex = _calcValProofIndex(validatorIndex);

        (bytes memory blockRootProof, bytes memory validatorFieldsProof) = _calcConvergentProofs({
            shortProof: new bytes(HISTSUMMARY_PROOF_LEN),
            shortIndex: blockHeaderIndex,
            shortLeaf: withdrawalProof.blockRoot,
            longProof: new bytes(VAL_FIELDS_PROOF_LEN),
            longIndex: validatorProofIndex,
            longLeaf: validatorRoot
        });

        bytes32 beaconStateRoot = Merkle.processInclusionProofSha256({
            proof: blockRootProof,
            leaf: withdrawalProof.blockRoot,
            index: blockHeaderIndex
        });

        bytes32 expectedRoot = Merkle.processInclusionProofSha256({
            proof: validatorFieldsProof,
            leaf: validatorRoot,
            index: validatorProofIndex
        });

        require(beaconStateRoot == expectedRoot, "_genBeaconStateRootProofs: mismatched roots");
        
        return (beaconStateRoot, validatorFieldsProof, blockRootProof);
    }

    // Returns true if a and b are sibling indices in the same sub-tree.
    //
    // i.e this returns true if these indices represent two child nodes in
    // sequential positions:
    // [A, B] or [B, A]
    function _areSiblings(uint a, uint b) internal pure returns (bool) {
        return 
            (a % 2 == 0 && b == a + 1) || (b % 2 == 0 && a == b + 1);
    }

    function _calcConvergentProofs(
        bytes memory shortProof,
        uint shortIndex,
        bytes32 shortLeaf,
        bytes memory longProof,
        uint longIndex,
        bytes32 longLeaf
    ) internal view returns (bytes memory, bytes memory) {
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

        // 
        {
            // Now that we've calculated the validatorFields sub-tree, merklize both trees simultaneously.
            // When we reach two leaf indices s.t. A is even and B == A + 1, we know we have found the point
            // where the two sub-trees converge.
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

        return (shortProof, longProof);
    }

    function _calcBlockHeaderIndex(BeaconChainProofs.WithdrawalProof memory withdrawalProof) internal view returns (uint) {
        return 
            HIST_SUMMARIES_PROOF_INDEX |
            (uint(withdrawalProof.historicalSummaryIndex) << (BeaconChainProofs.BLOCK_ROOTS_TREE_HEIGHT + 1)) |
            (BeaconChainProofs.BLOCK_SUMMARY_ROOT_INDEX << BeaconChainProofs.BLOCK_ROOTS_TREE_HEIGHT) |
            uint(withdrawalProof.blockRootIndex);
    }

    function _calcValProofIndex(uint40 validatorIndex) internal pure returns (uint) {
        return 
            (BeaconChainProofs.VALIDATOR_TREE_ROOT_INDEX << (BeaconChainProofs.VALIDATOR_TREE_HEIGHT + 1)) | 
            uint(validatorIndex);
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
}