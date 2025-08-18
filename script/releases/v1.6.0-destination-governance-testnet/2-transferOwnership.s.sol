// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./1-deployGovernance.s.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "../../releases/Env.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/// @notice Transfers ownership of the proxyAdmin to the executorMultisig for testnet
/// @dev Assumes that the proxyAdmin has been deployed and that the opsMultisig has been configured
contract TransferOwnership is MultisigBuilder, DeployGovernance {
    using Env for *;

    /// forgefmt: disable-next-item
    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        if (!Env.isDestinationChain() || !Env._strEq(Env.env(), "testnet-base-sepolia")) {
            return;
        }

        // Transfer ownership of the proxyAdmin to the executorMultisig
        ProxyAdmin(address(Env.proxyAdmin())).transferOwnership(Env.executorMultisig());
    }

    function testScript() public virtual {
        if (!Env.isDestinationChain() || !Env._strEq(Env.env(), "testnet-base-sepolia")) {
            return;
        }

        // Complete the first step of the upgrade
        runAsEOA();

        execute();

        assertEq(
            ProxyAdmin(address(Env.proxyAdmin())).owner(),
            Env.executorMultisig(),
            "proxyAdmin.owner() != executorMultisig"
        );
    }
}
