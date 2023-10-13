// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/utils/math/Math.sol";

import "./mocks/MiddlewareRegistryMock.sol";
import "./mocks/ServiceManagerMock.sol";
import "./Delegation.t.sol";

contract WithdrawalTests is DelegationTests {

    // packed info used to help handle stack-too-deep errors
    struct DataForTestWithdrawal {
        IStrategy[] delegatorStrategies;
        uint256[] delegatorShares;
        address withdrawer;
        uint96 nonce;
    }

    MiddlewareRegistryMock public generalReg1;
    ServiceManagerMock public generalServiceManager1;

    MiddlewareRegistryMock public generalReg2;
    ServiceManagerMock public generalServiceManager2;

    function initializeGeneralMiddlewares() public {
        generalServiceManager1 = new ServiceManagerMock(slasher);

        generalReg1 = new MiddlewareRegistryMock(
             generalServiceManager1,
             strategyManager
        );
        
        generalServiceManager2 = new ServiceManagerMock(slasher);

        generalReg2 = new MiddlewareRegistryMock(
             generalServiceManager2,
             strategyManager
        );
    }

    //This function helps with stack too deep issues with "testWithdrawal" test
    function testWithdrawalWrapper(
            address operator, 
            address depositor,
            address withdrawer, 
            uint96 ethAmount,
            uint96 eigenAmount,
            bool withdrawAsTokens,
            bool RANDAO
        ) 
            public 
            fuzzedAddress(operator) 
            fuzzedAddress(depositor) 
            fuzzedAddress(withdrawer) 
        {
            cheats.assume(depositor != operator);
            cheats.assume(ethAmount >= 1 && ethAmount <= 1e18); 
            cheats.assume(eigenAmount >= 1 && eigenAmount <= 1e18); 

            initializeGeneralMiddlewares();

            if(RANDAO) {
                _testWithdrawalAndDeregistration(operator, depositor, withdrawer, ethAmount, eigenAmount, withdrawAsTokens);
            }
            else{
                _testWithdrawalWithStakeUpdate(operator, depositor, withdrawer, ethAmount, eigenAmount, withdrawAsTokens);
            }

        }

    /// @notice test staker's ability to undelegate/withdraw from an operator.
    /// @param operator is the operator being delegated to.
    /// @param depositor is the staker delegating stake to the operator.
    function _testWithdrawalAndDeregistration(
            address operator, 
            address depositor,
            address withdrawer, 
            uint96 ethAmount,
            uint96 eigenAmount,
            bool withdrawAsTokens
        ) 
            internal 
        {

        testDelegation(operator, depositor, ethAmount, eigenAmount);

        cheats.startPrank(operator);
        slasher.optIntoSlashing(address(generalServiceManager1));
        cheats.stopPrank();

        generalReg1.registerOperator(operator, uint32(block.timestamp) + 3 days);

        address delegatedTo = delegation.delegatedTo(depositor);

        // packed data structure to deal with stack-too-deep issues
        DataForTestWithdrawal memory dataForTestWithdrawal;

        // scoped block to deal with stack-too-deep issues
        {
            //delegator-specific information
            (IStrategy[] memory delegatorStrategies, uint256[] memory delegatorShares) =
                strategyManager.getDeposits(depositor);
            dataForTestWithdrawal.delegatorStrategies = delegatorStrategies;
            dataForTestWithdrawal.delegatorShares = delegatorShares;
            dataForTestWithdrawal.withdrawer = withdrawer;
            // harcoded nonce value
            dataForTestWithdrawal.nonce = 0;
        }

        uint256[] memory strategyIndexes = new uint256[](2);
        IERC20[] memory tokensArray = new IERC20[](2);
        {
            // hardcoded values
            strategyIndexes[0] = 0;
            strategyIndexes[1] = 0;
            tokensArray[0] = weth;
            tokensArray[1] = eigenToken;
        }

        cheats.warp(uint32(block.timestamp) + 1 days);
        cheats.roll(uint32(block.timestamp) + 1 days);

        _testQueueWithdrawal(
            depositor,
            strategyIndexes,
            dataForTestWithdrawal.delegatorStrategies,
            dataForTestWithdrawal.delegatorShares,
            withdrawer
        );
        uint32 queuedWithdrawalBlock = uint32(block.number);
        
        //now withdrawal block time is before deregistration
        cheats.warp(uint32(block.timestamp) + 2 days);
        cheats.roll(uint32(block.timestamp) + 2 days);
        
        generalReg1.deregisterOperator(operator);
        {
            //warp past the serve until time, which is 3 days from the beginning.  THis puts us at 4 days past that point
            cheats.warp(uint32(block.timestamp) + 4 days);
            cheats.roll(uint32(block.timestamp) + 4 days);

            uint256 middlewareTimeIndex =  1;
            if (withdrawAsTokens) {
                _testCompleteQueuedWithdrawalTokens(
                    depositor,
                    dataForTestWithdrawal.delegatorStrategies,
                    tokensArray,
                    dataForTestWithdrawal.delegatorShares,
                    delegatedTo,
                    dataForTestWithdrawal.withdrawer,
                    dataForTestWithdrawal.nonce,
                    queuedWithdrawalBlock,
                    middlewareTimeIndex
                );
            } else {
                _testCompleteQueuedWithdrawalShares(
                    depositor,
                    dataForTestWithdrawal.delegatorStrategies,
                    tokensArray,
                    dataForTestWithdrawal.delegatorShares,
                    delegatedTo,
                    dataForTestWithdrawal.withdrawer,
                    dataForTestWithdrawal.nonce,
                    queuedWithdrawalBlock,
                    middlewareTimeIndex
                );
            }
        }
    }

    /// @notice test staker's ability to undelegate/withdraw from an operator.
    /// @param operator is the operator being delegated to.
    /// @param depositor is the staker delegating stake to the operator.
    function _testWithdrawalWithStakeUpdate(
            address operator, 
            address depositor,
            address withdrawer, 
            uint96 ethAmount,
            uint96 eigenAmount,
            bool withdrawAsTokens
        ) 
            public 
        {
        testDelegation(operator, depositor, ethAmount, eigenAmount);

        cheats.startPrank(operator);
        slasher.optIntoSlashing(address(generalServiceManager1));
        slasher.optIntoSlashing(address(generalServiceManager2));
        cheats.stopPrank();

        // emit log_named_uint("Linked list element 1", uint256(uint160(address(generalServiceManager1))));
        // emit log_named_uint("Linked list element 2", uint256(uint160(address(generalServiceManager2))));
        // emit log("________________________________________________________________");
        generalReg1.registerOperator(operator, uint32(block.timestamp) + 5 days);
        // emit log_named_uint("Middleware 1 Update Block", uint32(block.number));

        cheats.warp(uint32(block.timestamp) + 1 days);
        cheats.roll(uint32(block.number) + 1);

        generalReg2.registerOperator(operator, uint32(block.timestamp) + 5 days);
        // emit log_named_uint("Middleware 2 Update Block", uint32(block.number));

        address delegatedTo = delegation.delegatedTo(depositor);

        // packed data structure to deal with stack-too-deep issues
        DataForTestWithdrawal memory dataForTestWithdrawal;

        // scoped block to deal with stack-too-deep issues
        {
            //delegator-specific information
            (IStrategy[] memory delegatorStrategies, uint256[] memory delegatorShares) =
                strategyManager.getDeposits(depositor);
            dataForTestWithdrawal.delegatorStrategies = delegatorStrategies;
            dataForTestWithdrawal.delegatorShares = delegatorShares;
            dataForTestWithdrawal.withdrawer = withdrawer;
            // harcoded nonce value
            dataForTestWithdrawal.nonce = 0;
        }

        uint256[] memory strategyIndexes = new uint256[](2);
        IERC20[] memory tokensArray = new IERC20[](2);
        {
            // hardcoded values
            strategyIndexes[0] = 0;
            strategyIndexes[1] = 0;
            tokensArray[0] = weth;
            tokensArray[1] = eigenToken;
        }

        cheats.warp(uint32(block.timestamp) + 1 days);
        cheats.roll(uint32(block.number) + 1);

        _testQueueWithdrawal(
            depositor,
            strategyIndexes,
            dataForTestWithdrawal.delegatorStrategies,
            dataForTestWithdrawal.delegatorShares,
            dataForTestWithdrawal.withdrawer
        );
        uint32 queuedWithdrawalBlock = uint32(block.number);
        
        //now withdrawal block time is before deregistration
        cheats.warp(uint32(block.timestamp) + 2 days);
        cheats.roll(uint32(block.number) + 2);

        
        uint256 prevElement = uint256(uint160(address(generalServiceManager2)));
        generalReg1.propagateStakeUpdate(operator, uint32(block.number), prevElement);

        cheats.warp(uint32(block.timestamp) + 1 days);
        cheats.roll(uint32(block.number) + 1);

        prevElement = uint256(uint160(address(generalServiceManager1)));
        generalReg2.propagateStakeUpdate(operator, uint32(block.number), prevElement);
        
        {
            //warp past the serve until time, which is 3 days from the beginning.  THis puts us at 4 days past that point
            cheats.warp(uint32(block.timestamp) + 4 days);
            cheats.roll(uint32(block.number) + 4);

            uint256 middlewareTimeIndex =  3;
            if (withdrawAsTokens) {
                _testCompleteQueuedWithdrawalTokens(
                    depositor,
                    dataForTestWithdrawal.delegatorStrategies,
                    tokensArray,
                    dataForTestWithdrawal.delegatorShares,
                    delegatedTo,
                    dataForTestWithdrawal.withdrawer,
                    dataForTestWithdrawal.nonce,
                    queuedWithdrawalBlock,
                    middlewareTimeIndex
                );
            } else {
                _testCompleteQueuedWithdrawalShares(
                    depositor,
                    dataForTestWithdrawal.delegatorStrategies,
                    tokensArray,
                    dataForTestWithdrawal.delegatorShares,
                    delegatedTo,
                    dataForTestWithdrawal.withdrawer,
                    dataForTestWithdrawal.nonce,
                    queuedWithdrawalBlock,
                    middlewareTimeIndex
                );
            }
        }
    }

    // @notice This function tests to ensure that a delegator can re-delegate to an operator after undelegating.
    // @param operator is the operator being delegated to.
    // @param staker is the staker delegating stake to the operator.
    function testRedelegateAfterWithdrawal(
            address operator, 
            address depositor, 
            address withdrawer, 
            uint96 ethAmount, 
            uint96 eigenAmount,
            bool withdrawAsShares
        ) 
            public
            fuzzedAddress(operator) 
            fuzzedAddress(depositor)
            fuzzedAddress(withdrawer)
        {
        cheats.assume(depositor != operator);
        //this function performs delegation and subsequent withdrawal
        testWithdrawalWrapper(operator, depositor, withdrawer, ethAmount, eigenAmount, withdrawAsShares, true);

        cheats.prank(depositor);
        delegation.undelegate(depositor);

        //warps past fraudproof time interval
        cheats.warp(block.timestamp + 7 days + 1);
        testDelegation(operator, depositor, ethAmount, eigenAmount);
    }

    // onlyNotFrozen modifier is not used in current DelegationManager implementation.
    // commented out test case for now
    // /// @notice test to see if an operator who is slashed/frozen
    // ///         cannot be undelegated from by their stakers.
    // /// @param operator is the operator being delegated to.
    // /// @param staker is the staker delegating stake to the operator.
    // function testSlashedOperatorWithdrawal(address operator, address staker, uint96 ethAmount, uint96 eigenAmount)
    //     public
    //     fuzzedAddress(operator)
    //     fuzzedAddress(staker)
    // {
    //     cheats.assume(staker != operator);
    //     testDelegation(operator, staker, ethAmount, eigenAmount);

    //     {
    //         address slashingContract = slasher.owner(); 

    //         cheats.startPrank(operator);
    //         slasher.optIntoSlashing(address(slashingContract));
    //         cheats.stopPrank();

    //         cheats.startPrank(slashingContract);
    //         slasher.freezeOperator(operator);
    //         cheats.stopPrank();
    //     }

    //     (IStrategy[] memory updatedStrategies, uint256[] memory updatedShares) =
    //         strategyManager.getDeposits(staker);

    //     uint256[] memory strategyIndexes = new uint256[](2);
    //     strategyIndexes[0] = 0;
    //     strategyIndexes[1] = 1;

    //     IERC20[] memory tokensArray = new IERC20[](2);
    //     tokensArray[0] = weth;
    //     tokensArray[0] = eigenToken;

    //     //initiating queued withdrawal
    //     cheats.expectRevert(
    //         bytes("StrategyManager.onlyNotFrozen: staker has been frozen and may be subject to slashing")
    //     );
    //     _testQueueWithdrawal(staker, strategyIndexes, updatedStrategies, updatedShares, staker);
    // }
}
