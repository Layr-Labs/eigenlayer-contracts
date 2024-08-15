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
    using Checkpoints for Checkpoints.Trace208;
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
    /// @notice challenge period for stake roots
    uint32 public immutable challengePeriod;
    /// @notice The minimum amount of time between calculationTimestamps of consecutive claims
    uint32 public immutable minTimeSinceLastClaim;

    /// @notice length of the operatorSets array over time
    Checkpoints.Trace208 internal operatorSetsLength;
    /// @notice map from operator set to a trace of their index over time
    mapping(address => mapping(uint32 => Checkpoints.Trace208)) internal operatorSetToIndex;
    /// @notice list of operator sets that have been configured to be in the StakeTree
    IAVSDirectory.OperatorSet[] public operatorSets;
    /// @notice the number of operator sets that have been configured to be in the StakeTree
    mapping(address => mapping(uint32 => EnumerableMap.AddressToUintMap)) internal operatorSetToStrategyAndMultipliers;

    /// @notice whether an address is allowed to make stake root claims
    mapping(address => bool) public isClaimer;
    struct StakeRootClaim {
      bytes32 stakeRoot;
      address claimer;
      uint32 calculationTimestamp; // the timestamp the was generated against
      uint32 claimedTimestamp; // the timestamp the claim was made
      bool provenFraudlent; // whether the claim has been proven fraudulent
    }
    /// @notice map to claimed stake roots
    StakeRootClaim[] public stakeRootClaims;

    /// @notice the verifier contract that will be used to verify snark proofs
    address public verifier;
    /// @notice the id of the program being verified when roots are posted
    bytes32 public imageId;

    event ClaimerChanged(address claimer, bool allowed);

    constructor(
        IDelegationManager _delegationManager,
        IAVSDirectory _avsDirectory,
        uint32 _challengePeriod,
        uint32 _minTimeSinceLastClaim
    ) {
        // _disableInitializers();
        delegationManager = _delegationManager;
        avsDirectory = _avsDirectory;
        challengePeriod = _challengePeriod;
        minTimeSinceLastClaim = _minTimeSinceLastClaim;
    }

    function initialize(address initialOwner, address[] calldata claimers) public initializer {
        __Ownable_init();
        _transferOwnership(initialOwner);
        for (uint256 i = 0; i < claimers.length; i++) {
            _setClaimerStatus(claimers[i], true);
        }
    }

    /// CALLED BY AVS

    /**
     * @notice called by an AVS to set their strategies and multipliers used to determine stakes for stake roots
     * @param operatorSetId the id of the operator set to set the strategies and multipliers for
     * @param strategiesAndMultipliers the strategies and multipliers to set for the operator set
     * @dev msg.sender is used as the AVS in determining the operator set
     */
    function addStrategiesAndMultipliers(
        uint32 operatorSetId,
        StrategyAndMultiplier[] calldata strategiesAndMultipliers
    ) external {
        require(avsDirectory.isOperatorSet(msg.sender, operatorSetId), "StakeRootCompendium.setStrategiesAndMultipliers: operator set does not exist");
        uint256 lengthBefore = operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].length();
        // if the operator set has been configured to have a positive number of strategies, increment the number of configured operator sets
        if (lengthBefore == 0) {
            require(operatorSets.length < MAX_NUM_OPERATOR_SETS, "StakeRootCompendium.setStrategiesAndMultipliers: too many operator sets");
            operatorSets.push(IAVSDirectory.OperatorSet(msg.sender, operatorSetId));
            operatorSetToIndex[msg.sender][operatorSetId].push(uint32(block.timestamp), uint208(operatorSets.length - 1));
            operatorSetsLength.push(uint32(block.timestamp), uint208(operatorSets.length));
        }
        
        // set the strategies and multipliers for the operator set
        for (uint256 i = 0; i < strategiesAndMultipliers.length; i++) {
            operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].set(address(strategiesAndMultipliers[i].strategy), uint256(strategiesAndMultipliers[i].multiplier));
        }
        require(operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].length() <= MAX_NUM_STRATEGIES, "StakeRootCompendium.setStrategiesAndMultipliers: too many strategies");
    }

    /**
     * @notice called by an AVS to remove their strategies and multipliers used to determine stakes for stake roots
     * @param operatorSetId the id of the operator set to remove the strategies and multipliers for
     * @param strategies the strategies to remove for the operator set
     * @dev msg.sender is used as the AVS in determining the operator set
     */
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
            uint208 operatorSetIndex = operatorSetToIndex[msg.sender][operatorSetId].latest();
            operatorSets[operatorSetIndex] = operatorSet;
            operatorSets.pop();
            
            // update the index of the operator set
            operatorSetToIndex[msg.sender][operatorSetId].push(uint32(block.timestamp), REMOVED_INDEX);
            operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].push(uint32(block.timestamp), operatorSetIndex);
            operatorSetsLength.push(uint32(block.timestamp), uint208(operatorSets.length));
        }
    }

    // STAKE ROOT CALCULATION

    /**
     * @notice called offchain with the operator set roots ordered by the operator set index at the timestamp to calculate the stake root
     * @param timestamp the timestamp of the stake root
     * @param operatorSetRoots the ordered operator set roots (not verified)
     * @dev operatorSetRoots must be ordered by the operator set index at the timestamp
     */
    function getStakeRoot(uint32 timestamp, bytes32[] calldata operatorSetRoots) public view returns (bytes32) {
        require(operatorSetRoots.length == operatorSetsLength.upperLookupRecent(timestamp), "AVSSyncTree.getStakeRoot: more leaves than AVSs that have set their strategies and multipliers");
        return Merkle.merkleizeKeccak256(operatorSetRoots);
    }

    /**
	 * @notice Returns the operatorSet root for the given operatorSet with the given witnesses
	 * @param operatorSet the operatorSet to calculate the operator set root for
	 * @param operators the operators in the operatorSet
	 * @return the operatorSet root
	 * @dev the operators are verified to be distinct and the length must be the total 
	 * number of operators in the operatorSet at the time of calling
	 */
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

            uint256[] memory shares = delegationManager.getOperatorShares(operators[i], strategies);
            uint256 stake = 0;
            for (uint256 j = 0; j < strategies.length; j++) {
                stake += shares[j] * multipliers[j];
            }

            operatorLeaves[i] =  keccak256(abi.encodePacked(operators[i], stake));    
        }
        return Merkle.merkleizeKeccak256(operatorLeaves);
    }

    /// CALLED BY CLAIMER

    /**
     * @notice called by the claimer to claim a stake root
     * @param calculationTimestamp the timestamp of the state the stakeRoot was calculated against
     * @param operatorSetRoots the operator set roots ordered by the operator set index
     * @dev only callable by the claimer
     */
    function claimStakeRoot(uint32 calculationTimestamp, bytes32[] calldata operatorSetRoots) external {
        require(isClaimer[msg.sender], "StakeRootCompendium.claimStakeRoot: only claimers can claim stake roots");
        require(
            stakeRootClaims.length == 0 ||
            calculationTimestamp > stakeRootClaims[stakeRootClaims.length - 1].calculationTimestamp + minTimeSinceLastClaim, 
            "StakeRootCompendium.claimStakeRoot: stake root already claimed"
        );
        bytes32 stakeRoot = getStakeRoot(calculationTimestamp, operatorSetRoots);
        stakeRootClaims.push(StakeRootClaim({
            stakeRoot: stakeRoot,
            claimer: msg.sender,
            calculationTimestamp: calculationTimestamp,
            claimedTimestamp: uint32(block.timestamp),
            provenFraudlent: false
        }));
    }

    /**
     * @notice called by the claimer to challenge a stake root claim
     * @param claimIndex the index of the claim to challenge
     * @param operatorSet an operatorSet which an invalid operatorSetRoot was claimed
     * @param operatorSetRoot an incorrect operatorSetRoot that was claimed
     * @param merkleProof the merkle proof of the operatorSetRoot against the stakeRoot
     * @param journal the output of the program being verified (verifiedTimestamp, verifiedOperatorSet, verifiedOperatorSetRoot)
     * @param seal the snark proof of the program being verified
     */
    function challengeStakeRootClaim(
        uint32 claimIndex,
        IAVSDirectory.OperatorSet calldata operatorSet,
        bytes32 operatorSetRoot,
        bytes calldata merkleProof,
        bytes calldata journal,
        bytes calldata seal
    ) external {
        // verify that the claim is valid
        StakeRootClaim memory claim = stakeRootClaims[claimIndex];
        require(!claim.provenFraudlent, "StakeRootCompendium.verifySnarkProof: claim already proven fraudulent");

        // verify that the operator set root is correct
        uint256 operatorSetIndex = operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].upperLookupRecent(claim.calculationTimestamp);
        require(Merkle.verifyInclusionKeccak(
            merkleProof,
            claim.stakeRoot,
            operatorSetRoot,
            operatorSetIndex
        ), "StakeRootCompendium.verifySnarkProof: invalid operator set root");

        IRiscZeroVerifier(verifier).verify(
                seal,
                imageId,
                sha256(journal)
            );

        // TODO
        require(sha256(journal) != operatorSetRoot, "StakeRootCompendium.verifySnarkProof: claimed operator set root is correct");
        stakeRootClaims[claimIndex].provenFraudlent = true;
    }

    /// PERMISSIONED SETTERS

    /**
     * @notice changes the status of whether an address is allowed to make stake root claims
     * @param claimer the address to change the status of
     * @param allowed whether the address is allowed to make stake root claims
     * @dev only callable by the owner
     */
    function setClaimerStatus(address claimer, bool allowed) external onlyOwner {
        _setClaimerStatus(claimer, allowed);
    }

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

    function _setClaimerStatus(address claimer, bool allowed) internal {
        isClaimer[claimer] = allowed;
        emit ClaimerChanged(claimer, allowed);
    }
}