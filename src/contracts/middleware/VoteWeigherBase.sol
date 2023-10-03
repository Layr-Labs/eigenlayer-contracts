// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../interfaces/IStrategyManager.sol";
import "./VoteWeigherBaseStorage.sol";

/**
 * @title A simple implementation of the `IVoteWeigher` interface.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice This contract is used for
 * - computing the total weight of an operator for any of the quorums that are considered
 * by the middleware
 * - addition and removal of strategies and the associated weighting criteria that are assigned
 * by the middleware for each of the quorum(s)
 * @dev
 */
contract VoteWeigherBase is VoteWeigherBaseStorage {
    /// @notice when applied to a function, ensures that the function is only callable by the current `owner` of the `serviceManager`
    modifier onlyServiceManagerOwner() {
        require(msg.sender == serviceManager.owner(), "VoteWeigherBase.onlyServiceManagerOwner: caller is not the owner of the serviceManager");
        _;
    }

    /// @notice when applied to a function, ensures that the `quorumNumber` corresponds to a valid quorum added to the VoteWeigher
    modifier validQuorumNumber(uint8 quorumNumber) {
        require(quorumNumber < quorumCount, "VoteWeigherBase.validQuorumNumber: quorumNumber is not valid");
        _;
    }

    /// @notice Sets the (immutable) `strategyManager` and `serviceManager` addresses
    constructor(
        IStrategyManager _strategyManager,
        IServiceManager _serviceManager
    ) VoteWeigherBaseStorage(_strategyManager, _serviceManager) {}

    /// @notice Returns the strategy and weight multiplier for the `index`'th strategy in the quorum `quorumNumber`
    function strategyAndWeightingMultiplierForQuorumByIndex(uint8 quorumNumber, uint256 index)
        public
        view
        returns (StrategyAndWeightingMultiplier memory)
    {
        return strategiesConsideredAndMultipliers[quorumNumber][index];
    }

    /**
     * @notice This function computes the total weight of the @param operator in the quorum @param quorumNumber.
     * @dev reverts in the case that `quorumNumber` is greater than or equal to `quorumCount`
     */
    function weightOfOperatorForQuorumView(uint8 quorumNumber, address operator) public virtual view validQuorumNumber(quorumNumber) returns (uint96) {
        uint96 weight;
        uint256 stratsLength = strategiesConsideredAndMultipliersLength(quorumNumber);
        StrategyAndWeightingMultiplier memory strategyAndMultiplier;

        for (uint256 i = 0; i < stratsLength;) {
            // accessing i^th StrategyAndWeightingMultiplier struct for the quorumNumber
            strategyAndMultiplier = strategiesConsideredAndMultipliers[quorumNumber][i];

            // shares of the operator in the strategy
            uint256 sharesAmount = delegation.operatorShares(operator, strategyAndMultiplier.strategy);

            // add the weight from the shares for this strategy to the total weight
            if (sharesAmount > 0) {
                weight += uint96(sharesAmount * strategyAndMultiplier.multiplier / WEIGHTING_DIVISOR);
            }

            unchecked {
                ++i;
            }
        }

        return weight;
    }

    /**
     * @notice This function computes the total weight of the @param operator in the quorum @param quorumNumber.
     * @dev reverts in the case that `quorumNumber` is greater than or equal to `quorumCount`
     * @dev a version of weightOfOperatorForQuorumView that can change state if needed
     */
    function weightOfOperatorForQuorum(uint8 quorumNumber, address operator) public virtual validQuorumNumber(quorumNumber) returns (uint96) {
        return weightOfOperatorForQuorumView(quorumNumber, operator);
    }

    /// @notice Create a new quorum and add the strategies and their associated weights to the quorum.
    function createQuorum(
        StrategyAndWeightingMultiplier[] memory _strategiesConsideredAndMultipliers
    ) external virtual onlyServiceManagerOwner {
        _createQuorum(_strategiesConsideredAndMultipliers);
    }

    /// @notice Adds new strategies and the associated multipliers to the @param quorumNumber.
    function addStrategiesConsideredAndMultipliers(
        uint8 quorumNumber,
        StrategyAndWeightingMultiplier[] memory _newStrategiesConsideredAndMultipliers
    ) external virtual onlyServiceManagerOwner validQuorumNumber(quorumNumber) {
        _addStrategiesConsideredAndMultipliers(quorumNumber, _newStrategiesConsideredAndMultipliers);
    }

    /**
     * @notice This function is used for removing strategies and their associated weights from the
     * mapping strategiesConsideredAndMultipliers for a specific @param quorumNumber.
     * @dev higher indices should be *first* in the list of @param indicesToRemove, since otherwise
     * the removal of lower index entries will cause a shift in the indices of the other strategiesToRemove
     */
    function removeStrategiesConsideredAndMultipliers(
        uint8 quorumNumber,
        uint256[] calldata indicesToRemove
    ) external virtual onlyServiceManagerOwner validQuorumNumber(quorumNumber) {
        uint256 indicesToRemoveLength = indicesToRemove.length;
        require(indicesToRemoveLength > 0, "VoteWeigherBase.removeStrategiesConsideredAndMultipliers: no indices to remove provided");
        for (uint256 i = 0; i < indicesToRemoveLength;) {
            emit StrategyRemovedFromQuorum(quorumNumber, strategiesConsideredAndMultipliers[quorumNumber][indicesToRemove[i]].strategy);
            // remove strategy and its associated multiplier
            strategiesConsideredAndMultipliers[quorumNumber][indicesToRemove[i]] = strategiesConsideredAndMultipliers[
                quorumNumber
            ][strategiesConsideredAndMultipliers[quorumNumber].length - 1];
            strategiesConsideredAndMultipliers[quorumNumber].pop();

            unchecked {
                ++i;
            }
        }
    }

    /**
     * @notice This function is used for modifying the weights of strategies that are already in the
     * mapping strategiesConsideredAndMultipliers for a specific
     * @param quorumNumber is the quorum number to change the strategy for
     * @param strategyIndices are the indices of the strategies to change
     * @param newMultipliers are the new multipliers for the strategies
     */
    function modifyStrategyWeights(
        uint8 quorumNumber,
        uint256[] calldata strategyIndices,
        uint96[] calldata newMultipliers
    ) external virtual onlyServiceManagerOwner validQuorumNumber(quorumNumber) {
        uint256 numStrats = strategyIndices.length;
        require(numStrats > 0, "VoteWeigherBase.modifyStrategyWeights: no strategy indices provided");
        // sanity check on input lengths
        require(newMultipliers.length == numStrats, "VoteWeigherBase.modifyStrategyWeights: input length mismatch");

        for (uint256 i = 0; i < numStrats; ) {
            // change the strategy's associated multiplier
            strategiesConsideredAndMultipliers[quorumNumber][strategyIndices[i]].multiplier = newMultipliers[i];
            emit StrategyMultiplierUpdated(quorumNumber, strategiesConsideredAndMultipliers[quorumNumber][strategyIndices[i]].strategy, newMultipliers[i]);
            unchecked {
                ++i;
            }
        }
    }

    /// @notice Returns the length of the dynamic array stored in `strategiesConsideredAndMultipliers[quorumNumber]`.
    function strategiesConsideredAndMultipliersLength(uint8 quorumNumber) public view returns (uint256) {
        return strategiesConsideredAndMultipliers[quorumNumber].length;
    }

    /**
     * @notice Creates a quorum with the given_strategiesConsideredAndMultipliers.
     */
    function _createQuorum(
        StrategyAndWeightingMultiplier[] memory _strategiesConsideredAndMultipliers
    ) internal {
        uint16 quorumCountMem = quorumCount;
        require(quorumCountMem < 192, "VoteWeigherBase._createQuorum: number of quorums cannot 192");
        uint8 quorumNumber = uint8(quorumCountMem);
        // increment quorumCount
        quorumCount = quorumCountMem + 1;
        // add the strategies and their associated weights to the quorum
        _addStrategiesConsideredAndMultipliers(quorumNumber, _strategiesConsideredAndMultipliers);
        // emit event
        emit QuorumCreated(quorumNumber);
    }

    /** 
     * @notice Adds `_newStrategiesConsideredAndMultipliers` to the `quorumNumber`-th quorum.
     * @dev Checks to make sure that the *same* strategy cannot be added multiple times (checks against both against existing and new strategies).
     * @dev This function has no check to make sure that the strategies for a single quorum have the same underlying asset. This is a concious choice,
     * since a middleware may want, e.g., a stablecoin quorum that accepts USDC, USDT, DAI, etc. as underlying assets and trades them as "equivalent".
     */
    function _addStrategiesConsideredAndMultipliers(
        uint8 quorumNumber,
        StrategyAndWeightingMultiplier[] memory _newStrategiesConsideredAndMultipliers
    ) internal {
        require(_newStrategiesConsideredAndMultipliers.length > 0, "VoteWeigherBase._addStrategiesConsideredAndMultipliers: no strategies provided");
        uint256 numStratsToAdd = _newStrategiesConsideredAndMultipliers.length;
        uint256 numStratsExisting = strategiesConsideredAndMultipliers[quorumNumber].length;
        require(
            numStratsExisting + numStratsToAdd <= MAX_WEIGHING_FUNCTION_LENGTH,
            "VoteWeigherBase._addStrategiesConsideredAndMultipliers: exceed MAX_WEIGHING_FUNCTION_LENGTH"
        );
        for (uint256 i = 0; i < numStratsToAdd; ) {
            // fairly gas-expensive internal loop to make sure that the *same* strategy cannot be added multiple times
            for (uint256 j = 0; j < (numStratsExisting + i); ) {
                require(
                    strategiesConsideredAndMultipliers[quorumNumber][j].strategy !=
                        _newStrategiesConsideredAndMultipliers[i].strategy,
                    "VoteWeigherBase._addStrategiesConsideredAndMultipliers: cannot add same strategy 2x"
                );
                unchecked {
                    ++j;
                }
            }
            require(
                _newStrategiesConsideredAndMultipliers[i].multiplier > 0,
                "VoteWeigherBase._addStrategiesConsideredAndMultipliers: cannot add strategy with zero weight"
            );
            strategiesConsideredAndMultipliers[quorumNumber].push(_newStrategiesConsideredAndMultipliers[i]);
            emit StrategyAddedToQuorum(quorumNumber, _newStrategiesConsideredAndMultipliers[i].strategy, _newStrategiesConsideredAndMultipliers[i].multiplier);
            unchecked {
                ++i;
            }
        }
    }
}
