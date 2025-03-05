// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/contracts/permissions/Pausable.sol";

contract PausableHarness is Pausable {
    // getters
    function isPauser(address pauser) external view returns (bool) {
        return pauserRegistry.isPauser(pauser);
    }

    function unpauser() external view returns (address) {
        return pauserRegistry.unpauser();
    }

    // bitwise operations

}
