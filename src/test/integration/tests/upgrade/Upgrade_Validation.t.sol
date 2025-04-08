// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/UpgradeTest.t.sol";

contract Integration_ValidateMainnetUpgrade is UpgradeTest {
    
    string constant EXPECTED_VERSION = "1.3.0";

    StrategyBase pepeStrategy = StrategyBase(0x99A05F4e3Fa886A5104630e8a4b01159867ad9a1);
    IERC20 pepeToken = IERC20(0x6982508145454Ce325dDbE47a25d4ec3d2311933);

    /// @notice Test upgrade setup is correct
    function test_mainnet_upgrade_setup() public {
        // 1. Check proper state pre-upgrade
        _verifyContractPointers();
        _verifyImplementations();
        _verifyInitializationParams();
        uint totalSharesBefore = _validateExistingLongtailStrat();

        // 2. Upgrade mainnet contracts
        _upgradeEigenLayerContracts();

        // 3. Verify upgrade setup
        _verifyContractPointers();
        _verifyImplementations();
        _verifyContractsInitialized(false);
        _verifyInitializationParams();

        uint totalSharesAfter = _validateExistingLongtailStrat();
        _validateImplConstructors();
        _validateVersion();

        assertEq(totalSharesBefore, totalSharesAfter, "upgrade share count differs");
    }

    function _validateExistingLongtailStrat() internal view returns (uint totalShares) {
        assertTrue(strategyManager.strategyIsWhitelistedForDeposit(pepeStrategy), "pepe not whitelisted");
        assertTrue(pepeStrategy.pauserRegistry() == eigenLayerPauserReg, "pepe pauser misconfigured");
        assertTrue(pepeStrategy.strategyManager() == strategyManager, "pepe pauser misconfigured");
        assertTrue(pepeStrategy.underlyingToken() == pepeToken, "pepe pauser misconfigured");

        return pepeStrategy.totalShares();
    }

    /// @notice Ensure contracts point at each other correctly via constructors
    /// override to remove ethPOSDeposit contract check
    function _verifyContractPointers() internal view virtual override {
        // AVSDirectory
        require(avsDirectory.delegation() == delegationManager, "avsDirectory: delegationManager address not set correctly");
        // DelegationManager
        require(delegationManager.strategyManager() == strategyManager, "delegationManager: strategyManager address not set correctly");
        require(delegationManager.eigenPodManager() == eigenPodManager, "delegationManager: eigenPodManager address not set correctly");
        // StrategyManager
        require(strategyManager.delegation() == delegationManager, "strategyManager: delegationManager address not set correctly");
        // EPM
        require(eigenPodManager.eigenPodBeacon() == eigenPodBeacon, "eigenPodManager: eigenPodBeacon contract address not set correctly");
        require(
            eigenPodManager.delegationManager() == delegationManager,
            "eigenPodManager: delegationManager contract address not set correctly"
        );
    }

    function _validateImplConstructors() internal view {
        {
            /// permissions/

            PauserRegistry pauserRegistry = eigenLayerPauserReg;
            assertTrue(pauserRegistry.isPauser(pauserMultisig), "pauser multisig should be pauser");
            assertTrue(pauserRegistry.isPauser(operationsMultisig), "ops multisig should be pauser");
            assertTrue(pauserRegistry.isPauser(executorMultisig), "executor multisig should be pauser");
            assertTrue(pauserRegistry.unpauser() == executorMultisig, "executor multisig should be unpauser");

            /// PermissionController has no initial storage
        }

        {
            /// core/

            assertTrue(allocationManager.delegation() == delegationManager, "alm.dm invalid");
            assertTrue(allocationManager.pauserRegistry() == eigenLayerPauserReg, "alm.pR invalid");
            assertTrue(allocationManager.permissionController() == permissionController, "alm.pc invalid");
            assertTrue(allocationManager.DEALLOCATION_DELAY() == DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS, "alm.deallocDelay invalid");
            assertTrue(
                allocationManager.ALLOCATION_CONFIGURATION_DELAY() == ALLOCATION_CONFIGURATION_DELAY,
                "alm.configDelay invalid"
            );

            assertTrue(avsDirectory.delegation() == delegationManager, "avsD.dm invalid");
            assertTrue(avsDirectory.pauserRegistry() == eigenLayerPauserReg, "avsD.pR invalid");

            assertTrue(delegationManager.strategyManager() == strategyManager, "dm.sm invalid");
            assertTrue(delegationManager.eigenPodManager() == eigenPodManager, "dm.epm invalid");
            assertTrue(delegationManager.allocationManager() == allocationManager, "dm.alm invalid");
            assertTrue(delegationManager.pauserRegistry() == eigenLayerPauserReg, "dm.pR invalid");
            assertTrue(delegationManager.permissionController() == permissionController, "dm.pc invalid");
            assertTrue(
                delegationManager.minWithdrawalDelayBlocks() == DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS, "dm.withdrawalDelay invalid"
            );

            assertTrue(rewardsCoordinator.delegationManager() == delegationManager, "rc.dm invalid");
            assertTrue(rewardsCoordinator.strategyManager() == strategyManager, "rc.sm invalid");
            assertTrue(rewardsCoordinator.allocationManager() == allocationManager, "rc.alm invalid");
            assertTrue(rewardsCoordinator.pauserRegistry() == eigenLayerPauserReg, "rc.pR invalid");
            assertTrue(rewardsCoordinator.permissionController() == permissionController, "rc.pc invalid");
            assertTrue(
                rewardsCoordinator.CALCULATION_INTERVAL_SECONDS() == REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS, "rc.calcInterval invalid"
            );
            assertTrue(rewardsCoordinator.MAX_REWARDS_DURATION() == REWARDS_COORDINATOR_MAX_REWARDS_DURATION, "rc.rewardsDuration invalid");
            assertTrue(rewardsCoordinator.MAX_RETROACTIVE_LENGTH() == REWARDS_COORDINATOR_MAX_RETROACTIVE_LENGTH, "rc.retroLength invalid");
            assertTrue(rewardsCoordinator.MAX_FUTURE_LENGTH() == REWARDS_COORDINATOR_MAX_FUTURE_LENGTH, "rc.futureLength invalid");
            assertTrue(rewardsCoordinator.GENESIS_REWARDS_TIMESTAMP() == REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP, "rc.genesis invalid");

            assertTrue(strategyManager.delegation() == delegationManager, "sm.dm invalid");
            assertTrue(strategyManager.pauserRegistry() == eigenLayerPauserReg, "sm.pR invalid");
        }

        {
            /// pods/
            assertTrue(address(eigenPodImplementation.ethPOS()) == ETHPOSDepositAddress, "ep.ethPOS invalid");
            assertTrue(eigenPodImplementation.eigenPodManager() == eigenPodManager, "ep.epm invalid");
            assertTrue(eigenPodImplementation.GENESIS_TIME() == EIGENPOD_GENESIS_TIME, "ep.genesis invalid");

            assertTrue(address(eigenPodManager.ethPOS()) == ETHPOSDepositAddress, "epm.ethPOS invalid");
            assertTrue(eigenPodManager.eigenPodBeacon() == eigenPodBeacon, "epm.epBeacon invalid");
            assertTrue(eigenPodManager.delegationManager() == delegationManager, "epm.dm invalid");
            assertTrue(eigenPodManager.pauserRegistry() == eigenLayerPauserReg, "epm.pR invalid");
        }

        {
            /// strategies/
            assertTrue(eigenStrategy.strategyManager() == strategyManager, "eigStrat.sm invalid");
            assertTrue(eigenStrategy.pauserRegistry() == eigenLayerPauserReg, "eigStrat.pR invalid");
            assertTrue(strategyFactoryBeaconImplementation.strategyManager() == strategyManager, "stratBase.sm invalid");
            assertTrue(strategyFactoryBeaconImplementation.pauserRegistry() == eigenLayerPauserReg, "stratBase.pR invalid");

            assertTrue(
                baseStrategyImplementation.strategyManager() == strategyManager, "stratBaseTVL.sm invalid"
            );
            assertTrue(baseStrategyImplementation.pauserRegistry() == eigenLayerPauserReg, "stratBaseTVL.pR invalid");

            assertTrue(strategyFactory.strategyManager() == strategyManager, "sFact.sm invalid");
            assertTrue(strategyFactory.pauserRegistry() == eigenLayerPauserReg, "sFact.pR invalid");
        }
    }

    function _validateVersion() internal view {
        // permissions/
        assertEq(permissionController.version(), EXPECTED_VERSION, "permissionController version mismatch");

        {
            /// core/
            assertEq(allocationManager.version(), EXPECTED_VERSION, "allocationManager version mismatch");
            assertEq(avsDirectory.version(), EXPECTED_VERSION, "avsDirectory version mismatch");
            assertEq(delegationManager.version(), EXPECTED_VERSION, "delegationManager version mismatch");
            assertEq(rewardsCoordinator.version(), EXPECTED_VERSION, "rewardsCoordinator version mismatch");
            assertEq(strategyManager.version(), EXPECTED_VERSION, "strategyManager version mismatch");
        }

        {
            /// pods/
            assertEq(eigenPodImplementation.version(), EXPECTED_VERSION, "eigenPod version mismatch");
            assertEq(eigenPodManager.version(), EXPECTED_VERSION, "eigenPodManager version mismatch");
        }

        {
            /// strategies/
            assertEq(eigenStrategy.version(), EXPECTED_VERSION, "eigenStrategy version mismatch");
            assertEq(strategyFactoryBeaconImplementation.version(), EXPECTED_VERSION, "strategyBase version mismatch");
            assertEq(baseStrategyImplementation.version(), EXPECTED_VERSION, "strategyBaseTVLLimits version mismatch");
            assertEq(strategyFactory.version(), EXPECTED_VERSION, "strategyFactory version mismatch");
        }
    }
}
