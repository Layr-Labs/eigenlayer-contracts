# BLSPublicKeyCompendium

This contract is shared by all AVSs and serves as a single place for operators to connect their execution layer address to a bn254 public key. This contract also prevents [rogue key attacks](https://xn--2-umb.com/22/bls-signatures/#rogue-key-attack).

## Flows

There is only one flow for this contract, which is a call from an operator to register a bn254 public key as controlled by their execution layer address.

### Cryptographic Verification

Operators provide the contract with 
- A BLS signature $\sigma \in \mathbb{G}_1$ of $M = (msg.sender, chain.id, \text{"EigenLayer\_BN254\_Pubkey\_Registration"})$
- Their public keys $pk_1 \in \mathbb{G}_1$ and $pk_2 \in \mathbb{G}_2$

The contract then 
- Calculates $\gamma = keccak256(\sigma, pk_1, pk_2, M)$
- Verifies the paring $e(\sigma + \gamma pk_1, [1]_2) = e(H(m) + \gamma[1]_1, pk_2)$

This verifies that the operator owns the secret key corresponding to the public keys and that the $pk_1$ and $pk_2$ have the same discrete logarithm according to their respective curve's generators. 

We do this particular verification because aggregation of public keys and hasing to the curve is cheap in $\mathbb{G}_1$ on ethereum, and the above scheme allows for both! (aggregation to be done in the [BLSSignatureChecker](./BLSSignatureChecker.md)) More detailed notes exist [here](https://geometry.xyz/notebook/Optimized-BLS-multisignatures-on-EVM).

The contract then stores a map from the execution layer address to the hash of the operator's $\mathbb{G}_1$ public key and the other way around.

### Upstream Dependencies

The [BLSPubkeyRegistry](./BLSPubkeyRegistry.md) looks up the public key hashes in this contract when operators register with a certain public key.