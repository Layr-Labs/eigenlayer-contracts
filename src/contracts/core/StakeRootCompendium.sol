// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./StakeRootCompendiumStorage.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import "../libraries/Checkpoints.sol";

contract StakeRootCompendium is StakeRootCompendiumStorage {
    using Checkpoints for Checkpoints.History;
    using EnumerableMap for EnumerableMap.AddressToUintMap;
    using EnumerableSet for EnumerableSet.Bytes32Set;

    constructor(
        IDelegationManager _delegationManager,
        IAVSDirectory _avsDirectory,
        uint32 _proofInterval,
        uint32 _blacklistWindow
    ) StakeRootCompendiumStorage(_delegationManager, _avsDirectory, _proofInterval, _blacklistWindow) {}

    function initialize(address initialOwner) public initializer {
        __Ownable_init();
        _transferOwnership(initialOwner);
    }

    /// OPERATORSET CONFIGURATION

    /// @inheritdoc IStakeRootCompendium
    function depositForOperatorSet(IAVSDirectory.OperatorSet calldata operatorSet) external payable {
        require(
            avsDirectory.isOperatorSet(operatorSet.avs, operatorSet.operatorSetId),
            "StakeRootCompendium.depositForOperatorSet: operator set does not exist"
        );
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
        require(
            strategiesAndMultipliers.length > 0,
            "StakeRootCompendium.setStrategiesAndMultipliers: no strategies and multipliers provided"
        );
        require(
            avsDirectory.isOperatorSet(msg.sender, operatorSetId),
            "StakeRootCompendium.setStrategiesAndMultipliers: operator set does not exist"
        );

        IAVSDirectory.OperatorSet memory operatorSet =
            IAVSDirectory.OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});

        // update the deposit balance for the operator set whenever number of strategies is changed
        _updateDepositBalanceInfo(operatorSet, true);

        uint256 numStrategiesBefore = operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].length();
        // if the operator set has been configured to have a positive number of strategies, increment the number of configured operator sets
        if (numStrategiesBefore == 0) {
            require(
                operatorSets.length < MAX_NUM_OPERATOR_SETS,
                "StakeRootCompendium.setStrategiesAndMultipliers: too many operator sets"
            );
            operatorSetToIndex[msg.sender][operatorSetId].push(uint32(block.timestamp), uint208(operatorSets.length));
            operatorSets.push(operatorSet);
        }

        // set the strategies and multipliers for the operator set
        for (uint256 i = 0; i < strategiesAndMultipliers.length; i++) {
            operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].set(
                address(strategiesAndMultipliers[i].strategy), uint256(strategiesAndMultipliers[i].multiplier)
            );
        }
        require(
            operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].length() <= MAX_NUM_STRATEGIES, 
            "StakeRootCompendium.setStrategiesAndMultipliers: too many strategies"
        );

        _updateTotals(numStrategiesBefore, operatorSetToStrategyAndMultipliers[msg.sender][operatorSetId].length());
    }

    /// @inheritdoc IStakeRootCompendium
    function removeStrategiesAndMultipliers(uint32 operatorSetId, IStrategy[] calldata strategies) external {
        IAVSDirectory.OperatorSet memory operatorSet =
            IAVSDirectory.OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
        // update the deposit balance for the operator set whenever number of strategies is changed
        _updateDepositBalanceInfo(operatorSet, true);

        // remove the strategies and multipliers for the operator set
        _removeStrategiesAndMultipliers(operatorSet, strategies);
    }

    /// @inheritdoc IStakeRootCompendium
    function setOperatorSetExtraData(uint32 operatorSetId, bytes32 extraData) external {
        (bool exists,, uint224 index) = operatorSetToIndex[msg.sender][operatorSetId].latestCheckpoint();
        require(exists && index != REMOVED_INDEX, "StakeRootCompendium.setOperatorSetExtraData: operatorSet is not in stakeTree");
        operatorSetExtraDatas[msg.sender][operatorSetId] = extraData;
    }

    /// @inheritdoc IStakeRootCompendium
    function setOperatorExtraData(uint32 operatorSetId, address operator, bytes32 extraData) external {
        (bool exists,, uint224 index) = operatorSetToIndex[msg.sender][operatorSetId].latestCheckpoint();
        require(exists && index != REMOVED_INDEX, "StakeRootCompendium.setOperatorExtraData: operatorSet is not in stakeTree");
        operatorExtraDatas[msg.sender][operatorSetId][operator] = extraData;
    }

    /// POSTING ROOTS AND BLACKLISTING

    /// @inheritdoc IStakeRootCompendium
    function verifyStakeRoot(
        uint32 calculationTimestamp,
        bytes32 stakeRoot,
        address chargeRecipient,
        Proof calldata proof
    ) external {
        // TODO: verify proof

        _postStakeRoot(calculationTimestamp, stakeRoot, chargeRecipient, false);
    }

    /// @inheritdoc IStakeRootCompendium
    function blacklistStakeRoot(uint32 submissionIndex) external onlyOwner {
        // TODO: this should not be onlyOwner
        require(
            !stakeRootSubmissions[submissionIndex].blacklisted,
            "StakeRootCompendium.blacklistStakeRoot: stakeRoot already blacklisted"
        );
        require(
            block.timestamp < stakeRootSubmissions[submissionIndex].blacklistableBefore,
            "StakeRootCompendium.blacklistStakeRoot: stakeRoot cannot be blacklisted"
        );
        require(
            !stakeRootSubmissions[submissionIndex].forcePosted,
            "StakeRootCompendium.blacklistStakeRoot: stakeRoot was force posted"
        );
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
        for (uint256 i = 0; i < operatorSetsToUpdate.length; i++) {
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
                block.timestamp < stakeRootSubmission.blacklistableBefore || stakeRootSubmission.blacklisted
                    || stakeRootSubmission.forcePosted
            ) {
                continue;
            }
            // if the charge recipient has changed, transfer the total charge to the previous recipient
            if (stakeRootSubmission.chargeRecipient != prevRecipient) {
                if (totalCharge > 0) {
                    payable(prevRecipient).transfer(totalCharge);
                }
                totalCharge = 0;
                prevRecipient = stakeRootSubmission.chargeRecipient;
            }
            // total charge is the charge per strategy per proof times the number of strategies at the time of proof
            totalCharge += totalChargeSnapshot.upperLookup(stakeRootSubmissions[i].calculationTimestamp);
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

    function setChargeForNumStrategies(uint64 newConstantChargePerProof, uint64 newLinearChargePerProof) external onlyOwner {
        _setChargePerProof(newConstantChargePerProof, newLinearChargePerProof);
    }

    /// VIEW FUNCTIONS

    /// @inheritdoc IStakeRootCompendium
    function getNumOperatorSets() external view returns (uint256) {
        return operatorSets.length;
    }

    /// @inheritdoc IStakeRootCompendium
    function getStakes(IAVSDirectory.OperatorSet calldata operatorSet, address operator)
        external
        view
        returns (uint256 delegatedStake, uint256 slashableStake)
    {
        (IStrategy[] memory strategies, uint256[] memory multipliers) = _getStrategiesAndMultipliers(operatorSet);
        return _getStakes(operatorSet, strategies, multipliers, operator);
    }

    /// @inheritdoc IStakeRootCompendium
    function getStakeRootSubmission(uint32 index) external view returns (StakeRootSubmission memory) {
        return stakeRootSubmissions[index];
    }

    /// @inheritdoc IStakeRootCompendium
    function getNumStakeRootSubmissions() external view returns (uint256) {
        return stakeRootSubmissions.length;
    }

    /// @inheritdoc IStakeRootCompendium
    function getOperatorSetIndexAtTimestamp(
        IAVSDirectory.OperatorSet calldata operatorSet,
        uint32 timestamp
    ) external view returns (uint32) {
        return uint32(operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].upperLookupRecent(timestamp));
    }

    /// @inheritdoc IStakeRootCompendium
    function getDepositBalance(IAVSDirectory.OperatorSet memory operatorSet)
        public
        view
        returns (uint256 balance, uint256 penalty)
    {
        (uint224 previousConstantCumulativeCharge, uint224 previousLinearCumulativeCharge) = _getPreviousCumulativeCharges(
            depositBalanceInfo[operatorSet.avs][operatorSet.operatorSetId].latestUpdateTime
        );
        (uint224 currentConstantCumulativeCharge, uint224 currentLinearCumulativeCharge) = _getCurrentCumulativeCharges(_getLatestCalculationTimestamp());

        uint256 pendingCharge =
            uint256(currentConstantCumulativeCharge - previousConstantCumulativeCharge) +
            uint256(currentLinearCumulativeCharge - previousLinearCumulativeCharge) * 
            operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length();
        uint256 storedBalance = depositBalanceInfo[operatorSet.avs][operatorSet.operatorSetId].balance;

        // if the charge would take the balance below the minimum deposit balance, return 0
        balance = storedBalance < pendingCharge + MIN_DEPOSIT_BALANCE ? 0 : storedBalance - pendingCharge;
        // if the balance is 0, then the charger may be able to claim the penalty for removing the operatorSet from the stakeTree
        penalty = balance == 0 ? pendingCharge + MIN_DEPOSIT_BALANCE - storedBalance : 0;
        return (balance, penalty);
    }

    /// Misc

    receive() external payable {}

    /// INTERNAL FUNCTIONS

    function _setChargePerProof(uint64 newConstantChargePerProof, uint64 newLinearChargePerProof) internal {
        _updateCumulativeCharges();
        constantChargePerProof = newConstantChargePerProof;
        linearChargePerProof = newLinearChargePerProof;
        _updateTotals(0, 0);
    }

    function _setProofInterval(uint32 proofInterval) internal {
        require(
            stakeRootSubmissions.length == 0 || _getLatestSubmittedCalculationTimestamp() == _getLatestCalculationTimestamp(), 
            "StakeRootCompendium._setProofInterval: make sure there are no pending proofs"
        );
        _updateCumulativeCharges();
        proofInterval = proofInterval;
    }

    function _updateCumulativeCharges() internal {
        uint32 latestCalculationTimestamp = _getLatestCalculationTimestamp();
        (uint224 currentConstantCumulativeCharge, uint224 currentLinearCumulativeCharge) = _getCurrentCumulativeCharges(latestCalculationTimestamp);

        // update the cumulative charge snapshots
        cumulativeContantChargeSnapshot.push(
            latestCalculationTimestamp,
            currentConstantCumulativeCharge
        );
        cumulativeLinearChargeSnapshot.push(
            latestCalculationTimestamp,
            currentLinearCumulativeCharge
        );
    }

    function _getCurrentCumulativeCharges(uint32 latestCalculationTimestamp) internal view returns(uint224, uint224) {
        return (
            _getCurrentCumulativeCharge(cumulativeContantChargeSnapshot, constantChargePerProof, latestCalculationTimestamp),
            _getCurrentCumulativeCharge(cumulativeLinearChargeSnapshot, linearChargePerProof, latestCalculationTimestamp)
        );
    }

    function _getPreviousCumulativeCharges(uint32 timestamp) internal view returns(uint224, uint224) {
        return (
            cumulativeContantChargeSnapshot.upperLookup(timestamp),
            cumulativeLinearChargeSnapshot.upperLookup(timestamp)
        );
    }

    function _getCurrentCumulativeCharge(Checkpoints.History storage cumulativeChargeSnapshot, uint64 currentCharge, uint32 latestCalculationTimestamp) internal view returns (uint224) {
        (bool exists, uint32 latestSnapshotTimestamp, uint224 latestSnapshottedCumulativeCharge) = cumulativeChargeSnapshot.latestCheckpoint();
        // return the latest snapshot if it has been accounted for
        if(latestSnapshotTimestamp == latestCalculationTimestamp) {
            return latestSnapshottedCumulativeCharge;
        }

        // use "now" if this is the first time the charge is set
        if (!exists) {
            latestSnapshotTimestamp = latestCalculationTimestamp;
            latestSnapshottedCumulativeCharge = 0;
        }

        // keep track of the cumulative charge that would have been charged since the last change
        // this is: 
        // latestSnapshottedCumulativeCharge + proofs since latest snapshot * (charge / proof)
        // latestSnapshottedCumulativeCharge + (latestCalculationTimestamp - latestSnapshotTimestamp) * (1 proof / proofInterval time) * (charge / proof)
        // latestSnapshottedCumulativeCharge + (charge / proof) * (time since snapshot) * (1 proof / proofInterval time)
        return latestSnapshottedCumulativeCharge + uint224(currentCharge * (latestCalculationTimestamp - latestSnapshotTimestamp) / proofInterval);
    }

    function _removeStrategiesAndMultipliers(IAVSDirectory.OperatorSet memory operatorSet, IStrategy[] memory strategies) internal {
        uint256 numStrategiesBefore = operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length();

        for (uint256 i = 0; i < strategies.length; i++) {
            require(
                operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].remove(address(strategies[i])),
                "StakeRootCompendium.removeStrategiesAndMultipliers: strategy not found"
            );
        }

        _updateTotals(numStrategiesBefore, operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length());
    }

    function _removeOperatorSet(IAVSDirectory.OperatorSet memory operatorSet) internal {
        if(operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length() != 0) {
            bytes32[] memory strategyBytes = operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId]._inner._keys.values();
            IStrategy[] memory strategies = new IStrategy[](strategyBytes.length);
            // todo: better way to do this?
            for (uint256 i = 0; i < strategyBytes.length; i++) {
                strategies[i] = IStrategy(address(uint160(uint256(strategyBytes[i]))));
            }

            // remove the strategies and multipliers for the operator set
            _removeStrategiesAndMultipliers(
                operatorSet,
                strategies
            );
        }

        IAVSDirectory.OperatorSet memory substituteOperatorSet = operatorSets[operatorSets.length - 1];
        uint224 operatorSetIndex = operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].latest();
        operatorSets[operatorSetIndex] = substituteOperatorSet;
        operatorSets.pop();

        // update the index of the operator sets
        operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].push(uint32(block.timestamp), REMOVED_INDEX);
        operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].push(uint32(block.timestamp), operatorSetIndex);
    }

    // updates the deposit balance for the operator set and returns the penalty if the operator set has fallen below the minimum deposit balance
    function _updateDepositBalanceInfo(
        IAVSDirectory.OperatorSet memory operatorSet,
        bool sendPenalty
    ) internal returns (uint256) {
        _updateCumulativeCharges();
        (uint256 depositBalance, uint256 penalty) = getDepositBalance(operatorSet);
        depositBalanceInfo[operatorSet.avs][operatorSet.operatorSetId].balance = depositBalance;
        depositBalanceInfo[operatorSet.avs][operatorSet.operatorSetId].latestUpdateTime = uint32(block.timestamp);
        // if the operatorSet has fallen below the minimum deposit balance, remove it from the stakeTree
        if (penalty > 0) {
            _removeOperatorSet(operatorSet);
            if (sendPenalty) {
                payable(msg.sender).transfer(penalty);
            }
        }

        return penalty;
    }

    // updates total strategies and the total charge per proof whenever an operator set's number of strategies changes
    function _updateTotals(uint256 numStrategiesBefore, uint256 numStrategiesAfter) internal {
        totalStrategies = totalStrategies - numStrategiesBefore + numStrategiesAfter;
        totalChargeSnapshot.push(
            uint32(block.timestamp), 
            uint224(totalStrategies * linearChargePerProof + constantChargePerProof)
        );
    }

    function _postStakeRoot(
        uint32 calculationTimestamp,
        bytes32 stakeRoot,
        address chargeRecipient,
        bool forcePosted
    ) internal {
        require(
            calculationTimestamp % proofInterval == 0,
            "StakeRootCompendium._postStakeRoot: calculationTimestamp must be a multiple of proofInterval"
        );

        uint256 stakeRootSubmissionsLength = stakeRootSubmissions.length;
        if (stakeRootSubmissionsLength != 0) {
            require(
                stakeRootSubmissions[stakeRootSubmissionsLength - 1].calculationTimestamp + proofInterval == calculationTimestamp,
                "StakeRootCompendium._postStakeRoot: calculationTimestamp must be greater than the last posted calculationTimestamp"
            );
        }

        stakeRootSubmissions.push(
            StakeRootSubmission({
                stakeRoot: stakeRoot,
                chargeRecipient: msg.sender,
                calculationTimestamp: calculationTimestamp,
                blacklistableBefore: uint32(block.timestamp) + blacklistWindow,
                blacklisted: false,
                crossPosted: false,
                forcePosted: forcePosted
            })
        );

        // todo: emit events
    }

    /// @notice gets the latest calculation timestamp, whether a stakeRoot has been posted or not
    function _getLatestCalculationTimestamp() internal view returns (uint32) {
        uint256 stakeRootSubmissionsLength = stakeRootSubmissions.length;
        require(stakeRootSubmissionsLength > 0, "StakeRootCompendium._getLatestCalculationTimestamp: first empty stakeRoot must be posted");
        uint32 latestCalculationTimestamp = stakeRootSubmissions[stakeRootSubmissionsLength - 1].calculationTimestamp;
        if (latestCalculationTimestamp + proofInterval < block.timestamp) {
            latestCalculationTimestamp += proofInterval;
        }
        return latestCalculationTimestamp;
    }

    function _getLatestSubmittedCalculationTimestamp() internal view returns (uint32) {
        uint256 stakeRootSubmissionsLength = stakeRootSubmissions.length;
        require(stakeRootSubmissionsLength > 0, "StakeRootCompendium._getLatestCalculationTimestamp: first empty stakeRoot must be posted");
        return stakeRootSubmissions[stakeRootSubmissionsLength - 1].calculationTimestamp;
    }

    function _getStrategiesAndMultipliers(IAVSDirectory.OperatorSet memory operatorSet)
        internal
        view
        returns (IStrategy[] memory, uint256[] memory)
    {
        IStrategy[] memory strategies = new IStrategy[](
            operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length()
        );
        uint256[] memory multipliers = new uint256[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            (address strategy, uint256 multiplier) =
                operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].at(i);
            strategies[i] = IStrategy(strategy);
            multipliers[i] = multiplier;
        }
        return (strategies, multipliers);
    }

    function _getStakes(IAVSDirectory.OperatorSet memory operatorSet, IStrategy[] memory strategies, uint256[] memory multipliers, address operator) internal view returns (uint256, uint256) {
        // calculate the weighted sum of the operator's shares for the strategies given the multipliers
        uint256 delegatedStake = 0;
        uint256 slashableStake = 0;
        {
            uint256[] memory delegatedShares = delegationManager.getOperatorScaledShares(operator, strategies);

            (uint64[] memory totalMagnitudes, uint64[] memory allocatedMagnitudes) = 
                avsDirectory.getTotalAndAllocatedMagnitudes(operator, operatorSet, strategies);

            for (uint256 i = 0; i < strategies.length; i++) {
                uint256 multipliedDelegatedShares = delegatedShares[i] * multipliers[i] / 1 ether;
                delegatedStake += multipliedDelegatedShares * totalMagnitudes[i];
                slashableStake += multipliedDelegatedShares * allocatedMagnitudes[i];
            }
        }
        return (delegatedStake, slashableStake);
    }

    // STAKE ROOT CALCULATION

    /// @inheritdoc IStakeRootCompendium
    function getStakeRoot(
        IAVSDirectory.OperatorSet[] calldata operatorSetsInStakeTree,
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

    /// @inheritdoc IStakeRootCompendium
    function getOperatorSetLeaves(
        uint256 operatorSetIndex,
        uint256 startOperatorIndex,
        uint256 numOperators
    ) external view returns (IAVSDirectory.OperatorSet memory, address[] memory, OperatorLeaf[] memory) {
        require(
            operatorSetIndex < operatorSets.length,
            "StakeRootCompendium.getOperatorSetLeaves: operator set index out of bounds"
        );
        IAVSDirectory.OperatorSet memory operatorSet = operatorSets[operatorSetIndex];
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

    /// @inheritdoc IStakeRootCompendium
    function getOperatorSetRoot(
        IAVSDirectory.OperatorSet calldata operatorSet,
        address[] calldata operators,
        OperatorLeaf[] calldata operatorLeaves
    ) external view returns (bytes32) {
        require(
            avsDirectory.isOperatorSet(operatorSet.avs, operatorSet.operatorSetId),
            "StakeRootCompendium.getOperatorSetRoot: operator set does not exist"
        );
        require(operatorLeaves.length <= MAX_OPERATOR_SET_SIZE, "AVSSyncTree._verifyOperatorStatus: operator set too large");
        require(
            operatorLeaves.length == avsDirectory.getNumOperatorsInOperatorSet(operatorSet),
            "AVSSyncTree.getOperatorSetRoot: operator set size mismatch"
        );

        uint256 totalDelegatedStake;
        uint256 totalSlashableStake;
        address prevOperator;
        bytes32[] memory operatorLeavesHashes = new bytes32[](operatorLeaves.length);
        for (uint256 i = 0; i < operatorLeaves.length; i++) {
            require(uint160(prevOperator) < uint160(operators[i]), "AVSSyncTree.getOperatorSetRoot: operators not sorted");
            prevOperator = operators[i];

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
