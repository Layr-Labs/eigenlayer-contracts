methods
{
    function verifyValidatorFields(bytes32, bytes32[], bytes, uint40) external envfree;
    function verifyValidatorBalance(bytes32, uint40, BeaconChainProofs.BalanceProof proof) external envfree;
}

/// @title For a specific validatorIndex group and a specific balanceContainerRoot, 
///    there is only one proof.balanceRoot such that verifyValidatorBalance doesn't revert
//validatorIndex0 == validatorIndex1 && balanceContainerRoot0 == balanceContainerRoot1 
//  && verifyValidatorBalance(balanceContainerRoot0, validatorIndex0, proof0) didnt revert 
//  && verifyValidatorBalance(balanceContainerRoot1, validatorIndex1, proof1) didnt revert, 
//  then proof0.balanceRoot == proof1.balanceRoot
rule verifyValidatorBalance_balanceRootUnique()
{
    bytes32 balanceContainerRoot1; bytes32 balanceContainerRoot2;
    uint40 validatorIndex1; uint40 validatorIndex2;
    BeaconChainProofs.BalanceProof proof1; BeaconChainProofs.BalanceProof proof2;

    verifyValidatorBalance(balanceContainerRoot1, validatorIndex1, proof1);
    verifyValidatorBalance(balanceContainerRoot2, validatorIndex2, proof2);

    assert (validatorIndex1 / 4 == validatorIndex2 / 4 && 
        balanceContainerRoot1 == balanceContainerRoot2) =>
        proof1.balanceRoot == proof2.balanceRoot;
}

// TODO the same for verifyValidatorFields
