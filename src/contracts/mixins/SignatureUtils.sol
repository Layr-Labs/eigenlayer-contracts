// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "@openzeppelin-upgrades/contracts/utils/cryptography/SignatureCheckerUpgradeable.sol";

import "../interfaces/ISignatureUtils.sol";

/// @title SignatureUtils
/// @notice A mixin to provide EIP-712 signature validation utilities.
/// @dev Domain name is hardcoded to "EigenLayer".
abstract contract SignatureUtils is ISignatureUtils {
    using SignatureCheckerUpgradeable for address;

    /// CONSTANTS

    /// @notice The EIP-712 typehash for the contract's domain.
    bytes32 internal constant EIP712_DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");

    /// @dev Returns the original chain ID from the time the contract was deployed.
    uint256 internal immutable _INITIAL_CHAIN_ID;

    /// @dev Returns the original domain separator from the time the contract was deployed.
    bytes32 internal immutable _INITIAL_DOMAIN_SEPARATOR;

    /// CONSTRUCTION

    constructor() {
        _INITIAL_CHAIN_ID = block.chainid;
        _INITIAL_DOMAIN_SEPARATOR = _calculateDomainSeparator();
    }

    /// EXTERNAL FUNCTIONS

    /**
     * @notice Returns the current EIP-712 domain separator for this contract.
     *
     * @dev The domain separator will change in the event of a fork that changes the ChainID.
     * @dev By introducing a domain separator the DApp developers are guaranteed that there can be no signature collision.
     * for more detailed information please read EIP-712.
     * @dev Use `_calculateDomainSeparator` rather than using this function.
     */
    function domainSeparator() public view virtual returns (bytes32) {
        /// forgefmt: disable-next-item
        return block.chainid == _INITIAL_CHAIN_ID
            // If the chain ID is the same, return the original domain separator.
            ? _INITIAL_DOMAIN_SEPARATOR
            // If the chain ID is different, return the new domain separator.
            : _calculateDomainSeparator();
    }

    /// INTERNAL HELPERS

    /// @dev Helper for calculating the contract's domain separator.
    function _calculateDomainSeparator() internal view returns (bytes32) {
        /// forgefmt: disable-next-item
        return 
            keccak256(
                abi.encode(
                    EIP712_DOMAIN_TYPEHASH, 
                    keccak256(bytes("EigenLayer")), 
                    block.chainid, 
                    address(this)
                )
            );
    }

    /// @dev Helper for creating valid EIP-712 signable digests.
    function _calculateSignableDigest(
        bytes32 hash
    ) internal view returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", domainSeparator(), hash));
    }

    /// @dev Helper for checking if a signature is valid, reverts if not valid.
    function _checkIsValidSignatureNow(
        address signer,
        bytes32 signableDigest,
        bytes memory signature,
        uint256 expiry
    ) internal view {
        require(expiry >= block.timestamp, SignatureExpired());
        require(signer.isValidSignatureNow(signableDigest, signature), InvalidSignature());
    }
}
