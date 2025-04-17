// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/mocks/BeaconChainMock_Deneb.t.sol";
import "src/test/integration/IntegrationChecks.t.sol";

abstract contract UpgradeTest is IntegrationChecks {
    /// Only run upgrade tests on mainnet forks
    function setUp() public virtual override {
        // We do not run upgrade tests on local envs or non-mainnet forks...
        if (!forkConfig.supportUpgradeTests) {
            cheats.skip(true);
        } else {
            isUpgraded = false;
            _setUp(false);

            cheats.etch(address(beaconChain), type(BeaconChainMock_DenebForkable).runtimeCode);
            beaconChain.initialize(eigenPodManager(), BEACON_GENESIS_TIME);
        }
    }

    function _upgradeEigenLayerContracts() public virtual {
        require(!isUpgraded, "_upgradeEigenLayerContracts: already performed upgrade");

        _upgradeMainnetContracts();
        _handlePectraFork();

        // Bump block.timestamp forward to allow verifyWC proofs for migrated pods
        beaconChain.advanceEpoch_NoRewards();

        isUpgraded = true;
    }

    // Set the fork timestamp sufficiently in the future to keep using Deneb proofs
    // `Prooftra.t.sol` will handle the Deneb -> Pectra transition
    function _handlePectraFork() internal {
        // 1. Set proof timestamp setter to operations multisig
        cheats.prank(eigenPodManager().owner());
        eigenPodManager().setProofTimestampSetter(address(operationsMultisig()));

        // 2. Set Proof timestamp
        cheats.prank(eigenPodManager().proofTimestampSetter());
        eigenPodManager().setPectraForkTimestamp(type(uint64).max);
    }
}
