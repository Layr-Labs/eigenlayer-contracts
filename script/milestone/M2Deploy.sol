// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "../../src/contracts/interfaces/IETHPOSDeposit.sol";
import "../../src/contracts/interfaces/IBeaconChainOracle.sol";

import "../../src/contracts/core/StrategyManager.sol";
import "../../src/contracts/core/Slasher.sol";
import "../../src/contracts/core/DelegationManager.sol";

import "../../src/contracts/strategies/StrategyBaseTVLLimits.sol";

import "../../src/contracts/pods/EigenPod.sol";
import "../../src/contracts/pods/EigenPodManager.sol";
import "../../src/contracts/pods/DelayedWithdrawalRouter.sol";

import "../../src/contracts/permissions/PauserRegistry.sol";

import "../../src/test/mocks/EmptyContract.sol";
import "../../src/test/mocks/ETHDepositMock.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/upgrade/M2Deploy.s.sol:M2Deploy --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv
contract M2Deploy is Script, Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    string public m1DeploymentOutputPath = string(bytes("script/output/M1_deployment_goerli_2023_3_23.json"));

    IETHPOSDeposit public ethPOS;

    ISlasher public slasher;
    IDelegationManager public delegation;
    DelegationManager public delegationImplementation;
    IStrategyManager public strategyManager;
    StrategyManager public strategyManagerImplementation;
    IEigenPodManager public eigenPodManager;
    EigenPodManager public eigenPodManagerImplementation;
    IDelayedWithdrawalRouter public delayedWithdrawalRouter;
    IBeacon public eigenPodBeacon;
    EigenPod public eigenPodImplementation;

    function run() external {
        // read and log the chainID
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        if (chainId == 1) {
            m1DeploymentOutputPath = string(bytes("script/output/M1_deployment_mainnet_2023_6_9.json"));
        }

        // READ JSON DEPLOYMENT DATA
        string memory deployment_data = vm.readFile(m1DeploymentOutputPath);
        slasher = Slasher(stdJson.readAddress(deployment_data, ".addresses.slasher"));
        delegation = slasher.delegation();
        strategyManager = slasher.strategyManager();
        eigenPodManager = strategyManager.eigenPodManager();
        eigenPodBeacon = eigenPodManager.eigenPodBeacon();
        ethPOS = eigenPodManager.ethPOS();
        delayedWithdrawalRouter = EigenPod(payable(eigenPodBeacon.implementation())).delayedWithdrawalRouter();

        vm.startBroadcast();
        delegationImplementation = new DelegationManager(strategyManager, slasher, eigenPodManager);
        strategyManagerImplementation = new StrategyManager(delegation, eigenPodManager, slasher);
        eigenPodManagerImplementation = new EigenPodManager(
            ethPOS,
            eigenPodBeacon,
            strategyManager,
            slasher,
            delegation
        );
        eigenPodImplementation = new EigenPod({
            _ethPOS: ethPOS,
            _delayedWithdrawalRouter: delayedWithdrawalRouter,
            _eigenPodManager: eigenPodManager,
            _MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR: 31 gwei,
            _RESTAKED_BALANCE_OFFSET_GWEI: 0.5 gwei,
            _GENESIS_TIME: 1616508000
        });

        // write the output to a contract
        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(deployed_addresses, "slasher", address(slasher));
        vm.serializeAddress(deployed_addresses, "delegation", address(delegation));
        vm.serializeAddress(deployed_addresses, "strategyManager", address(strategyManager));
        vm.serializeAddress(deployed_addresses, "delayedWithdrawalRouter", address(delayedWithdrawalRouter));
        vm.serializeAddress(deployed_addresses, "eigenPodManager", address(eigenPodManager));
        vm.serializeAddress(deployed_addresses, "eigenPodBeacon", address(eigenPodBeacon));
        vm.serializeAddress(deployed_addresses, "ethPOS", address(ethPOS));

        vm.serializeAddress(deployed_addresses, "delegationImplementation", address(delegationImplementation));
        vm.serializeAddress(
            deployed_addresses,
            "strategyManagerImplementation",
            address(strategyManagerImplementation)
        );
        vm.serializeAddress(
            deployed_addresses,
            "eigenPodManagerImplementation",
            address(eigenPodManagerImplementation)
        );
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "eigenPodImplementation",
            address(eigenPodImplementation)
        );

        string memory chain_info = "chainInfo";
        vm.serializeUint(chain_info, "deploymentBlock", block.number);
        string memory chain_info_output = vm.serializeUint(chain_info, "chainId", chainId);

        vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);
        // serialize all the data
        string memory finalJson = vm.serializeString(parent_object, chain_info, chain_info_output);
        vm.writeJson(finalJson, "script/output/M2_deployment_data.json");
    }
}
