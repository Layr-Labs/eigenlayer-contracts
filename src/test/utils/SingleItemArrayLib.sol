// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/interfaces/IAllocationManager.sol";

/// @dev Helper library for simplifying the syntax for creating single item arrays for inputs.
library SingleItemArrayLib {
    /// -----------------------------------------------------------------------
    /// Native Types
    /// -----------------------------------------------------------------------

    function toArrayU16(
        uint16 x
    ) internal pure returns (uint16[] memory array) {
        array = new uint16[](1);
        array[0] = x;
    }
    
    function toArrayU32(
        uint32 x
    ) internal pure returns (uint32[] memory array) {
        array = new uint32[](1);
        array[0] = x;
    }

    function toArrayU64(
        uint64 x
    ) internal pure returns (uint64[] memory array) {
        array = new uint64[](1);
        array[0] = x;
    }

    function toArrayU256(
        uint256 x
    ) internal pure returns (uint256[] memory array) {
        array = new uint256[](1);
        array[0] = x;
    }


    function toArrayU16(
        uint16 x,
        uint256 len
    ) internal pure returns (uint16[] memory array) {
        array = new uint16[](len);
        for (uint256 i; i < len; ++i) {
            array[i] = x;
        }
    }
    
    function toArrayU32(
        uint32 x,
        uint256 len
    ) internal pure returns (uint32[] memory array) {
        array = new uint32[](len);
        for (uint256 i; i < len; ++i) {
            array[i] = x;
        }
    }

    function toArrayU64(
        uint64 x,
        uint256 len
    ) internal pure returns (uint64[] memory array) {
        array = new uint64[](len);
        for (uint256 i; i < len; ++i) {
            array[i] = x;
        }
    }

    function toArrayU256(
        uint256 x,
        uint256 len
    ) internal pure returns (uint256[] memory array) {
        array = new uint256[](len);
        for (uint256 i; i < len; ++i) {
            array[i] = x;
        }
    }

    function toArray(
        address x
    ) internal pure returns (address[] memory array) {
        array = new address[](1);
        array[0] = x;
    }

    /// -----------------------------------------------------------------------
    /// EigenLayer Types
    /// -----------------------------------------------------------------------

    function toArray(
        IERC20 token
    ) internal pure returns (IERC20[] memory array) {
        array = new IERC20[](1);
        array[0] = token;
    }

    function toArray(
        IStrategy strategy
    ) internal pure returns (IStrategy[] memory array) {
        array = new IStrategy[](1);
        array[0] = strategy;
    }

    function toArray(
        OperatorSet memory operatorSet
    ) internal pure returns (OperatorSet[] memory array) {
        array = new OperatorSet[](1);
        array[0] = operatorSet;
    }

    function toArray(
        IAllocationManagerTypes.CreateSetParams memory createSetParams
    ) internal pure returns (IAllocationManagerTypes.CreateSetParams[] memory array) {
        array = new IAllocationManagerTypes.CreateSetParams[](1);
        array[0] = createSetParams;
    }

    function toArray(
        IAllocationManagerTypes.AllocateParams memory allocateParams
    ) internal pure returns (IAllocationManagerTypes.AllocateParams[] memory array) {
        array = new IAllocationManagerTypes.AllocateParams[](1);
        array[0] = allocateParams;
    }
}