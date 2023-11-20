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

struct WithdrawalProofs {
    uint64 oracleTimestamp;
    BeaconChainProofs.StateRootProof stateRootProof;
    BeaconChainProofs.WithdrawalProof[] withdrawalProofs;
    bytes[] validatorFieldsProofs;
    bytes32[][] validatorFields;
    bytes32[][] withdrawalFields;
}

contract BeaconChainMock {

    struct Validator {
        bytes32 pubkeyHash;
        uint40 validatorIndex;
        bytes withdrawalCreds;
        uint64 effectiveBalanceGwei;
    }
    
    uint40 nextIndex = 0;
    uint64 nextTimestamp;
    
    // Sequential list of created Validators
    // Validator[] validators;
    // mapping(uint40 => Validator) validators;

    mapping(uint40 => Validator) validators;

    BeaconChainOracleMock oracle;

    uint constant GWEI_TO_WEI = 1e9;
    uint immutable BLOCKROOT_PROOF_LEN = 32 * BeaconChainProofs.BEACON_BLOCK_HEADER_FIELD_TREE_HEIGHT;
    uint immutable FIELDS_PROOF_LEN = 
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
        uint40 validatorIndex = nextIndex;
        nextIndex++;

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

    function exitValidator(uint40 validatorIndex) public returns (WithdrawalProofs memory) {
        Validator memory validator = validators[validatorIndex];

        validators[validatorIndex].effectiveBalanceGwei = 0;

        return (_genExitProof(validator));
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
        proof.validatorFieldsProofs[0] = new bytes(FIELDS_PROOF_LEN);
        bytes32 validatorRoot = Merkle.merkleizeSha256(proof.validatorFields[0]);
        uint index = 
            (BeaconChainProofs.VALIDATOR_TREE_ROOT_INDEX << (BeaconChainProofs.VALIDATOR_TREE_HEIGHT + 1)) | 
            uint(validator.validatorIndex);

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

    function _genExitProof(Validator memory validator) internal returns (WithdrawalProofs memory) {
        revert("exitValidator: unimplemented");
        WithdrawalProofs memory proof;
        // uint64 withdrawalEpoch = uint64(block.timestamp);

        // // withdrawalFields and proof of withdrawalFields
        // proof.withdrawalFields = new bytes32[][](1);
        // proof.withdrawalFields[0] = new bytes32[](2 ** BeaconChainProofs.WITHDRAWAL_FIELD_TREE_HEIGHT);
        // proof.withdrawalFields[0][BeaconChainProofs.WITHDRAWAL_VALIDATOR_INDEX_INDEX] =
        //     uint40(_toLittleEndianUint64(validator.validatorIndex));
        // proof.withdrawalFields[0][BeaconChainProofs.WITHDRAWAL_VALIDATOR_AMOUNT_INDEX] =
        //     _toLittleEndianUint64(validator.effectiveBalanceGwei);

        // BeaconChainProofs.WithdrawalProof memory withdrawalProof;

        // // validatorFields and proof of validatorFields
        // proof.validatorFields = new bytes32[][](1);
        // proof.validatorFields[0] = new bytes32[](2 ** BeaconChainProofs.VALIDATOR_FIELD_TREE_HEIGHT);
        // proof.validatorFields[0][BeaconChainProofs.VALIDATOR_PUBKEY_INDEX] = validator.pubkeyHash;
        // proof.validatorFields[0][BeaconChainProofs.VALIDATOR_WITHDRAWABLE_EPOCH_INDEX] = 
        //     _toLittleEndianUint64(withdrawalEpoch);

        // // Calculate blockRoot using beaconStateRoot and an empty proof:
        // bytes memory blockRootProof = new bytes(BLOCKROOT_PROOF_LEN);
        // bytes32 blockRoot = Merkle.processInclusionProofSha256({
        //     proof: blockRootProof,
        //     leaf: beaconStateRoot,
        //     index: BeaconChainProofs.STATE_ROOT_INDEX
        // });

        // proof.stateRootProof = BeaconChainProofs.StateRootProof({
        //     beaconStateRoot: beaconStateRoot,
        //     proof: blockRootProof
        // });


        // // Send the block root to the oracle and increment timestamp:
        // proof.oracleTimestamp = uint64(nextTimestamp);
        // oracle.setBlockRoot(nextTimestamp, blockRoot);
        // nextTimestamp++;
        
        return proof;
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