// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * Purpose: use an EOA to deploy all of the new source chain contracts for this upgrade.
 */
contract DeploySourceChain is EOADeployer {
    using Env for *;

    /// forgefmt: disable-next-item
    function _runAsEOA() internal override {
        // If we're not on a source chain, we don't need to deploy any contracts
        if (!Env.isSourceChain()) {
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

        // Deploy ReleaseManager implementation
        deployImpl({
            name: type(ReleaseManager).name,
            deployedTo: address(
                new ReleaseManager({_permissionController: Env.proxy.permissionController(), _version: Env.deployVersion()})
            )
        });

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        if (!Env.isSourceChain()) {
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

        // KeyRegistrar
        assertion(
            Env._getProxyImpl(address(Env.proxy.keyRegistrar())) == address(Env.impl.keyRegistrar()),
            "keyRegistrar impl failed"
        );

        // CrossChainRegistry
        assertion(
            Env._getProxyImpl(address(Env.proxy.crossChainRegistry())) == address(Env.impl.crossChainRegistry()),
            "crossChainRegistry impl failed"
        );

        // ReleaseManager
        assertion(
            Env._getProxyImpl(address(Env.proxy.releaseManager())) == address(Env.impl.releaseManager()),
            "releaseManager impl failed"
        );
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

        // KeyRegistrar and ReleaseManager don't have initialize functions
    }

    function _validateVersion() internal view {
        // On future upgrades, just tick the major/minor/patch to validate
        string memory expected = Env.deployVersion();

        assertEq(Env.impl.keyRegistrar().version(), expected, "keyRegistrar version mismatch");
        assertEq(Env.impl.crossChainRegistry().version(), expected, "crossChainRegistry version mismatch");
        assertEq(Env.impl.releaseManager().version(), expected, "releaseManager version mismatch");
    }

    function _assertTrue(bool b, string memory err) internal pure {
        assertTrue(b, err);
    }

    function _assertFalse(bool b, string memory err) internal pure {
        assertFalse(b, err);
    }
}
