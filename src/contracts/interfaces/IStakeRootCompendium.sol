// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IAVSDirectory.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IStrategy.sol";

interface IStakeRootCompendium {
    struct StrategyAndMultiplier {
        IStrategy strategy;
        uint96 multiplier;
    }

    // TODO
    struct Proof {
        uint32 x;
    }

    struct DepositBalanceInfo {
        uint32 lastUpdatedAt;
        uint256 balance;
    }

    struct StakeRootSubmission {
        bytes32 stakeRoot;
        address chargeRecipient; // the address to send the charge to
        uint32 calculationTimestamp; // the timestamp the was generated against
        uint32 blacklistableBefore; // the timestamp the proof submission was submitted to the contract
        bool blacklisted; // whether the submission has been blacklisted by governance
        bool forcePosted; // whether the submission was posted without proof by governance
    }

    event SnarkProofVerified(bytes journal, bytes seal);
    event VerifierChanged(address oldVerifier, address newVerifier);
    event ImageIdChanged(bytes32 oldImageId, bytes32 newImageId);

    function MAX_OPERATOR_SET_SIZE() external view returns (uint32);
    function MAX_NUM_OPERATOR_SETS() external view returns (uint32);
    function MAX_NUM_STRATEGIES() external view returns (uint32);

    /// @notice the minimum balance that must be maintained for an operatorSet
    function MIN_DEPOSIT_BALANCE() external view returns (uint256);

    function delegationManager() external view returns (IDelegationManager);
    function avsDirectory() external view returns (IAVSDirectory);
    function verifier() external view returns (address);
    function imageId() external view returns (bytes32);

    /// @notice charge per strategy per proof
    function charge() external view returns (uint256);

    /// @notice the interval at which proofs can be posted, to not overcharge the operatorSets
    function proofInterval() external view returns (uint32);

    /// @notice the stake root submissions that have been posted
    function getStakeRootSubmission(uint32 index) external view returns (StakeRootSubmission memory);

    /// @notice return the index of an operatorSet at a certain timestamp
    function getOperatorSetIndexAtTimestamp(
        IAVSDirectory.OperatorSet calldata operatorSet,
        uint32 timestamp
    ) external view returns (uint32);

    /**
     * @notice called offchain with the operatorSet roots ordered by the operatorSet index at the timestamp to calculate the stake root
     * @param operatorSetsInStakeTree the operatorSets that each of the operatorSetRoots correspond to. must be the same as operatorSets storage var at the time of call
     * @param operatorSetRoots the ordered operatorSet roots (not verified)
     * @dev operatorSetsInStakeTree must be the same as operatorSets storage var at the time of call
     * @dev operatorSetRoots must be ordered by the operatorSet index at the time of call
     */
    function getStakeRoot(
        IAVSDirectory.OperatorSet[] calldata operatorSetsInStakeTree,
        bytes32[] calldata operatorSetRoots
    ) external view returns (bytes32);

    /**
     * @notice Returns the operatorSet root for the given operatorSet with the given witnesses
     * @param operatorSet the operatorSet to calculate the operator set root for
     * @param operators the operators in the operatorSet
     * @return the operatorSet root
     * @dev the operators are verified to be distinct and the length must be the total
     * number of operators in the operatorSet at the time of calling
     */
    function getOperatorSetRoot(
        IAVSDirectory.OperatorSet calldata operatorSet,
        address[] calldata operators
    ) external view returns (bytes32);

    /**
     * @notice deposits funds for an operator set
     * @param operatorSet the operator set to deposit for
     * @dev must be called before adding strategies and multipliers to the operator set
     * @dev the operator set must have a minimum balance of 2 * MIN_DEPOSIT_BALANCE to disallow joining with minimal cost after removal
     * @dev permissionless to deposit
     */
    function depositForOperatorSet(IAVSDirectory.OperatorSet calldata operatorSet) external payable;

    /**
     * @notice called by an AVS to set their strategies and multipliers used to determine stakes for stake roots
     * @param operatorSetId the id of the operatorSet to set the strategies and multipliers for
     * @param strategiesAndMultipliers the strategies and multipliers to set for the operatorSet
     * @dev msg.sender is used as the AVS in determining the operatorSet
     */
    function addStrategiesAndMultipliers(
        uint32 operatorSetId,
        StrategyAndMultiplier[] calldata strategiesAndMultipliers
    ) external;

    /**
     * @notice called by an AVS to remove their strategies and multipliers used to determine stakes for stake roots
     * @param operatorSetId the id of the operatorSet to remove the strategies and multipliers for
     * @param strategies the strategies to remove for the operatorSet
     * @dev msg.sender is used as the AVS in determining the operatorSet
     */
    function removeStrategiesAndMultipliers(uint32 operatorSetId, IStrategy[] calldata strategies) external;

    /**
     * @notice called by an AVS to add extraData to an operator for an operatorSet to be added to leaves of the tree
     * @param operatorSetId the id of the operatorSet to set the extraData for
     * @param operator the operator to set the extraData for
     * @param timestamp the time at which the extraData takes effect, in case delayed updates to extradata are required
     * @param extraData the extraData
     */
    function setExtraData(uint32 operatorSetId, address operator, uint32 timestamp, bytes32 extraData) external;

    /**
     * @notice called by watchers to update the deposit balance infos for operatorSets, usually those that have
     * fallen below the minimum balance, in order to remove them from the stakeTree
     * @param operatorSetsToUpdate the operatorSets to update the deposit balance infos for
     * @dev sends the caller the leftover after charging if the balance is below the minimum
     */
    function updateDepositBalanceInfos(IAVSDirectory.OperatorSet[] calldata operatorSetsToUpdate) external;

    /**
     * @notice Process charges for the next numToCharge stakeRootSubmissions that have not been redeemed
     * @param numToCharge the number of charges to process
     */
    function processCharges(uint256 numToCharge) external;

    /**
     * @notice called by the claimer to claim a stake root
     * @param calculationTimestamp the timestamp of the state the stakeRoot was calculated against
     * @param stakeRoot the stakeRoot at calculationTimestamp
     * @param chargeRecipient the address to send the charge to when processed
     * @param proof todo
     * @dev permissionless to call
     */
    function verifyStakeRoot(
        uint32 calculationTimestamp,
        bytes32 stakeRoot,
        address chargeRecipient,
        Proof calldata proof
    ) external;

    /**
     * @notice called by governance to blacklist a stakeRoot in case it's incorrect
     * @param submissionIndex the index of the stakeRoot submission to blacklist
     * @dev called in case there's a bug in the verifier stack or program to allow an incorrect stakeRoot to be proven
     * @dev only callable by governance
     * @dev must be called within blacklistWindow of the stakeRoot being posted
     */
    function blacklistStakeRoot(uint32 submissionIndex) external;

    /**
     * @notice called by governance
     * @param calculationTimestamp the timestamp of the state the stakeRoot was calculated against
     * @param stakeRoot the stakeRoot at calculationTimestamp
     * @dev only callable by governance when a root has not been posted in forcePostWindow
     */
    function forcePostStakeRoot(uint32 calculationTimestamp, bytes32 stakeRoot) external;

    /**
     * @notice sets the verifier contract that will be used to verify snark proofs
     * @param _verifier the address of the verifier contract
     * @dev only callable by the owner
     */
    function setVerifier(address _verifier) external;

    /**
     * @notice sets/changes the id of the program being verified when roots are posted
     * @param _imageId the new imageId to set
     * @dev only callable by the owner
     */
    function setImageId(bytes32 _imageId) external;

    /**
     * @notice get the deposit balance for the operator set
     * @param operatorSet the operator set to get the deposit balance for
     * @return balance the deposit balance for the operator set
     * @return penalty the penalty to be received by calling updateDepositBalanceInfos if the operator set has fallen below the minimum deposit balance
     */
    function getDepositBalance(IAVSDirectory.OperatorSet memory operatorSet)
        external
        view
        returns (uint256 balance, uint256 penalty);
}
