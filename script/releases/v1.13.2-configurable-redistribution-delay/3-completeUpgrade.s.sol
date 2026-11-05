// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {QueueUpgrade} from "./2-queueUpgrade.s.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import {IStrategyManagerErrors} from "src/contracts/interfaces/IStrategyManager.sol";
import "../Env.sol";
import "../TestUtils.sol";

/// Purpose: execute the queued StrategyManager upgrade (configurable slash resolution delay) via the timelock.
contract ExecuteUpgrade is QueueUpgrade {
    using Env for *;

    function _runAsMultisig() internal virtual override prank(Env.protocolCouncilMultisig()) {
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
        if (!Env.isCoreProtocolDeployed() || Env._versionGte(Env.envVersion(), Env.deployVersion())) {
            return;
        }

        StrategyManager strategyManager = Env.proxy.strategyManager();

        // Capture pre-upgrade StrategyManager state.
        uint256 pausedStatusBefore = strategyManager.paused();
        address ownerBefore = strategyManager.owner();
        address strategyWhitelisterBefore = strategyManager.strategyWhitelister();
        OperatorSet[] memory pendingOperatorSets = strategyManager.getPendingOperatorSets();
        uint256[][] memory pendingSlashIds = new uint256[][](pendingOperatorSets.length);
        uint32[][] memory resolutionBlocks = new uint32[][](pendingOperatorSets.length);
        for (uint256 i = 0; i < pendingOperatorSets.length; ++i) {
            pendingSlashIds[i] = strategyManager.getPendingSlashIds(pendingOperatorSets[i]);
            resolutionBlocks[i] = new uint32[](pendingSlashIds[i].length);
            for (uint256 j = 0; j < pendingSlashIds[i].length; ++j) {
                resolutionBlocks[i][j] =
                    strategyManager.getSlashResolutionBlock(pendingOperatorSets[i], pendingSlashIds[i][j]);
            }
        }

        // Deploy the implementation (from previous step 1).
        super.runAsEOA();

        // Queue the upgrade (from previous step 2).
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
        QueueUpgrade._runAsMultisig();
        _unsafeResetHasPranked(); // reset hasPranked so we can use it again

        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
        assertFalse(timelock.isOperationReady(txHash), "Transaction should NOT be ready for execution.");
        assertFalse(timelock.isOperationDone(txHash), "Transaction should NOT be complete.");

        // Warp past the timelock delay.
        vm.warp(block.timestamp + timelock.getMinDelay());
        assertEq(timelock.isOperationReady(txHash), true, "Transaction should be executable.");

        // Execute the queued upgrade.
        execute();

        assertTrue(timelock.isOperationDone(txHash), "Transaction should be complete.");
        _validateStrategyManagerUpgrade(pausedStatusBefore, ownerBefore, strategyWhitelisterBefore);
        _validatePendingSlashesPreserved(pendingOperatorSets, pendingSlashIds, resolutionBlocks);
        _validateSlashResolutionDelayConfig();

        // Run standard proxy validations.
        TestUtils.validateProxyAdmins();
        TestUtils.validateProxyConstructors();
        TestUtils.validateProxyStorage();
        TestUtils.validateImplAddressesMatchProxy();

        // Run last since it calls pauseAll().
        TestUtils.validateProtocolRegistry();
    }

    function _validateStrategyManagerUpgrade(
        uint256 pausedStatusBefore,
        address ownerBefore,
        address strategyWhitelisterBefore
    ) internal view {
        StrategyManager strategyManager = Env.proxy.strategyManager();

        // Validate implementation address.
        address actualImpl = ProxyAdmin(Env.proxyAdmin())
            .getProxyImplementation(ITransparentUpgradeableProxy(payable(address(strategyManager))));
        assertEq(actualImpl, address(Env.impl.strategyManager()), "StrategyManager implementation incorrect");

        // Validate version and storage continuity.
        assertEq(strategyManager.version(), Env.deployVersion(), "StrategyManager version incorrect");
        assertEq(strategyManager.owner(), ownerBefore, "StrategyManager owner changed");
        assertEq(
            strategyManager.strategyWhitelister(), strategyWhitelisterBefore, "StrategyManager whitelister changed"
        );
        TestUtils.validateStrategyManagerImmutables(strategyManager);
        TestUtils.validateStrategyManagerPausedStatusUnchanged(strategyManager, pausedStatusBefore);
    }

    /// @dev Pending slashes recorded before the upgrade keep their bookkeeping and their resolution block.
    function _validatePendingSlashesPreserved(
        OperatorSet[] memory pendingOperatorSets,
        uint256[][] memory pendingSlashIds,
        uint32[][] memory resolutionBlocks
    ) internal view {
        StrategyManager strategyManager = Env.proxy.strategyManager();

        OperatorSet[] memory pendingOperatorSetsAfter = strategyManager.getPendingOperatorSets();
        assertEq(pendingOperatorSetsAfter.length, pendingOperatorSets.length, "pending operator set count changed");

        for (uint256 i = 0; i < pendingOperatorSets.length; ++i) {
            assertEq(pendingOperatorSetsAfter[i].avs, pendingOperatorSets[i].avs, "pending operator set avs changed");
            assertEq(pendingOperatorSetsAfter[i].id, pendingOperatorSets[i].id, "pending operator set id changed");

            uint256[] memory slashIdsAfter = strategyManager.getPendingSlashIds(pendingOperatorSets[i]);
            assertEq(slashIdsAfter.length, pendingSlashIds[i].length, "pending slash id count changed");

            for (uint256 j = 0; j < pendingSlashIds[i].length; ++j) {
                assertEq(slashIdsAfter[j], pendingSlashIds[i][j], "pending slash id changed");
                assertEq(
                    strategyManager.getSlashResolutionBlock(pendingOperatorSets[i], pendingSlashIds[i][j]),
                    resolutionBlocks[i][j],
                    "pending slash resolution block changed"
                );
            }
        }
    }

    /// @dev The executor batch names the ops multisig as setter; the delay starts at zero until the ops multisig sets it.
    function _validateSlashResolutionDelayConfig() internal {
        StrategyManager strategyManager = Env.proxy.strategyManager();

        TestUtils.validateStrategyManagerSlashResolutionDelay(strategyManager);
        assertEq(
            strategyManager.slashResolutionDelaySetter(),
            Env.opsMultisig(),
            "slashResolutionDelaySetter should be the ops multisig"
        );
        assertEq(strategyManager.SLASH_RESOLUTION_DELAY_BLOCKS(), 0, "slash resolution delay should start at zero");

        // Only the ops multisig can change the delay; the executor (owner) cannot.
        vm.prank(Env.executorMultisig());
        vm.expectRevert(IStrategyManagerErrors.OnlySlashResolutionDelaySetter.selector);
        strategyManager.setSlashResolutionDelayBlocks(50_400);

        vm.prank(Env.opsMultisig());
        strategyManager.setSlashResolutionDelayBlocks(50_400);
        assertEq(
            strategyManager.SLASH_RESOLUTION_DELAY_BLOCKS(), 50_400, "ops multisig should be able to set the delay"
        );

        // Restore the as-deployed value so later validations observe post-upgrade state.
        vm.prank(Env.opsMultisig());
        strategyManager.setSlashResolutionDelayBlocks(0);
        assertEq(
            strategyManager.SLASH_RESOLUTION_DELAY_BLOCKS(), 0, "slash resolution delay should be restored to zero"
        );
    }
}
