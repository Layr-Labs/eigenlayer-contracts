// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// Core contracts
import "src/contracts/multichain/OperatorTableUpdater.sol";
import "src/contracts/interfaces/IBN254CertificateVerifier.sol";
import "src/contracts/libraries/BN254.sol";

// Forge
import "forge-std/Script.sol";
import "forge-std/Test.sol";
import "forge-std/console.sol";

// Run with:
// forge script script/deploy/devnet/multichain/call_confirmGlobalTableRoot.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run()"
contract CallConfirmGlobalTableRoot is Script, Test {
    // OperatorTableUpdater contract address
    OperatorTableUpdater public operatorTableUpdater = OperatorTableUpdater(0xd7230B89E5E2ed1FD068F0FF9198D7960243f12a);
    
    // Provided calldata
    bytes constant CALLDATA = hex"6ab40904000000000000000000000000000000000000000000000000000000000000006050322e11fb32a88de6146a5b8ce223b8047b86d685d5622cdd72e7f43b8831e00000000000000000000000000000000000000000000000000000000068479dd000000000000000000000000000000000000000000000000000000000684361bc50322e11fb32a88de6146a5b8ce223b8047b86d685d5622cdd72e7f43b8831e02d4380fbb410907800dac8fb23f6f024e840f83c067fdcc72b9694d4d4af9d660537d2c036169f7571ec1b587322b0d3dcbc80dc7656748d7f0821aeda8093920fc08688615a24d95eefdb1dca981928034e2226fb59a34f61f74a2ae136eb311dae3c560de03a89d456116d98ea3d820b8fe7083f28b10cd2497452a78084721e7491635ac0b0784a8d03d111b696f4501809a5bfa0327e8a537917a3719ff30667734cb7258c62640addb0a49df5704ebed3768af31804f40fbdbe2238d38f00000000000000000000000000000000000000000000000000000000000001200000000000000000000000000000000000000000000000000000000000000000";

    function run() public {
        // Remove function selector (first 4 bytes)
        bytes memory params = new bytes(CALLDATA.length - 4);
        for (uint i = 0; i < params.length; i++) {
            params[i] = CALLDATA[i + 4];
        }

        // Decode the parameters
        (
            IBN254CertificateVerifierTypes.BN254Certificate memory globalTableRootCert,
            bytes32 globalTableRoot,
            uint32 referenceTimestamp
        ) = abi.decode(params, (IBN254CertificateVerifierTypes.BN254Certificate, bytes32, uint32));

        // Log the decoded parameters
        console.log("=== Decoded Parameters ===");
        console.log("Global Table Root:", vm.toString(globalTableRoot));
        console.log("Reference Timestamp:", referenceTimestamp);
        
        console.log("\n=== BN254 Certificate Details ===");
        console.log("Certificate Reference Timestamp:", globalTableRootCert.referenceTimestamp);
        console.log("Message Hash:", vm.toString(globalTableRootCert.messageHash));
        
        console.log("\nAPK (Aggregate Public Key):");
        console.log("  X[0]:", globalTableRootCert.apk.X[0]);
        console.log("  X[1]:", globalTableRootCert.apk.X[1]);
        console.log("  Y[0]:", globalTableRootCert.apk.Y[0]);
        console.log("  Y[1]:", globalTableRootCert.apk.Y[1]);
        
        console.log("\nSignature:");
        console.log("  X:", globalTableRootCert.signature.X);
        console.log("  Y:", globalTableRootCert.signature.Y);
        
        console.log("\nNumber of Non-Signer Witnesses:", globalTableRootCert.nonSignerWitnesses.length);
        
        // Log non-signer witness details
        for (uint i = 0; i < globalTableRootCert.nonSignerWitnesses.length; i++) {
            console.log("\nNon-Signer Witness", i, ":");
            console.log("  Operator Index:", globalTableRootCert.nonSignerWitnesses[i].operatorIndex);
            console.log("  Public Key X:", globalTableRootCert.nonSignerWitnesses[i].operatorInfo.pubkey.X);
            console.log("  Public Key Y:", globalTableRootCert.nonSignerWitnesses[i].operatorInfo.pubkey.Y);
            console.log("  Number of Weights:", globalTableRootCert.nonSignerWitnesses[i].operatorInfo.weights.length);
            console.log("  Proof Length:", globalTableRootCert.nonSignerWitnesses[i].operatorInfoProof.length);
        }

        // Call the function
        vm.startBroadcast();
        
        console.log("\n=== Calling confirmGlobalTableRoot ===");
        console.log("Contract Address:", address(operatorTableUpdater));
        globalTableRootCert.referenceTimestamp = 1749246396;
        // globalTableRootCert.signature.X = 2812592133056428006798188102254787725159759460097192535498392252255581386194;
        // globalTableRootCert.signature.Y = 21677677408920167616725754635464008119772374805151169167895924883821892308305;
        
        operatorTableUpdater.confirmGlobalTableRoot(
            globalTableRootCert,
            globalTableRoot,
            referenceTimestamp
        );
        
        console.log("Transaction sent successfully!");
        
        vm.stopBroadcast();
    }
} 