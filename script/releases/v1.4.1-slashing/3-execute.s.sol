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
        DelegationManager delegation = Env.proxy.delegationManager();
        assertTrue(delegation.strategyManager() == Env.proxy.strategyManager(), "dm.sm invalid");
        assertTrue(delegation.eigenPodManager() == Env.proxy.eigenPodManager(), "dm.epm invalid");
        assertTrue(delegation.allocationManager() == Env.proxy.allocationManager(), "dm.alm invalid");
        assertTrue(delegation.pauserRegistry() == Env.impl.pauserRegistry(), "dm.pR invalid");
        assertTrue(delegation.permissionController() == Env.proxy.permissionController(), "dm.pc invalid");
        assertTrue(delegation.minWithdrawalDelayBlocks() == Env.MIN_WITHDRAWAL_DELAY(), "dm.withdrawalDelay invalid");

        UpgradeableBeacon eigenPodBeacon = Env.beacon.eigenPod();
        assertTrue(eigenPodBeacon.implementation() == address(Env.impl.eigenPod()), "eigenPodBeacon.impl invalid");

        EigenPodManager eigenPodManager = Env.proxy.eigenPodManager();
        assertTrue(eigenPodManager.ethPOS() == Env.ethPOS(), "epm.ethPOS invalid");
        assertTrue(eigenPodManager.eigenPodBeacon() == Env.beacon.eigenPod(), "epm.epBeacon invalid");
        assertTrue(eigenPodManager.delegationManager() == Env.proxy.delegationManager(), "epm.dm invalid");
        assertTrue(eigenPodManager.pauserRegistry() == Env.impl.pauserRegistry(), "epm.pR invalid");
    }

    /// @dev Call initialize on all proxies to ensure they are initialized
    /// Additionally, validate initialization variables
    function _validateProxiesInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        DelegationManager delegation = Env.proxy.delegationManager();
        vm.expectRevert(errInit);
        delegation.initialize(address(0), 0);
        assertTrue(delegation.owner() == Env.executorMultisig(), "dm.owner invalid");
        assertTrue(delegation.paused() == 0, "dm.paused invalid");

        EigenPodManager eigenPodManager = Env.proxy.eigenPodManager();
        vm.expectRevert(errInit);
        eigenPodManager.initialize(address(0), 0);
        assertTrue(eigenPodManager.owner() == Env.executorMultisig(), "epm.owner invalid");
        assertTrue(eigenPodManager.paused() == 0, "epm.paused invalid");
    }
}
