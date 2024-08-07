// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IStakeRootCompendium.sol";
import "../interfaces/IStrategy.sol";
import "../libraries/Merkle.sol";

import "@risc0-ethereum/contracts/src/IRiscZeroVerifier.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";


contract StakeRootCompendium is IStakeRootCompendium, OwnableUpgradeable {

    uint32 public constant MAX_OPERATOR_SET_SIZE = 2048;
    uint32 public constant MAX_NUM_OPERATOR_SETS = 2048;
    uint32 public constant MAX_NUM_STRATEGIES = 16;

    IDelegationManager public delegation;
    IAVSDirectory public avsDirectory;

    address public verifier;
    bytes32 public imageId;

    mapping(uint32 => mapping(IStrategy => uint96)) public operatorSetIdToStrategyToMultiplier;

    modifier isOperatorSetAVS(address avs) {
        require(avsDirectory.isOperatorSetAVS(avs), "StakeRootCompendium: is not operator set AVS");
        _;
    }

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
        require(operatorSetRoots.length <= MAX_NUM_OPERATOR_SETS, "AVSSyncTree.getStakeRoot: operatorSet size limit exceeded");
    
        bytes32[] memory operatorSetLeaves = new bytes32[](avss.length);
        for (uint256 i = 0; i < avss.length; i++) {
            operatorSetLeaves[i] = keccak256(abi.encodePacked(avss[i], operatorSetRoots[i]));
        }
        return Merkle.merkleizeKeccak256(operatorSetLeaves);
    }

    function getOperatorSetRoot(address avs, uint32 operatorSetId, address[] calldata operators, IStrategy[] calldata strategies) external view isOperatorSetAVS(avs) returns (bytes32) {
        require(operators.length <= MAX_OPERATOR_SET_SIZE, "AVSSyncTree._verifyOperatorStatus: operator set too large");

        bytes32[] memory operatorLeaves = new bytes32[](operators.length);
        for (uint256 i = 0; i < operators.length; i++) {
            require(delegation.isOperator(operators[i]), "AVSSyncTree.getOperatorSetRoot: operator not registered");
            require(avsDirectory.avsOperatorStatus(avs,operators[i]) == IAVSDirectory.OperatorAVSRegistrationStatus.REGISTERED, "AVSSyncTree.getOperatorSetRoot: operator not registered to AVS");
            require(avsDirectory.isMember(avs,operators[i],operatorSetId), "AVSSyncTree.getOperatorSetRoot: operator not in operator set");

            // shares associated with this operator and these strategies
            uint256 weightedStrategyShareSum = _calculateWeightedStrategyShareSum(operators[i], strategies, getMultipliers(operatorSetId, strategies));
            bytes32 operatorRoot = keccak256(abi.encodePacked(operators[i], weightedStrategyShareSum));
            operatorLeaves[i] = keccak256(abi.encodePacked(operators[i], operatorRoot));     
        }
        return Merkle.merkleizeKeccak256(operatorLeaves);
    }

    function verifySnarkProof(
        bytes calldata _journal,
        bytes calldata _seal
    ) external {
        IRiscZeroVerifier(verifier).verify(
                _seal,
                imageId,
                sha256(_journal)
            );
        
        emit SnarkProofVerified(_journal, _seal);
    }

    //TODO: getting an error here:
    // Error: Unimplemented feature (/solidity/libsolidity/codegen/ArrayUtils.cpp:240):
    // Copying of type struct IStakeRootCompendium.StrategyAndMultiplier calldata[] calldata to storage not yet supported.
    function setStrategiesAndMultipliers(
        uint32 operatorSetId,
        StrategyAndMultiplier[] calldata strategiesAndMultipliers
    ) external isOperatorSetAVS(msg.sender) {
        for (uint256 i = 0; i < strategiesAndMultipliers.length; i++) {
            operatorSetIdToStrategyToMultiplier[operatorSetId][strategiesAndMultipliers[i].strategy] = strategiesAndMultipliers[i].multiplier;
        }
    }


    function getMultipliers(
        uint32 operatorSetId,
        IStrategy[] calldata strategies
    ) public view returns (uint96[] memory) {
        uint96[] memory multipliers = new uint96[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            multipliers[i] = operatorSetIdToStrategyToMultiplier[operatorSetId][strategies[i]];
        }
        return multipliers;
    }

    function setVerifier(address _verifier) external onlyOwner {
        address oldVerifier = verifier; 
        verifier = _verifier;
        emit VerifierChanged(oldVerifier, verifier);
    }

    function setImageId(bytes32 _imageId) external onlyOwner {
        bytes32 oldImageId = imageId;
        imageId = _imageId;
        emit ImageIdChanged(oldImageId, imageId);
    }

    function _calculateWeightedStrategyShareSum(address operator, IStrategy[] memory strategies, uint96[] memory multipliers) internal view returns (uint256) {
        require(strategies.length == multipliers.length, "AVSSyncTree._calculateWeightedStrategyShareSum: strategies and multipliers length mismatch");
        uint256[] memory shares = _retrieveStrategyShares(operator, strategies);
        uint256 weightedSum = 0;
        for (uint256 i = 0; i < shares.length; i++) {
            weightedSum += shares[i] * multipliers[i];
        }
        return weightedSum;
    }


    function _retrieveStrategyShares(address operator, IStrategy[] memory strategies) internal view returns (uint256[] memory) {
        require(strategies.length <= MAX_NUM_STRATEGIES, "AVSSyncTree._retrieveStrategyShares: too many strategies");
        return delegation.getOperatorShares(operator, strategies);
    }
}