// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// Core contracts and interfaces
import "src/contracts/interfaces/IBN254CertificateVerifier.sol";
import "src/contracts/interfaces/IBN254TableCalculator.sol";
import "src/contracts/libraries/BN254.sol";

// Forge
import "forge-std/Script.sol";
import "forge-std/Test.sol";

// Run with:
// forge script script/deploy/devnet/multichain/generate_bn254_certificate.s.sol --rpc-url $RPC_URL --sig "run()"
contract GenerateBN254Certificate is Script, Test {
    using BN254 for BN254.G1Point;
        
    function run() public {
        console.log("=== Generating BN254 Certificate ===\n");
        
        // Read wallet data from JSON
        string memory walletPath = "script/output/devnet/multichain/operatorWallet.json";
        string memory walletJson = vm.readFile(walletPath);
        
        // Parse BLS wallet data
        uint256 privateKey = vm.parseJsonUint(walletJson, ".blsWallet.privateKey");
        
        // Parse G1 public key
        uint256 pubkeyG1X = vm.parseJsonUint(walletJson, ".blsWallet.publicKeyG1.x");
        uint256 pubkeyG1Y = vm.parseJsonUint(walletJson, ".blsWallet.publicKeyG1.y");
        BN254.G1Point memory pubkeyG1 = BN254.G1Point({X: pubkeyG1X, Y: pubkeyG1Y});
        
        // Parse G2 public key
        uint256 pubkeyG2X0 = vm.parseJsonUint(walletJson, ".blsWallet.publicKeyG2.x0");
        uint256 pubkeyG2X1 = vm.parseJsonUint(walletJson, ".blsWallet.publicKeyG2.x1");
        uint256 pubkeyG2Y0 = vm.parseJsonUint(walletJson, ".blsWallet.publicKeyG2.y0");
        uint256 pubkeyG2Y1 = vm.parseJsonUint(walletJson, ".blsWallet.publicKeyG2.y1");
        BN254.G2Point memory pubkeyG2 = BN254.G2Point({
            X: [pubkeyG2X0, pubkeyG2X1],
            Y: [pubkeyG2Y0, pubkeyG2Y1]
        });
        
        // Parse operator address
        address operatorAddress = vm.parseJsonAddress(walletJson, ".wallet.address");
        
        console.log("Loaded wallet data for operator:", operatorAddress);
        console.log("Private key loaded (last 8 digits):", privateKey % 100000000);
        console.log("");
        
        // Generate signature: signature = H(message) * privateKey
        bytes32 MESSAGE_HASH = 0x50322e11fb32a88de6146a5b8ce223b8047b86d685d5622cdd72e7f43b8831e0;
        BN254.G1Point memory messagePoint = BN254.hashToG1(MESSAGE_HASH);
        BN254.G1Point memory signature = messagePoint.scalar_mul(privateKey);
        
        console.log("Message Hash:", vm.toString(MESSAGE_HASH));
        console.log("Message Point (H(message)):");
        console.log("  X:", messagePoint.X);
        console.log("  Y:", messagePoint.Y);
        console.log("");
        
        console.log("Signature (H(message) * privateKey):");
        console.log("  X:", signature.X);
        console.log("  Y:", signature.Y);
        console.log("");
        
        // Create operator info for non-signer witness
        // Since there are no signers, we include this operator as a non-signer
        IBN254TableCalculatorTypes.BN254OperatorInfo memory operatorInfo = 
            IBN254TableCalculatorTypes.BN254OperatorInfo({
                pubkey: pubkeyG1,
                weights: new uint256[](1) // Initialize with one weight of 0
            });
        operatorInfo.weights[0] = 1000; // Set some example weight
        
        // Create non-signer witness
        IBN254CertificateVerifierTypes.BN254OperatorInfoWitness memory nonSignerWitness = 
            IBN254CertificateVerifierTypes.BN254OperatorInfoWitness({
                operatorIndex: 0,
                operatorInfoProof: hex"", // Empty proof for this example
                operatorInfo: operatorInfo
            });
        
        // Create array of non-signer witnesses
        IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[] memory nonSignerWitnesses = 
            new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](1);
        nonSignerWitnesses[0] = nonSignerWitness;
        
        // Create the certificate
        uint32 referenceTimestamp = uint32(block.timestamp);
        
        IBN254CertificateVerifierTypes.BN254Certificate memory certificate = 
            IBN254CertificateVerifierTypes.BN254Certificate({
                referenceTimestamp: referenceTimestamp,
                messageHash: MESSAGE_HASH,
                signature: signature,
                apk: pubkeyG2, // APK is the same as pubkeyG2 as requested
                nonSignerWitnesses: nonSignerWitnesses
            });
        
        console.log("=== BN254 Certificate Created ===");
        console.log("Reference Timestamp:", certificate.referenceTimestamp);
        console.log("Message Hash:", vm.toString(certificate.messageHash));
        console.log("");
        
        console.log("Signature:");
        console.log("  X:", certificate.signature.X);
        console.log("  Y:", certificate.signature.Y);
        console.log("");
        
        console.log("APK (Aggregate Public Key - G2):");
        console.log("  X[0]:", certificate.apk.X[0]);
        console.log("  X[1]:", certificate.apk.X[1]);
        console.log("  Y[0]:", certificate.apk.Y[0]);
        console.log("  Y[1]:", certificate.apk.Y[1]);
        console.log("");
        
        console.log("Non-Signer Witnesses: ", certificate.nonSignerWitnesses.length);
        for (uint i = 0; i < certificate.nonSignerWitnesses.length; i++) {
            console.log("\nNon-Signer Witness", i, ":");
            console.log("  Operator Index:", certificate.nonSignerWitnesses[i].operatorIndex);
            console.log("  Operator Public Key:");
            console.log("    X:", certificate.nonSignerWitnesses[i].operatorInfo.pubkey.X);
            console.log("    Y:", certificate.nonSignerWitnesses[i].operatorInfo.pubkey.Y);
            console.log("  Weights:", certificate.nonSignerWitnesses[i].operatorInfo.weights.length);
            if (certificate.nonSignerWitnesses[i].operatorInfo.weights.length > 0) {
                console.log("    Weight[0]:", certificate.nonSignerWitnesses[i].operatorInfo.weights[0]);
            }
        }
        
        // Encode the certificate to bytes for use in transactions
        bytes memory encodedCertificate = abi.encode(certificate);
        console.log("\n=== Encoded Certificate ===");
        console.log("Length:", encodedCertificate.length, "bytes");
        console.log("Encoded (first 128 chars):");
        console.logBytes(encodedCertificate);
        
        // Save certificate to JSON
        string memory output_object = "certificate";
        
        // Serialize certificate data
        vm.serializeUint(output_object, "referenceTimestamp", certificate.referenceTimestamp);
        vm.serializeBytes32(output_object, "messageHash", certificate.messageHash);
        
        // Serialize signature
        string memory signature_obj = "signature";
        vm.serializeUint(signature_obj, "X", certificate.signature.X);
        string memory signature_output = vm.serializeUint(signature_obj, "Y", certificate.signature.Y);
        
        // Serialize APK
        string memory apk_obj = "apk";
        vm.serializeUint(apk_obj, "X0", certificate.apk.X[0]);
        vm.serializeUint(apk_obj, "X1", certificate.apk.X[1]);
        vm.serializeUint(apk_obj, "Y0", certificate.apk.Y[0]);
        string memory apk_output = vm.serializeUint(apk_obj, "Y1", certificate.apk.Y[1]);
        
        // Combine outputs
        vm.serializeString(output_object, "signature", signature_output);
        vm.serializeString(output_object, "apk", apk_output);
        vm.serializeUint(output_object, "nonSignerWitnessesCount", certificate.nonSignerWitnesses.length);
        string memory finalJson = vm.serializeBytes(output_object, "encodedCertificate", encodedCertificate);
        
        // Write to file
        vm.writeJson(
            finalJson,
            string.concat(
                "script/output/devnet/multichain/bn254_certificate_",
                vm.toString(block.timestamp),
                ".json"
            )
        );
        
        console.log("\nCertificate saved to JSON file");
    }
} 