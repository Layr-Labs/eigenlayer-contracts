// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/interfaces/IBLSStakeRegistryCoordinator.sol";
import "../../contracts/interfaces/IBLSPubkeyRegistry.sol";
import "../../contracts/interfaces/IStakeRegistry.sol";



contract RegistryCoordinatorMock is IBLSStakeRegistryCoordinator {

    // Add the necessary variables here
    IStakeRegistry public stakeRegistry;
    IBLSPubkeyRegistry public blsPubkeyRegistry;

    
    constructor(IStakeRegistry _stakeRegistry, IBLSPubkeyRegistry _blsPubkeyRegistry) {
        // Initialize your variables here
        stakeRegistry = _stakeRegistry;
        blsPubkeyRegistry = _blsPubkeyRegistry;
    }

    function getOperatorId(address operator) override external view returns (bytes32) {
        // Implement your logic here
    }

    function getFromTaskNumberForOperator(address operator) override external view returns (uint32) {
        // Implement your logic here
    }

    function operatorIdToQuorumBitmap(bytes32 operatorId) override external view returns (uint256) {
        // Implement your logic here
    }

    function registries(uint256 index) override external view returns (address) {
        // Implement your logic here
    }

    function numRegistries() override external view returns (uint256) {
        // Implement your logic here
    }

    function registerOperator(bytes memory quorumNumbers, bytes calldata interactionData) override external {
        // Implement your logic here
    }

    function deregisterOperator(bytes calldata interactionData) override external {
        // Implement your logic here
    }
}
