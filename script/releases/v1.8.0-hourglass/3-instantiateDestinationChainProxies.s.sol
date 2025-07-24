// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {DeployDestinationChainProxies} from "./1-deployDestinationChainProxies.s.sol";
import {DeployDestinationChainImpls} from "./2-deployDestinationChainImpls.s.sol";
import {CrosschainDeployLib} from "script/releases/CrosschainDeployLib.sol";
import "../Env.sol";

/**
 * Purpose: Upgrade TaskMailbox proxy to point to implementation and transfer control to ProxyAdmin.
 */
contract InstantiateDestinationChainProxies is DeployDestinationChainImpls {
    using Env for *;

    /// forgefmt: disable-next-item
    function _runAsMultisig() internal override prank(Env.multichainDeployerMultisig()) {
        // If we're not on a destination chain, we don't need to do anything
        if (!Env.isDestinationChain()) {
            return;
        }

        // TaskMailbox - we also initialize this contract
        ITransparentUpgradeableProxy taskMailboxProxy =
            ITransparentUpgradeableProxy(payable(address(Env.proxy.taskMailbox())));
        
        // Get initialization parameters
        TaskMailboxInitParams memory initParams = _getTaskMailboxInitParams();
        
        taskMailboxProxy.upgradeToAndCall(
            address(Env.impl.taskMailbox()),
            abi.encodeCall(
                TaskMailbox.initialize,
                (
                    initParams.owner,
                    initParams.feeSplit,
                    initParams.feeSplitCollector
                )
            )
        );

        // Transfer proxy admin ownership
        taskMailboxProxy.changeAdmin(Env.proxyAdmin());
    }

    function testScript() public virtual override {
        if (!Env.isDestinationChain()) {
            return;
        }

        // 1. Deploy the destination chain contracts
        // If proxies are not deployed, deploy them
        if (!_areProxiesDeployed()) {
            DeployDestinationChainProxies._runAsMultisig();
            _unsafeResetHasPranked(); // reset hasPranked so we can use it in the execute()
        } else {
            // Since the proxies are already deployed, we need to update the env with the proper addresses
            _addContractsToEnv();
        }

        // 2. Deploy the destination chain impls
        _mode = OperationalMode.EOA; // Set to EOA mode so we can deploy the impls in the EOA script
        DeployDestinationChainImpls._runAsEOA();

        // 3. Instantiate the destination chain proxies
        execute();

        // 4. Validate the destination chain
        _validateStorage();
        _validateProxyAdmins();
        _validateProxyConstructors();
        _validateProxiesInitialized();
    }

    /// @dev Validate that storage variables are set correctly
    function _validateStorage() internal view {
        // Validate TaskMailbox
        TaskMailbox taskMailbox = Env.proxy.taskMailbox();
        TaskMailboxInitParams memory initParams = _getTaskMailboxInitParams();

        assertTrue(taskMailbox.owner() == initParams.owner, "taskMailbox.owner invalid");
        assertEq(taskMailbox.feeSplit(), initParams.feeSplit, "taskMailbox.feeSplit invalid");
        assertTrue(
            taskMailbox.feeSplitCollector() == initParams.feeSplitCollector, "taskMailbox.feeSplitCollector invalid"
        );
    }

    /// @dev Ensure each deployed TUP/beacon is owned by the proxyAdmin/executorMultisig
    function _validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();

        assertTrue(Env._getProxyAdmin(address(Env.proxy.taskMailbox())) == pa, "taskMailbox proxyAdmin incorrect");
    }

    function _validateProxyConstructors() internal view {
        TaskMailbox taskMailbox = Env.proxy.taskMailbox();
        assertEq(taskMailbox.version(), Env.deployVersion(), "taskMailbox version mismatch");
        assertTrue(
            taskMailbox.BN254_CERTIFICATE_VERIFIER() == address(Env.proxy.bn254CertificateVerifier()),
            "taskMailbox.BN254_CERTIFICATE_VERIFIER mismatch"
        );
        assertTrue(
            taskMailbox.ECDSA_CERTIFICATE_VERIFIER() == address(Env.proxy.ecdsaCertificateVerifier()),
            "taskMailbox.ECDSA_CERTIFICATE_VERIFIER mismatch"
        );
    }

    function _validateProxiesInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// TaskMailbox
        TaskMailbox taskMailbox = Env.proxy.taskMailbox();

        vm.expectRevert(errInit);
        taskMailbox.initialize(
            address(0), // owner
            0, // feeSplit
            address(0) // feeSplitCollector
        );
    }

    function _assertTrue(bool b, string memory err) private pure {
        assertTrue(b, err);
    }

    function _assertFalse(bool b, string memory err) private pure {
        assertFalse(b, err);
    }

    // Define TaskMailbox initialization parameters
    function _getTaskMailboxInitParams() internal view returns (TaskMailboxInitParams memory) {
        TaskMailboxInitParams memory initParams;

        initParams.owner = Env.opsMultisig();
        initParams.feeSplit = 0; // 0% fee split initially
        initParams.feeSplitCollector = Env.opsMultisig(); // Initially set to opsMultisig

        return initParams;
    }

    struct TaskMailboxInitParams {
        address owner;
        uint16 feeSplit;
        address feeSplitCollector;
    }
}
