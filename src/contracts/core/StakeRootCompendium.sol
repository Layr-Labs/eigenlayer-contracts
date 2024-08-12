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



contract StakeRootCompendium is IStakeRootCompendium, OwnableUpgradeable {
    using EnumerableMap for EnumerableMap.AddressToUintMap;


    /// @notice the maximum number of operators that can be in an operator set in the StakeTree
    uint32 public constant MAX_OPERATOR_SET_SIZE = 2048;
    /// @notice the maximum number of operator sets that can be in the StakeTree
    uint32 public constant MAX_NUM_OPERATOR_SETS = 2048;
    /// @notice the maximum number of strategies that each operator set in the StakeTree can use to weight their operator stakes
    uint32 public constant MAX_NUM_STRATEGIES = 20;

    /// @notice the delegation manager contract
    IDelegationManager public delegation;
    /// @notice the AVS directory contract
    IAVSDirectory public avsDirectory;

    /// @notice the verifier contract that will be used to verify snark proofs
    address public verifier;
    /// @notice the id of the program being verified when roots are posted
    bytes32 public imageId;

    /// @notice the number of operator sets that have been configured to be in the StakeTree
    uint256 public numConfiguredOperatorSets;
    /// @notice the number of operator sets that have been configured to be in the StakeTree
    mapping(address => mapping(uint32 => EnumerableMap.AddressToUintMap)) internal operatorSetToStrategyAndMultipliers;

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
        require(stakeRootLeaves.length == numConfiguredOperatorSets, "AVSSyncTree.getStakeRoot: more leaves than AVSs that have set their strategies and multipliers");
    
        bytes32[] memory operatorSetLeaves = new bytes32[](stakeRootLeaves.length);

        IAVSDirectory.OperatorSet memory prevOperatorSet;
        for (uint256 i = 0; i < stakeRootLeaves.length; i++) {
            // ensure that stakeRootLeaves are sorted, first by AVS and then by operatorSetId
            require(
                stakeRootLeaves[i].operatorSet.avs > prevOperatorSet.avs 
                || (
                    stakeRootLeaves[i].operatorSet.avs == prevOperatorSet.avs 
                    && stakeRootLeaves[i].operatorSet.operatorSetId > prevOperatorSet.operatorSetId
                ), 
                "AVSSyncTree.getStakeRoot: stakeRootLeaves not sorted"
            );
            prevOperatorSet = stakeRootLeaves[i].operatorSet;

            operatorSetLeaves[i] = keccak256(abi.encodePacked(stakeRootLeaves[i].operatorSet.avs, stakeRootLeaves[i].operatorSet.operatorSetId, stakeRootLeaves[i].operatorSetRoot));
        }
        return Merkle.merkleizeKeccak256(operatorSetLeaves);
    }

    /**
     * @notice calculates the root of the operator set at the time of calling
     * @param operatorSet the operator set to get the root for
     * @param operators the operators in the operator set at the time of calling
     * @return the root of the operator set
     * @dev the operators must be sorted
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
            require(avsDirectory.isMember(operatorSet.avs, operators[i], operatorSet.operatorSetId), "AVSSyncTree.getOperatorSetRoot: operator not in operator set");
            
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

            uint256[] memory shares = delegation.getOperatorShares(operators[i], strategies);
            uint256 stake = 0;
            for (uint256 j = 0; j < strategies.length; j++) {
                stake += shares[j] * multipliers[j];
            }

            operatorLeaves[i] =  keccak256(abi.encodePacked(operators[i], stake));    
        }
        return Merkle.merkleizeKeccak256(operatorLeaves);
    }

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
        
        // set the strategies and multipliers for the operator set
        for (uint256 i = 0; i < strategiesAndMultipliers.length; i++) {
            operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].set(address(strategiesAndMultipliers[i].strategy), uint256(strategiesAndMultipliers[i].multiplier));
        }

        require(operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].length() <= MAX_NUM_STRATEGIES, "StakeRootCompendium.setStrategiesAndMultipliers: too many strategies");

        // if the operator set has been configured to have a positive number of strategies, increment the number of configured operator sets
        if (lengthBefore == 0) {
            numConfiguredOperatorSets++;
        }
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
            numConfiguredOperatorSets--;
        }
    }

    /// SNARK RELATED FUNCTIONS

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

    // /**
    //  * @notice 
    //  * @param _journal 
    //  * @param _seal 
    //  */
    // function verifySnarkProof(
    //     bytes calldata _journal,
    //     bytes calldata _seal
    // ) external {
    //     IRiscZeroVerifier(verifier).verify(
    //             _seal,
    //             imageId,
    //             sha256(_journal)
    //         );
        
    //     emit SnarkProofVerified(_journal, _seal);
    // }

    /// INTERNAL FUNCTIONS

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