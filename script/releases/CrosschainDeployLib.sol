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

library CrosschainDeployLib {
    using CrosschainDeployLib for *;

    /// -----------------------------------------------------------------------
    /// Write
    /// -----------------------------------------------------------------------

    /*
     * @notice Deploys a crosschain empty contract.
     * @dev The empty contract MUST stay consistent across all chains/deployments.
     * @dev The empty contract MUST always be deployed with the same salt.
     */
    function deployEmptyContract(
        address deployer
    ) internal returns (address) {
        return _deployCrosschain(deployer, type(EmptyContract).creationCode, type(EmptyContract).name);
    }

    /*
     * @notice Deploys a crosschain `TransparentUpgradeableProxy` using CreateX.
     * @dev The initial admin is the deployer.
     * @dev The implementation MUST also be deterministic to ensure the contract can be deployed on all chains.
     * @dev The salt MUST be unique for each proxy deployment sharing the same implementation otherwise address collisions WILL occur.
     * @dev The `admin` is also assumed to be the deployer.
     *
     * @dev Example usage:
     * ```solidity
     * bytes11 salt = bytes11(uint88(0xffffffffffffffffffffff));
     * address emptyContract = type(EmptyContract).creationCode.deployCrosschain(deployer);
     * address proxy = emptyContract.deployCrosschainProxy(deployer, salt); 
     * ITransparentUpgradeableProxy(address(proxy)).upgradeTo(address(implementation));
     * ITransparentUpgradeableProxy(address(proxy)).changeAdmin(address(admin));
     * ```
     */
    function deployCrosschainProxy(
        address adminAndDeployer,
        address implementation,
        string memory name
    ) internal returns (ITransparentUpgradeableProxy) {
        return ITransparentUpgradeableProxy(
            _deployCrosschain(adminAndDeployer, computeUpgradeableProxyInitCode(implementation, adminAndDeployer), name)
        );
    }

    /*
     * @notice Deploys a crosschain contract with CreateX.
     *
     * @dev Example usage:
     * ```solidity
     * type(EmptyContract).creationCode.deployCrosschain(deployer, EMPTY_CONTRACT_SALT)
     * ```
     */
    function _deployCrosschain(address deployer, bytes memory initCode, string memory name) private returns (address) {
        return createx.deployCreate2(computeProtectedSalt(deployer, name), initCode);
    }

    /// -----------------------------------------------------------------------
    /// Helpers
    /// -----------------------------------------------------------------------

    /*
     * @notice Returns an encoded CreateX salt.
     * @dev The deployer is the only account that can use this salt via CreateX hence the name "protected".
     * @dev The salt is structured as: Deployer EOA (20 bytes) | Cross-chain flag (1 byte) | Entropy (11 bytes)
     * @dev Example: 0xbebebebebebebebebebebebebebebebebebebebe|ff|1212121212121212121212
     */
    function computeProtectedSalt(address deployer, string memory name) internal pure returns (bytes32) {
        return bytes32(
            bytes.concat(
                bytes20(deployer),
                bytes1(uint8(0)), // Cross-chain redeploy protection enabled (0: false, 1: true)
                bytes11(keccak256(bytes(name))) // salt
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
     * @notice Computes the deterministic address where a crosschain contract will be deployed.
     * @dev Uses CreateX's computeCreate2Address with a protected salt to ensure deterministic deployment.
     * @param deployer The address of the deployer account.
     * @param initCodeHash The keccak256 hash of the contract's initialization code.
     * @param name The name used to generate the protected salt.
     * @return The deterministic address where the contract will be deployed.
     */
    function computeCrosschainAddress(
        address deployer,
        bytes32 initCodeHash,
        string memory name
    ) internal view returns (address) {
        return createx.computeCreate2Address(
            keccak256(abi.encodePacked(bytes32(uint256(uint160(deployer))), computeProtectedSalt(deployer, name))),
            initCodeHash
        );
    }

    /*
     * @notice Computes the deterministic address where a crosschain upgradeable proxy will be deployed.
     * @dev Computes the init code hash for a transparent upgradeable proxy and calls computeCrosschainAddress.
     * @param adminAndDeployer The address that will be both the admin and deployer of the proxy.
     * @param implementation The address of the implementation contract.
     * @param name The name used to generate the protected salt.
     * @return The deterministic address where the upgradeable proxy will be deployed.
     */
    function computeCrosschainUpgradeableProxyAddress(
        address adminAndDeployer,
        address implementation,
        string memory name
    ) internal view returns (address) {
        return computeCrosschainAddress(
            adminAndDeployer, keccak256(computeUpgradeableProxyInitCode(implementation, adminAndDeployer)), name
        );
    }
}
