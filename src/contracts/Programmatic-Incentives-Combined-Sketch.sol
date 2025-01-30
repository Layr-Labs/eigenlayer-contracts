// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

struct VectorEntry {
    address key;
    uint96 value;
}

struct VectorStorage {
    VectorEntry[] vector;
    mapping(address => bool) isInVector;
}

library LinearWeightedVector {

    error CannotAddDuplicateKey();
    error KeyNotInVector();
    error IncorrectIndexInput();

    event EntryAddedToVector(address indexed newKey, uint96 value);
    event KeyRemovedFromVector(address indexed keyRemoved);

    function addEntry(VectorStorage storage vectorStorage, VectorEntry memory newEntry) internal {
        require(!vectorStorage.isInVector[newEntry.key], CannotAddDuplicateKey());
        emit EntryAddedToVector(newEntry.key, newEntry.value);
        vectorStorage.isInVector[newEntry.key] = true;
        vectorStorage.vector.push(newEntry);
    }

    function removeKey(VectorStorage storage vectorStorage, address keyToRemove, uint256 indexOfKey) internal {
        require(vectorStorage.isInVector[keyToRemove], KeyNotInVector());
        require(vectorStorage.vector[indexOfKey].key == keyToRemove, IncorrectIndexInput());
        // swap and pop
        emit KeyRemovedFromVector(keyToRemove);
        vectorStorage.isInVector[keyToRemove] = false;
        vectorStorage.vector[indexOfKey] = vectorStorage.vector[vectorStorage.vector.length - 1];
        vectorStorage.vector.pop();
    }

    function findKeyIndex(VectorStorage storage vectorStorage, address key) internal view returns (uint256) {
        require(vectorStorage.isInVector[key], KeyNotInVector());
        uint256 length = vectorStorage.vector.length;
        uint256 index = 0;
        for (; index < length; ++index) {
            if (key == vectorStorage.vector[index].key) {
                break;
            }
        }
        return index;
    }
}

abstract contract SingleLinearWeightedVectorUser {
    using LinearWeightedVector for *;

    error InputLengthMismatch();

    VectorStorage internal _vectorStorage;

    function vector() external view returns (VectorEntry[] memory) {
        return _vectorStorage.vector;
    }

    function isInVector(address key) external view returns (bool) {
        return _vectorStorage.isInVector[key];
    }

    function keyIndex(address key) external view returns (uint256) {
        return _vectorStorage.findKeyIndex(key);
    }

    function _addEntries(VectorEntry[] memory newEntries) internal {
        for (uint256 i = 0; i < newEntries.length; ++i) {
            _vectorStorage.addEntry(newEntries[i]);
        }
    }

    function _removeKeys(address[] memory keysToRemove, uint256[] memory keyIndices) internal {
        require(keysToRemove.length == keyIndices.length, InputLengthMismatch());
        for (uint256 i = 0; i < keysToRemove.length; ++i) {
            _vectorStorage.removeKey(keysToRemove[i], keyIndices[i]);
        }
    }
}

contract OwnableLinearWeightedVector is SingleLinearWeightedVectorUser,  Ownable {
    function addEntries(VectorEntry[] memory newEntries) external onlyOwner {
        _addEntries(newEntries);
    }

    function removeKeys(address[] memory keysToRemove, uint256[] memory keyIndices) external onlyOwner {
        _removeKeys(keysToRemove, keyIndices);
    }
}

contract ProgrammaticIncentivesConfig is Initializable, OwnableUpgradeable {

    event AVSAddedToWhitelist(address indexed avs);
    event AVSRemovedFromWhitelist(address indexed avs);

    // stream => bips out of total
    mapping(uint8 => uint256) public bipsPerStream;

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

    /// @notice emitted when a new stream is created
    event StreamCreated(uint8 indexed streamNumber);
}

// module for managing EIGEN token inflation
interface IEIGENInflationModule {
    // TODO: events

    // TODO: import IERC20
    // @notice the eigen token, which is minted by this module
    function EIGEN() external view returns (IERC20);

    // TBD: decide what timespan inflation rates are defined over. natural choices are one second or one year
    function maxInflationRate() external view returns (uint256);

    // @notice current total inflation rate
    function totalInflationRate() external view returns (uint256);

    // @notice the inflation rate going to the `recipient`
    function inflationRate(address recipient) external view returns (uint256);

    // @notice adjust the inflation rate going to the `recipient` to the `newInflationRate`
    function setInflationRate(address recipient, uint256 newInflationRate) external;

    // @notice the latest UTC timestamp at whichn `recipient` claimed their inflationary tokens
    function lastClaimedTimestamp(address recipient) external view returns (uint256);

    // @notice amount of tokens that will be minted to `recipient` when tokens are claimed for them
    function tokensToClaim(address recipient) external view returns (uint256);

    // TBD: access control
    // returns the amount of tokens claimed
    function claimTokens(address recipient) external returns (uint256);
}

// TODO: decide about substreams + how to manage
abstract contract EIGENInflationModule is IEIGENInflationModule, Initializable, OwnableUpgradeable {

    error MaxRateLessThanCurrentRate();
    error TotalRateGreaterThanMax();

    event MaxInflationRateSet(uint256 newMaxInflationRate);
    event TotalInflationRateSet(uint256 newTotalInflationRate);
    event InflationRateSet(address indexed recipient, uint256 newInflationRate);

    uint256 public maxInflationRate;
    uint256 public totalInflationRate;
    mapping(address => uint256) public inflationRate;
    mapping(address => uint256) public lastClaimedTimestamp;

    function setInflationRate(address recipient, uint256 newInflationRate) external {
        uint256 currentInflationRate = inflationRate[recipient];
        // TODO: store interim amount to mint
        uint256 newTotalInflationRate = totalInflationRate + newInflationRate - currentInflationRate;
        require(newTotalInflationRate <= maxInflationRate, TotalRateGreaterThanMax());
        emit InflationRateSet(recipient, newInflationRate);
        emit TotalInflationRateSet(newTotalInflationRate);
        inflationRate[recipient] = newInflationRate;
        totalInflationRate = newTotalInflationRate;
    }

    function setMaxInflationRate(uint256 newMaxInflationRate) external onlyOwner {
        _setMaxInflationRate(newMaxInflationRate);
    }

    function _setMaxInflationRate(uint256 newMaxInflationRate) internal {
        require(newMaxInflationRate >= totalInflationRate, MaxRateLessThanCurrentRate());
        emit MaxInflationRateSet(newMaxInflationRate);
        maxInflationRate = newMaxInflationRate;
    }
}

interface IMintableToken {
    function mint(address recipient, uint256 amount) external;
}

interface IRecipient {
    function funder() external view returns (address);
    function onMint(address token, uint256 amount) external;
}

uint256 constant TIMESCALE = 4 weeks;

struct Substream {
    uint256 weight;
    uint256 rewardDebt;
}

// @dev non-normalization refers to the fact that totalWeight (the divisor for substream size) is floating, not fixed
// i.e. substreams are a fraction of the total stream size, with the fraction defined by (substreams[address] / totalWeight)
struct NonNormalizedStream {
    address token;
    // @dev note that this is in terms of TIMESCALE
    uint256 rate;
    // @dev note that this is in terms of TIMESCALE
    uint256 lastUpdated;
    // TODO: enforce some max total weight?
    uint256 totalWeight;
    uint256 scaledCumulativeRewardDebtPerWeight;
    mapping(address => Substream) substreams;
}


// TODO: deal with totalWeight of zero!
// TODO: events
library StreamMath {

    function substreamRate(
        NonNormalizedStream storage stream,
        address substreamRecipient
    ) internal view returns (uint256) {
        Substream storage substream = stream.substreams[substreamRecipient];

        return stream.rate * substream.weight / stream.totalWeight;
    }

    function pendingAmountToClaim(
        NonNormalizedStream storage stream,
        address substreamRecipient
    ) internal view returns (uint256) {
        Substream storage substream = stream.substreams[substreamRecipient];

        // calculate pending increase to scaled cumulative reward debt per weight
        uint256 _scaledCumulativeRewardDebtPerWeight = stream.scaledCumulativeRewardDebtPerWeight;
        uint256 _totalWeight = stream.totalWeight;
        if (_totalWeight !=0 ) {
            _scaledCumulativeRewardDebtPerWeight = _scaledCumulativeRewardDebtPerWeight + 
                ((1e18 * ((block.timestamp / TIMESCALE) - stream.lastUpdated) * stream.rate) / _totalWeight);
        }   
        uint256 cumulativeAmount = (substream.weight * _scaledCumulativeRewardDebtPerWeight) / 1e18;
        uint256 _rewardDebt = substream.rewardDebt;
        // TODO: consider rounding and the possibility of underflows
        // TODO: will making cumulativeAmount and rewardDebt also scaled up help solve this?
        if (cumulativeAmount > _rewardDebt) {
            return (cumulativeAmount - _rewardDebt);
        } else {
            return 0;
        }
    }

    // @dev returns updated value of scaledCumulativeRewardDebtPerWeight
    function updateStream(NonNormalizedStream storage stream) internal returns (uint256) {
        // increase scaled cumulative reward debt per weight
        uint256 _scaledCumulativeRewardDebtPerWeight = stream.scaledCumulativeRewardDebtPerWeight;
        uint256 _totalWeight = stream.totalWeight;
        if (_totalWeight !=0 ) {
            // TODO: event
            _scaledCumulativeRewardDebtPerWeight = _scaledCumulativeRewardDebtPerWeight + 
                ((1e18 * ((block.timestamp / TIMESCALE) - stream.lastUpdated) * stream.rate) / _totalWeight);
            stream.scaledCumulativeRewardDebtPerWeight = _scaledCumulativeRewardDebtPerWeight;
            stream.lastUpdated = block.timestamp / TIMESCALE;
        }   
        return _scaledCumulativeRewardDebtPerWeight;
    }

    function claimForSubstream(
        NonNormalizedStream storage stream,
        address substreamRecipient,
        address streamToken
    ) internal {
        Substream storage substream = stream.substreams[substreamRecipient];

        // increase scaled cumulative reward debt per weight
        uint256 _scaledCumulativeRewardDebtPerWeight = updateStream(stream);

        uint256 cumulativeAmount = (substream.weight * _scaledCumulativeRewardDebtPerWeight) / 1e18;
        uint256 _rewardDebt = substream.rewardDebt;
        // TODO: consider rounding and the possibility of underflows
        if (cumulativeAmount > _rewardDebt) {
            uint256 amountToMint = cumulativeAmount - _rewardDebt;
            IMintableToken(streamToken).mint(substreamRecipient, amountToMint);
            substream.rewardDebt = cumulativeAmount;
            // TODO: more on this interface?
            IRecipient(substreamRecipient).onMint(streamToken, amountToMint);
        }
    }

    function updateSubstreamWeight(
        NonNormalizedStream storage stream,
        address substreamRecipient,
        uint256 newWeight,
        address streamToken
    ) internal {
        Substream storage substream = stream.substreams[substreamRecipient];

        claimForSubstream(stream, substreamRecipient, streamToken);
        // TODO: event
        uint256 _totalWeight = stream.totalWeight - substream.weight + newWeight;
        stream.totalWeight = _totalWeight;

        // TODO: event
        substream.weight = newWeight;
        // adjust rewardDebt to make pending rewards correct
        uint256 cumulativeAmount = (newWeight * stream.scaledCumulativeRewardDebtPerWeight) / 1e18;
        substream.rewardDebt = cumulativeAmount;
    }

}

contract TokenInflationNexus is OwnableUpgradeable {
    using StreamMath for *;

    error MaxRateLessThanCurrentRate();
    error TotalRateGreaterThanMax();

    event MaxInflationRateSet(uint256 newMaxInflationRate);
    event TotalInflationRateSet(uint256 newTotalInflationRate);
    event StreamRateUpdated(uint256 indexed streamID, uint256 newStreamRate);

    address public bEIGEN;

    uint256 public nextStreamID;
    // @dev the implicit TIMESCALE for this is TIMESCALE
    uint256 public maxInflationRate;
    // @dev the implicit TIMESCALE for this is TIMESCALE
    uint256 public totalInflationRate;
    // mapping: streamID => Stream
    mapping(uint256 streamID => NonNormalizedStream stream) public streams;
    // mapping: address => streamID => access to edit substreams?
    mapping(address editor => mapping(uint256 streamID => bool hasAccess)) public hasSubstreamEditAccess;

    function createNewStream(
        // @dev note that this is in terms of the TIMESCALE
        uint256 rate
    ) external onlyOwner {
        uint256 streamID = nextStreamID;
        ++nextStreamID;
        streams[streamID].rate = rate;
        streams[streamID].lastUpdated = block.timestamp / TIMESCALE;
        // TODO: initiate substreams?
    }

// TODO: fix issue -- problem is that updating a stream with a long timespan may make the stream *about to mint a bunch of tokens*
// Ex: create new stream with 1 day duration right before turn of day, mint, then zero out that stream and create a new stream with (1 day + 1 minute) duration
    function _setStreamRate(uint256 streamID, uint256 newStreamRate) external onlyOwner {
        NonNormalizedStream storage stream = streams[streamID];

        // update stream, ignoring return value (we do not care about the new value of scaledCumulativeRewardDebtPerWeight)
        stream.updateStream();

        uint256 newTotalInflationRate = totalInflationRate + newStreamRate - stream.rate;
        require(newTotalInflationRate <= maxInflationRate, TotalRateGreaterThanMax());
        emit StreamRateUpdated(streamID, newStreamRate);
        emit TotalInflationRateSet(newTotalInflationRate);
        totalInflationRate = newTotalInflationRate;
        stream.rate = newStreamRate;
    }

    function setMaxInflationRate(uint256 newMaxInflationRate) external onlyOwner {
        _setMaxInflationRate(newMaxInflationRate);
    }

    function _setMaxInflationRate(uint256 newMaxInflationRate) internal {
        // TODO: alt implementation option is to allow reducing below current inflation, but require elsewhere that the new rate EITHER
        // is below the max OR is <= the current rate. The existing approach gives more 'strictly correct' behavior to the `maxInflationRate`
        require(newMaxInflationRate >= totalInflationRate, MaxRateLessThanCurrentRate());
        emit MaxInflationRateSet(newMaxInflationRate);
        maxInflationRate = newMaxInflationRate;
    }

    function claimForSubstream(
        uint256 streamID,
        address substreamRecipient
    ) external {
        // TODO: access control?
        streams[streamID].claimForSubstream(substreamRecipient, bEIGEN);
    }

    function updateSubstreamWeight(
        uint256 streamID,
        address substreamRecipient,
        uint256 newWeight
    ) external {
        // TODO: check auth!
        streams[streamID].updateSubstreamWeight(substreamRecipient, newWeight, bEIGEN);
    }

}






















