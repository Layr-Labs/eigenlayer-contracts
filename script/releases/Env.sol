// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Vm.sol";
import "zeus-templates/utils/ZEnvHelpers.sol";

import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

/// core/
import "src/contracts/core/AllocationManager.sol";
import "src/contracts/core/AVSDirectory.sol";
import "src/contracts/core/DelegationManager.sol";
import "src/contracts/core/RewardsCoordinator.sol";
import "src/contracts/interfaces/IRewardsCoordinator.sol";
import "src/contracts/core/StrategyManager.sol";
import "src/contracts/core/ReleaseManager.sol";

/// permissions/
import "src/contracts/permissions/PauserRegistry.sol";
import "src/contracts/permissions/PermissionController.sol";
import "src/contracts/permissions/KeyRegistrar.sol";

/// pods/
import "src/contracts/pods/EigenPod.sol";
import "src/contracts/pods/EigenPodManager.sol";

/// strategies/
import "src/contracts/strategies/EigenStrategy.sol";
import "src/contracts/strategies/StrategyBase.sol";
import "src/contracts/strategies/StrategyBaseTVLLimits.sol";
import "src/contracts/strategies/StrategyFactory.sol";

/// token/
import "src/contracts/interfaces/IEigen.sol";
import "src/contracts/interfaces/IBackingEigen.sol";
import "src/contracts/token/Eigen.sol";
import "src/contracts/token/BackingEigen.sol";

/// multichain/
import "src/contracts/multichain/CrossChainRegistry.sol";
import "src/contracts/multichain/OperatorTableUpdater.sol";
import "src/contracts/multichain/ECDSACertificateVerifier.sol";
import "src/contracts/multichain/BN254CertificateVerifier.sol";

/// avs/
import "src/contracts/avs/task/TaskMailbox.sol";

// For destination chains
import "src/test/mocks/EmptyContract.sol";

library Env {
    using ZEnvHelpers for *;

    /// Dummy types and variables to facilitate syntax, e.g: `Env.proxy.delegationManager()`
    enum DeployedProxy {
        A
    }
    enum DeployedBeacon {
        A
    }
    enum DeployedImpl {
        A
    }
    enum DeployedInstance {
        A
    }

    DeployedProxy internal constant proxy = DeployedProxy.A;
    DeployedBeacon internal constant beacon = DeployedBeacon.A;
    DeployedImpl internal constant impl = DeployedImpl.A;
    DeployedInstance internal constant instance = DeployedInstance.A;

    /**
     * env
     */
    function env() internal view returns (string memory) {
        return _string("ZEUS_ENV");
    }

    function envVersion() internal view returns (string memory) {
        return _string("ZEUS_ENV_VERSION");
    }

    function deployVersion() internal view returns (string memory) {
        return _string("ZEUS_DEPLOY_TO_VERSION");
    }

    function executorMultisig() internal view returns (address) {
        return _envAddress("executorMultisig");
    }

    function opsMultisig() internal view returns (address) {
        return _envAddress("operationsMultisig");
    }

    function communityMultisig() internal view returns (address) {
        return _envAddress("communityMultisig");
    }

    function protocolCouncilMultisig() internal view returns (address) {
        return _envAddress("protocolCouncilMultisig");
    }

    function pauserMultisig() internal view returns (address) {
        return _envAddress("pauserMultisig");
    }

    function multichainDeployerMultisig() internal view returns (address) {
        return _envAddress("multichainDeployerMultisig");
    }

    function createX() internal view returns (address) {
        return _envAddress("createX");
    }

    function proxyAdmin() internal view returns (address) {
        return _envAddress("proxyAdmin");
    }

    function ethPOS() internal view returns (IETHPOSDeposit) {
        return IETHPOSDeposit(_envAddress("ethPOS"));
    }

    function timelockController() internal view returns (TimelockController) {
        return TimelockController(payable(_envAddress("timelockController")));
    }

    function multiSendCallOnly() internal view returns (address) {
        return _envAddress("MultiSendCallOnly");
    }

    function MIN_WITHDRAWAL_DELAY() internal view returns (uint32) {
        return _envU32("MIN_WITHDRAWAL_DELAY");
    }

    function ALLOCATION_CONFIGURATION_DELAY() internal view returns (uint32) {
        return _envU32("ALLOCATION_CONFIGURATION_DELAY");
    }

    function REWARDS_UPDATER() internal view returns (address) {
        return _envAddress("REWARDS_COORDINATOR_UPDATER");
    }

    function ACTIVATION_DELAY() internal view returns (uint32) {
        return _envU32("REWARDS_COORDINATOR_ACTIVATION_DELAY");
    }

    function DEFAULT_SPLIT_BIPS() internal view returns (uint16) {
        return _envU16("REWARDS_COORDINATOR_DEFAULT_OPERATOR_SPLIT_BIPS");
    }

    function CALCULATION_INTERVAL_SECONDS() internal view returns (uint32) {
        return _envU32("REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS");
    }

    function MAX_REWARDS_DURATION() internal view returns (uint32) {
        return _envU32("REWARDS_COORDINATOR_MAX_REWARDS_DURATION");
    }

    function MAX_RETROACTIVE_LENGTH() internal view returns (uint32) {
        return _envU32("REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH");
    }

    function MAX_FUTURE_LENGTH() internal view returns (uint32) {
        return _envU32("REWARDS_COORDINATOR_MAX_FUTURE_LENGTH");
    }

    function GENESIS_REWARDS_TIMESTAMP() internal view returns (uint32) {
        return _envU32("REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP");
    }

    function REWARDS_PAUSE_STATUS() internal view returns (uint256) {
        return _envU256("REWARDS_COORDINATOR_PAUSE_STATUS");
    }

    function SLASH_ESCROW_DELAY() internal view returns (uint32) {
        return _envU32("SLASH_ESCROW_DELAY");
    }

    function CROSS_CHAIN_REGISTRY_PAUSE_STATUS() internal view returns (uint256) {
        return _envU256("CROSS_CHAIN_REGISTRY_INIT_PAUSE_STATUS");
    }

    function TABLE_UPDATE_CADENCE() internal view returns (uint32) {
        return _envU32("TABLE_UPDATE_CADENCE");
    }

    function MAX_TASK_SLA() internal view returns (uint96) {
        // MAX_TASK_SLA is stored as uint256 in zeus env, cast to uint96
        return uint96(_envU256("MAX_TASK_SLA"));
    }

    function isSourceChain() internal view returns (bool) {
        return _envBool("SOURCE_CHAIN");
    }

    function isDestinationChain() internal view returns (bool) {
        return _envBool("DESTINATION_CHAIN");
    }

    /**
     * core/
     */
    function allocationManager(
        DeployedProxy
    ) internal view returns (AllocationManager) {
        return AllocationManager(_deployedProxy(type(AllocationManager).name));
    }

    function allocationManager(
        DeployedImpl
    ) internal view returns (AllocationManager) {
        return AllocationManager(_deployedImpl(type(AllocationManager).name));
    }

    function avsDirectory(
        DeployedProxy
    ) internal view returns (AVSDirectory) {
        return AVSDirectory(_deployedProxy(type(AVSDirectory).name));
    }

    function avsDirectory(
        DeployedImpl
    ) internal view returns (AVSDirectory) {
        return AVSDirectory(_deployedImpl(type(AVSDirectory).name));
    }

    function delegationManager(
        DeployedProxy
    ) internal view returns (DelegationManager) {
        return DelegationManager(_deployedProxy(type(DelegationManager).name));
    }

    function delegationManager(
        DeployedImpl
    ) internal view returns (DelegationManager) {
        return DelegationManager(_deployedImpl(type(DelegationManager).name));
    }

    function rewardsCoordinator(
        DeployedProxy
    ) internal view returns (RewardsCoordinator) {
        return RewardsCoordinator(_deployedProxy(type(RewardsCoordinator).name));
    }

    function rewardsCoordinator(
        DeployedImpl
    ) internal view returns (RewardsCoordinator) {
        return RewardsCoordinator(_deployedImpl(type(RewardsCoordinator).name));
    }

    function strategyManager(
        DeployedProxy
    ) internal view returns (StrategyManager) {
        return StrategyManager(_deployedProxy(type(StrategyManager).name));
    }

    function strategyManager(
        DeployedImpl
    ) internal view returns (StrategyManager) {
        return StrategyManager(_deployedImpl(type(StrategyManager).name));
    }

    function releaseManager(
        DeployedProxy
    ) internal view returns (ReleaseManager) {
        return ReleaseManager(_deployedProxy(type(ReleaseManager).name));
    }

    function releaseManager(
        DeployedImpl
    ) internal view returns (ReleaseManager) {
        return ReleaseManager(_deployedImpl(type(ReleaseManager).name));
    }

    /**
     * permissions/
     */
    function pauserRegistry(
        DeployedImpl
    ) internal view returns (PauserRegistry) {
        return PauserRegistry(_deployedImpl(type(PauserRegistry).name));
    }

    function permissionController(
        DeployedProxy
    ) internal view returns (PermissionController) {
        return PermissionController(_deployedProxy(type(PermissionController).name));
    }

    function permissionController(
        DeployedImpl
    ) internal view returns (PermissionController) {
        return PermissionController(_deployedImpl(type(PermissionController).name));
    }

    function keyRegistrar(
        DeployedProxy
    ) internal view returns (KeyRegistrar) {
        return KeyRegistrar(_deployedProxy(type(KeyRegistrar).name));
    }

    function keyRegistrar(
        DeployedImpl
    ) internal view returns (KeyRegistrar) {
        return KeyRegistrar(_deployedImpl(type(KeyRegistrar).name));
    }

    /**
     * pods/
     */
    function eigenPod(
        DeployedBeacon
    ) internal view returns (UpgradeableBeacon) {
        return UpgradeableBeacon(_deployedBeacon(type(EigenPod).name));
    }

    function eigenPod(
        DeployedImpl
    ) internal view returns (EigenPod) {
        return EigenPod(payable(_deployedImpl(type(EigenPod).name)));
    }

    function eigenPodManager(
        DeployedProxy
    ) internal view returns (EigenPodManager) {
        return EigenPodManager(_deployedProxy(type(EigenPodManager).name));
    }

    function eigenPodManager(
        DeployedImpl
    ) internal view returns (EigenPodManager) {
        return EigenPodManager(_deployedImpl(type(EigenPodManager).name));
    }

    /**
     * strategies/
     */
    function eigenStrategy(
        DeployedProxy
    ) internal view returns (EigenStrategy) {
        return EigenStrategy(_deployedProxy(type(EigenStrategy).name));
    }

    function eigenStrategy(
        DeployedImpl
    ) internal view returns (EigenStrategy) {
        return EigenStrategy(_deployedImpl(type(EigenStrategy).name));
    }

    // Beacon proxy
    function strategyBase(
        DeployedBeacon
    ) internal view returns (UpgradeableBeacon) {
        return UpgradeableBeacon(_deployedBeacon(type(StrategyBase).name));
    }

    // Beaconed impl
    function strategyBase(
        DeployedImpl
    ) internal view returns (StrategyBase) {
        return StrategyBase(_deployedImpl(type(StrategyBase).name));
    }

    // Returns the number of proxy instances
    function strategyBaseTVLLimits_Count(
        DeployedInstance
    ) internal view returns (uint256) {
        return _deployedInstanceCount(type(StrategyBaseTVLLimits).name);
    }

    // Returns the proxy instance at index `i`
    function strategyBaseTVLLimits(DeployedInstance, uint256 i) internal view returns (StrategyBaseTVLLimits) {
        return StrategyBaseTVLLimits(_deployedInstance(type(StrategyBaseTVLLimits).name, i));
    }

    function strategyBaseTVLLimits(
        DeployedImpl
    ) internal view returns (StrategyBaseTVLLimits) {
        return StrategyBaseTVLLimits(_deployedImpl(type(StrategyBaseTVLLimits).name));
    }

    function strategyFactory(
        DeployedProxy
    ) internal view returns (StrategyFactory) {
        return StrategyFactory(_deployedProxy(type(StrategyFactory).name));
    }

    function strategyFactory(
        DeployedImpl
    ) internal view returns (StrategyFactory) {
        return StrategyFactory(_deployedImpl(type(StrategyFactory).name));
    }

    /**
     * token/
     */
    function eigen(
        DeployedProxy
    ) internal view returns (Eigen) {
        return Eigen(_deployedProxy(type(Eigen).name));
    }

    function eigen(
        DeployedImpl
    ) internal view returns (Eigen) {
        return Eigen(_deployedImpl(type(Eigen).name));
    }

    function beigen(
        DeployedProxy
    ) internal view returns (IBackingEigen) {
        return IBackingEigen(_deployedProxy(type(BackingEigen).name));
    }

    function beigen(
        DeployedImpl
    ) internal view returns (IBackingEigen) {
        return IBackingEigen(_deployedImpl(type(BackingEigen).name));
    }

    /**
     * multichain/
     */
    function emptyContract(
        DeployedImpl
    ) internal view returns (EmptyContract) {
        return EmptyContract(_deployedImpl(type(EmptyContract).name));
    }

    function crossChainRegistry(
        DeployedProxy
    ) internal view returns (CrossChainRegistry) {
        return CrossChainRegistry(_deployedProxy(type(CrossChainRegistry).name));
    }

    function crossChainRegistry(
        DeployedImpl
    ) internal view returns (CrossChainRegistry) {
        return CrossChainRegistry(_deployedImpl(type(CrossChainRegistry).name));
    }

    function operatorTableUpdater(
        DeployedProxy
    ) internal view returns (OperatorTableUpdater) {
        return OperatorTableUpdater(_deployedProxy(type(OperatorTableUpdater).name));
    }

    function operatorTableUpdater(
        DeployedImpl
    ) internal view returns (OperatorTableUpdater) {
        return OperatorTableUpdater(_deployedImpl(type(OperatorTableUpdater).name));
    }

    function ecdsaCertificateVerifier(
        DeployedProxy
    ) internal view returns (ECDSACertificateVerifier) {
        return ECDSACertificateVerifier(_deployedProxy(type(ECDSACertificateVerifier).name));
    }

    function ecdsaCertificateVerifier(
        DeployedImpl
    ) internal view returns (ECDSACertificateVerifier) {
        return ECDSACertificateVerifier(_deployedImpl(type(ECDSACertificateVerifier).name));
    }

    function bn254CertificateVerifier(
        DeployedProxy
    ) internal view returns (BN254CertificateVerifier) {
        return BN254CertificateVerifier(_deployedProxy(type(BN254CertificateVerifier).name));
    }

    function bn254CertificateVerifier(
        DeployedImpl
    ) internal view returns (BN254CertificateVerifier) {
        return BN254CertificateVerifier(_deployedImpl(type(BN254CertificateVerifier).name));
    }

    function taskMailbox(
        DeployedProxy
    ) internal view returns (TaskMailbox) {
        return TaskMailbox(_deployedProxy(type(TaskMailbox).name));
    }

    function taskMailbox(
        DeployedImpl
    ) internal view returns (TaskMailbox) {
        return TaskMailbox(_deployedImpl(type(TaskMailbox).name));
    }

    /**
     * Helpers
     */
    function _deployedInstance(string memory name, uint256 idx) private view returns (address) {
        return ZEnvHelpers.state().deployedInstance(name, idx);
    }

    function _deployedInstanceCount(
        string memory name
    ) private view returns (uint256) {
        return ZEnvHelpers.state().deployedInstanceCount(name);
    }

    function _deployedProxy(
        string memory name
    ) private view returns (address) {
        return ZEnvHelpers.state().deployedProxy(name);
    }

    function _deployedBeacon(
        string memory name
    ) private view returns (address) {
        return ZEnvHelpers.state().deployedBeacon(name);
    }

    function _deployedImpl(
        string memory name
    ) private view returns (address) {
        return ZEnvHelpers.state().deployedImpl(name);
    }

    function _envAddress(
        string memory key
    ) private view returns (address) {
        return ZEnvHelpers.state().envAddress(key);
    }

    function _envU256(
        string memory key
    ) private view returns (uint256) {
        return ZEnvHelpers.state().envU256(key);
    }

    function _envU64(
        string memory key
    ) private view returns (uint64) {
        return ZEnvHelpers.state().envU64(key);
    }

    function _envU32(
        string memory key
    ) private view returns (uint32) {
        return ZEnvHelpers.state().envU32(key);
    }

    function _envU16(
        string memory key
    ) private view returns (uint16) {
        return ZEnvHelpers.state().envU16(key);
    }

    function _envBool(
        string memory key
    ) private view returns (bool) {
        return ZEnvHelpers.state().envBool(key);
    }

    address internal constant VM_ADDRESS = address(uint160(uint256(keccak256("hevm cheat code"))));
    Vm internal constant vm = Vm(VM_ADDRESS);

    function _string(
        string memory key
    ) private view returns (string memory) {
        return vm.envString(key);
    }

    /**
     * Test Helpers
     */

    /// @dev Query and return `proxyAdmin.getProxyImplementation(proxy)`
    function _getProxyImpl(
        address _proxy
    ) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyImplementation(ITransparentUpgradeableProxy(_proxy));
    }

    /// @dev Query and return `proxyAdmin.getProxyAdmin(proxy)`
    function _getProxyAdmin(
        address _proxy
    ) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyAdmin(ITransparentUpgradeableProxy(_proxy));
    }

    function _strEq(string memory a, string memory b) internal pure returns (bool) {
        return keccak256(bytes(a)) == keccak256(bytes(b));
    }
}
