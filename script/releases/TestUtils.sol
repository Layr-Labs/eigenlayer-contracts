// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./Env.sol";
import "forge-std/Vm.sol";
import "src/contracts/mixins/SplitContractMixin.sol";

/// @notice Utility library for testing contract deployments
/// @dev This library exposes the following test functions:
/// @dev - validateProxyAdmins - Check that proxy admins are correctly set.
/// @dev - validateImplConstructors - Check that implementation constructors are correctly set.
/// @dev - validateImplsInitialized - Check that implementation cannot be initialized.
/// @dev - validateImplAddressesMatchProxy - Check that implementation addresses match the proxy admin's reported implementation address.
/// @dev - validateProtocolVersion - Check that the protocol version is correctly set.
library TestUtils {
    using Env for *;

    address internal constant VM_ADDRESS = address(uint160(uint256(keccak256("hevm cheat code"))));
    Vm internal constant vm = Vm(VM_ADDRESS);

    /**
     *
     *                         PROXY ADMIN VALIDATION
     *
     */

    /// @dev This function is run on *all* deployed contracts to ensure that the proxyAdmin is correctly set.
    function validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();
        /// permissions/
        assertTrue(
            _getProxyAdmin(address(Env.proxy.permissionController())) == pa, "permissionController proxyAdmin incorrect"
        );
        assertTrue(_getProxyAdmin(address(Env.proxy.keyRegistrar())) == pa, "keyRegistrar proxyAdmin incorrect");

        /// core/
        assertTrue(
            _getProxyAdmin(address(Env.proxy.allocationManager())) == pa, "allocationManager proxyAdmin incorrect"
        );
        assertTrue(_getProxyAdmin(address(Env.proxy.avsDirectory())) == pa, "avsDirectory proxyAdmin incorrect");
        assertTrue(
            _getProxyAdmin(address(Env.proxy.delegationManager())) == pa, "delegationManager proxyAdmin incorrect"
        );
        assertTrue(_getProxyAdmin(address(Env.proxy.protocolRegistry())) == pa, "protocolRegistry proxyAdmin incorrect");
        assertTrue(_getProxyAdmin(address(Env.proxy.releaseManager())) == pa, "releaseManager proxyAdmin incorrect");
        assertTrue(
            _getProxyAdmin(address(Env.proxy.rewardsCoordinator())) == pa, "rewardsCoordinator proxyAdmin incorrect"
        );
        assertTrue(_getProxyAdmin(address(Env.proxy.strategyManager())) == pa, "strategyManager proxyAdmin incorrect");

        /// pods/
        assertTrue(Env.beacon.eigenPod().owner() == Env.executorMultisig(), "eigenPod beacon owner incorrect");
        assertTrue(_getProxyAdmin(address(Env.proxy.eigenPodManager())) == pa, "eigenPodManager proxyAdmin incorrect");

        /// strategies/
        assertTrue(_getProxyAdmin(address(Env.proxy.eigenStrategy())) == pa, "eigenStrategy proxyAdmin incorrect");
        assertTrue(Env.beacon.strategyBase().owner() == Env.executorMultisig(), "strategyBase beacon owner incorrect");

        uint256 count = Env.instance.strategyBaseTVLLimits_Count();
        for (uint256 i = 0; i < count; i++) {
            assertTrue(
                _getProxyAdmin(address(Env.instance.strategyBaseTVLLimits(i))) == pa,
                "strategyBaseTVLLimits proxyAdmin incorrect"
            );
        }

        assertTrue(_getProxyAdmin(address(Env.proxy.strategyFactory())) == pa, "strategyFactory proxyAdmin incorrect");

        /// multichain/
        assertTrue(
            _getProxyAdmin(address(Env.proxy.crossChainRegistry())) == pa, "crossChainRegistry proxyAdmin incorrect"
        );
        assertTrue(
            _getProxyAdmin(address(Env.proxy.operatorTableUpdater())) == pa, "operatorTableUpdater proxyAdmin incorrect"
        );
        assertTrue(
            _getProxyAdmin(address(Env.proxy.ecdsaCertificateVerifier())) == pa,
            "ecdsaCertificateVerifier proxyAdmin incorrect"
        );
        assertTrue(
            _getProxyAdmin(address(Env.proxy.bn254CertificateVerifier())) == pa,
            "bn254CertificateVerifier proxyAdmin incorrect"
        );

        /// avs/
        assertTrue(_getProxyAdmin(address(Env.proxy.taskMailbox())) == pa, "taskMailbox proxyAdmin incorrect");
    }

    /**
     *
     *                        IMPLEMENTATION VALIDATION
     *
     */
    function validateImplConstructors() internal view {
        {
            /// permissions/
            PauserRegistry registry = Env.impl.pauserRegistry();
            assertTrue(registry.isPauser(Env.pauserMultisig()), "pauser multisig should be pauser");
            assertTrue(registry.isPauser(Env.opsMultisig()), "ops multisig should be pauser");
            assertTrue(registry.isPauser(Env.executorMultisig()), "executor multisig should be pauser");
            assertTrue(registry.unpauser() == Env.executorMultisig(), "executor multisig should be unpauser");

            // PermissionController has no initial storage

            KeyRegistrar keyRegistrar = Env.impl.keyRegistrar();
            assertTrue(
                keyRegistrar.permissionController() == Env.proxy.permissionController(),
                "keyRegistrar permissionController incorrect"
            );
            assertTrue(
                keyRegistrar.allocationManager() == Env.proxy.allocationManager(),
                "keyRegistrar allocationManager incorrect"
            );
        }

        {
            /// core/
            AllocationManagerView allocationManagerView = Env.impl.allocationManagerView();
            assertTrue(
                allocationManagerView.delegation() == Env.proxy.delegationManager(),
                "allocationManagerView delegation incorrect"
            );
            assertTrue(
                allocationManagerView.eigenStrategy() == Env.proxy.eigenStrategy(),
                "allocationManagerView eigenStrategy incorrect"
            );
            assertTrue(
                allocationManagerView.DEALLOCATION_DELAY() == Env.MIN_WITHDRAWAL_DELAY(),
                "allocationManagerView DEALLOCATION_DELAY incorrect"
            );
            assertTrue(
                allocationManagerView.ALLOCATION_CONFIGURATION_DELAY() == Env.ALLOCATION_CONFIGURATION_DELAY(),
                "allocationManagerView ALLOCATION_CONFIGURATION_DELAY incorrect"
            );

            AllocationManager allocationManager = Env.impl.allocationManager();
            assertTrue(
                allocationManager.viewImplementation() == address(Env.impl.allocationManagerView()),
                "allocationManager allocationManagerView incorrect"
            );
            assertTrue(
                allocationManager.delegation() == Env.proxy.delegationManager(),
                "allocationManager delegation incorrect"
            );
            assertTrue(
                allocationManager.eigenStrategy() == Env.proxy.eigenStrategy(),
                "allocationManager eigenStrategy incorrect"
            );
            assertTrue(
                allocationManager.pauserRegistry() == Env.impl.pauserRegistry(),
                "allocationManager pauserRegistry incorrect"
            );
            assertTrue(
                allocationManager.permissionController() == Env.proxy.permissionController(),
                "allocationManager permissionController incorrect"
            );
            assertTrue(
                allocationManager.DEALLOCATION_DELAY() == Env.MIN_WITHDRAWAL_DELAY(),
                "allocationManager DEALLOCATION_DELAY incorrect"
            );
            assertTrue(
                allocationManager.ALLOCATION_CONFIGURATION_DELAY() == Env.ALLOCATION_CONFIGURATION_DELAY(),
                "allocationManager ALLOCATION_CONFIGURATION_DELAY incorrect"
            );

            AVSDirectory avsDirectory = Env.impl.avsDirectory();
            assertTrue(avsDirectory.delegation() == Env.proxy.delegationManager(), "avsDirectory delegation incorrect");
            assertTrue(
                avsDirectory.pauserRegistry() == Env.impl.pauserRegistry(), "avsDirectory pauserRegistry incorrect"
            );

            DelegationManager delegation = Env.impl.delegationManager();
            assertTrue(
                delegation.strategyManager() == Env.proxy.strategyManager(),
                "delegationManager strategyManager incorrect"
            );
            assertTrue(
                delegation.eigenPodManager() == Env.proxy.eigenPodManager(),
                "delegationManager eigenPodManager incorrect"
            );
            assertTrue(
                delegation.allocationManager() == Env.proxy.allocationManager(),
                "delegationManager allocationManager incorrect"
            );
            assertTrue(
                delegation.pauserRegistry() == Env.impl.pauserRegistry(), "delegationManager pauserRegistry incorrect"
            );
            assertTrue(
                delegation.permissionController() == Env.proxy.permissionController(),
                "delegationManager permissionController incorrect"
            );
            assertTrue(
                delegation.minWithdrawalDelayBlocks() == Env.MIN_WITHDRAWAL_DELAY(),
                "delegationManager minWithdrawalDelayBlocks incorrect"
            );

            ReleaseManager releaseManager = Env.impl.releaseManager();
            assertTrue(
                releaseManager.permissionController() == Env.proxy.permissionController(),
                "releaseManager permissionController incorrect"
            );

            RewardsCoordinator rewardsCoordinator = Env.impl.rewardsCoordinator();
            assertTrue(
                rewardsCoordinator.delegationManager() == Env.proxy.delegationManager(),
                "rewardsCoordinator delegationManager incorrect"
            );
            assertTrue(
                rewardsCoordinator.strategyManager() == Env.proxy.strategyManager(),
                "rewardsCoordinator strategyManager incorrect"
            );
            assertTrue(
                rewardsCoordinator.allocationManager() == Env.proxy.allocationManager(),
                "rewardsCoordinator allocationManager incorrect"
            );
            assertTrue(
                rewardsCoordinator.pauserRegistry() == Env.impl.pauserRegistry(),
                "rewardsCoordinator pauserRegistry incorrect"
            );
            assertTrue(
                rewardsCoordinator.permissionController() == Env.proxy.permissionController(),
                "rewardsCoordinator permissionController incorrect"
            );
            assertTrue(
                rewardsCoordinator.CALCULATION_INTERVAL_SECONDS() == Env.CALCULATION_INTERVAL_SECONDS(),
                "rewardsCoordinator CALCULATION_INTERVAL_SECONDS incorrect"
            );
            assertTrue(
                rewardsCoordinator.MAX_REWARDS_DURATION() == Env.MAX_REWARDS_DURATION(),
                "rewardsCoordinator MAX_REWARDS_DURATION incorrect"
            );
            assertTrue(
                rewardsCoordinator.MAX_RETROACTIVE_LENGTH() == Env.MAX_RETROACTIVE_LENGTH(),
                "rewardsCoordinator MAX_RETROACTIVE_LENGTH incorrect"
            );
            assertTrue(
                rewardsCoordinator.MAX_FUTURE_LENGTH() == Env.MAX_FUTURE_LENGTH(),
                "rewardsCoordinator MAX_FUTURE_LENGTH incorrect"
            );
            assertTrue(
                rewardsCoordinator.GENESIS_REWARDS_TIMESTAMP() == Env.GENESIS_REWARDS_TIMESTAMP(),
                "rewardsCoordinator GENESIS_REWARDS_TIMESTAMP incorrect"
            );

            StrategyManager strategyManager = Env.impl.strategyManager();
            assertTrue(
                strategyManager.allocationManager() == Env.proxy.allocationManager(),
                "strategyManager allocationManager incorrect"
            );
            assertTrue(
                strategyManager.delegation() == Env.proxy.delegationManager(), "strategyManager delegation incorrect"
            );
            assertTrue(
                strategyManager.pauserRegistry() == Env.impl.pauserRegistry(),
                "strategyManager pauserRegistry incorrect"
            );
        }

        {
            /// pods/
            EigenPod eigenPod = Env.impl.eigenPod();
            assertTrue(eigenPod.ethPOS() == Env.ethPOS(), "eigenPod ethPOS incorrect");
            assertTrue(eigenPod.eigenPodManager() == Env.proxy.eigenPodManager(), "eigenPod eigenPodManager incorrect");

            EigenPodManager eigenPodManager = Env.impl.eigenPodManager();
            assertTrue(eigenPodManager.ethPOS() == Env.ethPOS(), "eigenPodManager ethPOS incorrect");
            assertTrue(
                eigenPodManager.eigenPodBeacon() == Env.beacon.eigenPod(), "eigenPodManager eigenPodBeacon incorrect"
            );
            assertTrue(
                eigenPodManager.delegationManager() == Env.proxy.delegationManager(),
                "eigenPodManager delegationManager incorrect"
            );
            assertTrue(
                eigenPodManager.pauserRegistry() == Env.impl.pauserRegistry(),
                "eigenPodManager pauserRegistry incorrect"
            );
        }

        {
            /// strategies/
            EigenStrategy eigenStrategy = Env.impl.eigenStrategy();
            assertTrue(
                eigenStrategy.strategyManager() == Env.proxy.strategyManager(),
                "eigenStrategy strategyManager incorrect"
            );
            assertTrue(
                eigenStrategy.pauserRegistry() == Env.impl.pauserRegistry(), "eigenStrategy pauserRegistry incorrect"
            );

            StrategyBase strategyBase = Env.impl.strategyBase();
            assertTrue(strategyBase.strategyManager() == Env.proxy.strategyManager(), "stratBase.sm invalid");
            assertTrue(strategyBase.pauserRegistry() == Env.impl.pauserRegistry(), "stratBase.pR invalid");

            StrategyBaseTVLLimits strategyBaseTVLLimits = Env.impl.strategyBaseTVLLimits();
            assertTrue(
                strategyBaseTVLLimits.strategyManager() == Env.proxy.strategyManager(), "stratBaseTVL.sm invalid"
            );
            assertTrue(strategyBaseTVLLimits.pauserRegistry() == Env.impl.pauserRegistry(), "stratBaseTVL.pR invalid");

            StrategyFactory strategyFactory = Env.impl.strategyFactory();
            assertTrue(strategyFactory.strategyManager() == Env.proxy.strategyManager(), "sFact.sm invalid");
            assertTrue(strategyFactory.pauserRegistry() == Env.impl.pauserRegistry(), "sFact.pR invalid");
        }

        {
            /// multichain/
            CrossChainRegistry crossChainRegistry = Env.impl.crossChainRegistry();
            assertTrue(
                crossChainRegistry.allocationManager() == Env.proxy.allocationManager(),
                "crossChainRegistry.alm invalid"
            );
            assertTrue(crossChainRegistry.keyRegistrar() == Env.proxy.keyRegistrar(), "crossChainRegistry.kr invalid");
            assertTrue(
                crossChainRegistry.permissionController() == Env.proxy.permissionController(),
                "crossChainRegistry.pc invalid"
            );
            assertTrue(
                crossChainRegistry.pauserRegistry() == Env.impl.pauserRegistry(), "crossChainRegistry.pR invalid"
            );

            OperatorTableUpdater operatorTableUpdater = Env.impl.operatorTableUpdater();
            assertTrue(
                operatorTableUpdater.bn254CertificateVerifier() == Env.proxy.bn254CertificateVerifier(),
                "opTable.bn254CV invalid"
            );
            assertTrue(
                operatorTableUpdater.ecdsaCertificateVerifier() == Env.proxy.ecdsaCertificateVerifier(),
                "opTable.ecdsaCV invalid"
            );
            assertTrue(operatorTableUpdater.pauserRegistry() == Env.impl.pauserRegistry(), "opTable.pR invalid");

            ECDSACertificateVerifier ecdsaCertificateVerifier = Env.impl.ecdsaCertificateVerifier();
            assertTrue(
                ecdsaCertificateVerifier.operatorTableUpdater() == Env.proxy.operatorTableUpdater(),
                "ecdsaCV.opTable invalid"
            );

            BN254CertificateVerifier bn254CertificateVerifier = Env.impl.bn254CertificateVerifier();
            assertTrue(
                bn254CertificateVerifier.operatorTableUpdater() == Env.proxy.operatorTableUpdater(),
                "bn254CV.opTable invalid"
            );
        }

        {
            /// avs/
            TaskMailbox taskMailbox = Env.impl.taskMailbox();
            assertTrue(
                taskMailbox.BN254_CERTIFICATE_VERIFIER() == address(Env.proxy.bn254CertificateVerifier()),
                "taskMailbox.bn254CV invalid"
            );
            assertTrue(
                taskMailbox.ECDSA_CERTIFICATE_VERIFIER() == address(Env.proxy.ecdsaCertificateVerifier()),
                "taskMailbox.ecdsaCV invalid"
            );
            assertTrue(taskMailbox.MAX_TASK_SLA() == Env.MAX_TASK_SLA(), "taskMailbox.maxTaskSLA invalid");
        }
    }

    function validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// permissions/
        // KeyRegistrar and PermissionController are initializable, but do not expose the `initialize` function.

        {
            /// core/
            AllocationManager allocationManager = Env.impl.allocationManager();
            vm.expectRevert(errInit);
            allocationManager.initialize(0);

            AVSDirectory avsDirectory = Env.impl.avsDirectory();
            vm.expectRevert(errInit);
            avsDirectory.initialize(address(0), 0);

            DelegationManager delegationManager = Env.impl.delegationManager();
            vm.expectRevert(errInit);
            delegationManager.initialize(0);

            ProtocolRegistry protocolRegistry = Env.impl.protocolRegistry();
            vm.expectRevert(errInit);
            protocolRegistry.initialize(address(0), address(0));

            // ReleaseManager is initializable, but does not expose the `initialize` function.

            RewardsCoordinator rewardsCoordinator = Env.impl.rewardsCoordinator();
            vm.expectRevert(errInit);
            rewardsCoordinator.initialize(address(0), 0, address(0), 0, 0);

            StrategyManager strategyManager = Env.impl.strategyManager();
            vm.expectRevert(errInit);
            strategyManager.initialize(address(0), address(0), 0);
        }

        {
            /// pods/
            EigenPod eigenPod = Env.impl.eigenPod();
            vm.expectRevert(errInit);
            eigenPod.initialize(address(0));

            EigenPodManager eigenPodManager = Env.impl.eigenPodManager();
            vm.expectRevert(errInit);
            eigenPodManager.initialize(address(0), 0);
        }

        {
            /// strategies/
            EigenStrategy eigenStrategy = Env.impl.eigenStrategy();
            vm.expectRevert(errInit);
            eigenStrategy.initialize(IEigen(address(0)), IBackingEigen(address(0)));

            StrategyBase strategyBase = Env.impl.strategyBase();
            vm.expectRevert(errInit);
            strategyBase.initialize(IERC20(address(0)));

            StrategyBaseTVLLimits strategyBaseTVLLimits = Env.impl.strategyBaseTVLLimits();
            vm.expectRevert(errInit);
            strategyBaseTVLLimits.initialize(0, 0, IERC20(address(0)));

            StrategyFactory strategyFactory = Env.impl.strategyFactory();
            vm.expectRevert(errInit);
            strategyFactory.initialize(address(0), 0, UpgradeableBeacon(address(0)));
        }

        {
            /// multichain/
            CrossChainRegistry crossChainRegistry = Env.impl.crossChainRegistry();
            vm.expectRevert(errInit);
            crossChainRegistry.initialize(address(0), 0, 0);

            OperatorTableUpdater operatorTableUpdater = Env.impl.operatorTableUpdater();
            OperatorSet memory dummyOperatorSet = OperatorSet({avs: address(0), id: 0});
            IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory dummyBN254Info;
            vm.expectRevert(errInit);
            operatorTableUpdater.initialize(address(0), uint256(0), dummyOperatorSet, 0, dummyBN254Info);

            // BN254 and ECDSA certificate verifiers are initializable, but do not expose the `initialize` function.
        }

        {
            /// avs/
            TaskMailbox taskMailbox = Env.impl.taskMailbox();
            vm.expectRevert(errInit);
            taskMailbox.initialize(address(0), 0, address(0));
        }
    }

    /**
     * @notice After the upgrade is complete, call to _validateNewImplAddresses to ensure the impl addresses match the proxy admin's reported implementation address.
     */
    function validateImplAddressesMatchProxy() internal view {
        ///permissions/
        assertTrue(
            _getProxyImpl(address(Env.proxy.permissionController())) == address(Env.impl.permissionController()),
            "permissionController impl address mismatch"
        );
        assertTrue(
            _getProxyImpl(address(Env.proxy.keyRegistrar())) == address(Env.impl.keyRegistrar()),
            "keyRegistrar impl address mismatch"
        );

        ///core/
        assertTrue(
            _getProxyImpl(address(Env.proxy.allocationManager())) == address(Env.impl.allocationManager()),
            "allocationManager impl address mismatch"
        );
        assertTrue(
            _getProxyImpl(address(Env.proxy.avsDirectory())) == address(Env.impl.avsDirectory()),
            "avsDirectory impl address mismatch"
        );
        assertTrue(
            _getProxyImpl(address(Env.proxy.delegationManager())) == address(Env.impl.delegationManager()),
            "delegationManager impl address mismatch"
        );
        assertTrue(
            _getProxyImpl(address(Env.proxy.protocolRegistry())) == address(Env.impl.protocolRegistry()),
            "protocolRegistry impl address mismatch"
        );
        assertTrue(
            _getProxyImpl(address(Env.proxy.releaseManager())) == address(Env.impl.releaseManager()),
            "releaseManager impl address mismatch"
        );
        assertTrue(
            _getProxyImpl(address(Env.proxy.rewardsCoordinator())) == address(Env.impl.rewardsCoordinator()),
            "rewardsCoordinator impl address mismatch"
        );
        assertTrue(
            _getProxyImpl(address(Env.proxy.strategyManager())) == address(Env.impl.strategyManager()),
            "strategyManager impl address mismatch"
        );

        ///pods/
        assertTrue(
            Env.beacon.eigenPod().implementation() == address(Env.impl.eigenPod()), "eigenPod impl address mismatch"
        );
        assertTrue(
            _getProxyImpl(address(Env.proxy.eigenPodManager())) == address(Env.impl.eigenPodManager()),
            "eigenPodManager impl address mismatch"
        );

        ///strategies/
        assertTrue(
            _getProxyImpl(address(Env.proxy.eigenStrategy())) == address(Env.impl.eigenStrategy()),
            "eigenStrategy impl address mismatch"
        );
        assertTrue(
            Env.beacon.strategyBase().implementation() == address(Env.impl.strategyBase()),
            "strategyBase impl address mismatch"
        );
        uint256 count = Env.instance.strategyBaseTVLLimits_Count();
        for (uint256 i = 0; i < count; i++) {
            assertTrue(
                _getProxyImpl(address(Env.instance.strategyBaseTVLLimits(i)))
                    == address(Env.impl.strategyBaseTVLLimits()),
                "strategyBaseTVLLimits impl address mismatch"
            );
        }
        assertTrue(
            _getProxyImpl(address(Env.proxy.strategyFactory())) == address(Env.impl.strategyFactory()),
            "strategyFactory impl address mismatch"
        );

        ///multichain/
        assertTrue(
            _getProxyImpl(address(Env.proxy.operatorTableUpdater())) == address(Env.impl.operatorTableUpdater()),
            "crossChainRegistry impl address mismatch"
        );
        assertTrue(
            _getProxyImpl(address(Env.proxy.crossChainRegistry())) == address(Env.impl.crossChainRegistry()),
            "operatorTableUpdater impl address mismatch"
        );
        assertTrue(
            _getProxyImpl(address(Env.proxy.ecdsaCertificateVerifier())) == address(Env.impl.ecdsaCertificateVerifier()),
            "ecdsaCertificateVerifier impl address mismatch"
        );
        assertTrue(
            _getProxyImpl(address(Env.proxy.bn254CertificateVerifier())) == address(Env.impl.bn254CertificateVerifier()),
            "bn254CertificateVerifier impl address mismatch"
        );

        ///avs/
        assertTrue(
            _getProxyImpl(address(Env.proxy.taskMailbox())) == address(Env.impl.taskMailbox()),
            "taskMailbox impl address mismatch"
        );
    }

    /**
     *
     *                        VERSION VALIDATION FUNCTIONS
     *
     */

    /// @dev Validate versions of entire protocol
    /// @dev This should be called on every release
    function validateProtocolVersion() internal view {
        require(
            TestUtils._strEq(Env.proxy.protocolRegistry().version(), Env.deployVersion()), "protocol version incorrect"
        );
    }

    /// Validate versions of specific contracts
    /// @dev We need to validate versions of specific contracts because some contracts can have mismatched versions.
    function validateKeyRegistrarVersion() internal view {
        require(_strEq(Env.impl.keyRegistrar().version(), Env.deployVersion()), "keyRegistrar version mismatch");
    }

    function validateAVSDirectoryVersion() internal view {
        require(_strEq(Env.impl.avsDirectory().version(), Env.deployVersion()), "avsDirectory version incorrect");
    }

    function validateDelegationManagerVersion() internal view {
        require(
            _strEq(Env.impl.delegationManager().version(), Env.deployVersion()), "delegationManager version incorrect"
        );
    }

    function validateStrategyManagerVersion() internal view {
        require(_strEq(Env.impl.strategyManager().version(), Env.deployVersion()), "strategyManager version incorrect");
    }

    function validateECDSACertificateVerifierVersion() internal view {
        require(
            _strEq(Env.impl.ecdsaCertificateVerifier().version(), Env.deployVersion()),
            "ecdsaCertificateVerifier version incorrect"
        );
    }

    /// @dev Query and return `proxyAdmin.getProxyImplementation(proxy)`
    function _getProxyImpl(
        address proxy
    ) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyImplementation(ITransparentUpgradeableProxy(proxy));
    }

    /// @dev Query and return `proxyAdmin.getProxyAdmin(proxy)`
    function _getProxyAdmin(
        address proxy
    ) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyAdmin(ITransparentUpgradeableProxy(proxy));
    }

    function _strEq(string memory a, string memory b) internal pure returns (bool) {
        return keccak256(abi.encodePacked(a)) == keccak256(abi.encodePacked(b));
    }

    function assertTrue(bool b, string memory err) private pure {
        vm.assertTrue(b, err);
    }

    function assertFalse(bool b, string memory err) private pure {
        vm.assertFalse(b, err);
    }
}
