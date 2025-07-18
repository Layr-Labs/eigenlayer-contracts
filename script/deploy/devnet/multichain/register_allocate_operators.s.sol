// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

// Core Contracts
import "src/contracts/core/AllocationManager.sol";
import "src/contracts/interfaces/IAllocationManager.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

// Forge
import "forge-std/Script.sol";
import "forge-std/Test.sol";

// forge script script/deploy/devnet/mutlichain/register_allocate_operators.s.sol --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast --sig "run()"
contract RegisterAllocateOperators is Script, Test {
    using stdJson for string;

    // Admin that can perform actions on behalf of the operatorSet
    address superAdmin = 0x8D8A8D3f88f6a6DA2083D865062bFBE3f1cfc293;
    address avs = 0x8D8A8D3f88f6a6DA2083D865062bFBE3f1cfc293;
    uint32 operatorSetId = 50;

    // Contracts
    AllocationManager public allocationManager = AllocationManager(0xFdD5749e11977D60850E06bF5B13221Ad95eb6B4);
    IStrategy public strategy = IStrategy(0xD523267698C81a372191136e477fdebFa33D9FB4); // WETH strategy

    function run() public {
        // Create operators array
        address[] memory operators = new address[](2);
        operators[0] = 0xB37856E7086b999d34000FF458662f5041F7Ad32;
        operators[1] = 0xc9e9994Ba55e3Fc066Edb3F6a9B821e11bEF8446;

        // Create operator set struct
        OperatorSet memory operatorSet = OperatorSet({avs: avs, id: operatorSetId});

        // Prepare operatorSetIds array
        uint32[] memory operatorSetIds = new uint32[](1);
        operatorSetIds[0] = operatorSetId;

        // Register each operator to the operator set
        vm.startBroadcast();
        for (uint256 i = 0; i < operators.length; i++) {
            // Register operator to operator set
            IAllocationManagerTypes.RegisterParams memory registerParams =
                IAllocationManagerTypes.RegisterParams({avs: avs, operatorSetIds: operatorSetIds, data: ""});

            allocationManager.registerForOperatorSets(operators[i], registerParams);
        }

        // Allocate each operator to the operator set
        for (uint256 i = 0; i < operators.length; i++) {
            // Get the operator's max magnitude for the strategy
            uint64 maxMagnitude = allocationManager.getMaxMagnitude(operators[i], strategy);

            // Create strategies array
            IStrategy[] memory strategies = new IStrategy[](1);
            strategies[0] = strategy;

            // Create magnitudes array (allocate everything)
            uint64[] memory magnitudes = new uint64[](1);
            magnitudes[0] = maxMagnitude;

            // Create allocation params
            IAllocationManagerTypes.AllocateParams[] memory allocateParams =
                new IAllocationManagerTypes.AllocateParams[](1);
            allocateParams[0] = IAllocationManagerTypes.AllocateParams({
                operatorSet: operatorSet,
                strategies: strategies,
                newMagnitudes: magnitudes
            });

            // Modify allocations
            allocationManager.modifyAllocations(operators[i], allocateParams);
        }
        vm.stopBroadcast();

        console.log(
            "Successfully registered and allocated", operators.length, "operators to operator set", operatorSetId
        );
    }
}