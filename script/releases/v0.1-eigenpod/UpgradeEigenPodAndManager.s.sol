// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/templates/OpsTimelockBuilder.sol";

import {IUpgradeableBeacon} from "script/utils/Interfaces.sol";

import "src/contracts/interfaces/IStrategyFactory.sol";
import "src/contracts/pods/EigenPodManager.sol";

contract UpgradeEigenPodAndManager is OpsTimelockBuilder {

    using MultisigCallUtils for MultisigCall[];
    using SafeTxUtils for *;

    function _queue(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (MultisigCall[] memory) {

        // construct initialization data for eigenPodManager
        bytes memory eigenPodManagerData = abi.encodeWithSelector(
            EigenPodManager(addrs.eigenPodManager.pendingImpl).initialize.selector,
            addrs.operationsMultisig, // set opsMultisig as new direct owner
            addrs.pauserRegistry, // set pauser registry
            uint256(0) // set all 0 bits, nothing paused
        );

        // upgrade eigenPodManager
        _executorCalls.append({
            to: addrs.proxyAdmin,
            data: abi.encodeWithSelector(
                ProxyAdmin.upgradeAndCall.selector,
                addrs.eigenPodManager.proxy,
                addrs.eigenPodManager.pendingImpl,
                eigenPodManagerData // initialize impl here
            )
        });

        // upgrade eigenPod beacon implementation
        _executorCalls.append({
            to: addrs.eigenPod.beacon,
            data: abi.encodeWithSelector(
                IUpgradeableBeacon.upgradeTo.selector,
                addrs.eigenPod.pendingImpl
            )
        });

        return _executorCalls;
    }

    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (MultisigCall[] memory) {

        // steals logic from queue() to perform execute()
        // likely the first step of any _execute() after a _queue()
        bytes memory executorCalldata = makeExecutorCalldata(
            _queue(addrs, env, params),
            params.multiSendCallOnly,
            addrs.timelock
        );

        // execute queued transaction upgrading eigenPodManager and eigenPod
        _opsCalls.append({
            to: addrs.timelock,
            value: 0,
            data: abi.encodeWithSelector(
                ITimelock.executeTransaction.selector,
                executorCalldata
            )
        });

        // after queued transaction, renounce ownership from eigenPodManager
        _opsCalls.append({
            to: addrs.eigenPodManager.proxy,
            value: 0,
            data: abi.encodeWithSelector(
                EigenPodManager(addrs.eigenPodManager.proxy).renounceOwnership.selector
            )
        });

        return _opsCalls;
    }
}
