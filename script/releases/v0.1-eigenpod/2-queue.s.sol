// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "script/Release_Template.s.sol";
import {IUpgradeableBeacon} from "script/utils/Interfaces.sol";
import {QueuedTransactions} from "script/releases/v0.1-eigenpod/queued-transactions.sol";

import "src/contracts/interfaces/IStrategyFactory.sol";

contract Queue is OpsTimelockBuilder {

    using MultisigCallHelper for *;
    using TransactionHelper for *;

    function _queue(Addresses memory addrs, Environment memory env, Params memory params) internal override returns (MultisigCall[] memory) {
        return new QueuedTransactions().get(addrs, env, params);
    }

    function _testExecute(
        Addresses memory addrs,
        Environment memory env,
        Params memory params
    ) internal override {
       // TODO: fill in test (?)
       //      - assert some precondition about the state of the network prior to upgrade.
    }
}