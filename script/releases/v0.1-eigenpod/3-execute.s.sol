// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/Release_Template.s.sol";
import {IUpgradeableBeacon} from "script/utils/Interfaces.sol";
import {QueuedTransactions} from "script/releases/v0.1-eigenpod/queued-transactions.sol";


import "src/contracts/interfaces/IStrategyFactory.sol";

contract Execute is MultisigBuilder {

    using MultisigCallHelper for *;
    using TransactionHelper for *;

    MultisigCall[] _opsCalls;

    function _execute(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (MultisigCall[] memory) {
        //////////////////////////////////
        ////////////////////////////////
        MultisigCall[] memory txns;

        txns = new QueuedTransactions().getTransactions(addrs, env, params);
        bytes memory executorCalldata = new OpsTimelockBuilder().makeExecutorCalldata(
            txns,
            params.multiSendCallOnly
        );
         _opsCalls.append({
            to: addrs.timelock,
            value: 0,
            data: abi.encodeWithSelector(
                ITimelock.executeTransaction.selector,
                executorCalldata
            )
        });
        ////////////////////////////////////
        ////////////////////////////////////

        // TODO: add any additional calls to be performed...

        // steals logic from queue() to perform execute()
        

        // _executorCalls.append(multisigCalls, to, value, data);


        // e.g
        // _opsCalls.append({
        //     to: addrs.strategyFactory.proxy,
        //     value: 0,
        //     data: abi.encodeWithSelector(
        //         IStrategyFactory.deployNewStrategy.selector,
        //         address(0xE)
        //     )
        // });

        return _opsCalls;
    }

    function _testExecute(
        Addresses memory addrs,
        Environment memory env,
        Params memory params
    ) internal override {
        bytes memory data = _opsCalls.encodeMultisendTxs();

        vm.startBroadcast(addrs.operationsMultisig);
        params.multiSendCallOnly.delegatecall(data);
        vm.stopBroadcast();
    }
}