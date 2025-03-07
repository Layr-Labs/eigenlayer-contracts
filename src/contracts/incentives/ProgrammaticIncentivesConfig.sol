// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "./Vectors.sol";
import "./Streams.sol";

interface IProgrammaticIncentivesConfig {
    // events for rewards-boost whitelist
    event AVSAddedToWhitelist(address indexed avs);
    event AVSRemovedFromWhitelist(address indexed avs);
    event TokenAddedToWhitelist(address indexed token);
    event TokenRemovedFromWhitelist(address indexed token);

    error InputLengthMismatch();

    // @notice Returns the weighting vector at `vectorIndex` (currently used to compare Strategy shares)
    function vector(uint256 vectorIndex) external view returns (VectorEntry[] memory);

    /**
     * @notice Getter function for querying whether a `key` is in the `vectorIndex`-th vector
     * @dev Will not revert even in the event that the `vectorIndex`-th vector does not exist
     */
    function isInVector(uint256 vectorIndex, address key) external view returns (bool);

    // @notice Returns the total number of existing vectors
    function numberOfWeightingVectors() external view returns (uint256);

    /**
     * @notice Getter function for fetching the index of a `key` within the `vectorIndex`-th vector
     * @dev Reverts if the key is not in the vector
     */
    function keyIndex(uint256 vectorIndex, address key) external view returns (uint256);

    // TODO: determine if this function should be called by the recipient itself or can be called on behalf of them
    // @notice Claims all pending inflationary tokens for the `substreamRecipient`
    function claimForSubstream(address substreamRecipient) external;

    // contract owner-only functions
    // @notice mapping avs => whether or not the avs is considered for incentives types that are filtered by the AVS whitelist
    function avsIsWhitelisted(address avs) external view returns (bool);

    // @notice Called by the contract owner modify the whitelist status of an AVS
    function editAVSWhitelistStatus(address avs, bool newWhitelistStatus) external;

    // @notice mapping token => whether or not the token is considered for incentives types that are filtered by the token whitelist
    function tokenIsWhitelisted(address avs) external view returns (bool);

    // TODO: there may be additional inputs required, especially if e.g. this function is supposed to deploy a token oracle contract
    // @notice Called by the contract owner modify the whitelist status of a token
    function editTokenWhitelistStatus(address avs, bool newWhitelistStatus) external;

    // @notice Called by the contract owner to change the relative amount of inflation going to the `substreamRecipient`
    function updateSubstreamWeight(address substreamRecipient, uint256 newWeight) external;

    // @notice Called by the contract owner to create a new vector indicating how different tokens are weighted relative to one another
    function createNewVector(VectorEntry[] memory initialVectorEntries) external;

    // @notice Called by the contract owner to add additional entries to an existing vector
    function addEntries(uint256 vectorIndex, VectorEntry[] memory newEntries) external;

    // @notice Called by the contract owner to remove entries from an existing vector
    function removeKeys(uint256 vectorIndex, address[] memory keysToRemove, uint256[] memory keyIndices) external;


}

/**
 * @notice Central contract for managing the configuration of many Programmatic Incentives parameters.
 * Uses LinearVectors and the associated LinearVectorOps library to manage lists of strategies & weights used by Distributor contracts.
 */
contract ProgrammaticIncentivesConfig is Initializable, OwnableUpgradeable, IProgrammaticIncentivesConfig {
    using LinearVectorOps for *;
    using StreamMath for *;

    // 4% of initial supply of EIGEN
    uint256 constant internal CURRENT_YEARLY_INFLATION = 66945866731386400000000160;

    // @inheritdoc IProgrammaticIncentivesConfig
    mapping(address => bool) public avsIsWhitelisted;

    // @inheritdoc IProgrammaticIncentivesConfig
    mapping(address => bool) public tokenIsWhitelisted;

    LinearVector[] internal _weightingVectors;

    // @notice single inflationary stream of EIGEN tokens
    // TODO: figure out what automatic getter(s) for this look like!
    NonNormalizedStream public stream;

    address public bEIGEN;

    constructor() {
        _disableInitializers();
    }

    function initialize(address initialOwner, address _bEIGEN) external initializer {
        _transferOwnership(initialOwner);
        bEIGEN = _bEIGEN;
        stream.rate = CURRENT_YEARLY_INFLATION * TIMESCALE / 365 days;
        stream.lastUpdatedTimestamp = block.timestamp;
        // TODO: initiate substreams?
    }

    // @inheritdoc IProgrammaticIncentivesConfig
    function editAVSWhitelistStatus(address avs, bool newWhitelistStatus) external onlyOwner {
        if (avsIsWhitelisted[avs] && !newWhitelistStatus) {
            avsIsWhitelisted[avs] = false;
            emit AVSRemovedFromWhitelist(avs);
        } else if (!avsIsWhitelisted[avs] && newWhitelistStatus) {
            avsIsWhitelisted[avs] = true;
            emit AVSAddedToWhitelist(avs);
        }
    }

    // @inheritdoc IProgrammaticIncentivesConfig
    function editTokenWhitelistStatus(address token, bool newWhitelistStatus) external onlyOwner {
        if (tokenIsWhitelisted[token] && !newWhitelistStatus) {
            tokenIsWhitelisted[token] = false;
            emit TokenRemovedFromWhitelist(token);
        } else if (!tokenIsWhitelisted[token] && newWhitelistStatus) {
            tokenIsWhitelisted[token] = true;
            emit TokenAddedToWhitelist(token);
        }
    }

    // @inheritdoc IProgrammaticIncentivesConfig
    function claimForSubstream(address substreamRecipient) external {
        // TODO: access control? could just use msg.sender instead of having `substreamRecipient` input
        stream.claimForSubstream(substreamRecipient, bEIGEN);
    }

    // @inheritdoc IProgrammaticIncentivesConfig
    function updateSubstreamWeight(address substreamRecipient, uint256 newWeight) external onlyOwner {
        stream.updateSubstreamWeight(substreamRecipient, newWeight, bEIGEN);
    }

    // @inheritdoc IProgrammaticIncentivesConfig
    function createNewVector(VectorEntry[] memory initialVectorEntries) external onlyOwner {
        _weightingVectors.push();
        _addEntries({
            vectorIndex: _weightingVectors.length - 1,
            newEntries: initialVectorEntries
        });
    }

    // @inheritdoc IProgrammaticIncentivesConfig
    function numberOfWeightingVectors() external view returns (uint256) {
        return _weightingVectors.length;
    }

    // @inheritdoc IProgrammaticIncentivesConfig
    function vector(uint256 vectorIndex) external view returns (VectorEntry[] memory) {
        return _weightingVectors[vectorIndex].vector;
    }

    // @inheritdoc IProgrammaticIncentivesConfig
    function isInVector(uint256 vectorIndex, address key) external view returns (bool) {
        return _weightingVectors[vectorIndex].isInVector[key];
    }

    // @inheritdoc IProgrammaticIncentivesConfig
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

    // @inheritdoc IProgrammaticIncentivesConfig
    function addEntries(uint256 vectorIndex, VectorEntry[] memory newEntries) external onlyOwner {
        _addEntries(vectorIndex, newEntries);
    }

    // @inheritdoc IProgrammaticIncentivesConfig
    function removeKeys(uint256 vectorIndex, address[] memory keysToRemove, uint256[] memory keyIndices) external onlyOwner {
        _removeKeys(vectorIndex, keysToRemove, keyIndices);
    }
}























