// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {QueueUpgrade} from "./2-queueUpgrade.s.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import {DelegationManager} from "src/contracts/core/DelegationManager.sol";
import {EmissionsController} from "src/contracts/core/EmissionsController.sol";
import {RewardsCoordinator} from "src/contracts/core/RewardsCoordinator.sol";
import {EigenPod} from "src/contracts/pods/EigenPod.sol";
import {EigenPodManager} from "src/contracts/pods/EigenPodManager.sol";
import {BackingEigen} from "src/contracts/token/BackingEigen.sol";
import "../Env.sol";
import "../TestUtils.sol";

contract ExecuteUpgrade is QueueUpgrade {
    using Env for *;

    struct PreUpgradeState {
        bool needsIncentiveCouncil;
        uint256 delegationManagerPaused;
        uint32 minWithdrawalDelayBlocks;
        address eigenPodManagerOwner;
        uint256 eigenPodManagerPaused;
        address proofTimestampSetter;
        uint64 pectraForkTimestamp;
        address emissionsControllerOwner;
        address incentiveCouncil;
        uint256 emissionsControllerPaused;
        uint16 totalWeight;
    }

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

    function testScript() public virtual override onlyIfUpgradeRequired("1.14.0") {
        PreUpgradeState memory state = _capturePreUpgradeState();

        super.runAsEOA();

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
        _unsafeResetHasPranked();

        assertTrue(timelock.isOperationPending(txHash), "Transaction should be queued.");
        assertFalse(timelock.isOperationReady(txHash), "Transaction should NOT be ready for execution.");
        assertFalse(timelock.isOperationDone(txHash), "Transaction should NOT be complete.");

        vm.warp(block.timestamp + timelock.getMinDelay());
        assertTrue(timelock.isOperationReady(txHash), "Transaction should be executable.");

        execute();

        assertTrue(timelock.isOperationDone(txHash), "Transaction should be complete.");
        _validateDelegationManagerUpgrade(state);
        _validateEigenPodManagerUpgrade(state);
        _validateEigenPodUpgrade();
        _validateEmissionsControllerUpgrade(state);
        _validateBackingEigenPermissions();
        _validateIncentiveCouncilCatchUp(state.needsIncentiveCouncil);

        TestUtils.validateProxyAdmins();
        if (_strategyManagerHasSlashResolutionDelayApi()) {
            TestUtils.validateProxyConstructors();
        }
        TestUtils.validateImplAddressesMatchProxy();

        // Run last since it calls pauseAll().
        TestUtils.validateProtocolRegistry();
    }

    function _capturePreUpgradeState() internal view returns (PreUpgradeState memory state) {
        state.needsIncentiveCouncil = _needsIncentiveCouncilUpgrade();

        DelegationManager delegationManager = Env.proxy.delegationManager();
        state.delegationManagerPaused = delegationManager.paused();
        state.minWithdrawalDelayBlocks = delegationManager.minWithdrawalDelayBlocks();

        EigenPodManager eigenPodManager = Env.proxy.eigenPodManager();
        state.eigenPodManagerOwner = eigenPodManager.owner();
        state.eigenPodManagerPaused = eigenPodManager.paused();
        state.proofTimestampSetter = eigenPodManager.proofTimestampSetter();
        state.pectraForkTimestamp = eigenPodManager.pectraForkTimestamp();

        if (!state.needsIncentiveCouncil) {
            EmissionsController emissionsController = Env.proxy.emissionsController();
            state.emissionsControllerOwner = emissionsController.owner();
            state.incentiveCouncil = emissionsController.incentiveCouncil();
            state.emissionsControllerPaused = emissionsController.paused();
            state.totalWeight = emissionsController.totalWeight();
        }
    }

    function _validateDelegationManagerUpgrade(
        PreUpgradeState memory state
    ) internal view {
        DelegationManager delegationManager = Env.proxy.delegationManager();
        address actualImpl = ProxyAdmin(Env.proxyAdmin())
            .getProxyImplementation(ITransparentUpgradeableProxy(payable(address(delegationManager))));

        assertEq(actualImpl, address(Env.impl.delegationManager()), "DelegationManager implementation incorrect");
        assertEq(delegationManager.version(), Env.deployVersion(), "DelegationManager version incorrect");
        assertEq(delegationManager.paused(), state.delegationManagerPaused, "DelegationManager paused status changed");
        assertEq(
            delegationManager.minWithdrawalDelayBlocks(),
            state.minWithdrawalDelayBlocks,
            "DelegationManager withdrawal delay changed"
        );
        TestUtils.validateDelegationManagerImmutables(delegationManager);
    }

    function _validateEigenPodManagerUpgrade(
        PreUpgradeState memory state
    ) internal view {
        EigenPodManager eigenPodManager = Env.proxy.eigenPodManager();
        address actualImpl = ProxyAdmin(Env.proxyAdmin())
            .getProxyImplementation(ITransparentUpgradeableProxy(payable(address(eigenPodManager))));

        assertEq(actualImpl, address(Env.impl.eigenPodManager()), "EigenPodManager implementation incorrect");
        assertEq(eigenPodManager.owner(), state.eigenPodManagerOwner, "EigenPodManager owner changed");
        assertEq(eigenPodManager.paused(), state.eigenPodManagerPaused, "EigenPodManager paused status changed");
        assertEq(
            eigenPodManager.proofTimestampSetter(),
            state.proofTimestampSetter,
            "EigenPodManager proof timestamp setter changed"
        );
        assertEq(
            eigenPodManager.pectraForkTimestamp(),
            state.pectraForkTimestamp,
            "EigenPodManager Pectra fork timestamp changed"
        );
        TestUtils.validateEigenPodManagerImmutables(eigenPodManager);
    }

    function _validateEigenPodUpgrade() internal view {
        assertEq(
            Env.beacon.eigenPod().implementation(),
            address(Env.impl.eigenPod()),
            "EigenPod beacon implementation incorrect"
        );
        TestUtils.validateEigenPodImmutables(EigenPod(payable(address(Env.impl.eigenPod()))));
    }

    function _validateEmissionsControllerUpgrade(
        PreUpgradeState memory state
    ) internal view {
        EmissionsController ec = Env.proxy.emissionsController();
        ProxyAdmin proxyAdmin = ProxyAdmin(Env.proxyAdmin());
        ITransparentUpgradeableProxy proxy = ITransparentUpgradeableProxy(payable(address(ec)));
        address actualImpl = proxyAdmin.getProxyImplementation(proxy);

        assertEq(actualImpl, address(Env.impl.emissionsController()), "EC implementation incorrect");
        if (state.needsIncentiveCouncil) {
            assertEq(ec.owner(), Env.opsMultisig(), "EC owner incorrect");
            assertEq(ec.incentiveCouncil(), Env.incentiveCouncilMultisig(), "EC incentive council incorrect");
            assertEq(ec.paused(), 0, "EC paused status incorrect");
            assertEq(ec.totalWeight(), 0, "EC total weight incorrect");
        } else {
            assertEq(ec.owner(), state.emissionsControllerOwner, "EC owner changed");
            assertEq(ec.incentiveCouncil(), state.incentiveCouncil, "EC incentive council changed");
            assertEq(ec.paused(), state.emissionsControllerPaused, "EC paused status changed");
            assertEq(ec.totalWeight(), state.totalWeight, "EC total weight changed");
        }
        TestUtils.validateEmissionsControllerImmutables(ec);
    }

    function _validateBackingEigenPermissions() internal view {
        EmissionsController ec = Env.proxy.emissionsController();
        BackingEigen backingEigen = BackingEigen(address(Env.proxy.beigen()));

        assertTrue(backingEigen.isMinter(address(ec)), "EC should be a bEIGEN minter");
        assertTrue(backingEigen.allowedFrom(address(ec)), "EC should be allowed to burn bEIGEN");
    }

    function _validateIncentiveCouncilCatchUp(
        bool wasApplied
    ) internal view {
        if (!wasApplied) return;

        RewardsCoordinator rc = Env.proxy.rewardsCoordinator();
        ProxyAdmin proxyAdmin = ProxyAdmin(Env.proxyAdmin());
        ITransparentUpgradeableProxy proxy = ITransparentUpgradeableProxy(payable(address(rc)));
        address actualImpl = proxyAdmin.getProxyImplementation(proxy);

        assertEq(actualImpl, address(Env.impl.rewardsCoordinator()), "RC implementation incorrect");
        assertEq(rc.feeRecipient(), Env.incentiveCouncilMultisig(), "RC fee recipient incorrect");
        TestUtils.validateRewardsCoordinatorImmutables(rc);

        BackingEigen backingEigen = BackingEigen(address(Env.proxy.beigen()));
        if (Env.legacyTokenHopper() != address(0)) {
            assertFalse(backingEigen.isMinter(Env.legacyTokenHopper()), "Legacy hopper should not be a bEIGEN minter");
        }
    }

    function _strategyManagerHasSlashResolutionDelayApi() internal view returns (bool) {
        (bool ok,) =
            address(Env.proxy.strategyManager()).staticcall(abi.encodeWithSignature("SLASH_RESOLUTION_DELAY_BLOCKS()"));
        return ok;
    }
}
