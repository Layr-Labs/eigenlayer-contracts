// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * Purpose: use an EOA to deploy all of the new destination chain contracts for this upgrade.
 */
contract DeployDestinationChain is EOADeployer {
    using Env for *;

    /// forgefmt: disable-next-item
    function _runAsEOA() internal override {
        // If we're not on a destination chain, we don't need to deploy any contracts
        if (!Env.isDestinationChain()) {
            return;
        }

        vm.startBroadcast();

        // Deploy OperatorTableUpdater implementation
        deployImpl({
            name: type(OperatorTableUpdater).name,
            deployedTo: address(
                new OperatorTableUpdater({
                    _bn254CertificateVerifier: Env.proxy.bn254CertificateVerifier(),
                    _ecdsaCertificateVerifier: Env.proxy.ecdsaCertificateVerifier(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _version: Env.deployVersion()
                })
            )
        });

        // Deploy ECDSACertificateVerifier implementation
        deployImpl({
            name: type(ECDSACertificateVerifier).name,
            deployedTo: address(
                new ECDSACertificateVerifier({
                    _operatorTableUpdater: Env.proxy.operatorTableUpdater(),
                    _version: Env.deployVersion()
                })
            )
        });

        // Deploy BN254CertificateVerifier implementation
        deployImpl({
            name: type(BN254CertificateVerifier).name,
            deployedTo: address(
                new BN254CertificateVerifier({
                    _operatorTableUpdater: Env.proxy.operatorTableUpdater(),
                    _version: Env.deployVersion()
                })
            )
        });

        // Deploy TaskMailbox implementation
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

    function testScript() public virtual {
        if (!Env.isDestinationChain()) {
            return;
        }

        // Set the mode to EOA so we can deploy the contracts
        runAsEOA();

        _validateNewImplAddresses({areMatching: false});
        _validateProxyAdmins();
        _validateImplConstructors();
        _validateImplsInitialized();
        _validateVersion();
    }

    /// @dev Validate that the `Env.impl` addresses are updated to be distinct from what the proxy
    /// admin reports as the current implementation address.
    ///
    /// Note: The upgrade script can call this with `areMatching == true` to check that these impl
    /// addresses _are_ matches.
    function _validateNewImplAddresses(
        bool areMatching
    ) internal view {
        function (bool, string memory) internal pure assertion = areMatching ? _assertTrue : _assertFalse;

        // OperatorTableUpdater
        assertion(
            Env._getProxyImpl(address(Env.proxy.operatorTableUpdater())) == address(Env.impl.operatorTableUpdater()),
            "operatorTableUpdater impl failed"
        );

        // ECDSACertificateVerifier
        assertion(
            Env._getProxyImpl(address(Env.proxy.ecdsaCertificateVerifier()))
                == address(Env.impl.ecdsaCertificateVerifier()),
            "ecdsaCertificateVerifier impl failed"
        );

        // BN254CertificateVerifier
        assertion(
            Env._getProxyImpl(address(Env.proxy.bn254CertificateVerifier()))
                == address(Env.impl.bn254CertificateVerifier()),
            "bn254CertificateVerifier impl failed"
        );

        // TaskMailbox
        assertion(
            Env._getProxyImpl(address(Env.proxy.taskMailbox())) == address(Env.impl.taskMailbox()),
            "taskMailbox impl failed"
        );
    }

    /// @dev Ensure each deployed TUP/beacon is owned by the proxyAdmin/executorMultisig
    function _validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();

        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.operatorTableUpdater())) == pa,
            "operatorTableUpdater proxyAdmin incorrect"
        );
        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.ecdsaCertificateVerifier())) == pa,
            "ecdsaCertificateVerifier proxyAdmin incorrect"
        );
        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.bn254CertificateVerifier())) == pa,
            "bn254CertificateVerifier proxyAdmin incorrect"
        );
        assertTrue(Env._getProxyAdmin(address(Env.proxy.taskMailbox())) == pa, "taskMailbox proxyAdmin incorrect");
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

        {
            /// TaskMailbox
            TaskMailbox taskMailbox = Env.impl.taskMailbox();
            assertTrue(
                taskMailbox.BN254_CERTIFICATE_VERIFIER() == address(Env.proxy.bn254CertificateVerifier()),
                "taskMailbox.BN254_CERTIFICATE_VERIFIER invalid"
            );
            assertTrue(
                taskMailbox.ECDSA_CERTIFICATE_VERIFIER() == address(Env.proxy.ecdsaCertificateVerifier()),
                "taskMailbox.ECDSA_CERTIFICATE_VERIFIER invalid"
            );
            assertEq(taskMailbox.MAX_TASK_SLA(), Env.MAX_TASK_SLA(), "taskMailbox.MAX_TASK_SLA failed");
            assertEq(taskMailbox.version(), Env.deployVersion(), "taskMailbox.version failed");
        }
    }

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// OperatorTableUpdater - dummy parameters
        OperatorTableUpdater operatorTableUpdater = Env.impl.operatorTableUpdater();
        OperatorSet memory dummyOperatorSet = OperatorSet({avs: address(0), id: 0});
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory dummyBN254Info;

        vm.expectRevert(errInit);
        operatorTableUpdater.initialize(
            address(0), // owner
            0, // initial paused status
            dummyOperatorSet, // globalRootConfirmerSet
            0, // globalRootConfirmationThreshold
            dummyBN254Info // globalRootConfirmerSetInfo
        );

        /// TaskMailbox - dummy parameters
        TaskMailbox taskMailbox = Env.impl.taskMailbox();
        vm.expectRevert(errInit);
        taskMailbox.initialize(
            address(0), // owner
            0, // feeSplit
            address(0) // feeSplitCollector
        );

        // ECDSACertificateVerifier and BN254CertificateVerifier don't have initialize functions
    }

    function _validateVersion() internal view {
        // On future upgrades, just tick the major/minor/patch to validate
        string memory expected = Env.deployVersion();

        assertEq(Env.impl.operatorTableUpdater().version(), expected, "operatorTableUpdater version mismatch");
        assertEq(Env.impl.ecdsaCertificateVerifier().version(), expected, "ecdsaCertificateVerifier version mismatch");
        assertEq(Env.impl.bn254CertificateVerifier().version(), expected, "bn254CertificateVerifier version mismatch");
        assertEq(Env.impl.taskMailbox().version(), expected, "taskMailbox version mismatch");
    }

    function _assertTrue(bool b, string memory err) internal pure {
        assertTrue(b, err);
    }

    function _assertFalse(bool b, string memory err) internal pure {
        assertFalse(b, err);
    }
}
