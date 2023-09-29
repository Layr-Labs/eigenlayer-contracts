// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import {ERC1271WalletMock, ERC1271MaliciousMock} from "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {IDelegationManager} from "src/contracts/interfaces/IDelegationManager.sol";
import {IStrategy} from "src/contracts/interfaces/IStrategy.sol";

import {EigenLayerTestHelper, EigenLayerDeployer} from "src/test/EigenLayerTestHelper.t.sol";
import {MiddlewareRegistryMock} from "src/test/mocks/MiddlewareRegistryMock.sol";
import {MiddlewareVoteWeigherMock, VoteWeigherBaseStorage} from "src/test/mocks/MiddlewareVoteWeigherMock.sol";
import {ServiceManagerMock} from "src/test/mocks/ServiceManagerMock.sol";

contract DelegationTests is EigenLayerTestHelper {
    uint256 internal PRIVATE_KEY = 420;
    uint32 internal serveUntil = 100;

    ServiceManagerMock internal serviceManager;
    MiddlewareVoteWeigherMock internal voteWeigher;
    MiddlewareVoteWeigherMock internal voteWeigherImplementation;

    bytes internal constant OPERATOR_NOT_REGISTERED_ERROR =
        bytes("DelegationManager._delegate: operator is not registered in EigenLayer");
    bytes internal constant SIGNATURE_EXPIRED_ERROR =
        bytes("DelegationManager.delegateToBySignature: staker signature expired");
    bytes internal constant SIGNATURE_INVALID_ERROR =
        bytes("EIP1271SignatureUtils.checkSignature_EIP1271: ERC1271 signature verification failed");
    bytes internal constant OPERATOR_ALREADY_REGISTERED_ERROR =
        bytes("DelegationManager.registerAsOperator: operator has already registered");
    bytes internal constant ALREADY_INITIALIZED_ERROR = bytes("Initializable: contract is already initialized");
    bytes internal constant EARNINGS_RECEIVER_ADDRESS_ZERO_ERROR =
        bytes("DelegationManager._setOperatorDetails: cannot set `earningsReceiver` to zero address");
    bytes internal constant OPERATOR_UNDELEGATION_ERROR =
        bytes("DelegationManager.undelegate: operators cannot undelegate from themselves");

    modifier fuzzedAmounts(uint256 ethAmount, uint256 eigenAmount) {
        vm.assume(ethAmount >= 0 && ethAmount <= 1e18);
        vm.assume(eigenAmount >= 0 && eigenAmount <= 1e18);
        _;
    }

    function setUp() public virtual override {
        EigenLayerDeployer.setUp();

        initializeMiddlewares();
    }

    function initializeMiddlewares() public {
        serviceManager = new ServiceManagerMock(slasher);
        voteWeigher = MiddlewareVoteWeigherMock(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        voteWeigherImplementation = new MiddlewareVoteWeigherMock(delegation, strategyManager, serviceManager);
        {
            uint96 multiplier = 1e18;
            uint8 _NUMBER_OF_QUORUMS = 2;
            uint256[] memory _quorumBips = new uint256[](_NUMBER_OF_QUORUMS);
            // split 60% ETH quorum, 40% EIGEN quorum
            _quorumBips[0] = 6000;
            _quorumBips[1] = 4000;
            VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] memory ethStratsAndMultipliers =
                new VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[](1);
            ethStratsAndMultipliers[0].strategy = wethStrat;
            ethStratsAndMultipliers[0].multiplier = multiplier;
            VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] memory eigenStratsAndMultipliers =
                new VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[](1);
            eigenStratsAndMultipliers[0].strategy = eigenStrat;
            eigenStratsAndMultipliers[0].multiplier = multiplier;

            vm.startPrank(eigenLayerProxyAdmin.owner());
            eigenLayerProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(voteWeigher))),
                address(voteWeigherImplementation),
                abi.encodeWithSelector(
                    MiddlewareVoteWeigherMock.initialize.selector,
                    _quorumBips,
                    ethStratsAndMultipliers,
                    eigenStratsAndMultipliers
                )
            );
            vm.stopPrank();
        }
    }

    /// @notice testing if an operator can register to themselves.
    function test_SelfOperatorRegister() public {
        _registerAdditionalOperator(0, serveUntil);
    }

    /// @notice testing if an operator can delegate to themselves.
    /// @param sender is the address of the operator.
    function testFuzz_SelfOperatorDelegate(address sender) public {
        vm.assume(sender != address(0));
        vm.assume(sender != address(eigenLayerProxyAdmin));
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: sender,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        _testRegisterAsOperator(sender, operatorDetails);
    }

    function test_TwoSelfOperatorsRegister() public {
        _registerAdditionalOperator(0, serveUntil);
        _registerAdditionalOperator(1, serveUntil);
    }

    /// @notice registers a fixed address as a delegate, delegates to it from a second address,
    ///         and checks that the delegate's voteWeights increase properly
    /// @param operator is the operator being delegated to.
    /// @param staker is the staker delegating stake to the operator.
    function testFuzz_Delegation(address operator, address staker, uint96 ethAmount, uint96 eigenAmount)
        public
        fuzzedAddress(operator)
        fuzzedAddress(staker)
        fuzzedAmounts(ethAmount, eigenAmount)
    {
        vm.assume(staker != operator);
        // base strategy will revert if these amounts are too small on first deposit
        vm.assume(ethAmount >= 1);
        vm.assume(eigenAmount >= 1);

        _testDelegation(operator, staker, ethAmount, eigenAmount, voteWeigher);
    }

    /// @notice tests that a when an operator is delegated to, that delegation is properly accounted for.
    function testFuzz_DelegationReceived(address _operator, address staker, uint64 ethAmount, uint64 eigenAmount)
        public
        fuzzedAddress(_operator)
        fuzzedAddress(staker)
        fuzzedAmounts(ethAmount, eigenAmount)
    {
        vm.assume(staker != _operator);
        vm.assume(ethAmount >= 1);
        vm.assume(eigenAmount >= 1);

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

        uint256[3] memory amountsBefore;
        amountsBefore[0] = voteWeigher.weightOfOperator(operator, 0);
        amountsBefore[1] = voteWeigher.weightOfOperator(operator, 1);
        amountsBefore[2] = delegation.operatorShares(operator, wethStrat);

        //making additional deposits to the  strategies
        assertTrue(!delegation.isDelegated(staker), "staker is not delegate");
        _testDepositWeth(staker, ethAmount);
        _testDepositEigen(staker, eigenAmount);
        _testDelegateToOperator(staker, operator);
        assertTrue(delegation.isDelegated(staker), "staker is not delegate");

        (IStrategy[] memory updatedStrategies, uint256[] memory updatedShares) = strategyManager.getDeposits(staker);

        {
            uint256 stakerEthWeight = strategyManager.stakerStrategyShares(staker, updatedStrategies[0]);
            uint256 stakerEigenWeight = strategyManager.stakerStrategyShares(staker, updatedStrategies[1]);

            uint256 operatorEthWeightAfter = voteWeigher.weightOfOperator(operator, 0);
            uint256 operatorEigenWeightAfter = voteWeigher.weightOfOperator(operator, 1);

            assertTrue(
                operatorEthWeightAfter - amountsBefore[0] == stakerEthWeight,
                "operatorEthWeight did not increment by the right amount"
            );
            assertTrue(
                operatorEigenWeightAfter - amountsBefore[1] == stakerEigenWeight,
                "Eigen weights did not increment by the right amount"
            );
        }
        {
            IStrategy _strat = wethStrat;
            // IStrategy _strat = strategyManager.stakerStrats(staker, 0);
            assertTrue(address(_strat) != address(0), "stakerStrats not updated correctly");

            assertTrue(
                delegation.operatorShares(operator, _strat) - updatedShares[0] == amountsBefore[2],
                "ETH operatorShares not updated correctly"
            );

            vm.startPrank(address(strategyManager));

            IDelegationManager.OperatorDetails memory expectedOperatorDetails = delegation.operatorDetails(operator);
            assertTrue(
                keccak256(abi.encode(expectedOperatorDetails)) == keccak256(abi.encode(operatorDetails)),
                "failed to set correct operator details"
            );
        }
    }

    /// @notice tests that a when an operator is undelegated from, that the staker is properly classified as undelegated.
    function testFuzz_Undelegation(address operator, address staker, uint96 ethAmount, uint96 eigenAmount)
        public
        fuzzedAddress(operator)
        fuzzedAddress(staker)
        fuzzedAmounts(ethAmount, eigenAmount)
    {
        vm.assume(staker != operator);
        // base strategy will revert if these amounts are too small on first deposit
        vm.assume(ethAmount >= 1);
        vm.assume(eigenAmount >= 1);

        _testDelegation(operator, staker, ethAmount, eigenAmount, voteWeigher);
        vm.startPrank(address(strategyManager));
        delegation.undelegate(staker);
        vm.stopPrank();

        require(delegation.delegatedTo(staker) == address(0), "undelegation unsuccessful");
    }

    /// @notice tests delegation from a staker to operator via ECDSA signature.
    function testFuzz_DelegateToBySignature(address operator, uint96 ethAmount, uint96 eigenAmount, uint256 expiry)
        public
        fuzzedAddress(operator)
    {
        address staker = vm.addr(PRIVATE_KEY);
        _registerOperatorAndDepositFromStaker(operator, staker, ethAmount, eigenAmount);

        uint256 nonceBefore = delegation.stakerNonce(staker);

        bytes32 structHash =
            keccak256(abi.encode(delegation.STAKER_DELEGATION_TYPEHASH(), staker, operator, nonceBefore, expiry));
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", delegation.domainSeparator(), structHash));

        bytes memory signature;
        {
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(PRIVATE_KEY, digestHash);
            signature = abi.encodePacked(r, s, v);
        }

        if (expiry < block.timestamp) {
            vm.expectRevert(SIGNATURE_EXPIRED_ERROR);
        }
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry =
            IDelegationManager.SignatureWithExpiry({signature: signature, expiry: expiry});
        delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, bytes32(0));
        if (expiry >= block.timestamp) {
            assertTrue(delegation.isDelegated(staker), "staker is not delegate");
            assertTrue(nonceBefore + 1 == delegation.stakerNonce(staker), "nonce not incremented correctly");
            assertTrue(delegation.delegatedTo(staker) == operator, "staker delegated to wrong operator");
        }
    }

    /// @notice tries delegating using a signature and an EIP 1271 compliant wallet
    function testFuzz_DelegateToBySignature_WithContractWallet_Successfully(
        address operator,
        uint96 ethAmount,
        uint96 eigenAmount
    ) public fuzzedAddress(operator) {
        address staker = vm.addr(PRIVATE_KEY);

        // deploy ERC1271WalletMock for staker to use
        vm.prank(staker);
        ERC1271WalletMock wallet = new ERC1271WalletMock(staker);

        staker = address(wallet);

        _registerOperatorAndDepositFromStaker(operator, staker, ethAmount, eigenAmount);

        uint256 nonceBefore = delegation.stakerNonce(staker);

        bytes32 structHash = keccak256(
            abi.encode(delegation.STAKER_DELEGATION_TYPEHASH(), staker, operator, nonceBefore, type(uint256).max)
        );
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", delegation.domainSeparator(), structHash));

        bytes memory signature;
        {
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(PRIVATE_KEY, digestHash);
            signature = abi.encodePacked(r, s, v);
        }

        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry =
            IDelegationManager.SignatureWithExpiry({signature: signature, expiry: type(uint256).max});
        delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, bytes32(0));
        assertTrue(delegation.isDelegated(staker), "staker is not delegate");
        assertTrue(nonceBefore + 1 == delegation.stakerNonce(staker), "nonce not incremented correctly");
        assertTrue(delegation.delegatedTo(staker) == operator, "staker delegated to wrong operator");
    }

    ///  @notice tries delegating using a signature and an EIP 1271 compliant wallet, *but* providing a bad signature
    function testFuzz_DelegateToBySignature_WithContractWallet_BadSignature(
        address operator,
        uint96 ethAmount,
        uint96 eigenAmount
    ) public fuzzedAddress(operator) {
        address staker = vm.addr(PRIVATE_KEY);

        // deploy ERC1271WalletMock for staker to use
        vm.startPrank(staker);
        ERC1271WalletMock wallet = new ERC1271WalletMock(staker);
        vm.stopPrank();
        staker = address(wallet);

        _registerOperatorAndDepositFromStaker(operator, staker, ethAmount, eigenAmount);

        uint256 nonceBefore = delegation.stakerNonce(staker);

        bytes32 structHash = keccak256(
            abi.encode(delegation.STAKER_DELEGATION_TYPEHASH(), staker, operator, nonceBefore, type(uint256).max)
        );
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", delegation.domainSeparator(), structHash));

        bytes memory signature;
        {
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(PRIVATE_KEY, digestHash);
            // mess up the signature by flipping v's parity
            v = (v == 27 ? 28 : 27);
            signature = abi.encodePacked(r, s, v);
        }

        vm.expectRevert(SIGNATURE_INVALID_ERROR);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry =
            IDelegationManager.SignatureWithExpiry({signature: signature, expiry: type(uint256).max});
        delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, bytes32(0));
    }

    /// @notice  tries delegating using a wallet that does not comply with EIP 1271
    function testFuzz_DelegateToBySignature_WithContractWallet_NonconformingWallet(
        address operator,
        uint96 ethAmount,
        uint96 eigenAmount,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) public fuzzedAddress(operator) {
        address staker = vm.addr(PRIVATE_KEY);

        // deploy non ERC1271-compliant wallet for staker to use
        vm.startPrank(staker);
        ERC1271MaliciousMock wallet = new ERC1271MaliciousMock();
        vm.stopPrank();
        staker = address(wallet);

        _registerOperatorAndDepositFromStaker(operator, staker, ethAmount, eigenAmount);

        vm.assume(staker != operator);

        bytes memory signature = abi.encodePacked(r, s, v);

        vm.expectRevert(); //SIGNATURE_INVALID_ERROR
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry =
            IDelegationManager.SignatureWithExpiry({signature: signature, expiry: type(uint256).max});
        delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, bytes32(0));
    }

    /// @notice tests delegation to EigenLayer via an ECDSA signatures with invalid signature
    /// @param operator is the operator being delegated to.
    function testFuzz_DelegateToByInvalidSignature(
        address operator,
        uint96 ethAmount,
        uint96 eigenAmount,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) public fuzzedAddress(operator) fuzzedAmounts(ethAmount, eigenAmount) {
        address staker = vm.addr(PRIVATE_KEY);
        _registerOperatorAndDepositFromStaker(operator, staker, ethAmount, eigenAmount);

        bytes memory signature = abi.encodePacked(r, s, v);

        vm.expectRevert(); //SIGNATURE_INVALID_ERROR
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry =
            IDelegationManager.SignatureWithExpiry({signature: signature, expiry: type(uint256).max});
        delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, bytes32(0));
    }

    /// @notice registers a fixed address as a delegate, delegates to it from a second address,
    /// and checks that the delegate's voteWeights increase properly
    /// @param operator is the operator being delegated to.
    /// @param staker is the staker delegating stake to the operator.
    function testFuzz_DelegationMultipleStrategies(uint8 numStratsToAdd, address operator, address staker)
        public
        fuzzedAddress(operator)
        fuzzedAddress(staker)
    {
        vm.assume(staker != operator);

        vm.assume(numStratsToAdd > 0 && numStratsToAdd <= 20);
        uint96 operatorEthWeightBefore = voteWeigher.weightOfOperator(operator, 0);
        uint96 operatorEigenWeightBefore = voteWeigher.weightOfOperator(operator, 1);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        _testRegisterAsOperator(operator, operatorDetails);
        _testDepositStrategies(staker, 1e18, numStratsToAdd);

        // add strategies to voteWeigher
        uint96 multiplier = 1e18;
        for (uint16 i = 0; i < numStratsToAdd; ++i) {
            VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] memory ethStratsAndMultipliers =
                new VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[](1);
            ethStratsAndMultipliers[0].strategy = strategies[i];
            ethStratsAndMultipliers[0].multiplier = multiplier;
            vm.startPrank(voteWeigher.serviceManager().owner());
            voteWeigher.addStrategiesConsideredAndMultipliers(0, ethStratsAndMultipliers);
            vm.stopPrank();
        }

        _testDepositEigen(staker, 1e18);
        _testDelegateToOperator(staker, operator);
        uint96 operatorEthWeightAfter = voteWeigher.weightOfOperator(operator, 0);
        uint96 operatorEigenWeightAfter = voteWeigher.weightOfOperator(operator, 1);
        assertTrue(operatorEthWeightAfter > operatorEthWeightBefore, "operatorEthWeight did not increase!");
        assertTrue(operatorEigenWeightAfter > operatorEigenWeightBefore, "operatorEthWeight did not increase!");
    }

    /// @notice This function tests to ensure that a delegation contract
    ///         cannot be intitialized multiple times
    function testCannotInitMultipleTimesDelegation() public cannotReinit {
        //delegation has already been initialized in the Deployer test contract
        delegation.initialize(address(this), eigenLayerPauserReg, 0);
    }

    /// @notice This function tests to ensure that a you can't register as a delegate multiple times
    /// @param operator is the operator being delegated to.
    function testFuzz_RegisterAsOperatorMultipleTimes(address operator) public fuzzedAddress(operator) {
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        _testRegisterAsOperator(operator, operatorDetails);
        vm.expectRevert(OPERATOR_ALREADY_REGISTERED_ERROR);
        _testRegisterAsOperator(operator, operatorDetails);
    }

    /// @notice This function tests to ensure that a staker cannot delegate to an unregistered operator
    /// @param delegate is the unregistered operator
    function testFuzz_DelegationToUnregisteredDelegate(address delegate) public fuzzedAddress(delegate) {
        //deposit into 1 strategy for getOperatorAddress(1), who is delegating to the unregistered operator
        _testDepositStrategies(getOperatorAddress(1), 1e18, 1);
        _testDepositEigen(getOperatorAddress(1), 1e18);

        vm.expectRevert(OPERATOR_NOT_REGISTERED_ERROR);
        vm.startPrank(getOperatorAddress(1));
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegation.delegateTo(delegate, signatureWithExpiry, bytes32(0));
        vm.stopPrank();
    }

    /// @notice This function tests to ensure that a delegation contract
    ///         cannot be intitialized multiple times, test with different caller addresses
    function testFuzz_CannotInitMultipleTimesDelegation(address _attacker) public {
        vm.assume(_attacker != address(eigenLayerProxyAdmin));
        //delegation has already been initialized in the Deployer test contract
        vm.prank(_attacker);
        vm.expectRevert(ALREADY_INITIALIZED_ERROR);
        delegation.initialize(_attacker, eigenLayerPauserReg, 0);
    }

    /// @notice This function tests that the earningsReceiver cannot be set to address(0)
    function test_CannotSetEarningsReceiverToZeroAddress() public {
        vm.expectRevert(EARNINGS_RECEIVER_ADDRESS_ZERO_ERROR);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: address(0),
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        string memory emptyStringForMetadataURI;
        delegation.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
    }

    /// @notice This function tests to ensure that an address can only call registerAsOperator() once
    function testFuzz_CannotRegisterAsOperatorTwice(address _operator, address _dt)
        public
        fuzzedAddress(_operator)
        fuzzedAddress(_dt)
    {
        vm.assume(_dt != address(0));
        vm.startPrank(_operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: msg.sender,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        string memory emptyStringForMetadataURI;
        delegation.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        vm.expectRevert(OPERATOR_ALREADY_REGISTERED_ERROR);
        delegation.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        vm.stopPrank();
    }

    /// @notice This function checks that you can only delegate to an address that is already registered.
    function testFuzz_DelegateToInvalidOperator(address _staker, address _unregisteredOperator)
        public
        fuzzedAddress(_staker)
    {
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        vm.startPrank(_staker);
        vm.expectRevert(OPERATOR_NOT_REGISTERED_ERROR);
        delegation.delegateTo(_unregisteredOperator, signatureWithExpiry, bytes32(0));
        vm.expectRevert(OPERATOR_NOT_REGISTERED_ERROR);
        delegation.delegateTo(_staker, signatureWithExpiry, bytes32(0));
        vm.stopPrank();
    }

    function testFuzz_Undelegate_SigP_Version(address _operator, address _staker, address _dt) public {
        vm.assume(_operator != address(0));
        vm.assume(_staker != address(0));
        vm.assume(_operator != _staker);
        vm.assume(_dt != address(0));
        vm.assume(_operator != address(eigenLayerProxyAdmin));
        vm.assume(_staker != address(eigenLayerProxyAdmin));

        //setup delegation
        vm.prank(_operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: _dt,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        string memory emptyStringForMetadataURI;
        delegation.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        vm.prank(_staker);
        IDelegationManager.SignatureWithExpiry memory signatureWithExpiry;
        delegation.delegateTo(_operator, signatureWithExpiry, bytes32(0));

        //operators cannot undelegate from themselves
        vm.prank(address(strategyManager));
        vm.expectRevert(OPERATOR_UNDELEGATION_ERROR);
        delegation.undelegate(_operator);

        //_staker cannot undelegate themselves
        vm.prank(_staker);
        vm.expectRevert();
        delegation.undelegate(_operator);

        //_operator cannot undelegate themselves
        vm.prank(_operator);
        vm.expectRevert();
        delegation.undelegate(_operator);

        //assert still delegated
        assertTrue(delegation.isDelegated(_staker));
        assertTrue(delegation.isOperator(_operator));

        //strategyManager can undelegate _staker
        vm.prank(address(strategyManager));
        delegation.undelegate(_staker);
        assertFalse(delegation.isDelegated(_staker));
    }

    function _registerAdditionalOperator(uint256 index, uint32 _serveUntil) internal {
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
        vm.startPrank(sender);

        //whitelist the serviceManager to slash the operator
        slasher.optIntoSlashing(address(serviceManager));

        voteWeigher.registerOperator(sender, _serveUntil);

        vm.stopPrank();
    }

    // registers the operator if they are not already registered, and deposits "WETH" + "EIGEN" on behalf of the staker.
    function _registerOperatorAndDepositFromStaker(
        address operator,
        address staker,
        uint96 ethAmount,
        uint96 eigenAmount
    ) internal {
        vm.assume(staker != operator);

        // if first deposit amount to base strategy is too small, it will revert. ignore that case here.
        vm.assume(ethAmount >= 1 && ethAmount <= 1e18);
        vm.assume(eigenAmount >= 1 && eigenAmount <= 1e18);

        if (!delegation.isOperator(operator)) {
            IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
                earningsReceiver: operator,
                delegationApprover: address(0),
                stakerOptOutWindowBlocks: 0
            });
            _testRegisterAsOperator(operator, operatorDetails);
        }

        //making additional deposits to the strategies
        assertTrue(!delegation.isDelegated(staker), "staker is not delegate");
        _testDepositWeth(staker, ethAmount);
        _testDepositEigen(staker, eigenAmount);
    }
}
