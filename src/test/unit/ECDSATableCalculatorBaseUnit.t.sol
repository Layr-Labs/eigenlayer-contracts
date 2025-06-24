// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/multichain/ECDSATableCalculatorBase.sol";
import "src/contracts/permissions/KeyRegistrar.sol";
import "src/contracts/interfaces/IKeyRegistrar.sol";
import "src/contracts/interfaces/IOperatorTableCalculator.sol";
import "src/contracts/interfaces/IECDSATableCalculator.sol";
import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/libraries/SlashingLib.sol";
import "src/test/utils/EigenLayerMultichainUnitTestSetup.sol";
import "src/test/mocks/AllocationManagerMock.sol";
import "src/test/utils/Random.sol";

// Mock implementation for testing abstract contract
contract ECDSATableCalculatorBaseHarness is ECDSATableCalculatorBase {
    // Storage for mock weights
    mapping(bytes32 => address[]) internal _mockOperators;
    mapping(bytes32 => uint[][]) internal _mockWeights;

    constructor(IKeyRegistrar _keyRegistrar) ECDSATableCalculatorBase(_keyRegistrar) {}

    function setMockOperatorWeights(OperatorSet calldata operatorSet, address[] memory operators, uint[][] memory weights) external {
        bytes32 key = operatorSet.key();
        _mockOperators[key] = operators;
        _mockWeights[key] = weights;
    }

    function _getOperatorWeights(OperatorSet calldata operatorSet)
        internal
        view
        override
        returns (address[] memory operators, uint[][] memory weights)
    {
        bytes32 key = operatorSet.key();
        operators = _mockOperators[key];
        weights = _mockWeights[key];
    }
}

/**
 * @title ECDSATableCalculatorBaseUnitTests
 * @notice Base contract for all ECDSATableCalculatorBase unit tests
 */
contract ECDSATableCalculatorBaseUnitTests is EigenLayerMultichainUnitTestSetup, IECDSATableCalculatorTypes, IKeyRegistrarTypes {
    using OperatorSetLib for OperatorSet;

    // Test contracts
    ECDSATableCalculatorBaseHarness public calculator;

    // Test addresses
    address public avs1 = address(0x1);
    address public avs2 = address(0x2);
    address public operator1 = address(0x3);
    address public operator2 = address(0x4);
    address public operator3 = address(0x5);

    // Test operator sets
    OperatorSet defaultOperatorSet;
    OperatorSet alternativeOperatorSet;

    // ECDSA test keys (private keys for signature generation)
    uint constant ECDSA_PRIV_KEY_1 = 0x1234567890123456789012345678901234567890123456789012345678901234;
    uint constant ECDSA_PRIV_KEY_2 = 0x9876543210987654321098765432109876543210987654321098765432109876;
    uint constant ECDSA_PRIV_KEY_3 = 0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890;

    // ECDSA addresses (public keys)
    address public ecdsaAddress1;
    address public ecdsaAddress2;
    address public ecdsaAddress3;

    // ECDSA key data (20-byte addresses)
    bytes public ecdsaKey1;
    bytes public ecdsaKey2;
    bytes public ecdsaKey3;

    function setUp() public virtual override {
        EigenLayerMultichainUnitTestSetup.setUp();

        // Deploy calculator with KeyRegistrar
        calculator = new ECDSATableCalculatorBaseHarness(keyRegistrar);

        // Set up operator sets
        defaultOperatorSet = OperatorSet({avs: avs1, id: 0});
        alternativeOperatorSet = OperatorSet({avs: avs2, id: 1});

        // Set up ECDSA addresses that correspond to the private keys
        ecdsaAddress1 = vm.addr(ECDSA_PRIV_KEY_1);
        ecdsaAddress2 = vm.addr(ECDSA_PRIV_KEY_2);
        ecdsaAddress3 = vm.addr(ECDSA_PRIV_KEY_3);

        // Set up ECDSA key data (20-byte addresses)
        ecdsaKey1 = abi.encodePacked(ecdsaAddress1);
        ecdsaKey2 = abi.encodePacked(ecdsaAddress2);
        ecdsaKey3 = abi.encodePacked(ecdsaAddress3);

        // Configure operator sets in AllocationManager
        allocationManagerMock.setAVSRegistrar(avs1, avs1);
        allocationManagerMock.setAVSRegistrar(avs2, avs2);

        // Configure operator sets for ECDSA
        vm.prank(avs1);
        keyRegistrar.configureOperatorSet(defaultOperatorSet, IKeyRegistrarTypes.CurveType.ECDSA);

        vm.prank(avs2);
        keyRegistrar.configureOperatorSet(alternativeOperatorSet, IKeyRegistrarTypes.CurveType.ECDSA);
    }

    // Helper functions
    function _registerOperatorKey(address operator, OperatorSet memory operatorSet, address ecdsaAddress, uint privKey) internal {
        bytes memory signature = _generateECDSASignature(operator, operatorSet, ecdsaAddress, privKey);

        vm.prank(operator);
        keyRegistrar.registerKey(operator, operatorSet, abi.encodePacked(ecdsaAddress), signature);
    }

    function _generateECDSASignature(address operator, OperatorSet memory operatorSet, address ecdsaAddress, uint privKey)
        internal
        view
        returns (bytes memory)
    {
        bytes32 messageHash = keyRegistrar.getECDSAKeyRegistrationMessageHash(operator, operatorSet, ecdsaAddress);

        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privKey, messageHash);
        return abi.encodePacked(r, s, v);
    }

    function _createSingleWeightArray(uint weight) internal pure returns (uint[][] memory) {
        uint[][] memory weights = new uint[][](1);
        weights[0] = new uint[](1);
        weights[0][0] = weight;
        return weights;
    }

    function _createMultiWeightArray(uint[] memory weightValues) internal pure returns (uint[][] memory) {
        uint[][] memory weights = new uint[][](1);
        weights[0] = weightValues;
        return weights;
    }
}

/**
 * @title ECDSATableCalculatorBaseUnitTests_calculateOperatorTable
 * @notice Unit tests for ECDSATableCalculatorBase.calculateOperatorTable
 */
contract ECDSATableCalculatorBaseUnitTests_calculateOperatorTable is ECDSATableCalculatorBaseUnitTests {
    function test_noOperators() public {
        // Set empty operators and weights
        address[] memory operators = new address[](0);
        uint[][] memory weights = new uint[][](0);
        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        ECDSAOperatorInfo[] memory infos = calculator.calculateOperatorTable(defaultOperatorSet);

        assertEq(infos.length, 0, "Should have 0 operators");
    }

    function test_operatorsWithNoRegisteredKeys() public {
        // Set operators without registering their keys
        address[] memory operators = new address[](2);
        operators[0] = operator1;
        operators[1] = operator2;

        uint[][] memory weights = new uint[][](2);
        weights[0] = _createSingleWeightArray(100)[0];
        weights[1] = _createSingleWeightArray(200)[0];

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        ECDSAOperatorInfo[] memory infos = calculator.calculateOperatorTable(defaultOperatorSet);

        // When no operators have registered keys, should return empty array
        assertEq(infos.length, 0, "Should have 0 operators when none are registered");
    }

    function test_allOperatorsRegistered() public {
        // Register operators
        _registerOperatorKey(operator1, defaultOperatorSet, ecdsaAddress1, ECDSA_PRIV_KEY_1);
        _registerOperatorKey(operator2, defaultOperatorSet, ecdsaAddress2, ECDSA_PRIV_KEY_2);

        // Set operators and weights
        address[] memory operators = new address[](2);
        operators[0] = operator1;
        operators[1] = operator2;

        uint[][] memory weights = new uint[][](2);
        weights[0] = _createSingleWeightArray(100)[0];
        weights[1] = _createSingleWeightArray(200)[0];

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        ECDSAOperatorInfo[] memory infos = calculator.calculateOperatorTable(defaultOperatorSet);

        assertEq(infos.length, 2, "Should have 2 operators");
        assertEq(infos[0].pubkey, ecdsaAddress1, "Operator1 pubkey mismatch");
        assertEq(infos[0].weights[0], 100, "Operator1 weight mismatch");
        assertEq(infos[1].pubkey, ecdsaAddress2, "Operator2 pubkey mismatch");
        assertEq(infos[1].weights[0], 200, "Operator2 weight mismatch");
    }

    function test_multipleWeightTypes() public {
        // Register operators
        _registerOperatorKey(operator1, defaultOperatorSet, ecdsaAddress1, ECDSA_PRIV_KEY_1);
        _registerOperatorKey(operator2, defaultOperatorSet, ecdsaAddress2, ECDSA_PRIV_KEY_2);

        // Set operators and weights with multiple types
        address[] memory operators = new address[](2);
        operators[0] = operator1;
        operators[1] = operator2;

        uint[][] memory weights = new uint[][](2);
        uint[] memory op1Weights = new uint[](3);
        op1Weights[0] = 100;
        op1Weights[1] = 150;
        op1Weights[2] = 50;
        weights[0] = op1Weights;

        uint[] memory op2Weights = new uint[](3);
        op2Weights[0] = 200;
        op2Weights[1] = 250;
        op2Weights[2] = 100;
        weights[1] = op2Weights;

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        ECDSAOperatorInfo[] memory infos = calculator.calculateOperatorTable(defaultOperatorSet);

        assertEq(infos.length, 2, "Should have 2 operators");
        assertEq(infos[0].weights.length, 3, "Operator1 should have 3 weight types");
        assertEq(infos[0].weights[0], 100, "Operator1 weight[0] mismatch");
        assertEq(infos[0].weights[1], 150, "Operator1 weight[1] mismatch");
        assertEq(infos[0].weights[2], 50, "Operator1 weight[2] mismatch");
        assertEq(infos[1].weights.length, 3, "Operator2 should have 3 weight types");
        assertEq(infos[1].weights[0], 200, "Operator2 weight[0] mismatch");
        assertEq(infos[1].weights[1], 250, "Operator2 weight[1] mismatch");
        assertEq(infos[1].weights[2], 100, "Operator2 weight[2] mismatch");
    }

    function test_mixedRegistrationStatus() public {
        // Register only operator1
        _registerOperatorKey(operator1, defaultOperatorSet, ecdsaAddress1, ECDSA_PRIV_KEY_1);

        // Set operators and weights
        address[] memory operators = new address[](3);
        operators[0] = operator1; // registered
        operators[1] = operator2; // not registered
        operators[2] = operator3; // not registered

        uint[][] memory weights = new uint[][](3);
        weights[0] = _createSingleWeightArray(100)[0];
        weights[1] = _createSingleWeightArray(200)[0];
        weights[2] = _createSingleWeightArray(300)[0];

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        ECDSAOperatorInfo[] memory infos = calculator.calculateOperatorTable(defaultOperatorSet);

        assertEq(infos.length, 1, "Should have 1 operator (only registered ones count)");
        assertEq(infos[0].pubkey, ecdsaAddress1, "Operator1 pubkey mismatch");
        assertEq(infos[0].weights[0], 100, "Operator1 weight mismatch");
    }

    function test_singleOperatorRegistered() public {
        // Test with 1 operator, 1 registered
        address newOperator = address(uint160(100));

        address[] memory operators = new address[](1);
        operators[0] = newOperator;
        uint[][] memory weights = new uint[][](1);
        weights[0] = _createSingleWeightArray(100)[0];

        _registerOperatorKey(newOperator, defaultOperatorSet, ecdsaAddress1, ECDSA_PRIV_KEY_1);
        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        ECDSAOperatorInfo[] memory infos = calculator.calculateOperatorTable(defaultOperatorSet);
        assertEq(infos.length, 1, "Should have 1 operator");
        assertEq(infos[0].pubkey, ecdsaAddress1, "Operator pubkey mismatch");
        assertEq(infos[0].weights[0], 100, "Operator weight mismatch");
    }

    function test_subsetRegisteredToAVS() public {
        // Register all operator keys
        _registerOperatorKey(operator1, defaultOperatorSet, ecdsaAddress1, ECDSA_PRIV_KEY_1);
        _registerOperatorKey(operator2, defaultOperatorSet, ecdsaAddress2, ECDSA_PRIV_KEY_2);
        _registerOperatorKey(operator3, defaultOperatorSet, ecdsaAddress3, ECDSA_PRIV_KEY_3); // Not in actual operator set

        // Set operators and weights
        address[] memory operators = new address[](2);
        operators[0] = operator1; // registered
        operators[1] = operator2; // registered

        uint[][] memory weights = new uint[][](2);
        weights[0] = _createSingleWeightArray(100)[0];
        weights[1] = _createSingleWeightArray(200)[0];

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        ECDSAOperatorInfo[] memory infos = calculator.calculateOperatorTable(defaultOperatorSet);

        assertEq(infos.length, 2, "Should have 2 operators");
        assertEq(infos[0].pubkey, ecdsaAddress1, "Operator1 pubkey mismatch");
        assertEq(infos[0].weights[0], 100, "Operator1 weight mismatch");
        assertEq(infos[1].pubkey, ecdsaAddress2, "Operator2 pubkey mismatch");
        assertEq(infos[1].weights[0], 200, "Operator2 weight mismatch");
    }

    function test_subsetRegisteredKey() public {
        // Register operator1 and operator3, but not operator2
        _registerOperatorKey(operator1, defaultOperatorSet, ecdsaAddress1, ECDSA_PRIV_KEY_1);
        _registerOperatorKey(operator3, defaultOperatorSet, ecdsaAddress3, ECDSA_PRIV_KEY_3);

        // Set operators and weights
        address[] memory operators = new address[](3);
        operators[0] = operator1; // registered
        operators[1] = operator2; // not registered
        operators[2] = operator3; // registered

        uint[][] memory weights = new uint[][](3);
        weights[0] = _createSingleWeightArray(100)[0];
        weights[1] = _createSingleWeightArray(200)[0]; // This weight won't be included
        weights[2] = _createSingleWeightArray(300)[0];

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        ECDSAOperatorInfo[] memory infos = calculator.calculateOperatorTable(defaultOperatorSet);

        assertEq(infos.length, 2, "Should have 2 operators (only registered ones)");
        assertEq(infos[0].pubkey, ecdsaAddress1, "Operator1 pubkey mismatch");
        assertEq(infos[0].weights[0], 100, "Operator1 weight mismatch");
        assertEq(infos[1].pubkey, ecdsaAddress3, "Operator3 pubkey mismatch");
        assertEq(infos[1].weights[0], 300, "Operator3 weight mismatch");
    }

    function test_emptyOperatorSetReturnsEmptyArray() public {
        // Test with operators that exist but none registered
        address[] memory operators = new address[](3);
        operators[0] = operator1;
        operators[1] = operator2;
        operators[2] = operator3;

        uint[][] memory weights = new uint[][](3);
        weights[0] = _createSingleWeightArray(100)[0];
        weights[1] = _createSingleWeightArray(200)[0];
        weights[2] = _createSingleWeightArray(300)[0];

        // Don't register any operators
        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        ECDSAOperatorInfo[] memory infos = calculator.calculateOperatorTable(defaultOperatorSet);

        // Verify empty array when no operators are registered
        assertEq(infos.length, 0, "Should have 0 operators when none are registered");
    }

    function testFuzz_calculateOperatorTable(Randomness r, uint8 numOperators, uint8 numRegistered) public rand(r) {
        numOperators = uint8(r.Uint256() % 10 + 1); // 1-10 operators
        numRegistered = uint8(r.Uint256() % (numOperators + 1)); // 0 to numOperators registered

        address[] memory operators = new address[](numOperators);
        uint[][] memory weights = new uint[][](numOperators);

        // Generate random operators and weights
        for (uint i = 0; i < numOperators; i++) {
            operators[i] = address(uint160(r.Uint256()));
            weights[i] = _createSingleWeightArray(r.Uint256() % 1000 + 1)[0];
        }

        // Register random subset of operators
        uint[] memory registeredIndices = new uint[](numRegistered);
        for (uint i = 0; i < numRegistered; i++) {
            uint idx = r.Uint256() % numOperators;
            // Ensure unique indices
            bool unique = true;
            for (uint j = 0; j < i; j++) {
                if (registeredIndices[j] == idx) {
                    unique = false;
                    break;
                }
            }
            if (unique) {
                registeredIndices[i] = idx;
                address ecdsaAddr = vm.addr(uint(keccak256(abi.encode(operators[idx], i))));
                _registerOperatorKey(operators[idx], defaultOperatorSet, ecdsaAddr, uint(keccak256(abi.encode(operators[idx], i))));
            }
        }

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);
        ECDSAOperatorInfo[] memory infos = calculator.calculateOperatorTable(defaultOperatorSet);

        // Count actual registered operators
        uint actualRegistered = 0;
        for (uint i = 0; i < numOperators; i++) {
            if (keyRegistrar.isRegistered(defaultOperatorSet, operators[i])) actualRegistered++;
        }

        assertEq(infos.length, actualRegistered, "Should have correct number of registered operators");
    }
}

/**
 * @title ECDSATableCalculatorBaseUnitTests_calculateOperatorTableBytes
 * @notice Unit tests for ECDSATableCalculatorBase.calculateOperatorTableBytes
 */
contract ECDSATableCalculatorBaseUnitTests_calculateOperatorTableBytes is ECDSATableCalculatorBaseUnitTests {
    function test_encodesCorrectly() public {
        // Register operator
        _registerOperatorKey(operator1, defaultOperatorSet, ecdsaAddress1, ECDSA_PRIV_KEY_1);

        // Set operators and weights
        address[] memory operators = new address[](1);
        operators[0] = operator1;
        uint[][] memory weights = _createSingleWeightArray(100);

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        bytes memory tableBytes = calculator.calculateOperatorTableBytes(defaultOperatorSet);

        // Decode and verify
        ECDSAOperatorInfo[] memory decodedInfos = abi.decode(tableBytes, (ECDSAOperatorInfo[]));

        assertEq(decodedInfos.length, 1, "Should have 1 operator");
        assertEq(decodedInfos[0].weights[0], 100, "Total weight should be 100");
        assertEq(decodedInfos[0].pubkey, ecdsaAddress1, "Pubkey mismatch");
    }

    function test_multipleOperatorsEncodedCorrectly() public {
        // Register multiple operators
        _registerOperatorKey(operator1, defaultOperatorSet, ecdsaAddress1, ECDSA_PRIV_KEY_1);
        _registerOperatorKey(operator2, defaultOperatorSet, ecdsaAddress2, ECDSA_PRIV_KEY_2);
        _registerOperatorKey(operator3, defaultOperatorSet, ecdsaAddress3, ECDSA_PRIV_KEY_3);

        // Set operators and weights
        address[] memory operators = new address[](3);
        operators[0] = operator1;
        operators[1] = operator2;
        operators[2] = operator3;

        uint[][] memory weights = new uint[][](3);
        weights[0] = _createSingleWeightArray(100)[0];
        weights[1] = _createSingleWeightArray(200)[0];
        weights[2] = _createSingleWeightArray(300)[0];

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        bytes memory tableBytes = calculator.calculateOperatorTableBytes(defaultOperatorSet);

        // Decode and verify
        ECDSAOperatorInfo[] memory decodedInfos = abi.decode(tableBytes, (ECDSAOperatorInfo[]));

        assertEq(decodedInfos.length, 3, "Should have 3 operators");
        assertEq(decodedInfos[0].pubkey, ecdsaAddress1, "Operator1 pubkey mismatch");
        assertEq(decodedInfos[0].weights[0], 100, "Operator1 weight mismatch");
        assertEq(decodedInfos[1].pubkey, ecdsaAddress2, "Operator2 pubkey mismatch");
        assertEq(decodedInfos[1].weights[0], 200, "Operator2 weight mismatch");
        assertEq(decodedInfos[2].pubkey, ecdsaAddress3, "Operator3 pubkey mismatch");
        assertEq(decodedInfos[2].weights[0], 300, "Operator3 weight mismatch");
    }

    function test_emptyOperatorSetEncodesEmptyArray() public {
        // Don't register any operators
        address[] memory operators = new address[](0);
        uint[][] memory weights = new uint[][](0);

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        bytes memory tableBytes = calculator.calculateOperatorTableBytes(defaultOperatorSet);

        // Decode and verify
        ECDSAOperatorInfo[] memory decodedInfos = abi.decode(tableBytes, (ECDSAOperatorInfo[]));

        assertEq(decodedInfos.length, 0, "Should encode empty array");
    }

    function testFuzz_encodesCorrectly(Randomness r, uint weight) public rand(r) {
        weight = r.Uint256() % 1e18 + 1; // 1 to 1e18

        // Register operator
        _registerOperatorKey(operator1, defaultOperatorSet, ecdsaAddress1, ECDSA_PRIV_KEY_1);

        // Set operators and weights
        address[] memory operators = new address[](1);
        operators[0] = operator1;
        uint[][] memory weights = _createSingleWeightArray(weight);

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        bytes memory tableBytes = calculator.calculateOperatorTableBytes(defaultOperatorSet);

        // Decode and verify
        ECDSAOperatorInfo[] memory decodedInfos = abi.decode(tableBytes, (ECDSAOperatorInfo[]));

        assertEq(decodedInfos[0].weights[0], weight, "Weight mismatch");
    }

    function testFuzz_multipleWeightTypesEncoded(Randomness r, uint8 numWeightTypes) public rand(r) {
        numWeightTypes = uint8(r.Uint256() % 5 + 1); // 1-5 weight types

        // Register operator
        _registerOperatorKey(operator1, defaultOperatorSet, ecdsaAddress1, ECDSA_PRIV_KEY_1);

        // Create random weights
        uint[] memory weightValues = new uint[](numWeightTypes);
        for (uint i = 0; i < numWeightTypes; i++) {
            weightValues[i] = r.Uint256() % 1000 + 1;
        }

        address[] memory operators = new address[](1);
        operators[0] = operator1;
        uint[][] memory weights = new uint[][](1);
        weights[0] = weightValues;

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        bytes memory tableBytes = calculator.calculateOperatorTableBytes(defaultOperatorSet);

        // Decode and verify
        ECDSAOperatorInfo[] memory decodedInfos = abi.decode(tableBytes, (ECDSAOperatorInfo[]));

        assertEq(decodedInfos.length, 1, "Should have 1 operator");
        assertEq(decodedInfos[0].weights.length, numWeightTypes, "Weight types mismatch");

        for (uint i = 0; i < numWeightTypes; i++) {
            assertEq(decodedInfos[0].weights[i], weightValues[i], "Weight value mismatch");
        }
    }
}

/**
 * @title ECDSATableCalculatorBaseUnitTests_getOperatorWeights
 * @notice Unit tests for ECDSATableCalculatorBase.getOperatorWeights
 */
contract ECDSATableCalculatorBaseUnitTests_getOperatorWeights is ECDSATableCalculatorBaseUnitTests {
    function test_returnsImplementationResult() public {
        // Set mock weights
        address[] memory expectedOperators = new address[](2);
        expectedOperators[0] = operator1;
        expectedOperators[1] = operator2;

        uint[][] memory expectedWeights = new uint[][](2);
        expectedWeights[0] = _createSingleWeightArray(100)[0];
        expectedWeights[1] = _createSingleWeightArray(200)[0];

        calculator.setMockOperatorWeights(defaultOperatorSet, expectedOperators, expectedWeights);

        (address[] memory operators, uint[][] memory weights) = calculator.getOperatorWeights(defaultOperatorSet);

        assertEq(operators.length, expectedOperators.length, "Operators length mismatch");
        assertEq(weights.length, expectedWeights.length, "Weights length mismatch");

        for (uint i = 0; i < operators.length; i++) {
            assertEq(operators[i], expectedOperators[i], "Operator address mismatch");
            assertEq(weights[i][0], expectedWeights[i][0], "Weight value mismatch");
        }
    }

    function test_emptyOperatorSet() public {
        // Set empty operators and weights
        address[] memory expectedOperators = new address[](0);
        uint[][] memory expectedWeights = new uint[][](0);

        calculator.setMockOperatorWeights(defaultOperatorSet, expectedOperators, expectedWeights);

        (address[] memory operators, uint[][] memory weights) = calculator.getOperatorWeights(defaultOperatorSet);

        assertEq(operators.length, 0, "Should return empty operators array");
        assertEq(weights.length, 0, "Should return empty weights array");
    }

    function test_alternativeOperatorSet() public {
        // Set weights for alternative operator set
        address[] memory expectedOperators = new address[](1);
        expectedOperators[0] = operator3;

        uint[][] memory expectedWeights = new uint[][](1);
        expectedWeights[0] = _createSingleWeightArray(500)[0];

        calculator.setMockOperatorWeights(alternativeOperatorSet, expectedOperators, expectedWeights);

        (address[] memory operators, uint[][] memory weights) = calculator.getOperatorWeights(alternativeOperatorSet);

        assertEq(operators.length, 1, "Operators length mismatch");
        assertEq(operators[0], operator3, "Operator address mismatch");
        assertEq(weights[0][0], 500, "Weight value mismatch");
    }

    function testFuzz_returnsImplementationResult(Randomness r, uint8 numOperators) public rand(r) {
        numOperators = uint8(r.Uint256() % 20); // 0-19 operators

        address[] memory expectedOperators = new address[](numOperators);
        uint[][] memory expectedWeights = new uint[][](numOperators);

        for (uint i = 0; i < numOperators; i++) {
            expectedOperators[i] = address(uint160(r.Uint256()));
            expectedWeights[i] = _createSingleWeightArray(r.Uint256() % 1000 + 1)[0];
        }

        calculator.setMockOperatorWeights(defaultOperatorSet, expectedOperators, expectedWeights);

        (address[] memory operators, uint[][] memory weights) = calculator.getOperatorWeights(defaultOperatorSet);

        assertEq(operators.length, numOperators, "Operators length mismatch");
        assertEq(weights.length, numOperators, "Weights length mismatch");

        for (uint i = 0; i < numOperators; i++) {
            assertEq(operators[i], expectedOperators[i], "Operator address mismatch");
            assertEq(weights[i][0], expectedWeights[i][0], "Weight value mismatch");
        }
    }
}

/**
 * @title ECDSATableCalculatorBaseUnitTests_getOperatorWeight
 * @notice Unit tests for ECDSATableCalculatorBase.getOperatorWeight
 */
contract ECDSATableCalculatorBaseUnitTests_getOperatorWeight is ECDSATableCalculatorBaseUnitTests {
    function test_operatorExists() public {
        // Set operators and weights
        address[] memory operators = new address[](3);
        operators[0] = operator1;
        operators[1] = operator2;
        operators[2] = operator3;

        uint[][] memory weights = new uint[][](3);
        weights[0] = _createSingleWeightArray(100)[0];
        weights[1] = _createSingleWeightArray(200)[0];
        weights[2] = _createSingleWeightArray(300)[0];

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        assertEq(calculator.getOperatorWeight(defaultOperatorSet, operator1), 100, "Operator1 weight mismatch");
        assertEq(calculator.getOperatorWeight(defaultOperatorSet, operator2), 200, "Operator2 weight mismatch");
        assertEq(calculator.getOperatorWeight(defaultOperatorSet, operator3), 300, "Operator3 weight mismatch");
    }

    function test_operatorDoesNotExist() public {
        // Set operators and weights
        address[] memory operators = new address[](2);
        operators[0] = operator1;
        operators[1] = operator2;

        uint[][] memory weights = new uint[][](2);
        weights[0] = _createSingleWeightArray(100)[0];
        weights[1] = _createSingleWeightArray(200)[0];

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        assertEq(calculator.getOperatorWeight(defaultOperatorSet, operator3), 0, "Non-existent operator should return 0");
        assertEq(calculator.getOperatorWeight(defaultOperatorSet, address(0xdead)), 0, "Random address should return 0");
    }

    function test_emptyOperatorSet() public {
        // Set empty operators and weights
        address[] memory operators = new address[](0);
        uint[][] memory weights = new uint[][](0);

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        assertEq(calculator.getOperatorWeight(defaultOperatorSet, operator1), 0, "Should return 0 for empty set");
    }

    function test_zeroWeight() public {
        // Set operator with zero weight
        address[] memory operators = new address[](1);
        operators[0] = operator1;

        uint[][] memory weights = new uint[][](1);
        weights[0] = _createSingleWeightArray(0)[0];

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        assertEq(calculator.getOperatorWeight(defaultOperatorSet, operator1), 0, "Should return 0 for zero weight");
    }

    function test_multipleWeightTypes() public {
        // Set operator with multiple weight types
        address[] memory operators = new address[](1);
        operators[0] = operator1;

        uint[][] memory weights = new uint[][](1);
        uint[] memory multiWeights = new uint[](3);
        multiWeights[0] = 100;
        multiWeights[1] = 200;
        multiWeights[2] = 300;
        weights[0] = multiWeights;

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        // getOperatorWeight returns first weight type
        assertEq(calculator.getOperatorWeight(defaultOperatorSet, operator1), 100, "Should return first weight type");
    }

    function testFuzz_getOperatorWeight(Randomness r, address operator, uint weight) public rand(r) {
        weight = r.Uint256() % 1e18; // 0 to 1e18

        // Set single operator
        address[] memory operators = new address[](1);
        operators[0] = operator;

        uint[][] memory weights = _createSingleWeightArray(weight);

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        assertEq(calculator.getOperatorWeight(defaultOperatorSet, operator), weight, "Weight mismatch");

        // Different operator should return 0
        address differentOperator = address(uint160(uint(uint160(operator)) + 1));
        assertEq(calculator.getOperatorWeight(defaultOperatorSet, differentOperator), 0, "Different operator should return 0");
    }

    function testFuzz_multipleOperators(Randomness r, uint8 numOperators) public rand(r) {
        numOperators = uint8(r.Uint256() % 10 + 1); // 1-10 operators

        address[] memory operators = new address[](numOperators);
        uint[][] memory weights = new uint[][](numOperators);
        uint[] memory expectedWeights = new uint[](numOperators);

        for (uint i = 0; i < numOperators; i++) {
            operators[i] = address(uint160(r.Uint256()));
            expectedWeights[i] = r.Uint256() % 1000 + 1;
            weights[i] = _createSingleWeightArray(expectedWeights[i])[0];
        }

        calculator.setMockOperatorWeights(defaultOperatorSet, operators, weights);

        // Verify each operator's weight
        for (uint i = 0; i < numOperators; i++) {
            assertEq(calculator.getOperatorWeight(defaultOperatorSet, operators[i]), expectedWeights[i], "Weight mismatch");
        }

        // Non-existent operator should return 0
        address nonExistent = address(uint160(r.Uint256()));
        bool exists = false;
        for (uint i = 0; i < numOperators; i++) {
            if (operators[i] == nonExistent) {
                exists = true;
                break;
            }
        }
        if (!exists) assertEq(calculator.getOperatorWeight(defaultOperatorSet, nonExistent), 0, "Non-existent operator should return 0");
    }
}
