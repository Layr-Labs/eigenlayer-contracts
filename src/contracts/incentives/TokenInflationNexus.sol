// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "../interfaces/IEigen.sol";
import "./Streams.sol";

// @notice Module for managing EIGEN token inflation
interface ITokenInflationNexus {

    error MaxRateLessThanCurrentRate();
    error TotalRateGreaterThanMax();
    error ZeroAddressForbidden();
    error OnlyStreamCreator();
    error NoSubstreamEditAccess();

    event MaxInflationRateSet(uint256 newMaxInflationRate);
    event TotalInflationRateSet(uint256 newTotalInflationRate);
    event StreamRateUpdated(uint256 indexed streamID, uint256 newStreamRate);

    event StreamCreatorSet(address indexed newStreamCreator);

    // @notice the bEIGEN token, which is minted by this module
    function bEIGEN() external view returns (address);

    function maxInflationRate() external view returns (uint256);

    // @notice current total inflation rate
    function totalInflationRate() external view returns (uint256);

    // @notice the inflation rate going to the `recipient`
    // TODO: write this function or equivalent
    // function inflationRate(address recipient) external view returns (uint256);

    // TODO: add more functions
}


// TODO: improved access control, or else make the maxInflationRate fixed
/**
 * @notice A module for managing bEIGEN token inflation. The contract owner can create and modify streams,
 * each having a monthly inflation rate, with the sum of all inflation rates not allowed to exceed the maxInflationRate.
 * Within each stream, substreams can be defined which determine the fraction of the stream going to a recipient address,
 * and do accounting to track pending mintable amounts for each substream.
 * @dev The fraction of minting rights that a substream is given is equal to (substream weight / stream totalWeight)
 * @dev This system has a "heartbeat" equal to the defined constant TIMESCALE; Changes to streams and substreams can be
 * _made_ at any time, but effectively become _live_ at the end of the current TIMESCALE. The beginning of TIMESCALE-unit
 * time is defined as the Unix Epoch (i.e. 00:00:00 UTC on 1 January 1970).
 */
contract TokenInflationNexus is Initializable, OwnableUpgradeable, ITokenInflationNexus {
    using StreamMath for *;

    address public bEIGEN;
    /**
     * @notice Can create new streams and edit existing streams and/or substreams.
     * @dev Cannot edit streams to exceed the `maxInflationRate` set by the contract owner.
     */
    address public streamCreator;

    uint256 public nextStreamID;
    // @dev the implicit TIMESCALE for this is TIMESCALE
    uint256 public maxInflationRate;
    // @dev the implicit TIMESCALE for this is TIMESCALE
    uint256 public totalInflationRate;
    // mapping: streamID => Stream
    mapping(uint256 streamID => NonNormalizedStream stream) public streams;
    // mapping: address => streamID => access to edit substreams?
    mapping(address editor => mapping(uint256 streamID => bool hasAccess)) internal _substreamEditAccessGranted;

    // @notice Getter function; `streamCreator` has access to edit all substreams
    function hasSubstreamEditAccess(address editor, uint256 streamID) public view returns (bool) {
        return (_substreamEditAccessGranted[editor][streamID] || editor == streamCreator);
    }


    modifier onlyStreamCreator() {
        require(msg.sender == streamCreator, OnlyStreamCreator());
        _;
    }

    function initialize(address _streamCreator) external initializer {
        _setStreamCreator(_streamCreator);
    }

    function _setStreamCreator(address _streamCreator) internal {
        require(_streamCreator != address(0), ZeroAddressForbidden());
        emit StreamCreatorSet(_streamCreator);
        streamCreator = _streamCreator;
    }

    function setStreamCreator(address _streamCreator) external onlyOwner {
        _setStreamCreator(_streamCreator);
    }


    function createNewStream(
        // @dev note that this is in terms of the TIMESCALE
        uint256 rate
    ) external onlyStreamCreator {
        uint256 streamID = nextStreamID;
        ++nextStreamID;
        streams[streamID].rate = rate;
        streams[streamID].lastUpdated = block.timestamp / TIMESCALE;
        // TODO: initiate substreams?
    }

    function setStreamRate(uint256 streamID, uint256 newStreamRate) external onlyStreamCreator {
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
        require(hasSubstreamEditAccess(msg.sender, streamID), NoSubstreamEditAccess());
        streams[streamID].updateSubstreamWeight(substreamRecipient, newWeight, bEIGEN);
    }

}






















