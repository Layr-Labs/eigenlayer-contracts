// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

import "src/contracts/core/DelegationManager.sol";
import "src/contracts/strategies/StrategyBase.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";
import "src/contracts/libraries/SlashingLib.sol";
import "src/test/utils/SingleItemArrayLib.sol";

// TODO: add upgrade tests for completing withdrawals queued before upgrade in integration tests
// TODO: add slashing cases for withdrawing as shares (can also be in integration tests)

/**
 * @notice Unit testing of the DelegationManager contract. Withdrawals are tightly coupled
 * with EigenPodManager and StrategyManager and are part of integration tests.
 * Contracts tested: DelegationManager
 * Contracts not mocked: StrategyBase, PauserRegistry
 */
contract DelegationManagerUnitTests is EigenLayerUnitTestSetup, IDelegationManagerEvents, IDelegationManagerErrors {
    using SlashingLib for *; 
    using SingleItemArrayLib for *;

    // Contract under test
    DelegationManager delegationManager;
    DelegationManager delegationManagerImplementation;

    // Helper to use in storage
    DepositScalingFactor dsf;

    // Mocks
    StrategyBase strategyImplementation;
    StrategyBase strategyMock;

    IERC20 tokenMock;
    uint256 tokenMockInitialSupply = 10e50;

    uint32 constant MIN_WITHDRAWAL_DELAY_BLOCKS = 126_000; // 17.5 days in blocks
    uint256 MAX_STRATEGY_SHARES = 1e38 - 1;

    // Max shares in a strategy, see StrategyBase.sol
    uint256 MAX_STRATEGY_SHARES = 1e38 - 1;
    uint256 MAX_ETH_SUPPLY = 120_400_000 ether;

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

        delegationManager = DelegationManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        // Redeploy StrategyManagerMock with DM
        strategyManagerMock = StrategyManagerMock(payable(address(new StrategyManagerMock(delegationManager))));

        // Deploy DelegationManager implmentation and upgrade proxy
        delegationManagerImplementation = new DelegationManager(
            IAVSDirectory(address(avsDirectoryMock)), 
            IStrategyManager(address(strategyManagerMock)), 
            IEigenPodManager(address(eigenPodManagerMock)), 
            IAllocationManager(address(allocationManagerMock)), 
            pauserRegistry,
            MIN_WITHDRAWAL_DELAY_BLOCKS
        );

        eigenLayerProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(delegationManager))),
            address(delegationManagerImplementation),
            abi.encodeWithSelector(
                DelegationManager.initialize.selector,
                address(this),
                0 // 0 is initial paused status
            )
        );

        // Deploy mock token and strategy
        tokenMock = new ERC20PresetFixedSupply("Mock Token", "MOCK", tokenMockInitialSupply, address(this));
        strategyImplementation = new StrategyBase(IStrategyManager(address(strategyManagerMock)), pauserRegistry);
        strategyMock = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, tokenMock)
                )
            )
        );

        // Roll blocks forward so that block.number - MIN_WITHDRAWAL_DELAY_BLOCKS doesn't revert
        // in _getSlashableSharesInQueue
        cheats.roll(MIN_WITHDRAWAL_DELAY_BLOCKS);

        // Exclude delegation manager from fuzzed tests
        isExcludedFuzzAddress[address(delegationManager)] = true;
        isExcludedFuzzAddress[address(strategyManagerMock)] = true;
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
                tokenMockInitialSupply,
                address(this)
            );
            strategies[i] = StrategyBase(
                address(
                    new TransparentUpgradeableProxy(
                        address(strategyImplementation),
                        address(eigenLayerProxyAdmin),
                        abi.encodeWithSelector(StrategyBase.initialize.selector, token)
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

    function _registerOperatorWithBaseDetails(address operator) internal {
        OperatorDetails memory operatorDetails = OperatorDetails({
            __deprecated_earningsReceiver: operator,
            delegationApprover: address(0),
            __deprecated_stakerOptOutWindowBlocks: 0
        });
        _registerOperator(operator, operatorDetails, emptyStringForMetadataURI);
    }

    function _registerOperatorWithDelegationApprover(address operator) internal {
        OperatorDetails memory operatorDetails = OperatorDetails({
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

        OperatorDetails memory operatorDetails = OperatorDetails({
            __deprecated_earningsReceiver: operator,
            delegationApprover: address(wallet),
            __deprecated_stakerOptOutWindowBlocks: 0
        });
        _registerOperator(operator, operatorDetails, emptyStringForMetadataURI);

        return wallet;
    }

    function _registerOperator(
        address operator,
        OperatorDetails memory operatorDetails,
        string memory metadataURI
    ) internal filterFuzzedAddressInputs(operator) {
        cheats.prank(operator);
        delegationManager.registerAsOperator(operatorDetails, 0, metadataURI);
    }

    /**
     * @notice Using this helper function to fuzz withdrawalAmounts since fuzzing two dynamic sized arrays of equal lengths
     * reject too many inputs. 
     */
    function _fuzzDepositWithdrawalAmounts(uint256[] memory fuzzAmounts) internal view returns (uint256[] memory, uint256[] memory) {
        cheats.assume(fuzzAmounts.length > 0);
        uint256[] memory withdrawalAmounts = new uint256[](fuzzAmounts.length);
        // We want to bound deposits amounts as well
        uint256[] memory depositAmounts = new uint256[](fuzzAmounts.length);
        for (uint256 i = 0; i < fuzzAmounts.length; i++) {
            depositAmounts[i] = bound(
                uint256(keccak256(abi.encodePacked(fuzzAmounts[i]))),
                1,
                1e38 - 1
            );
            // generate withdrawal amount within range s.t withdrawAmount <= depositAmount
            withdrawalAmounts[i] = bound(
                uint256(keccak256(abi.encodePacked(depositAmounts[i]))),
                0,
                depositAmounts[i]
            );
        }
        return (depositAmounts, withdrawalAmounts);
    }

    function _setUpQueueWithdrawalsSingleStrat(
        address staker,
        address withdrawer,
        IStrategy strategy,
        uint256 depositSharesToWithdraw
    ) internal view returns (
        QueuedWithdrawalParams[] memory,
        Withdrawal memory,
        bytes32
    ) {
        IStrategy[] memory strategyArray = strategy.toArray();
        QueuedWithdrawalParams[] memory queuedWithdrawalParams = new QueuedWithdrawalParams[](1);
        {
            uint256[] memory withdrawalAmounts = depositSharesToWithdraw.toArrayU256();
            queuedWithdrawalParams[0] = QueuedWithdrawalParams({
                strategies: strategyArray,
                depositShares: withdrawalAmounts,
                withdrawer: withdrawer
            });
        }

        // Get scaled shares to withdraw
        uint256[] memory scaledSharesArray = _getScaledShares(staker, strategy, depositSharesToWithdraw).toArrayU256();

        Withdrawal memory withdrawal = Withdrawal({
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
        uint256[] memory depositWithdrawalAmounts
    ) internal view returns (
        QueuedWithdrawalParams[] memory,
        Withdrawal memory,
        bytes32
    ) {
        QueuedWithdrawalParams[] memory queuedWithdrawalParams = new QueuedWithdrawalParams[](1);
        {
            queuedWithdrawalParams[0] = QueuedWithdrawalParams({
                strategies: strategies,
                depositShares: depositWithdrawalAmounts,
                withdrawer: withdrawer
            });
        }

        // Get scaled shares to withdraw
        uint256[] memory scaledSharesArray = new uint256[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            scaledSharesArray[i] = _getScaledShares(staker, strategies[i], depositWithdrawalAmounts[i]);
        }
        
        Withdrawal memory withdrawal = Withdrawal({
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

    function _getScaledShares(address staker, IStrategy strategy, uint256 depositSharesToWithdraw) internal view returns (uint256) {
        // Setup vars
        address operator = delegationManager.delegatedTo(staker);
        IStrategy[] memory strategyArray = new IStrategy[](1);
        strategyArray[0] = strategy;

        // Calculate the amount of slashing to apply
        uint64 maxMagnitude = allocationManagerMock.getMaxMagnitudes(operator, strategyArray)[0];
        uint256 slashingFactor = _getSlashingFactor(staker, strategy, maxMagnitude);

        DepositScalingFactor memory _dsf = DepositScalingFactor(delegationManager.depositScalingFactor(staker, strategy));
        uint256 sharesToWithdraw = _dsf.calcWithdrawable(depositSharesToWithdraw, slashingFactor);

        uint256 scaledShares = SlashingLib.scaleForQueueWithdrawal({
            sharesToWithdraw: sharesToWithdraw,
            slashingFactor: slashingFactor
        });

        return scaledShares;
    }

    function _getSlashingFactor(
        address staker,
        IStrategy strategy,
        uint64 operatorMaxMagnitude
    ) internal view returns (uint256) {
        if (strategy == beaconChainETHStrategy) {
            uint64 beaconChainSlashingFactor = delegationManager.getBeaconChainSlashingFactor(staker);
            return operatorMaxMagnitude.mulWad(beaconChainSlashingFactor);
        }

        return operatorMaxMagnitude;
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
    ) internal returns (Withdrawal memory, IERC20[] memory, bytes32) {
        uint256[] memory depositAmounts = depositAmount.toArrayU256();
        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(staker, depositAmounts);
        (
            QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: staker,
            withdrawer: withdrawer,
            strategy: strategies[0],
            depositSharesToWithdraw: withdrawalAmount
        });

        cheats.prank(staker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
        // Set the current deposits to be the depositAmount - withdrawalAmount
        uint256[] memory currentAmounts = uint256(depositAmount - withdrawalAmount).toArrayU256();
        strategyManagerMock.setDeposits(staker, strategies, currentAmounts);

        IERC20[] memory tokens = new IERC20[](strategies.length);
        for (uint i = 0; i < tokens.length; i++) {
            tokens[i] = strategies[i].underlyingToken();
        }

        return (withdrawal, tokens, withdrawalRoot);
    }

    /**
     * Deploy and deposit staker into a single strategy, then set up multiple queued withdrawals for the staker
     * Assumptions: 
     * - operator is already a registered operator.
     * - total deposit amount = depositAmount * numWithdrawals
     * - this will fully withdraw from the single strategy
     */
    function _setUpCompleteQueuedWithdrawalsSingleStrat(
        address staker,
        address withdrawer,
        uint256 depositAmount,
        uint256 numWithdrawals
    ) internal returns (
        IDelegationManagerTypes.Withdrawal[] memory withdrawals, 
        IERC20[][] memory tokens, 
        bytes32[] memory withdrawalRoots
    ) {
        uint256[] memory depositAmounts = new uint256[](1);
        depositAmounts[0] = depositAmount * numWithdrawals;
        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(staker, depositAmounts);

        withdrawals = new Withdrawal[](numWithdrawals);
        tokens = new IERC20[][](numWithdrawals);
        withdrawalRoots = new bytes32[](numWithdrawals);

        for (uint i = 0; i < numWithdrawals; i++) {
            (
                IDelegationManagerTypes.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
                IDelegationManagerTypes.Withdrawal memory withdrawal,
                bytes32 withdrawalRoot
            ) = _setUpQueueWithdrawalsSingleStrat({
                staker: staker,
                withdrawer: withdrawer,
                strategy: strategies[0],
                depositSharesToWithdraw: depositAmount
            });

            cheats.prank(staker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);

            withdrawals[i] = withdrawal;
            tokens[i] = new IERC20[](1);
            tokens[i][0] = strategies[0].underlyingToken();
            withdrawalRoots[i] = withdrawalRoot;
        }

        {
            // Set the current deposits to be 0
            uint256[] memory currentAmounts = new uint256[](1);
            currentAmounts[0] = 0;
            strategyManagerMock.setDeposits(staker, strategies, currentAmounts);
        }
        

        return (withdrawals, tokens, withdrawalRoots);
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
    ) internal returns (Withdrawal memory, IERC20[] memory, bytes32) {
        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(staker, depositAmounts);

        IERC20[] memory tokens = new IERC20[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            tokens[i] = strategies[i].underlyingToken();
        }

        (
            QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawals({
            staker: staker,
            withdrawer: withdrawer,
            strategies: strategies,
            depositWithdrawalAmounts: withdrawalAmounts
        });

        cheats.prank(staker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        return (withdrawal, tokens, withdrawalRoot);
    }

    function _setOperatorMagnitude(
        address operator,
        IStrategy strategy,
        uint64 magnitude
    ) internal returns (uint256 sharesToDecrease) {
        allocationManagerMock.setMaxMagnitude(operator, strategy, magnitude);
    }

    /**
     * Given shares, current operator max magnitude, and wad to slash, calculate the slashed shares and new magnitude
     * This function simulates the slashOperator operations in the AllocationManager
     */
    function _calculateSlashedSharesAndMagnitude(
        uint256 shares,
        uint64 operatorMagnitude,
        uint256 wadToSlash
    ) internal pure returns (uint256, uint64, uint256) {
        uint64 slashedMagnitude = uint64(operatorMagnitude.mulWadRoundUp(wadToSlash));
        uint256 sharesWadSlashed = uint256(slashedMagnitude).divWad(operatorMagnitude);

        uint256 sharesSlashed = shares.mulWad(sharesWadSlashed);
        uint64 newMagnitude = uint64(operatorMagnitude - slashedMagnitude);
        return (sharesSlashed, newMagnitude, sharesWadSlashed);
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
        delegationManager.initialize(address(this), 0);
    }
}

contract DelegationManagerUnitTests_RegisterModifyOperator is DelegationManagerUnitTests {
    function test_registerAsOperator_revert_paused() public {
        // set the pausing flag
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_NEW_DELEGATION);

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.registerAsOperator(
            OperatorDetails({
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
        OperatorDetails memory operatorDetails
    ) public {
        // Register once
        cheats.startPrank(defaultOperator);
        delegationManager.registerAsOperator(operatorDetails, 0, emptyStringForMetadataURI);

        // Expect revert when register again
        cheats.expectRevert(ActivelyDelegated.selector);
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
        OperatorDetails memory operatorDetails,
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

    /// @notice Register two separate operators shouldn't revert
    function testFuzz_registerAsOperator_TwoSeparateOperatorsRegister(
        address operator1,
        IDelegationManagerTypes.OperatorDetails memory operatorDetails1,
        address operator2,
        IDelegationManagerTypes.OperatorDetails memory operatorDetails2
    ) public {
        cheats.assume(operator1 != operator2);
        _registerOperatorWithBaseDetails(operator1);
        _registerOperatorWithBaseDetails(operator2);
        assertTrue(
            delegationManager.isOperator(operator1),
            "operator1 not registered"
        );
        assertTrue(
            delegationManager.isOperator(operator2),
            "operator2 not registered"
        );
    }


    // @notice Verifies that a staker who is actively delegated to an operator cannot register as an operator (without first undelegating, at least)
    function testFuzz_Revert_registerAsOperator_cannotRegisterWhileDelegated(
        address staker,
        OperatorDetails memory operatorDetails
    ) public filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != defaultOperator);

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, emptySalt);

        // expect revert if attempt to register as operator
        cheats.expectRevert(ActivelyDelegated.selector);
        delegationManager.registerAsOperator(operatorDetails, 0, emptyStringForMetadataURI);

        cheats.stopPrank();
    }
    
    /// @notice Add test for registerAsOperator where the operator has existing deposits in strategies
    /// Assert:
    ///     depositShares == operatorShares == withdrawableShares
    ///     check operatorDetails hash encode matches the operatorDetails hash stored (call view function)
    function testFuzz_registerAsOperator_withDeposits(
        uint128 shares
    ) public {
        // Set staker shares in StrategyManager
        IStrategy[] memory strategiesToReturn = new IStrategy[](1);
        strategiesToReturn[0] = strategyMock;
        uint256[] memory sharesToReturn = new uint256[](1);
        sharesToReturn[0] = shares;
        strategyManagerMock.setDeposits(defaultOperator, strategiesToReturn, sharesToReturn);
        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);

        // register operator, their own staker depositShares should increase their operatorShares
        _registerOperatorWithBaseDetails(defaultOperator);
        uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);

        // check depositShares == operatorShares == withdrawableShares
        assertEq(operatorSharesAfter, shares, "operator shares not set correctly");
        (uint256[] memory withdrawableShares, ) = delegationManager.getWithdrawableShares(defaultOperator, strategiesToReturn);
        assertEq(
            withdrawableShares[0],
            shares,
            "withdrawable shares not set correctly"
        );
        assertEq(
            strategyManagerMock.stakerDepositShares(defaultOperator, strategyMock),
            shares,
            "staker deposit shares not set correctly"
        );
    }

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
        OperatorDetails memory initialOperatorDetails,
        OperatorDetails memory modifiedOperatorDetails
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
        cheats.expectRevert(OperatorNotRegistered.selector);
        delegationManager.updateOperatorMetadataURI(emptyStringForMetadataURI);
    }

    /**
     * @notice Verifies that a staker cannot call cannot modify their `OperatorDetails` without first registering as an operator
     * @dev This is an important check to ensure that our definition of 'operator' remains consistent, in particular for preserving the
     * invariant that 'operators' are always delegated to themselves
     */
    function testFuzz_Revert_updateOperatorMetadataUri_notOperator(
        OperatorDetails memory operatorDetails
    ) public {
        cheats.expectRevert(OperatorNotRegistered.selector);
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
    using SingleItemArrayLib for *;
    using SlashingLib for *;

    function test_Revert_WhenPaused() public {
        cheats.prank(defaultOperator);
        delegationManager.registerAsOperator(
            OperatorDetails({
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
        cheats.expectRevert(ActivelyDelegated.selector);
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
        cheats.expectRevert(OperatorNotRegistered.selector);
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
        IStrategy[] memory strategiesToReturn = strategyMock.toArray();
        uint256[] memory sharesToReturn = shares.toArrayU256();
        strategyManagerMock.setDeposits(staker, strategiesToReturn, sharesToReturn);
        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
        // delegate from the `staker` to the operator
        cheats.prank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DepositScalingFactorUpdated(staker, strategyMock, WAD);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(defaultOperator, staker, strategyMock, shares);
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
        IStrategy[] memory strategiesToReturn = strategyMock.toArray();
        uint256[] memory sharesToReturn = shares.toArrayU256();
        strategyManagerMock.setDeposits(staker, strategiesToReturn, sharesToReturn);

        // Set the operators magnitude to be 0
        _setOperatorMagnitude(defaultOperator, strategyMock, 0);

        // delegate from the `staker` to the operator
        cheats.prank(staker);
        cheats.expectRevert(FullySlashed.selector);
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
        cheats.expectRevert(FullySlashed.selector);
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
    function testFuzz_OperatorWhoAcceptsAllStakers_AlreadySlashed_StrategyManagerShares(
        address staker,
        uint128 shares,
        uint64 maxMagnitude
    ) public filterFuzzedAddressInputs(staker) {
        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);
        maxMagnitude = uint64(bound(uint256(maxMagnitude), 1, WAD));
        shares = uint128(bound(uint256(shares), 1000, MAX_STRATEGY_SHARES));

        // Set empty sig+salt
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        bytes32 salt;

        _registerOperatorWithBaseDetails(defaultOperator);

        // Set staker shares in StrategyManager
        IStrategy[] memory strategiesToReturn = strategyMock.toArray();
        uint256[] memory sharesToReturn = shares.toArrayU256();
        strategyManagerMock.setDeposits(staker, strategiesToReturn, sharesToReturn);
        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);

        // Set the operators magnitude
        _setOperatorMagnitude(defaultOperator, strategyMock, maxMagnitude);

        // Expected staker scaling factor
        uint256 stakerScalingFactor = uint256(WAD).divWad(maxMagnitude);

        // delegate from the `staker` to the operator
        cheats.prank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(defaultOperator, staker, strategyMock, shares);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DepositScalingFactorUpdated(staker, strategyMock, stakerScalingFactor);
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

        (uint256[] memory withdrawableShares, ) = delegationManager.getWithdrawableShares(staker, strategiesToReturn);
        // max delta of 0.1% given minimum shares bound is 1000
        assertApproxEqRel(withdrawableShares[0], shares, 1e15, "staker shares not set correctly");
    }

    /**
     * @notice `staker` delegates to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * via the `staker` calling `DelegationManager.delegateTo`
     * The function should pass with any `operatorSignature` input (since it should be unused)
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * OperatorSharesIncreased event should only be emitted if beaconShares is > 0. Since a staker can have negative shares nothing should happen in that case
     */
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
        // Set the operators magnitude
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, WAD);

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
            emit DepositScalingFactorUpdated(staker, beaconChainETHStrategy, WAD);
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(defaultOperator, staker, beaconChainETHStrategy, uint256(beaconShares));
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
    function testFuzz_OperatorWhoAcceptsAllStakers_AlreadySlashed_BeaconChainStrategyShares(
        address staker,
        int256 beaconShares,
        uint64 maxMagnitude
    ) public filterFuzzedAddressInputs(staker) {
        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);
        maxMagnitude = uint64(bound(uint256(maxMagnitude), 1, WAD));
        // Bound and ensure beaconShares rounded down to gwei
        beaconShares = int256(bound(uint256(beaconShares), 1 gwei, MAX_ETH_SUPPLY));
        beaconShares = beaconShares - (beaconShares % 1 gwei);

        // Set empty sig+salt
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        bytes32 salt;

        // Register and set operator's magnitude
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, maxMagnitude);
        // Expected staker depositScalingFactor
        uint256 stakerScalingFactor = uint256(WAD).divWad(maxMagnitude);

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
            emit DepositScalingFactorUpdated(staker, beaconChainETHStrategy, stakerScalingFactor);
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

        IStrategy[] memory strategiesToReturn = beaconChainETHStrategy.toArray();
        (uint256[] memory withdrawableShares, ) = delegationManager.getWithdrawableShares(staker, strategiesToReturn);
        if (beaconShares > 0) {
            // max delta of 0.01% given minimum shares bound of 1 gwei
            assertApproxEqRel(withdrawableShares[0], uint256(beaconShares), 1e14, "staker shares not set correctly");
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
        IStrategy[] memory strategiesToReturn = strategyMock.toArray();
        uint256[] memory sharesToReturn = shares.toArrayU256();
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
    function testFuzz_OperatorWhoAcceptsAllStakers_AlreadySlashed_BeaconChainAndStrategyManagerShares(
        int256 beaconShares,
        uint128 shares,
        uint64 maxMagnitudeBeacon,
        uint64 maxMagnitudeStrategy
    ) public {
        shares = uint128(bound(uint256(shares), 10000, MAX_STRATEGY_SHARES));
        maxMagnitudeBeacon = uint64(bound(uint256(maxMagnitudeBeacon), 1, WAD));
        maxMagnitudeStrategy = uint64(bound(uint256(maxMagnitudeStrategy), 1, WAD));
        // Bound and ensure beaconShares rounded down to gwei
        beaconShares = int256(bound(uint256(beaconShares), 1 gwei, MAX_ETH_SUPPLY));
        beaconShares = beaconShares - (beaconShares % 1 gwei);

        _registerOperatorWithBaseDetails(defaultOperator);

        // Set empty sig+salt
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        bytes32 salt;

        // Set the operators magnitude to be 50%
        _setOperatorMagnitude(defaultOperator, beaconChainETHStrategy, maxMagnitudeBeacon);
        _setOperatorMagnitude(defaultOperator, strategyMock, maxMagnitudeStrategy);

        // Set staker shares in BeaconChainStrategy and StrategyMananger
        IStrategy[] memory strategiesToReturn = strategyMock.toArray();
        uint256[] memory sharesToReturn = shares.toArrayU256();
        strategyManagerMock.setDeposits(staker, strategiesToReturn, sharesToReturn);
        eigenPodManagerMock.setPodOwnerShares(staker, beaconShares);
        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
        uint256 beaconSharesBefore = delegationManager.operatorShares(defaultStaker, beaconChainETHStrategy);
        // delegate from the `staker` to the operator
        cheats.startPrank(defaultStaker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(defaultStaker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(defaultOperator, defaultStaker, strategyMock, shares);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DepositScalingFactorUpdated(defaultStaker, strategyMock, uint256(WAD).divWad(maxMagnitudeStrategy));
        if (beaconShares > 0) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(defaultOperator, defaultStaker, beaconChainETHStrategy, uint256(beaconShares));
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit DepositScalingFactorUpdated(defaultStaker, beaconChainETHStrategy, uint256(WAD).divWad(maxMagnitudeBeacon));
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
        assertTrue(delegationManager.isOperator(defaultOperator), "defaultStaker not registered as operator");
        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "defaultStaker delegated to the wrong address");
        assertFalse(delegationManager.isOperator(defaultStaker), "staker incorrectly registered as operator");
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
        (uint256[] memory withdrawableShares, ) = delegationManager.getWithdrawableShares(defaultStaker, strategiesToCheck);
        if (beaconShares > 0) {
            // max delta of 0.01% given minimum shares bound of 1 gwei
            assertApproxEqRel(withdrawableShares[0], uint256(beaconShares), 1e14, "staker shares not set correctly");
        } else {
            assertEq(withdrawableShares[0], 0, "staker beacon chain shares not set correctly");
        }
        assertApproxEqRel(withdrawableShares[1], shares, 1e14, "staker shares not set correctly");
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
        cheats.expectRevert(ISignatureUtils.SignatureExpired.selector);
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
        cheats.expectRevert(SaltSpent.selector);
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
        IStrategy[] memory strategiesToReturn = strategyMock.toArray();
        uint256[] memory sharesToReturn = shares.toArrayU256();
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
            IStrategy[] memory strategiesToReturn = strategyMock.toArray();
            uint256[] memory sharesToReturn = shares.toArrayU256();
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
        cheats.expectRevert(ISignatureUtils.SignatureExpired.selector);
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
        cheats.expectRevert(SaltSpent.selector);
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

        OperatorDetails memory operatorDetails = OperatorDetails({
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

        OperatorDetails memory operatorDetails = OperatorDetails({
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

contract DelegationManagerUnitTests_ShareAdjustment is DelegationManagerUnitTests {
    using SingleItemArrayLib for *;
    using SlashingLib for *;

    // @notice Verifies that `DelegationManager.increaseDelegatedShares` reverts if not called by the StrategyManager nor EigenPodManager
    function testFuzz_increaseDelegatedShares_revert_invalidCaller(
        address invalidCaller,
        uint256 shares
    ) public filterFuzzedAddressInputs(invalidCaller) {
        cheats.assume(invalidCaller != address(strategyManagerMock));
        cheats.assume(invalidCaller != address(eigenPodManagerMock));
        cheats.assume(invalidCaller != address(eigenLayerProxyAdmin));

        cheats.expectRevert(OnlyStrategyManagerOrEigenPodManager.selector);
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

        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DepositScalingFactorUpdated(staker, strategyMock, WAD);
        if (delegationManager.isDelegated(staker)) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(defaultOperator, staker, strategyMock, shares);
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
            uint256 slashingFactor = _getSlashingFactor(staker, strategyMock, magnitude);
            dsf.update(0, shares, slashingFactor);
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit DepositScalingFactorUpdated(staker, strategyMock, dsf.scalingFactor());
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
        cheats.expectRevert(FullySlashed.selector);
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
        IStrategy[] memory strategiesDeposited = strategyMock.toArray();
        uint256[] memory sharesToReturn = existingShares.toArrayU256();
        strategyManagerMock.setDeposits(staker, strategiesDeposited, sharesToReturn);
        (uint256[] memory withdrawableShares, ) = delegationManager.getWithdrawableShares(staker, strategiesDeposited);
        if (existingShares < 1e18) {
            // Check that withdrawable shares are within 1 share for amounts < 1e18
            // TODO @michael
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
        cheats.expectRevert(FullySlashed.selector);
        delegationManager.increaseDelegatedShares(staker, strategyMock, existingShares, shares);

        (withdrawableShares, ) = delegationManager.getWithdrawableShares(staker, strategiesDeposited);
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

        uint256 slashingFactor = _getSlashingFactor(staker, strategyMock, magnitude);
        dsf.update(0, shares, slashingFactor);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DepositScalingFactorUpdated(staker, strategy, dsf.scalingFactor());
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(defaultOperator, staker, strategy, shares);

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

    /// @notice Verifies that `DelegationManager.decreaseAndBurnOperatorShares` reverts if not called by the AllocationManager
    function testFuzz_Revert_decreaseAndBurnOperatorShares_invalidCaller(
        address invalidCaller
    ) public filterFuzzedAddressInputs(invalidCaller) {
        cheats.assume(invalidCaller != address(allocationManagerMock));

        cheats.startPrank(invalidCaller);
        cheats.expectRevert(IDelegationManagerErrors.OnlyAllocationManager.selector);
        delegationManager.decreaseAndBurnOperatorShares(invalidCaller, strategyMock, 0, 0);
    }

    /// @notice Verifies that there is no change in shares if the staker is not delegatedd
    function testFuzz_Revert_decreaseAndBurnOperatorShares_noop() public {
        _registerOperatorWithBaseDetails(defaultOperator);

        cheats.prank(address(allocationManagerMock));
        delegationManager.decreaseAndBurnOperatorShares(defaultOperator, strategyMock, WAD, WAD/2);
        assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "shares should not have changed");
    }

    function testFuzz_Revert_decreaseAndBurnOperatorShares_InvalidNewMaxMagnitude(
        uint64 prevMaxMagnitude,
        uint64 newMaxMagnitude
    ) public {
        prevMaxMagnitude = bound(prevMaxMagnitude, 1, WAD);
        newMaxMagnitude = bound(newMaxMagnitude, 0, prevMaxMagnitude - 1);

        cheats.prank(address(allocationManagerMock));
        cheats.expectRevert(IDelegationManagerErrors.MaxMagnitudeCantIncrease.selector);
        delegationManager.decreaseAndBurnOperatorShares(invalidCaller, strategyMock, prevMaxMagnitude, newMaxMagnitude);        
    }

    /**
     * @notice Verifies that `DelegationManager.decreaseAndBurnOperatorShares` properly decreases the delegated `shares` that the operator
     * who the `defaultStaker` is delegated to has in the strategies
     */
    function testFuzz_decreaseAndBurnOperatorShares_slashedOperator(
        IStrategy[] memory strategies,
        uint128 shares,
        uint64 previousMaxMagnitude,
        uint256 wadToSlash
    ) public {
        // sanity-filtering on fuzzed input length & staker
        cheats.assume(strategies.length > 0 && strategies.length <= 16);
        shares = uint128(bound(shares, 1, MAX_STRATEGY_SHARES));
        previousMaxMagnitude = uint64(bound(previousMaxMagnitude, WAD, WAD));
        wadToSlash = bound(wadToSlash, 1, WAD);
        
        bool hasBeaconChainStrategy = false;
        for(uint256 i = 0; i < numStrats; i++) {
            if (strategies[i] == beaconChainETHStrategy) {
                hasBeaconChainStrategy = true;
                // Swap beacon chain strategy to the end of the array
                strategies[i] = strategies[numStrats - 1];
                strategies[numStrats - 1] = beaconChainETHStrategy;
                
                // Resize
                assembly {
                    mstore(strategies, sub(mload(strategies), 1))
                }
                break;
            }
        }

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // Set the staker deposits in the strategies
        uint256[] memory sharesToSet = new uint256[](strategies.length);
        for(uint256 i = 0; i < strategies.length; i++) {
            sharesToSet[i] = shares;
            _setOperatorMagnitude(defaultOperator, strategies[i], previousMaxMagnitude);
        }

        // Okay to set beacon chain shares in SM mock, wont' be called by DM
        strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
        if (hasBeaconChainStrategy) {
            eigenPodManagerMock.setPodOwnerShares(defaultStaker, int256(uint256(shares)));
        }

        // events expected emitted for each strategy
        for (uint256 i = 0; i < strategies.length; ++i) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(defaultOperator, defaultStaker, strategies[i], shares);
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit DepositScalingFactorUpdated(defaultStaker, strategies[i], uint256(WAD).divWad(uint256(previousMaxMagnitude)));
        }
        // delegate from the `staker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        address delegatedTo = delegationManager.delegatedTo(defaultStaker);

        // check shares before call to `decreaseOperatorShares`
        for (uint256 i = 0; i < strategies.length; ++i) {
            // store delegated shares in a mapping
            delegatedSharesBefore[strategies[i]] = delegationManager.operatorShares(delegatedTo, strategies[i]);
            // also construct an array which we'll use in another loop
            totalSharesForStrategyInArray[address(strategies[i])] += shares;
        }

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
                    delegationManager.burnOperatorShares(defaultOperator, strategies[i], WAD, WAD / 2);
                    totalSharesDecreasedForStrategy[strategies[i]] += currentShares / 2;
                }
                cheats.expectEmit(true, true, true, true, address(delegationManager));
                emit OperatorSharesDecreased(
                    defaultOperator,
                    address(0),
                    strategies[i],
                    sharesToDecrease
                );
                // operator shares decrease as a result of maxMagnitude decreasing
                delegationManager.decreaseOperatorShares(
                    defaultOperator,
                    strategies[i],
                    sharesWadSlashed
                );
                // Also update maxMagnitude in ALM mock
                allocationManagerMock.setMaxMagnitude(defaultOperator, strategies[i], newMaxMagnitude);

                totalSharesDecreasedForStrategy[strategies[i]] += sharesToDecrease;
            }
            cheats.stopPrank();
        }

        // check shares after call to `burnOperatorShares`
        (uint256[] memory withdrawableShares, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        for (uint256 i = 0; i < strategies.length; ++i) {
            uint256 delegatedSharesAfter = delegationManager.operatorShares(delegatedTo, strategies[i]);
            assertEq(
                delegatedSharesAfter,
                delegatedSharesBefore[strategies[i]] - totalSharesDecreasedForStrategy[strategies[i]],
                "delegated shares did not decrement correctly"
            );

            assertLe(
                withdrawableShares[i],
                delegatedSharesAfter,
                "withdrawable shares for staker should be <= operatorShares due to potential rounding"
            );
        }
    }
}

contract DelegationManagerUnitTests_Undelegate is DelegationManagerUnitTests {
    using SlashingLib for uint256;
    using SingleItemArrayLib for *;

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
        cheats.expectRevert(NotActivelyDelegated.selector);
        delegationManager.undelegate(undelegatedStaker);
    }

    // @notice Verifies that an operator cannot undelegate from themself (this should always be forbidden)
    function testFuzz_undelegate_revert_stakerIsOperator(address operator) public filterFuzzedAddressInputs(operator) {
        _registerOperatorWithBaseDetails(operator);

        cheats.prank(operator);
        cheats.expectRevert(OperatorsCannotUndelegate.selector);
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
        cheats.expectRevert(OperatorsCannotUndelegate.selector);
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
        cheats.expectRevert(CallerCannotUndelegate.selector);
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
        IStrategy[] memory strategies = strategyMock.toArray();
        uint256[] memory sharesToSet = shares.toArrayU256();
        strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);

        // register *this contract* as an operator and delegate from the `staker` to them
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        
        // Format queued withdrawal
        (
            QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategyMock,
            depositSharesToWithdraw: shares
        });

        // Undelegate the staker
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(defaultStaker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(defaultOperator, defaultStaker, strategyMock, shares);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, queuedWithdrawalParams[0].depositShares);
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
        (uint256[] memory stakerWithdrawableShares, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        assertEq(stakerWithdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");
    }

    /**
     * @notice Verifies that the `undelegate` function properly queues a withdrawal for all shares of the staker
     * @notice The operator should have its shares slashed prior to the staker's deposit
     */
    function testFuzz_undelegate_preSlashedOperator(
        uint256 shares,
        uint64 operatorMagnitude
    ) public {
        shares = bound(shares, 1, MAX_STRATEGY_SHARES);
        operatorMagnitude = uint64(bound(operatorMagnitude, 1, WAD));

        // register *this contract* as an operator & set its slashed magnitude
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, operatorMagnitude);
    
        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = strategyMock.toArray();
        {
            uint256[] memory sharesToSet = shares.toArrayU256();
            strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
        }

        // delegate from the `staker` to them
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        uint256 depositScalingFactor = delegationManager.depositScalingFactor(defaultStaker, strategyMock);
        assertTrue(depositScalingFactor > WAD, "bad test setup");
        
        // Format queued withdrawal
        (
            QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategyMock,
            depositSharesToWithdraw: shares
        });

        // Calculate operatorShares decreased, may be off of shares due to rounding
        StakerScalingFactors memory ssf = StakerScalingFactors({
            depositScalingFactor: depositScalingFactor,
            isBeaconChainScalingFactorSet: false,
            beaconChainScalingFactor: 0
        });
        uint256 operatorSharesDecreased = shares.toShares(ssf, operatorMagnitude);
        assertLe(operatorSharesDecreased, shares, "operatorSharesDecreased should be <= shares");

        // Undelegate the staker
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(defaultStaker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(defaultOperator, defaultStaker, strategyMock, operatorSharesDecreased);
        cheats.expectEmit(true, true, true, true, address(delegationManager));  
        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, queuedWithdrawalParams[0].depositShares);
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
        assertEq(
            delegationManager.operatorShares(defaultOperator, strategyMock),
            shares - operatorSharesDecreased,
            "operator shares not decreased correctly"
        );
        (uint256[] memory stakerWithdrawableShares, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        assertEq(stakerWithdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");
    }

    /**
     * @notice Verifies that the `undelegate` function properly queues a withdrawal for all shares of the staker
     * @notice The operator should have its shares slashed prior to the staker's deposit
     */
    function testFuzz_undelegate_slashedWhileStaked(
        uint256 shares,
        uint64 operatorMagnitude,
        uint256 wadToSlash
    ) public {
        shares = bound(shares, 1, MAX_STRATEGY_SHARES);
        operatorMagnitude = uint64(bound(operatorMagnitude, 2, WAD));
        wadToSlash = bound(wadToSlash, 1, WAD - 1);

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, operatorMagnitude);
    
        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = strategyMock.toArray();
        {
            uint256[] memory sharesToSet = shares.toArrayU256();
            strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
        }
        
        // delegate from the `defaultStaker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        assertEq(
            delegationManager.operatorShares(defaultOperator, strategyMock),
            shares,
            "operatorShares should increment correctly"
        );

        // Set operator magnitude
        uint256 operatorSharesAfterSlash;
        {
            (
                uint256 sharesToDecrement,
                uint64 newOperatorMagnitude,
                uint256 wadSlashed
            ) = _calculateSlashedSharesAndMagnitude(shares, operatorMagnitude, wadToSlash);
            // even though wadToSlash isn't 100%, it could round upwards to end up with 0 leftover magnitude
            cheats.assume(newOperatorMagnitude > 0);
            _setOperatorMagnitude(defaultOperator, strategyMock, newOperatorMagnitude);

            cheats.prank(address(allocationManagerMock));
            delegationManager.burnOperatorShares(defaultOperator, strategyMock, WAD, operatorMagnitude);
            operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategyMock);
            // assertEq(operatorSharesAfterSlash, operatorSharesBefore / 2, "operator shares not properly updated");
            assertEq(
                shares - sharesToDecrement,
                operatorSharesAfterSlash
            );
        }

        (uint256 depositScalingFactor,,) = delegationManager.stakerScalingFactor(defaultStaker, strategyMock);
        assertEq(depositScalingFactor, uint256(WAD).divWad(operatorMagnitude), "bad test setup");

        // Get withdrawable shares
        (uint256[] memory withdrawableSharesBefore, uint256[] memory depositShares) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        
        // Format queued withdrawal
        (
            ,
            Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategyMock,
            depositSharesToWithdraw: shares
        });

        uint256[] memory sharesToWithdraw = uint256(shares / 2).toArrayU256();
        // Undelegate the staker
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(defaultStaker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(defaultOperator, defaultStaker, strategyMock, withdrawableSharesBefore[0]);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, sharesToWithdraw);
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
        assertEq(
            delegationManager.operatorShares(defaultOperator, strategyMock),
            operatorSharesAfterSlash - withdrawableSharesBefore[0], // Note that this is not exactly 0 due to rounding
            "operator shares not decreased correctly"
        );
        (uint256[] memory stakerWithdrawableShares, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        assertEq(stakerWithdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");
    }

    /**
     * @notice Verifies that the `undelegate` function properly undelegates a staker even though their shares
     * were slashed entirely.
     */
    function testFuzz_undelegate_slashedOperator100PercentWhileStaked(uint128 shares) public {
        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
    
        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = strategyMock.toArray();
        {
            uint256[] memory sharesToSet = shares.toArrayU256();
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
            delegationManager.burnOperatorShares(defaultOperator, strategyMock, WAD, 0);
            operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategyMock);
            assertEq(operatorSharesAfterSlash, 0, "operator shares not fully slashed");
        }

        uint256 depositScalingFactor = delegationManager.depositScalingFactor(defaultStaker, strategyMock);
        assertEq(depositScalingFactor, WAD, "bad test setup");

        // Get withdrawable shares
        (uint256[] memory withdrawableSharesBefore, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        assertEq(
            withdrawableSharesBefore[0],
            0,
            "withdrawable shares should be 0 after being slashed fully"
        );

        // Undelegate the staker
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(defaultStaker, defaultOperator);
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
        (uint256[] memory stakerWithdrawableShares, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        assertEq(stakerWithdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");
    }

    function testFuzz_undelegate_slashedOperatorCloseTo100(
        uint128 shares,
        uint64 initialMagnitude,
        uint256 wadToSlashRequest,
        address[] memory stakers,
        uint8 numStakers
    ) public {
        cheats.assume(stakers.length > 8);
        numStakers = uint8(bound(uint256(numStakers), 1, 8));
        initialMagnitude = uint64(bound(uint256(initialMagnitude), 1, WAD));
        wadToSlashRequest = bound(wadToSlashRequest, 1, WAD - 1);
        numStakers = 0;

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, initialMagnitude);

        console.log("INITIAL OP MAGNITUDE", initialMagnitude);
    
        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = strategyMock;
        {
            uint256[] memory sharesToSet = new uint256[](1);
            sharesToSet[0] = shares;

            strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);

            for (uint256 i = 0; i < numStakers; ++i) {
                stakers[i] = address(uint160(uint256(bytes32(keccak256(abi.encode(stakers[i], i))))));
                sharesToSet[0] = uint256(bytes32(keccak256(abi.encode(stakers[i], shares))));
                sharesToSet[0] = bound(sharesToSet[0], 1, 1e38 - 1);
                console.log("STAKER ADDRESS", i);
                console.log(stakers[i]);

                strategyManagerMock.setDeposits(
                    stakers[i],
                    strategies,
                    sharesToSet
                );
            }
        }

        // delegate from the `defaultStaker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        for (uint256 i = 0; i < numStakers; ++i) {
            _delegateToOperatorWhoAcceptsAllStakers(stakers[i], defaultOperator);
        }

        uint256 totalWithdrawable = 0;
        for (uint256 i = 0; i < numStakers; ++i) {
            (uint256[] memory withdrawableSharesBefore, ) = delegationManager.getWithdrawableShares(stakers[i], strategies);

            totalWithdrawable += withdrawableSharesBefore[0];
        }
        {
            (uint256[] memory withdrawableSharesBefore, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
            totalWithdrawable += withdrawableSharesBefore[0];
        }
        assertLe(
            totalWithdrawable, delegationManager.operatorShares(defaultOperator, strategyMock), "should be <= op shares due to rounding"
        );

        // Set operator magnitude
        uint256 operatorSharesAfterSlash;
        {
            uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);
            console.log("OP SHARES BEFORE");
            console.log(operatorSharesBefore);
            (
                uint256 sharesToDecrement,
                uint64 newOperatorMagnitude,
                uint256 wadSlashed
            ) = _calculateSlashedSharesAndMagnitude(shares, initialMagnitude, wadToSlashRequest);
            cheats.assume(newOperatorMagnitude != 0);
            _setOperatorMagnitude(defaultOperator, strategyMock, newOperatorMagnitude);

            cheats.prank(address(allocationManagerMock));
            delegationManager.decreaseOperatorShares(defaultOperator, strategyMock, wadSlashed);
            operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategyMock);

            (uint256[] memory withdrawableSharesBefore, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);

            console.log("STAKER WITHDRAWABLE");
            console.log(
                withdrawableSharesBefore[0]
            );
            console.log("OP SHARES");
            console.log(operatorSharesAfterSlash);
        }

        // (uint256 depositScalingFactor,,) = delegationManager.stakerScalingFactor(defaultStaker, strategyMock);
        // assertEq(depositScalingFactor, WAD, "bad test setup");

        // Undelegate the staker
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(defaultStaker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit DepositScalingFactorUpdated(defaultStaker, strategyMock, WAD);
        cheats.prank(defaultStaker);
        delegationManager.undelegate(defaultStaker);

        for (uint256 i = 0; i < numStakers; ++i) {
            console.log("UNDELEGATING ", i);
            cheats.prank(stakers[i]);
            delegationManager.undelegate(stakers[i]);
        }

        // Checks - delegation status
        assertEq(
            delegationManager.delegatedTo(defaultStaker),
            address(0),
            "undelegated staker should be delegated to zero address"
        );
        assertFalse(delegationManager.isDelegated(defaultStaker), "staker not undelegated");

        // Checks - operator & staker shares
        // assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "operator shares not decreased correctly");
        // (uint256[] memory stakerWithdrawableShares, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        // assertEq(stakerWithdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");
        // (uint256 newDepositScalingFactor,,) = delegationManager.stakerScalingFactor(defaultStaker, strategyMock);
        // assertEq(newDepositScalingFactor, WAD, "staker scaling factor not reset correctly");
    }

    function testFuzz_undelegate_slashedRepeatedly(
        // uint128 shares,
        uint64 initialMagnitude
    ) public {
        initialMagnitude = uint64(bound(uint256(initialMagnitude), 10000, WAD));
        initialMagnitude = 999999999999990009;
        uint256 shares = 44182209037560531097078597505;

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, initialMagnitude);

        console.log("INITIAL OP MAGNITUDE", initialMagnitude);
    
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

        {
            (uint256[] memory withdrawableSharesBefore, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);

                console.log("BEFORE SLASH STAKER WITHDRAWABLE");
                console.log(
                    withdrawableSharesBefore[0]
                );
                console.log("BEFORE SLASH OP SHARES");
                console.log(delegationManager.operatorShares(defaultOperator, strategyMock));
        }

        // Set operator magnitude
        uint256 operatorSharesAfterSlash;
        {
            uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);

            uint64 newOperatorMagnitude = initialMagnitude;
            uint256 wadSlashed;
            for (uint256 i = 0; i < 1000; ++i) {
                (
                    ,
                    newOperatorMagnitude,
                    wadSlashed
                ) = _calculateSlashedSharesAndMagnitude(shares, newOperatorMagnitude, 1);
                cheats.assume(newOperatorMagnitude != 0);
                _setOperatorMagnitude(defaultOperator, strategyMock, newOperatorMagnitude);

                cheats.prank(address(allocationManagerMock));
                delegationManager.decreaseOperatorShares(defaultOperator, strategyMock, wadSlashed);


                operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategyMock);

                (uint256[] memory withdrawableSharesBefore, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);

                console.log("SLASHED NUMBER OF TIMES: ", i);

                console.log("STAKER WITHDRAWABLE");
                console.log(
                    withdrawableSharesBefore[0]
                );
                console.log("OP SHARES");
                console.log(operatorSharesAfterSlash);

                console.log("==========> SHARES DRIFT", operatorSharesAfterSlash - withdrawableSharesBefore[0]);
                console.log("INITIAL MAG", initialMagnitude);
                console.log("CURR MAG", newOperatorMagnitude);
            }


            // (
            //     uint256 sharesToDecrement,
            //     uint64 newOperatorMagnitude,
            //     uint256 wadSlashed
            // ) = _calculateSlashedSharesAndMagnitude(shares, initialMagnitude, 1);
            // cheats.assume(newOperatorMagnitude != 0);
            // _setOperatorMagnitude(defaultOperator, strategyMock, newOperatorMagnitude);

            // cheats.prank(address(allocationManagerMock));
            // delegationManager.decreaseOperatorShares(defaultOperator, strategyMock, wadSlashed);
            // operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategyMock);

            // console.log("SHARES DRIFT", operatorSharesAfterSlash - withdrawableSharesBefore[0]);
            // (uint256[] memory withdrawableSharesBefore, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);

            // console.log("STAKER WITHDRAWABLE");
            // console.log(
            //     withdrawableSharesBefore[0]
            // );
            // console.log("OP SHARES");
            // console.log(operatorSharesAfterSlash);

            // console.log("==========> SHARES DRIFT", operatorSharesAfterSlash - withdrawableSharesBefore[0]);
        }

        // (uint256 depositScalingFactor,,) = delegationManager.stakerScalingFactor(defaultStaker, strategyMock);
        // assertEq(depositScalingFactor, WAD, "bad test setup");

        // Undelegate the staker
        // cheats.expectEmit(true, true, true, true, address(delegationManager));
        // emit StakerUndelegated(defaultStaker, defaultOperator);
        // cheats.expectEmit(true, true, true, true, address(delegationManager));
        // emit DepositScalingFactorUpdated(defaultStaker, strategyMock, WAD);
        // cheats.prank(defaultStaker);
        // delegationManager.undelegate(defaultStaker);

        // for (uint256 i = 0; i < numStakers; ++i) {
        //     console.log("UNDELEGATING ", i);
        //     cheats.prank(stakers[i]);
        //     delegationManager.undelegate(stakers[i]);
        // }

        // // Checks - delegation status
        // assertEq(
        //     delegationManager.delegatedTo(defaultStaker),
        //     address(0),
        //     "undelegated staker should be delegated to zero address"
        // );
        // assertFalse(delegationManager.isDelegated(defaultStaker), "staker not undelegated");

        // Checks - operator & staker shares
        // assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "operator shares not decreased correctly");
        // (uint256[] memory stakerWithdrawableShares, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        // assertEq(stakerWithdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");
        // (uint256 newDepositScalingFactor,,) = delegationManager.stakerScalingFactor(defaultStaker, strategyMock);
        // assertEq(newDepositScalingFactor, WAD, "staker scaling factor not reset correctly");
    }


    function testFuzz_undelegate_depositRepeatedly(
        // uint128 shares,
        uint64 initialMagnitude
    ) public {
        initialMagnitude = uint64(bound(uint256(initialMagnitude), 10000, WAD));
        initialMagnitude = 999999999999990009;
        // uint256 shares = 44182209037560531097078597505;
        uint256 shares = 4418220903756053;
        // uint256 shares = 1e38 - 1000;
        // uint256 shares = 1000000000000000000;

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, initialMagnitude);

        console.log("INITIAL OP MAGNITUDE", initialMagnitude);
    
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

        {
            (uint256[] memory withdrawableSharesBefore, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);

            console.log("BEFORE SLASH STAKER WITHDRAWABLE");
            console.log(withdrawableSharesBefore[0]);
            console.log("BEFORE SLASH OP SHARES");
            console.log(delegationManager.operatorShares(defaultOperator, strategyMock));
        }

        // Set operator magnitude
        uint256 operatorSharesAfterSlash;
        {
            uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);

            uint64 newOperatorMagnitude = initialMagnitude;
            uint256 wadSlashed;

            uint256 newDepositShares = shares;
            for (uint256 i = 0; i < 1000; ++i) {
                // (
                //     ,
                //     newOperatorMagnitude,
                //     wadSlashed
                // ) = _calculateSlashedSharesAndMagnitude(shares, newOperatorMagnitude, 1);
                // cheats.assume(newOperatorMagnitude != 0);
                // _setOperatorMagnitude(defaultOperator, strategyMock, newOperatorMagnitude);
                cheats.prank(address(strategyManagerMock));
                delegationManager.increaseDelegatedShares(defaultStaker, strategyMock, newDepositShares, 1000);
                newDepositShares += 1000;

                uint256[] memory newDepositSharesArray = new uint256[](1);
                newDepositSharesArray[0] = newDepositShares;

                strategyManagerMock.setDeposits(defaultStaker, strategies, newDepositSharesArray);


                operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategyMock);

                (uint256[] memory withdrawableSharesBefore, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);

                console.log("SLASHED NUMBER OF TIMES: ", i);

                console.log("STAKER WITHDRAWABLE");
                console.log(
                    withdrawableSharesBefore[0]
                );
                // console.log("OP SHARES");
                // console.log(operatorSharesAfterSlash);

                console.log("==========> SHARES DRIFT", operatorSharesAfterSlash - withdrawableSharesBefore[0]);
                // console.log("INITIAL MAG", initialMagnitude);
                // console.log("CURR MAG", newOperatorMagnitude);
            }


            // (
            //     uint256 sharesToDecrement,
            //     uint64 newOperatorMagnitude,
            //     uint256 wadSlashed
            // ) = _calculateSlashedSharesAndMagnitude(shares, initialMagnitude, 1);
            // cheats.assume(newOperatorMagnitude != 0);
            // _setOperatorMagnitude(defaultOperator, strategyMock, newOperatorMagnitude);

            // cheats.prank(address(allocationManagerMock));
            // delegationManager.decreaseOperatorShares(defaultOperator, strategyMock, wadSlashed);
            // operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategyMock);

            // console.log("SHARES DRIFT", operatorSharesAfterSlash - withdrawableSharesBefore[0]);
            // (uint256[] memory withdrawableSharesBefore, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);

            // console.log("STAKER WITHDRAWABLE");
            // console.log(
            //     withdrawableSharesBefore[0]
            // );
            // console.log("OP SHARES");
            // console.log(operatorSharesAfterSlash);

            // console.log("==========> SHARES DRIFT", operatorSharesAfterSlash - withdrawableSharesBefore[0]);
        }
    }

    /**
     * @notice Given an operator with slashed magnitude, delegate, undelegate, and then delegate back to the same operator with
     * completing withdrawals as shares. This should result in the operatorShares after the second delegation being <= the shares from the first delegation.
     */
    function testFuzz_undelegate_delegateAgainWithRounding(uint128 shares) public {
        // set magnitude to 66% to ensure rounding when calculating `toShares`
        uint64 operatorMagnitude = 333333333333333333;

        // register *this contract* as an operator & set its slashed magnitude
        _registerOperatorWithBaseDetails(defaultOperator);
        _setOperatorMagnitude(defaultOperator, strategyMock, operatorMagnitude);
    
        // Set the staker deposits in the strategies
        IStrategy[] memory strategies = strategyMock.toArray();
        {
            uint256[] memory sharesToSet = shares.toArrayU256();
            strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
        }

        // delegate from the `staker` to them
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        uint256 depositScalingFactor = delegationManager.depositScalingFactor(defaultStaker, strategyMock);
        assertEq(depositScalingFactor, uint256(WAD).divWad(uint256(operatorMagnitude)), "first deposit should result in k value of (1 / magnitude)");
        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, strategyMock);

        // Format queued withdrawal
        (
            ,
            Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategyMock,
            depositSharesToWithdraw: shares
        });

        DepositScalingFactor memory _dsf = DepositScalingFactor(depositScalingFactor);
        uint256 slashingFactor = _getSlashingFactor(defaultStaker, strategyMock, operatorMagnitude);
        uint256 operatorSharesDecreased = _dsf.calcWithdrawable(shares, slashingFactor);
        uint256[] memory sharesToWithdraw = operatorSharesDecreased.toArrayU256();

        // Undelegate the staker
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(defaultStaker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(defaultOperator, defaultStaker, strategyMock, operatorSharesDecreased);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, sharesToWithdraw);
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
        (uint256[] memory stakerWithdrawableShares, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        assertEq(stakerWithdrawableShares[0], 0, "staker withdrawable shares not calculated correctly");

        // // Re-delegate the staker to the operator again. The shares should have increased but may be less than from before due to rounding
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        // complete withdrawal as shares, should add back delegated shares to operator due to delegating again
        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = IERC20(strategies[0].underlyingToken());
        cheats.roll(withdrawal.startBlock + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS());
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, false);

        uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, strategyMock);
        assertLe(operatorSharesAfter, operatorSharesBefore, "operator shares should be less than or equal to before due to potential rounding");
    }

    // // TODO: fix old Withdrawals.t.sol test
    // // @notice This function tests to ensure that a delegator can re-delegate to an operator after undelegating.
    // // @param operator is the operator being delegated to.
    // // @param staker is the staker delegating stake to the operator.
    // function testFuzz_RedelegateAfterWithdrawal(
    //     Randomness r
    // ) public rand(r) {
    //     address operator = r.Address();
    //     address depositor = r.Address();
    //     uint96 ethAmount = r.Uint96();
    //     uint96 eigenAmount = r.Uint96();
    //     bool withdrawAsShares = r.Boolean();
    //     //this function performs delegation and subsequent withdrawal
    //     testWithdrawalWrapper(operator, depositor, ethAmount, eigenAmount, withdrawAsShares, true);

    //     cheats.prank(depositor);
    //     delegationManager.undelegate(depositor);

    //     //warps past fraudproof time interval
    //     cheats.warp(block.timestamp + 7 days + 1);
    //     _initiateDelegation(operator, depositor, ethAmount, eigenAmount);
    // }
}

contract DelegationManagerUnitTests_queueWithdrawals is DelegationManagerUnitTests {
    using SlashingLib for *;
    using SingleItemArrayLib for *;

    function test_Revert_WhenEnterQueueWithdrawalsPaused() public {
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_ENTER_WITHDRAWAL_QUEUE);
        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, , ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategyMock,
            depositSharesToWithdraw: 100
        });
        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
    }

    function test_Revert_WhenQueueWithdrawalParamsLengthMismatch() public {
        IStrategy[] memory strategyArray = strategyMock.toArray();
        uint256[] memory shareAmounts = new uint256[](2);
        shareAmounts[0] = 100;
        shareAmounts[1] = 100;

        QueuedWithdrawalParams[] memory queuedWithdrawalParams = new QueuedWithdrawalParams[](1);
        queuedWithdrawalParams[0] = QueuedWithdrawalParams({
            strategies: strategyArray,
            depositShares: shareAmounts,
            withdrawer: defaultStaker
        });

        cheats.expectRevert(InputArrayLengthMismatch.selector);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
    }

    function test_Revert_WhenNotStakerWithdrawer(address withdrawer) public {
        cheats.assume(withdrawer != defaultStaker);

        (QueuedWithdrawalParams[] memory queuedWithdrawalParams, , ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: withdrawer,
            strategy: strategyMock,
            depositSharesToWithdraw: 100
        });
        cheats.expectRevert(WithdrawerNotStaker.selector);
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
    }

    function test_Revert_WhenEmptyStrategiesArray() public {
        IStrategy[] memory strategyArray = new IStrategy[](0);
        uint256[] memory shareAmounts = new uint256[](0);
        address withdrawer = defaultOperator;

        QueuedWithdrawalParams[] memory queuedWithdrawalParams = new QueuedWithdrawalParams[](1);
        queuedWithdrawalParams[0] = QueuedWithdrawalParams({
            strategies: strategyArray,
            depositShares: shareAmounts,
            withdrawer: withdrawer
        });

        cheats.expectRevert(InputArrayLengthZero.selector);
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
        Randomness r
    ) public rand(r) {
        uint128 depositAmount = r.Uint128();
        uint128 withdrawalAmount = r.Uint128(1, depositAmount);
        uint256[] memory sharesAmounts = depositAmount.toArrayU256();
        // sharesAmounts is single element so returns single strategy
        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(defaultStaker, sharesAmounts);
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        (
            QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategies[0],
            depositSharesToWithdraw: withdrawalAmount
        });
        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");
        uint256 nonceBefore = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        uint256 delegatedSharesBefore = delegationManager.operatorShares(defaultOperator, strategies[0]);

        // queueWithdrawals
        cheats.prank(defaultStaker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, queuedWithdrawalParams[0].depositShares);
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
     */
    function testFuzz_queueWithdrawal_SingleStrat_preSlashedOperator(
        uint256 depositAmount,
        uint256 withdrawalAmount,
        uint256 wadToSlash
    ) public {
        depositAmount = bound(depositAmount, 1, MAX_STRATEGY_SHARES);
        withdrawalAmount = bound(withdrawalAmount, 1, depositAmount);
        wadToSlash = bound(wadToSlash, 1, WAD - 1);

        // Slash the operator
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            uint256 sharesToDecrement,
            uint64 newOperatorMagnitude,
            uint256 wadSlashed
        ) = _calculateSlashedSharesAndMagnitude(depositAmount, WAD, wadToSlash);
        _setOperatorMagnitude(defaultOperator, strategyMock, newOperatorMagnitude);

        // Deposit for staker & delegate
        IStrategy[] memory strategies = strategyMock.toArray();
        {
            uint256[] memory sharesToSet = depositAmount.toArrayU256();
            strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
            _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        }

        (
            QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategies[0],
            depositSharesToWithdraw: withdrawalAmount
        });

        (uint256 depositScalingFactor,,) = delegationManager.stakerScalingFactor(defaultStaker, strategyMock);
        StakerScalingFactors memory ssf = StakerScalingFactors({
            depositScalingFactor: depositScalingFactor,
            isBeaconChainScalingFactorSet: false,
            beaconChainScalingFactor: 0
        });
        uint256 expectedWithdrawnShares = withdrawalAmount.toShares(ssf, newOperatorMagnitude);

        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");
        uint256 nonceBefore = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        uint256 delegatedSharesBefore = delegationManager.operatorShares(defaultOperator, strategies[0]);

        // queueWithdrawals
        cheats.prank(defaultStaker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, queuedWithdrawalParams[0].depositShares);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        uint256 nonceAfter = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        uint256 delegatedSharesAfter = delegationManager.operatorShares(defaultOperator, strategies[0]);
        assertEq(nonceBefore + 1, nonceAfter, "staker nonce should have incremented");
        assertApproxEqRel(
            delegatedSharesBefore - expectedWithdrawnShares,
            delegatedSharesAfter,
            1e16, // max delta of 1%, given shares amount > 1000
            "delegated shares not decreased correctly"
        );
    }

    /**
     * @notice Verifies that `DelegationManager.queueWithdrawals` properly queues a withdrawal for the `withdrawer`
     * from the `strategy` for the `sharesAmount`. Operator is slashed while the staker is deposited
     * - Asserts that staker is delegated to the operator
     * - Asserts that shares for delegatedTo operator are decreased by `sharesAmount`
     * - Asserts that staker cumulativeWithdrawalsQueued nonce is incremented
     * - Checks that event was emitted with correct withdrawalRoot and withdrawal
     */
    function testFuzz_queueWithdrawal_SingleStrat_slashedWhileStaked(
        uint256 depositAmount,
        uint256 withdrawalAmount,
        uint64 operatorMagnitude,
        uint256 wadToSlash
    ) public {
        depositAmount = bound(depositAmount, 1, MAX_STRATEGY_SHARES);
        withdrawalAmount = bound(withdrawalAmount, 1, depositAmount);
        operatorMagnitude = uint64(bound(operatorMagnitude, 2, WAD));
        wadToSlash = bound(wadToSlash, 1, WAD - 1);

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // Deposit for staker & delegate
        IStrategy[] memory strategies = strategyMock.toArray();
        {
            uint256[] memory sharesToSet = depositAmount.toArrayU256();
            strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
            _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        }

        // Slash the operator
        (
            uint256 sharesToDecrement,
            uint64 newOperatorMagnitude,
        ) = _calculateSlashedSharesAndMagnitude(depositAmount, WAD, wadToSlash);
        _setOperatorMagnitude(defaultOperator, strategyMock, newOperatorMagnitude);
        cheats.prank(address(allocationManagerMock));
        delegationManager.burnOperatorShares(defaultOperator, strategyMock, WAD, operatorMagnitude);

        (
            QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategies[0],
            depositSharesToWithdraw: withdrawalAmount
        });

        uint256[] memory sharesToWithdraw = uint256(withdrawalAmount / 2).toArrayU256();
        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");
        uint256 nonceBefore = delegationManager.cumulativeWithdrawalsQueued(defaultStaker);
        uint256 delegatedSharesBefore = delegationManager.operatorShares(defaultOperator, strategies[0]);

        // queueWithdrawals
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, sharesToWithdraw);
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        uint256 delegatedSharesAfter = delegationManager.operatorShares(defaultOperator, strategies[0]);

        uint256 slashingFactor = _getSlashingFactor(defaultStaker, strategyMock, operatorMagnitude);
        dsf = DepositScalingFactor(delegationManager.depositScalingFactor(defaultStaker, strategyMock));

        uint256 sharesWithdrawn = dsf.calcWithdrawable(withdrawalAmount, slashingFactor);
        assertEq(nonceBefore + 1, nonceAfter, "staker nonce should have incremented");
        assertApproxEqRel(
            delegatedSharesBefore - sharesWithdrawn,
            delegatedSharesAfter,
            "delegated shares not decreased correctly"
        );
    }

    /**
     * @notice Verifies that `DelegationManager.queueWithdrawals` queues an empty withdrawal for the `withdrawer`
     * from the `strategy` for the `sharesAmount` since the Operator is slashed 100% while the staker is deposited
     * - Asserts that queuing a withdrawal results in an empty withdrawal when the operator is slashed 100%
     * - Asserts that staker withdrawableShares after is 0
     * - Checks that event was emitted with correct withdrawalRoot and withdrawal
     */
    function testFuzz_queueWithdrawal_SingleStrat_slashed100PercentWhileStaked(
        uint128 depositAmount,
        uint128 withdrawalAmount
    ) public {
        withdrawalAmount = uint128(bound(withdrawalAmount, 0, depositAmount));
        
        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // Deposit for staker & delegate
        IStrategy[] memory strategies = strategyMock.toArray();
        {
            uint256[] memory sharesToSet = depositAmount.toArrayU256();
            strategyManagerMock.setDeposits(defaultStaker, strategies, sharesToSet);
            _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        }

        // Slash the operator
        uint64 operatorMagnitude = 0;
        _setOperatorMagnitude(defaultOperator, strategyMock, operatorMagnitude);
        cheats.prank(address(allocationManagerMock));
        delegationManager.burnOperatorShares(defaultOperator, strategyMock, WAD, 0);

        // Attempt to withdraw for the strategy that was slashed 100% for the operator
        QueuedWithdrawalParams[] memory queuedWithdrawalParams = new QueuedWithdrawalParams[](1);
        {
            uint256[] memory withdrawalAmounts = withdrawalAmount.toArrayU256();
            queuedWithdrawalParams[0] = QueuedWithdrawalParams({
                strategies: strategies,
                depositShares: withdrawalAmounts,
                withdrawer: defaultStaker
            });
        }

        assertEq(delegationManager.delegatedTo(defaultStaker), defaultOperator, "staker should be delegated to operator");

        // queueWithdrawals should result in an empty withdrawal
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        (uint256[] memory withdrawableShares, ) = delegationManager.getWithdrawableShares(defaultStaker, strategies);
        assertEq(
            withdrawableShares[0],
            0,
            "withdrawable shares should be 0 after being slashed fully"
        );

        (IDelegationManagerTypes.Withdrawal[] memory withdrawals, uint[][] memory shares)
            = delegationManager.getQueuedWithdrawals(defaultStaker);

        assertEq(withdrawals.length, 1, "staker should have a single withdrawal");
        assertEq(shares.length, 1, "output arrays should have equal length");
        assertEq(shares[0].length, 1, "withdrawal should consider a single strategy");
        assertEq(shares[0][0], 0, "withdrawal should be for 0 shares");
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
        uint256[] memory fuzzAmounts
    ) public {
        (
            uint256[] memory depositAmounts,
            uint256[] memory withdrawalAmounts
        ) = _fuzzDepositWithdrawalAmounts(fuzzAmounts);

        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(defaultStaker, depositAmounts);
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        (
            QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawals({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategies: strategies,
            depositWithdrawalAmounts: withdrawalAmounts
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
        emit SlashingWithdrawalQueued(withdrawalRoot, withdrawal, queuedWithdrawalParams[0].depositShares);
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
    using SingleItemArrayLib for *;
    using SlashingLib for *;

    function test_Revert_WhenExitWithdrawalQueuePaused() public {
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_EXIT_WITHDRAWAL_QUEUE);
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            Withdrawal memory withdrawal,
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

        IERC20[][] memory tokensArray = new IERC20[][](1);
        tokensArray[0] = tokens;

        bool[] memory receiveAsTokens = new bool[](1);
        receiveAsTokens[0] = false;

        cheats.expectRevert(IPausable.CurrentlyPaused.selector);
        delegationManager.completeQueuedWithdrawals(tokensArray,  receiveAsTokens, 1);
    }

    function test_Revert_WhenInputArrayLengthMismatch() public {
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            Withdrawal memory withdrawal,
            IERC20[] memory tokens,
            /* bytes32 withdrawalRoot */
        ) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            depositAmount: 100,
            withdrawalAmount: 100
        });
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        // Roll to completion block
        cheats.roll(withdrawal.startBlock + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS());

        // resize tokens array
        IERC20[] memory newTokens = new IERC20[](0);

        cheats.prank(defaultStaker);
        cheats.expectRevert(InputArrayLengthMismatch.selector);
        delegationManager.completeQueuedWithdrawal(withdrawal, newTokens,  false);

        IERC20[][] memory tokensArray = new IERC20[][](1);
        tokensArray[0] = newTokens;

        bool[] memory receiveAsTokens = new bool[](1);
        receiveAsTokens[0] = true;

        cheats.prank(defaultStaker);
        cheats.expectRevert(InputArrayLengthMismatch.selector);
        delegationManager.completeQueuedWithdrawals(tokensArray,  receiveAsTokens, 1);

        // check that the withdrawal completes otherwise
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens,  true);
    }
    
    function test_Revert_WhenWithdrawerNotCaller(address invalidCaller) filterFuzzedAddressInputs(invalidCaller) public {
        cheats.assume(invalidCaller != defaultStaker);

        _registerOperatorWithBaseDetails(defaultOperator);
        (
            Withdrawal memory withdrawal,
            IERC20[] memory tokens,
        ) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            depositAmount: 100,
            withdrawalAmount: 100
        });
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        cheats.expectRevert(WithdrawerNotCaller.selector);
        cheats.prank(invalidCaller);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens,  false);
    }

    function test_Revert_WhenInvalidWithdrawalRoot() public {
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            Withdrawal memory withdrawal,
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

        cheats.expectRevert(WithdrawalNotQueued.selector);
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens,  false);
    }

    /**
     * @notice should revert if minWithdrawalDelayBlocks has not passed, and if
     * delegationManager.getCompletableTimestamp returns a value greater than minWithdrawalDelayBlocks
     * then it should revert if the validBlockNumber has not passed either.
     */
    function test_Revert_WhenWithdrawalDelayNotPassed(
        uint256[] memory fuzzAmounts,
        bool receiveAsTokens
    ) public {
        (
            uint256[] memory depositAmounts,
            uint256[] memory withdrawalAmounts
        ) = _fuzzDepositWithdrawalAmounts(fuzzAmounts);
        
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            Withdrawal memory withdrawal,
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
        cheats.expectRevert(WithdrawalDelayNotElapsed.selector);
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, receiveAsTokens);

        IERC20[][] memory tokensArray = new IERC20[][](1);
        tokensArray[0] = tokens;

        bool[] memory receiveAsTokensArray = new bool[](1);
        receiveAsTokensArray[0] = false;

        cheats.expectRevert(WithdrawalDelayNotElapsed.selector);
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawals(tokensArray,  receiveAsTokensArray, 1);
    }

    /**
     * Test completing multiple queued withdrawals for a single strategy by passing in the withdrawals
     */
    function test_completeQueuedWithdrawals_MultipleWithdrawals(
        address staker,
        uint64 depositAmount,
        uint numWithdrawals
    ) public filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != defaultOperator);
        cheats.assume(depositAmount > 0);
        cheats.assume(numWithdrawals > 1 && numWithdrawals < 20);

        (
            IDelegationManagerTypes.Withdrawal[] memory withdrawals,
            IERC20[][] memory tokens,
            bytes32[] memory withdrawalRoots
        ) = _setUpCompleteQueuedWithdrawalsSingleStrat({
            staker: staker,
            withdrawer: staker,
            depositAmount: depositAmount,
            numWithdrawals: numWithdrawals
        });

        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, withdrawals[0].strategies[0]);

        for (uint i = 0; i < withdrawalRoots.length; i++) {
            assertTrue(delegationManager.pendingWithdrawals(withdrawalRoots[i]), "withdrawalRoots should be pending");
        }

        bool[] memory receiveAsTokens = new bool[](withdrawals.length);
        for (uint i = 0; i < withdrawals.length; i++) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit SlashingWithdrawalCompleted(withdrawalRoots[i]);
            receiveAsTokens[i] = true;
        }

        // completeQueuedWithdrawal
        cheats.roll(withdrawals[0].startBlock + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS());
        cheats.prank(staker);
        delegationManager.completeQueuedWithdrawals(withdrawals, tokens, receiveAsTokens);

        uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, withdrawals[0].strategies[0]);
        assertEq(operatorSharesAfter, operatorSharesBefore, "operator shares should be unchanged");

        for (uint i = 0; i < withdrawals.length; i++) {
            assertFalse(delegationManager.pendingWithdrawals(withdrawalRoots[i]), "withdrawalRoot should be completed and marked false now");
        }
    }

    /**
     * Test completing multiple queued withdrawals for a single strategy without passing in the withdrawals
     */
    function test_completeQueuedWithdrawals_NumToComplete(
        address staker,
        uint64 depositAmount,
        uint numWithdrawals,
        uint numToComplete
    ) public filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != defaultOperator);
        cheats.assume(depositAmount > 0);
        numWithdrawals = bound(numWithdrawals, 2, 20);
        numToComplete = bound(numToComplete, 1, numWithdrawals);

        (
            IDelegationManagerTypes.Withdrawal[] memory withdrawals,
            IERC20[][] memory tokens,
            bytes32[] memory withdrawalRoots
        ) = _setUpCompleteQueuedWithdrawalsSingleStrat({
            staker: staker,
            withdrawer: staker,
            depositAmount: depositAmount,
            numWithdrawals: numWithdrawals
        });

        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        uint256 operatorSharesBefore = delegationManager.operatorShares(defaultOperator, withdrawals[0].strategies[0]);

        for (uint i = 0; i < withdrawalRoots.length; i++) {
            assertTrue(delegationManager.pendingWithdrawals(withdrawalRoots[i]), "withdrawalRoots should be pending");
        }

        bool[] memory receiveAsTokens = new bool[](withdrawals.length);
        for (uint i = 0; i < numToComplete; i++) {
            // cheats.expectEmit(true, true, true, true, address(delegationManager));
            // emit SlashingWithdrawalCompleted(withdrawalRoots[i]);
            receiveAsTokens[i] = true;
        }

        // completeQueuedWithdrawal
        cheats.roll(withdrawals[0].startBlock + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS());
        cheats.prank(staker);
        delegationManager.completeQueuedWithdrawals(tokens, receiveAsTokens, numToComplete);

        uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, withdrawals[0].strategies[0]);
        assertEq(operatorSharesAfter, operatorSharesBefore, "operator shares should be unchanged");

        for (uint i = 0; i < numToComplete; i++) {
            assertFalse(delegationManager.pendingWithdrawals(withdrawalRoots[i]), "withdrawalRoot should be completed and marked false now");
        }

        for (uint i = numToComplete; i < numWithdrawals; i++) {
            assertTrue(delegationManager.pendingWithdrawals(withdrawalRoots[i]), "withdrawalRoot should still be pending");
        }
    }

    /**
     * @notice Verifies that `DelegationManager.completeQueuedWithdrawal` properly completes a queued withdrawal for the `withdrawer`
     * for a single strategy. Withdraws as tokens so there are no operator shares increase.
     * - Asserts that the withdrawalRoot is True before `completeQueuedWithdrawal` and False after
     * - Asserts operatorShares is unchanged after `completeQueuedWithdrawal`
     * - Checks that event `WithdrawalCompleted` is emitted with withdrawalRoot
     */
    function test_completeQueuedWithdrawal_SingleStratWithdrawAsTokens(
        Randomness r
    ) public rand(r) {
        address staker = r.Address();
        uint128 depositAmount = r.Uint128();
        uint128 withdrawalAmount = r.Uint128(1, depositAmount);

        _registerOperatorWithBaseDetails(defaultOperator);
        (
            Withdrawal memory withdrawal,
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
        Randomness r
    ) public rand(r) {
        uint128 depositAmount = r.Uint128();
        uint128 withdrawalAmount = r.Uint128(1, depositAmount);

        // Deposit Staker
        uint256[] memory depositAmounts = depositAmount.toArrayU256();
        IStrategy[] memory strategies = strategyMock.toArray();
        strategyManagerMock.setDeposits(defaultStaker, strategies, depositAmounts);

        // Register operator and delegate to it
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        uint256 operatorSharesBeforeQueue = delegationManager.operatorShares(defaultOperator, strategyMock);

        // Queue withdrawal
        (
            QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategyMock,
            depositSharesToWithdraw: withdrawalAmount
        });
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
        assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");
        uint256 operatorSharesAfterQueue = delegationManager.operatorShares(defaultOperator, strategyMock);

        assertEq(operatorSharesAfterQueue, operatorSharesBeforeQueue - withdrawalAmount, "operator shares should be decreased after queue");

        // Slash operator while staker has queued withdrawal
        (
            uint256 sharesToDecrement,
            uint64 newOperatorMagnitude,
        ) = _calculateSlashedSharesAndMagnitude(depositAmount - withdrawalAmount, WAD, wadToSlash);
        _setOperatorMagnitude(defaultOperator, strategyMock, newOperatorMagnitude);
        cheats.prank(address(allocationManagerMock));
        delegationManager.burnOperatorShares(defaultOperator, withdrawal.strategies[0], WAD, operatorMagnitude);
        uint256 operatorSharesAfterSlash = delegationManager.operatorShares(defaultOperator, strategyMock);
        assertApproxEqAbs(operatorSharesAfterSlash, operatorSharesAfterQueue / 2, 1, "operator shares should be decreased after slash");

        // Complete queue withdrawal
        {
            IERC20[] memory tokens = new IERC20[](1);
            tokens[0] = IERC20(strategyMock.underlyingToken());
            cheats.roll(withdrawal.startBlock + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS());
            cheats.prank(defaultStaker);
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit SlashingWithdrawalCompleted(withdrawalRoot);
            delegationManager.completeQueuedWithdrawal(withdrawal, tokens,  true);    
        }

        // Checks: operator shares
        uint256 operatorSharesAfterWithdrawalComplete = delegationManager.operatorShares(defaultOperator, strategyMock);
        assertEq(operatorSharesAfterWithdrawalComplete, operatorSharesAfterSlash, "operator shares should be unchanged from slash to withdrawal completion");    
        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be completed and marked false now");

        // Checks: staker shares: 
        uint256 stakerSharesWithdrawn = strategyManagerMock.strategySharesWithdrawn(defaultStaker, strategyMock);
        {
            (uint256 depositScalingFactor,,) = delegationManager.stakerScalingFactor(defaultStaker, strategyMock);
            ssf = StakerScalingFactors({
                depositScalingFactor: depositScalingFactor,
                isBeaconChainScalingFactorSet: false,
                beaconChainScalingFactor: 0
            });    
        }

        uint256 actualSharesWithdrawn = withdrawalAmount.toShares(ssf, newOperatorMagnitude);
        assertEq(stakerSharesWithdrawn, actualSharesWithdrawn, "staker shares withdrawn not calculated correctly");
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
        Randomness r
    ) public rand(r) {
        uint128 depositAmount = r.Uint128();
        uint128 withdrawalAmount = r.Uint128(1, depositAmount);

        // Deposit Staker
        eigenPodManagerMock.setPodOwnerShares(defaultStaker, int256(uint256(depositAmount)));

        // Register operator and delegate to it
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        uint256 operatorSharesBeforeQueue = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);

        // Queue withdrawal
        (
            QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: beaconChainETHStrategy,
            depositSharesToWithdraw: withdrawalAmount
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
        Randomness r
    ) public rand(r) {
        uint128 depositAmount = r.Uint128();
        uint128 withdrawalAmount = r.Uint128(1, depositAmount);

        // Deposit Staker
        eigenPodManagerMock.setPodOwnerShares(defaultStaker, int256(uint256(depositAmount)));

        // Register operator and delegate to it
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);
        uint256 operatorSharesBeforeQueue = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);

        // Queue withdrawal
        (
            QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: beaconChainETHStrategy,
            depositSharesToWithdraw: withdrawalAmount
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
            delegationManager.burnOperatorShares(defaultOperator, withdrawal.strategies[0], WAD, operatorMagnitude);
            operatorSharesAfterAVSSlash = delegationManager.operatorShares(defaultOperator, beaconChainETHStrategy);
            assertApproxEqAbs(operatorSharesAfterAVSSlash, operatorSharesAfterBeaconSlash / 2, 1, "operator shares should be decreased after AVS slash");
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
    function testFuzz_completeQueuedWithdrawal_SingleStratWithdrawAsShares_nonSlashedOperator(
        Randomness r
    ) public rand(r) {
        address staker = r.Address();
        uint128 depositAmount = r.Uint128();
        uint128 withdrawalAmount = r.Uint128(1, depositAmount);
        
        _registerOperatorWithBaseDetails(defaultOperator);

        (
            Withdrawal memory withdrawal,
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
}

contract DelegationManagerUnitTests_burningShares is DelegationManagerUnitTests {
    using SingleItemArrayLib for *;

    /**
     * @notice Test burning shares for an operator with no queued withdrawals
     * - Asserts slashable shares before and after in queue is 0
     * - Asserts operator shares are decreased by half
     */
    function testFuzz_burnOperatorShares_NoQueuedWithdrawals(Randomness r) public {
        address operator = r.Address();
        address staker = r.Address();
        uint64 initMagnitude = WAD;
        uint64 newMagnitude = 5e17;
        uint256 shares = r.Uint256(1, MAX_STRATEGY_SHARES);

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(operator);
        _setOperatorMagnitude(operator, strategyMock, initMagnitude);
        // Set the staker deposits in the strategies
        IStrategy[] memory strategyArray = strategyMock.toArray();
        uint256[] memory sharesArray = shares.toArrayU256();
        strategyManagerMock.setDeposits(staker, strategyArray, sharesArray);
        // delegate from the `staker` to the operator
        _delegateToOperatorWhoAcceptsAllStakers(staker, operator);
        uint256 operatorSharesBefore = delegationManager.operatorShares(operator, strategyMock);
        uint256 queuedSlashableSharesBefore = delegationManager.getSlashableSharesInQueue(operator, strategyMock);

        // calculate burned shares, should be halved
        uint256 sharesToBurn = shares/2;

        // Burn shares
        cheats.prank(address(allocationManagerMock));
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(operator, address(0), strategyMock, sharesToBurn);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesBurned(operator, strategyMock, sharesToBurn);
        delegationManager.burnOperatorShares({
            operator: operator,
            strategy: strategyMock,
            prevMaxMagnitude: initMagnitude,
            newMaxMagnitude: newMagnitude
        });

        uint256 queuedSlashableSharesAfter = delegationManager.getSlashableSharesInQueue(operator, strategyMock);
        uint256 operatorSharesAfter = delegationManager.operatorShares(operator, strategyMock);
        assertEq(queuedSlashableSharesBefore, 0, "there should be no slashable shares in queue");
        assertEq(queuedSlashableSharesAfter, 0, "there should be no slashable shares in queue");
        assertEq(operatorSharesAfter, operatorSharesBefore - sharesToBurn, "operator shares should be decreased by sharesToBurn");
    }

    /**
     * @notice Test burning shares for an operator with no slashable queued withdrawals in past MIN_WITHDRAWAL_DELAY_BLOCKS window.
     * There does exist past queued withdrawals but nothing in the queue is slashable.
     * - Asserts slashable shares in queue right after queuing a withdrawal is the withdrawal amount
     * and then checks that after the withdrawal window the slashable shares is 0 again.
     * - Asserts operator shares are decreased by half after burning
     * - Asserts that the slashable shares in queue before/after burning are 0
     */
    function testFuzz_burnOperatorShares_NoQueuedWithdrawalsInWindow(Randomness r) public {
        // 1. Randomize operator and staker info
        // Operator info
        address operator = r.Address();
        uint64 newMagnitude = 5e17;
        // First staker
        address staker1 = r.Address();
        uint256 shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        // Second Staker, will queue withdraw shares
        address staker2 = r.Address();
        uint256 depositAmount = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint256 withdrawAmount = r.Uint256(1, depositAmount);

        // 2. Register the operator, set the staker deposits, and delegate the 2 stakers to them
        _registerOperatorWithBaseDetails(operator);
        {
            // Set the first staker deposits in the strategies
            IStrategy[] memory strategyArray = strategyMock.toArray();
            uint256[] memory sharesArray = shares.toArrayU256();
            uint256[] memory depositArray = depositAmount.toArrayU256();
            strategyManagerMock.setDeposits(staker1, strategyArray, sharesArray);
            // Set the second staker's deposits in the strategies
            strategyManagerMock.setDeposits(staker2, strategyArray, depositArray);
        }
        _delegateToOperatorWhoAcceptsAllStakers(staker1, operator);
        _delegateToOperatorWhoAcceptsAllStakers(staker2, operator);

        // 3. Queue withdrawal for staker2 and roll blocks forward so that the withdrawal is not slashable
        {
            (
                QueuedWithdrawalParams[] memory queuedWithdrawalParams,
                Withdrawal memory withdrawal,
            ) = _setUpQueueWithdrawalsSingleStrat({
                staker: staker2,
                withdrawer: staker2,
                strategy: strategyMock,
                depositSharesToWithdraw: withdrawAmount
            });
            cheats.prank(staker2);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                withdrawAmount,
                "there should be withdrawAmount slashable shares in queue"
            );
            cheats.roll(withdrawal.startBlock + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS());
        }

        uint256 operatorSharesBefore = delegationManager.operatorShares(operator, strategyMock);
        uint256 queuedSlashableSharesBefore = delegationManager.getSlashableSharesInQueue(operator, strategyMock);

        // calculate burned shares, should be halved
        // staker2 queue withdraws shares and we roll blocks to after the withdrawal is no longer slashable.
        // Therefore amount of shares to burn should be what the staker still has remaining + staker1 shares and then
        // divided by 2 since the operator was slashed 50%
        uint256 sharesToBurn = (shares + depositAmount - withdrawAmount) / 2;

        // 4. Burn shares
        _setOperatorMagnitude(operator, strategyMock, newMagnitude);
        cheats.prank(address(allocationManagerMock));
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(operator, address(0), strategyMock, sharesToBurn);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesBurned(operator, strategyMock, sharesToBurn);
        delegationManager.burnOperatorShares({
            operator: operator,
            strategy: strategyMock,
            prevMaxMagnitude: WAD,
            newMaxMagnitude: newMagnitude
        });

        // 5. Assert expected values
        uint256 queuedSlashableSharesAfter = delegationManager.getSlashableSharesInQueue(operator, strategyMock);
        uint256 operatorSharesAfter = delegationManager.operatorShares(operator, strategyMock);
        assertEq(queuedSlashableSharesBefore, 0, "there should be no slashable shares in queue");
        assertEq(queuedSlashableSharesAfter, 0, "there should be no slashable shares in queue");
        assertEq(operatorSharesAfter, operatorSharesBefore - sharesToBurn, "operator shares should be decreased by sharesToBurn");
    }

    /**
     * @notice Test burning shares for an operator with slashable queued withdrawals in past MIN_WITHDRAWAL_DELAY_BLOCKS window.
     * There exists a single withdrawal that is slashable.
     */
    function testFuzz_burnOperatorShares_SingleSlashableWithdrawal(Randomness r) public {
        // 1. Randomize operator and staker info
        // Operator info
        address operator = r.Address();
        uint64 newMagnitude = 25e16;
        // First staker
        address staker1 = r.Address();
        uint256 shares = r.Uint256(1, MAX_STRATEGY_SHARES);
        // Second Staker, will queue withdraw shares
        address staker2 = r.Address();
        uint256 depositAmount = r.Uint256(1, MAX_STRATEGY_SHARES);
        uint256 withdrawAmount = r.Uint256(1, depositAmount);

        // 2. Register the operator, set the staker deposits, and delegate the 2 stakers to them
        _registerOperatorWithBaseDetails(operator);
        {
            // Set the first staker deposits in the strategies
            IStrategy[] memory strategyArray = strategyMock.toArray();
            uint256[] memory sharesArray = shares.toArrayU256();
            uint256[] memory depositArray = depositAmount.toArrayU256();
            strategyManagerMock.setDeposits(staker1, strategyArray, sharesArray);
            // Set the second staker's deposits in the strategies
            strategyManagerMock.setDeposits(staker2, strategyArray, depositArray);
        }
        _delegateToOperatorWhoAcceptsAllStakers(staker1, operator);
        _delegateToOperatorWhoAcceptsAllStakers(staker2, operator);

        // 3. Queue withdrawal for staker2 so that the withdrawal is slashable
        {
            (
                QueuedWithdrawalParams[] memory queuedWithdrawalParams,
                Withdrawal memory withdrawal,
            ) = _setUpQueueWithdrawalsSingleStrat({
                staker: staker2,
                withdrawer: staker2,
                strategy: strategyMock,
                depositSharesToWithdraw: withdrawAmount
            });
            cheats.prank(staker2);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                withdrawAmount,
                "there should be withdrawAmount slashable shares in queue"
            );
        }

        uint256 operatorSharesBefore = delegationManager.operatorShares(operator, strategyMock);
        uint256 queuedSlashableSharesBefore = delegationManager.getSlashableSharesInQueue(operator, strategyMock);

        // calculate burned shares, should be 3/4 of the original shares
        // staker2 queue withdraws shares
        // Therefore amount of shares to burn should be what the staker still has remaining + staker1 shares and then
        // divided by 2 since the operator was slashed 50%
        uint256 sharesToDecrease = (shares + depositAmount - withdrawAmount) * 3 / 4;
        uint256 sharesToBurn = sharesToDecrease + withdrawAmount * 3 / 4;

        // 4. Burn shares
        _setOperatorMagnitude(operator, strategyMock, newMagnitude);
        cheats.prank(address(allocationManagerMock));
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(operator, address(0), strategyMock, sharesToDecrease);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesBurned(operator, strategyMock, sharesToBurn);
        delegationManager.burnOperatorShares({
            operator: operator,
            strategy: strategyMock,
            prevMaxMagnitude: WAD,
            newMaxMagnitude: newMagnitude
        });

        // 5. Assert expected values
        uint256 queuedSlashableSharesAfter = delegationManager.getSlashableSharesInQueue(operator, strategyMock);
        uint256 operatorSharesAfter = delegationManager.operatorShares(operator, strategyMock);
        assertEq(queuedSlashableSharesBefore, withdrawAmount, "Slashable shares in queue should be full withdraw amount");
        assertEq(queuedSlashableSharesAfter, withdrawAmount / 4, "Slashable shares in queue should be 1/4 withdraw amount after slashing");
        assertEq(operatorSharesAfter, operatorSharesBefore - sharesToDecrease, "operator shares should be decreased by sharesToBurn");
    }

    /**
     * @notice Test burning shares for an operator with slashable queued withdrawals in past MIN_WITHDRAWAL_DELAY_BLOCKS window.
     * There exists multiple withdrawals that are slashable.
     */
    function testFuzz_burnOperatorShares_MultipleSlashableWithdrawals(Randomness r) public {
        // 1. Randomize operator and staker info
        // Operator info
        address operator = r.Address();
        uint64 newMagnitude = 25e16;
        // Staker and withdrawing amounts
        address staker = r.Address();
        uint256 depositAmount = r.Uint256(3, MAX_STRATEGY_SHARES);
        uint256 withdrawAmount1 = r.Uint256(2, depositAmount);
        uint256 withdrawAmount2 = r.Uint256(1, depositAmount - withdrawAmount1);

        // 2. Register the operator, set the staker deposits, and delegate the 2 stakers to them
        _registerOperatorWithBaseDetails(operator);
        {
            // Set the first staker deposits in the strategies
            IStrategy[] memory strategyArray = strategyMock.toArray();
            uint256[] memory sharesArray = depositAmount.toArrayU256();
            strategyManagerMock.setDeposits(staker, strategyArray, sharesArray);
        }
        _delegateToOperatorWhoAcceptsAllStakers(staker, operator);

        // 3. Queue withdrawal for staker and roll blocks forward so that the withdrawal is not slashable
        {
            (
                QueuedWithdrawalParams[] memory queuedWithdrawalParams,
                Withdrawal memory withdrawal,
            ) = _setUpQueueWithdrawalsSingleStrat({
                staker: staker,
                withdrawer: staker,
                strategy: strategyMock,
                depositSharesToWithdraw: withdrawAmount1
            });
            cheats.prank(staker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                withdrawAmount1,
                "there should be withdrawAmount slashable shares in queue"
            );

            (
                queuedWithdrawalParams,
                withdrawal,
            ) = _setUpQueueWithdrawalsSingleStrat({
                staker: staker,
                withdrawer: staker,
                strategy: strategyMock,
                depositSharesToWithdraw: withdrawAmount2
            });
            cheats.prank(staker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                withdrawAmount2 + withdrawAmount1,
                "there should be withdrawAmount slashable shares in queue"
            );
        }

        uint256 operatorSharesBefore = delegationManager.operatorShares(operator, strategyMock);
        uint256 queuedSlashableSharesBefore = delegationManager.getSlashableSharesInQueue(operator, strategyMock);

        // calculate burned shares, should be halved for both operatorShares and slashable shares in queue
        // staker queue withdraws shares twice and both withdrawals should be slashed 75%.
        uint256 sharesToDecrease = (depositAmount - withdrawAmount1 - withdrawAmount2) * 3 / 4;
        uint256 sharesToBurn = sharesToDecrease + (delegationManager.getSlashableSharesInQueue(operator, strategyMock) * 3 / 4);

        // 4. Burn shares
        _setOperatorMagnitude(operator, strategyMock, newMagnitude);
        cheats.prank(address(allocationManagerMock));
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(operator, address(0), strategyMock, sharesToDecrease);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesBurned(operator, strategyMock, sharesToBurn);
        delegationManager.burnOperatorShares({
            operator: operator,
            strategy: strategyMock,
            prevMaxMagnitude: WAD,
            newMaxMagnitude: newMagnitude
        });

        // 5. Assert expected values
        uint256 queuedSlashableSharesAfter = delegationManager.getSlashableSharesInQueue(operator, strategyMock);
        uint256 operatorSharesAfter = delegationManager.operatorShares(operator, strategyMock);
        assertEq(queuedSlashableSharesBefore, (withdrawAmount1 + withdrawAmount2), "Slashable shares in queue should be full withdraw amount");
        assertEq(queuedSlashableSharesAfter, (withdrawAmount1 + withdrawAmount2) / 4, "Slashable shares in queue should be 1/4 withdraw amount after slashing");
        assertEq(operatorSharesAfter, operatorSharesBefore - sharesToDecrease, "operator shares should be decreased by sharesToBurn");
    }

    /**
     * @notice TODO Test burning shares for an operator with slashable queued withdrawals in past MIN_WITHDRAWAL_DELAY_BLOCKS window.
     * There exists multiple withdrawals that are slashable but queued with different maxMagnitudes at
     * time of queuing.
     *
     * Test Setup:  
     * - staker1 deposits, queues withdrawal for some amount,
     * - operator slashed 50% 
     * - staker 2 deposits, queues withdrawal for some amount
     * - operator is then slashed another 50%
     * slashed amount for staker 1 should be 75% and staker 2 should be 50% where the total
     * slashed amount is the sum of both
     */
    function testFuzz_burnOperatorShares_MultipleWithdrawalsMultipleSlashings(Randomness r) public {
        address operator = r.Address();
        address staker = r.Address();
        uint256 depositAmount = r.Uint256(3, MAX_STRATEGY_SHARES);
        uint256 depositSharesToWithdraw1 = r.Uint256(1, depositAmount);
        uint256 depositSharesToWithdraw2 = r.Uint256(1, depositAmount - depositSharesToWithdraw1);

        uint64 newMagnitude = 5e17;

        // 2. Register the operator, set the staker deposits, and delegate the 2 stakers to them
        _registerOperatorWithBaseDetails(operator);
        {
            // Set the first staker deposits in the strategies
            IStrategy[] memory strategyArray = strategyMock.toArray();
            uint256[] memory depositArray = depositAmount.toArrayU256();
            strategyManagerMock.setDeposits(staker, strategyArray, depositArray);
        }
        _delegateToOperatorWhoAcceptsAllStakers(staker, operator);

        // 3. Queue withdrawal for staker and slash operator for 50%
        {
            (
                QueuedWithdrawalParams[] memory queuedWithdrawalParams,
                Withdrawal memory withdrawal,
            ) = _setUpQueueWithdrawalsSingleStrat({
                staker: staker,
                withdrawer: staker,
                strategy: strategyMock,
                depositSharesToWithdraw: depositSharesToWithdraw1
            });

            // 3.1 queue a withdrawal for the staker
            cheats.prank(staker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            uint256 operatorSharesBefore = delegationManager.operatorShares(operator, strategyMock);
            uint256 queuedSlashableSharesBefore = delegationManager.getSlashableSharesInQueue(operator, strategyMock);

            uint256 sharesToDecrease = (depositAmount - depositSharesToWithdraw1) / 2;
            uint256 sharesToBurn = sharesToDecrease + depositSharesToWithdraw1/2;

            // 3.2 Burn shares
            _setOperatorMagnitude(operator, strategyMock, newMagnitude);
            cheats.prank(address(allocationManagerMock));
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesDecreased(operator, address(0), strategyMock, sharesToDecrease);
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesBurned(operator, strategyMock, sharesToBurn);
            delegationManager.burnOperatorShares({
                operator: operator,
                strategy: strategyMock,
                prevMaxMagnitude: WAD,
                newMaxMagnitude: newMagnitude
            });

            // 3.3 Assert slashable shares and operator shares
            assertEq(
                queuedSlashableSharesBefore,
                depositSharesToWithdraw1,
                "Slashable shares in queue should be full withdraw1 amount"
            );
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                depositSharesToWithdraw1 / 2,
                "Slashable shares in queue should be 1/2 withdraw1 amount after slashing"
            );
            assertEq(
                delegationManager.operatorShares(operator, strategyMock),
                operatorSharesBefore - sharesToDecrease,
                "operator shares should be decreased by sharesToBurn"
            );
        }

        // 4. Queue withdrawal for staker and slash operator for 50% again
        newMagnitude = newMagnitude/2;
        {
            (
                QueuedWithdrawalParams[] memory queuedWithdrawalParams,
                Withdrawal memory withdrawal,
            ) = _setUpQueueWithdrawalsSingleStrat({
                staker: staker,
                withdrawer: staker,
                strategy: strategyMock,
                depositSharesToWithdraw: depositSharesToWithdraw2
            });

            // actual withdrawn shares are half of the deposit shares because of first slashing
            uint256 withdrawAmount2 = depositSharesToWithdraw2 / 2;

            // 4.1 queue a withdrawal for the staker
            cheats.prank(staker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            uint256 operatorSharesBefore = delegationManager.operatorShares(operator, strategyMock);
            uint256 queuedSlashableSharesBefore = delegationManager.getSlashableSharesInQueue(operator, strategyMock);

            uint256 sharesToDecrease = operatorSharesBefore / 2;
            uint256 sharesToBurn = sharesToDecrease + (withdrawAmount2 + depositSharesToWithdraw1/2)/2;

            // 4.2 Burn shares
            _setOperatorMagnitude(operator, strategyMock, newMagnitude);
            cheats.prank(address(allocationManagerMock));
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesDecreased(operator, address(0), strategyMock, sharesToDecrease);
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesBurned(operator, strategyMock, sharesToBurn);
            delegationManager.burnOperatorShares({
                operator: operator,
                strategy: strategyMock,
                prevMaxMagnitude: newMagnitude*2,
                newMaxMagnitude: newMagnitude
            });

            // 4.3 Assert slashable shares and operator shares
            assertEq(
                queuedSlashableSharesBefore,
                withdrawAmount2 + depositSharesToWithdraw1/2,
                "Slashable shares in queue before should be withdrawAmount1 / 2 + withdrawAmount2"
            );
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                (withdrawAmount2 + depositSharesToWithdraw1/2)/2,
                "Slashable shares in queue should be (withdrawAmount2 + depositSharesToWithdraw1/2)/2 after slashing"
            );
            assertEq(
                delegationManager.operatorShares(operator, strategyMock),
                operatorSharesBefore - sharesToDecrease,
                "operator shares should be decreased by sharesToBurn"
            );
        }

    }

    /**
     * @notice Ensure that when a withdrawal is completable then there are no slashable shares in the queue.
     * However if the withdrawal is not completable and the withdrawal delay hasn't elapsed, then the withdrawal
     * should be counted as slashable.
     */
    function testFuzz_burnOperatorShares_Timings(Randomness r) public {
        // 1. Randomize operator and staker info
        // Operator info
        address operator = r.Address();
        uint64 newMagnitude = 25e16;
        // staker
        address staker = r.Address();
        uint256 depositAmount = r.Uint256(1, MAX_STRATEGY_SHARES);

        // 2. Register the operator, set the staker deposits, and delegate the staker to them
        _registerOperatorWithBaseDetails(operator);
        {
            // Set the first staker deposits in the strategies
            IStrategy[] memory strategyArray = strategyMock.toArray();
            uint256[] memory depositArray = depositAmount.toArrayU256();
            strategyManagerMock.setDeposits(staker, strategyArray, depositArray);
        }
        _delegateToOperatorWhoAcceptsAllStakers(staker, operator);

        // 3. Queue withdrawal for staker and roll blocks forward so that the withdrawal is completable
        uint256 completableBlock;
        {
            (
                QueuedWithdrawalParams[] memory queuedWithdrawalParams,
                Withdrawal memory withdrawal,
            ) = _setUpQueueWithdrawalsSingleStrat({
                staker: staker,
                withdrawer: staker,
                strategy: strategyMock,
                depositSharesToWithdraw: depositAmount
            });
            cheats.startPrank(staker);
            delegationManager.queueWithdrawals(queuedWithdrawalParams);
            // 3.1 after queuing the withdrawal, check that there are slashable shares in queue
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                depositAmount,
                "there should be depositAmount slashable shares in queue"
            );
            // Check slashable shares in queue before and when the withdrawal is completable
            completableBlock = withdrawal.startBlock + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS();
            IERC20[] memory tokenArray = strategyMock.underlyingToken().toArray();

            // 3.2 roll to right before withdrawal is completable, check that slashable shares are still there
            // attempting to complete a withdrawal should revert
            cheats.roll(completableBlock - 1);
            cheats.expectRevert(WithdrawalDelayNotElapsed.selector);
            delegationManager.completeQueuedWithdrawal(withdrawal, tokenArray, true);
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                depositAmount,
                "there should still be depositAmount slashable shares in queue"
            );

            // 3.3 roll to blocknumber that the withdrawal is completable, there should be no slashable shares in queue
            cheats.roll(completableBlock);
            delegationManager.completeQueuedWithdrawal(withdrawal, tokenArray, true);
            assertEq(
                delegationManager.getSlashableSharesInQueue(operator, strategyMock),
                0,
                "there should be no slashable shares in queue when the withdrawal is completable"
            );

            cheats.stopPrank();

        }

        uint256 operatorSharesBefore = delegationManager.operatorShares(operator, strategyMock);

        // 4. Burn 0 shares when new magnitude is set
        _setOperatorMagnitude(operator, strategyMock, newMagnitude);
        cheats.prank(address(allocationManagerMock));
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesDecreased(operator, address(0), strategyMock, 0);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesBurned(operator, strategyMock, 0);
        delegationManager.burnOperatorShares({
            operator: operator,
            strategy: strategyMock,
            prevMaxMagnitude: WAD,
            newMaxMagnitude: newMagnitude
        });

        // 5. Assert expected values
        uint256 operatorSharesAfter = delegationManager.operatorShares(operator, strategyMock);
        assertEq(
            delegationManager.getSlashableSharesInQueue(operator, strategyMock),
            0,
            "there should still be no slashable shares in queue after burning 0 shares"
        );
        assertEq(operatorSharesAfter, operatorSharesBefore, "operator shares should be unchanged and equal to 0");
        assertEq(operatorSharesBefore, 0, "operator shares should be unchanged and equal to 0");
    }
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
   9. Invariant check getWithdrawableShares = sum(deposits), Multiple deposits with operator who has never been slashed
   10. Invariant check getWithdrawableShares = sum(deposits), Multiple deposits with operator who HAS been been slashed
 */

contract DelegationManagerUnitTests_Lifecycle is DelegationManagerUnitTests {
    using SingleItemArrayLib for *;

    // 2. RegisterOperator, Deposit, Delegate, Queue, Complete
    function test_register_operator_deposit_delegate_queue_complete(Randomness r) public rand(r) {
        address operator = r.Address();
        address staker = r.Address();
        IStrategy[] memory strategies = strategyMock.toArray();
        uint256[] memory depositShares = uint256(100 ether).toArrayU256();

        // 1) Register operator.
        _registerOperatorWithBaseDetails(operator);

        // 2) Mock deposit into SM.
        strategyManagerMock.setDeposits(staker, strategies, depositShares);
        
        // 3) Staker delegates to operator.
        _delegateToOperatorWhoAcceptsAllStakers(staker, operator);
        
        // 3) Staker queues withdrawals.
        QueuedWithdrawalParams[] memory queuedWithdrawalParams = new QueuedWithdrawalParams[](1);
        queuedWithdrawalParams[0] = QueuedWithdrawalParams({
            strategies: strategies,
            depositShares: depositShares,
            withdrawer: staker
        });
        cheats.prank(staker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        // 4) Complete queued withdrawals.
        Withdrawal memory withdrawal = Withdrawal({
            staker: staker,
            delegatedTo: operator,
            withdrawer: staker,
            nonce: 0,
            startBlock: uint32(block.number),
            strategies: strategies,
            scaledShares: depositShares
        });

        bytes32 withdrawalRoot = keccak256(abi.encode(withdrawal));
        assertTrue(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be pending");
        cheats.roll(block.number + delegationManager.MIN_WITHDRAWAL_DELAY_BLOCKS());

        cheats.prank(staker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokenMock.toArray(),  false);

        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should not be pending");

        // Checks
        assertEq(delegationManager.cumulativeWithdrawalsQueued(staker),  1, "staker nonce should have incremented");
        assertEq(delegationManager.operatorShares(operator, strategies[0]), 100 ether, "operator shares should be 0 after withdrawal");
    }
}