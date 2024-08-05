import "./EigenPodMethodsAndSimplifications.spec";

ghost mathint validatorsActivated {
    init_state axiom validatorsActivated == 0;
}

hook Sstore _validatorPubkeyHashToInfo[KEY bytes32 validatorPubkeyHash].status IEigenPod.VALIDATOR_STATUS newValue (IEigenPod.VALIDATOR_STATUS oldValue) 
{
    if (oldValue != IEigenPod.VALIDATOR_STATUS.ACTIVE && newValue == IEigenPod.VALIDATOR_STATUS.ACTIVE) validatorsActivated = validatorsActivated + 1;
    if (oldValue == IEigenPod.VALIDATOR_STATUS.ACTIVE && newValue != IEigenPod.VALIDATOR_STATUS.ACTIVE) validatorsActivated = validatorsActivated - 1;
}

rule activeValidatorsCount_correctness(env e, method f) filtered { f -> !f.isView && !isIgnoredMethod(f) }
{
    require validatorsActivated == 0;
    mathint activeValsBefore = activeValidatorCount();
    bytes32 validatorHash;
    calldataarg args;
    f(e, args);
    mathint activeValsAfter = activeValidatorCount();

    assert activeValsAfter - activeValsBefore == validatorsActivated;
}

