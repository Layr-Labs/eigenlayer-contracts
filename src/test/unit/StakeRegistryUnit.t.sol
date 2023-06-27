// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "../../contracts/core/Slasher.sol";
import "../../contracts/permissions/PauserRegistry.sol";
import "../../contracts/interfaces/IStrategyManager.sol";
import "../../contracts/interfaces/IStakeRegistry.sol";
import "../../contracts/interfaces/IServiceManager.sol";
import "../../contracts/interfaces/IVoteWeigher.sol";

import "../../contracts/libraries/BitmapUtils.sol";

import "../mocks/StrategyManagerMock.sol";
import "../mocks/EigenPodManagerMock.sol";
import "../mocks/ServiceManagerMock.sol";
import "../mocks/OwnableMock.sol";
import "../mocks/DelegationMock.sol";
import "../mocks/SlasherMock.sol";

import "../harnesses/StakeRegistryHarness.sol";

import "forge-std/Test.sol";

contract StakeRegistryUnitTests is Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;

    ISlasher public slasher = ISlasher(address(uint160(uint256(keccak256("slasher")))));

    Slasher public slasherImplementation;
    StakeRegistryHarness public stakeRegistryImplementation;
    StakeRegistryHarness public stakeRegistry;

    ServiceManagerMock public serviceManagerMock;
    StrategyManagerMock public strategyManagerMock;
    DelegationMock public delegationMock;
    EigenPodManagerMock public eigenPodManagerMock;

    address public serviceManagerOwner = address(uint160(uint256(keccak256("serviceManagerOwner"))));
    address public registryCoordinator = address(uint160(uint256(keccak256("registryCoordinator"))));
    address public pauser = address(uint160(uint256(keccak256("pauser"))));
    address public unpauser = address(uint160(uint256(keccak256("unpauser"))));

    address defaultOperator = address(uint160(uint256(keccak256("defaultOperator"))));
    bytes32 defaultOperatorId = keccak256("defaultOperatorId");
    uint8 defaultQuorumNumber = 0;
    uint8 numQuorums = 192;

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

        cheats.startPrank(serviceManagerOwner);
        // make the serviceManagerOwner the owner of the serviceManager contract
        serviceManagerMock = new ServiceManagerMock(slasher);
        stakeRegistryImplementation = new StakeRegistryHarness(
            IRegistryCoordinator(registryCoordinator),
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
                        quorumStrategiesConsideredAndMultipliers // initialize with 0ed out 128 quorums
                    )
                )
            )
        );

        cheats.stopPrank();
    }

    function testCorrectConstruction() public {
        // make sure the contract intializers are disabled
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        stakeRegistryImplementation.initialize(new uint96[](0), new IVoteWeigher.StrategyAndWeightingMultiplier[][](0));
    }

    function testSetMinimumStakeForQuorum_NotFromServiceManager_Reverts() public {
        cheats.expectRevert("VoteWeigherBase.onlyServiceManagerOwner: caller is not the owner of the serviceManager");
        stakeRegistry.setMinimumStakeForQuorum(defaultQuorumNumber, 0);
    }

    function testSetMinimumStakeForQuorum_Valid(uint8 quorumNumber, uint96 minimumStakeForQuorum) public {
        // set the minimum stake for quorum
        cheats.prank(serviceManagerOwner);
        stakeRegistry.setMinimumStakeForQuorum(quorumNumber, minimumStakeForQuorum);

        // make sure the minimum stake for quorum is as expected
        assertEq(stakeRegistry.minimumStakeForQuorum(quorumNumber), minimumStakeForQuorum);
    }

    function testRegisterOperator_NotFromRegistryCoordinator_Reverts() public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        cheats.expectRevert("StakeRegistry.onlyRegistryCoordinator: caller is not the RegistryCoordinator");
        stakeRegistry.registerOperator(defaultOperator, defaultOperatorId, quorumNumbers);
    }

    function testRegisterOperator_MoreQuorumsThanQuorumCount_Reverts() public {
        bytes memory quorumNumbers = new bytes(numQuorums+1);
        for (uint i = 0; i < quorumNumbers.length; i++) {
            quorumNumbers[i] = bytes1(uint8(i));
        }

        // expect that it reverts when you register
        cheats.expectRevert("StakeRegistry._registerOperator: greatest quorumNumber must be less than quorumCount");
        cheats.prank(registryCoordinator);
        stakeRegistry.registerOperator(defaultOperator, defaultOperatorId, quorumNumbers);
    }

    function testRegisterOperator_LessThanMinimumStakeForQuorum_Reverts(
        uint96[] memory stakesForQuorum
    ) public {
        cheats.assume(stakesForQuorum.length > 0);

        // set the weights of the operator
        // stakeRegistry.setOperatorWeight()

        bytes memory quorumNumbers = new bytes(stakesForQuorum.length > 128 ? 128 : stakesForQuorum.length);
        for (uint i = 0; i < quorumNumbers.length; i++) {
            quorumNumbers[i] = bytes1(uint8(i));
        }

        stakesForQuorum[stakesForQuorum.length - 1] = stakeRegistry.minimumStakeForQuorum(uint8(quorumNumbers.length - 1)) - 1;

        // expect that it reverts when you register
        cheats.expectRevert("StakeRegistry._registerOperator: Operator does not meet minimum stake requirement for quorum");
        cheats.prank(registryCoordinator);
        stakeRegistry.registerOperator(defaultOperator, defaultOperatorId, quorumNumbers);
    }

    function testFirstRegisterOperator_Valid(
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

    function testRegisterManyOperators_Valid(
        uint256[] memory quorumBitmaps,
        uint80[][] memory stakesForQuorums,
        uint24[] memory blocksPassed
    ) public {
        cheats.assume(quorumBitmaps.length > 0);
        cheats.assume(quorumBitmaps.length <= blocksPassed.length);
        cheats.assume(quorumBitmaps.length <= stakesForQuorums.length);
        cheats.assume(quorumBitmaps.length == 1);

        uint32 initialBlockNumber = 100;
        cheats.roll(initialBlockNumber);
        uint32 cumulativeBlockNumber = initialBlockNumber;

        uint96[][] memory paddedStakesForQuorums = new uint96[][](quorumBitmaps.length);
        for (uint256 i = 0; i < quorumBitmaps.length; i++) {
            emit log_named_uint("quorumBitmaps[i]", quorumBitmaps[i]);
            quorumBitmaps[i] = quorumBitmaps[i] & 0xf;
            paddedStakesForQuorums[i] = _registerOperatorValid(_incrementAddress(defaultOperator, i), _incrementBytes32(defaultOperatorId, i), quorumBitmaps[i], stakesForQuorums[i]);

            cumulativeBlockNumber += blocksPassed[i];
            cheats.roll(cumulativeBlockNumber);
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
    
    function testUpdateOperatorStake_Valid(
        uint24[] memory blocksPassed,
        uint96[] memory stakes
    ) public {
        cheats.assume(blocksPassed.length > 0);
        cheats.assume(blocksPassed.length <= stakes.length);
        // initialize at a non-zero block number
        uint32 intialBlockNumber = 100;
        cheats.roll(intialBlockNumber);
        uint32 cumulativeBlockNumber = intialBlockNumber;
        // loop through each one of the blocks passed, roll that many blocks, set the weight in the given quorum to the stake, and trigger a stake update
        for (uint256 i = 0; i < blocksPassed.length; i++) {
            stakeRegistry.setOperatorWeight(defaultQuorumNumber, defaultOperator, stakes[i]);

            cheats.expectEmit(true, true, true, true, address(stakeRegistry));
            emit StakeUpdate(defaultOperatorId, defaultQuorumNumber, stakes[i]);
            stakeRegistry.updateOperatorStake(defaultOperator, defaultOperatorId, defaultQuorumNumber);

            cumulativeBlockNumber += blocksPassed[i];
            cheats.roll(cumulativeBlockNumber);
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

    function testRecordTotalStakeUpdate_Valid(
        uint24[] memory blocksPassed,
        uint96[] memory stakes
    ) public {
        cheats.assume(blocksPassed.length > 0);
        cheats.assume(blocksPassed.length <= stakes.length);
        // initialize at a non-zero block number
        uint32 intialBlockNumber = 100;
        cheats.roll(intialBlockNumber);
        uint32 cumulativeBlockNumber = intialBlockNumber;
        // loop through each one of the blocks passed, roll that many blocks, create an Operator Stake Update for total stake, and trigger a total stake update
        for (uint256 i = 0; i < blocksPassed.length; i++) {
            IStakeRegistry.OperatorStakeUpdate memory totalStakeUpdate;
            totalStakeUpdate.stake = stakes[i];

            stakeRegistry.recordTotalStakeUpdate(defaultQuorumNumber, totalStakeUpdate);

            cumulativeBlockNumber += blocksPassed[i];
            cheats.roll(cumulativeBlockNumber);
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

    function _registerOperatorValid(
        address operator,
        bytes32 operatorId,
        uint256 quorumBitmap,
        uint80[] memory stakesForQuorum
    ) internal returns(uint96[] memory){
        cheats.assume(quorumBitmap > 0);

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
        cheats.prank(registryCoordinator);
        stakeRegistry.registerOperator(operator, operatorId, quorumNumbers);

        return paddedStakesForQuorum;
    }

    function _incrementAddress(address start, uint256 inc) internal pure returns(address) {
        return address(uint160(uint256(uint160(start) + inc)));
    }

    function _incrementBytes32(bytes32 start, uint256 inc) internal pure returns(bytes32) {
        return bytes32(uint256(start) + inc);
    }
}