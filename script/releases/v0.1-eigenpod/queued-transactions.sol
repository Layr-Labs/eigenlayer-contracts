// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/Release_Template.s.sol";

contract QueuedTransactions {

    MultisigCall[] _executorCalls;

    /**
     * Multisig Transactions to queue.
     */
    function get(Addresses memory addrs, Environment memory env, Params memory params) public returns (MultisigCall[] memory) {
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
}