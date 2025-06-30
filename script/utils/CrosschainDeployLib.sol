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
    /// -----------------------------------------------------------------------
    /// Write
    /// -----------------------------------------------------------------------

    function deployCrosschain(address deployer, bytes memory initCode, string memory name) internal returns (address) {
        require(msg.sender == deployer, "CrosschainDeployLib.deployCrosschain: deployer mismatch");
        return createx.deployCreate2(computeProtectedSalt(deployer, name), initCode);
    }

    function deployEmptyContract(
        address deployer
    ) internal returns (address) {
        require(msg.sender == deployer, "CrosschainDeployLib.deployEmptyContract: deployer mismatch");
        address computedAddress = computeCrosschainAddress({
            deployer: deployer,
            initCodeHash: keccak256(type(EmptyContract).creationCode),
            name: "EmptyContract"
        });
        if (computedAddress.code.length != 0) return computedAddress;
        return deployCrosschain(deployer, type(EmptyContract).creationCode, "EmptyContract");
    }

    function deployCrosschainProxy(
        address deployer,
        address implementation,
        string memory name
    ) internal returns (ITransparentUpgradeableProxy) {
        require(msg.sender == deployer, "CrosschainDeployLib.deployCrosschainProxy: deployer mismatch");
        return ITransparentUpgradeableProxy(
            deployCrosschain(deployer, computeUpgradeableProxyInitCode(implementation, deployer), name)
        );
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
                bytes1(uint8(1)), // Cross-chain deployments are allowed (0: false, 1: true)
                bytes11(keccak256(bytes(name)))
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
        string memory name
    ) internal view returns (address) {
        bytes32 salt = computeProtectedSalt(deployer, name);
        // CreateX's _guard function checks if first 20 bytes equal msg.sender (which is the test contract when we call it)
        // Since they don't match, CreateX falls through to: guardedSalt = keccak256(abi.encode(salt))
        bytes32 guardedSalt = keccak256(abi.encode(deployer, salt));
        return createx.computeCreate2Address(guardedSalt, initCodeHash);
    }

    /*
     * @notice Returns the predicted address of a `TransparentUpgradeableProxy` deployed with CreateX.
     */
    function computeCrosschainUpgradeableProxyAddress(
        address implementation,
        address admin,
        string memory name
    ) internal view returns (address) {
        return computeCrosschainAddress(admin, keccak256(computeUpgradeableProxyInitCode(implementation, admin)), name);
    }
}
