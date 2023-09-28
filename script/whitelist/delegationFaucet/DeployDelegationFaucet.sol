// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/IStrategyManager.sol";
import "src/contracts/strategies/StrategyBase.sol";
import "src/contracts/permissions/PauserRegistry.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "./DelegationFaucet.sol";

import "forge-std/Script.sol";
import "forge-std/StdJson.sol";

contract DeployDelegationFaucet is Script, DSTest {
    // EigenLayer contracts
    ProxyAdmin public eigenLayerProxyAdmin;
    PauserRegistry public eigenLayerPauserReg;

    IDelegationManager public delegation;
    IStrategyManager public strategyManager;

    // M2 testing/mock contracts
    IERC20 public stakeToken;
    StrategyBase public stakeTokenStrat;
    StrategyBase public baseStrategyImplementation;

    address eigenLayerProxyAdminAddress;
    address eigenLayerPauserRegAddress;
    address delegationAddress;
    address strategyManagerAddress;
    address operationsMultisig;
    address executorMultisig;

    function run() external {
        string memory goerliDeploymentConfig = vm.readFile("script/output/M1_deployment_goerli_2023_3_23.json");
        _setAddresses(goerliDeploymentConfig);

        vm.startBroadcast(msg.sender);
        // Deploy ERC20 stakeToken
        stakeToken = new ERC20PresetMinterPauser("StakeToken", "STK");

        // Deploy StrategyBase for stakeToken & whitelist it
        baseStrategyImplementation = new StrategyBase(strategyManager);
        stakeTokenStrat = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(baseStrategyImplementation),
                    eigenLayerProxyAdminAddress,
                    abi.encodeWithSelector(StrategyBase.initialize.selector, stakeToken, eigenLayerPauserRegAddress)
                )
            )
        );

        // Deploy DelegationFaucet, grant it admin/mint/pauser roles, etc.

        vm.stopBroadcast();
    }

    function _setAddresses(string memory config) internal {
        eigenLayerProxyAdminAddress = stdJson.readAddress(config, ".addresses.eigenLayerProxyAdmin");
        eigenLayerPauserRegAddress = stdJson.readAddress(config, ".addresses.eigenLayerPauserReg");
        delegationAddress = stdJson.readAddress(config, ".addresses.delegation");
        strategyManagerAddress = stdJson.readAddress(config, ".addresses.strategyManager");
        operationsMultisig = stdJson.readAddress(config, ".parameters.operationsMultisig");
        executorMultisig = stdJson.readAddress(config, ".parameters.executorMultisig");
    }
}
