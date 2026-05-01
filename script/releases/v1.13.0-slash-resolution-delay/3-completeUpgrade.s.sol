// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {QueueUpgrade} from "./2-queueUpgrade.s.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";
import {IProxyAdmin} from "zeus-templates/interfaces/IProxyAdmin.sol";
import {
    ITransparentUpgradeableProxy as ITransparentProxy
} from "zeus-templates/interfaces/ITransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "../Env.sol";
import "../TestUtils.sol";

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
        if (!Env.isCoreProtocolDeployed()) {
            return;
        }

        // Capture pre-upgrade state
        uint256 pausedStatusBefore = Env.proxy.strategyManager().paused();
        address ownerBefore = Env.proxy.strategyManager().owner();
        address strategyWhitelisterBefore = Env.proxy.strategyManager().strategyWhitelister();

        // Deploy the implementations (from previous step 1)
        super.runAsEOA();

        // Queue the upgrade (from previous step 2)
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

        // Warp past delay
        vm.warp(block.timestamp + timelock.getMinDelay());
        assertEq(timelock.isOperationReady(txHash), true, "Transaction should be executable.");

        // Execute
        execute();

        assertTrue(timelock.isOperationDone(txHash), "Transaction should be complete.");
        _validateStrategyManagerUpgrade(pausedStatusBefore, ownerBefore, strategyWhitelisterBefore);

        // Run standard proxy validations
        TestUtils.validateProxyAdmins();
        TestUtils.validateProxyConstructors();
        TestUtils.validateProxyStorage();
        TestUtils.validateImplAddressesMatchProxy();

        // Run last since it calls pauseAll()
        TestUtils.validateProtocolRegistry();
    }

    function _validateStrategyManagerUpgrade(
        uint256 pausedStatusBefore,
        address ownerBefore,
        address strategyWhitelisterBefore
    ) internal view {
        StrategyManager strategyManager = Env.proxy.strategyManager();

        // Validate implementation address
        address actualImpl = ProxyAdmin(Env.proxyAdmin())
            .getProxyImplementation(ITransparentUpgradeableProxy(payable(address(strategyManager))));
        assertEq(actualImpl, address(Env.impl.strategyManager()), "StrategyManager implementation incorrect");
        assertEq(strategyManager.version(), Env.deployVersion(), "StrategyManager version incorrect");
        assertEq(strategyManager.owner(), ownerBefore, "StrategyManager owner changed");
        assertEq(
            strategyManager.strategyWhitelister(), strategyWhitelisterBefore, "StrategyManager whitelister changed"
        );
        TestUtils.validateStrategyManagerImmutables(strategyManager);
        TestUtils.validateStrategyManagerPausedStatusUnchanged(strategyManager, pausedStatusBefore);
        TestUtils.validateStrategyManagerSlashResolutionDelay(strategyManager);
    }
}
