// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {QueueAndUnpause} from "./2-queueUpgradeAndUnpause.s.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract Execute is QueueAndUnpause {
    using Env for *;

    function _runAsMultisig() prank(Env.protocolCouncilMultisig()) internal override {
        bytes memory call = _getCalldataToExecutor();
        TimelockController timelock = Env.timelockController();
        timelock.execute(
            Env.executorMultisig(),
            0,
            call,
            0,
            bytes32(0)
        );
    }

    function testDeploy() override public {}

    function testExecute() public { 
        // 1- run queueing logic
        vm.startPrank(Env.opsMultisig());
        super._runAsMultisig();
        vm.stopPrank();

        TimelockController timelock = Env.timelockController();
        bytes memory call = _getCalldataToExecutor();
        bytes32 txHash = timelock.hashOperation(Env.executorMultisig(), 0, call, 0, 0);
        assertEq(timelock.isOperationPending(txHash), true, "Transaction should be queued and pending.");


        // 2- warp past delay?
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA
        assertEq(timelock.isOperationReady(txHash), true, "Transaction should be executable.");
        
        // 3- execute
        execute();

        // 3. TODO: assert that the execute did something
    }
}
