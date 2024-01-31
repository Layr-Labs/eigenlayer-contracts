// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../src/contracts/interfaces/IETHPOSDeposit.sol";

import "../../src/contracts/core/StrategyManager.sol";
import "../../src/contracts/core/Slasher.sol";
import "../../src/contracts/core/DelegationManager.sol";
import "../../src/contracts/core/AVSDirectory.sol";

import "../../src/contracts/pods/EigenPod.sol";
import "../../src/contracts/pods/EigenPodManager.sol";
import "../../src/contracts/pods/DelayedWithdrawalRouter.sol";

import "forge-std/Script.sol";
import "forge-std/Test.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/upgrade/GoerliUpgrade2.s.sol:GoerliUpgrade2 --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv

// NOTE: ONLY WORKS ON GOERLI
// CommitHash: 7257364d03d255ea8c855f36317ce0e892b78497
contract GoerliUpgrade2 is Script, Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    string public deploymentOutputPath = string(bytes("script/output/M1_deployment_goerli_2023_3_23.json"));

    function run() external {
        // read and log the chainID
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        string memory config_data = vm.readFile(deploymentOutputPath);

        IDelayedWithdrawalRouter delayedWithdrawalRouter = IDelayedWithdrawalRouter(stdJson.readAddress(config_data, ".addresses.delayedWithdrawalRouter"));
        IDelegationManager delegation = IDelegationManager(stdJson.readAddress(config_data, ".addresses.delegation"));
        IEigenPodManager eigenPodManager = IEigenPodManager(stdJson.readAddress(config_data, ".addresses.eigenPodManager"));
        IStrategyManager strategyManager = IStrategyManager(stdJson.readAddress(config_data, ".addresses.strategyManager"));
        ISlasher slasher = ISlasher(stdJson.readAddress(config_data, ".addresses.slasher"));
        IBeacon eigenPodBeacon = IBeacon(stdJson.readAddress(config_data, ".addresses.eigenPodBeacon"));

        vm.startBroadcast();

        address avsDirectoryImplementation = address(
            new AVSDirectory(
                delegation
            )
        );

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

        emit log_named_address("AVSDirectoryImplementation", avsDirectoryImplementation);
        //   AVSDirectoryImplementation:
        emit log_named_address("DelegationImplementation", delegationImplementation);
        //   DelegationImplementation: 
        emit log_named_address("SlasherImplementation", slasherImplementation);
        //   SlasherImplementation: 
        emit log_named_address("StrategyManagerImplementation", strategyManagerImplementation);
        //   StrategyManagerImplementation: 
        emit log_named_address("DelayedWithdrawalRouterImplementation", delayedWithdrawalRouterImplementation);
        //   DelayedWithdrawalRouterImplementation: 
        emit log_named_address("EigenPodImplementation", eigenPodImplementation);
        //   EigenPodImplementation: 
        emit log_named_address("EigenPodManagerImplementation", eigenPodManagerImplementation);
        //   EigenPodManagerImplementation: 
    }
}