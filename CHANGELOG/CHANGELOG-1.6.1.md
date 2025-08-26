# v1.6.1 Moocow Updates

## Release Manager

@ypatil12 

## Highlights

🔧 **Improvements**
- Make `EigenPod.validatorStatus` public to use internally as a helper
- Update the `EigenPod.verifyWithdrawalCredentials` function to only accept `beaconTimestamps` that are after the `latestCheckpointTimestamp`. This enables the eigenpod state machine to be easier to be reasoned about 
- Update the `EigenPod.requestWithdrawal` function to ensure that validators are pointed to the pod, matching the behavior of `requestConsolidation`

## Changelog

TODO