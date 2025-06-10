// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// Core contracts
import "src/contracts/interfaces/IBN254TableCalculator.sol";
import "src/contracts/interfaces/IKeyRegistrar.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";
import "src/contracts/libraries/BN254.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

// Forge
import "forge-std/Script.sol";
import "forge-std/Test.sol";
import "forge-std/console.sol";

// Run with:
// forge script script/deploy/devnet/multichain/debug_decode_operator_table.s.sol --sig "run()"
contract DebugDecodeOperatorTable is Script, Test {
    
    // The operator table bytes to decode
    bytes constant VAL = hex"0000000000000000000000008d8a8d3f88f6a6da2083d865062bfbe3f1cfc293000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000008d8a8d3f88f6a6da2083d865062bfbe3f1cfc293000000000000000000000000000000000000000000000000000000000001518000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000020e8a1d3bfe8bfcb220cc753cf480ea2dc5c36edeacd39cbb8b9185d0765332bf900000000000000000000000000000000000000000000000000000000000000020d5501288c7ceb688f3d1783a5c7ce5dafcd1bcea53bd6167b93f86752094cc409540f25f6f74ce94276987ee3bf0c84020bbc08a2bd2745dd3cb27eecb0d8d700000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000de0b6b3a7640000";

    function run() public view {
        console.log("=== Debugging Operator Table Decoding ===\n");
        
        // Step 1: Decode just the OperatorSet
        console.log("Step 1: Decoding OperatorSet only...");
        (OperatorSet memory operatorSet) = abi.decode(VAL, (OperatorSet));
        console.log("  Success! AVS:", operatorSet.avs);
        console.log("  ID:", operatorSet.id);
        
        // Step 2: Decode OperatorSet + CurveType
        console.log("\nStep 2: Decoding OperatorSet + CurveType...");
        (OperatorSet memory operatorSet2, IKeyRegistrarTypes.CurveType curveType) = 
            abi.decode(VAL, (OperatorSet, IKeyRegistrarTypes.CurveType));
        console.log("  Success! CurveType:", uint256(curveType));
        
        // Step 3: Decode OperatorSet + CurveType + OperatorSetConfig
        console.log("\nStep 3: Decoding OperatorSet + CurveType + OperatorSetConfig...");
        (
            OperatorSet memory operatorSet3,
            IKeyRegistrarTypes.CurveType curveType2,
            ICrossChainRegistryTypes.OperatorSetConfig memory config
        ) = abi.decode(VAL, (OperatorSet, IKeyRegistrarTypes.CurveType, ICrossChainRegistryTypes.OperatorSetConfig));
        console.log("  Success! Config owner:", config.owner);
        console.log("  Max staleness period:", config.maxStalenessPeriod);
        
        // Step 4: Manually inspect the bytes
        console.log("\nStep 4: Manual byte inspection...");
        console.log("Total VAL length:", VAL.length, "bytes");
        
        // Extract first few 32-byte words
        bytes32 word0;
        bytes32 word1;
        bytes32 word2;
        bytes32 word3;
        bytes32 word4;
        
        assembly {
            word0 := mload(add(VAL, 0x20))
            word1 := mload(add(VAL, 0x40))
            word2 := mload(add(VAL, 0x60))
            word3 := mload(add(VAL, 0x80))
            word4 := mload(add(VAL, 0xA0))
        }
        
        console.log("\nFirst 5 words (32 bytes each):");
        console.logBytes32(word0);
        console.logBytes32(word1);
        console.logBytes32(word2);
        console.logBytes32(word3);
        console.logBytes32(word4);
        
        // Step 5: Try full decode
        console.log("\nStep 5: Attempting full decode...");
        (
            OperatorSet memory operatorSet4,
            IKeyRegistrarTypes.CurveType curveType3,
            ICrossChainRegistryTypes.OperatorSetConfig memory config2,
            IBN254TableCalculatorTypes.BN254OperatorSetInfo memory info
        ) = abi.decode(VAL, (
            OperatorSet, 
            IKeyRegistrarTypes.CurveType, 
            ICrossChainRegistryTypes.OperatorSetConfig, 
            IBN254TableCalculatorTypes.BN254OperatorSetInfo
        ));
        
        console.log("  Success! Full decode worked.");
        console.log("  Operator Info Tree Root:", vm.toString(info.operatorInfoTreeRoot));
        console.log("  Num Operators:", info.numOperators);
        console.log("  APK X:", info.aggregatePubkey.X);
        console.log("  APK Y:", info.aggregatePubkey.Y);
        console.log("  Total Weights length:", info.totalWeights.length);
        
        // Step 6: Check if it's an offset issue
        console.log("\nStep 6: Checking for dynamic data offsets...");
        
        // The BN254OperatorSetInfo contains a dynamic array (totalWeights)
        // Let's check if the data has proper offsets
        uint256 offset = 32 * 6; // Skip OperatorSet (2 words) + CurveType (1 word) + Config (2 words) + offset to BN254OperatorSetInfo
        if (VAL.length > offset + 32) {
            bytes32 possibleOffset;
            assembly {
                possibleOffset := mload(add(VAL, add(0x20, offset)))
            }
            console.log("  Possible offset at position", offset, ":");
            console.logBytes32(possibleOffset);
        }
    }
} 