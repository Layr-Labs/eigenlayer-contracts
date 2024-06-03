// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "src/contracts/libraries/BeaconChainProofs.sol";

/// @notice This contract is used to test offchain proof generation
contract BeaconChainProofsWrapper {

    function verifyStateRoot(
        bytes32 beaconBlockRoot,
        BeaconChainProofs.StateRootProof calldata proof
    ) external view {
        BeaconChainProofs.verifyStateRoot(beaconBlockRoot, proof);
    }

    function verifyValidatorFields(
        bytes32 beaconStateRoot,
        bytes32[] calldata validatorFields,
        bytes calldata validatorFieldsProof,
        uint40 validatorIndex
    ) external view {
        BeaconChainProofs.verifyValidatorFields(beaconStateRoot, validatorFields, validatorFieldsProof, validatorIndex);
    }

    function verifyBalanceContainer(
        bytes32 beaconBlockRoot,
        BeaconChainProofs.BalanceContainerProof calldata proof
    ) external view {
        BeaconChainProofs.verifyBalanceContainer(beaconBlockRoot, proof);
    }

    function verifyValidatorBalance(
        bytes32 balanceContainerRoot,
        uint40 validatorIndex,
        BeaconChainProofs.BalanceProof calldata proof
    ) external view {
        BeaconChainProofs.verifyValidatorBalance(balanceContainerRoot, validatorIndex, proof);
    }


}