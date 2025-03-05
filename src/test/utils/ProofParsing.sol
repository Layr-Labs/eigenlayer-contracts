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

    function getSlot() public view returns (uint256) {
        return stdJson.readUint(proofConfigJson, ".slot");
    }

    function getValidatorIndex() public view returns (uint256) {
        return stdJson.readUint(proofConfigJson, ".validatorIndex");
    }





    function getBeaconStateRoot() public view returns (bytes32) {
        return stdJson.readBytes32(proofConfigJson, ".beaconStateRoot");
    }

    function getBlockRoot() public view returns (bytes32) {
        return stdJson.readBytes32(proofConfigJson, ".blockHeaderRoot");
    }















    

    function getValidatorFields() public returns(bytes32[] memory) {
        bytes32[] memory validatorFields = new bytes32[](8);
        for (uint i = 0; i < 8; i++) {
            prefix = string.concat(".ValidatorFields[", string.concat(vm.toString(i), "]"));
            validatorFields[i] = (stdJson.readBytes32(proofConfigJson, prefix)); 
        }
        return validatorFields;
    }


    function getWithdrawalCredentialProof() public returns(bytes memory) {
        bytes32[] memory withdrawalCredentialProof = new bytes32[](46);
        for (uint i = 0; i < 46; i++) {
            prefix = string.concat(".WithdrawalCredentialProof[", string.concat(vm.toString(i), "]"));
            withdrawalCredentialProof[i] = (stdJson.readBytes32(proofConfigJson, prefix)); 
        }
        return abi.encodePacked(withdrawalCredentialProof);
    }

}
