// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";

contract Deploy is EOADeployer {
    Deployment[] private _deployments;

    function _deploy() internal override returns (Deployment[] memory) {
        vm.startBroadcast();

        //////////////////////////
        // deploy your contracts here
        //////////////////////////

        vm.stopBroadcast();

        return _deployments;
    }

    function zeusTest() public override {
        // Test function implementation
    }
}
