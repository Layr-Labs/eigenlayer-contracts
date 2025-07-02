// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// Core contracts
import "src/contracts/permissions/KeyRegistrar.sol";
import "src/contracts/interfaces/IKeyRegistrar.sol";
import "src/contracts/core/AllocationManager.sol";
import "src/contracts/interfaces/IAllocationManager.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

// Forge
import "forge-std/Script.sol";
import "forge-std/Test.sol";
import "forge-std/console.sol";

// Run with:
// forge script script/deploy/devnet/multichain/generate_registerKey_calldata_v2.s.sol --rpc-url $RPC_URL --sig "run()"
contract GenerateRegisterKeyCalldataV2 is Script, Test {
    // Contract addresses
    KeyRegistrar public keyRegistrar = KeyRegistrar(0x1C84Bb62fE7791e173014A879C706445fa893BbE);
    AllocationManager public allocationManager = AllocationManager(0xFdD5749e11977D60850E06bF5B13221Ad95eb6B4);
    
    // Configuration
    address public constant OPERATOR = 0x6B58f6762689DF33fe8fa3FC40Fb5a3089D3a8cc;
    address public constant AVS = 0xCE2Ac75bE2E0951F1F7B288c7a6A9BfB6c331DC4;
    uint32 public constant OPERATOR_SET_ID = 0;
    uint256 public constant PRIVATE_KEY = 0x3dd7c381f27775d9945f0fcf5bb914484c4d01681824603c71dd762259f43214;
    
    function run() public view {
        console.log("=== Generating Calldata for Operator Registration ===\n");
        
        // Create operator set
        OperatorSet memory operatorSet = OperatorSet({
            avs: AVS,
            id: OPERATOR_SET_ID
        });
        
        // For ECDSA, the key data is the operator address packed into bytes
        bytes memory keyData = abi.encodePacked(OPERATOR);
        
        console.log("Configuration:");
        console.log("  Operator:", OPERATOR);
        console.log("  AVS:", AVS);
        console.log("  Operator Set ID:", OPERATOR_SET_ID);
        console.log("  Key Address (ECDSA):", OPERATOR);
        console.log("  Private Key: [REDACTED]");
        
        // Generate ECDSA signature for registerKey
        bytes memory signature = _generateECDSASignature(OPERATOR, operatorSet, OPERATOR, PRIVATE_KEY);
        
        console.log("\nGenerated Key Registration Data:");
        console.log("  Key Data Length:", keyData.length, "bytes");
        console.log("  Key Data:", vm.toString(keyData));
        console.log("  Signature Length:", signature.length, "bytes");
        console.log("  Signature:", vm.toString(signature));
        
        // Generate registerKey calldata
        bytes memory registerKeyCalldata = abi.encodeWithSignature(
            "registerKey(address,(address,uint32),bytes,bytes)",
            OPERATOR,
            operatorSet,
            keyData,
            signature
        );
        
        console.log("\n=== Generated registerKey Calldata ===");
        console.log("Contract Address:", address(keyRegistrar));
        console.log("Function: registerKey(address,(address,uint32),bytes,bytes)");
        console.log("Selector: 0xd40cda16");
        console.log("Calldata Length:", registerKeyCalldata.length, "bytes");
        console.log("\nCalldata (hex):");
        console.log(vm.toString(registerKeyCalldata));
        
        // Generate registerForOperatorSets calldata
        uint32[] memory operatorSetIds = new uint32[](1);
        operatorSetIds[0] = OPERATOR_SET_ID;
        
        IAllocationManagerTypes.RegisterParams memory registerParams = IAllocationManagerTypes.RegisterParams({
            avs: AVS,
            operatorSetIds: operatorSetIds,
            data: ""
        });
        
        bytes memory registerOperatorCalldata = abi.encodeWithSignature(
            "registerForOperatorSets(address,(address,uint32[],bytes))",
            OPERATOR,
            registerParams
        );
        
        console.log("\n=== Generated registerForOperatorSets Calldata ===");
        console.log("Contract Address:", address(allocationManager));
        console.log("Function: registerForOperatorSets(address,(address,uint32[],bytes))");
        console.log("Selector: 0xadc2e3d9");
        console.log("Calldata Length:", registerOperatorCalldata.length, "bytes");
        console.log("\nCalldata (hex):");
        console.log(vm.toString(registerOperatorCalldata));
        
        console.log("\n=== Order of Operations ===");
        console.log("1. First call registerForOperatorSets on AllocationManager");
        console.log("2. Then call registerKey on KeyRegistrar");
        
        // Verify both calldatas
        console.log("\n=== Verification ===");
        _verifyRegisterKeyCalldata(registerKeyCalldata);
        _verifyRegisterOperatorCalldata(registerOperatorCalldata);
    }
    
    function _generateECDSASignature(
        address _operator,
        OperatorSet memory operatorSet,
        address keyAddress,
        uint256 privKey
    ) internal view returns (bytes memory) {
        // Get the typehash from KeyRegistrar
        // bytes32 ECDSA_KEY_REGISTRATION_TYPEHASH = keyRegistrar.ECDSA_KEY_REGISTRATION_TYPEHASH();
        
        // // Create the struct hash
        // bytes32 structHash = keccak256(
        //     abi.encode(ECDSA_KEY_REGISTRATION_TYPEHASH, _operator, operatorSet.avs, operatorSet.id, keyAddress)
        // );
        
        // // Get domain separator and create full message hash
        // bytes32 domainSeparator = keyRegistrar.domainSeparator();
        // bytes32 messageHash = keccak256(abi.encodePacked("\x19\x01", domainSeparator, structHash));
        
        // console.log("\nSignature Generation Details:");
        // console.log("  ECDSA_KEY_REGISTRATION_TYPEHASH:", vm.toString(ECDSA_KEY_REGISTRATION_TYPEHASH));
        // console.log("  Struct Hash:", vm.toString(structHash));
        // console.log("  Domain Separator:", vm.toString(domainSeparator));
        // console.log("  Message Hash:", vm.toString(messageHash));

        bytes32 messageHash = keyRegistrar.getECDSAKeyRegistrationMessageHash(_operator, operatorSet, keyAddress);
        console.log("Message Hash:", vm.toString(messageHash));
        
        // Sign the message
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privKey, messageHash);
        
        console.log("  Signature components:");
        console.log("    v:", v);
        console.log("    r:", vm.toString(r));
        console.log("    s:", vm.toString(s));
        
        return abi.encodePacked(r, s, v);
    }
    
    function _verifyRegisterKeyCalldata(bytes memory encodedCall) internal pure {
        // Skip function selector (first 4 bytes)
        bytes memory params = new bytes(encodedCall.length - 4);
        for (uint i = 0; i < params.length; i++) {
            params[i] = encodedCall[i + 4];
        }
        
        // Decode the parameters
        (
            address decodedOperator,
            OperatorSet memory decodedOperatorSet,
            bytes memory decodedKeyData,
            bytes memory decodedSignature
        ) = abi.decode(params, (address, OperatorSet, bytes, bytes));
        
        console.log("RegisterKey - Decoded parameters match:");
        console.log("  Operator matches:", decodedOperator == OPERATOR);
        console.log("  AVS matches:", decodedOperatorSet.avs == AVS);
        console.log("  ID matches:", decodedOperatorSet.id == OPERATOR_SET_ID);
        console.log("  Key data length:", decodedKeyData.length);
        console.log("  Signature length:", decodedSignature.length);
    }
    
    function _verifyRegisterOperatorCalldata(bytes memory encodedCall) internal pure {
        // Skip function selector (first 4 bytes)
        bytes memory params = new bytes(encodedCall.length - 4);
        for (uint i = 0; i < params.length; i++) {
            params[i] = encodedCall[i + 4];
        }
        
        // Decode the parameters
        (
            address decodedOperator,
            IAllocationManagerTypes.RegisterParams memory decodedParams
        ) = abi.decode(params, (address, IAllocationManagerTypes.RegisterParams));
        
        console.log("RegisterForOperatorSets - Decoded parameters match:");
        console.log("  Operator matches:", decodedOperator == OPERATOR);
        console.log("  AVS matches:", decodedParams.avs == AVS);
        console.log("  Number of operator set IDs:", decodedParams.operatorSetIds.length);
        if (decodedParams.operatorSetIds.length > 0) {
            console.log("  First operator set ID matches:", decodedParams.operatorSetIds[0] == OPERATOR_SET_ID);
        }
        console.log("  Data length:", bytes(decodedParams.data).length);
    }
} 