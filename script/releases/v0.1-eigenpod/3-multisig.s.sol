// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "zeus-templates/templates/OpsTimelockBuilder.sol";

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {IUpgradeableBeacon} from "script/interfaces/IUpgradeableBeacon.sol";
import "src/contracts/interfaces/IStrategyFactory.sol";
import "src/contracts/pods/EigenPodManager.sol";

import "./2-multisig.s.sol";

contract ExecuteEigenPodAndManager is MultisigBuilder {

    using MultisigCallUtils for MultisigCall[];
    using SafeTxUtils for *;

    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (MultisigCall[] memory) {

        QueueEigenPodAndManager queue = new QueueEigenPodAndManager();

        MultisigCall[] memory _executorCalls = queue.queue(addrs, env, params);

        // steals logic from queue() to perform execute()
        // likely the first step of any _execute() after a _queue()
        bytes memory executorCalldata = queue.makeExecutorCalldata(
            _executorCalls,
            params.multiSendCallOnly,
            addrs.timelock
        );

        // execute queued transaction upgrading eigenPodManager and eigenPod
        _multisigCalls.append({
            to: addrs.timelock,
            value: 0,
            data: abi.encodeWithSelector(
                ITimelock.executeTransaction.selector,
                executorCalldata
            )
        });

        // after queued transaction, renounce ownership from eigenPodManager
        _multisigCalls.append({
            to: addrs.eigenPodManager.proxy,
            value: 0,
            data: abi.encodeWithSelector(
                EigenPodManager(addrs.eigenPodManager.proxy).renounceOwnership.selector
            )
        });

        return _multisigCalls;
    }
}
