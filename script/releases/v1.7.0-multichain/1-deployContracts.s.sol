// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

/**
 * Purpose: use an EOA to deploy all of the new contracts for this upgrade.
 */
contract Deploy is EOADeployer {
    using Env for *;

    /// forgefmt: disable-next-item
    function _runAsEOA() internal override {
        vm.startBroadcast();

        _deploySourceChainContracts();
        _deployDestinationChainContracts();

        vm.stopBroadcast();
    }

    function _deploySourceChainContracts() internal {
        // If we're not on a source chain, we don't need to deploy any contracts
        if (!Env.isSourceChain()) {
            return;
        }

        // Deploy KeyRegistrar implementation
        deployImpl({
            name: type(KeyRegistrar).name,
            deployedTo: address(
                new KeyRegistrar({
                    _permissionController: Env.proxy.permissionController(),
                    _allocationManager: Env.proxy.allocationManager(),
                    _version: Env.deployVersion()
                })
            )
        });

        // Deploy KeyRegistrar proxy
        deployProxy({
            name: type(KeyRegistrar).name,
            deployedTo: address(
                new TransparentUpgradeableProxy({
                    _logic: address(Env.impl.keyRegistrar()),
                    admin_: Env.proxyAdmin(),
                    _data: "" // No initialization needed for KeyRegistrar
                })
            )
        });

        // Deploy CrossChainRegistry implementation
        deployImpl({
            name: type(CrossChainRegistry).name,
            deployedTo: address(
                new CrossChainRegistry({
                    _allocationManager: Env.proxy.allocationManager(),
                    _keyRegistrar: Env.proxy.keyRegistrar(),
                    _permissionController: Env.proxy.permissionController(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _version: Env.deployVersion()
                })
            )
        });

        // Deploy CrossChainRegistry proxy
        deployProxy({
            name: type(CrossChainRegistry).name,
            deployedTo: address(
                new TransparentUpgradeableProxy({
                    _logic: address(Env.impl.crossChainRegistry()),
                    admin_: Env.proxyAdmin(),
                    _data: abi.encodeCall(
                        CrossChainRegistry.initialize,
                        (
                            Env.opsMultisig(), // initialOwner
                            Env.CROSS_CHAIN_REGISTRY_PAUSE_STATUS()
                        )
                    )
                })
            )
        });

        // Deploy ReleaseManager implementation
        deployImpl({
            name: type(ReleaseManager).name,
            deployedTo: address(
                new ReleaseManager({
                    _permissionController: Env.proxy.permissionController(),
                    _version: Env.deployVersion()
                })
            )
        });

        // Deploy ReleaseManager proxy
        deployProxy({
            name: type(ReleaseManager).name,
            deployedTo: address(
                new TransparentUpgradeableProxy({
                    _logic: address(Env.impl.releaseManager()),
                    admin_: Env.proxyAdmin(),
                    _data: "" // No initialize function for ReleaseManager
                })
            )
        });
    }

    function _deployDestinationChainContracts() internal {
        // If we're not on a destination chain, we don't need to deploy any contracts
        if (!Env.isDestinationChain()) {
            return;
        }

        // Deploy or get the empty contract for the destination chain
        address emptyContract = address(Env.impl.emptyContract());

        // 1. Deploy the proxies pointing to an empty contract
        deployProxy({
            name: type(OperatorTableUpdater).name,
            deployedTo: address(new TransparentUpgradeableProxy({_logic: emptyContract, admin_: msg.sender, _data: ""}))
        });

        deployProxy({
            name: type(ECDSACertificateVerifier).name,
            deployedTo: address(new TransparentUpgradeableProxy({_logic: emptyContract, admin_: msg.sender, _data: ""}))
        });

        deployProxy({
            name: type(BN254CertificateVerifier).name,
            deployedTo: address(new TransparentUpgradeableProxy({_logic: emptyContract, admin_: msg.sender, _data: ""}))
        });

        // 2. Deploy the implementations
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

        // 3. Upgrade the proxies to point to the new implementations
        // TODO: remove these dummy parameters
        OperatorSet memory dummyOperatorSet = OperatorSet({avs: address(0), id: 0});
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory dummyBN254Info;
        ICrossChainRegistry.OperatorSetConfig memory dummyConfig;

        ITransparentUpgradeableProxy operatorTableUpdaterProxy =
            ITransparentUpgradeableProxy(payable(address(Env.proxy.operatorTableUpdater())));
        operatorTableUpdaterProxy.upgradeToAndCall(
            address(Env.impl.operatorTableUpdater()),
            abi.encodeCall(
                OperatorTableUpdater.initialize,
                (Env.opsMultisig(), dummyOperatorSet, 0, 0, dummyBN254Info, dummyConfig)
            )
        );

        ITransparentUpgradeableProxy ecdsaCertificateVerifierProxy =
            ITransparentUpgradeableProxy(payable(address(Env.proxy.ecdsaCertificateVerifier())));
        ecdsaCertificateVerifierProxy.upgradeTo(address(Env.impl.ecdsaCertificateVerifier()));

        ITransparentUpgradeableProxy bn254CertificateVerifierProxy =
            ITransparentUpgradeableProxy(payable(address(Env.proxy.bn254CertificateVerifier())));
        bn254CertificateVerifierProxy.upgradeTo(address(Env.impl.bn254CertificateVerifier()));

        // 4. Transfer proxy admin ownership
        ProxyAdmin(msg.sender).changeProxyAdmin(operatorTableUpdaterProxy, Env.proxyAdmin());
        ProxyAdmin(msg.sender).changeProxyAdmin(ecdsaCertificateVerifierProxy, Env.proxyAdmin());
        ProxyAdmin(msg.sender).changeProxyAdmin(bn254CertificateVerifierProxy, Env.proxyAdmin());
    }

    function testScript() public virtual {
        _runAsEOA();

        _validateStorage();
        _validateProxyAdmins();
        _validateImplConstructors();
        _validateImplsInitialized();
        _validateVersion();
        _validateProxyConstructors();
        _validateProxiesInitialized();
    }

    /// @dev Validate that storage variables are set correctly
    function _validateStorage() internal view {
        if (Env.isSourceChain()) {
            // Validate KeyRegistrar
            KeyRegistrar keyRegistrar = Env.proxy.keyRegistrar();
            assertTrue(address(keyRegistrar) != address(0), "keyRegistrar not deployed");

            // Validate CrossChainRegistry
            CrossChainRegistry crossChainRegistry = Env.proxy.crossChainRegistry();
            assertTrue(crossChainRegistry.owner() == Env.opsMultisig(), "ccr.owner invalid");
            assertTrue(crossChainRegistry.paused() == Env.CROSS_CHAIN_REGISTRY_PAUSE_STATUS(), "ccr.paused invalid");

            // Validate ReleaseManager
            ReleaseManager releaseManager = Env.proxy.releaseManager();
            assertTrue(address(releaseManager) != address(0), "releaseManager not deployed");
        }

        if (Env.isDestinationChain()) {
            // Validate OperatorTableUpdater
            OperatorTableUpdater operatorTableUpdater = Env.proxy.operatorTableUpdater();
            assertTrue(operatorTableUpdater.owner() == Env.opsMultisig(), "otu.owner invalid");
            // TODO: add checks on global root confirmer set

            // Validate ECDSACertificateVerifier
            ECDSACertificateVerifier ecdsaCertificateVerifier = Env.proxy.ecdsaCertificateVerifier();
            assertTrue(address(ecdsaCertificateVerifier) != address(0), "ecdsaCertificateVerifier not deployed");

            // Validate BN254CertificateVerifier
            BN254CertificateVerifier bn254CertificateVerifier = Env.proxy.bn254CertificateVerifier();
            assertTrue(address(bn254CertificateVerifier) != address(0), "bn254CertificateVerifier not deployed");
        }
    }

    /// @dev Ensure each deployed TUP/beacon is owned by the proxyAdmin/executorMultisig
    function _validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();

        if (Env.isSourceChain()) {
            assertTrue(Env._getProxyAdmin(address(Env.proxy.keyRegistrar())) == pa, "keyRegistrar proxyAdmin incorrect");
            assertTrue(
                Env._getProxyAdmin(address(Env.proxy.crossChainRegistry())) == pa,
                "crossChainRegistry proxyAdmin incorrect"
            );
            assertTrue(
                Env._getProxyAdmin(address(Env.proxy.releaseManager())) == pa, "releaseManager proxyAdmin incorrect"
            );
        }

        if (Env.isDestinationChain()) {
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
        }
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        if (Env.isSourceChain()) {
            {
                /// KeyRegistrar
                KeyRegistrar keyRegistrar = Env.impl.keyRegistrar();
                assertTrue(
                    address(keyRegistrar.permissionController()) == address(Env.proxy.permissionController()),
                    "kr.permissionController invalid"
                );
                assertTrue(
                    address(keyRegistrar.allocationManager()) == address(Env.proxy.allocationManager()),
                    "kr.allocationManager invalid"
                );
                assertEq(keyRegistrar.version(), Env.deployVersion(), "kr.version failed");
            }

            {
                /// CrossChainRegistry
                CrossChainRegistry crossChainRegistry = Env.impl.crossChainRegistry();
                assertTrue(
                    address(crossChainRegistry.allocationManager()) == address(Env.proxy.allocationManager()),
                    "ccr.allocationManager invalid"
                );
                assertTrue(
                    address(crossChainRegistry.keyRegistrar()) == address(Env.proxy.keyRegistrar()),
                    "ccr.keyRegistrar invalid"
                );
                assertTrue(
                    address(crossChainRegistry.permissionController()) == address(Env.proxy.permissionController()),
                    "ccr.permissionController invalid"
                );
                assertTrue(
                    address(crossChainRegistry.pauserRegistry()) == address(Env.impl.pauserRegistry()),
                    "ccr.pauserRegistry invalid"
                );
                assertEq(crossChainRegistry.version(), Env.deployVersion(), "ccr.version failed");
            }

            {
                /// ReleaseManager
                ReleaseManager releaseManager = Env.impl.releaseManager();
                assertTrue(
                    releaseManager.permissionController() == Env.proxy.permissionController(),
                    "rm.permissionController invalid"
                );
                assertEq(releaseManager.version(), Env.deployVersion(), "rm.version failed");
            }
        }

        if (Env.isDestinationChain()) {
            {
                /// OperatorTableUpdater
                OperatorTableUpdater operatorTableUpdater = Env.impl.operatorTableUpdater();
                assertTrue(
                    address(operatorTableUpdater.bn254CertificateVerifier())
                        == address(Env.proxy.bn254CertificateVerifier()),
                    "otu.bn254CertificateVerifier invalid"
                );
                assertTrue(
                    address(operatorTableUpdater.ecdsaCertificateVerifier())
                        == address(Env.proxy.ecdsaCertificateVerifier()),
                    "otu.ecdsaCertificateVerifier invalid"
                );
                assertEq(operatorTableUpdater.version(), Env.deployVersion(), "otu.version failed");
            }

            {
                /// ECDSACertificateVerifier
                ECDSACertificateVerifier ecdsaCertificateVerifier = Env.impl.ecdsaCertificateVerifier();
                assertTrue(
                    address(ecdsaCertificateVerifier.operatorTableUpdater())
                        == address(Env.proxy.operatorTableUpdater()),
                    "ecv.operatorTableUpdater invalid"
                );
                assertEq(ecdsaCertificateVerifier.version(), Env.deployVersion(), "ecv.version failed");
            }

            {
                /// BN254CertificateVerifier
                BN254CertificateVerifier bn254CertificateVerifier = Env.impl.bn254CertificateVerifier();
                assertTrue(
                    address(bn254CertificateVerifier.operatorTableUpdater())
                        == address(Env.proxy.operatorTableUpdater()),
                    "b254cv.operatorTableUpdater invalid"
                );
                assertEq(bn254CertificateVerifier.version(), Env.deployVersion(), "b254cv.version failed");
            }
        }
    }

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        if (Env.isSourceChain()) {
            /// CrossChainRegistry
            CrossChainRegistry crossChainRegistry = Env.impl.crossChainRegistry();
            vm.expectRevert(errInit);
            crossChainRegistry.initialize(address(0), 0);
        }

        if (Env.isDestinationChain()) {
            /// OperatorTableUpdater - dummy parameters
            OperatorTableUpdater operatorTableUpdater = Env.impl.operatorTableUpdater();
            OperatorSet memory dummyOperatorSet = OperatorSet({avs: address(0), id: 0});
            IBN254TableCalculatorTypes.BN254OperatorSetInfo memory dummyBN254Info;
            ICrossChainRegistry.OperatorSetConfig memory dummyConfig;

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
    }

    function _validateVersion() internal view {
        string memory expected = Env.deployVersion();

        if (Env.isSourceChain()) {
            assertEq(Env.impl.keyRegistrar().version(), expected, "keyRegistrar version mismatch");
            assertEq(Env.impl.crossChainRegistry().version(), expected, "crossChainRegistry version mismatch");
            assertEq(Env.impl.releaseManager().version(), expected, "releaseManager version mismatch");
        }

        if (Env.isDestinationChain()) {
            assertEq(Env.impl.operatorTableUpdater().version(), expected, "operatorTableUpdater version mismatch");
            assertEq(
                Env.impl.ecdsaCertificateVerifier().version(), expected, "ecdsaCertificateVerifier version mismatch"
            );
            assertEq(
                Env.impl.bn254CertificateVerifier().version(), expected, "bn254CertificateVerifier version mismatch"
            );
        }
    }

    function _validateProxyConstructors() internal view {
        if (Env.isSourceChain()) {
            KeyRegistrar keyRegistrar = Env.proxy.keyRegistrar();
            assertEq(keyRegistrar.version(), Env.deployVersion(), "keyRegistrar version mismatch");
            assertTrue(
                keyRegistrar.permissionController() == Env.proxy.permissionController(),
                "keyRegistrar permissionController mismatch"
            );
            assertTrue(
                keyRegistrar.allocationManager() == Env.proxy.allocationManager(),
                "keyRegistrar allocationManager mismatch"
            );

            CrossChainRegistry crossChainRegistry = Env.proxy.crossChainRegistry();
            assertEq(crossChainRegistry.version(), Env.deployVersion(), "crossChainRegistry version mismatch");
            assertTrue(
                crossChainRegistry.allocationManager() == Env.proxy.allocationManager(),
                "crossChainRegistry allocationManager mismatch"
            );
            assertTrue(
                crossChainRegistry.keyRegistrar() == Env.proxy.keyRegistrar(),
                "crossChainRegistry keyRegistrar mismatch"
            );
            assertTrue(
                crossChainRegistry.permissionController() == Env.proxy.permissionController(),
                "crossChainRegistry permissionController mismatch"
            );
            assertTrue(
                crossChainRegistry.pauserRegistry() == Env.impl.pauserRegistry(),
                "crossChainRegistry pauserRegistry mismatch"
            );

            ReleaseManager releaseManager = Env.proxy.releaseManager();
            assertEq(releaseManager.version(), Env.deployVersion(), "releaseManager version mismatch");
            assertTrue(
                releaseManager.permissionController() == Env.proxy.permissionController(),
                "releaseManager permissionController mismatch"
            );
        }

        if (Env.isDestinationChain()) {
            OperatorTableUpdater operatorTableUpdater = Env.proxy.operatorTableUpdater();
            assertEq(operatorTableUpdater.version(), Env.deployVersion(), "operatorTableUpdater version mismatch");
            assertTrue(
                operatorTableUpdater.bn254CertificateVerifier() == Env.proxy.bn254CertificateVerifier(),
                "otu.bn254CertificateVerifier mismatch"
            );
            assertTrue(
                operatorTableUpdater.ecdsaCertificateVerifier() == Env.proxy.ecdsaCertificateVerifier(),
                "otu.ecdsaCertificateVerifier mismatch"
            );

            ECDSACertificateVerifier ecdsaCertificateVerifier = Env.proxy.ecdsaCertificateVerifier();
            assertEq(
                ecdsaCertificateVerifier.version(), Env.deployVersion(), "ecdsaCertificateVerifier version mismatch"
            );
            assertTrue(
                ecdsaCertificateVerifier.operatorTableUpdater() == Env.proxy.operatorTableUpdater(),
                "ecdsaCertificateVerifier operatorTableUpdater mismatch"
            );

            BN254CertificateVerifier bn254CertificateVerifier = Env.proxy.bn254CertificateVerifier();
            assertEq(
                bn254CertificateVerifier.version(), Env.deployVersion(), "bn254CertificateVerifier version mismatch"
            );
            assertTrue(
                bn254CertificateVerifier.operatorTableUpdater() == Env.proxy.operatorTableUpdater(),
                "bn254CertificateVerifier operatorTableUpdater mismatch"
            );
        }
    }

    function _validateProxiesInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        if (Env.isSourceChain()) {
            /// CrossChainRegistry
            CrossChainRegistry crossChainRegistry = Env.proxy.crossChainRegistry();
            vm.expectRevert(errInit);
            crossChainRegistry.initialize(address(0), 0);

            // ReleaseManager and KeyRegistrar don't have initialize functions
        }

        if (Env.isDestinationChain()) {
            /// OperatorTableUpdater - dummy parameters
            OperatorTableUpdater operatorTableUpdater = Env.proxy.operatorTableUpdater();
            OperatorSet memory dummyOperatorSet = OperatorSet({avs: address(0), id: 0});
            IBN254TableCalculatorTypes.BN254OperatorSetInfo memory dummyBN254Info;
            ICrossChainRegistry.OperatorSetConfig memory dummyConfig;

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
    }

    function _assertTrue(bool b, string memory err) private pure {
        assertTrue(b, err);
    }

    function _assertFalse(bool b, string memory err) private pure {
        assertFalse(b, err);
    }

    function _strEq(string memory a, string memory b) private pure returns (bool) {
        return keccak256(bytes(a)) == keccak256(bytes(b));
    }
}
