// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "../interfaces/multichain/IECDSATableCalculator.sol";
import "../interfaces/multichain/IOperatorWeightCalculator.sol";
import "../interfaces/IKeyRegistrar.sol";
import "./ECDSATableCalculatorStorage.sol";

/**
 * @title ECDSATableCalculator
 * @notice Contract that calculates ECDSA operator tables for a given operatorSet
 * @dev This contract uses an IOperatorWeightCalculator to get operator weights and formats them into the required table structure
 */
contract ECDSATableCalculator is Initializable, OwnableUpgradeable, ECDSATableCalculatorStorage {
    constructor(
        IOperatorWeightCalculator _operatorWeightCalculator,
        IKeyRegistrar _keyRegistrar
    ) ECDSATableCalculatorStorage(_operatorWeightCalculator, _keyRegistrar) {
        _disableInitializers();
    }

    function initialize(
        address initialOwner
    ) external initializer {
        _transferOwnership(initialOwner);
    }

    /// @inheritdoc IECDSATableCalculator
    function calculateOperatorTable(
        OperatorSet calldata operatorSet
    ) external view virtual returns (ECDSAOperatorInfo[] memory operatorInfos) {
        return _calculateOperatorTable(operatorSet);
    }

    /// @inheritdoc IECDSATableCalculator
    function calculateOperatorTableBytes(
        OperatorSet calldata operatorSet
    ) external view virtual returns (bytes memory operatorTableBytes) {
        return abi.encode(_calculateOperatorTable(operatorSet));
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
        (address[] memory operators, uint256[][] memory weights) =
            operatorWeightCalculator.getOperatorWeights(operatorSet);

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
