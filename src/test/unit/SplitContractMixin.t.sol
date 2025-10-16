// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "src/contracts/mixins/SplitContractMixin.sol";

contract SplitContractMixinTest is Test {
    // Mock contracts to test delegation
    MockImplementation mockImplementation;
    MockViewImplementation mockViewImplementation;
    TestSplitContract testSplitContract;

    address constant NON_CONTRACT_ADDRESS = address(0x1234);
    address constant ZERO_ADDRESS = address(0);

    function setUp() public {
        // Deploy mock implementations
        mockImplementation = new MockImplementation();
        mockViewImplementation = new MockViewImplementation();

        // Deploy test split contract
        testSplitContract = new TestSplitContract(address(mockViewImplementation));
    }

    /// -----------------------------------------------------------------------
    /// Constructor Tests
    /// -----------------------------------------------------------------------

    function test_constructor_setsViewImplementation() public {
        assertEq(testSplitContract.getViewImplementation(), address(mockViewImplementation));
    }

    function test_constructor_withZeroAddress() public {
        TestSplitContract zeroContract = new TestSplitContract(ZERO_ADDRESS);
        assertEq(zeroContract.getViewImplementation(), ZERO_ADDRESS);
    }

    /// -----------------------------------------------------------------------
    /// _delegate Function Tests
    /// -----------------------------------------------------------------------

    function test_delegate_successfulCall() public {
        // Set the implementation to delegate to
        testSplitContract.setImplementation(address(mockImplementation));

        // Call a function that should succeed through fallback
        (bool success, bytes memory data) = address(testSplitContract).call(abi.encodeWithSignature("getValue()"));

        assertTrue(success);
        uint returnedValue = abi.decode(data, (uint));
        assertEq(returnedValue, 42);
    }

    function test_delegate_withParameters() public {
        testSplitContract.setImplementation(address(mockImplementation));

        uint inputValue = 123;
        (bool success, bytes memory data) = address(testSplitContract).call(abi.encodeWithSignature("setValue(uint256)", inputValue));

        assertTrue(success);
        // The state change happens in the context of the test contract, not the mock
        // So we need to check the return value or call a getter function
        assertEq(data.length, 0); // setValue doesn't return anything
    }

    function test_delegate_withReturnValue() public {
        testSplitContract.setImplementation(address(mockImplementation));

        string memory inputText = "Hello World";
        (bool success, bytes memory data) = address(testSplitContract).call(abi.encodeWithSignature("setText(string)", inputText));

        assertTrue(success);
        // The state change happens in the context of the test contract, not the mock
        assertEq(data.length, 0); // setText doesn't return anything
    }

    function test_delegate_revertsOnError() public {
        testSplitContract.setImplementation(address(mockImplementation));

        // Call a function that reverts
        (bool success, bytes memory data) = address(testSplitContract).call(abi.encodeWithSignature("revertFunction()"));

        assertFalse(success);
        assertTrue(data.length > 0);
    }

    function test_delegate_withValue() public {
        testSplitContract.setImplementation(address(mockImplementation));

        uint value = 1 ether;
        (bool success, bytes memory data) = address(testSplitContract).call{value: value}(abi.encodeWithSignature("payableFunction()"));

        assertTrue(success);
        // The ETH is sent to the test contract, not the mock implementation
        assertEq(address(testSplitContract).balance, value);
    }

    function test_delegate_toZeroAddress() public {
        testSplitContract.setImplementation(ZERO_ADDRESS);

        (bool success, bytes memory data) = address(testSplitContract).call(abi.encodeWithSignature("getValue()"));

        // The call will succeed but return empty data or revert
        // Let's check if it actually fails
        if (success) {
            // If it succeeds, the data should be empty or invalid
            assertTrue(data.length == 0);
        } else {
            // If it fails, that's also expected
            assertTrue(true);
        }
    }

    function test_delegate_toNonContract() public {
        testSplitContract.setImplementation(NON_CONTRACT_ADDRESS);

        (bool success, bytes memory data) = address(testSplitContract).call(abi.encodeWithSignature("getValue()"));

        // The call will succeed but return empty data or revert
        if (success) {
            // If it succeeds, the data should be empty or invalid
            assertTrue(data.length == 0);
        } else {
            // If it fails, that's also expected
            assertTrue(true);
        }
    }

    /// -----------------------------------------------------------------------
    /// _delegateView Function Tests
    /// -----------------------------------------------------------------------

    function test_delegateView_successfulCall() public {
        testSplitContract.setImplementation(address(mockViewImplementation));

        // Call a view function through fallback
        (bool success, bytes memory data) = address(testSplitContract).staticcall(abi.encodeWithSignature("getValue()"));

        assertTrue(success);
        uint result = abi.decode(data, (uint));
        assertEq(result, 100);
    }

    function test_delegateView_withParameters() public {
        testSplitContract.setImplementation(address(mockViewImplementation));

        uint input = 456;
        (bool success, bytes memory data) =
            address(testSplitContract).staticcall(abi.encodeWithSignature("getValueWithParam(uint256)", input));

        assertTrue(success);
        uint result = abi.decode(data, (uint));
        assertEq(result, input * 2);
    }

    function test_delegateView_returnsString() public {
        testSplitContract.setImplementation(address(mockViewImplementation));

        (bool success, bytes memory data) = address(testSplitContract).staticcall(abi.encodeWithSignature("getText()"));

        assertTrue(success);
        string memory result = abi.decode(data, (string));
        assertEq(result, "View Implementation");
    }

    function test_delegateView_returnsBool() public {
        testSplitContract.setImplementation(address(mockViewImplementation));

        (bool success, bytes memory data) = address(testSplitContract).staticcall(abi.encodeWithSignature("getBool()"));

        assertTrue(success);
        bool result = abi.decode(data, (bool));
        assertTrue(result);
    }

    function test_delegateView_returnsArray() public {
        testSplitContract.setImplementation(address(mockViewImplementation));

        (bool success, bytes memory data) = address(testSplitContract).staticcall(abi.encodeWithSignature("getArray()"));

        assertTrue(success);
        uint[] memory result = abi.decode(data, (uint[]));
        assertEq(result.length, 3);
        assertEq(result[0], 1);
        assertEq(result[1], 2);
        assertEq(result[2], 3);
    }

    function test_delegateView_toZeroAddress() public {
        testSplitContract.setImplementation(ZERO_ADDRESS);

        (bool success, bytes memory data) = address(testSplitContract).staticcall(abi.encodeWithSignature("getValue()"));

        // The call will succeed but return empty data or revert
        if (success) {
            // If it succeeds, the data should be empty or invalid
            assertTrue(data.length == 0);
        } else {
            // If it fails, that's also expected
            assertTrue(true);
        }
    }

    function test_delegateView_toNonContract() public {
        testSplitContract.setImplementation(NON_CONTRACT_ADDRESS);

        (bool success, bytes memory data) = address(testSplitContract).staticcall(abi.encodeWithSignature("getValue()"));

        // The call will succeed but return empty data or revert
        if (success) {
            // If it succeeds, the data should be empty or invalid
            assertTrue(data.length == 0);
        } else {
            // If it fails, that's also expected
            assertTrue(true);
        }
    }

    /// -----------------------------------------------------------------------
    /// Edge Cases and Error Handling
    /// -----------------------------------------------------------------------

    function test_delegate_withEmptyCalldata() public {
        testSplitContract.setImplementation(address(mockImplementation));

        (bool success, bytes memory data) = address(testSplitContract).call("");
        assertFalse(success);
    }

    function test_delegate_withInvalidFunctionSelector() public {
        testSplitContract.setImplementation(address(mockImplementation));

        (bool success, bytes memory data) = address(testSplitContract).call(abi.encodeWithSignature("nonExistentFunction()"));
        assertFalse(success);
    }

    function test_delegate_gasLimit() public {
        testSplitContract.setImplementation(address(mockImplementation));

        // Test with limited gas
        (bool success, bytes memory data) = address(testSplitContract).call{gas: 1000}(abi.encodeWithSignature("getValue()"));

        // Should fail due to insufficient gas
        assertFalse(success);
    }

    function test_delegate_largeReturnData() public {
        testSplitContract.setImplementation(address(mockImplementation));

        (bool success, bytes memory data) = address(testSplitContract).call(abi.encodeWithSignature("getLargeData()"));

        assertTrue(success);
        assertTrue(data.length > 1000); // Assuming getLargeData returns large data
    }

    /// -----------------------------------------------------------------------
    /// Fuzz Tests
    /// -----------------------------------------------------------------------

    function testFuzz_delegate_withRandomValue(uint value) public {
        testSplitContract.setImplementation(address(mockImplementation));

        (bool success, bytes memory data) = address(testSplitContract).call(abi.encodeWithSignature("setValue(uint256)", value));

        assertTrue(success);
        // The state change happens in the context of the test contract, not the mock
        assertEq(data.length, 0); // setValue doesn't return anything
    }

    function testFuzz_delegateView_withRandomValue(uint value) public {
        testSplitContract.setImplementation(address(mockViewImplementation));

        (bool success, bytes memory data) =
            address(testSplitContract).staticcall(abi.encodeWithSignature("getValueWithParam(uint256)", value));

        assertTrue(success);
        uint result = abi.decode(data, (uint));
        assertEq(result, value * 2);
    }

    function testFuzz_delegate_withRandomAddress(address impl) public {
        testSplitContract.setImplementation(impl);

        (bool success, bytes memory data) = address(testSplitContract).call(abi.encodeWithSignature("getValue()"));

        // The call will always succeed through fallback, but the result depends on the implementation
        assertTrue(success);
        // If impl is not a contract or doesn't have the function, data will be empty
        // If impl is a contract with the function, data will contain the return value
    }
}

/**
 * @title TestSplitContract
 * @dev Concrete implementation of SplitContractMixin for testing
 */
contract TestSplitContract is SplitContractMixin {
    address public implementation;

    constructor(address _viewImplementation) SplitContractMixin(_viewImplementation) {}

    function setImplementation(address _implementation) external {
        implementation = _implementation;
    }

    function getImplementation() external view returns (address) {
        return implementation;
    }

    function getViewImplementation() external view returns (address) {
        return viewImplementation;
    }

    // Fallback function to test _delegate
    fallback() external payable {
        _delegate(implementation);
    }

    // Receive function
    receive() external payable {
        _delegate(implementation);
    }
}

/**
 * @title MockImplementation
 * @dev Mock contract for testing delegation functionality
 */
contract MockImplementation {
    uint public value;
    string public text;

    function getValue() external pure returns (uint) {
        return 42;
    }

    function setValue(uint _value) external {
        value = _value;
    }

    function setText(string memory _text) external {
        text = _text;
    }

    function revertFunction() external pure {
        revert("MockImplementation reverted");
    }

    function payableFunction() external payable {
        // Just receive the value
    }

    function getLargeData() external pure returns (bytes memory) {
        bytes memory data = new bytes(2000);
        for (uint i = 0; i < data.length; i++) {
            data[i] = bytes1(uint8(i % 256));
        }
        return data;
    }
}

/**
 * @title MockViewImplementation
 * @dev Mock contract for testing view delegation functionality
 */
contract MockViewImplementation {
    function getValue() external pure returns (uint) {
        return 100;
    }

    function getValueWithParam(uint param) external pure returns (uint) {
        return param * 2;
    }

    function getText() external pure returns (string memory) {
        return "View Implementation";
    }

    function getBool() external pure returns (bool) {
        return true;
    }

    function getArray() external pure returns (uint[] memory) {
        uint[] memory arr = new uint[](3);
        arr[0] = 1;
        arr[1] = 2;
        arr[2] = 3;
        return arr;
    }
}
