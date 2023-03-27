// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/libraries/BN254.sol";
import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";


contract ProofParsing is Test{
    string internal proofConfigJson;
    string prefix;

    bytes32[18] blockHeaderProof;
    bytes32[3] slotProof;
    bytes32[9] withdrawalProof;
    bytes32[46] validatorProof;



    bytes32[7] executionPayloadProof;
    bytes32[4] blockNumberProofs;


    bytes32 slotRoot;
    bytes32 executionPayloadRoot;
    bytes32 blockNumberRoots;

    function setJSON(string memory path) public {
        proofConfigJson = vm.readFile(path);
    }

    function getSlot() public returns(uint256) {
        return stdJson.readUint(proofConfigJson, ".slot");
    }

    function getValidatorIndex() public returns(uint256){
        return stdJson.readUint(proofConfigJson, ".validatorIndex");
    }

    function getWithdrawalIndex() public returns(uint256){
        return stdJson.readUint(proofConfigJson, ".withdrawalIndex");
    }

    function getBlockHeaderRootIndex() public returns(uint256){
        return stdJson.readUint(proofConfigJson, ".blockHeaderRootIndex");
    }

    function getBeaconStateRoot() public returns(bytes32){
        return stdJson.readBytes32(proofConfigJson, ".beaconStateRoot");
    }

    function getBlockHeaderRoot() public returns(bytes32){
        return stdJson.readBytes32(proofConfigJson, ".blockHeaderRoot");
    }

    function getBlockBodyRoot() public returns(bytes32){
        return stdJson.readBytes32(proofConfigJson, ".blockBodyRoot");
    }

    function getSlotRoot() public returns(bytes32){
        return stdJson.readBytes32(proofConfigJson, ".slotRoot");
    }

    function getBalanceRoot() public returns(bytes32){
        return stdJson.readBytes32(proofConfigJson, ".balanceRoot");
    }

    function getBlockNumberRoot() public returns(bytes32){
        return stdJson.readBytes32(proofConfigJson, ".blockNumberRoot");
    }

    function getExecutionPayloadRoot() public returns(bytes32){
        return stdJson.readBytes32(proofConfigJson, ".executionPayloadRoot");
    }
    function getExecutionPayloadProof () public returns(bytes32[7] memory){
        for (uint i = 0; i < 7; i++) {
            prefix = string.concat(".ExecutionPayloadProof[", string.concat(vm.toString(i), "]"));
            executionPayloadProof[i] = (stdJson.readBytes32(proofConfigJson, prefix)); 
        }
        return executionPayloadProof;
    }

    function getBlockNumberProof () public returns(bytes32[4] memory){
        for (uint i = 0; i < 4; i++) {
            prefix = string.concat(".BlockNumberProof[", string.concat(vm.toString(i), "]"));
            blockNumberProofs[i] = (stdJson.readBytes32(proofConfigJson, prefix)); 
        }
        return blockNumberProofs;
    }

    function getBlockHeaderProof() public returns(bytes32[18] memory){
        for (uint i = 0; i < 18; i++) {
            prefix = string.concat(".BlockHeaderProof[", string.concat(vm.toString(i), "]"));
            blockHeaderProof[i] = (stdJson.readBytes32(proofConfigJson, prefix)); 
        }
        return blockHeaderProof;
    }

    function getSlotProof() public returns(bytes32[3] memory){
        for (uint i = 0; i < 3; i++) {
            prefix = string.concat(".SlotProof[", string.concat(vm.toString(i), "]"));
            slotProof[i] = (stdJson.readBytes32(proofConfigJson, prefix)); 
        }
        return slotProof;
    }

    function getWithdrawalProof() public returns(bytes32[9] memory){
        for (uint i = 0; i < 9; i++) {
            prefix = string.concat(".WithdrawalProof[", string.concat(vm.toString(i), "]"));
            withdrawalProof[i] = (stdJson.readBytes32(proofConfigJson, prefix)); 
        }
        return withdrawalProof;
    }

    function getValidatorProof() public returns(bytes32[46] memory){
        for (uint i = 0; i < 46; i++) {
            prefix = string.concat(".ValidatorProof[", string.concat(vm.toString(i), "]"));
            validatorProof[i] = (stdJson.readBytes32(proofConfigJson, prefix)); 
        }
        return validatorProof;
    }
    
    function getWithdrawalFields() public returns(bytes32[] memory){
        bytes32[] memory withdrawalFields = new bytes32[](4);
        for (uint i = 0; i < 4; i++) {
            prefix = string.concat(".WithdrawalFields[", string.concat(vm.toString(i), "]"));
            emit log_named_bytes32("prefix", stdJson.readBytes32(proofConfigJson, prefix));
            withdrawalFields[i] = (stdJson.readBytes32(proofConfigJson, prefix)); 
        }
         emit log_named_uint("length withdrawal firle", withdrawalFields.length);
         return withdrawalFields;

    }

    function getValidatorFields() public returns(bytes32[] memory){
        bytes32[] memory validatorFields = new bytes32[](8);
        for (uint i = 0; i < 8; i++) {
            prefix = string.concat(".ValidatorFields[", string.concat(vm.toString(i), "]"));
            validatorFields[i] = (stdJson.readBytes32(proofConfigJson, prefix)); 
        }
        return validatorFields;
    }

    function getValidatorBalanceProof() public returns(bytes32[] memory){
        bytes32[] memory validatorBalanceProof = new bytes32[](44);
        for (uint i = 0; i < 44; i++) {
            prefix = string.concat(".ValidatorBalanceProof[", string.concat(vm.toString(i), "]"));
            validatorBalanceProof[i] = (stdJson.readBytes32(proofConfigJson, prefix)); 
        }
        return validatorBalanceProof;
    }

    function getWithdrawalCredentialProof() public returns(bytes32[] memory){
        bytes32[] memory withdrawalCredenitalProof = new bytes32[](46);
        for (uint i = 0; i < 46; i++) {
            prefix = string.concat(".WithdrawalCredenitalProof[", string.concat(vm.toString(i), "]"));
            withdrawalCredenitalProof[i] = (stdJson.readBytes32(proofConfigJson, prefix)); 
        }
        return withdrawalCredenitalProof;
    }
}