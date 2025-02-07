// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

interface IMintableToken {
    function mint(address recipient, uint256 amount) external;
}

interface IRecipient {
    function funder() external view returns (address);
    function onMint(address token, uint256 amount) external;
}

uint256 constant TIMESCALE = 4 weeks;

// @notice Determines the relative size of a substream within a token stream, and accounts for the tokens already minted by the substream
struct Substream {
    uint256 weight;
    uint256 rewardDebt;
}

/**
 * @notice Defines a token stream, which may contain any number of substreams, each accumulating a fraction of the `rate`
 * @dev Non-normalization refers to the fact that totalWeight (the divisor for substream size) is floating, not fixed
 * i.e. substreams are a fraction of the total stream size, with the fraction defined by (substreams[address] / totalWeight)
 */
struct NonNormalizedStream {
    // @dev note that this is in terms of TIMESCALE
    uint256 rate;
    // @dev note that this is in terms of TIMESCALE
    uint256 lastUpdated;
    // TODO: enforce some max total weight?
    uint256 totalWeight;
    uint256 scaledCumulativeRewardDebtPerWeight;
    mapping(address => Substream) substreams;
}

uint256 constant NORMALIZED_STREAM_TOTAL_WEIGHT = 1e18;

/**
 * @notice Defines a token stream, which may contain any number of substreams, each accumulating a fraction of the `rate`
 * @dev This is identified as normalized because the totalWeight is designed to be fixed at NORMALIZED_STREAM_TOTAL_WEIGHT.
 * Each substream is a fraction of the total stream size, with the fraction defined by (substreams[address] / NORMALIZED_STREAM_TOTAL_WEIGHT)
 */
struct NormalizedStream {
    // @dev note that this is in terms of TIMESCALE
    uint256 rate;
    // @dev note that this is in terms of TIMESCALE
    uint256 lastUpdated;
    uint256 unassignedWeight;
    uint256 scaledCumulativeRewardDebtPerWeight;
    mapping(address => Substream) substreams;
}


// TODO: deal with totalWeight of zero!
// TODO: events
library StreamMath {

// functions for NonNormalizedStream
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





// functions for NormalizedStream

    function substreamRate(
        NormalizedStream storage stream,
        address substreamRecipient
    ) internal view returns (uint256) {
        Substream storage substream = stream.substreams[substreamRecipient];

        return stream.rate * substream.weight / NORMALIZED_STREAM_TOTAL_WEIGHT;
    }

    function pendingAmountToClaim(
        NormalizedStream storage stream,
        address substreamRecipient
    ) internal view returns (uint256) {
        Substream storage substream = stream.substreams[substreamRecipient];

        // calculate pending increase to scaled cumulative reward debt per weight
        uint256 _scaledCumulativeRewardDebtPerWeight = stream.scaledCumulativeRewardDebtPerWeight;
        _scaledCumulativeRewardDebtPerWeight = _scaledCumulativeRewardDebtPerWeight + 
            ((1e18 * ((block.timestamp / TIMESCALE) - stream.lastUpdated) * stream.rate) / NORMALIZED_STREAM_TOTAL_WEIGHT);
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
    function updateStream(NormalizedStream storage stream) internal returns (uint256) {
        // increase scaled cumulative reward debt per weight
        uint256 _scaledCumulativeRewardDebtPerWeight = stream.scaledCumulativeRewardDebtPerWeight;
        // TODO: event
        _scaledCumulativeRewardDebtPerWeight = _scaledCumulativeRewardDebtPerWeight + 
            ((1e18 * ((block.timestamp / TIMESCALE) - stream.lastUpdated) * stream.rate) / NORMALIZED_STREAM_TOTAL_WEIGHT);
        stream.scaledCumulativeRewardDebtPerWeight = _scaledCumulativeRewardDebtPerWeight;
        stream.lastUpdated = block.timestamp / TIMESCALE;
        return _scaledCumulativeRewardDebtPerWeight;
    }


    function claimForSubstream(
        NormalizedStream storage stream,
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
        NormalizedStream storage stream,
        address substreamRecipient,
        uint256 newWeight,
        address streamToken
    ) internal {
        Substream storage substream = stream.substreams[substreamRecipient];

        claimForSubstream(stream, substreamRecipient, streamToken);
        // TODO: event
        uint256 _unassignedWeight = stream.unassignedWeight + substream.weight - newWeight;
        stream.unassignedWeight = _unassignedWeight;

        // TODO: event
        substream.weight = newWeight;
        // adjust rewardDebt to make pending rewards correct
        uint256 cumulativeAmount = (newWeight * stream.scaledCumulativeRewardDebtPerWeight) / 1e18;
        substream.rewardDebt = cumulativeAmount;
    }
}






















