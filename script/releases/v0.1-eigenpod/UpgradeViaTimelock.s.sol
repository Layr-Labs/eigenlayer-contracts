// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/Release_Template.s.sol";
import {IUpgradeableBeacon} from "script/utils/Interfaces.sol";

contract UpgradeCounter is OpsTimelockBuilder {

    using MultisigCallHelper for MultisigCall;

    FinalExecutorCall[] finalCalls;

    MultisigCall[] opsCalls;

    function _makeTimelockTxns(Addresses memory addrs, Environment memory env, Params memory params) internal virtual returns (FinalExecutorCall[] memory) {
        finalCalls.append({
            to: addrs.eigenPod.beacon,
            data: abi.encodeWithSelector(
                IUpgradeableBeacon.upgradeTo.selector,
                addrs.eigenPod.pendingImpl
            )
        });

        finalCalls.append({
            to: addrs.proxyAdmin,
            data: abi.encodeWithSelector(
                ProxyAdmin.upgrade.selector,
                addrs.eigenPodManager,
                addrs.eigenPod.pendingImpl
            )
        });

        return finalCalls;
    }

    function _queue(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (MultisigCall[] memory) {
        _makeTimelockTxns(addrs, env, params);

        opsCalls.append({
            to: addrs.admin.timelock,
            data: EncTimelock.queueTransaction(finalCalls)
        });

        return opsCalls;
    }

    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (MultisigCall[] memory) {
        opsCalls.append({
            to: addrs.admin.timelock,
            data: EncTimelock.executeTransaction(),
            data: _makeTimelockTxns(addrs, env, params)
        });

        opsCalls.append({
            to: addrs.strategyFactory.proxy,
            data: IStrategyFactory.whitelistThing(details)
        });

        return opsCalls;
    }

    function _test_Execute(
        Addresses memory addrs,
        Environment memory env,
        Params memory params
    ) internal override {
        bytes memory data = encodeMultisendTxs(arr);

        vm.startBroadcast(addrs.admin.opsMultisig);
        addrs.admin.multiSend.delegatecall(data);
        vm.stopBroadcast();
    }
}