// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/mocks/BeaconChainMock.t.sol";
import "src/test/integration/mocks/LibProofGen.t.sol";

/// @notice A backwards-compatible BeaconChain Mock that updates containers & proofs from Deneb to Pectra
contract BeaconChainMock_DenebForkable is BeaconChainMock {
    using StdStyle for *;
    using print for *;
    using LibValidator for *;
    using LibProofGen for *;

    // Denotes whether the beacon chain has been forked to Pectra
    bool isPectra;

    // The timestamp of the Pectra hard fork
    uint64 public pectraForkTimestamp;

    constructor(EigenPodManager _eigenPodManager, uint64 _genesisTime) 
        BeaconChainMock(_eigenPodManager, _genesisTime) 
    {
        LibProofGen.useDencun();
    }

    function NAME() public pure override returns (string memory) {
        return "BeaconChain_PectraForkable";
    }

    /// @dev Always return 32 ETH in gwei
    function _getMaxEffectiveBalanceGwei(Validator storage v) internal override view returns (uint64) {
        return isPectra ? super._getMaxEffectiveBalanceGwei(v) : MIN_ACTIVATION_BALANCE_GWEI;
    }

    /// @notice Forks the beacon chain to Pectra
    /// @dev Test battery should warp to the fork timestamp after calling this method
    function forkToPectra(uint64 _pectraForkTimestamp) public {
        isPectra = true;
        LibProofGen.usePectra();

        cheats.warp(_pectraForkTimestamp);
        pectraForkTimestamp = _pectraForkTimestamp;
    }
}
