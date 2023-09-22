# BLSSignatureChecker

This contract is deployed per AVS. It verifies the signatures of operators in an efficient way given the rest of the registry architecture. A lot of EigenLayer AVSs can be summarized as a quorum signature on a message and slashing if some quality of that message and other state is true.

## Flows

### checkSignatures
```
function checkSignatures(
        bytes32 msgHash, 
        bytes calldata quorumNumbers,
        uint32 referenceBlockNumber, 
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    )

struct NonSignerStakesAndSignature {
        uint32[] nonSignerQuorumBitmapIndices;
        BN254.G1Point[] nonSignerPubkeys;
        BN254.G1Point[] quorumApks;
        BN254.G2Point apkG2;
        BN254.G1Point sigma;
        uint32[] quorumApkIndices;
        uint32[] totalStakeIndices;  
        uint32[][] nonSignerStakeIndices; // nonSignerStakeIndices[quorumNumberIndex][nonSignerIndex]
    }
```
This function is called by a AVS aggregator when confirming that a `msgHash` is signed by certain `quorumNumbers`.

The function calculates the sum (`apk`) of the aggregate public keys for all the quorums in question using the provided `quorumApkIndices` which point to hashes of the quorum aggregate public keys at the `referenceBlockNumber` of the signature of which `quorumApks` are the preimages. Since there may be nodes from the quorums that don't sign, the `nonSignerPubkeys` are subtracted from `apk`. There is a detail here that since an operator may serve more than one of the quorums in question, the number of quorums in `quorumNumbers` that the nonsigner served at the `referenceBlockNumber` is calculated using the provided `nonSignerQuorumBitmapIndices` and their public key is multiplied by the number before it is subtracted from `apk`. This gets rid of the duplicate additions of their public key because their public key is in more than one of the added `quorumApks`. Now the contract has `apk` set to the claimed aggregate public key of the signers.

Next, the contract fetches the total stakes of each of the `quorumNumbers` at the `referenceBlockNumber` using the provided `totalStakeIndices`. The stakes of each of the nonsigners for each of the quorums are fetched using `nonSignerStakeIndices` and subtracted from the total stakes. Now the contract has the claimed signing stake for each of the quorums.

Finally, the contract does a similar check to the [BLSPublicKeyCompendium](./BLSPublicKeyCompendium.md):

- Calculates $\gamma = keccak256(apk, apkG2, sigma)$
- Verifies the paring $e(\sigma + \gamma apk, [1]_2) = e(H(msgHash) + \gamma[1]_1, apkG2)$

More detailed notes exist on the signature check [here](https://geometry.xyz/notebook/Optimized-BLS-multisignatures-on-EVM).

If it checks out, the contract returns the stake that signed the message for each quorum and the hash of the reference block number and the list of public key hashes of the nonsigners for future use. 

## Upstream Dependencies

AVSs are expected to use this contract's method for their specific tasks. For example, EigenDA uses this function in their contracts when confirming batches of blobs on their DA layer onchain.
