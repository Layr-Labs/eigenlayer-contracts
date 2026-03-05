# Duration Vault E2E (Hoodi)

This folder contains a minimal end-to-end smoke test script intended to be run against live environments (e.g. `preprod-hoodi`).

## What it tests

- **Vault lifecycle**: deploy → delegate+deposit → lock
- **Slashing**: slash the vault (as the AVS slasher) and confirm redistributed funds reach the insurance recipient
- **Rewards plumbing**: create an on-chain AVS rewards submission so the Sidecar can pick it up and later submit a distribution root

## Prereqs

- A Hoodi RPC URL (export `RPC_HOODI=...`)
- A funded EOA private key that will be used as:
  - **vault admin**
  - **AVS address**
  - **operator-set slasher**
  - **staker** (delegates + deposits)

## Run (preprod-hoodi)

### 0) Get core contract addresses

These scripts are normal Foundry scripts; you just need the 5 core addresses for the environment you’re targeting:

- `ALLOCATION_MANAGER`
- `DELEGATION_MANAGER`
- `STRATEGY_MANAGER`
- `STRATEGY_FACTORY`
- `REWARDS_COORDINATOR`

Convenient way to grab them (copy/paste) is still Zeus:

```bash
zeus env show preprod-hoodi
```

Then export the 5 addresses in your shell (or pass them as args via `--sig`).

### 1) Fork/simulate first

Run the script on a fork first (no `--broadcast`).
Tip: pass a funded sender (your EOA) so forked execution doesn’t revert due to `insufficient funds for gas`:

```bash
forge script script/operations/e2e/DurationVaultHoodiE2E.s.sol \
  --fork-url "$RPC_HOODI" \
  --sender 0xYourEOA \
  -vvvv
```

### 2) Broadcast on Hoodi

Then broadcast on Hoodi:

```bash
forge script script/operations/e2e/DurationVaultHoodiE2E.s.sol \
  --rpc-url "$RPC_HOODI" \
  --private-key "$PRIVATE_KEY" \
  --broadcast \
  -vvvv
```

### Re-running without a new EOA

- If you want to re-run **everything from scratch**, keep the same EOA but set a fresh `E2E_OPERATOR_SET_ID` (e.g. `2`, `3`, …) so `createRedistributingOperatorSets` doesn’t revert.
- If you only need a **fresh rewards submission** (e.g. to fix a bad time window), you can reuse the existing deployed vault from `e2e-state.json`:

```bash
export E2E_PHASE=rewards
forge script script/operations/e2e/DurationVaultHoodiE2E.s.sol \
  --rpc-url "$RPC_HOODI" \
  --private-key "$PRIVATE_KEY" \
  --broadcast \
  -vvvv
```

## Optional env overrides

You can override parameters using env vars:

- `E2E_OPERATOR_SET_ID` (uint, default `1`)
- `E2E_INSURANCE_RECIPIENT` (address, default = your EOA)
- `E2E_MAX_PER_DEPOSIT` (uint, default `200 ether`)
- `E2E_STAKE_CAP` (uint, default `1000 ether`)
- `E2E_VAULT_DURATION_SECONDS` (uint, default `120`)
- `E2E_DEPOSIT_AMOUNT` (uint, default `100 ether`)
- `E2E_SLASH_WAD` (uint, default `0.25e18`)
- `E2E_REWARD_AMOUNT` (uint, default `10 ether`)
- `E2E_REWARDS_START_TIMESTAMP` (uint, default = **current** `RewardsCoordinator.CALCULATION_INTERVAL_SECONDS` boundary)
- `E2E_REWARDS_DURATION_SECONDS` (uint, default = `RewardsCoordinator.CALCULATION_INTERVAL_SECONDS` (1 day on `preprod-hoodi`))

## Validating rewards end-to-end (with Sidecar)

This script only creates a `createAVSRewardsSubmission()` entry. To complete the “real” rewards E2E:

- **Sidecar** should index the new submission and compute a distribution root for the relevant time window.
- The configured **RewardsUpdater** (often a Sidecar-controlled key) should call `RewardsCoordinator.submitRoot(...)`.
- Once the root is activated, the earner can call `RewardsCoordinator.processClaim(...)` with the Sidecar-produced proof.

If you tell me which Sidecar instance / pipeline you’re using on `preprod-hoodi`, I can add a short checklist of the exact on-chain events + fields to confirm (submission hash, root index, activation timestamp, claim balance deltas).

## Claiming rewards (optional, after Sidecar posts a root)

Once Sidecar has posted a **claimable** distribution root and you have a claim JSON from your coworker, run the claim-only script:

```bash
forge script script/operations/e2e/ClaimRewards.s.sol \
  --rpc-url "$RPC_HOODI" --private-key "$PRIVATE_KEY" --broadcast \
  --sig "run(string,address,address)" \
  -- "path/to/claim.json" 0xEarner 0xRecipient
```

### Claim JSON format

The claim file must be parseable by Foundry `vm.parseJson` and match the `RewardsMerkleClaim` structure:

```json
{
  "rootIndex": 0,
  "earnerIndex": 0,
  "earnerTreeProof": "0x",
  "earnerLeaf": { "earner": "0x0000000000000000000000000000000000000000", "earnerTokenRoot": "0x..." },
  "tokenIndices": [0],
  "tokenTreeProofs": ["0x"],
  "tokenLeaves": [
    { "token": "0x0000000000000000000000000000000000000000", "cumulativeEarnings": "10000000000000000000" }
  ]
}
```

Notes:
- The script will **revert** unless the root at `rootIndex` exists and `block.timestamp >= activatedAt`.
- The script will **revert** if `earnerLeaf.earner != <earner arg>` to avoid accidental mismatched claims. In practice you should pass the same address as the EOA you’re broadcasting from.


