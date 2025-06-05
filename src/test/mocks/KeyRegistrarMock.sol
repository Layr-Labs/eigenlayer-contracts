// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/interfaces/IKeyRegistrar.sol";
import "src/contracts/libraries/OperatorSetLib.sol";
import "src/contracts/interfaces/ISignatureUtilsMixin.sol";
import "src/contracts/libraries/BN254.sol";

contract KeyRegistrarMock is IKeyRegistrar {
    using OperatorSetLib for OperatorSet;

    mapping(bytes32 => IKeyRegistrar.CurveType) internal _operatorSetCurveTypes;

    function getOperatorSetCurveType(OperatorSet memory operatorSet) external view returns (IKeyRegistrar.CurveType) {
        return _operatorSetCurveTypes[operatorSet.key()];
    }

    function setOperatorSetCurveType(OperatorSet memory operatorSet, IKeyRegistrar.CurveType curveType) external {
        _operatorSetCurveTypes[operatorSet.key()] = curveType;
    }

    // Stub implementations for remaining interface methods
    function checkKey(OperatorSet memory, address) external pure returns (bool) {
        return true;
    }

    function configureOperatorSet(OperatorSet memory, CurveType) external {}

    function deregisterKey(address, OperatorSet memory) external {}

    function getBN254Key(OperatorSet memory, address) external pure returns (BN254.G1Point memory, BN254.G2Point memory) {
        return (BN254.G1Point(0, 0), BN254.G2Point([uint(0), uint(0)], [uint(0), uint(0)]));
    }

    function getECDSAAddress(OperatorSet memory, address) external pure returns (address) {
        return address(0);
    }

    function getECDSAKey(OperatorSet memory, address) external pure returns (bytes memory) {
        return "";
    }

    function getKeyHash(OperatorSet memory, address) external pure returns (bytes32) {
        return bytes32(0);
    }

    function initialize(address) external {}

    function isKeyGloballyRegistered(bytes32) external pure returns (bool) {
        return false;
    }

    function isRegistered(OperatorSet memory, address) external pure returns (bool) {
        return true;
    }

    function registerKey(address, OperatorSet memory, bytes calldata, bytes calldata) external {}

    function version() external pure returns (string memory) {
        return "1.0.0";
    }
}
