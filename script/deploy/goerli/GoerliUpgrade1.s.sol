// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";


import "../../../src/contracts/interfaces/IETHPOSDeposit.sol";
import "../../../src/contracts/interfaces/IBeaconChainOracle.sol";

import "../../../src/contracts/core/StrategyManager.sol";
import "../../../src/contracts/core/Slasher.sol";
import "../../../src/contracts/core/DelegationManager.sol";

import "../../../src/contracts/strategies/StrategyBase.sol";

import "../../../src/contracts/pods/EigenPod.sol";
import "../../../src/contracts/pods/EigenPodManager.sol";
import "../../../src/contracts/pods/DelayedWithdrawalRouter.sol";


import "forge-std/Script.sol";
import "forge-std/Test.sol";

// # To load the variables in the .env file
// source .env

// # To deploy and verify our contract
// forge script script/misc/DeployStrategy.s.sol:DeployStrategy --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv

// NOTE: ONLY WORKS ON GOERLI
// CommitHash: eccdfd43bb882d66a68cad8875dde2979e204546
contract GoerliUpgrade1 is Script, Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    string public deploymentOutputPath = string(bytes("script/output/M1_deployment_goerli_2023_3_23.json"));

    // EigenLayer Contract


    function run() external {
        // read and log the chainID
        uint256 chainId = block.chainid;
        emit log_named_uint("You are deploying on ChainID", chainId);

        string memory config_data = vm.readFile(deploymentOutputPath);
        IStrategyManager strategyManager = IStrategyManager(stdJson.readAddress(config_data, ".addresses.strategyManager"));
        IDelegationManager delegation = IDelegationManager(stdJson.readAddress(config_data, ".addresses.delegation"));
        IEigenPodManager eigenPodManager = IEigenPodManager(stdJson.readAddress(config_data, ".addresses.eigenPodManager"));
        // IBeacon eigenPodBeacon = IBeacon(stdJson.readAddress(config_data, ".addresses.eigenPodBeacon"));
        ISlasher slasher = ISlasher(stdJson.readAddress(config_data, ".addresses.slasher"));
        IDelayedWithdrawalRouter delayedWithdrawalRouter = IDelayedWithdrawalRouter(stdJson.readAddress(config_data, ".addresses.delayedWithdrawalRouter"));

        vm.startBroadcast();

        address strategyManagerImplementation = address(
            new StrategyManager(
                delegation,
                eigenPodManager,
                slasher
            )
        );

        address slasherImplementation = address(
            new Slasher(
                strategyManager,
                delegation
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

        vm.stopBroadcast();

        emit log_named_address("StrategyManagerImplementation", strategyManagerImplementation);
        emit log_named_address("SlasherImplementation", slasherImplementation);
        emit log_named_address("EigenPodImplementation", eigenPodImplementation);

        //   StrategyManagerImplementation: 0x1b8a566357c21b8b7b7c738a6963e2374718ea94
        //   SlasherImplementation: 0x2f82092969d156da92f0b787525042735fc4774a
        //   EigenPodImplementation: 0x4dd49853a27e3d4a0557876fe225ffce9b6b5d7a
    }
}
