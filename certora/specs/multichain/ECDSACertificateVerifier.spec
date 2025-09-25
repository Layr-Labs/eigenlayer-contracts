using OperatorSetHelper as OperatorSetHelper;

methods {
    function ECDSACertificateVerifier._parseSignatures(bytes32 signableDigest, bytes memory signatures) internal returns (address[] memory) => cvlParseSignatures(signatures);
    function OperatorSetHelper.getOperatorSetKey(ECDSACertificateVerifier.OperatorSet os) external returns (bytes32) envfree;
}

// Needed because the copy loop in _parseSignatures seems to be heavy for the prover.
function cvlParseSignatures(bytes signatures) returns (address[]) {
    uint256 signatureCount = require_uint256(signatures.length / 65);
    address[] signers;
    require signers.length == signatureCount;
    return signers;
}

use builtin rule sanity filtered { f -> f.contract == currentContract }


rule ecdsa_updateOperatorTable_revert_conditions(env e) {
  ECDSACertificateVerifier.OperatorSet os;
  uint32 ts;
  IOperatorTableCalculatorTypes.ECDSAOperatorInfo[] operatorInfos;
  ICrossChainRegistryTypes.OperatorSetConfig cfg;

  bytes32 k = OperatorSetHelper.getOperatorSetKey(os);
  bool notUpdater = e.msg.sender != currentContract.operatorTableUpdater;
  bool stale      = ts <= currentContract._latestReferenceTimestamps[k];

  updateOperatorTable@withrevert(e, os, ts, operatorInfos, cfg);
  assert ((notUpdater || stale) => lastReverted),
         "onlyTableUpdater && ts must be strictly newer";
}

rule ecdsa_updateOperatorTable_updates_and_frames(env e) {
  ECDSACertificateVerifier.OperatorSet os;
  uint32 ts;
  IOperatorTableCalculatorTypes.ECDSAOperatorInfo[] operatorInfos;
  ICrossChainRegistryTypes.OperatorSetConfig cfg;

  bytes32 k  = OperatorSetHelper.getOperatorSetKey(os);
  ECDSACertificateVerifier.OperatorSet other; 
  bytes32 ok = OperatorSetHelper.getOperatorSetKey(other);

  // snapshots
  uint32 latest_before = currentContract._latestReferenceTimestamps[k];
  uint32 other_before  = currentContract._latestReferenceTimestamps[ok];

  updateOperatorTable(e, os, ts, operatorInfos, cfg);

  // effects
  assert currentContract._latestReferenceTimestamps[k] == ts;
  assert currentContract._latestReferenceTimestamps[k] > latest_before;
  assert currentContract._numOperators[k][ts] == operatorInfos.length;
  assert currentContract._operatorSetOwners[k] == cfg.owner;
  assert currentContract._maxStalenessPeriods[k] == cfg.maxStalenessPeriod;

  // frame: other sets unchanged
  assert (ok == k) || (currentContract._latestReferenceTimestamps[ok] == other_before);
}

rule ecdsa_verifyCertificate_disabled_root_reverts(env e) {
  ECDSACertificateVerifier.OperatorSet os;
  IECDSACertificateVerifierTypes.ECDSACertificate cert;

  bool disabled = !currentContract.operatorTableUpdater.isRootValidByTimestamp(e, cert.referenceTimestamp);

  verifyCertificate@withrevert(e, os, cert);
  assert (disabled => lastReverted), "disabled root => revert";
}

rule ecdsa_verifyCertificate_staleness_window_reverts(env e) {
  ECDSACertificateVerifier.OperatorSet os;
  IECDSACertificateVerifierTypes.ECDSACertificate cert;

  bytes32 k = OperatorSetHelper.getOperatorSetKey(os);
  bool stale = currentContract._maxStalenessPeriods[k] != 0 && e.block.timestamp > cert.referenceTimestamp + currentContract._maxStalenessPeriods[k];

  verifyCertificate@withrevert(e, os, cert);
  assert (stale => lastReverted), "cert beyond staleness window => revert";
}


rule ecdsa_verify_paths_do_not_mutate_state(env e) {
  ECDSACertificateVerifier.OperatorSet os;
  IECDSACertificateVerifierTypes.ECDSACertificate cert;
  uint16[] p; uint256[] n;

  bytes32 k = OperatorSetHelper.getOperatorSetKey(os);

  uint32 latest_before      = currentContract._latestReferenceTimestamps[k];
  uint32 maxStale_before    = currentContract._maxStalenessPeriods[k];
  address owner_before      = currentContract._operatorSetOwners[k];

  verifyCertificate(e, os, cert);
  verifyCertificateProportion(e, os, cert, p);
  verifyCertificateNominal(e, os, cert, n);

  assert currentContract._latestReferenceTimestamps[k] == latest_before;
  assert currentContract._maxStalenessPeriods[k] == maxStale_before;
  assert currentContract._operatorSetOwners[k] == owner_before;
}