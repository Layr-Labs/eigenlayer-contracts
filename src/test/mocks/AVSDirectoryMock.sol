// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IAVSDirectory.sol";

contract AVSDirectoryMock is IAVSDirectory, Test {
    function createOperatorSets(uint32[] calldata operatorSetIds) external {}

    function becomeOperatorSetAVS() external {}

    function migrateOperatorsToOperatorSets(
        address[] calldata operators,
        uint32[][] calldata operatorSetIds
    ) external {}

    function registerOperatorToOperatorSets(
        address operator,
        uint32[] calldata operatorSetIds,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external {}

    function deregisterOperatorFromOperatorSets(address operator, uint32[] calldata operatorSetIds) external {}

    function forceDeregisterFromOperatorSets(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external {}

    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external {}

    function deregisterOperatorFromAVS(address operator) external {}

    function allocate(
        address operator,
        MagnitudeAdjustment[] calldata allocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external {}

    function queueDeallocate(
        address operator,
        MagnitudeAdjustment[] calldata deallocations,
        SignatureWithSaltAndExpiry calldata operatorSignature
    ) external {}

    function completeDeallocations(
        address operator,
        IStrategy[] calldata strategies,
        OperatorSet[][] calldata operatorSets
    ) external {}

    function slashOperator(
        address operator,
        uint32 operatorSetId,
        IStrategy[] calldata strategies,
        uint16 bipsToSlash
    ) external {}

    function updateAVSMetadataURI(string calldata metadataURI) external {}

    function cancelSalt(bytes32 salt) external {}

    function operatorSaltIsSpent(address operator, bytes32 salt) external view returns (bool) {}

    function isMember(address avs, address operator, uint32 operatorSetId) external view returns (bool) {}

    function isOperatorSetAVS(address avs) external view returns (bool) {}

    function isOperatorSet(address avs, uint32 operatorSetId) external view returns (bool) {
        return true;
    }

    function isOperatorSlashable(address operator, OperatorSet memory operatorSet) external view returns (bool) {}

    function getSlashableBips(
        address operator,
        OperatorSet calldata operatorSet,
        IStrategy strategy,
        uint32 timestamp
    ) external view returns (uint16) {}

    function calculateOperatorAVSRegistrationDigestHash(
        address operator,
        address avs,
        bytes32 salt,
        uint256 expiry
    ) external view returns (bytes32) {}

    function calculateOperatorSetRegistrationDigestHash(
        address avs,
        uint32[] calldata operatorSetIds,
        bytes32 salt,
        uint256 expiry
    ) external view returns (bytes32) {}

    function calculateOperatorSetForceDeregistrationTypehash(
        address avs,
        uint32[] calldata operatorSetIds,
        bytes32 salt,
        uint256 expiry
    ) external view returns (bytes32) {}

    /// @notice Getter function for the current EIP-712 domain separator for this contract.
    /// @dev The domain separator will change in the event of a fork that changes the ChainID.
    function domainSeparator() external view returns (bytes32) {}

    /// @notice The EIP-712 typehash for the Registration struct used by the contract.
    function OPERATOR_AVS_REGISTRATION_TYPEHASH() external view returns (bytes32) {}

    /// @notice The EIP-712 typehash for the OperatorSetRegistration struct used by the contract.
    function OPERATOR_SET_REGISTRATION_TYPEHASH() external view returns (bytes32) {}

    function isMember(address operator, IAVSDirectory.OperatorSet memory operatorSet) external view returns (bool) {}
}