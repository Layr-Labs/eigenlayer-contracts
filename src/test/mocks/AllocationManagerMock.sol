// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IStrategy.sol";
import "../../contracts/libraries/Snapshots.sol";

contract AllocationManagerMock is Test {
    using Snapshots for Snapshots.DefaultWadHistory;

    mapping(address => bool) internal _isAVS;

    receive() external payable {}
    fallback() external payable {}

    mapping(address => mapping(IStrategy => Snapshots.DefaultWadHistory)) internal _maxMagnitudeHistory;

    function setMaxMagnitudes(
        address operator,
        IStrategy[] calldata strategies,
        uint64[] calldata maxMagnitudes
    ) external {
        for (uint256 i = 0; i < strategies.length; ++i) {
            setMaxMagnitude(operator, strategies[i], maxMagnitudes[i]);
        }
    }

    function setMaxMagnitude(
        address operator,
        IStrategy strategy,
        uint64 maxMagnitude
    ) public {
        _maxMagnitudeHistory[operator][strategy].push({
            key: uint32(block.number),
            value: maxMagnitude
        });
    }

    function getMaxMagnitude(
        address operator,
        IStrategy strategy
    ) external view returns (uint64) {
        return _maxMagnitudeHistory[operator][strategy].latest();
    }

    function getMaxMagnitudes(
        address operator,
        IStrategy[] calldata strategies
    ) external view returns (uint64[] memory) {
        uint64[] memory maxMagnitudes = new uint64[](strategies.length);

        for (uint256 i = 0; i < strategies.length; ++i) {
            maxMagnitudes[i] = _maxMagnitudeHistory[operator][strategies[i]].latest();
        }

        return maxMagnitudes;
    }

    function getMaxMagnitudesAtBlock(
        address operator,
        IStrategy[] calldata strategies,
        uint32 blockNumber
    ) external view returns (uint64[] memory) {
        uint64[] memory maxMagnitudes = new uint64[](strategies.length);

        for (uint256 i = 0; i < strategies.length; ++i) {
            maxMagnitudes[i] = _maxMagnitudeHistory[operator][strategies[i]].upperLookup({
                key: blockNumber
            });
        }

        return maxMagnitudes;
    }

    function isAVS(address avs) external view returns (bool) {
        return _isAVS[avs];
    }

    function setIsAVS(address avs, bool isAVS) external {
        _isAVS[avs] = isAVS;
    }
}