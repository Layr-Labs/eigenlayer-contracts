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

abstract contract StakeRootCompendiumStorage is IStakeRootCompendium, OwnableUpgradeable {
    using Checkpoints for Checkpoints.History;
    using EnumerableMap for EnumerableMap.AddressToUintMap;
    using EnumerableSet for EnumerableSet.Bytes32Set;

    /// @notice the delegation manager contract
    IDelegationManager public immutable delegationManager;
    /// @notice the AVS directory contract
    IAVSDirectory public immutable avsDirectory;

    /// @notice the maximum total charge for a proof
    uint256 immutable public MAX_TOTAL_CHARGE;

    /// @notice the minimum balance that must be maintained for an operatorSet
    /// @dev this balance compensates gas costs to deregister an operatorSet
    uint256 immutable public MIN_BALANCE_THRESHOLD;

    /// @notice the placeholder index used for operator sets that are removed from the StakeTree
    uint32 public constant REMOVED_INDEX = type(uint32).max;

    /// @notice the minimum number of proofs that an operatorSet's deposit balance needs to cover and 
    /// the number of proofs they must pay for since their latest reconfiguration
    /// @dev this prevents de-registering an operatorSet immediately after reconfiguring
    uint256 immutable public MIN_PROOFS_DURATION;

    /// @notice the verifier contract that will be used to verify snark proofs
    address public immutable verifier;
    /// @notice the id of the program being verified when roots are posted
    bytes32 public immutable imageId;

    /// @notice the interval in seconds at which proofs can be posted
    uint32 public proofIntervalSeconds;
    /// @notice the address allowed to confirm roots
    address public rootConfirmer;
    /// @notice the linear charge per proof in the number of strategies
    uint96 public chargePerOperatorSet;
    /// @notice the constant charge per proof
    uint96 public chargePerStrategy;

    uint32 public totalChargeLastUpdatedTimestamp;
    /// @notice the total constant charge per proof since deployment
    uint96 public totalChargePerOperatorSetLastUpdate;
    /// @notice the total linear charge per proof since deployment
    uint96 public totalChargePerStrategyLastUpdate;
    /// @notice deposit balance to be deducted for operatorSets
    mapping(address => mapping(uint32 => DepositInfo)) public depositInfos;

    /// @notice map from operator set to a trace of their index over time
    mapping(address => mapping(uint32 => Checkpoints.History)) internal operatorSetToIndex;
    /// @notice list of operator sets that have been configured to be in the StakeTree
    IAVSDirectory.OperatorSet[] public operatorSets;
    
    /// @notice the total number of strategies among all operator sets (with duplicates)
    uint256 public totalStrategies;
    /// @notice the total charge for a proofs at a certain time depending on the number of strategies
    Checkpoints.History internal chargePerProofHistory;

    /// @notice the strategies and multipliers for each operator set
    mapping(address => mapping(uint32 => EnumerableMap.AddressToUintMap)) internal operatorSetToStrategyAndMultipliers;
    /// @notice the extraData for each operatorSet
    mapping(address => mapping(uint32 => bytes32)) internal operatorSetExtraDatas;
    /// @notice the extraData for each operator in each operator set
    mapping(address => mapping(uint32 => mapping(address => bytes32))) internal operatorExtraDatas;

    /// @notice the stake root submissions
    IStakeRootCompendium.StakeRootSubmission[] public stakeRootSubmissions;
    
    constructor(
        IDelegationManager _delegationManager,
        IAVSDirectory _avsDirectory,
        uint256 _maxTotalCharge,
        uint256 _minBalanceThreshold,
        uint256 _minProofsDuration,
        address _verifier,
        bytes32 _imageId
    ) {
        delegationManager = _delegationManager;
        avsDirectory = _avsDirectory;
        MAX_TOTAL_CHARGE = _maxTotalCharge;
        MIN_BALANCE_THRESHOLD = _minBalanceThreshold;
        MIN_PROOFS_DURATION = _minProofsDuration;
        verifier = _verifier;
        imageId = _imageId;
    }
}
