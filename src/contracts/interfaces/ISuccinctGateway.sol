// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

interface ISuccinctGatewayEvents {
    event RequestCallback(
        uint32 indexed nonce,
        bytes32 indexed functionId,
        bytes input,
        bytes context,
        address callbackAddress,
        bytes4 callbackSelector,
        uint32 callbackGasLimit,
        uint256 feeAmount
    );
    event RequestCall(
        bytes32 indexed functionId,
        bytes input,
        address entryAddress,
        bytes entryCalldata,
        uint32 entryGasLimit,
        address sender,
        uint256 feeAmount
    );
    event RequestFulfilled(
        uint32 indexed nonce, bytes32 indexed functionId, bytes32 inputHash, bytes32 outputHash
    );
    event Call(bytes32 indexed functionId, bytes32 inputHash, bytes32 outputHash);
    event SetFeeVault(address indexed oldFeeVault, address indexed newFeeVault);
    event ProverUpdated(address indexed prover, bool added);
}

interface ISuccinctGatewayErrors {
    error InvalidRequest(uint32 nonce, bytes32 expectedRequestHash, bytes32 requestHash);
    error CallbackFailed(bytes4 callbackSelector, bytes output, bytes context);
    error InvalidCall(bytes32 functionId, bytes input);
    error CallFailed(address callbackAddress, bytes callbackData);
    error InvalidProof(address verifier, bytes32 inputHash, bytes32 outputHash, bytes proof);
    error ReentrantFulfill();
    error OnlyProver(address sender);
}

interface ISuccinctGateway is ISuccinctGatewayEvents, ISuccinctGatewayErrors {
    function requestCallback(
        bytes32 functionId,
        bytes memory input,
        bytes memory context,
        bytes4 callbackSelector,
        uint32 callbackGasLimit
    ) external payable returns (bytes32);

    function requestCall(
        bytes32 functionId,
        bytes memory input,
        address entryAddress,
        bytes memory entryData,
        uint32 entryGasLimit
    ) external payable;

    function verifiedCall(bytes32 functionId, bytes memory input)
        external
        returns (bytes memory);

    function isCallback() external view returns (bool);
}