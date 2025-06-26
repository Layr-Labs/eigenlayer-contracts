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
    using SlashingLib for *;

    uint printLast = 3;
    uint printEvery = 100;
    bool writeJSON = false;
    string writeLocation = "src/test/integration/output.json";

    struct InitTOML {
        uint printEvery;
        uint printLast;
        bool writeJSON;
        string writeLocation;
    }

    function _init() internal override {
        string memory tomlRAW = cheats.readFile("src/test/integration/params.toml");
        bytes memory data = vm.parseToml(tomlRAW, ".init");
        InitTOML memory it = abi.decode(data, (InitTOML));

        printLast = it.printLast;
        printEvery = it.printEvery;
        writeJSON = it.writeJSON;
        writeLocation = it.writeLocation;

        console.log("");
        console.log("==Init==".magenta());

        // console.log(" - printing final %d inputs", printLast);
        // console.log(" - printing other inputs every %d iterations", printEvery);
        if (writeJSON) {
            console.log(" - writing to file: %s".green(), writeLocation);
        } else {
            console.log(" - writing to file: false");
        }
    }

    struct GenerateSettings {
        // Initial run params
        uint totalShares;
        uint wadsInWithdrawalQueue;
        uint64 maxMagnitude;
        uint numDataPoints;

        // Derived params
        uint depositSharesToWithdraw;
        uint operatorShares;
        uint scaledShares;

        // Plot/Line/Labels
        string plotName;
        string lineLabel;

        // File IO
        bool writeToFile;
        string writeLocation;
    }

    struct GenerateOutput {
        uint[] x;
        uint[] y;
    }

    function test_Generate() public rand(0) {
        GenerateSettings memory settings = _readGenerateTOML();

        GenerateOutput memory data = GenerateOutput({
            x: new uint[](settings.numDataPoints),
            y: new uint[](settings.numDataPoints)
        });

        for (uint i = 0; i < settings.numDataPoints; i++) {
            /// 3. Calc how much to slash, proportional to the number of 
            uint wadsToSlash = (i + 1).mulDiv(WAD, settings.numDataPoints + 1);

            /// 4. Commit offense and get slashed
            uint64 slashedMagnitude = uint64(uint(settings.maxMagnitude).mulWadRoundUp(wadsToSlash));
            uint64 newMaxMag = settings.maxMagnitude - slashedMagnitude;

            /// 5. How much gets redistributed?
            // From operator shares...
            uint amountOpShares = settings.operatorShares.calcSlashedAmount({
                prevMaxMagnitude: settings.maxMagnitude,
                newMaxMagnitude: newMaxMag
            });
            // From withdrawal queue...
            uint amountWQShares = settings.scaledShares.scaleForBurning({
                prevMaxMagnitude: settings.maxMagnitude,
                newMaxMagnitude: newMaxMag
            });
            // Total:
            uint totalRedistributed = amountOpShares + amountWQShares;
            if (totalRedistributed > settings.totalShares) {
                revert("minted shares!!".red().bold());
            }

            // 6. Add points to data
            data.x[i] = uint(slashedMagnitude);
            data.y[i] = totalRedistributed;
        }

        _writeGenerateJSON(settings, data);

        revert("finished".green().bold());
    }

    struct GenerateTOML {
        uint totalShares;
        uint wadsInWithdrawalQueue;
        uint maxMagnitudeExp;
        uint numDataPoints;
        string plotName;
        string lineLabel;
        bool writeToFile;
        string writeLocationBase;
        string fileName;
    }

    function _readGenerateTOML() internal returns (GenerateSettings memory settings) {
        string memory tomlRAW = cheats.readFile("src/test/integration/params.toml");
        bytes memory data = vm.parseToml(tomlRAW, ".generate");
        GenerateTOML memory gt = abi.decode(data, (GenerateTOML));
        
        settings = GenerateSettings({
            totalShares: gt.totalShares,
            wadsInWithdrawalQueue: gt.wadsInWithdrawalQueue,
            maxMagnitude: uint64(10 ** gt.maxMagnitudeExp),
            numDataPoints: gt.numDataPoints,
            depositSharesToWithdraw: 0,
            operatorShares: 0,
            scaledShares: 0,
            plotName: gt.plotName,
            lineLabel: gt.lineLabel,
            writeToFile: gt.writeToFile,
            writeLocation: string.concat(gt.writeLocationBase, gt.fileName)
        });

        // Adjust/calculate
        if (settings.numDataPoints > uint(settings.maxMagnitude)) {
            settings.numDataPoints = uint(settings.maxMagnitude);
        }

        console.log("");
        console.log("==Generate Settings==".magenta());

        console.log(" plot: '%s'".yellow(), settings.plotName);
        console.log(" line: '%s'".yellow(), settings.lineLabel);
        if (settings.writeToFile) {
            console.log(" - writing to file: %s".green(), settings.writeLocation);
        } else {
            console.log(" - writing to file: false".green().dim());
        }

        console.log("");
        console.log("==Run Info==".magenta());

        /// 1. depositIntoStrategy: Simulate a deposit of totalShares into operator with given magnitude
        scratchDSF.reset();
        scratchDSF.update({
            prevDepositShares: 0,
            addedShares: settings.totalShares,
            slashingFactor: settings.maxMagnitude
        });

        /// 2. queueWithdrawal: Queue a proportion of deposit shares for withdrawal
        // Amount of deposit shares to withdraw
        settings.depositSharesToWithdraw = settings.totalShares.mulWad(settings.wadsInWithdrawalQueue);
        // Amount to reduce operator's shares by
        uint withdrawable_Mid = scratchDSF.calcWithdrawable(settings.depositSharesToWithdraw, settings.maxMagnitude);
        settings.operatorShares = settings.totalShares - withdrawable_Mid;
        // Scaled shares added to the withdrawal queue
        settings.scaledShares = scratchDSF.scaleForQueueWithdrawal(settings.depositSharesToWithdraw);

        
        // console.log(" -- depositSh. to queue:   %s (%d)".dim().yellow().italic(), _pretty(settings.depositSharesToWithdraw), settings.depositSharesToWithdraw);
        console.log(" - numDataPoints:          %d".dim().yellow(), settings.numDataPoints);
        console.log(" - maxMagnitude:           %s (%d)".dim().yellow(), _pretty(settings.maxMagnitude), settings.maxMagnitude);
        console.log(" - totalShares:            %s (%d)".dim().yellow(), _pretty(settings.totalShares), settings.totalShares);
        console.log(" -- % in withdrawal queue: %s (raw wads: %d)".dim().yellow(), _prettyPerc(gt.wadsInWithdrawalQueue, 1e18), gt.wadsInWithdrawalQueue);
        console.log(" -- operatorShares:        %s (%d)".dim().yellow(), _pretty(settings.operatorShares), settings.operatorShares);
        console.log(" -- queuedShares (scaled): %s (%d)".dim().yellow(), _pretty(settings.scaledShares), settings.scaledShares);

        console.log("");

        return settings;
    }

    function _writeGenerateJSON(
        GenerateSettings memory settings, 
        GenerateOutput memory output
    ) internal {
        require(output.x.length == output.y.length, "length err");

        if (!settings.writeToFile) return;

        console.log("");
        console.log(" writing output to file: %s...".green().dim(), settings.writeLocation);

        // Serialize x and y points
        string memory data = cheats.serializeUint(settings.lineLabel, "dataX", output.x);
        data = cheats.serializeUint(settings.lineLabel, "dataY", output.y);

        // Get key for which data will be replaced
        string memory valueKey = string.concat(".", settings.plotName, ".", settings.lineLabel);
        console.log(" overwriting existing value for key: %s".green().dim(), valueKey);

        // Write to file
        cheats.writeJson(data, settings.writeLocation, valueKey);
        console.log(" done!\n".green());
    }

    function test_WTF(uint8 sharesExp, uint64 diff) public rand(0) {
        // prevMM = uint64(bound(prevMM, 1, WAD));
        // newMM = prevMM == 1 ? 0 : uint64(bound(newMM, 0, prevMM-1));
        diff = uint64(bound(diff, 1, WAD));
        // uint shares = uint(WAD / diff);
        // shares > 2 ? shares-- : shares;
        sharesExp = uint8(bound(sharesExp, 1, 32));
        // uint shares = bound(shares, 10, 1e18);
        uint shares = 10 ** sharesExp;

        uint64 prevMM = WAD;
        uint64 newMM = prevMM - uint64(diff);

        console.log(" - shares:      %s (%d)", _pretty(shares), shares);
        console.log(" - prevMM:      %s (%d)", _pretty(prevMM), prevMM);
        console.log(" - newMM:       %s (%d)", _pretty(newMM), newMM);

        uint opShares = shares.calcSlashedAmount(prevMM, newMM);
        uint wqShares = shares.scaleForBurning(prevMM, newMM);

        console.log(" - opShares:  %s (%d)", _pretty(opShares), opShares);
        console.log(" - wqShares:  %s (%d)", _pretty(wqShares), wqShares);

        assertEq(opShares, wqShares, "fail".green().bold());
    }

    uint OPERATOR_SHARES;
    uint WADS_TO_SLASH;

    struct Params {
        // Random:
        uint64 prevMM;
        uint wadsInWQ;
        // Derived:
        uint64 newMM;
        uint depositSharesToWithdraw;
        uint operatorShares;
        uint scaledShares;

        // Derived - withdrawable:
        uint dsf;
        uint withdrawableStart;

        // Results:
        uint expectedRedistribution;
        uint amountRedistributed;
        uint completeWithdrawal;
        uint withdrawRemaining;
        uint completeExit;
        uint totalWithdrawn;
        uint totalMissing;
        uint abs_Error;
    }

    DepositScalingFactor scratchDSF;
    Params[] points;

    uint min_Amt = type(uint).max;
    uint max_Amt = 0;

    function test_MaxExtract() public rand(0) {
        _readExtractTOML();

        // 3 different strategies to explore:
        //
        // 1. Minimize the amount the AVS can redistribute
        // 2. Maximize the amount the operator can recover via queued withdrawals
        // 3. Minimize the amount the operator can recover via queued withdrawals
        // -- this last one is to simulate "operator griefing stakers"
        //
        // The idea is that "the operator expects a 10% slash", and sets up:
        // - prevMM (their allocation)
        // -- Might want to parameterize this somewhat, assuming AVSs don't accept
        //    allocations below some size (MINIMUM_MAGNITUDE_ACCEPTED)
        // - the amount of their assets in the withdrawal queue vs out (operatorShares vs queuedShares)
        // - their DSF !? 
        // -- we can lower this pretty easily... (but this would apply to 2, not 3)
        // -- what if this starts with intentionally slashing themselves and depositing to get a positive DSF?

        Params memory params;

        for (uint i = 0; i < NUM_ITERATIONS; i++) {
            scratchDSF.reset();

            // Generate random params:
            params.prevMM = uint64(cheats.randomUint(1, WAD));
            params.wadsInWQ = cheats.randomUint(1, WAD);

            {
                /// Calc derived params:

                /// 1. Deposit OPERATOR_SHARES into operator with max mag == prevMM
                scratchDSF.update({
                    prevDepositShares: 0,
                    addedShares: OPERATOR_SHARES,
                    slashingFactor: params.prevMM
                });
                params.dsf = scratchDSF._scalingFactor;
                params.withdrawableStart = scratchDSF.calcWithdrawable(OPERATOR_SHARES, params.prevMM);

                /// 2. Queue a proportion of the shares for withdrawal
                params.depositSharesToWithdraw = OPERATOR_SHARES.mulWad(params.wadsInWQ);
                // Amount to reduce operator's shares by
                uint withdrawable_Mid = scratchDSF.calcWithdrawable(params.depositSharesToWithdraw, params.prevMM);
                params.operatorShares = OPERATOR_SHARES - withdrawable_Mid;
                // Amount we add to withdrawal queue (and withdrawal object)
                params.scaledShares = scratchDSF.scaleForQueueWithdrawal(params.depositSharesToWithdraw);

                /// 3. Commit offense and get slashed
                uint64 slashedMagnitude = uint64(uint(params.prevMM).mulWadRoundUp(WADS_TO_SLASH));
                params.newMM = params.prevMM - slashedMagnitude;           
                params.expectedRedistribution = OPERATOR_SHARES.calcSlashedAmount(params.prevMM, params.newMM);
            }

            {
                // How much actually gets redistributed?
                uint amountOpShares = params.operatorShares.calcSlashedAmount(params.prevMM, params.newMM);
                uint amountWQShares = params.scaledShares.scaleForBurning(params.prevMM, params.newMM);

                params.amountRedistributed = amountOpShares + amountWQShares;
            }

            {
                // How much can the operator withdraw by completing their existing withdrawal
                params.completeWithdrawal = params.scaledShares.scaleForCompleteWithdrawal(params.newMM);
            }

            {
                // How much can the operator withdraw by withdrawing the remainder of their deposit shares?
                uint remainingDepositShares = OPERATOR_SHARES - params.depositSharesToWithdraw;
                uint scaleQWD = scratchDSF.scaleForQueueWithdrawal(remainingDepositShares);

                params.withdrawRemaining = scaleQWD.scaleForCompleteWithdrawal(params.newMM);
            }

            {
                /// Calculate totals
                // 1. Amount exited by the operator
                params.completeExit = params.completeWithdrawal + params.withdrawRemaining;
                // 2. Amount withdrawn (redistribution + operator exit)
                params.totalWithdrawn = params.amountRedistributed + params.completeExit;
                require(params.totalWithdrawn <= OPERATOR_SHARES, "end > start");
                // 3. Amount missing after all's said and done
                params.totalMissing = OPERATOR_SHARES - params.totalWithdrawn;
            }

            // Error in amountRedistributed
            if (params.amountRedistributed < params.expectedRedistribution) {
                params.abs_Error = params.expectedRedistribution - params.amountRedistributed;
            } else {
                params.abs_Error = params.amountRedistributed - params.expectedRedistribution;
            }

            // Check results - if we have a new local min/max, push to results list
            // if (params.abs_Error > max_Amt) {
            //     max_Amt = params.abs_Error;
            //     points.push(params);
            // }

            // if (params.amountRedistributed > max_Amt) {
            //     max_Amt = params.amountRedistributed;
            //     points.push(params);
            // } 

            if (params.amountRedistributed < min_Amt) {
                min_Amt = params.amountRedistributed;
                points.push(params);
            } 
        }

        // Print results:
        for (uint i = 0; i < points.length; i++) {
            params = points[i];

            console.log("");
            console.log("(point %d)".dim().italic(), i);
            
            console.log(" - prevMM => newMM                   %s => %s".dim(), _pretty(params.prevMM), _pretty(params.newMM));
            console.log(" - prev | new                        %d | %d".dim(), params.prevMM, params.newMM);
            console.log(" - % in withdrawal queue:     %s".dim(), _prettyPerc(WAD - params.wadsInWQ, WAD));

            // Deposit info
            console.log(" - Deposit Shares: ".dim().yellow());
            uint remainingDepositShares = OPERATOR_SHARES - params.depositSharesToWithdraw;
            console.log(" -- queued deposit shares:    %s (%d)".dim().italic(), _pretty(params.depositSharesToWithdraw), params.depositSharesToWithdraw);
            console.log(" -- remaining deposit shares: %s (%d)".dim().italic(), _pretty(remainingDepositShares), remainingDepositShares);
            
            // Operator/Queue info
            console.log(" - Operator/Queue Shares: ".dim().yellow());
            console.log(" -- new operator shares:      %s (%d)".dim().italic(), _pretty(params.operatorShares), params.operatorShares);
            console.log(" -- scaledShares (queued):    %s (%d)".dim().italic(), _pretty(params.scaledShares), params.scaledShares);

            // Results:
            console.log(" - Results: ".dim().magenta());
            console.log(" -- expected redistribution:  %s (%d)".green(), _pretty(params.expectedRedistribution), params.expectedRedistribution); 
            console.log(" -- redistributed (avs):      %s (%d)".yellow(), _pretty(params.amountRedistributed), params.amountRedistributed);
            console.log(" -- withdrawn (operator):     %s (%d)".yellow(), _pretty(params.completeExit), params.completeExit);
            console.log(" -- total assets (end):       %s (%d)".yellow(), _pretty(params.totalWithdrawn), params.totalWithdrawn);
            // Where'd the rounding error go?
            int excessRedistribution = int(params.amountRedistributed) - int(params.expectedRedistribution);
            if (excessRedistribution > 0) {
                console.log(" -- avs gets extra:           %s (%d)".magenta(), _pretty(uint(excessRedistribution)), uint(excessRedistribution));
            } else {
                console.log(" -- dust lost:                %s (%d)".cyan(), _pretty(uint(-excessRedistribution)), uint(-excessRedistribution));
            }
            console.log(" -- missing assets:           %s (%d)".red(), _pretty(params.totalMissing), params.totalMissing);
            console.log(" -- absolute error:           %s (%d)".red(), _pretty(params.abs_Error), params.abs_Error);
        }

        console.log("");
        console.log("==Setup==".green());
        console.log(" - NUM_ITERATIONS:          %d", NUM_ITERATIONS);
        console.log(" - START ASSETS:            %s (%d)", _pretty(OPERATOR_SHARES), OPERATOR_SHARES);
        console.log(" - SLASH PERCENT:           %s", _prettyPerc(WAD - WADS_TO_SLASH, WAD));
        
        // console.log(" - EXPECTED REDISTRIBUTION: %s (%d)".green(), _pretty(expectedRedistribution), expectedRedistribution);

        revert("finished".green());
    }

    uint NUM_ITERATIONS = 0;
    uint GENERATE_METHOD = 0;
    uint C_NUM_SHARES = 0;
    uint C_MAG_DIFF = 0;

    struct MaxInput {        
        uint slashedFromOpShares;

        uint shares;
        uint64 prevMM;
        uint64 newMM;
    }

    function _randomS() internal returns (uint shares, uint diff) {
        shares = cheats.randomUint(1, WAD);

        diff = uint(WAD) / shares;
        require(shares * diff < uint(WAD), "condition not met");
    }

    function _randomD() internal returns (uint shares, uint diff) {
        diff = cheats.randomUint(1, WAD);

        shares = uint(WAD) / diff;
        require(shares * diff < uint(WAD), "condition not met");
    }

    function _contrivedS() internal returns (uint shares, uint diff) {
        shares = C_NUM_SHARES;

        diff = uint(WAD) / shares;
        require(shares * diff < uint(WAD), "condition not met");
    }

    function _contrivedD() internal returns (uint shares, uint diff) {
        diff = C_MAG_DIFF;

        shares = uint(WAD) / diff;
        if (shares * diff == uint(WAD)) shares--;

        require(shares * diff < uint(WAD), "condition not met");
    }

    function _genInputs() internal returns (uint, uint) {
        return
            GENERATE_METHOD == 0 ? _randomS() :
            GENERATE_METHOD == 1 ? _randomD() :
            GENERATE_METHOD == 2 ? _contrivedS() :
            GENERATE_METHOD == 3 ? _contrivedD() :
            (0, 0);
    }

    function test_Temp() public rand(0) {
        uint operatorShares = 3;
        uint withdrawalQueueShares = 7;

        uint64 prevMaxMag = 3;
        uint64 newMaxMag = 2;

        uint slashFromOpShares = operatorShares.calcSlashedAmount(prevMaxMag, newMaxMag);
        uint slashFromWQ = withdrawalQueueShares.scaleForBurning(prevMaxMag, newMaxMag);

        console.log("slashFromOP: %d", slashFromOpShares);
        console.log("slashFromWQ: %d", slashFromWQ);

        revert();
    }

    function test_Scratch() public rand(0) {
        _readScratchTOML();

        uint idx;
        MaxInput[] memory max = new MaxInput[](NUM_ITERATIONS);
        
        // Generate shares, prevMM, and newMM s.t.
        //
        // shares * (prevMM - newMM) < WAD
        //
        // Given scaleForBurning, this should guarantee that shares slashed from withdrawal queue 
        // are zero.
        //
        // ... then we check how much is slashed from the operator's shares in the same scenario.
        // The maximum slashed from the operator's shares is the maximum amount we can "avoid"
        // slashing by queueing a withdrawal.
        for (uint i = 0; i < NUM_ITERATIONS; i++) {
            (uint shares, uint diff) = _genInputs();

            // uint64 prevMM = WAD;
            // uint64 newMM = WAD - uint64(diff);

            uint64 prevMM = diff == WAD ? uint64(diff) : uint64(cheats.randomUint(diff, WAD));
            uint64 newMM = prevMM - uint64(diff);

            // uint64 prevMM = diff == 1 ? 1 : uint64(cheats.randomUint(1, diff));
            // uint64 newMM = prevMM == 1 ? 0 : uint64(cheats.randomUint(0, prevMM-1));

            uint slashedFromOpShares = shares.calcSlashedAmount(prevMM, newMM);
            uint slashedFromWQueue = shares.scaleForBurning(prevMM, newMM);
            require(slashedFromWQueue == 0, "slashed nonzero amt from wq");

            if (slashedFromOpShares > max[idx].slashedFromOpShares) {
                idx++;
                max[idx].slashedFromOpShares = slashedFromOpShares;
                max[idx].shares = shares;
                max[idx].prevMM = prevMM;
                max[idx].newMM = newMM;
            }
        }

        // Manually correct length before printing max inputs
        assembly { mstore(max, add(1, idx)) }

        for (uint i = 0; i < max.length; i++) {
            MaxInput memory m = max[i];
            uint64 diff = m.prevMM - m.newMM;

            console.log("");
            console.log("(# %d)".dim().italic(), i);
            console.log("Max slashed from op shares:  %s (%d)".yellow(), _pretty(m.slashedFromOpShares), m.slashedFromOpShares);
            console.log(" - input shares:             %s (%d)".dim(), _pretty(m.shares), m.shares);
            console.log(" - magnitude diff:           %s (%d)".dim(), _pretty(diff), diff);
            console.log(" -- prevMM:                  %s (%d)".dim().italic(), _pretty(m.prevMM), m.prevMM);
            console.log(" -- newMM:                   %s (%d)".dim().italic(), _pretty(m.newMM), m.newMM);
        }

        revert("finished".green().bold());
    }

    function _readExtractTOML() internal {
        string memory tomlRAW = cheats.readFile("src/test/integration/params.toml");

        NUM_ITERATIONS = cheats.parseTomlUint(tomlRAW, ".numIterations");
        // OPERATOR_SHARES = cheats.parseTomlUint(tomlRAW, ".operatorShares");
        OPERATOR_SHARES = (10 ** cheats.parseTomlUint(tomlRAW, ".operatorSharesExp"));
        WADS_TO_SLASH = cheats.parseTomlUint(tomlRAW, ".wadsToSlash");
    }

    function _readScratchTOML() internal {
        string memory tomlRAW = cheats.readFile("src/test/integration/params.toml");

        NUM_ITERATIONS = cheats.parseTomlUint(tomlRAW, ".numIterations");
        GENERATE_METHOD = cheats.parseTomlUint(tomlRAW, ".genMethod");
        C_NUM_SHARES = cheats.parseTomlUint(tomlRAW, ".contrivedShares");
        C_MAG_DIFF = cheats.parseTomlUint(tomlRAW, ".contrivedDiff");

        console.log("");
        console.log("==Setup==".green());
        console.log(" - NUM_ITERATIONS: %d", NUM_ITERATIONS);
        console.log(" - GENERATE_METHOD: %d", GENERATE_METHOD);

        if (GENERATE_METHOD == 2) {
            console.log(" -- CONTRIVED_SHARES:    %s (%d)", _pretty(C_NUM_SHARES), C_NUM_SHARES);
        } else if (GENERATE_METHOD == 3) {
            console.log(" -- CONTRIVED_DIFF:      %s (%d)", _pretty(C_MAG_DIFF), C_MAG_DIFF);
        }
    }

    struct DSFState {
        DepositScalingFactor dsf;
        uint depositShares;
        uint operatorMaxMag;

        uint prevScalingFactor;
        uint prevDepositShares;
        uint addedShares;
    }

    DSFState state;

    uint[] DSFs;

    enum Operation {
        DEPOSIT,
        WITHDRAW,
        REPEAT_PREV,
        REPEAT_PREV2
    }

    struct Action {
        uint amount;
        Operation op;
    }

    Action prev;
    Action prevprev;

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
        uint contIterations = _readFindOptimalTOML();

        for (uint i = 0; i < finalIteration; i++) {
            Action memory action = Action({
                op: Operation.DEPOSIT,
                amount: 1                
            });
            
            state.prevScalingFactor = state.dsf._scalingFactor;
            state.prevDepositShares = state.depositShares;

            _applyAction(action);

            // Iterate until decrease is around 50%
            uint scaledDecrease = 1e5 - ((1e5 * state.dsf._scalingFactor) / state.prevScalingFactor);
            if (scaledDecrease >= 49e3) {
                console.log("Peak found at %d iterations!".green().bold(), i+1);
                printInputs = true;
                finalIteration = i + printLast;
            }

            _print(action, i+1);
        }

        revert();
    }

    function test_DSF_FuzzOptimalSequence() public rand(0) {
        (
            DSFState memory startState,
            uint fuzzRuns,
            uint sequenceLength
        ) = _readFuzzSequenceTOML();

        uint minDSF_Overall = WAD;
        Action[] memory best = new Action[](sequenceLength);

        Action[] memory actions = new Action[](sequenceLength);
        uint finalActionIdx;

        for (uint i = 0; i < fuzzRuns; i++) {
            // One full run:
            
            // Reset start state
            state.dsf._scalingFactor = startState.dsf._scalingFactor;
            state.depositShares = startState.depositShares;
            state.operatorMaxMag = startState.operatorMaxMag;

            state.prevScalingFactor = startState.prevScalingFactor;
            state.prevDepositShares = startState.prevDepositShares;

            // Track lowest DSF within a given run, in case the final iterations kill a good run
            uint minDSF_WithinRun = state.dsf._scalingFactor;
            uint minDSF_Location = 0;

            for (uint j = 0; j < actions.length; j++) {
                state.prevScalingFactor = state.dsf._scalingFactor;
                state.prevDepositShares = state.depositShares;

                // Select [DEPOSIT,WITHDRAW]
                // - Start with a deposit
                // - Never withdraw twice in a row (pointless expansion of search space)
                // - Only withdraw up to the amount previously deposited - 1
                //
                // TODO - consider making withdraw a 1:4?
                // I have no idea what ratio would be good... maybe we do exhaustive rather than fuzzing?
                actions[j].op = Operation.DEPOSIT;
                // if (j == 0 || prev.amount >= 2 || prev.op == Operation.WITHDRAW) {
                //     actions[j].op = Operation.DEPOSIT;
                // } else {
                //     actions[j].op = cheats.randomBool() ? Operation.DEPOSIT : Operation.WITHDRAW;
                // }

                // Select AMOUNT
                // - deposit amounts less than sequence length
                // - withdraw amounts less than total balance
                actions[j].amount = cheats.randomUint(1, 2);
                // if (actions[j].op == Operation.DEPOSIT) {
                //     actions[j].amount = sequenceLength == 1 ? 1 : cheats.randomUint(1, sequenceLength-j);
                // } else {
                //     actions[j].amount = prev.amount == 1 ? 1 : cheats.randomUint(1, prev.amount-1);
                // }

                // Apply selected action, updating DSF
                _applyAction(actions[j]);

                if (state.dsf._scalingFactor < minDSF_WithinRun) {
                    minDSF_WithinRun = state.dsf._scalingFactor;
                    minDSF_Location = j;
                }

                prev = actions[j];
            }

            // If we found a new low, update
            if (minDSF_WithinRun < minDSF_Overall) {
                minDSF_Overall = minDSF_WithinRun;
                console.log("Found new minimum DSF: %s (%d)".dim(), _pretty(minDSF_Overall), minDSF_Overall);
                console.log("- Iterations: %d".dim().italic(), minDSF_Location+1);
                _copySeq({ from: actions, to: best, finalActionIdx: minDSF_Location });
                finalActionIdx = minDSF_Location;
                // _printSeq(best, minDSF_Location, minDSF_WithinRun);
            } else if (minDSF_WithinRun == minDSF_Overall && minDSF_Location < finalActionIdx) {
                minDSF_Overall = minDSF_WithinRun;
                console.log("Found faster DSF drop for DSF: %s (%d)".dim(), _pretty(minDSF_Overall), minDSF_Overall);
                console.log("- (%d => %d)".dim().italic(), finalActionIdx+1, minDSF_Location+1);
                _copySeq({ from: actions, to: best, finalActionIdx: minDSF_Location });
                finalActionIdx = minDSF_Location;
                // _printSeq(best, minDSF_Location, minDSF_WithinRun);
            }
        }

        _printSeq(best, finalActionIdx, minDSF_Overall);

        revert("finished");
    }

    function _printSeq(Action[] memory actions, uint finalActionIdx, uint minDSF) internal {
        console.log("==Best Sequence==".yellow());
        console.log("- iterations: %d".dim(), finalActionIdx+1);
        console.log("- DSF:               %s (%d)", _pretty(minDSF), minDSF);
        for (uint i = 0; i <= finalActionIdx; i++) {
            _logDepositOrWithdraw(actions[i]);
        }
    }

    function _copySeq(Action[] memory from, Action[] memory to, uint finalActionIdx) internal pure {
        require(from.length == to.length && finalActionIdx < from.length);
        for (uint i = 0; i <= finalActionIdx; i++) {
            to[i].amount = from[i].amount;
            to[i].op = from[i].op;
        }
    }

    function _applyAction(Action memory action) internal {
        if (action.op == Operation.DEPOSIT) {
            state.dsf.update({
                prevDepositShares: state.depositShares,
                addedShares: action.amount,
                slashingFactor: state.operatorMaxMag
            });

            state.depositShares += action.amount;
        } else if (action.op == Operation.WITHDRAW) {
            state.depositShares -= action.amount;
        }
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

        // console.log("prevDepositShares: %d", state.prevDepositShares);
        state.prevDepositShares += state.addedShares;
        state.operatorMaxMag -= maxMagDecrease;

        bool peakFound;
        for (uint i = startIteration; i < startIteration+contIterations; i++) {
            _print(i);
            if (peakFound) printInputs = false;

            state.addedShares = 1;
            state.prevScalingFactor = state.dsf._scalingFactor;

            // console.log("prevDepositShares: %d", state.prevDepositShares);

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

    function test_DSF_Sequence() public rand(0) {
        Action[] memory actions = _readSequenceTOML();

        uint iterations = 0;
        for (uint i = 0; i < actions.length; i++) {
            Action memory mainAction = actions[i];
            if (i == 0) {
                require(mainAction.op != Operation.REPEAT_PREV && mainAction.op != Operation.REPEAT_PREV2);
            } else if (i == 1) {
                require(mainAction.op != Operation.REPEAT_PREV2);
            }

            // number of times to apply action
            uint applyTimes = 
                mainAction.op == Operation.REPEAT_PREV ? mainAction.amount :
                mainAction.op == Operation.REPEAT_PREV2 ? 2 * mainAction.amount :
                1;

            state.prevScalingFactor = state.dsf._scalingFactor;
            state.prevDepositShares = state.depositShares;

            for (uint j = 0; j < applyTimes; j++) {
                // Cursed switch syntax determines if we're doing the main action or repeating a prev
                Action memory action =
                    mainAction.op == Operation.REPEAT_PREV ? prev :
                    mainAction.op == Operation.REPEAT_PREV2 ? 
                        (j % 2 == 0 ? prevprev : prev) :
                    mainAction;
                
                _applyAction(action);

                
            }

            iterations += applyTimes;
            _print(mainAction, iterations);

            prevprev = prev;
            prev = mainAction;
        }

        revert();
    }

    function _readSequenceTOML() internal returns (Action[] memory actions) {
        string memory tomlRAW = cheats.readFile("src/test/integration/params.toml");
        bytes memory data = vm.parseToml(tomlRAW, ".sequence");
        actions = abi.decode(data, (Action[]));

        // Also setup start state
        state.dsf._scalingFactor = cheats.parseTomlUint(tomlRAW, ".seqStartDSF");
        state.depositShares = cheats.parseTomlUint(tomlRAW, ".seqStartShares");
        state.operatorMaxMag = WAD - 1;

        state.prevScalingFactor = state.dsf._scalingFactor;
        state.prevDepositShares = state.depositShares;

        printInputs = true;
        finalIteration = actions.length;

        _printStartState();

        return actions;
    }

    function _readTOML() internal returns (uint, uint) {
        string memory tomlRAW = cheats.readFile("src/test/integration/params.toml");

        return (
            cheats.parseTomlUint(tomlRAW, ".prevDepositShares"),
            cheats.parseTomlUint(tomlRAW, ".iterations")
        );
    }

    function _readFindOptimalTOML() internal returns (uint) {
        string memory tomlRAW = cheats.readFile("src/test/integration/params.toml");

        // Setup start state
        state.dsf._scalingFactor = cheats.parseTomlUint(tomlRAW, ".startDSF");
        state.depositShares = cheats.parseTomlUint(tomlRAW, ".startShares");
        state.operatorMaxMag = WAD - 1;

        state.prevScalingFactor = state.dsf._scalingFactor;
        state.prevDepositShares = state.depositShares;

        printInputs = true;
        finalIteration = cheats.parseTomlUint(tomlRAW, ".contIterations");

        _printStartState();

        return (
            finalIteration
        );
    }

    function _readFuzzSequenceTOML() internal returns (DSFState memory startState, uint, uint) {
        string memory tomlRAW = cheats.readFile("src/test/integration/params.toml");

        // Setup start state
        startState.dsf._scalingFactor = cheats.parseTomlUint(tomlRAW, ".seqStartDSF");
        startState.depositShares = cheats.parseTomlUint(tomlRAW, ".seqStartShares");
        startState.operatorMaxMag = WAD - 1;

        startState.prevScalingFactor = startState.dsf._scalingFactor;
        startState.prevDepositShares = startState.depositShares;

        // Reset start state
        state.dsf._scalingFactor = startState.dsf._scalingFactor;
        state.depositShares = startState.depositShares;
        state.operatorMaxMag = startState.operatorMaxMag;

        state.prevScalingFactor = startState.prevScalingFactor;
        state.prevDepositShares = startState.prevDepositShares;

        _printStartState();

        // printInputs = true;
        // finalIteration = cheats.parseTomlUint(tomlRAW, ".contIterations");
        // _printStartState();

        return (
            startState,
            cheats.parseTomlUint(tomlRAW, ".fuzzRuns"),
            cheats.parseTomlUint(tomlRAW, ".sequenceLength")
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

    uint finalIteration = type(uint).max;
    bool printInputs = false;

    function _printStartState() internal {
        console.log("");
        console.log("==StartState==".yellow());

        console.log(" - startDSF:              %s (%d)".dim(), _pretty(state.dsf._scalingFactor), state.dsf._scalingFactor);
        console.log(" - startDepositShares:    %d".dim(), state.prevDepositShares);
        console.log(" - startOpMaxMag:         %s (%d)".dim(), _pretty(state.operatorMaxMag), state.operatorMaxMag);

        console.log("");
    }

    function _print(Action memory action, uint iter) internal {
        _printIteration(action, iter);  
        
        if (action.op == Operation.DEPOSIT) {
            if (printInputs) {
                console.log("==Inputs==".yellow());

                console.log(" - prevDepositShares:     %s (%d)".dim(), _pretty(state.prevDepositShares), state.prevDepositShares);
                console.log(" - addedShares:           %d".dim(), action.amount);
                console.log(" - prevDSF:               %s (%d)".dim(), _pretty(state.prevScalingFactor), state.prevScalingFactor);

                console.log("=currentShares=".yellow().dim());

                uint pDS = state.prevDepositShares;
                uint pDSMulDSF = pDS * state.prevScalingFactor;
                uint numerator = pDS.mulWad(state.prevScalingFactor) * state.operatorMaxMag;
                uint currentShares = pDS.mulWad(state.prevScalingFactor).mulWad(state.operatorMaxMag);

                console.log(" - prevDepositShares*DSF: %s (%d)".dim(), _pretty(pDSMulDSF), pDSMulDSF);
                console.log(" - numerator:             %s (%d)".dim(), _pretty(numerator), numerator);
                console.log(" - currentShares:         %d".dim(), currentShares);
            }

            _prettyDSFDelta();
        } else if (action.op == Operation.WITHDRAW) {
            
        } else {
            _prettyDSFDelta();
        }

        console.log("=newDepositShares=".yellow().dim());
        console.log(" - depositShares:         %s (%d)".dim(), _pretty(state.depositShares), state.depositShares);

        console.log("");
    }

    // Ex:
    // [iter: 5] 
    // - deposited shares:   0e01 (1)
    function _printIteration(Action memory action, uint iter) internal {
        console.log("[iter: %d]".dim(), iter);

        if (action.op == Operation.DEPOSIT) {
            _logDepositOrWithdraw(action);
        } else if (action.op == Operation.WITHDRAW) {
            _logDepositOrWithdraw(action);
        } else if (action.op == Operation.REPEAT_PREV) {
            console.log("Repeated last action %d times:".dim(), action.amount);
            _logDepositOrWithdraw(prev);
        } else if (action.op == Operation.REPEAT_PREV2) {
            console.log("Repeated last two actions %d times:".dim(), action.amount);
            _logDepositOrWithdraw(prevprev);
            _logDepositOrWithdraw(prev);
        }
    }

    function _logDepositOrWithdraw(Action memory action) internal {
        if (action.op == Operation.DEPOSIT) {
            console.log(
                "-- deposited shares:       %s (%d)".dim().italic(),
                _pretty(action.amount), 
                action.amount
            );
        } else if (action.op == Operation.WITHDRAW) {
            console.log(
                "-- withdrew shares:       %s (%d)".dim().italic(),
                _pretty(action.amount), 
                action.amount
            );
        }
    }

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
            console.log(" - prevDepositShares:     %d".dim(), state.prevDepositShares - state.addedShares);
            // console.log(" - prevDepositShares:     %d".dim(), state.prevDepositShares);
            console.log(" - addedShares:           %d".dim(), state.addedShares);
            console.log(" - prevDSF:               %s (%d)".dim(), _pretty(state.prevScalingFactor), state.prevScalingFactor);

            {
                console.log("=currentShares=".yellow().dim());

                uint pDS = state.prevDepositShares - state.addedShares;
                // uint pDS = state.prevDepositShares;
                uint pDSMulDSF = pDS * state.prevScalingFactor;
                uint numerator = pDS.mulWad(state.prevScalingFactor) * state.operatorMaxMag;
                uint currentShares = pDS.mulWad(state.prevScalingFactor).mulWad(state.operatorMaxMag);

                console.log(" - prevDepositShares*DSF: %s (%d)".dim(), _pretty(pDSMulDSF), pDSMulDSF);
                console.log(" - numerator:             %s (%d)".dim(), _pretty(numerator), numerator);
                console.log(" - currentShares:         %d".dim(), currentShares);
            }
        } 
        
        _prettyDSFDelta();
    }

    // Print DSF as red under 1e12, yellow under 1e15, green otherwise
    // Print decrease/increase/no change with %
    function _prettyDSFDelta() internal {
        console.log("=newDSF=".yellow().dim());
        
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
}