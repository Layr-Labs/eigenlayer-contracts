// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Vm.sol";

import "src/contracts/interfaces/IAllocationManager.sol";
import "src/contracts/interfaces/IStrategy.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

type Randomness is uint256;

library Random {
    VmSafe constant cheats = VmSafe(address(bytes20(uint160(uint256(keccak256("hevm cheat code"))))));

    /// -----------------------------------------------------------------------
    /// Native Types
    /// -----------------------------------------------------------------------

    function Address() internal returns (address) {
        return cheats.randomAddress();
    }

    function Bytes32() internal returns (bytes32) {
        return bytes32(cheats.randomUint());
    }

    function Boolean() internal returns (bool) {
        return cheats.randomUint(0, 1) == 0;
    }

    /// -----------------------------------------------------------------------
    /// Unsigned Integers
    /// -----------------------------------------------------------------------

    function Uint256(uint256 min, uint256 max) internal returns (uint256) {
        return cheats.randomUint(min, max);
    }

    function Uint256() internal returns (uint256) {
        return cheats.randomUint();
    }

    function Uint128(uint128 min, uint128 max) internal returns (uint128) {
        return uint128(cheats.randomUint(min, max));
    }

    function Uint128() internal view returns (uint128) {
        return uint128(cheats.randomUint({bits: 128}));
    }

    function Uint64(uint64 min, uint64 max) internal returns (uint64) {
        return uint64(cheats.randomUint(min, max));
    }

    function Uint64() internal view returns (uint64) {
        return uint64(cheats.randomUint({bits: 64}));
    }

    function Uint32(uint32 min, uint32 max) internal returns (uint32) {
        return uint32(cheats.randomUint(min, max));
    }

    function Uint32() internal view returns (uint32) {
        return uint32(cheats.randomUint({bits: 32}));
    }

    /// -----------------------------------------------------------------------
    /// Signed Integers
    /// -----------------------------------------------------------------------

    function bound(int256 r, int256 min, int256 max) public pure returns (int256) {
        require(min <= max, "Invalid range");

        int256 range = max - min + 1;

        if (r >= 0) {
            r = min + int256(uint256(r) % uint256(range));
        } else {
            r = min + int256(uint256(-r) % uint256(range));
        }

        return r;
    }

    function Int256(int256 min, int256 max) internal view returns (int256) {
        return bound(cheats.randomInt(), min, max);
    }

    function Int256() internal view returns (int256) {
        return cheats.randomInt();
    }

    function Int128(int128 min, int128 max) internal view returns (int128) {
        return int128(Int256(min, max));
    }

    function Int128() internal view returns (int128) {
        return int128(cheats.randomInt({bits: 128}));
    }

    function Int64(int64 min, int64 max) internal view returns (int64) {
        return int64(Int256(min, max));
    }

    function Int64() internal view returns (int64) {
        return int64(cheats.randomInt({bits: 64}));
    }

    function Int32(int32 min, int32 max) internal view returns (int32) {
        return int32(Int256(min, max));
    }

    function Int32() internal view returns (int32) {
        return int32(cheats.randomInt({bits: 32}));
    }

    /// -----------------------------------------------------------------------
    /// General Types
    /// -----------------------------------------------------------------------

    function StrategyArray(
        uint256 len
    ) internal returns (IStrategy[] memory strategies) {
        strategies = new IStrategy[](len);
        for (uint256 i; i < len; ++i) {
            strategies[i] = IStrategy(Address());
        }
    }

    function OperatorSetArray(address avs, uint256 len) internal view returns (OperatorSet[] memory operatorSets) {
        operatorSets = new OperatorSet[](len);
        for (uint256 i; i < len; ++i) {
            operatorSets[i] = OperatorSet(avs, Uint32());
        }
    }

    /// -----------------------------------------------------------------------
    /// `AllocationManager` Types
    /// -----------------------------------------------------------------------

    /// @dev Usage: `createSetParams(r, numOpSets, numStrats)`.
    function CreateSetParams(
        uint256 numOpSets,
        uint256 numStrats
    ) internal returns (IAllocationManagerTypes.CreateSetParams[] memory params) {
        params = new IAllocationManagerTypes.CreateSetParams[](numOpSets);
        for (uint256 i; i < numOpSets; ++i) {
            params[i].operatorSetId = Uint32(1, type(uint32).max);
            params[i].strategies = StrategyArray(numStrats);
        }
    }

    /// @dev Usage:
    /// ```
    /// AllocateParams[] memory allocateParams = allocateParams(avs, numAllocations, numStrats);
    /// cheats.prank(avs);
    /// allocationManager.createOperatorSets(createSetParams(allocateParams));
    /// ```
    function CreateSetParams(
        IAllocationManagerTypes.AllocateParams[] memory allocateParams
    ) internal pure returns (IAllocationManagerTypes.CreateSetParams[] memory params) {
        params = new IAllocationManagerTypes.CreateSetParams[](allocateParams.length);
        for (uint256 i; i < allocateParams.length; ++i) {
            params[i] =
                IAllocationManagerTypes.CreateSetParams(allocateParams[i].operatorSet.id, allocateParams[i].strategies);
        }
    }

    /// @dev Usage:
    /// ```
    /// AllocateParams[] memory allocateParams = allocateParams(avs, numAllocations, numStrats);
    /// CreateSetParams[] memory createSetParams = createSetParams(allocateParams);
    ///
    /// cheats.prank(avs);
    /// allocationManager.createOperatorSets(createSetParams);
    ///
    /// cheats.prank(operator);
    /// allocationManager.modifyAllocations(allocateParams);
    /// ```
    function AllocateParams(
        address avs,
        uint256 numAllocations,
        uint256 numStrats
    ) internal returns (IAllocationManagerTypes.AllocateParams[] memory allocateParams) {
        allocateParams = new IAllocationManagerTypes.AllocateParams[](numAllocations);

        for (uint256 i; i < numAllocations; ++i) {
            allocateParams[i].operatorSet = OperatorSet(avs, Uint32());
            allocateParams[i].strategies = StrategyArray(numStrats);
            allocateParams[i].newMagnitudes = new uint64[](numStrats);

            uint64 remaining = WAD;
            uint64 baseMagnitude = uint64(WAD / numStrats); // Even split
            uint64 adjustment = uint64(WAD % numStrats); // Remainder for adjustment

            for (uint256 j; j < numStrats; ++j) {
                allocateParams[i].newMagnitudes[j] = baseMagnitude;

                // Distribute the adjustment remainder incrementally
                if (adjustment > 0) {
                    allocateParams[i].newMagnitudes[j]++;
                    adjustment--;
                }

                remaining -= allocateParams[i].newMagnitudes[j];
            }

            // Ensure the remaining is fully distributed and total matches WAD
            assert(remaining == 0);
        }
    }

    /// @dev Usage:
    /// ```
    /// AllocateParams[] memory allocateParams = allocateParams(avs, numAllocations, numStrats);
    /// AllocateParams[] memory deallocateParams = deallocateParams(allocateParams);
    /// CreateSetParams[] memory createSetParams = createSetParams(allocateParams);
    ///
    /// cheats.prank(avs);
    /// allocationManager.createOperatorSets(createSetParams);
    ///
    /// cheats.prank(operator);
    /// allocationManager.modifyAllocations(allocateParams);
    ///
    /// cheats.prank(operator)
    /// allocationManager.modifyAllocations(deallocateParams);
    /// ```
    function DeallocateParams(
        IAllocationManagerTypes.AllocateParams[] memory allocateParams
    ) internal returns (IAllocationManagerTypes.AllocateParams[] memory deallocateParams) {
        uint256 numDeallocations = allocateParams.length;

        deallocateParams = new IAllocationManagerTypes.AllocateParams[](numDeallocations);

        for (uint256 i; i < numDeallocations; ++i) {
            deallocateParams[i].operatorSet = allocateParams[i].operatorSet;
            deallocateParams[i].strategies = allocateParams[i].strategies;

            deallocateParams[i].newMagnitudes = new uint64[](allocateParams[i].strategies.length);
            for (uint256 j; j < allocateParams[i].strategies.length; ++j) {
                deallocateParams[i].newMagnitudes[j] = Uint64(0, allocateParams[i].newMagnitudes[j] - 1);
            }
        }
    }

    /// @dev Usage:
    /// ```
    /// AllocateParams[] memory allocateParams = allocateParams(avs, numAllocations, numStrats);
    /// CreateSetParams[] memory createSetParams = createSetParams(allocateParams);
    /// RegisterParams memory registerParams = registerParams(allocateParams);
    ///
    /// cheats.prank(avs);
    /// allocationManager.createOperatorSets(createSetParams);
    ///
    /// cheats.prank(operator);
    /// allocationmanager.registerForOperatorSets(registerParams);
    /// ```
    function RegisterParams(
        IAllocationManagerTypes.AllocateParams[] memory allocateParams
    ) internal returns (IAllocationManagerTypes.RegisterParams memory params) {
        params.avs = allocateParams[0].operatorSet.avs;
        params.operatorSetIds = new uint32[](allocateParams.length);
        for (uint256 i; i < allocateParams.length; ++i) {
            params.operatorSetIds[i] = allocateParams[i].operatorSet.id;
        }
        params.data = abi.encode(Bytes32());
    }

    function SlashingParams(
        address operator,
        IAllocationManagerTypes.AllocateParams memory allocateParams
    ) internal returns (IAllocationManagerTypes.SlashingParams memory params) {
        params.operator = operator;
        params.operatorSetId = allocateParams.operatorSet.id;
        params.wadToSlash = Uint256(0.01 ether, 1 ether);
        params.description = "test";
    }
}
