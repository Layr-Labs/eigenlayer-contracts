// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IAllocationManager.sol";
import "../libraries/Snapshots.sol";

abstract contract CompressorRegistryStorage is OwnableUpgradeable {
    using Snapshots for Snapshots.History;
    using EnumerableMap for EnumerableMap.AddressToUintMap;
    using EnumerableSet for EnumerableSet.Bytes32Set;

    // TODO: Rename these...
    error NonexistentOperatorSet();
    error NonexistentStrategy();
    error InsufficientDepositBalance();
    error NotInStakeTree();
    error OperatorSetNotOldEnough();
    error EthTransferFailed();
    error TimestampNotMultipleOfProofInterval();
    error TimestampAlreadyPosted();
    error TimestampOfIndexChargePerProofIsGreaterThanCalculationTimestamp();
    error IndexChargePerProofNotValid();
    error OnlyRootConfirmerCanConfirm();
    error StakeRootDoesNotMatch();
    error TimestampAlreadyConfirmed();
    error MaxTotalChargeMustBeGreaterThanTheCurrentTotalCharge();
    error NoProofsThatHaveBeenChargedButNotSubmitted();
    error ChargePerProofExceedsMax();
    error InputArrayLengthMismatch();
    error InputCorrelatedVariableMismatch();
    error OperatorSetIndexOutOfBounds();
    error OperatorSetMustExist();
    error OperatorSetSizeMismatch();

    struct StrategyAndMultiplier {
        IStrategy strategy;
        uint96 multiplier;
    }

    /**
     * @dev Struct containing charges for operator sets, strategies, and max total charge.
     * @param chargePerOperatorSet The linear charge per proof in the number of strategies.
     * @param chargePerStrategy The constant charge per proof.
     * @param The max total charge for a stakeroot proof, used to bound computation offchain.
     */
    struct ChargeParams {
        uint96 chargePerOperatorSet;
        uint96 chargePerStrategy;
        uint96 maxChargePerProof;
    }
    
    /**
     * @dev Struct containing info about cumulative charges.
     * @param chargePerOperatorSet The cumulative constant charge per operator set since deployment.
     * @param chargePerStrategy The cumulative linear charge per strategy per operator set since deployment.
     * @param lastUpdateTimestamp The last time cumulative charges were updated.
     * @param proofIntervalSeconds The interval in seconds at which proofs can be posted.
     */

    struct CumulativeChargeParams {
        uint96 chargePerOperatorSet;
        uint96 chargePerStrategy;
        uint32 lastUpdateTimestamp;
        uint32 proofIntervalSeconds;
    }

    struct DepositInfo {
        // the balance of the operatorSet (includes pending deductions)
        uint96 balance;
        // the timestamp of the operatorSets latest deposit or increase in number of strategies.
        // withdrawals of deposit balance are bounded by paying for MIN_PREPAID_PROOFS proofs since
        // ones latest demand increase
        uint32 lastDemandIncreaseTimestamp;
        uint32 countStrategies;
        // the cumulativeChargePerOperatorSet at the time of the lastest deduction from the deposit balance
        // used in making further deductions
        uint96 cumulativeChargePerOperatorSetLastPaid;
        // the cumulativeChargePerStrategy at the time of the lastest deduction from the deposit balance
        // used in making further deductions
        uint96 cumulativeChargePerStrategyLastPaid;
    }

    /// @notice the delegation manager contract
    IDelegationManager public immutable delegationManager;
    /// @notice the AVS directory contract
    IAVSDirectory public immutable avsDirectory;
    /// @notice the allocation manager contract
    IAllocationManager public immutable allocationManager;

    /// @notice the minimum balance that must be maintained for an operatorSet
    /// @dev this balance compensates gas costs to deregister an operatorSet
    uint256 public immutable MIN_BALANCE_THRESHOLD;

    /// @notice the minimum number of proofs that an operatorSet's deposit balance needs to cover and
    /// the number of proofs they must pay for since their latest reconfiguration
    /// @dev this prevents de-registering an operatorSet immediately after reconfiguring
    uint256 public immutable MIN_PREPAID_PROOFS;

    /// @notice the total number of strategies among all operator sets (with duplicates)
    uint256 public totalStrategies;

    address public compressor;

    /// @notice the total charge for a proofs at a certain time depending on the number of strategies
    Snapshots.History internal chargePerProof;
    /// @dev Contains cumulative charges for operator sets, strategies, and max total charge.
    ChargeParams public chargeParams;
    /// @dev Contains info about cumulative charges.
    CumulativeChargeParams public cumulativeChargeParams;
    /// @notice deposit balance to be deducted for operatorSets
    mapping(address => mapping(uint32 => DepositInfo)) public depositInfos;
    /// @notice the extraData for each operatorSet

    constructor(
        IDelegationManager _delegationManager,
        IAVSDirectory _avsDirectory,
        IAllocationManager _allocationManager,
        uint256 _minBalanceThreshold,
        uint256 _minPrepaidProofs
    ) {
        delegationManager = _delegationManager;
        avsDirectory = _avsDirectory;
        allocationManager = _allocationManager;
        MIN_BALANCE_THRESHOLD = _minBalanceThreshold;
        MIN_PREPAID_PROOFS = _minPrepaidProofs;
    }
}
