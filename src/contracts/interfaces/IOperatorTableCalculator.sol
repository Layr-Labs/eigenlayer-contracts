// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../libraries/BN254.sol";
import "../libraries/OperatorSetLib.sol";

interface IOperatorTableCalculatorTypes {
    /// BN254 Table Calculator Types
    /**
     * @notice A struct that contains information about a single operator
     * @param pubkey The G1 public key of the operator.
     * @param weights The weights of the operator for a single operatorSet.
     * @dev The `weights` array can be defined as a list of arbitrary groupings. For example,
     * it can be [slashable_stake, delegated_stake, strategy_i_stake, ...]
     */
    struct BN254OperatorInfo {
        BN254.G1Point pubkey;
        uint256[] weights;
    }

    /**
     * @notice A struct that contains information about all operators for a given operatorSet
     * @param operatorInfoTreeRoot The root of the operatorInfo tree.
     * @param numOperators The number of operators in the operatorSet.
     * @param aggregatePubkey The aggregate G1 public key of the operators in the operatorSet.
     * @param totalWeights The total weights of the operators in the operatorSet.
     *
     * @dev The operatorInfoTreeRoot is the root of a merkle tree that contains the operatorInfos for each operator in the operatorSet.
     * It is calculated in this function and used by the `IBN254CertificateVerifier` to verify stakes against the non-signing operators
     *
     * @dev Retrieval of the `aggregatePubKey` depends on maintaining a key registry contract, see `BLSAPKRegistry` for an example implementation.
     *
     * @dev The `totalWeights` array should be the same length as each individual `weights` array in `operatorInfos`.
     */
    struct BN254OperatorSetInfo {
        bytes32 operatorInfoTreeRoot;
        uint256 numOperators;
        BN254.G1Point aggregatePubkey;
        uint256[] totalWeights;
    }

    /// ECDSA Table Calculator Types

    /**
     * @notice A struct that contains information about a single operator
     * @param pubkey The address of the signing ECDSA key of the operator and not the operator address itself.
     * This is read from the KeyRegistrar contract.
     * @param weights The weights of the operator for a single operatorSet
     * @dev The `weights` array can be defined as a list of arbitrary groupings. For example,
     * it can be [slashable_stake, delegated_stake, strategy_i_stake, ...]
     */
    struct ECDSAOperatorInfo {
        address pubkey;
        uint256[] weights;
    }
}

/// @notice A base operator table calculator that all operator table calculators (ECDSA, BN254) must implement
interface IOperatorTableCalculator {
    /**
     * @notice Calculates the operator table for a given operatorSet
     * @param operatorSet the operatorSet to calculate the operator table for
     * @return operatorTableBytes the operatorTableBytes for the given operatorSet
     */
    function calculateOperatorTableBytes(
        OperatorSet calldata operatorSet
    ) external view returns (bytes memory operatorTableBytes);

    /**
     * @notice Get the operator weights for a given operatorSet based on the slashable stake.
     * @param operatorSet The operatorSet to get the weights for
     * @return operators The addresses of the operators in the operatorSet
     * @return weights The weights for each operator in the operatorSet, this is a 2D array where the first index is the operator
     * and the second index is the type of weight. In this case its of length 1 and returns the slashable stake for the operatorSet.
     */
    function getOperatorWeights(
        OperatorSet calldata operatorSet
    ) external view returns (address[] memory operators, uint256[][] memory weights);

    /**
     * @notice Get the weight for a given operator in a given operatorSet based on the slashable stake.
     * @param operatorSet The operatorSet to get the weight for
     * @param operator The operator to get the weight for
     * @return weight The weight for the operator in the operatorSet
     */
    function getOperatorWeight(
        OperatorSet calldata operatorSet,
        address operator
    ) external view returns (uint256 weight);
}
