// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./1-deployGovernance.s.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "../../releases/Env.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/// @notice Transfers ownership of the proxyAdmin to the executorMultisig
/// @dev Assumes that the proxyAdmin has been deployed and that the opsMultisig has been configured
/// @dev For base, the ops multisig is 3/n, so we have to manually propose the transfer of ownership
///      In this case, the ownership will have to be updated via the safe UI and script validated here
///      to complete the upgrade
contract TransferOwnership is MultisigBuilder, DeployGovernance {
    using Env for *;

    /// forgefmt: disable-next-item
    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        if (!Env.isDestinationChain()) { 
            return;
        }

        // For base, the ops multisig is 3/n, so we have to manually propose the transfer of ownership, due to 
        // being unable to do a delegateCall from the SAFE UI
        if (Env._strEq(Env.env(), "base")) { 
            return;
        }
        // Transfer ownership of the proxyAdmin to the executorMultisig
        ProxyAdmin(address(Env.proxyAdmin())).transferOwnership(Env.executorMultisig());
    }

    function testScript() public virtual {
        if (!Env.isDestinationChain()) {
            return;
        }

        if (
            Env._strEq(Env.env(), "preprod") || Env._strEq(Env.env(), "testnet-sepolia")
                || Env._strEq(Env.env(), "mainnet") || Env._strEq(Env.env(), "testnet-holesky")
                || Env._strEq(Env.env(), "testnet-hoodi")
        ) {
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
