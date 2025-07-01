// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// Core contracts
import "src/contracts/permissions/KeyRegistrar.sol";
import "src/contracts/interfaces/IKeyRegistrar.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

// Forge
import "forge-std/Script.sol";
import "forge-std/Test.sol";
import "forge-std/console.sol";

// Run with:
// forge script script/deploy/devnet/multichain/call_registerKey.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run()"
contract CallRegisterKey is Script, Test {
    // KeyRegistrar contract address
    KeyRegistrar public keyRegistrar = KeyRegistrar(0x1C84Bb62fE7791e173014A879C706445fa893BbE);
    
    // Provided calldata
    bytes constant CALLDATA = hex"d40cda160000000000000000000000006b58f6762689df33fe8fa3fc40fb5a3089d3a8cc000000000000000000000000ce2ac75be2e0951f1f7b288c7a6a9bfb6c331dc4000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000146b58f6762689df33fe8fa3fc40fb5a3089d3a8cc000000000000000000000000000000000000000000000000000000000000000000000000000000000000004118dbe8a4bd8b7298942f0f4a8ab75f995b62b0f678aededea622a4b1b8bfd8ad27c8d11933356105a55221eead4e5c32c8cc6f4aee93d2e8bbca0d41fdbb8b3800000000000000000000000000000000000000000000000000000000000000";

    function run() public {
        // Remove function selector (first 4 bytes)
        bytes memory params = new bytes(CALLDATA.length - 4);
        for (uint i = 0; i < params.length; i++) {
            params[i] = CALLDATA[i + 4];
        }

        // Decode the parameters
        (
            address operator,
            OperatorSet memory operatorSet,
            bytes memory keyData,
            bytes memory signature
        ) = abi.decode(params, (address, OperatorSet, bytes, bytes));

        // Log the decoded parameters
        console.log("=== Decoded Parameters ===");
        console.log("Operator:", operator);
        console.log("\nOperatorSet:");
        console.log("  AVS:", operatorSet.avs);
        console.log("  ID:", operatorSet.id);
        
        console.log("\nKey Data:");
        console.log("  Length:", keyData.length, "bytes");
        if (keyData.length == 20) {
            // ECDSA key - decode as address
            address keyAddress = address(bytes20(keyData));
            console.log("  Key Address:", keyAddress);
            console.log("  (This appears to be an ECDSA key)");
        } else {
            console.log("  Raw data:", vm.toString(keyData));
        }
        
        console.log("\nSignature:");
        console.log("  Length:", signature.length, "bytes");
        console.log("  Raw signature:", vm.toString(signature));
        
        // Check if operator set is configured
        console.log("\n=== Pre-checks ===");
        IKeyRegistrarTypes.CurveType curveType = keyRegistrar.getOperatorSetCurveType(operatorSet);
        console.log("Operator Set Curve Type:", uint256(curveType));
        if (curveType == IKeyRegistrarTypes.CurveType.NONE) {
            console.log("WARNING: Operator set is not configured!");
        } else if (curveType == IKeyRegistrarTypes.CurveType.ECDSA) {
            console.log("Operator set is configured for ECDSA keys");
        } else if (curveType == IKeyRegistrarTypes.CurveType.BN254) {
            console.log("Operator set is configured for BN254 keys");
        }
        
        // Check if key is already registered
        bool isAlreadyRegistered = keyRegistrar.isRegistered(operatorSet, operator);
        console.log("Is key already registered?", isAlreadyRegistered);
        
        // Call the function
        // vm.startBroadcast();
        
        console.log("\n=== Calling registerKey ===");
        console.log("Contract Address:", address(keyRegistrar));

        vm.prank(0x6B58f6762689DF33fe8fa3FC40Fb5a3089D3a8cc);
        
        keyRegistrar.registerKey(
            operator,
            operatorSet,
            keyData,
            signature
        );
        
        console.log("Transaction sent successfully!");
        
        // Verify registration
        bool isRegisteredAfter = keyRegistrar.isRegistered(operatorSet, operator);
        console.log("\n=== Post-registration check ===");
        console.log("Is key registered now?", isRegisteredAfter);
        
        if (curveType == IKeyRegistrarTypes.CurveType.ECDSA) {
            address registeredAddress = keyRegistrar.getECDSAAddress(operatorSet, operator);
            console.log("Registered ECDSA address:", registeredAddress);
        }
        
        // vm.stopBroadcast();
    }
} 