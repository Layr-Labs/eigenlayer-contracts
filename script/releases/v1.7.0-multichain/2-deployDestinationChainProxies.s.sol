// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";


/**
 * Purpose: Deploy proxy contracts for the destination chain using a multisig.
 */
contract DeployDestinationChainProxies is MultisigBuilder {
    using Env for *;

    /// forgefmt: disable-next-item
    function _runAsMultisig() internal override prank(Env.multichainDeployerMultisig()) {
        // If we're not on a destination chain, we don't need to deploy any contracts
        if (!Env.isDestinationChain()) {
            return;
        }

        vm.startBroadcast();

        // Deploy or get the empty contract for the destination chain
        // address emptyContract = address(Env.impl.emptyContract());

        // Deploy the proxies pointing to an empty contract
        // TODO: deploy using createX
        // deployProxy({
        //     name: type(OperatorTableUpdater).name,
        //     deployedTo: address(new TransparentUpgradeableProxy({_logic: emptyContract, admin_: Env.multichainDeployerMultisig(), _data: ""}))
        // });

        // deployProxy({
        //     name: type(ECDSACertificateVerifier).name,
        //     deployedTo: address(new TransparentUpgradeableProxy({_logic: emptyContract, admin_: Env.multichainDeployerMultisig(), _data: ""}))
        // });

        // deployProxy({
        //     name: type(BN254CertificateVerifier).name,
        //     deployedTo: address(new TransparentUpgradeableProxy({_logic: emptyContract, admin_: Env.multichainDeployerMultisig(), _data: ""}))
        // });

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        if (!Env.isDestinationChain()) {
            return;
        }

        _runAsMultisig();

        _validateProxiesDeployed();
        _validateProxyAdminIsMultisig();
    }

    /// @dev Validate that proxy contracts are deployed
    function _validateProxiesDeployed() internal view {
        assertTrue(address(Env.proxy.operatorTableUpdater()) != address(0), "operatorTableUpdater proxy not deployed");
        assertTrue(address(Env.proxy.ecdsaCertificateVerifier()) != address(0), "ecdsaCertificateVerifier proxy not deployed");
        assertTrue(address(Env.proxy.bn254CertificateVerifier()) != address(0), "bn254CertificateVerifier proxy not deployed");
    }

    /// @dev Validate that proxies are owned by the multisig (temporarily)
    function _validateProxyAdminIsMultisig() internal view {
        address multisig = Env.multichainDeployerMultisig();

        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.operatorTableUpdater())) == multisig,
            "operatorTableUpdater proxyAdmin should be multisig"
        );
        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.ecdsaCertificateVerifier())) == multisig,
            "ecdsaCertificateVerifier proxyAdmin should be multisig"
        );
        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.bn254CertificateVerifier())) == multisig,
            "bn254CertificateVerifier proxyAdmin should be multisig"
        );
    }
} 