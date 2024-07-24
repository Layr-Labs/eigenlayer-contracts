// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";
import "../core/AVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";
import "../libraries/Merkle.sol";
import "../interfaces/IStrategy.sol";


contract AVSSyncTree {

    uint32 public constant MAX_OPERATOR_SET_SIZE = 512;
    uint32 public constant MAX_NUM_STRATEGIES = 32;

    IDelegationManager public delegation;
    AVSDirectory public avsDirectory;
    constructor(
        IDelegationManager _delegation,
        AVSDirectory _avsDirectory
    ) {
        delegation = _delegation;
        avsDirectory = _avsDirectory;
    }

    function getOperatorSetRoot(address avs, address[] calldata operators, IStrategy[] calldata strategies) external view returns (bytes32) {

        // verify that provided operators are members of provided operator set and are registered to the AVS
        _verifyOperatorStatus(avs, operators);

        bytes32[] memory operatorLeaves = new bytes32[](operators.length);
        for (uint256 i = 0; i < operators.length; i++) {
            // shares associated with this operator and these strategies
            uint256[] memory shares = _retrieveStrategyShares(operators[i], strategies);
            bytes32 operatorRoot = _computeOperatorRoot(strategies, shares);
            operatorLeaves[i] = keccak256(abi.encodePacked(operators[i], operatorRoot));     
        }
        return Merkle.merkleizeSha256(operatorLeaves);
    }

    function _computeOperatorRoot(IStrategy[] memory  strategies, uint256[] memory shares) internal pure returns (bytes32) {
        bytes32[] memory leaves = new bytes32[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            leaves[i] = keccak256(abi.encodePacked(strategies[i], shares[i]));
        }
        return Merkle.merkleizeSha256(leaves);
    }

    function _verifyOperatorStatus(address avs, address[] memory operators) internal view {
        require(operators.length <= MAX_OPERATOR_SET_SIZE, "AVSSyncTree._verifyOperatorStatus: operator set too large");
        for (uint256 i = 0; i < operators.length; i++) {
            require(delegation.isOperator(operators[i]), "AVSSyncTree.getOperatorSetRoot: operator not registered");
            require(avsDirectory.avsOperatorStatus(avs,operators[i]) == IAVSDirectory.OperatorAVSRegistrationStatus.REGISTERED, "AVSSyncTree.getOperatorSetRoot: operator not registered to AVS");
        }
    }
    function _retrieveStrategyShares(address operator, IStrategy[] memory strategies) internal view returns (uint256[] memory) {
        require(strategies.length <= MAX_NUM_STRATEGIES, "AVSSyncTree._retrieveStrategyShares: too many strategies");
        return delegation.getOperatorShares(operator, strategies);
    }

}