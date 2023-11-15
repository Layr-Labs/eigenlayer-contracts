// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/test/integration/GlobalRefs.sol";

contract Staker {
    GlobalRefs globalRefs;

    constructor(GlobalRefs _globalRefs) {
        globalRefs = _globalRefs;
    }

    function depositIntoStrategy() public {}

    function delegate() public {}

    function queueWithdrawal() public {}

    function completeQueuedWithdrawal() public {} 
}