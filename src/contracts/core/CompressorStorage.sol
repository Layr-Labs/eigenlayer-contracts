// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";
import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IAllocationManager.sol";
import "../interfaces/IStakeRootCompendium.sol";
import "../libraries/Snapshots.sol";

abstract contract CompressorStorage is OwnableUpgradeable {
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

    struct OperatorLeaf {
        uint256 delegatedStake;
        uint256 slashableStake;
        bytes32 extraData;
    }

    /// @notice the placeholder index used for operator sets that are removed from the StakeTree
    uint32 public constant REMOVED_INDEX = type(uint32).max;

    /// @notice the delegation manager contract
    IDelegationManager public immutable delegationManager;
    /// @notice the AVS directory contract
    IAVSDirectory public immutable avsDirectory;
    /// @notice the allocation manager contract
    IAllocationManager public immutable allocationManager;

    /// @notice the verifier contract that will be used to verify snark proofs
    address public immutable verifier;
    /// @notice the id of the program being verified when roots are posted
    bytes32 public immutable imageId;

    /// @notice the address allowed to confirm
    address public confirmer;

    address public manager;

    uint32 public proofIntervalSeconds;
    
    /// @notice list of operator sets that have been configured to be in the StakeTree
    OperatorSet[] public operatorSets;

    /// @notice the extraData for each operatorSet
    mapping(address => mapping(uint32 => bytes32)) internal operatorSetExtraDatas;
    /// @notice map from operator set to a trace of their index over time
    mapping(address => mapping(uint32 => Snapshots.History)) internal operatorSetToIndex;
    /// @notice the extraData for each operator in each operator set
    mapping(address => mapping(uint32 => mapping(address => bytes32))) internal operatorExtraDatas;
    /// @notice the strategies and multipliers for each operator set
    mapping(address => mapping(uint32 => EnumerableMap.AddressToUintMap)) internal operatorSetToStrategyAndMultipliers;

    struct CompressedState {
        bytes32 root;
        uint32 calculationTimestamp; // the timestamp of the state the stakeRoot was calculated against
        uint32 submissionTimestamp; // the timestamp of the submission
        bool confirmed; // whether the rootConfimer has confirmed the root
    }

    /// @notice the stake root submissions
    CompressedState[] public compressedStates;

    constructor(
        IDelegationManager _delegationManager,
        IAVSDirectory _avsDirectory,
        IAllocationManager _allocationManager,
        address _verifier,
        bytes32 _imageId
    ) {
        delegationManager = _delegationManager;
        avsDirectory = _avsDirectory;
        allocationManager = _allocationManager;
        // note verifier and imageId are immutable and set by implementation contract
        // since proof verification is in the hot path, this is a gas optimization to avoid calling the storage contract for verifier and imageId
        verifier = _verifier;
        imageId = _imageId;
    }
}
