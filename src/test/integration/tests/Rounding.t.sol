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

    OperatorSet mOpSet; // "manipOpSet" used for magnitude manipulation
    OperatorSet rOpSet; // redistributable opset used to exploit precision loss and trigger redistribution
    
    function _init() internal override {
        _configAssetTypes(HOLDS_LST);

        attacker = new User("Attacker");
        badAVS = new AVS("BadAVS");
        strategy = lstStrats[0];
        token = IERC20Metadata(address(strategy.underlyingToken()));

        // Register attacker as operator and create attacker-controller AVS/OpSets
        attacker.registerAsOperator(0);
        rollForward({blocks: ALLOCATION_CONFIGURATION_DELAY + 1});
        badAVS.updateAVSMetadataURI("https://example.com");
        mOpSet = badAVS.createOperatorSet(strategy.toArray());
        rOpSet = badAVS.createRedistributingOperatorSet(strategy.toArray(), address(attacker));

        // Register for both opsets
        attacker.registerForOperatorSet(mOpSet);
        attacker.registerForOperatorSet(rOpSet);

        _print("setup");
    }

    function test_rounding() public rand(0) {
        _magnitudeManipulation();
        _setupFinal();
        _final();
    }
    
    // TODO - another way to mess with rounding/precision loss is to manipulate DSF
    function _magnitudeManipulation() internal {
        attacker.modifyAllocations(AllocateParams({
            operatorSet: mOpSet,
            strategies: strategy.toArray(),
            newMagnitudes: WAD.toArrayU64()
        }));

        _print("allocate");

        badAVS.slashOperator(SlashingParams({
            operator: address(attacker),
            operatorSetId: mOpSet.id,
            strategies: strategy.toArray(),
            wadsToSlash: uint(WAD-1).toArrayU256(),
            description: ""
        }));

        _print("slash");
    }

    function _setupFinal() internal {
        // create redistributable opset
        // deposit assets
    }

    function _final() internal {
        // perform final slash on redistributable opset and check for profit
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
