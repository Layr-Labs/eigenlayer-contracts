// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

abstract contract DelegationManagerErrors {
    error DelegationManager_CallerNotStrategyManagerOrEigenPodManager();

    error DelegationManager_registerAsOperator_CallerAlreadyDelegated();

    error DelegationManager_modifyOperatorDetails_OperatorNotRegistered();

    error DelegationManager_updateOperatorMetadataURI_OperatorNotRegistered();

    error DelegationManager_delegateTo_StakerAlreadyDelegated();
    error DelegationManager_delegateTo_OperatorNotRegistered();

    error DelegationManager_delegateToBySignature_StakerSignatureExpired();
    error DelegationManager_delegateToBySignature_StakerAlreadyDelegated();
    error DelegationManager_delegateToBySignature_OperatorNotRegistered();

    error DelegationManager_undelegate_StakerNotDelegated();
    error DelegationManager_undelegate_OperatorCannotUndelegate();
    error DelegationManager_undelegate_CannotUndelegateToZeroAddress();
    error DelegationManager_undelegate_CallerCannotUndelegateStaker();

    error DelegationManager_queueWithdrawal_InputLengthMismatch();
    error DelegationManager_queueWithdrawal_CallerMustBeStaker();

    error DelegationManager_setOperatorDetails_StakerOptOutWindowBlocksExeedsMax();
    error DelegationManager_setOperatorDetails_StakerOptOutWindowBlocksCannotBeDecreased();

    error DelegationManager_initializeAllocationDelay_OperatorNotRegistered();
    error DelegationManager_initializeAllocationDelay_AllocationDelayAlreadySet();

    error DelegationManager_delegate_ApproverSignatureExpired();
    error DelegationManager_delegate_ApproverAlreadySpentSalt();

    error DelegationManager_completeQueuedWithdrawal_ActionNotQueued();
    error DelegationManager_completeQueuedWithdrawal_MinWithdrawalDelayMustElapse();
    error DelegationManager_completeQueuedWithdrawal_CallerNotWithdrawer();
    error DelegationManager_completeQueuedWithdrawal_InputLengthsMismatch();

    error DelegationManager_completeReceiveAsTokens_WithdrawalDelayMustElapseForStrategy();

    error DelegationManager_completeReceiveAsShares_WithdrawalDelayMustElapseForStrategy();

    error DelegationManager_removeSharesAndQueueWithdrawal_StakerCannotBeZeroAddress();
    error DelegationManager_removeSharesAndQueueWithdrawal_StrategiesNotProvided();
    error DelegationManager_removeSharesAndQueueWithdrawal_StakerForbidsThirdPartyTransfers();

    error DelegationManager_setMinWithdrawalDelayBlocks_MinWithdrawalDelayBlocksExceedsMax();

    error DelegationManager_setStrategyWithdrawalDelayBlocks_InputLengthsMismatch();
    error DelegationManager_setStrategyWithdrawalDelayBlocks_WithdrawalDelayBlocksExeedsMax();

    error DelegationManager_setStrategyWithdrawalDelay_InputLengthsMismatch();
    error DelegationManager_setStrategyWithdrawalDelay_WithdrawalDelayExceedsMax();
}
