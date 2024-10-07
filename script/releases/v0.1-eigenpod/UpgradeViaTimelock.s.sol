// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/Release_Template.s.sol";
import {IUpgradeableBeacon} from "script/utils/Interfaces.sol";

import "src/contracts/interfaces/IStrategyFactory.sol";

contract UpgradeViaTimelock is OpsTimelockBuilder {

    using MultisigCallHelper for *;
    using TransactionHelper for *;

    MultisigCall[] _executorCalls;

    MultisigCall[] _opsCalls;

    function _queue(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (MultisigCall[] memory) {

        _executorCalls.append({
            to: addrs.eigenPod.beacon,
            data: abi.encodeWithSelector(
                IUpgradeableBeacon.upgradeTo.selector,
                addrs.eigenPod.pendingImpl
            )
        });

        _executorCalls.append({
            to: addrs.proxyAdmin,
            data: abi.encodeWithSelector(
                ProxyAdmin.upgrade.selector,
                addrs.eigenPodManager,
                addrs.eigenPod.pendingImpl
            )
        });

        return _executorCalls;
    }

    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (MultisigCall[] memory) {

        // steals logic from queue() to perform execute()
        bytes memory executorCalldata = makeExecutorCalldata(
            _queue(addrs, env, params),
            params.multiSendCallOnly,
            addrs.timelock
        );

        // _executorCalls.append(multisigCalls, to, value, data);

        _opsCalls.append({
            to: addrs.timelock,
            value: 0,
            data: abi.encodeWithSelector(
                ITimelock.executeTransaction.selector,
                executorCalldata
            )
        });

        _opsCalls.append({
            to: addrs.strategyFactory.proxy,
            value: 0,
            data: abi.encodeWithSelector(
                IStrategyFactory.deployNewStrategy.selector,
                address(0xE)
            )
        });

        return _opsCalls;
    }

    function _testExecute(
        Addresses memory addrs,
        Environment memory env,
        Params memory params
    ) internal override {
        bytes memory data = _opsCalls.encodeMultisendTxs();

        vm.startBroadcast(addrs.operationsMultisig);
        params.multiSendCallOnly.delegatecall(data);
        vm.stopBroadcast();
    }
}