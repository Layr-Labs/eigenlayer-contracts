// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./../EigenPod.t.sol";
import "./../utils/ProofParsing.sol";
import "./../mocks/StrategyManagerMock.sol";
import "./../mocks/SlasherMock.sol";
import "./../mocks/DelegationManagerMock.sol";
import "./../mocks/DelayedWithdrawalRouterMock.sol";

import "forge-std/Test.sol";

contract EigenPodUnitTests is Test, ProofParsing {
    
    using BytesLib for bytes;

    uint256 internal constant GWEI_TO_WEI = 1e9;

    bytes pubkey =
        hex"88347ed1c492eedc97fc8c506a35d44d81f27a0c7a1c661b35913cfd15256c0cccbd34a83341f505c7de2983292f2cab";
    uint40 validatorIndex0 = 0;
    uint40 validatorIndex1 = 1;

    address podOwner = address(42000094993494);

    Vm cheats = Vm(HEVM_ADDRESS);

    ProxyAdmin public eigenLayerProxyAdmin;
    IEigenPodManager public eigenPodManager;
    IEigenPod public podImplementation;
    IDelayedWithdrawalRouter public delayedWithdrawalRouter;
    IETHPOSDeposit public ethPOSDeposit;
    IBeacon public eigenPodBeacon;
    EPInternalFunctions public podInternalFunctionTester;

    IDelegationManager public delegation;
    IStrategyManager public strategyManager;
    ISlasher public slasher;
    IEigenPod public pod;
    PauserRegistry public pauserReg;

    BeaconChainOracleMock public beaconChainOracle;
    address[] public slashingContracts;
    address pauser = address(69);
    address unpauser = address(489);
    address podManagerAddress = 0x212224D2F2d262cd093eE13240ca4873fcCBbA3C;
    address podAddress = address(123);
    uint256 stakeAmount = 32e18;
    mapping(address => bool) fuzzedAddressMapping;
    bytes signature;
    bytes32 depositDataRoot;

    bytes32[] withdrawalFields;
    bytes32[] validatorFields;

    uint32 WITHDRAWAL_DELAY_BLOCKS = 7 days / 12 seconds;
    uint64 MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR = 31e9;
    uint64 RESTAKED_BALANCE_OFFSET_GWEI = 75e7;
    uint64 internal constant GOERLI_GENESIS_TIME = 1616508000;
    uint64 internal constant SECONDS_PER_SLOT = 12;

    EigenPodTests test;

    // EIGENPOD EVENTS
   /// @notice Emitted when an ETH validator stakes via this eigenPod
    event EigenPodStaked(bytes pubkey);

    /// @notice Emitted when an ETH validator's withdrawal credentials are successfully verified to be pointed to this eigenPod
    event ValidatorRestaked(uint40 validatorIndex);

    /// @notice Emitted when an ETH validator's balance is updated in EigenLayer
    event ValidatorBalanceUpdated(uint40 validatorIndex, uint64 balanceTimestamp, uint64 newBalanceGwei);

    /// @notice Emitted when an ETH validator is prove to have withdrawn from the beacon chain
    event FullWithdrawalRedeemed(
        uint40 validatorIndex,
        uint64 withdrawalTimestamp,
        address indexed recipient,
        uint64 withdrawalAmountGwei
    );

    /// @notice Emitted when a partial withdrawal claim is successfully redeemed
    event PartialWithdrawalRedeemed(
        uint40 validatorIndex,
        uint64 withdrawalTimestamp,
        address indexed recipient,
        uint64 partialWithdrawalAmountGwei
    );

    modifier fuzzedAddress(address addr) virtual {
        cheats.assume(fuzzedAddressMapping[addr] == false);
        _;
    }

    function setUp() public {
        ethPOSDeposit = new ETHPOSDepositMock();
        beaconChainOracle = new BeaconChainOracleMock();
        EmptyContract emptyContract = new EmptyContract();
         // deploy proxy admin for ability to upgrade proxy contracts
        eigenLayerProxyAdmin = new ProxyAdmin();

        // this contract is deployed later to keep its address the same (for these tests)
        eigenPodManager = EigenPodManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(eigenLayerProxyAdmin), ""))
        );

        delayedWithdrawalRouter = new DelayedWithdrawalRouterMock();

        podImplementation = new EigenPod(
            ethPOSDeposit,
            delayedWithdrawalRouter,
            eigenPodManager,
            MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            GOERLI_GENESIS_TIME
        );


        eigenPodBeacon = new UpgradeableBeacon(address(podImplementation));

        test = new EigenPodTests();
        eigenPodBeacon = new UpgradeableBeacon(address(podImplementation));
        strategyManager = new StrategyManagerMock();
        slasher = new SlasherMock();
        delegation = new DelegationManagerMock();
         // deploy pauser registry
        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserReg = new PauserRegistry(pausers, unpauser);

        EigenPodManager eigenPodManagerImplementation = new EigenPodManager(
            ethPOSDeposit,
            eigenPodBeacon,
            strategyManager,
            slasher,
            delegation
        );

        eigenLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(eigenPodManager))),
            address(eigenPodManagerImplementation),
            abi.encodeWithSelector(
                EigenPodManager.initialize.selector,
                type(uint256).max, // maxPods
                beaconChainOracle,
                address(this),
                pauserReg,
                0 /*initialPausedStatus*/
            )
        );

        cheats.startPrank(podOwner);
        eigenPodManager.createPod();
        cheats.stopPrank();
        pod = eigenPodManager.getPod(podOwner);
    }

    function testStakingWithInvalidAmount () public {
        cheats.deal(address(eigenPodManager), 10e18);
        cheats.startPrank(address(eigenPodManager));
        cheats.expectRevert(bytes("EigenPod.stake: must initially stake for any validator with 32 ether"));
        pod.stake{value: 10e18}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
    }

    function testFullWithdrawalProofWithWrongWithdrawalFields(bytes32[] memory wrongWithdrawalFields) public {
        Relayer relay = new Relayer();
        uint256  WITHDRAWAL_FIELD_TREE_HEIGHT = 2;

        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        BeaconChainProofs.WithdrawalProof memory proofs = _getWithdrawalProof();
        bytes32 beaconStateRoot = getBeaconStateRoot();
        cheats.assume(wrongWithdrawalFields.length !=  2 ** WITHDRAWAL_FIELD_TREE_HEIGHT);
        validatorFields = getValidatorFields();

        cheats.expectRevert(bytes("BeaconChainProofs.verifyWithdrawal: withdrawalFields has incorrect length"));
        relay.verifyWithdrawal(beaconStateRoot, wrongWithdrawalFields, proofs);
    }

    function testCheckThatHasRestakedIsSetToTrue() public returns (IEigenPod){
        require(pod.hasRestaked() == true, "Pod should not be restaked");
        return pod;
    }

    function testActivateRestakingWithM2Pods() external {
        IEigenPod pod = testCheckThatHasRestakedIsSetToTrue();
        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        pod.activateRestaking();
        cheats.stopPrank(); 
    }

    function testWithdrawBeforeRestakingWithM2Pods() external {
        IEigenPod pod = testCheckThatHasRestakedIsSetToTrue();
        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        pod.withdrawBeforeRestaking();
        cheats.stopPrank(); 
    }

    function testAttemptedWithdrawalAfterVerifyingWithdrawalCredentials() public {
        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        IEigenPod(pod).withdrawBeforeRestaking();
        cheats.stopPrank();
    }

    // function testTooSoonBalanceUpdate(uint64 oracleTimestamp, uint64 mostRecentBalanceUpdateTimestamp) external {
    //     cheats.assume(oracleTimestamp < mostRecentBalanceUpdateTimestamp);
    //     _deployInternalFunctionTester();

    //     setJSON("./src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
    //     validatorFields = getValidatorFields();

    //     uint40[] memory validatorIndices = new uint40[](1);
    //     validatorIndices[0] = uint40(getValidatorIndex());
    //     BeaconChainProofs.BalanceUpdateProof memory proof = _getBalanceUpdateProof();

    //     bytes32 newBeaconStateRoot = getBeaconStateRoot();
    //     emit log_named_bytes32("newBeaconStateRoot", newBeaconStateRoot);
    //     BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(newBeaconStateRoot);

    //     cheats.expectRevert(
    //         bytes("EigenPod.verifyBalanceUpdate: Validators balance has already been updated for this timestamp")
    //     );
    //     podInternalFunctionTester.verifyBalanceUpdate(
    //         oracleTimestamp,
    //         0,
    //         bytes32(0),
    //         proof,
    //         validatorFields,
    //         mostRecentBalanceUpdateTimestamp
    //     );
    // }

    function testProcessFullWithdrawalForLessThanMaxRestakedBalance(uint64 withdrawalAmount) public {
        _deployInternalFunctionTester();
        cheats.assume(withdrawalAmount > 0 && withdrawalAmount < MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR);
        IEigenPod.ValidatorInfo memory validatorInfo = IEigenPod.ValidatorInfo({
            validatorIndex: 0,
            restakedBalanceGwei: 0,
            mostRecentBalanceUpdateTimestamp: 0,
            status: IEigenPod.VALIDATOR_STATUS.ACTIVE
        });
        uint64 balanceBefore = podInternalFunctionTester.withdrawableRestakedExecutionLayerGwei();
        podInternalFunctionTester.processFullWithdrawal(0, bytes32(0), 0, podOwner, withdrawalAmount, validatorInfo);
        require(podInternalFunctionTester.withdrawableRestakedExecutionLayerGwei() - balanceBefore == withdrawalAmount, "withdrawableRestakedExecutionLayerGwei hasn't been updated correctly");
    }
    function testWithdrawBeforeRestakingAfterRestaking() public {
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        cheats.startPrank(podOwner);
        pod.withdrawBeforeRestaking();
        cheats.stopPrank();
    }

    //post M2, all new pods deployed will have "hasRestaked = true".  THis tests that
    function testDeployedPodIsRestaked(address podOwner) public fuzzedAddress(podOwner) {
        cheats.startPrank(podOwner);
        eigenPodManager.createPod();
        cheats.stopPrank();
        IEigenPod pod = eigenPodManager.getPod(podOwner);
        require(pod.hasRestaked() == true, "Pod should be restaked");
    }
    function testTryToActivateRestakingAfterHasRestakedIsSet() public {
        require(pod.hasRestaked() == true, "Pod should be restaked");
        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        pod.activateRestaking();
    }
    function testTryToWithdrawBeforeRestakingAfterHasRestakedIsSet() public {
        require(pod.hasRestaked() == true, "Pod should be restaked");
        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        pod.withdrawBeforeRestaking();
    }
    function testMismatchedWithdrawalProofInputs(uint64 numValidators, uint64 numValidatorProofs) external {
        cheats.assume(numValidators < numValidatorProofs && numValidatorProofs < 5);
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        bytes[] memory validatorFieldsProofArray = new bytes[](numValidatorProofs);
        for (uint256 index = 0; index < numValidators; index++) {
            validatorFieldsProofArray[index] = abi.encodePacked(getValidatorProof());
        }
        bytes32[][] memory validatorFieldsArray = new bytes32[][](numValidators);
        for (uint256 index = 0; index < validatorFieldsArray.length; index++) {
             validatorFieldsArray[index] = getValidatorFields();
        }
        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();
        BeaconChainProofs.WithdrawalProof[] memory withdrawalProofsArray = new BeaconChainProofs.WithdrawalProof[](1);

        withdrawalProofsArray[0] = _getWithdrawalProof();

        bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);
        withdrawalFieldsArray[0] = withdrawalFields;

        cheats.expectRevert(bytes("EigenPod.verifyAndProcessWithdrawals: inputs must be same length"));
        pod.verifyAndProcessWithdrawals(0, stateRootProofStruct, withdrawalProofsArray, validatorFieldsProofArray, validatorFieldsArray, withdrawalFieldsArray);
    }
    function testProveWithdrawalFromBeforeLastWithdrawBeforeRestaking() external {
        cheats.store(address(pod), bytes32(uint256(52)), bytes32(uint256(1)));
        require(pod.hasRestaked() != true, "Pod should not be restaked");
        setJSON("./src/test/test-data/fullWithdrawalProof_Latest.json");
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(getLatestBlockRoot());
        BeaconChainProofs.WithdrawalProof[] memory withdrawalProofsArray = new BeaconChainProofs.WithdrawalProof[](1);
        withdrawalProofsArray[0] = _getWithdrawalProof();

        uint64 timestampOfWithdrawal = Endian.fromLittleEndianUint64(withdrawalProofsArray[0].timestampRoot);
        uint256 newTimestamp = timestampOfWithdrawal + 2500;

        cheats.warp(newTimestamp);
        cheats.startPrank(podOwner);
        pod.withdrawBeforeRestaking();
        cheats.stopPrank();

        bytes[] memory validatorFieldsProofArray = new bytes[](1);
        validatorFieldsProofArray[0] = abi.encodePacked(getValidatorProof());        
        bytes32[][] memory validatorFieldsArray = new bytes32[][](1);
        validatorFieldsArray[0] = getValidatorFields();
        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();
        bytes32[][] memory withdrawalFieldsArray = new bytes32[][](1);
        withdrawalFieldsArray[0] = withdrawalFields;

        cheats.warp(timestampOfWithdrawal);
        cheats.expectRevert(bytes("EigenPod.proofIsForValidTimestamp: beacon chain proof must be for timestamp after mostRecentWithdrawalTimestamp"));
        pod.verifyAndProcessWithdrawals(0, stateRootProofStruct, withdrawalProofsArray, validatorFieldsProofArray, validatorFieldsArray, withdrawalFieldsArray);
    }
    function testPodReceiveFallBack(uint256 amountETH) external {
        cheats.assume(amountETH > 0);
        cheats.deal(address(this), amountETH);
        Address.sendValue(payable(address(pod)), amountETH);
        require(address(pod).balance == amountETH, "Pod should have received ETH");
    }
    /**
    * This is a regression test for a bug (EIG-14) found by Hexens.  Lets say podOwner sends 32 ETH to the EigenPod, 
    * the nonBeaconChainETHBalanceWei increases by 32 ETH. podOwner calls withdrawBeforeRestaking, which 
    * will simply send the entire ETH balance (32 ETH) to the owner. The owner activates restaking, 
    * creates a validator and verifies the withdrawal credentials, receiving 32 ETH in shares.  
    * They can exit the validator, the pod gets the 32ETH and they can call withdrawNonBeaconChainETHBalanceWei
    * And simply withdraw the 32ETH because nonBeaconChainETHBalanceWei is 32ETH.  This was an issue because 
    * nonBeaconChainETHBalanceWei was never zeroed out in _processWithdrawalBeforeRestaking
     */
    function testValidatorBalanceCannotBeRemovedFromPodViaNonBeaconChainETHBalanceWei() external {
        
        uint256 amount = 32 ether;
         cheats.store(address(pod), bytes32(uint256(52)), bytes32(0));
        cheats.deal(address(this), amount);
        // simulate a withdrawal processed on the beacon chain, pod balance goes to 32 ETH
        Address.sendValue(payable(address(pod)), amount);
        require(pod.nonBeaconChainETHBalanceWei() == amount, "nonBeaconChainETHBalanceWei should be 32 ETH");
        //simulate that hasRestaked is set to false, so that we can test withdrawBeforeRestaking for pods deployed before M2 activation
        cheats.store(address(pod), bytes32(uint256(52)), bytes32(uint256(1)));
        //this is an M1 pod so hasRestaked should be false
        require(pod.hasRestaked() == false, "Pod should be restaked");
        cheats.startPrank(podOwner);
        pod.activateRestaking();
         cheats.stopPrank();
        require(pod.nonBeaconChainETHBalanceWei() == 0, "nonBeaconChainETHBalanceWei should be 32 ETH");
    }
    /**
    * Regression test for a bug that allowed balance updates to be made for withdrawn validators.  Thus
    * the validator's balance could be maliciously proven to be 0 before the validator themselves are
    * able to prove their withdrawal.
    */
    // function testBalanceUpdateMadeAfterWithdrawableEpochFails() external {
    //     _deployInternalFunctionTester();
    //     cheats.roll(block.number + 1);
    //     // ./solidityProofGen "BalanceUpdateProof" 302913 true 0 "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_overCommitted_302913.json"
    //     setJSON("./src/test/test-data/balanceUpdateProof_overCommitted_302913.json");
    //     bytes32[] memory validatorFields= getValidatorFields();

    //     uint40 validatorIndex = uint40(getValidatorIndex());

    //     BeaconChainProofs.BalanceUpdateProof memory proof = _getBalanceUpdateProof();
    //     bytes32 newLatestBlockRoot = getLatestBlockRoot();

    //     BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(newLatestBlockRoot);

    //     BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof(); 
    //     proof.balanceRoot = bytes32(uint256(0));     

    //     validatorFields[7] = bytes32(uint256(0));
    //     cheats.warp(GOERLI_GENESIS_TIME + 1 days);
    //     uint64 oracleTimestamp = uint64(block.timestamp);

    //     podInternalFunctionTester.setValidatorStatus(validatorFields[0], IEigenPod.VALIDATOR_STATUS.ACTIVE);
    //     cheats.expectRevert(bytes("EigenPod.verifyBalanceUpdate: validator is withdrawable but has not withdrawn"));
    //     podInternalFunctionTester.verifyBalanceUpdate(oracleTimestamp, validatorIndex, newLatestBlockRoot, proof, validatorFields, 0);
    // }

    function testWithdrawlBeforeRestakingFromNonPodOwnerAddress(uint256 amount, address nonPodOwner) external {
        cheats.assume(nonPodOwner != podOwner);
        uint256 amount = 32 ether;

        cheats.store(address(pod), bytes32(uint256(52)), bytes32(0));

        cheats.startPrank(nonPodOwner);
        cheats.expectRevert(bytes("EigenPod.onlyEigenPodOwner: not podOwner"));
        pod.withdrawBeforeRestaking();
        cheats.stopPrank();  
    }

    function testFullWithdrawalAmounts(bytes32 pubkeyHash, uint64 withdrawalAmount) external {
        _deployInternalFunctionTester();
        IEigenPod.ValidatorInfo memory validatorInfo = IEigenPod.ValidatorInfo({
            validatorIndex: 0,
            restakedBalanceGwei: 0,
            mostRecentBalanceUpdateTimestamp: 0,
            status: IEigenPod.VALIDATOR_STATUS.ACTIVE
        });
        IEigenPod.VerifiedWithdrawal memory vw = podInternalFunctionTester.processFullWithdrawal(0, pubkeyHash, 0, podOwner, withdrawalAmount, validatorInfo);

        if(withdrawalAmount > podInternalFunctionTester.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR()){
            require(vw.amountToSendGwei == withdrawalAmount - podInternalFunctionTester.MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR(), "newAmount should be MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR");
        }
        else{
            require(vw.amountToSendGwei == 0, "newAmount should be withdrawalAmount");
        }
    }

    function testProcessPartialWithdrawal(
        uint40 validatorIndex,
        uint64 withdrawalTimestamp,
        address recipient,
        uint64 partialWithdrawalAmountGwei
    ) external {
        _deployInternalFunctionTester();
        cheats.expectEmit(true, true, true, true, address(podInternalFunctionTester));
        emit PartialWithdrawalRedeemed(
            validatorIndex,
            withdrawalTimestamp,
            recipient,
            partialWithdrawalAmountGwei
        );
        IEigenPod.VerifiedWithdrawal memory vw = podInternalFunctionTester.processPartialWithdrawal(validatorIndex, withdrawalTimestamp, recipient, partialWithdrawalAmountGwei);

        require(vw.amountToSendGwei == partialWithdrawalAmountGwei, "newAmount should be partialWithdrawalAmountGwei");
    }

    function testRecoverTokens(uint256 amount, address recipient) external {
        cheats.assume(amount > 0 && amount < 1e30);  
        IERC20 randomToken = new ERC20PresetFixedSupply(
            "rand",
            "RAND",
            1e30,
            address(this)
        );

        IERC20[] memory tokens = new IERC20[](1);
        tokens[0] = randomToken;
        uint256[] memory amounts = new uint256[](1);
        amounts[0] = amount;

        randomToken.transfer(address(pod), amount);
        require(randomToken.balanceOf(address(pod)) == amount, "randomToken balance should be amount");

        uint256 recipientBalanceBefore = randomToken.balanceOf(recipient);

        cheats.startPrank(podOwner);
        pod.recoverTokens(tokens, amounts, recipient);
        cheats.stopPrank();
        require(randomToken.balanceOf(address(recipient)) - recipientBalanceBefore == amount, "recipient should have received amount");
    }

    function testRecoverTokensMismatchedInputs() external {
        uint256 tokenListLen = 5;
        uint256 amountsToWithdrawLen = 2;

        IERC20[] memory tokens = new IERC20[](tokenListLen);
        uint256[] memory amounts = new uint256[](amountsToWithdrawLen);

        cheats.expectRevert(bytes("EigenPod.recoverTokens: tokenList and amountsToWithdraw must be same length"));
        cheats.startPrank(podOwner);
        pod.recoverTokens(tokens, amounts, address(this));
        cheats.stopPrank();
    }

    function _deployInternalFunctionTester() internal {
        podInternalFunctionTester = new EPInternalFunctions(
            ethPOSDeposit,
            delayedWithdrawalRouter,
            IEigenPodManager(podManagerAddress),
            MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
            GOERLI_GENESIS_TIME
        );
    }

    function _getStateRootProof() internal returns (BeaconChainProofs.StateRootProof memory) {
        return BeaconChainProofs.StateRootProof(getBeaconStateRoot(), abi.encodePacked(getStateRootProof()));
    }

    /// @notice this function just generates a valid proof so that we can test other functionalities of the withdrawal flow
    function _getWithdrawalProof() internal returns (BeaconChainProofs.WithdrawalProof memory) {

        {
            bytes32 blockRoot = getBlockRoot();
            bytes32 slotRoot = getSlotRoot();
            bytes32 timestampRoot = getTimestampRoot();
            bytes32 executionPayloadRoot = getExecutionPayloadRoot();

            return
                BeaconChainProofs.WithdrawalProof(
                    abi.encodePacked(getWithdrawalProof()),
                    abi.encodePacked(getSlotProof()),
                    abi.encodePacked(getExecutionPayloadProof()),
                    abi.encodePacked(getTimestampProof()),
                    abi.encodePacked(getHistoricalSummaryProof()),
                    uint64(getBlockRootIndex()),
                    uint64(getHistoricalSummaryIndex()),
                    uint64(getWithdrawalIndex()),
                    blockRoot,
                    slotRoot,
                    timestampRoot,
                    executionPayloadRoot
                );
        }
    }

    function _deployPod() internal {
        cheats.startPrank(podOwner);
        eigenPodManager.createPod();
        cheats.stopPrank();
    }
}