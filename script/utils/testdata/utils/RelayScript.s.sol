// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "forge-std/Script.sol";
import "forge-std/Vm.sol";

import "./Relay.sol";

abstract contract RelayScript is Script {
    bool private started;
    bool private broadcasted;
    mapping(address => bool) private isRelay;

    uint256 internal broadcastForkId;

    modifier relayBroadcast {
        _start();
        _;
        _broadcast();
    }

    function _start() internal {
        require(!started, "RelayScript._start: already started");
        started = true;

        // create a fork of the current state to broadcast on later
        broadcastForkId = vm.activeFork();

        // create a fork of the current state to prank on now
        vm.createSelectFork(vm.envString("ETH_RPC_URL"));
        vm.startStateDiffRecording();
    }

    function _prank(address relay) internal {
        require(started, "RelayScript._prank: already started");
        vm.prank(relay);
        isRelay[relay] = true;
    }

    function _broadcast() internal {
        require(started, "RelayScript._broadcast: not started");
        require(!broadcasted, "RelayScript._broadcast: already broadcasted");
        broadcasted = true;
        
        Vm.AccountAccess[] memory accesses = vm.stopAndReturnStateDiff();

        // broadcast relay txs on broadcast fork
        vm.selectFork(broadcastForkId);
        vm.startBroadcast();
        // Process each access to find relay calls
        for (uint256 i = 0; i < accesses.length; i++) {
            Vm.AccountAccess memory access = accesses[i];
            
            // Only process Call type accesses
            // TODO: why doesn't Vm.AccountAccessKind.Call work here?
            if (uint8(access.kind) != uint8(0)) continue;
            
            // Check if the accessor is one of our relays
            if (!isRelay[access.accessor]) continue;

            // Forward the call through the Pod interface
            Relay(payable(access.accessor)).execute(
                access.account,
                access.value,
                access.data
            );
        }
    }
}
