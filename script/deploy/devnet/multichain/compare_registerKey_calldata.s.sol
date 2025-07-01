// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// Core contracts
import "src/contracts/interfaces/IKeyRegistrar.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

// Forge
import "forge-std/Script.sol";
import "forge-std/Test.sol";
import "forge-std/console.sol";

// Run with:
// forge script script/deploy/devnet/multichain/compare_registerKey_calldata.s.sol --sig "run()"
contract CompareRegisterKeyCalldata is Script, Test {
    
    // The two calldata strings to compare
    bytes constant CALLDATA_WORKS = hex"d40cda160000000000000000000000006b58f6762689df33fe8fa3fc40fb5a3089d3a8cc000000000000000000000000ce2ac75be2e0951f1f7b288c7a6a9bfb6c331dc4000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000146b58f6762689df33fe8fa3fc40fb5a3089d3a8cc00000000000000000000000000000000000000000000000000000000000000000000000000000000000000412964ee5e454d64347cecdee91e872230b556bae02f81df4689bd1226535de3f600a35b9b380b2cda991b4a3aa3618d6f94b12eddec8862b94eb6acb6d04c5c9c1b00000000000000000000000000000000000000000000000000000000000000";
    
    bytes constant CALLDATA_FAILS = hex"d40cda160000000000000000000000006b58f6762689df33fe8fa3fc40fb5a3089d3a8cc000000000000000000000000ce2ac75be2e0951f1f7b288c7a6a9bfb6c331dc4000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000146b58f6762689df33fe8fa3fc40fb5a3089d3a8cc000000000000000000000000000000000000000000000000000000000000000000000000000000000000004118dbe8a4bd8b7298942f0f4a8ab75f995b62b0f678aededea622a4b1b8bfd8ad27c8d11933356105a55221eead4e5c32c8cc6f4aee93d2e8bbca0d41fdbb8b380000000000000000000000000000000000000000000000000000000000000000";
    
    function run() public view {
        console.log("=== Comparing RegisterKey Calldata ===\n");
        
        // Extract function selectors
        bytes4 selector1 = bytes4(CALLDATA_WORKS);
        bytes4 selector2 = bytes4(CALLDATA_FAILS);
        
        console.log("Function Selectors:");
        console.log("  Working calldata:", vm.toString(selector1));
        console.log("  Failing calldata:", vm.toString(selector2));
        console.log("  Selectors match:", selector1 == selector2);
        
        // Decode both calldatas
        console.log("\n=== Decoding Working Calldata ===");
        (
            address operator1,
            OperatorSet memory operatorSet1,
            bytes memory keyData1,
            bytes memory signature1
        ) = _decodeCalldata(CALLDATA_WORKS);
        
        _printDecodedData("Working", operator1, operatorSet1, keyData1, signature1);
        
        console.log("\n=== Decoding Failing Calldata ===");
        (
            address operator2,
            OperatorSet memory operatorSet2,
            bytes memory keyData2,
            bytes memory signature2
        ) = _decodeCalldata(CALLDATA_FAILS);
        
        _printDecodedData("Failing", operator2, operatorSet2, keyData2, signature2);
        
        // Compare the differences
        console.log("\n=== Comparison Analysis ===");
        console.log("Operators match:", operator1 == operator2);
        console.log("AVS addresses match:", operatorSet1.avs == operatorSet2.avs);
        console.log("Operator Set IDs match:", operatorSet1.id == operatorSet2.id);
        console.log("Key data matches:", keccak256(keyData1) == keccak256(keyData2));
        console.log("Signature length matches:", signature1.length == signature2.length);
        console.log("Signatures match:", keccak256(signature1) == keccak256(signature2));
        
        // Detailed signature analysis
        console.log("\n=== Signature Analysis ===");
        if (signature1.length == 65 && signature2.length == 65) {
            // Decode ECDSA signatures (r, s, v)
            bytes32 r1 = bytes32(slice(signature1, 0, 32));
            bytes32 s1 = bytes32(slice(signature1, 32, 32));
            uint8 v1 = uint8(signature1[64]);
            
            bytes32 r2 = bytes32(slice(signature2, 0, 32));
            bytes32 s2 = bytes32(slice(signature2, 32, 32));
            uint8 v2 = uint8(signature2[64]);
            
            console.log("Working signature components:");
            console.log("  r:", vm.toString(r1));
            console.log("  s:", vm.toString(s1));
            console.log("  v:", v1);
            
            console.log("\nFailing signature components:");
            console.log("  r:", vm.toString(r2));
            console.log("  s:", vm.toString(s2));
            console.log("  v:", v2);
            
            console.log("\nComponent comparison:");
            console.log("  r values match:", r1 == r2);
            console.log("  s values match:", s1 == s2);
            console.log("  v values match:", v1 == v2);
        }
        
        // Raw hex comparison of signatures
        console.log("\n=== Raw Signature Hex ===");
        console.log("Working signature:");
        console.log(vm.toString(signature1));
        console.log("\nFailing signature:");
        console.log(vm.toString(signature2));
        
        // Byte-by-byte comparison
        if (signature1.length == signature2.length) {
            uint diffCount = 0;
            for (uint i = 0; i < signature1.length; i++) {
                if (signature1[i] != signature2[i]) {
                    if (diffCount < 10) { // Only show first 10 differences
                        console.log("  Difference at byte", i, "- Working:", uint8(signature1[i]), "Failing:", uint8(signature2[i]));
                    }
                    diffCount++;
                }
            }
            console.log("\nTotal bytes different:", diffCount);
        }
    }
    
    function _decodeCalldata(bytes memory calldata_) internal pure returns (
        address operator,
        OperatorSet memory operatorSet,
        bytes memory keyData,
        bytes memory signature
    ) {
        // Skip function selector (first 4 bytes)
        bytes memory params = new bytes(calldata_.length - 4);
        for (uint i = 0; i < params.length; i++) {
            params[i] = calldata_[i + 4];
        }
        
        // Decode the parameters
        (operator, operatorSet, keyData, signature) = abi.decode(
            params,
            (address, OperatorSet, bytes, bytes)
        );
    }
    
    function _printDecodedData(
        string memory label,
        address operator,
        OperatorSet memory operatorSet,
        bytes memory keyData,
        bytes memory signature
    ) internal pure {
        console.log(label);
        console.log("  Operator:", operator);
        console.log("  AVS:", operatorSet.avs);
        console.log("  Operator Set ID:", operatorSet.id);
        console.log("  Key Data Length (bytes):", keyData.length);
        
        if (keyData.length == 20) {
            address keyAddress = address(bytes20(keyData));
            console.log("  Key Address:", keyAddress);
        }
        
        console.log("  Signature Length:", signature.length, "bytes");
    }
    
    function slice(bytes memory data, uint start, uint len) internal pure returns (bytes memory) {
        bytes memory result = new bytes(len);
        for (uint i = 0; i < len; i++) {
            result[i] = data[start + i];
        }
        return result;
    }
} 