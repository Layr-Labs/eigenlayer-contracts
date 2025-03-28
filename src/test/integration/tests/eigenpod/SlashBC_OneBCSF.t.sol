// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/mocks/BeaconChainMock.t.sol";
import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/harnesses/EigenPodManagerWrapper.sol";

/// @notice Testing the rounding behavior when beacon chain slashing factor is initially 1
contract Integration_SlashBC_OneBCSF is IntegrationCheckUtils {
    using ArrayLib for *;

    AVS avs;
    OperatorSet operatorSet;

    User operator;
    IAllocationManagerTypes.AllocateParams allocateParams;

    User staker;
    IStrategy[] strategies;
    uint[] initDepositShares;
    uint40[] validators;
    uint64 beaconBalanceGwei;
    uint64 slashedGwei;

    /**
     * Shared setup:
     * 1. Update the EPM implementation to manually set the beaconChainSlashingFactor to 1
     * Note: Slashing the EigenPod enough (in units of gwei) to reduce the beaconChainSlashingFactor to 1 without
     * rounding down to 0 in a single slash even with increased MaxEB to 2048 is not possible and would require several
     * iterations of slashing to do so. Using a harness to set the beaconChainSlashingFactor to 1 is the easiest way for this test.
     * 2. create a new staker, operator, and avs
     * 3. start validators and verify withdrawal credentials
     */
    function _init() internal override {
        // 1. etch a new implementation to set staker's beaconChainSlashingFactor to 1
        EigenPodManagerWrapper eigenPodManagerWrapper =
            new EigenPodManagerWrapper(DEPOSIT_CONTRACT, eigenPodBeacon, delegationManager, eigenLayerPauserReg, "v9.9.9");
        address targetAddr = address(eigenPodManagerImplementation);
        cheats.etch(targetAddr, address(eigenPodManagerWrapper).code);

        // 2. create a new staker, operator, and avs
        _configAssetTypes(HOLDS_ETH);
        (staker, strategies, initDepositShares) = _newRandomStaker();
        (operator,,) = _newRandomOperator();
        (avs,) = _newRandomAVS();

        cheats.assume(initDepositShares[0] >= 64 ether);

        EigenPodManagerWrapper(address(eigenPodManager)).setBeaconChainSlashingFactor(address(staker), 1);
        assertEq(eigenPodManager.beaconChainSlashingFactor(address(staker)), 1);

        // 3. start validators and verify withdrawal credentials
        (validators, beaconBalanceGwei,) = staker.startValidators();
        beaconChain.advanceEpoch_NoWithdrawNoRewards();
        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);
    }

    /// @notice Test that a staker can still verify WC, start/complete CP even if the operator has 1 magnitude remaining
    function test_verifyWC_startCompleteCP(uint24 _r) public rand(_r) {
        // 4. start validators and verify withdrawal credentials
        (validators, beaconBalanceGwei,) = staker.startValidators(uint8(_randUint(1, 10)));
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // 5. start/complete checkpoint
        beaconChain.advanceEpoch();
        staker.startCheckpoint();
        beaconChain.advanceEpoch();

        staker.completeCheckpoint();
        check_CompleteCheckpoint_State(staker);

        // 6. Assert BCSF is still 1
        assertEq(eigenPodManager.beaconChainSlashingFactor(address(staker)), 1);
    }

    /// @notice Test that a staker is slashed to 0 BCSF from a minor slash and that they can't deposit more shares
    /// from their EigenPod (either through verifyWC or start/complete CP)
    function test_slashFullyBC_revert_deposit(uint24 _r) public rand(_r) {
        // 4. slash validators on beacon chain (start/complete checkpoint)
        uint40[] memory slashedValidators = _choose(validators);
        slashedGwei = beaconChain.slashValidators(slashedValidators, BeaconChainMock.SlashType.Minor);
        beaconChain.advanceEpoch_NoWithdrawNoRewards();

        staker.startCheckpoint();
        staker.completeCheckpoint();
        check_CompleteCheckpoint_WithCLSlashing_HandleRoundDown_State(staker, slashedGwei);
        // BCSF should be 0 now
        assertEq(eigenPodManager.beaconChainSlashingFactor(address(staker)), 0);

        // 5. deposit expecting revert (randomly pick to verifyWC, start/complete CP)
        if (_randBool()) {
            // Verify WC
            (validators,,) = staker.startValidators(uint8(_randUint(3, 10)));
            beaconChain.advanceEpoch();

            cheats.expectRevert(IDelegationManagerErrors.FullySlashed.selector);
            staker.verifyWithdrawalCredentials(validators);
        } else {
            // Start/complete CP
            // Ensure that not all validators were slashed so that some rewards can be generated when
            // we advance epoch
            cheats.assume(slashedValidators.length < validators.length);
            beaconChain.advanceEpoch();
            staker.startCheckpoint();

            cheats.expectRevert(IDelegationManagerErrors.FullySlashed.selector);
            staker.completeCheckpoint();
        }
    }

    /**
     * @notice Test that a staker with 1 maxMagnitude and 1 BCSF has slashingFactor rounded down to 0
     * and further deposits are not possible
     * Test sequence following _init()
     * 4. Create an operator set and register an operator
     * 5. slash operator to 1 magnitude remaining
     * 6. delegate to operator
     * 7. deposit expecting revert (randomly pick to verifyWC, start/complete CP)
     */
    function test_slashAVS_delegate_revert_startCompleteCP(uint24 _r) public rand(_r) {
        // 4. Create an operator set and register an operator
        operatorSet = avs.createOperatorSet(strategies);
        operator.registerForOperatorSet(operatorSet);
        check_Registration_State_NoAllocation(operator, operatorSet, strategies);
        // Allocate all magnitude to the operator set
        allocateParams = _genAllocation_AllAvailable(operator, operatorSet, strategies);
        operator.modifyAllocations(allocateParams);
        check_IncrAlloc_State_Slashable_NoDelegatedStake(operator, allocateParams);
        _rollBlocksForCompleteAllocation(operator, operatorSet, strategies);

        // 5. slash operator to 1 magnitude remaining
        SlashingParams memory slashParams = _genSlashing_Custom(operator, operatorSet, WAD - 1);
        avs.slashOperator(slashParams);
        check_Base_Slashing_State(operator, allocateParams, slashParams);

        // assert operator has 1 magnitude remaining
        assertEq(allocationManager.getMaxMagnitude(address(operator), beaconChainETHStrategy), 1);

        // 6. delegate to operator
        staker.delegateTo(operator);
        check_Delegation_State(staker, operator, strategies, initDepositShares);

        // 7. deposit expecting revert (randomly pick to verifyWC, start/complete CP)
        if (_randBool()) {
            // Verify WC
            (validators,,) = staker.startValidators(uint8(_randUint(1, 10)));
            beaconChain.advanceEpoch();

            cheats.expectRevert(IDelegationManagerErrors.FullySlashed.selector);
            staker.verifyWithdrawalCredentials(validators);
        } else {
            // Start/complete CP
            beaconChain.advanceEpoch();
            staker.startCheckpoint();

            cheats.expectRevert(IDelegationManagerErrors.FullySlashed.selector);
            staker.completeCheckpoint();
        }
    }
}
