// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/interfaces/IAllocationManager.sol";
import "src/contracts/interfaces/IDelegationManager.sol";

library ArrayLib {
    using ArrayLib for *;
    using ArrayLib for uint256[];
    using ArrayLib for address[];

    /// -----------------------------------------------------------------------
    /// Single Item Arrays
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

    function toArray(
        bool x
    ) internal pure returns (bool[] memory array) {
        array = new bool[](1);
        array[0] = x;
    }

    function toArray(
        bool x,
        uint256 len
    ) internal pure returns (bool[] memory array) {
        array = new bool[](len);
        for (uint256 i; i < len; ++i) {
            array[i] = x;
        }
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

    function toArray(
        IDelegationManagerTypes.Withdrawal memory withdrawal
    ) internal pure returns (IDelegationManagerTypes.Withdrawal[] memory array) {
        array = new IDelegationManagerTypes.Withdrawal[](1);
        array[0] = withdrawal;
    }

    function toArray(
        IDelegationManagerTypes.QueuedWithdrawalParams memory queuedWithdrawalParams
    ) internal pure returns (IDelegationManagerTypes.QueuedWithdrawalParams[] memory array) {
        array = new IDelegationManagerTypes.QueuedWithdrawalParams[](1);
        array[0] = queuedWithdrawalParams;
    }

    /// -----------------------------------------------------------------------
    /// Length Updates
    /// -----------------------------------------------------------------------

    function setLength(
        uint16[] memory array,
        uint256 len
    ) internal pure returns (uint16[] memory) {
        assembly {
            mstore(array, len)
        }
        return array;
    }

    function setLength(
        uint32[] memory array,
        uint256 len
    ) internal pure returns (uint32[] memory) {
        assembly {
            mstore(array, len)
        }
        return array;
    }

    function setLength(
        uint64[] memory array,
        uint256 len
    ) internal pure returns (uint64[] memory) {
        assembly {
            mstore(array, len)
        }
        return array;
    }

    function setLength(
        uint256[] memory array,
        uint256 len
    ) internal pure returns (uint256[] memory) {
        assembly {
            mstore(array, len)
        }
        return array;
    }

    function setLength(
        address[] memory array,
        uint256 len
    ) internal pure returns (address[] memory) {
        assembly {
            mstore(array, len)
        }
        return array;
    }

    function setLength(
        IERC20[] memory array,
        uint256 len
    ) internal pure returns (IERC20[] memory) {
        assembly {
            mstore(array, len)
        }
        return array;
    }

    function setLength(
        IStrategy[] memory array,
        uint256 len
    ) internal pure returns (IStrategy[] memory) {
        assembly {
            mstore(array, len)
        }
        return array;
    }

    function setLength(
        OperatorSet[] memory array,
        uint256 len
    ) internal pure returns (OperatorSet[] memory) {
        assembly {
            mstore(array, len)
        }
        return array;
    }

    function setLength(
        IDelegationManagerTypes.Withdrawal[] memory array,
        uint256 len
    ) internal pure returns (IDelegationManagerTypes.Withdrawal[] memory) {
        assembly {
            mstore(array, len)
        }
        return array;
    }

    /// -----------------------------------------------------------------------
    /// Contains
    /// -----------------------------------------------------------------------

    function contains(
        uint16[] memory array,
        uint16 x
    ) internal pure returns (bool) {
        for (uint256 i; i < array.length; ++i) {
            if (array[i] == x) return true;
        }
        return false;
    }

    function contains(
        uint32[] memory array,
        uint32 x
    ) internal pure returns (bool) {
        for (uint256 i; i < array.length; ++i) {
            if (array[i] == x) return true;
        }
        return false;
    }

    function contains(
        uint64[] memory array,
        uint64 x
    ) internal pure returns (bool) {
        for (uint256 i; i < array.length; ++i) {
            if (array[i] == x) return true;
        }
        return false;
    }

    function contains(
        uint256[] memory array,
        uint256 x
    ) internal pure returns (bool) {
        for (uint256 i; i < array.length; ++i) {
            if (array[i] == x) return true;
        }
        return false;
    }

    function contains(
        address[] memory array,
        address x
    ) internal pure returns (bool) {
        for (uint256 i; i < array.length; ++i) {
            if (array[i] == x) return true;
        }
        return false;
    }

    function contains(
        IERC20[] memory array,
        IERC20 x
    ) internal pure returns (bool) {
        for (uint256 i; i < array.length; ++i) {
            if (array[i] == x) return true;
        }
        return false;
    }

    function contains(
        IStrategy[] memory array,
        IStrategy x
    ) internal pure returns (bool) {
        for (uint256 i; i < array.length; ++i) {
            if (array[i] == x) return true;
        }
        return false;
    }

    function contains(
        OperatorSet[] memory array,
        OperatorSet memory x
    ) internal pure returns (bool) {
        for (uint256 i; i < array.length; ++i) {
            if (keccak256(abi.encode(array[i])) == keccak256(abi.encode(x))) return true;
        }
        return false;
    }

    function contains(
        IDelegationManagerTypes.Withdrawal[] memory array,
        IDelegationManagerTypes.Withdrawal memory x
    ) internal pure returns (bool) {
        for (uint256 i; i < array.length; ++i) {
            if (keccak256(abi.encode(array[i])) == keccak256(abi.encode(x))) return true;
        }
        return false;
    }

    /// -----------------------------------------------------------------------
    /// indexOf
    /// -----------------------------------------------------------------------

    function indexOf(
        uint16[] memory array,
        uint16 x
    ) internal pure returns (uint256) {
        for (uint256 i; i < array.length; ++i) {
            if (array[i] == x) return i;
        }
        return type(uint256).max;
    }

    function indexOf(
        uint32[] memory array,
        uint32 x
    ) internal pure returns (uint256) {
        for (uint256 i; i < array.length; ++i) {
            if (array[i] == x) return i;
        }
        return type(uint256).max;
    }

    function indexOf(
        uint64[] memory array,
        uint64 x
    ) internal pure returns (uint256) {
        for (uint256 i; i < array.length; ++i) {
            if (array[i] == x) return i;
        }
        return type(uint256).max;
    }

    function indexOf(
        uint256[] memory array,
        uint256 x
    ) internal pure returns (uint256) {
        for (uint256 i; i < array.length; ++i) {
            if (array[i] == x) return i;
        }
        return type(uint256).max;
    }

    function indexOf(
        address[] memory array,
        address x
    ) internal pure returns (uint256) {
        for (uint256 i; i < array.length; ++i) {
            if (array[i] == x) return i;
        }
        return type(uint256).max;
    }

    function indexOf(
        IERC20[] memory array,
        IERC20 x
    ) internal pure returns (uint256) {
        for (uint256 i; i < array.length; ++i) {
            if (array[i] == x) return i;
        }
        return type(uint256).max;
    }

    function indexOf(
        IStrategy[] memory array,
        IStrategy x
    ) internal pure returns (uint256) {
        for (uint256 i; i < array.length; ++i) {
            if (array[i] == x) return i;
        }
        return type(uint256).max;
    }

    function indexOf(
        OperatorSet[] memory array,
        OperatorSet memory x
    ) internal pure returns (uint256) {
        for (uint256 i; i < array.length; ++i) {
            if (keccak256(abi.encode(array[i])) == keccak256(abi.encode(x))) return i;
        }
        return type(uint256).max;
    }

    function indexOf(
        IDelegationManagerTypes.Withdrawal[] memory array,
        IDelegationManagerTypes.Withdrawal memory x
    ) internal pure returns (uint256) {
        for (uint256 i; i < array.length; ++i) {
            if (keccak256(abi.encode(array[i])) == keccak256(abi.encode(x))) return i;
        }
        return type(uint256).max;
    }

    /// -----------------------------------------------------------------------
    /// Sorting
    /// -----------------------------------------------------------------------
    
    function sort(
        IStrategy[] memory array
    ) internal pure returns (IStrategy[] memory) {
        if (array.length <= 1) {
            return array;
        }

        for (uint256 i = 1; i < array.length; i++) {
            IStrategy key = array[i];
            uint256 j = i - 1;
            
            while (j > 0 && uint256(uint160(address(array[j]))) > uint256(uint160(address(key)))) {
                array[j + 1] = array[j];
                j--;
            }
            
            // Special case for the first element
            if (j == 0 && uint256(uint160(address(array[j]))) > uint256(uint160(address(key)))) {
                array[j + 1] = array[j];
                array[j] = key;
            } else if (j < i - 1) {
                array[j + 1] = key;
            }
        }
        
        return array;
    }
}