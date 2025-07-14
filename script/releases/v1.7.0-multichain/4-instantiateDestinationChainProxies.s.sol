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
        // If we're not on a destination chain, we don't need to do anything
        if (!Env.isDestinationChain()) {
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
        OperatorTableUpdaterInitParams memory initParams = _getTableUpdaterInitParams();
        ITransparentUpgradeableProxy operatorTableUpdaterProxy =
            ITransparentUpgradeableProxy(payable(address(Env.proxy.operatorTableUpdater())));
        operatorTableUpdaterProxy.upgradeToAndCall(
            address(Env.impl.operatorTableUpdater()),
            abi.encodeCall(
                OperatorTableUpdater.initialize,
                (
                    Env.opsMultisig(),
                    0, // initial paused status
                    initParams.globalRootConfirmerSet,
                    initParams.globalRootConfirmationThreshold,
                    initParams.globalRootConfirmerSetInfo
                )
            )
        );

        // Transfer proxy admin ownership
        operatorTableUpdaterProxy.changeAdmin(Env.proxyAdmin());
        ecdsaCertificateVerifierProxy.changeAdmin(Env.proxyAdmin());
        bn254CertificateVerifierProxy.changeAdmin(Env.proxyAdmin());
    }

    function testScript() public virtual override {
        if (!Env.isDestinationChain()) {
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
        OperatorTableUpdater operatorTableUpdater = Env.proxy.operatorTableUpdater();
        assertTrue(operatorTableUpdater.owner() == Env.opsMultisig(), "operatorTableUpdater.owner invalid");
        assertTrue(operatorTableUpdater.paused() == 0, "operatorTableUpdater.paused invalid");

        // Checks on the generator
        OperatorTableUpdaterInitParams memory initParams = _getTableUpdaterInitParams();
        assertEq(
            operatorTableUpdater.getGenerator().key(),
            initParams.globalRootConfirmerSet.key(),
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
        assertEq(
            operatorTableUpdater.getLatestReferenceTimestamp(),
            0,
            "operatorTableUpdater.latestReferenceTimestamp invalid"
        );
        assertTrue(
            operatorTableUpdater.isRootValid(operatorTableUpdater.GENERATOR_GLOBAL_TABLE_ROOT()),
            "operatorTableUpdater.generatorGlobalTableRoot invalid"
        );
        assertTrue(
            operatorTableUpdater.isRootValidByTimestamp(operatorTableUpdater.GENERATOR_REFERENCE_TIMESTAMP()),
            "operatorTableUpdater.generatorGlobalTableRoot invalid"
        );
        ICrossChainRegistryTypes.OperatorSetConfig memory generatorConfig = operatorTableUpdater.getGeneratorConfig();
        assertEq(generatorConfig.maxStalenessPeriod, 0, "generatorConfig.maxStalenessPeriod invalid");
        assertEq(generatorConfig.owner, address(operatorTableUpdater), "generatorConfig.owner invalid");

        // Validate ECDSACertificateVerifier
        ECDSACertificateVerifier ecdsaCertificateVerifier = Env.proxy.ecdsaCertificateVerifier();
        assertTrue(address(ecdsaCertificateVerifier) != address(0), "ecdsaCertificateVerifier not deployed");

        // Validate BN254CertificateVerifier
        BN254CertificateVerifier bn254CertificateVerifier = Env.proxy.bn254CertificateVerifier();
        assertTrue(address(bn254CertificateVerifier) != address(0), "bn254CertificateVerifier not deployed");
        assertEq(
            bn254CertificateVerifier.latestReferenceTimestamp(initParams.globalRootConfirmerSet),
            1,
            "bn254CertificateVerifier.latestReferenceTimestamp invalid"
        );
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
    }

    function _validateProxyConstructors() internal view {
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

        ECDSACertificateVerifier ecdsaCertificateVerifier = Env.proxy.ecdsaCertificateVerifier();
        assertEq(ecdsaCertificateVerifier.version(), Env.deployVersion(), "ecdsaCertificateVerifier version mismatch");
        assertTrue(
            ecdsaCertificateVerifier.operatorTableUpdater() == Env.proxy.operatorTableUpdater(),
            "ecdsaCertificateVerifier operatorTableUpdater mismatch"
        );

        BN254CertificateVerifier bn254CertificateVerifier = Env.proxy.bn254CertificateVerifier();
        assertEq(bn254CertificateVerifier.version(), Env.deployVersion(), "bn254CertificateVerifier version mismatch");
        assertTrue(
            bn254CertificateVerifier.operatorTableUpdater() == Env.proxy.operatorTableUpdater(),
            "bn254CertificateVerifier operatorTableUpdater mismatch"
        );
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
            dummyOperatorSet, // globalRootConfirmerSet
            0, // globalRootConfirmationThreshold
            dummyBN254Info // globalRootConfirmerSetInfo
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
            initParams = _parseToml("script/releases/v1.7.0-multichain/configs/preprod.toml");
        }
        // NOTE: For testnet-holesky and testnet-hoodi, the operator table updater is not used
        else if (Env._strEq(Env.env(), "testnet-sepolia") || Env._strEq(Env.env(), "testnet-base-sepolia")) {
            initParams = _parseToml("script/releases/v1.7.0-multichain/configs/testnet.toml");
        } else if (Env._strEq(Env.env(), "mainnet") || Env._strEq(Env.env(), "mainnet-base")) {
            initParams = _parseToml("script/releases/v1.7.0-multichain/configs/mainnet.toml");
        }

        return initParams;
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

        // Parse globalRootConfirmerSet
        address avs = toml.readAddress(".globalRootConfirmerSet.avs");
        uint32 id = uint32(toml.readUint(".globalRootConfirmerSet.id"));
        initParams.globalRootConfirmerSet = OperatorSet({avs: avs, id: id});

        // Parse globalRootConfirmerSetInfo
        initParams.globalRootConfirmerSetInfo.numOperators =
            uint256(toml.readUint(".globalRootConfirmerSetInfo.numOperators"));
        initParams.globalRootConfirmerSetInfo.operatorInfoTreeRoot =
            toml.readBytes32(".globalRootConfirmerSetInfo.operatorInfoTreeRoot");
        initParams.globalRootConfirmerSetInfo.totalWeights =
            toml.readUintArray(".globalRootConfirmerSetInfo.totalWeights");
        uint256 apkX = toml.readUint(".globalRootConfirmerSetInfo.aggregatePubkey.X");
        uint256 apkY = toml.readUint(".globalRootConfirmerSetInfo.aggregatePubkey.Y");
        initParams.globalRootConfirmerSetInfo.aggregatePubkey = BN254.G1Point({X: apkX, Y: apkY});

        return initParams;
    }

    struct OperatorTableUpdaterInitParams {
        uint16 globalRootConfirmationThreshold;
        OperatorSet globalRootConfirmerSet;
        IOperatorTableCalculatorTypes.BN254OperatorSetInfo globalRootConfirmerSetInfo;
    }
}
