// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// Contracts
import "../../src/contracts/core/DelegationManager.sol";
import "../../src/contracts/core/StrategyManager.sol";
import "../../src/contracts/core/AVSDirectory.sol";
import "../../src/contracts/core/AllocationManager.sol";
import "../../src/contracts/strategies/StrategyBase.sol";

// Test
import "forge-std/Test.sol";

/// @notice Tests deployed contracts as part of the public devnet
/// Run with: forge test --mc Devnet_Lifecycle_Test --rpc-url $RPC_HOLESKY
contract Devnet_Lifecycle_Test is Test {
    
    // Contracts
    DelegationManager public delegationManager;
    StrategyManager public strategyManager;
    AVSDirectory public avsDirectory;
    AllocationManager public allocationManager;
    StrategyBase public wethStrategy;
    IERC20 public weth;

    Vm cheats = Vm(VM_ADDRESS);

    // Addresses
    address public staker = address(0x1);
    address public operator;
    uint256 operatorPk = 420;
    address public avs = address(0x3);
    uint32 public operatorSetId = 1;
    uint256 public wethAmount = 100 ether;
    uint256 public wethShares = 100 ether;
    OperatorSet public operatorSet;

    // Values
    uint64 public magnitudeToSet = 1e18;

    function setUp() public {
        // Set contracts
        delegationManager = DelegationManager(0x3391eBafDD4b2e84Eeecf1711Ff9FC06EF9Ed182);
        strategyManager = StrategyManager(0x70f8bC2Da145b434de66114ac539c9756eF64fb3);
        avsDirectory = AVSDirectory(0xCa839541648D3e23137457b1Fd4A06bccEADD33a);
        allocationManager = AllocationManager(0xAbD5Dd30CaEF8598d4EadFE7D45Fd582EDEade15);
        wethStrategy = StrategyBase(0x4f812633943022fA97cb0881683aAf9f318D5Caa);
        weth = IERC20(0x94373a4919B3240D86eA41593D5eBa789FEF3848);

        // Set operator
        operator = cheats.addr(operatorPk);
        operatorSet = OperatorSet({avs: avs, id: operatorSetId});
    }

    function _getOperatorSetArray() internal view returns (uint32[] memory) {
        uint32[] memory operatorSets = new uint32[](1);
        operatorSets[0] = operatorSetId;
        return operatorSets;
    }
    
    function _getOperatorSetsArray() internal view returns (OperatorSet[] memory) {
        OperatorSet[] memory operatorSets = new OperatorSet[](1);
        operatorSets[0] = OperatorSet({avs: avs, id: operatorSetId});
        return operatorSets;
    }

    function test() public {
        if (block.chainid == 17000) {
            // Seed staker with WETH
            StdCheats.deal(address(weth), address(staker), wethAmount);
            _run_lifecycle();
        }
    }

    function _run_lifecycle() internal {
        // Staker <> Operator Relationship
        _depositIntoStrategy();
        _registerOperator();
        _delegateToOperator();

        // Operator <> AVS Relationship
        _registerAVS();
        _registerOperatorToAVS();
        _setMagnitude();

        // Slash operator
        _slashOperator();

        // Withdraw staker
        _withdrawStaker();
    }

    function _depositIntoStrategy() internal {
        // Approve WETH
        cheats.startPrank(staker);
        weth.approve(address(strategyManager), wethAmount);

        // Deposit WETH into strategy
        strategyManager.depositIntoStrategy(wethStrategy, weth, wethAmount);
        cheats.stopPrank();

        // Check staker balance
        assertEq(weth.balanceOf(staker), 0);

        // Check staker shares
        assertEq(strategyManager.stakerDepositShares(staker, wethStrategy), wethAmount);
    }

    function _registerOperator() internal {
        // Register operator
        string memory emptyStringForMetadataURI;
        cheats.prank(operator);
        delegationManager.registerAsOperator(address(0), 1, emptyStringForMetadataURI);
        // Warp passed configuration delay
        cheats.roll(block.number + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS());

        // Validate storage
        assertTrue(delegationManager.isOperator(operator));
    }

    function _delegateToOperator() internal {
        // Delegate to operator
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry;
        cheats.prank(staker);
        delegationManager.delegateTo(operator, signatureWithExpiry, bytes32(0));

        // Validate storage
        assertTrue(delegationManager.isDelegated(staker));
        assertEq(delegationManager.delegatedTo(staker), operator);

        // Validate operator shares
        assertEq(delegationManager.operatorShares(operator, wethStrategy), wethShares);
    }

    function _registerAVS() internal {
        cheats.startPrank(avs);
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = wethStrategy;

        IAllocationManagerTypes.CreateSetParams memory createSetParams = IAllocationManagerTypes.CreateSetParams({
            operatorSetId: operatorSetId,
            strategies: strategies
        });

        IAllocationManagerTypes.CreateSetParams[] memory array = new IAllocationManagerTypes.CreateSetParams[](1);
        array[0] = createSetParams;

        allocationManager.createOperatorSets(avs, array);
        cheats.stopPrank();
    }

    function _registerOperatorToAVS() public {
        cheats.prank(operator);
        
        uint32[] memory operatorSetIds = new uint32[](1);
        operatorSetIds[0] = operatorSetId;

        allocationManager.registerForOperatorSets(operator, IAllocationManagerTypes.RegisterParams(avs, operatorSetIds, ""));

        assertEq(allocationManager.getMembers(OperatorSet(avs, operatorSetId))[0], operator);
    }

    function _setMagnitude() public {
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = wethStrategy;

        uint64[] memory magnitudes = new uint64[](1);
        magnitudes[0] = magnitudeToSet;

        IAllocationManagerTypes.AllocateParams[] memory allocations = new IAllocationManagerTypes.AllocateParams[](1);
        allocations[0] = IAllocationManagerTypes.AllocateParams({
            operatorSet: operatorSet,
            strategies: strategies,
            newMagnitudes: magnitudes
        });

        cheats.prank(operator);
        allocationManager.modifyAllocations(operator, allocations);

        // Assert storage
        IAllocationManagerTypes.Allocation memory info = allocationManager.getAllocation(operator, operatorSet, wethStrategy);
        assertEq(info.currentMagnitude, 0);
        assertEq(info.pendingDiff, int128(uint128(magnitudeToSet)));
        assertEq(info.effectBlock, block.number + 1);

        // Warp to effect block
        cheats.roll(block.number + 1);

        // Check allocation
        info = allocationManager.getAllocation(operator, operatorSet, wethStrategy);
        assertEq(info.currentMagnitude, magnitudeToSet);
    }

    function _slashOperator() public {
        // Get slashing params
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = wethStrategy;
        IAllocationManagerTypes.SlashingParams memory slashingParams = IAllocationManagerTypes.SlashingParams({
            operator: operator,
            operatorSetId: 1,
            wadToSlash: 5e17,
            description: "test"
        });

        // Slash operator
        cheats.prank(avs);
        allocationManager.slashOperator(avs, slashingParams);

        // Assert storage
        IAllocationManagerTypes.Allocation memory info = allocationManager.getAllocation(operator, operatorSet, wethStrategy);
        assertEq(info.currentMagnitude, magnitudeToSet - 5e17);
    }

    function _withdrawStaker() public {
        // Generate queued withdrawal params
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = wethStrategy;
        (uint256[] memory withdrawableShares, ) = delegationManager.getWithdrawableShares(staker, strategies);
        IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawals = new IDelegationManagerTypes.QueuedWithdrawalParams[](1);
        queuedWithdrawals[0] = IDelegationManagerTypes.QueuedWithdrawalParams({
            strategies: strategies,
            depositShares: withdrawableShares,
            withdrawer: staker
        });

        // Generate withdrawal params
        uint256[] memory scaledShares = new uint256[](1);
        scaledShares[0] = 100e18;
        IDelegationManagerTypes.Withdrawal memory withdrawal = IDelegationManagerTypes.Withdrawal({
            staker: staker,
            delegatedTo: operator,
            withdrawer: staker,
            nonce: delegationManager.cumulativeWithdrawalsQueued(staker),
            startBlock: uint32(block.number),
            strategies: strategies,
            scaledShares: scaledShares
        });
        // bytes32 withdrawalRoot = delegationManager.calculateWithdrawalRoot(withdrawal);
        // Generate complete withdrawal params

        cheats.startPrank(staker);
        delegationManager.queueWithdrawals(queuedWithdrawals);

        // Roll passed withdrawal delay
        cheats.roll(block.number + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS());

        // Complete withdrawal
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = weth;
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, true);

        // Assert tokens
        assertEq(weth.balanceOf(staker), wethAmount / 2);
    }
}