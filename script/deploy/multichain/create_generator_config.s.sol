// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/Strings.sol";
import "src/contracts/interfaces/IOperatorTableCalculator.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";
import "src/contracts/libraries/BN254.sol";
import "src/contracts/libraries/Merkle.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// forge script script/deploy/multichain/create_generator_config.s.sol --sig "run(string memory,uint256,uint256)" $NETWORK $X_COORD $Y_COORD
contract CreateGeneratorConfig is Script, Test {
    using Strings for *;
    using Merkle for bytes32[];
    using BN254 for BN254.G1Point;

    address internal constant AVS = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;

    function run(string memory network, uint256 xCoord, uint256 yCoord) public {
        /**
         *
         *                     VALIDATION
         *
         */
        require(_strEq(network, "preprod") || _strEq(network, "testnet"), "Invalid network");
        
        // Create G1Point from provided coordinates
        BN254.G1Point memory publicKeyG1 = BN254.G1Point({
            X: xCoord,
            Y: yCoord
        });

        /**
         *
         *                     Create the `BN254OperatorSetInfo` struct
         *
         */

        // 1. Generate the `BN254OperatorSetInfo` struct
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo;

        // 2. Set the numOperators and totalWeights
        operatorSetInfo.numOperators = 1;
        uint256[] memory weights = new uint256[](1);
        weights[0] = 1;
        operatorSetInfo.totalWeights = weights;

        // 3. Set the apk (aggregate public key)
        operatorSetInfo.aggregatePubkey = publicKeyG1;

        // 4. Set the operatorInfoTreeRoot
        bytes32[] memory operatorInfoLeaves = new bytes32[](1);
        operatorInfoLeaves[0] = keccak256(
            abi.encode(
                IOperatorTableCalculatorTypes.BN254OperatorInfo({
                    pubkey: publicKeyG1,
                    weights: weights
                })
            )
        );
        operatorSetInfo.operatorInfoTreeRoot = operatorInfoLeaves.merkleizeKeccak();

        /**
         *
         *                     OUTPUT - OPERATOR SET INFO (TOML FORMAT)
         *
         */

        // Write operator set info to TOML file
        _writeOperatorSetToml(network, operatorSetInfo);
        
        // Log the generated config
        console.log("Generated operator set config for network:", network);
        console.log("Public key G1 - X:", xCoord);
        console.log("Public key G1 - Y:", yCoord);
        console.log("Config written to:", string.concat("script/releases/v1.7.0-multichain/configs/", network, ".toml"));
    }

    function _writeOperatorSetToml(
        string memory network,
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo
    ) internal {
        // Build JSON object using serializeJson
        string memory json_obj = "toml_output";

        // Top level fields
        vm.serializeUint(json_obj, "globalRootConfirmationThreshold", 10_000);

        // globalRootConfirmerSet object
        string memory confirmerSet_obj = "globalRootConfirmerSet";
        vm.serializeString(confirmerSet_obj, "avs", AVS.toHexString());
        string memory confirmerSetOutput = vm.serializeUint(confirmerSet_obj, "id", 0);
        vm.serializeString(json_obj, "globalRootConfirmerSet", confirmerSetOutput);

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
