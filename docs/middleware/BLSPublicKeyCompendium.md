# BLSPublicKeyCompendium

This contract is shared by all AVSs and serves a single place for operators to connect their execution layer address to a bn254 public key. When operators opt into AVSs, the AVS contracts can read from the BLSPublicKeyCompendium in order to get the BLS public key of the operator. This contract also prevents against [rogue key attacks](https://xn--2-umb.com/22/bls-signatures/#rogue-key-attack).

## Flows

There is only one flow for this contract, which is a call from an operator to register their execution layer address with thier bn254 public key.

### Cryptographic Verification

Operators provide the contract with 
- A BLS signature $\sigma \in \mathbb{G}_1$ of $M = (msg.sender, chain.id, \text{"EigenLayer\_BN254\_Pubkey\_Registration"})$
- Their public keys $pk_1 \in \mathbb{G}_1$ and $pk_2 \in \mathbb{G}_2$

The contract then 
- Calculates $\gamma = keccak256(\sigma, pk_1, pk_2, M)$
- Verifies $e(\sigma + \gamma pk_1, [1]_2) = e(H(m) + \gamma[1]_1, pk_2)$

This verifies that the operator owns the secret key corresponding to the public keys and that the $pk_1$ and $pk_2$ have the same discrete logarithm according to their respective curve's generators.

### Integrations

The contract stores a map from the execution layer address to the hash of the $\mathbb{G}_1$ public key and the other way around. This means that the $\mathbb{G}_1$ public key needs to sent in calldata onchain, hashed, and verified when AVSs want to do operations with individual operator public keys (remember to link here).