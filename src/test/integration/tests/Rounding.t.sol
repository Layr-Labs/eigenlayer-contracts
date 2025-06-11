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

    OperatorSet mOpSet; // "manipOpSet" used for magnitude manipulation
    OperatorSet rOpSet; // redistributable opset used to exploit precision loss and trigger redistribution
    
    function _init() internal override {
        _configAssetTypes(HOLDS_LST);

        attacker = new User("Attacker"); // attacker can be both operator and staker
        badAVS = new AVS("BadAVS");
        strategy = lstStrats[0];
        token = IERC20Metadata(address(strategy.underlyingToken()));
        deal(address(token), address(attacker), uint256(1000000000000000000)); // TODO: make the balance a fuzzing param
        
        // good staker and operator setup
        goodStaker = new User("goodStaker");
        deal(address(token), address(goodStaker), uint256(1000000000000000000)); // TODO: make the balance a fuzzing param
        
        // Register attacker as operator and create attacker-controller AVS/OpSets
        attacker.registerAsOperator(0);
        rollForward({blocks: ALLOCATION_CONFIGURATION_DELAY + 1});
        badAVS.updateAVSMetadataURI("https://example.com");
        mOpSet = badAVS.createOperatorSet(strategy.toArray()); // setup low mag operator
        rOpSet = badAVS.createRedistributingOperatorSet(strategy.toArray(), address(attacker)); // execute exploit

        // Register for both opsets
        attacker.registerForOperatorSet(mOpSet);
        attacker.registerForOperatorSet(rOpSet);
        
        // TODO: considerations
        // - assets of other users within the opSet
        // - calculate total assets entering and leaving the protocol, estimate precision loss (+ or -)

        _print("setup");
    }

    // TODO: parameterize deposits
    // TODO: consider manual fuzzing from 1 up to WAD - 1
    function test_rounding(uint64 wadToSlash) public rand(0) { 
        // bound wadToSlash to 1 < wadToSlash < WAD - 1
        wadToSlash = uint64(bound(wadToSlash, 1, WAD - 1));
        
        _magnitudeManipulation(wadToSlash); // get operator to low mag
        _setupFinal(wadToSlash); // 
        _final(wadToSlash);
    }
    
    // TODO - another way to mess with rounding/precision loss is to manipulate DSF
    function _magnitudeManipulation(uint64 wadToSlash) internal {
        
        // allocate magnitude to operator set
        attacker.modifyAllocations(AllocateParams({
            operatorSet: mOpSet,
            strategies: strategy.toArray(),
            newMagnitudes: WAD.toArrayU64()
        }));
        
        // TODO: print "newMagnitudes"

        _print("allocate");

        // slash operator to low mag
        badAVS.slashOperator(SlashingParams({
            operator: address(attacker),
            operatorSetId: mOpSet.id,
            strategies: strategy.toArray(),
            wadsToSlash: uint(wadToSlash).toArrayU256(), // TODO: Make WAD-1 a fuzzing param
            description: "manipulation!"
        }));
        
        // TODO: print "wadsToSlash"

        _print("slash");
        
        // deallocate magnitude from operator set
        attacker.modifyAllocations(AllocateParams({
            operatorSet: mOpSet,
            strategies: strategy.toArray(),
            newMagnitudes: 0.toArrayU64()
        }));
        
        rollForward({blocks: DEALLOCATION_DELAY + 1});
        
        _print("deallocate");
    }

    function _setupFinal(uint64 wadToSlash) internal {
        // allocate to redistributable opset
        attacker.modifyAllocations(AllocateParams({
            operatorSet: rOpSet,
            strategies: strategy.toArray(),
            newMagnitudes: (WAD - wadToSlash).toArrayU64()
        }));
        
        // deposit assets and delegate to attacker
        attacker.depositIntoEigenlayer(strategy.toArray(), token.balanceOf(address(attacker)).toArrayU256());
        
        
        goodStaker.depositIntoEigenlayer(strategy.toArray(), token.balanceOf(address(goodStaker)).toArrayU256());
        
        // TODO: print "initTokenBalances"
        
        _print("deposit");
    }

    function _final(uint64 wadToSlash) internal {
        // perform final slash on redistributable opset and check for profit
        badAVS.slashOperator(SlashingParams({
            operator: address(attacker),
            operatorSetId: rOpSet.id,
            strategies: strategy.toArray(),
            wadsToSlash: uint(WAD - wadToSlash).toArrayU256(),
            description: "final slash"
        }));
        
        _print("slash");
        
        // roll forward past the escrow delay
        rollForward({blocks: slashEscrowFactory.getGlobalEscrowDelay() + 1});
        
        // release funds
        vm.prank(address(attacker));
        slashEscrowFactory.releaseSlashEscrow(rOpSet, 1); // 1 is used as it's the first slashId
        
        _print("release");
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
