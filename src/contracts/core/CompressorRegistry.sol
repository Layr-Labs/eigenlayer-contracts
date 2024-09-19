// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import "../interfaces/IStrategy.sol";
import "../libraries/Merkle.sol";
import "./CompressorRegistryStorage.sol";

contract CompressorRegistry is CompressorRegistryStorage {
    using Snapshots for Snapshots.History;
    using EnumerableMap for EnumerableMap.AddressToUintMap;
    using EnumerableSet for EnumerableSet.Bytes32Set;

    constructor(
        IDelegationManager _delegationManager,
        IAVSDirectory _avsDirectory,
        IAllocationManager _allocationManager,
        uint256 _minBalanceThreshold,
        uint256 _minPrepaidProofs
    )
        CompressorRegistryStorage(
            _delegationManager,
            _avsDirectory,
            _allocationManager,
            _minBalanceThreshold,
            _minPrepaidProofs
        )
    {}

    function initialize(
        address _owner,
        address _compressor,
        uint32 _proofIntervalSeconds,
        ChargeParams memory _chargeParams
    ) public initializer {
        __Ownable_init();
        _transferOwnership(_owner);

        compressor = _compressor;
        chargeParams = _chargeParams;
        cumulativeChargeParams.proofIntervalSeconds = _proofIntervalSeconds;
    }

    /// OPERATORSET CONFIGURATION

    
    function deposit(
        OperatorSet calldata operatorSet
    ) external payable {
        DepositInfo storage depositInfo = depositInfos[operatorSet.avs][operatorSet.operatorSetId];
        depositInfo.balance += uint96(msg.value);
    }

    function register(OperatorSet calldata operatorSet) external {
        (, uint256 cumulativeChargePerOperatorSet, uint256 cumulativeChargePerStrategy) =
            _calculateCumulativeCharges();

        DepositInfo storage depositInfo = depositInfos[operatorSet.avs][operatorSet.operatorSetId];
        depositInfo.lastDemandIncreaseTimestamp = uint32(block.timestamp);
        depositInfo.cumulativeChargePerOperatorSetLastPaid = uint96(cumulativeChargePerOperatorSet);
        depositInfo.cumulativeChargePerStrategyLastPaid = uint96(cumulativeChargePerStrategy);

        _updateTotalStrategies(
            0, depositInfo.countStrategies
        );

        // note that they've shown increased demand for proofs
        depositInfo.lastDemandIncreaseTimestamp = uint32(block.timestamp);
        // make sure they have enough to pay for MIN_PROOFS_PREPAID
        require(
            depositInfo.balance
                >= minBalance(
                    depositInfo.countStrategies
                ),
            InsufficientDepositBalance()
        );

        // todo: call compressor.register(operatorSet);
    }

    
    function addOrModifyStrategiesAndMultipliers(
        uint32 operatorSetId,
        StrategyAndMultiplier[] calldata _strategiesAndMultipliers
    ) external {
        OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
        // update the deposit balance for the operator set whenever number of strategies is changed
        _updateDepositInfo(operatorSet);

        uint256 numStrategiesBefore;
        uint256 numStrategiesAfter;

        // todo: call (numStrategiesBefore, numStrategiesAfter) = compressor.addOrModifyStrategiesAndMultipliers(operatorSet, _strategiesAndMultipliers);

        DepositInfo storage depositInfo = depositInfos[operatorSet.avs][operatorSet.operatorSetId];

        // make sure they have enough to pay for MIN_PREPAID_PROOFS
        require(depositInfo.balance >= minBalance(numStrategiesAfter), InsufficientDepositBalance());
        // note that they've shown increased demand for proofs
        depositInfo.countStrategies = uint32(numStrategiesAfter);
        depositInfo.lastDemandIncreaseTimestamp = uint32(block.timestamp);

        // only adding new strategies to count
        _updateTotalStrategies(numStrategiesBefore, numStrategiesAfter);
    }

    
    function removeStrategiesAndMultipliers(uint32 operatorSetId, IStrategy[] calldata strategies) external {
        OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
        // update the deposit balance for the operator set whenever number of strategies is changed
        _updateDepositInfo(operatorSet);

        uint256 numStrategiesBefore;
        uint256 numStrategiesAfter;
        // todo: call (numStrategiesBefore, numStrategiesAfter) = compressor.removeStrategiesAndMultipliers(operatorSet, strategies);

        _updateTotalStrategies(numStrategiesBefore, numStrategiesAfter);
    }

    
    function setOperatorSetExtraData(uint32 operatorSetId, bytes32 extraData) external {
        // todo call compressor.setOperatorSetExtraData(OperatorSet({avs: msg.sender, operatorSetId: operatorSetId}), extraData);
    }

    
    function setOperatorExtraData(uint32 operatorSetId, address operator, bytes32 extraData) external {
        // todo call compressor.setOperatorExtraData(OperatorSet({avs: msg.sender, operatorSetId: operatorSetId}), operator, extraData);
    }


    function withdraw(uint32 operatorSetId, uint256 amount) external payable returns (uint256) {
        OperatorSet memory operatorSet = OperatorSet({avs: msg.sender, operatorSetId: operatorSetId});
        require(canWithdraw(operatorSet), OperatorSetNotOldEnough());

        // debt any pending charge
        uint256 balance = _updateDepositInfo(operatorSet);
        if (balance < MIN_BALANCE_THRESHOLD + amount) {
            // withdraw all to avoid deposit balance dropping below minimum
            amount = balance;
            // remove from stake tree if applicable
            // todo call compressor.removeOperatorSetsFromStakeTree(operatorSet);
        }
        // debt the withdraw amount
        depositInfos[msg.sender][operatorSetId].balance -= uint96(amount);

        _safeTransferETH(msg.sender, amount);

        return amount;
    }

    /// CHARGE MANAGEMENT

    function removeOperatorSetsFromStakeTree(
        OperatorSet[] calldata operatorSetsToRemove
    ) external {
        uint256 penalty = 0;
        for (uint256 i = 0; i < operatorSetsToRemove.length; i++) {
            uint256 depositBalance = _updateDepositInfo(operatorSetsToRemove[i]);
            // remove from stake tree if their deposit balance is below the minimum
            if (depositBalance < MIN_BALANCE_THRESHOLD) {
                    // todo call compressor.removeOperatorSetsFromStakeTree(operatorSetsToRemove[i]);
                penalty += depositBalance;
            }
        }
        // note this reward is subject to a race condition where anyone can claim the penalty if they submit their transaction first
        // it is the caller's responsibility to use private mempool to protect themselves from reverted transactions
        _safeTransferETH(msg.sender, penalty);
    }

    /// POSTING ROOTS AND BLACKLISTING

    function withdrawPayment(
    ) external {
        // todo handle payment of rewards for compression service providers
    }

    /// SET FUNCTIONS

    function setMaxChargePerProof(
        uint96 _maxChargePerProof
    ) external onlyOwner {
        require(_maxChargePerProof >= chargePerProof.latest(), MaxTotalChargeMustBeGreaterThanTheCurrentTotalCharge());
        chargeParams.maxChargePerProof = _maxChargePerProof;
    }

    function setChargePerProof(uint96 _chargePerStrategy, uint96 _chargePerOperatorSet) external onlyOwner {
        _updateCumulativeCharge();
        ChargeParams storage params = chargeParams;
        params.chargePerStrategy = _chargePerStrategy;
        params.chargePerOperatorSet = _chargePerOperatorSet;
        _updateChargePerProof();
    }

    function setProofIntervalSeconds(
        uint32 _proofIntervalSeconds
    ) external onlyOwner {
        _updateCumulativeCharge();
        // todo we must not interrupt pending proof calculations by rugging the outstanding calculationTimestamps
        // todo call compressor.setProofIntervalSeconds(_proofIntervalSeconds);
        cumulativeChargeParams.proofIntervalSeconds = _proofIntervalSeconds;
    }

    function setConfirmer(
        address _confirmer
    ) public onlyOwner {
        // todo call compressor.setConfirmer(_confirmer);
    }

    /// INTERNAL FUNCTIONS

    function _updateTotalStrategies(uint256 _countStrategiesBefore, uint256 _countStrategiesAfter) internal {
        totalStrategies = totalStrategies - _countStrategiesBefore + _countStrategiesAfter;
        _updateChargePerProof();
    }

    function _updateChargePerProof() internal {
        // note if totalStrategies is 0, the charge per proof will be 0, and provers should not post a proof
        uint256 countOperatorSet;
        
        // todo call countOperatorSet = compressor.getOperatorSetCount();
        uint256 _chargePerProof =
            countOperatorSet * chargeParams.chargePerOperatorSet + totalStrategies * chargeParams.chargePerStrategy;
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
        (, uint256 cumulativeChargePerOperatorSet, uint256 cumulativeChargePerStrategy) = _calculateCumulativeCharges();
        DepositInfo storage depositInfo = depositInfos[operatorSet.avs][operatorSet.operatorSetId];

        // subtract new total charge from last paid total charge
        uint256 pendingCharge = cumulativeChargePerOperatorSet - depositInfo.cumulativeChargePerOperatorSetLastPaid
            + (
                (cumulativeChargePerStrategy - depositInfo.cumulativeChargePerStrategyLastPaid)
                    * depositInfo.countStrategies
            );

        balance = depositInfo.balance > pendingCharge ? depositInfo.balance - pendingCharge : 0;

        depositInfo.balance = uint96(balance);
        depositInfo.cumulativeChargePerOperatorSetLastPaid = uint96(cumulativeChargePerOperatorSet);
        depositInfo.cumulativeChargePerOperatorSetLastPaid = uint96(cumulativeChargePerStrategy);
    }

    /// VIEW FUNCTIONS

    function minBalance(
        uint256 numStrategies
    ) public view returns (uint256) {
        ChargeParams memory charges = chargeParams;
        return (numStrategies * charges.chargePerStrategy + charges.chargePerOperatorSet) * MIN_PREPAID_PROOFS;
    }

    function canWithdraw(
        OperatorSet memory operatorSet
    ) public view returns (bool) {
        // they must have paid for all of their prepaid proofs before withdrawing after a demand increase
        return block.timestamp
            > depositInfos[operatorSet.avs][operatorSet.operatorSetId].lastDemandIncreaseTimestamp
                + MIN_PREPAID_PROOFS * cumulativeChargeParams.proofIntervalSeconds;
    }

    function getBalance(
        OperatorSet memory operatorSet
    ) external view returns (uint256 balance) {
        DepositInfo memory depositInfo = depositInfos[operatorSet.avs][operatorSet.operatorSetId];
        (, uint96 cumulativeChargePerOperatorSet, uint96 cumulativeChargePerStrategy) = _calculateCumulativeCharges();
        uint256 pendingCharge = uint256(
            cumulativeChargePerOperatorSet - depositInfo.cumulativeChargePerOperatorSetLastPaid
        )
            + uint256(cumulativeChargePerStrategy - depositInfo.cumulativeChargePerStrategyLastPaid)
                * depositInfo.countStrategies;

        return depositInfo.balance > pendingCharge ? depositInfo.balance - pendingCharge : 0;
    }

    function _safeTransferETH(address to, uint256 amount) internal {
        (bool success,) = to.call{value: amount}("");
        require(success, EthTransferFailed());
    }
}
