// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/Release_Template.s.sol";
import {IUpgradeableBeacon} from "script/utils/Interfaces.sol";

contract UpgradeCounter is MultisigBuilder {

    using MultisigCallHelper for *;

    MultisigCall[] _multisigCalls;

    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (MultisigCall[] memory) {
        _multisigCalls.append({
            to: addrs.admin.timelock,
            data: abi.encodeWithSelector(
                ITimelock.executeTransaction.selector({
                    to: addrs.proxyAdmin,
                    value: 0,
                    signature: bytes(0),
                    data: bytes(0),
                    eta: uint(0)
                })
            )
        });

        _multisigCalls.append({
            to: addrs.proxyAdmin,
            data: abi.encodeWithSelector(
                ProxyAdmin.upgrade.selector,
                addrs.eigenPodManager,
                addrs.eigenPod.pendingImpl
            )
        });

        return _multisigCalls;
    }

    function _test_Execute(
        Addresses memory addrs,
        Environment memory env,
        Params memory params
    ) internal override {
        bytes memory data = encodeMultisendTxs(_multisigCalls);

        vm.startBroadcast(addrs.admin.opsMultisig);
        addrs.admin.multiSend.delegatecall(data);
        vm.stopBroadcast();


    }
}