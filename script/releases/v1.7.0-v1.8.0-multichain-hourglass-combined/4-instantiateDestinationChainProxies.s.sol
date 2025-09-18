// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import {DeployDestinationChainProxies} from "./2-deployDestinationChainProxies.s.sol";
import {DeployDestinationChainImpls} from "./3-deployDestinationChainImpls.s.sol";
import {CrosschainDeployLib} from "script/releases/CrosschainDeployLib.sol";
import "../Env.sol";

import "src/contracts/interfaces/ICrossChainRegistry.sol";
import "src/contracts/interfaces/IOperatorTableCalculator.sol";
import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/libraries/BN254.sol";

// For TOML parsing
import {stdToml} from "forge-std/StdToml.sol";

/**
 * Purpose: Upgrade proxies to point to implementations and transfer control to ProxyAdmin.
 */
contract InstantiateDestinationChainProxies is DeployDestinationChainImpls {
    using Env for *;
    using OperatorSetLib for OperatorSet;
    using stdToml for string;

    /// forgefmt: disable-next-item
    function _runAsMultisig() internal override prank(Env.multichainDeployerMultisig()) {
        // If we're not on a destination chain or we're on a version that already has these contracts deployed, we don't need to do anything
        if (!Env.isDestinationChain() || _isAlreadyDeployed()) {
            return;
        }

        // Upgrade the proxies to point to the actual implementations
        // ECDSACertificateVerifier
        ITransparentUpgradeableProxy ecdsaCertificateVerifierProxy =
            ITransparentUpgradeableProxy(payable(address(Env.proxy.ecdsaCertificateVerifier())));
        ecdsaCertificateVerifierProxy.upgradeTo(address(Env.impl.ecdsaCertificateVerifier()));

        // BN254CertificateVerifier
        ITransparentUpgradeableProxy bn254CertificateVerifierProxy =
            ITransparentUpgradeableProxy(payable(address(Env.proxy.bn254CertificateVerifier())));
        bn254CertificateVerifierProxy.upgradeTo(address(Env.impl.bn254CertificateVerifier()));

        // OperatorTableUpdater - we also initialize this contract
        OperatorTableUpdaterInitParams memory operatorTableUpdaterInitParams = _getTableUpdaterInitParams();
        ITransparentUpgradeableProxy operatorTableUpdaterProxy =
            ITransparentUpgradeableProxy(payable(address(Env.proxy.operatorTableUpdater())));
        operatorTableUpdaterProxy.upgradeToAndCall(
            address(Env.impl.operatorTableUpdater()),
            abi.encodeCall(
                OperatorTableUpdater.initialize,
                (
                    Env.opsMultisig(),
                    0, // initial paused status
                    operatorTableUpdaterInitParams.generator,
                    operatorTableUpdaterInitParams.globalRootConfirmationThreshold,
                    operatorTableUpdaterInitParams.generatorInfo
                )
            )
        );

        // TaskMailbox - we also initialize this contract
        TaskMailboxInitParams memory taskMailboxInitParams = _getTaskMailboxInitParams();
        ITransparentUpgradeableProxy taskMailboxProxy =
            ITransparentUpgradeableProxy(payable(address(Env.proxy.taskMailbox())));
        taskMailboxProxy.upgradeToAndCall(
            address(Env.impl.taskMailbox()),
            abi.encodeCall(
                TaskMailbox.initialize,
                (
                    taskMailboxInitParams.owner,
                    taskMailboxInitParams.feeSplit,
                    taskMailboxInitParams.feeSplitCollector
                )
            )
        );


        // Transfer proxy admin ownership
        operatorTableUpdaterProxy.changeAdmin(Env.proxyAdmin());
        ecdsaCertificateVerifierProxy.changeAdmin(Env.proxyAdmin());
        bn254CertificateVerifierProxy.changeAdmin(Env.proxyAdmin());
        taskMailboxProxy.changeAdmin(Env.proxyAdmin());
    }

    function testScript() public virtual override {
        if (!Env.isDestinationChain() || _isAlreadyDeployed()) {
            return;
        }

        // 1. Deploy the destination chain contracts
        // If proxies are not deployed, deploy them
        if (!_areProxiesDeployed()) {
            DeployDestinationChainProxies._runAsMultisig();
            _unsafeResetHasPranked(); // reset hasPranked so we can use it in the execute()
        } else {
            // Since the proxies are already deployed, we need to update the env with the proper addresses
            _addContractsToEnv();
        }

        // 2. Deploy the destination chain impls
        _mode = OperationalMode.EOA; // Set to EOA mode so we can deploy the impls in the EOA script
        DeployDestinationChainImpls._runAsEOA();

        // 3. Instantiate the destination chain proxies
        execute();

        // 4. Validate the destination chain
        _validateStorage();
        _validateProxyAdmins();
        _validateProxyConstructors();
        _validateProxiesInitialized();
    }

    /// @dev Validate that storage variables are set correctly
    function _validateStorage() internal view {
        // Validate OperatorTableUpdater
        {
            OperatorTableUpdater operatorTableUpdater = Env.proxy.operatorTableUpdater();
            assertTrue(operatorTableUpdater.owner() == Env.opsMultisig(), "operatorTableUpdater.owner invalid");
            assertTrue(operatorTableUpdater.paused() == 0, "operatorTableUpdater.paused invalid");

            // Checks on the generator
            OperatorTableUpdaterInitParams memory initParams = _getTableUpdaterInitParams();
            OperatorSet memory generator = operatorTableUpdater.getGenerator();
            assertEq(generator.key(), initParams.generator.key(), "operatorTableUpdater.generator invalid");
            assertEq(
                generator.avs,
                _getGeneratorAddress(), // The generator is set to the ops multisig of the *source* chain, hence we cannot use Env.opsMultisig().
                "operatorTableUpdater.generator.avs invalid"
            );
            assertEq(
                generator.id,
                0, // The generator is set to 0 on initialization
                "operatorTableUpdater.generator.id invalid"
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
            // latestReferenceTimestamp is set to block.timestamp during initialization
            assertTrue(
                operatorTableUpdater.getLatestReferenceTimestamp() > 0,
                "operatorTableUpdater.latestReferenceTimestamp should be > 0"
            );
            assertTrue(
                operatorTableUpdater.isRootValid(operatorTableUpdater.GENERATOR_GLOBAL_TABLE_ROOT()),
                "operatorTableUpdater.generatorGlobalTableRoot invalid"
            );
            assertTrue(
                operatorTableUpdater.isRootValidByTimestamp(operatorTableUpdater.GENERATOR_REFERENCE_TIMESTAMP()),
                "operatorTableUpdater.generatorGlobalTableRoot invalid"
            );
            ICrossChainRegistryTypes.OperatorSetConfig memory generatorConfig =
                operatorTableUpdater.getGeneratorConfig();
            assertEq(generatorConfig.maxStalenessPeriod, 0, "generatorConfig.maxStalenessPeriod invalid");
            assertEq(generatorConfig.owner, address(operatorTableUpdater), "generatorConfig.owner invalid");

            // Check the global root confirmation threshold is 10000
            assertEq(
                operatorTableUpdater.globalRootConfirmationThreshold(),
                10_000,
                "operatorTableUpdater.globalRootConfirmationThreshold invalid"
            );
        }

        // Validate ECDSACertificateVerifier
        {
            ECDSACertificateVerifier ecdsaCertificateVerifier = Env.proxy.ecdsaCertificateVerifier();
            assertTrue(address(ecdsaCertificateVerifier) != address(0), "ecdsaCertificateVerifier not deployed");
        }

        // Validate BN254CertificateVerifier
        {
            OperatorTableUpdaterInitParams memory initParams = _getTableUpdaterInitParams();
            BN254CertificateVerifier bn254CertificateVerifier = Env.proxy.bn254CertificateVerifier();
            assertTrue(address(bn254CertificateVerifier) != address(0), "bn254CertificateVerifier not deployed");

            // Check the generator info
            OperatorSet memory generator = initParams.generator;
            assertEq(
                bn254CertificateVerifier.latestReferenceTimestamp(generator),
                1,
                "bn254CertificateVerifier.latestReferenceTimestamp invalid"
            );
            IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory generatorInfo = initParams.generatorInfo;
            IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory returnedGeneratorInfo =
                bn254CertificateVerifier.getOperatorSetInfo(generator, 1);
            assertEq(
                returnedGeneratorInfo.operatorInfoTreeRoot,
                generatorInfo.operatorInfoTreeRoot,
                "bn254CertificateVerifier.operatorSetInfo.operatorInfoTreeRoot invalid"
            );
            assertEq(
                returnedGeneratorInfo.numOperators,
                generatorInfo.numOperators,
                "bn254CertificateVerifier.operatorSetInfo.numOperators invalid"
            );
            assertEq(
                returnedGeneratorInfo.aggregatePubkey.X,
                generatorInfo.aggregatePubkey.X,
                "bn254CertificateVerifier.operatorSetInfo.aggregatePubkey.X invalid"
            );
            assertEq(
                returnedGeneratorInfo.aggregatePubkey.Y,
                generatorInfo.aggregatePubkey.Y,
                "bn254CertificateVerifier.operatorSetInfo.aggregatePubkey.Y invalid"
            );
            assertEq(
                returnedGeneratorInfo.totalWeights.length,
                generatorInfo.totalWeights.length,
                "bn254CertificateVerifier.operatorSetInfo.totalWeights.length invalid"
            );
            for (uint256 i = 0; i < returnedGeneratorInfo.totalWeights.length; i++) {
                assertEq(
                    returnedGeneratorInfo.totalWeights[i],
                    generatorInfo.totalWeights[i],
                    "bn254CertificateVerifier.operatorSetInfo.totalWeights invalid"
                );
            }
            assertEq(
                bn254CertificateVerifier.getOperatorSetOwner(generator),
                address(Env.proxy.operatorTableUpdater()), // OperatorTableUpdater is the owner of the generator
                "bn254CertificateVerifier.operatorSetOwner invalid"
            );
            assertEq(
                bn254CertificateVerifier.maxOperatorTableStaleness(generator),
                0,
                "bn254CertificateVerifier.maxOperatorTableStaleness invalid"
            );
            assertEq(
                bn254CertificateVerifier.isReferenceTimestampSet(generator, 1),
                true,
                "bn254CertificateVerifier.isReferenceTimestampSet invalid"
            );
        }

        // Validate TaskMailbox
        {
            TaskMailbox taskMailbox = Env.proxy.taskMailbox();
            TaskMailboxInitParams memory initParams = _getTaskMailboxInitParams();
            assertTrue(address(taskMailbox) != address(0), "taskMailbox not deployed");

            assertTrue(taskMailbox.owner() == initParams.owner, "taskMailbox.owner invalid");
            assertEq(taskMailbox.feeSplit(), initParams.feeSplit, "taskMailbox.feeSplit invalid");
            assertTrue(
                taskMailbox.feeSplitCollector() == initParams.feeSplitCollector, "taskMailbox.feeSplitCollector invalid"
            );
        }
    }

    /// @dev Ensure each deployed TUP/beacon is owned by the proxyAdmin/executorMultisig
    function _validateProxyAdmins() internal view {
        address pa = Env.proxyAdmin();

        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.operatorTableUpdater())) == pa,
            "operatorTableUpdater proxyAdmin incorrect"
        );
        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.ecdsaCertificateVerifier())) == pa,
            "ecdsaCertificateVerifier proxyAdmin incorrect"
        );
        assertTrue(
            Env._getProxyAdmin(address(Env.proxy.bn254CertificateVerifier())) == pa,
            "bn254CertificateVerifier proxyAdmin incorrect"
        );
        assertTrue(Env._getProxyAdmin(address(Env.proxy.taskMailbox())) == pa, "taskMailbox proxyAdmin incorrect");
    }

    function _validateProxyConstructors() internal view {
        /// OperatorTableUpdater
        {
            OperatorTableUpdater operatorTableUpdater = Env.proxy.operatorTableUpdater();
            assertEq(operatorTableUpdater.version(), Env.deployVersion(), "operatorTableUpdater version mismatch");
            assertTrue(
                operatorTableUpdater.bn254CertificateVerifier() == Env.proxy.bn254CertificateVerifier(),
                "out.bn254CertificateVerifier mismatch"
            );
            assertTrue(
                operatorTableUpdater.ecdsaCertificateVerifier() == Env.proxy.ecdsaCertificateVerifier(),
                "out.ecdsaCertificateVerifier mismatch"
            );
        }

        /// ECDSACertificateVerifier
        {
            ECDSACertificateVerifier ecdsaCertificateVerifier = Env.proxy.ecdsaCertificateVerifier();
            assertEq(
                ecdsaCertificateVerifier.version(), Env.deployVersion(), "ecdsaCertificateVerifier version mismatch"
            );
            assertTrue(
                ecdsaCertificateVerifier.operatorTableUpdater() == Env.proxy.operatorTableUpdater(),
                "ecdsaCertificateVerifier operatorTableUpdater mismatch"
            );
        }

        /// BN254CertificateVerifier
        {
            BN254CertificateVerifier bn254CertificateVerifier = Env.proxy.bn254CertificateVerifier();
            assertEq(
                bn254CertificateVerifier.version(), Env.deployVersion(), "bn254CertificateVerifier version mismatch"
            );
            assertTrue(
                bn254CertificateVerifier.operatorTableUpdater() == Env.proxy.operatorTableUpdater(),
                "bn254CertificateVerifier operatorTableUpdater mismatch"
            );
        }

        /// TaskMailbox
        {
            TaskMailbox taskMailbox = Env.proxy.taskMailbox();
            assertEq(taskMailbox.version(), Env.deployVersion(), "taskMailbox version mismatch");
            assertTrue(
                taskMailbox.BN254_CERTIFICATE_VERIFIER() == address(Env.proxy.bn254CertificateVerifier()),
                "taskMailbox.BN254_CERTIFICATE_VERIFIER mismatch"
            );
            assertTrue(
                taskMailbox.ECDSA_CERTIFICATE_VERIFIER() == address(Env.proxy.ecdsaCertificateVerifier()),
                "taskMailbox.ECDSA_CERTIFICATE_VERIFIER mismatch"
            );
            assertEq(taskMailbox.MAX_TASK_SLA(), Env.MAX_TASK_SLA(), "taskMailbox.MAX_TASK_SLA mismatch");
        }
    }

    function _validateProxiesInitialized() internal {
        bytes memory errInit = "Initializable: contract is already initialized";

        /// OperatorTableUpdater - dummy parameters
        OperatorTableUpdater operatorTableUpdater = Env.proxy.operatorTableUpdater();
        OperatorSet memory dummyOperatorSet = OperatorSet({avs: address(0), id: 0});
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo memory dummyBN254Info;

        vm.expectRevert(errInit);
        operatorTableUpdater.initialize(
            address(0), // owner
            0, // initial paused status
            dummyOperatorSet, // generator
            0, // globalRootConfirmationThreshold
            dummyBN254Info // generatorInfo
        );

        /// TaskMailbox
        TaskMailbox taskMailbox = Env.proxy.taskMailbox();

        vm.expectRevert(errInit);
        taskMailbox.initialize(
            address(0), // owner
            0, // feeSplit
            address(0) // feeSplitCollector
        );

        // ECDSACertificateVerifier and BN254CertificateVerifier don't have initialize functions
    }

    function _assertTrue(bool b, string memory err) private pure {
        assertTrue(b, err);
    }

    function _assertFalse(bool b, string memory err) private pure {
        assertFalse(b, err);
    }

    // In order to not clutter the Zeus Env, we define our operator table updater init params here
    function _getTableUpdaterInitParams() internal view returns (OperatorTableUpdaterInitParams memory) {
        OperatorTableUpdaterInitParams memory initParams;

        if (Env._strEq(Env.env(), "preprod")) {
            initParams = _parseToml("script/releases/v1.7.0-v1.8.0-multichain-hourglass-combined/configs/preprod.toml");
        }
        // NOTE: For testnet-holesky and testnet-hoodi, the operator table updater is not used
        else if (Env._strEq(Env.env(), "testnet-sepolia") || Env._strEq(Env.env(), "testnet-base-sepolia")) {
            initParams = _parseToml("script/releases/v1.7.0-v1.8.0-multichain-hourglass-combined/configs/testnet.toml");
        } else if (Env._strEq(Env.env(), "mainnet") || Env._strEq(Env.env(), "base")) {
            initParams = _parseToml("script/releases/v1.7.0-v1.8.0-multichain-hourglass-combined/configs/mainnet.toml");
        }

        return initParams;
    }

    function _getGeneratorAddress() internal view returns (address generatorAddress) {
        if (Env._strEq(Env.env(), "preprod")) {
            generatorAddress = 0x6d609cD2812bDA02a75dcABa7DaafE4B20Ff5608;
        } else if (Env._strEq(Env.env(), "testnet-sepolia") || Env._strEq(Env.env(), "testnet-base-sepolia")) {
            generatorAddress = 0xb094Ba769b4976Dc37fC689A76675f31bc4923b0;
        } else if (Env._strEq(Env.env(), "mainnet") || Env._strEq(Env.env(), "base")) {
            generatorAddress = 0xBE1685C81aA44FF9FB319dD389addd9374383e90;
        }
        require(generatorAddress != address(0), "Invalid network");
        return generatorAddress;
    }

    function _parseToml(
        string memory path
    ) internal view returns (OperatorTableUpdaterInitParams memory initParams) {
        // Read the TOML file
        string memory root = vm.projectRoot();
        string memory fullPath = string.concat(root, "/", path);
        string memory toml = vm.readFile(fullPath);

        // Parse globalRootConfirmationThreshold
        initParams.globalRootConfirmationThreshold = uint16(toml.readUint(".globalRootConfirmationThreshold"));

        // Parse generator
        address avs = toml.readAddress(".generator.avs");
        uint32 id = uint32(toml.readUint(".generator.id"));
        initParams.generator = OperatorSet({avs: avs, id: id});

        // Parse generatorInfo
        initParams.generatorInfo.numOperators = uint256(toml.readUint(".generatorInfo.numOperators"));
        initParams.generatorInfo.operatorInfoTreeRoot = toml.readBytes32(".generatorInfo.operatorInfoTreeRoot");
        initParams.generatorInfo.totalWeights = toml.readUintArray(".generatorInfo.totalWeights");
        uint256 apkX = toml.readUint(".generatorInfo.aggregatePubkey.X");
        uint256 apkY = toml.readUint(".generatorInfo.aggregatePubkey.Y");
        initParams.generatorInfo.aggregatePubkey = BN254.G1Point({X: apkX, Y: apkY});

        return initParams;
    }

    struct OperatorTableUpdaterInitParams {
        uint16 globalRootConfirmationThreshold;
        OperatorSet generator;
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo generatorInfo;
    }

    // Define TaskMailbox initialization parameters
    function _getTaskMailboxInitParams() internal view returns (TaskMailboxInitParams memory) {
        TaskMailboxInitParams memory initParams;

        initParams.owner = Env.opsMultisig();
        initParams.feeSplit = 0; // 0% fee split initially
        initParams.feeSplitCollector = Env.opsMultisig(); // Initially set to opsMultisig

        return initParams;
    }

    struct TaskMailboxInitParams {
        address owner;
        uint16 feeSplit;
        address feeSplitCollector;
    }
}
