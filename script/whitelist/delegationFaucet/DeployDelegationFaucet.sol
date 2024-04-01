// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/IStrategyManager.sol";
import "src/contracts/strategies/StrategyBase.sol";
import "src/contracts/permissions/PauserRegistry.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "./DelegationFaucet.sol";

import "forge-std/Script.sol";
import "forge-std/StdJson.sol";

/**
 * @notice Deploys the following contracts:
 * - ERC20 Dummy token for testing
 * - StrategyBase to be added to the StrategyManager whitelist
 * - DelegationFaucet contract
 */
contract DeployDelegationFaucet is Script, DSTest {
    // EigenLayer contracts
    ProxyAdmin public eigenLayerProxyAdmin;
    PauserRegistry public eigenLayerPauserReg;
    IDelegationManager public delegation;
    IStrategyManager public strategyManager;

    DelegationFaucet public delegationFaucet;

    // M2 testing/mock contracts
    ERC20PresetMinterPauser public stakeToken;
    StrategyBase public stakeTokenStrat;
    StrategyBase public baseStrategyImplementation;

    address eigenLayerProxyAdminAddress;
    address eigenLayerPauserRegAddress;
    address delegationAddress;
    address strategyManagerAddress;
    address operationsMultisig;
    address executorMultisig;

    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    function run() external {
        string memory goerliDeploymentConfig = vm.readFile("script/output/M1_deployment_goerli_2023_3_23.json");
        _setAddresses(goerliDeploymentConfig);

        vm.startBroadcast();
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

        // Needs to be strategyManager.strategyWhitelister() to add STK strategy
        // IStrategy[] memory _strategy = new IStrategy[](1);
        // _strategy[0] = stakeTokenStrat;
        // strategyManager.addStrategiesToDepositWhitelist(_strategy);

        // Deploy DelegationFaucet, grant it admin/mint/pauser roles, etc.
        delegationFaucet = new DelegationFaucet(strategyManager, delegation, stakeToken, stakeTokenStrat);
        stakeToken.grantRole(MINTER_ROLE, address(delegationFaucet));
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
