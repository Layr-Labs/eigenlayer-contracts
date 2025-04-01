// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Script.sol";
import "src/test/utils/Logger.t.sol";
import "./IDeploymentParser.sol";
import "../releases/Env.sol";

contract ZeusDeploymentParser is IDeploymentParser, Script, Logger {
    using stdJson for string;

    function NAME() public view virtual override returns (string memory) {
        return "ZeusDeploymentParser";
    }

    // Core contracts
    function allocationManager() external view returns (AllocationManager) {
        return Env.allocationManager(Env.proxy);
    }

    function avsDirectory() external view returns (AVSDirectory) {
        return Env.avsDirectory(Env.proxy);
    }

    function delegationManager() external view returns (DelegationManager) {
        return Env.delegationManager(Env.proxy);
    }

    function rewardsCoordinator() external view returns (RewardsCoordinator) {
        return Env.rewardsCoordinator(Env.proxy);
    }

    function strategyManager() external view returns (StrategyManager) {
        return Env.strategyManager(Env.proxy);
    }

    function eigenPodManager() external view returns (EigenPodManager) {
        return Env.eigenPodManager(Env.proxy);
    }

    function permissionController() external view returns (PermissionController) {
        return Env.permissionController(Env.proxy);
    }

    function pauserRegistry() external view returns (PauserRegistry) {
        return Env.pauserRegistry(Env.impl);
    }

    // Strategy contracts
    function strategyFactory() external view returns (StrategyFactory) {
        return Env.strategyFactory(Env.proxy);
    }

    function eigenStrategy() external view returns (EigenStrategy) {
        return Env.eigenStrategy(Env.proxy);
    }

    function strategyBase() external view returns (StrategyBase) {
        return Env.strategyBase(Env.impl);
    }

    function deployedStrategyArray(uint256 index) external view returns (StrategyBase) {
        return Env.strategyBaseTVLLimits(Env.instance, index);
    }

    // Token contracts
    function EIGEN() external view returns (IEigen) {
        return Env.eigen(Env.proxy);
    }

    function bEIGEN() external view returns (IBackingEigen) {
        return Env.beigen(Env.proxy);
    }

    // Admin contracts
    function eigenLayerProxyAdmin() external view returns (ProxyAdmin) {
        return ProxyAdmin(Env.proxyAdmin());
    }

    function tokenProxyAdmin() external view returns (ProxyAdmin) {
        return ProxyAdmin(Env.proxyAdmin());
    }

    function eigenPodBeacon() external view returns (UpgradeableBeacon) {
        return Env.eigenPod(Env.beacon);
    }

    // Implementation contracts
    function allocationManagerImplementation() external view returns (AllocationManager) {
        return Env.allocationManager(Env.impl);
    }

    function avsDirectoryImplementation() external view returns (AVSDirectory) {
        return Env.avsDirectory(Env.impl);
    }

    function delegationManagerImplementation() external view returns (DelegationManager) {
        return Env.delegationManager(Env.impl);
    }

    function rewardsCoordinatorImplementation() external view returns (RewardsCoordinator) {
        return Env.rewardsCoordinator(Env.impl);
    }

    function strategyManagerImplementation() external view returns (StrategyManager) {
        return Env.strategyManager(Env.impl);
    }

    function eigenPodManagerImplementation() external view returns (EigenPodManager) {
        return Env.eigenPodManager(Env.impl);
    }

    function permissionControllerImplementation() external view returns (PermissionController) {
        return Env.permissionController(Env.impl);
    }

    function strategyFactoryImplementation() external view returns (StrategyFactory) {
        return Env.strategyFactory(Env.impl);
    }

    function eigenStrategyImplementation() external view returns (EigenStrategy) {
        return Env.eigenStrategy(Env.impl);
    }

    function baseStrategyImplementation() external view returns (StrategyBase) {
        return Env.strategyBase(Env.impl);
    }

    function eigenPodImplementation() external view returns (EigenPod) {
        return Env.eigenPod(Env.impl);
    }

    function EIGENImplementation() external view returns (IEigen) {
        return Env.eigen(Env.impl);
    }

    function bEIGENImplementation() external view returns (IBackingEigen) {
        return Env.beigen(Env.impl);
    }

    // Multisig addresses
    function executorMultisig() external view returns (address) {
        return Env.executorMultisig();
    }

    function operationsMultisig() external view returns (address) {
        return Env.opsMultisig();
    }

    function communityMultisig() external view returns (address) {
        return Env.protocolCouncilMultisig();
    }

    function pauserMultisig() external view returns (address) {
        return Env.pauserMultisig();
    }

    function timelock() external view returns (address) {
        return address(Env.timelockController());
    }

    // Verification functions
    function verifyContractPointers() external view {
        // AVSDirectory
        assertTrue(
            this.avsDirectory().delegation() == this.delegationManager(),
            "avsDirectory: delegationManager address not set correctly"
        );
        // RewardsCoordinator
        assertTrue(
            this.rewardsCoordinator().delegationManager() == this.delegationManager(),
            "rewardsCoordinator: delegationManager address not set correctly"
        );
        assertTrue(
            this.rewardsCoordinator().strategyManager() == this.strategyManager(),
            "rewardsCoordinator: strategyManager address not set correctly"
        );
        // DelegationManager
        assertTrue(
            this.delegationManager().strategyManager() == this.strategyManager(),
            "delegationManager: strategyManager address not set correctly"
        );
        assertTrue(
            this.delegationManager().eigenPodManager() == this.eigenPodManager(),
            "delegationManager: eigenPodManager address not set correctly"
        );
        // StrategyManager
        assertTrue(
            this.strategyManager().delegation() == this.delegationManager(),
            "strategyManager: delegationManager address not set correctly"
        );
        // EPM
        assertTrue(
            address(this.eigenPodManager().ethPOS()) == address(Env.ethPOS()),
            "eigenPodManager: ethPOSDeposit contract address not set correctly"
        );
        assertTrue(
            this.eigenPodManager().eigenPodBeacon() == this.eigenPodBeacon(),
            "eigenPodManager: eigenPodBeacon contract address not set correctly"
        );
        assertTrue(
            this.eigenPodManager().delegationManager() == this.delegationManager(),
            "eigenPodManager: delegationManager contract address not set correctly"
        );
    }

    function verifyImplementations() external view {
        assertEq(
            this.eigenLayerProxyAdmin().getProxyImplementation(ITransparentUpgradeableProxy(payable(address(this.avsDirectory())))),
            address(this.avsDirectoryImplementation()),
            "avsDirectory: implementation set incorrectly"
        );
        assertEq(
            this.eigenLayerProxyAdmin().getProxyImplementation(
                ITransparentUpgradeableProxy(payable(address(this.rewardsCoordinator())))
            ),
            address(this.rewardsCoordinatorImplementation()),
            "rewardsCoordinator: implementation set incorrectly"
        );
        assertEq(
            this.eigenLayerProxyAdmin().getProxyImplementation(
                ITransparentUpgradeableProxy(payable(address(this.delegationManager())))
            ),
            address(this.delegationManagerImplementation()),
            "delegationManager: implementation set incorrectly"
        );
        assertEq(
            this.eigenLayerProxyAdmin().getProxyImplementation(ITransparentUpgradeableProxy(payable(address(this.strategyManager())))),
            address(this.strategyManagerImplementation()),
            "strategyManager: implementation set incorrectly"
        );
        assertEq(
            this.eigenLayerProxyAdmin().getProxyImplementation(ITransparentUpgradeableProxy(payable(address(this.eigenPodManager())))),
            address(this.eigenPodManagerImplementation()),
            "eigenPodManager: implementation set incorrectly"
        );

        assertEq(
            this.eigenPodBeacon().implementation(),
            address(this.eigenPodImplementation()),
            "eigenPodBeacon: implementation set incorrectly"
        );
    }

    function verifyContractsInitialized(bool isInitialDeployment) external {
        // AVSDirectory
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        this.avsDirectory().initialize(address(0), 0);
        // RewardsCoordinator
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        this.rewardsCoordinator().initialize(
            address(0),
            0, // initialPausedStatus
            address(0), // rewardsUpdater
            0, // activationDelay
            0 // defaultSplitBips
        );
        // DelegationManager
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        this.delegationManager().initialize(address(0), 0);
        // StrategyManager
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        this.strategyManager().initialize(address(0), address(0), 0);
        // EigenPodManager
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        this.eigenPodManager().initialize(address(0), 0);
    }

    function verifyInitializationParams() external view {
        // AVSDirectory
        assertTrue(
            this.avsDirectory().pauserRegistry() == this.pauserRegistry(),
            "avsdirectory: pauser registry not set correctly"
        );
        assertEq(this.avsDirectory().owner(), this.executorMultisig(), "avsdirectory: owner not set correctly");

        // RewardsCoordinator
        assertTrue(
            this.rewardsCoordinator().pauserRegistry() == this.pauserRegistry(),
            "rewardsCoordinator: pauser registry not set correctly"
        );
        assertEq(
            this.rewardsCoordinator().MAX_REWARDS_DURATION(),
            Env.MAX_REWARDS_DURATION(),
            "rewardsCoordinator: maxRewardsDuration not set correctly"
        );
        assertEq(
            this.rewardsCoordinator().MAX_RETROACTIVE_LENGTH(),
            Env.MAX_RETROACTIVE_LENGTH(),
            "rewardsCoordinator: maxRetroactiveLength not set correctly"
        );
        assertEq(
            this.rewardsCoordinator().MAX_FUTURE_LENGTH(),
            Env.MAX_FUTURE_LENGTH(),
            "rewardsCoordinator: maxFutureLength not set correctly"
        );
        assertEq(
            this.rewardsCoordinator().GENESIS_REWARDS_TIMESTAMP(),
            Env.GENESIS_REWARDS_TIMESTAMP(),
            "rewardsCoordinator: genesisRewardsTimestamp not set correctly"
        );
        assertEq(
            this.rewardsCoordinator().activationDelay(),
            Env.ACTIVATION_DELAY(),
            "rewardsCoordinator: activationDelay not set correctly"
        );
        assertEq(
            this.rewardsCoordinator().CALCULATION_INTERVAL_SECONDS(),
            Env.CALCULATION_INTERVAL_SECONDS(),
            "rewardsCoordinator: CALCULATION_INTERVAL_SECONDS not set correctly"
        );
        assertEq(
            this.rewardsCoordinator().defaultOperatorSplitBips(),
            Env.DEFAULT_SPLIT_BIPS(),
            "rewardsCoordinator: defaultSplitBips not set correctly"
        );

        // DelegationManager
        assertTrue(
            this.delegationManager().pauserRegistry() == this.pauserRegistry(),
            "delegationManager: pauser registry not set correctly"
        );
        assertEq(this.delegationManager().owner(), this.executorMultisig(), "delegationManager: owner not set correctly");

        // StrategyManager
        assertTrue(
            this.strategyManager().pauserRegistry() == this.pauserRegistry(),
            "strategyManager: pauser registry not set correctly"
        );
        assertEq(this.strategyManager().owner(), this.executorMultisig(), "strategyManager: owner not set correctly");

        // EigenPodManager
        assertTrue(
            this.eigenPodManager().pauserRegistry() == this.pauserRegistry(),
            "eigenPodManager: pauser registry not set correctly"
        );
        assertEq(this.eigenPodManager().owner(), this.executorMultisig(), "eigenPodManager: owner not set correctly");
        assertEq(
            address(this.eigenPodManager().ethPOS()),
            address(Env.ethPOS()),
            "eigenPodManager: ethPOS not set correctly"
        );

        // EigenPodBeacon
        assertEq(this.eigenPodBeacon().owner(), this.executorMultisig(), "eigenPodBeacon: owner not set correctly");

        // EigenPodImplementation
        assertEq(
            this.eigenPodImplementation().GENESIS_TIME(),
            Env.EIGENPOD_GENESIS_TIME(),
            "eigenPodImplementation: GENESIS TIME not set correctly"
        );
        assertEq(
            address(this.eigenPodImplementation().ethPOS()),
            address(Env.ethPOS()),
            "eigenPodImplementation: ethPOS not set correctly"
        );

        // Pausing Permissions
        assertTrue(this.pauserRegistry().isPauser(this.operationsMultisig()), "pauserRegistry: operationsMultisig is not pauser");
        assertTrue(this.pauserRegistry().isPauser(this.executorMultisig()), "pauserRegistry: executorMultisig is not pauser");
        assertTrue(this.pauserRegistry().isPauser(this.pauserMultisig()), "pauserRegistry: pauserMultisig is not pauser");
        assertEq(this.pauserRegistry().unpauser(), this.executorMultisig(), "pauserRegistry: unpauser not set correctly");
    }

    function initialize() external {
        // No initialization needed for Zeus implementation
        // All contract addresses are loaded directly from Env.sol
    }
} 