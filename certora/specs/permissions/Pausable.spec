
methods {
    // external calls to PauserRegistry
    isPauser() returns (bool) => DISPATCHER(true)
	unpauser() returns (address) => DISPATCHER(true)

    // envfree functions
    paused() returns (uint256) envfree
    paused(uint8 index) returns (bool) envfree
    pauserRegistry() returns (address) envfree

    // harnessed functions
    pauser() returns (address) envfree
    unpauser() returns (address) envfree
    bitwise_not(uint256) returns (uint256) envfree
    bitwise_and(uint256, uint256) returns (uint256) envfree
}

rule onlyPauserCanPauseAndOnlyUnpauserCanUnpause() {
    method f;
    env e;
    uint256 pausedStatusBefore = paused();
    address pauser;
    require(isPauser(pauser));
    address unpauser = unpauser();
    if (f.selector == pause(uint256).selector) {
        uint256 newPausedStatus;
        pause(e, newPausedStatus);
        uint256 pausedStatusAfter = paused();
        if (e.msg.sender == pauser && bitwise_and(pausedStatusBefore, newPausedStatus) == pausedStatusBefore) {
            assert(pausedStatusAfter == newPausedStatus, "pausedStatusAfter != newPausedStatus");
        } else {
            assert(pausedStatusAfter == pausedStatusBefore, "pausedStatusAfter != pausedStatusBefore");
        }
    } else if (f.selector == pauseAll().selector) {
        pauseAll(e);
        uint256 pausedStatusAfter = paused();
       if (e.msg.sender == pauser) {
            // assert(pausedStatusAfter == type(uint256).max, "pausedStatusAfter != newPausedStatus");
            assert(pausedStatusAfter == 115792089237316195423570985008687907853269984665640564039457584007913129639935,
                "pausedStatusAfter != newPausedStatus");
        } else {
            assert(pausedStatusAfter == pausedStatusBefore, "pausedStatusAfter != pausedStatusBefore");
        }
    } else if (f.selector == unpause(uint256).selector) {
        uint256 newPausedStatus;
        unpause(e, newPausedStatus);
        uint256 pausedStatusAfter = paused();
        if (e.msg.sender == unpauser && bitwise_and(bitwise_not(pausedStatusBefore), bitwise_not(newPausedStatus)) == bitwise_not(pausedStatusBefore)) {
            assert(pausedStatusAfter == newPausedStatus, "pausedStatusAfter != newPausedStatus");
        } else {
            assert(pausedStatusAfter == pausedStatusBefore, "pausedStatusAfter != pausedStatusBefore");
        }
    } else {
        calldataarg arg;
        f(e,arg);
        uint256 pausedStatusAfter = paused();
        assert(pausedStatusAfter == pausedStatusBefore, "pausedStatusAfter != pausedStatusBefore");
    }
}