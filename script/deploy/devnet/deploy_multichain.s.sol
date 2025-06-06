// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

// Core Contracts
import "../../../src/contracts/core/AllocationManager.sol";
import "../../../src/contracts/permissions/PermissionController.sol";

// Multichain Contracts
import "../../../src/contracts/multichain/BN254CertificateVerifier.sol";
import "../../../src/contracts/multichain/OperatorTableUpdater.sol";
import "../../../src/contracts/multichain/BN254TableCalculator.sol";
import "../../../src/contracts/permissions/KeyRegistrar.sol";
import "../../../src/contracts/interfaces/IECDSACertificateVerifier.sol";
import "../../../src/contracts/permissions/PauserRegistry.sol";
import "../../../src/test/mocks/EmptyContract.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

contract DeployMultichain is Script, Test {
    Vm cheats = Vm(VM_ADDRESS);

    // EigenLayer Contracts
    EmptyContract public emptyContract;
    AllocationManager public allocationManager = AllocationManager(0x948a420b8CC1d6BFd0B6087C2E7c344a2CD0bc39);
    PermissionController public permissionController = PermissionController(0x25E5F8B1E7aDf44518d35D5B2271f114e081f0E5);

    // Multichain Contracts
    ProxyAdmin public proxyAdmin;
    KeyRegistrar public keyRegistrar;
    KeyRegistrar public keyRegistrarImplementation;
    BN254CertificateVerifier public bn254CertificateVerifier;
    BN254CertificateVerifier public bn254CertificateVerifierImplementation;
    OperatorTableUpdater public operatorTableUpdater;
    OperatorTableUpdater public operatorTableUpdaterImplementation;
    BN254TableCalculator public bn254TableCalculator;

    function run() public {
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        emptyContract = new EmptyContract();
        proxyAdmin = new ProxyAdmin();

        // Key Registrar
        keyRegistrar = KeyRegistrar(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));

        // BN254 Certificate Verifier
        bn254CertificateVerifier = BN254CertificateVerifier(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));

        // Operator Table Updater
        operatorTableUpdater = OperatorTableUpdater(address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")));

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        keyRegistrarImplementation = new KeyRegistrar(permissionController, allocationManager, "9.9.9");
        bn254CertificateVerifierImplementation = new BN254CertificateVerifier(address(operatorTableUpdater));
        operatorTableUpdaterImplementation = new OperatorTableUpdater(address(bn254CertificateVerifier), address(keyRegistrar), "9.9.9");

        // Third, upgrade the proxies to point to the new implementations
        proxyAdmin.upgrade(ITransparentUpgradeableProxy(payable(address(keyRegistrar))), address(keyRegistrarImplementation));
        proxyAdmin.upgrade(ITransparentUpgradeableProxy(payable(address(bn254CertificateVerifier))), address(bn254CertificateVerifierImplementation));
        proxyAdmin.upgrade(ITransparentUpgradeableProxy(payable(address(operatorTableUpdater))), address(operatorTableUpdaterImplementation));

    }
}