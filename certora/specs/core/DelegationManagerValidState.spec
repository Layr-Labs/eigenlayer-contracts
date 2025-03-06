import "../ptaHelpers.spec";

using DelegationManager as DelegationManager;

methods {
    //// External Calls
	// external calls to DelegationManager 
    function undelegate(address) external;
    function decreaseDelegatedShares(address,address,uint256) external;
	function increaseDelegatedShares(address,address,uint256) external;
    function getWithdrawableShares(address, address[]) external;

    // external calls from DelegationManager to ServiceManager
    function _.updateStakes(address[]) external => NONDET;

	// external calls to StrategyManager
    function _.getDeposits(address) external => DISPATCHER(true);
    function _.withdrawSharesAsTokens(address,address,address,uint256) external => DISPATCHER(true);

	// external calls to EigenPodManager
    function _.addShares(address,address,uint256) external => DISPATCHER(true);
    // function _.withdrawSharesAsTokens(address,address,address,uint256) external => DISPATCHER(true); // Already summarized in section for StrategyManager
    function _.podOwnerShares(address) external => NONDET;

    // external calls to EigenPod
	function _.withdrawRestakedBeaconChainETH(address,uint256) external => DISPATCHER(true);

    // external calls to PauserRegistry
    function _.isPauser(address) external => DISPATCHER(true);
	function _.unpauser() external => DISPATCHER(true);

    // external calls to Strategy contracts
    function _.withdraw(address,address,uint256) external => DISPATCHER(true);
    function _.deposit(address,uint256) external => DISPATCHER(true);
    function _.underlyingToken() external => DISPATCHER(true);

    // external calls to ERC20
    function _.balanceOf(address) external => DISPATCHER(true);
    function _.transfer(address, uint256) external => DISPATCHER(true);
    function _.transferFrom(address, address, uint256) external => DISPATCHER(true);

    // external calls to ERC1271 (can import OpenZeppelin mock implementation)
    // isValidSignature(bytes32 hash, bytes memory signature) returns (bytes4 magicValue) => DISPATCHER(true)
    function _.isValidSignature(bytes32, bytes) external => DISPATCHER(true);

    //// Harnessed Functions
    // Harnessed getters
    function get_operatorShares(address,address) external returns (uint256) envfree;
    function get_stakerDelegateableShares(address,address) external returns (uint256) envfree;
    function get_min_withdrawal_delay_blocks() external returns (uint32) envfree;

    //envfree functions
    function delegatedTo(address) external returns (address) envfree;
    function delegationApprover(address operator) external returns (address) envfree;
    function operatorShares(address operator, address strategy) external returns (uint256) envfree;
    function isDelegated(address staker) external returns (bool) envfree;
    function isOperator(address operator) external returns (bool) envfree;
    function delegationApproverSaltIsSpent(address delegationApprover, bytes32 salt) external returns (bool) envfree;
    function owner() external returns (address) envfree;
    function strategyManager() external returns (address) envfree;
    function eigenPodManager() external returns (address) envfree;
    function calculateWithdrawalRoot(IDelegationManagerTypes.Withdrawal) external returns (bytes32) envfree;
    function pendingWithdrawals(bytes32) external returns (bool) envfree;
    function depositScalingFactor(address, address) external returns (uint256) envfree;

    function _.beaconChainSlashingFactor(address staker) external => beaconChainSlashingFactorSummary(staker) expect uint64;
}

// -----------------------------------------
// External Calls Summary
// -----------------------------------------

ghost mapping(address => uint64) beaconChainSlashingFactorGhost {
    init_state axiom forall address staker . beaconChainSlashingFactorGhost[staker] == 0;
    // The following is proven for the eigenpod manager (beaconChainSlashingFactor is at most one WAD):
    axiom forall address staker . beaconChainSlashingFactorGhost[staker] <= WAD();
}

function beaconChainSlashingFactorSummary(address staker) returns uint64 {
    return beaconChainSlashingFactorGhost[staker];
}

// -----------------------------------------
// Constants and Definitions
// -----------------------------------------

definition WAD() returns uint64 = 1000000000000000000; // definition uint64 WAD = 1e18

definition inRootsSets(address staker, bytes32 value) returns bool = (stakerQueuedWithdrawalRootsSetsIndexes[staker][value] != 0);

// ----------------------------------------- 
// Ghost Mappings and Mirroring
// ----------------------------------------- 

// Mirrors `_cumulativeScaledSharesHistory` keys
ghost mapping(address => mapping(address => mapping(uint256 => uint32))) cumulativeScaledSharesHistoryKeys {
    init_state axiom forall address operator . forall address strategy . forall uint256 index . cumulativeScaledSharesHistoryKeys[operator][strategy][index] == 0;
}
// Mirrors `_cumulativeScaledSharesHistory` values
ghost mapping(address => mapping(address => mapping(mathint => uint224))) cumulativeScaledSharesHistoryValues {
    init_state axiom forall address operator . forall address strategy . forall mathint index . cumulativeScaledSharesHistoryValues[operator][strategy][index] == 0;
}
// Mirrors `_cumulativeScaledSharesHistory` lengths
ghost mapping(address => mapping(address => uint256)) cumulativeScaledSharesHistoryLengths {
    init_state axiom forall address operator . forall address strategy . cumulativeScaledSharesHistoryLengths[operator][strategy] == 0;
}

// Mirrors `operatorShares`
ghost mapping(address => mapping(address => uint256)) operatorSharesGhost {
    init_state axiom forall address operator . forall address strategy . operatorSharesGhost[operator][strategy] == 0;
}

// Mirror `_depositScalingFactor`
ghost mapping(address => mapping(address => uint256)) scalingFactorsGhost {
    init_state axiom forall address staker . forall address strategy . scalingFactorsGhost[staker][strategy] == 0;
}

// Mirrors `_stakerQueuedWithdrawalRoots` values
ghost mapping(address => mapping(uint256 => bytes32)) stakerQueuedWithdrawalRootsSetsValues {
    init_state axiom forall address staker . forall uint256 index . stakerQueuedWithdrawalRootsSetsValues[staker][index] == to_bytes32(0);
}
// Mirrors `_stakerQueuedWithdrawalRoots` indexes
ghost mapping(address => mapping(bytes32 => uint256)) stakerQueuedWithdrawalRootsSetsIndexes {
    init_state axiom forall address staker . forall bytes32 value . stakerQueuedWithdrawalRootsSetsIndexes[staker][value] == 0;
}
// Mirrors `_stakerQueuedWithdrawalRoots` lengths
ghost mapping(address => uint256) stakerQueuedWithdrawalRootsGhostLengths {
    init_state axiom forall address staker . stakerQueuedWithdrawalRootsGhostLengths[staker] == 0;
    axiom forall address staker . stakerQueuedWithdrawalRootsGhostLengths[staker] < 0xffffffffffffffffffffffffffffffff;
}

// Mirror `pendingWithdrawals`
ghost mapping(bytes32 => bool) pendingWithdrawalsMirror {
    init_state axiom forall bytes32 root . pendingWithdrawalsMirror[root] == false;
}

// Mirrors `queuedWithdrawals[root].staker`
ghost mapping(bytes32 => address) queuedWithdrawalsStakesrGhost {
    init_state axiom forall bytes32 root . queuedWithdrawalsStakesrGhost[root] == 0;
}

// Mirrors `_operatorDetails[operator].delegationApprover`
ghost mapping(address => address) operatorDetailsDelegationApproversMirror {
    init_state axiom forall address operator . operatorDetailsDelegationApproversMirror[operator] == 0;
}

// -----------------------------------------
// Hooks for Synchronizing Ghost Mappings
// -----------------------------------------

// Sync when `_operatorDetails[operator].delegationApprover` is updated
hook Sstore DelegationManager._operatorDetails[KEY address operator].delegationApprover address newDelegationApprover (address oldDelegationApprover) {
    require oldDelegationApprover == operatorDetailsDelegationApproversMirror[operator];
    operatorDetailsDelegationApproversMirror[operator] = newDelegationApprover;
}
hook Sload address delegationApprover DelegationManager._operatorDetails[KEY address operator].delegationApprover {
    require delegationApprover == operatorDetailsDelegationApproversMirror[operator];
}

// Sync when `queuedWithdrawals[root].staker` is updated
hook Sstore DelegationManager._queuedWithdrawals[KEY bytes32 root].staker address newStaker (address oldStaker) {
    require oldStaker == queuedWithdrawalsStakesrGhost[root];
    queuedWithdrawalsStakesrGhost[root] = newStaker;
}
hook Sload address staker DelegationManager._queuedWithdrawals[KEY bytes32 root].staker {
    require staker == queuedWithdrawalsStakesrGhost[root];
}

// Sync keys when `_cumulativeScaledSharesHistory` is updated
hook Sstore DelegationManager._cumulativeScaledSharesHistory[KEY address operator][KEY address strategy]._snapshots[INDEX uint256 index]._key uint32 newKey (uint32 oldKey) {
    require oldKey == cumulativeScaledSharesHistoryKeys[operator][strategy][index];
    cumulativeScaledSharesHistoryKeys[operator][strategy][index] = newKey;
}
// Sync values when `_cumulativeScaledSharesHistory` is updated
hook Sstore DelegationManager._cumulativeScaledSharesHistory[KEY address operator][KEY address strategy]._snapshots[INDEX uint256 index]._value uint224 newValue (uint224 oldValue) {
    require oldValue == cumulativeScaledSharesHistoryValues[operator][strategy][index];
    cumulativeScaledSharesHistoryValues[operator][strategy][index] = newValue;
}
// Sync lengths when `_cumulativeScaledSharesHistory` is updated
hook Sstore DelegationManager._cumulativeScaledSharesHistory[KEY address operator][KEY address strategy]._snapshots.length uint256 newLength (uint256 oldLength) {
    require oldLength == cumulativeScaledSharesHistoryLengths[operator][strategy];
    cumulativeScaledSharesHistoryLengths[operator][strategy] = newLength;
}
// Validate snapshot length, keys and values during access
hook Sload uint256 length DelegationManager._cumulativeScaledSharesHistory[KEY address operator][KEY address strategy]._snapshots.length {
    require length == cumulativeScaledSharesHistoryLengths[operator][strategy];
}

hook Sload uint32 key DelegationManager._cumulativeScaledSharesHistory[KEY address operator][KEY address strategy]._snapshots[INDEX uint256 index]._key {
    require key == cumulativeScaledSharesHistoryKeys[operator][strategy][index];
}

hook Sload uint224 value DelegationManager._cumulativeScaledSharesHistory[KEY address operator][KEY address strategy]._snapshots[INDEX uint256 index]._value {
    require value == cumulativeScaledSharesHistoryValues[operator][strategy][index];
}

// Sync data when `operatorShares` is updated
hook Sstore DelegationManager.operatorShares[KEY address operator][KEY address strategy] uint256 newShares (uint256 oldShares) {
    require oldShares == operatorSharesGhost[operator][strategy];
    operatorSharesGhost[operator][strategy] = newShares;
}
hook Sload uint256 shares DelegationManager.operatorShares[KEY address operator][KEY address strategy] {
    require shares == operatorSharesGhost[operator][strategy];
}

// Sync when `pendingWithdrawals` is updated
hook Sstore DelegationManager.pendingWithdrawals[KEY bytes32 root] bool newPending (bool oldPending) {
    require oldPending == pendingWithdrawalsMirror[root];
    pendingWithdrawalsMirror[root] = newPending;
}
hook Sload bool pending DelegationManager.pendingWithdrawals[KEY bytes32 root] {
    require pending == pendingWithdrawalsMirror[root];
}

// Sync scalingFactor when `_depositScalingFactor` is updated
hook Sstore DelegationManager._depositScalingFactor[KEY address operator][KEY address strategy]._scalingFactor uint256 newScalingFactor (uint256 oldScalingFactor) {
    require oldScalingFactor == scalingFactorsGhost[operator][strategy];
    scalingFactorsGhost[operator][strategy] = newScalingFactor;
}
hook Sload uint256 scalingFactor DelegationManager._depositScalingFactor[KEY address operator][KEY address strategy]._scalingFactor {
    require scalingFactor == scalingFactorsGhost[operator][strategy];
}

// Sync lengths when `_stakerQueuedWithdrawalRoots` is updated
hook Sstore DelegationManager._stakerQueuedWithdrawalRoots[KEY address operator].(offset 0) uint256 newLength {
    stakerQueuedWithdrawalRootsGhostLengths[operator] = newLength;
}
// Sync values when `_stakerQueuedWithdrawalRoots` is updated
hook Sstore DelegationManager._stakerQueuedWithdrawalRoots[KEY address operator]._inner._values[INDEX uint256 index] bytes32 newValue {
    stakerQueuedWithdrawalRootsSetsValues[operator][index] = newValue;
}
// Sync indexes when `_stakerQueuedWithdrawalRoots` is updated
hook Sstore DelegationManager._stakerQueuedWithdrawalRoots[KEY address operator]._inner._indexes[KEY bytes32 value] uint256 newIndex {
    stakerQueuedWithdrawalRootsSetsIndexes[operator][value] = newIndex;
}
// Validate length, indexes and values during access
hook Sload uint256 length DelegationManager._stakerQueuedWithdrawalRoots[KEY address operator].(offset 0) {
    require stakerQueuedWithdrawalRootsGhostLengths[operator] == length;
}

hook Sload bytes32 value DelegationManager._stakerQueuedWithdrawalRoots[KEY address operator]._inner._values[INDEX uint256 index] {
    require stakerQueuedWithdrawalRootsSetsValues[operator][index] == value;
}

hook Sload uint256 index DelegationManager._stakerQueuedWithdrawalRoots[KEY address operator]._inner._indexes[KEY bytes32 value] {
    require stakerQueuedWithdrawalRootsSetsIndexes[operator][value] == index;
}

// -----------------------------------------
// Invariants for Protocol Rules
// -----------------------------------------

function requireDelegationManagerValidState() {
    requireInvariant operatorDelegatesToThemselves();
    requireInvariant stakerQueuedWithdrawalRootsInvariant();
    requireInvariant cumulativeScaledSharesHistoryKeysMonotonicInc();
    requireInvariant cumulativeScaledSharesHistoryPastLengthNullified();
    requireInvariant cumulativeScaledSharesMonotonicallyIncreasing();
    requireInvariant queuedWithdrawalsCorrect();
}

/// @title 
/* 
Ensures consistency and correctness of the `stakerQueuedWithdrawalRoots` mapping:
1. `stakerQueuedWithdrawalRootsSetsIndexes` must correctly map to `stakerQueuedWithdrawalRootsSetsValues`.
2. Non-zero indexes must lie within [1, stakerQueuedWithdrawalRootsGhostLengths[operator]].
3. Zero index indicates the value is not in the set.
*/
/// @property Staker Queued Withdrawal Roots Consistency
invariant stakerQueuedWithdrawalRootsInvariant()
    (forall address operator . forall uint256 index. 0 <= index && index < stakerQueuedWithdrawalRootsGhostLengths[operator] => 
    to_mathint(stakerQueuedWithdrawalRootsSetsIndexes[operator][stakerQueuedWithdrawalRootsSetsValues[operator][index]]) == index + 1)
    && (forall address operator . forall bytes32 value. stakerQueuedWithdrawalRootsSetsIndexes[operator][value] == 0 || 
         ((forall uint256 tmp. to_mathint(tmp) == (stakerQueuedWithdrawalRootsSetsIndexes[operator][value] - 1) => stakerQueuedWithdrawalRootsSetsValues[operator][tmp] == value)
          && stakerQueuedWithdrawalRootsSetsIndexes[operator][value] >= 1 && stakerQueuedWithdrawalRootsSetsIndexes[operator][value] <= stakerQueuedWithdrawalRootsGhostLengths[operator]));

/// @title _cumulativeScaledSharesHistory.keys are monotonic increasing
/// @property CumulativeScaledSharesHistory Keys are valid - Snapshots' timestamps are sorted in increasing order.
// _cumulativeScaledSharesHistory.key are the block number at the update time, should be monotonic increasing
invariant cumulativeScaledSharesHistoryKeysMonotonicInc()
    forall address operator . forall address strategy . forall uint256 index1 . forall uint256 index2 .
        index1 < cumulativeScaledSharesHistoryLengths[operator][strategy] && index2 <= index1 =>
        cumulativeScaledSharesHistoryKeys[operator][strategy][index2] <= cumulativeScaledSharesHistoryKeys[operator][strategy][index1];

// Ensures that entries beyond `cumulativeScaledSharesHistoryLengths` are nullified:
// 1. For each operator and strategy, all keys and values at indices >= `cumulativeScaledSharesHistoryLengths` must be zero.
invariant cumulativeScaledSharesHistoryPastLengthNullified()
    forall address operator . forall address strategy . forall uint256 index . index >= cumulativeScaledSharesHistoryLengths[operator][strategy] =>
        (cumulativeScaledSharesHistoryKeys[operator][strategy][index] == 0 && cumulativeScaledSharesHistoryValues[operator][strategy][index] == 0);

/// @title  Ensures all keys in `cumulativeScaledSharesHistoryKeys` are less than or equal to the current block number.
/// @property CumulativeScaledSharesHistory Keys are valid
invariant CumulativeScaledSharesHistoryKeysLessThanCurrentBlock(env e)
    forall address operator . forall address strategy . forall uint256 index . index < cumulativeScaledSharesHistoryLengths[operator][strategy] =>
        cumulativeScaledSharesHistoryKeys[operator][strategy][index] <= e.block.number 
        {
            preserved with (env e1) {
                require e1.block.number == e.block.number;
            }
        }

/// @title CumulativeScaledShares is monotonically increasing
/// @property CumulativeScaledShares is monotonically increasing
invariant cumulativeScaledSharesMonotonicallyIncreasing()
    forall address operator . forall address strategy . forall uint256 index1 . forall uint256 index2 . (index1 >= index2 =>
        cumulativeScaledSharesHistoryValues[operator][strategy][index1] >= cumulativeScaledSharesHistoryValues[operator][strategy][index2])
    {
        preserved with (env e) {
            requireInvariant cumulativeScaledSharesHistoryKeysMonotonicInc();
            requireInvariant cumulativeScaledSharesHistoryPastLengthNullified();
            requireInvariant CumulativeScaledSharesHistoryKeysLessThanCurrentBlock(e);
        }
    }

/// @title
/*
For a given staker, all roots existing in _stakerQueuedWithdrawalRoots also exists in storage:
_queuedWithdrawals[withdrawalRoot] returns the Withdrawal struct / staker != address(0)
pendingWithdrawals[withdrawalRoot] returns True
*/
/// @property Queued Withdrawal Registration Consistency Invariant
invariant queuedWithdrawalsCorrect()
    forall address staker . forall bytes32 root . inRootsSets(staker, root) => 
                pendingWithdrawalsMirror[root] && queuedWithdrawalsStakesrGhost[root] != 0
    {
        preserved {
            requireInvariant stakerQueuedWithdrawalRootsInvariant();
        }
    }

/// @title Verifies that operators delegates to themselves.
/// @property Operators delegates to themselves
invariant operatorDelegatesToThemselves()
    forall address operator . operatorDetailsDelegationApproversMirror[operator] != 0 => DelegationManager.delegatedTo[operator] == operator;


/// @title Operator Cannot Deregister
/// @property Operator Cannot Deregister
rule operatorCannotDeregister(method f, address account)
filtered{f -> !f.isView}
{
    env e;
    calldataarg args;

    require account != 0;

    address delegatedToPre = delegatedTo(account);

        f(e, args);

    address delegatedToPost = delegatedTo(account);    

    assert delegatedToPre == account => delegatedToPost == account;
}

/// @title depositScalingFactor is initialized as WAD and only increases on further deposits.
/// @property depositScalingFactor Monotonic Increasing
rule depositScalingFactorMonotonicInc(method f, address staker, address strategy) 
filtered{f -> !f.isView}
{
    env e;
    calldataarg args;

    uint256 depositScalingFactorPre = scalingFactorsGhost[staker][strategy];
    require depositScalingFactorPre >= WAD();

        f(e, args);

    uint256 depositScalingFactorPost = scalingFactorsGhost[staker][strategy];

    assert depositScalingFactorPost >= depositScalingFactorPre;
}

// /// @title For a given (Staker, Strategy), withdrawableShares <= depositShares.
// /// @property Withdrawable Shares Cannot Exceed Deposited Shares.
// rule withdrawableSharesGrEqDepositShares(method f, address staker, address strategy) {
//     env e;
//     calldataarg args;

//     address[] strategies = [strategy];
//     require strategy != 0;
//     require scalingFactorsGhost[staker][strategy] >=  WAD(); // safe assumption - was proved

//     uint256[] withdrawableSharesPre; 
//     uint256[] depositSharesPre;
//     uint256[] withdrawableSharesPost; 
//     uint256[] depositSharesPost;

//     withdrawableSharesPre, depositSharesPre = getWithdrawableShares(e, staker, strategies);

//     require withdrawableSharesPre[0] <= depositSharesPre[0];

//     f(e,args);

//     withdrawableSharesPost, depositSharesPost = getWithdrawableShares(e, staker, strategies);

//     assert withdrawableSharesPost[0] <= depositSharesPost[0];
// }

/// @title
/*
depositScalingFactor gets updated for a staker on every added deposit even if the staker is not delegated. Added shares deposits occur through
1. increaseDelegatedShares called by deposits through the StrategyManager/EigenPodManager
2. delegateTo
3. completeQueuedWithdrawals as shares
*/
/// @property Authorized Methods for Deposit Scaling Factor Modification Invariant
rule whoCanChangeDepositScalingFactor(method f, address staker, address strategy) 
filtered{f -> !f.isView}
{
    env e;
    calldataarg args;

    uint256 depositScalingFactorPre = scalingFactorsGhost[staker][strategy];

        f(e, args);

    uint256 depositScalingFactorPost = scalingFactorsGhost[staker][strategy];

    assert depositScalingFactorPost != depositScalingFactorPre =>
        f.selector == sig:increaseDelegatedShares(address, address, uint256, uint256).selector ||
        f.selector == sig:delegateTo(address, ISignatureUtilsMixinTypes.SignatureWithExpiry, bytes32).selector ||
        f.selector == sig:completeQueuedWithdrawals(IDelegationManagerTypes.Withdrawal[], address[][], bool[]).selector;
}
