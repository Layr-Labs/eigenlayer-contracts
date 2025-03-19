// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Script, stdJson} from "forge-std/Script.sol";

import {IRewardsCoordinator, IRewardsCoordinatorTypes} from "src/contracts/interfaces/IRewardsCoordinator.sol";
import {IAllocationManager} from "src/contracts/interfaces/IAllocationManager.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IStrategy} from "src/contracts/interfaces/IStrategy.sol";
import {OperatorSet} from "src/contracts/libraries/OperatorSetLib.sol";

contract TestOpSetRewards is Script {
    IRewardsCoordinator rewardsCoordinator = IRewardsCoordinator(0xb22Ef643e1E067c994019A4C19e403253C05c2B0);
    IAllocationManager allocationManager = IAllocationManager(0xFdD5749e11977D60850E06bF5B13221Ad95eb6B4);

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
    // forge script script/TestOpSetRewards.s.sol:TestOpSetRewards --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_3()' --private-key '<0xDA29BB71669f46F2a779b4b62f03644A84eE3479_PRIV_KEY>' -vvvv --broadcast
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

    // Set Operator Set Split:
    // forge script script/TestOpSetRewards.s.sol:TestOpSetRewards --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_4()' --private-key '<0x8D8A8D3f88f6a6DA2083D865062bFBE3f1cfc293_PRIV_KEY>' -vvvv --broadcast
    function tx_4() public {
        OperatorSet memory operatorSet = OperatorSet({avs: AVS_A, id: OPERATOR_SET_ID_0});

        vm.startBroadcast();
        rewardsCoordinator.setOperatorSetSplit(OPERATOR_3, operatorSet, 3000);
        vm.stopBroadcast();
    }

    // Test Operator Directed Operator Set Rewards Submission: Test Operator Set Split
    // forge script script/TestOpSetRewards.s.sol:TestOpSetRewards --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_5()' --private-key '<0x8D8A8D3f88f6a6DA2083D865062bFBE3f1cfc293_PRIV_KEY>' -vvvv --broadcast
    function tx_5() public {
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
            startTimestamp: uint32(1_738_627_200), // 2025-02-04 00:00:00 UTC
            duration: uint32(432_000), // 5 days
            description: "test opset rewards"
        });

        vm.startBroadcast();
        STETH.approve(address(rewardsCoordinator), totalAmount);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, rewardsSubmissions);
        vm.stopBroadcast();
    }

    // Remove Strategies from Operator Set:
    // forge script script/TestOpSetRewards.s.sol:TestOpSetRewards --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_6()' --private-key '<0x8D8A8D3f88f6a6DA2083D865062bFBE3f1cfc293_PRIV_KEY>' -vvvv --broadcast
    function tx_6() public {
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = STRATEGY_SLASHING_1;

        vm.startBroadcast();
        allocationManager.removeStrategiesFromOperatorSet(AVS_A, OPERATOR_SET_ID_0, strategies);
        vm.stopBroadcast();
    }

    // Test Operator Directed Operator Set Rewards Submission
    // forge script script/TestOpSetRewards.s.sol:TestOpSetRewards --rpc-url '<HOLESKY_RPC_URL>' --sig 'tx_7()' --private-key '<0x8D8A8D3f88f6a6DA2083D865062bFBE3f1cfc293_PRIV_KEY>' -vvvv --broadcast
    function tx_7() public {
        IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory defaultStrategyAndMultipliers =
            new IRewardsCoordinatorTypes.StrategyAndMultiplier[](1);

        defaultStrategyAndMultipliers[0] =
            IRewardsCoordinatorTypes.StrategyAndMultiplier({strategy: STRATEGY_SLASHING_1, multiplier: 1e18});

        IRewardsCoordinatorTypes.OperatorReward[] memory operatorRewards =
            new IRewardsCoordinatorTypes.OperatorReward[](1);
        operatorRewards[0] = IRewardsCoordinatorTypes.OperatorReward({
            operator: OPERATOR_3,
            amount: 1e17 // 0.1 WETH
        });

        uint256 totalAmount = _calculateTotalAmount(operatorRewards);

        OperatorSet memory operatorSet = OperatorSet({avs: AVS_A, id: OPERATOR_SET_ID_0});

        IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory rewardsSubmissions =
            new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
            strategiesAndMultipliers: defaultStrategyAndMultipliers,
            token: STETH,
            operatorRewards: operatorRewards,
            startTimestamp: uint32(1_741_996_800), // 2025-03-15 00:00:00 UTC
            duration: uint32(345_600), // 4 days
            description: "test opset rewards"
        });

        vm.startBroadcast();
        STETH.approve(address(rewardsCoordinator), totalAmount);
        rewardsCoordinator.createOperatorDirectedOperatorSetRewardsSubmission(operatorSet, rewardsSubmissions);
        vm.stopBroadcast();
    }
}
