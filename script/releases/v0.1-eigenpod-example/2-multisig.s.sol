// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "zeus-templates/templates/OpsTimelockBuilder.sol";

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {IUpgradeableBeacon} from "script/interfaces/IUpgradeableBeacon.sol";
import "src/contracts/pods/EigenPodManager.sol";

contract QueueEigenPodAndManager is OpsTimelockBuilder {

    using MultisigCallUtils for MultisigCall[];
    using SafeTxUtils for SafeTx;

    MultisigCall[] internal _executorCalls;

    function queue(Addresses memory addrs, Environment memory env, Params memory params) public override returns (MultisigCall[] memory) {

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
}
