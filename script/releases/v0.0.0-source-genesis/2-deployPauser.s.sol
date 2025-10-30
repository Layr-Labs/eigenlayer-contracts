// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {DeployGovernance} from "script/releases/v0.0.0-source-genesis/1-deployGovernance.s.sol";
import {MultisigDeployLib} from "../MultisigDeployLib.sol";
import "../Env.sol";


/// @dev This script is used to deploy the operations contracts on a testnet environment.
/// This script deploys the following contracts/msigs:
/// - pauserMultisig
/// - pauserRegistry
/// - proxyAdmin
contract DeployPauser is DeployGovernance {
    using Env for *;

    function _runAsEOA() internal virtual override {
        // Setup safe deploy parameters
        uint256 salt = uint256(keccak256(abi.encode(block.chainid, block.timestamp + 1))); // Pseudo-random salt; We add 1 to the timestamp to ensure the salt is different from the governance script
        address[] memory initialOwners = new address[](1);
        initialOwners[0] = TESTNET_OWNER;

        vm.startBroadcast();

        // Deploy pauserMultisig
        address pauserMultisig =
            MultisigDeployLib.deployMultisig({initialOwners: initialOwners, initialThreshold: TESTNET_THRESHOLD, salt: ++salt});

        // Deploy pauserRegistry
        address[] memory pausers = new address[](2);
        pausers[0] = Env.opsMultisig();
        pausers[0] = Env.executorMultisig();
        pausers[1] = pauserMultisig;

        deployImpl({
            name: type(PauserRegistry).name,
            deployedTo: address(new PauserRegistry({_pausers: pausers, _unpauser: Env.executorMultisig()}))
        });

        vm.stopBroadcast();

        // Update config
        zUpdate("pauserMultisig", pauserMultisig);
    }

    function testScript() public virtual override {
        // Run the deploy governance script, since we need the executor multisig to be deployed
        DeployGovernance._runAsEOA();
        
        // Run the deploy operations script
        runAsEOA();

        // Check the pauser multisig
        assertNotEq(Env.pauserMultisig(), address(0));
        assertEq(MultisigDeployLib.getThreshold(Env.pauserMultisig()), TESTNET_THRESHOLD);

        // Assert the owners of each multisig
        address[] memory expectedOwners = new address[](1);
        expectedOwners[0] = TESTNET_OWNER;
        for (uint256 i = 0; i < expectedOwners.length; i++) {
            assertTrue(MultisigDeployLib.isOwner(Env.pauserMultisig(), expectedOwners[i]));
        }
        assertEq(
            expectedOwners.length,
            MultisigDeployLib.getOwners(Env.pauserMultisig()).length,
            "pauserMultisigOwners length mismatch"
        );

        // Assert that the pauserRegistry is non-zero, and that the pausers are set correctly
        assertTrue(Env.impl.pauserRegistry().isPauser(Env.opsMultisig()));
        assertTrue(Env.impl.pauserRegistry().isPauser(Env.pauserMultisig()));
        assertTrue(Env.impl.pauserRegistry().isPauser(Env.executorMultisig()));
        assertEq(Env.impl.pauserRegistry().unpauser(), Env.executorMultisig());
    }
}