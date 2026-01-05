// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "src/contracts/interfaces/IRewardsCoordinator.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

contract RewardsCoordinatorMock is Test {
    receive() external payable {}
    fallback() external payable {}

    struct SetOperatorSetSplitCall {
        address operator;
        OperatorSet operatorSet;
        uint16 split;
    }

    SetOperatorSetSplitCall internal _lastSetOperatorSetSplitCall;
    uint public setOperatorSetSplitCallCount;

    function setOperatorSetSplit(address operator, OperatorSet calldata operatorSet, uint16 split) external {
        setOperatorSetSplitCallCount++;
        _lastSetOperatorSetSplitCall = SetOperatorSetSplitCall({operator: operator, operatorSet: operatorSet, split: split});
    }

    function lastSetOperatorSetSplitCall() external view returns (SetOperatorSetSplitCall memory) {
        return _lastSetOperatorSetSplitCall;
    }
}
