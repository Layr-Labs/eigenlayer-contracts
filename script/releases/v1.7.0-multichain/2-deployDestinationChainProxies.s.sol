// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {CrosschainDeployLib} from "script/releases/CrosschainDeployLib.sol";
import "../Env.sol";

/**
 * Purpose: Deploy proxy contracts for the destination chain using a multisig.
 */
contract DeployDestinationChainProxies is MultisigBuilder {
    using Env for *;

    /// forgefmt: disable-next-item
    function _runAsMultisig() internal virtual override {
        // If we're not on a destination chain or we're on a version that already has these contracts deployed, we don't need to deploy any contracts
        if (!Env.isDestinationChain() || _isAlreadyDeployed()) {
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
        _unsafeAddImplContract(type(EmptyContract).name, emptyContract);
        _unsafeAddProxyContract(type(OperatorTableUpdater).name, address(operatorTableUpdaterProxy));
        _unsafeAddProxyContract(type(ECDSACertificateVerifier).name, address(ecdsaCertificateVerifierProxy));
        _unsafeAddProxyContract(type(BN254CertificateVerifier).name, address(bn254CertificateVerifierProxy));
    }

    function testScript() public virtual {
        if (!Env.isDestinationChain() || _isAlreadyDeployed()) {
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

    /// @dev Check if the proxies are deployed by checking if the empty contract is deployed
    function _areProxiesDeployed() internal view returns (bool) {
        address expectedEmptyContract = CrosschainDeployLib.computeCrosschainAddress({
            deployer: Env.multichainDeployerMultisig(),
            initCodeHash: keccak256(type(EmptyContract).creationCode),
            name: type(EmptyContract).name
        });

        // If the empty contract is deployed, then the proxies are deployed
        if (expectedEmptyContract.code.length > 0) {
            return true;
        }
        return false;
    }

    /// @dev Add the contracts to the env
    function _addContractsToEnv() internal {
        address emptyContract = CrosschainDeployLib.computeCrosschainAddress({
            deployer: Env.multichainDeployerMultisig(),
            initCodeHash: keccak256(type(EmptyContract).creationCode),
            name: type(EmptyContract).name
        });
        _unsafeAddProxyContract(
            type(OperatorTableUpdater).name,
            _computeExpectedProxyAddress(type(OperatorTableUpdater).name, emptyContract)
        );
        _unsafeAddProxyContract(
            type(ECDSACertificateVerifier).name,
            _computeExpectedProxyAddress(type(ECDSACertificateVerifier).name, emptyContract)
        );
        _unsafeAddProxyContract(
            type(BN254CertificateVerifier).name,
            _computeExpectedProxyAddress(type(BN254CertificateVerifier).name, emptyContract)
        );
    }

    /// @dev Compute the expected proxy address for a given name and empty contract
    function _computeExpectedProxyAddress(string memory name, address emptyContract) internal view returns (address) {
        return CrosschainDeployLib.computeCrosschainUpgradeableProxyAddress({
            adminAndDeployer: Env.multichainDeployerMultisig(),
            implementation: emptyContract,
            name: name
        });
    }

    /// @dev Check if the version is already deployed
    function _isAlreadyDeployed() internal view returns (bool) {
        if (keccak256(bytes(Env.envVersion())) == keccak256(bytes("1.7.0-rc.0"))) {
            return true;
        } else if (keccak256(bytes(Env.envVersion())) == keccak256(bytes("1.7.0-rc.1"))) {
            return true;
        } else if (keccak256(bytes(Env.envVersion())) == keccak256(bytes("1.7.0-rc.2"))) {
            return true;
        } else if (keccak256(bytes(Env.envVersion())) == keccak256(bytes("1.8.0-rc.0"))) {
            return true;
        }
        return false;
    }
}
