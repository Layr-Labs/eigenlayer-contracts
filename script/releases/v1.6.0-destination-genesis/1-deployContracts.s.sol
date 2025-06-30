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

        // Setup safe seploy parameters
        uint256 salt = uint256(keccak256(abi.encode(block.timestamp, block.chainid))); // Pseudo-random salt
        address[] memory initialOwners = new address[](0);
        initialOwners[0] = msg.sender;

        vm.startBroadcast();

        // NOTE: `multichainDeployerMultisig` is already deployed

        // Deploy opsMultisig
        address opsMultisig = deployMultisig({
            initialOwners: initialOwners,
            initialThreshold: 1,
            salt: ++salt
        });

        // Deploy pauserMultisig
        address pauserMultisig = deployMultisig({
            initialOwners: initialOwners, 
            initialThreshold: 1,
            salt: ++salt
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

    function deployMultisig(
        address[] memory initialOwners,
        uint256 initialThreshold,
        uint256 salt
    ) internal returns (address) {
        // addresses taken from https://github.com/safe-global/safe-smart-account/blob/main/CHANGELOG.md#expected-addresses-with-deterministic-deployment-proxy-default
        // NOTE: double check these addresses are correct on each chain
        address safeFactory = 0x4e1DCf7AD4e460CfD30791CCC4F9c8a4f820ec67;
        address safeSingleton = 0x41675C099F32341bf84BFc5382aF534df5C7461a;
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
}