
methods {
    //// External Calls
	// external calls to DelegationManager 
    function _.undelegate(address) external;
    function _.decreaseDelegatedShares(address,address,uint256) external;
	function _.increaseDelegatedShares(address,address,uint256) external;

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
	
    // TODO: envfree functions

    // harnessed functions
    function get_podOwnerShares(address) external returns (int256) envfree;
}

invariant podOwnerSharesAlwaysWholeGweiAmount(address podOwner)
    get_podOwnerShares(podOwner) % 1000000000 == 0;