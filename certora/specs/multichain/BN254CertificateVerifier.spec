import "../libraries/BN254-nondet.spec";
using OperatorSetHelper as OperatorSetHelper;

use builtin rule sanity filtered { f -> f.contract == currentContract }

methods{
    function OperatorSetHelper.getOperatorSetKey(BN254CertificateVerifier.OperatorSet os) external returns (bytes32) envfree;
}

/*
 * msg.sender != address(operatorTableUpdater) => revert
 * referenceTimestamp == 0 => revert
 *
 * What it means: Only the designated operatorTableUpdater address can call this function, all other callers must be rejected
 *
 * Why it should hold: The contract has a specific modifier 'onlyTableUpdater' and the function is critical for updating operator information. Unauthorized access could corrupt the operator registry
 *
 * Possible consequences: Unauthorized operator table manipulation, state corruption, denial of service by malicious actors updating operator information
 */
rule updateOperatorTable_revert_conditions(env e) {
    BN254CertificateVerifier.OperatorSet operatorSet;
    uint32 referenceTimestamp;
    IOperatorTableCalculatorTypes.BN254OperatorSetInfo operatorSetInfo;
    ICrossChainRegistryTypes.OperatorSetConfig operatorSetConfig;

    bool zeroRefTS = referenceTimestamp == 0;
    // call function under test
    updateOperatorTable@withrevert(e, operatorSet, referenceTimestamp, operatorSetInfo, operatorSetConfig);
    bool updateOperatorTable_reverted = lastReverted;

    // verify integrity
    assert ((e.msg.sender != currentContract.operatorTableUpdater || zeroRefTS) => updateOperatorTable_reverted), "msg.sender != address(operatorTableUpdater) => revert";
}




rule updateOperatorTable_updates_latest_for_set_and_frames(env e) {
    BN254CertificateVerifier.OperatorSet os;
    uint32 ts;
    IOperatorTableCalculatorTypes.BN254OperatorSetInfo info;
    ICrossChainRegistryTypes.OperatorSetConfig cfg;

    // keys
    bytes32 k  = OperatorSetHelper.getOperatorSetKey(os);
    BN254CertificateVerifier.OperatorSet otherOs;
    bytes32 ok = OperatorSetHelper.getOperatorSetKey(otherOs);

    // snapshots
    uint32 latest_before   = currentContract._latestReferenceTimestamps[k];
    uint32 other_before    = currentContract._latestReferenceTimestamps[ok];

    // require minimal success preconditions you know the code enforces:
    require e.msg.sender == currentContract.operatorTableUpdater;
    require ts != 0;
    // (If the implementation also checks ts <= block.timestamp, keep it below instead.)

    updateOperatorTable(e, os, ts, info, cfg);

    // effect on this set
    assert currentContract._latestReferenceTimestamps[k] == ts;
    assert currentContract._latestReferenceTimestamps[k] >  latest_before;

    // frame: unrelated sets unchanged
    assert (ok == k) || (currentContract._latestReferenceTimestamps[ok] == other_before);
}




/*
 * !operatorTableUpdater.isRootValidByTimestamp(cert.referenceTimestamp) => revert
 *
 * What it means: The function must revert if the operator table updater indicates that the root corresponding to the reference timestamp is disabled or invalid
 *
 * Why it should hold: The contract explicitly checks this condition in _validateCertificateTimestamp() with 'require(operatorTableUpdater.isRootValidByTimestamp(referenceTimestamp), RootDisabled())'. This ensures certificates are only verified against valid, non-disabled operator sets.
 *
 * Possible consequences: Accepting certificates based on outdated or compromised operator sets, potential for replay attacks using disabled roots
 */
rule verifyCertificate_disabled_root(env e) {
    BN254CertificateVerifier.OperatorSet operatorSet;
    IBN254CertificateVerifierTypes.BN254Certificate cert;
    uint256[] signedStakes;

    // call function under test
    verifyCertificate@withrevert(e, operatorSet, cert);
    bool verifyCertificate_reverted = lastReverted;

    // verify integrity
    assert (!currentContract.operatorTableUpdater.isRootValidByTimestamp(e, cert.referenceTimestamp) => verifyCertificate_reverted), "!operatorTableUpdater.isRootValidByTimestamp(cert.referenceTimestamp) => revert";
}

/*
 * !operatorTableUpdater.isRootValidByTimestamp(cert.referenceTimestamp) => revert
 * * totalStakeNominalThresholds.length > 100 => revert
 * What it means: The function must revert when the operator table updater indicates that the root for the given reference timestamp is disabled or invalid
 *
 * Why it should hold: This is a critical security check that ensures only valid, non-revoked operator sets can be used for verification. Disabled roots indicate compromised or outdated operator information
 *
 * Possible consequences: Security bypass allowing verification against compromised or revoked operator sets, potential fund loss or unauthorized operations
 */
rule verifyCertificateNominal_revert_conditions(env e) {
    BN254CertificateVerifier.OperatorSet operatorSet;
    IBN254CertificateVerifierTypes.BN254Certificate cert;
    uint256[] totalStakeNominalThresholds;
    bool result;

    bool disabledRoot = !currentContract.operatorTableUpdater.isRootValidByTimestamp(e, cert.referenceTimestamp);
    bool maxThresholdPassed = totalStakeNominalThresholds.length > 100;

    // call function under test
    verifyCertificateNominal@withrevert(e, operatorSet, cert, totalStakeNominalThresholds);
    bool verifyCertificateNominal_reverted = lastReverted;

    // verify integrity
    assert ((disabledRoot || maxThresholdPassed) => verifyCertificateNominal_reverted), "!operatorTableUpdater.isRootValidByTimestamp(cert.referenceTimestamp) => revert";
}
