// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/users/User.t.sol";
import {console} from "forge-std/console.sol";
import {BN254} from "src/contracts/libraries/BN254.sol";

// TODO: instead of hardcoding bn254 keys, generate randomly per operator
contract BN254OperatorSet_CalculateWeights is IntegrationCheckUtils {
    using BN254 for BN254.G1Point;

    AVS avs;
    OperatorSet operatorSet;
    BN254TableCalculator tableCalculator;

    address[] operators;
    User operator1;
    User operator2;
    AllocateParams allocateParams;
    AllocateParams allocateParams2;

    User staker1;
    User staker2;
    IStrategy[] strategies;
    // IStrategy[] strategies2;
    uint[] initTokenBalances1;
    uint[] initTokenBalances2;
    uint[] initDepositShares1;
    uint[] initDepositShares2;

    // -----------------------------------------------------------------------
    // Test keys for BN254
    // key 1
    uint public bn254PrivKey1 = 69;
    BN254.G1Point bn254G1Key1;
    BN254.G2Point bn254G2Key1;
    bytes public bn254Key1;
    // key 2
    uint public bn254PrivKey2 = 123;
    BN254.G1Point bn254G1Key2;
    BN254.G2Point bn254G2Key2;
    bytes public bn254Key2;

    function _init() internal override {
        _configUserTypes(DEFAULT);

        // Initialize table calculator
        tableCalculator = new BN254TableCalculator(
            IKeyRegistrar(address(keyRegistrar)), IAllocationManager(address(allocationManager)), BN254_TABLE_CALCULATOR_LOOKAHEAD_BLOCKS
        );

        // Setup stakers and operators
        (staker1, strategies, initTokenBalances1) = _newRandomStaker();
        (staker2, initTokenBalances2) = _newStaker(strategies);
        (operator1,,) = _newRandomOperator();
        (operator2,,) = _newRandomOperator();
        operators = new address[](2);
        operators[0] = address(operator1);
        operators[1] = address(operator2);
        (avs,) = _newRandomAVS();

        // 1. Deposit Into Strategies for both stakers
        initDepositShares1 = _calculateExpectedShares(strategies, initTokenBalances1);
        staker1.depositIntoEigenlayer(strategies, initTokenBalances1);
        check_Deposit_State(staker1, strategies, initDepositShares1);

        initDepositShares2 = _calculateExpectedShares(strategies, initTokenBalances2);
        staker2.depositIntoEigenlayer(strategies, initTokenBalances2);
        check_Deposit_State(staker2, strategies, initDepositShares2);

        // 2. Delegate both stakers to their respective operators
        staker1.delegateTo(operator1);
        check_Delegation_State(staker1, operator1, strategies, initDepositShares1);

        staker2.delegateTo(operator2);
        check_Delegation_State(staker2, operator2, strategies, initDepositShares2);

        // 3. Create operator set with configured BN254 keytype
        operatorSet = avs.createOperatorSet(strategies);
        avs.configureOperatorSet(operatorSet, IKeyRegistrarTypes.CurveType.BN254);

        // 4. Register both operators for operator set
        operator1.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator1, operatorSet, allStrats);

        operator2.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator2, operatorSet, allStrats);

        // 5. Allocate both operators to operator set
        allocateParams = _genAllocation_AllAvailable(operator1, operatorSet);
        operator1.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable(operator1, allocateParams);

        allocateParams2 = _genAllocation_AllAvailable(operator2, operatorSet);
        operator2.modifyAllocations(allocateParams2);
        check_IncrAlloc_State_Slashable(operator2, allocateParams2);

        // 6. Wait for allocations to complete
        // note: both operators have same allocationDelay so just roll for operator1
        _rollBlocksForCompleteAllocation(operator1, operatorSet, strategies);

        // 7. Register BN254 keys for both operators
        // Operator 1 key setup
        bn254G1Key1 = BN254.generatorG1().scalar_mul(bn254PrivKey1);
        bn254G2Key1.X[1] = 19_101_821_850_089_705_274_637_533_855_249_918_363_070_101_489_527_618_151_493_230_256_975_900_223_847;
        bn254G2Key1.X[0] = 5_334_410_886_741_819_556_325_359_147_377_682_006_012_228_123_419_628_681_352_847_439_302_316_235_957;
        bn254G2Key1.Y[1] = 354_176_189_041_917_478_648_604_979_334_478_067_325_821_134_838_555_150_300_539_079_146_482_658_331;
        bn254G2Key1.Y[0] = 4_185_483_097_059_047_421_902_184_823_581_361_466_320_657_066_600_218_863_748_375_739_772_335_928_910;
        operator1.registerBN254Key(operatorSet, bn254G1Key1, bn254G2Key1, bn254PrivKey1);

        // Operator 2 key setup
        bn254G1Key2 = BN254.generatorG1().scalar_mul(bn254PrivKey2);
        bn254G2Key2.X[1] = 19_276_105_129_625_393_659_655_050_515_259_006_463_014_579_919_681_138_299_520_812_914_148_935_621_072;
        bn254G2Key2.X[0] = 14_066_454_060_412_929_535_985_836_631_817_650_877_381_034_334_390_275_410_072_431_082_437_297_539_867;
        bn254G2Key2.Y[1] = 12_642_665_914_920_339_463_975_152_321_804_664_028_480_770_144_655_934_937_445_922_690_262_428_344_269;
        bn254G2Key2.Y[0] = 10_109_651_107_942_685_361_120_988_628_892_759_706_059_655_669_161_016_107_907_096_760_613_704_453_218;
        operator2.registerBN254Key(operatorSet, bn254G1Key2, bn254G2Key2, bn254PrivKey2);
    }

    function testFuzz_calculateOperatorWeights(uint24 _random) public rand(_random) {
        // Calculate operator weights
        uint[][] memory weights;
        (operators, weights) = tableCalculator.getOperatorWeights(operatorSet);

        // Verify results
        assertEq(operators.length, 2, "Should have exactly two operators");
        assertEq(operators[0], address(operator1), "Operator address should match");
        assertEq(operators[1], address(operator2), "Operator address should match");
        assertEq(weights.length, 2, "Should have weights for two operators");
        assertEq(weights[0].length, 1, "Should have one weight type (slashable stake)");
        assertEq(weights[1].length, 1, "Should have one weight type (slashable stake)");

        // Get the expected weight from allocation manager
        uint[][] memory slashableStake = allocationManager.getMinimumSlashableStake(
            operatorSet, operators, strategies, uint32(block.number + tableCalculator.LOOKAHEAD_BLOCKS())
        );

        // The BN254TableCalculator simply sums all the slashable stake so this should the expected weight
        // should simply be the sum of all the slashable stake from AllocationManager.getMinimumSlashableStake
        for (uint i = 0; i < slashableStake.length; i++) {
            // Loop through each operator and calculate their weight
            uint operatorWeight = 0;
            for (uint j = 0; j < slashableStake[i].length; j++) {
                operatorWeight += slashableStake[i][j];
            }
            assertEq(operatorWeight, weights[i][0], "Operator weight should match");
        }

        // verify both view functions return the same weight
        // `getOperatorWeight` matches same weight from `getOperatorWeights`
        assertEq(tableCalculator.getOperatorWeight(operatorSet, address(operator1)), weights[0][0], "Operator weight should match");
        assertEq(tableCalculator.getOperatorWeight(operatorSet, address(operator2)), weights[1][0], "Operator weight should match");
    }

    function testFuzz_calculateOperatorTable(uint24 _random) public rand(_random) {
        // Calculate operator table
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = tableCalculator.calculateOperatorTable(operatorSet);

        // Verify results
        assertEq(operatorSetInfo.numOperators, 2, "Should have exactly two operators");
        assertEq(operatorSetInfo.totalWeights.length, 1, "Should have one weight type");

        // Get the expected weight from allocation manager
        uint[][] memory slashableStake = allocationManager.getMinimumSlashableStake(
            operatorSet, operators, strategies, uint32(block.number + tableCalculator.LOOKAHEAD_BLOCKS())
        );

        uint expectedTotalWeight = 0;
        for (uint i = 0; i < slashableStake.length; i++) {
            // Loop through each operator and calculate their weight
            uint operatorWeight = 0;
            for (uint j = 0; j < slashableStake[i].length; j++) {
                operatorWeight += slashableStake[i][j];
            }
            expectedTotalWeight += operatorWeight;
        }

        assertEq(
            operatorSetInfo.totalWeights[0], expectedTotalWeight, "Total weight should match minimum slashable stake of both operators"
        );

        // Verify aggregate pubkey matches operator's pubkey
        BN254.G1Point memory expectedApk = bn254G1Key1.plus(bn254G1Key2);
        assertEq(operatorSetInfo.aggregatePubkey.X, expectedApk.X, "Aggregate pubkey X should match");
        assertEq(operatorSetInfo.aggregatePubkey.Y, expectedApk.Y, "Aggregate pubkey Y should match");
        assertEq(operatorSetInfo.numOperators, 2, "Number of operators should match");
    }

    function testFuzz_getOperatorInfos(uint24 _random) public rand(_random) {
        // Get operator infos
        IBN254TableCalculatorTypes.BN254OperatorInfo[] memory operatorInfos = tableCalculator.getOperatorInfos(operatorSet);

        // Verify results
        assertEq(operatorInfos.length, 2, "Should have exactly two operators");

        // Verify operator1 info
        assertEq(operatorInfos[0].pubkey.X, bn254G1Key1.X, "Operator1 pubkey X should match");
        assertEq(operatorInfos[0].pubkey.Y, bn254G1Key1.Y, "Operator1 pubkey Y should match");
        assertEq(operatorInfos[0].weights.length, 1, "Operator1 should have one weight type");

        // Verify operator2 info
        assertEq(operatorInfos[1].pubkey.X, bn254G1Key2.X, "Operator2 pubkey X should match");
        assertEq(operatorInfos[1].pubkey.Y, bn254G1Key2.Y, "Operator2 pubkey Y should match");
        assertEq(operatorInfos[1].weights.length, 1, "Operator2 should have one weight type");

        // Verify weights match getOperatorWeight for each operator
        assertEq(
            operatorInfos[0].weights[0],
            tableCalculator.getOperatorWeight(operatorSet, address(operator1)),
            "Operator1 weight should match getOperatorWeight"
        );
        assertEq(
            operatorInfos[1].weights[0],
            tableCalculator.getOperatorWeight(operatorSet, address(operator2)),
            "Operator2 weight should match getOperatorWeight"
        );

        // Verify weights match minimum slashable stake from allocation manager
        uint[][] memory slashableStake = allocationManager.getMinimumSlashableStake(
            operatorSet, operators, strategies, uint32(block.number + tableCalculator.LOOKAHEAD_BLOCKS())
        );

        // Calculate and verify weights for each operator
        for (uint i = 0; i < slashableStake.length; i++) {
            uint operatorWeight = 0;
            for (uint j = 0; j < slashableStake[i].length; j++) {
                operatorWeight += slashableStake[i][j];
            }
            assertEq(operatorWeight, operatorInfos[i].weights[0], "Operator weight should match minimum slashable stake");
        }
    }

    function testFuzz_getOperatorWeight(uint24 _random) public rand(_random) {
        // Get operator weight
        uint weight = tableCalculator.getOperatorWeight(operatorSet, address(operator1));

        // Get the expected weight from allocation manager
        operators = new address[](1);
        operators[0] = address(operator1);
        uint[][] memory slashableStake = allocationManager.getMinimumSlashableStake(
            operatorSet, operators, strategies, uint32(block.number + tableCalculator.LOOKAHEAD_BLOCKS())
        );

        uint expectedWeight = 0;
        for (uint i = 0; i < slashableStake.length; i++) {
            for (uint j = 0; j < slashableStake[i].length; j++) {
                expectedWeight += slashableStake[i][j];
            }
        }

        assertEq(weight, expectedWeight, "Weight should match minimum slashable stake");
    }

    function testFuzz_deallocate_calculateOperatorTable(uint24 _random) public rand(_random) {
        // 1. First verify initial state is same as calculateOperatorTable
        IBN254TableCalculatorTypes.BN254OperatorSetInfo memory operatorSetInfo = tableCalculator.calculateOperatorTable(operatorSet);

        // Verify initial results
        assertEq(operatorSetInfo.numOperators, 2, "Should have exactly two operators");
        assertEq(operatorSetInfo.totalWeights.length, 1, "Should have one weight type");

        // Get the expected weight from allocation manager
        uint[][] memory slashableStake = allocationManager.getMinimumSlashableStake(
            operatorSet, operators, strategies, uint32(block.number + tableCalculator.LOOKAHEAD_BLOCKS())
        );

        uint expectedTotalWeight = 0;
        uint[] memory operatorWeights = new uint[](2);
        for (uint i = 0; i < slashableStake.length; i++) {
            // Loop through each operator and calculate their weight
            for (uint j = 0; j < slashableStake[i].length; j++) {
                operatorWeights[i] += slashableStake[i][j];
            }
            expectedTotalWeight += operatorWeights[i];
        }

        assertEq(
            operatorSetInfo.totalWeights[0], expectedTotalWeight, "Total weight should match minimum slashable stake of both operators"
        );

        // Verify aggregate pubkey matches operator's pubkey
        BN254.G1Point memory expectedApk = bn254G1Key1.plus(bn254G1Key2);
        assertEq(operatorSetInfo.aggregatePubkey.X, expectedApk.X, "Aggregate pubkey X should match");
        assertEq(operatorSetInfo.aggregatePubkey.Y, expectedApk.Y, "Aggregate pubkey Y should match");

        // 2. Now deallocate operator1's stake
        AllocateParams memory deallocateParams = _genDeallocation_Full(operator1, operatorSet);
        operator1.modifyAllocations(deallocateParams);

        // Immediately after deallocation, weight should still be the same due to delay
        operatorSetInfo = tableCalculator.calculateOperatorTable(operatorSet);
        // check operator weight != 0
        assertGt(tableCalculator.getOperatorWeight(operatorSet, address(operator1)), 0, "Operator1 weight should not be 0");
        assertGt(tableCalculator.getOperatorWeight(operatorSet, address(operator2)), 0, "Operator2 weight should not be 0");
        // check total weight != 0
        assertEq(operatorSetInfo.totalWeights[0], expectedTotalWeight, "Total weight should still match before delay period");
        assertEq(
            tableCalculator.getOperatorWeight(operatorSet, address(operator1)),
            operatorWeights[0],
            "Operator1 weight should still match before delay period"
        );

        // Roll blocks forward to complete deallocation
        _rollBlocksForCompleteAllocation(operator1, operatorSet, strategies);

        // 3. After delay period, operator1's weight should be 0
        operatorSetInfo = tableCalculator.calculateOperatorTable(operatorSet);
        uint expectedWeightAfterDeallocation = expectedTotalWeight - operatorWeights[0];
        assertEq(operatorSetInfo.totalWeights[0], expectedWeightAfterDeallocation, "Total weight should be reduced after deallocation");
        assertEq(tableCalculator.getOperatorWeight(operatorSet, address(operator1)), 0, "Operator1 weight should be 0 after deallocation");

        // Verify operator2's weight remains unchanged
        assertEq(
            tableCalculator.getOperatorWeight(operatorSet, address(operator2)),
            operatorWeights[1],
            "Operator2 weight should remain unchanged"
        );

        // Verify aggregate pubkey is still the same even though operator1's weight is 0
        assertEq(operatorSetInfo.aggregatePubkey.X, expectedApk.X, "Aggregate pubkey X should match");
        assertEq(operatorSetInfo.aggregatePubkey.Y, expectedApk.Y, "Aggregate pubkey Y should match");
    }
}
