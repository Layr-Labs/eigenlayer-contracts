// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "../../src/contracts/core/StrategyManager.sol";
import "../../src/contracts/core/Slasher.sol";
import "../../src/contracts/core/DelegationManager.sol";

import "../../src/contracts/strategies/StrategyBase.sol";

import "../../src/contracts/pods/EigenPod.sol";
import "../../src/contracts/pods/EigenPodManager.sol";
import "../../src/contracts/pods/DelayedWithdrawalRouter.sol";

import "../../src/contracts/permissions/PauserRegistry.sol";

import "../../src/test/mocks/EmptyContract.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

contract ExistingDeploymentParser is Script, Test {

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

    function _parseDeployedContracts(string memory existingDeploymentInfoPath) internal {
        // read and log the chainID
        uint256 currentChainId = block.chainid;
        emit log_named_uint("You are parsing on ChainID", currentChainId);

        // READ JSON CONFIG DATA
        string memory existingDeploymentData = vm.readFile(existingDeploymentInfoPath);

        // check that the chainID matches the one in the config
        uint256 configChainId = stdJson.readUint(existingDeploymentData, ".chainInfo.chainId");
        require(configChainId == currentChainId, "You are on the wrong chain for this config");

        // read all of the deployed addresses
        communityMultisig = stdJson.readAddress(existingDeploymentData, ".multisig_addresses.communityMultisig");
        teamMultisig = stdJson.readAddress(existingDeploymentData, ".multisig_addresses.teamMultisig");
        
        eigenLayerProxyAdmin = ProxyAdmin(stdJson.readAddress(existingDeploymentData, ".addresses.eigenLayerProxyAdmin"));
        eigenLayerPauserReg = PauserRegistry(stdJson.readAddress(existingDeploymentData, ".addresses.eigenLayerPauserReg"));
        slasher = Slasher(stdJson.readAddress(existingDeploymentData, ".addresses.slasher"));
        slasherImplementation = Slasher(stdJson.readAddress(existingDeploymentData, ".addresses.slasherImplementation"));
        delegation = DelegationManager(stdJson.readAddress(existingDeploymentData, ".addresses.delegation"));
        delegationImplementation = DelegationManager(stdJson.readAddress(existingDeploymentData, ".addresses.delegationImplementation"));
        strategyManager = StrategyManager(stdJson.readAddress(existingDeploymentData, ".addresses.strategyManager"));
        strategyManagerImplementation = StrategyManager(stdJson.readAddress(existingDeploymentData, ".addresses.strategyManagerImplementation"));
        eigenPodManager = EigenPodManager(stdJson.readAddress(existingDeploymentData, ".addresses.eigenPodManager"));
        eigenPodManagerImplementation = EigenPodManager(stdJson.readAddress(existingDeploymentData, ".addresses.eigenPodManagerImplementation"));
        delayedWithdrawalRouter = DelayedWithdrawalRouter(stdJson.readAddress(existingDeploymentData, ".addresses.delayedWithdrawalRouter"));
        delayedWithdrawalRouterImplementation = 
            DelayedWithdrawalRouter(stdJson.readAddress(existingDeploymentData, ".addresses.delayedWithdrawalRouterImplementation"));
        eigenPodBeacon = UpgradeableBeacon(stdJson.readAddress(existingDeploymentData, ".addresses.eigenPodBeacon"));
        eigenPodImplementation = EigenPod(stdJson.readAddress(existingDeploymentData, ".addresses.eigenPodImplementation"));
        baseStrategyImplementation = StrategyBase(stdJson.readAddress(existingDeploymentData, ".addresses.baseStrategyImplementation"));
        emptyContract = EmptyContract(stdJson.readAddress(existingDeploymentData, ".addresses.emptyContract"));

        // load strategy list
        bytes memory strategyListRaw = stdJson.parseRaw(existingDeploymentData, ".addresses.strategies");
        address[] memory strategyList = abi.decode(strategyListRaw, (address[]));
        for (uint256 i = 0; i < strategyList.length; ++i) {
            deployedStrategyArray.push(StrategyBase(strategyList[i]));
        }
    }
}
