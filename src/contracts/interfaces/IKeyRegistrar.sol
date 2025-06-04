// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

interface IKeyRegistrarTypes {
    /// @dev Enum defining supported curve types
    enum CurveType {
        NONE,
        ECDSA,
        BN254
    }

    /// @dev Structure to store key information
    struct KeyInfo {
        bool isRegistered;
        bytes keyData; // Flexible storage for different curve types
    }
}
