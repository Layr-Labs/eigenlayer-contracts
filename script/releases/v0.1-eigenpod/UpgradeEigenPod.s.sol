// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/Release_Template.s.sol";

contract UpgradeCounter is MultisigBuilder {

    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (Tx[] memory) {
        Tx[] memory txs = new Tx[](2);

        // txs[0] = Tx({
        //     to: eigenPodBeacon,
        //     value: 0,
        //     data: abi.encodeWithSelector(
        //         IUpgradeableBeacon.upgradeTo.selector, newEigenPodImpl
        //     )
        // });

        // txs[1] = Tx({
        //     to: eigenLayerProxyAdmin,
        //     value: 0,
        //     data: abi.encodeWithSelector(ProxyAdmin.upgrade.selector, eigenPodManager, newEigenPodManagerImpl)
        // });

        bytes memory calldata_to_executor; // upgrade data

        Transaction({
            to: addrs.proxyAdmin,
            value: 0,
            data: calldata_to_executor,
            op: EncGnosisSafe.Operation.DelegateCall
        });

        return txs;
    }
}