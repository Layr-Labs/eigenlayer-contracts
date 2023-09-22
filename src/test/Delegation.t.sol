// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/utils/math/Math.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/mocks/ERC1271WalletMock.sol";
import "src/contracts/interfaces/ISignatureUtils.sol";

import "../test/EigenLayerTestHelper.t.sol";

import "./mocks/MiddlewareRegistryMock.sol";
import "./mocks/ServiceManagerMock.sol";

import "./harnesses/StakeRegistryHarness.sol";

contract DelegationTests is EigenLayerTestHelper {
    using Math for uint256;

    uint256 public PRIVATE_KEY = 420;

    uint32 serveUntil = 100;

    address public registryCoordinator = address(uint160(uint256(keccak256("registryCoordinator"))));
    ServiceManagerMock public serviceManager;
    StakeRegistryHarness public stakeRegistry;
    StakeRegistryHarness public stakeRegistryImplementation;
    uint8 defaultQuorumNumber = 0;
    bytes32 defaultOperatorId = bytes32(uint256(0));

    modifier fuzzedAmounts(uint256 ethAmount, uint256 eigenAmount) {
        cheats.assume(ethAmount >= 0 && ethAmount <= 1e18);
        cheats.assume(eigenAmount >= 0 && eigenAmount <= 1e18);
        _;
    }

    function setUp() public virtual override {
        EigenLayerDeployer.setUp();

        initializeMiddlewares();
    }

    function initializeMiddlewares() public {
        serviceManager = new ServiceManagerMock(slasher);

        stakeRegistry = StakeRegistryHarness(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );
        stakeRegistryImplementation = new StakeRegistryHarness(
            IRegistryCoordinator(registryCoordinator),
            strategyManager,
            serviceManager
        );

        {
            uint96 multiplier = 1e18;
            uint8 _NUMBER_OF_QUORUMS = 2;
            // uint256[] memory _quorumBips = new uint256[](_NUMBER_OF_QUORUMS);
            // // split 60% ETH quorum, 40% EIGEN quorum
            // _quorumBips[0] = 6000;
            // _quorumBips[1] = 4000;
            // IVoteWeigher.StrategyAndWeightingMultiplier[] memory ethStratsAndMultipliers =
            //     new IVoteWeigher.StrategyAndWeightingMultiplier[](1);
            // ethStratsAndMultipliers[0].strategy = wethStrat; 
            // ethStratsAndMultipliers[0].multiplier = multiplier;
            // IVoteWeigher.StrategyAndWeightingMultiplier[] memory eigenStratsAndMultipliers =
            //     new IVoteWeigher.StrategyAndWeightingMultiplier[](1);
            // eigenStratsAndMultipliers[0].strategy = eigenStrat;
            // eigenStratsAndMultipliers[0].multiplier = multiplier;

            cheats.startPrank(eigenLayerProxyAdmin.owner());

            // setup the dummy minimum stake for quorum
            uint96[] memory minimumStakeForQuorum = new uint96[](_NUMBER_OF_QUORUMS);
            for (uint256 i = 0; i < minimumStakeForQuorum.length; i++) {
                minimumStakeForQuorum[i] = uint96(i+1);
            }

            // setup the dummy quorum strategies
            IVoteWeigher.StrategyAndWeightingMultiplier[][] memory quorumStrategiesConsideredAndMultipliers =
                new IVoteWeigher.StrategyAndWeightingMultiplier[][](2);
            quorumStrategiesConsideredAndMultipliers[0] = new IVoteWeigher.StrategyAndWeightingMultiplier[](1);
            quorumStrategiesConsideredAndMultipliers[0][0] = IVoteWeigher.StrategyAndWeightingMultiplier(
                wethStrat,
                multiplier
            );
            quorumStrategiesConsideredAndMultipliers[1] = new IVoteWeigher.StrategyAndWeightingMultiplier[](1);
            quorumStrategiesConsideredAndMultipliers[1][0] = IVoteWeigher.StrategyAndWeightingMultiplier(
                eigenStrat,
                multiplier
            );

            eigenLayerProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(payable(address(stakeRegistry))),
                address(stakeRegistryImplementation),
                abi.encodeWithSelector(StakeRegistry.initialize.selector, minimumStakeForQuorum, quorumStrategiesConsideredAndMultipliers) 
            );
            cheats.stopPrank();
        
        }
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
    function testDelegation(address operator, address staker, uint96 ethAmount, uint96 eigenAmount)
        public
        fuzzedAddress(operator)
        fuzzedAddress(staker)
        fuzzedAmounts(ethAmount, eigenAmount)
    {
        cheats.assume(staker != operator);
        // base strategy will revert if these amounts are too small on first deposit
        cheats.assume(ethAmount >= 1);
        cheats.assume(eigenAmount >= 2);
        
        // Set weights ahead of the helper function call
        bytes memory quorumNumbers = new bytes(2);
        quorumNumbers[0] = bytes1(uint8(0));
        quorumNumbers[0] = bytes1(uint8(1));
        stakeRegistry.setOperatorWeight(0, operator, ethAmount);
        stakeRegistry.setOperatorWeight(1, operator, eigenAmount);
        stakeRegistry.registerOperator(operator, defaultOperatorId, quorumNumbers);
        _testDelegation(operator, staker, ethAmount, eigenAmount, quorumNumbers, stakeRegistry);
    }

    /// @notice tests that a when an operator is delegated to, that delegation is properly accounted for.
    function testDelegationReceived(address _operator, address staker, uint64 ethAmount, uint64 eigenAmount)
        public
        fuzzedAddress(_operator)
        fuzzedAddress(staker)
        fuzzedAmounts(ethAmount, eigenAmount)
    {
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

        uint256[3] memory amountsBefore;
        amountsBefore[0] = stakeRegistry.weightOfOperatorForQuorum(0, operator);
        amountsBefore[1] = stakeRegistry.weightOfOperatorForQuorum(1, operator);
        amountsBefore[2] = delegation.operatorShares(operator, wethStrat);

        //making additional deposits to the  strategies
        assertTrue(!delegation.isDelegated(staker) == true, "testDelegation: staker is not delegate");
        _testDepositWeth(staker, ethAmount);
        _testDepositEigen(staker, eigenAmount);
        _testDelegateToOperator(staker, operator);
        stakeRegistry.setOperatorWeight(0, operator, ethAmount);
        stakeRegistry.setOperatorWeight(1, operator, eigenAmount);
        assertTrue(delegation.isDelegated(staker) == true, "testDelegation: staker is not delegate");

        (IStrategy[] memory updatedStrategies, uint256[] memory updatedShares) =
            strategyManager.getDeposits(staker);

        {
            uint256 stakerEthWeight = strategyManager.stakerStrategyShares(staker, updatedStrategies[0]);
            uint256 stakerEigenWeight = strategyManager.stakerStrategyShares(staker, updatedStrategies[1]);

            uint256 operatorEthWeightAfter = stakeRegistry.weightOfOperatorForQuorum(0, operator);
            uint256 operatorEigenWeightAfter = stakeRegistry.weightOfOperatorForQuorum(1, operator);

            assertTrue(
                operatorEthWeightAfter - amountsBefore[0] == stakerEthWeight,
                "testDelegation: operatorEthWeight did not increment by the right amount"
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

            cheats.startPrank(address(strategyManager));

            IDelegationManager.OperatorDetails memory expectedOperatorDetails = delegation.operatorDetails(operator);
            assertTrue(keccak256(abi.encode(expectedOperatorDetails)) == keccak256(abi.encode(operatorDetails)),
            "failed to set correct operator details");
        }
    }

    /// @notice tests that a when an operator is undelegated from, that the staker is properly classified as undelegated.
    function testUndelegation(address operator, address staker, uint96 ethAmount, uint96 eigenAmount)
        public
        fuzzedAddress(operator)
        fuzzedAddress(staker)
        fuzzedAmounts(ethAmount, eigenAmount)
    {
        cheats.assume(staker != operator);
        // base strategy will revert if these amounts are too small on first deposit
        cheats.assume(ethAmount >= 1);
        cheats.assume(eigenAmount >= 1);
        bytes memory quorumNumbers = new bytes(2);
        quorumNumbers[0] = bytes1(uint8(0));
        quorumNumbers[0] = bytes1(uint8(1));
        _testDelegation(operator, staker, ethAmount, eigenAmount, quorumNumbers, stakeRegistry);
        cheats.startPrank(address(strategyManager));
        delegation.undelegate(staker);
        cheats.stopPrank();

        require(delegation.delegatedTo(staker) == address(0), "undelegation unsuccessful");
    }

    /// @notice tests delegation from a staker to operator via ECDSA signature.  
    function testDelegateToBySignature(address operator, uint96 ethAmount, uint96 eigenAmount, uint256 expiry)
        public
        fuzzedAddress(operator)
    {
        address staker = cheats.addr(PRIVATE_KEY);
        _registerOperatorAndDepositFromStaker(operator, staker, ethAmount, eigenAmount); 

        uint256 nonceBefore = delegation.stakerNonce(staker);

        bytes32 structHash = keccak256(abi.encode(delegation.STAKER_DELEGATION_TYPEHASH(), staker, operator, nonceBefore, expiry));
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
    function testDelegateToBySignature_WithContractWallet_Successfully(address operator, uint96 ethAmount, uint96 eigenAmount)
        public
        fuzzedAddress(operator)
    {
        address staker = cheats.addr(PRIVATE_KEY);

        // deploy ERC1271WalletMock for staker to use
        cheats.startPrank(staker);
        ERC1271WalletMock wallet = new ERC1271WalletMock(staker);
        cheats.stopPrank();
        staker = address(wallet);

        _registerOperatorAndDepositFromStaker(operator, staker, ethAmount, eigenAmount); 
    
        uint256 nonceBefore = delegation.stakerNonce(staker);

        bytes32 structHash = keccak256(abi.encode(delegation.STAKER_DELEGATION_TYPEHASH(), staker, operator, nonceBefore, type(uint256).max));
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
    function testDelegateToBySignature_WithContractWallet_BadSignature(address operator, uint96 ethAmount, uint96 eigenAmount)
        public
        fuzzedAddress(operator)
    {
        address staker = cheats.addr(PRIVATE_KEY);

        // deploy ERC1271WalletMock for staker to use
        cheats.startPrank(staker);
        ERC1271WalletMock wallet = new ERC1271WalletMock(staker);
        cheats.stopPrank();
        staker = address(wallet);

        _registerOperatorAndDepositFromStaker(operator, staker, ethAmount, eigenAmount); 
    
        uint256 nonceBefore = delegation.stakerNonce(staker);

        bytes32 structHash = keccak256(abi.encode(delegation.STAKER_DELEGATION_TYPEHASH(), staker, operator, nonceBefore, type(uint256).max));
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", delegation.domainSeparator(), structHash));

        bytes memory signature;
        {
            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(PRIVATE_KEY, digestHash);
            // mess up the signature by flipping v's parity
            v = (v == 27 ? 28 : 27);
            signature = abi.encodePacked(r, s, v);
        }
    
        cheats.expectRevert(bytes("EIP1271SignatureUtils.checkSignature_EIP1271: ERC1271 signature verification failed"));
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry = ISignatureUtils.SignatureWithExpiry({
            signature: signature,
            expiry: type(uint256).max
        });
        delegation.delegateToBySignature(staker, operator, signatureWithExpiry, signatureWithExpiry, bytes32(0));
    }

    /// @notice  tries delegating using a wallet that does not comply with EIP 1271
    function testDelegateToBySignature_WithContractWallet_NonconformingWallet(address operator, uint96 ethAmount, uint96 eigenAmount, uint8 v, bytes32 r, bytes32 s)
        public
        fuzzedAddress(operator)
    {
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
    )
        public
        fuzzedAddress(operator)
        fuzzedAmounts(ethAmount, eigenAmount)
    {
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

    /// @notice registers a fixed address as a delegate, delegates to it from a second address,
    /// and checks that the delegate's voteWeights increase properly
    /// @param operator is the operator being delegated to.
    /// @param staker is the staker delegating stake to the operator.
    function testDelegationMultipleStrategies(uint8 numStratsToAdd, address operator, address staker)
    public
        fuzzedAddress(operator)
        fuzzedAddress(staker)
    {
        cheats.assume(staker != operator);

        cheats.assume(numStratsToAdd > 0 && numStratsToAdd <= 20);
        uint256 amountToDeposit = 1e18;
        uint96 operatorEthWeightBefore = stakeRegistry.weightOfOperatorForQuorum(0, operator);
        uint96 operatorEigenWeightBefore = stakeRegistry.weightOfOperatorForQuorum(1, operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        _testRegisterAsOperator(operator, operatorDetails);
        _testDepositStrategies(staker, amountToDeposit, numStratsToAdd);

        // add strategies to voteWeigher
        uint96 multiplier = 1e18;
        for (uint16 i = 0; i < numStratsToAdd; ++i) {
            IVoteWeigher.StrategyAndWeightingMultiplier[] memory ethStratsAndMultipliers =
            new IVoteWeigher.StrategyAndWeightingMultiplier[](1);
            ethStratsAndMultipliers[0].strategy = strategies[i];
            ethStratsAndMultipliers[0].multiplier = multiplier;
            cheats.startPrank(stakeRegistry.serviceManager().owner());
            stakeRegistry.addStrategiesConsideredAndMultipliers(0, ethStratsAndMultipliers);
            cheats.stopPrank();
        }
        _testDepositEigen(staker, amountToDeposit);
        _testDelegateToOperator(staker, operator);
        stakeRegistry.setOperatorWeight(0, operator, uint96(amountToDeposit));
        stakeRegistry.setOperatorWeight(1, operator, uint96(amountToDeposit));
        uint96 operatorEthWeightAfter = stakeRegistry.weightOfOperatorForQuorum(0, operator);
        uint96 operatorEigenWeightAfter = stakeRegistry.weightOfOperatorForQuorum(1, operator);
        assertTrue(
            operatorEthWeightAfter > operatorEthWeightBefore, "testDelegation: operatorEthWeight did not increase!"
        );
        assertTrue(
            operatorEigenWeightAfter > operatorEigenWeightBefore, "testDelegation: operatorEthWeight did not increase!"
        );
}

    /// @notice This function tests to ensure that a delegation contract
    ///         cannot be intitialized multiple times
    function testCannotInitMultipleTimesDelegation() public cannotReinit {
        //delegation has already been initialized in the Deployer test contract
        delegation.initialize(address(this), eigenLayerPauserReg, 0);
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
        delegation.initialize(_attacker, eigenLayerPauserReg, 0);
    }

    /// @notice This function tests that the earningsReceiver cannot be set to address(0)
    function testCannotSetEarningsReceiverToZeroAddress() public{
        cheats.expectRevert(bytes("DelegationManager._setOperatorDetails: cannot set `earningsReceiver` to zero address"));
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: address(0),
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        string memory emptyStringForMetadataURI;
        delegation.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
    }

    /// @notice This function tests to ensure that an address can only call registerAsOperator() once
    function testCannotRegisterAsOperatorTwice(address _operator, address _dt) public fuzzedAddress(_operator) fuzzedAddress(_dt) {
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
    function testDelegateToInvalidOperator(address _staker, address _unregisteredoperator) public fuzzedAddress(_staker) {
        vm.startPrank(_staker);
        cheats.expectRevert(bytes("DelegationManager._delegate: operator is not registered in EigenLayer"));
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry;
        delegation.delegateTo(_unregisteredoperator, signatureWithExpiry, bytes32(0));
        cheats.expectRevert(bytes("DelegationManager._delegate: operator is not registered in EigenLayer"));
        delegation.delegateTo(_staker, signatureWithExpiry, bytes32(0));
        cheats.stopPrank();
    }

    function testUndelegate_SigP_Version(address _operator,address _staker,address _dt) public {

        vm.assume(_operator != address(0));
        vm.assume(_staker != address(0));
        vm.assume(_operator != _staker);
        vm.assume(_dt != address(0));
        vm.assume(_operator != address(eigenLayerProxyAdmin));
        vm.assume(_staker != address(eigenLayerProxyAdmin));

        //setup delegation
        vm.prank(_operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver:_dt,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        string memory emptyStringForMetadataURI;
        delegation.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        vm.prank(_staker);
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry;
        delegation.delegateTo(_operator, signatureWithExpiry, bytes32(0));

        //operators cannot undelegate from themselves
        vm.prank(address(strategyManager));
        cheats.expectRevert(bytes("DelegationManager.undelegate: operators cannot undelegate from themselves"));
        delegation.undelegate(_operator);

        //_staker cannot undelegate themselves
        vm.prank(_staker);
        cheats.expectRevert();
        delegation.undelegate(_operator);
        
        //_operator cannot undelegate themselves
        vm.prank(_operator);
        cheats.expectRevert();
        delegation.undelegate(_operator);

        //assert still delegated
        assertTrue(delegation.isDelegated(_staker));
        assertFalse(!delegation.isDelegated(_staker));
        assertTrue(delegation.isOperator(_operator));

        //strategyManager can undelegate _staker
        vm.prank(address(strategyManager));
        delegation.undelegate(_staker);
        assertFalse(delegation.isDelegated(_staker));
        assertTrue(!delegation.isDelegated(_staker));

    }

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

        //whitelist the serviceManager to slash the operator
        slasher.optIntoSlashing(address(serviceManager));
        // bytes memory defaultQuorumNumber = abi.encodePacked(uint8(0));
        // stakeRegistry.registerOperator(sender, bytes32(index), defaultQuorumNumber);

        cheats.stopPrank();
    }


    // registers the operator if they are not already registered, and deposits "WETH" + "EIGEN" on behalf of the staker.
    function _registerOperatorAndDepositFromStaker(address operator, address staker, uint96 ethAmount, uint96 eigenAmount) internal {
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
