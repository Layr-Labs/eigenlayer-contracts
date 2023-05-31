// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../interfaces/IRegistryCoordinator.sol";
import "./StakeRegistry.sol";

contract BLSIndexRegistryCoordinator is StakeRegistry, IRegistryCoordinator {
    function registerOperator(bytes memory quorumNumbers, bytes calldata) external returns (bytes32);

}