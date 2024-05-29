// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

/**
 * @title DelegationManagerErrors
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice Library for DelegationManager Error messages
 */
library DelegationManagerErrors {
    /// onlyStrategyManagerOrEigenPodManager()
    string internal constant ONLY_SM_OR_EPM = "DelegationManager: onlyStrategyManagerOrEigenPodManager";

    /// registerAsOperator()
    string internal constant ALREADY_DELEGATED = "DelegationManager: caller is already actively delegated";

    /// modifyOperatorDetails()
    /// updateOperatorMetadataURI()
    string internal constant NOT_AN_OPERATOR = "DelegationManager: caller must be an operator";
    string internal constant NOT_OWNER = "DelegationManager.updateOperatorMetadataURI: caller must be an operator";
    string internal constant STAKER_NOT_DELEGATED = "DelegationManager: staker must be delegated to undelegate";

    /// delegateToBySignature()
    string internal constant EXPIRED_SIG = "DelegationManager.delegateToBySignature: staker signature expired";

    /// undelegate()
    string internal constant STAKER_CANT_BE_OPERATOR = "DelegationManager.undelegate: staker can't be an operator";
    string internal constant ZERO_ADDRESS = "DelegationManager: address must be non-zero";
    string internal constant UNAUTHORIZED_UNDELEGATE = "DelegationManager.undelegate: caller cannot undelegate staker";

    /// queueWithdrawals()
    string internal constant INVALID_INPUT_ARRAY_LENGTHS = "DelegationManager.queueWithdrawal: input length mismatch";
    string internal constant UNAUTHORIZED_WITHDRAWER = "DelegationManager.queueWithdrawal: withdrawer must be staker";

    /// _setOperatorDetails()
    string internal constant GREATER_THAN_MAX_OPT_OUT_WINDOW_BLOCKS = "DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be > MAX_STAKER_OPT_OUT_WINDOW_BLOCKS";
    string internal constant DECREASING_OPT_OUT_WINDOW_BLOCKS = "DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be decreased";

    /// _delegate()
    string internal constant DELEGATE_APPROVE_SIG_EXPIRE = "DelegationManager._delegate: approver signature expired";
    string internal constant DELEGATE_APPROVE_SALT_SPENT = "DelegationManager._delegate: approverSalt already spent";

    /// _completeQueuedWithdrawal()
    string internal constant WITHDRAWAL_NOT_PENDING = "DelegationManager._completeQueuedWithdrawal: action is not in queue";
    string internal constant WITHDRAWAL_PERIOD_NOT_PASSED = "DelegationManager._completeQueuedWithdrawal: minWithdrawalDelayBlocks period has not yet passed";
    string internal constant STRATEGY_WITHDRAWAL_PERIOD_NOT_PASSED = "DelegationManager._completeQueuedWithdrawal: withdrawalDelayBlocks period has not yet passed for this strategy";
    string internal constant UNAUTHORIZED_MSG_SENDER = "DelegationManager._completeQueuedWithdrawal: only withdrawer can complete action";
    string internal constant WITHDRAWAL_INVALID_INPUT_ARRAY_LENGTHS = "DelegationManager._completeQueuedWithdrawal: input length mismatch";

    /// _removeSharesAndQueueWithdrawal()
    string internal constant REMOVE_SHARES_INVALID_INPUT_ARRAY_LENGTHS = "DelegationManager._removeSharesAndQueueWithdrawal: strategies cannot be empty";
    string internal constant REMOVE_SHARES_THIRD_PARTY_TRANSFERS_FORBIDDEN = "DelegationManager._removeSharesAndQueueWithdrawal: withdrawer must be same address as staker if thirdPartyTransfersForbidden are set";

    /// _setMinWithdrawalDelayBlocks()
    string internal constant MAX_WITHDRAWAL_DELAY_BLOCKS_EXCEEDED = "DelegationManager._setMinWithdrawalDelayBlocks: _minWithdrawalDelayBlocks cannot be > MAX_WITHDRAWAL_DELAY_BLOCKS";

    /// _setWithdrawalDelayBlocks()
    string internal constant SET_STRATEGY_DELAY_INVALID_ARRAY_LENGTHS = "DelegationManager._setStrategyWithdrawalDelayBlocks: input length mismatch";
    string internal constant MAX_STRATEGY_WITHDRAWAL_DELAY_BLOCKS_EXCEEDED = "DelegationManager._setStrategyWithdrawalDelayBlocks: _withdrawalDelayBlocks cannot be > MAX_WITHDRAWAL_DELAY_BLOCKS";
}