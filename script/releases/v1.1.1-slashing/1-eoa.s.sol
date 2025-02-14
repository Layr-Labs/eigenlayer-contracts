// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

contract Deploy is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        // Deploy Allocation Manager
        deployImpl({
            name: type(AllocationManager).name,
            deployedTo: address(
                new AllocationManager({
                    _delegation: Env.proxy.delegationManager(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _permissionController: Env.proxy.permissionController(),
                    _DEALLOCATION_DELAY: Env.MIN_WITHDRAWAL_DELAY(),
                    _ALLOCATION_CONFIGURATION_DELAY: Env.ALLOCATION_CONFIGURATION_DELAY()
                })
            )
        });

        vm.stopBroadcast();
    }

    function testDeploy() public virtual {
        _runAsEOA();
        _validateNewImplAddresses(false);
        _validateImplConstructors();
        _validateImplsInitialized();
    }

    /// @dev Validate that the `Env.impl` addresses are updated to be distinct from what the proxy
    /// admin reports as the current implementation address.
    ///
    /// Note: The upgrade script can call this with `areMatching == true` to check that these impl
    /// addresses _are_ matches.
    function _validateNewImplAddresses(
        bool areMatching
    ) internal view {
        function(address, address, string memory)
            internal
            pure assertion = areMatching ? _assertMatch : _assertNotMatch;

        assertion(
            _getProxyImpl(address(Env.proxy.allocationManager())),
            address(Env.impl.allocationManager()),
            "allocationManager impl failed"
        );
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        AllocationManager allocationManager = Env.impl.allocationManager();
        assertTrue(allocationManager.delegation() == Env.proxy.delegationManager(), "dm invalid");
        assertTrue(allocationManager.pauserRegistry() == Env.impl.pauserRegistry(), "pR invalid");
        assertTrue(allocationManager.permissionController() == Env.proxy.permissionController(), "pc invalid");
        assertTrue(allocationManager.DEALLOCATION_DELAY() == Env.MIN_WITHDRAWAL_DELAY(), "deallocationDelay invalid");
        assertTrue(
            allocationManager.ALLOCATION_CONFIGURATION_DELAY() == Env.ALLOCATION_CONFIGURATION_DELAY(),
            "allocationConfigurationDelay invalid"
        );
    }

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        AllocationManager allocationManager = Env.impl.allocationManager();
        vm.expectRevert(errInit);
        allocationManager.initialize(address(0), 0);
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

    function _assertMatch(address a, address b, string memory err) private pure {
        assertEq(a, b, err);
    }

    function _assertNotMatch(address a, address b, string memory err) private pure {
        assertNotEq(a, b, err);
    }
}
