// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/utils/math/Math.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "src/test/integration/IntegrationChecks.t.sol";

/// Not sure this test is actually useful, mostly used this to play around with the numbers.
/// The conclusion is self-evident from the replicated math...
contract Integration_Rounding_Deallocation is IntegrationCheckUtils {

    using ArrayLib for *;
    using StdStyle for *;
    using Strings for *;
    using Math for *;

    struct Magnitude {
        uint64 max;
        uint64 cur;
        uint64 available;
        uint64 pendDealloc;
    }

    enum ValueMod {
        SUB,
        NONE,
        ADD
    }

    struct Slash {
        uint8 mantissa;
        uint8 exponent;
        ValueMod vMod;
    }

    Magnitude mag;
    Slash slash;

    using SlashingLib for *;

    DepositScalingFactor dsf;

    function test_DSF_Manip(uint scalingFactor, uint prevDepositShares, uint addedShares, uint slashingFactor) public rand(0) {
        scalingFactor = bound(scalingFactor, 1, WAD);
        prevDepositShares = bound(prevDepositShares, 0, uint(type(uint8).max));
        addedShares = bound(addedShares, 0, uint(type(uint8).max));
        slashingFactor = bound(slashingFactor, 1, 1e18);

        dsf._scalingFactor = scalingFactor;
        // prevDepositShares = 2;
        // addedShares = 2;
        // slashingFactor = 628157962353078290;

        console.log("WAD:               ", uint(1e18));
        console.log("PrevDepositShares: ", prevDepositShares);
        console.log("AddedShares:       ", addedShares);
        console.log("slashingFactor:    ", slashingFactor);
        console.log("scalingFactor:     ", dsf._scalingFactor);

        dsf.update(prevDepositShares, addedShares, slashingFactor);

        console.log("newScalingFactor:  ", dsf._scalingFactor);

        assertTrue(dsf._scalingFactor >= scalingFactor, "scalingFactor decreased!");
        assertTrue(dsf._scalingFactor != 0, "scalingFactor hit 0!");
    }

    struct DSFState {
        DepositScalingFactor dsf;
        uint prevScalingFactor;
        uint operatorMaxMag;
        uint prevDepositShares;
        uint addedShares;
    }

    DSFState state;
    DepositScalingFactor dummy;

    bool writeJSON = false;

    function _init() internal override {
        string memory tomlRAW = cheats.readFile("src/test/integration/params.toml");

        printLast = cheats.parseTomlUint(tomlRAW, ".printLast");
        printEvery = cheats.parseTomlUint(tomlRAW, ".printEvery");
        writeJSON = cheats.parseTomlUint(tomlRAW, ".writeJSON") == 1;
        console.log("Printing final %d inputs", printLast);
        console.log("Printing other inputs every %d iterations", printEvery);
        console.log("Writing to file: %s", writeJSON ? "true" : "false");
    }

    uint[] DSFs;

    function test_DSF_DownToZero() public rand(0) {
        (uint prevDepositShares, uint iterations) = _readTOML();
        finalIteration = iterations;

        console.log("Start Balance: %s (%d)", _pretty(prevDepositShares), prevDepositShares);
        console.log("Iterations: %d", iterations);

        state.dsf._scalingFactor = WAD;
        state.prevScalingFactor = state.dsf._scalingFactor;
        state.operatorMaxMag = WAD-1;
        state.addedShares = 1;
        state.prevDepositShares = prevDepositShares;

        for (uint i = 0; i < iterations; i++) {
            _print(i);

            state.prevScalingFactor = state.dsf._scalingFactor;

            state.dsf.update({
                prevDepositShares: state.prevDepositShares,
                addedShares: state.addedShares,
                slashingFactor: state.operatorMaxMag
            });

            DSFs.push(state.dsf._scalingFactor);

            state.prevDepositShares += state.addedShares;
        }

        _print(iterations);
        _writeJSON(".DTZ");

        revert();
    }

    function test_DSF_FindOptimalIterations() public rand(0) {
        (
            uint startDSF,
            uint startShares, 
            uint startIteration,
            uint contIterations
        ) = _readFindContinuationTOML();
        finalIteration = startIteration + contIterations;
        // printInputs = true;

        console.log("Start Balance: %s (%d)", _pretty(startShares), startShares);
        // console.log("Iterations: %d", iterations);

        state.dsf._scalingFactor = startDSF;
        state.prevScalingFactor = startDSF;
        state.operatorMaxMag = WAD-1;
        state.addedShares = 1;
        state.prevDepositShares = startShares;

        for (uint i = startIteration; i < finalIteration; i++) {
            _print(i);

            state.prevScalingFactor = state.dsf._scalingFactor;

            state.dsf.update({
                prevDepositShares: state.prevDepositShares,
                addedShares: state.addedShares,
                slashingFactor: state.operatorMaxMag
            });

            // if (i + 1 == 104) printInputs = true;

            // Iterate until decrease is around 50%
            uint scaledDecrease = 1e5 - ((1e5 * state.dsf._scalingFactor) / state.prevScalingFactor);
            if (scaledDecrease >= 49e3) {
                console.log("Peak found at %d iterations!".green().bold(), i+1);
                printInputs = true;
                finalIteration = i + printLast;
            }

            state.prevDepositShares += state.addedShares;
        }

        _print(finalIteration);

        revert();
    }

    function test_DSF_ContinueIteration() public rand(0) {
        (
            uint startDSF,
            uint startShares,
            uint startSharesToAdd,
            uint maxMagDecrease,
            uint startIteration,
            uint contIterations
        ) = _readContinueTOML();

        {
            state.dsf._scalingFactor = startDSF;
            state.prevScalingFactor = startDSF;
            state.operatorMaxMag = WAD - 1;

            // console.log("operatorMaxMag: %s (%d)", _pretty(state.operatorMaxMag), state.operatorMaxMag);

            state.prevDepositShares = startShares;
            state.addedShares = startSharesToAdd;
            console.log("prevDepositShares: %d", state.prevDepositShares);
        
            printInputs = true;
            finalIteration = startIteration + contIterations;
        }

        state.dsf.update({
            prevDepositShares: state.prevDepositShares,
            addedShares: state.addedShares,
            slashingFactor: state.operatorMaxMag
        });

        DSFs.push(state.dsf._scalingFactor);

        console.log("prevDepositShares: %d", state.prevDepositShares);
        state.prevDepositShares += state.addedShares;
        state.operatorMaxMag -= maxMagDecrease;

        bool peakFound;
        for (uint i = startIteration; i < startIteration+contIterations; i++) {
            _print(i);
            if (peakFound) printInputs = false;

            state.addedShares = 1;
            state.prevScalingFactor = state.dsf._scalingFactor;

            console.log("prevDepositShares: %d", state.prevDepositShares);

            state.dsf.update({
                prevDepositShares: state.prevDepositShares,
                addedShares: state.addedShares,
                slashingFactor: state.operatorMaxMag
            });

            DSFs.push(state.dsf._scalingFactor);

            // Print inputs until peak is found, then stop printing inputs
            uint scaledDecrease = 1e5 - ((1e5 * state.dsf._scalingFactor) / state.prevScalingFactor);
            if (scaledDecrease >= 49e3) {
                console.log("Peak found at %d iterations".green().dim(), i+1);
                peakFound = true;
            }

            state.prevDepositShares += state.addedShares;
        }

        _print(startIteration+contIterations);
        _writeJSON(".continued");
        revert();
    }

    function _readTOML() internal returns (uint, uint) {
        string memory tomlRAW = cheats.readFile("src/test/integration/params.toml");

        return (
            cheats.parseTomlUint(tomlRAW, ".prevDepositShares"),
            cheats.parseTomlUint(tomlRAW, ".iterations")
        );
    }

    function _readFindContinuationTOML() internal returns (uint, uint, uint, uint) {
        string memory tomlRAW = cheats.readFile("src/test/integration/params.toml");

        return (
            cheats.parseTomlUint(tomlRAW, ".startDSF"),
            cheats.parseTomlUint(tomlRAW, ".startShares"),
            cheats.parseTomlUint(tomlRAW, ".iterations") + 1,
            cheats.parseTomlUint(tomlRAW, ".contIterations")
        );
    }

    function _readContinueTOML() internal returns (uint, uint, uint, uint, uint, uint) {
        string memory tomlRAW = cheats.readFile("src/test/integration/params.toml");

        return (
            cheats.parseTomlUint(tomlRAW, ".startDSF"),
            cheats.parseTomlUint(tomlRAW, ".startShares"),
            cheats.parseTomlUint(tomlRAW, ".startSharesToAdd"),
            cheats.parseTomlUint(tomlRAW, ".maxMagDecrease"),
            cheats.parseTomlUint(tomlRAW, ".iterations") + 1,
            cheats.parseTomlUint(tomlRAW, ".contIterations")
        );
    }

    function _writeJSON(string memory arrKey) internal {
        if (!writeJSON) return;

        string memory outputJSON = cheats.serializeUint("", "", DSFs);

        cheats.writeJson(outputJSON, "src/test/integration/output.json", arrKey);
    }

    /// Minimum:
    ///
    /// DSF:            999000999000999
    /// operatorMaxMag: 999999999999999999
    /// depositShares:  1001

    /// Deposit-only method:
    ///
    /// numDeposits:      1000
    /// - optimalStart:   993
    /// - final DSF:      5.01e14 (501756146512794)
    ///
    /// numDeposits:      2000
    /// - optimalStart:   1993
    /// - final DSF:      2.50e14 (250438266967192)
    ///
    /// numDeposits:      3000
    /// - optimalStart:   2989
    /// - final DSF:      1.67e14 (166944908180300)
    ///
    /// numDeposits:      4000
    /// - optimalStart:   3989
    /// - final DSF:      1.25e14 (125172111653523)

    function test_DSF_SpecificUpdate() public rand(0) {
        dummy._scalingFactor = 1;
        dummy.update({
            prevDepositShares: WAD,
            addedShares: 1,
            slashingFactor: 1e18 - 1
        });

        console.log("finalDSF:                %s (%d)", _pretty(dummy._scalingFactor), dummy._scalingFactor);
        assertTrue(dummy._scalingFactor != 0, "found 0");
    }

    uint finalIteration = type(uint).max;
    uint printLast = 3;
    uint printEvery = 100;
    bool printInputs = false;

    function _print(uint iter) internal {
        // print every 100 iterations, but also print the final 20 iterations
        if (!printInputs && iter % printEvery != 0) {
            if (printLast <= finalIteration && iter < finalIteration - printLast) {
                return;
            }
        }

        // if (!printInputs && iter % 10000000 != 0 && iter < finalIteration - printLast) return; 
        
        console.log("(%d)".dim(), iter);

        // if (printInputs || iter > finalIteration - printLast) {
        if (true) {
            console.log("==Inputs==".yellow());
            // console.log(" - prevDepositShares:     %d".dim(), state.prevDepositShares - state.addedShares);
            console.log(" - prevDepositShares:     %d".dim(), state.prevDepositShares);
            console.log(" - addedShares:           %d".dim(), state.addedShares);
            console.log(" - prevDSF:               %s (%d)".dim(), _pretty(state.prevScalingFactor), state.prevScalingFactor);

            {
                console.log("=currentShares=".yellow().dim());

                // uint pDS = state.prevDepositShares - state.addedShares;
                uint pDS = state.prevDepositShares;
                uint pDSMulDSF = pDS * state.prevScalingFactor;
                uint numerator = pDS.mulWad(state.prevScalingFactor) * state.operatorMaxMag;
                uint currentShares = pDS.mulWad(state.prevScalingFactor).mulWad(state.operatorMaxMag);

                console.log(" - prevDepositShares*DSF: %s (%d)".dim(), _pretty(pDSMulDSF), pDSMulDSF);
                console.log(" - numerator:             %s (%d)".dim(), _pretty(numerator), numerator);
                console.log(" - currentShares:         %d".dim(), currentShares);
            }

            console.log("=newDSF=".yellow().dim());
        } 
        
        // uint totalDepositShares = state.prevDepositShares + state.addedShares;
        // uint totalWithdrawable = state.dsf.calcWithdrawable(totalDepositShares, state.operatorMaxMag);
        // console.log(" - totalDepositShares:     %d", totalDepositShares);
        // console.log(" - totalWithdrawable:      %d", totalWithdrawable);
        
        // console.log(" - operatorMaxMag:     %s (%d)", _pretty(state.operatorMaxMag), state.operatorMaxMag);
        if (state.dsf._scalingFactor < 1e15) {
            console.log(
                " - DSF:                   %s (%d)".yellow(),
                _pretty(state.dsf._scalingFactor), 
                state.dsf._scalingFactor
            );
        } else if (state.dsf._scalingFactor < 1e12) {
            console.log(
                " - DSF:                   %s (%d)".red(),
                _pretty(state.dsf._scalingFactor), 
                state.dsf._scalingFactor
            );
        } else {
            console.log(
                " - DSF:                   %s (%d)".green(),
                _pretty(state.dsf._scalingFactor), 
                state.dsf._scalingFactor
            );
        }

        if (state.prevScalingFactor < state.dsf._scalingFactor) {
            uint increase = state.dsf._scalingFactor - state.prevScalingFactor;

            console.log(
                " -- increase: %s (%s)".green().dim(),
                _prettyPerc({_old: state.dsf._scalingFactor, _new: state.prevScalingFactor}),
                _pretty(increase)
            );
        } else if (state.prevScalingFactor == state.dsf._scalingFactor) {
            console.log(" -- no change".yellow().dim());
        } else {
            uint decrease = state.prevScalingFactor - state.dsf._scalingFactor;
            
            console.log(
                " -- decrease: %s (%s)".red().dim(), 
                _prettyPerc({_new: state.dsf._scalingFactor, _old: state.prevScalingFactor}),
                _pretty(decrease)
            );
        }

        // console.log(" - 1 WAD:              %s (%d)".dim(), _pretty(WAD), WAD);
    }

    /// Expresses a % decrease with 3 decimals; e.g: 99.999%
    function _prettyPerc(uint _new, uint _old) internal view returns (string memory) {
        // Scale by 1e5 so we can extract 3 decimals (e.g. 99999 = 99.999%)
        uint scaled = 1e5 - ((1e5 * _new) / _old);

        uint whole = scaled / 1000;          // 99
        uint decimal = scaled % 1000;        // 999

        // Pad decimal to 3 digits
        string memory paddedDecimal;
        if (decimal < 10) {
            paddedDecimal = string(abi.encodePacked("00", decimal.toString()));
        } else if (decimal < 100) {
            paddedDecimal = string(abi.encodePacked("0", decimal.toString()));
        } else {
            paddedDecimal = decimal.toString();
        }

        return string(
            abi.encodePacked(
                whole.toString(), ".", 
                paddedDecimal, "%"
            )
        );
    }

    // pretty print a wad-based value, e.g. "4.23e17"
    function _pretty(uint value) internal view returns (string memory) {
        if (value == 0) {
            return "0.00e0";
        }

        uint temp = value;
        uint digits = 0;
        while (temp >= 10) {
            temp /= 10;
            digits++;
        }

        // Extract first 5 significant digits (rounded down)
        uint divisor = 10 ** (digits > 4 ? digits - 4 : 0);
        uint leading = value / divisor; // e.g. 56789 for 56789112

        // Format: x.yyyy
        uint first = leading / 10000;            // x
        uint second = (leading / 1000) % 10;     // y
        uint third = (leading / 100) % 10;       // y
        uint fourth = (leading / 10) % 10;       // y
        uint fifth = leading % 10;               // y

        return string(
            abi.encodePacked(
                first.toString(), ".", 
                second.toString(), 
                third.toString(), 
                fourth.toString(), 
                fifth.toString(),
                "e", 
                digits.toString()
            )
        );
    }

    /// @dev Mocks out math performed in AllocationManager.slashOperator. This examines precision drift
    /// caused by slashing applied when the operator has a pending deallocation, as slashOperator applies
    /// slashes equally to two differently sized values:
    ///
    /// uint64 slashedMagnitude = uint64(uint256(allocation.currentMagnitude).mulWadRoundUp(params.wadsToSlash[i]));
    /// 
    /// vs
    /// 
    /// uint64 slashedPending =
    ///             uint64(uint256(uint128(-allocation.pendingDiff)).mulWadRoundUp(params.wadsToSlash[i]));
    ///
    /// The goal was to determine how precision drift impacts an operator's max magnitude over time. In conclusion,
    /// any precision drift/rounding that occurs causes the operator to have _less_ max magnitude than they should,
    /// reducing the ability of an attacker to leverage magnitude for an attack. We can rule out pending deallocations 
    /// when considering slash/redistro attacks.
    function test_deallocMath(
        uint64 maxMag,
        uint64 curMag, 
        uint64 pendDealloc, 
        uint8 slashMantissa, 
        uint8 slashExponent,
        uint8 vMod
    ) public rand(0) {
        maxMag = uint64(bound(maxMag, 2, 1e18));
        curMag = uint64(bound(curMag, 1, maxMag));
        pendDealloc = uint64(bound(pendDealloc, 1, curMag));

        // Slashes are XeY +/- 1
        slashMantissa = uint8(bound(slashMantissa, 1, 9));
        slashExponent = uint8(bound(slashExponent, 0, 17));
        vMod = uint8(bound(vMod, uint(type(ValueMod).min), uint(type(ValueMod).max)));

        mag = Magnitude({
            max: maxMag,
            cur: curMag,
            available: maxMag - curMag,
            pendDealloc: pendDealloc
        });

        slash = Slash({
           mantissa: slashMantissa,
           exponent: slashExponent,
           vMod: ValueMod(vMod)
        });

        _printMagInfo();

        for (uint i = slashExponent; i > 0; i--) {
            _doSlash();
            _printMagInfo();

            slash.exponent--;
        }

        revert();
    }

    function _doSlash() internal {
        console.log("");
        console.log("===Slash Info===".magenta());

        uint wadsToSlash = uint(slash.mantissa) * 10 ** uint(slash.exponent);
        if (slash.vMod == ValueMod.SUB && wadsToSlash > 1) {
            wadsToSlash--;
            console.log(" - wadsToSlash: %de%d-1 (%d)", slash.mantissa, slash.exponent, wadsToSlash);
        } else if (slash.vMod == ValueMod.ADD && wadsToSlash < uint(WAD)) {
            wadsToSlash++;
            console.log(" - wadsToSlash: %de%d+1 (%d)", slash.mantissa, slash.exponent, wadsToSlash);
        } else {
            console.log(" - wadsToSlash: %de%d (%d)", slash.mantissa, slash.exponent, wadsToSlash);
        }

        // Apply the same slashing to current allocation and pending deallocation
        uint64 slashedMagnitude = uint64(SlashingLib.mulWadRoundUp(mag.cur, wadsToSlash));
        uint64 slashedPending = uint64(SlashingLib.mulWadRoundUp(mag.pendDealloc, wadsToSlash));

        // Check to see if rounding occurred
        bool sMRounding = SlashingLib.mulWad(mag.cur, wadsToSlash) != slashedMagnitude;
        bool sPRounding = SlashingLib.mulWad(mag.pendDealloc, wadsToSlash) != slashedPending;
        console.log(" - slashedMagnitude: %d (rounding: %s)", slashedMagnitude, sMRounding ? "true".magenta() : "false");
        console.log(" - slashedPending:   %d (rounding: %s)", slashedPending, sPRounding ? "true".magenta() : "false");

        // Update magnitude info according to slash
        mag.max -= slashedMagnitude;
        mag.cur -= slashedMagnitude;
        mag.available = mag.max - mag.cur;
        mag.pendDealloc -= slashedPending;

        // Simulate applying deallocation and guage how much magnitude the operator is now able to allocate
        // uint64 newAvailable = mag.available + mag.pendDealloc;
        // uint64 newCurMag = mag.cur - mag.pendDealloc;
        // uint64 newMaxMag = newAvailable + newCurMag;
        uint64 newCurMag = mag.cur - mag.pendDealloc;
        uint64 newAvailable = mag.max - newCurMag;
        uint64 effectiveMaxMag = newAvailable + newCurMag;

        console.log(" - after deallocation:".yellow());
        console.log(" -- new cur mag:     ", newCurMag);
        console.log(" -- new available:   ", newAvailable);

        if (mag.max == effectiveMaxMag) {
            console.log(" -- max mag loss: %s", "none".green());
        } else if (mag.max > effectiveMaxMag) {
            console.log(
                " -- max mag loss: %s (%d)", 
                "lost magnitude".yellow(),
                mag.max - effectiveMaxMag
            );

            revert("lossy");
        } else {
            console.log(
                " -- max mag loss: %s (%d)", 
                "gained magnitude".red(),
                effectiveMaxMag - mag.max
            );

            revert("gain");
        }
    }

    function _printMagInfo() internal {
        console.log("");    
        console.log("===Magnitude Info===".cyan());

        console.log(" - initial max mag:  ", WAD);
        console.log(" - max magnitude:    ", mag.max);
        console.log(" -- available:       ", mag.available);
        console.log(" -- allocated:       ", mag.cur);
        console.log(" -- pending dealloc: ", mag.pendDealloc);
    }
}