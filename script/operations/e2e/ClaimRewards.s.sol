// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Script.sol";

import "src/contracts/core/RewardsCoordinator.sol";

/// @notice Claim RewardsCoordinator rewards using an externally-provided claim JSON.
///
/// Intended usage:
/// - Your coworker/Sidecar produces a claim JSON (already in the correct shape/encoding).
/// - You run this script against preprod-hoodi (or any Zeus env) to call `processClaim`.
///
/// Example:
/// forge script script/operations/e2e/ClaimRewards.s.sol \
///   --rpc-url "$RPC_HOODI" --private-key "$PRIVATE_KEY" --broadcast \
///   --sig "run(string,address,address)" \
///   -- "path/to/claim.json" 0xEarner 0xRecipient
contract ClaimRewards is Script {
    /// @notice Claim using the RewardsCoordinator address from env var `REWARDS_COORDINATOR`.
    function run(
        string memory claimFile,
        address earner,
        address recipient
    ) public {
        address rc = vm.envAddress("REWARDS_COORDINATOR");
        vm.startBroadcast();
        _claim(RewardsCoordinator(rc), claimFile, earner, recipient);
        vm.stopBroadcast();
    }

    /// @notice Claim using an explicitly provided RewardsCoordinator address.
    function runWithRewardsCoordinator(
        address rewardsCoordinator,
        string memory claimFile,
        address earner,
        address recipient
    ) public {
        vm.startBroadcast();
        _claim(RewardsCoordinator(rewardsCoordinator), claimFile, earner, recipient);
        vm.stopBroadcast();
    }

    function _claim(
        RewardsCoordinator rc,
        string memory claimFile,
        address earner,
        address recipient
    ) internal {
        IRewardsCoordinatorTypes.RewardsMerkleClaim memory claim = _loadClaimFromFile(claimFile);
        require(claim.earnerLeaf.earner == earner, "claim earner mismatch");
        rc.processClaim(claim, recipient);
    }

    function _loadClaimFromFile(
        string memory claimPath
    ) internal view returns (IRewardsCoordinatorTypes.RewardsMerkleClaim memory claim) {
        string memory fullPath = _resolvePath(claimPath);
        string memory json = vm.readFile(fullPath);

        // Accept either:
        // - direct object: { rootIndex, earnerIndex, earnerTreeProof, earnerLeaf, tokenIndices, tokenTreeProofs, tokenLeaves }
        // - wrapped: { proof: { ... } } (Sidecar-style)
        if (_hasJsonPath(json, ".proof")) {
            claim.rootIndex = uint32(abi.decode(vm.parseJson(json, ".proof.rootIndex"), (uint256)));
            claim.earnerIndex = uint32(abi.decode(vm.parseJson(json, ".proof.earnerIndex"), (uint256)));
            claim.earnerTreeProof = abi.decode(vm.parseJson(json, ".proof.earnerTreeProof"), (bytes));
            claim.earnerLeaf =
                abi.decode(vm.parseJson(json, ".proof.earnerLeaf"), (IRewardsCoordinatorTypes.EarnerTreeMerkleLeaf));

            // If the proof contains scalar token fields, wrap them into single-element arrays.
            claim.tokenIndices = new uint32[](1);
            claim.tokenIndices[0] = uint32(abi.decode(vm.parseJson(json, ".proof.tokenIndices"), (uint256)));
            claim.tokenTreeProofs = new bytes[](1);
            claim.tokenTreeProofs[0] = abi.decode(vm.parseJson(json, ".proof.tokenTreeProofs"), (bytes));
            claim.tokenLeaves = new IRewardsCoordinatorTypes.TokenTreeMerkleLeaf[](1);
            claim.tokenLeaves[0] =
                abi.decode(vm.parseJson(json, ".proof.tokenLeaves"), (IRewardsCoordinatorTypes.TokenTreeMerkleLeaf));
            return claim;
        }

        claim.rootIndex = uint32(abi.decode(vm.parseJson(json, ".rootIndex"), (uint256)));
        claim.earnerIndex = uint32(abi.decode(vm.parseJson(json, ".earnerIndex"), (uint256)));
        claim.earnerTreeProof = abi.decode(vm.parseJson(json, ".earnerTreeProof"), (bytes));
        claim.earnerLeaf =
            abi.decode(vm.parseJson(json, ".earnerLeaf"), (IRewardsCoordinatorTypes.EarnerTreeMerkleLeaf));
        claim.tokenIndices = abi.decode(vm.parseJson(json, ".tokenIndices"), (uint32[]));
        claim.tokenTreeProofs = abi.decode(vm.parseJson(json, ".tokenTreeProofs"), (bytes[]));
        claim.tokenLeaves =
            abi.decode(vm.parseJson(json, ".tokenLeaves"), (IRewardsCoordinatorTypes.TokenTreeMerkleLeaf[]));
    }

    function _hasJsonPath(
        string memory json,
        string memory path
    ) internal pure returns (bool) {
        try vm.parseJson(json, path) returns (bytes memory raw) {
            return raw.length != 0;
        } catch {
            return false;
        }
    }

    function _resolvePath(
        string memory path
    ) internal view returns (string memory) {
        bytes memory b = bytes(path);
        if (b.length > 0 && b[0] == bytes1("/")) return path;
        return string.concat(vm.projectRoot(), "/", path);
    }
}

