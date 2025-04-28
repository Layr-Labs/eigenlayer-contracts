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

        // We are upgrading 1 contract: AllocationManager
        deployImpl({
            name: type(AllocationManager).name,
            deployedTo: address(
                new AllocationManager({
                    _delegation: Env.proxy.delegationManager(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _permissionController: Env.proxy.permissionController(),
                    _DEALLOCATION_DELAY: Env.MIN_WITHDRAWAL_DELAY(),
                    _ALLOCATION_CONFIGURATION_DELAY: Env.ALLOCATION_CONFIGURATION_DELAY(),
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
            _getProxyImpl(address(Env.proxy.allocationManager())) == address(Env.impl.allocationManager()),
            "allocationManager impl failed"
        );
    }

    /// @dev Ensure each deployed TUP/beacon is owned by the proxyAdmin/executorMultisig
    function _validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();

        assertTrue(
            _getProxyAdmin(address(Env.proxy.allocationManager())) == pa, "allocationManager proxyAdmin incorrect"
        );
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        AllocationManager allocationManager = Env.impl.allocationManager();
        assertTrue(allocationManager.delegation() == Env.proxy.delegationManager(), "am.dm invalid");
        assertTrue(allocationManager.pauserRegistry() == Env.impl.pauserRegistry(), "am.pr invalid");
        assertTrue(allocationManager.permissionController() == Env.proxy.permissionController(), "am.pc invalid");
        assertTrue(allocationManager.DEALLOCATION_DELAY() == Env.MIN_WITHDRAWAL_DELAY(), "am.deallocDelay invalid");
        assertTrue(
            allocationManager.ALLOCATION_CONFIGURATION_DELAY() == Env.ALLOCATION_CONFIGURATION_DELAY(),
            "am.configDelay invalid"
        );
    }

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        AllocationManager allocationManager = Env.impl.allocationManager();
        vm.expectRevert(errInit);
        allocationManager.initialize(address(0), 0);
    }

    function _validateVersion() internal view {
        // On future upgrades, just tick the major/minor/patch to validate
        string memory expected = Env.deployVersion();

        assertEq(Env.impl.allocationManager().version(), expected, "allocationManager version mismatch");
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
