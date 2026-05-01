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
import {IBackingEigen} from "src/contracts/interfaces/IBackingEigen.sol";
import {EmissionsController} from "src/contracts/core/EmissionsController.sol";
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
        bool needsIncentiveCouncil = _needsIncentiveCouncilUpgrade();

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
        _validateStrategyFactoryUpgrade();
        _validateDurationVaultStrategyUpgrade();
        _validateIncentiveCouncilCatchUp(needsIncentiveCouncil);

        // Run standard proxy validations
        TestUtils.validateProxyAdmins();
        TestUtils.validateProxyConstructors();
        TestUtils.validateProxyStorage();
        TestUtils.validateImplAddressesMatchProxy();

        // Duration vault validations
        TestUtils.validateDurationVaultStrategyProxyAdmin();
        TestUtils.validateDurationVaultStrategyStorage();
        TestUtils.validateDurationVaultStrategyImplConstructors();
        TestUtils.validateDurationVaultStrategyImplAddressesMatchProxy();
        TestUtils.validateDurationVaultStrategyProtocolRegistry();

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

    function _validateStrategyFactoryUpgrade() internal view {
        StrategyFactory strategyFactory = Env.proxy.strategyFactory();

        // Validate implementation address
        address actualImpl = ProxyAdmin(Env.proxyAdmin())
            .getProxyImplementation(ITransparentUpgradeableProxy(payable(address(strategyFactory))));
        assertEq(actualImpl, address(Env.impl.strategyFactory()), "StrategyFactory implementation incorrect");

        // Validate new EIGEN/bEIGEN immutables
        TestUtils.validateStrategyFactoryImmutables(strategyFactory);

        // Validate storage unchanged
        assertEq(strategyFactory.owner(), Env.opsMultisig(), "StrategyFactory owner changed");
        assertEq(strategyFactory.paused(), 0, "StrategyFactory paused status changed");
    }

    function _validateDurationVaultStrategyUpgrade() internal view {
        // Validate beacon points to new implementation
        assertEq(
            Env.beacon.durationVaultStrategy().implementation(),
            address(Env.impl.durationVaultStrategy()),
            "DurationVaultStrategy beacon implementation incorrect"
        );

        // Validate new implementation immutables
        TestUtils.validateDurationVaultStrategyImmutables(Env.impl.durationVaultStrategy());
    }

    /// @notice Validates v1.12.0 incentive council catch-up if it was applied.
    function _validateIncentiveCouncilCatchUp(
        bool wasApplied
    ) internal view {
        if (!wasApplied) return;

        // RewardsCoordinator should now have correct impl and feeRecipient
        RewardsCoordinator rc = Env.proxy.rewardsCoordinator();
        address rcImpl =
            ProxyAdmin(Env.proxyAdmin()).getProxyImplementation(ITransparentUpgradeableProxy(payable(address(rc))));
        assertEq(rcImpl, address(Env.impl.rewardsCoordinator()), "RewardsCoordinator implementation incorrect");
        assertEq(rc.feeRecipient(), Env.incentiveCouncilMultisig(), "RewardsCoordinator feeRecipient incorrect");
        TestUtils.validateRewardsCoordinatorImmutables(rc);

        // EmissionsController should be initialized
        EmissionsController ec = Env.proxy.emissionsController();
        address ecImpl =
            ProxyAdmin(Env.proxyAdmin()).getProxyImplementation(ITransparentUpgradeableProxy(payable(address(ec))));
        assertEq(ecImpl, address(Env.impl.emissionsController()), "EmissionsController implementation incorrect");
        assertEq(ec.owner(), Env.opsMultisig(), "EmissionsController owner incorrect");
        TestUtils.validateEmissionsControllerImmutables(ec);

        // bEIGEN minting rights transferred
        IBackingEigen beigen = Env.proxy.beigen();
        assertTrue(beigen.isMinter(address(ec)), "EmissionsController should have minting rights");
        if (Env.legacyTokenHopper() != address(0)) {
            assertFalse(beigen.isMinter(Env.legacyTokenHopper()), "Legacy hopper should NOT have minting rights");
        }
    }
}
