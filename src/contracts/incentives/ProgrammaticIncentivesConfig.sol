// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "./Vectors.sol";

interface IProgrammaticIncentivesConfig {
    function vector(uint256 vectorIndex) external view returns (VectorEntry[] memory);
}

/**
 * @notice Central contract for managing the configuration of many Programmatic Incentives parameters.
 * Uses LinearVectors and the associated LinearVectorOps library to manage lists of strategies & weights used by Distributor contracts.
 */
contract ProgrammaticIncentivesConfig is Initializable, OwnableUpgradeable, IProgrammaticIncentivesConfig {
    using LinearVectorOps for *;

    event AVSAddedToWhitelist(address indexed avs);
    event AVSRemovedFromWhitelist(address indexed avs);

        // calculate EIGEN from bipsPerStream and total rate
    function rewardsCoordinator_distributeProgrammaticIncentives(uint8 streamNumber, uint256 amountEIGEN) external {
        // read list of whitelisted AVSs
        // pull list of strategies and multipliers
        // get distributions from all AVSs (including ghost distributions)
        // filter distributions for stream inclusion (some difficulty with partial inclusion)
        // calculate total distributions in stream for each AVS (using price oracles to convert distributed tokens into common basis)
        // sum AVS distributions to get grand total
        // use grand total to get an EIGEN/($ or ETH, whatever unit of account is) for the stream
        // multiply rate by AVS sum to get EIGEN for the AVS
        // distribute EIGEN to Stakers and Operators of the AVS based on Strategy Weights in stream (a la RewardsV1)
        // TODO: figure out if this is for the same period or "trailing" -- EIGEN incentives distribution is for stakers in _this period_ based on distributions in last period
    }
    // also exclude any AVSs from stream which do not value any of the strategies

    mapping(address => bool) public avsIsWhitelisted;

    function editAVSWhitelistStatus(address avs, bool newWhitelistStatus) external onlyOwner {
        if (avsIsWhitelisted[avs] && !newWhitelistStatus) {
            avsIsWhitelisted[avs] = false;
            emit AVSRemovedFromWhitelist(avs);
        } else if (!avsIsWhitelisted[avs] && newWhitelistStatus) {
            avsIsWhitelisted[avs] = true;
            emit AVSAddedToWhitelist(avs);
        }
    }

    error InputLengthMismatch();

    LinearVector[] internal _weightingVectors;

    function createNewVector(VectorEntry[] memory initialVectorEntries) external onlyOwner {
        _weightingVectors.push();
        _addEntries({
            vectorIndex: _weightingVectors.length - 1,
            newEntries: initialVectorEntries
        });
    }

    function numberOfWeightingVectors() external view returns (uint256) {
        return _weightingVectors.length;
    }

    function vector(uint256 vectorIndex) external view returns (VectorEntry[] memory) {
        return _weightingVectors[vectorIndex].vector;
    }

    function isInVector(uint256 vectorIndex, address key) external view returns (bool) {
        return _weightingVectors[vectorIndex].isInVector[key];
    }

    function keyIndex(uint256 vectorIndex, address key) external view returns (uint256) {
        return _weightingVectors[vectorIndex].findKeyIndex(key);
    }

    function _addEntries(uint256 vectorIndex, VectorEntry[] memory newEntries) internal {
        for (uint256 i = 0; i < newEntries.length; ++i) {
            _weightingVectors[vectorIndex].addEntry(newEntries[i]);
        }
    }

    function _removeKeys(uint256 vectorIndex, address[] memory keysToRemove, uint256[] memory keyIndices) internal {
        require(keysToRemove.length == keyIndices.length, InputLengthMismatch());
        for (uint256 i = 0; i < keysToRemove.length; ++i) {
            _weightingVectors[vectorIndex].removeKey(keysToRemove[i], keyIndices[i]);
        }
    }

    function addEntries(uint256 vectorIndex, VectorEntry[] memory newEntries) external onlyOwner {
        _addEntries(vectorIndex, newEntries);
    }

    function removeKeys(uint256 vectorIndex, address[] memory keysToRemove, uint256[] memory keyIndices) external onlyOwner {
        _removeKeys(vectorIndex, keysToRemove, keyIndices);
    }
}























