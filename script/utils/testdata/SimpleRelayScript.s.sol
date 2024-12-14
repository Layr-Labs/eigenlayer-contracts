// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../../../src/contracts/core/DelegationManager.sol";

import "./utils/RelayScript.s.sol";

contract SimpleRelayScript is RelayScript {
    Relay internal operator;
    
    function run() public relayBroadcast {
        _broadcastEOA(msg.sender);
        // Deploy a new Relay
        operator = new Relay();

        _broadcastRelay(address(operator));
        DelegationManager(0x75dfE5B44C2E530568001400D3f704bC8AE350CC).registerAsOperator(
            address(0),
            0,
            "https://example.com"
        );
    }
} 