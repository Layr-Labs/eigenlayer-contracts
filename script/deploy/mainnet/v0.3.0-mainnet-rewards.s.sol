// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/utils/Create2.sol";
import "../../utils/ExistingDeploymentParser.sol";

/**
 * @notice Script used for the first deployment of EigenLayer core contracts to Holesky
 * FORK LOCAL
 * anvil --fork-url $RPC_MAINNET
 * forge script script/deploy/mainnet/v0.3.0-mainnet-rewards.s.sol:MainnetRewardsCoordinatorDeploy --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 *
 * MAINNET
 * forge script script/deploy/mainnet/v0.3.0-mainnet-rewards.s.sol:MainnetRewardsCoordinatorDeploy --rpc-url $RPC_MAINNET --private-key $PRIVATE_KEY --verify --broadcast -vvvv
 *
 */
contract MainnetRewardsCoordinatorDeploy is ExistingDeploymentParser {
    function run() external virtual {
        _parseInitialDeploymentParams(
            "script/configs/mainnet/v0.3.0-mainnet-rewards.config.json"
        );
        _parseDeployedContracts(
            "script/configs/mainnet/v0.3.0-eigenlayer-addresses.config.json"
        );

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        emit log_named_address("Deployer Address", msg.sender);

        _deployRewardsCoordinator();

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // Sanity Checks
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized();
        _verifyInitializationParams();

        logAndOutputContractAddresses("script/output/mainnet/v0.3.0-mainnet-rewards.output.json");
    }

    /**
     * @notice Deploy RewardsCoordinator for Holesky
     */
    function _deployRewardsCoordinator() internal {


        // Deploy RewardsCoordinator proxy and implementation
        rewardsCoordinatorImplementation = new RewardsCoordinator(
            delegationManager,
            strategyManager,
            REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS,
            REWARDS_COORDINATOR_MAX_REWARDS_DURATION,
            REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH,
            REWARDS_COORDINATOR_MAX_FUTURE_LENGTH,
            REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP
        );
        rewardsCoordinator = RewardsCoordinator(
            address(
                new TransparentUpgradeableProxy(
                    address(rewardsCoordinatorImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        RewardsCoordinator.initialize.selector,
                        executorMultisig,
                        eigenLayerPauserReg,
                        REWARDS_COORDINATOR_INIT_PAUSED_STATUS,
                        REWARDS_COORDINATOR_UPDATER,
                        REWARDS_COORDINATOR_ACTIVATION_DELAY,
                        REWARDS_COORDINATOR_GLOBAL_OPERATOR_COMMISSION_BIPS
                    )
                )
            )
        );
    }

    /**
     * @notice Deploy RewardsCoordinator Implementation for Holesky and upgrade the proxy
     */
    function _upgradeRewardsCoordinator() internal {
        // Deploy RewardsCoordinator proxy and implementation
        rewardsCoordinatorImplementation = new RewardsCoordinator(
            delegationManager,
            strategyManager,
            REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS,
            REWARDS_COORDINATOR_MAX_REWARDS_DURATION,
            REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH,
            REWARDS_COORDINATOR_MAX_FUTURE_LENGTH,
            REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP
        );

        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(rewardsCoordinator))),
            address(rewardsCoordinatorImplementation)
        );
    }
}
