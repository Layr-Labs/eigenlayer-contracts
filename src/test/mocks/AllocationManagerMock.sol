// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "src/contracts/interfaces/IStrategy.sol";
import "src/contracts/libraries/Snapshots.sol";
import "src/contracts/libraries/OperatorSetLib.sol";

contract AllocationManagerMock is Test {
    using Snapshots for Snapshots.DefaultWadHistory;
    using OperatorSetLib for OperatorSet;

    receive() external payable {}
    fallback() external payable {}

    mapping(bytes32 operatorSetKey => bool) public _isOperatorSet;
    mapping(address avs => uint) public getOperatorSetCount;
    mapping(address => mapping(IStrategy => Snapshots.DefaultWadHistory)) internal _maxMagnitudeHistory;

    function setIsOperatorSet(OperatorSet memory operatorSet, bool boolean) external {
        _isOperatorSet[operatorSet.key()] = boolean;
    }

    function isOperatorSet(OperatorSet memory operatorSet) external view returns (bool) {
        return _isOperatorSet[operatorSet.key()];
    }

    function setMaxMagnitudes(address operator, IStrategy[] calldata strategies, uint64[] calldata maxMagnitudes) external {
        for (uint i = 0; i < strategies.length; ++i) {
            setMaxMagnitude(operator, strategies[i], maxMagnitudes[i]);
        }
    }

    function setMaxMagnitude(address operator, IStrategy strategy, uint64 maxMagnitude) public {
        _maxMagnitudeHistory[operator][strategy].push({key: uint32(block.number), value: maxMagnitude});
    }

    function getMaxMagnitude(address operator, IStrategy strategy) external view returns (uint64) {
        return _maxMagnitudeHistory[operator][strategy].latest();
    }

    function getMaxMagnitudes(address operator, IStrategy[] calldata strategies) external view returns (uint64[] memory) {
        uint64[] memory maxMagnitudes = new uint64[](strategies.length);

        for (uint i = 0; i < strategies.length; ++i) {
            maxMagnitudes[i] = _maxMagnitudeHistory[operator][strategies[i]].latest();
        }

        return maxMagnitudes;
    }

    function getMaxMagnitudesAtBlock(address operator, IStrategy[] calldata strategies, uint32 blockNumber)
        external
        view
        returns (uint64[] memory)
    {
        uint64[] memory maxMagnitudes = new uint64[](strategies.length);

        for (uint i = 0; i < strategies.length; ++i) {
            maxMagnitudes[i] = _maxMagnitudeHistory[operator][strategies[i]].upperLookup({key: blockNumber});
        }

        return maxMagnitudes;
    }

    function setAVSSetCount(address avs, uint numSets) external {
        getOperatorSetCount[avs] = numSets;
    }
}
