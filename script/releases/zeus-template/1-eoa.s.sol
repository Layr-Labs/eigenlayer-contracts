// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Addresses, Environment, Params, Deployment, EOADeployer} from "zeus-templates/templates/EOADeployer.sol";

contract Deploy is EOADeployer {

    function _deploy(Addresses memory, Environment memory, Params memory) internal override returns (Deployment[] memory) {

        vm.startBroadcast();

        //////////////////////////
        // deploy your contracts here
        //////////////////////////

        vm.stopBroadcast();

        return _deployments;
    }
}