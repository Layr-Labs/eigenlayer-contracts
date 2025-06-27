// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {DeployDestinationChainProxies} from "./2-deployDestinationChainProxies.s.sol";
import "../Env.sol";

/**
 * Purpose: Deploy implementation contracts for the destination chain using an EOA.
 */
contract DeployDestinationChainImpls is EOADeployer, DeployDestinationChainProxies {
    using Env for *;

    function _runAsEOA() internal override {
        // If we're not on a destination chain, we don't need to deploy any contracts
        if (!Env.isDestinationChain()) {
            return;
        }

        vm.startBroadcast();

        // Deploy the implementations
        deployImpl({
            name: type(OperatorTableUpdater).name,
            deployedTo: address(
                new OperatorTableUpdater({
                    _bn254CertificateVerifier: Env.proxy.bn254CertificateVerifier(),
                    _ecdsaCertificateVerifier: Env.proxy.ecdsaCertificateVerifier(),
                    _version: Env.deployVersion()
                })
            )
        });

        deployImpl({
            name: type(ECDSACertificateVerifier).name,
            deployedTo: address(
                new ECDSACertificateVerifier({
                    _operatorTableUpdater: Env.proxy.operatorTableUpdater(),
                    _version: Env.deployVersion()
                })
            )
        });

        deployImpl({
            name: type(BN254CertificateVerifier).name,
            deployedTo: address(
                new BN254CertificateVerifier({
                    _operatorTableUpdater: Env.proxy.operatorTableUpdater(),
                    _version: Env.deployVersion()
                })
            )
        });

        vm.stopBroadcast();
    }

    function testScript() public virtual override {
        if (!Env.isDestinationChain()) {
            return;
        }

        _runAsEOA();
    
        _validateImplsDeployed();
        _validateImplConstructors();
        _validateImplsInitialized();
        _validateVersion();
    }

    /// @dev Validate that implementation contracts are deployed
    function _validateImplsDeployed() internal view {
        assertTrue(address(Env.impl.operatorTableUpdater()) != address(0), "operatorTableUpdater impl not deployed");
        assertTrue(address(Env.impl.ecdsaCertificateVerifier()) != address(0), "ecdsaCertificateVerifier impl not deployed");
        assertTrue(address(Env.impl.bn254CertificateVerifier()) != address(0), "bn254CertificateVerifier impl not deployed");
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        {
            /// OperatorTableUpdater
            OperatorTableUpdater operatorTableUpdater = Env.impl.operatorTableUpdater();
            assertTrue(
                address(operatorTableUpdater.bn254CertificateVerifier())
                    == address(Env.proxy.bn254CertificateVerifier()),
                "out.bn254CertificateVerifier invalid"
            );
            assertTrue(
                address(operatorTableUpdater.ecdsaCertificateVerifier())
                    == address(Env.proxy.ecdsaCertificateVerifier()),
                "out.ecdsaCertificateVerifier invalid"
            );
            assertEq(operatorTableUpdater.version(), Env.deployVersion(), "out.version failed");
        }

        {
            /// ECDSACertificateVerifier
            ECDSACertificateVerifier ecdsaCertificateVerifier = Env.impl.ecdsaCertificateVerifier();
            assertTrue(
                address(ecdsaCertificateVerifier.operatorTableUpdater()) == address(Env.proxy.operatorTableUpdater()),
                "ecv.operatorTableUpdater invalid"
            );
            assertEq(ecdsaCertificateVerifier.version(), Env.deployVersion(), "ecv.version failed");
        }

        {
            /// BN254CertificateVerifier
            BN254CertificateVerifier bn254CertificateVerifier = Env.impl.bn254CertificateVerifier();
            assertTrue(
                address(bn254CertificateVerifier.operatorTableUpdater()) == address(Env.proxy.operatorTableUpdater()),
                "b254cv.operatorTableUpdater invalid"
            );
            assertEq(bn254CertificateVerifier.version(), Env.deployVersion(), "b254cv.version failed");
        }
    }

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// OperatorTableUpdater - dummy parameters
        OperatorTableUpdater operatorTableUpdater = Env.impl.operatorTableUpdater();
        OperatorSet memory dummyOperatorSet = OperatorSet({avs: address(0), id: 0});
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory dummyBN254Info;
        ICrossChainRegistryTypes.OperatorSetConfig memory dummyConfig;

        vm.expectRevert(errInit);
        operatorTableUpdater.initialize(
            address(0), // owner
            dummyOperatorSet, // globalRootConfirmerSet
            0, // globalRootConfirmationThreshold
            0, // referenceTimestamp
            dummyBN254Info, // globalRootConfirmerSetInfo
            dummyConfig // globalRootConfirmerSetConfig
        );

        // ECDSACertificateVerifier and BN254CertificateVerifier don't have initialize functions
    }

    function _validateVersion() internal view {
        string memory expected = Env.deployVersion();

        assertEq(Env.impl.operatorTableUpdater().version(), expected, "operatorTableUpdater version mismatch");
        assertEq(Env.impl.ecdsaCertificateVerifier().version(), expected, "ecdsaCertificateVerifier version mismatch");
        assertEq(Env.impl.bn254CertificateVerifier().version(), expected, "bn254CertificateVerifier version mismatch");
    }
} 