// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {EigenLabsUpgrade} from "../EigenLabsUpgrade.s.sol";

/**
 * Purpose: use an EOA to deploy all of the new contracts for this upgrade. 
 */
contract Deploy is EOADeployer {
    using EigenLabsUpgrade for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();
        
        // TODO(alex): deploy all of the slashing contracts, recording the address with `deploySingleton(addr, name)`.
        // for `name`, likely use `this.impl(type(Contract).name)`.

        vm.stopBroadcast();
    }

    function testDeploy() public {
        _runAsEOA();
        Deployment[] memory deploys = deploys();

        // TODO(alex): assert on all of our deployments.
    }
}
