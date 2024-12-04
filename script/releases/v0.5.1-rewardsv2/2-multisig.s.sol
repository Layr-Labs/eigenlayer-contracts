// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigCall, MultisigCallUtils, MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {Deploy} from "./1-eoa.s.sol";
import {RewardsCoordinator} from "src/contracts/core/RewardsCoordinator.sol";
import {EigenLabsUpgrade} from "../EigenLabsUpgrade.s.sol";
import {IPauserRegistry} from "src/contracts/interfaces/IPauserRegistry.sol";
import {ITimelock} from "zeus-templates/interfaces/ITimelock.sol";
import {console} from "forge-std/console.sol";
import {EncGnosisSafe} from "zeus-templates/utils/EncGnosisSafe.sol";
import {MultisigCallUtils, MultisigCall} from "zeus-templates/utils/MultisigCallUtils.sol";
import {IMultiSend} from "zeus-templates/interfaces/IMultiSend.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

/**
 * Purpose: enqueue a multisig transaction which tells the ProxyAdmin to upgrade RewardsCoordinator.
 */
contract Queue is MultisigBuilder, Deploy {
    using MultisigCallUtils for MultisigCall[];
    using EigenLabsUpgrade for *;
    using EncGnosisSafe for *;
    using MultisigCallUtils for *;

    MultisigCall[] private _executorCalls;
    MultisigCall[] private _opsCalls;

    function options() internal virtual override view returns (MultisigOptions memory) {
        return MultisigOptions(
            this._operationsMultisig(),
            Operation.Call
        );
    }

    function _getMultisigTransactionCalldata() internal view returns (bytes memory) {
        ProxyAdmin pa = ProxyAdmin(this._proxyAdmin());

        bytes memory proxyAdminCalldata = abi.encodeCall(
            pa.upgrade,
            (
                TransparentUpgradeableProxy(payable(zDeployedProxy(type(RewardsCoordinator).name))),
                zDeployedImpl(type(RewardsCoordinator).name)
            )
        );

        bytes memory executorMultisigCalldata = address(this._timelock()).calldataToExecTransaction(
            this._proxyAdmin(),
            proxyAdminCalldata,
            EncGnosisSafe.Operation.Call
        );

        return (executorMultisigCalldata);
    }

    function runAsMultisig() internal virtual override {
        bytes memory executorMultisigCalldata = _getMultisigTransactionCalldata();

        TimelockController timelock = TimelockController(payable(this._timelock()));
        timelock.schedule(
            this._executorMultisig(),
            0 /* value */,
            executorMultisigCalldata,
            0 /* predecessor */,
            bytes32(0) /* salt */,
            timelock.getMinDelay()
        );
    }

    function testDeploy() public virtual override {
        runAsEOA();

        execute();
        TimelockController timelock = TimelockController(payable(this._timelock()));

        bytes memory multisigTxnData = _getMultisigTransactionCalldata();
        bytes32 txHash = timelock.hashOperation(this._executorMultisig(), 0, multisigTxnData, 0, 0);

        assertEq(timelock.isOperationPending(txHash), true, "Transaction should be queued.");
    }
}
