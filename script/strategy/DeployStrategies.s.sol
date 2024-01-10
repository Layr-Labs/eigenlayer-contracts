// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "../../src/contracts/strategies/StrategyBaseTVLLimits.sol";

import "../utils/ExistingDeploymentParser.sol";
import "../utils/DateTime.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

contract DeployStrategies is ExistingDeploymentParser {
    Vm cheats = Vm(HEVM_ADDRESS);

    // struct used to encode token info in config file
    struct StrategyConfig {
        uint256 maxDeposits;
        uint256 maxPerDeposit;
        address tokenAddress;
        string tokenSymbol;
    }

    // strategies deployed
    StrategyBaseTVLLimits[] public deployedStrategyArray;

    string public configPath = string(bytes("script/strategy/configs/config.json"));

    function run() external {
        _parseDeployedContracts();

        // READ JSON CONFIG/DEPLOYMENT FILE
        string memory config_data = vm.readFile(configPath);

        // tokens to deploy strategies for
        StrategyConfig[] memory strategyConfigs;
        bytes memory strategyConfigsRaw = stdJson.parseRaw(config_data, ".strategies");
        strategyConfigs = abi.decode(strategyConfigsRaw, (StrategyConfig[]));

        cheats.startBroadcast();

        // create upgradeable proxies that each point to the implementation and initialize them
        for (uint256 i = 0; i < strategyConfigs.length; ++i) {
            deployedStrategyArray.push(
                StrategyBaseTVLLimits(
                    address(
                        new TransparentUpgradeableProxy(
                            address(baseStrategyImplementation),
                            address(eigenLayerProxyAdmin),
                            abi.encodeWithSelector(
                                StrategyBaseTVLLimits.initialize.selector,
                                strategyConfigs[i].maxPerDeposit,
                                strategyConfigs[i].maxDeposits,
                                IERC20(strategyConfigs[i].tokenAddress),
                                eigenLayerPauserReg
                            )
                        )
                    )
                )
            );
        }

        cheats.stopBroadcast();

        _verifyImplementations();
        _verifyPauserInitializations();
        _verifyInitializationParams();

        // write the deployment info to a file

        string memory deployed_strategies = "strategies";
        for (uint256 i = 0; i < strategyConfigs.length; ++i) {
            vm.serializeAddress(deployed_strategies, strategyConfigs[i].tokenSymbol, address(deployedStrategyArray[i]));
        }
        string memory deployed_strategies_output = vm.serializeAddress(
            deployed_strategies,
            strategyConfigs[strategyConfigs.length - 1].tokenSymbol,
            address(deployedStrategyArray[strategyConfigs.length - 1])
        );

        string memory parent_object = "deployment";

        string memory chain_info = "chainInfo";
        vm.serializeUint(chain_info, "deploymentBlock", block.number);
        string memory chain_info_output = vm.serializeUint(chain_info, "chainId", block.chainid);

        vm.serializeString(parent_object, chain_info, chain_info_output);
        parent_object = vm.serializeString(parent_object, deployed_strategies, deployed_strategies_output);

        (uint256 year, uint256 month, uint256 day) = DateTime.timestampToDate(block.timestamp);
        string memory date = string(abi.encodePacked(vm.toString(year), "_", vm.toString(month), "_", vm.toString(day)));
        string memory output_file = string(abi.encodePacked("script/output/strategy/strategy_deployment_", chainName, "_", date, ".json"));

        vm.writeJson(parent_object, output_file);
    }

    function _verifyImplementations() internal view {
         for (uint256 i = 0; i < deployedStrategyArray.length; ++i) {
            require(
                eigenLayerProxyAdmin.getProxyImplementation(
                    TransparentUpgradeableProxy(payable(address(deployedStrategyArray[i])))
                ) == address(baseStrategyImplementation),
                "strategy: implementation set incorrectly"
            );
        }
    }

    function _verifyPauserInitializations() internal view {
        for (uint256 i = 0; i < deployedStrategyArray.length; ++i) {
            require(
                deployedStrategyArray[i].pauserRegistry() == eigenLayerPauserReg,
                "StrategyBaseTVLLimits: pauser registry not set correctly"
            );
            require(
                deployedStrategyArray[i].paused() == 0,
                "StrategyBaseTVLLimits: init paused status set incorrectly"
            );
        }
    }

    function _verifyInitializationParams() internal {
        string memory config_data = vm.readFile(configPath);
        for (uint i = 0; i < deployedStrategyArray.length; i++) {
            uint256 maxPerDeposit = stdJson.readUint(
                config_data,
                string.concat(".strategies[", vm.toString(i), "].max_per_deposit")
            );
            uint256 maxDeposits = stdJson.readUint(
                config_data,
                string.concat(".strategies[", vm.toString(i), "].max_deposits")
            );
            (uint256 setMaxPerDeposit, uint256 setMaxDeposits) = deployedStrategyArray[i].getTVLLimits();
            require(setMaxPerDeposit == maxPerDeposit, "setMaxPerDeposit not set correctly");
            require(setMaxDeposits == maxDeposits, "setMaxDeposits not set correctly");
        }
    }
}
