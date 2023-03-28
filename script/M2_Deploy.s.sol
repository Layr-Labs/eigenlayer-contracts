// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "../src/contracts/core/StrategyManager.sol";
import "../src/contracts/core/Slasher.sol";
import "../src/contracts/core/DelegationManager.sol";

import "../src/contracts/strategies/StrategyBase.sol";

import "../src/contracts/pods/EigenPod.sol";
import "../src/contracts/pods/EigenPodManager.sol";
import "../src/contracts/pods/DelayedWithdrawalRouter.sol";
import "../src/contracts/pods/BeaconChainOracle.sol";

import "../src/contracts/permissions/PauserRegistry.sol";

import "../src/test/mocks/EmptyContract.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";



// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/M2_Deploy.s.sol:Deployer_M2 --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv
contract Deployer_M2 is Script, Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    string public deployConfigPath = string(bytes("script/M2_deploy.config.json"));

    // EigenLayer Contracts
    ProxyAdmin public eigenLayerProxyAdmin;
    PauserRegistry public eigenLayerPauserReg;
    Slasher public slasher;
    Slasher public slasherImplementation;
    DelegationManager public delegation;
    DelegationManager public delegationImplementation;
    StrategyManager public strategyManager;
    StrategyManager public strategyManagerImplementation;
    EigenPodManager public eigenPodManager;
    EigenPodManager public eigenPodManagerImplementation;
    DelayedWithdrawalRouter public delayedWithdrawalRouter;
    DelayedWithdrawalRouter public delayedWithdrawalRouterImplementation;
    UpgradeableBeacon public eigenPodBeacon;
    EigenPod public eigenPodImplementation;
    StrategyBase public baseStrategyImplementation;

    EmptyContract public emptyContract;

    address communityMultisig;
    address teamMultisig;

    // strategies deployed
    StrategyBase[] public deployedStrategyArray;


    BeaconChainOracle public beaconChainOracle;

    function run() external {
        // read and log the chainID
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        // READ JSON CONFIG DATA
        string memory config_data = vm.readFile(deployConfigPath);

        // check that the chainID matches the one in the config
        uint256 configChainId = stdJson.readUint(config_data, ".chainInfo.chainId");
        require(configChainId == chainId, "You are on the wrong chain for this config");

        // read all of the deployed addresses
        communityMultisig = stdJson.readAddress(config_data, ".multisig_addresses.communityMultisig");
        teamMultisig = stdJson.readAddress(config_data, ".multisig_addresses.teamMultisig");
        
        eigenLayerProxyAdmin = ProxyAdmin(stdJson.readAddress(config_data, ".addresses.eigenLayerProxyAdmin"));
        eigenLayerPauserReg = PauserRegistry(stdJson.readAddress(config_data, ".addresses.eigenLayerPauserReg"));
        slasher = Slasher(stdJson.readAddress(config_data, ".addresses.slasher"));
        slasherImplementation = Slasher(stdJson.readAddress(config_data, ".addresses.slasherImplementation"));
        delegation = DelegationManager(stdJson.readAddress(config_data, ".addresses.delegation"));
        delegationImplementation = DelegationManager(stdJson.readAddress(config_data, ".addresses.delegationImplementation"));
        strategyManager = StrategyManager(stdJson.readAddress(config_data, ".addresses.strategyManager"));
        strategyManagerImplementation = StrategyManager(stdJson.readAddress(config_data, ".addresses.strategyManagerImplementation"));
        eigenPodManager = EigenPodManager(stdJson.readAddress(config_data, ".addresses.eigenPodManager"));
        eigenPodManagerImplementation = EigenPodManager(stdJson.readAddress(config_data, ".addresses.eigenPodManagerImplementation"));
        delayedWithdrawalRouter = DelayedWithdrawalRouter(stdJson.readAddress(config_data, ".addresses.delayedWithdrawalRouter"));
        delayedWithdrawalRouterImplementation = 
            DelayedWithdrawalRouter(stdJson.readAddress(config_data, ".addresses.delayedWithdrawalRouterImplementation"));
        eigenPodBeacon = UpgradeableBeacon(stdJson.readAddress(config_data, ".addresses.eigenPodBeacon"));
        eigenPodImplementation = EigenPod(stdJson.readAddress(config_data, ".addresses.eigenPodImplementation"));
        baseStrategyImplementation = StrategyBase(stdJson.readAddress(config_data, ".addresses.baseStrategyImplementation"));
        emptyContract = EmptyContract(stdJson.readAddress(config_data, ".addresses.emptyContract"));

        // load strategy list
        bytes memory strategyListRaw = stdJson.parseRaw(config_data, ".addresses.strategies");
        address[] memory strategyList = abi.decode(strategyListRaw, (address[]));
        for (uint256 i = 0; i < strategyList.length; ++i) {
            deployedStrategyArray.push(StrategyBase(strategyList[i]));
        }

        address oracleInitialOwner = communityMultisig;
        uint256 initialThreshold = 2;
        bytes memory oracleSignerListRaw = stdJson.parseRaw(config_data, ".oracleInitialization.signers");
        address[] memory initialOracleSigners = abi.decode(oracleSignerListRaw, (address[]));

        require(initialThreshold <= initialOracleSigners.length, "invalid initialThreshold");

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        beaconChainOracle = new BeaconChainOracle(oracleInitialOwner, initialThreshold, initialOracleSigners);

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();
    }
}



    

