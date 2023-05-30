// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "../../contracts/permissions/PauserRegistry.sol";
import "../../contracts/interfaces/IStrategyManager.sol";
import "../../contracts/interfaces/IServiceManager.sol";
import "../../contracts/interfaces/IVoteWeigher.sol";
import "../../contracts/middleware/VoteWeigherBase.sol";

import "../harnesses/VoteWeigherBaseHarness.sol";

import "../mocks/StrategyManagerMock.sol";
import "../mocks/OwnableMock.sol";
import "../mocks/DelegationMock.sol";

import "forge-std/Test.sol";

contract VoteWeigherBaseUnitTests is Test {

    Vm cheats = Vm(HEVM_ADDRESS);

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;
    IStrategyManager public strategyManager;

    address serviceManagerOwner;
    IServiceManager public serviceManager;

    DelegationMock delegationMock;

    VoteWeigherBaseHarness public voteWeigher;

    VoteWeigherBaseHarness public voteWeigherImplementation;

    address public pauser = address(555);
    address public unpauser = address(999);

    uint256 initialSupply = 1e36;
    address initialOwner = address(this);

    /// @notice emitted when a new quorum is created
    event QuorumCreated(uint8 indexed quorumNumber);
    /// @notice emitted when `strategy` has been added to the array at `strategiesConsideredAndMultipliers[quorumNumber]` with the `multiplier`
    event StrategyAddedToQuorum(uint8 indexed quorumNumber, IStrategy strategy, uint96 multiplier);
    /// @notice emitted when `strategy` has removed from the array at `strategiesConsideredAndMultipliers[quorumNumber]`
    event StrategyRemovedFromQuorum(uint8 indexed quorumNumber, IStrategy strategy);
    /// @notice emitted when `strategy` has its `multiplier` updated in the array at `strategiesConsideredAndMultipliers[quorumNumber]`
    event StrategyMultiplierUpdated(uint8 indexed quorumNumber, IStrategy strategy, uint256 multiplier);

    function setUp() virtual public {
        proxyAdmin = new ProxyAdmin();

        pauserRegistry = new PauserRegistry(pauser, unpauser);

        StrategyManagerMock strategyManagerMock = new StrategyManagerMock();
        delegationMock = new DelegationMock();
        strategyManagerMock.setAddresses(
            delegationMock,
            IEigenPodManager(address(uint160(uint256(keccak256(abi.encodePacked("eigenPodManager")))))),
            ISlasher(address(uint160(uint256(keccak256(abi.encodePacked("slasher"))))))
        );

        strategyManager = IStrategyManager(address(strategyManagerMock));

        // make the serviceManagerOwner the owner of the serviceManager contract
        cheats.prank(serviceManagerOwner);
        serviceManager = IServiceManager(address(new OwnableMock()));

        voteWeigherImplementation = new VoteWeigherBaseHarness(strategyManager, serviceManager);

        voteWeigher = VoteWeigherBaseHarness(address(new TransparentUpgradeableProxy(address(voteWeigherImplementation), address(proxyAdmin), "")));
    }

    function testCorrectConstructionParameters() public {
        assertEq(address(voteWeigherImplementation.strategyManager()), address(strategyManager));
        assertEq(address(voteWeigherImplementation.slasher()), address(strategyManager.slasher()));
        assertEq(address(voteWeigherImplementation.delegation()), address(strategyManager.delegation()));
        assertEq(address(voteWeigherImplementation.serviceManager()), address(serviceManager));
    }

    function testCreateQuorum_Valid(IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers) public {
        strategiesAndWeightingMultipliers = _convertToValidStrategiesAndWeightingMultipliers(strategiesAndWeightingMultipliers);
        // create a quorum from the serviceManagerOwner
        // get the quorum count before the quorum is created
        uint8 quorumCountBefore = voteWeigher.quorumCount();
        cheats.prank(serviceManagerOwner);
        // expect each strategy to be added to the quorum
        for (uint i = 0; i < strategiesAndWeightingMultipliers.length; i++) {
            cheats.expectEmit(true, true, true, true, address(voteWeigher));
            emit StrategyAddedToQuorum(quorumCountBefore, strategiesAndWeightingMultipliers[i].strategy, strategiesAndWeightingMultipliers[i].multiplier);
        }
        // created quorum will have quorum number of the count before it was created
        cheats.expectEmit(true, true, true, true, address(voteWeigher));
        emit QuorumCreated(quorumCountBefore);
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers);
        
        assertEq(voteWeigher.quorumCount(), quorumCountBefore + 1);
        // check that all of the weights are correct
        for (uint i = 0; i < strategiesAndWeightingMultipliers.length; i++) {
            IVoteWeigher.StrategyAndWeightingMultiplier memory strategyAndWeightingMultiplier = voteWeigher.strategyAndWeightingMultiplierForQuorumByIndex(quorumCountBefore, i);
            assertEq(address(strategyAndWeightingMultiplier.strategy), address(strategiesAndWeightingMultipliers[i].strategy));
            assertEq(strategyAndWeightingMultiplier.multiplier, strategiesAndWeightingMultipliers[i].multiplier);
        }
    }

    function testCreateQuorum_FromNotServiceManagerOwner_Reverts(
        address notServiceManagerOwner,
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers
    ) public {
        cheats.assume(notServiceManagerOwner != serviceManagerOwner);
        cheats.prank(notServiceManagerOwner);
        cheats.expectRevert("VoteWeigherBase.onlyServiceManagerOwner: caller is not the owner of the serviceManager");
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers);
    }

    function testCreateQuorum_StrategiesAndWeightingMultipliers_LengthGreaterThanMaxAllowed_Reverts(
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers
    ) public {
        strategiesAndWeightingMultipliers = _removeDuplicates(strategiesAndWeightingMultipliers);
        _assumeNoZeroWeights(strategiesAndWeightingMultipliers);

        cheats.assume(strategiesAndWeightingMultipliers.length > voteWeigher.getMaxWeighingFunctionLength());
        cheats.prank(serviceManagerOwner);
        cheats.expectRevert("VoteWeigherBase._addStrategiesConsideredAndMultipliers: exceed MAX_WEIGHING_FUNCTION_LENGTH");
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers);
    }

    function testCreateQuorum_StrategiesAndWeightingMultipliers_WithDuplicateStrategies_Reverts(
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers
    ) public {
        cheats.assume(strategiesAndWeightingMultipliers.length <= voteWeigher.getMaxWeighingFunctionLength());
        cheats.assume(strategiesAndWeightingMultipliers.length > 1);
        _assumeNoZeroWeights(strategiesAndWeightingMultipliers);

        // make sure a duplicate strategy exists
        bool duplicateExists;
        for (uint256 i = 0; i < strategiesAndWeightingMultipliers.length; i++) {
            for (uint256 j = i + 1; j < strategiesAndWeightingMultipliers.length; j++) {
                if (strategiesAndWeightingMultipliers[i].strategy == strategiesAndWeightingMultipliers[j].strategy) {
                    duplicateExists = true;
                    break;
                }
            }
        }
        cheats.assume(duplicateExists);

        cheats.prank(serviceManagerOwner);
        cheats.expectRevert("VoteWeigherBase._addStrategiesConsideredAndMultipliers: cannot add same strategy 2x");
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers);
    }

    function testCreateQuorum_EmptyStrategiesAndWeightingMultipliers_Reverts() public {
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers;
        cheats.prank(serviceManagerOwner);
        cheats.expectRevert("VoteWeigherBase._addStrategiesConsideredAndMultipliers: no strategies provided");
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers);
    }

    function testCreateQuorum_StrategiesAndWeightingMultipliers_WithZeroWeight(
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers
    ) public {
        strategiesAndWeightingMultipliers = _removeDuplicates(strategiesAndWeightingMultipliers);
        cheats.assume(strategiesAndWeightingMultipliers.length <= voteWeigher.getMaxWeighingFunctionLength());
        cheats.assume(strategiesAndWeightingMultipliers.length > 0);
        // make sure a zero weight exists
        bool zeroWeightExists;
        for (uint256 i = 0; i < strategiesAndWeightingMultipliers.length; i++) {
            if (strategiesAndWeightingMultipliers[i].multiplier == 0) {
                zeroWeightExists = true;
                break;
            }
        }
        cheats.assume(zeroWeightExists);

        cheats.prank(serviceManagerOwner);
        cheats.expectRevert("VoteWeigherBase._addStrategiesConsideredAndMultipliers: cannot add strategy with zero weight");
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers);
    }

    function testCreateQuorum_MoreThan256Quorums_Reverts(
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers
    ) public {
        strategiesAndWeightingMultipliers = _convertToValidStrategiesAndWeightingMultipliers(strategiesAndWeightingMultipliers);
            
        cheats.startPrank(serviceManagerOwner);
        for (uint i = 0; i < 255; i++) {
            voteWeigher.createQuorum(strategiesAndWeightingMultipliers);
        }
        cheats.expectRevert("VoteWeigherBase._createQuorum: number of quorums cannot 256");
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers); 
    }

    function testAddStrategiesConsideredAndMultipliers_Valid(
        uint256 randomSplit,
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers
    ) public {
        strategiesAndWeightingMultipliers = _convertToValidStrategiesAndWeightingMultipliers(strategiesAndWeightingMultipliers);
        // make sure there is at least 2 strategies
        cheats.assume(strategiesAndWeightingMultipliers.length > 1);
        // we need at least 1 strategy in each side of the split
        randomSplit = randomSplit % (strategiesAndWeightingMultipliers.length - 1) + 1;
        // create 2 arrays, 1 with the first randomSplit elements and 1 with the rest
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers1 = new IVoteWeigher.StrategyAndWeightingMultiplier[](randomSplit);
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers2 = new IVoteWeigher.StrategyAndWeightingMultiplier[](strategiesAndWeightingMultipliers.length - randomSplit);
        for (uint256 i = 0; i < strategiesAndWeightingMultipliers.length; i++) {
            if (i < randomSplit) {
                strategiesAndWeightingMultipliers1[i] = strategiesAndWeightingMultipliers[i];
            } else {
                strategiesAndWeightingMultipliers2[i - randomSplit] = strategiesAndWeightingMultipliers[i];
            }
        }
        uint8 quorumNumber = voteWeigher.quorumCount();
        // create quorum with the first randomSplit elements
        cheats.startPrank(serviceManagerOwner);
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers1);

        // add the rest of the strategies
        for (uint i = 0; i < strategiesAndWeightingMultipliers2.length; i++) {
            cheats.expectEmit(true, true, true, true, address(voteWeigher));
            emit StrategyAddedToQuorum(quorumNumber, strategiesAndWeightingMultipliers2[i].strategy, strategiesAndWeightingMultipliers2[i].multiplier);
        }
        voteWeigher.addStrategiesConsideredAndMultipliers(quorumNumber, strategiesAndWeightingMultipliers2);

        // check that the quorum was created and strategies were added correctly
        for (uint i = 0; i < strategiesAndWeightingMultipliers.length; i++) {
            IVoteWeigher.StrategyAndWeightingMultiplier memory strategyAndWeightingMultiplier = voteWeigher.strategyAndWeightingMultiplierForQuorumByIndex(quorumNumber, i);
            assertEq(address(strategyAndWeightingMultiplier.strategy), address(strategiesAndWeightingMultipliers[i].strategy));
            assertEq(strategyAndWeightingMultiplier.multiplier, strategiesAndWeightingMultipliers[i].multiplier);
        }
    }

    function testAddStrategiesConsideredAndMultipliers_NotFromServiceManagerOwner_Reverts(
        address notServiceManagerOwner,
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers
    ) public {
        cheats.assume(notServiceManagerOwner != serviceManagerOwner);
        strategiesAndWeightingMultipliers = _convertToValidStrategiesAndWeightingMultipliers(strategiesAndWeightingMultipliers);
        // we need 2 or more strategies to create the quorums and then add one
        cheats.assume(strategiesAndWeightingMultipliers.length > 1);

        // separate strategiesAndWeightingMultipliers into 2 arrays, 1 with the last element and 1 with the rest
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers1 = new IVoteWeigher.StrategyAndWeightingMultiplier[](strategiesAndWeightingMultipliers.length - 1);
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers2 = new IVoteWeigher.StrategyAndWeightingMultiplier[](1);
        for (uint256 i = 0; i < strategiesAndWeightingMultipliers.length - 1; i++) {
            strategiesAndWeightingMultipliers1[i] = strategiesAndWeightingMultipliers[i];
        }
        strategiesAndWeightingMultipliers2[0] = strategiesAndWeightingMultipliers[strategiesAndWeightingMultipliers.length - 1];

        // create quorum with all but the last element
        uint8 quorumNumber = voteWeigher.quorumCount();
        cheats.prank(serviceManagerOwner);
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers1);

        // add the last element
        cheats.prank(notServiceManagerOwner);
        cheats.expectRevert("VoteWeigherBase.onlyServiceManagerOwner: caller is not the owner of the serviceManager");
        voteWeigher.addStrategiesConsideredAndMultipliers(quorumNumber, strategiesAndWeightingMultipliers2);
    }

    function testAddStrategiesConsideredAndMultipliers_ForNonExistantQuorum_Reverts(
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers
    ) public {
        strategiesAndWeightingMultipliers = _convertToValidStrategiesAndWeightingMultipliers(strategiesAndWeightingMultipliers);
        // we need 2 or more strategies to create the quorums and then add one
        cheats.assume(strategiesAndWeightingMultipliers.length > 1);

        // separate strategiesAndWeightingMultipliers into 2 arrays, 1 with the last element and 1 with the rest
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers1 = new IVoteWeigher.StrategyAndWeightingMultiplier[](strategiesAndWeightingMultipliers.length - 1);
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers2 = new IVoteWeigher.StrategyAndWeightingMultiplier[](1);
        for (uint256 i = 0; i < strategiesAndWeightingMultipliers.length - 1; i++) {
            strategiesAndWeightingMultipliers1[i] = strategiesAndWeightingMultipliers[i];
        }
        strategiesAndWeightingMultipliers2[0] = strategiesAndWeightingMultipliers[strategiesAndWeightingMultipliers.length - 1];

        // create quorum with all but the last element
        uint8 quorumNumber = voteWeigher.quorumCount();
        cheats.startPrank(serviceManagerOwner);
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers1);

        // add the last element
        cheats.expectRevert("VoteWeigherBase._addStrategiesConsideredAndMultipliers: quorumNumber exceeds quorumCount");
        voteWeigher.addStrategiesConsideredAndMultipliers(quorumNumber+1, strategiesAndWeightingMultipliers2);        
    }

    function testRemoveStrategiesConsideredAndMultipliers_Valid(
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers,
        uint256[] memory indicesToRemove
    ) public {
        strategiesAndWeightingMultipliers = _convertToValidStrategiesAndWeightingMultipliers(strategiesAndWeightingMultipliers);
        // take indices modulo length
        for (uint256 i = 0; i < indicesToRemove.length; i++) {
            indicesToRemove[i] = indicesToRemove[i] % strategiesAndWeightingMultipliers.length;
        }
        indicesToRemove = _removeDuplicatesUint256(indicesToRemove);
        cheats.assume(indicesToRemove.length > 0);
        cheats.assume(indicesToRemove.length < strategiesAndWeightingMultipliers.length);
        indicesToRemove = _sortInDescendingOrder(indicesToRemove);

        // create the quorum
        uint8 quorumNumber = voteWeigher.quorumCount();
        cheats.startPrank(serviceManagerOwner);
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers);

        // remove certain strategies
        // make sure events are emmitted
        for (uint i = 0; i < indicesToRemove.length; i++) {
            cheats.expectEmit(true, true, true, true, address(voteWeigher));
            emit StrategyRemovedFromQuorum(quorumNumber, strategiesAndWeightingMultipliers[indicesToRemove[i]].strategy);
        }
        voteWeigher.removeStrategiesConsideredAndMultipliers(quorumNumber, indicesToRemove);

        // check that the strategies were removed
        
        // check that the strategies that were not removed are still there
        // get all strategies and multipliers form the contracts
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliersFromContract = new IVoteWeigher.StrategyAndWeightingMultiplier[](voteWeigher.strategiesConsideredAndMultipliersLength(quorumNumber));
        for (uint256 i = 0; i < strategiesAndWeightingMultipliersFromContract.length; i++) {
            strategiesAndWeightingMultipliersFromContract[i] = voteWeigher.strategyAndWeightingMultiplierForQuorumByIndex(quorumNumber, i);
        }

        // remove indicesToRemove from local strategiesAndWeightingMultipliers
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliersLocal = new IVoteWeigher.StrategyAndWeightingMultiplier[](strategiesAndWeightingMultipliers.length - indicesToRemove.length);
        uint256 j = 0;
        for (uint256 i = 0; i < strategiesAndWeightingMultipliers.length; i++) {
            for (uint256 k = 0; k < indicesToRemove.length; k++) {
                if (i == indicesToRemove[k]) {
                    break;
                }
                if (k == indicesToRemove.length - 1) {
                    strategiesAndWeightingMultipliersLocal[j] = strategiesAndWeightingMultipliers[i];
                    j++;
                }
            }
        }

        // sort both arrays by strategy so that they are easily comaparable
        strategiesAndWeightingMultipliersFromContract = _sortStrategiesAndWeightingMultipliersByStrategy(strategiesAndWeightingMultipliersFromContract);
        strategiesAndWeightingMultipliersLocal = _sortStrategiesAndWeightingMultipliersByStrategy(strategiesAndWeightingMultipliersLocal);

        // check that the arrays are the same
        assertEq(strategiesAndWeightingMultipliersFromContract.length, strategiesAndWeightingMultipliersLocal.length);
        for (uint256 i = 0; i < strategiesAndWeightingMultipliersFromContract.length; i++) {
            assertEq(address(strategiesAndWeightingMultipliersFromContract[i].strategy), address(strategiesAndWeightingMultipliersLocal[i].strategy));
            assertEq(strategiesAndWeightingMultipliersFromContract[i].multiplier, strategiesAndWeightingMultipliersLocal[i].multiplier);
        }

    }

    function testRemoveStrategiesConsideredAndMultipliers_NotFromServiceManagerOwner_Reverts(
        address notServiceManagerOwner,
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers,
        uint256[] memory indicesToRemove
    ) public {
        cheats.assume(notServiceManagerOwner != serviceManagerOwner);
        strategiesAndWeightingMultipliers = _convertToValidStrategiesAndWeightingMultipliers(strategiesAndWeightingMultipliers);

        // create a valid quorum
        uint8 quorumNumber = voteWeigher.quorumCount();
        cheats.prank(serviceManagerOwner);
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers);

        // remove certain strategies
        cheats.prank(notServiceManagerOwner);
        cheats.expectRevert("VoteWeigherBase.onlyServiceManagerOwner: caller is not the owner of the serviceManager");
        voteWeigher.removeStrategiesConsideredAndMultipliers(quorumNumber, indicesToRemove);
    }

    function testRemoveStrategiesConsideredAndMultipliers_ForNonExistantQuorum_Reverts(
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers,
        uint256[] memory indicesToRemove
    ) public {
        strategiesAndWeightingMultipliers = _convertToValidStrategiesAndWeightingMultipliers(strategiesAndWeightingMultipliers);

        // create a valid quorum
        uint8 quorumNumber = voteWeigher.quorumCount();
        cheats.startPrank(serviceManagerOwner);
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers);

        // remove strategies from a non-existant quorum
        cheats.expectRevert("VoteWeigherBase.removeStrategiesConsideredAndMultipliers: quorumNumber is greater than or equal to quorumCount");
        voteWeigher.removeStrategiesConsideredAndMultipliers(quorumNumber + 1, indicesToRemove);
    }

    function testRemoveStrategiesConsideredAndMultipliers_EmptyIndicesToRemove_HasNoEffect(
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers
    ) public {
        strategiesAndWeightingMultipliers = _convertToValidStrategiesAndWeightingMultipliers(strategiesAndWeightingMultipliers);

        // create a valid quorum
        uint8 quorumNumber = voteWeigher.quorumCount();
        cheats.startPrank(serviceManagerOwner);
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers);

        // remove no strategies
        voteWeigher.removeStrategiesConsideredAndMultipliers(quorumNumber, new uint256[](0));
        // make sure the quorum strategies and weights haven't changed
        for (uint i = 0; i < strategiesAndWeightingMultipliers.length; i++) {
            IVoteWeigher.StrategyAndWeightingMultiplier memory strategyAndWeightingMultiplier = voteWeigher.strategyAndWeightingMultiplierForQuorumByIndex(quorumNumber, i);
            assertEq(address(strategyAndWeightingMultiplier.strategy), address(strategiesAndWeightingMultipliers[i].strategy));
            assertEq(strategyAndWeightingMultiplier.multiplier, strategiesAndWeightingMultipliers[i].multiplier);
        }
    }

    function testModifyStrategyWeights_NotFromServiceManagerOwner_Reverts(
        address notServiceManagerOwner,
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers,
        uint256[] memory strategyIndices,
        uint96[] memory newWeights
    ) public {
        cheats.assume(notServiceManagerOwner != serviceManagerOwner);
        strategiesAndWeightingMultipliers = _convertToValidStrategiesAndWeightingMultipliers(strategiesAndWeightingMultipliers);

        // create a valid quorum
        uint8 quorumNumber = voteWeigher.quorumCount();
        cheats.prank(serviceManagerOwner);
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers);

        // modify certain strategies
        cheats.prank(notServiceManagerOwner);
        cheats.expectRevert("VoteWeigherBase.onlyServiceManagerOwner: caller is not the owner of the serviceManager");
        voteWeigher.modifyStrategyWeights(quorumNumber, strategyIndices, newWeights);
    }

    function testModifyStrategyWeights_Valid(
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers,
        uint256[] memory strategyIndices,
        uint96[] memory newWeights
    ) public {
        strategiesAndWeightingMultipliers = _convertToValidStrategiesAndWeightingMultipliers(strategiesAndWeightingMultipliers);

        // take indices modulo length
        for (uint256 i = 0; i < strategyIndices.length; i++) {
            strategyIndices[i] = strategyIndices[i] % strategiesAndWeightingMultipliers.length;
        }
        strategyIndices = _removeDuplicatesUint256(strategyIndices);
        cheats.assume(strategyIndices.length > 0);
        cheats.assume(strategyIndices.length < strategiesAndWeightingMultipliers.length);

        // trim the provided weights to the length of the strategyIndices
        uint96[] memory newWeightsTrim = new uint96[](strategyIndices.length);
        for (uint256 i = 0; i < strategyIndices.length; i++) {
            if(i < newWeights.length) {
                newWeightsTrim[i] = newWeights[i];
            } else {
                newWeightsTrim[i] = strategiesAndWeightingMultipliers[strategyIndices[i]].multiplier - 1;
            }
        }
        newWeights = newWeightsTrim;

        // create a valid quorum
        uint8 quorumNumber = voteWeigher.quorumCount();
        cheats.startPrank(serviceManagerOwner);
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers);

        // modify certain strategies
        for (uint i = 0; i < strategyIndices.length; i++) {
            cheats.expectEmit(true, true, true, true, address(voteWeigher));
            emit StrategyMultiplierUpdated(quorumNumber, strategiesAndWeightingMultipliers[strategyIndices[i]].strategy, newWeights[i]);
        }
        voteWeigher.modifyStrategyWeights(quorumNumber, strategyIndices, newWeights);

        // convert the strategies and weighting multipliers to the modified
        for (uint i = 0; i < strategyIndices.length; i++) {
            strategiesAndWeightingMultipliers[strategyIndices[i]].multiplier = newWeights[i];
        }
        // make sure the quorum strategies and weights have changed
        for (uint i = 0; i < strategiesAndWeightingMultipliers.length; i++) {
            IVoteWeigher.StrategyAndWeightingMultiplier memory strategyAndWeightingMultiplier = voteWeigher.strategyAndWeightingMultiplierForQuorumByIndex(quorumNumber, i);
            assertEq(address(strategyAndWeightingMultiplier.strategy), address(strategiesAndWeightingMultipliers[i].strategy));
            assertEq(strategyAndWeightingMultiplier.multiplier, strategiesAndWeightingMultipliers[i].multiplier);
        }
    }

    function testModifyStrategyWeights_ForNonExistantQuorum_Reverts(
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers,
        uint256[] memory strategyIndices,
        uint96[] memory newWeights
    ) public {
        strategiesAndWeightingMultipliers = _convertToValidStrategiesAndWeightingMultipliers(strategiesAndWeightingMultipliers);

        // create a valid quorum
        uint8 quorumNumber = voteWeigher.quorumCount();
        cheats.startPrank(serviceManagerOwner);
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers);

        // modify certain strategies of a non-existant quorum
        cheats.expectRevert("VoteWeigherBase.modifyStrategyWeights: quorumNumber is greater than or equal to quorumCount");
        voteWeigher.modifyStrategyWeights(quorumNumber + 1, strategyIndices, newWeights);
    }

    function testModifyStrategyWeights_InconsistentStrategyAndWeightArrayLengths_Reverts(
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers,
        uint256[] memory strategyIndices,
        uint96[] memory newWeights
    ) public {
        strategiesAndWeightingMultipliers = _convertToValidStrategiesAndWeightingMultipliers(strategiesAndWeightingMultipliers);

        // make sure the arrays are of different lengths
        cheats.assume(strategyIndices.length != newWeights.length);

        // create a valid quorum
        uint8 quorumNumber = voteWeigher.quorumCount();
        cheats.startPrank(serviceManagerOwner);
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers);

        // modify certain strategies
        cheats.expectRevert("VoteWeigherBase.modifyStrategyWeights: input length mismatch");
        voteWeigher.modifyStrategyWeights(quorumNumber, strategyIndices, newWeights);
    }

    function testModifyStrategyWeights_EmptyStrategyIndicesAndWeights_HasNoEffect(
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers
    ) public {
        strategiesAndWeightingMultipliers = _convertToValidStrategiesAndWeightingMultipliers(strategiesAndWeightingMultipliers);

        // create a valid quorum
        uint8 quorumNumber = voteWeigher.quorumCount();
        cheats.startPrank(serviceManagerOwner);
        voteWeigher.createQuorum(strategiesAndWeightingMultipliers);

        // modify no strategies
        voteWeigher.modifyStrategyWeights(quorumNumber, new uint256[](0), new uint96[](0));
        // make sure the quorum strategies and weights haven't changed
        for (uint i = 0; i < strategiesAndWeightingMultipliers.length; i++) {
            IVoteWeigher.StrategyAndWeightingMultiplier memory strategyAndWeightingMultiplier = voteWeigher.strategyAndWeightingMultiplierForQuorumByIndex(quorumNumber, i);
            assertEq(address(strategyAndWeightingMultiplier.strategy), address(strategiesAndWeightingMultipliers[i].strategy));
            assertEq(strategyAndWeightingMultiplier.multiplier, strategiesAndWeightingMultipliers[i].multiplier);
        }
    }

    function testWeightOfOperator(
        address operator,
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndMultipliers,
        uint96[] memory shares
    ) public {
        strategiesAndMultipliers = _convertToValidStrategiesAndWeightingMultipliers(strategiesAndMultipliers);
        cheats.assume(shares.length >= strategiesAndMultipliers.length);
        for (uint i = 0; i < strategiesAndMultipliers.length; i++) {
            cheats.assume(uint256(shares[i]) * uint256(strategiesAndMultipliers[i].multiplier) <= type(uint96).max);
        }

        // set the operator shares
        for (uint i = 0; i < strategiesAndMultipliers.length; i++) {
            delegationMock.setOperatorShares(operator, strategiesAndMultipliers[i].strategy, shares[i]);
        }

        // create a valid quorum
        uint8 quorumNumber = voteWeigher.quorumCount();
        cheats.startPrank(serviceManagerOwner);
        voteWeigher.createQuorum(strategiesAndMultipliers);

        // make sure the weight of the operator is correct
        uint256 expectedWeight = 0;
        for (uint i = 0; i < strategiesAndMultipliers.length; i++) {
            
            expectedWeight += shares[i] * strategiesAndMultipliers[i].multiplier / voteWeigher.getWeightingDivisor();
        }

        assertEq(voteWeigher.weightOfOperator(quorumNumber, operator), expectedWeight);
    }

    function _removeDuplicates(IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers) 
        internal pure 
        returns(IVoteWeigher.StrategyAndWeightingMultiplier[] memory)
    {
        IVoteWeigher.StrategyAndWeightingMultiplier[] memory deduplicatedStrategiesAndWeightingMultipliers = new IVoteWeigher.StrategyAndWeightingMultiplier[](strategiesAndWeightingMultipliers.length);
        uint256 numUniqueStrategies = 0;
        for (uint i = 0; i < strategiesAndWeightingMultipliers.length; i++) {
            bool isDuplicate = false;
            for (uint j = 0; j < deduplicatedStrategiesAndWeightingMultipliers.length; j++) {
                if (strategiesAndWeightingMultipliers[i].strategy == deduplicatedStrategiesAndWeightingMultipliers[j].strategy) {
                    isDuplicate = true;
                    break;
                }
            }
            if(!isDuplicate) {
                deduplicatedStrategiesAndWeightingMultipliers[numUniqueStrategies] = strategiesAndWeightingMultipliers[i];
                numUniqueStrategies++;
            }
        }

        IVoteWeigher.StrategyAndWeightingMultiplier[] memory trimmedStrategiesAndWeightingMultipliers = new IVoteWeigher.StrategyAndWeightingMultiplier[](numUniqueStrategies);
        for (uint i = 0; i < numUniqueStrategies; i++) {
            trimmedStrategiesAndWeightingMultipliers[i] = deduplicatedStrategiesAndWeightingMultipliers[i];
        }
        return trimmedStrategiesAndWeightingMultipliers;
    }

    function _assumeNoZeroWeights(IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers) internal view {
        for (uint256 i = 0; i < strategiesAndWeightingMultipliers.length; i++) {
            cheats.assume(strategiesAndWeightingMultipliers[i].multiplier != 0);
        }
    }

    function _sortStrategiesAndWeightingMultipliersByStrategy(IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers) 
        internal pure returns(IVoteWeigher.StrategyAndWeightingMultiplier[] memory) 
    {
        uint256 n = strategiesAndWeightingMultipliers.length;
        for (uint256 i = 0; i < n - 1; i++) {
            uint256 min_idx = i;
            for (uint256 j = i + 1; j < n; j++) {
                if (uint160(address(strategiesAndWeightingMultipliers[j].strategy)) < uint160(address(strategiesAndWeightingMultipliers[min_idx].strategy))) {
                    min_idx = j;
                }
            }
            IVoteWeigher.StrategyAndWeightingMultiplier memory temp = strategiesAndWeightingMultipliers[min_idx];
            strategiesAndWeightingMultipliers[min_idx] = strategiesAndWeightingMultipliers[i];
            strategiesAndWeightingMultipliers[i] = temp;
        }
        return strategiesAndWeightingMultipliers;
    }

    function _sortInDescendingOrder(uint256[] memory arr) internal pure returns(uint256[] memory) {
        uint256 n = arr.length;
        for (uint256 i = 0; i < n - 1; i++) {
            uint256 max_idx = i;
            for (uint256 j = i + 1; j < n; j++) {
                if (arr[j] > arr[max_idx]) {
                    max_idx = j;
                }
            }
            uint256 temp = arr[max_idx];
            arr[max_idx] = arr[i];
            arr[i] = temp;
        }
        return arr;
    }

    function _removeDuplicatesUint256(uint256[] memory arr) internal pure returns(uint256[] memory) {
        uint256[] memory deduplicatedArr = new uint256[](arr.length);
        uint256 numUniqueElements = 0;
        for (uint i = 0; i < arr.length; i++) {
            bool isDuplicate = false;
            for (uint j = 0; j < deduplicatedArr.length; j++) {
                if (arr[i] == deduplicatedArr[j]) {
                    isDuplicate = true;
                    break;
                }
            }
            if(!isDuplicate) {
                deduplicatedArr[numUniqueElements] = arr[i];
                numUniqueElements++;
            }
        }

        uint256[] memory trimmedArr = new uint256[](numUniqueElements);
        for (uint i = 0; i < numUniqueElements; i++) {
            trimmedArr[i] = deduplicatedArr[i];
        }
        return trimmedArr;
    }

    function _convertToValidStrategiesAndWeightingMultipliers(IVoteWeigher.StrategyAndWeightingMultiplier[] memory strategiesAndWeightingMultipliers) internal view returns (IVoteWeigher.StrategyAndWeightingMultiplier[] memory) {
        strategiesAndWeightingMultipliers = _removeDuplicates(strategiesAndWeightingMultipliers);
        cheats.assume(strategiesAndWeightingMultipliers.length <= voteWeigher.getMaxWeighingFunctionLength());
        cheats.assume(strategiesAndWeightingMultipliers.length > 0);
        _assumeNoZeroWeights(strategiesAndWeightingMultipliers);
        return strategiesAndWeightingMultipliers;
    }
}