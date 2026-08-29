// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {QueueUpgrade} from "./2-queueUpgrade.s.sol";
import {CrosschainDeployLib, createx} from "../CrosschainDeployLib.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import {DelegationManager} from "src/contracts/core/DelegationManager.sol";
import "../Env.sol";
import "../TestUtils.sol";

contract ExecuteUpgrade is QueueUpgrade {
    using Env for *;

    function _runAsMultisig() internal virtual override prank(Env.protocolCouncilMultisig()) {
        address implementation = createx.deployCreate2(
            CrosschainDeployLib.computeProtectedSalt(Env.protocolCouncilMultisig(), type(DelegationManager).name),
            _delegationManagerInitCode()
        );
        require(implementation == address(Env.impl.delegationManager()), "unexpected implementation address");

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

        // Capture pre-upgrade DelegationManager state.
        DelegationManager delegationManager = Env.proxy.delegationManager();
        uint256 pausedStatusBefore = delegationManager.paused();
        uint32 minWithdrawalDelayBefore = delegationManager.minWithdrawalDelayBlocks();

        // Precompute and register the implementation address (from previous step 1).
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
        assertTrue(address(Env.impl.delegationManager()).code.length != 0, "implementation was not deployed");
        TestUtils.validateDelegationManagerInitialized(Env.impl.delegationManager());
        TestUtils.validateDelegationManagerVersion();
        _validateDelegationManagerUpgrade(pausedStatusBefore, minWithdrawalDelayBefore);

        // Run standard proxy validations.
        TestUtils.validateProxyAdmins();

        // `validateProxyConstructors` and `validateProxyStorage` assume a post-v1.13.0
        // StrategyManager ABI (SLASH_RESOLUTION_DELAY_BLOCKS). If v1.13.0 is still queued
        // in an environment, these checks will revert and are not relevant to this release's
        // DelegationManager-only upgrade.
        if (_strategyManagerHasSlashResolutionDelayApi()) {
            TestUtils.validateProxyConstructors();
            TestUtils.validateProxyStorage();
        }

        TestUtils.validateImplAddressesMatchProxy();

        // Run last since it calls pauseAll().
        TestUtils.validateProtocolRegistry();
    }

    function _validateDelegationManagerUpgrade(
        uint256 pausedStatusBefore,
        uint32 minWithdrawalDelayBefore
    ) internal view {
        DelegationManager delegationManager = Env.proxy.delegationManager();

        // Validate implementation address.
        address actualImpl = ProxyAdmin(Env.proxyAdmin())
            .getProxyImplementation(ITransparentUpgradeableProxy(payable(address(delegationManager))));
        assertEq(actualImpl, address(Env.impl.delegationManager()), "DelegationManager implementation incorrect");

        // Validate version and storage continuity.
        assertEq(delegationManager.version(), Env.deployVersion(), "DelegationManager version incorrect");
        assertEq(delegationManager.paused(), pausedStatusBefore, "DelegationManager paused status changed");
        assertEq(
            delegationManager.minWithdrawalDelayBlocks(),
            minWithdrawalDelayBefore,
            "DelegationManager min withdrawal delay changed"
        );

        // Validate constructor immutables.
        TestUtils.validateDelegationManagerImmutables(delegationManager);
    }

    function _strategyManagerHasSlashResolutionDelayApi() internal view returns (bool) {
        (bool ok,) =
            address(Env.proxy.strategyManager()).staticcall(abi.encodeWithSignature("SLASH_RESOLUTION_DELAY_BLOCKS()"));
        return ok;
    }
}
