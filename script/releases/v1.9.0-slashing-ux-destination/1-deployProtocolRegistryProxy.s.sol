// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "../Env.sol";
import {CrosschainDeployLib} from "script/releases/CrosschainDeployLib.sol";

/// Purpose: Deploy the Protocol Registry contract
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
        if (Env.isCoreProtocolDeployed()) {
            return;
        }

        // Destination chains may already have the deterministic CREATE2 proxy deployed
        // (e.g. if this step was run previously). In that case, attempting to deploy
        // again will revert with a CreateX collision. Instead, attach the expected
        // address into the Zeus env and validate invariants.
        if (_areProxiesDeployed()) {
            _addContractsToEnv();
        } else {
            execute();
        }

        _validateProxyAdminIsExpected();
        _validateExpectedProxyAddress();
    }

    /// @dev Validate that proxies are owned by the multichain deployer multisig (temporarily)
    function _validateProxyAdminIsExpected() internal view {
        address admin = Env.getProxyAdminBySlot(address(Env.proxy.protocolRegistry()));
        assertTrue(
            admin == Env.proxyAdmin() || admin == Env.multichainDeployerMultisig(),
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
        return expectedProtocolRegistry.code.length > 0;
    }

    /// @dev Add the contracts to the env
    function _addContractsToEnv() internal {
        address expectedProtocolRegistry =
            _computeExpectedProxyAddress(type(ProtocolRegistry).name, address(Env.impl.emptyContract()));
        _unsafeAddProxyContract(type(ProtocolRegistry).name, expectedProtocolRegistry);
    }

    /// @dev Compute the expected proxy address for a given name and empty contract
    function _computeExpectedProxyAddress(
        string memory name,
        address emptyContract
    ) internal view returns (address) {
        return CrosschainDeployLib.computeCrosschainUpgradeableProxyAddress({
            adminAndDeployer: Env.multichainDeployerMultisig(),
            implementation: emptyContract,
            name: name
        });
    }
}
