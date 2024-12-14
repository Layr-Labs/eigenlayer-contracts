// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/access/Ownable.sol";

import "./solady/Pod.sol";

/**
 * @title Relay
 * @dev A contract that combines Pod functionality with Ownable, where ownership controls
 * override the mothership-related functions
 */
contract Relay is Pod, Ownable {
    constructor() Pod() Ownable() {}

    /**
     * @dev Override the Pod's mothership function to return the owner instead
     * @return address The owner's address
     */
    function mothership() public view override returns (address) {
        return owner();
    }
}
