// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "./Env.sol";

/// @title CoreContractsDeployer
/// @notice Provides reusable helpers for deploying individual core contract implementations.
/// Usage:
/// ```solidity
/// vm.startBroadcast();
/// deployPermissionController();
/// deployKeyRegistrar();
/// vm.stopBroadcast();
/// ```
abstract contract CoreContractsDeployer is EOADeployer {
    using Env for *;

    /// permissions/
    function deployPermissionController() internal onlyEOA returns (PermissionController deployed) {
        deployed = new PermissionController();
        deployImpl({name: type(PermissionController).name, deployedTo: address(deployed)});
    }

    function deployKeyRegistrar() internal onlyEOA returns (KeyRegistrar deployed) {
        deployed = new KeyRegistrar({
            _permissionController: Env.proxy.permissionController(),
            _allocationManager: Env.proxy.allocationManager(),
            _version: Env.deployVersion()
        });
        deployImpl({name: type(KeyRegistrar).name, deployedTo: address(deployed)});
    }

    /// core/
    function deployAllocationManagerView() internal onlyEOA returns (AllocationManagerView deployed) {
        deployed = new AllocationManagerView({
            _delegation: Env.proxy.delegationManager(),
            _eigenStrategy: Env.proxy.eigenStrategy(),
            _DEALLOCATION_DELAY: Env.MIN_WITHDRAWAL_DELAY(),
            _ALLOCATION_CONFIGURATION_DELAY: Env.ALLOCATION_CONFIGURATION_DELAY()
        });
        deployImpl({name: type(AllocationManagerView).name, deployedTo: address(deployed)});
    }

    function deployAllocationManager() internal onlyEOA returns (AllocationManager deployed) {
        deployed = new AllocationManager({
            _allocationManagerView: Env.impl.allocationManagerView(),
            _delegation: Env.proxy.delegationManager(),
            _eigenStrategy: Env.proxy.eigenStrategy(),
            _pauserRegistry: Env.impl.pauserRegistry(),
            _permissionController: Env.proxy.permissionController(),
            _DEALLOCATION_DELAY: Env.MIN_WITHDRAWAL_DELAY(),
            _ALLOCATION_CONFIGURATION_DELAY: Env.ALLOCATION_CONFIGURATION_DELAY()
        });
        deployImpl({name: type(AllocationManager).name, deployedTo: address(deployed)});
    }

    function deployAVSDirectory() internal onlyEOA returns (AVSDirectory deployed) {
        deployed = new AVSDirectory({
            _delegation: Env.proxy.delegationManager(),
            _pauserRegistry: Env.impl.pauserRegistry(),
            _version: Env.deployVersion()
        });
        deployImpl({name: type(AVSDirectory).name, deployedTo: address(deployed)});
    }

    function deployDelegationManager() internal onlyEOA returns (DelegationManager deployed) {
        deployed = new DelegationManager({
            _strategyManager: Env.proxy.strategyManager(),
            _eigenPodManager: Env.proxy.eigenPodManager(),
            _allocationManager: Env.proxy.allocationManager(),
            _pauserRegistry: Env.impl.pauserRegistry(),
            _permissionController: Env.proxy.permissionController(),
            _MIN_WITHDRAWAL_DELAY: Env.MIN_WITHDRAWAL_DELAY(),
            _version: Env.deployVersion()
        });
        deployImpl({name: type(DelegationManager).name, deployedTo: address(deployed)});
    }

    function deployProtocolRegistry() internal onlyEOA returns (ProtocolRegistry deployed) {
        deployed = new ProtocolRegistry();
        deployImpl({name: type(ProtocolRegistry).name, deployedTo: address(deployed)});
    }

    function deployReleaseManager() internal onlyEOA returns (ReleaseManager deployed) {
        deployed = new ReleaseManager({_permissionController: Env.proxy.permissionController()});
        deployImpl({name: type(ReleaseManager).name, deployedTo: address(deployed)});
    }

    function deployRewardsCoordinator() internal onlyEOA returns (RewardsCoordinator deployed) {
        deployed = new RewardsCoordinator({
            params: IRewardsCoordinatorTypes.RewardsCoordinatorConstructorParams({
                delegationManager: Env.proxy.delegationManager(),
                strategyManager: Env.proxy.strategyManager(),
                allocationManager: Env.proxy.allocationManager(),
                pauserRegistry: Env.impl.pauserRegistry(),
                permissionController: Env.proxy.permissionController(),
                CALCULATION_INTERVAL_SECONDS: Env.CALCULATION_INTERVAL_SECONDS(),
                MAX_REWARDS_DURATION: Env.MAX_REWARDS_DURATION(),
                MAX_RETROACTIVE_LENGTH: Env.MAX_RETROACTIVE_LENGTH(),
                MAX_FUTURE_LENGTH: Env.MAX_FUTURE_LENGTH(),
                GENESIS_REWARDS_TIMESTAMP: Env.GENESIS_REWARDS_TIMESTAMP()
            })
        });
        deployImpl({name: type(RewardsCoordinator).name, deployedTo: address(deployed)});
    }

    function deployStrategyManager() internal onlyEOA returns (StrategyManager deployed) {
        deployed = new StrategyManager({
            _allocationManager: Env.proxy.allocationManager(),
            _delegation: Env.proxy.delegationManager(),
            _pauserRegistry: Env.impl.pauserRegistry(),
            _version: Env.deployVersion()
        });
        deployImpl({name: type(StrategyManager).name, deployedTo: address(deployed)});
    }

    /// pods/
    function deployEigenPodManager() internal onlyEOA returns (EigenPodManager deployed) {
        deployed = new EigenPodManager({
            _ethPOS: Env.ethPOS(),
            _eigenPodBeacon: Env.beacon.eigenPod(),
            _delegationManager: Env.proxy.delegationManager(),
            _pauserRegistry: Env.impl.pauserRegistry()
        });
        deployImpl({name: type(EigenPodManager).name, deployedTo: address(deployed)});
    }

    function deployEigenPod() internal onlyEOA returns (EigenPod deployed) {
        deployed = new EigenPod({_ethPOS: Env.ethPOS(), _eigenPodManager: Env.proxy.eigenPodManager()});
        deployImpl({name: type(EigenPod).name, deployedTo: address(deployed)});
    }

    /// strategies/
    function deployEigenStrategy() internal onlyEOA returns (EigenStrategy deployed) {
        deployed = new EigenStrategy({
            _strategyManager: Env.proxy.strategyManager(),
            _pauserRegistry: Env.impl.pauserRegistry()
        });
        deployImpl({name: type(EigenStrategy).name, deployedTo: address(deployed)});
    }

    function deployStrategyBase() internal onlyEOA returns (StrategyBase deployed) {
        deployed = new StrategyBase({
            _strategyManager: Env.proxy.strategyManager(),
            _pauserRegistry: Env.impl.pauserRegistry()
        });
        deployImpl({name: type(StrategyBase).name, deployedTo: address(deployed)});
    }

    function deployStrategyBaseTVLLimits() internal onlyEOA returns (StrategyBaseTVLLimits deployed) {
        deployed = new StrategyBaseTVLLimits({
            _strategyManager: Env.proxy.strategyManager(),
            _pauserRegistry: Env.impl.pauserRegistry()
        });
        deployImpl({name: type(StrategyBaseTVLLimits).name, deployedTo: address(deployed)});
    }

    function deployStrategyFactory() internal onlyEOA returns (StrategyFactory deployed) {
        deployed = new StrategyFactory({
            _strategyManager: Env.proxy.strategyManager(),
            _pauserRegistry: Env.impl.pauserRegistry()
        });
        deployImpl({name: type(StrategyFactory).name, deployedTo: address(deployed)});
    }

    /// multichain/
    function deployBN254CertificateVerifier() internal onlyEOA returns (BN254CertificateVerifier deployed) {
        deployed = new BN254CertificateVerifier({_operatorTableUpdater: Env.proxy.operatorTableUpdater()});
        deployImpl({name: type(BN254CertificateVerifier).name, deployedTo: address(deployed)});
    }

    function deployCrossChainRegistry() internal onlyEOA returns (CrossChainRegistry deployed) {
        deployed = new CrossChainRegistry({
            _allocationManager: Env.proxy.allocationManager(),
            _keyRegistrar: Env.proxy.keyRegistrar(),
            _permissionController: Env.proxy.permissionController(),
            _pauserRegistry: Env.impl.pauserRegistry()
        });
        deployImpl({name: type(CrossChainRegistry).name, deployedTo: address(deployed)});
    }

    function deployECDSACertificateVerifier() internal onlyEOA returns (ECDSACertificateVerifier deployed) {
        deployed = new ECDSACertificateVerifier({
            _operatorTableUpdater: Env.proxy.operatorTableUpdater(),
            _version: Env.deployVersion()
        });
        deployImpl({name: type(ECDSACertificateVerifier).name, deployedTo: address(deployed)});
    }

    function deployOperatorTableUpdater() internal onlyEOA returns (OperatorTableUpdater deployed) {
        deployed = new OperatorTableUpdater({
            _bn254CertificateVerifier: Env.proxy.bn254CertificateVerifier(),
            _ecdsaCertificateVerifier: Env.proxy.ecdsaCertificateVerifier(),
            _pauserRegistry: Env.impl.pauserRegistry()
        });
        deployImpl({name: type(OperatorTableUpdater).name, deployedTo: address(deployed)});
    }

    /// avs/
    function deployTaskMailbox() internal onlyEOA returns (TaskMailbox deployed) {
        deployed = new TaskMailbox({
            _bn254CertificateVerifier: address(Env.proxy.bn254CertificateVerifier()),
            _ecdsaCertificateVerifier: address(Env.proxy.ecdsaCertificateVerifier()),
            _maxTaskSLA: Env.MAX_TASK_SLA()
        });
        deployImpl({name: type(TaskMailbox).name, deployedTo: address(deployed)});
    }
}
