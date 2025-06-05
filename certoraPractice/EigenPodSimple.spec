
using EigenPodManager as eigenPodManager;

methods {
    function getParentBlockRoot(uint64) internal returns (bytes32) => NONDET;

    function BeaconChainProofs.verifyStateRoot(bytes32 beaconBlockRoot, BeaconChainProofs.StateRootProof calldata proof) internal => NONDET;

    function BeaconChainProofs.verifyValidatorFields(
        BeaconChainProofs.ProofVersion proofVersion,
        bytes32 beaconStateRoot,
        bytes32[] calldata validatorFields,
        bytes calldata validatorFieldsProof,
        uint40 validatorIndex
    ) internal => NONDET;

    function BeaconChainProofs.verifyBalanceContainer(
        BeaconChainProofs.ProofVersion proofVersion,
        bytes32 beaconBlockRoot,
        BeaconChainProofs.BalanceContainerProof calldata proof
    ) internal => NONDET;

    function BeaconChainProofs.verifyValidatorBalance(
        bytes32 balanceContainerRoot,
        uint40 validatorIndex,
        BeaconChainProofs.BalanceProof calldata proof
    ) internal returns (uint64) => NONDET;

    function BeaconChainProofs.isValidatorSlashed(
        bytes32[] memory validatorFields
    ) internal returns (bool) => NONDET;

    function _.deposit(
        bytes pubkey,
        bytes withdrawal_credentials,
        bytes signature,
        bytes32 deposit_data_root) external => doNothing() expect void;
    
    // envfree functions
    function podOwner() external returns (address) envfree;
    function validatorStatus(bytes32 pubkeyHash) external returns (IEigenPodTypes.VALIDATOR_STATUS) envfree;
    function proofSubmitter() external returns (address) envfree;
}

function doNothing() {}

rule eigenPodSanity(env e, method f) {
    calldataarg args;
    f(e, args);
    satisfy true;
}

rule initializePodOwner(env e) {
    address podOwner;
    initialize(e, podOwner);
    assert podOwner() == podOwner;
}

rule initializeSetsPodOwnerWithRevert(env e) {
    address podOwner;
    initialize@withrevert(e, podOwner);
    bool initializeReverted = lastReverted;
    assert (podOwner() == podOwner && podOwner() != 0) || initializeReverted;
}

rule initializeCannotBeCalledTwice() {
    env e;
    address podOwner;
    initialize(e, podOwner);
    env e2;
    address someoneElse;
    initialize@withrevert(e2, someoneElse);
    bool initializeReverted = lastReverted;
    assert lastReverted;
    assert podOwner() == podOwner;
}

rule onlyPodOwnerCanCall(env e, method f) filtered {
    f -> f.selector == sig:recoverTokens(address[], uint256[], address).selector
    || f.selector == sig:setProofSubmitter(address).selector
} {
        calldataarg args;
        f(e, args);
        assert e.msg.sender == podOwner();
}

ghost bool wasInitialized {
    init_state axiom wasInitialized == false;
}

hook Sstore currentContract.podOwner address newOwner (address oldOwner) {
    if (oldOwner == 0 && newOwner != 0) {
        wasInitialized = true;
    }
}

invariant podOwnerCannotBeZeroA()
    wasInitialized => podOwner() != 0;

// Exercise 1
definition GWEI_TO_WEI() returns uint32 = 1000000000; // 1e9

rule cannotWithdrawMoreThanRestaked(env e) {
    // Transform initialRestaked to gwei
    uint256 initialRestakedGwei = require_uint256(currentContract.restakedExecutionLayerGwei * GWEI_TO_WEI());

    address recipient;
    uint256 amountToWithdrawWei;

    // Require amount to withdraw is too large
    require require_uint64(amountToWithdrawWei / GWEI_TO_WEI()) > currentContract.restakedExecutionLayerGwei; 

    withdrawRestakedBeaconChainETH@withrevert(e, recipient, amountToWithdrawWei);

    assert lastReverted;
}

/// Exercise 2
invariant inactiveValidatorsHaveEmptyInfo(bytes32 pubkeyHash)
    (validatorStatus(pubkeyHash) == IEigenPodTypes.VALIDATOR_STATUS.INACTIVE) => (
        currentContract._validatorPubkeyHashToInfo[pubkeyHash].validatorIndex == 0
        && currentContract._validatorPubkeyHashToInfo[pubkeyHash].restakedBalanceGwei == 0
        && currentContract._validatorPubkeyHashToInfo[pubkeyHash].lastCheckpointedAt == 0
    );


/// Exercise 3
rule validatorIndexSetOnlyOnce(bytes32 pubkeyHash) {
    uint64 validatorIndexBefore = currentContract._validatorPubkeyHashToInfo[pubkeyHash].validatorIndex;
    // perform arbitrary function call
    method f;
    env e;
    calldataarg args;
    f(e,args);
    
    uint64 validatorIndexAfter = currentContract._validatorPubkeyHashToInfo[pubkeyHash].validatorIndex;
    assert(validatorIndexBefore == 0 || validatorIndexAfter == validatorIndexBefore,
        "validator index modified from nonzero value");
}


/// Exercise 4
definition validatorStatusTransitionAllowed(IEigenPodTypes.VALIDATOR_STATUS statusBefore, IEigenPodTypes.VALIDATOR_STATUS statusAfter) returns bool =
    (statusBefore == IEigenPodTypes.VALIDATOR_STATUS.INACTIVE && statusAfter == IEigenPodTypes.VALIDATOR_STATUS.ACTIVE)
    || (statusBefore == IEigenPodTypes.VALIDATOR_STATUS.ACTIVE && statusAfter == IEigenPodTypes.VALIDATOR_STATUS.WITHDRAWN);

rule onlyValidStatusStateTransitions(env e, method f, calldataarg args, bytes32 pubkeyHash) {
    IEigenPodTypes.VALIDATOR_STATUS statusBefore = currentContract._validatorPubkeyHashToInfo[pubkeyHash].status;

    f(e,args);

    IEigenPodTypes.VALIDATOR_STATUS statusAfter = currentContract._validatorPubkeyHashToInfo[pubkeyHash].status;

    assert(
        statusBefore == statusAfter ||
        (statusBefore == IEigenPodTypes.VALIDATOR_STATUS.INACTIVE && statusAfter == IEigenPodTypes.VALIDATOR_STATUS.ACTIVE) ||
        (statusBefore == IEigenPodTypes.VALIDATOR_STATUS.ACTIVE && statusAfter == IEigenPodTypes.VALIDATOR_STATUS.WITHDRAWN)
    );
}

// Invariants around proofSubmitter

/**
 * 1. Owner or proof submitter can submit a checkpoint
 * 2. Owner or proof submitter can verify withdrawal credentials
 */

rule onlyPodOwnerOrProofSubmitterCanCall(env e, method f) filtered {
    f -> f.selector == sig:startCheckpoint(bool).selector
    || f.selector == sig:verifyWithdrawalCredentials(uint64, BeaconChainProofs.StateRootProof, uint40[], bytes[], bytes32[][]).selector
} {
    calldataarg args;
    f(e, args);
    assert e.msg.sender == podOwner() || e.msg.sender == proofSubmitter();
}

/// Helper Invariant
rule checkVerifyCheckpointProofs(env e){
    require eigenPodManager.podOwnerDepositShares[podOwner()] % GWEI_TO_WEI() == 0;
    require currentContract.restakedExecutionLayerGwei * GWEI_TO_WEI() <= eigenPodManager.podOwnerDepositShares[podOwner()];

    BeaconChainProofs.BalanceContainerProof balanceContainerProof;
    BeaconChainProofs.BalanceProof[] proofs;
    verifyCheckpointProofs(e, balanceContainerProof, proofs);
    
    assert currentContract._currentCheckpoint.proofsRemaining == 0 => currentContract.restakedExecutionLayerGwei * GWEI_TO_WEI() <= eigenPodManager.podOwnerDepositShares[podOwner()];
}