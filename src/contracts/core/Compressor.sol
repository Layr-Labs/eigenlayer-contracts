// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import "../interfaces/IStrategy.sol";
import "../libraries/Merkle.sol";
import "./CompressorStorage.sol";

contract Compressor is CompressorStorage {
    using Snapshots for Snapshots.History;
    using EnumerableMap for EnumerableMap.AddressToUintMap;
    using EnumerableSet for EnumerableSet.Bytes32Set;

    constructor(
        IDelegationManager _delegationManager,
        IAVSDirectory _avsDirectory,
        IAllocationManager _allocationManager,
        address _verifier,
        bytes32 _imageId
    )
        CompressorStorage(
            _delegationManager,
            _avsDirectory,
            _allocationManager,
            _verifier,
            _imageId
        )
    {}

    function initialize(
        address _owner,
        address _manager,
        address _confirmer,
        uint32 _proofIntervalSeconds
    ) public initializer {
        __Ownable_init();
        _transferOwnership(_owner);

        manager = _manager;
        confirmer = _confirmer;
        proofIntervalSeconds = _proofIntervalSeconds;

        compressedStates.push(CompressedState({
            root: bytes32(0),
            calculationTimestamp: 0,
            submissionTimestamp: uint32(block.timestamp),
            confirmed: false
        }));
    }

    modifier onlyManager() {
        require(msg.sender == manager, "Compressor: only manager.");
        _;
    }

    modifier onlyRegistered(OperatorSet memory operatorSet, bool isRegister) {
        require(
            isRegister == isRegistered(operatorSet),
            "Compressor: operator set is not in the corrrect registration state."
        );
        _;
    }

    /// OPERATORSET CONFIGURATION

    function register(
        OperatorSet calldata operatorSet
    ) external onlyManager onlyRegistered(operatorSet, false) {
        operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].push(
            uint32(block.timestamp), uint224(operatorSets.length)
        );
        operatorSets.push(operatorSet);
    }

    function addOrModifyStrategiesAndMultipliers(
        OperatorSet memory operatorSet,
        StrategyAndMultiplier[] calldata _strategiesAndMultipliers
    ) external onlyManager onlyRegistered(operatorSet, true) returns (uint256, uint256){
        EnumerableMap.AddressToUintMap storage strategiesAndMultipliers =
            operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId];

        uint256 numStrategiesBefore = strategiesAndMultipliers.length();
        // set the strategies and multipliers for the operator set
        for (uint256 i = 0; i < _strategiesAndMultipliers.length; i++) {
            strategiesAndMultipliers.set(
                address(_strategiesAndMultipliers[i].strategy), uint256(_strategiesAndMultipliers[i].multiplier)
            );
        }

        return (numStrategiesBefore, strategiesAndMultipliers.length());
    }

    
    function removeStrategiesAndMultipliers(
        OperatorSet memory operatorSet, 
        IStrategy[] calldata strategies
    ) external onlyManager onlyRegistered(operatorSet, true) returns (uint256, uint256) {
        EnumerableMap.AddressToUintMap storage strategiesAndMultipliers =
            operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId];

        // note below either all strategies are removed or none are removed and transaction reverts
        uint256 numStrategiesBefore = strategiesAndMultipliers.length();
        for (uint256 i = 0; i < strategies.length; i++) {
            require(strategiesAndMultipliers.remove(address(strategies[i])), NonexistentStrategy());
        }

        return(numStrategiesBefore, numStrategiesBefore - strategies.length);
    }

    
    function setOperatorSetExtraData(OperatorSet memory operatorSet, bytes32 extraData) external onlyManager onlyRegistered(operatorSet, true) {
        operatorSetExtraDatas[operatorSet.avs][operatorSet.operatorSetId] = extraData;
    }

    
    function setOperatorExtraData(OperatorSet memory operatorSet, address operator, bytes32 extraData) external onlyManager onlyRegistered(operatorSet, true) {
        operatorExtraDatas[operatorSet.avs][operatorSet.operatorSetId][operator] = extraData;
    }

    /// POSTING ROOTS AND BLACKLISTING

    function verifyCompression(
        uint256 calculationTimestamp,
        bytes32 root,
        address chargeRecipient,
        uint256 indexChargePerProof,
        bytes calldata proof
    ) external {
        require(
            calculationTimestamp % proofIntervalSeconds == 0,
            TimestampNotMultipleOfProofInterval()
        );
        // no length check here is ok because the initializer adds a default submission
        require(
            compressedStates[compressedStates.length - 1].calculationTimestamp != calculationTimestamp,
            TimestampAlreadyPosted()
        );
        
        //todo charging

        compressedStates.push(
            CompressedState({
                root: root,
                calculationTimestamp: uint32(calculationTimestamp),
                submissionTimestamp: uint32(block.timestamp),
                confirmed: false
            })
        );

        // interactions

        // note verify will be an external call, so adding to the end to apply the check, effect, interaction pattern to avoid reentrancy
        // TODO: verify proof
        // TODO: prevent race incentives and public mempool sniping, eg embed chargeRecipient in the proof
    }

    
    function confirmCompression(uint32 index, bytes32 root) external {
        require(msg.sender == confirmer, OnlyRootConfirmerCanConfirm());
        CompressedState storage compressedState = compressedStates[index];
        require(compressedState.root == root, StakeRootDoesNotMatch());
        require(!compressedState.confirmed, TimestampAlreadyConfirmed());
        compressedState.confirmed = true;
    }

    /// SET FUNCTIONS


    
    function setProofIntervalSeconds(
        uint32 _proofIntervalSeconds
    ) external onlyManager {
        // todo we must not interrupt pending proof calculations by rugging the outstanding calculationTimestamps
        proofIntervalSeconds = _proofIntervalSeconds;
    }

    
    function setConfirmer(
        address _confirmer
    ) external onlyManager {
        confirmer = _confirmer;
    }

    /// INTERNAL FUNCTIONS

    function deregister(
        OperatorSet memory operatorSet
    ) external onlyManager onlyRegistered(operatorSet, true) {
        OperatorSet memory substituteOperatorSet = operatorSets[operatorSets.length - 1];
        Snapshots.History storage indexHistory = operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId];
        uint224 operatorSetIndex = indexHistory.latest();
        operatorSets[operatorSetIndex] = substituteOperatorSet;
        operatorSets.pop();

        // update the index of the operator sets
        // note when there is only one operator set left, the index will not be updated as the operator set will be removed in the next step
        operatorSetToIndex[substituteOperatorSet.avs][substituteOperatorSet.operatorSetId].push(
            uint32(block.timestamp), operatorSetIndex
        );
        indexHistory.push(uint32(block.timestamp), REMOVED_INDEX);
    }

    function _getStrategiesAndMultipliers(
        OperatorSet memory operatorSet
    ) internal view returns (IStrategy[] memory, uint256[] memory) {
        IStrategy[] memory strategies =
            new IStrategy[](operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length());
        uint256[] memory multipliers = new uint256[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            (address strategy, uint256 multiplier) =
                operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].at(i);
            strategies[i] = IStrategy(strategy);
            multipliers[i] = multiplier;
        }
        return (strategies, multipliers);
    }

    function _getStakes(
        OperatorSet memory operatorSet,
        IStrategy[] memory strategies,
        uint256[] memory multipliers,
        address operator
    ) internal view returns (uint256, uint256) {
        // calculate the weighted sum of the operator's shares for the strategies given the multipliers
        uint256 delegatedStake = 0;
        uint256 slashableStake = 0;
        {
            uint256[] memory delegatedShares = delegationManager.getOperatorScaledShares(operator, strategies);

            (uint64[] memory totalMagnitudes, uint64[] memory allocatedMagnitudes) =
                allocationManager.getTotalAndAllocatedMagnitudes(operator, operatorSet, strategies);

            for (uint256 i = 0; i < strategies.length; i++) {
                delegatedStake += delegatedShares[i] * totalMagnitudes[i] / 1 ether * multipliers[i];
                slashableStake += delegatedShares[i] * allocatedMagnitudes[i] / 1 ether * multipliers[i];
            }
        }
        return (delegatedStake, slashableStake);
    }

    // VIEW FUNCTIONS
    function isRegistered(
        OperatorSet memory operatorSet
    ) public view returns (bool) {
        (bool exists,, uint224 index) = operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].latestSnapshot();
        return exists && index != REMOVED_INDEX;
    }

    
    function getNumStakeRootSubmissions() external view returns (uint256) {
        return compressedStates.length;
    }

    
    function getNumOperatorSets() external view returns (uint256) {
        return operatorSets.length;
    }

    
    function getStakes(
        OperatorSet calldata operatorSet,
        address operator
    ) external view returns (uint256 delegatedStake, uint256 slashableStake) {
        (IStrategy[] memory strategies, uint256[] memory multipliers) = _getStrategiesAndMultipliers(operatorSet);
        return _getStakes(operatorSet, strategies, multipliers, operator);
    }

    // STAKE ROOT CALCULATION

    
    function getStakeRoot(
        OperatorSet[] calldata operatorSetsInStakeTree,
        bytes32[] calldata operatorSetRoots
    ) external view returns (bytes32) {
        require(
            operatorSets.length == operatorSetsInStakeTree.length,
            "StakeRootCompendium.getStakeRoot: operatorSets vs. operatorSetsInStakeTree length mismatch"
        );
        require(
            operatorSetsInStakeTree.length == operatorSetRoots.length,
            "StakeRootCompendium.getStakeRoot: operatorSetsInStakeTree vs. operatorSetRoots mismatch"
        );
        for (uint256 i = 0; i < operatorSetsInStakeTree.length; i++) {
            require(
                operatorSets[i].avs == operatorSetsInStakeTree[i].avs,
                "StakeRootCompendium.getStakeRoot: operatorSets vs. operatorSetsInStakeTree avs mismatch"
            );
            require(
                operatorSets[i].operatorSetId == operatorSetsInStakeTree[i].operatorSetId,
                "StakeRootCompendium.getStakeRoot: operatorSets vs. operatorSetsInStakeTree operatorSetId mismatch"
            );
        }
        return Merkle.merkleizeKeccak256(operatorSetRoots);
    }

    
    function getOperatorSetLeaves(
        uint256 operatorSetIndex,
        uint256 startOperatorIndex,
        uint256 numOperators
    ) external view returns (OperatorSet memory, address[] memory, OperatorLeaf[] memory) {
        require(
            operatorSetIndex < operatorSets.length,
            "StakeRootCompendium.getOperatorSetLeaves: operator set index out of bounds"
        );
        OperatorSet memory operatorSet = operatorSets[operatorSetIndex];
        address[] memory operators = avsDirectory.getOperatorsInOperatorSet(operatorSet, startOperatorIndex, numOperators);
        (IStrategy[] memory strategies, uint256[] memory multipliers) = _getStrategiesAndMultipliers(operatorSet);

        OperatorLeaf[] memory operatorLeaves = new OperatorLeaf[](operators.length);
        for (uint256 i = 0; i < operatorLeaves.length; i++) {
            // calculate the weighted sum of the operator's shares for the strategies given the multipliers
            (uint256 delegatedStake, uint256 slashableStake) = _getStakes(operatorSet, strategies, multipliers, operators[i]);

            operatorLeaves[i] = OperatorLeaf({
                delegatedStake: delegatedStake,
                slashableStake: slashableStake,
                extraData: operatorExtraDatas[operatorSet.avs][operatorSet.operatorSetId][operators[i]]
            });
        }
        return (operatorSet, operators, operatorLeaves);
    }

    
    function getOperatorSetRoot(
        OperatorSet calldata operatorSet,
        OperatorLeaf[] calldata operatorLeaves
    ) external view returns (bytes32) {
        require(
            avsDirectory.isOperatorSet(operatorSet.avs, operatorSet.operatorSetId),
            "StakeRootCompendium.getOperatorSetRoot: operator set does not exist"
        );
        require(
            operatorLeaves.length == avsDirectory.getNumOperatorsInOperatorSet(operatorSet),
            "AVSSyncTree.getOperatorSetRoot: operator set size mismatch"
        );

        uint256 totalDelegatedStake;
        uint256 totalSlashableStake;
        bytes32 prevExtraData;
        bytes32[] memory operatorLeavesHashes = new bytes32[](operatorLeaves.length);
        for (uint256 i = 0; i < operatorLeaves.length; i++) {
            require(uint256(prevExtraData) < uint256(operatorLeaves[i].extraData), "StakeRootCompendium.getOperatorSetRoot: extraData not sorted");
            prevExtraData = operatorLeaves[i].extraData;

            operatorLeavesHashes[i] = keccak256(
                abi.encodePacked(
                    operatorLeaves[i].delegatedStake,
                    operatorLeaves[i].slashableStake,
                    operatorLeaves[i].extraData
                )
            );

            totalDelegatedStake += operatorLeaves[i].delegatedStake;
            totalSlashableStake += operatorLeaves[i].slashableStake;
        }

        bytes32 operatorTreeRoot = Merkle.merkleizeKeccak256(operatorLeavesHashes);
        return keccak256(
            abi.encodePacked(
                operatorTreeRoot,
                keccak256(
                    abi.encodePacked(
                        totalDelegatedStake, 
                        totalSlashableStake, 
                        operatorSetExtraDatas[operatorSet.avs][operatorSet.operatorSetId]
                    )
                )
            )
        );
    }
}
