// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/IStrategy.sol";

contract AllocatorLogic_Sketch {

    struct OperatorSet {
        address avs;
        uint32 id;
    }

    // Mapping: operator => Strategy => list of active allocators (some debate about whether this should be OpSet-specific)
    mapping(address => mapping(IStrategy => address[])) public activeAllocators;

    // Mapping: allocator => Strategy => list of operators with magnitude allocated to them
    mapping(address => mapping(IStrategy => address[])) public allocatedOperators;

    // Mapping: allocator => Strategy => operator => OpSet (i.e. AVS, ID) => magnitude
    mapping(address => mapping(IStrategy => mapping(address => mapping(bytes32 => uint256)))) public allocatedMagnitudes;

    // Mapping: allocator => Strategy => totalMagnitude (summed over operators & OpSets)
    mapping(address => mapping(IStrategy => uint256)) public totalMagnitude;

    mapping(address => mapping(IStrategy => uint256)) public nonSlashableMagnitude;

    // Mapping: allocator => strategy => operator => totalAllocatedMagnitudeToOperator (used for multiplication when updated delegatedShares)
    mapping(address => mapping(IStrategy => mapping(address => uint256))) public totalMagnitudeAllocatedToOperator;

    // This is summed over OpSets but for a single operator
    // Mapping: operator => strategy => delegatedShares (from allocators)
    mapping(address => mapping(IStrategy => uint256)) public delegatedShares;

    // Mapping: allocator => strategy => delegatedShares (from stakers)
    mapping(address => mapping(IStrategy => uint256)) public allocatorShares;

    // Mapping: operator => allocator (approved or not)
    mapping(address => mapping(address => bool)) public approvedToAllocate;

    function getSlashableStake(address operator, IStrategy strategy, bytes32 opSet) public view returns (uint256 total) {
        address[] memory allocatorList = activeAllocators[operator][strategy];
        for (uint256 i = 0; i < allocatorList.length; ++i) {
            address allocator = allocatorList[i];
            uint256 allocatedMagnitude = allocatedMagnitudes[allocator][strategy][operator][opSet];
            uint256 magnitudeDivisor = totalMagnitude[allocator][strategy];
            uint256 userShares = allocatorShares[allocator][strategy]; 
            total += (userShares * allocatedMagnitude) / magnitudeDivisor;
        }
        return total;
    }

    function _slashOperator(IStrategy strategy, address operator, bytes32 opSet, uint256 bipsToSlash) internal {
        address[] memory allocatorList = activeAllocators[operator][strategy];
        for (uint256 i = 0; i < allocatorList.length; ++i) {
            address allocator = allocatorList[i];
            uint256 allocatedMagnitude = allocatedMagnitudes[allocator][strategy][operator][opSet];
            uint256 magnitudeDivisor = totalMagnitude[allocator][strategy];
            // TODO -- NEED TO ROUND THIS *UP*, NOT DOWN
            uint256 magnitudeToDecrement = (bipsToSlash * allocatedMagnitude) / (magnitudeDivisor * 10000);
            _decreaseMagnitude({
                allocator: allocator,
                strategy: strategy,
                operator: operator,
                opSet:opSet,
                amountToDecrease: magnitudeToDecrement
            });
            uint256 updatedTotalMagnitude = magnitudeDivisor - magnitudeToDecrement;
            totalMagnitude[allocator][strategy] = updatedTotalMagnitude;
        }
    }

    function _onDeposit(address allocator, IStrategy strategy, uint256 shares) internal {
        // perform delegated shares update
        address[] memory operatorList = allocatedOperators[allocator][strategy];
        for (uint256 i = 0; i < operatorList.length; ++i) {
            address operator = operatorList[i];
            uint256 operatorMagnitude = totalMagnitudeAllocatedToOperator[allocator][strategy][operator];
            delegatedShares[operator][strategy] += (shares * operatorMagnitude) / totalMagnitude[allocator][strategy];
        }
    }

    function _onWithdrawal(address allocator, IStrategy strategy, uint256 shares) internal {
        // perform delegated shares update
        address[] memory operatorList = allocatedOperators[allocator][strategy];
        for (uint256 i = 0; i < operatorList.length; ++i) {
            address operator = operatorList[i];
            uint256 operatorMagnitude = totalMagnitudeAllocatedToOperator[allocator][strategy][operator];
            delegatedShares[operator][strategy] -= (shares * operatorMagnitude) / totalMagnitude[allocator][strategy];
        }
    }

    function _increaseMagnitude(address allocator, IStrategy strategy, address operator, bytes32 opSet, uint256 amountToIncrease) internal {
        uint256 nonSlashableMagnitudeBefore = nonSlashableMagnitude[allocator][strategy];
        uint256 nonSlashableMagnitudeAfter = nonSlashableMagnitudeBefore - amountToIncrease;
        uint256 allocatedMagnitudeBefore = allocatedMagnitudes[allocator][strategy][operator][opSet];
        uint256 allocatedMagnitudeAfter = allocatedMagnitudeBefore + amountToIncrease;
        uint256 totalMagnitudeAllocatedToOperatorBefore = totalMagnitudeAllocatedToOperator[allocator][strategy][operator];
        uint256 totalMagnitudeAllocatedToOperatorAfter = totalMagnitudeAllocatedToOperatorBefore + amountToIncrease;

        // perform delegated shares update
        uint256 userShares = allocatorShares[allocator][strategy]; 
        delegatedShares[operator][strategy] += (userShares * amountToIncrease) / totalMagnitude[allocator][strategy];

        // totalMagnitude doesn't change, since this is just a move from non-slashable to slashable
        nonSlashableMagnitude[allocator][strategy] = nonSlashableMagnitudeAfter;
        allocatedMagnitudes[allocator][strategy][operator][opSet] = allocatedMagnitudeAfter;
        totalMagnitudeAllocatedToOperator[allocator][strategy][operator] = totalMagnitudeAllocatedToOperatorAfter;

        if (totalMagnitudeAllocatedToOperatorBefore == 0 && amountToIncrease != 0) {
            if (!addressIsInList(allocator, activeAllocators[operator][strategy])) {
                // these two lists are always updated in tandem. they represent two views of the same information
                activeAllocators[operator][strategy].push(allocator);
                allocatedOperators[allocator][strategy].push(operator);
            }
        }
    }

    function _decreaseMagnitude(address allocator, IStrategy strategy, address operator, bytes32 opSet, uint256 amountToDecrease) internal {
        uint256 nonSlashableMagnitudeBefore = nonSlashableMagnitude[allocator][strategy];
        uint256 nonSlashableMagnitudeAfter = nonSlashableMagnitudeBefore + amountToDecrease;
        uint256 allocatedMagnitudeBefore = allocatedMagnitudes[allocator][strategy][operator][opSet];
        uint256 allocatedMagnitudeAfter = allocatedMagnitudeBefore - amountToDecrease;
        uint256 totalMagnitudeAllocatedToOperatorBefore = totalMagnitudeAllocatedToOperator[allocator][strategy][operator];
        uint256 totalMagnitudeAllocatedToOperatorAfter = totalMagnitudeAllocatedToOperatorBefore - amountToDecrease;

        // perform delegated shares update
        uint256 userShares = allocatorShares[allocator][strategy]; 
        delegatedShares[operator][strategy] -= (userShares * amountToDecrease) / totalMagnitude[allocator][strategy];

        // totalMagnitude doesn't change, since this is just a move from non-slashable to slashable
        nonSlashableMagnitude[allocator][strategy] = nonSlashableMagnitudeAfter;
        allocatedMagnitudes[allocator][strategy][operator][opSet] = allocatedMagnitudeAfter;
        totalMagnitudeAllocatedToOperator[allocator][strategy][operator] = totalMagnitudeAllocatedToOperatorAfter;

        if (totalMagnitudeAllocatedToOperatorAfter == 0 && totalMagnitudeAllocatedToOperatorBefore != 0) {
            // these two lists are always updated in tandem. they represent two views of the same information
            uint256 allocatorListLocation = findLocationInList(allocator, activeAllocators[operator][strategy]);
            activeAllocators[operator][strategy][allocatorListLocation] =
                activeAllocators[operator][strategy][activeAllocators[operator][strategy].length - 1];
            activeAllocators[operator][strategy].pop();

            uint256 operatorListLocation = findLocationInList(allocator, allocatedOperators[allocator][strategy]);
            allocatedOperators[allocator][strategy][operatorListLocation] = 
                allocatedOperators[allocator][strategy][allocatedOperators[allocator][strategy].length - 1];
            allocatedOperators[allocator][strategy].pop();
        }
    }

    function addressIsInList(address toCheck, address[] memory list) public pure returns (bool) {
        for (uint256 i = 0; i < list.length; ++i) {
            if (toCheck == list[i]) {
                return true;
            }
        }
        return false;
    }


    function findLocationInList(address toCheck, address[] memory list) public pure returns (uint256) {
        for (uint256 i = 0; i < list.length; ++i) {
            if (toCheck == list[i]) {
                return i;
            }
        }
        revert("not in list");
    }
}
