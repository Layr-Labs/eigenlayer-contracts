// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

/**
 * @title The interface for common signature utilities.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
interface ISignatureUtils {
    // @notice Struct that bundles together a signature and an expiration time for the signature. Used primarily for stack management.
    struct SignatureWithExpiry {
        // the signature itself, formatted as a single bytes object
        bytes signature;
        // the expiration timestamp (UTC) of the signature
        uint256 expiry;
    }

    // @notice Struct that bundles together a signature, a salt for uniqueness, and an expiration time for the signature. Used primarily for stack management.
    struct SignatureWithSaltAndExpiry {
        // the signature itself, formatted as a single bytes object
        bytes signature;
        // the salt used to generate the signature
        bytes32 salt;
        // the expiration timestamp (UTC) of the signature
        uint256 expiry;
    }
}