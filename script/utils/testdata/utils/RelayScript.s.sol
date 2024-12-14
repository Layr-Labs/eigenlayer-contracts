// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "forge-std/Script.sol";
import "forge-std/Vm.sol";

import "./Relay.sol";

abstract contract RelayScript is Script {
    bool private started;
    mapping(address => bool) private isRelay;
    mapping(address => bool) private isEOA;
    uint256 internal broadcastForkId;

    bool internal everStarted;
    uint256 internal relayForkId;

    modifier relayBroadcast {
        _start();
        _;
        _flushBroadcasts();
    }

    function _start() internal {
        require(!started, "RelayScript._start: already started");

        // create a fork of the current state to broadcast on later
        broadcastForkId = vm.activeFork();

        // create a fork of the current state to prank on now
        if (!everStarted) {
            relayForkId = vm.createFork(vm.envString("ETH_RPC_URL"));
        }
        vm.selectFork(relayForkId);
        vm.startStateDiffRecording();

        started = true;
        everStarted = true;
    }

    function _flushQueuedBroadcasts() internal {
        // stop broadcasting
        (Vm.CallerMode callerMode, address caller, ) = vm.readCallers();
        // TODO: why doesn't Vm.CallerMode.NONE work here?
        // if we are in the middle of a broadcast, stop it
        if (uint8(callerMode) != uint8(0)) {
            _stopBroadcast();
        }

        // flush the queued broadcasts
        _flushBroadcasts();
        // restart relaying
        _start();
        
        // resume broadcasting
        if (uint8(callerMode) != uint8(0)) {
            if (isRelay[caller]) {
                _startBroadcastRelay(caller);
            } else if (isEOA[caller]) {
                _startBroadcastEOA(caller);
            }
        }
    }

    function _broadcastRelay(address relay) internal {
        require(started, "RelayScript._broadcastRelay: not started");
        isRelay[relay] = true;
        vm.prank(relay);
    }

    function _broadcastEOA(address eoa) internal {
        require(started, "RelayScript._broadcastEOA: not started");
        isEOA[eoa] = true;
        vm.prank(eoa);
    }

    function _startBroadcastRelay(address relay) internal {
        require(started, "RelayScript._broadcastRelay: not started");
        isRelay[relay] = true;
        vm.startPrank(relay);
    }
    
    function _startBroadcastEOA(address eoa) internal {
        require(started, "RelayScript._broadcastEOA: not started");
        isEOA[eoa] = true;
        vm.startPrank(eoa);
    }

    function _stopBroadcast() internal {
        vm.stopPrank();
    }

    function _flushBroadcasts() internal {
        require(started, "RelayScript._broadcast: not started");
        started = false;
        
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
