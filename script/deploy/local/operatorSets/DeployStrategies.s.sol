// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "../../../utils/ExistingDeploymentParser.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
contract DeployStrategies is ExistingDeploymentParser {

    function run() external {
        _parseDeployedContracts("script/output/devnet/M2_from_scratch_deployment_data.json");

        vm.startBroadcast();
        StrategyDeployer strategyDeployer = new StrategyDeployer(
            executorMultisig,
            address(baseStrategyImplementation)
        );
        uint256 batches = 1;
        uint256 batchSize = 100;

        IStrategy[] memory strategies = new IStrategy[](batchSize * batches);

        for (uint256 i = 0; i < batches; i++) {
            IStrategy[] memory strategiesJustDeployed = strategyDeployer.createManyStrategies(batchSize);
            for (uint256 j = 0; j < batchSize; j++) {
                strategies[i * batchSize + j] = strategiesJustDeployed[j];
            }
            strategyManager.addStrategiesToDepositWhitelist(strategiesJustDeployed);
        }

        vm.stopBroadcast();

        address[] memory strategyAddresses;
        assembly {
            strategyAddresses := strategies
        }
        string memory deployed_strategies = vm.serializeAddress("", "strategies", strategyAddresses);
        
        vm.writeJson(deployed_strategies, "script/output/devnet/deployed_strategies.json");
    }
}


contract StrategyDeployer {
    address immutable beneficiary;
    address immutable baseStrategyImplementation;

    constructor(address _beneficiary, address _baseStrategyImplementation) {
        beneficiary = _beneficiary;
        baseStrategyImplementation = _baseStrategyImplementation;
    }

    function createManyStrategies(uint256 numStrategies) external returns(IStrategy[] memory) {
        IStrategy[] memory strategies = new IStrategy[](numStrategies);
        for (uint256 i = 0; i < numStrategies; i++) {
            // create a strategy
            strategies[i] = 
                StrategyBaseTVLLimits(
                    address(
                        new TransparentUpgradeableProxy(
                            address(baseStrategyImplementation),
                            address(1),
                            abi.encodeWithSelector(
                                StrategyBaseTVLLimits.initialize.selector,
                                type(uint256).max,
                                type(uint256).max,
                                new ERC20PresetFixedSupply("Test", "TST", uint256(type(uint128).max), beneficiary),
                                address(1)
                            )
                        )
                    )
                );
        }
        return strategies;
    }
}