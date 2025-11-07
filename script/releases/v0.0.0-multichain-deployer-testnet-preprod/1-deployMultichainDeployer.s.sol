// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {MultisigDeployLib, IMultisig} from "script/releases/MultisigDeployLib.sol";
import "forge-std/console.sol";
import "../../releases/Env.sol";

/// @notice Deploy the multichain deployer multisig
/// @dev This script is used to deploy the multichain deployer multisig on the destination chain
/// @dev This script should ONLY be used for testnet environments
/// @dev The SAFE version is 1.4.1
contract DeployMultichainDeployer is EOADeployer {
    using Env for *;

    /// @notice The type of environment we are deploying to
    enum EnvType {
        NULL,
        PREPROD,
        TESTNET
    }

    /// @dev The environment type
    EnvType public envType = EnvType.NULL;

    /// @dev Setup the environment type
    modifier setupEnv() {
        if (Env._strEq(Env.env(), "preprod-hoodi")) {
            envType = EnvType.PREPROD;
        } else if (
            Env._strEq(Env.env(), "testnet-sepolia") || Env._strEq(Env.env(), "testnet-base-sepolia")
                || Env._strEq(Env.env(), "testnet-hoodi")
        ) {
            envType = EnvType.TESTNET;
        }
        _;
    }

    /// @dev The expected address of the multichain deployer on the destination chain
    address public constant MULTICHAIN_DEPLOYER_EXPECTED_ADDRESS_TESTNET = 0xA591635DE4C254BD3fa9C9Db9000eA6488344C28;
    address public constant MULTICHAIN_DEPLOYER_EXPECTED_ADDRESS_PREPROD = 0x011aF0ABDBcffBF879de4E1F1F8e346bDCAc0653;

    /// @dev Salt for deploying the multichain deployer multisig
    uint256 public constant SALT_TESTNET = 5;
    uint256 public constant SALT_PREPROD = 6;

    /// @dev Initial threshold for the multichain deployer multisig - same for testnet and preprod
    uint256 public constant INITIAL_THRESHOLD = 1;

    /// @dev Initial owner of the multichain deployer multisig - same for testnet and preprod
    address public constant INITIAL_OWNER = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;

    /// @dev Payment receiver for the multichain deployer multisig - same for testnet and preprod
    /// @dev These addresses were originally deployed via the UI, which included a payment receiver
    address public constant PAYMENT_RECEIVER = 0x5afe7A11E7000000000000000000000000000000;

    function _runAsEOA() internal override setupEnv {
        // Dont' run if we're on base or mainnet
        if (Env._strEq(Env.env(), "base") || Env._strEq(Env.env(), "mainnet")) {
            return;
        }

        vm.startBroadcast();

        address[] memory initialOwners = new address[](1);
        initialOwners[0] = INITIAL_OWNER;

        address multichainDeployerMultisig = MultisigDeployLib.deployMultisigWithPaymentReceiver({
            initialOwners: initialOwners,
            initialThreshold: INITIAL_THRESHOLD,
            salt: _getSalt(),
            paymentReceiver: PAYMENT_RECEIVER
        });

        vm.stopBroadcast();

        // Update config
        zUpdate("multichainDeployerMultisig", multichainDeployerMultisig);
    }

    function testScript() public virtual setupEnv {
        // Dont' run if we're on base or mainnet
        if (Env._strEq(Env.env(), "base") || Env._strEq(Env.env(), "mainnet")) {
            return;
        }

        // If the multichain deployer multisig is already deployed, we need to add the contracts to the env
        if (_isDeployerMultisigDeployed()) {
            _addContractsToEnv();
        } else {
            // Otherwise, we need to deploy the multichain deployer multisig
            super.runAsEOA();
        }

        // Check that the multichain deployer multisig is deployed at the expected address
        assertEq(Env.multichainDeployerMultisig(), _getExpectedDeployerMultisig());

        // Check the threshold is 1
        assertEq(IMultisig(address(Env.multichainDeployerMultisig())).getThreshold(), 1);

        // Check owners
        address[] memory owners = IMultisig(address(Env.multichainDeployerMultisig())).getOwners();
        assertEq(owners.length, 1, "Expected 1 owner");
        assertEq(owners[0], INITIAL_OWNER, "Expected initial owner");

        // Check the version is 1.4.1
        assertEq(IMultisig(address(Env.multichainDeployerMultisig())).VERSION(), "1.4.1");
    }

    /// @dev Check if the multichain deployer multisig is deployed for a given environment type
    function _isDeployerMultisigDeployed() internal view returns (bool) {
        address deployer = _getExpectedDeployerMultisig();
        return deployer.code.length > 0;
    }

    /// @dev Add contracts to the env
    /// @dev This function should only be called if the multichain deployer multisig is already deployed
    function _addContractsToEnv() internal {
        address deployer = _getExpectedDeployerMultisig();
        require(_isDeployerMultisigDeployed(), "Multichain deployer multisig not deployed");
        zUpdate("multichainDeployerMultisig", deployer);
    }

    /// @dev Get the salt for a given environment type
    function _getSalt() internal view returns (uint256) {
        if (envType == EnvType.PREPROD) {
            return SALT_PREPROD;
        } else if (envType == EnvType.TESTNET) {
            return SALT_TESTNET;
        }
        revert("Invalid environment type");
    }

    function _getExpectedDeployerMultisig() internal view returns (address) {
        if (envType == EnvType.PREPROD) {
            return MULTICHAIN_DEPLOYER_EXPECTED_ADDRESS_PREPROD;
        } else if (envType == EnvType.TESTNET) {
            return MULTICHAIN_DEPLOYER_EXPECTED_ADDRESS_TESTNET;
        }
        revert("Invalid environment type");
    }
}
