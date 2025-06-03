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
        _validateEigenPodETHPos();
    }

    /// @dev Mirrors the checks done in 1-deployContracts, but now we check each contract's
    /// proxy, as the upgrade should mean that each proxy can see these methods/immutables
    function _validateProxyConstructors() internal view {
        {
            EigenPodManager eigenPodManager = Env.proxy.eigenPodManager();
            assertTrue(eigenPodManager.ethPOS() == Env.ethPOS(), "epm.ethPOS invalid");
            // Validate hoodi specific ethPOS
            if (block.chainid == 560_048) {
                assertTrue(
                    address(eigenPodManager.ethPOS()) == 0x00000000219ab540356cBB839Cbe05303d7705Fa,
                    "hoodi ETH POS invalid"
                );
            }
            assertTrue(eigenPodManager.eigenPodBeacon() == Env.beacon.eigenPod(), "epm.epBeacon invalid");
            assertTrue(eigenPodManager.delegationManager() == Env.proxy.delegationManager(), "epm.dm invalid");
            assertTrue(eigenPodManager.pauserRegistry() == Env.impl.pauserRegistry(), "epm.pR invalid");
        }

        {
            UpgradeableBeacon eigenPodBeacon = Env.beacon.eigenPod();
            assertTrue(eigenPodBeacon.implementation() == address(Env.impl.eigenPod()), "eigenPodBeacon.impl invalid");

            /// EigenPod
            EigenPod eigenPod = Env.impl.eigenPod();
            assertTrue(eigenPod.ethPOS() == Env.ethPOS(), "ep.ethPOS invalid");
            assertTrue(eigenPod.eigenPodManager() == Env.proxy.eigenPodManager(), "ep.epm invalid");
            assertEq(eigenPod.version(), Env.deployVersion(), "ep.version failed");
        }
    }

    /// @dev Call initialize on all proxies to ensure they are initialized
    /// Additionally, validate initialization variables
    function _validateProxiesInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        EigenPodManager eigenPodManager = Env.proxy.eigenPodManager();
        vm.expectRevert(errInit);
        eigenPodManager.initialize(address(0), 0);
        assertTrue(eigenPodManager.owner() == Env.executorMultisig(), "epm.owner invalid");
        // EigenPods are paused on sepolia
        if (block.chainid == 11155111) {
            assertTrue(eigenPodManager.paused() == 487, "epm.paused invalid");
        } else {
            assertTrue(eigenPodManager.paused() == 0, "epm.paused invalid");
        }
    }

    /// @dev Validates the ethPOS on the eigenPodImpl, an exiting EigenPod, and if a new pod is deployed
    function _validateEigenPodETHPos() internal {
        // Ignore non-Hoodi chains
        if (block.chainid != 560_048) return;

        // Check the ethPOS on the eigenPodImpl
        EigenPod eigenPod = EigenPod(payable(Env.beacon.eigenPod().implementation()));
        assertTrue(
            address(eigenPod.ethPOS()) == 0x00000000219ab540356cBB839Cbe05303d7705Fa, "hoodi ETH POS invalid"
        );

        // Check the ethPOS on an existing EigenPod
        EigenPod existingEigenPod = EigenPod(payable(0x916966e12bB58bdd17c70514A3d9DeFD65D294Dc));
        assertTrue(
            address(existingEigenPod.ethPOS()) == 0x00000000219ab540356cBB839Cbe05303d7705Fa,
            "hoodi ETH POS invalid"
        );

        // Deploy a new pod
        EigenPodManager eigenPodManager = Env.proxy.eigenPodManager();
        address newPod = eigenPodManager.createPod();
        EigenPod newEigenPod = EigenPod(payable(newPod));
        assertTrue(
            address(newEigenPod.ethPOS()) == 0x00000000219ab540356cBB839Cbe05303d7705Fa, "hoodi ETH POS invalid"
        );
    }
}
