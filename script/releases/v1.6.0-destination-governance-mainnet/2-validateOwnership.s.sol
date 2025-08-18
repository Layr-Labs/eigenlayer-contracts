// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./1-deployGovernance.s.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "../../releases/Env.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/// @notice Validates ownership of the proxyAdmin for mainnet (test script only)
/// @dev This script only validates the ownership transfer was completed correctly
/// @dev For mainnet, the ownership transfer must be done manually via the Safe UI
contract ValidateOwnership is MultisigBuilder, DeployGovernance {
    using Env for *;

    /// @notice This function doesn't actually transfer ownership, it's for manual processes
    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        if (!Env._strEq(Env.env(), "base")) {
            return;
        }

        // For mainnet, ownership transfer must be done manually via Safe UI
        // This function serves as a placeholder for the manual process
        // The actual transfer is validated in testScript()
    }

    function testScript() public virtual {
        if (!Env._strEq(Env.env(), "base")) {
            return;
        }

        // Complete the first step of the upgrade
        runAsEOA();

        // Validate that ownership has been transferred to executorMultisig
        assertEq(
            ProxyAdmin(address(Env.proxyAdmin())).owner(),
            Env.executorMultisig(),
            "proxyAdmin.owner() != executorMultisig - ownership transfer must be completed manually via Safe UI"
        );
    }
}
