// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./Deploy_Test_RewardsCoordinator.s.sol";
import "script/utils/TimelockEncoding.sol";

/**
 * @notice Script used for the first deployment of EigenLayer core contracts to Holesky
 * anvil --fork-url $RPC_HOLESKY
 *
 * Holesky testnet: Deploy/Upgrade RewardsCoordinator
 * forge script script/deploy/holesky/v042-upgrade_testnet_rewardsCoordinator.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv --verify --etherscan-api-key $ETHERSCAN_API_KEY
 *
 */
contract Upgrade_Testnet_RewardsCoordinator is Deploy_Test_RewardsCoordinator, TimelockEncoding {
    function run() external virtual override {
        _parseInitialDeploymentParams("script/configs/holesky/eigenlayer_testnet.config.json");
        _parseDeployedContracts("script/configs/holesky/eigenlayer_addresses_testnet.config.json");

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

        // Create Upgrade Tx via Community Multisig
        bytes memory calldata_to_proxy_admin = abi.encodeWithSelector(
            ProxyAdmin.upgrade.selector,
            TransparentUpgradeableProxy(payable(address(rewardsCoordinator))),
            rewardsCoordinatorImplementation
        );
        
        bytes memory final_calldata_to_executor_multisig = encodeForExecutor(
            communityMultisig, //from
            address(eigenLayerProxyAdmin), //to 
            0, // value
            calldata_to_proxy_admin, // data
            ISafe.Operation.Call // operation
        );

        // Simulate Transaction
        vm.prank(communityMultisig);
        (bool success, ) = address(executorMultisig).call(final_calldata_to_executor_multisig);
        require(success, "Transaction failed");

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

    function setRewardForAllSubmitter() external {
        address hopper;

        require(hopper != address(0), "Hopper address is not set");

        // Set reward for all submitters
        vm.startBroadcast();
        rewardsCoordinator.setRewardsForAllSubmitter(hopper, true);
        vm.stopBroadcast();

        require(rewardsCoordinator.isRewardsForAllSubmitter(hopper), "Hopper is not set as rewards for all submitter");
    }
}