// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "../../src/contracts/core/StrategyManager.sol";
import "../../src/contracts/core/Slasher.sol";
import "../../src/contracts/core/DelegationManager.sol";
import "../../src/contracts/core/AVSDirectory.sol";

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
    AVSDirectory public avsDirectory;
    AVSDirectory public avsDirectoryImplementation;
    DelegationManager public delegationManager;
    DelegationManager public delegationManagerImplementation;
    StrategyManager public strategyManager;
    StrategyManager public strategyManagerImplementation;
    EigenPodManager public eigenPodManager;
    EigenPodManager public eigenPodManagerImplementation;
    DelayedWithdrawalRouter public delayedWithdrawalRouter;
    DelayedWithdrawalRouter public delayedWithdrawalRouterImplementation;
    UpgradeableBeacon public eigenPodBeacon;
    EigenPod pod;
    EigenPod public eigenPodImplementation;
    StrategyBase public baseStrategyImplementation;

    EmptyContract public emptyContract;

    address executorMultisig;
    address operationsMultisig;

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
        executorMultisig = stdJson.readAddress(existingDeploymentData, ".parameters.executorMultisig");
        operationsMultisig = stdJson.readAddress(existingDeploymentData, ".parameters.operationsMultisig");
        
        eigenLayerProxyAdmin = ProxyAdmin(stdJson.readAddress(existingDeploymentData, ".addresses.eigenLayerProxyAdmin"));
        eigenLayerPauserReg = PauserRegistry(stdJson.readAddress(existingDeploymentData, ".addresses.eigenLayerPauserReg"));
        slasher = Slasher(stdJson.readAddress(existingDeploymentData, ".addresses.slasher"));
        slasherImplementation = Slasher(stdJson.readAddress(existingDeploymentData, ".addresses.slasherImplementation"));
        delegationManager = DelegationManager(stdJson.readAddress(existingDeploymentData, ".addresses.delegation"));
        delegationManagerImplementation = DelegationManager(stdJson.readAddress(existingDeploymentData, ".addresses.delegationImplementation"));
        avsDirectory = AVSDirectory(stdJson.readAddress(existingDeploymentData, ".addresses.avsDirectory"));
        avsDirectoryImplementation = AVSDirectory(stdJson.readAddress(existingDeploymentData, ".addresses.avsDirectoryImplementation"));
        strategyManager = StrategyManager(stdJson.readAddress(existingDeploymentData, ".addresses.strategyManager"));
        strategyManagerImplementation = StrategyManager(stdJson.readAddress(existingDeploymentData, ".addresses.strategyManagerImplementation"));
        eigenPodManager = EigenPodManager(stdJson.readAddress(existingDeploymentData, ".addresses.eigenPodManager"));
        eigenPodManagerImplementation = EigenPodManager(stdJson.readAddress(existingDeploymentData, ".addresses.eigenPodManagerImplementation"));
        delayedWithdrawalRouter = DelayedWithdrawalRouter(stdJson.readAddress(existingDeploymentData, ".addresses.delayedWithdrawalRouter"));
        delayedWithdrawalRouterImplementation = 
            DelayedWithdrawalRouter(stdJson.readAddress(existingDeploymentData, ".addresses.delayedWithdrawalRouterImplementation"));
        eigenPodBeacon = UpgradeableBeacon(stdJson.readAddress(existingDeploymentData, ".addresses.eigenPodBeacon"));
        eigenPodImplementation = EigenPod(payable(stdJson.readAddress(existingDeploymentData, ".addresses.eigenPodImplementation")));
        baseStrategyImplementation = StrategyBase(stdJson.readAddress(existingDeploymentData, ".addresses.baseStrategyImplementation"));
        emptyContract = EmptyContract(stdJson.readAddress(existingDeploymentData, ".addresses.emptyContract"));

        /*
        commented out -- needs JSON formatting of the form:
        strategies": [
      {"WETH": "0x7CA911E83dabf90C90dD3De5411a10F1A6112184"},
      {"rETH": "0x879944A8cB437a5f8061361f82A6d4EED59070b5"},
      {"tsETH": "0xcFA9da720682bC4BCb55116675f16F503093ba13"},
      {"wstETH": "0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff"}]
        // load strategy list
        bytes memory strategyListRaw = stdJson.parseRaw(existingDeploymentData, ".addresses.strategies");
        address[] memory strategyList = abi.decode(strategyListRaw, (address[]));
        for (uint256 i = 0; i < strategyList.length; ++i) {
            deployedStrategyArray.push(StrategyBase(strategyList[i]));
        }
        */
    }

    /**
     * @notice Verify that constructor parameters are set correctly from implementation upgrades
     * Should be called after calling _parseDeployedContracts and performing upgrades. Note these are just basic upgrade checks,
     * and more detailed storage checks using script/validateStorage.ts and manual storage checks should also be performed.
     */
    function _verifyContractPointers() internal {
        // verify DelegationManager contracts
        require(delegationManager.strategyManager() == strategyManager, "DelegationManager.strategyManager not set");
        require(delegationManager.slasher() == slasher, "DelegationManager.slasher not set");
        require(delegationManager.eigenPodManager() == eigenPodManager, "DelegationManager.eigenPodManager not set");

        // verify StrategyManager contracts
        require(strategyManager.delegation() == delegationManager, "StrategyManager.delegationManager not set");
        require(strategyManager.eigenPodManager() == eigenPodManager, "StrategyManager.eigenPodManager not set");
        require(strategyManager.slasher() == slasher, "StrategyManager.slasher not set");

        // verify EigenPodManager contracts
        require(eigenPodManager.eigenPodBeacon() == eigenPodBeacon, "EigenPodManager.eigenPodBeacon not set");
        require(eigenPodManager.delegationManager() == delegationManager, "EigenPodManager.delegationManager not set");
        require(eigenPodManager.strategyManager() == strategyManager, "EigenPodManager.strategyManager not set");
        require(eigenPodManager.slasher() == slasher, "EigenPodManager.slasher not set");

        // verify delayedWithdrawalRouter contracts
        require(delayedWithdrawalRouter.eigenPodManager() == eigenPodManager, "DelayedWithdrawalRouter.eigenPodManager not set");

        // verify eigenPod implementation contract
        require(eigenPodImplementation.eigenPodManager() == eigenPodManager, "EigenPod.eigenPodManager not set");
        require(eigenPodImplementation.delayedWithdrawalRouter() == delayedWithdrawalRouter, "EigenPod.delayedWithdrawalRouter not set");

        // verify AVSDirectory contracts
        require(avsDirectory.delegation() == delegationManager, "AVSDirectory.delegationManager not set");
    }
}
