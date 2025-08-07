// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
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
        address opsMultisig =
            deployMultisig({initialOwners: opsMultisigInitialOwners, initialThreshold: 3, salt: ++salt});

        // Deploy pauserMultisig
        address pauserMultisig =
            deployMultisig({initialOwners: pauserMultisigInitialOwners, initialThreshold: 1, salt: ++salt});

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
        assertEq(Multisig(address(Env.opsMultisig())).getThreshold(), 3);
        assertEq(Multisig(address(Env.pauserMultisig())).getThreshold(), 1);

        // Assert the owners of each multisig
        address[] memory opsMultisigOwners =
            _getMultisigOwner("script/releases/v1.6.0-multichain-genesis/opsOwners.toml");
        address[] memory pauserMultisigOwners =
            _getMultisigOwner("script/releases/v1.6.0-multichain-genesis/pauserOwners.toml");
        for (uint256 i = 0; i < opsMultisigOwners.length; i++) {
            assertTrue(Multisig(address(Env.opsMultisig())).isOwner(opsMultisigOwners[i]));
        }
        for (uint256 i = 0; i < pauserMultisigOwners.length; i++) {
            assertTrue(Multisig(address(Env.pauserMultisig())).isOwner(pauserMultisigOwners[i]));
        }

        // Assert owner counts are correct
        assertEq(
            opsMultisigOwners.length,
            Multisig(address(Env.opsMultisig())).getOwners().length,
            "opsMultisigOwners length mismatch"
        );
        assertEq(
            pauserMultisigOwners.length,
            Multisig(address(Env.pauserMultisig())).getOwners().length,
            "pauserMultisigOwners length mismatch"
        );

        // Assert that the pauserRegistry is non-zero, and that the pausers are set correctly
        assertTrue(Env.impl.pauserRegistry().isPauser(Env.opsMultisig()));
        assertTrue(Env.impl.pauserRegistry().isPauser(Env.pauserMultisig()));
        assertEq(Env.impl.pauserRegistry().unpauser(), Env.opsMultisig());
    }

    function deployMultisig(
        address[] memory initialOwners,
        uint256 initialThreshold,
        uint256 salt
    ) internal returns (address) {
        // addresses taken from https://github.com/safe-global/safe-smart-account/blob/main/CHANGELOG.md#expected-addresses-with-deterministic-deployment-proxy-default
        // NOTE: double check these addresses are correct on each chain
        address safeFactory = 0x4e1DCf7AD4e460CfD30791CCC4F9c8a4f820ec67;
        address safeSingleton = 0x29fcB43b46531BcA003ddC8FCB67FFE91900C762; // Gnosis safe L2 singleton
        address safeFallbackHandler = 0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99;

        bytes memory emptyData;

        bytes memory initializerData = abi.encodeWithSignature(
            "setup(address[],uint256,address,bytes,address,address,uint256,address)",
            initialOwners, /* signers */
            initialThreshold, /* threshold */
            address(0), /* to (used in setupModules) */
            emptyData, /* data (used in setupModules) */
            safeFallbackHandler,
            address(0), /* paymentToken */
            0, /* payment */
            payable(address(0)) /* paymentReceiver */
        );

        bytes memory calldataToFactory =
            abi.encodeWithSignature("createProxyWithNonce(address,bytes,uint256)", safeSingleton, initializerData, salt);

        (bool success, bytes memory returndata) = safeFactory.call(calldataToFactory);
        require(success, "multisig deployment failed");
        address deployedMultisig = abi.decode(returndata, (address));
        require(deployedMultisig != address(0), "something wrong in multisig deployment, zero address returned");
        return deployedMultisig;
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

interface Multisig {
    function getThreshold() external view returns (uint256);
    function getOwners() external view returns (address[] memory);
    function isOwner(
        address owner
    ) external view returns (bool);
}
