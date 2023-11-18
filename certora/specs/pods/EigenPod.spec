
methods {
    //// External Calls
	// external calls to DelegationManager 
    function _.undelegate(address) external;
    function _.decreaseDelegatedShares(address,address,uint256) external;
	function _.increaseDelegatedShares(address,address,uint256) external;

    // external calls from DelegationManager to ServiceManager
    function _.updateStakes(address[]) external => NONDET;

	// external calls to Slasher
    function _.isFrozen(address) external => DISPATCHER(true);
	function _.canWithdraw(address,uint32,uint256) external => DISPATCHER(true);

	// external calls to StrategyManager
    function _.getDeposits(address) external => DISPATCHER(true);
    function _.slasher() external => DISPATCHER(true);
    function _.addShares(address,address,uint256) external => DISPATCHER(true);
    function _.removeShares(address,address,uint256) external => DISPATCHER(true);
    function _.withdrawSharesAsTokens(address, address, uint256, address) external => DISPATCHER(true);

	// external calls to EigenPodManager
    function _.addShares(address,uint256) external => DISPATCHER(true);
    function _.removeShares(address,uint256) external => DISPATCHER(true);
    function _.withdrawSharesAsTokens(address, address, uint256) external => DISPATCHER(true);

    // external calls to EigenPod
	function _.withdrawRestakedBeaconChainETH(address,uint256) external => DISPATCHER(true);
    
    // external calls to DelayedWithdrawalRouter (from EigenPod)
    function _.createDelayedWithdrawal(address, address) external => DISPATCHER(true);

    // external calls to PauserRegistry
    function _.isPauser(address) external => DISPATCHER(true);
	function _.unpauser() external => DISPATCHER(true);
	
    // envfree functions
    function MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR() external returns (uint64) envfree;
    function withdrawableRestakedExecutionLayerGwei() external returns (uint64) envfree;
    function nonBeaconChainETHBalanceWei() external returns (uint256) envfree;
    function eigenPodManager() external returns (address) envfree;
    function podOwner() external returns (address) envfree;
    function hasRestaked() external returns (bool) envfree;
    function mostRecentWithdrawalTimestamp() external returns (uint64) envfree;
    function validatorPubkeyHashToInfo(bytes32 validatorPubkeyHash) external returns (IEigenPod.ValidatorInfo) envfree;
    function provenWithdrawal(bytes32 validatorPubkeyHash, uint64 slot) external returns (bool) envfree;
    function validatorStatus(bytes32 pubkeyHash) external returns (IEigenPod.VALIDATOR_STATUS) envfree;
    function delayedWithdrawalRouter() external returns (address) envfree;
    function nonBeaconChainETHBalanceWei() external returns (uint256) envfree;

    // harnessed functions
    function get_validatorIndex(bytes32 pubkeyHash) external returns (uint64) envfree;
    function get_restakedBalanceGwei(bytes32 pubkeyHash) external returns (uint64) envfree;
    function get_mostRecentBalanceUpdateTimestamp(bytes32 pubkeyHash) external returns (uint64) envfree;
    function get_podOwnerShares() external returns (int256) envfree;
    function get_withdrawableRestakedExecutionLayerGwei() external returns (uint256) envfree;
}

// defines the allowed validator status transitions
definition validatorStatusTransitionAllowed(IEigenPod.VALIDATOR_STATUS statusBefore, IEigenPod.VALIDATOR_STATUS statusAfter) returns bool =
    (statusBefore == IEigenPod.VALIDATOR_STATUS.INACTIVE && statusAfter == IEigenPod.VALIDATOR_STATUS.ACTIVE)
    || (statusBefore == IEigenPod.VALIDATOR_STATUS.ACTIVE && statusAfter == IEigenPod.VALIDATOR_STATUS.WITHDRAWN);

// verifies that only the 2 allowed transitions of validator status occur
rule validatorStatusTransitionsCorrect(bytes32 pubkeyHash) {
    IEigenPod.VALIDATOR_STATUS statusBefore = validatorStatus(pubkeyHash);
    method f;
    env e;
    calldataarg args;
    f(e,args);
    IEigenPod.VALIDATOR_STATUS statusAfter = validatorStatus(pubkeyHash);
    assert(
        (statusBefore == statusAfter)
        || validatorStatusTransitionAllowed(statusBefore, statusAfter),
        "disallowed validator status transition occurred"
    );
}

// verifies that _validatorPubkeyHashToInfo[validatorPubkeyHash].mostRecentBalanceUpdateTimestamp can ONLY increase (or remain the same)
rule mostRecentBalanceUpdateTimestampOnlyIncreases(bytes32 validatorPubkeyHash) {
    IEigenPod.ValidatorInfo validatorInfoBefore = validatorPubkeyHashToInfo(validatorPubkeyHash);
    method f;
    env e;
    calldataarg args;
    f(e,args);
    IEigenPod.ValidatorInfo validatorInfoAfter = validatorPubkeyHashToInfo(validatorPubkeyHash);
    assert(validatorInfoAfter.mostRecentBalanceUpdateTimestamp >= validatorInfoBefore.mostRecentBalanceUpdateTimestamp,
        "mostRecentBalanceUpdateTimestamp decreased");
}

// verifies that if a validator is marked as 'INACTIVE', then it has no other entries set in its ValidatorInfo
invariant inactiveValidatorsHaveEmptyInfo(bytes32 pubkeyHash)
    (validatorStatus(pubkeyHash) == IEigenPod.VALIDATOR_STATUS.INACTIVE) => (
        get_validatorIndex(pubkeyHash) == 0
        && get_restakedBalanceGwei(pubkeyHash) == 0
        && get_mostRecentBalanceUpdateTimestamp(pubkeyHash) == 0);

// verifies that _validatorPubkeyHashToInfo[validatorPubkeyHash].validatorIndex can be set initially but otherwise can't change
// this can be understood as the only allowed transitions of index being of the form: 0 => anything (otherwise the index must stay the same)
rule validatorIndexSetOnlyOnce(bytes32 pubkeyHash) {
    requireInvariant inactiveValidatorsHaveEmptyInfo(pubkeyHash);
    uint64 validatorIndexBefore = get_validatorIndex(pubkeyHash);
    // perform arbitrary function call
    method f;
    env e;
    calldataarg args;
    f(e,args);
    uint64 validatorIndexAfter = get_validatorIndex(pubkeyHash);
    assert(validatorIndexBefore == 0 || validatorIndexAfter == validatorIndexBefore,
        "validator index modified from nonzero value");
}

// verifies that once a validator has its status set to WITHDRAWN, its ‘restakedBalanceGwei’ is *and always remains* zero
invariant withdrawnValidatorsHaveZeroRestakedGwei(bytes32 pubkeyHash)
    (validatorStatus(pubkeyHash) == IEigenPod.VALIDATOR_STATUS.INACTIVE) =>
        (get_restakedBalanceGwei(pubkeyHash) == 0);

// // TODO: see if this draft rule can be salvaged
// // draft rule to capture the following behavior (or at least most of it):
// // The core invariant that ought to be maintained across the EPM and the EPs is that
// // podOwnerShares[podOwner] + sum(sharesInQueuedWithdrawals) =
// // sum(_validatorPubkeyHashToInfo[validatorPubkeyHash].restakedBalanceGwei) + withdrawableRestakedExecutionLayerGwei

// // idea: if we ignore shares in queued withdrawals and rearrange, then we have:
// // sum(_validatorPubkeyHashToInfo[validatorPubkeyHash].restakedBalanceGwei) = 
// // EigenPodManager.podOwnerShares(podOwner) - withdrawableRestakedExecutionLayerGwei
// // we can track changes to the '_validatorPubkeyHashToInfo' mapping and check this with ghost variables

// // based on Certora's example here https://github.com/Certora/Tutorials/blob/michael/ethcc/EthCC/Ghosts/ghostTest.spec
// ghost mathint sumOfValidatorRestakedbalancesWei {
//     init_state axiom sumOfValidatorRestakedbalancesWei == to_mathint(get_podOwnerShares()) - to_mathint(get_withdrawableRestakedExecutionLayerGwei() * 1000000000);
// }

// // hook Sstore _validatorPubkeyHashToInfo[KEY bytes32 validatorPubkeyHash] IEigenPod.ValidatorInfo newValue (IEigenPod.ValidatorInfo oldValue) STORAGE {
// //     sumOfValidatorRestakedbalancesWei = (
// //         sumOfValidatorRestakedbalancesWei + 
// //         to_mathint(newValue.restakedBalanceGwei) * 1000000000 -
// //         to_mathint(oldValue.restakedBalanceGwei) * 1000000000
// //     );
// // }
// // struct ValidatorInfo {
// //     // index of the validator in the beacon chain
// //     uint64 validatorIndex;
// //     // amount of beacon chain ETH restaked on EigenLayer in gwei
// //     uint64 restakedBalanceGwei;
// //     //timestamp of the validator's most recent balance update
// //     uint64 mostRecentBalanceUpdateTimestamp;
// //     // status of the validator
// //     VALIDATOR_STATUS status;
// // }

// // NOTE: this fails with the error:
// // CRITICAL: Found errors
// // CRITICAL: [main] ERROR ALWAYS - Error in spec file (EigenPod.spec:153:1): Slot pattern EigenPodHarness._validatorPubkeyHashToInfo[KEY bytes32 validatorPubkeyHash] is not an integral type: IEigenPod.ValidatorInfo
// // CRITICAL: Encountered an error running Certora Prover:
// // CVL specification syntax and type check failed
// // it would seem that some workaround may be necessary to make this struct storage work with a 'hook'
// hook Sstore _validatorPubkeyHashToInfo[KEY bytes32 validatorPubkeyHash] uint256 newValue (uint256 oldValue) STORAGE {
//     sumOfValidatorRestakedbalancesWei = (
//         sumOfValidatorRestakedbalancesWei + 
//         // extract the restakedBalanceGwei and multiply by 1e9 to get wei
//         to_mathint((newValue << 184) >> 192) * 1000000000 -
//         to_mathint((newValue << 184) >> 192) * 1000000000
//     );
// }

// rule baseInvariant() {
//     // perform arbitrary function call
//     method f;
//     env e;
//     calldataarg args;
//     f(e,args);
//     assert(sumOfValidatorRestakedbalancesWei == get_podOwnerShares() - to_mathint(get_withdrawableRestakedExecutionLayerGwei()),
//         "base invariant violated");
// }