// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

import "src/contracts/core/DelegationManager.sol";
import "src/contracts/strategies/StrategyBase.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/contracts/libraries/SlashingLib.sol";

/**
 * @notice Unit testing of the DelegationManager contract. Withdrawals are tightly coupled
 * with EigenPodManager and StrategyManager and are part of integration tests.
 * Contracts tested: DelegationManager
 * Contracts not mocked: StrategyBase, PauserRegistry
 */
contract DelegationManagerUnitTests is EigenLayerUnitTestSetup, IDelegationManagerEvents {
    using SlashingLib for *;

    // Contract under test
    DelegationManager delegationManager;
    DelegationManager delegationManagerImplementation;

    // Helper to use in storage
    StakerScalingFactors ssf;

    // Mocks
    StrategyBase strategyImplementation;
    StrategyBase strategyMock;

    IERC20 mockToken;
    uint256 mockTokenInitialSupply = 10e50;

    uint32 constant MIN_WITHDRAWAL_DELAY_BLOCKS = 126_000; // 17.5 days in blocks

    // Delegation signer
    uint256 delegationSignerPrivateKey = uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);
    uint256 stakerPrivateKey = uint256(123_456_789);

    // empty string reused across many tests
    string emptyStringForMetadataURI;

    // "empty" / zero salt, reused across many tests
    bytes32 emptySalt;

    // reused in various tests. in storage to help handle stack-too-deep errors
    address defaultStaker = cheats.addr(uint256(123_456_789));
    address defaultOperator = address(this);
    address defaultApprover = cheats.addr(delegationSignerPrivateKey);
    address defaultAVS = address(this);

    // 604800 seconds in week / 12 = 50,400 blocks
    uint256 minWithdrawalDelayBlocks = 50400;
    IStrategy[] public initializeStrategiesToSetDelayBlocks;
    uint256[] public initializeWithdrawalDelayBlocks;

    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    // Index for flag that pauses new delegations when set.
    uint8 internal constant PAUSED_NEW_DELEGATION = 0;

    // Index for flag that pauses queuing new withdrawals when set.
    uint8 internal constant PAUSED_ENTER_WITHDRAWAL_QUEUE = 1;

    // Index for flag that pauses completing existing withdrawals when set.
    uint8 internal constant PAUSED_EXIT_WITHDRAWAL_QUEUE = 2;

    // the number of 12-second blocks in 30 days (60 * 60 * 24 * 30 / 12 = 216,000)
    uint256 public constant MAX_WITHDRAWAL_DELAY_BLOCKS = 216000;

    /// @notice mappings used to handle duplicate entries in fuzzed address array input
    mapping(address => uint256) public totalSharesForStrategyInArray;
    mapping(IStrategy => uint256) public totalSharesDecreasedForStrategy;
    mapping(IStrategy => uint256) public delegatedSharesBefore;

    function setUp() public virtual override {
        // Setup
        EigenLayerUnitTestSetup.setUp();

        // Deploy DelegationManager implmentation and proxy
        initializeStrategiesToSetDelayBlocks = new IStrategy[](0);
        initializeWithdrawalDelayBlocks = new uint256[](0);
        delegationManagerImplementation = new DelegationManager(
            IAVSDirectory(address(avsDirectoryMock)), 
            IStrategyManager(address(strategyManagerMock)), 
            IEigenPodManager(address(eigenPodManagerMock)), 
            IAllocationManager(address(allocationManagerMock)), 
            MIN_WITHDRAWAL_DELAY_BLOCKS
        );
        delegationManager = DelegationManager(
            address(
                new TransparentUpgradeableProxy(
                    address(delegationManagerImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        DelegationManager.initialize.selector,
                        address(this),
                        pauserRegistry,
                        0, // 0 is initialPausedStatus
                        minWithdrawalDelayBlocks,
                        initializeStrategiesToSetDelayBlocks,
                        initializeWithdrawalDelayBlocks
                    )
                )
            )
        );

        // Deploy mock token and strategy
        mockToken = new ERC20PresetFixedSupply("Mock Token", "MOCK", mockTokenInitialSupply, address(this));
        strategyImplementation = new StrategyBase(IStrategyManager(address(strategyManagerMock)));
        strategyMock = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, mockToken, pauserRegistry)
                )
            )
        );

        // Exclude delegation manager from fuzzed tests
        isExcludedFuzzAddress[address(delegationManager)] = true;
        isExcludedFuzzAddress[defaultApprover] = true;
    }

    /**
     * INTERNAL / HELPER FUNCTIONS
     */

    /**
     * @notice internal function to deploy mock tokens and strategies and have the staker deposit into them. 
     * Since we are mocking the strategyManager we call strategyManagerMock.setDeposits so that when
     * DelegationManager calls getDeposits, we can have these share amounts returned.
     */
    function _deployAndDepositIntoStrategies(
        address staker,
        uint256[] memory sharesAmounts
    ) internal returns (IStrategy[] memory) {
        uint256 numStrats = sharesAmounts.length;
        IStrategy[] memory strategies = new IStrategy[](numStrats);
        uint256[] memory withdrawalDelayBlocks = new uint256[](strategies.length);
        for (uint8 i = 0; i < numStrats; i++) {
            withdrawalDelayBlocks[i] = bound(uint256(keccak256(abi.encode(staker, i))), 0, MAX_WITHDRAWAL_DELAY_BLOCKS);
            ERC20PresetFixedSupply token = new ERC20PresetFixedSupply(
                string(abi.encodePacked("Mock Token ", i)),
                string(abi.encodePacked("MOCK", i)),
                mockTokenInitialSupply,
                address(this)
            );
            strategies[i] = StrategyBase(
                address(
                    new TransparentUpgradeableProxy(
                        address(strategyImplementation),
                        address(eigenLayerProxyAdmin),
                        abi.encodeWithSelector(StrategyBase.initialize.selector, token, pauserRegistry)
                    )
                )
            );
        }
        // delegationManager.setStrategyWithdrawalDelayBlocks(strategies, withdrawalDelayBlocks);
        strategyManagerMock.setDeposits(staker, strategies, sharesAmounts);
        return strategies;
    }

    /**
     * @notice internal function for calculating a signature from the delegationSigner corresponding to `_delegationSignerPrivateKey`, approving
     * the `staker` to delegate to `operator`, with the specified `salt`, and expiring at `expiry`.
     */
    function _getApproverSignature(
        uint256 _delegationSignerPrivateKey,
        address staker,
        address operator,
        bytes32 salt,
        uint256 expiry
    ) internal view returns (ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry) {
        approverSignatureAndExpiry.expiry = expiry;
        {
            bytes32 digestHash = delegationManager.calculateDelegationApprovalDigestHash(
                staker,
                operator,
                delegationManager.delegationApprover(operator),
                salt,
                expiry
            );
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(_delegationSignerPrivateKey, digestHash);
            approverSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }
        return approverSignatureAndExpiry;
    }

    /**
     * @notice internal function for calculating a signature from the staker corresponding to `_stakerPrivateKey`, delegating them to
     * the `operator`, and expiring at `expiry`.
     */
    function _getStakerSignature(
        uint256 _stakerPrivateKey,
        address operator,
        uint256 expiry
    ) internal view returns (ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry) {
        address staker = cheats.addr(stakerPrivateKey);
        stakerSignatureAndExpiry.expiry = expiry;
        {
            bytes32 digestHash = delegationManager.calculateCurrentStakerDelegationDigestHash(staker, operator, expiry);
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(_stakerPrivateKey, digestHash);
            stakerSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }
        return stakerSignatureAndExpiry;
    }

    // @notice Assumes operator does not have a delegation approver & staker != approver
    function _delegateToOperatorWhoAcceptsAllStakers(address staker, address operator) internal {
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        cheats.prank(staker);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);
    }

    function _delegateToOperatorWhoRequiresSig(address staker, address operator, bytes32 salt) internal {
        uint256 expiry = type(uint256).max;
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(
            delegationSignerPrivateKey,
            staker,
            operator,
            salt,
            expiry
        );
        cheats.prank(staker);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);
    }

    function _delegateToOperatorWhoRequiresSig(address staker, address operator) internal {
        _delegateToOperatorWhoRequiresSig(staker, operator, emptySalt);
    }

    function _delegateToBySignatureOperatorWhoAcceptsAllStakers(
        address staker,
        address caller,
        address operator,
        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry,
        bytes32 salt
    ) internal {
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        cheats.prank(caller);
        delegationManager.delegateToBySignature(
            staker,
            operator,
            stakerSignatureAndExpiry,
            approverSignatureAndExpiry,
            salt
        );
    }

    function _delegateToBySignatureOperatorWhoRequiresSig(
        address staker,
        address caller,
        address operator,
        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry,
        bytes32 salt
    ) internal {
        uint256 expiry = type(uint256).max;
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(
            delegationSignerPrivateKey,
            staker,
            operator,
            salt,
            expiry
        );
        cheats.prank(caller);
        delegationManager.delegateToBySignature(
            staker,
            operator,
            stakerSignatureAndExpiry,
            approverSignatureAndExpiry,
            salt
        );
    }

    function _registerOperatorWithBaseDetails(address operator) internal {
        IDelegationManagerTypes.OperatorDetails memory operatorDetails = IDelegationManagerTypes.OperatorDetails({
            __deprecated_earningsReceiver: operator,
            delegationApprover: address(0),
            __deprecated_stakerOptOutWindowBlocks: 0
        });
        _registerOperator(operator, operatorDetails, emptyStringForMetadataURI);
    }

    function _registerOperatorWithDelegationApprover(address operator) internal {
        IDelegationManagerTypes.OperatorDetails memory operatorDetails = IDelegationManagerTypes.OperatorDetails({
            __deprecated_earningsReceiver: operator,
            delegationApprover: defaultApprover,
            __deprecated_stakerOptOutWindowBlocks: 0
        });
        _registerOperator(operator, operatorDetails, emptyStringForMetadataURI);
    }

    function _registerOperatorWith1271DelegationApprover(address operator) internal returns (ERC1271WalletMock) {
        address delegationSigner = defaultApprover;
        /**
         * deploy a ERC1271WalletMock contract with the `delegationSigner` address as the owner,
         * so that we can create valid signatures from the `delegationSigner` for the contract to check when called
         */
        ERC1271WalletMock wallet = new ERC1271WalletMock(delegationSigner);

        IDelegationManagerTypes.OperatorDetails memory operatorDetails = IDelegationManagerTypes.OperatorDetails({
            __deprecated_earningsReceiver: operator,
            delegationApprover: address(wallet),
            __deprecated_stakerOptOutWindowBlocks: 0
        });
        _registerOperator(operator, operatorDetails, emptyStringForMetadataURI);

        return wallet;
    }

    function _registerOperator(
        address operator,
        IDelegationManagerTypes.OperatorDetails memory operatorDetails,
        string memory metadataURI
    ) internal filterFuzzedAddressInputs(operator) {
        cheats.prank(operator);
        delegationManager.registerAsOperator(operatorDetails, 0, metadataURI);
    }

    /**
     * @notice Using this helper function to fuzz withdrawalAmounts since fuzzing two dynamic sized arrays of equal lengths
     * reject too many inputs. 
     */
    function _fuzzWithdrawalAmounts(uint256[] memory depositAmounts) internal view returns (uint256[] memory) {
        uint256[] memory withdrawalAmounts = new uint256[](depositAmounts.length);
        for (uint256 i = 0; i < depositAmounts.length; i++) {
            cheats.assume(depositAmounts[i] > 0);
            // generate withdrawal amount within range s.t withdrawAmount <= depositAmount
            withdrawalAmounts[i] = bound(
                uint256(keccak256(abi.encodePacked(depositAmounts[i]))),
                0,
                depositAmounts[i]
            );
        }
        return withdrawalAmounts;
    }

    function _setUpQueueWithdrawalsSingleStrat(
        address staker,
        address withdrawer,
        IStrategy strategy,
        uint256 sharesToWithdraw
    ) internal view returns (
        IDelegationManagerTypes.QueuedWithdrawalParams[] memory,
        IDelegationManagerTypes.Withdrawal memory,
        bytes32
    ) {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        strategyArray[0] = strategy;
        IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams = new IDelegationManagerTypes.QueuedWithdrawalParams[](1);
        {
            uint256[] memory withdrawalAmounts = new uint256[](1);
            withdrawalAmounts[0] = sharesToWithdraw;

            queuedWithdrawalParams[0] = IDelegationManagerTypes.QueuedWithdrawalParams({
                strategies: strategyArray,
                shares: withdrawalAmounts,
                withdrawer: withdrawer
            });
        }

        // Get scaled shares to withdraw
        uint256[] memory scaledSharesArray = new uint256[](1);
        scaledSharesArray[0] = _getScaledShares(staker, strategy, sharesToWithdraw);

        IDelegationManagerTypes.Withdrawal memory withdrawal = IDelegationManagerTypes.Withdrawal({
            staker: staker,
            delegatedTo: delegationManager.delegatedTo(staker),
            withdrawer: withdrawer,
            nonce: delegationManager.cumulativeWithdrawalsQueued(staker),
            startBlock: uint32(block.number),
            strategies: strategyArray,
            scaledShares: scaledSharesArray
        });
        bytes32 withdrawalRoot = delegationManager.calculateWithdrawalRoot(withdrawal);
        
        return (queuedWithdrawalParams, withdrawal, withdrawalRoot);
    }

    function _setUpQueueWithdrawals(
        address staker,
        address withdrawer,
        IStrategy[] memory strategies,
        uint256[] memory withdrawalAmounts
    ) internal view returns (
        IDelegationManagerTypes.QueuedWithdrawalParams[] memory,
        IDelegationManagerTypes.Withdrawal memory,
        bytes32
    ) {
        IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams = new IDelegationManagerTypes.QueuedWithdrawalParams[](1);
        {
            queuedWithdrawalParams[0] = IDelegationManagerTypes.QueuedWithdrawalParams({
                strategies: strategies,
                shares: withdrawalAmounts,
                withdrawer: withdrawer
            });
        }

        // Get scaled shares to withdraw
        uint256[] memory scaledSharesArray = new uint256[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            scaledSharesArray[i] = _getScaledShares(staker, strategies[i], withdrawalAmounts[i]);
        }
        
        IDelegationManagerTypes.Withdrawal memory withdrawal = IDelegationManagerTypes.Withdrawal({
            staker: staker,
            delegatedTo: delegationManager.delegatedTo(staker),
            withdrawer: withdrawer,
            nonce: delegationManager.cumulativeWithdrawalsQueued(staker),
            startBlock: uint32(block.number),
            strategies: strategies,
            scaledShares: scaledSharesArray
        });
        bytes32 withdrawalRoot = delegationManager.calculateWithdrawalRoot(withdrawal);
        
        return (queuedWithdrawalParams, withdrawal, withdrawalRoot);
    }

    function _getScaledShares(address staker, IStrategy strategy, uint256 sharesToWithdraw) internal view returns (uint256) {
        // Setup vars
        address operator = delegationManager.delegatedTo(staker);
        IStrategy[] memory strategyArray = new IStrategy[](1);
        strategyArray[0] = strategy;

        // Set scaling factors
        (uint256 depositScalingFactor, bool isBeaconChainScalingFactorSet, uint64 beaconChainScalingFactor) = delegationManager.stakerScalingFactor(staker, strategy);
        StakerScalingFactors memory stakerScalingFactor = StakerScalingFactors({
            depositScalingFactor: depositScalingFactor,
            isBeaconChainScalingFactorSet: isBeaconChainScalingFactorSet,
            beaconChainScalingFactor: beaconChainScalingFactor
        });

        uint256 scaledShares = SlashingLib.scaleSharesForQueuedWithdrawal(
            sharesToWithdraw,
            stakerScalingFactor,
            allocationManagerMock.getMaxMagnitudes(operator, strategyArray)[0]
        );

        return scaledShares;
    }

    /**
     * Deploy and deposit staker into a single strategy, then set up a queued withdrawal for the staker
     * Assumptions: 
     * - operator is already a registered operator.
     * - withdrawalAmount <= depositAmount
     */
    function _setUpCompleteQueuedWithdrawalSingleStrat(
        address staker,
        address withdrawer,
        uint256 depositAmount,
        uint256 withdrawalAmount
    ) internal returns (IDelegationManagerTypes.Withdrawal memory, IERC20[] memory, bytes32) {
        uint256[] memory depositAmounts = new uint256[](1);
        depositAmounts[0] = depositAmount;
        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(staker, depositAmounts);
        (
            IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: staker,
            withdrawer: withdrawer,
            strategy: strategies[0],
            sharesToWithdraw: withdrawalAmount
        });

        cheats.prank(staker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
        // Set the current deposits to be the depositAmount - withdrawalAmount
        uint256[] memory currentAmounts = new uint256[](1);
        currentAmounts[0] = depositAmount - withdrawalAmount;
        strategyManagerMock.setDeposits(staker, strategies, currentAmounts);

        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = strategies[0].underlyingToken();
        return (withdrawal, tokens, withdrawalRoot);
    }

    /**
     * Deploy and deposit staker into strategies, then set up a queued withdrawal for the staker
     * Assumptions: 
     * - operator is already a registered operator.
     * - for each i, withdrawalAmount[i] <= depositAmount[i] (see filterFuzzedDepositWithdrawInputs above)
     */
    function _setUpCompleteQueuedWithdrawal(
        address staker,
        address withdrawer,
        uint256[] memory depositAmounts,
        uint256[] memory withdrawalAmounts
    ) internal returns (IDelegationManagerTypes.Withdrawal memory, IERC20[] memory, bytes32) {
        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(staker, depositAmounts);

        IERC20[] memory tokens = new IERC20[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            tokens[i] = strategies[i].underlyingToken();
        }

        (
            IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawals({
            staker: staker,
            withdrawer: withdrawer,
            strategies: strategies,
            withdrawalAmounts: withdrawalAmounts
        });

        cheats.prank(staker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        return (withdrawal, tokens, withdrawalRoot);
    }

    function _setOperatorMagnitude(address operator, IStrategy strategy, uint64 magnitude) internal {
        allocationManagerMock.setMaxMagnitude(operator, strategy, magnitude);
    }
}

contract DelegationManagerUnitTests_Initialization_Setters is DelegationManagerUnitTests {
    function test_initialization() public view {
        assertEq(
            address(delegationManager.strategyManager()),
            address(strategyManagerMock),
            "constructor / initializer incorrect, strategyManager set wrong"
        );
        assertEq(
            address(delegationManager.pauserRegistry()),
            address(pauserRegistry),
            "constructor / initializer incorrect, pauserRegistry set wrong"
        );
        assertEq(
            address(delegationManager.eigenPodManager()),
            address(eigenPodManagerMock),
            "constructor / initializer incorrect, eigenPodManager set wrong"
        );
        assertEq(
            address(delegationManager.allocationManager()),
            address(allocationManagerMock),
            "constructor / initializer incorrect, allocationManager set wrong"
        );
        assertEq(
            delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS(),
            MIN_WITHDRAWAL_DELAY_BLOCKS,
            "constructor / initializer incorrect, MIN_WITHDRAWAL_DELAY set wrong"
        );
        assertEq(delegationManager.owner(), address(this), "constructor / initializer incorrect, owner set wrong");
        assertEq(delegationManager.paused(), 0, "constructor / initializer incorrect, paused status set wrong");
    }

    /// @notice Verifies that the DelegationManager cannot be iniitalized multiple times
    function test_initialize_revert_reinitialization() public {
        cheats.expectRevert("Initializable: contract is already initialized");
        delegationManager.initialize(
            address(this),
            pauserRegistry,
            0
        );
    }
}

contract DelegationManagerUnitTests_RegisterModifyOperator is DelegationManagerUnitTests {
    function test_registerAsOperator_revert_paused() public {
        // set the pausing flag
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_NEW_DELEGATION);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.registerAsOperator(
            IDelegationManagerTypes.OperatorDetails({
                __deprecated_earningsReceiver: defaultOperator,
                delegationApprover: address(0),
                __deprecated_stakerOptOutWindowBlocks: 0
            }),
            0,
            emptyStringForMetadataURI
        );
    }

    // @notice Verifies that someone cannot successfully call `DelegationManager.registerAsOperator(operatorDetails)` again after registering for the first time
    function testFuzz_registerAsOperator_revert_cannotRegisterMultipleTimes(
        IDelegationManagerTypes.OperatorDetails memory operatorDetails
    ) public {
        // Register once
        cheats.startPrank(defaultOperator);
        delegationManager.registerAsOperator(operatorDetails, 0, emptyStringForMetadataURI);

        // Expect revert when register again
        cheats.expectRevert(IDelegationManagerErrors.ActivelyDelegated.selector);
        delegationManager.registerAsOperator(operatorDetails, 0, emptyStringForMetadataURI);
        cheats.stopPrank();
    }

    /**
     * @notice `operator` registers via calling `DelegationManager.registerAsOperator(operatorDetails, metadataURI)`
     * Should be able to set any parameters, other than too high value for `stakerOptOutWindowBlocks`
     * The set parameters should match the desired parameters (correct storage update)
     * Operator becomes delegated to themselves
     * Properly emits events – especially the `OperatorRegistered` event, but also `StakerDelegated` & `OperatorDetailsModified` events
     * Reverts appropriately if operator was already delegated to someone (including themselves, i.e. they were already an operator)
     * @param operator and @param operatorDetails are fuzzed inputs
     */
    function testFuzz_registerAsOperator(
        address operator,
        IDelegationManagerTypes.OperatorDetails memory operatorDetails,
        string memory metadataURI
    ) public filterFuzzedAddressInputs(operator) {
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorDetailsModified(operator, operatorDetails);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(operator, operator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorRegistered(operator, operatorDetails);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorMetadataURIUpdated(operator, metadataURI);

        cheats.prank(operator);
        delegationManager.registerAsOperator(operatorDetails, 0, metadataURI);

        // Storage checks
        assertEq(
            operatorDetails.delegationApprover,
            delegationManager.delegationApprover(operator),
            "delegationApprover not set correctly"
        );
        assertEq(delegationManager.delegatedTo(operator), operator, "operator not delegated to self");
    }

    /// TODO: registerAsOperator 2 separate addresses
    /// function testTwoSelfOperatorsRegister() public {}


    // @notice Verifies that a staker who is actively delegated to an operator cannot register as an operator (without first undelegating, at least)
    function testFuzz_Revert_registerAsOperator_cannotRegisterWhileDelegated(
        address staker,
        IDelegationManagerTypes.OperatorDetails memory operatorDetails
    ) public filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != defaultOperator);

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, emptySalt);

        // expect revert if attempt to register as operator
        cheats.expectRevert(IDelegationManagerErrors.ActivelyDelegated.selector);
        delegationManager.registerAsOperator(operatorDetails, 0, emptyStringForMetadataURI);

        cheats.stopPrank();
    }
    
    /// TODO: Add test for registerAsOperator where the operator has existing deposits in strategies
    /// Assert:
    ///     depositShares == operatorShares == withdrawableShares
    ///     check operatorDetails hash encode matches the operatorDetails hash stored (call view function)
    function testFuzz_registerAsOperator_withDeposits() public {}

    /**
     * @notice Tests that an operator can modify their OperatorDetails by calling `DelegationManager.modifyOperatorDetails`
     * Should be able to set any parameters, other than setting their `earningsReceiver` to the zero address or too high value for `stakerOptOutWindowBlocks`
     * The set parameters should match the desired parameters (correct storage update)
     * Properly emits an `OperatorDetailsModified` event
     * Reverts appropriately if the caller is not an operator
     * Reverts if operator tries to decrease their `stakerOptOutWindowBlocks` parameter
     * @param initialOperatorDetails and @param modifiedOperatorDetails are fuzzed inputs
     */
    function testFuzz_modifyOperatorParameters(
        IDelegationManagerTypes.OperatorDetails memory initialOperatorDetails,
        IDelegationManagerTypes.OperatorDetails memory modifiedOperatorDetails
    ) public {
        _registerOperator(defaultOperator, initialOperatorDetails, emptyStringForMetadataURI);

        cheats.startPrank(defaultOperator);

        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorDetailsModified(defaultOperator, modifiedOperatorDetails);
        delegationManager.modifyOperatorDetails(modifiedOperatorDetails);

        assertEq(
            modifiedOperatorDetails.delegationApprover,
            delegationManager.delegationApprover(defaultOperator),
            "delegationApprover not set correctly"
        );
        assertEq(delegationManager.delegatedTo(defaultOperator), defaultOperator, "operator not delegated to self");
        // or else the transition is disallowed

        cheats.stopPrank();
    }

    // @notice Tests that an address which is not an operator cannot successfully call `updateOperatorMetadataURI`.
    function test_updateOperatorMetadataUri_notRegistered() public {
        assertFalse(delegationManager.isOperator(defaultOperator), "bad test setup");

        cheats.prank(defaultOperator);
        cheats.expectRevert(IDelegationManagerErrors.OperatorNotRegistered.selector);
        delegationManager.updateOperatorMetadataURI(emptyStringForMetadataURI);
    }

    /**
     * @notice Verifies that a staker cannot call cannot modify their `OperatorDetails` without first registering as an operator
     * @dev This is an important check to ensure that our definition of 'operator' remains consistent, in particular for preserving the
     * invariant that 'operators' are always delegated to themselves
     */
    function testFuzz_Revert_updateOperatorMetadataUri_notOperator(
        IDelegationManagerTypes.OperatorDetails memory operatorDetails
    ) public {
        cheats.expectRevert(IDelegationManagerErrors.OperatorNotRegistered.selector);
        delegationManager.modifyOperatorDetails(operatorDetails);
    }

    // @notice Tests that an operator who calls `updateOperatorMetadataURI` will correctly see an `OperatorMetadataURIUpdated` event emitted with their input
    function testFuzz_UpdateOperatorMetadataURI(string memory metadataURI) public {
        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // call `updateOperatorMetadataURI` and check for event
        cheats.prank(defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorMetadataURIUpdated(defaultOperator, metadataURI);
        delegationManager.updateOperatorMetadataURI(metadataURI);
    }
}

contract DelegationManagerUnitTests_delegateTo is DelegationManagerUnitTests {
    function test_Revert_WhenPaused() public {
        cheats.prank(defaultOperator);
        delegationManager.registerAsOperator(
            IDelegationManagerTypes.OperatorDetails({
                __deprecated_earningsReceiver: defaultOperator,
                delegationApprover: address(0),
                __deprecated_stakerOptOutWindowBlocks: 0
            }),
            0,
            emptyStringForMetadataURI
        );

        // set the pausing flag
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_NEW_DELEGATION);

        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        cheats.prank(defaultStaker);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, emptySalt);
    }

    /**
     * @notice Delegates from `staker` to an operator, then verifies that the `staker` cannot delegate to another `operator` (at least without first undelegating)
     */
    function testFuzz_Revert_WhenDelegateWhileDelegated(
        address staker,
        address operator,
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 salt
    ) public filterFuzzedAddressInputs(staker) filterFuzzedAddressInputs(operator) {
        // filter out input since if the staker tries to delegate again after registering as an operator, we will revert earlier than this test is designed to check
        cheats.assume(staker != operator);

        // delegate from the staker to an operator
        cheats.assume(operator != address(this));
        _registerOperatorWithBaseDetails(operator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, operator);

        // try to delegate again and check that the call reverts
        cheats.prank(staker);
        cheats.expectRevert(IDelegationManagerErrors.ActivelyDelegated.selector);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);
    }

    /// @notice Verifies that `staker` cannot delegate to an unregistered `operator`
    function testFuzz_Revert_WhenDelegateToUnregisteredOperator(
        address staker,
        address operator
    ) public filterFuzzedAddressInputs(staker) filterFuzzedAddressInputs(operator) {
        assertFalse(delegationManager.isOperator(operator), "incorrect test input?");

        // try to delegate and check that the call reverts
        cheats.prank(staker);
        cheats.expectRevert(IDelegationManagerErrors.OperatorNotRegistered.selector);
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);
    }

    /**
     * @notice `staker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * via the `staker` calling `DelegationManager.delegateTo`
     * The function should pass with any `operatorSignature` input (since it should be unused)
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_OperatorWhoAcceptsAllStakers_StrategyManagerShares(
        address staker,
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 salt,
        uint128 shares
    ) public filterFuzzedAddressInputs(staker) {
        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        _registerOperatorWithBaseDetails(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
        // Set staker shares in StrategyManager
        IStrategy[] memory strategiesToReturn = new IStrategy[](1);
        strategiesToReturn[0] = strategyMock;
        uint256[] memory sharesToReturn = new uint256[](1);
        sharesToReturn[0] = shares;
        strategyManagerMock.setDeposits(staker, strategiesToReturn, sharesToReturn);
        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
        // delegate from the `staker` to the operator
        cheats.prank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(defaultOperator, staker, strategyMock, shares);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DepositScalingFactorUpdated(staker, strategyMock, WAD);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);

        assertEq(operatorSharesBefore + shares, operatorSharesAfter, "operator shares not increased correctly");
        assertTrue(delegationManager.isOperator(defaultOperator), "staker not registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
    }

    /**
     * @notice `staker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * but it should revert as the strategy has been fully slashed for the operator.
     */
    function testFuzz_Revert_OperatorWhoAcceptsAllStakers_AlreadySlashed100Percent_StrategyManagerShares(
        address staker,
        uint128 shares
    ) public filterFuzzedAddressInputs(staker) {
        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        // Set empty sig+salt
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        bytes32 salt;

        _registerOperatorWithBaseDetails(defaultOperator);

        // Set staker shares in StrategyManager
        IStrategy[] memory strategiesToReturn = new IStrategy[](1);
        strategiesToReturn[0] = strategyMock;
        uint256[] memory sharesToReturn = new uint256[](1);
        sharesToReturn[0] = shares;
        strategyManagerMock.setDeposits(staker, strategiesToReturn, sharesToReturn);

        // Set the operators magnitude to be 0
        _setOperatorMagnitude(defaultOperator, strategyMock, 0);

        // delegate from the `staker` to the operator
        cheats.prank(staker);
        cheats.expectRevert(IDelegationManagerErrors.FullySlashed.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);

        assertTrue(delegationManager.delegatedTo(staker) != defaultOperator, "staker should not be delegated to the operator");
        assertFalse(delegationManager.isDelegated(staker), "staker should not be delegated");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
    }

    /**
     * @notice `staker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * but it should revert as the beaconChainStrategy has been fully slashed for the operator.
     * The function should pass with any `operatorSignature` input (since it should be unused)
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * OperatorSharesIncreased event should only be emitted if beaconShares is > 0. Since a staker can have negative shares nothing should happen in that case
     */
    function testFuzz_Revert_OperatorWhoAcceptsAllStakers_AlreadySlashed100Percent_BeaconChainStrategyShares(
        address staker,
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 salt,
        int256 beaconShares
    ) public filterFuzzedAddressInputs(staker) {
        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);
        cheats.assume(beaconShares > 0);

        _registerOperatorWithBaseDetails(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
        // Set staker shares in BeaconChainStrategy
        eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);
        uint256 beaconSharesBefore = delegationManager.operatorShares(staker, beaconChainETHStrategy);

        // Set the operators magnitude for native restaking to be 0
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, 0);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectRevert(IDelegationManagerErrors.FullySlashed.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        uint256 beaconSharesAfter = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
        
        assertEq(
            beaconSharesBefore,
            beaconSharesAfter,
            "operator beaconchain shares should not have increased with negative shares"
        );
        assertTrue(delegationManager.delegatedTo(staker) != defaultOperator, "staker should not be delegated to the operator");
        assertFalse(delegationManager.isDelegated(staker), "staker should not be delegated");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
    }

    /// @notice Same test as above, except operator has a magnitude < WAD for the given strategies
    /// TODO: fuzz the magnitude
    function testFuzz_OperatorWhoAcceptsAllStakers_AlreadySlashed_StrategyManagerShares(
        address staker,
        uint128 shares
    ) public filterFuzzedAddressInputs(staker) {
        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        // Set empty sig+salt
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        bytes32 salt;

        _registerOperatorWithBaseDetails(defaultOperator);

        // Set staker shares in StrategyManager
        IStrategy[] memory strategiesToReturn = new IStrategy[](1);
        strategiesToReturn[0] = strategyMock;
        uint256[] memory sharesToReturn = new uint256[](1);
        sharesToReturn[0] = shares;
        strategyManagerMock.setDeposits(staker, strategiesToReturn, sharesToReturn);
        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);

        // Set the operators magnitude to be 50%
        _setOperatorMagnitude(defaultOperator, strategyMock, 5e17);

        // delegate from the `staker` to the operator
        cheats.prank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(defaultOperator, staker, strategyMock, shares);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DepositScalingFactorUpdated(staker, strategyMock, 2e18);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);


        assertEq(operatorSharesBefore + shares, operatorSharesAfter, "operator shares not increased correctly");
        assertTrue(delegationManager.isOperator(defaultOperator), "staker not registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );

        (uint256[] memory withdrawableShares) = delegationManager.getWithdrawableShares(staker, strategiesToReturn);
        assertEq(withdrawableShares[0], shares, "staker shares not set correctly");
    }

    /**
     * @notice `staker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * via the `staker` calling `DelegationManager.delegateTo`
     * The function should pass with any `operatorSignature` input (since it should be unused)
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * OperatorSharesIncreased event should only be emitted if beaconShares is > 0. Since a staker can have negative shares nothing should happen in that case
     */
    // TODO: fuzz the magnitude
    function testFuzz_OperatorWhoAcceptsAllStakers_BeaconChainStrategyShares(
        address staker,
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 salt,
        int256 beaconShares
    ) public filterFuzzedAddressInputs(staker) {
        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        _registerOperatorWithBaseDetails(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
        // Set staker shares in BeaconChainStrategy
        eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);
        uint256 beaconSharesBefore = delegationManager.operatorShares(staker, beaconChainETHStrategy);
        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        if (beaconShares > 0) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(defaultOperator, staker, beaconChainETHStrategy, uint256(beaconShares));
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit DepositScalingFactorUpdated(staker, beaconChainETHStrategy, WAD);
        }
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        uint256 beaconSharesAfter = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
        if (beaconShares <= 0) {
            assertEq(
                beaconSharesBefore,
                beaconSharesAfter,
                "operator beaconchain shares should not have increased with negative shares"
            );
        } else {
            assertEq(
                beaconSharesBefore + uint256(beaconShares),
                beaconSharesAfter,
                "operator beaconchain shares not increased correctly"
            );
        }
        assertTrue(delegationManager.isOperator(defaultOperator), "staker not registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
    }

    /// @notice Same test as above, except operator has a magnitude < WAD for the given strategies
    /// TODO: fuzz the magnitude
    function testFuzz_OperatorWhoAcceptsAllStakers_AlreadySlashed_BeaconChainStrategyShares(
        address staker,
        int256 beaconShares
    ) public filterFuzzedAddressInputs(staker) {
        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        // Set empty sig+salt
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        bytes32 salt;

        _registerOperatorWithBaseDetails(defaultOperator);

        // Set staker shares in BeaconChainStrategy
        eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);
        uint256 beaconSharesBefore = delegationManager.operatorShares(staker, beaconChainETHStrategy);

        // Set the operators magnitude to be 50%
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, 5e17);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        if (beaconShares > 0) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(defaultOperator, staker, beaconChainETHStrategy, uint256(beaconShares));
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit DepositScalingFactorUpdated(staker, beaconChainETHStrategy, 2e18);
        }
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        uint256 beaconSharesAfter = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
        if (beaconShares <= 0) {
            assertEq(
                beaconSharesBefore,
                beaconSharesAfter,
                "operator beaconchain shares should not have increased with negative shares"
            );
        } else {
            assertEq(
                beaconSharesBefore + uint256(beaconShares),
                beaconSharesAfter,
                "operator beaconchain shares not increased correctly"
            );
        }
        assertTrue(delegationManager.isOperator(defaultOperator), "staker not registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );

        IStrategy[] memory strategiesToReturn = new IStrategy[](1);
        strategiesToReturn[0] = beaconChainETHStrategy;
        (uint256[] memory withdrawableShares) = delegationManager.getWithdrawableShares(staker, strategiesToReturn);
        if (beaconShares > 0) {
            assertEq(withdrawableShares[0], uint256(beaconShares), "staker shares not set correctly");
        } else {
            assertEq(withdrawableShares[0], 0, "staker shares not set correctly");
        }
    }

    /**
     * @notice `staker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * via the `staker` calling `DelegationManager.delegateTo`
     * Similar to tests above but now with staker who has both EigenPod and StrategyManager shares.
     */
    function testFuzz_OperatorWhoAcceptsAllStakers_BeaconChainAndStrategyManagerShares(
        address staker,
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 salt,
        int256 beaconShares,
        uint256 shares
    ) public filterFuzzedAddressInputs(staker) {
        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        _registerOperatorWithBaseDetails(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
        // Set staker shares in BeaconChainStrategy and StrategyMananger
        IStrategy[] memory strategiesToReturn = new IStrategy[](1);
        strategiesToReturn[0] = strategyMock;
        uint256[] memory sharesToReturn = new uint256[](1);
        sharesToReturn[0] = shares;
        strategyManagerMock.setDeposits(staker, strategiesToReturn, sharesToReturn);
        eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);
        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
        uint256 beaconSharesBefore = delegationManager.operatorShares(staker, beaconChainETHStrategy);
        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(defaultOperator, staker, strategyMock, shares);
        if (beaconShares > 0) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(defaultOperator, staker, beaconChainETHStrategy, uint256(beaconShares));
        }
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
        uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);
        uint256 beaconSharesAfter = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
        if (beaconShares <= 0) {
            assertEq(
                beaconSharesBefore,
                beaconSharesAfter,
                "operator beaconchain shares should not have increased with negative shares"
            );
        } else {
            assertEq(
                beaconSharesBefore + uint256(beaconShares),
                beaconSharesAfter,
                "operator beaconchain shares not increased correctly"
            );
        }
        assertEq(operatorSharesBefore + shares, operatorSharesAfter, "operator shares not increased correctly");
        assertTrue(delegationManager.isOperator(defaultOperator), "staker not registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
    }

    /**
     * @notice `staker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * via the `staker` calling `DelegationManager.delegateTo`
     * Similar to tests above but now with staker who has both EigenPod and StrategyManager shares.
     */
    //TODO: fuzz magnitude
    function testFuzz_OperatorWhoAcceptsAllStakers_AlreadySlashed_BeaconChainAndStrategyManagerShares(
        address staker,
        int256 beaconShares,
        uint128 shares
    ) public filterFuzzedAddressInputs(staker) {
        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        _registerOperatorWithBaseDetails(defaultOperator);

        // Set empty sig+salt
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        bytes32 salt;

        // Set the operators magnitude to be 50%
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, 5e17);
        _setOperatorMagnitude(defaultOperator, strategyMock, 5e17);

        // Set staker shares in BeaconChainStrategy and StrategyMananger
        IStrategy[] memory strategiesToReturn = new IStrategy[](1);
        strategiesToReturn[0] = strategyMock;
        uint256[] memory sharesToReturn = new uint256[](1);
        sharesToReturn[0] = shares;
        strategyManagerMock.setDeposits(staker, strategiesToReturn, sharesToReturn);
        eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);
        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
        uint256 beaconSharesBefore = delegationManager.operatorShares(staker, beaconChainETHStrategy);
        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(defaultOperator, staker, strategyMock, shares);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DepositScalingFactorUpdated(staker, strategyMock, 2e18);
        if (beaconShares > 0) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(defaultOperator, staker, beaconChainETHStrategy, uint256(beaconShares));
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit DepositScalingFactorUpdated(staker, beaconChainETHStrategy, 2e18);
        }
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
        uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);
        uint256 beaconSharesAfter = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
        if (beaconShares <= 0) {
            assertEq(
                beaconSharesBefore,
                beaconSharesAfter,
                "operator beaconchain shares should not have increased with negative shares"
            );
        } else {
            assertEq(
                beaconSharesBefore + uint256(beaconShares),
                beaconSharesAfter,
                "operator beaconchain shares not increased correctly"
            );
        }
        assertEq(operatorSharesBefore + shares, operatorSharesAfter, "operator shares not increased correctly");
        assertTrue(delegationManager.isOperator(defaultOperator), "staker not registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );

        IStrategy[] memory strategiesToCheck = new IStrategy[](2);
        strategiesToCheck[0] = beaconChainETHStrategy;
        strategiesToCheck[1] = strategyMock;
        (uint256[] memory withdrawableShares) = delegationManager.getWithdrawableShares(staker, strategiesToCheck);
        if (beaconShares > 0) {
            assertEq(withdrawableShares[0], uint256(beaconShares), "staker beacon chain shares not set correctly");
        } else {
            assertEq(withdrawableShares[0], 0, "staker beacon chain shares not set correctly");
        }
        assertEq(withdrawableShares[1], shares, "staker strategy shares not set correctly");
    }

    /**
     * @notice `staker` delegates to a operator who does not require any signature verification similar to test above.
     * In this scenario, staker doesn't have any delegatable shares and operator shares should not increase. Staker
     * should still be correctly delegated to the operator after the call.
     */
    function testFuzz_OperatorWhoAcceptsAllStakers_ZeroDelegatableShares(
        address staker,
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 salt,
        uint64 operatorMagnitude
    ) public filterFuzzedAddressInputs(staker) {
        // Bound magnitude
        operatorMagnitude = uint64(bound(operatorMagnitude, 1, uint64(WAD)));

        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        _registerOperatorWithBaseDetails(defaultOperator);

        _setOperatorMagnitude(defaultOperator, strategyMock, operatorMagnitude);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();

        assertTrue(delegationManager.isOperator(defaultOperator), "staker not registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
    }

    /**
     * @notice Like `testDelegateToOperatorWhoRequiresECDSASignature` but using an invalid expiry on purpose and checking that reversion occurs
     */
    function testFuzz_Revert_WhenOperatorWhoRequiresECDSASignature_ExpiredDelegationApproverSignature(
        address staker,
        bytes32 salt,
        uint256 expiry
    ) public filterFuzzedAddressInputs(staker) {
        // roll to a very late timestamp
        skip(type(uint256).max / 2);
        // filter to only *invalid* `expiry` values
        expiry = bound(expiry, 0, block.timestamp - 1);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        _registerOperatorWithDelegationApprover(defaultOperator);

        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(
            delegationSignerPrivateKey,
            staker,
            defaultOperator,
            salt,
            expiry
        );

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectRevert(IDelegationManagerErrors.SignatureExpired.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
    }

    /**
     * @notice Like `testDelegateToOperatorWhoRequiresECDSASignature` but undelegating after delegating and trying the same approveSignature
     * and checking that reversion occurs with the same salt
     */
    function testFuzz_Revert_WhenOperatorWhoRequiresECDSASignature_PreviouslyUsedSalt(
        address staker,
        bytes32 salt,
        uint256 expiry
    ) public filterFuzzedAddressInputs(staker) {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);

        // filter inputs, since this will fail when the staker is already registered as an operator
        // staker also must not be the delegationApprover so that signature verification process takes place
        cheats.assume(staker != defaultOperator);
        cheats.assume(staker != defaultApprover);

        _registerOperatorWithDelegationApprover(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(
            delegationSignerPrivateKey,
            staker,
            defaultOperator,
            salt,
            expiry
        );

        // delegate from the `staker` to the operator, undelegate, and then try to delegate again with same approversalt
        // to check that call reverts
        cheats.startPrank(staker);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        assertTrue(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent not spent?"
        );
        delegationManager.undelegate(staker);
        cheats.expectRevert(IDelegationManagerErrors.SaltSpent.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
    }

    /**
     * @notice Like `testDelegateToOperatorWhoRequiresECDSASignature` but using an incorrect signature on purpose and checking that reversion occurs
     */
    function testFuzz_Revert_WhenOperatorWhoRequiresECDSASignature_WithBadSignature(
        address staker,
        uint256 expiry
    ) public filterFuzzedAddressInputs(staker) {
        // filter to only valid `expiry` values
        expiry = bound(expiry, block.timestamp + 1, type(uint256).max);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator && staker != defaultApprover);

        _registerOperatorWithDelegationApprover(defaultOperator);

        // calculate the signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        approverSignatureAndExpiry.expiry = expiry;
        {
            bytes32 digestHash = delegationManager.calculateDelegationApprovalDigestHash(
                staker,
                defaultOperator,
                delegationManager.delegationApprover(defaultOperator),
                emptySalt,
                expiry
            );
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(delegationSignerPrivateKey, digestHash);
            // mess up the signature by flipping v's parity
            v = (v == 27 ? 28 : 27);
            approverSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }

        // try to delegate from the `staker` to the operator, and check reversion
        cheats.startPrank(staker);
        cheats.expectRevert(ISignatureUtils.InvalidSignature.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, emptySalt);
        cheats.stopPrank();
    }

    /**
     * @notice `staker` delegates to an operator who requires signature verification through an EOA (i.e. the operator’s `delegationApprover` address is set to a nonzero EOA)
     * via the `staker` calling `DelegationManager.delegateTo`
     * The function should pass *only with a valid ECDSA signature from the `delegationApprover`, OR if called by the operator or their delegationApprover themselves
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_OperatorWhoRequiresECDSASignature(
        address staker,
        bytes32 salt,
        uint256 expiry
    ) public filterFuzzedAddressInputs(staker) {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        _registerOperatorWithDelegationApprover(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(
            delegationSignerPrivateKey,
            staker,
            defaultOperator,
            salt,
            expiry
        );

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();

        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        if (staker == delegationManager.delegationApprover(defaultOperator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertFalse(
                delegationManager.delegationApproverSaltIsSpent(
                    delegationManager.delegationApprover(defaultOperator),
                    salt
                ),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(
                    delegationManager.delegationApprover(defaultOperator),
                    salt
                ),
                "salt somehow spent not spent?"
            );
        }
    }

    /**
     * @notice `staker` delegates to an operator who requires signature verification through an EOA (i.e. the operator’s `delegationApprover` address is set to a nonzero EOA)
     * via the `staker` calling `DelegationManager.delegateTo`
     * The function should pass *only with a valid ECDSA signature from the `delegationApprover`, OR if called by the operator or their delegationApprover themselves
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Operator shares should increase by the amount of shares delegated
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_OperatorWhoRequiresECDSASignature_StrategyManagerShares(
        address staker,
        bytes32 salt,
        uint256 expiry,
        uint128 shares
    ) public filterFuzzedAddressInputs(staker) {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        _registerOperatorWithDelegationApprover(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(
            delegationSignerPrivateKey,
            staker,
            defaultOperator,
            salt,
            expiry
        );

        // Set staker shares in StrategyManager
        IStrategy[] memory strategiesToReturn = new IStrategy[](1);
        strategiesToReturn[0] = strategyMock;
        uint256[] memory sharesToReturn = new uint256[](1);
        sharesToReturn[0] = shares;
        strategyManagerMock.setDeposits(staker, strategiesToReturn, sharesToReturn);
        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
        uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);
        assertEq(operatorSharesBefore + shares, operatorSharesAfter, "operator shares not increased correctly");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        if (staker == delegationManager.delegationApprover(defaultOperator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertFalse(
                delegationManager.delegationApproverSaltIsSpent(
                    delegationManager.delegationApprover(defaultOperator),
                    salt
                ),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(
                    delegationManager.delegationApprover(defaultOperator),
                    salt
                ),
                "salt somehow spent not spent?"
            );
        }
    }

    /**
     * @notice `staker` delegates to an operator who requires signature verification through an EOA (i.e. the operator’s `delegationApprover` address is set to a nonzero EOA)
     * via the `staker` calling `DelegationManager.delegateTo`
     * The function should pass *only with a valid ECDSA signature from the `delegationApprover`, OR if called by the operator or their delegationApprover themselves
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Operator beaconShares should increase by the amount of shares delegated if beaconShares > 0
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_OperatorWhoRequiresECDSASignature_BeaconChainStrategyShares(
        address staker,
        bytes32 salt,
        uint256 expiry,
        int256 beaconShares
    ) public filterFuzzedAddressInputs(staker) {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        _registerOperatorWithDelegationApprover(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(
            delegationSignerPrivateKey,
            staker,
            defaultOperator,
            salt,
            expiry
        );

        // Set staker shares in BeaconChainStrategy
        eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);
        uint256 beaconSharesBefore = delegationManager.operatorShares(staker, beaconChainETHStrategy);
        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        if (beaconShares > 0) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(defaultOperator, staker, beaconChainETHStrategy, uint256(beaconShares));
        }
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
        uint256 beaconSharesAfter = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
        if (beaconShares <= 0) {
            assertEq(
                beaconSharesBefore,
                beaconSharesAfter,
                "operator beaconchain shares should not have increased with negative shares"
            );
        } else {
            assertEq(
                beaconSharesBefore + uint256(beaconShares),
                beaconSharesAfter,
                "operator beaconchain shares not increased correctly"
            );
        }
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        if (staker == delegationManager.delegationApprover(defaultOperator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertFalse(
                delegationManager.delegationApproverSaltIsSpent(
                    delegationManager.delegationApprover(defaultOperator),
                    salt
                ),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(
                    delegationManager.delegationApprover(defaultOperator),
                    salt
                ),
                "salt somehow spent not spent?"
            );
        }
    }

    /**
     * @notice `staker` delegates to an operator who requires signature verification through an EOA (i.e. the operator’s `delegationApprover` address is set to a nonzero EOA)
     * via the `staker` calling `DelegationManager.delegateTo`
     * The function should pass *only with a valid ECDSA signature from the `delegationApprover`, OR if called by the operator or their delegationApprover themselves
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Operator beaconshares should increase by the amount of beaconShares delegated if beaconShares > 0
     * Operator strategy manager shares should icnrease by amount of shares
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_OperatorWhoRequiresECDSASignature_BeaconChainAndStrategyManagerShares(
        address staker,
        bytes32 salt,
        uint256 expiry,
        int256 beaconShares,
        uint128 shares
    ) public filterFuzzedAddressInputs(staker) {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);

        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);
        _registerOperatorWithDelegationApprover(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(
            delegationSignerPrivateKey,
            staker,
            defaultOperator,
            salt,
            expiry
        );

        // Set staker shares in BeaconChainStrategy and StrategyMananger
        {
            IStrategy[] memory strategiesToReturn = new IStrategy[](1);
            strategiesToReturn[0] = strategyMock;
            uint256[] memory sharesToReturn = new uint256[](1);
            sharesToReturn[0] = shares;
            strategyManagerMock.setDeposits(staker, strategiesToReturn, sharesToReturn);
            eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);
        }
        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
        uint256 beaconSharesBefore = delegationManager.operatorShares(staker, beaconChainETHStrategy);
        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(defaultOperator, staker, strategyMock, shares);
        if (beaconShares > 0) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(defaultOperator, staker, beaconChainETHStrategy, uint256(beaconShares));
        }
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
        uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);
        uint256 beaconSharesAfter = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
        if (beaconShares <= 0) {
            assertEq(
                beaconSharesBefore,
                beaconSharesAfter,
                "operator beaconchain shares should not have increased with negative shares"
            );
        } else {
            assertEq(
                beaconSharesBefore + uint256(beaconShares),
                beaconSharesAfter,
                "operator beaconchain shares not increased correctly"
            );
        }
        assertEq(operatorSharesBefore + shares, operatorSharesAfter, "operator shares not increased correctly");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        if (staker == delegationManager.delegationApprover(defaultOperator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertFalse(
                delegationManager.delegationApproverSaltIsSpent(
                    delegationManager.delegationApprover(defaultOperator),
                    salt
                ),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(
                    delegationManager.delegationApprover(defaultOperator),
                    salt
                ),
                "salt somehow spent not spent?"
            );
        }
    }

    /**
     * @notice delegateTo test with operator's delegationApprover address set to a contract address
     * and check that reversion occurs when the signature is expired
     */
    function testFuzz_Revert_WhenOperatorWhoRequiresEIP1271Signature_ExpiredDelegationApproverSignature(
        address staker,
        uint256 expiry
    ) public filterFuzzedAddressInputs(staker) {
        // roll to a very late timestamp
        skip(type(uint256).max / 2);
        // filter to only *invalid* `expiry` values
        expiry = bound(expiry, 0, block.timestamp - 1);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        _registerOperatorWithDelegationApprover(defaultOperator);

        // create the signature struct
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        approverSignatureAndExpiry.expiry = expiry;

        // try to delegate from the `staker` to the operator, and check reversion
        cheats.startPrank(staker);
        cheats.expectRevert(IDelegationManagerErrors.SignatureExpired.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, emptySalt);
        cheats.stopPrank();
    }

    /**
     * @notice delegateTo test with operator's delegationApprover address set to a contract address
     * and check that reversion occurs when the signature approverSalt is already used.
     * Performed by delegating to operator, undelegating, and trying to reuse the same signature
     */
    function testFuzz_Revert_WhenOperatorWhoRequiresEIP1271Signature_PreviouslyUsedSalt(
        address staker,
        bytes32 salt,
        uint256 expiry
    ) public filterFuzzedAddressInputs(staker) {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);

        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        ERC1271WalletMock wallet = _registerOperatorWith1271DelegationApprover(defaultOperator);
        cheats.assume(staker != address(wallet) && staker != defaultOperator);

        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(
            delegationSignerPrivateKey,
            staker,
            defaultOperator,
            salt,
            expiry
        );

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        delegationManager.undelegate(staker);
        // Reusing same signature should revert with salt already being used
        cheats.expectRevert(IDelegationManagerErrors.SaltSpent.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
    }

    /**
     * @notice delegateTo test with operator's delegationApprover address set to a contract address that
     * is non compliant with EIP1271
     */
    function testFuzz_Revert_WhenOperatorWhoRequiresEIP1271Signature_NonCompliantWallet(
        address staker,
        uint256 expiry
    ) public filterFuzzedAddressInputs(staker) {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);

        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        // deploy a ERC1271MaliciousMock contract that will return an incorrect value when called
        ERC1271MaliciousMock wallet = new ERC1271MaliciousMock();

        // filter fuzzed input, since otherwise we can get a flaky failure here. if the caller itself is the 'delegationApprover'
        // then we don't even trigger the signature verification call, so we won't get a revert as expected
        cheats.assume(staker != address(wallet));

        IDelegationManagerTypes.OperatorDetails memory operatorDetails = IDelegationManagerTypes.OperatorDetails({
            __deprecated_earningsReceiver: defaultOperator,
            delegationApprover: address(wallet),
            __deprecated_stakerOptOutWindowBlocks: 0
        });
        _registerOperator(defaultOperator, operatorDetails, emptyStringForMetadataURI);

        // create the signature struct
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        approverSignatureAndExpiry.expiry = expiry;

        // try to delegate from the `staker` to the operator, and check reversion
        cheats.startPrank(staker);
        // because the ERC1271MaliciousMock contract returns the wrong amount of data, we get a low-level "EvmError: Revert" message here rather than the error message bubbling up
        cheats.expectRevert();
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, emptySalt);
        cheats.stopPrank();
    }

    /**
     * @notice delegateTo test with operator's delegationApprover address set to a contract address that
     * returns a value other than the EIP1271 "magic bytes" and checking that reversion occurs appropriately
     */
    function testFuzz_Revert_WhenOperatorWhoRequiresEIP1271Signature_IsValidSignatureFails(
        address staker,
        bytes32 salt,
        uint256 expiry
    ) public filterFuzzedAddressInputs(staker) {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);

        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        // deploy a ERC1271WalletMock contract that will return an incorrect value when called
        // owner is the 0 address
        ERC1271WalletMock wallet = new ERC1271WalletMock(address(1));

        // filter fuzzed input, since otherwise we can get a flaky failure here. if the caller itself is the 'delegationApprover'
        // then we don't even trigger the signature verification call, so we won't get a revert as expected
        cheats.assume(staker != address(wallet));

        IDelegationManagerTypes.OperatorDetails memory operatorDetails = IDelegationManagerTypes.OperatorDetails({
            __deprecated_earningsReceiver: defaultOperator,
            delegationApprover: address(wallet),
            __deprecated_stakerOptOutWindowBlocks: 0
        });
        _registerOperator(defaultOperator, operatorDetails, emptyStringForMetadataURI);

        // calculate the delegationSigner's but this is not the correct signature from the wallet contract
        // since the wallet owner is address(1)
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(
            delegationSignerPrivateKey,
            staker,
            defaultOperator,
            salt,
            expiry
        );

        // try to delegate from the `staker` to the operator, and check reversion
        cheats.startPrank(staker);
        // Signature should fail as the wallet will not return EIP1271_MAGICVALUE
        cheats.expectRevert(ISignatureUtils.InvalidSignature.selector);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, emptySalt);
        cheats.stopPrank();
    }

    /**
     * @notice `staker` delegates to an operator who requires signature verification through an EIP1271-compliant contract (i.e. the operator’s `delegationApprover` address is
     * set to a nonzero and code-containing address) via the `staker` calling `DelegationManager.delegateTo`
     * The function uses OZ's ERC1271WalletMock contract, and thus should pass *only when a valid ECDSA signature from the `owner` of the ERC1271WalletMock contract,
     * OR if called by the operator or their delegationApprover themselves
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_OperatorWhoRequiresEIP1271Signature(
        address staker,
        bytes32 salt,
        uint256 expiry
    ) public filterFuzzedAddressInputs(staker) {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);

        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        _registerOperatorWith1271DelegationApprover(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(
            delegationSignerPrivateKey,
            staker,
            defaultOperator,
            salt,
            expiry
        );

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();

        assertTrue(delegationManager.isDelegated(staker), "staker not delegated correctly");
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // check that the nonce incremented appropriately
        if (staker == defaultOperator || staker == delegationManager.delegationApprover(defaultOperator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertFalse(
                delegationManager.delegationApproverSaltIsSpent(
                    delegationManager.delegationApprover(defaultOperator),
                    salt
                ),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(
                    delegationManager.delegationApprover(defaultOperator),
                    salt
                ),
                "salt somehow spent not spent?"
            );
        }
    }
}

contract DelegationManagerUnitTests_delegateToBySignature is DelegationManagerUnitTests {
    function test_revert_paused() public {
        cheats.prank(defaultOperator);
        delegationManager.registerAsOperator(
            IDelegationManagerTypes.OperatorDetails({
                __deprecated_earningsReceiver: defaultOperator,
                delegationApprover: address(0),
                __deprecated_stakerOptOutWindowBlocks: 0
            }),
            0,
            emptyStringForMetadataURI
        );

        // set the pausing flag
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_NEW_DELEGATION);

        uint256 expiry = type(uint256).max;
        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(
            stakerPrivateKey,
            defaultOperator,
            expiry
        );
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.delegateToBySignature(
            defaultStaker,
            defaultOperator,
            stakerSignatureAndExpiry,
            approverSignatureAndExpiry,
            emptySalt
        );
    }

    /// @notice Checks that `DelegationManager.delegateToBySignature` reverts if the staker's signature has expired
    function testFuzz_Revert_WhenStakerSignatureExpired(
        address staker,
        address operator,
        uint256 expiry,
        bytes memory signature
    ) public filterFuzzedAddressInputs(staker) filterFuzzedAddressInputs(operator) {
        expiry = bound(expiry, 0, block.timestamp - 1);
        cheats.expectRevert(IDelegationManagerErrors.SignatureExpired.selector);
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry = ISignatureUtils.SignatureWithExpiry({
            signature: signature,
            expiry: expiry
        });
        delegationManager.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, emptySalt);
    }

    /// @notice Checks that `DelegationManager.delegateToBySignature` reverts if the staker's ECDSA signature verification fails
    function test_Revert_EOAStaker_WhenStakerSignatureVerificationFails() public {
        address invalidStaker = address(1000);
        address caller = address(2000);
        uint256 expiry = type(uint256).max;

        _registerOperatorWithBaseDetails(defaultOperator);

        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(
            stakerPrivateKey,
            defaultOperator,
            expiry
        );

        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        // Should revert from invalid signature as staker is not set as the address of signer
        cheats.startPrank(caller);
        cheats.expectRevert(ISignatureUtils.InvalidSignature.selector);
        // use an empty approver signature input since none is needed / the input is unchecked
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateToBySignature(
            invalidStaker,
            defaultOperator,
            stakerSignatureAndExpiry,
            approverSignatureAndExpiry,
            emptySalt
        );
        cheats.stopPrank();
    }

    /// @notice Checks that `DelegationManager.delegateToBySignature` reverts if the staker's contract signature verification fails
    function test_Revert_ERC1271Staker_WhenStakerSignatureVerficationFails() public {
        address staker = address(new ERC1271WalletMock(address(1)));
        address caller = address(2000);
        uint256 expiry = type(uint256).max;

        _registerOperatorWithBaseDetails(defaultOperator);

        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(
            stakerPrivateKey,
            defaultOperator,
            expiry
        );

        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        // Should revert from invalid signature as staker is not set as the address of signer
        cheats.startPrank(caller);
        cheats.expectRevert(ISignatureUtils.InvalidSignature.selector);
        // use an empty approver signature input since none is needed / the input is unchecked
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateToBySignature(
            staker,
            defaultOperator,
            stakerSignatureAndExpiry,
            approverSignatureAndExpiry,
            emptySalt
        );
        cheats.stopPrank();
    }

    // TODO: Fix test from Delegation.t.sol
    //     /// @notice  tries delegating using a wallet that does not comply with EIP 1271
    // function testDelegateToBySignature_WithContractWallet_NonconformingWallet(
    //     address operator,
    //     uint96 ethAmount,
    //     uint96 eigenAmount,
    //     uint8 v,
    //     bytes32 r,
    //     bytes32 s
    // ) public fuzzedAddress(operator) {
    //     address staker = cheats.addr(PRIVATE_KEY);

    //     // deploy non ERC1271-compliant wallet for staker to use
    //     cheats.startPrank(staker);
    //     ERC1271MaliciousMock wallet = new ERC1271MaliciousMock();
    //     cheats.stopPrank();
    //     staker = address(wallet);

    //     _registerOperatorAndDepositFromStaker(operator, staker, ethAmount, eigenAmount);

    //     cheats.assume(staker != operator);

    //     bytes memory signature = abi.encodePacked(r, s, v);

    //     cheats.expectRevert();
    //     ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry = ISignatureUtils.SignatureWithExpiry({
    //         signature: signature,
    //         expiry: type(uint256).max
    //     });
    //     delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, bytes32(0));
    // }

    /// @notice Checks that `DelegationManager.delegateToBySignature` reverts when the staker is already delegated
    function test_Revert_Staker_WhenActivelyDelegated() public {
        address staker = cheats.addr(stakerPrivateKey);
        address caller = address(2000);
        uint256 expiry = type(uint256).max;

        _registerOperatorWithBaseDetails(defaultOperator);

        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(
            stakerPrivateKey,
            defaultOperator,
            expiry
        );

        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        // Should revert as `staker` has already delegated to `operator`
        cheats.startPrank(caller);
        cheats.expectRevert(IDelegationManagerErrors.ActivelyDelegated.selector);
        // use an empty approver signature input since none is needed / the input is unchecked
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateToBySignature(
            staker,
            defaultOperator,
            stakerSignatureAndExpiry,
            approverSignatureAndExpiry,
            emptySalt
        );
        cheats.stopPrank();
    }

    /// @notice Checks that `delegateToBySignature` reverts when operator is not registered after successful staker signature verification
    function test_Revert_EOAStaker_OperatorNotRegistered() public {
        address staker = cheats.addr(stakerPrivateKey);
        address caller = address(2000);
        uint256 expiry = type(uint256).max;

        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(
            stakerPrivateKey,
            defaultOperator,
            expiry
        );

        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        // Should revert as `operator` is not registered
        cheats.startPrank(caller);
        cheats.expectRevert(IDelegationManagerErrors.OperatorNotRegistered.selector);
        // use an empty approver signature input since none is needed / the input is unchecked
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateToBySignature(
            staker,
            defaultOperator,
            stakerSignatureAndExpiry,
            approverSignatureAndExpiry,
            emptySalt
        );
        cheats.stopPrank();
    }

    /**
     * @notice Checks that `DelegationManager.delegateToBySignature` reverts if the delegationApprover's signature has expired
     * after successful staker signature verification
     */
    function testFuzz_Revert_WhenDelegationApproverSignatureExpired(
        address caller,
        uint256 stakerExpiry,
        uint256 delegationApproverExpiry
    ) public filterFuzzedAddressInputs(caller) {
        cheats.assume(caller != defaultOperator);
        
        // roll to a very late timestamp
        skip(type(uint256).max / 2);

        // filter to only valid `stakerExpiry` values
        stakerExpiry = bound(stakerExpiry, block.timestamp + 1, type(uint256).max);
        // filter to only *invalid* `delegationApproverExpiry` values
        delegationApproverExpiry = bound(delegationApproverExpiry, 0, block.timestamp - 1);

        console.log("timestamp: %s", block.timestamp);
        console.log(stakerExpiry);
        console.log(delegationApproverExpiry);


        _registerOperatorWithDelegationApprover(defaultOperator);

        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(
            delegationSignerPrivateKey,
            defaultStaker,
            defaultOperator,
            emptySalt,
            delegationApproverExpiry
        );

        // calculate the staker signature
        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(
            stakerPrivateKey,
            defaultOperator,
            stakerExpiry
        );

        // try delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`, and check for reversion
        cheats.startPrank(caller);
        cheats.expectRevert(IDelegationManagerErrors.SignatureExpired.selector);
        delegationManager.delegateToBySignature(
            defaultStaker,
            defaultOperator,
            stakerSignatureAndExpiry,
            approverSignatureAndExpiry,
            emptySalt
        );
        cheats.stopPrank();
    }

    /**
     * @notice `staker` becomes delegated to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * via the `caller` calling `DelegationManager.delegateToBySignature`
     * The function should pass with any `operatorSignature` input (since it should be unused)
     * The function should pass only with a valid `stakerSignatureAndExpiry` input
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * BeaconChainStrategy and StrategyManager operator shares should increase for operator
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_EOAStaker_OperatorWhoAcceptsAllStakers(
        address caller,
        uint256 expiry,
        int256 beaconShares,
        uint256 shares
    ) public filterFuzzedAddressInputs(caller) {
        cheats.assume(expiry >= block.timestamp);
        cheats.assume(shares > 0);

        _registerOperatorWithBaseDetails(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                emptySalt
            ),
            "salt somehow spent too early?"
        );
        {
            // Set staker shares in BeaconChainStrategy and StrategyMananger
            IStrategy[] memory strategiesToReturn = new IStrategy[](1);
            strategiesToReturn[0] = strategyMock;
            uint256[] memory sharesToReturn = new uint256[](1);
            sharesToReturn[0] = shares;
            strategyManagerMock.setDeposits(defaultStaker, strategiesToReturn, sharesToReturn);
            eigenPodManagerMock.setPodOwnerShares(defaultStaker, beaconShares);
        }

        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
        uint256 beaconSharesBefore = delegationManager.operatorShares(defaultStaker, beaconChainETHStrategy);
        // fetch the staker's current nonce
        uint256 currentStakerNonce = delegationManager.stakerNonce(defaultStaker);
        // calculate the staker signature
        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(
            stakerPrivateKey,
            defaultOperator,
            expiry
        );

        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(defaultStaker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(defaultOperator, defaultStaker, strategyMock, shares);
        if (beaconShares > 0) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(defaultOperator, defaultStaker, beaconChainETHStrategy, uint256(beaconShares));
        }
        _delegateToBySignatureOperatorWhoAcceptsAllStakers(
            defaultStaker,
            caller,
            defaultOperator,
            stakerSignatureAndExpiry,
            emptySalt
        );

        // Check operator shares increases
        uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);
        uint256 beaconSharesAfter = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
        if (beaconShares <= 0) {
            assertEq(
                beaconSharesBefore,
                beaconSharesAfter,
                "operator beaconchain shares should not have increased with negative shares"
            );
        } else {
            assertEq(
                beaconSharesBefore + uint256(beaconShares),
                beaconSharesAfter,
                "operator beaconchain shares not increased correctly"
            );
        }
        assertEq(operatorSharesBefore + shares, operatorSharesAfter, "operator shares not increased correctly");
        // check all the delegation status changes
        assertTrue(delegationManager.isDelegated(defaultStaker), "staker not delegated correctly");
        assertEq(
            delegationManager.delegatedTo(defaultStaker),
            defaultOperator,
            "staker delegated to the wrong address"
        );
        assertFalse(delegationManager.isOperator(defaultStaker), "staker incorrectly registered as operator");
        // check that the staker nonce incremented appropriately
        assertEq(
            delegationManager.stakerNonce(defaultStaker),
            currentStakerNonce + 1,
            "staker nonce did not increment"
        );
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                emptySalt
            ),
            "salt somehow spent too incorrectly?"
        );
    }

    /**
     * @notice `staker` becomes delegated to an operator who requires signature verification through an EOA (i.e. the operator’s `delegationApprover` address is set to a nonzero EOA)
     * via the `caller` calling `DelegationManager.delegateToBySignature`
     * The function should pass *only with a valid ECDSA signature from the `delegationApprover`, OR if called by the operator or their delegationApprover themselves
     * AND with a valid `stakerSignatureAndExpiry` input
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * BeaconChainStrategy and StrategyManager operator shares should increase for operator
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_EOAStaker_OperatorWhoRequiresECDSASignature(
        address caller,
        bytes32 salt,
        uint256 expiry,
        int256 beaconShares,
        uint256 shares
    ) public filterFuzzedAddressInputs(caller) {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);
        cheats.assume(shares > 0);

        _registerOperatorWithDelegationApprover(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
        {
            // Set staker shares in BeaconChainStrategy and StrategyMananger
            IStrategy[] memory strategiesToReturn = new IStrategy[](1);
            strategiesToReturn[0] = strategyMock;
            uint256[] memory sharesToReturn = new uint256[](1);
            sharesToReturn[0] = shares;
            strategyManagerMock.setDeposits(defaultStaker, strategiesToReturn, sharesToReturn);
            eigenPodManagerMock.setPodOwnerShares(defaultStaker, beaconShares);
        }

        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
        uint256 beaconSharesBefore = delegationManager.operatorShares(defaultStaker, beaconChainETHStrategy);
        // fetch the staker's current nonce
        uint256 currentStakerNonce = delegationManager.stakerNonce(defaultStaker);
        // calculate the staker signature
        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(
            stakerPrivateKey,
            defaultOperator,
            expiry
        );

        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(defaultStaker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(defaultOperator, defaultStaker, strategyMock, shares);
        if (beaconShares > 0) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(defaultOperator, defaultStaker, beaconChainETHStrategy, uint256(beaconShares));
        }
        _delegateToBySignatureOperatorWhoRequiresSig(
            defaultStaker,
            caller,
            defaultOperator,
            stakerSignatureAndExpiry,
            salt
        );
        {
            // Check operator shares increases
            uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);
            uint256 beaconSharesAfter = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
            if (beaconShares <= 0) {
                assertEq(
                    beaconSharesBefore,
                    beaconSharesAfter,
                    "operator beaconchain shares should not have increased with negative shares"
                );
            } else {
                assertEq(
                    beaconSharesBefore + uint256(beaconShares),
                    beaconSharesAfter,
                    "operator beaconchain shares not increased correctly"
                );
            }
            assertEq(operatorSharesBefore + shares, operatorSharesAfter, "operator shares not increased correctly");
        }
        assertTrue(delegationManager.isDelegated(defaultStaker), "staker not delegated correctly");
        assertEq(
            delegationManager.delegatedTo(defaultStaker),
            defaultOperator,
            "staker delegated to the wrong address"
        );
        assertFalse(delegationManager.isOperator(defaultStaker), "staker incorrectly registered as operator");

        // check that the delegationApprover nonce incremented appropriately
        if (caller == defaultOperator || caller == delegationManager.delegationApprover(defaultOperator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertFalse(
                delegationManager.delegationApproverSaltIsSpent(
                    delegationManager.delegationApprover(defaultOperator),
                    salt
                ),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(
                    delegationManager.delegationApprover(defaultOperator),
                    salt
                ),
                "salt somehow spent not spent?"
            );
        }

        // check that the staker nonce incremented appropriately
        assertEq(
            delegationManager.stakerNonce(defaultStaker),
            currentStakerNonce + 1,
            "staker nonce did not increment"
        );
    }

    /**
     * @notice `staker` becomes delegated to an operatorwho requires signature verification through an EIP1271-compliant contract (i.e. the operator’s `delegationApprover` address is
     * set to a nonzero and code-containing address) via the `caller` calling `DelegationManager.delegateToBySignature`
     * The function uses OZ's ERC1271WalletMock contract, and thus should pass *only when a valid ECDSA signature from the `owner` of the ERC1271WalletMock contract,
     * OR if called by the operator or their delegationApprover themselves
     * AND with a valid `stakerSignatureAndExpiry` input
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_EOAStaker_OperatorWhoRequiresEIP1271Signature(
        address caller,
        bytes32 salt,
        uint256 expiry,
        int256 beaconShares,
        uint256 shares
    ) public filterFuzzedAddressInputs(caller) {
        cheats.assume(expiry >= block.timestamp);
        cheats.assume(shares > 0);

        _registerOperatorWith1271DelegationApprover(defaultOperator);

        // verify that the salt hasn't been used before
        assertFalse(
            delegationManager.delegationApproverSaltIsSpent(
                delegationManager.delegationApprover(defaultOperator),
                salt
            ),
            "salt somehow spent too early?"
        );
        {
            // Set staker shares in BeaconChainStrategy and StrategyMananger
            IStrategy[] memory strategiesToReturn = new IStrategy[](1);
            strategiesToReturn[0] = strategyMock;
            uint256[] memory sharesToReturn = new uint256[](1);
            sharesToReturn[0] = shares;
            strategyManagerMock.setDeposits(defaultStaker, strategiesToReturn, sharesToReturn);
            eigenPodManagerMock.setPodOwnerShares(defaultStaker, beaconShares);
        }

        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
        uint256 beaconSharesBefore = delegationManager.operatorShares(defaultStaker, beaconChainETHStrategy);
        // fetch the staker's current nonce
        uint256 currentStakerNonce = delegationManager.stakerNonce(defaultStaker);
        // calculate the staker signature
        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(
            stakerPrivateKey,
            defaultOperator,
            expiry
        );

        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(defaultStaker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(defaultOperator, defaultStaker, strategyMock, shares);
        if (beaconShares > 0) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(defaultOperator, defaultStaker, beaconChainETHStrategy, uint256(beaconShares));
        }
        _delegateToBySignatureOperatorWhoRequiresSig(
            defaultStaker,
            caller,
            defaultOperator,
            stakerSignatureAndExpiry,
            salt
        );

        {
            // Check operator shares increases
            uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);
            uint256 beaconSharesAfter = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
            if (beaconShares <= 0) {
                assertEq(
                    beaconSharesBefore,
                    beaconSharesAfter,
                    "operator beaconchain shares should not have increased with negative shares"
                );
            } else {
                assertEq(
                    beaconSharesBefore + uint256(beaconShares),
                    beaconSharesAfter,
                    "operator beaconchain shares not increased correctly"
                );
            }
            assertEq(operatorSharesBefore + shares, operatorSharesAfter, "operator shares not increased correctly");
        }
        assertTrue(delegationManager.isDelegated(defaultStaker), "staker not delegated correctly");
        assertEq(
            delegationManager.delegatedTo(defaultStaker),
            defaultOperator,
            "staker delegated to the wrong address"
        );
        assertFalse(delegationManager.isOperator(defaultStaker), "staker incorrectly registered as operator");

        // check that the delegationApprover nonce incremented appropriately
        if (caller == defaultOperator || caller == delegationManager.delegationApprover(defaultOperator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertFalse(
                delegationManager.delegationApproverSaltIsSpent(
                    delegationManager.delegationApprover(defaultOperator),
                    salt
                ),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(
                    delegationManager.delegationApprover(defaultOperator),
                    salt
                ),
                "salt somehow spent not spent?"
            );
        }

        // check that the staker nonce incremented appropriately
        assertEq(
            delegationManager.stakerNonce(defaultStaker),
            currentStakerNonce + 1,
            "staker nonce did not increment"
        );
    }

    /**
     * @notice Calls same delegateToBySignature test but with the staker address being a ERC1271WalletMock
     * Generates valid signatures from the staker to delegate to operator `defaultOperator`
     */
    function testFuzz_ERC1271Staker_OperatorWhoAcceptsAllStakers(
        address caller,
        uint256 expiry,
        int256 beaconShares,
        uint256 shares
    ) public filterFuzzedAddressInputs(caller) {
        defaultStaker = address(ERC1271WalletMock(cheats.addr(stakerPrivateKey)));
        testFuzz_EOAStaker_OperatorWhoAcceptsAllStakers(caller, expiry, beaconShares, shares);
    }

    /**
     * @notice Calls same delegateToBySignature test but with the staker address being a ERC1271WalletMock
     * Generates valid signatures from the staker to delegate to operator `defaultOperator` who has
     * a delegationApprover address set to a nonzero EOA
     */
    function testFuzz_ERC1271Staker_OperatorWhoRequiresECDSASignature(
        address caller,
        bytes32 salt,
        uint256 expiry,
        int256 beaconShares,
        uint256 shares
    ) public filterFuzzedAddressInputs(caller) {
        // Call same test but with the staker address being a ERC1271WalletMock
        defaultStaker = address(ERC1271WalletMock(cheats.addr(stakerPrivateKey)));
        testFuzz_EOAStaker_OperatorWhoRequiresECDSASignature(caller, salt, expiry, beaconShares, shares);
    }

    /**
     * @notice Calls same delegateToBySignature test but with the staker address being a ERC1271WalletMock
     * Generates valid signatures from the staker to delegate to operator `defaultOperator` who has
     * a delegationApprover address set to a nonzero ERC1271 compliant contract
     */
    function testFuzz_ERC1271Staker_OperatorWhoRequiresEIP1271Signature(
        address caller,
        bytes32 salt,
        uint256 expiry,
        int256 beaconShares,
        uint256 shares
    ) public filterFuzzedAddressInputs(caller) {
        // Call same test but with the staker address being a ERC1271WalletMock
        defaultStaker = address(ERC1271WalletMock(cheats.addr(stakerPrivateKey)));
        testFuzz_EOAStaker_OperatorWhoRequiresEIP1271Signature(caller, salt, expiry, beaconShares, shares);
    }
}

contract DelegationManagerUnitTests_ShareAdjustment is DelegationManagerUnitTests {
    // @notice Verifies that `DelegationManager.increaseDelegatedShares` reverts if not called by the StrategyManager nor EigenPodManager
    function testFuzz_increaseDelegatedShares_revert_invalidCaller(
        address invalidCaller,
        uint256 shares
    ) public filterFuzzedAddressInputs(invalidCaller) {
        cheats.assume(invalidCaller != address(strategyManagerMock));
        cheats.assume(invalidCaller != address(eigenPodManagerMock));
        cheats.assume(invalidCaller != address(eigenLayerProxyAdmin));

        cheats.expectRevert(IDelegationManagerErrors.OnlyStrategyManagerOrEigenPodManager.selector);
        delegationManager.increaseDelegatedShares(invalidCaller, strategyMock, 0, shares);
    }

    // @notice Verifies that there is no change in shares if the staker is not delegated
    function testFuzz_increaseDelegatedShares_noop(address staker) public {
        cheats.assume(staker != defaultOperator);
        _registerOperatorWithBaseDetails(defaultOperator);
        assertFalse(delegationManager.isDelegated(staker), "bad test setup");

        cheats.prank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, strategyMock, 0, 0);
        assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "shares should not have changed");
    }

    /**
     * @notice Verifies that `DelegationManager.increaseDelegatedShares` properly increases the delegated `shares` that the operator
     * who the `staker` is delegated to has in the strategy
     * @dev Checks that there is no change if the staker is not delegated
     */
    function testFuzz_increaseDelegatedShares(
        address staker,
        uint128 shares,
        bool delegateFromStakerToOperator
    ) public filterFuzzedAddressInputs(staker) {
        // filter inputs, since delegating to the operator will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // delegate from the `staker` to the operator *if `delegateFromStakerToOperator` is 'true'*
        if (delegateFromStakerToOperator) {
            _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        }

        uint256 _delegatedSharesBefore = delegationManager.operatorShares(
            delegationManager.delegatedTo(staker),
            strategyMock
        );

        if (delegationManager.isDelegated(staker)) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(defaultOperator, staker, strategyMock, shares);
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit DepositScalingFactorUpdated(staker, strategyMock, WAD);
        }

        cheats.prank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, strategyMock, 0, shares);

        uint256 delegatedSharesAfter = delegationManager.operatorShares(
            delegationManager.delegatedTo(staker),
            strategyMock
        );

        if (delegationManager.isDelegated(staker)) {
            assertEq(
                delegatedSharesAfter,
                _delegatedSharesBefore + shares,
                "delegated shares did not increment correctly"
            );
        } else {
            assertEq(delegatedSharesAfter, _delegatedSharesBefore, "delegated shares incremented incorrectly");
            assertEq(_delegatedSharesBefore, 0, "nonzero shares delegated to zero address!");
        }
    }

    /**
     * @notice Verifies that `DelegationManager.increaseDelegatedShares` properly increases the delegated `shares` that the operator
     * who the `staker` is delegated to has in the strategy
     * @dev Checks that there is no change if the staker is not delegated
     */
    function testFuzz_increaseDelegatedShares_slashedOperator(
        address staker,
        uint128 shares,
        uint64 magnitude,
        bool delegateFromStakerToOperator
    ) public filterFuzzedAddressInputs(staker) { // remeber to filter fuzz inputs
        cheats.assume(staker != defaultOperator);
        magnitude = uint64(bound(magnitude, 1, WAD));

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);
        
        // Set operator magnitude
        _setOperatorMagnitude(defaultOperator, strategyMock, magnitude);


        // delegate from the `staker` to the operator *if `delegateFromStakerToOperator` is 'true'*
        if (delegateFromStakerToOperator) {
            _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        }

        uint256 _delegatedSharesBefore = delegationManager.operatorShares(
            delegationManager.delegatedTo(staker),
            strategyMock
        );

        if (delegationManager.isDelegated(staker)) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(defaultOperator, staker, strategyMock, shares);
            ssf.updateDepositScalingFactor(0, shares, magnitude);
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit DepositScalingFactorUpdated(staker, strategyMock, ssf.depositScalingFactor);
        }

        cheats.prank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, strategyMock, 0, shares);

        uint256 delegatedSharesAfter = delegationManager.operatorShares(
            delegationManager.delegatedTo(staker),
            strategyMock
        );

        if (delegationManager.isDelegated(staker)) {
            assertEq(
                delegatedSharesAfter,
                _delegatedSharesBefore + shares,
                "delegated shares did not increment correctly"
            );
        } else {
            assertEq(delegatedSharesAfter, _delegatedSharesBefore, "delegated shares incremented incorrectly");
            assertEq(_delegatedSharesBefore, 0, "nonzero shares delegated to zero address!");
        }
    }

    /**
     * @notice Verifies that `DelegationManager.increaseDelegatedShares` reverts when operator slashed 100% for a strategy
     * and the staker has deposits in that strategy
     */
    function testFuzz_Revert_increaseDelegatedShares_slashedOperator100Percent(
        address staker,
        uint128 shares
    ) public filterFuzzedAddressInputs(staker) { // remeber to filter fuzz inputs
        cheats.assume(staker != defaultOperator);

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);
        // Set operator magnitude
        _setOperatorMagnitude({
            operator: defaultOperator,
            strategy: strategyMock,
            magnitude: 0
        });
        // delegate from the `staker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);

        uint256 _delegatedSharesBefore = delegationManager.operatorShares(
            delegationManager.delegatedTo(staker),
            strategyMock
        );

        cheats.prank(address(strategyManagerMock));
        cheats.expectRevert(IDelegationManagerErrors.FullySlashed.selector);
        delegationManager.increaseDelegatedShares(staker, strategyMock, 0, shares);

        uint256 delegatedSharesAfter = delegationManager.operatorShares(
            delegationManager.delegatedTo(staker),
            strategyMock
        );

        assertEq(delegatedSharesAfter, _delegatedSharesBefore, "delegated shares incremented incorrectly");
        assertEq(_delegatedSharesBefore, 0, "nonzero shares delegated to zero address!");
    }

    /**
     * @notice Verifies that `DelegationManager.increaseDelegatedShares` reverts when operator slashed 100% for a strategy
     * and the staker has deposits in that strategy. In this test case, the staker was initially deposited and delegated
     * to the operator before the operator was slashed 100%.
     * @dev Checks that withdrawable shares after 100% slashed is 0
     * @dev Checks that as a staker, redepositing after 100% slashed reverts
     */
    function testFuzz_Revert_increaseDelegatedShares_slashedOperator100PercentWithExistingStaker(
        address staker,
        uint64 initialMagnitude,
        uint128 existingShares,
        uint128 shares
    ) public filterFuzzedAddressInputs(staker) { // remeber to filter fuzz inputs
        cheats.assume(staker != defaultOperator);
        cheats.assume(shares > existingShares);
        initialMagnitude = uint64(bound(initialMagnitude, 1, WAD));

        // 1. Register operator with initial operator magnitude and delegate staker to operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude({
            operator: defaultOperator,
            strategy: strategyMock,
            magnitude: initialMagnitude
        });
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        // 2. set staker initial shares and increase delegated shares
        cheats.prank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, strategyMock, 0, existingShares);
        IStrategy[] memory strategiesDeposited = new IStrategy[](1);
        strategiesDeposited[0] = strategyMock;
        uint256[] memory sharesToReturn = new uint256[](1);
        sharesToReturn[0] = existingShares;
        strategyManagerMock.setDeposits(staker, strategiesDeposited, sharesToReturn);
        uint256[] memory withdrawableShares = delegationManager.getWithdrawableShares(staker, strategiesDeposited);
        if (existingShares < 1e18) {
            // Check that withdrawable shares are within 1 share for amounts < 1e18
            assertApproxEqAbs(
                withdrawableShares[0],
                existingShares,
                1,
                "Existing shares should be set correctly"
            );
        } else {
            // check that withdrawable shares are rounded within  0.01% for amounts >= 1e18 
            assertApproxEqRel(
                withdrawableShares[0],
                existingShares,
                1e14,
                "Existing shares should be set correctly"
            );
        }
        // 3. Now set operator magnitude to 0 (100% slashed)
        _setOperatorMagnitude({
            operator: defaultOperator,
            strategy: strategyMock,
            magnitude: 0
        });

        // 4. Try to "redeposit" and expect a revert since strategy is 100% slashed
        // staker's withdrawable shares should also be 0 now
        cheats.prank(address(strategyManagerMock));
        cheats.expectRevert(IDelegationManagerErrors.FullySlashed.selector);
        delegationManager.increaseDelegatedShares(staker, strategyMock, existingShares, shares);

        withdrawableShares = delegationManager.getWithdrawableShares(staker, strategiesDeposited);
        assertEq(
            withdrawableShares[0],
            0,
            "All existing shares should be slashed"
        );
    }

    /**
     * @notice Verifies that `DelegationManager.increaseDelegatedShares` doesn't revert when operator slashed 100% for a strategy
     * and the staker has deposits in a separate strategy
     */
    function testFuzz_increaseDelegatedShares_slashedOperator100Percent(
        address staker,
        uint128 shares,
        uint64 magnitude,
        IStrategy strategy
    ) public filterFuzzedAddressInputs(staker) { // remeber to filter fuzz inputs
        cheats.assume(staker != defaultOperator);
        cheats.assume(address(strategy) != address(strategyMock));

        magnitude = uint64(bound(magnitude, 1, WAD));

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);
        // Set operator magnitude for 100% slashed strategy
        _setOperatorMagnitude({
            operator: defaultOperator,
            strategy: strategyMock,
            magnitude: 0
        });
        // Set operator magnitude for non-100% slashed strategy
        _setOperatorMagnitude({
            operator: defaultOperator,
            strategy: strategy,
            magnitude: magnitude
        });
        // delegate from the `staker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);

        uint256 _delegatedSharesBefore = delegationManager.operatorShares(
            delegationManager.delegatedTo(staker),
            strategy
        );

        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(defaultOperator, staker, strategy, shares);
        ssf.updateDepositScalingFactor(0, shares, magnitude);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DepositScalingFactorUpdated(staker, strategy, ssf.depositScalingFactor);

        cheats.prank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, strategy, 0, shares);

        uint256 delegatedSharesAfter = delegationManager.operatorShares(
            delegationManager.delegatedTo(staker),
            strategy
        );

        assertEq(
            delegatedSharesAfter,
            _delegatedSharesBefore + shares,
            "delegated shares did not increment correctly"
        );
    }

    /// @notice Verifies that `DelegationManager.decreaseOperatorShares` reverts if not called by the AllocationManager
    function testFuzz_decreaseOperatorShares_revert_invalidCaller(
        address invalidCaller
    ) public filterFuzzedAddressInputs(invalidCaller) {
        cheats.assume(invalidCaller != address(allocationManagerMock));

        cheats.startPrank(invalidCaller);
        cheats.expectRevert(IDelegationManagerErrors.OnlyAllocationManager.selector);
        delegationManager.decreaseOperatorShares(invalidCaller, strategyMock, 0, 0);
    }

    /// @notice Verifies that there is no change in shares if the staker is not delegatedd
    function testFuzz_decreaseOperatorShares_noop() public {
        _registerOperatorWithBaseDetails(defaultOperator);

        cheats.prank(address(allocationManagerMock));
        delegationManager.decreaseOperatorShares(defaultOperator, strategyMock, WAD, WAD);
        assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "shares should not have changed");
    }

    /**
     * @notice Verifies that `DelegationManager.decreaseOperatorShares` properly decreases the delegated `shares` that the operator
     * who the `defaultStaker` is delegated to has in the strategies
     * @dev Checks that there is no change if the staker is not delegated
     * TODO: fuzz magnitude
     */
    function testFuzz_decreaseOperatorShares_slashedOperator(
        IStrategy[] memory strategies,
        uint128 shares,
        bool delegateFromStakerToOperator
    ) public {
        // sanity-filtering on fuzzed input length & staker
        cheats.assume(strategies.length <= 16);
        // TODO: remove, handles rounding on division
        cheats.assume(shares % 2 == 0);

        bool hasBeaconChainStrategy = false;
        for(uint256 i = 0; i < strategies.length; i++) {
            if (strategies[i] == beaconChainETHStrategy) {
                hasBeaconChainStrategy = true;
                break;
            }
        }

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // Set the staker deposits in the strategies
        uint256[] memory sharesToSet = new uint256[](strategies.length);
        for(uint256 i = 0; i < strategies.length; i++) {
            sharesToSet[i] = shares;
        }
        // Okay to set beacon chain shares in SM mock, wont' be called by DM
        strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
        if (hasBeaconChainStrategy) {
            eigenPodManagerMock.setPodOwnerShares(defaultStaker, int256(uint256(shares)));
        }

        // delegate from the `staker` to the operator *if `delegateFromStakerToOperator` is 'true'*
        if (delegateFromStakerToOperator) {
            _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        }

        address delegatedTo = delegationManager.delegatedTo(defaultStaker);

        // for each strategy in `strategies`, increase delegated shares by `shares`
        // noop if the staker is not delegated
        cheats.startPrank(address(strategyManagerMock));
        for (uint256 i = 0; i < strategies.length; ++i) {
            if (delegateFromStakerToOperator) {
                cheats.expectEmit(true, true, true, true, address(delegationManager));
                emit OperatorSharesIncreased(defaultOperator, defaultStaker, strategies[i], shares);
                cheats.expectEmit(true, true, true, true, address(delegationManager));
                emit DepositScalingFactorUpdated(defaultStaker, strategies[i], WAD);
            }
            delegationManager.increaseDelegatedShares(defaultStaker, strategies[i], 0, shares);
            // store delegated shares in a mapping
            delegatedSharesBefore[strategies[i]] = delegationManager.operatorShares(delegatedTo, strategies[i]);
            // also construct an array which we'll use in another loop
            totalSharesForStrategyInArray[address(strategies[i])] += shares;
        }
        cheats.stopPrank();

        // for each strategy in `strategies`, decrease delegated shares by `shares`
        {
            cheats.startPrank(address(allocationManagerMock));
            if (delegateFromStakerToOperator) {
                for (uint256 i = 0; i < strategies.length; ++i) {
                    uint256 currentShares = delegationManager.operatorShares(defaultOperator, strategies[i]);
                    cheats.expectEmit(true, true, true, true, address(delegationManager));
                    emit OperatorSharesDecreased(
                        defaultOperator,
                        address(0),
                        strategies[i],
                        currentShares / 2
                    );
                    delegationManager.decreaseOperatorShares(defaultOperator, strategies[i], WAD, WAD / 2);
                    totalSharesDecreasedForStrategy[strategies[i]] += currentShares / 2;
                }
            }
            cheats.stopPrank();
        }

        // check shares after call to `decreaseOperatorShares`
        uint256[] memory withdrawableShares = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        for (uint256 i = 0; i < strategies.length; ++i) {
            uint256 delegatedSharesAfter = delegationManager.operatorShares(delegatedTo, strategies[i]);

            if (delegateFromStakerToOperator) {
                assertEq(
                    delegatedSharesAfter,
                    delegatedSharesBefore[strategies[i]] - totalSharesDecreasedForStrategy[strategies[i]],
                    "delegated shares did not decrement correctly"
                );
                assertEq(
                    withdrawableShares[i],
                    delegatedSharesAfter,
                    "withdrawable shares for staker not calculated correctly"
                );
            } else {
                assertEq(
                    delegatedSharesAfter,
                    delegatedSharesBefore[strategies[i]],
                    "delegated shares decremented incorrectly"
                );
                assertEq(delegatedSharesBefore[strategies[i]], 0, "nonzero shares delegated to zero address!");
            }
        }
    }
}

contract DelegationManagerUnitTests_Undelegate is DelegationManagerUnitTests {
    // @notice Verifies that undelegating is not possible when the "undelegation paused" switch is flipped
    function test_undelegate_revert_paused(address staker) public filterFuzzedAddressInputs(staker) {
        // set the pausing flag
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_ENTER_WITHDRAWAL_QUEUE);

        cheats.prank(staker);
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.undelegate(staker);
    }

    function testFuzz_undelegate_revert_notDelegated(
        address undelegatedStaker
    ) public filterFuzzedAddressInputs(undelegatedStaker) {
        cheats.assume(undelegatedStaker != defaultOperator);
        assertFalse(delegationManager.isDelegated(undelegatedStaker), "bad test setup");

        cheats.prank(undelegatedStaker);
        cheats.expectRevert(IDelegationManagerErrors.NotActivelyDelegated.selector);
        delegationManager.undelegate(undelegatedStaker);
    }

    // @notice Verifies that an operator cannot undelegate from themself (this should always be forbidden)
    function testFuzz_undelegate_revert_stakerIsOperator(address operator) public filterFuzzedAddressInputs(operator) {
        _registerOperatorWithBaseDetails(operator);

        cheats.prank(operator);
        cheats.expectRevert(IDelegationManagerErrors.OperatorsCannotUndelegate.selector);
        delegationManager.undelegate(operator);
    }

    /**
     * @notice verifies that `DelegationManager.undelegate` reverts if trying to undelegate an operator from themselves
     * @param callFromOperatorOrApprover -- calls from the operator if 'false' and the 'approver' if true
     */
    function testFuzz_undelegate_operatorCannotForceUndelegateThemself(
        address delegationApprover,
        bool callFromOperatorOrApprover
    ) public filterFuzzedAddressInputs(delegationApprover) {
        // register *this contract* as an operator with the default `delegationApprover`
        _registerOperatorWithDelegationApprover(defaultOperator);

        address caller;
        if (callFromOperatorOrApprover) {
            caller = delegationApprover;
        } else {
            caller = defaultOperator;
        }

        // try to call the `undelegate` function and check for reversion
        cheats.prank(caller);
        cheats.expectRevert(IDelegationManagerErrors.OperatorsCannotUndelegate.selector);
        delegationManager.undelegate(defaultOperator);
    }

    function test_undelegate_revert_zeroAddress() public {
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(address(0), defaultOperator);

        cheats.prank(address(0));
        cheats.expectRevert(IPausable.InputAddressZero.selector);
        delegationManager.undelegate(address(0));
    }

    /**
     * @notice Verifies that the `undelegate` function has proper access controls (can only be called by the operator who the `staker` has delegated
     * to or the operator's `delegationApprover`), or the staker themselves
     */
    function testFuzz_undelegate_revert_invalidCaller(
        address invalidCaller
    ) public filterFuzzedAddressInputs(invalidCaller) {
        address staker = address(0x123);
        // filter out addresses that are actually allowed to call the function
        cheats.assume(invalidCaller != staker);
        cheats.assume(invalidCaller != defaultOperator);
        cheats.assume(invalidCaller != defaultApprover);

        _registerOperatorWithDelegationApprover(defaultOperator);
        _delegateToOperatorWhoRequiresSig(staker, defaultOperator);

        cheats.prank(invalidCaller);
        cheats.expectRevert(IDelegationManagerErrors.CallerCannotUndelegate.selector);
        delegationManager.undelegate(staker);
    }

    /**
     * Staker is undelegated from an operator, via a call to `undelegate`, properly originating from the staker's address.
     * Reverts if the staker is themselves an operator (i.e. they are delegated to themselves)
     * Does nothing if the staker is already undelegated
     * Properly undelegates the staker, i.e. the staker becomes “delegated to” the zero address, and `isDelegated(staker)` returns ‘false’
     * Emits a `StakerUndelegated` event
     */
    function testFuzz_undelegate_noDelegateableShares(address staker) public filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != defaultOperator);
        
        // register *this contract* as an operator and delegate from the `staker` to them
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);

        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(staker, delegationManager.delegatedTo(staker));
        cheats.prank(staker);
        bytes32[] memory withdrawalRoots = delegationManager.undelegate(staker);

        assertEq(withdrawalRoots.length, 0, "withdrawalRoot should be an empty array");
        assertEq(
            delegationManager.delegatedTo(staker),
            address(0),
            "undelegated staker should be delegated to zero address"
        );
        assertFalse(delegationManager.isDelegated(staker), "staker not undelegated");
    }

    /**
     * @notice Verifies that the `undelegate` function allows for a force undelegation
     */
    function testFuzz_undelegate_forceUndelegation_noDelegateableShares(
        address staker,
        bytes32 salt,
        bool callFromOperatorOrApprover
    ) public filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != defaultOperator);

        _registerOperatorWithDelegationApprover(defaultOperator);
        _delegateToOperatorWhoRequiresSig(staker, defaultOperator, salt);

        address caller;
        if (callFromOperatorOrApprover) {
            caller = defaultApprover;
        } else {
            caller = defaultOperator;
        }

        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(staker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerForceUndelegated(staker, defaultOperator);
        cheats.prank(caller);
        bytes32[] memory withdrawalRoots = delegationManager.undelegate(staker);

        assertEq(withdrawalRoots.length, 0, "withdrawalRoot should be an empty array");
        assertEq(
            delegationManager.delegatedTo(staker),
            address(0),
            "undelegated staker should be delegated to zero address"
        );
        assertFalse(delegationManager.isDelegated(staker), "staker not undelegated");
    }

    /**
     * @notice Verifies that the `undelegate` function properly queues a withdrawal for all shares of the staker
     */
    function testFuzz_undelegate_nonSlashedOperator(uint128 shares) public {
        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        uint256[] memory sharesToSet = new uint256[](1);
        sharesToSet[0] = shares;
        strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);

        // register *this contract* as an operator and delegate from the `staker` to them
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        
        // Format queued withdrawal
        (
            IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategyMock,
            sharesToWithdraw: shares
        });

        // Undelegate the staker
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(defaultStaker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(defaultOperator, defaultStaker, strategyMock, shares);
        cheats.expectEmit(true, true, true, true, address(delegationManager));  
        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, queuedWithdrawalParams[0].shares);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DepositScalingFactorUpdated(defaultStaker, strategyMock, WAD);
        cheats.prank(defaultStaker);
        delegationManager.undelegate(defaultStaker);

        // Checks - delegation status
        assertEq(
            delegationManager.delegatedTo(defaultStaker),
            address(0),
            "undelegated staker should be delegated to zero address"
        );
        assertFalse(delegationManager.isDelegated(defaultStaker), "staker not undelegated");

        // Checks - operator & staker shares
        assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "operator shares not decreased correctly");
        uint256 stakerWithdrawableShares = delegationManager.getWithdrawableShares(defaultStaker, strategies)[0];
        assertEq(stakerWithdrawableShares, 0, "staker withdrawable shares not calculated correctly");
    }

    /**
     * @notice Verifies that the `undelegate` function properly queues a withdrawal for all shares of the staker
     * @notice The operator should have its shares slashed prior to the staker's deposit
     * TODO: fuzz magnitude
     */
    function testFuzz_undelegate_preSlashedOperator(uint128 shares) public {
        // TODO: remove this assumption & properly handle rounding on division
        cheats.assume(shares % 2 == 0);
        uint64 operatorMagnitude = 5e17;

        // register *this contract* as an operator & set its slashed magnitude
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, operatorMagnitude);
    
        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        {
            uint256[] memory sharesToSet = new uint256[](1);
            sharesToSet[0] = shares;
            strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
        }

        // delegate from the `staker` to them
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        (uint256 depositScalingFactor,,) = delegationManager.stakerScalingFactor(defaultStaker, strategyMock);
        assertTrue(depositScalingFactor > WAD, "bad test setup");
        
        // Format queued withdrawal
        (
            IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategyMock,
            sharesToWithdraw: shares
        });

        // Undelegate the staker
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(defaultStaker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(defaultOperator, defaultStaker, strategyMock, shares);
        cheats.expectEmit(true, true, true, true, address(delegationManager));  
        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, queuedWithdrawalParams[0].shares);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DepositScalingFactorUpdated(defaultStaker, strategyMock, WAD);
        cheats.prank(defaultStaker);
        delegationManager.undelegate(defaultStaker);

        // Checks - delegation status
        assertEq(
            delegationManager.delegatedTo(defaultStaker),
            address(0),
            "undelegated staker should be delegated to zero address"
        );
        assertFalse(delegationManager.isDelegated(defaultStaker), "staker not undelegated");

        // Checks - operator & staker shares
        assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "operator shares not decreased correctly");
        uint256 stakerWithdrawableShares = delegationManager.getWithdrawableShares(defaultStaker, strategies)[0];
        assertEq(stakerWithdrawableShares, 0, "staker withdrawable shares not calculated correctly");
        (uint256 newDepositScalingFactor,,) = delegationManager.stakerScalingFactor(defaultStaker, strategyMock);
        assertEq(newDepositScalingFactor, WAD, "staker scaling factor not reset correctly");
    }

    /**
     * @notice Verifies that the `undelegate` function properly queues a withdrawal for all shares of the staker
     * @notice The operator should have its shares slashed prior to the staker's deposit
     * TODO: fuzz magnitude
     */
    function testFuzz_undelegate_slashedWhileStaked(uint128 shares) public {
        // TODO: remove this assumption & properly handle rounding on division
        cheats.assume(shares % 2 == 0);

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
    
        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        {
            uint256[] memory sharesToSet = new uint256[](1);
            sharesToSet[0] = shares;
            strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
        }
        
        // delegate from the `defaultStaker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Set operator magnitude
        uint64 operatorMagnitude = 5e17;
        uint256 operatorSharesAfterSlash;
        {
            uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
            _setOperatorMagnitude(defaultOperator, strategyMock, operatorMagnitude);
            cheats.prank(address(allocationManagerMock));
            delegationManager.decreaseOperatorShares(defaultOperator, strategyMock, WAD, operatorMagnitude);
            operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategyMock);
            assertEq(operatorSharesAfterSlash, operatorSharesBefore / 2, "operator shares not properly updated");
        }

        (uint256 depositScalingFactor,,) = delegationManager.stakerScalingFactor(defaultStaker, strategyMock);
        assertEq(depositScalingFactor, WAD, "bad test setup");

        // Get withdrawable shares
        uint256 withdrawableSharesBefore = delegationManager.getWithdrawableShares(defaultStaker, strategies)[0];
        
        // Format queued withdrawal
        (
            IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategyMock,
            sharesToWithdraw: withdrawableSharesBefore
        });

        // Undelegate the staker
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(defaultStaker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(defaultOperator, defaultStaker, strategyMock, operatorSharesAfterSlash);
        cheats.expectEmit(true, true, true, true, address(delegationManager));  
        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, queuedWithdrawalParams[0].shares);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DepositScalingFactorUpdated(defaultStaker, strategyMock, WAD);
        cheats.prank(defaultStaker);
        delegationManager.undelegate(defaultStaker);

        // Checks - delegation status
        assertEq(
            delegationManager.delegatedTo(defaultStaker),
            address(0),
            "undelegated staker should be delegated to zero address"
        );
        assertFalse(delegationManager.isDelegated(defaultStaker), "staker not undelegated");

        // Checks - operator & staker shares
        assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "operator shares not decreased correctly");
        uint256 stakerWithdrawableShares = delegationManager.getWithdrawableShares(defaultStaker, strategies)[0];
        assertEq(stakerWithdrawableShares, 0, "staker withdrawable shares not calculated correctly");
        (uint256 newDepositScalingFactor,,) = delegationManager.stakerScalingFactor(defaultStaker, strategyMock);
        assertEq(newDepositScalingFactor, WAD, "staker scaling factor not reset correctly");
    }

    /**
     * @notice Verifies that the `undelegate` function properly undelegates a staker even though their shares
     * were slashed entirely.
     */
    function testFuzz_undelegate_slashedOperator100PercentWhileStaked(uint128 shares) public {
        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
    
        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        {
            uint256[] memory sharesToSet = new uint256[](1);
            sharesToSet[0] = shares;
            strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
        }

        // delegate from the `defaultStaker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Set operator magnitude
        uint64 operatorMagnitude = 0;
        uint256 operatorSharesAfterSlash;
        {
            _setOperatorMagnitude(defaultOperator, strategyMock, operatorMagnitude);
            cheats.prank(address(allocationManagerMock));
            delegationManager.decreaseOperatorShares(defaultOperator, strategyMock, WAD, operatorMagnitude);
            operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategyMock);
            assertEq(operatorSharesAfterSlash, 0, "operator shares not fully slashed");
        }

        (uint256 depositScalingFactor,,) = delegationManager.stakerScalingFactor(defaultStaker, strategyMock);
        assertEq(depositScalingFactor, WAD, "bad test setup");

        // Get withdrawable shares
        uint256 withdrawableSharesBefore = delegationManager.getWithdrawableShares(defaultStaker, strategies)[0];
        assertEq(
            withdrawableSharesBefore,
            0,
            "withdrawable shares should be 0 after being slashed fully"
        );

        // Undelegate the staker
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(defaultStaker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DepositScalingFactorUpdated(defaultStaker, strategyMock, WAD);
        cheats.prank(defaultStaker);
        delegationManager.undelegate(defaultStaker);

        // Checks - delegation status
        assertEq(
            delegationManager.delegatedTo(defaultStaker),
            address(0),
            "undelegated staker should be delegated to zero address"
        );
        assertFalse(delegationManager.isDelegated(defaultStaker), "staker not undelegated");

        // Checks - operator & staker shares
        assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "operator shares not decreased correctly");
        uint256 stakerWithdrawableShares = delegationManager.getWithdrawableShares(defaultStaker, strategies)[0];
        assertEq(stakerWithdrawableShares, 0, "staker withdrawable shares not calculated correctly");
        (uint256 newDepositScalingFactor,,) = delegationManager.stakerScalingFactor(defaultStaker, strategyMock);
        assertEq(newDepositScalingFactor, WAD, "staker scaling factor not reset correctly");
    }

    // TODO: fix old Withdrawals.t.sol test
    //     // @notice This function tests to ensure that a delegator can re-delegate to an operator after undelegating.
    // // @param operator is the operator being delegated to.
    // // @param staker is the staker delegating stake to the operator.
    // function testRedelegateAfterWithdrawal(
    //     address operator,
    //     address depositor,
    //     uint96 ethAmount,
    //     uint96 eigenAmount,
    //     bool withdrawAsShares
    // ) public fuzzedAddress(operator) fuzzedAddress(depositor) {
    //     cheats.assume(depositor != operator);
    //     //this function performs delegation and subsequent withdrawal
    //     testWithdrawalWrapper(operator, depositor, ethAmount, eigenAmount, withdrawAsShares, true);

    //     cheats.prank(depositor);
    //     delegation.undelegate(depositor);

    //     //warps past fraudproof time interval
    //     cheats.warp(block.timestamp + 7 days + 1);
    //     _initiateDelegation(operator, depositor, ethAmount, eigenAmount);
    // }
}

contract DelegationManagerUnitTests_queueWithdrawals is DelegationManagerUnitTests {
    function test_Revert_WhenEnterQueueWithdrawalsPaused() public {
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_ENTER_WITHDRAWAL_QUEUE);
        (IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams, , ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategyMock,
            sharesToWithdraw: 100
        });
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
    }

    function test_Revert_WhenQueueWithdrawalParamsLengthMismatch() public {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        strategyArray[0] = strategyMock;
        uint256[] memory shareAmounts = new uint256[](2);
        shareAmounts[0] = 100;
        shareAmounts[1] = 100;

        IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams = new IDelegationManagerTypes.QueuedWithdrawalParams[](1);
        queuedWithdrawalParams[0] = IDelegationManagerTypes.QueuedWithdrawalParams({
            strategies: strategyArray,
            shares: shareAmounts,
            withdrawer: defaultStaker
        });

        cheats.expectRevert(IDelegationManagerErrors.InputArrayLengthMismatch.selector);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
    }

    function test_Revert_WhenNotStakerWithdrawer(address withdrawer) public {
        cheats.assume(withdrawer != defaultStaker);

        (IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams, , ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: withdrawer,
            strategy: strategyMock,
            sharesToWithdraw: 100
        });
        cheats.expectRevert(IDelegationManagerErrors.WithdrawerNotStaker.selector);
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
    }

    function test_Revert_WhenEmptyStrategiesArray() public {
        IStrategy[] memory strategyArray = new IStrategy[](0);
        uint256[] memory shareAmounts = new uint256[](0);
        address withdrawer = defaultOperator;

        IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams = new IDelegationManagerTypes.QueuedWithdrawalParams[](1);
        queuedWithdrawalParams[0] = IDelegationManagerTypes.QueuedWithdrawalParams({
            strategies: strategyArray,
            shares: shareAmounts,
            withdrawer: withdrawer
        });

        cheats.expectRevert(IDelegationManagerErrors.InputArrayLengthZero.selector);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
    }

    /**
     * @notice Verifies that `DelegationManager.queueWithdrawals` properly queues a withdrawal for the `withdrawer`
     * from the `strategy` for the `sharesAmount`. 
     * - Asserts that staker is delegated to the operator
     * - Asserts that shares for delegatedTo operator are decreased by `sharesAmount`
     * - Asserts that staker cumulativeWithdrawalsQueued nonce is incremented
     * - Checks that event was emitted with correct withdrawalRoot and withdrawal
     */
    function testFuzz_queueWithdrawal_SingleStrat_nonSlashedOperator(
        uint128 depositAmount,
        uint128 withdrawalAmount
    ) public {
        cheats.assume(defaultStaker != defaultOperator);
        cheats.assume(withdrawalAmount > 0 && withdrawalAmount <= depositAmount);
        uint256[] memory sharesAmounts = new uint256[](1);
        sharesAmounts[0] = depositAmount;
        // sharesAmounts is single element so returns single strategy
        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(defaultStaker, sharesAmounts);
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        (
            IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategies[0],
            sharesToWithdraw: withdrawalAmount
        });
        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");
        uint256 nonceBefore = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        uint256 delegatedSharesBefore = delegationManager.operatorShares(defaultOperator, strategies[0]);

        // queueWithdrawals
        cheats.prank(defaultStaker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, queuedWithdrawalParams[0].shares);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        uint256 nonceAfter = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        uint256 delegatedSharesAfter = delegationManager.operatorShares(defaultOperator, strategies[0]);
        assertEq(nonceBefore + 1, nonceAfter, "staker nonce should have incremented");
        assertEq(delegatedSharesBefore - withdrawalAmount, delegatedSharesAfter, "delegated shares not decreased correctly");
    }

    /**
     * @notice Verifies that `DelegationManager.queueWithdrawals` properly queues a withdrawal for the `withdrawer`
     * from the `strategy` for the `sharesAmount`. Operator is slashed prior to the staker's deposit
     * - Asserts that staker is delegated to the operator
     * - Asserts that shares for delegatedTo operator are decreased by `sharesAmount`
     * - Asserts that staker cumulativeWithdrawalsQueued nonce is incremented
     * - Checks that event was emitted with correct withdrawalRoot and withdrawal
     * TODO: fuzz magnitude
     */
    function testFuzz_queueWithdrawal_SingleStrat_preSlashedOperator(
        uint128 depositAmount,
        uint128 withdrawalAmount
    ) public {
        // TODO: remove these assumptions & properly handle rounding on division
        cheats.assume(depositAmount % 2 == 0);
        cheats.assume(withdrawalAmount % 2 == 0);
        cheats.assume(withdrawalAmount > 0 && withdrawalAmount <= depositAmount);

        // Slash the operator
        uint64 operatorMagnitude = 5e17;
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, operatorMagnitude);

        // Deposit for staker & delegate
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        {
            uint256[] memory sharesToSet = new uint256[](1);
            sharesToSet[0] = depositAmount;
            strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
            _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        }

        (
            IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategies[0],
            sharesToWithdraw: withdrawalAmount
        });

        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");
        uint256 nonceBefore = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        uint256 delegatedSharesBefore = delegationManager.operatorShares(defaultOperator, strategies[0]);

        // queueWithdrawals
        cheats.prank(defaultStaker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, queuedWithdrawalParams[0].shares);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        uint256 nonceAfter = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        uint256 delegatedSharesAfter = delegationManager.operatorShares(defaultOperator, strategies[0]);
        assertEq(nonceBefore + 1, nonceAfter, "staker nonce should have incremented");
        assertEq(delegatedSharesBefore - withdrawalAmount, delegatedSharesAfter, "delegated shares not decreased correctly");
    }

    /**
     * @notice Verifies that `DelegationManager.queueWithdrawals` properly queues a withdrawal for the `withdrawer`
     * from the `strategy` for the `sharesAmount`. Operator is slashed while the staker is deposited
     * - Asserts that staker is delegated to the operator
     * - Asserts that shares for delegatedTo operator are decreased by `sharesAmount`
     * - Asserts that staker cumulativeWithdrawalsQueued nonce is incremented
     * - Checks that event was emitted with correct withdrawalRoot and withdrawal
     * TODO: fuzz magnitude
     */
    function testFuzz_queueWithdrawal_SingleStrat_slashedWhileStaked(
        uint128 depositAmount,
        uint128 withdrawalAmount
    ) public {
        // TODO: remove these assumptions & properly handle rounding on division
        cheats.assume(depositAmount % 2 == 0);
        cheats.assume(withdrawalAmount % 2 == 0);
        cheats.assume(withdrawalAmount > 0 && withdrawalAmount <= depositAmount);

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // Deposit for staker & delegate
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        {
            uint256[] memory sharesToSet = new uint256[](1);
            sharesToSet[0] = depositAmount;
            strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
            _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        }

        // Slash the operator
        uint64 operatorMagnitude = 5e17;
        _setOperatorMagnitude(defaultOperator, strategyMock, operatorMagnitude);
        cheats.prank(address(allocationManagerMock));
        delegationManager.decreaseOperatorShares(defaultOperator, strategyMock, WAD, operatorMagnitude);


        (
            IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategies[0],
            sharesToWithdraw: withdrawalAmount
        });

        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");
        uint256 nonceBefore = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        uint256 delegatedSharesBefore = delegationManager.operatorShares(defaultOperator, strategies[0]);
        uint256 withdrawableShares = delegationManager.getWithdrawableShares(defaultStaker, strategies)[0];

        // queueWithdrawals
        if (withdrawalAmount > withdrawableShares) {
            cheats.expectRevert(IDelegationManagerErrors.WithdrawalExceedsMax.selector);
        } else {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, queuedWithdrawalParams[0].shares);
        }
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        if (withdrawalAmount > withdrawableShares) {
            assertEq(delegationManager.cumulativeWithdrawalsQueued(defaultStaker), nonceBefore, "staker nonce should not have incremented");
        } else {
            uint256 nonceAfter = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
            uint256 delegatedSharesAfter = delegationManager.operatorShares(defaultOperator, strategies[0]);
            assertEq(nonceBefore + 1, nonceAfter, "staker nonce should have incremented");
            assertEq(delegatedSharesBefore - withdrawalAmount, delegatedSharesAfter, "delegated shares not decreased correctly");
        }
    }

    /**
     * @notice Verifies that `DelegationManager.queueWithdrawals` reverts when queuing a withdrawal for the `withdrawer`
     * from the `strategy` for the `sharesAmount` since the Operator is slashed 100% while the staker is deposited
     * - Asserts that queuing a withdrawal reverts when the operator is slashed 100%
     * - Asserts that staker withdrawableShares after is 0
     * - Checks that event was emitted with correct withdrawalRoot and withdrawal
     */
    function testFuzz_Revert_queueWithdrawal_SingleStrat_slashed100PercentWhileStaked(
        uint128 depositAmount,
        uint128 withdrawalAmount
    ) public {
        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // Deposit for staker & delegate
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        {
            uint256[] memory sharesToSet = new uint256[](1);
            sharesToSet[0] = depositAmount;
            strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
            _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        }

        // Slash the operator
        uint64 operatorMagnitude = 0;
        _setOperatorMagnitude(defaultOperator, strategyMock, operatorMagnitude);
        cheats.prank(address(allocationManagerMock));
        delegationManager.decreaseOperatorShares(defaultOperator, strategyMock, WAD, 0);

        // Attempt to withdraw for the strategy that was slashed 100% for the operator
        IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams = new IDelegationManagerTypes.QueuedWithdrawalParams[](1);
        {
            uint256[] memory withdrawalAmounts = new uint256[](1);
            withdrawalAmounts[0] = withdrawalAmount;

            queuedWithdrawalParams[0] = IDelegationManagerTypes.QueuedWithdrawalParams({
                strategies: strategies,
                shares: withdrawalAmounts,
                withdrawer: defaultStaker
            });
        }

        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");

        // queueWithdrawals should revert from the 100% slashing
        cheats.prank(defaultStaker);
        cheats.expectRevert(IDelegationManagerErrors.FullySlashed.selector);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        uint256 withdrawableShares = delegationManager.getWithdrawableShares(defaultStaker, strategies)[0];
        assertEq(
            withdrawableShares,
            0,
            "withdrawable shares should be 0 after being slashed fully"
        );
    }

    /**
     * @notice Verifies that `DelegationManager.queueWithdrawals` properly queues a withdrawal for the `withdrawer`
     * with multiple strategies and sharesAmounts. Depending on length sharesAmounts, deploys corresponding number of strategies
     * and deposits sharesAmounts into each strategy for the staker and delegates to operator.
     * For each strategy, withdrawAmount <= depositAmount
     * - Asserts that staker is delegated to the operator
     * - Asserts that shares for delegatedTo operator are decreased by `sharesAmount`
     * - Asserts that staker cumulativeWithdrawalsQueued nonce is incremented
     * - Checks that event was emitted with correct withdrawalRoot and withdrawal
     */
    function testFuzz_queueWithdrawal_MultipleStrats__nonSlashedOperator(
        uint128[] memory depositAmountsUint128
    ) public {
        cheats.assume(depositAmountsUint128.length > 0 && depositAmountsUint128.length <= 32);

        uint256[] memory depositAmounts = new uint256[](depositAmountsUint128.length);
        for (uint256 i = 0; i < depositAmountsUint128.length; i++) {
            depositAmounts[i] = depositAmountsUint128[i];
        }
        uint256[] memory withdrawalAmounts = _fuzzWithdrawalAmounts(depositAmounts);

        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(defaultStaker, depositAmounts);
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        (
            IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawals({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategies: strategies,
            withdrawalAmounts: withdrawalAmounts
        });
        // Before queueWithdrawal state values
        uint256 nonceBefore = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");
        uint256[] memory delegatedSharesBefore = new uint256[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            delegatedSharesBefore[i] = delegationManager.operatorShares(defaultOperator, strategies[i]);
        }

        // queueWithdrawals
        cheats.prank(defaultStaker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, queuedWithdrawalParams[0].shares);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        // Post queueWithdrawal state values
        for (uint256 i = 0; i < strategies.length; i++) {
            assertEq(
                delegatedSharesBefore[i] - withdrawalAmounts[i], // Shares before - withdrawal amount
                delegationManager.operatorShares(defaultOperator, strategies[i]), // Shares after
                "delegated shares not decreased correctly"
            );
        }
        uint256 nonceAfter = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        assertEq(nonceBefore + 1, nonceAfter, "staker nonce should have incremented");
    }
}

contract DelegationManagerUnitTests_completeQueuedWithdrawal is DelegationManagerUnitTests {
    // TODO: add upgrade tests for completing withdrawals queued before upgrade in integration tests

    function test_Revert_WhenExitWithdrawalQueuePaused() public {
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_EXIT_WITHDRAWAL_QUEUE);
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            IERC20[] memory tokens,
            /* bytes32 withdrawalRoot */
        ) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            depositAmount: 100,
            withdrawalAmount: 100
        });
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens,  false);
    }

    function test_Revert_WhenInputArrayLengthMismatch() public {
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            IERC20[] memory tokens,
            /* bytes32 withdrawalRoot */
        ) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            depositAmount: 100,
            withdrawalAmount: 100
        });
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // resize tokens array
        tokens = new IERC20[](0);

        cheats.expectRevert(IDelegationManagerErrors.InputArrayLengthMismatch.selector);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens,  false);
    }

    function test_Revert_WhenWithdrawerNotCaller(address invalidCaller) filterFuzzedAddressInputs(invalidCaller) public {
        cheats.assume(invalidCaller != defaultStaker);

        _registerOperatorWithBaseDetails(defaultOperator);
        (
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            IERC20[] memory tokens,
        ) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            depositAmount: 100,
            withdrawalAmount: 100
        });
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        cheats.expectRevert(IDelegationManagerErrors.WithdrawerNotCaller.selector);
        cheats.prank(invalidCaller);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens,  false);
    }

    function test_Revert_WhenInvalidWithdrawalRoot() public {
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            IERC20[] memory tokens,
            bytes32 withdrawalRoot
        ) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            depositAmount: 100,
            withdrawalAmount: 100
        });
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");
        cheats.roll(withdrawal.startBlock + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS());
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens,  true);
        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be completed and marked false now");

        cheats.expectRevert(IDelegationManagerErrors.WithdrawalNotQueued.selector);
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens,  false);
    }

    /**
     * @notice should revert if minWithdrawalDelayBlocks has not passed, and if
     * delegationManager.getCompletableTimestamp returns a value greater than minWithdrawalDelayBlocks
     * then it should revert if the validBlockNumber has not passed either.
     */
    function test_Revert_WhenWithdrawalDelayNotPassed(
        uint256[] memory depositAmounts,
        bool receiveAsTokens
    ) public {
        cheats.assume(depositAmounts.length > 0 && depositAmounts.length <= 32);
        uint256[] memory withdrawalAmounts = _fuzzWithdrawalAmounts(depositAmounts);
        
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            IERC20[] memory tokens,
            /* bytes32 withdrawalRoot */
        ) = _setUpCompleteQueuedWithdrawal({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            depositAmounts: depositAmounts,
            withdrawalAmounts: withdrawalAmounts
        });

        // prank as withdrawer address
        cheats.roll(withdrawal.startBlock + minWithdrawalDelayBlocks - 1);
        cheats.expectRevert(IDelegationManagerErrors.WithdrawalDelayNotElapsed.selector);
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens,  receiveAsTokens);
    }

    /**
     * @notice Verifies that `DelegationManager.completeQueuedWithdrawal` properly completes a queued withdrawal for the `withdrawer`
     * for a single strategy. Withdraws as tokens so there are no operator shares increase.
     * - Asserts that the withdrawalRoot is True before `completeQueuedWithdrawal` and False after
     * - Asserts operatorShares is unchanged after `completeQueuedWithdrawal`
     * - Checks that event `WithdrawalCompleted` is emitted with withdrawalRoot
     */
    function test_completeQueuedWithdrawal_SingleStratWithdrawAsTokens(
        address staker,
        uint128 depositAmount,
        uint128 withdrawalAmount
    ) public filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != defaultOperator);
        cheats.assume(withdrawalAmount > 0 && withdrawalAmount <= depositAmount);
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            IERC20[] memory tokens,
            bytes32 withdrawalRoot
        ) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: staker,
            withdrawer: staker,
            depositAmount: depositAmount,
            withdrawalAmount: withdrawalAmount
        });
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, withdrawal.strategies[0]);
        assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");

        // completeQueuedWithdrawal
        cheats.roll(withdrawal.startBlock + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS());
        cheats.prank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit SlashingWithdrawalCompleted(withdrawalRoot);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens,  true);

        uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, withdrawal.strategies[0]);
        assertEq(operatorSharesAfter, operatorSharesBefore, "operator shares should be unchanged");
        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be completed and marked false now");
    }

    /**
     * @notice Verifies that `DelegationManager.completeQueuedWithdrawal` properly completes a queued withdrawal for the `withdrawer`
     * for a single strategy. Withdraws as tokens so there are no operator shares increase.
     * - Asserts that the withdrawalRoot is True before `completeQueuedWithdrawal` and False after
     * - Asserts operatorShares is decreased after the operator is slashed
     * - Checks that event `WithdrawalCompleted` is emitted with withdrawalRoot
     * - Asserts that the shares the staker completed withdrawal for are less than what is expected since its operator is slashed
     */
    function test_completeQueuedWithdrawal_SingleStratWithdrawAsTokens_slashOperatorDuringQueue(
        uint128 depositAmount,
        uint128 withdrawalAmount
    ) public {
        // TODO: remove these assumptions & properly handle rounding on division
        cheats.assume(depositAmount % 2 == 0);
        cheats.assume(withdrawalAmount % 2 == 0);
        cheats.assume(withdrawalAmount > 0 && withdrawalAmount <= depositAmount);


        // Deposit Staker
        uint256[] memory depositAmounts = new uint256[](1);
        depositAmounts[0] = depositAmount;
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        strategyManagerMock.setDeposits(defaultStaker, strategies, depositAmounts);

        // Register operator and delegate to it
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        uint256 operatorSharesBeforeQueue = delegationManager.operatorShares(defaultOperator, strategyMock);

        // Queue withdrawal
        (
            IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategyMock,
            sharesToWithdraw: withdrawalAmount
        });
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
        assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");
        uint256 operatorSharesAfterQueue = delegationManager.operatorShares(defaultOperator, strategyMock);

        assertEq(operatorSharesAfterQueue, operatorSharesBeforeQueue - withdrawalAmount, "operator shares should be decreased after queue");

        // Slash operator while staker has queued withdrawal
        uint64 operatorMagnitude = 5e17;
        _setOperatorMagnitude(defaultOperator, withdrawal.strategies[0], operatorMagnitude);
        cheats.prank(address(allocationManagerMock));
        delegationManager.decreaseOperatorShares(defaultOperator, withdrawal.strategies[0], WAD, operatorMagnitude);
        uint256 operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategyMock);
        assertEq(operatorSharesAfterSlash, operatorSharesAfterQueue / 2, "operator shares should be decreased after slash");

        // Complete queue withdrawal
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = IERC20(strategies[0].underlyingToken());
        cheats.roll(withdrawal.startBlock + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS());
        cheats.prank(defaultStaker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit SlashingWithdrawalCompleted(withdrawalRoot);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens,  true);

        // Checks: operator shares
        uint256 operatorSharesAfterWithdrawalComplete = delegationManager.operatorShares(defaultOperator, withdrawal.strategies[0]);
        assertEq(operatorSharesAfterWithdrawalComplete, operatorSharesAfterSlash, "operator shares should be unchanged from slash to withdrawal completion");
        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be completed and marked false now");

        // Checks: staker shares: 
        uint256 stakerSharesWithdrawn = strategyManagerMock.strategySharesWithdrawn(defaultStaker, strategyMock);
        assertEq(stakerSharesWithdrawn, withdrawalAmount / 2, "staker shares withdrawn should be half of expected since operator is slashed by half");
    }

    /**
     * @notice Verifies that `DelegationManager.completeQueuedWithdrawal` properly completes a queued withdrawal for the `withdrawer`
     * for the BeaconChainStrategy. Withdraws as tokens so there are no operator shares increase.
     * - Asserts that the withdrawalRoot is True before `completeQueuedWithdrawal` and False after
     * - Asserts operatorShares is decreased after staker is slashed
     * - Checks that event `WithdrawalCompleted` is emitted with withdrawalRoot
     * - Asserts that the shares the staker completed withdrawal for are less than what is expected since the staker is slashed during queue
     */
    // TODO: fuzz the beacon chain magnitude
    function test_completeQueuedWithdrawal_BeaconStratWithdrawAsTokens_slashStakerDuringQueue(
        uint128 depositAmount,
        uint128 withdrawalAmount
    ) public {
        // TODO: remove these assumptions & properly handle rounding on division
        cheats.assume(depositAmount % 2 == 0);
        cheats.assume(withdrawalAmount % 2 == 0);
        cheats.assume(withdrawalAmount > 0 && withdrawalAmount <= depositAmount);

        // Deposit Staker
        eigenPodManagerMock.setPodOwnerShares(defaultStaker, int256(uint256(depositAmount)));

        // Register operator and delegate to it
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        uint256 operatorSharesBeforeQueue = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);

        // Queue withdrawal
        (
            IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: beaconChainETHStrategy,
            sharesToWithdraw: withdrawalAmount
        });
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
        assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");
        uint256 operatorSharesAfterQueue = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
        assertEq(operatorSharesAfterQueue, operatorSharesBeforeQueue - withdrawalAmount, "operator shares should be decreased after queue");

        // Slash the staker for beacon chain shares while it has queued a withdrawal
        uint256 beaconSharesBeforeSlash = uint256(eigenPodManagerMock.podOwnerShares(defaultStaker));
        uint64 stakerBeaconChainScalingFactor = 5e17;
        cheats.prank(address(eigenPodManagerMock));
        delegationManager.decreaseBeaconChainScalingFactor(defaultStaker, beaconSharesBeforeSlash, stakerBeaconChainScalingFactor);
        uint256 operatorSharesAfterBeaconSlash = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
        assertEq(operatorSharesAfterBeaconSlash, operatorSharesAfterQueue / 2, "operator shares should be decreased after beaconChain slash");

        // Complete queue withdrawal
        IERC20[] memory tokens = new IERC20[](1);
        cheats.roll(withdrawal.startBlock + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS());
        cheats.prank(defaultStaker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit SlashingWithdrawalCompleted(withdrawalRoot);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens,  true);

        // Checks: operator shares
        uint256 operatorSharesAfterWithdrawalComplete = delegationManager.operatorShares(defaultOperator, withdrawal.strategies[0]);
        assertEq(operatorSharesAfterWithdrawalComplete, operatorSharesAfterBeaconSlash, "operator shares should be unchanged from slash to withdrawal completion");
        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be completed and marked false now");

        // Checks: staker shares
        uint256 stakerBeaconSharesWithdrawn = eigenPodManagerMock.podOwnerSharesWithdrawn(defaultStaker);
        assertEq(stakerBeaconSharesWithdrawn, withdrawalAmount / 2, "staker shares withdrawn should be half of expected it is slashed by half");
    }

    /**
     * @notice Verifies that `DelegationManager.completeQueuedWithdrawal` properly completes a queued withdrawal for the `withdrawer`
     * for the BeaconChainStrategy. Withdraws as tokens so there are no operator shares increase.
     * - Asserts that the withdrawalRoot is True before `completeQueuedWithdrawal` and False after
     * - Asserts operatorShares is decreased after staker is slashed and after the operator is slashed
     * - Checks that event `WithdrawalCompleted` is emitted with withdrawalRoot
     * - Asserts that the shares the staker completed withdrawal for are less than what is expected since both the staker and its operator are slashed during queue
     */
    // TODO: fuzz the beacon chain magnitude & operator magnitude
    function test_completeQueuedWithdrawal_BeaconStratWithdrawAsTokens_slashStakerAndOperator(
        uint128 depositAmount,
        uint128 withdrawalAmount
    ) public {
        // TODO: remove these assumptions & properly handle rounding on division
        cheats.assume(depositAmount % 2 == 0);
        cheats.assume(withdrawalAmount % 2 == 0);
        cheats.assume(withdrawalAmount > 0 && withdrawalAmount <= depositAmount);

        // Deposit Staker
        eigenPodManagerMock.setPodOwnerShares(defaultStaker, int256(uint256(depositAmount)));

        // Register operator and delegate to it
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        uint256 operatorSharesBeforeQueue = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);

        // Queue withdrawal
        (
            IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: beaconChainETHStrategy,
            sharesToWithdraw: withdrawalAmount
        });

        uint256 operatorSharesAfterAVSSlash;
        {
            cheats.prank(defaultStaker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");
            uint256 operatorSharesAfterQueue = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
            assertEq(operatorSharesAfterQueue, operatorSharesBeforeQueue - withdrawalAmount, "operator shares should be decreased after queue");

            // Slash the staker for beacon chain shares while it has queued a withdrawal
            uint256 beaconSharesBeforeSlash = uint256(eigenPodManagerMock.podOwnerShares(defaultStaker));
            uint64 stakerBeaconChainScalingFactor = 5e17;
            cheats.prank(address(eigenPodManagerMock));
            delegationManager.decreaseBeaconChainScalingFactor(defaultStaker, beaconSharesBeforeSlash, stakerBeaconChainScalingFactor);
            uint256 operatorSharesAfterBeaconSlash = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
            assertEq(operatorSharesAfterBeaconSlash, operatorSharesAfterQueue / 2, "operator shares should be decreased after beaconChain slash");

            // Slash the operator for beacon chain shares
            uint64 operatorMagnitude = 5e17;
            _setOperatorMagnitude(defaultOperator, withdrawal.strategies[0], operatorMagnitude);
            cheats.prank(address(allocationManagerMock));
            delegationManager.decreaseOperatorShares(defaultOperator, withdrawal.strategies[0], WAD, operatorMagnitude);
            operatorSharesAfterAVSSlash = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
            assertEq(operatorSharesAfterAVSSlash, operatorSharesAfterBeaconSlash / 2, "operator shares should be decreased after AVS slash");
        }
        operatorSharesAfterAVSSlash = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);

        // Complete queue withdrawal
        IERC20[] memory tokens = new IERC20[](1);
        cheats.roll(withdrawal.startBlock + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS());
        cheats.prank(defaultStaker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit SlashingWithdrawalCompleted(withdrawalRoot);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens,  true);

        // Checks: operator shares
        uint256 operatorSharesAfterWithdrawalComplete = delegationManager.operatorShares(defaultOperator, withdrawal.strategies[0]);
        assertEq(operatorSharesAfterWithdrawalComplete, operatorSharesAfterAVSSlash, "operator shares should be unchanged from slash to withdrawal completion");
        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be completed and marked false now");

        // Checks: staker shares
        uint256 stakerBeaconSharesWithdrawn = eigenPodManagerMock.podOwnerSharesWithdrawn(defaultStaker);
        assertEq(stakerBeaconSharesWithdrawn, withdrawalAmount / 4, "staker shares withdrawn should be 1/4th of expected it is slashed by half twice");
    }


    /**
     * @notice Verifies that `DelegationManager.completeQueuedWithdrawal` properly completes a queued withdrawal for the `withdrawer`
     * for a single strategy. Withdraws as shares so if the withdrawer is delegated, operator shares increase. In the test case, this only
     * happens if staker and withdrawer are fuzzed the same address (i.e. staker == withdrawer)
     * - Asserts that the withdrawalRoot is True before `completeQueuedWithdrawal` and False after
     * - Asserts if staker == withdrawer, operatorShares increase, otherwise operatorShares are unchanged
     * - Checks that event `WithdrawalCompleted` is emitted with withdrawalRoot
     */
    function test_completeQueuedWithdrawal_SingleStratWithdrawAsShares_nonSlashedOperator(
        address staker,
        uint128 depositAmount,
        uint128 withdrawalAmount
    ) public filterFuzzedAddressInputs(staker) {
        // TODO: remove these assumptions & properly handle rounding on division
        cheats.assume(staker != defaultOperator);
        cheats.assume(withdrawalAmount > 0 && withdrawalAmount <= depositAmount);
        _registerOperatorWithBaseDetails(defaultOperator);

        (
            IDelegationManagerTypes.Withdrawal memory withdrawal,
            IERC20[] memory tokens,
            bytes32 withdrawalRoot
        ) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: staker,
            withdrawer: staker,
            depositAmount: depositAmount,
            withdrawalAmount: withdrawalAmount
        });
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, withdrawal.strategies[0]);
        assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");

        // Set delegationManager on strategyManagerMock so it can call back into delegationManager
        strategyManagerMock.setDelegationManager(delegationManager);

        // completeQueuedWithdrawal
        cheats.roll(withdrawal.startBlock + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS());
        cheats.prank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit SlashingWithdrawalCompleted(withdrawalRoot);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens,  false);

        uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, withdrawal.strategies[0]);
        // Since staker is delegated, operatorShares get incremented
        assertEq(operatorSharesAfter, operatorSharesBefore + withdrawalAmount, "operator shares not increased correctly");
        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be completed and marked false now");
    }

    // TODO: add slashing cases for withdrawing as shares (can also be in integration tests)
}


/**
 * @notice TODO Lifecycle tests - These tests combine multiple functionalities of the DelegationManager
   1. Old SigP test - registerAsOperator, separate staker delegate to operator, as operator undelegate (reverts),
    checks that staker is still delegated and operator still registered, staker undelegates, checks staker not delegated and operator
    is still registered
   2. RegisterOperator, Deposit, Delegate, Queue, Complete
   3. RegisterOperator, Mock Slash(set maxMagnitudes), Deposit/Delegate, Queue, Complete
   4. RegisterOperator, Deposit/Delegate, Mock Slash(set maxMagnitudes), Queue, Complete
   5. RegisterOperator, Mock Slash(set maxMagnitudes), Deposit/Delegate, Queue, Mock Slash(set maxMagnitudes), Complete
   7. RegisterOperator, Deposit/Delegate, Mock Slash 100% (set maxMagnitudes), Undelegate, Complete non 100% slashed strategies
   8. RegisterOperator, Deposit/Delegate, Undelegate, Re delegate to another operator, Mock Slash 100% (set maxMagnitudes), Complete as shares
    (withdrawals should have been slashed even though delegated to a new operator)
 */
