// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "src/contracts/interfaces/IStrategy.sol";
import {IAllocationManagerTypes} from "src/contracts/interfaces/IAllocationManager.sol";

Vm constant cheats = Vm(address(uint160(uint(keccak256("hevm cheat code")))));

IStrategy constant BEACONCHAIN_ETH_STRAT = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);
IERC20 constant NATIVE_ETH = IERC20(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

uint constant MIN_BALANCE = 1e6;
uint constant MAX_BALANCE = 5e6;
uint constant GWEI_TO_WEI = 1e9;

uint constant FLAG = 1;

/// @dev Types representing the different types of assets a ranomized users can hold.
uint constant NO_ASSETS = (FLAG << 0); // will have no assets
uint constant HOLDS_LST = (FLAG << 1); // will hold some random amount of LSTs
uint constant HOLDS_ETH = (FLAG << 2); // will hold some random amount of ETH
uint constant HOLDS_ALL = (FLAG << 3); // will always hold ETH, and some LSTs
uint constant HOLDS_MAX = (FLAG << 4); // will hold every LST and ETH (used for testing max strategies)

/// @dev Types representing the different types of users that can be created.
uint constant DEFAULT = (FLAG << 0);
uint constant ALT_METHODS = (FLAG << 1);

/// @dev Types representing the different types of forks that can be simulated.
uint constant LOCAL = (FLAG << 0);
uint constant MAINNET = (FLAG << 1);
uint constant HOLESKY = (FLAG << 2);

abstract contract Logger is Test {
    using StdStyle for *;

    /// -----------------------------------------------------------------------
    /// Storage
    /// -----------------------------------------------------------------------

    bool public logging = true;

    /// -----------------------------------------------------------------------
    /// Modifiers
    /// -----------------------------------------------------------------------

    // Address used to store a trace counter to allow us to use noTracing
    // across any contract that inherits Logger
    address constant LOG_STATE_ADDR = address(0xDEADBEEF);
    bytes32 constant LOG_STATE_SLOT = bytes32(0);

    modifier noTracing() {
        uint traceCounter = _getTraceCounter();
        if (traceCounter == 0) cheats.pauseTracing();

        traceCounter++;
        _setTraceCounter(traceCounter);

        _;

        traceCounter = _getTraceCounter();
        traceCounter--;
        _setTraceCounter(traceCounter);

        if (traceCounter == 0) cheats.resumeTracing();
    }

    modifier noLogging() {
        logging = false;
        _;
        logging = true;
    }

    /// -----------------------------------------------------------------------
    /// Must Override
    /// -----------------------------------------------------------------------

    /// @dev Provide a name for the inheriting contract.
    function NAME() public view virtual returns (string memory);

    /// -----------------------------------------------------------------------
    /// Colored Names
    /// -----------------------------------------------------------------------

    /// @dev Returns `NAME` colored based logging the inheriting contract's role.
    function NAME_COLORED() public view returns (string memory) {
        return colorByRole(NAME());
    }

    /// @dev Returns `name` colored based logging its role.
    function colorByRole(string memory name) public pure returns (string memory colored) {
        bool isOperator = _contains(name, "operator");
        bool isStaker = _contains(name, "staker");
        bool isAVS = _contains(name, "avs");

        if (isOperator) colored = name.blue();
        else if (isStaker) colored = name.cyan();
        else if (isAVS) colored = name.magenta();
        else colored = name.yellow();
    }

    /// @dev Returns `true` if `needle` is found in `haystack`.
    function _contains(string memory haystack, string memory needle) internal pure returns (bool) {
        return cheats.indexOf(haystack, needle) != type(uint).max;
    }

    /// -----------------------------------------------------------------------
    /// Cheats
    /// -----------------------------------------------------------------------

    function rollForward(uint blocks) internal {
        cheats.roll(block.number + blocks);
        console.log("%s.roll(+ %d blocks)", colorByRole("cheats"), blocks);
    }

    function rollBackward(uint blocks) internal {
        cheats.roll(block.number - blocks);
        console.log("%s.roll(- %d blocks)", colorByRole("cheats"), blocks);
    }

    /// -----------------------------------------------------------------------
    /// Logging
    /// -----------------------------------------------------------------------

    function _toggleLog() internal {
        logging = !logging;
        console.log("\n%s logging %s...", NAME_COLORED(), logging ? "enabled" : "disabled");
    }

    /// -----------------------------------------------------------------------
    /// Trace Counter get/set
    /// -----------------------------------------------------------------------

    function _getTraceCounter() internal view returns (uint) {
        return uint(cheats.load(LOG_STATE_ADDR, LOG_STATE_SLOT));
    }

    function _setTraceCounter(uint _newValue) internal {
        cheats.store(LOG_STATE_ADDR, LOG_STATE_SLOT, bytes32(_newValue));
    }
}

/// @dev Assumes the user is a `Logger`.
library print {
    using print for *;
    using StdStyle for *;

    /// -----------------------------------------------------------------------
    /// Logging
    /// -----------------------------------------------------------------------

    function method(string memory m) internal view {
        if (!_logging()) return;
        console.log("%s.%s()", _name(), m.italic());
    }

    function method(string memory m, string memory args) internal view {
        if (!_logging()) return;
        console.log("%s.%s(%s)", _name(), m.italic(), args);
    }

    function user(string memory name, uint assetType, uint userType, IStrategy[] memory strategies, uint[] memory tokenBalances)
        internal
        view
    {
        if (!_logging()) return;
        console.log("\nCreated %s %s who holds %s.", userType.asUserType(), _logger().colorByRole(name), assetType.asAssetType());

        console.log("   Balances:");
        for (uint i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            if (strat == BEACONCHAIN_ETH_STRAT) {
                console.log("       Native ETH: %s", print.asWad(tokenBalances[i]));
            } else {
                IERC20 underlyingToken = strat.underlyingToken();
                console.log("       %s: %s", IERC20Metadata(address(underlyingToken)).name(), print.asGwei(tokenBalances[i]));
            }
        }
    }

    function gasUsed() internal {
        uint used = cheats.snapshotGasLastCall("gasUsed");
        if (!_logging()) return;
        console.log("   Gas used: %d".dim().bold(), used);
    }

    /// -----------------------------------------------------------------------
    /// Logging
    /// -----------------------------------------------------------------------

    function createOperatorSets(IAllocationManagerTypes.CreateSetParams[] memory p) internal pure {
        console.log("Creating operator sets:");
        for (uint i; i < p.length; ++i) {
            console.log("   operatorSet%d:".yellow(), p[i].operatorSetId);
            for (uint j; j < p[i].strategies.length; ++j) {
                console.log("       strategy%s: %s", cheats.toString(j), address(p[i].strategies[j]));
            }
        }
    }

    function deregisterFromOperatorSets(IAllocationManagerTypes.DeregisterParams memory p) internal pure {
        console.log("Deregistering operator: %s", address(p.operator));
        console.log("   from operator sets:");
        for (uint i; i < p.operatorSetIds.length; ++i) {
            console.log("       operatorSet%d:".yellow(), p.operatorSetIds[i]);
        }
    }

    /// -----------------------------------------------------------------------
    /// Strings
    /// -----------------------------------------------------------------------

    function asAssetType(uint t) internal pure returns (string memory s) {
        if (t == HOLDS_ALL) s = "ALL_ASSETS";
        else if (t == HOLDS_LST) s = "LST";
        else if (t == HOLDS_ETH) s = "ETH";
        else if (t == NO_ASSETS) s = "NO_ASSETS";
    }

    function asUserType(uint t) internal pure returns (string memory s) {
        if (t == DEFAULT) s = "DEFAULT";
        else if (t == ALT_METHODS) s = "ALT_METHODS";
    }

    function asForkType(uint t) internal pure returns (string memory s) {
        if (t == LOCAL) s = "LOCAL";
        else if (t == MAINNET) s = "MAINNET";
        else if (t == HOLESKY) s = "HOLESKY";
    }

    function asGwei(uint x) internal pure returns (string memory) {
        return x.asDecimal(9, " gwei");
    }

    function asWad(uint x) internal pure returns (string memory) {
        return x.asDecimal(18, " wad");
    }

    function asDecimal(uint x, uint8 decimals, string memory denomination) internal pure returns (string memory s) {
        if (x == 0) return string.concat("0.0", denomination);

        s = cheats.toString(x);

        while (bytes(s).length < decimals) s = string.concat("0", s);

        uint len = bytes(s).length;
        bytes memory b = bytes(s);
        bytes memory left = new bytes(len > decimals ? len - decimals : 0);
        bytes memory right = new bytes(decimals);

        for (uint i; i < left.length; ++i) {
            left[i] = b[i];
        }

        for (uint i; i < decimals; ++i) {
            right[i] = b[len - decimals + i];
        }

        return string.concat(left.length > 0 ? string(left) : "0", ".", string(right), denomination);
    }

    /// -----------------------------------------------------------------------
    /// Helpers
    /// -----------------------------------------------------------------------

    function _name() internal view returns (string memory) {
        try _logger().NAME_COLORED() returns (string memory name) {
            return name;
        } catch {
            revert("This contract is not a `Logger`.");
        }
    }

    function _logger() internal view returns (Logger) {
        return Logger(address(this));
    }

    function _logging() internal view returns (bool) {
        return _logger().logging();
    }
}
