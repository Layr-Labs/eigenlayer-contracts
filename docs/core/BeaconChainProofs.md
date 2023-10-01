#### `BeaconChainProofs.verifyValidatorFields`

```solidity
 function verifyValidatorFields(
        bytes32 beaconStateRoot,
        bytes32[] calldata validatorFields,
        bytes calldata validatorFieldsProof,
        uint40 validatorIndex
    ) internal
```
Verifies the proof of a provided [validator container](https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator) against the beacon state root.  This proof can be used to verify any field in the validator container.  Below is a diagram that illustrates exactly how the proof is structured relative to the [beacon state object](https://github.com/ethereum/consensus-specs/blob/dev/specs/capella/beacon-chain.md#beaconstate).  

[Verify Validator Fields Proof Structure](../images/Withdrawal Credential Proof.png)


#### `BeaconChainProofs.verifyValidatorBalance`

```solidity
function verifyValidatorBalance(
        bytes32 beaconStateRoot,
        bytes32 balanceRoot,
        bytes calldata validatorBalanceProof,
        uint40 validatorIndex
    ) internal
```
Verifies the proof of a [validator's](https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#validator) balance against the beacon state root.  Validator's balances are stored separately in "balances" field of the [beacon state object](https://github.com/ethereum/consensus-specs/blob/dev/specs/capella/beacon-chain.md#beaconstate), with each entry corresponding to the appropriate validator, based on index.  Below is a diagram that illustrates exactly how the proof is structured relative to the [beacon state object](https://github.com/ethereum/consensus-specs/blob/dev/specs/capella/beacon-chain.md#beaconstate).  

[Verify Validator Fields Proof Structure](../images/Balance Proof.png)

#### `BeaconChainProofs.verifyStateRootAgainstLatestBlockRoot`

```solidity
function verifyStateRootAgainstLatestBlockRoot(
        bytes32 latestBlockRoot,
        bytes32 beaconStateRoot,
        bytes calldata stateRootProof
    ) internal
```
Verifies the proof of a beacon state root against the oracle provded block root.  Every [beacon block](https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#beaconblock) in the beacon state contains the state root corresponding with that block.  Thus to prove anything against a state root, we must first prove the state root against the corresponding oracle block root.

[Verify State Root Proof Structure](../images/staterootproof.png)


#### `BeaconChainProofs.verifyWithdrawal`

```solidity
function verifyWithdrawal(
        bytes32 beaconStateRoot,
        bytes32[] calldata withdrawalFields,
        WithdrawalProof calldata withdrawalProof
    ) internal
```
Verifies a withdrawal, either [full or partial](https://eth2book.info/capella/part2/deposits-withdrawals/withdrawal-processing/#partial-and-full-withdrawals), of a validator.  There are a maximum of 16 withdrawals per block in the consensus layer.  This proof proves the inclusion of a given [withdrawal](https://github.com/ethereum/consensus-specs/blob/dev/specs/capella/beacon-chain.md#withdrawal) in the block for a given slot.  

One important note is that we use [`historical_summaries`](https://github.com/ethereum/consensus-specs/blob/dev/specs/capella/beacon-chain.md#withdrawal) to prove the blocks that contain withdrawals.  Each new [historical summary](https://github.com/ethereum/consensus-specs/blob/dev/specs/capella/beacon-chain.md#historicalsummary) is added every 8192 slots, i.e., if `slot % 8192 = 0`, then `slot.state_roots` and `slot.block_roots` are merkleized and are used to create the latest `historical_summaries` entry.  

[Verify Withdrawal Proof Structure](../images/Withdrawal Proof.png)



