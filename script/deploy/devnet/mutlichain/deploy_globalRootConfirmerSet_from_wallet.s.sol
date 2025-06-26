// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/Strings.sol";
import "src/test/utils/OperatorWalletLib.sol";
import "src/contracts/interfaces/IBN254TableCalculator.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";
import "src/contracts/libraries/BN254.sol";
import "src/contracts/libraries/Merkle.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// forge script script/deploy/devnet/mutlichain/deploy_globalRootConfirmerSet_from_wallet.s.sol --sig "run()"
contract DeployGlobalRootConfirmerSetFromWallet is Script, Test {
    using Strings for uint256;
    using Merkle for bytes32[];
    using BN254 for BN254.G1Point;

    struct WalletData {
        address addr;
        uint256 privateKey;
    }

    struct BlsWalletData {
        uint256 privateKey;
        BN254.G1Point publicKeyG1;
        BN254.G2Point publicKeyG2;
    }

    struct OperatorData {
        WalletData wallet;
        BlsWalletData blsWallet;
    }

    function run() public {
        /**
         *
         *                     READ WALLET DATA FROM JSON
         *
         */

        // Read the JSON file
        string memory json = vm.readFile("script/output/devnet/multichain/operatorWallet.json");

        // Parse the wallet data
        OperatorData memory operatorData;
        
        // Parse regular wallet
        operatorData.wallet.addr = vm.parseJsonAddress(json, ".wallet.address");
        operatorData.wallet.privateKey = vm.parseJsonUint(json, ".wallet.privateKey");
        
        // Parse BLS wallet
        operatorData.blsWallet.privateKey = vm.parseJsonUint(json, ".blsWallet.privateKey");
        
        // Parse publicKeyG1
        operatorData.blsWallet.publicKeyG1.X = vm.parseJsonUint(json, ".blsWallet.publicKeyG1.x");
        operatorData.blsWallet.publicKeyG1.Y = vm.parseJsonUint(json, ".blsWallet.publicKeyG1.y");
        
        // Parse publicKeyG2
        operatorData.blsWallet.publicKeyG2.X[0] = vm.parseJsonUint(json, ".blsWallet.publicKeyG2.x0");
        operatorData.blsWallet.publicKeyG2.X[1] = vm.parseJsonUint(json, ".blsWallet.publicKeyG2.x1");
        operatorData.blsWallet.publicKeyG2.Y[0] = vm.parseJsonUint(json, ".blsWallet.publicKeyG2.y0");
        operatorData.blsWallet.publicKeyG2.Y[1] = vm.parseJsonUint(json, ".blsWallet.publicKeyG2.y1");

        /**
         *
         *                     Create the `BN254OperatorInfo` struct
         *
         */

        // 1. Generate the `BN254OperatorInfo` struct
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo;

        // 2. Set the numOperators and totalWeights
        operatorSetInfo.numOperators = 1;
        uint256[] memory weights = new uint256[](1);
        weights[0] = 1;
        operatorSetInfo.totalWeights = weights;

        // 3. Set the apk (aggregate public key)
        BN254.G1Point memory aggregatePubkey;
        aggregatePubkey = aggregatePubkey.plus(operatorData.blsWallet.publicKeyG1);
        operatorSetInfo.aggregatePubkey = aggregatePubkey;

        // 4. Set the operatorInfoTreeRoot
        bytes32[] memory operatorInfoLeaves = new bytes32[](1);
        operatorInfoLeaves[0] = keccak256(
            abi.encode(
                IBN254TableCalculatorTypes.BN254OperatorInfo({
                    pubkey: operatorData.blsWallet.publicKeyG1, 
                    weights: weights
                })
            )
        );
        operatorSetInfo.operatorInfoTreeRoot = operatorInfoLeaves.merkleizeKeccak();

        /**
         *
         *                     Create the `operatorSetConfig` struct
         *
         */
        ICrossChainRegistry.OperatorSetConfig memory operatorSetConfig;
        operatorSetConfig.owner = operatorData.wallet.addr;
        operatorSetConfig.maxStalenessPeriod = 0;

        /**
         *
         *                     OUTPUT - OPERATOR SET INFO
         *
         */
        string memory parent_object = "parent object";

        // Serialize operatorSetInfo
        string memory operatorSetInfo_object = "operatorSetInfo";
        vm.serializeBytes32(operatorSetInfo_object, "operatorInfoTreeRoot", operatorSetInfo.operatorInfoTreeRoot);
        vm.serializeUint(operatorSetInfo_object, "numOperators", operatorSetInfo.numOperators);

        // Serialize apk as nested object
        string memory apk_object = "apk";
        vm.serializeUint(apk_object, "x", operatorSetInfo.aggregatePubkey.X);
        string memory apkOutput = vm.serializeUint(apk_object, "y", operatorSetInfo.aggregatePubkey.Y);
        vm.serializeString(operatorSetInfo_object, "apk", apkOutput);

        // Serialize totalWeights array
        string memory operatorSetInfoOutput =
            vm.serializeUint(operatorSetInfo_object, "totalWeights", operatorSetInfo.totalWeights);

        // Serialize operatorSetConfig
        string memory operatorSetConfig_object = "operatorSetConfig";
        vm.serializeAddress(operatorSetConfig_object, "owner", operatorSetConfig.owner);
        string memory operatorSetConfigOutput =
            vm.serializeUint(operatorSetConfig_object, "maxStalenessPeriod", operatorSetConfig.maxStalenessPeriod);

        // Combine both objects into final output
        vm.serializeString(parent_object, "operatorSetInfo", operatorSetInfoOutput);
        string memory finalJson = vm.serializeString(parent_object, "operatorSetConfig", operatorSetConfigOutput);

        // Write to the output file
        vm.writeJson(finalJson, "script/output/devnet/multichain/globalRootOperatorInfo_from_wallet.json");

        console.log("Successfully generated operator set info from wallet data");
        console.log("Output written to: script/output/devnet/multichain/globalRootOperatorInfo_from_wallet.json");
    }
} 