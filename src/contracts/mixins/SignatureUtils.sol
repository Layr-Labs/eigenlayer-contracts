// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "@openzeppelin-upgrades/contracts/utils/cryptography/SignatureCheckerUpgradeable.sol";

import "../interfaces/ISignatureUtils.sol";

bytes32 constant EIP712_DOMAIN_TYPEHASH =
    keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");

/// @title SignatureUtils
/// @notice A mixin to provide EIP-712 signature validation utilities.
/// @dev Domain name is hardcoded to "EigenLayer".
abstract contract SignatureUtils is ISignatureUtils {
    using SignatureCheckerUpgradeable for address;

    /// EXTERNAL FUNCTIONS

    /**
     * @notice Returns the current EIP-712 domain separator for this contract.
     *
     * @dev The domain separator will change in the event of a fork that changes the ChainID.
     * @dev The domain separator is always recomputed to avoid the need for storage or immutable variables.
     */
    function domainSeparator() public view virtual returns (bytes32) {
        // forgefmt: disable-next-item
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

    /// INTERNAL HELPERS

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
