// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {CrosschainDeployLib} from "script/releases/CrosschainDeployLib.sol";
import "../Env.sol";

/**
 * Purpose: Deploy proxy contracts for the TaskMailbox on destination chains using a multisig.
 */
contract DeployDestinationChainProxies is MultisigBuilder {
    using Env for *;

    /// forgefmt: disable-next-item
    function _runAsMultisig() internal virtual override {
        // If we're not on a destination chain or we're on a version that already has these contracts deployed, we don't need to deploy any contracts
        if (!Env.isDestinationChain() || _isAlreadyDeployed()) {
            return;
        }

        // We don't use the prank modifier here, since we have to write to the env
        _startPrank(Env.multichainDeployerMultisig());

        // Deploy the proxy pointing to an empty contract
        ITransparentUpgradeableProxy taskMailboxProxy = CrosschainDeployLib.deployCrosschainProxy({
            implementation: address(Env.impl.emptyContract()),
            adminAndDeployer: Env.multichainDeployerMultisig(),
            name: type(TaskMailbox).name
        });

        // Stop pranking
        _stopPrank();

        // Save all the contracts to the env
        _unsafeAddProxyContract(type(TaskMailbox).name, address(taskMailboxProxy));
    }

    function testScript() public virtual {
        if (!Env.isDestinationChain() || _isAlreadyDeployed()) {
            return;
        }

        // 1. Deploy the destination chain proxies
        execute();

        _validateExpectedProxyAddress();
        _validateProxyAdminIsMultisig();
    }

    /// @dev Validate that the expected proxy address is deployed
    function _validateExpectedProxyAddress() internal view {
        address expectedProxy = _computeExpectedProxyAddress(type(TaskMailbox).name, address(Env.impl.emptyContract()));
        address actualProxy = address(Env.proxy.taskMailbox());
        assertEq(expectedProxy, actualProxy, "taskMailbox proxy address mismatch");
    }

    /// @dev Validate that proxies are owned by the multichain deployer multisig (temporarily)
    function _validateProxyAdminIsMultisig() internal view {
        address multisig = Env.multichainDeployerMultisig();

        assertTrue(
            _getProxyAdminBySlot(address(Env.proxy.taskMailbox())) == multisig,
            "taskMailbox proxyAdmin should be multisig"
        );
    }

    /// @dev We have to use the slot directly since _getProxyAdmin expects the caller to be the actual proxy admin
    function _getProxyAdminBySlot(
        address _proxy
    ) internal view returns (address) {
        bytes32 adminSlot = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
        address admin = address(uint160(uint256(vm.load(address(_proxy), adminSlot))));
        return admin;
    }

    /// @dev Check if the TaskMailbox proxy is deployed by checking if it's code size is greater than 0
    function _areProxiesDeployed() internal view returns (bool) {
        EmptyContract emptyContract = Env.impl.emptyContract();
        address taskMailboxProxy = _computeExpectedProxyAddress(type(TaskMailbox).name, address(emptyContract));

        if (taskMailboxProxy.code.length > 0) {
            return true;
        }
        return false;
    }

    /// @dev Add the contracts to the env
    function _addContractsToEnv() internal {
        EmptyContract emptyContract = Env.impl.emptyContract();
        _unsafeAddProxyContract(
            type(TaskMailbox).name, _computeExpectedProxyAddress(type(TaskMailbox).name, address(emptyContract))
        );
    }

    /// @dev Compute the expected proxy address for a given name and empty contract
    function _computeExpectedProxyAddress(string memory name, address emptyContract) internal view returns (address) {
        return CrosschainDeployLib.computeCrosschainUpgradeableProxyAddress({
            adminAndDeployer: Env.multichainDeployerMultisig(),
            implementation: emptyContract,
            name: name
        });
    }

    /// @dev Check if the version is already deployed
    function _isAlreadyDeployed() internal view returns (bool) {
        if (keccak256(bytes(Env.envVersion())) == keccak256(bytes("1.8.0-rc.0"))) {
            return true;
        }
        return false;
    }
}
