// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import "../interfaces/IStrategy.sol";
import "../libraries/Merkle.sol";
import "./StakeRootCompendiumStorage.sol";

contract StakeRootCompendium is StakeRootCompendiumStorage {
    using Snapshots for Snapshots.History;
    using EnumerableMap for EnumerableMap.AddressToUintMap;
    using EnumerableSet for EnumerableSet.Bytes32Set;

    constructor(
        IDelegationManager _delegationManager,
        IAVSDirectory _avsDirectory,
        IAllocationManager _allocationManager,
        uint256 _minBalanceThreshold,
        uint256 _minPrepaidProofs,
        address _verifier,
        bytes32 _imageId
    )
        StakeRootCompendiumStorage(
            _delegationManager,
            _avsDirectory,
            _allocationManager,
            _minBalanceThreshold,
            _minPrepaidProofs,
            _verifier,
            _imageId
        )
    {}

    function initialize(
        address _owner,
        address _rootConfirmer,
        uint32 _proofIntervalSeconds,
        ChargeParams calldata _chargeParams
    ) public initializer {
        __Ownable_init();
        _transferOwnership(_owner);

        rootConfirmer = _rootConfirmer;
        chargeParams = _chargeParams;
        cumulativeChargeParams.proofIntervalSeconds = _proofIntervalSeconds;

        stakeRootSubmissions.push(
            StakeRootSubmission({
                calculationTimestamp: 0,
                submissionTimestamp: uint32(block.timestamp),
                stakeRoot: bytes32(0),
                confirmed: false
            })
        );
    }

    /// OPERATORSET CONFIGURATION

    /// @inheritdoc IStakeRootCompendium
    function deposit(
        OperatorSet calldata operatorSet
    ) external payable {
        DepositInfo storage depositInfo = depositInfos[operatorSet.avs][operatorSet.operatorSetId];

        // If the operator set is not in the stake tree, add it.
        if (!_isInStakeTree(operatorSet)) {
            // Compute updated cumulative charges.
            (, uint256 cumulativeChargePerOperatorSet, uint256 cumulativeChargePerStrategy) =
                _calculateCumulativeCharges();

            // Update the deposit info for the given operator set.
            depositInfo.lastDemandIncreaseTimestamp = uint32(block.timestamp);
            depositInfo.cumulativeChargePerOperatorSetLastPaid = uint96(cumulativeChargePerOperatorSet);
            depositInfo.cumulativeChargePerStrategyLastPaid = uint96(cumulativeChargePerStrategy);

            // Update the total amount of strategies.
            _updateTotalStrategies(
                0, operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length()
            );

            // TODO: What is this variables overall purpose? Maybe explain the why?
            operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].push(
                uint32(block.timestamp), uint224(operatorSets.length)
            );

            // Add operator set to the list of operator sets contained within the stake tree.
            operatorSets.push(operatorSet);
        }

        // Increase the operator set's deposit balance by call value.
        depositInfo.balance += uint96(msg.value);

        // Update the last demand increase timestamp to now.
        depositInfo.lastDemandIncreaseTimestamp = uint32(block.timestamp);

        // Assert that the operator set's deposit balance is greater than the minimum deposit balance.
        require(
            depositInfo.balance
                >= minDepositBalance(
                    operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length()
                ),
            InsufficientDepositBalance()
        );
    }

    /// @inheritdoc IStakeRootCompendium
    function addOrModifyStrategiesAndMultipliers(
        uint32 operatorSetId,
        StrategyAndMultiplier[] calldata _strategiesAndMultipliers
    ) external {
        OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});

        // Update the deposit info for the given operator set.
        _updateDepositInfo(operatorSet);

        EnumerableMap.AddressToUintMap storage strategiesAndMultipliers =
            operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId];

        // Get the number of strategies before the update.
        uint256 numStrategiesBefore = strategiesAndMultipliers.length();

        // Add or modify the strategies and multipliers.
        for (uint256 i = 0; i < _strategiesAndMultipliers.length; ++i) {
            strategiesAndMultipliers.set(
                address(_strategiesAndMultipliers[i].strategy), uint256(_strategiesAndMultipliers[i].multiplier)
            );
        }

        // Get the number of strategies after the update.
        uint256 numStrategiesAfter = strategiesAndMultipliers.length();

        DepositInfo storage depositInfo = depositInfos[operatorSet.avs][operatorSet.operatorSetId];

        // Assert that the operator set's balance is greater than the minimum deposit balance.
        require(depositInfo.balance >= minDepositBalance(numStrategiesAfter), InsufficientDepositBalance());

        // Update the last demand increase timestamp to now.
        depositInfo.lastDemandIncreaseTimestamp = uint32(block.timestamp);

        // Update the total number of strategies.
        _updateTotalStrategies(numStrategiesBefore, numStrategiesAfter);
    }

    /// @inheritdoc IStakeRootCompendium
    function removeStrategiesAndMultipliers(uint32 operatorSetId, IStrategy[] calldata strategies) external {
        OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});

        // Update the deposit info for the given operator set.
        _updateDepositInfo(operatorSet);

        EnumerableMap.AddressToUintMap storage strategiesAndMultipliers =
            operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSetId];

        // Remove the strategies and multipliers.
        for (uint256 i = 0; i < strategies.length; ++i) {
            // Assert the strategy exists, and remove it.
            require(strategiesAndMultipliers.remove(address(strategies[i])), NonexistentStrategy());
        }

        // Get the number of strategies before the update.
        uint256 numStrategiesBefore = strategiesAndMultipliers.length();

        // Update the total number of strategies.
        _updateTotalStrategies(numStrategiesBefore, numStrategiesBefore - strategies.length);
    }

    /// @inheritdoc IStakeRootCompendium
    function setOperatorSetExtraData(uint32 operatorSetId, bytes32 extraData) external {
        // Assert that the operator set is in the stake tree.
        require(_isInStakeTree(OperatorSet({avs: msg.sender, operatorSetId: operatorSetId})), NotInStakeTree());

        // Update the operator set's extra data.
        operatorSetExtraDatas[msg.sender][operatorSetId] = extraData;
    }

    /// @inheritdoc IStakeRootCompendium
    function setOperatorExtraData(uint32 operatorSetId, address operator, bytes32 extraData) external {
        // Assert that the operator set is in the stake tree.
        require(_isInStakeTree(OperatorSet({avs: msg.sender, operatorSetId: operatorSetId})), NotInStakeTree());

        // Update the operator's extra data.
        operatorExtraDatas[msg.sender][operatorSetId][operator] = extraData;
    }

    /// @inheritdoc IStakeRootCompendium
    function withdraw(uint32 operatorSetId, uint256 amount) public payable returns (uint256) {
        OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
        // Assert that the operator set is old enough to withdraw.
        require(canWithdrawDepositBalance(operatorSet), OperatorSetNotOldEnough());
        // Update the deposit info for the given operator set.
        uint256 balance = _updateDepositInfo(operatorSet);
        // If the balance is less than the minimum balance threshold withdraw entire balance.
        if (balance < MIN_BALANCE_THRESHOLD + amount) {
            // Withdraw the entire balance to avoid deposit balance dropping below minimum.
            amount = balance;
            // Remove from the stake tree, if applicable.
            _removeFromStakeTree(operatorSet);
        }
        // Decrease the operator set's balance by the amount withdrawn.
        depositInfos[msg.sender][operatorSetId].balance -= uint96(amount);
        // Transfer the amount to the caller.
        _safeTransferETH(msg.sender, amount);

        return amount;
    }

    /// CHARGE MANAGEMENT

    /// @inheritdoc IStakeRootCompendium
    function removeOperatorSetsFromStakeTree(
        OperatorSet[] calldata operatorSetsToRemove
    ) external {
        // Track the sum of the balances of the operator sets that are removed from the stake tree.
        uint256 penalty;
        // Iterate over the operator sets to remove.
        for (uint256 i = 0; i < operatorSetsToRemove.length; ++i) {
            // If the operator set is in the stake tree, update the deposit info.
            if (_isInStakeTree(operatorSetsToRemove[i])) {
                // Update the deposit info for the given operator set.
                uint256 depositBalance = _updateDepositInfo(operatorSetsToRemove[i]);
                // If the deposit balance is less than the minimum balance threshold, remove from the stake tree.
                if (depositBalance < MIN_BALANCE_THRESHOLD) {
                    _removeFromStakeTree(operatorSetsToRemove[i]);
                    penalty += depositBalance;
                }
            }
        }
        // Note: This reward can be claimed by anyone who submits first.
        // Use a private mempool to avoid failed/hanging transactions.
        _safeTransferETH(msg.sender, penalty);
    }

    /// POSTING ROOTS AND BLACKLISTING

    /// @inheritdoc IStakeRootCompendium
    function verifyStakeRoot(
        uint256 calculationTimestamp,
        bytes32 stakeRoot,
        address chargeRecipient,
        uint256 indexChargePerProof,
        Proof calldata _proof
    ) external {
        // Assert that `calculationTimestamp` is a multiple of the proof interval.
        require(
            calculationTimestamp % cumulativeChargeParams.proofIntervalSeconds == 0,
            TimestampNotMultipleOfProofInterval()
        );
        // Assert that the stake root has not already been posted.
        require(
            stakeRootSubmissions[stakeRootSubmissions.length - 1].calculationTimestamp != calculationTimestamp,
            TimestampAlreadyPosted()
        );

        Snapshots.Snapshot memory _snapshot = chargePerProof._snapshots[indexChargePerProof];

        // TODO: Error name needs modified, and a comment better explaining this check.
        require(
            _snapshot._key <= calculationTimestamp, TimestampOfIndexChargePerProofIsGreaterThanCalculationTimestamp()
        );

        // TODO: Error name needs modified, and a comment better explaining this check.
        require(
            chargePerProof.length() == indexChargePerProof + 1
                || uint256(chargePerProof._snapshots[indexChargePerProof + 1]._key) > calculationTimestamp,
            IndexChargePerProofNotValid()
        );

        // Push the new stake root submission.
        stakeRootSubmissions.push(
            StakeRootSubmission({
                calculationTimestamp: uint32(calculationTimestamp),
                submissionTimestamp: uint32(block.timestamp),
                stakeRoot: stakeRoot,
                confirmed: false
            })
        );

        // TODO: Comment better explaining _snapshot._value's relation to the transfer.
        _safeTransferETH(chargeRecipient, _snapshot._value);

        // TODO: verify proof
        // TODO: prevent race incentives and public mempool sniping, eg embed chargeRecipient in the proof
    }

    /// @inheritdoc IStakeRootCompendium
    function confirmStakeRoot(uint32 index, bytes32 stakeRoot) external {
        // Assert the caller is the root confirmer.
        require(msg.sender == rootConfirmer, OnlyRootConfirmerCanConfirm());
        IStakeRootCompendium.StakeRootSubmission storage stakeRootSubmission = stakeRootSubmissions[index];
        // Assert the stake root matches the stake root submission.
        require(stakeRootSubmission.stakeRoot == stakeRoot, StakeRootDoesNotMatch());
        // Assert the stake root has not already been confirmed.
        require(!stakeRootSubmission.confirmed, TimestampAlreadyConfirmed());
        // Confirm the stake root.
        stakeRootSubmission.confirmed = true;
    }

    /// SET FUNCTIONS

    /// @inheritdoc IStakeRootCompendium
    function setMaxChargePerProof(
        uint96 _maxChargePerProof
    ) external onlyOwner {
        // Assert that the new max charge per proof is greater than the current max charge per proof.
        require(_maxChargePerProof >= chargePerProof.latest(), MaxTotalChargeMustBeGreaterThanTheCurrentTotalCharge());
        // Update the max charge per proof.
        chargeParams.maxChargePerProof = _maxChargePerProof;
    }

    /// @inheritdoc IStakeRootCompendium
    function setChargePerProof(uint96 _chargePerStrategy, uint96 _chargePerOperatorSet) external onlyOwner {
        // Update cumulative charge params.
        _updateCumulativeCharge();
        // Update charge params.
        ChargeParams storage params = chargeParams;
        params.chargePerStrategy = _chargePerStrategy;
        params.chargePerOperatorSet = _chargePerOperatorSet;
        // Update charge per proof.
        _updateChargePerProof();
    }

    /// @inheritdoc IStakeRootCompendium
    function setProofIntervalSeconds(
        uint32 _proofIntervalSeconds
    ) external onlyOwner {
        // Update cumulative charge params.
        _updateCumulativeCharge();
        // Assert that the last calculation timestamp is the same as the last update timestamp.
        // TODO: Better comment explaing the "why' for this check.
        require(
            stakeRootSubmissions[stakeRootSubmissions.length - 1].calculationTimestamp
                == cumulativeChargeParams.lastUpdateTimestamp,
            NoProofsThatHaveBeenChargedButNotSubmitted()
        );
        // Update proof interval seconds.
        cumulativeChargeParams.proofIntervalSeconds = _proofIntervalSeconds;
    }

    /// @inheritdoc IStakeRootCompendium
    function setRootConfirmer(
        address _rootConfirmer
    ) public onlyOwner {
        rootConfirmer = _rootConfirmer;
    }

    /// INTERNAL FUNCTIONS

    function _removeFromStakeTree(
        OperatorSet memory operatorSet
    ) internal {
        OperatorSet memory substituteOperatorSet = operatorSets[operatorSets.length - 1];
        Snapshots.History storage snapshotHistory = operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId];
        uint224 operatorSetIndex = snapshotHistory.latest();
        operatorSets[operatorSetIndex] = substituteOperatorSet;
        operatorSets.pop();

        // update the index of the operator sets
        // note when there is only one operator set left, the index will not be updated as the operator set will be removed in the next step
        operatorSetToIndex[substituteOperatorSet.avs][substituteOperatorSet.operatorSetId].push(
            uint32(block.timestamp), operatorSetIndex
        );
        snapshotHistory.push(uint32(block.timestamp), REMOVED_INDEX);

        depositInfos[operatorSet.avs][operatorSet.operatorSetId].balance = 0;

        _updateTotalStrategies(
            operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length(), 0
        );
    }

    function _updateTotalStrategies(uint256 _countStrategiesBefore, uint256 _countStrategiesAfter) internal {
        totalStrategies = totalStrategies - _countStrategiesBefore + _countStrategiesAfter;
        _updateChargePerProof();
    }

    function _updateChargePerProof() internal {
        // note if totalStrategies is 0, the charge per proof will be 0, and provers should not post a proof
        uint256 _chargePerProof =
            operatorSets.length * chargeParams.chargePerOperatorSet + totalStrategies * chargeParams.chargePerStrategy;
        require(_chargePerProof <= chargeParams.maxChargePerProof, ChargePerProofExceedsMax());
        chargePerProof.push(uint32(block.timestamp), uint224(_chargePerProof));
    }

    function _calculateCumulativeCharges() internal view returns (uint32, uint96, uint96) {
        // calculate the total charge since the last update up until the latest calculation timestamp
        // note that there may be no corresponding stakeRootSubmission for the latest calculation timestamp
        // but if the calculationTimestamp is in the past, then it should be charged for, since proofs are being generated
        uint32 latestCalculationTimestamp =
            uint32(block.timestamp) - uint32(block.timestamp % cumulativeChargeParams.proofIntervalSeconds);

        if (cumulativeChargeParams.lastUpdateTimestamp == latestCalculationTimestamp) {
            return (
                latestCalculationTimestamp,
                cumulativeChargeParams.chargePerOperatorSet,
                cumulativeChargeParams.chargePerStrategy
            );
        }

        uint256 numProofs = (latestCalculationTimestamp - cumulativeChargeParams.lastUpdateTimestamp)
            / cumulativeChargeParams.proofIntervalSeconds;

        return (
            latestCalculationTimestamp,
            uint96(cumulativeChargeParams.chargePerOperatorSet + chargeParams.chargePerOperatorSet * numProofs),
            uint96(cumulativeChargeParams.chargePerStrategy + chargeParams.chargePerStrategy * numProofs)
        );
    }

    function _updateCumulativeCharge() internal {
        (uint32 lastUpdateTimestamp, uint96 cumulativeChargePerOperatorSet, uint96 cumulativeChargePerStrategy) =
            _calculateCumulativeCharges();

        cumulativeChargeParams = CumulativeChargeParams({
            chargePerOperatorSet: cumulativeChargePerOperatorSet,
            chargePerStrategy: cumulativeChargePerStrategy,
            lastUpdateTimestamp: lastUpdateTimestamp,
            proofIntervalSeconds: cumulativeChargeParams.proofIntervalSeconds
        });
    }

    // updates the deposit balance for the operator set and returns the penalty if the operator set has fallen below the minimum deposit balance
    function _updateDepositInfo(
        OperatorSet memory operatorSet
    ) internal returns (uint256 balance) {
        require(_isInStakeTree(operatorSet), NotInStakeTree());

        (, uint256 cumulativeChargePerOperatorSet, uint256 cumulativeChargePerStrategy) = _calculateCumulativeCharges();
        DepositInfo storage depositInfo = depositInfos[operatorSet.avs][operatorSet.operatorSetId];

        // subtract new total charge from last paid total charge
        uint256 pendingCharge = cumulativeChargePerOperatorSet - depositInfo.cumulativeChargePerOperatorSetLastPaid
            + (
                (cumulativeChargePerStrategy - depositInfo.cumulativeChargePerStrategyLastPaid)
                    * operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length()
            );

        balance = depositInfo.balance > pendingCharge ? depositInfo.balance - pendingCharge : 0;

        depositInfo.balance = uint96(balance);
        depositInfo.cumulativeChargePerOperatorSetLastPaid = uint96(cumulativeChargePerOperatorSet);
        depositInfo.cumulativeChargePerOperatorSetLastPaid = uint96(cumulativeChargePerStrategy);
    }

    function _getStrategiesAndMultipliers(
        OperatorSet memory operatorSet
    ) internal view returns (IStrategy[] memory, uint256[] memory) {
        IStrategy[] memory strategies =
            new IStrategy[](operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length());
        uint256[] memory multipliers = new uint256[](strategies.length);
        for (uint256 i = 0; i < strategies.length; ++i) {
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
            uint256[] memory delegatedShares = delegationManager.getOperatorDelegatedShares(operator, strategies);

            (uint64[] memory totalMagnitudes, uint64[] memory allocatedMagnitudes) =
                allocationManager.getTotalAndAllocatedMagnitudes(operator, operatorSet, strategies);

            for (uint256 i = 0; i < strategies.length; ++i) {
                delegatedStake += delegatedShares[i] * totalMagnitudes[i] / 1 ether * multipliers[i];
                slashableStake += delegatedShares[i] * allocatedMagnitudes[i] / 1 ether * multipliers[i];
            }
        }
        return (delegatedStake, slashableStake);
    }

    function _isInStakeTree(
        OperatorSet memory operatorSet
    ) internal view returns (bool) {
        (bool exists,, uint224 index) = operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].latestSnapshot();
        return exists && index != REMOVED_INDEX;
    }

    // VIEW FUNCTIONS

    /// VIEW FUNCTIONS

    /// @inheritdoc IStakeRootCompendium
    function minDepositBalance(
        uint256 numStrategies
    ) public view returns (uint256) {
        ChargeParams memory charges = chargeParams;
        return (numStrategies * charges.chargePerStrategy + charges.chargePerOperatorSet) * MIN_PREPAID_PROOFS;
    }

    /// @inheritdoc IStakeRootCompendium
    function canWithdrawDepositBalance(
        OperatorSet memory operatorSet
    ) public view returns (bool) {
        // they must have paid for all of their prepaid proofs before withdrawing after a demand increase
        return block.timestamp
            > depositInfos[operatorSet.avs][operatorSet.operatorSetId].lastDemandIncreaseTimestamp
                + MIN_PREPAID_PROOFS * cumulativeChargeParams.proofIntervalSeconds;
    }

    /// @inheritdoc IStakeRootCompendium
    function getNumStakeRootSubmissions() external view returns (uint256) {
        return stakeRootSubmissions.length;
    }

    /// @inheritdoc IStakeRootCompendium
    function getStakeRootSubmission(
        uint32 index
    ) external view returns (StakeRootSubmission memory) {
        return stakeRootSubmissions[index];
    }

    /// @inheritdoc IStakeRootCompendium
    function getNumOperatorSets() external view returns (uint256) {
        return operatorSets.length;
    }

    /// @inheritdoc IStakeRootCompendium
    function getOperatorSetIndexAtTimestamp(OperatorSet memory operatorSet, uint32 timestamp)
        external
        view
        returns (uint224)
    {
        return operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].upperLookup(timestamp);
    }

    /// @inheritdoc IStakeRootCompendium
    function getStakes(
        OperatorSet calldata operatorSet,
        address operator
    ) external view returns (uint256 delegatedStake, uint256 slashableStake) {
        (IStrategy[] memory strategies, uint256[] memory multipliers) = _getStrategiesAndMultipliers(operatorSet);
        return _getStakes(operatorSet, strategies, multipliers, operator);
    }

    /// @inheritdoc IStakeRootCompendium
    function getDepositBalance(
        OperatorSet memory operatorSet
    ) external view returns (uint256 balance) {
        DepositInfo memory depositInfo = depositInfos[operatorSet.avs][operatorSet.operatorSetId];
        (, uint96 cumulativeChargePerOperatorSet, uint96 cumulativeChargePerStrategy) = _calculateCumulativeCharges();
        uint256 pendingCharge = uint256(
            cumulativeChargePerOperatorSet - depositInfo.cumulativeChargePerOperatorSetLastPaid
        )
            + uint256(cumulativeChargePerStrategy - depositInfo.cumulativeChargePerStrategyLastPaid)
                * operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length();

        return depositInfo.balance > pendingCharge ? depositInfo.balance - pendingCharge : 0;
    }

    /// @inheritdoc IStakeRootCompendium
    function proofIntervalSeconds() external view returns (uint32) {
        return cumulativeChargeParams.proofIntervalSeconds;
    }

    // STAKE ROOT CALCULATION

    /// @inheritdoc IStakeRootCompendium
    function getStakeRoot(
        OperatorSet[] calldata operatorSetsInStakeTree,
        bytes32[] calldata operatorSetRoots
    ) external view returns (bytes32) {
        require(operatorSets.length == operatorSetsInStakeTree.length, OperatorSetSizeMismatch());
        require(operatorSetsInStakeTree.length == operatorSetRoots.length, InputArrayLengthMismatch());
        for (uint256 i = 0; i < operatorSetsInStakeTree.length; ++i) {
            require(operatorSets[i].avs == operatorSetsInStakeTree[i].avs, InputCorrelatedVariableMismatch());
            require(
                operatorSets[i].operatorSetId == operatorSetsInStakeTree[i].operatorSetId,
                InputCorrelatedVariableMismatch()
            );
        }
        return Merkle.merkleizeKeccak256(operatorSetRoots);
    }

    /// @inheritdoc IStakeRootCompendium
    function getOperatorSetLeaves(
        uint256 operatorSetIndex,
        uint256 startOperatorIndex,
        uint256 numOperators
    ) external view returns (OperatorSet memory, address[] memory, OperatorLeaf[] memory) {
        require(operatorSetIndex < operatorSets.length, OperatorSetIndexOutOfBounds());
        OperatorSet memory operatorSet = operatorSets[operatorSetIndex];
        address[] memory operators =
            avsDirectory.getOperatorsInOperatorSet(operatorSet, startOperatorIndex, numOperators);
        (IStrategy[] memory strategies, uint256[] memory multipliers) = _getStrategiesAndMultipliers(operatorSet);

        OperatorLeaf[] memory operatorLeaves = new OperatorLeaf[](operators.length);
        for (uint256 i = 0; i < operatorLeaves.length; ++i) {
            // calculate the weighted sum of the operator's shares for the strategies given the multipliers
            (uint256 delegatedStake, uint256 slashableStake) =
                _getStakes(operatorSet, strategies, multipliers, operators[i]);

            operatorLeaves[i] = OperatorLeaf({
                delegatedStake: delegatedStake,
                slashableStake: slashableStake,
                extraData: operatorExtraDatas[operatorSet.avs][operatorSet.operatorSetId][operators[i]]
            });
        }
        return operatorLeaves;
    }

    /// @inheritdoc IStakeRootCompendium
    function getOperatorSetRoot(
        uint32 operatorSetIndex,
        OperatorLeaf[] calldata operatorLeaves
    ) external view returns (bytes32) {
        OperatorSet memory operatorSet = operatorSets[operatorSetIndex];
        require(avsDirectory.isOperatorSet(operatorSet.avs, operatorSet.operatorSetId), OperatorSetMustExist());
        require(
            operatorLeaves.length == avsDirectory.getNumOperatorsInOperatorSet(operatorSet), InputArrayLengthMismatch()
        );
        uint256 totalDelegatedStake;
        for (uint256 i = 0; i < operatorLeaves.length; ++i) {
            require(uint256(prevExtraData) < uint256(operatorLeaves[i].extraData), ExtraDataNotSorted());
            prevExtraData = operatorLeaves[i].extraData;

            operatorLeavesHashes[i] = keccak256(
                abi.encodePacked(
                    operatorLeaves[i].delegatedStake, operatorLeaves[i].slashableStake, operatorLeaves[i].extraData
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

    function _safeTransferETH(address to, uint256 amount) internal {
        (bool success,) = to.call{value: amount}("");
        require(success, EthTransferFailed());
    }

    // in case of charge problems
    receive() external payable {}
}
