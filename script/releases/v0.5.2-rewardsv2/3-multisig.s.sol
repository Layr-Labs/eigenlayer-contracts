// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "zeus-templates/utils/Encode.sol";
import {Queue} from "./2-multisig.s.sol";
import {RewardsCoordinator} from "src/contracts/core/RewardsCoordinator.sol";
import {IPauserRegistry} from "src/contracts/interfaces/IPauserRegistry.sol";
import {DelegationManager} from "src/contracts/core/DelegationManager.sol";
import {StrategyManager} from "src/contracts/core/StrategyManager.sol";
import "../Env.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract Execute is Queue {
    using Env for *;
    using Encode for *;
    using ZEnvHelpers for *;

    event Upgraded(address indexed implementation);

    /**
     * @dev Overrides the previous _execute function to execute the queued transactions.
     */
    function _runAsMultisig() prank(Env.protocolCouncilMultisig()) internal override {
        bytes memory executorMultisigCalldata = _getCalldataToExecutor();
        TimelockController timelock = Env.timelockController();
        timelock.execute(
            Env.executorMultisig(),
            0 /* value */,
            executorMultisigCalldata,
            0 /* predecessor */,
            bytes32(0) /* salt */
        );
    }

    function testDeploy() public override {
        // save the previous implementation address to assert its change later
        address prevRewardsCoordinator = address(Env.impl.rewardsCoordinator());

        // 0. Deploy the Implementation contract.
        runAsEOA();

        // 1. run the queue script.
        super._runAsMultisig();

        State storage state = ZEnvHelpers.state();

        RewardsCoordinator rewardsCoordinatorProxy = Env.proxy.rewardsCoordinator();
        uint256 pausedStatusBefore = rewardsCoordinatorProxy.paused();
        TimelockController timelock = Env.timelockController();

        // 2. run the execute script above.
        bytes memory multisigTxnData = _getCalldataToExecutor();
        bytes32 txHash = timelock.hashOperation(Env.executorMultisig(), 0, multisigTxnData, 0, 0);

        assertEq(timelock.isOperationPending(txHash), true, "Transaction should be queued and pending.");
        vm.warp(block.timestamp + timelock.getMinDelay()); // 1 tick after ETA.

        assertEq(timelock.isOperationReady(txHash), true, "Transaction should be executable.");

        vm.expectEmit(true, true, true, true, address(rewardsCoordinatorProxy));
        emit Upgraded(address(Env.impl.rewardsCoordinator()));
        execute();

        // 3. assert that the execute did something
        assertEq(timelock.isOperationDone(txHash), true, "Transaction should be executed.");

        // assert that the proxy implementation was updated.
        ProxyAdmin admin = ProxyAdmin(Env.proxyAdmin());
        address rewardsCoordinatorImpl = admin.getProxyImplementation(
            ITransparentUpgradeableProxy(payable(address(Env.proxy.rewardsCoordinator())))
        );
        assertEq(rewardsCoordinatorImpl, address(Env.impl.rewardsCoordinator()));
        assertNotEq(prevRewardsCoordinator, rewardsCoordinatorImpl, "expected rewardsCoordinatorImpl to be different");

        uint256 pausedStatusAfter = rewardsCoordinatorProxy.paused();
        address owner = Env.opsMultisig();
        IPauserRegistry pauserRegistry = IPauserRegistry(Env.impl.pauserRegistry());
        uint64 initPausedStatus = state.envU64("REWARDS_COORDINATOR_INIT_PAUSED_STATUS");
        address rewardsUpdater = state.envAddress("REWARDS_COORDINATOR_UPDATER");
        uint32 activationDelay = state.envU32("REWARDS_COORDINATOR_ACTIVATION_DELAY");
        uint16 defaultOperatorSplitBips = state.envU16("REWARDS_COORDINATOR_DEFAULT_OPERATOR_SPLIT_BIPS");

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
            address(Env.proxy.delegationManager()),
            "expected delegationManager"
        );
        assertEq(
            address(rewardsCoordinatorProxy.strategyManager()),
            address(Env.proxy.strategyManager()),
            "expected strategyManager"
        );

        assertEq(
            rewardsCoordinatorProxy.CALCULATION_INTERVAL_SECONDS(),
            state.envU32("REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"),
            "expected REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"
        );
        assertGt(
            rewardsCoordinatorProxy.CALCULATION_INTERVAL_SECONDS(),
            0,
            "expected non-zero REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"
        );

        assertEq(rewardsCoordinatorProxy.MAX_REWARDS_DURATION(), state.envU32("REWARDS_COORDINATOR_MAX_REWARDS_DURATION"));
        assertGt(rewardsCoordinatorProxy.MAX_REWARDS_DURATION(), 0);

        assertEq(
            rewardsCoordinatorProxy.MAX_RETROACTIVE_LENGTH(),
            state.envU32("REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH")
        );
        assertGt(rewardsCoordinatorProxy.MAX_RETROACTIVE_LENGTH(), 0);

        assertEq(rewardsCoordinatorProxy.MAX_FUTURE_LENGTH(), state.envU32("REWARDS_COORDINATOR_MAX_FUTURE_LENGTH"));
        assertGt(rewardsCoordinatorProxy.MAX_FUTURE_LENGTH(), 0);

        assertEq(
            rewardsCoordinatorProxy.GENESIS_REWARDS_TIMESTAMP(),
            state.envU32("REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP")
        );
        assertGt(rewardsCoordinatorProxy.GENESIS_REWARDS_TIMESTAMP(), 0);
    }
}
