// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../Env.sol";
import {QueueUpgrade} from "./2-queueUpgrade.s.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract Execute is QueueUpgrade {
    using Env for *;

    function _runAsMultisig() internal override prank(Env.protocolCouncilMultisig()) {
        bytes memory calldata_to_executor = _getCalldataToExecutor();

        TimelockController timelock = Env.timelockController();
        timelock.execute({
            target: Env.executorMultisig(),
            value: 0,
            payload: calldata_to_executor,
            predecessor: 0,
            salt: 0
        });
    }

    function testScript() public virtual override {
        runAsEOA();

        TimelockController timelock = Env.timelockController();
        bytes memory calldata_to_executor = _getCalldataToExecutor();
        bytes32 txHash = timelock.hashOperation({
            target: Env.executorMultisig(),
            value: 0,
            data: calldata_to_executor,
            predecessor: 0,
            salt: 0
        });

        assertFalse(timelock.isOperationPending(txHash), "Transaction should NOT be queued.");

        // 1- run queueing logic
        QueueUpgrade._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
        assertFalse(timelock.isOperationReady(txHash), "Transaction should NOT be ready for execution.");
        assertFalse(timelock.isOperationDone(txHash), "Transaction should NOT be complete.");

        // 2- warp past delay
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA
        assertEq(timelock.isOperationReady(txHash), true, "Transaction should be executable.");

        // 3- execute
        execute();

        assertTrue(timelock.isOperationDone(txHash), "Transaction should be complete.");

        _validateNewImplAddresses({areMatching: true});
        _validateProxyAdmins();
        _validateProxyConstructors();
        _validateProxiesInitialized();
    }

    /// @dev Mirrors the checks done in 1-deployContracts, but now we check each contract's
    /// proxy, as the upgrade should mean that each proxy can see these methods/immutables
    function _validateProxyConstructors() internal view {
        SlashEscrowFactory slashEscrowFactory = Env.proxy.slashEscrowFactory();
        assertTrue(slashEscrowFactory.allocationManager() == Env.proxy.allocationManager(), "sef.alm invalid");
        assertTrue(slashEscrowFactory.strategyManager() == Env.proxy.strategyManager(), "sef.sm invalid");
        assertTrue(slashEscrowFactory.pauserRegistry() == Env.impl.pauserRegistry(), "sef.pR invalid");
        assertTrue(slashEscrowFactory.slashEscrowImplementation() == Env.impl.slashEscrow(), "sef.se invalid");
        // Check slashEscrowFactory local vars again for sanity
        assertTrue(slashEscrowFactory.owner() == Env.executorMultisig(), "sef.owner invalid");
        assertTrue(slashEscrowFactory.paused() == 0, "sef.paused invalid");
        assertTrue(slashEscrowFactory.getGlobalEscrowDelay() == Env.SLASH_ESCROW_DELAY(), "sef.globalDelay invalid");

        AllocationManager allocationManager = Env.proxy.allocationManager();
        assertTrue(allocationManager.delegation() == Env.proxy.delegationManager(), "alm.dm invalid");
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
    function _validateProxiesInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        // SlashEscrow is immutable and uninitializable
        SlashEscrowFactory slashEscrowFactory = Env.proxy.slashEscrowFactory();
        vm.expectRevert(errInit);
        slashEscrowFactory.initialize(address(0), 0, 0);
        assertTrue(slashEscrowFactory.owner() == Env.executorMultisig(), "sef.owner invalid");
        assertTrue(slashEscrowFactory.paused() == 0, "sef.paused invalid");

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
        assertTrue(eigenStrategy.EIGEN() == Env.proxy.eigen(), "eigenStrat.EIGEN invalid");
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
}
