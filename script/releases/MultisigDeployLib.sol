// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

/// @dev A library for deploying Safe multisigs
library MultisigDeployLib {
    using MultisigDeployLib for *;

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
        // addresses taken from https://github.com/safe-global/safe-smart-account/blob/main/CHANGELOG.md#expected-addresses-with-deterministic-deployment-proxy-default
        // NOTE: double check these addresses are correct on each chain
        address safeFactory = 0x4e1DCf7AD4e460CfD30791CCC4F9c8a4f820ec67;
        // @dev Gnosis safe L2 singleton - This value is ignored on L1 deployments. 
        address safeSingleton = 0x29fcB43b46531BcA003ddC8FCB67FFE91900C762; 
        address safeFallbackHandler = 0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99;

        bytes memory emptyData;

        bytes memory initializerData = abi.encodeWithSignature(
            "setup(address[],uint256,address,bytes,address,address,uint256,address)",
            initialOwners, /* signers */
            initialThreshold, /* threshold */
            address(0), /* to (used in setupModules) */
            emptyData, /* data (used in setupModules) */
            safeFallbackHandler,
            address(0), /* paymentToken */
            0, /* payment */
            payable(address(0)) /* paymentReceiver */
        );

        bytes memory calldataToFactory =
            abi.encodeWithSignature("createProxyWithNonce(address,bytes,uint256)", safeSingleton, initializerData, salt);

        (bool success, bytes memory returndata) = safeFactory.call(calldataToFactory);
        require(success, "multisig deployment failed");
        address deployedMultisig = abi.decode(returndata, (address));
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

    function isOwner(
        address multisig,
        address owner
    ) internal view returns (bool) {
        return IMultisig(multisig).isOwner(owner);
    }
}

interface IMultisig {
    function getThreshold() external view returns (uint256);
    function getOwners() external view returns (address[] memory);
    function isOwner(address owner) external view returns (bool);
}