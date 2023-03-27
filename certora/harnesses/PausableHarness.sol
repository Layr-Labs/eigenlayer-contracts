// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../munged/permissions/Pausable.sol";

contract PausableHarness is Pausable {
    // getters
    function pauser() external view returns (address) {
        return pauserRegistry.pauser();
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