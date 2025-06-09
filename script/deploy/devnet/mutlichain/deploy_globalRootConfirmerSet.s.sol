// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/Strings.sol";
import "src/test/utils/OperatorWalletLib.sol";
import "src/contracts/interfaces/IBN254TableCalculator.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";

import "src/contracts/libraries/Merkle.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// forge script script/deploy/devnet/mutlichain/deploy_globalRootConfirmerSet.s.sol --sig "run(string memory)" $SEED
contract DeployGlobalRootConfirmerSet is Script, Test {
    using Strings for uint256;
    using Merkle for bytes32[];
    using BN254 for BN254.G1Point;

    function run(
        string memory salt
    ) public {
        /**
         *
         *                     WALLET CREATION
         *
         */

        // 1. Create a BN254 Wallet
        Operator memory operator = OperatorWalletLib.createOperator(salt);

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

        // 3. Set the apk
        BN254.G1Point memory aggregatePubkey;
        aggregatePubkey = aggregatePubkey.plus(operator.signingKey.publicKeyG1);
        operatorSetInfo.aggregatePubkey = aggregatePubkey;

        // 4. Set the operatorInfoTreeRoot
        bytes32[] memory operatorInfoLeaves = new bytes32[](1);
        operatorInfoLeaves[0] = keccak256(
            abi.encode(
                IBN254TableCalculatorTypes.BN254OperatorInfo({pubkey: operator.signingKey.publicKeyG1, weights: weights})
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
        operatorSetConfig.maxStalenessPeriod = 1 days;

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

        vm.writeJson(finalJson, "script/output/devnet/multichain/globalRootOperatorInfo.json");

        /**
         *
         *                     OUTPUT - BLS WALLET
         *
         */

        // Write operator data to a separate function to avoid stack too deep
        _writeOperatorData(operator);
    }

    function _writeOperatorData(
        Operator memory operator
    ) internal {
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
        vm.writeJson(operatorOutput, "script/output/devnet/multichain/operatorWallet.json");
    }
}
