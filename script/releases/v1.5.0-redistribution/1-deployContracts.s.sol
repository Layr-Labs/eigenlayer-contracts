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

    /// forgefmt: disable-next-item
    function _runAsEOA() internal override {
        vm.startBroadcast();

        // We are upgrading 2 contracts: AllocationManager and the StrategyManager
        deployImpl({
            name: type(AllocationManager).name,
            deployedTo: address(
                new AllocationManager({
                    _delegation: Env.proxy.delegationManager(),
                    _eigenStrategy: Env.proxy.eigenStrategy(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _permissionController: Env.proxy.permissionController(),
                    _DEALLOCATION_DELAY: 1 days, // TODO
                    _ALLOCATION_CONFIGURATION_DELAY: 1 days, // TODO
                    _version: Env.deployVersion()
                })
            )
        });

        deployImpl({
            name: type(StrategyManager).name,
            deployedTo: address(
                new StrategyManager({
                   _allocationManager: Env.proxy.allocationManager(),
                   _delegation: Env.proxy.delegationManager(),
                   _pauserRegistry: Env.impl.pauserRegistry(),
                   _version: Env.deployVersion()
                })
            )
        });

        // TODO: Blacklist EIGEN from StrategyFactory.

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        _runAsEOA();

        _validateNewImplAddresses({areMatching: false});
        _validateProxyAdmins();
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
        function (bool, string memory) internal pure assertion = areMatching ? _assertTrue : _assertFalse;

        assertion(
            _getProxyImpl(address(Env.proxy.allocationManager())) == address(Env.impl.allocationManager()),
            "allocationManager impl failed"
        );

        assertion(
            _getProxyImpl(address(Env.proxy.strategyManager())) == address(Env.impl.strategyManager()),
            "strategyManager impl failed"
        );
    }

    /// @dev Ensure each deployed TUP/beacon is owned by the proxyAdmin/executorMultisig
    function _validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();

        assertTrue(
            _getProxyAdmin(address(Env.proxy.allocationManager())) == pa, "allocationManager proxyAdmin incorrect"
        );

        assertTrue(_getProxyAdmin(address(Env.proxy.strategyManager())) == pa, "strategyManager proxyAdmin incorrect");
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        {
            /// AllocationManager
            AllocationManager allocationManager = Env.impl.allocationManager();
            assertTrue(
                allocationManager.delegation() == Env.proxy.delegationManager(), "allocationManager.delegation invalid"
            );
            assertTrue(
                allocationManager.eigenStrategy() == Env.proxy.eigenStrategy(),
                "allocationManager.eigenStrategy invalid"
            );
            assertTrue(
                allocationManager.pauserRegistry() == Env.impl.pauserRegistry(),
                "allocationManager.pauserRegistry invalid"
            );
            assertTrue(
                allocationManager.permissionController() == Env.proxy.permissionController(),
                "allocationManager.permissionController invalid"
            );
            assertTrue(
                allocationManager.DEALLOCATION_DELAY() == 1 days, // TODO
                "allocationManager.DEALLOCATION_DELAY invalid"
            );
            assertTrue(
                allocationManager.ALLOCATION_CONFIGURATION_DELAY() == 1 days, // TODO
                "allocationManager.ALLOCATION_CONFIGURATION_DELAY invalid"
            );
            assertTrue(_strEq(allocationManager.version(), Env.deployVersion()), "allocationManager.version failed");
        }

        {
            /// StrategyManager
            StrategyManager strategyManager = Env.impl.strategyManager();
            assertTrue(
                strategyManager.allocationManager() == Env.proxy.allocationManager(),
                "strategyManager.allocationManager invalid"
            );
            assertTrue(
                strategyManager.delegation() == Env.proxy.delegationManager(), "strategyManager.delegation invalid"
            );
            assertTrue(
                strategyManager.pauserRegistry() == Env.impl.pauserRegistry(), "strategyManager.pauserRegistry invalid"
            );
            assertTrue(_strEq(strategyManager.version(), Env.deployVersion()), "strategyManager.version failed");
        }
    }

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// AllocationManager
        AllocationManager allocationManager = Env.impl.allocationManager();
        vm.expectRevert(errInit);
        allocationManager.initialize(0);

        /// StrategyManager
        StrategyManager strategyManager = Env.impl.strategyManager();
        vm.expectRevert(errInit);
        strategyManager.initialize(address(0), address(0), 0);
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

    function _strEq(string memory a, string memory b) private pure returns (bool) {
        return keccak256(bytes(a)) == keccak256(bytes(b));
    }
}
