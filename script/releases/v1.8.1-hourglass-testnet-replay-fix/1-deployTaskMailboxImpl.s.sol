// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

/**
 * @title DeployTaskMailboxImpl
 * @notice Deploy new TaskMailbox implementation with task replay fix for destination chains.
 *         This fixes a vulnerability where certificates could be replayed across different tasks
 *         directed at the same operator set by including the taskHash in the messageHash.
 */
contract DeployTaskMailboxImpl is EOADeployer {
    using Env for *;

    /// forgefmt: disable-next-item
    function _runAsEOA() internal override {
        // If we're not on a destination chain and not on version 1.8.0, we don't need to deploy any contracts
        if (!(Env.isDestinationChain() && Env._strEq(Env.envVersion(), "1.8.0"))) {
            return;
        }

        vm.startBroadcast();

        // Deploy TaskMailbox implementation with the replay fix
        deployImpl({
            name: type(TaskMailbox).name,
            deployedTo: address(
                new TaskMailbox({
                    _bn254CertificateVerifier: address(Env.proxy.bn254CertificateVerifier()),
                    _ecdsaCertificateVerifier: address(Env.proxy.ecdsaCertificateVerifier()),
                    _maxTaskSLA: Env.MAX_TASK_SLA(),
                    _version: Env.deployVersion()
                })
            )
        });

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        if (!(Env.isDestinationChain() && Env._strEq(Env.envVersion(), "1.8.0"))) {
            return;
        }

        // Deploy the new TaskMailbox implementation
        runAsEOA();

        _validateNewImplAddress();
        _validateProxyAdmin();
        _validateImplConstructor();
        _validateImplInitialized();
        _validateVersion();
    }

    /// @dev Validate that the new TaskMailbox impl address is distinct from the current one
    function _validateNewImplAddress() internal view {
        address currentImpl = Env._getProxyImpl(address(Env.proxy.taskMailbox()));
        address newImpl = address(Env.impl.taskMailbox());

        assertFalse(currentImpl == newImpl, "TaskMailbox impl should be different from current implementation");
    }

    /// @dev Validate that the TaskMailbox proxy is still owned by the correct ProxyAdmin
    function _validateProxyAdmin() internal view {
        address pa = Env.proxyAdmin();

        assertTrue(Env._getProxyAdmin(address(Env.proxy.taskMailbox())) == pa, "TaskMailbox proxyAdmin incorrect");
    }

    /// @dev Validate the immutables set in the new TaskMailbox implementation constructor
    function _validateImplConstructor() internal view {
        TaskMailbox taskMailboxImpl = Env.impl.taskMailbox();

        // Validate version
        assertEq(
            keccak256(bytes(taskMailboxImpl.version())),
            keccak256(bytes(Env.deployVersion())),
            "TaskMailbox impl version mismatch"
        );

        // Validate certificate verifiers
        assertTrue(
            taskMailboxImpl.BN254_CERTIFICATE_VERIFIER() == address(Env.proxy.bn254CertificateVerifier()),
            "TaskMailbox BN254_CERTIFICATE_VERIFIER mismatch"
        );
        assertTrue(
            taskMailboxImpl.ECDSA_CERTIFICATE_VERIFIER() == address(Env.proxy.ecdsaCertificateVerifier()),
            "TaskMailbox ECDSA_CERTIFICATE_VERIFIER mismatch"
        );

        // Validate MAX_TASK_SLA
        assertEq(taskMailboxImpl.MAX_TASK_SLA(), Env.MAX_TASK_SLA(), "TaskMailbox MAX_TASK_SLA mismatch");
    }

    /// @dev Validate that the new implementation cannot be initialized (should revert)
    function _validateImplInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        TaskMailbox taskMailboxImpl = Env.impl.taskMailbox();

        vm.expectRevert(errInit);
        taskMailboxImpl.initialize(
            address(0), // owner
            0, // feeSplit
            address(0) // feeSplitCollector
        );
    }

    /// @dev Validate the version is correctly set
    function _validateVersion() internal view {
        assertEq(
            keccak256(bytes(Env.impl.taskMailbox().version())),
            keccak256(bytes(Env.deployVersion())),
            "TaskMailbox version should match deploy version"
        );
    }
}
