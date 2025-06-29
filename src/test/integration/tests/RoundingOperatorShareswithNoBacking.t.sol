// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_Rounding is IntegrationCheckUtils {
    using ArrayLib for *;
    using StdStyle for *;

    string logFilename;  //  Store filename
    bool headerWritten = false;  //  Add flag to write header only once
    User attacker;
    // Number of users to create
    User[4] users;
    AVS badAVS;
    IStrategy strategy;
    IERC20Metadata token;
    User goodStaker;
    uint64 initTokenBalance;

    OperatorSet mOpSet; // "manipOpSet" used for magnitude manipulation
    OperatorSet rOpSet; // Redistributable opset used to exploit precision loss and trigger redistribution
    
    uint numSlashes = 1;

    function _init() internal override {
        _configAssetTypes(HOLDS_LST);
        //  Generate unique filename but don't write header in _init
        uint256 uniqueId = uint256(keccak256(abi.encodePacked(block.timestamp, block.difficulty, gasleft()))) % 1000000;
        logFilename = string(abi.encodePacked("./fuzz_results_", vm.toString(uniqueId), ".csv"));
        //  Removed header writing from here

        attacker = new User("Attacker"); // Attacker serves as both operator and staker
        badAVS = new AVS("BadAVS"); // AVS is also attacker-controlled
        strategy = lstStrats[0];
        token = IERC20Metadata(address(strategy.underlyingToken()));
        
        // Create users in a loop and give them some tokens
        for (uint256 i = 0; i < users.length; i++) {
            users[i] = new User(string(abi.encodePacked("randomUser", vm.toString(i + 1))));
            deal(address(token), address(users[i]), 1e18);
        }
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
        
        _print("setup");
    }

    // TODO: consider incremental manual fuzzing from 1 up to WAD - 1
    // Removed test that test's rounding with full mags slashed - all interesting results come from partial mags
    function test_rounding_partialMagsSlashed(
        uint64 initialWadToSlash, 
        uint64 _initTokenBalance
    ) public rand(0) {
        vm.pauseGasMetering();
        
        //  Write header only on first fuzz run
        if (!headerWritten) {
            vm.writeLine(logFilename, "initial_slash,init_balance,final_slash,operator_shares,status");
            headerWritten = true;
        }
        
        // Don't slash 100% as we will do multiple slashes
        initialWadToSlash = uint64(bound(initialWadToSlash, 1, WAD - 1));
        // Ensure attacker has at least one token
        initTokenBalance = _initTokenBalance > 0 ? _initTokenBalance : _initTokenBalance + 1;
        deal(address(token), address(attacker), initTokenBalance);
        
        // Use modifyAllocation+slashOperator to arbitrarily set operator max magnitude
        _magnitudeManipulation(initialWadToSlash);
        _deposit(initTokenBalance);
        
        // Perform slashes to gradually whittle down operator magnitude, as well as produce slash escrows.
        // Since we're doing multiple slashes, never slash 100%.
        //for (uint16 i = 0; i < numSlashes; i++) {
        //    uint64 wadToSlash = uint64(cheats.randomUint(1, WAD - 1));
        //}
        uint64 wadToSlash = uint64(cheats.randomUint(1, WAD - 1));
        _slash(wadToSlash);

        
        // Release all slashing escrows to all users and attacker.
        _release();
        // Withdraw all user and attacker deposits after withdrawal delay
        _withdraw();

        uint256 finalOperatorShares = delegationManager.getOperatorShares(address(attacker), strategy.toArray())[0];
        // Log ONLY the key metrics for each run (one line per test)
        string memory runData = string(abi.encodePacked(
        vm.toString(initialWadToSlash), ",",
        vm.toString(initTokenBalance), ",", 
        vm.toString(wadToSlash), ",",
        vm.toString(finalOperatorShares), ",",
        finalOperatorShares > 20 ? "EXPLOIT" : "CLEAN"
        ));
        vm.writeLine(logFilename, runData);
        
        // Only console.log for exploits (will show in terminal)
        if(finalOperatorShares > 20) {
            console.log("EXPLOIT FOUND!");
            console.log("Params: initial_slash=%d, init_balance=%d, final_slash=%d", 
                    initialWadToSlash, initTokenBalance, wadToSlash);
            console.log("Final operator shares: %d", finalOperatorShares);
            revert("Operator shares greater than 20!");
        }
    }
    
    // TODO - another way to mess with rounding/precision loss is to manipulate DSF
    function _magnitudeManipulation(uint64 wadToSlash) internal {
        // Allocate all magnitude to operator set.
        attacker.modifyAllocations(AllocateParams({
            operatorSet: mOpSet,
            strategies: strategy.toArray(),
            newMagnitudes: WAD.toArrayU64()
        }));
        
        // TODO: print "newMagnitudes"

        _print("allocate");

        // Slash operator to arbitrary mag.
        badAVS.slashOperator(SlashingParams({
            operator: address(attacker),
            operatorSetId: mOpSet.id,
            strategies: strategy.toArray(),
            wadsToSlash: uint(wadToSlash).toArrayU256(),
            description: "manipulation!"
        }));
        
        // TODO: print "wadsToSlash"

        _print("slash");
        
        // Deallocate magnitude from operator set.
        attacker.modifyAllocations(AllocateParams({
            operatorSet: mOpSet,
            strategies: strategy.toArray(),
            newMagnitudes: 0.toArrayU64()
        }));
        
        rollForward({blocks: DEALLOCATION_DELAY + 1});
        
        _print("deallocate");
    }

    function _deposit(uint64 initTokenBalance) internal {
        // Allocate all remaining magnitude to redistributable opset.
        uint64 allocatableMagnitude = allocationManager.getAllocatableMagnitude(address(attacker), strategy);
        attacker.modifyAllocations(AllocateParams({
            operatorSet: rOpSet,
            strategies: strategy.toArray(),
            newMagnitudes: (allocatableMagnitude).toArrayU64()
        }));
        
        // case where we randomly add more tokens to the attacker wallet,
        // its a boolean that we randomly set to true
        bool addMoreTokens = cheats.randomBool();
        if(addMoreTokens) {
            deal(address(token), address(attacker),initTokenBalance);
        }


        
        // Deposit all attacker assets into Eigenlayer.
        attacker.depositIntoEigenlayer(strategy.toArray(), token.balanceOf(address(attacker)).toArrayU256());
        // Deposit all users into Eigenlayer
        for (uint256 i = 0; i < users.length; i++) {
            users[i].depositIntoEigenlayer(strategy.toArray(), token.balanceOf(address(users[i])).toArrayU256());
        }

        _print("deposit");
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
        
        _print("slash");
    }
    // release slashed funds from slashing escrow
    function _release() internal {
        // Roll forward past the escrow delay.
        rollForward({blocks: slashEscrowFactory.getGlobalEscrowDelay() + 1});
        
               // Release funds.
        vm.prank(address(attacker));
        slashEscrowFactory.releaseSlashEscrow(mOpSet, 1);

        // Release funds.
        vm.prank(address(attacker));
        slashEscrowFactory.releaseSlashEscrow(rOpSet, 1);
     
        _print("release");
    }

    function _withdraw() internal {

        (, uint256[] memory depositShares) = strategyManager.getDeposits(address(attacker));
        
        Withdrawal[] memory withdrawals = attacker.queueWithdrawals(strategy.toArray(), depositShares);

        Withdrawal[][] memory userWithdrawals = new Withdrawal[][](users.length);
        
        for (uint256 i = 0; i < users.length; i++) {
            (, uint256[] memory depositShares) = strategyManager.getDeposits(address(users[i]));
            userWithdrawals[i] = users[i].queueWithdrawals(strategy.toArray(), depositShares);
        }

        rollForward({blocks: DELEGATION_MANAGER_MIN_WITHDRAWAL_DELAY_BLOCKS + 1});
        
        attacker.completeWithdrawalsAsTokens(withdrawals);
        for (uint256 i = 0; i < users.length; i++) {
            users[i].completeWithdrawalsAsTokens(userWithdrawals[i]);
        }

        
        _print("withdraw");

    }
    function _print(string memory phaseName) internal {
        address a = address(attacker);

        console.log("");    
        console.log("===Attacker Info: %s phase===".cyan(), phaseName);

        {
            console.log("\nRaw Assets:".magenta());
            console.log(" - token: %s", token.symbol());
            console.log(" - held balance: %d", token.balanceOf(a));
            // TODO - amt deposited, possibly keep track of this separately?
        }

        {
            console.log("\nShares:".magenta());

            (uint[] memory withdrawableArr, uint[] memory depositArr) 
                = delegationManager.getWithdrawableShares(a, strategy.toArray());
            uint withdrawableShares = withdrawableArr.length == 0 ? 0 : withdrawableArr[0];
            uint depositShares = depositArr.length == 0 ? 0 : depositArr[0];
            console.log(" - deposit shares: %d", depositShares);
            console.log(" - withdrawable shares: %d", withdrawableShares);
            console.log(" - operator shares: %d", delegationManager.operatorShares(a, strategy));
        }

        {
            console.log("\nScaling:".magenta());

            Allocation memory mAlloc = allocationManager.getAllocation(a, mOpSet, strategy);
            Allocation memory rAlloc = allocationManager.getAllocation(a, rOpSet, strategy);

            console.log(" - Init Mag: %d", WAD);
            console.log(
                " - Max Mag:  %d\n   -- Total Allocated: %d\n   -- Total Available: %d", 
                allocationManager.getMaxMagnitude(a, strategy),
                allocationManager.getEncumberedMagnitude(a, strategy),
                allocationManager.getAllocatableMagnitude(a, strategy)
            );
            console.log(" - Allocated to mOpSet: %d", mAlloc.currentMagnitude);
            console.log(" - Allocated to rOpSet: %d", rAlloc.currentMagnitude);
            console.log(" - DSF: %d", delegationManager.depositScalingFactor(a, strategy));
        }

        console.log("\n  ===\n".cyan());
    }
}