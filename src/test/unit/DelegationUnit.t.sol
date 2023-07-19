// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";

import "forge-std/Test.sol";

import "../mocks/StrategyManagerMock.sol";
import "../mocks/SlasherMock.sol";
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

    uint256 hardhatAccountZeroPrivateKey = uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);
    uint256 stakerPrivateKey = uint256(123456789);

    // @notice Emitted when a new operator registers in EigenLayer and provides their OperatorDetails.
    event OperatorRegistered(address indexed operator, IDelegationManager.OperatorDetails operatorDetails);

    // @notice Emitted when an operator updates their OperatorDetails to @param newOperatorDetails
    event OperatorDetailsModified(address indexed operator, IDelegationManager.OperatorDetails newOperatorDetails);

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

        delegationManagerImplementation = new DelegationManager(strategyManagerMock, slasherMock);

        cheats.startPrank(eigenLayerProxyAdmin.owner());
        eigenLayerProxyAdmin.upgrade(TransparentUpgradeableProxy(payable(address(delegationManager))), address(delegationManagerImplementation));
        cheats.stopPrank();
        
        delegationManager.initialize(address(this), eigenLayerPauserReg, 0);

        strategyImplementation = new StrategyBase(strategyManager);

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
    }

    /**
     * @notice `operator` registers via calling `DelegationManager.registerAsOperator(operatorDetails)`
     * Should be able to set any parameters, other than setting their `earningsReceiver` to the zero address or too high value for `stakerOptOutWindowBlocks`
     * The set parameters should match the desired parameters (correct storage update)
     * Operator becomes delegated to themselves
     * Properly emits events – especially the `OperatorRegistered` event, but also `StakerDelegated` & `OperatorDetailsModified` events
     * Reverts appropriately if operator was already delegated to someone (including themselves, i.e. they were already an operator)
     * @param operator and @param operatorDetails are fuzzed inputs
     */
    function testRegisterAsOperator(address operator, IDelegationManager.OperatorDetails memory operatorDetails) public {
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

        delegationManager.registerAsOperator(operatorDetails);

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
        delegationManager.registerAsOperator(operatorDetails);
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
        delegationManager.registerAsOperator(operatorDetails);
        cheats.stopPrank();
    }

    // @notice Verifies that someone cannot successfully call `DelegationManager.registerAsOperator(operatorDetails)` again after registering for the first time
    function testCannotRegisterAsOperatorMultipleTimes(address operator, IDelegationManager.OperatorDetails memory operatorDetails) public 
        filterFuzzedAddressInputs(operator)
    {
        testRegisterAsOperator(operator, operatorDetails);
        cheats.startPrank(operator);
        cheats.expectRevert(bytes("DelegationManager.registerAsOperator: operator has already registered"));
        delegationManager.registerAsOperator(operatorDetails);
        cheats.stopPrank();
    }

    // @notice Verifies that a staker who is actively delegated to an operator cannot register as an operator (without first undelegating, at least)
    function testCannotRegisterAsOperatorWhileDelegated(address staker, IDelegationManager.OperatorDetails memory operatorDetails) public 
        filterFuzzedAddressInputs(staker)
    {
        // filter out disallowed stakerOptOutWindowBlocks values
        cheats.assume(operatorDetails.stakerOptOutWindowBlocks <= delegationManager.MAX_STAKER_OPT_OUT_WINDOW_BLOCKS());

        // register *this contract* as an operator
        address operator = address(this);
        IDelegationManager.OperatorDetails memory _operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, _operatorDetails);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateTo(operator, approverSignatureAndExpiry);        

        cheats.expectRevert(bytes("DelegationManager._delegate: staker is already actively delegated"));
        delegationManager.registerAsOperator(operatorDetails);

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
        address operator = address(this);
        testRegisterAsOperator(operator, initialOperatorDetails);
        // filter out zero address since people can't set their earningsReceiver address to the zero address (special test case to verify)
        cheats.assume(modifiedOperatorDetails.earningsReceiver != address(0));

        cheats.startPrank(operator);

        // either it fails for trying to set the stakerOptOutWindowBlocks
        if (modifiedOperatorDetails.stakerOptOutWindowBlocks > delegationManager.MAX_STAKER_OPT_OUT_WINDOW_BLOCKS()) {
            cheats.expectRevert(bytes("DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be > MAX_STAKER_OPT_OUT_WINDOW_BLOCKS"));
            delegationManager.modifyOperatorDetails(modifiedOperatorDetails);
        // or the transition is allowed,
        } else if (modifiedOperatorDetails.stakerOptOutWindowBlocks >= initialOperatorDetails.stakerOptOutWindowBlocks) {
            cheats.expectEmit(true, true, true, true, address(delegationManager));
            emit OperatorDetailsModified(operator, modifiedOperatorDetails);
            delegationManager.modifyOperatorDetails(modifiedOperatorDetails);

            require(modifiedOperatorDetails.earningsReceiver == delegationManager.earningsReceiver(operator), "earningsReceiver not set correctly");
            require(modifiedOperatorDetails.delegationApprover == delegationManager.delegationApprover(operator), "delegationApprover not set correctly");
            require(modifiedOperatorDetails.stakerOptOutWindowBlocks == delegationManager.stakerOptOutWindowBlocks(operator), "stakerOptOutWindowBlocks not set correctly");
            require(delegationManager.delegatedTo(operator) == operator, "operator not delegated to self");
        // or else the transition is disallowed
        } else {
            cheats.expectRevert(bytes("DelegationManager._setOperatorDetails: stakerOptOutWindowBlocks cannot be decreased"));
            delegationManager.modifyOperatorDetails(modifiedOperatorDetails);
        }

        cheats.stopPrank();
    }

    /**
     * @notice Verifies that an operator cannot modify their `earningsReceiver` address to set it to the zero address
     * @dev This is an important check since we check `earningsReceiver != address(0)` to check if an address is an operator!
     */
    function testCannotModifyEarningsReceiverAddressToZeroAddress() public {
        // register *this contract* as an operator
        address operator = address(this);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails);

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
    function testDelegateToOperatorWhoAcceptsAllStakers(address staker, IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry) public 
        filterFuzzedAddressInputs(staker)
    {
        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails);

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, operator);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry);        
        cheats.stopPrank();

        require(delegationManager.isDelegated(staker), "staker not delegated correctly");
        require(delegationManager.delegatedTo(staker) == operator, "staker delegated to the wrong address");
        require(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");
    }

    /**
     * @notice Delegates from `staker` to an operator, then verifies that the `staker` cannot delegate to another `operator` (at least without first undelegating)
     */
    function testCannotDelegateWhileDelegated(address staker, address operator, IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry) public 
        filterFuzzedAddressInputs(staker)
        filterFuzzedAddressInputs(operator)
    {
        // delegate from the staker to an operator
        testDelegateToOperatorWhoAcceptsAllStakers(staker, approverSignatureAndExpiry);

        // register another operator
        // filter out this contract, since we already register it as an operator in the above step
        cheats.assume(operator != address(this));
        IDelegationManager.OperatorDetails memory _operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, _operatorDetails);

        // try to delegate again and check that the call reverts
        cheats.startPrank(staker);
        cheats.expectRevert(bytes("DelegationManager._delegate: staker is already actively delegated"));
        delegationManager.delegateTo(operator, approverSignatureAndExpiry);        
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
        delegationManager.delegateTo(operator, approverSignatureAndExpiry);        
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
    function testDelegateToOperatorWhoRequiresECDSASignature(address staker, uint256 expiry) public 
        filterFuzzedAddressInputs(staker)
    {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);

        uint256 delegationApproverPrivateKey = hardhatAccountZeroPrivateKey;
        address delegationApprover = cheats.addr(delegationApproverPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails);

        // calculate the signature
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        approverSignatureAndExpiry.expiry = expiry;
        uint256 currentApproverNonce = delegationManager.delegationApproverNonce(delegationApprover);
        {
            bytes32 digestHash = delegationManager.calculateApproverDigestHash(operator, expiry);
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(delegationApproverPrivateKey, digestHash);
            approverSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, operator);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry);        
        cheats.stopPrank();

        require(delegationManager.isDelegated(staker), "staker not delegated correctly");
        require(delegationManager.delegatedTo(staker) == operator, "staker delegated to the wrong address");
        require(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        if (staker == operator || staker == delegationManager.delegationApprover(operator)) {
            require(delegationManager.delegationApproverNonce(delegationManager.delegationApprover(operator)) == currentApproverNonce,
                "delegationApprover nonce incremented inappropriately");
        } else {
            require(delegationManager.delegationApproverNonce(delegationManager.delegationApprover(operator)) == currentApproverNonce + 1,
                "delegationApprover nonce did not increment");
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

        uint256 delegationApproverPrivateKey = hardhatAccountZeroPrivateKey;
        address delegationApprover = cheats.addr(delegationApproverPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails);

        // calculate the signature
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        approverSignatureAndExpiry.expiry = expiry;
        {
            bytes32 digestHash = delegationManager.calculateApproverDigestHash(operator, expiry);
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(delegationApproverPrivateKey, digestHash);
            // mess up the signature by flipping v's parity
            v = (v == 27 ? 28 : 27);
            approverSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }

        // try to delegate from the `staker` to the operator, and check reversion
        cheats.startPrank(staker);
        cheats.expectRevert(bytes("EIP1271SignatureUtils.checkSignature_EIP1271: signature not from signer"));
        delegationManager.delegateTo(operator, approverSignatureAndExpiry);        
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

        uint256 delegationApproverPrivateKey = hardhatAccountZeroPrivateKey;
        address delegationApprover = cheats.addr(delegationApproverPrivateKey);

        // register *this contract* as an operator
        address operator = address(this);
        // filter inputs, since this will fail when the staker is already registered as an operator
        cheats.assume(staker != operator);

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: delegationApprover,
            stakerOptOutWindowBlocks: 0
        });
        testRegisterAsOperator(operator, operatorDetails);

        // calculate the signature
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        approverSignatureAndExpiry.expiry = expiry;
        {
            bytes32 digestHash = delegationManager.calculateApproverDigestHash(operator, expiry);
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(delegationApproverPrivateKey, digestHash);
            approverSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectRevert(bytes("DelegationManager._delegate: approver signature expired"));
        delegationManager.delegateTo(operator, approverSignatureAndExpiry);        
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
    function testDelegateToOperatorWhoRequiresEIP1271Signature(address staker, uint256 expiry) public 
        filterFuzzedAddressInputs(staker)
    {
        // filter to only valid `expiry` values
        cheats.assume(expiry >= block.timestamp);

        uint256 delegationSignerPrivateKey = hardhatAccountZeroPrivateKey;
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
        testRegisterAsOperator(operator, operatorDetails);

        // calculate the signature
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        approverSignatureAndExpiry.expiry = expiry;
        uint256 currentApproverNonce = delegationManager.delegationApproverNonce(delegationManager.delegationApprover(operator));
        {
            bytes32 digestHash = delegationManager.calculateApproverDigestHash(operator, expiry);
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(delegationSignerPrivateKey, digestHash);
            approverSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }

        // delegate from the `staker` to the operator
        cheats.startPrank(staker);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, operator);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry);        
        cheats.stopPrank();

        require(delegationManager.isDelegated(staker), "staker not delegated correctly");
        require(delegationManager.delegatedTo(staker) == operator, "staker delegated to the wrong address");
        require(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // check that the nonce incremented appropriately
        if (staker == operator || staker == delegationManager.delegationApprover(operator)) {
            require(delegationManager.delegationApproverNonce(delegationManager.delegationApprover(operator)) == currentApproverNonce,
                "delegationApprover nonce incremented inappropriately");
        } else {
            require(delegationManager.delegationApproverNonce(delegationManager.delegationApprover(operator)) == currentApproverNonce + 1,
                "delegationApprover nonce did not increment");
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
        testRegisterAsOperator(operator, operatorDetails);

        // create the signature struct
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        approverSignatureAndExpiry.expiry = expiry;

        // try to delegate from the `staker` to the operator, and check reversion
        cheats.startPrank(staker);
        // because the contract returns the wrong amount of data, we get a low-level "EvmError: Revert" message here rather than the error message bubbling up
        // cheats.expectRevert(bytes("EIP1271SignatureUtils.checkSignature_EIP1271: ERC1271 signature verification failed"));
        cheats.expectRevert();
        delegationManager.delegateTo(operator, approverSignatureAndExpiry);        
        cheats.stopPrank();
    }

    /**
     * @notice `staker` becomes delegated to an operator who does not require any signature verification (i.e. the operator’s `delegationApprover` address is set to the zero address)
     * via the `caller calling `DelegationManager.delegateToBySignature`
     * The function should pass with any `operatorSignature` input (since it should be unused)
     * The function should pass only with a valid `stakerSignatureAndExpiry` input
     * Properly emits a `StakerDelegated` event
     * Staker is correctly delegated after the call (i.e. correct storage update)
     * Reverts if the staker is already delegated (to the operator or to anyone else)
     * Reverts if the ‘operator’ is not actually registered as an operator
     */
    function testDelegateBySignatureToOperatorWhoWhoAcceptsAllStakers(address caller, uint256 expiry) public 
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
        testRegisterAsOperator(operator, operatorDetails);

        // calculate the staker signature
        IDelegationManager.SignatureWithExpiry memory stakerSignatureAndExpiry;
        stakerSignatureAndExpiry.expiry = expiry;
        uint256 currentStakerNonce = delegationManager.stakerNonce(staker);
        {
            bytes32 digestHash = delegationManager.calculateStakerDigestHash(staker, operator, expiry);
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(stakerPrivateKey, digestHash);
            stakerSignatureAndExpiry.signature = abi.encodePacked(r, s, v);
        }

        // delegate from the `staker` to the operator, via having the `caller` call `DelegationManager.delegateToBySignature`
        cheats.startPrank(caller);
        cheats.expectEmit(true, true, true, true, address(delegationManager));
        emit StakerDelegated(staker, operator);
        // use an empty approver signature input since none is needed
        IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry;
        delegationManager.delegateToBySignature(staker, operator, stakerSignatureAndExpiry, approverSignatureAndExpiry);        
        cheats.stopPrank();

        // check all the delegation status changes
        require(delegationManager.isDelegated(staker), "staker not delegated correctly");
        require(delegationManager.delegatedTo(staker) == operator, "staker delegated to the wrong address");
        require(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");

        // check that the nonce incremented appropriately
        require(delegationManager.stakerNonce(staker) == currentStakerNonce + 1,
            "delegationApprover nonce did not increment");
    }


    // function testDelegateToOperatorWhoAcceptsAllStakers(address staker, IDelegationManager.SignatureWithExpiry memory approverSignatureAndExpiry) public 
    //     filterFuzzedAddressInputs(staker)
    // {
    //     // register *this contract* as an operator
    //     address operator = address(this);
    //     // filter inputs, since this will fail when the staker is already registered as an operator
    //     cheats.assume(staker != operator);

    //     IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
    //         earningsReceiver: operator,
    //         delegationApprover: address(0),
    //         stakerOptOutWindowBlocks: 0
    //     });
    //     testRegisterAsOperator(operator, operatorDetails);

    //     // delegate from the `staker` to the operator
    //     cheats.startPrank(staker);
    //     cheats.expectEmit(true, true, true, true, address(delegationManager));
    //     emit StakerDelegated(staker, operator);
    //     delegationManager.delegateTo(operator, approverSignatureAndExpiry);        
    //     cheats.stopPrank();

    //     require(delegationManager.isDelegated(staker), "staker not delegated correctly");
    //     require(delegationManager.delegatedTo(staker) == operator, "staker delegated to the wrong address");
    //     require(!delegationManager.isOperator(staker), "staker incorrectly registered as operator");
    // }
















    // @notice Verifies that the DelegationManager cannot be iniitalized multiple times
    function testCannotReinitializeDelegationManager() public {
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        delegationManager.initialize(address(this), eigenLayerPauserReg, 0);
    }

    function testBadECDSASignatureExpiry(address staker, address operator, uint256 expiry, bytes memory signature) public{
        cheats.assume(expiry < block.timestamp);
        cheats.expectRevert(bytes("DelegationManager.delegateToBySignature: staker signature expired"));
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry = IDelegationManager.SignatureWithExpiry({
            signature: signature,
            expiry: expiry
        });
        delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry);
    }

    function testUndelegateFromNonStrategyManagerAddress(address undelegator) public fuzzedAddress(undelegator) {
        cheats.assume(undelegator != address(strategyManagerMock));
        cheats.expectRevert(bytes("onlyStrategyManager"));
        cheats.startPrank(undelegator);
        delegationManager.undelegate(address(this));
    }

    function testUndelegateByOperatorFromThemselves(address operator) public fuzzedAddress(operator) {
        cheats.startPrank(operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        delegationManager.registerAsOperator(operatorDetails);
        cheats.stopPrank();
        cheats.expectRevert(bytes("DelegationManager.undelegate: operators cannot undelegate from themselves"));
        
        cheats.startPrank(address(strategyManagerMock));
        delegationManager.undelegate(operator);
        cheats.stopPrank();
    }

    function testIncreaseDelegatedSharesFromNonStrategyManagerAddress(address operator, uint256 shares) public fuzzedAddress(operator) {
        cheats.assume(operator != address(strategyManagerMock));
        cheats.expectRevert(bytes("onlyStrategyManager"));
        cheats.startPrank(operator);
        delegationManager.increaseDelegatedShares(operator, strategyMock, shares);
    }

    function testDecreaseDelegatedSharesFromNonStrategyManagerAddress(
        address operator,  
        IStrategy[] memory strategies,  
        uint256[] memory shareAmounts
    ) public fuzzedAddress(operator) {
        cheats.assume(operator != address(strategyManagerMock));
        cheats.expectRevert(bytes("onlyStrategyManager"));
        cheats.startPrank(operator);
        delegationManager.decreaseDelegatedShares(operator, strategies, shareAmounts);
    }

    function testDelegateWhenOperatorIsFrozen(address operator, address staker) public fuzzedAddress(operator) fuzzedAddress(staker) {
        cheats.assume(operator != staker);
        
        cheats.startPrank(operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        delegationManager.registerAsOperator(operatorDetails);
        cheats.stopPrank();

        slasherMock.setOperatorFrozenStatus(operator, true);
        cheats.expectRevert(bytes("DelegationManager._delegate: cannot delegate to a frozen operator"));
        cheats.startPrank(staker);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry);
        cheats.stopPrank();
    }

    function testDelegateWhenStakerHasExistingDelegation(address staker, address operator, address operator2) public
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
        delegationManager.registerAsOperator(operatorDetails);
        cheats.stopPrank();

        cheats.startPrank(operator2);
        delegationManager.registerAsOperator(operatorDetails);
        cheats.stopPrank();

        cheats.startPrank(staker);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry);
        cheats.stopPrank();

        cheats.startPrank(staker);
        cheats.expectRevert(bytes("DelegationManager._delegate: staker is already actively delegated"));
        delegationManager.delegateTo(operator2, signatureWithExpiry);
        cheats.stopPrank();
    }

    function testDelegationToUnregisteredOperator(address operator) public{
        cheats.expectRevert(bytes("DelegationManager._delegate: operator is not registered in EigenLayer"));
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry);
    }

    function testDelegationWhenPausedNewDelegationIsSet(address operator, address staker) public fuzzedAddress(operator) fuzzedAddress(staker) {
        cheats.startPrank(pauser);
        delegationManager.pause(1);
        cheats.stopPrank();

        cheats.startPrank(staker);
        cheats.expectRevert(bytes("Pausable: index is paused"));
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegationManager.delegateTo(operator, signatureWithExpiry);
        cheats.stopPrank();
    }
}