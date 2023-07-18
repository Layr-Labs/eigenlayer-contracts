// SPDX-License-Identifier: BUSL-1.1

pragma solidity =0.8.12;

/**
 * @title Library of functions shared across Middlewares.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
library MiddlewareUtils {
    uint256 internal constant MAX_QUORUM_BITMAP = type(uint192).max;

    /// @notice Finds the `signatoryRecordHash`, used for fraudproofs.
    function computeSignatoryRecordHash(
        uint32 referenceBlockNumber,
        bytes32[] memory nonSignerPubkeyHashes
    ) internal pure returns (bytes32) {
        return keccak256(
            abi.encodePacked(referenceBlockNumber, nonSignerPubkeyHashes)
        );
    }
}
