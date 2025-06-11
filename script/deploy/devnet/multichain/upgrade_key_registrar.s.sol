// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// OpenZeppelin Contracts
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

// Core Contracts
import "src/contracts/permissions/KeyRegistrar.sol";
import "src/contracts/interfaces/IPermissionController.sol";
import "src/contracts/interfaces/IAllocationManager.sol";
import "src/contracts/interfaces/IKeyRegistrar.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

// Forge
import "forge-std/Script.sol";

// Run with:
// forge script script/deploy/devnet/multichain/upgrade_key_registrar.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run()"
contract UpgradeKeyRegistrar is Script {
    // ========== STATE VARIABLES ==========
    
    // Proxy Admin (controls upgrades)
    ProxyAdmin public proxyAdmin = ProxyAdmin(0xC5dc0d145a21FDAD791Df8eDC7EbCB5330A3FdB5);
    
    // Existing Proxy Address for KeyRegistrar
    // TODO: Replace with actual deployed proxy address
    KeyRegistrar public keyRegistrarProxy = KeyRegistrar(0x1C84Bb62fE7791e173014A879C706445fa893BbE);
    
    // Dependencies for KeyRegistrar constructor
    // TODO: Replace with actual deployed addresses
    IPermissionController public permissionController = IPermissionController(0xa2348c77802238Db39f0CefAa500B62D3FDD682b);
    IAllocationManager public allocationManager = IAllocationManager(0xFdD5749e11977D60850E06bF5B13221Ad95eb6B4);
    
    // Version for the upgrade
    string public constant VERSION = "1.0.1"; // Update version for the upgrade
    
    // New Implementation contract (will be deployed)
    KeyRegistrar public keyRegistrarImplementation;
    
    function run() public {
        vm.startBroadcast();
        
        // Deploy new KeyRegistrar implementation
        keyRegistrarImplementation = new KeyRegistrar(
            permissionController,
            allocationManager,
            VERSION
        );
        
        // Upgrade KeyRegistrar proxy to new implementation
        proxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(keyRegistrarProxy))),
            address(keyRegistrarImplementation)
        );
        
        vm.stopBroadcast();
    }
} 