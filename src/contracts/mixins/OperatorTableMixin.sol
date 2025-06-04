// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {OperatorSetLib} from "../libraries/OperatorSetLib.sol";
import "../interfaces/ICrossChainRegistry.sol";
import "../interfaces/IBN254TableCalculator.sol";
import "../interfaces/IECDSATableCalculator.sol";

/// @notice A mixin that provides functionality for encoding and decoding operator tables
/// @dev This mixin acts as a library to enable inheritance of all struct and enum types from the above interfaces
abstract contract OperatorTableMixin is
    ICrossChainRegistryTypes,
    IBN254TableCalculatorTypes,
    IECDSATableCalculatorTypes
{
    using OperatorSetLib for OperatorSet;

    enum KeyType {
        NONE,
        BN254,
        ECDSA
    }

    /// @notice Thrown when an invalid key type is provided
    error InvalidKeyType();

    /**
     * @notice Gets the operator table info from a bytes array
     * @param tableInfo The bytes containing the operator table info
     * @return operatorSet The operator set
     * @return keyType The key type
     * @return operatorSetConfig The operator set config
     */
    function _getOperatorTableInfo(
        bytes calldata tableInfo
    )
        internal
        pure
        returns (OperatorSet memory operatorSet, KeyType keyType, OperatorSetConfig memory operatorSetConfig)
    {
        (operatorSet, keyType, operatorSetConfig) = abi.decode(tableInfo, (OperatorSet, KeyType, OperatorSetConfig));
    }

    /**
     * @notice Gets the BN254 operator set info from a bytes array
     * @param tableInfo The bytes containing the operator table info
     * @return operatorSetInfo The BN254 operator set info
     */
    function _getBN254OperatorSetInfo(
        bytes calldata tableInfo
    ) internal pure returns (BN254OperatorSetInfo memory operatorSetInfo) {
        (,,, operatorSetInfo) = abi.decode(tableInfo, (OperatorSet, KeyType, OperatorSetConfig, BN254OperatorSetInfo));
    }

    /**
     * @notice Gets the ECDSA operator set info from a bytes array
     * @param tableInfo The bytes containing the operator table info
     * @return operatorSetInfo The ECDSAoperator set info
     */
    function _getECDSAOperatorSetInfo(
        bytes calldata tableInfo
    ) internal pure returns (ECDSAOperatorInfo[] memory operatorSetInfo) {
        (,,, operatorSetInfo) = abi.decode(tableInfo, (OperatorSet, KeyType, OperatorSetConfig, ECDSAOperatorInfo[]));
    }
}
