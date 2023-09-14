# BLSOperatorStateRetriever

This contract is deployed once and is intended to be used by all AVSs that use the provided registry contracts. It is used as a utility for offchain AVS actors to fetch the details of AVSs operators, their stakes in the StakeRegistry, and their indices in the IndexRegistry.

## Flows

This contract is a bunch of view functions that are very gas expensive, they are meant to only be called by offchain actors.

### getOperatorState

This gets the ordered list of operators and their stakes for the provided quorum numbers for the AVS with the provided registry coordinator at the provided block number

There is an overloaded version of this function that takes an operator id and uses the quorum numbers they were registered for instead of having the caller provide quorum numbers.

### getCheckSignaturesIndices

This function is particularly called by AVS aggregators, who get BLS signatures from AVS operators to confirm their signature onchain. Using the provided block number, registry coordinator, quorum numbers, and non signing operator ids, it returns the correct various indices required to confirm signatures with the [BLSSignatureChecker](./BLSSignatureChecker.md).

## Upstream Dependencies

Again, this is called by offchain actors during requests from AVS offchain actors. For example, in EigenDA, a disperser sends data to operators, and those operators are expected to call functions in this contract to make sure they have received the correct amount of data, among other things.