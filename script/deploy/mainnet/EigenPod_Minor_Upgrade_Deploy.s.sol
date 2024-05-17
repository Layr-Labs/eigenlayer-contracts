// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";

import "../../../src/contracts/interfaces/IETHPOSDeposit.sol";
import "../../../src/contracts/interfaces/IBeaconChainOracle.sol";

import "../../../src/contracts/core/StrategyManager.sol";
import "../../../src/contracts/core/Slasher.sol";
import "../../../src/contracts/core/DelegationManager.sol";

import "../../../src/contracts/pods/EigenPod.sol";
import "../../../src/contracts/pods/EigenPodManager.sol";
import "../../../src/contracts/pods/DelayedWithdrawalRouter.sol";

import "../../../src/contracts/permissions/PauserRegistry.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/deploy/mainnet/EigenPod_Minor_Upgrade_Deploy.s.sol:EigenPod_Minor_Upgrade_Deploy --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv
contract EigenPod_Minor_Upgrade_Deploy is Script, Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    string public m2DeploymentOutputPath;
    string public freshOutputPath;

    // EigenLayer core contracts
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

    // Eigenlayer Proxy Admin
    ProxyAdmin public eigenLayerProxyAdmin;

    // BeaconChain deposit contract & beacon chain oracle
    IETHPOSDeposit public ethPOS;
    address public beaconChainOracle;

    // RPC url to fork from for pre-upgrade state change tests
    string public rpcUrl;

    uint64 public genesisTimeBefore;
    uint64 public maxRestakedBalanceBefore;

    function run() external {
        // Read and log the chain ID
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        // Update deployment path addresses if on mainnet
        if (chainId == 1) {
            m2DeploymentOutputPath = "script/output/mainnet/M2_mainnet_upgrade.output.json";
            freshOutputPath = "script/output/mainnet/eigenpod_minor_upgrade_deploy.json";
            rpcUrl = "RPC_MAINNET";
        } else {
            revert("Chain not supported");
        }

        // Read json data
        string memory deployment_data = vm.readFile(m2DeploymentOutputPath);
        slasher = Slasher(stdJson.readAddress(deployment_data, ".addresses.slasher"));
        delegation = DelegationManager(stdJson.readAddress(deployment_data, ".addresses.delegationManager"));
        strategyManager = DelegationManager(address(delegation)).strategyManager();
        eigenPodManager = strategyManager.eigenPodManager();
        eigenPodBeacon = eigenPodManager.eigenPodBeacon();
        ethPOS = eigenPodManager.ethPOS();

        eigenLayerProxyAdmin = ProxyAdmin(stdJson.readAddress(deployment_data, ".addresses.eigenLayerProxyAdmin"));

        genesisTimeBefore = EigenPod(payable(eigenPodBeacon.implementation())).GENESIS_TIME();
        maxRestakedBalanceBefore = EigenPod(payable(eigenPodBeacon.implementation())).MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR();
        delayedWithdrawalRouter = EigenPod(payable(eigenPodBeacon.implementation())).delayedWithdrawalRouter();

        // Begin deployment
        vm.startBroadcast();

        // Deploy new implmementation contracts
        eigenPodImplementation = new EigenPod({
            _ethPOS: ethPOS,
            _delayedWithdrawalRouter: delayedWithdrawalRouter,
            _eigenPodManager: eigenPodManager,
            _MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR: maxRestakedBalanceBefore,
            _GENESIS_TIME: genesisTimeBefore
        });

        vm.stopBroadcast();

        // Write json data out
        string memory parent_object = "parent object";
        string memory deployed_addresses = "addresses";

        // Add chain info
        string memory chain_info = "chainInfo";
        vm.serializeUint(chain_info, "deploymentBlock", block.number);
        string memory chain_info_output = vm.serializeUint(chain_info, "chainId", chainId);

        // Serialize new implementation addresses
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "eigenPodImplementation",
            address(eigenPodImplementation)
        );

        // Save addresses
        vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);
        string memory finalJson = vm.serializeString(parent_object, chain_info, chain_info_output);

        // Write output to file
        vm.writeJson(finalJson, freshOutputPath);

        // Perform post-upgrade tests
        simulatePerformingUpgrade();
        checkUpgradeCorrectness();
    }

    function simulatePerformingUpgrade() public {
        // Upgrade beacon
        cheats.prank(UpgradeableBeacon(address(eigenPodBeacon)).owner());
        UpgradeableBeacon(address(eigenPodBeacon)).upgradeTo(address(eigenPodImplementation));
    }

    function checkUpgradeCorrectness() public view {
        _verifyEigenPodCorrectness();
    }

    function _verifyEigenPodCorrectness() public view {
        // Check that state is correct
        require(eigenPodBeacon.implementation() == address(eigenPodImplementation),
            "implementation set incorrectly");
        require(eigenPodImplementation.delayedWithdrawalRouter() == delayedWithdrawalRouter,
            "delayedWithdrawalRouter set incorrectly");
        require(eigenPodImplementation.ethPOS() == ethPOS,
            "ethPOS set incorrectly");
        require(eigenPodImplementation.eigenPodManager() == eigenPodManager,
            "eigenPodManager set incorrectly");
        // check that values are unchanged
        require(eigenPodImplementation.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR() == maxRestakedBalanceBefore,
            "MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR set incorrectly");
        require(eigenPodImplementation.GENESIS_TIME() == genesisTimeBefore,
            "GENESIS_TIME set incorrectly");
        // redundant checks on correct values
        require(eigenPodImplementation.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR() == 32 gwei,
            "MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR set incorrectly");
        require(eigenPodImplementation.GENESIS_TIME() == 1606824023,
            "GENESIS_TIME set incorrectly");


        require(address(EigenPod(payable(eigenPodBeacon.implementation())).delayedWithdrawalRouter())  == 0x7Fe7E9CC0F274d2435AD5d56D5fa73E47F6A23D8);
        require(address(EigenPod(payable(eigenPodBeacon.implementation())).eigenPodManager())  == 0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338);
        require(address(EigenPod(payable(eigenPodBeacon.implementation())).ethPOS())  == 0x00000000219ab540356cBB839Cbe05303d7705Fa);
    }
}