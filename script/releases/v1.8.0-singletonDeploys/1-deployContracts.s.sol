// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

/**
 * Purpose: use an EOA to deploy just the `KeyRegistrar` and `ReleaseManager` contracts.
 */
contract DeploySourceChain is EOADeployer {
    using Env for *;

    /// forgefmt: disable-next-item
    function _runAsEOA() internal override {
        // If we ARE on a source or destination chain, we don't need to deploy any contracts
        if (Env.isSourceChain() || Env.isDestinationChain()) {
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
        // If we ARE on a source or destination chain, we don't need to deploy any contracts
        if (Env.isSourceChain() || Env.isDestinationChain()) {   
            return;
        }

        // Set the mode to EOA so we can deploy the contracts
        super.runAsEOA();

        // Note: there are no initialize functions for KeyRegistrar or ReleaseManager, hence no need to validate them
        _validateStorage();
        _validateProxyAdmins();
        _validateImplConstructors();
        _validateProxyConstructors();
    }

    /// @dev Validate that storage variables are set correctly
    function _validateStorage() internal view {
        // Validate KeyRegistrar
        KeyRegistrar keyRegistrar = Env.proxy.keyRegistrar();
        assertTrue(address(keyRegistrar) != address(0), "keyRegistrar not deployed");

        // Validate ReleaseManager
        ReleaseManager releaseManager = Env.proxy.releaseManager();
        assertTrue(address(releaseManager) != address(0), "releaseManager not deployed");
    }

    /// @dev Ensure each deployed TUP/beacon is owned by the proxyAdmin/executorMultisig
    function _validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();

        assertTrue(Env._getProxyAdmin(address(Env.proxy.keyRegistrar())) == pa, "keyRegistrar proxyAdmin incorrect");
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
            /// ReleaseManager
            ReleaseManager releaseManager = Env.impl.releaseManager();
            assertTrue(
                releaseManager.permissionController() == Env.proxy.permissionController(),
                "rm.permissionController invalid"
            );
            assertEq(releaseManager.version(), Env.deployVersion(), "rm.version failed");
        }
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

        ReleaseManager releaseManager = Env.proxy.releaseManager();
        assertEq(releaseManager.version(), Env.deployVersion(), "releaseManager version mismatch");
        assertTrue(
            releaseManager.permissionController() == Env.proxy.permissionController(),
            "releaseManager permissionController mismatch"
        );
    }
    function _assertTrue(bool b, string memory err) private pure {
        assertTrue(b, err);
    }

    function _assertFalse(bool b, string memory err) private pure {
        assertFalse(b, err);
    }
}
