# Metadata URI Standard

Below is the new metadataURI standard for AVSs that call `setMetadataURI` in the `AllocationManager`

```plaintext
 {
  "name": "AVS",
  "website": "https.avs.xyz/",
  "description": "Some description about",
  "logo": "http://github.com/logo.png",
  "twitter": "https://twitter.com/avs"
  "operatorSets": [
        {
          "name":"ETH Set",
          "id":"1", // Note: we use this param to match the opSet id in the Allocation Manager
          "description":"The ETH operatorSet for AVS",
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
            "slashingConditions: ["Condition A", "Condition B"]
        },
        {
          "name":"EIGEN Set",
          "id":"2", // Note: we use this param to match the opSet id in the Allocation Manager
          "description":"The EIGEN operatorSet for AVS",
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
            "slashingConditions: ["Condition A", "Condition B"]
        }
   ]
}
```