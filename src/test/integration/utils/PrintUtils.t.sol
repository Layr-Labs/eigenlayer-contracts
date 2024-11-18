// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

abstract contract PrintUtils is Test {
    using StdStyle for *;

    /// @dev Inheriting contracts implement this method
    function NAME() public view virtual returns (string memory);

    function NAME_COLORED() internal view returns (string memory) {
        bool isOperator = vm.indexOf(NAME(), "operator") != type(uint256).max;
        bool isStaker = vm.indexOf(NAME(), "staker") != type(uint256).max;

        if (isOperator) {
            return NAME().blue();
        } else if (isStaker){
            return NAME().magenta();
        } else {
            return NAME().yellow();
        }
    }

    function _logM(
        string memory method
    ) internal {
        console.log("\n%s.%s()", NAME_COLORED(), method.italic());
    }

    function _logM(string memory method, string memory arg) internal {
        console.log("\n%s.%s(%s)", NAME_COLORED(), method.italic(), arg.dim());
    }
    
    // function console.log(string memory key, bool value) internal {
    //     emit log_named_string(key, value ? "true".green() : "false".magenta());
    // }
}
