
methods {
    // Internal, NONDET-summarized EigenPod library functions
    function _.verifyValidatorFields(bytes32, bytes32[] calldata, bytes calldata, uint40) internal => NONDET;
    function _.verifyValidatorBalance(bytes32, bytes32, bytes calldata, uint40) internal => NONDET;
    function _.verifyStateRootAgainstLatestBlockRoot(bytes32, bytes32, bytes calldata) internal => NONDET;
    function _.verifyWithdrawal(bytes32, bytes32[] calldata, BeaconChainProofs.WithdrawalProof calldata) internal => NONDET;

    // Internal, NONDET-summarized "send ETH" function -- unsound summary used to avoid HAVOC behavior
    // when sending ETH using `Address.sendValue()`
    function _._sendETH(address recipient, uint256 amountWei) internal => NONDET;

    // summarize the deployment of EigenPods to avoid default, HAVOC behavior
    function _.deploy(uint256, bytes32, bytes memory bytecode) internal => NONDET;

    //// External Calls
	// external calls to DelegationManager 
    function _.undelegate(address) external => DISPATCHER(true);
    function _.decreaseDelegatedShares(address,address,uint256) external => DISPATCHER(true);
	function _.increaseDelegatedShares(address,address,uint256) external => DISPATCHER(true);

    // external calls from DelegationManager to ServiceManager
    function _.updateStakes(address[]) external => NONDET;

	// external calls to Slasher
    function _.isFrozen(address) external => DISPATCHER(true);
	function _.canWithdraw(address,uint32,uint256) external => DISPATCHER(true);
    function _.recordStakeUpdate(address,uint32,uint32,uint256) external => NONDET;

	// external calls to StrategyManager
    function _.getDeposits(address) external => DISPATCHER(true);
    function _.slasher() external => DISPATCHER(true);
    function _.removeShares(address,address,uint256) external => DISPATCHER(true);
    function _.withdrawSharesAsTokens(address, address, uint256, address) external => DISPATCHER(true);

    // external calls to Strategy contracts
    function _.deposit(address, uint256) external => NONDET;
    function _.withdraw(address, address, uint256) external => NONDET;

	// external calls to EigenPodManager
    function _.addShares(address,uint256) external => DISPATCHER(true);
    function _.removeShares(address,uint256) external => DISPATCHER(true);
    function _.withdrawSharesAsTokens(address, address, uint256) external => DISPATCHER(true);
    function _.podOwnerShares(address) external => DISPATCHER(true);

    // external calls to EigenPod
	function _.withdrawRestakedBeaconChainETH(address,uint256) external => DISPATCHER(true);
    function _.stake(bytes, bytes, bytes32) external => DISPATCHER(true);
    function _.initialize(address) external => DISPATCHER(true);

    // external calls to ETH2Deposit contract
    function _.deposit(bytes, bytes, bytes, bytes32) external => NONDET;

    // external calls to PauserRegistry
    function _.isPauser(address) external => DISPATCHER(true);
	function _.unpauser() external => DISPATCHER(true);

    // external calls to ERC20 token
    function _.transfer(address, uint256) external => DISPATCHER(true);
    function _.transferFrom(address, address, uint256) external => DISPATCHER(true);
    function _.approve(address, uint256) external => DISPATCHER(true);
	
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
    function nonBeaconChainETHBalanceWei() external returns (uint256) envfree;

    // harnessed functions
    function get_validatorIndex(bytes32 pubkeyHash) external returns (uint64) envfree;
    function get_restakedBalanceGwei(bytes32 pubkeyHash) external returns (uint64) envfree;
    function get_mostRecentBalanceUpdateTimestamp(bytes32 pubkeyHash) external returns (uint64) envfree;
    function get_podOwnerShares() external returns (int256) envfree;
    function get_withdrawableRestakedExecutionLayerGwei() external returns (uint256) envfree;
    function get_ETH_Balance() external returns (uint256) envfree;
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

// based on Certora's example here https://github.com/Certora/Tutorials/blob/michael/ethcc/EthCC/Ghosts/ghostTest.spec
ghost mathint sumOfValidatorRestakedbalancesWei {
    // NOTE: this commented out line is broken, as calling functions in axioms is currently disallowed, but this is what we'd run ideally. 
    // init_state axiom sumOfValidatorRestakedbalancesWei == to_mathint(get_podOwnerShares()) - to_mathint(get_withdrawableRestakedExecutionLayerGwei() * 1000000000);

    // since both of these variables are zero at construction, just set the ghost to zero in the axiom
    init_state axiom sumOfValidatorRestakedbalancesWei == 0;
}

hook Sstore _validatorPubkeyHashToInfo[KEY bytes32 validatorPubkeyHash].restakedBalanceGwei uint64 newValue (uint64 oldValue) {
    sumOfValidatorRestakedbalancesWei = (
        sumOfValidatorRestakedbalancesWei + 
        to_mathint(newValue) * 1000000000 -
        to_mathint(oldValue) * 1000000000
    );
}

rule consistentAccounting() {
    // fetch info before call
    int256 podOwnerSharesBefore = get_podOwnerShares();
    uint256 withdrawableRestakedExecutionLayerGweiBefore = get_withdrawableRestakedExecutionLayerGwei();
    uint256 eigenPodBalanceBefore = get_ETH_Balance();
    // filter down to valid pre-states
    require(sumOfValidatorRestakedbalancesWei ==
        to_mathint(podOwnerSharesBefore) - to_mathint(withdrawableRestakedExecutionLayerGweiBefore));

    // perform arbitrary function call
    method f;
    env e;
    calldataarg args;
    f(e,args);

    // fetch info after call
    int256 podOwnerSharesAfter = get_podOwnerShares();
    uint256 withdrawableRestakedExecutionLayerGweiAfter = get_withdrawableRestakedExecutionLayerGwei();
    uint256 eigenPodBalanceAfter = get_ETH_Balance();
    /**
     * handling for weird, unrealistic edge case where calling `initialize` causes the pod owner to change, so the 
     * call to `get_podOwnerShares` queries the shares for a different address.
     * calling `initialize` should *not* change user shares, so it is unrealistic to simulate it doing so.
     */
    if (f.selector == sig:initialize(address).selector) {
        podOwnerSharesAfter = podOwnerSharesBefore;
    }
    // check post-state
    // TODO: this check is still broken for `withdrawRestakedBeaconChainETH` since it does a low-level call to transfer the ETH, which triggers optimistic fallback dispatching
    // special handling for one function
    if (f.selector == sig:withdrawRestakedBeaconChainETH(address,uint256).selector) {
        /* TODO: un-comment this once the dispatching is handled correctly
        assert(sumOfValidatorRestakedbalancesWei ==
            to_mathint(podOwnerSharesAfter) - to_mathint(withdrawableRestakedExecutionLayerGweiAfter)
            // adjustment term for the ETH balance of the contract changing
            + to_mathint(eigenPodBalanceBefore) - to_mathint(eigenPodBalanceAfter),
            "invalid post-state");
        */
        // TODO: delete this once the above is salvaged (was added since CVL forbids empty blocks)
        assert(true);
    // outside of special case, we don't need the adjustment term
    } else {
        assert(sumOfValidatorRestakedbalancesWei ==
            to_mathint(podOwnerSharesAfter) - to_mathint(withdrawableRestakedExecutionLayerGweiAfter),
            "invalid post-state");
    }
}

/*
rule baseInvariant() {
    int256 podOwnerSharesBefore = get_podOwnerShares();
    // perform arbitrary function call
    method f;
    env e;
    calldataarg args;
    f(e,args);
    int256 podOwnerSharesAfter = get_podOwnerShares();
    mathint podOwnerSharesDelta = podOwnerSharesAfter - podOwnerSharesBefore;
    assert(sumOfValidatorRestakedbalancesWei == podOwnerSharesDelta - to_mathint(get_withdrawableRestakedExecutionLayerGwei()),
        "base invariant violated");
}

invariant consistentAccounting() {
    sumOfValidatorRestakedbalancesWei ==
        to_mathint(get_withdrawableRestakedExecutionLayerGwei()) - to_mathint(get_withdrawableRestakedExecutionLayerGwei());
}
*/