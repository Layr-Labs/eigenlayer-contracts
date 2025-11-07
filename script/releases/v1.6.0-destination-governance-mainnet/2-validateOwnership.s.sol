// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./1-deployGovernance.s.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "../../releases/Env.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/// @notice Validates ownership of the proxyAdmin for mainnet
contract ValidateOwnership is MultisigBuilder, DeployGovernance {
    using Env for *;

    /// @notice This function doesn't actually transfer ownership, it's for manual processes
    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        if (!Env._strEq(Env.env(), "base") || !Env._strEq(Env.envVersion(), "0.0.0")) {
            return;
        }

        ProxyAdmin proxyAdmin = ProxyAdmin(address(Env.proxyAdmin()));
        proxyAdmin.transferOwnership(Env.executorMultisig());
    }

    function testScript() public virtual {
        if (!Env._strEq(Env.env(), "base") || !Env._strEq(Env.envVersion(), "0.0.0")) {
            return;
        }

        runAsEOA();

        // Complete upgrade
        execute();

        // Validate that ownership has been transferred to executorMultisig
        assertEq(
            ProxyAdmin(address(Env.proxyAdmin())).owner(),
            Env.executorMultisig(),
            "proxyAdmin.owner() != executorMultisig - ownership transfer not completed"
        );
    }
}
