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
    uint256 public numAVSs;

    mapping(uint32 => mapping(IStrategy => uint96)) public operatorSetIdToStrategyToMultiplier;
    mapping(address => bool) public isAVS;

    modifier isOperatorSet(address avs, uint32 operatorSetId) {
        require(avsDirectory.isOperatorSet(avs, operatorSetId), "StakeRootCompendium: operator set does not exist");
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

     function getStakeRoot(StakeRootLeaf[] calldata stakeRootLeaves) external view returns (bytes32) {
        require(stakeRootLeaves.length <= MAX_NUM_OPERATOR_SETS, "AVSSyncTree.getStakeRoot: operatorSet size limit exceeded");
        require(stakeRootLeaves.length == numAVSs, "AVSSyncTree.getStakeRoot: more leaves than AVSs that have set their strategies and multipliers");
    
        bytes32[] memory operatorSetLeaves = new bytes32[](stakeRootLeaves.length);
        for (uint256 i = 0; i < stakeRootLeaves.length; i++) {
            operatorSetLeaves[i] = keccak256(abi.encodePacked(stakeRootLeaves[i].avs, stakeRootLeaves[i].operatorSetRoot, stakeRootLeaves[i].operatorSetId));
        }
        return Merkle.merkleizeKeccak256(operatorSetLeaves);
    }

    //Note: assumes that the operator is registered to the AVS and is registered with the delegationManager.
    function getOperatorSetRoot(
        address avs,
        uint32 operatorSetId, 
        address[] calldata operators, 
        IStrategy[] calldata strategies
    )
        external view 
        isOperatorSet(avs, operatorSetId) 
        returns (bytes32) 
    {
         
        require(operators.length <= MAX_OPERATOR_SET_SIZE, "AVSSyncTree._verifyOperatorStatus: operator set too large");
        require(operators.length == avsDirectory.operatorSetMemberCount(avs, operatorSetId), "AVSSyncTree.getOperatorSetRoot: operator set size mismatch");

        bytes32[] memory operatorLeaves = new bytes32[](operators.length);
        uint160 prevOperator = 0;
        for (uint256 i = 0; i < operators.length; i++) {
            require(avsDirectory.isMember(avs,operators[i], operatorSetId), "AVSSyncTree.getOperatorSetRoot: operator not in operator set");
            require(uint160(operators[i]) > prevOperator, "AVSSyncTree.getOperatorSetRoot: operators not sorted");
            
            uint96[] memory multipliers = new uint96[](strategies.length);
            for (uint256 j = 0; j < strategies.length; j++) {
                multipliers[j] = operatorSetIdToStrategyToMultiplier[operatorSetId][strategies[i]];
            }
            // shares associated with this operator and these strategies
            operatorLeaves[i] =  keccak256(abi.encodePacked(operators[i], _calculateWeightedStrategyShareSum(operators[i], strategies, multipliers)));    
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
    ) external isOperatorSet(msg.sender, operatorSetId) {
        if(!isAVS[msg.sender]) {
            isAVS[msg.sender] = true;
            //if the AVS has never set its strategies and multipliers before, increment the number of AVSs
            numAVSs++;
        }
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
        require(strategies.length <= MAX_NUM_STRATEGIES, "AVSSyncTree._retrieveStrategyShares: too many strategies");
        uint256[] memory shares = delegation.getOperatorShares(operator, strategies);
        uint256 weightedSum = 0;
        for (uint256 i = 0; i < shares.length; i++) {
            weightedSum += shares[i] * multipliers[i];
        }
        return weightedSum;
    }
}