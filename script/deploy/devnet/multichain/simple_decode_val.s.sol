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
// forge script script/deploy/devnet/multichain/simple_decode_val.s.sol --sig "run()"
contract SimpleDecodeVal is Script, Test {
    
    // The operator table bytes to decode
    bytes constant VAL = hex"0000000000000000000000008d8a8d3f88f6a6da2083d865062bfbe3f1cfc293000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000008d8a8d3f88f6a6da2083d865062bfbe3f1cfc293000000000000000000000000000000000000000000000000000000000001518000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000020e8a1d3bfe8bfcb220cc753cf480ea2dc5c36edeacd39cbb8b9185d0765332bf900000000000000000000000000000000000000000000000000000000000000020d5501288c7ceb688f3d1783a5c7ce5dafcd1bcea53bd6167b93f86752094cc409540f25f6f74ce94276987ee3bf0c84020bbc08a2bd2745dd3cb27eecb0d8d700000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000de0b6b3a7640000";

    function run() public view {
        console.log("=== Decoding VAL ===\n");
        console.log("VAL length:", VAL.length, "bytes");
        
        // Decode all at once
        (
            OperatorSet memory operatorSet,
            IKeyRegistrarTypes.CurveType curveType,
            ICrossChainRegistryTypes.OperatorSetConfig memory config,
            IBN254TableCalculatorTypes.BN254OperatorSetInfo memory info
        ) = abi.decode(VAL, (
            OperatorSet, 
            IKeyRegistrarTypes.CurveType, 
            ICrossChainRegistryTypes.OperatorSetConfig, 
            IBN254TableCalculatorTypes.BN254OperatorSetInfo
        ));
        
        console.log("=== Decoded Results ===\n");
        
        console.log("OperatorSet:");
        console.log("  AVS:", operatorSet.avs);
        console.log("  ID:", operatorSet.id);
        
        console.log("\nCurve Type:", uint256(curveType));
        
        console.log("\nOperatorSet Config:");
        console.log("  Owner:", config.owner);
        console.log("  Max Staleness Period:", config.maxStalenessPeriod);
        
        console.log("\nBN254 Operator Set Info:");
        console.log("  Operator Info Tree Root:", vm.toString(info.operatorInfoTreeRoot));
        console.log("  Number of Operators:", info.numOperators);
        console.log("  Aggregate Public Key:");
        console.log("    X:", info.aggregatePubkey.X);
        console.log("    Y:", info.aggregatePubkey.Y);
        console.log("  Total Weights length:", info.totalWeights.length);
        
        if (info.totalWeights.length > 0) {
            console.log("  Weights:");
            for (uint i = 0; i < info.totalWeights.length && i < 5; i++) {
                console.log("    [", i, "]:", info.totalWeights[i]);
            }
        }
    }
} 