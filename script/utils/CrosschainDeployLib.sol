// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

ICreateX constant createx = ICreateX(0xba5Ed099633D3B313e4D5F7bdc1305d3c28ba5Ed);

interface ICreateX {
    function deployCreate2(bytes32 salt, bytes memory initCode) external payable returns (address newContract);
    function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) external view returns (address computedAddress);
}

library CrosschainDeployLib {
    /// @notice Deploys a contract with CreateX.
    /// @param initCode The initialization code for the contract.
    /// @param salt The entropy for the deployment.
    /// @return The address of the deployed contract.
    function deployCrosschain(bytes memory initCode, bytes11 salt) internal returns (address) {
        return createx.deployCreate2(computeSalt(msg.sender, salt), initCode);
    }

    /// @notice Deploys a contract with CreateX.
    /// @param initCode The initialization code for the contract.
    /// @param value The amount of ETH to send with the deployment.
    /// @return The address of the deployed contract.
    function deployCrosschain(bytes memory initCode) internal returns (address) {
        return deployCrosschain(initCode, bytes11(0xE15E4));
    }

    /// @notice Returns the cross-chain address of a contract.
    /// @param eoa The EOA address of the deployer.
    /// @param salt The entropy for the deployment.
    /// @param initCodeHash The hash of the initialization code.
    /// @return The predicted address of the contract.
    function getCrosschainAddress(address eoa, bytes11 salt, bytes32 initCodeHash) internal view returns (address) {
        return createx.computeCreate2Address(computeSalt(eoa, salt), initCodeHash);
    }

    /// @notice Returns the encoded salt for a CreateX deployment.
    /// @param eoa The EOA address of the deployer.
    /// @param salt The entropy for the deployment.
    /// @return The encoded salt.
    /// Deployer EOA (20 bytes) | Cross-chain flag (1 byte) | Entropy (11 bytes)
    /// 0xbebebebebebebebebebebebebebebebebebebebe|ff|1212121212121212121212
    function computeSalt(address eoa, bytes11 salt) internal view returns (bytes32) {
        return bytes32(
            bytes.concat(
                bytes20(eoa),
                bytes1(1), // Cross-chain deployments are allowed (0: false, 1: true)
                salt
            )
        );
    }
}
