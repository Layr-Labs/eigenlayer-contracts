// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Script, stdJson} from "forge-std/Script.sol";

import {IRewardsCoordinator, IRewardsCoordinatorTypes} from "src/contracts/interfaces/IRewardsCoordinator.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IStrategy} from "src/contracts/interfaces/IStrategy.sol";
import {OperatorSet} from "src/contracts/libraries/OperatorSetLib.sol";

contract TestOpSetRewards is Script {
    IRewardsCoordinator rewardsCoordinator = IRewardsCoordinator(0xb22Ef643e1E067c994019A4C19e403253C05c2B0);

    IERC20 STETH = IERC20(0x3F1c547b21f65e10480dE3ad8E19fAAC46C95034);
    IERC20 WETH = IERC20(0x94373a4919B3240D86eA41593D5eBa789FEF3848);

    // admin
    address ADMIN = 0x8D8A8D3f88f6a6DA2083D865062bFBE3f1cfc293;

    // operators
    address OPERATOR_3 = 0x93A797473810c125ECe22F25a2087b6cEb8cE886;

    // avs
    address AVS_A = 0xA7BAFB187a85f41c85B70f359D75F3cdaC62CD71;

    // Operator Set ID
    uint32 OPERATOR_SET_ID_0 = 0;

    // strategies
    IStrategy STRATEGY_STETH = IStrategy(0x5C8b55722f421556a2AAfb7A3EA63d4c3e514312);
    IStrategy STRATEGY_SLASHING_1 = IStrategy(0x947e522010E22856071f8fb03e735fEDFCcD6E9F);

    function _setupStrategyAndMultiplier()
        internal
        view
        returns (IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory)
    {
        IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory defaultStrategyAndMultipliers =
            new IRewardsCoordinatorTypes.StrategyAndMultiplier[](2);

        defaultStrategyAndMultipliers[0] =
            IRewardsCoordinatorTypes.StrategyAndMultiplier({strategy: STRATEGY_STETH, multiplier: 2e18});

        defaultStrategyAndMultipliers[1] =
            IRewardsCoordinatorTypes.StrategyAndMultiplier({strategy: STRATEGY_SLASHING_1, multiplier: 1e18});

        return defaultStrategyAndMultipliers;
    }

    function _calculateTotalAmount(
        IRewardsCoordinatorTypes.OperatorReward[] memory operatorRewards
    ) internal pure returns (uint256) {
        uint256 totalAmount = 0;
        for (uint256 i = 0; i < operatorRewards.length; i++) {
            totalAmount += operatorRewards[i].amount;
        }
        return totalAmount;
    }

    // Test Operator Directed Operator Set Rewards Submission
    // forge script script/TestOpSetRewards.s.sol:TestOpSetRewards --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_1()' --private-key '<0x8D8A8D3f88f6a6DA2083D865062bFBE3f1cfc293_PRIV_KEY>' -vvvv --broadcast
    function tx_1() public {
        IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategyAndMultipliers = _setupStrategyAndMultiplier();

        IRewardsCoordinatorTypes.OperatorReward[] memory operatorRewards =
            new IRewardsCoordinatorTypes.OperatorReward[](1);
        operatorRewards[0] = IRewardsCoordinatorTypes.OperatorReward({
            operator: OPERATOR_3,
            amount: 9e16 // 0.09 stETH
        });

        uint256 totalAmount = _calculateTotalAmount(operatorRewards);

        OperatorSet memory operatorSet = OperatorSet({avs: AVS_A, id: OPERATOR_SET_ID_0});

        IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory rewardsSubmissions =
            new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: strategyAndMultipliers,
            token: STETH,
            operatorRewards: operatorRewards,
            startTimestamp: uint32(1_734_048_000), // 2024-12-13 00:00:00 UTC
            duration: uint32(432_000), // 5 days
            description: "test opset rewards"
        });

        vm.startBroadcast();
        STETH.approve(address(rewardsCoordinator), totalAmount);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, rewardsSubmissions);
        vm.stopBroadcast();
    }

    // Test Operator Directed Operator Set Rewards Submission: AVS refund
    // forge script script/TestOpSetRewards.s.sol:TestOpSetRewards --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_2()' --private-key '<0x8D8A8D3f88f6a6DA2083D865062bFBE3f1cfc293_PRIV_KEY>' -vvvv --broadcast
    function tx_2() public {
        IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategyAndMultipliers = _setupStrategyAndMultiplier();

        IRewardsCoordinatorTypes.OperatorReward[] memory operatorRewards =
            new IRewardsCoordinatorTypes.OperatorReward[](1);
        operatorRewards[0] = IRewardsCoordinatorTypes.OperatorReward({
            operator: OPERATOR_3,
            amount: 1e17 // 0.1 stETH
        });

        uint256 totalAmount = _calculateTotalAmount(operatorRewards);

        OperatorSet memory operatorSet = OperatorSet({avs: AVS_A, id: OPERATOR_SET_ID_0});

        IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory rewardsSubmissions =
            new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: strategyAndMultipliers,
            token: STETH,
            operatorRewards: operatorRewards,
            startTimestamp: uint32(1_733_788_800), // 2024-12-10 00:00:00 UTC
            duration: uint32(172_800), // 2 days
            description: "test opset rewards"
        });

        vm.startBroadcast();
        STETH.approve(address(rewardsCoordinator), totalAmount);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, rewardsSubmissions);
        vm.stopBroadcast();
    }

    // Test PI Rewards Submission:
    // forge script script/TestOpSetRewards.s.sol:TestOpSetRewards --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_3()' --private-key '<0x8D8A8D3f88f6a6DA2083D865062bFBE3f1cfc293_PRIV_KEY>' -vvvv --broadcast
    function tx_3() public {
        IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategyAndMultipliers = _setupStrategyAndMultiplier();

        uint256 amount = 1e17; // 0.1 WETH

        IRewardsCoordinatorTypes.RewardsSubmission[] memory rewardsSubmissions =
            new IRewardsCoordinatorTypes.RewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinatorTypes.RewardsSubmission({
            strategiesAndMultipliers: strategyAndMultipliers,
            token: WETH,
            amount: amount,
            startTimestamp: uint32(1_734_048_000), // 2024-12-13 00:00:00 UTC
            duration: uint32(432_000) // 5 days
        });

        vm.startBroadcast();
        WETH.approve(address(rewardsCoordinator), amount);
        rewardsCoordinator.createRewardsForAllEarners(rewardsSubmissions);
        vm.stopBroadcast();
    }

    // // Test Operator Directed Rewards Submission: Operator-avs split activated after startTimestamp and before duration end
    // // forge script script/TestRewardsV2.s.sol:TestRewardsV2 --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_6()' --private-key '<0xDA29BB71669f46F2a779b4b62f03644A84eE3479_PRIV_KEY>' -vvvv --broadcast
    // function tx_6() public {
    //     IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategyAndMultipliers = _setupStrategyAndMultiplier();

    //     IRewardsCoordinatorTypes.OperatorReward[] memory operatorRewards = new IRewardsCoordinatorTypes.OperatorReward[](1);
    //     operatorRewards[0] = IRewardsCoordinatorTypes.OperatorReward({
    //         operator: OPERATOR_EIGENYIELDS,
    //         amount: 1.2e18 // 1.2 WETH
    //     });

    //     uint256 totalAmount = _calculateTotalAmount(operatorRewards);

    //     IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory rewardsSubmissions =
    //         new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](1);
    //     rewardsSubmissions[0] = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
    //         strategiesAndMultipliers: strategyAndMultipliers,
    //         token: WETH,
    //         operatorRewards: operatorRewards,
    //         startTimestamp: uint32(1_733_788_800), // 2024-12-10 00:00:00 UTC
    //         duration: uint32(518_400), // 6 days
    //         description: ""
    //     });

    //     vm.startBroadcast();
    //     WETH.approve(address(eigenDAServiceManager), totalAmount);
    //     eigenDAServiceManager.createOperatorDirectedAVSRewardsSubmission(rewardsSubmissions);
    //     vm.stopBroadcast();
    // }

    // // Test Operator Directed Rewards Submission: Operator not registered to avs for entire duration
    // // forge script script/TestRewardsV2.s.sol:TestRewardsV2 --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_7()' --private-key '<0xDA29BB71669f46F2a779b4b62f03644A84eE3479_PRIV_KEY>' -vvvv --broadcast
    // function tx_7() public {
    //     IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategyAndMultipliers = _setupStrategyAndMultiplier();

    //     IRewardsCoordinatorTypes.OperatorReward[] memory operatorRewards = new IRewardsCoordinatorTypes.OperatorReward[](1);
    //     operatorRewards[0] = IRewardsCoordinatorTypes.OperatorReward({
    //         operator: OPERATOR_XYZ,
    //         amount: 0.9e18 // 0.9 WETH
    //     });

    //     uint256 totalAmount = _calculateTotalAmount(operatorRewards);

    //     IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory rewardsSubmissions =
    //         new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](1);
    //     rewardsSubmissions[0] = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
    //         strategiesAndMultipliers: strategyAndMultipliers,
    //         token: WETH,
    //         operatorRewards: operatorRewards,
    //         startTimestamp: uint32(1_734_048_000), // 2024-12-13 00:00:00 UTC
    //         duration: uint32(259_200), // 3 days
    //         description: ""
    //     });

    //     vm.startBroadcast();
    //     WETH.approve(address(eigenDAServiceManager), totalAmount);
    //     eigenDAServiceManager.createOperatorDirectedAVSRewardsSubmission(rewardsSubmissions);
    //     vm.stopBroadcast();
    // }

    // // Test Operator Directed Rewards Submission: Operator not registered to avs for partial duration
    // // forge script script/TestRewardsV2.s.sol:TestRewardsV2 --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_8()' --private-key '<0xDA29BB71669f46F2a779b4b62f03644A84eE3479_PRIV_KEY>' -vvvv --broadcast
    // function tx_8() public {
    //     IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategyAndMultipliers = _setupStrategyAndMultiplier();

    //     IRewardsCoordinatorTypes.OperatorReward[] memory operatorRewards = new IRewardsCoordinatorTypes.OperatorReward[](1);
    //     operatorRewards[0] = IRewardsCoordinatorTypes.OperatorReward({
    //         operator: OPERATOR_XYZ,
    //         amount: 1.4e18 // 1.4 WETH
    //     });

    //     uint256 totalAmount = _calculateTotalAmount(operatorRewards);

    //     IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory rewardsSubmissions =
    //         new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](1);
    //     rewardsSubmissions[0] = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
    //         strategiesAndMultipliers: strategyAndMultipliers,
    //         token: WETH,
    //         operatorRewards: operatorRewards,
    //         startTimestamp: uint32(1_733_702_400), // 2024-12-09 00:00:00 UTC
    //         duration: uint32(604_800), // 7 days
    //         description: ""
    //     });

    //     vm.startBroadcast();
    //     WETH.approve(address(eigenDAServiceManager), totalAmount);
    //     eigenDAServiceManager.createOperatorDirectedAVSRewardsSubmission(rewardsSubmissions);
    //     vm.stopBroadcast();
    // }

    // // Test Operator Directed Rewards Submission: Staker (0x9999Ee8B5cBA0688DbD4c4Cbbb821800758FbCDD) has withdrawal queued during the duration.
    // // forge script script/TestRewardsV2.s.sol:TestRewardsV2 --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_9()' --private-key '<0xDA29BB71669f46F2a779b4b62f03644A84eE3479_PRIV_KEY>' -vvvv --broadcast
    // function tx_9() public {
    //     IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategyAndMultipliers = _setupStrategyAndMultiplier();

    //     IRewardsCoordinatorTypes.OperatorReward[] memory operatorRewards = new IRewardsCoordinatorTypes.OperatorReward[](1);
    //     operatorRewards[0] = IRewardsCoordinatorTypes.OperatorReward({
    //         operator: OPERATOR_GALAXY,
    //         amount: 1.2e18 // 1.2 WETH
    //     });

    //     uint256 totalAmount = _calculateTotalAmount(operatorRewards);

    //     IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory rewardsSubmissions =
    //         new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](1);
    //     rewardsSubmissions[0] = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
    //         strategiesAndMultipliers: strategyAndMultipliers,
    //         token: WETH,
    //         operatorRewards: operatorRewards,
    //         startTimestamp: uint32(1_733_788_800), // 2024-12-10 00:00:00 UTC
    //         duration: uint32(518_400), // 6 days
    //         description: ""
    //     });

    //     vm.startBroadcast();
    //     WETH.approve(address(eigenDAServiceManager), totalAmount);
    //     eigenDAServiceManager.createOperatorDirectedAVSRewardsSubmission(rewardsSubmissions);
    //     vm.stopBroadcast();
    // }

    // // Test Operator Directed Rewards Submission: Staker (0x17C1c083e46F3924C33da32e4Aa117724DEcdc33) is undelegated during the duration.
    // // forge script script/TestRewardsV2.s.sol:TestRewardsV2 --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_10()' --private-key '<0xDA29BB71669f46F2a779b4b62f03644A84eE3479_PRIV_KEY>' -vvvv --broadcast
    // function tx_10() public {
    //     IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategyAndMultipliers = _setupStrategyAndMultiplier();

    //     IRewardsCoordinatorTypes.OperatorReward[] memory operatorRewards = new IRewardsCoordinatorTypes.OperatorReward[](1);
    //     operatorRewards[0] = IRewardsCoordinatorTypes.OperatorReward({
    //         operator: OPERATOR_SINOPMM,
    //         amount: 1.2e18 // 1.2 WETH
    //     });

    //     uint256 totalAmount = _calculateTotalAmount(operatorRewards);

    //     IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory rewardsSubmissions =
    //         new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](1);
    //     rewardsSubmissions[0] = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
    //         strategiesAndMultipliers: strategyAndMultipliers,
    //         token: WETH,
    //         operatorRewards: operatorRewards,
    //         startTimestamp: uint32(1_733_788_800), // 2024-12-10 00:00:00 UTC
    //         duration: uint32(518_400), // 6 days
    //         description: ""
    //     });

    //     vm.startBroadcast();
    //     WETH.approve(address(eigenDAServiceManager), totalAmount);
    //     eigenDAServiceManager.createOperatorDirectedAVSRewardsSubmission(rewardsSubmissions);
    //     vm.stopBroadcast();
    // }

    // // Test Operator Directed Rewards Submission: Dust payment of 1 wei over 1 day duration
    // // forge script script/TestRewardsV2.s.sol:TestRewardsV2 --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_11()' --private-key '<0xDA29BB71669f46F2a779b4b62f03644A84eE3479_PRIV_KEY>' -vvvv --broadcast
    // function tx_11() public {
    //     IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategyAndMultipliers = _setupStrategyAndMultiplier();

    //     IRewardsCoordinatorTypes.OperatorReward[] memory operatorRewards = new IRewardsCoordinatorTypes.OperatorReward[](1);
    //     operatorRewards[0] = IRewardsCoordinatorTypes.OperatorReward({
    //         operator: OPERATOR_STAKELY,
    //         amount: 1 // 1e-18 WETH (1 wei)
    //     });

    //     uint256 totalAmount = _calculateTotalAmount(operatorRewards);

    //     IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory rewardsSubmissions =
    //         new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](1);
    //     rewardsSubmissions[0] = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
    //         strategiesAndMultipliers: strategyAndMultipliers,
    //         token: WETH,
    //         operatorRewards: operatorRewards,
    //         startTimestamp: uint32(1_734_220_800), // 2024-12-15 00:00:00 UTC
    //         duration: uint32(86_400), // 1 day
    //         description: ""
    //     });

    //     vm.startBroadcast();
    //     WETH.approve(address(eigenDAServiceManager), totalAmount);
    //     eigenDAServiceManager.createOperatorDirectedAVSRewardsSubmission(rewardsSubmissions);
    //     vm.stopBroadcast();
    // }

    // // Test Operator Directed Rewards Submission: Dust payment of 1 wei over 6 days duration
    // // forge script script/TestRewardsV2.s.sol:TestRewardsV2 --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_12()' --private-key '<0xDA29BB71669f46F2a779b4b62f03644A84eE3479_PRIV_KEY>' -vvvv --broadcast
    // function tx_12() public {
    //     IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategyAndMultipliers = _setupStrategyAndMultiplier();

    //     IRewardsCoordinatorTypes.OperatorReward[] memory operatorRewards = new IRewardsCoordinatorTypes.OperatorReward[](1);
    //     operatorRewards[0] = IRewardsCoordinatorTypes.OperatorReward({
    //         operator: OPERATOR_STAKELY,
    //         amount: 1 // 1e-18 WETH (1 wei)
    //     });

    //     uint256 totalAmount = _calculateTotalAmount(operatorRewards);

    //     IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory rewardsSubmissions =
    //         new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](1);
    //     rewardsSubmissions[0] = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
    //         strategiesAndMultipliers: strategyAndMultipliers,
    //         token: WETH,
    //         operatorRewards: operatorRewards,
    //         startTimestamp: uint32(1_733_788_800), // 2024-12-10 00:00:00 UTC
    //         duration: uint32(518_400), // 6 days
    //         description: ""
    //     });

    //     vm.startBroadcast();
    //     WETH.approve(address(eigenDAServiceManager), totalAmount);
    //     eigenDAServiceManager.createOperatorDirectedAVSRewardsSubmission(rewardsSubmissions);
    //     vm.stopBroadcast();
    // }

    // // Test Operator Directed Rewards Submission: Dust payment of 60 wei over 6 days duration
    // // forge script script/TestRewardsV2.s.sol:TestRewardsV2 --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_13()' --private-key '<0xDA29BB71669f46F2a779b4b62f03644A84eE3479_PRIV_KEY>' -vvvv --broadcast
    // function tx_13() public {
    //     IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategyAndMultipliers = _setupStrategyAndMultiplier();

    //     IRewardsCoordinatorTypes.OperatorReward[] memory operatorRewards = new IRewardsCoordinatorTypes.OperatorReward[](1);
    //     operatorRewards[0] = IRewardsCoordinatorTypes.OperatorReward({
    //         operator: OPERATOR_STAKELY,
    //         amount: 60 // 60e-18 WETH (1 wei)
    //     });

    //     uint256 totalAmount = _calculateTotalAmount(operatorRewards);

    //     IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory rewardsSubmissions =
    //         new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](1);
    //     rewardsSubmissions[0] = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
    //         strategiesAndMultipliers: strategyAndMultipliers,
    //         token: WETH,
    //         operatorRewards: operatorRewards,
    //         startTimestamp: uint32(1_733_788_800), // 2024-12-10 00:00:00 UTC
    //         duration: uint32(518_400), // 6 days
    //         description: ""
    //     });

    //     vm.startBroadcast();
    //     WETH.approve(address(eigenDAServiceManager), totalAmount);
    //     eigenDAServiceManager.createOperatorDirectedAVSRewardsSubmission(rewardsSubmissions);
    //     vm.stopBroadcast();
    // }

    // // Test Operator Directed Rewards Submission: 1. Edge case commissions of 0 and 10000 bips. 2. Two split setting in the same day.
    // // forge script script/TestRewardsV2.s.sol:TestRewardsV2 --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_14()' --private-key '<0xDA29BB71669f46F2a779b4b62f03644A84eE3479_PRIV_KEY>' -vvvv --broadcast
    // function tx_14() public {
    //     IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategyAndMultipliers = _setupStrategyAndMultiplier();

    //     IRewardsCoordinatorTypes.OperatorReward[] memory operatorRewards = new IRewardsCoordinatorTypes.OperatorReward[](1);
    //     operatorRewards[0] = IRewardsCoordinatorTypes.OperatorReward({
    //         operator: OPERATOR_EIGENLABS,
    //         amount: 0.9e18 // 0.9 WETH
    //     });

    //     uint256 totalAmount = _calculateTotalAmount(operatorRewards);

    //     IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory rewardsSubmissions =
    //         new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](1);
    //     rewardsSubmissions[0] = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
    //         strategiesAndMultipliers: strategyAndMultipliers,
    //         token: WETH,
    //         operatorRewards: operatorRewards,
    //         startTimestamp: uint32(1_733_788_800), // 2024-12-10 00:00:00 UTC
    //         duration: uint32(950_400), // 11 days
    //         description: ""
    //     });

    //     vm.startBroadcast();
    //     WETH.approve(address(eigenDAServiceManager), totalAmount);
    //     eigenDAServiceManager.createOperatorDirectedAVSRewardsSubmission(rewardsSubmissions);
    //     vm.stopBroadcast();
    // }

    // // Test Rewards v1 Submission: 1. Edge case commissions of 0 and 10000 bips. 2. Two split setting in the same day.
    // // forge script script/TestRewardsV2.s.sol:TestRewardsV2 --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_15()' --private-key '<0xDA29BB71669f46F2a779b4b62f03644A84eE3479_PRIV_KEY>' -vvvv --broadcast
    // function tx_15() public {
    //     IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategyAndMultipliers = _setupStrategyAndMultiplier();

    //     uint256 amount = 1e18; // 1 WETH

    //     IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions =
    //         new IRewardsCoordinator.RewardsSubmission[](1);
    //     rewardsSubmissions[0] = IRewardsCoordinator.RewardsSubmission({
    //         strategiesAndMultipliers: strategyAndMultipliers,
    //         token: WETH,
    //         amount: amount,
    //         startTimestamp: uint32(1_733_788_800), // 2024-12-10 00:00:00 UTC
    //         duration: uint32(950_400) // 11 days
    //     });

    //     vm.startBroadcast();
    //     WETH.approve(address(eigenDAServiceManager), amount);
    //     eigenDAServiceManager.createAVSRewardsSubmission(rewardsSubmissions);
    //     vm.stopBroadcast();
    // }
}
