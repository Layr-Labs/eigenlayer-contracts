# v1.6.1 Electra Timing Fix

Fixes a bug on EigenPods regarding partial withdrawals for Electra. Please read our [forum]() for a detailed description of the bug. **No action is needed and no customer funds are at risk.**

## Release Manager

@ypatil12 

## Highlights

🐛 *Bug Fixes*
- Update the `EigenPod.requestWithdrawal` function to ensure that validators are pointed to the pod, matching the behavior of `requestConsolidation`

🔧 *Improvements*
- Update the `EigenPod.verifyWithdrawalCredentials` function to only accept `beaconTimestamps` that are after the `latestCheckpointTimestamp`. This enables the eigenpod state machine to be easier to be reasoned about 

