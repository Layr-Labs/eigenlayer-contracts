// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

// Just upgrade StrategyManager
contract Deploy is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        // Deploy DM
        deployImpl({
            name: type(DelegationManager).name,
            deployedTo: address(new DelegationManager({
                _strategyManager: Env.proxy.strategyManager(),
                _eigenPodManager: Env.proxy.eigenPodManager(),
                _allocationManager: Env.proxy.allocationManager(),
                _pauserRegistry: Env.impl.pauserRegistry(),
                _permissionController: Env.proxy.permissionController(),
                _MIN_WITHDRAWAL_DELAY: Env.MIN_WITHDRAWAL_DELAY()
            }))
        });

        // Deploy AVSD
        deployImpl({
            name: type(AVSDirectory).name,
            deployedTo: address(new AVSDirectory({
                _delegation: Env.proxy.delegationManager(),
                _pauserRegistry: Env.impl.pauserRegistry()
            }))
        });

        // Deploy SM
        deployImpl({
            name: type(StrategyManager).name,
            deployedTo: address(new StrategyManager({
                _delegation: Env.proxy.delegationManager(),
                _pauserRegistry: Env.impl.pauserRegistry()
            }))
        });

        // Deploy RC
        deployImpl({
            name: type(RewardsCoordinator).name,
            deployedTo: address(new RewardsCoordinator({
                _delegationManager: Env.proxy.delegationManager(),
                _strategyManager: Env.proxy.strategyManager(),
                _allocationManager: Env.proxy.allocationManager(),
                _pauserRegistry: Env.impl.pauserRegistry(),
                _permissionController: Env.proxy.permissionController(),
                _CALCULATION_INTERVAL_SECONDS: Env.CALCULATION_INTERVAL_SECONDS(),
                _MAX_REWARDS_DURATION: Env.MAX_REWARDS_DURATION(),
                _MAX_RETROACTIVE_LENGTH: Env.MAX_RETROACTIVE_LENGTH(),
                _MAX_FUTURE_LENGTH: Env.MAX_FUTURE_LENGTH(),
                _GENESIS_REWARDS_TIMESTAMP: Env.GENESIS_REWARDS_TIMESTAMP()
            }))
        });

        vm.stopBroadcast();
    }

    function testDeploy() public virtual {
        _runAsEOA();
        _validateDomainSeparatorNonZero();
        _validateNewImplAddresses(false);
        _validateImplConstructors();
        _validateImplsInitialized();
        _validateRCValues();
    }

    function _validateDomainSeparatorNonZero() internal view {
        bytes32 zeroDomainSeparator = bytes32(0);

        assertFalse(Env.impl.avsDirectory().domainSeparator() == zeroDomainSeparator, "avsD.domainSeparator is zero");
        assertFalse(Env.impl.delegationManager().domainSeparator() == zeroDomainSeparator, "dm.domainSeparator is zero");
        assertFalse(Env.impl.strategyManager().domainSeparator() == zeroDomainSeparator, "rc.domainSeparator is zero");
    }


    /// @dev Validate that the `Env.impl` addresses are updated to be distinct from what the proxy
    /// admin reports as the current implementation address.
    ///
    /// Note: The upgrade script can call this with `areMatching == true` to check that these impl
    /// addresses _are_ matches.
    function _validateNewImplAddresses(bool areMatching) internal view {
        function (address, address, string memory) internal pure assertion =
            areMatching ? _assertMatch : _assertNotMatch;


        assertion(
            _getProxyImpl(address(Env.proxy.strategyManager())),
            address(Env.impl.strategyManager()),
            "strategyManager impl failed"
        );

        assertion(
            _getProxyImpl(address(Env.proxy.delegationManager())),
            address(Env.impl.delegationManager()),
            "delegationManager impl failed"
        );

        assertion(
            _getProxyImpl(address(Env.proxy.avsDirectory())),
            address(Env.impl.avsDirectory()),
            "avsdirectory impl failed"
        );

        assertion(
            _getProxyImpl(address(Env.proxy.rewardsCoordinator())),
            address(Env.impl.rewardsCoordinator()),
            "rewardsCoordinator impl failed"
        );
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        AVSDirectory avsDirectory = Env.impl.avsDirectory();
        assertTrue(avsDirectory.delegation() == Env.proxy.delegationManager(), "avsD.dm invalid");
        assertTrue(avsDirectory.pauserRegistry() == Env.impl.pauserRegistry(), "avsD.pR invalid");

        DelegationManager delegation = Env.impl.delegationManager();
        assertTrue(delegation.strategyManager() == Env.proxy.strategyManager(), "dm.sm invalid");
        assertTrue(delegation.eigenPodManager() == Env.proxy.eigenPodManager(), "dm.epm invalid");
        assertTrue(delegation.allocationManager() == Env.proxy.allocationManager(), "dm.alm invalid");
        assertTrue(delegation.pauserRegistry() == Env.impl.pauserRegistry(), "dm.pR invalid");
        assertTrue(delegation.permissionController() == Env.proxy.permissionController(), "dm.pc invalid");
        assertTrue(delegation.minWithdrawalDelayBlocks() == Env.MIN_WITHDRAWAL_DELAY(), "dm.withdrawalDelay invalid");

        RewardsCoordinator rewards = Env.impl.rewardsCoordinator();
        assertTrue(rewards.delegationManager() == Env.proxy.delegationManager(), "rc.dm invalid");
        assertTrue(rewards.strategyManager() == Env.proxy.strategyManager(), "rc.sm invalid");
        assertTrue(rewards.allocationManager() == Env.proxy.allocationManager(), "rc.alm invalid");
        assertTrue(rewards.pauserRegistry() == Env.impl.pauserRegistry(), "rc.pR invalid");
        assertTrue(rewards.permissionController() == Env.proxy.permissionController(), "rc.pc invalid");
        assertTrue(rewards.CALCULATION_INTERVAL_SECONDS() == Env.CALCULATION_INTERVAL_SECONDS(), "rc.calcInterval invalid");
        assertTrue(rewards.MAX_REWARDS_DURATION() == Env.MAX_REWARDS_DURATION(), "rc.rewardsDuration invalid");
        assertTrue(rewards.MAX_RETROACTIVE_LENGTH() == Env.MAX_RETROACTIVE_LENGTH(), "rc.retroLength invalid");
        assertTrue(rewards.MAX_FUTURE_LENGTH() == Env.MAX_FUTURE_LENGTH(), "rc.futureLength invalid");
        assertTrue(rewards.GENESIS_REWARDS_TIMESTAMP() == Env.GENESIS_REWARDS_TIMESTAMP(), "rc.genesis invalid");
        
        StrategyManager strategyManager = Env.impl.strategyManager();
        assertTrue(strategyManager.delegation() == Env.proxy.delegationManager(), "sm.dm invalid");
        assertTrue(strategyManager.pauserRegistry() == Env.impl.pauserRegistry(), "sm.pR invalid");
    }

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        AVSDirectory avsDirectory = Env.impl.avsDirectory();
        vm.expectRevert(errInit);
        avsDirectory.initialize(address(0), 0);

        DelegationManager delegation = Env.impl.delegationManager();
        vm.expectRevert(errInit);
        delegation.initialize(address(0), 0);

        RewardsCoordinator rewards = Env.impl.rewardsCoordinator();
        vm.expectRevert(errInit);
        rewards.initialize(address(0), 0, address(0), 0, 0);

        StrategyManager strategyManager = Env.impl.strategyManager();
        vm.expectRevert(errInit);
        strategyManager.initialize(address(0), address(0), 0);
    }

    function _validateRCValues() internal view {

        RewardsCoordinator rewardsCoordinatorImpl = Env.impl.rewardsCoordinator();
        assertEq(
            rewardsCoordinatorImpl.CALCULATION_INTERVAL_SECONDS(),
            Env.CALCULATION_INTERVAL_SECONDS(),
            "expected REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"
        );
        assertEq(
            rewardsCoordinatorImpl.CALCULATION_INTERVAL_SECONDS(),
            1 days,
            "expected REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"
        );
        assertGt(
            rewardsCoordinatorImpl.CALCULATION_INTERVAL_SECONDS(),
            0,
            "expected non-zero REWARDS_COORDINATOR_CALCULATION_INTERVAL_SECONDS"
        );

        assertEq(rewardsCoordinatorImpl.MAX_REWARDS_DURATION(), Env.MAX_REWARDS_DURATION());
        assertGt(rewardsCoordinatorImpl.MAX_REWARDS_DURATION(), 0);

        assertEq(
            rewardsCoordinatorImpl.MAX_RETROACTIVE_LENGTH(),
            Env.MAX_RETROACTIVE_LENGTH()
        );
        assertGt(rewardsCoordinatorImpl.MAX_RETROACTIVE_LENGTH(), 0);

        assertEq(rewardsCoordinatorImpl.MAX_FUTURE_LENGTH(), Env.MAX_FUTURE_LENGTH());
        assertGt(rewardsCoordinatorImpl.MAX_FUTURE_LENGTH(), 0);

        assertEq(
            rewardsCoordinatorImpl.GENESIS_REWARDS_TIMESTAMP(),
            Env.GENESIS_REWARDS_TIMESTAMP()
        );
        assertGt(rewardsCoordinatorImpl.GENESIS_REWARDS_TIMESTAMP(), 0);
    }   

    /// @dev Query and return `proxyAdmin.getProxyImplementation(proxy)`
    function _getProxyImpl(address proxy) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyImplementation(ITransparentUpgradeableProxy(proxy));
    }

    /// @dev Query and return `proxyAdmin.getProxyAdmin(proxy)`
    function _getProxyAdmin(address proxy) internal view returns (address) {
        return ProxyAdmin(Env.proxyAdmin()).getProxyAdmin(ITransparentUpgradeableProxy(proxy));
    }

    function _assertMatch(address a, address b, string memory err) private pure {
        assertEq(a, b, err);
    }

    function _assertNotMatch(address a, address b, string memory err) private pure {
        assertNotEq(a, b, err);
    }
}