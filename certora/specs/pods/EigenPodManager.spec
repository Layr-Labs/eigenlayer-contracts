
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
    function _.removeShares(address,address,uint256) external => DISPATCHER(true);
    function _.withdrawSharesAsTokens(address, address, uint256, address) external => DISPATCHER(true);

	// external calls to EigenPodManager
    function _.addShares(address,uint256) external => DISPATCHER(true);
    function _.removeShares(address,uint256) external => DISPATCHER(true);
    function _.withdrawSharesAsTokens(address, address, uint256) external => DISPATCHER(true);

    // external calls to EigenPod
	function _.withdrawRestakedBeaconChainETH(address,uint256) external => DISPATCHER(true);
    // external calls to PauserRegistry
    function _.isPauser(address) external => DISPATCHER(true);
	function _.unpauser() external => DISPATCHER(true);
	
    // envfree functions
    function ownerToPod(address podOwner) external returns (address) envfree;
    function getPod(address podOwner) external returns (address) envfree;
    function ethPOS() external returns (address) envfree;
    function eigenPodBeacon() external returns (address) envfree;
    function getBlockRootAtTimestamp(uint64 timestamp) external returns (bytes32) envfree;
    function strategyManager() external returns (address) envfree;
    function slasher() external returns (address) envfree;
    function hasPod(address podOwner) external returns (bool) envfree;
    function numPods() external returns (uint256) envfree;
    function podOwnerShares(address podOwner) external returns (int256) envfree;
    function beaconChainETHStrategy() external returns (address) envfree;

    // harnessed functions
    function get_podOwnerShares(address) external returns (int256) envfree;
    function get_podByOwner(address) external returns (address) envfree;
}

// verifies that podOwnerShares[podOwner] is never a non-whole Gwei amount
invariant podOwnerSharesAlwaysWholeGweiAmount(address podOwner)
    get_podOwnerShares(podOwner) % 1000000000 == 0;

// verifies that ownerToPod[podOwner] is set once (when podOwner deploys a pod), and can otherwise never be updated
rule podAddressNeverChanges(address podOwner) {
    address podAddressBefore = get_podByOwner(podOwner);
    // perform arbitrary function call
    method f;
    env e;
    calldataarg args;
    f(e,args);
    address podAddressAfter = get_podByOwner(podOwner);
    assert(podAddressBefore == 0 || podAddressBefore == podAddressAfter,
        "pod address changed after being set!");
}

// verifies that podOwnerShares[podOwner] can become negative (i.e. go from zero/positive to negative)
// ONLY as a result of a call to `recordBeaconChainETHBalanceUpdate`
rule limitationOnNegativeShares(address podOwner) {
    int256 podOwnerSharesBefore = get_podOwnerShares(podOwner);
    // perform arbitrary function call
    method f;
    env e;
    calldataarg args;
    f(e,args);
    int256 podOwnerSharesAfter = get_podOwnerShares(podOwner);
    if (podOwnerSharesAfter < 0) {
        if (podOwnerSharesBefore >= 0) {
            assert(f.selector == sig:recordBeaconChainETHBalanceUpdate(address, int256).selector,
                "pod owner shares became negative from calling an unqualified function!");
        } else {
            assert(
                (podOwnerSharesAfter >= podOwnerSharesBefore) ||
                (f.selector == sig:recordBeaconChainETHBalanceUpdate(address, int256).selector),
                "pod owner had negative shares decrease inappropriately"
            );
        }
    }
    // need this line to keep the prover happy :upside_down_face:
    assert(true);
}
