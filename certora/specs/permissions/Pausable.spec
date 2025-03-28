
methods {
    // envfree functions
    function paused() external returns (uint256) envfree;
    function paused(uint8 index) external returns (bool) envfree;
    function pauserRegistry() external returns (address) envfree;

    // harnessed functions
    function isPauser(address) external returns (bool) envfree;
    function unpauser() external returns (address) envfree;
    function bitwise_not(uint256) external returns (uint256) envfree;
    function bitwise_and(uint256, uint256) external returns (uint256) envfree;
}

rule onlyPauserCanPauseAndOnlyUnpauserCanUnpause() {
    method f;
    env e;
    uint256 pausedStatusBefore = paused();
    address unpauser = unpauser();
    if (f.selector == sig:pause(uint256).selector) {
        uint256 newPausedStatus;
        pause(e, newPausedStatus);
        uint256 pausedStatusAfter = paused();
        if (isPauser(e.msg.sender) && bitwise_and(pausedStatusBefore, newPausedStatus) == pausedStatusBefore) {
            assert(pausedStatusAfter == newPausedStatus, "pausedStatusAfter != newPausedStatus");
        } else {
            assert(pausedStatusAfter == pausedStatusBefore, "pausedStatusAfter != pausedStatusBefore");
        }
    } else if (f.selector == sig:pauseAll().selector) {
        pauseAll(e);
        uint256 pausedStatusAfter = paused();
       if (isPauser(e.msg.sender)) {
            // assert(pausedStatusAfter == type(uint256).max, "pausedStatusAfter != newPausedStatus");
            assert(pausedStatusAfter == 115792089237316195423570985008687907853269984665640564039457584007913129639935,
                "pausedStatusAfter != newPausedStatus");
        } else {
            assert(pausedStatusAfter == pausedStatusBefore, "pausedStatusAfter != pausedStatusBefore");
        }
    } else if (f.selector == sig:unpause(uint256).selector) {
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