// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/interfaces/IAllocationManager.sol";
import "src/contracts/interfaces/IStrategy.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

type Randomness is uint;

using Random for Randomness global;

library Random {
    /// -----------------------------------------------------------------------
    /// Constants
    /// -----------------------------------------------------------------------

    /// @dev Equivalent to: `uint256(keccak256("RANDOMNESS.SEED"))`.
    uint constant SEED = 0x93bfe7cafd9427243dc4fe8c6e706851eb6696ba8e48960dd74ecc96544938ce;

    /// @dev Equivalent to: `uint256(keccak256("RANDOMNESS.SLOT"))`.
    uint constant SLOT = 0xd0660badbab446a974e6a19901c78a2ad88d7e4f1710b85e1cfc0878477344fd;

    /// -----------------------------------------------------------------------
    /// Helpers
    /// -----------------------------------------------------------------------

    function set(Randomness r) internal returns (Randomness) {
        /// @solidity memory-safe-assembly
        assembly {
            sstore(SLOT, r)
        }
        return r;
    }

    function shuffle(Randomness r) internal returns (Randomness) {
        /// @solidity memory-safe-assembly
        assembly {
            mstore(0x00, sload(SLOT))
            mstore(0x20, r)
            r := keccak256(0x00, 0x20)
        }
        return r.set();
    }

    /// -----------------------------------------------------------------------
    /// Native Types
    /// -----------------------------------------------------------------------

    function Int256(Randomness r, int min, int max) internal returns (int) {
        return max <= min ? min : r.Int256() % (max - min) + min;
    }

    function Int256(Randomness r) internal returns (int) {
        return r.unwrap() % 2 == 0 ? int(r.Uint256()) : -int(r.Uint256());
    }

    function Int128(Randomness r, int128 min, int128 max) internal returns (int128) {
        return int128(Int256(r, min, max));
    }

    function Int128(Randomness r) internal returns (int128) {
        return int128(Int256(r));
    }

    function Int64(Randomness r, int64 min, int64 max) internal returns (int64) {
        return int64(Int256(r, min, max));
    }

    function Int64(Randomness r) internal returns (int64) {
        return int64(Int256(r));
    }

    function Int32(Randomness r, int32 min, int32 max) internal returns (int32) {
        return int32(Int256(r, min, max));
    }

    function Uint256(Randomness r, uint min, uint max) internal returns (uint) {
        return max <= min ? min : r.Uint256() % (max - min) + min;
    }

    function Uint256(Randomness r) internal returns (uint) {
        return r.shuffle().unwrap();
    }

    function Uint128(Randomness r, uint128 min, uint128 max) internal returns (uint128) {
        return uint128(Uint256(r, min, max));
    }

    function Uint128(Randomness r) internal returns (uint128) {
        return uint128(Uint256(r));
    }

    function Uint64(Randomness r, uint64 min, uint64 max) internal returns (uint64) {
        return uint64(Uint256(r, min, max));
    }

    function Uint64(Randomness r) internal returns (uint64) {
        return uint64(Uint256(r));
    }

    function Uint32(Randomness r, uint32 min, uint32 max) internal returns (uint32) {
        return uint32(Uint256(r, min, max));
    }

    function Uint32(Randomness r) internal returns (uint32) {
        return uint32(Uint256(r));
    }

    function Bytes32(Randomness r) internal returns (bytes32) {
        return bytes32(r.Uint256());
    }

    function Address(Randomness r) internal returns (address) {
        return address(uint160(r.Uint256(1, type(uint160).max)));
    }

    function Boolean(Randomness r) internal returns (bool) {
        return r.Uint256() % 2 == 0;
    }

    /// -----------------------------------------------------------------------
    /// Arrays
    /// -----------------------------------------------------------------------

    function Int256Array(Randomness r, uint len, int min, int max) internal returns (int[] memory arr) {
        arr = new int[](len);
        for (uint i; i < len; ++i) {
            arr[i] = r.Int256(min, max);
        }
    }

    function Int128Array(Randomness r, uint len, int128 min, int128 max) internal returns (int128[] memory arr) {
        arr = new int128[](len);
        for (uint i; i < len; ++i) {
            arr[i] = r.Int128(min, max);
        }
    }

    function Int64Array(Randomness r, uint len, int64 min, int64 max) internal returns (int64[] memory arr) {
        arr = new int64[](len);
        for (uint i; i < len; ++i) {
            arr[i] = r.Int64(min, max);
        }
    }

    function Int32Array(Randomness r, uint len, int32 min, int32 max) internal returns (int32[] memory arr) {
        arr = new int32[](len);
        for (uint i; i < len; ++i) {
            arr[i] = r.Int32(min, max);
        }
    }

    function Uint256Array(Randomness r, uint len, uint min, uint max) internal returns (uint[] memory arr) {
        arr = new uint[](len);
        for (uint i; i < len; ++i) {
            arr[i] = uint(r.Uint256(min, max));
        }
    }

    /// -----------------------------------------------------------------------
    /// General Types
    /// -----------------------------------------------------------------------

    function StakerArray(Randomness r, uint len) internal returns (address[] memory stakers) {
        stakers = new address[](len);
        for (uint i; i < len; ++i) {
            stakers[i] = r.Address();
        }
    }

    function StrategyArray(Randomness r, uint len) internal returns (IStrategy[] memory strategies) {
        strategies = new IStrategy[](len);
        for (uint i; i < len; ++i) {
            strategies[i] = IStrategy(r.Address());
        }
    }

    function OperatorSetArray(Randomness r, address avs, uint len) internal returns (OperatorSet[] memory operatorSets) {
        operatorSets = new OperatorSet[](len);
        for (uint i; i < len; ++i) {
            operatorSets[i] = OperatorSet(avs, r.Uint32());
        }
    }

    /// -----------------------------------------------------------------------
    /// `AllocationManager` Types
    /// -----------------------------------------------------------------------

    /// @dev Usage: `r.createSetParams(r, numOpSets, numStrats)`.
    function CreateSetParams(Randomness r, uint numOpSets, uint numStrats)
        internal
        returns (IAllocationManagerTypes.CreateSetParams[] memory params)
    {
        params = new IAllocationManagerTypes.CreateSetParams[](numOpSets);
        for (uint i; i < numOpSets; ++i) {
            params[i].operatorSetId = r.Uint32(1, type(uint32).max);
            params[i].strategies = r.StrategyArray(numStrats);
        }
    }

    /// @dev Usage:
    /// ```
    /// AllocateParams[] memory allocateParams = r.allocateParams(avs, numAllocations, numStrats);
    /// cheats.prank(avs);
    /// allocationManager.createOperatorSets(r.createSetParams(allocateParams));
    /// ```
    function CreateSetParams(Randomness, IAllocationManagerTypes.AllocateParams[] memory allocateParams)
        internal
        pure
        returns (IAllocationManagerTypes.CreateSetParams[] memory params)
    {
        params = new IAllocationManagerTypes.CreateSetParams[](allocateParams.length);
        for (uint i; i < allocateParams.length; ++i) {
            params[i] = IAllocationManagerTypes.CreateSetParams(allocateParams[i].operatorSet.id, allocateParams[i].strategies);
        }
    }

    /// @dev Usage:
    /// ```
    /// AllocateParams[] memory allocateParams = r.allocateParams(avs, numAllocations, numStrats);
    /// CreateSetParams[] memory createSetParams = r.createSetParams(allocateParams);
    ///
    /// cheats.prank(avs);
    /// allocationManager.createOperatorSets(createSetParams);
    ///
    /// cheats.prank(operator);
    /// allocationManager.modifyAllocations(allocateParams);
    /// ```
    function AllocateParams(Randomness r, address avs, uint numAllocations, uint numStrats)
        internal
        returns (IAllocationManagerTypes.AllocateParams[] memory allocateParams)
    {
        allocateParams = new IAllocationManagerTypes.AllocateParams[](numAllocations);

        // TODO: Randomize magnitudes such that they sum to 1e18 (100%).
        uint64 magnitudePerSet = uint64(WAD / numStrats);

        for (uint i; i < numAllocations; ++i) {
            allocateParams[i].operatorSet = OperatorSet(avs, r.Uint32());
            allocateParams[i].strategies = r.StrategyArray(numStrats);
            allocateParams[i].newMagnitudes = new uint64[](numStrats);

            for (uint j; j < numStrats; ++j) {
                allocateParams[i].newMagnitudes[j] = magnitudePerSet;
            }
        }
    }

    /// @dev Usage:
    /// ```
    /// AllocateParams[] memory allocateParams = r.allocateParams(avs, numAllocations, numStrats);
    /// AllocateParams[] memory deallocateParams = r.deallocateParams(allocateParams);
    /// CreateSetParams[] memory createSetParams = r.createSetParams(allocateParams);
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
    function DeallocateParams(Randomness r, IAllocationManagerTypes.AllocateParams[] memory allocateParams)
        internal
        returns (IAllocationManagerTypes.AllocateParams[] memory deallocateParams)
    {
        uint numDeallocations = allocateParams.length;

        deallocateParams = new IAllocationManagerTypes.AllocateParams[](numDeallocations);

        for (uint i; i < numDeallocations; ++i) {
            deallocateParams[i].operatorSet = allocateParams[i].operatorSet;
            deallocateParams[i].strategies = allocateParams[i].strategies;

            deallocateParams[i].newMagnitudes = new uint64[](allocateParams[i].strategies.length);
            for (uint j; j < allocateParams[i].strategies.length; ++j) {
                deallocateParams[i].newMagnitudes[j] = r.Uint64(0, allocateParams[i].newMagnitudes[j] - 1);
            }
        }
    }

    /// @dev Usage:
    /// ```
    /// AllocateParams[] memory allocateParams = r.allocateParams(avs, numAllocations, numStrats);
    /// CreateSetParams[] memory createSetParams = r.createSetParams(allocateParams);
    /// RegisterParams memory registerParams = r.registerParams(allocateParams);
    ///
    /// cheats.prank(avs);
    /// allocationManager.createOperatorSets(createSetParams);
    ///
    /// cheats.prank(operator);
    /// allocationmanager.registerForOperatorSets(registerParams);
    /// ```
    function RegisterParams(Randomness r, IAllocationManagerTypes.AllocateParams[] memory allocateParams)
        internal
        returns (IAllocationManagerTypes.RegisterParams memory params)
    {
        params.avs = allocateParams[0].operatorSet.avs;
        params.operatorSetIds = new uint32[](allocateParams.length);
        for (uint i; i < allocateParams.length; ++i) {
            params.operatorSetIds[i] = allocateParams[i].operatorSet.id;
        }
        params.data = abi.encode(r.Bytes32());
    }

    function SlashingParams(Randomness r, address operator, IAllocationManagerTypes.AllocateParams memory allocateParams)
        internal
        returns (IAllocationManagerTypes.SlashingParams memory params)
    {
        IStrategy[] memory strategies = allocateParams.strategies;

        params.operator = operator;
        params.operatorSetId = allocateParams.operatorSet.id;

        // Randomly select a subset of strategies to slash.
        uint len = r.Uint256({min: 1, max: strategies.length});

        // Update length of strategies array.
        assembly {
            mstore(strategies, len)
        }

        params.strategies = strategies;
        params.wadsToSlash = new uint[](len);

        // Randomly select a `wadToSlash` for each strategy.
        for (uint i; i < len; ++i) {
            params.wadsToSlash[i] = r.Uint256({min: 0.001 ether, max: 1 ether});
        }
    }

    /// -----------------------------------------------------------------------
    /// Helpers
    /// -----------------------------------------------------------------------

    function wrap(uint r) internal pure returns (Randomness) {
        return Randomness.wrap(r);
    }

    function unwrap(Randomness r) internal pure returns (uint) {
        return Randomness.unwrap(r);
    }
}
