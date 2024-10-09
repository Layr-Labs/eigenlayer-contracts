// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/Release_Template.s.sol";
import "src/contracts/pods/EigenPod.sol";
import "src/contracts/pods/EigenPodManager.sol";
import "src/contracts/interfaces/IEigenPodManager.sol";

import "script/utils/StringUtils.sol";

contract QueueEigenPodAndManager is QueueBuilder {
    using MultisigCallUtils for MultisigCall[];


    function _queue(Addresses memory addrs, Environment memory env, Params memory params) public override returns (MultisigCall[] memory) {

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