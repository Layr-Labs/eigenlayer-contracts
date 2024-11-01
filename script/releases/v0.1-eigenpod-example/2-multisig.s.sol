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
        // construct initialization data for eigenPodManager
        bytes memory eigenPodManagerData = abi.encodeWithSelector(
            EigenPodManager.initialize.selector,
            _operationsMultisig(), // set opsMultisig as new direct owner
            _pauserRegistry(), // set pauser registry
            uint256(0) // set all 0 bits, nothing paused
        );

        // upgrade eigenPodManager
        _executorCalls.append({
            to: _proxyAdmin(),
            data: abi.encodeWithSelector(
                ProxyAdmin.upgradeAndCall.selector,
                _eigenPodManagerProxy(),
                _eigenPodManagerPendingImpl(),
                eigenPodManagerData // initialize impl here
            )
        });

        // upgrade eigenPod beacon implementation
        _executorCalls.append({
            to: _eigenPodBeacon(),
            data: abi.encodeWithSelector(IUpgradeableBeacon.upgradeTo.selector, _eigenPodPendingImpl())
        });

        return _executorCalls;
    }

    function _execute() internal virtual override returns (MultisigCall[] memory) {
        // get the queue data
        MultisigCall[] memory calls = _queue();

        // encode calls for executor
        bytes memory executorCalldata = calls.makeExecutorCalldata(_multiSendCallOnly(), _timelock());

        // encode executor data for timelock
        bytes memory timelockCalldata = abi.encodeWithSelector(
            ITimelock.queueTransaction.selector, _executorMultisig(), 0, "", executorCalldata, type(uint256).max
        );

        _opsCalls.append(_timelock(), timelockCalldata);

        // encode timelock data for ops multisig
        return _opsCalls;
    }

    function zeusTest() public virtual override {
        // Test function implementation
    }
}
