// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {QueueUpgrade} from "./2-queueUpgrade.s.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "zeus-templates/utils/Encode.sol";

/// @notice Executes the upgrade for redistribution v1.5.0
contract Execute is QueueUpgrade {
    using Env for *;
    using Encode for *;

    function _runAsMultisig() internal override prank(Env.protocolCouncilMultisig()) {
        bytes memory calldata_to_executor_v1_5_queue = _getCalldataToExecutor_v1_5_queue();
        TimelockController timelock = Env.timelockController();

        if (_isMainnet()) {
            timelock.execute({
                target: Env.executorMultisig(),
                value: 0,
                payload: calldata_to_executor_v1_5_queue,
                predecessor: 0,
                salt: 0
            });
        }
    }

    /// @dev Get the calldata to be sent from the timelock to the executor
    /// Copied over from v1.5.0 release script
    function _getCalldataToExecutor_v1_5_queue() internal returns (bytes memory) {
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls().append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.delegationManager()),
                impl: address(Env.impl.delegationManager())
            })
        }).append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.allocationManager()),
                impl: address(Env.impl.allocationManager())
            })
        }).append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.strategyManager()),
                impl: address(Env.impl.strategyManager())
            })
        }).append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.eigenPodManager()),
                impl: address(Env.impl.eigenPodManager())
            })
        }).append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.eigenStrategy()),
                impl: address(Env.impl.eigenStrategy())
            })
        }).append({
            to: address(Env.beacon.strategyBase()),
            data: Encode.upgradeableBeacon.upgradeTo({newImpl: address(Env.impl.strategyBase())})
        });

        // Add call to upgrade each pre-longtail strategy instance
        uint256 count = Env.instance.strategyBaseTVLLimits_Count();
        for (uint256 i = 0; i < count; i++) {
            address proxyInstance = address(Env.instance.strategyBaseTVLLimits(i));

            executorCalls.append({
                to: Env.proxyAdmin(),
                data: Encode.proxyAdmin.upgrade({proxy: proxyInstance, impl: address(Env.impl.strategyBaseTVLLimits())})
            });
        }

        return Encode.gnosisSafe.execTransaction({
            from: address(Env.timelockController()),
            to: Env.multiSendCallOnly(),
            op: Encode.Operation.DelegateCall,
            data: Encode.multiSend(executorCalls)
        });
    }

    function testScript() public virtual override {
        if (!_isMainnet()) {
            return;
        }

        TimelockController timelock = Env.timelockController();
        bytes memory calldata_to_executor_v1_5_queue = _getCalldataToExecutor_v1_5_queue();
        bytes32 txHash_v1_5 = timelock.hashOperation({
            target: Env.executorMultisig(),
            value: 0,
            data: calldata_to_executor_v1_5_queue,
            predecessor: 0,
            salt: 0
        });
        // 1 - Deploy. The contracts have been deployed in the redistro upgrade script

        /// 2 - Queue. The contracts have been deployed in the redistribution upgrade script.
        /// At the time of writing, the operation IS ready
        assertEq(timelock.isOperationReady(txHash_v1_5), true, "v1.5 txn should be executable.");

        // 3 - execute
        execute();
        assertTrue(timelock.isOperationDone(txHash_v1_5), "v1.5 txn should be complete.");

        // 4. Validate
        _validateNewImplAddresses_v1_5({areMatching: true});
        _validateProxyAdmins_v1_5();
        _validateProxyConstructors_v1_5();
        _validateProxiesInitialized_v1_5();
    }

    /// @dev Validate that the `Env.impl` addresses are updated to be distinct from what the proxy
    /// admin reports as the current implementation address.
    ///
    /// Note: The upgrade script can call this with `areMatching == true` to check that these impl
    /// addresses _are_ matches.
    /// Copied over from v1.5.0 release script
    function _validateNewImplAddresses_v1_5(
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
    /// Copied over from v1.5.0 release script
    function _validateProxyAdmins_v1_5() internal view {
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

    /// @dev Mirrors the checks done in 1-deployContracts, but now we check each contract's
    /// proxy, as the upgrade should mean that each proxy can see these methods/immutables
    /// Copied over from v1.5.0 release script
    function _validateProxyConstructors_v1_5() internal view {
        AllocationManager allocationManager = Env.proxy.allocationManager();
        assertTrue(allocationManager.delegation() == Env.proxy.delegationManager(), "alm.dm invalid");
        assertTrue(allocationManager.eigenStrategy() == Env.proxy.eigenStrategy(), "alm.es invalid");
        assertTrue(allocationManager.pauserRegistry() == Env.impl.pauserRegistry(), "alm.pR invalid");
        assertTrue(allocationManager.permissionController() == Env.proxy.permissionController(), "alm.pc invalid");
        assertTrue(allocationManager.DEALLOCATION_DELAY() == Env.MIN_WITHDRAWAL_DELAY(), "alm.deallocDelay invalid");
        assertTrue(
            allocationManager.ALLOCATION_CONFIGURATION_DELAY() == Env.ALLOCATION_CONFIGURATION_DELAY(),
            "alm.configDelay invalid"
        );

        DelegationManager delegation = Env.proxy.delegationManager();
        assertTrue(delegation.strategyManager() == Env.proxy.strategyManager(), "dm.sm invalid");
        assertTrue(delegation.eigenPodManager() == Env.proxy.eigenPodManager(), "dm.epm invalid");
        assertTrue(delegation.allocationManager() == Env.proxy.allocationManager(), "dm.alm invalid");
        assertTrue(delegation.pauserRegistry() == Env.impl.pauserRegistry(), "dm.pR invalid");
        assertTrue(delegation.permissionController() == Env.proxy.permissionController(), "dm.pc invalid");
        assertTrue(delegation.minWithdrawalDelayBlocks() == Env.MIN_WITHDRAWAL_DELAY(), "dm.withdrawalDelay invalid");

        StrategyManager strategyManager = Env.proxy.strategyManager();
        assertTrue(strategyManager.delegation() == Env.proxy.delegationManager(), "sm.dm invalid");
        assertTrue(strategyManager.pauserRegistry() == Env.impl.pauserRegistry(), "sm.pR invalid");

        EigenPodManager eigenPodManager = Env.proxy.eigenPodManager();
        assertTrue(eigenPodManager.ethPOS() == Env.ethPOS(), "epm.ethPOS invalid");
        assertTrue(eigenPodManager.eigenPodBeacon() == Env.beacon.eigenPod(), "epm.epBeacon invalid");
        assertTrue(eigenPodManager.delegationManager() == Env.proxy.delegationManager(), "epm.dm invalid");
        assertTrue(eigenPodManager.pauserRegistry() == Env.impl.pauserRegistry(), "epm.pR invalid");

        /// strategies/
        EigenStrategy eigenStrategy = Env.proxy.eigenStrategy();
        assertTrue(eigenStrategy.strategyManager() == Env.proxy.strategyManager(), "eigStrat.sm invalid");
        assertTrue(eigenStrategy.pauserRegistry() == Env.impl.pauserRegistry(), "eigStrat.pR invalid");

        UpgradeableBeacon strategyBeacon = Env.beacon.strategyBase();
        assertTrue(strategyBeacon.implementation() == address(Env.impl.strategyBase()), "strategyBeacon.impl invalid");

        uint256 count = Env.instance.strategyBaseTVLLimits_Count();
        for (uint256 i = 0; i < count; i++) {
            StrategyBaseTVLLimits strategy = Env.instance.strategyBaseTVLLimits(i);

            assertTrue(strategy.strategyManager() == Env.proxy.strategyManager(), "sFact.sm invalid");
            assertTrue(strategy.pauserRegistry() == Env.impl.pauserRegistry(), "sFact.pR invalid");
        }
    }

    /// @dev Call initialize on all proxies to ensure they are initialized
    /// Additionally, validate initialization variables
    /// Copied over from v1.5.0 release script
    function _validateProxiesInitialized_v1_5() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        AllocationManager allocationManager = Env.proxy.allocationManager();
        vm.expectRevert(errInit);
        allocationManager.initialize(0);
        assertTrue(allocationManager.paused() == 0, "alm.paused invalid");

        DelegationManager delegation = Env.proxy.delegationManager();
        vm.expectRevert(errInit);
        delegation.initialize(0);
        assertTrue(delegation.paused() == 0, "dm.paused invalid");

        StrategyManager strategyManager = Env.proxy.strategyManager();
        vm.expectRevert(errInit);
        strategyManager.initialize(address(0), address(0), 0);
        assertTrue(strategyManager.owner() == Env.executorMultisig(), "sm.owner invalid");
        assertTrue(strategyManager.paused() == 0, "sm.paused invalid");
        assertTrue(
            strategyManager.strategyWhitelister() == address(Env.proxy.strategyFactory()), "sm.whitelister invalid"
        );

        EigenPodManager eigenPodManager = Env.proxy.eigenPodManager();
        vm.expectRevert(errInit);
        eigenPodManager.initialize(address(0), 0);
        assertTrue(eigenPodManager.owner() == Env.executorMultisig(), "epm.owner invalid");
        // For sepolia, eigenpodmanager is paused
        if (block.chainid != 11_155_111) {
            assertTrue(eigenPodManager.paused() == 0, "epm.paused invalid");
        } else {
            assertTrue(eigenPodManager.paused() == 487, "epm.paused invalid");
        }

        EigenStrategy eigenStrategy = Env.proxy.eigenStrategy();
        vm.expectRevert(errInit);
        eigenStrategy.initialize(IEigen(address(0)), IBackingEigen(address(0)));
        assertTrue(eigenStrategy.paused() == 0, "eigenStrat.paused invalid");
        assertTrue(address(eigenStrategy.EIGEN()) == address(Env.proxy.eigen()), "eigenStrat.EIGEN invalid");
        assertTrue(eigenStrategy.underlyingToken() == Env.proxy.beigen(), "eigenStrat.underlying invalid");

        // StrategyBase proxies are initialized when deployed by factory

        uint256 count = Env.instance.strategyBaseTVLLimits_Count();
        for (uint256 i = 0; i < count; i++) {
            StrategyBaseTVLLimits strategy = Env.instance.strategyBaseTVLLimits(i);

            emit log_named_address("strat", address(strategy));

            vm.expectRevert(errInit);
            strategy.initialize(0, 0, IERC20(address(0)));
            assertTrue(strategy.maxPerDeposit() == type(uint256).max, "stratTVLLim.maxPerDeposit invalid");
            assertTrue(strategy.maxTotalDeposits() == type(uint256).max, "stratTVLLim.maxPerDeposit invalid");
        }
    }

    function _isMainnet() internal view returns (bool) {
        return _strEq(Env.env(), "mainnet");
    }
}
