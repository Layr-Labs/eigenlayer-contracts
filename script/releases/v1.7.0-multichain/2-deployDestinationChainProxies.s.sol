// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ScriptHelpers} from "zeus-templates/utils/ScriptHelpers.sol";
import {ZEnvHelpers} from "zeus-templates/utils/ZEnvHelpers.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {CrosschainDeployLib} from "script/releases/CrosschainDeployLib.sol";
import "forge-std/console.sol";
import "../Env.sol";

/**
 * Purpose: Deploy proxy contracts for the destination chain using a multisig.
 */
contract DeployDestinationChainProxies is MultisigBuilder {
    using ZEnvHelpers for *;
    using ScriptHelpers for *;
    using Env for *;

    /// forgefmt: disable-next-item
    function _runAsMultisig() internal virtual override {
        // If we're not on a destination chain, we don't need to deploy any contracts
        if (!Env.isDestinationChain()) {
            return;
        }

        // We don't use the prank modifier here, since we have to write to the env
        _startPrank(Env.multichainDeployerMultisig());

        // Deploy empty contract
        address emptyContract = CrosschainDeployLib.deployEmptyContract(Env.multichainDeployerMultisig());

        // Deploy the proxies pointing to an empty contract
        ITransparentUpgradeableProxy operatorTableUpdaterProxy = CrosschainDeployLib.deployCrosschainProxy({
            implementation: emptyContract,
            adminAndDeployer: Env.multichainDeployerMultisig(),
            name: type(OperatorTableUpdater).name
        });

        ITransparentUpgradeableProxy ecdsaCertificateVerifierProxy = CrosschainDeployLib.deployCrosschainProxy({
            implementation: emptyContract,
            adminAndDeployer: Env.multichainDeployerMultisig(),
            name: type(ECDSACertificateVerifier).name
        });

        ITransparentUpgradeableProxy bn254CertificateVerifierProxy = CrosschainDeployLib.deployCrosschainProxy({
            implementation: emptyContract,
            adminAndDeployer: Env.multichainDeployerMultisig(),
            name: type(BN254CertificateVerifier).name
        });

        // Stop pranking
        _stopPrank();

        // Save all the contracts to the env
        // NOTE: This is an antipattern, we should update the ZEnvHelpers to support this
        ZEnvHelpers.state().__updateContract(type(EmptyContract).name.impl(), address(emptyContract));
        ZEnvHelpers.state().__updateContract(type(OperatorTableUpdater).name.proxy(), address(operatorTableUpdaterProxy));
        ZEnvHelpers.state().__updateContract(type(ECDSACertificateVerifier).name.proxy(), address(ecdsaCertificateVerifierProxy));
        ZEnvHelpers.state().__updateContract(type(BN254CertificateVerifier).name.proxy(), address(bn254CertificateVerifierProxy));
    }

    function testScript() public virtual {
        if (!Env.isDestinationChain()) {
            return;
        }

        execute();

        _validateProxyAdminIsMultisig();
    }

    /// @dev Validate that proxies are owned by the multichain deployer multisig (temporarily)
    function _validateProxyAdminIsMultisig() internal view {
        address multisig = Env.multichainDeployerMultisig();

        assertTrue(
            _getProxyAdminBySlot(address(Env.proxy.operatorTableUpdater())) == multisig,
            "operatorTableUpdater proxyAdmin should be multisig"
        );
        assertTrue(
            _getProxyAdminBySlot(address(Env.proxy.ecdsaCertificateVerifier())) == multisig,
            "ecdsaCertificateVerifier proxyAdmin should be multisig"
        );
        assertTrue(
            _getProxyAdminBySlot(address(Env.proxy.bn254CertificateVerifier())) == multisig,
            "bn254CertificateVerifier proxyAdmin should be multisig"
        );
    }

    /// @dev We have to use the slot directly since _getProxyAdmin expects the caller to be the actual proxy admin
    function _getProxyAdminBySlot(
        address _proxy
    ) internal view returns (address) {
        bytes32 adminSlot = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
        address admin = address(uint160(uint256(vm.load(address(_proxy), adminSlot))));
        return admin;
    }
}
