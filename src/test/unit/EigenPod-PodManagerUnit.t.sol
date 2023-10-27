///@notice Placeholder for future unit tests that combine interaction between the EigenPod & EigenPodManager

// TODO: salvage / re-implement a check for reentrancy guard on functions, as possible
    // function testRecordBeaconChainETHBalanceUpdateFailsWhenReentering() public {
    //     uint256 amount = 1e18;
    //     uint256 amount2 = 2e18;
    //     address staker = address(this);
    //     uint256 beaconChainETHStrategyIndex = 0;

    //     _beaconChainReentrancyTestsSetup();

    //     testRestakeBeaconChainETHSuccessfully(staker, amount);        

    //     address targetToUse = address(strategyManager);
    //     uint256 msgValueToUse = 0;

    //     int256 amountDelta = int256(amount2 - amount);
    //     // reference: function recordBeaconChainETHBalanceUpdate(address podOwner, uint256 beaconChainETHStrategyIndex, uint256 sharesDelta, bool isNegative)
    //     bytes memory calldataToUse = abi.encodeWithSelector(StrategyManager.recordBeaconChainETHBalanceUpdate.selector, staker, beaconChainETHStrategyIndex, amountDelta);
    //     reenterer.prepare(targetToUse, msgValueToUse, calldataToUse, bytes("ReentrancyGuard: reentrant call"));

    //     cheats.startPrank(address(reenterer));
    //     eigenPodManager.recordBeaconChainETHBalanceUpdate(staker, amountDelta);
    //     cheats.stopPrank();
    // }

    // function _beaconChainReentrancyTestsSetup() internal {
    //     // prepare EigenPodManager with StrategyManager and Delegation replaced with a Reenterer contract
    //     reenterer = new Reenterer();
    //     eigenPodManagerImplementation = new EigenPodManager(
    //         ethPOSMock,
    //         eigenPodBeacon,
    //         IStrategyManager(address(reenterer)),
    //         slasherMock,
    //         IDelegationManager(address(reenterer))
    //     );
    //     eigenPodManager = EigenPodManager(
    //         address(
    //             new TransparentUpgradeableProxy(
    //                 address(eigenPodManagerImplementation),
    //                 address(proxyAdmin),
    //                 abi.encodeWithSelector(
    //                     EigenPodManager.initialize.selector,
    //                     type(uint256).max /*maxPods*/,
    //                     IBeaconChainOracle(address(0)) /*beaconChainOracle*/,
    //                     initialOwner,
    //                     pauserRegistry,
    //                     0 /*initialPausedStatus*/
    //                 )
    //             )
    //         )
    //     );
    // }
