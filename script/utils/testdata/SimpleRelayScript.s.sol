// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../../../src/contracts/core/DelegationManager.sol";

import "./utils/RelayScript.s.sol";

contract SimpleRelayScript is RelayScript {
    Relay internal relay;
    
    function run() public {
        vm.broadcast();
        // Deploy a new Relay
        relay = new Relay();

        _run();
    }

    function _run() internal relayBroadcast {
        _prank(address(relay));
        DelegationManager(0x75dfE5B44C2E530568001400D3f704bC8AE350CC).registerAsOperator(
            address(0),
            0,
            "https://example.com"
        );
    }


} 