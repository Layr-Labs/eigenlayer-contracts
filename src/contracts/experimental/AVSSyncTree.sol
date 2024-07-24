// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./AVSDirectory.sol";
import "../libraries/Merkle.sol";


contract AVSSyncTree is IAVSDirectory, Merkle {

    uint32 public constant MAX_OPERATOR_SET_SIZE = 512;
    uint32 public constant MAX_NUM_STRATEGIES = 32;


    function getOperatorSetRoot(address avs, uint32 operatorSetId, address[] operators, address[] strategies) external view returns (bytes32) {
        require(isOperatorSetAVS[avs], "AVSSyncTree.getOperatorSetRoot: is not operator set AVS");

        // verify that provided operators are members of provided operator set and are registered to the AVS
        _verifyOperatorStatus(operatorSetId, operators);

        bytes32[] memory operatorLeaves = new bytes32[](operators.length);
        for (uint256 i = 0; i < operators.length; i++) {
            // shares associated with this operator and these strategies
            uint256[] memory shares = _retrieveStrategyShares(operators[i], strategies);
            bytes32 operatorRoot = _computeOperatorRoot(strategies, shares);
            operatorLeaves[i] = keccak256(abi.encodePacked(operators[i], operatorRoot));     
        }
        return merkleizeSha256(operatorLeaves);
    }

    function _computeOperatorRoot(address[] strategies, uint256[] shares) internal pure returns (bytes32) {
        bytes32[] memory leaves = new bytes32[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            leaves[i] = keccak256(abi.encodePacked(strategies[i], shares[i]));
        }
        return merkleizeSha256(leaves);
    }

    function _computeOperatorSetRoot()

    function _verifyOperatorStatus(uint32 operatorSetId, address[] operators) internal view {
        require(operators.length <= MAX_OPERATOR_SET_SIZE, "AVSSyncTree._verifyOperatorStatus: operator set too large");
        for (uint256 i = 0; i < operators.length; i++) {
            require(delegation.isOperator(operators[i]), "AVSSyncTree.getOperatorSetRoot: operator not registered");
            require(avsOperatorStatus[avs][operators[i]] == OperatorAVSRegistrationStatus.Registered, "AVSSyncTree.getOperatorSetRoot: operator not registered to AVS");
            require(isMember[avs][operators[i]][operatorSetId], "AVSSyncTree.getOperatorSetRoot: operator not in operator set");
        }
    }
    function _retrieveStrategyShares(address operator, address[] strategies) internal view returns (uint256[] memory) {
        require(strategies.length <= MAX_NUM_STRATEGIES, "AVSSyncTree._retrieveStrategyShares: too many strategies");
        return delegation.getOperatorShares(operator, strategies);
    }

}