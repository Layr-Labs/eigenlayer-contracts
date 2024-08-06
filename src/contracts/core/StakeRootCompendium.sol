// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";
import "../libraries/Merkle.sol";

import "@risc0-ethereum/contracts/src/IRiscZeroVerifier.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";


contract StakeRootCompendium is IStakeRootCompendium, OwnableUpgradeable {

    uint32 public constant MAX_OPERATOR_SET_SIZE = 512;
    uint32 public constant MAX_NUM_STRATEGIES = 32;

    IDelegationManager public delegation;
    IAVSDirectory public avsDirectory;
    constructor(
        IDelegationManager _delegation,
        IAVSDirectory _avsDirectory
    ) {
        _disableInitializers();
        delegation = _delegation;
        avsDirectory = _avsDirectory;
    }

     function getStakeRoot(address[] calldata avss, bytes32[] calldata operatorSetRoots) external view returns (bytes32) {
        require(avss.length == operatorSetRoots.length, "AVSSyncTree.getStakeRoot: AVS and operator set roots length mismatch");
    
        bytes32[] memory operatorSetLeaves = new bytes32[](avss.length);
        for (uint256 i = 0; i < avss.length; i++) {
            operatorSetLeaves[i] = keccak256(abi.encodePacked(avss[i], operatorSetRoots[i]));
        }
        return Merkle.merkleizeKeccak256(operatorSetLeaves);
    }

    function getOperatorSetRoot(address avs, uint32 operatorSetId, address[] operators, address[] strategies) external view returns (bytes32) {
        require(avsDirectory.isOperatorSetAVS[avs], "AVSSyncTree.getOperatorSetRoot: is not operator set AVS");
        require(operators.length <= MAX_OPERATOR_SET_SIZE, "AVSSyncTree._verifyOperatorStatus: operator set too large");

        // verify that provided operators are members of provided operator set and are registered to the AVS
        _verifyOperatorStatus(avs, operatorSetId, operators);

        bytes32[] memory operatorLeaves = new bytes32[](operators.length);
        for (uint256 i = 0; i < operators.length; i++) {
            require(delegation.isOperator(operators[i]), "AVSSyncTree.getOperatorSetRoot: operator not registered");
            require(avsOperatorStatus[avs][operators[i]] == OperatorAVSRegistrationStatus.REGISTERED, "AVSSyncTree.getOperatorSetRoot: operator not registered to AVS");
            require(isMember[avs][operators[i]][operatorSetId], "AVSSyncTree.getOperatorSetRoot: operator not in operator set");

            // shares associated with this operator and these strategies
            uint256[] memory shares = _retrieveStrategyShares(operators[i], strategies);
            bytes32 operatorRoot = _computeOperatorRoot(strategies, shares);
            operatorLeaves[i] = keccak256(abi.encodePacked(operators[i], operatorRoot));     
        }
        return Merkle.merkleizeKeccak256(operatorLeaves);
    }

    function _computeOperatorRoot(address[] strategies, uint256[] shares) internal pure returns (bytes32) {
        bytes32[] memory leaves = new bytes32[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            leaves[i] = keccak256(abi.encodePacked(strategies[i], shares[i]));
        }
        return Merkle.merkleizeKeccak256(leaves);
    }

    function _retrieveStrategyShares(address operator, address[] strategies) internal view returns (uint256[] memory) {
        require(strategies.length <= MAX_NUM_STRATEGIES, "AVSSyncTree._retrieveStrategyShares: too many strategies");
        return delegation.getOperatorShares(operator, strategies);
    }

}