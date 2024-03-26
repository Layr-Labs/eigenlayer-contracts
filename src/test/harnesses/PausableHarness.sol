// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../contracts/permissions/Pausable.sol";

// wrapper around the Pausable contract that exposes the internal `_initializePauser` function.
contract PausableHarness is Pausable {
    function initializePauser(IPauserRegistry _pauserRegistry, uint256 initPausedStatus) external {
        _initializePauser(_pauserRegistry, initPausedStatus);
    }
}
