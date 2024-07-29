// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IAVSSyncTree.sol";
import "../core/AVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";
import "../libraries/Merkle.sol";
import "../interfaces/IStrategy.sol";
import "@risc0-ethereum/contracts/src/IRiscZeroVerifier.sol";


contract AVSSyncTree is IAVSSyncTree {

    uint32 public immutable MAX_OPERATOR_SET_SIZE = 512;
    uint32 public immutable MAX_NUM_STRATEGIES = 32;

    address public verifier;

    IDelegationManager public delegation;
    AVSDirectory public avsDirectory;

    event VerifierChanged(address oldVerifier, address newVerifier);
    event SnarkProofVerified(bytes journal, bytes seal);
    constructor(
        IDelegationManager _delegation,
        AVSDirectory _avsDirectory
    ) {
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


    function getOperatorSetRoot(address avs, address[] calldata operators, address[] calldata strategies) public view returns (bytes32) {
        require(strategies.length <= MAX_NUM_STRATEGIES, "AVSSyncTree._retrieveStrategyShares: too many strategies");

        IStrategy[] memory strategiesConverted = new IStrategy[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            strategiesConverted[i] = IStrategy(strategies[i]);
        }

        bytes32[] memory operatorLeaves = new bytes32[](operators.length);
        for (uint256 i = 0; i < operators.length; i++) {
            // verify that provided operators are members of provided operator set and are registered to the AVS
            require(delegation.isOperator(operators[i]), "AVSSyncTree.getOperatorSetRoot: operator not registered");
            require(avsDirectory.avsOperatorStatus(avs,operators[i]) == IAVSDirectory.OperatorAVSRegistrationStatus.REGISTERED, "AVSSyncTree.getOperatorSetRoot: operator not registered to AVS");
            
            // shares associated with this operator and these strategies
            uint256[] memory shares = delegation.getOperatorShares(operators[i], strategiesConverted);
            bytes32 operatorRoot = _computeOperatorRoot(strategies, shares);
            
            operatorLeaves[i] = keccak256(abi.encodePacked(operators[i], operatorRoot));     
        }
        return Merkle.merkleizeKeccak256(operatorLeaves);
    }

    function _computeOperatorRoot(address[] memory  strategies, uint256[] memory shares) internal pure returns (bytes32) {
        bytes32[] memory leaves = new bytes32[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            leaves[i] = keccak256(abi.encodePacked(strategies[i], shares[i]));
        }
        return Merkle.merkleizeKeccak256(leaves);
    }

    function convertToStrategy(address[] memory addresses) public pure returns (IStrategy[] memory) {
        IStrategy[] memory strategies = new IStrategy[](addresses.length);
        for (uint256 i = 0; i < addresses.length; i++) {
            strategies[i] = IStrategy(addresses[i]);
        }
        return strategies;
    }

    //TODO: make imageID only settable by owner
    function verifySnarkProof(
        bytes calldata _journal,
        bytes32 imageId,
        bytes calldata _seal
    ) external {
        IRiscZeroVerifier(verifier).verify(
                _seal,
                imageId,
                sha256(_journal)
            );
        
        emit SnarkProofVerified(_journal, _seal);
    }

    function setVerifier(address _verifier) external {
        address oldVerifier = verifier; 
        verifier = _verifier;
        emit VerifierChanged(oldVerifier, verifier);
    }

}