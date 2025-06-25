## OperatorTableCalculator

| File | Type | Notes |
| -------- | -------- | 
| [`ECDSATableCalculatorBase.sol`](../../../src/contracts/multichain/ECDSATableCalculatorBase.sol) | Abstract | Base functionality for ECDSA operator tables |
| [`BN254TableCalculatorBase.sol`](../../../src/contracts/multichain/BN254TableCalculatorBase.sol) | Abstract | Base functionality for BN254 operator tables |

Interfaces:

| File | Notes |
| -------- | -------- |
| [`IOperatorTableCalculator.sol`](../../../src/contracts/interfaces/IOperatorTableCalculator.sol) | Base interface for all calculators |
| [`IECDSATableCalculator.sol`](../../../src/contracts/interfaces/IECDSATableCalculator.sol) | ECDSA-specific interface |
| [`IBN254TableCalculator.sol`](../../../src/contracts/interfaces/IBN254TableCalculator.sol) | BN254-specific interface |

---

## Overview

The OperatorTableCalculator contracts are responsible for calculating operator tables that are used in the multichain protocol. These tables contain operator information (public keys and weights) that are read from  the [`CrossChainRegistry`](./CrossChainRegistry.md) to transport stakes.

The base contracts (`ECDSATableCalculatorBase` and `BN254TableCalculatorBase`) provide the core logic for table calculation, while leaving weight calculation as an unimplemented method to be implemented by derived contracts. 

---

## ECDSATableCalculatorBase

The `ECDSATableCalculatorBase` provides base functionality for calculating ECDSA operator tables. It handles operator key retrieval and table construction. 

### Core Functions

#### `calculateOperatorTable`

```solidity
/**
 * @notice A struct that contains information about a single operator
 * @param pubkey The address of the signing ECDSA key of the operator and not the operator address itself.
 * This is read from the KeyRegistrar contract.
 * @param weights The weights of the operator for a single operatorSet
 * @dev The `weights` array can be defined as a list of arbitrary groupings. For example,
 * it can be [slashable_stake, delegated_stake, strategy_i_stake, ...]
 * @dev The `weights` array should be the same length for each operator in the operatorSet.
 */
struct ECDSAOperatorInfo {
    address pubkey;
    uint256[] weights;
}

/**
 * @notice calculates the operatorInfos for a given operatorSet
 * @param operatorSet the operatorSet to calculate the operator table for
 * @return operatorInfos the list of operatorInfos for the given operatorSet
 * @dev The output of this function is converted to bytes via the `calculateOperatorTableBytes` function
 */
function calculateOperatorTable(
    OperatorSet calldata operatorSet
) external view returns (ECDSAOperatorInfo[] memory operatorInfos);
```

Calculates and returns an array of `ECDSAOperatorInfo` structs containing public keys and weights for all operators in the operatorSet who have registered ECDSA keys.

*Effects*:
* None (view function)

*Process*:
* Calls `_getOperatorWeights` to retrieve operator addresses and their weights
* For each operator with a registered ECDSA key:
  * Retrieves the ECDSA address (public key) from the `KeyRegistrar`
  * Creates an `ECDSAOperatorInfo` struct with the public key and weights
* Returns only operators with registered keys

#### `calculateOperatorTableBytes`

```solidity
/**
 * @notice Calculates the operator table bytes for a given operatorSet
 * @param operatorSet The operatorSet to calculate the operator table for
 * @return operatorTableBytes The encoded operator table bytes
 */
function calculateOperatorTableBytes(
    OperatorSet calldata operatorSet
) external view returns (bytes memory operatorTableBytes);
```

Returns the ABI-encoded bytes representation of the operator table, which is used by the `CrossChainRegistry` to calculate the operatorTable. 

*Returns*:
* ABI-encoded array of `ECDSAOperatorInfo` structs

### Abstract Methods

#### `_getOperatorWeights`

```solidity
/**
 * @notice Abstract function to get the operator weights for a given operatorSet
 * @param operatorSet The operatorSet to get the weights for
 * @return operators The addresses of the operators in the operatorSet
 * @return weights The weights for each operator in the operatorSet
 */
function _getOperatorWeights(
    OperatorSet calldata operatorSet
) internal view virtual returns (address[] memory operators, uint256[][] memory weights);
```

Must be implemented by derived contracts to define the weight calculation strategy. The weights are returned as a 2D array indexed by operator. It is up to the AVS to define each operator's weights array. Some examples include:

- A single array evaluated purely on slashable stake `[slashable_stake]` o
- An array of 2 values can be used for evaluating on slashable and delegated stake `[slashable_stake, delegated_stake]`
- An array of several values can be used for evaluating stake on multiple  strategies `[slashable_stake_stETH, slashable_stake_USDC, slashable_stake_EIGEN]` 

In addition, an AVS can have custom calculation methodologies that include:
- Capping the stake of an operator
- Using oracles to price stake

An example integration is defined by [`ECDSATableCalculator`](../../../src/contracts/multichain/ECDSATableCalculator.sol)

---

## BN254TableCalculatorBase

The `BN254TableCalculatorBase` provides base functionality for calculating BN254 operator tables.

### Core Functions

#### `calculateOperatorTable`

```solidity
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
 * @param operatorInfoTreeRoot The root of the operatorInfo tree. Each leaf is a `BN254OperatorInfo` struct
 * @param numOperators The number of operators in the operatorSet.
 * @param aggregatePubkey The aggregate G1 public key of the operators in the operatorSet.
 * @param totalWeights The total weights of the operators in the operatorSet.
 *
 * @dev The operatorInfoTreeRoot is the root of a merkle tree that contains the operatorInfos for each operator in the operatorSet.
 * It is calculated in this function and used by the `IBN254CertificateVerifier` to verify stakes against the non-signing operators
 *
 * @dev Retrieval of the `aggregatePubKey` depends on maintaining a key registry contract, see `BN254TableCalculatorBase` for an example implementation.
 *
 * @dev The `totalWeights` array should be the same length as each individual `weights` array in `operatorInfos`.
 */
struct BN254OperatorSetInfo {
    bytes32 operatorInfoTreeRoot;
    uint256 numOperators;
    BN254.G1Point aggregatePubkey;
    uint256[] totalWeights;
}

/**
 * @notice calculates the operatorInfos for a given operatorSet
 * @param operatorSet the operatorSet to calculate the operator table for
 * @return operatorSetInfo the operatorSetInfo for the given operatorSet
 * @dev The output of this function is converted to bytes via the `calculateOperatorTableBytes` function
 */
function calculateOperatorTable(
    OperatorSet calldata operatorSet
) external view returns (BN254OperatorSetInfo memory operatorSetInfo);
```

Calculates and returns a `BN254OperatorSetInfo` struct containing:
- A merkle tree root of operator information
- The total number of operators
- An aggregate BN254 public key
- Total weights across all operators

*Effects*:
* None (view function)

*Process*:
* Calls `_getOperatorWeights` to retrieve operator addresses and their weights
* For each operator with a registered BN254 key:
  * Retrieves the BN254 G1 point from the `KeyRegistrar`
  * Adds the operator's weights to the total weights
  * Creates a merkle leaf from the operator info
  * Adds the G1 point to the aggregate public key
* Constructs a merkle tree from all operator info leaves
* Returns the complete operator set information

BN254 tables take advantage of signature aggregation. As such, we add operator's weights to the total weights. We generate a merkle root that contains individual operator stakes (`BN254OperatorInfo`) to lower transport costs. See [`BN254CertificateVerifier`](../destination/CertificateVerifier.md#bn254certificateverifier) for more information on the caching and verification scheme. 

#### `calculateOperatorTableBytes`

```solidity
/**
 * @notice Calculates the operator table bytes for a given operatorSet
 * @param operatorSet The operatorSet to calculate the operator table for
 * @return operatorTableBytes The encoded operator table bytes
 */
function calculateOperatorTableBytes(
    OperatorSet calldata operatorSet
) external view returns (bytes memory operatorTableBytes);
```

Returns the ABI-encoded bytes representation of the operator table, which is used by the `CrossChainRegistry` to calculate the operatorTable.

*Returns*:
* ABI-encoded `BN254OperatorSetInfo` struct

#### `getOperatorInfos`

```solidity
/**
 * @notice Get the operatorInfos for a given operatorSet
 * @param operatorSet the operatorSet to get the operatorInfos for
 * @return operatorInfos the operatorInfos for the given operatorSet
 */
function getOperatorInfos(
    OperatorSet calldata operatorSet
) external view returns (BN254OperatorInfo[] memory operatorInfos);
```

Returns an array of `BN254OperatorInfo` structs for all operators in the operatorSet who have registered BN254 keys.

*Effects*:
* None (view function)

### Abstract Methods

#### `_getOperatorWeights`

```solidity
/**
 * @notice Abstract function to get the operator weights for a given operatorSet
 * @param operatorSet The operatorSet to get the weights for
 * @return operators The addresses of the operators in the operatorSet
 * @return weights The weights for each operator in the operatorSet
 */
function _getOperatorWeights(
    OperatorSet calldata operatorSet
) internal view virtual returns (address[] memory operators, uint256[][] memory weights);
```

Must be implemented by derived contracts to define the weight calculation strategy. Similar to ECDSA, weights are a 2D array supporting multiple weight types per operator.

An example integration is defined by [`BN254TableCalculator`](../../../src/contracts/multichain/BN254TableCalculator.sol)
