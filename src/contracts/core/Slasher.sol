// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/ISlasher.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IStrategyManager.sol";
import "../libraries/StructuredLinkedList.sol";
import "../permissions/Pausable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

import "../libraries/SlashingAccountingUtils.sol";

/**
 * @notice This contract is not in use as of the Eigenlayer M2 release.
 *
 * Although many contracts reference it as an immutable variable, they do not
 * interact with it and it is effectively dead code. The Slasher was originally
 * deployed during Eigenlayer M1, but remained paused and unused for the duration
 * of that release as well.
 *
 * Eventually, slashing design will be finalized and the Slasher will be finished
 * and more fully incorporated into the core contracts. For now, you can ignore this
 * file. If you really want to see what the deployed M1 version looks like, check
 * out the `init-mainnet-deployment` branch under "releases".
 *
 * This contract is a stub that maintains its original interface for use in testing
 * and deploy scripts. Otherwise, it does nothing.
 */ 
contract Slasher is Initializable, OwnableUpgradeable, ISlasher, Pausable {
     
    /**
     * @notice Mapping: operator => strategy => share scalingFactor,
     * stored in the "SHARE_CONVERSION_SCALE", i.e. scalingFactor = 2 * SHARE_CONVERSION_SCALE indicates a scalingFactor of "2".
     * Note that a value of zero is treated as one, since this means that the operator has never been slashed
     */
    mapping(address => mapping(IStrategy => uint256)) internal _shareScalingFactor;

    function shareScalingFactor(address operator, IStrategy strategy) public view returns (uint256) {
        uint256 scalingFactor = _shareScalingFactor[operator][strategy];
        if (scalingFactor == 0) {
            return SlashingAccountingUtils.SHARE_CONVERSION_SCALE;
        } else {
            return scalingFactor;
        }
    }

    // @notice Mapping: operator => strategy => epochs in which the strategy was slashed for the operator
    // TODO: note that since default will be 0, we should probably make the "first epoch" actually be epoch 1 or something
    mapping(address => mapping(IStrategy => int256[])) public slashedEpochHistory;
    function lastSlashed(address operator, IStrategy strategy) public view returns (int256) {
        uint256 slashedEpochHistoryLength = slashedEpochHistory[operator][strategy].length;
        if (slashedEpochHistoryLength == 0) {
            return 0;
        } else {
            return slashedEpochHistory[operator][strategy][slashedEpochHistoryLength - 1];
        }
    }

    function _slashShares(address operator, IStrategy strategy, int256 epoch, uint256 bipsToSlash) internal {
        require(bipsToSlash != 0, "cannot slash 0%");
        require(bipsToSlash < SlashingAccountingUtils.BIPS_FACTOR, "cannot slash more than 99.99% in a single epoch");
        int256 lastSlashedEpoch = lastSlashed(operator, strategy);
        // TODO: again note that we need something like the first epoch being epoch 1 here
        require(epoch > lastSlashedEpoch, "slashing must occur in strictly ascending epoch order");
        uint256 scalingFactorBefore = shareScalingFactor(operator, strategy);
        uint256 scalingFactorAfter;
        // deal with edge case of operator being slashed repeatedly, inflating scalingFactor to max uint size
        // TODO: figure out more nuanced / appropriate way to handle this 'edge case', e.g. deciding if deposits should be blocked when close to limit
        if (SlashingAccountingUtils.MAX_SCALING_FACTOR / scalingFactorBefore >= bipsToSlash) {
            scalingFactorAfter = type(uint256).max;
        } else {
            scalingFactorAfter = scalingFactorBefore * SlashingAccountingUtils.BIPS_FACTOR / (SlashingAccountingUtils.BIPS_FACTOR - bipsToSlash);
        }
        // update storage to reflect the slashing
        slashedEpochHistory[operator][strategy].push(epoch);
        _shareScalingFactor[operator][strategy] = scalingFactorAfter;
        // TODO: events!
    }

    constructor(IStrategyManager, IDelegationManager) {}

    function initialize(
        address,
        IPauserRegistry,
        uint256
    ) external {}

    function optIntoSlashing(address) external {}

    function freezeOperator(address) external {}

    function resetFrozenStatus(address[] calldata) external {}

    function recordFirstStakeUpdate(address, uint32) external {}

    function recordStakeUpdate(
        address,
        uint32,
        uint32,
        uint256
    ) external {}

    function recordLastStakeUpdateAndRevokeSlashingAbility(address, uint32) external {}

    function strategyManager() external view returns (IStrategyManager) {}

    function delegation() external view returns (IDelegationManager) {}

    function isFrozen(address) external view returns (bool) {}

    function canSlash(address, address) external view returns (bool) {}

    function contractCanSlashOperatorUntilBlock(
        address,
        address
    ) external view returns (uint32) {}

    function latestUpdateBlock(address, address) external view returns (uint32) {}

    function getCorrectValueForInsertAfter(address, uint32) external view returns (uint256) {}

    function canWithdraw(
        address,
        uint32,
        uint256
    ) external returns (bool) {}

    function operatorToMiddlewareTimes(
        address,
        uint256
    ) external view returns (MiddlewareTimes memory) {}

    function middlewareTimesLength(address) external view returns (uint256) {}

    function getMiddlewareTimesIndexStalestUpdateBlock(address, uint32) external view returns (uint32) {}

    function getMiddlewareTimesIndexServeUntilBlock(address, uint32) external view returns (uint32) {}

    function operatorWhitelistedContractsLinkedListSize(address) external view returns (uint256) {}

    function operatorWhitelistedContractsLinkedListEntry(
        address,
        address
    ) external view returns (bool, uint256, uint256) {}

    function whitelistedContractDetails(
        address,
        address
    ) external view returns (MiddlewareDetails memory) {}

}
