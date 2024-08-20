// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../interfaces/ISlasher.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IStrategyManager.sol";
import "../permissions/Pausable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";

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
    constructor(IStrategyManager, IDelegationManager) {}

    function initialize(address, IPauserRegistry, uint256) external {}

    function optIntoSlashing(address) external {}

    function freezeOperator(address) external {}

    function resetFrozenStatus(address[] calldata) external {}

    function recordFirstStakeUpdate(address, uint32) external {}

    function recordStakeUpdate(address, uint32, uint32, uint256) external {}

    function recordLastStakeUpdateAndRevokeSlashingAbility(address, uint32) external {}

    function strategyManager() external view returns (IStrategyManager) {}

    function delegation() external view returns (IDelegationManager) {}

    function isFrozen(address) external view returns (bool) {}

    function canSlash(address, address) external view returns (bool) {}

    function contractCanSlashOperatorUntilBlock(address, address) external view returns (uint32) {}

    function latestUpdateBlock(address, address) external view returns (uint32) {}

    function getCorrectValueForInsertAfter(address, uint32) external view returns (uint256) {}

    function canWithdraw(address, uint32, uint256) external returns (bool) {}

    function operatorToMiddlewareTimes(address, uint256) external view returns (MiddlewareTimes memory) {}

    function middlewareTimesLength(address) external view returns (uint256) {}

    function getMiddlewareTimesIndexStalestUpdateBlock(address, uint32) external view returns (uint32) {}

    function getMiddlewareTimesIndexServeUntilBlock(address, uint32) external view returns (uint32) {}

    function operatorWhitelistedContractsLinkedListSize(address) external view returns (uint256) {}

    function operatorWhitelistedContractsLinkedListEntry(
        address,
        address
    ) external view returns (bool, uint256, uint256) {}

    function whitelistedContractDetails(address, address) external view returns (MiddlewareDetails memory) {}
}
