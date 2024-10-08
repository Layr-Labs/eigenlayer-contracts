// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/Release_Template.s.sol";

contract QueuedTransactions {
    /**
     * Multisig Transactions to queue.
     */
    function get(Addresses memory addrs, Environment memory env, Params memory params) public returns (MultisigCall[] memory) {
        MultisigCall[] memory _executorCalls = new MultisigCall[](2);
        _executorCalls[0] = MultisigCall({
            to: addrs.eigenPod.beacon,
            data: abi.encodeWithSelector(
                IUpgradeableBeacon.upgradeTo.selector,
                addrs.eigenPod.pendingImpl
            ),
            value: 0
        });

        _executorCalls[1] = MultisigCall({
            to: addrs.proxyAdmin,
            data: abi.encodeWithSelector(
                ProxyAdmin.upgrade.selector,
                addrs.eigenPodManager,
                addrs.eigenPod.pendingImpl
            ),
            value: 0
        });

        return _executorCalls;
    }
}