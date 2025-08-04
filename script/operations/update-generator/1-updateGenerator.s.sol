// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../releases/Env.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";

// Types
import "src/contracts/interfaces/IOperatorTableCalculator.sol";
import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/libraries/BN254.sol";
import "src/contracts/interfaces/IBaseCertificateVerifier.sol";

// For TOML parsing
import {stdToml} from "forge-std/StdToml.sol";

/**
 * Purpose: Update the generator on a TESTNET environment
 */
contract QueueTransferProxyAdmin is MultisigBuilder {
    using Env for *;
    using OperatorSetLib for OperatorSet;
    using stdToml for string;

    string private constant TESTNET_CONFIG_PATH = "script/releases/v1.7.0-multichain/configs/testnet.toml";

    function _runAsMultisig() internal virtual override prank(Env.opsMultisig()) {
        GeneratorParams memory generatorParams = _getGeneratorParams(TESTNET_CONFIG_PATH);
        Env.proxy.operatorTableUpdater().updateGenerator(generatorParams.generator, generatorParams.generatorInfo);
    }

    function testScript() public virtual {
        // Require that the environment is a testnet environment supported by multichain
        require(
            Env._strEq(Env.env(), "testnet-sepolia") || Env._strEq(Env.env(), "testnet-base-sepolia"),
            "Environment must be a testnet environment"
        );

        // Update the generator
        execute();

        // Get the generator params
        GeneratorParams memory generatorParams = _getGeneratorParams(TESTNET_CONFIG_PATH);

        // Check that the generator has been updated in the operator table updater
        OperatorTableUpdater operatorTableUpdater = Env.proxy.operatorTableUpdater();
        assertEq(
            operatorTableUpdater.getGenerator().key(),
            generatorParams.generator.key(),
            "operatorTableUpdater.generator invalid"
        );
        assertEq(
            operatorTableUpdater.getGeneratorReferenceTimestamp(),
            operatorTableUpdater.GENERATOR_REFERENCE_TIMESTAMP(),
            "operatorTableUpdater.generatorReferenceTimestamp invalid"
        );
        assertEq(
            operatorTableUpdater.getGeneratorReferenceTimestamp(),
            1,
            "operatorTableUpdater.generatorReferenceTimestamp invalid"
        );
        assertEq(
            operatorTableUpdater.getGlobalTableRootByTimestamp(1),
            operatorTableUpdater.GENERATOR_GLOBAL_TABLE_ROOT(),
            "operatorTableUpdater.generatorGlobalTableRoot invalid"
        );

        // Check that the generator has been updated in the certificate verifier
        BN254CertificateVerifier certificateVerifier = Env.proxy.bn254CertificateVerifier();
        assertEq(
            certificateVerifier.latestReferenceTimestamp(generatorParams.generator),
            operatorTableUpdater.GENERATOR_REFERENCE_TIMESTAMP(),
            "certificateVerifier.latestReferenceTimestamp invalid"
        );
        assertEq(
            IBaseCertificateVerifier(address(certificateVerifier)).maxOperatorTableStaleness(generatorParams.generator),
            operatorTableUpdater.GENERATOR_MAX_STALENESS_PERIOD(),
            "certificateVerifier.maxStalenessPeriod invalid"
        );
        assertEq(
            certificateVerifier.getOperatorSetOwner(generatorParams.generator),
            address(operatorTableUpdater),
            "certificateVerifier.operatorSetOwner invalid"
        );
        // Get the operatorSetInfo
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = certificateVerifier
            .getOperatorSetInfo(generatorParams.generator, operatorTableUpdater.GENERATOR_REFERENCE_TIMESTAMP());
        assertEq(
            operatorSetInfo.numOperators,
            generatorParams.generatorInfo.numOperators,
            "certificateVerifier.getOperatorSetInfo.numOperators invalid"
        );
        assertEq(
            operatorSetInfo.operatorInfoTreeRoot,
            generatorParams.generatorInfo.operatorInfoTreeRoot,
            "certificateVerifier.getOperatorSetInfo.operatorInfoTreeRoot invalid"
            "certificateVerifier.getOperatorSetInfo invalid"
        );
        assertEq(
            operatorSetInfo.aggregatePubkey.X,
            generatorParams.generatorInfo.aggregatePubkey.X,
            "certificateVerifier.getOperatorSetInfo.aggregatePubkey.X invalid"
        );
        assertEq(
            operatorSetInfo.aggregatePubkey.Y,
            generatorParams.generatorInfo.aggregatePubkey.Y,
            "certificateVerifier.getOperatorSetInfo.aggregatePubkey.Y invalid"
        );
        assertEq(
            operatorSetInfo.totalWeights,
            generatorParams.generatorInfo.totalWeights,
            "certificateVerifier.getOperatorSetInfo.totalWeights invalid"
        );
    }

    function _getGeneratorParams(
        string memory path
    ) internal view returns (GeneratorParams memory generatorParams) {
        // Read the TOML file
        string memory root = vm.projectRoot();
        string memory fullPath = string.concat(root, "/", path);
        string memory toml = vm.readFile(fullPath);

        // Parse globalRootConfirmerSet
        address avs = toml.readAddress(".globalRootConfirmerSet.avs");
        uint32 id = uint32(toml.readUint(".globalRootConfirmerSet.id"));
        generatorParams.generator = OperatorSet({avs: avs, id: id});

        // Parse globalRootConfirmerSetInfo
        generatorParams.generatorInfo.numOperators = uint256(toml.readUint(".globalRootConfirmerSetInfo.numOperators"));
        generatorParams.generatorInfo.operatorInfoTreeRoot =
            toml.readBytes32(".globalRootConfirmerSetInfo.operatorInfoTreeRoot");
        generatorParams.generatorInfo.totalWeights = toml.readUintArray(".globalRootConfirmerSetInfo.totalWeights");
        uint256 apkX = toml.readUint(".globalRootConfirmerSetInfo.aggregatePubkey.X");
        uint256 apkY = toml.readUint(".globalRootConfirmerSetInfo.aggregatePubkey.Y");
        generatorParams.generatorInfo.aggregatePubkey = BN254.G1Point({X: apkX, Y: apkY});

        return generatorParams;
    }

    struct GeneratorParams {
        OperatorSet generator;
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo generatorInfo;
    }
}
