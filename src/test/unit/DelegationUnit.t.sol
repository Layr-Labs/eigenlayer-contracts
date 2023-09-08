// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "forge-std/Test.sol";

import "../mocks/StrategyManagerMock.sol";
import "../mocks/SlasherMock.sol";
import "../mocks/EigenPodManagerMock.sol";
import "../EigenLayerTestHelper.t.sol";
import "../mocks/ERC20Mock.sol";
import "../Delegation.t.sol";

contract DelegationUnitTests is EigenLayerTestHelper {

    StrategyManagerMock strategyManagerMock;
    SlasherMock slasherMock;
    DelegationManager delegationManager;
    DelegationManager delegationManagerImplementation;
    StrategyBase strategyImplementation;
    StrategyBase strategyMock;
    EigenPodManagerMock eigenPodManagerMock;

    uint256 delegationSignerPrivateKey = uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);
    uint256 stakerPrivateKey = uint256(123456789);

    // empty string reused across many tests
    string emptyStringForMetadataURI;

    // "empty" / zero salt, reused across many tests
    bytes32 emptySalt;

    // reused in various tests. in storage to help handle stack-too-deep errors
    address _operator = address(this);

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

    // @notice reuseable modifier + associated mapping for filtering out weird fuzzed inputs, like making calls from the ProxyAdmin or the zero address
    mapping(address => bool) public addressIsExcludedFromFuzzedInputs;
    modifier filterFuzzedAddressInputs(address fuzzedAddress) {
        cheats.assume(!addressIsExcludedFromFuzzedInputs[fuzzedAddress]);
        _;
    }

    function setUp() override virtual public{
        EigenLayerDeployer.setUp();

        slasherMock = new SlasherMock();
        delegationManager = DelegationManager(address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), "")));
        strategyManagerMock = new StrategyManagerMock();
        eigenPodManagerMock = new EigenPodManagerMock();

        delegationManagerImplementation = new DelegationManager(strategyManagerMock, slasherMock, eigenPodManagerMock);

        cheats.startPrank(eigenLayerProxyAdmin.owner());
        eigenLayerProxyAdmin.upgrade(TransparentUpgradeableProxy(payable(address(delegationManager))), address(delegationManagerImplementation));
        cheats.stopPrank();
        
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
        require(delegationManager.strategyManager() == strategyManagerMock,
            "constructor / initializer incorrect, strategyManager set wrong");
        require(delegationManager.slasher() == slasherMock,
            "constructor / initializer incorrect, slasher set wrong");
        require(delegationManager.pauserRegistry() == eigenLayerPauserReg,
            "constructor / initializer incorrect, pauserRegistry set wrong");
        require(delegationManager.owner() == initalOwner,
            "constructor / initializer incorrect, owner set wrong");
        require(delegationManager.paused() == initialPausedStatus,
            "constructor / initializer incorrect, paused status set wrong");
    }

    // @notice Verifies that the DelegationManager cannot be iniitalized multiple times
    function testCannotReinitializeDelegationManager() public {
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        delegationManager.initialize(address(this), eigenLayerPauserReg, 0);
    }

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

        cheats.startPrank(operator);
        cheats.expectRevert(bytes("DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be > MAX_STAKER_OPT_OUT_WINDOW_BLOCKS"));
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        cheats.stopPrank();
    }
    
    /**
     * @notice Verifies that an operator cannot register with `earningsReceiver` set to the zero address
     * @dev This is an important check since we check `earningsReceiver != address(0)` to check if an address is an operator!
     */
    function testCannotRegisterAsOperatorWithZeroAddressAsEarningsReceiver() public {
        cheats.startPrank(operator);
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

        // register *this contract* as an operator
        IDelegationManager.OperatorDetails memory _operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: _operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(_operator, _operatorDetails, emptyStringForMetadataURI);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateTo(_operator, approverSignatureAndExpiry, emptySalt);        

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
        testRegisterAsOperator(_operator, initialOperatorDetails, emptyStringForMetadataURI);
        // filter out zero address since people can't set their earningsReceiver address to the zero address (special test case to verify)
        cheats.assume(modifiedOperatorDetails.earningsReceiver != address(0));

        cheats.startPrank(_operator);

        // either it fails for trying to set the stakerOptOutWindowBlocks
        if (modifiedOperatorDetails.stakerOptOutWindowBlocks > delegationManager.MAX_STAKER_OPT_OUT_WINDOW_BLOCKS()) {
            cheats.expectRevert(bytes("DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be > MAX_STAKER_OPT_OUT_WINDOW_BLOCKS"));
            delegationManager.modifyOperatorDetails(modifiedOperatorDetails);
        // or the transition is allowed,
        } else if (modifiedOperatorDetails.stakerOptOutWindowBlocks >= initialOperatorDetails.stakerOptOutWindowBlocks) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorDetailsModified(_operator, modifiedOperatorDetails);
            delegationManager.modifyOperatorDetails(modifiedOperatorDetails);

            require(modifiedOperatorDetails.earningsReceiver == delegationManager.earningsReceiver(_operator), "earningsReceiver not set correctly");
            require(modifiedOperatorDetails.delegationApprover == delegationManager.delegationApprover(_operator), "delegationApprover not set correctly");
            require(modifiedOperatorDetails.stakerOptOutWindowBlocks == delegationManager.stakerOptOutWindowBlocks(_operator), "stakerOptOutWindowBlocks not set correctly");
            require(delegationManager.delegatedTo(_operator) == _operator, "operator not delegated to self");
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
            earningsReceiver: _operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(_operator, operatorDetails, emptyStringForMetadataURI);

        // call `updateOperatorMetadataURI` and check for event
        cheats.startPrank(_operator);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit OperatorMetadataURIUpdated(_operator, metadataURI);
        delegationManager.updateOperatorMetadataURI(metadataURI);
        cheats.stopPrank();
    }

    // @notice Tests that an address which is not an operator cannot successfully call `updateOperatorMetadataURI`.
    function testCannotUpdateOperatorMetadataURIWithoutRegisteringFirst() public {
        require(!delegationManager.isOperator(_operator), "bad test setup");

        cheats.startPrank(_operator);
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
            earningsReceiver: _operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(_operator, operatorDetails, emptyStringForMetadataURI);

        operatorDetails.earningsReceiver = address(0);
        cheats.expectRevert(bytes("DelegationManager._setOperatorDetails: cannot set `earningsReceiver` to zero address"));
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
    function testDelegateToOperatorWhoAcceptsAllStakers(address staker, IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry, bytes32 salt) public 
        filterFuzzedAddressInputs(staker)
    {
        // register *this contract* as an operator
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != _operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: _operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(_operator, operatorDetails, emptyStringForMetadataURI);

        // verify that the salt hasn't been used before
        require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(_operator), salt), "salt somehow spent too early?");

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, _operator);
        delegationManager.delegateTo(_operator, approverSignatureAndExpiry, salt);        
        cheats.stopPrank();

        require(delegationManager.isDelegated(staker), "staker not delegated correctly");
        require(delegationManager.delegatedTo(staker) == _operator, "staker delegated to the wrong address");
        require(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // verify that the salt is still marked as unused (since it wasn't checked or used)
        require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(_operator), salt), "salt somehow spent too early?");
    }

    /**
     * @notice Delegates from `staker` to an operator, then verifies that the `staker` cannot delegate to another `operator` (at least without first undelegating)
     */
    function testCannotDelegateWhileDelegated(address staker, address operator, IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry, bytes32 salt) public 
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
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
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
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(delegationSignerPrivateKey, staker, operator, salt, expiry);

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
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
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
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(delegationSignerPrivateKey, staker, operator, salt, expiry);

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
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(delegationSignerPrivateKey, staker, operator, salt, expiry);

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

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(wallet),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails, emptyStringForMetadataURI);

        // create the signature struct
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
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
        IDelegationManager.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(stakerPrivateKey, operator, expiry);

        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        cheats.startPrank(caller);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, operator);
        // use an empty approver signature input since none is needed / the input is unchecked
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
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
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(delegationSignerPrivateKey, staker, operator, salt, expiry);

        // fetch the staker's current nonce
        uint256 currentStakerNonce = delegationManager.stakerNonce(staker);
        // calculate the staker signature
        IDelegationManager.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(stakerPrivateKey, operator, expiry);

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
        cheats.assume(staker != _operator);

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
        testRegisterAsOperator(_operator, operatorDetails, emptyStringForMetadataURI);

        // verify that the salt hasn't been used before
        require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(_operator), salt), "salt somehow spent too early?");
        // calculate the delegationSigner's signature
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry = _getApproverSignature(delegationSignerPrivateKey, staker, _operator, salt, expiry);

        // fetch the staker's current nonce
        uint256 currentStakerNonce = delegationManager.stakerNonce(staker);
        // calculate the staker signature
        IDelegationManager.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(stakerPrivateKey, _operator, expiry);

        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        cheats.startPrank(caller);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, _operator);
        delegationManager.delegateToBySignature(staker, _operator, stakerSignatureAndExpiry, approverSignatureAndExpiry, salt);        
        cheats.stopPrank();

        require(delegationManager.isDelegated(staker), "staker not delegated correctly");
        require(delegationManager.delegatedTo(staker) == _operator, "staker delegated to the wrong address");
        require(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // check that the delegationApprover nonce incremented appropriately
        if (caller == _operator || caller == delegationManager.delegationApprover(_operator)) {
            // verify that the salt is still marked as unused (since it wasn't checked or used)
            require(!delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(_operator), salt), "salt somehow spent too incorrectly?");
        } else {
            // verify that the salt is marked as used
            require(delegationManager.delegationApproverSaltIsSpent(delegationManager.delegationApprover(_operator), salt), "salt somehow spent not spent?");
        }

        // check that the staker nonce incremented appropriately
        require(delegationManager.stakerNonce(staker) == currentStakerNonce + 1,
            "staker nonce did not increment");
    }

    // @notice Checks that `DelegationManager.delegateToBySignature` reverts if the staker's signature has expired
    function testDelegateBySignatureRevertsWhenStakerSignatureExpired(address staker, address operator, uint256 expiry, bytes memory signature) public{
        cheats.assume(expiry < block.timestamp);
        cheats.expectRevert(bytes("DelegationManager.delegateToBySignature: staker signature expired"));
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry = IDelegationManager.SignatureWithExpiry({
            signature: signature,
            expiry: expiry
        });
        delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, emptySalt);
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
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, operator, emptySalt, delegationApproverExpiry);

        // calculate the staker signature
        IDelegationManager.SignatureWithExpiry memory stakerSignatureAndExpiry = _getStakerSignature(stakerPrivateKey, operator, stakerExpiry);

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
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker, operator, emptySalt, expiry);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectRevert(bytes("DelegationManager._delegate: approver signature expired"));
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, emptySalt);        
        cheats.stopPrank();
    }

    /**
     * Staker is undelegated from an operator, via a call to `undelegate`, properly originating from the StrategyManager address.
     * Reverts if called by any address that is not the StrategyManager
     * Reverts if the staker is themselves an operator (i.e. they are delegated to themselves)
     * Does nothing if the staker is already undelegated
     * Properly undelegates the staker, i.e. the staker becomes “delegated to” the zero address, and `isDelegated(staker)` returns ‘false’
     * Emits a `StakerUndelegated` event
     */
    function testUndelegateFromOperator(address staker) public {
        // register *this contract* as an operator and delegate from the `staker` to them (already filters out case when staker is the operator since it will revert)
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        testDelegateToOperatorWhoAcceptsAllStakers(staker, approverSignatureAndExpiry, emptySalt);

        cheats.startPrank(address(strategyManagerMock));
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerUndelegated(staker, delegationManager.delegatedTo(staker));
        delegationManager.undelegate(staker);
        cheats.stopPrank();

        require(!delegationManager.isDelegated(staker), "staker not undelegated!");
        require(delegationManager.delegatedTo(staker) == address(0), "undelegated staker should be delegated to zero address");
    }

    // @notice Verifies that an operator cannot undelegate from themself (this should always be forbidden)
    function testOperatorCannotUndelegateFromThemself(address operator) public fuzzedAddress(operator) {
        cheats.startPrank(operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        cheats.stopPrank();
        cheats.expectRevert(bytes("DelegationManager.undelegate: operators cannot undelegate from themselves"));
        
        cheats.startPrank(address(strategyManagerMock));
        delegationManager.undelegate(operator);
        cheats.stopPrank();
    }

    // @notice Verifies that `DelegationManager.undelegate` reverts if not called by the StrategyManager
    function testCannotCallUndelegateFromNonStrategyManagerAddress(address caller) public fuzzedAddress(caller) {
        cheats.assume(caller != address(strategyManagerMock));
        cheats.expectRevert(bytes("onlyStrategyManager"));
        cheats.startPrank(caller);
        delegationManager.undelegate(address(this));
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
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
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

        uint256 delegatedSharesBefore = delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategy);

        cheats.startPrank(address(strategyManagerMock));
        delegationManager.increaseDelegatedShares(staker, strategy, shares);
        cheats.stopPrank();

        uint256 delegatedSharesAfter = delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategy);

        if (delegationManager.isDelegated(staker)) {
            require(delegatedSharesAfter == delegatedSharesBefore + shares, "delegated shares did not increment correctly");
        } else {
            require(delegatedSharesAfter == delegatedSharesBefore, "delegated shares incremented incorrectly");
            require(delegatedSharesBefore == 0, "nonzero shares delegated to zero address!");
        }
    }

    /**
     * @notice Verifies that `DelegationManager.decreaseDelegatedShares` properly decreases the delegated `shares` that the operator
     * who the `staker` is delegated to has in the strategies
     * @dev Checks that there is no change if the staker is not delegated
     */
    function testDecreaseDelegatedShares(address staker, IStrategy[] memory strategies, uint256 shares, bool delegateFromStakerToOperator) public {
        // sanity-filtering on fuzzed input length
        cheats.assume(strategies.length <= 64);
        // register *this contract* as an operator
        address operator = address(this);
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
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

        uint256[] memory delegatedSharesBefore = new uint256[](strategies.length);
        uint256[] memory sharesInputArray = new uint256[](strategies.length);

        // for each strategy in `strategies`, increase delegated shares by `shares`
        cheats.startPrank(address(strategyManagerMock));
        for (uint256 i = 0; i < strategies.length; ++i) {
            delegationManager.increaseDelegatedShares(staker, strategies[i], shares); 
            delegatedSharesBefore[i] = delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategies[i]);   
            // also construct an array which we'll use in another loop
            sharesInputArray[i] = shares;
        }
        cheats.stopPrank();

        // for each strategy in `strategies`, decrease delegated shares by `shares`
        cheats.startPrank(address(strategyManagerMock));
        delegationManager.decreaseDelegatedShares(delegationManager.delegatedTo(staker), strategies, sharesInputArray);
        cheats.stopPrank();

        // check shares after call to `decreaseDelegatedShares`
        for (uint256 i = 0; i < strategies.length; ++i) {
            uint256 delegatedSharesAfter = delegationManager.operatorShares(delegationManager.delegatedTo(staker), strategies[i]); 

            if (delegationManager.isDelegated(staker)) {
                require(delegatedSharesAfter == delegatedSharesBefore[i] - sharesInputArray[i], "delegated shares did not decrement correctly");
            } else {
                require(delegatedSharesAfter == delegatedSharesBefore[i], "delegated shares decremented incorrectly");
                require(delegatedSharesBefore[i] == 0, "nonzero shares delegated to zero address!");
            }
        }
    }

    // @notice Verifies that `DelegationManager.increaseDelegatedShares` reverts if not called by the StrategyManager nor EigenPodManager
    function testCannotCallIncreaseDelegatedSharesFromNonPermissionedAddress(address operator, uint256 shares) public fuzzedAddress(operator) {
        cheats.assume(operator != address(strategyManagerMock));
        cheats.assume(operator != address(eigenPodManagerMock));
        cheats.expectRevert(bytes("DelegationManager: onlyStrategyManagerOrEigenPodManager"));
        cheats.startPrank(operator);
        delegationManager.increaseDelegatedShares(operator, strategyMock, shares);
    }

    // @notice Verifies that `DelegationManager.decreaseDelegatedShares` reverts if not called by the StrategyManager nor EigenPodManager
    function testCannotCallDecreaseDelegatedSharesFromNonPermissionedAddress(
        address operator,  
        IStrategy[] memory strategies,  
        uint256[] memory shareAmounts
    ) public fuzzedAddress(operator) {
        cheats.assume(operator != address(strategyManagerMock));
        cheats.assume(operator != address(eigenPodManagerMock));
        cheats.expectRevert(bytes("DelegationManager: onlyStrategyManagerOrEigenPodManager"));
        cheats.startPrank(operator);
        delegationManager.decreaseDelegatedShares(operator, strategies, shareAmounts);
    }

    // @notice Verifies that it is not possible for a staker to delegate to an operator when the operator is frozen in EigenLayer
    function testCannotDelegateWhenOperatorIsFrozen(address operator, address staker) public fuzzedAddress(operator) fuzzedAddress(staker) {
        cheats.assume(operator != staker);
        
        cheats.startPrank(operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        delegationManager.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        cheats.stopPrank();

        slasherMock.setOperatorFrozenStatus(operator, true);
        cheats.expectRevert(bytes("DelegationManager._delegate: cannot delegate to a frozen operator"));
        cheats.startPrank(staker);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry, emptySalt);
        cheats.stopPrank();
    }

    // @notice Verifies that it is not possible for a staker to delegate to an operator when they are already delegated to an operator
    function testCannotDelegateWhenStakerHasExistingDelegation(address staker, address operator, address operator2) public
        fuzzedAddress(staker)
        fuzzedAddress(operator)
        fuzzedAddress(operator2)
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
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
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
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry, emptySalt);
    }

    // @notice Verifies that delegating is not possible when the "new delegations paused" switch is flipped
    function testCannotDelegateWhenPausedNewDelegationIsSet(address operator, address staker) public fuzzedAddress(operator) fuzzedAddress(staker) {
        cheats.startPrank(pauser);
        delegationManager.pause(1);
        cheats.stopPrank();

        cheats.startPrank(staker);
        cheats.expectRevert(bytes("Pausable: index is paused"));
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry, emptySalt);
        cheats.stopPrank();
    }

    // special event purely used in the StrategyManagerMock contract, inside of `testForceUndelegation` to verify that the correct call is made
    event ForceTotalWithdrawalCalled(address staker);

    /**
     * @notice Verifies that the `forceUndelegation` function properly calls `strategyManager.forceTotalWithdrawal`
     * @param callFromOperatorOrApprover -- calls from the operator if 'false' and the 'approver' if true
     */
    function testForceUndelegation(address staker, bytes32 salt, bool callFromOperatorOrApprover) public
        fuzzedAddress(staker)
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

        // call the `forceUndelegation` function and check that the correct calldata is forwarded by looking for an event emitted by the StrategyManagerMock contract
        cheats.startPrank(caller);
        cheats.expectEmit(true, true, true, true, address(strategyManagerMock));
        emit ForceTotalWithdrawalCalled(staker);
        (bytes32 returnValue_1, bytes32 returnValue_2) = delegationManager.forceUndelegation(staker);
        // check that the return values are empty, as specified in the mock contract
        require(returnValue_1 == bytes32(uint256(0)), "mock contract returned wrong return value");
        require(returnValue_2 == bytes32(uint256(0)), "mock contract returned wrong return value");
        cheats.stopPrank();
    }

    /**
     * @notice Verifies that the `forceUndelegation` function has proper access controls (can only be called by the operator who the `staker` has delegated
     * to or the operator's `delegationApprover`)
     */
    function testCannotCallForceUndelegationFromImproperAddress(address staker, address caller) public
        fuzzedAddress(staker)
        fuzzedAddress(caller)
    {
        address delegationApprover = cheats.addr(delegationSignerPrivateKey);
        address operator = address(this);

        // filtering since you can't delegate to yourself after registering as an operator
        cheats.assume(staker != operator);

        // filter out addresses that are actually allowed to call the function
        cheats.assume(caller != operator);
        cheats.assume(caller != delegationApprover);

        // register this contract as an operator and delegate from the staker to it
        uint256 expiry = type(uint256).max;
        testDelegateToOperatorWhoRequiresECDSASignature(staker, emptySalt, expiry);

        // try to call the `forceUndelegation` function and check for reversion
        cheats.startPrank(caller);
        cheats.expectRevert(bytes("DelegationManager.forceUndelegation: caller must be operator or their delegationApprover"));
        delegationManager.forceUndelegation(staker);
        cheats.stopPrank();
    }

    /**
     * @notice verifies that `DelegationManager.forceUndelegation` reverts if trying to undelegate an operator from themselves
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

        // try to call the `forceUndelegation` function and check for reversion
        cheats.startPrank(caller);
        cheats.expectRevert(bytes("DelegationManager.forceUndelegation: operators cannot be force-undelegated"));
        delegationManager.forceUndelegation(operator);
        cheats.stopPrank();
    }

    /**
     * @notice Verifies that the reversion occurs when trying to reuse an 'approverSalt'
     */
    function test_Revert_WhenTryingToReuseSalt(address staker_one, address staker_two, bytes32 salt) public
        fuzzedAddress(staker_one)
        fuzzedAddress(staker_two)
    {
        // address delegationApprover = cheats.addr(delegationSignerPrivateKey);
        address operator = address(this);

        // filtering since you can't delegate to yourself after registering as an operator
        cheats.assume(staker_one != operator);
        cheats.assume(staker_two != operator);

        // register this contract as an operator and delegate from `staker_one` to it, using the `salt`
        uint256 expiry = type(uint256).max;
        testDelegateToOperatorWhoRequiresECDSASignature(staker_one, salt, expiry);

        // calculate the delegationSigner's signature
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry =
            _getApproverSignature(delegationSignerPrivateKey, staker_two, operator, salt, expiry);

        // try to delegate to the operator from `staker_two`, and verify that the call reverts for the proper reason (trying to reuse a salt)
        cheats.startPrank(staker_two);
        cheats.expectRevert(bytes("DelegationManager._delegate: approverSalt already spent"));
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, salt);
        cheats.stopPrank();
    }

    /**
     * @notice internal function for calculating a signature from the delegationSigner corresponding to `_delegationSignerPrivateKey`, approving
     * the `staker` to delegate to `operator`, with the specified `salt`, and expiring at `expiry`.
     */
    function _getApproverSignature(uint256 _delegationSignerPrivateKey, address staker, address operator, bytes32 salt, uint256 expiry)
        internal view returns (IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry)
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
        internal view returns (IDelegationManager.SignatureWithExpiry memory stakerSignatureAndExpiry)
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