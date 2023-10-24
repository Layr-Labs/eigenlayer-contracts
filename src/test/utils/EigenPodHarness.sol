// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/pods/EigenPod.sol";

contract EPInternalFunctions is EigenPod {

    constructor(
        IETHPOSDeposit _ethPOS,
        IDelayedWithdrawalRouter _delayedWithdrawalRouter,
        IEigenPodManager _eigenPodManager,
        uint64 _MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
        uint64 _RESTAKED_BALANCE_OFFSET_GWEI,
        uint64 _GENESIS_TIME
    ) EigenPod(
        _ethPOS,
        _delayedWithdrawalRouter,
        _eigenPodManager,
        _MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR,
        _RESTAKED_BALANCE_OFFSET_GWEI,
        _GENESIS_TIME
    ) {}

    function processFullWithdrawal(
        uint40 validatorIndex,
        bytes32 validatorPubkeyHash,
        uint64 withdrawalHappenedTimestamp,
        address recipient,
        uint64 withdrawalAmountGwei,
        ValidatorInfo memory validatorInfo
    ) public returns(IEigenPod.VerifiedWithdrawal memory) {
        return _processFullWithdrawal(
            validatorIndex,
            validatorPubkeyHash,
            withdrawalHappenedTimestamp,
            recipient,
            withdrawalAmountGwei,
            validatorInfo
        );
    }

    function processPartialWithdrawal(
        uint40 validatorIndex,
        uint64 withdrawalHappenedTimestamp,
        address recipient,
        uint64 withdrawalAmountGwei
    ) public returns(IEigenPod.VerifiedWithdrawal memory) {
        return _processPartialWithdrawal(
            validatorIndex,
            withdrawalHappenedTimestamp,
            recipient,
            withdrawalAmountGwei
        );
    }

    function verifyBalanceUpdate(
        uint64 oracleTimestamp,
        uint40 validatorIndex,
        bytes32 beaconStateRoot,
        BeaconChainProofs.BalanceUpdateProof calldata balanceUpdateProof,
        bytes32[] calldata validatorFields,
        uint64 mostRecentBalanceUpdateTimestamp
    )
        public
    {
        bytes32 pkhash = validatorFields[0];
        _validatorPubkeyHashToInfo[pkhash].mostRecentBalanceUpdateTimestamp = mostRecentBalanceUpdateTimestamp;
        _verifyBalanceUpdate(
            oracleTimestamp,
            validatorIndex,
            beaconStateRoot,
            balanceUpdateProof,
            validatorFields
        );
    }

    function setValidatorStatus(bytes32 pkhash, VALIDATOR_STATUS status) public {
        _validatorPubkeyHashToInfo[pkhash].status = status;
    }
 }