// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {MultisigDeployLib, IMultisig} from "script/releases/MultisigDeployLib.sol";
import "../Env.sol";

// For TOML parsing
import {stdToml} from "forge-std/StdToml.sol";

contract DeployDestinationGenesis is EOADeployer {
    using Env for *;
    using stdToml for string;

    function _runAsEOA() internal override {
        if (!Env.isDestinationChain()) {
            return;
        }

        // Setup safe seploy parameters
        uint256 salt = uint256(keccak256(abi.encode(block.chainid, block.timestamp))); // Pseudo-random salt

        // Setup ops multisig initial owners
        address[] memory opsMultisigInitialOwners;
        opsMultisigInitialOwners = _getMultisigOwner("script/releases/v1.6.0-destination-genesis/opsOwners.toml");

        // Setup pauser multisig initial owners
        address[] memory pauserMultisigInitialOwners;
        pauserMultisigInitialOwners = _getMultisigOwner("script/releases/v1.6.0-destination-genesis/pauserOwners.toml");

        vm.startBroadcast();

        // Deploy opsMultisig
        address opsMultisig = MultisigDeployLib.deployMultisig({
            initialOwners: opsMultisigInitialOwners,
            initialThreshold: 3,
            salt: ++salt
        });

        // Deploy pauserMultisig
        address pauserMultisig = MultisigDeployLib.deployMultisig({
            initialOwners: pauserMultisigInitialOwners,
            initialThreshold: 1,
            salt: ++salt
        });

        // Deploy pauserRegistry
        address[] memory pausers = new address[](2);
        pausers[0] = opsMultisig;
        pausers[1] = pauserMultisig;

        deployImpl({
            name: type(PauserRegistry).name,
            deployedTo: address(new PauserRegistry({_pausers: pausers, _unpauser: opsMultisig}))
        });

        // Deploy proxyAdmin
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        proxyAdmin.transferOwnership(opsMultisig);

        vm.stopBroadcast();

        // Update config
        zUpdate("proxyAdmin", address(proxyAdmin));
        zUpdate("operationsMultisig", opsMultisig);
        zUpdate("pauserMultisig", pauserMultisig);
    }

    function testScript() public virtual {
        if (!Env.isDestinationChain()) {
            return;
        }

        super.runAsEOA();

        // Check proxyAdmin owner
        assertEq(ProxyAdmin(Env.proxyAdmin()).owner(), Env.opsMultisig());

        // Check that the multisigs are non-zero
        assertNotEq(Env.opsMultisig(), address(0));
        assertNotEq(Env.pauserMultisig(), address(0));

        // Check the threshold of each multisig
        assertEq(IMultisig(address(Env.opsMultisig())).getThreshold(), 3);
        assertEq(IMultisig(address(Env.pauserMultisig())).getThreshold(), 1);

        // Assert the owners of each multisig
        address[] memory opsMultisigOwners =
            _getMultisigOwner("script/releases/v1.6.0-destination-genesis/opsOwners.toml");
        address[] memory pauserMultisigOwners =
            _getMultisigOwner("script/releases/v1.6.0-destination-genesis/pauserOwners.toml");
        for (uint256 i = 0; i < opsMultisigOwners.length; i++) {
            assertTrue(IMultisig(address(Env.opsMultisig())).isOwner(opsMultisigOwners[i]));
        }
        for (uint256 i = 0; i < pauserMultisigOwners.length; i++) {
            assertTrue(IMultisig(address(Env.pauserMultisig())).isOwner(pauserMultisigOwners[i]));
        }

        // Assert owner counts are correct
        assertEq(
            opsMultisigOwners.length,
            IMultisig(address(Env.opsMultisig())).getOwners().length,
            "opsMultisigOwners length mismatch"
        );
        assertEq(
            pauserMultisigOwners.length,
            IMultisig(address(Env.pauserMultisig())).getOwners().length,
            "pauserMultisigOwners length mismatch"
        );

        // Assert that the pauserRegistry is non-zero, and that the pausers are set correctly
        assertTrue(Env.impl.pauserRegistry().isPauser(Env.opsMultisig()));
        assertTrue(Env.impl.pauserRegistry().isPauser(Env.pauserMultisig()));
        assertEq(Env.impl.pauserRegistry().unpauser(), Env.opsMultisig());
    }

    function _getMultisigOwner(
        string memory path
    ) internal view returns (address[] memory) {
        // Read the TOML file
        string memory root = vm.projectRoot();
        string memory fullPath = string.concat(root, "/", path);
        string memory toml = vm.readFile(fullPath);

        // Parse the owners array from the TOML
        address[] memory owners = toml.readAddressArray(".owners");

        return owners;
    }
}
