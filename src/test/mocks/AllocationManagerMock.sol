// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.9;

import "forge-std/Test.sol";
import "src/contracts/interfaces/IStrategy.sol";
import "src/contracts/libraries/Snapshots.sol";
import "src/contracts/libraries/OperatorSetLib.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

contract AllocationManagerMock is Test {
    address constant DEFAULT_BURN_ADDRESS = address(0x00000000000000000000000000000000000E16E4);

    using Snapshots for Snapshots.DefaultWadHistory;
    using OperatorSetLib for OperatorSet;

    receive() external payable {}
    fallback() external payable {}

    mapping(bytes32 operatorSetKey => bool) public _isOperatorSet;
    mapping(address avs => uint) public getOperatorSetCount;
    mapping(address => mapping(IStrategy => Snapshots.DefaultWadHistory)) internal _maxMagnitudeHistory;
    mapping(address => address) internal _avsRegistrar;
    mapping(bytes32 operatorSetKey => address) public _getRedistributionRecipient;
    mapping(bytes32 operatorSetKey => uint) public _getSlashCount;
    mapping(bytes32 operatorSetKey => address[] members) internal _members;
    mapping(bytes32 operatorSetKey => IStrategy[] strategies) internal _strategies;
    mapping(bytes32 operatorSetKey => mapping(address operator => mapping(IStrategy strategy => uint minimumSlashableStake))) internal
        _minimumSlashableStake;

    function getSlashCount(OperatorSet memory operatorSet) external view returns (uint) {
        return _getSlashCount[operatorSet.key()];
    }

    function setSlashCount(OperatorSet memory operatorSet, uint slashCount) external {
        _getSlashCount[operatorSet.key()] = slashCount;
    }

    function getRedistributionRecipient(OperatorSet memory operatorSet) external view returns (address recipient) {
        recipient = _getRedistributionRecipient[operatorSet.key()];

        if (recipient == address(0)) recipient = DEFAULT_BURN_ADDRESS;
    }

    function setRedistributionRecipient(OperatorSet memory operatorSet, address recipient) external {
        _getRedistributionRecipient[operatorSet.key()] = recipient;
    }

    function isRedistributingOperatorSet(OperatorSet memory operatorSet) external view returns (bool) {
        return true;
    }

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

    function setAVSRegistrar(address avs, address avsRegistrar) external {
        _avsRegistrar[avs] = avsRegistrar;
    }

    function getAVSRegistrar(address avs) external view returns (address) {
        return _avsRegistrar[avs];
    }

    function getMembers(OperatorSet memory operatorSet) external view returns (address[] memory) {
        return _members[operatorSet.key()];
    }

    function setMembersInOperatorSet(OperatorSet memory operatorSet, address[] memory members) external {
        _members[operatorSet.key()] = members;
    }

    function setStrategiesInOperatorSet(OperatorSet memory operatorSet, IStrategy[] memory strategies) external {
        _strategies[operatorSet.key()] = strategies;
    }

    function getStrategiesInOperatorSet(OperatorSet memory operatorSet) external view returns (IStrategy[] memory) {
        return _strategies[operatorSet.key()];
    }

    function setMinimumSlashableStake(
        OperatorSet memory operatorSet,
        address[] memory operators,
        IStrategy[] memory strategies,
        uint[][] memory minimumSlashableStake
    ) external {
        for (uint i = 0; i < operators.length; ++i) {
            for (uint j = 0; j < strategies.length; ++j) {
                _minimumSlashableStake[operatorSet.key()][operators[i]][strategies[j]] = minimumSlashableStake[i][j];
            }
        }
    }

    function getMinimumSlashableStake(
        OperatorSet memory operatorSet,
        address[] memory operators,
        IStrategy[] memory strategies /*uint32 futureBlock*/
    ) external pure returns (uint[][] memory) {
        uint[][] memory minimumSlashableStake = new uint[][](operators.length);

        for (uint i = 0; i < operators.length; ++i) {
            minimumSlashableStake[i] = new uint[](strategies.length);
        }

        return minimumSlashableStake;
    }
}
