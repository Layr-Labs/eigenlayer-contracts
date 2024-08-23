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

    /// @notice the minimum balance that must be maintained for an operatorSet
    uint256 public constant MIN_DEPOSIT_BALANCE = 0.1 ether;

    /// @notice the delegation manager contract
    IDelegationManager public immutable delegationManager;
    /// @notice the AVS directory contract
    IAVSDirectory public immutable avsDirectory;

    /// @notice the interval at which proofs can be posted, to not overcharge the operatorSets
    uint32 public immutable proofInterval;
    /// @notice the period of time within which a root can be marked as blacklisted
    uint32 public immutable blacklistWindow;

    /// @notice charge per strategy per proof
    uint256 public charge;

    /// @notice deposit balance to be deducted for operatorSets
    mapping(address => mapping(uint32 => DepositBalanceInfo)) public depositBalanceInfo;

    /// @notice map from operator set to a trace of their index over time
    mapping(address => mapping(uint32 => Checkpoints.History)) internal operatorSetToIndex;
    /// @notice list of operator sets that have been configured to be in the StakeTree
    IAVSDirectory.OperatorSet[] public operatorSets;

    /// @notice the total number of strategies across all operator sets over time
    Checkpoints.History internal totalStrategies;
    /// @notice the number of operator sets that have been configured to be in the StakeTree
    mapping(address => mapping(uint32 => EnumerableMap.AddressToUintMap)) internal operatorSetToStrategyAndMultipliers;
    /// @notice the extraData for each operator in each operator set
    mapping(address => mapping(uint32 => mapping(address => bytes32))) internal extraDatas;

    /// @notice the verifier contract that will be used to verify snark proofs
    address public verifier;
    /// @notice the id of the program being verified when roots are posted
    bytes32 public imageId;

    /// @notice the index of the latest stake root submission that has been charged
    uint256 public latestChargedSubmissionIndex;
    /// @notice the stake root submissions that have been posted
    StakeRootSubmission[] public stakeRootSubmissions;

    constructor(
        IDelegationManager _delegationManager,
        IAVSDirectory _avsDirectory,
        uint32 _proofInterval,
        uint32 _blacklistWindow
    ) {
        // _disableInitializers();
        delegationManager = _delegationManager;
        avsDirectory = _avsDirectory;
        proofInterval = _proofInterval;
        blacklistWindow = _blacklistWindow;
    }

    function initialize(address initialOwner, uint256 _charge) public initializer {
        __Ownable_init();
        _transferOwnership(initialOwner);
        _setCharge(_charge);
    }

    /// OPERATORSET CONFIGURATION

    /// @inheritdoc IStakeRootCompendium
    function depositForOperatorSet(IAVSDirectory.OperatorSet calldata operatorSet) external payable {
        require(avsDirectory.isOperatorSet(operatorSet.avs, operatorSet.operatorSetId), "StakeRootCompendium.depositForOperatorSet: operator set does not exist");
        depositBalanceInfo[operatorSet.avs][operatorSet.operatorSetId].balance += msg.value;
        require(
            depositBalanceInfo[operatorSet.avs][operatorSet.operatorSetId].balance >= 2 * MIN_DEPOSIT_BALANCE, 
            "StakeRootCompendium.depositForOperatorSet: depositer must have at least 2x the minimum balance on deposit"
        );

        // update the deposit balance for the operator set whenever a deposit is made
        _updateDepositBalanceInfo(operatorSet, true);
    }

    /// @inheritdoc IStakeRootCompendium
    function addStrategiesAndMultipliers(
        uint32 operatorSetId,
        StrategyAndMultiplier[] calldata strategiesAndMultipliers
    ) external {
        require(strategiesAndMultipliers.length > 0, "StakeRootCompendium.setStrategiesAndMultipliers: no strategies and multipliers provided");
        require(avsDirectory.isOperatorSet(msg.sender, operatorSetId), "StakeRootCompendium.setStrategiesAndMultipliers: operator set does not exist");

        IAVSDirectory.OperatorSet memory operatorSet = IAVSDirectory.OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
        uint224 lengthBefore = uint224(operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].length());
        // if the operator set has been configured to have a positive number of strategies, increment the number of configured operator sets
        if (lengthBefore == 0) {
            require(operatorSets.length < MAX_NUM_OPERATOR_SETS, "StakeRootCompendium.setStrategiesAndMultipliers: too many operator sets");
            operatorSetToIndex[msg.sender][operatorSetId].push(uint32(block.timestamp), uint208(operatorSets.length));
            operatorSets.push(operatorSet);
        }
        
        // update the deposit balance for the operator set whenever number of strategies is changed
        _updateDepositBalanceInfo(operatorSet, true);
        
        // set the strategies and multipliers for the operator set
        for (uint256 i = 0; i < strategiesAndMultipliers.length; i++) {
            operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].set(address(strategiesAndMultipliers[i].strategy), uint256(strategiesAndMultipliers[i].multiplier));
        }
        uint224 lengthAfter = uint224(operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].length());
        require(lengthAfter <= MAX_NUM_STRATEGIES, "StakeRootCompendium.setStrategiesAndMultipliers: too many strategies");
        totalStrategies.push(uint32(block.timestamp), totalStrategies.latest() + lengthAfter - lengthBefore);
    }

    /// @inheritdoc IStakeRootCompendium
    function removeStrategiesAndMultipliers(
        uint32 operatorSetId,
        IStrategy[] calldata strategies
    ) external {   
        IAVSDirectory.OperatorSet memory operatorSet = IAVSDirectory.OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
        // update the deposit balance for the operator set whenever number of strategies is changed
        _updateDepositBalanceInfo(operatorSet, true);

        uint224 lengthBefore = uint224(operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].length());
        // remove the strategies and multipliers for the operator set
        for (uint256 i = 0; i < strategies.length; i++) {
            require(operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].remove(address(strategies[i])), "StakeRootCompendium.removeStrategiesAndMultipliers: strategy not found");
        }

        uint224 lengthAfter = uint224(operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].length());
        // if the operator set has been configured to have no strategies, decrement the number of configured operator sets
        if(lengthAfter == 0) {
            _removeOperatorSet(operatorSet);
        }
        totalStrategies.push(uint32(block.timestamp), totalStrategies.latest() - lengthBefore + lengthAfter);
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
        _updateDepositBalanceInfo(IAVSDirectory.OperatorSet({avs: msg.sender, operatorSetId: operatorSetId}), true);
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
        address prevOperator;
        for (uint256 i = 0; i < operators.length; i++) {
            require(avsDirectory.isMember(operators[i], operatorSet), "AVSSyncTree.getOperatorSetRoot: operator not in operator set");
            
            // ensure that operators are sorted
            require(operators[i] > prevOperator, "AVSSyncTree.getOperatorSetRoot: operators not sorted");
            prevOperator = operators[i];

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
    function verifyStakeRoot(uint32 calculationTimestamp, bytes32 stakeRoot, address chargeRecipient, Proof calldata proof) external {
        // TODO: verify proof

        _postStakeRoot(calculationTimestamp, stakeRoot, chargeRecipient, false);
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
        _postStakeRoot(calculationTimestamp, stakeRoot, address(0), true);
    }

    /// CHARGE MANAGEMENT

    /// @inheritdoc IStakeRootCompendium
    function updateDepositBalanceInfos(IAVSDirectory.OperatorSet[] calldata operatorSetsToUpdate) external {
        uint256 penalty = 0;
        for(uint256 i = 0; i < operatorSetsToUpdate.length; i++) {
            penalty += _updateDepositBalanceInfo(operatorSetsToUpdate[i], false);
        }
        if (penalty > 0) {
            payable(msg.sender).transfer(penalty);
        }
    }

    /// @inheritdoc IStakeRootCompendium
    function processCharges(uint256 numToCharge) external {
        uint256 latestChargedSubmissionIndexMemory = latestChargedSubmissionIndex;
        uint256 endIndex = latestChargedSubmissionIndexMemory + numToCharge;
        if (endIndex > stakeRootSubmissions.length) {
            endIndex = stakeRootSubmissions.length;
        }

        address prevRecipient;
        uint256 totalCharge;
        for (uint256 i = latestChargedSubmissionIndexMemory; i < endIndex; i++) {
            StakeRootSubmission memory stakeRootSubmission = stakeRootSubmissions[i];
            // if the stakeRootSubmission is blacklisted or force posted, skip it
            if (
                block.timestamp < stakeRootSubmission.submissionTimestamp + blacklistWindow ||
                stakeRootSubmission.blacklisted || 
                stakeRootSubmission.forcePosted
            ) {
                continue;
            }
            // if the charge recipient has changed, transfer the total charge to the previous recipient
            if(stakeRootSubmission.chargeRecipient != prevRecipient) {
                if (totalCharge > 0) {
                    payable(prevRecipient).transfer(totalCharge);
                }
                totalCharge = 0;
                prevRecipient = stakeRootSubmission.chargeRecipient;
            }
            // total charge is the charge per strategy per proof times the number of strategies at the time of proof
            totalCharge += charge * totalStrategies.upperLookup(stakeRootSubmissions[i].calculationTimestamp);
        }
    }

    /// PERMISSIONED SETTERS

    /// @inheritdoc IStakeRootCompendium
    function setVerifier(address _verifier) external onlyOwner {
        address oldVerifier = verifier; 
        verifier = _verifier;
        emit VerifierChanged(oldVerifier, verifier);
    }

    /// @inheritdoc IStakeRootCompendium
    function setImageId(bytes32 _imageId) external onlyOwner {
        bytes32 oldImageId = imageId;
        imageId = _imageId;
        emit ImageIdChanged(oldImageId, imageId);
    }

    /// VIEW FUNCTIONS

    /// @notice the stake root submissions that have been posted
    function getStakeRootSubmissionAt(uint32 index) external view returns (StakeRootSubmission memory) {
        return stakeRootSubmissions[index];
    }

    /// @inheritdoc IStakeRootCompendium
    function getDepositBalance(IAVSDirectory.OperatorSet memory operatorSet) public view returns (uint256 balance, uint256 penalty) {
        // the total charge for the operator set is the charge per strategy per proof
        uint256 totalCharge = 
            charge * 
            operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length() * 
            (block.timestamp - depositBalanceInfo[operatorSet.avs][operatorSet.operatorSetId].lastUpdatedAt) / proofInterval;
        uint256 storedBalance = depositBalanceInfo[operatorSet.avs][operatorSet.operatorSetId].balance;

        // if the charge would take the balance below the minimum deposit balance, return 0
        balance = storedBalance < totalCharge + MIN_DEPOSIT_BALANCE ? 0 : storedBalance - totalCharge;
        // if the balance is 0, then the charger may be able to claim the penalty for removing the operatorSet from the stakeTree
        penalty = balance == 0 ? totalCharge + MIN_DEPOSIT_BALANCE - storedBalance : 0;
        return (balance, penalty);
    }

    /// Misc

    receive() external payable {}

    /// INTERNAL FUNCTIONS

    function _removeOperatorSet(IAVSDirectory.OperatorSet memory operatorSet) internal {
        IAVSDirectory.OperatorSet memory substituteOperatorSet = operatorSets[operatorSets.length - 1];
        uint224 operatorSetIndex = operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].latest();
        operatorSets[operatorSetIndex] = substituteOperatorSet;
        operatorSets.pop();
        
        // update the index of the operator sets
        operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].push(uint32(block.timestamp), REMOVED_INDEX);
        operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].push(uint32(block.timestamp), operatorSetIndex);
    }

    function _setCharge(uint256 _charge) internal {
        charge = _charge;
        // TODO: emit event
    }

    // updates the deposit balance for the operator set and returns the penalty if the operator set has fallen below the minimum deposit balance
    function _updateDepositBalanceInfo(IAVSDirectory.OperatorSet memory operatorSet, bool sendPenalty) internal returns(uint256) {
        (uint256 depositBalance, uint256 penalty) = getDepositBalance(operatorSet);
        depositBalanceInfo[operatorSet.avs][operatorSet.operatorSetId].balance = depositBalance;
        depositBalanceInfo[operatorSet.avs][operatorSet.operatorSetId].lastUpdatedAt = uint32(block.timestamp);

        // if the operatorSet has fallen below the minimum deposit balance, remove it from the stakeTree
        if (penalty > 0) {
            _removeOperatorSet(operatorSet);
            if (sendPenalty) {
                payable(msg.sender).transfer(penalty);
            }
        }

        return penalty;
    }

    function _postStakeRoot(uint32 calculationTimestamp, bytes32 stakeRoot, address chargeRecipient, bool forcePosted) internal {
        require(calculationTimestamp % proofInterval == 0, "StakeRootCompendium._postStakeRoot: calculationTimestamp must be a multiple of proofInterval");

        uint256 stakeRootSubmissionsLength = stakeRootSubmissions.length;
        if (stakeRootSubmissionsLength != 0) {
            require(
                stakeRootSubmissions[stakeRootSubmissionsLength - 1].calculationTimestamp < calculationTimestamp, 
                "StakeRootCompendium._postStakeRoot: calculationTimestamp must be greater than the last posted calculationTimestamp"
            );
        }

        stakeRootSubmissions.push(StakeRootSubmission({
            stakeRoot: stakeRoot,
            chargeRecipient: msg.sender,
            calculationTimestamp: calculationTimestamp,
            submissionTimestamp: uint32(block.timestamp),
            blacklisted: false,
            forcePosted: forcePosted
        }));

        // todo: emit events
    }
}