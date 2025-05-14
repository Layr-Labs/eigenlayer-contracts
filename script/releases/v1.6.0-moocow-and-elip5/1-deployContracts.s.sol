// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * Purpose: use an EOA to deploy all of the new contracts for this upgrade.
 */
contract Deploy is EOADeployer {
    using Env for *;

    /// forgefmt: disable-next-item
    function _runAsEOA() internal override {
        vm.startBroadcast();


        // We are upgrading 2 contracts: EigenPod and the Eigen token
        deployImpl({
            name: type(EigenPod).name,
            deployedTo: address(
                new EigenPod({
                    _ethPOS: Env.ethPOS(),
                    _eigenPodManager: Env.proxy.eigenPodManager(),
                    _version: Env.deployVersion()
                })
            )
        });

        deployImpl({
            name: type(Eigen).name,
            deployedTo: address(
                new Eigen({
                   _bEIGEN: Env.proxy.beigen(),
                   _version: Env.deployVersion()
                })
            )
        });

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        _runAsEOA();

        _validateNewImplAddresses({areMatching: false});
        _validateProxyAdmins();
        _validateImplConstructors();
        _validateImplsInitialized();
        _validateVersion();
    }

    /// @dev Validate that the `Env.impl` addresses are updated to be distinct from what the proxy
    /// admin reports as the current implementation address.
    ///
    /// Note: The upgrade script can call this with `areMatching == true` to check that these impl
    /// addresses _are_ matches.
    function _validateNewImplAddresses(
        bool areMatching
    ) internal view {
        function (bool, string memory) internal pure assertion = areMatching ? _assertTrue : _assertFalse;

        assertion(Env.beacon.eigenPod().implementation() == address(Env.impl.eigenPod()), "eigenPod impl failed");

        assertion(_getProxyImpl(address(Env.proxy.eigen())) == address(Env.impl.eigen()), "eigen impl failed");
    }

    /// @dev Ensure each deployed TUP/beacon is owned by the proxyAdmin/executorMultisig
    function _validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();

        assertTrue(Env.beacon.eigenPod().owner() == Env.executorMultisig(), "eigenPod beacon owner incorrect");

        assertTrue(_getProxyAdmin(address(Env.proxy.eigen())) == pa, "eigen proxyAdmin incorrect");
    }

    /// @dev Validate the immutables set in the new implementation constructors
    function _validateImplConstructors() internal view {
        {
            /// EigenPod
            EigenPod eigenPod = Env.impl.eigenPod();
            assertTrue(eigenPod.ethPOS() == Env.ethPOS(), "ep.ethPOS invalid");
            assertTrue(eigenPod.eigenPodManager() == Env.proxy.eigenPodManager(), "ep.epm invalid");
            assertTrue(_strEq(eigenPod.version(), Env.deployVersion()), "ep.version failed");
        }

        {
            /// Eigen
            Eigen eigen = Eigen(address(Env.impl.eigen()));
            assertTrue(address(eigen.bEIGEN()) == address(Env.proxy.beigen()), "eigen.beigen invalid");
        }
    }

    /// @dev Call initialize on all deployed implementations to ensure initializers are disabled
    function _validateImplsInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// EigenPod
        EigenPod eigenPod = Env.impl.eigenPod();
        vm.expectRevert(errInit);
        eigenPod.initialize(address(0));

        /// Eigen
        Eigen eigen = Eigen(address(Env.impl.eigen()));
        vm.expectRevert(errInit);
        eigen.initialize(address(0), new address[](0), new uint256[](0), new uint256[](0));
    }

    function _validateVersion() internal view {
        // On future upgrades, just tick the major/minor/patch to validate
        string memory expected = Env.deployVersion();

        assertEq(Env.impl.eigenPod().version(), expected, "eigenPod version mismatch");
        assertEq(Env.impl.eigen().version(), expected, "eigen version mismatch");
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

    function _assertTrue(bool b, string memory err) private pure {
        assertTrue(b, err);
    }

    function _assertFalse(bool b, string memory err) private pure {
        assertFalse(b, err);
    }

    function _strEq(string memory a, string memory b) private pure returns (bool) {
        return keccak256(bytes(a)) == keccak256(bytes(b));
    }
}
