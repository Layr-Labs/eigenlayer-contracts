# Storage Validation Scripts
This script uses cast and forge inspect to get the layouts of the local and on-chain contract. Your `ETHERSCAN_API_KEY` must be set as an environment variable. The storage layouts are saved into csv files and passed into the typescript helper validateStorage.ts, which takes paths to two csv layouts and validates the storage slots.

## Run validation
To validate the storage of an upgradeable deployed contract against a local one, run the script:
```bash
bash script/upgrade/validateUpgrade.sh -n <network> -c <contract> -a <implementation address>
```

The supported networks are `goerli` and `mainnet`. The supported contracts are `strategyManager`, `delegation`, `eigenPod`, `eigenPodManager`, and `slasher`.

The above script generates two csv files, `localLayout.csv` and `onChainLayout.csv`. To keep these csv files after validating storage, add a `-k` flag to the above command

Additionally, one can validate the storage of two csv files outputted by the `forge inspect` command by running 

```js
npx ts-node script/upgrade/validateStorage.ts --old <path_to_old_layout_csv> --new <path_to_new_layout_csv> --keep
```

## Limitations
Storage slot validation is NOT comprehensive, and errs on the side of caution. We recommend using this script as a tool along with manual storage slot verification. The validation is opinionated on storage for each contract consuming 50 slots and gaps being sized accordingly.

The script does not validate legal type changes (ie. from bool to uint8) and errors out if the types of slots have updated, including having different struct names. A manual check will need to be done to validate this conversion. In addition, the script does not support non-contiguous gaps.