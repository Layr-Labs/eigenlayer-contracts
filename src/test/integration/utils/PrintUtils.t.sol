// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";

import "@openzeppelin/contracts/utils/Strings.sol";

abstract contract PrintUtils is Test {

    using Strings for *;
    using StdStyle for *;

    string constant HEADER_DELIMITER = "==================================================";
    string constant SECTION_DELIMITER = "======";

    /// @dev Inheriting contracts implement this method
    function NAME() public virtual view returns (string memory);

    function _logHeader(string memory key) internal {
        emit log(HEADER_DELIMITER);

        emit log(key);

        emit log(HEADER_DELIMITER);
    }

    function _logHeader(string memory key, address a) internal {
        emit log(HEADER_DELIMITER);

        emit log_named_string(key.cyan(), a.yellow());
        // emit log_named_address(key.cyan(), a);

        emit log(HEADER_DELIMITER);
    }

    function _logSection(string memory key) internal {
        emit log(string.concat(
            SECTION_DELIMITER,
            key,
            SECTION_DELIMITER
        ));
    }

    function _logSection(string memory key, address a) internal {
        emit log(string.concat(
            SECTION_DELIMITER,
            key.cyan(),
            ": ",
            a.yellow().dim(),
            SECTION_DELIMITER
        ));
    }

    function _logAction(string memory key, string memory action) internal {
        emit log_named_string(
            key.cyan(),
            action.italic()
        );
    }

    /// @dev Log method name
    function _logM(string memory method) internal {
        emit log(string.concat(
            NAME().cyan(),
            ".",
            method.italic()
        ));
    }

    function _logM(string memory method, string memory arg) internal {
        emit log(string.concat(
            NAME().cyan(),
            ".",
            method.italic(),
            ":",
            arg
        ));
    }

    function _log(string memory s) internal {
        emit log(s);
    }

    function _logGreen(string memory s) internal {
        emit log(s.green());
    }

    function _logGreen(string memory s, string memory value) internal {
        emit log_named_string(s, value.green());
    }

    function _logYellow(string memory s, string memory value) internal {
        emit log_named_string(s, value.yellow());
    }

    function _log(string memory key, string memory value) internal {
        emit log_named_string(key, value);
    }

    function _log(string memory key, uint value) internal {
        emit log_named_uint(key, value);
    }

    function _log(string memory key, address value) internal {
        emit log_named_string(key, value.yellow());
    }

    function _logDim(string memory key, address value) internal {
        emit log_named_string(key.dim(), value.yellow().dim());
    }

    function _log(string memory key, bytes32 value) internal {
        emit log_named_string(key, value.dimBytes32());
    }

    function _log(string memory key, bool value) internal {
        emit log_named_string(key, value ? "true".green() : "false".magenta());
    }
}