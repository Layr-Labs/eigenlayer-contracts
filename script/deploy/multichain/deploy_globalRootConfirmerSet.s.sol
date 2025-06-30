// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/Strings.sol";
import "src/test/utils/OperatorWalletLib.sol";
import "src/test/utils/Random.sol";
import "src/contracts/interfaces/IOperatorTableCalculator.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";

import "src/contracts/libraries/Merkle.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// forge script script/deploy/multichain/deploy_globalRootConfirmerSet.s.sol --sig "run(string memory)" $NETWORK
contract DeployGlobalRootConfirmerSet is Script, Test {
    using Strings for *;
    using Merkle for bytes32[];
    using BN254 for BN254.G1Point;

    address internal constant AVS = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;

    function run(string memory network, string memory salt) public {
        /**
         *
         *                     WALLET CREATION
         *
         */
        require(_strEq(network, "preprod") || _strEq(network, "testnet"), "Invalid network");

        // 1. Create a BN254 Wallet using random salt
        Operator memory operator = OperatorWalletLib.createOperator(salt);

        /**
         *
         *                     Create the `BN254OperatorInfo` struct
         *
         */

        // 1. Generate the `BN254OperatorInfo` struct
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo;

        // 2. Set the numOperators and totalWeights
        operatorSetInfo.numOperators = 1;
        uint256[] memory weights = new uint256[](1);
        weights[0] = 1;
        operatorSetInfo.totalWeights = weights;

        // 3. Set the apk
        BN254.G1Point memory aggregatePubkey;
        aggregatePubkey = aggregatePubkey.plus(operator.signingKey.publicKeyG1);
        operatorSetInfo.aggregatePubkey = aggregatePubkey;

        // 4. Set the operatorInfoTreeRoot
        bytes32[] memory operatorInfoLeaves = new bytes32[](1);
        operatorInfoLeaves[0] = keccak256(
            abi.encode(
                IOperatorTableCalculatorTypes.BN254OperatorInfo({
                    pubkey: operator.signingKey.publicKeyG1,
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
        operatorSetConfig.owner = operator.key.addr;
        operatorSetConfig.maxStalenessPeriod = 0;

        /**
         *
         *                     OUTPUT - OPERATOR SET INFO (TOML FORMAT)
         *
         */

        // Write operator set info to TOML file
        _writeOperatorSetToml(network, operatorSetInfo, operatorSetConfig);

        /**
         *
         *                     OUTPUT - BLS WALLET
         *
         */

        // Write operator data to a separate function to avoid stack too deep
        _writeOperatorData(operator, network);
    }

    function _writeOperatorData(Operator memory operator, string memory network) internal {
        string memory operator_object = "operator";

        // Serialize regular wallet info
        string memory wallet_object = "wallet";
        vm.serializeUint(wallet_object, "privateKey", operator.key.privateKey);
        string memory walletOutput = vm.serializeAddress(wallet_object, "address", operator.key.addr);

        // Serialize BLS wallet info
        string memory blsWallet_object = "blsWallet";
        vm.serializeUint(blsWallet_object, "privateKey", operator.signingKey.privateKey);

        // Serialize publicKeyG1
        string memory publicKeyG1_object = "publicKeyG1";
        vm.serializeUint(publicKeyG1_object, "x", operator.signingKey.publicKeyG1.X);
        string memory publicKeyG1Output = vm.serializeUint(publicKeyG1_object, "y", operator.signingKey.publicKeyG1.Y);
        vm.serializeString(blsWallet_object, "publicKeyG1", publicKeyG1Output);

        // Serialize publicKeyG2
        string memory publicKeyG2_object = "publicKeyG2";
        vm.serializeUint(publicKeyG2_object, "x0", operator.signingKey.publicKeyG2.X[0]);
        vm.serializeUint(publicKeyG2_object, "x1", operator.signingKey.publicKeyG2.X[1]);
        vm.serializeUint(publicKeyG2_object, "y0", operator.signingKey.publicKeyG2.Y[0]);
        string memory publicKeyG2Output =
            vm.serializeUint(publicKeyG2_object, "y1", operator.signingKey.publicKeyG2.Y[1]);
        string memory blsWalletOutput = vm.serializeString(blsWallet_object, "publicKeyG2", publicKeyG2Output);

        // Combine wallet and blsWallet into operator object
        vm.serializeString(operator_object, "wallet", walletOutput);
        string memory operatorOutput = vm.serializeString(operator_object, "blsWallet", blsWalletOutput);

        // Write to separate file
        string memory walletOutputPath = string.concat("script/deploy/multichain/", network, ".wallet.json");
        vm.writeJson(operatorOutput, walletOutputPath);
    }

    function _writeOperatorSetToml(
        string memory network,
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo,
        ICrossChainRegistry.OperatorSetConfig memory operatorSetConfig
    ) internal {
        // Build JSON object using serializeJson
        string memory json_obj = "toml_output";

        // Top level fields
        vm.serializeUint(json_obj, "globalRootConfirmationThreshold", 10_000);
        vm.serializeUint(json_obj, "referenceTimestamp", block.timestamp);

        // globalRootConfirmerSet object
        string memory confirmerSet_obj = "globalRootConfirmerSet";
        vm.serializeString(confirmerSet_obj, "avs", AVS.toHexString());
        string memory confirmerSetOutput = vm.serializeUint(confirmerSet_obj, "id", 0);
        vm.serializeString(json_obj, "globalRootConfirmerSet", confirmerSetOutput);

        // globalRootConfirmerSetConfig object
        string memory confirmerSetConfig_obj = "globalRootConfirmerSetConfig";
        vm.serializeUint(confirmerSetConfig_obj, "maxStalenessPeriod", operatorSetConfig.maxStalenessPeriod);
        string memory confirmerSetConfigOutput =
            vm.serializeAddress(confirmerSetConfig_obj, "owner", operatorSetConfig.owner);
        vm.serializeString(json_obj, "globalRootConfirmerSetConfig", confirmerSetConfigOutput);

        // globalRootConfirmerSetInfo object
        string memory confirmerSetInfo_obj = "globalRootConfirmerSetInfo";
        vm.serializeUint(confirmerSetInfo_obj, "numOperators", operatorSetInfo.numOperators);
        vm.serializeBytes32(confirmerSetInfo_obj, "operatorInfoTreeRoot", operatorSetInfo.operatorInfoTreeRoot);
        vm.serializeUint(confirmerSetInfo_obj, "totalWeights", operatorSetInfo.totalWeights);

        // aggregatePubkey nested object
        string memory aggregatePubkey_obj = "aggregatePubkey";
        vm.serializeString(aggregatePubkey_obj, "X", operatorSetInfo.aggregatePubkey.X.toString());
        string memory aggregatePubkeyOutput =
            vm.serializeString(aggregatePubkey_obj, "Y", operatorSetInfo.aggregatePubkey.Y.toString());

        string memory confirmerSetInfoOutput =
            vm.serializeString(confirmerSetInfo_obj, "aggregatePubkey", aggregatePubkeyOutput);
        string memory finalJson = vm.serializeString(json_obj, "globalRootConfirmerSetInfo", confirmerSetInfoOutput);

        // Write TOML file using writeToml
        string memory outputPath = string.concat("script/releases/v1.7.0-multichain/configs/", network, ".toml");
        vm.writeToml(finalJson, outputPath);
    }

    function _strEq(string memory a, string memory b) internal pure returns (bool) {
        return keccak256(abi.encodePacked(a)) == keccak256(abi.encodePacked(b));
    }
}
