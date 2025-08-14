// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {DeployDestinationChainProxies} from "./1-deployDestinationChainProxies.s.sol";
import {CrosschainDeployLib} from "script/releases/CrosschainDeployLib.sol";
import "../Env.sol";

/**
 * Purpose: Deploy implementation contracts for the TaskMailbox on destination chains using an EOA.
 */
contract DeployDestinationChainImpls is EOADeployer, DeployDestinationChainProxies {
    using Env for *;

    function _runAsEOA() internal override {
        // If we're not on a destination chain or we're on a version that already has these contracts deployed, we don't need to deploy any contracts
        if (!Env.isDestinationChain() || _isAlreadyDeployed()) {
            return;
        }

        vm.startBroadcast();

        // Deploy the implementation
        deployImpl({
            name: type(TaskMailbox).name,
            deployedTo: address(
                new TaskMailbox({
                    _bn254CertificateVerifier: address(Env.proxy.bn254CertificateVerifier()),
                    _ecdsaCertificateVerifier: address(Env.proxy.ecdsaCertificateVerifier()),
                    _maxTaskSLA: Env.MAX_TASK_SLA(),
                    _version: Env.deployVersion()
                })
            )
        });

        vm.stopBroadcast();
    }

    function testScript() public virtual override {
        if (!Env.isDestinationChain() || _isAlreadyDeployed()) {
            return;
        }

        // 1. Deploy destination chain proxies
        // Only deploy the proxies if they haven't been deployed yet
        /// @dev This is needed in the production environment tests since this step would fail if the proxies are already deployed
        if (!_areProxiesDeployed()) {
            DeployDestinationChainProxies._runAsMultisig();
            _unsafeResetHasPranked(); // reset hasPranked so we can use it in the execute()
        } else {
            // Since the proxies are already deployed, we need to update the env with the proper addresses
            _addContractsToEnv();
        }

        // 2. Deploy destination chain impls
        runAsEOA();

        // Validate the destination chain
        _validateImplConstructors();
        _validateImplsInitialized();
        _validateVersion();
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        /// TaskMailbox
        TaskMailbox taskMailbox = Env.impl.taskMailbox();
        address expectedBN254 = address(Env.proxy.bn254CertificateVerifier());
        address expectedECDSA = address(Env.proxy.ecdsaCertificateVerifier());

        require(expectedBN254 != address(0), "BN254CertificateVerifier proxy not found for validation");
        require(expectedECDSA != address(0), "ECDSACertificateVerifier proxy not found for validation");

        assertTrue(
            taskMailbox.BN254_CERTIFICATE_VERIFIER() == expectedBN254, "taskMailbox.BN254_CERTIFICATE_VERIFIER invalid"
        );
        assertTrue(
            taskMailbox.ECDSA_CERTIFICATE_VERIFIER() == expectedECDSA, "taskMailbox.ECDSA_CERTIFICATE_VERIFIER invalid"
        );
        assertEq(taskMailbox.MAX_TASK_SLA(), Env.MAX_TASK_SLA(), "taskMailbox.MAX_TASK_SLA failed");
        assertEq(taskMailbox.version(), Env.deployVersion(), "taskMailbox.version failed");
    }

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// TaskMailbox
        TaskMailbox taskMailbox = Env.impl.taskMailbox();

        vm.expectRevert(errInit);
        taskMailbox.initialize(
            address(0), // owner
            0, // feeSplit
            address(0) // feeSplitCollector
        );
    }

    function _validateVersion() internal view {
        string memory expected = Env.deployVersion();

        assertEq(Env.impl.taskMailbox().version(), expected, "taskMailbox version mismatch");
    }
}
