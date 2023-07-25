// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

import "../interfaces/IStrategyManager.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/ISlasher.sol";
import "../interfaces/IBLSRegistryCoordinatorWithIndices.sol";
import "../interfaces/IAVSStateViewer.sol";
import "../interfaces/IVoteWeigher.sol";

import "../libraries/BitmapUtils.sol";

/**
 * @title Contract for an AVS viewer which should implement all of the functions required to display the AVS details on the EigenLayer UI.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 */
abstract contract AVSStateViewer is Initializable, IAVSStateViewer {

    IBLSRegistryCoordinatorWithIndices public immutable registryCoordinator;

    IVoteWeigher public immutable voteWeigher;

    constructor(
        IBLSRegistryCoordinatorWithIndices _registryCoordinator
    ) {
        registryCoordinator = _registryCoordinator;
        voteWeigher = IVoteWeigher(address(_registryCoordinator.stakeRegistry()));
        _disableInitializers();
    }

    /**
     * @notice Returns 
     * 1) an array of the quorums that are being served by an operator
     * 2) an array of strategies for each quorum in sequence in a 2 dimensional array
     * @param operator the address of the operator to get the quorum details for
     */
    function strategiesStakedForEachQuorum(address operator) external view returns (uint256[] memory, IStrategy[][] memory) {
        // get the operator id for the operator address
        bytes32 operatorId = registryCoordinator.getOperatorId(operator);
        // get the current quorum bitmap for the operator
        uint256 quorumBitmap = registryCoordinator.getCurrentQuorumBitmapByOperatorId(operatorId);
        // convert the quorum bitmap to an array of uint256s
        bytes memory quorumNumberBytes = BitmapUtils.bitmapToBytesArray(quorumBitmap);
        uint256[] memory quorumNumbers = new uint256[](quorumNumberBytes.length);
        for (uint256 i = 0; i < quorumNumberBytes.length; i++) {
            quorumNumbers[i] = uint256(uint8(quorumNumberBytes[i]));
        }

        // get the strategies for each quorum
        IStrategy[][] memory strategies = new IStrategy[][](quorumNumbers.length);
        for (uint256 i = 0; i < quorumNumbers.length; i++) {
            // get the number of strategies for the quorum
            uint256 strategiesAndWeightedMultipliers = voteWeigher.strategiesConsideredAndMultipliersLength(uint8(quorumNumbers[i]));
            strategies[i] = new IStrategy[](strategiesAndWeightedMultipliers);
            // get the strategies for the quorum
            for (uint j = 0; j < strategiesAndWeightedMultipliers; j++) {
                strategies[i][j] = voteWeigher.strategyAndWeightingMultiplierForQuorumByIndex(uint8(quorumNumbers[i]), j).strategy;
            }
        }

        return (quorumNumbers, strategies);
    }
}