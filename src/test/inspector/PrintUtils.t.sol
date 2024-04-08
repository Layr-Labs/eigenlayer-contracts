// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "forge-std/Test.sol";

import "@openzeppelin/contracts/utils/Strings.sol";

contract PrintUtils is Test {

    using Strings for *;
    using StdStyle for *;

    string constant HEADER_DELIMITER = "========================================================================";
    string constant SECTION_DELIMITER = "======";

    function _logHeader(string memory name) internal {
        emit log(HEADER_DELIMITER);

        emit log(name);

        emit log(HEADER_DELIMITER);
    }

    function _logHeader(string memory name, address a) internal {
        emit log(HEADER_DELIMITER);

        emit log_named_string(name.cyan(), a.yellow());
        // emit log_named_address(name.cyan(), a);

        emit log(HEADER_DELIMITER);
    }

    function _logSection(string memory name) internal {
        emit log(string.concat(
            SECTION_DELIMITER,
            name,
            SECTION_DELIMITER
        ));
    }

    function _logSection(string memory name, address a) internal {
        emit log(string.concat(
            SECTION_DELIMITER,
            name.cyan(),
            ": ",
            a.yellow().dim(),
            SECTION_DELIMITER
        ));
    }

    function _logAction(string memory name, string memory action) internal {
        emit log_named_string(
            name.cyan(),
            action.italic()
        );
    }

    function _logEth(string memory name, uint value) internal {
        uint lhs;
        uint rhs;
        string memory unit;

        if (value >= 1 ether / 1000) {
            lhs = value / 1 ether;
            rhs = value % 1 ether;
            unit = "ether";
        } else if (value >= 1 gwei) {
            lhs = value / 1 gwei;
            rhs = value % 1 gwei;
            unit = "gwei";
        } else {
            lhs = value;
            unit = "wei";
        }

        string memory amount = string.concat(
            lhs.toString(),
            string("."),
            rhs.toString(),
            string(" "),
            unit
        );

        emit log_named_string(name, amount);
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

    function _log(string memory name, string memory value) internal {
        emit log_named_string(name, value);
    }

    function _log(string memory name, uint value) internal {
        emit log_named_uint(name, value);
    }

    function _log(string memory name, address value) internal {
        emit log_named_string(name, value.yellow());
    }

    function _logDim(string memory name, address value) internal {
        emit log_named_string(name.dim(), value.yellow().dim());
    }

    function _log(string memory name, bytes32 value) internal {
        emit log_named_string(name, value.dimBytes32());
    }

    function _log(string memory name, bool value) internal {
        emit log_named_string(name, value ? "true".green() : "false".magenta());
    }
}