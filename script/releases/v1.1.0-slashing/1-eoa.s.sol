// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

contract Deploy is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        // Deploy RC
        deployImpl({
            name: type(RewardsCoordinator).name,
            deployedTo: address(
                new RewardsCoordinator({
                    _delegationManager: Env.proxy.delegationManager(),
                    _strategyManager: Env.proxy.strategyManager(),
                    _allocationManager: Env.proxy.allocationManager(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _permissionController: Env.proxy.permissionController(),
                    _CALCULATION_INTERVAL_SECONDS: Env
                        .CALCULATION_INTERVAL_SECONDS(),
                    _MAX_REWARDS_DURATION: Env.MAX_REWARDS_DURATION(),
                    _MAX_RETROACTIVE_LENGTH: Env.MAX_RETROACTIVE_LENGTH(),
                    _MAX_FUTURE_LENGTH: Env.MAX_FUTURE_LENGTH(),
                    _GENESIS_REWARDS_TIMESTAMP: Env.GENESIS_REWARDS_TIMESTAMP()
                })
            )
        });

        vm.stopBroadcast();
    }

    function testDeploy() public virtual {
        _runAsEOA();
        _validateNewImplAddresses(false);
        _validateImplConstructors();
        _validateImplsInitialized();
    }

    /// @dev Validate that the `Env.impl` addresses are updated to be distinct from what the proxy
    /// admin reports as the current implementation address.
    ///
    /// Note: The upgrade script can call this with `areMatching == true` to check that these impl
    /// addresses _are_ matches.
    function _validateNewImplAddresses(bool areMatching) internal view {
        function(address, address, string memory)
            internal
            pure assertion = areMatching ? _assertMatch : _assertNotMatch;

        assertion(
            _getProxyImpl(address(Env.proxy.rewardsCoordinator())),
            address(Env.impl.rewardsCoordinator()),
            "rewardsCoordinator impl failed"
        );
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        RewardsCoordinator rewards = Env.impl.rewardsCoordinator();
        assertTrue(
            rewards.delegationManager() == Env.proxy.delegationManager(),
            "rc.dm invalid"
        );
        assertTrue(
            rewards.strategyManager() == Env.proxy.strategyManager(),
            "rc.sm invalid"
        );
        assertTrue(
            rewards.allocationManager() == Env.proxy.allocationManager(),
            "rc.alm invalid"
        );
        assertTrue(
            rewards.pauserRegistry() == Env.impl.pauserRegistry(),
            "rc.pR invalid"
        );
        assertTrue(
            rewards.permissionController() == Env.proxy.permissionController(),
            "rc.pc invalid"
        );
        assertTrue(
            rewards.CALCULATION_INTERVAL_SECONDS() ==
                Env.CALCULATION_INTERVAL_SECONDS(),
            "rc.calcInterval invalid"
        );
        assertTrue(
            rewards.MAX_REWARDS_DURATION() == Env.MAX_REWARDS_DURATION(),
            "rc.rewardsDuration invalid"
        );
        assertTrue(
            rewards.MAX_RETROACTIVE_LENGTH() == Env.MAX_RETROACTIVE_LENGTH(),
            "rc.retroLength invalid"
        );
        assertTrue(
            rewards.MAX_FUTURE_LENGTH() == Env.MAX_FUTURE_LENGTH(),
            "rc.futureLength invalid"
        );
        assertTrue(
            rewards.GENESIS_REWARDS_TIMESTAMP() ==
                Env.GENESIS_REWARDS_TIMESTAMP(),
            "rc.genesis invalid"
        );
    }

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        RewardsCoordinator rewards = Env.impl.rewardsCoordinator();
        vm.expectRevert(errInit);
        rewards.initialize(address(0), 0, address(0), 0, 0);
    }

    /// @dev Query and return `proxyAdmin.getProxyImplementation(proxy)`
    function _getProxyImpl(address proxy) internal view returns (address) {
        return
            ProxyAdmin(Env.proxyAdmin()).getProxyImplementation(
                ITransparentUpgradeableProxy(proxy)
            );
    }

    /// @dev Query and return `proxyAdmin.getProxyAdmin(proxy)`
    function _getProxyAdmin(address proxy) internal view returns (address) {
        return
            ProxyAdmin(Env.proxyAdmin()).getProxyAdmin(
                ITransparentUpgradeableProxy(proxy)
            );
    }

    function _assertMatch(
        address a,
        address b,
        string memory err
    ) private pure {
        assertEq(a, b, err);
    }

    function _assertNotMatch(
        address a,
        address b,
        string memory err
    ) private pure {
        assertNotEq(a, b, err);
    }
}
