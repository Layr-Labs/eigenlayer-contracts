// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/contracts/interfaces/IDelegationManager.sol";

interface IDelegationManagerEvents {
    // @notice Emitted when a new operator registers in EigenLayer and provides their OperatorDetails.
    event OperatorRegistered(address indexed operator, IDelegationManager.OperatorDetails operatorDetails);

    // @notice Emitted when an operator updates their OperatorDetails to @param newOperatorDetails
    event OperatorDetailsModified(address indexed operator, IDelegationManager.OperatorDetails newOperatorDetails);

    /**
     * @notice Emitted when @param operator indicates that they are updating their MetadataURI string
     * @dev Note that these strings are *never stored in storage* and are instead purely emitted in events for off-chain indexing
     */
    event OperatorMetadataURIUpdated(address indexed operator, string metadataURI);

    /**
     * @notice Emitted when @param avs indicates that they are updating their MetadataURI string
     * @dev Note that these strings are *never stored in storage* and are instead purely emitted in events for off-chain indexing
     */
    event AVSMetadataURIUpdated(address indexed avs, string metadataURI);

    /// @notice Enum representing the status of an operator's registration with an AVS
    enum OperatorAVSRegistrationStatus {
        UNREGISTERED,       // Operator not registered to AVS
        REGISTERED          // Operator registered to AVS
    }

    /// @notice Emitted when an operator's registration status for an AVS is updated
    event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, OperatorAVSRegistrationStatus status);

    /// @notice Emitted whenever an operator's shares are increased for a given strategy
    event OperatorSharesIncreased(address indexed operator, address staker, IStrategy strategy, uint256 shares);

    /// @notice Emitted whenever an operator's shares are decreased for a given strategy
    event OperatorSharesDecreased(address indexed operator, address staker, IStrategy strategy, uint256 shares);

    // @notice Emitted when @param staker delegates to @param operator.
    event StakerDelegated(address indexed staker, address indexed operator);

    // @notice Emitted when @param staker undelegates from @param operator.
    event StakerUndelegated(address indexed staker, address indexed operator);

    /// @notice Emitted when @param staker is undelegated via a call not originating from the staker themself
    event StakerForceUndelegated(address indexed staker, address indexed operator);

    /**
     * @notice Emitted when a new withdrawal is queued.
     * @param withdrawalRoot Is the hash of the `withdrawal`.
     * @param withdrawal Is the withdrawal itself.
     */
    event WithdrawalQueued(bytes32 withdrawalRoot, IDelegationManager.Withdrawal withdrawal);

    /// @notice Emitted when a queued withdrawal is completed
    event WithdrawalCompleted(bytes32 withdrawalRoot);

    /// @notice Emitted when the `strategyWithdrawalDelayBlocks` variable is modified from `previousValue` to `newValue`.
    event StrategyWithdrawalDelayBlocksSet(IStrategy strategy, uint256 previousValue, uint256 newValue);
}
