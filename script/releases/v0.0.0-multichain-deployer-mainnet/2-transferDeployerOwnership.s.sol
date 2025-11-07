// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./1-deployMultichainDeployer.s.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {IMultisig} from "script/releases/MultisigDeployLib.sol";
import "../../releases/Env.sol";

// For TOML parsing
import {stdToml} from "forge-std/StdToml.sol";

/**
 * Purpose: Deploy proxy contracts for the destination chain using a multisig.
 */
contract TransferDeployerOwnership is MultisigBuilder, DeployMultichainDeployer {
    using stdToml for string;
    using Env for *;

    /// @dev The new threshold for the multichain deployer multisig
    uint256 public constant NEW_THRESHOLD = 3;

    /// forgefmt: disable-next-item
    function _runAsMultisig() internal virtual override prank(Env.multichainDeployerMultisig()) {
        // If we are not on base or mainnet, return
        if (!Env._strEq(Env.env(), "base") && !Env._strEq(Env.env(), "mainnet")) {
            return;
        }

        // Get owners
        address[] memory owners = _getMultisigOwners();

        // Add owners to the multisig
        for (uint256 i = 0; i < owners.length; i++) {
            IMultisig(address(Env.multichainDeployerMultisig())).addOwnerWithThreshold(owners[i], INITIAL_THRESHOLD);
        }

        // Make threshold 3/7
        IMultisig(address(Env.multichainDeployerMultisig())).changeThreshold(NEW_THRESHOLD);
    }

    function testScript() public virtual override {
        // If we are not on base or mainnet, return
        if (!Env._strEq(Env.env(), "base") && !Env._strEq(Env.env(), "mainnet")) {
            return;
        }

        // 1. Deploy destination chain contracts
        DeployMultichainDeployer.testScript();

        // 2. Execute the transfer of ownership & threshold changes
        // If the threshold is 1, we can skip the execution
        if (IMultisig(address(Env.multichainDeployerMultisig())).getThreshold() == 1) {
            execute();
        }

        // 3. Checks

        // Check that the threshold is correct
        assertEq(IMultisig(address(Env.multichainDeployerMultisig())).getThreshold(), NEW_THRESHOLD);

        // Check that the owners are correct
        address[] memory expectedOwners = _getMultisigOwners();
        for (uint256 i = 0; i < expectedOwners.length; i++) {
            assertTrue(
                IMultisig(address(Env.multichainDeployerMultisig())).isOwner(expectedOwners[i]), "Owner not found"
            );
        }
        // Assert that the initial owner is still in the multisig
        assertTrue(
            IMultisig(address(Env.multichainDeployerMultisig())).isOwner(INITIAL_OWNER), "Initial owner not found"
        );

        // Assert that the owners length is correct
        address[] memory actualOwners = IMultisig(address(Env.multichainDeployerMultisig())).getOwners();
        assertEq(actualOwners.length, expectedOwners.length + 1, "Expected owners length mismatch"); // We add one since the initial owner is not removed
    }

    /// @dev Get the owners from the TOML file
    function _getMultisigOwners() internal view returns (address[] memory) {
        string memory path = "script/releases/v0.0.0-multichain-deployer-mainnet/owners.toml";
        // Read the TOML file
        string memory root = vm.projectRoot();
        string memory fullPath = string.concat(root, "/", path);
        string memory toml = vm.readFile(fullPath);

        // Parse the owners array from the TOML
        address[] memory owners = toml.readAddressArray(".owners");

        return owners;
    }
}
