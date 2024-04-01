// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "src/contracts/interfaces/ISignatureUtils.sol";

import "../test/EigenLayerTestHelper.t.sol";

contract DelegationTests is EigenLayerTestHelper {
    uint256 public PRIVATE_KEY = 420;

    uint32 serveUntil = 100;

    address public registryCoordinator = address(uint160(uint256(keccak256("registryCoordinator"))));
    uint8 defaultQuorumNumber = 0;
    bytes32 defaultOperatorId = bytes32(uint256(0));

    modifier fuzzedAmounts(uint256 ethAmount, uint256 eigenAmount) {
        cheats.assume(ethAmount >= 0 && ethAmount <= 1e18);
        cheats.assume(eigenAmount >= 0 && eigenAmount <= 1e18);
        _;
    }

    function setUp() public virtual override {
        EigenLayerDeployer.setUp();
    }

    /// @notice testing if an operator can register to themselves.
    function testSelfOperatorRegister() public {
        _testRegisterAdditionalOperator(0);
    }

    /// @notice testing if an operator can delegate to themselves.
    /// @param sender is the address of the operator.
    function testSelfOperatorDelegate(address sender) public {
        cheats.assume(sender != address(0));
        cheats.assume(sender != address(eigenLayerProxyAdmin));
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: sender,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        _testRegisterAsOperator(sender, operatorDetails);
    }

    function testTwoSelfOperatorsRegister() public {
        _testRegisterAdditionalOperator(0);
        _testRegisterAdditionalOperator(1);
    }

    /// @notice registers a fixed address as a delegate, delegates to it from a second address,
    ///         and checks that the delegate's voteWeights increase properly
    /// @param operator is the operator being delegated to.
    /// @param staker is the staker delegating stake to the operator.
    function testDelegation(
        address operator,
        address staker,
        uint96 ethAmount,
        uint96 eigenAmount
    ) public fuzzedAddress(operator) fuzzedAddress(staker) fuzzedAmounts(ethAmount, eigenAmount) {
        cheats.assume(staker != operator);
        // base strategy will revert if these amounts are too small on first deposit
        cheats.assume(ethAmount >= 1);
        cheats.assume(eigenAmount >= 2);

        // Set weights ahead of the helper function call
        bytes memory quorumNumbers = new bytes(2);
        quorumNumbers[0] = bytes1(uint8(0));
        quorumNumbers[0] = bytes1(uint8(1));
        _testDelegation(operator, staker, ethAmount, eigenAmount);
    }

    /// @notice tests that a when an operator is delegated to, that delegation is properly accounted for.
    function testDelegationReceived(
        address _operator,
        address staker,
        uint64 ethAmount,
        uint64 eigenAmount
    ) public fuzzedAddress(_operator) fuzzedAddress(staker) fuzzedAmounts(ethAmount, eigenAmount) {
        cheats.assume(staker != _operator);
        cheats.assume(ethAmount >= 1);
        cheats.assume(eigenAmount >= 2);

        // use storage to solve stack-too-deep
        operator = _operator;

        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        if (!delegation.isOperator(operator)) {
            _testRegisterAsOperator(operator, operatorDetails);
        }

        uint256 amountBefore = delegation.operatorShares(operator, wethStrat);

        //making additional deposits to the  strategies
        assertTrue(!delegation.isDelegated(staker) == true, "testDelegation: staker is not delegate");
        _testDepositWeth(staker, ethAmount);
        _testDepositEigen(staker, eigenAmount);
        _testDelegateToOperator(staker, operator);
        assertTrue(delegation.isDelegated(staker) == true, "testDelegation: staker is not delegate");

        (/*IStrategy[] memory updatedStrategies*/, uint256[] memory updatedShares) = strategyManager.getDeposits(staker);

        {
            IStrategy _strat = wethStrat;
            // IStrategy _strat = strategyManager.stakerStrats(staker, 0);
            assertTrue(address(_strat) != address(0), "stakerStrats not updated correctly");

            assertTrue(
                delegation.operatorShares(operator, _strat) - updatedShares[0] == amountBefore,
                "ETH operatorShares not updated correctly"
            );

            cheats.startPrank(address(strategyManager));

            IDelegationManager.OperatorDetails memory expectedOperatorDetails = delegation.operatorDetails(operator);
            assertTrue(
                keccak256(abi.encode(expectedOperatorDetails)) == keccak256(abi.encode(operatorDetails)),
                "failed to set correct operator details"
            );
        }
    }

    /// @notice tests that a when an operator is undelegated from, that the staker is properly classified as undelegated.
    function testUndelegation(
        address operator,
        address staker,
        uint96 ethAmount,
        uint96 eigenAmount
    ) public fuzzedAddress(operator) fuzzedAddress(staker) fuzzedAmounts(ethAmount, eigenAmount) {
        cheats.assume(staker != operator);
        // base strategy will revert if these amounts are too small on first deposit
        cheats.assume(ethAmount >= 1);
        cheats.assume(eigenAmount >= 1);
        _testDelegation(operator, staker, ethAmount, eigenAmount);

        (IStrategy[] memory strategyArray, uint256[] memory shareAmounts) = strategyManager.getDeposits(staker);
        uint256[] memory strategyIndexes = new uint256[](strategyArray.length);

        // withdraw shares
        _testQueueWithdrawal(staker, strategyIndexes, strategyArray, shareAmounts, staker /*withdrawer*/);

        cheats.startPrank(staker);
        delegation.undelegate(staker);
        cheats.stopPrank();

        require(delegation.delegatedTo(staker) == address(0), "undelegation unsuccessful");
    }

    /// @notice tests delegation from a staker to operator via ECDSA signature.
    function testDelegateToBySignature(
        address operator,
        uint96 ethAmount,
        uint96 eigenAmount,
        uint256 expiry
    ) public fuzzedAddress(operator) {
        address staker = cheats.addr(PRIVATE_KEY);
        _registerOperatorAndDepositFromStaker(operator, staker, ethAmount, eigenAmount);

        uint256 nonceBefore = delegation.stakerNonce(staker);

        bytes32 structHash = keccak256(
            abi.encode(delegation.STAKER_DELEGATION_TYPEHASH(), staker, operator, nonceBefore, expiry)
        );
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", delegation.domainSeparator(), structHash));

        bytes memory signature;
        {
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(PRIVATE_KEY, digestHash);
            signature = abi.encodePacked(r, s, v);
        }

        if (expiry < block.timestamp) {
            cheats.expectRevert("DelegationManager.delegateToBySignature: staker signature expired");
        }
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry = ISignatureUtils.SignatureWithExpiry({
            signature: signature,
            expiry: expiry
        });
        delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, bytes32(0));
        if (expiry >= block.timestamp) {
            assertTrue(delegation.isDelegated(staker) == true, "testDelegation: staker is not delegate");
            assertTrue(nonceBefore + 1 == delegation.stakerNonce(staker), "nonce not incremented correctly");
            assertTrue(delegation.delegatedTo(staker) == operator, "staker delegated to wrong operator");
        }
    }

    /// @notice tries delegating using a signature and an EIP 1271 compliant wallet
    function testDelegateToBySignature_WithContractWallet_Successfully(
        address operator,
        uint96 ethAmount,
        uint96 eigenAmount
    ) public fuzzedAddress(operator) {
        address staker = cheats.addr(PRIVATE_KEY);

        // deploy ERC1271WalletMock for staker to use
        cheats.startPrank(staker);
        ERC1271WalletMock wallet = new ERC1271WalletMock(staker);
        cheats.stopPrank();
        staker = address(wallet);

        _registerOperatorAndDepositFromStaker(operator, staker, ethAmount, eigenAmount);

        uint256 nonceBefore = delegation.stakerNonce(staker);

        bytes32 structHash = keccak256(
            abi.encode(delegation.STAKER_DELEGATION_TYPEHASH(), staker, operator, nonceBefore, type(uint256).max)
        );
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", delegation.domainSeparator(), structHash));

        bytes memory signature;
        {
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(PRIVATE_KEY, digestHash);
            signature = abi.encodePacked(r, s, v);
        }

        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry = ISignatureUtils.SignatureWithExpiry({
            signature: signature,
            expiry: type(uint256).max
        });
        delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, bytes32(0));
        assertTrue(delegation.isDelegated(staker) == true, "testDelegation: staker is not delegate");
        assertTrue(nonceBefore + 1 == delegation.stakerNonce(staker), "nonce not incremented correctly");
        assertTrue(delegation.delegatedTo(staker) == operator, "staker delegated to wrong operator");
    }

    ///  @notice tries delegating using a signature and an EIP 1271 compliant wallet, *but* providing a bad signature
    function testDelegateToBySignature_WithContractWallet_BadSignature(
        address operator,
        uint96 ethAmount,
        uint96 eigenAmount
    ) public fuzzedAddress(operator) {
        address staker = cheats.addr(PRIVATE_KEY);

        // deploy ERC1271WalletMock for staker to use
        cheats.startPrank(staker);
        ERC1271WalletMock wallet = new ERC1271WalletMock(staker);
        cheats.stopPrank();
        staker = address(wallet);

        _registerOperatorAndDepositFromStaker(operator, staker, ethAmount, eigenAmount);

        uint256 nonceBefore = delegation.stakerNonce(staker);

        bytes32 structHash = keccak256(
            abi.encode(delegation.STAKER_DELEGATION_TYPEHASH(), staker, operator, nonceBefore, type(uint256).max)
        );
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", delegation.domainSeparator(), structHash));

        bytes memory signature;
        {
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(PRIVATE_KEY, digestHash);
            // mess up the signature by flipping v's parity
            v = (v == 27 ? 28 : 27);
            signature = abi.encodePacked(r, s, v);
        }

        cheats.expectRevert(
            bytes("EIP1271SignatureUtils.checkSignature_EIP1271: ERC1271 signature verification failed")
        );
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry = ISignatureUtils.SignatureWithExpiry({
            signature: signature,
            expiry: type(uint256).max
        });
        delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, bytes32(0));
    }

    /// @notice  tries delegating using a wallet that does not comply with EIP 1271
    function testDelegateToBySignature_WithContractWallet_NonconformingWallet(
        address operator,
        uint96 ethAmount,
        uint96 eigenAmount,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) public fuzzedAddress(operator) {
        address staker = cheats.addr(PRIVATE_KEY);

        // deploy non ERC1271-compliant wallet for staker to use
        cheats.startPrank(staker);
        ERC1271MaliciousMock wallet = new ERC1271MaliciousMock();
        cheats.stopPrank();
        staker = address(wallet);

        _registerOperatorAndDepositFromStaker(operator, staker, ethAmount, eigenAmount);

        cheats.assume(staker != operator);

        bytes memory signature = abi.encodePacked(r, s, v);

        cheats.expectRevert();
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry = ISignatureUtils.SignatureWithExpiry({
            signature: signature,
            expiry: type(uint256).max
        });
        delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, bytes32(0));
    }

    /// @notice tests delegation to EigenLayer via an ECDSA signatures with invalid signature
    /// @param operator is the operator being delegated to.
    function testDelegateToByInvalidSignature(
        address operator,
        uint96 ethAmount,
        uint96 eigenAmount,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) public fuzzedAddress(operator) fuzzedAmounts(ethAmount, eigenAmount) {
        address staker = cheats.addr(PRIVATE_KEY);
        _registerOperatorAndDepositFromStaker(operator, staker, ethAmount, eigenAmount);

        bytes memory signature = abi.encodePacked(r, s, v);

        cheats.expectRevert();
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry = ISignatureUtils.SignatureWithExpiry({
            signature: signature,
            expiry: type(uint256).max
        });
        delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, bytes32(0));
    }

    /// @notice This function tests to ensure that a delegation contract
    ///         cannot be intitialized multiple times
    function testCannotInitMultipleTimesDelegation() public cannotReinit {
        //delegation has already been initialized in the Deployer test contract
        delegation.initialize(
            address(this),
            eigenLayerPauserReg,
            0,
            minWithdrawalDelayBlocks,
            initializeStrategiesToSetDelayBlocks,
            initializeWithdrawalDelayBlocks
        );
    }

    /// @notice This function tests to ensure that a you can't register as a delegate multiple times
    /// @param operator is the operator being delegated to.
    function testRegisterAsOperatorMultipleTimes(address operator) public fuzzedAddress(operator) {
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        _testRegisterAsOperator(operator, operatorDetails);
        cheats.expectRevert(bytes("DelegationManager.registerAsOperator: operator has already registered"));
        _testRegisterAsOperator(operator, operatorDetails);
    }

    /// @notice This function tests to ensure that a staker cannot delegate to an unregistered operator
    /// @param delegate is the unregistered operator
    function testDelegationToUnregisteredDelegate(address delegate) public fuzzedAddress(delegate) {
        //deposit into 1 strategy for getOperatorAddress(1), who is delegating to the unregistered operator
        _testDepositStrategies(getOperatorAddress(1), 1e18, 1);
        _testDepositEigen(getOperatorAddress(1), 1e18);

        cheats.expectRevert(bytes("DelegationManager._delegate: operator is not registered in EigenLayer"));
        cheats.startPrank(getOperatorAddress(1));
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry;
        delegation.delegateTo(delegate, signatureWithExpiry, bytes32(0));
        cheats.stopPrank();
    }

    /// @notice This function tests to ensure that a delegation contract
    ///         cannot be intitialized multiple times, test with different caller addresses
    function testCannotInitMultipleTimesDelegation(address _attacker) public {
        cheats.assume(_attacker != address(eigenLayerProxyAdmin));
        //delegation has already been initialized in the Deployer test contract
        vm.prank(_attacker);
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        delegation.initialize(
            _attacker,
            eigenLayerPauserReg,
            0,
            0, // minWithdrawalDelayBLocks
            initializeStrategiesToSetDelayBlocks,
            initializeWithdrawalDelayBlocks
        );
    }

    /// @notice This function tests that the earningsReceiver cannot be set to address(0)
    function testCannotSetEarningsReceiverToZeroAddress() public {
        cheats.expectRevert(
            bytes("DelegationManager._setOperatorDetails: cannot set `earningsReceiver` to zero address")
        );
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: address(0),
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        string memory emptyStringForMetadataURI;
        delegation.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
    }

    /// @notice This function tests to ensure that an address can only call registerAsOperator() once
    function testCannotRegisterAsOperatorTwice(
        address _operator,
        address _dt
    ) public fuzzedAddress(_operator) fuzzedAddress(_dt) {
        vm.assume(_dt != address(0));
        vm.startPrank(_operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: msg.sender,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        string memory emptyStringForMetadataURI;
        delegation.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        vm.expectRevert("DelegationManager.registerAsOperator: operator has already registered");
        delegation.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        cheats.stopPrank();
    }

    /// @notice this function checks that you can only delegate to an address that is already registered.
    function testDelegateToInvalidOperator(
        address _staker,
        address _unregisteredoperator
    ) public fuzzedAddress(_staker) {
        vm.startPrank(_staker);
        cheats.expectRevert(bytes("DelegationManager._delegate: operator is not registered in EigenLayer"));
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry;
        delegation.delegateTo(_unregisteredoperator, signatureWithExpiry, bytes32(0));
        cheats.expectRevert(bytes("DelegationManager._delegate: operator is not registered in EigenLayer"));
        delegation.delegateTo(_staker, signatureWithExpiry, bytes32(0));
        cheats.stopPrank();
    }

    function testUndelegate_SigP_Version(address _operator, address _staker, address _dt) public {
        vm.assume(_operator != address(0));
        vm.assume(_staker != address(0));
        vm.assume(_operator != _staker);
        vm.assume(_dt != address(0));
        vm.assume(_operator != address(eigenLayerProxyAdmin));
        vm.assume(_staker != address(eigenLayerProxyAdmin));

        // setup delegation
        vm.prank(_operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: _dt,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        string memory emptyStringForMetadataURI;
        delegation.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        vm.prank(_staker);
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry;
        delegation.delegateTo(_operator, signatureWithExpiry, bytes32(0));

        // operators cannot undelegate from themselves
        vm.prank(_operator);
        cheats.expectRevert(bytes("DelegationManager.undelegate: operators cannot be undelegated"));
        delegation.undelegate(_operator);

        // assert still delegated
        assertTrue(delegation.isDelegated(_staker));
        assertTrue(delegation.isOperator(_operator));

        // _staker *can* undelegate themselves
        vm.prank(_staker);
        delegation.undelegate(_staker);

        // assert undelegated
        assertTrue(!delegation.isDelegated(_staker));
        assertTrue(delegation.isOperator(_operator));
    }

    /**************************************
     * 
     *  Withdrawals Tests with StrategyManager, using actual SM contract instead of Mock to test
     * 
     **************************************/

    
    // function testQueueWithdrawalRevertsMismatchedSharesAndStrategyArrayLength() external {
    //     IStrategy[] memory strategyArray = new IStrategy[](1);
    //     uint256[] memory shareAmounts = new uint256[](2);

    //     {
    //         strategyArray[0] = eigenPodManagerMock.beaconChainETHStrategy();
    //         shareAmounts[0] = 1;
    //         shareAmounts[1] = 1;
    //     }

    //     IDelegationManager.QueuedWithdrawalParams[] memory params = new IDelegationManager.QueuedWithdrawalParams[](1);
        
    //     params[0] = IDelegationManager.QueuedWithdrawalParams({
    //         strategies: strategyArray,
    //         shares: shareAmounts,
    //         withdrawer: address(this)
    //     });

    //     cheats.expectRevert(bytes("DelegationManager.queueWithdrawal: input length mismatch"));
    //     delegationManager.queueWithdrawals(params);
    // }

    // function testQueueWithdrawalRevertsWithZeroAddressWithdrawer() external {
    //     IStrategy[] memory strategyArray = new IStrategy[](1);
    //     uint256[] memory shareAmounts = new uint256[](1);

    //     IDelegationManager.QueuedWithdrawalParams[] memory params = new IDelegationManager.QueuedWithdrawalParams[](1);
        
    //     params[0] = IDelegationManager.QueuedWithdrawalParams({
    //         strategies: strategyArray,
    //         shares: shareAmounts,
    //         withdrawer: address(0)
    //     });

    //     cheats.expectRevert(bytes("DelegationManager.queueWithdrawal: must provide valid withdrawal address"));
    //     delegationManager.queueWithdrawals(params);
    // }

    // function testQueueWithdrawal_ToSelf(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // )
    //     public
    //     returns (
    //         IDelegationManager.Withdrawal memory /* queuedWithdrawal */,
    //         IERC20[] memory /* tokensArray */,
    //         bytes32 /* withdrawalRoot */
    //     )
    // {
    //     _setUpWithdrawalTests();
    //     StrategyBase strategy = strategyMock;
    //     IERC20 token = strategy.underlyingToken();

    //     // filtering of fuzzed inputs
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);

    //     _tempStrategyStorage = strategy;

    //     _depositIntoStrategySuccessfully(strategy, /*staker*/ address(this), depositAmount);

    //     (
    //         IDelegationManager.Withdrawal memory withdrawal,
    //         IERC20[] memory tokensArray,
    //         bytes32 withdrawalRoot
    //     ) = _setUpWithdrawalStructSingleStrat(
    //             /*staker*/ address(this),
    //             /*withdrawer*/ address(this),
    //             token,
    //             _tempStrategyStorage,
    //             withdrawalAmount
    //         );

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(/*staker*/ address(this), _tempStrategyStorage);
    //     uint256 nonceBefore = delegationManager.cumulativeWithdrawalsQueued(/*staker*/ address(this));

    //     require(!delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRootPendingBefore is true!");

    //     {
    //         cheats.expectEmit(true, true, true, true, address(delegationManager));
    //         emit WithdrawalQueued(
    //             withdrawalRoot,
    //             withdrawal
    //         );

    //         IDelegationManager.QueuedWithdrawalParams[] memory params = new IDelegationManager.QueuedWithdrawalParams[](1);
            
    //         params[0] = IDelegationManager.QueuedWithdrawalParams({
    //             strategies: withdrawal.strategies,
    //             shares: withdrawal.shares,
    //             withdrawer: address(this)
    //         });
    //         delegationManager.queueWithdrawals(params);
    //     }

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(/*staker*/ address(this), _tempStrategyStorage);
    //     uint256 nonceAfter = delegationManager.cumulativeWithdrawalsQueued(/*staker*/ address(this));

    //     require(delegationManager.pendingWithdrawals(withdrawalRoot), "pendingWithdrawalsAfter is false!");
    //     require(sharesAfter == sharesBefore - withdrawalAmount, "sharesAfter != sharesBefore - withdrawalAmount");
    //     require(nonceAfter == nonceBefore + 1, "nonceAfter != nonceBefore + 1");

    //     return (withdrawal, tokensArray, withdrawalRoot);
    // }

    // function testQueueWithdrawal_ToSelf_TwoStrategies(
    //     uint256[2] memory depositAmounts,
    //     uint256[2] memory withdrawalAmounts
    // )
    //     public
    //     returns (
    //         IDelegationManager.Withdrawal memory /* withdrawal */,
    //         bytes32 /* withdrawalRoot */
    //     )
    // {
    //     _setUpWithdrawalTests();
    //     // filtering of fuzzed inputs
    //     cheats.assume(withdrawalAmounts[0] != 0 && withdrawalAmounts[0] < depositAmounts[0]);
    //     cheats.assume(withdrawalAmounts[1] != 0 && withdrawalAmounts[1] < depositAmounts[1]);
    //     address staker = address(this);

    //     IStrategy[] memory strategies = new IStrategy[](2);
    //     strategies[0] = IStrategy(strategyMock);
    //     strategies[1] = IStrategy(strategyMock2);

    //     IERC20[] memory tokens = new IERC20[](2);
    //     tokens[0] = strategyMock.underlyingToken();
    //     tokens[1] = strategyMock2.underlyingToken();

    //     uint256[] memory amounts = new uint256[](2);
    //     amounts[0] = withdrawalAmounts[0];
    //     amounts[1] = withdrawalAmounts[1];

    //     _depositIntoStrategySuccessfully(strategies[0], staker, depositAmounts[0]);
    //     _depositIntoStrategySuccessfully(strategies[1], staker, depositAmounts[1]);

    //     (
    //         IDelegationManager.Withdrawal memory withdrawal,
    //         bytes32 withdrawalRoot
    //     ) = _setUpWithdrawalStruct_MultipleStrategies(
    //             /* staker */ staker,
    //             /* withdrawer */ staker,
    //             strategies,
    //             amounts
    //         );

    //     uint256[] memory sharesBefore = new uint256[](2);
    //     sharesBefore[0] = strategyManager.stakerStrategyShares(staker, strategies[0]);
    //     sharesBefore[1] = strategyManager.stakerStrategyShares(staker, strategies[1]);
    //     uint256 nonceBefore = delegationManager.cumulativeWithdrawalsQueued(staker);

    //     require(!delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRootPendingBefore is true!");

    //     {
    //         cheats.expectEmit(true, true, true, true, address(delegationManager));
    //         emit WithdrawalQueued(
    //             withdrawalRoot,
    //             withdrawal
    //         );

    //         IDelegationManager.QueuedWithdrawalParams[] memory params = new IDelegationManager.QueuedWithdrawalParams[](1);
        
    //         params[0] = IDelegationManager.QueuedWithdrawalParams({
    //             strategies: withdrawal.strategies,
    //             shares: withdrawal.shares,
    //             withdrawer: staker
    //         });

    //         delegationManager.queueWithdrawals(
    //             params
    //         );
    //     }

    //     uint256[] memory sharesAfter = new uint256[](2);
    //     sharesAfter[0] = strategyManager.stakerStrategyShares(staker, strategies[0]);
    //     sharesAfter[1] = strategyManager.stakerStrategyShares(staker, strategies[1]);
    //     uint256 nonceAfter = delegationManager.cumulativeWithdrawalsQueued(staker);

    //     require(delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRootPendingAfter is false!");
    //     require(
    //         sharesAfter[0] == sharesBefore[0] - withdrawalAmounts[0],
    //         "Strat1: sharesAfter != sharesBefore - withdrawalAmount"
    //     );
    //     require(
    //         sharesAfter[1] == sharesBefore[1] - withdrawalAmounts[1],
    //         "Strat2: sharesAfter != sharesBefore - withdrawalAmount"
    //     );
    //     require(nonceAfter == nonceBefore + 1, "nonceAfter != nonceBefore + 1");

    //     return (withdrawal, withdrawalRoot);
    // }

    // function testQueueWithdrawalPartiallyWithdraw(uint128 amount) external {
    //     testQueueWithdrawal_ToSelf(uint256(amount) * 2, amount);
    //     require(!delegationManager.isDelegated(address(this)), "should still be delegated failed");
    // }

    // function testQueueWithdrawal_ToDifferentAddress(
    //     address withdrawer,
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external filterFuzzedAddressInputs(withdrawer) {
    //     _setUpWithdrawalTests();
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     address staker = address(this);

    //     _depositIntoStrategySuccessfully(strategyMock, staker, depositAmount);
    //     (
    //         IDelegationManager.Withdrawal memory withdrawal,
    //         ,
    //         bytes32 withdrawalRoot
    //     ) = _setUpWithdrawalStructSingleStrat(
    //             staker,
    //             withdrawer,
    //             /*token*/ strategyMock.underlyingToken(),
    //             strategyMock,
    //             withdrawalAmount
    //         );

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategyMock);
    //     uint256 nonceBefore = delegationManager.cumulativeWithdrawalsQueued(staker);

    //     require(!delegationManager.pendingWithdrawals(withdrawalRoot), "pendingWithdrawalsBefore is true!");

    //     cheats.expectEmit(true, true, true, true, address(delegationManager));
    //     emit WithdrawalQueued(
    //         withdrawalRoot,
    //         withdrawal
    //     );

    //     IDelegationManager.QueuedWithdrawalParams[] memory params = new IDelegationManager.QueuedWithdrawalParams[](1);
        
    //     params[0] = IDelegationManager.QueuedWithdrawalParams({
    //         strategies: withdrawal.strategies,
    //         shares: withdrawal.shares,
    //         withdrawer: withdrawer
    //     });

    //     delegationManager.queueWithdrawals(
    //         params
    //     );

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategyMock);
    //     uint256 nonceAfter = delegationManager.cumulativeWithdrawalsQueued(staker);

    //     require(delegationManager.pendingWithdrawals(withdrawalRoot), "pendingWithdrawalsAfter is false!");
    //     require(sharesAfter == sharesBefore - withdrawalAmount, "sharesAfter != sharesBefore - amount");
    //     require(nonceAfter == nonceBefore + 1, "nonceAfter != nonceBefore + 1");
    // }

    // function testCompleteQueuedWithdrawalRevertsWhenAttemptingReentrancy(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     // replace dummyStrat with Reenterer contract
    //     reenterer = new Reenterer();
    //     strategyMock = StrategyBase(address(reenterer));

    //     // whitelist the strategy for deposit
    //     cheats.startPrank(strategyManager.owner());
    //     IStrategy[] memory _strategy = new IStrategy[](1);
    //     _strategy[0] = strategyMock;
    //     strategyManager.addStrategiesToDepositWhitelist(_strategy);
    //     cheats.stopPrank();

    //     _tempStakerStorage = address(this);
    //     IStrategy strategy = strategyMock;

    //     reenterer.prepareReturnData(abi.encode(depositAmount));


    //     IStrategy[] memory strategyArray = new IStrategy[](1);
    //     IERC20[] memory tokensArray = new IERC20[](1);
    //     uint256[] memory shareAmounts = new uint256[](1);
    //     {
    //         strategyArray[0] = strategy;
    //         shareAmounts[0] = withdrawalAmount;
    //         tokensArray[0] = mockToken;
    //     }

    //     (
    //         IDelegationManager.Withdrawal memory withdrawal,
    //         /* tokensArray */,
    //         /* withdrawalRoot */
    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = false;

    //     address targetToUse = address(strategyManager);
    //     uint256 msgValueToUse = 0;
    //     bytes memory calldataToUse = abi.encodeWithSelector(
    //         DelegationManager.completeQueuedWithdrawal.selector,
    //         withdrawal,
    //         tokensArray,
    //         middlewareTimesIndex,
    //         receiveAsTokens
    //     );
    //     reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));
    //     delegationManager.completeQueuedWithdrawal(withdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);
    // }

    // function testCompleteQueuedWithdrawalRevertsWhenNotCallingFromWithdrawerAddress(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     _tempStakerStorage = address(this);

    //     (
    //         IDelegationManager.Withdrawal memory withdrawal,
    //         IERC20[] memory tokensArray,
    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy strategy = withdrawal.strategies[0];
    //     IERC20 token = tokensArray[0];

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceBefore = token.balanceOf(address(_tempStakerStorage));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = false;

    //     cheats.startPrank(address(123456));
    //     cheats.expectRevert(
    //         bytes(
    //             "DelegationManager.completeQueuedAction: only withdrawer can complete action"
    //         )
    //     );
    //     delegationManager.completeQueuedWithdrawal(withdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);
    //     cheats.stopPrank();

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceAfter = token.balanceOf(address(_tempStakerStorage));

    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    // }

    // function testCompleteQueuedWithdrawalRevertsWhenTryingToCompleteSameWithdrawal2X(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     _tempStakerStorage = address(this);

    //     (
    //         IDelegationManager.Withdrawal memory withdrawal,
    //         IERC20[] memory tokensArray,
    //         bytes32 withdrawalRoot
    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy strategy = withdrawal.strategies[0];
    //     IERC20 token = tokensArray[0];

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceBefore = token.balanceOf(address(_tempStakerStorage));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = false;

    //     cheats.expectEmit(true, true, true, true, address(delegationManager));
    //     emit WithdrawalCompleted(withdrawalRoot);
    //     delegationManager.completeQueuedWithdrawal(withdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceAfter = token.balanceOf(address(_tempStakerStorage));

    //     require(sharesAfter == sharesBefore + withdrawalAmount, "sharesAfter != sharesBefore + withdrawalAmount");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    //     sharesBefore = sharesAfter;
    //     balanceBefore = balanceAfter;

    //     cheats.expectRevert(
    //         bytes(
    //             "DelegationManager.completeQueuedAction: action is not in queue"
    //         )
    //     );
    //     delegationManager.completeQueuedWithdrawal(withdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     balanceAfter = token.balanceOf(address(_tempStakerStorage));

    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    // }

    // function testCompleteQueuedWithdrawalRevertsWhenWithdrawalDelayBlocksHasNotPassed(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     _tempStakerStorage = address(this);

    //     (
    //         IDelegationManager.Withdrawal memory withdrawal,
    //         IERC20[] memory tokensArray,
    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     uint256 valueToSet = 1;
    //     // set the `withdrawalDelayBlocks` variable
    //     cheats.startPrank(strategyManager.owner());
    //     uint256 previousValue = delegationManager.withdrawalDelayBlocks();
    //     cheats.expectEmit(true, true, true, true, address(delegationManager));
    //     emit WithdrawalDelayBlocksSet(previousValue, valueToSet);
    //     delegationManager.setWithdrawalDelayBlocks(valueToSet);
    //     cheats.stopPrank();
    //     require(
    //         delegationManager.withdrawalDelayBlocks() == valueToSet,
    //         "delegationManager.withdrawalDelayBlocks() != valueToSet"
    //     );

    //     cheats.expectRevert(
    //         bytes("DelegationManager.completeQueuedAction: withdrawalDelayBlocks period has not yet passed")
    //     );
    //     delegationManager.completeQueuedWithdrawal(withdrawal, tokensArray, /* middlewareTimesIndex */ 0, /* receiveAsTokens */ false);
    // }

    // function testCompleteQueuedWithdrawalRevertsWhenWithdrawalDoesNotExist() external {
    //     uint256 withdrawalAmount = 1e18;
    //     IStrategy strategy = strategyMock;
    //     IERC20 token = strategy.underlyingToken();

    //     (IDelegationManager.Withdrawal memory withdrawal, IERC20[] memory tokensArray, ) = _setUpWithdrawalStructSingleStrat(
    //             /*staker*/ address(this),
    //             /*withdrawer*/ address(this),
    //             token,
    //             strategy,
    //             withdrawalAmount
    //         );

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceBefore = token.balanceOf(address(_tempStakerStorage));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = false;

    //     cheats.expectRevert(bytes("DelegationManager.completeQueuedAction: action is not in queue"));
    //     delegationManager.completeQueuedWithdrawal(withdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceAfter = token.balanceOf(address(_tempStakerStorage));

    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    // }

    // function testCompleteQueuedWithdrawalRevertsWhenWithdrawalsPaused(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     _tempStakerStorage = address(this);

    //     (
    //         IDelegationManager.Withdrawal memory withdrawal,
    //         IERC20[] memory tokensArray,
    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy strategy = withdrawal.strategies[0];
    //     IERC20 token = tokensArray[0];

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceBefore = token.balanceOf(address(_tempStakerStorage));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = false;

    //     // pause withdrawals
    //     cheats.startPrank(pauser);
    //     delegationManager.pause(2 ** PAUSED_EXIT_WITHDRAWAL_QUEUE);
    //     cheats.stopPrank();

    //     cheats.expectRevert(bytes("Pausable: index is paused"));
    //     delegationManager.completeQueuedWithdrawal(withdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceAfter = token.balanceOf(address(_tempStakerStorage));

    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    // }

    // function testCompleteQueuedWithdrawalFailsWhenTokensInputLengthMismatch(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     _tempStakerStorage = address(this);

    //     (
    //         IDelegationManager.Withdrawal memory withdrawal,
    //         IERC20[] memory tokensArray,
    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy strategy = withdrawal.strategies[0];
    //     IERC20 token = tokensArray[0];

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceBefore = token.balanceOf(address(_tempStakerStorage));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = true;
    //     // mismatch tokens array by setting tokens array to empty array
    //     tokensArray = new IERC20[](0);

    //     cheats.expectRevert(bytes("DelegationManager.completeQueuedAction: input length mismatch"));
    //     delegationManager.completeQueuedWithdrawal(withdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceAfter = token.balanceOf(address(_tempStakerStorage));

    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    // }

    // function testCompleteQueuedWithdrawalWithNonzeroWithdrawalDelayBlocks(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount,
    //     uint16 valueToSet
    // ) external {
    //     // filter fuzzed inputs to allowed *and nonzero* amounts
    //     cheats.assume(valueToSet <= delegationManager.MAX_WITHDRAWAL_DELAY_BLOCKS() && valueToSet != 0);
    //     cheats.assume(depositAmount != 0 && withdrawalAmount != 0);
    //     cheats.assume(depositAmount >= withdrawalAmount);
    //     address staker = address(this);

    //     (
    //         IDelegationManager.Withdrawal memory withdrawal,
    //         IERC20[] memory tokensArray,
    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy strategy = withdrawal.strategies[0];
    //     IERC20 token = tokensArray[0];
    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = false;

    //     // set the `withdrawalDelayBlocks` variable
    //     cheats.startPrank(delegationManager.owner());
    //     uint256 previousValue = delegationManager.withdrawalDelayBlocks();
    //     cheats.expectEmit(true, true, true, true, address(delegationManager));
    //     emit WithdrawalDelayBlocksSet(previousValue, valueToSet);
    //     delegationManager.setWithdrawalDelayBlocks(valueToSet);
    //     cheats.stopPrank();
    //     require(
    //         delegationManager.withdrawalDelayBlocks() == valueToSet,
    //         "strategyManager.withdrawalDelayBlocks() != valueToSet"
    //     );

    //     cheats.expectRevert(
    //         bytes("DelegationManager.completeQueuedAction: withdrawalDelayBlocks period has not yet passed")
    //     );
    //     delegationManager.completeQueuedWithdrawal(withdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceBefore = token.balanceOf(address(staker));

    //     // roll block number forward to one block before the withdrawal should be completeable and attempt again
    //     uint256 originalBlockNumber = block.number;
    //     cheats.roll(originalBlockNumber + valueToSet - 1);
    //     cheats.expectRevert(
    //         bytes("DelegationManager.completeQueuedAction: withdrawalDelayBlocks period has not yet passed")
    //     );
    //     delegationManager.completeQueuedWithdrawal(withdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     // roll block number forward to the block at which the withdrawal should be completeable, and complete it
    //     cheats.roll(originalBlockNumber + valueToSet);
    //     delegationManager.completeQueuedWithdrawal(withdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceAfter = token.balanceOf(address(staker));

    //     require(sharesAfter == sharesBefore + withdrawalAmount, "sharesAfter != sharesBefore + withdrawalAmount");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    // }


    // function testCompleteQueuedWithdrawal_ReceiveAsTokensMarkedFalse(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     address staker = address(this);

    //     // testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     (
    //         IDelegationManager.Withdrawal memory withdrawal,
    //         IERC20[] memory tokensArray,
    //         bytes32 withdrawalRoot
    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy strategy = withdrawal.strategies[0];
    //     IERC20 token = tokensArray[0];

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceBefore = token.balanceOf(address(staker));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = false;
    //     cheats.expectEmit(true, true, true, true, address(delegationManager));
    //     emit WithdrawalCompleted(withdrawalRoot);
    //     delegationManager.completeQueuedWithdrawal(withdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(address(this), strategy);
    //     uint256 balanceAfter = token.balanceOf(address(staker));

    //     require(sharesAfter == sharesBefore + withdrawalAmount, "sharesAfter != sharesBefore + withdrawalAmount");
    //     require(balanceAfter == balanceBefore, "balanceAfter != balanceBefore");
    // }

    // function testCompleteQueuedWithdrawal_ReceiveAsTokensMarkedTrue(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     cheats.assume(withdrawalAmount != 0 && withdrawalAmount <= depositAmount);
    //     address staker = address(this);

    //     (
    //         IDelegationManager.Withdrawal memory withdrawal,
    //         IERC20[] memory tokensArray,
    //         bytes32 withdrawalRoot
    //     ) = testQueueWithdrawal_ToSelf(depositAmount, withdrawalAmount);

    //     IStrategy strategy = withdrawal.strategies[0];
    //     IERC20 token = tokensArray[0];

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
    //     uint256 balanceBefore = token.balanceOf(address(staker));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = true;
    //     cheats.expectEmit(true, true, true, true, address(delegationManager));
    //     emit WithdrawalCompleted(withdrawalRoot);
    //     delegationManager.completeQueuedWithdrawal(withdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
    //     uint256 balanceAfter = token.balanceOf(address(staker));

    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(balanceAfter == balanceBefore + withdrawalAmount, "balanceAfter != balanceBefore + withdrawalAmount");
    //     if (depositAmount == withdrawalAmount) {
    //         // Since receiving tokens instead of shares, if withdrawal amount is entire deposit, then strategy will be removed
    //         // with sharesAfter being 0
    //         require(
    //             !_isDepositedStrategy(staker, strategy),
    //             "Strategy still part of staker's deposited strategies"
    //         );
    //         require(sharesAfter == 0, "staker shares is not 0");
    //     }
    // }

    // function testCompleteQueuedWithdrawalFullyWithdraw(uint256 amount) external {
    //     address staker = address(this);
    //     (
    //         IDelegationManager.Withdrawal memory withdrawal,
    //         IERC20[] memory tokensArray,
    //         bytes32 withdrawalRoot
    //     ) = testQueueWithdrawal_ToSelf(amount, amount);

    //     IStrategy strategy = withdrawal.strategies[0];
    //     IERC20 token = tokensArray[0];

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
    //     uint256 balanceBefore = token.balanceOf(address(staker));

    //     uint256 middlewareTimesIndex = 0;
    //     bool receiveAsTokens = true;
    //     cheats.expectEmit(true, true, true, true, address(delegationManager));
    //     emit WithdrawalCompleted(withdrawalRoot);
    //     delegationManager.completeQueuedWithdrawal(withdrawal, tokensArray, middlewareTimesIndex, receiveAsTokens);

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
    //     uint256 balanceAfter = token.balanceOf(address(staker));

    //     require(sharesAfter == sharesBefore, "sharesAfter != sharesBefore");
    //     require(balanceAfter == balanceBefore + amount, "balanceAfter != balanceBefore + withdrawalAmount");
    //     require(
    //         !_isDepositedStrategy(staker, strategy),
    //         "Strategy still part of staker's deposited strategies"
    //     );
    //     require(sharesAfter == 0, "staker shares is not 0");
    // }

    // function test_removeSharesRevertsWhenShareAmountIsZero(uint256 depositAmount) external {
    //     _setUpWithdrawalTests();
    //     address staker = address(this);
    //     uint256 withdrawalAmount = 0;

    //     _depositIntoStrategySuccessfully(strategyMock, staker, depositAmount);

    //     (IDelegationManager.Withdrawal memory withdrawal, , ) = _setUpWithdrawalStructSingleStrat(
    //             /*staker*/ address(this),
    //             /*withdrawer*/ address(this),
    //             mockToken,
    //             strategyMock,
    //             withdrawalAmount
    //         );

    //     IDelegationManager.QueuedWithdrawalParams[] memory params = new IDelegationManager.QueuedWithdrawalParams[](1);
        
    //     params[0] = IDelegationManager.QueuedWithdrawalParams({
    //         strategies: withdrawal.strategies,
    //         shares: withdrawal.shares,
    //         withdrawer: staker
    //     });

    //     cheats.expectRevert(bytes("StrategyManager._removeShares: shareAmount should not be zero!"));
    //     delegationManager.queueWithdrawals(
    //         params
    //     );
    // }

    // function test_removeSharesRevertsWhenShareAmountIsTooLarge(
    //     uint256 depositAmount,
    //     uint256 withdrawalAmount
    // ) external {
    //     _setUpWithdrawalTests();
    //     cheats.assume(depositAmount > 0 && withdrawalAmount > depositAmount);
    //     address staker = address(this);

    //     _depositIntoStrategySuccessfully(strategyMock, staker, depositAmount);

    //     (IDelegationManager.Withdrawal memory withdrawal, ,  ) = _setUpWithdrawalStructSingleStrat(
    //             /*staker*/ address(this),
    //             /*withdrawer*/ address(this),
    //             mockToken,
    //             strategyMock,
    //             withdrawalAmount
    //         );

    //     IDelegationManager.QueuedWithdrawalParams[] memory params = new IDelegationManager.QueuedWithdrawalParams[](1);
        
    //     params[0] = IDelegationManager.QueuedWithdrawalParams({
    //         strategies: withdrawal.strategies,
    //         shares: withdrawal.shares,
    //         withdrawer: address(this)
    //     });

    //     cheats.expectRevert(bytes("StrategyManager._removeShares: shareAmount too high"));
    //     delegationManager.queueWithdrawals(
    //         params
    //     );
    // }

    // /**
    //  * Testing queueWithdrawal of 3 strategies, fuzzing the deposit and withdraw amounts. if the withdrawal amounts == deposit amounts
    //  * then the strategy should be removed from the staker StrategyList
    //  */
    // function test_removeStrategyFromStakerStrategyList(uint256[3] memory depositAmounts, uint256[3] memory withdrawalAmounts) external {
    //     _setUpWithdrawalTests();
    //     // filtering of fuzzed inputs
    //     cheats.assume(withdrawalAmounts[0] > 0 && withdrawalAmounts[0] <= depositAmounts[0]);
    //     cheats.assume(withdrawalAmounts[1] > 0 && withdrawalAmounts[1] <= depositAmounts[1]);
    //     cheats.assume(withdrawalAmounts[2] > 0 && withdrawalAmounts[2] <= depositAmounts[2]);
    //     address staker = address(this);

    //     // Setup input params
    //     IStrategy[] memory strategies = new IStrategy[](3);
    //     strategies[0] = strategyMock;
    //     strategies[1] = strategyMock2;
    //     strategies[2] = strategyMock3;
    //     uint256[] memory amounts = new uint256[](3);
    //     amounts[0] = withdrawalAmounts[0];
    //     amounts[1] = withdrawalAmounts[1];
    //     amounts[2] = withdrawalAmounts[2];

    //     _depositIntoStrategySuccessfully(strategies[0], staker, depositAmounts[0]);
    //     _depositIntoStrategySuccessfully(strategies[1], staker, depositAmounts[1]);
    //     _depositIntoStrategySuccessfully(strategies[2], staker, depositAmounts[2]);

    //     ( ,bytes32 withdrawalRoot) = _setUpWithdrawalStruct_MultipleStrategies(
    //             /* staker */ staker,
    //             /* withdrawer */ staker,
    //             strategies,
    //             amounts
    //         );
    //     require(!delegationManager.pendingWithdrawals(withdrawalRoot), "withdrawalRootPendingBefore is true!");
    //     delegationManager.cumulativeWithdrawalsQueued(staker);
    //     uint256[] memory sharesBefore = new uint256[](3);
    //     sharesBefore[0] = strategyManager.stakerStrategyShares(staker, strategies[0]);
    //     sharesBefore[1] = strategyManager.stakerStrategyShares(staker, strategies[1]);
    //     sharesBefore[2] = strategyManager.stakerStrategyShares(staker, strategies[2]);

    //     IDelegationManager.QueuedWithdrawalParams[] memory params = new IDelegationManager.QueuedWithdrawalParams[](1);
        
    //     params[0] = IDelegationManager.QueuedWithdrawalParams({
    //         strategies: strategies,
    //         shares: amounts,
    //         withdrawer: address(this)
    //     });

    //     delegationManager.queueWithdrawals(
    //         params
    //     );

    //     uint256[] memory sharesAfter = new uint256[](3);
    //     sharesAfter[0] = strategyManager.stakerStrategyShares(staker, strategies[0]);
    //     sharesAfter[1] = strategyManager.stakerStrategyShares(staker, strategies[1]);
    //     sharesAfter[2] = strategyManager.stakerStrategyShares(staker, strategies[2]);
    //     require(sharesBefore[0] == sharesAfter[0] + withdrawalAmounts[0], "Strat1: sharesBefore != sharesAfter + withdrawalAmount");
    //     if (depositAmounts[0] == withdrawalAmounts[0]) {
    //         require(!_isDepositedStrategy(staker, strategies[0]), "Strategy still part of staker's deposited strategies");
    //     }
    //     require(sharesBefore[1] == sharesAfter[1] + withdrawalAmounts[1], "Strat2: sharesBefore != sharesAfter + withdrawalAmount");
    //     if (depositAmounts[1] == withdrawalAmounts[1]) {
    //         require(!_isDepositedStrategy(staker, strategies[1]), "Strategy still part of staker's deposited strategies");
    //     }
    //     require(sharesBefore[2] == sharesAfter[2] + withdrawalAmounts[2], "Strat3: sharesBefore != sharesAfter + withdrawalAmount");
    //     if (depositAmounts[2] == withdrawalAmounts[2]) {
    //         require(!_isDepositedStrategy(staker, strategies[2]), "Strategy still part of staker's deposited strategies");
    //     }
    // }

    // // ensures that when the staker and withdrawer are different and a withdrawal is completed as shares (i.e. not as tokens)
    // // that the shares get added back to the right operator
    // function test_completingWithdrawalAsSharesAddsSharesToCorrectOperator() external {
    //     address staker = address(this);
    //     address withdrawer = address(1000);
    //     address operator_for_staker = address(1001);
    //     address operator_for_withdrawer = address(1002);

    //     // register operators
    //     bytes32 salt;
    //     IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
    //         earningsReceiver: operator_for_staker,
    //         delegationApprover: address(0),
    //         stakerOptOutWindowBlocks: 0
    //     });
    //     testRegisterAsOperator(operator_for_staker, operatorDetails, emptyStringForMetadataURI);
    //     testRegisterAsOperator(operator_for_withdrawer, operatorDetails, emptyStringForMetadataURI);

    //     // delegate from the `staker` and withdrawer to the operators
    //     ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;
    //     cheats.startPrank(staker);
    //     delegationManager.delegateTo(operator_for_staker, approverSignatureAndExpiry, salt);        
    //     cheats.stopPrank();
    //     cheats.startPrank(withdrawer);
    //     delegationManager.delegateTo(operator_for_withdrawer, approverSignatureAndExpiry, salt);        
    //     cheats.stopPrank();

    //     // Setup input params
    //     IStrategy[] memory strategies = new IStrategy[](3);
    //     strategies[0] = strategyMock;
    //     strategies[1] = delegationManager.beaconChainETHStrategy();
    //     strategies[2] = strategyMock3;
    //     uint256[] memory amounts = new uint256[](3);
    //     amounts[0] = 1e18;
    //     amounts[1] = 2e18;
    //     amounts[2] = 3e18;

    //     (IDelegationManager.Withdrawal memory withdrawal, ) = _setUpWithdrawalStruct_MultipleStrategies({
    //             staker: staker,
    //             withdrawer: withdrawer,
    //             strategyArray: strategies,
    //             shareAmounts: amounts
    //         });

    //     // give both the operators a bunch of delegated shares, so we can decrement them when queuing the withdrawal
    //     cheats.startPrank(address(delegationManager.strategyManager()));
    //     for (uint256 i = 0; i < strategies.length; ++i) {
    //         delegationManager.increaseDelegatedShares(staker, strategies[i], amounts[i]);
    //         delegationManager.increaseDelegatedShares(withdrawer, strategies[i], amounts[i]);
    //     }
    //     cheats.stopPrank();

    //     IDelegationManager.QueuedWithdrawalParams[] memory params = new IDelegationManager.QueuedWithdrawalParams[](1);
        
    //     params[0] = IDelegationManager.QueuedWithdrawalParams({
    //         strategies: strategies,
    //         shares: amounts,
    //         withdrawer: withdrawer
    //     });

    //     // queue the withdrawal
    //     cheats.startPrank(staker);
    //     delegationManager.queueWithdrawals(params);        
    //     cheats.stopPrank();

    //     for (uint256 i = 0; i < strategies.length; ++i) {
    //         require(delegationManager.operatorShares(operator_for_staker, strategies[i]) == 0,
    //             "staker operator shares incorrect after queueing");
    //         require(delegationManager.operatorShares(operator_for_withdrawer, strategies[i]) == amounts[i],
    //             "withdrawer operator shares incorrect after queuing");
    //     }

    //     // complete the withdrawal
    //     cheats.startPrank(withdrawer);
    //     IERC20[] memory tokens;
    //     delegationManager.completeQueuedWithdrawal(
    //         withdrawal,
    //         tokens,
    //         0 /*middlewareTimesIndex*/,
    //         false /*receiveAsTokens*/
    //     );
    //     cheats.stopPrank();

    //     for (uint256 i = 0; i < strategies.length; ++i) {
    //         if (strategies[i] != delegationManager.beaconChainETHStrategy()) {
    //             require(delegationManager.operatorShares(operator_for_staker, strategies[i]) == 0,
    //                 "staker operator shares incorrect after completing withdrawal");
    //             require(delegationManager.operatorShares(operator_for_withdrawer, strategies[i]) == 2 * amounts[i],
    //                 "withdrawer operator shares incorrect after completing withdrawal");
    //         } else {
    //             require(delegationManager.operatorShares(operator_for_staker, strategies[i]) == amounts[i],
    //                 "staker operator beaconChainETHStrategy shares incorrect after completing withdrawal");
    //             require(delegationManager.operatorShares(operator_for_withdrawer, strategies[i]) == amounts[i],
    //                 "withdrawer operator beaconChainETHStrategy shares incorrect after completing withdrawal");
    //         }
    //     }
    // }

    // /**
    //  * Setup DelegationManager and StrategyManager contracts for testing instead of using StrategyManagerMock
    //  * since we need to test the actual contracts together for the withdrawal queueing tests
    //  */
    // function _setUpWithdrawalTests() internal {
    //     delegationManagerImplementation = new DelegationManager(strategyManager, slasherMock, eigenPodManagerMock);
    //     cheats.startPrank(eigenLayerProxyAdmin.owner());
    //     eigenLayerProxyAdmin.upgrade(TransparentUpgradeableProxy(payable(address(delegationManager))), address(delegationManagerImplementation));
    //     cheats.stopPrank();


    //     strategyImplementation = new StrategyBase(strategyManager);
    //     mockToken = new ERC20Mock();
    //     strategyMock = StrategyBase(
    //         address(
    //             new TransparentUpgradeableProxy(
    //                 address(strategyImplementation),
    //                 address(eigenLayerProxyAdmin),
    //                 abi.encodeWithSelector(StrategyBase.initialize.selector, mockToken, pauserRegistry)
    //             )
    //         )
    //     );
    //     strategyMock2 = StrategyBase(
    //         address(
    //             new TransparentUpgradeableProxy(
    //                 address(strategyImplementation),
    //                 address(eigenLayerProxyAdmin),
    //                 abi.encodeWithSelector(StrategyBase.initialize.selector, mockToken, pauserRegistry)
    //             )
    //         )
    //     );
    //     strategyMock3 = StrategyBase(
    //         address(
    //             new TransparentUpgradeableProxy(
    //                 address(strategyImplementation),
    //                 address(eigenLayerProxyAdmin),
    //                 abi.encodeWithSelector(StrategyBase.initialize.selector, mockToken, pauserRegistry)
    //             )
    //         )
    //     );

    //     // whitelist the strategy for deposit
    //     cheats.startPrank(strategyManager.owner());
    //     IStrategy[] memory _strategies = new IStrategy[](3);
    //     _strategies[0] = strategyMock;
    //     _strategies[1] = strategyMock2;
    //     _strategies[2] = strategyMock3;
    //     strategyManager.addStrategiesToDepositWhitelist(_strategies);
    //     cheats.stopPrank();

    //     require(delegationManager.strategyManager() == strategyManager,
    //         "constructor / initializer incorrect, strategyManager set wrong");
    // }

    // function _depositIntoStrategySuccessfully(
    //     IStrategy strategy,
    //     address staker,
    //     uint256 amount
    // ) internal {
    //     IERC20 token = strategy.underlyingToken();
    //     // IStrategy strategy = strategyMock;

    //     // filter out zero case since it will revert with "StrategyManager._addShares: shares should not be zero!"
    //     cheats.assume(amount != 0);
    //     // filter out zero address because the mock ERC20 we are using will revert on using it
    //     cheats.assume(staker != address(0));
    //     // filter out the strategy itself from fuzzed inputs
    //     cheats.assume(staker != address(strategy));
    //     // sanity check / filter
    //     cheats.assume(amount <= token.balanceOf(address(this)));
    //     cheats.assume(amount >= 1);

    //     uint256 sharesBefore = strategyManager.stakerStrategyShares(staker, strategy);
    //     uint256 stakerStrategyListLengthBefore = strategyManager.stakerStrategyListLength(staker);

    //     // needed for expecting an event with the right parameters
    //     uint256 expectedShares = strategy.underlyingToShares(amount);

    //     cheats.startPrank(staker);
    //     cheats.expectEmit(true, true, true, true, address(strategyManager));
    //     emit Deposit(staker, token, strategy, expectedShares);
    //     uint256 shares = strategyManager.depositIntoStrategy(strategy, token, amount);
    //     cheats.stopPrank();

    //     uint256 sharesAfter = strategyManager.stakerStrategyShares(staker, strategy);
    //     uint256 stakerStrategyListLengthAfter = strategyManager.stakerStrategyListLength(staker);

    //     require(sharesAfter == sharesBefore + shares, "sharesAfter != sharesBefore + shares");
    //     if (sharesBefore == 0) {
    //         require(
    //             stakerStrategyListLengthAfter == stakerStrategyListLengthBefore + 1,
    //             "stakerStrategyListLengthAfter != stakerStrategyListLengthBefore + 1"
    //         );
    //         require(
    //             strategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1) == strategy,
    //             "strategyManager.stakerStrategyList(staker, stakerStrategyListLengthAfter - 1) != strategy"
    //         );
    //     }
    // }

    // /**
    //  * @notice internal function to help check if a strategy is part of list of deposited strategies for a staker
    //  * Used to check if removed correctly after withdrawing all shares for a given strategy
    //  */
    // function _isDepositedStrategy(address staker, IStrategy strategy) internal view returns (bool) {
    //     uint256 stakerStrategyListLength = strategyManager.stakerStrategyListLength(staker);
    //     for (uint256 i = 0; i < stakerStrategyListLength; ++i) {
    //         if (strategyManager.stakerStrategyList(staker, i) == strategy) {
    //             return true;
    //         }
    //     }
    //     return false;
    // }

    function _testRegisterAdditionalOperator(uint256 index) internal {
        address sender = getOperatorAddress(index);

        //register as both ETH and EIGEN operator
        uint256 wethToDeposit = 1e18;
        uint256 eigenToDeposit = 1e10;
        _testDepositWeth(sender, wethToDeposit);
        _testDepositEigen(sender, eigenToDeposit);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: sender,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        _testRegisterAsOperator(sender, operatorDetails);
        cheats.startPrank(sender);

        cheats.stopPrank();
    }

    // registers the operator if they are not already registered, and deposits "WETH" + "EIGEN" on behalf of the staker.
    function _registerOperatorAndDepositFromStaker(
        address operator,
        address staker,
        uint96 ethAmount,
        uint96 eigenAmount
    ) internal {
        cheats.assume(staker != operator);

        // if first deposit amount to base strategy is too small, it will revert. ignore that case here.
        cheats.assume(ethAmount >= 1 && ethAmount <= 1e18);
        cheats.assume(eigenAmount >= 1 && eigenAmount <= 1e18);

        if (!delegation.isOperator(operator)) {
            IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
                earningsReceiver: operator,
                delegationApprover: address(0),
                stakerOptOutWindowBlocks: 0
            });
            _testRegisterAsOperator(operator, operatorDetails);
        }

        //making additional deposits to the strategies
        assertTrue(!delegation.isDelegated(staker) == true, "testDelegation: staker is not delegate");
        _testDepositWeth(staker, ethAmount);
        _testDepositEigen(staker, eigenAmount);
    }
}
