// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {DeployGovernance} from "script/releases/v0.0.0-source-genesis/1-deployGovernance.s.sol";
import {DeployOperations} from "script/releases/v0.0.0-source-genesis/2-deployOperations.s.sol";
import "../Env.sol";

contract DeployCore is DeployOperations {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        

        vm.stopBroadcast();
    }

    function testScript() public virtual override {
        // Run the governance and operations scripts
        DeployGovernance._runAsEOA();
        DeployOperations._runAsEOA();

        // Run the core script
        runAsEOA();
    }
}