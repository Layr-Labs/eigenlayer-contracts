// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "../../src/contracts/core/StrategyManager.sol";
import "../../src/contracts/core/DelegationManager.sol";
import "../../src/contracts/core/AVSDirectory.sol";
import "../../src/contracts/core/RewardsCoordinator.sol";
import "../../src/contracts/core/AllocationManager.sol";
import "../../src/contracts/permissions/PermissionController.sol";
import "../../src/contracts/permissions/PauserRegistry.sol";

import "../../src/contracts/strategies/StrategyFactory.sol";
import "../../src/contracts/strategies/StrategyBase.sol";
import "../../src/contracts/strategies/StrategyBaseTVLLimits.sol";
import "../../src/contracts/strategies/EigenStrategy.sol";

import "../../src/contracts/pods/EigenPod.sol";
import "../../src/contracts/pods/EigenPodManager.sol";

import "../../src/contracts/interfaces/IBackingEigen.sol";
import "../../src/contracts/interfaces/IEigen.sol";

interface IDeploymentParser {
    // Core contracts
    function allocationManager() external view returns (AllocationManager);
    function avsDirectory() external view returns (AVSDirectory);
    function delegationManager() external view returns (DelegationManager);
    function rewardsCoordinator() external view returns (RewardsCoordinator);
    function strategyManager() external view returns (StrategyManager);
    function eigenPodManager() external view returns (EigenPodManager);
    function permissionController() external view returns (PermissionController);
    function pauserRegistry() external view returns (PauserRegistry);

    // Strategy contracts
    function strategyFactory() external view returns (StrategyFactory);
    function eigenStrategy() external view returns (EigenStrategy);
    function strategyBase() external view returns (StrategyBase);
    function deployedStrategyArray(
        uint256 index
    ) external view returns (StrategyBase);

    // Token contracts
    function EIGEN() external view returns (IEigen);
    function bEIGEN() external view returns (IBackingEigen);

    // Admin contracts
    function eigenLayerProxyAdmin() external view returns (ProxyAdmin);
    function tokenProxyAdmin() external view returns (ProxyAdmin);
    function eigenPodBeacon() external view returns (UpgradeableBeacon);

    // Implementation contracts
    function allocationManagerImplementation() external view returns (AllocationManager);
    function avsDirectoryImplementation() external view returns (AVSDirectory);
    function delegationManagerImplementation() external view returns (DelegationManager);
    function rewardsCoordinatorImplementation() external view returns (RewardsCoordinator);
    function strategyManagerImplementation() external view returns (StrategyManager);
    function eigenPodManagerImplementation() external view returns (EigenPodManager);
    function permissionControllerImplementation() external view returns (PermissionController);
    function strategyFactoryImplementation() external view returns (StrategyFactory);
    function eigenStrategyImplementation() external view returns (EigenStrategy);
    function baseStrategyImplementation() external view returns (StrategyBase);
    function eigenPodImplementation() external view returns (EigenPod);
    function EIGENImplementation() external view returns (IEigen);
    function bEIGENImplementation() external view returns (IBackingEigen);

    // Multisig addresses
    function executorMultisig() external view returns (address);
    function operationsMultisig() external view returns (address);
    function communityMultisig() external view returns (address);
    function pauserMultisig() external view returns (address);
    function timelock() external view returns (address);

    // Verification functions
    function verifyContractPointers() external view;
    function verifyImplementations() external view;
    function verifyContractsInitialized(
        bool isInitialDeployment
    ) external;
    function verifyInitializationParams() external view;

    // Initialization
    function initialize() external;
}
