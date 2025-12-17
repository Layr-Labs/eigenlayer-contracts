import "../libraries/BN254-nondet.spec";

methods {
}

use builtin rule sanity filtered { f -> f.contract == currentContract  }


// A tiny “frame” helper for maps: if key != touchedKey then unchanged.
function unchangedMappingUint32(uint32 before, uint32 after, bool keysEqual) returns (bool) {
    return keysEqual || (after == before);
}

// A tiny “frame” helper for maps: if key != touchedKey then unchanged.
function unchangedMappingBytes32(bytes32 before, bytes32 after, bool keysEqual) returns (bool) {
    return keysEqual || (after == before);
}

rule confirmGlobalTableRoot_reverts_conditions(env e) {
    IBN254CertificateVerifierTypes.BN254Certificate cert;
    bytes32 root; 
    uint32 ts; 
    uint32 blk;

    bool futureTS = ts > e.block.timestamp;
    bool staleTS = ts <= currentContract._latestReferenceTimestamp; 
    bool zeroTS = ts == 0;
    bool invalidMsgHash = currentContract.getGlobalTableUpdateMessageHash(e, root, ts, blk) != cert.messageHash;
    // PAUSED_GLOBAL_ROOT_UPDATE = 0
    uint256 mask = 1 << 0;
    bool paused = (currentContract._paused & mask) == mask;

    confirmGlobalTableRoot@withrevert(e, cert, root, ts, blk);
    bool reverted = lastReverted;

    assert ((futureTS || staleTS || invalidMsgHash || zeroTS || paused) => reverted),
        "future || stale || bad messageHash => revert";
}

rule confirmGlobalTableRoot_success_updates_and_frames(env e) {
    IBN254CertificateVerifierTypes.BN254Certificate cert;
    bytes32 root; uint32 ts; uint32 blk;

    // Snapshots
    uint32 latest_before = currentContract._latestReferenceTimestamp;

    // Other keys for frame checks
    uint32 otherTs;
    uint32 otherBlk;
    bytes32 otherRoot;

    bytes32 roots_other_before = currentContract._globalTableRoots[otherTs];
    uint32 blks_other_before = currentContract._referenceBlockNumbers[otherTs];
    uint32 ts_other_before = currentContract._referenceTimestamps[otherBlk];
    bool valid_other_before = currentContract._isRootValid[otherRoot];
    
    address _owner_before = currentContract._owner;
    uint16 threshold_before = currentContract.globalRootConfirmationThreshold;
    uint256 paused_before = currentContract._paused;


    // Call
    confirmGlobalTableRoot(e, cert, root, ts, blk);

    // Post
    assert (
        currentContract._latestReferenceTimestamp == ts
        && currentContract._referenceBlockNumbers[ts] == blk
        && currentContract._referenceTimestamps[blk] == ts
        && currentContract._globalTableRoots[ts] == root
        && currentContract._isRootValid[root] == true
    );

    // Frame: no unintended writes
    assert (unchangedMappingBytes32(roots_other_before,
                             currentContract._globalTableRoots[otherTs],
                             otherTs == ts));
    assert (unchangedMappingUint32(blks_other_before,
                             currentContract._referenceBlockNumbers[otherTs],
                             otherTs == ts));
    assert (unchangedMappingUint32(ts_other_before,
                             currentContract._referenceTimestamps[otherBlk],
                             otherBlk == blk));
    assert ( (otherRoot == root) || (currentContract._isRootValid[otherRoot] == valid_other_before) );

    // Governance params don’t change here
    assert currentContract.globalRootConfirmationThreshold == threshold_before;
    assert currentContract._owner == _owner_before;
    assert currentContract._paused == paused_before;
}

rule confirmGlobalTableRoot_strictly_monotonic(env e) {
    IBN254CertificateVerifierTypes.BN254Certificate cert;
    bytes32 root; uint32 ts; uint32 blk;

    uint32 latest_before = currentContract._latestReferenceTimestamp;
    confirmGlobalTableRoot(e, cert, root, ts, blk);
    assert currentContract._latestReferenceTimestamp > latest_before;
}


rule updateOperatorTable_revert_conditions(env e) {
    uint32 referenceTimestamp; 
    bytes32 root; 
    uint32 idx; 
    bytes proof; 
    bytes operatorTable;

    // Snapshots to read pre-state in asserts
    bool rootValid_before = currentContract._isRootValid[root];
    bool emptyTable = operatorTable.length == 0;

    // an unset reference timestamp does not cause a revert but noop
    bool referenceTimestampSet = isReferenceTimestampSet(e, operatorTable, referenceTimestamp); 

    // the following revert cases only trigger if referenceTimestampSet == true
    bool staleTS = referenceTimestamp <= latestTsForOperatorTableBytes(e, operatorTable);
    bool mismatchedRoot = root != currentContract._globalTableRoots[referenceTimestamp];
    bool invalidCurvetype = !validCurvetype(e, operatorTable);

    uint256 _paused_before = currentContract._paused;
    uint256 mask = 1 << 1; // PAUSED_OPERATOR_TABLE_UPDATE == 1
    bool paused = (currentContract._paused & mask) == mask;

    updateOperatorTable@withrevert(e, referenceTimestamp, root, idx, proof, operatorTable);
    bool reverted = lastReverted;

    assert ((!rootValid_before || paused || emptyTable ) => reverted),
        "!_isRootValid[root] || _globalTableRoots[ts] != root || paused => revert";
    assert !referenceTimestampSet => ((mismatchedRoot || staleTS || invalidCurvetype) => reverted);
}

rule updateOperatorTable_frame_conditions(env e) {
    uint32 ts; bytes32 root; uint32 idx; bytes proof; bytes table;

    // Snapshots
    uint32 latest_before = currentContract._latestReferenceTimestamp;
    bytes32 mapped_before = currentContract._globalTableRoots[ts];
    uint16 thresh_before = currentContract.globalRootConfirmationThreshold;
    address owner_before = currentContract._owner;
    uint256 paused_before = currentContract._paused;
    bool rootValid_before_state = currentContract._isRootValid[root];

    updateOperatorTable(e, ts, root, idx, proof, table);

    // No writes to these in implementation
    assert currentContract._latestReferenceTimestamp == latest_before;
    assert currentContract._globalTableRoots[ts] == mapped_before;
    assert currentContract.globalRootConfirmationThreshold == thresh_before;
    assert currentContract._owner == owner_before;
    assert currentContract._paused == paused_before;
    assert currentContract._isRootValid[root] == rootValid_before_state;
}

rule disableRoot_then_updateOperatorTable_must_revert(env e) {
    bytes32 root; uint32 ts; uint32 idx; bytes proof; bytes table;

    // Disable
    disableRoot(e, root);
    assert !currentContract._isRootValid[root];

    // Now any attempt to use it must revert (independent of other args)
    updateOperatorTable@withrevert(e, ts, root, idx, proof, table);
    assert lastReverted;
}


rule setGlobalRootConfirmationThreshold_revert_conditions(env e) {
    uint16 bps;

    // onlyOwner
    setGlobalRootConfirmationThreshold@withrevert(e, bps);
    bool reverted = lastReverted; 

    assert ((e.msg.sender != currentContract._owner || bps > currentContract.MAX_BPS(e)) => reverted), "onlyOwner";
}

rule setGlobalRootConfirmationThreshold_success_and_frame_conditions(env e) {
    uint16 bps;
    uint32 someTs;
    uint32 someBlk;

    address owner = currentContract._owner;
    uint32 referenceTimestamp = currentContract._latestReferenceTimestamp;
    uint32 refBlockNum_before   = currentContract._referenceBlockNumbers[someTs];
    uint32 refTimestamp_before  = currentContract._referenceTimestamps[someBlk];

    // success updates
    setGlobalRootConfirmationThreshold(e, bps);

    // integrity check
    assert (currentContract.globalRootConfirmationThreshold == bps);

    // frame on unrelated state
    assert currentContract._latestReferenceTimestamp == referenceTimestamp;
    assert currentContract._owner == owner;
    assert currentContract._referenceBlockNumbers[someTs] == refBlockNum_before;
    assert currentContract._referenceTimestamps[someBlk]  == refTimestamp_before;
}


rule disableRoot_revert_conditions(env e){
    bytes32 root;

    bool valid_before = currentContract._isRootValid[root];
    bool pauser = currentContract.pauserRegistry.isPauser(e, e.msg.sender);

    disableRoot@withrevert(e, root);

    assert ((!pauser || !valid_before) => lastReverted), "valid root || onlyPauser";
}

rule disableRoot_success_and_frame_conditions(env e) {
    bytes32 root;
    bytes32 other;

    bool other_before = currentContract._isRootValid[other];

    // success effect
    disableRoot(e, root);

    assert (currentContract._isRootValid[root] == false);
    assert ( (other == root) || (currentContract._isRootValid[other] == other_before) );
}