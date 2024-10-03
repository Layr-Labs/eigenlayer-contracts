// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/Release_Template.s.sol";
import {IUpgradeableBeacon} from "script/utils/Interfaces.sol";

contract UpgradeCounter is MultisigBuilder {

    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (Tx[] memory) {
        Tx[] memory txs = new Tx[](2);

        txs[0] = Tx({
            to: addrs.eigenPod.beacon,
            value: 0,
            data: abi.encodeWithSelector(
                IUpgradeableBeacon.upgradeTo.selector,
                addrs.eigenPod.pendingImpl
            )
        });

        txs[1] = Tx({
            to: addrs.proxyAdmin,
            value: 0,
            data: abi.encodeWithSelector(
                ProxyAdmin.upgrade.selector,
                addrs.eigenPodManager,
                addrs.eigenPod.pendingImpl
            )
        });

        return txs;
    }
}