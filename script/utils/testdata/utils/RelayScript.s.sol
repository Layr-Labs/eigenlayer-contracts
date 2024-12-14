// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "forge-std/Script.sol";
import "forge-std/Vm.sol";

import "./Relay.sol";

abstract contract RelayScript is Script {
    bool private started;
    bool private broadcasted;
    mapping(address => bool) private isRelay;
    mapping(address => bool) private isEOA;
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

    function _broadcastRelay(address relay) internal {
        require(started, "RelayScript._broadcastRelay: not started");
        vm.prank(relay);
        isRelay[relay] = true;
    }

    function _broadcastEOA(address eoa) internal {
        require(started, "RelayScript._broadcastEOA: not started");
        vm.prank(eoa);
        isEOA[eoa] = true;
    }

    function _broadcast() internal {
        require(started, "RelayScript._broadcast: not started");
        require(!broadcasted, "RelayScript._broadcast: already broadcasted");
        broadcasted = true;
        
        Vm.AccountAccess[] memory accesses = vm.stopAndReturnStateDiff();

        // broadcast relay txs on broadcast fork
        vm.selectFork(broadcastForkId);
        // Process each access to find relay calls
        for (uint256 i = 0; i < accesses.length; i++) {
            Vm.AccountAccess memory access = accesses[i];
            
            // Check if the accessor is one of our relays
            if (isRelay[access.accessor] && uint8(access.kind) == uint8(0)) {
                Relay relay = Relay(payable(access.accessor));
                vm.broadcast(relay.mothership());
                // Forward the call through the Pod interface
                relay.execute(
                    access.account,
                    access.value,
                    access.data
                );
            } else if (isEOA[access.accessor] && uint8(access.kind) == uint8(0)) {
                vm.broadcast(access.accessor);
                (bool success, ) = access.account.call{value: access.value}(access.data);
                require(success, "RelayScript._broadcast: EOA call failed");
            } else if (isEOA[access.accessor] && uint8(access.kind) == uint8(4)) {
                vm.broadcast(access.accessor);
                bytes memory data = access.data;
                uint256 value = access.value;
                assembly {
                    let x := create(value, add(data, 0x20), mload(data))
                }
            }
        }
    }
}
