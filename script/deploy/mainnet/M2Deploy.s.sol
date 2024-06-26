// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "../../../src/contracts/interfaces/IETHPOSDeposit.sol";

import "../../../src/contracts/core/StrategyManager.sol";
import "../../../src/contracts/core/Slasher.sol";
import "../../../src/contracts/core/DelegationManager.sol";

import "../../../src/contracts/pods/EigenPod.sol";
import "../../../src/contracts/pods/EigenPodManager.sol";

import "../../../src/contracts/permissions/PauserRegistry.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/milestone/M2Deploy.s.sol:M2Deploy --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv
contract M2Deploy is Script, Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    string public m1DeploymentOutputPath;
    string public m2DeploymentOutputPath;

    // EigenLayer core contracts
    ISlasher public slasher;
    IDelegationManager public delegation;
    DelegationManager public delegationImplementation;
    IStrategyManager public strategyManager;
    StrategyManager public strategyManagerImplementation;
    IEigenPodManager public eigenPodManager;
    EigenPodManager public eigenPodManagerImplementation;
    IBeacon public eigenPodBeacon;
    EigenPod public eigenPodImplementation;

    // Eigenlayer Proxy Admin
    ProxyAdmin public eigenLayerProxyAdmin;

    // BeaconChain deposit contract
    IETHPOSDeposit public ethPOS;

    // RPC url to fork from for pre-upgrade state change tests
    string public rpcUrl;

    // Pre-upgrade values to check post-upgrade
    address public strategyWhitelister;
    bytes32 public withdrawalDelayBlocksStorageSlot = bytes32(uint256(204)); // 0xcc == 204
    uint256 public withdrawalsQueuedStorageSlot = 208; //0xd0 = 208
    uint256 public withdrawalDelayBlocks;
    bytes32 public delegationManagerDomainSeparator;
    uint256 public numPods;

    // Pointers to pre-upgrade values for lstDepositor
    address public lstDepositor;
    uint256 public stakerStrategyListLength;
    uint256[] public stakerStrategyShares; // Array of shares in each strategy
    IStrategy[] public stakerStrategyList; // Array of strategies staker has deposited into
    IERC20[] public tokensToWithdraw;

    // Pointers to pre-upgrade values for eigenPodDepositor
    address public eigenPodDepositor;
    IEigenPod public eigenPod;
    address public eigenPodOwner;
    bool public hasPod;
    uint64 public mostRecentWithdrawalBlock;

    function run() external {
        // Read and log the chain ID
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        // Update deployment path addresses if on mainnet
        if (chainId == 1) {
            m1DeploymentOutputPath = string(bytes("script/output/M1_deployment_mainnet_2023_6_9.json"));
            m2DeploymentOutputPath = "script/output/M2_deployment_data_mainnet.json";
            rpcUrl = "RPC_MAINNET";
        } else if (chainId == 5) {
            m1DeploymentOutputPath = string(bytes("script/output/M1_deployment_goerli_2023_3_23.json"));
            m2DeploymentOutputPath = "script/output/M2_deployment_data_goerli.json";
            rpcUrl = "RPC_GOERLI";
        } else {
            revert("Chain not supported");
        }

        // Read json data
        string memory deployment_data = vm.readFile(m1DeploymentOutputPath);
        slasher = Slasher(stdJson.readAddress(deployment_data, ".addresses.slasher"));
        delegation = slasher.delegation();
        strategyManager = slasher.strategyManager();
        eigenPodManager = strategyManager.eigenPodManager();
        eigenPodBeacon = eigenPodManager.eigenPodBeacon();
        ethPOS = eigenPodManager.ethPOS();

        eigenLayerProxyAdmin = ProxyAdmin(stdJson.readAddress(deployment_data, ".addresses.eigenLayerProxyAdmin"));

        // Store pre-upgrade values to check against later
        strategyWhitelister = strategyManager.strategyWhitelister();
        delegationManagerDomainSeparator = IDelegationManagerV0(address(delegation)).DOMAIN_SEPARATOR();
        numPods = eigenPodManager.numPods();

        // Set chain-specific values
        IStrategy[] memory strategyArray = new IStrategy[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        if (chainId == 1) {
            // no-op for now
        } else if (chainId == 5) {
            // Set LST Depositor values
            lstDepositor = 0x01e453D2465cEC1BD2ac9aed06115Fbf28482b33;
            strategyArray[0] = IStrategy(0x879944A8cB437a5f8061361f82A6d4EED59070b5);
            shareAmounts[0] = 188647761812080108;
            tokensToWithdraw.push(IERC20(0x178E141a0E3b34152f73Ff610437A7bf9B83267A));

            // Set eigenPod owner values
            eigenPodDepositor = 0xE9D04433bac1bd584B0493cbaBa170CCCBDA8F00;
        } else {
            revert("chain ID not supported");
        }

        // Store LST depositor pre-upgrade values
        stakerStrategyListLength = strategyManager.stakerStrategyListLength(lstDepositor);
        (stakerStrategyList, stakerStrategyShares) = strategyManager.getDeposits(lstDepositor);

        // Store eigenPod owner pre-ugprade values
        eigenPod = eigenPodManager.ownerToPod(eigenPodDepositor);
        require(address(eigenPod).balance > 0, "eigenPod to test has balance of 0");
        hasPod = eigenPodManager.hasPod(eigenPodDepositor);
        eigenPodOwner = eigenPod.podOwner();
        mostRecentWithdrawalBlock = m1EigenPod(address(eigenPod)).mostRecentWithdrawalBlockNumber();

        // Begin deployment
        vm.startBroadcast();

        // Deploy new implmementation contracts
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
            _eigenPodManager: eigenPodManager,
            _GENESIS_TIME: 1616508000
        });

        vm.stopBroadcast();

        // Write json data out
        string memory parent_object = "parent object";
        string memory deployed_addresses = "addresses";

        // Serialize proxy and non-deployed addresses
        vm.serializeAddress(deployed_addresses, "slasher", address(slasher));
        vm.serializeAddress(deployed_addresses, "delegation", address(delegation));
        vm.serializeAddress(deployed_addresses, "strategyManager", address(strategyManager));
        vm.serializeAddress(deployed_addresses, "eigenPodManager", address(eigenPodManager));
        vm.serializeAddress(deployed_addresses, "eigenPodBeacon", address(eigenPodBeacon));
        vm.serializeAddress(deployed_addresses, "ethPOS", address(ethPOS));

        // Serialize new implementation addresses
        vm.serializeAddress(deployed_addresses, "delegationImplementation", address(delegationImplementation));
        vm.serializeAddress(
            deployed_addresses,
            "strategyManagerImplementation",
            address(strategyManagerImplementation)
        );
        vm.serializeAddress(
            deployed_addresses,
            "eigenPodManagerImplementation",
            address(eigenPodManagerImplementation)
        );
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "eigenPodImplementation",
            address(eigenPodImplementation)
        );

        // Add chain info
        string memory chain_info = "chainInfo";
        vm.serializeUint(chain_info, "deploymentBlock", block.number);
        string memory chain_info_output = vm.serializeUint(chain_info, "chainId", chainId);

        // Save addresses
        vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);
        string memory finalJson = vm.serializeString(parent_object, chain_info, chain_info_output);

        // Write output to file
        vm.writeJson(finalJson, m2DeploymentOutputPath);

        // Perform post-upgrade tests
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

        // Upgrade beacon
        cheats.prank(UpgradeableBeacon(address(eigenPodBeacon)).owner());
        UpgradeableBeacon(address(eigenPodBeacon)).upgradeTo(address(eigenPodImplementation));
    }

    function checkUpgradeCorrectness() public {
        _verifyStorageSlots();

        _verifyContractsInitialized();

        _verifyEigenPodCorrectness();
    }

    // Call contracts to ensure that all simple view functions return the same values (e.g. the return value of `StrategyManager.delegation()` hasn’t changed)
    // StrategyManager: delegation, eigenPodManager, slasher, strategyWhitelister, withdrawalDelayBlocks all unchanged
    // DelegationManager: DOMAIN_SEPARATOR, strategyManager, slasher, eigenPodManager  all unchanged
    // EigenPodManager: ethPOS, eigenPodBeacon, strategyManager, slasher, numPods  all unchanged
    // delegationManager is now correct (added immutable)
    // Call contracts to make sure they are still “initialized” (ensure that trying to call initializer reverts)
    function _verifyStorageSlots() internal view {
        // StrategyManager: Check view functions return pre-upgraded values
        require(strategyManager.delegation() == delegation, "strategyManager.delegation incorrect");
        require(strategyManager.eigenPodManager() == eigenPodManager, "strategyManager.eigenPodManager incorrect");
        require(strategyManager.slasher() == slasher, "strategyManager.slasher incorrect");
        require(
            strategyManager.strategyWhitelister() == strategyWhitelister,
            "strategyManager.strategyWhitelister incorrect"
        );
        require(
            cheats.load(address(strategyManager), withdrawalDelayBlocksStorageSlot) == bytes32(withdrawalDelayBlocks),
            "strategyManager.withdrawalDelayBlocks incorrect"
        );
        // DelegationManager: Check view functions return pre-upgraded values
        require(DelegationManagerStorage(address(delegation)).strategyManager() == strategyManager, "delegation.strategyManager incorrect");
        require(
            delegation.domainSeparator() == delegationManagerDomainSeparator,
            "delegation.domainSeparator incorrect"
        );
        require(DelegationManagerStorage(address(delegation)).slasher() == slasher, "delegation.slasher incorrect");
        require(DelegationManagerStorage(address(delegation)).eigenPodManager() == eigenPodManager, "delegation.eigenPodManager incorrect");
        // EigenPodManager: check view functions return pre-upgraded values
        require(eigenPodManager.ethPOS() == ethPOS, "eigenPodManager.ethPOS incorrect");
        require(eigenPodManager.eigenPodBeacon() == eigenPodBeacon, "eigenPodManager.eigenPodBeacon incorrect");
        require(eigenPodManager.strategyManager() == strategyManager, "eigenPodManager.strategyManager incorrect");
        require(eigenPodManager.slasher() == slasher, "eigenPodManager.slasher incorrect");
        require(eigenPodManager.numPods() == numPods, "eigenPodManager.numPods incorrect");
        require(EigenPodManagerStorage(address(eigenPodManager)).delegationManager() == delegation, "eigenPodManager.delegationManager incorrect");
    }

    function _verifyContractsInitialized() internal {
        // Check that contracts are unable to be re-initialized
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        StrategyManager(address(strategyManager)).initialize(
            address(this),
            address(this),
            PauserRegistry(address(this)),
            0
        );

        IStrategy[] memory strategyArray = new IStrategy[](0);
        uint256[] memory withdrawalDelayBlocksArray = new uint256[](0);
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        DelegationManager(address(delegation)).initialize(
            address(this),
            PauserRegistry(address(this)),
            0, // initialPausedStatus
            0, // minWithdrawalDelayBLocks
            strategyArray,
            withdrawalDelayBlocksArray
        );

        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        EigenPodManager(address(eigenPodManager)).initialize(
            address(this),
            PauserRegistry(address(this)),
            0
        );
    }

    function _verifyEigenPodCorrectness() public {
        // Check that state is correct
        require(
            address(eigenPodManager.ownerToPod(eigenPodDepositor)) == address(eigenPod),
            "eigenPodManager.ownerToPod incorrect"
        );
        require(eigenPodManager.hasPod(eigenPodDepositor) == hasPod, "eigenPodManager.hasPod incorrect");
        require(eigenPod.podOwner() == eigenPodOwner, "eigenPod.podOwner incorrect");

        // Unpause eigenpods verify credentials
        uint256 paused = IPausable(address(eigenPodManager)).paused();
        cheats.prank(IPauserRegistry(IPausable(address(eigenPodManager)).pauserRegistry()).unpauser());
        IPausable(address(eigenPodManager)).unpause(paused ^ (1 << 2)); // eigenpods verify credentials on 2nd bit

        cheats.prank(eigenPodOwner);
        eigenPod.startCheckpoint(false);
    }
}

interface IDelegationManagerV0 {
    function DOMAIN_SEPARATOR() external view returns (bytes32);
}

interface m1EigenPod {
    function mostRecentWithdrawalBlockNumber() external view returns (uint64);
}
