// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import {ERC1271WalletMock, ERC1271MaliciousMock} from "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {DelegationManager} from "src/contracts/core/DelegationManager.sol";
import {IDelegationManager} from "src/contracts/interfaces/IDelegationManager.sol";
import {IStrategy} from "src/contracts/interfaces/IStrategy.sol";
import {StrategyBase} from "src/contracts/strategies/StrategyBase.sol";
import {EigenLayerTestHelper, EigenLayerDeployer} from "src/test/EigenLayerTestHelper.t.sol";
import {StrategyManagerMock} from "src/test/mocks/StrategyManagerMock.sol";
import {SlasherMock} from "src/test/mocks/SlasherMock.sol";
import {ERC20Mock} from "src/test/mocks/ERC20Mock.sol";

contract DelegationUnitTests is EigenLayerTestHelper {
    StrategyManagerMock internal strategyManagerMock;
    SlasherMock internal slasherMock;
    DelegationManager internal delegationManager;
    DelegationManager internal delegationManagerImplementation;
    StrategyBase internal strategyImplementation;
    StrategyBase internal strategyMock;
    bytes internal constant ALREADY_INITIALIZED_ERROR = bytes("Initializable: contract is already initialized");

    bytes internal constant STAKER_OPT_OUT_WINDOW_GT_MAX_ERROR = bytes(
        "DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be > MAX_STAKER_OPT_OUT_WINDOW_BLOCKS"
    );
    bytes internal constant EARNINGS_RECEIVER_ADDRESS_ZERO_ERROR =
        bytes("DelegationManager._setOperatorDetails: cannot set `earningsReceiver` to zero address");
    bytes internal constant OPERATOR_ALREADY_REGISTER_ERROR =
        bytes("DelegationManager.registerAsOperator: operator has already registered");
    bytes internal constant STAKER_ALREADY_DELEGATED_ERROR =
        bytes("DelegationManager._delegate: staker is already actively delegated");
    bytes internal constant OPT_OUT_WINDOW_INCREASE_ERROR =
        bytes("DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be decreased");
    bytes internal constant NOT_OPERATOR_ERROR =
        bytes("DelegationManager.updateOperatorMetadataURI: caller must be an operator");
    bytes internal constant OPERATOR_NOT_REGISTERED_ERROR =
        bytes("DelegationManager._delegate: operator is not registered in EigenLayer");
    bytes internal constant INVALID_SIGNATURE_ERROR =
        bytes("EIP1271SignatureUtils.checkSignature_EIP1271: signature not from signer");
    bytes internal constant APPROVER_SIGNATURE_EXPIRED_ERROR =
        bytes("DelegationManager._delegate: approver signature expired");
    bytes internal constant OPERATOR_UNDELEGATE_SELF_ERROR =
        bytes("DelegationManager.undelegate: operators cannot undelegate from themselves");
    bytes internal constant STAKER_SIGNATURE_EXPIRED_ERROR =
        bytes("DelegationManager.delegateToBySignature: staker signature expired");
    bytes internal constant NOT_STRATEGY_MANAGER_ERROR = bytes("onlyStrategyManager");
    bytes internal constant OPERATOR_FROZEN_ERROR =
        bytes("DelegationManager._delegate: cannot delegate to a frozen operator");
    bytes internal constant PAUSED_INDEX_ERROR = bytes("Pausable: index is paused");
    bytes internal constant NOT_OPERATOR_OR_APPOVER_ERROR =
        bytes("DelegationManager.forceUndelegation: caller must be operator or their delegationApprover");
    bytes internal constant FORCE_UNDELEGATE_OPERATOR_ERROR =
        bytes("DelegationManager.forceUndelegation: operators cannot be force-undelegated");
    bytes internal constant APPROVER_SALT_SPENT_ERROR = bytes("DelegationManager._delegate: approverSalt already spent");
    // mapping for filtering out weird fuzzed inputs, like making calls from the ProxyAdmin or the zero address
    mapping(address => bool) public addressIsExcludedFromFuzzedInputs;

    uint256 internal delegationSignerPrivateKey =
        uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);
    uint256 internal stakerPrivateKey = uint256(123456789);

    // empty string reused across many tests
    string internal emptyStringForMetadataURI;

    // "empty" / zero salt, reused across many tests
    bytes32 internal emptySalt;

    // reused in various tests. in storage to help handle stack-too-deep errors
    address internal _operator = address(this);

    // @notice Emitted when a new operator registers in EigenLayer and provides their OperatorDetails.
    event OperatorRegistered(address indexed operator, IDelegationManager.OperatorDetails operatorDetails);

    // @notice Emitted when an operator updates their OperatorDetails to @param newOperatorDetails
    event OperatorDetailsModified(address indexed operator, IDelegationManager.OperatorDetails newOperatorDetails);

    /**
     * @notice Emitted when @param operator indicates that they are updating their MetadataURI string
     * @dev Note that these strings are *never stored in storage* and are instead purely emitted in events for off-chain indexing
     */
    event OperatorMetadataURIUpdated(address indexed operator, string metadataURI);

    // @notice Emitted when @param staker delegates to @param operator.
    event StakerDelegated(address indexed staker, address indexed operator);

    // @notice Emitted when @param staker undelegates from @param operator.
    event StakerUndelegated(address indexed staker, address indexed operator);

    // @notice reuseable modifier
    modifier filterFuzzedAddressInputs(address fuzzedAddress) {
        vm.assume(!addressIsExcludedFromFuzzedInputs[fuzzedAddress]);
        _;
    }

    function setUp() public virtual override {
        EigenLayerDeployer.setUp();

        slasherMock = new SlasherMock();
        delegationManager = DelegationManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        strategyManagerMock = new StrategyManagerMock();

        delegationManagerImplementation = new DelegationManager(strategyManagerMock, slasherMock);

        vm.startPrank(eigenLayerProxyAdmin.owner());
        eigenLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(delegationManager))), address(delegationManagerImplementation)
        );
        vm.stopPrank();

        address initalOwner = address(this);
        uint256 initialPausedStatus = 0;
        delegationManager.initialize(initalOwner, eigenLayerPauserReg, initialPausedStatus);

        strategyImplementation = new StrategyBase(strategyManagerMock);

        strategyMock = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(strategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, weth, eigenLayerPauserReg)
                )
            )
        );

        // excude the zero address and the proxyAdmin from fuzzed inputs
        addressIsExcludedFromFuzzedInputs[address(0)] = true;
        addressIsExcludedFromFuzzedInputs[address(eigenLayerProxyAdmin)] = true;
        addressIsExcludedFromFuzzedInputs[address(strategyManagerMock)] = true;
        addressIsExcludedFromFuzzedInputs[address(delegationManager)] = true;
        addressIsExcludedFromFuzzedInputs[address(slasherMock)] = true;

        // check setup (constructor + initializer)
        assertTrue(delegationManager.strategyManager() == strategyManagerMock, "strategyManager set wrong");
        assertTrue(delegationManager.slasher() == slasherMock, "slasher set wrong");
        assertTrue(delegationManager.pauserRegistry() == eigenLayerPauserReg, "pauserRegistry set wrong");
        assertTrue(delegationManager.owner() == initalOwner, "owner set wrong");
        assertTrue(delegationManager.paused() == initialPausedStatus, "paused status set wrong");
    }

    // @notice Verifies that the DelegationManager cannot be iniitalized multiple times
    function test_RevertsWhen_AlreadyInitialized_Initialize() public {
        vm.expectRevert(ALREADY_INITIALIZED_ERROR);
        delegationManager.initialize(address(this), eigenLayerPauserReg, 0);
    }

    function test_RegisterAsOperator() public {}

    function test_RevertsWhen_AlreadyOperator_RegisterAsOperator() public {}

    function test_ModifyOperatorDetails() public {}

    function test_RevertsWhen_EarningsReceiverAddress0_ModifyOperatorDetails() public {}

    function test_RevertsWhen_StakerOptOutWindowBlocksGreaterThanMaxOptOutWindowBlocks_ModifyOperatorDetails() public {}

    function test_RevertsWhen_StakerOptOutWindowBlocksDecreased_ModifyOperatorDetails() public {}

    function test_UpdateOperatorMetdataURI() public {}

    function test_RevertsWhen_NotOperator_UpdateOperatorMetdataURI() public {}

    function test_RevertsWhen_InvalidSignature_DelegateTo() public {}

    function test_RevertsWhen_OperatorNotRegistered_DelegateTo() public {}

    function test_RevertsWhen_AlreadyDelegated_DelegateTo() public {}

    function test_DelegateToBySignature() public {}

    function test_RevertsWhen_SignatureExpired_DelegateToBySignature() public {}

    function test_RevertsWhen_SignatureInvalidWrongSigner_DelegateToBySignature() public {}

    function test_RevertsWhen_StakerAlreadyDelegated_DelegateToBySignature() public {}

    function test_RevertsWhen_OperatorNotRegistered_DelegateToBySignature() public {}

    function test_RevertsWhen_OperatorIsFrozen_DelegateToBySignature() public {}

    function test_RevertsWhen_SignatureInvalidWrongNonce_DelegateToBySignature() public {}

    function test_RevertsWhen_DelegateApproverSignatureExpired_DelegateToBySignature() public {}

    function test_RevertsWhen_SignatureSaltPreviouslyUsed_DelegateToBySignature() public {}

    /// TODO: Look more into if we can bypass the ApproveSig check for DelegateToBySignature

    function test_Undelegate() public {}

    function test_RevertsWhen_OperatorUndelegatesFromSelf_Undelegate() public {}

    function test_RevertsWhen_NotStrategyManager_Undelegate() public {}

    function test_ForceUndelegation() public {}
    /// TODO: Come back to this one

    function test_IncreaseDelegatedShares() public {}

    function test_RevertsWhen_NotStrategyManager_IncreaseDelegatedShares() public {}

    function test_StrategyMangerCallsSucceedsWithoutIncreasingShares_IncreaseDelegatedShares() public {}

    function test_DereaseDelegatedShares() public {}

    function test_RevertsWhen_NotStrategyManager_DecreaseDelegatedShares() public {}

    function test_StrategyMangerCallsSucceedsWithoutDecreasingShares_DecreaseDelegatedShares() public {}

    function test_DomainSeperator() public {}

    function test_IsDelegated() public {}

    function test_IsNotDelegated() public {}

    function test_IsOperator() public {}

    function test_IsNotOperator() public {}

    function test_DelegateTo() public {}

    function test_CalculateCurrentStakerDelegationDigestHash() public {}

    function test_CalculateDelegationApprovalDigestHash() public {}

    /**
     * @notice `operator` registers via calling `DelegationManager.registerAsOperator(operatorDetails, metadataURI)`
     * Should be able to set any parameters, other than setting their `earningsReceiver` to the zero address or too high value for `stakerOptOutWindowBlocks`
     * The set parameters should match the desired parameters (correct storage update)
     * Operator becomes delegated to themselves
     * Properly emits events – especially the `OperatorRegistered` event, but also `StakerDelegated` & `OperatorDetailsModified` events
     * Reverts appropriately if operator was already delegated to someone (including themselves, i.e. they were already an operator)
     * @param operator and @param operatorDetails are fuzzed inputs
     */
    function testFuzz_RegisterAsOperator(
        address operator,
        IDelegationManager.OperatorDetails memory operatorDetails,
        string memory metadataURI
    ) public filterFuzzedAddressInputs(operator) {
        // filter out zero address since people can't delegate to the zero address and operators are delegated to themselves
        vm.assume(operator != address(0));
        // filter out zero address since people can't set their earningsReceiver address to the zero address (special test case to verify)
        vm.assume(operatorDetails.earningsReceiver != address(0));
        // filter out disallowed stakerOptOutWindowBlocks values
        vm.assume(operatorDetails.stakerOptOutWindowBlocks <= delegationManager.MAX_STAKER_OPT_OUT_WINDOW_BLOCKS());

        vm.startPrank(operator);
        vm.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorDetailsModified(operator, operatorDetails);
        vm.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(operator, operator);
        vm.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorRegistered(operator, operatorDetails);
        vm.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorMetadataURIUpdated(operator, metadataURI);

        delegationManager.registerAsOperator(operatorDetails, metadataURI);

        assertTrue(
            operatorDetails.earningsReceiver == delegationManager.earningsReceiver(operator),
            "earningsReceiver not set correctly"
        );
        assertTrue(
            operatorDetails.delegationApprover == delegationManager.delegationApprover(operator),
            "delegationApprover not set correctly"
        );
        assertTrue(
            operatorDetails.stakerOptOutWindowBlocks == delegationManager.stakerOptOutWindowBlocks(operator),
            "stakerOptOutWindowBlocks not set correctly"
        );
        assertTrue(delegationManager.delegatedTo(operator) == operator, "operator not delegated to self");
        vm.stopPrank();
    }

    /**
     * @notice Verifies that an operator cannot register with `stakerOptOutWindowBlocks` set larger than `MAX_STAKER_OPT_OUT_WINDOW_BLOCKS`
     * @param operatorDetails is a fuzzed input
     */
    function testFuzz_CannotRegisterAsOperatorWithDisallowedStakerOptOutWindowBlocks(
        IDelegationManager.OperatorDetails memory operatorDetails
    ) public {
        // filter out zero address since people can't set their earningsReceiver address to the zero address (special test case to verify)
        vm.assume(operatorDetails.earningsReceiver != address(0));
        // filter out *allowed* stakerOptOutWindowBlocks values
        vm.assume(operatorDetails.stakerOptOutWindowBlocks > delegationManager.MAX_STAKER_OPT_OUT_WINDOW_BLOCKS());

        vm.startPrank(operator);
        vm.expectRevert(STAKER_OPT_OUT_WINDOW_GT_MAX_ERROR);
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        vm.stopPrank();
    }

    /**
     * @notice Verifies that an operator cannot register with `earningsReceiver` set to the zero address
     * @dev This is an important check since we check `earningsReceiver != address(0)` to check if an address is an operator!
     */
    function test_CannotRegisterAsOperatorWithZeroAddressAsEarningsReceiver() public {
        vm.startPrank(operator);
        IDelegationManager.OperatorDetails memory operatorDetails;
        vm.expectRevert(EARNINGS_RECEIVER_ADDRESS_ZERO_ERROR);
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        vm.stopPrank();
    }

    // @notice Verifies that someone cannot successfully call `DelegationManager.registerAsOperator(operatorDetails)` again after registering for the first time
    function testFuzz_CannotRegisterAsOperatorMultipleTimes(
        address operator,
        IDelegationManager.OperatorDetails memory operatorDetails
    ) public filterFuzzedAddressInputs(operator) {
        testFuzz_RegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);
        vm.startPrank(operator);
        vm.expectRevert(OPERATOR_ALREADY_REGISTER_ERROR);
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        vm.stopPrank();
    }

    // @notice Verifies that a staker who is actively delegated to an operator cannot register as an operator (without first undelegating, at least)
    function testFuzz_CannotRegisterAsOperatorWhileDelegated(
        address staker,
        IDelegationManager.OperatorDetails memory operatorDetails
    ) public filterFuzzedAddressInputs(staker) {
        // filter out disallowed stakerOptOutWindowBlocks values
        vm.assume(operatorDetails.stakerOptOutWindowBlocks <= delegationManager.MAX_STAKER_OPT_OUT_WINDOW_BLOCKS());
        // filter out zero address since people can't set their earningsReceiver address to the zero address (special test case to verify)
        vm.assume(operatorDetails.earningsReceiver != address(0));

        // register *this contract* as an operator
        IDelegationManager.OperatorDetails memory _operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: _operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(_operator, _operatorDetails, emptyStringForMetadataURI);

        // delegate from the `staker` to the operator
        vm.startPrank(staker);
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateTo(_operator, approverSignatureAndExpiry, emptySalt);

        vm.expectRevert(STAKER_ALREADY_DELEGATED_ERROR);
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);

        vm.stopPrank();
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
    function testFuzz_ModifyOperatorParameters(
        IDelegationManager.OperatorDetails memory initialOperatorDetails,
        IDelegationManager.OperatorDetails memory modifiedOperatorDetails
    ) public {
        testFuzz_RegisterAsOperator(_operator, initialOperatorDetails, emptyStringForMetadataURI);
        // filter out zero address since people can't set their earningsReceiver address to the zero address (special test case to verify)
        vm.assume(modifiedOperatorDetails.earningsReceiver != address(0));

        vm.startPrank(_operator);

        // either it fails for trying to set the stakerOptOutWindowBlocks
        if (modifiedOperatorDetails.stakerOptOutWindowBlocks > delegationManager.MAX_STAKER_OPT_OUT_WINDOW_BLOCKS()) {
            vm.expectRevert(STAKER_OPT_OUT_WINDOW_GT_MAX_ERROR);
            delegationManager.modifyOperatorDetails(modifiedOperatorDetails);
            // or the transition is allowed,
        } else if (modifiedOperatorDetails.stakerOptOutWindowBlocks >= initialOperatorDetails.stakerOptOutWindowBlocks)
        {
            vm.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorDetailsModified(_operator, modifiedOperatorDetails);
            delegationManager.modifyOperatorDetails(modifiedOperatorDetails);

            assertTrue(
                modifiedOperatorDetails.earningsReceiver == delegationManager.earningsReceiver(_operator),
                "earningsReceiver not set correctly"
            );
            assertTrue(
                modifiedOperatorDetails.delegationApprover == delegationManager.delegationApprover(_operator),
                "delegationApprover not set correctly"
            );
            assertTrue(
                modifiedOperatorDetails.stakerOptOutWindowBlocks
                    == delegationManager.stakerOptOutWindowBlocks(_operator),
                "stakerOptOutWindowBlocks not set correctly"
            );
            assertTrue(delegationManager.delegatedTo(_operator) == _operator, "operator not delegated to self");
            // or else the transition is disallowed
        } else {
            vm.expectRevert(OPT_OUT_WINDOW_INCREASE_ERROR);
            delegationManager.modifyOperatorDetails(modifiedOperatorDetails);
        }

        vm.stopPrank();
    }

    // @notice Tests that an operator who calls `updateOperatorMetadataURI` will correctly see an `OperatorMetadataURIUpdated` event emitted with their input
    function testFuzz_UpdateOperatorMetadataURI(string memory metadataURI) public {
        // register *this contract* as an operator
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: _operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(_operator, operatorDetails, emptyStringForMetadataURI);

        // call `updateOperatorMetadataURI` and check for event
        vm.startPrank(_operator);
        vm.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorMetadataURIUpdated(_operator, metadataURI);
        delegationManager.updateOperatorMetadataURI(metadataURI);
        vm.stopPrank();
    }

    // @notice Tests that an address which is not an operator cannot successfully call `updateOperatorMetadataURI`.
    function test_CannotUpdateOperatorMetadataURIWithoutRegisteringFirst() public {
        assertTrue(!delegationManager.isOperator(_operator), "bad test setup");

        vm.prank(_operator);
        vm.expectRevert(NOT_OPERATOR_ERROR);
        delegationManager.updateOperatorMetadataURI(emptyStringForMetadataURI);
    }

    /**
     * @notice Verifies that an operator cannot modify their `earningsReceiver` address to set it to the zero address
     * @dev This is an important check since we check `earningsReceiver != address(0)` to check if an address is an operator!
     */
    function test_CannotModifyEarningsReceiverAddressToZeroAddress() public {
        // register *this contract* as an operator
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: _operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(_operator, operatorDetails, emptyStringForMetadataURI);

        operatorDetails.earningsReceiver = address(0);
        vm.expectRevert(EARNINGS_RECEIVER_ADDRESS_ZERO_ERROR);
        delegationManager.modifyOperatorDetails(operatorDetails);
    }

    /**
     * @notice `staker` delegates to an operator who does not assertTrue any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * via the `staker` calling `DelegationManager.delegateTo`
     * The function should pass with any `operatorSignature` input (since it should be unused)
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_DelegateToOperatorWhoAcceptsAllStakers(
        address staker,
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 salt
    ) public filterFuzzedAddressInputs(staker) {
        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        vm.assume(staker != _operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: _operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(_operator, operatorDetails, emptyStringForMetadataURI);

        // verify that the salt hasn't been used before
        assertTrue(
            !delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(_operator), salt),
            "salt somehow spent too early?"
        );

        // delegate from the `staker` to the operator
        vm.startPrank(staker);
        vm.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, _operator);
        delegationManager.delegateTo(_operator, approverSignatureAndExpiry, salt);
        vm.stopPrank();

        assertTrue(delegationManager.isDelegated(staker), "staker not delegated correctly");
        assertTrue(delegationManager.delegatedTo(staker) == _operator, "staker delegated to the wrong address");
        assertTrue(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertTrue(
            !delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(_operator), salt),
            "salt somehow spent too early?"
        );
    }

    /**
     * @notice Delegates from `staker` to an operator, then verifies that the `staker` cannot delegate to another `operator` (at least without first undelegating)
     */
    function testFuzz_CannotDelegateWhileDelegated(
        address staker,
        address operator,
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry,
        bytes32 salt
    ) public filterFuzzedAddressInputs(staker) filterFuzzedAddressInputs(operator) {
        // filter out input since if the staker tries to delegate again after registering as an operator, we will revert earlier than this test is designed to check
        vm.assume(staker != operator);

        // delegate from the staker to an operator
        testFuzz_DelegateToOperatorWhoAcceptsAllStakers(staker, approverSignatureAndExpiry, salt);

        // register another operator
        // filter out this contract, since we already register it as an operator in the above step
        vm.assume(operator != address(this));
        IDelegationManager.OperatorDetails memory _operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(operator, _operatorDetails, emptyStringForMetadataURI);

        // try to delegate again and check that the call reverts
        vm.startPrank(staker);
        vm.expectRevert(STAKER_ALREADY_DELEGATED_ERROR);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);
        vm.stopPrank();
    }

    // @notice Verifies that `staker` cannot delegate to an unregistered `operator`
    function testFuzz_CannotDelegateToUnregisteredOperator(address staker, address operator)
        public
        filterFuzzedAddressInputs(staker)
        filterFuzzedAddressInputs(operator)
    {
        assertTrue(!delegationManager.isOperator(operator), "incorrect test input?");

        // try to delegate and check that the call reverts
        vm.startPrank(staker);
        vm.expectRevert(OPERATOR_NOT_REGISTERED_ERROR);
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);
        vm.stopPrank();
    }

    /**
     * @notice `staker` delegates to an operator who assertTrues signature verification through an EOA (i.e. the operator’s `delegationApprover` address is set to a nonzero EOA)
     * via the `staker` calling `DelegationManager.delegateTo`
     * The function should pass *only with a valid ECDSA signature from the `delegationApprover`, OR if called by the operator or their delegationApprover themselves
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_DelegateToOperatorWhoassertTruesECDSASignature(address staker, bytes32 salt, uint256 expiry)
        public
        filterFuzzedAddressInputs(staker)
    {
        // filter to only valid `expiry` values
        vm.assume(expiry >= block.timestamp);

        address delegationApprover = vm.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        vm.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // verify that the salt hasn't been used before
        assertTrue(
            !delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, operator, salt, expiry);

        // delegate from the `staker` to the operator
        vm.startPrank(staker);
        vm.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, operator);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);
        vm.stopPrank();

        assertTrue(delegationManager.isDelegated(staker), "staker not delegated correctly");
        assertTrue(delegationManager.delegatedTo(staker) == operator, "staker delegated to the wrong address");
        assertTrue(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        if (staker == operator || staker == delegationManager.delegationApprover(operator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertTrue(
                !delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt),
                "salt somehow spent not spent?"
            );
        }
    }

    /**
     * @notice Like `testFuzz_DelegateToOperatorWhoassertTruesECDSASignature` but using an incorrect signature on purpose and checking that reversion occurs
     */
    function testFuzz_DelegateToOperatorWhoassertTruesECDSASignature_RevertsWithBadSignature(address staker, uint256 expiry)
        public
        filterFuzzedAddressInputs(staker)
    {
        // filter to only valid `expiry` values
        vm.assume(expiry >= block.timestamp);

        address delegationApprover = vm.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        vm.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // calculate the signature
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        approverSignatureAndExpiry.expiry = expiry;
        {
            bytes32 digestHash = delegationManager.calculateDelegationApprovalDigestHash(
                staker, operator, delegationManager.delegationApprover(operator), emptySalt, expiry
            );
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(delegationSignerPrivateKey, digestHash);
            // mess up the signature by flipping v's parity
            v = (v == 27 ? 28 : 27);
            approverSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }

        // try to delegate from the `staker` to the operator, and check reversion
        vm.startPrank(staker);
        vm.expectRevert(INVALID_SIGNATURE_ERROR);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);
        vm.stopPrank();
    }

    /**
     * @notice Like `testFuzz_DelegateToOperatorWhoassertTruesECDSASignature` but using an invalid expiry on purpose and checking that reversion occurs
     */
    function testFuzz_DelegateToOperatorWhoassertTruesECDSASignature_RevertsWithExpiredDelegationApproverSignature(
        address staker,
        bytes32 salt,
        uint256 expiry
    ) public filterFuzzedAddressInputs(staker) {
        // roll to a very late timestamp
        vm.roll(type(uint256).max / 2);
        // filter to only *invalid* `expiry` values
        vm.assume(expiry < block.timestamp);

        address delegationApprover = vm.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        vm.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // calculate the delegationSigner's signature
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, operator, salt, expiry);

        // delegate from the `staker` to the operator
        vm.startPrank(staker);
        vm.expectRevert(APPROVER_SIGNATURE_EXPIRED_ERROR);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);
        vm.stopPrank();
    }

    /**
     * @notice `staker` delegates to an operator who assertTrues signature verification through an EIP1271-compliant contract (i.e. the operator’s `delegationApprover` address is
     * set to a nonzero and code-containing address) via the `staker` calling `DelegationManager.delegateTo`
     * The function uses OZ's ERC1271WalletMock contract, and thus should pass *only when a valid ECDSA signature from the `owner` of the ERC1271WalletMock contract,
     * OR if called by the operator or their delegationApprover themselves
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_DelegateToOperatorWhoassertTruesEIP1271Signature(address staker, bytes32 salt, uint256 expiry)
        public
        filterFuzzedAddressInputs(staker)
    {
        // filter to only valid `expiry` values
        vm.assume(expiry >= block.timestamp);

        address delegationSigner = vm.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        vm.assume(staker != operator);

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
        testFuzz_RegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // verify that the salt hasn't been used before
        assertTrue(
            !delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, operator, salt, expiry);

        // delegate from the `staker` to the operator
        vm.startPrank(staker);
        vm.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, operator);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);
        vm.stopPrank();

        assertTrue(delegationManager.isDelegated(staker), "staker not delegated correctly");
        assertTrue(delegationManager.delegatedTo(staker) == operator, "staker delegated to the wrong address");
        assertTrue(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // check that the nonce incremented appropriately
        if (staker == operator || staker == delegationManager.delegationApprover(operator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertTrue(
                !delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt),
                "salt somehow spent not spent?"
            );
        }
    }

    /**
     * @notice Like `testFuzz_DelegateToOperatorWhoassertTruesEIP1271Signature` but using a contract that
     * returns a value other than the EIP1271 "magic bytes" and checking that reversion occurs appropriately
     */
    function testFuzz_DelegateToOperatorWhoassertTruesEIP1271Signature_RevertsOnBadReturnValue(
        address staker,
        uint256 expiry
    ) public filterFuzzedAddressInputs(staker) {
        // filter to only valid `expiry` values
        vm.assume(expiry >= block.timestamp);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        vm.assume(staker != operator);

        // deploy a ERC1271MaliciousMock contract that will return an incorrect value when called
        ERC1271MaliciousMock wallet = new ERC1271MaliciousMock();

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(wallet),
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // create the signature struct
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        approverSignatureAndExpiry.expiry = expiry;

        // try to delegate from the `staker` to the operator, and check reversion
        vm.startPrank(staker);
        // because the ERC1271MaliciousMock contract returns the wrong amount of data, we get a low-level "EvmError: Revert" message here rather than the error message bubbling up
        // vm.expectRevert(bytes("EIP1271SignatureUtils.checkSignature_EIP1271: ERC1271 signature verification failed"));
        vm.expectRevert();
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);
        vm.stopPrank();
    }

    /**
     * @notice `staker` becomes delegated to an operator who does not assertTrue any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * via the `caller` calling `DelegationManager.delegateToBySignature`
     * The function should pass with any `operatorSignature` input (since it should be unused)
     * The function should pass only with a valid `stakerSignatureAndExpiry` input
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_DelegateBySignatureToOperatorWhoAcceptsAllStakers(address caller, bytes32 salt, uint256 expiry)
        public
        filterFuzzedAddressInputs(caller)
    {
        // filter to only valid `expiry` values
        vm.assume(expiry >= block.timestamp);

        address staker = vm.addr(stakerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        vm.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // verify that the salt hasn't been used before
        assertTrue(
            !delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt),
            "salt somehow spent too early?"
        );
        // fetch the staker's current nonce
        uint256 currentStakerNonce = delegationManager.stakerNonce(staker);
        // calculate the staker signature
        IDelegationManager.SignatureWithExpiry memory stakerSignatureAndExpiry =
            _getStakerSignature(stakerPrivateKey, operator, expiry);

        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        vm.startPrank(caller);
        vm.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, operator);
        // use an empty approver signature input since none is needed / the input is unchecked
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateToBySignature(
            staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, emptySalt
        );
        vm.stopPrank();

        // check all the delegation status changes
        assertTrue(delegationManager.isDelegated(staker), "staker not delegated correctly");
        assertTrue(delegationManager.delegatedTo(staker) == operator, "staker delegated to the wrong address");
        assertTrue(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // check that the staker nonce incremented appropriately
        assertTrue(delegationManager.stakerNonce(staker) == currentStakerNonce + 1, "staker nonce did not increment");
        // verify that the salt is still marked as unused (since it wasn't checked or used)
        assertTrue(
            !delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt),
            "salt somehow spent too incorrectly?"
        );
    }

    /**
     * @notice `staker` becomes delegated to an operator who assertTrues signature verification through an EOA (i.e. the operator’s `delegationApprover` address is set to a nonzero EOA)
     * via the `caller` calling `DelegationManager.delegateToBySignature`
     * The function should pass *only with a valid ECDSA signature from the `delegationApprover`, OR if called by the operator or their delegationApprover themselves
     * AND with a valid `stakerSignatureAndExpiry` input
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_DelegateBySignatureToOperatorWhoassertTruesECDSASignature(address caller, bytes32 salt, uint256 expiry)
        public
        filterFuzzedAddressInputs(caller)
    {
        // filter to only valid `expiry` values
        vm.assume(expiry >= block.timestamp);

        address staker = vm.addr(stakerPrivateKey);
        address delegationApprover = vm.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        vm.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // verify that the salt hasn't been used before
        assertTrue(
            !delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, operator, salt, expiry);

        // fetch the staker's current nonce
        uint256 currentStakerNonce = delegationManager.stakerNonce(staker);
        // calculate the staker signature
        IDelegationManager.SignatureWithExpiry memory stakerSignatureAndExpiry =
            _getStakerSignature(stakerPrivateKey, operator, expiry);

        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        vm.startPrank(caller);
        vm.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, operator);
        delegationManager.delegateToBySignature(
            staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, salt
        );
        vm.stopPrank();

        assertTrue(delegationManager.isDelegated(staker), "staker not delegated correctly");
        assertTrue(delegationManager.delegatedTo(staker) == operator, "staker delegated to the wrong address");
        assertTrue(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // check that the delegationApprover nonce incremented appropriately
        if (caller == operator || caller == delegationManager.delegationApprover(operator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertTrue(
                !delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(operator), salt),
                "salt somehow spent not spent?"
            );
        }

        // check that the staker nonce incremented appropriately
        assertTrue(delegationManager.stakerNonce(staker) == currentStakerNonce + 1, "staker nonce did not increment");
    }

    /**
     * @notice `staker` becomes delegated to an operatorwho assertTrues signature verification through an EIP1271-compliant contract (i.e. the operator’s `delegationApprover` address is
     * set to a nonzero and code-containing address) via the `caller` calling `DelegationManager.delegateToBySignature`
     * The function uses OZ's ERC1271WalletMock contract, and thus should pass *only when a valid ECDSA signature from the `owner` of the ERC1271WalletMock contract,
     * OR if called by the operator or their delegationApprover themselves
     * AND with a valid `stakerSignatureAndExpiry` input
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testFuzz_DelegateBySignatureToOperatorWhoassertTruesEIP1271Signature(
        address caller,
        bytes32 salt,
        uint256 expiry
    ) public filterFuzzedAddressInputs(caller) {
        // filter to only valid `expiry` values
        vm.assume(expiry >= block.timestamp);

        address staker = vm.addr(stakerPrivateKey);
        address delegationSigner = vm.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        vm.assume(staker != _operator);

        /**
         * deploy a ERC1271WalletMock contract with the `delegationSigner` address as the owner,
         * so that we can create valid signatures from the `delegationSigner` for the contract to check when called
         */
        ERC1271WalletMock wallet = new ERC1271WalletMock(delegationSigner);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: _operator,
            delegationApprover: address(wallet),
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(_operator, operatorDetails, emptyStringForMetadataURI);

        // verify that the salt hasn't been used before
        assertTrue(
            !delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(_operator), salt),
            "salt somehow spent too early?"
        );
        // calculate the delegationSigner's signature
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, _operator, salt, expiry);

        // fetch the staker's current nonce
        uint256 currentStakerNonce = delegationManager.stakerNonce(staker);
        // calculate the staker signature
        IDelegationManager.SignatureWithExpiry memory stakerSignatureAndExpiry =
            _getStakerSignature(stakerPrivateKey, _operator, expiry);

        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        vm.startPrank(caller);
        vm.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, _operator);
        delegationManager.delegateToBySignature(
            staker, _operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, salt
        );
        vm.stopPrank();

        assertTrue(delegationManager.isDelegated(staker), "staker not delegated correctly");
        assertTrue(delegationManager.delegatedTo(staker) == _operator, "staker delegated to the wrong address");
        assertTrue(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // check that the delegationApprover nonce incremented appropriately
        if (caller == _operator || caller == delegationManager.delegationApprover(_operator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            assertTrue(
                !delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(_operator), salt),
                "salt somehow spent too incorrectly?"
            );
        } else {
            // verify that the salt is marked as used
            assertTrue(
                delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(_operator), salt),
                "salt somehow spent not spent?"
            );
        }

        // check that the staker nonce incremented appropriately
        assertTrue(delegationManager.stakerNonce(staker) == currentStakerNonce + 1, "staker nonce did not increment");
    }

    // @notice Checks that `DelegationManager.delegateToBySignature` reverts if the staker's signature has expired
    function testFuzz_DelegateBySignatureRevertsWhenStakerSignatureExpired(
        address staker,
        address operator,
        uint256 expiry,
        bytes memory signature
    ) public {
        vm.assume(expiry < block.timestamp);
        vm.expectRevert(STAKER_SIGNATURE_EXPIRED_ERROR);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry =
            IDelegationManager.SignatureWithExpiry({signature: signature, expiry: expiry});
        delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, emptySalt);
    }

    // @notice Checks that `DelegationManager.delegateToBySignature` reverts if the delegationApprover's signature has expired and their signature is checked
    function testFuzz_DelegateBySignatureRevertsWhenDelegationApproverSignatureExpired(
        address caller,
        uint256 stakerExpiry,
        uint256 delegationApproverExpiry
    ) public filterFuzzedAddressInputs(caller) {
        // filter to only valid `stakerExpiry` values
        vm.assume(stakerExpiry >= block.timestamp);
        // roll to a very late timestamp
        vm.roll(type(uint256).max / 2);
        // filter to only *invalid* `delegationApproverExpiry` values
        vm.assume(delegationApproverExpiry < block.timestamp);

        address staker = vm.addr(stakerPrivateKey);
        address delegationApprover = vm.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        vm.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // calculate the delegationSigner's signature
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, operator, emptySalt, delegationApproverExpiry);

        // calculate the staker signature
        IDelegationManager.SignatureWithExpiry memory stakerSignatureAndExpiry =
            _getStakerSignature(stakerPrivateKey, operator, stakerExpiry);

        // try delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`, and check for reversion
        vm.startPrank(caller);
        vm.expectRevert(APPROVER_SIGNATURE_EXPIRED_ERROR);
        delegationManager.delegateToBySignature(
            staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, emptySalt
        );
        vm.stopPrank();
    }

    /**
     * @notice Like `testFuzz_DelegateToOperatorWhoassertTruesECDSASignature` but using an invalid expiry on purpose and checking that reversion occurs
     */
    function testFuzz_DelegateToOperatorWhoassertTruesECDSASignature_RevertsWithExpiredSignature(
        address staker,
        uint256 expiry
    ) public filterFuzzedAddressInputs(staker) {
        // roll to a very late timestamp
        vm.roll(type(uint256).max / 2);
        // filter to only *invalid* `expiry` values
        vm.assume(expiry < block.timestamp);

        address delegationApprover = vm.addr(delegationSignerPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        vm.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // calculate the delegationSigner's signature
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, operator, emptySalt, expiry);

        // delegate from the `staker` to the operator
        vm.startPrank(staker);
        vm.expectRevert(APPROVER_SIGNATURE_EXPIRED_ERROR);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);
        vm.stopPrank();
    }

    /**
     * Staker is undelegated from an operator, via a call to `undelegate`, properly originating from the StrategyManager address.
     * Reverts if called by any address that is not the StrategyManager
     * Reverts if the staker is themselves an operator (i.e. they are delegated to themselves)
     * Does nothing if the staker is already undelegated
     * Properly undelegates the staker, i.e. the staker becomes “delegated to” the zero address, and `isDelegated(staker)` returns ‘false’
     * Emits a `StakerUndelegated` event
     */
    function testFuzz_UndelegateFromOperator(address staker) public {
        // register *this contract* as an operator and delegate from the `staker` to them (already filters out case when staker is the operator since it will revert)
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        testFuzz_DelegateToOperatorWhoAcceptsAllStakers(staker, approverSignatureAndExpiry, emptySalt);

        vm.startPrank(address(strategyManagerMock));
        vm.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(staker, delegationManager.delegatedTo(staker));
        delegationManager.undelegate(staker);
        vm.stopPrank();

        assertTrue(!delegationManager.isDelegated(staker), "staker not undelegated!");
        assertTrue(
            delegationManager.delegatedTo(staker) == address(0),
            "undelegated staker should be delegated to zero address"
        );
    }

    // @notice Verifies that an operator cannot undelegate from themself (this should always be forbidden)
    function testFuzz_OperatorCannotUndelegateFromThemself(address operator) public fuzzedAddress(operator) {
        vm.startPrank(operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        vm.stopPrank();
        vm.expectRevert(OPERATOR_UNDELEGATE_SELF_ERROR);

        vm.startPrank(address(strategyManagerMock));
        delegationManager.undelegate(operator);
        vm.stopPrank();
    }

    // @notice Verifies that `DelegationManager.undelegate` reverts if not called by the StrategyManager
    function testFuzz_CannotCallUndelegateFromNonStrategyManagerAddress(address caller) public fuzzedAddress(caller) {
        vm.assume(caller != address(strategyManagerMock));
        vm.expectRevert(NOT_STRATEGY_MANAGER_ERROR);
        vm.startPrank(caller);
        delegationManager.undelegate(address(this));
    }

    /**
     * @notice Verifies that `DelegationManager.increaseDelegatedShares` properly increases the delegated `shares` that the operator
     * who the `staker` is delegated to has in the strategy
     * @dev Checks that there is no change if the staker is not delegated
     */
    function testFuzz_IncreaseDelegatedShares(address staker, uint256 shares, bool delegateFromStakerToOperator) public {
        IStrategy strategy = strategyMock;

        // register *this contract* as an operator
        address operator = address(this);
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        // filter inputs, since delegating to the operator will fail when the staker is already registered as an operator
        vm.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // delegate from the `staker` to the operator *if `delegateFromStakerToOperator` is 'true'*
        if (delegateFromStakerToOperator) {
            vm.startPrank(staker);
            vm.expectEmit(true, true, true, true, address(delegationManager));
            emit StakerDelegated(staker, operator);
            delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);
            vm.stopPrank();
        }

        uint256 delegatedSharesBefore =
            delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategy);

        vm.startPrank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, strategy, shares);
        vm.stopPrank();

        uint256 delegatedSharesAfter = delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategy);

        if (delegationManager.isDelegated(staker)) {
            assertTrue(
                delegatedSharesAfter == delegatedSharesBefore + shares, "delegated shares did not increment correctly"
            );
        } else {
            assertTrue(delegatedSharesAfter == delegatedSharesBefore, "delegated shares incremented incorrectly");
            assertTrue(delegatedSharesBefore == 0, "nonzero shares delegated to zero address!");
        }
    }

    /**
     * @notice Verifies that `DelegationManager.decreaseDelegatedShares` properly decreases the delegated `shares` that the operator
     * who the `staker` is delegated to has in the strategies
     * @dev Checks that there is no change if the staker is not delegated
     */
    function testFuzz_DecreaseDelegatedShares(
        address staker,
        IStrategy[] memory strategies,
        uint256 shares,
        bool delegateFromStakerToOperator
    ) public {
        // sanity-filtering on fuzzed input length
        vm.assume(strategies.length <= 64);
        // register *this contract* as an operator
        address operator = address(this);
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        // filter inputs, since delegating to the operator will fail when the staker is already registered as an operator
        vm.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // delegate from the `staker` to the operator *if `delegateFromStakerToOperator` is 'true'*
        if (delegateFromStakerToOperator) {
            vm.startPrank(staker);
            vm.expectEmit(true, true, true, true, address(delegationManager));
            emit StakerDelegated(staker, operator);
            delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);
            vm.stopPrank();
        }

        uint256[] memory delegatedSharesBefore = new uint256[](strategies.length);
        uint256[] memory sharesInputArray = new uint256[](strategies.length);

        // for each strategy in `strategies`, increase delegated shares by `shares`
        vm.startPrank(address(strategyManagerMock));
        for (uint256 i = 0; i < strategies.length; ++i) {
            delegationManager.increaseDelegatedShares(staker, strategies[i], shares);
            delegatedSharesBefore[i] =
                delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategies[i]);
            // also construct an array which we'll use in another loop
            sharesInputArray[i] = shares;
        }
        vm.stopPrank();

        // for each strategy in `strategies`, decrease delegated shares by `shares`
        vm.startPrank(address(strategyManagerMock));
        delegationManager.decreaseDelegatedShares(delegationManager.delegatedTo(staker), strategies, sharesInputArray);
        vm.stopPrank();

        // check shares after call to `decreaseDelegatedShares`
        for (uint256 i = 0; i < strategies.length; ++i) {
            uint256 delegatedSharesAfter =
                delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategies[i]);

            if (delegationManager.isDelegated(staker)) {
                assertTrue(
                    delegatedSharesAfter == delegatedSharesBefore[i] - sharesInputArray[i],
                    "delegated shares did not decrement correctly"
                );
            } else {
                assertTrue(delegatedSharesAfter == delegatedSharesBefore[i], "delegated shares decremented incorrectly");
                assertTrue(delegatedSharesBefore[i] == 0, "nonzero shares delegated to zero address!");
            }
        }
    }

    // @notice Verifies that `DelegationManager.increaseDelegatedShares` reverts if not called by the StrategyManager
    function testFuzz_CannotCallIncreaseDelegatedSharesFromNonStrategyManagerAddress(address operator, uint256 shares)
        public
        fuzzedAddress(operator)
    {
        vm.assume(operator != address(strategyManagerMock));
        vm.expectRevert(NOT_STRATEGY_MANAGER_ERROR);
        vm.startPrank(operator);
        delegationManager.increaseDelegatedShares(operator, strategyMock, shares);
    }

    // @notice Verifies that `DelegationManager.decreaseDelegatedShares` reverts if not called by the StrategyManager
    function testFuzz_CannotCallDecreaseDelegatedSharesFromNonStrategyManagerAddress(
        address operator,
        IStrategy[] memory strategies,
        uint256[] memory shareAmounts
    ) public fuzzedAddress(operator) {
        vm.assume(operator != address(strategyManagerMock));
        vm.expectRevert(NOT_STRATEGY_MANAGER_ERROR);
        vm.startPrank(operator);
        delegationManager.decreaseDelegatedShares(operator, strategies, shareAmounts);
    }

    // @notice Verifies that it is not possible for a staker to delegate to an operator when the operator is frozen in EigenLayer
    function testFuzz_CannotDelegateWhenOperatorIsFrozen(address operator, address staker)
        public
        fuzzedAddress(operator)
        fuzzedAddress(staker)
    {
        vm.assume(operator != staker);

        vm.startPrank(operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        vm.stopPrank();

        slasherMock.setOperatorFrozenStatus(operator, true);
        vm.expectRevert(OPERATOR_FROZEN_ERROR);
        vm.startPrank(staker);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry, emptySalt);
        vm.stopPrank();
    }

    // @notice Verifies that it is not possible for a staker to delegate to an operator when they are already delegated to an operator
    function testFuzz_CannotDelegateWhenStakerHasExistingDelegation(address staker, address operator, address operator2)
        public
        fuzzedAddress(staker)
        fuzzedAddress(operator)
        fuzzedAddress(operator2)
    {
        vm.assume(operator != operator2);
        vm.assume(staker != operator);
        vm.assume(staker != operator2);

        vm.startPrank(operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        vm.stopPrank();

        vm.startPrank(operator2);
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        vm.stopPrank();

        vm.startPrank(staker);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry, emptySalt);
        vm.stopPrank();

        vm.startPrank(staker);
        vm.expectRevert(STAKER_ALREADY_DELEGATED_ERROR);
        delegationManager.delegateTo(operator2, signatureWithExpiry, emptySalt);
        vm.stopPrank();
    }

    // @notice Verifies that it is not possible to delegate to an unregistered operator
    function testFuzz_CannotDelegateToUnregisteredOperator(address operator) public {
        vm.expectRevert(OPERATOR_NOT_REGISTERED_ERROR);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry, emptySalt);
    }

    // @notice Verifies that delegating is not possible when the "new delegations paused" switch is flipped
    function testFuzz_CannotDelegateWhenPausedNewDelegationIsSet(address operator, address staker)
        public
        fuzzedAddress(operator)
        fuzzedAddress(staker)
    {
        vm.startPrank(pauser);
        delegationManager.pause(1);
        vm.stopPrank();

        vm.startPrank(staker);
        vm.expectRevert(PAUSED_INDEX_ERROR);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry, emptySalt);
        vm.stopPrank();
    }

    // special event purely used in the StrategyManagerMock contract, inside of `testForceUndelegation` to verify that the correct call is made
    event ForceTotalWithdrawalCalled(address staker);

    /**
     * @notice Verifies that the `forceUndelegation` function properly calls `strategyManager.forceTotalWithdrawal`
     * @param callFromOperatorOrApprover -- calls from the operator if 'false' and the 'approver' if true
     */
    function testFuzz_ForceUndelegation(address staker, bytes32 salt, bool callFromOperatorOrApprover)
        public
        fuzzedAddress(staker)
    {
        address delegationApprover = vm.addr(delegationSignerPrivateKey);
        address operator = address(this);

        // filtering since you can't delegate to yourself after registering as an operator
        vm.assume(staker != operator);

        // register this contract as an operator and delegate from the staker to it
        uint256 expiry = type(uint256).max;
        testFuzz_DelegateToOperatorWhoassertTruesECDSASignature(staker, salt, expiry);

        address caller;
        if (callFromOperatorOrApprover) {
            caller = delegationApprover;
        } else {
            caller = operator;
        }

        // call the `forceUndelegation` function and check that the correct calldata is forwarded by looking for an event emitted by the StrategyManagerMock contract
        vm.startPrank(caller);
        vm.expectEmit(true, true, true, true, address(strategyManagerMock));
        emit ForceTotalWithdrawalCalled(staker);
        bytes32 returnValue = delegationManager.forceUndelegation(staker);
        // check that the return value is empty, as specified in the mock contract
        assertTrue(returnValue == bytes32(uint256(0)), "mock contract returned wrong return value");
        vm.stopPrank();
    }

    /**
     * @notice Verifies that the `forceUndelegation` function has proper access controls (can only be called by the operator who the `staker` has delegated
     * to or the operator's `delegationApprover`)
     */
    function testFuzz_CannotCallForceUndelegationFromImproperAddress(address staker, address caller)
        public
        fuzzedAddress(staker)
        fuzzedAddress(caller)
    {
        address delegationApprover = vm.addr(delegationSignerPrivateKey);
        address operator = address(this);

        // filtering since you can't delegate to yourself after registering as an operator
        vm.assume(staker != operator);

        // filter out addresses that are actually allowed to call the function
        vm.assume(caller != operator);
        vm.assume(caller != delegationApprover);

        // register this contract as an operator and delegate from the staker to it
        uint256 expiry = type(uint256).max;
        testFuzz_DelegateToOperatorWhoassertTruesECDSASignature(staker, emptySalt, expiry);

        // try to call the `forceUndelegation` function and check for reversion
        vm.startPrank(caller);
        vm.expectRevert(NOT_OPERATOR_OR_APPOVER_ERROR);
        delegationManager.forceUndelegation(staker);
        vm.stopPrank();
    }

    /**
     * @notice verifies that `DelegationManager.forceUndelegation` reverts if trying to undelegate an operator from themselves
     * @param callFromOperatorOrApprover -- calls from the operator if 'false' and the 'approver' if true
     */
    function testFuzz_OperatorCannotForceUndelegateThemself(address delegationApprover, bool callFromOperatorOrApprover)
        public
    {
        // register *this contract* as an operator
        address operator = address(this);
        IDelegationManager.OperatorDetails memory _operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testFuzz_RegisterAsOperator(operator, _operatorDetails, emptyStringForMetadataURI);

        address caller;
        if (callFromOperatorOrApprover) {
            caller = delegationApprover;
        } else {
            caller = operator;
        }

        // try to call the `forceUndelegation` function and check for reversion
        vm.startPrank(caller);
        vm.expectRevert(FORCE_UNDELEGATE_OPERATOR_ERROR);
        delegationManager.forceUndelegation(operator);
        vm.stopPrank();
    }

    /**
     * @notice Verifies that the reversion occurs when trying to reuse an 'approverSalt'
     */
    function testFuzz_Revert_WhenTryingToReuseSalt(address staker_one, address staker_two, bytes32 salt)
        public
        fuzzedAddress(staker_one)
        fuzzedAddress(staker_two)
    {
        // address delegationApprover = vm.addr(delegationSignerPrivateKey);
        address operator = address(this);

        // filtering since you can't delegate to yourself after registering as an operator
        vm.assume(staker_one != operator);
        vm.assume(staker_two != operator);

        // register this contract as an operator and delegate from `staker_one` to it, using the `salt`
        uint256 expiry = type(uint256).max;
        testFuzz_DelegateToOperatorWhoassertTruesECDSASignature(staker_one, salt, expiry);

        // calculate the delegationSigner's signature
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker_two, operator, salt, expiry);

        // try to delegate to the operator from `staker_two`, and verify that the call reverts for the proper reason (trying to reuse a salt)
        vm.startPrank(staker_two);
        vm.expectRevert(APPROVER_SALT_SPENT_ERROR);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);
        vm.stopPrank();
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
    ) internal view returns (IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry) {
        approverSignatureAndExpiry.expiry = expiry;
        {
            bytes32 digestHash = delegationManager.calculateDelegationApprovalDigestHash(
                staker, operator, delegationManager.delegationApprover(operator), salt, expiry
            );
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(_delegationSignerPrivateKey, digestHash);
            approverSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }
        return approverSignatureAndExpiry;
    }

    /**
     * @notice internal function for calculating a signature from the staker corresponding to `_stakerPrivateKey`, delegating them to
     * the `operator`, and expiring at `expiry`.
     */
    function _getStakerSignature(uint256 _stakerPrivateKey, address operator, uint256 expiry)
        internal
        view
        returns (IDelegationManager.SignatureWithExpiry memory stakerSignatureAndExpiry)
    {
        address staker = vm.addr(stakerPrivateKey);
        stakerSignatureAndExpiry.expiry = expiry;
        {
            bytes32 digestHash = delegationManager.calculateCurrentStakerDelegationDigestHash(staker, operator, expiry);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(_stakerPrivateKey, digestHash);
            stakerSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }
        return stakerSignatureAndExpiry;
    }
}
