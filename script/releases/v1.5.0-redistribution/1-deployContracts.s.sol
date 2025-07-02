// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * Purpose: use an EOA to deploy all of the new contracts for this upgrade.
 * This upgrades:
 * - `AllocationManager`
 * - `DelegationManager`
 * - `StrategyManager`
 * - `EigenPodManager`
 * - Strategies (`EigenStrategy`, `StrategyBase`, `StrategyBaseTVLLimits`)
 */
contract Deploy is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        // Core contracts: AllocationManager, DelegationManager, StrategyManager, EigenPodManager, Strategies

        // AllocationManager
        deployImpl({
            name: type(AllocationManager).name,
            deployedTo: address(
                new AllocationManager({
                    _delegation: Env.proxy.delegationManager(),
                    _eigenStrategy: Env.proxy.eigenStrategy(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _permissionController: Env.proxy.permissionController(),
                    _DEALLOCATION_DELAY: Env.MIN_WITHDRAWAL_DELAY(),
                    _ALLOCATION_CONFIGURATION_DELAY: Env.ALLOCATION_CONFIGURATION_DELAY(),
                    _version: Env.deployVersion()
                })
            )
        });

        // DelegationManager
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

        // StrategyManager
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

        // EigenPodManager
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

        // Strategies: EigenStrategy, StrategyBase (for strategies deployed via factory), StrategyBaseTVLLimits
        deployImpl({
            name: type(EigenStrategy).name,
            deployedTo: address(
                new EigenStrategy({
                    _strategyManager: Env.proxy.strategyManager(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _version: Env.deployVersion()
                })
            )
        });

        // for strategies deployed via factory
        deployImpl({
            name: type(StrategyBase).name,
            deployedTo: address(
                new StrategyBase({
                    _strategyManager: Env.proxy.strategyManager(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _version: Env.deployVersion()
                })
            )
        });

        deployImpl({
            name: type(StrategyBaseTVLLimits).name,
            deployedTo: address(
                new StrategyBaseTVLLimits({
                    _strategyManager: Env.proxy.strategyManager(),
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

        // AllocationManager
        assertion(
            _getProxyImpl(address(Env.proxy.allocationManager())) == address(Env.impl.allocationManager()),
            "allocationManager impl failed"
        );

        // DelegationManager
        assertion(
            _getProxyImpl(address(Env.proxy.delegationManager())) == address(Env.impl.delegationManager()),
            "delegationManager impl failed"
        );

        // StrategyManager
        assertion(
            _getProxyImpl(address(Env.proxy.strategyManager())) == address(Env.impl.strategyManager()),
            "strategyManager impl failed"
        );

        // EigenPodManager
        assertion(
            _getProxyImpl(address(Env.proxy.eigenPodManager())) == address(Env.impl.eigenPodManager()),
            "eigenPodManager impl failed"
        );

        // Strategies
        assertion(
            _getProxyImpl(address(Env.proxy.eigenStrategy())) == address(Env.impl.eigenStrategy()),
            "eigenStrategy impl failed"
        );

        assertion(
            Env.beacon.strategyBase().implementation() == address(Env.impl.strategyBase()), "strategyBase impl failed"
        );

        uint256 count = Env.instance.strategyBaseTVLLimits_Count();
        for (uint256 i = 0; i < count; i++) {
            assertion(
                _getProxyImpl(address(Env.instance.strategyBaseTVLLimits(i)))
                    == address(Env.impl.strategyBaseTVLLimits()),
                "strategyBaseTVLLimits impl failed"
            );
        }
    }

    /// @dev Ensure each deployed TUP/beacon is owned by the proxyAdmin/executorMultisig
    function _validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();

        // AllocationManager
        assertTrue(
            _getProxyAdmin(address(Env.proxy.allocationManager())) == pa, "allocationManager proxyAdmin incorrect"
        );

        // DelegationManager
        assertTrue(
            _getProxyAdmin(address(Env.proxy.delegationManager())) == pa, "delegationManager proxyAdmin incorrect"
        );

        // StrategyManager
        assertTrue(_getProxyAdmin(address(Env.proxy.strategyManager())) == pa, "strategyManager proxyAdmin incorrect");

        // EigenPodManager
        assertTrue(_getProxyAdmin(address(Env.proxy.eigenPodManager())) == pa, "eigenPodManager proxyAdmin incorrect");

        // Strategies
        assertTrue(_getProxyAdmin(address(Env.proxy.eigenStrategy())) == pa, "eigenStrategy proxyAdmin incorrect");

        assertTrue(Env.beacon.strategyBase().owner() == Env.executorMultisig(), "strategyBase beacon owner incorrect");

        uint256 count = Env.instance.strategyBaseTVLLimits_Count();
        for (uint256 i = 0; i < count; i++) {
            assertTrue(
                _getProxyAdmin(address(Env.instance.strategyBaseTVLLimits(i))) == pa,
                "strategyBaseTVLLimits proxyAdmin incorrect"
            );
        }
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        // AllocationManager
        {
            AllocationManager allocationManager = Env.impl.allocationManager();
            assertTrue(allocationManager.delegation() == Env.proxy.delegationManager(), "alm.dm invalid");
            assertTrue(allocationManager.eigenStrategy() == Env.proxy.eigenStrategy(), "alm.es invalid");
            assertTrue(allocationManager.pauserRegistry() == Env.impl.pauserRegistry(), "alm.pR invalid");
            assertTrue(allocationManager.permissionController() == Env.proxy.permissionController(), "alm.pc invalid");
            assertTrue(allocationManager.DEALLOCATION_DELAY() == Env.MIN_WITHDRAWAL_DELAY(), "alm.deallocDelay invalid");
            assertTrue(
                allocationManager.ALLOCATION_CONFIGURATION_DELAY() == Env.ALLOCATION_CONFIGURATION_DELAY(),
                "alm.configDelay invalid"
            );
        }

        // DelegationManager
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

        // StrategyManager
        {
            StrategyManager strategyManager = Env.impl.strategyManager();
            assertTrue(strategyManager.delegation() == Env.proxy.delegationManager(), "sm.dm invalid");
            assertTrue(strategyManager.pauserRegistry() == Env.impl.pauserRegistry(), "sm.pR invalid");
        }

        // EigenPodManager
        {
            EigenPodManager eigenPodManager = Env.impl.eigenPodManager();
            assertTrue(eigenPodManager.ethPOS() == Env.ethPOS(), "epm.ethPOS invalid");
            assertTrue(eigenPodManager.eigenPodBeacon() == Env.beacon.eigenPod(), "epm.epBeacon invalid");
            assertTrue(eigenPodManager.delegationManager() == Env.proxy.delegationManager(), "epm.dm invalid");
            assertTrue(eigenPodManager.pauserRegistry() == Env.impl.pauserRegistry(), "epm.pR invalid");
        }

        // Strategies
        {
            EigenStrategy eigenStrategy = Env.impl.eigenStrategy();
            assertTrue(eigenStrategy.strategyManager() == Env.proxy.strategyManager(), "eigStrat.sm invalid");
            assertTrue(eigenStrategy.pauserRegistry() == Env.impl.pauserRegistry(), "eigStrat.pR invalid");

            StrategyBase strategyBase = Env.impl.strategyBase();
            assertTrue(strategyBase.strategyManager() == Env.proxy.strategyManager(), "stratBase.sm invalid");
            assertTrue(strategyBase.pauserRegistry() == Env.impl.pauserRegistry(), "stratBase.pR invalid");

            StrategyBaseTVLLimits strategyBaseTVLLimits = Env.impl.strategyBaseTVLLimits();
            assertTrue(
                strategyBaseTVLLimits.strategyManager() == Env.proxy.strategyManager(), "stratBaseTVL.sm invalid"
            );
            assertTrue(strategyBaseTVLLimits.pauserRegistry() == Env.impl.pauserRegistry(), "stratBaseTVL.pR invalid");
        }
    }

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        // AllocationManager
        AllocationManager allocationManager = Env.impl.allocationManager();
        vm.expectRevert(errInit);
        allocationManager.initialize(0);

        // DelegationManager
        DelegationManager delegation = Env.impl.delegationManager();
        vm.expectRevert(errInit);
        delegation.initialize(0);

        // StrategyManager
        StrategyManager strategyManager = Env.impl.strategyManager();
        vm.expectRevert(errInit);
        strategyManager.initialize(address(0), address(0), 0);

        // EigenPodManager
        EigenPodManager eigenPodManager = Env.impl.eigenPodManager();
        vm.expectRevert(errInit);
        eigenPodManager.initialize(address(0), 0);

        // Strategies
        EigenStrategy eigenStrategy = Env.impl.eigenStrategy();
        vm.expectRevert(errInit);
        eigenStrategy.initialize(IEigen(address(0)), IBackingEigen(address(0)));

        StrategyBase strategyBase = Env.impl.strategyBase();
        vm.expectRevert(errInit);
        strategyBase.initialize(IERC20(address(0)));

        StrategyBaseTVLLimits strategyBaseTVLLimits = Env.impl.strategyBaseTVLLimits();
        vm.expectRevert(errInit);
        strategyBaseTVLLimits.initialize(0, 0, IERC20(address(0)));
    }

    function _validateVersion() internal view {
        // On future upgrades, just tick the major/minor/patch to validate
        string memory expected = Env.deployVersion();

        assertEq(Env.impl.allocationManager().version(), expected, "allocationManager version mismatch");
        assertEq(Env.impl.delegationManager().version(), expected, "delegationManager version mismatch");
        assertEq(Env.impl.strategyManager().version(), expected, "strategyManager version mismatch");
        assertEq(Env.impl.eigenPodManager().version(), expected, "eigenPodManager version mismatch");
        assertEq(Env.impl.eigenStrategy().version(), expected, "eigenStrategy version mismatch");
        assertEq(Env.impl.strategyBase().version(), expected, "strategyBase version mismatch");
        assertEq(Env.impl.strategyBaseTVLLimits().version(), expected, "strategyBaseTVLLimits version mismatch");
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
