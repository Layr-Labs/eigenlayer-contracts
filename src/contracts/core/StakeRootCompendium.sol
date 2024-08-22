// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IStrategy.sol";
import "../interfaces/IStakeRootCompendium.sol";
import "../libraries/Merkle.sol";

import "@risc0-ethereum/IRiscZeroVerifier.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import "../libraries/Checkpoints.sol";

contract StakeRootCompendium is IStakeRootCompendium, OwnableUpgradeable {
    using Checkpoints for Checkpoints.History;
    using EnumerableMap for EnumerableMap.AddressToUintMap;

    /// @notice the maximum number of operators that can be in an operator set in the StakeTree
    uint32 public constant MAX_OPERATOR_SET_SIZE = 2048;
    /// @notice the maximum number of operator sets that can be in the StakeTree
    uint32 public constant MAX_NUM_OPERATOR_SETS = 2048;
    /// @notice the maximum number of strategies that each operator set in the StakeTree can use to weight their operator stakes
    uint32 public constant MAX_NUM_STRATEGIES = 20;
    /// @notice the placeholder index used for operator sets that are removed from the StakeTree
    uint32 public constant REMOVED_INDEX = type(uint32).max;

    /// @notice the delegation manager contract
    IDelegationManager public immutable delegationManager;
    /// @notice the AVS directory contract
    IAVSDirectory public immutable avsDirectory;
    /// @notice the period of time within which a root can be marked as blacklisted
    uint32 public immutable blacklistWindow;

    /// @notice map from operator set to a trace of their index over time
    mapping(address => mapping(uint32 => Checkpoints.History)) internal operatorSetToIndex;
    /// @notice list of operator sets that have been configured to be in the StakeTree
    IAVSDirectory.OperatorSet[] public operatorSets;
    /// @notice the number of operator sets that have been configured to be in the StakeTree
    mapping(address => mapping(uint32 => EnumerableMap.AddressToUintMap)) internal operatorSetToStrategyAndMultipliers;

    /// @notice the extraData for each operator in each operator set
    mapping(address => mapping(uint32 => mapping(address => bytes32))) internal extraDatas;

    /// @notice the verifier contract that will be used to verify snark proofs
    address public verifier;
    /// @notice the id of the program being verified when roots are posted
    bytes32 public imageId;

    StakeRootSubmission[] public stakeRootSubmissions;

    constructor(
        IDelegationManager _delegationManager,
        IAVSDirectory _avsDirectory,
        uint32 _blacklistWindow
    ) {
        // _disableInitializers();
        delegationManager = _delegationManager;
        avsDirectory = _avsDirectory;
        blacklistWindow = _blacklistWindow;
    }

    function initialize(address initialOwner) public initializer {
        __Ownable_init();
        _transferOwnership(initialOwner);
    }

    /// CALLED BY AVS

    /// @inheritdoc IStakeRootCompendium
    function addStrategiesAndMultipliers(
        uint32 operatorSetId,
        StrategyAndMultiplier[] calldata strategiesAndMultipliers
    ) external {
        require(strategiesAndMultipliers.length > 0, "StakeRootCompendium.setStrategiesAndMultipliers: no strategies and multipliers provided");
        require(avsDirectory.isOperatorSet(msg.sender, operatorSetId), "StakeRootCompendium.setStrategiesAndMultipliers: operator set does not exist");
        uint256 lengthBefore = operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].length();
        // if the operator set has been configured to have a positive number of strategies, increment the number of configured operator sets
        if (lengthBefore == 0) {
            require(operatorSets.length < MAX_NUM_OPERATOR_SETS, "StakeRootCompendium.setStrategiesAndMultipliers: too many operator sets");
            operatorSets.push(IAVSDirectory.OperatorSet(msg.sender, operatorSetId));
            operatorSetToIndex[msg.sender][operatorSetId].push(uint32(block.timestamp), uint208(operatorSets.length - 1));
        }
        
        // set the strategies and multipliers for the operator set
        for (uint256 i = 0; i < strategiesAndMultipliers.length; i++) {
            operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].set(address(strategiesAndMultipliers[i].strategy), uint256(strategiesAndMultipliers[i].multiplier));
        }
        require(operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].length() <= MAX_NUM_STRATEGIES, "StakeRootCompendium.setStrategiesAndMultipliers: too many strategies");
    }

    /// @inheritdoc IStakeRootCompendium
    function removeStrategiesAndMultipliers(
        uint32 operatorSetId,
        IStrategy[] calldata strategies
    ) external {        
        // remove the strategies and multipliers for the operator set
        for (uint256 i = 0; i < strategies.length; i++) {
            require(operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].remove(address(strategies[i])), "StakeRootCompendium.removeStrategiesAndMultipliers: strategy not found");
        }

        // if the operator set has been configured to have no strategies, decrement the number of configured operator sets
        if(operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].length() == 0) {
            IAVSDirectory.OperatorSet memory operatorSet = operatorSets[operatorSets.length - 1];
            uint224 operatorSetIndex = operatorSetToIndex[msg.sender][operatorSetId].latest();
            operatorSets[operatorSetIndex] = operatorSet;
            operatorSets.pop();
            
            // update the index of the operator set
            operatorSetToIndex[msg.sender][operatorSetId].push(uint32(block.timestamp), REMOVED_INDEX);
            operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].push(uint32(block.timestamp), operatorSetIndex);
        }
    }

    /// @inheritdoc IStakeRootCompendium
    function setExtraData(
	   uint32 operatorSetId,
	   address operator,
	   uint32 timestamp,
	   bytes32 extraData
	) external {
        (bool exists,,uint224 index) = operatorSetToIndex[msg.sender][operatorSetId].latestCheckpoint();
        require(exists && index != REMOVED_INDEX, "StakeRootCompendium.setExtraData: operatorSet is not in stakeTree");
        extraDatas[msg.sender][operatorSetId][operator] = extraData;
    }

    // STAKE ROOT CALCULATION

    /// @inheritdoc IStakeRootCompendium
    function getStakeRoot(IAVSDirectory.OperatorSet[] calldata operatorSetsInStakeTree, bytes32[] calldata operatorSetRoots) external view returns (bytes32) {
        require(operatorSets.length == operatorSetsInStakeTree.length, "StakeRootCompendium.getStakeRoot: operatorSets vs. operatorSetsInStakeTree length mismatch");
        require(operatorSetsInStakeTree.length == operatorSetRoots.length, "StakeRootCompendium.getStakeRoot: operatorSetsInStakeTree vs. operatorSetRoots mismatch");
        for (uint256 i = 0; i < operatorSetsInStakeTree.length; i++) {
            require(operatorSets[i].avs == operatorSetsInStakeTree[i].avs, "StakeRootCompendium.getStakeRoot: operatorSets vs. operatorSetsInStakeTree avs mismatch");
            require(operatorSets[i].operatorSetId == operatorSetsInStakeTree[i].operatorSetId, "StakeRootCompendium.getStakeRoot: operatorSets vs. operatorSetsInStakeTree operatorSetId mismatch");
        }
        return Merkle.merkleizeKeccak256(operatorSetRoots);
    }

    /// @inheritdoc IStakeRootCompendium
    function getOperatorSetRoot(
        IAVSDirectory.OperatorSet calldata operatorSet, 
        address[] calldata operators
    )
        external view 
        returns (bytes32) 
    {
        require(avsDirectory.isOperatorSet(operatorSet.avs, operatorSet.operatorSetId), "StakeRootCompendium.getOperatorSetRoot: operator set does not exist");
        require(operators.length <= MAX_OPERATOR_SET_SIZE, "AVSSyncTree._verifyOperatorStatus: operator set too large");
        require(operators.length == avsDirectory.operatorSetMemberCount(operatorSet.avs, operatorSet.operatorSetId), "AVSSyncTree.getOperatorSetRoot: operator set size mismatch");

        bytes32[] memory operatorLeaves = new bytes32[](operators.length);
        uint160 prevOperator = 0;
        for (uint256 i = 0; i < operators.length; i++) {
            require(avsDirectory.isMember(operators[i], operatorSet), "AVSSyncTree.getOperatorSetRoot: operator not in operator set");
            
            // ensure that operators are sorted
            require(uint160(operators[i]) > prevOperator, "AVSSyncTree.getOperatorSetRoot: operators not sorted");
            prevOperator = uint160(operators[i]);

            // calculate the weighted sum of the operator's shares for the strategies given the multipliers
            IStrategy[] memory strategies = new IStrategy[](operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length());
            uint256[] memory multipliers = new uint256[](strategies.length);
            for (uint256 j = 0; j < strategies.length; j++) {
                (address strategy, uint256 multiplier) = operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].at(j);
                strategies[j] = IStrategy(strategy);
                multipliers[j] = multiplier;
            }

            uint256 delegatedStake = 0;
            uint256 slashableStake = 0;
            {
                uint256[] memory delegatedShares = delegationManager.getOperatorShares(operators[i], strategies);
                uint24[] memory operatorSlashablePPM = avsDirectory.getSlashablePPM(operators[i], operatorSet, strategies, uint32(block.timestamp), true);
                for (uint256 j = 0; j < strategies.length; j++) {
                    delegatedStake += delegatedShares[j] * multipliers[j];
                    slashableStake += delegatedShares[j] * multipliers[j] * operatorSlashablePPM[j] / 1e6;
                }
            }

            operatorLeaves[i] =  keccak256(abi.encodePacked(operators[i], delegatedStake, slashableStake, extraDatas[operatorSet.avs][operatorSet.operatorSetId][operators[i]]));    
        }
        return Merkle.merkleizeKeccak256(operatorLeaves);
    }

    /// POSTING ROOTS AND BLACKLISTING

    /// @inheritdoc IStakeRootCompendium
    function verifyStakeRoot(uint32 calculationTimestamp, bytes32 stakeRoot, Proof calldata proof) external {
        // TODO: verify proof

        _postStakeRoot(calculationTimestamp, stakeRoot, false);
    }
    
    
    /// @inheritdoc IStakeRootCompendium
    function blacklistStakeRoot(uint32 submissionIndex) external onlyOwner {
        // TODO: this should not be onlyOwner
        require(!stakeRootSubmissions[submissionIndex].blacklisted, "StakeRootCompendium.blacklistStakeRoot: stakeRoot already blacklisted");
        require(stakeRootSubmissions[submissionIndex].submissionTimestamp + blacklistWindow >= block.timestamp, "StakeRootCompendium.blacklistStakeRoot: stakeRoot cannot be blacklisted");
        require(!stakeRootSubmissions[submissionIndex].forcePosted, "StakeRootCompendium.blacklistStakeRoot: stakeRoot was force posted");
        stakeRootSubmissions[submissionIndex].blacklisted = true;
    }
    
    /// @inheritdoc IStakeRootCompendium
    function forcePostStakeRoot(uint32 calculationTimestamp, bytes32 stakeRoot) external onlyOwner {
        _postStakeRoot(calculationTimestamp, stakeRoot, true);
    }

    /// PERMISSIONED SETTERS

    /**
     * @notice sets the verifier contract that will be used to verify snark proofs
     * @param _verifier the address of the verifier contract
     * @dev only callable by the owner
     */
    function setVerifier(address _verifier) external onlyOwner {
        address oldVerifier = verifier; 
        verifier = _verifier;
        emit VerifierChanged(oldVerifier, verifier);
    }

    /**
     * @notice sets/changes the id of the program being verified when roots are posted
     * @param _imageId the new imageId to set
     * @dev only callable by the owner
     */
    function setImageId(bytes32 _imageId) external onlyOwner {
        bytes32 oldImageId = imageId;
        imageId = _imageId;
        emit ImageIdChanged(oldImageId, imageId);
    }

    /// INTERNAL FUNCTIONS

    function _postStakeRoot(uint32 calculationTimestamp, bytes32 stakeRoot, bool forcePosted) internal {
        uint256 stakeRootSubmissionsLength = stakeRootSubmissions.length;
        if (stakeRootSubmissionsLength != 0) {
            require(stakeRootSubmissions[stakeRootSubmissionsLength - 1].calculationTimestamp < calculationTimestamp, "StakeRootCompendium._postStakeRoot: calculationTimestamp must be greater than the last posted calculationTimestamp");
        }

        stakeRootSubmissions.push(StakeRootSubmission({
            stakeRoot: stakeRoot,
            calculationTimestamp: calculationTimestamp,
            submissionTimestamp: uint32(block.timestamp),
            blacklisted: false,
            forcePosted: forcePosted
        }));

        // todo: emit events
    }
}