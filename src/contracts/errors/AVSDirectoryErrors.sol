// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

abstract contract AVSDirectoryErrors {
    error AVSDirectory_createOperatorSets_OperatorSetAlreadyExists();

    error AVSDirectory_becomeOperatorSetAVS_CallerAlreadyOperatorSetAVS();

    error AVSDirectory_migrateOperatorsToOperatorSets_CallerNotOperatorSetAVS();
    error AVSDirectory_migrateOperatorsToOperatorSets_OperatorAlreadyNonLegacyOperator();

    error AVSDirectory_registerOperatorToOperatorSets_OperatorSignatureExpired();
    error AVSDirectory_registerOperatorToOperatorSets_OperatorNotRegistered();
    error AVSDirectory_registerOperatorToOperatorSets_CallerNotOperatorSetAVS();
    error AVSDirectory_registerOperatorToOperatorSets_OperatorAlreadySpentSalt();

    error AVSDirectory_forceDeregisterFromOperatorSets_CallerMustBeOperatorIfSignatureNotProvided();
    error AVSDirectory_forceDeregisterFromOperatorSets_OperatorSignatureExpired();
    error AVSDirectory_forceDeregisterFromOperatorSets_OperatorAlreadySpentSalt();

    error AVSDirectory_modifyAllocations_OperatorNotRegistered();
    error AVSDirectory_modifyAllocations_OperatorMustInitializeAllocationDelay();
    error AVSDirectory_modifyAllocations_OperatorTotalExpectedMagnitudeMismatch();

    error AVSDirectory_updateFreeMagnitude_OperatorNotRegistered();
    error AVSDirectory_updateFreeMagnitude_ArrayLengthMismatch();

    error AVSDirectory_slashOperator_OperatorNotSlashableForOperatorSet();

    error AVSDirectory_registerOperatorToAVS_OperatorSignatureExpired();
    error AVSDirectory_registerOperatorToAVS_CallerIsOperatorSetAVS();
    error AVSDirectory_registerOperatorToAVS_OperatorAlreadyRegistered();
    error AVSDirectory_registerOperatorToAVS_OperatorAlreadySpentSalt();
    error AVSDirectory_registerOperatorToAVS_OperatorNotRegistered();

    error AVSDirectory_deregisterOperatorFromAVS_OperatorNotRegistered();
    error AVSDirectory_deregisterOperatorFromAVS_CallerIsOperatorSetAVS();

    error AVSDirectory_registerToOperatorSets_InvalidOperatorSet();
    error AVSDirectory_registerToOperatorSets_OperatorAlreadyRegistered();

    error AVSDirectory_deregisterFromOperatorSets_OperatorNotRegistered();

    error AVSDirectory_modifyAllocations_OperatorSetsAndMagnitudesLengthMismatch();
    error AVSDirectory_modifyAllocations_InvalidOperatorSet();
    error AVSDirectory_modifyAllocations_OperatorSetsNotInAscendingOrder();
    error AVSDirectory_modifyAllocations_CannotSetMagnitudeWithPendingAllocationOrDeallocation();
    error AVSDirectory_modifyAllocations_InsuffientFreeAllocatableMagnitude();

    error AVSDirectory_verifyOperatorSignature_OperatorSignatureExpired();
    error AVSDirectory_verifyOperatorSignature_OperatorAlreadySpentSalt();
}
