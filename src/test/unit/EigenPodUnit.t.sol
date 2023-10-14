// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./../EigenPod.t.sol";

contract EigenPodUnitTests is EigenPodTests {
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
        testStaking();
        IEigenPod pod = eigenPodManager.getPod(podOwner);
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
        testDeployAndVerifyNewEigenPod();
        IEigenPod pod = eigenPodManager.getPod(podOwner);
        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        IEigenPod(pod).withdrawBeforeRestaking();
        cheats.stopPrank();
    }

    function testBalanceProofWithWrongTimestamp(uint64 timestamp) public {
        cheats.assume(timestamp > GOERLI_GENESIS_TIME);
        // ./solidityProofGen "BalanceUpdateProof" 302913 false 0 "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_notOverCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
        IEigenPod newPod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);

         // ./solidityProofGen "BalanceUpdateProof" 302913 true 0 "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_overCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_overCommitted_302913.json");
        // prove overcommitted balance
        cheats.warp(timestamp);
        _proveOverCommittedStake(newPod);

        
        validatorFields = getValidatorFields();
        uint40 validatorIndex = uint40(getValidatorIndex());
        bytes32 newLatestBlockRoot = getLatestBlockRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(newLatestBlockRoot);
        BeaconChainProofs.BalanceUpdateProof memory proofs = _getBalanceUpdateProof();
        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof();      

        cheats.expectRevert(bytes("EigenPod.verifyBalanceUpdate: Validators balance has already been updated for this timestamp"));
        newPod.verifyBalanceUpdate(uint64(block.timestamp - 1), validatorIndex, stateRootProofStruct, proofs, validatorFields);
    }

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
        // ./solidityProofGen "ValidatorFieldsProof" 302913 true "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "withdrawal_credential_proof_510257.json"
         setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");

        IEigenPod pod = _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);

        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        cheats.startPrank(podOwner);
        pod.withdrawBeforeRestaking();
        cheats.stopPrank();
    }

    //post M2, all new pods deployed will have "hasRestaked = true".  THis tests that
    function testDeployedPodIsRestaked(address podOwner) public {
        cheats.startPrank(podOwner);
        eigenPodManager.createPod();
        cheats.stopPrank();

        IEigenPod pod = eigenPodManager.getPod(podOwner);
        require(pod.hasRestaked() == true, "Pod should be restaked");
    }

    function testTryToActivateRestakingAfterHasRestakedIsSet() public {
       cheats.startPrank(podOwner);
        eigenPodManager.createPod();
        cheats.stopPrank();

        IEigenPod pod = eigenPodManager.getPod(podOwner);
        require(pod.hasRestaked() == true, "Pod should be restaked");

        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        pod.activateRestaking();

    }

    function testTryToWithdrawBeforeRestakingAfterHasRestakedIsSet() public {
        cheats.startPrank(podOwner);
        eigenPodManager.createPod();
        cheats.stopPrank();

        IEigenPod pod = eigenPodManager.getPod(podOwner);
        require(pod.hasRestaked() == true, "Pod should be restaked");

        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        pod.withdrawBeforeRestaking();
    }

    function testMismatchedWithdrawalProofInputs(uint64 numValidators, uint64 numValidatorProofs) external {
        cheats.assume(numValidators < numValidatorProofs && numValidatorProofs < 5);

        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

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
        newPod.verifyAndProcessWithdrawals(0, stateRootProofStruct, withdrawalProofsArray, validatorFieldsProofArray, validatorFieldsArray, withdrawalFieldsArray);
    }

    function testProveWithdrawalFromBeforeLastWithdrawBeforeRestaking() external {
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod pod = eigenPodManager.getPod(podOwner);

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
        setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod pod = eigenPodManager.getPod(podOwner);
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
        cheats.startPrank(podOwner);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
        
        uint256 amount = 32 ether;

         cheats.store(address(newPod), bytes32(uint256(52)), bytes32(0));
        cheats.deal(address(this), amount);
        // simulate a withdrawal processed on the beacon chain, pod balance goes to 32 ETH
        Address.sendValue(payable(address(newPod)), amount);
        require(newPod.nonBeaconChainETHBalanceWei() == amount, "nonBeaconChainETHBalanceWei should be 32 ETH");
        //simulate that hasRestaked is set to false, so that we can test withdrawBeforeRestaking for pods deployed before M2 activation
        cheats.store(address(newPod), bytes32(uint256(52)), bytes32(uint256(1)));
        //this is an M1 pod so hasRestaked should be false
        require(newPod.hasRestaked() == false, "Pod should be restaked");
        cheats.startPrank(podOwner);
        newPod.activateRestaking();
         cheats.stopPrank();
        require(newPod.nonBeaconChainETHBalanceWei() == 0, "nonBeaconChainETHBalanceWei should be 32 ETH");
    }

    /**
    * Regression test for a bug that allowed balance updates to be made for withdrawn validators.  Thus
    * the validator's balance could be maliciously proven to be 0 before the validator themselves are
    * able to prove their withdrawal.
    */
    function testBalanceUpdateMadeAfterWithdrawableEpochFails() external {
        //make initial deposit
        // ./solidityProofGen "BalanceUpdateProof" 302913 false 0 "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_notOverCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_notOverCommitted_302913.json");
        _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);

        cheats.roll(block.number + 1);
        // ./solidityProofGen "BalanceUpdateProof" 302913 true 0 "data/withdrawal_proof_goerli/goerli_slot_6399999.json"  "data/withdrawal_proof_goerli/goerli_slot_6399998.json" "balanceUpdateProof_overCommitted_302913.json"
        setJSON("./src/test/test-data/balanceUpdateProof_overCommitted_302913.json");
        validatorFields = getValidatorFields();
        uint40 validatorIndex = uint40(getValidatorIndex());
        bytes32 newLatestBlockRoot = getLatestBlockRoot();
        BeaconChainOracleMock(address(beaconChainOracle)).setOracleBlockRootAtTimestamp(newLatestBlockRoot);
        BeaconChainProofs.BalanceUpdateProof memory proofs = _getBalanceUpdateProof();
        BeaconChainProofs.StateRootProof memory stateRootProofStruct = _getStateRootProof(); 
        proofs.balanceRoot = bytes32(uint256(0));     

        validatorFields[7] = bytes32(uint256(0));
        cheats.warp(GOERLI_GENESIS_TIME + 1 days);
        uint64 oracleTimestamp = uint64(block.timestamp);
        cheats.expectRevert(bytes("EigenPod.verifyBalanceUpdate: validator is withdrawable but has not withdrawn"));
        newPod.verifyBalanceUpdate(oracleTimestamp, validatorIndex, stateRootProofStruct, proofs, validatorFields);
    }

    function testWithdrawlBeforeRestakingFromNonPodOwnerAddress(uint256 amount, address nonPodOwner) external {
        cheats.assume(nonPodOwner != podOwner);
        cheats.startPrank(podOwner);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
        
        uint256 amount = 32 ether;

        cheats.store(address(newPod), bytes32(uint256(52)), bytes32(0));

        cheats.startPrank(nonPodOwner);
        cheats.expectRevert(bytes("EigenPod.onlyEigenPodOwner: not podOwner"));
        newPod.withdrawBeforeRestaking();
        cheats.stopPrank();  
    }

    function testDelayedWithdrawalIsCreatedByWithdrawBeforeRestaking(uint256 amount) external {
        cheats.startPrank(podOwner);
        IEigenPod newPod = eigenPodManager.getPod(podOwner);
        cheats.expectEmit(true, true, true, true, address(newPod));
        emit EigenPodStaked(pubkey);
        eigenPodManager.stake{value: stakeAmount}(pubkey, signature, depositDataRoot);
        cheats.stopPrank();
        
        uint256 amount = 32 ether;

        cheats.store(address(newPod), bytes32(uint256(52)), bytes32(0));
        cheats.deal(address(this), amount);
        // simulate a withdrawal processed on the beacon chain, pod balance goes to 32 ETH
        Address.sendValue(payable(address(newPod)), amount);
        require(newPod.nonBeaconChainETHBalanceWei() == amount, "nonBeaconChainETHBalanceWei should be 32 ETH");

        cheats.startPrank(podOwner);
        newPod.withdrawBeforeRestaking();
        cheats.stopPrank();

        require(_getLatestDelayedWithdrawalAmount(podOwner) == amount, "Payment amount should be stake amount");
    }

    function testFullWithdrawalsProvenOutOfOrder(bytes32 pubkeyHash, uint64 withdrawalAmount) external {
        uint64 timestamp = 5;
        uint64 lesserTimestamp = 3;
        _deployInternalFunctionTester();
        IEigenPod.ValidatorInfo memory validatorInfo = IEigenPod.ValidatorInfo({
            validatorIndex: 0,
            restakedBalanceGwei: 0,
            mostRecentBalanceUpdateTimestamp: 0,
            status: IEigenPod.VALIDATOR_STATUS.ACTIVE
        });
        podInternalFunctionTester.processFullWithdrawal(0, pubkeyHash, timestamp += 5 , podOwner, withdrawalAmount, validatorInfo);
        IEigenPod.ValidatorInfo memory info = podInternalFunctionTester.validatorPubkeyHashToInfo(pubkeyHash);

        podInternalFunctionTester.processFullWithdrawal(0, pubkeyHash, lesserTimestamp , podOwner, withdrawalAmount, podInternalFunctionTester.validatorPubkeyHashToInfo(pubkeyHash));

        IEigenPod.ValidatorInfo memory info2 = podInternalFunctionTester.validatorPubkeyHashToInfo(pubkeyHash);

        require(info2.mostRecentBalanceUpdateTimestamp == timestamp, "mostRecentBalanceUpdateTimestamp should not have been updated");
    }
}