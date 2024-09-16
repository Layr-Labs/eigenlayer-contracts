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

    struct OperatorLeaf {
        uint256 delegatedStake;
        uint256 slashableStake;
        bytes32 extraData;
    }

    struct DepositInfo {
        uint96 balance;
        uint32 lastUpdatedTimestamp;
        uint96 totalChargePerOperatorSetLastPaid;
        uint96 totalChargePerStrategyLastPaid;
    }

    struct StakeRootSubmission {
        bytes32 stakeRoot;
        uint32 calculationTimestamp;
        bool confirmed; // whether the submission was posted without proof by governance
    }

    event SnarkProofVerified(bytes journal, bytes seal);
    event VerifierSet(address newVerifier);
    event ImageIdSet(bytes32 newImageId);

    function MIN_BALANCE_THRESHOLD() external view returns (uint256);

    function delegationManager() external view returns (IDelegationManager);
    function avsDirectory() external view returns (IAVSDirectory);
    function verifier() external view returns (address);
    function imageId() external view returns (bytes32);

    /// @notice the number of operator sets in the StakeTree
    function getNumOperatorSets() external view returns (uint256);

    /// @notice the interval at which proofs can be posted, to not overcharge the operatorSets
    function proofIntervalSeconds() external view returns (uint32);

    /**
     * @notice returns the stake root submission at the given index
     * @param index the index of the stake root submission
     */
    function getStakeRootSubmission(uint32 index) external view returns (StakeRootSubmission memory);

    /**
     * @notice returns the delegated and slashable stakes for an operator in an operatorSet
     * @param operatorSet the operatorSet to get the stakes for
     * @param operator the operator to get the stakes for
     * @return delegatedStake the delegated stake for the operator
     * @return slashableStake the slashable stake for the operator
     */
    function getStakes(IAVSDirectory.OperatorSet calldata operatorSet, address operator)
        external
        view
        returns (uint256 delegatedStake, uint256 slashableStake);

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
     * @param operatorLeaves the operator leaves in the operatorSet
     */
    function getOperatorSetRoot(
        IAVSDirectory.OperatorSet calldata operatorSet,
        address[] calldata operators,
        OperatorLeaf[] calldata operatorLeaves
    ) external view returns (bytes32);

    /**
     * @notice Returns the operatorSet leaves for the operatorSet at a given index for a certain set of operators
     * @param operatorSetIndex the index of the operatorSet within the SRC's operatorSets list to calculate the operator set leaves for
     * @param startOperatorIndex the index of the first operator to get the leaves for
     * @param numOperators the number of operators to get the leaves for
     * @return the operatorSet leaves
     */
    function getOperatorSetLeaves(
        uint256 operatorSetIndex,
        uint256 startOperatorIndex,
        uint256 numOperators
    ) external view returns (OperatorLeaf[] memory);

    /**
     * @notice deposits funds for an operator set
     * @param operatorSet the operator set to deposit for
     * @dev must be called before adding strategies and multipliers to the operator set
     * @dev the operator set must have a minimum balance of 2 * MIN_DEPOSIT_BALANCE to disallow joining with minimal cost after removal
     * @dev permissionless to deposit
     */
    function deposit(IAVSDirectory.OperatorSet calldata operatorSet) external payable;

    /**
     * @notice called by an AVS to add strategies and multipliers or modify multipliers used to determine stakes for stake roots
     * @param operatorSetId the id of the operatorSet to set the strategies and multipliers for
     * @param strategiesAndMultipliers the strategies and multipliers to set for the operatorSet
     * @dev msg.sender is used as the AVS in determining the operatorSet
     */
    function addOrModifyStrategiesAndMultipliers(
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
     * @notice called by an AVS to add extraData to an operatorSet to be added to right subtree of the operatorSetTree
     * @param operatorSetId the id of the operatorSet to set the extraData for
     * @param extraData the extraData
     */
    function setOperatorSetExtraData(uint32 operatorSetId, bytes32 extraData) external;

    /**
     * @notice called by an AVS to add extraData to an operator for an operatorSet to be added to leaves of the tree
     * @param operatorSetId the id of the operatorSet to set the extraData for
     * @param operator the operator to set the extraData for
     * @param extraData the extraData
     */
    function setOperatorExtraData(uint32 operatorSetId, address operator, bytes32 extraData) external;

    /**
     * @notice called by watchers to update the deposit balance infos for operatorSets that have
     * fallen below the minimum balance, in order to remove them from the stakeTree
     * @param operatorSetsToRemove the operatorSets to update the deposit balance infos for
     * @dev sends the caller the leftover after charging if the balance is below the minimum
     */
    function removeOperatorSetsFromStakeTree(IAVSDirectory.OperatorSet[] calldata operatorSetsToRemove) external;

    /**
     * @notice called by the claimer to claim a stake root
     * @param _calculationTimestamp the timestamp of the state the stakeRoot was calculated against
     * @param _stakeRoot the stakeRoot at calculationTimestamp
     * @param _chargeRecipient the address to send the charge to when processed
     * @param _indexChargePerProof the index of the charge per proof to use
     * @param _proof todo
     * @dev permissionless to call
     */
    function verifyStakeRoot(
        uint256 _calculationTimestamp,
        bytes32 _stakeRoot,
        address _chargeRecipient,
        uint256 _indexChargePerProof,
        Proof calldata _proof
    ) external;

    /**
     * @notice called by governance
     * @param calculationTimestamp the timestamp of the state the stakeRoot was calculated against
     * @param stakeRoot the stakeRoot at calculationTimestamp
     * @dev only callable by governance when a root has not been posted in forcePostWindow
     */
    function confirmStakeRoot(uint32 calculationTimestamp, bytes32 stakeRoot) external;

    /**
     * @param numStrategies the number of strategies the operatorSet has
     * @return the minimum deposit balance required for the operatorSet, which is just enough to pay for a certain number of proofs
     * @dev this is enforced upon deposits and additions or modifications of strategies and multipliers
     */
    function minDepositBalance(uint256 numStrategies) external view returns (uint256);

    /**
     * @param operatorSet the operatorSet to check withdrawability for
     * @return whether or not the operatorSet can withdraw any of their deposit balance
     */
    function canWithdrawDepositBalance(IAVSDirectory.OperatorSet memory operatorSet) external view returns (bool);

    /**
     * @notice get the deposit balance for the operator set
     * @param operatorSet the operator set to get the deposit balance for
     * @return balance the deposit balance for the operator set
     */
    function getDepositBalance(IAVSDirectory.OperatorSet memory operatorSet)
        external
        view
        returns (uint256 balance);
}
