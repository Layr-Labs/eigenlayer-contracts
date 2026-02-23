// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/interfaces/IAVSRegistrar.sol";

/// @notice Minimal AVS registrar for E2E testing.
/// @dev AllocationManager will call `registerOperator`/`deregisterOperator` during operator set
/// registration/deregistration. This registrar simply records membership and emits events.
contract MockAVSRegistrar is IAVSRegistrar {
    error UnsupportedAVS();

    event OperatorRegistered(address indexed operator, address indexed avs, uint32[] operatorSetIds, bytes data);
    event OperatorDeregistered(address indexed operator, address indexed avs, uint32[] operatorSetIds);

    address public immutable supportedAVS;

    /// operator => operatorSetId => registered
    mapping(address => mapping(uint32 => bool)) public isRegistered;

    constructor(
        address _supportedAVS
    ) {
        supportedAVS = _supportedAVS;
    }

    function supportsAVS(
        address avs
    ) external view returns (bool) {
        return avs == supportedAVS;
    }

    function registerOperator(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds,
        bytes calldata data
    ) external {
        if (avs != supportedAVS) revert UnsupportedAVS();
        for (uint256 i = 0; i < operatorSetIds.length; i++) {
            isRegistered[operator][operatorSetIds[i]] = true;
        }
        emit OperatorRegistered(operator, avs, operatorSetIds, data);
    }

    function deregisterOperator(
        address operator,
        address avs,
        uint32[] calldata operatorSetIds
    ) external {
        if (avs != supportedAVS) revert UnsupportedAVS();
        for (uint256 i = 0; i < operatorSetIds.length; i++) {
            isRegistered[operator][operatorSetIds[i]] = false;
        }
        emit OperatorDeregistered(operator, avs, operatorSetIds);
    }
}

