// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// Core contracts
import "src/contracts/multichain/OperatorTableUpdater.sol";
import "src/contracts/interfaces/IBN254CertificateVerifier.sol";
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
// forge script script/deploy/devnet/multichain/call_updateOperatorTable.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run()"
contract CallUpdateOperatorTable is Script, Test {
    // OperatorTableUpdater contract address
    OperatorTableUpdater public operatorTableUpdater = OperatorTableUpdater(0xd7230B89E5E2ed1FD068F0FF9198D7960243f12a);
    
    // Provided calldata
    bytes constant CALLDATA = hex"9ea94778000000000000000000000000000000000000000000000000000000006851b2b450322e11fb32a88de6146a5b8ce223b8047b86d685d5622cdd72e7f43b8831e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001e00000000000000000000000008d8a8d3f88f6a6da2083d865062bfbe3f1cfc293000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000008d8a8d3f88f6a6da2083d865062bfbe3f1cfc293000000000000000000000000000000000000000000000000000000000001518000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000020e8a1d3bfe8bfcb220cc753cf480ea2dc5c36edeacd39cbb8b9185d0765332bf900000000000000000000000000000000000000000000000000000000000000020d5501288c7ceb688f3d1783a5c7ce5dafcd1bcea53bd6167b93f86752094cc409540f25f6f74ce94276987ee3bf0c84020bbc08a2bd2745dd3cb27eecb0d8d700000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000de0b6b3a7640000";

    function run() public {
        // Remove function selector (first 4 bytes)
        bytes memory params = new bytes(CALLDATA.length - 4);
        for (uint i = 0; i < params.length; i++) {
            params[i] = CALLDATA[i + 4];
        }

        // Decode the parameters
        (
            uint32 referenceTimestamp,
            bytes32 globalTableRoot,
            uint32 operatorSetIndex,
            bytes memory proof,
            bytes memory operatorTableBytes
        ) = abi.decode(params, (uint32, bytes32, uint32, bytes, bytes));

        // Log the decoded parameters
        console.log("=== Decoded Parameters ===");
        console.log("Reference Timestamp:", referenceTimestamp);
        console.log("Global Table Root:", vm.toString(globalTableRoot));
        console.log("Operator Set Index:", operatorSetIndex);
        console.log("Proof length:", proof.length, "bytes");
        console.log("Operator Table Bytes length:", operatorTableBytes.length, "bytes");
        
        // Decode the operator table info
        console.log("\n=== Operator Table Details ===");

        (
            OperatorSet memory operatorSet,
            IKeyRegistrarTypes.CurveType curveType,
            ICrossChainRegistryTypes.OperatorSetConfig memory operatorSetConfig,
            bytes memory operatorTableInfoBytes
        ) = abi.decode(operatorTableBytes, (OperatorSet, IKeyRegistrarTypes.CurveType, ICrossChainRegistryTypes.OperatorSetConfig, bytes));
        
        console.log("OperatorSet:");
        console.log("  AVS:", operatorSet.avs);
        console.log("  ID:", operatorSet.id);
        console.log("Curve Type:", uint256(curveType));
        console.log("OperatorSet Config:");
        console.log("  Owner:", operatorSetConfig.owner);
        console.log("  Max Staleness Period:", operatorSetConfig.maxStalenessPeriod);
        
        // If it's BN254, decode the operator set info
        if (curveType == IKeyRegistrarTypes.CurveType.BN254) {
            // console.log("\n=== BN254 Operator Set Info ===");
            (IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo) = 
                abi.decode(operatorTableInfoBytes, (IBN254TableCalculatorTypes.BN254OperatorSetInfo));
            
            console.log("Operator Info Tree Root:", vm.toString(operatorSetInfo.operatorInfoTreeRoot));
            console.log("Number of Operators:", operatorSetInfo.numOperators);
            console.log("Aggregate Public Key:");
            console.log("  X:", operatorSetInfo.aggregatePubkey.X);
            console.log("  Y:", operatorSetInfo.aggregatePubkey.Y);
            console.log("Total Weights length:", operatorSetInfo.totalWeights.length);
            for (uint i = 0; i < operatorSetInfo.totalWeights.length && i < 5; i++) {
                console.log("  Weight[", i, "]:", operatorSetInfo.totalWeights[i]);
            }
        }
        
        // Log proof data (first few bytes)
        if (proof.length > 0) {
            console.log("\n=== Proof Data (first 64 bytes) ===");
            bytes memory proofSnippet = new bytes(proof.length < 64 ? proof.length : 64);
            for (uint i = 0; i < proofSnippet.length; i++) {
                proofSnippet[i] = proof[i];
            }
            console.logBytes(proofSnippet);
        }

        // Call the function
        vm.startBroadcast();
        
        console.log("\n=== Calling updateOperatorTable ===");
        console.log("Contract Address:", address(operatorTableUpdater));
        
        operatorTableUpdater.updateOperatorTable(
            referenceTimestamp,
            globalTableRoot,
            operatorSetIndex,
            proof,
            operatorTableBytes
        );
        
        console.log("Transaction sent successfully!");
        
        vm.stopBroadcast();
    }
} 