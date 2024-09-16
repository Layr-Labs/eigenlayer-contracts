// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// Contracts
import "../../src/contracts/core/DelegationManager.sol";
import "../../src/contracts/core/StrategyManager.sol";
import "../../src/contracts/core/AVSDirectory.sol";
import "../../src/contracts/core/AllocationManager.sol";
import "../../src/contracts/strategies/StrategyBase.sol";

<<<<<<< HEAD
import "src/test/utils/ArrayLib.sol";

=======
>>>>>>> b0193bfe (feat: alm tests)
// Test
import "forge-std/Test.sol";

/// @notice Tests deployed contracts as part of the public devnet
/// Run with: forge test --mc Devnet_Lifecycle_Test --rpc-url $RPC_HOLESKY
<<<<<<< HEAD
contract Devnet_Lifecycle_Test is Test, IAllocationManagerTypes {
    using ArrayLib for *;
=======
contract Devnet_Lifecycle_Test is Test {
>>>>>>> b0193bfe (feat: alm tests)
    
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
<<<<<<< HEAD
        return operatorSetId.toArrayU32();
    }
    
    function _getOperatorSetsArray() internal view returns (OperatorSet[] memory) {
        return  OperatorSet({avs: avs, id: operatorSetId}).toArray();
=======
        uint32[] memory operatorSets = new uint32[](1);
        operatorSets[0] = operatorSetId;
        return operatorSets;
    }
    
    function _getOperatorSetsArray() internal view returns (OperatorSet[] memory) {
        OperatorSet[] memory operatorSets = new OperatorSet[](1);
        operatorSets[0] = OperatorSet({avs: avs, id: operatorSetId});
        return operatorSets;
>>>>>>> b0193bfe (feat: alm tests)
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
<<<<<<< HEAD

       CreateSetParams memory createSetParams =CreateSetParams({
            operatorSetId: operatorSetId,
            strategies: wethStrategy.toArray()
        });

        allocationManager.createOperatorSets(avs, createSetParams.toArray());
=======
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = wethStrategy;

        IAllocationManagerTypes.CreateSetParams memory createSetParams = IAllocationManagerTypes.CreateSetParams({
            operatorSetId: operatorSetId,
            strategies: strategies
        });

        IAllocationManagerTypes.CreateSetParams[] memory array = new IAllocationManagerTypes.CreateSetParams[](1);
        array[0] = createSetParams;

        allocationManager.createOperatorSets(avs, array);
>>>>>>> b0193bfe (feat: alm tests)
        cheats.stopPrank();
    }

    function _registerOperatorToAVS() public {
        cheats.prank(operator);
<<<<<<< HEAD
        allocationManager.registerForOperatorSets(operator,RegisterParams(avs, operatorSetId.toArrayU32(), ""));
=======
        
        uint32[] memory operatorSetIds = new uint32[](1);
        operatorSetIds[0] = operatorSetId;

        allocationManager.registerForOperatorSets(operator, IAllocationManagerTypes.RegisterParams(avs, operatorSetIds, ""));

>>>>>>> b0193bfe (feat: alm tests)
        assertEq(allocationManager.getMembers(OperatorSet(avs, operatorSetId))[0], operator);
    }

    function _setMagnitude() public {
<<<<<<< HEAD
       AllocateParams[] memory allocations = new AllocateParams[](1);
        allocations[0] = AllocateParams({
            operatorSet: operatorSet,
            strategies: wethStrategy.toArray(),
            newMagnitudes: magnitudeToSet.toArrayU64()
=======
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = wethStrategy;

        uint64[] memory magnitudes = new uint64[](1);
        magnitudes[0] = magnitudeToSet;

        IAllocationManagerTypes.AllocateParams[] memory allocations = new IAllocationManagerTypes.AllocateParams[](1);
        allocations[0] = IAllocationManagerTypes.AllocateParams({
            operatorSet: operatorSet,
            strategies: strategies,
            newMagnitudes: magnitudes
>>>>>>> b0193bfe (feat: alm tests)
        });

        cheats.prank(operator);
        allocationManager.modifyAllocations(operator, allocations);

        // Assert storage
<<<<<<< HEAD
       Allocation memory info = allocationManager.getAllocation(operator, operatorSet, wethStrategy);
=======
        IAllocationManagerTypes.Allocation memory info = allocationManager.getAllocation(operator, operatorSet, wethStrategy);
>>>>>>> b0193bfe (feat: alm tests)
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
<<<<<<< HEAD
       SlashingParams memory slashingParams = SlashingParams({
            operator: operator,
            operatorSetId: 1,
            strategies: wethStrategy.toArray(),
            wadsToSlash: 5e17.toArrayU256(),
=======
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = wethStrategy;
        IAllocationManagerTypes.SlashingParams memory slashingParams = IAllocationManagerTypes.SlashingParams({
            operator: operator,
            operatorSetId: 1,
            wadToSlash: 5e17,
>>>>>>> b0193bfe (feat: alm tests)
            description: "test"
        });

        // Slash operator
        cheats.prank(avs);
        allocationManager.slashOperator(avs, slashingParams);

        // Assert storage
<<<<<<< HEAD
       Allocation memory info = allocationManager.getAllocation(operator, operatorSet, wethStrategy);
=======
        IAllocationManagerTypes.Allocation memory info = allocationManager.getAllocation(operator, operatorSet, wethStrategy);
>>>>>>> b0193bfe (feat: alm tests)
        assertEq(info.currentMagnitude, magnitudeToSet - 5e17);
    }

    function _withdrawStaker() public {
        // Generate queued withdrawal params
<<<<<<< HEAD
        IStrategy[] memory strategies = wethStrategy.toArray();
=======
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = wethStrategy;
>>>>>>> b0193bfe (feat: alm tests)
        (uint256[] memory withdrawableShares, ) = delegationManager.getWithdrawableShares(staker, strategies);
        IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawals = new IDelegationManagerTypes.QueuedWithdrawalParams[](1);
        queuedWithdrawals[0] = IDelegationManagerTypes.QueuedWithdrawalParams({
            strategies: strategies,
            depositShares: withdrawableShares,
            withdrawer: staker
        });

        // Generate withdrawal params
<<<<<<< HEAD
=======
        uint256[] memory scaledShares = new uint256[](1);
        scaledShares[0] = 100e18;
>>>>>>> b0193bfe (feat: alm tests)
        IDelegationManagerTypes.Withdrawal memory withdrawal = IDelegationManagerTypes.Withdrawal({
            staker: staker,
            delegatedTo: operator,
            withdrawer: staker,
            nonce: delegationManager.cumulativeWithdrawalsQueued(staker),
            startBlock: uint32(block.number),
            strategies: strategies,
<<<<<<< HEAD
            scaledShares: 100e18.toArrayU256()
=======
            scaledShares: scaledShares
>>>>>>> b0193bfe (feat: alm tests)
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