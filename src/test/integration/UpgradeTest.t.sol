// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/IntegrationDeployer.t.sol";
import "src/test/integration/IntegrationChecks.t.sol";
import "src/test/integration/mocks/BeaconChainMock_Deneb.t.sol";

abstract contract UpgradeTest is IntegrationCheckUtils {
    /// Only run upgrade tests on mainnet forks
    function setUp() public virtual override {
        if (!isForktest()) {
            cheats.skip(true);
        } else {
            isUpgraded = false;
            super.setUp();

            // Use Deneb Beacon Chain Mock as Pectra state is not live on mainnet
            beaconChain = BeaconChainMock(new BeaconChainMock_DenebForkable(eigenPodManager, BEACON_GENESIS_TIME));
        }
    }

    /// Deploy current implementation contracts and upgrade existing proxies
    function _upgradeEigenLayerContracts() public virtual {
        require(forkType == MAINNET, "_upgradeEigenLayerContracts: somehow running upgrade test locally");
        require(!isUpgraded, "_upgradeEigenLayerContracts: already performed upgrade");

        emit log("_upgradeEigenLayerContracts: upgrading mainnet to slashing");

        _upgradeMainnetContracts();
        _handlePectraFork();

        // Bump block.timestamp forward to allow verifyWC proofs for migrated pods
        emit log("advancing block time to start of next epoch:");

        beaconChain.advanceEpoch_NoRewards();

        emit log("======");

        isUpgraded = true;
        emit log("_upgradeEigenLayerContracts: slashing upgrade complete");
    }

    // Set the fork timestamp sufficiently in the future to keep using Deneb proofs
    // `Prooftra.t.sol` will handle the Deneb -> Pectra transition
    function _handlePectraFork() internal {
        // 1. Set proof timestamp setter to operations multisig
        cheats.prank(eigenPodManager.owner());
        eigenPodManager.setProofTimestampSetter(address(operationsMultisig));

        // 2. Set Proof timestamp
        cheats.prank(eigenPodManager.proofTimestampSetter());
        eigenPodManager.setPectraForkTimestamp(type(uint64).max);
    }
}
