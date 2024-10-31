// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigCall, MultisigCallUtils, MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {SafeTx, SafeTxUtils} from "zeus-templates/utils/SafeTxUtils.sol";

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {IUpgradeableBeacon} from "script/interfaces/IUpgradeableBeacon.sol";
import {ITimelock} from "zeus-templates/interfaces/ITimelock.sol";
import "src/contracts/pods/EigenPodManager.sol";

contract QueueEigenPodAndManager is MultisigBuilder {
    using MultisigCallUtils for MultisigCall[];
    using SafeTxUtils for SafeTx;

    MultisigCall[] private _executorCalls;
    MultisigCall[] private _opsCalls;

    function _queue() internal returns (MultisigCall[] memory) {
        address eigenPodManagerPendingImpl = zeusAddress("EigenPodManager_pendingImpl");
        address operationsMultisig = zeusAddress("OperationsMultisig");
        address pauserRegistry = zeusAddress("PauserRegistry");

        // construct initialization data for eigenPodManager
        bytes memory eigenPodManagerData = abi.encodeWithSelector(
            EigenPodManager(eigenPodManagerPendingImpl).initialize.selector,
            operationsMultisig, // set opsMultisig as new direct owner
            pauserRegistry, // set pauser registry
            uint256(0) // set all 0 bits, nothing paused
        );

        address proxyAdmin = zeusAddress("ProxyAdmin");
        address eigenPodManagerProxy = zeusAddress("EigenPodManager_proxy");

        // upgrade eigenPodManager
        _executorCalls.append({
            to: proxyAdmin,
            data: abi.encodeWithSelector(
                ProxyAdmin.upgradeAndCall.selector,
                eigenPodManagerProxy,
                eigenPodManagerPendingImpl,
                eigenPodManagerData // initialize impl here
            )
        });

        address eigenPodBeacon = zeusAddress("EigenPod_beacon");
        address eigenPodPendingImpl = zeusAddress("EigenPod_pendingImpl");

        // upgrade eigenPod beacon implementation
        _executorCalls.append({
            to: eigenPodBeacon,
            data: abi.encodeWithSelector(IUpgradeableBeacon.upgradeTo.selector, eigenPodPendingImpl)
        });

        return _executorCalls;
    }

    function _execute() internal virtual override returns (MultisigCall[] memory) {
        // get the queue data
        MultisigCall[] memory calls = _queue();

        address multiSendCallOnly = zeusAddress("MultiSendCallOnly");
        address timelock = zeusAddress("Timelock");

        // encode calls for executor
        bytes memory executorCalldata = calls.makeExecutorCalldata(multiSendCallOnly, timelock);

        address executorMultisig = zeusAddress("ExecutorMultisig");

        // encode executor data for timelock
        bytes memory timelockCalldata = abi.encodeWithSelector(
            ITimelock.queueTransaction.selector, executorMultisig, 0, "", executorCalldata, type(uint256).max
        );

        _opsCalls.append(timelock, timelockCalldata);

        // encode timelock data for ops multisig
        return _opsCalls;
    }

    function zeusTest() public override virtual {
        // Test function implementation
    }
}
