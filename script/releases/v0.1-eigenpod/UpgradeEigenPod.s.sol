// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/templates/MultisigBuilder.sol";

import {IUpgradeableBeacon} from "script/utils/Interfaces.sol";

contract UpgradeEigenPod is MultisigBuilder {

    using MultisigCallUtils for MultisigCall[];

    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (MultisigCall[] memory) {
        _multisigCalls.append({
            to: addrs.timelock,
            value: 0,
            data: abi.encodeWithSelector(
                    ITimelock.executeTransaction.selector,
                    addrs.proxyAdmin,
                    0,
                    "",
                    bytes(""),
                    uint(0)
                )
            }
        );

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

    function _testExecute(
        Addresses memory addrs,
        Environment memory env,
        Params memory params
    ) internal override {
        bytes memory data = _multisigCalls.encodeMultisendTxs();

        vm.startBroadcast(addrs.operationsMultisig);
        params.multiSendCallOnly.delegatecall(data);
        vm.stopBroadcast();

        emit log_bytes(data);
    }
}
