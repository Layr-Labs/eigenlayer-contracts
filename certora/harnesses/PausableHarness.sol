// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

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
    function bitwise_not(uint256 input) external pure returns (uint256) {
        return (~input);
    }

    function bitwise_and(uint256 input_1, uint256 input_2) external pure returns (uint256) {
        return (input_1 & input_2);
    }
}
