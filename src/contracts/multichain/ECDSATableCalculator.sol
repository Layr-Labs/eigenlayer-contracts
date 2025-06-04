// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/multichain/IECDSATableCalculator.sol";
import "../interfaces/IKeyRegistrar.sol";
import "../interfaces/IAllocationManager.sol";

/**
 * @title ECDSATableCalculator
 * @notice Contract that calculates ECDSA operator tables for a given operatorSet
 * @dev This contract uses an IOperatorWeightCalculator to get operator weights and formats them into the required table structure
 */
contract ECDSATableCalculator is IECDSATableCalculator {
    // Immutables & Constants
    /// @notice KeyRegistrar contract for managing operator keys
    IKeyRegistrar public immutable keyRegistrar;
    /// @notice AllocationManager contract for managing operator allocations
    IAllocationManager public immutable allocationManager;
    /// @notice The default lookahead blocks for the slashable stake lookup
    uint256 public constant LOOKAHEAD_BLOCKS = 100_800;

    constructor(IKeyRegistrar _keyRegistrar, IAllocationManager _allocationManager) {
        keyRegistrar = _keyRegistrar;
        allocationManager = _allocationManager;
    }

    /// @inheritdoc IECDSATableCalculator
    function calculateOperatorTable(
        OperatorSet calldata operatorSet
    ) external view virtual returns (ECDSAOperatorInfo[] memory operatorInfos) {
        return _calculateOperatorTable(operatorSet);
    }

    /// @inheritdoc IOperatorTableCalculator
    function calculateOperatorTableBytes(
        OperatorSet calldata operatorSet
    ) external view virtual returns (bytes memory operatorTableBytes) {
        return abi.encode(_calculateOperatorTable(operatorSet));
    }

    /// @inheritdoc IOperatorTableCalculator
    function getOperatorWeights(
        OperatorSet calldata operatorSet
    ) external view virtual returns (address[] memory operators, uint256[][] memory weights) {
        return _getOperatorWeights(operatorSet);
    }

    /// @inheritdoc IOperatorTableCalculator
    function getOperatorWeight(
        OperatorSet calldata operatorSet,
        address operator
    ) external view virtual returns (uint256 weight) {
        (address[] memory operators, uint256[][] memory weights) = _getOperatorWeights(operatorSet);

        // Find the index of the operator in the operators array
        for (uint256 i = 0; i < operators.length; i++) {
            if (operators[i] == operator) {
                return weights[i][0];
            }
        }

        return 0;
    }

    /**
     * @notice Get the operator weights for a given operatorSet based on the slashable stake.
     * @param operatorSet The operatorSet to get the weights for
     * @return operators The addresses of the operators in the operatorSet
     * @return weights The weights for each operator in the operatorSet, this is a 2D array where the first index is the operator
     * and the second index is the type of weight. In this case its of length 1 and returns the slashable stake for the operatorSet.
     */
    function _getOperatorWeights(
        OperatorSet calldata operatorSet
    ) internal view virtual returns (address[] memory operators, uint256[][] memory weights) {
        // Get all operators & strategies in the operatorSet
        operators = allocationManager.getMembers(operatorSet);
        IStrategy[] memory strategies = allocationManager.getStrategiesInOperatorSet(operatorSet);

        // Get the minimum slashable stake for each operator
        uint256[][] memory minSlashableStake = allocationManager.getMinimumSlashableStake({
            operatorSet: operatorSet,
            operators: operators,
            strategies: strategies,
            futureBlock: uint32(block.number + LOOKAHEAD_BLOCKS)
        });

        weights = new uint256[][](operators.length);
        for (uint256 operatorIndex = 0; operatorIndex < operators.length; ++operatorIndex) {
            // Initialize operator weights array of length 1 just for slashable stake
            weights[operatorIndex] = new uint256[](1);
            // 1. For the given operator, loop through the strategies and sum together to calculate the operator's weight for the operatorSet
            for (uint256 stratIndex = 0; stratIndex < strategies.length; ++stratIndex) {
                weights[operatorIndex][0] += minSlashableStake[operatorIndex][stratIndex];
            }
        }

        return (operators, weights);
    }

    /**
     * @notice Internal function to calculate the operator table for a given operatorSet
     * @param operatorSet The operatorSet to calculate the operator table for
     * @return operatorInfos The list of operatorInfos for the given operatorSet
     * @dev This function:
     * 1. Gets operator weights from the weight calculator
     * 2. Creates operator info entries with ECDSA public keys
     *    - assumes that the operator has a registered key
     */
    function _calculateOperatorTable(
        OperatorSet calldata operatorSet
    ) internal view returns (ECDSAOperatorInfo[] memory operatorInfos) {
        (address[] memory operators, uint256[][] memory weights) = _getOperatorWeights(operatorSet);

        operatorInfos = new ECDSAOperatorInfo[](operators.length);
        uint256 numOperators = 0;

        for (uint256 i = 0; i < operators.length; i++) {
            if (!keyRegistrar.isRegistered(operatorSet, operators[i])) {
                continue;
            }

            // Get the ECDSA public key for the operator
            bytes memory pubkey = keyRegistrar.getECDSAKey(operatorSet, operators[i]);
            address pubkeyAddress = keyRegistrar.computeAddressFromPubkey(pubkey);
            operatorInfos[numOperators] = ECDSAOperatorInfo({pubkey: pubkeyAddress, weights: weights[i]});
            ++numOperators;
        }

        // Resize the operatorInfos array to the number of operators
        assembly {
            mstore(operatorInfos, numOperators)
        }
    }
}
