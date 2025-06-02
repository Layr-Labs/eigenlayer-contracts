// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {BN254} from "../libraries/BN254.sol";
import {OperatorSet} from "../libraries/OperatorSetLib.sol";

interface IBN254TableCalculatorTypes {
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
}

interface IBN254TableCalculator is IBN254TableCalculatorTypes {
    /**
     * @notice calculates the operatorInfos for a given operatorSet
     * @param operatorSet the operatorSet to calculate the operator table for
     * @return operatorSetInfo the operatorSetInfo for the given operatorSet
     * @dev The output of this function is converted to bytes via the `calculateOperatorTableBytes` function
     */
    function calculateOperatorTable(
        OperatorSet calldata operatorSet
    ) external view returns (BN254OperatorSetInfo memory operatorSetInfo);

    /**
     * @notice calculates the operatorInfos for a given operatorSet
     * @param operatorSet the operatorSet to calculate the operator table for
     * @return operatorTableBytes the operatorTableBytes for the given operatorSet
     */
    function calculateOperatorTableBytes(
        OperatorSet calldata operatorSet
    ) external view returns (bytes memory operatorTableBytes);

    /**
     * @notice Get the operatorInfos for a given operatorSet
     * @param operatorSet the operatorSet to get the operatorInfos for
     * @return operatorInfos the operatorInfos for the given operatorSet
     */
    function getOperatorInfos(
        OperatorSet calldata operatorSet
    ) external view returns (BN254OperatorInfo[] memory operatorInfos);
}
