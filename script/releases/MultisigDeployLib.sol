// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

/// @notice A library for deploying Safe multisigs
/// @dev Double check the constant addresses are correct on each chain prior to deployment
library MultisigDeployLib {
    using MultisigDeployLib for *;

    /// @dev Safe proxy factory, this should be the same for all chains
    /// @dev See https://github.com/safe-global/safe-smart-account/blob/main/CHANGELOG.md#version-141 for more details
    /// @dev DOUBLE CHECK THESE ADDRESSES ARE CORRECT ON EACH CHAIN PRIOR TO DEPLOYMENT
    address public constant SAFE_PROXY_FACTORY = 0x4e1DCf7AD4e460CfD30791CCC4F9c8a4f820ec67;
    address public constant SAFE_SINGLETON = 0x41675C099F32341bf84BFc5382aF534df5C7461a;
    address public constant SAFE_FALLBACK_HANDLER = 0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99;

    /// @dev L2 Singletons still need to be passed into the L1 deployment
    /// @dev `SAFE_TO_L2_SETUP` does a no-op if the chain is mainnet
    /// @dev See: https://github.com/safe-global/safe-smart-account/blob/0095f1aa113255d97b476e625760514cc7d10982/contracts/libraries/SafeToL2Setup.sol#L59-L69
    /// @dev Even mainnet testnet chains (eg. sepolia, hoodi) should use the SAFE_L2_SINGLETON
    address public constant SAFE_L2_SINGLETON = 0x29fcB43b46531BcA003ddC8FCB67FFE91900C762;
    address public constant SAFE_TO_L2_SETUP = 0xBD89A1CE4DDe368FFAB0eC35506eEcE0b1fFdc54;
    bytes public constant SETUP_TO_L2_DATA = abi.encodeWithSignature("setupToL2(address)", SAFE_L2_SINGLETON);

    /**
     * @notice Deploys a Safe multisig
     * @param initialOwners The initial owners of the multisig
     * @param initialThreshold The initial threshold of the multisig
     * @param salt The salt to use for the deployment
     * @return The address of the deployed multisig
     */
    function deployMultisig(
        address[] memory initialOwners,
        uint256 initialThreshold,
        uint256 salt
    ) internal returns (address) {
        bytes memory initializerData = abi.encodeWithSelector(
            IMultisig.setup.selector,
            initialOwners, /* signers */
            initialThreshold, /* threshold */
            SAFE_TO_L2_SETUP, /* to (used in setupModules) */
            SETUP_TO_L2_DATA, /* data (used in setupModules) */
            SAFE_FALLBACK_HANDLER, /* fallbackHandler */
            address(0), /* paymentToken */
            0, /* payment */
            payable(address(0)) /* paymentReceiver */
        );

        address deployedMultisig =
            ISafeProxyFactory(SAFE_PROXY_FACTORY).createProxyWithNonce(SAFE_SINGLETON, initializerData, salt);
        require(deployedMultisig != address(0), "something wrong in multisig deployment, zero address returned");
        return deployedMultisig;
    }

    function getThreshold(
        address multisig
    ) internal view returns (uint256) {
        return IMultisig(multisig).getThreshold();
    }

    function getOwners(
        address multisig
    ) internal view returns (address[] memory) {
        return IMultisig(multisig).getOwners();
    }

    function isOwner(address multisig, address owner) internal view returns (bool) {
        return IMultisig(multisig).isOwner(owner);
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
