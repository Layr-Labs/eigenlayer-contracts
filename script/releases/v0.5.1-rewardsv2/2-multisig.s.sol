// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {RewardsCoordinator} from "src/contracts/core/RewardsCoordinator.sol";
import {Deploy} from "./1-eoa.s.sol";
import "zeus-templates/utils/Encode.sol";
import "../Env.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/**
 * Purpose: enqueue a multisig transaction which tells the ProxyAdmin to upgrade RewardsCoordinator.
 */
contract Queue is MultisigBuilder, Deploy {
    using Env for *;
    using Encode for *;

    function _getCalldataToExecutor() internal returns (bytes memory) {
        ProxyAdmin pa = ProxyAdmin(Env.proxyAdmin());
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls()
            .append({
                to: Env.proxyAdmin(),
                data: abi.encodeCall(
                    pa.upgrade,
                    (
                        ITransparentUpgradeableProxy(payable(address(Env.proxy.rewardsCoordinator()))),
                        address(Env.impl.rewardsCoordinator())
                    )
                )
            });

        return Encode.gnosisSafe.execTransaction({
            from: address(Env.timelockController()),
            to: Env.multiSendCallOnly(),
            op: Encode.Operation.DelegateCall,
            data: Encode.multiSend(executorCalls)
        });
    }

    function _runAsMultisig() prank(Env.opsMultisig()) internal virtual override {
        bytes memory executorMultisigCalldata = _getCalldataToExecutor();

        TimelockController timelock = Env.timelockController();
        timelock.schedule(
            Env.executorMultisig(),
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
        TimelockController timelock = Env.timelockController();

        bytes memory multisigTxnData = _getCalldataToExecutor();
        bytes32 txHash = timelock.hashOperation(Env.executorMultisig(), 0, multisigTxnData, 0, 0);

        assertEq(timelock.isOperationPending(txHash), true, "Transaction should be queued.");
    }
}
