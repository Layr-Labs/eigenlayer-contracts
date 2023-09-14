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
import "../../src/contracts/middleware/BLSPublicKeyCompendium.sol";

import "../../src/test/mocks/EmptyContract.sol";
import "../../src/test/mocks/ETHDepositMock.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/upgrade/GoerliM2Deployment.s.sol:GoerliM2Deployment --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv
contract GoerliM2Deployment is Script, Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    string public deploymentOutputPath = string(bytes("script/output/M1_deployment_goerli_2023_3_23.json"));

    address executorMultisig;
    address operationsMultisig;
    address pauserMultisig;

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

        // READ JSON DEPLOYMENT DATA
        string memory deployment_data = vm.readFile(deploymentOutputPath);
        slasher = Slasher(stdJson.readAddress(deployment_data, ".addresses.slasher"));
        delegation = slasher.delegation();
        strategyManager = slasher.strategyManager();
        eigenPodManager = strategyManager.eigenPodManager();
        delayedWithdrawalRouter = DelayedWithdrawalRouter(stdJson.readAddress(deployment_data, ".addresses.delayedWithdrawalRouter"));
        eigenPodBeacon = eigenPodManager.eigenPodBeacon();
        ethPOS = eigenPodManager.ethPOS();


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
        eigenPodImplementation = new EigenPod(
            ethPOS,
            delayedWithdrawalRouter,
            eigenPodManager,
            31 gwei,
            0.5 gwei
        );

        // write the output to a contract
        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(deployed_addresses, "delegationImplementation", address(delegationImplementation));
        vm.serializeAddress(deployed_addresses, "strategyManagerImplementation", address(strategyManagerImplementation));
        vm.serializeAddress(deployed_addresses, "eigenPodManagerImplementation", address(eigenPodManagerImplementation));
        vm.serializeAddress(deployed_addresses, "eigenPodImplementation", address(eigenPodImplementation));

        string memory chain_info = "chainInfo";
        vm.serializeUint(chain_info, "deploymentBlock", block.number);
        string memory chain_info_output = vm.serializeUint(chain_info, "chainId", chainId);

        // serialize all the data
        vm.serializeString(parent_object, chain_info, chain_info_output);
        vm.writeJson(parent_object, "script/output/M2_deployment_data.json");
    }
}
