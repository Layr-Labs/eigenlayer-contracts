// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "forge-std/StdJson.sol";

contract ProofParsing is Test {
    string internal proofConfigJson;
    string prefix;

    bytes32[18] blockHeaderProof;
    bytes32[3] slotProof;
    bytes32[10] withdrawalProofDeneb;
    bytes32[9] withdrawalProofCapella;
    bytes32[46] validatorProof;
    bytes32[44] historicalSummaryProof;

    bytes32[7] executionPayloadProof;
    bytes32[5] timestampProofsCapella;
    bytes32[4] timestampProofsDeneb;

    bytes32 slotRoot;
    bytes32 executionPayloadRoot;

    function setJSON(string memory path) public {
        proofConfigJson = vm.readFile(path);
    }

    function getSlot() public view returns (uint) {
        return stdJson.readUint(proofConfigJson, ".slot");
    }

    function getValidatorIndex() public view returns (uint) {
        return stdJson.readUint(proofConfigJson, ".validatorIndex");
    }

    function getValidatorPubkeyHash() public view returns (bytes32) {
        return stdJson.readBytes32(proofConfigJson, ".ValidatorFields[0]");
    }

    function getWithdrawalIndex() public view returns (uint) {
        return stdJson.readUint(proofConfigJson, ".withdrawalIndex");
    }

    function getBlockRootIndex() public view returns (uint) {
        return stdJson.readUint(proofConfigJson, ".blockHeaderRootIndex");
    }

    function getHistoricalSummaryIndex() public view returns (uint) {
        return stdJson.readUint(proofConfigJson, ".historicalSummaryIndex");
    }

    function getBeaconStateRoot() public view returns (bytes32) {
        return stdJson.readBytes32(proofConfigJson, ".beaconStateRoot");
    }

    function getBlockRoot() public view returns (bytes32) {
        return stdJson.readBytes32(proofConfigJson, ".blockHeaderRoot");
    }

    function getSlotRoot() public view returns (bytes32) {
        return stdJson.readBytes32(proofConfigJson, ".slotRoot");
    }

    function getTimestampRoot() public view returns (bytes32) {
        return stdJson.readBytes32(proofConfigJson, ".timestampRoot");
    }

    function getExecutionPayloadRoot() public view returns (bytes32) {
        return stdJson.readBytes32(proofConfigJson, ".executionPayloadRoot");
    }

    function getLatestBlockRoot() public view returns (bytes32) {
        return stdJson.readBytes32(proofConfigJson, ".latestBlockHeaderRoot");
    }

    function getExecutionPayloadProof() public returns (bytes32[7] memory) {
        for (uint i = 0; i < 7; i++) {
            prefix = string.concat(".ExecutionPayloadProof[", string.concat(vm.toString(i), "]"));
            executionPayloadProof[i] = (stdJson.readBytes32(proofConfigJson, prefix));
        }
        return executionPayloadProof;
    }

    function getTimestampProofDeneb() public returns (bytes32[5] memory) {
        for (uint i = 0; i < 5; i++) {
            prefix = string.concat(".TimestampProof[", string.concat(vm.toString(i), "]"));
            timestampProofsCapella[i] = (stdJson.readBytes32(proofConfigJson, prefix));
        }
        return timestampProofsCapella;
    }

    function getTimestampProofCapella() public returns (bytes32[4] memory) {
        for (uint i = 0; i < 4; i++) {
            prefix = string.concat(".TimestampProof[", string.concat(vm.toString(i), "]"));
            timestampProofsDeneb[i] = (stdJson.readBytes32(proofConfigJson, prefix));
        }
        return timestampProofsDeneb;
    }

    function getBlockHeaderProof() public returns (bytes32[18] memory) {
        for (uint i = 0; i < 18; i++) {
            prefix = string.concat(".BlockHeaderProof[", string.concat(vm.toString(i), "]"));
            blockHeaderProof[i] = (stdJson.readBytes32(proofConfigJson, prefix));
        }
        return blockHeaderProof;
    }

    function getSlotProof() public returns (bytes32[3] memory) {
        for (uint i = 0; i < 3; i++) {
            prefix = string.concat(".SlotProof[", string.concat(vm.toString(i), "]"));
            slotProof[i] = (stdJson.readBytes32(proofConfigJson, prefix));
        }
        return slotProof;
    }

    function getStateRootProof() public returns (bytes memory) {
        bytes32[] memory stateRootProof = new bytes32[](3);
        for (uint i = 0; i < 3; i++) {
            prefix = string.concat(".StateRootAgainstLatestBlockHeaderProof[", string.concat(vm.toString(i), "]"));
            stateRootProof[i] = (stdJson.readBytes32(proofConfigJson, prefix));
        }
        return abi.encodePacked(stateRootProof);
    }

    function getWithdrawalProofDeneb() public returns (bytes32[10] memory) {
        for (uint i = 0; i < 10; i++) {
            prefix = string.concat(".WithdrawalProof[", string.concat(vm.toString(i), "]"));
            withdrawalProofDeneb[i] = (stdJson.readBytes32(proofConfigJson, prefix));
        }
        return withdrawalProofDeneb;
    }

    function getWithdrawalProofCapella() public returns (bytes32[9] memory) {
        for (uint i = 0; i < 9; i++) {
            prefix = string.concat(".WithdrawalProof[", string.concat(vm.toString(i), "]"));
            withdrawalProofCapella[i] = (stdJson.readBytes32(proofConfigJson, prefix));
        }
        return withdrawalProofCapella;
    }

    function getValidatorProof() public returns (bytes32[46] memory) {
        for (uint i = 0; i < 46; i++) {
            prefix = string.concat(".ValidatorProof[", string.concat(vm.toString(i), "]"));
            validatorProof[i] = (stdJson.readBytes32(proofConfigJson, prefix));
        }
        return validatorProof;
    }

    function getHistoricalSummaryProof() public returns (bytes32[44] memory) {
        for (uint i = 0; i < 44; i++) {
            prefix = string.concat(".HistoricalSummaryProof[", string.concat(vm.toString(i), "]"));
            historicalSummaryProof[i] = (stdJson.readBytes32(proofConfigJson, prefix));
        }
        return historicalSummaryProof;
    }

    function getWithdrawalFields() public returns (bytes32[] memory) {
        bytes32[] memory withdrawalFields = new bytes32[](4);
        for (uint i = 0; i < 4; i++) {
            prefix = string.concat(".WithdrawalFields[", string.concat(vm.toString(i), "]"));
            withdrawalFields[i] = (stdJson.readBytes32(proofConfigJson, prefix));
        }
        return withdrawalFields;
    }

    function getValidatorFields() public returns (bytes32[] memory) {
        bytes32[] memory validatorFields = new bytes32[](8);
        for (uint i = 0; i < 8; i++) {
            prefix = string.concat(".ValidatorFields[", string.concat(vm.toString(i), "]"));
            validatorFields[i] = (stdJson.readBytes32(proofConfigJson, prefix));
        }
        return validatorFields;
    }

    function getBalanceUpdateProof() public returns (bytes memory) {
        // Balance update proofs are the same as withdrawal credential proofs
        return getWithdrawalCredentialProof();
    }

    function getWithdrawalCredentialProof() public returns (bytes memory) {
        bytes32[] memory withdrawalCredentialProof = new bytes32[](46);
        for (uint i = 0; i < 46; i++) {
            prefix = string.concat(".WithdrawalCredentialProof[", string.concat(vm.toString(i), "]"));
            withdrawalCredentialProof[i] = (stdJson.readBytes32(proofConfigJson, prefix));
        }
        return abi.encodePacked(withdrawalCredentialProof);
    }

    function getValidatorFieldsProof() public returns (bytes32[] memory) {
        bytes32[] memory validatorFieldsProof = new bytes32[](46);
        for (uint i = 0; i < 46; i++) {
            prefix = string.concat(".ValidatorFieldsProof[", string.concat(vm.toString(i), "]"));
            validatorFieldsProof[i] = (stdJson.readBytes32(proofConfigJson, prefix));
        }
        return validatorFieldsProof;
    }
}
