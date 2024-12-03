// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigCall, MultisigCallUtils, MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {SafeTx, SafeTxUtils} from "zeus-templates/utils/SafeTxUtils.sol";
import {Queue} from "./2-multisig.s.sol";
import {EigenLabsUpgrade} from "../EigenLabsUpgrade.s.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import {RewardsCoordinator} from "src/contracts/core/RewardsCoordinator.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {IPauserRegistry} from "src/contracts/interfaces/IPauserRegistry.sol";
import {DelegationManager} from "src/contracts/core/DelegationManager.sol";
import {StrategyManager} from "src/contracts/core/StrategyManager.sol";
import {console} from "forge-std/console.sol";

contract Execute is Queue {
    using MultisigCallUtils for MultisigCall[];
    using SafeTxUtils for SafeTx;
    using EigenLabsUpgrade for *;

    /**
     * @dev Overrides the previous _execute function to execute the queued transactions.
     */
    function _runAsMultisig() internal override {
        bytes memory executorMultisigCalldata = _getMultisigTransactionCalldata();
        TimelockController timelock = TimelockController(payable(this._timelock()));
        timelock.execute(
            this._executorMultisig(),
            0 /* value */,
            executorMultisigCalldata,
            0 /* predecessor */,
            bytes32(0) /* salt */
        );
    }

    function testDeploy() public override {
        // save the previous implementation address to assert its change later
        address prevRewardsCoordinator = zDeployedImpl(type(RewardsCoordinator).name);

        // 0. Deploy the Implementation contract.
        runAsEOA();

        // 1. run the queue script.
        vm.startPrank(this._operationsMultisig());
        super._runAsMultisig();
        vm.stopPrank();

        RewardsCoordinator rewardsCoordinatorProxy = RewardsCoordinator(zDeployedProxy(type(RewardsCoordinator).name));
        uint256 pausedStatusBefore = rewardsCoordinatorProxy.paused();
        TimelockController timelock = this._timelock();

        // 2. run the execute script above.
        bytes memory multisigTxnData = _getMultisigTransactionCalldata();
        bytes32 txHash = timelock.hashOperation(this._executorMultisig(), 0, multisigTxnData, 0, 0);

        assertEq(timelock.isOperationPending(txHash), true, "Transaction should be queued and pending.");
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA.

        assertEq(timelock.isOperationReady(txHash), true, "Transaction should be executable.");

        zSetMultisigContext(this._protocolCouncilMultisig());
        execute();

        // 3. assert that the execute did something
        assertEq(timelock.isOperationDone(txHash), true, "Transaction should be executed.");

        // assert that the proxy implementation was updated.
        ProxyAdmin admin = ProxyAdmin(this._proxyAdmin());
        address rewardsCoordinatorImpl = admin.getProxyImplementation(
            TransparentUpgradeableProxy(payable(zDeployedProxy(type(RewardsCoordinator).name)))
        );
        assertEq(rewardsCoordinatorImpl, zDeployedImpl(type(RewardsCoordinator).name));
        assertNotEq(prevRewardsCoordinator, rewardsCoordinatorImpl, "expected rewardsCoordinatorImpl to be different");

        uint256 pausedStatusAfter = rewardsCoordinatorProxy.paused();
        address owner = this._operationsMultisig();
        IPauserRegistry pauserRegistry = IPauserRegistry(this._pauserRegistry());
        uint64 initPausedStatus = zUint64("REWARDS_COORDINATOR_INIT_PAUSED_STATUS");
        address rewardsUpdater = zAddress("REWARDS_COORDINATOR_UPDATER");
        uint32 activationDelay = zUint32("REWARDS_COORDINATOR_ACTIVATION_DELAY");
        uint16 defaultOperatorSplitBips = zUint16("REWARDS_COORDINATOR_DEFAULT_OPERATOR_SPLIT_BIPS");

        // Ensure that the proxy contract cannot be re-initialized.
        vm.expectRevert("Initializable: contract is already initialized");
        rewardsCoordinatorProxy.initialize(
            owner,
            pauserRegistry,
            initPausedStatus,
            rewardsUpdater,
            activationDelay,
            defaultOperatorSplitBips
        );

        // Assert Immutables and State Variables set through initialize
        assertEq(rewardsCoordinatorProxy.owner(), owner, "expected owner");
        assertEq(address(rewardsCoordinatorProxy.pauserRegistry()), address(pauserRegistry), "expected pauserRegistry");
        assertEq(address(rewardsCoordinatorProxy.rewardsUpdater()), rewardsUpdater, "expected rewardsUpdater");
        assertEq(rewardsCoordinatorProxy.activationDelay(), activationDelay, "expected activationDelay");
        assertEq(
            rewardsCoordinatorProxy.defaultOperatorSplitBips(),
            defaultOperatorSplitBips,
            "expected defaultOperatorSplitBips"
        );
        assertEq(
            pausedStatusBefore,
            pausedStatusAfter,
            "expected paused status to be the same before and after initialization"
        );
        assertEq(
            address(rewardsCoordinatorProxy.delegationManager()),
            zDeployedProxy(type(DelegationManager).name),
            "expected delegationManager"
        );
        assertEq(
            address(rewardsCoordinatorProxy.strategyManager()),
            zDeployedProxy(type(StrategyManager).name),
            "expected strategyManager"
        );

        assertEq(
            rewardsCoordinatorProxy.CALCULATION_INTERVAL_SECONDS(),
            zUint32("REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"),
            "expected REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"
        );
        assertGt(
            rewardsCoordinatorProxy.CALCULATION_INTERVAL_SECONDS(),
            0,
            "expected non-zero REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"
        );

        assertEq(rewardsCoordinatorProxy.MAX_REWARDS_DURATION(), zUint32("REWARDS_COORDINATOR_MAX_REWARDS_DURATION"));
        assertGt(rewardsCoordinatorProxy.MAX_REWARDS_DURATION(), 0);

        assertEq(
            rewardsCoordinatorProxy.MAX_RETROACTIVE_LENGTH(),
            zUint32("REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH")
        );
        assertGt(rewardsCoordinatorProxy.MAX_RETROACTIVE_LENGTH(), 0);

        assertEq(rewardsCoordinatorProxy.MAX_FUTURE_LENGTH(), zUint32("REWARDS_COORDINATOR_MAX_FUTURE_LENGTH"));
        assertGt(rewardsCoordinatorProxy.MAX_FUTURE_LENGTH(), 0);

        assertEq(
            rewardsCoordinatorProxy.GENESIS_REWARDS_TIMESTAMP(),
            zUint32("REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP")
        );
        assertGt(rewardsCoordinatorProxy.GENESIS_REWARDS_TIMESTAMP(), 0);
    }
}
