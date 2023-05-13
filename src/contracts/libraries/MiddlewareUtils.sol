// SPDX-License-Identifier: BUSL-1.1

pragma solidity =0.8.12;

/**
 * @title Library of functions shared across DataLayr.
 * @author Layr Labs, Inc.
 */
library MiddlewareUtils {
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
