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

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "../../src/test/mocks/EmptyContract.sol";
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

    string public deploymentOutputPath = string(bytes("script/output/M2_preprod_deployment_from_scratch.json"));

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
        DelegationImplementation: 0x934eB3E2b6D5C2E1601B29B7180026D71438F20D
        SlasherImplementation: 0x05c235183e8b9dFb7113Cf92bbDc3f5085324158
        StrategyManagerImplementation: 0xb9B69504f1a727E783F4B4248A115D56F4080DF8
        DelayedWithdrawalRouterImplementation: 0x44a40C60857b4B420Ad3D8b9646FefEBF2D0dB86
        EigenPodImplementation: 0x83cbB48391F428878Bc5DD97C9792a8dbCAa0729
        EigenPodManagerImplementation: 0xEEdCC9dB001fB8429721FE21426F51f0Cdd329EC
        */
    }
}