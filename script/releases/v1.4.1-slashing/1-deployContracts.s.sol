// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * Purpose: use an EOA to deploy all of the new contracts for this upgrade.
 */
contract Deploy is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        // We are upgrading 3 contracts: DelegationManager, EigenPodManager, and EigenPod
        deployImpl({
            name: type(DelegationManager).name,
            deployedTo: address(
                new DelegationManager({
                    _strategyManager: Env.proxy.strategyManager(),
                    _eigenPodManager: Env.proxy.eigenPodManager(),
                    _allocationManager: Env.proxy.allocationManager(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _permissionController: Env.proxy.permissionController(),
                    _MIN_WITHDRAWAL_DELAY: Env.MIN_WITHDRAWAL_DELAY(),
                    _version: Env.deployVersion()
                })
            )
        });

        deployImpl({
            name: type(EigenPodManager).name,
            deployedTo: address(
                new EigenPodManager({
                    _ethPOS: Env.ethPOS(),
                    _eigenPodBeacon: Env.beacon.eigenPod(),
                    _delegationManager: Env.proxy.delegationManager(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _version: Env.deployVersion()
                })
            )
        });

        deployImpl({
            name: type(EigenPod).name,
            deployedTo: address(
                new EigenPod({
                    _ethPOS: Env.ethPOS(),
                    _eigenPodManager: Env.proxy.eigenPodManager(),
                    _GENESIS_TIME: Env.EIGENPOD_GENESIS_TIME(),
                    _version: Env.deployVersion()
                })
            )
        });

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        _runAsEOA();

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
        /// core/

        function (bool, string memory) internal pure assertion = areMatching ? _assertTrue : _assertFalse;

        assertion(
            _getProxyImpl(address(Env.proxy.delegationManager())) == address(Env.impl.delegationManager()),
            "delegationManager impl failed"
        );

        assertion(Env.beacon.eigenPod().implementation() == address(Env.impl.eigenPod()), "eigenPod impl failed");

        assertion(
            _getProxyImpl(address(Env.proxy.eigenPodManager())) == address(Env.impl.eigenPodManager()),
            "eigenPodManager impl failed"
        );
    }

    /// @dev Ensure each deployed TUP/beacon is owned by the proxyAdmin/executorMultisig
    function _validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();

        assertTrue(
            _getProxyAdmin(address(Env.proxy.delegationManager())) == pa, "delegationManager proxyAdmin incorrect"
        );

        assertTrue(Env.beacon.eigenPod().owner() == Env.executorMultisig(), "eigenPod beacon owner incorrect");

        assertTrue(_getProxyAdmin(address(Env.proxy.eigenPodManager())) == pa, "eigenPodManager proxyAdmin incorrect");
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        {
            DelegationManager delegation = Env.impl.delegationManager();
            assertTrue(delegation.strategyManager() == Env.proxy.strategyManager(), "dm.sm invalid");
            assertTrue(delegation.eigenPodManager() == Env.proxy.eigenPodManager(), "dm.epm invalid");
            assertTrue(delegation.allocationManager() == Env.proxy.allocationManager(), "dm.alm invalid");
            assertTrue(delegation.pauserRegistry() == Env.impl.pauserRegistry(), "dm.pR invalid");
            assertTrue(delegation.permissionController() == Env.proxy.permissionController(), "dm.pc invalid");
            assertTrue(
                delegation.minWithdrawalDelayBlocks() == Env.MIN_WITHDRAWAL_DELAY(), "dm.withdrawalDelay invalid"
            );
        }

        {
            /// pods/
            EigenPod eigenPod = Env.impl.eigenPod();
            assertTrue(eigenPod.ethPOS() == Env.ethPOS(), "ep.ethPOS invalid");
            assertTrue(eigenPod.eigenPodManager() == Env.proxy.eigenPodManager(), "ep.epm invalid");
            assertTrue(eigenPod.GENESIS_TIME() == Env.EIGENPOD_GENESIS_TIME(), "ep.genesis invalid");

            EigenPodManager eigenPodManager = Env.impl.eigenPodManager();
            assertTrue(eigenPodManager.ethPOS() == Env.ethPOS(), "epm.ethPOS invalid");
            assertTrue(eigenPodManager.eigenPodBeacon() == Env.beacon.eigenPod(), "epm.epBeacon invalid");
            assertTrue(eigenPodManager.delegationManager() == Env.proxy.delegationManager(), "epm.dm invalid");
            assertTrue(eigenPodManager.pauserRegistry() == Env.impl.pauserRegistry(), "epm.pR invalid");
        }
    }

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        DelegationManager delegation = Env.impl.delegationManager();
        vm.expectRevert(errInit);
        delegation.initialize(address(0), 0);

        /// pods/
        EigenPod eigenPod = Env.impl.eigenPod();
        vm.expectRevert(errInit);
        eigenPod.initialize(address(0));

        EigenPodManager eigenPodManager = Env.impl.eigenPodManager();
        vm.expectRevert(errInit);
        eigenPodManager.initialize(address(0), 0);
    }

    function _validateVersion() internal view {
        // On future upgrades, just tick the major/minor/patch to validate
        string memory expected = Env.deployVersion();

        assertEq(Env.impl.delegationManager().version(), expected, "delegationManager version mismatch");
        assertEq(Env.impl.eigenPod().version(), expected, "eigenPod version mismatch");
        assertEq(Env.impl.eigenPodManager().version(), expected, "eigenPodManager version mismatch");
    }

    /// @dev Query and return `proxyAdmin.getProxyImplementation(proxy)`
    function _getProxyImpl(
        address proxy
    ) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyImplementation(ITransparentUpgradeableProxy(proxy));
    }

    /// @dev Query and return `proxyAdmin.getProxyAdmin(proxy)`
    function _getProxyAdmin(
        address proxy
    ) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyAdmin(ITransparentUpgradeableProxy(proxy));
    }

    function _assertTrue(bool b, string memory err) private pure {
        assertTrue(b, err);
    }

    function _assertFalse(bool b, string memory err) private pure {
        assertFalse(b, err);
    }
}
