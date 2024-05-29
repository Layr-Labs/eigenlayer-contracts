// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

import "src/contracts/core/DelegationManager.sol";
import "src/contracts/strategies/StrategyBase.sol";

import "src/test/events/IDelegationManagerEvents.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

/**
 * @notice Unit testing of the DelegationManager contract. Withdrawals are tightly coupled
 * with EigenPodManager and StrategyManager and are part of integration tests.
 * Contracts tested: DelegationManager
 * Contracts not mocked: StrategyBase, PauserRegistry
 */
contract DelegationManagerUnitTests is EigenLayerUnitTestSetup, IDelegationManagerEvents {
    // Contract under test
    DelegationManager delegationManager;
    DelegationManager delegationManagerImplementation;

    // Mocks
    StrategyBase strategyImplementation;
    StrategyBase strategyMock;
    IERC20 mockToken;
    uint256 mockTokenInitialSupply = 10e50;

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
    mapping(IStrategy => uint256) public delegatedSharesBefore;

    function setUp() public virtual override {
        // Setup
        EigenLayerUnitTestSetup.setUp();

        // Deploy DelegationManager implmentation and proxy
        initializeStrategiesToSetDelayBlocks = new IStrategy[](0);
        initializeWithdrawalDelayBlocks = new uint256[](0);
        delegationManagerImplementation = new DelegationManager(strategyManagerMock, slasherMock, eigenPodManagerMock);
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
        strategyImplementation = new StrategyBase(strategyManagerMock);
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
        addressIsExcludedFromFuzzedInputs[address(delegationManager)] = true;
        addressIsExcludedFromFuzzedInputs[defaultApprover] = true;
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
        delegationManager.setStrategyWithdrawalDelayBlocks(strategies, withdrawalDelayBlocks);
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
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            __deprecated_earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        _registerOperator(operator, operatorDetails, emptyStringForMetadataURI);
    }

    function _registerOperatorWithDelegationApprover(address operator) internal {
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            __deprecated_earningsReceiver: operator,
            delegationApprover: defaultApprover,
            stakerOptOutWindowBlocks: 0
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

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            __deprecated_earningsReceiver: operator,
            delegationApprover: address(wallet),
            stakerOptOutWindowBlocks: 0
        });
        _registerOperator(operator, operatorDetails, emptyStringForMetadataURI);

        return wallet;
    }

    function _registerOperator(
        address operator,
        IDelegationManager.OperatorDetails memory operatorDetails,
        string memory metadataURI
    ) internal filterFuzzedAddressInputs(operator) {
        _filterOperatorDetails(operator, operatorDetails);
        cheats.prank(operator);
        delegationManager.registerAsOperator(operatorDetails, metadataURI);
    }

    function _filterOperatorDetails(
        address operator,
        IDelegationManager.OperatorDetails memory operatorDetails
    ) internal view {
        // filter out zero address since people can't delegate to the zero address and operators are delegated to themselves
        cheats.assume(operator != address(0));
        // filter out disallowed stakerOptOutWindowBlocks values
        cheats.assume(operatorDetails.stakerOptOutWindowBlocks <= delegationManager.MAX_STAKER_OPT_OUT_WINDOW_BLOCKS());
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
        uint256 withdrawalAmount
    ) internal view returns (
        IDelegationManager.QueuedWithdrawalParams[] memory,
        IDelegationManager.Withdrawal memory,
        bytes32
    ) {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        strategyArray[0] = strategy;
        uint256[] memory withdrawalAmounts = new uint256[](1);
        withdrawalAmounts[0] = withdrawalAmount;

        IDelegationManager.QueuedWithdrawalParams[] memory queuedWithdrawalParams = new IDelegationManager.QueuedWithdrawalParams[](1);
        queuedWithdrawalParams[0] = IDelegationManager.QueuedWithdrawalParams({
            strategies: strategyArray,
            shares: withdrawalAmounts,
            withdrawer: withdrawer
        });

        IDelegationManager.Withdrawal memory withdrawal = IDelegationManager.Withdrawal({
            staker: staker,
            delegatedTo: delegationManager.delegatedTo(staker),
            withdrawer: withdrawer,
            nonce: delegationManager.cumulativeWithdrawalsQueued(staker),
            startBlock: uint32(block.number),
            strategies: strategyArray,
            shares: withdrawalAmounts
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
        IDelegationManager.QueuedWithdrawalParams[] memory,
        IDelegationManager.Withdrawal memory,
        bytes32
    ) {
        IDelegationManager.QueuedWithdrawalParams[] memory queuedWithdrawalParams = new IDelegationManager.QueuedWithdrawalParams[](1);
        queuedWithdrawalParams[0] = IDelegationManager.QueuedWithdrawalParams({
            strategies: strategies,
            shares: withdrawalAmounts,
            withdrawer: withdrawer
        });
        
        IDelegationManager.Withdrawal memory withdrawal = IDelegationManager.Withdrawal({
            staker: staker,
            delegatedTo: delegationManager.delegatedTo(staker),
            withdrawer: withdrawer,
            nonce: delegationManager.cumulativeWithdrawalsQueued(staker),
            startBlock: uint32(block.number),
            strategies: strategies,
            shares: withdrawalAmounts
        });
        bytes32 withdrawalRoot = delegationManager.calculateWithdrawalRoot(withdrawal);
        
        return (queuedWithdrawalParams, withdrawal, withdrawalRoot);
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
    ) internal returns (IDelegationManager.Withdrawal memory, IERC20[] memory, bytes32) {
        uint256[] memory depositAmounts = new uint256[](1);
        depositAmounts[0] = depositAmount;
        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(staker, depositAmounts);
        (
            IDelegationManager.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManager.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: staker,
            withdrawer: withdrawer,
            strategy: strategies[0],
            withdrawalAmount: withdrawalAmount
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
     * Deploy and deposit staker into a single strategy, then set up a queued withdrawal for the staker
     * Assumptions: 
     * - operator is already a registered operator.
     * - withdrawalAmount <= depositAmount
     */
    function _setUpCompleteQueuedWithdrawalBeaconStrat(
        address staker,
        address withdrawer,
        uint256 depositAmount,
        uint256 withdrawalAmount
    ) internal returns (IDelegationManager.Withdrawal memory, IERC20[] memory, bytes32) {
        uint256[] memory depositAmounts = new uint256[](1);
        depositAmounts[0] = depositAmount;
        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = beaconChainETHStrategy;
        (
            IDelegationManager.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManager.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: staker,
            withdrawer: withdrawer,
            strategy: strategies[0],
            withdrawalAmount: withdrawalAmount
        });

        cheats.prank(staker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
        // Set the current deposits to be the depositAmount - withdrawalAmount
        uint256[] memory currentAmounts = new uint256[](1);
        currentAmounts[0] = depositAmount - withdrawalAmount;
        strategyManagerMock.setDeposits(staker, strategies, currentAmounts);

        IERC20[] memory tokens;
        // tokens[0] = strategies[0].underlyingToken();
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
    ) internal returns (IDelegationManager.Withdrawal memory, IERC20[] memory, bytes32) {
        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(staker, depositAmounts);

        IERC20[] memory tokens = new IERC20[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            tokens[i] = strategies[i].underlyingToken();
        }

        (
            IDelegationManager.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManager.Withdrawal memory withdrawal,
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
}

contract DelegationManagerUnitTests_Initialization_Setters is DelegationManagerUnitTests {
    function test_initialization() public {
        assertEq(
            address(delegationManager.strategyManager()),
            address(strategyManagerMock),
            "constructor / initializer incorrect, strategyManager set wrong"
        );
        assertEq(
            address(delegationManager.slasher()),
            address(slasherMock),
            "constructor / initializer incorrect, slasher set wrong"
        );
        assertEq(
            address(delegationManager.pauserRegistry()),
            address(pauserRegistry),
            "constructor / initializer incorrect, pauserRegistry set wrong"
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
            0,
            0, // minWithdrawalDelayBlocks
            initializeStrategiesToSetDelayBlocks,
            initializeWithdrawalDelayBlocks
        );
    }

    function testFuzz_setMinWithdrawalDelayBlocks_revert_notOwner(
        address invalidCaller
    ) public filterFuzzedAddressInputs(invalidCaller) {
        cheats.assume(invalidCaller != delegationManager.owner());
        cheats.prank(invalidCaller);
        cheats.expectRevert("Ownable: caller is not the owner");
        delegationManager.setMinWithdrawalDelayBlocks(0);
    }

    function testFuzz_setMinWithdrawalDelayBlocks_revert_tooLarge(uint256 newMinWithdrawalDelayBlocks) external {
        // filter fuzzed inputs to disallowed amounts
        cheats.assume(newMinWithdrawalDelayBlocks > delegationManager.MAX_WITHDRAWAL_DELAY_BLOCKS());

        // attempt to set the `minWithdrawalDelayBlocks` variable
        cheats.expectRevert("DelegationManager._setMinWithdrawalDelayBlocks: _minWithdrawalDelayBlocks cannot be > MAX_WITHDRAWAL_DELAY_BLOCKS");
        delegationManager.setMinWithdrawalDelayBlocks(newMinWithdrawalDelayBlocks);
    }

    function testFuzz_initialize_Revert_WhenWithdrawalDelayBlocksTooLarge(
        uint256[] memory withdrawalDelayBlocks,
        uint256 invalidStrategyIndex
    ) public {
        // set withdrawalDelayBlocks to be too large
        cheats.assume(withdrawalDelayBlocks.length > 0);
        uint256 numStrats = withdrawalDelayBlocks.length;
        IStrategy[] memory strategiesToSetDelayBlocks = new IStrategy[](numStrats);
        for (uint256 i = 0; i < numStrats; i++) {
            strategiesToSetDelayBlocks[i] = IStrategy(address(uint160(uint256(keccak256(abi.encode(strategyMock, i))))));
        }

        // set at least one index to be too large for withdrawalDelayBlocks
        invalidStrategyIndex = invalidStrategyIndex % numStrats;
        withdrawalDelayBlocks[invalidStrategyIndex] = MAX_WITHDRAWAL_DELAY_BLOCKS + 1;

        // Deploy DelegationManager implmentation and proxy
        delegationManagerImplementation = new DelegationManager(strategyManagerMock, slasherMock, eigenPodManagerMock);
        cheats.expectRevert(
            "DelegationManager._setStrategyWithdrawalDelayBlocks: _withdrawalDelayBlocks cannot be > MAX_WITHDRAWAL_DELAY_BLOCKS"
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
                        strategiesToSetDelayBlocks,
                        withdrawalDelayBlocks
                    )
                )
            )
        );
    }
}

contract DelegationManagerUnitTests_RegisterModifyOperator is DelegationManagerUnitTests {
    function test_registerAsOperator_revert_paused() public {
        // set the pausing flag
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_NEW_DELEGATION);

        cheats.expectRevert("Pausable: index is paused");
        delegationManager.registerAsOperator(
            IDelegationManager.OperatorDetails({
                __deprecated_earningsReceiver: defaultOperator,
                delegationApprover: address(0),
                stakerOptOutWindowBlocks: 0
            }),
            emptyStringForMetadataURI
        );
    }

    // @notice Verifies that someone cannot successfully call `DelegationManager.registerAsOperator(operatorDetails)` again after registering for the first time
    function testFuzz_registerAsOperator_revert_cannotRegisterMultipleTimes(
        address operator,
        IDelegationManager.OperatorDetails memory operatorDetails
    ) public filterFuzzedAddressInputs(operator) {
        _filterOperatorDetails(operator, operatorDetails);

        // Register once
        cheats.startPrank(operator);
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);

        // Expect revert when register again
        cheats.expectRevert("DelegationManager.registerAsOperator: caller is already actively delegated");
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        cheats.stopPrank();
    }

    /**
     * @notice Verifies that an operator cannot register with `stakerOptOutWindowBlocks` set larger than `MAX_STAKER_OPT_OUT_WINDOW_BLOCKS`
     */
    function testFuzz_registerAsOperator_revert_optOutBlocksTooLarge(
        IDelegationManager.OperatorDetails memory operatorDetails
    ) public {
        // filter out *allowed* stakerOptOutWindowBlocks values
        cheats.assume(operatorDetails.stakerOptOutWindowBlocks > delegationManager.MAX_STAKER_OPT_OUT_WINDOW_BLOCKS());

        cheats.prank(defaultOperator);
        cheats.expectRevert("DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be > MAX_STAKER_OPT_OUT_WINDOW_BLOCKS");
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
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
        IDelegationManager.OperatorDetails memory operatorDetails,
        string memory metadataURI
    ) public filterFuzzedAddressInputs(operator) {
        _filterOperatorDetails(operator, operatorDetails);

        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorDetailsModified(operator, operatorDetails);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(operator, operator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorRegistered(operator, operatorDetails);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorMetadataURIUpdated(operator, metadataURI);

        cheats.prank(operator);
        delegationManager.registerAsOperator(operatorDetails, metadataURI);

        // Storage checks
        assertEq(
            operatorDetails.delegationApprover,
            delegationManager.delegationApprover(operator),
            "delegationApprover not set correctly"
        );
        assertEq(
            operatorDetails.stakerOptOutWindowBlocks,
            delegationManager.stakerOptOutWindowBlocks(operator),
            "stakerOptOutWindowBlocks not set correctly"
        );
        assertEq(delegationManager.delegatedTo(operator), operator, "operator not delegated to self");
    }

    // @notice Verifies that a staker who is actively delegated to an operator cannot register as an operator (without first undelegating, at least)
    function testFuzz_registerAsOperator_cannotRegisterWhileDelegated(
        address staker,
        IDelegationManager.OperatorDetails memory operatorDetails
    ) public filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != defaultOperator);
        // Staker becomes an operator, so filter against staker's address
        _filterOperatorDetails(staker, operatorDetails);

        // register *this contract* as an operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, emptySalt);

        // expect revert if attempt to register as operator
        cheats.expectRevert("DelegationManager.registerAsOperator: caller is already actively delegated");
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);

        cheats.stopPrank();
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
        IDelegationManager.OperatorDetails memory initialOperatorDetails,
        IDelegationManager.OperatorDetails memory modifiedOperatorDetails
    ) public {
        _registerOperator(defaultOperator, initialOperatorDetails, emptyStringForMetadataURI);

        cheats.startPrank(defaultOperator);

        // either it fails for trying to set the stakerOptOutWindowBlocks
        if (modifiedOperatorDetails.stakerOptOutWindowBlocks > delegationManager.MAX_STAKER_OPT_OUT_WINDOW_BLOCKS()) {
            cheats.expectRevert(
                "DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be > MAX_STAKER_OPT_OUT_WINDOW_BLOCKS"
            );
            delegationManager.modifyOperatorDetails(modifiedOperatorDetails);
            // or the transition is allowed,
        } else if (
            modifiedOperatorDetails.stakerOptOutWindowBlocks >= initialOperatorDetails.stakerOptOutWindowBlocks
        ) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorDetailsModified(defaultOperator, modifiedOperatorDetails);
            delegationManager.modifyOperatorDetails(modifiedOperatorDetails);

            assertEq(
                modifiedOperatorDetails.delegationApprover,
                delegationManager.delegationApprover(defaultOperator),
                "delegationApprover not set correctly"
            );
            assertEq(
                modifiedOperatorDetails.stakerOptOutWindowBlocks,
                delegationManager.stakerOptOutWindowBlocks(defaultOperator),
                "stakerOptOutWindowBlocks not set correctly"
            );
            assertEq(delegationManager.delegatedTo(defaultOperator), defaultOperator, "operator not delegated to self");
            // or else the transition is disallowed
        } else {
            cheats.expectRevert("DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be decreased");
            delegationManager.modifyOperatorDetails(modifiedOperatorDetails);
        }

        cheats.stopPrank();
    }

    // @notice Tests that an address which is not an operator cannot successfully call `updateOperatorMetadataURI`.
    function test_updateOperatorMetadataUri_notRegistered() public {
        assertFalse(delegationManager.isOperator(defaultOperator), "bad test setup");

        cheats.prank(defaultOperator);
        cheats.expectRevert("DelegationManager.updateOperatorMetadataURI: caller must be an operator");
        delegationManager.updateOperatorMetadataURI(emptyStringForMetadataURI);
    }

    /**
     * @notice Verifies that a staker cannot call cannot modify their `OperatorDetails` without first registering as an operator
     * @dev This is an important check to ensure that our definition of 'operator' remains consistent, in particular for preserving the
     * invariant that 'operators' are always delegated to themselves
     */
    function testFuzz_updateOperatorMetadataUri_revert_notOperator(
        IDelegationManager.OperatorDetails memory operatorDetails
    ) public {
        cheats.expectRevert("DelegationManager.modifyOperatorDetails: caller must be an operator");
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
            IDelegationManager.OperatorDetails({
                __deprecated_earningsReceiver: defaultOperator,
                delegationApprover: address(0),
                stakerOptOutWindowBlocks: 0
            }),
            emptyStringForMetadataURI
        );

        // set the pausing flag
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_NEW_DELEGATION);

        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        cheats.prank(defaultStaker);
        cheats.expectRevert("Pausable: index is paused");
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
        cheats.startPrank(staker);
        cheats.expectRevert("DelegationManager.delegateTo: staker is already actively delegated");
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
    }

    // @notice Verifies that `staker` cannot delegate to an unregistered `operator`
    function testFuzz_Revert_WhenDelegateToUnregisteredOperator(
        address staker,
        address operator
    ) public filterFuzzedAddressInputs(staker) filterFuzzedAddressInputs(operator) {
        assertFalse(delegationManager.isOperator(operator), "incorrect test input?");

        // try to delegate and check that the call reverts
        cheats.startPrank(staker);
        cheats.expectRevert("DelegationManager.delegateTo: operator is not registered in EigenLayer");
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);
        cheats.stopPrank();
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
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(defaultOperator, staker, strategyMock, shares);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
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
     * @notice `staker` delegates to a operator who does not require any signature verification similar to test above.
     * In this scenario, staker doesn't have any delegatable shares and operator shares should not increase. Staker
     * should still be correctly delegated to the operator after the call.
     */
    function testFuzz_OperatorWhoAcceptsAllStakers_ZeroDelegatableShares(
        address staker,
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 salt
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
        cheats.expectRevert("DelegationManager._delegate: approver signature expired");
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
        cheats.expectRevert("DelegationManager._delegate: approverSalt already spent");
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
        cheats.expectRevert("EIP1271SignatureUtils.checkSignature_EIP1271: signature not from signer");
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
        uint256 shares
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
        uint256 shares
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
        cheats.expectRevert("DelegationManager._delegate: approver signature expired");
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
        cheats.expectRevert("DelegationManager._delegate: approverSalt already spent");
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

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            __deprecated_earningsReceiver: defaultOperator,
            delegationApprover: address(wallet),
            stakerOptOutWindowBlocks: 0
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

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            __deprecated_earningsReceiver: defaultOperator,
            delegationApprover: address(wallet),
            stakerOptOutWindowBlocks: 0
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
        cheats.expectRevert("EIP1271SignatureUtils.checkSignature_EIP1271: ERC1271 signature verification failed");
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
            IDelegationManager.OperatorDetails({
                __deprecated_earningsReceiver: defaultOperator,
                delegationApprover: address(0),
                stakerOptOutWindowBlocks: 0
            }),
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
        cheats.expectRevert("Pausable: index is paused");
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
        cheats.expectRevert("DelegationManager.delegateToBySignature: staker signature expired");
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
        cheats.expectRevert("EIP1271SignatureUtils.checkSignature_EIP1271: signature not from signer");
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
        cheats.expectRevert("EIP1271SignatureUtils.checkSignature_EIP1271: ERC1271 signature verification failed");
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

    /// @notice Checks that `DelegationManager.delegateToBySignature` reverts when the staker is already delegated
    function test_Revert_Staker_WhenAlreadyDelegated() public {
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
        cheats.expectRevert("DelegationManager.delegateToBySignature: staker is already actively delegated");
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
        cheats.expectRevert("DelegationManager.delegateToBySignature: operator is not registered in EigenLayer");
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
        cheats.expectRevert("DelegationManager._delegate: approver signature expired");
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

        cheats.expectRevert("DelegationManager: onlyStrategyManagerOrEigenPodManager");
        delegationManager.increaseDelegatedShares(invalidCaller, strategyMock, shares);
    }

    // @notice Verifies that there is no change in shares if the staker is not delegated
    function testFuzz_increaseDelegatedShares_noop(address staker) public {
        cheats.assume(staker != defaultOperator);
        _registerOperatorWithBaseDetails(defaultOperator);
        assertFalse(delegationManager.isDelegated(staker), "bad test setup");

        cheats.prank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, strategyMock, 1);
        assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "shares should not have changed");
    }

    /**
     * @notice Verifies that `DelegationManager.increaseDelegatedShares` properly increases the delegated `shares` that the operator
     * who the `staker` is delegated to has in the strategy
     * @dev Checks that there is no change if the staker is not delegated
     */
    function testFuzz_increaseDelegatedShares(
        address staker,
        uint256 shares,
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
        }

        cheats.prank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, strategyMock, shares);

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

    // @notice Verifies that `DelegationManager.decreaseDelegatedShares` reverts if not called by the StrategyManager nor EigenPodManager
    function testFuzz_decreaseDelegatedShares_revert_invalidCaller(
        address invalidCaller,
        uint256 shares
    ) public filterFuzzedAddressInputs(invalidCaller) {
        cheats.assume(invalidCaller != address(strategyManagerMock));
        cheats.assume(invalidCaller != address(eigenPodManagerMock));

        cheats.startPrank(invalidCaller);
        cheats.expectRevert("DelegationManager: onlyStrategyManagerOrEigenPodManager");
        delegationManager.decreaseDelegatedShares(invalidCaller, strategyMock, shares);
    }

    // @notice Verifies that there is no change in shares if the staker is not delegated
    function testFuzz_decreaseDelegatedShares_noop(address staker) public {
        cheats.assume(staker != defaultOperator);
        _registerOperatorWithBaseDetails(defaultOperator);
        assertFalse(delegationManager.isDelegated(staker), "bad test setup");

        cheats.prank(address(strategyManagerMock));
        delegationManager.decreaseDelegatedShares(staker, strategyMock, 1);
        assertEq(delegationManager.operatorShares(defaultOperator, strategyMock), 0, "shares should not have changed");
    }

    /**
     * @notice Verifies that `DelegationManager.decreaseDelegatedShares` properly decreases the delegated `shares` that the operator
     * who the `staker` is delegated to has in the strategies
     * @dev Checks that there is no change if the staker is not delegated
     */
    function testFuzz_decreaseDelegatedShares(
        address staker,
        IStrategy[] memory strategies,
        uint128 shares,
        bool delegateFromStakerToOperator
    ) public filterFuzzedAddressInputs(staker) {
        // sanity-filtering on fuzzed input length & staker
        cheats.assume(strategies.length <= 32);
        cheats.assume(staker != defaultOperator);

        // Register operator
        _registerOperatorWithBaseDetails(defaultOperator);

        // delegate from the `staker` to the operator *if `delegateFromStakerToOperator` is 'true'*
        if (delegateFromStakerToOperator) {
            _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        }

        uint256[] memory sharesInputArray = new uint256[](strategies.length);

        address delegatedTo = delegationManager.delegatedTo(staker);

        // for each strategy in `strategies`, increase delegated shares by `shares`
        // noop if the staker is not delegated
        cheats.startPrank(address(strategyManagerMock));
        for (uint256 i = 0; i < strategies.length; ++i) {
            delegationManager.increaseDelegatedShares(staker, strategies[i], shares);
            // store delegated shares in a mapping
            delegatedSharesBefore[strategies[i]] = delegationManager.operatorShares(delegatedTo, strategies[i]);
            // also construct an array which we'll use in another loop
            sharesInputArray[i] = shares;
            totalSharesForStrategyInArray[address(strategies[i])] += sharesInputArray[i];
        }
        cheats.stopPrank();

        bool isDelegated = delegationManager.isDelegated(staker);

        // for each strategy in `strategies`, decrease delegated shares by `shares`
        {
            cheats.startPrank(address(strategyManagerMock));
            address operatorToDecreaseSharesOf = delegationManager.delegatedTo(staker);
            if (isDelegated) {
                for (uint256 i = 0; i < strategies.length; ++i) {
                    cheats.expectEmit(true, true, true, true, address(delegationManager));
                    emit OperatorSharesDecreased(
                        operatorToDecreaseSharesOf,
                        staker,
                        strategies[i],
                        sharesInputArray[i]
                    );
                    delegationManager.decreaseDelegatedShares(staker, strategies[i], sharesInputArray[i]);
                }
            }
            cheats.stopPrank();
        }

        // check shares after call to `decreaseDelegatedShares`
        for (uint256 i = 0; i < strategies.length; ++i) {
            uint256 delegatedSharesAfter = delegationManager.operatorShares(delegatedTo, strategies[i]);

            if (isDelegated) {
                assertEq(
                    delegatedSharesAfter + totalSharesForStrategyInArray[address(strategies[i])],
                    delegatedSharesBefore[strategies[i]],
                    "delegated shares did not decrement correctly"
                );
                assertEq(delegatedSharesAfter, 0, "nonzero shares delegated to");
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
        cheats.expectRevert("Pausable: index is paused");
        delegationManager.undelegate(staker);
    }

    function testFuzz_undelegate_revert_notDelgated(
        address undelegatedStaker
    ) public filterFuzzedAddressInputs(undelegatedStaker) {
        cheats.assume(undelegatedStaker != defaultOperator);
        assertFalse(delegationManager.isDelegated(undelegatedStaker), "bad test setup");

        cheats.prank(undelegatedStaker);
        cheats.expectRevert("DelegationManager.undelegate: staker must be delegated to undelegate");
        delegationManager.undelegate(undelegatedStaker);
    }

    // @notice Verifies that an operator cannot undelegate from themself (this should always be forbidden)
    function testFuzz_undelegate_revert_stakerIsOperator(address operator) public filterFuzzedAddressInputs(operator) {
        _registerOperatorWithBaseDetails(operator);

        cheats.prank(operator);
        cheats.expectRevert("DelegationManager.undelegate: operators cannot be undelegated");
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
        cheats.expectRevert("DelegationManager.undelegate: operators cannot be undelegated");
        delegationManager.undelegate(defaultOperator);
    }

    //TODO: verify that this check is even needed
    function test_undelegate_revert_zeroAddress() public {
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(address(0), defaultOperator);

        cheats.prank(address(0));
        cheats.expectRevert("DelegationManager.undelegate: cannot undelegate zero address");
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
        cheats.expectRevert("DelegationManager.undelegate: caller cannot undelegate staker");
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
        emit StakerForceUndelegated(staker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(staker, defaultOperator);
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
}

contract DelegationManagerUnitTests_queueWithdrawals is DelegationManagerUnitTests {
    function test_Revert_WhenEnterQueueWithdrawalsPaused() public {
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_ENTER_WITHDRAWAL_QUEUE);
        (IDelegationManager.QueuedWithdrawalParams[] memory queuedWithdrawalParams, , ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            strategy: strategyMock,
            withdrawalAmount: 100
        });
        cheats.expectRevert("Pausable: index is paused");
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
    }

    function test_Revert_WhenQueueWithdrawalParamsLengthMismatch() public {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        strategyArray[0] = strategyMock;
        uint256[] memory shareAmounts = new uint256[](2);
        shareAmounts[0] = 100;
        shareAmounts[1] = 100;

        IDelegationManager.QueuedWithdrawalParams[] memory queuedWithdrawalParams = new IDelegationManager.QueuedWithdrawalParams[](1);
        queuedWithdrawalParams[0] = IDelegationManager.QueuedWithdrawalParams({
            strategies: strategyArray,
            shares: shareAmounts,
            withdrawer: defaultStaker
        });

        cheats.expectRevert("DelegationManager.queueWithdrawal: input length mismatch");
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
    }

    function test_Revert_WhenNotStakerWithdrawer(address withdrawer) public {
        cheats.assume(withdrawer != defaultStaker);

        (IDelegationManager.QueuedWithdrawalParams[] memory queuedWithdrawalParams, , ) = _setUpQueueWithdrawalsSingleStrat({
            staker: defaultStaker,
            withdrawer: withdrawer,
            strategy: strategyMock,
            withdrawalAmount: 100
        });
        cheats.expectRevert("DelegationManager.queueWithdrawal: withdrawer must be staker");
        cheats.prank(defaultStaker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
    }

    function test_Revert_WhenEmptyStrategiesArray() public {
        IStrategy[] memory strategyArray = new IStrategy[](0);
        uint256[] memory shareAmounts = new uint256[](0);
        address withdrawer = defaultOperator;

        IDelegationManager.QueuedWithdrawalParams[] memory queuedWithdrawalParams = new IDelegationManager.QueuedWithdrawalParams[](1);
        queuedWithdrawalParams[0] = IDelegationManager.QueuedWithdrawalParams({
            strategies: strategyArray,
            shares: shareAmounts,
            withdrawer: withdrawer
        });

        cheats.expectRevert("DelegationManager._removeSharesAndQueueWithdrawal: strategies cannot be empty");
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
    function testFuzz_queueWithdrawal_SingleStrat(
        address staker,
        uint256 depositAmount,
        uint256 withdrawalAmount
    ) public filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != defaultOperator);
        cheats.assume(withdrawalAmount > 0 && withdrawalAmount <= depositAmount);
        uint256[] memory sharesAmounts = new uint256[](1);
        sharesAmounts[0] = depositAmount;
        // sharesAmounts is single element so returns single strategy
        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(staker, sharesAmounts);
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        (
            IDelegationManager.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManager.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawalsSingleStrat({
            staker: staker,
            withdrawer: staker,
            strategy: strategies[0],
            withdrawalAmount: withdrawalAmount
        });
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker should be delegated to operator");
        uint256 nonceBefore = delegationManager.cumulativeWithdrawalsQueued(staker);
        uint256 delegatedSharesBefore = delegationManager.operatorShares(defaultOperator, strategies[0]);

        // queueWithdrawals
        cheats.prank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit WithdrawalQueued(withdrawalRoot, withdrawal);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        uint256 nonceAfter = delegationManager.cumulativeWithdrawalsQueued(staker);
        uint256 delegatedSharesAfter = delegationManager.operatorShares(defaultOperator, strategies[0]);
        assertEq(nonceBefore + 1, nonceAfter, "staker nonce should have incremented");
        assertEq(delegatedSharesBefore - withdrawalAmount, delegatedSharesAfter, "delegated shares not decreased correctly");
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
    function testFuzz_queueWithdrawal_MultipleStrats(
        address staker,
        uint256[] memory depositAmounts
    ) public filterFuzzedAddressInputs(staker){
        cheats.assume(staker != defaultOperator);
        cheats.assume(depositAmounts.length > 0 && depositAmounts.length <= 32);
        uint256[] memory withdrawalAmounts = _fuzzWithdrawalAmounts(depositAmounts);

        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(staker, depositAmounts);
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        (
            IDelegationManager.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManager.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawals({
            staker: staker,
            withdrawer: staker,
            strategies: strategies,
            withdrawalAmounts: withdrawalAmounts
        });
        // Before queueWithdrawal state values
        uint256 nonceBefore = delegationManager.cumulativeWithdrawalsQueued(staker);
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker should be delegated to operator");
        uint256[] memory delegatedSharesBefore = new uint256[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            delegatedSharesBefore[i] = delegationManager.operatorShares(defaultOperator, strategies[i]);
        }

        // queueWithdrawals
        cheats.prank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit WithdrawalQueued(withdrawalRoot, withdrawal);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        // Post queueWithdrawal state values
        for (uint256 i = 0; i < strategies.length; i++) {
            assertEq(
                delegatedSharesBefore[i] - withdrawalAmounts[i], // Shares before - withdrawal amount
                delegationManager.operatorShares(defaultOperator, strategies[i]), // Shares after
                "delegated shares not decreased correctly"
            );
        }
        uint256 nonceAfter = delegationManager.cumulativeWithdrawalsQueued(staker);
        assertEq(nonceBefore + 1, nonceAfter, "staker nonce should have incremented");
    }

    /**
     * @notice Verifies that `DelegationManager.queueWithdrawals` properly queues a withdrawal for the `withdrawer`
     * with multiple strategies and sharesAmounts and with thirdPartyTransfersForbidden for one of the strategies.
     * Queuing a withdrawal should pass as the `withdrawer` address is the same as the staker.
     * 
     * Depending on length sharesAmounts, deploys corresponding number of strategies
     * and deposits sharesAmounts into each strategy for the staker and delegates to operator.
     * For each strategy, withdrawAmount <= depositAmount
     * - Asserts that staker is delegated to the operator
     * - Asserts that shares for delegatedTo operator are decreased by `sharesAmount`
     * - Asserts that staker cumulativeWithdrawalsQueued nonce is incremented
     * - Checks that event was emitted with correct withdrawalRoot and withdrawal
     */
    function testFuzz_queueWithdrawal_ThirdPartyTransfersForbidden(
        address staker,
        uint256[] memory depositAmounts,
        uint256 randSalt
    ) public filterFuzzedAddressInputs(staker){
        cheats.assume(depositAmounts.length > 0 && depositAmounts.length <= 32);
        cheats.assume(staker != defaultOperator);
        uint256[] memory withdrawalAmounts = _fuzzWithdrawalAmounts(depositAmounts);

        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(staker, depositAmounts);
        // Randomly set strategy true for thirdPartyTransfersForbidden
        uint256 randStrategyIndex = randSalt % strategies.length;
        strategyManagerMock.setThirdPartyTransfersForbidden(strategies[randStrategyIndex], true);

        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        (
            IDelegationManager.QueuedWithdrawalParams[] memory queuedWithdrawalParams,
            IDelegationManager.Withdrawal memory withdrawal,
            bytes32 withdrawalRoot
        ) = _setUpQueueWithdrawals({
            staker: staker,
            withdrawer: staker,
            strategies: strategies,
            withdrawalAmounts: withdrawalAmounts
        });
        // Before queueWithdrawal state values
        uint256 nonceBefore = delegationManager.cumulativeWithdrawalsQueued(staker);
        assertEq(delegationManager.delegatedTo(staker), defaultOperator, "staker should be delegated to operator");
        uint256[] memory delegatedSharesBefore = new uint256[](strategies.length);
        for (uint256 i = 0; i < strategies.length; i++) {
            delegatedSharesBefore[i] = delegationManager.operatorShares(defaultOperator, strategies[i]);
        }

        // queueWithdrawals
        cheats.prank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit WithdrawalQueued(withdrawalRoot, withdrawal);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);

        // Post queueWithdrawal state values
        for (uint256 i = 0; i < strategies.length; i++) {
            assertEq(
                delegatedSharesBefore[i] - withdrawalAmounts[i], // Shares before - withdrawal amount
                delegationManager.operatorShares(defaultOperator, strategies[i]), // Shares after
                "delegated shares not decreased correctly"
            );
        }
        uint256 nonceAfter = delegationManager.cumulativeWithdrawalsQueued(staker);
        assertEq(nonceBefore + 1, nonceAfter, "staker nonce should have incremented");
    }

    /**
     * @notice Randomly selects one of the strategies to set thirdPartyTransfersForbidden to true.
     * Verifies that `DelegationManager.queueWithdrawals` properly reverts a queuedWithdrawal since the `withdrawer`
     * is not the same as the `staker`.
     */
    function testFuzz_queueWithdrawal_Revert_WhenThirdPartyTransfersForbidden(
        address staker,
        address withdrawer,
        uint256[] memory depositAmounts,
        uint256 randSalt
    ) public filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != withdrawer && staker != defaultOperator);
        cheats.assume(depositAmounts.length > 0 && depositAmounts.length <= 32);
        uint256[] memory withdrawalAmounts = _fuzzWithdrawalAmounts(depositAmounts);

        IStrategy[] memory strategies = _deployAndDepositIntoStrategies(staker, depositAmounts);
        // Randomly set strategy true for thirdPartyTransfersForbidden
        uint256 randStrategyIndex = randSalt % strategies.length;
        strategyManagerMock.setThirdPartyTransfersForbidden(strategies[randStrategyIndex], true);
        _registerOperatorWithBaseDetails(defaultOperator);
        _delegateToOperatorWhoAcceptsAllStakers(staker, defaultOperator);
        (IDelegationManager.QueuedWithdrawalParams[] memory queuedWithdrawalParams, , ) = _setUpQueueWithdrawals({
            staker: staker,
            withdrawer: withdrawer,
            strategies: strategies,
            withdrawalAmounts: withdrawalAmounts
        });

        // queueWithdrawals
        // NOTE: Originally, you could queue a withdrawal to a different address, which would fail with a specific error
        // if third party transfers were forbidden. Now, withdrawing to a different address is forbidden regardless
        // of third party transfer status.
        cheats.expectRevert(
            "DelegationManager.queueWithdrawal: withdrawer must be staker"
        );
        cheats.prank(staker);
        delegationManager.queueWithdrawals(queuedWithdrawalParams);
    }
}

contract DelegationManagerUnitTests_completeQueuedWithdrawal is DelegationManagerUnitTests {
    function test_Revert_WhenExitWithdrawalQueuePaused() public {
        cheats.prank(pauser);
        delegationManager.pause(2 ** PAUSED_EXIT_WITHDRAWAL_QUEUE);
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            IDelegationManager.Withdrawal memory withdrawal,
            IERC20[] memory tokens,
            /* bytes32 withdrawalRoot */
        ) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            depositAmount: 100,
            withdrawalAmount: 100
        });
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        cheats.expectRevert("Pausable: index is paused");
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, 0 /* middlewareTimesIndex */, false);
    }

    function test_Revert_WhenInvalidWithdrawalRoot() public {
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            IDelegationManager.Withdrawal memory withdrawal,
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
        cheats.roll(block.number + delegationManager.getWithdrawalDelay(withdrawal.strategies));
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, 0 /* middlewareTimesIndex */, false);
        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be completed and marked false now");

        cheats.roll(block.number + delegationManager.getWithdrawalDelay(withdrawal.strategies));
        cheats.expectRevert("DelegationManager._completeQueuedWithdrawal: action is not in queue");
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, 0 /* middlewareTimesIndex */, false);
    }

    /**
     * @notice should revert if minWithdrawalDelayBlocks has not passed, and if
     * delegationManager.getWithdrawalDelay returns a value greater than minWithdrawalDelayBlocks
     * then it should revert if the validBlockNumber has not passed either.
     */
    function test_Revert_WhenWithdrawalDelayBlocksNotPassed(
        uint256[] memory depositAmounts,
        uint256 randSalt,
        bool receiveAsTokens
    ) public {
        cheats.assume(depositAmounts.length > 0 && depositAmounts.length <= 32);
        uint256[] memory withdrawalAmounts = _fuzzWithdrawalAmounts(depositAmounts);
        
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            IDelegationManager.Withdrawal memory withdrawal,
            IERC20[] memory tokens,
            /* bytes32 withdrawalRoot */
        ) = _setUpCompleteQueuedWithdrawal({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            depositAmounts: depositAmounts,
            withdrawalAmounts: withdrawalAmounts
        });

        // prank as withdrawer address
        cheats.startPrank(defaultStaker);

        cheats.expectRevert(
            "DelegationManager._completeQueuedWithdrawal: minWithdrawalDelayBlocks period has not yet passed"
        );
        cheats.roll(block.number + minWithdrawalDelayBlocks - 1);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, 0 /* middlewareTimesIndex */, receiveAsTokens);

        uint256 validBlockNumber = delegationManager.getWithdrawalDelay(withdrawal.strategies);
        if (validBlockNumber > minWithdrawalDelayBlocks) {
            cheats.expectRevert(
                "DelegationManager._completeQueuedWithdrawal: withdrawalDelayBlocks period has not yet passed for this strategy"
            );
            cheats.roll(validBlockNumber - 1);
            delegationManager.completeQueuedWithdrawal(withdrawal, tokens, 0 /* middlewareTimesIndex */, receiveAsTokens);
        }

        cheats.stopPrank();
    }

    /**
     * @notice should revert when the withdrawalDelayBlocks period has not yet passed for the
     * beacon chain strategy
     */
    function test_Revert_WhenWithdrawalDelayBlocksNotPassed_BeaconStrat(
        uint256 depositAmount,
        uint256 withdrawalAmount,
        uint256 beaconWithdrawalDelay
    ) public {
        cheats.assume(depositAmount > 1 && withdrawalAmount <= depositAmount);
        beaconWithdrawalDelay = bound(beaconWithdrawalDelay, minWithdrawalDelayBlocks, MAX_WITHDRAWAL_DELAY_BLOCKS);
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            IDelegationManager.Withdrawal memory withdrawal,
            IERC20[] memory tokens,
            bytes32 withdrawalRoot
        ) = _setUpCompleteQueuedWithdrawalBeaconStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            depositAmount: depositAmount,
            withdrawalAmount: withdrawalAmount
        });

        IStrategy[] memory strategies = new IStrategy[](1);
        strategies[0] = beaconChainETHStrategy;
        uint256[] memory withdrawalDelayBlocks = new uint256[](1);
        delegationManager.setStrategyWithdrawalDelayBlocks(withdrawal.strategies, withdrawalDelayBlocks);

        // prank as withdrawer address
        cheats.startPrank(defaultStaker);

        cheats.expectRevert(
            "DelegationManager._completeQueuedWithdrawal: minWithdrawalDelayBlocks period has not yet passed"
        );
        cheats.roll(block.number + minWithdrawalDelayBlocks - 1);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, 0 /* middlewareTimesIndex */, false);

        uint256 validBlockNumber = delegationManager.getWithdrawalDelay(withdrawal.strategies);
        if (validBlockNumber > minWithdrawalDelayBlocks) {
            cheats.expectRevert(
                "DelegationManager._completeQueuedWithdrawal: withdrawalDelayBlocks period has not yet passed for this strategy"
            );
            cheats.roll(validBlockNumber - 1);
            delegationManager.completeQueuedWithdrawal(withdrawal, tokens, 0 /* middlewareTimesIndex */, false);
        }

        cheats.stopPrank();
    }

    function test_Revert_WhenNotCalledByWithdrawer() public {
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            IDelegationManager.Withdrawal memory withdrawal,
            IERC20[] memory tokens,
            /*bytes32 withdrawalRoot*/
        ) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            depositAmount: 100,
            withdrawalAmount: 100
        });
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        cheats.roll(block.number + delegationManager.getWithdrawalDelay(withdrawal.strategies));
        cheats.expectRevert("DelegationManager._completeQueuedWithdrawal: only withdrawer can complete action");
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, 0 /* middlewareTimesIndex */, false);
    }

    function test_Revert_WhenTokensArrayLengthMismatch() public {
        _registerOperatorWithBaseDetails(defaultOperator);
        (IDelegationManager.Withdrawal memory withdrawal, , ) = _setUpCompleteQueuedWithdrawalSingleStrat({
            staker: defaultStaker,
            withdrawer: defaultStaker,
            depositAmount: 100,
            withdrawalAmount: 100
        });
        _delegateToOperatorWhoAcceptsAllStakers(defaultStaker, defaultOperator);

        IERC20[] memory tokens = new IERC20[](0);
        cheats.roll(block.number + delegationManager.getWithdrawalDelay(withdrawal.strategies));
        cheats.expectRevert("DelegationManager._completeQueuedWithdrawal: input length mismatch");
        cheats.prank(defaultStaker);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, 0 /* middlewareTimesIndex */, true);
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
        uint256 depositAmount,
        uint256 withdrawalAmount
    ) public filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != defaultOperator);
        cheats.assume(withdrawalAmount > 0 && withdrawalAmount <= depositAmount);
        _registerOperatorWithBaseDetails(defaultOperator);
        (
            IDelegationManager.Withdrawal memory withdrawal,
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
        cheats.roll(block.number + delegationManager.getWithdrawalDelay(withdrawal.strategies));
        cheats.prank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit WithdrawalCompleted(withdrawalRoot);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, 0 /* middlewareTimesIndex */, true);

        uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, withdrawal.strategies[0]);
        assertEq(operatorSharesAfter, operatorSharesBefore, "operator shares should be unchanged");
        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be completed and marked false now");
    }

    /**
     * @notice Verifies that `DelegationManager.completeQueuedWithdrawal` properly completes a queued withdrawal for the `withdrawer`
     * for a single strategy. Withdraws as shares so if the withdrawer is delegated, operator shares increase. In the test case, this only
     * happens if staker and withdrawer are fuzzed the same address (i.e. staker == withdrawer)
     * - Asserts that the withdrawalRoot is True before `completeQueuedWithdrawal` and False after
     * - Asserts if staker == withdrawer, operatorShares increase, otherwise operatorShares are unchanged
     * - Checks that event `WithdrawalCompleted` is emitted with withdrawalRoot
     */
    function test_completeQueuedWithdrawal_SingleStratWithdrawAsShares(
        address staker,
        uint256 depositAmount,
        uint256 withdrawalAmount
    ) public filterFuzzedAddressInputs(staker) {
        cheats.assume(staker != defaultOperator);
        cheats.assume(withdrawalAmount > 0 && withdrawalAmount <= depositAmount);
        _registerOperatorWithBaseDetails(defaultOperator);

        (
            IDelegationManager.Withdrawal memory withdrawal,
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
        cheats.roll(block.number + delegationManager.getWithdrawalDelay(withdrawal.strategies));
        cheats.prank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit WithdrawalCompleted(withdrawalRoot);
        delegationManager.completeQueuedWithdrawal(withdrawal, tokens, 0 /* middlewareTimesIndex */, false);

        uint256 operatorSharesAfter = delegationManager.operatorShares(defaultOperator, withdrawal.strategies[0]);
        // Since staker is delegated, operatorShares get incremented
        assertEq(operatorSharesAfter, operatorSharesBefore + withdrawalAmount, "operator shares not increased correctly");
        assertFalse(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRoot should be completed and marked false now");
    }
}
