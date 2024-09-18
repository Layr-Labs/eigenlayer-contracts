// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./StakeRootCompendiumStorage.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import "../libraries/Snapshots.sol";

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
        uint96 _maxTotalCharge,
        uint96 _chargePerStrategy,
        uint96 _chargePerOperatorSet
    ) public initializer {
        __Ownable_init();
        _transferOwnership(_owner);

        rootConfirmer = _rootConfirmer;

        maxTotalCharge = _maxTotalCharge;
        proofIntervalSeconds = _proofIntervalSeconds;
        chargePerOperatorSet = _chargePerOperatorSet;
        chargePerStrategy = _chargePerStrategy;

        stakeRootSubmissions.push(
            StakeRootSubmission({calculationTimestamp: 0, stakeRoot: bytes32(0), confirmed: false})
        );
    }

    /// OPERATORSET CONFIGURATION

    /// @inheritdoc IStakeRootCompendium
    function deposit(
        OperatorSet calldata operatorSet
    ) external payable {
        if (!_isInStakeTree(operatorSet)) {
            (, uint256 cumulativeChargePerOperatorSet, uint256 cumulativeChargePerStrategy) = _cumulativeCharges();
            depositInfos[operatorSet.avs][operatorSet.operatorSetId] = DepositInfo({
                balance: 0, // balance will be updated outer context
                lastDemandIncreaseTimestamp: uint32(block.timestamp),
                cumulativeChargePerOperatorSetLastPaid: uint96(cumulativeChargePerOperatorSet),
                cumulativeChargePerStrategyLastPaid: uint96(cumulativeChargePerStrategy)
            });

            _updateTotalStrategies(
                0, operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length()
            );
            operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].push(
                uint32(block.timestamp), uint224(operatorSets.length)
            );
            operatorSets.push(operatorSet);
        }
        depositInfos[operatorSet.avs][operatorSet.operatorSetId].balance += uint96(msg.value);
        // note that they've shown increased demand for proofs
        depositInfos[operatorSet.avs][operatorSet.operatorSetId].lastDemandIncreaseTimestamp = uint32(block.timestamp);
        // make sure they have enough to pay for MIN_PROOFS_PREPAID
        require(
            depositInfos[operatorSet.avs][operatorSet.operatorSetId].balance
                >= minDepositBalance(
                    operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length()
                ),
            "StakeRootCompendium.addOrModifyStrategiesAndMultipliers: insufficient deposit balance"
        );
    }

    /// @inheritdoc IStakeRootCompendium
    function addOrModifyStrategiesAndMultipliers(
        uint32 operatorSetId,
        StrategyAndMultiplier[] calldata strategiesAndMultipliers
    ) external {
        OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
        // update the deposit balance for the operator set whenever number of strategies is changed
        _updateDepositInfo(operatorSet);

        uint256 numStrategiesBefore =
            operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length();
        // set the strategies and multipliers for the operator set
        for (uint256 i = 0; i < strategiesAndMultipliers.length; i++) {
            operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].set(
                address(strategiesAndMultipliers[i].strategy), uint256(strategiesAndMultipliers[i].multiplier)
            );
        }
        uint256 numStrategiesAfter =
            operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length();

        // make sure they have enough to pay for MIN_PREPAID_PROOFS
        require(
            depositInfos[operatorSet.avs][operatorSet.operatorSetId].balance >= minDepositBalance(numStrategiesAfter),
            "StakeRootCompendium.addOrModifyStrategiesAndMultipliers: insufficient deposit balance"
        );
        // note that they've shown increased demand for proofs
        depositInfos[operatorSet.avs][operatorSet.operatorSetId].lastDemandIncreaseTimestamp = uint32(block.timestamp);

        // only adding new strategies to count
        _updateTotalStrategies(numStrategiesBefore, numStrategiesAfter);
    }

    /// @inheritdoc IStakeRootCompendium
    function removeStrategiesAndMultipliers(uint32 operatorSetId, IStrategy[] calldata strategies) external {
        OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
        // update the deposit balance for the operator set whenever number of strategies is changed
        _updateDepositInfo(operatorSet);

        // note below either all strategies are removed or none are removed and transaction reverts
        uint256 numStrategiesBefore = operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSetId].length();
        for (uint256 i = 0; i < strategies.length; i++) {
            require(
                operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSetId].remove(address(strategies[i])),
                "StakeRootCompendium.removeStrategiesAndMultipliers: strategy not found"
            );
        }

        _updateTotalStrategies(numStrategiesBefore, numStrategiesBefore - strategies.length);
    }

    /// @inheritdoc IStakeRootCompendium
    function setOperatorSetExtraData(uint32 operatorSetId, bytes32 extraData) external {
        require(
            _isInStakeTree(OperatorSet({avs: msg.sender, operatorSetId: operatorSetId})),
            "StakeRootCompendium.setOperatorSetExtraData: operatorSet is not in stakeTree"
        );
        operatorSetExtraDatas[msg.sender][operatorSetId] = extraData;
    }

    /// @inheritdoc IStakeRootCompendium
    function setOperatorExtraData(uint32 operatorSetId, address operator, bytes32 extraData) external {
        require(
            _isInStakeTree(OperatorSet({avs: msg.sender, operatorSetId: operatorSetId})),
            "StakeRootCompendium.setOperatorExtraData: operatorSet is not in stakeTree"
        );
        operatorExtraDatas[msg.sender][operatorSetId][operator] = extraData;
    }

    /// @inheritdoc IStakeRootCompendium
    function withdraw(uint32 operatorSetId, uint256 amount) external payable returns (uint256) {
        OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
        require(
            canWithdrawDepositBalance(operatorSet),
            "StakeRootCompendium.withdrawForOperatorSet: operator set is not old enough"
        );

        // debt any pending charge
        uint256 balance = _updateDepositInfo(operatorSet);
        if (balance < MIN_BALANCE_THRESHOLD + amount) {
            // withdraw all to avoid deposit balance dropping below minimum
            amount = balance;
            // remove from stake tree if applicable
            _removeFromStakeTree(operatorSet);
        }
        // debt the withdraw amount
        depositInfos[msg.sender][operatorSetId].balance -= uint96(amount);

        (bool success,) = payable(msg.sender).call{value: amount}("");
        require(success, "StakeRootCompendium.withdrawForOperatorSet: eth transfer failed");

        return amount;
    }

    /// CHARGE MANAGEMENT

    /// @inheritdoc IStakeRootCompendium
    function removeOperatorSetsFromStakeTree(
        OperatorSet[] calldata operatorSetsToRemove
    ) external {
        uint256 penalty = 0;
        for (uint256 i = 0; i < operatorSetsToRemove.length; i++) {
            if (_isInStakeTree(operatorSetsToRemove[i])) {
                uint256 depositBalance = _updateDepositInfo(operatorSetsToRemove[i]);
                // remove from stake tree if their deposit balance is below the minimum
                if (depositBalance < MIN_BALANCE_THRESHOLD) {
                    _removeFromStakeTree(operatorSetsToRemove[i]);
                    penalty += depositBalance;
                }
            }
        }
        // note this reward is subject to a race condition where anyone can claim the penalty if they submit their transaction first
        // it is the caller's responsibility to use private mempool to protect themselves from reverted transactions
        (bool success,) = payable(msg.sender).call{value: penalty}("");
        require(success, "StakeRootCompendium.updateBalances: eth transfer failed");
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
        require(
            calculationTimestamp % proofIntervalSeconds == 0,
            "StakeRootCompendium._postStakeRoot: timestamp must be a multiple of proofInterval"
        );
        // no length check here is ok because the initializer adds a default submission
        require(
            stakeRootSubmissions[stakeRootSubmissions.length - 1].calculationTimestamp != calculationTimestamp,
            "StakeRootCompendium._postStakeRoot: timestamp already posted"
        );
        // credit the charge recipient
        Snapshots.Snapshot memory _snapshot = totalChargeHistory._snapshots[indexChargePerProof];
        require(
            _snapshot._key <= calculationTimestamp,
            "StakeRootCompendium._postStakeRoot: timestamp of indexChargePerProof is greater than the calculationTimestamp"
        );
        require(
            totalChargeHistory.length() == indexChargePerProof + 1
                || uint256(totalChargeHistory._snapshots[indexChargePerProof + 1]._key) > calculationTimestamp,
            "StakeRootCompendium._postStakeRoot: indexChargePerProof is not valid"
        );
        stakeRootSubmissions.push(
            StakeRootSubmission({
                calculationTimestamp: uint32(calculationTimestamp),
                stakeRoot: stakeRoot,
                confirmed: false
            })
        );

        (bool success,) = payable(chargeRecipient).call{value: _snapshot._value}("");
        require(success, "StakeRootCompendium.withdrawForChargeRecipient: eth transfer failed");

        // interactions

        // note verify will be an external call, so adding to the end to apply the check, effect, interaction pattern to avoid reentrancy
        // TODO: verify proof
        // TODO: prevent race incentives and public mempool sniping, eg embed chargeRecipient in the proof
    }

    /// @inheritdoc IStakeRootCompendium
    function confirmStakeRoot(uint32 index, bytes32 stakeRoot) external {
        require(msg.sender == rootConfirmer, "StakeRootCompendium.confirmStakeRoot: only rootConfirmer can confirm");
        require(
            stakeRootSubmissions[index].stakeRoot == stakeRoot,
            "StakeRootCompendium.confirmStakeRoot: stake root does not match"
        );
        require(
            !stakeRootSubmissions[index].confirmed, "StakeRootCompendium.confirmStakeRoot: timestamp already confirmed"
        );
        stakeRootSubmissions[index].confirmed = true;
    }

    /// SET FUNCTIONS

    /// @inheritdoc IStakeRootCompendium
    function setMaxTotalCharge(
        uint96 _maxTotalCharge
    ) external onlyOwner {
        require(
            _maxTotalCharge >= totalChargeHistory.latest(),
            "StakeRootCompendium.setMaxTotalCharge: max total charge must be greater the current totalCharge"
        );
        maxTotalCharge = _maxTotalCharge;
    }

    /// @inheritdoc IStakeRootCompendium
    function setChargePerProof(uint96 _chargePerStrategy, uint96 _chargePerOperatorSet) external onlyOwner {
        _updateTotalCharge();
        chargePerStrategy = _chargePerStrategy;
        chargePerOperatorSet = _chargePerOperatorSet;
        _updateTotalCharge();
    }

    /// @inheritdoc IStakeRootCompendium
    function setProofIntervalSeconds(
        uint32 proofIntervalSeconds
    ) external onlyOwner {
        _updateTotalCharge();
        // we must not interrupt pending proof calculations by rugging the outstanding calculationTimestamps
        uint32 latestSubmittedCalculationTimestamp =
            stakeRootSubmissions[stakeRootSubmissions.length - 1].calculationTimestamp;
        require(
            latestSubmittedCalculationTimestamp == cumulativeChargeLastUpdatedTimestamp,
            "StakeRootCompendium.setProofIntervalSeconds: no proofs that have been charged but have not been submitteed"
        );
        proofIntervalSeconds = proofIntervalSeconds;
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
        uint224 operatorSetIndex = operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].latest();
        operatorSets[operatorSetIndex] = substituteOperatorSet;
        operatorSets.pop();

        // update the index of the operator sets
        // note when there is only one operator set left, the index will not be updated as the operator set will be removed in the next step
        operatorSetToIndex[substituteOperatorSet.avs][substituteOperatorSet.operatorSetId].push(
            uint32(block.timestamp), operatorSetIndex
        );
        operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].push(uint32(block.timestamp), REMOVED_INDEX);

        depositInfos[operatorSet.avs][operatorSet.operatorSetId].balance = 0;

        _updateTotalStrategies(
            operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length(), 0
        );
    }

    function _updateTotalStrategies(uint256 _countStrategiesBefore, uint256 _countStrategiesAfter) internal {
        totalStrategies = totalStrategies - _countStrategiesBefore + _countStrategiesAfter;
        _updateTotalCharge();
    }

    function _updateTotalCharge() internal {
        // note if totalStrategies is 0, the charge per proof will be 0, and provers should not post a proof
        uint256 totalCharge = operatorSets.length * chargePerOperatorSet + totalStrategies * chargePerStrategy;
        require(
            totalCharge <= maxTotalCharge,
            "StakeRootCompendium._updateChargePerProof: charge per proof exceeds max total charge"
        );
        totalChargeHistory.push(uint32(block.timestamp), uint224(totalCharge));
    }

    function _cumulativeCharges() internal view returns (uint32, uint96, uint96) {
        // calculate the total charge since the last update up until the latest calculation timestamp
        // note that there may be no corresponding stakeRootSubmission for the latest calculation timestamp
        // but if the calculationTimestamp is in the past, then it should be charged for, since proofs are being generated
        uint32 latestCalculationTimestamp = uint32(block.timestamp) - uint32(block.timestamp % proofIntervalSeconds);
        if (cumulativeChargeLastUpdatedTimestamp == latestCalculationTimestamp) {
            return (
                latestCalculationTimestamp,
                cumulativeChargePerOperatorSetLastUpdate,
                cumulativeChargePerStrategyLastUpdate
            );
        }
        uint256 numProofs = (latestCalculationTimestamp - cumulativeChargeLastUpdatedTimestamp) / proofIntervalSeconds;
        return (
            latestCalculationTimestamp,
            uint96(cumulativeChargePerOperatorSetLastUpdate + chargePerOperatorSet * numProofs),
            uint96(cumulativeChargePerStrategyLastUpdate + chargePerStrategy * numProofs)
        );
    }

    function _updateCumulativeCharge() internal {
        (uint32 latestCalculationTimestamp, uint96 cumulativeChargePerOperatorSet, uint96 cumulativeChargePerStrategy) =
            _cumulativeCharges();
        cumulativeChargeLastUpdatedTimestamp = latestCalculationTimestamp;
        cumulativeChargePerOperatorSetLastUpdate = cumulativeChargePerOperatorSet;
        cumulativeChargePerStrategyLastUpdate = cumulativeChargePerStrategy;
    }

    // updates the deposit balance for the operator set and returns the penalty if the operator set has fallen below the minimum deposit balance
    function _updateDepositInfo(
        OperatorSet memory operatorSet
    ) internal returns (uint256) {
        require(_isInStakeTree(operatorSet), "StakeRootCompendium._updateDepositInfo: operatorSet is not in stakeTree");

        (, uint256 cumulativeChargePerOperatorSet, uint256 cumulativeChargePerStrategy) = _cumulativeCharges();
        DepositInfo memory depositInfo = depositInfos[operatorSet.avs][operatorSet.operatorSetId];

        // subtract new total charge from last paid total charge
        uint256 pendingCharge = cumulativeChargePerOperatorSet - depositInfo.cumulativeChargePerOperatorSetLastPaid
            + (
                (cumulativeChargePerStrategy - depositInfo.cumulativeChargePerStrategyLastPaid)
                    * operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length()
            );

        uint256 balance = depositInfo.balance > pendingCharge ? depositInfo.balance - pendingCharge : 0;

        depositInfos[operatorSet.avs][operatorSet.operatorSetId] = DepositInfo({
            balance: uint96(balance),
            lastDemandIncreaseTimestamp: depositInfo.lastDemandIncreaseTimestamp,
            cumulativeChargePerOperatorSetLastPaid: uint96(cumulativeChargePerOperatorSet),
            cumulativeChargePerStrategyLastPaid: uint96(cumulativeChargePerStrategy)
        });

        return balance;
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
        return (numStrategies * chargePerStrategy + chargePerOperatorSet) * MIN_PREPAID_PROOFS;
    }

    /// @inheritdoc IStakeRootCompendium
    function canWithdrawDepositBalance(
        OperatorSet memory operatorSet
    ) public view returns (bool) {
        // they must have paid for all of their prepaid proofs before withdrawing after a demand increase
        return block.timestamp
            > depositInfos[operatorSet.avs][operatorSet.operatorSetId].lastDemandIncreaseTimestamp
                + MIN_PREPAID_PROOFS * proofIntervalSeconds;
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
        (, uint96 cumulativeChargePerOperatorSet, uint96 cumulativeChargePerStrategy) = _cumulativeCharges();
        uint256 pendingCharge = uint256(
            cumulativeChargePerOperatorSet - depositInfo.cumulativeChargePerOperatorSetLastPaid
        )
            + uint256(cumulativeChargePerStrategy - depositInfo.cumulativeChargePerStrategyLastPaid)
                * operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length();

        return depositInfo.balance > pendingCharge ? depositInfo.balance - pendingCharge : 0;
    }

    // STAKE ROOT CALCULATION

    /// @inheritdoc IStakeRootCompendium
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

    /// @inheritdoc IStakeRootCompendium
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
        address[] memory operators =
            avsDirectory.getOperatorsInOperatorSet(operatorSet, startOperatorIndex, numOperators);
        (IStrategy[] memory strategies, uint256[] memory multipliers) = _getStrategiesAndMultipliers(operatorSet);

        OperatorLeaf[] memory operatorLeaves = new OperatorLeaf[](operators.length);
        for (uint256 i = 0; i < operatorLeaves.length; i++) {
            // calculate the weighted sum of the operator's shares for the strategies given the multipliers
            (uint256 delegatedStake, uint256 slashableStake) =
                _getStakes(operatorSet, strategies, multipliers, operators[i]);

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
        OperatorSet calldata operatorSet,
        address[] calldata operators,
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
        address prevOperator;
        bytes32[] memory operatorLeavesHashes = new bytes32[](operatorLeaves.length);
        for (uint256 i = 0; i < operatorLeaves.length; i++) {
            require(
                uint160(prevOperator) < uint160(operators[i]), "AVSSyncTree.getOperatorSetRoot: operators not sorted"
            );
            prevOperator = operators[i];

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

    // in case of charge problems
    receive() external payable {}
}
