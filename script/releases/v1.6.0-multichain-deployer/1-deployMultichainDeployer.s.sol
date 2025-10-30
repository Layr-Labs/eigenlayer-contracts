// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {MultisigDeployLib, IMultisig} from "script/releases/MultisigDeployLib.sol";
import "forge-std/console.sol";
import "../../releases/Env.sol";

/// @notice Deploy the multichain deployer multisig
/// @dev This script is used to deploy the multichain deployer multisig on the destination chain
/// @dev This script should ONLY be used for mainnet environments. Testnet environments should follow our notion guide
/// TODO: Add a testnet version of this script
/// @dev The SAFE version is 1.4.1
contract DeployMultichainDeployer is EOADeployer {
    using Env for *;

    /// @dev The expected address of the multichain deployer on the destination chain
    address public constant MULTICHAIN_DEPLOYER_EXPECTED_ADDRESS = 0xa3053EF25F1F7d9D55a7655372B8a31D0f40eCA9;

    /// @dev Salt for deploying the multichain deployer multisig
    uint256 public constant SALT = 0;

    /// @dev Initial threshold for the multichain deployer multisig
    uint256 public constant INITIAL_THRESHOLD = 1;

    /// @dev Initial owner of the multichain deployer multisig
    address public constant INITIAL_OWNER = 0x792e42f05E87Fb9D8b8F9FdFC598B1de20507964;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        address[] memory initialOwners = new address[](1);
        initialOwners[0] = INITIAL_OWNER;

        address multichainDeployerMultisig = MultisigDeployLib.deployMultisig({
            initialOwners: initialOwners,
            initialThreshold: INITIAL_THRESHOLD,
            salt: SALT
        });

        vm.stopBroadcast();

        // Update config
        zUpdate("multichainDeployerMultisig", multichainDeployerMultisig);
    }

    function testScript() public virtual {
        // If the multichain deployer multisig is already deployed, we need to add the contracts to the env
        if (_isDeployerMultisigDeployed()) {
            _addContractsToEnv();
        } else {
            // Otherwise, we need to deploy the multichain deployer multisig
            super.runAsEOA();
        }

        // Check that the multichain deployer multisig is deployed at the expected address
        assertEq(Env.multichainDeployerMultisig(), MULTICHAIN_DEPLOYER_EXPECTED_ADDRESS);

        // Check the threshold is 1
        assertEq(IMultisig(address(Env.multichainDeployerMultisig())).getThreshold(), 1);

        // Check owners
        address[] memory owners = IMultisig(address(Env.multichainDeployerMultisig())).getOwners();
        assertEq(owners.length, 1, "Expected 1 owner");
        assertEq(owners[0], INITIAL_OWNER, "Expected initial owner");

        // Check the version is 1.4.1
        assertEq(IMultisig(address(Env.multichainDeployerMultisig())).VERSION(), "1.4.1");
    }

    function _isDeployerMultisigDeployed() internal view returns (bool) {
        return MULTICHAIN_DEPLOYER_EXPECTED_ADDRESS.code.length > 0;
    }

    /// @dev Add contracts to the env
    /// @dev This function should only be called if the multichain deployer multisig is already deployed
    function _addContractsToEnv() internal {
        require(MULTICHAIN_DEPLOYER_EXPECTED_ADDRESS.code.length > 0, "Multichain deployer multisig not deployed");
        zUpdate("multichainDeployerMultisig", MULTICHAIN_DEPLOYER_EXPECTED_ADDRESS);
    }
}
