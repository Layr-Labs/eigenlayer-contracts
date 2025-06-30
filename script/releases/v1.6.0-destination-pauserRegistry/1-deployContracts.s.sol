// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

contract DeployDestinationGenesis is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        if (!Env.isDestinationChain()) {
            return;
        }

        vm.startBroadcast();

        // Deploy pauserRegistry
        address[] memory pausers = new address[](2);
        pausers[0] = Env.opsMultisig();
        pausers[1] = Env.pauserMultisig();

        deployImpl({
            name: type(PauserRegistry).name,
            deployedTo: address(new PauserRegistry({_pausers: pausers, _unpauser: Env.opsMultisig()}))
        });

        vm.stopBroadcast(); 
    }

    function testScript() public virtual {
        if (!Env.isDestinationChain()) {
            return;
        }

        // Assert that the pauserRegistry is non-zero, and that the pausers are set correctly
        assertTrue(Env.impl.pauserRegistry().isPauser(Env.opsMultisig()));
        assertTrue(Env.impl.pauserRegistry().isPauser(Env.pauserMultisig()));
        assertEq(Env.impl.pauserRegistry().unpauser(), Env.opsMultisig());
    }
}