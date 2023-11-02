
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

    // harnessed functions
    function get_validatorIndex(bytes32 pubkeyHash) external returns (uint64) envfree;
    function get_restakedBalanceGwei(bytes32 pubkeyHash) external returns (uint64) envfree;
    function get_mostRecentBalanceUpdateTimestamp(bytes32 pubkeyHash) external returns (uint64) envfree;
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
    method f;
    env e;
    calldataarg args;
    f(e,args);
    uint64 validatorIndexAfter = get_validatorIndex(pubkeyHash);
    assert(validatorIndexBefore == 0 || validatorIndexAfter == validatorIndexBefore,
        "validator index modified from nonzero value");
}