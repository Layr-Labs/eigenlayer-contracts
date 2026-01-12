// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./Env.sol";
import "forge-std/Vm.sol";
import "src/contracts/mixins/SplitContractMixin.sol";

/// @notice Utility library for testing contract deployments
/// @dev This library exposes the following test functions:
/// @dev - validateProxyAdmins/validateDestinationProxyAdmins - Check that proxy admins are correctly set.
/// @dev - validateProxyConstructors/validateDestinationProxyConstructors - Check that proxy constructors are correctly set.
/// @dev - validateProxiesAlreadyInitialized/validateDestinationProxiesAlreadyInitialized - Check that proxies are already initialized.
/// @dev - validateProxyStorage/validateDestinationProxyStorage - Check that proxy storage is correctly set.
/// @dev - validateImplConstructors/validateDestinationImplConstructors - Check that implementation constructors are correctly set.
/// @dev - validateImplsNotInitializable/validateDestinationImplsNotInitializable - Check that implementation cannot be initialized.
/// @dev - validateImplAddressesMatchProxy/validateDestinationImplAddressesMatchProxy - Check that implementation addresses match the proxy admin's reported implementation address.
/// @dev - validateProtocolRegistry/validateDestinationProtocolRegistry - Check that the protocol version is correctly set.
library TestUtils {
    using Env for *;

    bytes constant errInit = "Initializable: contract is already initialized";
    address internal constant VM_ADDRESS = address(uint160(uint256(keccak256("hevm cheat code"))));
    Vm internal constant vm = Vm(VM_ADDRESS);

    ///
    ///                         PROXY ADMIN VALIDATION
    ///

    /// @dev This function is run on *all* deployed contracts to ensure that the proxyAdmin is correctly set.
    function validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();
        /// pemissions/
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
        // Protocol registry tested in the `validateDestinationProxyAdmins` function
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

        validateDestinationProxyAdmins();
    }

    function validateDestinationProxyAdmins() internal view {
        address pa = Env.proxyAdmin();
        /// core/
        assertTrue(_getProxyAdmin(address(Env.proxy.protocolRegistry())) == pa, "protocolRegistry proxyAdmin incorrect");

        /// multichain/
        assertTrue(
            _getProxyAdmin(address(Env.proxy.bn254CertificateVerifier())) == pa,
            "bn254CertificateVerifier proxyAdmin incorrect"
        );
        assertTrue(
            _getProxyAdmin(address(Env.proxy.ecdsaCertificateVerifier())) == pa,
            "ecdsaCertificateVerifier proxyAdmin incorrect"
        );
        assertTrue(
            _getProxyAdmin(address(Env.proxy.operatorTableUpdater())) == pa, "operatorTableUpdater proxyAdmin incorrect"
        );

        /// avs/
        assertTrue(_getProxyAdmin(address(Env.proxy.taskMailbox())) == pa, "taskMailbox proxyAdmin incorrect");
    }

    ///
    ///                         PROXY VALIDATION
    ///

    /// @dev Validate that the proxy constructors are correctly set
    function validateProxyConstructors() internal view {
        /// pemissions/
        // PermissionController has no constructor
        validateKeyRegistrarImmutables(Env.proxy.keyRegistrar());

        /// core/
        validateAllocationManagerImmutables(Env.proxy.allocationManager());
        validateAVSDirectoryImmutables(Env.proxy.avsDirectory());
        validateDelegationManagerImmutables(Env.proxy.delegationManager());
        // ProtocolRegistry has no constructor
        validateReleaseManagerImmutables(Env.proxy.releaseManager());
        validateRewardsCoordinatorImmutables(Env.proxy.rewardsCoordinator());
        validateStrategyManagerImmutables(Env.proxy.strategyManager());

        /// pods/
        // EigenPod beacon doesn't have immutable, only implementation does
        validateEigenPodManagerImmutables(Env.proxy.eigenPodManager());

        /// strategies/
        validateEigenStrategyImmutables(Env.proxy.eigenStrategy());
        // StrategyBase beacon doesn't have immutables, only implementation does
        uint256 count = Env.instance.strategyBaseTVLLimits_Count();
        for (uint256 i = 0; i < count; i++) {
            validateStrategyBaseTVLLimitsImmutables(Env.instance.strategyBaseTVLLimits(i));
        }
        validateStrategyFactoryImmutables(Env.proxy.strategyFactory());

        /// multichain/
        validateCrossChainRegistryImmutables(Env.proxy.crossChainRegistry());

        validateDestinationProxyConstructors();
    }

    function validateDestinationProxyConstructors() internal view {
        /// multichain/
        validateBN254CertificateVerifierImmutables(Env.proxy.bn254CertificateVerifier());
        validateECDSACertificateVerifierImmutables(Env.proxy.ecdsaCertificateVerifier());
        validateOperatorTableUpdaterImmutables(Env.proxy.operatorTableUpdater());

        /// avs/
        validateTaskMailboxImmutables(Env.proxy.taskMailbox());
    }

    /// @dev Validate that the proxies are already initialized.
    function validateProxiesAlreadyInitialized() internal {
        /// pemissions/
        // KeyRegistrar and PermissionController are initializable, but do not expose the `initialize` function.

        /// core/
        validateAllocationManagerInitialized(Env.proxy.allocationManager());
        validateAVSDirectoryInitialized(Env.proxy.avsDirectory());
        validateDelegationManagerInitialized(Env.proxy.delegationManager());
        // Protocol registry tested in the `validateDestinationProxiesAlreadyInitialized` function
        // ReleaseManager is initializable, but does not expose the `initialize` function.
        validateRewardsCoordinatorInitialized(Env.proxy.rewardsCoordinator());
        validateStrategyManagerInitialized(Env.proxy.strategyManager());

        /// pods/
        // EigenPod proxies are initialized by individual users
        validateEigenPodManagerInitialized(Env.proxy.eigenPodManager());

        /// strategies/
        validateEigenStrategyInitialized(Env.proxy.eigenStrategy());
        // StrategyBase proxies are initialized when deployed by factory
        uint256 count = Env.instance.strategyBaseTVLLimits_Count();
        for (uint256 i = 0; i < count; i++) {
            validateStrategyBaseTVLLimitsInitialized(Env.instance.strategyBaseTVLLimits(i));
        }
        validateStrategyFactoryInitialized(Env.proxy.strategyFactory());

        /// multichain/
        validateCrossChainRegistryInitialized(Env.proxy.crossChainRegistry());

        validateDestinationProxiesAlreadyInitialized();
    }

    function validateDestinationProxiesAlreadyInitialized() internal {
        /// core/
        validateProtocolRegistryInitialized(Env.proxy.protocolRegistry());

        /// multichain/
        validateOperatorTableUpdaterInitialized(Env.proxy.operatorTableUpdater());
        // BN254 and ECDSA certificate verifiers are initializable, but do not expose the `initialize` function.

        /// avs/
        validateTaskMailboxInitialized(Env.proxy.taskMailbox());
    }

    function validateProxyStorage() internal view {
        {
        /// permissions/

        // PauserRegistry is also deployed on destination chain and tested in the `validateDestinationProxyStorage` function
        // PermissionController and KeyRegistrar have no initial storage
        }

        {
            /// core/
            AllocationManager allocationManager = Env.proxy.allocationManager();
            assertTrue(allocationManager.paused() == 0, "alm.paused invalid");

            AVSDirectory avsDirectory = Env.proxy.avsDirectory();
            assertTrue(avsDirectory.owner() == Env.executorMultisig(), "avsD.owner invalid");
            assertTrue(avsDirectory.paused() == 0, "avsD.paused invalid");

            DelegationManager delegation = Env.proxy.delegationManager();
            assertTrue(delegation.paused() == 0, "dm.paused invalid");

            // Protocol registry tested in the `validateDestinationProxyStorage` function

            // ReleaseManager has no initial storage

            RewardsCoordinator rewards = Env.proxy.rewardsCoordinator();
            assertTrue(rewards.owner() == Env.opsMultisig(), "rc.owner invalid");
            assertTrue(rewards.paused() == Env.REWARDS_PAUSE_STATUS(), "rc.paused invalid");
            assertTrue(rewards.rewardsUpdater() == Env.REWARDS_UPDATER(), "rc.updater invalid");
            assertTrue(rewards.activationDelay() == Env.ACTIVATION_DELAY(), "rc.activationDelay invalid");
            assertTrue(rewards.defaultOperatorSplitBips() == Env.DEFAULT_SPLIT_BIPS(), "rc.splitBips invalid");

            StrategyManager strategyManager = Env.proxy.strategyManager();
            assertTrue(strategyManager.owner() == Env.executorMultisig(), "sm.owner invalid");
            assertTrue(strategyManager.paused() == 0, "sm.paused invalid");
            assertTrue(
                strategyManager.strategyWhitelister() == address(Env.proxy.strategyFactory()), "sm.whitelister invalid"
            );
        }

        {
            /// pods/
            // EigenPod proxies are initialized by individual users

            EigenPodManager eigenPodManager = Env.proxy.eigenPodManager();
            assertTrue(eigenPodManager.owner() == Env.executorMultisig(), "epm.owner invalid");
            if (Env.supportsEigenPods()) {
                assertTrue(eigenPodManager.paused() == 0, "epm.paused invalid");
            }
        }

        {
            /// strategies/
            EigenStrategy eigenStrategy = Env.proxy.eigenStrategy();
            assertTrue(eigenStrategy.paused() == 0, "eigenStrat.paused invalid");
            assertTrue(address(eigenStrategy.EIGEN()) == address(Env.proxy.eigen()), "eigenStrat.EIGEN invalid");
            assertTrue(
                address(eigenStrategy.underlyingToken()) == address(Env.proxy.beigen()), "eigenStrat.underlying invalid"
            );

            // StrategyBase proxies are initialized when deployed by factory

            uint256 count = Env.instance.strategyBaseTVLLimits_Count();
            for (uint256 i = 0; i < count; i++) {
                StrategyBaseTVLLimits strategy = Env.instance.strategyBaseTVLLimits(i);

                assertTrue(strategy.maxPerDeposit() == type(uint256).max, "stratTVLLim.maxPerDeposit invalid");
                assertTrue(strategy.maxTotalDeposits() == type(uint256).max, "stratTVLLim.maxPerDeposit invalid");
            }

            StrategyFactory strategyFactory = Env.proxy.strategyFactory();
            assertTrue(strategyFactory.owner() == Env.opsMultisig(), "sFact.owner invalid");
            assertTrue(strategyFactory.paused() == 0, "sFact.paused invalid");
            assertTrue(strategyFactory.strategyBeacon() == Env.beacon.strategyBase(), "sFact.beacon invalid");
        }
        {
            /// multichain/
            // Operator table updater and certificate verifies do not have initial storage
            CrossChainRegistry crossChainRegistry = Env.proxy.crossChainRegistry();
            assertTrue(crossChainRegistry.owner() == Env.opsMultisig(), "crossChainRegistry owner invalid");
            assertTrue(
                crossChainRegistry.getTableUpdateCadence() == Env.TABLE_UPDATE_CADENCE(),
                "crossChainRegistry table update cadence invalid"
            );
            if (Env.isSource()) {
                assertTrue(crossChainRegistry.paused() == 0, "crossChainRegistry paused invalid");
            } else {
                // For hoodi, the paused status is 31 (all flags set)
                assertTrue(crossChainRegistry.paused() == 31, "crossChainRegistry paused invalid");
            }
        }

        validateDestinationProxyStorage();
    }

    function validateDestinationProxyStorage() internal view {
        {
            /// permissions/
            // PauserRegistry is also deployed on destination chain
            PauserRegistry registry = Env.impl.pauserRegistry();
            assertTrue(registry.isPauser(Env.pauserMultisig()), "pauser multisig should be pauser");
            assertTrue(registry.isPauser(Env.opsMultisig()), "ops multisig should be pauser");
            assertTrue(registry.isPauser(address(Env.proxy.protocolRegistry())), "protocol registry should be pauser");
            if (Env.isCoreProtocolDeployed()) {
                assertTrue(registry.isPauser(Env.executorMultisig()), "executor multisig should be pauser");
                assertTrue(registry.unpauser() == Env.executorMultisig(), "executor multisig should be unpauser");
            } else {
                assertTrue(registry.unpauser() == Env.opsMultisig(), "ops multisig should be unpauser");
            }
        }

        {
            /// core/
            ProtocolRegistry protocolRegistry = Env.proxy.protocolRegistry();
            assertTrue(
                protocolRegistry.hasRole(protocolRegistry.PAUSER_ROLE(), Env.pauserMultisig()),
                "pr.pauserMultisig invalid"
            );
            assertTrue(
                protocolRegistry.hasRole(protocolRegistry.DEFAULT_ADMIN_ROLE(), Env.executorMultisig()),
                "pr.defaultAdmin invalid"
            );
        }

        {
            /// multichain/
            OperatorTableUpdater operatorTableUpdater = Env.proxy.operatorTableUpdater();
            assertTrue(operatorTableUpdater.owner() == Env.opsMultisig(), "operatorTableUpdater owner invalid");
            assertTrue(
                operatorTableUpdater.globalRootConfirmationThreshold() == 10_000,
                "operatorTableUpdater globalRootConfirmationThreshold invalid"
            );
            // Reset of params are dependent on per chain state (eg. generator state)
        }

        {
            /// avs/
            TaskMailbox taskMailbox = Env.proxy.taskMailbox();
            assertTrue(taskMailbox.owner() == Env.opsMultisig(), "taskMailbox owner invalid");
            assertTrue(taskMailbox.feeSplit() == 0, "taskMailbox feeSplit invalid");
            assertTrue(taskMailbox.feeSplitCollector() == Env.opsMultisig(), "taskMailbox feeSplitCollector invalid");
        }
    }

    ///
    ///                        IMPLEMENTATION VALIDATION
    ///

    /// @dev Validate that the implementation constructors are correctly set
    function validateImplConstructors() internal view {
        /// pemissions/
        // PermissionController has no constructor
        validateKeyRegistrarImmutables(Env.impl.keyRegistrar());

        /// core/
        validateAllocationManagerViewImmutables(Env.impl.allocationManagerView());
        validateAllocationManagerImmutables(Env.impl.allocationManager());
        validateAVSDirectoryImmutables(Env.impl.avsDirectory());
        validateDelegationManagerImmutables(Env.impl.delegationManager());
        // ProtocolRegistry has no constructor
        validateReleaseManagerImmutables(Env.impl.releaseManager());
        validateRewardsCoordinatorImmutables(Env.impl.rewardsCoordinator());
        validateStrategyManagerImmutables(Env.impl.strategyManager());

        /// pods/
        validateEigenPodImmutables(Env.impl.eigenPod());
        validateEigenPodManagerImmutables(Env.impl.eigenPodManager());

        /// strategies/
        validateEigenStrategyImmutables(Env.impl.eigenStrategy());
        validateStrategyBaseImmutables(Env.impl.strategyBase());
        validateStrategyBaseTVLLimitsImmutables(Env.impl.strategyBaseTVLLimits());
        validateStrategyFactoryImmutables(Env.impl.strategyFactory());

        /// multichain/
        validateCrossChainRegistryImmutables(Env.impl.crossChainRegistry());

        validateDestinationImplConstructors();
    }

    function validateDestinationImplConstructors() internal view {
        /// multichain/
        validateBN254CertificateVerifierImmutables(Env.impl.bn254CertificateVerifier());
        validateECDSACertificateVerifierImmutables(Env.impl.ecdsaCertificateVerifier());
        validateOperatorTableUpdaterImmutables(Env.impl.operatorTableUpdater());

        /// avs/
        validateTaskMailboxImmutables(Env.impl.taskMailbox());
    }

    /// @dev Validate that the implementation contracts are not initializable.
    /// @dev Each function checks that initializing the contract will revert.
    function validateImplsNotInitializable() internal {
        /// pemissions/
        // KeyRegistrar and PermissionController are initializable, but do not expose the `initialize` function.

        /// core/
        // AllocationManagerView is initializable, but does not expose the `initialize` function.
        validateAllocationManagerInitialized(Env.impl.allocationManager());
        validateAVSDirectoryInitialized(Env.impl.avsDirectory());
        validateDelegationManagerInitialized(Env.impl.delegationManager());
        // Protocol registry tested in the `validateDestinationImplsNotInitializable` function
        // ReleaseManager is initializable, but does not expose the `initialize` function.
        validateRewardsCoordinatorInitialized(Env.impl.rewardsCoordinator());
        validateStrategyManagerInitialized(Env.impl.strategyManager());

        /// pods/
        // EigenPod implementations are initialized by individual users
        validateEigenPodManagerInitialized(Env.impl.eigenPodManager());

        /// strategies/
        validateEigenStrategyInitialized(Env.impl.eigenStrategy());
        // StrategyBase implementations are initialized when deployed by factory
        validateStrategyBaseTVLLimitsInitialized(Env.impl.strategyBaseTVLLimits());
        validateStrategyFactoryInitialized(Env.impl.strategyFactory());

        /// multichain/
        validateCrossChainRegistryInitialized(Env.impl.crossChainRegistry());

        validateDestinationImplsNotInitializable();
    }

    function validateDestinationImplsNotInitializable() internal {
        /// core/
        validateProtocolRegistryInitialized(Env.impl.protocolRegistry());

        /// multichain/
        validateOperatorTableUpdaterInitialized(Env.impl.operatorTableUpdater());
        // BN254 and ECDSA certificate verifiers are initializable, but do not expose the `initialize` function.

        /// avs/
        validateTaskMailboxInitialized(Env.impl.taskMailbox());
    }

    /// @notice After the upgrade is complete, call to _validateNewImplAddresses to ensure the impl addresses match the proxy admin's reported implementation address.
    function validateImplAddressesMatchProxy() internal view {
        /// pemissions/
        assertTrue(
            _getProxyImpl(address(Env.proxy.permissionController())) == address(Env.impl.permissionController()),
            "permissionController impl address mismatch"
        );
        assertTrue(
            _getProxyImpl(address(Env.proxy.keyRegistrar())) == address(Env.impl.keyRegistrar()),
            "keyRegistrar impl address mismatch"
        );

        /// core/
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
        // Protocol registry tested in the `validateDestinationImplAddressesMatchProxy` function
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

        /// pods/
        assertTrue(
            Env.beacon.eigenPod().implementation() == address(Env.impl.eigenPod()), "eigenPod impl address mismatch"
        );
        assertTrue(
            _getProxyImpl(address(Env.proxy.eigenPodManager())) == address(Env.impl.eigenPodManager()),
            "eigenPodManager impl address mismatch"
        );

        /// strategies/
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

        /// multichain/
        assertTrue(
            _getProxyImpl(address(Env.proxy.crossChainRegistry())) == address(Env.impl.crossChainRegistry()),
            "operatorTableUpdater impl address mismatch"
        );

        validateDestinationImplAddressesMatchProxy();
    }

    function validateDestinationImplAddressesMatchProxy() internal view {
        /// core/
        /// @dev Skip the protocol registry validation as it will be deployed while the upgrade is in progress.
        /// TODO: Uncomment this after v1.9.0 is live
        // assertTrue(
        //     _getProxyImpl(address(Env.proxy.protocolRegistry())) == address(Env.impl.protocolRegistry()),
        //     "protocolRegistry impl address mismatch"
        // );

        /// multichain/
        assertTrue(
            _getProxyImpl(address(Env.proxy.bn254CertificateVerifier()))
                == address(Env.impl.bn254CertificateVerifier()),
            "bn254CertificateVerifier impl address mismatch"
        );
        assertTrue(
            _getProxyImpl(address(Env.proxy.ecdsaCertificateVerifier()))
                == address(Env.impl.ecdsaCertificateVerifier()),
            "ecdsaCertificateVerifier impl address mismatch"
        );
        assertTrue(
            _getProxyImpl(address(Env.proxy.operatorTableUpdater())) == address(Env.impl.operatorTableUpdater()),
            "crossChainRegistry impl address mismatch"
        );

        /// avs/
        assertTrue(
            _getProxyImpl(address(Env.proxy.taskMailbox())) == address(Env.impl.taskMailbox()),
            "taskMailbox impl address mismatch"
        );
    }

    ///
    ///                        VERSION VALIDATION FUNCTIONS
    ///
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

    ///
    ///                         VALIDATE IMMUTABLES
    ///
    /// @dev These functions are used to validate the immutables of either proxy or implementation contracts.

    /// pemissions/
    function validateKeyRegistrarImmutables(
        KeyRegistrar keyRegistrar
    ) internal view {
        assertTrue(
            keyRegistrar.permissionController() == Env.proxy.permissionController(),
            "keyRegistrar permissionController incorrect"
        );
        assertTrue(
            keyRegistrar.allocationManager() == Env.proxy.allocationManager(),
            "keyRegistrar allocationManager incorrect"
        );
    }

    // PermissionController has no immutables

    /// core/
    function validateAllocationManagerViewImmutables(
        AllocationManagerView allocationManagerView
    ) internal view {
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
        assertTrue(
            allocationManagerView.SLASHER_CONFIGURATION_DELAY() == Env.ALLOCATION_CONFIGURATION_DELAY(),
            "allocationManagerView SLASHER_CONFIGURATION_DELAY incorrect"
        );
    }

    function validateAllocationManagerImmutables(
        AllocationManager allocationManager
    ) internal view {
        assertTrue(
            allocationManager.viewImplementation() == address(Env.impl.allocationManagerView()),
            "allocationManager allocationManagerView incorrect"
        );
        assertTrue(
            allocationManager.delegation() == Env.proxy.delegationManager(), "allocationManager delegation incorrect"
        );
        assertTrue(
            allocationManager.eigenStrategy() == Env.proxy.eigenStrategy(), "allocationManager eigenStrategy incorrect"
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
        assertTrue(
            allocationManager.SLASHER_CONFIGURATION_DELAY() == Env.ALLOCATION_CONFIGURATION_DELAY(),
            "allocationManager SLASHER_CONFIGURATION_DELAY incorrect"
        );
    }

    function validateAVSDirectoryImmutables(
        AVSDirectory avsDirectory
    ) internal view {
        assertTrue(avsDirectory.delegation() == Env.proxy.delegationManager(), "avsDirectory delegation incorrect");
        assertTrue(avsDirectory.pauserRegistry() == Env.impl.pauserRegistry(), "avsDirectory pauserRegistry incorrect");
    }

    function validateDelegationManagerImmutables(
        DelegationManager delegation
    ) internal view {
        assertTrue(
            delegation.strategyManager() == Env.proxy.strategyManager(), "delegationManager strategyManager incorrect"
        );
        assertTrue(
            delegation.eigenPodManager() == Env.proxy.eigenPodManager(), "delegationManager eigenPodManager incorrect"
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
    }

    function validateReleaseManagerImmutables(
        ReleaseManager releaseManager
    ) internal view {
        assertTrue(
            releaseManager.permissionController() == Env.proxy.permissionController(),
            "releaseManager permissionController incorrect"
        );
    }

    function validateRewardsCoordinatorImmutables(
        RewardsCoordinator rewardsCoordinator
    ) internal view {
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
    }

    function validateStrategyManagerImmutables(
        StrategyManager strategyManager
    ) internal view {
        assertTrue(
            strategyManager.allocationManager() == Env.proxy.allocationManager(),
            "strategyManager allocationManager incorrect"
        );
        assertTrue(
            strategyManager.delegation() == Env.proxy.delegationManager(), "strategyManager delegation incorrect"
        );
        assertTrue(
            strategyManager.pauserRegistry() == Env.impl.pauserRegistry(), "strategyManager pauserRegistry incorrect"
        );
    }

    /// pods/
    function validateEigenPodImmutables(
        EigenPod eigenPod
    ) internal view {
        assertTrue(eigenPod.ethPOS() == Env.ethPOS(), "eigenPod ethPOS incorrect");
        assertTrue(eigenPod.eigenPodManager() == Env.proxy.eigenPodManager(), "eigenPod eigenPodManager incorrect");
    }

    function validateEigenPodManagerImmutables(
        EigenPodManager eigenPodManager
    ) internal view {
        assertTrue(eigenPodManager.ethPOS() == Env.ethPOS(), "eigenPodManager ethPOS incorrect");
        assertTrue(
            eigenPodManager.eigenPodBeacon() == Env.beacon.eigenPod(), "eigenPodManager eigenPodBeacon incorrect"
        );
        assertTrue(
            eigenPodManager.delegationManager() == Env.proxy.delegationManager(),
            "eigenPodManager delegationManager incorrect"
        );
        assertTrue(
            eigenPodManager.pauserRegistry() == Env.impl.pauserRegistry(), "eigenPodManager pauserRegistry incorrect"
        );
    }

    /// strategies/
    function validateEigenStrategyImmutables(
        EigenStrategy eigenStrategy
    ) internal view {
        assertTrue(
            eigenStrategy.strategyManager() == Env.proxy.strategyManager(), "eigenStrategy strategyManager incorrect"
        );
        assertTrue(
            eigenStrategy.pauserRegistry() == Env.impl.pauserRegistry(), "eigenStrategy pauserRegistry incorrect"
        );
    }

    function validateStrategyBaseImmutables(
        StrategyBase strategyBase
    ) internal view {
        assertTrue(
            strategyBase.strategyManager() == Env.proxy.strategyManager(), "strategyBase strategyManager incorrect"
        );
        assertTrue(strategyBase.pauserRegistry() == Env.impl.pauserRegistry(), "strategyBase pauserRegistry incorrect");
    }

    function validateStrategyBaseTVLLimitsImmutables(
        StrategyBaseTVLLimits strategyBaseTVLLimits
    ) internal view {
        assertTrue(
            strategyBaseTVLLimits.strategyManager() == Env.proxy.strategyManager(),
            "strategyBaseTVLLimits strategyManager incorrect"
        );
        assertTrue(
            strategyBaseTVLLimits.pauserRegistry() == Env.impl.pauserRegistry(),
            "strategyBaseTVLLimits pauserRegistry incorrect"
        );
    }

    function validateStrategyFactoryImmutables(
        StrategyFactory strategyFactory
    ) internal view {
        assertTrue(
            strategyFactory.strategyManager() == Env.proxy.strategyManager(),
            "strategyFactory strategyManager incorrect"
        );
        assertTrue(
            strategyFactory.pauserRegistry() == Env.impl.pauserRegistry(), "strategyFactory pauserRegistry incorrect"
        );
    }

    /// multichain/
    function validateBN254CertificateVerifierImmutables(
        BN254CertificateVerifier bn254CertificateVerifier
    ) internal view {
        assertTrue(
            bn254CertificateVerifier.operatorTableUpdater() == Env.proxy.operatorTableUpdater(),
            "bn254CertificateVerifier operatorTableUpdater incorrect"
        );
    }

    function validateCrossChainRegistryImmutables(
        CrossChainRegistry crossChainRegistry
    ) internal view {
        assertTrue(
            crossChainRegistry.allocationManager() == Env.proxy.allocationManager(),
            "crossChainRegistry allocationManager incorrect"
        );
        assertTrue(
            crossChainRegistry.keyRegistrar() == Env.proxy.keyRegistrar(), "crossChainRegistry keyRegistrar incorrect"
        );
        assertTrue(
            crossChainRegistry.permissionController() == Env.proxy.permissionController(),
            "crossChainRegistry permissionController incorrect"
        );
        assertTrue(
            crossChainRegistry.pauserRegistry() == Env.impl.pauserRegistry(),
            "crossChainRegistry pauserRegistry incorrect"
        );
    }

    function validateECDSACertificateVerifierImmutables(
        ECDSACertificateVerifier ecdsaCertificateVerifier
    ) internal view {
        assertTrue(
            ecdsaCertificateVerifier.operatorTableUpdater() == Env.proxy.operatorTableUpdater(),
            "ecdsaCertificateVerifier operatorTableUpdater incorrect"
        );
    }

    function validateOperatorTableUpdaterImmutables(
        OperatorTableUpdater operatorTableUpdater
    ) internal view {
        assertTrue(
            operatorTableUpdater.bn254CertificateVerifier() == Env.proxy.bn254CertificateVerifier(),
            "operatorTableUpdater bn254CertificateVerifier incorrect"
        );
        assertTrue(
            operatorTableUpdater.ecdsaCertificateVerifier() == Env.proxy.ecdsaCertificateVerifier(),
            "operatorTableUpdater ecdsaCertificateVerifier incorrect"
        );
        assertTrue(
            operatorTableUpdater.pauserRegistry() == Env.impl.pauserRegistry(),
            "operatorTableUpdater pauserRegistry incorrect"
        );
    }

    /// avs/
    function validateTaskMailboxImmutables(
        TaskMailbox taskMailbox
    ) internal view {
        assertTrue(
            taskMailbox.BN254_CERTIFICATE_VERIFIER() == address(Env.proxy.bn254CertificateVerifier()),
            "taskMailbox BN254_CERTIFICATE_VERIFIER incorrect"
        );
        assertTrue(
            taskMailbox.ECDSA_CERTIFICATE_VERIFIER() == address(Env.proxy.ecdsaCertificateVerifier()),
            "taskMailbox ECDSA_CERTIFICATE_VERIFIER incorrect"
        );
        assertTrue(taskMailbox.MAX_TASK_SLA() == Env.MAX_TASK_SLA(), "taskMailbox MAX_TASK_SLA incorrect");
    }

    ///
    ///                         VALIDATE INITIALIZED
    ///

    /// @dev These functions are used to validate the initialized state of either proxy or implementation contracts.

    /// pemissions/
    // KeyRegistrar and PermissionController are initializable, but do not expose the `initialize` function.

    /// core/
    function validateAllocationManagerInitialized(
        AllocationManager allocationManager
    ) internal {
        vm.expectRevert(errInit);
        allocationManager.initialize(0);
    }

    function validateAVSDirectoryInitialized(
        AVSDirectory avsDirectory
    ) internal {
        vm.expectRevert(errInit);
        avsDirectory.initialize(address(0), 0);
    }

    function validateDelegationManagerInitialized(
        DelegationManager delegationManager
    ) internal {
        vm.expectRevert(errInit);
        delegationManager.initialize(0);
    }

    function validateProtocolRegistryInitialized(
        ProtocolRegistry protocolRegistry
    ) internal {
        vm.expectRevert(errInit);
        protocolRegistry.initialize(address(0), address(0));
    }

    // ReleaseManager is initializable, but does not expose the `initialize` function.

    function validateRewardsCoordinatorInitialized(
        RewardsCoordinator rewardsCoordinator
    ) internal {
        vm.expectRevert(errInit);
        rewardsCoordinator.initialize(address(0), 0, address(0), 0, 0, address(0));
    }

    function validateStrategyManagerInitialized(
        StrategyManager strategyManager
    ) internal {
        vm.expectRevert(errInit);
        strategyManager.initialize(address(0), address(0), 0);
    }

    /// pods/
    // EigenPod proxies are initialized by individual users

    function validateEigenPodManagerInitialized(
        EigenPodManager eigenPodManager
    ) internal {
        vm.expectRevert(errInit);
        eigenPodManager.initialize(address(0), 0);
    }

    /// strategies/
    function validateEigenStrategyInitialized(
        EigenStrategy eigenStrategy
    ) internal {
        vm.expectRevert(errInit);
        eigenStrategy.initialize(IEigen(address(0)), IBackingEigen(address(0)));
    }

    // StrategyBase proxies are initialized when deployed by factory

    function validateStrategyBaseTVLLimitsInitialized(
        StrategyBaseTVLLimits strategyBaseTVLLimits
    ) internal {
        vm.expectRevert(errInit);
        strategyBaseTVLLimits.initialize(0, 0, IERC20(address(0)));
    }

    function validateStrategyFactoryInitialized(
        StrategyFactory strategyFactory
    ) internal {
        vm.expectRevert(errInit);
        strategyFactory.initialize(address(0), 0, UpgradeableBeacon(address(0)));
    }

    /// multichain/
    function validateCrossChainRegistryInitialized(
        CrossChainRegistry crossChainRegistry
    ) internal {
        vm.expectRevert(errInit);
        crossChainRegistry.initialize(address(0), 0, 0);
    }

    function validateOperatorTableUpdaterInitialized(
        OperatorTableUpdater operatorTableUpdater
    ) internal {
        OperatorSet memory dummyOperatorSet = OperatorSet({avs: address(0), id: 0});
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory dummyBN254Info;
        vm.expectRevert(errInit);
        operatorTableUpdater.initialize(address(0), uint256(0), dummyOperatorSet, 0, dummyBN254Info);
    }

    // BN254 and ECDSA certificate verifiers are initializable, but do not expose the `initialize` function.

    /// avs/
    function validateTaskMailboxInitialized(
        TaskMailbox taskMailbox
    ) internal {
        vm.expectRevert(errInit);
        taskMailbox.initialize(address(0), 0, address(0));
    }

    ///
    ///                         VALIDATE PROTOCOL REGISTRY
    ///

    /// @notice Validate the protocol registry by checking the version and all contracts
    /// @dev This should be called *after* an upgrade has been completed
    function validateProtocolRegistry() internal {
        // Version is checked in the `validateDestinationProtocolRegistry` function

        // Check the deployments
        address addr;
        IProtocolRegistryTypes.DeploymentConfig memory config;
        {
            /// permissions/
            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(KeyRegistrar).name);
            assertTrue(addr == address(Env.proxy.keyRegistrar()), "keyRegistrar address incorrect");
            assertFalse(config.pausable, "keyRegistrar should not be pausable");
            assertFalse(config.deprecated, "keyRegistrar should not be deprecated");

            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(PermissionController).name);
            assertTrue(addr == address(Env.proxy.permissionController()), "permissionController address incorrect");
            assertFalse(config.pausable, "permissionController should not be pausable");
            assertFalse(config.deprecated, "permissionController should not be deprecated");
        }

        {
            /// core/
            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(AllocationManager).name);
            assertTrue(addr == address(Env.proxy.allocationManager()), "allocationManager address incorrect");
            assertTrue(config.pausable, "allocationManager should be pausable");
            assertFalse(config.deprecated, "allocationManager should not be deprecated");

            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(AVSDirectory).name);
            assertTrue(addr == address(Env.proxy.avsDirectory()), "avsDirectory address incorrect");
            assertTrue(config.pausable, "avsDirectory should be pausable");
            assertFalse(config.deprecated, "avsDirectory should not be deprecated");

            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(DelegationManager).name);
            assertTrue(addr == address(Env.proxy.delegationManager()), "delegationManager address incorrect");
            assertTrue(config.pausable, "delegationManager should be pausable");
            assertFalse(config.deprecated, "delegationManager should not be deprecated");

            // Protocol registry tested in the `validateDestinationProtocolRegistry` function

            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(ReleaseManager).name);
            assertTrue(addr == address(Env.proxy.releaseManager()), "releaseManager address incorrect");
            assertFalse(config.pausable, "releaseManager should not be pausable");
            assertFalse(config.deprecated, "releaseManager should not be deprecated");

            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(RewardsCoordinator).name);
            assertTrue(addr == address(Env.proxy.rewardsCoordinator()), "rewardsCoordinator address incorrect");
            assertTrue(config.pausable, "rewardsCoordinator should be pausable");
            assertFalse(config.deprecated, "rewardsCoordinator should not be deprecated");

            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(StrategyManager).name);
            assertTrue(addr == address(Env.proxy.strategyManager()), "strategyManager address incorrect");
            assertTrue(config.pausable, "strategyManager should be pausable");
            assertFalse(config.deprecated, "strategyManager should not be deprecated");
        }

        {
            /// pods/
            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(EigenPodManager).name);
            assertTrue(addr == address(Env.proxy.eigenPodManager()), "eigenPodManager address incorrect");
            assertTrue(config.pausable, "eigenPodManager should be pausable");
            assertFalse(config.deprecated, "eigenPodManager should not be deprecated");

            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(EigenPod).name);
            assertTrue(addr == address(Env.beacon.eigenPod()), "eigenPod address incorrect");
            assertFalse(config.pausable, "eigenPod should not be pausable");
            assertFalse(config.deprecated, "eigenPod should not be deprecated");
        }

        {
            /// strategies/
            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(EigenStrategy).name);
            assertTrue(addr == address(Env.proxy.eigenStrategy()), "eigenStrategy address incorrect");
            assertTrue(config.pausable, "eigenStrategy should be pausable");
            assertFalse(config.deprecated, "eigenStrategy should not be deprecated");

            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(StrategyBase).name);
            assertTrue(addr == address(Env.beacon.strategyBase()), "strategyBase address incorrect");
            assertFalse(config.pausable, "strategyBase should not be pausable");
            assertFalse(config.deprecated, "strategyBase should not be deprecated");

            uint256 count = Env.instance.strategyBaseTVLLimits_Count();
            string memory baseName = type(StrategyBaseTVLLimits).name;
            for (uint256 i = 0; i < count; i++) {
                (addr, config) =
                    Env.proxy.protocolRegistry().getDeployment(string.concat(baseName, "_", Strings.toString(i)));
                assertTrue(
                    addr == address(Env.instance.strategyBaseTVLLimits(i)), "strategyBaseTVLLimits address incorrect"
                );
                assertTrue(config.pausable, "strategyBaseTVLLimits should be pausable");
                assertFalse(config.deprecated, "strategyBaseTVLLimits should not be deprecated");
            }

            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(StrategyFactory).name);
            assertTrue(addr == address(Env.proxy.strategyFactory()), "strategyFactory address incorrect");
            assertTrue(config.pausable, "strategyFactory should be pausable");
            assertFalse(config.deprecated, "strategyFactory should not be deprecated");
        }

        {
            /// token/
            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(BackingEigen).name);
            assertTrue(addr == address(Env.proxy.beigen()), "bEIGEN address incorrect");
            assertFalse(config.pausable, "bEIGEN should not be pausable");
            assertFalse(config.deprecated, "bEIGEN should not be deprecated");

            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(Eigen).name);
            assertTrue(addr == address(Env.proxy.eigen()), "eigenToken address incorrect");
            assertFalse(config.pausable, "eigenToken should not be pausable");
            assertFalse(config.deprecated, "eigenToken should not be deprecated");
        }

        {
            /// multichain/
            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(CrossChainRegistry).name);
            assertTrue(addr == address(Env.proxy.crossChainRegistry()), "crossChainRegistry address incorrect");
            assertTrue(config.pausable, "crossChainRegistry should be pausable");
            assertFalse(config.deprecated, "crossChainRegistry should not be deprecated");
        }

        validateDestinationProtocolRegistry();
    }

    function validateDestinationProtocolRegistry() internal {
        /// First, check the version of the registry
        assertTrue(
            _strEq(Env.proxy.protocolRegistry().version(), Env.deployVersion()), "protocol registry version incorrect"
        );

        // Then, check the deployments
        address addr;
        IProtocolRegistryTypes.DeploymentConfig memory config;

        {
            /// permissions/
            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(PauserRegistry).name);
            assertTrue(addr == address(Env.impl.pauserRegistry()), "pauserRegistry address incorrect");
            assertFalse(config.pausable, "pauserRegistry should not be pausable");
            assertFalse(config.deprecated, "pauserRegistry should not be deprecated");
        }

        {
            /// core/
            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(ProtocolRegistry).name);
            assertTrue(addr == address(Env.proxy.protocolRegistry()), "protocolRegistry address incorrect");
            assertFalse(config.pausable, "protocolRegistry should not be pausable");
            assertFalse(config.deprecated, "protocolRegistry should not be deprecated");
        }

        {
            /// multichain/
            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(BN254CertificateVerifier).name);
            assertTrue(
                addr == address(Env.proxy.bn254CertificateVerifier()), "bn254CertificateVerifier address incorrect"
            );
            assertFalse(config.pausable, "bn254CertificateVerifier should not be pausable");
            assertFalse(config.deprecated, "bn254CertificateVerifier should not be deprecated");

            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(ECDSACertificateVerifier).name);
            assertTrue(
                addr == address(Env.proxy.ecdsaCertificateVerifier()), "ecdsaCertificateVerifier address incorrect"
            );
            assertFalse(config.pausable, "ecdsaCertificateVerifier should not be pausable");
            assertFalse(config.deprecated, "ecdsaCertificateVerifier should not be deprecated");

            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(OperatorTableUpdater).name);
            assertTrue(addr == address(Env.proxy.operatorTableUpdater()), "operatorTableUpdater address incorrect");
            assertTrue(config.pausable, "operatorTableUpdater should not be pausable");
            assertFalse(config.deprecated, "operatorTableUpdater should not be deprecated");
        }

        {
            /// avs/
            (addr, config) = Env.proxy.protocolRegistry().getDeployment(type(TaskMailbox).name);
            assertTrue(addr == address(Env.proxy.taskMailbox()), "taskMailbox address incorrect");
            assertFalse(config.pausable, "taskMailbox should not be pausable");
            assertFalse(config.deprecated, "taskMailbox should not be deprecated");
        }

        // Lastly, attempt to call pauseAll on the protocol registry
        vm.prank(Env.pauserMultisig());
        Env.proxy.protocolRegistry().pauseAll();
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

    function _strEq(
        string memory a,
        string memory b
    ) internal pure returns (bool) {
        return keccak256(abi.encodePacked(a)) == keccak256(abi.encodePacked(b));
    }

    function assertTrue(
        bool b,
        string memory err
    ) private pure {
        vm.assertTrue(b, err);
    }

    function assertFalse(
        bool b,
        string memory err
    ) private pure {
        vm.assertFalse(b, err);
    }
}
