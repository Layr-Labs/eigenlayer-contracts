// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {Slasher, ISlasher} from "../../contracts/core/Slasher.sol";
import {PauserRegistry} from "../../contracts/permissions/PauserRegistry.sol";
import {IStakeRegistry} from "../../contracts/interfaces/IStakeRegistry.sol";
import {IServiceManager} from "../../contracts/interfaces/IServiceManager.sol";
import {IVoteWeigher} from "../../contracts/interfaces/IVoteWeigher.sol";

import {BitmapUtils} from "../../contracts/libraries/BitmapUtils.sol";

import {StrategyManagerMock, IStrategy} from "../mocks/StrategyManagerMock.sol";
import {EigenPodManagerMock} from "../mocks/EigenPodManagerMock.sol";
import {ServiceManagerMock} from "../mocks/ServiceManagerMock.sol";
import {DelegationMock} from "../mocks/DelegationMock.sol";
import {RegistryCoordinatorMock} from "../mocks/RegistryCoordinatorMock.sol";

import {StakeRegistryHarness, StakeRegistry, IRegistryCoordinator} from "../harnesses/StakeRegistryHarness.sol";

import {Test,console} from "forge-std/Test.sol";


contract StakeRegistryTest is Test {
    bytes constant internal ONLY_REGISTRY_COORDINATOR = "StakeRegistry.onlyRegistryCoordinator: caller is not the RegistryCoordinator";
    bytes constant internal ALREADY_INITIALIZED = "Initializable: contract is already initialized";
    bytes constant internal ONLY_SERVICE_MANAGER = "VoteWeigherBase.onlyServiceManagerOwner: caller is not the owner of the serviceManager";
    bytes constant internal GREATEST_QUORUM_GT_QUORUM_COUNT = "StakeRegistry._registerOperator: greatest quorumNumber must be less than quorumCount";
    bytes constant internal OPERATOR_INSUFFICIENT_STAKE_FOR_QUORUM = "StakeRegistry._registerOperator: Operator does not meet minimum stake requirement for quorum";
    bytes constant internal LENGTH_MISMATCH = "Registry._initialize: minimumStakeForQuorum length mismatch";
    address constant internal serviceManagerOwner = address(uint160(uint256(keccak256("serviceManagerOwner"))));
    address constant internal pauser = address(uint160(uint256(keccak256("pauser"))));
    address constant internal unpauser = address(uint160(uint256(keccak256("unpauser"))));
    address constant internal defaultOperator = address(uint160(uint256(keccak256("defaultOperator"))));

    ProxyAdmin internal proxyAdmin;
    PauserRegistry internal pauserRegistry;
    Slasher internal slasherImplementation;
    StakeRegistryHarness internal stakeRegistryImplementation;
    StakeRegistryHarness internal stakeRegistry;
    ISlasher internal slasher;

    RegistryCoordinatorMock internal registryCoordinatorMock;
    ServiceManagerMock internal serviceManagerMock;
    StrategyManagerMock internal strategyManagerMock;
    DelegationMock internal delegationMock;
    EigenPodManagerMock internal eigenPodManagerMock;

    bytes32 internal defaultOperatorId = keccak256("defaultOperatorId");
    uint8 internal defaultQuorumNumber = 0;
    uint8 internal numQuorums = 192;

    /// @notice emitted whenever the stake of `operator` is updated
    event StakeUpdate(
        bytes32 indexed operatorId,
        uint8 quorumNumber,
        uint96 stake
    );

    function setUp() virtual public {
        proxyAdmin = new ProxyAdmin();

        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserRegistry = new PauserRegistry(pausers, unpauser);

        registryCoordinatorMock = new RegistryCoordinatorMock();
        delegationMock = new DelegationMock();
        eigenPodManagerMock = new EigenPodManagerMock();
        strategyManagerMock = new StrategyManagerMock();
        slasherImplementation = new Slasher(strategyManagerMock, delegationMock);
        slasher = Slasher(
            address(
                new TransparentUpgradeableProxy(
                    address(slasherImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(Slasher.initialize.selector, msg.sender, pauserRegistry, 0/*initialPausedStatus*/)
                )
            )
        );

        strategyManagerMock.setAddresses(
            delegationMock,
            eigenPodManagerMock,
            slasher
        );

        vm.startPrank(serviceManagerOwner);
        // make the serviceManagerOwner the owner of the serviceManager contract
        serviceManagerMock = new ServiceManagerMock(slasher);
        stakeRegistryImplementation = new StakeRegistryHarness(
            IRegistryCoordinator(address(registryCoordinatorMock)),
            strategyManagerMock,
            IServiceManager(address(serviceManagerMock))
        );

        // setup the dummy minimum stake for quorum
        uint96[] memory minimumStakeForQuorum = new uint96[](numQuorums);
        for (uint256 i = 0; i < minimumStakeForQuorum.length; i++) {
            minimumStakeForQuorum[i] = uint96(i+1);
        }

        // setup the dummy quorum strategies
        IVoteWeigher.StrategyAndWeightingMultiplier[][] memory quorumStrategiesConsideredAndMultipliers =
            new IVoteWeigher.StrategyAndWeightingMultiplier[][](numQuorums);
        for (uint256 i = 0; i < quorumStrategiesConsideredAndMultipliers.length; i++) {
            quorumStrategiesConsideredAndMultipliers[i] = new IVoteWeigher.StrategyAndWeightingMultiplier[](1);
            quorumStrategiesConsideredAndMultipliers[i][0] = IVoteWeigher.StrategyAndWeightingMultiplier(
                IStrategy(address(uint160(i))),
                uint96(i+1)
            );
        }

        stakeRegistry = StakeRegistryHarness(
            address(
                new TransparentUpgradeableProxy(
                    address(stakeRegistryImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        StakeRegistry.initialize.selector,
                        minimumStakeForQuorum,
                        quorumStrategiesConsideredAndMultipliers
                    )
                )
            )
        );

        vm.stopPrank();
    }
}
contract Initialize is StakeRegistryTest {
    function test_RevertsIf_AlreadyInitialized_Initialize() public {
        // make sure the contract intializers are disabled
        vm.expectRevert(ALREADY_INITIALIZED);
        stakeRegistryImplementation.initialize(new uint96[](0), new IVoteWeigher.StrategyAndWeightingMultiplier[][](0));
    }

    function test_RevertsIf_LengthMismatch_Initialize() public {
        vm.expectRevert(LENGTH_MISMATCH);
        new TransparentUpgradeableProxy(
            address(stakeRegistryImplementation),
            address(proxyAdmin),
            abi.encodeWithSelector(
                StakeRegistry.initialize.selector,
                new uint96[](0),
                new IVoteWeigher.StrategyAndWeightingMultiplier[][](1)
            )
        );

    }
}

contract UpdateStakes is StakeRegistryTest {
    function test_UpdateStakes() public {
        address[] memory operators = new address[](2);
        bytes32[] memory operatorIds = new bytes32[](2);
        uint256[] memory prevElements = new uint256[](2);

        (operators[0],operators[1]) = (defaultOperator, address(0xbad));
        (operatorIds[0], operatorIds[1]) = (defaultOperatorId, bytes32(uint256(0xbad)));
        vm.prank(serviceManagerOwner);
        stakeRegistry.setMinimumStakeForQuorum(defaultQuorumNumber, 0);

        stakeRegistry.setOperatorWeight(defaultQuorumNumber, defaultOperator, 1);
        stakeRegistry.setOperatorWeight(defaultQuorumNumber, address(0xbad), 2);
        stakeRegistry.updateStakes(operators, operatorIds, prevElements);
    }

    function test_RevertsIf_OperatorAndOperatorIdMismatch() public {
        address[] memory operators = new address[](2);
        bytes32[] memory operatorIds = new bytes32[](2);
        uint256[] memory prevElements = new uint256[](2);

        (operators[0],operators[1]) = (defaultOperator, address(0xbad));
        (operatorIds[1], operatorIds[0]) = (defaultOperatorId, bytes32(uint256(0xbad)));
        vm.prank(serviceManagerOwner);
        stakeRegistry.setMinimumStakeForQuorum(defaultQuorumNumber, 0);

        stakeRegistry.setOperatorWeight(defaultQuorumNumber, defaultOperator, 1);
        stakeRegistry.setOperatorWeight(defaultQuorumNumber, address(0xbad), 2);
        stakeRegistry.updateStakes(operators, operatorIds, prevElements);
    }

    function test_RevertsIf_OperatorStateIsAfterBlock_UpdateStake() public {}

    function test_RevertsIf_NewerOperatorStakeUpdateAvailable_UpdateStake() public {}
}
contract SetMinimumStakeForQuorum is StakeRegistryTest {
    function test_RevertsIf_NotServiceManager_SetMinimumStakeForQuorum() public {
        vm.expectRevert(ONLY_SERVICE_MANAGER);
        stakeRegistry.setMinimumStakeForQuorum(defaultQuorumNumber, 0);
    }
}

contract RegisterOperator is StakeRegistryTest{
    function test_RegisterOperator() public {
        stakeRegistry.setOperatorWeight(defaultQuorumNumber, defaultOperator, 1);
        stakeRegistry.updateOperatorStake(defaultOperator, defaultOperatorId, defaultQuorumNumber );
        vm.prank(serviceManagerOwner);
        stakeRegistry.setMinimumStakeForQuorum(defaultQuorumNumber, 0);
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        vm.prank(address(registryCoordinatorMock));
        stakeRegistry.registerOperator(defaultOperator, defaultOperatorId, quorumNumbers);
    }

    function test_RevertsIf_NotRegistryCoordinator_RegisterOperator() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        vm.expectRevert(ONLY_REGISTRY_COORDINATOR);
        stakeRegistry.registerOperator(defaultOperator, defaultOperatorId, quorumNumbers);
    }

    function test_RevertsIf_MoreQuorumsThanQuorumCount_RegisterOperator() public {
        bytes memory quorumNumbers = new bytes(numQuorums+1);
        for (uint i = 0; i < quorumNumbers.length; i++) {
            quorumNumbers[i] = bytes1(uint8(i));
        }

        // expect that it reverts when you register
        vm.expectRevert(GREATEST_QUORUM_GT_QUORUM_COUNT);
        vm.prank(address(registryCoordinatorMock));
        stakeRegistry.registerOperator(defaultOperator, defaultOperatorId, quorumNumbers);
    }
}

contract GetStakeUpdateIndexForOperatorIdForQuorumAtBlockNumber is StakeRegistryTest {
    function test_GetStakeUpdateIndexForOperatorIdForQuorumAtBlockNumber() public {
        bytes memory quorumNumbers = new bytes(1);

        vm.prank(address(registryCoordinatorMock));
        stakeRegistry.registerOperator(defaultOperator, defaultOperatorId, quorumNumbers);
        // stakeRegistry.getStakeUpdateForQuorumFromOperatorIdAndIndex(0, defaultOperatorId, 0);
    }

    function test_RevertsIf_OperatorHasNoStakeAtBlock_GetStakeUpdateIndexForOperatorIdForQuorumAtBlockNumber() public {
        /// if it the index is out of bounds 
        /// we probably want to be explicit about why
        // ie if there are multiple arrays that get access it's nice to quickly know where the index was out of bounds
        vm.expectRevert();
        stakeRegistry.getStakeUpdateForQuorumFromOperatorIdAndIndex(0, defaultOperatorId, 0);
    }

    function test_RevertsIf_NoStakeForOperatorAndQuorumAtBlock_GetStakeUpdateIndexForOperatorIdForQuorumAtBlockNumber() public {}

}

contract DereigsterOperator is StakeRegistryTest {
    function test_DeregisterOperator() public {}
    
    function test_RevertsIf_NotRegistryCoordinator_DeregisterOperator() public {}

    function test_RevertsIf_MoreQuorumsThanQuorumCounter_DeregisterOperator() public {}
}
contract CheckOperatorInactiveAtBlockNumber is StakeRegistryTest{
    function test_checkOperatorInactiveAtBlockNumber() public {}
}

contract StakeRegistryFuzzTests is StakeRegistryTest{ 
    function testFuzz_SetMinimumStakeForQuorum(uint8 quorumNumber, uint96 minimumStakeForQuorum) public {
        // set the minimum stake for quorum
        vm.prank(serviceManagerOwner);
        stakeRegistry.setMinimumStakeForQuorum(quorumNumber, minimumStakeForQuorum);

        // make sure the minimum stake for quorum is as expected
        assertEq(stakeRegistry.minimumStakeForQuorum(quorumNumber), minimumStakeForQuorum);
    }

    function testFuzz_RevertsIf_LessThanMinimumStakeForQuorum_RegisterOperator(
        uint96[] memory stakesForQuorum
    ) public {
        vm.assume(stakesForQuorum.length > 0);

        // set the weights of the operator
        // stakeRegistry.setOperatorWeight()

        bytes memory quorumNumbers = new bytes(stakesForQuorum.length > 128 ? 128 : stakesForQuorum.length);
        for (uint i = 0; i < quorumNumbers.length; i++) {
            quorumNumbers[i] = bytes1(uint8(i));
        }

        stakesForQuorum[stakesForQuorum.length - 1] = stakeRegistry.minimumStakeForQuorum(uint8(quorumNumbers.length - 1)) - 1;

        // expect that it reverts when you register
        vm.expectRevert(OPERATOR_INSUFFICIENT_STAKE_FOR_QUORUM);
        vm.prank(address(registryCoordinatorMock));
        stakeRegistry.registerOperator(defaultOperator, defaultOperatorId, quorumNumbers);
    }

    function testFuzz_RegisterFirstOperator(
        uint256 quorumBitmap,
        uint80[] memory stakesForQuorum
    ) public {

        uint96[] memory paddedStakesForQuorum = _registerOperatorValid(defaultOperator, defaultOperatorId, quorumBitmap, stakesForQuorum);

        uint8 quorumNumberIndex = 0;
        for (uint8 i = 0; i < 192; i++) {
            if (quorumBitmap >> i & 1 == 1) {
                // check that the operator has 1 stake update in the quorum numbers they registered for
                assertEq(stakeRegistry.getStakeHistoryLengthForQuorumNumber(defaultOperatorId, i), 1);
                // make sure that the stake update is as expected
                IStakeRegistry.OperatorStakeUpdate memory stakeUpdate =
                    stakeRegistry.getStakeUpdateForQuorumFromOperatorIdAndIndex(i, defaultOperatorId, 0);
                emit log_named_uint("length  of paddedStakesForQuorum", paddedStakesForQuorum.length);
                assertEq(stakeUpdate.stake, paddedStakesForQuorum[quorumNumberIndex]);
                assertEq(stakeUpdate.updateBlockNumber, uint32(block.number));
                assertEq(stakeUpdate.nextUpdateBlockNumber, 0);

                // make the analogous check for total stake history
                assertEq(stakeRegistry.getLengthOfTotalStakeHistoryForQuorum(i), 1);
                // make sure that the stake update is as expected
                stakeUpdate = stakeRegistry.getTotalStakeUpdateForQuorumFromIndex(i, 0);
                assertEq(stakeUpdate.stake, paddedStakesForQuorum[quorumNumberIndex]);
                assertEq(stakeUpdate.updateBlockNumber, uint32(block.number));
                assertEq(stakeUpdate.nextUpdateBlockNumber, 0);

                quorumNumberIndex++;
            } else {
                // check that the operator has 0 stake updates in the quorum numbers they did not register for
                assertEq(stakeRegistry.getStakeHistoryLengthForQuorumNumber(defaultOperatorId, i), 0);
                // make the analogous check for total stake history
                assertEq(stakeRegistry.getLengthOfTotalStakeHistoryForQuorum(i), 0);
            }
        }
    }

    function testFuzz_RegisterManyOperators_Valid(
        uint256 pseudoRandomNumber,
        uint8 numOperators,
        uint24[] memory blocksPassed
    ) public {
        vm.assume(numOperators > 0 && numOperators <= 15);
        // modulo so no overflow
        pseudoRandomNumber = pseudoRandomNumber % type(uint128).max;

        uint256[] memory quorumBitmaps = new uint256[](numOperators);

        // append to blocksPassed as needed
        uint24[] memory appendedBlocksPassed = new uint24[](quorumBitmaps.length);
        for (uint256 i = blocksPassed.length; i < quorumBitmaps.length; i++) {
            appendedBlocksPassed[i] = 0;
        }
        blocksPassed = appendedBlocksPassed;
        
        uint32 initialBlockNumber = 100;
        vm.roll(initialBlockNumber);
        uint32 cumulativeBlockNumber = initialBlockNumber;

        uint96[][] memory paddedStakesForQuorums = new uint96[][](quorumBitmaps.length);
        for (uint256 i = 0; i < quorumBitmaps.length; i++) {
            (quorumBitmaps[i], paddedStakesForQuorums[i]) = _registerOperatorRandomValid(_incrementAddress(defaultOperator, i), _incrementBytes32(defaultOperatorId, i), pseudoRandomNumber + i);

            cumulativeBlockNumber += blocksPassed[i];
            vm.roll(cumulativeBlockNumber);
        }
        
        // for each bit in each quorumBitmap, increment the number of operators in that quorum
        uint32[] memory numOperatorsInQuorum = new uint32[](192);
        for (uint256 i = 0; i < quorumBitmaps.length; i++) {
            for (uint256 j = 0; j < 192; j++) {
                if (quorumBitmaps[i] >> j & 1 == 1) {
                    numOperatorsInQuorum[j]++;
                }
            }
        }

        // operatorQuorumIndices is an array of iindices within the quorum numbers that each operator registered for
        // used for accounting in the next loops
        uint32[] memory operatorQuorumIndices = new uint32[](quorumBitmaps.length);

        // for each quorum
        for (uint8 i = 0; i < 192; i++) {
            uint32 operatorCount = 0;
            // reset the cumulative block number
            cumulativeBlockNumber = initialBlockNumber;

            // make sure the number of total stake updates is as expected
            assertEq(stakeRegistry.getLengthOfTotalStakeHistoryForQuorum(i), numOperatorsInQuorum[i]);

            uint96 cumulativeStake = 0;
            // for each operator
            for (uint256 j = 0; j < quorumBitmaps.length; j++) {
                // if the operator is in the quorum
                if (quorumBitmaps[j] >> i & 1 == 1) {
                    cumulativeStake += paddedStakesForQuorums[j][operatorQuorumIndices[j]];
                    // make sure the number of stake updates is as expected
                    assertEq(stakeRegistry.getStakeHistoryLengthForQuorumNumber(_incrementBytes32(defaultOperatorId, j), i), 1);

                    // make sure that the stake update is as expected
                    IStakeRegistry.OperatorStakeUpdate memory totalStakeUpdate =
                        stakeRegistry.getTotalStakeUpdateForQuorumFromIndex(i, operatorCount);

                    assertEq(totalStakeUpdate.stake, cumulativeStake);
                    assertEq(totalStakeUpdate.updateBlockNumber, cumulativeBlockNumber);
                    // make sure that the next update block number of the previous stake update is as expected
                    if (operatorCount != 0) {
                        IStakeRegistry.OperatorStakeUpdate memory prevTotalStakeUpdate =
                            stakeRegistry.getTotalStakeUpdateForQuorumFromIndex(i, operatorCount - 1);
                        assertEq(prevTotalStakeUpdate.nextUpdateBlockNumber, cumulativeBlockNumber);
                    }

                    operatorQuorumIndices[j]++;
                    operatorCount++;
                } else {
                    // make sure the number of stake updates is as expected
                    assertEq(stakeRegistry.getStakeHistoryLengthForQuorumNumber(_incrementBytes32(defaultOperatorId, j), i), 0);
                }
                cumulativeBlockNumber += blocksPassed[j];
            }   
        }
    }

    function testFuzz_DeregisterFirstOperator_Valid(
        uint256 pseudoRandomNumber,
        uint256 quorumBitmap,
        uint256 deregistrationQuorumsFlag,
        uint80[] memory stakesForQuorum,
        uint8 numOperatorsRegisterBefore
    ) public {
        // modulo so no overflow
        pseudoRandomNumber = pseudoRandomNumber % type(uint128).max;
        // register a bunch of operators
        vm.roll(100);
        uint32 cumulativeBlockNumber = 100;

        uint256 numOperators = 1 + 2*numOperatorsRegisterBefore;
        uint256[] memory quorumBitmaps = new uint256[](numOperators);

        // register
        for (uint i = 0; i < numOperatorsRegisterBefore; i++) {
            (quorumBitmaps[i],) = _registerOperatorRandomValid(_incrementAddress(defaultOperator, i), _incrementBytes32(defaultOperatorId, i), pseudoRandomNumber + i);
            
            cumulativeBlockNumber += 1;
            vm.roll(cumulativeBlockNumber);
        }

        // register the operator to be deregistered
        quorumBitmaps[numOperatorsRegisterBefore] = quorumBitmap;
        bytes32 operatorIdToDeregister = _incrementBytes32(defaultOperatorId, numOperatorsRegisterBefore);
        uint96[] memory paddedStakesForQuorum;
        {
            address operatorToDeregister = _incrementAddress(defaultOperator, numOperatorsRegisterBefore);
            paddedStakesForQuorum = _registerOperatorValid(operatorToDeregister, operatorIdToDeregister, quorumBitmap, stakesForQuorum);
        }
        // register the rest of the operators
        for (uint i = numOperatorsRegisterBefore + 1; i < 2*numOperatorsRegisterBefore; i++) {
            cumulativeBlockNumber += 1;
            vm.roll(cumulativeBlockNumber);

            (quorumBitmaps[i],) = _registerOperatorRandomValid(_incrementAddress(defaultOperator, i), _incrementBytes32(defaultOperatorId, i), pseudoRandomNumber + i);
        }

        {
            bool shouldPassBlockBeforeDeregistration  = uint256(keccak256(abi.encodePacked(pseudoRandomNumber, "shouldPassBlockBeforeDeregistration"))) & 1 == 1;
            if (shouldPassBlockBeforeDeregistration) {
                cumulativeBlockNumber += 1;
                vm.roll(cumulativeBlockNumber);
            }
        }

        // deregister the operator from a subset of the quorums
        uint256 deregistrationQuroumBitmap = quorumBitmap & deregistrationQuorumsFlag;
        _deregisterOperatorValid(operatorIdToDeregister, quorumBitmap);

        // for each bit in each quorumBitmap, increment the number of operators in that quorum
        uint32[] memory numOperatorsInQuorum = new uint32[](192);
        for (uint256 i = 0; i < quorumBitmaps.length; i++) {
            for (uint256 j = 0; j < 192; j++) {
                if (quorumBitmaps[i] >> j & 1 == 1) {
                    numOperatorsInQuorum[j]++;
                }
            }
        }

        uint8 quorumNumberIndex = 0;
        for (uint8 i = 0; i < 192; i++) {
            if (deregistrationQuroumBitmap >> i & 1 == 1) {
                // check that the operator has 2 stake updates in the quorum numbers they registered for
                assertEq(stakeRegistry.getStakeHistoryLengthForQuorumNumber(operatorIdToDeregister, i), 2);
                // make sure that the last stake update is as expected
                IStakeRegistry.OperatorStakeUpdate memory lastStakeUpdate =
                    stakeRegistry.getStakeUpdateForQuorumFromOperatorIdAndIndex(i, operatorIdToDeregister, 1);
                assertEq(lastStakeUpdate.stake, 0);
                assertEq(lastStakeUpdate.updateBlockNumber, cumulativeBlockNumber);
                assertEq(lastStakeUpdate.nextUpdateBlockNumber, 0);

                // make the analogous check for total stake history
                assertEq(stakeRegistry.getLengthOfTotalStakeHistoryForQuorum(i), numOperatorsInQuorum[i] + 1);
                // make sure that the last stake update is as expected
                IStakeRegistry.OperatorStakeUpdate memory lastTotalStakeUpdate 
                    = stakeRegistry.getTotalStakeUpdateForQuorumFromIndex(i, numOperatorsInQuorum[i]);
                assertEq(lastTotalStakeUpdate.stake, 
                    stakeRegistry.getTotalStakeUpdateForQuorumFromIndex(i, numOperatorsInQuorum[i] - 1).stake // the previous total stake
                        - paddedStakesForQuorum[quorumNumberIndex] // minus the stake that was deregistered
                );
                assertEq(lastTotalStakeUpdate.updateBlockNumber, cumulativeBlockNumber);
                assertEq(lastTotalStakeUpdate.nextUpdateBlockNumber, 0);
                quorumNumberIndex++;
            } else if (quorumBitmap >> i & 1 == 1) {
                assertEq(stakeRegistry.getStakeHistoryLengthForQuorumNumber(operatorIdToDeregister, i), 1);
                assertEq(stakeRegistry.getLengthOfTotalStakeHistoryForQuorum(i), numOperatorsInQuorum[i]);
                quorumNumberIndex++;
            } else {
                // check that the operator has 0 stake updates in the quorum numbers they did not register for
                assertEq(stakeRegistry.getStakeHistoryLengthForQuorumNumber(operatorIdToDeregister, i), 0);
            }
        }
    }
    
    function testFuzz_UpdateOperatorStake_Valid(
        uint24[] memory blocksPassed,
        uint96[] memory stakes
    ) public {
        vm.assume(blocksPassed.length > 0);
        vm.assume(blocksPassed.length <= stakes.length);
        // initialize at a non-zero block number
        uint32 intialBlockNumber = 100;
        vm.roll(intialBlockNumber);
        uint32 cumulativeBlockNumber = intialBlockNumber;
        // loop through each one of the blocks passed, roll that many blocks, set the weight in the given quorum to the stake, and trigger a stake update
        for (uint256 i = 0; i < blocksPassed.length; i++) {
            stakeRegistry.setOperatorWeight(defaultQuorumNumber, defaultOperator, stakes[i]);

            vm.expectEmit(true, true, true, true, address(stakeRegistry));
            emit StakeUpdate(defaultOperatorId, defaultQuorumNumber, stakes[i]);
            stakeRegistry.updateOperatorStake(defaultOperator, defaultOperatorId, defaultQuorumNumber);

            cumulativeBlockNumber += blocksPassed[i];
            vm.roll(cumulativeBlockNumber);
        }

        // reset for checking indices
        cumulativeBlockNumber = intialBlockNumber;
        // make sure that the stake updates are as expected
        for (uint256 i = 0; i < blocksPassed.length - 1; i++) {
            IStakeRegistry.OperatorStakeUpdate memory operatorStakeUpdate = stakeRegistry.getStakeUpdateForQuorumFromOperatorIdAndIndex(defaultQuorumNumber, defaultOperatorId, i);
            
            uint96 expectedStake = stakes[i];
            if (expectedStake < stakeRegistry.minimumStakeForQuorum(defaultQuorumNumber)) {
                expectedStake = 0;
            }

            assertEq(operatorStakeUpdate.stake, expectedStake);
            assertEq(operatorStakeUpdate.updateBlockNumber, cumulativeBlockNumber);
            cumulativeBlockNumber += blocksPassed[i];
            assertEq(operatorStakeUpdate.nextUpdateBlockNumber, cumulativeBlockNumber);
        }

        // make sure that the last stake update is as expected
        IStakeRegistry.OperatorStakeUpdate memory lastOperatorStakeUpdate = stakeRegistry.getStakeUpdateForQuorumFromOperatorIdAndIndex(defaultQuorumNumber, defaultOperatorId, blocksPassed.length - 1);
        assertEq(lastOperatorStakeUpdate.stake, stakes[blocksPassed.length - 1]);
        assertEq(lastOperatorStakeUpdate.updateBlockNumber, cumulativeBlockNumber);
        assertEq(lastOperatorStakeUpdate.nextUpdateBlockNumber, uint32(0));
    }

    function testFuzz_RecordTotalStakeUpdate_Valid(
        uint24[] memory blocksPassed,
        uint96[] memory stakes
    ) public {
        vm.assume(blocksPassed.length > 0);
        vm.assume(blocksPassed.length <= stakes.length);
        // initialize at a non-zero block number
        uint32 intialBlockNumber = 100;
        vm.roll(intialBlockNumber);
        uint32 cumulativeBlockNumber = intialBlockNumber;
        // loop through each one of the blocks passed, roll that many blocks, create an Operator Stake Update for total stake, and trigger a total stake update
        for (uint256 i = 0; i < blocksPassed.length; i++) {
            IStakeRegistry.OperatorStakeUpdate memory totalStakeUpdate;
            totalStakeUpdate.stake = stakes[i];

            stakeRegistry.recordTotalStakeUpdate(defaultQuorumNumber, totalStakeUpdate);

            cumulativeBlockNumber += blocksPassed[i];
            vm.roll(cumulativeBlockNumber);
        }

        // reset for checking indices
        cumulativeBlockNumber = intialBlockNumber;
        // make sure that the total stake updates are as expected
        for (uint256 i = 0; i < blocksPassed.length - 1; i++) {
            IStakeRegistry.OperatorStakeUpdate memory totalStakeUpdate = stakeRegistry.getTotalStakeUpdateForQuorumFromIndex(defaultQuorumNumber, i);

            assertEq(totalStakeUpdate.stake, stakes[i]);
            assertEq(totalStakeUpdate.updateBlockNumber, cumulativeBlockNumber);
            cumulativeBlockNumber += blocksPassed[i];
            assertEq(totalStakeUpdate.nextUpdateBlockNumber, cumulativeBlockNumber);
        }
    }

    // utility function for registering an operator with a valid quorumBitmap and stakesForQuorum using provided randomness
    function _registerOperatorRandomValid(
        address operator,
        bytes32 operatorId,
        uint256 psuedoRandomNumber
    ) internal returns(uint256, uint96[] memory){
        // generate uint256 quorumBitmap from psuedoRandomNumber
        uint256 quorumBitmap = uint256(keccak256(abi.encodePacked(psuedoRandomNumber, "quorumBitmap")));
        // generate uint80[] stakesForQuorum from psuedoRandomNumber
        uint80[] memory stakesForQuorum = new uint80[](BitmapUtils.countNumOnes(quorumBitmap));
        for(uint i = 0; i < stakesForQuorum.length; i++) {
            stakesForQuorum[i] = uint80(uint256(keccak256(abi.encodePacked(psuedoRandomNumber, i, "stakesForQuorum"))));
        }

        return (quorumBitmap, _registerOperatorValid(operator, operatorId, quorumBitmap, stakesForQuorum));
    }

    // utility function for registering an operator
    function _registerOperatorValid(
        address operator,
        bytes32 operatorId,
        uint256 quorumBitmap,
        uint80[] memory stakesForQuorum
    ) internal returns(uint96[] memory){
        vm.assume(quorumBitmap != 0);
        quorumBitmap = quorumBitmap & type(uint192).max;

        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        // pad the stakesForQuorum array with the minimum stake for the quorums 
        uint96[] memory paddedStakesForQuorum = new uint96[](BitmapUtils.countNumOnes(quorumBitmap));
        for(uint i = 0; i < paddedStakesForQuorum.length; i++) {
            uint96 minimumStakeForQuorum = stakeRegistry.minimumStakeForQuorum(uint8(quorumNumbers[i]));
            // make sure the operator has at least the mininmum stake in each quorum it is registering for
            if (i >= stakesForQuorum.length || stakesForQuorum[i] < minimumStakeForQuorum) {
                paddedStakesForQuorum[i] = minimumStakeForQuorum;
            } else {
                paddedStakesForQuorum[i] = stakesForQuorum[i];
            }
        }

        // set the weights of the operator
        for(uint i = 0; i < paddedStakesForQuorum.length; i++) {
            stakeRegistry.setOperatorWeight(uint8(quorumNumbers[i]), operator, paddedStakesForQuorum[i]);
        }

        // register operator
        vm.prank(address(registryCoordinatorMock));
        stakeRegistry.registerOperator(operator, operatorId, quorumNumbers);
        return paddedStakesForQuorum;
    }

    // utility function for deregistering an operator
    function _deregisterOperatorValid(
        bytes32 operatorId,
        uint256 quorumBitmap
    ) internal {
        quorumBitmap = quorumBitmap & type(uint192).max;

        bytes memory quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        // deregister operator
        vm.prank(address(registryCoordinatorMock));
        stakeRegistry.deregisterOperator(operatorId, quorumNumbers);
    }

    function _incrementAddress(address start, uint256 inc) internal pure returns(address) {
        return address(uint160(uint256(uint160(start) + inc)));
    }

    function _incrementBytes32(bytes32 start, uint256 inc) internal pure returns(bytes32) {
        return bytes32(uint256(start) + inc);
    }

}

