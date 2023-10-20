// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "src/contracts/interfaces/ISignatureUtils.sol";

import "../test/EigenLayerTestHelper.t.sol";

import "./mocks/MiddlewareRegistryMock.sol";
import "./mocks/ServiceManagerMock.sol";
import "./mocks/RegistryCoordinatorMock.sol";

import "./harnesses/StakeRegistryHarness.sol";

contract DelegationTests is EigenLayerTestHelper {
    using Math for uint256;

    uint256 public PRIVATE_KEY = 420;

    uint32 serveUntil = 100;

    // address public registryCoordinator = address(uint160(uint256(keccak256("registryCoordinator"))));
    RegistryCoordinatorMock registryCoordinator = new RegistryCoordinatorMock();

    ServiceManagerMock public serviceManager;
    StakeRegistryHarness public stakeRegistry;
    StakeRegistryHarness public stakeRegistryImplementation;
    uint8 defaultQuorumNumber = 0;
    bytes32 defaultOperatorId = bytes32(uint256(0));

    modifier fuzzedAmounts(uint256 ethAmount, uint256 eigenAmount) {
        cheats.assume(ethAmount >= 0 && ethAmount <= 1e18);
        cheats.assume(eigenAmount >= 0 && eigenAmount <= 1e18);
        _;
    }

    function setUp() public virtual override {
        EigenLayerDeployer.setUp();

        initializeMiddlewares();
    }

    function initializeMiddlewares() public {
        serviceManager = new ServiceManagerMock(slasher);

        stakeRegistry = StakeRegistryHarness(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        stakeRegistryImplementation = new StakeRegistryHarness(
            IRegistryCoordinator(registryCoordinator),
            strategyManager,
            serviceManager
        );

        {
            uint96 multiplier = 1e18;
            uint8 _NUMBER_OF_QUORUMS = 10;
            cheats.startPrank(eigenLayerProxyAdmin.owner());

            // setup the dummy minimum stake for quorum
            uint96[] memory minimumStakeForQuorum = new uint96[](_NUMBER_OF_QUORUMS);
            for (uint256 i = 0; i < minimumStakeForQuorum.length; i++) {
                minimumStakeForQuorum[i] = uint96(i + 1);
            }


            // setup the dummy quorum strategies
            IVoteWeigher.StrategyAndWeightingMultiplier[][]
                memory quorumStrategiesConsideredAndMultipliers = new IVoteWeigher.StrategyAndWeightingMultiplier[][](_NUMBER_OF_QUORUMS);
            for (uint256 i = 0; i < _NUMBER_OF_QUORUMS; ++i) {
                quorumStrategiesConsideredAndMultipliers[i] = new IVoteWeigher.StrategyAndWeightingMultiplier[](1);
                quorumStrategiesConsideredAndMultipliers[i][0] = IVoteWeigher.StrategyAndWeightingMultiplier(
                    eigenStrat,
                    multiplier
                );
            }

            // // setup the dummy quorum strategies
            // IVoteWeigher.StrategyAndWeightingMultiplier[][]
            //     memory quorumStrategiesConsideredAndMultipliers = new IVoteWeigher.StrategyAndWeightingMultiplier[][](
            //         2
            //     );
            quorumStrategiesConsideredAndMultipliers[0] = new IVoteWeigher.StrategyAndWeightingMultiplier[](3);
            quorumStrategiesConsideredAndMultipliers[0][0] = IVoteWeigher.StrategyAndWeightingMultiplier(
                stETHStrat,
                multiplier
            );
            quorumStrategiesConsideredAndMultipliers[0][1] = IVoteWeigher.StrategyAndWeightingMultiplier(
                rETHStrat,
                multiplier
            );
            quorumStrategiesConsideredAndMultipliers[0][2] = IVoteWeigher.StrategyAndWeightingMultiplier(
                cETHStrat,
                multiplier
            );

            eigenLayerProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(stakeRegistry))),
                address(stakeRegistryImplementation),
                abi.encodeWithSelector(
                    StakeRegistry.initialize.selector,
                    minimumStakeForQuorum,
                    quorumStrategiesConsideredAndMultipliers
                )
            );

            delegation.setStakeRegistry(stakeRegistry);

            cheats.stopPrank();
            assertEq(address(stakeRegistry.delegation()), address(delegation));
        }
    }

    function testDelegationMultipleStrategies(address operator, address staker) public fuzzedAddress(operator) fuzzedAddress(staker) {
        cheats.assume(staker != operator);
        // Adds strats to quorum 0, which already has stETH, rETH, cETH strats
        // total strats = numStratsToAdd + 3
        uint8 numStratsToAdd = 10;

        uint256 amountToDeposit = 1e18;
        uint96 operatorQuorum0WeightBefore = stakeRegistry.weightOfOperatorForQuorum(0, operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        _testRegisterAsOperator(operator, operatorDetails);
        _testDepositStrategies(staker, amountToDeposit, numStratsToAdd);

        // add strategies to voteWeigher
        uint96 multiplier = 1e18;
        for (uint16 i = 0; i < numStratsToAdd; ++i) {
            IVoteWeigher.StrategyAndWeightingMultiplier[]
                memory ethStratsAndMultipliers = new IVoteWeigher.StrategyAndWeightingMultiplier[](1);
            ethStratsAndMultipliers[0].strategy = strategies[i];
            ethStratsAndMultipliers[0].multiplier = multiplier;
            cheats.startPrank(stakeRegistry.serviceManager().owner());
            stakeRegistry.addStrategiesConsideredAndMultipliers(0, ethStratsAndMultipliers);
            cheats.stopPrank();
        }
        
        // Deposit to stETH strategy which is part of quorum 0
        _testDepositToStrategy(staker, amountToDeposit, stETH, stETHStrat);
        _testDepositToStrategy(staker, amountToDeposit, rETH, rETHStrat);
        _testDelegateToOperator(staker, operator);

        uint96 operatorQuorum0WeightAfter = stakeRegistry.weightOfOperatorForQuorum(0, operator);
        assertTrue(
            operatorQuorum0WeightAfter > operatorQuorum0WeightBefore,
            "testDelegation: operatorStETHWeight did not increase!"
        );
        // Test unDelegate
    }
}
