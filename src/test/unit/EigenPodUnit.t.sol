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
    
    function testCheckThatHasRestakedIsSetToTrue() public {
        testStaking();
        IEigenPod pod = eigenPodManager.getPod(podOwner);
        require(pod.hasRestaked() == true, "Pod should not be restaked");
    }

    function testAttemptedWithdrawalAfterVerifyingWithdrawalCredentials() public {
        testDeployAndVerifyNewEigenPod();
        IEigenPod pod = eigenPodManager.getPod(podOwner);
        cheats.startPrank(podOwner);
        cheats.expectRevert(bytes("EigenPod.hasNeverRestaked: restaking is enabled"));
        IEigenPod(pod).withdrawBeforeRestaking();
        cheats.stopPrank();
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
        cheats.assume(numValidators < numValidatorProofs && numValidatorProofs < 100);

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

    // function testDecrementMoreThanRestakedExecutionLayerGwei(uint256 largerAmount) external {
    //     setJSON("./src/test/test-data/withdrawal_credential_proof_302913.json");
    //     _testDeployAndVerifyNewEigenPod(podOwner, signature, depositDataRoot);
    //     IEigenPod pod = eigenPodManager.getPod(podOwner);

    //     cheats.startPrank(address(eigenPodManager));

    //     cheats.expectRevert(bytes("EigenPod.decrementWithdrawableRestakedExecutionLayerGwei: amount to decrement is greater than current withdrawableRestakedRxecutionLayerGwei balance"));
    //     pod.decrementWithdrawableRestakedExecutionLayerGwei(largerAmount);
    // }

}