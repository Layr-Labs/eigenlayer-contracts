import "../optimizations.spec";

using AllocationManager as AllocationManager;
methods {
    function AllocationManager.DEALLOCATION_DELAY() external returns(uint32) envfree;
    function AllocationManager.getMaxMagnitude(address,address) external returns (uint64) envfree;

    // external calls to AVSRegistrar.  Note that the source does not have a proper implementation, the one available always reverts
    function _.registerOperator(address,address,uint32[],bytes) external => DISPATCHER(true);
    function _.deregisterOperator(address,address,uint32[]) external => DISPATCHER(true);
    function _.supportsAVS(address) external => DISPATCHER(true);

    // external calls to Strategy contracts
    function _.underlyingToken() external => DISPATCHER(true);
    function _.withdraw(address,address,uint256) external => DISPATCHER(true);

    // external calls to ERC20
    function _.balanceOf(address) external => DISPATCHER(true);
    function _.transfer(address,uint256) external => DISPATCHER(true);

    function OperatorSetLib.key(OperatorSetLib.OperatorSet memory os) internal returns (bytes32) => returnOperatorSetKey(os); // expect (bytes32); // return unique bytes32 that is not zero.

    // internal math summary to avoid overflows from the tool;
    // function AllocationManager._addInt128(uint64 a, int128 b) internal returns (uint64) => cvlAddInt128(a, b);
}

function cvlAddInt128(uint64 a, int128 b) returns uint64 {
    require(b >= 0 || to_mathint(a) > to_mathint(-b)); // Prevent underflow
    require(b <= 0 || a < to_mathint(max_uint64) -to_mathint(b)); // Prevent overflow
    return require_uint64(to_mathint(a) + to_mathint(b));
}

function returnOperatorSetKey(OperatorSetLib.OperatorSet os) returns bytes32 {
    return idToKey[os.id];
}

ghost mapping(uint32 => bytes32) idToKey {
    axiom forall uint32 id1 . forall uint32 id2 . (idToKey[id1] != to_bytes32(0) && idToKey[id2] != to_bytes32(0)) &&
        (id1 != id2 => idToKey[id1] != idToKey[id2]);
}

// -----------------------------------------
// Constants and Definitions
// -----------------------------------------

definition WAD() returns uint64 = 1000000000000000000; // definition uint64 WAD = 1e18

definition inRegisteredSets(address operator, bytes32 value) returns bool = (registeredSetsIndexes[operator][value] != 0);

definition inAllocatedSets(address operator, bytes32 value) returns bool = (allocatedSetsIndexes[operator][value] != 0);

// checks if a given index is within the dealocationQ bounds for a given operator and strategy.
definition inBound(address operator, address strategy, int128 index) returns bool = 
    index >= deallocationQueueBeginGhost[operator][strategy] && index < deallocationQueueEndGhost[operator][strategy];

// -----------------------------------------
// Ghost Mappings and Mirroring
// -----------------------------------------

// Mirrors `_maxMagnitudeHistory` keys
ghost mapping(address => mapping(address => mapping(uint256 => uint32))) maxMagnitudeHistoryKeys {
    init_state axiom forall address operator . forall address strategy . forall uint256 index . maxMagnitudeHistoryKeys[operator][strategy][index] == 0;
}
// Mirrors `_maxMagnitudeHistory` values
ghost mapping(address => mapping(address => mapping(mathint => uint224))) maxMagnitudeHistoryValues {
    init_state axiom forall address operator . forall address strategy . forall mathint index . maxMagnitudeHistoryValues[operator][strategy][index] == 0;
}
// Mirrors `_maxMagnitudeHistory` lengths
ghost mapping(address => mapping(address => uint256)) maxMagnitudeHistoryLengths {
    init_state axiom forall address operator . forall address strategy . maxMagnitudeHistoryLengths[operator][strategy] == 0;
}

// Mirrors `_operatorSets` values
ghost mapping(address => mapping(uint256 => bytes32)) operatorSetsValues {
    init_state axiom forall address avs . forall uint256 index . operatorSetsValues[avs][index] == to_bytes32(0);
}
// Mirrors `_operatorSets` indexes
ghost mapping(address => mapping(bytes32 => uint256)) operatorSetsIndexes {
    init_state axiom forall address avs . forall bytes32 value . operatorSetsIndexes[avs][value] == 0;
}
// Mirrors `_operatorSets` lengths
ghost mapping(address => uint256) ghostOperatorSetsLengths {
    init_state axiom forall address avs . ghostOperatorSetsLengths[avs] == 0;
    axiom forall address avs . ghostOperatorSetsLengths[avs] < 0xffffffffffffffffffffffffffffffff;
}

// Mirrors `registeredSets` values
ghost mapping(address => mapping(uint256 => bytes32)) registeredSetsValues {
    init_state axiom forall address operator . forall uint256 index . registeredSetsValues[operator][index] == to_bytes32(0);
}
// Mirrors `registeredSets` indexes
ghost mapping(address => mapping(bytes32 => uint256)) registeredSetsIndexes {
    init_state axiom forall address operator . forall bytes32 value . registeredSetsIndexes[operator][value] == 0;
}
// Mirrors `registeredSets` lengths
ghost mapping(address => uint256) ghostLengths {
    init_state axiom forall address operator . ghostLengths[operator] == 0;
    axiom forall address operator . ghostLengths[operator] < 0xffffffffffffffffffffffffffffffff;
}

// Mirrors `registrationStatus` registered
ghost mapping(address => mapping(bytes32 => bool)) registrationStatusRegistered {
    init_state axiom forall address operator . forall bytes32 operatorSetKey . registrationStatusRegistered[operator][operatorSetKey] == false;
}
// Mirrors `registrationStatus` slashableUntil
ghost mapping(address => mapping(bytes32 => uint32)) registrationStatusSlashableUntil {
    init_state axiom forall address operator . forall bytes32 operatorSetKey . registrationStatusSlashableUntil[operator][operatorSetKey] == 0;
}

// Mirrors `allocatedSets` values
ghost mapping(address => mapping(uint256 => bytes32)) allocatedSetsValues {
    init_state axiom forall address operator . forall uint256 index . allocatedSetsValues[operator][index] == to_bytes32(0);
}
// Mirrors `allocatedSets` indexes
ghost mapping(address => mapping(bytes32 => uint256)) allocatedSetsIndexes {
    init_state axiom forall address operator . forall bytes32 value . allocatedSetsIndexes[operator][value] == 0;
}
// Mirrors `allocatedSets` lengths
ghost mapping(address => uint256) allocatedSetsGhostLengths {
    init_state axiom forall address operator . allocatedSetsGhostLengths[operator] == 0;
    axiom forall address operator . allocatedSetsGhostLengths[operator] < 0xffffffffffffffffffffffffffffffff;
}

// Mirrors `deallocationQueue` data
ghost mapping(address => mapping(address => mapping(int128 => bytes32))) deallocationQueueDataGhost {
    init_state axiom forall address operator . forall address strategy . forall int128 index . deallocationQueueDataGhost[operator][strategy][index] == to_bytes32(0);
}
 // Mirrors `deallocationQueue` begin
ghost mapping(address => mapping(address => int128)) deallocationQueueBeginGhost {
    init_state axiom forall address operator . forall address strategy . deallocationQueueBeginGhost[operator][strategy] == 0;
}
// Mirrors `deallocationQueue` end
ghost mapping(address => mapping(address => int128)) deallocationQueueEndGhost {
    init_state axiom forall address operator . forall address strategy . deallocationQueueEndGhost[operator][strategy] == 0;
}

// mirroring mapping(address operator => mapping(bytes32 operatorSetKey => mapping(IStrategy strategy => Allocation))) allocations;
// Mirrors `allocations` currentMagnitude
ghost mapping(address => mapping(bytes32 => mapping(address => uint64))) allocationsCurrentMagnitude {
    init_state axiom forall address operator . forall bytes32 setKey . forall address strategy . allocationsCurrentMagnitude[operator][setKey][strategy] == 0;
}
// Holds the sum of all current magnitudes per operator and strategy (across setKeys)
ghost mapping(address => mapping(address => mathint)) sumAllocationsCurrentMagnitude {
    init_state axiom forall address operator . forall address strategy . sumAllocationsCurrentMagnitude[operator][strategy] == 0;
}
/// The initial value is being updated as we access the allocations current magnitudes one-by-one.
/// Should only be used as an initial value, never post-action!
ghost mapping(address => mapping(address => mathint)) sumAllocationsCurrentMagnitude_init {
    init_state axiom forall address operator . forall address strategy . sumAllocationsCurrentMagnitude_init[operator][strategy] == 0;
}

ghost mapping(address => mapping(bytes32 => mapping(address => bool))) didAccessOpAndStrategy;

function SumTrackingSetup() {
    require forall address operator . forall address strategy . sumAllocationsCurrentMagnitude[operator][strategy] == 
        sumAllocationsCurrentMagnitude_init[operator][strategy] + sumAllocationsPendindDiff_init[operator][strategy];
    require forall address operator . forall bytes32 key . forall address strategy . !didAccessOpAndStrategy[operator][key][strategy];
    require forall address operator . forall bytes32 key . forall address strategy . !didAccessOpAndStrategyDiff[operator][key][strategy];
}

// Mirrors `allocations` pendingDiff
ghost mapping(address => mapping(bytes32 => mapping(address => int128))) allocationsPendingDiff {
    init_state axiom forall address operator . forall bytes32 setKey . forall address strategy . allocationsPendingDiff[operator][setKey][strategy] == 0;
}
// Holds the sum of all pendingDiffs per operator and strategy (across setKeys)
ghost mapping(address => mapping(address => mathint)) sumAllocationsPendindDiff { // only positive
    init_state axiom forall address operator . forall address strategy . sumAllocationsPendindDiff[operator][strategy] == 0;
}
/// The initial value is being updated as we access the allocations pending diffs one-by-one.
/// Should only be used as an initial value, never post-action!
ghost mapping(address => mapping(address => mathint)) sumAllocationsPendindDiff_init { // only positive
    init_state axiom forall address operator . forall address strategy . sumAllocationsPendindDiff_init[operator][strategy] == 0;
}

ghost mapping(address => mapping(bytes32 => mapping(address => bool))) didAccessOpAndStrategyDiff;

// Mirrors `allocations` effectBlock
ghost mapping(address => mapping(bytes32 => mapping(address => uint32))) allocationsEffectBlock {
    init_state axiom forall address operator . forall bytes32 setKey . forall address strategy . allocationsEffectBlock[operator][setKey][strategy] == 0;
}

// Mirrors `encumberedMagnitude`
ghost mapping(address => mapping(address => uint64)) encumberedMagnitudeMirror {
    init_state axiom forall address operator . forall address strategy . encumberedMagnitudeMirror[operator][strategy] == 0;
}

// -----------------------------------------
// Hooks for Synchronizing Ghost Mappings
// -----------------------------------------

// Sync keys when `_maxMagnitudeHistory` is updated
hook Sstore AllocationManager._maxMagnitudeHistory[KEY address operator][KEY address strategy]._snapshots[INDEX uint256 index]._key uint32 newKey (uint32 oldKey) {
    require oldKey == maxMagnitudeHistoryKeys[operator][strategy][index];
    maxMagnitudeHistoryKeys[operator][strategy][index] = newKey;
}
// Sync values when `_maxMagnitudeHistory` is updated
hook Sstore AllocationManager._maxMagnitudeHistory[KEY address operator][KEY address strategy]._snapshots[INDEX uint256 index]._value uint224 newValue (uint224 oldValue) {
    require oldValue == maxMagnitudeHistoryValues[operator][strategy][index];
    maxMagnitudeHistoryValues[operator][strategy][index] = newValue;
}
// Sync lengths when `_maxMagnitudeHistory` is updated
hook Sstore AllocationManager._maxMagnitudeHistory[KEY address operator][KEY address strategy]._snapshots.length uint256 newLength (uint256 oldLength) {
    require oldLength == maxMagnitudeHistoryLengths[operator][strategy];
    maxMagnitudeHistoryLengths[operator][strategy] = newLength;
}
// Validate snapshot length, keys and values during access
hook Sload uint256 length AllocationManager._maxMagnitudeHistory[KEY address operator][KEY address strategy]._snapshots.length {
    require length == maxMagnitudeHistoryLengths[operator][strategy];
}

hook Sload uint32 key AllocationManager._maxMagnitudeHistory[KEY address operator][KEY address strategy]._snapshots[INDEX uint256 index]._key {
    require key == maxMagnitudeHistoryKeys[operator][strategy][index];
}

hook Sload uint224 value AllocationManager._maxMagnitudeHistory[KEY address operator][KEY address strategy]._snapshots[INDEX uint256 index]._value {
    require value == maxMagnitudeHistoryValues[operator][strategy][index];
}

// Sync lengths when `_operatorSets` is updated
hook Sstore AllocationManager._operatorSets[KEY address operator].(offset 0) uint256 newLength {
    ghostOperatorSetsLengths[operator] = newLength;
}
// Sync values when `_operatorSets` is updated
hook Sstore AllocationManager._operatorSets[KEY address operator]._inner._values[INDEX uint256 index] bytes32 newValue {
    operatorSetsValues[operator][index] = newValue;
}
// Sync indexes when `_operatorSets` is updated
hook Sstore AllocationManager._operatorSets[KEY address operator]._inner._indexes[KEY bytes32 value] uint256 newIndex {
    operatorSetsIndexes[operator][value] = newIndex;
}
// Validate length, indexes and values during access
hook Sload uint256 length AllocationManager._operatorSets[KEY address operator].(offset 0) {
    require ghostOperatorSetsLengths[operator] == length;
}

hook Sload bytes32 value AllocationManager._operatorSets[KEY address operator]._inner._values[INDEX uint256 index] {
    require operatorSetsValues[operator][index] == value;
}

hook Sload uint256 index AllocationManager._operatorSets[KEY address operator]._inner._indexes[KEY bytes32 value] {
    require operatorSetsIndexes[operator][value] == index;
}

// Sync lengths when `registeredSets` is updated
hook Sstore AllocationManager.registeredSets[KEY address operator].(offset 0) uint256 newLength {
    ghostLengths[operator] = newLength;
}
// Sync values when `registeredSets` is updated
hook Sstore AllocationManager.registeredSets[KEY address operator]._inner._values[INDEX uint256 index] bytes32 newValue {
    registeredSetsValues[operator][index] = newValue;
}
// Sync indexes when `registeredSets` is updated
hook Sstore AllocationManager.registeredSets[KEY address operator]._inner._indexes[KEY bytes32 value] uint256 newIndex {
    registeredSetsIndexes[operator][value] = newIndex;
}
// Validate length, indexes and values during access
hook Sload uint256 length AllocationManager.registeredSets[KEY address operator].(offset 0) {
    require ghostLengths[operator] == length;
}

hook Sload bytes32 value AllocationManager.registeredSets[KEY address operator]._inner._values[INDEX uint256 index] {
    require registeredSetsValues[operator][index] == value;
}

hook Sload uint256 index AllocationManager.registeredSets[KEY address operator]._inner._indexes[KEY bytes32 value] {
    require registeredSetsIndexes[operator][value] == index;
}

// Sync registered when `registrationStatus` is updated
hook Sstore AllocationManager.registrationStatus[KEY address operator][KEY bytes32 operatorSetKey].registered bool newRegistered {
    registrationStatusRegistered[operator][operatorSetKey] = newRegistered;
}
// Sync slashableUntil when `registrationStatus` is updated
hook Sstore AllocationManager.registrationStatus[KEY address operator][KEY bytes32 operatorSetKey].slashableUntil uint32 newSlashableUntil {
    registrationStatusSlashableUntil[operator][operatorSetKey] = newSlashableUntil;
}
// Validate slashableUntil and registered during access
hook Sload bool registered AllocationManager.registrationStatus[KEY address operator][KEY bytes32 operatorSetKey].registered {
    require registrationStatusRegistered[operator][operatorSetKey] == registered;
}

hook Sload uint32 slashableUntil AllocationManager.registrationStatus[KEY address operator][KEY bytes32 operatorSetKey].slashableUntil {
    require registrationStatusSlashableUntil[operator][operatorSetKey] == slashableUntil;
}

// Sync lengths when `allocatedSets` is updated
hook Sstore AllocationManager.allocatedSets[KEY address operator].(offset 0) uint256 newLength {
    allocatedSetsGhostLengths[operator] = newLength;
}
// Sync values when `allocatedSets` is updated
hook Sstore AllocationManager.allocatedSets[KEY address operator]._inner._values[INDEX uint256 index] bytes32 newValue {
    allocatedSetsValues[operator][index] = newValue;
}
// Sync indexes when `allocatedSets` is updated
hook Sstore AllocationManager.allocatedSets[KEY address operator]._inner._indexes[KEY bytes32 value] uint256 newIndex {
    allocatedSetsIndexes[operator][value] = newIndex;
}
// Validate length, indexes and values during access
hook Sload uint256 length AllocationManager.allocatedSets[KEY address operator].(offset 0) {
    require allocatedSetsGhostLengths[operator] == length;
}

hook Sload bytes32 value AllocationManager.allocatedSets[KEY address operator]._inner._values[INDEX uint256 index] {
    require allocatedSetsValues[operator][index] == value;
}

hook Sload uint256 index AllocationManager.allocatedSets[KEY address operator]._inner._indexes[KEY bytes32 value] {
    require allocatedSetsIndexes[operator][value] == index;
}

// Sync data when `deallocationQueue` is updated
hook Sstore AllocationManager.deallocationQueue[KEY address operator][KEY address strategy]._data[KEY int128 index] bytes32 newValue (bytes32 oldValue) {
    require oldValue == deallocationQueueDataGhost[operator][strategy][index];
    deallocationQueueDataGhost[operator][strategy][index] = newValue;
}
hook Sload bytes32 value AllocationManager.deallocationQueue[KEY address operator][KEY address strategy]._data[KEY int128 index] {
    require value == deallocationQueueDataGhost[operator][strategy][index];
}

// Sync begin when `deallocationQueue` is updated
hook Sstore AllocationManager.deallocationQueue[KEY address operator][KEY address strategy]._begin int128 newBegin (int128 oldBegin) {
    require oldBegin ==  deallocationQueueBeginGhost[operator][strategy];
    deallocationQueueBeginGhost[operator][strategy] = newBegin;
}
hook Sload int128 begin AllocationManager.deallocationQueue[KEY address operator][KEY address strategy]._begin {
    require begin == deallocationQueueBeginGhost[operator][strategy];
}

// Sync end when `deallocationQueue` is updated
hook Sstore AllocationManager.deallocationQueue[KEY address operator][KEY address strategy]._end int128 newEnd (int128 oldEnd) {
    require oldEnd == deallocationQueueEndGhost[operator][strategy];
    deallocationQueueEndGhost[operator][strategy] = newEnd;
}
hook Sload int128 end AllocationManager.deallocationQueue[KEY address operator][KEY address strategy]._end {
    require end == deallocationQueueEndGhost[operator][strategy];
}

// Sync currentMagnitude when `allocations` is updated
hook Sstore AllocationManager.allocations[KEY address operator][KEY bytes32 setKey][KEY address strategy].currentMagnitude uint64 newCurMag (uint64 oldCurMag) {
    require oldCurMag == allocationsCurrentMagnitude[operator][setKey][strategy];
    if (!didAccessOpAndStrategy[operator][setKey][strategy]){
        didAccessOpAndStrategy[operator][setKey][strategy] = true;
        sumAllocationsCurrentMagnitude_init[operator][strategy] = sumAllocationsCurrentMagnitude_init[operator][strategy] - oldCurMag;
        require sumAllocationsCurrentMagnitude_init[operator][strategy] >= 0;
    }
    sumAllocationsCurrentMagnitude[operator][strategy] = sumAllocationsCurrentMagnitude[operator][strategy] + newCurMag - oldCurMag;
    allocationsCurrentMagnitude[operator][setKey][strategy] = newCurMag;
}
hook Sload uint64 curMag AllocationManager.allocations[KEY address operator][KEY bytes32 setKey][KEY address strategy].currentMagnitude {
    require curMag == allocationsCurrentMagnitude[operator][setKey][strategy];
    if (!didAccessOpAndStrategy[operator][setKey][strategy]){
        didAccessOpAndStrategy[operator][setKey][strategy] = true;
        sumAllocationsCurrentMagnitude_init[operator][strategy] = sumAllocationsCurrentMagnitude_init[operator][strategy] - curMag;
        require sumAllocationsCurrentMagnitude_init[operator][strategy] >= 0;
    }
}

// keep track of pendingDiff valid state transition. if true, then an invalid state transition occurred.
ghost bool didMakeInvalidPendingDiffTransition;

// Sync pendingDiff when `allocations` is updated
hook Sstore AllocationManager.allocations[KEY address operator][KEY bytes32 setKey][KEY address strategy].pendingDiff int128 newPenDiff (int128 oldPenDiff) {
    require oldPenDiff == allocationsPendingDiff[operator][setKey][strategy];
    // summing only positive diff (as negative doesnt count until the slashing is enforced).
    // this sum is correct because of pendingDiffStateTransformation.
    /*
    For every storage load of the pendingDiff for some operator, setKey and strategy for the first time, 
    we deduct the storage amount from the initial value and require that the result is non-negative. 

    Example:
    Suppose that at first the initial sum of pendingDiffs that are going to be accessed in the flow is x.
    sumAllocationsPendindDiff_init[operator][strategy] = x.
    The function flow is going to access:
        allocations[operator][setKey1][strategy].pendingDiff = 1
        allocations[operator][setKey2][strategy].pendingDiff = 2
        allocations[operator][setKey3][strategy].pendingDiff = 3

    x(0) = x
    At first access [0x1]:    
        x(1) -> x(0) - 1 = x - 1 ; require x - 1 >=0
    At second access [0x2]:    
        x(2) -> x(1) - 2 = x - 3 ; require x - 3 >=0
    At third access [0x3]:
        x(3) -> x(2) - 3 = x - 6 ; require x - 6 >=0

    The result is that x = sumAllocationsPendindDiff_init[operator][strategy] >= 6 = 1 + 2 + 3, as expected.
    */
    if (newPenDiff >= 0) {
        if (oldPenDiff >= 0) {
            if (!didAccessOpAndStrategyDiff[operator][setKey][strategy]){
                didAccessOpAndStrategyDiff[operator][setKey][strategy] = true;
                sumAllocationsPendindDiff_init[operator][strategy] = sumAllocationsPendindDiff_init[operator][strategy] - oldPenDiff;
                require sumAllocationsPendindDiff_init[operator][strategy] >= 0;
            }
            sumAllocationsPendindDiff[operator][strategy] = sumAllocationsPendindDiff[operator][strategy] + newPenDiff - oldPenDiff;
        } else {
            if (!didAccessOpAndStrategyDiff[operator][setKey][strategy]){
                didAccessOpAndStrategyDiff[operator][setKey][strategy] = true;
                require sumAllocationsPendindDiff_init[operator][strategy] >= 0;
            }
            sumAllocationsPendindDiff[operator][strategy] = sumAllocationsPendindDiff[operator][strategy] + newPenDiff;
        }
    } else {
         if (oldPenDiff >= 0) {
            if (!didAccessOpAndStrategyDiff[operator][setKey][strategy]){
                didAccessOpAndStrategyDiff[operator][setKey][strategy] = true;
                sumAllocationsPendindDiff_init[operator][strategy] = sumAllocationsPendindDiff_init[operator][strategy] - oldPenDiff;
                require sumAllocationsPendindDiff_init[operator][strategy] >= 0;
            }
            sumAllocationsPendindDiff[operator][strategy] = sumAllocationsPendindDiff[operator][strategy] - oldPenDiff;
        } 
    }
    // keep track of pendingDiff valid state transition:
    // allocation.pendingDiff state transformation - from 0 to negative or positive, from positive zero only, from negative to negative or to zero.
    if (oldPenDiff > 0) {
        didMakeInvalidPendingDiffTransition = !(newPenDiff == 0 || oldPenDiff == newPenDiff) || didMakeInvalidPendingDiffTransition;
    } else if (oldPenDiff < 0) {
        didMakeInvalidPendingDiffTransition = newPenDiff > 0 || didMakeInvalidPendingDiffTransition;
    }
    allocationsPendingDiff[operator][setKey][strategy] = newPenDiff;
}
hook Sload int128 penDiff AllocationManager.allocations[KEY address operator][KEY bytes32 setKey][KEY address strategy].pendingDiff {
    require penDiff == allocationsPendingDiff[operator][setKey][strategy];
    if (!didAccessOpAndStrategyDiff[operator][setKey][strategy]){
        didAccessOpAndStrategyDiff[operator][setKey][strategy] = true;
        sumAllocationsPendindDiff_init[operator][strategy] = sumAllocationsPendindDiff_init[operator][strategy] - penDiff;
        require sumAllocationsPendindDiff_init[operator][strategy] >= 0;
    }
}

// Sync effectBlock when `allocations` is updated
hook Sstore AllocationManager.allocations[KEY address operator][KEY bytes32 setKey][KEY address strategy].effectBlock uint32 newEffectBlock (uint32 oldEffectBlock) {
    require oldEffectBlock == allocationsEffectBlock[operator][setKey][strategy];
    allocationsEffectBlock[operator][setKey][strategy] = newEffectBlock;
}
hook Sload uint32 effectBlock AllocationManager.allocations[KEY address operator][KEY bytes32 setKey][KEY address strategy].effectBlock {
    require effectBlock == allocationsEffectBlock[operator][setKey][strategy];
}

// Sync when `encumberedMagnitude` is updated
hook Sstore AllocationManager.encumberedMagnitude[KEY address operator][KEY address strategy] uint64 newEncumberedMagnitude (uint64 oldEncumberedMagnitude) {
    require oldEncumberedMagnitude == encumberedMagnitudeMirror[operator][strategy];
    encumberedMagnitudeMirror[operator][strategy] = newEncumberedMagnitude;
}
hook Sload uint64 encumberedMagnitude AllocationManager.encumberedMagnitude[KEY address operator][KEY address strategy] {
    require encumberedMagnitudeMirror[operator][strategy] == encumberedMagnitude;
}

// -----------------------------------------
// Invariants for Protocol Rules
// -----------------------------------------

function requireValidState() {
    requireInvariant maxMagnitudeHistoryKeysMonotonicInc();
    requireInvariant SetInRegisteredIFFStatusIsTrue();
    requireInvariant registeredSetsInvariant();
    requireInvariant allocatedSetsInvariant();
    requireInvariant operatorSetsInvariant();
    requireInvariant maxMagnitudeMonotonicallyDecreasing();
    requireInvariant maxMagnitudeHistoryPastLengthNullified();
    requireInvariant encumberedMagnitudeEqSumOfCurrentMagnitudesAndPositivePending();
    requireInvariant negativePendingDiffAtMostCurrentMagnitude();
    requireInvariant noZeroKeyInDealocationQ();
    requireInvariant EndGreaterThenBegin();
    requireInvariant deallocationQueueDataOutOfBoundsAreNullified();
    requireInvariant noPositivePendingDiffInDeallocationQ();
    requireInvariant effectBlockZeroHasNoPendingDiff();
    requireInvariant deallocationQueueDataUniqueness();
}

// Ensures consistency and correctness of the `registeredSets` mapping:
// 1. `registeredSetsIndexes` must correctly map to `registeredSetsValues`.
// 2. Non-zero indexes must lie within [1, ghostLengths[operator]].
// 3. Zero index indicates the value is not in the set.
invariant registeredSetsInvariant()
    (forall address operator . forall uint256 index. 0 <= index && index < ghostLengths[operator] => to_mathint(registeredSetsIndexes[operator][registeredSetsValues[operator][index]]) == index + 1)
    && (forall address operator . forall bytes32 value. registeredSetsIndexes[operator][value] == 0 || 
         ((forall uint256 tmp. to_mathint(tmp) == (registeredSetsIndexes[operator][value] - 1) => registeredSetsValues[operator][tmp] == value)
          && registeredSetsIndexes[operator][value] >= 1 && registeredSetsIndexes[operator][value] <= ghostLengths[operator]));

// Ensures consistency and correctness of the `allocatedSets` mapping:
// 1. `allocatedSetsIndexes` must correctly map to `allocatedSetsValues`.
// 2. Non-zero indexes must lie within [1, allocatedSetsGhostLengths[operator]].
// 3. Zero index indicates the value is not in the set.
invariant allocatedSetsInvariant()
    (forall address operator . forall uint256 index. 0 <= index && index < allocatedSetsGhostLengths[operator] => to_mathint(allocatedSetsIndexes[operator][allocatedSetsValues[operator][index]]) == index + 1)
    && (forall address operator . forall bytes32 value. allocatedSetsIndexes[operator][value] == 0 || 
         ((forall uint256 tmp. to_mathint(tmp) == (allocatedSetsIndexes[operator][value] - 1) => allocatedSetsValues[operator][tmp] == value)
          && allocatedSetsIndexes[operator][value] >= 1 && allocatedSetsIndexes[operator][value] <= allocatedSetsGhostLengths[operator]));

// Ensures consistency and correctness of the `_operatorSets` mapping:
// 1. `operatorSetsIndexes` must correctly map to `operatorSetsValues`.
// 2. Non-zero indexes must lie within [1, ghostOperatorSetsLengths[operator]].
// 3. Zero index indicates the value is not in the set.
invariant operatorSetsInvariant() // needed for deallocationQueueDataUniqueness
    (forall address avs . forall uint256 index. 0 <= index && index < ghostOperatorSetsLengths[avs] => to_mathint(operatorSetsIndexes[avs][operatorSetsValues[avs][index]]) == index + 1)
    && (forall address avs . forall bytes32 value. operatorSetsIndexes[avs][value] == 0 || 
         ((forall uint256 tmp. to_mathint(tmp) == (operatorSetsIndexes[avs][value] - 1) => operatorSetsValues[avs][tmp] == value)
          && operatorSetsIndexes[avs][value] >= 1 && operatorSetsIndexes[avs][value] <= ghostOperatorSetsLengths[avs]));

/// @title _maxMagnitudeHistory.keys are monotonic increasing
/// @property MagnitudeHistory Keys are valid - Snapshots' timestamps are sorted in increasing order.
// _maxMagnitudeHistory.key are the block number at the update time, should be monotonic increasing
invariant maxMagnitudeHistoryKeysMonotonicInc()
    forall address operator . forall address strategy . forall uint256 index1 . forall uint256 index2 .
        index1 < maxMagnitudeHistoryLengths[operator][strategy] && index2 <= index1 =>
        maxMagnitudeHistoryKeys[operator][strategy][index2] <= maxMagnitudeHistoryKeys[operator][strategy][index1];

// Ensures that entries beyond `maxMagnitudeHistoryLengths` are nullified:
// 1. For each operator and strategy, all keys and values at indices >= `maxMagnitudeHistoryLengths` must be zero.
invariant maxMagnitudeHistoryPastLengthNullified()
    forall address operator . forall address strategy . forall uint256 index . index >= maxMagnitudeHistoryLengths[operator][strategy] =>
        (maxMagnitudeHistoryKeys[operator][strategy][index] == 0 && maxMagnitudeHistoryValues[operator][strategy][index] == 0)
        {
            preserved {
                // prevent overflows for loop iter = 2
                require forall address operator . forall address strategy . maxMagnitudeHistoryLengths[operator][strategy] < max_uint256 - 1;
            }
        }

/// @title  Ensures all keys in `maxMagnitudeHistoryKeys` are less than or equal to the current block number.
/// @property MagnitudeHistory Keys are valid
invariant maxMagnitudeHistoryKeysLessThanCurrentBlock(env e)
    forall address operator . forall address strategy . forall uint256 index . index < maxMagnitudeHistoryLengths[operator][strategy] =>
        maxMagnitudeHistoryKeys[operator][strategy][index] <= e.block.number 
        {
            preserved with (env e1) {
                require e1.block.number == e.block.number;
                requireInvariant registeredSetsInvariant();
                requireInvariant allocatedSetsInvariant();
            }
        }

/// @title maxMagnitude is monotonically decreasing
/// @property maxMagnitude is monotonically decreasing
invariant maxMagnitudeMonotonicallyDecreasing()
    forall address operator . forall address strategy . forall uint256 index1 . forall uint256 index2 . (index1 >= index2 =>
        maxMagnitudeHistoryValues[operator][strategy][index1] <= maxMagnitudeHistoryValues[operator][strategy][index2]) 
        && (maxMagnitudeHistoryLengths[operator][strategy] > 0 => maxMagnitudeHistoryValues[operator][strategy][0] <= WAD())
    {
        preserved with (env e) {
            requireInvariant maxMagnitudeHistoryKeysMonotonicInc();
            requireInvariant registeredSetsInvariant();
            requireInvariant allocatedSetsInvariant();
            requireInvariant maxMagnitudeHistoryPastLengthNullified();
            requireInvariant maxMagnitudeHistoryKeysLessThanCurrentBlock(e);
        }
    }

/// @title All setKeys in registeredSets should have registrationStatus as true per operator and vice versa
/// @property Ensure consistency between registered sets and their registration statuses
invariant SetInRegisteredIFFStatusIsTrue()
    forall address operator . forall bytes32 operatorSetKey . inRegisteredSets(operator, operatorSetKey) <=> registrationStatusRegistered[operator][operatorSetKey] 
    {
        preserved {
            requireInvariant registeredSetsInvariant();
        }
    }

// encumberedMagnitude per strategy S = sum all allocation.current magnitude per strategy S + sum all allocation.pending per strategy S if pending > 0;
/// @title encumberedMagnitude equals the sum of currentMagnitudes and positive pendingDiffs.
/// @property Magnitudes Solvency
invariant encumberedMagnitudeEqSumOfCurrentMagnitudesAndPositivePending() 
    forall address operator . forall address strategy. encumberedMagnitudeMirror[operator][strategy] == 
        sumAllocationsCurrentMagnitude[operator][strategy] + sumAllocationsPendindDiff[operator][strategy]
        {
             preserved {
                SumTrackingSetup();
                requireInvariant registeredSetsInvariant();
                requireInvariant allocatedSetsInvariant();
            }
        }

/// @title For each operator and strategy tuple, max magnitude is greater or equal to the encumbered magnitude.
/// @property Magnitudes Solvency
invariant maxMagnitudeGEencumberedMagnitude(address operator, address strategy)
   AllocationManager.getMaxMagnitude(operator, strategy) >= encumberedMagnitudeMirror[operator][strategy] 
        {
            preserved with (env e) {
                requireValidState();
                SumTrackingSetup();
                requireInvariant maxMagnitudeHistoryKeysLessThanCurrentBlock(e);
                // assuming deallocation block indices dont overflow
                require forall address _operator . forall address _strategy . (deallocationQueueEndGhost[_operator][_strategy] < max_uint64 - 1);
                // would happen around the year 2833 to get block number equal to half of max_uint32
                require e.block.number < max_uint32 + AllocationManager.DEALLOCATION_DELAY + 1;  
                require e.block.number > 0;
                // prevent overflows for loop iter = 2
                require forall address operator_ . forall address strategy_ . maxMagnitudeHistoryLengths[operator_][strategy_] < max_uint256 - 1;
            }
        }

/// @title Cant have a neggative pendingDiff that is greater than the current magnitude. 
/// @property Magnitudes solvency
invariant negativePendingDiffAtMostCurrentMagnitude()
    forall address operator . forall bytes32 key. forall address strategy . allocationsPendingDiff[operator][key][strategy] < 0 => 
        -allocationsPendingDiff[operator][key][strategy] <= allocationsCurrentMagnitude[operator][key][strategy]
    {
        preserved with (env e) {
            requireValidState();
        } 
    }

/// @title At most 1 pending allocation/deallocation can be queued for a given (Operator, OperatorSet, Strategy) tuple
/// @property DeallocationQueue allocations uniqueness
invariant deallocationQueueDataUniqueness()
    forall address operator . forall address strategy . forall int128 index1 . forall int128 index2 .
        (index1 != index2) && inBound(operator, strategy, index1) && inBound(operator, strategy, index2) => 
            deallocationQueueDataGhost[operator][strategy][index1] != deallocationQueueDataGhost[operator][strategy][index2]
    {
        preserved with (env e) {
            requireValidState();
            requireInvariant maxMagnitudeHistoryKeysLessThanCurrentBlock(e);
            requireInvariant deallocationQueueEffectBlocLessThanCurrBlockNumberPlushDelayPlusOne(e);
            // assuming deallocation block indices dont overflow
            require forall address _operator . forall address _strategy . (deallocationQueueEndGhost[_operator][_strategy] < max_uint64 - 1);
            // would happen around the year 2833 to get block number equal to half of max_uint32
            require e.block.number < max_uint32 - AllocationManager.DEALLOCATION_DELAY - 1;  
            require e.block.number > 0;
            requireInvariant deallocationQueueEffectBlockAscesndingOrder(e);
        } 
    }

// Out-of-Bounds Deallocation Queue Data Nullification Invariant - helper Lemma
invariant deallocationQueueDataOutOfBoundsAreNullified() 
    forall address operator . forall address strategy . forall int128 index . (!inBound(operator, strategy, index)) <=> deallocationQueueDataGhost[operator][strategy][index] == to_bytes32(0)
    {
        preserved {
            // assuming deallocation block indices dont overflow
            require forall address operator . forall address strategy . (deallocationQueueEndGhost[operator][strategy] < max_uint64 - 1);
        }
    }

// Deallocation Queue Ordering: End Index ≥ Begin Index Invariant - helper Lemma
invariant EndGreaterThenBegin()
    forall address operator . forall address strategy . deallocationQueueBeginGhost[operator][strategy] <= deallocationQueueEndGhost[operator][strategy]
    {
        preserved {
            // assuming deallocation block indices dont overflow
            require forall address operator . forall address strategy . (deallocationQueueEndGhost[operator][strategy] < max_uint64 - 1);
        }
    }

/// @title Non-Zero Keys in Deallocation Queue Invariant
/// @property DeallocationQueue allocations uniqueness
invariant noZeroKeyInDealocationQ()
    forall address operator . forall address strategy . forall int128 index . inBound(operator, strategy, index) => deallocationQueueDataGhost[operator][strategy][index] != to_bytes32(0);

/// @title Allocations in the deallocation queue have their effect blocks equal at most the latest block number + DEALLOCATION_DELAY + 1.
/// @property Deallocation Queue Effect Block Timing Bound Invariant.
invariant deallocationQueueEffectBlocLessThanCurrBlockNumberPlushDelayPlusOne(env e)
    forall address operator . forall address strategy . forall int128 index . 
        inBound(operator, strategy, index) => (allocationsEffectBlock[operator][deallocationQueueDataGhost[operator][strategy][index]][strategy] <= e.block.number + AllocationManager.DEALLOCATION_DELAY + 1) 
        {
            preserved with (env e1) {
                requireValidState();
                SumTrackingSetup();
                requireInvariant maxMagnitudeHistoryKeysLessThanCurrentBlock(e);
                // assuming deallocation block indices dont overflow
                require forall address _operator . forall address _strategy . (deallocationQueueEndGhost[_operator][_strategy] < max_uint64 - 1);
                require e1.block.number == e.block.number;
                // would happen around the year 2833 to get block number equal to half of max_uint32
                require e1.block.number < max_uint32 - AllocationManager.DEALLOCATION_DELAY - 1;
                require e1.block.number > 0;  

                // assuming deallocation block indices dont overflow
                require forall address operator . forall address strategy . (deallocationQueueEndGhost[operator][strategy] < max_uint64 - 1);
                requireInvariant deallocationQueueEffectBlockAscesndingOrder(e1);
            }
        }

/// @title 
/*
Ensures that effect blocks in the deallocation queue are sorted in ascending order, 
except when an effect block is zero, which must have been less than or equal to the current block number before. 
This prevents out-of-order deallocations and ensures a valid execution sequence.
*/
/// @property Deallocation Queue Effect Block Ascending Order Invariant
invariant deallocationQueueEffectBlockAscesndingOrder(env e1)
    forall address operator . forall address strategy . forall int128 index1 . forall int128 index2 .
        (index1 <= index2) && inBound(operator, strategy, index1) && inBound(operator, strategy, index2) => 
            ((allocationsEffectBlock[operator][deallocationQueueDataGhost[operator][strategy][index1]][strategy] <= 
            allocationsEffectBlock[operator][deallocationQueueDataGhost[operator][strategy][index2]][strategy]) || 
            (allocationsEffectBlock[operator][deallocationQueueDataGhost[operator][strategy][index2]][strategy] == 0 &&
            allocationsEffectBlock[operator][deallocationQueueDataGhost[operator][strategy][index1]][strategy] <= e1.block.number))
    {
        preserved with (env e) {
            requireValidState();
            SumTrackingSetup();
            requireInvariant maxMagnitudeHistoryKeysLessThanCurrentBlock(e);
            // requireInvariant deallocationQueueDataOutOfBoundsAreNullified();
            requireInvariant deallocationQueueEffectBlocLessThanCurrBlockNumberPlushDelayPlusOne(e);
            // assuming deallocation block indices dont overflow
            require forall address _operator . forall address _strategy . (deallocationQueueEndGhost[_operator][_strategy] < max_uint64 - 1);
            // would happen around the year 2833 to get block number equal to half of max_uint32
            require e.block.number < max_uint32 - AllocationManager.DEALLOCATION_DELAY - 1;  
            require e.block.number > 0;
            require e.block.number == e1.block.number;
        }
    }

/// @title Ensures that if an allocation's effectBlock is zero, its pendingDiff must also be zero, preventing invalid pending modifications for inactive allocations.
/// @property No Pending Diff for Zero Effect Block
invariant effectBlockZeroHasNoPendingDiff()
    forall address operator . forall bytes32 setKey . forall address strategy .
        allocationsEffectBlock[operator][setKey][strategy] == 0 => allocationsPendingDiff[operator][setKey][strategy] == 0
    {
        preserved with (env e) {
            requireValidState();
            requireInvariant maxMagnitudeHistoryKeysLessThanCurrentBlock(e);
            // requireInvariant deallocationQueueDataOutOfBoundsAreNullified();
            requireInvariant deallocationQueueEffectBlocLessThanCurrBlockNumberPlushDelayPlusOne(e);
            // assuming deallocation block indices dont overflow
            require forall address _operator . forall address _strategy . (deallocationQueueEndGhost[_operator][_strategy] < max_uint64 - 1);
            // would happen around the year 2833 to get block number equal to half of max_uint32
            require e.block.number < max_uint32 - AllocationManager.DEALLOCATION_DELAY - 1;  
            require e.block.number > 0;
        }
    }

/// @title Ensures that all allocations in the deallocation queue have a pendingDiff value that is zero or negative, preventing improper positive adjustments during deallocation.
/// @property No Positive Pending Diff in Deallocation Queue
invariant noPositivePendingDiffInDeallocationQ() 
    forall address operator . forall address strategy . forall int128 index . 
        inBound(operator, strategy, index) => (allocationsPendingDiff[operator][deallocationQueueDataGhost[operator][strategy][index]][strategy] <= 0)
    {
        preserved with (env e) {
            requireValidState();
            requireInvariant maxMagnitudeHistoryKeysLessThanCurrentBlock(e);
            // requireInvariant deallocationQueueDataOutOfBoundsAreNullified();
            requireInvariant deallocationQueueEffectBlocLessThanCurrBlockNumberPlushDelayPlusOne(e);
            // assuming deallocation block indices dont overflow
            require forall address _operator . forall address _strategy . (deallocationQueueEndGhost[_operator][_strategy] < max_uint64 - 1);
            // would happen around the year 2833 to get block number equal to half of max_uint32
            require e.block.number < max_uint32 - AllocationManager.DEALLOCATION_DELAY - 1;  
            require e.block.number > 0;
            requireInvariant deallocationQueueEffectBlockAscesndingOrder(e);
        }
    }
