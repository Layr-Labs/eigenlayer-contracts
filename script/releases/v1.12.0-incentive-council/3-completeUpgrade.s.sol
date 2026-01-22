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

import {EmissionsController} from "src/contracts/core/EmissionsController.sol";
import {RewardsCoordinator} from "src/contracts/core/RewardsCoordinator.sol";
import {IProtocolRegistry, IProtocolRegistryTypes} from "src/contracts/interfaces/IProtocolRegistry.sol";
import {IBackingEigen} from "src/contracts/interfaces/IBackingEigen.sol";

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

        // Warp Past delay
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA
        assertEq(timelock.isOperationReady(txHash), true, "Transaction should be executable.");

        // Execute
        execute();

        // Run Tests
        TestUtils.validateProxyAdmins();
        TestUtils.validateProxyConstructors();
        // TestUtils.validateProxiesAlreadyInitialized();
        TestUtils.validateProxyStorage();
        TestUtils.validateImplAddressesMatchProxy();

        // Additional v9.9.9-specific validation (must run before validateProtocolRegistry which pauses everything)
        _validateEmissionsControllerUpgrade();
        _validateRewardsCoordinatorUpgrade();
        _validateMintingRights();

        // Run last since it calls pauseAll()
        TestUtils.validateProtocolRegistry();
    }

    /// @notice Validates that the EmissionsController was upgraded and initialized correctly
    function _validateEmissionsControllerUpgrade() internal view {
        EmissionsController ec = Env.proxy.emissionsController();

        // Validate proxy admin
        address expectedProxyAdmin = Env.proxyAdmin();
        address actualProxyAdmin =
            ProxyAdmin(expectedProxyAdmin).getProxyAdmin(ITransparentUpgradeableProxy(payable(address(ec))));
        assertEq(actualProxyAdmin, expectedProxyAdmin, "EC: proxy admin incorrect");

        // Validate implementation
        address expectedImpl = address(Env.impl.emissionsController());
        address actualImpl =
            ProxyAdmin(expectedProxyAdmin).getProxyImplementation(ITransparentUpgradeableProxy(payable(address(ec))));
        assertEq(actualImpl, expectedImpl, "EC: implementation incorrect");

        // Validate initialization state
        assertEq(ec.owner(), Env.opsMultisig(), "EC: owner incorrect");
        assertEq(ec.incentiveCouncil(), Env.incentiveCouncilMultisig(), "EC: incentiveCouncil incorrect");
        assertEq(ec.paused(), 0, "EC: paused status incorrect");

        // Validate immutables
        TestUtils.validateEmissionsControllerImmutables(ec);

        // Validate protocol registry entry
        (address addr, IProtocolRegistryTypes.DeploymentConfig memory config) =
            Env.proxy.protocolRegistry().getDeployment(type(EmissionsController).name);
        assertEq(addr, address(ec), "EC: protocol registry address incorrect");
        assertTrue(config.pausable, "EC: should be pausable in registry");
        assertFalse(config.deprecated, "EC: should not be deprecated in registry");
    }

    /// @notice Validates that the RewardsCoordinator was upgraded and reinitialized correctly
    function _validateRewardsCoordinatorUpgrade() internal view {
        RewardsCoordinator rc = Env.proxy.rewardsCoordinator();

        // Validate implementation
        address expectedImpl = address(Env.impl.rewardsCoordinator());
        address actualImpl =
            ProxyAdmin(Env.proxyAdmin()).getProxyImplementation(ITransparentUpgradeableProxy(payable(address(rc))));
        assertEq(actualImpl, expectedImpl, "RC: implementation incorrect");

        // Validate reinitialization state
        assertEq(rc.owner(), Env.opsMultisig(), "RC: owner incorrect");
        assertEq(rc.rewardsUpdater(), Env.REWARDS_UPDATER(), "RC: rewardsUpdater incorrect");
        assertEq(rc.activationDelay(), Env.ACTIVATION_DELAY(), "RC: activationDelay incorrect");
        assertEq(rc.defaultOperatorSplitBips(), Env.DEFAULT_SPLIT_BIPS(), "RC: defaultOperatorSplitBips incorrect");
        assertEq(rc.feeRecipient(), Env.incentiveCouncilMultisig(), "RC: feeRecipient incorrect");
        assertEq(rc.paused(), 0, "RC: paused status incorrect");

        // Validate immutables still intact
        TestUtils.validateRewardsCoordinatorImmutables(rc);
    }

    /// @notice Validates that minting rights were correctly transferred
    function _validateMintingRights() internal view {
        IBackingEigen beigen = Env.proxy.beigen();

        // Validate EmissionsController has minting rights
        assertTrue(
            beigen.isMinter(address(Env.proxy.emissionsController())), "EmissionsController should have minting rights"
        );

        // Validate old hopper does NOT have minting rights
        assertFalse(beigen.isMinter(Env.legacyTokenHopper()), "Old hopper should NOT have minting rights");
    }
}

