// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../utils/ExistingDeploymentParser.sol";

/**
 * @notice Script used for the first deployment of EigenLayer core contracts to Holesky
 * forge script script/deploy/holesky/M2_Deploy_Fixed_EigenPod.s.sol --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 * forge script script/deploy/holesky/M2_Deploy_Fixed_EigenPod.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv
 *
 */
contract M2_Deploy_Fixed_EigenPod is ExistingDeploymentParser {
    function run() external virtual {
        _parseDeployedContracts("script/output/holesky/M2_deploy_from_scratch.holesky.config.json");
        _parseInitialDeploymentParams("script/configs/holesky/M2_deploy_fixed_EigenPod.holesky.config.json");

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        emit log_named_address("Deployer Address", msg.sender);

        _deployFixedEigenPod();
        // _deployAdditionalStrategies();

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        // Sanity Checks
        // _verifyContractPointers();
        // _verifyContractsInitialized({isInitialDeployment: true});
        // _verifyInitializationParams();

        // logAndOutputContractAddresses("script/output/holesky/M2_deploy_fixed_EigenPod.holesky.config.json");
    }

    /**
     * @notice Initially deployed with WEI value for EIGENPOD_MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR
     * This deploys a new eigenPod and upgrades the upgradeable beacon implementation
     */
    function _deployFixedEigenPod() internal {
        require(
            EIGENPOD_MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR == 32000000000,
            "EIGENPOD_MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR not set properly"
        );

        eigenPodManager = EigenPodManager(0xB8d8952f572e67B11e43bC21250967772fa883Ff);
        delayedWithdrawalRouter = DelayedWithdrawalRouter(0xC4BC46a87A67a531eCF7f74338E1FA79533334Fa);

        // Deploy EigenPod Contracts
        eigenPodImplementation = new EigenPod(
            IETHPOSDeposit(ETHPOSDepositAddress),
            delayedWithdrawalRouter,
            eigenPodManager,
            EIGENPOD_MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            EIGENPOD_GENESIS_TIME
        );

        emit log_named_address("EigenPod Implementation", address(eigenPodImplementation));
    }

    /// @notice Deploy additional strategies for testnetwork
    function _deployAdditionalStrategies() internal {
        for (uint256 i = 0; i < numStrategiesToDeploy; i++) {
            StrategyUnderlyingTokenConfig memory strategyConfig = strategiesToDeploy[i];

            // Deploy and upgrade strategy
            StrategyBaseTVLLimits strategy = StrategyBaseTVLLimits(
                address(
                    new TransparentUpgradeableProxy(
                        address(baseStrategyImplementation),
                        address(eigenLayerProxyAdmin),
                        abi.encodeWithSelector(
                            StrategyBaseTVLLimits.initialize.selector,
                            STRATEGY_MAX_PER_DEPOSIT,
                            STRATEGY_MAX_TOTAL_DEPOSITS,
                            IERC20(strategyConfig.tokenAddress),
                            eigenLayerPauserReg
                        )
                    )
                )
            );

            deployedStrategyArray.push(strategy);
        }
    }
}
