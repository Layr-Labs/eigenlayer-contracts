// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "../../src/contracts/interfaces/IETHPOSDeposit.sol";
import "../../src/contracts/interfaces/IBeaconChainOracle.sol";

import "../../src/contracts/core/StrategyManager.sol";
import "../../src/contracts/core/Slasher.sol";
import "../../src/contracts/core/DelegationManager.sol";

import "../../src/contracts/pods/EigenPod.sol";
import "../../src/contracts/pods/EigenPodManager.sol";
import "../../src/contracts/pods/DelayedWithdrawalRouter.sol";

import "../../src/contracts/permissions/PauserRegistry.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

interface IDelegationManagerV0 {
    function DOMAIN_SEPARATOR() external view returns (bytes32);
}


// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/milestone/M2Deploy.s.sol:M2Deploy --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv
contract M2Deploy is Script, Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    string public m1DeploymentOutputPath = string(bytes("script/output/M1_deployment_goerli_2023_3_23.json"));

    IETHPOSDeposit public ethPOS;

    ISlasher public slasher;
    IDelegationManager public delegation;
    DelegationManager public delegationImplementation;
    IStrategyManager public strategyManager;
    StrategyManager public strategyManagerImplementation;
    IEigenPodManager public eigenPodManager;
    EigenPodManager public eigenPodManagerImplementation;
    IDelayedWithdrawalRouter public delayedWithdrawalRouter;
    IBeacon public eigenPodBeacon;
    EigenPod public eigenPodImplementation;

    ProxyAdmin public eigenLayerProxyAdmin;
    address public strategyWhitelister;
    uint256 public withdrawalDelayBlocks;
    bytes32 public delegationManagerDomainSeparator;
    uint256 public numPods;
    uint256 public maxPods;

    function run() external {
        // read and log the chainID
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        if(chainId == 1) {
            m1DeploymentOutputPath = string(bytes("script/output/M1_deployment_mainnet_2023_6_9.json"));
        }

        // READ JSON DEPLOYMENT DATA
        string memory deployment_data = vm.readFile(m1DeploymentOutputPath);
        slasher = Slasher(stdJson.readAddress(deployment_data, ".addresses.slasher"));
        delegation = slasher.delegation();
        strategyManager = slasher.strategyManager();
        eigenPodManager = strategyManager.eigenPodManager();
        eigenPodBeacon = eigenPodManager.eigenPodBeacon();
        ethPOS = eigenPodManager.ethPOS();

        eigenLayerProxyAdmin = ProxyAdmin(stdJson.readAddress(deployment_data, ".addresses.eigenLayerProxyAdmin"));

        // store pre-upgrade values to check against later
        strategyWhitelister = strategyManager.strategyWhitelister();
        // withdrawalDelayBlocks = strategyManager.withdrawalDelayBlocks();
        delegationManagerDomainSeparator = IDelegationManagerV0(address(delegation)).DOMAIN_SEPARATOR();
        numPods = eigenPodManager.numPods();
        maxPods = eigenPodManager.maxPods();
        delayedWithdrawalRouter = EigenPod(payable(eigenPodBeacon.implementation())).delayedWithdrawalRouter();

        vm.startBroadcast();
        delegationImplementation = new DelegationManager(strategyManager, slasher, eigenPodManager);
        strategyManagerImplementation = new StrategyManager(delegation, eigenPodManager, slasher);
        eigenPodManagerImplementation = new EigenPodManager(
            ethPOS,
            eigenPodBeacon,
            strategyManager,
            slasher,
            delegation
        );
        eigenPodImplementation = new EigenPod({
            _ethPOS: ethPOS,
            _delayedWithdrawalRouter: delayedWithdrawalRouter,
            _eigenPodManager: eigenPodManager,
            _MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR: 31 gwei, 
            _RESTAKED_BALANCE_OFFSET_GWEI: 0.5 gwei,
            _GENESIS_TIME: 1616508000
        });

        // write the output to a contract
        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(deployed_addresses, "slasher", address(slasher));
        vm.serializeAddress(deployed_addresses, "delegation", address(delegation));
        vm.serializeAddress(deployed_addresses, "strategyManager", address(strategyManager));
        vm.serializeAddress(deployed_addresses, "delayedWithdrawalRouter", address(delayedWithdrawalRouter));
        vm.serializeAddress(deployed_addresses, "eigenPodManager", address(eigenPodManager));
        vm.serializeAddress(deployed_addresses, "eigenPodBeacon", address(eigenPodBeacon));
        vm.serializeAddress(deployed_addresses, "ethPOS", address(ethPOS));

        vm.serializeAddress(deployed_addresses, "delegationImplementation", address(delegationImplementation));
        vm.serializeAddress(deployed_addresses, "strategyManagerImplementation", address(strategyManagerImplementation));
        vm.serializeAddress(deployed_addresses, "eigenPodManagerImplementation", address(eigenPodManagerImplementation));
        string memory deployed_addresses_output = vm.serializeAddress(deployed_addresses, "eigenPodImplementation", address(eigenPodImplementation));

        string memory chain_info = "chainInfo";
        vm.serializeUint(chain_info, "deploymentBlock", block.number);
        string memory chain_info_output = vm.serializeUint(chain_info, "chainId", chainId);

        vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);
        // serialize all the data
        string memory finalJson = vm.serializeString(parent_object, chain_info, chain_info_output);
        vm.writeJson(finalJson, "script/output/M2_deployment_data.json");

        vm.stopBroadcast();

        // perform automated testing
        simulatePerformingUpgrade();
        checkUpgradeCorrectness();
    }

    function simulatePerformingUpgrade() public {
        cheats.startPrank(eigenLayerProxyAdmin.owner());
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(delegation))),
            address(delegationImplementation)
        );
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(strategyManager))),
            address(strategyManagerImplementation)
        );
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation)
        );
        cheats.stopPrank();
        cheats.prank(UpgradeableBeacon(address(eigenPodBeacon)).owner());
        UpgradeableBeacon(address(eigenPodBeacon)).upgradeTo(address(eigenPodImplementation));
        // cheats.prank(Ownable(address(eigenPodManager)).owner());
        // eigenPodManager.updateBeaconChainOracle(IBeaconChainOracle(beaconChainOracleGoerli));
    }

    // Call contracts to ensure that all simple view functions return the same values (e.g. the return value of `StrategyManager.delegation()` hasn’t changed)
    // StrategyManager: delegation, eigenPodManager, slasher, strategyWhitelister, withdrawalDelayBlocks all unchanged
    // DelegationManager: DOMAIN_SEPARATOR, strategyManager, slasher  all unchanged
    // EigenPodManager: ethPOS, eigenPodBeacon, strategyManager, slasher, beaconChainOracle, numPods, maxPods  all unchanged
    // delegationManager is now correct (added immutable)
    // Call contracts to make sure they are still “initialized” (ensure that trying to call initializer reverts)
    function checkUpgradeCorrectness() public {
        require(strategyManager.delegation() == delegation, "strategyManager.delegation incorrect");
        require(strategyManager.eigenPodManager() == eigenPodManager, "strategyManager.eigenPodManager incorrect");
        require(strategyManager.slasher() == slasher, "strategyManager.slasher incorrect");
        require(strategyManager.strategyWhitelister() == strategyWhitelister, "strategyManager.strategyWhitelister incorrect");
        // require(strategyManager.withdrawalDelayBlocks() == withdrawalDelayBlocks, "strategyManager.withdrawalDelayBlocks incorrect");

        require(delegation.domainSeparator() == delegationManagerDomainSeparator, "delegation.domainSeparator incorrect");
        require(delegation.slasher() == slasher, "delegation.slasher incorrect");
        require(delegation.eigenPodManager() == eigenPodManager, "delegation.eigenPodManager incorrect");

        require(eigenPodManager.ethPOS() == ethPOS, "eigenPodManager.ethPOS incorrect");
        require(eigenPodManager.eigenPodBeacon() == eigenPodBeacon, "eigenPodManager.eigenPodBeacon incorrect");
        require(eigenPodManager.strategyManager() == strategyManager, "eigenPodManager.strategyManager incorrect");
        require(eigenPodManager.slasher() == slasher, "eigenPodManager.slasher incorrect");
        // require(address(eigenPodManager.beaconChainOracle()) == beaconChainOracleGoerli, "eigenPodManager.beaconChainOracle incorrect");
        require(eigenPodManager.numPods() == numPods, "eigenPodManager.numPods incorrect");
        require(eigenPodManager.maxPods() == maxPods, "eigenPodManager.maxPods incorrect");
        require(eigenPodManager.delegationManager() == delegation, "eigenPodManager.delegationManager incorrect");

        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        StrategyManager(address(strategyManager)).initialize(address(this), address(this), PauserRegistry(address(this)), 0);

        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        DelegationManager(address(delegation)).initialize(address(this), PauserRegistry(address(this)), 0);

        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        EigenPodManager(address(eigenPodManager)).initialize(0, IBeaconChainOracle(address(this)), address(this), PauserRegistry(address(this)), 0);
    }

// Existing LST depositor – ensure that strategy length and shares are all identical
// Existing LST depositor – ensure that an existing queued withdrawal remains queued
// Check from stored root, and recalculate root and make sure it matches
// Check that completing the withdrawal results in the same behavior (same transfer of ERC20 tokens)
// Check that staker nonce & numWithdrawalsQueued remains the same as before the upgrade
// Existing LST depositor – queuing a withdrawal before/after the upgrade has the same effects (same decrease in shares, resultant withdrawal root)
// Existing EigenPod owner – EigenPodManager.ownerToPod remains the same
// Existing EigenPod owner –  EigenPodManager.hasPod remains the same
// Existing EigenPod owner –  EigenPod.podOwner remains the same
// Existing EigenPod owner –  EigenPod.mostRecentWithdrawalTimestamp (after upgrade) == EigenPod.mostRecentWithdrawalBlock (before upgrade)
// Existing EigenPod owner – EigenPod.hasRestaked remains false
// Can call EigenPod.activateRestaking and it correctly:
// Sends all funds in EigenPod (need to make sure it has nonzero balance beforehand)
// Sets `hasRestaked` to ‘true’
// Emits a ‘RestakingActivated’ event
// EigenPod.mostRecentWithdrawalTimestamp updates correctly
// EigenPod: ethPOS, delayedWithdrawalRouter, eigenPodManager

}
