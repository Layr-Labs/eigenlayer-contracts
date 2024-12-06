// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigCall, MultisigCallUtils, MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {EigenLabsUpgrade} from "../EigenLabsUpgrade.s.sol";
import {EncGnosisSafe} from "zeus-templates/utils/EncGnosisSafe.sol";
import {MultisigCallUtils, MultisigCall} from "zeus-templates/utils/MultisigCallUtils.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

/**
 * Purpose: 
 *      * enqueue a multisig transaction which;
 *             - upgrades all the relevant contracts, and
 *             - unpauses the system.
 *  This should be run via the protocol council multisig.
 */
contract Queue is MultisigBuilder {
    using MultisigCallUtils for MultisigCall[];
    using EigenLabsUpgrade for *;
    using EncGnosisSafe for *;
    using MultisigCallUtils for *;

    MultisigCall[] private _executorCalls;
    MultisigCall[] private _opsCalls;

    function _getMultisigTransactionCalldata() internal view returns (bytes memory) {
        ProxyAdmin pa = ProxyAdmin(this._proxyAdmin());

        // TODO(alex): multisig transaction calldata for upgrading all contracts from phase 1.
        bytes memory executorMultisigCalldata = new bytes(0);
        // for syntax, see rewardsv2 upgrade script

        return (executorMultisigCalldata);
    }

    function _runAsMultisig() internal virtual override {
        (bytes memory call) = _getMultisigTransactionCalldata();

        TimelockController timelock = TimelockController(payable(this._timelock()));
        timelock.schedule(
            this._executorMultisig(), 
            0, 
            call,
            0,
            bytes32(0),  
            timelock.getMinDelay()
        );
    }

    function testDeploy() virtual public {
        zSetMultisigContext(this._protocolCouncilMultisig());
        execute();

        TimelockController timelock = TimelockController(payable(this._timelock()));
        bytes memory call = _getMultisigTransactionCalldata();
        bytes32 txHash = timelock.hashOperation(this._executorMultisig(), 0, call, 0, 0);
        assertEq(timelock.isOperationPending(txHash), true, "Transaction should be queued.");
    }
}
