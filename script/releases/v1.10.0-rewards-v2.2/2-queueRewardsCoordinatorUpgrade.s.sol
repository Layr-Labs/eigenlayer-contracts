// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {DeployRewardsCoordinatorImpl} from "./1-deployRewardsCoordinatorImpl.s.sol";
import "../Env.sol";
import "../TestUtils.sol";

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {MultisigCall, Encode} from "zeus-templates/utils/Encode.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";

/**
 * @title QueueRewardsCoordinatorUpgrade
 * @notice Queue the RewardsCoordinator upgrade transaction in the Timelock via the Operations Multisig.
 *         This queues the upgrade to add Rewards v2.2 support:
 *         - Unique stake rewards (linear to allocated unique stake)
 *         - Total stake rewards (linear to total stake)
 *         - Updated MAX_FUTURE_LENGTH to 730 days
 */
contract QueueRewardsCoordinatorUpgrade is MultisigBuilder, DeployRewardsCoordinatorImpl {
    using Env for *;
    using Encode for *;
    using TestUtils for *;

    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        if (!(Env.isSourceChain() && Env._strEq(Env.envVersion(), "1.9.0"))) {
            return;
        }

        bytes memory calldata_to_executor = _getCalldataToExecutor();

        TimelockController timelock = Env.timelockController();
        timelock.schedule({
            target: Env.executorMultisig(),
            value: 0,
            data: calldata_to_executor,
            predecessor: 0,
            salt: 0,
            delay: timelock.getMinDelay()
        });
    }

    /// @dev Get the calldata to be sent from the timelock to the executor
    function _getCalldataToExecutor() internal returns (bytes memory) {
        /// forgefmt: disable-next-item
        MultisigCall[] storage executorCalls = Encode.newMultisigCalls().append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.rewardsCoordinator()),
                impl: address(Env.impl.rewardsCoordinator())
            })
        });

        return Encode.gnosisSafe.execTransaction({
            from: address(Env.timelockController()),
            to: Env.multiSendCallOnly(),
            op: Encode.Operation.DelegateCall,
            data: Encode.multiSend(executorCalls)
        });
    }

    function testScript() public virtual override {
        if (!(Env.isSourceChain() && Env._strEq(Env.envVersion(), "1.9.0"))) {
            return;
        }

        // 1 - Deploy. The new RewardsCoordinator implementation has been deployed
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

        // Ensure transaction is not already queued
        assertFalse(timelock.isOperationPending(txHash), "Transaction should not be queued yet");
        assertFalse(timelock.isOperationReady(txHash), "Transaction should not be ready");
        assertFalse(timelock.isOperationDone(txHash), "Transaction should not be complete");

        // 2 - Queue the transaction
        _runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        // Verify transaction is queued
        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued");
        assertFalse(timelock.isOperationReady(txHash), "Transaction should NOT be ready immediately");
        assertFalse(timelock.isOperationDone(txHash), "Transaction should NOT be complete");

        // Validate that timelock delay is properly configured
        uint256 minDelay = timelock.getMinDelay();
        assertTrue(minDelay > 0, "Timelock delay should be greater than 0");

        // Validate proxy state before upgrade
        _validatePreUpgradeState();
    }

    /// @dev Validate the state before the upgrade
    function _validatePreUpgradeState() internal view {
        RewardsCoordinator rewardsCoordinator = Env.proxy.rewardsCoordinator();

        // Validate current implementation is different from new implementation
        address currentImpl = TestUtils._getProxyImpl(address(rewardsCoordinator));
        address newImpl = address(Env.impl.rewardsCoordinator());
        assertTrue(currentImpl != newImpl, "Current and new implementations should be different");

        // Validate current MAX_FUTURE_LENGTH is different from new value
        // Note: We access this through the proxy, which still points to the old implementation
        uint32 currentMaxFutureLength = rewardsCoordinator.MAX_FUTURE_LENGTH();
        assertTrue(currentMaxFutureLength != 63_072_000, "Current MAX_FUTURE_LENGTH should not be 730 days yet");

        // Note: Cannot check version() on old implementation as it may not have this function
        // The upgrade is needed to add new functions and update MAX_FUTURE_LENGTH

        // Validate that we're upgrading from the correct version
        // We can't directly call them since they don't exist, but we can verify the upgrade is needed
        // by checking that we're indeed on the right version
        assertEq(Env.envVersion(), "1.9.0", "Should be on version 1.9.0 before upgrade");
    }
}
