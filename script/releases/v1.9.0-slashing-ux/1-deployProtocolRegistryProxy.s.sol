// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "../Env.sol";
import {CrosschainDeployLib} from "script/releases/CrosschainDeployLib.sol";

/**
 * Purpose: Deploy the Protocol Registry contract
 */
contract DeployProtocolRegistryProxy is MultisigBuilder {
    using Env for *;
    using CrosschainDeployLib for *;

    /// forgefmt: disable-next-item
    function _runAsMultisig() internal virtual override {
        // We don't use the prank modifier here, since we have to write to the env
        _startPrank(Env.multichainDeployerMultisig());

        // Deploy Protocol Registry Proxy
        ITransparentUpgradeableProxy protocolRegistryProxy = CrosschainDeployLib.deployCrosschainProxy({
            implementation: address(Env.impl.emptyContract()),
            adminAndDeployer: Env.multichainDeployerMultisig(),
            name: type(ProtocolRegistry).name
        });

        // Stop pranking
        _stopPrank();

        // Save all the contracts to the env
        _unsafeAddProxyContract(type(ProtocolRegistry).name, address(protocolRegistryProxy));
    }

    function testScript() public virtual {
        execute();

        _validateProxyAdminIsMultisig();
        _validateExpectedProxyAddress();
    }

    /// @dev Validate that proxies are owned by the multichain deployer multisig (temporarily)
    function _validateProxyAdminIsMultisig() internal view {
        address multisig = Env.multichainDeployerMultisig();

        assertTrue(
            _getProxyAdminBySlot(address(Env.proxy.protocolRegistry())) == multisig,
            "protocolRegistry proxyAdmin should be multisig"
        );
    }

    /// @dev Validate that the expected proxy address is deployed
    function _validateExpectedProxyAddress() internal view {
        address expectedProxy =
            _computeExpectedProxyAddress(type(ProtocolRegistry).name, address(Env.impl.emptyContract()));
        address actualProxy = address(Env.proxy.protocolRegistry());
        assertTrue(expectedProxy == actualProxy, "protocolRegistry proxy address mismatch");
    }

    /// @dev Check if the proxies are deployed by checking if the empty contract is deployed
    function _areProxiesDeployed() internal view returns (bool) {
        address expectedProtocolRegistry =
            _computeExpectedProxyAddress(type(ProtocolRegistry).name, address(Env.impl.emptyContract()));

        // If the empty contract is deployed, then the proxies are deployed
        if (expectedProtocolRegistry.code.length > 0) {
            return true;
        }
        return false;
    }

    /// @dev Add the contracts to the env
    function _addContractsToEnv() internal {
        address expectedProtocolRegistry =
            _computeExpectedProxyAddress(type(ProtocolRegistry).name, address(Env.impl.emptyContract()));
        _unsafeAddProxyContract(type(ProtocolRegistry).name, expectedProtocolRegistry);
    }

    /// @dev Compute the expected proxy address for a given name and empty contract
    function _computeExpectedProxyAddress(string memory name, address emptyContract) internal view returns (address) {
        return CrosschainDeployLib.computeCrosschainUpgradeableProxyAddress({
            adminAndDeployer: Env.multichainDeployerMultisig(),
            implementation: emptyContract,
            name: name
        });
    }

    /// @dev We have to use the slot directly since _getProxyAdmin expects the caller to be the actual proxy admin
    function _getProxyAdminBySlot(
        address _proxy
    ) internal view returns (address) {
        bytes32 adminSlot = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
        address admin = address(uint160(uint256(vm.load(address(_proxy), adminSlot))));
        return admin;
    }
}
