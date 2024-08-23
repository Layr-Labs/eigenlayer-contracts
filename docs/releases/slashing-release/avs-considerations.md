# AVS Considerations

The operatorSet release solves 4 painpoints for AVSs:

1. Quorums cannot be displayed on the UI with their proper TVL and strategies
2. MetaAVSs are not supported
3. Quorums cannot properly be rewarded for the in the current rewards released
4. Configurable operator commissions per AVS/operatorSet

### Quorum (OperatorSet) Display

AVSs can now register operatorSets and their associated strategies in the AVSDirectory. These associations are emitted via events on the AVSDirectory contract.

### MetaAVS Support

MetaAVSs currently have to create a new AVS contract per "meta"-instance. The operatorSet release solves for this by enabling AVSs to create an operatorSet per "meta"-instance.

MetaAVSs are supported by the following metadata URI spec:

```plaintext
{
  "name": "EigenDA",
  "website": "https://eigenda.xyz/",
  "description": "Some description about eigenda",
  "logo": "http://github.com/logo.png",
  "twitter": "https://twitter.com/eigenda"
  "operatorSets": [
        {
          "name":"ETH Set",
          "id":"1",
          "description":"The ETH quorum for EigenDA",
          "software":[
              {
              "name": "NetworkMonitor",
              "description": "",
              "url": "https://link-to-binary-or-github.com"
              },
              {
              "name": "ValidatorClient",
              "description": "",
              "url": "https://link-to-binary-or-github.com"
              }
            ],
            "strategies": [
                "SWETH Strategy Address",
                "STETH Strategy Address",
                "BeaconETH Strategy Address"
             ]
        },
        {
          "name":"Eigen Set",
          "id":"2",
          "description":"The Eigen quorum for EigenDA",
          "software":[
              {
              "name": "NetworkMonitor",
              "description": "",
              "url": "https://link-to-binary-or-github.com"
              },
              {
              "name": "ValidatorClient",
              "description": "",
              "url": "https://link-to-binary-or-github.com"
              }
            ],
            "strategies": [
                "Eigen Strategy Address",
             ]
        }
      ]
}
```

### OperatorSet Payments

The operatorSet release comes with a revamped rewards interface, where AVSs can make payments not only to all their operators but also their operators in a specific operatorSet. This enables AVSs with overlapping strategies in their operatorSets to accurately make payments to operators. In addition, if different operators are doing different tasks (e.g. the MetaAVS model), then AVSs can pay them differently.

### Configurable Operator Commissions

Several AVSs have noted about the high operational costs associated with operating their AVS and a fixed 10% global commission for operators is not viable for them. The OperatorSets Release will enable configuable operator commissions for each RewardType for each AVS (operatorSet) to allow for more flexibility here.

# Migration

AVSs can migrate operators to operatorSets that have already been registered (via M2 registration) for the AVS. This functionality is added for convenience to the AVS. Operators may only be migrated ONCE to a list of operator set IDs for a given AVS. Operators can unilaterally deregister from the operator set if they disagree with the migration. Once an AVS begins migration, it can no longer use the legacy M2 registration pathway.
