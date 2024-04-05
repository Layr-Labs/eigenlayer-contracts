// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../../src/contracts/interfaces/IETHPOSDeposit.sol";

import "../../../src/contracts/core/StrategyManager.sol";
import "../../../src/contracts/core/Slasher.sol";
import "../../../src/contracts/core/DelegationManager.sol";
import "../../../src/contracts/core/AVSDirectory.sol";

import "../../../src/contracts/pods/EigenPod.sol";
import "../../../src/contracts/pods/EigenPodManager.sol";
import "../../../src/contracts/pods/DelayedWithdrawalRouter.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "../../../src/test/mocks/EmptyContract.sol";
import "forge-std/Script.sol";
import "forge-std/Test.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/upgrade/GoerliUpgrade2.s.sol:GoerliUpgrade2 --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv

// NOTE: ONLY WORKS ON GOERLI
// CommitHash: 6de01c6c16d6df44af15f0b06809dc160eac0ebf
contract GoerliUpgrade2 is Script, Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    string public deploymentOutputPath = string(bytes("script/output/M1_deployment_goerli_2023_3_23.json"));

    IDelayedWithdrawalRouter delayedWithdrawalRouter;
    IDelegationManager delegation;
    IEigenPodManager eigenPodManager;
    IStrategyManager strategyManager;
    ISlasher slasher;
    IBeacon eigenPodBeacon;
    EmptyContract emptyContract;
    ProxyAdmin eigenLayerProxyAdmin;

    function run() external {
        // read and log the chainID
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        string memory config_data = vm.readFile(deploymentOutputPath);

        delayedWithdrawalRouter = IDelayedWithdrawalRouter(stdJson.readAddress(config_data, ".addresses.delayedWithdrawalRouter"));
        delegation = IDelegationManager(stdJson.readAddress(config_data, ".addresses.delegation"));
        eigenPodManager = IEigenPodManager(stdJson.readAddress(config_data, ".addresses.eigenPodManager"));
        strategyManager = IStrategyManager(stdJson.readAddress(config_data, ".addresses.strategyManager"));
        slasher = ISlasher(stdJson.readAddress(config_data, ".addresses.slasher"));
        eigenPodBeacon = IBeacon(stdJson.readAddress(config_data, ".addresses.eigenPodBeacon"));
        emptyContract = EmptyContract(stdJson.readAddress(config_data, ".addresses.emptyContract"));
        eigenLayerProxyAdmin = ProxyAdmin(stdJson.readAddress(config_data, ".addresses.eigenLayerProxyAdmin"));

        vm.startBroadcast();

        address delegationImplementation = address(
            new DelegationManager(
                strategyManager,
                slasher,
                eigenPodManager
            )
        );

        address slasherImplementation = address(
            new Slasher(
                strategyManager,
                delegation
            )
        );

        address strategyManagerImplementation = address(
            new StrategyManager(
                delegation,
                eigenPodManager,
                slasher
            )
        );

        address delayedWithdrawalRouterImplementation = address(
            new DelayedWithdrawalRouter(
                eigenPodManager
            )
        );  
        
        address eigenPodImplementation = address(
            new EigenPod(
                IETHPOSDeposit(0xff50ed3d0ec03aC01D4C79aAd74928BFF48a7b2b),
                delayedWithdrawalRouter,
                eigenPodManager,
                32e9,
                1616508000
            )
        );

        address eigenPodManagerImplementation = address(
            new EigenPodManager(
                IETHPOSDeposit(0xff50ed3d0ec03aC01D4C79aAd74928BFF48a7b2b),
                eigenPodBeacon,
                strategyManager,
                slasher,
                delegation
            )
        );

        vm.stopBroadcast();

        emit log_named_address("DelegationImplementation", delegationImplementation);
        emit log_named_address("SlasherImplementation", slasherImplementation);
        emit log_named_address("StrategyManagerImplementation", strategyManagerImplementation);
        emit log_named_address("DelayedWithdrawalRouterImplementation", delayedWithdrawalRouterImplementation);
        emit log_named_address("EigenPodImplementation", eigenPodImplementation);
        emit log_named_address("EigenPodManagerImplementation", eigenPodManagerImplementation);

        /*
        == Logs ==
        You are deploying on ChainID: 5
        DelegationImplementation: 0x56652542926444Ebce46Fd97aFd80824ed51e58C
        SlasherImplementation: 0x89C5e6e98f79be658e830Ec66b61ED3EE910D262
        StrategyManagerImplementation: 0x506C21f43e81D9d231d8A13831b42A2a2B5540E4
        DelayedWithdrawalRouterImplementation: 0xE576731194EC3d8Ba92E7c2B578ea74238772878
        EigenPodImplementation: 0x16a0d8aD2A2b12f3f47d0e8F5929F9840e29a426
        EigenPodManagerImplementation: 0xDA9B60D3dC7adD40C0e35c628561Ff71C13a189f
        */
    }
}
