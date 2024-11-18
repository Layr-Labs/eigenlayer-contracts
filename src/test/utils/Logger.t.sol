// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

abstract contract Logger is Test {
    using StdStyle for *;

    /// -----------------------------------------------------------------------
    /// Must Override
    /// -----------------------------------------------------------------------

    /// @dev Provide a name for the inheriting contract.
    function NAME() public view virtual returns (string memory);

    /// -----------------------------------------------------------------------
    /// Storage
    /// -----------------------------------------------------------------------

    bool on = true;

    /// -----------------------------------------------------------------------
    /// Colored Names
    /// -----------------------------------------------------------------------

    /// @dev Returns `NAME` colored based on the inheriting contract's role.
    function NAME_COLORED() internal view returns (string memory) {
        vm.pauseTracing();
        bool isOperator = _contains(NAME(), "operator");
        bool isStaker = _contains(NAME(), "staker");
        vm.resumeTracing();

        if (isOperator) {
            return NAME().blue();
        } else if (isStaker) {
            return NAME().magenta();
        } else {
            return NAME().yellow();
        }
    }

    /// -----------------------------------------------------------------------
    /// Logging
    /// -----------------------------------------------------------------------

    modifier noTracing() {
        vm.pauseTracing();
        _;
        vm.resumeTracing();
    }

    modifier noLogging() {
        _pauseLogging();
        _;
        _resumeLogging();
    }

    function _pauseLogging() internal {
        console.log("\n%s logging paused...", NAME_COLORED());
        on = false;
    }

    function _resumeLogging() internal {
        console.log("\n%s logging unpaused...", NAME_COLORED());
        on = true;
    }

    /// -----------------------------------------------------------------------
    /// Logging
    /// -----------------------------------------------------------------------

    function _logM(
        string memory method
    ) internal view {
        console.log("\n%s.%s()", NAME_COLORED(), method.italic());
    }

    function _logM(string memory method, string memory arg) internal view {
        console.log("\n%s.%s(%s)", NAME_COLORED(), method.italic(), arg.dim());
    }

    /// -----------------------------------------------------------------------
    /// Helpers
    /// -----------------------------------------------------------------------

    /// @dev Returns `true` if `needle` is found in `haystack`.
    function _contains(string memory haystack, string memory needle) internal pure returns (bool) {
        return vm.indexOf(haystack, needle) != type(uint256).max;
    }

    /// @dev Returns a string representation of a `uint256` in wad format (with decimal place).
    function _toStringWad(
        uint256 wad
    ) public pure returns (string memory s) {
        if (wad == 0) return "0.0 (wad)";

        s = vm.toString(wad);

        while (bytes(s).length < 18) s = string.concat("0", s);

        uint256 len = bytes(s).length;
        bytes memory b = bytes(s);
        bytes memory left = new bytes(len > 18 ? len - 18 : 0);
        bytes memory right = new bytes(18);

        for (uint256 i; i < left.length; ++i) {
            left[i] = b[i];
        }

        for (uint256 i; i < 18; ++i) {
            right[i] = b[len - 18 + i];
        }

        return string.concat(left.length > 0 ? string(left) : "0", ".", string(right), " (wad)");
    }
}
