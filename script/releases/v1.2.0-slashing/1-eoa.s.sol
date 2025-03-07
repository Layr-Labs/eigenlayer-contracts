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

        // Deploy EigenPodManager Impl
        deployImpl({
            name: type(EigenPodManager).name,
            deployedTo: address(
                new EigenPodManager({
                    _ethPOS: Env.ethPOS(),
                    _eigenPodBeacon: Env.beacon.eigenPod(),
                    _delegationManager: Env.proxy.delegationManager(),
                    _pauserRegistry: Env.impl.pauserRegistry(),
                    _version: "v1.2.0"
                })
            )
        });

        // Deploy EigenPodBeacon Impl
        deployImpl({
            name: type(EigenPod).name,
            deployedTo: address(
                new EigenPod({
                    _ethPOS: Env.ethPOS(),
                    _eigenPodManager: Env.proxy.eigenPodManager(),
                    _GENESIS_TIME: Env.EIGENPOD_GENESIS_TIME(),
                    _version: "v1.2.0"
                })
            )
        });

        vm.stopBroadcast();
    }

    function testDeploy() public virtual {
        _runAsEOA();
        _validateNewImplAddresses(false);
        _validateProxyAdmins();
        _validateImplConstructors();
        _validateImplsInitialized();
    }

    /// @dev Validate that the `Env.impl` addresses are updated to be distinct from what the proxy
    /// admin reports as the current implementation address.
    ///
    /// Note: The upgrade script can call this with `areMatching == true` to check that these impl
    /// addresses _are_ matches.
    function _validateNewImplAddresses(
        bool areMatching
    ) internal view {
        function(address, address, string memory)
            internal
            pure assertion = areMatching ? _assertMatch : _assertNotMatch;

        assertion(Env.beacon.eigenPod().implementation(), address(Env.impl.eigenPod()), "eigenPod impl failed");

        assertion(
            _getProxyImpl(address(Env.proxy.eigenPodManager())),
            address(Env.impl.eigenPodManager()),
            "eigenPodManager impl failed"
        );
    }

    /// @dev Ensure each deployed TUP/beacon is owned by the proxyAdmin/executorMultisig
    function _validateProxyAdmins() internal view {
        assertTrue(Env.beacon.eigenPod().owner() == Env.executorMultisig(), "eigenPod beacon owner incorrect");

        assertTrue(
            _getProxyAdmin(address(Env.proxy.eigenPodManager())) == Env.proxyAdmin(),
            "eigenPodManager proxyAdmin incorrect"
        );
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        /// pods/
        EigenPod eigenPod = Env.impl.eigenPod();
        assertTrue(eigenPod.ethPOS() == Env.ethPOS(), "ep.ethPOS invalid");
        assertTrue(eigenPod.eigenPodManager() == Env.proxy.eigenPodManager(), "ep.epm invalid");
        assertTrue(eigenPod.GENESIS_TIME() == Env.EIGENPOD_GENESIS_TIME(), "ep.genesis invalid");

        /// manager/
        EigenPodManager eigenPodManager = Env.impl.eigenPodManager();
        assertTrue(eigenPodManager.ethPOS() == Env.ethPOS(), "epm.ethPOS invalid");
        assertTrue(eigenPodManager.eigenPodBeacon() == Env.beacon.eigenPod(), "epm.epBeacon invalid");
        assertTrue(eigenPodManager.delegationManager() == Env.proxy.delegationManager(), "epm.dm invalid");
        assertTrue(eigenPodManager.pauserRegistry() == Env.impl.pauserRegistry(), "epm.pR invalid");
    }

    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        EigenPod eigenPod = Env.impl.eigenPod();
        vm.expectRevert(errInit);
        eigenPod.initialize(address(0));

        EigenPodManager eigenPodManager = Env.impl.eigenPodManager();
        vm.expectRevert(errInit);
        eigenPodManager.initialize(address(0), 0);
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

    function _assertMatch(address a, address b, string memory err) private pure {
        assertEq(a, b, err);
    }

    function _assertNotMatch(address a, address b, string memory err) private pure {
        assertNotEq(a, b, err);
    }
}
