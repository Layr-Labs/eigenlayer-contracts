// SPDX-License-Identifier: MIT
// contract is *modified* from Seaport's test file: https://github.com/ProjectOpenSea/seaport/blob/891b5d4f52b58eb7030597fbb22dca67fd86c4c8/contracts/test/Reenterer.sol
pragma solidity ^0.8.9;

import "forge-std/Test.sol";

contract Reenterer is Test {
    Vm cheats = Vm(VM_ADDRESS);

    address public target;
    uint public msgValue;
    bytes public callData;
    bytes public expectedRevertData;
    bytes public dataToReturn;

    event Reentered(bytes returnData);

    function prepare(address targetToUse, uint msgValueToUse, bytes memory callDataToUse) external {
        target = targetToUse;
        msgValue = msgValueToUse;
        callData = callDataToUse;
    }

    // added function that allows writing to `expectedRevertData`
    function prepare(address targetToUse, uint msgValueToUse, bytes memory callDataToUse, bytes memory expectedRevertDataToUse) external {
        target = targetToUse;
        msgValue = msgValueToUse;
        callData = callDataToUse;
        expectedRevertData = expectedRevertDataToUse;
    }

    // added function that allows writing to `dataToReturn`
    function prepareReturnData(bytes memory returnDataToUse) external {
        dataToReturn = returnDataToUse;
    }

    receive() external payable {
        // added expectrevert logic
        if (expectedRevertData.length != 0) cheats.expectRevert(expectedRevertData);
        (bool success, bytes memory returnData) = target.call{value: msgValue}(callData);

        if (!success) {
            assembly {
                returndatacopy(0, 0, returndatasize())
                revert(0, returndatasize())
            }
        }

        emit Reentered(returnData);

        // added dataToReturn logic
        uint dataToReturnLength = dataToReturn.length;
        if (dataToReturnLength > 0) {
            bytes memory _dataToReturn = dataToReturn;
            assembly {
                return(add(_dataToReturn, 32), dataToReturnLength)
            }
        }
    }

    // added fallback function that is a copy of the `receive` function
    fallback() external payable {
        // added expectRevert logic
        if (expectedRevertData.length != 0) cheats.expectRevert(expectedRevertData);
        (bool success, bytes memory returnData) = target.call{value: msgValue}(callData);

        if (!success) {
            assembly {
                returndatacopy(0, 0, returndatasize())
                revert(0, returndatasize())
            }
        }

        emit Reentered(returnData);

        // added dataToReturn logic
        uint dataToReturnLength = dataToReturn.length;
        if (dataToReturnLength > 0) {
            bytes memory _dataToReturn = dataToReturn;
            assembly {
                return(add(_dataToReturn, 32), dataToReturnLength)
            }
        }
    }
}
