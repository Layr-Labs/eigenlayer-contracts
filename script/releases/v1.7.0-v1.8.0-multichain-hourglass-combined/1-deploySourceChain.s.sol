// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

/**
 * Purpose: use an EOA to deploy all of the new source chain contracts for this upgrade.
 */
contract DeploySourceChain is EOADeployer {
    using Env for *;

    /// forgefmt: disable-next-item
    function _runAsEOA() internal override {
        // If we're not on a source chain or we're on a version that already has these contracts deployed, we don't need to deploy any contracts
        if (!Env.isSourceChain() || _isAlreadyDeployedSourceChain()) {
            return;
        }

        vm.startBroadcast();

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
                            Env.TABLE_UPDATE_CADENCE(),
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
                new ReleaseManager({_permissionController: Env.proxy.permissionController(), _version: Env.deployVersion()})
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

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        if (!Env.isSourceChain() || _isAlreadyDeployedSourceChain()) {
            return;
        }

        // Set the mode to EOA so we can deploy the contracts
        super.runAsEOA();

        _validateStorage();
        _validateProxyAdmins();
        _validateImplConstructors();
        _validateImplsInitialized();
        _validateProxyConstructors();
        _validateProxiesInitialized();
    }

    /// @dev Validate that storage variables are set correctly
    function _validateStorage() internal view {
        // Validate KeyRegistrar
        KeyRegistrar keyRegistrar = Env.proxy.keyRegistrar();
        assertTrue(address(keyRegistrar) != address(0), "keyRegistrar not deployed");

        // Validate CrossChainRegistry
        CrossChainRegistry crossChainRegistry = Env.proxy.crossChainRegistry();
        assertTrue(crossChainRegistry.owner() == Env.opsMultisig(), "ccr.owner invalid");
        assertTrue(crossChainRegistry.paused() == Env.CROSS_CHAIN_REGISTRY_PAUSE_STATUS(), "ccr.paused invalid");
        assertEq(
            crossChainRegistry.getTableUpdateCadence(), Env.TABLE_UPDATE_CADENCE(), "ccr.tableUpdateCadence invalid"
        );

        // Validate ReleaseManager
        ReleaseManager releaseManager = Env.proxy.releaseManager();
        assertTrue(address(releaseManager) != address(0), "releaseManager not deployed");
    }

    /// @dev Ensure each deployed TUP/beacon is owned by the proxyAdmin/executorMultisig
    function _validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();

        assertTrue(Env._getProxyAdmin(address(Env.proxy.keyRegistrar())) == pa, "keyRegistrar proxyAdmin incorrect");
        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.crossChainRegistry())) == pa, "crossChainRegistry proxyAdmin incorrect"
        );
        assertTrue(Env._getProxyAdmin(address(Env.proxy.releaseManager())) == pa, "releaseManager proxyAdmin incorrect");
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
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

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// CrossChainRegistry
        CrossChainRegistry crossChainRegistry = Env.impl.crossChainRegistry();
        vm.expectRevert(errInit);
        crossChainRegistry.initialize(address(0), 1 days, 0);
    }

    function _validateProxyConstructors() internal view {
        KeyRegistrar keyRegistrar = Env.proxy.keyRegistrar();
        assertEq(keyRegistrar.version(), Env.deployVersion(), "keyRegistrar version mismatch");
        assertTrue(
            keyRegistrar.permissionController() == Env.proxy.permissionController(),
            "keyRegistrar permissionController mismatch"
        );
        assertTrue(
            keyRegistrar.allocationManager() == Env.proxy.allocationManager(), "keyRegistrar allocationManager mismatch"
        );

        CrossChainRegistry crossChainRegistry = Env.proxy.crossChainRegistry();
        assertEq(crossChainRegistry.version(), Env.deployVersion(), "crossChainRegistry version mismatch");
        assertTrue(
            crossChainRegistry.allocationManager() == Env.proxy.allocationManager(),
            "crossChainRegistry allocationManager mismatch"
        );
        assertTrue(
            crossChainRegistry.keyRegistrar() == Env.proxy.keyRegistrar(), "crossChainRegistry keyRegistrar mismatch"
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

    function _validateProxiesInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// CrossChainRegistry
        CrossChainRegistry crossChainRegistry = Env.proxy.crossChainRegistry();
        vm.expectRevert(errInit);
        crossChainRegistry.initialize(address(0), 1 days, 0);

        // ReleaseManager and KeyRegistrar don't have initialize functions
    }

    function _assertTrue(bool b, string memory err) private pure {
        assertTrue(b, err);
    }

    function _assertFalse(bool b, string memory err) private pure {
        assertFalse(b, err);
    }

    /// @dev Check if the version is already deployed
    function _isAlreadyDeployedSourceChain() internal view returns (bool) {
        if (Env._strEq(Env.envVersion(), "1.8.0")) {
            return true;
        }
        return false;
    }
}
