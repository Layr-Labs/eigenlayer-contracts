// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

interface ISignatureUtilsMixinErrors {
    /// @notice Thrown when a signature is invalid.
    error InvalidSignature();
    /// @notice Thrown when a signature has expired.
    error SignatureExpired();
}

interface ISignatureUtilsMixinTypes {
    /// @notice Struct that bundles together a signature and an expiration time for the signature.
    /// @dev Used primarily for stack management.
    struct SignatureWithExpiry {
        // the signature itself, formatted as a single bytes object
        bytes signature;
        // the expiration timestamp (UTC) of the signature
        uint256 expiry;
    }

    /// @notice Struct that bundles together a signature, a salt for uniqueness, and an expiration time for the signature.
    /// @dev Used primarily for stack management.
    struct SignatureWithSaltAndExpiry {
        // the signature itself, formatted as a single bytes object
        bytes signature;
        // the salt used to generate the signature
        bytes32 salt;
        // the expiration timestamp (UTC) of the signature
        uint256 expiry;
    }
}

/**
 * @title The interface for common signature utilities.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
interface ISignatureUtilsMixin is ISignatureUtilsMixinErrors, ISignatureUtilsMixinTypes {
    /**
     * @notice Returns the version of the contract.
     *
     * @dev Must be overridden by the inheriting contract.
     */
    function version() external view returns (string memory);

    /**
     * @notice Returns the current EIP-712 domain separator for this contract.
     *
     * @dev The domain separator will change in the event of a fork that changes the ChainID.
     * @dev The domain separator is always recomputed to avoid the need for storage or immutable variables.
     */
    function domainSeparator() external view returns (bytes32);
}
