// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_Rounding is IntegrationCheckUtils {
    using ArrayLib for *;
    using StdStyle for *;

    User attacker;
    AVS badAVS;
    IStrategy strategy;
    IERC20Metadata token;
    User goodStaker;
    uint64 initTokenBalance;

    OperatorSet mOpSet; // "manipOpSet" used for magnitude manipulation
    OperatorSet rOpSet; // Redistributable opset used to exploit precision loss and trigger redistribution

    uint16 numSlashes = 100; // TODO: transform into envvar
    
    // CSV file tracking
    string private csvPath;
    uint256 private testRunId;
    bool private csvInitialized;
    
    // Struct to hold CSV data to avoid stack too deep
    struct CSVData {
        uint256 tokenBalance;
        uint256 depositShares;
        uint256 withdrawableShares;
        uint256 operatorShares;
        uint256 maxMag;
        uint256 totalAllocated;
        uint256 totalAvailable;
        uint256 allocatedMOpSet;
        uint256 allocatedROpSet;
        uint256 dsf;
    }

    function _init() internal override {
        _configAssetTypes(HOLDS_LST);

        attacker = new User("Attacker"); // Attacker serves as both operator and staker
        badAVS = new AVS("BadAVS"); // AVS is also attacker-controlled
        strategy = lstStrats[0];
        token = IERC20Metadata(address(strategy.underlyingToken()));

        // Prepares to add non-attacker stake into the protocol. Can be any amount > 0.
        // Note that the honest stake does not need to be allocated anywhere, so long as it's in the same strategy.
        goodStaker = new User("GoodStaker");
        deal(address(token), address(goodStaker), uint256(1e18));
        goodStaker.depositIntoEigenlayer(strategy.toArray(), 1e18.toArrayU256());

        // Register attacker as operator and create attacker-controlled AVS/OpSets
        attacker.registerAsOperator(0);
        rollForward({blocks: ALLOCATION_CONFIGURATION_DELAY + 1});
        badAVS.updateAVSMetadataURI("https://example.com");
        mOpSet = badAVS.createOperatorSet(strategy.toArray()); // setup low mag operator
        rOpSet = badAVS.createRedistributingOperatorSet(strategy.toArray(), address(attacker)); // execute exploit

        // Register for both opsets
        attacker.registerForOperatorSet(mOpSet);
        attacker.registerForOperatorSet(rOpSet);

        // Initialize CSV tracking
        _initializeCSV();

        _writePhaseData("setup", 0, 0);
    }

    function _initializeCSV() private {
        // Check if CSV filename is provided via environment variable
        try vm.envString("CSV_FILENAME") returns (string memory customFilename) {
            csvPath = customFilename;
        } catch {
            // Generate unique filename with timestamp and random seed
            testRunId = uint256(keccak256(abi.encodePacked(block.timestamp, block.prevrandao)));
            csvPath = string(abi.encodePacked(
                "rounding_test_results_",
                vm.toString(testRunId % 1000000), // Use last 6 digits for shorter filename
                ".csv"
            ));
        }
        
        // Write CSV headers
        string memory headers = "test_name,phase,initial_max_mag,init_token_balance,token_balance,deposit_shares,withdrawable_shares,operator_shares,init_mag,max_mag,total_allocated,total_available,allocated_mOpSet,allocated_rOpSet,dsf,slash_count,precision_loss,test_result";
        vm.writeLine(csvPath, headers);
        csvInitialized = true;
    }

    /**
     *
     *                         TESTS
     *
     */

    // TODO: consider incremental manual fuzzing from 1 up to WAD - 1
    /**
     * @notice Tests rounding behavior when all operator magnitudes are slashed
     * @param _initialMaxMag The initial maximum magnitude to set for the operator, bounded between 1 and WAD-1
     * @param _initTokenBalance The initial token balance to give the attacker, must be > 0
     * @dev This test verifies that when an operator's magnitude is gradually reduced through multiple slashes
     * and finally completely slashed, there is no precision loss in the redistribution of tokens.
     */
    function test_rounding_allMagsSlashed(
        uint64 _initialMaxMag,
        uint64 _initTokenBalance
    ) public rand(0) {
        vm.pauseGasMetering();
        // Don't slash 100% as we will do multiple slashes
        _initialMaxMag = uint64(bound(_initialMaxMag, 1, WAD - 1));
        // Ensure attacker has at least one token
        initTokenBalance = _initTokenBalance > 0 ? _initTokenBalance : _initTokenBalance + 1;
        deal(address(token), address(attacker), initTokenBalance);

        // Use modifyAllocation+slashOperator to arbitrarily set operator max magnitude
        _magnitudeManipulation(_initialMaxMag);

        // Deposit all attacker assets into Eigenlayer.
        _depositAll();

        // Perform slashes to gradually whittle down operator magnitude, as well as produce slash escrows.
        // Since we're doing multiple slashes, never slash 100%.
        for (uint16 i = 0; i < numSlashes; i++) {
            uint64 wadToSlash = uint64(cheats.randomUint(1, WAD - 1));
            _slash(wadToSlash);
        }

        // Perform final 100% slash to extract any remaining tokens.
        _slash(WAD);

        // Release all escrows to the redistributionRecipient (attacker).
        _release();

        // Check for precision loss.
        // Note: No precision loss expected after all magnitude is slashed.
        (bool success, uint256 actualLoss) = _checkForPrecisionLoss(0);
        
        // Write final result to CSV
        _writeTestResult("test_rounding_allMagsSlashed", "final", _initialMaxMag, actualLoss, success ? "PASS" : "FAIL");
    }

    /**
     * @notice Tests rounding behavior when operator magnitudes are partially slashed
     * @param _initialMaxMag The initial maximum magnitude to set for the operator, bounded between 1 and WAD-1
     * @param _initTokenBalance The initial token balance to give the attacker, must be > 0
     * @dev This test verifies that when an operator's magnitude is gradually reduced through multiple slashes
     * and finally completely slashed, there is minimal precision loss.
     */
    function test_rounding_partialMagsSlashed(
        uint64 _initialMaxMag,
        uint64 _initTokenBalance
    ) public rand(0) {
        vm.pauseGasMetering();
        // Don't slash 100% as we will do multiple slashes
        _initialMaxMag = uint64(bound(_initialMaxMag, 1, WAD - 1));
        // Ensure attacker has at least one token
        initTokenBalance = _initTokenBalance > 0 ? _initTokenBalance : _initTokenBalance + 1;
        deal(address(token), address(attacker), initTokenBalance);

        // Use modifyAllocation+slashOperator to arbitrarily set operator max magnitude
        _magnitudeManipulation(_initialMaxMag);

        // Deposit all attacker assets into Eigenlayer.
        _depositAll();

        // Perform slashes to gradually whittle down operator magnitude, as well as produce slash escrows.
        // Since we're doing multiple slashes, never slash 100%.
        for (uint16 i = 0; i < numSlashes; i++) {
            uint64 wadToSlash = uint64(cheats.randomUint(1, WAD - 1));
            _slash(wadToSlash);
        }

        // Release all escrows to the redistributionRecipient (attacker).
        _release();

        // Withdraw all attacker deposits. Necessary as operator has partial mags remaining.
        _withdraw();

        // Note: Precision loss seems to be a consequence of the DSF, rather than slashing precision loss.
        // Max precision loss for this test is observed to correspond to residual operator shares.
        // TODO: Explore root cause of this precision loss.
        uint operatorShares = delegationManager.getOperatorShares(address(attacker), strategy.toArray())[0];
        (bool success, uint256 actualLoss) = _checkForPrecisionLoss(operatorShares);
        
        // Write final result to CSV
        _writeTestResult("test_rounding_partialMagsSlashed", "final", _initialMaxMag, actualLoss, success ? "PASS" : "FAIL");
    }

    // TODO: consider parameterizing "honest" stake to see how strategy exchange rate is affected
    /**
     * @notice Tests rounding behavior when the attacker manipulates the strategy's exchange rate with all magnitudes slashed
     * @param _initialMaxMag The initial maximum magnitude to set for the operator, bounded between 1 and WAD-1
     * @param _initTokenBalance The initial token balance to give the attacker, must be > 0
     * @dev This test verifies that when the attacker manipulates the strategy's exchange rate, by depositing funds directly into the strategy, that the attacker cannot profit from the manipulation, even when all magnitudes are slashed.
     */
    function test_rounding_strategySharesManipulation_allMagsSlashed(
        uint64 _initialMaxMag,
        uint64 _initTokenBalance
    ) public rand(0) {
        vm.pauseGasMetering();
        // Don't slash 100% initially as we will do multiple slashes
        _initialMaxMag = uint64(bound(_initialMaxMag, 1, WAD - 1));
        // Ensure attacker has at least one token
        initTokenBalance = _initTokenBalance > 0 ? _initTokenBalance : _initTokenBalance + 1;
        deal(address(token), address(attacker), initTokenBalance);

        // Use modifyAllocation+slashOperator to arbitrarily set operator max magnitude
        _magnitudeManipulation(_initialMaxMag);

        // Allocate all remaining magnitude to redistributable opset.
        uint64 allocatableMagnitude = allocationManager.getAllocatableMagnitude(address(attacker), strategy);
        attacker.modifyAllocations(AllocateParams({
            operatorSet: rOpSet,
            strategies: strategy.toArray(),
            newMagnitudes: (allocatableMagnitude).toArrayU64()
        }));

        // Deposit a random amount of tokens into the strategy.
        attacker.depositIntoEigenlayer(strategy.toArray(), cheats.randomUint(1, token.balanceOf(address(attacker))).toArrayU256());

        // Perform slashes to gradually whittle down operator magnitude, as well as produce slash escrows.
        // Since we're doing multiple slashes, never slash 100%.
        for (uint16 i = 0; i < numSlashes; i++) {
            uint64 wadToSlash = uint64(cheats.randomUint(1, WAD - 1));
            // Randomly decide whether to add funds to the strategy between slashes.
            if (cheats.randomBool() && token.balanceOf(address(attacker)) > 0) {
                uint fundsToAdd = cheats.randomUint(1, token.balanceOf(address(attacker)));
                // Deal funds directly to the strategy to affect the exchange rate.
                vm.prank(address(attacker));
                token.transfer(address(strategy), fundsToAdd);
            }
            _slash(wadToSlash);
        }

        // Perform final 100% slash to extract any remaining tokens.
        _slash(WAD);

        // Release all escrows to the redistributionRecipient (attacker).
        _release();

        // Note: In this test case, the attacker burns funds to attempt to manipulate the strategy's exchange rate. As such, severe token loss is expected.
        (bool success, uint256 actualLoss) = _checkForPrecisionLoss(initTokenBalance);
        
        // Write final result to CSV
        _writeTestResult("test_rounding_strategySharesManipulation_allMagsSlashed", "final", _initialMaxMag, actualLoss, success ? "PASS" : "FAIL");
    }

    /**
     * @notice Tests rounding behavior when the attacker manipulates the strategy's exchange rate
     * @param _initialMaxMag The initial maximum magnitude to set for the operator, bounded between 1 and WAD-1
     * @param _initTokenBalance The initial token balance to give the attacker, must be > 0
     * @dev This test verifies that when the attacker manipulates the strategy's exchange rate, by depositing funds directly into the strategy, that the attacker cannot profit from the manipulation.
     */
    function test_rounding_strategySharesManipulation_partialMagsSlashed(
        uint64 _initialMaxMag,
        uint64 _initTokenBalance
    ) public rand(0) {
        vm.pauseGasMetering();
        // Don't slash 100% as we will do multiple slashes
        _initialMaxMag = uint64(bound(_initialMaxMag, 1, WAD - 1));
        // Ensure attacker has at least one token
        initTokenBalance = _initTokenBalance > 0 ? _initTokenBalance : _initTokenBalance + 1;
        deal(address(token), address(attacker), initTokenBalance);

        // Use modifyAllocation+slashOperator to arbitrarily set operator max magnitude
        _magnitudeManipulation(_initialMaxMag);

        // Allocate all remaining magnitude to redistributable opset.
        uint64 allocatableMagnitude = allocationManager.getAllocatableMagnitude(address(attacker), strategy);
        attacker.modifyAllocations(AllocateParams({
            operatorSet: rOpSet,
            strategies: strategy.toArray(),
            newMagnitudes: (allocatableMagnitude).toArrayU64()
        }));

        // Deposit a random amount of tokens into the strategy.
        attacker.depositIntoEigenlayer(strategy.toArray(), cheats.randomUint(1, token.balanceOf(address(attacker))).toArrayU256());

        // Perform slashes to gradually whittle down operator magnitude, as well as produce slash escrows.
        // Since we're doing multiple slashes, never slash 100%.
        for (uint16 i = 0; i < numSlashes; i++) {
            uint64 wadToSlash = uint64(cheats.randomUint(1, WAD - 1));
            // Randomly decide whether to add funds to the strategy between slashes.
            if (cheats.randomBool() && token.balanceOf(address(attacker)) > 0) {
                uint fundsToAdd = cheats.randomUint(1, token.balanceOf(address(attacker)));
                // Deal funds directly to the strategy to affect the exchange rate.
                vm.prank(address(attacker));
                token.transfer(address(strategy), fundsToAdd);
            }
            _slash(wadToSlash);
        }

        // Release all escrows to the redistributionRecipient (attacker).
        _release();

        // Withdraw all attacker deposits. Necessary as operator has partial mags remaining.
        _withdraw();

        // Note: In this test case, the attacker burns funds to attempt to manipulate the strategy's exchange rate. As such, severe token loss is expected.
        (bool success, uint256 actualLoss) = _checkForPrecisionLoss(initTokenBalance);
        
        // Write final result to CSV
        _writeTestResult("test_rounding_strategySharesManipulation_partialMagsSlashed", "final", _initialMaxMag, actualLoss, success ? "PASS" : "FAIL");
    }

    /**
     *
     *                         INTERNAL FUNCTIONS
     *
     */

    // @notice Check for any precision loss.
    //         First case means attacker gained money. If found, we've proven the existence of an exploit.
    //         Second case means attacker lost money. This demonstrates precision loss, with maxLoss as the upper bound.
    // @dev Returns (success, actualLoss) where success indicates test passed
    function _checkForPrecisionLoss(uint256 maxLoss) internal returns (bool success, uint256 actualLoss) {
        if (token.balanceOf(address(attacker)) > initTokenBalance) {
            uint64 diff = uint64(token.balanceOf(address(attacker))) - initTokenBalance;
            actualLoss = 0; // Actually a gain
            success = false;
            // ANY tokens gained is an exploit.
            revert("Rounding error exploit found!");
        } else if (token.balanceOf(address(attacker)) < initTokenBalance) {
            uint64 diff = uint64(initTokenBalance - token.balanceOf(address(attacker)));
            actualLoss = diff;
            success = diff <= maxLoss;
            // Check against provided tolerance.
            assertLe(diff, maxLoss, "Tokens lost!");
        } else {
            actualLoss = 0;
            success = true;
        }
    }

    // TODO - another way to mess with rounding/precision loss is to manipulate DSF
    function _magnitudeManipulation(uint64 _initialMaxMag) internal {
        // Allocate all magnitude to operator set.
        attacker.modifyAllocations(AllocateParams({
            operatorSet: mOpSet,
            strategies: strategy.toArray(),
            newMagnitudes: WAD.toArrayU64()
        }));

        _writePhaseData("allocate", _initialMaxMag, 0);

        uint64 wadToSlash = WAD - _initialMaxMag;
        // Slash operator to arbitrary mag.
        badAVS.slashOperator(SlashingParams({
            operator: address(attacker),
            operatorSetId: mOpSet.id,
            strategies: strategy.toArray(),
            wadsToSlash: uint(wadToSlash).toArrayU256(),
            description: "manipulation!"
        }));

        _writePhaseData("slash", _initialMaxMag, 1);

        // Deallocate magnitude from operator set.
        attacker.modifyAllocations(AllocateParams({
            operatorSet: mOpSet,
            strategies: strategy.toArray(),
            newMagnitudes: 0.toArrayU64()
        }));

        rollForward({blocks: DEALLOCATION_DELAY + 1});

        _writePhaseData("deallocate", _initialMaxMag, 0);
    }

    function _depositAll() internal {
        // Allocate all remaining magnitude to redistributable opset.
        uint64 allocatableMagnitude = allocationManager.getAllocatableMagnitude(address(attacker), strategy);
        attacker.modifyAllocations(AllocateParams({
            operatorSet: rOpSet,
            strategies: strategy.toArray(),
            newMagnitudes: (allocatableMagnitude).toArrayU64()
        }));

        // Deposit all attacker assets into Eigenlayer.
        attacker.depositIntoEigenlayer(strategy.toArray(), token.balanceOf(address(attacker)).toArrayU256());

        _writePhaseData("depositAll", 0, 0);
    }

    function _slash(uint64 wadToSlash) internal {
        // Perform final slash on redistributable opset and check for profit. Slash all operator magnitude.
        badAVS.slashOperator(SlashingParams({
            operator: address(attacker),
            operatorSetId: rOpSet.id,
            strategies: strategy.toArray(),
            wadsToSlash: uint(wadToSlash).toArrayU256(),
            description: "final slash"
        }));

        _writePhaseData("slash", 0, 0);
    }

    function _release() internal {
        // Roll forward past the escrow delay.
        rollForward({blocks: slashEscrowFactory.getGlobalEscrowDelay() + 1});

        // Release funds.
        for (uint32 i = 1; i <= numSlashes; i++) {
            vm.prank(address(attacker));
            slashEscrowFactory.releaseSlashEscrow(rOpSet, i);
        }

        // Release final escrow.
        vm.prank(address(attacker));
        slashEscrowFactory.releaseSlashEscrow(rOpSet, uint256(numSlashes) + 1);

        _writePhaseData("release", 0, 0);
    }

    function _withdraw() internal {
        (, uint256[] memory depositShares) = strategyManager.getDeposits(address(attacker));

        Withdrawal[] memory withdrawals = attacker.queueWithdrawals(strategy.toArray(), depositShares);

        rollForward({blocks: DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS + 1});

        attacker.completeWithdrawalsAsTokens(withdrawals);

        _writePhaseData("withdraw", 0, 0);
    }
    
    function _collectCSVData() internal returns (CSVData memory data) {
        address a = address(attacker);
        
        // Get share data
        (uint[] memory withdrawableArr, uint[] memory depositArr) = delegationManager.getWithdrawableShares(a, strategy.toArray());
        data.withdrawableShares = withdrawableArr.length == 0 ? 0 : withdrawableArr[0];
        data.depositShares = depositArr.length == 0 ? 0 : depositArr[0];
        data.operatorShares = delegationManager.operatorShares(a, strategy);
        
        // Get allocation data
        Allocation memory mAlloc = allocationManager.getAllocation(a, mOpSet, strategy);
        Allocation memory rAlloc = allocationManager.getAllocation(a, rOpSet, strategy);
        data.allocatedMOpSet = mAlloc.currentMagnitude;
        data.allocatedROpSet = rAlloc.currentMagnitude;
        
        // Get other data
        data.tokenBalance = token.balanceOf(a);
        data.maxMag = allocationManager.getMaxMagnitude(a, strategy);
        data.totalAllocated = allocationManager.getEncumberedMagnitude(a, strategy);
        data.totalAvailable = allocationManager.getAllocatableMagnitude(a, strategy);
        data.dsf = delegationManager.depositScalingFactor(a, strategy);
    }

    function _writePhaseData(string memory phaseName, uint64 _initialMaxMag, uint256 slashCount) internal {
        if (!csvInitialized) return;
        
        // Collect all data in struct
        CSVData memory data = _collectCSVData();
        
        // Write row part 1
        string memory row = string(abi.encodePacked(
            "current_test,",
            phaseName, ",",
            vm.toString(_initialMaxMag), ",",
            vm.toString(initTokenBalance), ","
        ));
        
        // Write row part 2
        row = string(abi.encodePacked(
            row,
            vm.toString(data.tokenBalance), ",",
            vm.toString(data.depositShares), ",",
            vm.toString(data.withdrawableShares), ",",
            vm.toString(data.operatorShares), ","
        ));
        
        // Write row part 3
        row = string(abi.encodePacked(
            row,
            vm.toString(WAD), ",",
            vm.toString(data.maxMag), ",",
            vm.toString(data.totalAllocated), ",",
            vm.toString(data.totalAvailable), ","
        ));
        
        // Write row part 4
        row = string(abi.encodePacked(
            row,
            vm.toString(data.allocatedMOpSet), ",",
            vm.toString(data.allocatedROpSet), ",",
            vm.toString(data.dsf), ",",
            vm.toString(slashCount), ",",
            "0,RUNNING"
        ));
        
        vm.writeLine(csvPath, row);
    }
    
    function _writeTestResult(
        string memory testName, 
        string memory phase, 
        uint64 _initialMaxMag, 
        uint256 precisionLoss, 
        string memory result
    ) internal {
        if (!csvInitialized) return;
        
        // Collect all data in struct
        CSVData memory data = _collectCSVData();
        
        // Write row in parts
        string memory row = string(abi.encodePacked(
            testName, ",",
            phase, ",",
            vm.toString(_initialMaxMag), ",",
            vm.toString(initTokenBalance), ","
        ));
        
        row = string(abi.encodePacked(
            row,
            vm.toString(data.tokenBalance), ",",
            vm.toString(data.depositShares), ",",
            vm.toString(data.withdrawableShares), ",",
            vm.toString(data.operatorShares), ","
        ));
        
        row = string(abi.encodePacked(
            row,
            vm.toString(WAD), ",",
            vm.toString(data.maxMag), ",",
            vm.toString(data.totalAllocated), ",",
            vm.toString(data.totalAvailable), ","
        ));
        
        row = string(abi.encodePacked(
            row,
            "0,0,", // mOpSet and rOpSet allocations (likely 0 at end)
            vm.toString(data.dsf), ",",
            vm.toString(numSlashes + 1), ",",
            vm.toString(precisionLoss), ",",
            result
        ));
        
        vm.writeLine(csvPath, row);
    }
}