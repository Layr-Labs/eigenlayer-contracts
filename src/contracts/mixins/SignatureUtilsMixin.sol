// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "@openzeppelin-upgrades/contracts/utils/ShortStringsUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/utils/cryptography/SignatureCheckerUpgradeable.sol";

import "../interfaces/ISignatureUtilsMixin.sol";
import "./SemVerMixin.sol";

/// @dev The EIP-712 domain type hash used for computing the domain separator
///      See https://eips.ethereum.org/EIPS/eip-712#definition-of-domainseparator
bytes32 constant EIP712_DOMAIN_TYPEHASH =
    keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)");

/// @title SignatureUtilsMixin
/// @notice A mixin contract that provides utilities for validating signatures according to EIP-712 and EIP-1271 standards.
/// @dev The domain name is hardcoded to "EigenLayer". This contract implements signature validation functionality that can be
///      inherited by other contracts. The domain separator uses the major version (e.g., "v1") to maintain EIP-712
///      signature compatibility across minor and patch version updates.
abstract contract SignatureUtilsMixin is ISignatureUtilsMixin, SemVerMixin {
    using SignatureCheckerUpgradeable for address;

    /// @notice Initializes the contract with a semantic version string.
    /// @param _version The SemVer-formatted version string (e.g., "v1.1.1") to use for this contract's domain separator.
    /// @dev Version should follow SemVer 2.0.0 format with 'v' prefix: vMAJOR.MINOR.PATCH.
    ///      Only the major version component is used in the domain separator to maintain signature compatibility
    ///      across minor and patch version updates.
    constructor(
        string memory _version
    ) SemVerMixin(_version) {}

    /// EXTERNAL FUNCTIONS ///

    /// @inheritdoc ISignatureUtilsMixin
    function domainSeparator() public view virtual returns (bytes32) {
        // forgefmt: disable-next-item
        return 
            keccak256(
                abi.encode(
                    EIP712_DOMAIN_TYPEHASH, 
                    keccak256(bytes("EigenLayer")),
                    keccak256(bytes(_majorVersion())),
                    block.chainid, 
                    address(this)
                )
            );
    }

    /// INTERNAL HELPERS ///

    /// @notice Creates a digest that can be signed using EIP-712.
    /// @dev Prepends the EIP-712 prefix ("\x19\x01") and domain separator to the input hash.
    ///      This follows the EIP-712 specification for creating structured data hashes.
    ///      See https://eips.ethereum.org/EIPS/eip-712#specification.
    /// @param hash The hash of the typed data to be signed.
    /// @return The complete digest that should be signed according to EIP-712.
    function _calculateSignableDigest(
        bytes32 hash
    ) internal view returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", domainSeparator(), hash));
    }

    /// @notice Validates a signature against a signer and digest, with an expiry timestamp.
    /// @dev Reverts if the signature is invalid or expired. Uses EIP-1271 for smart contract signers.
    ///      For EOA signers, validates ECDSA signatures directly.
    ///      For contract signers, calls isValidSignature according to EIP-1271.
    ///      See https://eips.ethereum.org/EIPS/eip-1271#specification.
    /// @param signer The address that should have signed the digest.
    /// @param signableDigest The digest that was signed, created via _calculateSignableDigest.
    /// @param signature The signature bytes to validate.
    /// @param expiry The timestamp after which the signature is no longer valid.
    function _checkIsValidSignatureNow(
        address signer,
        bytes32 signableDigest,
        bytes memory signature,
        uint256 expiry
    ) internal view {
        // First, check if the signature has expired by comparing the expiry timestamp
        // against the current block timestamp.
        require(expiry >= block.timestamp, SignatureExpired());

        // Next, verify that the signature is valid for the given signer and digest.
        // For EOA signers, this performs standard ECDSA signature verification.
        // For contract signers, this calls the EIP-1271 isValidSignature method.
        require(signer.isValidSignatureNow(signableDigest, signature), InvalidSignature());
    }
}
