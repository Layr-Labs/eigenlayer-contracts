// summaries for unresolved calls
methods {
    // IERC1271
    function _.isValidSignature(bytes32, bytes) external => DISPATCHER(true); 

    // IEigenPod
    function _.withdrawRestakedBeaconChainETH(address, uint256) external => DISPATCHER(true);

    // IStrategy
    function _.withdraw(address, address, uint256) external => DISPATCHER(true);
    function _.deposit(address, uint256) external => DISPATCHER(true);

    // PauserRegistry
    function _.isPauser(address) external => DISPATCHER(true);
    function _.unpauser() external  => DISPATCHER(true);

    // IETHPOSDeposit
    function _.deposit(bytes, bytes, bytes, bytes32) external => DISPATCHER(true);
}