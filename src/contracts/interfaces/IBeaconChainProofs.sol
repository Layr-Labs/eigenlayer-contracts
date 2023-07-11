//SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


interface IBeaconChainProofs {

    /**
     * @notice This function verifies merkle proofs of the balance of a certain validator against a beacon chain state root
     * @param validatorIndex the index of the proven validator
     * @param validatorRoot is the serialized balance used to prove the balance of the validator (refer to `getBalanceFromBalanceRoot` above for detailed explanation)
     */
    function verifyValidatorFields(uint40 validatorIndex, bytes32 validatorRoot) external view returns (bool);
}

