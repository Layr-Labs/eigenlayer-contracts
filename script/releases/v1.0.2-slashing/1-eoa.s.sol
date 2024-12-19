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
        deployImpl({
            name: type(StrategyManager).name,
            deployedTo: address(new StrategyManager({
                _delegation: Env.proxy.delegationManager(),
                _pauserRegistry: Env.impl.pauserRegistry()
            }))
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
        function (address, address, string memory) internal pure assertion =
            areMatching ? _assertMatch : _assertNotMatch;


        assertion(
            _getProxyImpl(address(Env.proxy.strategyManager())),
            address(Env.impl.strategyManager()),
            "strategyManager impl failed"
        );
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        StrategyManager strategyManager = Env.impl.strategyManager();
        assertTrue(strategyManager.delegation() == Env.proxy.delegationManager(), "sm.dm invalid");
        assertTrue(strategyManager.pauserRegistry() == Env.impl.pauserRegistry(), "sm.pR invalid");
    }

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        StrategyManager strategyManager = Env.impl.strategyManager();
        vm.expectRevert(errInit);
        strategyManager.initialize(address(0), address(0), 0);
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