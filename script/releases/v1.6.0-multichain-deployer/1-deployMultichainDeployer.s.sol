// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "forge-std/console.sol";
import "../../releases/Env.sol";

/// @notice Deploy the multichain deployer multisig
/// @dev This script is used to deploy the multichain deployer multisig on the destination chain
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

    /// @dev Safe proxy factory, this should be the same for all chains
    /// @dev See https://github.com/safe-global/safe-smart-account/blob/main/CHANGELOG.md#version-141 for more details
    /// @dev DOUBLE CHECK THESE ADDRESSES ARE CORRECT ON EACH CHAIN PRIOR TO DEPLOYMENT
    address public constant SAFE_PROXY_FACTORY = 0x4e1DCf7AD4e460CfD30791CCC4F9c8a4f820ec67;
    address public constant SAFE_SINGLETON = 0x41675C099F32341bf84BFc5382aF534df5C7461a;
    address public constant SAFE_FALLBACK_HANDLER = 0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99;

    /// @dev L2 Singletons still need to be passed into the L1 deployment
    /// @dev `SAFE_TO_L2_SETUP` does a no-op if the chain is mainnet
    /// @dev See: https://github.com/safe-global/safe-smart-account/blob/0095f1aa113255d97b476e625760514cc7d10982/contracts/libraries/SafeToL2Setup.sol#L59-L69
    address public constant SAFE_L2_SINGLETON = 0x29fcB43b46531BcA003ddC8FCB67FFE91900C762;
    address public constant SAFE_TO_L2_SETUP = 0xBD89A1CE4DDe368FFAB0eC35506eEcE0b1fFdc54;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        address multichainDeployerMultisig = ISafeProxyFactory(SAFE_PROXY_FACTORY).createProxyWithNonce({
            _singleton: SAFE_SINGLETON,
            initializer: getInitializationData(),
            saltNonce: SALT
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

    function getInitializationData() internal pure returns (bytes memory) {
        // Setup initial owners
        address[] memory initialOwners = new address[](1);
        initialOwners[0] = INITIAL_OWNER;

        // Setup the multisig
        return abi.encodeWithSelector(
            IMultisig.setup.selector,
            initialOwners, /* signers */
            INITIAL_THRESHOLD, /* threshold */
            SAFE_TO_L2_SETUP, /* to (used in setupModules) */
            abi.encodeWithSignature("setupToL2(address)", SAFE_L2_SINGLETON), /* data (used in setupModules) */
            SAFE_FALLBACK_HANDLER, /* fallbackHandler */
            address(0), /* paymentToken */
            0, /* payment */
            payable(address(0)) /* paymentReceiver */
        );
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

interface IMultisig {
    function VERSION() external view returns (string memory);
    function getThreshold() external view returns (uint256);
    function getOwners() external view returns (address[] memory);
    function isOwner(
        address owner
    ) external view returns (bool);

    /**
     * @notice Sets an initial storage of the Safe contract.
     * @dev This method can only be called once.
     *      If a proxy was created without setting up, anyone can call setup and claim the proxy.
     * @param _owners List of Safe owners.
     * @param _threshold Number of required confirmations for a Safe transaction.
     * @param to Contract address for optional delegate call.
     * @param data Data payload for optional delegate call.
     * @param fallbackHandler Handler for fallback calls to this contract
     * @param paymentToken Token that should be used for the payment (0 is ETH)
     * @param payment Value that should be paid
     * @param paymentReceiver Address that should receive the payment (or 0 if tx.origin)
     */
    function setup(
        address[] calldata _owners,
        uint256 _threshold,
        address to,
        bytes calldata data,
        address fallbackHandler,
        address paymentToken,
        uint256 payment,
        address paymentReceiver
    ) external;

    function addOwnerWithThreshold(address owner, uint256 threshold) external;

    function changeThreshold(
        uint256 threshold
    ) external;

    /**
     * @notice Removes the owner `owner` from the Safe and updates the threshold to `_threshold`.
     * @dev This can only be done via a Safe transaction.
     * @param prevOwner Owner that pointed to the owner to be removed in the linked list
     * @param owner Owner address to be removed.
     * @param _threshold New threshold.
     */
    function removeOwner(address prevOwner, address owner, uint256 _threshold) external;
}

interface ISafeProxyFactory {
    /**
     * @notice Deploys a new proxy with `_singleton` singleton and `saltNonce` salt. Optionally executes an initializer call to a new proxy.
     * @param _singleton Address of singleton contract. Must be deployed at the time of execution.
     * @param initializer Payload for a message call to be sent to a new proxy contract.
     * @param saltNonce Nonce that will be used to generate the salt to calculate the address of the new proxy contract.
     */
    function createProxyWithNonce(
        address _singleton,
        bytes memory initializer,
        uint256 saltNonce
    ) external returns (address proxy);
}
