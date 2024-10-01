// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./Release_Template.s.sol";

// zeus new "pepe" opsmultisig

contract PEPE_Upgrade is OpsTimelockBuilder {

    struct TimelockData {
        address target;
        uint value;
        string signature;
        bytes data;
        uint eta;
    }

    function _makeTimelockData(
        Addresses memory addrs,
        Environment memory env,
        Params memory params
    ) internal override returns (TimelockData memory) {
        TimelockData data = _newTimelockData();
        data.eta = env.isMainnet() ? 12351235 : 0;

        data.appendCall({
            to: addrs.eigenPod.beacon,
            data: EncBeacon.upgradeTo(addrs.eigenPod.getPending())
        });

        data.appendCall({
            to: addrs.proxyAdmin,
            data: EncProxyAdmin.upgrade(addrs.eigenPodManager.proxy, addrs.eigenPodManager.getPending())
        });

        return data.encode();
    }
    
    function _queue(
        Addresses memory addrs,
        Environment memory env,
        Params memory params
    ) internal override returns (Calls[] memory) {
        TimelockData memory data = _makeTimelockData(addrs, env, params);

        return opsMultisig.queue(txns);
    }
    
    function _execute(
        Addresses memory addrs,
        Environment memory env,
        Params memory params
    ) internal override returns (Calls[]) {
        TimelockTxn[] memory txns = _makeTimelockTxns(addrs, env, params);

        Calls[] memory final = opsMultsig.execute(txns)
            .append({
                to: awfe
                data: awef
            });

        return final;
    }
}

contract PEPE_Deploy is EOABuilder {

    function _deploy(
        Addresses memory addrs,
        Environment memory env,
        Params memory params
    ) internal override returns (Deployment[]) {
        
    }
}