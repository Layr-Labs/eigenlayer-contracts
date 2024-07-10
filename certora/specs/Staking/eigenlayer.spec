methods {
    // Strategy manager
    function _.withdrawalRootPending(bytes32 _withdrawalRoot) external => NONDET; // expect (bool);
    function _.numWithdrawalsQueued(address _user) external => NONDET; // expect (uint96);
    function _.pauserRegistry() external => HAVOC_ECF; // expect (IPauserRegistry);
    function _.paused(uint8 index) external => NONDET; // expect (bool);
    function _.unpause(uint256 newPausedStatus) external => HAVOC_ECF; // expect void

    // interface IEigenPod
    function _.withdrawBeforeRestaking() external => HAVOC_ECF; // expect void
    
    // interface IEigenPodManager
    function _.getPod(address podOwner) external => NONDET; // expect address; // (IEigenPod)
    function _.createPod() external => HAVOC_ECF; // expect (address);
}