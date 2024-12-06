// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigCall, MultisigCallUtils} from "zeus-templates/templates/MultisigBuilder.sol";
import {SafeTx, SafeTxUtils} from "zeus-templates/utils/SafeTxUtils.sol";
import {Queue} from "./2-queueUpgradeAndUnpause.s.sol";
import {EigenLabsUpgrade} from "../EigenLabsUpgrade.s.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract Execute is Queue {
    using MultisigCallUtils for MultisigCall[];
    using SafeTxUtils for SafeTx;
    using EigenLabsUpgrade for *;

    /**
     * @dev Overrides the previous _execute function to execute the queued transactions.
     */
    function _runAsMultisig() internal override {
        bytes memory call = _getMultisigTransactionCalldata();
        TimelockController timelock = TimelockController(payable(this._timelock()));
        timelock.execute(
            this._executorMultisig(),
            0,
            call,
            0,
            bytes32(0)
        );
    }

    function testDeploy() override public {}

    function testExecute() public { 
        // 1- run queueing logic
        vm.startPrank(this._operationsMultisig());
        super._runAsMultisig();
        vm.stopPrank();

        TimelockController timelock = this._timelock();
        bytes memory call = _getMultisigTransactionCalldata();
        bytes32 txHash = timelock.hashOperation(this._executorMultisig(), 0, call, 0, 0);
        assertEq(timelock.isOperationPending(txHash), true, "Transaction should be queued and pending.");


        // 2- warp past delay?
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA
        assertEq(timelock.isOperationReady(txHash), true, "Transaction should be executable.");
        
        // 3- execute
        zSetMultisigContext(this._protocolCouncilMultisig());
        execute();

        // 3. TODO: assert that the execute did something
    }
}
