// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

contract EmptyContract {
    uint256 public stateVar;

    function increment() public returns (uint256) {
        stateVar++;
        return stateVar;
    }
}