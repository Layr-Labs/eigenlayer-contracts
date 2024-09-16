// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../utils/ExistingDeploymentParser.sol";
import "../../utils/Multisend.sol";
import "script/utils/TimelockEncoding.sol";

/**
 *
 * Mainnet: Deploy/Upgrade RewardsCoordinator
 * forge script script/deploy/mainnet/v0.4.3-upgrade-rewardsCoordinator.s.sol --rpc-url $MAINNET_RPC --private-key $PRIVATE_KEY --broadcast -vvvv --verify --etherscan-api-key $ETHERSCAN_API_KEY
 *
 * Test: forge test --mc Upgrade_Mainnet_RewardsCoordinator --mt test_set_reward_for_all_submitter --rpc-url $MAINNET_RPC -vv
 */
contract Upgrade_Mainnet_RewardsCoordinator is ExistingDeploymentParser, TimelockEncoding {

    // CALLDATA FOR CALL TO TIMELOCK
    // TUESDAY, SEPTEMBER 27 2024 22:00:00 GMT (6pm EST/3pm PST)
    uint256 timelockEta = 1727474400;

    uint256 dayToQueueAction = 1726610400;

    // Calldatas for upgrading RC
    bytes final_calldata_to_executor_multisig;
    bytes calldata_to_timelock_queuing_action;
    bytes calldata_to_timelock_executing_action;

    // Calldatas for setting reward for all submitter
    bytes hopper_setter_final_calldata_to_ops_multisig;

    modifier parseState() {
        _parseInitialDeploymentParams("script/configs/mainnet/mainnet-config.config.json");
        _parseDeployedContracts("script/configs/mainnet/mainnet-addresses.config.json");
        _;
    }

    function run() public parseState {
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        if (chainId != 1) {
            revert("Chain not supported");
        }
        
        RewardsCoordinator oldRewardsCoordinator = rewardsCoordinatorImplementation;

        // Deploy Rewards Coordinator
        vm.startBroadcast();
        rewardsCoordinatorImplementation = new RewardsCoordinator(
            delegationManager,
            strategyManager,
            REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS,
            REWARDS_COORDINATOR_MAX_REWARDS_DURATION,
            REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH,
            REWARDS_COORDINATOR_MAX_FUTURE_LENGTH,
            REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP
        );
        vm.stopBroadcast();

        _sanityCheckImplementations(oldRewardsCoordinator, rewardsCoordinatorImplementation);

        emit log_named_address("Rewards Coordinator Implementation", address(rewardsCoordinatorImplementation));

        // Create Upgrade Txs via Operations Multisig to Timelock to:
        // 1. Upgrade RewardsCoordinator
        // 2. Set owner of RewardsCoordinator to OperationsMultisig
        Tx[] memory txs = new Tx[](2);
        address timelockTarget = executorMultisig;

        // 1. Upgrade Rewards Coordiantor
        txs[0] = Tx({
            to: address(eigenLayerProxyAdmin),
            value: 0,
            data: abi.encodeWithSelector(
                ProxyAdmin.upgrade.selector,
                TransparentUpgradeableProxy(payable(address(rewardsCoordinator))),
                rewardsCoordinatorImplementation
            )
        });

        // 2. Set owner of RewardsCoordinator to OperationsMultisig
        txs[1] = Tx({
            to: address(rewardsCoordinator),
            value: 0,
            data: abi.encodeWithSelector(
                Ownable.transferOwnership.selector,
                address(operationsMultisig)
            )
        });

        bytes memory calldata_to_multisend_contract = abi.encodeWithSelector(MultiSendCallOnly.multiSend.selector, encodeMultisendTxs(txs));
        emit log_named_bytes("calldata_to_multisend_contract", calldata_to_multisend_contract);
        
        final_calldata_to_executor_multisig = encodeForExecutor({
            from: timelock,
            to: multiSendCallOnly,
            value: 0,
            data: calldata_to_multisend_contract,
            operation: ISafe.Operation.DelegateCall
        });

        calldata_to_timelock_queuing_action = abi.encodeWithSelector(ITimelock.queueTransaction.selector,
            timelockTarget,
            timelockValue,
            timelockSignature,
            final_calldata_to_executor_multisig,
            timelockEta
        );

        emit log_named_bytes("calldata_to_timelock_queuing_action", calldata_to_timelock_queuing_action);

        calldata_to_timelock_executing_action = abi.encodeWithSelector(ITimelock.executeTransaction.selector,
            timelockTarget,
            timelockValue,
            timelockSignature,
            final_calldata_to_executor_multisig,
            timelockEta
        );

        emit log_named_bytes("calldata_to_timelock_executing_action", calldata_to_timelock_executing_action);
    }

    function test_mainnet_rc_upgrade() public {
        run();  

        vm.warp(dayToQueueAction);

        // Queue Transaction
        vm.prank(operationsMultisig);
        (bool success, ) = address(timelock).call(calldata_to_timelock_queuing_action);
        require(success, "Timelock queueTransaction failed");

        // Fast forwart to after ETA
        vm.warp(timelockEta + 1);
        vm.prank(operationsMultisig);
        (success, ) = address(timelock).call(calldata_to_timelock_executing_action);
        require(success, "Timelock executeTransaction failed");

        // Assert owner
        assertEq(address(rewardsCoordinator.owner()), address(operationsMultisig), "RewardsCoordinator owner is not OperationsMultisig");

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized();
        _verifyInitializationParams();
    }

    function _sanityCheckImplementations(RewardsCoordinator oldRc, RewardsCoordinator newRc) internal {
        // Verify configs between both rewardsCoordinatorImplementations
        assertEq(
            address(oldRc.delegationManager()),
            address(newRc.delegationManager()),
            "DM mismatch"
        );
        assertEq(
            address(oldRc.strategyManager()),
            address(newRc.strategyManager()),
            "SM mismatch"
        );
        assertEq(
            oldRc.CALCULATION_INTERVAL_SECONDS(),
            newRc.CALCULATION_INTERVAL_SECONDS(),
            "CALCULATION_INTERVAL_SECONDS mismatch"
        );
        assertEq(
            oldRc.MAX_REWARDS_DURATION(),
            newRc.MAX_REWARDS_DURATION(),
            "MAX_REWARDS_DURATION mismatch"
        );
        assertEq(
            oldRc.MAX_RETROACTIVE_LENGTH(),
            newRc.MAX_RETROACTIVE_LENGTH(),
            "MAX_RETROACTIVE_LENGTH mismatch"
        );
        assertEq(
            oldRc.MAX_FUTURE_LENGTH(),
            newRc.MAX_FUTURE_LENGTH(),
            "MAX_FUTURE_LENGTH mismatch"
        );
        assertEq(
            oldRc.GENESIS_REWARDS_TIMESTAMP(),
            newRc.GENESIS_REWARDS_TIMESTAMP(),
            "GENESIS_REWARDS_TIMESTAMP mismatch"
        );
    }

    function test_set_reward_for_all_submitter(address hopper) public {
        test_mainnet_rc_upgrade();
    
        // Set reward for all submitters
        vm.prank(operationsMultisig);
        rewardsCoordinator.setRewardsForAllSubmitter(hopper, true);
        
        assertTrue(rewardsCoordinator.isRewardsForAllSubmitter(hopper), "Hopper not set for all submitters");
    }
}