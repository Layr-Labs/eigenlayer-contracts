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
        {
            UpgradeableBeacon eigenPodBeacon = Env.beacon.eigenPod();
            assertTrue(eigenPodBeacon.implementation() == address(Env.impl.eigenPod()), "eigenPodBeacon.impl invalid");

            /// EigenPod
            EigenPod eigenPod = Env.impl.eigenPod();
            assertTrue(eigenPod.ethPOS() == Env.ethPOS(), "ep.ethPOS invalid");
            assertTrue(eigenPod.eigenPodManager() == Env.proxy.eigenPodManager(), "ep.epm invalid");
            assertEq(eigenPod.version(), Env.deployVersion(), "ep.version failed");
        }

        {
            /// Eigen
            Eigen eigen = Eigen(address(Env.proxy.eigen()));
            assertTrue(address(eigen.bEIGEN()) == address(Env.proxy.beigen()), "eigen.beigen invalid");
        }
    }

    /// @dev Call initialize on all proxies to ensure they are initialized
    /// Additionally, validate initialization variables
    function _validateProxiesInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// Eigen
        Eigen eigen = Eigen(address(Env.proxy.eigen()));
        vm.expectRevert(errInit);
        eigen.initialize(address(0), new address[](0), new uint256[](0), new uint256[](0));
    }
}
