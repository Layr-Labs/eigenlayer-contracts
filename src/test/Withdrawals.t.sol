// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../test/EigenLayerTestHelper.t.sol";

contract WithdrawalTests is EigenLayerTestHelper {
    // packed info used to help handle stack-too-deep errors
    struct DataForTestWithdrawal {
        IStrategy[] delegatorStrategies;
        uint256[] delegatorShares;
        address withdrawer;
        uint96 nonce;
    }

    bytes32 defaultOperatorId = bytes32(uint256(0));

    function setUp() public virtual override {
        EigenLayerDeployer.setUp();
    }

    //This function helps with stack too deep issues with "testWithdrawal" test
    function testWithdrawalWrapper(
        address operator,
        address depositor,
        uint96 ethAmount,
        uint96 eigenAmount,
        bool withdrawAsTokens,
        bool RANDAO
    ) public fuzzedAddress(operator) fuzzedAddress(depositor) {
        cheats.assume(depositor != operator);
        cheats.assume(ethAmount >= 1 && ethAmount <= 1e18);
        cheats.assume(eigenAmount >= 1 && eigenAmount <= 1e18);

        address withdrawer = depositor;

        if (RANDAO) {
            _testWithdrawalAndDeregistration(operator, depositor, withdrawer, ethAmount, eigenAmount, withdrawAsTokens);
        } else {
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
    ) internal {
        _initiateDelegation(operator, depositor, ethAmount, eigenAmount);

        address delegatedTo = delegation.delegatedTo(depositor);

        // packed data structure to deal with stack-too-deep issues
        DataForTestWithdrawal memory dataForTestWithdrawal;

        // scoped block to deal with stack-too-deep issues
        {
            //delegator-specific information
            (IStrategy[] memory delegatorStrategies, uint256[] memory delegatorShares) = strategyManager.getDeposits(
                depositor
            );
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

        {
            //warp past the serve until time, which is 3 days from the beginning.  THis puts us at 4 days past that point
            cheats.warp(uint32(block.timestamp) + 4 days);
            cheats.roll(uint32(block.timestamp) + 4 days);

            uint256 middlewareTimeIndex = 1;
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
    ) public {
        _initiateDelegation(operator, depositor, ethAmount, eigenAmount);

        // emit log_named_uint("Linked list element 1", uint256(uint160(address(generalServiceManager1))));
        // emit log_named_uint("Linked list element 2", uint256(uint160(address(generalServiceManager2))));
        // emit log("________________________________________________________________");
        // emit log_named_uint("Middleware 1 Update Block", uint32(block.number));

        cheats.warp(uint32(block.timestamp) + 1 days);
        cheats.roll(uint32(block.number) + 1);

        // emit log_named_uint("Middleware 2 Update Block", uint32(block.number));

        address delegatedTo = delegation.delegatedTo(depositor);

        // packed data structure to deal with stack-too-deep issues
        DataForTestWithdrawal memory dataForTestWithdrawal;

        // scoped block to deal with stack-too-deep issues
        {
            //delegator-specific information
            (IStrategy[] memory delegatorStrategies, uint256[] memory delegatorShares) = strategyManager.getDeposits(
                depositor
            );
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

        // uint256 prevElement = uint256(uint160(address(generalServiceManager2)));

        cheats.warp(uint32(block.timestamp) + 1 days);
        cheats.roll(uint32(block.number) + 1);

        // prevElement = uint256(uint160(address(generalServiceManager1)));

        {
            //warp past the serve until time, which is 3 days from the beginning.  THis puts us at 4 days past that point
            cheats.warp(uint32(block.timestamp) + 4 days);
            cheats.roll(uint32(block.number) + 4);

            uint256 middlewareTimeIndex = 3;
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
        uint96 ethAmount,
        uint96 eigenAmount,
        bool withdrawAsShares
    ) public fuzzedAddress(operator) fuzzedAddress(depositor) {
        cheats.assume(depositor != operator);
        //this function performs delegation and subsequent withdrawal
        testWithdrawalWrapper(operator, depositor, ethAmount, eigenAmount, withdrawAsShares, true);

        cheats.prank(depositor);
        delegation.undelegate(depositor);

        //warps past fraudproof time interval
        cheats.warp(block.timestamp + 7 days + 1);
        _initiateDelegation(operator, depositor, ethAmount, eigenAmount);
    }

    // Helper function to begin a delegation
    function _initiateDelegation(
        address operator,
        address staker,
        uint96 ethAmount,
        uint96 eigenAmount
    ) internal fuzzedAddress(operator) fuzzedAddress(staker) fuzzedAmounts(ethAmount, eigenAmount) {
        cheats.assume(staker != operator);
        // base strategy will revert if these amounts are too small on first deposit
        cheats.assume(ethAmount >= 1);
        cheats.assume(eigenAmount >= 2);

        _testDelegation(operator, staker, ethAmount, eigenAmount);
    }

    modifier fuzzedAmounts(uint256 ethAmount, uint256 eigenAmount) {
        cheats.assume(ethAmount >= 0 && ethAmount <= 1e18);
        cheats.assume(eigenAmount >= 0 && eigenAmount <= 1e18);
        _;
    }
}

