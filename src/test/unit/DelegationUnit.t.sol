// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

import "src/contracts/core/DelegationManager.sol";
import "src/contracts/strategies/StrategyBase.sol";

import "src/test/events/IDelegationManagerEvents.sol";
import "src/test/mocks/StakeRegistryStub.sol";
import "src/test/mocks/Reenterer.sol";
import "src/test/mocks/EmptyContract.sol";
import "src/test/utils/EigenLayerUnitTestSetup.sol";

/**
 * @notice Unit testing of the DelegationManager contract. Withdrawals are tightly coupled
 * with EigenPodManager and StrategyManager and are part of integration tests.
 * Contracts tested: DelegationMananger
 * Contracts not mocked: StrategyBase, PauserRegistry
 */
contract DelegationManangerUnitTests is EigenLayerUnitTestSetup, IDelegationManagerEvents {
    // Contract under test
    DelegationManager delegationManager;
    DelegationManager delegationManagerImplementation;
    
    // Mocks
    StrategyBase strategyImplementation;
    StrategyBase strategyMock;
    StrategyBase strategyMock2;
    StrategyBase strategyMock3;
    IERC20 mockToken;
    uint256 mockTokenInitialSupply = 10e50;
    StakeRegistryStub stakeRegistryMock;
    Reenterer public reenterer;
    EmptyContract public emptyContract;

    // Below vars used as transient storage to fix stack-too-deep errors
    IStrategy public _tempStrategyStorage;
    address public _tempStakerStorage;
    
    // Delegation signer 
    uint256 delegationSignerPrivateKey = uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);
    uint256 stakerPrivateKey = uint256(123456789);

    // empty string reused across many tests
    string emptyStringForMetadataURI;

    // "empty" / zero salt, reused across many tests
    bytes32 emptySalt;

    // reused in various tests. in storage to help handle stack-too-deep errors
    address defaultOperator = address(this);

    // Index for flag that pauses new delegations when set.
    uint8 internal constant PAUSED_NEW_DELEGATION = 0;

    // Index for flag that pauses queuing new withdrawals when set.
    uint8 internal constant PAUSED_ENTER_WITHDRAWAL_QUEUE = 1;

    // Index for flag that pauses completing existing withdrawals when set.
    uint8 internal constant PAUSED_EXIT_WITHDRAWAL_QUEUE = 2;
    
    /// @notice mappings used to handle duplicate entries in fuzzed address array input
    mapping(address => uint256) public totalSharesForStrategyInArray;
    mapping(IStrategy => uint256) public delegatedSharesBefore;

    function setUp() public virtual override {
        // Setup
        EigenLayerUnitTestSetup.setUp();

        // Deploy DelegationManager implmentation and proxy
        delegationManagerImplementation = new DelegationManager(strategyManagerMock, slasherMock, eigenPodManagerMock);
        delegationManager = DelegationManager(
            address(
                new TransparentUpgradeableProxy(
                    address(delegationManagerImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(DelegationManager.initialize.selector, address(this), pauserRegistry, 0) // 0 is initialPausedStatus
                )
            )
        );

        // Deploy mock stake registry and set
        stakeRegistryMock = new StakeRegistryStub();
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakeRegistrySet(stakeRegistryMock);
        delegationManager.setStakeRegistry(stakeRegistryMock);

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
    }


    /**
    * INTERNAL / HELPER FUNCTIONS
    */

    function _setUpWithdrawalStructSingleStrat(address staker, address withdrawer, IERC20 token, IStrategy strategy, uint256 shareAmount)
        internal view returns (IDelegationManager.Withdrawal memory queuedWithdrawal, IERC20[] memory tokensArray, bytes32 withdrawalRoot)
    {
        IStrategy[] memory strategyArray = new IStrategy[](1);
        tokensArray = new IERC20[](1);
        uint256[] memory shareAmounts = new uint256[](1);
        strategyArray[0] = strategy;
        tokensArray[0] = token;
        shareAmounts[0] = shareAmount;
        queuedWithdrawal = 
            IDelegationManager.Withdrawal({
                strategies: strategyArray,
                shares: shareAmounts,
                staker: staker,
                withdrawer: withdrawer,
                nonce: delegationManager.cumulativeWithdrawalsQueued(staker),
                startBlock: uint32(block.number),
                delegatedTo: delegationManager.delegatedTo(staker)
            }
        );
        // calculate the withdrawal root
        withdrawalRoot = delegationManager.calculateWithdrawalRoot(queuedWithdrawal);
        return (queuedWithdrawal, tokensArray, withdrawalRoot);
    }

    function _setUpWithdrawalStruct_MultipleStrategies(
        address staker,
        address withdrawer,
        IStrategy[] memory strategyArray,
        uint256[] memory shareAmounts
    )
        internal view returns (IDelegationManager.Withdrawal memory withdrawal, bytes32 withdrawalRoot)
    {
        withdrawal = 
            IDelegationManager.Withdrawal({
                strategies: strategyArray,
                shares: shareAmounts,
                staker: staker,
                withdrawer: withdrawer,
                nonce: delegationManager.cumulativeWithdrawalsQueued(staker),
                startBlock: uint32(block.number),
                delegatedTo: delegationManager.delegatedTo(staker)
            }
        );
        // calculate the withdrawal root
        withdrawalRoot = delegationManager.calculateWithdrawalRoot(withdrawal);
        return (withdrawal, withdrawalRoot);
    }

    function _setUpWithdrawalStructSingleStrat_MultipleStrategies(
        address staker,
        address withdrawer,
        IStrategy[] memory strategyArray,
        uint256[] memory shareAmounts
    )
        internal view returns (IDelegationManager.Withdrawal memory queuedWithdrawal, bytes32 withdrawalRoot)
    {
        queuedWithdrawal = 
            IDelegationManager.Withdrawal({
                staker: staker,
                withdrawer: withdrawer,
                nonce: delegationManager.cumulativeWithdrawalsQueued(staker),
                startBlock: uint32(block.number),
                delegatedTo: delegationManager.delegatedTo(staker),
                strategies: strategyArray,
                shares: shareAmounts
            }
        );
        // calculate the withdrawal root
        withdrawalRoot = delegationManager.calculateWithdrawalRoot(queuedWithdrawal);
        return (queuedWithdrawal, withdrawalRoot);
    }

    /**
     * @notice internal function for calculating a signature from the delegationSigner corresponding to `_delegationSignerPrivateKey`, approving
     * the `staker` to delegate to `operator`, with the specified `salt`, and expiring at `expiry`.
     */
    function _getApproverSignature(uint256 _delegationSignerPrivateKey, address staker, address operator, bytes32 salt, uint256 expiry)
        internal view returns (ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry)
    {
        approverSignatureAndExpiry.expiry = expiry;
        {
            bytes32 digestHash =
                delegationManager.calculateDelegationApprovalDigestHash(staker, operator, delegationManager.delegationApprover(operator), salt, expiry);
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(_delegationSignerPrivateKey, digestHash);
            approverSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }
        return approverSignatureAndExpiry;
    }

    /**
     * @notice internal function for calculating a signature from the staker corresponding to `_stakerPrivateKey`, delegating them to
     * the `operator`, and expiring at `expiry`.
     */
    function _getStakerSignature(uint256 _stakerPrivateKey, address operator, uint256 expiry)
        internal view returns (ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry)
    {
        address staker = cheats.addr(stakerPrivateKey);
        stakerSignatureAndExpiry.expiry = expiry;
        {
            bytes32 digestHash = delegationManager.calculateCurrentStakerDelegationDigestHash(staker, operator, expiry);
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(_stakerPrivateKey, digestHash);
            stakerSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }
        return stakerSignatureAndExpiry;
    }
}

contract  DelegationManagerUnitTests_Initialization_Setters is DelegationManangerUnitTests {
    function test_initialization() public {
        assertEq(address(delegationManager.strategyManager()), address(strategyManagerMock), "constructor / initializer incorrect, strategyManager set wrong");
        assertEq(address(delegationManager.slasher()), address(slasherMock), "constructor / initializer incorrect, slasher set wrong");
        assertEq(address(delegationManager.pauserRegistry()), address(pauserRegistry), "constructor / initializer incorrect, pauserRegistry set wrong");
        assertEq(delegationManager.owner(), address(this), "constructor / initializer incorrect, owner set wrong");
        assertEq(delegationManager.paused(), 0, "constructor / initializer incorrect, paused status set wrong");
    }

    /// @notice Verifies that the DelegationManager cannot be iniitalized multiple times
    function test_initialize_revert_reinitialization() public {
        cheats.expectRevert("Initializable: contract is already initialized");
        delegationManager.initialize(address(this), pauserRegistry, 0);
    }

    /// @notice Verifies that the stakeRegistry cannot be set after it has already been set
    function test_setStakeRegistry_revert_alreadySet() public {
        cheats.expectRevert("DelegationManager.setStakeRegistry: stakeRegistry already set");
        delegationManager.setStakeRegistry(stakeRegistryMock);
    }

    function testFuzz_setWithdrawalDelayBlocks_revert_notOwner(address invalidCaller) public filterFuzzedAddressInputs(invalidCaller) {
        cheats.prank(invalidCaller);
        cheats.expectRevert("Ownable: caller is not the owner");
        delegationManager.setWithdrawalDelayBlocks(0);
    }

    function testFuzz_setWithdrawalDelayBlocks_revert_tooLarge(uint256 newWithdrawalDelayBlocks) external {
        // filter fuzzed inputs to disallowed amounts
        cheats.assume(newWithdrawalDelayBlocks > delegationManager.MAX_WITHDRAWAL_DELAY_BLOCKS());

        // attempt to set the `withdrawalDelayBlocks` variable
        cheats.expectRevert("DelegationManager.setWithdrawalDelayBlocks: newWithdrawalDelayBlocks too high");
        delegationManager.setWithdrawalDelayBlocks(newWithdrawalDelayBlocks);
    }

    function testFuzz_setWithdrawalDelayBlocks(uint256 newWithdrawalDelayBlocks) public {
        cheats.assume(newWithdrawalDelayBlocks <= delegationManager.MAX_WITHDRAWAL_DELAY_BLOCKS());

        // set the `withdrawalDelayBlocks` variable
        uint256 previousDelayBlocks = delegationManager.withdrawalDelayBlocks();
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit WithdrawalDelayBlocksSet(previousDelayBlocks, newWithdrawalDelayBlocks);
        delegationManager.setWithdrawalDelayBlocks(newWithdrawalDelayBlocks);
        
        // Check storage
        assertEq(delegationManager.withdrawalDelayBlocks(), newWithdrawalDelayBlocks, "withdrawalDelayBlocks not set correctly");
    }
}

contract DelegationManagerUnitTests_Old is DelegationManangerUnitTests {
    /**
     * @notice `operator` registers via calling `DelegationManager.registerAsOperator(operatorDetails, metadataURI)`
     * Should be able to set any parameters, other than setting their `earningsReceiver` to the zero address or too high value for `stakerOptOutWindowBlocks`
     * The set parameters should match the desired parameters (correct storage update)
     * Operator becomes delegated to themselves
     * Properly emits events – especially the `OperatorRegistered` event, but also `StakerDelegated` & `OperatorDetailsModified` events
     * Reverts appropriately if operator was already delegated to someone (including themselves, i.e. they were already an operator)
     * @param operator and @param operatorDetails are fuzzed inputs
     */
    function testRegisterAsOperator(address operator, IDelegationManager.OperatorDetails memory operatorDetails, string memory metadataURI) public 
        filterFuzzedAddressInputs(operator)
    {
        // filter out zero address since people can't delegate to the zero address and operators are delegated to themselves
        cheats.assume(operator != address(0));
        // filter out zero address since people can't set their earningsReceiver address to the zero address (special test case to verify)
        cheats.assume(operatorDetails.earningsReceiver != address(0));
        // filter out disallowed stakerOptOutWindowBlocks values
        cheats.assume(operatorDetails.stakerOptOutWindowBlocks <= delegationManager.MAX_STAKER_OPT_OUT_WINDOW_BLOCKS());

        cheats.startPrank(operator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorDetailsModified(operator, operatorDetails);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(operator, operator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorRegistered(operator, operatorDetails);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorMetadataURIUpdated(operator, metadataURI);

        delegationManager.registerAsOperator(operatorDetails, metadataURI);

        require(operatorDetails.earningsReceiver == delegationManager.earningsReceiver(operator), "earningsReceiver not set correctly");
        require(operatorDetails.delegationApprover == delegationManager.delegationApprover(operator), "delegationApprover not set correctly");
        require(operatorDetails.stakerOptOutWindowBlocks == delegationManager.stakerOptOutWindowBlocks(operator), "stakerOptOutWindowBlocks not set correctly");
        require(delegationManager.delegatedTo(operator) == operator, "operator not delegated to self");
        cheats.stopPrank();
    }

    /**
     * @notice Verifies that an operator cannot register with `stakerOptOutWindowBlocks` set larger than `MAX_STAKER_OPT_OUT_WINDOW_BLOCKS`
     * @param operatorDetails is a fuzzed input
     */
    function testCannotRegisterAsOperatorWithDisallowedStakerOptOutWindowBlocks(IDelegationManager.OperatorDetails memory operatorDetails) public {
        // filter out zero address since people can't set their earningsReceiver address to the zero address (special test case to verify)
        cheats.assume(operatorDetails.earningsReceiver != address(0));
        // filter out *allowed* stakerOptOutWindowBlocks values
        cheats.assume(operatorDetails.stakerOptOutWindowBlocks > delegationManager.MAX_STAKER_OPT_OUT_WINDOW_BLOCKS());

        cheats.startPrank(defaultOperator);
        cheats.expectRevert(bytes("DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be > MAX_STAKER_OPT_OUT_WINDOW_BLOCKS"));
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        cheats.stopPrank();
    }
    
    /**
     * @notice Verifies that an operator cannot register with `earningsReceiver` set to the zero address
     * @dev This is an important check since we check `earningsReceiver != address(0)` to check if an address is an operator!
     */
    function testCannotRegisterAsOperatorWithZeroAddressAsEarningsReceiver() public {
        cheats.startPrank(defaultOperator);
        IDelegationManager.OperatorDetails memory operatorDetails;
        cheats.expectRevert(bytes("DelegationManager._setOperatorDetails: cannot set `earningsReceiver` to zero address"));
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        cheats.stopPrank();
    }

    // @notice Verifies that someone cannot successfully call `DelegationManager.registerAsOperator(operatorDetails)` again after registering for the first time
    function testCannotRegisterAsOperatorMultipleTimes(address operator, IDelegationManager.OperatorDetails memory operatorDetails) public 
        filterFuzzedAddressInputs(operator)
    {
        testRegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);
        cheats.startPrank(operator);
        cheats.expectRevert(bytes("DelegationManager.registerAsOperator: operator has already registered"));
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        cheats.stopPrank();
    }

    // @notice Verifies that a staker who is actively delegated to an operator cannot register as an operator (without first undelegating, at least)
    function testCannotRegisterAsOperatorWhileDelegated(address staker, IDelegationManager.OperatorDetails memory operatorDetails) public 
        filterFuzzedAddressInputs(staker)
    {
        // filter out disallowed stakerOptOutWindowBlocks values
        cheats.assume(operatorDetails.stakerOptOutWindowBlocks <= delegationManager.MAX_STAKER_OPT_OUT_WINDOW_BLOCKS());
        // filter out zero address since people can't set their earningsReceiver address to the zero address (special test case to verify)
        cheats.assume(operatorDetails.earningsReceiver != address(0));
        // filter out case where staker *is* the operator
        cheats.assume(staker != defaultOperator);

        // register *this contract* as an operator
        IDelegationManager.OperatorDetails memory _operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: defaultOperator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(defaultOperator, _operatorDetails, emptyStringForMetadataURI);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, emptySalt);        

        cheats.expectRevert(bytes("DelegationManager._delegate: staker is already actively delegated"));
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
    function testModifyOperatorParameters(
        IDelegationManager.OperatorDetails memory initialOperatorDetails,
        IDelegationManager.OperatorDetails memory modifiedOperatorDetails
    ) public {
        testRegisterAsOperator(defaultOperator, initialOperatorDetails, emptyStringForMetadataURI);
        // filter out zero address since people can't set their earningsReceiver address to the zero address (special test case to verify)
        cheats.assume(modifiedOperatorDetails.earningsReceiver != address(0));

        cheats.startPrank(defaultOperator);

        // either it fails for trying to set the stakerOptOutWindowBlocks
        if (modifiedOperatorDetails.stakerOptOutWindowBlocks > delegationManager.MAX_STAKER_OPT_OUT_WINDOW_BLOCKS()) {
            cheats.expectRevert(bytes("DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be > MAX_STAKER_OPT_OUT_WINDOW_BLOCKS"));
            delegationManager.modifyOperatorDetails(modifiedOperatorDetails);
        // or the transition is allowed,
        } else if (modifiedOperatorDetails.stakerOptOutWindowBlocks >= initialOperatorDetails.stakerOptOutWindowBlocks) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorDetailsModified(defaultOperator, modifiedOperatorDetails);
            delegationManager.modifyOperatorDetails(modifiedOperatorDetails);

            require(modifiedOperatorDetails.earningsReceiver == delegationManager.earningsReceiver(defaultOperator), "earningsReceiver not set correctly");
            require(modifiedOperatorDetails.delegationApprover == delegationManager.delegationApprover(defaultOperator), "delegationApprover not set correctly");
            require(modifiedOperatorDetails.stakerOptOutWindowBlocks == delegationManager.stakerOptOutWindowBlocks(defaultOperator), "stakerOptOutWindowBlocks not set correctly");
            require(delegationManager.delegatedTo(defaultOperator) == defaultOperator, "operator not delegated to self");
        // or else the transition is disallowed
        } else {
            cheats.expectRevert(bytes("DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be decreased"));
            delegationManager.modifyOperatorDetails(modifiedOperatorDetails);
        }

        cheats.stopPrank();
    }

    // @notice Tests that an operator who calls `updateOperatorMetadataURI` will correctly see an `OperatorMetadataURIUpdated` event emitted with their input
    function testUpdateOperatorMetadataURI(string memory metadataURI) public {
        // register *this contract* as an operator
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: defaultOperator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(defaultOperator, operatorDetails, emptyStringForMetadataURI);

        // call `updateOperatorMetadataURI` and check for event
        cheats.startPrank(defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorMetadataURIUpdated(defaultOperator, metadataURI);
        delegationManager.updateOperatorMetadataURI(metadataURI);
        cheats.stopPrank();
    }

    // @notice Tests that an address which is not an operator cannot successfully call `updateOperatorMetadataURI`.
    function testCannotUpdateOperatorMetadataURIWithoutRegisteringFirst() public {
        require(!delegationManager.isOperator(defaultOperator), "bad test setup");

        cheats.startPrank(defaultOperator);
        cheats.expectRevert(bytes("DelegationManager.updateOperatorMetadataURI: caller must be an operator"));
        delegationManager.updateOperatorMetadataURI(emptyStringForMetadataURI);
        cheats.stopPrank();
    }

    /**
     * @notice Verifies that an operator cannot modify their `earningsReceiver` address to set it to the zero address
     * @dev This is an important check since we check `earningsReceiver != address(0)` to check if an address is an operator!
     */
    function testCannotModifyEarningsReceiverAddressToZeroAddress() public {
        // register *this contract* as an operator
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: defaultOperator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(defaultOperator, operatorDetails, emptyStringForMetadataURI);

        operatorDetails.earningsReceiver = address(0);
        cheats.expectRevert(bytes("DelegationManager._setOperatorDetails: cannot set `earningsReceiver` to zero address"));
        delegationManager.modifyOperatorDetails(operatorDetails);
    }

    /**
     * @notice Verifies that a staker cannot call cannot modify their `OperatorDetails` without first registering as an operator
     * @dev This is an important check to ensure that our definition of 'operator' remains consistent, in particular for preserving the
     * invariant that 'operators' are always delegated to themselves
     */
    function testCannotModifyOperatorDetailsWithoutRegistering(IDelegationManager.OperatorDetails memory operatorDetails) public {
        cheats.expectRevert(bytes("DelegationManager.modifyOperatorDetails: caller must be an operator"));
        delegationManager.modifyOperatorDetails(operatorDetails);
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
    function testDelegateToOperatorWhoAcceptsAllStakers(address staker, ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry, bytes32 salt) public 
        filterFuzzedAddressInputs(staker)
    {
        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: defaultOperator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(defaultOperator, operatorDetails, emptyStringForMetadataURI);

        // verify that the salt hasn't been used before
        require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt), "salt somehow spent too early?");

        IStrategy[] memory strategiesToReturn = new IStrategy[](1);
        strategiesToReturn[0] = strategyMock;
        uint256[] memory sharesToReturn = new uint256[](1);
        sharesToReturn[0] = 1;
        strategyManagerMock.setDeposits(strategiesToReturn, sharesToReturn);
        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorSharesIncreased(defaultOperator, staker, strategyMock, 1);
        delegationManager.delegateTo(defaultOperator, approverSignatureAndExpiry, salt);        
        cheats.stopPrank();

        require(delegationManager.isDelegated(staker), "staker not delegated correctly");
        require(delegationManager.delegatedTo(staker) == defaultOperator, "staker delegated to the wrong address");
        require(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // verify that the salt is still marked as unused (since it wasn't checked or used)
        require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt), "salt somehow spent too early?");
    }

    /**
     * @notice Delegates from `staker` to an operator, then verifies that the `staker` cannot delegate to another `operator` (at least without first undelegating)
     */
    function testCannotDelegateWhileDelegated(address staker, address operator, ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry, bytes32 salt) public 
        filterFuzzedAddressInputs(staker)
        filterFuzzedAddressInputs(operator)
    {
        // filter out input since if the staker tries to delegate again after registering as an operator, we will revert earlier than this test is designed to check
        cheats.assume(staker != operator);

        // delegate from the staker to an operator
        testDelegateToOperatorWhoAcceptsAllStakers(staker, approverSignatureAndExpiry, salt);

        // register another operator
        // filter out this contract, since we already register it as an operator in the above step
        cheats.assume(operator != address(this));
        IDelegationManager.OperatorDetails memory _operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, _operatorDetails, emptyStringForMetadataURI);

        // try to delegate again and check that the call reverts
        cheats.startPrank(staker);
        cheats.expectRevert(bytes("DelegationManager._delegate: staker is already actively delegated"));
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);        
        cheats.stopPrank();
    }

    // @notice Verifies that `staker` cannot delegate to an unregistered `operator`
    function testCannotDelegateToUnregisteredOperator(address staker, address operator) public 
        filterFuzzedAddressInputs(staker)
        filterFuzzedAddressInputs(operator)
    {
        require(!delegationManager.isOperator(operator), "incorrect test input?");

        // try to delegate and check that the call reverts
        cheats.startPrank(staker);
        cheats.expectRevert(bytes("DelegationManager._delegate: operator is not registered in EigenLayer"));
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);        
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
    function testDelegateToOperatorWhoRequiresECDSASignature(address staker, bytes32 salt, uint256 expiry) public 
        filterFuzzedAddressInputs(staker)
    {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);

        address delegationApprover = cheats.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // verify that the salt hasn't been used before
        require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt), "salt somehow spent too early?");
        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(delegationSignerPrivateKey, staker, operator, salt, expiry);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, operator);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);        
        cheats.stopPrank();

        require(delegationManager.isDelegated(staker), "staker not delegated correctly");
        require(delegationManager.delegatedTo(staker) == operator, "staker delegated to the wrong address");
        require(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        if (staker == operator || staker == delegationManager.delegationApprover(operator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt), "salt somehow spent too incorrectly?");
        } else {
            // verify that the salt is marked as used
            require(delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt), "salt somehow spent not spent?");
        }
    }

    /**
     * @notice Like `testDelegateToOperatorWhoRequiresECDSASignature` but using an incorrect signature on purpose and checking that reversion occurs
     */
    function testDelegateToOperatorWhoRequiresECDSASignature_RevertsWithBadSignature(address staker, uint256 expiry)  public 
        filterFuzzedAddressInputs(staker)
    {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);

        address delegationApprover = cheats.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // calculate the signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        approverSignatureAndExpiry.expiry = expiry;
        {
            bytes32 digestHash =
                delegationManager.calculateDelegationApprovalDigestHash(staker, operator, delegationManager.delegationApprover(operator), emptySalt, expiry);
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(delegationSignerPrivateKey, digestHash);
            // mess up the signature by flipping v's parity
            v = (v == 27 ? 28 : 27);
            approverSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }

        // try to delegate from the `staker` to the operator, and check reversion
        cheats.startPrank(staker);
        cheats.expectRevert(bytes("EIP1271SignatureUtils.checkSignature_EIP1271: signature not from signer"));
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);        
        cheats.stopPrank();
    }

    /**
     * @notice Like `testDelegateToOperatorWhoRequiresECDSASignature` but using an invalid expiry on purpose and checking that reversion occurs
     */
    function testDelegateToOperatorWhoRequiresECDSASignature_RevertsWithExpiredDelegationApproverSignature(address staker, bytes32 salt, uint256 expiry)  public 
        filterFuzzedAddressInputs(staker)
    {
        // roll to a very late timestamp
        cheats.roll(type(uint256).max / 2);
        // filter to only *invalid* `expiry` values
        cheats.assume(expiry < block.timestamp);

        address delegationApprover = cheats.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(delegationSignerPrivateKey, staker, operator, salt, expiry);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectRevert(bytes("DelegationManager._delegate: approver signature expired"));
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);        
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
    function testDelegateToOperatorWhoRequiresEIP1271Signature(address staker, bytes32 salt, uint256 expiry) public 
        filterFuzzedAddressInputs(staker)
    {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);

        address delegationSigner = cheats.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != operator);

        /**
         * deploy a ERC1271WalletMock contract with the `delegationSigner` address as the owner,
         * so that we can create valid signatures from the `delegationSigner` for the contract to check when called
         */
        ERC1271WalletMock wallet = new ERC1271WalletMock(delegationSigner);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(wallet),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // verify that the salt hasn't been used before
        require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt), "salt somehow spent too early?");
        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(delegationSignerPrivateKey, staker, operator, salt, expiry);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, operator);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);        
        cheats.stopPrank();

        require(delegationManager.isDelegated(staker), "staker not delegated correctly");
        require(delegationManager.delegatedTo(staker) == operator, "staker delegated to the wrong address");
        require(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // check that the nonce incremented appropriately
        if (staker == operator || staker == delegationManager.delegationApprover(operator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt), "salt somehow spent too incorrectly?");
        } else {
            // verify that the salt is marked as used
            require(delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt), "salt somehow spent not spent?");
        }
    }

    /**
     * @notice Like `testDelegateToOperatorWhoRequiresEIP1271Signature` but using a contract that
     * returns a value other than the EIP1271 "magic bytes" and checking that reversion occurs appropriately
     */
    function testDelegateToOperatorWhoRequiresEIP1271Signature_RevertsOnBadReturnValue(address staker, uint256 expiry) public 
        filterFuzzedAddressInputs(staker)
    {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != operator);

        // deploy a ERC1271MaliciousMock contract that will return an incorrect value when called
        ERC1271MaliciousMock wallet = new ERC1271MaliciousMock();

        // filter fuzzed input, since otherwise we can get a flaky failure here. if the caller itself is the 'delegationApprover'
        // then we don't even trigger the signature verification call, so we won't get a revert as expected
        cheats.assume(staker != address(wallet));

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(wallet),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // create the signature struct
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        approverSignatureAndExpiry.expiry = expiry;

        // try to delegate from the `staker` to the operator, and check reversion
        cheats.startPrank(staker);
        // because the ERC1271MaliciousMock contract returns the wrong amount of data, we get a low-level "EvmError: Revert" message here rather than the error message bubbling up
        // cheats.expectRevert(bytes("EIP1271SignatureUtils.checkSignature_EIP1271: ERC1271 signature verification failed"));
        cheats.expectRevert();
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);        
        cheats.stopPrank();
    }

    /**
     * @notice `staker` becomes delegated to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * via the `caller` calling `DelegationManager.delegateToBySignature`
     * The function should pass with any `operatorSignature` input (since it should be unused)
     * The function should pass only with a valid `stakerSignatureAndExpiry` input
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testDelegateBySignatureToOperatorWhoAcceptsAllStakers(address caller, bytes32 salt, uint256 expiry) public 
        filterFuzzedAddressInputs(caller)
    {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);

        address staker = cheats.addr(stakerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // verify that the salt hasn't been used before
        require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt), "salt somehow spent too early?");
        // fetch the staker's current nonce
        uint256 currentStakerNonce = delegationManager.stakerNonce(staker);
        // calculate the staker signature
        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(stakerPrivateKey, operator, expiry);

        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        cheats.startPrank(caller);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, operator);
        // use an empty approver signature input since none is needed / the input is unchecked
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateToBySignature(staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, emptySalt);        
        cheats.stopPrank();

        // check all the delegation status changes
        require(delegationManager.isDelegated(staker), "staker not delegated correctly");
        require(delegationManager.delegatedTo(staker) == operator, "staker delegated to the wrong address");
        require(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // check that the staker nonce incremented appropriately
        require(delegationManager.stakerNonce(staker) == currentStakerNonce + 1,
            "staker nonce did not increment");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt), "salt somehow spent too incorrectly?");
    }

    /**
     * @notice `staker` becomes delegated to an operator who requires signature verification through an EOA (i.e. the operator’s `delegationApprover` address is set to a nonzero EOA)
     * via the `caller` calling `DelegationManager.delegateToBySignature`
     * The function should pass *only with a valid ECDSA signature from the `delegationApprover`, OR if called by the operator or their delegationApprover themselves
     * AND with a valid `stakerSignatureAndExpiry` input
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testDelegateBySignatureToOperatorWhoRequiresECDSASignature(address caller, bytes32 salt, uint256 expiry) public 
        filterFuzzedAddressInputs(caller)
    {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);

        address staker = cheats.addr(stakerPrivateKey);
        address delegationApprover = cheats.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // verify that the salt hasn't been used before
        require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt), "salt somehow spent too early?");
        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(delegationSignerPrivateKey, staker, operator, salt, expiry);

        // fetch the staker's current nonce
        uint256 currentStakerNonce = delegationManager.stakerNonce(staker);
        // calculate the staker signature
        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(stakerPrivateKey, operator, expiry);

        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        cheats.startPrank(caller);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, operator);
        delegationManager.delegateToBySignature(staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, salt);        
        cheats.stopPrank();

        require(delegationManager.isDelegated(staker), "staker not delegated correctly");
        require(delegationManager.delegatedTo(staker) == operator, "staker delegated to the wrong address");
        require(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // check that the delegationApprover nonce incremented appropriately
        if (caller == operator || caller == delegationManager.delegationApprover(operator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt), "salt somehow spent too incorrectly?");
        } else {
            // verify that the salt is marked as used
            require(delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt), "salt somehow spent not spent?");
        }

        // check that the staker nonce incremented appropriately
        require(delegationManager.stakerNonce(staker) == currentStakerNonce + 1,
            "staker nonce did not increment");
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
    function testDelegateBySignatureToOperatorWhoRequiresEIP1271Signature(address caller, bytes32 salt, uint256 expiry) public 
        filterFuzzedAddressInputs(caller)
    {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);

        address staker = cheats.addr(stakerPrivateKey);
        address delegationSigner = cheats.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != defaultOperator);

        /**
         * deploy a ERC1271WalletMock contract with the `delegationSigner` address as the owner,
         * so that we can create valid signatures from the `delegationSigner` for the contract to check when called
         */
        ERC1271WalletMock wallet = new ERC1271WalletMock(delegationSigner);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: defaultOperator,
            delegationApprover: address(wallet),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(defaultOperator, operatorDetails, emptyStringForMetadataURI);

        // verify that the salt hasn't been used before
        require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt), "salt somehow spent too early?");
        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(delegationSignerPrivateKey, staker, defaultOperator, salt, expiry);

        // fetch the staker's current nonce
        uint256 currentStakerNonce = delegationManager.stakerNonce(staker);
        // calculate the staker signature
        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(stakerPrivateKey, defaultOperator, expiry);

        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        cheats.startPrank(caller);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, defaultOperator);
        delegationManager.delegateToBySignature(staker, defaultOperator, stakerSignatureAndExpiry, approverSignatureAndExpiry, salt);        
        cheats.stopPrank();

        require(delegationManager.isDelegated(staker), "staker not delegated correctly");
        require(delegationManager.delegatedTo(staker) == defaultOperator, "staker delegated to the wrong address");
        require(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // check that the delegationApprover nonce incremented appropriately
        if (caller == defaultOperator || caller == delegationManager.delegationApprover(defaultOperator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt), "salt somehow spent too incorrectly?");
        } else {
            // verify that the salt is marked as used
            require(delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(defaultOperator), salt), "salt somehow spent not spent?");
        }

        // check that the staker nonce incremented appropriately
        require(delegationManager.stakerNonce(staker) == currentStakerNonce + 1,
            "staker nonce did not increment");
    }

    // @notice Checks that `DelegationManager.delegateToBySignature` reverts if the staker's signature has expired
    function testDelegateBySignatureRevertsWhenStakerSignatureExpired(address staker, address operator, uint256 expiry, bytes memory signature) public{
        cheats.assume(expiry < block.timestamp);
        cheats.expectRevert(bytes("DelegationManager.delegateToBySignature: staker signature expired"));
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry = ISignatureUtils.SignatureWithExpiry({
            signature: signature,
            expiry: expiry
        });
        delegationManager.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, emptySalt);
    }

    // @notice Checks that `DelegationManager.delegateToBySignature` reverts if the delegationApprover's signature has expired and their signature is checked
    function testDelegateBySignatureRevertsWhenDelegationApproverSignatureExpired(address caller, uint256 stakerExpiry, uint256 delegationApproverExpiry) public 
        filterFuzzedAddressInputs(caller)
    {
        // filter to only valid `stakerExpiry` values
        cheats.assume(stakerExpiry >= block.timestamp);
        // roll to a very late timestamp
        cheats.roll(type(uint256).max / 2);
        // filter to only *invalid* `delegationApproverExpiry` values
        cheats.assume(delegationApproverExpiry < block.timestamp);

        address staker = cheats.addr(stakerPrivateKey);
        address delegationApprover = cheats.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, operator, emptySalt, delegationApproverExpiry);

        // calculate the staker signature
        ISignatureUtils.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(stakerPrivateKey, operator, stakerExpiry);

        // try delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`, and check for reversion
        cheats.startPrank(caller);
        cheats.expectRevert(bytes("DelegationManager._delegate: approver signature expired"));
        delegationManager.delegateToBySignature(staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, emptySalt);        
        cheats.stopPrank();
    }

    /**
     * @notice Like `testDelegateToOperatorWhoRequiresECDSASignature` but using an invalid expiry on purpose and checking that reversion occurs
     */
    function testDelegateToOperatorWhoRequiresECDSASignature_RevertsWithExpiredSignature(address staker, uint256 expiry)  public 
        filterFuzzedAddressInputs(staker)
    {
        // roll to a very late timestamp
        cheats.roll(type(uint256).max / 2);
        // filter to only *invalid* `expiry` values
        cheats.assume(expiry < block.timestamp);

        address delegationApprover = cheats.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, operator, emptySalt, expiry);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectRevert(bytes("DelegationManager._delegate: approver signature expired"));
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);        
        cheats.stopPrank();
    }

    /**
     * Staker is undelegated from an operator, via a call to `undelegate`, properly originating from the staker's address.
     * Reverts if the staker is themselves an operator (i.e. they are delegated to themselves)
     * Does nothing if the staker is already undelegated
     * Properly undelegates the staker, i.e. the staker becomes “delegated to” the zero address, and `isDelegated(staker)` returns ‘false’
     * Emits a `StakerUndelegated` event
     */
    function testUndelegateFromOperator(address staker) public {
        // register *this contract* as an operator and delegate from the `staker` to them (already filters out case when staker is the operator since it will revert)
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        testDelegateToOperatorWhoAcceptsAllStakers(staker, approverSignatureAndExpiry, emptySalt);

        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(staker, delegationManager.delegatedTo(staker));
        delegationManager.undelegate(staker);
        cheats.stopPrank();

        require(!delegationManager.isDelegated(staker), "staker not undelegated!");
        require(delegationManager.delegatedTo(staker) == address(0), "undelegated staker should be delegated to zero address");
    }

    // @notice Verifies that an operator cannot undelegate from themself (this should always be forbidden)
    function testOperatorCannotUndelegateFromThemself(address operator) public filterFuzzedAddressInputs(operator) {
        cheats.startPrank(operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        cheats.stopPrank();
        cheats.expectRevert(bytes("DelegationManager.undelegate: operators cannot be undelegated"));
        
        cheats.startPrank(operator);
        delegationManager.undelegate(operator);
        cheats.stopPrank();
    }

    /**
     * @notice Verifies that `DelegationManager.increaseDelegatedShares` properly increases the delegated `shares` that the operator
     * who the `staker` is delegated to has in the strategy
     * @dev Checks that there is no change if the staker is not delegated
     */
    function testIncreaseDelegatedShares(address staker, uint256 shares, bool delegateFromStakerToOperator) public {
        IStrategy strategy = strategyMock;

        // register *this contract* as an operator
        address operator = address(this);
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        // filter inputs, since delegating to the operator will fail when the staker is already registered as an operator
        cheats.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // delegate from the `staker` to the operator *if `delegateFromStakerToOperator` is 'true'*
        if (delegateFromStakerToOperator) {
            cheats.startPrank(staker);
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit StakerDelegated(staker, operator);
            delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);        
            cheats.stopPrank();            
        }

        uint256 _delegatedSharesBefore = delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategy);

        if(delegationManager.isDelegated(staker)) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorSharesIncreased(operator, staker, strategy, shares);
        }

        cheats.startPrank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, strategy, shares);
        cheats.stopPrank();

        uint256 delegatedSharesAfter = delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategy);

        if (delegationManager.isDelegated(staker)) {
            require(delegatedSharesAfter == _delegatedSharesBefore + shares, "delegated shares did not increment correctly");
        } else {
            require(delegatedSharesAfter == _delegatedSharesBefore, "delegated shares incremented incorrectly");
            require(_delegatedSharesBefore == 0, "nonzero shares delegated to zero address!");
        }
    }

    /**
     * @notice Verifies that `DelegationManager.decreaseDelegatedShares` properly decreases the delegated `shares` that the operator
     * who the `staker` is delegated to has in the strategies
     * @dev Checks that there is no change if the staker is not delegated
     */
    function testDecreaseDelegatedShares(address staker, IStrategy[] memory strategies, uint128 shares, bool delegateFromStakerToOperator) public filterFuzzedAddressInputs(staker) {
        // sanity-filtering on fuzzed input length
        cheats.assume(strategies.length <= 32);
        // register *this contract* as an operator
        address operator = address(this);
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
        // filter inputs, since delegating to the operator will fail when the staker is already registered as an operator
        cheats.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // delegate from the `staker` to the operator *if `delegateFromStakerToOperator` is 'true'*
        if (delegateFromStakerToOperator) {
            cheats.startPrank(staker);
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit StakerDelegated(staker, operator);
            delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);        
            cheats.stopPrank();            
        }

        uint256[] memory sharesInputArray = new uint256[](strategies.length);

        address delegatedTo = delegationManager.delegatedTo(staker);

        // for each strategy in `strategies`, increase delegated shares by `shares`
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

        // for each strategy in `strategies`, decrease delegated shares by `shares`
        {
            cheats.startPrank(address(strategyManagerMock));
            address operatorToDecreaseSharesOf = delegationManager.delegatedTo(staker);
            if (delegationManager.isDelegated(staker)) {
                for (uint256 i = 0; i < strategies.length;  ++i) {
                    cheats.expectEmit(true, true, true, true, address(delegationManager));
                    emit OperatorSharesDecreased(operatorToDecreaseSharesOf, staker, strategies[i], sharesInputArray[i]);
                    delegationManager.decreaseDelegatedShares(staker, strategies[i], sharesInputArray[i]);
                }
            }
            cheats.stopPrank();
        }

        // check shares after call to `decreaseDelegatedShares`
        bool isDelegated =  delegationManager.isDelegated(staker);
        for (uint256 i = 0; i < strategies.length; ++i) {
            uint256 delegatedSharesAfter = delegationManager.operatorShares(delegatedTo, strategies[i]); 

            if (isDelegated) {
                require(delegatedSharesAfter + totalSharesForStrategyInArray[address(strategies[i])] == delegatedSharesBefore[strategies[i]],
                    "delegated shares did not decrement correctly");
            } else {
                require(delegatedSharesAfter == delegatedSharesBefore[strategies[i]], "delegated shares decremented incorrectly");
                require(delegatedSharesBefore[strategies[i]] == 0, "nonzero shares delegated to zero address!");
            }
        }
    }

    // @notice Verifies that `DelegationManager.increaseDelegatedShares` reverts if not called by the StrategyManager nor EigenPodManager
    function testCannotCallIncreaseDelegatedSharesFromNonPermissionedAddress(address operator, uint256 shares) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(strategyManagerMock));
        cheats.assume(operator != address(eigenPodManagerMock));
        cheats.expectRevert(bytes("DelegationManager: onlyStrategyManagerOrEigenPodManager"));
        cheats.startPrank(operator);
        delegationManager.increaseDelegatedShares(operator, strategyMock, shares);
    }

    // @notice Verifies that `DelegationManager.decreaseDelegatedShares` reverts if not called by the StrategyManager nor EigenPodManager
    function testCannotCallDecreaseDelegatedSharesFromNonPermissionedAddress(
        address operator,  
        IStrategy strategy,  
        uint256 shares
    ) public filterFuzzedAddressInputs(operator) {
        cheats.assume(operator != address(strategyManagerMock));
        cheats.assume(operator != address(eigenPodManagerMock));
        cheats.expectRevert(bytes("DelegationManager: onlyStrategyManagerOrEigenPodManager"));
        cheats.startPrank(operator);
        delegationManager.decreaseDelegatedShares(operator, strategy, shares);
    }

    // @notice Verifies that it is not possible for a staker to delegate to an operator when they are already delegated to an operator
    function testCannotDelegateWhenStakerHasExistingDelegation(address staker, address operator, address operator2) public
        filterFuzzedAddressInputs(staker)
        filterFuzzedAddressInputs(operator)
        filterFuzzedAddressInputs(operator2)
    {
        cheats.assume(operator != operator2);
        cheats.assume(staker != operator);
        cheats.assume(staker != operator2);

        cheats.startPrank(operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        cheats.stopPrank();

        cheats.startPrank(operator2);
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        cheats.stopPrank();

        cheats.startPrank(staker);
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry, emptySalt);
        cheats.stopPrank();

        cheats.startPrank(staker);
        cheats.expectRevert(bytes("DelegationManager._delegate: staker is already actively delegated"));
        delegationManager.delegateTo(operator2, signatureWithExpiry, emptySalt);
        cheats.stopPrank();
    }

    // @notice Verifies that it is not possible to delegate to an unregistered operator
    function testCannotDelegateToUnregisteredOperator(address operator) public {
        cheats.expectRevert(bytes("DelegationManager._delegate: operator is not registered in EigenLayer"));
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry, emptySalt);
    }

    // @notice Verifies that delegating is not possible when the "new delegations paused" switch is flipped
    function testCannotDelegateWhenPausedNewDelegationIsSet(address operator, address staker) public filterFuzzedAddressInputs(operator) filterFuzzedAddressInputs(staker) {
        // set the pausing flag
        cheats.startPrank(pauser);
        delegationManager.pause(2 ** PAUSED_NEW_DELEGATION);
        cheats.stopPrank();

        cheats.startPrank(staker);
        cheats.expectRevert(bytes("Pausable: index is paused"));
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry, emptySalt);
        cheats.stopPrank();
    }

    // @notice Verifies that undelegating is not possible when the "undelegation paused" switch is flipped
    function testCannotUndelegateWhenPausedUndelegationIsSet(address operator, address staker) public filterFuzzedAddressInputs(operator) filterFuzzedAddressInputs(staker) {
        // register *this contract* as an operator and delegate from the `staker` to them (already filters out case when staker is the operator since it will revert)
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        testDelegateToOperatorWhoAcceptsAllStakers(staker, approverSignatureAndExpiry, emptySalt);

        // set the pausing flag
        cheats.startPrank(pauser);
        delegationManager.pause(2 ** PAUSED_ENTER_WITHDRAWAL_QUEUE);
        cheats.stopPrank();

        cheats.startPrank(staker);
        cheats.expectRevert(bytes("Pausable: index is paused"));
        delegationManager.undelegate(staker);
        cheats.stopPrank();
    }

    // special event purely used in the StrategyManagerMock contract, inside of `undelegate` function to verify that the correct call is made
    event ForceTotalWithdrawalCalled(address staker);

    /**
     * @notice Verifies that the `undelegate` function properly calls `strategyManager.forceTotalWithdrawal` when necessary
     * @param callFromOperatorOrApprover -- calls from the operator if 'false' and the 'approver' if true
     */
    function testForceUndelegation(address staker, bytes32 salt, bool callFromOperatorOrApprover) public
        filterFuzzedAddressInputs(staker)
    {
        address delegationApprover = cheats.addr(delegationSignerPrivateKey);
        address operator = address(this);

        // filtering since you can't delegate to yourself after registering as an operator
        cheats.assume(staker != operator);

        // register this contract as an operator and delegate from the staker to it
        uint256 expiry = type(uint256).max;
        testDelegateToOperatorWhoRequiresECDSASignature(staker, salt, expiry);

        address caller;
        if (callFromOperatorOrApprover) {
            caller = delegationApprover;
        } else {
            caller = operator;
        }

        // call the `undelegate` function
        cheats.startPrank(caller);
        // check that the correct calldata is forwarded by looking for an event emitted by the StrategyManagerMock contract
        if (strategyManagerMock.stakerStrategyListLength(staker) != 0) {
            cheats.expectEmit(true, true, true, true, address(strategyManagerMock));
            emit ForceTotalWithdrawalCalled(staker);
        }
        // withdrawal root
        (IStrategy[] memory strategies, uint256[] memory shares) = delegationManager.getDelegatableShares(staker);
        IDelegationManager.Withdrawal memory fullWithdrawal = IDelegationManager.Withdrawal({
            staker: staker,
            delegatedTo: operator,
            withdrawer: staker,
            nonce: delegationManager.cumulativeWithdrawalsQueued(staker),
            startBlock: uint32(block.number),
            strategies: strategies,
            shares: shares
        });
        bytes32 withdrawalRoot = delegationManager.calculateWithdrawalRoot(fullWithdrawal);

        (bytes32 returnValue) = delegationManager.undelegate(staker);

        if (strategies.length == 0) {
            withdrawalRoot = bytes32(0);
        }

        // check that the return value is the withdrawal root
        require(returnValue == withdrawalRoot, "contract returned wrong return value");
        cheats.stopPrank();
    }

    /**
     * @notice Verifies that the `undelegate` function has proper access controls (can only be called by the operator who the `staker` has delegated
     * to or the operator's `delegationApprover`), or the staker themselves
     */
    function testCannotCallUndelegateFromImproperAddress(address staker, address caller) public
        filterFuzzedAddressInputs(staker)
        filterFuzzedAddressInputs(caller)
    {
        address delegationApprover = cheats.addr(delegationSignerPrivateKey);
        address operator = address(this);

        // filtering since you can't delegate to yourself after registering as an operator
        cheats.assume(staker != operator);

        // filter out addresses that are actually allowed to call the function
        cheats.assume(caller != operator);
        cheats.assume(caller != delegationApprover);
        cheats.assume(caller != staker);

        // register this contract as an operator and delegate from the staker to it
        uint256 expiry = type(uint256).max;
        testDelegateToOperatorWhoRequiresECDSASignature(staker, emptySalt, expiry);

        // try to call the `undelegate` function and check for reversion
        cheats.startPrank(caller);
        cheats.expectRevert(bytes("DelegationManager.undelegate: caller cannot undelegate staker"));
        delegationManager.undelegate(staker);
        cheats.stopPrank();
    }

    /**
     * @notice verifies that `DelegationManager.undelegate` reverts if trying to undelegate an operator from themselves
     * @param callFromOperatorOrApprover -- calls from the operator if 'false' and the 'approver' if true
     */
    function testOperatorCannotForceUndelegateThemself(address delegationApprover, bool callFromOperatorOrApprover) public {
        // register *this contract* as an operator
        address operator = address(this);
        IDelegationManager.OperatorDetails memory _operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, _operatorDetails, emptyStringForMetadataURI);

        address caller;
        if (callFromOperatorOrApprover) {
            caller = delegationApprover;
        } else {
            caller = operator;
        }

        // try to call the `undelegate` function and check for reversion
        cheats.startPrank(caller);
        cheats.expectRevert(bytes("DelegationManager.undelegate: operators cannot be undelegated"));
        delegationManager.undelegate(operator);
        cheats.stopPrank();
    }

    /**
     * @notice Verifies that the reversion occurs when trying to reuse an 'approverSalt'
     */
    function test_Revert_WhenTryingToReuseSalt(address staker_one, address staker_two, bytes32 salt) public
        filterFuzzedAddressInputs(staker_one)
        filterFuzzedAddressInputs(staker_two)
    {
        // address delegationApprover = cheats.addr(delegationSignerPrivateKey);
        address operator = address(this);

        // filtering since you can't delegate to yourself after registering as an operator
        cheats.assume(staker_one != operator);
        cheats.assume(staker_two != operator);

        // filtering since you can't delegate twice
        cheats.assume(staker_one != staker_two);

        address delegationApprover = cheats.addr(delegationSignerPrivateKey);
        // filter out the case where `staker` *is* the 'delegationApprover', since in this case the salt won't get used
        cheats.assume(staker_one != delegationApprover);
        cheats.assume(staker_two != delegationApprover);

        // register this contract as an operator and delegate from `staker_one` to it, using the `salt`
        uint256 expiry = type(uint256).max;
        testDelegateToOperatorWhoRequiresECDSASignature(staker_one, salt, expiry);

        // calculate the delegationSigner's signature
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker_two, operator, salt, expiry);

        // try to delegate to the operator from `staker_two`, and verify that the call reverts for the proper reason (trying to reuse a salt)
        cheats.startPrank(staker_two);
        cheats.expectRevert(bytes("DelegationManager._delegate: approverSalt already spent"));
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
    }
}
