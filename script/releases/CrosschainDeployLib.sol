// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "src/test/mocks/EmptyContract.sol";

/// @dev https://github.com/pcaversaccio/createx/tree/main
ICreateX constant createx = ICreateX(0xba5Ed099633D3B313e4D5F7bdc1305d3c28ba5Ed);

interface ICreateX {
    function deployCreate2(bytes32 salt, bytes memory initCode) external payable returns (address newContract);
    function computeCreate2Address(
        bytes32 salt,
        bytes32 initCodeHash
    ) external view returns (address computedAddress);
}

bytes11 constant EMPTY_CONTRACT_SALT = bytes11(uint88(0xffffffffffffffffffffff));

library CrosschainDeployLib {
    using CrosschainDeployLib for *;

    /// -----------------------------------------------------------------------
    /// Write
    /// -----------------------------------------------------------------------

    /**
     * @notice Deploys a crosschain empty contract.
     * @dev The empty contract MUST stay consistent across all chains/deployments.
     * @dev The empty contract MUST always be deployed with the same salt.
     */
    function deployEmptyContract(address deployer) internal returns (address) {
        address computedAddress =
            computeCrosschainAddress(deployer, keccak256(type(EmptyContract).creationCode), EMPTY_CONTRACT_SALT);
        if (computedAddress.code.length != 0) return computedAddress;
        return type(EmptyContract).creationCode.deployCrosschain(deployer, EMPTY_CONTRACT_SALT);
    }

    /*
     * @notice Deploys a crosschain contract with CreateX.
     *
     * @dev Example usage:
     * ```solidity
     * type(EmptyContract).creationCode.deployCrosschain(EMPTY_CONTRACT_SALT)
     * ```
     */
    function deployCrosschain(bytes memory initCode, address deployer, bytes11 salt) internal returns (address) {
        return createx.deployCreate2(computeProtectedSalt(deployer, salt), initCode);
    }

    // /*
    //  * @notice Deploys a crosschain contract using CreateX with the `DEFAULT_SALT`.
    //  *
    //  * @dev Example usage:
    //  * ```solidity
    //  * address emptyContract = type(EmptyContract).creationCode.deployCrosschain();
    //  * ```
    //  */
    // function deployCrosschain(
    //     bytes memory initCode
    // ) internal returns (address) {
    //     return deployCrosschain(initCode, EMPTY_CONTRACT_SALT);
    // }

    /*
     * @notice Deploys a crosschain `TransparentUpgradeableProxy` using CreateX.
     * @dev The initial admin is msg.sender.
     * @dev The implementation MUST also be deterministic to ensure the contract can be deployed on all chains.
     * @dev The salt MUST be unique for each proxy deployment sharing the same implementation otherwise address collisions WILL occur.
     *
     * @dev Example usage:
     * ```solidity
     * bytes11 salt = bytes11(uint88(0xffffffffffffffffffffff));
     * address emptyContract = type(EmptyContract).creationCode.deployCrosschain();
     * address proxy = emptyContract.deployCrosschainProxy(salt); 
     * ITransparentUpgradeableProxy(address(proxy)).upgradeTo(address(implementation));
     * ITransparentUpgradeableProxy(address(proxy)).changeAdmin(address(admin));
     * ```
     */
    // function deployCrosschainProxy(
    //     address implementation,
    //     bytes11 salt
    // ) internal returns (ITransparentUpgradeableProxy) {
    //     return ITransparentUpgradeableProxy(
    //         deployCrosschain(computeUpgradeableProxyInitCode(implementation, msg.sender), salt)
    //     );
    // }

    // /**
    //  * @notice Deploys a crosschain name salted `TransparentUpgradeableProxy` using CreateX.
    //  */
    // function deployCrosschainProxy(
    //     address implementation,
    //     string memory name
    // ) internal returns (ITransparentUpgradeableProxy) {
    //     return deployCrosschainProxy(implementation, bytes11(keccak256(bytes(name))));
    // }

    /// -----------------------------------------------------------------------
    /// Helpers
    /// -----------------------------------------------------------------------

    /*
     * @notice Returns an encoded CreateX salt.
     * @dev The deployer is the only account that can use this salt via CreateX hence the name "protected".
     * @dev The salt is structured as: Deployer EOA (20 bytes) | Cross-chain flag (1 byte) | Entropy (11 bytes)
     * @dev Example: 0xbebebebebebebebebebebebebebebebebebebebe|ff|1212121212121212121212
     */
    function computeProtectedSalt(address deployer, bytes11 salt) internal pure returns (bytes32) {
        return bytes32(
            bytes.concat(
                bytes20(deployer),
                bytes1(uint8(1)), // Cross-chain deployments are allowed (0: false, 1: true)
                bytes11(salt)
            )
        );
    }

    /*
     * @notice Returns the initialization code for a transparent upgradeable proxy.
     * @dev The returned init code does not include metadata typically appended by the compiler.
     */
    function computeUpgradeableProxyInitCode(
        address implementation,
        address admin
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(type(TransparentUpgradeableProxy).creationCode, abi.encode(implementation, admin, ""));
    }

    /*
     * @notice Returns the predicted address of a contract deployed with CreateX.
     */
    function computeCrosschainAddress(
        address deployer,
        bytes32 initCodeHash,
        bytes11 salt
    ) internal view returns (address) {
        return createx.computeCreate2Address(computeProtectedSalt(deployer, salt), initCodeHash);
    }

    /*
     * @notice Returns the predicted address of a `TransparentUpgradeableProxy` deployed with CreateX.
     */
    function computeCrosschainUpgradeableProxyAddress(
        address implementation,
        address admin,
        bytes11 salt
    ) internal view returns (address) {
        return computeCrosschainAddress(admin, keccak256(computeUpgradeableProxyInitCode(implementation, admin)), salt);
    }
}
