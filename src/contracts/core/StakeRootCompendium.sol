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
        uint96 _maxTotalCharge,
        uint96 _chargePerStrategy,
        uint96 _chargePerOperatorSet
    ) public initializer {
        __Ownable_init();
        _transferOwnership(_owner);

        rootConfirmer = _rootConfirmer;

        stakerootCharges = StakerootCharges({
            chargePerOperatorSet: _chargePerOperatorSet,
            chargePerStrategy: _chargePerStrategy,
            maxTotalCharge: _maxTotalCharge
        });

        stakerootCumulativeCharges.proofIntervalSeconds = _proofIntervalSeconds;

        stakeRootSubmissions.push(
            StakeRootSubmission({calculationTimestamp: 0, stakeRoot: bytes32(0), confirmed: false})
        );
    }

    /// OPERATORSET CONFIGURATION

    /// @inheritdoc IStakeRootCompendium
    function deposit(
        OperatorSet calldata operatorSet
    ) external payable {
        DepositInfo storage depositInfo = depositInfos[operatorSet.avs][operatorSet.operatorSetId];
        if (!_isInStakeTree(operatorSet)) {
            (, uint256 cumulativeChargePerOperatorSet, uint256 cumulativeChargePerStrategy) =
                _calculateCumulativeCharges(stakerootCharges, stakerootCumulativeCharges);

            depositInfo.lastDemandIncreaseTimestamp = uint32(block.timestamp);
            depositInfo.cumulativeChargePerOperatorSetLastPaid = uint96(cumulativeChargePerOperatorSet);
            depositInfo.cumulativeChargePerStrategyLastPaid = uint96(cumulativeChargePerStrategy);

            _updateTotalStrategies(
                0, operatorSetToStrategyAndMultipliers[operatorSet.avs][operatorSet.operatorSetId].length()
            );

            operatorSetToIndex[operatorSet.avs][operatorSet.operatorSetId].push(
                uint32(block.timestamp), uint224(operatorSets.length)
            );

            operatorSets.push(operatorSet);
        }

        depositInfo.balance += uint96(msg.value);
        // note that they've shown increased demand for proofs
        depositInfo.lastDemandIncreaseTimestamp = uint32(block.timestamp);
        // make sure they have enough to pay for MIN_PROOFS_PREPAID
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

        DepositInfo storage depositInfo = depositInfos[operatorSet.avs][operatorSet.operatorSetId];

        // make sure they have enough to pay for MIN_PREPAID_PROOFS
        require(depositInfo.balance >= minDepositBalance(numStrategiesAfter), InsufficientDepositBalance());
        // note that they've shown increased demand for proofs
        depositInfo.lastDemandIncreaseTimestamp = uint32(block.timestamp);

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
                NonexistentStrategy()
            );
        }

        _updateTotalStrategies(numStrategiesBefore, numStrategiesBefore - strategies.length);
    }

    /// @inheritdoc IStakeRootCompendium
    function setOperatorSetExtraData(uint32 operatorSetId, bytes32 extraData) external {
        require(
            _isInStakeTree(OperatorSet({avs: msg.sender, operatorSetId: operatorSetId})),
            StakeTreeMustIncludeOperatorSet()
        );
        operatorSetExtraDatas[msg.sender][operatorSetId] = extraData;
    }

    /// @inheritdoc IStakeRootCompendium
    function setOperatorExtraData(uint32 operatorSetId, address operator, bytes32 extraData) external {
        require(
            _isInStakeTree(OperatorSet({avs: msg.sender, operatorSetId: operatorSetId})),
            StakeTreeMustIncludeOperatorSet()
        );
        operatorExtraDatas[msg.sender][operatorSetId][operator] = extraData;
    }

    /// @inheritdoc IStakeRootCompendium
    function withdraw(uint32 operatorSetId, uint256 amount) external payable returns (uint256) {
        OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
        require(canWithdrawDepositBalance(operatorSet), OperatorSetNotOldEnough());

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

        safeTransferETH(msg.sender, amount);

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
        safeTransferETH(msg.sender, penalty);
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
            calculationTimestamp % stakerootCumulativeCharges.proofIntervalSeconds == 0,
            TimestampNotMultipleOfProofInterval()
        );
        // no length check here is ok because the initializer adds a default submission
        require(
            stakeRootSubmissions[stakeRootSubmissions.length - 1].calculationTimestamp != calculationTimestamp,
            TimestampAlreadyPosted()
        );
        // credit the charge recipient
        Snapshots.Snapshot memory _snapshot = totalChargeHistory._snapshots[indexChargePerProof];
        require(
            _snapshot._key <= calculationTimestamp, TimestampOfIndexChargePerProofIsGreaterThanCalculationTimestamp()
        );
        require(
            totalChargeHistory.length() == indexChargePerProof + 1
                || uint256(totalChargeHistory._snapshots[indexChargePerProof + 1]._key) > calculationTimestamp,
            IndexChargePerProofNotValid()
        );
        stakeRootSubmissions.push(
            StakeRootSubmission({
                calculationTimestamp: uint32(calculationTimestamp),
                stakeRoot: stakeRoot,
                confirmed: false
            })
        );

        safeTransferETH(chargeRecipient, _snapshot._value);
        // interactions

        // note verify will be an external call, so adding to the end to apply the check, effect, interaction pattern to avoid reentrancy
        // TODO: verify proof
        // TODO: prevent race incentives and public mempool sniping, eg embed chargeRecipient in the proof
    }

    /// @inheritdoc IStakeRootCompendium
    function confirmStakeRoot(uint32 index, bytes32 stakeRoot) external {
        require(msg.sender == rootConfirmer, OnlyRootConfirmerCanConfirm());
        require(stakeRootSubmissions[index].stakeRoot == stakeRoot, StakeRootDoesNotMatch());
        require(!stakeRootSubmissions[index].confirmed, TimestampAlreadyConfirmed());
        stakeRootSubmissions[index].confirmed = true;
    }

    /// SET FUNCTIONS

    /// @inheritdoc IStakeRootCompendium
    function setMaxTotalCharge(
        uint96 _maxTotalCharge
    ) external onlyOwner {
        require(_maxTotalCharge >= totalChargeHistory.latest(), MaxTotalChargeMustBeGreaterThanTheCurrentTotalCharge());
        stakerootCharges.maxTotalCharge = _maxTotalCharge;
    }

    /// @inheritdoc IStakeRootCompendium
    function setChargePerProof(uint96 _chargePerStrategy, uint96 _chargePerOperatorSet) external onlyOwner {
        StakerootCharges storage charges = stakerootCharges;
        _updateTotalCharge(charges);
        charges.chargePerStrategy = _chargePerStrategy;
        charges.chargePerOperatorSet = _chargePerOperatorSet;
        _updateTotalCharge(charges);
    }

    /// @inheritdoc IStakeRootCompendium
    function setProofIntervalSeconds(
        uint32 proofIntervalSeconds
    ) external onlyOwner {
        StakerootCumulativeCharges storage cumulativeCharges = stakerootCumulativeCharges;
        _updateTotalCharge(stakerootCharges);
        // we must not interrupt pending proof calculations by rugging the outstanding calculationTimestamps
        require(
            stakeRootSubmissions[stakeRootSubmissions.length - 1].calculationTimestamp
                == cumulativeCharges.lastUpdateTimestamp,
            NoProofsThatHaveBeenChargedButNotSubmitted()
        );
        cumulativeCharges.proofIntervalSeconds = proofIntervalSeconds;
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
        _updateTotalCharge(stakerootCharges);
    }

    function _updateTotalCharge(
        StakerootCharges memory charges
    ) internal {
        // note if totalStrategies is 0, the charge per proof will be 0, and provers should not post a proof
        uint256 totalCharge =
            operatorSets.length * charges.chargePerOperatorSet + totalStrategies * charges.chargePerStrategy;
        require(totalCharge <= charges.maxTotalCharge, ChargePerProofExceedsMaxTotalCharge());
        totalChargeHistory.push(uint32(block.timestamp), uint224(totalCharge));
    }

    function _calculateCumulativeCharges(
        StakerootCharges memory charges,
        StakerootCumulativeCharges memory cumulativeCharges
    ) internal view returns (uint32, uint96, uint96) {
        // calculate the total charge since the last update up until the latest calculation timestamp
        // note that there may be no corresponding stakeRootSubmission for the latest calculation timestamp
        // but if the calculationTimestamp is in the past, then it should be charged for, since proofs are being generated
        uint32 latestCalculationTimestamp =
            uint32(block.timestamp) - uint32(block.timestamp % cumulativeCharges.proofIntervalSeconds);

        if (cumulativeCharges.lastUpdateTimestamp == latestCalculationTimestamp) {
            return (
                latestCalculationTimestamp, cumulativeCharges.chargePerOperatorSet, cumulativeCharges.chargePerStrategy
            );
        }

        uint256 numProofs = (latestCalculationTimestamp - cumulativeCharges.lastUpdateTimestamp)
            / cumulativeCharges.proofIntervalSeconds;

        return (
            latestCalculationTimestamp,
            uint96(cumulativeCharges.chargePerOperatorSet + charges.chargePerOperatorSet * numProofs),
            uint96(cumulativeCharges.chargePerStrategy + charges.chargePerStrategy * numProofs)
        );
    }

    function _updateCumulativeCharge() internal {
        (uint32 lastUpdateTimestamp, uint96 cumulativeChargePerOperatorSet, uint96 cumulativeChargePerStrategy) =
            _calculateCumulativeCharges(stakerootCharges, stakerootCumulativeCharges);

        stakerootCumulativeCharges = StakerootCumulativeCharges({
            chargePerOperatorSet: cumulativeChargePerOperatorSet,
            chargePerStrategy: cumulativeChargePerStrategy,
            lastUpdateTimestamp: lastUpdateTimestamp,
            proofIntervalSeconds: stakerootCumulativeCharges.proofIntervalSeconds
        });
    }

    // updates the deposit balance for the operator set and returns the penalty if the operator set has fallen below the minimum deposit balance
    function _updateDepositInfo(
        OperatorSet memory operatorSet
    ) internal returns (uint256 balance) {
        require(_isInStakeTree(operatorSet), StakeTreeMustIncludeOperatorSet());

        (, uint256 cumulativeChargePerOperatorSet, uint256 cumulativeChargePerStrategy) =
            _calculateCumulativeCharges(stakerootCharges, stakerootCumulativeCharges);
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
        StakerootCharges memory charges = stakerootCharges;
        return (numStrategies * charges.chargePerStrategy + charges.chargePerOperatorSet) * MIN_PREPAID_PROOFS;
    }

    /// @inheritdoc IStakeRootCompendium
    function canWithdrawDepositBalance(
        OperatorSet memory operatorSet
    ) public view returns (bool) {
        // they must have paid for all of their prepaid proofs before withdrawing after a demand increase
        return block.timestamp
            > depositInfos[operatorSet.avs][operatorSet.operatorSetId].lastDemandIncreaseTimestamp
                + MIN_PREPAID_PROOFS * stakerootCumulativeCharges.proofIntervalSeconds;
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
        (, uint96 cumulativeChargePerOperatorSet, uint96 cumulativeChargePerStrategy) =
            _calculateCumulativeCharges(stakerootCharges, stakerootCumulativeCharges);
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
        address avs,
        uint32[] calldata operatorSetIdsInStakeTree,
        bytes32[] calldata operatorSetRoots
    ) external view returns (bytes32) {
        // TODO: This fn should revert if mismatched parameters are passed in due to out of bounds
        // array access, see if these checks can be removed.
        require(operatorSets.length == operatorSetIdsInStakeTree.length, InputArrayLengthMismatch());
        require(operatorSetIdsInStakeTree.length == operatorSetRoots.length, InputArrayLengthMismatch());

        for (uint256 i = 0; i < operatorSetIdsInStakeTree.length; i++) {
            require(operatorSets[i].operatorSetId == operatorSetIdsInStakeTree[i], InputCorrelatedVariableMismatch());
        }

        return Merkle.merkleizeKeccak256(operatorSetRoots);
    }

    /// @inheritdoc IStakeRootCompendium
    function getOperatorSetLeaves(
        uint256 operatorSetIndex,
        uint256 startOperatorIndex,
        uint256 numOperators
    ) external view returns (OperatorSet memory, address[] memory, OperatorLeaf[] memory) {
        require(operatorSetIndex < operatorSets.length, OutOfBounds());
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
        require(avsDirectory.isOperatorSet(operatorSet.avs, operatorSet.operatorSetId), OperatorSetMustExist());
        require(
            operatorLeaves.length == avsDirectory.getNumOperatorsInOperatorSet(operatorSet), OperatorSetSizeMismatch()
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

    function safeTransferETH(address to, uint256 amount) internal {
        /// @solidity memory-safe-assembly
        assembly {
            if iszero(call(gas(), to, amount, codesize(), 0x00, codesize(), 0x00)) {
                mstore(0x00, 0xb12d13eb) // `ETHTransferFailed()`.
                revert(0x1c, 0x04)
            }
        }
    }

    // in case of charge problems
    receive() external payable {}
}
