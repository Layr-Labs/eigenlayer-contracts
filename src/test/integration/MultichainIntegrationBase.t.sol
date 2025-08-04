// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationBase.t.sol";
import "src/contracts/multichain/BN254CertificateVerifier.sol";
import "src/contracts/multichain/ECDSACertificateVerifier.sol";
import "src/contracts/multichain/OperatorTableUpdater.sol";
import "src/contracts/multichain/CrossChainRegistry.sol";
import "src/contracts/permissions/KeyRegistrar.sol";
import "src/test/mocks/OperatorTableCalculatorMock.sol";
import "src/contracts/libraries/BN254.sol";
import "src/contracts/libraries/Merkle.sol";
import "src/contracts/interfaces/IOperatorTableCalculator.sol";
import "src/contracts/interfaces/IBN254CertificateVerifier.sol";
import "src/contracts/interfaces/IECDSACertificateVerifier.sol";
import "src/contracts/interfaces/IKeyRegistrar.sol";
import "src/contracts/interfaces/ICrossChainRegistry.sol";
import "src/contracts/interfaces/IAllocationManager.sol";
import "src/contracts/interfaces/IOperatorTableUpdater.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

/**
 * @title MultichainIntegrationBase
 * @notice Base contract for multichain integration tests
 * @dev Provides deployment, setup, and utility methods for multichain contracts
 */
abstract contract MultichainIntegrationBase is IntegrationBase {
    using StdStyle for *;
    using BN254 for BN254.G1Point;

    // Multichain contracts
    CrossChainRegistry public crossChainRegistry;
    CrossChainRegistry public crossChainRegistryImplementation;
    OperatorTableUpdater public operatorTableUpdater;
    OperatorTableUpdater public operatorTableUpdaterImplementation;
    BN254CertificateVerifier public bn254CertificateVerifier;
    BN254CertificateVerifier public bn254CertificateVerifierImplementation;
    ECDSACertificateVerifier public ecdsaCertificateVerifier;
    ECDSACertificateVerifier public ecdsaCertificateVerifierImplementation;
    OperatorTableCalculatorMock public operatorTableCalculatorMock;

    // Test constants
    uint constant CHAIN_ID_SOURCE = 1;
    uint constant CHAIN_ID_DESTINATION = 42_161; // Arbitrum
    uint constant DEFAULT_GENERATION_ID = 1;
    uint32 constant DEFAULT_QUORUM_THRESHOLD = 6667; // 66.67%

    // Operator key generation helpers
    uint constant BN254_PRIVATE_KEY_1 = 0x1;
    uint constant BN254_PRIVATE_KEY_2 = 0x2;
    uint constant BN254_PRIVATE_KEY_3 = 0x3;

    // ECDSA private keys for testing
    uint constant ECDSA_PRIVATE_KEY_1 = 0x11;
    uint constant ECDSA_PRIVATE_KEY_2 = 0x22;
    uint constant ECDSA_PRIVATE_KEY_3 = 0x33;

    /**
     * Deployment and Setup Methods
     */

    /**
     * @notice Deploy all multichain contracts
     */
    function _deployContracts() internal {
        // The base contracts are already deployed by IntegrationDeployer
        // We just need to deploy the multichain-specific contracts
        _deployMultichainContracts();
    }

    /**
     * @notice Deploy multichain-specific contracts
     */
    function _deployMultichainContracts() internal {
        // Deploy OperatorTableCalculatorMock
        operatorTableCalculatorMock = new OperatorTableCalculatorMock();

        // Deploy CrossChainRegistry with required dependencies
        crossChainRegistryImplementation =
            new CrossChainRegistry(allocationManager, keyRegistrar, permissionController, eigenLayerPauserReg, "1.0.0");

        crossChainRegistry = CrossChainRegistry(
            address(new TransparentUpgradeableProxy(address(crossChainRegistryImplementation), address(eigenLayerProxyAdmin), ""))
        );

        // Initialize CrossChainRegistry
        crossChainRegistry.initialize(
            address(this), // owner
            1 hours, // initialTableUpdateCadence (1 hour = 3600 seconds)
            0 // initialPausedStatus
        );

        // Deploy OperatorTableUpdater (placeholder addresses for certificate verifiers initially)
        emptyContract = new EmptyContract();
        operatorTableUpdater =
            OperatorTableUpdater(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));

        // Deploy BN254CertificateVerifier
        bn254CertificateVerifierImplementation = new BN254CertificateVerifier(IOperatorTableUpdater(address(operatorTableUpdater)), "1.0.0");

        bn254CertificateVerifier = BN254CertificateVerifier(
            address(new TransparentUpgradeableProxy(address(bn254CertificateVerifierImplementation), address(eigenLayerProxyAdmin), ""))
        );

        // Deploy ECDSACertificateVerifier
        ecdsaCertificateVerifierImplementation = new ECDSACertificateVerifier(IOperatorTableUpdater(address(operatorTableUpdater)), "1.0.0");

        ecdsaCertificateVerifier = ECDSACertificateVerifier(
            address(new TransparentUpgradeableProxy(address(ecdsaCertificateVerifierImplementation), address(eigenLayerProxyAdmin), ""))
        );

        // Now deploy new OperatorTableUpdater with correct certificate verifier addresses
        OperatorTableUpdater newOperatorTableUpdaterImplementation = new OperatorTableUpdater(
            IBN254CertificateVerifier(address(bn254CertificateVerifier)),
            IECDSACertificateVerifier(address(ecdsaCertificateVerifier)),
            eigenLayerPauserReg,
            "1.0.0"
        );

        // Upgrade the proxy to use the new implementation with correct addresses
        cheats.prank(address(eigenLayerProxyAdmin.owner()));
        eigenLayerProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(operatorTableUpdater))), address(newOperatorTableUpdaterImplementation)
        );

        // Create a real BN254OperatorSetInfo using a single operator
        uint aggSignerPrivKey = 69; // Global root confirmer private key

        uint[] memory globalRootConfirmerPrivKeys = new uint[](1);
        globalRootConfirmerPrivKeys[0] = aggSignerPrivKey;

        // Generate G1 public key using scalar multiplication
        BN254.G1Point memory aggregatePubkey = BN254.generatorG1().scalar_mul(globalRootConfirmerPrivKeys[0]);

        // Create real operator info array
        IOperatorTableCalculatorTypes.BN254OperatorInfo[] memory realOperatorInfos =
            new IOperatorTableCalculatorTypes.BN254OperatorInfo[](1);

        // Operator info
        uint[] memory operatorWeights = new uint[](1);
        operatorWeights[0] = 10_000; // Weight
        realOperatorInfos[0] = IOperatorTableCalculatorTypes.BN254OperatorInfo({pubkey: aggregatePubkey, weights: operatorWeights});

        // Create merkle tree of real operator infos
        bytes32[] memory operatorLeaves = new bytes32[](1);
        operatorLeaves[0] = keccak256(abi.encode(realOperatorInfos[0]));
        bytes32 operatorInfoTreeRoot = Merkle.merkleizeKeccak(operatorLeaves);

        // Calculate total weights
        uint[] memory totalWeights = new uint[](1);
        totalWeights[0] = operatorWeights[0];

        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory initialOperatorSetInfo = IOperatorTableCalculatorTypes
            .BN254OperatorSetInfo({
            operatorInfoTreeRoot: operatorInfoTreeRoot,
            numOperators: 1,
            aggregatePubkey: aggregatePubkey,
            totalWeights: totalWeights
        });

        ICrossChainRegistryTypes.OperatorSetConfig memory initialOperatorSetConfig =
            ICrossChainRegistryTypes.OperatorSetConfig({owner: address(0xDEADBEEF), maxStalenessPeriod: 0});

        OperatorSet memory globalRootConfirmerSet = OperatorSet({avs: address(0x6), id: 1});

        // Compute the initial global table root containing the global root confirmer set
        bytes memory globalRootConfirmerTable = abi.encode(
            globalRootConfirmerSet, IKeyRegistrarTypes.CurveType.BN254, initialOperatorSetConfig, abi.encode(initialOperatorSetInfo)
        );
        bytes32 globalRootConfirmerLeafHash = keccak256(globalRootConfirmerTable);

        // Create a simple merkle tree with just the global root confirmer set
        bytes32[] memory initialLeaves = new bytes32[](1);
        initialLeaves[0] = globalRootConfirmerLeafHash;
        bytes32 initialGlobalTableRoot = Merkle.merkleizeKeccak(initialLeaves);

        // Initialize OperatorTableUpdater with correct parameter order
        operatorTableUpdater.initialize(
            address(this), // owner
            0, // initialPausedStatus
            globalRootConfirmerSet, // globalRootConfirmerSet
            6600, // globalRootConfirmationThreshold (66%)
            initialOperatorSetInfo // globalRootConfirmerSetInfo
        );

        console.log("Multichain contracts deployed successfully");
    }

    function _setupMultichainContracts() internal {
        _deployMultichainContracts();

        // Register this test contract as an AVS (needed for AllocationManager.createOperatorSets)
        allocationManager.updateAVSMetadataURI(address(this), "https://test-avs.com");

        console.log("Multichain contracts setup completed");
    }

    /**
     * User and Operator Helper Methods
     */

    // Operator set IDs for different curve types (each operator set supports only one curve type)
    uint32 constant BN254_OPERATOR_SET_ID = 1;
    uint32 constant ECDSA_OPERATOR_SET_ID = 2;

    // Known valid private keys from unit tests (with corresponding G2 points)
    uint[] private validBN254PrivateKeys = [69, 123];
    uint private bn254KeyIndex = 0;
    uint private ecdsaPrivateKeyCounter = 0x1234567890123456789012345678901234567890123456789012345678901234;

    /// @dev Create a new operator with BN254 key registered
    function _newOperatorWithBN254Key(string memory name) internal returns (User operator, BN254.G1Point memory pubkey) {
        require(bn254KeyIndex < validBN254PrivateKeys.length, "Not enough valid BN254 private keys");

        operator = _newRandomOperator_NoAssets();

        // Use known valid BN254 private key
        uint privateKey = validBN254PrivateKeys[bn254KeyIndex++];
        pubkey = BN254.generatorG1().scalar_mul(privateKey);

        // Register BN254 key
        _registerBN254Key(operator, pubkey, privateKey);

        return (operator, pubkey);
    }

    /// @dev Create a new operator with ECDSA key registered
    function _newOperatorWithECDSAKey(string memory name) internal returns (User operator, address pubkey) {
        operator = _newRandomOperator_NoAssets();

        // Generate unique ECDSA keypair
        uint privateKey = ecdsaPrivateKeyCounter++;
        pubkey = vm.addr(privateKey);

        // Register ECDSA key
        _registerECDSAKey(operator, pubkey, privateKey);

        return (operator, pubkey);
    }

    /// @dev Register a BN254 key for an operator using the real KeyRegistrar
    function _registerBN254Key(User operator, BN254.G1Point memory pubkey, uint privateKey) internal {
        // Create operator set for BN254 testing
        OperatorSet memory operatorSet = OperatorSet({avs: address(this), id: BN254_OPERATOR_SET_ID});

        // Configure the operator set for BN254 curve type (only once)
        try keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254) {
            // Configuration successful
        } catch {
            // Already configured, which is fine
        }

        // Get the corresponding G2 point for this private key
        BN254.G2Point memory g2Point;
        if (privateKey == 69) {
            // G2 point for private key 69
            g2Point.X[1] = 19_101_821_850_089_705_274_637_533_855_249_918_363_070_101_489_527_618_151_493_230_256_975_900_223_847;
            g2Point.X[0] = 5_334_410_886_741_819_556_325_359_147_377_682_006_012_228_123_419_628_681_352_847_439_302_316_235_957;
            g2Point.Y[1] = 354_176_189_041_917_478_648_604_979_334_478_067_325_821_134_838_555_150_300_539_079_146_482_658_331;
            g2Point.Y[0] = 4_185_483_097_059_047_421_902_184_823_581_361_466_320_657_066_600_218_863_748_375_739_772_335_928_910;
        } else if (privateKey == 123) {
            // G2 point for private key 123
            g2Point.X[1] = 19_276_105_129_625_393_659_655_050_515_259_006_463_014_579_919_681_138_299_520_812_914_148_935_621_072;
            g2Point.X[0] = 14_066_454_060_412_929_535_985_836_631_817_650_877_381_034_334_390_275_410_072_431_082_437_297_539_867;
            g2Point.Y[1] = 12_642_665_914_920_339_463_975_152_321_804_664_028_480_770_144_655_934_937_445_922_690_262_428_344_269;
            g2Point.Y[0] = 10_109_651_107_942_685_361_120_988_628_892_759_706_059_655_669_161_016_107_907_096_760_613_704_453_218;
        } else {
            revert("Unknown private key - no corresponding G2 point available");
        }

        // Encode the key data
        bytes memory keyData = abi.encode(pubkey.X, pubkey.Y, g2Point.X, g2Point.Y);

        // Generate BN254 signature
        bytes32 messageHash = keyRegistrar.getBN254KeyRegistrationMessageHash(address(operator), operatorSet, keyData);
        BN254.G1Point memory msgPoint = BN254.hashToG1(messageHash);
        BN254.G1Point memory signature = msgPoint.scalar_mul(privateKey);
        bytes memory signatureBytes = abi.encode(signature.X, signature.Y);

        // Register the key directly with KeyRegistrar
        vm.prank(address(operator));
        keyRegistrar.registerKey(address(operator), operatorSet, keyData, signatureBytes);
    }

    /// @dev Register an ECDSA key for an operator using the real KeyRegistrar
    function _registerECDSAKey(User operator, address pubkey, uint privateKey) internal {
        // Create operator set for ECDSA testing
        OperatorSet memory operatorSet = OperatorSet({avs: address(this), id: ECDSA_OPERATOR_SET_ID});

        // Configure the operator set for ECDSA curve type (only once)
        try keyRegistrar.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.ECDSA) {
            // Configuration successful
        } catch {
            // Already configured, which is fine
        }

        // Encode the key data
        bytes memory keyData = abi.encodePacked(pubkey);

        // Generate ECDSA signature
        bytes32 messageHash = keyRegistrar.getECDSAKeyRegistrationMessageHash(address(operator), operatorSet, pubkey);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privateKey, messageHash);
        bytes memory signatureBytes = abi.encodePacked(r, s, v);

        // Register the key directly with KeyRegistrar
        vm.prank(address(operator));
        keyRegistrar.registerKey(address(operator), operatorSet, keyData, signatureBytes);
    }

    /**
     * Table Generation and Transport Simulation
     */

    /**
     * @notice Generate BN254 operator table for testing
     */
    function _generateBN254OperatorTable(OperatorSet memory operatorSet, User[] memory operators)
        internal
        view
        returns (IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory)
    {
        // Create operator infos array
        IOperatorTableCalculatorTypes.BN254OperatorInfo[] memory operatorInfos =
            new IOperatorTableCalculatorTypes.BN254OperatorInfo[](operators.length);

        // Aggregate public key (sum of all operator public keys)
        BN254.G1Point memory aggregatePubkey = BN254.G1Point({X: 0, Y: 0});

        // Total weights for each weight type
        uint[] memory totalWeights = new uint[](2);

        for (uint i = 0; i < operators.length; i++) {
            // Get operator's registered BN254 key
            (BN254.G1Point memory g1Point, BN254.G2Point memory g2Point) = keyRegistrar.getBN254Key(operatorSet, address(operators[i]));

            // Create mock weights for testing
            uint[] memory weights = new uint[](2);
            weights[0] = 1000 + i * 100; // First weight type
            weights[1] = 2000 + i * 200; // Second weight type

            operatorInfos[i] = IOperatorTableCalculatorTypes.BN254OperatorInfo({pubkey: g1Point, weights: weights});

            // Add to aggregate public key
            aggregatePubkey = BN254.plus(aggregatePubkey, g1Point);

            // Add to total weights
            totalWeights[0] += weights[0];
            totalWeights[1] += weights[1];
        }

        // Create merkle tree of operator infos
        bytes32[] memory leaves = new bytes32[](operatorInfos.length);
        for (uint i = 0; i < operatorInfos.length; i++) {
            leaves[i] = keccak256(abi.encode(operatorInfos[i]));
        }
        bytes32 operatorInfoTreeRoot = Merkle.merkleizeKeccak(leaves);

        return IOperatorTableCalculatorTypes.BN254OperatorSetInfo({
            operatorInfoTreeRoot: operatorInfoTreeRoot,
            numOperators: operatorInfos.length,
            aggregatePubkey: aggregatePubkey,
            totalWeights: totalWeights
        });
    }

    /**
     * @notice Generate ECDSA operator table for testing
     */
    function _generateECDSAOperatorTable(OperatorSet memory operatorSet, User[] memory operators)
        internal
        view
        returns (IOperatorTableCalculatorTypes.ECDSAOperatorInfo[] memory)
    {
        IOperatorTableCalculatorTypes.ECDSAOperatorInfo[] memory operatorInfos =
            new IOperatorTableCalculatorTypes.ECDSAOperatorInfo[](operators.length);

        for (uint i = 0; i < operators.length; i++) {
            // Get operator's registered ECDSA key
            address pubkey = keyRegistrar.getECDSAAddress(operatorSet, address(operators[i]));

            // Create mock weights for testing
            uint[] memory weights = new uint[](2);
            weights[0] = 1000 + i * 100; // First weight type
            weights[1] = 2000 + i * 200; // Second weight type

            operatorInfos[i] = IOperatorTableCalculatorTypes.ECDSAOperatorInfo({pubkey: pubkey, weights: weights});
        }

        return operatorInfos;
    }

    /// @dev Simulate transport of operator table to destination chain
    function _simulateTableTransport(uint sourceChainId, uint destinationChainId, uint generationId) internal {
        console.log("Simulating table transport from chain %d to chain %d", sourceChainId, destinationChainId);
        console.log("Generation ID: %d", generationId);

        // In a real implementation, this would involve cross-chain messaging
        // For testing, we just log the simulation
    }

    /**
     * Certificate Generation and Verification Helpers
     */

    /**
     * @notice Generate BN254 certificate for testing
     */
    function _generateBN254Certificate(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp,
        bytes32 messageHash,
        User[] memory operators,
        uint[] memory privateKeys
    ) internal view returns (IBN254CertificateVerifierTypes.BN254Certificate memory) {
        // For testing, assume all operators are signers (no non-signers)
        // This simplifies the certificate generation

        // Calculate aggregate signature
        BN254.G1Point memory aggregateSignature = BN254.G1Point({X: 0, Y: 0});
        BN254.G2Point memory aggregatePubkey;

        // Use appropriate G2 public key based on the signing private key
        if (privateKeys.length == 1 && privateKeys[0] == 123) {
            // G2 point for private key 123
            aggregatePubkey.X[1] = 19_276_105_129_625_393_659_655_050_515_259_006_463_014_579_919_681_138_299_520_812_914_148_935_621_072;
            aggregatePubkey.X[0] = 14_066_454_060_412_929_535_985_836_631_817_650_877_381_034_334_390_275_410_072_431_082_437_297_539_867;
            aggregatePubkey.Y[1] = 12_642_665_914_920_339_463_975_152_321_804_664_028_480_770_144_655_934_937_445_922_690_262_428_344_269;
            aggregatePubkey.Y[0] = 10_109_651_107_942_685_361_120_988_628_892_759_706_059_655_669_161_016_107_907_096_760_613_704_453_218;
        } else {
            // Default to private key 69
            aggregatePubkey.X[1] = 19_101_821_850_089_705_274_637_533_855_249_918_363_070_101_489_527_618_151_493_230_256_975_900_223_847;
            aggregatePubkey.X[0] = 5_334_410_886_741_819_556_325_359_147_377_682_006_012_228_123_419_628_681_352_847_439_302_316_235_957;
            aggregatePubkey.Y[1] = 354_176_189_041_917_478_648_604_979_334_478_067_325_821_134_838_555_150_300_539_079_146_482_658_331;
            aggregatePubkey.Y[0] = 4_185_483_097_059_047_421_902_184_823_581_361_466_320_657_066_600_218_863_748_375_739_772_335_928_910;
        }

        // Generate signatures for each operator
        for (uint i = 0; i < operators.length; i++) {
            BN254.G1Point memory msgHashG1 = BN254.hashToG1(messageHash);
            BN254.G1Point memory signature = BN254.scalar_mul(msgHashG1, privateKeys[i]);
            aggregateSignature = BN254.plus(aggregateSignature, signature);
        }

        // No non-signers in this test scenario
        IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[] memory nonSignerWitnesses =
            new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](0);

        return IBN254CertificateVerifierTypes.BN254Certificate({
            referenceTimestamp: referenceTimestamp,
            messageHash: messageHash,
            signature: aggregateSignature,
            apk: aggregatePubkey,
            nonSignerWitnesses: nonSignerWitnesses
        });
    }

    /**
     * @notice Generate ECDSA certificate for testing
     */
    function _generateECDSACertificate(
        OperatorSet memory operatorSet,
        uint32 referenceTimestamp,
        bytes32 messageHash,
        User[] memory operators,
        uint[] memory privateKeys
    ) internal view returns (IECDSACertificateVerifierTypes.ECDSACertificate memory cert) {
        require(operators.length == privateKeys.length, "Mismatched operators and keys");

        // Use the contract's digest calculation (same as unit tests)
        bytes32 signableDigest = ecdsaCertificateVerifier.calculateCertificateDigest(referenceTimestamp, messageHash);

        // Collect signers and their private keys
        uint numSigners = privateKeys.length;
        address[] memory signerAddresses = new address[](numSigners);
        bytes[] memory signaturesArr = new bytes[](numSigners);

        // Get signer addresses and create signatures (same pattern as unit tests)
        for (uint i = 0; i < numSigners; i++) {
            signerAddresses[i] = vm.addr(privateKeys[i]);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(privateKeys[i], signableDigest);
            signaturesArr[i] = abi.encodePacked(r, s, v);
        }

        // Sort signers and signatures by address (ascending) - exact unit test pattern
        for (uint i = 0; i < numSigners; i++) {
            for (uint j = i + 1; j < numSigners; j++) {
                if (signerAddresses[j] < signerAddresses[i]) {
                    // Swap addresses
                    address tmpAddr = signerAddresses[i];
                    signerAddresses[i] = signerAddresses[j];
                    signerAddresses[j] = tmpAddr;
                    // Swap signatures
                    bytes memory tmpSig = signaturesArr[i];
                    signaturesArr[i] = signaturesArr[j];
                    signaturesArr[j] = tmpSig;
                }
            }
        }

        // Concatenate signatures in sorted order (same as unit tests)
        bytes memory signatures;
        for (uint i = 0; i < numSigners; i++) {
            signatures = bytes.concat(signatures, signaturesArr[i]);
        }

        cert = IECDSACertificateVerifierTypes.ECDSACertificate({
            referenceTimestamp: referenceTimestamp,
            messageHash: messageHash,
            sig: signatures
        });
    }

    /**
     * Cross-Chain Registry Helpers
     */

    /// @dev Create a generation reservation using the real CrossChainRegistry
    function _createGenerationReservation(uint chainId, uint generationId, User operator, uint32 operatorSetId) internal {
        // Reference the operator set for the specific curve type
        OperatorSet memory operatorSet = OperatorSet({avs: address(this), id: operatorSetId});

        // Check if operator set exists, if not create it
        if (!allocationManager.isOperatorSet(operatorSet)) {
            IAllocationManagerTypes.CreateSetParams[] memory createParams = new IAllocationManagerTypes.CreateSetParams[](1);
            createParams[0] = IAllocationManagerTypes.CreateSetParams({
                operatorSetId: operatorSetId,
                strategies: new IStrategy[](0) // Empty strategies for now
            });

            allocationManager.createOperatorSets(address(this), createParams);
        }

        // Create chain IDs array
        uint[] memory chainIDs = new uint[](1);
        chainIDs[0] = chainId;

        // First, whitelist the chain ID (as owner)
        address[] memory operatorTableUpdaters = new address[](1);
        operatorTableUpdaters[0] = address(operatorTableUpdater);
        crossChainRegistry.addChainIDsToWhitelist(chainIDs, operatorTableUpdaters);

        // Check if generation reservation already exists
        OperatorSet[] memory activeReservations = crossChainRegistry.getActiveGenerationReservations();
        bool reservationExists = false;
        for (uint i = 0; i < activeReservations.length; i++) {
            if (activeReservations[i].avs == operatorSet.avs && activeReservations[i].id == operatorSet.id) {
                reservationExists = true;
                break;
            }
        }

        if (!reservationExists) {
            // Create operator set config
            ICrossChainRegistryTypes.OperatorSetConfig memory config = ICrossChainRegistryTypes.OperatorSetConfig({
                owner: address(operator),
                maxStalenessPeriod: 86_400 // 1 day
            });

            // Create generation reservation
            crossChainRegistry.createGenerationReservation(operatorSet, operatorTableCalculatorMock, config);
        }
    }

    /// @dev Update operator table in the certificate verifier after transport
    function _updateOperatorTableBN254(
        uint chainId,
        uint generationId,
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo,
        User operator,
        uint32 operatorSetId
    ) internal {
        console.log("Updating BN254 operator table after transport to chain %d", chainId);
        console.log("Generation %d", generationId);

        OperatorSet memory operatorSet = OperatorSet({avs: address(this), id: operatorSetId});

        ICrossChainRegistryTypes.OperatorSetConfig memory config = ICrossChainRegistryTypes.OperatorSetConfig({
            owner: address(operator),
            maxStalenessPeriod: 86_400 // 1 day
        });

        // Call the real updateOperatorTable function
        // Note: We need to prank as the OperatorTableUpdater since only it can call this function
        vm.prank(address(operatorTableUpdater));
        bn254CertificateVerifier.updateOperatorTable(operatorSet, uint32(block.timestamp), operatorSetInfo, config);
    }

    /// @dev Update operator table in the certificate verifier after transport
    function _updateOperatorTableECDSA(
        uint chainId,
        uint generationId,
        IOperatorTableCalculatorTypes.ECDSAOperatorInfo[] memory operatorInfos,
        User operator,
        uint32 operatorSetId
    ) internal {
        console.log("Updating ECDSA operator table after transport to chain %d", chainId);
        console.log("Generation %d", generationId);

        OperatorSet memory operatorSet = OperatorSet({avs: address(this), id: operatorSetId});

        ICrossChainRegistryTypes.OperatorSetConfig memory config = ICrossChainRegistryTypes.OperatorSetConfig({
            owner: address(operator),
            maxStalenessPeriod: 86_400 // 1 day
        });

        // Call the real updateOperatorTable function
        // Note: We need to prank as the OperatorTableUpdater since only it can call this function
        vm.prank(address(operatorTableUpdater));
        ecdsaCertificateVerifier.updateOperatorTable(operatorSet, uint32(block.timestamp), operatorInfos, config);
    }

    /**
     * Utility Methods
     */

    /// @dev Calculate hash of BN254 operator table
    function _hashBN254Table(IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo) internal pure returns (bytes32) {
        return keccak256(abi.encode(operatorSetInfo));
    }

    /// @dev Calculate hash of ECDSA operator table
    function _hashECDSATable(IOperatorTableCalculatorTypes.ECDSAOperatorInfo[] memory operatorInfos) internal pure returns (bytes32) {
        return keccak256(abi.encode(operatorInfos));
    }

    /// @dev Create test message hash
    function _createTestMessageHash(string memory message) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(message));
    }

    /// @dev Roll forward blocks to simulate time passage
    function _rollForwardBlocks(uint blocks) internal {
        cheats.roll(block.number + blocks);
    }

    /// @dev Roll forward time to simulate time passage
    function _rollForwardTime(uint seconds_) internal {
        cheats.warp(block.timestamp + seconds_);
    }

    /**
     * HELPER FUNCTIONS FOR MULTICHAIN TESTS
     */

    /// @dev Create and configure test operator set
    function _createTestOperatorSet(uint32 operatorSetId) internal {
        OperatorSet memory operatorSet = OperatorSet({avs: address(this), id: operatorSetId});

        bool exists = allocationManager.isOperatorSet(operatorSet);
        if (!exists) {
            IAllocationManagerTypes.CreateSetParams[] memory createSetParams = new IAllocationManagerTypes.CreateSetParams[](1);
            createSetParams[0] = IAllocationManagerTypes.CreateSetParams({operatorSetId: operatorSetId, strategies: new IStrategy[](0)});
            allocationManager.createOperatorSets(address(this), createSetParams);
        }
    }

    /// @dev Setup AVS registration and chain whitelist
    function _setupAVSAndChains() internal {
        // Register as AVS
        allocationManager.updateAVSMetadataURI(address(this), "test-avs");

        // Whitelist chain IDs for cross-chain operations
        uint[] memory chainIDs = new uint[](2);
        chainIDs[0] = 1; // Source chain
        chainIDs[1] = 137; // Destination chain

        address[] memory operatorTableUpdaters = new address[](2);
        operatorTableUpdaters[0] = address(operatorTableUpdater);
        operatorTableUpdaters[1] = address(operatorTableUpdater);

        crossChainRegistry.addChainIDsToWhitelist(chainIDs, operatorTableUpdaters);
    }

    /// @dev Create generation reservation for operator set
    function _createGenerationReservation(OperatorSet memory operatorSet) internal {
        // Check if generation reservation already exists
        OperatorSet[] memory activeReservations = crossChainRegistry.getActiveGenerationReservations();
        bool reservationExists = false;
        for (uint j = 0; j < activeReservations.length; j++) {
            if (activeReservations[j].avs == operatorSet.avs && activeReservations[j].id == operatorSet.id) {
                reservationExists = true;
                break;
            }
        }

        if (!reservationExists) {
            IOperatorTableCalculator mockCalculator = IOperatorTableCalculator(address(operatorTableCalculatorMock));

            crossChainRegistry.createGenerationReservation(
                operatorSet, mockCalculator, ICrossChainRegistryTypes.OperatorSetConfig({owner: address(this), maxStalenessPeriod: 3600})
            );
        }
    }

    /// @dev Register BN254 keys for operators in operator set
    function _registerBN254Keys(OperatorSet memory operatorSet) internal returns (User[] memory operators) {
        operators = new User[](1);
        uint[] memory bn254PrivKeys = _getBN254PrivateKeys();

        for (uint i = 0; i < 1; i++) {
            operators[i] = _newRandomOperator_NoAssets();
            _registerSingleBN254Key(operators[i], operatorSet, bn254PrivKeys[i]);
        }
    }

    /// @dev Register ECDSA keys for operators in operator set
    function _registerECDSAKeys(OperatorSet memory operatorSet) internal returns (User[] memory operators) {
        operators = new User[](3);
        uint[] memory ecdsaPrivKeys = _getECDSAPrivateKeys();

        for (uint i = 0; i < 3; i++) {
            operators[i] = _newRandomOperator_NoAssets();
            _registerSingleECDSAKey(operators[i], operatorSet, ecdsaPrivKeys[i]);
        }
    }

    /// @dev Register single BN254 key for operator
    function _registerSingleBN254Key(User operator, OperatorSet memory operatorSet, uint privKey) internal {
        // Generate G1 point and get corresponding G2 point
        BN254.G1Point memory pubkey = BN254.generatorG1().scalar_mul(privKey);
        BN254.G2Point memory g2Pubkey = _getBN254G2Point(privKey);

        // Encode key data for registration
        bytes memory keyData = abi.encode(pubkey.X, pubkey.Y, g2Pubkey.X, g2Pubkey.Y);

        // Generate registration signature
        bytes32 messageHash = keyRegistrar.getBN254KeyRegistrationMessageHash(address(operator), operatorSet, keyData);
        BN254.G1Point memory msgPoint = BN254.hashToG1(messageHash);
        BN254.G1Point memory signature = msgPoint.scalar_mul(privKey);
        bytes memory signatureBytes = abi.encode(signature.X, signature.Y);

        // Register key with KeyRegistrar
        vm.prank(address(operator));
        keyRegistrar.registerKey(address(operator), operatorSet, keyData, signatureBytes);
    }

    /// @dev Register single ECDSA key for operator
    function _registerSingleECDSAKey(User operator, OperatorSet memory operatorSet, uint privKey) internal {
        // Generate ECDSA key pair
        address pubkey = vm.addr(privKey);
        bytes memory keyData = abi.encodePacked(pubkey);

        // Generate registration signature
        bytes32 messageHash = keyRegistrar.getECDSAKeyRegistrationMessageHash(address(operator), operatorSet, pubkey);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privKey, messageHash);
        bytes memory signature = abi.encodePacked(r, s, v);

        // Register key with KeyRegistrar
        vm.prank(address(operator));
        keyRegistrar.registerKey(address(operator), operatorSet, keyData, signature);
    }

    /// @dev Update global table root and BN254 operator table
    function _confirmGlobalTableRootAndUpdateBN254(
        OperatorSet memory operatorSet,
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo,
        uint32 referenceTimestamp
    ) internal {
        // Create the operator table data for transport
        bytes memory operatorTable = abi.encode(
            operatorSet,
            IKeyRegistrarTypes.CurveType.BN254,
            ICrossChainRegistryTypes.OperatorSetConfig({owner: address(this), maxStalenessPeriod: 3600}),
            abi.encode(operatorSetInfo)
        );

        // Create global table root containing the operator table
        bytes32 operatorSetLeafHash = operatorTableUpdater.calculateOperatorTableLeaf(operatorTable);
        bytes32[] memory leaves = new bytes32[](1);
        leaves[0] = operatorSetLeafHash;
        bytes32 globalTableRoot = Merkle.merkleizeKeccak(leaves);

        // Update the global table root with confirmation
        _updateGlobalTableRoot(globalTableRoot, referenceTimestamp);

        // Update the operator table using the confirmed global root
        uint32 operatorSetIndex = 0; // Single leaf in merkle tree
        bytes memory proof = new bytes(0); // Empty proof for single leaf

        operatorTableUpdater.updateOperatorTable(referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorTable);
    }

    /// @dev Update global table root and ECDSA operator table
    function _confirmGlobalTableRootAndUpdateECDSA(
        OperatorSet memory operatorSet,
        IOperatorTableCalculatorTypes.ECDSAOperatorInfo[] memory operatorInfos,
        uint32 referenceTimestamp
    ) internal {
        // Create the operator table data for transport
        bytes memory operatorTable = abi.encode(
            operatorSet,
            IKeyRegistrarTypes.CurveType.ECDSA,
            ICrossChainRegistryTypes.OperatorSetConfig({owner: address(this), maxStalenessPeriod: 3600}),
            abi.encode(operatorInfos)
        );

        // Create global table root containing the operator table
        bytes32 operatorSetLeafHash = operatorTableUpdater.calculateOperatorTableLeaf(operatorTable);
        bytes32[] memory leaves = new bytes32[](1);
        leaves[0] = operatorSetLeafHash;
        bytes32 globalTableRoot = Merkle.merkleizeKeccak(leaves);

        // Update the global table root with confirmation
        _updateGlobalTableRoot(globalTableRoot, referenceTimestamp);

        // Update the operator table using the confirmed global root
        uint32 operatorSetIndex = 0; // Single leaf in merkle tree
        bytes memory proof = new bytes(0); // Empty proof for single leaf

        operatorTableUpdater.updateOperatorTable(referenceTimestamp, globalTableRoot, operatorSetIndex, proof, operatorTable);
    }

    /// @dev Updates the global table root with certificate verification
    function _updateGlobalTableRoot(bytes32 globalTableRoot, uint32 confirmationTimestamp) internal {
        uint32 referenceBlockNumber = uint32(block.number);

        // Get the actual global root confirmer set from the OperatorTableUpdater
        OperatorSet memory globalRootConfirmerSet = operatorTableUpdater.getGenerator();

        // Get the global root confirmer set's reference timestamp (from initialization)
        uint32 globalRootConfirmerTimestamp = operatorTableUpdater.getGeneratorReferenceTimestamp();

        // Use the proper message hash from the OperatorTableUpdater
        bytes32 messageHash =
            operatorTableUpdater.getGlobalTableUpdateMessageHash(globalTableRoot, confirmationTimestamp, referenceBlockNumber);

        // Create certificate for global table root confirmation using the correct timestamp
        IBN254CertificateVerifierTypes.BN254Certificate memory confirmationCertificate =
            _generateGlobalRootConfirmationCertificate(globalRootConfirmerSet, globalRootConfirmerTimestamp, messageHash);

        // Verify certificate to confirm global table root
        operatorTableUpdater.confirmGlobalTableRoot(confirmationCertificate, globalTableRoot, confirmationTimestamp, referenceBlockNumber);
    }

    /// @dev Generate BN254 certificate for global root confirmation
    function _generateGlobalRootConfirmationCertificate(OperatorSet memory operatorSet, uint32 referenceTimestamp, bytes32 messageHash)
        internal
        view
        returns (IBN254CertificateVerifierTypes.BN254Certificate memory)
    {
        uint aggSignerPrivKey = 69; // Private key for global root confirmer
        BN254.G1Point memory aggregateSignature = BN254.hashToG1(messageHash).scalar_mul(aggSignerPrivKey);

        // Use the hardcoded G2 aggregate public key for private key 69
        BN254.G2Point memory aggregateG2Pubkey;
        aggregateG2Pubkey.X[1] = 19_101_821_850_089_705_274_637_533_855_249_918_363_070_101_489_527_618_151_493_230_256_975_900_223_847;
        aggregateG2Pubkey.X[0] = 5_334_410_886_741_819_556_325_359_147_377_682_006_012_228_123_419_628_681_352_847_439_302_316_235_957;
        aggregateG2Pubkey.Y[1] = 354_176_189_041_917_478_648_604_979_334_478_067_325_821_134_838_555_150_300_539_079_146_482_658_331;
        aggregateG2Pubkey.Y[0] = 4_185_483_097_059_047_421_902_184_823_581_361_466_320_657_066_600_218_863_748_375_739_772_335_928_910;

        // No non-signers in this test scenario
        IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[] memory nonSignerWitnesses =
            new IBN254CertificateVerifierTypes.BN254OperatorInfoWitness[](0);

        return IBN254CertificateVerifierTypes.BN254Certificate({
            referenceTimestamp: referenceTimestamp,
            messageHash: messageHash,
            signature: aggregateSignature,
            apk: aggregateG2Pubkey,
            nonSignerWitnesses: nonSignerWitnesses
        });
    }

    /**
     * PRIVATE KEY HELPERS
     */
    function _getBN254PrivateKeys() internal pure returns (uint[] memory) {
        uint[] memory keys = new uint[](1);
        keys[0] = 123; // Main test operator key
        return keys;
    }

    function _getECDSAPrivateKeys() internal pure returns (uint[] memory) {
        uint[] memory keys = new uint[](3);
        keys[0] = 0x1234567890123456789012345678901234567890123456789012345678901234;
        keys[1] = 0x9876543210987654321098765432109876543210987654321098765432109876;
        keys[2] = 0x1111111111111111111111111111111111111111111111111111111111111111;
        return keys;
    }

    function _getGlobalRootConfirmerPrivateKeys() internal pure returns (uint[] memory) {
        uint[] memory keys = new uint[](1);
        keys[0] = 69; // Global root confirmer key (avoids collision with main test)
        return keys;
    }

    function _getBN254G2Point(uint privKey) internal pure returns (BN254.G2Point memory g2Point) {
        if (privKey == 69) {
            // G2 point for private key 69
            g2Point.X[1] = 19_101_821_850_089_705_274_637_533_855_249_918_363_070_101_489_527_618_151_493_230_256_975_900_223_847;
            g2Point.X[0] = 5_334_410_886_741_819_556_325_359_147_377_682_006_012_228_123_419_628_681_352_847_439_302_316_235_957;
            g2Point.Y[1] = 354_176_189_041_917_478_648_604_979_334_478_067_325_821_134_838_555_150_300_539_079_146_482_658_331;
            g2Point.Y[0] = 4_185_483_097_059_047_421_902_184_823_581_361_466_320_657_066_600_218_863_748_375_739_772_335_928_910;
        } else if (privKey == 123) {
            // G2 point for private key 123
            g2Point.X[1] = 19_276_105_129_625_393_659_655_050_515_259_006_463_014_579_919_681_138_299_520_812_914_148_935_621_072;
            g2Point.X[0] = 14_066_454_060_412_929_535_985_836_631_817_650_877_381_034_334_390_275_410_072_431_082_437_297_539_867;
            g2Point.Y[1] = 12_642_665_914_920_339_463_975_152_321_804_664_028_480_770_144_655_934_937_445_922_690_262_428_344_269;
            g2Point.Y[0] = 10_109_651_107_942_685_361_120_988_628_892_759_706_059_655_669_161_016_107_907_096_760_613_704_453_218;
        } else {
            revert("Unknown private key");
        }
    }

    /**
     * Initialization Override
     */
    function setUp() public virtual override {
        super.setUp();

        // Initialize userTypes and assetTypes arrays to avoid underflow in random functions
        _configUserTypes(DEFAULT | ALT_METHODS);
        _configAssetTypes(HOLDS_LST | HOLDS_ETH | HOLDS_ALL);

        _setupMultichainContracts();
    }
}
