// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../contracts/permissions/Pausable.sol";

// wrapper around the Pausable contract that exposes the internal `_setPausedStatus` function.
contract PausableHarness is Pausable {
    constructor(IPauserRegistry _pauserRegistry) Pausable(_pauserRegistry) {}

    function initializePauser(uint initPausedStatus) external {
        _setPausedStatus(initPausedStatus);
    }
}
