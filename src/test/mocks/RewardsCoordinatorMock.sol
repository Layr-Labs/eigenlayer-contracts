// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "src/contracts/interfaces/IRewardsCoordinator.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

contract RewardsCoordinatorMock is Test {
    receive() external payable {}
    fallback() external payable {}

    struct SetOperatorAVSSplitCall {
        address operator;
        address avs;
        uint16 split;
    }

    struct SetOperatorSetSplitCall {
        address operator;
        OperatorSet operatorSet;
        uint16 split;
    }

    struct SetClaimerForCall {
        address earner;
        address claimer;
    }

    SetOperatorAVSSplitCall internal _lastSetOperatorAVSSplitCall;
    SetOperatorSetSplitCall internal _lastSetOperatorSetSplitCall;
    SetClaimerForCall internal _lastSetClaimerForCall;
    uint public setOperatorAVSSplitCallCount;
    uint public setOperatorSetSplitCallCount;
    uint public setClaimerForCallCount;

    function setOperatorAVSSplit(address operator, address avs, uint16 split) external {
        setOperatorAVSSplitCallCount++;
        _lastSetOperatorAVSSplitCall = SetOperatorAVSSplitCall({operator: operator, avs: avs, split: split});
    }

    function setOperatorSetSplit(address operator, OperatorSet calldata operatorSet, uint16 split) external {
        setOperatorSetSplitCallCount++;
        _lastSetOperatorSetSplitCall = SetOperatorSetSplitCall({operator: operator, operatorSet: operatorSet, split: split});
    }

    function setClaimerFor(address earner, address claimer) external {
        setClaimerForCallCount++;
        _lastSetClaimerForCall = SetClaimerForCall({earner: earner, claimer: claimer});
    }

    function lastSetOperatorAVSSplitCall() external view returns (SetOperatorAVSSplitCall memory) {
        return _lastSetOperatorAVSSplitCall;
    }

    function lastSetOperatorSetSplitCall() external view returns (SetOperatorSetSplitCall memory) {
        return _lastSetOperatorSetSplitCall;
    }

    function lastSetClaimerForCall() external view returns (SetClaimerForCall memory) {
        return _lastSetClaimerForCall;
    }
}
