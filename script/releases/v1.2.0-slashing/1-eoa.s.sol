// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "src/contracts/libraries/BeaconChainProofs.sol";

contract Deploy is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        // Deploy EigenPodBeacon Impl
        deployImpl({
            name: type(EigenPod).name,
            deployedTo: address(new EigenPod({
                _ethPOS: Env.ethPOS(),
                _eigenPodManager: Env.proxy.eigenPodManager(),
                _GENESIS_TIME: Env.EIGENPOD_GENESIS_TIME()
            }))
        });

        vm.stopBroadcast();
    }

    function testDeploy() public virtual {
        _runAsEOA();
        _validateNewImplAddresses(false);
        _validateImplConstructors();
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
            Env.beacon.eigenPod().implementation(),
            address(Env.impl.eigenPod()),
            "eigenPod impl failed"
        );
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        /// pods/
        EigenPod eigenPod = Env.impl.eigenPod();
        assertTrue(eigenPod.ethPOS() == Env.ethPOS(), "ep.ethPOS invalid");
        assertTrue(eigenPod.eigenPodManager() == Env.proxy.eigenPodManager(), "ep.epm invalid");
        assertTrue(eigenPod.GENESIS_TIME() == Env.EIGENPOD_GENESIS_TIME(), "ep.genesis invalid");
        if (block.chainid == 5) {
            assertEq(BeaconChainProofs.PECTRA_FORK_TIMESTAMP, 1_739_352_768);
        } 
        
        // TODO: set this to 0 for mainnet, till the fork timestamp is provided
        if (block.chainid == 1) {
            assertEq(BeaconChainProofs.PECTRA_FORK_TIMESTAMP, 0);
        }
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
